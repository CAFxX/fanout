package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 uint, p1 int8, p2 int8) int32 {
    var v5 bool = ((true || true) && false)
    p0 = (^(p0))
    if v5 {
        var v6 uint8 = 1
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    var v3 uint64 = 1
    var v4 int8 = 42
    _ = len(v1)
    v3 -= v3
    switch 127 {
    case -9223372036854775808:
        if true {
            var v7 int8 = v4
            _ = v7
        }
    }
    var v8 string = ""
    var v9 float32 = ((1.5 * 3.14159) * 1e10)
    var v10 uint8 = 128
    v1["hello"] = 2
    if ((1 >= 9223372036854775807) && ((v10 > v10) || ((!v2) && true))) {
        v11 := GenS1{F0: v10}
        v12 := GenS2{F0: (v4 ^ v4), F1: 42}
        _ = v11
        _ = v12
    }
    v3 = 0
    var v13 float32 = v9
    v14 := GenS2{F0: v4, F1: 1}
    v15 := []int{1, 9223372036854775807}
    var v16 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
