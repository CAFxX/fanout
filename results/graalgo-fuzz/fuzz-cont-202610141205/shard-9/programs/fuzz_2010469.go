package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint16 = 42
        _ = fi4
        _ = v5
    }
    var v6 uint64 = 648
    _ = v6
    return v6
}
func genFn1(p0 int32, p1 int64, p2 bool) int64 {
    v7 := []int{0, 2147483647}
    switch 930 {
    case 9223372036854775807:
        switch "foo bar" {
        case "\n\t":
            switch ("a" + "\n\t") {
            case "foo bar":
                fmt.Println(p0, p1)
            case "":
                var v8 int64 = 1
                _ = v8
            default:
                v7 = append(v7, -435)
            }
        case "a":
            fmt.Println(p0)
        case "hello":
            switch 998 {
            case 9223372036854775807:
                if true {
                    v9 := GenS1{F0: 0, F1: 2}
                    if p2 {
                        fmt.Println(p1)
                        v10 := GenS1{F0: -65, F1: 267}
                        fmt.Println(p1, p0)
                        _ = v10
                    }
                    _ = v9
                }
            case 4294967295:
                fmt.Println(p0, p1)
            case 127:
                p2 = p2
            }
        }
    }
    v11 := GenS1{F0: -126, F1: 803}
    var v12 int32 = p0
    switch -2147483648 {
    case 0:
        switch -155 {
        case -1:
            var v13 int8 = 0
            _ = v13
        default:
            v14 := GenS1{F0: -1, F1: 42}
            _ = v14
        }
    case 32767:
        p0 = (^(v12))
    case -9223372036854775808:
        p1 += (p1 - (-(p1)))
    default:
        p1 += (p1 - -155)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v7
    return 1
}

func main() {
    v1 := []string{("foo bar" + "foo bar"), "hello", ("hello" + "a"), "a"}
    var v2 float64 = ((((1.5 / 3.14159) * -2.25) * 1e10) / (((3.14159 * 0.0) * 0.1) - (1e10 - -2.25)))
    v3 := []int{165, 127}
    if false {
        if (!false) {
            _ = len(v3)
            fmt.Println(v2)
            if (18446744073709551615 < 919) {
                if false {
                    if (621 == genFn0()) {
                        switch "" {
                        case "hello":
                            fmt.Println(v2)
                        case "\n\t":
                            var v15 uint32 = 1
                            _ = v15
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v16 uint32 = 0
                _ = v16
            } else {
                v17 := GenS2{F0: -1, F1: 2}
                _ = len(v1)
                _ = v17
            }
        } else {
            var v18 int8 = -1
            v18 = (^(v18))
            v19 := GenS1{F0: v18, F1: 2}
            _ = v18
            _ = v19
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v20 uint64 = 219
        _ = v20
    }
    var v21 int16 = 1
    for fi22 := 0; fi22 < 2; fi22++ {
        v21 += -954
        var v23 uint32 = 1
        var v24 int64 = 9223372036854775807
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 float32 = (1.5 / 1e10)
    var v26 int8 = 0
    var v27 string = (("" + ("" + "\n\t")) + "hello")
    if (559 > 2) {
        for fi28 := 0; fi28 < 13; fi28++ {
            switch "foo bar" {
            case "":
                for fi29 := 0; fi29 < 6; fi29++ {
                    var v30 int32 = -150
                    _ = fi29
                    _ = v30
                }
            case "a":
                var v31 uint64 = 581
                _ = v31
            case "hello":
                if (false && true) {
                    v32 := GenS2{F0: -9223372036854775808, F1: 0}
                    v26 = v26
                    _ = v32
                } else {
                    fmt.Println(v27, v2)
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 float64 = v2
                        _ = fi33
                        _ = v34
                    }
                    v27 = (v27 + "foo bar")
                }
            }
            _ = fi28
        }
    }
    var v35 int = 4294967295
    var v36 int = v35
    v25 = (float32(v2) / (v25 - 1.5))
    fmt.Println(v21, v25, v2)
    for fi37 := 0; fi37 < 18; fi37++ {
        v25 = float32(float64(v25))
        v38 := GenS2{F0: -350, F1: 2}
        v21 += v21
        _ = fi37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
}
