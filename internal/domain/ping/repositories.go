package ping

import (
	"gorm.io/gorm"
)

type PingRepository interface {
	GetPing() (*Ping, error)
}

type PingRepositoryImpl struct {
	db *gorm.DB
}

func NewPingRepository(db *gorm.DB) *PingRepositoryImpl {
	return &PingRepositoryImpl{db: db}
}

func (r *PingRepositoryImpl) GetPing() (*Ping, error) {
	// Aquí puedes implementar la lógica para obtener datos de la base de datos
	return &Ping{Message: "pong"}, nil
}
