package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int32 = 0
    var v2 bool = ((v1 >= (v1 & -1)) && (42 != 0))
    var v3 int32 = v1
    v4 := []int{-2147483648, -186, 2147483647}
    var v5 string = "foo bar"
    v6 := make(map[int]string)
    var v7 uint32 = 4294967295
    v7 = (v7 * v7)
    switch -595 {
    case -768:
        var v8 int16 = 452
        _ = v8
    case 361:
        v1 += v3
    case 9223372036854775807:
        fmt.Println(v2, v5)
    }
    var v9 float32 = 3.14159
    var v10 int16 = 99
    v11 := GenS1{F0: true, F1: 65535}
    fmt.Println(v9, v10)
    if ((!(907 != 0)) && v2) {
        v1 = (-(v1))
    } else {
        if v2 {
            v12 := GenS2{F0: -59}
            _ = v12
        } else {
            v13 := GenS3{F0: 2}
            v14 := GenS1{F0: (v7 <= (^(v7))), F1: 0}
            _ = v13
            _ = v14
        }
        if false {
            var v15 int64 = -649
            var v16 string = v5
            fmt.Println(v9)
            _ = v15
            _ = v16
        }
        var v17 uint8 = 255
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
