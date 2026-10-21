package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 bool) uint {
    var v5 uint16 = 255
    if (0 >= 4294967295) {
        v6 := GenS1{F0: 257, F1: -1, F2: 4294967295}
        _ = v6
    }
    v7 := GenS2{F0: 17}
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 1
}
func genFn1(p0 uint8) uint64 {
    v8 := GenS1{F0: -421, F1: -1, F2: 42}
    p0 += p0
    var v9 bool = true
    fmt.Println(p0, v9)
    var v10 string = "hello"
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 36
}

func main() {
    v1 := []int{127, -9223372036854775808, 4294967295, 2147483647}
    v2 := make(map[string]uint32)
    var v3 uint16 = 0
    var v4 int16 = -714
    v11 := GenS2{F0: 42}
    switch -2147483648 {
    case 2147483647:
        var v12 uint = 170
        _ = v12
    case 9223372036854775807:
        var v13 bool = (15 >= 31)
        _ = v13
    }
    delete(v2, (("\n\t" + "a") + "hello"))
    v14 := GenS2{F0: 42}
    var v15 uint64 = 944
    var v16 int64 = -1
    switch "" {
    case "\n\t":
        v16 = (v16 & (-(v16)))
    default:
        v3 = v3
    }
    fmt.Println(v4, v16, v4)
    var v17 float32 = 3.14159
    v4 = v4
    v17 -= float32(-2.25)
    v18 := GenS2{F0: 100}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
