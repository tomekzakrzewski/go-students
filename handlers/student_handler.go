package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tomekzakrzewski/go-students/db"
	"github.com/tomekzakrzewski/go-students/error"
	"github.com/tomekzakrzewski/go-students/models"
)

type StudentHandler struct {
	studentStore *db.StudentStore
}

func NewStudentHandler(studentStore *db.StudentStore) *StudentHandler {
	return &StudentHandler{
		studentStore: studentStore,
	}
}

// HandleAddStudent adds a new student.
//
//	@Summary		Add a new student
//	@Description	Add a new student to the database.
//	@ID				add-student
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			studentParams	body		models.CreateStudentParams	true	"Student parameters"
//	@Success		201				{object}	models.Student				"Successful operation"
//	@Failure		400				{object}	error.Http					"Invalid request body"
//	@Failure		422				{object}	error.Http					"Invalid validation"
//	@Router			/students [post]
func (h *StudentHandler) HandleAddStudent(c *gin.Context) {
	var studentParams models.CreateStudentParams
	if err := c.ShouldBindJSON(&studentParams); err != nil {
		c.Error(error.InvalidRequestBody())
		return
	}
	errors := studentParams.Validate()
	if len(errors) > 0 {
		c.Error(error.InvalidValidation(errors))
		return
	}
	student := models.NewStudentFromParams(&studentParams)
	insertedUser, err := h.studentStore.Insert(student)
	if err != nil {
		c.Error(err)
		return
	}
	log.Println("inserted student into database")
	c.JSON(http.StatusCreated, insertedUser)
}

// HandleGetStudents retrieves all students.
//
//	@Summary		Retrieve all students
//	@Description	Retrieve details of all students.
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Student	"Successful operation"
//	@Failure		500	{object}	error.Http		"Internal server error"
//	@Router			/students [get]
func (h *StudentHandler) HandleGetStudents(c *gin.Context) {
	students, err := h.studentStore.FindAll()
	if err != nil {
		c.Error(err)
		return
	}
	log.Println("found all students from database")
	c.JSON(http.StatusOK, students)
}

// HandleGetStudent retrieves a student by ID.
//
//	@Summary		Retrieve a student by ID
//	@Description	Retrieve a student's details by their ID.
//	@ID				get-student-by-id
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"Student ID"
//	@Success		200	{object}	models.Student	"Successful operation"
//	@Failure		400	{object}	error.Http		"Invalid ID format"
//	@Failure		404	{object}	error.Http		"Student not found"
//	@Router			/students/{id} [get]
func (h *StudentHandler) HandleGetStudent(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.Error(error.InvalidId())
		return
	}
	student, err := h.studentStore.FindById(id)
	if err != nil {
		c.Error(error.NotFound())
		return
	}
	log.Println("found student with id: ", idString)
	c.JSON(http.StatusOK, student)
}

// HandleUpdateStudent updates a student by ID.
//
//	@Summary		Update a student by ID
//	@Description	Update a student's details by their ID.
//	@ID				update-student-by-id
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int							true	"Student ID"
//	@Param			studentParams	body		models.CreateStudentParams	true	"Student parameters"
//	@Success		200				{string}	string						"updated user with id {id}"
//	@Failure		400				{object}	error.Http					"Invalid ID or request body"
//	@Failure		404				{object}	error.Http					"Student not found"
//	@Router			/students/{id} [put]
func (h *StudentHandler) HandleUpdateStudent(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.Error(error.InvalidId())
		return
	}
	var studentParams models.CreateStudentParams
	if err := c.ShouldBindJSON(&studentParams); err != nil {
		c.Error(error.InvalidRequestBody())
		return
	}
	errors := studentParams.ValidateUpdate()
	if len(errors) > 0 {
		c.Error(error.InvalidValidation(errors))
		return
	}
	student := models.NewStudentFromParams(&studentParams)
	err = h.studentStore.Update(id, *student)
	if err != nil {
		c.Error(error.NotFound())
		return
	}
	log.Println("updated user with id ", id)
	c.JSON(http.StatusOK, "updated user with id "+idString)
}

// HandleDeleteStudent deletes a student by ID.
//
//	@Summary		Delete a student by ID
//	@Description	Delete a student identified by their ID.
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true	"Student ID"
//	@Success		200	{object}	map[string]interface{}	"Successful operation"
//	@Failure		400	{object}	error.Http				"Invalid ID format"
//	@Failure		404	{object}	error.Http				"Student not found"
//	@Router			/students/{id} [delete]
func (h *StudentHandler) HandleDeleteStudent(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.Error(error.InvalidId())
		return
	}
	err = h.studentStore.Delete(id)
	if err != nil {
		c.Error(error.NotFound())
		return
	}
	log.Println("deleted user with id ", id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted user with id " + idString})
}
