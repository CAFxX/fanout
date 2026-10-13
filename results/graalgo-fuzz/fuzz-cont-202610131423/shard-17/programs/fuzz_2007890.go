package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint64; F1 string; F2 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint8 = 1
    var v2 int16 = 0
    var v3 int16 = v2
    switch ("hello" + "foo bar") {
    case "foo bar":
        v1 = (^(v1))
    }
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 uint32 = 2147483648
        if (true && false) {
            var v6 uint16 = uint16(v5)
            for fi7 := 0; fi7 < 5; fi7++ {
                v1 -= 128
                _ = fi7
            }
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v2 += v3
    v8 := []int{-722, 128, 2147483647}
    v9 := GenS1{F0: 18446744073709551615, F1: 825, F2: false}
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int16 = (v2 - v3)
        if (65535 < uint32(v1)) {
            v3 = 154
        } else {
            if (241 == 4294967295) {
                switch "" {
                case "\n\t":
                    if len(v8) > 0 {
                        v8[0] = fi10
                    }
                default:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
            }
            switch ("\n\t" + ("\n\t" + "hello")) {
            case "\n\t":
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 int8 = 0
                    if (false || true) {
                        v14 := GenS2{F0: false, F1: 65535}
                        switch "a" {
                        case "foo bar":
                            fmt.Println(fi10)
                        case "":
                            if len(v8) > 0 {
                                v8[0] = fi10
                            }
                        case "hello":
                            v13 = v13
                        }
                        if (!false) {
                            v1 = 255
                        }
                        _ = v14
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            v1 = v1
                        default:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                        var v15 int = (fi12 % ((fi12) | 1))
                        _ = len(v8)
                        switch fi10 {
                        case -9223372036854775808:
                            var v16 int32 = 0
                            _ = v16
                        case 32767:
                            v1 -= (v1 * (v1 | 42))
                            var v17 int = v15
                            _ = v17
                        case 0:
                            var v18 int8 = (v13 & v13)
                            _ = v18
                        }
                        switch "a" {
                        case "foo bar":
                            v8 = append(v8, -2147483648)
                        case "a":
                            var v19 int64 = 0
                            _ = v19
                        case "hello":
                            v3 = (v11 >> ((114) & 63))
                        }
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
            case "foo bar":
                var v20 int32 = 0
                _ = v20
            case "hello":
                switch ("" + ((((("hello" + "hello") + "\n\t") + "a") + "") + "hello")) {
                case "a":
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 float32 = 1e10
                        var v23 uint8 = 199
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                default:
                    v24 := GenS1{F0: 18446744073709551615, F1: 32767, F2: false}
                    _ = v24
                }
            default:
                v3 = (v3 << ((4294967295) & 63))
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        if false {
            fmt.Println(v2, v3)
            for fi25 := 0; fi25 < 4; fi25++ {
                fmt.Println(v2, v1, v1)
                if false {
                    v1 += (^(v1))
                } else {
                    var v26 uint = 0
                    var v27 uint32 = 737
                    var v28 int8 = 8
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
            fmt.Println(v1, v11, fi10)
        }
        _ = fi10
        _ = v11
    }
    v29 := GenS3{F0: 42, F1: "\n\t", F2: 18446744073709551615}
    for fi30 := 0; fi30 < 19; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            if (2 < 352) {
                switch "hello" {
                case "":
                    var v32 string = "hello"
                    _ = v32
                default:
                    var v33 uint64 = 42
                    _ = v33
                }
                fmt.Println(fi30, v2, fi31)
            } else {
                fmt.Println(v2, v2, v3)
                _ = len(v8)
                v8 = append(v8, -861)
                var v34 bool = (true || (32767 >= 866))
                _ = v34
            }
            _ = len(v8)
            _ = fi31
        }
        _ = fi30
    }
    if false {
        fmt.Println(v3)
        v35 := GenS4{F0: -128, F1: (v1 / (((^(v1))) | 1)), F2: 13}
        _ = v35
    } else {
        var v36 uint64 = 916
        if false {
            v37 := GenS2{F0: (0.0 <= 1e10), F1: 835}
            _ = v37
        }
        v1 = (v1 >> ((800) & 63))
        _ = v36
    }
    if false {
        if (v2 == v3) {
            switch 307 {
            case -489:
                fmt.Println(v1)
            case 71:
                var v38 int32 = 2147483647
                _ = v38
            case -418:
                var v39 uint64 = 333
                _ = v39
            }
            for fi40 := 0; fi40 < 12; fi40++ {
                v3 -= (v2 ^ (v2 + (v3 * 960)))
                var v41 uint16 = 2
                if ((v3 * (v3 >> ((65535) & 63))) > 42) {
                    for fi42 := 0; fi42 < 15; fi42++ {
                        fmt.Println(v2, v2)
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        v2 += 1000
                        v1 = (v1 * 30)
                        var v43 uint16 = v41
                        _ = fi42
                        _ = v43
                    }
                } else {
                    v41 = (v41 % ((v41) | 1))
                    v2 += 835
                }
                _ = fi40
                _ = v41
            }
        } else {
            v44 := GenS5{F0: 1000}
            _ = v44
        }
        v1 -= v1
    } else {
        var v45 uint32 = 2147483647
        _ = v45
    }
    var v46 int = -507
    var v47 uint8 = 214
    for fi48 := 0; fi48 < 18; fi48++ {
        v1 -= (v47 % (((^(v47))) | 1))
        _ = fi48
    }
    v49 := GenS1{F0: 1, F1: 42, F2: (1 <= 2)}
    if false {
        for fi50 := 0; fi50 < 7; fi50++ {
            switch -474 {
            case -1:
                for fi51 := 0; fi51 < 11; fi51++ {
                    if false {
                        v47 += 76
                    } else {
                        v46 = fi51
                        var v52 float32 = -2.25
                        switch "" {
                        case "foo bar":
                            var v53 int = fi51
                            _ = v53
                        }
                        _ = v52
                    }
                    v1 += v1
                    _ = fi51
                }
            default:
                var v54 float64 = -2.25
                _ = v54
            }
            var v55 uint = uint(v47)
            fmt.Println(v47, v3, fi50)
            _ = fi50
            _ = v55
        }
        var v56 uint32 = 0
        _ = v56
    } else {
        fmt.Println(v1, v46, v46)
        v2 -= -1
    }
    var v57 uint8 = (v47 + v47)
    var v58 int = v46
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v46
    _ = v47
    _ = v49
    _ = v57
    _ = v58
    _ = v8
    _ = v9
}
