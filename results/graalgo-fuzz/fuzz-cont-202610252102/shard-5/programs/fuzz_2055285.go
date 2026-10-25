package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
func genFn0(p0 int32) int16 {
    p0 -= p0
    v4 := []int{42}
    _ = len(v4)
    var v5 uint16 = 0
    v6 := GenS1{F0: false, F1: 2147483647, F2: 1000}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -792
}

func main() {
    var v1 string = ("hello" + "hello")
    v2 := []int{32767, 128, 128, 32768}
    v3 := make(map[int]bool)
    var v7 uint = 18446744073709551615
    var v8 uint = v7
    v7 -= (v8 / (((^(v7))) | 1))
    var v9 bool = false
    if (240 >= 128) {
        switch 128 {
        case -9223372036854775808:
            v8 = (^(v8))
        case -718:
            var v10 uint64 = 1
            _ = v10
        }
        var v11 int32 = 1
        var v12 int8 = -128
        _ = v11
        _ = v12
    }
    var v13 uint16 = 481
    v14 := v3[32768]
    for fi15 := 0; fi15 < 15; fi15++ {
        v13 = (^(v13))
        _ = fi15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS1{F0: v9, F1: -2147483648, F2: -32768}
        _ = len(v3)
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
