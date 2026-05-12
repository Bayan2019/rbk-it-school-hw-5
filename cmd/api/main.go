package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Bayan2019/rbk-it-school-hw-5/internal/auth"
	"github.com/Bayan2019/rbk-it-school-hw-5/internal/client"
	"github.com/Bayan2019/rbk-it-school-hw-5/internal/config"
	"github.com/Bayan2019/rbk-it-school-hw-5/internal/handler"
	"github.com/Bayan2019/rbk-it-school-hw-5/internal/model"
	"github.com/Bayan2019/rbk-it-school-hw-5/internal/repository/postgres"
	"github.com/Bayan2019/rbk-it-school-hw-5/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	middle "github.com/Bayan2019/rbk-it-school-hw-5/internal/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("warning: assuming default configuration: .env unreadable: %v\n", err)
	}

	cfg := config.MustLoad()

	db, err := postgres.NewDB(cfg.Database)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	cityRepo := postgres.NewCityRepository(db)
	weatherRepo := postgres.NewWeatherRepository(db)

	osmClient := client.NewOsmClient(cfg.Api)
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	weatherClient := client.NewWeatherClient(httpClient)

	userService := service.NewUserService(userRepo)
	cityService := service.NewCityService(cityRepo, osmClient)
	weatherService := service.NewWeatherService(weatherRepo, weatherClient)

	jwtManager := auth.NewJWTManager([]byte(cfg.App.JwtSecret))

	userHandler := handler.NewUserHandler(userService, jwtManager)
	cityHandler := handler.NewCityHandler(cityService)
	weatherHandler := handler.NewWeatherHandler(cityService, weatherService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	// 1. Аутентификация
	r.Post("/auth/register", userHandler.Register)
	r.Post("/auth/login", userHandler.Login)

	// 4. Защита маршрутов
	r.Group(func(r chi.Router) {
		// Все операции должны работать через текущего пользователя из JWT.
		r.Use(middle.AuthMiddleware(userHandler.JwtManager))
		// Убрать user_id из URL.
		r.Post("/cities", cityHandler.Add2User)
		r.Get("/cities", cityHandler.ListOfUser)
		r.Delete("/cities/{city_id}", cityHandler.DeleteFromUser)
		r.Get("/weather", weatherHandler.GetWeatherOfUserCities)
		r.Get("/weather/history", weatherHandler.GetWeatherHistoryOfUser)

		// 8. Новый endpoint
		r.Get("/users/me", userHandler.Profile)
		// 4. Защита маршрутов
		// Убрать user_id из URL.
		// Все операции должны работать через текущего пользователя из JWT.
		r.Put("/users/me", userHandler.Update)

		// 5. Авторизация (Roles)
		r.Group(func(r chi.Router) {
			// Использовать middleware RequireRole("admin")
			r.Use(middle.RequireRole(model.RolesAdmin))
			// Только admin может:
			r.Get("/users", userHandler.List)
			r.Get("/users/{id}", userHandler.GetByID)
			r.Delete("/users/{id}", userHandler.Delete)
		})
	})

	log.Printf("server started on %s", cfg.App.Port)
	if err := http.ListenAndServe(cfg.App.Port, r); err != nil {
		log.Fatal(err)
	}
}
