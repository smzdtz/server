package main

import (
	cron "smzdtz-server/corn"
	"smzdtz-server/routes"
)

func main() {
	// Our server will live in the routes package
	cron.RunCronJobs(true)
	routes.Run()
}
