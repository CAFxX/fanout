package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint32; F2 int64 }
func genFn0() uint8 {
    var v4 uint8 = 42
    var v5 uint = 0
    v5 = (^(v5))
    _ = v4
    _ = v5
    return (v4 - 42)
}
func genFn1(p0 uint8, p1 string, p2 int64) uint64 {
    v6 := []int{-2147483648}
    var v7 int8 = -58
    var v8 uint = 1
    var v9 uint32 = 566
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 802
}

func main() {
    v1 := []uint32{}
    v2 := []int{}
    v3 := make(map[int]bool)
    if (534 >= 611) {
        var v10 uint16 = 118
        _ = v10
    } else {
        var v11 bool = true
        v11 = v11
        _ = v11
    }
    v12 := GenS1{F0: "foo bar", F1: 4294967295}
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS2{F0: -2147483648, F1: 42, F2: 100}
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 2
    _ = len(v3)
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS3{F0: 2}
            _ = fi17
            _ = v18
        }
        var v19 int32 = 42
        _ = len(v1)
        _ = fi16
        _ = v19
    }
    v20 := GenS4{F0: 118, F1: 0, F2: 42}
    fmt.Println(v15)
    fmt.Println(v15)
    for fi21 := 0; fi21 < 1; fi21++ {
        fmt.Println(v15, fi21)
        _ = fi21
    }
    switch -872 {
    case 4294967295:
        var v22 string = "foo bar"
        _ = v22
    case 32768:
        var v23 int8 = -43
        _ = v23
    case -446:
        v15 = v15
    }
    if false {
        if (645 == -9223372036854775808) {
            v15 -= (v15 / ((v15) | 1))
        } else {
            v15 -= 42
            v15 -= (v15 * (v15 - (v15 - genFn1(2, "\n\t", 9223372036854775807))))
            v24 := v3[-1]
            _ = v24
        }
        var v25 uint16 = 0
        v25 -= v25
        _ = v25
    } else {
        switch 32767 {
        case -660:
            for fi26 := 0; fi26 < 5; fi26++ {
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 int64 = 1
                    if len(v2) > 0 {
                        v2[0] = fi26
                    }
                    v15 += (v15 >> ((618) & 63))
                    _ = fi27
                    _ = v28
                }
                var v29 bool = false
                _ = fi26
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 9; fi30++ {
            if true {
                _ = len(v3)
                fmt.Println(fi30)
            }
            v31 := GenS2{F0: 2147483647, F1: 42, F2: 127}
            _ = fi30
            _ = v31
        }
        delete(v3, 2147483647)
    }
    var v32 uint64 = (v15 * genFn1(92, "", -426))
    var v33 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
}
