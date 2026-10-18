package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
func genFn0() uint64 {
    switch ((("a" + "foo bar") + "hello") + "hello") {
    case "a":
        var v7 uint64 = 1
        _ = v7
    case "\n\t":
        switch 2147483647 {
        case -1:
            var v8 float64 = 1.5
            _ = v8
        default:
            var v9 int32 = -352
            _ = v9
        }
    case "":
        var v10 float64 = 0.0
        _ = v10
    }
    if true {
        var v11 bool = true
        _ = v11
    } else {
        switch 32767 {
        case -58:
            var v12 int32 = 1
            _ = v12
        case 32767:
            var v13 int32 = -1
            _ = v13
        default:
            var v14 uint32 = 4294967295
            _ = v14
        }
    }
    v15 := []int{32767}
    if len(v15) > 0 {
        v15[0] = 127
    }
    if true {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        v15 = append(v15, 42)
    } else {
        _ = len(v15)
        if len(v15) > 0 {
            v15[0] = -9223372036854775808
        }
    }
    _ = v15
    return 42
}
func genFn1(p0 uint64, p1 bool, p2 uint) uint32 {
    v16 := GenS1{F0: 52}
    fmt.Println(p2, p2, p0)
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            switch fi18 {
            case 42:
                p0 = p0
            case 32768:
                v19 := GenS2{F0: fi17}
                _ = v19
            default:
                v20 := []int{}
                _ = v20
            }
            fmt.Println(fi17)
            p0 = (p0 + p0)
            fmt.Println(p0, fi17, fi17)
            _ = fi18
        }
        p2 -= p2
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 2147483647
}

func main() {
    var v1 uint32 = 2
    v2 := []string{"foo bar"}
    v3 := []uint32{v1}
    var v4 int64 = 9223372036854775807
    var v5 int32 = 357
    var v6 int32 = 42
    var v21 float32 = -2.25
    v1 = genFn1(4294967295, (1.5 <= 1.5), 382)
    v4 = (^(v4))
    v21 -= float32((-2.25 * -2.25))
    fmt.Println(v5, v1, v21)
    if (32767 > 9223372036854775807) {
        if len(v3) > 0 {
            v3[0] = (v1 - (v1 ^ v1))
        }
        v22 := GenS1{F0: 255}
        if len(v3) > 0 {
            v3[0] = (^(v1))
        }
        _ = v22
    } else {
        v1 = genFn1(42, false, 812)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS3{F0: 42, F1: -128, F2: -137}
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 int32 = (^(v5))
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 int8 = 42
    _ = len(v2)
    var v28 uint8 = 128
    if (1000 <= 129) {
        var v29 uint32 = v1
        var v30 int32 = v6
        _ = v29
        _ = v30
    }
    fmt.Println(v27, v21, v27)
    if true {
        _ = len(v3)
    } else {
        fmt.Println(v4)
        for fi31 := 0; fi31 < 13; fi31++ {
            _ = len(v2)
            v32 := GenS1{F0: v28}
            v4 = v4
            _ = fi31
            _ = v32
        }
    }
    v21 = (v21 - v21)
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
