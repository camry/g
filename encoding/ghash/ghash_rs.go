package ghash

// RS 实现经典的 RS 哈希算法32位。
func RS(str []byte) uint32 {
    var (
        b    uint32 = 378551
        a    uint32 = 63689
        hash uint32 = 0
    )
    for i := 0; i < len(str); i++ {
        hash = hash*a + uint32(str[i])
        a *= b
    }
    return hash
}

// RS64 实现经典的 RS 哈希算法64位。
func RS64(str []byte) uint64 {
    var (
        b    uint64 = 378551
        a    uint64 = 63689
        hash uint64 = 0
    )
    for i := 0; i < len(str); i++ {
        hash = hash*a + uint64(str[i])
        a *= b
    }
    return hash
}
