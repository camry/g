package gutil

// InArray 将在数组中搜索任意类型的元素，返回匹配元素的布尔值。
// element 是要搜索的元素，collection 是要搜索的值的切片。
func InArray[T comparable](element T, collection []T) bool {
    for _, item := range collection {
        if item == element {
            return true
        }
    }
    return false
}
