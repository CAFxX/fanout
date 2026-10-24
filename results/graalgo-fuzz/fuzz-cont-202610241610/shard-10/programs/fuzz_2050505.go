package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint = 0
    v3 := make(map[string]string)
    v1 -= (^(v1))
    for fi4 := 0; fi4 < 7; fi4++ {
        v3["\n\t"] = "\n\t"
        _ = fi4
    }
    v1 += (v1 + -9223372036854775808)
    var v5 uint8 = uint8(v2)
    v1 = (v1 % (((v1 << ((v2) & 63))) | 1))
    var v6 bool = false
    fmt.Println(v5, v1, v1)
    if ("\n\t" >= "foo bar") {
        v7 := GenS1{F0: 255, F1: (0.1 > float32(1.5))}
        var v8 uint8 = (v5 >> ((358) & 63))
        _ = len(v3)
        _ = v7
        _ = v8
    } else {
        switch "\n\t" {
        case "a":
            for fi9 := 0; fi9 < 16; fi9++ {
                v1 += int64(v2)
                if v6 {
                    var v10 int16 = -124
                    v5 = v5
                    _ = v10
                }
                _ = fi9
            }
        default:
            var v11 float32 = ((1.5 - (0.1 + (float32(((float64(0.0) - 1e10) - 0.0)) / 1.0))) - 1e10)
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := []int{}
                fmt.Println(v6)
                fmt.Println(v2, fi12, fi13)
                _ = fi13
                _ = v14
            }
            var v15 uint = (v2 << ((v2) & 63))
            var v16 float32 = (1.5 * 0.1)
            _ = fi12
            _ = v15
            _ = v16
        }
        var v17 int32 = 2147483647
        _ = v17
    }
    if v6 {
        var v18 uint64 = 211
        var v19 uint8 = (v5 | (v5 / ((42) | 1)))
        _ = v18
        _ = v19
    }
    var v20 int64 = v1
    var v21 float64 = (3.14159 / 1.0)
    if (v21 <= v21) {
        v20 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v5
    _ = v6
}
