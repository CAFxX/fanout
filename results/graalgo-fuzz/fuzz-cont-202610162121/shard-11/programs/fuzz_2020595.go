package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 int8; F2 int }
type GenS9 struct { F0 int32 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    var v3 int32 = -2147483648
    var v4 int = 127
    v5 := []uint32{65535, 0, 65535}
    v6 := make(map[uint32]uint32)
    if false {
        _ = len(v2)
    } else {
        var v7 int16 = 32767
        var v8 float32 = 1.5
        switch "" {
        case "\n\t":
            fmt.Println(v7, v7, v3)
        case "foo bar":
            fmt.Println(v7, v3)
        default:
            _ = len(v5)
        }
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: v4, F1: -1, F2: 235}
    v10 := GenS1{F0: (v4 >> ((255) & 63)), F1: 1, F2: 255}
    fmt.Println(v1, v4, v3)
    var v11 int32 = v3
    if (287 == 255) {
        var v12 float32 = 0.0
        if false {
            if false {
                if false {
                    fmt.Println(v4, v11)
                } else {
                    v12 = v12
                    var v13 uint64 = 0
                    switch v4 {
                    case 127:
                        if (true && false) {
                            fmt.Println(v4, v4)
                            v12 = float32(float64(1.5))
                        } else {
                            v14 := GenS1{F0: (v4 * 941), F1: 127, F2: 1}
                            switch 42 {
                            case 2147483647:
                                if (!(100 > -128)) {
                                    var v15 bool = true
                                    switch -113 {
                                    case 2147483647:
                                        fmt.Println(v1, v1)
                                    case -875:
                                        var v16 bool = (809 >= 761)
                                        _ = v16
                                    case 0:
                                        v4 += (v4 | v4)
                                    }
                                    _ = v15
                                }
                            case -442:
                                var v17 int = v4
                                _ = v17
                            case 4294967295:
                                v1 = v1
                            default:
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    if true {
                                        switch (fi18 * v4) {
                                        case 0:
                                            if false {
                                                var v19 float64 = ((float64(1.5) / (-2.25 * ((0.1 - ((3.14159 * -2.25) - v1)) - v1))) + (float64(float32(v1)) + 0.0))
                                                var v20 int8 = 100
                                                _ = v19
                                                _ = v20
                                            }
                                        case 576:
                                            _ = len(v6)
                                        }
                                        v21 := GenS1{F0: v4, F1: 0, F2: 128}
                                        switch v4 {
                                        case -1:
                                            v1 = -2.25
                                        case -9223372036854775808:
                                            if (-1 >= 118) {
                                                for fi22 := 0; fi22 < 1; fi22++ {
                                                    if true {
                                                        v23 := GenS2{F0: -9223372036854775808, F1: 0, F2: 0}
                                                        _ = v23
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = fi22
                                                }
                                                delete(v6, 157)
                                                var v24 int = fi18
                                                _ = v24
                                            }
                                        case 32768:
                                            v3 -= (^(v3))
                                        }
                                        _ = v21
                                    } else {
                                        v11 -= v3
                                        delete(v2, (v4 & fi18))
                                        if true {
                                            var v25 uint8 = 66
                                            if false {
                                                fmt.Println(v11, fi18, v1)
                                                _ = len(v2)
                                            } else {
                                                switch fi18 {
                                                case 9223372036854775807:
                                                    var v26 int16 = 1000
                                                    _ = v26
                                                }
                                            }
                                            _ = v25
                                        } else {
                                            for fi27 := 0; fi27 < 7; fi27++ {
                                                var v28 float32 = v12
                                                _ = fi27
                                                _ = v28
                                            }
                                            v13 += (v13 >> ((65535) & 63))
                                        }
                                    }
                                    fmt.Println(v3, v13, v12)
                                    var v29 bool = (!(!(0 > 1)))
                                    _ = fi18
                                    _ = v29
                                }
                            }
                            v30 := GenS1{F0: v4, F1: 0, F2: 127}
                            _ = v14
                            _ = v30
                        }
                    case 32768:
                        if false {
                            if false {
                                v31 := GenS1{F0: v4, F1: 42, F2: 255}
                                _ = v31
                            } else {
                                if true {
                                    v13 += v13
                                    var v32 uint8 = 42
                                    _ = v32
                                } else {
                                    if true {
                                        fmt.Println(v4, v1)
                                        switch (("a" + "hello") + "foo bar") {
                                        case "hello":
                                            for fi33 := 0; fi33 < 6; fi33++ {
                                                for fi34 := 0; fi34 < 16; fi34++ {
                                                    switch fi34 {
                                                    case -324:
                                                        v13 -= v13
                                                    case 128:
                                                        var v35 uint64 = (v13 - 271)
                                                        _ = v35
                                                    default:
                                                        switch (^(fi33)) {
                                                        case -2147483648:
                                                            v12 = (v12 + (v12 * 1e10))
                                                        case 390:
                                                            if (65535 <= 255) {
                                                                v36 := GenS2{F0: 12, F1: -872, F2: 1}
                                                                var v37 float64 = 1e10
                                                                fmt.Println(fi34)
                                                                _ = v36
                                                                _ = v37
                                                            } else {
                                                                var v38 int16 = -552
                                                                if (false || (false || true)) {
                                                                    fmt.Println(fi34)
                                                                } else {
                                                                    fmt.Println(fi34, v13)
                                                                }
                                                                _ = v38
                                                            }
                                                        }
                                                    }
                                                    switch fi33 {
                                                    case -407:
                                                        switch 127 {
                                                        case -2147483648:
                                                            switch (fi33 >> ((0) & 63)) {
                                                            case 537:
                                                                var v39 uint = 255
                                                                _ = v39
                                                            }
                                                        case -1:
                                                            fmt.Println(v13, v3, v12)
                                                        }
                                                    }
                                                    var v40 uint16 = 1
                                                    v1 -= (v1 * (v1 * v1))
                                                    var v41 uint32 = 2147483647
                                                    _ = fi34
                                                    _ = v40
                                                    _ = v41
                                                }
                                                _ = fi33
                                            }
                                        }
                                    }
                                    for fi42 := 0; fi42 < 1; fi42++ {
                                        v4 = v4
                                        _ = fi42
                                    }
                                    v4 = (^(v4))
                                }
                                v4 = (v4 / (((^(v4))) | 1))
                                for fi43 := 0; fi43 < 9; fi43++ {
                                    v44 := GenS3{F0: v11, F1: true, F2: (9 != 128)}
                                    v45 := v2[v4]
                                    var v46 int32 = (^(v11))
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            }
                            v47 := GenS1{F0: v4, F1: -128, F2: 208}
                            var v48 uint32 = 2147483647
                            _ = v47
                            _ = v48
                        } else {
                            var v49 int16 = 701
                            v50 := GenS2{F0: -1, F1: -1, F2: 817}
                            _ = v49
                            _ = v50
                        }
                    case -325:
                        var v51 float64 = 0.1
                        _ = v51
                    default:
                        fmt.Println(v12)
                    }
                    _ = v13
                }
                _ = len(v5)
            }
        } else {
            if true {
                v4 = v4
                if (false || (!true)) {
                    v11 += v11
                    if (true && (false || true)) {
                        if (v3 != (v11 + -751)) {
                            var v52 float32 = v12
                            switch "\n\t" {
                            case "":
                                for fi53 := 0; fi53 < 20; fi53++ {
                                    var v54 int64 = -1
                                    v55 := GenS3{F0: v3, F1: (-32768 <= int16(v4)), F2: true}
                                    fmt.Println(v4, v54, v1)
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                }
                            case "hello":
                                for fi56 := 0; fi56 < 20; fi56++ {
                                    for fi57 := 0; fi57 < 14; fi57++ {
                                        var v58 float32 = ((0.0 + (float32((v1 + v1)) + float32(((3.14159 / v1) / (((((1e10 * ((v1 - ((v1 + ((v1 + float64((v12 + v52))) / (v1 * v1))) * 1.5)) + 0.0)) - (3.14159 + v1)) / 0.1) + float64((3.14159 + ((((float32(v1) + 1e10) + 3.14159) - 3.14159) + v12)))) - v1))))) + (0.1 * -2.25))
                                        _ = fi57
                                        _ = v58
                                    }
                                    if true {
                                        v11 -= v3
                                        if (!false) {
                                            var v59 uint16 = 32767
                                            _ = v59
                                        } else {
                                            fmt.Println(v11)
                                        }
                                    } else {
                                        var v60 uint16 = 1
                                        _ = v60
                                    }
                                    switch ("hello" + ("foo bar" + "\n\t")) {
                                    case "foo bar":
                                        if false {
                                            var v61 int32 = 577
                                            _ = v61
                                        }
                                    case "a":
                                        if true {
                                            if false {
                                                delete(v2, v4)
                                                v62 := GenS4{F0: 42, F1: ("" + "")}
                                                v63 := GenS1{F0: -1, F1: -128, F2: 128}
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                    }
                                    _ = fi56
                                }
                            }
                            if (-629 < 42) {
                                v64 := GenS2{F0: 42, F1: -1, F2: 18446744073709551615}
                                v1 = 3.14159
                                v65 := GenS5{F0: 2, F1: 831}
                                _ = v64
                                _ = v65
                            }
                            _ = v52
                        }
                    } else {
                        fmt.Println(v4)
                    }
                    var v66 int = v4
                    _ = v66
                }
            } else {
                var v67 float32 = v12
                v12 -= v12
                fmt.Println(v11, v67, v1)
                _ = v67
            }
            var v68 bool = (243 > 0)
            if v68 {
                if (-9223372036854775808 <= 980) {
                    var v69 bool = v68
                    _ = v69
                } else {
                    var v70 int = 1
                    v1 = 3.14159
                    fmt.Println(v3)
                    _ = v70
                }
                switch ("foo bar" + (("a" + "") + "")) {
                case "a":
                    if len(v5) > 0 {
                        v5[0] = 2
                    }
                case "foo bar":
                    v12 = ((1e10 * (v12 + (v12 - 1e10))) - 1.5)
                case "":
                    v71 := GenS4{F0: (^(v3)), F1: "\n\t"}
                    _ = v71
                default:
                    _ = len(v2)
                }
                v5 = append(v5, 2)
            } else {
                var v72 uint = 819
                v3 = 1
                if v68 {
                    if (!(!((v1 > v1) && v68))) {
                        v73 := GenS5{F0: 1, F1: 18446744073709551615}
                        for fi74 := 0; fi74 < 10; fi74++ {
                            v11 = v11
                            v68 = v68
                            switch (fi74 >> (((^(v72))) & 63)) {
                            case 32768:
                                v75 := GenS6{F0: -128}
                                _ = v75
                            case 0:
                                fmt.Println(v3, v3, v1)
                            default:
                                v76 := GenS3{F0: v11, F1: (v68 || (!true)), F2: v68}
                                _ = v76
                            }
                            _ = fi74
                        }
                        _ = v73
                    } else {
                        v77 := GenS7{F0: v11}
                        _ = v77
                    }
                    var v78 uint8 = 97
                    _ = v78
                } else {
                    switch v4 {
                    case -875:
                        switch (^(v4)) {
                        case 32767:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case -1:
                        if true {
                            if true {
                                switch "" {
                                case "":
                                    delete(v2, v4)
                                case "hello":
                                    v79 := GenS8{F0: "", F1: -1, F2: (v4 & (v4 * v4))}
                                    _ = v79
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    v68 = (234 < 0)
                                default:
                                    var v80 bool = (("" <= "foo bar") || v68)
                                    _ = v80
                                }
                            }
                            var v81 int8 = 61
                            if (109 > 0) {
                                switch "a" {
                                case "\n\t":
                                    v72 = (v72 & (v72 ^ 1))
                                case "hello":
                                    var v82 uint8 = 127
                                    _ = v82
                                default:
                                    for fi83 := 0; fi83 < 6; fi83++ {
                                        fmt.Println(v72, v72, v3)
                                        _ = fi83
                                    }
                                }
                                var v84 uint32 = 42
                                _ = v84
                            } else {
                                for fi85 := 0; fi85 < 14; fi85++ {
                                    var v86 uint = (v72 ^ v72)
                                    _ = fi85
                                    _ = v86
                                }
                            }
                            _ = v81
                        }
                    case 4294967295:
                        v87 := GenS4{F0: (-(v3)), F1: "a"}
                        _ = v87
                    default:
                        v11 += v3
                    }
                    var v88 uint64 = 42
                    v11 -= v11
                    _ = v88
                }
                _ = v72
            }
            _ = v68
        }
        v4 = -2147483648
        _ = v12
    } else {
        delete(v2, v4)
        v89 := GenS2{F0: 998, F1: 9223372036854775807, F2: 942}
        _ = v89
    }
    v90 := GenS9{F0: v11}
    if false {
        v2[v4] = 122
        fmt.Println(v11, v4, v3)
    }
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
    _ = v90
}
