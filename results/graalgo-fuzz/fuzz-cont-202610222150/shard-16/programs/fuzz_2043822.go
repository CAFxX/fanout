package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }
func genFn0(p0 int8, p1 uint8, p2 uint8) int {
    p1 = 127
    fmt.Println(p2)
    var v6 uint32 = 156
    var v7 int8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 398
}
func genFn1(p0 int16, p1 int64) uint32 {
    switch "a" {
    case "":
        var v8 int16 = 1
        _ = v8
    default:
        p0 = (^(p0))
    }
    var v9 string = "a"
    var v10 uint32 = 1
    v11 := []int{4294967295, 36}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = -36
    var v2 int32 = 42
    var v3 uint8 = 2
    var v4 int64 = -9223372036854775808
    v5 := []string{}
    var v12 uint = 18446744073709551615
    if true {
        v3 = v3
    }
    var v13 int16 = 1000
    v5 = append(v5, ((("hello" + "") + "hello") + ""))
    v1 = 22
    fmt.Println(v13)
    fmt.Println(v3, v13)
    v5 = append(v5, "\n\t")
    var v14 uint = (v12 & (v12 ^ v12))
    var v15 string = "a"
    if false {
        v15 = v15
        var v16 uint32 = uint32(v3)
        _ = v16
    } else {
        v12 = (^(v14))
        v17 := GenS1{F0: -9223372036854775808, F1: (v4 * (v4 % ((v4) | 1))), F2: 2}
        var v18 int16 = 104
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
