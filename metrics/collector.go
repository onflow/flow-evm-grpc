package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Collector interface {
	ApiErrorOccurred()
	TraceDownloadFailed()
	ServerPanicked(err error)
	EvmBlockHeightUpdated(height uint64)
	EvmAccountCalled(address string)
	MeasureRequestDuration(start time.Time, method string)
}

type DefaultCollector struct {
	// TODO: for now we cannot differentiate which api request failed number of times
	apiErrorsCounter          prometheus.Counter
	traceDownloadErrorCounter prometheus.Counter
	serverPanicsCounters      *prometheus.CounterVec
	evmBlockHeight            prometheus.Gauge
	evmAccountCallCounters    *prometheus.CounterVec
	requestDurations          *prometheus.HistogramVec
}

func NewCollector(factory promauto.Factory) Collector {
	apiErrors := factory.NewCounter(prometheus.CounterOpts{
		Name: "api_errors_total",
		Help: "Total number of errors returned by the endpoint resolvers",
	})

	traceDownloadErrorCounter := factory.NewCounter(prometheus.CounterOpts{
		Name: "trace_download_errors_total",
		Help: "Total number of trace download errors",
	})

	serverPanicsCounters := factory.NewCounterVec(prometheus.CounterOpts{
		Name: "api_server_panics_total",
		Help: "Total number of panics handled by server",
	}, []string{"error"})

	evmBlockHeight := factory.NewGauge(prometheus.GaugeOpts{
		Name: "evm_block_height",
		Help: "Current EVM block height",
	})

	evmAccountCallCounters := factory.NewCounterVec(prometheus.CounterOpts{
		Name: "evm_account_calls_total",
		Help: "Total number of calls to specific evm account",
	}, []string{"address"})

	// TODO: Think of adding 'status_code'
	requestDurations := factory.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api_request_duration_seconds",
		Help:    "Duration of requests made to the endpoint resolvers",
		Buckets: prometheus.DefBuckets,
	}, []string{"method"})

	return &DefaultCollector{
		apiErrorsCounter:          apiErrors,
		traceDownloadErrorCounter: traceDownloadErrorCounter,
		serverPanicsCounters:      serverPanicsCounters,
		evmBlockHeight:            evmBlockHeight,
		evmAccountCallCounters:    evmAccountCallCounters,
		requestDurations:          requestDurations,
	}
}

func (c *DefaultCollector) ApiErrorOccurred() {
	c.apiErrorsCounter.Inc()
}

func (c *DefaultCollector) TraceDownloadFailed() {
	c.traceDownloadErrorCounter.Inc()
}

func (c *DefaultCollector) ServerPanicked(err error) {
	c.serverPanicsCounters.With(prometheus.Labels{"error": err.Error()}).Inc()
}

func (c *DefaultCollector) EvmBlockHeightUpdated(height uint64) {
	c.evmBlockHeight.Set(float64(height))
}

func (c *DefaultCollector) EvmAccountCalled(address string) {
	c.evmAccountCallCounters.With(prometheus.Labels{"address": address}).Inc()

}

func (c *DefaultCollector) MeasureRequestDuration(start time.Time, method string) {
	duration := time.Since(start)
	c.requestDurations.With(prometheus.Labels{"method": method}).Observe(float64(duration))
}
