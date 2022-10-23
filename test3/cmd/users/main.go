package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dingowd/CyberZilla/test3/internal/app"
	"github.com/dingowd/CyberZilla/test3/internal/config"
	"github.com/dingowd/CyberZilla/test3/internal/httpserver"
	"github.com/dingowd/CyberZilla/test3/internal/logger"
	"github.com/dingowd/CyberZilla/test3/internal/logger/lrus"
	"github.com/dingowd/CyberZilla/test3/internal/logger/standart"
	"github.com/dingowd/CyberZilla/test3/internal/storage"
	"github.com/dingowd/CyberZilla/test3/internal/storage/mysql"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")

	// init config
	conf := config.NewConfig()
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Fprintln(os.Stdout, "ошибка чтения toml файла "+err.Error()+", установка параметров по умолчанию")
		conf = config.Default()
	}

	// init output
	var output io.Writer
	file, err := os.OpenFile(conf.Logger.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err == nil {
		output = file
		defer file.Close()
	} else {
		output = os.Stdout
	}

	// init logger
	var logg logger.Logger
	switch conf.LogName {
	case "lrus":
		logg = lrus.New(conf.Logger.Level, output)
	default:
		logg = standart.New(conf.Logger.Level, output)
	}

	// init storage
	var store storage.Storage
	store = mysql.New(logg)
	if err := store.Connect(conf.DSN); err != nil {
		logg.Error("failed to connect database" + err.Error())
		os.Exit(1) // nolint:gocritic
	}
	defer store.Close()

	// init application
	users := app.New(logg, store)

	// init http server
	server := httpserver.NewServer(users, conf.HTTPSrv)

	// graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	go func() {
		<-ctx.Done()
		logg.Info("Users service stopping...")
		server.Stop()
		logg.Info("Users service stopped")
		time.Sleep(5 * time.Second)
	}()

	logg.Info("Users service is running...")

	// start http server
	server.Start()
}
