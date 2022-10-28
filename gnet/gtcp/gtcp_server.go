package gtcp

import (
    "context"
    "crypto/tls"
    "errors"
    "fmt"
    "net"
    "sync"
)

// Server 定义 TCP 服务包装器。
type Server struct {
    err       error
    mu        sync.Mutex   // 用于 Server.listener 并发安全。
    listener  net.Listener // 网络监听器。
    network   string       // 服务器网络协议。
    address   string       // 服务器监听地址。
    handler   func(*Conn)  // 连接处理器。
    tlsConfig *tls.Config  // TLS 配置。
}

// NewServer 新建 TCP 服务器。
func NewServer(address string, handler func(*Conn), tlsConfig *tls.Config) *Server {
    srv := &Server{
        network:   "tcp",
        address:   address,
        handler:   handler,
        tlsConfig: tlsConfig,
    }
    srv.err = srv.listen()
    return srv
}

// Run 启动 TCP 服务器。
func (s *Server) Run(ctx context.Context) (err error) {
    if s.err != nil {
        return s.err
    }
    if s.listener == nil {
        err = errors.New("start running failed: socket Listener not defined")
        return
    }
    if s.handler == nil {
        err = errors.New("start running failed: socket handler not defined")
        return
    }
    for {
        var conn net.Conn
        if conn, err = s.listener.Accept(); err != nil {
            err = fmt.Errorf(`Listener.Accept failed`)
            return err
        } else if conn != nil {
            go s.handler(NewConnByNetConn(conn))
        }
    }
}

// Stop 停止 TCP 服务器。
func (s *Server) Stop(ctx context.Context) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if s.listener == nil {
        return nil
    }
    return s.listener.Close()
}

// listen 网络监听。
func (s *Server) listen() (err error) {
    if s.tlsConfig != nil {
        s.mu.Lock()
        s.listener, err = tls.Listen(s.network, s.address, s.tlsConfig)
        s.mu.Unlock()
        if err != nil {
            err = fmt.Errorf(`tls.Listen failed for address "%s"`, s.address)
            return
        }
    } else {
        var tcpAddr *net.TCPAddr
        if tcpAddr, err = net.ResolveTCPAddr(s.network, s.address); err != nil {
            err = fmt.Errorf(`net.ResolveTCPAddr failed for address "%s"`, s.address)
            return
        }
        s.mu.Lock()
        s.listener, err = net.ListenTCP(s.network, tcpAddr)
        s.mu.Unlock()
        if err != nil {
            err = fmt.Errorf(`net.ListenTCP failed for address "%s"`, s.address)
            return
        }
    }
    return nil
}
