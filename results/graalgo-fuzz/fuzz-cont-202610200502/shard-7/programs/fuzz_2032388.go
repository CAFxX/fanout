package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 bool; F1 int8 }
func genFn0(p0 uint64) uint16 {
    var v5 float32 = ((((float32((3.14159 - 3.14159)) / (((0.0 + float32(float64((3.14159 * -2.25)))) - float32(float64(0.0))) - float32(1e10))) - -2.25) + ((1e10 * float32(1.5)) + (1.5 - -2.25))) / 1.5)
    v6 := []int{-660, -335, 764, 42}
    _ = p0
    _ = v5
    _ = v6
    return 255
}
func genFn1() uint64 {
    var v7 uint64 = 921
    var v8 uint64 = v7
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 158
    var v3 int8 = 1
    v4 := []string{("a" + ("foo bar" + "")), "", "a", "hello"}
    delete(v1, 0)
    var v9 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint64 = 493
            v13 := GenS1{F0: 937, F1: 276}
            fmt.Println(v2, v3, fi10)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 float64 = 3.14159
    delete(v1, 677)
    var v15 uint8 = 6
    fmt.Println(v9)
    v2 += v2
    if false {
        var v16 float32 = 0.1
        var v17 uint32 = 2
        v2 += 37
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: 0, F1: 4294967295}
        _ = v18
    }
    var v19 uint64 = 791
    fmt.Println(v19)
    v3 -= (v3 - -128)
    v20 := GenS1{F0: v19, F1: (v19 & v19)}
    for fi21 := 0; fi21 < 14; fi21++ {
        switch (fi21 + 42) {
        case 894:
            v22 := GenS1{F0: v19, F1: 4294967295}
            _ = v22
        case 0:
            v23 := GenS2{F0: false, F1: v3}
            _ = v23
        default:
            for fi24 := 0; fi24 < 15; fi24++ {
                v1[1] = 4
                _ = fi24
            }
        }
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
