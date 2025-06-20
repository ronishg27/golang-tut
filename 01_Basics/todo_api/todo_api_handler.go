package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var AllTodos = []Todo{
	{ID: 1, Title: "Learn golang", Description: "Learn golang from the basics", Completed: true},
	{ID: 2, Title: "Learn Docker", Description: "Learn Docker in 5 days", Completed: false},
	{ID: 3, Title: "Learn Kubernetes", Description: "Learn Kubernetes in 5 days", Completed: false},
	{ID: 4, Title: "Learn Terraform", Description: "Learn Terraform in 5 days", Completed: false},
	{ID: 5, Title: "Learn Ansible", Description: "Learn Ansible in 5 days", Completed: false},
	{ID: 6, Title: "Learn Packer", Description: "Learn Packer in 5 days", Completed: false},
}

func CreateTodoHandler(ctx *gin.Context) {
	var myTodo Todo
	err := ctx.ShouldBindJSON(&myTodo)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Failed to convert from JSON to Todo struct",
			"error":   err,
		})
		return
	}
	AllTodos = append(AllTodos, myTodo)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Todo Created successfully",
		"todo":    AllTodos,
	})
}

func GetAllTodosHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo Fetched successfully",
		"todos":   AllTodos,
		"size":    len(AllTodos),
		"status":  200,
	})
}

func GetTodoByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoUInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Todo ID",
		})
		return
	}
	for i := range AllTodos {

		if AllTodos[i].ID == todoUInt {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Todo Fetched successfully",
				"todo":    AllTodos[i],
				"status":  200,
			})
			return
		}

	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Todo Not Found",
	})
}

func UpdateTodoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoUInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Todo ID",
		})
		return
	}
	var todoToUpdate Todo
	err = ctx.ShouldBindJSON(&todoToUpdate)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid todo data",
			"error":   err.Error(),
		})
		return
	}

	for i := range AllTodos {
		if AllTodos[i].ID == todoUInt {
			AllTodos[i].Title = todoToUpdate.Title
			AllTodos[i].Description = todoToUpdate.Description
			AllTodos[i].Completed = todoToUpdate.Completed

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Todo Updated successfully",
				"todo":    AllTodos[i],
				"status":  200,
			})
			return
		}
	}
}

func DeleteTodoHandler(ctx *gin.Context) {
	todoIdParam := ctx.Param("id")
	todoUInt, err := strconv.Atoi(todoIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Todo ID",
		})
		return
	}
	for i, t := range AllTodos {
		if t.ID == todoUInt {
			AllTodos = append(AllTodos[:i], AllTodos[i+1:]...)
			// AllTodos = slices.Delete(AllTodos, i, i+1)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Todo Deleted successfully",
				"todo":    AllTodos,
				"status":  200,
			})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Todo Not Found",
	})

}
