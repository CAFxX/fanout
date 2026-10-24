package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (-2.25 >= (((((-2.25 + float64(-2.25)) + 1e10) - 0.1) + ((3.14159 * 0.1) / 1.0)) + (-2.25 / 1.0)))
    v3 := make(map[uint32]int)
    var v4 int = -215
    v5 := []int{488, -9223372036854775808, (v4 * v4), v4}
    var v6 int32 = 725
    v7 := v1["a"]
    v8 := GenS1{F0: -9223372036854775808}
    delete(v3, 0)
    fmt.Println(v2)
    var v9 bool = v2
    switch "a" {
    case "\n\t":
        fmt.Println(v2, v6)
    default:
        fmt.Println(v4, v4, v6)
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 bool = v7
        _ = fi10
        _ = v11
    }
    fmt.Println(v6)
    v12 := GenS2{F0: 255}
    fmt.Println(v6)
    v9 = (v9 && (!(!(!((0.1 != 0.1) && (v6 != (-(v6))))))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
