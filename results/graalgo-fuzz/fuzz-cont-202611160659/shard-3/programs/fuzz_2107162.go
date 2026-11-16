package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 int }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0() int16 {
    var v5 string = "\n\t"
    var v6 int32 = 0
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 int8, p1 int8) string {
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 -= p0
        _ = fi7
    }
    v8 := []int{-755, 9223372036854775807, 0, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v8
    return "a"
}

func main() {
    v1 := []uint32{}
    v2 := []string{}
    v3 := []int{1, 2147483647, -9223372036854775808, 42}
    var v4 uint16 = 65535
    if (0.0 != (1e10 + (float32((1.5 - (1e10 * 0.1))) - 0.0))) {
        var v9 float64 = (float64(0.0) + (float64(0.0) * (3.14159 / 3.14159)))
        switch "foo bar" {
        case "hello":
            var v10 int32 = 582
            _ = v10
        case "\n\t":
            v9 = 3.14159
        }
        _ = v9
    }
    var v11 uint8 = 0
    fmt.Println(v4, v11)
    v4 = (v4 + v4)
    for fi12 := 0; fi12 < 15; fi12++ {
        if len(v1) > 0 {
            v1[0] = 879
        }
        var v13 int16 = 341
        fmt.Println(v13)
        _ = fi12
        _ = v13
    }
    v11 -= (v11 >> ((42) & 63))
    switch "foo bar" {
    case "hello":
        v4 = (^(v4))
    case "a":
        if (false || true) {
            v14 := GenS1{F0: "a", F1: 128}
            _ = v14
        } else {
            fmt.Println(v11)
            v4 -= 632
        }
    }
    var v15 float64 = ((((1.5 - -2.25) * float64(3.14159)) / 1e10) * 1e10)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v4 = v4
        var v17 int16 = -552
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 int64 = 42
            var v20 int8 = 100
            if len(v1) > 0 {
                v1[0] = 42
            }
            v21 := GenS2{F0: 2}
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                fmt.Println(v17, v4)
            case "":
                _ = len(v3)
            }
            switch 65 {
            case 32768:
                var v22 uint8 = (v11 % ((v11) | 1))
                _ = v22
            case 0:
                v23 := GenS3{F0: uint64(v4), F1: 661, F2: int(v4)}
                _ = v23
            case -2147483648:
                switch fi16 {
                case -1:
                    v24 := GenS4{F0: genFn0(), F1: v19}
                    _ = v24
                }
            default:
                v11 -= (v11 / ((v11) | 1))
            }
            var v25 uint32 = 2
            var v26 int16 = v17
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
            _ = v25
            _ = v26
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
