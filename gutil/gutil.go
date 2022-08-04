package gutil

import "reflect"

// InArray 将在数组中搜索任意类型的元素，返回匹配元素的布尔值。
// needle 是要搜索的元素，haystack 是要搜索的值的切片。
func InArray(needle any, haystack any) bool {
    switch reflect.TypeOf(haystack).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(haystack)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
                return true
            }
        }
    }
    return false
}
