package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 bool }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint64; F1 int }
func genFn0(p0 uint32) int32 {
    v6 := []int{}
    p0 += p0
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{1}
    v3 := make(map[string]int)
    var v4 float64 = (-2.25 + 1.5)
    var v5 uint8 = 127
    if ("foo bar" < "") {
        v7 := GenS1{F0: ("" + "hello")}
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float32 = 1e10
            v5 = (v5 & 0)
            var v10 int = fi8
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        v11 := GenS2{F0: 7, F1: uint32(v5)}
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS3{F0: (^(fi13)), F1: 669, F2: false}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v5 = (v5 * 128)
    v15 := GenS4{F0: true, F1: 65535, F2: 70}
    if (v5 == (v5 - (v5 * 0))) {
        v16 := GenS5{F0: -1}
        _ = v16
    }
    v5 -= (v5 / ((255) | 1))
    if (!true) {
        v4 -= 0.0
        var v17 uint64 = 736
        if true {
            var v18 int8 = 25
            var v19 uint64 = (v17 * 270)
            for fi20 := 0; fi20 < 2; fi20++ {
                for fi21 := 0; fi21 < 6; fi21++ {
                    switch ("a" + "foo bar") {
                    case "hello":
                        v4 = (v4 * float64(0.1))
                    default:
                        _ = len(v2)
                    }
                    if true {
                        fmt.Println(fi21, v19, v19)
                        v22 := GenS3{F0: 32767, F1: 65535, F2: true}
                        _ = v22
                    }
                    _ = fi21
                }
                _ = fi20
            }
            _ = v18
            _ = v19
        } else {
            v23 := GenS6{F0: 319, F1: "\n\t"}
            switch (("hello" + (("" + ("foo bar" + "hello")) + "foo bar")) + "foo bar") {
            case "a":
                v24 := GenS5{F0: 608}
                _ = v24
            case "foo bar":
                var v25 uint = 255
                _ = v25
            case "\n\t":
                switch -9223372036854775808 {
                case 0:
                    v26 := GenS7{F0: 18446744073709551615, F1: -1}
                    _ = v26
                case 376:
                    fmt.Println(v5, v4, v4)
                case 867:
                    v27 := GenS1{F0: ("a" + "\n\t")}
                    _ = v27
                default:
                    fmt.Println(v5, v4, v5)
                }
            }
            var v28 bool = true
            _ = v23
            _ = v28
        }
        _ = v17
    } else {
        if false {
            fmt.Println(v4, v5)
            for fi29 := 0; fi29 < 5; fi29++ {
                if ((fi29 < (fi29 & fi29)) && true) {
                    fmt.Println(v4)
                    switch "\n\t" {
                    case "a":
                        var v30 int16 = -1
                        _ = v30
                    case "hello":
                        v4 += v4
                    case "foo bar":
                        v5 = 15
                    default:
                        fmt.Println(fi29)
                    }
                } else {
                    v5 = (v5 ^ 0)
                    v4 += v4
                }
                v31 := v1[(fi29 << ((2) & 63))]
                _ = fi29
                _ = v31
            }
        }
        var v32 uint32 = 579
        _ = v32
    }
    var v33 int16 = -353
    if false {
        v34 := GenS7{F0: 42, F1: -2147483648}
        _ = v34
    } else {
        switch "a" {
        case "":
            for fi35 := 0; fi35 < 4; fi35++ {
                delete(v1, fi35)
                v36 := GenS5{F0: genFn0(337)}
                fmt.Println(v4, v4, v33)
                _ = fi35
                _ = v36
            }
        default:
            v37 := GenS7{F0: 0, F1: -9223372036854775808}
            _ = v37
        }
        v5 = 67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
