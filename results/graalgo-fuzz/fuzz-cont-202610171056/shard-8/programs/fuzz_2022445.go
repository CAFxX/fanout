package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 uint32 = 446
    var v2 uint = uint(v1)
    var v3 uint16 = 398
    fmt.Println(v2)
    v4 := []int{32767, -9223372036854775808, 0}
    var v5 float64 = (3.14159 * 1.5)
    v3 += (v3 % ((v3) | 1))
    v6 := GenS1{F0: 177, F1: -57}
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 += (v3 - v3)
            fmt.Println(fi7, v3, v1)
            _ = fi8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v2)
            _ = fi9
        }
        switch "" {
        case "foo bar":
            var v10 int16 = 32767
            _ = v10
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi7
    }
    v2 = v2
    v2 += v2
    var v11 uint16 = (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
