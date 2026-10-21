package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
func genFn0(p0 uint8, p1 int64, p2 string) int32 {
    var v5 uint8 = p0
    switch -978 {
    case 94:
        var v6 int = 32767
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 281
}

func main() {
    var v1 int = 32768
    v2 := make(map[string]string)
    var v3 int64 = 9223372036854775807
    var v4 int8 = -128
    v7 := GenS1{F0: 1000, F1: (-(v3)), F2: ((0 > 897) && (false && (128 != 0)))}
    var v8 uint32 = 2147483648
    fmt.Println(v4)
    var v9 uint8 = 128
    switch -1 {
    case 848:
        v10 := GenS1{F0: 1000, F1: v3, F2: (v4 == v4)}
        _ = v10
    case 0:
        v1 += (^(v1))
    default:
        var v11 string = ("" + "\n\t")
        _ = v11
    }
    v1 = v1
    if true {
        var v12 uint32 = 55
        _ = v12
    }
    v13 := GenS1{F0: -657, F1: int64(v4), F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
