package grpc

type RocketService struct {
}

type Handler struct {
	RocketService RocketService
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}
