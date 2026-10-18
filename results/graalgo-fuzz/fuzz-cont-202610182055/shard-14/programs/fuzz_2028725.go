package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 int32) uint16 {
    var v5 string = ""
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 + (fi6 | (fi6 ^ (fi6 ^ fi6)))), (fi6 & fi6), fi6}
        v5 = v5
        var v8 uint32 = 2147483648
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: -657, F2: (p0 << ((220) & 63))}
    _ = p0
    _ = v5
    _ = v9
    return 65535
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    v2 = (v2 / ((v2) | 1))
    v2 = v2
    if true {
        v2 = (v2 / ((v2) | 1))
    }
    v2 = 1
    for fi10 := 0; fi10 < 15; fi10++ {
        v4["a"] = int(v1)
        fmt.Println(v2, fi10, fi10)
        v11 := GenS1{F0: 1000, F1: (v2 | (v2 & (v2 / ((v2) | 1)))), F2: -290}
        _ = fi10
        _ = v11
    }
    fmt.Println(v2)
    var v12 float64 = (0.1 + 1e10)
    v13 := GenS2{F0: -128, F1: -913}
    var v14 int32 = -2147483648
    fmt.Println(v14)
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
