package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint64; F1 int; F2 int64 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int32; F1 int16; F2 int8 }
type GenS10 struct { F0 uint32; F1 int16; F2 uint }
type GenS11 struct { F0 int8; F1 int16; F2 uint32 }
type GenS12 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2147483647, 2147483648, 0}
    var v4 int8 = -76
    var v5 uint32 = 739
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint = 42
        var v8 float32 = -2.25
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = len(v1)
    v3 = append(v3, (^(v5)))
    fmt.Println(v4, v4)
    switch "" {
    case "\n\t":
        var v9 int = 0
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                v4 += v4
                fmt.Println(fi10, fi10)
                var v12 int16 = 42
                _ = fi11
                _ = v12
            }
            switch (fi10 << ((886) & 63)) {
            case 2147483647:
                if (((!false) || false) && false) {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 float32 = 3.14159
                        v15 := GenS1{F0: "hello", F1: (v5 >> ((2) & 63))}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                } else {
                    v3 = append(v3, 65535)
                    v5 -= v5
                    for fi16 := 0; fi16 < 9; fi16++ {
                        if (!((((3.14159 * (1.5 / ((0.1 - (0.1 / -2.25)) + ((3.14159 * ((1e10 - 3.14159) + 1e10)) + 3.14159)))) * 1e10) + 3.14159) < 1e10)) {
                            v5 += v5
                            fmt.Println(v5)
                            v4 = (v4 & (v4 & 42))
                        } else {
                            if ((1.5 / -2.25) <= (1.5 + -2.25)) {
                                v5 += v5
                            } else {
                                v5 = (^(v5))
                                delete(v1, "hello")
                            }
                            v17 := GenS1{F0: "foo bar", F1: v5}
                            v4 -= v4
                            v4 = v4
                            fmt.Println(fi16, v5, v4)
                            _ = v17
                        }
                        var v18 uint32 = (v5 % (((v5 * (v5 << ((42) & 63)))) | 1))
                        v4 += 1
                        _ = fi16
                        _ = v18
                    }
                }
            default:
                v5 += v5
            }
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                if (790 <= 66) {
                    if true {
                        for fi19 := 0; fi19 < 4; fi19++ {
                            var v20 uint32 = (v5 >> ((934) & 63))
                            v5 = v20
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        if false {
                            v4 = -1
                            v21 := GenS2{F0: 100}
                            _ = v21
                        }
                        var v22 uint = 18446744073709551615
                        v5 -= v5
                        _ = v22
                    }
                } else {
                    v4 = v4
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v5 = v5
                        if (18446744073709551615 != 65535) {
                            if ((255 != 233) && (65535 != 830)) {
                                v24 := GenS3{F0: 1000, F1: v4}
                                v5 = 4294967295
                                var v25 uint8 = 42
                                _ = v24
                                _ = v25
                            } else {
                                var v26 float64 = float64(((-2.25 / 0.1) * 3.14159))
                                _ = v26
                            }
                        } else {
                            var v27 float32 = -2.25
                            _ = v27
                        }
                        _ = fi23
                    }
                    v5 += v5
                }
            case "foo bar":
                v4 = (v4 << ((4294967295) & 63))
            case "a":
                v5 -= v5
            default:
                if (uint8(v5) > 2) {
                    var v28 uint8 = 127
                    v29 := GenS4{F0: 65535, F1: fi10}
                    v30 := GenS4{F0: (v5 >> ((42) & 63)), F1: 897}
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    var v31 float32 = (0.1 * 3.14159)
                    v4 = -127
                    _ = v31
                }
            }
            _ = fi10
        }
    case "":
        switch "foo bar" {
        case "foo bar":
            v32 := GenS5{F0: 42}
            _ = v32
        case "":
            switch ("a" + "\n\t") {
            case "\n\t":
                var v33 string = ""
                _ = v33
            }
        default:
            var v34 int32 = 42
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        switch (-(fi35)) {
        case -9223372036854775808:
            fmt.Println(v4, v5, v4)
        case 0:
            if true {
                var v36 uint8 = 157
                v4 += (v4 | (v4 << ((42) & 63)))
                delete(v1, ("\n\t" + "a"))
                _ = v36
            } else {
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 uint64 = 1
                    v38 = v38
                    v38 -= 749
                    _ = fi37
                    _ = v38
                }
            }
        default:
            var v39 uint8 = 0
            _ = v39
        }
        v40 := GenS6{F0: 1, F1: 851}
        if (false && false) {
            switch 9223372036854775807 {
            case 780:
                var v41 int32 = 2147483647
                _ = v41
            case 391:
                if true {
                    var v42 uint = 584
                    _ = v42
                } else {
                    if len(v3) > 0 {
                        v3[0] = v5
                    }
                    for fi43 := 0; fi43 < 13; fi43++ {
                        var v44 bool = false
                        if v44 {
                            var v45 int64 = 1
                            var v46 uint64 = 2
                            var v47 int32 = 42
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        _ = fi43
                        _ = v44
                    }
                }
            case 2:
                v4 -= 23
            default:
                if true {
                    for fi48 := 0; fi48 < 4; fi48++ {
                        switch 9223372036854775807 {
                        case 32768:
                            if (0 <= 2147483647) {
                                v49 := GenS5{F0: -1}
                                v5 = 42
                                if false {
                                    v5 = v5
                                    if (!(false && true)) {
                                        if ("foo bar" < "foo bar") {
                                            var v50 uint32 = 42
                                            var v51 uint8 = 1
                                            if true {
                                                v51 -= v51
                                                fmt.Println(v5)
                                                _ = len(v1)
                                            } else {
                                                fmt.Println(v50)
                                                v50 -= (v50 + v5)
                                            }
                                            var v52 bool = (231 > v51)
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        }
                                        v53 := GenS7{F0: 308, F1: fi48, F2: -9223372036854775808}
                                        fmt.Println(v4, v4)
                                        _ = v53
                                    }
                                    v54 := GenS1{F0: (("" + "") + ("foo bar" + "hello")), F1: 1}
                                    _ = v54
                                }
                                _ = v49
                            }
                        default:
                            delete(v1, "a")
                        }
                        var v55 uint64 = 0
                        fmt.Println(v4, v5)
                        _ = fi48
                        _ = v55
                    }
                    v2[2147483648] = 42
                    for fi56 := 0; fi56 < 14; fi56++ {
                        if (!((51 <= 219) || false)) {
                            switch "foo bar" {
                            case "":
                                v5 -= 1
                            case "foo bar":
                                switch fi35 {
                                case -1:
                                    if (!(!(0 == -72))) {
                                        if (int64(v4) < 829) {
                                            fmt.Println(v5)
                                            if false {
                                                v4 = (v4 % ((v4) | 1))
                                                v57 := GenS2{F0: (v4 / (((v4 | 127)) | 1))}
                                                _ = v57
                                            }
                                        } else {
                                            switch "\n\t" {
                                            case "hello":
                                                v5 += 473
                                            case "\n\t":
                                                if (!("hello" == "hello")) {
                                                    v4 += -128
                                                    v58 := GenS6{F0: 42, F1: 0}
                                                    _ = v58
                                                }
                                            default:
                                                if (!(v5 == 505)) {
                                                    fmt.Println(fi35, v5)
                                                    v59 := GenS4{F0: v5, F1: 127}
                                                    var v60 float32 = 0.0
                                                    if (v60 > (0.1 - v60)) {
                                                        if (false && ((true && false) || true)) {
                                                            delete(v2, (^(v5)))
                                                            v4 = v4
                                                            var v61 float32 = (v60 / v60)
                                                            _ = v61
                                                        }
                                                        var v62 uint8 = 127
                                                        _ = v62
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            delete(v1, "a")
                                                        default:
                                                            fmt.Println(v4, fi56, v5)
                                                        }
                                                        var v63 uint64 = 864
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        fmt.Println(v60)
                                                        v3 = append(v3, 129)
                                                        if (false || ((!((^(v5)) >= (v5 << ((801) & 63)))) || true)) {
                                                            var v64 uint64 = 18446744073709551615
                                                            var v65 int64 = 1
                                                            var v66 uint64 = (v64 - 1)
                                                            if false {
                                                                _ = len(v2)
                                                                if false {
                                                                    fmt.Println(fi56, v66)
                                                                    v67 := GenS4{F0: 42, F1: fi56}
                                                                    fmt.Println(v64, fi35, fi35)
                                                                    _ = v67
                                                                } else {
                                                                    v68 := GenS8{F0: 2, F1: int64(v4)}
                                                                    v64 = v63
                                                                    fmt.Println(v65)
                                                                    _ = v68
                                                                }
                                                                v65 += int64(v4)
                                                            } else {
                                                                delete(v1, "a")
                                                            }
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        v69 := GenS9{F0: 829, F1: 42, F2: (^(v4))}
                                                        _ = v63
                                                        _ = v69
                                                    }
                                                    var v70 bool = (false || false)
                                                    _ = v59
                                                    _ = v60
                                                    _ = v70
                                                }
                                            }
                                            v4 = 100
                                            if true {
                                                switch (fi56 + (fi35 ^ (fi35 << ((42) & 63)))) {
                                                case -596:
                                                    v5 = v5
                                                case 128:
                                                    var v71 bool = true
                                                    _ = v71
                                                default:
                                                    var v72 int64 = 9223372036854775807
                                                    _ = v72
                                                }
                                            } else {
                                                v5 = 4294967295
                                            }
                                        }
                                        var v73 float32 = (((1.5 * 1.5) - float32(1.5)) / (0.0 + (0.1 / (((float32(1e10) / 1.0) - (((0.1 + 0.0) + 0.0) / 1.0)) - 1.5))))
                                        _ = v73
                                    } else {
                                        v74 := GenS10{F0: 42, F1: 0, F2: 65535}
                                        var v75 uint16 = 1
                                        _ = v74
                                        _ = v75
                                    }
                                case 0:
                                    v4 = (v4 - (v4 / ((v4) | 1)))
                                    fmt.Println(v4)
                                    if ((1000 >= -83) && ((971 >= 176) && (2 >= uint16(v5)))) {
                                        var v76 int32 = 0
                                        _ = v76
                                    } else {
                                        var v77 float32 = 1e10
                                        var v78 uint16 = 1
                                        _ = v77
                                        _ = v78
                                    }
                                case 538:
                                    v79 := GenS7{F0: 191, F1: -493, F2: -173}
                                    _ = v79
                                default:
                                    var v80 int16 = 1000
                                    _ = v80
                                }
                            default:
                                delete(v2, (v5 * 744))
                            }
                        }
                        v4 += v4
                        var v81 uint32 = v5
                        _ = fi56
                        _ = v81
                    }
                } else {
                    var v82 string = (("" + ("" + "a")) + "")
                    v83 := GenS9{F0: 0, F1: 747, F2: v4}
                    _ = v82
                    _ = v83
                }
            }
            switch fi35 {
            case 946:
                if (419 != 65535) {
                    v84 := GenS11{F0: v4, F1: 88, F2: v5}
                    _ = v84
                }
            case -689:
                var v85 int64 = -22
                _ = v85
            case 4294967295:
                switch (-(fi35)) {
                case -9223372036854775808:
                    v4 = (v4 - v4)
                case -570:
                    switch -9223372036854775808 {
                    case 127:
                        switch (fi35 - int(v5)) {
                        case 1:
                            fmt.Println(fi35, fi35)
                        }
                    }
                default:
                    switch "\n\t" {
                    case "foo bar":
                        v5 += 731
                    case "a":
                        fmt.Println(v4, fi35, v5)
                    }
                }
            default:
                for fi86 := 0; fi86 < 5; fi86++ {
                    var v87 uint16 = 0
                    _ = fi86
                    _ = v87
                }
            }
            v4 = v4
        } else {
            var v88 int32 = 2147483647
            v89 := GenS6{F0: -398, F1: 42}
            var v90 uint16 = 603
            _ = v88
            _ = v89
            _ = v90
        }
        _ = fi35
        _ = v40
    }
    v91 := GenS11{F0: (v4 >> ((82) & 63)), F1: 1, F2: 65535}
    var v92 float32 = float32(((0.0 - (0.0 + ((float64(0.0) + 0.1) / ((1e10 * 1e10) + 1.5)))) + 0.0))
    v92 += ((v92 + v92) * v92)
    v5 += 4294967295
    var v93 bool = true
    switch 127 {
    case 32768:
        var v94 int64 = 84
        _ = v94
    case 4294967295:
        for fi95 := 0; fi95 < 13; fi95++ {
            v96 := GenS6{F0: 859, F1: 864}
            if ((v93 || (4294967295 > 696)) || true) {
                v3 = append(v3, v5)
                var v97 uint16 = 2
                if false {
                    if (!(697 <= 0)) {
                        switch (fi95 - (fi95 ^ fi95)) {
                        case 32767:
                            switch ("" + ("hello" + ((("" + ("foo bar" + "\n\t")) + ("\n\t" + ("" + ("foo bar" + "hello")))) + ""))) {
                            case "\n\t":
                                v2[v5] = v5
                            default:
                                var v98 string = "hello"
                                _ = v98
                            }
                        case -551:
                            var v99 int8 = v4
                            _ = v99
                        }
                        v4 = v4
                    } else {
                        for fi100 := 0; fi100 < 17; fi100++ {
                            var v101 uint32 = v5
                            _ = fi100
                            _ = v101
                        }
                        v97 = v97
                        v93 = true
                    }
                } else {
                    for fi102 := 0; fi102 < 20; fi102++ {
                        v3 = append(v3, 352)
                        switch "a" {
                        case "a":
                            v5 += (v5 | (^(v5)))
                        default:
                            v103 := GenS12{F0: 9223372036854775807}
                            _ = v103
                        }
                        _ = fi102
                    }
                }
                _ = v97
            }
            _ = fi95
            _ = v96
        }
    default:
        for fi104 := 0; fi104 < 17; fi104++ {
            v92 = 1e10
            _ = len(v2)
            _ = fi104
        }
    }
    v5 += 4294967295
    fmt.Println(v92)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
    _ = v92
    _ = v93
}
