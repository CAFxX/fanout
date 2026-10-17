package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint8 = 114
    var v2 uint16 = 0
    var v3 int64 = 9223372036854775807
    var v4 uint64 = uint64(v2)
    var v5 uint8 = (v1 - 42)
    v6 := []int{}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int16 = -533
        v8 = (v8 ^ v8)
        if len(v6) > 0 {
            v6[0] = fi7
        }
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    switch "foo bar" {
    case "":
        var v10 string = ("foo bar" + (((("\n\t" + "foo bar") + (("a" + "a") + ("foo bar" + "a"))) + ("foo bar" + "hello")) + "hello"))
        _ = v10
    case "hello":
        v3 = 9223372036854775807
    case "foo bar":
        fmt.Println(v1, v3)
    default:
        switch "a" {
        case "\n\t":
            fmt.Println(v5)
        }
    }
    if (128 <= 1) {
        switch 42 {
        case 128:
            var v11 int32 = 957
            _ = v11
        case 32767:
            var v12 int8 = 0
            _ = v12
        case 2147483647:
            v5 -= v1
        }
        _ = len(v6)
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        if (!(!(v2 != (v2 - (v2 - 65535))))) {
            fmt.Println(v3, v2, v2)
            fmt.Println(v9, v4, fi13)
        }
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    if false {
        var v15 float32 = 3.14159
        for fi16 := 0; fi16 < 13; fi16++ {
            switch "" {
            case "a":
                switch fi16 {
                case 128:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
            if ((^(v4)) <= 825) {
                v17 := GenS1{F0: (1 < -128)}
                switch ("\n\t" + (("" + "a") + ("\n\t" + "\n\t"))) {
                case "a":
                    var v18 string = "a"
                    _ = v18
                case "\n\t":
                    v5 = 235
                case "foo bar":
                    v19 := GenS1{F0: true}
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 18; fi20++ {
                        if len(v6) > 0 {
                            v6[0] = 32768
                        }
                        fmt.Println(v5, v1)
                        var v21 int64 = (v3 % ((v3) | 1))
                        _ = fi20
                        _ = v21
                    }
                }
                _ = v17
            }
            _ = fi16
        }
        _ = v15
    } else {
        v22 := GenS1{F0: false}
        _ = v22
    }
    switch -99 {
    case 9223372036854775807:
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS1{F0: false}
            if ((false || false) || true) {
                var v25 bool = (-156 <= 2147483647)
                _ = v25
            } else {
                if true {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi26 := 0; fi26 < 18; fi26++ {
                        fmt.Println(v9, v4, fi23)
                        if (v9 > (v4 << ((255) & 63))) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if (0.0 <= (1e10 * 3.14159)) {
                                v2 -= (v2 & 42)
                                fmt.Println(v2)
                            }
                            v9 += 1
                            switch (fi23 % (((fi26 / (((fi23 % ((fi26) | 1))) | 1))) | 1)) {
                            case 4294967295:
                                v2 -= (v2 & uint16(v4))
                                if true {
                                    switch ("\n\t" + "\n\t") {
                                    case "\n\t":
                                        fmt.Println(v3, v3, v3)
                                    }
                                }
                            }
                            var v27 uint64 = v9
                            v28 := GenS1{F0: (!(!false))}
                            _ = v27
                            _ = v28
                        } else {
                            fmt.Println(v4, v5, v9)
                            var v29 uint32 = 42
                            v9 = 217
                            _ = v29
                        }
                        if false {
                            fmt.Println(v3, v2, v1)
                            v4 = (v9 & (v9 & (v9 % (((v4 & v9)) | 1))))
                        }
                        _ = fi26
                    }
                    v5 = 0
                }
                if true {
                    if true {
                        var v30 float32 = (float32(0.0) / float32(3.14159))
                        fmt.Println(v9)
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 uint = 255
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                    if len(v6) > 0 {
                        v6[0] = (fi23 >> ((301) & 63))
                    }
                    v33 := GenS1{F0: false}
                    _ = v33
                } else {
                    if (100 > 100) {
                        if true {
                            var v34 int32 = 379
                            var v35 uint16 = (v2 / (((^(v2))) | 1))
                            var v36 uint8 = v5
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        v37 := GenS1{F0: (1.5 > float32(0.0))}
                        v38 := GenS1{F0: false}
                        _ = v37
                        _ = v38
                    }
                }
            }
            _ = fi23
            _ = v24
        }
    case 32767:
        v9 = (v4 | (v9 << ((4294967295) & 63)))
    }
    switch "foo bar" {
    case "\n\t":
        for fi39 := 0; fi39 < 6; fi39++ {
            v3 -= (^(v3))
            _ = fi39
        }
    default:
        if false {
            if (uint(v2) == 18446744073709551615) {
                v40 := GenS2{F0: (v5 | 128), F1: -9223372036854775808, F2: v4}
                var v41 uint8 = (v5 >> ((uint(v5)) & 63))
                _ = v40
                _ = v41
            } else {
                if (1 == 44) {
                    v42 := GenS3{F0: ("foo bar" + ""), F1: 840}
                    if (0 == 554) {
                        var v43 string = "hello"
                        _ = v43
                    } else {
                        for fi44 := 0; fi44 < 13; fi44++ {
                            for fi45 := 0; fi45 < 7; fi45++ {
                                v6 = append(v6, fi45)
                                var v46 int16 = -32768
                                fmt.Println(v1)
                                _ = fi45
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 11; fi47++ {
                                switch fi47 {
                                case 18:
                                    v5 -= v1
                                case 9223372036854775807:
                                    fmt.Println(fi44)
                                case 32767:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                _ = fi47
                            }
                            _ = fi44
                        }
                        for fi48 := 0; fi48 < 13; fi48++ {
                            v1 = 123
                            _ = fi48
                        }
                        v4 -= (v4 % ((v4) | 1))
                    }
                    _ = v42
                } else {
                    switch ("\n\t" + "a") {
                    case "hello":
                        fmt.Println(v3)
                    case "foo bar":
                        v9 = 0
                    case "a":
                        var v49 uint16 = uint16(v9)
                        _ = v49
                    default:
                        v9 -= (v4 & (^(v9)))
                    }
                }
                _ = len(v6)
                if (true && false) {
                    if (!(!(42 == 90))) {
                        fmt.Println(v5, v1, v3)
                    } else {
                        if (!false) {
                            if (v3 == (v3 + v3)) {
                                v4 -= (v9 / (((v9 + (^(v9)))) | 1))
                                for fi50 := 0; fi50 < 17; fi50++ {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi50
                                }
                                if true {
                                    v4 = v9
                                } else {
                                    v1 += (v5 - v1)
                                    v1 = v5
                                }
                            }
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                var v51 uint = uint(v1)
                                _ = v51
                            case "\n\t":
                                v6 = append(v6, -9223372036854775808)
                            case "hello":
                                for fi52 := 0; fi52 < 20; fi52++ {
                                    v53 := GenS3{F0: "foo bar", F1: 42}
                                    v54 := GenS3{F0: ("\n\t" + ""), F1: 635}
                                    var v55 int64 = (v3 ^ (v3 ^ v3))
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                            }
                            for fi56 := 0; fi56 < 13; fi56++ {
                                if true {
                                    v5 = v5
                                }
                                var v57 uint64 = (v4 ^ v9)
                                _ = fi56
                                _ = v57
                            }
                            var v58 uint32 = 2
                            _ = v58
                        }
                        v59 := GenS4{F0: (2 >= 42), F1: 941}
                        _ = v59
                    }
                } else {
                    v60 := GenS1{F0: true}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi61 := 0; fi61 < 6; fi61++ {
                        for fi62 := 0; fi62 < 5; fi62++ {
                            v63 := GenS3{F0: "", F1: 42}
                            _ = fi62
                            _ = v63
                        }
                        _ = fi61
                    }
                    _ = v60
                }
            }
        }
    }
    v1 -= v5
    if (4294967295 >= 2147483648) {
        if ((v4 & (v4 / (((v4 << ((uint(v2)) & 63))) | 1))) != v9) {
            v4 -= (^(v9))
        }
    }
    var v64 float64 = float64((0.1 / 1.5))
    for fi65 := 0; fi65 < 16; fi65++ {
        fmt.Println(v5, v3, fi65)
        _ = len(v6)
        v66 := GenS4{F0: false, F1: 42}
        _ = fi65
        _ = v66
    }
    v9 += 978
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v9
}
