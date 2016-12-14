package stats_test

import (
	"cpumonitor/fakes"
	"cpumonitor/stats"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stats", func() {
	Context("Start", func() {
		It("returns 200 StatusOK", func() {
			cpustat := new(fakes.FakeCPUStats)
			statsHandler := stats.NewStatHandler(cpustat)
			testServer := httptest.NewServer(http.HandlerFunc(statsHandler.Start))
			resp, err := http.Get(testServer.URL)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns non 200 if error occurs", func() {
			cpustat := new(fakes.FakeCPUStats)
			cpustat.RunReturns(errors.New("error"))
			statsHandler := stats.NewStatHandler(cpustat)
			testServer := httptest.NewServer(http.HandlerFunc(statsHandler.Start))
			resp, err := http.Get(testServer.URL)

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
		})
	})
	Context("Collect", func() {
		It("returns a slice of cpu percentages and timestamps", func() {
			collector := stats.NewCollector()
			err := collector.Collect()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
