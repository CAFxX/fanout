package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 bool }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    var v3 int64 = 9223372036854775807
    v4 := []string{}
    _ = len(v4)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 42, F1: -2147483648, F2: (v3 % ((v3) | 1))}
        switch "hello" {
        case "foo bar":
            fmt.Println(v2)
        default:
            var v7 int32 = 2147483647
            _ = v7
        }
        v3 = v3
        _ = fi5
        _ = v6
    }
    v8 := GenS2{F0: 4294967295, F1: 2147483647, F2: (-306 != -75)}
    var v9 int = 128
    fmt.Println(v9)
    v10 := GenS2{F0: 647, F1: 4294967295, F2: true}
    v9 = 128
    v3 = v3
    fmt.Println(v3, v2, v2)
    if (250 >= 2147483648) {
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 float64 = 0.0
                v3 += (v3 * -9223372036854775808)
                _ = fi12
                _ = v13
            }
            fmt.Println(v3)
            v14 := v1[v9]
            _ = fi11
            _ = v14
        }
    }
    v3 -= 565
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
