package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint32; F2 int32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int64; F2 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8; F1 uint }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint32 = 42
    var v2 int = -941
    var v3 uint16 = 42
    v4 := make(map[uint32]bool)
    var v5 int64 = 42
    v2 += 32767
    var v6 int64 = v5
    for fi7 := 0; fi7 < 9; fi7++ {
        switch "\n\t" {
        case "hello":
            v8 := []int{(fi7 >> ((560) & 63))}
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: 128, F1: 42, F2: 42}
            _ = v9
        default:
            if ((v6 * -1) < v5) {
                var v10 int = v2
                v11 := GenS1{F0: 42, F1: -1, F2: 13}
                if (255 != 703) {
                    _ = len(v4)
                    var v12 uint16 = v3
                    _ = v12
                } else {
                    if false {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            v14 := GenS1{F0: 127, F1: -966, F2: -1}
                            _ = fi13
                            _ = v14
                        }
                        v3 = 255
                        for fi15 := 0; fi15 < 16; fi15++ {
                            fmt.Println(fi7, fi7, v2)
                            v16 := GenS1{F0: 0, F1: -1, F2: 100}
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v17 := GenS2{F0: false, F1: 1, F2: v3}
                        fmt.Println(fi7)
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 10; fi18++ {
                        var v19 float32 = 1.5
                        _ = fi18
                        _ = v19
                    }
                }
                _ = v10
                _ = v11
            } else {
                v20 := []int{(v2 << ((2) & 63)), 127, fi7, (v2 << ((uint(v1)) & 63))}
                var v21 string = ("a" + "hello")
                v6 -= v6
                _ = v20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 13; fi22++ {
            if false {
                v23 := GenS2{F0: true, F1: 58, F2: (v3 / ((42) | 1))}
                if false {
                    switch "hello" {
                    case "\n\t":
                        v24 := GenS2{F0: true, F1: -1, F2: 0}
                        _ = v24
                    case "":
                        var v25 bool = true
                        _ = v25
                    case "hello":
                        if false {
                            v26 := GenS1{F0: 127, F1: 2147483647, F2: 42}
                            if (-143 == 32767) {
                                v27 := GenS1{F0: 155, F1: 50, F2: int8(v2)}
                                _ = v27
                            } else {
                                switch "\n\t" {
                                case "":
                                    fmt.Println(v6, v3)
                                case "foo bar":
                                    switch fi7 {
                                    case -826:
                                        switch ("\n\t" + "foo bar") {
                                        case "\n\t":
                                            if true {
                                                v6 = (v6 * 1)
                                                var v28 int8 = int8(v5)
                                                v29 := GenS1{F0: 255, F1: -1, F2: (v28 / ((v28) | 1))}
                                                _ = v28
                                                _ = v29
                                            }
                                        case "a":
                                            v2 -= (-(fi7))
                                            var v30 uint64 = 18446744073709551615
                                            _ = v30
                                        case "hello":
                                            if (318 <= 42) {
                                                v1 += (v1 >> ((uint(v3)) & 63))
                                            }
                                        }
                                    default:
                                        if ((((v2 ^ (v2 ^ -2147483648)) > (fi22 | int(v1))) && (-1 == 404)) && true) {
                                            var v31 uint64 = 908
                                            switch (("a" + "hello") + ("hello" + "hello")) {
                                            case "\n\t":
                                                if false {
                                                    if false {
                                                        if true {
                                                            v32 := GenS2{F0: true, F1: 42, F2: 2}
                                                            _ = v32
                                                        } else {
                                                            if (!false) {
                                                                v33 := GenS1{F0: 128, F1: -294, F2: 55}
                                                                fmt.Println(fi7)
                                                                _ = v33
                                                            } else {
                                                                switch fi7 {
                                                                case -9223372036854775808:
                                                                    var v34 int64 = (v6 << ((2) & 63))
                                                                    _ = v34
                                                                }
                                                            }
                                                            v35 := []int{fi7, v2, fi7}
                                                            _ = v35
                                                        }
                                                    } else {
                                                        var v36 string = (((("foo bar" + "") + ("\n\t" + (("hello" + "") + ""))) + "foo bar") + ("hello" + ("foo bar" + "\n\t")))
                                                        v37 := v4[2147483647]
                                                        fmt.Println(v1)
                                                        v5 -= int64(v3)
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    fmt.Println(v3, v2, v31)
                                                } else {
                                                    if false {
                                                        var v38 uint16 = (v3 >> ((uint(v31)) & 63))
                                                        switch ("a" + ("foo bar" + (("hello" + "foo bar") + ""))) {
                                                        case "hello":
                                                            v39 := GenS3{F0: uint(v3), F1: 2147483647, F2: int32(v3)}
                                                            _ = v39
                                                        case "\n\t":
                                                            v3 -= (v3 ^ v38)
                                                        default:
                                                            v38 = (^(v38))
                                                        }
                                                        v5 = (^(v6))
                                                        _ = v38
                                                    } else {
                                                        switch "a" {
                                                        case "a":
                                                            v40 := GenS2{F0: (!true), F1: 39, F2: 65535}
                                                            _ = v40
                                                        case "\n\t":
                                                            v41 := []int{fi7, (fi22 << ((4294967295) & 63))}
                                                            if (!(1000 < 596)) {
                                                                if true {
                                                                    if false {
                                                                        v31 = 1
                                                                    } else {
                                                                        var v42 float32 = (-2.25 + 0.1)
                                                                        var v43 int16 = 32767
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                    if true {
                                                                        var v44 int64 = v5
                                                                        v41 = append(v41, -9223372036854775808)
                                                                        v41 = append(v41, fi7)
                                                                        _ = v44
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v1 -= (v1 << ((1) & 63))
                                                                        v5 -= 0
                                                                    } else {
                                                                        switch fi7 {
                                                                        case 0:
                                                                            v1 = 0
                                                                        case 42:
                                                                            fmt.Println(v2)
                                                                        case -1:
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v45 uint32 = v1
                                                                                _ = v45
                                                                            default:
                                                                                v6 = (v6 % ((v5) | 1))
                                                                            }
                                                                        default:
                                                                            switch (fi7 | v2) {
                                                                            case 4294967295:
                                                                                if true {
                                                                                    fmt.Println(v1, fi22)
                                                                                    var v46 int64 = (^(v5))
                                                                                    v41 = append(v41, (fi22 % (((fi7 % ((127) | 1))) | 1)))
                                                                                    _ = v46
                                                                                }
                                                                            case -1:
                                                                                switch (fi7 | (-(fi7))) {
                                                                                case 32768:
                                                                                    if len(v41) > 0 {
                                                                                        v41[0] = (fi7 | (fi22 & v2))
                                                                                    }
                                                                                case -2147483648:
                                                                                    v31 = (v31 / ((42) | 1))
                                                                                case -926:
                                                                                    switch (v2 >> ((27) & 63)) {
                                                                                    case 540:
                                                                                        v47 := v4[2147483647]
                                                                                        _ = v47
                                                                                    }
                                                                                    if (1 != 255) {
                                                                                        var v48 int16 = 32767
                                                                                        if false {
                                                                                            var v49 float64 = 1.5
                                                                                            _ = v49
                                                                                        } else {
                                                                                            var v50 uint8 = 205
                                                                                            _ = v50
                                                                                        }
                                                                                        v2 += v2
                                                                                        _ = v48
                                                                                    } else {
                                                                                        switch (fi7 & fi22) {
                                                                                        case 32767:
                                                                                            v51 := GenS3{F0: 780, F1: v1, F2: 2147483647}
                                                                                            _ = v51
                                                                                        case 42:
                                                                                            switch 32768 {
                                                                                            case -598:
                                                                                                if true {
                                                                                                    var v52 string = ("foo bar" + "hello")
                                                                                                    _ = v52
                                                                                                } else {
                                                                                                    if len(v41) >= 2 {
                                                                                                        v41 = v41[1:len(v41)]
                                                                                                    }
                                                                                                    v4[(v1 - 394)] = (!((v1 >> ((4294967295) & 63)) <= (v1 & (v1 % (((^(v1))) | 1)))))
                                                                                                }
                                                                                            case 42:
                                                                                                v53 := GenS3{F0: uint(v1), F1: (v1 & v1), F2: 1}
                                                                                                v54 := GenS1{F0: 119, F1: -2147483648, F2: -1}
                                                                                                _ = len(v4)
                                                                                                var v55 uint = 42
                                                                                                _ = v53
                                                                                                _ = v54
                                                                                                _ = v55
                                                                                            default:
                                                                                                v56 := v4[(v1 | v1)]
                                                                                                _ = v56
                                                                                            }
                                                                                        }
                                                                                        if false {
                                                                                            v57 := GenS1{F0: 80, F1: 2147483647, F2: int8(v31)}
                                                                                            var v58 int32 = 328
                                                                                            var v59 uint32 = 2147483648
                                                                                            if (170 <= 941) {
                                                                                                if true {
                                                                                                    v1 = (v59 + 2)
                                                                                                }
                                                                                                var v60 uint64 = v31
                                                                                                _ = v60
                                                                                            }
                                                                                            _ = len(v4)
                                                                                            _ = v57
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                        } else {
                                                                                            if (-1 > (v6 - (-(v6)))) {
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    v61 := GenS3{F0: 255, F1: (v1 + v1), F2: 0}
                                                                                                    _ = v61
                                                                                                case "a":
                                                                                                    if false {
                                                                                                        v62 := v4[v1]
                                                                                                        switch (fi22 + int(v3)) {
                                                                                                        case 1:
                                                                                                            var v63 uint32 = v1
                                                                                                            _ = v63
                                                                                                        default:
                                                                                                            if len(v41) > 0 {
                                                                                                                v41[0] = fi22
                                                                                                            }
                                                                                                        }
                                                                                                        v1 = v1
                                                                                                        switch (fi22 << ((uint(v1)) & 63)) {
                                                                                                        case -865:
                                                                                                            v64 := GenS3{F0: 4294967295, F1: 839, F2: 0}
                                                                                                            _ = v64
                                                                                                        case -140:
                                                                                                            var v65 int16 = 32767
                                                                                                            _ = v65
                                                                                                        }
                                                                                                        _ = v62
                                                                                                    } else {
                                                                                                        var v66 bool = false
                                                                                                        v41 = append(v41, (^(fi22)))
                                                                                                        _ = v66
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v67 := GenS3{F0: 27, F1: (v1 / ((2147483647) | 1)), F2: 1}
                                                                                                    _ = v67
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if ((-2.25 + (0.0 + -2.25)) < ((3.14159 + (float32((1e10 / 3.14159)) - 3.14159)) - 0.1)) {
                                                                                            var v68 uint64 = v31
                                                                                            v4[(v1 >> ((1) & 63))] = true
                                                                                            _ = v68
                                                                                        } else {
                                                                                            v6 = -773
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v69 int16 = int16(v1)
                                                                                _ = v69
                                                                            }
                                                                            if false {
                                                                                v70 := GenS4{F0: false, F1: (v31 & v31)}
                                                                                fmt.Println(v5, v1)
                                                                                _ = v70
                                                                            } else {
                                                                                var v71 float64 = ((0.1 + float64(float32((0.1 + 1e10)))) / 1e10)
                                                                                _ = v71
                                                                            }
                                                                            v4[uint32(v3)] = true
                                                                            v72 := GenS2{F0: true, F1: 0, F2: v3}
                                                                            _ = v72
                                                                        }
                                                                        v2 = fi22
                                                                    }
                                                                }
                                                            } else {
                                                                v1 -= v1
                                                            }
                                                            var v73 bool = (!(int64(v1) > v5))
                                                            _ = v41
                                                            _ = v73
                                                        case "":
                                                            v74 := []int{v2, (fi22 ^ fi22)}
                                                            _ = v74
                                                        }
                                                        if ("hello" > "") {
                                                            v3 += v3
                                                            if ((-1 != v2) || false) {
                                                                v75 := []int{}
                                                                var v76 bool = (42 >= 2147483647)
                                                                _ = v75
                                                                _ = v76
                                                            } else {
                                                                v6 += (-(v5))
                                                            }
                                                            v31 = v31
                                                        } else {
                                                            v77 := GenS5{F0: false, F1: v2}
                                                            _ = v77
                                                        }
                                                        if (250 <= (v3 - v3)) {
                                                            v3 = (v3 | v3)
                                                            v3 += v3
                                                            var v78 int = (^(fi7))
                                                            v79 := []int{fi22, (v2 + (^(v2))), (fi7 % ((fi22) | 1))}
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        if ((v31 >> ((548) & 63)) > (^(v31))) {
                                                            v1 = v1
                                                            v1 += 757
                                                        }
                                                        switch 1 {
                                                        case 4294967295:
                                                            v80 := GenS5{F0: (true && false), F1: v2}
                                                            _ = v80
                                                        default:
                                                            if false {
                                                                if false {
                                                                    var v81 int8 = 36
                                                                    fmt.Println(fi22)
                                                                    v1 -= 42
                                                                    _ = v81
                                                                } else {
                                                                    v3 = (v3 | v3)
                                                                    fmt.Println(v6, v2)
                                                                }
                                                                v4[v1] = (!(!true))
                                                                v31 += v31
                                                            } else {
                                                                v6 -= v5
                                                                v82 := []int{178, fi7, v2, v2}
                                                                _ = v82
                                                            }
                                                        }
                                                        fmt.Println(v31)
                                                        v3 += 255
                                                        switch fi7 {
                                                        case 468:
                                                            v4[v1] = false
                                                        case 32767:
                                                            var v83 int64 = int64(v3)
                                                            _ = v83
                                                        case -1:
                                                            var v84 float32 = 0.0
                                                            _ = v84
                                                        }
                                                        var v85 uint64 = (v31 ^ v31)
                                                        _ = v85
                                                    }
                                                    if (43 > 127) {
                                                        var v86 uint8 = 0
                                                        v87 := GenS6{F0: 42}
                                                        v2 += -9223372036854775808
                                                        _ = v86
                                                        _ = v87
                                                    } else {
                                                        fmt.Println(v2)
                                                        var v88 bool = false
                                                        fmt.Println(v31, v5, v2)
                                                        _ = v88
                                                    }
                                                    v3 = (^(v3))
                                                    delete(v4, (^(v1)))
                                                }
                                            case "":
                                                if (!((uint16(v1) == (v3 << ((124) & 63))) && false)) {
                                                    switch int(v3) {
                                                    case 42:
                                                        switch "hello" {
                                                        case "":
                                                            v89 := GenS2{F0: ((576 <= 1000) || (!true)), F1: -128, F2: 42}
                                                            _ = v89
                                                        case "a":
                                                            v90 := GenS5{F0: false, F1: v2}
                                                            _ = v90
                                                        case "hello":
                                                            if true {
                                                                v6 += (-(v5))
                                                                var v91 uint32 = v1
                                                                var v92 int8 = 42
                                                                fmt.Println(v2, fi7)
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                        }
                                                    case 0:
                                                        var v93 uint32 = (v1 - v1)
                                                        _ = v93
                                                    case 2147483647:
                                                        v94 := GenS6{F0: 116}
                                                        _ = v94
                                                    }
                                                    if (true || false) {
                                                        switch (fi7 / ((v2) | 1)) {
                                                        case 128:
                                                            switch (v2 << ((1) & 63)) {
                                                            case 127:
                                                                v2 = 127
                                                            case -240:
                                                                if true {
                                                                    switch -2147483648 {
                                                                    case 32767:
                                                                        var v95 string = "hello"
                                                                        _ = v95
                                                                    case 128:
                                                                        if true {
                                                                            var v96 int64 = -977
                                                                            _ = v96
                                                                        } else {
                                                                            var v97 uint64 = v31
                                                                            if true {
                                                                                fmt.Println(v2, v1)
                                                                                var v98 float64 = (3.14159 - 1e10)
                                                                                var v99 uint32 = (v1 / ((2147483648) | 1))
                                                                                _ = v98
                                                                                _ = v99
                                                                            } else {
                                                                                if false {
                                                                                    v2 = 32768
                                                                                    switch ("hello" + "a") {
                                                                                    case "foo bar":
                                                                                        v100 := GenS7{F0: 127, F1: (v5 - (^(v5))), F2: (v3 >> ((2) & 63))}
                                                                                        _ = v100
                                                                                    case "\n\t":
                                                                                        if (!((1e10 + 1e10) != ((0.0 - (0.1 - -2.25)) * (float64(1.5) / ((((0.0 * 1.5) / 3.14159) - 0.1) * float64((float32(float64(0.0)) + 1e10))))))) {
                                                                                            v31 -= 4294967295
                                                                                            v101 := GenS8{F0: 2}
                                                                                            _ = v101
                                                                                        } else {
                                                                                            if false {
                                                                                                var v102 float64 = 0.0
                                                                                                _ = v102
                                                                                            }
                                                                                            var v103 float64 = (1.5 - float64((1.5 - (3.14159 * -2.25))))
                                                                                            var v104 string = "foo bar"
                                                                                            v2 -= (^(v2))
                                                                                            v105 := []int{}
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                            _ = v105
                                                                                        }
                                                                                    case "hello":
                                                                                        fmt.Println(fi22)
                                                                                    }
                                                                                } else {
                                                                                    switch -26 {
                                                                                    case -235:
                                                                                        _ = len(v4)
                                                                                    case 2147483647:
                                                                                        v5 = v5
                                                                                    case 127:
                                                                                        v6 = 1
                                                                                    }
                                                                                    v106 := GenS6{F0: 127}
                                                                                    v3 = (v3 / (((v3 % ((1) | 1))) | 1))
                                                                                    _ = v106
                                                                                }
                                                                                v107 := GenS4{F0: (v1 <= (v1 | (^(v1)))), F1: v31}
                                                                                v1 -= v1
                                                                                _ = v107
                                                                            }
                                                                            _ = v97
                                                                        }
                                                                    default:
                                                                        switch 642 {
                                                                        case 0:
                                                                            if false {
                                                                                v108 := GenS9{F0: -128, F1: 2}
                                                                                v109 := GenS10{F0: ((("\n\t" + "") + "foo bar") + "hello")}
                                                                                _ = v108
                                                                                _ = v109
                                                                            }
                                                                        default:
                                                                            v4[0] = true
                                                                        }
                                                                    }
                                                                    var v110 int = fi7
                                                                    v6 -= int64(v31)
                                                                    _ = v110
                                                                } else {
                                                                    if (v31 != (^(v31))) {
                                                                        var v111 uint8 = 70
                                                                        _ = v111
                                                                    } else {
                                                                        var v112 int = (fi22 | v2)
                                                                        _ = v112
                                                                    }
                                                                }
                                                            default:
                                                                if true {
                                                                    var v113 bool = true
                                                                    v114 := []int{}
                                                                    var v115 uint32 = 4294967295
                                                                    _ = v113
                                                                    _ = v114
                                                                    _ = v115
                                                                }
                                                            }
                                                            if (!(!((false || true) || true))) {
                                                                v116 := []int{fi22, (fi22 - (^(v2)))}
                                                                v5 += (v5 << ((56) & 63))
                                                                var v117 uint32 = (^(v1))
                                                                v117 = (v117 | v1)
                                                                v5 += v6
                                                                _ = v116
                                                                _ = v117
                                                            }
                                                        case -457:
                                                            v2 = (fi22 ^ fi7)
                                                        case -9223372036854775808:
                                                            v6 -= (v6 | (v6 | v6))
                                                            if (false || (false && ("hello" < ("foo bar" + "hello")))) {
                                                                fmt.Println(v5, v1, fi7)
                                                                v5 = v5
                                                                v4[v1] = true
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                v2 += fi22
                                                            case "a":
                                                                delete(v4, v1)
                                                            default:
                                                                v118 := []int{fi7, (v2 << ((18446744073709551615) & 63)), fi7}
                                                                _ = v118
                                                            }
                                                            var v119 uint64 = v31
                                                            switch (((("a" + ("foo bar" + (("a" + "foo bar") + ("\n\t" + "a")))) + "\n\t") + ((("a" + ("hello" + ("a" + (((("a" + "foo bar") + "hello") + "hello") + "\n\t")))) + "\n\t") + "a")) + "foo bar") {
                                                            case "hello":
                                                                if ("foo bar" < "\n\t") {
                                                                    var v120 int16 = int16(v5)
                                                                    _ = v120
                                                                }
                                                            }
                                                            _ = v119
                                                        }
                                                    } else {
                                                        if (!true) {
                                                            switch (("a" + "hello") + "a") {
                                                            case "\n\t":
                                                                if false {
                                                                    v121 := GenS7{F0: (-(fi7)), F1: v5, F2: (v3 % ((259) | 1))}
                                                                    v122 := GenS8{F0: 255}
                                                                    v31 = (v31 >> ((255) & 63))
                                                                    _ = v121
                                                                    _ = v122
                                                                } else {
                                                                    var v123 uint64 = (^(v31))
                                                                    _ = v123
                                                                }
                                                            }
                                                            if false {
                                                                v3 -= 2
                                                            } else {
                                                                var v124 int64 = v5
                                                                _ = v124
                                                            }
                                                        }
                                                        switch ((("foo bar" + ("\n\t" + "foo bar")) + (("foo bar" + ("foo bar" + "a")) + "foo bar")) + ("foo bar" + "")) {
                                                        case "foo bar":
                                                            var v125 int8 = 1
                                                            _ = v125
                                                        case "\n\t":
                                                            if true {
                                                                var v126 bool = (-1 != (^(v5)))
                                                                v127 := GenS7{F0: v2, F1: v6, F2: v3}
                                                                var v128 int16 = 1000
                                                                _ = v126
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                        case "hello":
                                                            var v129 bool = true
                                                            _ = v129
                                                        default:
                                                            _ = len(v4)
                                                        }
                                                        if ((fi7 | 0) <= fi7) {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(v31, fi7, v2)
                                                            case "":
                                                                var v130 uint = 4294967295
                                                                _ = v130
                                                            default:
                                                                switch (fi7 ^ (fi22 + (v2 / ((v2) | 1)))) {
                                                                case 32768:
                                                                    switch ("a" + "\n\t") {
                                                                    case "\n\t":
                                                                        v5 = -1
                                                                    default:
                                                                        v6 = (v6 * (v6 | v6))
                                                                    }
                                                                case 127:
                                                                    var v131 uint = 905
                                                                    _ = v131
                                                                default:
                                                                    var v132 int = v2
                                                                    _ = v132
                                                                }
                                                            }
                                                            v1 -= 2147483648
                                                            var v133 int64 = -9223372036854775808
                                                            _ = v133
                                                        }
                                                        if false {
                                                            v6 = 9223372036854775807
                                                            v134 := []int{(fi22 / ((-551) | 1)), 903, (fi7 >> ((1) & 63)), (^(fi7))}
                                                            if len(v134) >= 2 {
                                                                v134 = v134[1:len(v134)]
                                                            }
                                                            _ = v134
                                                        }
                                                        v135 := []int{}
                                                        v1 += (v1 + (v1 + uint32(v31)))
                                                        v3 = (v3 >> ((501) & 63))
                                                        _ = v135
                                                    }
                                                } else {
                                                    var v136 int8 = -1
                                                    var v137 string = ("foo bar" + "foo bar")
                                                    v1 = v1
                                                    _ = v136
                                                    _ = v137
                                                }
                                            case "foo bar":
                                                v3 += 2
                                            }
                                            _ = v31
                                        } else {
                                            v3 = v3
                                            if false {
                                                v138 := GenS4{F0: true, F1: 2}
                                                _ = v138
                                            } else {
                                                _ = len(v4)
                                                v5 -= v5
                                                if false {
                                                    var v139 uint8 = 0
                                                    if false {
                                                        if (2147483647 <= 1) {
                                                            v140 := []int{}
                                                            var v141 bool = (42 <= -54)
                                                            _ = v140
                                                            _ = v141
                                                        } else {
                                                            switch 32767 {
                                                            case 128:
                                                                v142 := GenS11{F0: 42}
                                                                _ = v142
                                                            case 32768:
                                                                v143 := GenS4{F0: (245 != 0), F1: 617}
                                                                _ = v143
                                                            case 1:
                                                                v144 := []int{}
                                                                _ = v144
                                                            }
                                                            var v145 uint = 949
                                                            v146 := []int{fi7, fi22, (v2 * 1)}
                                                            _ = v145
                                                            _ = v146
                                                        }
                                                    }
                                                    _ = v139
                                                }
                                            }
                                            fmt.Println(fi22)
                                        }
                                    }
                                }
                                v5 = -1
                            }
                            _ = v26
                        }
                        fmt.Println(fi22, v6)
                    default:
                        if (65535 == (v1 & 774)) {
                            v6 = v5
                        } else {
                            v147 := GenS10{F0: ("" + "\n\t")}
                            v148 := GenS2{F0: false, F1: 1, F2: 65535}
                            var v149 uint32 = 2147483648
                            _ = v147
                            _ = v148
                            _ = v149
                        }
                    }
                } else {
                    switch v2 {
                    case 0:
                        var v150 bool = (230 < 255)
                        _ = v150
                    case 686:
                        if (((0.1 * 1.5) / -2.25) <= -2.25) {
                            v1 = 979
                            v5 -= v6
                        }
                    default:
                        delete(v4, (v1 * v1))
                    }
                    v151 := []int{32768}
                    delete(v4, 2147483648)
                    _ = v151
                }
                var v152 string = "a"
                _ = v152
                _ = v23
            }
            var v153 int64 = (v5 & v6)
            _ = fi22
            _ = v153
        }
        _ = fi7
    }
    switch "foo bar" {
    case "hello":
        var v154 int64 = v6
        _ = v154
    case "":
        var v155 uint32 = (v1 >> ((18446744073709551615) & 63))
        _ = v155
    case "a":
        for fi156 := 0; fi156 < 14; fi156++ {
            var v157 bool = false
            v3 += v3
            _ = fi156
            _ = v157
        }
    }
    fmt.Println(v2, v5)
    v158 := v4[(v1 >> ((1) & 63))]
    if (1 < -1) {
        v2 = 127
        if (!((!(true || v158)) && (3.14159 == ((((-2.25 + (0.1 + 1.5)) / 1e10) + 3.14159) * -2.25)))) {
            for fi159 := 0; fi159 < 16; fi159++ {
                v5 = (-(v5))
                _ = fi159
            }
        } else {
            var v160 int16 = -32768
            _ = v160
        }
        if (!(!v158)) {
            _ = len(v4)
            switch (("" + (("" + "a") + ("" + ("" + ("" + "\n\t"))))) + (("foo bar" + ("hello" + ("hello" + (("\n\t" + "hello") + "\n\t")))) + "hello")) {
            case "\n\t":
                switch v2 {
                case -698:
                    if v158 {
                        var v161 int64 = (v6 >> ((2) & 63))
                        v5 = (v5 ^ (v5 ^ (-(v161))))
                        v162 := GenS10{F0: (("\n\t" + ("a" + "")) + (("a" + (("" + "foo bar") + ("\n\t" + "hello"))) + "hello"))}
                        _ = v161
                        _ = v162
                    }
                }
            case "":
                var v163 int = v2
                _ = v163
            }
            delete(v4, v1)
        }
    } else {
        v164 := GenS11{F0: 0}
        _ = v164
    }
    if v158 {
        var v165 uint32 = 65535
        _ = v165
    }
    switch ("foo bar" + "hello") {
    case "":
        delete(v4, (^(v1)))
    default:
        if ((!(45 != -1)) && v158) {
            v2 -= v2
            switch "\n\t" {
            case "\n\t":
                switch 128 {
                case -9223372036854775808:
                    v166 := GenS1{F0: 127, F1: -716, F2: -1}
                    _ = v166
                case 127:
                    var v167 uint = 793
                    _ = v167
                }
            default:
                fmt.Println(v2)
            }
            var v168 uint16 = (v3 % ((255) | 1))
            _ = v168
        }
    }
    v1 -= v1
    v169 := GenS12{F0: 18446744073709551615, F1: v6}
    var v170 int64 = v6
    v171 := GenS2{F0: false, F1: -1, F2: (v3 * v3)}
    var v172 float32 = (0.0 - (1e10 + float32(-2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("v158:", v158)
    fmt.Println("v170:", v170)
    fmt.Println("v172:", v172)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v158
    _ = v169
    _ = v170
    _ = v171
    _ = v172
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
