package gudp

import (
    "context"
    "errors"
    "fmt"
    "net"
    "sync"
)

// Server 定义 UDP 服务器。
type Server struct {
    err     error
    mu      sync.Mutex  // 用于 Server.Conn 并发安全。
    conn    *Conn       // UDP 服务器连接对象。
    network string      // UDP 服务器网络协议。
    address string      // UDP 服务器监听地址。
    handler func(*Conn) // UDP 连接的处理程序。
}

// NewServer 新建 UDP 服务器。
func NewServer(address string, handler func(*Conn)) *Server {
    srv := &Server{
        network: "udp",
        address: address,
        handler: handler,
    }
    srv.err = srv.listen()
    return srv
}

// Run 启动 UDP 服务器。
func (s *Server) Run(ctx context.Context) (err error) {
    if s.err != nil {
        return s.err
    }
    if s.handler == nil {
        err = errors.New("start running failed: socket handler not defined")
        return
    }
    s.handler(s.conn)
    return nil
}

// Close 关闭 UDP 服务器。
func (s *Server) Close(ctx context.Context) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.conn.Close()
}

// listen 网络监听。
func (s *Server) listen() error {
    addr, err := net.ResolveUDPAddr(s.network, s.address)
    if err != nil {
        return fmt.Errorf(`net.ResolveUDPAddr failed for address "%s"`, s.address)
    }
    conn, err := net.ListenUDP(s.network, addr)
    if err != nil {
        return fmt.Errorf(`net.ListenUDP failed for address "%s"`, s.address)
    }
    s.mu.Lock()
    s.conn = NewConnByNetConn(conn)
    s.mu.Unlock()
    return nil
}

// Conn UDP 服务器连接对象
func (s *Server) Conn() *Conn {
    return s.conn
}
