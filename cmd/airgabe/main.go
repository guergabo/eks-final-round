package main

import (
	"fmt"
	"os"

	"github.com/guergabo/eks-final-round/internal/core/services/airgabesrv"
	"github.com/guergabo/eks-final-round/internal/handlers/airgabehdl"
	"github.com/guergabo/eks-final-round/internal/repositories/airgaberepo"
)

func main() {
	// wire
	airplaneRepository := airgaberepo.NewLocalFile()
	airplaneService := airgabesrv.New(airplaneRepository)
	airplaneHandler := airgabehdl.NewCLHandler(airplaneService)

	// collect command line arguments - pkg
	argsWithoutProg := os.Args[1:]

	// parse arguments and form request
	fmt.Println(airplaneHandler.Run(argsWithoutProg).Status)
}
