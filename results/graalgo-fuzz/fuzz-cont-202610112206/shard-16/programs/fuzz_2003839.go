package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 string; F2 bool }
func genFn0(p0 uint16, p1 uint64) int {
    v7 := GenS1{F0: (false && false), F1: 341, F2: "a"}
    v8 := GenS1{F0: false, F1: 0, F2: ""}
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: true, F1: 2, F2: "a"}
        p1 = (p1 & (^(p1)))
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32768
}
func genFn1(p0 int64, p1 int8, p2 int16) uint32 {
    var v11 int = 32767
    v12 := GenS2{F0: -128}
    var v13 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = 100
    var v3 int8 = (v2 % ((v2) | 1))
    var v4 uint64 = 0
    var v5 float64 = (3.14159 / 1e10)
    var v6 int16 = 1000
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 uint = 255
        v16 := []int{(fi14 / (((fi14 + (fi14 >> ((1) & 63)))) | 1)), fi14}
        var v17 uint64 = (v4 % ((v4) | 1))
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    switch -881 {
    case 2147483647:
        var v18 uint32 = genFn1(9223372036854775807, 42, v6)
        _ = v18
    case 128:
        if ((0.1 == (((-2.25 / v5) + 3.14159) * v5)) && (genFn1(0, v3, (-(v6))) > 65535)) {
            v19 := []int{-2147483648}
            if len(v19) > 0 {
                v19[0] = 42
            }
            _ = v19
        }
    }
    v20 := []int{}
    switch ("a" + "") {
    case "":
        var v21 float64 = (0.0 * float64(v1))
        _ = v21
    case "foo bar":
        if (v2 >= -1) {
            v22 := GenS2{F0: 100}
            var v23 uint8 = 2
            v20 = append(v20, 1)
            _ = v22
            _ = v23
        }
    default:
        var v24 float64 = 1e10
        _ = v24
    }
    v5 -= 1e10
    v2 = 100
    var v25 int = -9223372036854775808
    for fi26 := 0; fi26 < 12; fi26++ {
        if false {
            var v27 int64 = 9223372036854775807
            _ = v27
        }
        _ = fi26
    }
    var v28 int8 = (v3 << ((255) & 63))
    if false {
        fmt.Println(v28, v25)
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    var v29 uint64 = v4
    v30 := GenS3{F0: v25, F1: "hello", F2: true}
    if true {
        for fi31 := 0; fi31 < 2; fi31++ {
            if false {
                for fi32 := 0; fi32 < 19; fi32++ {
                    _ = len(v20)
                    v29 = 645
                    _ = fi32
                }
                v3 += v2
            }
            v2 = v3
            _ = fi31
        }
    } else {
        v29 = 2
        v33 := GenS2{F0: v2}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
