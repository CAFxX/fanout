package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int8; F1 uint32; F2 int32 }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64; F1 uint64; F2 int }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string; F1 int; F2 uint32 }
type GenS15 struct { F0 int16; F1 int; F2 int32 }
type GenS16 struct { F0 bool }
func genFn0(p0 int16, p1 string) int64 {
    if (-1 > -799) {
        for fi4 := 0; fi4 < 16; fi4++ {
            v5 := GenS1{F0: 1, F1: false, F2: 255}
            var v6 int16 = 95
            _ = fi4
            _ = v5
            _ = v6
        }
        p1 = p1
    } else {
        v7 := GenS1{F0: 9223372036854775807, F1: true, F2: 255}
        _ = v7
    }
    v8 := GenS1{F0: 9223372036854775807, F1: true, F2: 127}
    for fi9 := 0; fi9 < 7; fi9++ {
        switch "" {
        case "a":
            for fi10 := 0; fi10 < 5; fi10++ {
                p1 = p1
                p0 += (p0 % (((p0 << ((1) & 63))) | 1))
                var v11 uint8 = 255
                _ = fi10
                _ = v11
            }
        case "hello":
            v12 := GenS1{F0: 9223372036854775807, F1: (0 == 0), F2: 146}
            _ = v12
        case "foo bar":
            p1 = (p1 + p1)
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 9223372036854775807
}
func genFn1() bool {
    var v13 float32 = 0.0
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := []int{}
            var v17 int64 = genFn0(0, ("foo bar" + "foo bar"))
            switch 32767 {
            case 127:
                if true {
                    v18 := GenS2{F0: 2147483647, F1: (v17 < (v17 ^ 1)), F2: 682}
                    _ = v18
                } else {
                    if (1.5 == 0.0) {
                        v13 = v13
                        var v19 int = (fi14 * (fi14 * (-(fi14))))
                        if len(v16) > 0 {
                            v16[0] = 32768
                        }
                        _ = v19
                    } else {
                        switch fi15 {
                        case 127:
                            v20 := GenS1{F0: (v17 | -113), F1: false, F2: 110}
                            _ = v20
                        case 0:
                            var v21 uint64 = 1
                            _ = v21
                        case -224:
                            switch "\n\t" {
                            case "":
                                v22 := GenS3{F0: (!false), F1: 255}
                                _ = v22
                            default:
                                v13 -= v13
                            }
                        }
                        var v23 int8 = 0
                        _ = v23
                    }
                    v17 = (v17 ^ v17)
                }
            default:
                switch (fi15 * (fi15 % (((fi14 & fi14)) | 1))) {
                case -584:
                    if (("a" + "a") >= "hello") {
                        v24 := GenS1{F0: v17, F1: true, F2: 2}
                        v16 = append(v16, 127)
                        v16 = append(v16, fi15)
                        var v25 uint8 = 255
                        if (false && (!true)) {
                            var v26 uint16 = 32767
                            v27 := GenS3{F0: (0.0 > 0.1), F1: v26}
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                        _ = v25
                    }
                default:
                    if (!(127 <= 1)) {
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(fi15, v13, fi15)
                        case "a":
                            var v28 float64 = float64((1.5 - 3.14159))
                            _ = v28
                        default:
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        }
                        if (!true) {
                            v13 = ((v13 + v13) - v13)
                            v29 := GenS2{F0: -1, F1: ((!false) || false), F2: 4294967295}
                            switch fi14 {
                            case -9223372036854775808:
                                var v30 uint32 = 776
                                _ = v30
                            }
                            _ = v29
                        } else {
                            var v31 float64 = float64(float32(float64(((v13 / (float32(0.0) + float32((1e10 + (-2.25 - 0.1))))) - 3.14159))))
                            v13 -= float32(((v31 + (v31 / 1e10)) + (((v31 * (v31 - float64(float32((v31 / 1.0))))) / 1.5) + v31)))
                            var v32 bool = true
                            _ = v31
                            _ = v32
                        }
                        if false {
                            v33 := GenS2{F0: 2147483647, F1: false, F2: 2147483647}
                            _ = v33
                        } else {
                            var v34 uint8 = 127
                            _ = v34
                        }
                        v35 := GenS4{F0: v17}
                        _ = v35
                    } else {
                        fmt.Println(v13, v13)
                        switch "a" {
                        case "\n\t":
                            v16 = append(v16, 1)
                            if false {
                                var v36 uint16 = 255
                                if (2147483648 == 453) {
                                    var v37 int64 = genFn0(-159, "foo bar")
                                    _ = v37
                                } else {
                                    v16 = append(v16, fi15)
                                    v36 -= 852
                                }
                                if false {
                                    if (!true) {
                                        switch 4294967295 {
                                        case -582:
                                            if true {
                                                v38 := GenS4{F0: v17}
                                                _ = v38
                                            }
                                        case 128:
                                            var v39 int = (^(fi14))
                                            _ = v39
                                        }
                                        var v40 uint = 2
                                        _ = v40
                                    }
                                    v41 := GenS2{F0: -1, F1: false, F2: 65535}
                                    _ = v41
                                } else {
                                    var v42 uint64 = 18446744073709551615
                                    if true {
                                        v43 := GenS5{F0: 0}
                                        v13 += v13
                                        _ = v43
                                    } else {
                                        if false {
                                            if (v42 != (v42 * (^(v42)))) {
                                                v16 = append(v16, int(v36))
                                                var v44 int64 = (v17 << ((965) & 63))
                                                _ = v44
                                            }
                                            fmt.Println(v42, fi14)
                                        }
                                    }
                                    if true {
                                        if (true && (!(255 < 0))) {
                                            fmt.Println(v13, fi15)
                                        }
                                        if (float64(float32((3.14159 / 1.5))) != 3.14159) {
                                            v42 = (^(v42))
                                        } else {
                                            if false {
                                                v42 += (^(v42))
                                                var v45 uint = 255
                                                fmt.Println(v42, v13, fi14)
                                                _ = v45
                                            } else {
                                                var v46 uint8 = 78
                                                v16 = append(v16, fi14)
                                                var v47 uint = 255
                                                _ = v46
                                                _ = v47
                                            }
                                            var v48 uint64 = v42
                                            v16 = append(v16, 0)
                                            _ = v48
                                        }
                                        var v49 uint = 1
                                        switch -2147483648 {
                                        case 42:
                                            var v50 float32 = -2.25
                                            _ = v50
                                        case 1:
                                            var v51 int32 = int32(v36)
                                            _ = v51
                                        default:
                                            if ((v36 % ((v36) | 1)) > (v36 & uint16(v42))) {
                                                var v52 int = (fi15 / ((fi14) | 1))
                                                var v53 bool = (!true)
                                                _ = v52
                                                _ = v53
                                            } else {
                                                switch fi15 {
                                                case 32768:
                                                    fmt.Println(fi14, v36)
                                                }
                                                if false {
                                                    switch fi15 {
                                                    case 127:
                                                        if false {
                                                            v54 := GenS3{F0: false, F1: (v36 % ((v36) | 1))}
                                                            var v55 int = 32767
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    case 128:
                                                        _ = len(v16)
                                                    default:
                                                        var v56 float32 = v13
                                                        _ = v56
                                                    }
                                                    var v57 string = (("a" + ("a" + "hello")) + "foo bar")
                                                    _ = v57
                                                } else {
                                                    v16 = append(v16, (fi14 ^ 4294967295))
                                                    var v58 uint = (v49 / ((18446744073709551615) | 1))
                                                    _ = v58
                                                }
                                                if len(v16) > 0 {
                                                    v16[0] = fi14
                                                }
                                                if (true && false) {
                                                    var v59 int16 = -32768
                                                    _ = v59
                                                }
                                                fmt.Println(fi15, fi14, v42)
                                            }
                                        }
                                        if true {
                                            fmt.Println(v17, v42, v17)
                                            v17 = 337
                                        } else {
                                            var v60 int64 = v17
                                            v61 := GenS1{F0: v60, F1: ("a" <= ((("a" + "a") + "") + "")), F2: 42}
                                            fmt.Println(v49, v13, fi14)
                                            v36 = (v36 / (((v36 % (((v36 * 32767)) | 1))) | 1))
                                            if ((float32((3.14159 + -2.25)) - 0.1) != (v13 * 0.0)) {
                                                v62 := GenS5{F0: 0}
                                                v63 := GenS3{F0: true, F1: (v36 >> (((v49 + (v49 - uint(v42)))) & 63))}
                                                _ = v62
                                                _ = v63
                                            }
                                            _ = v60
                                            _ = v61
                                        }
                                        _ = v49
                                    } else {
                                        if false {
                                            v64 := GenS5{F0: 161}
                                            v42 -= v42
                                            var v65 float32 = v13
                                            _ = v64
                                            _ = v65
                                        } else {
                                            var v66 int32 = -973
                                            var v67 uint64 = v42
                                            v68 := GenS6{F0: (v66 * (-(v66))), F1: true}
                                            if true {
                                                var v69 uint32 = 957
                                                if (!(!(fi14 <= fi14))) {
                                                    if len(v16) >= 2 {
                                                        v16 = v16[1:len(v16)]
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        var v70 int64 = (v17 - (v17 | 42))
                                                        v67 += (v67 & (^(v42)))
                                                        v71 := GenS7{F0: 42, F1: v69, F2: -1}
                                                        _ = v70
                                                        _ = v71
                                                    case "foo bar":
                                                        fmt.Println(v13)
                                                    }
                                                }
                                                _ = v69
                                            }
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                        v17 -= (v17 % (((-(v17))) | 1))
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                    }
                                    _ = v42
                                }
                                _ = v36
                            } else {
                                switch (fi15 / (((fi14 % ((fi14) | 1))) | 1)) {
                                case 286:
                                    v16 = append(v16, (fi15 * (fi14 - (fi15 - (fi15 % ((fi14) | 1))))))
                                    v17 = 9223372036854775807
                                    fmt.Println(v17)
                                case -1:
                                    var v72 uint64 = 42
                                    _ = v72
                                case 0:
                                    v73 := GenS8{F0: genFn0(-435, ""), F1: (fi15 - (^(fi15)))}
                                    _ = v73
                                default:
                                    if false {
                                        if len(v16) > 0 {
                                            v16[0] = (fi15 % ((fi15) | 1))
                                        }
                                        if len(v16) > 0 {
                                            v16[0] = (^(fi15))
                                        }
                                        var v74 int = 32768
                                        v16 = append(v16, fi15)
                                        _ = v74
                                    } else {
                                        v13 -= v13
                                        var v75 string = "a"
                                        _ = v75
                                    }
                                }
                            }
                        case "foo bar":
                            switch (fi15 & (fi15 << ((0) & 63))) {
                            case 9223372036854775807:
                                switch (fi15 << ((1) & 63)) {
                                case 9223372036854775807:
                                    v16 = append(v16, (fi15 >> ((1) & 63)))
                                case -980:
                                    v17 -= (-(v17))
                                case 4294967295:
                                    v17 = v17
                                default:
                                    if true {
                                        switch fi14 {
                                        case 127:
                                            var v76 uint = 42
                                            _ = v76
                                        case -62:
                                            v13 += v13
                                        }
                                        if (-32768 < -427) {
                                            var v77 int8 = -50
                                            if false {
                                                if (0.0 != 0.0) {
                                                    if (0 <= 108) {
                                                        v78 := GenS7{F0: 1, F1: 2147483648, F2: 1}
                                                        if true {
                                                            v79 := GenS8{F0: genFn0(-32768, ("foo bar" + ("" + (("a" + "") + (("foo bar" + ("" + (("\n\t" + "foo bar") + "hello"))) + (("hello" + "\n\t") + "\n\t")))))), F1: -1}
                                                            var v80 int = fi15
                                                            var v81 float64 = (-2.25 * -2.25)
                                                            var v82 int32 = -1
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        } else {
                                                            v77 = 127
                                                            v77 += -128
                                                            fmt.Println(v17, v17)
                                                        }
                                                        _ = v78
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v83 uint16 = 0
                                                        _ = v83
                                                    default:
                                                        if true {
                                                            if (int64(fi14) < (v17 ^ v17)) {
                                                                v17 += v17
                                                            }
                                                            var v84 uint64 = 4294967295
                                                            var v85 uint8 = 217
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                    }
                                                    var v86 uint = 65535
                                                    _ = v86
                                                }
                                                if (((v77 & (v77 / ((v77) | 1))) > -58) || true) {
                                                    fmt.Println(fi15, v77, fi15)
                                                    if false {
                                                        if true {
                                                            var v87 float64 = float64((float32(3.14159) + v13))
                                                            _ = v87
                                                        }
                                                        v88 := GenS8{F0: -9223372036854775808, F1: fi15}
                                                        _ = v88
                                                    } else {
                                                        v16 = append(v16, (fi14 & 127))
                                                        var v89 uint8 = 217
                                                        _ = v89
                                                    }
                                                    var v90 bool = false
                                                    _ = v90
                                                } else {
                                                    var v91 int16 = -32768
                                                    var v92 uint = 42
                                                    _ = v91
                                                    _ = v92
                                                }
                                            } else {
                                                switch 128 {
                                                case -822:
                                                    var v93 uint = 42
                                                    _ = v93
                                                default:
                                                    if true {
                                                        switch fi15 {
                                                        case -1:
                                                            if false {
                                                                var v94 uint8 = 1
                                                                v13 -= 3.14159
                                                                _ = v94
                                                            }
                                                            v17 = 9223372036854775807
                                                            var v95 uint64 = 569
                                                            var v96 string = "\n\t"
                                                            var v97 float64 = 3.14159
                                                            _ = v95
                                                            _ = v96
                                                            _ = v97
                                                        case 128:
                                                            var v98 int64 = genFn0(-1, ("\n\t" + (("" + ("" + "a")) + "")))
                                                            _ = v98
                                                        }
                                                        v99 := GenS9{F0: 0, F1: 1}
                                                        v100 := GenS10{F0: 65535, F1: 832, F2: v77}
                                                        fmt.Println(v17)
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                }
                                                v16 = append(v16, fi15)
                                                var v101 float64 = (0.0 - ((0.1 / 1e10) / 0.1))
                                                _ = v101
                                            }
                                            v13 = v13
                                            _ = v77
                                        }
                                        var v102 int = (^(fi15))
                                        _ = v102
                                    }
                                }
                            }
                        }
                        if (2147483647 != 0) {
                            _ = len(v16)
                        }
                    }
                }
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v13 = 0.0
    var v103 int8 = 0
    v13 = v13
    _ = v103
    _ = v13
    return (!(42 > 2))
}

func main() {
    var v1 uint = 42
    var v2 uint = (v1 | v1)
    var v3 uint8 = 128
    for fi104 := 0; fi104 < 5; fi104++ {
        for fi105 := 0; fi105 < 2; fi105++ {
            v3 = v3
            v1 += v2
            _ = fi105
        }
        _ = fi104
    }
    for fi106 := 0; fi106 < 9; fi106++ {
        var v107 string = "foo bar"
        fmt.Println(fi106)
        _ = fi106
        _ = v107
    }
    var v108 uint32 = 902
    if true {
        var v109 uint32 = (v108 | v108)
        v2 += (v1 / ((v2) | 1))
        fmt.Println(v3, v3)
        _ = v109
    }
    v110 := []int{2147483647, 350, -783}
    var v111 int = 42
    for fi112 := 0; fi112 < 10; fi112++ {
        var v113 int8 = -115
        if true {
            fmt.Println(fi112, v113, v111)
            v114 := GenS11{F0: (^(v111))}
            _ = v114
        } else {
            switch "\n\t" {
            case "\n\t":
                switch ("" + ("foo bar" + "foo bar")) {
                case "":
                    var v115 int16 = 42
                    _ = v115
                case "foo bar":
                    v108 = 2147483648
                default:
                    if len(v110) > 0 {
                        v110[0] = (fi112 % ((int(v3)) | 1))
                    }
                }
            }
        }
        for fi116 := 0; fi116 < 16; fi116++ {
            v3 = (^(v3))
            _ = fi116
        }
        _ = fi112
        _ = v113
    }
    switch (v111 % ((v111) | 1)) {
    case 0:
        v110 = append(v110, v111)
    case 32767:
        v117 := GenS9{F0: 42, F1: 716}
        _ = v117
    case -740:
        if true {
            fmt.Println(v3)
            v118 := GenS12{F0: 42, F1: 42, F2: 9223372036854775807}
            var v119 uint16 = 0
            _ = v118
            _ = v119
        } else {
            v120 := GenS7{F0: 127, F1: v108, F2: 286}
            _ = v120
        }
    default:
        if (255 == 1) {
            switch "" {
            case "hello":
                fmt.Println(v111, v2, v2)
            case "a":
                _ = len(v110)
            case "\n\t":
                for fi121 := 0; fi121 < 19; fi121++ {
                    v122 := GenS13{F0: v3}
                    if (false || false) {
                        v3 += 1
                        var v123 bool = ("" != "foo bar")
                        for fi124 := 0; fi124 < 2; fi124++ {
                            if (1 < 18446744073709551615) {
                                if v123 {
                                    v2 += v1
                                    v125 := GenS14{F0: "foo bar", F1: fi124, F2: (v108 & (v108 ^ 787))}
                                    var v126 uint8 = v3
                                    if (1e10 == (3.14159 + 0.1)) {
                                        var v127 uint8 = (v126 - 1)
                                        _ = v127
                                    }
                                    var v128 int = int(v126)
                                    _ = v125
                                    _ = v126
                                    _ = v128
                                }
                            }
                            _ = fi124
                        }
                        _ = v123
                    }
                    v108 += v108
                    _ = fi121
                    _ = v122
                }
            }
            v129 := GenS3{F0: (127 != 42), F1: 32767}
            v130 := GenS12{F0: -527, F1: 0, F2: v111}
            _ = v129
            _ = v130
        } else {
            for fi131 := 0; fi131 < 11; fi131++ {
                for fi132 := 0; fi132 < 18; fi132++ {
                    v133 := GenS15{F0: 0, F1: -355, F2: 319}
                    _ = len(v110)
                    if ("hello" > "a") {
                        v3 += (v3 | v3)
                    }
                    _ = fi132
                    _ = v133
                }
                var v134 uint16 = 0
                var v135 int64 = int64(fi131)
                _ = fi131
                _ = v134
                _ = v135
            }
            var v136 uint64 = 42
            var v137 uint = v1
            _ = v136
            _ = v137
        }
    }
    if false {
        if (1 == 533) {
            var v138 uint = v1
            _ = v138
        } else {
            v139 := GenS2{F0: 42, F1: (!true), F2: 1}
            for fi140 := 0; fi140 < 10; fi140++ {
                var v141 float64 = 0.1
                v142 := GenS16{F0: (false && false)}
                _ = fi140
                _ = v141
                _ = v142
            }
            _ = v139
        }
    }
    var v143 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v110:", len(v110), v110)
    fmt.Println("v111:", v111)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v108
    _ = v110
    _ = v111
    _ = v143
    _ = v2
    _ = v3
}
