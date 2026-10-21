package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0() uint32 {
    if false {
        var v5 uint = 1
        _ = v5
    }
    var v6 int64 = 42
    v7 := GenS1{F0: 18446744073709551615, F1: -2147483648}
    var v8 int8 = -41
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint16 = 1
    var v3 uint8 = 1
    var v4 int8 = 75
    v9 := []int{128, 557, 1}
    v10 := GenS2{F0: 2}
    var v11 int = 9223372036854775807
    var v12 int = v11
    fmt.Println(v12)
    switch (v12 + 127) {
    case 2147483647:
        v9 = append(v9, (-(v12)))
    default:
        v13 := GenS3{F0: 1, F1: -9223372036854775808, F2: 0}
        _ = v13
    }
    var v14 bool = ((v4 / (((-(v4))) | 1)) == 127)
    if len(v9) > 0 {
        v9[0] = (-(v12))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
