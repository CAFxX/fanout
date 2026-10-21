package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int8 }
type GenS2 struct { F0 int; F1 int64; F2 int32 }
func genFn0(p0 uint16, p1 int32) string {
    fmt.Println(p1)
    p1 = 42
    _ = p0
    _ = p1
    return ""
}
func genFn1() int16 {
    var v7 int32 = 218
    if false {
        v8 := []int{-9223372036854775808, 32767, 128, 4294967295}
        fmt.Println(v7)
        switch 32767 {
        case -984:
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = v8
    } else {
        v7 = (v7 / (((v7 | (v7 - v7))) | 1))
        switch 765 {
        case 32767:
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS1{F0: -1, F1: -32768, F2: 127}
                if (fi9 != (fi9 * fi9)) {
                    v7 = (^(v7))
                } else {
                    fmt.Println(fi9, fi9)
                }
                var v11 uint = 65535
                _ = fi9
                _ = v10
                _ = v11
            }
        default:
            var v12 int = -889
            _ = v12
        }
    }
    v7 = v7
    v13 := GenS2{F0: -9223372036854775808, F1: 432, F2: (v7 / ((-2147483648) | 1))}
    var v14 uint16 = 245
    _ = v13
    _ = v14
    _ = v7
    return 445
}

func main() {
    v1 := []string{"", ((("hello" + "hello") + ("\n\t" + ("foo bar" + (("a" + "a") + "foo bar")))) + ("hello" + "hello"))}
    var v2 int = 32767
    var v3 int16 = -1
    v4 := make(map[string]uint32)
    var v5 uint64 = 0
    var v6 string = "foo bar"
    if (false && true) {
        v3 += (v3 ^ genFn1())
    } else {
        if len(v1) > 0 {
            v1[0] = v6
        }
        if (!((v5 & 1) < 2)) {
            fmt.Println(v3)
            v2 = 2147483647
        } else {
            _ = len(v1)
            v15 := GenS1{F0: 42, F1: 1000, F2: 100}
            var v16 int32 = 1
            _ = v15
            _ = v16
        }
        var v17 float64 = -2.25
        _ = v17
    }
    var v18 bool = (9223372036854775807 < 1)
    var v19 int16 = v3
    v20 := GenS2{F0: v2, F1: -99, F2: 38}
    v5 += (v5 | 776)
    for fi21 := 0; fi21 < 6; fi21++ {
        v19 = -462
        fmt.Println(v6, v6)
        _ = fi21
    }
    fmt.Println(v3)
    var v22 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
