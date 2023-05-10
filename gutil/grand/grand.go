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

// RangeInt 随机整数方法返回 `min` 到 `max` 之间的随机整数，支持负数，包含边界，即：[min, max]。
func (r *GRand) RangeInt(min, max int) int {
    if min >= max {
        return min
    }
    if min >= 0 {
        return r.Intn(max-min+1) + min
    }
    // 由于 `Intn` 不支持负数，所以我们应该首先将值向右移动，然后调用 `Intn` 产生随机数，并将最终结果移回左侧。
    return r.Intn(max+(0-min)+1) - (0 - min)
}

// Hit 用于指定一个数 `num` 和总数 `total` ，往往 num<=total，并随机计算是否满足 num/total 的概率。
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

// HitProb 用于给定一个概率浮点数 `prob`，往往 prob<=1.0，并随机计算是否满足该概率。
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
