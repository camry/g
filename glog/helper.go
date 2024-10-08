package glog

import (
    "context"
    "fmt"
    "os"
)

// DefaultMessageKey 默认消息KEY。
var DefaultMessageKey = "msg"

// Option 日志助手选项。
type Option func(*Helper)

// Helper 是一个日志助手。
type Helper struct {
    logger  Logger
    msgKey  string
    sprint  func(...any) string
    sprintf func(format string, a ...any) string
}

// WithMessageKey 配置消息KEY。
func WithMessageKey(k string) Option {
    return func(opts *Helper) {
        opts.msgKey = k
    }
}

// WithSprint 配置 sprint。
func WithSprint(sprint func(...any) string) Option {
    return func(opts *Helper) {
        opts.sprint = sprint
    }
}

// WithSprintf 配置 sprintf。
func WithSprintf(sprintf func(format string, a ...any) string) Option {
    return func(opts *Helper) {
        opts.sprintf = sprintf
    }
}

// NewHelper 新建一个日志助手。
func NewHelper(logger Logger, opts ...Option) *Helper {
    options := &Helper{
        msgKey: DefaultMessageKey,
        logger: logger,
    }
    for _, o := range opts {
        o(options)
    }
    return options
}

// WithContext 返回被 h 的浅副本改变的上下文 ctx, 提供的 ctx 不能为空。
func (h *Helper) WithContext(ctx context.Context) *Helper {
    return &Helper{
        msgKey: h.msgKey,
        logger: WithContext(ctx, h.logger),
    }
}

// Enabled 如果给定级别高于此级别，则返回 true。它委托给底层 *Filter。
func (h *Helper) Enabled(level Level) bool {
    if l, ok := h.logger.(*Filter); ok {
        return level >= l.level
    }
    return true
}

// Log 按级别和键值打印日志。
func (h *Helper) Log(level Level, keyvals ...any) {
    _ = h.logger.Log(level, keyvals...)
}

// Debug 打印调试级别的日志。
func (h *Helper) Debug(a ...any) {
    if !h.Enabled(LevelDebug) {
        return
    }
    _ = h.logger.Log(LevelDebug, h.msgKey, fmt.Sprint(a...))
}

// Debugf 按 fmt.Sprintf 格式打印调试级别的日志。
func (h *Helper) Debugf(format string, a ...any) {
    if !h.Enabled(LevelDebug) {
        return
    }
    _ = h.logger.Log(LevelDebug, h.msgKey, fmt.Sprintf(format, a...))
}

// Debugw 按键值对打印调试级别的日志。
func (h *Helper) Debugw(keyvals ...any) {
    _ = h.logger.Log(LevelDebug, keyvals...)
}

// Info 打印信息级别的日志。
func (h *Helper) Info(a ...any) {
    if !h.Enabled(LevelInfo) {
        return
    }
    _ = h.logger.Log(LevelInfo, h.msgKey, fmt.Sprint(a...))
}

// Infof 按 fmt.Sprintf 格式打印信息级别的日志。
func (h *Helper) Infof(format string, a ...any) {
    if !h.Enabled(LevelInfo) {
        return
    }
    _ = h.logger.Log(LevelInfo, h.msgKey, fmt.Sprintf(format, a...))
}

// Infow 按键值对打印信息级别的日志。
func (h *Helper) Infow(keyvals ...any) {
    _ = h.logger.Log(LevelInfo, keyvals...)
}

// Warn 打印警告级别的日志。
func (h *Helper) Warn(a ...any) {
    if !h.Enabled(LevelWarn) {
        return
    }
    _ = h.logger.Log(LevelWarn, h.msgKey, fmt.Sprint(a...))
}

// Warnf 按 fmt.Sprintf 格式打印警告级别的日志。
func (h *Helper) Warnf(format string, a ...any) {
    if !h.Enabled(LevelWarn) {
        return
    }
    _ = h.logger.Log(LevelWarn, h.msgKey, fmt.Sprintf(format, a...))
}

// Warnw 按键值对打印警告级别的日志。
func (h *Helper) Warnw(keyvals ...any) {
    _ = h.logger.Log(LevelWarn, keyvals...)
}

// Error 打印错误级别的日志。
func (h *Helper) Error(a ...any) {
    if !h.Enabled(LevelError) {
        return
    }
    _ = h.logger.Log(LevelError, h.msgKey, fmt.Sprint(a...))
}

// Errorf 按 fmt.Sprintf 格式打印错误级别的日志。
func (h *Helper) Errorf(format string, a ...any) {
    if !h.Enabled(LevelError) {
        return
    }
    _ = h.logger.Log(LevelError, h.msgKey, fmt.Sprintf(format, a...))
}

// Errorw 按键值对打印错误级别的日志。
func (h *Helper) Errorw(keyvals ...any) {
    _ = h.logger.Log(LevelError, keyvals...)
}

// Fatal 打印致命级别的日志。
func (h *Helper) Fatal(a ...any) {
    _ = h.logger.Log(LevelFatal, h.msgKey, fmt.Sprint(a...))
    os.Exit(1)
}

// Fatalf 按 fmt.Sprintf 格式打印致命级别的日志。
func (h *Helper) Fatalf(format string, a ...any) {
    _ = h.logger.Log(LevelFatal, h.msgKey, fmt.Sprintf(format, a...))
    os.Exit(1)
}

// Fatalw 按键值对打印致命级别的日志。
func (h *Helper) Fatalw(keyvals ...any) {
    _ = h.logger.Log(LevelFatal, keyvals...)
    os.Exit(1)
}
