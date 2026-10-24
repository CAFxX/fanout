package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
func genFn0(p0 int, p1 int32, p2 int64) string {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 int64 = 1
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 42
    v7 := []int{}
    var v8 uint = 2
    p2 = (p2 ^ p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (("" + "hello") + ("" + "foo bar"))
}

func main() {
    var v1 float32 = 0.1
    var v2 float32 = (float32((1e10 / 1.0)) * 1e10)
    var v3 float32 = -2.25
    fmt.Println(v2, v2)
    fmt.Println(v2, v3, v2)
    var v9 float32 = v3
    for fi10 := 0; fi10 < 17; fi10++ {
        v3 = v1
        var v11 uint64 = 135
        _ = fi10
        _ = v11
    }
    v12 := []int{-458, 127, 630, -9223372036854775808}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            _ = len(v12)
            _ = fi14
        }
        _ = fi13
    }
    v1 = ((v3 - v1) + v1)
    var v15 bool = true
    v9 -= v2
    v3 -= -2.25
    v16 := GenS1{F0: 32767, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
