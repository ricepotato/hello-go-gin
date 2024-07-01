package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func SomeJson(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "go",
		"tag":  "<br>",
	}

	c.JSON(http.StatusOK, data)
}

func HtmlTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "template1.tmpl", gin.H{
		"name": "ricepotato",
	})
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	// 명시적으로 바인딩을 정의하여 multipart form을 바인딩 할 수 있습니다:
	// c.ShouldBindWith(&form, binding.Form)
	// 혹은 ShouldBind 메소드를 사용하여 간단하게 자동으로 바인딩을 할 수 있습니다:

	// curl -v --form user=user --form password=password http://localhost:8080/login

	var form LoginForm
	// 이 경우에는 자동으로 적절한 바인딩이 선택 됩니다
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

func DataStream(c *gin.Context) {
	response, err := http.Get("https://www.google.com/robots.txt")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contenType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="google-robots.txt"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contenType, reader, extraHeaders)
}

func SecureJson(c *gin.Context) {
	names := []string{"lena", "austin", "foo"}

	// 출력내용  :   while(1);["lena","austin","foo"]
	c.SecureJSON(http.StatusOK, names)
}

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func GetIdFromURI(c *gin.Context) {
	// uri에 있는 id와 name을 가져옵니다.
	// /users/987fbc97-4bed-5078-9f07-9141ba07c9f3/ricepoato 와 같이 요청
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": person.ID, "name": person.Name})
}

func SomeXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "Hello, go gin!"})
}

func SomeYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"message": "Hello, go gin!"})
}

var secrets = gin.H{
	"ricepotato": gin.H{"email": "ricepotato40@gmail.com", "phone": "010-1234-5678"},
}

func AdminInfo(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}
