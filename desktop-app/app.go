package main

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
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

func (a *App) GetCpuUsage() string {
	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		panic(err)
	}

	cpuUsageStr := "CPU Usage: "
	for _, cpuPercent := range cpuPercentages {
		cpuUsageStr += fmt.Sprintf("%.2f%% ", cpuPercent)
	}

	return cpuUsageStr
}

func (a *App) GetRamDetails() string {
	ramDetails, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	ramDetailsStr := ""
	ramDetailsStr += "Total: " + strconv.Itoa(int(ramDetails.Total/1024/1024)) + " MB\n"
	ramDetailsStr += "Available: " + strconv.Itoa(int(ramDetails.Available/1024/1024)) + " MB\n"
	ramDetailsStr += "Used: " + strconv.Itoa(int(ramDetails.Used/1024/1024)) + " MB\n"
	ramDetailsStr += "Free: " + strconv.Itoa(int(ramDetails.Free/1024/1024)) + " MB\n"

	return ramDetailsStr
}

func (a *App) GetGpuDetails() string {
	Info := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, err := Info.Output()
	if err != nil {
		panic(err)
	}

	return string(History)
}
