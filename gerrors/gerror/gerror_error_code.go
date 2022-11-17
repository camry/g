package gerror

import (
    "github.com/camry/g/gerrors/gcode"
)

// Code 获取错误码。
// 如果没有错误代码，则返回 `gcode.CodeNil`。
func (err *Error) Code() gcode.Code {
    if err == nil {
        return gcode.CodeNil
    }
    if err.code == gcode.CodeNil {
        return Code(err.Unwrap())
    }
    return err.code
}

// SetCode 使用指定 `code` 更新内部 `code` 。
func (err *Error) SetCode(code gcode.Code) {
    if err == nil {
        return
    }
    err.code = code
}
