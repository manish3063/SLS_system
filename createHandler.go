package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createHandler(c *gin.Context) {
	reqBody := Link{}
	err := c.Bind(&reqBody)

	if err != nil {
		if err != nil {
			ress := gin.H{
				"error": err.Error(),
			}

			c.JSON(http.StatusBadRequest, ress)
			return
		}
	}

	_, err_result, _ := CreateLink(reqBody)

	if err_result != "" {
		res := gin.H{
			"error": err_result,
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"success":   true,
		"long_link": reqBody.LongLink,
		//"short_link": "localhost:8082 "+short_link+"",
	}
	c.JSON(http.StatusOK, res)

}

func CreateLink(reqbody Link) (bool, string, int) {
	var result = true
	var err_responce = ""
	var id = 0

	sqlStatement := `
INSERT INTO sls_link(long_link,short_link)
VALUES ($1,$2) RETURNING id`
	err2 := DB.QueryRow(sqlStatement, reqbody.LongLink, "test").Scan(&id)

	fmt.Println(err2)
	fmt.Println("kkk", sqlStatement)
	if err2 != nil {
		err_responce = "Something went wrong"
		return false, err_responce, id
	}

	fmt.Println(id)

	result = false

	return result, err_responce, id

}
