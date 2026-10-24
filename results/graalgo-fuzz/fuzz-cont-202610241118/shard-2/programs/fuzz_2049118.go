package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int; F1 uint64; F2 string }

func main() {
    var v1 int8 = 1
    var v2 int32 = -2147483648
    var v3 float32 = 3.14159
    v4 := []int{1, -619, -9223372036854775808}
    var v5 int = 42
    switch "hello" {
    case "":
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: "\n\t", F1: 198, F2: false}
            v3 -= 0.1
            if false {
                var v8 uint8 = 2
                switch "a" {
                case "hello":
                    fmt.Println(v8, v3)
                case "\n\t":
                    v1 += v1
                default:
                    v3 += (float32(1e10) + v3)
                }
                _ = v8
            } else {
                v9 := GenS1{F0: "", F1: 128, F2: (0.0 >= 1.5)}
                v10 := GenS2{F0: v5, F1: 1, F2: "\n\t"}
                v3 = ((-2.25 / (float32(float64(v3)) * v3)) - v3)
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    default:
        fmt.Println(v2, v1, v5)
    }
    var v11 uint16 = 2
    for fi12 := 0; fi12 < 8; fi12++ {
        v5 -= (-(v5))
        v3 += v3
        v2 = v2
        _ = fi12
    }
    var v13 int8 = 100
    var v14 int32 = int32(v11)
    v11 += (v11 * v11)
    for fi15 := 0; fi15 < 3; fi15++ {
        switch v5 {
        case -1:
            v3 = v3
        case -240:
            v14 = 2147483647
        default:
            var v16 uint = 882
            _ = v16
        }
        _ = fi15
    }
    var v17 float32 = v3
    switch v5 {
    case 32767:
        if len(v4) > 0 {
            v4[0] = (-(v5))
        }
    case 128:
        fmt.Println(v13)
    }
    var v18 float32 = 1e10
    v11 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
