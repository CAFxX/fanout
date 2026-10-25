package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 string }
func genFn0() int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 float32 = float32(3.14159)
        v8 := GenS1{F0: 0}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int16 = 32767
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{(fi10 ^ -2147483648), (^(fi10)), (fi10 / ((fi10) | 1)), fi10}
        var v12 bool = (0 >= 0)
        v9 += 0
        _ = fi10
        _ = v11
        _ = v12
    }
    v9 = v9
    v13 := GenS2{F0: 0, F1: "\n\t"}
    _ = v13
    _ = v9
    return -128
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{32768, 32768}
    v3 := []int{-9223372036854775808, 42}
    var v4 uint = 4294967295
    var v5 uint8 = 255
    switch 2147483647 {
    case 0:
        for fi14 := 0; fi14 < 17; fi14++ {
            _ = len(v3)
            _ = fi14
        }
    }
    fmt.Println(v4, v4)
    v5 += 1
    v5 = (v5 ^ (v5 + v5))
    v5 -= (v5 + 255)
    if true {
        var v15 int8 = 42
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int32 = 817
            _ = fi16
            _ = v17
        }
        var v18 float64 = 0.0
        _ = v15
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 string = ""
        _ = fi19
        _ = v20
    }
    v1["foo bar"] = 1
    var v21 uint64 = 1
    v21 += (v21 | v21)
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 float64 = 0.1
        _ = fi22
        _ = v23
    }
    var v24 int = -8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
