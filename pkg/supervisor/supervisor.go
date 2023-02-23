package supervisor

import (
	"sync"

	"gcm/pkg/logger"
	"gcm/pkg/services"
)

type (
	Supervisor struct {
		config *Config

		listener   *services.Listener
		done       chan struct{}
		listenChan chan services.ListenChannel
		deleteChan chan string
		wg         sync.WaitGroup
		services   map[string]services.Service
	}

	Config struct {
		ListenAddress string
	}
)

var (
	log = logger.Module("supervisor")
)

func NewSupervisor(config *Config) *Supervisor {
	return &Supervisor{
		config:     config,
		done:       make(chan struct{}),
		listenChan: make(chan services.ListenChannel),
		deleteChan: make(chan string),
		services:   make(map[string]services.Service),
	}
}

func (s *Supervisor) Close(wg *sync.WaitGroup) {
	defer wg.Done()

	// 关闭 Listener
	wg.Add(1)
	s.listener.Close(wg)

	// 关闭 proxy
	wg.Add(len(s.services))
	for _, proxy := range s.services {
		proxy.Close(wg)
	}

	// 关闭 Supervisor
	close(s.done)
	s.wg.Wait()
}

func (s *Supervisor) Serve() {
	s.listener = services.NewListener(s.config.ListenAddress, s.listenChan)
	s.run()
}

func (s *Supervisor) run() {
	// 启动 Listener
	err := s.listener.Run()
	if err != nil {
		log.Info(err.Error())
		return
	}
	// 处理事件
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.done:
				log.Info("stopping supervisor")
				return
			case conn := <-s.listenChan:
				clientId := conn.Conn().RemoteAddr().String()
				conn.Conn().Write([]byte(clientId))
			case id := <-s.deleteChan:
				s.deleteClient(id)
			}
		}
	}()
}

func (s *Supervisor) deleteClient(clientId string) {
	if service, ok := s.services[clientId]; ok {
		delete(s.services, clientId)
		go func() {
			wg := sync.WaitGroup{}
			wg.Add(1)
			service.Close(&wg)
		}()
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (c *Config) Validate() error {
	return nil
}
