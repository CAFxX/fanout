package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    switch 465 {
    case 791:
        var v4 uint64 = 247
        _ = v4
    case -9223372036854775808:
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 uint64 = 1
            var v7 uint64 = v6
            v6 -= (v7 | (v6 % ((42) | 1)))
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 17; fi8++ {
            if ((("\n\t" + "") > "") && false) {
                v9 := GenS1{F0: 0, F1: 1}
                v10 := GenS1{F0: 985, F1: 2}
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    }
    delete(v3, 4294967295)
    var v11 uint8 = 42
    switch -822 {
    case 2147483647:
        v12 := []int{-244, 127, 1}
        _ = v12
    default:
        fmt.Println(v11)
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 float32 = float32(0.1)
            _ = fi14
            _ = v15
        }
        v11 -= v11
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int8 = 127
            var v18 float32 = (-2.25 / float32(3.14159))
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    v11 = v11
    fmt.Println(v11)
    v19 := []int{1}
    var v20 float32 = 1.5
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS1{F0: 42, F1: 32767}
        for fi23 := 0; fi23 < 7; fi23++ {
            _ = len(v19)
            v24 := v1[("a" + "hello")]
            _ = fi23
            _ = v24
        }
        var v25 uint = 569
        _ = fi21
        _ = v22
        _ = v25
    }
    switch "hello" {
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            var v26 int8 = 19
            _ = v26
        case "":
            v20 += float32((-2.25 + 0.0))
        case "hello":
            v11 = (v11 / (((v11 >> ((1) & 63))) | 1))
        }
    default:
        for fi27 := 0; fi27 < 14; fi27++ {
            _ = len(v2)
            v20 = 1e10
            fmt.Println(v11, v11, fi27)
            _ = fi27
        }
    }
    fmt.Println(v11, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
