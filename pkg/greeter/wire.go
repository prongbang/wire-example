//+build wireinject

package greeter

import (
	"github.com/google/wire"
)

func Initialize() (Greeter, error) {
	wire.Build(NewGreeter, NewUseCase)
	return &greeter{}, nil
}
