package greeter

type usecase struct {
}

type UseCase interface {
	GetAll() []string
}

func NewUseCase() UseCase {
	return &usecase{}
}

func (u *usecase) GetAll() []string {
	return []string{
		"greeter-1", "greeter-2", "greeter-3", "greeter-4", "greeter-5",
	}
}
