package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }

func main() {
    var v1 int = 32767
    var v2 uint = 0
    var v3 uint16 = 255
    if (!true) {
        v2 = v2
        switch "" {
        case "\n\t":
            v4 := GenS1{F0: true, F1: 163, F2: false}
            _ = v4
        default:
            v5 := []int{v1, (^(v1)), 538, (^(v1))}
            _ = v5
        }
        if true {
            v3 = 32767
            fmt.Println(v1)
        }
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int16 = 42
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 uint32 = 517
    var v10 int64 = 246
    if false {
        v11 := GenS2{F0: 0, F1: 42, F2: v1}
        switch ("\n\t" + "hello") {
        case "a":
            var v12 float64 = -2.25
            _ = v12
        default:
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS3{F0: 1, F1: v3, F2: (v2 >> ((uint(v3)) & 63))}
                v2 -= (v2 + 0)
                var v15 int8 = -128
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        var v16 uint32 = (v9 / ((793) | 1))
        _ = v11
        _ = v16
    } else {
        v17 := []int{(v1 + v1), v1, (v1 % ((v1) | 1)), v1}
        _ = v17
    }
    var v18 string = ""
    v19 := []int{v1, -435, v1, 2}
    v20 := GenS2{F0: uint8(v9), F1: 127, F2: v1}
    for fi21 := 0; fi21 < 10; fi21++ {
        _ = len(v19)
        var v22 int16 = 1
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
