package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 int64; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int; F1 string; F2 uint }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS15 struct { F0 bool; F1 int64 }

func main() {
    var v1 int64 = 0
    var v2 uint32 = 1
    var v3 int32 = -1
    var v4 uint16 = 255
    switch "" {
    case "\n\t":
        var v5 int64 = 42
        _ = v5
    case "":
        v6 := GenS1{F0: 0}
        _ = v6
    case "a":
        v2 -= v2
    }
    var v7 uint32 = 452
    v8 := []int{32767, 855}
    if (-1 > (-(v1))) {
        if false {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 int64 = (^(v1))
                _ = len(v8)
                _ = fi9
                _ = v10
            }
            v2 -= (v7 % (((v7 ^ 4294967295)) | 1))
            var v11 uint16 = v4
            _ = v11
        }
        var v12 int32 = -2147483648
        _ = v12
    } else {
        v7 = (v2 ^ (v2 << ((18446744073709551615) & 63)))
        var v13 int16 = 534
        v2 += 389
        _ = v13
    }
    var v14 float32 = (3.14159 * 3.14159)
    var v15 int16 = -1
    for fi16 := 0; fi16 < 12; fi16++ {
        v3 += (^(v3))
        var v17 uint64 = 18446744073709551615
        _ = fi16
        _ = v17
    }
    v15 = -1
    if true {
        var v18 bool = false
        _ = v18
    }
    switch -9223372036854775808 {
    case -617:
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v3)
            if (540 < 2) {
                v14 = v14
                fmt.Println(v2)
            } else {
                switch fi19 {
                case 32767:
                    _ = len(v8)
                case 0:
                    if true {
                        if false {
                            for fi20 := 0; fi20 < 17; fi20++ {
                                if (v2 > (v7 + (^(v7)))) {
                                    if (true && false) {
                                        switch (fi19 % (((-(fi19))) | 1)) {
                                        case -627:
                                            if ((-(v15)) == (v15 << ((46) & 63))) {
                                                if false {
                                                    v3 -= 0
                                                    v21 := GenS2{F0: 111, F1: (fi20 & 9223372036854775807)}
                                                    v14 = (float32(((1.5 - 0.1) / 3.14159)) * (v14 - v14))
                                                    _ = v21
                                                } else {
                                                    v8 = append(v8, int(v3))
                                                    v7 = v7
                                                    if ((("\n\t" + "a") + ((("\n\t" + ("foo bar" + "hello")) + (("\n\t" + "a") + "")) + "a")) != ("" + "foo bar")) {
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    } else {
                                                        if false {
                                                            var v22 uint16 = (v4 - (v4 & (v4 - 1)))
                                                            var v23 uint8 = 0
                                                            var v24 float64 = (-2.25 * 1e10)
                                                            _ = v22
                                                            _ = v23
                                                            _ = v24
                                                        } else {
                                                            var v25 float32 = 0.1
                                                            _ = v25
                                                        }
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v26 uint = 42
                                                            _ = v26
                                                        }
                                                        v2 = (^(v2))
                                                    }
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                }
                                                _ = len(v8)
                                                _ = len(v8)
                                            } else {
                                                if false {
                                                    switch fi20 {
                                                    case 9223372036854775807:
                                                        if true {
                                                            var v27 float32 = 0.1
                                                            var v28 int16 = v15
                                                            fmt.Println(v14, v4, v28)
                                                            _ = v27
                                                            _ = v28
                                                        } else {
                                                            v14 = (-2.25 * v14)
                                                        }
                                                    default:
                                                        if false {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                var v29 uint32 = (v2 / ((v7) | 1))
                                                                _ = v29
                                                            }
                                                            v15 += v15
                                                            fmt.Println(v14)
                                                        } else {
                                                            var v30 int = 127
                                                            if false {
                                                                v31 := GenS3{F0: (fi19 & (v30 >> ((0) & 63))), F1: uint32(v4), F2: 4294967295}
                                                                fmt.Println(v15)
                                                                _ = v31
                                                            } else {
                                                                if true {
                                                                    var v32 float32 = (0.1 / 1e10)
                                                                    var v33 uint32 = v2
                                                                    var v34 float32 = 0.1
                                                                    _ = v32
                                                                    _ = v33
                                                                    _ = v34
                                                                }
                                                                v2 += 2147483647
                                                                v8 = append(v8, 9223372036854775807)
                                                            }
                                                            _ = v30
                                                        }
                                                    }
                                                }
                                                var v35 int8 = 1
                                                if true {
                                                    v14 = float32(((1.5 + -2.25) / 1.0))
                                                    v14 = 0.0
                                                    var v36 uint16 = 0
                                                    _ = v36
                                                } else {
                                                    v37 := GenS1{F0: 673}
                                                    if (!true) {
                                                        v38 := GenS4{F0: v1}
                                                        switch "" {
                                                        case "a":
                                                            switch fi19 {
                                                            case 324:
                                                                v7 = 328
                                                            case 0:
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                            }
                                                        case "foo bar":
                                                            if true {
                                                                v14 = (3.14159 - (v14 + (v14 + 0.1)))
                                                                v2 = uint32(v4)
                                                            }
                                                        case "":
                                                            v1 = -9223372036854775808
                                                        default:
                                                            var v39 float32 = (v14 + (float32((1.5 - (float64(v14) + 0.1))) * (0.1 + v14)))
                                                            _ = v39
                                                        }
                                                        fmt.Println(v15, v3, v1)
                                                        switch ("" + ("foo bar" + "")) {
                                                        case "a":
                                                            var v40 string = ((("\n\t" + ("hello" + ((("foo bar" + (((("" + "foo bar") + ("" + "a")) + "\n\t") + ("a" + "hello"))) + ("" + ("" + ""))) + ("foo bar" + "a")))) + "\n\t") + "\n\t")
                                                            if true {
                                                                var v41 int32 = 2147483647
                                                                _ = v41
                                                            } else {
                                                                switch v40 {
                                                                case "hello":
                                                                    _ = len(v8)
                                                                case "\n\t":
                                                                    if (false && false) {
                                                                        v1 = (v1 & (v1 | 1))
                                                                        var v42 float64 = 0.1
                                                                        _ = v42
                                                                    } else {
                                                                        v43 := GenS5{F0: 65535, F1: (v4 % (((v4 / (((^(v4))) | 1))) | 1))}
                                                                        var v44 string = "foo bar"
                                                                        var v45 int64 = v1
                                                                        if len(v8) > 0 {
                                                                            v8[0] = 32767
                                                                        }
                                                                        _ = v43
                                                                        _ = v44
                                                                        _ = v45
                                                                    }
                                                                case "a":
                                                                    if true {
                                                                        var v46 uint16 = 2
                                                                        v46 = 429
                                                                        _ = v46
                                                                    } else {
                                                                        switch fi20 {
                                                                        case 0:
                                                                            if false {
                                                                                v8 = append(v8, fi20)
                                                                                switch int(v35) {
                                                                                case 2147483647:
                                                                                    switch fi20 {
                                                                                    case -1:
                                                                                        if len(v8) >= 2 {
                                                                                            v8 = v8[1:len(v8)]
                                                                                        }
                                                                                    case -185:
                                                                                        fmt.Println(v1, v40)
                                                                                    }
                                                                                    var v47 uint64 = 18446744073709551615
                                                                                    var v48 int32 = (v3 * v3)
                                                                                    if len(v8) >= 2 {
                                                                                        v8 = v8[1:len(v8)]
                                                                                    }
                                                                                    fmt.Println(v40, v2)
                                                                                    if true {
                                                                                        var v49 uint8 = 0
                                                                                        _ = v49
                                                                                    } else {
                                                                                        switch v40 {
                                                                                        case "foo bar":
                                                                                            v50 := GenS6{F0: (((((("" + ((("hello" + ((("\n\t" + v40) + (v40 + (v40 + v40))) + v40)) + v40) + (v40 + v40))) + v40) + (v40 + ("" + "hello"))) + "hello") + v40) + v40)}
                                                                                            _ = v50
                                                                                        }
                                                                                    }
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                }
                                                                                _ = len(v8)
                                                                                v2 = 42
                                                                                var v51 uint16 = (v4 - (v4 * 42))
                                                                                _ = v51
                                                                            }
                                                                        default:
                                                                            _ = len(v8)
                                                                        }
                                                                        if (!true) {
                                                                            switch (fi20 + -2147483648) {
                                                                            case 29:
                                                                                if ((fi20 ^ fi19) >= fi20) {
                                                                                    v3 -= v3
                                                                                    v8 = append(v8, (fi19 << ((340) & 63)))
                                                                                    if (42 != v4) {
                                                                                        fmt.Println(v3)
                                                                                        v52 := GenS7{F0: (v7 | (v7 << ((0) & 63))), F1: v1, F2: v15}
                                                                                        _ = v52
                                                                                    } else {
                                                                                        if (!(!false)) {
                                                                                            v8 = append(v8, (fi19 << ((1) & 63)))
                                                                                        } else {
                                                                                            v4 = v4
                                                                                        }
                                                                                        v53 := GenS1{F0: 1}
                                                                                        var v54 int = (fi20 % ((32767) | 1))
                                                                                        v40 = ("\n\t" + v40)
                                                                                        v55 := GenS1{F0: 18446744073709551615}
                                                                                        var v56 int16 = (v15 - v15)
                                                                                        _ = len(v8)
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                    }
                                                                                }
                                                                            case -2147483648:
                                                                                if (1 <= 255) {
                                                                                    v57 := GenS8{F0: 94, F1: 18446744073709551615, F2: 127}
                                                                                    if false {
                                                                                        if (((float32((0.0 - (float64(v14) - 0.1))) * v14) - 1e10) != float32((float64(1e10) * 1e10))) {
                                                                                            var v58 bool = false
                                                                                            _ = len(v8)
                                                                                            _ = v58
                                                                                        }
                                                                                        var v59 uint16 = uint16(v2)
                                                                                        fmt.Println(v7, v15, v40)
                                                                                        _ = v59
                                                                                    } else {
                                                                                        v3 += (v3 + (v3 / ((int32(v15)) | 1)))
                                                                                        if (-1 >= 93) {
                                                                                            fmt.Println(fi20)
                                                                                        } else {
                                                                                            v4 += 394
                                                                                            v15 += v15
                                                                                            switch int(v15) {
                                                                                            case -1:
                                                                                                if (v7 >= v2) {
                                                                                                    var v60 int = fi19
                                                                                                    _ = len(v8)
                                                                                                    _ = v60
                                                                                                } else {
                                                                                                    v3 = v3
                                                                                                    v2 = (^(v7))
                                                                                                }
                                                                                            default:
                                                                                                if len(v8) >= 2 {
                                                                                                    v8 = v8[1:len(v8)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v61 int = (fi20 * (fi19 + fi19))
                                                                                        _ = v61
                                                                                    }
                                                                                    if len(v8) > 0 {
                                                                                        v8[0] = fi20
                                                                                    }
                                                                                    _ = v57
                                                                                } else {
                                                                                    var v62 uint32 = v7
                                                                                    var v63 int8 = 127
                                                                                    if false {
                                                                                        switch 42 {
                                                                                        case 2147483647:
                                                                                            var v64 float32 = v14
                                                                                            _ = v64
                                                                                        default:
                                                                                            fmt.Println(v3, v35)
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                v62 += (^(v62))
                                                                                            case "foo bar":
                                                                                                var v65 float64 = 1.5
                                                                                                _ = v65
                                                                                            case "a":
                                                                                                v8 = append(v8, fi19)
                                                                                            }
                                                                                        }
                                                                                        v66 := GenS9{F0: (v2 - 65535), F1: v3}
                                                                                        var v67 uint64 = 1
                                                                                        _ = len(v8)
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    } else {
                                                                                        if false {
                                                                                            fmt.Println(v4)
                                                                                            v4 += v4
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        var v68 int8 = v63
                                                                                        _ = v68
                                                                                    } else {
                                                                                        if true {
                                                                                            v15 = (v15 + v15)
                                                                                        } else {
                                                                                            v62 -= v2
                                                                                        }
                                                                                        v69 := GenS10{F0: uint(v7)}
                                                                                        _ = v69
                                                                                    }
                                                                                    if true {
                                                                                        _ = len(v8)
                                                                                    } else {
                                                                                        var v70 uint16 = (v4 ^ (^(v4)))
                                                                                        var v71 int64 = (v1 | int64(v4))
                                                                                        var v72 uint8 = 1
                                                                                        fmt.Println(v2, v62)
                                                                                        var v73 int8 = (^(v35))
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                        _ = v72
                                                                                        _ = v73
                                                                                    }
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                }
                                                                                v74 := GenS11{F0: (v4 % ((255) | 1))}
                                                                                var v75 int16 = v15
                                                                                _ = v74
                                                                                _ = v75
                                                                            default:
                                                                                if true {
                                                                                    if (!false) {
                                                                                        var v76 string = v40
                                                                                        if len(v8) > 0 {
                                                                                            v8[0] = (-(fi20))
                                                                                        }
                                                                                        v77 := GenS12{F0: 0, F1: v76, F2: 0}
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                    }
                                                                                    fmt.Println(fi20)
                                                                                    var v78 float32 = (3.14159 / 1.5)
                                                                                    var v79 string = "\n\t"
                                                                                    v79 = v79
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                        } else {
                                                                            _ = len(v8)
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                var v80 float64 = 1.5
                                                                                _ = v80
                                                                            }
                                                                            var v81 uint = 18446744073709551615
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                default:
                                                                    if (!(false && (v14 != v14))) {
                                                                        fmt.Println(v1, v14, fi20)
                                                                        var v82 uint = 99
                                                                        _ = v82
                                                                    }
                                                                    if len(v8) > 0 {
                                                                        v8[0] = fi20
                                                                    }
                                                                    v15 = (v15 ^ v15)
                                                                }
                                                            }
                                                            _ = v40
                                                        case "":
                                                            v83 := GenS10{F0: 1}
                                                            _ = v83
                                                        case "\n\t":
                                                            var v84 float64 = 0.1
                                                            _ = v84
                                                        default:
                                                            v15 += 28
                                                        }
                                                        var v85 bool = (!((false && (!(255 < 1))) || ((1 < 4294967295) && (false || true))))
                                                        _ = v38
                                                        _ = v85
                                                    } else {
                                                        v14 = (-2.25 + float32(-2.25))
                                                        fmt.Println(fi19, v4)
                                                    }
                                                    _ = v37
                                                }
                                                _ = v35
                                            }
                                        case 9223372036854775807:
                                            var v86 int32 = int32(v7)
                                            _ = v86
                                        case 0:
                                            if ((4294967295 > (v7 & (v7 | (^(v7))))) && false) {
                                                v87 := GenS13{F0: 655, F1: -128}
                                                v88 := GenS13{F0: 42, F1: 127}
                                                if false {
                                                    v89 := GenS14{F0: 237, F1: 1, F2: 42}
                                                    var v90 int = (-(fi20))
                                                    _ = v89
                                                    _ = v90
                                                } else {
                                                    if (v4 != v4) {
                                                        v7 -= (v2 % ((v2) | 1))
                                                    }
                                                    if false {
                                                        var v91 float64 = 0.0
                                                        if (v7 != v7) {
                                                            v92 := GenS4{F0: int64(fi19)}
                                                            fmt.Println(v91, v14, v15)
                                                            _ = v92
                                                        }
                                                        if false {
                                                            v93 := GenS11{F0: v4}
                                                            var v94 uint8 = 1
                                                            var v95 uint16 = v4
                                                            _ = v93
                                                            _ = v94
                                                            _ = v95
                                                        } else {
                                                            var v96 int8 = 100
                                                            var v97 uint64 = 353
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        _ = v91
                                                    } else {
                                                        v8 = append(v8, fi19)
                                                    }
                                                }
                                                _ = v87
                                                _ = v88
                                            } else {
                                                v7 = 42
                                                switch (fi20 << ((2) & 63)) {
                                                case 32768:
                                                    if len(v8) > 0 {
                                                        v8[0] = fi19
                                                    }
                                                default:
                                                    var v98 uint = 910
                                                    _ = v98
                                                }
                                                v14 = v14
                                            }
                                        }
                                    }
                                    var v99 uint8 = 255
                                    var v100 int = fi19
                                    _ = v100
                                    _ = v99
                                } else {
                                    if (!(true && ("" < ((("a" + "foo bar") + ("a" + "foo bar")) + ("hello" + "a"))))) {
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        switch (fi20 * fi20) {
                                        case 4294967295:
                                            v101 := GenS9{F0: v2, F1: (v3 * v3)}
                                            _ = v101
                                        case -9223372036854775808:
                                            fmt.Println(fi20)
                                        case 76:
                                            var v102 int32 = v3
                                            _ = v102
                                        default:
                                            switch fi19 {
                                            case 128:
                                                var v103 float64 = -2.25
                                                _ = v103
                                            case 1:
                                                if true {
                                                    var v104 uint8 = 2
                                                    _ = v104
                                                }
                                            case 4294967295:
                                                if ((v4 ^ v4) != v4) {
                                                    fmt.Println(fi20, v14, v1)
                                                    switch fi20 {
                                                    case 855:
                                                        var v105 uint64 = 4294967295
                                                        _ = v105
                                                    case 127:
                                                        v106 := GenS12{F0: fi20, F1: "\n\t", F2: 65535}
                                                        _ = v106
                                                    }
                                                } else {
                                                    if true {
                                                        fmt.Println(v3)
                                                        var v107 float64 = float64(v14)
                                                        v108 := GenS7{F0: 309, F1: v1, F2: (v15 & -32768)}
                                                        _ = v107
                                                        _ = v108
                                                    } else {
                                                        var v109 uint = 100
                                                        switch fi19 {
                                                        case -1:
                                                            _ = len(v8)
                                                        case 4294967295:
                                                            var v110 int64 = -1
                                                            _ = v110
                                                        }
                                                        v3 -= 0
                                                        _ = v109
                                                    }
                                                }
                                            }
                                        }
                                        if true {
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            var v111 uint8 = 1
                                            _ = v111
                                        } else {
                                            var v112 uint32 = 858
                                            _ = v112
                                        }
                                    }
                                }
                                v113 := GenS15{F0: true, F1: int64(v3)}
                                v4 += (v4 - v4)
                                v7 -= v7
                                var v114 uint = 1
                                _ = fi20
                                _ = v113
                                _ = v114
                            }
                        }
                        v115 := GenS10{F0: 848}
                        _ = v115
                    } else {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        for fi116 := 0; fi116 < 8; fi116++ {
                            if ((!(!true)) && (!false)) {
                                var v117 uint16 = (v4 & (v4 / (((v4 * (v4 & v4))) | 1)))
                                var v118 uint8 = 166
                                v1 = 1
                                _ = v117
                                _ = v118
                            } else {
                                v119 := GenS10{F0: 457}
                                v120 := GenS13{F0: uint64(v4), F1: 100}
                                var v121 bool = (v1 >= v1)
                                _ = v119
                                _ = v120
                                _ = v121
                            }
                            switch fi19 {
                            case 0:
                                var v122 int16 = (v15 / ((1) | 1))
                                _ = v122
                            case 9223372036854775807:
                                v1 = 42
                            case -2147483648:
                                switch "a" {
                                case "":
                                    var v123 uint32 = v2
                                    _ = v123
                                case "a":
                                    var v124 int16 = v15
                                    _ = v124
                                }
                            default:
                                if (v7 > 2147483647) {
                                    v125 := GenS10{F0: uint(v2)}
                                    v2 += 2147483648
                                    v7 = v7
                                    v126 := GenS9{F0: 2147483648, F1: (v3 | v3)}
                                    fmt.Println(v1, v1, v7)
                                    _ = v125
                                    _ = v126
                                }
                            }
                            _ = fi116
                        }
                        var v127 int = 2147483647
                        _ = v127
                    }
                case 42:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    var v128 uint32 = uint32(v4)
                    _ = v128
                }
                v4 = 2
            }
            var v129 uint = 2
            _ = fi19
            _ = v129
        }
    default:
        v7 = 623
    }
    switch -738 {
    case 42:
        if true {
            switch -525 {
            case 128:
                v4 -= 709
            case 0:
                v15 = v15
            }
        }
    }
    if true {
        v7 += (^(v7))
    }
    var v130 float32 = v14
    fmt.Println(v14)
    v14 = v14
    var v131 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v130
    _ = v131
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
