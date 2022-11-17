package gcode

// Code 通用错误代码接口定义。
type Code interface {
    // Code 错误码。
    Code() int
    // Message 错误码简短信息。
    Message() string
    // Detail 错误码详细信息。
    Detail() any
}

// ================================================================================================================
// 公共错误码定义。
// 保留内部错误码: code < 1000。
// ================================================================================================================

var (
    CodeNil                      = localCode{-1, "", nil}
    CodeOK                       = localCode{0, "OK", nil}
    CodeInternalError            = localCode{50, "Internal Error", nil}
    CodeValidationFailed         = localCode{51, "Validation Failed", nil}
    CodeDbOperationError         = localCode{52, "Database Operation Error", nil}
    CodeInvalidParameter         = localCode{53, "Invalid Parameter", nil}
    CodeMissingParameter         = localCode{54, "Missing Parameter", nil}
    CodeInvalidOperation         = localCode{55, "Invalid Operation", nil}
    CodeInvalidConfiguration     = localCode{56, "Invalid Configuration", nil}
    CodeMissingConfiguration     = localCode{57, "Missing Configuration", nil}
    CodeNotImplemented           = localCode{58, "Not Implemented", nil}
    CodeNotSupported             = localCode{59, "Not Supported", nil}
    CodeOperationFailed          = localCode{60, "Operation Failed", nil}
    CodeNotAuthorized            = localCode{61, "Not Authorized", nil}
    CodeSecurityReason           = localCode{62, "Security Reason", nil}
    CodeServerBusy               = localCode{63, "Server Is Busy", nil}
    CodeUnknown                  = localCode{64, "Unknown Error", nil}
    CodeNotFound                 = localCode{65, "Not Found", nil}
    CodeInvalidRequest           = localCode{66, "Invalid Request", nil}
    CodeBusinessValidationFailed = localCode{300, "Business Validation Failed", nil}
)

// New 创建并返回错误码。
// 注：返回错误码 Code 接口对象。
func New(code int, message string, detail interface{}) Code {
    return localCode{
        code:    code,
        message: message,
        detail:  detail,
    }
}

// WithCode 根据指定的 Code 错误码创建并返回一个新的错误码。
func WithCode(code Code, detail interface{}) Code {
    return localCode{
        code:    code.Code(),
        message: code.Message(),
        detail:  detail,
    }
}
