package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
type GenS5 struct { F0 uint16; F1 bool }
func genFn0(p0 uint, p1 string) uint {
    v6 := []int{42, 90}
    if true {
        v7 := GenS1{F0: -830, F1: 2147483647}
        for fi8 := 0; fi8 < 17; fi8++ {
            if (-128 < 1) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    p1 = p1
                    _ = fi9
                }
            }
            fmt.Println(p0, p1)
            v10 := GenS2{F0: 32767, F1: 100}
            _ = fi8
            _ = v10
        }
        _ = v7
    } else {
        var v11 string = ((p1 + ((p1 + p1) + (p1 + ""))) + p1)
        if (2 < 42) {
            p0 = p0
            p0 = 42
            p0 -= p0
        } else {
            p0 = p0
            var v12 float32 = (float32((0.0 - 0.0)) - ((((0.0 + ((0.0 / ((-2.25 / (((1e10 / 1.0) * ((((1.5 / 1.0) / 1.5) + 0.0) * (0.1 + -2.25))) + (float32(((((0.1 + 1.5) / (-2.25 * 3.14159)) + 3.14159) - ((-2.25 + (3.14159 - 3.14159)) + 1e10))) - 0.0))) / (0.1 + (float32(float64((3.14159 * 0.1))) * (1.5 * 1.5))))) * 0.0)) * float32(0.1)) + (1.5 - -2.25)) / 1.0))
            _ = v12
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := []int{949, -2147483648, 2147483647}
    v2 := make(map[string]bool)
    var v3 uint64 = 4294967295
    v4 := []string{}
    v5 := make(map[uint32]uint32)
    var v13 float64 = -2.25
    for fi14 := 0; fi14 < 4; fi14++ {
        v3 = 782
        _ = fi14
    }
    if (!((255 <= 255) || true)) {
        v4 = append(v4, ("" + ""))
    } else {
        var v15 float32 = -2.25
        v15 = v15
        v15 = 0.0
        _ = v15
    }
    if (-1 >= -1) {
        v3 = 153
        v13 = v13
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS2{F0: 1, F1: 0}
            _ = fi16
            _ = v17
        }
    }
    var v18 uint32 = 4294967295
    fmt.Println(v3)
    var v19 bool = false
    v20 := GenS3{F0: -1}
    var v21 uint = 255
    var v22 float32 = (1.5 * 3.14159)
    fmt.Println(v21, v13)
    var v23 float32 = -2.25
    v19 = true
    v24 := GenS4{F0: 42, F1: 0, F2: v3}
    v25 := GenS5{F0: 0, F1: (-28 == -1)}
    var v26 bool = (!false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
