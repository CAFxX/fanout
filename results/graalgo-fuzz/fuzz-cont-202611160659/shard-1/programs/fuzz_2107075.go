package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0(p0 uint32, p1 int8, p2 uint8) uint {
    v5 := []int{995, 4294967295, 555, 0}
    _ = len(v5)
    var v6 float64 = ((-2.25 + (((3.14159 + ((((-2.25 * (0.0 + ((float64((3.14159 - 1.5)) + 0.0) * 1.5))) + 1.5) * float64(1.5)) + (-2.25 / -2.25))) * (((1e10 + 1.5) / float64((3.14159 + 0.0))) / 1.0)) - float64((0.0 * (0.1 + 3.14159))))) + 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 int8 = -110
    v1 += v1
    fmt.Println(v1, v1)
    if true {
        fmt.Println(v4)
        var v7 int = 1
        _ = v7
    }
    if true {
        v1 -= v1
    }
    switch (("foo bar" + "") + "foo bar") {
    case "\n\t":
        var v8 float64 = ((1e10 - 0.1) - 3.14159)
        _ = v8
    default:
        v4 = 15
    }
    v9 := GenS1{F0: ((("foo bar" + "a") < "\n\t") && true), F1: 42}
    v10 := []int{-588}
    var v11 uint16 = 0
    _ = len(v10)
    if (0.1 == 0.1) {
        var v12 uint16 = 860
        v4 = v4
        _ = v12
    }
    fmt.Println(v1, v4, v4)
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(v4, fi13, v11)
        fmt.Println(v1)
        var v14 uint = 412
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
