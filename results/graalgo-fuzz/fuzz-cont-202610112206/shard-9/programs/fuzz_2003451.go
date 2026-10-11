package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint8; F1 string }

func main() {
    v1 := []uint32{4294967295, 1, 65535, 2147483647}
    var v2 uint8 = 180
    var v3 float64 = float64(((-2.25 / float32(float64((1.5 - -2.25)))) * (float32((0.0 + 1.5)) * (-2.25 - 0.1))))
    var v4 bool = true
    var v5 bool = (-2147483648 != 1)
    v6 := []uint32{328, 4294967295, 232, 1}
    v2 += v2
    v7 := GenS1{F0: 42, F1: 42}
    var v8 uint8 = v2
    v9 := GenS2{F0: -312, F1: 0}
    var v10 uint = 247
    switch 128 {
    case -2147483648:
        v5 = v4
    case -858:
        fmt.Println(v4, v3, v5)
    case 380:
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS2{F0: -439, F1: 42}
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v8, v2)
    v5 = (1 != 42)
    var v13 int16 = -1
    v14 := GenS3{F0: (v8 + 2), F1: "foo bar"}
    v10 = v10
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint8 = v8
        v16 = 126
        _ = fi15
        _ = v16
    }
    v8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
