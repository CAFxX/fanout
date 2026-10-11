package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int32 }
type GenS2 struct { F0 int; F1 int8; F2 int8 }
func genFn0(p0 int32, p1 uint16, p2 string) uint16 {
    fmt.Println(p2, p0)
    p2 = p2
    p2 = (((p2 + "a") + ("" + p2)) + p2)
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint64) uint64 {
    if false {
        v4 := GenS1{F0: 100, F1: 0, F2: -383}
        _ = v4
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(fi5, p0)
        _ = fi5
    }
    switch 2147483647 {
    case 32767:
        var v6 float64 = (float64(((((3.14159 * 0.1) + 0.1) + (0.0 * (1e10 / 0.1))) * (3.14159 * -2.25))) - (((float64(0.1) * 0.0) * float64(float32(3.14159))) + 1.5))
        _ = v6
    case 895:
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint8 = 128
            _ = fi7
            _ = v8
        }
    case 1:
        fmt.Println(p0)
    }
    if (false && false) {
        var v9 int = 906
        _ = v9
    } else {
        p0 = p0
        p0 = p0
    }
    _ = p0
    return 83
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 int32 = -2147483648
    v10 := []int{128, 552}
    v11 := GenS1{F0: 65535, F1: -1, F2: v3}
    var v12 int32 = v3
    var v13 int32 = (v3 % (((v12 & v12)) | 1))
    v14 := GenS2{F0: 841, F1: 127, F2: -127}
    var v15 float64 = 1e10
    v13 = 1
    var v16 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
