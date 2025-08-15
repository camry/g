package gudp

import (
    "io"
    "time"

    "github.com/camry/g/v2/gerrors/gerror"
)

// ClientConn 客户端连接。
type ClientConn struct {
    *localConn
}

// NewClientConn 创建到 `remoteAddress` 的UDP连接
// 可选参数 `localAddress` 指定本地连接地址。
func NewClientConn(remoteAddress string, localAddress ...string) (*ClientConn, error) {
    udpConn, err := NewNetConn(remoteAddress, localAddress...)
    if err != nil {
        return nil, err
    }
    return &ClientConn{
        localConn: &localConn{
            UDPConn: udpConn,
        },
    }, nil
}

// Send 将数据写入远程地址。
func (c *ClientConn) Send(data []byte, retry ...Retry) (err error) {
    for {
        _, err = c.Write(data)
        if err == nil {
            return nil
        }
        // 连接关闭。
        if err == io.EOF {
            return err
        }
        // 重试后仍然失败。
        if len(retry) == 0 || retry[0].Count == 0 {
            return gerror.Wrap(err, `Write data failed`)
        }
        if len(retry) > 0 {
            retry[0].Count--
            if retry[0].Interval == 0 {
                retry[0].Interval = defaultRetryInterval
            }
            time.Sleep(retry[0].Interval)
            continue
        }
        return err
    }
}

// SendRecv 将数据写入连接并阻塞读取响应。
func (c *ClientConn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error) {
    if err := c.Send(data, retry...); err != nil {
        return nil, err
    }
    result, _, err := c.Recv(receive, retry...)
    return result, err
}
