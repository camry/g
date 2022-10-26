# 定时器

## 用法

```go
// Seconds field, required
gcron.New(gcron.WithSeconds())

// Seconds field, optional
gcron.New(cron.WithParser(gcron.NewParser(
gcron.SecondOptional | gcron.Minute | gcron.Hour | gcron.Dom | gcron.Month | gcron.Dow | gcron.Descriptor,
)))
```
