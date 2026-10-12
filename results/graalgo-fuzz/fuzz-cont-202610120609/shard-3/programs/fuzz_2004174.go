package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int = 127
    var v2 uint32 = 323
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    v5 := []string{"foo bar", "", "foo bar", ""}
    v6 := make(map[int]string)
    v7 := GenS1{F0: 0}
    v8 := GenS2{F0: 478}
    v1 -= (v1 - v1)
    v9 := GenS1{F0: 451}
    fmt.Println(v1, v1)
    v10 := GenS2{F0: 800}
    var v11 float64 = (((1.5 * (((0.1 + (float64(0.0) / 1.0)) - 0.0) * 1e10)) * (0.1 + (((1.5 + float64(0.1)) - 1.5) / 1e10))) / 1.0)
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint32 = (^(v2))
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
