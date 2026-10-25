package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }

func main() {
    v1 := []int{-1, -1, 42}
    var v2 int32 = -1
    v3 := []uint32{42, 2147483648, 2147483648}
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    switch 32767 {
    case -1:
        if ((true && true) && true) {
            var v6 bool = false
            _ = v6
        }
    case 128:
        v2 = (v2 | v2)
    default:
        var v7 int64 = -785
        _ = v7
    }
    if true {
        v8 := GenS1{F0: 766}
        var v9 int16 = 32767
        v10 := GenS2{F0: v9, F1: (!(!false)), F2: 127}
        _ = v10
        _ = v8
        _ = v9
    } else {
        var v11 int16 = 860
        var v12 float32 = 0.0
        _ = v11
        _ = v12
    }
    fmt.Println(v2)
    v2 += (v2 % ((v2) | 1))
    if len(v3) > 0 {
        v3[0] = 0
    }
    if (42 <= 42) {
        for fi13 := 0; fi13 < 10; fi13++ {
            if len(v1) > 0 {
                v1[0] = (fi13 ^ 2147483647)
            }
            v14 := GenS1{F0: 18446744073709551615}
            v2 -= (v2 / ((v2) | 1))
            _ = fi13
            _ = v14
        }
        var v15 string = ("a" + "foo bar")
        _ = v15
    } else {
        v16 := GenS1{F0: 18446744073709551615}
        _ = v16
    }
    var v17 uint8 = 128
    if false {
        var v18 int8 = 100
        var v19 uint = 2
        _ = v18
        _ = v19
    }
    var v20 uint = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
