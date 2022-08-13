package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//ping test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"encode":  "/encode",
			"decode":  "/edecode",
		})
	})

	// router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	//test
	router.GET("/encode", func(c *gin.Context) {
		q, ok := c.GetQuery("q")
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"你的输入是:": q,
				"你的编码是:": my_encode(q),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"你的输入不合法": "输入合法字符",
			})
		}
	})

	//dis var
	v1 := router.Group("/dis_var")
	{
		v1.GET("encode", my_encode_router)
	}

	//upload file，performability 20220813
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "D://1.png")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":8090")
}

func my_encode(s string) string {
	input := []byte(s)

	encodeString := base64.StdEncoding.EncodeToString(input)
	// fmt.Println(encodeString)

	return encodeString
}

func my_encode_router(c *gin.Context) {
	q, ok := c.GetQuery("q")
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"你的输入是:":    q,
			q + " 编码是:": my_encode(q),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"你的输入不合法": "输入合法字符",
		})
	}
}

func my_decode(s string) []byte {

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(decodeBytes))

	return decodeBytes
}

func my_base64(s string) []byte {
	input := []byte(s)

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	// fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	return uDec
}
