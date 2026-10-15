package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 uint }
type GenS4 struct { F0 uint; F1 int; F2 bool }
type GenS5 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]string)
    var v3 int8 = 100
    v4 := []uint32{}
    v4 = append(v4, 42)
    fmt.Println(v3)
    var v5 bool = true
    if v5 {
        v6 := GenS1{F0: 1, F1: (v3 % ((v3) | 1)), F2: 2147483647}
        switch 42 {
        case 1:
            switch ("foo bar" + "a") {
            case "\n\t":
                for fi7 := 0; fi7 < 8; fi7++ {
                    for fi8 := 0; fi8 < 1; fi8++ {
                        switch (("hello" + "") + "hello") {
                        case "hello":
                            fmt.Println(v3, fi7, v5)
                        case "foo bar":
                            if ((!v5) || false) {
                                if (!(!false)) {
                                    if len(v4) > 0 {
                                        v4[0] = 0
                                    }
                                    v9 := GenS2{F0: 1}
                                    var v10 int32 = 85
                                    _ = v10
                                    _ = v9
                                } else {
                                    v11 := GenS2{F0: 0}
                                    var v12 int32 = 781
                                    _ = v11
                                    _ = v12
                                }
                                switch ("" + "") {
                                case "a":
                                    v3 -= (v3 | (v3 - (v3 % (((v3 << ((0) & 63))) | 1))))
                                default:
                                    switch -9223372036854775808 {
                                    case 222:
                                        v5 = v5
                                    case 128:
                                        fmt.Println(v1, fi7, v3)
                                    default:
                                        v1 += v1
                                    }
                                }
                            }
                        }
                        _ = fi8
                    }
                    _ = fi7
                }
            case "foo bar":
                _ = len(v4)
            case "":
                var v13 bool = (-1 >= 42)
                _ = v13
            default:
                var v14 uint8 = 128
                _ = v14
            }
        case 127:
            v1 = v1
        default:
            switch "\n\t" {
            case "a":
                v5 = false
            case "hello":
                var v15 string = "hello"
                _ = v15
            case "":
                switch -842 {
                case 4294967295:
                    fmt.Println(v3, v1)
                case 1:
                    if true {
                        v4 = append(v4, 2)
                        v16 := GenS1{F0: 18446744073709551615, F1: (v3 >> ((255) & 63)), F2: -536}
                        _ = v16
                    } else {
                        var v17 int32 = 1
                        fmt.Println(v3, v1, v1)
                        v17 = v17
                        _ = v17
                    }
                }
            default:
                switch "a" {
                case "\n\t":
                    fmt.Println(v1, v5)
                case "foo bar":
                    var v18 string = ("" + ("" + ("\n\t" + "foo bar")))
                    _ = v18
                case "a":
                    var v19 string = "foo bar"
                    _ = v19
                }
            }
        }
        _ = v6
    } else {
        for fi20 := 0; fi20 < 19; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 bool = ("\n\t" > ("\n\t" + ""))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = len(v4)
                _ = fi21
                _ = v22
            }
            v5 = (2147483647 < 42)
            _ = fi20
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 float64 = ((0.0 * (0.0 + float64(((float32((0.1 - 0.1)) - 0.1) / (float32(-2.25) - 0.1))))) + 1e10)
        v25 := GenS3{F0: 42, F1: 95, F2: 4294967295}
        v26 := GenS4{F0: 221, F1: -357, F2: (!((42 > -1) && ((18446744073709551615 > 1) || true)))}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    if (113 != 492) {
        fmt.Println(v5)
        v1 = 651
        var v27 bool = false
        _ = v27
    } else {
        v28 := GenS5{F0: 523, F1: (v3 % ((v3) | 1)), F2: 65535}
        v29 := v2[-9223372036854775808]
        if (v5 && v5) {
            var v30 uint8 = 2
            v31 := GenS5{F0: uint16(v30), F1: (v3 * v3), F2: 682}
            for fi32 := 0; fi32 < 12; fi32++ {
                v3 = (-(v3))
                _ = fi32
            }
            _ = v30
            _ = v31
        } else {
            v5 = (4294967295 == 255)
        }
        _ = v28
        _ = v29
    }
    v33 := GenS6{F0: 2147483647}
    v5 = (!v5)
    var v34 int32 = -414
    v35 := GenS7{F0: 160, F1: 0}
    var v36 bool = false
    fmt.Println(v36, v34, v34)
    switch "foo bar" {
    case "a":
        v5 = v36
    case "foo bar":
        if true {
            fmt.Println(v5, v3)
            v34 += v34
            for fi37 := 0; fi37 < 15; fi37++ {
                _ = len(v2)
                _ = fi37
            }
        }
    case "\n\t":
        fmt.Println(v5)
    default:
        fmt.Println(v36, v5)
    }
    if v36 {
        fmt.Println(v3, v1)
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
