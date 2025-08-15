package gudp

import (
    "fmt"
    "net"
    "strings"
    "sync"

    "github.com/camry/g/v2/gerrors/gcode"
    "github.com/camry/g/v2/gerrors/gerror"
)

const (
    // FreePortAddress 标记服务器使用随机空闲端口侦听。
    FreePortAddress = ":0"
)

const (
    defaultServer = "default"
)

// Server is the UDP server.
type Server struct {
    // 用于服务器端。监听并发安全性。
    mu sync.Mutex

    // UDP 服务器连接对象。
    conn *ServerConn

    // UDP 服务器监听地址。
    address string

    // UDP 连接的处理程序。
    handler ServerHandler
}

// ServerHandler 处理所有服务器连接。
type ServerHandler func(conn *ServerConn)

// NewServer 创建并返回UDP服务器。
func NewServer(address string, handler ServerHandler) *Server {
    s := &Server{
        address: address,
        handler: handler,
    }
    return s
}

// SetAddress 设置UDP服务器地址。
func (s *Server) SetAddress(address string) {
    s.address = address
}

// SetHandler 设置UDP服务器的连接处理程序。
func (s *Server) SetHandler(handler ServerHandler) {
    s.handler = handler
}

// Close 关闭连接。
// 它将使服务器立即关闭。
func (s *Server) Close() (err error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    err = s.conn.Close()
    if err != nil {
        err = gerror.Wrap(err, "connection failed")
    }
    return
}

// Run 开始监听UDP连接。
func (s *Server) Run() error {
    if s.handler == nil {
        return gerror.NewCode(
            gcode.CodeMissingConfiguration,
            "start running failed: socket handler not defined",
        )
    }
    addr, err := net.ResolveUDPAddr("udp", s.address)
    if err != nil {
        err = gerror.Wrapf(err, `net.ResolveUDPAddr failed for address "%s"`, s.address)
        return err
    }
    listenedConn, err := net.ListenUDP("udp", addr)
    if err != nil {
        err = gerror.Wrapf(err, `net.ListenUDP failed for address "%s"`, s.address)
        return err
    }
    s.mu.Lock()
    s.conn = NewServerConn(listenedConn)
    s.mu.Unlock()
    s.handler(s.conn)
    return nil
}

// GetListenedAddress 获取当前服务器监听地址。
func (s *Server) GetListenedAddress() string {
    if !strings.Contains(s.address, FreePortAddress) {
        return s.address
    }
    var (
        address      = s.address
        listenedPort = s.GetListenedPort()
    )
    address = strings.Replace(address, FreePortAddress, fmt.Sprintf(`:%d`, listenedPort), -1)
    return address
}

// GetListenedPort 获取当前服务器监听端口。
func (s *Server) GetListenedPort() int {
    s.mu.Lock()
    defer s.mu.Unlock()
    if ln := s.conn; ln != nil {
        return ln.LocalAddr().(*net.UDPAddr).Port
    }
    return -1
}
