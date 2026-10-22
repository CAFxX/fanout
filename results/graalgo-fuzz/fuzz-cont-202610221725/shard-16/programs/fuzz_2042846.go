package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 float32 = (-2.25 - 1.5)
    var v4 uint8 = 128
    v5 := GenS1{F0: 4294967295, F1: int16(v4), F2: 0}
    v6 := GenS1{F0: 42, F1: -32768, F2: 127}
    var v7 int8 = 42
    v8 := []int{}
    var v9 uint16 = 65535
    var v10 uint8 = (v4 % (((^(v4))) | 1))
    if false {
        v9 -= v9
        v11 := GenS1{F0: 901, F1: 1, F2: v7}
        var v12 uint16 = v9
        _ = v11
        _ = v12
    } else {
        var v13 float32 = v3
        _ = v13
    }
    var v14 float32 = float32(0.1)
    if false {
        v15 := GenS1{F0: 42, F1: 32767, F2: (v7 << ((577) & 63))}
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
