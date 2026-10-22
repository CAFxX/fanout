package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 bool = true
    var v2 int64 = -1
    v3 := make(map[uint32]int)
    var v4 uint16 = 2
    switch -9223372036854775808 {
    case 4294967295:
        switch 32768 {
        case 18:
            v2 = (v2 * v2)
        default:
            for fi5 := 0; fi5 < 8; fi5++ {
                v2 += 357
                var v6 bool = (0 < v4)
                _ = fi5
                _ = v6
            }
        }
    default:
        switch 0 {
        case 9223372036854775807:
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS1{F0: 492, F1: uint64(v4)}
                _ = fi7
                _ = v8
            }
        default:
            v1 = (!v1)
        }
    }
    v9 := v3[827]
    var v10 bool = false
    fmt.Println(v4)
    v4 -= (v4 * v4)
    var v11 float64 = -2.25
    v12 := []int{(v9 ^ v9), 128, (v9 % (((v9 % ((v9) | 1))) | 1)), (-(v9))}
    v1 = (!false)
    var v13 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
