package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64) uint64 {
    p0 = (p0 % ((590) | 1))
    p0 = p0
    var v6 string = ("foo bar" + "hello")
    v7 := GenS1{F0: 576, F1: 1}
    v8 := []int{128, 478}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 int8 = 56
    var v2 int64 = -1
    v3 := make(map[string]uint32)
    var v4 string = "a"
    var v5 int16 = 0
    var v9 uint8 = 255
    var v10 float32 = 1e10
    v11 := GenS1{F0: 0, F1: 65535}
    var v12 int32 = 0
    if false {
        v13 := GenS1{F0: 65535, F1: 634}
        fmt.Println(v1)
        _ = v13
    } else {
        v14 := GenS1{F0: 2147483648, F1: 88}
        switch -1 {
        case -9223372036854775808:
            v15 := []int{81}
            _ = v15
        }
        v16 := GenS1{F0: 172, F1: 0}
        _ = v14
        _ = v16
    }
    _ = len(v3)
    var v17 int8 = 0
    fmt.Println(v5, v1, v10)
    if (!true) {
        var v18 int = 286
        v4 = (v4 + ((v4 + "") + "hello"))
        _ = v18
    } else {
        v9 = (v9 & 174)
        v17 = v17
    }
    v1 -= -40
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
