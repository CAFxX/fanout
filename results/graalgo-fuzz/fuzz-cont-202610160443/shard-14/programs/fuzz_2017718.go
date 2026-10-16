package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 int; F2 string }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 uint32; F1 int64; F2 int16 }
type GenS5 struct { F0 uint; F1 int }
func genFn0(p0 uint64, p1 bool, p2 int16) int32 {
    var v6 int32 = 1
    v7 := GenS1{F0: 694, F1: "\n\t", F2: v6}
    v8 := GenS2{F0: 42, F1: -2147483648, F2: "hello"}
    v9 := []int{4294967295, 1, 746}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v6 >> ((4294967295) & 63))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 712
    var v3 float32 = 0.0
    v4 := make(map[uint32]int)
    var v5 int64 = 42
    v10 := []int{}
    var v11 uint64 = 0
    v12 := GenS1{F0: (v2 | uint32(v11)), F1: ("" + "a"), F2: genFn0(18446744073709551615, true, 275)}
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int16 = 1
        _ = fi13
        _ = v14
    }
    if (859 < 42) {
        switch 127 {
        case 32767:
            if false {
                var v15 uint = 255
                fmt.Println(v2, v2, v11)
                _ = v15
            }
        case 42:
            var v16 uint = 255
            _ = v16
        }
    }
    v17 := GenS3{F0: 127, F1: v5, F2: false}
    var v18 bool = false
    v3 -= (v3 * v3)
    v19 := GenS1{F0: (v2 / ((v2) | 1)), F1: "\n\t", F2: 0}
    switch -9223372036854775808 {
    case 128:
        v2 = v2
    case 0:
        v10 = append(v10, -323)
    }
    switch 2147483647 {
    case 1:
        v11 += v11
    case 4294967295:
        v20 := GenS4{F0: v2, F1: v5, F2: -346}
        _ = v20
    case 32768:
        var v21 uint16 = 255
        _ = v21
    }
    v2 += (v2 + (v2 + (v2 >> ((0) & 63))))
    delete(v4, v2)
    for fi22 := 0; fi22 < 12; fi22++ {
        v10 = append(v10, (fi22 - 4294967295))
        _ = fi22
    }
    switch "foo bar" {
    case "hello":
        var v23 bool = v18
        _ = v23
    case "\n\t":
        v24 := GenS5{F0: 255, F1: 0}
        _ = v24
    case "a":
        var v25 int16 = 1
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
