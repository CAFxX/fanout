package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
func genFn0() bool {
    var v4 string = "a"
    var v5 uint64 = 2
    var v6 int16 = 936
    v7 := GenS1{F0: 1000, F1: -1}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := []int{128, -9223372036854775808, 127, 0}
    var v2 uint32 = 196
    var v3 int16 = -1
    var v8 uint = 65535
    v9 := GenS1{F0: 42, F1: 2147483647}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 uint8 = 140
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi10
        _ = v11
    }
    var v12 int16 = 42
    var v13 uint = (v8 - v8)
    fmt.Println(v3)
    var v14 int64 = 42
    v15 := GenS1{F0: (^(v3)), F1: int32(v14)}
    v16 := GenS1{F0: -1, F1: -1}
    for fi17 := 0; fi17 < 4; fi17++ {
        v8 -= v8
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi17
    }
    v3 -= (v12 / ((v12) | 1))
    v18 := GenS1{F0: (v3 >> ((2) & 63)), F1: 1}
    v3 = 903
    fmt.Println(v12, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
