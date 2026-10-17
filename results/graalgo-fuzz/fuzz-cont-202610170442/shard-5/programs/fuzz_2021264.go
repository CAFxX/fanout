package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
func genFn0(p0 uint32, p1 uint8, p2 int16) int32 {
    var v4 int64 = -70
    v5 := GenS1{F0: 4294967295, F1: -128, F2: 2147483647}
    for fi6 := 0; fi6 < 16; fi6++ {
        p2 = (p2 / ((p2) | 1))
        v7 := []int{fi6, (fi6 << ((667) & 63))}
        _ = fi6
        _ = v7
    }
    fmt.Println(v4)
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 594
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 2
    var v3 uint32 = 1
    for fi8 := 0; fi8 < 20; fi8++ {
        v2 = v2
        fmt.Println(v2, v2)
        var v9 int = 2147483647
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 2
    var v11 uint32 = 940
    var v12 string = ""
    v13 := GenS1{F0: (^(v10)), F1: 41, F2: 2147483647}
    var v14 int8 = -1
    _ = len(v1)
    fmt.Println(v14, v2)
    var v15 int64 = -9223372036854775808
    v12 = ("foo bar" + (v12 + v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
