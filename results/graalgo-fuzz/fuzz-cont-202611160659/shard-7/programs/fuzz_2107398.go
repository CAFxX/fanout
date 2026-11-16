package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }
func genFn0() uint32 {
    if (4294967295 == 18446744073709551615) {
        var v7 bool = (1 <= 2)
        switch ("foo bar" + "hello") {
        case "":
            v7 = (-83 > 0)
        case "hello":
            v8 := GenS1{F0: (42 == -71), F1: -870, F2: 0}
            _ = v8
        }
        _ = v7
    }
    switch 0 {
    case 32768:
        var v9 int64 = -411
        _ = v9
    case -1:
        var v10 int32 = -694
        _ = v10
    default:
        v11 := []int{128, 32768, 2147483647}
        _ = v11
    }
    v12 := []int{-539, 42, -976}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    switch "a" {
    case "a":
        var v13 uint = 0
        _ = v13
    }
    _ = v12
    return 854
}
func genFn1() uint32 {
    if (1.5 < (-2.25 * (1e10 + (1.5 / 0.1)))) {
        var v14 uint16 = 255
        _ = v14
    }
    var v15 int16 = -55
    var v16 int64 = -596
    _ = v15
    _ = v16
    return 490
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    v4 := make(map[string]bool)
    var v5 uint32 = 812
    var v6 uint32 = (v5 ^ 2)
    fmt.Println(v5, v6)
    if (8 >= 1) {
        v17 := GenS2{F0: -1, F1: -1, F2: 1}
        _ = v17
    }
    var v18 int64 = 1
    v19 := GenS1{F0: true, F1: 42, F2: 179}
    v5 = 2147483648
    v20 := GenS1{F0: false, F1: (v18 >> ((2) & 63)), F2: 317}
    v6 += genFn1()
    var v21 uint64 = 18446744073709551615
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
