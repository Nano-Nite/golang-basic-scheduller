/**
 * @author Dimas Ilyasa
 * @email dimas.ilyasa@soluix.ai
 */
package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/robfig/cron"
)

func main() {

	schedulerTime := "* * * * *"
	var wg sync.WaitGroup

	myFigure := figure.NewColorFigure("Soluix", "", "red", true)
	myFigure.Print()
	fmt.Printf("Date init : %v\n\n\n\n", time.Now().Format("02-01-2006 15:04:05"))

	wg.Add(1)
	defer wg.Wait()
	go func() {
		c := cron.New()
		if err := c.AddFunc(schedulerTime, func() {
			log.Printf("Running shceduller every second, now at : %v", time.Now().Format(time.RFC3339))
		}); err != nil {
			log.Println("Failed to start Automatch scheduler")
		}
		c.Start()
	}()

}
