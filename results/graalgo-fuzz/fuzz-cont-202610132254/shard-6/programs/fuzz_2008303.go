package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 int8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int8; F1 uint16; F2 int64 }
type GenS6 struct { F0 int8; F1 uint8; F2 string }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int64; F1 uint8; F2 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool; F1 uint8 }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 uint64; F1 int32 }
type GenS14 struct { F0 int64; F1 int8; F2 string }
type GenS15 struct { F0 int64; F1 int8 }
type GenS16 struct { F0 int32; F1 uint16 }

func main() {
    var v1 float64 = 1.5
    var v2 uint32 = 42
    var v3 uint = 255
    v4 := []uint32{v2, 1, v2}
    var v5 float32 = (0.1 + 0.0)
    switch (("foo bar" + "foo bar") + ("\n\t" + "\n\t")) {
    case "a":
        for fi6 := 0; fi6 < 19; fi6++ {
            if (18446744073709551615 > 4294967295) {
                var v7 uint16 = uint16(v2)
                _ = v7
            } else {
                v4 = append(v4, (v2 & v2))
                for fi8 := 0; fi8 < 12; fi8++ {
                    fmt.Println(v3, v3)
                    v3 += uint(v2)
                    if ("" > (("foo bar" + "foo bar") + ("foo bar" + ("foo bar" + "\n\t")))) {
                        switch (fi6 * (fi6 + fi8)) {
                        case -2147483648:
                            v4 = append(v4, v2)
                        case 9223372036854775807:
                            v2 -= v2
                        case 127:
                            v9 := GenS1{F0: 127}
                            _ = v9
                        default:
                            if len(v4) > 0 {
                                v4[0] = v2
                            }
                            v10 := GenS2{F0: fi6, F1: (v3 | 4294967295), F2: 127}
                            var v11 float32 = 0.1
                            _ = v10
                            _ = v11
                        }
                        v12 := GenS1{F0: 0}
                        var v13 float64 = (((v1 * 1e10) + float64((v5 / v5))) / 0.1)
                        _ = v12
                        _ = v13
                    }
                    _ = fi8
                }
                if false {
                    v14 := GenS3{F0: true, F1: 254}
                    v2 = (^(v2))
                    switch ("a" + "") {
                    case "hello":
                        if (57 <= 42) {
                            if true {
                                v1 = 1e10
                                v3 += (v3 & (v3 * v3))
                            }
                        } else {
                            for fi15 := 0; fi15 < 15; fi15++ {
                                _ = len(v4)
                                fmt.Println(v5, fi6)
                                _ = fi15
                            }
                            switch fi6 {
                            case 128:
                                switch (("a" + "a") + "a") {
                                case "foo bar":
                                    var v16 int8 = 101
                                    _ = v16
                                case "":
                                    fmt.Println(v3)
                                case "\n\t":
                                    for fi17 := 0; fi17 < 20; fi17++ {
                                        fmt.Println(v5)
                                        var v18 uint16 = 32767
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi17
                                        _ = v18
                                    }
                                }
                            case -2147483648:
                                if false {
                                    v19 := GenS3{F0: true, F1: v2}
                                    var v20 bool = false
                                    if (!true) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        v1 += (v1 / v1)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v19
                                    _ = v20
                                } else {
                                    if (!(true || false)) {
                                        if false {
                                            v2 += (v2 >> ((0) & 63))
                                        } else {
                                            var v21 uint32 = v2
                                            _ = v21
                                        }
                                        v1 += v1
                                        switch ("hello" + "hello") {
                                        case "":
                                            switch "hello" {
                                            case "\n\t":
                                                var v22 bool = (!(((0 >= 291) || ((v2 + (^(v2))) == v2)) || false))
                                                _ = v22
                                            case "":
                                                v4 = append(v4, (v2 / ((v2) | 1)))
                                            case "hello":
                                                v1 += 3.14159
                                            }
                                        default:
                                            var v23 uint = 577
                                            _ = v23
                                        }
                                    }
                                    if (fi6 < fi6) {
                                        var v24 uint16 = 685
                                        for fi25 := 0; fi25 < 5; fi25++ {
                                            var v26 uint = v3
                                            _ = fi25
                                            _ = v26
                                        }
                                        _ = v24
                                    } else {
                                        v2 = v2
                                        v1 = v1
                                        v2 += (v2 % ((v2) | 1))
                                    }
                                    var v27 int32 = 2147483647
                                    _ = v27
                                }
                            case 869:
                                if true {
                                    if (!true) {
                                        v2 += v2
                                        switch (fi6 >> ((v3) & 63)) {
                                        case 2147483647:
                                            for fi28 := 0; fi28 < 7; fi28++ {
                                                var v29 uint64 = 837
                                                var v30 uint32 = (v2 * (v2 / (((v2 >> ((v3) & 63))) | 1)))
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = fi28
                                                _ = v29
                                                _ = v30
                                            }
                                        case -461:
                                            if false {
                                                switch fi6 {
                                                case 9223372036854775807:
                                                    if true {
                                                        if false {
                                                            v31 := GenS2{F0: fi6, F1: v3, F2: 100}
                                                            _ = v31
                                                        } else {
                                                            if (!(true || ((4294967295 >= 65535) || true))) {
                                                                v32 := GenS4{F0: true, F1: 121}
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                v33 := GenS1{F0: 127}
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                            v4 = append(v4, v2)
                                                            var v34 int8 = -128
                                                            _ = v34
                                                        }
                                                        for fi35 := 0; fi35 < 12; fi35++ {
                                                            var v36 uint32 = 2
                                                            var v37 string = ("a" + (("" + "") + ""))
                                                            _ = fi35
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        v1 += (((((v1 + (float64(-2.25) * v1)) - 0.1) + (((((((((v1 * v1) * v1) * 1e10) - (float64(0.0) + v1)) * float64(v5)) - (v1 - (1e10 + (1e10 / 3.14159)))) / v1) + (0.0 - 0.1)) - 3.14159)) * 1e10) - v1)
                                                    } else {
                                                        for fi38 := 0; fi38 < 11; fi38++ {
                                                            v39 := GenS3{F0: false, F1: (v2 * (v2 / ((42) | 1)))}
                                                            _ = fi38
                                                            _ = v39
                                                        }
                                                        switch (fi6 + fi6) {
                                                        case 2147483647:
                                                            fmt.Println(v5, v3, fi6)
                                                        default:
                                                            var v40 uint32 = 370
                                                            _ = v40
                                                        }
                                                    }
                                                default:
                                                    var v41 int16 = -1
                                                    _ = v41
                                                }
                                                v42 := GenS3{F0: (!(2 != 255)), F1: 2147483647}
                                                switch ("foo bar" + ("\n\t" + "hello")) {
                                                case "\n\t":
                                                    v1 += v1
                                                case "foo bar":
                                                    if false {
                                                        v43 := GenS5{F0: 42, F1: 563, F2: 1}
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v43
                                                    } else {
                                                        switch (fi6 >> ((1) & 63)) {
                                                        case 128:
                                                            var v44 uint8 = 250
                                                            _ = v44
                                                        }
                                                    }
                                                }
                                                _ = v42
                                            }
                                        case -9223372036854775808:
                                            for fi45 := 0; fi45 < 14; fi45++ {
                                                var v46 uint = (v3 * (v3 | (v3 << ((v3) & 63))))
                                                _ = fi45
                                                _ = v46
                                            }
                                        default:
                                            v47 := GenS5{F0: -108, F1: 1, F2: 0}
                                            _ = v47
                                        }
                                        switch (("" + "a") + ("hello" + "\n\t")) {
                                        case "\n\t":
                                            v1 = ((0.1 - v1) * v1)
                                        default:
                                            for fi48 := 0; fi48 < 6; fi48++ {
                                                v2 = v2
                                                fmt.Println(v2)
                                                v2 -= v2
                                                _ = fi48
                                            }
                                        }
                                    } else {
                                        v49 := GenS1{F0: 2}
                                        _ = v49
                                    }
                                    var v50 string = ""
                                    for fi51 := 0; fi51 < 18; fi51++ {
                                        var v52 int8 = 127
                                        _ = len(v4)
                                        switch v50 {
                                        case "foo bar":
                                            v1 -= 0.0
                                        case "a":
                                            if (32767 < 255) {
                                                v53 := GenS2{F0: (fi6 + (fi51 >> ((uint(v2)) & 63))), F1: v3, F2: v52}
                                                fmt.Println(fi6, fi51)
                                                if (42 >= 42) {
                                                    v50 = ""
                                                    v54 := GenS6{F0: 100, F1: 0, F2: "hello"}
                                                    if (false || false) {
                                                        v55 := GenS7{F0: (v2 / ((v2) | 1)), F1: (v52 * (v52 % (((^(v52))) | 1)))}
                                                        switch (fi6 * -1) {
                                                        case -2147483648:
                                                            var v56 uint8 = 247
                                                            _ = v56
                                                        case 195:
                                                            var v57 int = (fi6 & fi6)
                                                            _ = v57
                                                        case 127:
                                                            var v58 int32 = 418
                                                            _ = v58
                                                        default:
                                                            v50 = v50
                                                        }
                                                        _ = v55
                                                    } else {
                                                        var v59 float64 = v1
                                                        if false {
                                                            v60 := GenS8{F0: -495, F1: (fi6 / ((fi51) | 1))}
                                                            _ = v60
                                                        } else {
                                                            v61 := GenS9{F0: -517, F1: 0, F2: 494}
                                                            var v62 uint32 = (v2 / ((665) | 1))
                                                            if (255 <= 0) {
                                                                var v63 float64 = (((1e10 - (v1 + v59)) * 0.0) * 0.1)
                                                                _ = v63
                                                            }
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        _ = v59
                                                    }
                                                    _ = v54
                                                } else {
                                                    var v64 string = v50
                                                    v4 = append(v4, (v2 >> ((v3) & 63)))
                                                    _ = v64
                                                }
                                                _ = v53
                                            } else {
                                                var v65 int = -2147483648
                                                v5 = (v5 - (((3.14159 * v5) - v5) + (v5 * ((v5 * (float32(1e10) - v5)) - float32((0.1 * 0.1))))))
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v66 := GenS1{F0: 119}
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                        _ = fi51
                                        _ = v52
                                    }
                                    _ = v50
                                }
                            default:
                                v67 := GenS4{F0: (0 < 42), F1: -2147483648}
                                _ = v67
                            }
                        }
                    case "foo bar":
                        var v68 int64 = 1
                        _ = v68
                    default:
                        if true {
                            var v69 int32 = -2147483648
                            switch "foo bar" {
                            case "a":
                                for fi70 := 0; fi70 < 14; fi70++ {
                                    fmt.Println(v3, v2, v2)
                                    var v71 uint = v3
                                    _ = fi70
                                    _ = v71
                                }
                            case "\n\t":
                                var v72 float64 = (v1 / v1)
                                _ = v72
                            case "":
                                for fi73 := 0; fi73 < 16; fi73++ {
                                    v74 := GenS10{F0: (3.14159 == v5)}
                                    if false {
                                        var v75 uint16 = 2
                                        var v76 int8 = 42
                                        if ((v1 - (v1 - float64(v5))) > float64(1e10)) {
                                            v77 := GenS9{F0: 1, F1: 156, F2: -449}
                                            if (false || (!false)) {
                                                v78 := GenS11{F0: false, F1: 1}
                                                var v79 int8 = (v76 << ((v3) & 63))
                                                v69 = (v69 - (v69 & (v69 / ((v69) | 1))))
                                                _ = v78
                                                _ = v79
                                            } else {
                                                if true {
                                                    if ((-(fi73)) < (fi6 % ((fi6) | 1))) {
                                                        fmt.Println(v76, v1, fi73)
                                                    }
                                                } else {
                                                    switch ("a" + "hello") {
                                                    case "foo bar":
                                                        v80 := GenS12{F0: 18446744073709551615, F1: (v75 << ((v3) & 63))}
                                                        _ = v80
                                                    default:
                                                        v2 -= 338
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = (^(v2))
                                                    }
                                                    var v81 string = ""
                                                    _ = v81
                                                }
                                                var v82 int16 = 32767
                                                v83 := GenS13{F0: 2, F1: 0}
                                                _ = v82
                                                _ = v83
                                            }
                                            if true {
                                                if ((v76 - (-(v76))) > v76) {
                                                    fmt.Println(v69, v3)
                                                    v84 := GenS14{F0: 9223372036854775807, F1: 33, F2: ""}
                                                    _ = v84
                                                }
                                                switch fi73 {
                                                case -9223372036854775808:
                                                    var v85 uint32 = uint32(v3)
                                                    _ = v85
                                                case 127:
                                                    fmt.Println(v1, v1)
                                                default:
                                                    v86 := GenS2{F0: fi73, F1: (v3 % (((^(v3))) | 1)), F2: 42}
                                                    _ = v86
                                                }
                                                var v87 bool = true
                                                _ = v87
                                            } else {
                                                var v88 bool = false
                                                _ = v88
                                            }
                                            _ = v77
                                        } else {
                                            v3 -= 65535
                                            _ = len(v4)
                                            if ((fi73 >> ((2) & 63)) <= 9223372036854775807) {
                                                fmt.Println(v5, v2)
                                            } else {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    _ = len(v4)
                                                case "":
                                                    if true {
                                                        v3 = 1
                                                    } else {
                                                        fmt.Println(v75)
                                                    }
                                                }
                                                if (!true) {
                                                    v5 += v5
                                                    v4 = append(v4, (v2 / ((4294967295) | 1)))
                                                    v69 = -819
                                                    switch (("a" + ("" + "")) + ("foo bar" + ("\n\t" + "a"))) {
                                                    case "a":
                                                        var v89 int8 = (v76 | v76)
                                                        _ = v89
                                                    case "":
                                                        v90 := GenS12{F0: 2, F1: (v75 ^ (v75 % ((v75) | 1)))}
                                                        _ = v90
                                                    case "hello":
                                                        var v91 uint8 = 128
                                                        _ = v91
                                                    default:
                                                        if true {
                                                            fmt.Println(v5, v1)
                                                            v5 = v5
                                                            if (("\n\t" < ("hello" + "\n\t")) && (4294967295 <= (fi73 - (fi73 << (((^(v3))) & 63))))) {
                                                                v4 = append(v4, 42)
                                                                fmt.Println(v69, v76, v76)
                                                            } else {
                                                                switch (fi73 - fi6) {
                                                                case -570:
                                                                    v92 := GenS7{F0: (v2 * 0), F1: 127}
                                                                    _ = v92
                                                                default:
                                                                    var v93 int = (fi6 % (((fi6 | fi6)) | 1))
                                                                    _ = v93
                                                                }
                                                                if (-2147483648 >= (v69 | 42)) {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = v2
                                                                    }
                                                                    var v94 uint32 = 2
                                                                    if false {
                                                                        var v95 uint64 = uint64(v3)
                                                                        if true {
                                                                            if true {
                                                                                v5 += v5
                                                                                v94 = 42
                                                                                _ = len(v4)
                                                                            }
                                                                            v69 -= (v69 & v69)
                                                                        } else {
                                                                            v75 = v75
                                                                        }
                                                                        var v96 int64 = -253
                                                                        v97 := GenS6{F0: (v76 - v76), F1: 128, F2: ("" + "\n\t")}
                                                                        _ = v95
                                                                        _ = v96
                                                                        _ = v97
                                                                    }
                                                                    _ = v94
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v98 := GenS9{F0: 42, F1: 127, F2: 0}
                                                    _ = v98
                                                }
                                            }
                                        }
                                        _ = v75
                                        _ = v76
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            var v99 uint32 = (^(v2))
                                            _ = v99
                                        case "hello":
                                            var v100 int = 282
                                            _ = v100
                                        }
                                        v101 := GenS5{F0: 42, F1: 255, F2: 0}
                                        _ = v101
                                    }
                                    _ = fi73
                                    _ = v74
                                }
                            default:
                                for fi102 := 0; fi102 < 6; fi102++ {
                                    if false {
                                        v5 += 0.1
                                        v69 = (v69 % (((v69 % (((v69 & 128)) | 1))) | 1))
                                    }
                                    _ = fi102
                                }
                            }
                            v103 := GenS2{F0: fi6, F1: 4294967295, F2: -100}
                            _ = v103
                            _ = v69
                        }
                    }
                    _ = v14
                }
            }
            var v104 int16 = 42
            v5 -= 1.5
            _ = fi6
            _ = v104
        }
    case "foo bar":
        var v105 int = -36
        _ = v105
    case "hello":
        if ((v2 >> ((v3) & 63)) >= v2) {
            v2 = (v2 % ((v2) | 1))
        }
    default:
        v3 = (^(v3))
    }
    v5 += v5
    var v106 uint32 = v2
    for fi107 := 0; fi107 < 11; fi107++ {
        var v108 int = (fi107 + fi107)
        switch (v108 & (-(v108))) {
        case 2147483647:
            fmt.Println(v1)
        case 32767:
            var v109 uint8 = 204
            _ = v109
        }
        var v110 uint16 = 255
        _ = fi107
        _ = v108
        _ = v110
    }
    if true {
        v3 = uint(v2)
        v111 := GenS2{F0: 128, F1: v3, F2: 100}
        var v112 int64 = 9223372036854775807
        _ = v111
        _ = v112
    } else {
        switch -779 {
        case -562:
            v5 = v5
        case 32768:
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
        case -983:
            var v113 int = 32768
            _ = v113
        default:
            v114 := GenS12{F0: uint64(v2), F1: 569}
            _ = v114
        }
    }
    v115 := GenS15{F0: 42, F1: 111}
    v106 -= v106
    v106 += v2
    if false {
        fmt.Println(v106, v3)
    }
    v116 := GenS16{F0: -1, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v106
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
