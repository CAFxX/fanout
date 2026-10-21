package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := []int{-1, 32768, 32767}
    var v4 float32 = 0.0
    v5 := []string{"foo bar", "hello", "", ""}
    var v6 int = 0
    switch v6 {
    case 42:
        v7 := v1[814]
        _ = v7
    case 9223372036854775807:
        for fi8 := 0; fi8 < 6; fi8++ {
            switch (fi8 & v6) {
            case 670:
                var v9 int8 = 53
                _ = v9
            case 812:
                v6 += fi8
            case -949:
                if true {
                    var v10 int32 = -1
                    fmt.Println(v10, v6)
                    if (73 > 42) {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            var v12 int16 = 42
                            var v13 uint8 = 106
                            delete(v2, 65535)
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v6 = fi8
                            v15 := GenS1{F0: (fi14 + fi8), F1: true, F2: 42}
                            v16 := GenS1{F0: -194, F1: (true && false), F2: 501}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        var v17 uint = 255
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 uint32 = 824
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    }
                    _ = v10
                } else {
                    if true {
                        fmt.Println(v6, v4)
                    } else {
                        v20 := GenS1{F0: 4294967295, F1: (3.14159 == 0.0), F2: 1}
                        fmt.Println(fi8, v6)
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if true {
                            var v22 int16 = -1
                            switch (("" + ("a" + "hello")) + (("a" + ("a" + ((("a" + ("" + ("" + "a"))) + ((("hello" + "hello") + "") + ("\n\t" + "a"))) + ""))) + "foo bar")) {
                            case "\n\t":
                                v6 -= fi21
                            }
                            _ = v22
                        } else {
                            v6 = (fi8 / (((fi21 << ((0) & 63))) | 1))
                        }
                        v6 -= (v6 & fi8)
                        _ = fi21
                    }
                }
            default:
                switch (-(v6)) {
                case 0:
                    var v23 int = (^(fi8))
                    _ = v23
                default:
                    v3 = append(v3, v6)
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v4, v6)
            _ = fi8
        }
    case 128:
        switch "\n\t" {
        case "a":
            if true {
                if false {
                    v6 -= (v6 ^ 53)
                    _ = len(v2)
                    if len(v5) > 0 {
                        v5[0] = (("\n\t" + ("" + "a")) + "a")
                    }
                }
            }
        case "\n\t":
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 string = "a"
                _ = fi24
                _ = v25
            }
        case "foo bar":
            var v26 uint64 = 4294967295
            _ = v26
        }
    default:
        v27 := GenS1{F0: (v6 << ((0) & 63)), F1: true, F2: 1}
        _ = v27
    }
    _ = len(v1)
    switch (v6 + v6) {
    case 32768:
        switch "foo bar" {
        case "foo bar":
            if false {
                v6 = v6
                var v28 uint = 65535
                switch v6 {
                case 2147483647:
                    var v29 bool = true
                    _ = v29
                case 101:
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v4 = v4
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        _ = fi30
                    }
                case -317:
                    var v31 uint16 = 226
                    _ = v31
                }
                _ = v28
            }
        default:
            var v32 bool = (!false)
            _ = v32
        }
    case -9223372036854775808:
        var v33 uint64 = 1
        _ = v33
    case 127:
        var v34 uint64 = 1
        _ = v34
    default:
        var v35 uint8 = 224
        _ = v35
    }
    fmt.Println(v4, v6)
    switch ((("hello" + ("foo bar" + "a")) + "a") + "") {
    case "hello":
        fmt.Println(v6)
    case "\n\t":
        var v36 uint8 = 127
        _ = v36
    case "foo bar":
        v4 = float32(0.0)
    default:
        for fi37 := 0; fi37 < 17; fi37++ {
            if (1 <= 1000) {
                for fi38 := 0; fi38 < 5; fi38++ {
                    if false {
                        var v39 uint16 = 255
                        var v40 int16 = 1000
                        v40 += v40
                        v41 := GenS1{F0: fi38, F1: (-381 != v40), F2: v39}
                        _ = v39
                        _ = v40
                        _ = v41
                    } else {
                        v42 := GenS1{F0: (v6 << ((42) & 63)), F1: false, F2: 2}
                        fmt.Println(fi37, v6, v6)
                        switch "\n\t" {
                        case "hello":
                            if (false || false) {
                                switch (-(fi38)) {
                                case -9223372036854775808:
                                    if (0 > 0) {
                                        v43 := GenS2{F0: 0, F1: 4294967295}
                                        _ = v43
                                    } else {
                                        var v44 string = "foo bar"
                                        _ = len(v5)
                                        switch fi37 {
                                        case 568:
                                            switch (fi37 << ((2) & 63)) {
                                            case 32767:
                                                var v45 int64 = 1
                                                _ = v45
                                            case -1:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            default:
                                                var v46 int64 = -1
                                                _ = v46
                                            }
                                        case 714:
                                            v4 += v4
                                        case 9223372036854775807:
                                            v44 = ""
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                v47 := GenS1{F0: (fi37 | v6), F1: false, F2: 65535}
                                                _ = v47
                                            case "\n\t":
                                                v4 = v4
                                                v6 = fi37
                                                fmt.Println(v4, v6, v44)
                                            }
                                        }
                                        _ = v44
                                    }
                                default:
                                    var v48 int32 = 2147483647
                                    _ = v48
                                }
                            } else {
                                if (!(!true)) {
                                    fmt.Println(fi38, v6)
                                    v4 = v4
                                    switch fi38 {
                                    case 1:
                                        var v49 float64 = 3.14159
                                        _ = v49
                                    default:
                                        var v50 int16 = -725
                                        _ = v50
                                    }
                                } else {
                                    v6 = 2147483647
                                }
                                v4 = 1e10
                                v6 -= -920
                            }
                        case "\n\t":
                            var v51 int16 = 0
                            _ = v51
                        case "a":
                            v2[0] = 398
                        }
                        _ = v42
                    }
                    switch "a" {
                    case "":
                        if true {
                            var v52 uint8 = 127
                            if (-1 <= 1000) {
                                var v53 int64 = -9223372036854775808
                                _ = v53
                            }
                            _ = v52
                        } else {
                            v4 -= v4
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = (fi37 % (((fi38 ^ (^(fi37)))) | 1))
                        }
                    }
                    _ = fi38
                }
                if (!true) {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v6, v6)
                    case "a":
                        var v54 int = v6
                        _ = v54
                    case "\n\t":
                        for fi55 := 0; fi55 < 11; fi55++ {
                            fmt.Println(v6, v4)
                            v56 := GenS2{F0: 2, F1: 4294967295}
                            var v57 uint16 = 255
                            v6 -= v6
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    }
                    switch v6 {
                    case -495:
                        v4 = (v4 - v4)
                    default:
                        if true {
                            _ = len(v2)
                            v6 = 42
                            if ((9223372036854775807 < -792) || true) {
                                if true {
                                    fmt.Println(fi37, v6, fi37)
                                    var v58 uint64 = 79
                                    _ = v58
                                }
                                var v59 int32 = 42
                                v59 = 2147483647
                                _ = v59
                            }
                        }
                    }
                    if true {
                        v4 = v4
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v60 string = ("hello" + "\n\t")
                    _ = v60
                }
                v6 += 0
            }
            _ = fi37
        }
    }
    v4 = ((0.1 + (v4 + ((float32(-2.25) * (3.14159 / v4)) - (((((float32(1.5) - v4) + v4) / (v4 - ((v4 / 1e10) + v4))) + (1.5 / -2.25)) * v4)))) / v4)
    switch (("hello" + "hello") + "a") {
    case "hello":
        v61 := GenS3{F0: 880}
        _ = v61
    default:
        switch v6 {
        case 0:
            var v62 int8 = 0
            _ = v62
        }
    }
    _ = len(v5)
    v6 = v6
    if false {
        v6 += 42
        fmt.Println(v6, v6)
    }
    if (float64(0.0) <= 1.5) {
        if false {
            if ((65535 <= 862) || (0.1 < (1.5 * -2.25))) {
                var v63 int = 9223372036854775807
                v64 := GenS1{F0: (^(v63)), F1: false, F2: 460}
                _ = v63
                _ = v64
            } else {
                delete(v2, 256)
                for fi65 := 0; fi65 < 14; fi65++ {
                    fmt.Println(v6, fi65)
                    v4 = (v4 / v4)
                    _ = fi65
                }
            }
            for fi66 := 0; fi66 < 7; fi66++ {
                var v67 uint8 = 42
                _ = fi66
                _ = v67
            }
            if len(v5) > 0 {
                v5[0] = ("hello" + ("" + "\n\t"))
            }
        }
        var v68 uint64 = 2
        _ = v68
    } else {
        if (!true) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if false {
                for fi69 := 0; fi69 < 17; fi69++ {
                    var v70 int64 = 608
                    if (32767 != -32768) {
                        v70 = v70
                    } else {
                        var v71 uint32 = 448
                        _ = v71
                    }
                    fmt.Println(fi69, fi69)
                    _ = fi69
                    _ = v70
                }
            }
            fmt.Println(v6)
        }
    }
    var v72 float32 = (v4 / v4)
    var v73 float32 = v4
    v6 -= 2147483647
    if false {
        for fi74 := 0; fi74 < 4; fi74++ {
            for fi75 := 0; fi75 < 4; fi75++ {
                v76 := GenS4{F0: 0, F1: -706}
                _ = fi75
                _ = v76
            }
            v73 = 0.0
            _ = fi74
        }
        v4 -= ((0.0 / v4) * (v4 + (v73 - ((0.1 * 0.1) - 1.5))))
        var v77 bool = (130 == 169)
        _ = v77
    } else {
        v78 := GenS5{F0: 0, F1: 42}
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
}
