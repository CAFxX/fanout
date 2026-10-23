package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 int32 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int32; F1 uint16; F2 bool }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 uint; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint; F1 bool; F2 uint8 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS20 struct { F0 int; F1 uint64; F2 int16 }

func main() {
    v1 := []uint32{889, 2147483647, 2147483647, 1}
    v2 := []string{}
    var v3 int16 = -1
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    if (!false) {
        fmt.Println(v3)
        fmt.Println(v3)
    } else {
        for fi6 := 0; fi6 < 13; fi6++ {
            if (222 < 255) {
                switch "foo bar" {
                case "a":
                    fmt.Println(v3, v3)
                default:
                    fmt.Println(v3, v3)
                }
                for fi7 := 0; fi7 < 11; fi7++ {
                    v8 := GenS1{F0: 0, F1: 32767}
                    delete(v5, "hello")
                    fmt.Println(v3)
                    _ = fi7
                    _ = v8
                }
                if (-374 > (v3 - (v3 / ((-21) | 1)))) {
                    if true {
                        v9 := GenS2{F0: 2147483647, F1: (fi6 | fi6)}
                        v3 += 32767
                        _ = v9
                    } else {
                        if ((false || false) || false) {
                            fmt.Println(fi6, fi6)
                            fmt.Println(v3, fi6)
                        } else {
                            var v10 uint32 = 2
                            v3 = v3
                            fmt.Println(v10, fi6)
                            _ = v10
                        }
                    }
                    fmt.Println(fi6, v3)
                } else {
                    var v11 uint8 = 255
                    switch fi6 {
                    case 32768:
                        _ = len(v1)
                    case 22:
                        switch "foo bar" {
                        case "hello":
                            var v12 int = 9223372036854775807
                            _ = v12
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            for fi13 := 0; fi13 < 8; fi13++ {
                                var v14 uint8 = (v11 >> ((18446744073709551615) & 63))
                                _ = fi13
                                _ = v14
                            }
                        }
                    }
                    _ = v11
                }
            } else {
                v3 = (v3 & 32767)
                var v15 uint64 = 18446744073709551615
                var v16 uint16 = 84
                _ = v15
                _ = v16
            }
            fmt.Println(fi6, v3)
            var v17 int16 = v3
            _ = fi6
            _ = v17
        }
        v18 := GenS3{F0: (v3 >> ((255) & 63))}
        if (true || false) {
            v19 := GenS1{F0: -3, F1: (v3 % ((v3) | 1))}
            _ = v19
        }
        _ = v18
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println(v3)
    for fi20 := 0; fi20 < 18; fi20++ {
        switch (fi20 - fi20) {
        case 42:
            switch 1 {
            case 1:
                v3 -= (^(v3))
            case 238:
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 uint64 = 0
                    v23 := GenS4{F0: int8(v22), F1: v3, F2: 813}
                    if false {
                        switch ("foo bar" + "") {
                        case "hello":
                            if true {
                                if true {
                                    switch (^(fi20)) {
                                    case 74:
                                        if false {
                                            switch "a" {
                                            case "":
                                                _ = len(v5)
                                            case "hello":
                                                v24 := GenS3{F0: -303}
                                                if (true || true) {
                                                    var v25 uint = 601
                                                    switch ("" + "") {
                                                    case "":
                                                        fmt.Println(v22, fi20)
                                                    case "hello":
                                                        if len(v2) > 0 {
                                                            v2[0] = "a"
                                                        }
                                                    }
                                                    _ = v25
                                                } else {
                                                    switch "" {
                                                    case "":
                                                        if true {
                                                            if true {
                                                                switch 127 {
                                                                case -341:
                                                                    var v26 float64 = float64(1e10)
                                                                    _ = v26
                                                                case -9223372036854775808:
                                                                    _ = len(v4)
                                                                case 127:
                                                                    fmt.Println(v22, fi20)
                                                                }
                                                            }
                                                            var v27 float32 = float32(-2.25)
                                                            _ = v27
                                                        } else {
                                                            if true {
                                                                if (1e10 >= 1e10) {
                                                                    var v28 int32 = 42
                                                                    v28 = (v28 - -1)
                                                                    _ = v28
                                                                } else {
                                                                    v3 -= (v3 ^ -32768)
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        var v29 int = fi21
                                                        _ = v29
                                                    }
                                                    v30 := GenS5{F0: false, F1: (("foo bar" + "a") + ("" + ("hello" + "hello")))}
                                                    _ = v30
                                                }
                                                var v31 string = "\n\t"
                                                _ = v24
                                                _ = v31
                                            case "foo bar":
                                                switch ("a" + "foo bar") {
                                                case "foo bar":
                                                    v3 += (v3 + v3)
                                                default:
                                                    if false {
                                                        v32 := GenS6{F0: 42, F1: 32767, F2: (128 >= 128)}
                                                        v33 := GenS2{F0: -1, F1: fi20}
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        var v34 int32 = int32(fi20)
                                                        _ = v34
                                                    }
                                                }
                                            default:
                                                v35 := GenS3{F0: v3}
                                                _ = v35
                                            }
                                            var v36 uint64 = 809
                                            switch fi21 {
                                            case 9223372036854775807:
                                                var v37 uint8 = 0
                                                if false {
                                                    v37 = (v37 | v37)
                                                    fmt.Println(v3, v22, v37)
                                                }
                                                _ = v37
                                            case 42:
                                                v36 -= v36
                                            case -9223372036854775808:
                                                v38 := GenS7{F0: false, F1: (v36 << ((196) & 63))}
                                                _ = v38
                                            default:
                                                switch 9223372036854775807 {
                                                case 2147483647:
                                                    v22 = (v36 + (v36 << ((255) & 63)))
                                                    var v39 int = (fi20 / ((fi20) | 1))
                                                    v40 := GenS8{F0: -52, F1: (v3 ^ 0)}
                                                    _ = len(v2)
                                                    if ((v39 % (((-(fi21))) | 1)) > 1) {
                                                        if len(v2) > 0 {
                                                            v2[0] = "a"
                                                        }
                                                        v39 += (v39 ^ 2147483647)
                                                        v41 := GenS7{F0: (3.14159 < -2.25), F1: v36}
                                                        v3 -= 1
                                                        switch (v39 - (fi21 << ((42) & 63))) {
                                                        case -372:
                                                            if (("a" <= "a") && (!false)) {
                                                                v42 := GenS1{F0: (fi20 ^ (fi20 * fi21)), F1: 42}
                                                                v3 += v3
                                                                _ = v42
                                                            }
                                                        case 561:
                                                            var v43 uint32 = 304
                                                            _ = v43
                                                        }
                                                        _ = v41
                                                    }
                                                    _ = v39
                                                    _ = v40
                                                case 58:
                                                    v44 := GenS6{F0: -2147483648, F1: 42, F2: false}
                                                    var v45 int16 = (v3 >> ((2) & 63))
                                                    v45 -= (v3 / ((-1) | 1))
                                                    _ = v44
                                                    _ = v45
                                                }
                                                if (((v36 % ((1) | 1)) == 1) && (32768 > (fi21 & fi21))) {
                                                    v46 := v5[("foo bar" + "a")]
                                                    v3 = 191
                                                    _ = v46
                                                }
                                                fmt.Println(fi21, v36, v22)
                                            }
                                            _ = v36
                                        }
                                    case 128:
                                        if ((2 >= 540) && true) {
                                            var v47 bool = (42 >= -2147483648)
                                            v3 -= v3
                                            v3 -= v3
                                            _ = v47
                                        } else {
                                            var v48 int8 = 8
                                            var v49 uint8 = 31
                                            _ = v48
                                            _ = v49
                                        }
                                    case 2147483647:
                                        var v50 int64 = 0
                                        _ = v50
                                    default:
                                        var v51 uint = 18446744073709551615
                                        _ = v51
                                    }
                                    fmt.Println(v3)
                                } else {
                                    switch "" {
                                    case "a":
                                        var v52 uint64 = v22
                                        _ = v52
                                    case "\n\t":
                                        v22 = (^(v22))
                                    }
                                    fmt.Println(fi20, fi21, fi20)
                                    v53 := GenS3{F0: -535}
                                    var v54 uint8 = 2
                                    _ = v53
                                    _ = v54
                                }
                                var v55 float32 = 3.14159
                                _ = v55
                            } else {
                                if false {
                                    switch 32767 {
                                    case 9223372036854775807:
                                        v3 = -1
                                    default:
                                        v56 := v4["\n\t"]
                                        _ = v56
                                    }
                                }
                                delete(v5, "\n\t")
                                var v57 bool = (4294967295 == 2)
                                _ = v57
                            }
                        case "":
                            if len(v1) > 0 {
                                v1[0] = 885
                            }
                        default:
                            var v58 uint32 = 42
                            _ = v58
                        }
                    } else {
                        var v59 int64 = 42
                        if true {
                            var v60 int8 = 0
                            var v61 uint64 = v22
                            _ = v60
                            _ = v61
                        } else {
                            switch ((((((("hello" + "hello") + ((("foo bar" + (("\n\t" + "\n\t") + ("a" + "a"))) + "a") + "hello")) + "foo bar") + ("\n\t" + ("hello" + "foo bar"))) + "a") + "\n\t") + "foo bar") {
                            case "\n\t":
                                fmt.Println(fi20, v3, fi20)
                            default:
                                if true {
                                    v62 := GenS9{F0: 42, F1: (1 >= 20)}
                                    var v63 uint8 = 255
                                    if (((("foo bar" + "a") < "hello") && false) || (65535 != uint32(v63))) {
                                        v22 = v22
                                        var v64 float32 = (-2.25 + float32(1.5))
                                        _ = v64
                                    } else {
                                        delete(v4, ("a" + "hello"))
                                        if false {
                                            var v65 bool = true
                                            if (18446744073709551615 < (v22 + 720)) {
                                                _ = len(v4)
                                                v63 = (v63 - 128)
                                                switch "" {
                                                case "\n\t":
                                                    if v65 {
                                                        var v66 int16 = -599
                                                        v1 = append(v1, 2)
                                                        _ = v66
                                                    }
                                                case "":
                                                    v4["\n\t"] = (("hello" + "\n\t") + "foo bar")
                                                }
                                            } else {
                                                v63 += (v63 ^ 2)
                                                if false {
                                                    v59 -= int64(fi20)
                                                }
                                            }
                                            switch fi21 {
                                            case 32768:
                                                if (!(-1 <= -2147483648)) {
                                                    if (33 > fi21) {
                                                        if v65 {
                                                            v2 = append(v2, "foo bar")
                                                            switch (fi20 ^ (fi21 >> ((2) & 63))) {
                                                            case 0:
                                                                v63 = 42
                                                            }
                                                            if false {
                                                                v63 += (^(v63))
                                                                var v67 int32 = 42
                                                                _ = v67
                                                            } else {
                                                                var v68 int = (fi21 ^ int(v59))
                                                                if len(v1) > 0 {
                                                                    v1[0] = 2
                                                                }
                                                                _ = v68
                                                            }
                                                        }
                                                        v69 := v4[("" + ("\n\t" + "hello"))]
                                                        v59 = v59
                                                        _ = v69
                                                    }
                                                    v59 -= v59
                                                    v70 := GenS7{F0: (("hello" + (("hello" + "\n\t") + "foo bar")) <= "hello"), F1: 250}
                                                    _ = v70
                                                }
                                            default:
                                                v71 := GenS10{F0: (v63 % ((v63) | 1))}
                                                _ = v71
                                            }
                                            _ = v65
                                        }
                                        if true {
                                            switch "a" {
                                            case "foo bar":
                                                v72 := GenS5{F0: (2147483647 != 172), F1: "hello"}
                                                var v73 uint = 4294967295
                                                var v74 bool = (v22 != uint64(v73))
                                                if (678 != 255) {
                                                    if (("hello" + "hello") < "") {
                                                        if (255 == (v73 + 1)) {
                                                            if false {
                                                                if (!false) {
                                                                    switch fi21 {
                                                                    case 214:
                                                                        _ = len(v2)
                                                                    default:
                                                                        v75 := GenS11{F0: 931}
                                                                        _ = v75
                                                                    }
                                                                    v76 := GenS8{F0: 100, F1: v3}
                                                                    v63 = v63
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        fmt.Println(v59, fi20, v73)
                                                                    case "\n\t":
                                                                        v63 = (v63 / ((v63) | 1))
                                                                    default:
                                                                        var v77 uint64 = (v22 + (v22 & v22))
                                                                        _ = v77
                                                                    }
                                                                    v78 := GenS5{F0: false, F1: "hello"}
                                                                    _ = v76
                                                                    _ = v78
                                                                }
                                                            }
                                                            fmt.Println(v63)
                                                            v2 = append(v2, ("hello" + ""))
                                                            v3 -= int16(v22)
                                                        } else {
                                                            switch "a" {
                                                            case "foo bar":
                                                                v73 -= 65535
                                                            }
                                                            var v79 int8 = 1
                                                            _ = v79
                                                        }
                                                        var v80 float64 = 3.14159
                                                        var v81 float64 = (v80 * float64(-2.25))
                                                        v80 = 0.1
                                                        v82 := GenS12{F0: 32767}
                                                        switch fi21 {
                                                        case 1:
                                                            if true {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 649
                                                                }
                                                                v80 = float64(1e10)
                                                                fmt.Println(fi21)
                                                            } else {
                                                                if (!v74) {
                                                                    if ("hello" >= "\n\t") {
                                                                        v80 = v80
                                                                        if (100 > 44) {
                                                                            var v83 uint64 = (v22 * (v22 << ((v73) & 63)))
                                                                            _ = v83
                                                                        }
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            if (!((v73 != v73) || true)) {
                                                                                fmt.Println(v81)
                                                                            }
                                                                        case "hello":
                                                                            switch (fi21 | (-(fi20))) {
                                                                            case 1:
                                                                                v80 -= ((v80 + 3.14159) / v81)
                                                                            }
                                                                        default:
                                                                            v59 -= v59
                                                                        }
                                                                        delete(v4, "foo bar")
                                                                        fmt.Println(v22, v22, v80)
                                                                        v22 = 207
                                                                    }
                                                                    v84 := GenS8{F0: 100, F1: (v3 | (v3 >> (((^(v73))) & 63)))}
                                                                    if ((v3 >> ((v73) & 63)) != (v3 % (((v3 - 0)) | 1))) {
                                                                        var v85 uint64 = v22
                                                                        v5["foo bar"] = 2147483647
                                                                        _ = v85
                                                                    } else {
                                                                        fmt.Println(v63, fi20, v74)
                                                                    }
                                                                    _ = v84
                                                                }
                                                                switch fi21 {
                                                                case 4294967295:
                                                                    var v86 uint16 = 255
                                                                    _ = v86
                                                                case 1:
                                                                    fmt.Println(v59)
                                                                default:
                                                                    var v87 int64 = 1
                                                                    if false {
                                                                        var v88 float32 = float32(float64(float32(v81)))
                                                                        v59 = v59
                                                                        var v89 uint32 = 0
                                                                        if false {
                                                                            v90 := GenS13{F0: (v63 / ((245) | 1))}
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v2 = append(v2, "a")
                                                                            }
                                                                            _ = v90
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                fmt.Println(v3, v74)
                                                                            case "a":
                                                                                v3 = (^(v3))
                                                                            }
                                                                        }
                                                                        if true {
                                                                            var v91 uint8 = v63
                                                                            v91 = uint8(v73)
                                                                            _ = v91
                                                                        }
                                                                        _ = v88
                                                                        _ = v89
                                                                    }
                                                                    var v92 int = (fi20 / ((-265) | 1))
                                                                    _ = v87
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                        var v93 int16 = (v3 / ((v3) | 1))
                                                        v1 = append(v1, 0)
                                                        _ = v80
                                                        _ = v81
                                                        _ = v82
                                                        _ = v93
                                                    }
                                                    var v94 int = (fi21 & fi20)
                                                    v95 := GenS12{F0: 2}
                                                    _ = v94
                                                    _ = v95
                                                }
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            case "\n\t":
                                                switch ("foo bar" + "\n\t") {
                                                case "\n\t":
                                                    var v96 uint64 = v22
                                                    _ = v96
                                                case "hello":
                                                    var v97 bool = ((65535 != 105) && false)
                                                    _ = v97
                                                case "a":
                                                    v5[("" + ("\n\t" + ("\n\t" + "")))] = 200
                                                }
                                            }
                                            fmt.Println(fi20)
                                        } else {
                                            v98 := GenS8{F0: -44, F1: (v3 + (v3 >> ((4294967295) & 63)))}
                                            fmt.Println(v59)
                                            _ = v98
                                        }
                                        v99 := GenS14{F0: 255}
                                        v63 = (v63 % (((v63 >> ((183) & 63))) | 1))
                                        if true {
                                            v3 += int16(v22)
                                        }
                                        _ = v99
                                    }
                                    _ = v62
                                    _ = v63
                                } else {
                                    v22 -= v22
                                    if false {
                                        var v100 uint16 = 32767
                                        v101 := GenS15{F0: 42, F1: (v100 >> ((1) & 63)), F2: (^(v3))}
                                        v3 = v3
                                        _ = v100
                                        _ = v101
                                    }
                                    fmt.Println(v3, v3, v3)
                                    var v102 float32 = 1e10
                                    var v103 uint32 = 90
                                    _ = v102
                                    _ = v103
                                }
                            }
                        }
                        var v104 uint = 18446744073709551615
                        _ = v104
                        _ = v59
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            default:
                var v105 bool = false
                _ = v105
            }
        default:
            v106 := GenS16{F0: (fi20 << ((2) & 63))}
            _ = v106
        }
        v3 = v3
        v3 = v3
        _ = fi20
    }
    if false {
        for fi107 := 0; fi107 < 18; fi107++ {
            v3 = v3
            v108 := GenS5{F0: (!(-2147483648 != 980)), F1: "hello"}
            if (fi107 <= fi107) {
                v109 := GenS17{F0: 1, F1: (false && (1 != 255)), F2: 255}
                v3 = 1
                _ = v109
            }
            _ = fi107
            _ = v108
        }
        if (!(true && (!(false || false)))) {
            var v110 uint16 = 1
            switch 153 {
            case 127:
                v110 = (^(v110))
            default:
                v110 -= v110
            }
            v111 := v4[""]
            _ = v110
            _ = v111
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 371
        }
        var v112 int32 = -1
        v113 := GenS8{F0: 7, F1: (v3 | 761)}
        _ = v112
        _ = v113
    }
    v114 := v4["hello"]
    v115 := GenS18{F0: (!true)}
    var v116 int16 = v3
    v117 := GenS8{F0: -81, F1: -32768}
    if true {
        if ((v3 & (v3 | (v3 ^ v3))) <= (v116 >> ((0) & 63))) {
            v114 = ("hello" + v114)
            v118 := GenS19{F0: 1, F1: 760, F2: 2147483647}
            _ = v118
        }
        switch -1 {
        case -361:
            switch "" {
            case "foo bar":
                v4[((v114 + "a") + (v114 + "foo bar"))] = v114
            case "hello":
                var v119 uint8 = 2
                _ = v119
            }
        case 1:
            var v120 float32 = 0.0
            _ = v120
        default:
            v121 := GenS20{F0: 2147483647, F1: 781, F2: v3}
            _ = v121
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v114
    _ = v115
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
