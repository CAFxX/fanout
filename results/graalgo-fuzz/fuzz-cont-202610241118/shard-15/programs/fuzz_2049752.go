package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 bool; F1 uint; F2 int32 }
func genFn0() int16 {
    var v4 uint16 = 1
    v5 := []int{42, 128, 32767}
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = -2.25
    var v3 uint32 = 0
    v6 := GenS1{F0: 42, F1: 42}
    fmt.Println(v2, v1)
    var v7 int32 = -419
    fmt.Println(v2, v7)
    fmt.Println(v7, v2)
    var v8 uint = 18446744073709551615
    if false {
        if false {
            fmt.Println(v2)
            var v9 int16 = 395
            _ = v9
        } else {
            v10 := []int{-357, (v1 << (((v8 << ((v8) & 63))) & 63)), (v1 << (((v8 * 1)) & 63))}
            _ = v10
        }
    }
    switch ("a" + "hello") {
    case "\n\t":
        if true {
            v8 = uint(v3)
            var v11 int = v1
            v12 := GenS1{F0: 42, F1: 18446744073709551615}
            _ = v11
            _ = v12
        }
    case "":
        switch ("" + "a") {
        case "a":
            v13 := GenS2{F0: 4294967295, F1: 61, F2: false}
            _ = v13
        }
    default:
        v14 := GenS3{F0: false, F1: (v8 | v8), F2: v7}
        _ = v14
    }
    v15 := []int{(-(v1))}
    fmt.Println(v3, v1)
    v2 -= v2
    v15 = append(v15, (-(v1)))
    if false {
        var v16 int8 = -47
        v16 += (v16 & 0)
        _ = v16
    }
    fmt.Println(v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
