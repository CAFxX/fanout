package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 uint8, p1 uint) bool {
    var v4 uint = 2
    var v5 uint16 = 32767
    var v6 uint8 = p0
    var v7 bool = true
    fmt.Println(v6, v5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (true && ((3.14159 <= 1.5) && (1 == 12)))
}
func genFn1() uint16 {
    v8 := []int{9223372036854775807}
    switch "\n\t" {
    case "foo bar":
        var v9 int16 = 1
        _ = v9
    case "a":
        var v10 int32 = 42
        _ = v10
    }
    var v11 uint64 = 2
    var v12 int32 = 1
    var v13 int16 = 42
    _ = v11
    _ = v12
    _ = v13
    _ = v8
    return 2
}

func main() {
    var v1 uint16 = 983
    v2 := make(map[uint32]uint32)
    v3 := []uint32{374, 65535}
    switch 0 {
    case 4294967295:
        fmt.Println(v1)
    default:
        _ = len(v3)
    }
    fmt.Println(v1)
    delete(v2, 42)
    v14 := GenS1{F0: true}
    v1 = v1
    var v15 int8 = 1
    switch 42 {
    case 32768:
        v16 := GenS2{F0: "foo bar", F1: 259}
        _ = v16
    case 42:
        v15 = (v15 ^ v15)
    case -2147483648:
        var v17 int8 = (v15 / ((v15) | 1))
        _ = v17
    default:
        fmt.Println(v15)
    }
    if true {
        v1 = 972
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
