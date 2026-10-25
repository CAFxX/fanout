package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }

func main() {
    var v1 string = "a"
    var v2 uint32 = 2147483647
    var v3 uint8 = 242
    var v4 uint = 18446744073709551615
    var v5 string = (v1 + v1)
    v6 := make(map[int]bool)
    v2 -= v2
    v5 = ""
    v7 := GenS1{F0: (1e10 >= 3.14159)}
    switch 140 {
    case 32768:
        v2 -= v2
    }
    switch 516 {
    case -9223372036854775808:
        v4 = v4
    default:
        fmt.Println(v4)
    }
    v8 := []int{202}
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: (884 >= 18446744073709551615)}
        fmt.Println(v1, v4)
        v5 = v1
        _ = fi9
        _ = v10
    }
    fmt.Println(v3, v2, v5)
    switch "hello" {
    case "hello":
        switch -636 {
        case 4294967295:
            var v11 uint64 = 874
            _ = v11
        case 1:
            v12 := GenS2{F0: v2, F1: (^(v2)), F2: (^(v2))}
            _ = v12
        case 128:
            fmt.Println(v2)
        }
    case "\n\t":
        for fi13 := 0; fi13 < 8; fi13++ {
            v3 = 26
            _ = fi13
        }
    case "a":
        _ = len(v8)
    }
    v14 := GenS2{F0: v2, F1: 544, F2: (^(v2))}
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
