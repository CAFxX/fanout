package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool) int16 {
    p0 = p0
    v6 := []int{-51}
    p0 = p0
    _ = p0
    _ = v6
    return 1000
}

func main() {
    var v1 int = 42
    v2 := make(map[string]int)
    var v3 float64 = ((0.0 + -2.25) / 1e10)
    var v4 bool = (!(true || false))
    var v5 uint8 = 221
    fmt.Println(v4, v4)
    v2["foo bar"] = 9223372036854775807
    v7 := []int{v1, (v1 + v1), 42, -9223372036854775808}
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            v4 = v4
            _ = fi9
        }
        fmt.Println(v1, v4)
        v3 = v3
        _ = fi8
    }
    v10 := GenS1{F0: 208}
    var v11 bool = (v4 || ((v5 / ((v5) | 1)) == 6))
    v4 = true
    v4 = false
    var v12 int16 = -1
    v4 = ((((3.14159 - v3) / (((3.14159 - v3) + float64(1.5)) / v3)) < v3) && (1 < 4294967295))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
