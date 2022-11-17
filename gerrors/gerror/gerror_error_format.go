package gerror

import (
    "fmt"
    "io"
)

// Format 根据 `fmt.Formatter` 接口格式化。
//
// %v, %s   : 打印所有错误字符串；
// %-v, %-s : 打印当前级别错误字符串；
// %+s      : 打印完整堆栈错误列表；
// %+v      : 打印错误字符串和完整堆栈错误列表
func (err *Error) Format(s fmt.State, verb rune) {
    switch verb {
    case 's', 'v':
        switch {
        case s.Flag('-'):
            if err.text != "" {
                _, _ = io.WriteString(s, err.text)
            } else {
                _, _ = io.WriteString(s, err.Error())
            }
        case s.Flag('+'):
            if verb == 's' {
                _, _ = io.WriteString(s, err.Stack())
            } else {
                _, _ = io.WriteString(s, err.Error()+"\n"+err.Stack())
            }
        default:
            _, _ = io.WriteString(s, err.Error())
        }
    }
}
