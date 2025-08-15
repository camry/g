package gerror

import (
    "fmt"

    "github.com/camry/g/v2/gerrors/gcode"
)

// New 用于创建一个自定义文本错误信息的 error 对象，并包含堆栈信息。
func New(text string) error {
    return &Error{
        stack: callers(),
        text:  text,
        code:  gcode.CodeNil,
    }
}

// Newf 用于创建一个自定义文本错误带参数信息的 error 对象，并包含堆栈信息。
func Newf(format string, args ...any) error {
    return &Error{
        stack: callers(),
        text:  fmt.Sprintf(format, args...),
        code:  gcode.CodeNil,
    }
}

// NewSkip 用于创建一个自定义错误信息的 error 对象，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。高级功能，一般开发者很少用得到。
// 参数 `skip` 指定堆栈跳过的层数。
func NewSkip(skip int, text string) error {
    return &Error{
        stack: callers(skip),
        text:  text,
        code:  gcode.CodeNil,
    }
}

// NewSkipf 用于创建一个自定义错误信息的error对象，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。
// 参数 `skip` 指定堆栈跳过的层数。
func NewSkipf(skip int, format string, args ...any) error {
    return &Error{
        stack: callers(skip),
        text:  fmt.Sprintf(format, args...),
        code:  gcode.CodeNil,
    }
}

// Wrap 用于包裹其他错误 error 对象，构造成多级的错误信息，包含堆栈信息。
// 注：它不会丢失包装错误的错误码，因为它从中继承了错误码。
func Wrap(err error, text string) error {
    if err == nil {
        return nil
    }
    return &Error{
        error: err,
        stack: callers(),
        text:  text,
        code:  Code(err),
    }
}

// Wrapf 用于包裹其他错误 error 对象，构造成多级的错误信息，包含堆栈信息。
func Wrapf(err error, format string, args ...any) error {
    if err == nil {
        return nil
    }
    return &Error{
        error: err,
        stack: callers(),
        text:  fmt.Sprintf(format, args...),
        code:  Code(err),
    }
}

// WrapSkip 用于包裹其他错误 error 对象，构造成多级的错误信息，包含堆栈信息，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。
func WrapSkip(skip int, err error, text string) error {
    if err == nil {
        return nil
    }
    return &Error{
        error: err,
        stack: callers(skip),
        text:  text,
        code:  Code(err),
    }
}

// WrapSkipf 用于包裹其他错误 error 对象，构造成多级的错误信息，包含堆栈信息，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。
func WrapSkipf(skip int, err error, format string, args ...any) error {
    if err == nil {
        return nil
    }
    return &Error{
        error: err,
        stack: callers(skip),
        text:  fmt.Sprintf(format, args...),
        code:  Code(err),
    }
}
