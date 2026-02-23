package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
	"error": "Internal server error",
})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
    id := c.Param("id")

    student, err := h.Service.GetStudentByID(id)
    if err != nil {

        if err == services.ErrStudentNotFound {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Student not found",
            })
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal server error",
        })
        return
    }

    c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
	"error": "Invalid request body",
    })
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
    })
    return
}

	c.JSON(http.StatusCreated, student)
}
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.Service.UpdateStudent(id, student)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	updatedStudent, _ := h.Service.GetStudentByID(id)
	c.JSON(http.StatusOK, updatedStudent)
}
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteStudent(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusNoContent)
}