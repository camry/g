package glog

import (
    "context"
    "fmt"
    "os"
    "sync"
)

// globalLogger 被设计为当前进程中的全局日志记录器。
var global = &loggerAppliance{}

// loggerAppliance 是 `Logger` 的代理器，使 logger 更改将影响所有子 logger.
type loggerAppliance struct {
    lock sync.Mutex
    Logger
}

// init 初始化全局默认 Logger。
func init() {
    global.SetLogger(DefaultLogger)
}

// SetLogger 配置 Logger。
func (a *loggerAppliance) SetLogger(in Logger) {
    a.lock.Lock()
    defer a.lock.Unlock()
    a.Logger = in
}

// SetLogger 应该在任何其他日志调用之前调用。
// 而且它不是线程安全的。
func SetLogger(logger Logger) {
    global.SetLogger(logger)
}

// GetLogger 将全局日志记录器器设备作为当前进程中的记录器返回。
func GetLogger() Logger {
    return global
}

// Log 按级别和键值打印日志。
func Log(level Level, keyvals ...any) {
    _ = global.Log(level, keyvals...)
}

// Context 配置上下文 logger。
func Context(ctx context.Context) *Helper {
    return NewHelper(WithContext(ctx, global.Logger))
}

// Debug 打印调试级别的日志。
func Debug(a ...any) {
    _ = global.Log(LevelDebug, DefaultMessageKey, fmt.Sprint(a...))
}

// Debugf 按 fmt.Sprintf 格式打印调试级别的日志。
func Debugf(format string, a ...any) {
    _ = global.Log(LevelDebug, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Debugw 按键值对打印调试级别的日志。
func Debugw(keyvals ...any) {
    _ = global.Log(LevelDebug, keyvals...)
}

// Info 打印信息级别的日志。
func Info(a ...any) {
    _ = global.Log(LevelInfo, DefaultMessageKey, fmt.Sprint(a...))
}

// Infof 按 fmt.Sprintf 格式打印信息级别的日志。
func Infof(format string, a ...any) {
    _ = global.Log(LevelInfo, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Infow 按键值对打印信息级别的日志。
func Infow(keyvals ...any) {
    _ = global.Log(LevelInfo, keyvals...)
}

// Warn 打印警告级别的日志。
func Warn(a ...any) {
    _ = global.Log(LevelWarn, DefaultMessageKey, fmt.Sprint(a...))
}

// Warnf 按 fmt.Sprintf 格式打印警告级别的日志。
func Warnf(format string, a ...any) {
    _ = global.Log(LevelWarn, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Warnw 按键值对打印警告级别的日志。
func Warnw(keyvals ...any) {
    _ = global.Log(LevelWarn, keyvals...)
}

// Error 打印错误级别的日志。
func Error(a ...any) {
    _ = global.Log(LevelError, DefaultMessageKey, fmt.Sprint(a...))
}

// Errorf 按 fmt.Sprintf 格式打印错误级别的日志。
func Errorf(format string, a ...any) {
    _ = global.Log(LevelError, DefaultMessageKey, fmt.Sprintf(format, a...))
}

// Errorw 按键值对打印错误级别的日志。
func Errorw(keyvals ...any) {
    _ = global.Log(LevelError, keyvals...)
}

// Fatal 打印致命级别的日志。
func Fatal(a ...any) {
    _ = global.Log(LevelFatal, DefaultMessageKey, fmt.Sprint(a...))
    os.Exit(1)
}

// Fatalf 按 fmt.Sprintf 格式打印致命级别的日志。
func Fatalf(format string, a ...any) {
    _ = global.Log(LevelFatal, DefaultMessageKey, fmt.Sprintf(format, a...))
    os.Exit(1)
}

// Fatalw 按键值对打印致命级别的日志。
func Fatalw(keyvals ...any) {
    _ = global.Log(LevelFatal, keyvals...)
    os.Exit(1)
}
