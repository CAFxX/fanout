package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint64; F1 bool; F2 bool }

func main() {
    var v1 float64 = float64(float32((1.5 - 1.5)))
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    fmt.Println(v1, v1)
    if true {
        v4 := []int{-166, -1}
        fmt.Println(v1, v2)
        _ = v4
    }
    v5 := GenS1{F0: 11}
    switch "foo bar" {
    case "foo bar":
        if false {
            for fi6 := 0; fi6 < 1; fi6++ {
                v7 := GenS2{F0: v2}
                if true {
                    var v8 uint32 = 576
                    var v9 int16 = -32768
                    _ = v8
                    _ = v9
                } else {
                    v10 := GenS3{F0: v2, F1: 0}
                    var v11 int16 = int16(fi6)
                    v12 := GenS1{F0: 593}
                    _ = v10
                    _ = v11
                    _ = v12
                }
                _ = fi6
                _ = v7
            }
        }
    case "hello":
        v1 += v1
    default:
        v2 = v2
    }
    switch v2 {
    case "":
        v2 = "hello"
    case "\n\t":
        for fi13 := 0; fi13 < 16; fi13++ {
            fmt.Println(fi13, v2)
            v1 = v1
            _ = fi13
        }
    case "hello":
        switch ((("a" + v2) + v2) + ("hello" + (v2 + "a"))) {
        case "hello":
            var v14 string = v2
            _ = v14
        }
    default:
        delete(v3, 42)
    }
    v15 := []int{-9223372036854775808, 0, 0}
    switch 32767 {
    case 42:
        v16 := GenS4{F0: 1, F1: (((-2.25 - ((v1 / (v1 / ((v1 * (0.1 / (((((-2.25 * 0.1) / 3.14159) - v1) - (v1 + (v1 - float64(1.5)))) * (v1 / (((((v1 * float64((((float32((v1 / 1.0)) - 0.0) + 3.14159) - 0.0))) + float64(0.1)) + v1) - 3.14159) - float64(float32((float64(3.14159) / ((v1 + (float64(0.0) - v1)) + (((v1 / v1) * v1) - ((v1 - v1) - (float64(0.0) / v1)))))))))))) - float64((1.5 + (0.1 * 0.0)))))) / v1)) / v1) == (v1 / 1e10)), F2: (664 > 1)}
        _ = v16
    case 273:
        v17 := GenS2{F0: v2}
        _ = v17
    }
    v2 = "\n\t"
    var v18 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v5
}
