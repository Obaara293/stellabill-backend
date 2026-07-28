package outbox

import "github.com/prometheus/client_golang/prometheus"

var (
	OutboxPublisherLag            *prometheus.GaugeVec
	OutboxPublisherInflight       prometheus.Gauge
	OutboxPublisherLimit          prometheus.Gauge
	ChaosOutboxCancellationsTotal prometheus.Counter

	// RetryBudgetAvailable reports the fraction (0.0–1.0) of the retry budget
	// still available. A value trending toward 0.0 signals that retries are
	// consuming too large a share of successful requests and the system is at
	// risk of a thundering-herd storm.
	RetryBudgetAvailable prometheus.Gauge

	// RetryDeniedTotal counts how many retries were denied because the retry
	// budget was exhausted. Spikes in this metric indicate that downstreams
	// are unhealthy and the budget is protecting them from further load.
	RetryDeniedTotal prometheus.Counter
)

func init() {
	OutboxPublisherLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "outbox_publisher_lag_seconds",
			Help: "Lag in seconds between event occurrence and publisher cursor position per publisher",
		},
		[]string{"publisher"},
	)
	_ = prometheus.Register(OutboxPublisherLag)

	OutboxPublisherInflight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "outbox_publisher_inflight",
			Help: "Number of in-flight publish requests to the downstream HTTP service",
		},
	)
	_ = prometheus.Register(OutboxPublisherInflight)

	OutboxPublisherLimit = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "outbox_publisher_limit",
			Help: "Current adaptive concurrency limit for the HTTP publisher",
		},
	)
	_ = prometheus.Register(OutboxPublisherLimit)

	ChaosOutboxCancellationsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "chaos_outbox_cancellations_total",
		Help: "Total number of outbox publish cancellations injected by the chaos hook (staging only)",
	})
	_ = prometheus.Register(ChaosOutboxCancellationsTotal)

	RetryBudgetAvailable = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "retry_budget_available",
		Help: "Fraction (0.0–1.0) of the retry budget still available; 0.0 means budget exhausted",
	})
	_ = prometheus.Register(RetryBudgetAvailable)

	RetryDeniedTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "retry_denied_total",
		Help: "Total number of retries denied due to budget exhaustion",
	})
	_ = prometheus.Register(RetryDeniedTotal)
}
