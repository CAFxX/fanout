package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 string; F1 uint; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 uint32 }
type GenS11 struct { F0 uint; F1 int16; F2 int32 }
type GenS12 struct { F0 int8; F1 string }
type GenS13 struct { F0 int64; F1 int8; F2 int8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int64; F1 int32; F2 int16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int32; F1 bool; F2 uint16 }
type GenS18 struct { F0 int; F1 int16; F2 uint8 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint16; F1 int64; F2 int16 }
type GenS21 struct { F0 int }
type GenS22 struct { F0 uint32; F1 int }
type GenS23 struct { F0 uint32; F1 bool; F2 int64 }
type GenS24 struct { F0 int32; F1 uint8 }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 uint64; F1 int64 }
type GenS27 struct { F0 bool }
type GenS28 struct { F0 uint32; F1 string }
type GenS29 struct { F0 int64; F1 int; F2 uint64 }
type GenS30 struct { F0 uint32; F1 uint8 }
type GenS31 struct { F0 int }

func main() {
    v1 := []string{"foo bar", ("hello" + "foo bar"), ("" + "hello")}
    var v2 bool = true
    var v3 int8 = 46
    if (v2 && ((!(42 != 0)) && (!v2))) {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        _ = len(v1)
        v4 := GenS1{F0: 1000, F1: 32767, F2: 42}
        _ = v4
    }
    _ = len(v1)
    v5 := GenS1{F0: 32767, F1: 1, F2: 463}
    if ((v3 & v3) > v3) {
        if ((v3 - (v3 - v3)) >= (-(v3))) {
            switch 32767 {
            case -1:
                v2 = ((1 <= 42) || (389 >= 2))
            case 1:
                var v6 float32 = 0.0
                _ = v6
            }
            fmt.Println(v3, v3)
            v3 = v3
        } else {
            switch 128 {
            case -2147483648:
                if true {
                    if (-1 <= -919) {
                        for fi7 := 0; fi7 < 8; fi7++ {
                            v3 += -1
                            _ = fi7
                        }
                        for fi8 := 0; fi8 < 1; fi8++ {
                            if true {
                                if v2 {
                                    switch fi8 {
                                    case 9223372036854775807:
                                        var v9 float64 = 0.0
                                        _ = v9
                                    case 32767:
                                        var v10 int32 = 833
                                        _ = v10
                                    case 482:
                                        var v11 uint32 = 1
                                        _ = v11
                                    }
                                }
                                if (78 > 2) {
                                    for fi12 := 0; fi12 < 3; fi12++ {
                                        var v13 bool = true
                                        v14 := GenS1{F0: -661, F1: 667, F2: 896}
                                        _ = fi12
                                        _ = v13
                                        _ = v14
                                    }
                                    var v15 string = ("" + "")
                                    _ = v15
                                } else {
                                    fmt.Println(v2, fi8)
                                    var v16 uint = 255
                                    _ = v16
                                }
                            }
                            fmt.Println(v2)
                            v17 := GenS1{F0: 42, F1: 983, F2: 928}
                            _ = fi8
                            _ = v17
                        }
                        var v18 uint16 = 32767
                        _ = v18
                    }
                } else {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 uint32 = 5
                        v3 -= (v3 ^ v3)
                        var v21 bool = false
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    var v22 uint64 = 289
                    if v2 {
                        var v23 bool = (2147483647 == 42)
                        if (((42 == -343) || (!(1 < v22))) && (1 > 0)) {
                            for fi24 := 0; fi24 < 20; fi24++ {
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    var v26 float64 = -2.25
                                    v22 = (v22 / ((v22) | 1))
                                    if (2147483647 < 42) {
                                        var v27 uint64 = v22
                                        v26 = v26
                                        if v2 {
                                            switch "\n\t" {
                                            case "":
                                                v28 := GenS1{F0: 1, F1: 255, F2: 4294967295}
                                                _ = v28
                                            case "\n\t":
                                                fmt.Println(fi24)
                                            }
                                            if (0 <= (v22 >> ((4294967295) & 63))) {
                                                v27 += 18446744073709551615
                                            }
                                        }
                                        _ = v27
                                    } else {
                                        fmt.Println(v22, v23, v26)
                                        fmt.Println(fi24, v2, fi24)
                                        v3 = -128
                                        switch fi24 {
                                        case 505:
                                            var v29 float32 = (0.0 - ((1.5 - float32(v26)) - -2.25))
                                            _ = v29
                                        }
                                        switch fi24 {
                                        case 4294967295:
                                            var v30 string = ""
                                            _ = v30
                                        case -936:
                                            switch 127 {
                                            case 585:
                                                v22 = v22
                                            case 2147483647:
                                                fmt.Println(v3, v2, v3)
                                            default:
                                                fmt.Println(fi25, v3)
                                            }
                                        }
                                        if v2 {
                                            var v31 uint16 = uint16(v22)
                                            _ = v31
                                        } else {
                                            switch fi25 {
                                            case 9223372036854775807:
                                                if true {
                                                    v23 = true
                                                } else {
                                                    v32 := GenS1{F0: 1000, F1: 0, F2: 4294967295}
                                                    _ = v32
                                                }
                                            case -128:
                                                _ = len(v1)
                                            case 128:
                                                if v2 {
                                                    v26 -= (-2.25 - v26)
                                                } else {
                                                    if v23 {
                                                        v33 := GenS1{F0: 42, F1: 492, F2: 0}
                                                        var v34 uint64 = (^(v22))
                                                        if true {
                                                            if true {
                                                                var v35 int32 = -762
                                                                fmt.Println(v3)
                                                                var v36 int16 = 0
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                            if (32 >= (v3 * (-(v3)))) {
                                                                v37 := GenS1{F0: 1000, F1: 0, F2: 442}
                                                                var v38 bool = (v2 || (!v23))
                                                                v39 := GenS1{F0: 0, F1: 481, F2: 728}
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            } else {
                                                                v40 := GenS1{F0: 1, F1: 840, F2: 495}
                                                                _ = v40
                                                            }
                                                            v3 = (v3 << ((65535) & 63))
                                                            v1 = append(v1, "")
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v2 = v23
                                                        }
                                                        switch (("\n\t" + "a") + "foo bar") {
                                                        case "\n\t":
                                                            switch (fi24 ^ fi24) {
                                                            case 491:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case 755:
                                                                _ = len(v1)
                                                            default:
                                                                v41 := GenS1{F0: 42, F1: 0, F2: 0}
                                                                _ = v41
                                                            }
                                                        case "foo bar":
                                                            v34 += (v22 << ((4294967295) & 63))
                                                        }
                                                        var v42 uint = 65535
                                                        _ = v33
                                                        _ = v34
                                                        _ = v42
                                                    } else {
                                                        var v43 uint64 = (v22 | v22)
                                                        if v2 {
                                                            v44 := GenS1{F0: 547, F1: 1, F2: 4294967295}
                                                            _ = v44
                                                        } else {
                                                            var v45 uint = 0
                                                            v26 -= -2.25
                                                            switch fi25 {
                                                            case 583:
                                                                v23 = true
                                                            case -9223372036854775808:
                                                                var v46 uint16 = 1
                                                                _ = v46
                                                            }
                                                            _ = v45
                                                        }
                                                        var v47 bool = (1 >= -1)
                                                        _ = v43
                                                        _ = v47
                                                    }
                                                }
                                                v1 = append(v1, ("hello" + ""))
                                                fmt.Println(v22, v22)
                                            default:
                                                v48 := GenS1{F0: 1, F1: 838, F2: 901}
                                                _ = v48
                                            }
                                            var v49 uint8 = 115
                                            switch (("" + ("hello" + ("hello" + ("" + "\n\t")))) + "a") {
                                            case "foo bar":
                                                var v50 int = fi25
                                                _ = v50
                                            case "":
                                                v51 := GenS2{F0: 2, F1: (v3 / ((v3) | 1)), F2: 255}
                                                _ = v51
                                            case "\n\t":
                                                v52 := GenS1{F0: -1, F1: 82, F2: 255}
                                                _ = v52
                                            default:
                                                var v53 string = (("hello" + ("a" + "")) + "\n\t")
                                                v2 = v23
                                                if v23 {
                                                    v53 = v53
                                                    if false {
                                                        fmt.Println(v3, v2, fi24)
                                                        switch fi25 {
                                                        case 128:
                                                            var v54 float64 = v26
                                                            _ = v54
                                                        case 32768:
                                                            var v55 uint64 = (v22 % (((^(v22))) | 1))
                                                            _ = v55
                                                        }
                                                        fmt.Println(v23, v2, v26)
                                                    }
                                                    v56 := GenS1{F0: 32767, F1: 622, F2: 22}
                                                    switch -631 {
                                                    case 680:
                                                        switch -2147483648 {
                                                        case 127:
                                                            v2 = (!(int(v49) < (fi25 & fi25)))
                                                        }
                                                        fmt.Println(v49, v2, fi24)
                                                        var v57 bool = (2147483647 >= 117)
                                                        _ = v57
                                                    case -628:
                                                        v49 -= 42
                                                    }
                                                    v1 = append(v1, (v53 + v53))
                                                    _ = v56
                                                }
                                                var v58 uint32 = 65535
                                                var v59 int8 = int8(fi25)
                                                var v60 uint = 87
                                                _ = v53
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                            v1 = append(v1, "hello")
                                            _ = v49
                                        }
                                        v61 := GenS2{F0: 42, F1: v3, F2: 128}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v62 int16 = 92
                                        var v63 uint64 = (v22 ^ v22)
                                        v62 -= v62
                                        fmt.Println(v22)
                                        switch "foo bar" {
                                        case "":
                                            v1 = append(v1, "")
                                            if len(v1) > 0 {
                                                v1[0] = ("a" + "")
                                            }
                                            var v64 int32 = -1
                                            _ = v64
                                        case "hello":
                                            v3 += (v3 + (v3 >> ((2) & 63)))
                                        case "\n\t":
                                            var v65 uint16 = 0
                                            _ = v65
                                        }
                                        if (!(((!((0 < 0) || true)) && ((v23 || v2) && ((-2147483648 >= 127) || ((v26 * -2.25) > (v26 - (v26 * -2.25)))))) || true)) {
                                            _ = len(v1)
                                            var v66 int16 = 853
                                            v67 := GenS3{F0: (v3 & (v3 % (((v3 * v3)) | 1))), F1: 686}
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                                _ = fi24
                            }
                        } else {
                            v68 := GenS4{F0: "foo bar", F1: 4294967295, F2: 128}
                            for fi69 := 0; fi69 < 14; fi69++ {
                                v3 += v3
                                _ = len(v1)
                                _ = fi69
                            }
                            _ = v68
                        }
                        v70 := GenS4{F0: "\n\t", F1: 0, F2: 2}
                        _ = v23
                        _ = v70
                    } else {
                        for fi71 := 0; fi71 < 12; fi71++ {
                            if v2 {
                                v22 += v22
                                v2 = v2
                                if true {
                                    v3 += (v3 >> ((2) & 63))
                                } else {
                                    v72 := GenS5{F0: ((4294967295 != 987) && ((-(fi71)) == fi71))}
                                    v1 = append(v1, ((("foo bar" + ("hello" + "a")) + "hello") + "a"))
                                    _ = v72
                                }
                            } else {
                                var v73 uint = 65535
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                                var v74 uint8 = 1
                                _ = v73
                                _ = v74
                            }
                            fmt.Println(v22, v2)
                            _ = fi71
                        }
                        v3 = v3
                        var v75 int16 = 796
                        _ = v75
                    }
                    _ = v22
                }
            case 0:
                for fi76 := 0; fi76 < 13; fi76++ {
                    v77 := GenS5{F0: v2}
                    switch (fi76 + 844) {
                    case 281:
                        v3 = v3
                    case 32768:
                        _ = len(v1)
                    default:
                        v3 += (^(v3))
                    }
                    if false {
                        fmt.Println(v3, v3, v3)
                        fmt.Println(fi76, v3)
                        v3 = v3
                    } else {
                        fmt.Println(v2, v2, v2)
                        var v78 float32 = (((1e10 / ((float32(1.5) / 1e10) / (-2.25 * (-2.25 * 0.1)))) / 1.0) - -2.25)
                        _ = v78
                    }
                    _ = fi76
                    _ = v77
                }
            case 32767:
                if v2 {
                    if v2 {
                        fmt.Println(v3, v2)
                        if v2 {
                            var v79 uint64 = 2
                            _ = len(v1)
                            v2 = false
                            _ = v79
                        } else {
                            if (!true) {
                                switch "" {
                                case "foo bar":
                                    v3 += v3
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                case "\n\t":
                                    _ = len(v1)
                                }
                                v3 += 1
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                            } else {
                                fmt.Println(v2, v2)
                            }
                            fmt.Println(v2, v2)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        switch "a" {
                        case "\n\t":
                            var v80 int32 = 1
                            _ = v80
                        case "foo bar":
                            var v81 uint16 = 2
                            _ = v81
                        case "hello":
                            fmt.Println(v3, v2)
                        }
                    }
                    for fi82 := 0; fi82 < 14; fi82++ {
                        for fi83 := 0; fi83 < 12; fi83++ {
                            fmt.Println(v2, fi82, v2)
                            var v84 uint16 = 1
                            _ = fi83
                            _ = v84
                        }
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        for fi85 := 0; fi85 < 15; fi85++ {
                            var v86 uint = 686
                            _ = fi85
                            _ = v86
                        }
                        _ = fi82
                    }
                    v3 += (v3 + v3)
                }
            }
        }
        if v2 {
            var v87 float64 = (((0.0 * ((0.0 - 1.5) + 1.5)) / 1e10) + -2.25)
            v3 += (-(v3))
            if v2 {
                v88 := GenS6{F0: 255, F1: 2, F2: 42}
                if true {
                    v89 := GenS7{F0: -1}
                    if false {
                        var v90 uint = 255
                        for fi91 := 0; fi91 < 17; fi91++ {
                            switch ("a" + "hello") {
                            case "foo bar":
                                switch (("\n\t" + "a") + "\n\t") {
                                case "\n\t":
                                    v87 += 3.14159
                                case "foo bar":
                                    v2 = (v2 && true)
                                default:
                                    if v2 {
                                        if (424 >= 2) {
                                            for fi92 := 0; fi92 < 17; fi92++ {
                                                if (-449 >= 456) {
                                                    fmt.Println(fi92)
                                                    fmt.Println(fi92, v2, v90)
                                                }
                                                fmt.Println(v90, v87)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = fi92
                                            }
                                            if (4294967295 <= 4294967295) {
                                                fmt.Println(v2, v90, v3)
                                                var v93 string = "foo bar"
                                                for fi94 := 0; fi94 < 16; fi94++ {
                                                    _ = len(v1)
                                                    var v95 int8 = 29
                                                    v93 = ((v93 + v93) + "")
                                                    _ = fi94
                                                    _ = v95
                                                }
                                                _ = v93
                                            }
                                        }
                                        switch fi91 {
                                        case 127:
                                            if (670 <= -538) {
                                                fmt.Println(v3, fi91)
                                                var v96 int = (fi91 + (fi91 + fi91))
                                                if (true && v2) {
                                                    for fi97 := 0; fi97 < 5; fi97++ {
                                                        if true {
                                                            switch -345 {
                                                            case -160:
                                                                var v98 uint16 = 2
                                                                v99 := GenS8{F0: v3}
                                                                if (4294967295 <= v96) {
                                                                    if false {
                                                                        if ((float64(1.5) + (0.1 * 1.5)) == v87) {
                                                                            var v100 int8 = (v3 | v3)
                                                                            v100 -= v100
                                                                            fmt.Println(fi97, v98, fi91)
                                                                            _ = v100
                                                                        } else {
                                                                            fmt.Println(v96)
                                                                            switch (fi91 ^ 9223372036854775807) {
                                                                            case 826:
                                                                                var v101 bool = v2
                                                                                _ = v101
                                                                            case 730:
                                                                                v102 := GenS6{F0: v98, F1: 0, F2: 42}
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                        var v103 uint32 = uint32(v98)
                                                                        if v2 {
                                                                            v104 := GenS9{F0: (v98 <= 42)}
                                                                            _ = v104
                                                                        } else {
                                                                            fmt.Println(fi97, v96)
                                                                        }
                                                                        _ = v103
                                                                    } else {
                                                                        switch (fi91 - 127) {
                                                                        case -9223372036854775808:
                                                                            var v105 bool = v2
                                                                            _ = v105
                                                                        case 32767:
                                                                            v98 = v98
                                                                        }
                                                                        fmt.Println(v90, v96)
                                                                        var v106 float32 = 1.5
                                                                        v3 -= (v3 * 1)
                                                                        switch -66 {
                                                                        case -647:
                                                                            v96 = fi91
                                                                            v107 := GenS10{F0: "hello", F1: 4294967295}
                                                                            var v108 string = "hello"
                                                                            if (!(v98 != v98)) {
                                                                                _ = len(v1)
                                                                                var v109 uint8 = 42
                                                                                fmt.Println(v108, fi97)
                                                                                _ = v109
                                                                            }
                                                                            if ((v98 >> ((v90) & 63)) >= (^(v98))) {
                                                                                v98 = v98
                                                                                v108 = (((v108 + v108) + v108) + v108)
                                                                            }
                                                                            _ = v107
                                                                            _ = v108
                                                                        default:
                                                                            v96 = 128
                                                                        }
                                                                        var v110 int = (-(fi91))
                                                                        v1 = append(v1, "foo bar")
                                                                        _ = v106
                                                                        _ = v110
                                                                    }
                                                                    if (v2 || (!false)) {
                                                                        if (!true) {
                                                                            var v111 int8 = 42
                                                                            var v112 uint = 18446744073709551615
                                                                            _ = v111
                                                                            _ = v112
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            if (-56 <= v3) {
                                                                                switch "" {
                                                                                case "hello":
                                                                                    switch (("" + (("hello" + "foo bar") + "")) + "foo bar") {
                                                                                    case "foo bar":
                                                                                        v113 := GenS9{F0: (!(!v2))}
                                                                                        _ = v113
                                                                                    }
                                                                                case "\n\t":
                                                                                    if ((^(v90)) >= (v90 / ((v90) | 1))) {
                                                                                        v114 := GenS9{F0: v2}
                                                                                        v115 := GenS11{F0: v90, F1: -1, F2: -1}
                                                                                        var v116 uint64 = 18446744073709551615
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                    } else {
                                                                                        var v117 int32 = -50
                                                                                        _ = v117
                                                                                    }
                                                                                default:
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        v3 += (v3 * (v3 & v3))
                                                                                    }
                                                                                }
                                                                                if v2 {
                                                                                    v2 = v2
                                                                                    var v118 uint32 = 2147483648
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    var v119 int32 = -93
                                                                                    _ = v118
                                                                                    _ = v119
                                                                                } else {
                                                                                    var v120 uint16 = (v98 ^ v98)
                                                                                    _ = v120
                                                                                }
                                                                                fmt.Println(v87, v98, v96)
                                                                            } else {
                                                                                _ = len(v1)
                                                                            }
                                                                        case "":
                                                                            if false {
                                                                                if ((^(v90)) != 18446744073709551615) {
                                                                                    var v121 float32 = ((float32(v87) - (3.14159 * 3.14159)) / 1.5)
                                                                                    _ = len(v1)
                                                                                    _ = v121
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v87, v3, v2)
                                                                    }
                                                                } else {
                                                                    fmt.Println(v87, v87)
                                                                    if v2 {
                                                                        switch fi91 {
                                                                        case 9223372036854775807:
                                                                            v98 -= 42
                                                                        case 855:
                                                                            var v122 int = v96
                                                                            _ = v122
                                                                        }
                                                                        v123 := GenS4{F0: "\n\t", F1: (v90 ^ 537), F2: 255}
                                                                        if false {
                                                                            var v124 float32 = 1.5
                                                                            _ = v124
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ""
                                                                        }
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            fmt.Println(v90, v90)
                                                                        case "hello":
                                                                            v98 = v98
                                                                        case "a":
                                                                            v3 = v3
                                                                        default:
                                                                            var v125 int16 = -1
                                                                            _ = v125
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ((("\n\t" + "a") + (("" + "hello") + ("a" + ("\n\t" + "a")))) + "\n\t")
                                                                        }
                                                                        _ = v123
                                                                    } else {
                                                                        _ = len(v1)
                                                                        v1 = append(v1, (("" + "a") + "hello"))
                                                                    }
                                                                }
                                                                _ = v98
                                                                _ = v99
                                                            case 4294967295:
                                                                switch (("\n\t" + "a") + "\n\t") {
                                                                case "a":
                                                                    fmt.Println(fi91, v90, fi97)
                                                                case "foo bar":
                                                                    var v126 float32 = 1.5
                                                                    _ = v126
                                                                case "\n\t":
                                                                    var v127 int16 = -1
                                                                    v128 := GenS12{F0: (v3 >> ((v90) & 63)), F1: "\n\t"}
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                            }
                                                            var v129 uint32 = 42
                                                            if len(v1) > 0 {
                                                                v1[0] = "\n\t"
                                                            }
                                                            _ = v129
                                                        } else {
                                                            v130 := GenS13{F0: 0, F1: -128, F2: v3}
                                                            if len(v1) > 0 {
                                                                v1[0] = ("foo bar" + "")
                                                            }
                                                            _ = v130
                                                        }
                                                        fmt.Println(fi97, v96)
                                                        var v131 uint16 = 969
                                                        if v2 {
                                                            switch 32767 {
                                                            case 31:
                                                                var v132 int64 = 1
                                                                _ = v132
                                                            case -1:
                                                                _ = len(v1)
                                                            case 755:
                                                                if len(v1) > 0 {
                                                                    v1[0] = "hello"
                                                                }
                                                            default:
                                                                if (((-2.25 - float32(((float64(0.1) - 1.5) - float64(0.1)))) != -2.25) && ((^(v131)) > v131)) {
                                                                    v131 = 255
                                                                }
                                                            }
                                                        } else {
                                                            v96 = 32768
                                                        }
                                                        _ = len(v1)
                                                        _ = fi97
                                                        _ = v131
                                                    }
                                                    v133 := GenS4{F0: "", F1: (v90 - (v90 ^ (v90 - 255))), F2: 43}
                                                    _ = v133
                                                }
                                                _ = v96
                                            } else {
                                                v134 := GenS8{F0: 127}
                                                _ = v134
                                            }
                                        case 9223372036854775807:
                                            for fi135 := 0; fi135 < 6; fi135++ {
                                                v136 := GenS14{F0: 171}
                                                if len(v1) > 0 {
                                                    v1[0] = (("\n\t" + ("hello" + "a")) + "\n\t")
                                                }
                                                if (false || v2) {
                                                    var v137 int = 4294967295
                                                    _ = v137
                                                } else {
                                                    var v138 int16 = 0
                                                    _ = v138
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    v2 = (!("foo bar" != "hello"))
                                                case "":
                                                    if ((fi135 & fi135) > (fi91 | fi135)) {
                                                        var v139 float64 = (v87 - ((1.5 / (v87 / ((float64(float32(-2.25)) - v87) + 1.5))) * (float64(-2.25) - (-2.25 - 0.1))))
                                                        var v140 bool = (!v2)
                                                        if (3.14159 <= 1e10) {
                                                            var v141 uint32 = 65535
                                                            v90 = v90
                                                            _ = v141
                                                        }
                                                        switch (^(fi91)) {
                                                        case 127:
                                                            if v2 {
                                                                if ((-(fi91)) != 852) {
                                                                    v90 += 4294967295
                                                                    switch "" {
                                                                    case "":
                                                                        v1 = append(v1, (("hello" + "foo bar") + "foo bar"))
                                                                        if (4294967295 == 927) {
                                                                            if v140 {
                                                                                var v142 uint32 = 177
                                                                                _ = v142
                                                                            } else {
                                                                                if (!(!(790 != 42))) {
                                                                                    var v143 uint8 = uint8(v90)
                                                                                    var v144 uint16 = 2
                                                                                    _ = v143
                                                                                    _ = v144
                                                                                } else {
                                                                                    fmt.Println(v2, fi91)
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = "\n\t"
                                                                                    }
                                                                                    fmt.Println(v140)
                                                                                }
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                switch (fi91 << ((v90) & 63)) {
                                                                                case 425:
                                                                                    v145 := GenS15{F0: -745, F1: 1, F2: -32768}
                                                                                    _ = v145
                                                                                }
                                                                            }
                                                                        }
                                                                        if true {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            v1 = append(v1, "\n\t")
                                                                        } else {
                                                                            v2 = (274 <= 9223372036854775807)
                                                                        }
                                                                        switch (fi91 >> ((v90) & 63)) {
                                                                        case 4294967295:
                                                                            _ = len(v1)
                                                                        case 42:
                                                                            if v2 {
                                                                                var v146 int16 = 1000
                                                                                switch ("foo bar" + ("\n\t" + "a")) {
                                                                                case "a":
                                                                                    var v147 int16 = -32768
                                                                                    _ = v147
                                                                                case "hello":
                                                                                    v1 = append(v1, "\n\t")
                                                                                case "foo bar":
                                                                                    v1 = append(v1, "foo bar")
                                                                                }
                                                                                _ = v146
                                                                            }
                                                                        }
                                                                    }
                                                                    if v2 {
                                                                        v87 = v139
                                                                        fmt.Println(v3)
                                                                        var v148 float64 = v139
                                                                        _ = v148
                                                                    } else {
                                                                        v87 -= ((v139 / v139) * (0.1 * v87))
                                                                    }
                                                                }
                                                                fmt.Println(v87, v87)
                                                            }
                                                        case -142:
                                                            v140 = (v139 <= -2.25)
                                                        default:
                                                            v90 = (v90 % ((255) | 1))
                                                        }
                                                        _ = v139
                                                        _ = v140
                                                    } else {
                                                        v2 = (4294967295 <= v90)
                                                        fmt.Println(v2)
                                                        switch ("\n\t" + "a") {
                                                        case "":
                                                            fmt.Println(fi135, fi135, fi135)
                                                        default:
                                                            fmt.Println(v90, v3)
                                                        }
                                                    }
                                                default:
                                                    var v149 float32 = 0.0
                                                    _ = v149
                                                }
                                                _ = fi135
                                                _ = v136
                                            }
                                        default:
                                            v150 := GenS14{F0: 388}
                                            _ = v150
                                        }
                                        v90 -= v90
                                    } else {
                                        fmt.Println(v2, v3)
                                    }
                                }
                            default:
                                for fi151 := 0; fi151 < 4; fi151++ {
                                    if v2 {
                                        v152 := GenS2{F0: 792, F1: 1, F2: 255}
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        fmt.Println(fi91, v3, v90)
                                        _ = v152
                                    } else {
                                        switch (-(fi91)) {
                                        case 1:
                                            fmt.Println(v2, v3)
                                        case 42:
                                            v153 := GenS16{F0: (-(fi91))}
                                            _ = v153
                                        case 436:
                                            v2 = (0 <= 1)
                                        default:
                                            switch "" {
                                            case "foo bar":
                                                var v154 uint64 = 2
                                                _ = v154
                                            case "hello":
                                                v2 = true
                                            case "\n\t":
                                                var v155 int = (fi91 * -2147483648)
                                                var v156 int8 = v3
                                                fmt.Println(v2)
                                                _ = v155
                                                _ = v156
                                            }
                                        }
                                        var v157 uint = (v90 % ((2) | 1))
                                        _ = v157
                                    }
                                    v90 = (^(v90))
                                    switch (fi91 << (((v90 - 512)) & 63)) {
                                    case 42:
                                        fmt.Println(v2, v3, v87)
                                    }
                                    _ = fi151
                                }
                            }
                            v1 = append(v1, "\n\t")
                            _ = fi91
                        }
                        _ = v90
                    } else {
                        v2 = (725 <= 9223372036854775807)
                        for fi158 := 0; fi158 < 2; fi158++ {
                            if (40 > 1) {
                                if (!((-(fi158)) >= fi158)) {
                                    if v2 {
                                        v2 = true
                                        for fi159 := 0; fi159 < 7; fi159++ {
                                            if (v3 <= -63) {
                                                var v160 uint = 0
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                                v161 := GenS17{F0: 806, F1: false, F2: 42}
                                                _ = v160
                                                _ = v161
                                            } else {
                                                fmt.Println(v87, fi158)
                                                v162 := GenS9{F0: false}
                                                var v163 string = ("" + ("\n\t" + ""))
                                                _ = v162
                                                _ = v163
                                            }
                                            switch fi158 {
                                            case -1:
                                                var v164 int64 = 101
                                                _ = v164
                                            case -146:
                                                v3 = 0
                                            }
                                            _ = fi159
                                        }
                                    } else {
                                        if (!(455 < 4294967295)) {
                                            v2 = true
                                            v165 := GenS8{F0: (v3 * v3)}
                                            for fi166 := 0; fi166 < 3; fi166++ {
                                                switch (fi166 % ((-451) | 1)) {
                                                case -152:
                                                    if len(v1) > 0 {
                                                        v1[0] = ""
                                                    }
                                                case 58:
                                                    var v167 bool = true
                                                    _ = v167
                                                case 0:
                                                    switch "\n\t" {
                                                    case "":
                                                        var v168 string = "hello"
                                                        var v169 uint64 = 2
                                                        _ = v168
                                                        _ = v169
                                                    }
                                                }
                                                v3 -= (v3 % (((v3 * (v3 - 0))) | 1))
                                                v170 := GenS9{F0: (255 > 18446744073709551615)}
                                                _ = fi166
                                                _ = v170
                                            }
                                            _ = v165
                                        }
                                        var v171 int8 = (^(v3))
                                        _ = v171
                                    }
                                    var v172 uint16 = 1
                                    _ = v172
                                } else {
                                    var v173 int32 = 42
                                    v174 := GenS18{F0: (fi158 / ((fi158) | 1)), F1: -1, F2: 42}
                                    v1 = append(v1, "hello")
                                    _ = v173
                                    _ = v174
                                }
                                switch 128 {
                                case 1:
                                    var v175 int16 = 767
                                    _ = v175
                                case 948:
                                    v176 := GenS3{F0: v3, F1: 255}
                                    _ = v176
                                case 127:
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v177 := GenS18{F0: 9223372036854775807, F1: -32768, F2: 40}
                                        if (1000 != -32768) {
                                            fmt.Println(v87)
                                            var v178 uint32 = 1
                                            _ = v178
                                        } else {
                                            v3 = v3
                                            var v179 uint16 = 377
                                            _ = v179
                                        }
                                        _ = v177
                                    } else {
                                        _ = len(v1)
                                        var v180 int = fi158
                                        fmt.Println(v2, v180)
                                        _ = v180
                                    }
                                default:
                                    if v2 {
                                        if len(v1) > 0 {
                                            v1[0] = ("a" + ("a" + ("a" + "a")))
                                        }
                                        for fi181 := 0; fi181 < 8; fi181++ {
                                            switch ("hello" + "") {
                                            case "a":
                                                var v182 bool = v2
                                                _ = v182
                                            default:
                                                var v183 float64 = 0.0
                                                _ = v183
                                            }
                                            var v184 string = "a"
                                            _ = fi181
                                            _ = v184
                                        }
                                    }
                                }
                            } else {
                                var v185 int64 = 9223372036854775807
                                var v186 uint32 = 953
                                if (("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))) <= "\n\t") {
                                    switch "hello" {
                                    case "hello":
                                        if v2 {
                                            v186 = (v186 | v186)
                                            fmt.Println(v2)
                                            fmt.Println(v186, fi158, v3)
                                        } else {
                                            fmt.Println(v87, fi158)
                                        }
                                    }
                                } else {
                                    for fi187 := 0; fi187 < 7; fi187++ {
                                        var v188 int = (fi187 / ((fi158) | 1))
                                        _ = fi187
                                        _ = v188
                                    }
                                }
                                _ = v185
                                _ = v186
                            }
                            switch "hello" {
                            case "\n\t":
                                _ = len(v1)
                            default:
                                for fi189 := 0; fi189 < 19; fi189++ {
                                    var v190 int16 = 882
                                    _ = fi189
                                    _ = v190
                                }
                            }
                            _ = fi158
                        }
                    }
                    _ = v89
                } else {
                    if (!true) {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                    } else {
                        for fi191 := 0; fi191 < 17; fi191++ {
                            if (255 > 18446744073709551615) {
                                for fi192 := 0; fi192 < 8; fi192++ {
                                    v2 = (0 < 32767)
                                    _ = fi192
                                }
                                fmt.Println(v87)
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            } else {
                                v3 = (v3 % (((v3 ^ v3)) | 1))
                            }
                            var v193 string = "a"
                            _ = fi191
                            _ = v193
                        }
                        var v194 float32 = (-2.25 + 3.14159)
                        _ = v194
                    }
                    v3 = (v3 | v3)
                }
                fmt.Println(v87, v3)
                _ = v88
            } else {
                v3 = -128
                v195 := GenS6{F0: 32767, F1: 4294967295, F2: -1}
                v1 = append(v1, "\n\t")
                _ = v195
            }
            _ = v87
        } else {
            switch 121 {
            case 130:
                switch 4294967295 {
                case 781:
                    v196 := GenS19{F0: 1}
                    _ = v196
                default:
                    v2 = v2
                }
            case 4294967295:
                if len(v1) > 0 {
                    v1[0] = ((("a" + "a") + "hello") + "\n\t")
                }
            case 42:
                if ((float32(((1.5 - ((0.1 * 0.0) * 1.5)) / float64((1.5 + ((0.1 * (0.1 - (1.5 / ((0.1 * 3.14159) / (3.14159 + 3.14159))))) * 3.14159))))) <= 0.0) || (v2 && false)) {
                    if (1 < 518) {
                        if false {
                            var v197 int16 = 727
                            _ = v197
                        }
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        if false {
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                        } else {
                            v3 = (v3 + (v3 - v3))
                            switch (("\n\t" + "\n\t") + "\n\t") {
                            case "foo bar":
                                if ("" != ("" + "\n\t")) {
                                    v3 = (v3 * (^(v3)))
                                }
                            case "\n\t":
                                var v198 uint = 147
                                _ = v198
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                    }
                    switch 234 {
                    case 128:
                        var v199 uint8 = 42
                        _ = v199
                    }
                } else {
                    if (v3 >= (v3 & v3)) {
                        v3 += 42
                        switch -218 {
                        case 226:
                            switch (("a" + (("hello" + "a") + "foo bar")) + ("\n\t" + "\n\t")) {
                            case "":
                                v1 = append(v1, "a")
                            case "hello":
                                switch "" {
                                case "foo bar":
                                    switch "\n\t" {
                                    case "":
                                        if (v2 || true) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v3 = 42
                                            default:
                                                v200 := GenS11{F0: 65535, F1: -32768, F2: -2147483648}
                                                _ = v200
                                            }
                                        } else {
                                            switch -1 {
                                            case 32767:
                                                var v201 uint16 = 868
                                                _ = v201
                                            case 32768:
                                                switch 128 {
                                                case 51:
                                                    if ("foo bar" == "a") {
                                                        var v202 int32 = 42
                                                        switch 42 {
                                                        case 32767:
                                                            if len(v1) > 0 {
                                                                v1[0] = "hello"
                                                            }
                                                        case 2147483647:
                                                            if (42 > -1) {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v203 := GenS7{F0: 42}
                                                                for fi204 := 0; fi204 < 12; fi204++ {
                                                                    if (42 <= 158) {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                    if v2 {
                                                                        for fi205 := 0; fi205 < 17; fi205++ {
                                                                            var v206 bool = (128 != 215)
                                                                            if true {
                                                                                v207 := GenS2{F0: 2, F1: (^(v3)), F2: 255}
                                                                                var v208 int32 = v202
                                                                                v206 = (-409 <= -996)
                                                                                v2 = v2
                                                                                _ = v207
                                                                                _ = v208
                                                                            } else {
                                                                                v209 := GenS17{F0: -1, F1: (v2 && true), F2: 976}
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    switch (fi204 << ((42) & 63)) {
                                                                                    case 128:
                                                                                        v210 := GenS20{F0: 1, F1: 42, F2: 1000}
                                                                                        _ = v210
                                                                                    case -767:
                                                                                        if true {
                                                                                            v211 := GenS3{F0: (v3 >> ((212) & 63)), F1: 32767}
                                                                                            var v212 int32 = -259
                                                                                            v213 := GenS21{F0: fi204}
                                                                                            _ = v211
                                                                                            _ = v212
                                                                                            _ = v213
                                                                                        } else {
                                                                                            if false {
                                                                                                var v214 uint32 = 851
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                v214 = v214
                                                                                                _ = v214
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v215 bool = (("hello" + ("foo bar" + "hello")) > "")
                                                                                        _ = v215
                                                                                    }
                                                                                case "\n\t":
                                                                                    v2 = v2
                                                                                default:
                                                                                    _ = len(v1)
                                                                                }
                                                                                var v216 uint = 255
                                                                                _ = v209
                                                                                _ = v216
                                                                            }
                                                                            _ = fi205
                                                                            _ = v206
                                                                        }
                                                                        v3 = (v3 / ((v3) | 1))
                                                                        v202 += -2147483648
                                                                    } else {
                                                                        if (v3 < (v3 * v3)) {
                                                                            v2 = (false || v2)
                                                                            v217 := GenS6{F0: 32767, F1: 171, F2: v202}
                                                                            var v218 uint64 = 4294967295
                                                                            _ = v217
                                                                            _ = v218
                                                                        } else {
                                                                            var v219 int32 = 868
                                                                            _ = v219
                                                                        }
                                                                        fmt.Println(fi204, v202)
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "foo bar"
                                                                        }
                                                                    }
                                                                    _ = fi204
                                                                }
                                                                _ = v203
                                                            }
                                                        case 0:
                                                            fmt.Println(v2)
                                                        default:
                                                            var v220 int = 1
                                                            _ = v220
                                                        }
                                                        if (2147483647 <= -9223372036854775808) {
                                                            var v221 uint64 = 2
                                                            _ = v221
                                                        } else {
                                                            for fi222 := 0; fi222 < 12; fi222++ {
                                                                if (2147483647 <= (v202 % (((v202 % ((0) | 1))) | 1))) {
                                                                    v202 += v202
                                                                } else {
                                                                    var v223 float32 = 1e10
                                                                    _ = v223
                                                                }
                                                                fmt.Println(v3, v202, fi222)
                                                                _ = fi222
                                                            }
                                                        }
                                                        _ = v202
                                                    } else {
                                                        if true {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        } else {
                                                            if (!(false && true)) {
                                                                v2 = (485 < 0)
                                                            }
                                                            var v224 float64 = 1.5
                                                            var v225 int64 = 42
                                                            _ = v224
                                                            _ = v225
                                                        }
                                                    }
                                                case 4294967295:
                                                    var v226 int8 = v3
                                                    _ = v226
                                                default:
                                                    switch 42 {
                                                    case 810:
                                                        v3 += 102
                                                    case 0:
                                                        v1 = append(v1, "a")
                                                    default:
                                                        for fi227 := 0; fi227 < 14; fi227++ {
                                                            var v228 uint64 = 4294967295
                                                            for fi229 := 0; fi229 < 13; fi229++ {
                                                                _ = len(v1)
                                                                _ = fi229
                                                            }
                                                            var v230 float32 = 0.1
                                                            _ = fi227
                                                            _ = v228
                                                            _ = v230
                                                        }
                                                    }
                                                }
                                            default:
                                                if (v2 && (18446744073709551615 < 0)) {
                                                    if ((0.0 / 0.1) >= ((0.0 - 3.14159) / 1.5)) {
                                                        var v231 float32 = 1e10
                                                        v2 = (9223372036854775807 == 4294967295)
                                                        if v2 {
                                                            for fi232 := 0; fi232 < 18; fi232++ {
                                                                switch (fi232 | (fi232 % ((fi232) | 1))) {
                                                                case 0:
                                                                    v231 = v231
                                                                }
                                                                _ = fi232
                                                            }
                                                        }
                                                        _ = v231
                                                    }
                                                    var v233 float32 = 1e10
                                                    var v234 uint64 = 4294967295
                                                    _ = v233
                                                    _ = v234
                                                } else {
                                                    for fi235 := 0; fi235 < 15; fi235++ {
                                                        switch ("" + "a") {
                                                        case "hello":
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v2 = v2
                                                            case "a":
                                                                if false {
                                                                    var v236 int8 = 100
                                                                    v237 := GenS14{F0: 443}
                                                                    var v238 int8 = v3
                                                                    _ = v236
                                                                    _ = v237
                                                                    _ = v238
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v239 := GenS22{F0: 65535, F1: fi235}
                                                            _ = v239
                                                        default:
                                                            var v240 float64 = -2.25
                                                            _ = v240
                                                        }
                                                        var v241 int32 = 1
                                                        fmt.Println(fi235, fi235, v241)
                                                        _ = fi235
                                                        _ = v241
                                                    }
                                                    for fi242 := 0; fi242 < 5; fi242++ {
                                                        v243 := GenS16{F0: fi242}
                                                        switch fi242 {
                                                        case 1:
                                                            if len(v1) > 0 {
                                                                v1[0] = (("hello" + "a") + "\n\t")
                                                            }
                                                        default:
                                                            if len(v1) > 0 {
                                                                v1[0] = "hello"
                                                            }
                                                        }
                                                        fmt.Println(fi242, fi242, fi242)
                                                        _ = fi242
                                                        _ = v243
                                                    }
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        _ = len(v1)
                                    }
                                case "":
                                    if true {
                                        _ = len(v1)
                                        var v244 int = -501
                                        _ = v244
                                    } else {
                                        v245 := GenS23{F0: 2147483648, F1: (!v2), F2: 9223372036854775807}
                                        _ = v245
                                    }
                                case "a":
                                    if (407 != -656) {
                                        fmt.Println(v2, v3)
                                        var v246 bool = v2
                                        _ = v246
                                    }
                                }
                            case "a":
                                v2 = (((("" + "") + "") > "") && (!v2))
                            }
                        case -247:
                            switch ("a" + "") {
                            case "hello":
                                var v247 int8 = -75
                                _ = v247
                            case "":
                                var v248 uint32 = 4294967295
                                _ = v248
                            case "foo bar":
                                v3 = (v3 & (v3 / ((1) | 1)))
                            }
                        case -2147483648:
                            v249 := GenS24{F0: -1, F1: 255}
                            _ = v249
                        default:
                            fmt.Println(v2, v3)
                        }
                    }
                    fmt.Println(v3, v3)
                    var v250 int = 0
                    _ = v250
                }
            }
            v2 = v2
            if (true && (true || (!(1.5 > -2.25)))) {
                var v251 int32 = 2147483647
                var v252 uint8 = 0
                _ = v251
                _ = v252
            } else {
                v2 = v2
            }
        }
        for fi253 := 0; fi253 < 14; fi253++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v254 := GenS25{F0: 1000}
            v3 += (v3 * (v3 << ((42) & 63)))
            _ = fi253
            _ = v254
        }
    }
    switch -773 {
    case -91:
        if v2 {
            var v255 uint8 = 128
            var v256 uint8 = v255
            switch "\n\t" {
            case "foo bar":
                v257 := GenS26{F0: 371, F1: 1}
                _ = v257
            default:
                if true {
                    if (116 > v3) {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        for fi258 := 0; fi258 < 1; fi258++ {
                            switch fi258 {
                            case 418:
                                var v259 int8 = (v3 ^ (v3 >> ((1) & 63)))
                                _ = v259
                            case 9223372036854775807:
                                for fi260 := 0; fi260 < 18; fi260++ {
                                    v261 := GenS27{F0: (255 < 255)}
                                    v255 -= (v255 & v256)
                                    if (!v2) {
                                        var v262 uint = 65535
                                        v2 = v2
                                        _ = v262
                                    } else {
                                        var v263 int8 = 100
                                        _ = v263
                                    }
                                    _ = fi260
                                    _ = v261
                                }
                            default:
                                for fi264 := 0; fi264 < 7; fi264++ {
                                    fmt.Println(v256, v255)
                                    _ = fi264
                                }
                            }
                            fmt.Println(v256)
                            _ = fi258
                        }
                        if v2 {
                            v1 = append(v1, "a")
                            switch 32767 {
                            case 2147483647:
                                var v265 string = "a"
                                _ = v265
                            case 32768:
                                v3 = (v3 << ((1) & 63))
                            }
                            if v2 {
                                v266 := GenS28{F0: 910, F1: "foo bar"}
                                _ = v266
                            } else {
                                v2 = (v2 && false)
                                if (v2 && (0 > 991)) {
                                    v255 += v255
                                    var v267 uint64 = 2
                                    for fi268 := 0; fi268 < 12; fi268++ {
                                        if (v2 || true) {
                                            if v2 {
                                                v269 := GenS2{F0: 1, F1: v3, F2: 239}
                                                var v270 int32 = 42
                                                _ = v269
                                                _ = v270
                                            } else {
                                                var v271 int8 = (^(v3))
                                                _ = v271
                                            }
                                            var v272 float64 = 1e10
                                            _ = v272
                                        } else {
                                            if (-2.25 <= float32(((3.14159 + 0.0) * (float64(1e10) / (((-2.25 - (1.5 / 1.0)) * 0.0) + 1.5))))) {
                                                switch "hello" {
                                                case "a":
                                                    fmt.Println(v2)
                                                }
                                                var v273 string = "foo bar"
                                                var v274 int16 = 1000
                                                _ = v273
                                                _ = v274
                                            }
                                        }
                                        fmt.Println(v256, v256, v3)
                                        switch fi268 {
                                        case 127:
                                            var v275 int32 = 42
                                            _ = v275
                                        case 9223372036854775807:
                                            v3 = v3
                                        case 162:
                                            v276 := GenS4{F0: "foo bar", F1: uint(v267), F2: (v255 % (((v256 * v255)) | 1))}
                                            _ = v276
                                        }
                                        _ = fi268
                                    }
                                    _ = v267
                                }
                                v277 := GenS29{F0: 42, F1: 1, F2: 2}
                                _ = v277
                            }
                        } else {
                            v2 = v2
                        }
                    } else {
                        for fi278 := 0; fi278 < 12; fi278++ {
                            _ = len(v1)
                            _ = fi278
                        }
                        for fi279 := 0; fi279 < 10; fi279++ {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = fi279
                        }
                    }
                    var v280 uint32 = 1
                    _ = v280
                } else {
                    v281 := GenS3{F0: (v3 + (v3 % ((v3) | 1))), F1: 1}
                    v3 = v3
                    if (!(((1 >= 1) || false) || v2)) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v1 = append(v1, "a")
                    }
                    _ = v281
                }
            }
            _ = v255
            _ = v256
        }
    default:
        if (42 < 0) {
            var v282 uint = 18446744073709551615
            var v283 float32 = 1.5
            _ = v282
            _ = v283
        }
    }
    switch 32768 {
    case -2147483648:
        v3 = v3
    case 42:
        fmt.Println(v2, v3)
    }
    var v284 uint8 = 2
    for fi285 := 0; fi285 < 1; fi285++ {
        var v286 uint8 = v284
        _ = fi285
        _ = v286
    }
    v287 := GenS30{F0: 2147483648, F1: (v284 ^ v284)}
    v284 = 42
    switch 32768 {
    case -539:
        switch -2147483648 {
        case 9223372036854775807:
            _ = len(v1)
        }
    case 2147483647:
        var v288 uint8 = 111
        _ = v288
    }
    switch "a" {
    case "":
        if v2 {
            v289 := GenS30{F0: 2147483647, F1: 2}
            var v290 int8 = v3
            _ = v289
            _ = v290
        } else {
            v2 = v2
            var v291 int8 = v3
            fmt.Println(v2, v3, v284)
            _ = v291
        }
    case "\n\t":
        if ("foo bar" > ("a" + ("a" + "\n\t"))) {
            if v2 {
                var v292 uint8 = (v284 / (((v284 << ((18446744073709551615) & 63))) | 1))
                for fi293 := 0; fi293 < 5; fi293++ {
                    fmt.Println(v284, v284, v2)
                    for fi294 := 0; fi294 < 17; fi294++ {
                        v295 := GenS16{F0: fi294}
                        v3 = v3
                        _ = fi294
                        _ = v295
                    }
                    _ = fi293
                }
                _ = v292
            } else {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                v1 = append(v1, "\n\t")
                if (804 == int(v284)) {
                    v296 := GenS28{F0: 65535, F1: "a"}
                    var v297 int8 = 0
                    _ = v296
                    _ = v297
                } else {
                    var v298 uint = 2
                    v3 = -128
                    _ = v298
                }
            }
            var v299 uint16 = 42
            for fi300 := 0; fi300 < 16; fi300++ {
                var v301 int64 = -1
                _ = fi300
                _ = v301
            }
            _ = v299
        }
    case "foo bar":
        v302 := GenS31{F0: 2147483647}
        _ = v302
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v284:", v284)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v284
    _ = v287
    _ = v3
    _ = v5
}
