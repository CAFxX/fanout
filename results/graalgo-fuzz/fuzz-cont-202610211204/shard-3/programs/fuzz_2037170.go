package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int32 = 693
    v2 := make(map[int]int)
    var v3 int64 = 430
    var v4 uint16 = 118
    var v5 uint32 = 806
    fmt.Println(v1)
    var v6 int32 = (v1 % (((-(v1))) | 1))
    fmt.Println(v4, v3, v4)
    var v7 int32 = v1
    for fi8 := 0; fi8 < 3; fi8++ {
        v5 -= 212
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{(-(fi8)), 9223372036854775807}
            v1 = (v7 | v1)
            _ = fi9
            _ = v10
        }
        v11 := []int{}
        _ = fi8
        _ = v11
    }
    v12 := GenS1{F0: 65535, F1: (v1 % (((v7 ^ (v1 | v1))) | 1)), F2: 2}
    if (false && true) {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint8 = 127
            _ = fi13
            _ = v14
        }
        var v15 int32 = int32(v4)
        _ = v15
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v5, v7)
    case "foo bar":
        switch 32767 {
        case 9223372036854775807:
            fmt.Println(v4, v6)
        }
    case "":
        var v16 bool = true
        _ = v16
    default:
        var v17 int = 32767
        _ = v17
    }
    var v18 int8 = 121
    if (32767 < 1000) {
        v4 += v4
        var v19 uint = 4294967295
        v6 -= (v6 & 959)
        _ = v19
    } else {
        _ = len(v2)
        delete(v2, -940)
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if ((-2.25 + 1e10) > ((0.1 - ((-2.25 * 3.14159) * 0.1)) + 1e10)) {
            v21 := GenS2{F0: "\n\t", F1: 1000}
            var v22 uint8 = 127
            _ = v21
            _ = v22
        } else {
            v23 := GenS2{F0: "\n\t", F1: 0}
            _ = v23
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            fmt.Println(v7, v5)
            v3 -= (v3 >> ((2) & 63))
            var v25 float64 = float64((float32(0.1) + (1.5 * -2.25)))
            v3 = (v3 % (((v3 + v3)) | 1))
            fmt.Println(v18, v4, v25)
            var v26 int32 = (v7 ^ v7)
            v27 := []int{127, fi20, (fi24 - fi24)}
            var v28 bool = ((!(220 == 44)) || false)
            v27 = append(v27, (fi24 * fi20))
            if (((v25 - (1e10 / v25)) + v25) <= (v25 * (1e10 * 1e10))) {
                v18 = v18
                v29 := GenS3{F0: 65}
                v25 -= float64(((3.14159 + 3.14159) - 0.1))
                _ = v29
            }
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
            _ = v28
        }
        var v30 uint64 = 395
        _ = fi20
        _ = v30
    }
    v7 += v7
    for fi31 := 0; fi31 < 20; fi31++ {
        switch "\n\t" {
        case "foo bar":
            for fi32 := 0; fi32 < 19; fi32++ {
                v1 = 42
                var v33 float32 = 1e10
                v33 -= v33
                _ = fi32
                _ = v33
            }
        }
        v34 := GenS4{F0: 903, F1: v5}
        var v35 int64 = v3
        _ = fi31
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
