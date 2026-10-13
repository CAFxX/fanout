package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    v1 := []int{-661, -9223372036854775808}
    var v2 uint32 = 42
    var v3 int8 = 42
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    var v6 uint32 = (v2 << ((4294967295) & 63))
    switch "foo bar" {
    case "\n\t":
        v7 := v4[-967]
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 10; fi8++ {
            v6 -= v2
            _ = fi8
        }
    default:
        v9 := GenS1{F0: (v3 >> ((4294967295) & 63)), F1: 1, F2: v2}
        _ = v9
    }
    var v10 bool = false
    v6 += (^(v6))
    if (1.5 == 3.14159) {
        v6 -= 2147483648
        v11 := GenS1{F0: v3, F1: 9223372036854775807, F2: (v6 % ((v2) | 1))}
        switch 47 {
        case 4:
            switch -596 {
            case -9223372036854775808:
                v12 := GenS2{F0: 1, F1: "foo bar"}
                _ = v12
            }
        }
        _ = v11
    } else {
        fmt.Println(v3, v10)
        v6 = v6
    }
    fmt.Println(v6, v3, v2)
    v2 += (v2 - 4294967295)
    _ = len(v4)
    v2 -= 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
