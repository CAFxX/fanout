package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint; F2 uint }

func main() {
    v1 := []uint32{4294967295, 155, 177, 2}
    var v2 uint = 277
    var v3 float64 = (float64(1.5) + 0.0)
    switch "\n\t" {
    case "foo bar":
        for fi4 := 0; fi4 < 1; fi4++ {
            v3 = (((v3 * (v3 * (((0.0 - 3.14159) * v3) / float64(float32((1.5 * v3)))))) / 3.14159) - v3)
            _ = fi4
        }
    case "hello":
        v3 += v3
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 uint16 = 906
            v2 = (^(v2))
            v7 := GenS1{F0: -128, F1: "\n\t"}
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    v2 -= v2
    for fi8 := 0; fi8 < 11; fi8++ {
        switch "" {
        case "\n\t":
            var v9 bool = false
            _ = v9
        case "":
            if len(v1) > 0 {
                v1[0] = 48
            }
        case "hello":
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint16 = 1
                v12 := GenS1{F0: 42, F1: ("hello" + "hello")}
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            var v13 int32 = 288
            _ = v13
        }
        _ = fi8
    }
    var v14 float32 = 0.0
    _ = len(v1)
    _ = len(v1)
    v15 := GenS1{F0: 80, F1: ("" + "foo bar")}
    v16 := GenS2{F0: 255}
    var v17 bool = (!(!(32767 < 4294967295)))
    v18 := GenS3{F0: 359, F1: 4294967295, F2: v2}
    v14 += v14
    v2 = v2
    var v19 uint8 = 127
    fmt.Println(v19, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
