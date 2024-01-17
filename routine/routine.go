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

	for idx := range r.Ac {
		// TODO: Add Multiple Threads for each course
		c, err := client.NewEClient(&r.Ac[idx])
		if err != nil {
			return nil, err
		}
		r.clients = append(r.clients, c)
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

	wg := &sync.WaitGroup{}
	for idx := range c.CourseNo {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for {
				if err := c.Select(idx); err != nil {
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
		}(idx)
	}
	wg.Wait()
}

func (r *Routine) Run() {
	wg := &sync.WaitGroup{}
	for idx := range r.clients {
		wg.Add(1)
		go func(c *client.EClient) {
			defer wg.Done()
			eroutine(c)
		}(r.clients[idx])
	}
	wg.Wait()
}
