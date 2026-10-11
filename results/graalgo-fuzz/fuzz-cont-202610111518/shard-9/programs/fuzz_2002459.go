package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
func genFn0() uint {
    if false {
        var v7 uint32 = 2147483648
        _ = v7
    }
    var v8 uint8 = 2
    v9 := []int{-1}
    fmt.Println(v8)
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{}
    var v2 int = -2147483648
    var v3 int64 = 0
    v4 := make(map[int]uint32)
    var v5 int64 = v3
    var v6 int64 = v3
    v5 += v5
    var v10 int = v2
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v11 := GenS1{F0: genFn0(), F1: 65535}
    v1 = append(v1, "foo bar")
    var v12 int32 = 42
    v10 -= (v10 % ((v2) | 1))
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := v4[-748]
        var v15 uint32 = (v14 & v14)
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := v4[(v2 | (v2 << ((0) & 63)))]
    v10 -= v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
