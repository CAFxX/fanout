package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = -2147483648
    var v2 uint = 1000
    v3 := make(map[string]int)
    var v4 bool = (0 <= 42)
    var v5 bool = ((783 < 255) && v4)
    v6 := make(map[uint32]string)
    v7 := []int{}
    var v8 int32 = 654
    v5 = v4
    v1 = (v1 % (((v1 * (v1 | v1))) | 1))
    v9 := GenS1{F0: -195}
    fmt.Println(v8, v5)
    switch "a" {
    case "\n\t":
        v2 -= v2
    case "":
        v8 -= -2147483648
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 float64 = 1.5
        switch "foo bar" {
        case "foo bar":
            v12 := GenS1{F0: -1}
            _ = v12
        case "\n\t":
            delete(v3, "\n\t")
        default:
            var v13 uint16 = 2
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v1 += (v1 | v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
