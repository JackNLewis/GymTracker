package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JackNLewis/gymtracker/db"
)

func main() {

	db.InitializeDB()

	r := gin.Default()
	r.GET("/getexercises", getExercises)
	r.Run() // listen and serve on 0.0.0.0:8080
}

// getExercises retrieves any
func getExercises(c *gin.Context) {

	muscle, ok := c.GetQuery("muscle")
	var ex []db.Exercise
	if !ok {
		ex = db.GetExercises("")
	} else {
		ex = db.GetExercises(muscle)
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, ex)
}
