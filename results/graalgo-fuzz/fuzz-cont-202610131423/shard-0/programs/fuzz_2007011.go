package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
func genFn0(p0 int64, p1 uint8, p2 uint) int8 {
    p1 -= 128
    var v7 int = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -71
}
func genFn1(p0 int32, p1 uint8, p2 uint32) int32 {
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: p0, F1: 2, F2: -1}
        var v10 int16 = 704
        v11 := []int{}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(p1)
    var v12 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (p0 + (-(p0)))
}

func main() {
    var v1 float64 = 1.5
    var v2 string = "hello"
    var v3 bool = (50 >= -45)
    v4 := make(map[uint32]int)
    var v5 uint = 2
    v6 := make(map[int]bool)
    v1 -= v1
    var v13 float32 = 1e10
    fmt.Println(v3, v2, v5)
    v14 := GenS2{F0: -32768}
    v15 := []int{9223372036854775807, 2147483647, 1}
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 float32 = ((3.14159 - v13) - float32(v1))
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 uint8 = 70
        switch (^(fi18)) {
        case 32767:
            fmt.Println(v13)
        case -2147483648:
            var v20 string = v2
            _ = v20
        default:
            fmt.Println(fi18, v13, v1)
        }
        _ = fi18
        _ = v19
    }
    v21 := v6[2147483647]
    v5 = (v5 | 65535)
    v21 = false
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            if v21 {
                var v24 int = 364
                _ = v24
            }
            _ = fi23
        }
        v3 = (107 <= 0)
        if v3 {
            switch (fi22 & fi22) {
            case 32768:
                var v25 uint16 = 751
                _ = v25
            default:
                fmt.Println(fi22, v2)
            }
        }
        _ = fi22
    }
    fmt.Println(v1, v21, v5)
    switch ("hello" + v2) {
    case "hello":
        v26 := GenS3{F0: (0 >= 817), F1: -84, F2: 69}
        _ = v26
    }
    var v27 bool = (-738 < -1)
    switch (v2 + v2) {
    case "a":
        v2 = (v2 + v2)
    case "":
        v1 = (1e10 * float64(float32((v1 - 0.0))))
    case "foo bar":
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
