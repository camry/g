package ghash

// ELF 实现经典的 ELF 哈希算法32位。
func ELF(str []byte) uint32 {
    var (
        hash uint32
        x    uint32
    )
    for i := 0; i < len(str); i++ {
        hash = (hash << 4) + uint32(str[i])
        if x = hash & 0xF0000000; x != 0 {
            hash ^= x >> 24
            hash &= ^x + 1
        }
    }
    return hash
}

// ELF64 实现经典的 ELF 哈希算法64位。
func ELF64(str []byte) uint64 {
    var (
        hash uint64
        x    uint64
    )
    for i := 0; i < len(str); i++ {
        hash = (hash << 4) + uint64(str[i])
        if x = hash & 0xF000000000000000; x != 0 {
            hash ^= x >> 24
            hash &= ^x + 1
        }
    }
    return hash
}
