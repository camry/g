package ghash

// PJW implements the classic PJW hash algorithm for 32 bits.
func PJW(str []byte) uint32 {
    var (
        BitsInUnsignedInt uint32 = 32 // 4 * 8
        ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
        OneEighth                = BitsInUnsignedInt / 8
        HighBits          uint32 = (0xFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
        hash              uint32
        test              uint32
    )
    for i := 0; i < len(str); i++ {
        hash = (hash << OneEighth) + uint32(str[i])
        if test = hash & HighBits; test != 0 {
            hash = (hash ^ (test >> ThreeQuarters)) & (^HighBits + 1)
        }
    }
    return hash
}

// PJW64 implements the classic PJW hash algorithm for 64 bits.
func PJW64(str []byte) uint64 {
    var (
        BitsInUnsignedInt uint64 = 32 // 4 * 8
        ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
        OneEighth                = BitsInUnsignedInt / 8
        HighBits          uint64 = (0xFFFFFFFFFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
        hash              uint64
        test              uint64
    )
    for i := 0; i < len(str); i++ {
        hash = (hash << OneEighth) + uint64(str[i])
        if test = hash & HighBits; test != 0 {
            hash = (hash ^ (test >> ThreeQuarters)) & (^HighBits + 1)
        }
    }
    return hash
}
