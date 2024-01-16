// Package routine controls the entire Process
package routine

import (
	"enrollment/client"
	"enrollment/conf"
	"enrollment/logger"
	"fmt"
	"sync"
	"time"
)

type Routine struct {
	*conf.Conf
	clients []*client.EClient
}

func NewRoutine() (*Routine, error) {
	r := &Routine{
		Conf:    conf.NewConfig(),
		clients: []*client.EClient{},
	}

	if err := r.LoadConfig(); err != nil {
		return nil, err
	}

	for _, acc := range r.Ac {
		for idx := range acc.CourseNo {
			// TODO: Add Multiple Threads for each course
			c, err := client.NewEClient(acc.No, acc.Pw, acc.Comment, acc.CourseNo[idx], acc.CourseComment[idx])
			if err != nil {
				return nil, err
			}
			r.clients = append(r.clients, c)
		}
	}

	return r, nil
}

func eroutine(c *client.EClient) {
	logger.Infof("eroutine launched for: %v", c)
	MaxRetry := 5
	err := fmt.Errorf("Init")
	for i := 0; i < MaxRetry; i++ {
		if err := c.Refresh(); err == nil {
			break
		}
		logger.Warnf("Login Error: %v, Retry Left: %v/%v", err, MaxRetry-i-1, MaxRetry)
	}

	if err != nil {
		logger.Errorf("Retry %v times Failed: %v", MaxRetry, err)
		return
	}

	for {
		if err := c.Select(); err != nil {
			logger.Warnf("encounter: %v", err)
			switch err {
			case client.ErrorSuccess:
				break
			case client.ErrorMax2:
				break
			case client.ErrorMultiple:
				break
			case client.ErrorSelected:
				break
			default:
				time.Sleep(100 * time.Millisecond)
				c.Refresh()
			}
		} else {
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func (r *Routine) Run() {
	wg := &sync.WaitGroup{}
	for _, c := range r.clients {
		wg.Add(1)
		go func(c *client.EClient) {
			defer wg.Done()
			// TODO: Remove Scan
			eroutine(c)
		}(c)
	}

	wg.Wait()
}
