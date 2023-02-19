package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/r4f3t/webapi/internal/config"
	inMemUser "github.com/r4f3t/webapi/internal/inMemoryCache/user"
	inMemUserController "github.com/r4f3t/webapi/internal/inMemoryCache/user/controller"
	"github.com/spf13/cobra"
)

type api struct {
	instance   *echo.Echo
	command    *cobra.Command
	Port       string
	DbSettings config.DbSettings
}

// apiCmd represents the api command
var apiCmd = &api{
	command: &cobra.Command{
		Use:   "api",
		Short: "",
		Long:  "",
	},
	Port: "5000",
}

func init() {
	RootCommand.AddCommand(apiCmd.command)
	//defines flags for api command specific
	apiCmd.command.Flags().StringVarP(&apiCmd.Port, "port", "p", "5000", "Service Port")
	apiCmd.instance = echo.New()

	apiCmd.command.RunE = func(cmd *cobra.Command, args []string) error {

		// construct repository
		inMemUserRepository := inMemUser.NewRepository()

		// construct user cache manager
		inMemUserCacheManager := inMemUser.NewUserCacheManager(inMemUserRepository)

		// construct service
		inMemUserService := inMemUser.NewService(inMemUserCacheManager)

		inMemUserController.MakeHandler(apiCmd.instance, inMemUserController.NewController(inMemUserService))

		apiCmd.instance.Logger.Fatal(apiCmd.instance.Start(":8080"))

		return nil
	}

}
