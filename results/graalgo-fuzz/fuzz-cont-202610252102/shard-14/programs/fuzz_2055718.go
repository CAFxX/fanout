package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint }
func genFn0(p0 bool, p1 int32) uint {
    p1 -= (p1 ^ (p1 & (p1 - 500)))
    v6 := GenS1{F0: 32768}
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := []int{128, -1, 2147483647, -2147483648}
    var v2 uint64 = 0
    var v3 uint = 261
    var v4 uint8 = 120
    var v5 uint8 = 0
    _ = len(v1)
    var v7 int64 = 9223372036854775807
    var v8 uint64 = (v2 >> (((v3 << ((v3) & 63))) & 63))
    var v9 int = -573
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int16 = 695
        _ = fi10
        _ = v11
    }
    v3 = (v3 & genFn0(true, 42))
    v12 := GenS1{F0: v9}
    if (((v9 * 0) == -2147483648) && true) {
        v13 := GenS2{F0: true, F1: 42, F2: 42}
        v1 = append(v1, (v9 ^ (v9 + (-(v9)))))
        _ = v13
    } else {
        v8 = (^(v2))
        switch "a" {
        case "foo bar":
            var v14 uint8 = 1
            _ = v14
        case "":
            for fi15 := 0; fi15 < 16; fi15++ {
                fmt.Println(v9, v5)
                _ = fi15
            }
        }
    }
    v16 := GenS2{F0: (2147483647 < 1), F1: 342, F2: 1}
    v4 = (v4 << (((v3 % ((v3) | 1))) & 63))
    if (2147483647 <= 934) {
        var v17 uint16 = 312
        v2 -= v8
        var v18 int = v9
        _ = v17
        _ = v18
    } else {
        fmt.Println(v8)
        var v19 bool = true
        v1 = append(v1, v9)
        _ = v19
    }
    fmt.Println(v3, v4, v2)
    if false {
        _ = len(v1)
        if len(v1) > 0 {
            v1[0] = (v9 << ((4294967295) & 63))
        }
    }
    v20 := GenS3{F0: (^(v3))}
    v9 += v9
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
