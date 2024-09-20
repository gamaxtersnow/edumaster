package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	OrderCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "order_count_total",
			Help: "Total number of orders placed.",
		},
		[]string{"status"},
	)
)

func init() {
	// 注册指标
	prometheus.MustRegister(OrderCount)
}
