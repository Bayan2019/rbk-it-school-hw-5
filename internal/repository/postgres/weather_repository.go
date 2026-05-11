package postgres

import (
	"context"
	"errors"
	"strings"

	"github.com/Bayan2019/rbk-it-school-hw-5/internal/dto"
	"github.com/jmoiron/sqlx"
)

type WeatherRepository struct {
	db *sqlx.DB
}

func NewWeatherRepository(db *sqlx.DB) *WeatherRepository {
	return &WeatherRepository{db: db}
}

////// methods
////// methods
////// methods

func (r *WeatherRepository) CreateHistory(ctx context.Context, userID int64, cityWeather dto.CityWeatherInput) (dto.WeatherHistoryResponse, error) {

	query := `
		INSERT INTO weather_history (user_id, city, temperature, description)
		VALUES (:user_id, :city, :temperature, :description)
		RETURNING city, temperature, description, requested_at
	`

	args := map[string]any{
		"user_id":     userID,
		"city":        cityWeather.City,
		"temperature": cityWeather.Temperature,
		"description": cityWeather.Description,
	}

	rows, err := r.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return dto.WeatherHistoryResponse{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var result dto.WeatherHistoryResponse
		if err := rows.StructScan(&result); err != nil {
			return dto.WeatherHistoryResponse{}, err
		}
		return result, nil
	}

	return dto.WeatherHistoryResponse{}, nil
}

func (r *WeatherRepository) WeatherHistoryOfUser(ctx context.Context, userID int64, filter dto.WeatherHistoryFilter) ([]dto.WeatherHistoryResponse, error) {
	builder := strings.Builder{}
	builder.WriteString(`
		SELECT user_id, city, temperature, description, requested_at
		FROM weather_history
		WHERE user_id = :user_id
	`)

	args := map[string]any{
		"user_id": userID,
	}

	if filter.City != "" {
		builder.WriteString(" AND city = :city")
		args["city"] = filter.City
	}

	builder.WriteString(" ORDER BY requested_at DESC")

	if filter.Limit != 0 {
		builder.WriteString(" LIMIT :limit")
		args["limit"] = filter.Limit
	}

	if filter.Offset != 0 {
		builder.WriteString(" OFFSET :offset")
		args["offset"] = filter.Offset
	}

	query, queryArgs, err := sqlx.Named(builder.String(), args)
	if err != nil {
		return nil, errors.New("sqlx.Named")
	}
	query = r.db.Rebind(query)

	var results []dto.WeatherHistoryResponse
	if err := r.db.SelectContext(ctx, &results, query, queryArgs...); err != nil {
		return nil, errors.New("r.db.SelectContext")
	}

	return results, nil
}

////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
////// accommodating functions
