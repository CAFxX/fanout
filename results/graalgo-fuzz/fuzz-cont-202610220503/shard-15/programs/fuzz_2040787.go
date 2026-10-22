package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 uint, p1 uint16, p2 int32) int16 {
    p0 = (p0 << (((p0 % (((^(p0))) | 1))) & 63))
    var v7 uint = p0
    fmt.Println(p0, p1, p0)
    v8 := GenS1{F0: (0.1 <= float32((-2.25 / ((((0.0 / 0.1) - -2.25) - (float64(1.5) - -2.25)) + (((float64(0.0) * 0.0) * 0.0) / float64(-2.25)))))), F1: v7}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint8) int {
    var v9 int16 = genFn0(4294967295, 335, -141)
    v10 := []int{0}
    v9 = v9
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int8 = -1
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: (uint32(p0) == 723), F1: 0}
    _ = p0
    _ = v10
    _ = v13
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 1
    var v3 uint32 = 65535
    v4 := []int{}
    var v5 bool = true
    v6 := []uint32{v3, (v3 + v3), 65535, v3}
    var v14 int64 = (v2 - (v2 >> ((0) & 63)))
    var v15 float32 = 0.1
    if len(v4) > 0 {
        v4[0] = -229
    }
    var v16 uint8 = 1
    fmt.Println(v2, v14, v5)
    if (!(!((42 == 879) || (128 > v16)))) {
        v1[("hello" + "a")] = (v3 * v3)
    }
    var v17 uint32 = v3
    if len(v6) > 0 {
        v6[0] = 42
    }
    switch "\n\t" {
    case "foo bar":
        v18 := GenS1{F0: false, F1: 724}
        _ = v18
    default:
        var v19 int8 = 0
        _ = v19
    }
    if v5 {
        if v5 {
            var v20 uint = 42
            var v21 int32 = 214
            var v22 uint16 = 65535
            _ = v20
            _ = v21
            _ = v22
        } else {
            fmt.Println(v16, v2)
        }
        v1["a"] = v3
        if (!v5) {
            v17 = (v17 << ((255) & 63))
            if false {
                v5 = (42 == 0)
                v23 := GenS2{F0: 0, F1: 1, F2: 1}
                switch ((("hello" + "") + "") + "\n\t") {
                case "foo bar":
                    fmt.Println(v3, v3)
                case "\n\t":
                    fmt.Println(v16)
                }
                _ = v23
            } else {
                v3 = v3
            }
            v5 = true
        }
    } else {
        v17 += v3
        v24 := GenS2{F0: 0, F1: 0, F2: 255}
        v25 := GenS2{F0: 65535, F1: 4294967295, F2: 0}
        _ = v24
        _ = v25
    }
    v14 = v14
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(v5)
        _ = fi26
    }
    switch genFn1((v16 << ((0) & 63))) {
    case 32767:
        for fi27 := 0; fi27 < 1; fi27++ {
            fmt.Println(v17, v5, v5)
            v28 := GenS1{F0: false, F1: 18446744073709551615}
            var v29 uint16 = uint16(v3)
            _ = fi27
            _ = v28
            _ = v29
        }
    case 2147483647:
        for fi30 := 0; fi30 < 11; fi30++ {
            fmt.Println(fi30)
            _ = fi30
        }
    }
    v31 := GenS3{F0: -469, F1: (v16 ^ uint8(v3))}
    v32 := GenS4{F0: 176, F1: 42, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
