package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32, p1 bool, p2 int32) bool {
    var v7 int16 = -32768
    v8 := GenS1{F0: 42, F1: p0}
    fmt.Println(v7, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}
func genFn1() int16 {
    switch 128 {
    case 519:
        var v9 float64 = -2.25
        _ = v9
    }
    if true {
        if (-2.25 < 0.1) {
            var v10 int = -794
            v11 := []int{}
            var v12 uint32 = 664
            _ = v10
            _ = v11
            _ = v12
        } else {
            v13 := []int{-1, 4294967295, 4294967295, 32767}
            if false {
                if (0 > 0) {
                    var v14 uint32 = 4294967295
                    _ = v14
                }
            }
            var v15 uint = 318
            _ = v13
            _ = v15
        }
        var v16 float32 = (-2.25 - 0.0)
        _ = v16
    }
    var v17 int8 = 0
    fmt.Println(v17)
    var v18 bool = (860 < int32(v17))
    _ = v17
    _ = v18
    return -32768
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = float32(-2.25)
    v3 := make(map[uint32]string)
    var v4 uint16 = 499
    var v5 float64 = 3.14159
    var v6 uint64 = 18446744073709551615
    switch -646 {
    case 127:
        delete(v3, 2147483647)
    case 128:
        fmt.Println(v2, v6)
    case -9223372036854775808:
        var v19 float32 = ((((float32((0.1 * (v5 * (((((float64((0.1 * -2.25)) * ((v5 - v5) * (-2.25 * (0.1 + -2.25)))) + float64(v2)) + v5) / v5) * float64((v2 / 1e10)))))) * v2) + (v2 + ((3.14159 - 1e10) * 3.14159))) + ((v2 + v2) + (v2 - 0.0))) - v2)
        _ = v19
    default:
        v20 := []int{826}
        _ = v20
    }
    var v21 uint = 4294967295
    v22 := GenS2{F0: 887, F1: 797, F2: -2147483648}
    if (false && true) {
        v23 := GenS3{F0: ""}
        v24 := []int{127, 9223372036854775807}
        var v25 int8 = -12
        _ = v23
        _ = v24
        _ = v25
    } else {
        var v26 uint8 = 240
        fmt.Println(v4, v5)
        v5 -= ((((((((v5 + v5) / v5) * -2.25) * 1.5) / (v5 - (v5 / v5))) + (1e10 / (0.0 * (1.5 / v5)))) - v5) * v5)
        _ = v26
    }
    fmt.Println(v5)
    for fi27 := 0; fi27 < 8; fi27++ {
        v6 -= v6
        var v28 uint32 = 850
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        v4 -= (v4 & (v4 | v4))
        if (!false) {
            v30 := GenS1{F0: (v4 | (v4 + (v4 & uint16(v21)))), F1: 42}
            for fi31 := 0; fi31 < 15; fi31++ {
                if true {
                    v32 := GenS2{F0: 0, F1: 2147483648, F2: (fi29 + (^(fi31)))}
                    _ = v32
                } else {
                    v33 := []int{}
                    _ = v33
                }
                _ = fi31
            }
            for fi34 := 0; fi34 < 9; fi34++ {
                v4 += 1
                v35 := GenS4{F0: 1}
                _ = fi34
                _ = v35
            }
            _ = v30
        } else {
            var v36 uint32 = 2
            for fi37 := 0; fi37 < 13; fi37++ {
                v5 += (float64(1e10) - (float64(v2) - (float64(v2) + 1e10)))
                v38 := v3[1]
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        _ = fi29
    }
    fmt.Println(v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
