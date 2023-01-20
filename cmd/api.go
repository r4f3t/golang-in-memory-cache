package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/r4f3t/webapi/internal/config"
	inMemUser "github.com/r4f3t/webapi/internal/inMemoryCache/user"
	inMemUserController "github.com/r4f3t/webapi/internal/inMemoryCache/user/controller"
	"github.com/r4f3t/webapi/internal/singletonObjectCache/user"
	"github.com/r4f3t/webapi/internal/singletonObjectCache/user/controller"
	"github.com/r4f3t/webapi/mockdata"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

		var dto []mockdata.UserModel
		db := ConnectDB("localhost", "postgres", "changeme", "postgres", "5432")
		db.Raw("SELECT * FROM public.\"Users\"").Scan(&dto)
		// construct repository
		userRepository := user.NewRepository()

		// construct repository
		inMemUserRepository := inMemUser.NewRepository()

		// construct user cache manager
		userCacheManager := user.NewUserCacheManager()

		// construct user cache manager
		inMemUserCacheManager := inMemUser.NewUserCacheManager()

		// construct service
		userService := user.NewService(userRepository, userCacheManager)

		// construct service
		inMemUserService := inMemUser.NewService(inMemUserRepository, inMemUserCacheManager)

		controller.MakeHandler(apiCmd.instance, controller.NewController(userService))

		inMemUserController.MakeHandler(apiCmd.instance, inMemUserController.NewController(inMemUserService))

		apiCmd.instance.Logger.Fatal(apiCmd.instance.Start(":8080"))

		return nil
	}

}

func ConnectDB(host, user, password, dbname, port string) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
	return db
}
