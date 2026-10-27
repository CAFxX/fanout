package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int32; F1 bool; F2 string }
type GenS17 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS18 struct { F0 int8; F1 uint8 }
type GenS19 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS20 struct { F0 int64; F1 int16 }
type GenS21 struct { F0 uint; F1 uint64; F2 int16 }
type GenS22 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint64, p1 uint, p2 uint8) int32 {
    v4 := []int{-499, 9223372036854775807}
    for fi5 := 0; fi5 < 16; fi5++ {
        if (988 > 2147483647) {
            p1 = p1
            var v6 int = int(p1)
            _ = v6
        }
        var v7 float64 = ((float64(float32(-2.25)) - (3.14159 - float64(0.0))) - (3.14159 / float64(3.14159)))
        switch ("foo bar" + ((("a" + "foo bar") + "a") + ("" + ""))) {
        case "a":
            fmt.Println(v7)
        case "hello":
            if true {
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 float32 = 0.0
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi8, p0)
                    _ = fi8
                    _ = v9
                }
            } else {
                for fi10 := 0; fi10 < 7; fi10++ {
                    p1 -= p1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi10
                }
            }
        case "foo bar":
            p2 = 1
        }
        _ = fi5
        _ = v7
    }
    if len(v4) > 0 {
        v4[0] = 837
    }
    v11 := GenS1{F0: 255}
    switch -2147483648 {
    case 1:
        fmt.Println(p1, p0, p2)
    default:
        p0 -= 2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    return 42
}
func genFn1(p0 int32, p1 uint) int {
    v12 := []int{}
    p1 = p1
    p1 = (p1 * (p1 / ((p1) | 1)))
    _ = p0
    _ = p1
    _ = v12
    return 127
}

func main() {
    var v1 float32 = 1e10
    v2 := []uint32{2147483647, 2147483647, 1}
    var v3 int = 128
    v2 = append(v2, 762)
    if (18446744073709551615 <= 1) {
        var v13 string = "\n\t"
        fmt.Println(v13)
        v14 := GenS2{F0: -349, F1: 1}
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            v2 = append(v2, 2)
            v16 := GenS2{F0: 0, F1: 65535}
            for fi17 := 0; fi17 < 1; fi17++ {
                _ = len(v2)
                var v18 int64 = 42
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
        var v19 int16 = -185
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v19
    }
    if (-9223372036854775808 <= -863) {
        fmt.Println(v1, v1)
        if true {
            fmt.Println(v1, v1)
        } else {
            v20 := GenS3{F0: (false && (0.1 == (0.1 + float64(v1)))), F1: -1, F2: 2}
            if false {
                if len(v2) > 0 {
                    v2[0] = 417
                }
                for fi21 := 0; fi21 < 9; fi21++ {
                    v22 := GenS4{F0: 2, F1: -80}
                    _ = fi21
                    _ = v22
                }
                fmt.Println(v1, v1)
            }
            if true {
                var v23 uint32 = 1
                if (false || (!((1 == 127) || (v23 == v23)))) {
                    v1 = (((0.1 - ((0.1 / 0.1) + float32((0.1 - 0.0)))) * v1) + ((1e10 / 3.14159) + v1))
                }
                _ = v23
            } else {
                var v24 uint64 = 2
                v24 = 511
                for fi25 := 0; fi25 < 15; fi25++ {
                    v3 -= (fi25 ^ genFn1(2147483647, 2))
                    _ = fi25
                }
                _ = v24
            }
            _ = v20
        }
    }
    v2 = append(v2, 599)
    for fi26 := 0; fi26 < 14; fi26++ {
        if ((766 != 42) && true) {
            if (v1 >= v1) {
                if false {
                    var v27 int64 = -9223372036854775808
                    _ = v27
                } else {
                    fmt.Println(fi26, v3)
                    for fi28 := 0; fi28 < 5; fi28++ {
                        if ((("hello" + "") + "") <= "foo bar") {
                            fmt.Println(fi26, fi28)
                            v1 -= 3.14159
                            switch 0 {
                            case 199:
                                v29 := GenS5{F0: 2}
                                _ = v29
                            case 0:
                                if true {
                                    fmt.Println(fi28, fi28)
                                    if (((0.0 - float64(((0.1 / float32((-2.25 - (0.1 / 1e10)))) - (v1 / (v1 - float32(0.1)))))) / (1.5 / (((1.5 * 1.5) / -2.25) + (0.1 + 0.1)))) < -2.25) {
                                        v3 = -660
                                        v1 = v1
                                        if false {
                                            v3 = (v3 & -1)
                                            var v30 uint = 65535
                                            _ = v30
                                        } else {
                                            fmt.Println(fi26, v3, v3)
                                            switch "\n\t" {
                                            case "hello":
                                                v3 -= (fi26 - v3)
                                            case "foo bar":
                                                v1 = 0.1
                                            default:
                                                v31 := GenS1{F0: 1}
                                                _ = v31
                                            }
                                            if true {
                                                fmt.Println(fi28, fi28, v3)
                                            } else {
                                                var v32 float32 = (float32(((1.5 / 3.14159) - float64(v1))) + v1)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v32
                                            }
                                        }
                                    } else {
                                        var v33 float32 = (float32((3.14159 * -2.25)) * ((v1 * (float32(0.1) * 3.14159)) / ((v1 / 1.0) + ((v1 - v1) / 3.14159))))
                                        if (804 != 292) {
                                            v3 += (^(fi28))
                                        }
                                        _ = v33
                                    }
                                } else {
                                    var v34 uint32 = 2147483647
                                    fmt.Println(v1, fi26)
                                    if (float64((v1 * float32(1.5))) >= 3.14159) {
                                        _ = len(v2)
                                    }
                                    _ = v34
                                }
                            case 32768:
                                v3 = genFn1(1, 1)
                            }
                        }
                        v35 := GenS6{F0: -232}
                        v3 = (v3 / (((fi26 & (fi26 % ((125) | 1)))) | 1))
                        _ = fi28
                        _ = v35
                    }
                    v1 = v1
                }
            } else {
                v1 += v1
            }
        } else {
            switch "" {
            case "hello":
                switch ("" + "\n\t") {
                case "a":
                    fmt.Println(v1, v1, v3)
                default:
                    v3 += fi26
                }
            case "\n\t":
                if (105 != -1) {
                    var v36 uint16 = 1
                    for fi37 := 0; fi37 < 4; fi37++ {
                        v38 := GenS7{F0: true}
                        v39 := GenS8{F0: 2, F1: 42, F2: 42}
                        if true {
                            var v40 int = genFn1(435, 1)
                            _ = v40
                        } else {
                            var v41 uint64 = 1
                            var v42 float64 = 1.5
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v41
                            _ = v42
                        }
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    var v43 int32 = 0
                    _ = v36
                    _ = v43
                } else {
                    if true {
                        var v44 uint64 = 430
                        switch genFn1(-2147483648, 1) {
                        case 127:
                            v45 := GenS9{F0: 7}
                            _ = v45
                        case -242:
                            var v46 float32 = float32(float64(0.1))
                            _ = v46
                        case 801:
                            v47 := GenS2{F0: -1, F1: uint(v44)}
                            _ = v47
                        default:
                            var v48 int64 = 1
                            _ = v48
                        }
                        v3 += (v3 << ((255) & 63))
                        _ = v44
                    }
                    if false {
                        switch "" {
                        case "foo bar":
                            if true {
                                if false {
                                    v49 := GenS10{F0: (("a" + ("a" + "foo bar")) + "foo bar")}
                                    var v50 int32 = genFn0(526, 1, 255)
                                    v51 := GenS9{F0: -32768}
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                                for fi52 := 0; fi52 < 14; fi52++ {
                                    var v53 uint = 2
                                    fmt.Println(fi52, fi52)
                                    if (500 > 670) {
                                        var v54 uint32 = 2147483647
                                        if (0 > genFn0(587, (v53 >> ((v53) & 63)), 127)) {
                                            switch fi52 {
                                            case 2147483647:
                                                if true {
                                                    v55 := GenS8{F0: v54, F1: 18446744073709551615, F2: -536}
                                                    switch "" {
                                                    case "a":
                                                        if true {
                                                            v54 = (v54 / (((v54 / (((^(v54))) | 1))) | 1))
                                                            fmt.Println(v54)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            fmt.Println(v53)
                                                        }
                                                    }
                                                    v3 = (v3 * (v3 & fi52))
                                                    _ = v55
                                                } else {
                                                    v56 := GenS11{F0: "a"}
                                                    _ = v56
                                                }
                                            default:
                                                var v57 int64 = -9223372036854775808
                                                _ = v57
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = 42
                                            }
                                            switch ("foo bar" + ("\n\t" + (("\n\t" + ("" + "hello")) + "\n\t"))) {
                                            case "\n\t":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "a":
                                                var v58 int = 127
                                                _ = v58
                                            case "hello":
                                                v53 += v53
                                            default:
                                                v53 -= v53
                                                var v59 int = fi52
                                                _ = v59
                                            }
                                        }
                                        _ = v54
                                    }
                                    _ = fi52
                                    _ = v53
                                }
                                switch 32767 {
                                case 42:
                                    var v60 int16 = -1
                                    _ = v60
                                default:
                                    switch (fi26 + 0) {
                                    case 32767:
                                        var v61 int32 = -1
                                        _ = v61
                                    case 0:
                                        for fi62 := 0; fi62 < 14; fi62++ {
                                            var v63 uint32 = 697
                                            switch (fi62 + -645) {
                                            case 932:
                                                var v64 bool = true
                                                var v65 float64 = 0.0
                                                if len(v2) > 0 {
                                                    v2[0] = v63
                                                }
                                                _ = v64
                                                _ = v65
                                            case 394:
                                                var v66 uint = 255
                                                if true {
                                                    switch "a" {
                                                    case "hello":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        if (!(v1 >= v1)) {
                                                            fmt.Println(v66)
                                                        } else {
                                                            fmt.Println(v66, v66)
                                                            _ = len(v2)
                                                            fmt.Println(v66)
                                                        }
                                                    }
                                                    if true {
                                                        v67 := GenS2{F0: genFn0(18446744073709551615, (v66 & (^(v66))), 0), F1: v66}
                                                        if false {
                                                            v1 += 1.5
                                                            if true {
                                                                v1 += v1
                                                                var v68 uint8 = 255
                                                                _ = v68
                                                            }
                                                        } else {
                                                            _ = len(v2)
                                                            v69 := GenS10{F0: "a"}
                                                            v2 = append(v2, v63)
                                                            _ = v69
                                                        }
                                                        if (45 <= 42) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = len(v2)
                                                        }
                                                        _ = v67
                                                    }
                                                    v1 += (0.1 + (float32(float64(v1)) - ((3.14159 + (v1 + v1)) - (1e10 * (0.1 * v1)))))
                                                } else {
                                                    v70 := GenS12{F0: 32767}
                                                    var v71 float64 = 1e10
                                                    _ = v70
                                                    _ = v71
                                                }
                                                var v72 int64 = 42
                                                switch 3 {
                                                case -9223372036854775808:
                                                    fmt.Println(fi62, v66)
                                                case 42:
                                                    v3 += fi62
                                                case -403:
                                                    fmt.Println(v3)
                                                }
                                                fmt.Println(v72)
                                                _ = v66
                                                _ = v72
                                            case -9223372036854775808:
                                                fmt.Println(fi26, v63, fi62)
                                            }
                                            v1 += (3.14159 * (0.0 + v1))
                                            _ = fi62
                                            _ = v63
                                        }
                                    default:
                                        switch (v3 | (-(v3))) {
                                        case 42:
                                            var v73 uint32 = 449
                                            _ = v73
                                        }
                                    }
                                }
                            } else {
                                switch (v3 / ((v3) | 1)) {
                                case -1:
                                    if len(v2) > 0 {
                                        v2[0] = 2147483647
                                    }
                                case -9223372036854775808:
                                    v3 -= v3
                                default:
                                    v74 := GenS2{F0: -1, F1: 4294967295}
                                    _ = v74
                                }
                            }
                        case "hello":
                            v75 := GenS13{F0: false}
                            _ = v75
                        case "\n\t":
                            for fi76 := 0; fi76 < 4; fi76++ {
                                if (42 != 1) {
                                    if true {
                                        var v77 uint16 = 1
                                        v1 += v1
                                        var v78 float32 = 1.5
                                        _ = v77
                                        _ = v78
                                    }
                                    v1 = 1.5
                                }
                                v79 := GenS6{F0: 42}
                                var v80 bool = false
                                _ = fi76
                                _ = v79
                                _ = v80
                            }
                        default:
                            for fi81 := 0; fi81 < 17; fi81++ {
                                var v82 float64 = -2.25
                                if false {
                                    v1 += -2.25
                                    v83 := GenS14{F0: 401, F1: 21}
                                    _ = v83
                                } else {
                                    v82 += (v82 + v82)
                                }
                                fmt.Println(fi81, fi81, fi81)
                                switch "\n\t" {
                                case "a":
                                    var v84 int64 = 1
                                    _ = v84
                                default:
                                    var v85 uint = 255
                                    _ = v85
                                }
                                _ = fi81
                                _ = v82
                            }
                        }
                        v1 = (3.14159 * v1)
                        _ = len(v2)
                    } else {
                        switch (fi26 - fi26) {
                        case 955:
                            if (231 < 339) {
                                v1 += ((((v1 * (0.1 * v1)) - 3.14159) - v1) / 0.1)
                            } else {
                                switch 1 {
                                case 2147483647:
                                    fmt.Println(fi26, v3)
                                }
                                if (2 <= 211) {
                                    for fi86 := 0; fi86 < 18; fi86++ {
                                        v1 = ((0.1 + (1e10 - (0.0 / 1.0))) / (float32(0.1) + (v1 * (v1 * ((v1 - -2.25) + ((0.0 + v1) * v1))))))
                                        switch (^(fi86)) {
                                        case 127:
                                            if (32767 < 670) {
                                                var v87 uint16 = 255
                                                _ = v87
                                            } else {
                                                switch fi86 {
                                                case -768:
                                                    v88 := GenS9{F0: 1000}
                                                    _ = v88
                                                case -48:
                                                    if false {
                                                        v89 := GenS11{F0: "a"}
                                                        switch "" {
                                                        case "foo bar":
                                                            switch fi26 {
                                                            case 2147483647:
                                                                v90 := GenS5{F0: 268}
                                                                if (684 >= 32767) {
                                                                    v3 += (-(fi26))
                                                                    var v91 int64 = 9223372036854775807
                                                                    v1 += v1
                                                                    _ = v91
                                                                }
                                                                var v92 int64 = 9223372036854775807
                                                                _ = v90
                                                                _ = v92
                                                            case -9223372036854775808:
                                                                v93 := GenS4{F0: 128, F1: -1}
                                                                _ = v93
                                                            case 127:
                                                                switch "\n\t" {
                                                                case "a":
                                                                    if true {
                                                                        if true {
                                                                            var v94 string = ("foo bar" + "")
                                                                            if true {
                                                                                v1 = ((v1 * -2.25) + (v1 - ((1e10 * v1) + v1)))
                                                                            } else {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 4294967295
                                                                                }
                                                                                var v95 int16 = -1
                                                                                _ = v95
                                                                            }
                                                                            v94 = "foo bar"
                                                                            _ = v94
                                                                        }
                                                                        v96 := GenS15{F0: 4294967295}
                                                                        v1 = (1.5 + float32((1.5 / 3.14159)))
                                                                        _ = v96
                                                                    }
                                                                case "foo bar":
                                                                    var v97 float64 = ((0.0 + 3.14159) - (float64(0.1) * float64(((((v1 + v1) - 1e10) / v1) - v1))))
                                                                    _ = v97
                                                                }
                                                            }
                                                        default:
                                                            v98 := GenS16{F0: 2147483647, F1: (9223372036854775807 <= 751), F2: ("a" + ("\n\t" + "foo bar"))}
                                                            _ = v98
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = 2147483648
                                                        }
                                                        _ = v89
                                                    }
                                                default:
                                                    v1 -= 0.1
                                                }
                                                v1 -= v1
                                                var v99 string = "a"
                                                _ = v99
                                            }
                                        case 403:
                                            v100 := GenS17{F0: 2, F1: 4294967295, F2: 18446744073709551615}
                                            v3 = 32768
                                            v101 := GenS2{F0: 1, F1: 0}
                                            _ = v100
                                            _ = v101
                                        case 2147483647:
                                            var v102 int32 = -445
                                            _ = v102
                                        default:
                                            fmt.Println(fi86)
                                        }
                                        if (((float32(0.1) / v1) / (v1 + ((v1 - 0.1) / v1))) < 3.14159) {
                                            v3 = fi86
                                            fmt.Println(v1, fi86)
                                            if (-1 <= 703) {
                                                v1 -= (v1 / v1)
                                                var v103 int64 = 1
                                                _ = v103
                                            } else {
                                                if false {
                                                    v104 := GenS18{F0: 42, F1: 1}
                                                    v3 = 1
                                                    _ = v104
                                                } else {
                                                    var v105 float32 = v1
                                                    _ = v105
                                                }
                                                var v106 int16 = -32768
                                                _ = v106
                                            }
                                            if true {
                                                var v107 uint32 = 2147483647
                                                _ = v107
                                            } else {
                                                v2 = append(v2, 2)
                                                if (-14 >= -128) {
                                                    v3 = 9223372036854775807
                                                    if (true && (2 > 42)) {
                                                        if true {
                                                            fmt.Println(fi86)
                                                        } else {
                                                            var v108 bool = (((1.5 * 0.1) < (float32(3.14159) / ((v1 / ((1e10 + v1) - (v1 + 0.0))) * (3.14159 + 1.5)))) && false)
                                                            fmt.Println(fi86)
                                                            _ = v108
                                                        }
                                                        v1 -= v1
                                                        if (2 < 357) {
                                                            if true {
                                                                var v109 uint16 = 0
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v109
                                                            } else {
                                                                var v110 uint32 = 133
                                                                var v111 uint16 = 32767
                                                                switch -865 {
                                                                case 0:
                                                                    v2 = append(v2, (^(v110)))
                                                                }
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                            v3 = (fi26 | (fi86 << ((65535) & 63)))
                                                            v1 -= v1
                                                        }
                                                    }
                                                    v112 := GenS19{F0: 2147483647, F1: 40, F2: 614}
                                                    _ = v112
                                                }
                                            }
                                        }
                                        fmt.Println(fi26)
                                        _ = fi86
                                    }
                                }
                                _ = len(v2)
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
            case "a":
                switch "foo bar" {
                case "":
                    switch genFn1(995, 454) {
                    case 562:
                        var v113 string = ("hello" + "foo bar")
                        _ = v113
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "hello":
                    v1 = 0.1
                case "foo bar":
                    v114 := GenS12{F0: 1000}
                    _ = v114
                }
            }
            if true {
                if true {
                    if (32767 == 42) {
                        switch 127 {
                        case 239:
                            var v115 uint64 = 2
                            _ = v115
                        }
                        var v116 float64 = 0.1
                        _ = v116
                    }
                }
                v1 = (v1 / (((float32((0.1 + 0.0)) / 1.0) - 1e10) - float32(3.14159)))
            } else {
                v117 := GenS14{F0: 0, F1: 18}
                _ = v117
            }
            if true {
                for fi118 := 0; fi118 < 16; fi118++ {
                    v1 -= (float32((float64(1.5) * 0.1)) * 0.1)
                    if (!(((3.14159 + 0.1) == (-2.25 - 1.5)) || false)) {
                        if false {
                            switch ((((("" + (("a" + "hello") + (("hello" + ("\n\t" + ("a" + "a"))) + "hello"))) + "a") + ("foo bar" + "a")) + ("foo bar" + "foo bar")) + "foo bar") {
                            case "\n\t":
                                _ = len(v2)
                            case "foo bar":
                                v119 := GenS20{F0: -1, F1: 32767}
                                v120 := GenS21{F0: 18446744073709551615, F1: 2, F2: -32768}
                                _ = v119
                                _ = v120
                            default:
                                switch -2147483648 {
                                case 2147483647:
                                    if ((1 == 0) && ((1.5 * 0.1) < 1e10)) {
                                        var v121 uint64 = 2
                                        _ = v121
                                    }
                                    v1 -= 0.1
                                    var v122 int64 = -1
                                    _ = v122
                                case 954:
                                    if len(v2) > 0 {
                                        v2[0] = 751
                                    }
                                }
                            }
                            v1 += (float32(1.5) + v1)
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 295
                            }
                        }
                        if false {
                            var v123 uint64 = 42
                            v1 += (1e10 - float32(-2.25))
                            if false {
                                v1 -= (v1 + (((v1 + (v1 / 1e10)) * v1) + 1e10))
                                v123 = v123
                                switch (("hello" + "") + "\n\t") {
                                case "foo bar":
                                    v123 -= v123
                                default:
                                    var v124 int8 = 0
                                    _ = v124
                                }
                            }
                            if (827 >= 551) {
                                if false {
                                    v1 -= 0.1
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v123, v123, v123)
                                    case "\n\t":
                                        v123 = 18446744073709551615
                                    case "foo bar":
                                        switch "a" {
                                        case "a":
                                            var v125 int64 = 47
                                            v123 = (v123 - 355)
                                            _ = len(v2)
                                            _ = v125
                                        case "hello":
                                            var v126 uint16 = 42
                                            v126 -= 255
                                            switch "" {
                                            case "\n\t":
                                                v127 := GenS7{F0: false}
                                                _ = v127
                                            default:
                                                var v128 uint8 = 2
                                                v129 := GenS8{F0: 0, F1: v123, F2: 42}
                                                _ = v128
                                                _ = v129
                                            }
                                            _ = v126
                                        case "foo bar":
                                            var v130 bool = true
                                            _ = v130
                                        }
                                    default:
                                        var v131 uint32 = 2
                                        _ = v131
                                    }
                                } else {
                                    var v132 float64 = 1.5
                                    if (-128 >= 1) {
                                        v133 := GenS17{F0: 2147483647, F1: 782, F2: (v123 & (v123 << ((42) & 63)))}
                                        _ = len(v2)
                                        var v134 uint64 = 641
                                        _ = v133
                                        _ = v134
                                    } else {
                                        fmt.Println(v3)
                                    }
                                    v132 += -2.25
                                    _ = v132
                                }
                                switch "\n\t" {
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                }
                                if false {
                                    fmt.Println(fi26)
                                }
                                v2 = append(v2, 42)
                                v3 = (fi118 % ((42) | 1))
                                v3 += (^(fi26))
                            }
                            _ = v123
                        }
                        fmt.Println(v1, v1)
                    } else {
                        var v135 uint64 = 0
                        var v136 string = ((((("a" + (("a" + "\n\t") + "\n\t")) + "\n\t") + "a") + (("" + ("a" + "foo bar")) + ("" + "a"))) + ("a" + "hello"))
                        _ = v135
                        _ = v136
                    }
                    _ = fi118
                }
            } else {
                var v137 float32 = v1
                switch "hello" {
                case "foo bar":
                    v1 -= v137
                case "a":
                    v138 := GenS9{F0: 1000}
                    _ = v138
                case "":
                    var v139 int = genFn1(0, 309)
                    _ = v139
                default:
                    v1 += v137
                }
                _ = v137
            }
        }
        _ = len(v2)
        switch genFn1(0, 65535) {
        case 4294967295:
            v1 -= (v1 - float32((float64((v1 / (v1 + (0.1 * v1)))) / (1e10 + (((((1e10 * 0.1) - (1.5 / (-2.25 / -2.25))) * 0.1) * 1.5) + (0.0 * 0.1))))))
        case 32767:
            v1 += (((((v1 * float32(1e10)) * 1.5) * v1) / 3.14159) / 1.5)
        case 128:
            if true {
                v1 -= v1
            }
        default:
            if false {
                v3 = 32767
                v140 := GenS22{F0: 65535, F1: 857}
                if true {
                    v1 -= (float32(1e10) * v1)
                    var v141 float64 = (1.5 / 3.14159)
                    _ = v141
                } else {
                    fmt.Println(v1, v3, v3)
                    var v142 int8 = 86
                    var v143 int32 = 0
                    _ = v142
                    _ = v143
                }
                _ = v140
            } else {
                if true {
                    fmt.Println(v1, v3, v3)
                } else {
                    v1 += 0.0
                    switch v3 {
                    case 128:
                        if ((861 > 1) && (false || (1 >= 42))) {
                            var v144 string = "a"
                            _ = v144
                        }
                    case 580:
                        for fi145 := 0; fi145 < 6; fi145++ {
                            v146 := GenS10{F0: "foo bar"}
                            var v147 int = -478
                            fmt.Println(fi26)
                            v147 = 127
                            _ = fi145
                            _ = v146
                            _ = v147
                        }
                    case 42:
                        v1 -= -2.25
                    }
                }
                v1 += -2.25
                v3 = (v3 & 1)
            }
        }
        _ = fi26
    }
    v1 += v1
    if (true || (!(42 >= 11))) {
        fmt.Println(v3)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v148 int64 = 1
        _ = v148
    }
    if true {
        switch genFn1(0, 65535) {
        case 9223372036854775807:
            _ = len(v2)
        }
        v1 += v1
    } else {
        v3 += v3
        switch "" {
        case "foo bar":
            switch v3 {
            case 0:
                if (!(!(0 > 280))) {
                    if true {
                        v1 += (v1 * v1)
                    }
                    switch (-(v3)) {
                    case 461:
                        v3 = -9223372036854775808
                    case 32768:
                        switch "" {
                        case "foo bar":
                            var v149 uint64 = 42
                            _ = v149
                        default:
                            for fi150 := 0; fi150 < 3; fi150++ {
                                var v151 int = (^(fi150))
                                if false {
                                    v3 = v3
                                    v151 += v3
                                } else {
                                    v151 -= genFn1(1, 1)
                                }
                                v3 = 113
                                _ = fi150
                                _ = v151
                            }
                        }
                    case 2147483647:
                        v3 += (v3 * (v3 + 42))
                    }
                    var v152 bool = (430 > 744)
                    _ = v152
                }
            case -300:
                for fi153 := 0; fi153 < 3; fi153++ {
                    for fi154 := 0; fi154 < 17; fi154++ {
                        v155 := GenS5{F0: 2}
                        _ = fi154
                        _ = v155
                    }
                    _ = fi153
                }
            default:
                v156 := GenS21{F0: 792, F1: 4294967295, F2: 115}
                _ = v156
            }
        case "\n\t":
            switch (v3 * v3) {
            case 32767:
                v1 += (v1 / 0.1)
            case -818:
                v157 := GenS18{F0: -1, F1: 127}
                _ = v157
            }
        case "a":
            var v158 bool = false
            _ = v158
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
