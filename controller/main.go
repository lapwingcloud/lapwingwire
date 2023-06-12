package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lapwingcloud/lapwingwire/controller/ent"
	"github.com/lapwingcloud/lapwingwire/controller/rest"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	// logWriter := zerolog.NewConsoleWriter()
	logWriter := os.Stdout
	logger := zerolog.New(logWriter).With().
		Timestamp().
		Str("service", "lapwingwire-controller").
		Str("server_hostname", hostname).
		Str("type", "app").
		Logger()

	sqlDriver, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/lapwingwire?parseTime=true")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to call sql.Open")
	}
	if err = sqlDriver.DB().Ping(); err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
	}
	db := ent.NewClient(ent.Driver(sqlDriver))
	defer db.Close()
	if err := db.Schema.Create(context.Background()); err != nil {
		logger.Fatal().Err(err).Msg("failed creating schema resources")
	}

	router := chi.NewRouter()
	router.Use(hlog.NewHandler(logger))
	router.Use(hlog.RequestIDHandler("request_id", "X-Request-Id"))
	router.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			clientIP = r.RemoteAddr
		}
		hlog.FromRequest(r).Info().
			Str("type", "access").
			Str("request_method", r.Method).
			Str("request_host", r.Host).
			Stringer("request_url", r.URL).
			Str("request_user_agent", r.Header.Get("User-Agent")).
			Dur("request_duration", duration).
			Int("response_status", status).
			Int("response_bytes", size).
			Str("client_ip", clientIP).
			Msg("")
	}))

	rest.NewHandler(db).RegisterRoutes(router)
	if err = http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal().Err(err).Msg("http server exited unexpectedly")
	}
}
