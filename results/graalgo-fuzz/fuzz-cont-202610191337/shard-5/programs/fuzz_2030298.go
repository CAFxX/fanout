package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 0
    var v3 uint = 18446744073709551615
    v4 := make(map[int]uint32)
    var v5 int8 = -75
    if (2 >= 426) {
        v5 -= (-(v5))
    } else {
        v6 := []int{}
        _ = v6
    }
    fmt.Println(v2, v2, v3)
    var v7 uint32 = 42
    var v8 uint64 = (^(v2))
    fmt.Println(v3, v5)
    v9 := []int{-9223372036854775808, -718, 127}
    v5 = v5
    v10 := GenS1{F0: (!false), F1: v5}
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            _ = len(v1)
            switch 2147483647 {
            case 128:
                v9 = append(v9, fi12)
            case 127:
                fmt.Println(fi12)
            default:
                var v13 string = ""
                v5 = 104
                v14 := GenS1{F0: (!false), F1: (^(v5))}
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        fmt.Println(v2)
        var v15 uint32 = v7
        _ = fi11
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
