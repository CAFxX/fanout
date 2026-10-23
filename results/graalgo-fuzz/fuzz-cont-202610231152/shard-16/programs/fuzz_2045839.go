package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{65535, 65535, 341}
    v3 := make(map[int]uint32)
    var v4 int8 = 42
    v5 := []uint32{528, 376}
    v1 = v1
    v4 = v4
    var v6 int32 = 42
    switch 32768 {
    case -9223372036854775808:
        v4 += (v4 % (((v4 + (^(v4)))) | 1))
    case 127:
        var v7 int8 = v4
        _ = v7
    }
    delete(v3, 127)
    v4 += v4
    var v8 int64 = 0
    switch int(v4) {
    case 65:
        fmt.Println(v4, v8, v6)
    case -427:
        v8 -= v8
    default:
        var v9 string = v1
        _ = v9
    }
    v10 := GenS1{F0: (v6 % ((155) | 1)), F1: 1}
    fmt.Println(v8)
    var v11 int64 = v8
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint32 = 259
        v11 = v11
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
