package gregex

import (
    "regexp"
    "sync"

    "github.com/camry/g/gerrors/gerror"
)

var (
    regexMu = sync.RWMutex{}

    // Cache for regex object.
    // Note that:
    // 1. It uses sync.RWMutex ensuring the concurrent safety.
    // 2. There's no expiring logic for this map.
    regexMap = make(map[string]*regexp.Regexp)
)

// getRegexp returns *regexp.Regexp object with given `pattern`.
// It uses cache to enhance the performance for compiling regular expression pattern,
// which means, it will return the same *regexp.Regexp object with the same regular
// expression pattern.
//
// It is concurrent-safe for multiple goroutines.
func getRegexp(pattern string) (regex *regexp.Regexp, err error) {
    // Retrieve the regular expression object using reading lock.
    regexMu.RLock()
    regex = regexMap[pattern]
    regexMu.RUnlock()
    if regex != nil {
        return
    }
    // If it does not exist in the cache,
    // it compiles the pattern and creates one.
    if regex, err = regexp.Compile(pattern); err != nil {
        err = gerror.Wrapf(err, `regexp.Compile failed for pattern "%s"`, pattern)
        return
    }
    // Cache the result object using writing lock.
    regexMu.Lock()
    regexMap[pattern] = regex
    regexMu.Unlock()
    return
}
