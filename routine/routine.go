// Package routine controls the entire Process
package routine

import (
	"enrolment/client"
	"enrolment/conf"
	"enrolment/logger"
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

	// Setting up multiple threads for one lesson enrolment
	// TODO: Add Mutex to control the same course
	for t := 0; t < r.Pg.Threads; t++ {
		for idx, a := range r.Ac {
			for k, v := range a.Courses {
				logger.Infof("Setting up Client for %v: %v", k, v)
				c, err := client.NewEClient(&r.Ac[idx], r.Ocr, k)
				if err != nil {
					return nil, err
				}
				r.clients = append(r.clients, c)
			}
		}
	}

	return r, nil
}

func eroutine(c *client.EClient) {
	logger.Infof("eroutine launched for: %v", c)
	MaxRetry := 5
	var err error
	for i := 0; i < MaxRetry; i++ {
		if err = c.Refresh(err); err == nil {
			break
		}
		logger.Warnf("Login Error: %v, Retry Left: %v/%v", err, MaxRetry-i-1, MaxRetry)
		time.Sleep(1000 * time.Millisecond)
	}

	if err != nil {
		logger.Errorf("Retry %v times Failed: %v", MaxRetry, err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err = c.Select(); err != nil {
				logger.Tracef("encounter: %v", err)
				switch err {
				case client.ErrorSuccess:
					goto f1
				case client.ErrorMax2:
					goto f1
				case client.ErrorMultiple:
					goto f1
				case client.ErrorSelected:
					goto f1
				case client.ErrorFast:
					time.Sleep(1000 * time.Millisecond)
				default:
					time.Sleep(10 * time.Millisecond)
					if err = c.Refresh(nil); err != nil {
						// Only 1 retry
						c.Refresh(nil)
					}
				}
			} else {
				time.Sleep(10 * time.Microsecond)
			}
		}
	f1:
		logger.Infof("[%s] 结束选课: %v", c.Notation(), err)
	}()
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
