package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0() int16 {
    v7 := GenS1{F0: ("" + "")}
    v8 := GenS1{F0: "hello"}
    v9 := GenS2{F0: true, F1: 2}
    if (!true) {
        if (42 < 255) {
            var v10 int32 = -1
            var v11 uint16 = 104
            _ = v10
            _ = v11
        }
        var v12 uint64 = 4294967295
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint64 = 4294967295
        var v15 string = ("foo bar" + "hello")
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 0
    var v3 int8 = -128
    v4 := []int{-9223372036854775808, 9223372036854775807, -816, 559}
    var v5 int32 = v2
    var v6 float64 = -2.25
    fmt.Println(v3)
    v3 = -128
    v6 -= (1.5 + (1e10 - v6))
    var v16 string = ("a" + "")
    var v17 int16 = 1000
    for fi18 := 0; fi18 < 7; fi18++ {
        v17 = (v17 ^ v17)
        _ = fi18
    }
    fmt.Println(v2, v16)
    fmt.Println(v2, v16)
    fmt.Println(v6, v3)
    var v19 float32 = 0.1
    fmt.Println(v2, v6)
    v20 := GenS2{F0: true, F1: 4294967295}
    v19 -= (v19 / (v19 - 3.14159))
    var v21 uint16 = 2
    v22 := GenS1{F0: v16}
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
