package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v5 := []int{}
    v6 := make(map[string]string)
    v2[-547] = "a"
    v7 := GenS1{F0: 195, F1: 2}
    switch 128 {
    case 9223372036854775807:
        v1 = v1
    case -218:
        var v8 uint32 = 430
        _ = v8
    }
    v1 = v1
    v9 := GenS2{F0: 790}
    if true {
        v1 = ((3.14159 / ((v1 - v1) * -2.25)) * (0.1 - 1.5))
        v10 := v2[32767]
        _ = v10
    } else {
        if false {
            delete(v6, "")
            v11 := GenS3{F0: 127}
            switch "\n\t" {
            case "":
                v12 := GenS3{F0: 2147483647}
                _ = v12
            default:
                _ = len(v6)
            }
            _ = v11
        }
    }
    if (492 != -9223372036854775808) {
        _ = len(v6)
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v1 += v1
        v1 += v1
        _ = fi13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v1 = (3.14159 + float64((((1e10 * 3.14159) * (0.0 - ((0.0 - (0.1 * 0.0)) - ((1.5 / (1e10 + (float32(1.5) + 0.1))) + 0.0)))) + (-2.25 + 0.0))))
        _ = fi14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 string = ""
        _ = fi15
        _ = v16
    }
    delete(v2, -756)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
