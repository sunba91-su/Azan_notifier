package main

import (
	"azan_notifier/controller"
	"azan_notifier/handlers"
)

func main() {
	handlers.LoadEnvs()
	controller.StartProgram()
}
