package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool; F1 bool; F2 int }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS8 struct { F0 int; F1 int8; F2 uint }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 string; F1 uint64; F2 bool }
type GenS11 struct { F0 int8; F1 int32; F2 bool }
type GenS12 struct { F0 int32; F1 int32; F2 int64 }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 int8; F1 uint16; F2 int32 }
type GenS15 struct { F0 int64; F1 uint32 }
type GenS16 struct { F0 int8; F1 string }
type GenS17 struct { F0 uint8; F1 int }

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 42
    v3 := []string{"hello", ("" + ""), "\n\t", "foo bar"}
    var v4 uint8 = 255
    var v5 bool = (true || ((^(v4)) <= v4))
    v1 += 1e10
    var v6 uint8 = v4
    if (!(2 != v6)) {
        var v7 int16 = 1
        v6 -= 215
        _ = v7
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 1, F1: "foo bar"}
            var v10 float32 = float32(3.14159)
            v11 := GenS1{F0: v6, F1: ""}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        if (1 != 360) {
            var v12 int = 1
            if (483 == 616) {
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                for fi13 := 0; fi13 < 1; fi13++ {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 uint8 = v6
                        _ = fi14
                        _ = v15
                    }
                    var v16 uint = 1
                    var v17 int = v12
                    _ = fi13
                    _ = v16
                    _ = v17
                }
            }
            switch "foo bar" {
            case "hello":
                v18 := GenS2{F0: v5, F1: (2 <= v6), F2: (v12 % (((v12 ^ (v12 ^ v12))) | 1))}
                _ = v18
            case "\n\t":
                var v19 uint = 1
                _ = v19
            case "":
                v2 = (v2 - v2)
            default:
                var v20 uint32 = 4294967295
                _ = v20
            }
            _ = v12
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        if v5 {
            v22 := GenS1{F0: (v4 / ((v4) | 1)), F1: ("foo bar" + "")}
            v4 = v6
            var v23 uint16 = (v2 ^ 1)
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            v1 = v1
            switch ("a" + "a") {
            case "\n\t":
                var v25 float64 = float64(-2.25)
                _ = v25
            case "a":
                fmt.Println(v2, fi24, v4)
            }
            var v26 int16 = 32767
            _ = fi24
            _ = v26
        }
        _ = fi21
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        if false {
            fmt.Println(v2, v5, v6)
        }
        _ = fi27
    }
    v28 := GenS2{F0: v5, F1: (!false), F2: 665}
    v29 := GenS1{F0: 0, F1: "a"}
    for fi30 := 0; fi30 < 18; fi30++ {
        fmt.Println(v2, v2, v5)
        var v31 float32 = 1.5
        _ = len(v3)
        _ = fi30
        _ = v31
    }
    if (1 < 971) {
        v5 = v5
    }
    if v5 {
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 float32 = v1
            for fi34 := 0; fi34 < 17; fi34++ {
                switch (fi32 % ((-1) | 1)) {
                case 2147483647:
                    if false {
                        _ = len(v3)
                    } else {
                        var v35 int64 = 1
                        if (-872 < 2147483647) {
                            var v36 uint = 4294967295
                            if (-68 > -113) {
                                if v5 {
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v6 = v6
                                    v5 = v5
                                } else {
                                    switch -480 {
                                    case 371:
                                        v37 := GenS2{F0: (!(v5 && v5)), F1: (((2 >= 411) || false) && (v5 && v5)), F2: (fi32 * 9223372036854775807)}
                                        _ = v37
                                    case 2147483647:
                                        v3 = append(v3, ((((("hello" + "foo bar") + "hello") + ("hello" + "hello")) + "\n\t") + "foo bar"))
                                    default:
                                        if v5 {
                                            switch -542 {
                                            case 32767:
                                                _ = len(v3)
                                            case 832:
                                                var v38 int16 = 42
                                                _ = v38
                                            case 536:
                                                if ((255 <= (v4 | 127)) && (true || v5)) {
                                                    fmt.Println(v35)
                                                } else {
                                                    fmt.Println(v4, v1, fi34)
                                                    v1 = float32(float64(0.1))
                                                }
                                            default:
                                                switch "hello" {
                                                case "\n\t":
                                                    v4 += 255
                                                case "foo bar":
                                                    switch (fi32 * fi32) {
                                                    case 915:
                                                        if v5 {
                                                            v39 := GenS1{F0: (^(v6)), F1: ("" + "\n\t")}
                                                            var v40 uint32 = 480
                                                            v3 = append(v3, "hello")
                                                            fmt.Println(v6, fi32, v36)
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    case 4294967295:
                                                        if true {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v41 uint32 = 4294967295
                                                                _ = v41
                                                            case "":
                                                                var v42 bool = (164 > 65535)
                                                                _ = v42
                                                            case "a":
                                                                if v5 {
                                                                    v35 += v35
                                                                    v4 = (v6 & (^(v4)))
                                                                    switch 478 {
                                                                    case 2147483647:
                                                                        v5 = (v5 || true)
                                                                    case -9223372036854775808:
                                                                        var v43 uint16 = v2
                                                                        _ = v43
                                                                    case 42:
                                                                        if false {
                                                                            v1 += (float32(-2.25) - (0.0 / float32((-2.25 / 1.5))))
                                                                            v44 := GenS2{F0: v5, F1: (-9223372036854775808 == 32768), F2: (fi32 % ((360) | 1))}
                                                                            _ = v44
                                                                        } else {
                                                                            if (v1 != v33) {
                                                                                v35 += (v35 | -746)
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v45 := GenS2{F0: false, F1: (v5 && (v5 && false)), F2: fi34}
                                                                                    _ = v45
                                                                                case "":
                                                                                    if (2147483648 < 0) {
                                                                                        fmt.Println(v6, v2)
                                                                                        v46 := GenS1{F0: v4, F1: "hello"}
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        _ = v46
                                                                                    } else {
                                                                                        v47 := GenS1{F0: v6, F1: "a"}
                                                                                        _ = v47
                                                                                    }
                                                                                case "foo bar":
                                                                                    switch -1 {
                                                                                    case 127:
                                                                                        if (42 == -678) {
                                                                                            v4 = (v6 % ((0) | 1))
                                                                                        }
                                                                                    case 128:
                                                                                        var v48 uint16 = 65535
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            if ((((32767 == 0) && v5) || (!true)) && ((v35 * (v35 * (v35 | (v35 - (v35 % (((^(v35))) | 1)))))) <= (v35 ^ v35))) {
                                                                                                switch fi32 {
                                                                                                case 42:
                                                                                                    v48 += v48
                                                                                                }
                                                                                                var v49 int = (fi34 << ((4294967295) & 63))
                                                                                                fmt.Println(v2, v1)
                                                                                                v50 := GenS2{F0: v5, F1: (-2.25 == 1.5), F2: 32768}
                                                                                                _ = v49
                                                                                                _ = v50
                                                                                            }
                                                                                        case "hello":
                                                                                            var v51 int8 = -128
                                                                                            _ = v51
                                                                                        case "foo bar":
                                                                                            var v52 float64 = (1e10 + float64((-2.25 - 0.1)))
                                                                                            _ = v52
                                                                                        }
                                                                                        _ = v48
                                                                                    case 0:
                                                                                        v53 := GenS1{F0: v6, F1: ("foo bar" + ("hello" + ((("\n\t" + "hello") + "hello") + ("hello" + "a"))))}
                                                                                        _ = v53
                                                                                    default:
                                                                                        switch fi34 {
                                                                                        case 9223372036854775807:
                                                                                            v5 = ((false || (fi32 <= int(v2))) && (v5 && false))
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v54 int16 = -977
                                                                                if (310 <= 234) {
                                                                                    v55 := GenS3{F0: (v2 << ((1) & 63)), F1: 42}
                                                                                    v4 = (^(v6))
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "hello"
                                                                                    }
                                                                                    if v5 {
                                                                                        v56 := GenS1{F0: v4, F1: "\n\t"}
                                                                                        fmt.Println(fi34)
                                                                                        _ = v56
                                                                                    }
                                                                                    _ = v55
                                                                                }
                                                                                _ = v54
                                                                            }
                                                                            fmt.Println(v5, v4, fi34)
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v36, v36, v2)
                                                            }
                                                        }
                                                    case -364:
                                                        if v5 {
                                                            var v57 float64 = (float64(((v1 / 3.14159) + 0.1)) * ((-2.25 + float64(3.14159)) + 3.14159))
                                                            if true {
                                                                var v58 int32 = int32(v35)
                                                                v59 := GenS3{F0: v2, F1: uint64(v2)}
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            _ = v57
                                                        }
                                                    default:
                                                        if (1.5 != ((1e10 * 0.0) + ((-2.25 * 0.0) + (0.0 * 3.14159)))) {
                                                            if (v5 && false) {
                                                                var v60 float32 = ((v33 / v1) / v33)
                                                                if (0 < (v4 ^ (v4 >> ((413) & 63)))) {
                                                                    fmt.Println(v36)
                                                                } else {
                                                                    v61 := GenS4{F0: -970}
                                                                    fmt.Println(v36, fi34, v36)
                                                                    var v62 bool = (v6 > (v4 * v6))
                                                                    switch fi34 {
                                                                    case 939:
                                                                        v63 := GenS4{F0: 0}
                                                                        _ = v63
                                                                    case -893:
                                                                        var v64 int64 = (v35 << ((uint(v2)) & 63))
                                                                        _ = v64
                                                                    case 42:
                                                                        v62 = false
                                                                    }
                                                                    v6 -= (v6 ^ v4)
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                v3 = append(v3, "foo bar")
                                                                _ = v60
                                                            }
                                                            _ = len(v3)
                                                            var v65 uint32 = 0
                                                            _ = v65
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            switch fi32 {
                                                            case 9223372036854775807:
                                                                v66 := GenS5{F0: 4294967295, F1: 419, F2: 88}
                                                                _ = v66
                                                            case 515:
                                                                if v5 {
                                                                    v67 := GenS3{F0: v2, F1: 18446744073709551615}
                                                                    var v68 uint16 = v2
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    fmt.Println(v6, v2, v4)
                                                                    v69 := GenS6{F0: v5}
                                                                    switch (^(fi34)) {
                                                                    case -163:
                                                                        v35 -= -9223372036854775808
                                                                    case 9223372036854775807:
                                                                        var v70 uint8 = (v6 / ((v4) | 1))
                                                                        _ = v70
                                                                    }
                                                                    _ = v69
                                                                }
                                                            case -9223372036854775808:
                                                                var v71 string = ""
                                                                fmt.Println(v71)
                                                                if false {
                                                                    v33 -= v1
                                                                    v72 := GenS7{F0: 620, F1: 30, F2: (v4 >> ((255) & 63))}
                                                                    var v73 int8 = 42
                                                                    _ = v72
                                                                    _ = v73
                                                                } else {
                                                                    var v74 int8 = 0
                                                                    v75 := GenS8{F0: fi32, F1: 100, F2: uint(v2)}
                                                                    v1 += v33
                                                                    _ = len(v3)
                                                                    _ = len(v3)
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                if (!false) {
                                                                    v4 += (v6 & v4)
                                                                    v35 += int64(fi34)
                                                                    var v76 uint64 = 42
                                                                    v3 = append(v3, ("a" + (((v71 + (v71 + "\n\t")) + v71) + ("a" + v71))))
                                                                    if v5 {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (((v71 + ("foo bar" + "foo bar")) + (v71 + (v71 + "\n\t"))) + v71)
                                                                        }
                                                                        var v77 float32 = (v1 / (0.1 / (-2.25 - v1)))
                                                                        if v5 {
                                                                            if v5 {
                                                                                var v78 float32 = float32(-2.25)
                                                                                if v5 {
                                                                                    if v5 {
                                                                                        v3 = append(v3, (v71 + ""))
                                                                                        v71 = "foo bar"
                                                                                        fmt.Println(v33)
                                                                                    }
                                                                                }
                                                                                var v79 uint8 = v4
                                                                                if (uint64(v2) <= (^(v76))) {
                                                                                    fmt.Println(v78)
                                                                                }
                                                                                v4 = (v79 / (((v4 & v4)) | 1))
                                                                                v76 = v76
                                                                                _ = v78
                                                                                _ = v79
                                                                            } else {
                                                                                var v80 float32 = float32(3.14159)
                                                                                v81 := GenS9{F0: (v35 % ((v35) | 1)), F1: v36}
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            v3 = append(v3, "")
                                                                        }
                                                                        _ = v77
                                                                    }
                                                                    _ = v76
                                                                } else {
                                                                    if (v5 || (((-2.25 / float64(float32(0.1))) * 0.0) > 1e10)) {
                                                                        if (!((v36 >> ((225) & 63)) > 255)) {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v71
                                                                            }
                                                                        } else {
                                                                            switch (fi32 - (^(fi32))) {
                                                                            case 42:
                                                                                fmt.Println(v1, v36, v1)
                                                                            case 32768:
                                                                                v6 -= v4
                                                                            case -1:
                                                                                v82 := GenS10{F0: "\n\t", F1: 0, F2: (v5 && true)}
                                                                                _ = v82
                                                                            }
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = ("\n\t" + v71)
                                                                                }
                                                                            default:
                                                                                v33 = 1.5
                                                                            }
                                                                            v83 := GenS9{F0: (v35 * int64(v4)), F1: (v36 & (v36 + v36))}
                                                                            _ = v83
                                                                        }
                                                                    } else {
                                                                        v84 := GenS3{F0: (^(v2)), F1: 18446744073709551615}
                                                                        v3 = append(v3, ((v71 + (v71 + v71)) + ("a" + v71)))
                                                                        var v85 float64 = 0.1
                                                                        fmt.Println(v36, v4, v85)
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                }
                                                                _ = v71
                                                            default:
                                                                if (966 < (v2 & (^(v2)))) {
                                                                    var v86 uint32 = 949
                                                                    v36 += uint(v2)
                                                                    v87 := GenS6{F0: ((uint8(v2) > v6) || (-688 == -2147483648))}
                                                                    _ = v86
                                                                    _ = v87
                                                                } else {
                                                                    v6 -= (v4 | (v4 ^ (v6 & 0)))
                                                                    fmt.Println(v6)
                                                                }
                                                            }
                                                            var v88 uint8 = (v4 << ((v36) & 63))
                                                            _ = v88
                                                        }
                                                    }
                                                case "hello":
                                                    var v89 uint8 = (v6 ^ (v6 + 127))
                                                    _ = v89
                                                default:
                                                    switch fi32 {
                                                    case 282:
                                                        if (2147483647 <= 1) {
                                                            switch (("" + "\n\t") + ("a" + ("\n\t" + "a"))) {
                                                            case "\n\t":
                                                                v36 = v36
                                                            case "hello":
                                                                _ = len(v3)
                                                            case "foo bar":
                                                                fmt.Println(v33)
                                                            }
                                                        }
                                                    case 42:
                                                        _ = len(v3)
                                                        if false {
                                                            v35 = (v35 & (v35 % ((539) | 1)))
                                                            if (!(!(2147483647 >= 398))) {
                                                                var v90 uint64 = 2
                                                                if v5 {
                                                                    if (1 < (v2 % ((uint16(v4)) | 1))) {
                                                                        if ("" < "") {
                                                                            if true {
                                                                                var v91 int64 = -9223372036854775808
                                                                                _ = v91
                                                                            } else {
                                                                                v2 = (v2 - (v2 | 990))
                                                                            }
                                                                        }
                                                                        var v92 uint16 = 42
                                                                        fmt.Println(v92)
                                                                        v93 := GenS6{F0: v5}
                                                                        _ = v92
                                                                        _ = v93
                                                                    }
                                                                    if (!(1 == -569)) {
                                                                        v90 -= v90
                                                                        v36 = (v36 - v36)
                                                                        var v94 int64 = v35
                                                                        _ = v94
                                                                    }
                                                                    v33 = v1
                                                                    if (v36 >= 1) {
                                                                        v35 += v35
                                                                    }
                                                                    v95 := GenS11{F0: 0, F1: 844, F2: true}
                                                                    _ = v95
                                                                } else {
                                                                    if (v5 || (true && (("" + ((("" + (("hello" + "foo bar") + "hello")) + "a") + "")) > "a"))) {
                                                                        v96 := GenS10{F0: "a", F1: v90, F2: false}
                                                                        _ = v96
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ""
                                                                        }
                                                                        v2 -= (v2 ^ v2)
                                                                    }
                                                                    v3 = append(v3, "a")
                                                                    if (-2147483648 >= int32(v36)) {
                                                                        if (-73 != 1) {
                                                                            fmt.Println(fi34, v90, fi34)
                                                                            if len(v3) > 0 {
                                                                                v3[0] = ""
                                                                            }
                                                                        } else {
                                                                            v97 := GenS7{F0: 20, F1: 42, F2: 255}
                                                                            v35 = (v35 * v35)
                                                                            var v98 int64 = (-(v35))
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                        v90 = v90
                                                                        fmt.Println(v1, v1, v5)
                                                                        v2 = v2
                                                                        v5 = false
                                                                        fmt.Println(v1, v90)
                                                                    } else {
                                                                        var v99 int32 = 1
                                                                        v100 := GenS3{F0: (v2 ^ v2), F1: 42}
                                                                        if v5 {
                                                                            v36 += v36
                                                                            v101 := GenS12{F0: 42, F1: 1, F2: (v35 << ((v36) & 63))}
                                                                            _ = v101
                                                                        }
                                                                        _ = v100
                                                                        _ = v99
                                                                    }
                                                                }
                                                                _ = len(v3)
                                                                _ = v90
                                                            }
                                                            v35 -= (v35 + int64(v2))
                                                            var v102 int32 = -2147483648
                                                            _ = v102
                                                        } else {
                                                            var v103 string = (((("" + ("" + "a")) + "a") + ("a" + "a")) + (("" + "foo bar") + (("\n\t" + ("" + "")) + ("a" + (("a" + ("hello" + "foo bar")) + "")))))
                                                            fmt.Println(v6)
                                                            v2 += 2
                                                            v104 := GenS8{F0: (fi32 >> (((v36 + (v36 & 432))) & 63)), F1: 1, F2: v36}
                                                            var v105 uint32 = 2
                                                            v105 = v105
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v33 += 1e10
                                                            v105 = v105
                                                            _ = v103
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                    }
                                                }
                                            }
                                            var v106 float32 = float32(3.14159)
                                            _ = v106
                                        } else {
                                            v6 -= (v6 << ((1) & 63))
                                            if (("a" + ("a" + "\n\t")) >= "") {
                                                if (((v35 / ((v35) | 1)) >= (v35 << ((v36) & 63))) || v5) {
                                                    v2 = v2
                                                }
                                                var v107 bool = ((!v5) || (!v5))
                                                _ = v107
                                            } else {
                                                var v108 int16 = 482
                                                switch "a" {
                                                case "hello":
                                                    var v109 bool = v5
                                                    _ = v109
                                                case "a":
                                                    v33 = (1e10 * v1)
                                                default:
                                                    v110 := GenS13{F0: 4294967295, F1: ""}
                                                    _ = v110
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    var v111 uint64 = 18446744073709551615
                                                    _ = v111
                                                case "hello":
                                                    switch "" {
                                                    case "":
                                                        var v112 uint64 = 0
                                                        _ = v112
                                                    case "foo bar":
                                                        v113 := GenS7{F0: 277, F1: int16(v4), F2: 128}
                                                        _ = v113
                                                    case "\n\t":
                                                        v108 = (v108 ^ (-(v108)))
                                                    }
                                                case "\n\t":
                                                    v33 -= ((v1 / v33) * v33)
                                                    switch ("foo bar" + "") {
                                                    case "a":
                                                        v36 = 65535
                                                    case "foo bar":
                                                        if ((417 >= 42) && false) {
                                                            v33 = v33
                                                        } else {
                                                            if v5 {
                                                                switch fi32 {
                                                                case 32768:
                                                                    v2 = (v2 + v2)
                                                                case 2147483647:
                                                                    var v114 bool = v5
                                                                    _ = v114
                                                                }
                                                                fmt.Println(v6, v108, v6)
                                                                var v115 uint32 = 4294967295
                                                                _ = v115
                                                            }
                                                            var v116 uint32 = 2
                                                            var v117 int = fi34
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                    case "hello":
                                                        if ((3.14159 - 1.5) > ((-2.25 / (float64(float32((3.14159 / 1.5))) + 1.5)) * 1e10)) {
                                                            fmt.Println(v2, v6, v2)
                                                            var v118 int = fi34
                                                            if v5 {
                                                                var v119 string = (("a" + "\n\t") + "")
                                                                var v120 string = ""
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                switch (fi32 * fi34) {
                                                                case -2147483648:
                                                                    var v121 uint16 = (v2 % ((833) | 1))
                                                                    _ = v121
                                                                default:
                                                                    switch fi32 {
                                                                    case 1:
                                                                        var v122 uint32 = 231
                                                                        _ = v122
                                                                    default:
                                                                        switch (fi34 << (((v36 << ((v36) & 63))) & 63)) {
                                                                        case -194:
                                                                            v123 := GenS14{F0: 100, F1: v2, F2: int32(v4)}
                                                                            _ = v123
                                                                        case 9223372036854775807:
                                                                            var v124 bool = v5
                                                                            fmt.Println(v2, v4, v35)
                                                                            if len(v3) > 0 {
                                                                                v3[0] = "a"
                                                                            }
                                                                            _ = v124
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(v33, v33, v1)
                                                                if v5 {
                                                                    v33 = ((v1 - ((1.5 + float32(3.14159)) + -2.25)) * float32(0.1))
                                                                    var v125 int16 = v108
                                                                    _ = v125
                                                                } else {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    _ = len(v3)
                                                                    var v126 bool = v5
                                                                    _ = len(v3)
                                                                    var v127 uint64 = 18446744073709551615
                                                                    _ = v126
                                                                    _ = v127
                                                                }
                                                                var v128 float32 = 3.14159
                                                                v119 = v120
                                                                _ = v119
                                                                _ = v120
                                                                _ = v128
                                                            } else {
                                                                if v5 {
                                                                    var v129 uint8 = v6
                                                                    _ = v129
                                                                } else {
                                                                    v130 := GenS3{F0: 42, F1: 4294967295}
                                                                    _ = v130
                                                                }
                                                                fmt.Println(fi32)
                                                            }
                                                            _ = v118
                                                        }
                                                    default:
                                                        var v131 int = 32768
                                                        _ = v131
                                                    }
                                                    v2 += (v2 % ((0) | 1))
                                                default:
                                                    var v132 float64 = 1.5
                                                    _ = v132
                                                }
                                                _ = v108
                                            }
                                        }
                                        switch fi34 {
                                        case 360:
                                            var v133 int32 = 109
                                            _ = v133
                                        case -748:
                                            if ((v2 * v2) < (v2 * (v2 ^ 316))) {
                                                fmt.Println(v36, v35)
                                                var v134 int64 = -1
                                                var v135 bool = (false && ("\n\t" > "foo bar"))
                                                _ = v134
                                                _ = v135
                                            }
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v136 := GenS13{F0: 42, F1: "a"}
                                    fmt.Println(v36, v4)
                                    _ = v136
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v35, v36, v2)
                            case "":
                                if true {
                                    var v137 uint32 = 2147483648
                                    switch (fi34 / (((fi32 ^ (fi34 % ((-725) | 1)))) | 1)) {
                                    case 127:
                                        v138 := GenS7{F0: uint32(v2), F1: 32767, F2: v6}
                                        v139 := GenS7{F0: 2147483647, F1: 0, F2: (^(v6))}
                                        if ((int64(v2) != v35) && v5) {
                                            v6 = v4
                                        }
                                        switch fi34 {
                                        case -9223372036854775808:
                                            if (v137 > v137) {
                                                switch "" {
                                                case "hello":
                                                    var v140 bool = v5
                                                    _ = v140
                                                case "a":
                                                    if (4294967295 <= 2) {
                                                        v5 = false
                                                        if (v35 != (v35 & v35)) {
                                                            var v141 int64 = v35
                                                            _ = v141
                                                        }
                                                        v36 = (v36 + 255)
                                                    } else {
                                                        if ((!v5) || true) {
                                                            var v142 int16 = 1000
                                                            v142 += v142
                                                            _ = v142
                                                        }
                                                        var v143 uint = 2
                                                        _ = v143
                                                    }
                                                }
                                                fmt.Println(fi34)
                                                switch ((((("" + "hello") + "\n\t") + "\n\t") + "") + "\n\t") {
                                                case "\n\t":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        fmt.Println(fi34, fi32, fi32)
                                                    case "":
                                                        switch "hello" {
                                                        case "":
                                                            switch "hello" {
                                                            case "a":
                                                                if len(v3) > 0 {
                                                                    v3[0] = "foo bar"
                                                                }
                                                            case "\n\t":
                                                                v33 -= -2.25
                                                            case "foo bar":
                                                                v3 = append(v3, "hello")
                                                            default:
                                                                var v144 int16 = int16(fi32)
                                                                _ = v144
                                                            }
                                                        case "hello":
                                                            fmt.Println(v6, v33)
                                                        default:
                                                            var v145 int16 = 764
                                                            v33 = v33
                                                            _ = v145
                                                        }
                                                    }
                                                }
                                            } else {
                                                v35 = -9223372036854775808
                                            }
                                        case 4294967295:
                                            v1 = 3.14159
                                        case -1:
                                            v35 = v35
                                        default:
                                            var v146 string = ("\n\t" + "\n\t")
                                            _ = v146
                                        }
                                        _ = v138
                                        _ = v139
                                    case -2147483648:
                                        var v147 float64 = (1e10 / -2.25)
                                        _ = v147
                                    case -9223372036854775808:
                                        switch (("a" + (("" + "hello") + "a")) + "hello") {
                                        case "\n\t":
                                            var v148 int32 = -2147483648
                                            _ = v148
                                        case "hello":
                                            v4 += 2
                                        }
                                    }
                                    _ = v137
                                } else {
                                    v149 := GenS10{F0: (("" + "a") + "foo bar"), F1: 42, F2: false}
                                    v2 -= 255
                                    _ = v149
                                }
                            case "a":
                                switch "foo bar" {
                                case "\n\t":
                                    switch "" {
                                    case "\n\t":
                                        if true {
                                            v150 := GenS15{F0: v35, F1: 1}
                                            _ = v150
                                        } else {
                                            v2 = (^(v2))
                                            var v151 uint8 = (v4 % ((2) | 1))
                                            _ = v151
                                        }
                                    default:
                                        if v5 {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            fmt.Println(v2, fi32)
                                            var v152 int64 = 42
                                            _ = v152
                                        } else {
                                            var v153 int = fi34
                                            var v154 uint64 = 427
                                            var v155 float32 = v1
                                            _ = v153
                                            _ = v154
                                            _ = v155
                                        }
                                    }
                                case "a":
                                    if (18446744073709551615 != 1) {
                                        var v156 int64 = 0
                                        var v157 uint8 = (v4 + v4)
                                        v4 = v157
                                        _ = v156
                                        _ = v157
                                    } else {
                                        var v158 int16 = 42
                                        var v159 uint8 = 153
                                        _ = v158
                                        _ = v159
                                    }
                                case "":
                                    fmt.Println(v5)
                                }
                            }
                            _ = v36
                        } else {
                            fmt.Println(v1)
                            v4 = 2
                        }
                        _ = v35
                    }
                case 737:
                    if (v5 && v5) {
                        var v160 uint32 = 1
                        _ = v160
                    } else {
                        var v161 string = "\n\t"
                        _ = v161
                    }
                case 128:
                    if (0 == 0) {
                        v162 := GenS16{F0: 42, F1: ("" + ("foo bar" + "hello"))}
                        switch (fi32 | (fi34 >> ((4294967295) & 63))) {
                        case -1:
                            v163 := GenS17{F0: 42, F1: 2147483647}
                            if true {
                                if (false || v5) {
                                    v6 -= v4
                                    if false {
                                        v2 -= v2
                                        var v164 int8 = -128
                                        v3 = append(v3, "hello")
                                        _ = v164
                                    }
                                } else {
                                    if (1 <= 42) {
                                        v2 = (^(v2))
                                    } else {
                                        var v165 int8 = 28
                                        _ = v165
                                    }
                                }
                                v166 := GenS1{F0: v6, F1: "a"}
                                if (0.1 != (((float64(v33) - (((0.1 * float64(v1)) / (float64((v33 - -2.25)) * (0.0 - 1e10))) / (1e10 * 0.1))) + 0.0) / float64(v33))) {
                                    if false {
                                        if (!(v5 && true)) {
                                            v4 += v4
                                        } else {
                                            v4 = v4
                                            _ = len(v3)
                                        }
                                        v167 := GenS11{F0: 127, F1: -1, F2: (v5 && v5)}
                                        _ = v167
                                    }
                                }
                                _ = v166
                            } else {
                                var v168 uint16 = (v2 | 32767)
                                _ = v168
                            }
                            _ = v163
                        case 1:
                            var v169 uint = 1
                            _ = v169
                        default:
                            var v170 int8 = int8(fi32)
                            _ = v170
                        }
                        _ = v162
                    } else {
                        v1 += (float32(((0.0 + 1.5) / 1.0)) + (-2.25 - v1))
                    }
                    switch ((("foo bar" + "a") + ((("foo bar" + "\n\t") + "\n\t") + "a")) + "hello") {
                    case "\n\t":
                        switch ("hello" + "hello") {
                        case "":
                            v1 = v33
                        case "\n\t":
                            v1 += float32(1.5)
                        }
                    }
                    var v171 int8 = 125
                    _ = v171
                }
                var v172 bool = false
                fmt.Println(fi34, fi32, v1)
                _ = fi34
                _ = v172
            }
            fmt.Println(v5, v1)
            _ = fi32
            _ = v33
        }
    } else {
        switch 32768 {
        case 2147483647:
            switch "foo bar" {
            case "hello":
                var v173 float32 = v1
                _ = v173
            default:
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            }
        default:
            fmt.Println(v4, v1, v4)
        }
    }
    var v174 int8 = 0
    v1 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v174
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
