package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint16, p1 int) uint8 {
    v5 := []int{(-(p1)), p1}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint16 = (^(p0))
        if len(v5) > 0 {
            v5[0] = (^(fi6))
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 128
}
func genFn1(p0 uint, p1 int8) uint16 {
    if false {
        v8 := []int{128, -1, 9223372036854775807, 2147483647}
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{}
        var v11 uint32 = 483
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = 324
    v13 := []int{}
    p1 += -128
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -1
    var v3 int32 = 42
    v4 := make(map[uint32]uint32)
    v3 += (v3 - -935)
    switch -383 {
    case -9223372036854775808:
        if ((v3 + (v3 >> ((2) & 63))) <= (v3 ^ v3)) {
            v14 := GenS1{F0: ("" + "")}
            fmt.Println(v2, v3)
            v3 = v3
            _ = v14
        }
    case 32768:
        var v15 uint16 = 2
        _ = v15
    default:
        v3 += (v3 >> ((600) & 63))
    }
    var v16 float64 = 0.1
    var v17 bool = (890 >= 597)
    v17 = true
    switch -548 {
    case -2147483648:
        v3 = 42
    default:
        v17 = (-306 == 371)
    }
    fmt.Println(v3)
    v18 := []int{4294967295}
    if v17 {
        fmt.Println(v2)
        if v17 {
            v3 -= (v3 & v3)
        }
    }
    if false {
        var v19 int64 = -9223372036854775808
        v20 := v4[4294967295]
        if v17 {
            v2 += (v2 * v2)
            v19 += -1
            var v21 int8 = 100
            _ = v21
        } else {
            var v22 int8 = 100
            switch "hello" {
            case "":
                v22 = (v22 + int8(v19))
            case "\n\t":
                if (!(-1 != v19)) {
                    fmt.Println(v16, v2)
                }
            default:
                v23 := GenS1{F0: "a"}
                _ = v23
            }
            _ = v22
        }
        _ = v19
        _ = v20
    }
    if v17 {
        v24 := GenS1{F0: ("hello" + ("hello" + ("\n\t" + "\n\t")))}
        fmt.Println(v17, v17)
        v17 = v17
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
