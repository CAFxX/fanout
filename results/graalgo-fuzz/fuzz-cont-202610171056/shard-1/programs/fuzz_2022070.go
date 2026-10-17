package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64; F1 int8; F2 int16 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int32; F1 string; F2 int32 }
type GenS7 struct { F0 uint64; F1 int8; F2 int16 }
type GenS8 struct { F0 string; F1 uint; F2 uint64 }
type GenS9 struct { F0 int64; F1 string; F2 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64; F1 uint32 }
type GenS12 struct { F0 uint64; F1 string; F2 uint8 }
func genFn0() int16 {
    var v5 int32 = 1
    var v6 int64 = 1
    v7 := []int{4294967295, 42, -731, 4294967295}
    v7 = append(v7, 735)
    _ = v5
    _ = v6
    _ = v7
    return 120
}

func main() {
    var v1 bool = (true && false)
    v2 := []uint32{2147483647, 874, 65535, 0}
    v3 := make(map[int]int)
    var v4 bool = true
    if len(v2) > 0 {
        v2[0] = 482
    }
    if (2147483648 > 2147483648) {
        v8 := GenS1{F0: 377}
        var v9 int8 = -1
        if v4 {
            switch "a" {
            case "a":
                var v10 float64 = float64(-2.25)
                _ = v10
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v4 = (1.5 == float64((-2.25 - (0.1 / 1e10))))
        } else {
            v11 := GenS1{F0: 122}
            v12 := GenS1{F0: 152}
            _ = v11
            _ = v12
        }
        _ = v8
        _ = v9
    } else {
        var v13 int16 = genFn0()
        _ = v13
    }
    switch 2147483647 {
    case 32767:
        var v14 float64 = (1e10 + float64(float32(-2.25)))
        _ = v14
    case -55:
        switch 2147483647 {
        case -309:
            if v1 {
                v1 = true
            } else {
                if false {
                    fmt.Println(v4, v4)
                    for fi15 := 0; fi15 < 15; fi15++ {
                        if true {
                            fmt.Println(fi15, v1, fi15)
                        }
                        _ = fi15
                    }
                    var v16 uint8 = 192
                    _ = v16
                }
                v4 = v1
                if false {
                    var v17 bool = false
                    _ = v17
                } else {
                    fmt.Println(v1, v4)
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v4 = (!false)
                        var v19 int64 = 1
                        _ = fi18
                        _ = v19
                    }
                    _ = len(v2)
                }
            }
        case 127:
            for fi20 := 0; fi20 < 2; fi20++ {
                switch "\n\t" {
                case "":
                    v21 := GenS2{F0: 1, F1: 127}
                    _ = v21
                case "foo bar":
                    v22 := GenS3{F0: -1, F1: -53, F2: 0}
                    _ = v22
                case "hello":
                    var v23 string = "hello"
                    _ = v23
                }
                delete(v3, fi20)
                _ = fi20
            }
        case 1:
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 uint8 = 2
                var v26 uint = 1
                _ = fi24
                _ = v25
                _ = v26
            }
        }
    case 128:
        var v27 uint64 = 112
        _ = v27
    default:
        if v1 {
            var v28 float32 = 0.0
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 uint64 = 1
                if false {
                    var v31 uint16 = 255
                    v28 -= v28
                    v32 := GenS1{F0: (^(v30))}
                    _ = v31
                    _ = v32
                } else {
                    if false {
                        for fi33 := 0; fi33 < 5; fi33++ {
                            if (!(!v1)) {
                                var v34 int32 = -2147483648
                                if (0 < 0) {
                                    var v35 int64 = -9223372036854775808
                                    _ = v35
                                }
                                fmt.Println(v4, v34)
                                _ = v34
                            }
                            _ = fi33
                        }
                        if (!(!(2147483648 < 2147483648))) {
                            switch 32767 {
                            case 1:
                                fmt.Println(v1, v4)
                            case 127:
                                for fi36 := 0; fi36 < 16; fi36++ {
                                    switch (fi29 / ((fi36) | 1)) {
                                    case 4294967295:
                                        v3[42] = fi29
                                    }
                                    _ = fi36
                                }
                            default:
                                v4 = (false && v1)
                            }
                            for fi37 := 0; fi37 < 2; fi37++ {
                                if false {
                                    v28 = float32(float64(v28))
                                    var v38 uint = 255
                                    v28 += v28
                                    _ = v38
                                }
                                fmt.Println(v30, v30)
                                v39 := GenS4{F0: 1, F1: -685}
                                switch "a" {
                                case "":
                                    var v40 uint8 = 238
                                    _ = v40
                                case "hello":
                                    var v41 int32 = 460
                                    _ = v41
                                default:
                                    if true {
                                        if v4 {
                                            v1 = (!(97 > 42))
                                        } else {
                                            v42 := GenS5{F0: genFn0(), F1: 1}
                                            _ = v42
                                        }
                                        fmt.Println(fi37, v1)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch fi37 {
                                        case 32768:
                                            fmt.Println(fi29, v28)
                                        case -1:
                                            var v43 float32 = (v28 - ((v28 / v28) + v28))
                                            _ = v43
                                        case 179:
                                            if (((!(18446744073709551615 == 0)) && v1) && (v1 || ((true && (2 < uint16(v30))) || false))) {
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(fi29)
                                                    if ((69 < int8(fi37)) && ((true || v1) || (0 > 274))) {
                                                        var v44 uint64 = (^(v30))
                                                        v45 := GenS6{F0: 1, F1: "", F2: 42}
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(fi37)
                                                    }
                                                    var v46 uint16 = 2
                                                    _ = v46
                                                case "a":
                                                    var v47 int8 = -128
                                                    _ = v47
                                                case "":
                                                    _ = len(v3)
                                                default:
                                                    v48 := GenS3{F0: 179, F1: -1, F2: 926}
                                                    _ = v48
                                                }
                                            }
                                        default:
                                            switch -1 {
                                            case 842:
                                                v49 := GenS6{F0: 1, F1: "foo bar", F2: -785}
                                                _ = v49
                                            case 4294967295:
                                                fmt.Println(fi37)
                                            default:
                                                var v50 int16 = 1
                                                _ = v50
                                            }
                                        }
                                    } else {
                                        switch int(v30) {
                                        case 4294967295:
                                            var v51 int8 = 23
                                            _ = v51
                                        case -36:
                                            v28 = v28
                                        case -2147483648:
                                            var v52 float32 = (v28 - float32(-2.25))
                                            _ = v52
                                        default:
                                            v4 = ((v30 << ((2) & 63)) != v30)
                                        }
                                        if (!(!v1)) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v30, v1)
                                            v28 = (float32(3.14159) + 1e10)
                                            fmt.Println(v1, v30)
                                            v28 = -2.25
                                            v53 := GenS7{F0: (v30 * v30), F1: 42, F2: 119}
                                            _ = v53
                                        }
                                    }
                                }
                                if v4 {
                                    switch "" {
                                    case "\n\t":
                                        var v54 uint64 = v30
                                        _ = v54
                                    }
                                    if ((v30 - v30) != (v30 / (((v30 / ((4294967295) | 1))) | 1))) {
                                        v55 := GenS5{F0: -1, F1: 0}
                                        v56 := GenS2{F0: 2147483648, F1: 255}
                                        _ = v55
                                        _ = v56
                                    }
                                    v30 += (^(v30))
                                }
                                _ = fi37
                                _ = v39
                            }
                            fmt.Println(v30)
                        } else {
                            if v1 {
                                if (!(((3.14159 + 1e10) < 0.0) || v4)) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        v57 := GenS8{F0: ("\n\t" + "foo bar"), F1: 1, F2: 0}
                                        _ = v57
                                    case "hello":
                                        var v58 uint64 = (v30 % (((v30 - v30)) | 1))
                                        _ = v58
                                    case "a":
                                        var v59 uint16 = 2
                                        _ = v59
                                    default:
                                        v60 := GenS4{F0: 42, F1: 9223372036854775807}
                                        _ = v60
                                    }
                                    v61 := GenS4{F0: -54, F1: 42}
                                    if (!false) {
                                        var v62 uint8 = 1
                                        var v63 float32 = v28
                                        v62 = v62
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v61
                                }
                                v2 = append(v2, 2147483647)
                            }
                            fmt.Println(v1)
                            v28 = (v28 / (-2.25 / v28))
                        }
                        if (42 > 604) {
                            v64 := GenS7{F0: v30, F1: 100, F2: 1}
                            _ = v64
                        }
                    }
                    fmt.Println(v4)
                }
                for fi65 := 0; fi65 < 4; fi65++ {
                    var v66 uint32 = 2147483647
                    _ = fi65
                    _ = v66
                }
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            switch (("a" + ("a" + "a")) + ("hello" + "a")) {
            case "foo bar":
                _ = len(v2)
            case "hello":
                var v67 int = 0
                _ = v67
            case "":
                v1 = (!false)
            default:
                switch "a" {
                case "":
                    var v68 uint16 = 255
                    _ = v68
                case "foo bar":
                    for fi69 := 0; fi69 < 18; fi69++ {
                        var v70 int32 = 0
                        _ = fi69
                        _ = v70
                    }
                case "hello":
                    v1 = (!v4)
                default:
                    v1 = v1
                }
            }
            if (65535 != 32767) {
                for fi71 := 0; fi71 < 18; fi71++ {
                    switch "hello" {
                    case "\n\t":
                        if (!true) {
                            switch fi71 {
                            case 2147483647:
                                switch "a" {
                                case "a":
                                    var v72 uint8 = 255
                                    _ = v72
                                default:
                                    var v73 uint8 = 127
                                    _ = v73
                                }
                            case -2147483648:
                                v1 = v1
                            case 0:
                                var v74 uint32 = 316
                                _ = v74
                            }
                            var v75 uint32 = 42
                            _ = v75
                        } else {
                            if ((!(1 != 18446744073709551615)) && false) {
                                if ("" >= (("" + "hello") + (("foo bar" + "hello") + "hello"))) {
                                    switch (fi71 & fi71) {
                                    case -9223372036854775808:
                                        v1 = false
                                    default:
                                        fmt.Println(v1)
                                    }
                                    var v76 int8 = -1
                                    _ = v76
                                }
                            } else {
                                v77 := GenS9{F0: 254, F1: "", F2: 384}
                                var v78 bool = ((((v1 || (0 == 0)) && true) || (!(2 > 0))) && (((0.1 / ((3.14159 / (1e10 / (1e10 * 1e10))) * float32(0.1))) + 1e10) >= 1.5))
                                _ = v77
                                _ = v78
                            }
                            var v79 uint8 = 42
                            _ = v79
                        }
                    default:
                        var v80 int8 = 100
                        _ = v80
                    }
                    for fi81 := 0; fi81 < 15; fi81++ {
                        var v82 uint16 = 2
                        v4 = (3.14159 > 1e10)
                        _ = fi81
                        _ = v82
                    }
                    var v83 uint64 = 1
                    _ = fi71
                    _ = v83
                }
                v4 = false
                var v84 string = "\n\t"
                _ = v84
            }
            v85 := GenS10{F0: 0}
            _ = v85
        }
    }
    var v86 float64 = 3.14159
    for fi87 := 0; fi87 < 13; fi87++ {
        var v88 uint = 293
        v4 = (0.0 > v86)
        _ = fi87
        _ = v88
    }
    v1 = (!v4)
    for fi89 := 0; fi89 < 6; fi89++ {
        var v90 int32 = -250
        if ((fi89 <= -1) && (v4 && (v86 != (v86 * ((((3.14159 * v86) / v86) + (1e10 * v86)) * v86))))) {
            var v91 bool = v1
            _ = v91
        } else {
            v92 := v3[9223372036854775807]
            v90 -= (^(v90))
            v1 = false
            _ = v92
        }
        v4 = (!true)
        _ = fi89
        _ = v90
    }
    v86 = 1.5
    v1 = true
    for fi93 := 0; fi93 < 11; fi93++ {
        v4 = (v4 || (!(fi93 < (fi93 >> ((1) & 63)))))
        _ = fi93
    }
    switch ("a" + "a") {
    case "hello":
        if v4 {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        }
    case "a":
        if true {
            v94 := GenS11{F0: 946, F1: 1}
            v95 := GenS12{F0: 1, F1: ("" + ""), F2: 242}
            v1 = v4
            _ = v94
            _ = v95
        } else {
            var v96 int64 = 42
            fmt.Println(v86, v96)
            _ = v96
        }
    case "\n\t":
        var v97 uint8 = 1
        _ = v97
    default:
        var v98 uint16 = 248
        _ = v98
    }
    var v99 uint8 = 246
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v86:", v86)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v86
    _ = v99
}
