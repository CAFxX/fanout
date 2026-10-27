package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int }
func genFn0() uint64 {
    v7 := GenS1{F0: -32768, F1: 42}
    for fi8 := 0; fi8 < 6; fi8++ {
        switch "" {
        case "a":
            var v9 string = "hello"
            _ = v9
        case "foo bar":
            var v10 uint32 = 351
            _ = v10
        default:
            var v11 int64 = 1
            _ = v11
        }
        _ = fi8
    }
    v12 := GenS1{F0: -1, F1: 18446744073709551615}
    var v13 uint64 = 375
    _ = v12
    _ = v13
    _ = v7
    return (v13 + (v13 - (v13 | 42)))
}

func main() {
    var v1 float64 = (3.14159 - (3.14159 + (0.0 + ((0.1 / 1.5) - 0.0))))
    v2 := []string{}
    var v3 int8 = -110
    v4 := make(map[uint32]bool)
    v5 := []int{-798, 0, 714, 9223372036854775807}
    var v6 int32 = 311
    switch "\n\t" {
    case "hello":
        var v14 float32 = (float32(float64(0.1)) + 1.5)
        _ = v14
    case "foo bar":
        v15 := GenS2{F0: 255}
        _ = v15
    }
    v3 = -128
    v4[65535] = false
    v16 := GenS3{F0: 2, F1: 9223372036854775807}
    var v17 bool = (0.0 >= 0.1)
    var v18 string = ((("" + "a") + "a") + "a")
    fmt.Println(v6)
    var v19 uint8 = 127
    switch 2147483647 {
    case -746:
        fmt.Println(v6, v17)
    }
    v3 -= 127
    v20 := GenS2{F0: uint(v19)}
    v3 -= v3
    _ = len(v4)
    var v21 uint = 255
    v22 := GenS2{F0: (v21 % (((v21 ^ (v21 ^ 0))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
