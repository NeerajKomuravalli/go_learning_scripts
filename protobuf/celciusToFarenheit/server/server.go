package server

import (
	"context"

	"github.com/NeerajKomuravalli/go_learning_scripts/protobuf/celciusToFarenheit/proto/temperature"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) ConvertFromCelsiusToFarenheit(ctx context.Context, temp *temperature.Temprature) (*temperature.Temprature, error) {
	return &temperature.Temprature{Temprature: ((1.8 * temp.Temprature) + 32)}, nil
}

func (server *Server) ConvertFromFarenheitToCelsius(ctx context.Context, temp *temperature.Temprature) (*temperature.Temprature, error) {
	return &temperature.Temprature{Temprature: (temp.Temprature - 30) / 2}, nil
}
