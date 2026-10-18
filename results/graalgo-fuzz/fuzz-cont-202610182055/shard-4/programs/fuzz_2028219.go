package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint32, p1 int32) int {
    for fi7 := 0; fi7 < 16; fi7++ {
        p0 += 0
        _ = fi7
    }
    p0 += p0
    var v8 uint16 = 0
    var v9 int = -308
    v8 -= uint16(p0)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 uint64) int64 {
    fmt.Println(p0)
    v10 := []int{0, -103}
    if false {
        fmt.Println(p0)
    } else {
        p0 = (p0 / (((^(p0))) | 1))
        fmt.Println(p0)
        var v11 int32 = 0
        _ = v11
    }
    v12 := GenS1{F0: 426, F1: 9223372036854775807, F2: -9223372036854775808}
    _ = p0
    _ = v10
    _ = v12
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 18446744073709551615
    var v3 string = "foo bar"
    var v4 uint64 = (v2 % (((v2 % ((2) | 1))) | 1))
    var v5 uint = uint(v2)
    v6 := []uint32{1, 281}
    v6 = append(v6, 1)
    v1 -= 255
    v4 = v2
    var v13 int16 = 1000
    if (65535 <= 65535) {
        fmt.Println(v5, v5)
        if false {
            v3 = v3
        }
    } else {
        var v14 int8 = -72
        _ = v14
    }
    v6 = append(v6, 1)
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS1{F0: 0, F1: 9223372036854775807, F2: -1}
        _ = len(v6)
        if (1.5 != -2.25) {
            v2 = v2
            v1 -= v1
        } else {
            v3 = "a"
            v17 := GenS1{F0: 0, F1: -9223372036854775808, F2: 810}
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    v3 = v3
    v18 := GenS1{F0: 42, F1: 42, F2: 424}
    v19 := GenS2{F0: 0, F1: -121}
    switch "foo bar" {
    case "foo bar":
        v13 = (v13 - v13)
    case "":
        if (!(!(v13 >= v13))) {
            var v20 uint16 = 2
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            fmt.Println(v5)
            _ = v20
        } else {
            v13 -= (v13 << ((431) & 63))
        }
    default:
        v4 = (v4 & 2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
