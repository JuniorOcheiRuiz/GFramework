# GFramework
GFramework is a Web Framework in Go and friendly for people coming from PHP Frameworks like Laravel, SlimPHP, etc. Still is under development, please be patient.

## Quick Start
#### Download and install
    go get github.com/JuniorOcheiRuiz/GFramework

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

#### Build and run
    go build main.go
    ./main

#### Go to [http://localhost:8080](http://localhost:8080)


## License

beego source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).