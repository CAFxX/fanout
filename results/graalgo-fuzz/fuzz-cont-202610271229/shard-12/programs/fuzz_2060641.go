package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int8; F2 int64 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string; F1 uint; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64; F1 int32 }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 int32; F1 string }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int64 }
func genFn0() string {
    switch 260 {
    case -2147483648:
        var v4 uint64 = 1
        _ = v4
    case 4294967295:
        var v5 int32 = 1
        _ = v5
    }
    if ((!(!(-1 != -32768))) || (0 < 2147483647)) {
        if (0.1 <= (((-2.25 / 1.0) * 3.14159) * 0.0)) {
            v6 := []int{42, 2147483647, -509}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 20; fi7++ {
                if (2147483648 <= 107) {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v9 := GenS1{F0: 65535}
                        fmt.Println(fi7)
                        v10 := []int{fi7, fi8}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
                _ = fi7
            }
            v11 := []int{9223372036854775807, 4294967295, 4294967295, 2147483647}
            for fi12 := 0; fi12 < 20; fi12++ {
                if (2 <= 4294967295) {
                    v13 := GenS1{F0: 255}
                    v14 := GenS2{F0: 0}
                    _ = len(v11)
                    _ = v13
                    _ = v14
                }
                if (32767 >= -815) {
                    v15 := GenS2{F0: -32768}
                    _ = v15
                } else {
                    switch (fi12 ^ -128) {
                    case -9223372036854775808:
                        var v16 int64 = -1
                        _ = v16
                    case 9223372036854775807:
                        v17 := GenS1{F0: 65535}
                        _ = v17
                    case 32768:
                        var v18 float64 = (-2.25 / (1e10 - float64(((3.14159 - 0.1) * float32(0.1)))))
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 11; fi19++ {
                            switch "hello" {
                            case "hello":
                                var v20 uint32 = 2147483647
                                _ = v20
                            default:
                                var v21 bool = false
                                _ = v21
                            }
                            var v22 int = fi19
                            _ = fi19
                            _ = v22
                        }
                    }
                    if true {
                        var v23 uint32 = 42
                        _ = v23
                    }
                    if (2 == 737) {
                        fmt.Println(fi12)
                        var v24 int16 = 1
                        _ = v24
                    } else {
                        v11 = append(v11, fi12)
                        if len(v11) > 0 {
                            v11[0] = -251
                        }
                        if true {
                            var v25 uint64 = 2
                            _ = v25
                        } else {
                            if (-125 <= 100) {
                                v26 := GenS3{F0: 1, F1: 0, F2: -680}
                                switch "" {
                                case "\n\t":
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v27 float32 = (0.1 - (1.5 - float32(0.1)))
                                        _ = v27
                                    case "hello":
                                        var v28 string = ("\n\t" + ("foo bar" + "hello"))
                                        _ = v28
                                    default:
                                        var v29 int32 = 23
                                        _ = v29
                                    }
                                default:
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                }
                                _ = v26
                            } else {
                                switch fi12 {
                                case 165:
                                    var v30 string = "hello"
                                    _ = v30
                                }
                                var v31 uint32 = 2
                                if false {
                                    v32 := GenS4{F0: -797, F1: (v31 - (v31 << ((4294967295) & 63)))}
                                    v33 := GenS5{F0: 1}
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 uint8 = 255
                                    var v35 int64 = 1
                                    switch 0 {
                                    case 52:
                                        var v36 int = (fi12 + fi12)
                                        _ = v36
                                    default:
                                        switch ("hello" + (("a" + "a") + "foo bar")) {
                                        case "foo bar":
                                            v35 = 9223372036854775807
                                        case "a":
                                            var v37 float64 = ((0.0 - (((1e10 / -2.25) * (0.1 * (-2.25 - (1.5 - ((0.1 - 3.14159) * (float64(0.1) / float64(float32(-2.25)))))))) + ((1.5 / 3.14159) / 3.14159))) * float64(((0.0 - 1e10) + ((0.1 / 1.0) - 1e10))))
                                            _ = v37
                                        }
                                    }
                                    _ = v34
                                    _ = v35
                                }
                                _ = v31
                            }
                        }
                    }
                }
                if false {
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v11 = append(v11, 32768)
                        var v39 uint16 = 462
                        v40 := GenS6{F0: 42}
                        fmt.Println(v39)
                        var v41 float64 = float64(float32(3.14159))
                        var v42 float32 = -2.25
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    var v43 int64 = 0
                    _ = v43
                }
                _ = fi12
            }
            _ = v11
        }
        switch 32768 {
        case 4294967295:
            v44 := []int{-2147483648}
            _ = v44
        case 0:
            switch ("hello" + "foo bar") {
            case "\n\t":
                v45 := GenS6{F0: 611}
                _ = v45
            }
        }
        var v46 uint = 4294967295
        _ = v46
    }
    v47 := GenS1{F0: 32767}
    for fi48 := 0; fi48 < 20; fi48++ {
        if true {
            v49 := GenS2{F0: 32767}
            _ = v49
        } else {
            var v50 uint32 = 379
            if (("hello" + "a") < "hello") {
                for fi51 := 0; fi51 < 16; fi51++ {
                    var v52 int = (fi48 + 735)
                    _ = fi51
                    _ = v52
                }
                v50 = (^(v50))
                var v53 int = (^(fi48))
                _ = v53
            } else {
                v54 := GenS5{F0: 9223372036854775807}
                v50 -= v50
                _ = v54
            }
            _ = v50
        }
        v55 := []int{}
        if ((((42 != fi48) || (true || ((fi48 + (fi48 * -1)) < fi48))) || ("\n\t" == "foo bar")) || (42 != 373)) {
            var v56 uint64 = 2
            if (0 > -468) {
                if false {
                    v56 = (^(v56))
                }
                var v57 uint64 = (v56 >> ((0) & 63))
                v58 := GenS7{F0: "", F1: 255, F2: false}
                _ = v57
                _ = v58
            } else {
                switch "" {
                case "hello":
                    for fi59 := 0; fi59 < 3; fi59++ {
                        v60 := GenS5{F0: fi48}
                        _ = fi59
                        _ = v60
                    }
                case "\n\t":
                    v56 -= v56
                default:
                    fmt.Println(fi48)
                }
                fmt.Println(v56, fi48)
                v55 = append(v55, 311)
            }
            _ = v56
        }
        _ = fi48
        _ = v55
    }
    _ = v47
    return "foo bar"
}

func main() {
    v1 := []string{"", "hello", "\n\t", ("a" + "")}
    var v2 uint16 = 0
    v3 := make(map[string]string)
    fmt.Println(v2)
    var v61 float64 = (float64(((0.1 * 0.0) + -2.25)) + ((1.5 / ((1.5 / 0.1) / 1.0)) * (0.1 / 1.5)))
    for fi62 := 0; fi62 < 12; fi62++ {
        v2 += 42
        if len(v1) > 0 {
            v1[0] = ""
        }
        for fi63 := 0; fi63 < 16; fi63++ {
            v61 = v61
            _ = fi63
        }
        _ = fi62
    }
    fmt.Println(v61, v2)
    v61 = ((3.14159 - 3.14159) - 1e10)
    switch -9223372036854775808 {
    case 2147483647:
        v1 = append(v1, "")
    case 1:
        fmt.Println(v2, v2)
    case 0:
        var v64 int64 = -849
        _ = v64
    default:
        v2 += (v2 / ((0) | 1))
    }
    v65 := GenS7{F0: "", F1: 255, F2: ((true && (false || (291 >= 42))) && true)}
    for fi66 := 0; fi66 < 17; fi66++ {
        v61 += ((float64(((1e10 / 1.5) + -2.25)) * v61) - v61)
        _ = fi66
    }
    for fi67 := 0; fi67 < 3; fi67++ {
        _ = len(v3)
        switch (("\n\t" + "hello") + "hello") {
        case "hello":
            var v68 int32 = -2147483648
            _ = v68
        }
        var v69 float32 = 0.1
        _ = fi67
        _ = v69
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "hello")
    }
    var v70 int16 = -1
    fmt.Println(v70, v2)
    for fi71 := 0; fi71 < 16; fi71++ {
        v70 = v70
        if (2147483648 == 4294967295) {
            v72 := GenS8{F0: "hello"}
            if (136 > 742) {
                v2 = (v2 << ((1) & 63))
            } else {
                v1 = append(v1, ("foo bar" + "\n\t"))
                if true {
                    fmt.Println(fi71)
                    var v73 uint32 = 572
                    if true {
                        if (!(!true)) {
                            for fi74 := 0; fi74 < 12; fi74++ {
                                if ((42 != 18446744073709551615) && true) {
                                    var v75 int8 = 0
                                    _ = v75
                                }
                                if false {
                                    if false {
                                        _ = len(v3)
                                        if false {
                                            v2 += v2
                                            var v76 uint64 = uint64(v2)
                                            _ = v76
                                        }
                                    } else {
                                        var v77 int32 = -718
                                        var v78 int = 32768
                                        _ = v77
                                        _ = v78
                                    }
                                }
                                v73 -= v73
                                _ = fi74
                            }
                        }
                    }
                    _ = v73
                }
            }
            v79 := GenS6{F0: 770}
            _ = v72
            _ = v79
        }
        if (42 <= 162) {
            for fi80 := 0; fi80 < 17; fi80++ {
                switch fi71 {
                case 675:
                    if false {
                        v81 := GenS7{F0: "", F1: 792, F2: (199 > uint8(v2))}
                        if true {
                            v61 = v61
                        }
                        var v82 float64 = ((v61 - 3.14159) * v61)
                        if (v61 != 1e10) {
                            fmt.Println(fi80, v61, fi71)
                            var v83 float32 = 0.1
                            _ = v83
                        } else {
                            switch fi80 {
                            case 9223372036854775807:
                                v61 = v82
                            default:
                                v2 -= 2
                            }
                            if (true && (v2 < v2)) {
                                v84 := GenS9{F0: -128}
                                _ = v84
                            }
                            var v85 bool = false
                            _ = v85
                        }
                        _ = v81
                        _ = v82
                    } else {
                        var v86 uint32 = 2
                        _ = len(v1)
                        var v87 int64 = -9223372036854775808
                        _ = v86
                        _ = v87
                    }
                case 124:
                    var v88 uint16 = v2
                    _ = v88
                case -547:
                    fmt.Println(fi80, v70)
                }
                var v89 bool = true
                if (!(166 < 2)) {
                    var v90 float32 = ((0.1 - 1.5) * 1.5)
                    switch "\n\t" {
                    case "\n\t":
                        switch (fi80 - (fi80 ^ fi80)) {
                        case 32768:
                            v61 -= float64(v90)
                        default:
                            v89 = (-75 > 59)
                        }
                    case "":
                        v90 -= (v90 * float32(((((v61 / (-2.25 / -2.25)) + ((v61 / v61) / v61)) - 3.14159) / 1.5)))
                    case "hello":
                        if v89 {
                            if (!(true || false)) {
                                if (1 >= 4294967295) {
                                    v89 = v89
                                    var v91 uint8 = 2
                                    var v92 bool = v89
                                    v61 += v61
                                    fmt.Println(v90, v89, fi71)
                                    _ = v91
                                    _ = v92
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch -162 {
                                        case 32768:
                                            fmt.Println(v90)
                                        case 431:
                                            fmt.Println(fi71, v70, v70)
                                        case 0:
                                            var v93 int64 = 42
                                            _ = v93
                                        default:
                                            v3["foo bar"] = "a"
                                        }
                                    case "a":
                                        v90 += v90
                                    case "\n\t":
                                        var v94 uint16 = (v2 + (v2 + 42))
                                        _ = v94
                                    default:
                                        v95 := GenS3{F0: -1, F1: 100, F2: 9223372036854775807}
                                        _ = v95
                                    }
                                    v96 := GenS1{F0: v2}
                                    var v97 uint32 = 2147483647
                                    _ = v96
                                    _ = v97
                                }
                            } else {
                                fmt.Println(v70)
                                v2 = (v2 | 429)
                                fmt.Println(fi80)
                                if (!true) {
                                    v98 := v3[("a" + "foo bar")]
                                    _ = v98
                                }
                                v3["a"] = "hello"
                                if v89 {
                                    fmt.Println(v90)
                                    if false {
                                        var v99 int16 = 42
                                        v70 += -32768
                                        var v100 int32 = 1
                                        _ = v100
                                        _ = v99
                                    } else {
                                        v89 = (v89 || true)
                                    }
                                    v89 = v89
                                } else {
                                    _ = len(v3)
                                    var v101 bool = v89
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v101
                                }
                            }
                            v3[(((("a" + "\n\t") + (("" + "a") + "")) + "a") + "hello")] = ""
                            v89 = ((v89 || v89) || v89)
                        } else {
                            delete(v3, "hello")
                        }
                    }
                    var v102 string = "foo bar"
                    _ = v102
                    _ = v90
                }
                _ = fi80
                _ = v89
            }
        }
        _ = fi71
    }
    switch ("foo bar" + ("\n\t" + "")) {
    case "":
        v103 := GenS9{F0: -64}
        _ = v103
    case "a":
        if (v2 == (v2 % ((2) | 1))) {
            var v104 int32 = 2147483647
            v105 := GenS3{F0: 42, F1: 24, F2: 407}
            _ = v104
            _ = v105
        }
    }
    v2 = (v2 | 43)
    var v106 float32 = 1e10
    if true {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            if false {
                var v107 uint64 = 1
                v108 := GenS1{F0: 0}
                var v109 float32 = 3.14159
                _ = v107
                _ = v108
                _ = v109
            } else {
                if false {
                    if (false || (562 != 700)) {
                        v1 = append(v1, "\n\t")
                    } else {
                        var v110 uint = 18446744073709551615
                        _ = v110
                    }
                } else {
                    _ = len(v1)
                    if false {
                        if true {
                            v2 = v2
                            if true {
                                if false {
                                    v111 := GenS8{F0: ""}
                                    _ = v111
                                } else {
                                    if true {
                                        switch 42 {
                                        case -2147483648:
                                            _ = len(v3)
                                        case 127:
                                            fmt.Println(v70)
                                        }
                                        switch 1 {
                                        case -112:
                                            var v112 int64 = 9223372036854775807
                                            _ = v112
                                        case -9223372036854775808:
                                            var v113 uint32 = 4294967295
                                            _ = v113
                                        case 0:
                                            if true {
                                                var v114 string = "a"
                                                fmt.Println(v2)
                                                _ = v114
                                            } else {
                                                v70 = 1
                                                fmt.Println(v106, v2, v70)
                                                v115 := GenS10{F0: 18446744073709551615, F1: 492}
                                                _ = v115
                                            }
                                        }
                                    }
                                    v61 = ((v61 * v61) / v61)
                                    v3[("\n\t" + (((((("hello" + ("a" + "hello")) + "a") + "") + "foo bar") + "\n\t") + ("foo bar" + "")))] = "foo bar"
                                }
                                for fi116 := 0; fi116 < 17; fi116++ {
                                    var v117 uint64 = 2
                                    if true {
                                        var v118 int = (-(fi116))
                                        _ = v118
                                    } else {
                                        for fi119 := 0; fi119 < 17; fi119++ {
                                            v61 -= (0.0 - (v61 / ((float64((v106 - (v106 * (0.1 + 1.5)))) / 1.0) / (v61 * v61))))
                                            v120 := GenS7{F0: (("" + "foo bar") + ("hello" + "a")), F1: uint(v117), F2: ("\n\t" > "")}
                                            switch (fi116 >> ((1) & 63)) {
                                            case 727:
                                                if false {
                                                    if (v70 <= v70) {
                                                        if false {
                                                            var v121 uint = 255
                                                            v2 = v2
                                                            _ = v121
                                                        } else {
                                                            var v122 float64 = (v61 * v61)
                                                            v106 = float32((((((v61 - float64(((float32(3.14159) / v106) * ((1.5 - float32(0.0)) / 0.1)))) + v61) / v61) * 0.0) / ((v122 * v122) - v122)))
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                _ = len(v3)
                                                            case "a":
                                                                var v123 uint32 = 247
                                                                _ = v123
                                                            case "\n\t":
                                                                v124 := GenS11{F0: v2, F1: "a"}
                                                                _ = v124
                                                            default:
                                                                v125 := GenS12{F0: 2147483647, F1: (v117 << ((0) & 63))}
                                                                _ = v125
                                                            }
                                                            v126 := GenS3{F0: -23, F1: 112, F2: -90}
                                                            v2 = (^(v2))
                                                            switch "hello" {
                                                            case "":
                                                                if (!(v2 != (v2 - (v2 / ((2) | 1))))) {
                                                                    var v127 float64 = v61
                                                                    _ = v127
                                                                } else {
                                                                    fmt.Println(v122)
                                                                    v61 = 0.1
                                                                    var v128 int8 = -128
                                                                    v117 = (^(v117))
                                                                    _ = v128
                                                                }
                                                            default:
                                                                if true {
                                                                    if (false || false) {
                                                                        var v129 int64 = -1
                                                                        v70 = (v70 >> ((686) & 63))
                                                                        if false {
                                                                            _ = len(v3)
                                                                            v129 += (v129 ^ (^(v129)))
                                                                            v130 := GenS10{F0: 857, F1: -155}
                                                                            if false {
                                                                                v117 = (v117 ^ v117)
                                                                                v129 -= v129
                                                                                v122 += ((1.5 + (v61 + ((((v122 * v122) / (float64(float32(-2.25)) * -2.25)) * ((0.0 - v122) - ((1e10 - (v61 * v122)) * 1e10))) + 3.14159))) - 1.5)
                                                                            }
                                                                            var v131 float64 = v61
                                                                            _ = v130
                                                                            _ = v131
                                                                        } else {
                                                                            v129 += (v129 + (v129 ^ v129))
                                                                            var v132 int8 = 42
                                                                            v129 = (v129 % ((v129) | 1))
                                                                            _ = v132
                                                                        }
                                                                        _ = v129
                                                                    } else {
                                                                        var v133 uint32 = 42
                                                                        fmt.Println(fi116)
                                                                        var v134 string = ((("" + "") + "") + "a")
                                                                        _ = v133
                                                                        _ = v134
                                                                    }
                                                                    var v135 string = ("" + ("hello" + "foo bar"))
                                                                    _ = v135
                                                                }
                                                            }
                                                            _ = v122
                                                            _ = v126
                                                        }
                                                        var v136 int64 = 594
                                                        var v137 uint16 = 2
                                                        v137 -= 272
                                                        if (v136 < (v136 * v136)) {
                                                            var v138 int32 = -944
                                                            _ = v138
                                                        }
                                                        if true {
                                                            var v139 int32 = -2147483648
                                                            if len(v1) > 0 {
                                                                v1[0] = ""
                                                            }
                                                            _ = v139
                                                        }
                                                        v140 := GenS13{F0: 1, F1: ""}
                                                        _ = v136
                                                        _ = v137
                                                        _ = v140
                                                    }
                                                    v1 = append(v1, "foo bar")
                                                    fmt.Println(v70, v61, v117)
                                                } else {
                                                    v61 += float64(v106)
                                                }
                                            case 960:
                                                var v141 int8 = 52
                                                _ = v141
                                            default:
                                                delete(v3, ("hello" + (("foo bar" + "hello") + "\n\t")))
                                            }
                                            _ = fi119
                                            _ = v120
                                        }
                                        if ((^(v70)) > 1000) {
                                            switch (-(fi116)) {
                                            case 127:
                                                if false {
                                                    var v142 int16 = (v70 << ((714) & 63))
                                                    _ = v142
                                                } else {
                                                    v3[("foo bar" + "hello")] = ("hello" + "hello")
                                                    switch (fi116 & (-(fi116))) {
                                                    case 128:
                                                        var v143 uint32 = 2147483647
                                                        _ = v143
                                                    case 1:
                                                        var v144 float32 = v106
                                                        _ = v144
                                                    case -9223372036854775808:
                                                        fmt.Println(v61, v2, v2)
                                                    default:
                                                        var v145 int64 = -9223372036854775808
                                                        _ = v145
                                                    }
                                                }
                                            case 9223372036854775807:
                                                var v146 float64 = 0.1
                                                _ = v146
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = (("foo bar" + (("foo bar" + "a") + ((((("a" + (("a" + "hello") + "\n\t")) + "") + "") + "foo bar") + ("" + (("" + "\n\t") + (("hello" + ("foo bar" + ("\n\t" + "hello"))) + ("hello" + "hello"))))))) + "foo bar")
                                                }
                                            }
                                            v117 -= 1
                                            v70 = (v70 * v70)
                                        } else {
                                            v2 -= v2
                                        }
                                        v147 := GenS14{F0: 1}
                                        _ = v147
                                    }
                                    if (2 != 971) {
                                        if (100 <= 0) {
                                            v148 := GenS8{F0: "foo bar"}
                                            _ = len(v1)
                                            _ = v148
                                        } else {
                                            _ = len(v3)
                                            v149 := GenS6{F0: 4294967295}
                                            v150 := GenS15{F0: -1}
                                            _ = v149
                                            _ = v150
                                        }
                                    } else {
                                        v2 = v2
                                        for fi151 := 0; fi151 < 20; fi151++ {
                                            v2 = 2
                                            fmt.Println(v106, v117)
                                            var v152 uint = 845
                                            _ = len(v3)
                                            var v153 int8 = -1
                                            _ = fi151
                                            _ = v152
                                            _ = v153
                                        }
                                        if true {
                                            if (!true) {
                                                fmt.Println(v117, fi116, v117)
                                                var v154 uint = 2
                                                _ = v154
                                            } else {
                                                var v155 uint = 42
                                                for fi156 := 0; fi156 < 9; fi156++ {
                                                    var v157 int64 = 1
                                                    var v158 uint64 = 1
                                                    v158 = v117
                                                    _ = fi156
                                                    _ = v157
                                                    _ = v158
                                                }
                                                _ = v155
                                            }
                                        }
                                    }
                                    _ = fi116
                                    _ = v117
                                }
                            }
                        } else {
                            var v159 uint8 = 151
                            _ = v159
                        }
                    }
                }
            }
        case "\n\t":
            v70 = (v70 ^ v70)
        case "hello":
            _ = len(v1)
        default:
            if false {
                var v160 uint16 = v2
                _ = v160
            } else {
                for fi161 := 0; fi161 < 17; fi161++ {
                    _ = len(v3)
                    _ = fi161
                }
            }
        }
    } else {
        v70 = v70
        var v162 uint16 = 475
        _ = v162
    }
    var v163 uint32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("v163:", v163)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v61:", v61)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v106
    _ = v163
    _ = v2
    _ = v3
    _ = v61
    _ = v65
    _ = v70
}
