package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/api/v1/madup/advertisers", advertiserList)
	router.POST("/api/v1/madup/advertisers", advertiserCreate)
	router.GET("/api/v1/madup/advertisers/:id", advertiserRetrieve)
	router.PUT("/api/v1/madup/advertisers/:id", advertiserUpdate)
	router.DELETE("/api/v1/madup/advertisers/:id", advertiserDelete)
	router.Run() // Listen and serve on 0.0.0.0:8000 (for windows "localhost:8080")
}

func advertiserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "광고주 리스트",
	})
}
func advertiserCreate(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "광고주 생성",
	})
}
func advertiserRetrieve(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "광고주 디테일",
	})
}
func advertiserUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "광고주 수정",
	})
}
func advertiserDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "광고주 삭제",
	})
}
