package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483647, 14, 115, 879}
    var v3 int8 = 0
    var v4 int8 = (-(v3))
    v5 := make(map[int]int)
    if false {
        if false {
            v6 := GenS1{F0: (^(v3)), F1: 255}
            v4 = 0
            _ = v6
        } else {
            v3 -= (v4 * v4)
            var v7 int8 = (-(v3))
            v8 := GenS1{F0: v4, F1: 2}
            _ = v7
            _ = v8
        }
    } else {
        var v9 string = "hello"
        var v10 uint16 = 255
        _ = v10
        _ = v9
    }
    if (false || true) {
        v11 := GenS1{F0: v4, F1: 971}
        var v12 float32 = 0.0
        var v13 int = 1
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint8 = 39
    var v15 int16 = 32767
    fmt.Println(v3, v14)
    var v16 uint64 = 4294967295
    v17 := GenS1{F0: (^(v3)), F1: 255}
    var v18 int = 333
    fmt.Println(v14, v4, v18)
    v19 := GenS2{F0: (v4 - (v4 + (^(v3)))), F1: true}
    if true {
        var v20 float64 = (0.0 / 0.1)
        if true {
            var v21 string = ((("a" + "\n\t") + (("" + "hello") + "hello")) + "foo bar")
            var v22 int16 = -1
            var v23 int8 = (v3 | -1)
            _ = v21
            _ = v22
            _ = v23
        }
        v18 -= -9223372036854775808
        _ = v20
    }
    v4 += (-(v3))
    var v24 bool = (0.1 == (1.5 * 3.14159))
    for fi25 := 0; fi25 < 4; fi25++ {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        var v26 bool = v24
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 uint64 = (v16 / ((1) | 1))
        v28 -= (^(v16))
        v24 = ((!true) || false)
        _ = fi27
        _ = v28
    }
    var v29 float64 = float64(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
