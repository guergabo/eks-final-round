package main

import (
	"fmt"

	"github.com/guergabo/eks-final-round/internal/core/services/airgabesrv"
	"github.com/guergabo/eks-final-round/internal/handlers/airgabehdl"
	"github.com/guergabo/eks-final-round/internal/repositories/airgaberepo"
	"github.com/guergabo/eks-final-round/pkg/utils"
)

func main() {
	// wire
	airplaneRepository := airgaberepo.NewLocalFile()
	airplaneService := airgabesrv.New(airplaneRepository)
	airplaneHandler := airgabehdl.NewCLHandler(airplaneService)

	// collect command line arguments - pkg
	argsWithoutProg := utils.GetCLIArgs()

	// parse arguments and form request
	fmt.Println(airplaneHandler.Run(argsWithoutProg).Status)
}
