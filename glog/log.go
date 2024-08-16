package glog

import (
    "context"
    "log"
)

// DefaultLogger 默认日志记录器。
var DefaultLogger = NewStdLogger(log.Writer())

// Logger 是一个日志接口。
type Logger interface {
    Log(level Level, keyvals ...any) error
}

type logger struct {
    logger    Logger
    prefix    []any
    hasValuer bool
    ctx       context.Context
}

// Log 按级别和键值打印日志。
func (c *logger) Log(level Level, keyvals ...any) error {
    kvs := make([]any, 0, len(c.prefix)+len(keyvals))
    kvs = append(kvs, c.prefix...)
    if c.hasValuer {
        bindValues(c.ctx, kvs)
    }
    kvs = append(kvs, keyvals...)
    return c.logger.Log(level, kvs...)
}

// With 配置日志字段。
func With(l Logger, kv ...any) Logger {
    c, ok := l.(*logger)
    if !ok {
        return &logger{logger: l, prefix: kv, hasValuer: containsValuer(kv), ctx: context.Background()}
    }
    kvs := make([]any, 0, len(c.prefix)+len(kv))
    kvs = append(kvs, c.prefix...)
    kvs = append(kvs, kv...)
    return &logger{
        logger:    c.logger,
        prefix:    kvs,
        hasValuer: containsValuer(kvs),
        ctx:       c.ctx,
    }
}

// WithContext 返回被 l 的浅副本改变的上下文 ctx, 提供的 ctx 不能为空。
func WithContext(ctx context.Context, l Logger) Logger {
    switch v := l.(type) {
    default:
        return &logger{logger: l, ctx: ctx}
    case *logger:
        lv := *v
        lv.ctx = ctx
        return &lv
    case *Filter:
        fv := *v
        fv.logger = WithContext(ctx, fv.logger)
        return &fv
    }
}
