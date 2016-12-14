package stats

import (
	"net/http"
	"time"
)

type Info struct {
	Percentage []float64
	TimeStamp  time.Time
}

//go:generate counterfeiter -o ../fakes/fake_cpustats.go . CPUStats

type CPUStats interface {
	Run() error
	Collect() ([]Info, error)
	Stop() error
}

type Collector struct {
}

func NewCollector() *Collector {
	return &Collector{}
}

type Handler struct {
	cpuStats CPUStats
}

func NewStatHandler(c CPUStats) *Handler {
	return &Handler{
		cpuStats: c,
	}
}

func (c *Handler) Start(w http.ResponseWriter, r *http.Request) {
	// TODO: will return an error so we should handle that case
	err := c.cpuStats.Run()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
