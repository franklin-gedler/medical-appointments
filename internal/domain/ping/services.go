package ping

type PingService struct {
	repo PingRepository
}

func NewPingService(repo PingRepository) *PingService {
	return &PingService{repo: repo}
}

func (s *PingService) GetPing() (*Ping, error) {
	return s.repo.GetPing()
}
