package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
func genFn0() uint32 {
    v5 := GenS1{F0: 42, F1: -9223372036854775808, F2: 905}
    switch -9223372036854775808 {
    case 501:
        var v6 float64 = (0.0 / float64(-2.25))
        _ = v6
    case -1:
        var v7 uint16 = 255
        _ = v7
    default:
        var v8 string = ("hello" + ("a" + ("\n\t" + "\n\t")))
        _ = v8
    }
    var v9 bool = false
    _ = v5
    _ = v9
    return 107
}

func main() {
    v1 := []int{2147483647, 1, 4294967295}
    var v2 int16 = 32767
    v3 := make(map[string]uint32)
    var v4 uint8 = 235
    v2 += v2
    v2 = (-(v2))
    fmt.Println(v2, v2)
    var v10 uint8 = 218
    var v11 uint64 = 79
    fmt.Println(v11)
    v12 := GenS2{F0: 1, F1: -918, F2: 0}
    var v13 uint32 = 42
    v14 := GenS2{F0: 0, F1: 2147483647, F2: 100}
    var v15 uint32 = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
