package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 int64, p1 int8, p2 int64) int16 {
    v7 := GenS1{F0: 32767}
    var v8 float32 = (3.14159 + 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 uint16) uint16 {
    v9 := GenS1{F0: genFn0(1, 1, 9223372036854775807)}
    var v10 bool = (((0 <= 4294967295) || true) || (42 == 1))
    var v11 int16 = 1
    var v12 int32 = 42
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 uint = 128
    var v4 string = "hello"
    var v5 int16 = 561
    var v6 int16 = 42
    v4 = v4
    for fi13 := 0; fi13 < 14; fi13++ {
        v3 = (v3 | v3)
        for fi14 := 0; fi14 < 2; fi14++ {
            v6 += 354
            var v15 string = v4
            var v16 float64 = (float64(0.0) + 1e10)
            if (2 < 754) {
                fmt.Println(fi14, v3)
                v17 := []int{(fi14 | (fi13 >> (((v3 | v3)) & 63))), (fi13 << ((v3) & 63)), fi14, -645}
                _ = v17
            } else {
                v18 := GenS1{F0: v5}
                v19 := []int{32767}
                v6 += v6
                v16 += -2.25
                _ = v18
                _ = v19
            }
            var v20 uint = 188
            _ = fi14
            _ = v15
            _ = v16
            _ = v20
        }
        if true {
            v21 := GenS1{F0: v5}
            var v22 uint64 = 189
            _ = v21
            _ = v22
        } else {
            var v23 float64 = float64((1.5 + 0.1))
            _ = v23
        }
        _ = fi13
    }
    v24 := []int{0, 4294967295, 2147483647, 2147483647}
    v25 := GenS2{F0: 42}
    var v26 float32 = (float32((0.0 - (((1e10 / 1e10) - 1.5) / 1.0))) * -2.25)
    v5 += v6
    if (v3 < (v3 * (v3 + (^(v3))))) {
        var v27 float64 = 0.0
        v28 := GenS3{F0: false, F1: 2147483647}
        fmt.Println(v27)
        _ = v27
        _ = v28
    } else {
        v3 -= 1
        var v29 uint = 4294967295
        v30 := GenS3{F0: false, F1: 179}
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        if false {
            if (fi31 > 200) {
                switch v4 {
                case "hello":
                    if ((4294967295 >= v3) || true) {
                        v6 -= v6
                    }
                }
                switch v4 {
                case "":
                    v3 = 4294967295
                default:
                    var v32 uint = v3
                    _ = v32
                }
            }
        } else {
            v4 = ("a" + v4)
            if (65535 > 0) {
                fmt.Println(v5)
                var v33 bool = (42 > 0)
                _ = v33
            }
        }
        fmt.Println(v3)
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
