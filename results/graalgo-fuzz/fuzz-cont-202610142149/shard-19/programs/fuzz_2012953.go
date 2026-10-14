package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
func genFn0() int32 {
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int32 = 1
            for fi8 := 0; fi8 < 4; fi8++ {
                switch "hello" {
                case "foo bar":
                    v9 := GenS1{F0: fi6, F1: 32767, F2: fi6}
                    _ = v9
                case "hello":
                    if (!true) {
                        var v10 string = ("hello" + "a")
                        _ = v10
                    }
                case "\n\t":
                    var v11 int64 = 651
                    _ = v11
                default:
                    var v12 float32 = 1.5
                    _ = v12
                }
                v13 := GenS1{F0: 9223372036854775807, F1: 42, F2: fi8}
                _ = fi8
                _ = v13
            }
            _ = fi6
            _ = v7
        }
        var v14 uint64 = 221
        fmt.Println(v14)
        _ = v14
    }
    var v15 bool = false
    v15 = v15
    v15 = (3.14159 == 0.0)
    v16 := GenS1{F0: 1, F1: 1, F2: 32767}
    _ = v15
    _ = v16
    return 327
}
func genFn1() uint16 {
    v17 := GenS2{F0: 1, F1: -923, F2: (!true)}
    var v18 int = 282
    var v19 int64 = -9223372036854775808
    _ = v17
    _ = v18
    _ = v19
    return 892
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 203
    v3 := []uint32{610}
    v4 := []int{}
    var v5 int = -2147483648
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint = 450
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 int8 = 2
        var v24 int32 = genFn0()
        _ = fi22
        _ = v23
        _ = v24
    }
    v5 -= (v5 ^ v5)
    fmt.Println(v5)
    v25 := GenS1{F0: v5, F1: 42, F2: (v5 - int(v2))}
    var v26 float32 = 3.14159
    v2 += v2
    var v27 int64 = v2
    var v28 int = v5
    fmt.Println(v5, v26)
    if len(v4) > 0 {
        v4[0] = (v28 << ((375) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
