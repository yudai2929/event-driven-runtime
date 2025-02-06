package cron

import (
	"log"

	cronpkg "github.com/robfig/cron/v3"
	"github.com/yudai2929/event-driven-runtime/metadata"
	"github.com/yudai2929/event-driven-runtime/runtime"
	"github.com/yudai2929/event-driven-runtime/storage"
)

var event = map[string]any{
	"message": "Hello, World!",
}

// StartTrigger is a function that starts the trigger
func StartTrigger(functionsDir string) {
	mc := metadata.NewClient()
	cronSetting := mc.ListCronSettings()
	storage := storage.NewFunctionStorage(functionsDir)

	c := cronpkg.New()

	// AddFunc adds a func to the Cron to be run on the given schedule.
	for _, cs := range cronSetting {
		// AddFunc adds a func to the Cron to be run on the given schedule.
		_, err := c.AddFunc(cs.Schedule, func() {
			filepath := storage.FilePath(cs.FunctionName)
			// Execute function
			_, err := runtime.Execute(filepath, event)
			if err != nil {
				log.Fatalf("failed to execute function: %v", err)
				return
			}

			log.Printf("Function executed with schedule %s", cs.Schedule)
		})
		if err != nil {
			log.Fatalf("failed to add function to cron: %v", err)
			return
		}
	}

	c.Start()
	log.Printf("Cron started with %d functions", len(cronSetting))

	select {}

}
