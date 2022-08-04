package g

type (
    Map        = map[string]any    // Map 是常用字典类型 map[string]any 的别名。
    MapAnyAny  = map[any]any       // MapAnyAny 是常用字典类型 map[any]any 的别名。
    MapAnyStr  = map[any]string    // MapAnyStr 是常用字典类型 map[any]string 的别名。
    MapAnyInt  = map[any]int       // MapAnyInt 是常用字典类型 map[any]int 的别名。
    MapStrAny  = map[string]any    // MapStrAny 是常用字典类型 map[string]any 的别名。
    MapStrStr  = map[string]string // MapStrStr 是常用字典类型 map[string]string 的别名。
    MapStrInt  = map[string]int    // MapStrInt 是常用字典类型 map[string]int 的别名。
    MapIntAny  = map[int]any       // MapIntAny 是常用字典类型 map[int]any 的别名。
    MapIntStr  = map[int]string    // MapIntStr 是常用字典类型 map[int]string 的别名。
    MapIntInt  = map[int]int       // MapIntInt 是常用字典类型 map[int]int 的别名。
    MapAnyBool = map[any]bool      // MapAnyBool 是常用字典类型 map[any]bool 的别名。
    MapStrBool = map[string]bool   // MapStrBool 是常用字典类型 map[string]bool 的别名。
    MapIntBool = map[int]bool      // MapIntBool 是常用字典类型 map[int]bool 的别名。
)

type (
    List        = []Map        // List 是常用切片类型 []Map 的别名。
    ListAnyAny  = []MapAnyAny  // ListAnyAny 是常用切片类型 []MapAnyAny 的别名。
    ListAnyStr  = []MapAnyStr  // ListAnyStr 是常用切片类型 []MapAnyStr 的别名。
    ListAnyInt  = []MapAnyInt  // ListAnyInt 是常用切片类型 []MapAnyInt 的别名。
    ListStrAny  = []MapStrAny  // ListStrAny 是常用切片类型 []MapStrAny 的别名。
    ListStrStr  = []MapStrStr  // ListStrStr 是常用切片类型 []MapStrStr 的别名。
    ListStrInt  = []MapStrInt  // ListStrInt 是常用切片类型 []MapStrInt 的别名。
    ListIntAny  = []MapIntAny  // ListIntAny 是常用切片类型 []MapIntAny 的别名。
    ListIntStr  = []MapIntStr  // ListIntStr 是常用切片类型 []MapIntStr 的别名。
    ListIntInt  = []MapIntInt  // ListIntInt 是常用切片类型 []MapIntInt 的别名。
    ListAnyBool = []MapAnyBool // ListAnyBool 是常用切片类型 []MapAnyBool 的别名。
    ListStrBool = []MapStrBool // ListStrBool 是常用切片类型 []MapStrBool 的别名。
    ListIntBool = []MapIntBool // ListIntBool 是常用切片类型 []MapIntBool 的别名。
)

type (
    Slice    = []any    // Slice 是常用切片类型 []any 的别名。
    SliceAny = []any    // SliceAny 是常用切片类型 []any 的别名。
    SliceStr = []string // SliceStr 是常用切片类型 []string 的别名。
    SliceInt = []int    // SliceInt 是常用切片类型 []int 的别名。
)

type (
    Array    = []any    // Array 是常用切片类型 []any 的别名。
    ArrayAny = []any    // ArrayAny 是常用切片类型 []any 的别名。
    ArrayStr = []string // ArrayStr 是常用切片类型 []string 的别名。
    ArrayInt = []int    // ArrayInt 是常用切片类型 []int 的别名。
)
