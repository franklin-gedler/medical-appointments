package controllers

import (
	"errors"
	"medical-appointments/internal/domain/ping"
	"net/http"

	"go.uber.org/zap"
)

type PingController struct {
	Service *ping.PingService
}

func NewPingController(service *ping.PingService) *PingController {
	return &PingController{
		Service: service,
	}
}

func (p *PingController) Ping(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log messages
		logger.Info("Mensaje de info")
		logger.Debug("Mensaje de depuración")
		logger.Error("Mensaje de error")
		logger.Warn("Mensaje de advertencia")

		// Log messages with fields
		logger.Info("Mensaje de info con tags", zap.Bool("Booleano", false), zap.Int("Entero", 2))
		logger.Debug("Mensaje de depuración con tags", zap.String("Cadena", "Cadena"), zap.Float64("Flotante", 3.14))
		err_test := errors.New("Cualquier mensaje de error")
		logger.Error(err_test.Error(), zap.Error(err_test))

		_, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	}
}
