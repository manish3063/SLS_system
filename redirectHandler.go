package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func redirectHandler(c *gin.Context) {

	id, _ := c.Params.Get("id")

	// if err != false {
	// 	ress := gin.H{
	// 		"error": "something went wrong",
	// 	}

	// 	c.JSON(http.StatusBadRequest, ress)
	// 	return
	// }
	fmt.Println(id)
	idd, _ := strconv.Atoi(id)
	long_url := getLongURL(idd)

	c.Redirect(http.StatusMovedPermanently, long_url)

}

func getLongURL(id int) string {

	long_url := ""
	SQL := `SELECT "long_link" FROM sls_link WHERE id=$1`

	rows := DB.QueryRow(SQL, id)

	rows.Scan(&long_url)
	fmt.Println(long_url)

	return long_url
}
