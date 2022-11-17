package gerror

import (
	"github.com/camry/g/gerrors/gcode"
)

// IIs Is 接口。
type IIs interface {
	Error() string
	Is(target error) bool
}

// IEqual Equal 接口。
type IEqual interface {
	Error() string
	Equal(target error) bool
}

// ICode Code 接口。
type ICode interface {
	Error() string
	Code() gcode.Code
}

// IStack Stack 接口。
type IStack interface {
	Error() string
	Stack() string
}

// ICause Cause 接口。
type ICause interface {
	Error() string
	Cause() error
}

// ICurrent Current 接口。
type ICurrent interface {
	Error() string
	Current() error
}

// IUnwrap Unwrap 接口。
type IUnwrap interface {
	Error() string
	Unwrap() error
}

const (
	// stackFilterKeyForG 过滤 G 模块路径堆栈。
	stackFilterKeyForG = "github.com/camry/g/"

	// separatorSpace 空间分隔符。
	separatorSpace = ", "
)

var (
	// IsUsingBriefStack 简短错误堆栈的开关。（默认为true）
	IsUsingBriefStack = true
)
