package main

import (
	"context"

	ui "github.com/gizak/termui/v3"
	"github.com/luqman-v1/orbit-pro-monitor/repo"
	"github.com/luqman-v1/orbit-pro-monitor/usecase"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	ctx := context.Background()
	repo := repo.NewOrbit(repo.OrbitConfig{
		URL:      "http://192.168.8.1/",
		Username: "",
		Password: "",
	})
	uc := usecase.NewMonitor(usecase.Monitor{
		Repo: repo,
	})
	err := uc.Monitor(ctx)
	if err != nil {
		return
	}
}
