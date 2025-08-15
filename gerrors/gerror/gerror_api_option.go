package gerror

import "github.com/camry/g/v2/gerrors/gcode"

// Option 自定义的错误创建。
type Option struct {
    Error error      // 包装错误（如果有）。
    Stack bool       // 是否将堆栈信息记录到错误中。
    Text  string     // 错误文本，由 New* 函数创建
    Code  gcode.Code // 如有必要，错误码。
}

// NewOption 用于自定义配置的错误对象创建。
func NewOption(option Option) error {
    err := &Error{
        error: option.Error,
        text:  option.Text,
        code:  option.Code,
    }
    if option.Stack {
        err.stack = callers()
    }
    return err
}
