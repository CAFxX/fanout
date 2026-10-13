package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 int32) uint {
    if (42 > 1) {
        var v6 int16 = 42
        var v7 int16 = 0
        v8 := []int{622, 856, 494}
        _ = v6
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch fi9 {
            case 0:
                var v10 uint16 = 65535
                _ = v10
            case 42:
                v11 := []int{127, 128, (fi9 * (-(fi9))), 42}
                _ = v11
            }
            _ = fi9
        }
        if true {
            p0 += (p0 >> ((1) & 63))
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(p0)
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: 0, F2: -956}
    var v15 int8 = 1
    var v16 uint8 = 2
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 556
}
func genFn1(p0 uint64) int64 {
    p0 -= 764
    switch (("a" + "foo bar") + "a") {
    case "hello":
        p0 = (p0 / (((p0 - 42)) | 1))
    default:
        for fi17 := 0; fi17 < 16; fi17++ {
            p0 = (p0 & 750)
            p0 = p0
            _ = fi17
        }
    }
    v18 := []int{-613}
    fmt.Println(p0)
    if (0 != 766) {
        v19 := GenS1{F0: p0, F1: 2147483647, F2: 0}
        var v20 string = "foo bar"
        var v21 int64 = -473
        _ = v19
        _ = v20
        _ = v21
    }
    _ = p0
    _ = v18
    return -9223372036854775808
}

func main() {
    var v1 int32 = -344
    v2 := make(map[uint32]bool)
    v3 := []int{-1, 32768}
    v4 := []uint32{851}
    var v5 float32 = 0.1
    var v22 uint32 = 42
    fmt.Println(v5, v22, v22)
    if (8 <= -91) {
        v23 := GenS1{F0: 18446744073709551615, F1: v22, F2: 42}
        var v24 int = 32768
        var v25 int32 = (v1 | int32(v24))
        _ = v23
        _ = v24
        _ = v25
    } else {
        switch "\n\t" {
        case "foo bar":
            if (!true) {
                var v26 int64 = genFn1(18446744073709551615)
                var v27 uint16 = 2
                var v28 float64 = 0.1
                _ = v26
                _ = v27
                _ = v28
            }
        case "\n\t":
            v22 += v22
        }
        v22 -= 1
        v1 -= 42
    }
    v2[(v22 % ((65535) | 1))] = true
    v5 = ((v5 / (v5 * (v5 * (v5 * (-2.25 * (v5 - v5)))))) - (((v5 / (v5 * (v5 - v5))) + 0.0) - v5))
    var v29 bool = (!(true || ((0.1 > 3.14159) && false)))
    var v30 uint8 = 2
    _ = len(v4)
    if ("hello" != "foo bar") {
        var v31 uint8 = (^(v30))
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 string = (("a" + (("\n\t" + "a") + "hello")) + "a")
        fmt.Println(v33, v29, v5)
        v34 := GenS1{F0: 18446744073709551615, F1: v22, F2: 0}
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
