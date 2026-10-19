package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int32; F1 string; F2 uint8 }
func genFn0(p0 bool) int64 {
    var v4 int32 = 0
    v5 := GenS1{F0: 494, F1: 42, F2: 42}
    fmt.Println(v4)
    switch 32767 {
    case 2147483647:
        v6 := GenS1{F0: (v4 | (v4 + 42)), F1: 89, F2: 1000}
        _ = v6
    default:
        var v7 string = "a"
        _ = v7
    }
    switch "foo bar" {
    case "hello":
        var v8 int8 = -91
        _ = v8
    case "":
        var v9 uint32 = 992
        _ = v9
    case "foo bar":
        var v10 uint32 = 0
        _ = v10
    }
    _ = p0
    _ = v4
    _ = v5
    return -675
}
func genFn1(p0 uint8, p1 uint16, p2 uint8) int64 {
    var v11 int16 = 1
    for fi12 := 0; fi12 < 4; fi12++ {
        v11 += 1000
        if true {
            fmt.Println(fi12, p0, v11)
        }
        _ = fi12
    }
    v13 := []int{}
    v13 = append(v13, -467)
    v14 := GenS2{F0: 2147483648, F1: -694}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    _ = v14
    return -776
}

func main() {
    var v1 bool = (!false)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v15 uint = 255
    var v16 int16 = 1000
    var v17 uint8 = 255
    var v18 int32 = -1
    v19 := GenS3{F0: 2147483647, F1: 2}
    v2["hello"] = 128
    fmt.Println(v1)
    v20 := GenS4{F0: (v18 ^ 42), F1: "foo bar", F2: (v17 >> (((v15 / (((v15 << ((v15) & 63))) | 1))) & 63))}
    if (2147483648 < 731) {
        var v21 uint8 = v17
        var v22 uint32 = 0
        _ = v21
        _ = v22
    }
    v15 += (v15 | (^(v15)))
    v23 := []int{127, 810, 100, -1}
    v23 = append(v23, 4294967295)
    v16 -= v16
    delete(v3, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
