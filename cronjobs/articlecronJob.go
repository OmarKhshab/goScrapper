package cronjobs

import (
	"fmt"

	"example.com/rest-api/routes"
	"github.com/robfig/cron"
)

func GetArticles() *cron.Cron {
	// Create a new cron instance
	c := cron.New()

	// Add a cron job that runs every 10 seconds
	c.AddFunc("@every 01h00m00s", func() { routes.GetAndUpdateArticles() })
	c.AddFunc("@weekly", func() { routes.SaveLogs() })

	// Start the cron scheduler
	c.Start()
	fmt.Println("Cron scheduler initialized")
	return c
}
