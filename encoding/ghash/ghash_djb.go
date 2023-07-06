package ghash

// DJB 实现经典的 DJB 哈希算法32位。
func DJB(str []byte) uint32 {
    var hash uint32 = 5381
    for i := 0; i < len(str); i++ {
        hash += (hash << 5) + uint32(str[i])
    }
    return hash
}

// DJB64 实现经典的 DJB 哈希算法64位。
func DJB64(str []byte) uint64 {
    var hash uint64 = 5381
    for i := 0; i < len(str); i++ {
        hash += (hash << 5) + uint64(str[i])
    }
    return hash
}
