package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32; F2 int }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int; F1 uint8; F2 int16 }
func genFn0() uint {
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: ("foo bar" + (("hello" + ("" + "\n\t")) + "\n\t"))}
            v10 := GenS1{F0: ""}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: ("a" + "a")}
        _ = fi7
        _ = v11
    }
    v12 := []int{-9223372036854775808, 774}
    var v13 uint16 = 79
    v13 -= v13
    _ = v12
    _ = v13
    return 2
}
func genFn1(p0 uint, p1 uint8, p2 string) int32 {
    p2 = p2
    p0 = p0
    v14 := []int{128, 4294967295, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 892
}

func main() {
    v1 := []int{4294967295}
    var v2 int32 = 1
    v3 := []uint32{910, 2}
    var v4 int16 = 0
    v5 := []int{1, 227, 2147483647, 1}
    v6 := []int{}
    var v15 int64 = -516
    v15 -= 0
    v16 := GenS1{F0: ("foo bar" + ("\n\t" + "a"))}
    for fi17 := 0; fi17 < 9; fi17++ {
        _ = len(v3)
        var v18 uint64 = 4294967295
        if true {
            v2 -= -2147483648
            v15 += (^(v15))
            if len(v1) > 0 {
                v1[0] = fi17
            }
        }
        _ = fi17
        _ = v18
    }
    var v19 float64 = 0.1
    v20 := GenS2{F0: v4, F1: v2, F2: 2147483647}
    var v21 uint32 = 459
    _ = len(v3)
    v22 := GenS2{F0: v4, F1: v2, F2: 32767}
    v23 := GenS3{F0: -1, F1: "foo bar"}
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 int16 = (v4 << ((255) & 63))
        _ = fi24
        _ = v25
    }
    var v26 uint32 = 133
    var v27 int = 42
    v21 = v21
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int32 = (v2 * genFn1(genFn0(), 1, "hello"))
        fmt.Println(v27)
        v29 += (v29 >> ((18446744073709551615) & 63))
        _ = fi28
        _ = v29
    }
    if (18446744073709551615 != 359) {
        v30 := GenS3{F0: 2147483647, F1: "hello"}
        switch "a" {
        case "foo bar":
            for fi31 := 0; fi31 < 5; fi31++ {
                v19 += v19
                fmt.Println(v19)
                _ = fi31
            }
        default:
            v32 := GenS4{F0: v27, F1: 1, F2: v4}
            _ = v32
        }
        fmt.Println(v19, v15, v15)
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
