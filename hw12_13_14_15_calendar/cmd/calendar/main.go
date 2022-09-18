package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/app"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/logger"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/models"
	internalhttp "github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/server/http"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "etc/calendar/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig()
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		fmt.Println("failed to parse config file: " + err.Error())
		os.Exit(1)
	}

	var eventStorage storage.EventsStorage
	if config.DB.InMemoryStorage {
		eventStorage = memorystorage.New(map[string]models.Event{})
	} else {
		db, err := sqlx.Open("pgx", config.DB.URI)
		if err != nil {
			fmt.Println("failed to open database: " + err.Error())
			os.Exit(1)
		}
		eventStorage = sqlstorage.New(*db)
	}
	logg := logger.New(config.Logger.Level)
	calendar := app.New(logg, eventStorage)

	server := internalhttp.NewServer(logg, calendar)
	server.Host = config.Server.Host
	server.Port = config.Server.Port

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
