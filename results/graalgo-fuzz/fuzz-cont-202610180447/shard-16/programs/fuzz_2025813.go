package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8) uint32 {
    var v6 uint = 215
    v6 -= (^(v6))
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 int8, p1 int32, p2 int8) int64 {
    v7 := []int{204, -2147483648}
    var v8 uint16 = 24
    var v9 uint32 = 832
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 672
}

func main() {
    v1 := []uint32{2, 42, 42, 2147483648}
    var v2 int8 = 1
    var v3 uint16 = 363
    var v4 int8 = (v2 / ((v2) | 1))
    var v5 int = 32767
    v5 = 1
    v5 = (v5 / (((v5 + (v5 % ((v5) | 1)))) | 1))
    var v10 int32 = 55
    switch ("a" + "") {
    case "foo bar":
        if false {
            v1 = append(v1, 879)
        } else {
            fmt.Println(v3, v5, v10)
            if len(v1) > 0 {
                v1[0] = 2
            }
            var v11 uint64 = 788
            _ = v11
        }
    }
    v12 := GenS1{F0: 32767}
    var v13 int8 = v2
    v2 = (v2 << ((846) & 63))
    fmt.Println(v3, v3, v4)
    if false {
        v13 -= v13
    }
    if true {
        var v14 uint32 = 91
        _ = v14
    }
    switch "a" {
    case "\n\t":
        for fi15 := 0; fi15 < 1; fi15++ {
            if false {
                v16 := GenS1{F0: 1}
                _ = v16
            }
            _ = fi15
        }
    case "a":
        v10 = v10
    }
    v17 := GenS1{F0: 923}
    switch 4294967295 {
    case -2147483648:
        var v18 uint = 2
        _ = v18
    case 128:
        var v19 uint16 = (v3 ^ (v3 ^ 42))
        _ = v19
    case 468:
        fmt.Println(v3, v13, v5)
    default:
        v20 := GenS1{F0: 55}
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS1{F0: int16(v10)}
        var v23 int16 = 0
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
