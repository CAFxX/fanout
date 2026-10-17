package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 int }
type GenS3 struct { F0 uint32; F1 bool; F2 int16 }
type GenS4 struct { F0 uint8; F1 int; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 int16; F1 uint16; F2 string }
type GenS10 struct { F0 int8 }
func genFn0(p0 bool, p1 int64, p2 uint8) string {
    if (-1 < 550) {
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                p1 += 700
                v9 := GenS1{F0: 118, F1: p1, F2: p1}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{fi10, (^(fi10))}
            fmt.Println(p1, fi10, fi10)
            _ = fi10
            _ = v11
        }
        p1 -= (^(p1))
    }
    fmt.Println(p0)
    p1 -= (-(p1))
    var v12 float32 = (0.1 - 1.5)
    var v13 int = 306
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return ""
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = 0.1
    var v3 bool = true
    var v4 uint16 = 32767
    var v5 uint = 2
    v6 := []uint32{0}
    switch 4294967295 {
    case 9223372036854775807:
        if (232 < -1) {
            switch ("hello" + ("a" + "a")) {
            case "a":
                fmt.Println(v1)
            case "foo bar":
                v14 := GenS1{F0: (v1 - 127), F1: -9223372036854775808, F2: -427}
                _ = v14
            case "hello":
                var v15 float32 = (1e10 + float32(v2))
                _ = v15
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            for fi16 := 0; fi16 < 4; fi16++ {
                if true {
                    v1 -= 2
                    fmt.Println(v4, v1)
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 uint16 = 0
                        _ = fi17
                        _ = v18
                    }
                }
                switch int(v4) {
                case 42:
                    if false {
                        v1 = 112
                        v6 = append(v6, 2147483648)
                        if false {
                            v19 := GenS2{F0: (fi16 ^ fi16), F1: 0, F2: (fi16 | fi16)}
                            if ((!((v4 | 2) <= v4)) || v3) {
                                v5 += (v5 << ((42) & 63))
                                v20 := GenS3{F0: 669, F1: v3, F2: 0}
                                _ = v20
                            }
                            _ = v19
                        } else {
                            if (((^(v4)) != v4) && ((!v3) && (!((^(v1)) != v1)))) {
                                v21 := GenS1{F0: (v1 % (((v1 | 1)) | 1)), F1: -9223372036854775808, F2: 0}
                                if false {
                                    if (v3 && (!(v5 < (^(v5))))) {
                                        fmt.Println(fi16, v2)
                                    } else {
                                        v22 := GenS1{F0: 152, F1: 240, F2: 0}
                                        v23 := GenS2{F0: fi16, F1: (v4 % (((v4 << ((uint(v4)) & 63))) | 1)), F2: (-(fi16))}
                                        fmt.Println(fi16)
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                fmt.Println(fi16, v5, v5)
                                _ = v21
                            } else {
                                if true {
                                    var v24 int8 = -128
                                    v25 := GenS1{F0: (^(v1)), F1: int64(v24), F2: 1}
                                    if false {
                                        for fi26 := 0; fi26 < 4; fi26++ {
                                            v27 := GenS3{F0: 155, F1: v3, F2: 0}
                                            if true {
                                                fmt.Println(v5, v4)
                                                switch "\n\t" {
                                                case "hello":
                                                    if false {
                                                        if v3 {
                                                            if len(v6) > 0 {
                                                                v6[0] = 42
                                                            }
                                                            v6 = append(v6, 442)
                                                        }
                                                        var v28 int64 = -9223372036854775808
                                                        _ = v28
                                                    }
                                                    switch int(v4) {
                                                    case 42:
                                                        var v29 uint64 = 4294967295
                                                        _ = v29
                                                    case 128:
                                                        var v30 uint16 = (v4 - 309)
                                                        _ = v30
                                                    case 900:
                                                        var v31 float64 = (-2.25 - ((float64(0.0) + v2) / v2))
                                                        _ = v31
                                                    }
                                                    fmt.Println(fi26, v3, fi16)
                                                default:
                                                    v32 := GenS4{F0: v1, F1: (fi26 + (fi26 % ((32768) | 1))), F2: fi16}
                                                    _ = v32
                                                }
                                            } else {
                                                var v33 float64 = v2
                                                var v34 uint32 = 962
                                                _ = v33
                                                _ = v34
                                            }
                                            var v35 float32 = 1e10
                                            _ = fi26
                                            _ = v27
                                            _ = v35
                                        }
                                        v1 = (v1 * (v1 % ((v1) | 1)))
                                        v36 := GenS5{F0: v1}
                                        _ = v36
                                    }
                                    _ = v24
                                    _ = v25
                                } else {
                                    if (42 <= -20) {
                                        switch "foo bar" {
                                        case "hello":
                                            switch (fi16 - fi16) {
                                            case 32768:
                                                v37 := GenS6{F0: (v4 ^ v4), F1: 187, F2: 865}
                                                _ = v37
                                            default:
                                                fmt.Println(v1, v4, v3)
                                            }
                                        }
                                        _ = len(v6)
                                        v2 += float64(0.0)
                                    } else {
                                        v6 = append(v6, 42)
                                        var v38 int16 = 1
                                        _ = v38
                                    }
                                    var v39 int16 = 32767
                                    _ = v39
                                }
                                var v40 int32 = -725
                                _ = v40
                            }
                            if (0 < 1) {
                                v1 = (^(v1))
                            }
                        }
                    } else {
                        var v41 int8 = 16
                        _ = v41
                    }
                }
                v42 := GenS5{F0: (v1 / (((v1 >> ((v5) & 63))) | 1))}
                _ = fi16
                _ = v42
            }
            if len(v6) > 0 {
                v6[0] = 70
            }
        }
    }
    var v43 float64 = v2
    v3 = (false || false)
    fmt.Println(v1, v4)
    switch (("hello" + ("a" + "hello")) + ("foo bar" + (("hello" + ("a" + "foo bar")) + (("" + "\n\t") + "foo bar")))) {
    case "hello":
        fmt.Println(v2, v43)
    case "foo bar":
        v6 = append(v6, 343)
    case "a":
        fmt.Println(v5, v3, v3)
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        fmt.Println(fi44, v5, v4)
        if (2 >= 65535) {
            switch fi44 {
            case -2147483648:
                v1 -= uint8(v5)
            }
            var v45 uint32 = 2147483648
            _ = v45
        } else {
            switch -346 {
            case -341:
                fmt.Println(fi44)
            case -1:
                _ = len(v6)
            case 32767:
                v43 = v43
            default:
                for fi46 := 0; fi46 < 2; fi46++ {
                    var v47 uint8 = v1
                    _ = fi46
                    _ = v47
                }
            }
            for fi48 := 0; fi48 < 11; fi48++ {
                v49 := GenS2{F0: fi44, F1: 65535, F2: (fi44 + -9223372036854775808)}
                v50 := GenS7{F0: 127}
                _ = fi48
                _ = v49
                _ = v50
            }
            v2 += float64(((0.0 / -2.25) / 3.14159))
        }
        _ = fi44
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        _ = len(v6)
        _ = len(v6)
        var v52 uint32 = 287
        _ = fi51
        _ = v52
    }
    v3 = v3
    if v3 {
        var v53 int8 = 127
        _ = v53
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch "hello" {
    case "foo bar":
        switch "foo bar" {
        case "a":
            for fi54 := 0; fi54 < 17; fi54++ {
                fmt.Println(v4, v2)
                if (776 <= 428) {
                    _ = len(v6)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    fmt.Println(v43, v4)
                    var v55 int = fi54
                    _ = v55
                }
                _ = fi54
            }
        default:
            if (v43 >= (v2 - -2.25)) {
                switch 530 {
                case -586:
                    for fi56 := 0; fi56 < 14; fi56++ {
                        var v57 int = (fi56 % ((426) | 1))
                        for fi58 := 0; fi58 < 7; fi58++ {
                            if ((!(2147483647 >= -147)) && (false || v3)) {
                                var v59 uint32 = 383
                                _ = v59
                            }
                            _ = fi58
                        }
                        _ = fi56
                        _ = v57
                    }
                case 0:
                    if (18446744073709551615 >= 42) {
                        if ("hello" == "") {
                            if (!false) {
                                switch 2147483647 {
                                case 32767:
                                    for fi60 := 0; fi60 < 1; fi60++ {
                                        var v61 uint64 = uint64(v5)
                                        switch fi60 {
                                        case 152:
                                            var v62 int = fi60
                                            _ = v62
                                        case 2147483647:
                                            v2 -= 1.5
                                        }
                                        switch "hello" {
                                        case "":
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v63 uint32 = 77
                                                _ = v63
                                            }
                                        }
                                        _ = fi60
                                        _ = v61
                                    }
                                default:
                                    for fi64 := 0; fi64 < 2; fi64++ {
                                        v65 := GenS1{F0: uint8(v4), F1: 9223372036854775807, F2: 1}
                                        for fi66 := 0; fi66 < 16; fi66++ {
                                            var v67 float32 = 3.14159
                                            _ = fi66
                                            _ = v67
                                        }
                                        _ = fi64
                                        _ = v65
                                    }
                                }
                                var v68 uint32 = 936
                                _ = v68
                            } else {
                                if (29 <= 2147483648) {
                                    v69 := GenS4{F0: v1, F1: 9223372036854775807, F2: 128}
                                    if (!false) {
                                        if true {
                                            fmt.Println(v5, v43)
                                            fmt.Println(v1)
                                        } else {
                                            var v70 uint16 = 0
                                            for fi71 := 0; fi71 < 14; fi71++ {
                                                var v72 uint = (v5 & v5)
                                                if (1.5 == ((0.1 / float32(v43)) / ((-2.25 / 0.1) + (3.14159 * 3.14159)))) {
                                                    v73 := GenS7{F0: 49}
                                                    _ = v73
                                                } else {
                                                    for fi74 := 0; fi74 < 6; fi74++ {
                                                        if (0 <= 42) {
                                                            v43 -= 0.1
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        } else {
                                                            var v75 uint32 = 42
                                                            var v76 uint8 = 128
                                                            var v77 string = (("" + ("foo bar" + (("hello" + ("foo bar" + "\n\t")) + "a"))) + (("" + "foo bar") + (("foo bar" + ("a" + "foo bar")) + "foo bar")))
                                                            _ = v75
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        _ = fi74
                                                    }
                                                }
                                                _ = fi71
                                                _ = v72
                                            }
                                            _ = v70
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            var v78 uint16 = (^(v4))
                                            _ = v78
                                        case "a":
                                            if (0 <= 2147483647) {
                                                v79 := GenS4{F0: (v1 + v1), F1: 624, F2: 32767}
                                                _ = v79
                                            }
                                        default:
                                            for fi80 := 0; fi80 < 4; fi80++ {
                                                v81 := GenS8{F0: uint(v1), F1: (^(v4))}
                                                v2 += 0.0
                                                _ = fi80
                                                _ = v81
                                            }
                                        }
                                    }
                                    _ = v69
                                } else {
                                    for fi82 := 0; fi82 < 9; fi82++ {
                                        for fi83 := 0; fi83 < 10; fi83++ {
                                            fmt.Println(v5)
                                            fmt.Println(fi82, fi83, v5)
                                            _ = fi83
                                        }
                                        v84 := GenS1{F0: (^(v1)), F1: 0, F2: 9223372036854775807}
                                        _ = fi82
                                        _ = v84
                                    }
                                }
                                v1 = (v1 << (((v5 % ((v5) | 1))) & 63))
                                for fi85 := 0; fi85 < 7; fi85++ {
                                    for fi86 := 0; fi86 < 8; fi86++ {
                                        switch 183 {
                                        case 86:
                                            if (!(((((0.0 * 0.1) * 0.1) / 1e10) <= 1.5) || (v5 > (v5 | 255)))) {
                                                if true {
                                                    v87 := GenS1{F0: (^(v1)), F1: 29, F2: 1}
                                                    _ = v87
                                                } else {
                                                    if v3 {
                                                        fmt.Println(fi86, v1, v4)
                                                        var v88 bool = (v3 && (!(v3 && ((1 != 445) && (false || (750 > 1))))))
                                                        _ = v88
                                                    }
                                                    v6 = append(v6, 132)
                                                }
                                                var v89 float32 = 3.14159
                                                _ = v89
                                            } else {
                                                if (false || true) {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    if false {
                                                        v43 = 0.1
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        switch -2147483648 {
                                                        case 0:
                                                            fmt.Println(v1, fi85, fi86)
                                                            v4 = (v4 + v4)
                                                            _ = len(v6)
                                                        case 42:
                                                            var v90 int32 = 0
                                                            _ = v90
                                                        }
                                                    }
                                                }
                                                v4 = 319
                                                v3 = true
                                            }
                                        }
                                        var v91 int32 = -1
                                        v43 += float64(float32(v2))
                                        _ = fi86
                                        _ = v91
                                    }
                                    v6 = append(v6, 65535)
                                    _ = len(v6)
                                    _ = fi85
                                }
                            }
                        } else {
                            for fi92 := 0; fi92 < 5; fi92++ {
                                if len(v6) > 0 {
                                    v6[0] = 42
                                }
                                var v93 uint32 = 42
                                for fi94 := 0; fi94 < 14; fi94++ {
                                    var v95 uint32 = (v93 | v93)
                                    v95 = v93
                                    if v3 {
                                        if false {
                                            v96 := GenS5{F0: 127}
                                            var v97 int16 = 1000
                                            v98 := GenS8{F0: v5, F1: (v4 >> ((923) & 63))}
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        }
                                        fmt.Println(fi94, v95)
                                        v1 -= (^(v1))
                                        v5 = (v5 << (((^(v5))) & 63))
                                    }
                                    v99 := GenS9{F0: 42, F1: v4, F2: "a"}
                                    _ = fi94
                                    _ = v95
                                    _ = v99
                                }
                                _ = fi92
                                _ = v93
                            }
                            v100 := GenS4{F0: (v1 >> ((v5) & 63)), F1: int(v5), F2: -2147483648}
                            _ = v100
                        }
                    } else {
                        if v3 {
                            var v101 string = "hello"
                            switch "\n\t" {
                            case "hello":
                                var v102 int64 = -47
                                _ = v102
                            case "\n\t":
                                for fi103 := 0; fi103 < 17; fi103++ {
                                    v101 = ("" + v101)
                                    _ = fi103
                                }
                            case "":
                                v2 = (v43 + (v2 - v43))
                            }
                            for fi104 := 0; fi104 < 18; fi104++ {
                                switch -2147483648 {
                                case -1:
                                    v2 += ((v2 + 0.1) + float64(float32(v43)))
                                case 907:
                                    for fi105 := 0; fi105 < 2; fi105++ {
                                        fmt.Println(v43, fi104, v43)
                                        v101 = v101
                                        _ = fi105
                                    }
                                }
                                v101 = ("hello" + "foo bar")
                                fmt.Println(v3, v4)
                                _ = fi104
                            }
                            _ = v101
                        }
                        if true {
                            v106 := GenS1{F0: (^(v1)), F1: -1, F2: 976}
                            if len(v6) > 0 {
                                v6[0] = 890
                            }
                            v107 := GenS9{F0: 32767, F1: v4, F2: "hello"}
                            _ = v106
                            _ = v107
                        } else {
                            _ = len(v6)
                            fmt.Println(v3, v2)
                        }
                    }
                default:
                    if (!(((v1 << (((v5 >> (((v5 + v5)) & 63))) & 63)) > 255) && ((!(9223372036854775807 > -118)) && v3))) {
                        if true {
                            var v108 uint = (^(v5))
                            fmt.Println(v4)
                            if (v3 && true) {
                                var v109 uint16 = (^(v4))
                                _ = v109
                            } else {
                                v110 := GenS1{F0: (v1 % ((v1) | 1)), F1: 0, F2: 21}
                                v2 += ((float64(1.5) / v2) * (v43 * v43))
                                v43 -= float64(0.0)
                                _ = v110
                            }
                            _ = v108
                        } else {
                            switch 32767 {
                            case 60:
                                v4 = v4
                            case 32767:
                                if (2 > v5) {
                                    for fi111 := 0; fi111 < 17; fi111++ {
                                        if len(v6) > 0 {
                                            v6[0] = 218
                                        }
                                        if v3 {
                                            v3 = v3
                                            var v112 int8 = 32
                                            switch "hello" {
                                            case "\n\t":
                                                v6 = append(v6, 508)
                                            }
                                            _ = v112
                                        }
                                        _ = fi111
                                    }
                                    v113 := GenS6{F0: v4, F1: 1, F2: 4294967295}
                                    if v3 {
                                        fmt.Println(v43, v1, v43)
                                    }
                                    _ = v113
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = uint32(v4)
                                    }
                                    for fi114 := 0; fi114 < 6; fi114++ {
                                        switch (("foo bar" + "foo bar") + "\n\t") {
                                        case "":
                                            if (!v3) {
                                                v6 = append(v6, 2)
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                        default:
                                            if v3 {
                                                var v115 int32 = -2147483648
                                                for fi116 := 0; fi116 < 14; fi116++ {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v117 float64 = (1.5 - v2)
                                                    _ = fi116
                                                    _ = v117
                                                }
                                                for fi118 := 0; fi118 < 13; fi118++ {
                                                    var v119 float32 = (float32(v43) * -2.25)
                                                    _ = fi118
                                                    _ = v119
                                                }
                                                _ = v115
                                            } else {
                                                var v120 uint64 = 42
                                                _ = v120
                                            }
                                        }
                                        v4 = 65535
                                        v2 = v43
                                        _ = fi114
                                    }
                                    fmt.Println(v5)
                                }
                            case -1:
                                _ = len(v6)
                            }
                            v3 = v3
                        }
                        var v121 uint = 18446744073709551615
                        switch ("\n\t" + "hello") {
                        case "a":
                            var v122 int8 = -1
                            _ = v122
                        default:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        _ = v121
                    }
                }
            } else {
                if false {
                    if v3 {
                        switch -9223372036854775808 {
                        case -2147483648:
                            var v123 uint = 696
                            _ = v123
                        case 2147483647:
                            if true {
                                for fi124 := 0; fi124 < 13; fi124++ {
                                    var v125 int = (fi124 - 2147483647)
                                    _ = fi124
                                    _ = v125
                                }
                            }
                        }
                        fmt.Println(v43, v3, v4)
                    } else {
                        v2 = float64(-2.25)
                        if ((v2 >= (v2 - v2)) || v3) {
                            v6 = append(v6, 2147483648)
                        }
                        for fi126 := 0; fi126 < 15; fi126++ {
                            v127 := GenS10{F0: 100}
                            var v128 uint32 = 65535
                            v6 = append(v6, v128)
                            _ = fi126
                            _ = v127
                            _ = v128
                        }
                    }
                } else {
                    var v129 uint64 = 0
                    var v130 uint64 = v129
                    _ = v129
                    _ = v130
                }
                fmt.Println(v4, v1)
                for fi131 := 0; fi131 < 19; fi131++ {
                    var v132 bool = ("" == ("" + "foo bar"))
                    v2 = (0.1 * (0.1 + (v43 * (float64(1e10) * (v43 + -2.25)))))
                    if true {
                        var v133 uint = v5
                        for fi134 := 0; fi134 < 14; fi134++ {
                            var v135 float64 = v43
                            v6 = append(v6, 42)
                            _ = fi134
                            _ = v135
                        }
                        _ = v133
                    }
                    _ = fi131
                    _ = v132
                }
            }
        }
    default:
        var v136 int = 2147483647
        _ = v136
    }
    var v137 uint8 = (v1 | v1)
    var v138 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v137
    _ = v138
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
