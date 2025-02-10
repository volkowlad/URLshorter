package save

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"url_rest_api/internal/http-server/api/response"
	"url_rest_api/internal/logger/sl"
)

	// TODO: move to config/bd
const (
	aliasLength = 5
)

type Request struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req) // распарсить запрос
		if err != nil {
			log.Error("failed to parse request", sl.Err(err))

			render.JSON(w, r, response.Error("failed to parse request"))

			return
		}

		log.Info("request Body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, response.Error("invalid request"))
			// TODO: 1:21:30
			return
		}

		alias := req.Alias
		if alias == "" {
			alias = 
		}
	}
}
