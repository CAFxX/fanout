package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64, p1 uint, p2 uint16) bool {
    p0 += 4294967295
    var v7 bool = (!false)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    var v1 float64 = 1e10
    v2 := []int{-996, 4294967295, 2147483647}
    v3 := make(map[int]uint32)
    var v4 string = ("a" + "a")
    var v5 uint64 = 42
    var v6 float32 = 3.14159
    v8 := GenS1{F0: ((false || (0 != 0)) || false), F1: v5}
    if false {
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    fmt.Println(v1)
    var v9 int16 = 32767
    var v10 string = ("foo bar" + v4)
    v9 = (v9 % ((v9) | 1))
    v9 -= v9
    fmt.Println(v5, v4, v5)
    if (0.0 > v6) {
        var v11 float64 = v1
        _ = v11
    }
    var v12 float64 = v1
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int8 = 1
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 bool = true
        v17 := v3[(fi15 | fi15)]
        if (v4 < v4) {
            var v18 int32 = int32(v9)
            _ = v18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    var v19 bool = (-1 >= -1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
