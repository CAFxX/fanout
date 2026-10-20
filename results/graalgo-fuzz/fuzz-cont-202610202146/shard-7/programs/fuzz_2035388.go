package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint) uint {
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(p1)
        p1 += (p1 / ((p1) | 1))
        fmt.Println(p1, p1, p1)
        _ = fi7
    }
    if (751 == 127) {
        fmt.Println(p0)
        v8 := []int{127, 2147483647, 1, 42}
        if false {
            var v9 uint8 = 2
            _ = v9
        } else {
            v8 = append(v8, 32767)
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = v8
    } else {
        v10 := GenS1{F0: 42}
        var v11 string = "a"
        if true {
            var v12 string = v11
            _ = v12
        }
        _ = v10
        _ = v11
    }
    p1 = (p1 ^ 2)
    fmt.Println(p0, p1)
    switch 1 {
    case 42:
        switch "\n\t" {
        case "a":
            var v13 int32 = 42
            _ = v13
        case "\n\t":
            v14 := GenS1{F0: 2}
            _ = v14
        case "hello":
            p1 = p1
        default:
            if (-9223372036854775808 != -9223372036854775808) {
                var v15 uint32 = 30
                _ = v15
            }
        }
    }
    _ = p0
    _ = p1
    return 70
}

func main() {
    v1 := []uint32{128, 4294967295, 245}
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]bool)
    v4 := []int{-955, 32767}
    var v5 uint32 = 0
    v6 := []uint32{(^(v5)), (^(v5))}
    v3[("a" + "")] = false
    if true {
        v16 := v3["a"]
        var v17 uint = 582
        _ = v16
        _ = v17
    }
    switch 559 {
    case 2147483647:
        v5 = 0
    case 9223372036854775807:
        v18 := GenS2{F0: (-327 != -978), F1: -204}
        _ = v18
    }
    switch "\n\t" {
    case "\n\t":
        v5 -= v5
    case "foo bar":
        fmt.Println(v5, v2)
    case "hello":
        v19 := GenS3{F0: 884}
        _ = v19
    }
    v2 -= (v2 + -1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (428 == 42) {
        v20 := GenS2{F0: false, F1: 9223372036854775807}
        _ = v20
    }
    fmt.Println(v5, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
