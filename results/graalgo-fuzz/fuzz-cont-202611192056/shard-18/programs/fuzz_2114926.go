package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint8 }
func genFn0() uint {
    var v4 float32 = 0.0
    var v5 uint = 1
    _ = v4
    _ = v5
    return (v5 | v5)
}
func genFn1() string {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 58}
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 1}
            switch "\n\t" {
            case "foo bar":
                var v10 int64 = 1
                _ = v10
            case "\n\t":
                v11 := []int{241, fi6}
                _ = v11
            default:
                fmt.Println(fi8, fi6)
            }
            var v12 float64 = 0.1
            v13 := []int{(fi8 & 4294967295), (fi8 + 1), fi8}
            _ = fi8
            _ = v12
            _ = v13
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if true {
        var v14 int8 = -128
        var v15 int32 = -690
        _ = v14
        _ = v15
    }
    switch "" {
    case "foo bar":
        var v16 int32 = 724
        _ = v16
    case "hello":
        var v17 uint = 451
        _ = v17
    case "":
        var v18 float64 = 1e10
        _ = v18
    }
    return ("hello" + "")
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -54
    v3 := []uint32{946, 2}
    v19 := GenS2{F0: 475}
    var v20 uint = 341
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 int16 = 42
    var v22 uint32 = 1
    v23 := GenS3{F0: -1, F1: 18446744073709551615}
    v21 = v21
    var v24 float32 = 1.5
    _ = len(v3)
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 bool = true
        fmt.Println(v26)
        v2 += (v2 + 0)
        _ = fi25
        _ = v26
    }
    v27 := GenS4{F0: 32767, F1: 42}
    v2 -= (-(v2))
    v28 := GenS5{F0: 127, F1: 69, F2: 2}
    _ = len(v1)
    v2 += 838
    v29 := GenS5{F0: 127, F1: 1, F2: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
