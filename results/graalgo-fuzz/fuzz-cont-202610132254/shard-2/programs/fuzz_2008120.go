package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8; F1 int16; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool }

func main() {
    var v1 string = ((("a" + "a") + "foo bar") + "\n\t")
    v2 := make(map[int]bool)
    var v3 int16 = 1
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]uint32)
    var v6 uint32 = 2147483647
    switch v1 {
    case "\n\t":
        v7 := GenS1{F0: 557, F1: (v6 << ((255) & 63))}
        _ = v7
    case "hello":
        delete(v4, (v6 % ((v6) | 1)))
    }
    v1 = ("foo bar" + "\n\t")
    switch 9223372036854775807 {
    case -9223372036854775808:
        v3 += v3
    case 0:
        if true {
            v3 -= -818
            v1 = (("a" + v1) + "hello")
            if false {
                _ = len(v4)
                var v8 int = -2147483648
                _ = v8
            }
        }
    case 2147483647:
        var v9 int = -9223372036854775808
        _ = v9
    default:
        if (129 != 949) {
            switch (v1 + "\n\t") {
            case "foo bar":
                for fi10 := 0; fi10 < 15; fi10++ {
                    v6 += 531
                    v3 += v3
                    _ = fi10
                }
            }
            v11 := []int{-9223372036854775808}
            _ = v11
        } else {
            v12 := GenS2{F0: 42, F1: int8(v6)}
            _ = v12
        }
    }
    var v13 float32 = 0.0
    fmt.Println(v6, v6, v13)
    if false {
        v14 := GenS1{F0: 4294967295, F1: 0}
        if (v13 > ((v13 / v13) * (0.0 - v13))) {
            if true {
                v3 = (-(v3))
                var v15 bool = false
                v6 = (v6 | (v6 & v6))
                _ = v15
            } else {
                for fi16 := 0; fi16 < 6; fi16++ {
                    _ = len(v2)
                    v17 := GenS2{F0: 42, F1: 100}
                    _ = fi16
                    _ = v17
                }
                if (((-2.25 / ((((1.5 + 0.1) - 0.0) / 0.1) / float64(3.14159))) - 1e10) <= 3.14159) {
                    var v18 uint32 = v6
                    v19 := GenS1{F0: 4294967295, F1: (^(v6))}
                    fmt.Println(v13)
                    _ = v18
                    _ = v19
                } else {
                    v1 = v1
                }
                v13 = (v13 - 1.5)
            }
            if (!true) {
                if (true && true) {
                    if (255 < 42) {
                        v13 = v13
                    }
                    var v20 uint = 42
                    v20 = (v20 / ((v20) | 1))
                    _ = v20
                } else {
                    var v21 uint8 = 255
                    var v22 uint32 = 2147483648
                    _ = v21
                    _ = v22
                }
            } else {
                if true {
                    v6 += 430
                    for fi23 := 0; fi23 < 6; fi23++ {
                        for fi24 := 0; fi24 < 20; fi24++ {
                            v25 := []int{32768, fi24}
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            _ = fi24
                            _ = v25
                        }
                        _ = fi23
                    }
                    v13 = v13
                } else {
                    v4[(v6 - 944)] = (^(v6))
                    var v26 int32 = -1
                    var v27 int8 = 1
                    _ = v26
                    _ = v27
                }
                switch -1 {
                case 144:
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 int16 = (v3 | v3)
                        switch (fi28 * fi28) {
                        case 32768:
                            v30 := GenS2{F0: 42, F1: -128}
                            _ = v30
                        case 32767:
                            for fi31 := 0; fi31 < 11; fi31++ {
                                var v32 string = v1
                                var v33 uint64 = 2
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                        _ = fi28
                        _ = v29
                    }
                case 9223372036854775807:
                    fmt.Println(v3)
                default:
                    v6 = v6
                }
                v3 = v3
            }
        } else {
            for fi34 := 0; fi34 < 15; fi34++ {
                fmt.Println(v1, fi34, v1)
                _ = fi34
            }
            fmt.Println(v13, v3, v13)
            if ((v6 ^ 2) == v6) {
                v13 = 3.14159
                switch 9223372036854775807 {
                case 771:
                    v35 := []int{128, 0}
                    _ = v35
                case 0:
                    v1 = (((("hello" + v1) + v1) + v1) + ("" + (v1 + v1)))
                case 32768:
                    var v36 float32 = 0.1
                    _ = v36
                default:
                    v1 = v1
                }
            } else {
                var v37 uint64 = 732
                _ = v37
            }
        }
        switch 128 {
        case -798:
            switch v1 {
            case "a":
                delete(v2, 886)
            case "":
                v38 := GenS3{F0: 127, F1: 32767, F2: 0}
                _ = v38
            case "foo bar":
                var v39 bool = (true && true)
                _ = v39
            default:
                switch int(v6) {
                case 9223372036854775807:
                    var v40 uint = 18446744073709551615
                    _ = v40
                case -783:
                    for fi41 := 0; fi41 < 16; fi41++ {
                        v3 += -1
                        v1 = (v1 + v1)
                        _ = fi41
                    }
                case 32767:
                    if (!false) {
                        switch (v1 + ((((("foo bar" + "a") + v1) + ((v1 + v1) + v1)) + (("\n\t" + v1) + "")) + (("" + v1) + ""))) {
                        case "a":
                            var v42 bool = true
                            _ = v42
                        case "hello":
                            v13 = v13
                        default:
                            var v43 uint8 = 1
                            _ = v43
                        }
                        switch -2147483648 {
                        case 915:
                            if ((!false) && false) {
                                switch v1 {
                                case "foo bar":
                                    v13 = (3.14159 - v13)
                                case "hello":
                                    for fi44 := 0; fi44 < 14; fi44++ {
                                        v45 := v5[(v1 + "hello")]
                                        v45 -= (v45 + v6)
                                        if false {
                                            fmt.Println(fi44)
                                            v6 = v45
                                        } else {
                                            switch int(v3) {
                                            case 42:
                                                switch fi44 {
                                                case 830:
                                                    v46 := GenS1{F0: 175, F1: v45}
                                                    _ = v46
                                                case 128:
                                                    var v47 uint16 = 0
                                                    _ = v47
                                                default:
                                                    v6 -= 0
                                                }
                                            default:
                                                var v48 uint = 1
                                                _ = v48
                                            }
                                            v49 := []int{}
                                            switch v1 {
                                            case "foo bar":
                                                var v50 int = int(v45)
                                                _ = v50
                                            }
                                            _ = v49
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                case "":
                                    switch 42 {
                                    case 32767:
                                        for fi51 := 0; fi51 < 12; fi51++ {
                                            v13 += v13
                                            var v52 uint = 2
                                            for fi53 := 0; fi53 < 7; fi53++ {
                                                v54 := GenS1{F0: (v52 / (((v52 - 1)) | 1)), F1: 4294967295}
                                                switch (("a" + v1) + "") {
                                                case "":
                                                    if ("" < "") {
                                                        var v55 uint64 = 297
                                                        v6 += 42
                                                        var v56 uint32 = uint32(v55)
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                case "foo bar":
                                                    switch v1 {
                                                    case "hello":
                                                        v57 := []int{fi53, fi53}
                                                        _ = v57
                                                    case "foo bar":
                                                        var v58 float32 = 1e10
                                                        _ = v58
                                                    }
                                                case "hello":
                                                    var v59 float32 = 1.5
                                                    _ = v59
                                                }
                                                fmt.Println(v52, v13)
                                                v52 = 255
                                                _ = fi53
                                                _ = v54
                                            }
                                            _ = fi51
                                            _ = v52
                                        }
                                    default:
                                        var v60 float32 = (-2.25 / 0.1)
                                        _ = v60
                                    }
                                default:
                                    if true {
                                        if (!(1 < int8(v3))) {
                                            var v61 int = 4294967295
                                            switch v61 {
                                            case -1:
                                                var v62 uint8 = 168
                                                _ = v62
                                            case 2147483647:
                                                switch v1 {
                                                case "\n\t":
                                                    v13 += float32((3.14159 - 1.5))
                                                case "a":
                                                    if false {
                                                        v63 := []int{v61}
                                                        for fi64 := 0; fi64 < 8; fi64++ {
                                                            v61 += v61
                                                            if true {
                                                                v65 := GenS4{F0: int16(v6)}
                                                                if true {
                                                                    v3 = 1000
                                                                    var v66 int64 = 321
                                                                    v61 -= (-(fi64))
                                                                    _ = v66
                                                                } else {
                                                                    if true {
                                                                        switch v1 {
                                                                        case "\n\t":
                                                                            switch 9223372036854775807 {
                                                                            case 852:
                                                                                v67 := GenS2{F0: 2147483647, F1: 82}
                                                                                _ = v67
                                                                            case -854:
                                                                                switch v1 {
                                                                                case "hello":
                                                                                    v61 -= v61
                                                                                case "":
                                                                                    switch fi64 {
                                                                                    case -453:
                                                                                        for fi68 := 0; fi68 < 15; fi68++ {
                                                                                            v63 = append(v63, fi64)
                                                                                            var v69 float64 = 0.0
                                                                                            _ = fi68
                                                                                            _ = v69
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(fi64, v6, v6)
                                                                        default:
                                                                            if (!(true || true)) {
                                                                                v6 -= 585
                                                                                var v70 int64 = 0
                                                                                _ = v70
                                                                            } else {
                                                                                for fi71 := 0; fi71 < 11; fi71++ {
                                                                                    var v72 int8 = 62
                                                                                    v73 := GenS3{F0: (v72 | v72), F1: v3, F2: 214}
                                                                                    _ = fi71
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                }
                                                                            }
                                                                        }
                                                                        v6 += 0
                                                                        v74 := GenS4{F0: v3}
                                                                        _ = v74
                                                                    }
                                                                    var v75 bool = (false && ((v1 + "") >= v1))
                                                                    _ = v75
                                                                }
                                                                _ = v65
                                                            }
                                                            if (false && true) {
                                                                v76 := GenS5{F0: 100}
                                                                delete(v2, (v61 | 918))
                                                                if len(v63) >= 2 {
                                                                    v63 = v63[1:len(v63)]
                                                                }
                                                                _ = v76
                                                            }
                                                            _ = fi64
                                                        }
                                                        _ = v63
                                                    } else {
                                                        v6 = v6
                                                    }
                                                case "foo bar":
                                                    for fi77 := 0; fi77 < 5; fi77++ {
                                                        v78 := GenS3{F0: 100, F1: (v3 / ((v3) | 1)), F2: 1}
                                                        _ = fi77
                                                        _ = v78
                                                    }
                                                default:
                                                    var v79 uint = 18446744073709551615
                                                    _ = v79
                                                }
                                            case -630:
                                                if false {
                                                    for fi80 := 0; fi80 < 15; fi80++ {
                                                        var v81 int32 = 0
                                                        _ = fi80
                                                        _ = v81
                                                    }
                                                    var v82 uint16 = 65535
                                                    _ = v82
                                                }
                                            default:
                                                switch "" {
                                                case "":
                                                    v83 := []int{v61, v61, v61, (v61 * v61)}
                                                    _ = v83
                                                default:
                                                    for fi84 := 0; fi84 < 10; fi84++ {
                                                        v85 := GenS5{F0: 89}
                                                        v86 := GenS2{F0: 729, F1: -128}
                                                        var v87 uint8 = 128
                                                        _ = fi84
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                }
                                            }
                                            v88 := GenS5{F0: 36}
                                            _ = v61
                                            _ = v88
                                        } else {
                                            if ((348 < 4294967295) && (575 == 128)) {
                                                var v89 uint32 = 467
                                                for fi90 := 0; fi90 < 12; fi90++ {
                                                    switch -865 {
                                                    case 9223372036854775807:
                                                        v91 := []int{fi90}
                                                        _ = v91
                                                    case 172:
                                                        v92 := GenS1{F0: 255, F1: v89}
                                                        _ = v92
                                                    default:
                                                        if false {
                                                            for fi93 := 0; fi93 < 17; fi93++ {
                                                                var v94 uint8 = 128
                                                                _ = fi93
                                                                _ = v94
                                                            }
                                                        } else {
                                                            v95 := []int{32767, fi90, (fi90 - fi90)}
                                                            _ = v95
                                                        }
                                                    }
                                                    _ = fi90
                                                }
                                                if true {
                                                    v96 := v4[v6]
                                                    _ = v96
                                                } else {
                                                    var v97 uint8 = 2
                                                    _ = v97
                                                }
                                                _ = v89
                                            }
                                            if true {
                                                var v98 float64 = ((float64(v13) * -2.25) + (1e10 / (((float64(float32(0.1)) + (float64(1.5) - 3.14159)) + (-2.25 * float64(1.5))) * 1.5)))
                                                _ = v98
                                            } else {
                                                v99 := GenS6{F0: 127, F1: 4294967295, F2: 0}
                                                v13 = v13
                                                var v100 int8 = 0
                                                _ = v100
                                                _ = v99
                                            }
                                            v101 := GenS7{F0: (v3 % (((^(v3))) | 1))}
                                            _ = v101
                                        }
                                        for fi102 := 0; fi102 < 2; fi102++ {
                                            for fi103 := 0; fi103 < 1; fi103++ {
                                                v104 := GenS2{F0: 2147483647, F1: -1}
                                                fmt.Println(fi102, v6, v13)
                                                v13 += float32(((0.1 + 1.5) / 3.14159))
                                                _ = fi103
                                                _ = v104
                                            }
                                            switch fi102 {
                                            case 1:
                                                fmt.Println(v1)
                                            case 4294967295:
                                                var v105 uint16 = 42
                                                _ = v105
                                            }
                                            fmt.Println(fi102, v6)
                                            _ = fi102
                                        }
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            var v106 int64 = 1
                                            _ = v106
                                        default:
                                            v107 := GenS2{F0: 0, F1: -1}
                                            _ = v107
                                        }
                                    }
                                }
                                var v108 uint8 = 127
                                _ = v108
                            }
                        default:
                            v3 -= (v3 & (-(v3)))
                        }
                        v109 := v2[128]
                        _ = v109
                    } else {
                        v110 := GenS1{F0: 255, F1: v6}
                        _ = v110
                    }
                default:
                    for fi111 := 0; fi111 < 16; fi111++ {
                        if true {
                            v1 = v1
                        }
                        v1 = ((("hello" + v1) + (((((v1 + ((("" + v1) + v1) + ("foo bar" + v1))) + (v1 + ((v1 + v1) + v1))) + (v1 + ("hello" + "\n\t"))) + (((v1 + v1) + ((v1 + (v1 + v1)) + "foo bar")) + v1)) + (v1 + v1))) + "hello")
                        switch v1 {
                        case "\n\t":
                            if ((v13 * v13) > ((v13 - -2.25) / (float32(3.14159) + ((((float32(float64((v13 / v13))) + v13) - (float32(3.14159) - v13)) * v13) / 1e10)))) {
                                _ = len(v2)
                            } else {
                                var v112 bool = (-2147483648 > 437)
                                _ = v112
                            }
                        }
                        _ = fi111
                    }
                }
            }
        case 4294967295:
            for fi113 := 0; fi113 < 8; fi113++ {
                switch v1 {
                case "":
                    fmt.Println(v1, v3, v13)
                case "a":
                    for fi114 := 0; fi114 < 10; fi114++ {
                        v115 := []int{}
                        _ = fi114
                        _ = v115
                    }
                }
                if false {
                    v116 := []int{}
                    v13 = float32((((0.1 / 0.1) / 3.14159) - 1.5))
                    var v117 bool = (-2.25 != (float32((3.14159 / ((3.14159 * 3.14159) * float64(v13)))) / float32((3.14159 - ((((0.1 + 1e10) + (0.1 + 0.1)) * (1e10 / 3.14159)) - 3.14159)))))
                    _ = v116
                    _ = v117
                }
                _ = fi113
            }
        }
        _ = v14
    }
    var v118 string = v1
    v13 += 0.0
    for fi119 := 0; fi119 < 8; fi119++ {
        for fi120 := 0; fi120 < 1; fi120++ {
            fmt.Println(v1)
            switch (fi120 | 127) {
            case 1:
                v13 -= ((v13 * -2.25) - (((v13 - v13) + ((v13 * ((float32(1e10) / 1e10) - (v13 * float32(-2.25)))) - v13)) * float32(1.5)))
            case 111:
                v1 = "hello"
            case -134:
                var v121 bool = (!true)
                _ = v121
            }
            switch -2147483648 {
            case 0:
                fmt.Println(v1)
            case 2147483647:
                if false {
                    var v122 float64 = -2.25
                    delete(v5, v1)
                    v6 += v6
                    if ((true || (false && (!(0.0 != -2.25)))) && (81 <= 2)) {
                        if (2147483647 != -2147483648) {
                            fmt.Println(v122, fi120, v1)
                        } else {
                            var v123 uint32 = (v6 * (v6 & v6))
                            _ = v123
                        }
                    }
                    var v124 int = (fi120 ^ fi119)
                    _ = v122
                    _ = v124
                }
            default:
                var v125 int32 = 1
                _ = v125
            }
            _ = fi120
        }
        v118 = v118
        var v126 int = (fi119 + (-(fi119)))
        _ = fi119
        _ = v126
    }
    if false {
        v13 = float32((1e10 * 0.0))
        var v127 uint64 = 18446744073709551615
        _ = v127
    } else {
        fmt.Println(v1, v118, v118)
        for fi128 := 0; fi128 < 11; fi128++ {
            var v129 int8 = 0
            v130 := []int{(fi128 | -868), fi128, fi128, 301}
            if true {
                v13 = -2.25
                var v131 uint8 = 42
                _ = v131
            } else {
                var v132 float32 = ((0.0 * 0.0) - (3.14159 * (v13 + v13)))
                _ = v132
            }
            _ = fi128
            _ = v129
            _ = v130
        }
    }
    fmt.Println(v3, v3, v118)
    for fi133 := 0; fi133 < 1; fi133++ {
        v118 = v1
        if (9223372036854775807 <= -1) {
            v118 = "hello"
            v6 = (^(v6))
        }
        switch 32768 {
        case -1:
            for fi134 := 0; fi134 < 19; fi134++ {
                fmt.Println(v118, v6, v118)
                var v135 uint16 = 0
                _ = fi134
                _ = v135
            }
        case -2147483648:
            v136 := v5[(v1 + "hello")]
            _ = v136
        default:
            var v137 int32 = 1
            _ = v137
        }
        _ = fi133
    }
    for fi138 := 0; fi138 < 3; fi138++ {
        if false {
            switch fi138 {
            case -355:
                v1 = "\n\t"
            case 32767:
                v1 = (v1 + "a")
            default:
                v139 := GenS3{F0: -26, F1: (^(v3)), F2: 192}
                _ = v139
            }
        }
        _ = fi138
    }
    fmt.Println(v1)
    v140 := GenS1{F0: 0, F1: v6}
    for fi141 := 0; fi141 < 12; fi141++ {
        for fi142 := 0; fi142 < 1; fi142++ {
            if true {
                fmt.Println(v1, fi141, v3)
            }
            var v143 uint16 = 656
            _ = fi142
            _ = v143
        }
        var v144 float64 = 1.5
        if false {
            if false {
                v145 := v5[""]
                v146 := v2[fi141]
                _ = v145
                _ = v146
            } else {
                var v147 bool = (float32(((-2.25 * ((v144 / v144) * v144)) + v144)) == (v13 + (v13 - -2.25)))
                _ = v147
            }
            switch fi141 {
            case -2147483648:
                v6 += v6
            case 4294967295:
                v148 := v4[(v6 * 4294967295)]
                _ = v148
            case 42:
                v149 := GenS8{F0: (!false)}
                _ = v149
            default:
                v150 := GenS8{F0: (1 == 1)}
                _ = v150
            }
            v151 := []int{(fi141 >> ((18446744073709551615) & 63))}
            _ = v151
        }
        _ = fi141
        _ = v144
    }
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v118
    _ = v13
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
