package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    var v3 uint64 = 0
    var v4 int64 = -9223372036854775808
    var v5 uint = 459
    var v6 int32 = 2147483647
    v3 += (v3 >> ((v5) & 63))
    v7 := []int{-503, 986, 42, 993}
    fmt.Println(v5, v5, v4)
    var v8 string = "\n\t"
    var v9 uint16 = 0
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    switch v8 {
    case "a":
        v12 := GenS1{F0: 169, F1: 2147483647, F2: v3}
        _ = v12
    case "foo bar":
        if (v3 <= v3) {
            for fi13 := 0; fi13 < 6; fi13++ {
                v6 -= (v6 ^ (-(v6)))
                v5 = v5
                var v14 int32 = 438
                _ = fi13
                _ = v14
            }
        }
    case "hello":
        fmt.Println(v3, v9)
    default:
        fmt.Println(v9, v4)
    }
    var v15 bool = true
    fmt.Println(v6, v3)
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint = v5
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: fi18, F1: 4294967295, F2: 1}
            var v20 float64 = 0.0
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 float64 = ((-2.25 * (1.5 / -2.25)) + 3.14159)
        _ = fi16
        _ = v17
        _ = v21
    }
    var v22 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
