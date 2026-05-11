package service

import (
	"context"

	"github.com/Bayan2019/rbk-it-school-hw-5/internal/dto"
	"github.com/Bayan2019/rbk-it-school-hw-5/internal/model"
)

type weatherProvider interface {
	GetCurrentWeather(ctx context.Context, lat, lon float64) (dto.ProviderWeatherResponse, error)
}

type weatherRepository interface {
	CreateHistory(ctx context.Context, userID int64, cityWeather dto.CityWeatherInput) error
	WeatherHistoryOfUser(ctx context.Context, userID int64, filter dto.WeatherHistoryFilter) ([]dto.WeatherHistoryResponse, error)
}

type WeatherService struct {
	repo     weatherRepository
	provider weatherProvider
}

func NewWeatherService(repo weatherRepository, provider weatherProvider) *WeatherService {
	return &WeatherService{
		repo:     repo,
		provider: provider,
	}
}

////// methods
////// methods
////// methods

func (s *WeatherService) CreateHistory(ctx context.Context, userID int64, city model.City) error {

	res, err := s.provider.GetCurrentWeather(ctx, city.Lat, city.Lon)
	if err != nil {
		// h.handleError(w, err)
		return err
	}
	// if err := cityWeather.NormalizeAndValidate(); err != nil {
	// 	return dto.WeatherHistoryResponse{}, err
	// }
	return s.repo.CreateHistory(ctx, userID, dto.CityWeatherInput{
		City:        city.City,
		Temperature: res.Temperature,
		Description: res.Description,
	})
}

func (s *WeatherService) WeatherHistoryOfUser(ctx context.Context, userID int64, filter dto.WeatherHistoryFilter) ([]dto.WeatherHistoryResponse, error) {

	filter.Normalize()
	return s.repo.WeatherHistoryOfUser(ctx, userID, filter)
}
