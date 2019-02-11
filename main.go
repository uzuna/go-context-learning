package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func main() {

	var Processies []Process

	fastProcess := &ProcessWithDuration{Name: "Fast", ProcessTime: time.Millisecond * 30}
	slowProcess := &ProcessWithDuration{Name: "Slow", ProcessTime: time.Millisecond * 90}
	clientTimeout := time.Millisecond * 70

	log.Println(fmt.Sprintf("Client timeout is [%s]", clientTimeout))
	Processies = append(Processies,
		fastProcess,
		slowProcess,
	)

	tasks := []Task{
		Task{
			Timeout: time.Millisecond * 20,
		}, Task{
			Timeout: time.Millisecond * 40,
		}, Task{
			Timeout: time.Millisecond * 60,
		}, Task{
			Timeout: time.Millisecond * 80,
		},
	}

	wg := &sync.WaitGroup{}
	for _, task := range tasks {
		for _, p := range Processies {
			wg.Add(1)
			go func(p Process, task Task) {
				defer wg.Done()
				ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
				defer cancel()
				r, err := p.Call(ctx, task)
				log.Println(r, err)
			}(p, task)
		}
	}
	wg.Wait()

	// withproxy
	log.Println("-- WithProxies --")

	var Proxies []Process
	for _, p := range Processies {
		fastProxy := &ProxyProcessWithTimeout{Name: "FastProxy", Server: p, Timeout: time.Millisecond * 45}
		slowProxy := &ProxyProcessWithTimeout{Name: "SlowProxy", Server: p, Timeout: time.Millisecond * 75}
		Proxies = append(Proxies, fastProxy, slowProxy)
	}

	for _, task := range tasks {
		for _, p := range Proxies {
			wg.Add(1)
			go func(p Process, task Task) {
				defer wg.Done()
				ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
				defer cancel()
				r, err := p.Call(ctx, task)
				log.Println(r, err)
			}(p, task)
		}
	}
	wg.Wait()

	// Heavy Loop Process
	log.Println("-- Heavy Process Simulation --")

	simProcess := &ProcessSimulation{Name: "l:1ms", Tick: time.Millisecond}
	fastProxy := &ProxyProcessWithTimeout{Name: "FastProxy", Server: simProcess, Timeout: time.Millisecond * 45}
	slowProxy := &ProxyProcessWithTimeout{Name: "SlowProxy", Server: simProcess, Timeout: time.Millisecond * 75}
	Proxies = append([]Process{}, fastProxy, slowProxy)

	tasks = []Task{
		Task{
			Timeout: time.Millisecond * 40,
			Value:   20,
		}, Task{
			Timeout: time.Millisecond * 40,
			Value:   40,
		}, Task{
			Timeout: time.Millisecond * 80,
			Value:   60,
		}, Task{
			Timeout: time.Millisecond * 80,
			Value:   80,
		},
	}

	for _, task := range tasks {
		for _, p := range Proxies {
			wg.Add(1)
			go func(p Process, task Task) {
				defer wg.Done()
				ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
				defer cancel()
				r, err := p.Call(ctx, task)
				log.Println(r, err)
			}(p, task)
		}
	}
	wg.Wait()
}

// Process is model of server
// Client can call the Call method with arguments context and task.
type Process interface {
	Call(ctx context.Context, task Task) (Responce, error)
}

// Task is simulate the task
type Task struct {
	Timeout time.Duration // Set the time it takes for your process to timeout
	Value   interface{}
}

type Responce string

// ProcessWithDuration is simple processor
// Having ProcessTIme parameter, and returning a response after that time.
type ProcessWithDuration struct {
	Name        string
	ProcessTime time.Duration
}

// Call execute the task or wait cancel signal from context.
func (p *ProcessWithDuration) Call(ctx context.Context, task Task) (Responce, error) {
	start := time.Now()
	select {
	case <-ctx.Done():
		// This case is cancel
		// Receive cancel signal from received context
		return Responce(fmt.Sprintf("%s: Cancel [%s/%s]", p.Name, time.Since(start), task.Timeout)), ctx.Err()
	case <-time.After(task.Timeout):
		// This case is timeout by request task parameter
		return Responce(fmt.Sprintf("%s: Timeout [%s/%s]", p.Name, time.Since(start), task.Timeout)), errors.Errorf("Timeout by Server")
	case <-time.After(p.ProcessTime):
		// This case is complete process
		// When Value has error interface then return error message
		if err, ok := task.Value.(error); ok {
			return Responce(fmt.Sprintf("%s: Error[%s: %s]", p.Name, time.Since(start), err.Error())), errors.WithStack(err)
		}
		return Responce(fmt.Sprintf("%s: Complete [%s/%s]", p.Name, time.Since(start), task.Timeout)), nil
	}
}

// ProxyProcessWithTimeout pass the request to the next process
// and interrupts the request by its own time out setting
type ProxyProcessWithTimeout struct {
	Name    string
	Server  Process
	Timeout time.Duration
}

func (p *ProxyProcessWithTimeout) Call(ctx context.Context, task Task) (Responce, error) {
	type ProxiedResponce struct {
		Responce Responce
		Error    error
	}
	// Generate goroutine
	ch := make(chan ProxiedResponce)

	// Get a cancel function to cancel by proxy
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		// Call server's method with cancelable context by proxy
		res, err := p.Server.Call(childCtx, task)
		ch <- ProxiedResponce{
			Responce: res,
			Error:    err,
		}
	}()
	go func() {
		// Call cancel after proxy timeout
		<-time.After(p.Timeout)
		cancel()
	}()

	// wait responce
	pres := <-ch
	res := Responce(fmt.Sprintf("%s: through[TO:%s] %s", p.Name, p.Timeout, string(pres.Responce)))
	return res, pres.Error
}

// ProcessSimulation simurate the real process.
// This is an example of a process that can be suspended from Context
type ProcessSimulation struct {
	Name string
	Tick time.Duration
}

func (p *ProcessSimulation) Call(ctx context.Context, task Task) (Responce, error) {
	start := time.Now()
	loopMax, ok := task.Value.(int)
	if !ok {
		return Responce(""), errors.Errorf("Task Value is not int")
	}
	if loopMax < 1 {
		return Responce(""), errors.Errorf("Task.Value must int type and 1 or more")
	}

	// Heavy loop process
	var counter int
	timeout := time.After(task.Timeout)
	for {
		if counter >= loopMax {
			return Responce(fmt.Sprintf("%s: Complete [loop: %d] [%s/%s]", p.Name, counter, time.Since(start), task.Timeout)), nil
		}
		counter++

		// Here is check Context and time out setting.
		select {
		case <-ctx.Done():
			// Add Rollback process when need
			return Responce(fmt.Sprintf("%s: Cancel [%s/%s]", p.Name, time.Since(start), task.Timeout)), ctx.Err()
		case <-timeout:
			// Add Rollback process when need
			return Responce(fmt.Sprintf("%s: Timeout [%s/%s]", p.Name, time.Since(start), task.Timeout)), errors.Errorf("Timeout by Simulation")
		case <-time.After(p.Tick):
		}
	}
}
