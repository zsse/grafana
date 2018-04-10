package sqlstore

import (
	"github.com/rkrikbaev/grafana/pkg/bus"
	m "github.com/rkrikbaev/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
