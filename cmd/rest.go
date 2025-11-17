package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	"github.com/nugrohoac/e-commerce/transport/rest/auth"
	prodRest "github.com/nugrohoac/e-commerce/transport/rest/product"
)

var cmdRest = &cobra.Command{
	Use:   "rest",
	Short: "start rest server",
	Run: func(cmd *cobra.Command, args []string) {
		echoInstance := echo.New()

		auth.RegisterHandler(echoInstance, authService)
		prodRest.RegisterHandler(echoInstance, productService)

		log.Infof("Starting HTTP Server at %v", cfg.Service.Port.REST)

		if err := echoInstance.Start(cfg.Service.Port.REST); err != nil {
			log.Fatalf("Failed to start server : %v", err)
		}
	},
}
