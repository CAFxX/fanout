package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 int32; F1 int16; F2 uint16 }
type GenS5 struct { F0 int16; F1 int8; F2 bool }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int64; F1 int32; F2 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 uint16; F2 string }
type GenS11 struct { F0 bool; F1 int16 }
type GenS12 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8, p2 bool) int16 {
    v7 := []int{-2147483648, 127}
    var v8 int = 135
    if p2 {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: p1, F1: true, F2: ("" + ("hello" + (("a" + "\n\t") + "\n\t")))}
            _ = fi9
            _ = v10
        }
        v7 = append(v7, v8)
    } else {
        v8 += 2147483647
        if (("foo bar" + "\n\t") < "foo bar") {
            v11 := GenS1{F0: p1, F1: true, F2: "foo bar"}
            switch "foo bar" {
            case "":
                var v12 uint = 1
                _ = v12
            }
            p2 = p2
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 bool) bool {
    v13 := []int{1, 9223372036854775807}
    var v14 int = 1
    if p0 {
        if ((-2.25 != 1.5) || (255 > 0)) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if (-627 < -9223372036854775808) {
                switch "foo bar" {
                case "a":
                    v14 -= v14
                default:
                    p0 = (42 >= 78)
                }
                if len(v13) > 0 {
                    v13[0] = v14
                }
                switch 4294967295 {
                case 0:
                    switch "a" {
                    case "a":
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 int32 = -2147483648
                            v17 := GenS2{F0: "\n\t", F1: 1}
                            var v18 string = ("foo bar" + "")
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    }
                case 2147483647:
                    v19 := GenS2{F0: "", F1: 2147483648}
                    _ = v19
                }
            } else {
                p0 = (0 != 332)
                var v20 int32 = 1
                v14 = 128
                _ = v20
            }
        }
    } else {
        switch "" {
        case "hello":
            switch (-(v14)) {
            case -2147483648:
                var v21 int = (^(v14))
                _ = v21
            default:
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
        case "":
            fmt.Println(v14, v14)
        default:
            fmt.Println(v14)
        }
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    _ = p0
    _ = v13
    _ = v14
    return p0
}

func main() {
    var v1 uint16 = 255
    var v2 int8 = -1
    var v3 int16 = 1000
    var v4 bool = (!(-1 != -1))
    v5 := []uint32{2, 2}
    var v6 string = (("\n\t" + ("foo bar" + ("hello" + "\n\t"))) + (("\n\t" + "foo bar") + ("\n\t" + ("\n\t" + ("a" + "hello")))))
    var v22 uint = 1
    if false {
        var v23 float32 = 1e10
        if v4 {
            for fi24 := 0; fi24 < 12; fi24++ {
                if (0.0 != (((v23 / (v23 * (float32((-2.25 / 3.14159)) + 1.5))) / ((v23 + (v23 * ((v23 / float32(-2.25)) + 0.0))) * (v23 - v23))) / v23)) {
                    v1 += v1
                    switch (fi24 % ((fi24) | 1)) {
                    case 0:
                        var v25 float64 = 1e10
                        _ = v25
                    case 2147483647:
                        if v4 {
                            var v26 uint32 = 65535
                            v27 := GenS1{F0: 128, F1: v4, F2: ((v6 + "\n\t") + "")}
                            _ = v26
                            _ = v27
                        }
                    case -9223372036854775808:
                        fmt.Println(fi24)
                    }
                    fmt.Println(v6, v6)
                }
                switch (fi24 << ((63) & 63)) {
                case 127:
                    switch (v6 + (("a" + "hello") + "hello")) {
                    case "":
                        switch v6 {
                        case "a":
                            v23 -= v23
                        }
                    case "hello":
                        v1 += 65535
                    default:
                        fmt.Println(v2, v4, v23)
                    }
                case -9223372036854775808:
                    if (-1 == (v3 + genFn0(v2, 42, v4))) {
                        v28 := GenS3{F0: 0, F1: v6, F2: -1}
                        v29 := GenS3{F0: (v2 & -97), F1: (v6 + v6), F2: -2147483648}
                        _ = v28
                        _ = v29
                    } else {
                        _ = len(v5)
                    }
                default:
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v31 := GenS2{F0: "hello", F1: 4294967295}
                        _ = fi30
                        _ = v31
                    }
                }
                var v32 uint8 = 112
                _ = fi24
                _ = v32
            }
            v2 = (v2 >> (((v22 ^ 65535)) & 63))
        }
        var v33 float64 = 1e10
        _ = v23
        _ = v33
    }
    v34 := GenS3{F0: (v2 << ((v22) & 63)), F1: (v6 + v6), F2: 875}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        if v4 {
            v36 := GenS4{F0: int32(v22), F1: (v3 / ((v3) | 1)), F2: (v1 & v1)}
            _ = v36
        } else {
            for fi37 := 0; fi37 < 18; fi37++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v38 bool = (2 == 817)
                v38 = v38
                _ = fi37
                _ = v38
            }
            var v39 int16 = v3
            v40 := GenS2{F0: v6, F1: 4294967295}
            _ = v39
            _ = v40
        }
        _ = len(v5)
        v22 = v22
        _ = fi35
    }
    v41 := GenS2{F0: ((v6 + v6) + v6), F1: 65535}
    _ = len(v5)
    v1 = (^(v1))
    var v42 bool = true
    if len(v5) > 0 {
        v5[0] = 51
    }
    if len(v5) > 0 {
        v5[0] = 530
    }
    var v43 int32 = 2147483647
    if ((!(((v42 && (v2 < (v2 & 1))) && (false || (v6 > v6))) && (1 == 268))) && ((128 == 59) || (829 >= -380))) {
        fmt.Println(v1)
        _ = len(v5)
        var v44 int = int(v43)
        _ = v44
    } else {
        for fi45 := 0; fi45 < 16; fi45++ {
            var v46 uint8 = 42
            for fi47 := 0; fi47 < 19; fi47++ {
                switch int(v2) {
                case 32767:
                    if (0.0 < -2.25) {
                        v48 := GenS1{F0: 177, F1: (0.1 == ((1e10 * 3.14159) / ((1.5 * 0.0) - 1e10))), F2: v6}
                        v1 += (v1 * (v1 ^ (v1 & (v1 - (^(v1))))))
                        _ = v48
                    } else {
                        v1 -= v1
                        v2 = (v2 ^ (-(v2)))
                        var v49 int64 = 42
                        _ = v49
                    }
                case 127:
                    if v4 {
                        if (0 <= 2) {
                            var v50 int64 = -9223372036854775808
                            _ = v50
                        } else {
                            var v51 uint32 = uint32(v1)
                            switch (((v6 + v6) + ("hello" + "foo bar")) + ("\n\t" + (v6 + v6))) {
                            case "":
                                v3 += v3
                            case "a":
                                var v52 uint64 = 2
                                _ = v52
                            default:
                                v51 = (^(v51))
                            }
                            switch v6 {
                            case "hello":
                                v1 = uint16(v46)
                            case "\n\t":
                                v53 := GenS1{F0: uint8(v51), F1: (v43 < (v43 << ((v22) & 63))), F2: v6}
                                _ = v53
                            case "a":
                                fmt.Println(v46, v51, v46)
                                if v42 {
                                    if ("a" == ("hello" + (("foo bar" + v6) + v6))) {
                                        if v4 {
                                            fmt.Println(v2, v1, v43)
                                        }
                                        v6 = (v6 + "\n\t")
                                        var v54 int = fi45
                                        _ = v54
                                    } else {
                                        v2 += v2
                                    }
                                    if ((1e10 / (float64(3.14159) - 0.1)) <= (-2.25 / ((-2.25 + ((1.5 + 1e10) + 0.0)) / 1e10))) {
                                        var v55 int8 = v2
                                        switch (^(fi47)) {
                                        case 2147483647:
                                            v55 = (v2 & (v55 << ((v22) & 63)))
                                            if true {
                                                v43 -= (-(v43))
                                                v56 := GenS5{F0: (v3 / ((0) | 1)), F1: v2, F2: v4}
                                                v57 := GenS6{F0: v22, F1: 255}
                                                if (false || (!(!((v2 & (v55 % (((v2 % ((v55) | 1))) | 1))) > (v2 << ((v22) & 63)))))) {
                                                    var v58 int32 = v43
                                                    _ = v58
                                                }
                                                _ = v56
                                                _ = v57
                                            } else {
                                                fmt.Println(v22, v42)
                                                _ = len(v5)
                                            }
                                        }
                                        if v42 {
                                            var v59 uint16 = 32767
                                            _ = v59
                                        } else {
                                            var v60 float32 = 1.5
                                            if (!((v1 >> (((^(v22))) & 63)) != (v1 | (v1 * v1)))) {
                                                if true {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    v61 := GenS5{F0: 42, F1: (v55 ^ -9), F2: v4}
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v61
                                                }
                                            }
                                            if v42 {
                                                switch "a" {
                                                case "hello":
                                                    v5 = append(v5, v51)
                                                default:
                                                    var v62 bool = ((1 != 100) || true)
                                                    _ = v62
                                                }
                                                if v42 {
                                                    var v63 uint64 = uint64(v46)
                                                    v55 += v2
                                                    switch (v6 + ("" + (v6 + v6))) {
                                                    case "hello":
                                                        v64 := GenS6{F0: v22, F1: v22}
                                                        _ = v64
                                                    }
                                                    _ = v63
                                                }
                                                var v65 int16 = v3
                                                if v42 {
                                                    if v42 {
                                                        var v66 int = fi47
                                                        _ = v66
                                                    } else {
                                                        var v67 uint16 = 598
                                                        _ = v67
                                                    }
                                                    v1 -= 523
                                                } else {
                                                    v1 = v1
                                                    switch ("\n\t" + v6) {
                                                    case "a":
                                                        switch (-(fi47)) {
                                                        case -1:
                                                            v68 := GenS5{F0: v65, F1: -1, F2: v4}
                                                            _ = v68
                                                        case -572:
                                                            switch 361 {
                                                            case 472:
                                                                var v69 uint32 = v51
                                                                _ = v69
                                                            default:
                                                                var v70 float64 = float64(v60)
                                                                _ = v70
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = 352
                                                            }
                                                            v60 -= ((float32(3.14159) - -2.25) - v60)
                                                        case 4294967295:
                                                            _ = len(v5)
                                                        default:
                                                            var v71 string = v6
                                                            _ = v71
                                                        }
                                                    case "hello":
                                                        switch (fi45 % ((32767) | 1)) {
                                                        case 9223372036854775807:
                                                            var v72 int = fi47
                                                            _ = v72
                                                        case 0:
                                                            _ = len(v5)
                                                            var v73 string = v6
                                                            v73 = v6
                                                            _ = v73
                                                        default:
                                                            if len(v5) > 0 {
                                                                v5[0] = (v51 % (((v51 ^ 514)) | 1))
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                    v2 = -1
                                                }
                                                switch ("\n\t" + v6) {
                                                case "hello":
                                                    if false {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if (0.1 == v60) {
                                                            switch (fi47 / ((32767) | 1)) {
                                                            case 32768:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            case 4294967295:
                                                                if v42 {
                                                                    var v74 float64 = (3.14159 * float64(((float32(0.1) * 1e10) - 0.1)))
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    _ = v74
                                                                }
                                                            }
                                                            var v75 uint32 = (v51 / ((v51) | 1))
                                                            if len(v5) > 0 {
                                                                v5[0] = v51
                                                            }
                                                            _ = v75
                                                        } else {
                                                            v65 += (v3 / (((v3 & 0)) | 1))
                                                            var v76 uint16 = 255
                                                            fmt.Println(v22)
                                                            if (!(0.0 > 0.0)) {
                                                                v77 := GenS7{F0: -823, F1: v43, F2: (v2 >> ((v22) & 63))}
                                                                _ = v77
                                                            } else {
                                                                switch v6 {
                                                                case "foo bar":
                                                                    v46 = 128
                                                                case "a":
                                                                    v5 = append(v5, v51)
                                                                case "":
                                                                    if (-905 >= 1) {
                                                                        v78 := GenS3{F0: (v2 << ((v22) & 63)), F1: "", F2: v43}
                                                                        v79 := GenS4{F0: v43, F1: (v65 * 42), F2: (v1 % (((v76 | v1)) | 1))}
                                                                        var v80 int64 = 9223372036854775807
                                                                        v46 = (v46 * v46)
                                                                        var v81 float64 = 0.1
                                                                        _ = v78
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                    }
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                            _ = v76
                                                        }
                                                        switch v6 {
                                                        case "":
                                                            var v82 int32 = v43
                                                            _ = v82
                                                        }
                                                    }
                                                }
                                                _ = v65
                                            } else {
                                                var v83 int64 = 396
                                                _ = v83
                                            }
                                            _ = v60
                                        }
                                        _ = v55
                                    } else {
                                        v22 += v22
                                        var v84 uint = 394
                                        v84 = (v22 << (((^(v22))) & 63))
                                        _ = v84
                                    }
                                    v85 := GenS7{F0: 235, F1: (v43 << (((v22 / (((v22 / (((v22 + v22)) | 1))) | 1))) & 63)), F2: (v2 * v2)}
                                    _ = v85
                                } else {
                                    var v86 int64 = 490
                                    v87 := GenS6{F0: 927, F1: (v22 << ((v22) & 63))}
                                    _ = v86
                                    _ = v87
                                }
                            }
                            if ((!v42) || false) {
                                fmt.Println(v42, v1)
                                if (int64(v3) != 42) {
                                    v88 := GenS8{F0: 42}
                                    var v89 bool = v4
                                    if ((v51 | v51) != (v51 / ((v51) | 1))) {
                                        if (int32(v1) >= int32(v2)) {
                                            fmt.Println(v2)
                                            if true {
                                                if (753 != uint64(v1)) {
                                                    v5 = append(v5, uint32(v22))
                                                    switch "\n\t" {
                                                    case "":
                                                        fmt.Println(v42, v1)
                                                    case "foo bar":
                                                        switch v6 {
                                                        case "\n\t":
                                                            v90 := GenS8{F0: v46}
                                                            _ = v90
                                                        }
                                                    case "hello":
                                                        if (v46 <= v46) {
                                                            fmt.Println(v42)
                                                            if v4 {
                                                                v91 := GenS9{F0: (-(v2))}
                                                                _ = v91
                                                            } else {
                                                                var v92 uint16 = (v1 % ((uint16(v46)) | 1))
                                                                if len(v5) > 0 {
                                                                    v5[0] = 2
                                                                }
                                                                _ = v92
                                                            }
                                                            v51 = v51
                                                        } else {
                                                            var v93 bool = v4
                                                            fmt.Println(v42)
                                                            _ = v93
                                                        }
                                                    }
                                                } else {
                                                    var v94 float32 = 0.1
                                                    var v95 int32 = v43
                                                    _ = v94
                                                    _ = v95
                                                }
                                                v1 -= (^(v1))
                                            } else {
                                                switch 4294967295 {
                                                case -967:
                                                    if len(v5) > 0 {
                                                        v5[0] = v51
                                                    }
                                                case 511:
                                                    if ((^(v22)) > v22) {
                                                        var v96 uint32 = v51
                                                        _ = v96
                                                    }
                                                default:
                                                    fmt.Println(v42)
                                                }
                                                v46 -= v46
                                                var v97 int64 = 0
                                                _ = v97
                                            }
                                        } else {
                                            var v98 int32 = v43
                                            fmt.Println(v6, v89, v2)
                                            fmt.Println(v6, v98)
                                            var v99 int64 = -1
                                            _ = v98
                                            _ = v99
                                        }
                                        v100 := GenS7{F0: 934, F1: v43, F2: (-(v2))}
                                        var v101 uint64 = 42
                                        _ = v100
                                        _ = v101
                                    } else {
                                        switch 128 {
                                        case -1:
                                            fmt.Println(v51, v1, v46)
                                        case 2147483647:
                                            if len(v5) > 0 {
                                                v5[0] = v51
                                            }
                                        case 4294967295:
                                            v3 = (v3 & (v3 + (v3 + (v3 & v3))))
                                        }
                                        v22 = (v22 * 255)
                                        switch v6 {
                                        case "\n\t":
                                            v51 += (^(v51))
                                        default:
                                            if (v42 && false) {
                                                if true {
                                                    v4 = v42
                                                    _ = len(v5)
                                                    switch 127 {
                                                    case -230:
                                                        var v102 int = -31
                                                        _ = v102
                                                    case 42:
                                                        var v103 int32 = (v43 ^ 145)
                                                        _ = v103
                                                    case 515:
                                                        var v104 uint = (v22 | v22)
                                                        _ = v104
                                                    default:
                                                        v51 -= 2
                                                    }
                                                } else {
                                                    v3 = (v3 ^ v3)
                                                    v105 := GenS10{F0: (v2 % ((v2) | 1)), F1: 1, F2: v6}
                                                    _ = v105
                                                }
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = v51
                                                }
                                                v106 := GenS6{F0: (v22 % ((uint(v1)) | 1)), F1: v22}
                                                if len(v5) > 0 {
                                                    v5[0] = v51
                                                }
                                                _ = v106
                                            }
                                            v107 := GenS11{F0: (!true), F1: 1000}
                                            _ = v107
                                        }
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = v51
                                    }
                                    _ = v88
                                    _ = v89
                                }
                            }
                            _ = v51
                        }
                        v43 = v43
                        v5 = append(v5, 65535)
                    }
                    v108 := GenS5{F0: int16(v43), F1: v2, F2: ((v3 * v3) < v3)}
                    _ = v108
                case 9223372036854775807:
                    var v109 float64 = 3.14159
                    _ = v109
                default:
                    v22 = 0
                }
                _ = fi47
            }
            fmt.Println(fi45, v42, v42)
            _ = fi45
            _ = v46
        }
        v110 := GenS12{F0: (v3 * int16(v43))}
        _ = v110
    }
    var v111 uint = (v22 ^ (v22 ^ 18446744073709551615))
    v4 = ((!false) || true)
    if (63 <= v2) {
        var v112 int16 = -32768
        _ = v112
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v111
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
