package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0() uint8 {
    var v4 uint16 = 65535
    var v5 uint8 = 138
    _ = v4
    _ = v5
    return v5
}
func genFn1(p0 string, p1 uint32, p2 int16) uint8 {
    var v6 int64 = -9223372036854775808
    var v7 uint = 65535
    switch -9223372036854775808 {
    case 9223372036854775807:
        p1 = p1
    case 32768:
        v8 := []int{0, 127}
        _ = v8
    case 0:
        if true {
            var v9 int32 = 1
            var v10 float64 = -2.25
            _ = v10
            _ = v9
        }
    default:
        var v11 int16 = p2
        _ = v11
    }
    var v12 bool = (!false)
    v13 := GenS1{F0: uint8(v7), F1: 93}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 42
    var v3 int32 = 2147483647
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v2)
        v15 := GenS1{F0: 128, F1: v2}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: v2}
    v1[65535] = ""
    v3 -= 1
    if (255 > genFn0()) {
        v17 := []int{127, 128, 128, 956}
        _ = v17
    }
    v18 := GenS1{F0: 42, F1: v2}
    v3 = v3
    v2 = 65535
    var v19 uint8 = 94
    var v20 int32 = int32(v2)
    var v21 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
