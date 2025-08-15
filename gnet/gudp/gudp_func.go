package gudp

import (
    "net"

    "github.com/camry/g/v2/gerrors/gerror"
)

// NewNetConn 创建并返回指定地址的 *net.UDPConn。
func NewNetConn(remoteAddress string, localAddress ...string) (*net.UDPConn, error) {
    var (
        err        error
        remoteAddr *net.UDPAddr
        localAddr  *net.UDPAddr
        network    = `udp`
    )
    remoteAddr, err = net.ResolveUDPAddr(network, remoteAddress)
    if err != nil {
        return nil, gerror.Wrapf(
            err,
            `net.ResolveUDPAddr failed for network "%s", address "%s"`,
            network, remoteAddress,
        )
    }
    if len(localAddress) > 0 {
        localAddr, err = net.ResolveUDPAddr(network, localAddress[0])
        if err != nil {
            return nil, gerror.Wrapf(
                err,
                `net.ResolveUDPAddr failed for network "%s", address "%s"`,
                network, localAddress[0],
            )
        }
    }
    conn, err := net.DialUDP(network, localAddr, remoteAddr)
    if err != nil {
        return nil, gerror.Wrapf(
            err,
            `net.DialUDP failed for network "%s", local "%s", remote "%s"`,
            network, localAddr.String(), remoteAddr.String(),
        )
    }
    return conn, nil
}

// Send 使用UDP连接将数据写入`address`，然后关闭连接。
// 请注意，它用于短连接使用。
func Send(address string, data []byte, retry ...Retry) error {
    conn, err := NewClientConn(address)
    if err != nil {
        return err
    }
    defer conn.Close()
    return conn.Send(data, retry...)
}

// SendRecv 使用UDP连接将数据写入`address`，读取响应，然后关闭连接。
// 请注意，它用于短连接使用。
func SendRecv(address string, data []byte, receive int, retry ...Retry) ([]byte, error) {
    conn, err := NewClientConn(address)
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    return conn.SendRecv(data, receive, retry...)
}

// MustGetFreePort 执行 GetFreePort，但发生任何错误都会 panic。
// Deprecated: 端口可能在返回后不久就会被使用，请使用 `:0` 作为监听地址，它会要求系统分配一个空闲端口。
func MustGetFreePort() (port int) {
    port, err := GetFreePort()
    if err != nil {
        panic(err)
    }
    return port
}

// GetFreePort 检索并返回一个空闲的端口。
// Deprecated: 端口可能在返回后不久就会被使用，请使用 `:0` 作为监听地址，它会要求系统分配一个空闲端口。
func GetFreePort() (port int, err error) {
    var (
        network = `udp`
        address = `:0`
    )
    resolvedAddr, err := net.ResolveUDPAddr(network, address)
    if err != nil {
        return 0, gerror.Wrapf(
            err,
            `net.ResolveUDPAddr failed for network "%s", address "%s"`,
            network, address,
        )
    }
    l, err := net.ListenUDP(network, resolvedAddr)
    if err != nil {
        return 0, gerror.Wrapf(
            err,
            `net.ListenUDP failed for network "%s", address "%s"`,
            network, resolvedAddr.String(),
        )
    }
    port = l.LocalAddr().(*net.UDPAddr).Port
    _ = l.Close()
    return
}

// GetFreePorts 检索并返回指定数量的空闲端口。
// Deprecated: 端口可能在返回后不久就会被使用，请使用 `:0` 作为监听地址，它会要求系统分配一个空闲端口。
func GetFreePorts(count int) (ports []int, err error) {
    var (
        network = `udp`
        address = `:0`
    )
    for i := 0; i < count; i++ {
        resolvedAddr, err := net.ResolveUDPAddr(network, address)
        if err != nil {
            return nil, gerror.Wrapf(
                err,
                `net.ResolveUDPAddr failed for network "%s", address "%s"`,
                network, address,
            )
        }
        l, err := net.ListenUDP(network, resolvedAddr)
        if err != nil {
            return nil, gerror.Wrapf(
                err,
                `net.ListenUDP failed for network "%s", address "%s"`,
                network, resolvedAddr.String(),
            )
        }
        ports = append(ports, l.LocalAddr().(*net.UDPAddr).Port)
        _ = l.Close()
    }
    return ports, nil
}
