package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 string; F2 int16 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 string; F1 uint32; F2 int32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint16; F1 bool }
type GenS16 struct { F0 int64; F1 uint16 }
type GenS17 struct { F0 uint16; F1 uint16 }
type GenS18 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS19 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 int }
type GenS22 struct { F0 uint; F1 uint; F2 string }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 uint8; F1 bool; F2 int }
type GenS25 struct { F0 uint; F1 uint32 }
type GenS26 struct { F0 uint }
type GenS27 struct { F0 int8; F1 int8 }
type GenS28 struct { F0 uint64; F1 uint }
type GenS29 struct { F0 uint64 }
type GenS30 struct { F0 uint8; F1 string; F2 int8 }
type GenS31 struct { F0 uint8 }
type GenS32 struct { F0 uint; F1 int32; F2 uint8 }
type GenS33 struct { F0 int; F1 uint64; F2 int8 }
type GenS34 struct { F0 int16; F1 uint8 }
func genFn0(p0 bool, p1 int32, p2 uint32) int16 {
    if p0 {
        p0 = p0
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 float32 = 1.5
            var v8 uint8 = 1
            v9 := GenS1{F0: ("hello" + "\n\t")}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    } else {
        fmt.Println(p0, p1, p2)
        var v10 int = 457
        _ = v10
    }
    var v11 uint16 = 0
    p2 -= p2
    if ((("a" + "") + ("" + "foo bar")) > (("" + "\n\t") + ("a" + "hello"))) {
        var v12 int = 294
        switch "a" {
        case "\n\t":
            var v13 uint32 = (p2 / (((^(p2))) | 1))
            _ = v13
        case "":
            if (((!(((("\n\t" + "foo bar") >= "foo bar") || p0) && (297 == (p2 << ((2) & 63))))) || p0) || p0) {
                var v14 uint16 = (v11 + (v11 - v11))
                _ = v14
            }
        }
        _ = v12
    }
    v15 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v15
    return 0
}

func main() {
    var v1 bool = (153 > 218)
    v2 := []int{-9223372036854775808, 959}
    var v3 uint8 = 2
    var v4 bool = (404 != 0)
    v5 := []int{int(v3), 32768, 32767}
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: 1}
        v18 := GenS3{F0: "\n\t"}
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 uint = 816
            _ = fi19
            _ = v20
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    v4 = v4
    v4 = (!false)
    for fi21 := 0; fi21 < 2; fi21++ {
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
        var v22 uint = 65535
        for fi23 := 0; fi23 < 1; fi23++ {
            v5 = append(v5, (fi21 | fi23))
            var v24 float64 = (1.5 + 0.1)
            if (v1 || v4) {
                if (!v4) {
                    var v25 string = ""
                    fmt.Println(fi21)
                    _ = v25
                } else {
                    var v26 uint8 = v3
                    v27 := GenS4{F0: -1}
                    fmt.Println(v26, v26)
                    var v28 float64 = 0.1
                    v29 := GenS5{F0: v22}
                    v24 = 1.5
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                switch fi21 {
                case 42:
                    v22 -= (v22 & 647)
                    if ((false && v4) && (("\n\t" + ("hello" + "a")) > ("a" + "\n\t"))) {
                        var v30 float64 = (v24 + (0.0 / v24))
                        v3 = (v3 % ((v3) | 1))
                        var v31 uint64 = 0
                        if (!(!v1)) {
                            fmt.Println(fi21, v30, v1)
                        }
                        _ = v30
                        _ = v31
                    } else {
                        switch int(v3) {
                        case 9223372036854775807:
                            switch "\n\t" {
                            case "foo bar":
                                v3 = v3
                            default:
                                var v32 int8 = -4
                                _ = v32
                            }
                        case -783:
                            v3 += v3
                        case 0:
                            if (994 < 1) {
                                switch (fi23 & (fi23 / ((fi21) | 1))) {
                                case 32767:
                                    v33 := GenS5{F0: (^(v22))}
                                    _ = v33
                                case 504:
                                    fmt.Println(v24, fi21)
                                default:
                                    fmt.Println(v24)
                                }
                                v1 = ("hello" != "")
                            } else {
                                if (127 != v3) {
                                    v3 += v3
                                    _ = len(v2)
                                    if v4 {
                                        v4 = v4
                                        if len(v5) > 0 {
                                            v5[0] = fi21
                                        }
                                        v22 -= v22
                                    }
                                    v34 := GenS5{F0: (v22 << (((v22 | v22)) & 63))}
                                    switch "" {
                                    case "":
                                        fmt.Println(v24, v3, v1)
                                    case "hello":
                                        v35 := GenS6{F0: v3, F1: ("" + (("hello" + "\n\t") + "")), F2: 42}
                                        _ = v35
                                    }
                                    _ = len(v5)
                                    _ = v34
                                } else {
                                    var v36 int = (-(fi21))
                                    if true {
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (true && (!v1)) {
                                                if true {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                } else {
                                                    v4 = (2147483647 >= -162)
                                                    switch (fi21 >> ((0) & 63)) {
                                                    case 32768:
                                                        var v37 string = "a"
                                                        _ = v37
                                                    case 42:
                                                        if true {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            if (v1 && v4) {
                                                                v1 = true
                                                            }
                                                        } else {
                                                            v5 = append(v5, (fi21 << (((v22 | 0)) & 63)))
                                                            v38 := GenS3{F0: "\n\t"}
                                                            switch -733 {
                                                            case 4294967295:
                                                                if (0 > 42) {
                                                                    v22 = v22
                                                                }
                                                            }
                                                            _ = v38
                                                        }
                                                        v1 = ((v4 || ((fi23 & (fi21 - 63)) >= v36)) || true)
                                                    default:
                                                        if (364 < -713) {
                                                            if (0.0 < -2.25) {
                                                                if v4 {
                                                                    var v39 int8 = 0
                                                                    if (872 <= 1) {
                                                                        v4 = ((v3 >> ((602) & 63)) == 42)
                                                                        fmt.Println(v3, v4)
                                                                        v22 += v22
                                                                    }
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if (true && (935 == 2)) {
                                                                            v36 -= fi23
                                                                            if (!(((0.0 - -2.25) - 1.5) >= 0.1)) {
                                                                                var v40 int8 = (v39 << (((v22 ^ v22)) & 63))
                                                                                _ = v40
                                                                            } else {
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                        var v41 uint32 = 976
                                                                        _ = v41
                                                                    case "foo bar":
                                                                        v3 -= (v3 | 2)
                                                                    default:
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            var v42 int = 42
                                                                            _ = v42
                                                                        case "":
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                fmt.Println(v22)
                                                                            case "hello":
                                                                                switch v36 {
                                                                                case 4294967295:
                                                                                    v2 = append(v2, (fi23 - (fi21 - (fi23 >> ((v22) & 63)))))
                                                                                case 2147483647:
                                                                                    v22 = v22
                                                                                default:
                                                                                    v3 = 127
                                                                                }
                                                                            default:
                                                                                v3 += 127
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v39
                                                                } else {
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                            var v43 uint32 = 0
                                                            v44 := GenS5{F0: v22}
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                    }
                                                    if v4 {
                                                        var v45 uint32 = 65535
                                                        _ = v45
                                                    } else {
                                                        var v46 int64 = 0
                                                        _ = v46
                                                    }
                                                }
                                                var v47 int16 = -485
                                                if true {
                                                    var v48 int8 = 100
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi23 | v36)
                                                    }
                                                    v5 = append(v5, (fi23 / ((v36) | 1)))
                                                    _ = v48
                                                }
                                                _ = v47
                                            } else {
                                                _ = len(v2)
                                            }
                                        case "hello":
                                            switch "foo bar" {
                                            case "":
                                                v49 := GenS4{F0: 42}
                                                _ = v49
                                            default:
                                                if v4 {
                                                    if false {
                                                        v22 -= v22
                                                        if v1 {
                                                            var v50 uint64 = 4294967295
                                                            _ = v50
                                                        } else {
                                                            v51 := GenS4{F0: -821}
                                                            _ = v51
                                                        }
                                                    } else {
                                                        var v52 float32 = 1e10
                                                        v53 := GenS7{F0: ("hello" + "hello"), F1: ""}
                                                        v54 := GenS3{F0: ""}
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    v55 := GenS8{F0: 86}
                                                    var v56 int64 = -1
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    var v57 int64 = 341
                                                    v5 = append(v5, v36)
                                                    _ = v57
                                                }
                                            }
                                        case "":
                                            var v58 int = (^(v36))
                                            _ = v58
                                        default:
                                            v4 = v1
                                        }
                                    }
                                    _ = len(v2)
                                    _ = v36
                                }
                                v4 = true
                                var v59 uint64 = 42
                                _ = v59
                            }
                        }
                    }
                case 2147483647:
                    if (("\n\t" + "foo bar") < (("a" + ("hello" + "")) + "foo bar")) {
                        var v60 uint = 704
                        fmt.Println(v24, fi23)
                        v24 -= 1e10
                        _ = v60
                    } else {
                        var v61 bool = (385 > 0)
                        _ = v61
                    }
                    fmt.Println(v4)
                    switch 32767 {
                    case -899:
                        fmt.Println(v1, v3, v24)
                    case 1:
                        v62 := GenS9{F0: 4294967295, F1: -32768}
                        _ = v62
                    case -2147483648:
                        _ = len(v2)
                    }
                    v3 += (v3 << ((v22) & 63))
                    var v63 int32 = 452
                    switch "\n\t" {
                    case "\n\t":
                        var v64 string = "a"
                        _ = v64
                    case "a":
                        v3 = v3
                    }
                    _ = v63
                default:
                    switch fi23 {
                    case -2147483648:
                        fmt.Println(fi21, v1, v3)
                    case 32767:
                        fmt.Println(fi23)
                    case 42:
                        _ = len(v2)
                    default:
                        var v65 int64 = -1
                        _ = v65
                    }
                }
                var v66 int32 = 42
                var v67 uint64 = 4294967295
                fmt.Println(v67, v1, fi23)
                _ = v66
                _ = v67
            } else {
                var v68 int64 = 1
                var v69 int = fi21
                switch (fi23 * 451) {
                case 2147483647:
                    if ((0.1 != -2.25) && true) {
                        var v70 int32 = 490
                        switch fi21 {
                        case 32767:
                            switch fi23 {
                            case 497:
                                v3 = v3
                            case -7:
                                switch "hello" {
                                case "a":
                                    v22 -= v22
                                case "foo bar":
                                    if len(v5) > 0 {
                                        v5[0] = fi23
                                    }
                                    if v1 {
                                        v68 = (v68 * -764)
                                        var v71 uint8 = (v3 + 1)
                                        _ = v71
                                    } else {
                                        if ((v24 + (1.5 - ((v24 * 0.1) - 0.1))) == (float64(3.14159) * ((3.14159 - 3.14159) - (0.0 - float64(1e10))))) {
                                            if (!(303 != 1000)) {
                                                if false {
                                                    var v72 int64 = v68
                                                    if v4 {
                                                        fmt.Println(v1)
                                                        v4 = ((!v1) || false)
                                                        var v73 uint32 = 2147483647
                                                        v5 = append(v5, v69)
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v74 int64 = v72
                                                        var v75 string = (("hello" + ("\n\t" + "a")) + "\n\t")
                                                        switch v75 {
                                                        case "\n\t":
                                                            _ = len(v2)
                                                        default:
                                                            var v76 uint32 = v73
                                                            _ = v76
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v73 = v73
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                    } else {
                                                        v77 := GenS10{F0: (fi23 | fi23)}
                                                        fmt.Println(fi21)
                                                        var v78 int16 = genFn0(false, 0, 4294967295)
                                                        v4 = (1 < (^(v3)))
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    _ = v72
                                                }
                                                fmt.Println(v22, v69)
                                                fmt.Println(fi21, v24, v3)
                                                if (("" > (("" + "\n\t") + (("hello" + ("\n\t" + "hello")) + ("" + "foo bar")))) || (!true)) {
                                                    fmt.Println(v22, v4)
                                                }
                                            }
                                            v24 = v24
                                            var v79 float64 = 1.5
                                            var v80 uint64 = 4294967295
                                            switch fi23 {
                                            case 9223372036854775807:
                                                v81 := GenS6{F0: (v3 - uint8(v80)), F1: ("" + ("" + "\n\t")), F2: 1000}
                                                _ = v81
                                            case -1:
                                                var v82 uint16 = 42
                                                _ = v82
                                            default:
                                                if false {
                                                    if false {
                                                        var v83 float32 = (float32(v79) / (1e10 / -2.25))
                                                        v5 = append(v5, fi21)
                                                        v2 = append(v2, (v69 >> ((v22) & 63)))
                                                        var v84 bool = false
                                                        v22 += 18446744073709551615
                                                        v70 = 2147483647
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        v3 += v3
                                                    }
                                                } else {
                                                    v5 = append(v5, 405)
                                                }
                                            }
                                            _ = v79
                                            _ = v80
                                        } else {
                                            v85 := GenS11{F0: v3}
                                            _ = v85
                                        }
                                    }
                                    if v1 {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                            default:
                                switch v69 {
                                case 127:
                                    fmt.Println(v70)
                                case 128:
                                    if v4 {
                                        if (44 != 26) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        } else {
                                            if v4 {
                                                v69 += (fi23 * fi23)
                                                switch "a" {
                                                case "foo bar":
                                                    if v4 {
                                                        switch 127 {
                                                        case 319:
                                                            v86 := GenS12{F0: "", F1: 1, F2: 2147483647}
                                                            _ = v86
                                                        case -710:
                                                            v5 = append(v5, 1)
                                                        }
                                                        var v87 int8 = 0
                                                        if (v24 != v24) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        if (v1 || v4) {
                                                            fmt.Println(fi21, v1)
                                                            var v88 float32 = 3.14159
                                                            _ = v88
                                                        } else {
                                                            var v89 uint32 = uint32(v22)
                                                            _ = v89
                                                        }
                                                        _ = v87
                                                    }
                                                case "\n\t":
                                                    v90 := GenS12{F0: "hello", F1: 739, F2: v70}
                                                    _ = v90
                                                case "hello":
                                                    var v91 uint = v22
                                                    _ = v91
                                                }
                                            }
                                            var v92 uint16 = 0
                                            v93 := GenS13{F0: 42}
                                            _ = v92
                                            _ = v93
                                        }
                                        switch "a" {
                                        case "a":
                                            var v94 bool = false
                                            v4 = (v4 || (v94 || ("hello" >= "")))
                                            _ = v94
                                        case "hello":
                                            v69 = (v69 * 9223372036854775807)
                                        case "foo bar":
                                            var v95 int8 = 127
                                            _ = v95
                                        }
                                        if ((!(v4 && ((((((1.5 * (0.0 - 0.1)) + 1.5) + 0.1) * ((1e10 / -2.25) + 3.14159)) / 3.14159) >= (0.1 / 3.14159)))) && (!(!true))) {
                                            fmt.Println(v4)
                                            v70 += (v70 ^ v70)
                                        }
                                    } else {
                                        v69 = 321
                                    }
                                case 1:
                                    _ = len(v2)
                                }
                            }
                        default:
                            if true {
                                v4 = (1 != -1)
                                switch "foo bar" {
                                case "a":
                                    if (v70 != v70) {
                                        var v96 string = "a"
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v96
                                    } else {
                                        if v4 {
                                            v68 += v68
                                            if v1 {
                                                v2 = append(v2, (^(v69)))
                                                if len(v5) > 0 {
                                                    v5[0] = fi23
                                                }
                                                if (v4 && false) {
                                                    if ("" != "hello") {
                                                        v3 -= (^(v3))
                                                        switch (fi21 + (fi23 % (((fi23 / (((v69 << (((v22 >> (((v22 / (((v22 % (((v22 << ((v22) & 63))) | 1))) | 1))) & 63))) & 63))) | 1))) | 1))) {
                                                        case 42:
                                                            var v97 int = (fi23 ^ (v69 * fi21))
                                                            _ = v97
                                                        case -558:
                                                            fmt.Println(v24, v69, v3)
                                                        }
                                                        fmt.Println(v69, v70)
                                                    } else {
                                                        v5 = append(v5, (v69 - (fi21 * (fi21 << ((v22) & 63)))))
                                                        fmt.Println(v24, v1)
                                                    }
                                                    v98 := GenS10{F0: fi23}
                                                    if (385 != 620) {
                                                        v1 = v1
                                                        fmt.Println(v3)
                                                    } else {
                                                        v68 = -9223372036854775808
                                                        var v99 int8 = 100
                                                        fmt.Println(v68, v24)
                                                        _ = v99
                                                    }
                                                    var v100 float32 = 0.1
                                                    if true {
                                                        fmt.Println(v22)
                                                    }
                                                    _ = v100
                                                    _ = v98
                                                } else {
                                                    v101 := GenS14{F0: (v68 <= (-(v68)))}
                                                    fmt.Println(v4)
                                                    var v102 uint = (v22 & v22)
                                                    var v103 int = fi23
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v101
                                                    _ = v102
                                                    _ = v103
                                                }
                                            } else {
                                                if (100 >= -119) {
                                                    var v104 uint64 = 0
                                                    v105 := GenS15{F0: 2, F1: ((false || (!false)) && (v3 >= v3))}
                                                    if v1 {
                                                        switch fi23 {
                                                        case -9223372036854775808:
                                                            v5 = append(v5, fi21)
                                                        case -466:
                                                            switch -1 {
                                                            case -321:
                                                                switch "a" {
                                                                case "a":
                                                                    var v106 float64 = v24
                                                                    _ = v106
                                                                case "foo bar":
                                                                    fmt.Println(v70, v22, v70)
                                                                }
                                                            case -2147483648:
                                                                var v107 float64 = 0.0
                                                                _ = v107
                                                            case 32767:
                                                                var v108 int = v69
                                                                _ = v108
                                                            default:
                                                                v69 -= (fi23 / ((-229) | 1))
                                                            }
                                                        }
                                                    } else {
                                                        switch v69 {
                                                        case 4294967295:
                                                            var v109 uint8 = v3
                                                            _ = v109
                                                        case 1:
                                                            fmt.Println(v24, fi23)
                                                        case 42:
                                                            var v110 uint8 = v3
                                                            _ = v110
                                                        default:
                                                            var v111 uint = 18446744073709551615
                                                            _ = v111
                                                        }
                                                        switch fi23 {
                                                        case 9223372036854775807:
                                                            _ = len(v5)
                                                        case -88:
                                                            switch "a" {
                                                            case "foo bar":
                                                                if false {
                                                                    v24 = ((v24 * v24) / -2.25)
                                                                }
                                                            case "\n\t":
                                                                var v112 uint16 = 580
                                                                _ = v112
                                                            default:
                                                                if (!(!v4)) {
                                                                    v1 = v4
                                                                    v5 = append(v5, (fi23 | (fi21 % ((fi23) | 1))))
                                                                    v113 := GenS7{F0: "a", F1: ("" + ("foo bar" + "a"))}
                                                                    _ = v113
                                                                }
                                                                fmt.Println(v68, v68)
                                                            }
                                                            switch ("a" + "") {
                                                            case "\n\t":
                                                                if v1 {
                                                                    switch ("" + "a") {
                                                                    case "":
                                                                        v114 := GenS11{F0: uint8(v22)}
                                                                        _ = v114
                                                                    case "hello":
                                                                        if len(v2) > 0 {
                                                                            v2[0] = -1
                                                                        }
                                                                    default:
                                                                        switch v69 {
                                                                        case -9223372036854775808:
                                                                            v104 = (v104 * 258)
                                                                            var v115 float32 = 0.1
                                                                            _ = v115
                                                                        case 32768:
                                                                            _ = len(v5)
                                                                        case -2147483648:
                                                                            if v4 {
                                                                                v70 += v70
                                                                                if v4 {
                                                                                    var v116 float32 = float32(3.14159)
                                                                                    v117 := GenS11{F0: v3}
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        if (127 >= 42) {
                                                                                            v68 += v68
                                                                                            if (!(42 <= 1)) {
                                                                                                if v4 {
                                                                                                    v1 = v1
                                                                                                } else {
                                                                                                    var v118 bool = (629 < 655)
                                                                                                    if (("\n\t" >= "a") && (!v118)) {
                                                                                                        fmt.Println(fi23)
                                                                                                        switch 0 {
                                                                                                        case 4294967295:
                                                                                                            if ((v68 & v68) < (v68 + (v68 | 735))) {
                                                                                                                var v119 float32 = v116
                                                                                                                fmt.Println(v1)
                                                                                                                var v120 uint8 = (v3 / ((v3) | 1))
                                                                                                                _ = v119
                                                                                                                _ = v120
                                                                                                            } else {
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = 32768
                                                                                                                }
                                                                                                                v118 = true
                                                                                                                v68 += -1
                                                                                                            }
                                                                                                        case 9223372036854775807:
                                                                                                            var v121 int32 = (-(v70))
                                                                                                            _ = v121
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v122 uint = (v22 & 440)
                                                                                                        v123 := GenS16{F0: (v68 ^ v68), F1: 1}
                                                                                                        var v124 int32 = (v70 ^ v70)
                                                                                                        _ = v122
                                                                                                        _ = v123
                                                                                                        _ = v124
                                                                                                    }
                                                                                                    fmt.Println(v1, v69, v116)
                                                                                                    v24 = (1e10 / v24)
                                                                                                    _ = v118
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v2 = append(v2, fi21)
                                                                                            var v125 float32 = v116
                                                                                            _ = v125
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            var v126 int = fi23
                                                                                            switch ("" + "") {
                                                                                            case "foo bar":
                                                                                                var v127 uint64 = v104
                                                                                                _ = v127
                                                                                            }
                                                                                            var v128 float64 = v24
                                                                                            _ = v126
                                                                                            _ = v128
                                                                                        } else {
                                                                                            var v129 uint32 = 2147483647
                                                                                            _ = v129
                                                                                        }
                                                                                    case "":
                                                                                        v2 = append(v2, fi23)
                                                                                    }
                                                                                    _ = v116
                                                                                    _ = v117
                                                                                }
                                                                                v130 := GenS17{F0: 42, F1: 2}
                                                                                _ = v130
                                                                            }
                                                                        }
                                                                    }
                                                                    v69 += v69
                                                                    v70 = (v70 + (v70 ^ v70))
                                                                    _ = len(v2)
                                                                } else {
                                                                    var v131 int64 = (v68 * v68)
                                                                    _ = v131
                                                                }
                                                            case "foo bar":
                                                                if v1 {
                                                                    fmt.Println(v3)
                                                                    v24 -= v24
                                                                    switch fi23 {
                                                                    case 4294967295:
                                                                        if len(v5) > 0 {
                                                                            v5[0] = v69
                                                                        }
                                                                        var v132 float64 = (v24 * v24)
                                                                        v133 := GenS10{F0: 0}
                                                                        _ = v132
                                                                        _ = v133
                                                                    case 1:
                                                                        v134 := GenS18{F0: 2147483647, F1: v3, F2: (v3 + (^(v3)))}
                                                                        _ = v134
                                                                    }
                                                                } else {
                                                                    var v135 uint16 = 1
                                                                    if true {
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            if (true || (v4 && (v24 != v24))) {
                                                                                var v136 bool = v4
                                                                                v4 = v4
                                                                                v137 := GenS12{F0: "\n\t", F1: 2147483648, F2: v70}
                                                                                _ = v136
                                                                                _ = v137
                                                                            } else {
                                                                                var v138 uint16 = (v135 % (((v135 - (v135 << ((v22) & 63)))) | 1))
                                                                                if v1 {
                                                                                    if v4 {
                                                                                        if (((1e10 - ((float32((-2.25 + (v24 * 3.14159))) / 1.5) / 0.1)) / 0.1) == 3.14159) {
                                                                                            var v139 uint = 65535
                                                                                            _ = v139
                                                                                        } else {
                                                                                            v104 = v104
                                                                                            switch "" {
                                                                                            case "a":
                                                                                                if v1 {
                                                                                                    var v140 uint16 = (v135 / ((v135) | 1))
                                                                                                    var v141 float64 = v24
                                                                                                    _ = v140
                                                                                                    _ = v141
                                                                                                }
                                                                                            case "hello":
                                                                                                var v142 uint64 = v104
                                                                                                v143 := GenS6{F0: 128, F1: ("a" + ""), F2: 42}
                                                                                                _ = v142
                                                                                                _ = v143
                                                                                            case "foo bar":
                                                                                                v1 = v1
                                                                                            default:
                                                                                                fmt.Println(v68, v70)
                                                                                            }
                                                                                        }
                                                                                        v1 = v4
                                                                                    }
                                                                                    v3 = 75
                                                                                    v144 := GenS2{F0: v138}
                                                                                    _ = v144
                                                                                }
                                                                                if true {
                                                                                    switch ("hello" + "a") {
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            fmt.Println(fi21)
                                                                                            v145 := GenS19{F0: 2147483647, F1: uint16(v104), F2: 157}
                                                                                            _ = v145
                                                                                        } else {
                                                                                            v146 := GenS20{F0: (v3 ^ (^(v3)))}
                                                                                            var v147 uint16 = v138
                                                                                            v148 := GenS16{F0: (v68 + (-(v68))), F1: 0}
                                                                                            v68 -= (v68 << (((v22 * (v22 >> (((v22 * v22)) & 63)))) & 63))
                                                                                            _ = v146
                                                                                            _ = v147
                                                                                            _ = v148
                                                                                        }
                                                                                        var v149 int8 = 1
                                                                                        fmt.Println(v104)
                                                                                        _ = v149
                                                                                    default:
                                                                                        var v150 int64 = v68
                                                                                        _ = v150
                                                                                    }
                                                                                    if (-128 == -8) {
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = (v69 + (fi21 << ((42) & 63)))
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch "a" {
                                                                                case "a":
                                                                                    var v151 bool = (false && ((^(v69)) == -2147483648))
                                                                                    _ = v151
                                                                                case "hello":
                                                                                    var v152 float64 = (3.14159 / 1.5)
                                                                                    _ = v152
                                                                                }
                                                                                v153 := GenS16{F0: v68, F1: (v138 + v135)}
                                                                                _ = v138
                                                                                _ = v153
                                                                            }
                                                                        case "a":
                                                                            var v154 int = fi21
                                                                            v155 := GenS3{F0: ("a" + "\n\t")}
                                                                            var v156 int32 = 660
                                                                            _ = v154
                                                                            _ = v155
                                                                            _ = v156
                                                                        case "foo bar":
                                                                            switch (("" + ((("hello" + "a") + "hello") + "a")) + ("" + (("\n\t" + "\n\t") + ""))) {
                                                                            case "\n\t":
                                                                                var v157 int32 = v70
                                                                                _ = v157
                                                                            }
                                                                        default:
                                                                            var v158 float64 = (((v24 - v24) * v24) / 3.14159)
                                                                            _ = v158
                                                                        }
                                                                        switch ("foo bar" + ("a" + "\n\t")) {
                                                                        case "hello":
                                                                            var v159 int8 = 0
                                                                            _ = v159
                                                                        case "a":
                                                                            v4 = v1
                                                                        case "":
                                                                            fmt.Println(v22, v70, v4)
                                                                        }
                                                                    } else {
                                                                        var v160 int32 = v70
                                                                        _ = v160
                                                                    }
                                                                    _ = v135
                                                                }
                                                            }
                                                        case -1:
                                                            switch "a" {
                                                            case "\n\t":
                                                                if (v68 == 865) {
                                                                    if (!true) {
                                                                        v161 := GenS21{F0: -1}
                                                                        var v162 uint8 = (v3 ^ v3)
                                                                        if true {
                                                                            v69 = (fi21 - 1)
                                                                            var v163 int64 = (v68 << ((v22) & 63))
                                                                            _ = v163
                                                                        } else {
                                                                            v162 += v3
                                                                            if v1 {
                                                                                var v164 uint32 = 65535
                                                                                v22 += v22
                                                                                v165 := GenS22{F0: uint(v162), F1: (^(v22)), F2: ""}
                                                                                _ = v164
                                                                                _ = v165
                                                                            } else {
                                                                                fmt.Println(v4, v69)
                                                                                var v166 int32 = (v70 << ((2) & 63))
                                                                                v4 = (651 != genFn0((42 != -1), 921, 2))
                                                                                v167 := GenS4{F0: (v166 ^ (v166 ^ v70))}
                                                                                _ = v166
                                                                                _ = v167
                                                                            }
                                                                        }
                                                                        _ = v161
                                                                        _ = v162
                                                                    }
                                                                    v3 += 228
                                                                    if ((fi21 + fi23) <= (^(fi21))) {
                                                                        v69 = fi21
                                                                    }
                                                                } else {
                                                                    v168 := GenS23{F0: -128}
                                                                    _ = v168
                                                                }
                                                                if true {
                                                                    var v169 float64 = v24
                                                                    v5 = append(v5, v69)
                                                                    var v170 uint8 = 39
                                                                    _ = v169
                                                                    _ = v170
                                                                } else {
                                                                    if (v3 > 0) {
                                                                        fmt.Println(v70, v3)
                                                                        fmt.Println(v104, fi23)
                                                                        var v171 int8 = 42
                                                                        fmt.Println(v70, v4, v3)
                                                                        v3 = v3
                                                                        _ = v171
                                                                    }
                                                                    if v1 {
                                                                        if true {
                                                                            _ = len(v2)
                                                                            v69 += (v69 / (((v69 - 9223372036854775807)) | 1))
                                                                            v69 -= -2147483648
                                                                            v172 := GenS24{F0: (v3 | (v3 >> ((v22) & 63))), F1: v4, F2: (fi21 & (fi23 + fi23))}
                                                                            var v173 uint = (v22 - v22)
                                                                            _ = v172
                                                                            _ = v173
                                                                        } else {
                                                                            v174 := GenS25{F0: v22, F1: 2147483648}
                                                                            if ("" >= "hello") {
                                                                                fmt.Println(v70)
                                                                                var v175 float32 = -2.25
                                                                                _ = v175
                                                                            } else {
                                                                                v3 = 20
                                                                                switch fi21 {
                                                                                case 32768:
                                                                                    v176 := GenS23{F0: -30}
                                                                                    _ = v176
                                                                                case 495:
                                                                                    if v1 {
                                                                                        v69 = v69
                                                                                        var v177 float64 = v24
                                                                                        _ = v177
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (!(!v4)) {
                                                                                fmt.Println(fi21, v68, fi21)
                                                                            } else {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (v69 / ((4294967295) | 1))
                                                                                }
                                                                                v68 -= 9223372036854775807
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                switch -2147483648 {
                                                                                case 2147483647:
                                                                                    var v178 int64 = (^(v68))
                                                                                    _ = v178
                                                                                case -954:
                                                                                    var v179 bool = v4
                                                                                    _ = v179
                                                                                case -9223372036854775808:
                                                                                    v3 = (v3 + (v3 | (^(v3))))
                                                                                }
                                                                            case "a":
                                                                                var v180 string = "a"
                                                                                _ = v180
                                                                            case "":
                                                                                fmt.Println(v70)
                                                                            }
                                                                            _ = v174
                                                                        }
                                                                        if ((v1 && v1) || false) {
                                                                            _ = len(v5)
                                                                            v68 = v68
                                                                        } else {
                                                                            if true {
                                                                                _ = len(v2)
                                                                            } else {
                                                                                v181 := GenS26{F0: (v22 - (v22 << ((42) & 63)))}
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = v181
                                                                            }
                                                                            switch 32767 {
                                                                            case 0:
                                                                                v182 := GenS18{F0: 42, F1: v3, F2: (v3 / ((v3) | 1))}
                                                                                _ = v182
                                                                            case -2147483648:
                                                                                var v183 uint = v22
                                                                                _ = v183
                                                                            case -9223372036854775808:
                                                                                var v184 uint8 = v3
                                                                                _ = v184
                                                                            }
                                                                        }
                                                                        v185 := GenS17{F0: 2, F1: 255}
                                                                        _ = v185
                                                                    }
                                                                    fmt.Println(fi23, v22, v104)
                                                                    var v186 bool = ((int32(v68) == v70) && (v22 > (v22 / ((v22) | 1))))
                                                                    _ = v186
                                                                }
                                                            case "foo bar":
                                                                v68 += v68
                                                            case "":
                                                                switch fi21 {
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        v2 = append(v2, v69)
                                                                    } else {
                                                                        var v187 uint32 = 0
                                                                        v22 = 227
                                                                        if false {
                                                                            v1 = v4
                                                                            v3 = 217
                                                                            fmt.Println(v69, v68, v70)
                                                                        } else {
                                                                            switch (v69 & (fi21 ^ (fi23 + fi23))) {
                                                                            case 842:
                                                                                switch -833 {
                                                                                case -9223372036854775808:
                                                                                    var v188 uint8 = (v3 - uint8(v22))
                                                                                    _ = v188
                                                                                }
                                                                            case 401:
                                                                                if ("a" < "") {
                                                                                    v22 = v22
                                                                                    var v189 string = "\n\t"
                                                                                    _ = v189
                                                                                } else {
                                                                                    var v190 int16 = -426
                                                                                    _ = v190
                                                                                }
                                                                            default:
                                                                                v104 = (v104 >> ((v22) & 63))
                                                                            }
                                                                            var v191 string = "foo bar"
                                                                            _ = v191
                                                                        }
                                                                        var v192 int8 = -1
                                                                        _ = v187
                                                                        _ = v192
                                                                    }
                                                                case 127:
                                                                    v1 = (!true)
                                                                }
                                                            default:
                                                                fmt.Println(v24, fi23)
                                                            }
                                                        default:
                                                            v193 := GenS27{F0: -79, F1: 0}
                                                            _ = v193
                                                        }
                                                        v194 := GenS2{F0: 0}
                                                        _ = v194
                                                    }
                                                    _ = v104
                                                    _ = v105
                                                }
                                                _ = len(v5)
                                                v3 = 128
                                            }
                                        } else {
                                            fmt.Println(v70, v68, v24)
                                        }
                                    }
                                case "hello":
                                    v195 := GenS28{F0: 441, F1: 65535}
                                    _ = v195
                                case "":
                                    switch int(v22) {
                                    case 127:
                                        if v1 {
                                            switch 577 {
                                            case -858:
                                                v5 = append(v5, (v69 | 9223372036854775807))
                                            default:
                                                fmt.Println(v4, v3, fi21)
                                            }
                                            var v196 string = ("hello" + "\n\t")
                                            if true {
                                                _ = len(v2)
                                                v197 := GenS29{F0: 1}
                                                v3 = 207
                                                _ = v197
                                            }
                                            _ = v196
                                        }
                                    case 32768:
                                        if (3.14159 >= float32((v24 / ((-2.25 / (v24 - v24)) + v24)))) {
                                            var v198 float64 = v24
                                            _ = v198
                                        } else {
                                            fmt.Println(v3, v69)
                                        }
                                    case -235:
                                        switch "foo bar" {
                                        case "":
                                            var v199 uint = v22
                                            if len(v5) > 0 {
                                                v5[0] = 127
                                            }
                                            var v200 uint16 = uint16(v199)
                                            _ = v199
                                            _ = v200
                                        case "a":
                                            v1 = v4
                                        case "\n\t":
                                            switch v69 {
                                            case 189:
                                                var v201 uint = v22
                                                _ = v201
                                            default:
                                                v202 := GenS30{F0: (v3 ^ (^(v3))), F1: "\n\t", F2: -1}
                                                _ = v202
                                            }
                                        default:
                                            v5 = append(v5, v69)
                                            if ((false && ((!v4) && v1)) && (-1 >= (v68 % ((v68) | 1)))) {
                                                if (v22 == (v22 % ((v22) | 1))) {
                                                    var v203 int16 = 42
                                                    fmt.Println(fi21, v24)
                                                    v68 -= (v68 >> (((v22 ^ v22)) & 63))
                                                    _ = v203
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v204 uint = 42
                                                    _ = v204
                                                }
                                                if ((v68 >> ((v22) & 63)) > int64(v70)) {
                                                    if v4 {
                                                        v4 = v1
                                                        var v205 float32 = 0.1
                                                        v4 = v1
                                                        _ = v205
                                                    } else {
                                                        if ("hello" > ("\n\t" + ("foo bar" + "foo bar"))) {
                                                            var v206 int64 = v68
                                                            switch "\n\t" {
                                                            case "a":
                                                                if v4 {
                                                                    var v207 int32 = -168
                                                                    _ = v207
                                                                }
                                                            case "\n\t":
                                                                var v208 int64 = v206
                                                                _ = v208
                                                            case "":
                                                                var v209 float64 = v24
                                                                _ = v209
                                                            }
                                                            v210 := GenS23{F0: -1}
                                                            _ = v206
                                                            _ = v210
                                                        }
                                                    }
                                                    v2 = append(v2, (fi21 << (((^(v22))) & 63)))
                                                    var v211 float64 = v24
                                                    _ = v211
                                                } else {
                                                    var v212 uint16 = 1
                                                    _ = v212
                                                }
                                            } else {
                                                var v213 uint = 42
                                                _ = v213
                                            }
                                            switch v69 {
                                            case 9223372036854775807:
                                                var v214 float64 = (v24 + v24)
                                                _ = v214
                                            case 42:
                                                v215 := GenS31{F0: 2}
                                                _ = v215
                                            case -56:
                                                fmt.Println(v3)
                                            }
                                        }
                                    }
                                }
                                var v216 uint32 = 4294967295
                                _ = v216
                            } else {
                                var v217 int = 1
                                v218 := GenS9{F0: 2, F1: genFn0(((v1 && v4) || (0 == 2)), v70, 42)}
                                if len(v2) > 0 {
                                    v2[0] = v69
                                }
                                _ = v217
                                _ = v218
                            }
                        }
                        _ = v70
                    } else {
                        if v4 {
                            _ = len(v2)
                            if v4 {
                                v219 := GenS11{F0: (v3 * 0)}
                                var v220 uint8 = v3
                                _ = v219
                                _ = v220
                            }
                            v22 = 65535
                            var v221 bool = true
                            _ = v221
                        }
                    }
                default:
                    fmt.Println(v68, v68)
                }
                _ = v68
                _ = v69
            }
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    var v222 uint32 = 326
    for fi223 := 0; fi223 < 20; fi223++ {
        for fi224 := 0; fi224 < 5; fi224++ {
            if len(v2) > 0 {
                v2[0] = (fi224 - fi224)
            }
            v1 = v1
            _ = len(v2)
            _ = fi224
        }
        var v225 float32 = 1e10
        v1 = (!false)
        _ = fi223
        _ = v225
    }
    var v226 int32 = 2147483647
    if (true || ((v226 | (v226 | (v226 & (v226 << ((2) & 63))))) <= v226)) {
        for fi227 := 0; fi227 < 15; fi227++ {
            for fi228 := 0; fi228 < 4; fi228++ {
                if len(v2) > 0 {
                    v2[0] = (fi227 << ((95) & 63))
                }
                if ((v222 | v222) > v222) {
                    v4 = false
                    v229 := GenS32{F0: 0, F1: (v226 + 2147483647), F2: 128}
                    var v230 float64 = 1.5
                    _ = v229
                    _ = v230
                }
                v4 = v1
                _ = fi228
            }
            _ = fi227
        }
        v222 = v222
    } else {
        for fi231 := 0; fi231 < 1; fi231++ {
            v4 = (816 >= 0)
            _ = fi231
        }
        if (-839 == 2147483647) {
            var v232 uint64 = uint64(v3)
            if false {
                v232 -= (v232 | v232)
                for fi233 := 0; fi233 < 8; fi233++ {
                    var v234 float64 = 1e10
                    fmt.Println(v226, fi233)
                    for fi235 := 0; fi235 < 5; fi235++ {
                        v226 -= v226
                        _ = fi235
                    }
                    _ = fi233
                    _ = v234
                }
                if (-32768 >= 42) {
                    if (-949 > 42) {
                        for fi236 := 0; fi236 < 20; fi236++ {
                            fmt.Println(v226, v232, fi236)
                            if ((v226 % ((-2147483648) | 1)) != (v226 / ((170) | 1))) {
                                var v237 uint32 = 1
                                for fi238 := 0; fi238 < 19; fi238++ {
                                    if v4 {
                                        fmt.Println(v237, v4)
                                        v2 = append(v2, 32768)
                                    } else {
                                        fmt.Println(v4, v4, fi236)
                                        var v239 int64 = 1
                                        v222 += 527
                                        _ = v239
                                    }
                                    if false {
                                        fmt.Println(v1, v232)
                                    }
                                    v237 += uint32(v232)
                                    _ = fi238
                                }
                                _ = v237
                            } else {
                                v240 := GenS6{F0: 255, F1: "\n\t", F2: 1}
                                switch (("hello" + "a") + "a") {
                                case "foo bar":
                                    v241 := GenS33{F0: (fi236 | fi236), F1: v232, F2: 42}
                                    _ = v241
                                }
                                _ = v240
                            }
                            _ = fi236
                        }
                        var v242 int8 = 1
                        _ = v242
                    } else {
                        v243 := GenS34{F0: -27, F1: v3}
                        _ = v243
                    }
                    if (!v4) {
                        var v244 int16 = 42
                        _ = v244
                    } else {
                        for fi245 := 0; fi245 < 16; fi245++ {
                            var v246 int16 = 0
                            v247 := GenS7{F0: "a", F1: (("a" + ((("" + ("hello" + ((("foo bar" + "foo bar") + "") + "hello"))) + "\n\t") + ("a" + "a"))) + "a")}
                            var v248 int16 = 32767
                            _ = fi245
                            _ = v246
                            _ = v247
                            _ = v248
                        }
                        var v249 float32 = (1e10 - -2.25)
                        _ = v249
                    }
                    if false {
                        var v250 bool = v1
                        fmt.Println(v222, v222)
                        var v251 int64 = 1
                        _ = v250
                        _ = v251
                    } else {
                        v232 = 427
                        var v252 int8 = 127
                        _ = v252
                    }
                } else {
                    v232 = (^(v232))
                }
            }
            var v253 bool = v1
            _ = v232
            _ = v253
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v222:", v222)
    fmt.Println("v226:", v226)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v222
    _ = v226
    _ = v3
    _ = v4
    _ = v5
}
