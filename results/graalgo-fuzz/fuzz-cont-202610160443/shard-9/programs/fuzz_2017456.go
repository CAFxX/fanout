package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
func genFn0(p0 uint64) uint8 {
    v7 := GenS1{F0: p0, F1: 1}
    fmt.Println(p0)
    p0 += 0
    p0 = 730
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 int64, p2 int32) int8 {
    fmt.Println(p0)
    v8 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    var v1 uint8 = 186
    var v2 int8 = 127
    v3 := []int{127}
    var v4 uint32 = 2
    var v5 uint16 = 1
    var v6 float64 = (0.1 + 0.0)
    switch "\n\t" {
    case "hello":
        if (v2 == genFn1((v1 - v1), 9223372036854775807, 1)) {
            v9 := GenS1{F0: 774, F1: 0}
            v5 = v5
            v6 += 0.1
            _ = v9
        } else {
            var v10 uint = 2
            v11 := GenS1{F0: 4294967295, F1: 304}
            _ = v10
            _ = v11
        }
    }
    var v12 uint = 42
    fmt.Println(v1, v2)
    v5 = (^(v5))
    var v13 uint = (^(v12))
    _ = len(v3)
    switch ("\n\t" + "hello") {
    case "hello":
        for fi14 := 0; fi14 < 6; fi14++ {
            v13 -= v12
            _ = fi14
        }
    case "":
        if true {
            var v15 int8 = v2
            _ = v15
        } else {
            v2 -= v2
        }
    case "foo bar":
        v16 := GenS1{F0: 146, F1: 449}
        _ = v16
    }
    var v17 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
