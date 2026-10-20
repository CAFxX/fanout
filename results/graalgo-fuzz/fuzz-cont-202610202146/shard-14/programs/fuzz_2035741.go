package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint64 }

func main() {
    var v1 int8 = 1
    var v2 int8 = (v1 * -128)
    v3 := make(map[int]string)
    var v4 int8 = v2
    v5 := make(map[int]bool)
    fmt.Println(v1, v4)
    v1 += 0
    var v6 uint32 = 2147483647
    switch "hello" {
    case "hello":
        var v7 bool = true
        _ = v7
    default:
        v8 := GenS1{F0: -773, F1: 42}
        _ = v8
    }
    v9 := []int{0, -2147483648}
    if (true && (7 < 4294967295)) {
        v2 += v2
        var v10 uint = 2
        _ = v10
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            v6 = 2
            v12 := GenS2{F0: 4294967295, F1: 18446744073709551615, F2: uint64(v6)}
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 uint16 = 32767
                v5[fi11] = false
                var v15 float64 = 1.5
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v2, v6, v1)
    fmt.Println(v1, v6, v4)
    v16 := GenS1{F0: -333, F1: 801}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
