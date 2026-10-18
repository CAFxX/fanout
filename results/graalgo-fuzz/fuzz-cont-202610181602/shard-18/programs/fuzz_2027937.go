package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    var v1 int32 = 1
    var v2 uint32 = 137
    var v3 uint8 = 127
    v4 := make(map[int]int)
    var v5 int16 = 32767
    var v6 int32 = v1
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint8 = (v3 / ((v3) | 1))
        fmt.Println(v8, v2, v5)
        var v9 uint64 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 = v2
    fmt.Println(v3)
    v1 = (v6 >> ((0) & 63))
    v4[-1] = -27
    if (0 <= 63) {
        v2 += (v2 - 1)
        fmt.Println(v6, v5)
    } else {
        v5 = v5
    }
    v10 := GenS1{F0: -9223372036854775808, F1: -1}
    var v11 int32 = v6
    var v12 int8 = 0
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 float64 = 3.14159
        v1 -= v1
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint16 = 42
        _ = len(v4)
        v1 = (v1 + v6)
        _ = fi17
        _ = v18
    }
    switch 32768 {
    case 1:
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 float64 = (3.14159 / 1.0)
            _ = fi19
            _ = v20
        }
    default:
        v6 = (v11 * (v11 - (v6 ^ int32(v2))))
    }
    fmt.Println(v3)
    var v21 int = -745
    switch "a" {
    case "hello":
        var v22 uint8 = v3
        _ = v22
    case "":
        var v23 int64 = 901
        _ = v23
    default:
        for fi24 := 0; fi24 < 11; fi24++ {
            v3 = v3
            if (false && true) {
                var v25 float64 = 0.0
                _ = v25
            } else {
                var v26 int64 = -9223372036854775808
                v21 += 32768
                _ = v26
            }
            v1 = (v11 % ((v1) | 1))
            _ = fi24
        }
    }
    v2 = uint32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
