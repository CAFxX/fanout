package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16 }
func genFn0(p0 int16) uint8 {
    var v6 int64 = 0
    if ((-(p0)) == p0) {
        p0 = (p0 & (p0 % ((p0) | 1)))
    } else {
        switch "hello" {
        case "":
            v7 := GenS1{F0: 32767, F1: v6}
            _ = v7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            v6 = (v6 | v6)
            _ = fi8
        }
    }
    if true {
        if (0.0 != 0.0) {
            switch "\n\t" {
            case "a":
                var v9 int32 = 42
                _ = v9
            case "\n\t":
                v10 := GenS1{F0: p0, F1: -547}
                _ = v10
            }
        }
        var v11 float64 = float64((0.1 + 0.1))
        _ = v11
    } else {
        switch (("a" + "foo bar") + "hello") {
        case "":
            v6 += v6
        default:
            p0 = (p0 | 484)
        }
        if false {
            p0 = p0
        }
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 uint8) uint16 {
    switch "hello" {
    case "hello":
        switch "a" {
        case "hello":
            fmt.Println(p0)
        default:
            p0 = 127
        }
    default:
        v12 := []int{2147483647, -1, 42, -464}
        _ = v12
    }
    switch 532 {
    case 2147483647:
        p0 -= (p0 / (((p0 % (((p0 >> ((4294967295) & 63))) | 1))) | 1))
    case -311:
        v13 := GenS1{F0: 1000, F1: 0}
        _ = v13
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 42
    var v3 uint32 = 0
    var v4 int16 = 0
    var v5 uint8 = 1
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int64 = -1
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v3 = 4294967295
        var v17 float64 = (float64((float32(-2.25) / 0.1)) - (0.1 + 1.5))
        var v18 uint16 = 416
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int8 = -1
    v2 = (v2 >> ((1) & 63))
    v20 := v1[v3]
    var v21 uint16 = 0
    v22 := GenS1{F0: int16(v5), F1: 42}
    _ = len(v1)
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := GenS2{F0: uint(v21), F1: (v19 / ((100) | 1))}
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 float32 = 1.5
            var v27 int32 = 163
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    v28 := GenS1{F0: (^(v4)), F1: int64(v2)}
    v21 = 642
    v21 += v21
    for fi29 := 0; fi29 < 13; fi29++ {
        delete(v1, (v3 - 0))
        v3 -= (^(v3))
        v30 := GenS2{F0: 1, F1: (v19 >> ((1) & 63))}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint = 0
        switch "a" {
        case "":
            switch "\n\t" {
            case "":
                _ = len(v1)
            case "foo bar":
                v20 = (-9223372036854775808 > 667)
            default:
                v33 := GenS3{F0: v19, F1: "\n\t"}
                _ = v33
            }
        case "a":
            v2 += (^(v2))
        default:
            v34 := GenS4{F0: (v4 % ((243) | 1))}
            _ = v34
        }
        var v35 string = ("" + "\n\t")
        _ = fi31
        _ = v32
        _ = v35
    }
    var v36 uint64 = uint64(v5)
    fmt.Println(v20, v4, v21)
    v19 += (v19 | (v19 << ((636) & 63)))
    v21 -= v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
