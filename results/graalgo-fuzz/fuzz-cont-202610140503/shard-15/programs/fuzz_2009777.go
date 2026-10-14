package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 int32, p1 uint32, p2 int16) uint8 {
    p1 = 7
    for fi6 := 0; fi6 < 13; fi6++ {
        p0 = p0
        _ = fi6
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 int8 = 100
    var v2 uint = 86
    var v3 int64 = 570
    var v4 bool = ((v2 ^ v2) == 444)
    var v5 int16 = 0
    if (!((0.1 + 0.0) >= 1e10)) {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := []int{}
            v9 := GenS1{F0: 32767, F1: fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{2147483647, 449, -9223372036854775808, 32768}
        if len(v10) > 0 {
            v10[0] = 2147483647
        }
        _ = v10
    }
    fmt.Println(v3, v4)
    v11 := []int{-9223372036854775808, 619, 128}
    v5 += (v5 ^ (v5 / ((v5) | 1)))
    fmt.Println(v1, v1)
    var v12 uint32 = 1
    var v13 uint16 = 238
    var v14 string = ""
    var v15 uint = v2
    v16 := GenS2{F0: v12, F1: (v1 << (((^(v2))) & 63))}
    v2 += (v15 * 0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
