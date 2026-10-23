package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() int8 {
    v7 := []int{2147483647, 127, -731, -2147483648}
    _ = len(v7)
    if (float64(-2.25) == ((0.0 / 1e10) * ((1e10 - 3.14159) + -2.25))) {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 255}
            var v10 int8 = 0
            fmt.Println(fi8, v10)
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := GenS1{F0: 2}
    v12 := GenS1{F0: 542}
    _ = v11
    _ = v12
    _ = v7
    return 116
}
func genFn1(p0 uint, p1 uint16, p2 uint8) string {
    v13 := []int{4294967295}
    v14 := GenS1{F0: (p1 * (p1 << (((^(p0))) & 63)))}
    var v15 float32 = (((3.14159 * (-2.25 + 0.0)) / 1.5) * (((float32(-2.25) - float32((float64(1e10) + 1.5))) + 1e10) / 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return "a"
}

func main() {
    var v1 uint16 = 65535
    var v2 int = 248
    var v3 uint32 = 808
    v4 := make(map[int]int)
    v5 := []uint32{v3}
    var v6 uint = 255
    v6 += (v6 - 312)
    fmt.Println(v1, v3)
    v16 := GenS1{F0: v1}
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            delete(v4, -9223372036854775808)
            _ = fi18
        }
        fmt.Println(v2, v2)
        _ = fi17
    }
    fmt.Println(v6, v3)
    v19 := GenS2{F0: -2147483648, F1: (v1 + (v1 ^ v1))}
    v3 = v3
    fmt.Println(v2)
    v6 += (v6 | v6)
    if (935 > -9223372036854775808) {
        if (!true) {
            v6 -= (^(v6))
        } else {
            var v20 uint32 = (^(v3))
            v1 = v1
            _ = v20
        }
        var v21 int16 = 32767
        _ = v21
    } else {
        var v22 uint = v6
        var v23 uint8 = 242
        v23 += v23
        _ = v22
        _ = v23
    }
    var v24 uint8 = 9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
