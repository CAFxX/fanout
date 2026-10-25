package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int8; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint; F2 string }
type GenS6 struct { F0 int16; F1 int16; F2 int16 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint64; F1 uint8; F2 int }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 int64; F1 uint64 }
type GenS14 struct { F0 uint8; F1 int64 }
type GenS15 struct { F0 uint32; F1 uint32 }
type GenS16 struct { F0 uint8; F1 uint64; F2 bool }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int32; F1 int; F2 int16 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint16, p1 uint32) int16 {
    p0 = p0
    p0 = (p0 | p0)
    v4 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = v4
    return int16(p0)
}

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]bool)
    v3 := []int{}
    if false {
        switch "a" {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        case "\n\t":
            for fi5 := 0; fi5 < 10; fi5++ {
                var v6 int32 = 0
                _ = fi5
                _ = v6
            }
        case "":
            v7 := GenS2{F0: -317}
            _ = v7
        default:
            v8 := GenS1{F0: false}
            _ = v8
        }
        var v9 float32 = 1.5
        _ = v9
    } else {
        var v10 int64 = 627
        if (811 == -2147483648) {
            fmt.Println(v10)
            var v11 int16 = 1
            for fi12 := 0; fi12 < 1; fi12++ {
                v11 = genFn0(391, 2)
                switch (^(fi12)) {
                case -564:
                    for fi13 := 0; fi13 < 19; fi13++ {
                        if (743 == 307) {
                            v14 := GenS3{F0: 18446744073709551615, F1: 100, F2: 127}
                            _ = v14
                        } else {
                            v11 += v11
                            fmt.Println(v11)
                        }
                        if false {
                            if true {
                                fmt.Println(v10)
                                fmt.Println(v11)
                                var v15 int32 = 0
                                fmt.Println(v15, v11)
                                switch "a" {
                                case "a":
                                    if true {
                                        var v16 uint = 403
                                        v17 := GenS3{F0: (v16 * 2), F1: 42, F2: 140}
                                        v11 = v11
                                        _ = v16
                                        _ = v17
                                    } else {
                                        v11 += v11
                                    }
                                default:
                                    fmt.Println(fi13)
                                }
                                if (v15 != v15) {
                                    var v18 uint16 = 0
                                    _ = v18
                                }
                                var v19 int64 = (v10 % ((-825) | 1))
                                var v20 uint = 1
                                _ = v15
                                _ = v19
                                _ = v20
                            } else {
                                var v21 bool = true
                                var v22 bool = v21
                                if v21 {
                                    v10 = (v10 * v10)
                                    var v23 int64 = (v10 & (-(v10)))
                                    _ = v23
                                } else {
                                    var v24 uint32 = 0
                                    v21 = (0 == 399)
                                    var v25 float64 = 1e10
                                    _ = v24
                                    _ = v25
                                }
                                _ = v21
                                _ = v22
                            }
                        }
                        switch "" {
                        case "hello":
                            if false {
                                switch fi12 {
                                case 42:
                                    v26 := GenS3{F0: 0, F1: -1, F2: 127}
                                    _ = v26
                                default:
                                    var v27 float64 = -2.25
                                    _ = v27
                                }
                            }
                        case "a":
                            v10 += 112
                        default:
                            if true {
                                switch fi12 {
                                case 128:
                                    v11 = (v11 * (-(v11)))
                                case 0:
                                    fmt.Println(fi13, v11)
                                default:
                                    v28 := GenS3{F0: 18446744073709551615, F1: -88, F2: 1}
                                    _ = v28
                                }
                                v11 = 1000
                            } else {
                                var v29 int8 = -110
                                _ = v29
                            }
                        }
                        if true {
                            if (!true) {
                                v11 = v11
                            }
                        } else {
                            v30 := GenS3{F0: 0, F1: 42, F2: 128}
                            v31 := GenS4{F0: fi12}
                            _ = v30
                            _ = v31
                        }
                        var v32 uint16 = 2
                        _ = fi13
                        _ = v32
                    }
                default:
                    var v33 uint16 = 32767
                    _ = v33
                }
                v10 -= (v10 & v10)
                _ = fi12
            }
            _ = v11
        } else {
            v10 = v10
            fmt.Println(v10)
        }
        switch 42 {
        case 128:
            v34 := GenS2{F0: -1}
            _ = v34
        case 32767:
            fmt.Println(v10)
        }
        _ = v10
    }
    if true {
        v35 := v2[42]
        _ = v35
    } else {
        var v36 string = ("foo bar" + "\n\t")
        _ = v36
    }
    var v37 int32 = 1
    switch (("foo bar" + "a") + ("\n\t" + "hello")) {
    case "":
        for fi38 := 0; fi38 < 12; fi38++ {
            for fi39 := 0; fi39 < 13; fi39++ {
                if (false && false) {
                    fmt.Println(v37, fi38, fi38)
                } else {
                    v40 := GenS5{F0: 42, F1: 65535, F2: "a"}
                    var v41 int = 32768
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v40
                    _ = v41
                }
                _ = fi39
            }
            v37 += v37
            switch "" {
            case "a":
                v37 = v37
            }
            _ = fi38
        }
    case "a":
        var v42 uint = 65535
        _ = v42
    case "foo bar":
        if false {
            if ((302 > 42) || (32767 >= -144)) {
                fmt.Println(v37)
                if (0 <= -1) {
                    switch 781 {
                    case 127:
                        delete(v2, 965)
                    case 720:
                        v37 += v37
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            v43 := v2[958]
                            _ = v43
                        case "foo bar":
                            var v44 uint16 = 255
                            _ = v44
                        case "":
                            v45 := GenS1{F0: (false || (true || ((-1 > 1) || true)))}
                            _ = v45
                        }
                    }
                    for fi46 := 0; fi46 < 6; fi46++ {
                        fmt.Println(fi46, v37)
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                switch "foo bar" {
                                case "a":
                                    if (!false) {
                                        var v47 string = "\n\t"
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                        v1 = append(v1, (((v47 + (v47 + "")) + (v47 + ((((v47 + "hello") + v47) + (v47 + v47)) + v47))) + ""))
                                        _ = v47
                                    }
                                case "":
                                    switch fi46 {
                                    case 32768:
                                        var v48 uint64 = 608
                                        _ = v48
                                    default:
                                        if true {
                                            fmt.Println(fi46)
                                            v37 -= v37
                                            for fi49 := 0; fi49 < 7; fi49++ {
                                                var v50 uint8 = 128
                                                if true {
                                                    switch 127 {
                                                    case 433:
                                                        fmt.Println(v50, v37, fi46)
                                                    default:
                                                        var v51 int32 = -850
                                                        _ = v51
                                                    }
                                                }
                                                var v52 uint32 = 2
                                                _ = fi49
                                                _ = v50
                                                _ = v52
                                            }
                                        } else {
                                            var v53 uint16 = 315
                                            _ = v53
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(v37)
                                default:
                                    v37 -= v37
                                }
                            case "a":
                                var v54 int32 = -1
                                _ = v54
                            case "hello":
                                v37 = (v37 / ((v37) | 1))
                            }
                        } else {
                            v55 := GenS1{F0: true}
                            for fi56 := 0; fi56 < 12; fi56++ {
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                v37 = v37
                                fmt.Println(v37)
                                v1 = append(v1, "hello")
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    if (true && (700 < 1)) {
                                        var v57 uint64 = 18446744073709551615
                                        if true {
                                            switch (fi46 * fi46) {
                                            case 657:
                                                fmt.Println(v57, v37)
                                            case 4294967295:
                                                fmt.Println(v37)
                                            case -2147483648:
                                                if true {
                                                    var v58 uint8 = 1
                                                    var v59 int8 = 100
                                                    var v60 float64 = 1e10
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    if (42 < 0) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v61 uint8 = 2
                                                        _ = v61
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = ((("\n\t" + "") + "hello") + "\n\t")
                                                    }
                                                    v62 := GenS5{F0: (v57 & v57), F1: 698, F2: "a"}
                                                    _ = v62
                                                }
                                            }
                                            switch ((("\n\t" + "\n\t") + "\n\t") + (("a" + (("\n\t" + ((("hello" + "hello") + "\n\t") + ("\n\t" + ""))) + "foo bar")) + "a")) {
                                            case "hello":
                                                if false {
                                                    _ = len(v1)
                                                    v63 := GenS6{F0: genFn0(969, uint32(v57)), F1: -38, F2: 1}
                                                    v64 := GenS1{F0: false}
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    var v65 int16 = genFn0(865, 2147483648)
                                                    _ = v65
                                                }
                                            case "\n\t":
                                                var v66 int = fi56
                                                _ = v66
                                            case "":
                                                fmt.Println(v37)
                                                var v67 int16 = 42
                                                var v68 int64 = 0
                                                if (int64(v57) > (v68 * 42)) {
                                                    v57 += (v57 | 294)
                                                    var v69 uint = 18446744073709551615
                                                    _ = v69
                                                } else {
                                                    _ = len(v3)
                                                    delete(v2, 42)
                                                }
                                                if ((v57 - 947) > 300) {
                                                    var v70 int64 = (v68 | (v68 / ((v68) | 1)))
                                                    _ = v70
                                                } else {
                                                    fmt.Println(v57)
                                                    delete(v2, 2)
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if false {
                                                            if (true && (-2.25 <= 0.1)) {
                                                                fmt.Println(fi46, v67, v68)
                                                                var v71 bool = (1 > 42)
                                                                _ = v71
                                                            } else {
                                                                var v72 int = (fi46 / (((fi56 ^ 4294967295)) | 1))
                                                                var v73 uint8 = 129
                                                                v57 += (v57 - (v57 / (((^(v57))) | 1)))
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        } else {
                                                            var v74 uint8 = 181
                                                            _ = v74
                                                        }
                                                    case "a":
                                                        v37 = v37
                                                    }
                                                }
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                        var v75 int16 = 1000
                                        _ = v57
                                        _ = v75
                                    }
                                    v76 := GenS6{F0: 0, F1: -871, F2: 1000}
                                    _ = v76
                                } else {
                                    v37 -= v37
                                    var v77 uint = 214
                                    _ = v77
                                }
                                _ = fi56
                            }
                            if ((0.0 / ((1e10 - 0.1) + float64(((-2.25 * 0.0) + (((1.5 * ((0.0 - (1.5 - ((1.5 * (1.5 / 1.0)) + -2.25))) * 3.14159)) - 0.1) * 1e10))))) >= (float64((((-2.25 / 1.5) / 0.1) * (1e10 - 1.5))) - 3.14159)) {
                                for fi78 := 0; fi78 < 11; fi78++ {
                                    v79 := GenS3{F0: 4294967295, F1: 42, F2: 0}
                                    v80 := GenS3{F0: 65535, F1: 42, F2: 26}
                                    switch "\n\t" {
                                    case "\n\t":
                                        if true {
                                            v81 := GenS3{F0: 1, F1: -30, F2: 255}
                                            var v82 bool = true
                                            if v82 {
                                                v37 = v37
                                                fmt.Println(fi46, fi46, v37)
                                                fmt.Println(v37, v37, v82)
                                                var v83 int16 = -32768
                                                v83 = 132
                                                _ = v83
                                            }
                                            _ = v81
                                            _ = v82
                                        }
                                    case "":
                                        v84 := GenS7{F0: (^(v37)), F1: ("foo bar" + "a")}
                                        _ = v84
                                    default:
                                        v85 := GenS5{F0: 293, F1: 80, F2: "hello"}
                                        _ = v85
                                    }
                                    _ = fi78
                                    _ = v79
                                    _ = v80
                                }
                                for fi86 := 0; fi86 < 17; fi86++ {
                                    var v87 uint32 = 596
                                    _ = fi86
                                    _ = v87
                                }
                            } else {
                                for fi88 := 0; fi88 < 14; fi88++ {
                                    v37 = 1
                                    v37 -= (v37 - v37)
                                    v37 += v37
                                    _ = fi88
                                }
                            }
                            _ = v55
                        }
                        _ = fi46
                    }
                } else {
                    v37 = v37
                    v3 = append(v3, 32767)
                    if true {
                        fmt.Println(v37)
                        fmt.Println(v37)
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                    }
                }
            } else {
                if (-59 >= 0) {
                    var v89 string = ""
                    _ = v89
                }
                for fi90 := 0; fi90 < 1; fi90++ {
                    switch "\n\t" {
                    case "a":
                        if (2147483648 <= 561) {
                            if (2 > 980) {
                                var v91 uint8 = 2
                                switch (fi90 | fi90) {
                                case -471:
                                    var v92 int32 = 42
                                    _ = v92
                                case 95:
                                    v37 -= -2147483648
                                }
                                for fi93 := 0; fi93 < 18; fi93++ {
                                    if ((v91 >> ((960) & 63)) != 219) {
                                        if (0 <= 42) {
                                            v94 := GenS3{F0: 0, F1: 100, F2: 255}
                                            if true {
                                                v37 = -721
                                                fmt.Println(fi93)
                                                v2[281] = (float32(((0.1 * (1.5 / 1e10)) / float64(1.5))) < 0.1)
                                                var v95 float32 = 1e10
                                                v95 = v95
                                                v37 -= (v37 & -1)
                                                if (-2.25 != ((1e10 * 1.5) - float64(v95))) {
                                                    if true {
                                                        var v96 float64 = 3.14159
                                                        _ = v96
                                                    } else {
                                                        if (-79 >= -1) {
                                                            v95 += 1e10
                                                            if false {
                                                                switch (fi93 * fi93) {
                                                                case -2147483648:
                                                                    var v97 bool = (!true)
                                                                    _ = v97
                                                                case -1:
                                                                    _ = len(v2)
                                                                default:
                                                                    var v98 int64 = -9223372036854775808
                                                                    _ = v98
                                                                }
                                                            } else {
                                                                var v99 uint = 602
                                                                var v100 int16 = 1
                                                                v37 += v37
                                                                if (true || false) {
                                                                    v100 = genFn0(42, 29)
                                                                }
                                                                _ = v100
                                                                _ = v99
                                                            }
                                                        } else {
                                                            switch 1 {
                                                            case -458:
                                                                var v101 float64 = -2.25
                                                                _ = v101
                                                            case 9223372036854775807:
                                                                v37 += int32(v91)
                                                            case -115:
                                                                var v102 uint16 = 974
                                                                _ = v102
                                                            }
                                                            if true {
                                                                var v103 uint32 = 293
                                                                var v104 uint16 = 255
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                v37 = (v37 * (v37 % (((v37 + -1)) | 1)))
                                                            }
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi93 - int(v37))
                                                        }
                                                        delete(v2, 78)
                                                    }
                                                } else {
                                                    if (true && true) {
                                                        v37 -= v37
                                                    } else {
                                                        v105 := GenS1{F0: true}
                                                        var v106 float32 = v95
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                    fmt.Println(fi93, fi93)
                                                }
                                                v37 += (v37 / (((v37 + v37)) | 1))
                                                _ = v95
                                            } else {
                                                fmt.Println(v91, v91)
                                                v91 += (v91 & v91)
                                                v107 := GenS8{F0: 112}
                                                _ = v107
                                            }
                                            _ = v94
                                        }
                                        v91 = 0
                                    }
                                    fmt.Println(fi93)
                                    if true {
                                        v37 += (v37 | v37)
                                        if (!false) {
                                            v108 := v2[2]
                                            var v109 int16 = -513
                                            _ = v108
                                            _ = v109
                                        }
                                        delete(v2, 606)
                                    }
                                    _ = fi93
                                }
                                _ = v91
                            } else {
                                v37 = v37
                                for fi110 := 0; fi110 < 10; fi110++ {
                                    if true {
                                        var v111 bool = (65535 <= 42)
                                        v3 = append(v3, (fi110 ^ 127))
                                        _ = v111
                                    }
                                    v37 = v37
                                    _ = fi110
                                }
                                for fi112 := 0; fi112 < 15; fi112++ {
                                    _ = len(v2)
                                    v113 := GenS9{F0: 864, F1: v37}
                                    v37 += v37
                                    _ = fi112
                                    _ = v113
                                }
                            }
                        }
                    case "hello":
                        for fi114 := 0; fi114 < 10; fi114++ {
                            delete(v2, 0)
                            var v115 bool = (18446744073709551615 != 25)
                            _ = fi114
                            _ = v115
                        }
                    }
                    _ = fi90
                }
            }
            for fi116 := 0; fi116 < 20; fi116++ {
                delete(v2, 2147483648)
                _ = fi116
            }
        }
    default:
        v117 := GenS7{F0: v37, F1: ("a" + "")}
        _ = v117
    }
    fmt.Println(v37)
    for fi118 := 0; fi118 < 20; fi118++ {
        switch fi118 {
        case -208:
            var v119 uint32 = 1
            _ = v119
        }
        v120 := GenS8{F0: -50}
        v121 := GenS10{F0: 2, F1: 127, F2: 32768}
        _ = fi118
        _ = v120
        _ = v121
    }
    _ = len(v1)
    if (-12 > -1) {
        if (!false) {
            v37 = v37
            v122 := GenS6{F0: 42, F1: genFn0(0, 1), F2: 32767}
            v123 := GenS7{F0: (v37 / (((v37 - -1)) | 1)), F1: "a"}
            _ = v122
            _ = v123
        }
        v37 += v37
        v124 := GenS11{F0: (v37 | v37), F1: 100}
        _ = v124
    }
    var v125 int = -76
    switch ("hello" + ("foo bar" + ("\n\t" + "hello"))) {
    case "foo bar":
        v125 -= v125
    case "hello":
        delete(v2, 42)
    case "a":
        v126 := GenS12{F0: 724, F1: 127}
        _ = v126
    }
    _ = len(v3)
    switch ("hello" + (("foo bar" + ("foo bar" + "a")) + ("foo bar" + (("foo bar" + "a") + "foo bar")))) {
    case "\n\t":
        for fi127 := 0; fi127 < 15; fi127++ {
            switch 2147483647 {
            case -2147483648:
                v125 = (fi127 % ((-1) | 1))
            case 550:
                v37 -= -1
            case 554:
                v128 := GenS7{F0: v37, F1: (("a" + (("foo bar" + "a") + "a")) + "a")}
                _ = v128
            default:
                v3 = append(v3, -9223372036854775808)
            }
            if true {
                for fi129 := 0; fi129 < 15; fi129++ {
                    var v130 int16 = 0
                    if (1.5 <= 3.14159) {
                        var v131 string = "hello"
                        v131 = (v131 + "foo bar")
                        if true {
                            if (2 < 180) {
                                v132 := GenS13{F0: -1, F1: 2}
                                _ = v132
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v130 = (v130 / (((v130 << ((274) & 63))) | 1))
                            }
                        } else {
                            if ((550 >= 42) || false) {
                                v133 := GenS7{F0: v37, F1: v131}
                                var v134 uint32 = 65535
                                v135 := GenS6{F0: v130, F1: v130, F2: (v130 ^ (v130 - v130))}
                                _ = v133
                                _ = v134
                                _ = v135
                            }
                            if false {
                                v130 = (v130 - (v130 - v130))
                                if true {
                                    if true {
                                        fmt.Println(v130, v125)
                                        v3 = append(v3, v125)
                                        v131 = v131
                                    }
                                    var v136 float64 = (3.14159 - 1e10)
                                    v125 = -339
                                    v37 += v37
                                    _ = len(v3)
                                    _ = v136
                                } else {
                                    var v137 uint32 = 264
                                    _ = v137
                                }
                                v37 = v37
                            }
                        }
                        var v138 float64 = float64(1e10)
                        v131 = v131
                        switch (fi129 << ((612) & 63)) {
                        case -2147483648:
                            v138 += (v138 - v138)
                        case 32767:
                            v139 := GenS8{F0: 127}
                            _ = v139
                        case 447:
                            if true {
                                v138 += (3.14159 * v138)
                                if (("" != v131) || true) {
                                    fmt.Println(v131, fi129)
                                    switch (("a" + (("foo bar" + v131) + v131)) + (("" + v131) + v131)) {
                                    case "foo bar":
                                        var v140 int8 = -128
                                        _ = v140
                                    }
                                } else {
                                    v37 = 0
                                    fmt.Println(v125, v138, v131)
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                }
                                v141 := GenS11{F0: (v37 / (((v37 + 1)) | 1)), F1: -1}
                                _ = v141
                            } else {
                                var v142 int32 = v37
                                v138 = (v138 / (((v138 / 0.1) * -2.25) * ((v138 / v138) - float64(0.1))))
                                if (218 <= 149) {
                                    var v143 int = 42
                                    fmt.Println(v143)
                                    _ = v143
                                }
                                var v144 bool = true
                                switch v131 {
                                case "\n\t":
                                    switch (("\n\t" + "") + (v131 + ((("a" + ("foo bar" + "a")) + "") + ""))) {
                                    case "foo bar":
                                        switch "" {
                                        case "\n\t":
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v131, v131, v131)
                                            case "hello":
                                                v145 := GenS10{F0: 2, F1: 167, F2: v125}
                                                _ = v145
                                            }
                                        case "foo bar":
                                            var v146 float32 = ((-2.25 + 0.1) / ((-2.25 - 3.14159) + 0.0))
                                            _ = v146
                                        default:
                                            if (2 > 793) {
                                                switch fi127 {
                                                case -1:
                                                    switch (^(fi127)) {
                                                    case 42:
                                                        fmt.Println(v138)
                                                    case -9223372036854775808:
                                                        v125 += (v125 * fi127)
                                                        var v147 uint64 = 1
                                                        _ = v147
                                                    default:
                                                        fmt.Println(v130, fi129, v131)
                                                    }
                                                case 2147483647:
                                                    v148 := v2[2147483648]
                                                    _ = v148
                                                default:
                                                    delete(v2, 1)
                                                }
                                                var v149 int = (fi129 / ((v125) | 1))
                                                v150 := GenS14{F0: 227, F1: 0}
                                                _ = v149
                                                _ = v150
                                            }
                                        }
                                    default:
                                        var v151 uint = 255
                                        _ = v151
                                    }
                                case "hello":
                                    var v152 uint64 = 2
                                    _ = v152
                                }
                                v142 -= (-(v142))
                                _ = v142
                                _ = v144
                            }
                        }
                        v153 := GenS15{F0: 42, F1: 42}
                        _ = v131
                        _ = v138
                        _ = v153
                    }
                    _ = fi129
                    _ = v130
                }
                v154 := GenS16{F0: 1, F1: 4294967295, F2: false}
                _ = v154
            }
            v37 = -49
            _ = fi127
        }
    case "":
        v155 := GenS12{F0: 42, F1: (v125 - 4294967295)}
        _ = v155
    case "foo bar":
        v156 := v2[2]
        _ = v156
    }
    switch "\n\t" {
    case "hello":
        v157 := GenS17{F0: 879}
        _ = v157
    case "a":
        v1 = append(v1, "")
    case "":
        if true {
            v158 := GenS2{F0: 1}
            fmt.Println(v125, v37)
            _ = v158
        } else {
            var v159 string = "foo bar"
            var v160 int8 = 42
            for fi161 := 0; fi161 < 1; fi161++ {
                if (0 >= 2) {
                    for fi162 := 0; fi162 < 17; fi162++ {
                        switch (fi162 / (((fi161 << ((4294967295) & 63))) | 1)) {
                        case -2147483648:
                            if true {
                                switch v159 {
                                case "hello":
                                    var v163 uint64 = 18446744073709551615
                                    _ = v163
                                case "a":
                                    if true {
                                        var v164 int32 = 1
                                        fmt.Println(v160)
                                        v37 = (v37 >> ((42) & 63))
                                        _ = v164
                                    }
                                }
                                var v165 uint64 = 383
                                v166 := v2[65535]
                                fmt.Println(v125, v37)
                                _ = v165
                                _ = v166
                            } else {
                                fmt.Println(v37)
                            }
                        case 4294967295:
                            v37 += v37
                        case 887:
                            var v167 bool = (1 != 978)
                            _ = len(v2)
                            _ = v167
                        default:
                            var v168 int16 = -1
                            _ = v168
                        }
                        var v169 uint32 = 0
                        fmt.Println(v169)
                        var v170 float64 = 1e10
                        var v171 int16 = 970
                        fmt.Println(v37)
                        fmt.Println(v37)
                        switch "" {
                        case "\n\t":
                            if (32767 >= 42) {
                                if false {
                                    v171 += v171
                                }
                                switch (v125 << ((2) & 63)) {
                                case 372:
                                    var v172 int16 = genFn0(130, (^(v169)))
                                    _ = v172
                                case -212:
                                    if false {
                                        if false {
                                            if (42 == 127) {
                                                v173 := GenS18{F0: 361, F1: fi161, F2: (v171 ^ (^(v171)))}
                                                _ = v173
                                            } else {
                                                var v174 int = (v125 | fi161)
                                                _ = v174
                                            }
                                            v175 := GenS19{F0: 128}
                                            _ = v175
                                        } else {
                                            var v176 float64 = (v170 + (v170 + -2.25))
                                            _ = v176
                                        }
                                        v177 := GenS20{F0: v169, F1: 42}
                                        var v178 int32 = v37
                                        _ = v177
                                        _ = v178
                                    } else {
                                        _ = len(v3)
                                    }
                                }
                            }
                        default:
                            if ((!true) || false) {
                                var v179 int32 = 0
                                v180 := GenS16{F0: 132, F1: 2, F2: true}
                                _ = v179
                                _ = v180
                            } else {
                                v171 = v171
                                delete(v2, 4294967295)
                                v181 := v2[v169]
                                _ = v181
                            }
                        }
                        _ = fi162
                        _ = v169
                        _ = v170
                        _ = v171
                    }
                    for fi182 := 0; fi182 < 8; fi182++ {
                        var v183 uint64 = 2
                        var v184 int = 2147483647
                        v37 = v37
                        v185 := v2[119]
                        switch "" {
                        case "hello":
                            switch "foo bar" {
                            case "foo bar":
                                var v186 int8 = (v160 & (v160 >> ((65535) & 63)))
                                _ = v186
                            }
                        case "":
                            fmt.Println(v184)
                        case "\n\t":
                            v183 = (v183 ^ v183)
                        }
                        _ = fi182
                        _ = v183
                        _ = v184
                        _ = v185
                    }
                    v37 = 981
                } else {
                    _ = len(v1)
                    fmt.Println(v160, v125, fi161)
                    var v187 float32 = (-2.25 / float32(1.5))
                    _ = v187
                }
                fmt.Println(v160, v160)
                var v188 uint = 0
                _ = fi161
                _ = v188
            }
            _ = v159
            _ = v160
        }
    }
    fmt.Println(v37, v125)
    v37 = 131
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v125:", v125)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v125
    _ = v2
    _ = v3
    _ = v37
}
