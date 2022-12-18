package main

import (
    "fmt"
    "os"
    "encoding/json"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/spf13/viper"
    "github.com/onosproject/onos-lib-go/pkg/logging"
    "github.com/amitinfo2k/k8s-examples/open-api/config"
    "github.com/amitinfo2k/k8s-examples/open-api/gorm"
    "github.com/amitinfo2k/k8s-examples/open-api/pkg/model"
)

var log = logging.GetLogger("model_0_0_0")

func main() {
    conf := loadConfig()
    db, err := gorm.Connect(conf)
    if err != nil {
        log.Panic("failed to connect database")
    }
    log.Infof("Successfully connected to database")
    gorm.RunMigration(db)
	    
    var s = model.AetherServer{ DB:db }
    // Echo instance
    echoRouter := echo.New()

    // Middleware
    echoRouter.Use(middleware.Logger())
    echoRouter.Use(middleware.Recover())

    model.RegisterHandlers(echoRouter, &s)

    // Start server
    echoRouter.Logger.Fatal(echoRouter.Start(":8080"))
}

func overrideUsingEnvVars(config *config.Config) {
    if host, present := os.LookupEnv("DB_HOST"); present {
        config.Database.Host = host
    }
}

func loadConfig() *config.Config {
    env := "local"
    if v, present := os.LookupEnv("ENV"); present {
        env = v
    }
    viper.SetConfigName(fmt.Sprintf("%s-env", env))
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config/")
    viper.AutomaticEnv()
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("fatal error config file: %w", err))
    }
    var conf config.Config
    err = viper.Unmarshal(&conf)
    if err != nil {
        panic(fmt.Errorf("unable to decode into struct: %w", err))
    }
    overrideUsingEnvVars(&conf)
    return &conf
}

func GetJSONRawBody(c echo.Context) map[string]interface{}  {

     jsonBody := make(map[string]interface{})
     err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
     if err != nil {

         log.Error("empty json body")
         return nil
     }

    return jsonBody
}
