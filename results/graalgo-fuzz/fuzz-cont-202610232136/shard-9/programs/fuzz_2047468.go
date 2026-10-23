package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
func genFn0(p0 int64) int8 {
    v4 := GenS1{F0: 1}
    var v5 int16 = 42
    switch "foo bar" {
    case "":
        var v6 int32 = -2147483648
        _ = v6
    case "hello":
        fmt.Println(p0)
    case "a":
        var v7 float64 = (0.0 * (1e10 - (0.1 * 0.1)))
        _ = v7
    }
    switch "\n\t" {
    case "hello":
        v5 += 859
    default:
        p0 -= p0
    }
    _ = p0
    _ = v4
    _ = v5
    return int8(v5)
}
func genFn1(p0 int16, p1 string, p2 int) int16 {
    var v8 uint = 768
    fmt.Println(p2, v8)
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{p2, fi9, 32767, fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 289
}

func main() {
    var v1 uint8 = 250
    v2 := make(map[int]string)
    var v3 int64 = -1
    v3 += v3
    v11 := GenS2{F0: ((float32((1e10 * (0.0 / 3.14159))) + float32(float64(1e10))) == ((3.14159 + 0.1) * 0.0)), F1: 670, F2: 2}
    v12 := []int{2147483647}
    if false {
        delete(v2, 32767)
    }
    v3 = v3
    if true {
        v13 := GenS2{F0: (464 < -9223372036854775808), F1: 18446744073709551615, F2: 42}
        v3 = (v3 * 585)
        v3 = v3
        _ = v13
    } else {
        var v14 float64 = -2.25
        _ = v14
    }
    v1 -= v1
    v3 += v3
    var v15 uint8 = 128
    var v16 float32 = (1.5 / (0.1 - 0.0))
    if false {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        delete(v2, 2147483647)
    }
    var v17 uint32 = uint32(v15)
    _ = len(v2)
    v17 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
