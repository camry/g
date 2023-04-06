package glog

import "io"

type writerWrapper struct {
    helper *Helper
    level  Level
}

type WriterOptionFn func(w *writerWrapper)

// WithWriterLevel 配置 Writer 包装器等级。
func WithWriterLevel(level Level) WriterOptionFn {
    return func(w *writerWrapper) {
        w.level = level
    }
}

// WithWriteMessageKey 配置 Writer 包装器助手消息 KEY。
func WithWriteMessageKey(key string) WriterOptionFn {
    return func(w *writerWrapper) {
        w.helper.msgKey = key
    }
}

// NewWriter 新建 Writer。
func NewWriter(logger Logger, opts ...WriterOptionFn) io.Writer {
    ww := &writerWrapper{
        helper: NewHelper(logger, WithMessageKey(DefaultMessageKey)),
        level:  LevelInfo, // 默认等级
    }
    for _, opt := range opts {
        opt(ww)
    }
    return ww
}

func (ww *writerWrapper) Write(p []byte) (int, error) {
    ww.helper.Log(ww.level, ww.helper.msgKey, string(p))
    return 0, nil
}
