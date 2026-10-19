package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int8; F1 int32; F2 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 uint16, p2 uint32) int32 {
    v5 := GenS1{F0: 100}
    p2 = 42
    if false {
        if true {
            var v6 int64 = -1
            _ = v6
        } else {
            fmt.Println(p0, p2, p1)
            var v7 uint = (p0 / ((p0) | 1))
            _ = v7
        }
        if (!false) {
            p1 += (p1 * (p1 + (p1 << (((p0 + 2)) & 63))))
        }
        v8 := GenS2{F0: false, F1: 0, F2: 2147483647}
        _ = v8
    }
    p2 -= 4294967295
    p2 -= (p2 - (p2 - (p2 >> ((18446744073709551615) & 63))))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 int) int8 {
    fmt.Println(p0)
    switch "foo bar" {
    case "foo bar":
        v9 := GenS3{F0: 0, F1: 66, F2: 2}
        _ = v9
    default:
        v10 := GenS1{F0: -1}
        _ = v10
    }
    v11 := []int{-1, (p0 + 453), p0, (^(p0))}
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 uint = 631
    var v2 uint8 = 2
    var v3 int = 42
    v4 := make(map[uint32]uint32)
    fmt.Println(v1)
    delete(v4, 0)
    var v12 uint32 = 42
    v3 += 32768
    switch 588 {
    case -358:
        var v13 uint8 = (v2 / ((v2) | 1))
        _ = v13
    case 32767:
        v14 := GenS1{F0: 1}
        _ = v14
    }
    v15 := []int{-960, -994}
    if false {
        v2 = v2
    } else {
        v16 := GenS4{F0: genFn1(v3)}
        _ = v16
    }
    v12 = 2
    if len(v15) > 0 {
        v15[0] = v3
    }
    var v17 uint32 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
