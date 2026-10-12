package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
func genFn0(p0 string, p1 int64, p2 int) uint32 {
    var v6 uint64 = 2
    var v7 uint64 = 211
    v8 := []int{(p2 << ((65535) & 63)), int(v6), p2, p2}
    for fi9 := 0; fi9 < 18; fi9++ {
        v6 = v7
        for fi10 := 0; fi10 < 4; fi10++ {
            _ = len(v8)
            _ = fi10
        }
        _ = fi9
    }
    if (753 <= 2) {
        var v11 bool = ((p1 << ((18446744073709551615) & 63)) <= p1)
        fmt.Println(p0, v6, v7)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 int8 = -57
    v2 := []uint32{42}
    v3 := make(map[int]bool)
    v4 := []int{-2147483648, 32768}
    var v5 uint32 = 4294967295
    v4 = append(v4, 0)
    v12 := GenS1{F0: 984}
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    switch -840 {
    case -2147483648:
        var v13 float64 = (3.14159 + (1.5 * 1.5))
        _ = v13
    case 990:
        v14 := GenS1{F0: 4294967295}
        _ = v14
    }
    var v15 int = -1
    fmt.Println(v5, v1, v1)
    v5 = v5
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint8 = 42
        var v18 int64 = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: -931, F1: 0}
    v5 = (v5 - v5)
    v20 := GenS3{F0: true, F1: (false || (0 > 1)), F2: 1}
    v1 -= v1
    var v21 uint32 = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 int = (v15 - (v15 % ((v15) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
