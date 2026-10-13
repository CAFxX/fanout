package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int8 }
func genFn0(p0 uint) string {
    v6 := []int{}
    v7 := GenS1{F0: (p0 >= 2), F1: 1, F2: -18}
    _ = p0
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := []uint32{42, 0, 0}
    var v4 int = 921
    var v5 int16 = 1000
    v8 := GenS1{F0: (1 >= 2147483647), F1: 0, F2: 0}
    v9 := GenS1{F0: (v4 < v4), F1: 61, F2: -128}
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v10 float64 = (0.0 / 1.0)
    v2[739] = "a"
    _ = len(v2)
    fmt.Println(v4, v10, v10)
    var v11 float64 = (1e10 / v10)
    for fi12 := 0; fi12 < 19; fi12++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v13 int = fi12
        v4 = -1
        _ = fi12
        _ = v13
    }
    switch "hello" {
    case "foo bar":
        v4 = v4
    case "hello":
        var v14 int32 = 1
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
