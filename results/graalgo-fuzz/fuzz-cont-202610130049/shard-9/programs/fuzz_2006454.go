package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 string) uint8 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{(^(fi5)), -478, fi5}
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p0)
    v7 := GenS1{F0: p1, F1: 2147483647}
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    var v1 int = 790
    v2 := make(map[string]bool)
    var v3 uint8 = 2
    var v4 bool = (42 <= 42)
    v1 = 127
    v4 = v4
    if (3.14159 == 0.0) {
        var v8 float64 = 1e10
        v4 = (!false)
        _ = v8
    }
    v4 = v4
    delete(v2, "")
    v4 = false
    v9 := []int{(^(v1)), (^(v1)), v1, -2147483648}
    v2["\n\t"] = (!true)
    v1 = v1
    v10 := v2["\n\t"]
    v1 += v1
    v10 = v4
    v11 := GenS2{F0: v1}
    v1 = (v1 / (((-(v1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
