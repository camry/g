package ghash

// BKDR 实现经典的 BKDR 哈希算法32位。
func BKDR(str []byte) uint32 {
    var (
        seed uint32 = 131 // 31 131 1313 13131 131313 etc..
        hash uint32 = 0
    )
    for i := 0; i < len(str); i++ {
        hash = hash*seed + uint32(str[i])
    }
    return hash
}

// BKDR64 实现经典的 BKDR 哈希算法64位。
func BKDR64(str []byte) uint64 {
    var (
        seed uint64 = 131 // 31 131 1313 13131 131313 etc..
        hash uint64 = 0
    )
    for i := 0; i < len(str); i++ {
        hash = hash*seed + uint64(str[i])
    }
    return hash
}
