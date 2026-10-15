package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 uint8, p2 uint32) uint8 {
    v7 := []int{199, -669, -709}
    for fi8 := 0; fi8 < 3; fi8++ {
        switch 32767 {
        case 947:
            var v9 int16 = 1000
            _ = v9
        default:
            p0 = (p0 * -47)
        }
        v10 := GenS1{F0: uint(p2), F1: fi8}
        var v11 float32 = (3.14159 * (1e10 * 0.0))
        _ = fi8
        _ = v10
        _ = v11
    }
    if false {
        if false {
            v12 := GenS1{F0: 478, F1: 32767}
            var v13 int64 = -9223372036854775808
            for fi14 := 0; fi14 < 20; fi14++ {
                v7 = append(v7, fi14)
                fmt.Println(p2, p0, p0)
                _ = fi14
            }
            _ = v12
            _ = v13
        }
        var v15 uint = 42
        fmt.Println(p1)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1() int8 {
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS2{F0: "hello"}
        _ = fi16
        _ = v17
    }
    var v18 float64 = 0.0
    v18 += ((float64(3.14159) / float64(-2.25)) + 0.0)
    var v19 int8 = 127
    _ = v18
    _ = v19
    return (v19 + (v19 * v19))
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{}
    var v3 uint = 18446744073709551615
    var v4 int64 = -1
    var v5 int16 = 42
    var v6 uint8 = 160
    if (!true) {
        v6 += (^(v6))
    }
    if ((true && ("" >= "hello")) && true) {
        var v20 string = "a"
        fmt.Println(v1, v5, v6)
        _ = v20
    }
    _ = len(v2)
    v4 = v4
    v6 = (v6 / (((v6 & (v6 + (^(v6))))) | 1))
    var v21 int = 32767
    fmt.Println(v3, v4)
    v2 = append(v2, v21)
    if true {
        switch 591 {
        case 2147483647:
            v22 := GenS1{F0: v3, F1: (v21 * 1)}
            _ = v22
        default:
            if false {
                v1 -= (v1 | (v1 * (v1 | 0)))
            } else {
                v3 = v3
                v23 := GenS1{F0: v3, F1: v21}
                if len(v2) > 0 {
                    v2[0] = (v21 + v21)
                }
                _ = v23
            }
        }
    } else {
        var v24 uint8 = v6
        for fi25 := 0; fi25 < 2; fi25++ {
            if len(v2) > 0 {
                v2[0] = 32768
            }
            switch fi25 {
            case 2147483647:
                fmt.Println(v6, v3)
            case 9223372036854775807:
                var v26 uint32 = 42
                _ = v26
            case 4294967295:
                if len(v2) > 0 {
                    v2[0] = (fi25 - (fi25 - (fi25 * 32767)))
                }
            default:
                var v27 float32 = ((1.5 - 3.14159) - 0.0)
                _ = v27
            }
            _ = fi25
        }
        fmt.Println(v21, v1)
        _ = v24
    }
    var v28 bool = true
    if len(v2) > 0 {
        v2[0] = (v21 >> ((v3) & 63))
    }
    v29 := GenS2{F0: "\n\t"}
    for fi30 := 0; fi30 < 5; fi30++ {
        switch -9223372036854775808 {
        case -536:
            fmt.Println(v6, v6, v3)
        }
        _ = fi30
    }
    var v31 uint8 = (v6 ^ v6)
    var v32 int16 = 42
    v4 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
