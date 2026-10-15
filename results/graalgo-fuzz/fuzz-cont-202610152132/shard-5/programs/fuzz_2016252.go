package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint64; F1 int32; F2 string }
func genFn0(p0 int, p1 uint16) uint8 {
    v4 := []int{(p0 % ((p0) | 1)), p0}
    var v5 float32 = -2.25
    var v6 uint32 = 4294967295
    fmt.Println(p1, v6)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int, p1 uint64, p2 uint16) int16 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 14; fi7++ {
        p1 = 432
        switch "foo bar" {
        case "\n\t":
            var v8 uint32 = 2
            _ = v8
        case "foo bar":
            var v9 uint8 = genFn0((fi7 - 2147483647), p2)
            _ = v9
        default:
            var v10 bool = (!(true && (-32768 != 0)))
            _ = v10
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 bool = (-194 <= 0)
            _ = fi11
            _ = v12
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 string = "foo bar"
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v1[65535] = 65535
    v15 := GenS1{F0: 106, F1: 0}
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 float32 = ((3.14159 + (1e10 - -2.25)) - 0.0)
        _ = fi16
        _ = v17
    }
    v3[2147483647] = 20
    if true {
        var v18 int8 = -17
        v19 := []int{42, 507, 42, 128}
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                switch "" {
                case "":
                    switch "hello" {
                    case "a":
                        delete(v1, 2)
                    }
                }
                _ = fi21
            }
            v18 -= v18
            var v22 int32 = -2147483648
            _ = fi20
            _ = v22
        }
        _ = v18
        _ = v19
    } else {
        var v23 string = (("" + "foo bar") + "foo bar")
        _ = v23
    }
    if (51 < 1) {
        var v24 uint64 = 2
        fmt.Println(v24)
        _ = v24
    }
    var v25 uint8 = 42
    switch 2147483647 {
    case -1:
        var v26 float64 = float64(float32(3.14159))
        _ = v26
    case -350:
        switch -48 {
        case -2147483648:
            switch 32767 {
            case 1:
                if (!true) {
                    v25 += v25
                    v25 += (v25 + v25)
                } else {
                    var v27 int16 = 1000
                    if (950 < -866) {
                        fmt.Println(v27)
                    } else {
                        fmt.Println(v27, v25)
                        switch 32768 {
                        case 9223372036854775807:
                            var v28 uint = 65535
                            _ = v28
                        case -962:
                            fmt.Println(v25)
                        }
                    }
                    v29 := v2[0]
                    _ = v27
                    _ = v29
                }
            case 42:
                v30 := GenS1{F0: 1, F1: 100}
                _ = v30
            default:
                v31 := GenS1{F0: 1, F1: 42}
                _ = v31
            }
        case 127:
            v25 -= (v25 >> ((808) & 63))
        }
    }
    if true {
        var v32 float64 = (((float64(1e10) / 0.1) - float64(1e10)) / ((0.1 - (0.1 * (0.1 + (1.5 + 0.0)))) / 1.0))
        fmt.Println(v32)
        _ = v32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        for fi34 := 0; fi34 < 8; fi34++ {
            v35 := GenS2{F0: v25, F1: 42}
            _ = fi34
            _ = v35
        }
        if true {
            switch "a" {
            case "":
                v25 -= (^(v25))
            case "a":
                v25 = (v25 | 128)
            }
            v36 := GenS3{F0: false, F1: 255}
            var v37 bool = ("foo bar" == "hello")
            _ = v36
            _ = v37
        }
        v25 -= v25
        _ = fi33
    }
    v38 := GenS4{F0: 207, F1: 1, F2: "foo bar"}
    v25 += genFn0(-9223372036854775808, 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v38
}
