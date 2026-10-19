package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint32, p2 int64) int8 {
    var v6 uint = 65535
    v7 := GenS1{F0: -32768, F1: (p0 >> ((v6) & 63))}
    switch 9223372036854775807 {
    case 32767:
        p0 = (p0 - 1)
    case 1:
        for fi8 := 0; fi8 < 8; fi8++ {
            v6 -= (v6 & 448)
            _ = fi8
        }
    case 596:
        v9 := []int{1, -362}
        _ = v9
    default:
        fmt.Println(p0, v6, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint16, p1 string, p2 uint) int {
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS2{F0: 877}
        p0 += p0
        _ = fi10
        _ = v11
    }
    p2 -= (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := []string{"\n\t", "a"}
    v2 := make(map[string]bool)
    var v3 int64 = -9223372036854775808
    var v4 bool = true
    var v5 string = ("foo bar" + ("foo bar" + "a"))
    var v12 float64 = 3.14159
    v4 = (0 == 65535)
    if v4 {
        v3 += 9223372036854775807
    }
    v4 = true
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 string = "a"
        _ = fi13
        _ = v14
    }
    if (true && (v4 && (!(!(2 < 201))))) {
        for fi15 := 0; fi15 < 16; fi15++ {
            v4 = v4
            var v16 int64 = (v3 * v3)
            if v4 {
                var v17 int32 = -1
                delete(v2, ("" + v5))
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        fmt.Println(v5)
    }
    fmt.Println(v4)
    v12 = v12
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: 2}
        _ = fi18
        _ = v19
    }
    if false {
        v20 := v2[((v5 + v5) + ((v5 + "") + v5))]
        switch 42 {
        case 0:
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 int64 = -1
                _ = fi21
                _ = v22
            }
        case 127:
            v3 = -9223372036854775808
        case -30:
            var v23 uint8 = 42
            _ = v23
        default:
            for fi24 := 0; fi24 < 5; fi24++ {
                v5 = "foo bar"
                _ = fi24
            }
        }
        _ = v20
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            v2[v5] = v4
            _ = fi25
        }
    }
    if len(v1) > 0 {
        v1[0] = v5
    }
    if false {
        v26 := GenS2{F0: 317}
        if (127 == genFn1(10, (v5 + ((v5 + v5) + v5)), 0)) {
            v4 = v4
            v27 := GenS2{F0: 2}
            if v4 {
                var v28 string = (v5 + "\n\t")
                _ = v28
            } else {
                if (!("foo bar" > v5)) {
                    fmt.Println(v12)
                    v5 = "foo bar"
                    if (true && (1 == 18446744073709551615)) {
                        v4 = true
                    } else {
                        v3 -= v3
                    }
                }
                v12 -= 3.14159
                var v29 string = (v5 + (((("foo bar" + v5) + ((v5 + (v5 + v5)) + v5)) + (((v5 + (v5 + "a")) + ("a" + ((v5 + "foo bar") + (v5 + ("a" + v5))))) + "\n\t")) + v5))
                _ = v29
            }
            _ = v27
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v26
    } else {
        var v30 string = "hello"
        var v31 int32 = 1
        if (v5 > "") {
            for fi32 := 0; fi32 < 9; fi32++ {
                v5 = (v5 + "\n\t")
                _ = fi32
            }
            v30 = (((v30 + v5) + v30) + (v30 + "hello"))
        }
        _ = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
