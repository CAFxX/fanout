package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64; F1 string; F2 string }
func genFn0(p0 int32) int32 {
    p0 = (-(p0))
    var v6 float32 = float32(3.14159)
    p0 = p0
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 665, F1: 0}
            v10 := GenS1{F0: 18446744073709551615, F1: 2}
            v11 := []int{-1, fi7, fi8, -1}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := []int{-1, (fi12 / (((fi12 / ((fi12) | 1))) | 1))}
        fmt.Println(v6)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v6
    return (p0 * (p0 * p0))
}

func main() {
    var v1 string = "a"
    var v2 int16 = 790
    var v3 uint = 4294967295
    v4 := make(map[string]int)
    var v5 float64 = 1.5
    if true {
        v1 = v1
    }
    var v14 uint8 = 255
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int = fi15
        _ = fi15
        _ = v16
    }
    switch 864 {
    case 9223372036854775807:
        if false {
            v17 := GenS1{F0: 67, F1: 115}
            v18 := GenS1{F0: 1, F1: 2}
            _ = v17
            _ = v18
        } else {
            _ = len(v4)
            v19 := GenS1{F0: 4294967295, F1: 0}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS2{F0: -1, F1: 28, F2: (v14 | (v14 + 255))}
        v1 = v1
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS1{F0: 2, F1: 916}
        _ = fi22
        _ = v23
    }
    var v24 bool = ((0.0 > 1e10) || (899 != 2))
    switch -1 {
    case 32767:
        fmt.Println(v24, v1)
    default:
        if v24 {
            v25 := []int{}
            _ = v25
        }
    }
    if v24 {
        v26 := GenS3{F0: -667, F1: (v1 + v1), F2: v1}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
