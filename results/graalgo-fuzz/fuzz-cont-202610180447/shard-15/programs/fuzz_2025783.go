package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint16; F1 int32; F2 int }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16) uint64 {
    p0 -= (p0 | 1)
    if (p0 > (p0 | (p0 % ((p0) | 1)))) {
        switch 462 {
        case 32768:
            var v5 uint64 = 42
            _ = v5
        case -477:
            if true {
                p0 -= (p0 + (-(p0)))
            }
        case 4294967295:
            fmt.Println(p0)
        default:
            switch "foo bar" {
            case "foo bar":
                v6 := []int{2147483647, 4294967295, -9223372036854775808, 1}
                _ = v6
            case "hello":
                if ((false || true) && (false || false)) {
                    switch 1 {
                    case -9223372036854775808:
                        v7 := GenS1{F0: 1, F1: -225}
                        _ = v7
                    }
                    var v8 int32 = 557
                    _ = v8
                }
            }
        }
        v9 := []int{0, 127}
        fmt.Println(p0)
        _ = v9
    }
    _ = p0
    return 7
}

func main() {
    v1 := []string{}
    var v2 int8 = 38
    v3 := make(map[int]bool)
    v4 := []string{("hello" + ("hello" + (("a" + "\n\t") + "hello"))), "\n\t"}
    if (false && (794 <= 869)) {
        var v10 uint16 = 902
        fmt.Println(v2, v2)
        v10 = v10
        _ = v10
    }
    switch "a" {
    case "foo bar":
        _ = len(v1)
    default:
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    }
    if false {
        if (true || true) {
            for fi11 := 0; fi11 < 8; fi11++ {
                v2 -= (v2 + (v2 * v2))
                var v12 uint8 = 42
                v12 -= (v12 - (v12 | v12))
                _ = fi11
                _ = v12
            }
        } else {
            switch "hello" {
            case "":
                var v13 uint32 = 65535
                _ = v13
            case "a":
                v1 = append(v1, "hello")
            case "foo bar":
                var v14 uint32 = 1
                _ = v14
            default:
                v15 := GenS2{F0: 255, F1: 769, F2: -2147483648}
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            if len(v1) > 0 {
                v1[0] = ("\n\t" + "")
            }
            v17 := GenS1{F0: 2, F1: -1}
            _ = fi16
            _ = v17
        }
        fmt.Println(v2)
    }
    v18 := GenS2{F0: 2, F1: 439, F2: 2147483647}
    v19 := v3[128]
    for fi20 := 0; fi20 < 8; fi20++ {
        _ = len(v4)
        _ = len(v3)
        v19 = true
        _ = fi20
    }
    v21 := GenS3{F0: 9223372036854775807, F1: (v2 + (-(v2))), F2: 0}
    for fi22 := 0; fi22 < 5; fi22++ {
        v2 -= (v2 * v2)
        fmt.Println(v19, v19)
        _ = fi22
    }
    var v23 int16 = -32768
    var v24 int = 42
    fmt.Println(v24)
    v25 := GenS4{F0: v2}
    var v26 int8 = -1
    var v27 float32 = (1e10 * ((((((1.5 / 1.5) - (3.14159 + float32(0.0))) * 0.1) * (1.5 + -2.25)) + (float32((0.0 / 1.5)) - 1e10)) / 1.5))
    v28 := GenS5{F0: 42}
    delete(v3, 4294967295)
    if ("\n\t" > "\n\t") {
        delete(v3, v24)
        v27 -= (-2.25 + (v27 - ((float32(0.0) + v27) / 1.0)))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
