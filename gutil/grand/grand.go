package grand

import (
    "math/rand"
)

type GRand struct {
    *rand.Rand
}

// NewRand 新建种子随机数 GRand。
// `seed` 随机种子值。
// `avgTimes` 多少次随机(一般取3~5次)？
func NewRand(seed int64) *GRand {
    return &GRand{
        rand.New(rand.NewSource(seed)),
    }
}

// Hit 用于指定一个数num和总数total，往往 num<=total，并随机计算是否满足num/total的概率。
// `randTimes` 取多少次随机值？
// 例如，Hit(1, 100)将会随机计算是否满足百分之一的概率。
func (r *GRand) Hit(num, total int, randTimes ...int) bool {
    rt := 1
    if len(randTimes) > 0 && randTimes[0] > 0 {
        rt = randTimes[0]
    }
    if rt > 1 {
        rv := 0
        for i := 0; i < rt; i++ {
            rv += r.Intn(total)
        }
        rv /= rt
        return rv < num
    }
    return r.Intn(total) < num
}

// HitProb 用于给定一个概率浮点数prob，往往 prob<=1.0，并随机计算是否满足该概率。
// `randTimes` 取多少次随机值？
// 例如，HitProb(0.005)将会随机计算是否满足千分之五的概率。
func (r *GRand) HitProb(prob float32, randTimes ...int) bool {
    rt := 1
    if len(randTimes) > 0 && randTimes[0] > 0 {
        rt = randTimes[0]
    }
    num := int(prob * 1e7)
    if rt > 1 {
        rv := 0
        for i := 0; i < rt; i++ {
            rv += r.Intn(1e7)
        }
        rv /= rt
        return rv < num
    }
    return r.Intn(1e7) < num
}
