package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint }
func genFn0() uint16 {
    var v5 uint = 52
    v6 := []int{-574, 962, 32767}
    var v7 int8 = 0
    for fi8 := 0; fi8 < 17; fi8++ {
        v7 -= v7
        v5 = (^(v5))
        for fi9 := 0; fi9 < 2; fi9++ {
            switch 38 {
            case -663:
                fmt.Println(fi8, v7)
            }
            v10 := GenS1{F0: (float64((float32(1e10) - float32((3.14159 - 3.14159)))) != (float64((0.0 * float32((float64(-2.25) - 3.14159)))) + (((-2.25 / 1.5) - 3.14159) + 3.14159))), F1: 0, F2: 42}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if true {
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int8 = v7
            var v13 uint64 = 760
            _ = fi11
            _ = v12
            _ = v13
        }
        if false {
            var v14 int16 = 800
            _ = v14
        }
    } else {
        if false {
            v7 = v7
        } else {
            var v15 float32 = 1.5
            fmt.Println(v7, v15, v5)
            v6 = append(v6, 128)
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS2{F0: false, F1: 0, F2: (v5 + (v5 % ((v5) | 1)))}
            fmt.Println(v5, v5)
            _ = len(v6)
            _ = fi16
            _ = v17
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -9223372036854775808
    v3 := []int{v2, -9223372036854775808}
    var v4 bool = (2147483647 < 680)
    var v18 uint = 0
    var v19 uint32 = 1
    v20 := GenS2{F0: v4, F1: 74, F2: (v18 | (v18 - (v18 + (^(v18)))))}
    var v21 int32 = 42
    for fi22 := 0; fi22 < 8; fi22++ {
        v19 -= v19
        v23 := GenS3{F0: (v18 ^ 4294967295), F1: 65535}
        v24 := GenS4{F0: 2, F1: 228}
        _ = fi22
        _ = v23
        _ = v24
    }
    v18 += (v18 / (((v18 >> ((958) & 63))) | 1))
    var v25 uint32 = (^(v19))
    var v26 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
