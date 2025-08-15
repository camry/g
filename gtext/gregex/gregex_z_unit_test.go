// go test *.go -bench=".*"

package gregex_test

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gtext/gregex"
)

var (
    PatternErr = `([\d+`
)

func Test_Quote(t *testing.T) {
    s1 := `[foo]` // `\[foo\]`
    assert.Equal(t, gregex.Quote(s1), `\[foo\]`)
}

func Test_Validate(t *testing.T) {
    var s1 = `(.+):(\d+)`
    assert.Equal(t, gregex.Validate(s1), nil)
    s1 = `((.+):(\d+)`
    assert.Equal(t, gregex.Validate(s1) == nil, false)
}

func Test_IsMatch(t *testing.T) {
    var pattern = `(.+):(\d+)`
    s1 := []byte(`sfs:2323`)
    assert.Equal(t, gregex.IsMatch(pattern, s1), true)
    s1 = []byte(`sfs2323`)
    assert.Equal(t, gregex.IsMatch(pattern, s1), false)
    s1 = []byte(`sfs:`)
    assert.Equal(t, gregex.IsMatch(pattern, s1), false)
    // error pattern
    assert.Equal(t, gregex.IsMatch(PatternErr, s1), false)
}

func Test_IsMatchString(t *testing.T) {
    var pattern = `(.+):(\d+)`
    s1 := `sfs:2323`
    assert.Equal(t, gregex.IsMatchString(pattern, s1), true)
    s1 = `sfs2323`
    assert.Equal(t, gregex.IsMatchString(pattern, s1), false)
    s1 = `sfs:`
    assert.Equal(t, gregex.IsMatchString(pattern, s1), false)
    // error pattern
    assert.Equal(t, gregex.IsMatchString(PatternErr, s1), false)
}

func Test_Match(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    s := "acbb" + wantSubs + "dd"
    subs, err := gregex.Match(re, []byte(s))
    assert.Nil(t, err)
    if string(subs[0]) != wantSubs {
        t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0], wantSubs)
    }
    if string(subs[1]) != "aab" {
        t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1], "aab")
    }
    // error pattern
    _, err = gregex.Match(PatternErr, []byte(s))
    assert.NotNil(t, err)
}

func Test_MatchString(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    s := "acbb" + wantSubs + "dd"
    subs, err := gregex.MatchString(re, s)
    assert.Nil(t, err)
    if string(subs[0]) != wantSubs {
        t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0], wantSubs)
    }
    if string(subs[1]) != "aab" {
        t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1], "aab")
    }
    // error pattern
    _, err = gregex.MatchString(PatternErr, s)
    assert.NotNil(t, err)
}

func Test_MatchAll(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    s := "acbb" + wantSubs + "dd"
    s = s + `其他的` + s
    subs, err := gregex.MatchAll(re, []byte(s))
    assert.Nil(t, err)
    if string(subs[0][0]) != wantSubs {
        t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0][0], wantSubs)
    }
    if string(subs[0][1]) != "aab" {
        t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[0][1], "aab")
    }

    if string(subs[1][0]) != wantSubs {
        t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[1][0], wantSubs)
    }
    if string(subs[1][1]) != "aab" {
        t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1][1], "aab")
    }
    // error pattern
    _, err = gregex.MatchAll(PatternErr, []byte(s))
    assert.NotNil(t, err)
}

func Test_MatchAllString(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    s := "acbb" + wantSubs + "dd"
    subs, err := gregex.MatchAllString(re, s+`其他的`+s)
    assert.Nil(t, err)
    if string(subs[0][0]) != wantSubs {
        t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0][0], wantSubs)
    }
    if string(subs[0][1]) != "aab" {
        t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[0][1], "aab")
    }

    if string(subs[1][0]) != wantSubs {
        t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[1][0], wantSubs)
    }
    if string(subs[1][1]) != "aab" {
        t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1][1], "aab")
    }
    // error pattern
    _, err = gregex.MatchAllString(PatternErr, s)
    assert.NotNil(t, err)
}

func Test_Replace(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    replace := "12345"
    s := "acbb" + wantSubs + "dd"
    wanted := "acbb" + replace + "dd"
    replacedStr, err := gregex.Replace(re, []byte(replace), []byte(s))
    assert.Nil(t, err)
    if string(replacedStr) != wanted {
        t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
    }
    // error pattern
    _, err = gregex.Replace(PatternErr, []byte(replace), []byte(s))
    assert.NotNil(t, err)
}

func Test_ReplaceString(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    replace := "12345"
    s := "acbb" + wantSubs + "dd"
    wanted := "acbb" + replace + "dd"
    replacedStr, err := gregex.ReplaceString(re, replace, s)
    assert.Nil(t, err)
    if replacedStr != wanted {
        t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
    }
    // error pattern
    _, err = gregex.ReplaceString(PatternErr, replace, s)
    assert.NotNil(t, err)
}

func Test_ReplaceFun(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    // replace :="12345"
    s := "acbb" + wantSubs + "dd"
    wanted := "acbb[x" + wantSubs + "y]dd"
    wanted = "acbb" + "3个a" + "dd"
    replacedStr, err := gregex.ReplaceFunc(re, []byte(s), func(s []byte) []byte {
        if strings.Contains(string(s), "aaa") {
            return []byte("3个a")
        }
        return []byte("[x" + string(s) + "y]")
    })
    assert.Nil(t, err)
    if string(replacedStr) != wanted {
        t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
    }
    // error pattern
    _, err = gregex.ReplaceFunc(PatternErr, []byte(s), func(s []byte) []byte {
        return []byte("")
    })
    assert.NotNil(t, err)
}

func Test_ReplaceFuncMatch(t *testing.T) {
    s := []byte("1234567890")
    p := `(\d{3})(\d{3})(.+)`
    s0, e0 := gregex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
        return match[0]
    })
    assert.Nil(t, e0)
    assert.Equal(t, s0, s)
    s1, e1 := gregex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
        return match[1]
    })
    assert.Nil(t, e1)
    assert.Equal(t, s1, []byte("123"))
    s2, e2 := gregex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
        return match[2]
    })
    assert.Nil(t, e2)
    assert.Equal(t, s2, []byte("456"))
    s3, e3 := gregex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
        return match[3]
    })
    assert.Nil(t, e3)
    assert.Equal(t, s3, []byte("7890"))
    // error pattern
    _, err := gregex.ReplaceFuncMatch(PatternErr, s, func(match [][]byte) []byte {
        return match[3]
    })
    assert.NotNil(t, err)
}

func Test_ReplaceStringFunc(t *testing.T) {
    re := "a(a+b+)b"
    wantSubs := "aaabb"
    // replace :="12345"
    s := "acbb" + wantSubs + "dd"
    wanted := "acbb[x" + wantSubs + "y]dd"
    wanted = "acbb" + "3个a" + "dd"
    replacedStr, err := gregex.ReplaceStringFunc(re, s, func(s string) string {
        if strings.Contains(s, "aaa") {
            return "3个a"
        }
        return "[x" + s + "y]"
    })
    assert.Nil(t, err)
    if replacedStr != wanted {
        t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
    }
    // error pattern
    _, err = gregex.ReplaceStringFunc(PatternErr, s, func(s string) string {
        return ""
    })
    assert.NotNil(t, err)
}

func Test_ReplaceStringFuncMatch(t *testing.T) {
    s := "1234567890"
    p := `(\d{3})(\d{3})(.+)`
    s0, e0 := gregex.ReplaceStringFuncMatch(p, s, func(match []string) string {
        return match[0]
    })
    assert.Nil(t, e0)
    assert.Equal(t, s0, s)
    s1, e1 := gregex.ReplaceStringFuncMatch(p, s, func(match []string) string {
        return match[1]
    })
    assert.Nil(t, e1)
    assert.Equal(t, s1, "123")
    s2, e2 := gregex.ReplaceStringFuncMatch(p, s, func(match []string) string {
        return match[2]
    })
    assert.Nil(t, e2)
    assert.Equal(t, s2, "456")
    s3, e3 := gregex.ReplaceStringFuncMatch(p, s, func(match []string) string {
        return match[3]
    })
    assert.Nil(t, e3)
    assert.Equal(t, s3, "7890")
    // error pattern
    _, err := gregex.ReplaceStringFuncMatch(PatternErr, s, func(match []string) string {
        return ""
    })
    assert.NotNil(t, err)
}

func Test_Split(t *testing.T) {
    re := "a(a+b+)b"
    matched := "aaabb"
    item0 := "acbb"
    item1 := "dd"
    s := item0 + matched + item1
    assert.Equal(t, gregex.IsMatchString(re, matched), true)
    items := gregex.Split(re, s) // split string with matched
    if items[0] != item0 {
        t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
    }
    if items[1] != item1 {
        t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
    }

    re1 := "a(a+b+)b"
    notmatched := "aaxbb"
    item00 := "acbb"
    item11 := "dd"
    s1 := item00 + notmatched + item11
    assert.Equal(t, gregex.IsMatchString(re1, notmatched), false)
    items1 := gregex.Split(re1, s1) // split string with notmatched then nosplitting
    if items1[0] != s1 {
        t.Fatalf("regex:%s,Split(%q) want %q", re1, s1, item00)
    }
    // error pattern
    items1 = gregex.Split(PatternErr, s1)
    assert.Nil(t, items1)
}
