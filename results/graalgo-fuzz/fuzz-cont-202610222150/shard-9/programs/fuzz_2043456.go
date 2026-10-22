package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 int16, p1 uint64, p2 uint32) uint {
    if (true || false) {
        var v5 int64 = -738
        v6 := GenS1{F0: 4294967295, F1: 2147483647}
        p0 += 32767
        _ = v5
        _ = v6
    }
    var v7 float64 = 1e10
    p2 -= 2147483648
    if true {
        for fi8 := 0; fi8 < 11; fi8++ {
            p0 = (-(p0))
            v9 := []int{(fi8 % ((fi8) | 1)), (fi8 + 2147483647), (fi8 & fi8), fi8}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint64 = p1
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 string, p1 uint64, p2 uint) bool {
    v12 := []int{}
    var v13 uint16 = 255
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return (42 != -1)
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 42
    v3 := make(map[uint32]string)
    var v4 float32 = (1.5 + ((0.1 + ((float32(0.0) - (-2.25 * 0.0)) - 3.14159)) * (-2.25 + 0.0)))
    var v14 int8 = -128
    var v15 uint64 = 2
    v16 := GenS1{F0: 2, F1: v1}
    v17 := GenS1{F0: 938, F1: v1}
    v1 += 42
    v4 += (v4 * float32(3.14159))
    v18 := GenS1{F0: 2, F1: 1}
    v15 -= v15
    var v19 float32 = (float32((1e10 * (1.5 * -2.25))) + (v4 / (float32(1.5) * float32(1e10))))
    if false {
        if (!((^(v15)) <= (v15 * v15))) {
            switch "" {
            case "a":
                v1 += (v1 >> ((1) & 63))
            default:
                v20 := []int{-2147483648}
                _ = v20
            }
        } else {
            v15 = (v15 | (v15 >> ((2) & 63)))
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 uint64 = v15
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
