package cmd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	prodSvc "github.com/nugrohoac/e-commerce/application/service/product"
	"github.com/nugrohoac/e-commerce/infrastructure/repository/product"
	"github.com/spf13/cobra"

	"github.com/nugrohoac/e-commerce/application/service/auth"
	"github.com/nugrohoac/e-commerce/infrastructure/repository/user"
	"github.com/nugrohoac/e-commerce/resource/config"
)

func init() {
	cobra.OnInitialize(initConfig)
	RootCMD.AddCommand(cmdRest)
}

var (
	RootCMD = &cobra.Command{
		Use:   "ecommerce",
		Short: "ecommerce service",
		Long:  "",
	}

	// repository
	userRepository    user.Repository
	productRepository product.Repository

	// service
	authService    auth.Service
	productService prodSvc.Service

	// config
	cfg  *config.Configuration
	cred *config.Credential
	err  error
)

func initConfig() {
	cfg, err = config.NewConfiguration("configuration.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cred, err = config.NewCredential("credential.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cred.Database.User,
		cred.Database.Password,
		cred.Database.Host,
		cred.Database.Port,
		cred.Database.Name,
	)

	dbConnection, err := sql.Open(cred.Database.Driver, dataSource)
	if err != nil {
		log.Fatal()
	}

	// repository
	userRepository = user.NewRepository(dbConnection)
	productRepository = product.NewRepository(dbConnection)

	// service
	authService = auth.NewService(userRepository)
	productService = prodSvc.NewService(productRepository)
}
