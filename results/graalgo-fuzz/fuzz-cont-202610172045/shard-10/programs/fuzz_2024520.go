package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 int8, p2 uint32) int {
    p2 -= (p2 & p2)
    fmt.Println(p1, p1, p0)
    var v6 int16 = 847
    switch p0 {
    case 0:
        p2 -= p2
    case 9223372036854775807:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1(p0 uint8) int64 {
    p0 = (p0 - (^(p0)))
    if false {
        p0 = (p0 / ((p0) | 1))
        p0 = (p0 / ((127) | 1))
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 uint = 255
    v4 := make(map[string]uint32)
    var v5 int16 = 32767
    fmt.Println(v3, v5)
    v7 := []int{-301, 9223372036854775807}
    for fi8 := 0; fi8 < 8; fi8++ {
        if true {
            v3 -= v3
            fmt.Println(v5, v3)
            v7 = append(v7, fi8)
        }
        var v9 uint16 = 65535
        v9 = (v9 * v9)
        _ = fi8
        _ = v9
    }
    if (!false) {
        v10 := GenS1{F0: 1, F1: "a", F2: 0}
        _ = v10
    }
    switch "\n\t" {
    case "\n\t":
        switch "a" {
        case "":
            switch -374 {
            case -1:
                fmt.Println(v5, v5)
            }
        case "hello":
            var v11 string = "\n\t"
            _ = v11
        }
    case "foo bar":
        var v12 int32 = -2147483648
        _ = v12
    case "hello":
        var v13 bool = (100 >= 0)
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int32 = -1
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 813}
    switch 42 {
    case -2147483648:
        var v17 int32 = 2147483647
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
