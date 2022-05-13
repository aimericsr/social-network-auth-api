package middlware

import (
	"log"
	"net/http"
	"time"

	"github.com/aimericsr/social-network-auth-api/util"
)

type LogMiddleware struct {
	logger *log.Logger
}

func NewLogMiddleware(logger *log.Logger) *LogMiddleware {
	return &LogMiddleware{logger: logger}
}

func (m *LogMiddleware) Func() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		logRespWriter := util.NewLogResponseWriter(w)

		//next.ServeHTTP(logRespWriter, r)

		m.logger.Printf(
			"duration=%s status=%d body=%s",
			time.Since(startTime).String(),
			logRespWriter.StatusCode,
			logRespWriter.Buf.String())
	})
}
