package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint64; F2 int }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int8; F1 int32 }
type GenS14 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS15 struct { F0 uint64; F1 uint8; F2 bool }
func genFn0(p0 bool) uint16 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{-9223372036854775808, (-(fi7))}
        var v9 int32 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    if (!(!p0)) {
        v10 := GenS1{F0: 113}
        var v11 int16 = 32767
        var v12 float32 = 0.1
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{127}
    if false {
        v14 := GenS1{F0: 42}
        p0 = (700 > 1)
        _ = v14
    }
    _ = p0
    _ = v13
    return 2
}
func genFn1(p0 uint16, p1 uint8) bool {
    if true {
        v15 := GenS2{F0: 569}
        if (true && false) {
            var v16 uint64 = 42
            var v17 int8 = 100
            _ = v16
            _ = v17
        }
        _ = v15
    } else {
        var v18 uint = 65535
        v18 -= v18
        var v19 int8 = -1
        _ = v18
        _ = v19
    }
    switch -964 {
    case 128:
        var v20 int8 = -1
        _ = v20
    case -219:
        v21 := []int{42, -981}
        _ = v21
    default:
        v22 := GenS3{F0: 1, F1: 0, F2: 18446744073709551615}
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        fmt.Println(p0, p1)
        if (65535 > 667) {
            for fi24 := 0; fi24 < 3; fi24++ {
                if true {
                    var v25 uint16 = 732
                    v25 = 255
                    v25 -= (v25 << ((18446744073709551615) & 63))
                    _ = v25
                }
                fmt.Println(fi23, fi24, fi24)
                var v26 int16 = -1
                _ = fi24
                _ = v26
            }
            var v27 bool = false
            if v27 {
                v28 := []int{(^(fi23)), fi23, -62}
                for fi29 := 0; fi29 < 9; fi29++ {
                    switch "a" {
                    case "hello":
                        fmt.Println(p1, p1, fi23)
                    case "a":
                        v30 := GenS4{F0: 15, F1: 2147483647, F2: -1}
                        _ = v30
                    case "":
                        var v31 float64 = (3.14159 / (float64(1.5) - float64((0.0 / 1.0))))
                        _ = v31
                    default:
                        v27 = ((false || (false && (!v27))) && v27)
                    }
                    var v32 uint32 = 2
                    fmt.Println(v27)
                    _ = fi29
                    _ = v32
                }
                _ = v28
            } else {
                if (!v27) {
                    fmt.Println(p1, p1, p0)
                } else {
                    switch ("a" + "foo bar") {
                    case "hello":
                        var v33 uint16 = (p0 - (p0 >> ((703) & 63)))
                        _ = v33
                    case "":
                        var v34 int64 = 9223372036854775807
                        _ = v34
                    }
                    p1 = (^(p1))
                    for fi35 := 0; fi35 < 8; fi35++ {
                        var v36 uint8 = (p1 - p1)
                        v27 = v27
                        v37 := []int{42}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
                var v38 float32 = (-2.25 + 1e10)
                _ = v38
            }
            _ = v27
        } else {
            for fi39 := 0; fi39 < 5; fi39++ {
                v40 := []int{}
                _ = fi39
                _ = v40
            }
        }
        _ = fi23
    }
    p1 -= 127
    _ = p0
    _ = p1
    return ((!true) || false)
}

func main() {
    var v1 string = "a"
    var v2 uint32 = 0
    v3 := make(map[string]int)
    var v4 float64 = (0.0 / (float64(1e10) - 1.5))
    var v5 string = v1
    v6 := make(map[uint32]uint32)
    switch 0 {
    case -2147483648:
        if (-1 != 32767) {
            switch (v5 + (v1 + "hello")) {
            case "foo bar":
                v41 := []int{-851, -1}
                _ = v41
            case "\n\t":
                for fi42 := 0; fi42 < 7; fi42++ {
                    for fi43 := 0; fi43 < 2; fi43++ {
                        var v44 int16 = -32768
                        v44 -= (v44 | v44)
                        _ = fi43
                        _ = v44
                    }
                    v45 := GenS5{F0: 42}
                    for fi46 := 0; fi46 < 12; fi46++ {
                        v47 := []int{fi46, (^(fi42))}
                        if true {
                            v2 += v2
                            var v48 uint16 = uint16(v2)
                            var v49 float32 = 0.1
                            _ = v48
                            _ = v49
                        } else {
                            v5 = (("" + v5) + "a")
                            if false {
                                v5 = v5
                                v50 := GenS4{F0: 51, F1: -2147483648, F2: 957}
                                _ = len(v6)
                                switch (fi42 ^ fi46) {
                                case 32767:
                                    var v51 string = (v1 + v1)
                                    _ = v51
                                case -2147483648:
                                    switch fi42 {
                                    case 786:
                                        v52 := v3["foo bar"]
                                        _ = v52
                                    case 4294967295:
                                        if false {
                                            v3["a"] = fi42
                                            v53 := GenS5{F0: 32767}
                                            fmt.Println(fi46, v2, fi42)
                                            _ = v53
                                        }
                                    default:
                                        v54 := GenS6{F0: -55, F1: 0, F2: (fi42 / ((fi46) | 1))}
                                        _ = v54
                                    }
                                case -28:
                                    fmt.Println(fi42)
                                default:
                                    v55 := GenS7{F0: "hello", F1: 0}
                                    _ = v55
                                }
                                var v56 string = "a"
                                _ = v50
                                _ = v56
                            }
                            var v57 int32 = 0
                            _ = v57
                        }
                        if (1 <= -451) {
                            var v58 int16 = 1
                            v5 = (("a" + ((v5 + "") + "")) + v1)
                            switch (v1 + "a") {
                            case "a":
                                v58 -= v58
                            case "hello":
                                if true {
                                    var v59 uint = 711
                                    v4 -= v4
                                    fmt.Println(v1, v58)
                                    v60 := GenS6{F0: -1, F1: 276, F2: (fi46 << ((42) & 63))}
                                    _ = v59
                                    _ = v60
                                } else {
                                    var v61 int64 = -911
                                    _ = v61
                                }
                            case "\n\t":
                                v62 := GenS1{F0: 254}
                                _ = v62
                            default:
                                fmt.Println(fi46, fi42)
                            }
                            _ = v58
                        }
                        v63 := GenS8{F0: -32768}
                        _ = fi46
                        _ = v47
                        _ = v63
                    }
                    _ = fi42
                    _ = v45
                }
            }
            var v64 float64 = (v4 / v4)
            v64 -= 1.5
            _ = v64
        } else {
            if (((82 >= 65535) || (0 > 65535)) || ((^(v2)) == 741)) {
                var v65 int32 = -2147483648
                _ = v65
            }
            v66 := GenS3{F0: 18446744073709551615, F1: 0, F2: 42}
            _ = v66
        }
    }
    var v67 string = ("hello" + "foo bar")
    switch 1 {
    case -9223372036854775808:
        for fi68 := 0; fi68 < 20; fi68++ {
            v69 := []int{(fi68 + fi68), -241}
            _ = fi68
            _ = v69
        }
    default:
        v70 := GenS5{F0: 1}
        _ = v70
    }
    v71 := []int{32767, 9223372036854775807, -9223372036854775808}
    if (((0.1 + (0.1 + (1.5 + 1.5))) + 1e10) == 3.14159) {
        for fi72 := 0; fi72 < 8; fi72++ {
            var v73 int = fi72
            v2 -= (v2 >> ((445) & 63))
            _ = fi72
            _ = v73
        }
        v67 = v1
    } else {
        var v74 float64 = v4
        switch v1 {
        case "foo bar":
            var v75 int32 = 363
            _ = v75
        case "":
            v4 -= v4
        }
        v76 := GenS8{F0: 104}
        _ = v74
        _ = v76
    }
    var v77 float64 = -2.25
    for fi78 := 0; fi78 < 6; fi78++ {
        v79 := GenS1{F0: 45}
        _ = fi78
        _ = v79
    }
    v80 := GenS1{F0: 7}
    for fi81 := 0; fi81 < 7; fi81++ {
        fmt.Println(v77)
        _ = fi81
    }
    for fi82 := 0; fi82 < 4; fi82++ {
        fmt.Println(v67)
        v77 = 0.1
        _ = fi82
    }
    v4 = v77
    v1 = (v5 + v67)
    switch v67 {
    case "foo bar":
        v2 -= v2
    case "hello":
        v83 := v6[v2]
        _ = v83
    case "":
        for fi84 := 0; fi84 < 3; fi84++ {
            for fi85 := 0; fi85 < 3; fi85++ {
                _ = len(v3)
                var v86 uint = 711
                _ = fi85
                _ = v86
            }
            if (-2.25 < (0.0 / 1.5)) {
                var v87 bool = (176 <= 32767)
                _ = v87
            } else {
                if ((v77 + v4) >= float64(float32((v4 / 1e10)))) {
                    v77 += v4
                    var v88 uint32 = (v2 / (((^(v2))) | 1))
                    switch (v1 + ((("\n\t" + "a") + v67) + v5)) {
                    case "\n\t":
                        v1 = "\n\t"
                    }
                    _ = v88
                }
                var v89 int8 = 44
                _ = v89
            }
            _ = fi84
        }
    default:
        v5 = v5
    }
    switch 2147483647 {
    case 2147483647:
        for fi90 := 0; fi90 < 12; fi90++ {
            v67 = ((v1 + v67) + v1)
            fmt.Println(v67)
            _ = fi90
        }
    case 962:
        v71 = append(v71, 511)
    default:
        delete(v3, v67)
    }
    for fi91 := 0; fi91 < 6; fi91++ {
        fmt.Println(v1)
        v92 := GenS3{F0: 50, F1: 65535, F2: 848}
        v93 := GenS3{F0: 1, F1: 2, F2: 42}
        _ = fi91
        _ = v92
        _ = v93
    }
    v6[v2] = v2
    for fi94 := 0; fi94 < 17; fi94++ {
        if (!((0 != 127) && false)) {
            var v95 uint32 = (^(v2))
            switch "" {
            case "\n\t":
                fmt.Println(v1)
            default:
                if (1e10 < float64((((((1.5 * 1.5) + (1e10 + (-2.25 * -2.25))) + 1.5) / (-2.25 - float32(float64((0.0 - 0.1))))) / ((0.1 * 0.1) + 3.14159)))) {
                    var v96 int8 = -128
                    v95 += 65535
                    _ = v96
                }
            }
            var v97 uint32 = v2
            _ = v95
            _ = v97
        } else {
            if true {
                var v98 uint = 18446744073709551615
                _ = v98
            } else {
                if true {
                    var v99 uint32 = (v2 % (((v2 + 2147483647)) | 1))
                    _ = v99
                } else {
                    if (167 == 0) {
                        v100 := GenS9{F0: 255, F1: v67}
                        if true {
                            if true {
                                v71 = append(v71, fi94)
                            }
                            if false {
                                var v101 uint64 = 925
                                _ = v101
                            } else {
                                v67 = (v5 + "a")
                                for fi102 := 0; fi102 < 2; fi102++ {
                                    if len(v71) > 0 {
                                        v71[0] = (fi102 - (^(fi94)))
                                    }
                                    _ = len(v3)
                                    _ = fi102
                                }
                            }
                        }
                        if (0 >= -521) {
                            for fi103 := 0; fi103 < 4; fi103++ {
                                v1 = (("" + ("foo bar" + "\n\t")) + "")
                                fmt.Println(v77, v1, v2)
                                _ = fi103
                            }
                            fmt.Println(fi94)
                        } else {
                            var v104 uint32 = (v2 & v2)
                            var v105 int16 = -32768
                            _ = v104
                            _ = v105
                        }
                        _ = v100
                    }
                    for fi106 := 0; fi106 < 1; fi106++ {
                        v107 := GenS1{F0: 41}
                        if (0 < 42) {
                            var v108 uint32 = 4294967295
                            _ = v108
                        }
                        switch v67 {
                        case "hello":
                            var v109 float32 = 1e10
                            switch v1 {
                            case "":
                                if true {
                                    var v110 string = "a"
                                    v111 := GenS10{F0: (fi106 + 117)}
                                    if false {
                                        if (false && (!false)) {
                                            if len(v71) > 0 {
                                                v71[0] = fi94
                                            }
                                            var v112 uint16 = 655
                                            v109 = 1e10
                                            _ = v112
                                        }
                                        v113 := GenS6{F0: 1, F1: 2, F2: fi106}
                                        v114 := GenS1{F0: 255}
                                        _ = v113
                                        _ = v114
                                    }
                                    _ = v110
                                    _ = v111
                                }
                            case "a":
                                v115 := GenS11{F0: true}
                                _ = v115
                            }
                            if (-487 <= -563) {
                                if false {
                                    var v116 float32 = 1e10
                                    if (!(true || true)) {
                                        v116 = 0.1
                                        fmt.Println(v109, v5, v5)
                                        var v117 float32 = v109
                                        _ = v117
                                    }
                                    v118 := GenS12{F0: v2}
                                    switch ("\n\t" + v5) {
                                    case "foo bar":
                                        switch (v5 + (v1 + v1)) {
                                        case "a":
                                            var v119 uint = 253
                                            _ = v119
                                        case "foo bar":
                                            fmt.Println(v109, v4, v5)
                                        }
                                    case "\n\t":
                                        var v120 uint32 = v2
                                        _ = v120
                                    case "hello":
                                        var v121 float64 = (1e10 - v77)
                                        _ = v121
                                    }
                                    _ = v116
                                    _ = v118
                                }
                            } else {
                                switch fi94 {
                                case -1:
                                    v77 -= 0.0
                                case 9223372036854775807:
                                    fmt.Println(v5, fi94, v1)
                                case 0:
                                    if (3.14159 == v109) {
                                        var v122 int32 = 0
                                        v5 = v1
                                        v123 := GenS10{F0: fi94}
                                        _ = v122
                                        _ = v123
                                    } else {
                                        if (!true) {
                                            v1 = v5
                                        }
                                        v77 = 0.1
                                        v2 += (v2 % ((v2) | 1))
                                    }
                                default:
                                    if true {
                                        v67 = "hello"
                                        switch 4294967295 {
                                        case -2147483648:
                                            v124 := GenS13{F0: -24, F1: -490}
                                            _ = v124
                                        case 956:
                                            if (((true || ((183 < 731) && false)) || false) || true) {
                                                if ((v4 / 0.1) < v4) {
                                                    if true {
                                                        if false {
                                                            if (true && false) {
                                                                if (v109 > ((float32(float64(-2.25)) + 0.1) * v109)) {
                                                                    if false {
                                                                        if false {
                                                                            switch v1 {
                                                                            case "foo bar":
                                                                                v109 = v109
                                                                            case "hello":
                                                                                switch fi106 {
                                                                                case 178:
                                                                                    v109 += v109
                                                                                case -434:
                                                                                    if false {
                                                                                        var v125 uint8 = 239
                                                                                        _ = v125
                                                                                    }
                                                                                case 184:
                                                                                    var v126 int16 = 1000
                                                                                    _ = v126
                                                                                }
                                                                            case "a":
                                                                                v127 := GenS8{F0: 42}
                                                                                _ = v127
                                                                            default:
                                                                                if ((!false) || false) {
                                                                                    if true {
                                                                                        fmt.Println(v2)
                                                                                    } else {
                                                                                        var v128 string = v67
                                                                                        var v129 float32 = 0.0
                                                                                        if true {
                                                                                            if ("" >= v67) {
                                                                                                var v130 uint16 = 818
                                                                                                if (!(!false)) {
                                                                                                    var v131 int32 = -2147483648
                                                                                                    _ = v131
                                                                                                } else {
                                                                                                    v2 = 0
                                                                                                }
                                                                                                if (-9223372036854775808 >= 0) {
                                                                                                    fmt.Println(v129, v130)
                                                                                                }
                                                                                                _ = v130
                                                                                            }
                                                                                            v77 -= 1.5
                                                                                            v132 := v3[v67]
                                                                                            _ = v132
                                                                                        }
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                    }
                                                                                    var v133 int64 = 9223372036854775807
                                                                                    _ = v133
                                                                                } else {
                                                                                    if false {
                                                                                        v1 = v5
                                                                                    }
                                                                                    var v134 uint64 = 18446744073709551615
                                                                                    _ = v134
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v5 = ""
                                                                        if false {
                                                                            v135 := GenS14{F0: 245, F1: -117, F2: 159}
                                                                            var v136 bool = false
                                                                            _ = v135
                                                                            _ = v136
                                                                        } else {
                                                                            v137 := GenS4{F0: 128, F1: -1, F2: -147}
                                                                            var v138 uint8 = 0
                                                                            _ = v137
                                                                            _ = v138
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                var v139 int = 501
                                                                var v140 string = v1
                                                                switch v139 {
                                                                case 1:
                                                                    fmt.Println(v139)
                                                                }
                                                                _ = v139
                                                                _ = v140
                                                            }
                                                            var v141 float64 = 1.5
                                                            v141 -= v141
                                                            if len(v71) >= 2 {
                                                                v71 = v71[1:len(v71)]
                                                            }
                                                            switch v67 {
                                                            case "a":
                                                                fmt.Println(v4, v77)
                                                            case "":
                                                                v142 := GenS7{F0: v1, F1: 42}
                                                                _ = v142
                                                            case "hello":
                                                                switch v5 {
                                                                case "hello":
                                                                    v143 := GenS5{F0: -1}
                                                                    _ = v143
                                                                }
                                                            default:
                                                                fmt.Println(v109, v5, v4)
                                                            }
                                                            fmt.Println(v109, fi94, v5)
                                                            v144 := GenS13{F0: -128, F1: 42}
                                                            _ = v141
                                                            _ = v144
                                                        }
                                                        v1 = (v67 + v5)
                                                    } else {
                                                        if true {
                                                            var v145 uint32 = (v2 | 122)
                                                            v145 += v145
                                                            _ = v145
                                                        } else {
                                                            fmt.Println(v77)
                                                        }
                                                        var v146 uint = 255
                                                        _ = v146
                                                    }
                                                    var v147 int = 32768
                                                    _ = v147
                                                }
                                            }
                                        case -503:
                                            v77 -= v4
                                        default:
                                            var v148 int = fi106
                                            _ = v148
                                        }
                                    }
                                }
                                var v149 uint16 = 0
                                var v150 uint16 = genFn0((42 > v149))
                                _ = v149
                                _ = v150
                            }
                            v67 = "\n\t"
                            if (fi106 > -474) {
                                switch (fi106 & -538) {
                                case -573:
                                    v151 := GenS15{F0: 356, F1: 128, F2: (36 < 9223372036854775807)}
                                    _ = v151
                                case 667:
                                    v6[2147483648] = (^(v2))
                                }
                                if false {
                                    v2 += (v2 | 2)
                                    v67 = v1
                                } else {
                                    v1 = v5
                                    _ = len(v3)
                                }
                                v152 := GenS1{F0: 28}
                                if true {
                                    v153 := GenS3{F0: 0, F1: 2, F2: 4294967295}
                                    var v154 int64 = 42
                                    _ = v153
                                    _ = v154
                                }
                                _ = v152
                            } else {
                                if false {
                                    v6[v2] = 65535
                                }
                            }
                            _ = v109
                        case "a":
                            v5 = ("a" + ("hello" + ("hello" + (v1 + v5))))
                        case "":
                            _ = len(v3)
                        default:
                            v155 := GenS12{F0: v2}
                            _ = v155
                        }
                        _ = fi106
                        _ = v107
                    }
                }
                var v156 string = "foo bar"
                _ = v156
            }
        }
        var v157 int = 9223372036854775807
        _ = fi94
        _ = v157
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v71
    _ = v77
    _ = v80
}
