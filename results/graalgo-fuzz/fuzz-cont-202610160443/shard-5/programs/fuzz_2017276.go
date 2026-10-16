package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0() int32 {
    if true {
        if (255 < 0) {
            var v4 int = -1
            v4 -= v4
            _ = v4
        } else {
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := GenS1{F0: 39, F1: 18446744073709551615}
                _ = fi5
                _ = v6
            }
            var v7 uint64 = 1
            v8 := []int{32768, -2147483648, 237, 612}
            _ = v7
            _ = v8
        }
        var v9 float64 = (1e10 * ((float64(((((1e10 / 1.5) * 0.1) * 1e10) / ((1.5 * 3.14159) - 0.0))) * 1e10) - -2.25))
        _ = v9
    }
    var v10 uint = 719
    v11 := GenS2{F0: false, F1: 155}
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 int, p1 bool) int32 {
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "" {
        case "foo bar":
            p0 = p0
        case "hello":
            v13 := []int{}
            _ = v13
        default:
            v14 := []int{}
            _ = v14
        }
        _ = fi12
    }
    var v15 uint8 = 43
    if (p1 && p1) {
        v15 -= (v15 / (((v15 >> ((974) & 63))) | 1))
        var v16 uint8 = v15
        _ = v16
    } else {
        v17 := GenS2{F0: ((float32((1.5 - float64(1e10))) != 3.14159) || (!(212 >= 65535))), F1: 1}
        v15 += (v15 ^ 255)
        _ = v17
    }
    var v18 uint = 2
    fmt.Println(v18, p0, v18)
    _ = p0
    _ = p1
    _ = v15
    _ = v18
    return -1
}

func main() {
    var v1 int32 = -955
    var v2 int16 = 37
    v3 := make(map[string]uint32)
    if ((18446744073709551615 >= 744) && (!false)) {
        v19 := []int{2147483647, -105, -732, -1}
        var v20 uint = 384
        _ = v19
        _ = v20
    }
    v21 := []int{-9223372036854775808, 199}
    delete(v3, "foo bar")
    v22 := v3[("foo bar" + "")]
    v23 := GenS2{F0: (0 > -14), F1: 255}
    v24 := GenS3{F0: v1, F1: ((v1 / ((genFn1(128, true)) | 1)) != (v1 | genFn0())), F2: 0}
    if true {
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int32 = (v1 | v1)
            _ = fi25
            _ = v26
        }
        var v27 float32 = ((0.0 + (-2.25 / ((-2.25 + (3.14159 * 1e10)) - -2.25))) / 1.5)
        _ = v27
    } else {
        for fi28 := 0; fi28 < 12; fi28++ {
            if false {
                v1 -= (v1 - v1)
            } else {
                fmt.Println(v22, v1)
            }
            _ = fi28
        }
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 int8 = 108
            v30 = 127
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            fmt.Println(v22, v2, fi31)
            _ = fi31
        }
    }
    switch (("hello" + "") + ("\n\t" + "\n\t")) {
    case "":
        var v32 uint32 = v22
        _ = v32
    default:
        fmt.Println(v22, v2, v22)
    }
    v2 = (v2 % ((-269) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
