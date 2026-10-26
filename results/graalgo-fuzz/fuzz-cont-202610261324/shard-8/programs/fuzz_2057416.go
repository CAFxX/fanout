package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = 0.0
    var v3 int64 = 1
    var v4 int = 32768
    var v5 uint32 = 1
    var v6 uint64 = 658
    switch (v4 >> ((4294967295) & 63)) {
    case -838:
        var v7 float32 = (0.1 - 3.14159)
        _ = v7
    }
    v8 := GenS1{F0: 669, F1: 829}
    fmt.Println(v2, v3, v4)
    v9 := []int{(-(v4)), (-(v4)), 0, 9223372036854775807}
    var v10 int32 = 1
    if true {
        fmt.Println(v5)
    }
    fmt.Println(v10, v5, v2)
    v4 = 0
    switch "a" {
    case "":
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(v1)
            if ((0 <= 74) || (v2 != 0.0)) {
                v4 = v4
            }
            _ = fi11
        }
    default:
        fmt.Println(v5, v6)
    }
    v12 := GenS2{F0: 86, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
