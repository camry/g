# 日志记录器

## 用法

```go
logger := glog.NewStdLogger(os.Stdout)
// fields & valuer
logger = glog.With(logger,
"service.name", "hellworld",
"service.version", "v1.0.0",
"ts", glog.DefaultTimestamp,
"caller", glog.DefaultCaller,
)
logger.Log(glog.LevelInfo, "key", "value")

// helper
helper := glog.NewHelper(logger)
helper.Log(glog.LevelInfo, "key", "value")
helper.Info("info message")
helper.Infof("info %s", "message")
helper.Infow("key", "value")

// filter
log := glog.NewHelper(glog.NewFilter(logger,
log.FilterLevel(glog.LevelInfo),
log.FilterKey("foo"),
log.FilterValue("bar"),
log.FilterFunc(customFilter),
))
log.Debug("debug log")
log.Info("info log")
log.Warn("warn log")
log.Error("warn log")
```
