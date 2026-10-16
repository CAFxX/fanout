package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 bool = (1000 >= 0)
    var v2 string = "\n\t"
    var v3 string = v2
    v4 := make(map[int]bool)
    var v5 uint64 = 18446744073709551615
    var v6 float32 = (-2.25 - 1e10)
    if (2147483647 >= 4294967295) {
        v3 = v2
    }
    v7 := []int{2147483647, 4294967295}
    fmt.Println(v6, v1, v2)
    var v8 uint64 = 516
    for fi9 := 0; fi9 < 11; fi9++ {
        _ = len(v4)
        _ = fi9
    }
    if true {
        var v10 float64 = 0.0
        for fi11 := 0; fi11 < 11; fi11++ {
            switch 2147483647 {
            case 2147483647:
                v5 = v8
            case 42:
                var v12 uint16 = 73
                _ = v12
            default:
                v13 := GenS1{F0: 42, F1: v3, F2: 713}
                _ = v13
            }
            _ = fi11
        }
        _ = v10
    }
    v6 = v6
    var v14 float64 = -2.25
    var v15 int8 = -128
    var v16 uint16 = 244
    fmt.Println(v14)
    v17 := GenS2{F0: 126}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
