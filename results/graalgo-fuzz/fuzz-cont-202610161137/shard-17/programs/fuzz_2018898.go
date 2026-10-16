package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 uint8) int32 {
    var v6 int = 977
    var v7 int8 = -1
    fmt.Println(p0, v6)
    v8 := []int{v6}
    p0 += p0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 uint) int32 {
    fmt.Println(p0)
    var v9 float64 = float64(3.14159)
    p0 += 42
    _ = p0
    _ = v9
    return -2147483648
}

func main() {
    var v1 int32 = 867
    var v2 int16 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    v5 := []uint32{42}
    fmt.Println(v1, v2)
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            v5 = append(v5, 65535)
            _ = fi10
        }
        fmt.Println(v2, v1)
        v11 := GenS1{F0: (v2 / ((v2) | 1)), F1: 568}
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int32 = v1
        var v14 float64 = -2.25
        v14 += (v14 * 3.14159)
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float32 = (1e10 * 0.1)
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint32 = 2
        _ = fi16
        _ = v17
    }
    _ = len(v3)
    v1 += (v1 << ((221) & 63))
    if (v15 >= float32(3.14159)) {
        v1 += genFn0(127)
        switch -857 {
        case -410:
            var v18 uint64 = 2
            _ = v18
        case 1:
            v1 = 2147483647
        }
        switch 4294967295 {
        case 128:
            v19 := GenS1{F0: v2, F1: 820}
            _ = v19
        case 32767:
            var v20 uint16 = 32767
            _ = v20
        }
    }
    fmt.Println(v1)
    v2 -= v2
    v2 += (v2 ^ 1)
    fmt.Println(v15, v1, v15)
    var v21 uint32 = 2147483647
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
