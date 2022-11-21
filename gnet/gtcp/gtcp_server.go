package gtcp

import (
    "context"
    "crypto/tls"
    "fmt"
    "github.com/camry/g/gerrors/gcode"
    "github.com/camry/g/gerrors/gerror"
    "net"
    "strings"
    "sync"
)

const (
    // FreePortAddress 使用随机免费端口标记服务器监听。
    FreePortAddress = ":0"
)

// Server 定义 TCP 服务包装器。
type Server struct {
    mu        sync.Mutex   // 用于 Server.listen 并发安全。
    listen    net.Listener // 网络监听器。
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
    return srv
}

// Close 关闭 TCP 服务器。
func (s *Server) Close() error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if s.listen == nil {
        return nil
    }
    return s.listen.Close()
}

// Run 启动 TCP 服务器。
func (s *Server) Run(ctx context.Context) (err error) {
    if s.handler == nil {
        err = gerror.NewCode(gcode.CodeMissingConfiguration, "start running failed: socket handler not defined")
        return
    }
    if s.tlsConfig != nil {
        // TLS Server
        s.mu.Lock()
        s.listen, err = tls.Listen("tcp", s.address, s.tlsConfig)
        s.mu.Unlock()
        if err != nil {
            err = gerror.Wrapf(err, `tls.Listen failed for address "%s"`, s.address)
            return
        }
    } else {
        // Normal Server
        var tcpAddr *net.TCPAddr
        if tcpAddr, err = net.ResolveTCPAddr("tcp", s.address); err != nil {
            err = gerror.Wrapf(err, `net.ResolveTCPAddr failed for address "%s"`, s.address)
            return err
        }
        s.mu.Lock()
        s.listen, err = net.ListenTCP("tcp", tcpAddr)
        s.mu.Unlock()
        if err != nil {
            err = gerror.Wrapf(err, `net.ListenTCP failed for address "%s"`, s.address)
            return err
        }
    }
    // Listening loop.
    for {
        var conn net.Conn
        if conn, err = s.listen.Accept(); err != nil {
            err = gerror.Wrapf(err, `Listener.Accept failed`)
            return err
        } else if conn != nil {
            go s.handler(NewConnByNetConn(conn))
        }
    }
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
    if ln := s.listen; ln != nil {
        return ln.Addr().(*net.TCPAddr).Port
    }
    return -1
}
