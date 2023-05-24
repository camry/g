package gcode

import "fmt"

// localCode 内部错误码实现。
type localCode struct {
    code    int    // Error code, usually an integer.
    message string // Brief message for this error code.
    detail  any    // As type of interface, it is mainly designed as an extension field for error code.
}

// Code 错误码。
func (c localCode) Code() int {
    return c.code
}

// Message 错误码简短信息。
func (c localCode) Message() string {
    return c.message
}

// Detail 错误码详细信息。
func (c localCode) Detail() any {
    return c.detail
}

// String 错误码字符串。
func (c localCode) String() string {
    if c.detail != nil {
        return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
    }
    if c.message != "" {
        return fmt.Sprintf(`%d:%s`, c.code, c.message)
    }
    return fmt.Sprintf(`%d`, c.code)
}
