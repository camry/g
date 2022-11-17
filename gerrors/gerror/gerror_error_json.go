package gerror

// MarshalJSON 实现 json.Marshal 接口。
// 注：这里不要使用指针作为其接收器。
func (err Error) MarshalJSON() ([]byte, error) {
	return []byte(`"` + err.Error() + `"`), nil
}
