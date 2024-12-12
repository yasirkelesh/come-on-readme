package main

import (
	"log"

	"github.com/come-on-readme/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "come_on_readme",
		//BodyLimit: //bir istek gövdesinin maksimum boyutunu belirler. Varsayılan değer 4 MB'dir.
		//Concurrency: 256*1024, //Eşzamanlı isteklerin maksimum sayısını belirler. Varsayılan değer 256'dır.
		//RequestMethods: []string{"GET", "POST", "PUT", "DELETE"}, //İzin verilen HTTP istek yöntemlerini belirler. Varsayılan değerler: "GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS", "TRACE".
		ServerHeader: "Fiber",
		//Prefork:      true, //SO_REUSEPORTsocket seçeneğinin kullanımını etkinleştirir. Bu, aynı bağlantı noktasını dinleyen birden fazla Go işlemi başlatır. soket parçalama hakkında daha fazla bilgi edinin.
	})

	app.Server()

	/* if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	} */

	app.Get("/", handlers.Hello)
	//api := app.Group("/api", handlers.Hello)

	//v1 := api.Group("/v1", handlers.Hello) // /api/v1
	//v1.Get("/list", handlers.Hello)        // /api/v1/list
	//v1.Get("/user", handlers.Hello)        // /api/v1/user

	//app.Route("/test", func(api fiber.Router) {
	//	api.Get("/foo", handlers.Hello).Name("foo") // /test/foo (name: test.foo)
	//	api.Get("/bar", handlers.Hello).Name("bar") // /test/bar (name: test.bar)
	//}, "test.")

	log.Fatal(app.Listen(":3000"))

	app.Shutdown()
}
