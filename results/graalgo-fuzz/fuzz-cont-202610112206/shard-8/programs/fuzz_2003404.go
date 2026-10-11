package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0(p0 uint64) uint {
    p0 = (^(p0))
    for fi4 := 0; fi4 < 1; fi4++ {
        fmt.Println(p0, p0)
        var v5 float32 = float32((1.5 - 1e10))
        _ = fi4
        _ = v5
    }
    if (76 < -1) {
        v6 := []int{9223372036854775807}
        _ = v6
    } else {
        var v7 uint32 = 2147483648
        _ = v7
    }
    v8 := GenS1{F0: 9223372036854775807, F1: 2147483647}
    _ = p0
    _ = v8
    return 65535
}
func genFn1(p0 uint8, p1 int, p2 uint64) uint32 {
    var v9 uint64 = p2
    v10 := GenS1{F0: 9223372036854775807, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 uint = 1
    var v2 uint64 = 954
    var v3 uint8 = 42
    var v11 float64 = 0.1
    v12 := GenS2{F0: 255, F1: 886}
    v13 := []int{32767}
    switch "hello" {
    case "foo bar":
        var v14 int = -390
        _ = v14
    case "\n\t":
        v15 := GenS1{F0: -1, F1: 0}
        _ = v15
    }
    _ = len(v13)
    switch 1 {
    case -9223372036854775808:
        v1 -= (v1 - (v1 ^ 899))
    case -908:
        v16 := GenS3{F0: 0}
        _ = v16
    default:
        v17 := GenS4{F0: 42}
        _ = v17
    }
    fmt.Println(v3)
    switch ("hello" + "\n\t") {
    case "a":
        for fi18 := 0; fi18 < 19; fi18++ {
            if (42 == 2147483647) {
                v19 := GenS4{F0: 1000}
                fmt.Println(fi18, fi18)
                v11 = (0.0 * v11)
                _ = v19
            } else {
                var v20 int32 = 42
                v11 += (v11 * ((((v11 / ((v11 - (0.1 / v11)) * 0.1)) / 1.0) + v11) + (-2.25 / -2.25)))
                fmt.Println(v3, v20)
                _ = v20
            }
            if false {
                var v21 uint = (v1 | v1)
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 float32 = float32(v11)
                    _ = fi22
                    _ = v23
                }
                _ = len(v13)
                _ = v21
            } else {
                fmt.Println(v1)
                if true {
                    v2 = v2
                }
                for fi24 := 0; fi24 < 7; fi24++ {
                    v13 = append(v13, fi18)
                    _ = fi24
                }
            }
            _ = fi18
        }
    case "\n\t":
        var v25 int = 2147483647
        _ = v25
    default:
        v13 = append(v13, 695)
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 18; fi27++ {
            fmt.Println(v1, v2)
            switch fi26 {
            case 128:
                fmt.Println(fi27, fi26)
            }
            switch int(v3) {
            case 516:
                v3 -= (^(v3))
            case 32768:
                var v28 string = ("foo bar" + ((("" + "foo bar") + "foo bar") + ("hello" + "")))
                _ = v28
            default:
                v2 = (v2 * v2)
            }
            _ = fi27
        }
        v29 := GenS5{F0: 0, F1: -1}
        _ = fi26
        _ = v29
    }
    v11 = 0.0
    fmt.Println(v1, v2)
    v11 -= v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
