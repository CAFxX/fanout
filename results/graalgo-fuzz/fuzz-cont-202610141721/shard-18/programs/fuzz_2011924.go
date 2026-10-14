package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 int32, p1 uint32, p2 uint16) int32 {
    if false {
        v6 := []int{-883, -1}
        var v7 int32 = (p0 + -2147483648)
        p0 = (v7 | (v7 ^ v7))
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (p1 ^ (p1 / ((2147483647) | 1))), F1: 255}
    v9 := GenS2{F0: -174, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p0
}
func genFn1() uint16 {
    var v10 int = 128
    v10 -= (v10 - (v10 & v10))
    _ = v10
    return 670
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 1
    v3 := []int{9223372036854775807}
    var v4 uint64 = 0
    v5 := []string{}
    fmt.Println(v4, v2)
    var v11 uint8 = 128
    if ((v4 - (v4 ^ (v4 / (((v4 ^ v4)) | 1)))) == v4) {
        var v12 int16 = 1
        v11 = (v11 & v11)
        _ = v12
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v2, v2, v2)
    case "":
        fmt.Println(v11, v11)
    default:
        switch 1 {
        case -448:
            var v13 uint = 0
            _ = v13
        }
    }
    v4 = v4
    v11 = 255
    v2 -= v2
    fmt.Println(v2, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
