package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	handlers "backend/handlers"
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
	/* if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	} */

	app.Get("/", handlers.Hello)

	log.Fatal(app.Listen(":3000"))
}
