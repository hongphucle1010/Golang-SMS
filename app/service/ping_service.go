package service

type IPingService interface {
	Pong() (string, error)
}

type PingService struct {
}

func (s *PingService) Pong() (string, error) {
	return "pong", nil
}