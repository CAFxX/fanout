package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
func genFn0() int64 {
    var v5 int32 = -1
    v5 = 470
    var v6 int64 = 1
    v6 += (v6 * (v6 - 1))
    v7 := GenS1{F0: 1}
    _ = v5
    _ = v6
    _ = v7
    return (-(v6))
}
func genFn1() int64 {
    var v8 int64 = -17
    switch "a" {
    case "\n\t":
        var v9 uint = 18446744073709551615
        _ = v9
    default:
        v10 := []int{2147483647, 624, -2147483648}
        _ = v10
    }
    if (0 > -1) {
        var v11 string = "hello"
        _ = v11
    } else {
        fmt.Println(v8)
    }
    var v12 int8 = -69
    _ = v12
    _ = v8
    return genFn0()
}

func main() {
    var v1 int16 = 583
    var v2 int8 = 77
    var v3 int64 = -234
    var v4 float64 = (((0.0 * (((1.5 + float64(1e10)) + (-2.25 + 0.0)) - -2.25)) + 1.5) + 1.5)
    fmt.Println(v4)
    v13 := GenS2{F0: "hello", F1: 1, F2: 694}
    v3 = 0
    var v14 uint8 = 128
    var v15 int32 = 42
    if true {
        var v16 float32 = 0.1
        fmt.Println(v1, v14, v14)
        v17 := GenS1{F0: v2}
        _ = v16
        _ = v17
    }
    v4 = (0.1 / ((float64(0.0) + ((3.14159 + 3.14159) / (float64(-2.25) - (((v4 - 1e10) + v4) + v4)))) * -2.25))
    if true {
        switch -2147483648 {
        case 250:
            if ((0 <= 42) && false) {
                v18 := []int{863, 32767, 42, -639}
                _ = v18
            } else {
                v19 := []int{0, 4294967295}
                fmt.Println(v15, v2)
                v4 = (((float64(float32((v4 - (float64(1e10) * v4)))) + -2.25) + (float64((-2.25 / float32((v4 * v4)))) / 3.14159)) - v4)
                _ = v19
            }
        case 0:
            v20 := GenS3{F0: (v14 ^ v14)}
            _ = v20
        default:
            var v21 float64 = 0.0
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := GenS4{F0: 0}
        _ = fi22
        _ = v23
    }
    v24 := GenS4{F0: v3}
    v25 := []int{-9223372036854775808}
    v3 -= 1
    fmt.Println(v15, v15, v1)
    if (false && true) {
        var v26 int = int(v2)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
