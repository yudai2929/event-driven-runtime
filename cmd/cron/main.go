package main

import "github.com/yudai2929/event-driven-runtime/triggers/cron"

const functionsDir = "functions/bin"

func main() {
	cron.StartTrigger(functionsDir)
}
