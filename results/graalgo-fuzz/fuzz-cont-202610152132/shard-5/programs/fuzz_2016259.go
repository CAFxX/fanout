package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 int64, p1 uint, p2 uint16) uint32 {
    p0 += p0
    v4 := GenS1{F0: -905, F1: 42}
    var v5 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int8, p1 uint32, p2 bool) int64 {
    var v6 int8 = (p0 + (p0 / (((-(p0))) | 1)))
    v7 := GenS1{F0: -325, F1: 127}
    if ("a" <= "foo bar") {
        v8 := []int{-607, -460, -152}
        p1 = (p1 & genFn0(-9223372036854775808, 2, 93))
        _ = v8
    } else {
        p1 -= 42
        switch "hello" {
        case "":
            var v9 int = -9223372036854775808
            _ = v9
        case "a":
            if true {
                v10 := GenS1{F0: 9223372036854775807, F1: 127}
                var v11 int = 4294967295
                _ = v10
                _ = v11
            }
        case "\n\t":
            v12 := GenS2{F0: v6, F1: (("" + "hello") + ("" + "")), F2: 0}
            _ = v12
        }
        var v13 uint16 = 964
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v14 int = -2147483648
    var v15 int = v14
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 float32 = float32(1.5)
        var v18 uint32 = 4294967295
        v19 := GenS3{F0: "hello"}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS4{F0: 680, F1: 67}
    delete(v2, 42)
    v14 = v15
    v21 := GenS5{F0: 494}
    v14 -= (-(v14))
    v22 := []int{}
    switch v14 {
    case 4294967295:
        v15 = (v14 >> ((42) & 63))
    }
    v23 := GenS6{F0: 915, F1: (v14 << ((778) & 63)), F2: 267}
    for fi24 := 0; fi24 < 16; fi24++ {
        if false {
            var v25 int = v15
            switch -1 {
            case 561:
                fmt.Println(v14, v15)
            case -818:
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 float64 = (0.1 / 1e10)
                    _ = fi26
                    _ = v27
                }
            }
            _ = v25
        }
        v15 = 32767
        var v28 uint = 42
        _ = fi24
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
