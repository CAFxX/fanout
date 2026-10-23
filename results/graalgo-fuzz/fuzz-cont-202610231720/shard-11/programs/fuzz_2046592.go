package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int32; F1 string; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint8; F1 uint8; F2 string }
type GenS8 struct { F0 int16 }
func genFn0(p0 uint32) int8 {
    var v6 uint64 = 0
    fmt.Println(p0, v6)
    var v7 uint8 = 255
    var v8 uint8 = 255
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 100
}
func genFn1(p0 uint32, p1 uint32, p2 uint16) int8 {
    v9 := GenS1{F0: 43, F1: 0, F2: 32768}
    p0 += (p0 << ((42) & 63))
    switch 0 {
    case 128:
        var v10 uint8 = 243
        _ = v10
    case 42:
        p0 = p1
    case -179:
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := []int{0, -9223372036854775808, -2147483648, -9223372036854775808}
    var v2 uint = 18446744073709551615
    var v3 bool = (-52 <= 0)
    var v4 uint16 = 329
    var v5 int8 = 1
    for fi11 := 0; fi11 < 8; fi11++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi11
    }
    switch "\n\t" {
    case "foo bar":
        var v12 float64 = float64(0.0)
        _ = v12
    default:
        fmt.Println(v4, v3)
    }
    v5 = genFn0(uint32(v4))
    var v13 int16 = 1
    var v14 uint = v2
    var v15 uint32 = 42
    if true {
        var v16 int8 = genFn1(0, (v15 + 734), (v4 << ((v14) & 63)))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 9; fi18++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v19 int64 = -313
                _ = fi18
                _ = v19
            }
            fmt.Println(v15)
            _ = len(v1)
            _ = fi17
        }
        var v20 int = 1
        var v21 int16 = (v13 / (((v13 * int16(v14))) | 1))
        _ = v20
        _ = v21
    }
    v13 = (v13 + v13)
    switch ("" + "\n\t") {
    case "\n\t":
        var v22 uint16 = (v4 - v4)
        _ = v22
    case "":
        v5 = 127
    }
    v13 += v13
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 int8 = (v5 << ((v2) & 63))
        var v25 uint16 = (^(v4))
        _ = fi23
        _ = v24
        _ = v25
    }
    if v3 {
        v26 := GenS2{F0: -32768}
        v1 = append(v1, 32767)
        _ = v26
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = -836
    }
    if false {
        var v27 uint8 = 128
        _ = v27
    }
    if (!true) {
        var v28 bool = (255 > 506)
        _ = v28
    } else {
        switch int(v13) {
        case -841:
            for fi29 := 0; fi29 < 19; fi29++ {
                v4 -= (v4 + (v4 + (v4 - 724)))
                v1 = append(v1, fi29)
                var v30 uint16 = v4
                _ = fi29
                _ = v30
            }
        case -753:
            switch -9223372036854775808 {
            case 128:
                if (129 <= 0) {
                    var v31 int16 = v13
                    var v32 int8 = genFn1(v15, v15, 2)
                    _ = v31
                    _ = v32
                }
            default:
                v3 = v3
            }
        case 127:
            for fi33 := 0; fi33 < 3; fi33++ {
                for fi34 := 0; fi34 < 9; fi34++ {
                    if true {
                        if (v4 < v4) {
                            if ((^(v15)) >= (v15 - uint32(v4))) {
                                var v35 uint8 = 128
                                _ = v35
                            } else {
                                var v36 float64 = 0.0
                                if false {
                                    v13 = int16(v2)
                                    if (((!false) || (!(42 != (-(fi33))))) && (genFn0(v15) <= 127)) {
                                        var v37 uint16 = v4
                                        v14 = 223
                                        fmt.Println(v3, fi34, v13)
                                        _ = v37
                                    } else {
                                        v38 := GenS1{F0: 255, F1: genFn0((v15 & v15)), F2: 4294967295}
                                        v39 := GenS2{F0: v13}
                                        _ = v38
                                        _ = v39
                                    }
                                    v13 -= (v13 / ((v13) | 1))
                                    v2 += (v14 | v14)
                                }
                                v15 -= 564
                                var v40 int16 = v13
                                _ = v36
                                _ = v40
                            }
                            if v3 {
                                var v41 int = (fi34 << ((v2) & 63))
                                switch ("\n\t" + ("hello" + "hello")) {
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = 1
                                    }
                                }
                                var v42 float32 = (-2.25 + 1e10)
                                v1 = append(v1, fi33)
                                _ = v41
                                _ = v42
                            } else {
                                v43 := GenS1{F0: 2, F1: 127, F2: fi33}
                                var v44 uint16 = (^(v4))
                                _ = v43
                                _ = v44
                            }
                            v5 = (v5 % ((v5) | 1))
                        } else {
                            v3 = true
                            v5 = int8(v13)
                        }
                        v2 -= (v14 * (v2 + 247))
                    } else {
                        v3 = true
                        if (847 <= v4) {
                            if (162 != 163) {
                                v4 += (v4 & v4)
                                if false {
                                    var v45 int32 = 2147483647
                                    _ = v45
                                } else {
                                    var v46 uint64 = 2
                                    _ = v46
                                }
                                v14 -= 1
                            } else {
                                switch ("hello" + "a") {
                                case "a":
                                    v3 = v3
                                    v13 = (v13 & v13)
                                    var v47 uint8 = 1
                                    if v3 {
                                        if v3 {
                                            switch int(v14) {
                                            case 0:
                                                v4 += 1
                                            case -331:
                                                var v48 bool = ((v47 == uint8(v4)) || (v5 <= genFn0((v15 << (((v2 << (((v14 << ((665) & 63))) & 63))) & 63)))))
                                                v49 := GenS1{F0: 0, F1: -1, F2: (fi33 % ((1) | 1))}
                                                v2 -= uint(v4)
                                                v1 = append(v1, -71)
                                                _ = v48
                                                _ = v49
                                            case 773:
                                                var v50 uint = 4294967295
                                                _ = v50
                                            default:
                                                v13 = (v13 << (((v2 << (((v2 * (^(v14)))) & 63))) & 63))
                                            }
                                            v1 = append(v1, fi33)
                                            var v51 uint = (^(v14))
                                            if (v14 > v14) {
                                                var v52 uint16 = v4
                                                if v3 {
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi34 ^ fi34)
                                                    }
                                                } else {
                                                    fmt.Println(v47, fi34)
                                                    if (0.0 > 0.0) {
                                                        if false {
                                                            if false {
                                                                if (!(!((^(v51)) <= (v2 << ((v14) & 63))))) {
                                                                    v53 := GenS2{F0: (^(v13))}
                                                                    _ = v53
                                                                }
                                                                var v54 uint64 = 4294967295
                                                                _ = v54
                                                            }
                                                            v5 += (-(v5))
                                                            var v55 int32 = 1
                                                            _ = v55
                                                        } else {
                                                            var v56 string = "hello"
                                                            var v57 int = (fi33 << (((v51 | 255)) & 63))
                                                            v15 += 730
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        v58 := GenS3{F0: -1, F1: true}
                                                        if v3 {
                                                            if (2 != (v51 - (v14 | v14))) {
                                                                if (v3 || ((v52 ^ 445) > 0)) {
                                                                    if (!(!(((0.0 + 1.5) * 1e10) >= -2.25))) {
                                                                        var v59 int64 = 1
                                                                        var v60 int64 = (v59 | (v59 | (^(v59))))
                                                                        v4 = v52
                                                                        _ = v59
                                                                        _ = v60
                                                                    } else {
                                                                        v61 := GenS4{F0: (v5 >> ((v2) & 63)), F1: (v52 | (v4 & 42))}
                                                                        _ = v61
                                                                    }
                                                                    fmt.Println(v52, v51)
                                                                    switch ("a" + "hello") {
                                                                    case "\n\t":
                                                                        fmt.Println(v52)
                                                                    case "":
                                                                        var v62 float64 = ((0.0 * 0.0) - 1e10)
                                                                        _ = v62
                                                                    }
                                                                    var v63 uint8 = uint8(v52)
                                                                    v13 = (v13 * 836)
                                                                    v47 = v63
                                                                    _ = v63
                                                                } else {
                                                                    var v64 uint16 = v52
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "":
                                                                        _ = len(v1)
                                                                    default:
                                                                        switch ("" + "\n\t") {
                                                                        case "a":
                                                                            if v3 {
                                                                                var v65 uint = (v2 & v51)
                                                                                fmt.Println(v52)
                                                                                _ = v65
                                                                            } else {
                                                                                var v66 int64 = -351
                                                                                _ = v66
                                                                            }
                                                                        default:
                                                                            _ = len(v1)
                                                                        }
                                                                    }
                                                                    v52 = v64
                                                                    var v67 string = "a"
                                                                    fmt.Println(v2, v2, v4)
                                                                    _ = v64
                                                                    _ = v67
                                                                }
                                                            }
                                                        } else {
                                                            v2 = (v51 - (v14 & 255))
                                                        }
                                                        if v3 {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                switch "" {
                                                                case "foo bar":
                                                                    v14 = (v2 << (((v2 / ((v2) | 1))) & 63))
                                                                case "hello":
                                                                    v68 := GenS5{F0: 1, F1: ("a" + "hello"), F2: v47}
                                                                    _ = v68
                                                                default:
                                                                    v3 = v3
                                                                }
                                                            case "a":
                                                                if true {
                                                                    switch "" {
                                                                    case "":
                                                                        fmt.Println(v13)
                                                                    case "hello":
                                                                        _ = len(v1)
                                                                    case "a":
                                                                        var v69 int64 = 42
                                                                        _ = v69
                                                                    default:
                                                                        v2 = (v2 & v14)
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            var v70 uint64 = 2
                                                                            _ = v70
                                                                        case "hello":
                                                                            v71 := GenS6{F0: 4294967295, F1: v5}
                                                                            _ = v71
                                                                        default:
                                                                            v15 = v15
                                                                        }
                                                                        if v3 {
                                                                            switch fi34 {
                                                                            case -9223372036854775808:
                                                                                switch 9223372036854775807 {
                                                                                case -9223372036854775808:
                                                                                    switch ("foo bar" + ("" + (("hello" + ((("hello" + ("a" + ("hello" + ""))) + ("\n\t" + "hello")) + "")) + (("foo bar" + "\n\t") + "hello")))) {
                                                                                    case "a":
                                                                                        var v72 float32 = (-2.25 / -2.25)
                                                                                        _ = v72
                                                                                    default:
                                                                                        if v3 {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = fi33
                                                                                            }
                                                                                            v73 := GenS7{F0: 52, F1: uint8(v14), F2: ""}
                                                                                            _ = v73
                                                                                        }
                                                                                    }
                                                                                case 481:
                                                                                    var v74 int64 = 892
                                                                                    _ = v74
                                                                                }
                                                                            default:
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    var v75 uint64 = uint64(v52)
                                                                                    fmt.Println(v14)
                                                                                    fmt.Println(v3, v2)
                                                                                    _ = v75
                                                                                case "foo bar":
                                                                                    fmt.Println(v14, v2, fi33)
                                                                                default:
                                                                                    var v76 uint32 = (v15 >> ((v51) & 63))
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                            if ((!(fi33 < 36)) && (fi34 >= fi34)) {
                                                                                _ = len(v1)
                                                                                v3 = (!true)
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    switch 32768 {
                                                                                    case 741:
                                                                                        var v77 int64 = 0
                                                                                        _ = v77
                                                                                    case 9223372036854775807:
                                                                                        v78 := GenS8{F0: (v13 / ((877) | 1))}
                                                                                        _ = v78
                                                                                    case 4294967295:
                                                                                        _ = len(v1)
                                                                                    }
                                                                                case "\n\t":
                                                                                    v2 = (v14 << (((v51 - 0)) & 63))
                                                                                case "foo bar":
                                                                                    var v79 int32 = 672
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                            var v80 int16 = (v13 + v13)
                                                                            _ = v80
                                                                        } else {
                                                                            var v81 uint32 = (v15 * 2147483647)
                                                                            v81 = 2
                                                                            if ((((-2.25 * 1e10) + 0.0) + ((1.5 - float64(0.0)) * -2.25)) == 0.1) {
                                                                                v47 += uint8(v81)
                                                                            } else {
                                                                                switch ((("" + "") + (("hello" + (("foo bar" + ("\n\t" + (("" + "hello") + "a"))) + (("" + "a") + ("foo bar" + "a")))) + "")) + "") {
                                                                                case "\n\t":
                                                                                    fmt.Println(v2)
                                                                                case "foo bar":
                                                                                    var v82 int16 = v13
                                                                                    _ = v82
                                                                                case "a":
                                                                                    var v83 int16 = v13
                                                                                    _ = v83
                                                                                }
                                                                            }
                                                                            _ = v81
                                                                        }
                                                                        var v84 uint16 = (v4 << (((^(v2))) & 63))
                                                                        v84 -= (v52 + v4)
                                                                        fmt.Println(v84, v3)
                                                                        _ = v84
                                                                    }
                                                                }
                                                            }
                                                            var v85 uint16 = (v4 & (v4 - 32767))
                                                            _ = v85
                                                        }
                                                        _ = v58
                                                    }
                                                }
                                                v13 = v13
                                                _ = v52
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v86 uint16 = (v4 ^ (v4 % (((v4 >> ((1) & 63))) | 1)))
                                                _ = v86
                                            }
                                            var v87 uint16 = 253
                                            if ((fi33 == 32768) || (!(!v3))) {
                                                _ = len(v1)
                                            }
                                            _ = v51
                                            _ = v87
                                        }
                                    } else {
                                        v88 := GenS3{F0: -2147483648, F1: v3}
                                        var v89 uint8 = 1
                                        if (0.1 > 3.14159) {
                                            v90 := GenS3{F0: -2147483648, F1: v3}
                                            _ = v90
                                        } else {
                                            v5 = v5
                                        }
                                        v89 = v47
                                        if ((v47 <= v89) && false) {
                                            if len(v1) > 0 {
                                                v1[0] = (fi33 * fi33)
                                            }
                                        } else {
                                            v91 := GenS5{F0: 0, F1: "\n\t", F2: (v89 / (((v47 - v47)) | 1))}
                                            _ = v91
                                        }
                                        fmt.Println(v14, v47, v4)
                                        fmt.Println(v3, v5)
                                        _ = v88
                                        _ = v89
                                    }
                                    _ = v47
                                }
                            }
                        }
                    }
                    var v92 uint16 = (^(v4))
                    v4 = uint16(v2)
                    _ = fi34
                    _ = v92
                }
                _ = fi33
            }
        default:
            switch 536 {
            case 32768:
                for fi93 := 0; fi93 < 2; fi93++ {
                    v13 = -1
                    var v94 int32 = 110
                    v3 = false
                    _ = fi93
                    _ = v94
                }
            case -9223372036854775808:
                for fi95 := 0; fi95 < 11; fi95++ {
                    var v96 uint32 = (v15 % ((v15) | 1))
                    _ = fi95
                    _ = v96
                }
            case -1:
                switch ("" + "") {
                case "":
                    var v97 float64 = 1e10
                    _ = v97
                }
            default:
                var v98 int16 = v13
                _ = v98
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
