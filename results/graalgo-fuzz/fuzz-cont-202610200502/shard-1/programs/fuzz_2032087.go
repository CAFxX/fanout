package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = float32(3.14159)
    var v3 int8 = -87
    var v4 float64 = 0.0
    var v5 int64 = 1
    v6 := make(map[int]string)
    if false {
        switch "\n\t" {
        case "foo bar":
            v7 := GenS1{F0: -306, F1: 981, F2: 297}
            _ = v7
        default:
            v1 = (v1 | (v1 ^ 2))
        }
    } else {
        delete(v6, -1)
    }
    var v8 int = 1
    if true {
        var v9 uint32 = (v1 - 658)
        if false {
            v10 := GenS1{F0: 2147483647, F1: (v8 | v8), F2: 32767}
            _ = v10
        }
        _ = v9
    }
    var v11 uint = 439
    v12 := GenS2{F0: v11, F1: ("" + "a")}
    switch (-(v8)) {
    case -2147483648:
        v13 := GenS3{F0: (v11 % ((616) | 1)), F1: 42}
        _ = v13
    case 128:
        var v14 int32 = 42
        _ = v14
    default:
        v15 := v6[127]
        _ = v15
    }
    v16 := []int{9223372036854775807, (v8 >> ((255) & 63)), v8, v8}
    fmt.Println(v3)
    v3 = (-(v3))
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 float64 = v4
        _ = fi17
        _ = v18
    }
    v3 = (v3 ^ v3)
    fmt.Println(v5, v5)
    v3 = v3
    v6[42] = "\n\t"
    v4 += (0.0 * 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
