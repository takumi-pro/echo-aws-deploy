package main

import (
	todo "echo-apprunner"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiHandler struct{}

type ResponseMessage struct {
	Message string `json:"mesage"`
}

func (h ApiHandler) GetTasks(ctx echo.Context, params todo.GetTasksParams) error {
	t := "title1"
	d := "des1"
	s := "test"
	return ctx.JSON(http.StatusOK, &todo.Task{
		Id:          1,
		Title:       &t,
		Description: &d,
		CreatedAt:   &s,
	})
}

// Create New Task
// (POST /tasks/create)
func (h ApiHandler) PostTasksCreate(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, ResponseMessage{
		Message: "PostTasksCreate",
	})
}

// Fetch Task by Id
// (GET /tasks/{taskId})
func (h ApiHandler) GetTasksTaskId(ctx echo.Context, taskId string) error {
	return ctx.JSON(http.StatusOK, ResponseMessage{
		Message: "GetTasksTaskId",
	})
}

// Delete Task
// (DELETE /tasks/{taskId}/delete)
func (h ApiHandler) DeleteTasksTaskIdDelete(ctx echo.Context, taskId string) error {
	return ctx.JSON(http.StatusOK, ResponseMessage{
		Message: "DeleteTasksTaskIdDelete",
	})
}

// Edit Task
// (PUT /tasks/{taskId}/edit)
func (h ApiHandler) PutTasksTaskIdEdit(ctx echo.Context, taskId string) error {
	return ctx.JSON(http.StatusOK, ResponseMessage{
		Message: "PutTasksTaskIdEdit",
	})
}

func main() {
	e := echo.New()

	api := ApiHandler{}
	todo.RegisterHandlers(e, api)
	e.Logger.Fatal(e.Start(":8080"))
}
