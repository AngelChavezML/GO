package handlers

import (
	"encoding/json"
	"go_web/post/internal"
	"go_web/post/platform/web"
	"io"
	"net/http"
)

type TaskRequestBody struct {
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

type TaskJSON struct {
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

type DefaultTask struct {
	TaskService internal.Greeting
}

func (d *DefaultTask) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body TaskRequestBody
		bytes, _ := io.ReadAll(r.Body)

		if err := json.Unmarshal(bytes, &body); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]interface{}{
				"error": "invalid request body",
			})
			return
		}
		data := TaskJSON(body)
		web.ResponseJSON(w, http.StatusCreated, map[string]any{
			"message": "Hello " + data.First_Name + " " + data.Last_Name,
		})
		//
	}

}
