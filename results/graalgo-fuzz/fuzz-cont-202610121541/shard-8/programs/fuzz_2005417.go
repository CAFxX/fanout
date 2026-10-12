package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0() int32 {
    var v7 float64 = 3.14159
    switch 2147483647 {
    case 0:
        v7 = v7
    case 2147483647:
        fmt.Println(v7)
    case 1:
        v7 -= 0.1
    }
    if false {
        v8 := []int{-2147483648, 2147483647, -938, 9223372036854775807}
        if len(v8) > 0 {
            v8[0] = 4294967295
        }
        var v9 string = ("foo bar" + "")
        _ = v8
        _ = v9
    } else {
        v7 = v7
    }
    v10 := GenS1{F0: 32767, F1: -1}
    _ = v10
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 int8) int {
    if (false || false) {
        var v11 bool = p0
        fmt.Println(p0, p0, p1)
        _ = v11
    }
    switch -2147483648 {
    case -2147483648:
        fmt.Println(p0, p0)
    case 2147483647:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := []uint32{362, 4294967295, 423, 0}
    var v2 int = -9223372036854775808
    v3 := make(map[uint32]bool)
    v4 := []uint32{664, 381}
    var v5 int = (v2 + v2)
    v6 := []uint32{42, 0, 142, 65535}
    delete(v3, 2147483648)
    v5 -= (v5 & (v5 - 4294967295))
    for fi12 := 0; fi12 < 9; fi12++ {
        if true {
            var v13 bool = false
            _ = v13
        }
        v14 := v3[902]
        var v15 uint8 = 160
        _ = fi12
        _ = v14
        _ = v15
    }
    v5 -= (v5 | (v5 | (^(v2))))
    v2 -= genFn1(false, -128)
    v2 -= genFn1((1 < 652), 42)
    var v16 float32 = -2.25
    fmt.Println(v16, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
