package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8; F1 int32; F2 int32 }
type GenS4 struct { F0 int16; F1 int8; F2 int32 }
func genFn0(p0 string, p1 uint64, p2 string) uint32 {
    v6 := []int{-2147483648, 9223372036854775807}
    p0 = (((p2 + p2) + p0) + p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}
func genFn1(p0 uint32, p1 string, p2 uint64) int64 {
    p0 += genFn0(p1, (p2 ^ (p2 << ((628) & 63))), p1)
    var v7 uint = 2
    v8 := []int{}
    var v9 float64 = float64((float32((((1e10 / 0.1) + 0.1) + (-2.25 / 3.14159))) * (1e10 / ((((3.14159 / 1.5) - 0.1) - -2.25) / 3.14159))))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = (((0.0 * (-2.25 / ((-2.25 / float32(1e10)) / float32(-2.25)))) + (1e10 / 0.1)) - 0.1)
    var v3 int = (v1 % (((v1 * v1)) | 1))
    var v4 int32 = 42
    var v5 float64 = 1.5
    if (-2147483648 < 535) {
        fmt.Println(v4, v1, v2)
        var v10 int64 = genFn1(498, "", 959)
        v11 := GenS1{F0: "\n\t"}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 float32 = 1e10
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v5, v2)
    for fi14 := 0; fi14 < 20; fi14++ {
        switch 32768 {
        case 1:
            if (!true) {
                v15 := GenS2{F0: 42, F1: 0}
                v1 += -511
                v5 = v5
                _ = v15
            }
        case 0:
            var v16 uint16 = 255
            _ = v16
        }
        v1 = fi14
        _ = fi14
    }
    var v17 int8 = 38
    var v18 bool = true
    for fi19 := 0; fi19 < 5; fi19++ {
        if v18 {
            var v20 uint = 709
            v1 -= v3
            var v21 int = fi19
            _ = v20
            _ = v21
        }
        var v22 int = 32767
        v23 := GenS3{F0: 222, F1: (v4 & v4), F2: (v4 * v4)}
        _ = fi19
        _ = v22
        _ = v23
    }
    v24 := []int{(-(v1)), (v1 & (v3 << ((2) & 63))), -1}
    v17 -= v17
    if (!((false && true) && (!true))) {
        v25 := GenS1{F0: "\n\t"}
        v3 = v3
        _ = v25
    } else {
        var v26 int16 = 127
        fmt.Println(v17, v17)
        _ = v26
    }
    v27 := GenS3{F0: 131, F1: 42, F2: v4}
    v28 := GenS4{F0: 32767, F1: v17, F2: v4}
    var v29 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
