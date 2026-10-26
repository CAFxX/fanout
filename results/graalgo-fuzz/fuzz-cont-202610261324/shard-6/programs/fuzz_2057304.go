package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 uint32; F2 int16 }
type GenS6 struct { F0 int32; F1 bool; F2 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 uint = 486
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 4294967295
    var v6 float32 = (-2.25 / 1e10)
    v3 -= 828
    fmt.Println(v6, v3)
    switch ("" + "\n\t") {
    case "\n\t":
        delete(v2, ("a" + ("foo bar" + "a")))
    default:
        v7 := []int{32767, 32767, 0}
        _ = v7
    }
    if true {
        switch (("\n\t" + "a") + "\n\t") {
        case "":
            var v8 int64 = (-(v4))
            _ = v8
        default:
            delete(v1, "hello")
        }
    }
    if false {
        switch -9223372036854775808 {
        case -743:
            var v9 float32 = v6
            _ = v9
        case 212:
            var v10 int = -1
            _ = v10
        }
        if false {
            if ((!false) && (!(!true))) {
                v4 -= v4
            } else {
                if ((!false) && (!(!true))) {
                    var v11 bool = true
                    v12 := []int{4294967295, 0, 112}
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(v3)
                }
            }
        }
        fmt.Println(v6, v6)
    } else {
        if (829 >= 444) {
            for fi13 := 0; fi13 < 7; fi13++ {
                v1["a"] = (("a" + "a") + ("\n\t" + ("a" + ("\n\t" + ("hello" + ("foo bar" + "\n\t"))))))
                _ = fi13
            }
        } else {
            var v14 uint8 = 172
            _ = v14
        }
        if true {
            v2["hello"] = 9223372036854775807
            var v15 uint = (v3 + (v3 - v3))
            switch 1 {
            case -179:
                switch ("" + ("a" + ("foo bar" + ""))) {
                case "":
                    fmt.Println(v4, v6, v15)
                default:
                    var v16 float64 = 0.0
                    _ = v16
                }
            default:
                if (1 == 978) {
                    v6 = v6
                }
            }
            _ = v15
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 string = ("\n\t" + ("hello" + ("\n\t" + ("" + ("foo bar" + "hello")))))
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 uint16 = 0
                var v21 int32 = 0
                v2["\n\t"] = fi17
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi17
            _ = v18
        }
    }
    v4 = v4
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS1{F0: 65535, F1: int8(v3), F2: -1}
        switch "a" {
        case "":
            if true {
                if true {
                    v6 = (v6 + float32(0.0))
                    v24 := GenS1{F0: 543, F1: -60, F2: 24}
                    _ = v24
                } else {
                    switch fi22 {
                    case 9223372036854775807:
                        _ = len(v1)
                    case 42:
                        var v25 bool = false
                        _ = v25
                    case 4294967295:
                        var v26 uint = v3
                        _ = v26
                    default:
                        v27 := GenS1{F0: 1, F1: 0, F2: 59}
                        _ = v27
                    }
                    var v28 uint64 = (v5 ^ 0)
                    _ = v28
                }
            } else {
                for fi29 := 0; fi29 < 11; fi29++ {
                    switch fi29 {
                    case 125:
                        v30 := GenS2{F0: -1, F1: 203}
                        _ = v30
                    case 127:
                        v6 = 3.14159
                    case -369:
                        v1["a"] = ("\n\t" + "\n\t")
                    }
                    v4 = (v4 << (((v3 & 42)) & 63))
                    v31 := GenS3{F0: "\n\t"}
                    _ = fi29
                    _ = v31
                }
            }
        case "foo bar":
            v32 := v2["foo bar"]
            _ = v32
        default:
            v5 = 42
        }
        _ = fi22
        _ = v23
    }
    var v33 int32 = 2147483647
    switch 9223372036854775807 {
    case 2147483647:
        var v34 uint32 = 153
        _ = v34
    case -583:
        delete(v1, ("" + ""))
    case 127:
        if true {
            v35 := GenS1{F0: 32767, F1: -128, F2: 100}
            if false {
                v33 = (v33 << ((0) & 63))
                v33 = v33
                switch -9223372036854775808 {
                case 127:
                    if (84 >= 194) {
                        v36 := []int{-191, 32768, 32768}
                        v33 -= (-(v33))
                        switch "\n\t" {
                        case "\n\t":
                            if false {
                                for fi37 := 0; fi37 < 6; fi37++ {
                                    var v38 uint = (v3 + (v3 + v3))
                                    v39 := GenS3{F0: "\n\t"}
                                    if true {
                                        for fi40 := 0; fi40 < 12; fi40++ {
                                            v33 = v33
                                            _ = fi40
                                        }
                                    } else {
                                        fmt.Println(fi37, v4)
                                        for fi41 := 0; fi41 < 13; fi41++ {
                                            switch (fi41 % ((fi41) | 1)) {
                                            case 9223372036854775807:
                                                fmt.Println(v3)
                                            default:
                                                if true {
                                                    if false {
                                                        v33 = v33
                                                        v42 := GenS4{F0: true}
                                                        v6 = v6
                                                        _ = v42
                                                    }
                                                    var v43 bool = (false && ("" > (((("a" + ("hello" + "foo bar")) + "") + ((("\n\t" + "\n\t") + "foo bar") + (("a" + "a") + ""))) + "a")))
                                                    _ = v43
                                                }
                                            }
                                            var v44 string = "\n\t"
                                            var v45 int32 = 550
                                            _ = fi41
                                            _ = v44
                                            _ = v45
                                        }
                                        v6 += (float32(-2.25) / v6)
                                    }
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            } else {
                                v6 -= (float32(float64(0.0)) * (v6 - (v6 / ((v6 + v6) * v6))))
                                v2[""] = 42
                            }
                        case "foo bar":
                            if len(v36) > 0 {
                                v36[0] = 54
                            }
                        default:
                            if (-49 < 57) {
                                v46 := GenS4{F0: true}
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    var v48 int = (fi47 ^ -137)
                                    _ = fi47
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 17; fi49++ {
                                    switch 149 {
                                    case 32767:
                                        for fi50 := 0; fi50 < 2; fi50++ {
                                            switch ("foo bar" + "\n\t") {
                                            case "a":
                                                if true {
                                                    var v51 string = ""
                                                    _ = v51
                                                } else {
                                                    v33 = (v33 >> (((v3 * v3)) & 63))
                                                    v4 = (v4 >> ((18446744073709551615) & 63))
                                                    var v52 int32 = v33
                                                    _ = v52
                                                }
                                            case "\n\t":
                                                fmt.Println(v6)
                                            }
                                            if (fi50 <= (fi50 / (((fi50 >> ((18446744073709551615) & 63))) | 1))) {
                                                v53 := GenS2{F0: (v33 + -361), F1: uint8(v5)}
                                                var v54 int32 = (v33 & (^(v33)))
                                                _ = v53
                                                _ = v54
                                            }
                                            if ((!(((!false) && false) || ((!false) || false))) && (!true)) {
                                                if (true || false) {
                                                    if (("foo bar" + "hello") != ("foo bar" + "a")) {
                                                        var v55 bool = (false || true)
                                                        if false {
                                                            var v56 int32 = (v33 / ((v33) | 1))
                                                            v57 := GenS2{F0: (v33 >> (((v3 + v3)) & 63)), F1: 226}
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        var v58 uint = v3
                                                        _ = v55
                                                        _ = v58
                                                    }
                                                    v6 += ((v6 / -2.25) - ((v6 + v6) - float32(1.5)))
                                                    fmt.Println(fi50, v6)
                                                    if true {
                                                        v33 += -860
                                                    } else {
                                                        switch -9223372036854775808 {
                                                        case 2147483647:
                                                            if len(v36) >= 2 {
                                                                v36 = v36[1:len(v36)]
                                                            }
                                                        default:
                                                            var v59 uint16 = 0
                                                            _ = v59
                                                        }
                                                        var v60 int64 = (v4 >> (((^(v3))) & 63))
                                                        v5 = (v5 + v5)
                                                        _ = v60
                                                    }
                                                    _ = len(v2)
                                                } else {
                                                    if len(v36) >= 2 {
                                                        v36 = v36[1:len(v36)]
                                                    }
                                                    if false {
                                                        if ((42 >= 0) && false) {
                                                            switch -1 {
                                                            case 0:
                                                                if false {
                                                                    v61 := GenS5{F0: false, F1: 4294967295, F2: -853}
                                                                    _ = v61
                                                                }
                                                                fmt.Println(v4, fi50, v4)
                                                            case -7:
                                                                if ((!((v6 <= 1e10) || (((3.14159 / (-2.25 / (-2.25 * v6))) * v6) >= 3.14159))) || true) {
                                                                    if ((4294967295 < 501) && (1 > -1)) {
                                                                        if false {
                                                                            _ = len(v36)
                                                                            v3 = (v3 / ((v3) | 1))
                                                                            fmt.Println(v3, fi49)
                                                                        }
                                                                        switch int(v5) {
                                                                        case 9223372036854775807:
                                                                            v62 := GenS6{F0: v33, F1: (float64((v6 - v6)) == 0.1), F2: 65535}
                                                                            _ = v62
                                                                        case 521:
                                                                            v6 = ((0.0 - (v6 / 1.0)) * float32(-2.25))
                                                                        case 553:
                                                                            if false {
                                                                                var v63 int = (fi50 * 605)
                                                                                if ((v6 >= float32(0.0)) || (true || true)) {
                                                                                    if (true || (0.0 != 1.5)) {
                                                                                        fmt.Println(v5, v4)
                                                                                        v64 := GenS7{F0: (v33 % (((v33 % ((382) | 1))) | 1))}
                                                                                        var v65 bool = (((("" + (((("foo bar" + ("a" + "hello")) + "") + ("\n\t" + "")) + "hello")) + ("foo bar" + "")) + "foo bar") > "a")
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    } else {
                                                                                        if true {
                                                                                            _ = len(v36)
                                                                                            v63 = 0
                                                                                            v63 = 869
                                                                                        } else {
                                                                                            var v66 uint = 65535
                                                                                            fmt.Println(v66, v33)
                                                                                            _ = v66
                                                                                        }
                                                                                    }
                                                                                    v2["foo bar"] = (v63 | fi50)
                                                                                    _ = len(v2)
                                                                                } else {
                                                                                    if true {
                                                                                        v67 := v2[("hello" + ("foo bar" + "a"))]
                                                                                        if false {
                                                                                            var v68 int64 = v4
                                                                                            _ = v68
                                                                                        } else {
                                                                                            if (42 >= 65535) {
                                                                                                var v69 int64 = (v4 | v4)
                                                                                                if (false || true) {
                                                                                                    v70 := GenS3{F0: "a"}
                                                                                                    _ = v70
                                                                                                } else {
                                                                                                    if ((true && false) && true) {
                                                                                                        switch (fi49 / ((32767) | 1)) {
                                                                                                        case 1:
                                                                                                            var v71 int32 = v33
                                                                                                            _ = v71
                                                                                                        case 831:
                                                                                                            v67 = fi50
                                                                                                        case -1:
                                                                                                            v4 = (v69 | -9223372036854775808)
                                                                                                        default:
                                                                                                            var v72 int8 = -1
                                                                                                            _ = v72
                                                                                                        }
                                                                                                        var v73 uint64 = (^(v5))
                                                                                                        v33 = 42
                                                                                                        _ = v73
                                                                                                    } else {
                                                                                                        v63 -= fi50
                                                                                                        if ((fi49 <= (fi50 - fi50)) || (2 <= 120)) {
                                                                                                            if false {
                                                                                                                var v74 uint64 = v5
                                                                                                                v75 := GenS2{F0: (v33 >> (((^(v3))) & 63)), F1: 194}
                                                                                                                _ = v74
                                                                                                                _ = v75
                                                                                                            }
                                                                                                            var v76 int = int(v69)
                                                                                                            v6 -= v6
                                                                                                            _ = v76
                                                                                                        }
                                                                                                        var v77 uint8 = 1
                                                                                                        v36 = append(v36, int(v3))
                                                                                                        var v78 float32 = (-2.25 - (v6 / v6))
                                                                                                        v5 += 42
                                                                                                        _ = v77
                                                                                                        _ = v78
                                                                                                    }
                                                                                                }
                                                                                                v79 := GenS4{F0: true}
                                                                                                v33 -= (v33 + (v33 >> (((^(v3))) & 63)))
                                                                                                _ = v69
                                                                                                _ = v79
                                                                                            } else {
                                                                                                v80 := GenS1{F0: 790, F1: -128, F2: -1}
                                                                                                v81 := GenS6{F0: v33, F1: ((^(v3)) <= v3), F2: (v3 & 621)}
                                                                                                if false {
                                                                                                    if false {
                                                                                                        v82 := GenS1{F0: 1, F1: 100, F2: -128}
                                                                                                        v83 := GenS8{F0: true, F1: 33}
                                                                                                        _ = v82
                                                                                                        _ = v83
                                                                                                    } else {
                                                                                                        if len(v36) >= 2 {
                                                                                                            v36 = v36[1:len(v36)]
                                                                                                        }
                                                                                                        if true {
                                                                                                            v1["hello"] = ("foo bar" + ("\n\t" + (("\n\t" + "hello") + ("foo bar" + "a"))))
                                                                                                        }
                                                                                                        v84 := GenS5{F0: true, F1: 2147483648, F2: 1}
                                                                                                        v5 -= v5
                                                                                                        if len(v36) > 0 {
                                                                                                            v36[0] = (fi50 - v63)
                                                                                                        }
                                                                                                        _ = v84
                                                                                                    }
                                                                                                    v4 -= (v4 - v4)
                                                                                                }
                                                                                                _ = v80
                                                                                                _ = v81
                                                                                            }
                                                                                        }
                                                                                        v85 := GenS9{F0: uint32(v3)}
                                                                                        _ = v67
                                                                                        _ = v85
                                                                                    }
                                                                                    if len(v36) >= 2 {
                                                                                        v36 = v36[1:len(v36)]
                                                                                    }
                                                                                    fmt.Println(v33)
                                                                                }
                                                                                v3 -= v3
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                    }
                                                                    var v86 uint32 = 2
                                                                    _ = v86
                                                                } else {
                                                                    var v87 uint = v3
                                                                    if false {
                                                                        v5 = v5
                                                                    } else {
                                                                        switch ("" + "\n\t") {
                                                                        case "a":
                                                                            if ((1.5 <= (float64(v6) - 0.0)) && (((-1 < -1) && (true && (!(1 < 1)))) || false)) {
                                                                                v3 -= (^(v3))
                                                                            }
                                                                        case "hello":
                                                                            if (!true) {
                                                                                fmt.Println(v5, v4)
                                                                                if false {
                                                                                    v33 = v33
                                                                                } else {
                                                                                    if (true && false) {
                                                                                        if (0 <= 0) {
                                                                                            v3 -= (v3 + v87)
                                                                                            var v88 int8 = 1
                                                                                            _ = v88
                                                                                        } else {
                                                                                            fmt.Println(fi50, v4)
                                                                                            v36 = append(v36, int(v33))
                                                                                            var v89 int = 9223372036854775807
                                                                                            _ = v89
                                                                                        }
                                                                                    }
                                                                                    v33 = (v33 - v33)
                                                                                    var v90 int32 = v33
                                                                                    _ = v90
                                                                                }
                                                                                v3 -= (v3 >> (((v87 >> (((v3 | (^(v3)))) & 63))) & 63))
                                                                            }
                                                                        case "":
                                                                            v36 = append(v36, fi49)
                                                                        }
                                                                        var v91 uint32 = 0
                                                                        _ = v91
                                                                    }
                                                                    _ = v87
                                                                }
                                                            case 2147483647:
                                                                v5 -= (v5 << ((920) & 63))
                                                            default:
                                                                _ = len(v1)
                                                            }
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v92 string = "hello"
                                                                _ = v92
                                                            case "a":
                                                                v6 = v6
                                                            case "hello":
                                                                _ = len(v1)
                                                            }
                                                        }
                                                        v4 = 42
                                                    } else {
                                                        var v93 int8 = -102
                                                        if true {
                                                            v94 := GenS3{F0: "a"}
                                                            v4 -= (v4 * v4)
                                                            _ = v94
                                                        }
                                                        _ = v93
                                                    }
                                                }
                                            } else {
                                                if false {
                                                    v6 = float32(0.0)
                                                } else {
                                                    fmt.Println(v33, v4)
                                                    if (!true) {
                                                        v5 -= v5
                                                    }
                                                }
                                            }
                                            _ = fi50
                                        }
                                    case 0:
                                        var v95 uint64 = (^(v5))
                                        _ = v95
                                    default:
                                        var v96 float32 = v6
                                        _ = v96
                                    }
                                    var v97 uint32 = 2
                                    v98 := GenS5{F0: (!(184 < 166)), F1: (v97 >> ((0) & 63)), F2: 1000}
                                    _ = fi49
                                    _ = v97
                                    _ = v98
                                }
                                _ = v46
                            }
                        }
                        _ = v36
                    } else {
                        if (!false) {
                            v33 += 935
                        } else {
                            for fi99 := 0; fi99 < 4; fi99++ {
                                for fi100 := 0; fi100 < 20; fi100++ {
                                    v101 := GenS2{F0: (v33 << ((18446744073709551615) & 63)), F1: 47}
                                    _ = len(v1)
                                    _ = fi100
                                    _ = v101
                                }
                                v5 = v5
                                _ = fi99
                            }
                            var v102 int64 = 668
                            v3 = (v3 * (^(v3)))
                            _ = v102
                        }
                    }
                case -17:
                    for fi103 := 0; fi103 < 19; fi103++ {
                        v104 := GenS4{F0: true}
                        _ = fi103
                        _ = v104
                    }
                case -1:
                    v6 += (v6 + v6)
                default:
                    var v105 int32 = (v33 - (v33 + v33))
                    _ = v105
                }
            }
            _ = v35
        } else {
            var v106 float64 = 1.5
            _ = v106
        }
    default:
        fmt.Println(v4, v33, v6)
    }
    var v107 uint8 = 2
    v1["a"] = ("a" + "")
    if ((v107 == (v107 | v107)) && ("\n\t" == "")) {
        v108 := []int{-924, 448, 0, -1}
        fmt.Println(v107, v3, v3)
        v5 -= v5
        _ = v108
    } else {
        switch 9223372036854775807 {
        case 988:
            switch "a" {
            case "":
                v109 := GenS4{F0: true}
                _ = v109
            case "\n\t":
                v6 -= -2.25
            case "a":
                for fi110 := 0; fi110 < 17; fi110++ {
                    if true {
                        var v111 int32 = 2147483647
                        if true {
                            v4 -= (v4 * int64(v5))
                            v112 := GenS2{F0: (v111 & v33), F1: 126}
                            _ = v112
                        }
                        _ = v111
                    } else {
                        v113 := []int{4294967295, (-(fi110))}
                        _ = v113
                    }
                    _ = fi110
                }
            default:
                v114 := []int{-48}
                _ = v114
            }
        case 0:
            var v115 float32 = float32((1.5 - 0.0))
            _ = v115
        case -284:
            v116 := GenS2{F0: (v33 ^ 2147483647), F1: v107}
            _ = v116
        default:
            v4 = v4
        }
        switch int(v3) {
        case 904:
            for fi117 := 0; fi117 < 5; fi117++ {
                v4 = v4
                for fi118 := 0; fi118 < 19; fi118++ {
                    var v119 uint32 = 462
                    v3 = (v3 | (v3 % ((uint(v5)) | 1)))
                    _ = fi118
                    _ = v119
                }
                _ = len(v2)
                _ = fi117
            }
        default:
            switch 229 {
            case 9223372036854775807:
                v5 -= (v5 - 128)
            case 128:
                var v120 int = 594
                _ = v120
            case -25:
                v107 = (^(v107))
            default:
                var v121 uint8 = 231
                _ = v121
            }
        }
    }
    v4 = 541
    fmt.Println(v5, v33, v5)
    var v122 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("v122:", v122)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v107
    _ = v122
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
