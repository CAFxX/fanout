package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 uint) bool {
    if false {
        var v6 int = 128
        v7 := GenS1{F0: 1}
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return true
}
func genFn1(p0 uint64) int32 {
    var v8 uint32 = 1
    if true {
        switch 128 {
        case -658:
            fmt.Println(v8, v8)
        }
    }
    fmt.Println(v8, p0)
    _ = p0
    _ = v8
    return 1
}

func main() {
    v1 := []int{}
    var v2 int32 = 238
    var v3 float64 = -2.25
    var v4 int64 = -9223372036854775808
    v5 := []uint32{797, 2147483648, 961, 153}
    v9 := GenS1{F0: -558}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 4294967295 {
    case 29:
        v10 := GenS2{F0: (false || true), F1: 762}
        _ = v10
    default:
        var v11 uint8 = 19
        _ = v11
    }
    switch "a" {
    case "\n\t":
        for fi12 := 0; fi12 < 18; fi12++ {
            v2 = v2
            var v13 uint32 = 220
            _ = fi12
            _ = v13
        }
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if true {
            var v14 float64 = 0.0
            v5 = append(v5, 398)
            switch ("\n\t" + "") {
            case "\n\t":
                v15 := GenS1{F0: -32768}
                _ = v15
            case "a":
                fmt.Println(v3)
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = v14
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v16 := GenS1{F0: -763}
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 float64 = 3.14159
        _ = fi17
        _ = v18
    }
    v3 = 0.1
    var v19 uint16 = 437
    v19 += 869
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    if ((v4 - v4) >= v4) {
        var v20 string = ""
        _ = v20
    } else {
        switch "a" {
        case "":
            v21 := GenS3{F0: -1, F1: 18446744073709551615}
            _ = v21
        case "foo bar":
            for fi22 := 0; fi22 < 8; fi22++ {
                v4 += 0
                var v23 float32 = (0.0 / 0.1)
                for fi24 := 0; fi24 < 11; fi24++ {
                    v25 := GenS1{F0: 42}
                    v26 := GenS3{F0: 47, F1: 255}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi22
                _ = v23
            }
        default:
            var v27 uint = 769
            _ = v27
        }
        var v28 uint64 = 70
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
