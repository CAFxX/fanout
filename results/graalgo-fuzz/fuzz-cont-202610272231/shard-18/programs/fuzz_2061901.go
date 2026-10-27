package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
func genFn0() uint32 {
    var v5 int8 = 1
    v6 := []int{589, -9223372036854775808, int(v5), 4}
    var v7 uint16 = 32767
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() int32 {
    v8 := []int{-397, 4294967295, -81}
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    _ = v8
    return 102
}

func main() {
    var v1 int16 = 607
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 uint64 = 4294967295
    if true {
        var v10 int8 = -1
        var v11 int = 273
        var v12 uint16 = 32767
        _ = v10
        _ = v11
        _ = v12
    } else {
        v4 = 42
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := []int{(fi13 << ((590) & 63)), fi13}
            v4 += v4
            _ = fi13
            _ = v14
        }
    }
    v15 := v3[933]
    delete(v3, v15)
    v16 := GenS1{F0: (v15 * (v15 * v15)), F1: 42}
    switch ("foo bar" + ((("" + "hello") + "") + ("foo bar" + "a"))) {
    case "a":
        v4 += 42
    case "hello":
        v17 := GenS2{F0: 2, F1: 255, F2: 42}
        _ = v17
    case "\n\t":
        var v18 string = "foo bar"
        _ = v18
    default:
        fmt.Println(v2, v4, v2)
    }
    var v19 float32 = (1e10 + (1.5 + 0.1))
    var v20 uint32 = 4294967295
    var v21 uint8 = 0
    if v2 {
        for fi22 := 0; fi22 < 12; fi22++ {
            v2 = ((984 >= 0) || (false && v2))
            _ = fi22
        }
        v21 = v21
    }
    var v23 float64 = 0.1
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
