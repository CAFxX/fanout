package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int16; F2 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 int8 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint; F1 int; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS9 struct { F0 int64; F1 int32; F2 int64 }
type GenS10 struct { F0 int; F1 uint32; F2 int64 }
func genFn0(p0 uint) uint16 {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{(-(fi4)), fi4, (fi4 + (fi4 + fi4))}
        var v6 uint64 = 665
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint = p0
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 int = 4294967295
    v2 := []uint32{353, 2147483648}
    v3 := make(map[int]string)
    if (!(42 >= -1)) {
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 string = "\n\t"
                switch (^(v1)) {
                case 9223372036854775807:
                    _ = len(v3)
                }
                _ = fi9
                _ = v10
            }
            if false {
                v11 := v3[(^(v1))]
                _ = v11
            }
            _ = fi8
        }
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: 627}
            v2 = append(v2, 2)
            var v15 int64 = -571
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(fi16, fi16)
            var v17 int32 = 2147483647
            _ = fi16
            _ = v17
        }
        _ = fi12
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        if (false && (255 < 32767)) {
            var v19 uint = 65535
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 float64 = (3.14159 * -2.25)
                _ = fi20
                _ = v21
            }
            var v22 float32 = float32(((float64(1.5) + 0.0) + 1.5))
            _ = v19
            _ = v22
        }
        switch (v1 + 32767) {
        case 0:
            for fi23 := 0; fi23 < 16; fi23++ {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v24 int64 = -9223372036854775808
                _ = fi23
                _ = v24
            }
        case 42:
            var v25 float64 = -2.25
            _ = v25
        case -633:
            var v26 int32 = 2147483647
            _ = v26
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi18
    }
    v27 := GenS2{F0: false}
    var v28 float64 = 0.1
    if (2147483648 <= 2147483648) {
        var v29 int = (v1 >> ((42) & 63))
        _ = v29
    } else {
        switch ("foo bar" + (("a" + "") + ("foo bar" + "\n\t"))) {
        case "a":
            for fi30 := 0; fi30 < 8; fi30++ {
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 uint8 = 255
                    _ = fi31
                    _ = v32
                }
                v33 := v3[(fi30 >> ((18446744073709551615) & 63))]
                fmt.Println(v1)
                _ = fi30
                _ = v33
            }
        case "":
            var v34 uint16 = 0
            _ = v34
        }
        switch 432 {
        case 127:
            fmt.Println(v28)
        case 2147483647:
            v1 += (v1 + -439)
        case 32768:
            var v35 uint64 = 4294967295
            _ = v35
        }
    }
    switch "foo bar" {
    case "hello":
        for fi36 := 0; fi36 < 4; fi36++ {
            for fi37 := 0; fi37 < 6; fi37++ {
                if ((true && (0 > 2)) || false) {
                    v38 := GenS3{F0: (!true), F1: -1, F2: 0}
                    var v39 int32 = 42
                    _ = v38
                    _ = v39
                }
                var v40 uint64 = 2
                if true {
                    fmt.Println(v40)
                    var v41 string = "\n\t"
                    var v42 int64 = 763
                    _ = v41
                    _ = v42
                } else {
                    v2 = append(v2, 42)
                    switch "hello" {
                    case "hello":
                        v2 = append(v2, 65535)
                    case "\n\t":
                        if false {
                            v43 := GenS4{F0: ("" + "hello"), F1: 127, F2: 100}
                            var v44 int64 = -9223372036854775808
                            v44 += v44
                            _ = v43
                            _ = v44
                        } else {
                            fmt.Println(v40, v1)
                            switch 32767 {
                            case 42:
                                switch v1 {
                                case -2147483648:
                                    switch (("hello" + "\n\t") + "foo bar") {
                                    case "foo bar":
                                        if false {
                                            if true {
                                                fmt.Println(fi37)
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v45 := GenS1{F0: 264}
                                            _ = v45
                                        } else {
                                            v3[v1] = "hello"
                                            if ((3.14159 != ((1e10 - ((((-2.25 - 0.1) * float32(v28)) / (0.0 + (float32(1.5) / float32((((1e10 + v28) + 3.14159) / v28))))) / -2.25)) + float32(v28))) || (true && (486 != fi36))) {
                                                switch "hello" {
                                                case "\n\t":
                                                    switch "foo bar" {
                                                    case "":
                                                        _ = len(v2)
                                                        v40 -= v40
                                                        if (!true) {
                                                            v3[fi36] = "\n\t"
                                                            delete(v3, fi37)
                                                        }
                                                        fmt.Println(fi36, v1, fi36)
                                                    case "a":
                                                        var v46 uint8 = 42
                                                        _ = v46
                                                    default:
                                                        if ((((-2147483648 == -1) || (!false)) || ((65535 <= 223) && true)) || false) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if false {
                                                                var v47 uint = 0
                                                                _ = v47
                                                            }
                                                        } else {
                                                            if (!false) {
                                                                switch "hello" {
                                                                case "":
                                                                    delete(v3, fi37)
                                                                case "hello":
                                                                    v48 := GenS4{F0: "\n\t", F1: 255, F2: 42}
                                                                    _ = v48
                                                                case "a":
                                                                    v49 := v3[fi37]
                                                                    _ = v49
                                                                }
                                                            }
                                                            fmt.Println(v40, v28)
                                                            v3[fi36] = ("" + "\n\t")
                                                            if (!false) {
                                                                var v50 int16 = -517
                                                                var v51 uint32 = 1
                                                                _ = v50
                                                                _ = v51
                                                            } else {
                                                                var v52 uint64 = (^(v40))
                                                                fmt.Println(v40, v1, v52)
                                                                _ = v52
                                                            }
                                                        }
                                                    }
                                                case "":
                                                    var v53 string = "foo bar"
                                                    _ = v53
                                                case "foo bar":
                                                    v54 := GenS5{F0: false, F1: uint16(v40)}
                                                    _ = v54
                                                default:
                                                    var v55 int32 = 0
                                                    _ = v55
                                                }
                                                v28 = -2.25
                                            }
                                        }
                                    case "hello":
                                        v1 -= v1
                                        v56 := GenS6{F0: 593, F1: fi36, F2: 42}
                                        v1 = (fi36 * v1)
                                        _ = v56
                                    }
                                case 647:
                                    v1 -= (v1 * 9223372036854775807)
                                case 128:
                                    var v57 float64 = v28
                                    _ = v57
                                default:
                                    var v58 int8 = 42
                                    _ = v58
                                }
                            case 2147483647:
                                switch "hello" {
                                case "":
                                    v28 = (v28 * (0.1 - 3.14159))
                                case "hello":
                                    v59 := GenS7{F0: 42, F1: 1, F2: genFn0(655)}
                                    _ = v59
                                case "\n\t":
                                    v40 = v40
                                    _ = len(v2)
                                default:
                                    v1 -= v1
                                }
                            case 32768:
                                if ((1 >= 62) || false) {
                                    switch fi36 {
                                    case -1:
                                        v1 = 32768
                                    case -2147483648:
                                        switch "\n\t" {
                                        case "hello":
                                            var v60 float64 = (float64(1.5) - 1.5)
                                            _ = v60
                                        case "a":
                                            var v61 uint = 0
                                            _ = v61
                                        default:
                                            v1 = fi36
                                        }
                                    case 128:
                                        fmt.Println(fi36)
                                    default:
                                        fmt.Println(v28, v40, v28)
                                    }
                                } else {
                                    v3[-9223372036854775808] = "hello"
                                    fmt.Println(v40)
                                }
                            default:
                                switch "\n\t" {
                                case "\n\t":
                                    v3[923] = "hello"
                                case "":
                                    var v62 bool = (-511 == 1)
                                    _ = v62
                                case "foo bar":
                                    if (!(!(0 != v1))) {
                                        v1 += (fi37 | -10)
                                    }
                                default:
                                    delete(v3, (v1 - 127))
                                }
                            }
                        }
                    case "a":
                        v1 -= -486
                        var v63 int = fi36
                        v40 = 937
                        _ = v63
                    }
                }
                _ = fi37
                _ = v40
            }
            _ = fi36
        }
    case "":
        v1 += (v1 | v1)
    case "foo bar":
        var v64 string = "hello"
        _ = v64
    default:
        var v65 int16 = 434
        _ = v65
    }
    v66 := GenS8{F0: 0, F1: 1, F2: 1}
    if len(v2) > 0 {
        v2[0] = 42
    }
    fmt.Println(v28, v1)
    _ = len(v2)
    var v67 int = (v1 ^ (v1 + -9223372036854775808))
    fmt.Println(v28, v28)
    switch "" {
    case "hello":
        for fi68 := 0; fi68 < 9; fi68++ {
            var v69 int64 = 1
            _ = fi68
            _ = v69
        }
    case "a":
        if true {
            switch (("foo bar" + "") + "\n\t") {
            case "\n\t":
                if false {
                    for fi70 := 0; fi70 < 5; fi70++ {
                        switch -2147483648 {
                        case 0:
                            for fi71 := 0; fi71 < 17; fi71++ {
                                var v72 bool = (false && (0.1 > 1.5))
                                var v73 int64 = -1
                                fmt.Println(v1, v28)
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                        case -1:
                            var v74 int64 = 9223372036854775807
                            _ = v74
                        default:
                            var v75 int32 = 351
                            _ = v75
                        }
                        v28 -= v28
                        var v76 int = v67
                        _ = fi70
                        _ = v76
                    }
                }
            case "a":
                if (false && (false || (v28 <= 0.0))) {
                    fmt.Println(v1, v1)
                    for fi77 := 0; fi77 < 7; fi77++ {
                        v78 := GenS9{F0: -1, F1: -2147483648, F2: 42}
                        v1 = fi77
                        _ = fi77
                        _ = v78
                    }
                    v67 += -9223372036854775808
                } else {
                    if true {
                        if (127 < 69) {
                            for fi79 := 0; fi79 < 6; fi79++ {
                                v80 := GenS1{F0: 85}
                                _ = fi79
                                _ = v80
                            }
                            if ("a" <= "a") {
                                if false {
                                    v1 += (v67 + (v67 + v1))
                                    fmt.Println(v1, v67, v67)
                                    fmt.Println(v67, v28)
                                }
                                var v81 string = "hello"
                                v1 = (v67 + (v67 * (v1 * (v67 & v1))))
                                _ = v81
                            } else {
                                fmt.Println(v1)
                                var v82 bool = true
                                _ = v82
                            }
                        }
                        var v83 int8 = -78
                        _ = v83
                    } else {
                        v1 = v1
                        v84 := GenS9{F0: 9223372036854775807, F1: -1, F2: int64(v1)}
                        if (1.5 == (0.0 - 0.0)) {
                            var v85 uint16 = 1
                            _ = v85
                        } else {
                            v1 = (v1 | 32767)
                        }
                        _ = v84
                    }
                }
            case "hello":
                for fi86 := 0; fi86 < 11; fi86++ {
                    v67 -= v1
                    _ = fi86
                }
            default:
                var v87 int64 = -701
                _ = v87
            }
            for fi88 := 0; fi88 < 2; fi88++ {
                if true {
                    v2 = append(v2, 2)
                }
                _ = fi88
            }
            var v89 uint = 165
            _ = v89
        } else {
            v90 := GenS10{F0: (v1 ^ 4294967295), F1: 0, F2: 541}
            _ = v90
        }
    case "\n\t":
        v1 -= (v67 | -449)
    default:
        for fi91 := 0; fi91 < 11; fi91++ {
            v67 += v1
            v67 = (fi91 << ((299) & 63))
            _ = fi91
        }
    }
    for fi92 := 0; fi92 < 14; fi92++ {
        if (-128 >= 42) {
            var v93 uint = 2
            _ = v93
        }
        switch ((((("a" + "hello") + "hello") + "") + "a") + "hello") {
        case "a":
            for fi94 := 0; fi94 < 3; fi94++ {
                var v95 uint = 0
                _ = fi94
                _ = v95
            }
        default:
            v1 -= 1
        }
        v67 += v1
        _ = fi92
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v66
    _ = v67
}
