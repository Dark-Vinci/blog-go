package middleware

import (
	"io/ioutil"
	"net/http"
	"bytes"

	"github.com/gin-gonic/gin"
)

func MaxSize(n int64) gin.HandlerFunc { 
	return func (c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, n)
		buff, errRead := c.GetRawData()

		if errRead != nil {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H {
				"Status": "error",
				"error_message": "too large: upload an image less than 8mb",
			})

			c.Abort()
			return
		}

		buf := bytes.NewBuffer(buff)
		c.Request.Body = ioutil.NopCloser(buf)
	}
}