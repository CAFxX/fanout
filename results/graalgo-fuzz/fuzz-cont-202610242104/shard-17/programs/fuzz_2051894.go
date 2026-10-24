package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
func genFn0(p0 int, p1 uint8) string {
    var v4 uint32 = 42
    p1 = (p1 >> ((65535) & 63))
    v5 := []int{(p0 ^ (p0 ^ (p0 | -2147483648)))}
    v6 := GenS1{F0: 211, F1: -2147483648}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 int64) uint64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 += (p0 & (p0 + p0))
        v8 := []int{(fi7 & (^(fi7))), fi7, 9223372036854775807}
        var v9 uint8 = 169
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{82, -190}
    _ = p0
    _ = v10
    return 553
}

func main() {
    var v1 uint32 = 0
    var v2 uint8 = 255
    var v3 int16 = 0
    fmt.Println(v1, v1, v1)
    v2 = v2
    var v11 uint = 256
    var v12 int = -1
    var v13 uint16 = 892
    var v14 uint = v11
    v13 += (v13 << ((uint(v1)) & 63))
    var v15 uint = v11
    var v16 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
