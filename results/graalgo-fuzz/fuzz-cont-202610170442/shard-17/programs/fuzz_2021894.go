package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 uint }

func main() {
    v1 := []int{-2147483648}
    var v2 int64 = -595
    var v3 uint16 = 686
    v4 := make(map[uint32]bool)
    var v5 float64 = 0.1
    var v6 uint64 = 208
    var v7 uint8 = 127
    var v8 float32 = float32((v5 * ((v5 * (((-2.25 + 1.5) / v5) * ((-2.25 * v5) * (-2.25 + 0.1)))) / 3.14159)))
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            if (false || (!false)) {
                if (false || false) {
                    v4[1] = false
                    var v11 uint32 = 0
                    v12 := GenS1{F0: -655, F1: 104}
                    _ = v11
                    _ = v12
                } else {
                    v7 -= v7
                }
            }
            fmt.Println(fi10, fi9, v7)
            var v13 uint64 = (v6 - (v6 & v6))
            _ = fi10
            _ = v13
        }
        v14 := GenS1{F0: 283, F1: 2}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi9
        _ = v14
    }
    _ = len(v4)
    v1 = append(v1, 128)
    if false {
        v8 = v8
        switch "\n\t" {
        case "":
            var v15 string = ("hello" + "")
            _ = v15
        case "\n\t":
            var v16 string = ""
            _ = v16
        case "hello":
            v17 := GenS2{F0: 0, F1: -1}
            _ = v17
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            if false {
                v2 -= 9223372036854775807
                v5 -= 3.14159
                v19 := GenS3{F0: v2, F1: 2}
                _ = v19
            }
            v6 += 363
            _ = fi18
        }
    } else {
        fmt.Println(v6, v2, v6)
        var v20 uint = uint(v6)
        _ = v20
    }
    fmt.Println(v5, v8)
    fmt.Println(v5)
    if (32767 >= 0) {
        for fi21 := 0; fi21 < 5; fi21++ {
            v5 = ((v5 * (float64(0.1) * v5)) + (v5 - v5))
            _ = fi21
        }
        fmt.Println(v8)
    } else {
        v22 := GenS2{F0: 2147483647, F1: 0}
        switch -2147483648 {
        case -541:
            v23 := GenS3{F0: v2, F1: 42}
            _ = v23
        case 32767:
            fmt.Println(v8)
        default:
            v3 += 854
        }
        _ = v22
    }
    if ((float64((v8 / v8)) * v5) != ((v5 - (v5 * 0.1)) * float64((v8 * v8)))) {
        var v24 float32 = v8
        _ = v24
    }
    var v25 uint32 = 174
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
