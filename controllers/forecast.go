package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Forecast() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("forecast")

		connStr := "postgres://postgres:password@localhost/?sslmode=disable"
		// connStr := "postgres://postgres:password@localhost/z?sslmode=disable"
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		} else {
			fmt.Println("connected")
		}

		rows, err := db.Query(`select * from Forecast srt`)

		for rows.Next() {
			fmt.Println(rows)
		}
	}
}
