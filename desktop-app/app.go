package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetCpuDetails() string {
	cpuDetails, err := cpu.Info()
	if err != nil {
		panic(err)
	}

	// Convert the CPU details to a string representation
	cpuDetailsStr := ""
	for _, cpuInfo := range cpuDetails {
		cpuDetailsStr += "Model: " + cpuInfo.ModelName + "\n"
		cpuDetailsStr += "Cores: " + strconv.Itoa(int(cpuInfo.Cores)) + "\n"
		cpuDetailsStr += "Speed: " + fmt.Sprintf("%.2f", cpuInfo.Mhz) + " MHz\n"
		cpuDetailsStr += "Cache Size: " + strconv.Itoa(int(cpuInfo.CacheSize)) + " KB\n"
	}

	return cpuDetailsStr
}
