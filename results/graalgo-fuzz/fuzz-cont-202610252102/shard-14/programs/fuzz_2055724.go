package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]int)
    var v3 uint64 = 2
    var v4 int16 = 1
    v3 = (v3 ^ 2)
    v5 := GenS1{F0: 1, F1: 42, F2: -268}
    v4 -= v4
    var v6 uint64 = v3
    var v7 int = 32767
    if (!((true && (-167 > 537)) && true)) {
        v8 := GenS1{F0: 164, F1: 1, F2: (v7 % (((v7 + v7)) | 1))}
        fmt.Println(v7)
        v7 += 2147483647
        _ = v8
    }
    fmt.Println(v4, v1)
    var v9 int8 = 1
    v6 += (v3 | (v6 + v6))
    v1 += -2.25
    switch (v7 << ((4294967295) & 63)) {
    case -2147483648:
        v10 := GenS2{F0: v9, F1: 721}
        _ = v10
    case 32767:
        v6 = 18446744073709551615
    case 63:
        v7 = v7
    default:
        v11 := GenS3{F0: 94}
        _ = v11
    }
    var v12 uint = 0
    fmt.Println(v1, v9)
    var v13 bool = false
    if false {
        var v14 bool = (!(!v13))
        var v15 bool = (v9 > 0)
        v6 = (^(v3))
        _ = v14
        _ = v15
    }
    var v16 string = ("" + ("a" + ("foo bar" + "a")))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
