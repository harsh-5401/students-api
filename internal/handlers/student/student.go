package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"students-api/internal/types"
	"students-api/internal/utils/responses"

	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}

		if err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError((err)))
		}

		// validation of request
		err = validator.New().Struct(student)

		if err != nil {

			validateErr := err.(validator.ValidationErrors)
			responses.WriteJson(w, http.StatusBadRequest, responses.ValidationError(validateErr))
		}

		slog.Info("creating a student")

		responses.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
