package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 uint16; F1 uint32 }
type GenS11 struct { F0 bool; F1 uint16 }
type GenS12 struct { F0 string; F1 string; F2 int16 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int32; F1 uint; F2 uint32 }
type GenS15 struct { F0 int16; F1 int }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int64; F1 uint; F2 int8 }
type GenS18 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0() int32 {
    if (737 >= 168) {
        var v7 float32 = float32(0.1)
        v8 := GenS1{F0: 681, F1: -9223372036854775808}
        var v9 bool = false
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 int64 = -9223372036854775808
        v11 := GenS1{F0: 726, F1: (v10 % ((v10) | 1))}
        if false {
            fmt.Println(v10)
        }
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 32767, F1: 1}
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := []int{fi13, fi13, 32767}
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 2, F1: 42}
    v16 := GenS1{F0: 740, F1: 683}
    _ = v12
    _ = v15
    _ = v16
    return -2147483648
}
func genFn1(p0 int16) int {
    var v17 uint32 = 1
    var v18 int8 = 1
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v18, v18)
            _ = fi20
        }
        _ = fi19
    }
    var v21 uint = 396
    if ((!true) && false) {
        v22 := GenS1{F0: 2, F1: -512}
        var v23 float64 = 3.14159
        switch -9223372036854775808 {
        case -868:
            v24 := GenS1{F0: 513, F1: -1}
            _ = v24
        case -242:
            v25 := []int{}
            _ = v25
        default:
            v26 := GenS2{F0: 255}
            _ = v26
        }
        _ = v22
        _ = v23
    } else {
        if (2 == 998) {
            p0 -= p0
            for fi27 := 0; fi27 < 14; fi27++ {
                if false {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v17 = v17
                        v29 := []int{(^(fi28))}
                        p0 -= p0
                        _ = fi28
                        _ = v29
                    }
                } else {
                    if (false && (1 <= 18446744073709551615)) {
                        v18 = v18
                    }
                    v30 := GenS3{F0: -789, F1: false}
                    _ = v30
                }
                _ = fi27
            }
        }
        switch "a" {
        case "a":
            var v31 int64 = 964
            _ = v31
        }
        var v32 float64 = 1.5
        _ = v32
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v21
    return -632
}

func main() {
    var v1 int8 = 33
    var v2 int8 = (v1 >> ((4294967295) & 63))
    v3 := []string{(("foo bar" + "foo bar") + ""), ("a" + "")}
    var v4 float64 = (3.14159 - float64(0.1))
    v5 := make(map[int]string)
    v6 := []int{}
    var v33 float32 = ((1.5 + 0.1) * 0.0)
    switch ((("hello" + "\n\t") + ("\n\t" + "\n\t")) + (("\n\t" + "") + "a")) {
    case "a":
        switch genFn1(-32768) {
        case 9223372036854775807:
            var v34 uint16 = 418
            _ = v34
        case 746:
            switch 32767 {
            case 518:
                v4 += v4
            case 2147483647:
                v2 = v2
            case 4294967295:
                v1 = v2
            default:
                for fi35 := 0; fi35 < 10; fi35++ {
                    v5[fi35] = "hello"
                    _ = len(v3)
                    _ = fi35
                }
            }
        default:
            v4 += (v4 * 1.5)
        }
    }
    v4 = v4
    fmt.Println(v2)
    if (!(0 == 155)) {
        var v36 uint64 = 42
        _ = v36
    } else {
        for fi37 := 0; fi37 < 12; fi37++ {
            var v38 uint = 18446744073709551615
            v39 := v5[genFn1(32767)]
            _ = fi37
            _ = v38
            _ = v39
        }
        if (!(false || false)) {
            delete(v5, 839)
            v3 = append(v3, "foo bar")
            if (v4 >= -2.25) {
                switch ("hello" + "a") {
                case "\n\t":
                    v33 = ((-2.25 / v33) + (float32(1.5) - ((0.0 * ((float32(v4) / -2.25) / (float32(-2.25) - 1.5))) + (float32(-2.25) - 0.0))))
                case "foo bar":
                    switch 781 {
                    case 400:
                        var v40 float32 = 1e10
                        _ = v40
                    case -1:
                        if (!(32767 == 820)) {
                            var v41 float64 = (float64(0.0) + (3.14159 + 0.1))
                            if true {
                                v42 := GenS3{F0: genFn0(), F1: false}
                                fmt.Println(v4, v2)
                                _ = v42
                            }
                            if true {
                                v4 = ((v4 * ((-2.25 + v4) * (v41 - 0.0))) + (v41 - (v41 / v41)))
                                for fi43 := 0; fi43 < 13; fi43++ {
                                    var v44 int16 = -1
                                    _ = fi43
                                    _ = v44
                                }
                            } else {
                                var v45 uint8 = 187
                                _ = v45
                            }
                            _ = v41
                        }
                    default:
                        var v46 int16 = -32768
                        _ = v46
                    }
                default:
                    v4 += float64(0.1)
                }
                for fi47 := 0; fi47 < 2; fi47++ {
                    var v48 float64 = (0.1 - (0.1 * v4))
                    if len(v6) > 0 {
                        v6[0] = (fi47 * fi47)
                    }
                    _ = fi47
                    _ = v48
                }
                v49 := GenS2{F0: 135}
                _ = v49
            } else {
                if true {
                    var v50 int64 = 878
                    var v51 uint64 = 42
                    _ = len(v3)
                    _ = v50
                    _ = v51
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            for fi52 := 0; fi52 < 19; fi52++ {
                fmt.Println(v33, v33, v4)
                _ = fi52
            }
        }
    }
    var v53 float32 = v33
    var v54 string = ("a" + ("\n\t" + "\n\t"))
    var v55 int8 = (v2 - 1)
    for fi56 := 0; fi56 < 14; fi56++ {
        for fi57 := 0; fi57 < 15; fi57++ {
            var v58 int32 = 1
            v53 = (v53 / 1e10)
            var v59 int64 = 42
            v33 = v33
            var v60 int = (^(fi57))
            v61 := GenS2{F0: 1}
            v33 = 3.14159
            v60 = (fi56 + genFn1(-1))
            var v62 uint16 = 1
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
            _ = v61
            _ = v62
        }
        _ = fi56
    }
    switch 1 {
    case 128:
        var v63 string = v54
        _ = v63
    case 32768:
        fmt.Println(v54)
    case 1:
        v4 = float64(v53)
    }
    _ = len(v3)
    if false {
        fmt.Println(v33, v55, v53)
        for fi64 := 0; fi64 < 11; fi64++ {
            if false {
                if ((v4 != (v4 * (0.1 - v4))) || (v54 != ("" + v54))) {
                    fmt.Println(fi64, fi64, v33)
                } else {
                    v55 = -128
                    if false {
                        if (!(v33 < 1.5)) {
                            v4 = (v4 / (v4 / 1.5))
                            if false {
                                for fi65 := 0; fi65 < 19; fi65++ {
                                    switch -679 {
                                    case -972:
                                        _ = len(v6)
                                    default:
                                        v5[-80] = v54
                                    }
                                    _ = fi65
                                }
                                for fi66 := 0; fi66 < 15; fi66++ {
                                    v54 = v54
                                    _ = fi66
                                }
                                if false {
                                    switch "foo bar" {
                                    case "a":
                                        v67 := GenS2{F0: 0}
                                        _ = v67
                                    case "\n\t":
                                        _ = len(v6)
                                    }
                                    for fi68 := 0; fi68 < 6; fi68++ {
                                        if false {
                                            var v69 float32 = 1e10
                                            fmt.Println(fi68, v4, v33)
                                            v70 := GenS3{F0: genFn0(), F1: true}
                                            _ = v69
                                            _ = v70
                                        } else {
                                            var v71 uint16 = 42
                                            if (fi64 == (-(fi64))) {
                                                if len(v3) > 0 {
                                                    v3[0] = "foo bar"
                                                }
                                                if (uint8(v71) >= 42) {
                                                    if (false || false) {
                                                        if true {
                                                            if false {
                                                                if (127 != 1) {
                                                                    v1 = v2
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ("" + v54)
                                                                    }
                                                                    var v72 float32 = v53
                                                                    _ = v72
                                                                }
                                                            } else {
                                                                v53 -= 0.0
                                                                var v73 uint = 65535
                                                                _ = v73
                                                            }
                                                            v74 := GenS1{F0: (^(v71)), F1: 0}
                                                            if false {
                                                                v5[(fi68 >> ((4294967295) & 63))] = v54
                                                                v55 += (v55 & 0)
                                                            }
                                                            var v75 uint = 1
                                                            _ = v74
                                                            _ = v75
                                                        } else {
                                                            var v76 int64 = 1
                                                            _ = v76
                                                        }
                                                        var v77 int64 = -9223372036854775808
                                                        _ = v77
                                                    }
                                                    var v78 int8 = (-(v55))
                                                    switch v54 {
                                                    case "":
                                                        v5[(fi68 - (fi64 & int(v71)))] = "hello"
                                                    case "\n\t":
                                                        v55 -= (^(v2))
                                                    default:
                                                        v79 := GenS4{F0: 997, F1: 0}
                                                        _ = v79
                                                    }
                                                    _ = v78
                                                } else {
                                                    var v80 int64 = 0
                                                    switch (v54 + "a") {
                                                    case "a":
                                                        if false {
                                                            v2 -= -128
                                                            v1 = 0
                                                            v2 -= (-(v2))
                                                        }
                                                    case "foo bar":
                                                        var v81 uint32 = 60
                                                        _ = v81
                                                    case "":
                                                        switch ("" + ("a" + v54)) {
                                                        case "hello":
                                                            if (1 != 4294967295) {
                                                                if true {
                                                                    fmt.Println(v4, v80, v2)
                                                                    if (!(18446744073709551615 < 255)) {
                                                                        v82 := GenS1{F0: (v71 * (^(v71))), F1: (v80 - 42)}
                                                                        switch fi68 {
                                                                        case -703:
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        case 42:
                                                                            if false {
                                                                                var v83 uint16 = (v71 >> ((18446744073709551615) & 63))
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = "\n\t"
                                                                                }
                                                                                _ = v83
                                                                            } else {
                                                                                v33 = v33
                                                                                switch (fi68 + -9223372036854775808) {
                                                                                case -14:
                                                                                    var v84 string = ((v54 + v54) + (v54 + v54))
                                                                                    _ = v84
                                                                                case 695:
                                                                                    fmt.Println(v54)
                                                                                default:
                                                                                    v85 := GenS2{F0: 2}
                                                                                    _ = v85
                                                                                }
                                                                                if false {
                                                                                    var v86 int32 = genFn0()
                                                                                    v87 := GenS2{F0: 128}
                                                                                    if (v71 != 42) {
                                                                                        v88 := GenS5{F0: 0, F1: 2, F2: -32768}
                                                                                        fmt.Println(v2, v71)
                                                                                        _ = v88
                                                                                    }
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                } else {
                                                                                    if false {
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            var v89 uint32 = 2
                                                                                            _ = v89
                                                                                        case "hello":
                                                                                            _ = len(v5)
                                                                                        case "a":
                                                                                            v90 := GenS6{F0: -265}
                                                                                            _ = v90
                                                                                        }
                                                                                        var v91 uint16 = (v71 ^ 255)
                                                                                        if (!(v55 != v55)) {
                                                                                            if (277 >= genFn0()) {
                                                                                                v92 := GenS7{F0: -2147483648}
                                                                                                v53 = (v33 - v33)
                                                                                                _ = v92
                                                                                            }
                                                                                            var v93 string = ""
                                                                                            switch (v93 + "hello") {
                                                                                            case "foo bar":
                                                                                                if false {
                                                                                                    v94 := GenS6{F0: -46}
                                                                                                    if false {
                                                                                                        v55 = (v1 / ((int8(fi68)) | 1))
                                                                                                        v95 := GenS5{F0: 1, F1: 146, F2: -276}
                                                                                                        fmt.Println(v2)
                                                                                                        _ = v95
                                                                                                    } else {
                                                                                                        v93 = v54
                                                                                                        v2 = v55
                                                                                                        if (v2 <= (-(v55))) {
                                                                                                            if true {
                                                                                                                v96 := GenS8{F0: (^(v71))}
                                                                                                                _ = v96
                                                                                                            } else {
                                                                                                                var v97 int = fi64
                                                                                                                _ = v97
                                                                                                            }
                                                                                                            var v98 int64 = v80
                                                                                                            v99 := GenS5{F0: 1, F1: 65, F2: 530}
                                                                                                            _ = v98
                                                                                                            _ = v99
                                                                                                        } else {
                                                                                                            v100 := GenS9{F0: v54, F1: (v55 & (v55 | (-(v55))))}
                                                                                                            _ = v100
                                                                                                        }
                                                                                                    }
                                                                                                    v101 := GenS10{F0: v91, F1: 1}
                                                                                                    _ = v101
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    v93 = v54
                                                                                                    if (!(false || true)) {
                                                                                                        fmt.Println(v93)
                                                                                                        delete(v5, fi64)
                                                                                                        var v102 uint = 4294967295
                                                                                                        _ = v102
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v91 = v91
                                                                                            _ = v93
                                                                                        }
                                                                                        _ = v91
                                                                                    } else {
                                                                                        if false {
                                                                                            v4 = v4
                                                                                            fmt.Println(v33, v80, v1)
                                                                                            if (!true) {
                                                                                                fmt.Println(v71, v71)
                                                                                            }
                                                                                            delete(v5, (^(fi64)))
                                                                                            var v103 string = (v54 + v54)
                                                                                            _ = v103
                                                                                        } else {
                                                                                            v33 += 1.5
                                                                                        }
                                                                                        if (v80 >= (v80 | -429)) {
                                                                                            var v104 int32 = 42
                                                                                            _ = v104
                                                                                        } else {
                                                                                            var v105 uint32 = 600
                                                                                            v4 += v4
                                                                                            _ = v105
                                                                                        }
                                                                                    }
                                                                                    if (true || true) {
                                                                                        var v106 uint16 = 255
                                                                                        switch v54 {
                                                                                        case "foo bar":
                                                                                            v55 = -76
                                                                                        case "hello":
                                                                                            var v107 int64 = (v80 << ((0) & 63))
                                                                                            _ = v107
                                                                                        default:
                                                                                            var v108 uint16 = (v106 & v71)
                                                                                            _ = v108
                                                                                        }
                                                                                        var v109 uint64 = 18446744073709551615
                                                                                        _ = v106
                                                                                        _ = v109
                                                                                    }
                                                                                    var v110 int16 = int16(v71)
                                                                                    _ = v110
                                                                                }
                                                                            }
                                                                        case 0:
                                                                            v71 = v71
                                                                            if true {
                                                                                v80 = v80
                                                                                var v111 uint16 = (v71 & v71)
                                                                                fmt.Println(v2, fi68, v71)
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = v54
                                                                                }
                                                                                if ((706 != 990) && false) {
                                                                                    v4 -= 3.14159
                                                                                    _ = len(v6)
                                                                                    v112 := GenS7{F0: 649}
                                                                                    if true {
                                                                                        fmt.Println(v71, v71)
                                                                                        v53 += ((v33 * 0.0) + v53)
                                                                                        v33 = v53
                                                                                        var v113 bool = (v53 == float32(((3.14159 - v4) + (1e10 - float64((((((v33 - v33) * 1.5) + (1e10 - float32((-2.25 - 1e10)))) * 3.14159) + v53))))))
                                                                                        v114 := GenS8{F0: 2}
                                                                                        _ = v113
                                                                                        _ = v114
                                                                                    } else {
                                                                                        var v115 int64 = (v80 >> ((42) & 63))
                                                                                        if (-2.25 >= (v33 * 0.1)) {
                                                                                            var v116 float32 = v33
                                                                                            v116 = (v53 * ((v33 / (v53 - 1e10)) * 3.14159))
                                                                                            v117 := v5[-133]
                                                                                            _ = v116
                                                                                            _ = v117
                                                                                        }
                                                                                        if false {
                                                                                            fmt.Println(fi68)
                                                                                            fmt.Println(v115, v2)
                                                                                            v55 = 0
                                                                                        } else {
                                                                                            v118 := GenS10{F0: v111, F1: 42}
                                                                                            _ = v118
                                                                                        }
                                                                                        _ = v115
                                                                                    }
                                                                                    _ = v112
                                                                                }
                                                                                v119 := GenS11{F0: ((v4 + ((float64((0.0 - (float32(v4) / (v53 / v33)))) - v4) * v4)) >= -2.25), F1: 255}
                                                                                switch fi68 {
                                                                                case 127:
                                                                                    var v120 uint8 = 128
                                                                                    _ = v120
                                                                                case 665:
                                                                                    v121 := GenS12{F0: "foo bar", F1: v54, F2: 42}
                                                                                    _ = v121
                                                                                case 9223372036854775807:
                                                                                    v71 = (v71 << ((255) & 63))
                                                                                }
                                                                                v54 = "foo bar"
                                                                                _ = v111
                                                                                _ = v119
                                                                            }
                                                                            switch 1 {
                                                                            case -9223372036854775808:
                                                                                var v122 uint64 = 0
                                                                                _ = v122
                                                                            }
                                                                            v71 -= (v71 & 2)
                                                                            switch v54 {
                                                                            case "hello":
                                                                                var v123 uint8 = 127
                                                                                _ = v123
                                                                            case "":
                                                                                v80 = v80
                                                                                _ = len(v3)
                                                                                v2 -= 0
                                                                            }
                                                                            if (!(true && (false && false))) {
                                                                                v53 -= 1e10
                                                                                v71 = 65535
                                                                            }
                                                                            v1 -= (v2 % ((v55) | 1))
                                                                            var v124 uint16 = v71
                                                                            _ = v124
                                                                        }
                                                                        fmt.Println(v33, v55, v53)
                                                                        _ = v82
                                                                    } else {
                                                                        v2 = v2
                                                                        v33 = (3.14159 / 1.0)
                                                                        if (!true) {
                                                                            v4 = (0.1 - -2.25)
                                                                            var v125 uint8 = 83
                                                                            _ = v125
                                                                        }
                                                                    }
                                                                    v1 += v55
                                                                    fmt.Println(v54, v4)
                                                                    v55 = -128
                                                                }
                                                                var v126 int64 = int64(v1)
                                                                _ = v126
                                                            } else {
                                                                v127 := v5[(fi64 << ((42) & 63))]
                                                                if (((!false) && (190 >= 241)) || (((3.14159 <= 1e10) && (821 != 255)) && false)) {
                                                                    v128 := v5[(fi64 >> ((18446744073709551615) & 63))]
                                                                    var v129 int = 9223372036854775807
                                                                    _ = len(v5)
                                                                    _ = v128
                                                                    _ = v129
                                                                }
                                                                _ = v127
                                                            }
                                                            v80 -= v80
                                                            v33 = v53
                                                        default:
                                                            var v130 int32 = 2147483647
                                                            var v131 uint8 = 1
                                                            fmt.Println(v130)
                                                            _ = v130
                                                            _ = v131
                                                        }
                                                    default:
                                                        v55 = (v1 + v2)
                                                    }
                                                    _ = v80
                                                }
                                                fmt.Println(v33, v4)
                                            }
                                            fmt.Println(v55)
                                            _ = v71
                                        }
                                        v132 := GenS3{F0: 2147483647, F1: (!(2147483648 < 2))}
                                        _ = fi68
                                        _ = v132
                                    }
                                    for fi133 := 0; fi133 < 5; fi133++ {
                                        if false {
                                            var v134 int64 = -775
                                            v1 += v2
                                            _ = v134
                                        }
                                        _ = fi133
                                    }
                                }
                            }
                            for fi135 := 0; fi135 < 9; fi135++ {
                                var v136 bool = (459 <= 2)
                                _ = fi135
                                _ = v136
                            }
                        } else {
                            fmt.Println(v33, v53)
                            var v137 uint32 = 1
                            var v138 float32 = v33
                            _ = v137
                            _ = v138
                        }
                        for fi139 := 0; fi139 < 1; fi139++ {
                            if false {
                                var v140 uint64 = 18446744073709551615
                                var v141 int = (fi64 % ((fi64) | 1))
                                v142 := GenS6{F0: -2}
                                _ = v140
                                _ = v141
                                _ = v142
                            } else {
                                var v143 int = fi139
                                if (v54 == "hello") {
                                    switch fi139 {
                                    case -637:
                                        if true {
                                            switch 127 {
                                            case 2147483647:
                                                delete(v5, 2147483647)
                                            default:
                                                var v144 uint8 = 1
                                                _ = v144
                                            }
                                            if (false || (float32(-2.25) > ((v53 + v33) + 0.1))) {
                                                v143 += genFn1(42)
                                            }
                                            fmt.Println(v54, v2, v33)
                                            var v145 int32 = 459
                                            v145 = (v145 * 0)
                                            _ = v145
                                        }
                                    default:
                                        switch genFn1(1) {
                                        case -661:
                                            var v146 int = 127
                                            _ = v146
                                        case -213:
                                            if true {
                                                if true {
                                                    if true {
                                                        var v147 uint64 = 42
                                                        var v148 int32 = -2147483648
                                                        _ = v147
                                                        _ = v148
                                                    } else {
                                                        switch 127 {
                                                        case 2147483647:
                                                            v1 = 1
                                                        }
                                                        v33 = (float32(1.5) * v53)
                                                    }
                                                } else {
                                                    v149 := v5[genFn1(0)]
                                                    if ("hello" >= "a") {
                                                        v4 = float64(v33)
                                                        fmt.Println(v55, v4, fi64)
                                                    }
                                                    v149 = (v54 + v54)
                                                    _ = v149
                                                }
                                                var v150 uint32 = 703
                                                _ = v150
                                            } else {
                                                if (128 > 170) {
                                                    var v151 bool = ((true || false) && true)
                                                    if v151 {
                                                        _ = len(v6)
                                                        var v152 int64 = 0
                                                        _ = v152
                                                    }
                                                    fmt.Println(v33)
                                                    _ = v151
                                                }
                                                var v153 uint = 2
                                                _ = v153
                                            }
                                        case 717:
                                            v1 = -128
                                            var v154 float32 = v33
                                            _ = v154
                                        default:
                                            if false {
                                                v53 -= v33
                                                _ = len(v5)
                                                fmt.Println(fi64)
                                            }
                                        }
                                    }
                                    v4 -= v4
                                    v155 := GenS13{F0: 9223372036854775807}
                                    _ = v155
                                }
                                if (!((^(v1)) >= 0)) {
                                    var v156 bool = false
                                    if (65535 != 255) {
                                        var v157 int16 = 42
                                        v158 := GenS2{F0: 1}
                                        v2 = 1
                                        _ = v157
                                        _ = v158
                                    } else {
                                        var v159 string = v54
                                        var v160 uint32 = 708
                                        var v161 uint32 = v160
                                        if len(v6) > 0 {
                                            v6[0] = (fi139 << ((65535) & 63))
                                        }
                                        _ = v159
                                        _ = v160
                                        _ = v161
                                    }
                                    switch v54 {
                                    case "hello":
                                        var v162 int8 = v1
                                        _ = v162
                                    default:
                                        if (!(true && (0 > 331))) {
                                            var v163 int32 = genFn0()
                                            v164 := GenS14{F0: 42, F1: 42, F2: 900}
                                            _ = v163
                                            _ = v164
                                        } else {
                                            v6 = append(v6, fi139)
                                        }
                                    }
                                    _ = v156
                                }
                                _ = v143
                            }
                            fmt.Println(v1, v2)
                            v165 := GenS9{F0: "a", F1: v2}
                            _ = fi139
                            _ = v165
                        }
                        v2 += v2
                    }
                }
            } else {
                switch fi64 {
                case 9223372036854775807:
                    var v166 float64 = 0.0
                    _ = v166
                default:
                    if (genFn1(int16(fi64)) < (fi64 / (((^(fi64))) | 1))) {
                        for fi167 := 0; fi167 < 5; fi167++ {
                            v4 = 0.0
                            _ = fi167
                        }
                    }
                }
                delete(v5, 437)
            }
            v168 := GenS15{F0: 0, F1: fi64}
            _ = len(v6)
            _ = fi64
            _ = v168
        }
        if false {
            for fi169 := 0; fi169 < 6; fi169++ {
                var v170 uint16 = 945
                _ = len(v3)
                if true {
                    if (!(true || (!true))) {
                        v171 := GenS7{F0: genFn0()}
                        v172 := v5[(fi169 / ((fi169) | 1))]
                        var v173 int8 = (v2 % ((0) | 1))
                        _ = v171
                        _ = v172
                        _ = v173
                    } else {
                        var v174 int32 = -2147483648
                        _ = v174
                    }
                }
                _ = fi169
                _ = v170
            }
            switch 32768 {
            case 42:
                if ((!("hello" < v54)) || ((!(true && false)) || true)) {
                    switch -65 {
                    case 579:
                        if false {
                            v2 = v1
                        } else {
                            v1 = v55
                        }
                    case 409:
                        v54 = ("" + (v54 + ("foo bar" + "\n\t")))
                    default:
                        v55 = v2
                    }
                } else {
                    for fi175 := 0; fi175 < 15; fi175++ {
                        v33 = (v53 / 1.5)
                        _ = fi175
                    }
                    v176 := GenS16{F0: (false && false)}
                    if (!(-735 <= 0)) {
                        var v177 uint = 0
                        switch "foo bar" {
                        case "\n\t":
                            v2 += v1
                        case "a":
                            v2 += (v1 << ((255) & 63))
                        case "hello":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        var v178 uint16 = 2
                        _ = v177
                        _ = v178
                    } else {
                        var v179 uint16 = 65535
                        _ = v179
                    }
                    _ = v176
                }
            case 9223372036854775807:
                fmt.Println(v53, v53)
            case -9223372036854775808:
                if (!(!(false && (!(!(42 != 32767)))))) {
                    switch 4294967295 {
                    case 0:
                        var v180 string = "foo bar"
                        _ = v180
                    default:
                        var v181 uint8 = 42
                        _ = v181
                    }
                    _ = len(v3)
                    switch (v54 + v54) {
                    case "hello":
                        var v182 uint = 1
                        _ = v182
                    case "a":
                        for fi183 := 0; fi183 < 11; fi183++ {
                            switch "a" {
                            case "a":
                                v54 = v54
                            case "hello":
                                for fi184 := 0; fi184 < 19; fi184++ {
                                    _ = len(v5)
                                    v185 := GenS1{F0: 255, F1: 0}
                                    _ = len(v6)
                                    _ = fi184
                                    _ = v185
                                }
                            }
                            v186 := GenS16{F0: (-18 > -9223372036854775808)}
                            _ = fi183
                            _ = v186
                        }
                    case "\n\t":
                        var v187 float32 = v33
                        _ = v187
                    }
                }
            default:
                v3 = append(v3, v54)
            }
        }
    } else {
        for fi188 := 0; fi188 < 13; fi188++ {
            for fi189 := 0; fi189 < 14; fi189++ {
                v33 = float32(v4)
                switch "hello" {
                case "\n\t":
                    var v190 uint64 = 18446744073709551615
                    _ = v190
                }
                v4 -= (v4 + (1e10 + (v4 - float64(1.5))))
                _ = fi189
            }
            _ = fi188
        }
    }
    if (-9223372036854775808 >= -9223372036854775808) {
        v33 -= 0.1
        v191 := v5[-2147483648]
        _ = v191
    } else {
        if (!true) {
            var v192 float32 = -2.25
            v33 = (v33 * (float32((v4 + v4)) / v192))
            _ = v192
        } else {
            fmt.Println(v1)
            for fi193 := 0; fi193 < 15; fi193++ {
                switch (-(fi193)) {
                case 32767:
                    switch fi193 {
                    case 4294967295:
                        v1 += (-(v55))
                    case -319:
                        _ = len(v5)
                    case 9223372036854775807:
                        v2 -= (v2 & v2)
                    default:
                        for fi194 := 0; fi194 < 8; fi194++ {
                            var v195 string = "a"
                            v196 := GenS7{F0: -677}
                            if (v2 != int8(fi193)) {
                                v195 = "\n\t"
                                v33 -= (v53 / v53)
                                v195 = (v54 + "foo bar")
                            } else {
                                v197 := GenS17{F0: -27, F1: 65535, F2: (v1 << ((77) & 63))}
                                _ = v197
                            }
                            _ = fi194
                            _ = v195
                            _ = v196
                        }
                    }
                case 2147483647:
                    v53 -= (((v53 + v53) - 1e10) * v33)
                case -9223372036854775808:
                    var v198 float32 = (v53 - float32((0.1 - 0.0)))
                    _ = v198
                default:
                    v199 := GenS16{F0: false}
                    _ = v199
                }
                if false {
                    v33 += 0.0
                    v33 -= 1e10
                    switch (fi193 % (((fi193 | 42)) | 1)) {
                    case -9223372036854775808:
                        v200 := GenS10{F0: 42, F1: 42}
                        _ = v200
                    case 127:
                        for fi201 := 0; fi201 < 16; fi201++ {
                            v33 = float32(1.5)
                            _ = fi201
                        }
                    default:
                        v2 = 100
                    }
                }
                _ = fi193
            }
            for fi202 := 0; fi202 < 17; fi202++ {
                v203 := GenS18{F0: 287, F1: 42, F2: 4294967295}
                _ = fi202
                _ = v203
            }
        }
    }
    for fi204 := 0; fi204 < 6; fi204++ {
        if true {
            v2 = (v1 & v55)
            v205 := GenS11{F0: (true && true), F1: 2}
            for fi206 := 0; fi206 < 3; fi206++ {
                v1 = v55
                v2 += v2
                var v207 uint32 = 369
                _ = fi206
                _ = v207
            }
            _ = v205
        } else {
            if (!false) {
                fmt.Println(v55)
                v208 := GenS13{F0: 9223372036854775807}
                _ = v208
            } else {
                var v209 int32 = 781
                v210 := v5[42]
                fmt.Println(v55, v2)
                _ = v209
                _ = v210
            }
        }
        _ = fi204
    }
    v33 += v33
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
