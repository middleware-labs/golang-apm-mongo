package mongo

import (
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func NewMonitor() {
	otelmongo.NewMonitor()
}
