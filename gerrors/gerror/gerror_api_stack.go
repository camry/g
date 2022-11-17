package gerror

import (
    "runtime"
)

// stack 代表堆栈程序计数器。
type stack []uintptr

const (
    // maxStackDepth 标记最大堆栈深度的错误后轨迹。
    maxStackDepth = 64
)

// Cause 获取根错误 error。
// 如果 `err` 是 nil, 则返回 nil。
func Cause(err error) error {
    if err == nil {
        return nil
    }
    if e, ok := err.(ICause); ok {
        return e.Cause()
    }
    if e, ok := err.(IUnwrap); ok {
        return Cause(e.Unwrap())
    }
    return err
}

// Stack 获取堆栈信息。
// 如果 `err` 是 nil, 则返回空字符串。
// 如果 `err` 不支持堆栈，它将直接返回错误字符串。
func Stack(err error) string {
    if err == nil {
        return ""
    }
    if e, ok := err.(IStack); ok {
        return e.Stack()
    }
    return err.Error()
}

// Current 获取当前 error。
// 如果当前错误是 nil, 则返回 nil。
func Current(err error) error {
    if err == nil {
        return nil
    }
    if e, ok := err.(ICurrent); ok {
        return e.Current()
    }
    return err
}

// Unwrap 获取下一层 error。
// 如果当前级别错误或下一个级别错误为 nil，则返回 nil。
func Unwrap(err error) error {
    if err == nil {
        return nil
    }
    if e, ok := err.(IUnwrap); ok {
        return e.Unwrap()
    }
    return nil
}

// HasStack 判断错误是否带堆栈，实现 `gerror.IStack` 接口。
func HasStack(err error) bool {
    _, ok := err.(IStack)
    return ok
}

// Equal 错误对象比较。
func Equal(err, target error) bool {
    if err == target {
        return true
    }
    if e, ok := err.(IEqual); ok {
        return e.Equal(target)
    }
    if e, ok := target.(IEqual); ok {
        return e.Equal(err)
    }
    return false
}

// Is 包含判断。
func Is(err, target error) bool {
    if e, ok := err.(IIs); ok {
        return e.Is(target)
    }
    return false
}

// HasError 包含判断，`gerror.Is` 的别名。
func HasError(err, target error) bool {
    return Is(err, target)
}

// callers 获取堆栈信息。
func callers(skip ...int) stack {
    var (
        pcs [maxStackDepth]uintptr
        n   = 3
    )
    if len(skip) > 0 {
        n += skip[0]
    }
    return pcs[:runtime.Callers(n, pcs[:])]
}
