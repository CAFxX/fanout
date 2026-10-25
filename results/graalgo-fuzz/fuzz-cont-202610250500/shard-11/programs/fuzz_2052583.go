package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint16 = 552
    var v2 uint8 = 218
    var v3 bool = true
    var v4 uint64 = 42
    v5 := make(map[uint32]int)
    var v6 int8 = 42
    var v7 uint = 2
    if (!false) {
        v4 = v4
        var v8 uint32 = 0
        var v9 float64 = 1.5
        _ = v8
        _ = v9
    } else {
        var v10 uint = uint(v1)
        _ = v10
    }
    switch "\n\t" {
    case "":
        fmt.Println(v4, v1, v7)
    case "hello":
        v1 = uint16(v2)
    }
    var v11 int16 = 32767
    v12 := GenS1{F0: -9223372036854775808}
    var v13 int16 = v11
    var v14 int64 = -316
    v15 := GenS2{F0: (v13 | v13), F1: 2147483648}
    var v16 float32 = float32(-2.25)
    v17 := GenS2{F0: 32767, F1: 532}
    v18 := []int{128, 32768}
    var v19 uint = v7
    fmt.Println(v4, v6)
    v13 += 1000
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
