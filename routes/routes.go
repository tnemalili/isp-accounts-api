package routes

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"isp.accounts.api/controllers"
	"os"
)

func ConfigureRouter() error {

	version := fmt.Sprintf("/api/%s/", os.Getenv("API_VERSION"))
	port := os.Getenv("API_PORT")
	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{ echo.POST,  echo.GET, echo.PUT,  echo.OPTIONS, echo.DELETE },
		AllowOrigins: []string{"*"},
	}))
	// DEBUG MODE (OPTIONAL)
	//router.Debug = true

	// JUST TO HIDE THE ECHO BANNER
	router.HideBanner = true

	// ADDING TRAILING SLASH TO REQUEST URI
	// router.Pre(middleware.AddTrailingSlash())

	api := router.Group(version)

	api.GET("heartbeat", controllers.HealthCheck)

	////////////////////////// ACCOUNTS SECTION ///////////////////////////////////

	api.POST("accounts", controllers.CREATEAccount)
	api.GET("accounts/:id", controllers.FETCHAccount)
	api.GET("accounts", controllers.FETCHAccounts)
	api.PUT("accounts/:id", controllers.UPDATEAccount)
	api.PUT("fund/:id", controllers.FUNDAccount)

	data, err := json.MarshalIndent(router.Routes(), "", "  ")
	if err != nil {
		return err
	}
	// LOG ALL REGISTERED ROUTES
	log.Info(string(data))
	log.Info("[ISP.ACCOUNTS.API STARTED!]")
	log.Info(router.Start(fmt.Sprintf(":%s", port)))
	return nil
}
