# GFramework
GFramework is a Web Framework in Go and friendly for people coming from PHP Frameworks like Laravel, SlimPHP, etc. Still is under development, please be patient.

## Quick Start
#### Download and install
    go get github.com/juniorocheiruiz/gframework

#### Create a file for example call it 'main.go'
```go
package main

import (
	"fmt"
	"github.com/juniorocheiruiz/gframework/http"
	"github.com/juniorocheiruiz/gframework/routing"
)

func main() {
	options := routing.RouterOptions{Port: 8080}
	router := routing.NewRouter(&options)

	router.Get("/", func(request *http.Request, response *http.Response) {
		response.WriteString("home")
	}).SetName("Home")

	router.Get("/name/{name:string}/age/{age?:int}", func(request *http.Request, response *http.Response) {
		name := request.Params.Get("name")
		age, _ := request.Params.GetInt("age")

		response.WriteString(fmt.Sprintf("%s - %d", name, age))
	}).SetName("user")

	router.Run()
}
```
##### Other examples:

```go
package main

import (
	"fmt"

	"github.com/juniorocheiruiz/gframework/http"
	"github.com/juniorocheiruiz/gframework/routing"
)

func main() {
	options := routing.RouterOptions{Port: 8080}
	router := routing.NewRouter(&options)

	router.Get("/", func(request *http.Request, response *http.Response) {
		response.WriteString("This is the home of the app.")
	}).SetName("Home")

	router.Get("/name/{name:string}/age/{age?:int}", func(request *http.Request, response *http.Response) {
		name := request.Params.Get("name")
		age, _ := request.Params.GetInt("age")

		response.WriteString(fmt.Sprintf("%s - %d", name, age))
	}).SetName("user")
	
	// Parsing JSON
	router.Post("/json", func(request *http.Request, response *http.Response) {
		/*
			json:
			{
				"name": "Junior",
				"age": 21,
				"cars": ["Renault", "Ford", "Ferrari"],
				"id": {
					"type": "DNI",
					"value": "12345678Z"
				}
			}
		*/
		jsonObject := struct {
			Name string
			Age  int
			Cars []string
			Id   struct {
				Type  string
				Value string
			}
		}{}

		err := request.ParseJSON(&jsonObject)

		if err == nil {
			response.WriteString(jsonObject.Name + "\n")
			response.WriteString(fmt.Sprintf("%d\n", jsonObject.Age))

			if jsonObject.Cars != nil {
				response.WriteString(jsonObject.Cars[0] + " - " + jsonObject.Cars[1] + " - " + jsonObject.Cars[2] + "\n")
			}

			response.WriteString(jsonObject.Id.Type + " - " + jsonObject.Id.Value)
		} else {
			response.SetStatusCode(400)
		}

	}).SetName("json")
	
	// Saving a file
	router.Post("/file", func(request *http.Request, response *http.Response) {
		files := request.UploadedFiles["image_logo"]

		err := files[0].Save("/some_path/image_logo.png")

		if err != nil {
			response.WriteString("Error uploading the file. Try again more later.")
		}
	}).SetName("file")

	router.Run()
}
```


#### Build and run
    go build main.go
    ./main

#### Go to [http://localhost:8080](http://localhost:8080)


## License

gframework source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
