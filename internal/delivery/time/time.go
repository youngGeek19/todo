package get_time

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/gofiber/fiber/v2/log"
	"log/slog"
	"net/http"
	"time"
)

func Time() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.With(slog.String("request_id", middleware.GetReqID(r.Context())))

		t := time.Now()

		render.JSON(w, r, t)

		log.Info("Time successfully sent!", slog.String("time", t.String()))

		return
	}
}
