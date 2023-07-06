package ghash

// DJB implements the classic DJB hash algorithm for 32 bits.
func DJB(str []byte) uint32 {
    var hash uint32 = 5381
    for i := 0; i < len(str); i++ {
        hash += (hash << 5) + uint32(str[i])
    }
    return hash
}

// DJB64 implements the classic DJB hash algorithm for 64 bits.
func DJB64(str []byte) uint64 {
    var hash uint64 = 5381
    for i := 0; i < len(str); i++ {
        hash += (hash << 5) + uint64(str[i])
    }
    return hash
}
