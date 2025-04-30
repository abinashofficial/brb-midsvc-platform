
package main

import (
	"brb-midsvc-platform/app"
	_ "brb-midsvc-platform/docs" // Swagger generated files
	"flag"
	"brb-midsvc-platform/tests"

)


// @title           BRB Microservice API
// @version         1.0
// @description     Mid-Level Backend Assessment Microservice for BRB.

// @host      localhost:8080
// @BasePath  /api
func main() {

	var runUnitTests bool

	flag.BoolVar(&runUnitTests, "runUnitTests", false, "Setting to true will run unit tests")
	flag.Parse()

	if runUnitTests {
		tests.Start()
	} else {
		app.Start()
	}
}
