package middleware

import (
	"myapp/data"

	"github.com/snirkop89/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
