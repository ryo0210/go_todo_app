package handler

import (
	"net/http"

	"github.com/ryo0210/go_todo_app/entity"
)

type ListTask struct {
	Service ListTasksService
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Service.ListTasks(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

// curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
// curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/bad_req.json.golden
// curl -i -XGET localhost:18000/tasks
