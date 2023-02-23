package supervisor

import (
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSupervisor(t *testing.T) {
	address := "127.0.0.1:11088"
	config := DefaultConfig()
	config.ListenAddress = address
	sv := NewSupervisor(config)
	sv.Serve()
	defer func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		sv.Close(wg)
		wg.Wait()
	}()

	conn, err := net.Dial("tcp", address)
	assert.Nil(t, err)
	id := []byte(conn.LocalAddr().String())
	buf := make([]byte, len(id))
	assert.Nil(t, conn.SetReadDeadline(time.Now().Add(time.Second*3)))
	n, err := conn.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(id), n)
	assert.Equal(t, id, buf)

	assert.Nil(t, conn.Close())
}
