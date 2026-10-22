package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint8; F2 bool }
func genFn0(p0 uint16, p1 int64, p2 int8) bool {
    p0 -= 220
    switch 1 {
    case 32768:
        v5 := []int{128, -9223372036854775808, 128, 42}
        _ = v5
    case 1:
        v6 := GenS1{F0: 624, F1: 644, F2: 255}
        _ = v6
    case 9223372036854775807:
        p0 -= p0
    }
    v7 := []int{127, 1}
    if true {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if ((p1 / (((p1 / ((p1) | 1))) | 1)) != 1) {
            fmt.Println(p1)
            p2 = int8(p0)
            var v8 string = (("a" + "hello") + "\n\t")
            _ = v8
        }
    } else {
        _ = len(v7)
        if (!(!true)) {
            if false {
                v9 := GenS2{F0: 32767, F1: 2147483647}
                var v10 uint16 = p0
                p2 -= p2
                _ = v10
                _ = v9
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        } else {
            var v11 uint64 = 2
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}
func genFn1(p0 bool) int32 {
    p0 = true
    var v12 uint16 = 0
    _ = p0
    _ = v12
    return -2147483648
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 1
    v3 := []int{128, 695}
    var v4 float32 = 0.0
    v13 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 18446744073709551615}
    v4 = float32(3.14159)
    var v14 uint = 4294967295
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(fi15)
        var v16 string = ""
        _ = fi15
        _ = v16
    }
    fmt.Println(v14, v2)
    var v17 int16 = -575
    var v18 float64 = (3.14159 - 0.1)
    var v19 int = -1
    if ((v14 / ((65535) | 1)) >= v14) {
        if (80 > 202) {
            fmt.Println(v17, v19, v14)
            v17 -= (v17 >> ((255) & 63))
        }
        var v20 int16 = v17
        switch "foo bar" {
        case "a":
            if len(v3) > 0 {
                v3[0] = (-(v19))
            }
        case "":
            v1 = append(v1, 42)
        case "hello":
            var v21 int8 = 100
            _ = v21
        default:
            var v22 int16 = v20
            _ = v22
        }
        _ = v20
    }
    var v23 uint64 = 18446744073709551615
    var v24 int64 = 1
    v25 := GenS1{F0: (v23 - v23), F1: (v19 >> ((v14) & 63)), F2: 537}
    v26 := GenS3{F0: 21}
    if false {
        var v27 int = (v19 & v19)
        _ = v27
    } else {
        _ = len(v1)
        var v28 float64 = 1e10
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(v19, fi29)
        var v30 uint32 = 0
        v31 := GenS4{F0: v30, F1: (v2 >> ((65535) & 63)), F2: false}
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
