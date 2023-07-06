package ghash

// JS 实现经典的 JS 哈希算法32位。
func JS(str []byte) uint32 {
    var hash uint32 = 1315423911
    for i := 0; i < len(str); i++ {
        hash ^= (hash << 5) + uint32(str[i]) + (hash >> 2)
    }
    return hash
}

// JS64 实现经典的 JS 哈希算法64位。
func JS64(str []byte) uint64 {
    var hash uint64 = 1315423911
    for i := 0; i < len(str); i++ {
        hash ^= (hash << 5) + uint64(str[i]) + (hash >> 2)
    }
    return hash
}
