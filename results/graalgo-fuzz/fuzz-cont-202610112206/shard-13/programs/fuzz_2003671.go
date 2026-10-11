package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32 }
func genFn0() uint16 {
    var v6 int32 = 0
    v7 := GenS1{F0: 2147483647, F1: 42}
    v8 := []int{32768, -2147483648, 514}
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []string{"hello"}
    v2 := []int{2147483647, -971}
    v3 := []string{"hello"}
    var v4 float64 = 0.0
    v5 := make(map[string]uint32)
    v4 = float64(1e10)
    var v9 uint8 = 146
    var v10 uint64 = 4294967295
    v10 = (v10 + v10)
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := v5[("\n\t" + "")]
        _ = fi11
        _ = v12
    }
    v13 := v5["a"]
    var v14 uint = 538
    var v15 uint64 = 1
    v16 := GenS2{F0: 2147483647}
    fmt.Println(v13)
    v17 := GenS1{F0: 9223372036854775807, F1: v15}
    v4 -= 1e10
    if (!(415 != (v14 / ((v14) | 1)))) {
        fmt.Println(v9)
        fmt.Println(v4, v9, v9)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v15 -= 18446744073709551615
    var v18 uint = uint(v10)
    v14 = (v18 | v18)
    var v19 int = -830
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
