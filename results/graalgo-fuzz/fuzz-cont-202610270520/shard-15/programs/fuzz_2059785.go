package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0(p0 int64) int {
    var v4 int8 = 42
    fmt.Println(v4, p0)
    v5 := []int{int(p0), -9223372036854775808}
    p0 -= (p0 - (p0 / ((311) | 1)))
    var v6 float32 = 0.1
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 960
}

func main() {
    var v1 int = -509
    var v2 bool = true
    var v3 int32 = -1
    v7 := GenS1{F0: 627}
    if (-2.25 >= 1.5) {
        v1 = v1
        if true {
            v8 := GenS1{F0: 876}
            fmt.Println(v3, v3)
            _ = v8
        } else {
            v9 := []int{-1, -329, 128}
            _ = v9
        }
        fmt.Println(v3)
    }
    v10 := []int{(v1 ^ v1)}
    fmt.Println(v1, v3, v3)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 float32 = 1.5
        _ = fi11
        _ = v12
    }
    v1 -= 42
    var v13 int = v1
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(fi14, v2, fi14)
        _ = fi14
    }
    v13 -= genFn0(865)
    v15 := GenS2{F0: genFn0(0), F1: -1}
    var v16 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
