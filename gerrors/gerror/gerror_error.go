package gerror

import (
    "errors"
    "fmt"
    "runtime"
    "strings"

    "github.com/camry/g/v2/gerrors/gcode"
)

// Error 自定义错误对象。
type Error struct {
    error error      // 包装错误。
    stack stack      // 堆栈数组，当创建或包装此错误时记录堆栈信息。
    text  string     // 创建错误时自定义错误文本。
    code  gcode.Code // 如有必要，错误码。
}

const (
    // stackFilterKeyLocal 过滤当前错误模块路径的键。
    stackFilterKeyLocal = "/gerrors/gerror/gerror"
)

var (
    // goRootForFilter 用于开发环境中的堆栈过滤。
    goRootForFilter = runtime.GOROOT()
)

func init() {
    if goRootForFilter != "" {
        goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
    }
}

// Error 实现错误的接口，它将所有错误返回为字符串。
func (err *Error) Error() string {
    if err == nil {
        return ""
    }
    errStr := err.text
    if errStr == "" && err.code != nil {
        errStr = err.code.Message()
    }
    if err.error != nil {
        if errStr != "" {
            errStr += ": "
        }
        errStr += err.error.Error()
    }
    return errStr
}

// Cause 获取根错误 error。
func (err *Error) Cause() error {
    if err == nil {
        return nil
    }
    loop := err
    for loop != nil {
        if loop.error != nil {
            if e, ok := loop.error.(*Error); ok {
                // 内部自定义错误。
                loop = e
            } else if e, ok := loop.error.(ICause); ok {
                // 实现 ApiCause 接口的其他错误。
                return e.Cause()
            } else {
                return loop.error
            }
        } else {
            // return loop
            //
            // 参考案例 https://github.com/pkg/errors。
            return errors.New(loop.text)
        }
    }
    return nil
}

// Current 获取当前 error。
// 如果当前错误是 nil, 则返回 nil。
func (err *Error) Current() error {
    if err == nil {
        return nil
    }
    return &Error{
        error: nil,
        stack: err.stack,
        text:  err.text,
        code:  err.code,
    }
}

// Unwrap 获取下一层 error。
func (err *Error) Unwrap() error {
    if err == nil {
        return nil
    }
    return err.error
}

// Equal 错误对象比较。
// 如果它们的 `code` 和 `text` 都相同，则认为错误相同。
func (err *Error) Equal(target error) bool {
    if err == target {
        return true
    }
    if err.code != Code(target) {
        return false
    }
    if err.text != fmt.Sprintf(`%-s`, target) {
        return false
    }
    return true
}

// Is 当前错误 `err` 的链接错误中是否包含错误 `target`。
func (err *Error) Is(target error) bool {
    if Equal(err, target) {
        return true
    }
    nextErr := err.Unwrap()
    if nextErr == nil {
        return false
    }
    if Equal(nextErr, target) {
        return true
    }
    if e, ok := nextErr.(IIs); ok {
        return e.Is(target)
    }
    return false
}
