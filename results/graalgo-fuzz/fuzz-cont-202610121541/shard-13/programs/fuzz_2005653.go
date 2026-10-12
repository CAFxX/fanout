package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 int16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 int64) int64 {
    var v5 int8 = 1
    var v6 int64 = p0
    v7 := []int{128, 9223372036854775807, 4294967295}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (-(p0))
}

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = (((0.0 - 0.1) + 0.1) / (0.1 / 1e10))
    var v3 int32 = 42
    v4 := make(map[uint32]uint32)
    var v8 float32 = (-2.25 / (float32((0.1 + (0.1 + (3.14159 - 0.1)))) + (((v2 * v2) * v2) + ((v2 + (0.0 * float32(float64(0.0)))) - float32(float64(float32(0.1)))))))
    v9 := GenS1{F0: 612, F1: 0}
    v8 = ((-2.25 + v8) - 0.0)
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v8, v3, v2)
        _ = fi10
    }
    v11 := []int{128, -992, -9223372036854775808}
    fmt.Println(v1, v1, v8)
    var v12 int32 = (v3 >> ((42) & 63))
    if true {
        switch 367 {
        case 42:
            if true {
                var v13 int8 = -88
                for fi14 := 0; fi14 < 6; fi14++ {
                    if true {
                        fmt.Println(v12, v13)
                    } else {
                        fmt.Println(v1)
                        var v15 float32 = 0.1
                        _ = v15
                    }
                    _ = fi14
                }
                _ = v13
            } else {
                fmt.Println(v8, v1, v12)
            }
        case 143:
            delete(v4, (v1 % ((v1) | 1)))
        case 32767:
            for fi16 := 0; fi16 < 6; fi16++ {
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS2{F0: "a", F1: 2, F2: -32768}
                    v11 = append(v11, fi17)
                    var v19 uint = 65535
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 uint64 = 4294967295
                _ = fi16
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 int64 = genFn0(9223372036854775807)
                if true {
                    v23 := GenS3{F0: -381, F1: (fi21 << ((42) & 63))}
                    _ = v23
                } else {
                    var v24 int64 = genFn0(v22)
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
        }
    } else {
        var v25 int32 = -1
        var v26 float64 = (-2.25 * -2.25)
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        delete(v4, 494)
        var v28 float64 = 0.0
        _ = fi27
        _ = v28
    }
    fmt.Println(v1, v3)
    for fi29 := 0; fi29 < 13; fi29++ {
        if true {
            v30 := GenS4{F0: fi29, F1: 32767}
            _ = v30
        }
        _ = fi29
    }
    _ = len(v4)
    v1 -= (v1 - (v1 / ((4294967295) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
