package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 bool }
func genFn0(p0 bool, p1 int8, p2 int32) uint8 {
    var v4 int16 = 0
    var v5 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 225
}
func genFn1(p0 int, p1 string) uint32 {
    v6 := GenS1{F0: 255, F1: 559, F2: true}
    p0 = p0
    p1 = ("a" + ((p1 + (p1 + p1)) + p1))
    fmt.Println(p1, p1)
    var v7 int64 = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 bool = (2 >= 32767)
    var v2 uint16 = 2
    var v3 int16 = 32767
    v8 := []int{347}
    if v1 {
        v3 = 42
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = len(v8)
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v2 = (v2 / ((v2) | 1))
    v1 = v1
    var v9 uint64 = 4294967295
    fmt.Println(v2, v3, v9)
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v1)
        var v11 float32 = (0.0 + -2.25)
        _ = fi10
        _ = v11
    }
    var v12 int32 = 1
    fmt.Println(v2, v12)
    fmt.Println(v3)
    fmt.Println(v2, v2)
    var v13 string = "hello"
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(v13, v9, v1)
        v15 := GenS1{F0: 127, F1: 2147483648, F2: (v1 && v1)}
        var v16 int32 = (^(v12))
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
