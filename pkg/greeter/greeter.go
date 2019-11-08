package greeter

type greeter struct {
	UseCase UseCase
}

type Greeter interface {
	GetAll() []string
}

func NewGreeter(useCase UseCase) Greeter {
	return &greeter{
		UseCase: useCase,
	}
}

func (g *greeter) GetAll() []string {
	return g.UseCase.GetAll()
}
