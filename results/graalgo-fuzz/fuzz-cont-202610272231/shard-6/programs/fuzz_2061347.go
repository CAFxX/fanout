package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }

func main() {
    var v1 float32 = ((float32(1.5) + 1.5) - 1e10)
    var v2 float32 = (v1 + v1)
    var v3 float32 = 3.14159
    var v4 bool = true
    fmt.Println(v3)
    v1 = float32(0.1)
    fmt.Println(v1, v3, v4)
    switch "\n\t" {
    case "foo bar":
        v1 = (float32((float64(float32(float64(0.1))) - 0.0)) * (((1e10 + (v2 * 1.5)) * (v3 / (1.5 / 3.14159))) + v1))
    case "hello":
        if v4 {
            var v5 uint = 65535
            _ = v5
        }
    case "\n\t":
        fmt.Println(v2, v2, v4)
    default:
        v1 = ((v3 / v1) * v2)
    }
    var v6 uint16 = 0
    v7 := []int{32767, 9223372036854775807, 0}
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v3)
        _ = fi8
    }
    v7 = append(v7, 305)
    var v9 int = 127
    switch "\n\t" {
    case "":
        var v10 int16 = 1000
        _ = v10
    default:
        fmt.Println(v3, v4, v1)
    }
    switch v9 {
    case -1:
        switch ("hello" + "") {
        case "hello":
            v4 = (!false)
        default:
            var v11 uint64 = 42
            _ = v11
        }
    case -433:
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS1{F0: 65535, F1: 2, F2: -980}
            _ = fi12
            _ = v13
        }
    case 42:
        if len(v7) > 0 {
            v7[0] = (v9 / ((v9) | 1))
        }
    default:
        if (993 > 252) {
            if true {
                v1 -= (v3 / v1)
                for fi14 := 0; fi14 < 17; fi14++ {
                    v9 += fi14
                    _ = fi14
                }
            }
            fmt.Println(v4, v3, v1)
        } else {
            var v15 int32 = -1
            v7 = append(v7, 9223372036854775807)
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint64 = 42
                switch "" {
                case "a":
                    v17 -= v17
                }
                var v18 int16 = -32768
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
