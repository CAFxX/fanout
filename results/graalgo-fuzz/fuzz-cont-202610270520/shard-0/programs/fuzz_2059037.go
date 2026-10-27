package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
func genFn0(p0 int8, p1 int) uint {
    p1 -= (p1 / ((p1) | 1))
    fmt.Println(p1)
    var v4 int16 = -346
    _ = p0
    _ = p1
    _ = v4
    return 255
}
func genFn1(p0 int32, p1 uint64, p2 uint16) int {
    switch "\n\t" {
    case "foo bar":
        p2 -= p2
    case "\n\t":
        p2 = p2
    }
    fmt.Println(p1, p2)
    for fi5 := 0; fi5 < 11; fi5++ {
        p1 += p1
        if (42 <= genFn0(127, fi5)) {
            p2 -= (^(p2))
        } else {
            fmt.Println(fi5)
            v6 := GenS1{F0: 42, F1: "a", F2: 2}
            v7 := []int{-961, (^(fi5)), fi5}
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 0
    v3 := make(map[uint32]int)
    v3[1] = -122
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: -875, F1: "hello", F2: 1}
        v10 := v1[(-(fi8))]
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v2 = -2147483648
        _ = fi11
    }
    v12 := GenS1{F0: 42, F1: "a", F2: 88}
    var v13 uint32 = 2147483647
    v2 = v2
    v13 += v13
    var v14 float32 = 3.14159
    switch "" {
    case "foo bar":
        var v15 uint = 0
        _ = v15
    case "":
        v16 := GenS1{F0: 1, F1: "", F2: 2}
        _ = v16
    default:
        var v17 int32 = (^(v2))
        _ = v17
    }
    var v18 int = genFn1(v2, 1, 65535)
    if false {
        var v19 int64 = 42
        _ = v19
    }
    fmt.Println(v14, v14)
    for fi20 := 0; fi20 < 19; fi20++ {
        delete(v1, fi20)
        v18 = (fi20 | (fi20 * 127))
        _ = fi20
    }
    v14 = 1.5
    if true {
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS2{F0: 255, F1: "\n\t", F2: 116}
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
