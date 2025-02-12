package save

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"url_rest_api/internal/http-server/api/response"
	"url_rest_api/internal/lib/random"
	"url_rest_api/internal/logger/sl"
	"url_rest_api/internal/storage"
)

// TODO: move to config/bd
const (
	aliasLength = 5
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) error
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
			alias = random.RandomURL(aliasLength)
		}

		err = urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExist) {
			log.Info("url already exists", slog.String("url", req.URL))

			render.JSON(w, r, response.Error("url already exists"))

			return
		}
		if err != nil {
			log.Error("failed to save url", sl.Err(err))

			render.JSON(w, r, response.Error("failed to save url"))

			return
		}

		//log.Info("saved url", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    alias,
		})
	}
}
