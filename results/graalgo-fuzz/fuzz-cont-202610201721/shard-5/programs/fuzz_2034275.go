package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint) int16 {
    p0 = 4294967295
    fmt.Println(p0)
    if true {
        v6 := []int{9223372036854775807, -170}
        _ = v6
    } else {
        v7 := GenS1{F0: true, F1: 32767}
        _ = v7
    }
    _ = p0
    return 1000
}

func main() {
    var v1 uint8 = 196
    var v2 int32 = 0
    var v3 string = ""
    var v4 int16 = 154
    v5 := make(map[int]int)
    fmt.Println(v2)
    var v8 float64 = float64(1e10)
    var v9 uint32 = 2147483648
    fmt.Println(v8, v8, v9)
    v10 := []int{127, 1, 856}
    for fi11 := 0; fi11 < 18; fi11++ {
        v1 = v1
        _ = fi11
    }
    v3 = ""
    v2 = (-(v2))
    v12 := GenS1{F0: (true && true), F1: 255}
    var v13 float64 = v8
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
