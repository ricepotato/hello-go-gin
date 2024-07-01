package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ricepotato/hello-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello, go gin!\n")
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	router.LoadHTMLGlob("templates/*")

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"ricepotato": "1234",
	}))
	authorized.GET("/secrets", controllers.AdminInfo)

	router.GET("/ping", controllers.Ping)
	router.GET("/somejson", controllers.SomeJson)
	router.GET("/html", controllers.HtmlTemplate)
	router.GET("/google-robots.txt", controllers.DataStream)
	router.POST("/login", controllers.Login)
	router.GET("/secureJson", controllers.SecureJson)
	router.GET("/users/:id/:name", controllers.GetIdFromURI)
	router.GET("/someXML", controllers.SomeXML)
	router.GET("/someYAML", controllers.SomeYAML)
	//router.Run(":8080") // 간단히 서버 실행

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe() // 서버 옵션과 함께 실행

	// graceful shutdown
	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }

	// go func() {
	// 	// 서비스 접속
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// // 5초의 타임아웃으로 인해 인터럽트 신호가 서버를 정상종료 할 때까지 기다립니다.
	// quit := make(chan os.Signal, 1)
	// // kill (파라미터 없음) 기본값으로 syscanll.SIGTERM를 보냅니다
	// // kill -2 는 syscall.SIGINT를 보냅니다
	// // kill -9 는 syscall.SIGKILL를 보내지만 캐치할수 없으므로, 추가할 필요가 없습니다.
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// // 5초의 타임아웃으로 ctx.Done()을 캐치합니다.
	// select {
	// case <-ctx.Done():
	// 	log.Println("timeout of 5 seconds.")
	// }
	// log.Println("Server exiting")
}
