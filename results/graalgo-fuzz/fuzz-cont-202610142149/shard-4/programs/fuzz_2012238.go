package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
func genFn0(p0 string, p1 int, p2 uint32) int64 {
    p2 -= p2
    v4 := []int{p1, 128, (p1 % ((p1) | 1))}
    p2 = (p2 >> ((1) & 63))
    v5 := GenS1{F0: p0, F1: ((2 < 255) && (!false))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int16 = 1
    v2 := []int{866, -1, -1}
    v3 := []int{-2147483648}
    if ((!false) && true) {
        var v6 uint16 = 42
        var v7 float64 = -2.25
        var v8 float32 = (1e10 / (0.0 - (1e10 * 3.14159)))
        _ = v6
        _ = v7
        _ = v8
    }
    fmt.Println(v1)
    fmt.Println(v1)
    v1 = v1
    fmt.Println(v1)
    if (828 <= 0) {
        if (false && false) {
            switch ("" + "a") {
            case "hello":
                _ = len(v3)
            case "\n\t":
                var v9 int = -149
                _ = v9
            case "foo bar":
                for fi10 := 0; fi10 < 15; fi10++ {
                    v1 = 1000
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v12 := GenS1{F0: ("" + "hello"), F1: false}
                        _ = fi11
                        _ = v12
                    }
                    fmt.Println(v1, v1)
                    _ = fi10
                }
            }
            var v13 float64 = 1.5
            switch ("hello" + "foo bar") {
            case "foo bar":
                if true {
                    if true {
                        var v14 float32 = 1.5
                        v15 := GenS1{F0: "\n\t", F1: (59 != 0)}
                        _ = v14
                        _ = v15
                    }
                    v1 = v1
                    var v16 uint32 = 2
                    _ = v16
                } else {
                    var v17 bool = (false || false)
                    v13 -= (v13 * float64(float32(-2.25)))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v17
                }
            case "hello":
                v1 += 0
            case "a":
                if (!false) {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
            default:
                var v18 uint8 = 0
                _ = v18
            }
            _ = v13
        }
    }
    if (!(!false)) {
        var v19 float64 = (-2.25 / (1.5 * 1e10))
        fmt.Println(v1)
        for fi20 := 0; fi20 < 17; fi20++ {
            v19 = v19
            for fi21 := 0; fi21 < 5; fi21++ {
                if true {
                    var v22 int = fi20
                    var v23 uint = 645
                    _ = v22
                    _ = v23
                } else {
                    if len(v3) > 0 {
                        v3[0] = 609
                    }
                    fmt.Println(v1, fi21)
                }
                var v24 string = (("foo bar" + "a") + "\n\t")
                _ = fi21
                _ = v24
            }
            _ = fi20
        }
        _ = v19
    } else {
        var v25 int64 = 578
        _ = len(v3)
        v1 = (v1 % ((v1) | 1))
        _ = v25
    }
    if (-9223372036854775808 == -1) {
        var v26 int8 = -4
        _ = v26
    } else {
        var v27 uint32 = 555
        fmt.Println(v1, v1)
        var v28 uint = 964
        _ = v27
        _ = v28
    }
    v1 = -220
    switch "a" {
    case "foo bar":
        v29 := GenS1{F0: "", F1: false}
        _ = v29
    case "\n\t":
        for fi30 := 0; fi30 < 20; fi30++ {
            fmt.Println(v1, fi30)
            _ = fi30
        }
    }
    v3 = append(v3, 378)
    v2 = append(v2, -9223372036854775808)
    for fi31 := 0; fi31 < 10; fi31++ {
        v2 = append(v2, (fi31 * 2147483647))
        if false {
            v32 := GenS2{F0: -60, F1: 607, F2: 0}
            _ = v32
        }
        v1 = (^(v1))
        _ = fi31
    }
    var v33 bool = (2147483647 >= 426)
    var v34 float64 = float64(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
