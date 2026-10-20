package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint16; F2 string }
type GenS6 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 uint16 }
type GenS12 struct { F0 int; F1 int32; F2 int8 }
type GenS13 struct { F0 uint8; F1 string }
type GenS14 struct { F0 uint64; F1 int8; F2 string }
type GenS15 struct { F0 int16; F1 int64; F2 int64 }
func genFn0() bool {
    var v6 uint8 = 0
    var v7 float32 = 1.5
    _ = v6
    _ = v7
    return ((true || true) || true)
}
func genFn1(p0 bool, p1 int64) int16 {
    switch 32767 {
    case -425:
        p0 = (0 == 42)
    case 9223372036854775807:
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 bool = true
            var v10 int8 = -128
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        v11 := GenS1{F0: 32767, F1: (-(p1))}
        _ = v11
    }
    v12 := GenS2{F0: 131, F1: 2, F2: (2147483648 == 2147483647)}
    var v13 uint8 = 1
    v13 = 42
    v14 := GenS2{F0: 222, F1: 65535, F2: (((-1 > 1) && false) && false)}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.1
    v3 := []string{("" + "hello"), "a"}
    var v4 int32 = 2147483647
    var v5 int = -252
    if false {
        var v15 int32 = (v4 / ((v4) | 1))
        var v16 float64 = v2
        var v17 uint = 1
        _ = v15
        _ = v16
        _ = v17
    } else {
        var v18 uint = 0
        if false {
            var v19 uint32 = 65535
            _ = v19
        }
        _ = v18
    }
    fmt.Println(v5)
    var v20 uint16 = 42
    if (-2147483648 > v4) {
        v3 = append(v3, "")
    } else {
        if true {
            v4 -= v4
        } else {
            fmt.Println(v2, v20, v4)
            switch "hello" {
            case "\n\t":
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS1{F0: 1, F1: 1}
                    _ = fi21
                    _ = v22
                }
            case "hello":
                for fi23 := 0; fi23 < 1; fi23++ {
                    if (255 < 208) {
                        if false {
                            for fi24 := 0; fi24 < 6; fi24++ {
                                v25 := GenS3{F0: 255}
                                _ = fi24
                                _ = v25
                            }
                            v4 = -900
                        } else {
                            var v26 uint8 = 255
                            switch fi23 {
                            case 0:
                                v3 = append(v3, "")
                            }
                            for fi27 := 0; fi27 < 9; fi27++ {
                                if false {
                                    var v28 int64 = 9223372036854775807
                                    fmt.Println(v26, v20, fi27)
                                    fmt.Println(v4, v5, v26)
                                    _ = v28
                                } else {
                                    v29 := GenS4{F0: 838}
                                    if true {
                                        if (2 < (v20 ^ (v20 << ((uint(v20)) & 63)))) {
                                            switch (fi27 | fi23) {
                                            case 918:
                                                v2 = (float64((0.0 - 0.0)) / ((1e10 * float64(-2.25)) + ((((v2 - 3.14159) + (0.1 * 0.1)) + 3.14159) * -2.25)))
                                            case -220:
                                                v1[868] = 0
                                            case -538:
                                                v3 = append(v3, ("hello" + "\n\t"))
                                            default:
                                                var v30 uint8 = v26
                                                _ = v30
                                            }
                                            fmt.Println(v20, fi27, v20)
                                            var v31 float32 = (-2.25 / -2.25)
                                            _ = v31
                                        }
                                    }
                                    _ = v29
                                }
                                var v32 float64 = 0.1
                                if true {
                                    var v33 uint = 65535
                                    _ = v33
                                } else {
                                    var v34 float32 = (3.14159 * -2.25)
                                    var v35 int64 = 526
                                    v1[1] = uint32(v26)
                                    _ = v34
                                    _ = v35
                                }
                                if len(v3) > 0 {
                                    v3[0] = ("\n\t" + (("foo bar" + (("a" + (("a" + "hello") + ("a" + "\n\t"))) + "foo bar")) + "foo bar"))
                                }
                                _ = fi27
                                _ = v32
                            }
                            _ = v26
                        }
                        var v36 string = ("a" + "foo bar")
                        _ = v36
                    }
                    if ((!(1 <= 144)) || (0.0 != ((v2 - float64(((float32(float64(1e10)) / 1.0) - float32(float64(0.1))))) - v2))) {
                        if (!(2 > 755)) {
                            delete(v1, 1)
                            v37 := GenS1{F0: -968, F1: -9223372036854775808}
                            var v38 bool = (false && ((-2.25 * (float64(3.14159) * 0.0)) < v2))
                            _ = v37
                            _ = v38
                        } else {
                            switch "hello" {
                            case "":
                                v4 -= v4
                            }
                        }
                        v3 = append(v3, "")
                    }
                    v5 += v5
                    _ = fi23
                }
            case "":
                var v39 uint16 = (v20 + (^(v20)))
                _ = v39
            default:
                v2 -= v2
            }
        }
        switch ("" + "") {
        case "foo bar":
            switch (-(v5)) {
            case 9223372036854775807:
                _ = len(v3)
            case 32768:
                v20 -= 332
            case 128:
                for fi40 := 0; fi40 < 1; fi40++ {
                    var v41 int = (fi40 - 2147483647)
                    _ = fi40
                    _ = v41
                }
            default:
                for fi42 := 0; fi42 < 3; fi42++ {
                    switch ("" + "\n\t") {
                    case "a":
                        var v43 uint16 = v20
                        _ = v43
                    }
                    v3 = append(v3, (("foo bar" + "") + "\n\t"))
                    _ = fi42
                }
            }
        case "hello":
            for fi44 := 0; fi44 < 18; fi44++ {
                if (true || (32767 < 42)) {
                    v4 = (v4 * v4)
                } else {
                    var v45 uint32 = 65535
                    _ = len(v1)
                    v46 := GenS1{F0: 42, F1: 42}
                    _ = v45
                    _ = v46
                }
                v4 = 42
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(fi44)
                }
                _ = fi44
            }
        case "\n\t":
            if true {
                for fi47 := 0; fi47 < 13; fi47++ {
                    var v48 bool = false
                    v1[484] = 2147483648
                    switch (fi47 | fi47) {
                    case 127:
                        v49 := GenS3{F0: 0}
                        _ = v49
                    case 135:
                        for fi50 := 0; fi50 < 8; fi50++ {
                            fmt.Println(v5, v20)
                            _ = fi50
                        }
                    case 4294967295:
                        for fi51 := 0; fi51 < 5; fi51++ {
                            v52 := GenS2{F0: 2, F1: v20, F2: false}
                            _ = fi51
                            _ = v52
                        }
                    }
                    _ = fi47
                    _ = v48
                }
            }
        default:
            var v53 int64 = 0
            _ = v53
        }
        if false {
            var v54 string = "foo bar"
            _ = v54
        }
    }
    v1[717] = 347
    v4 += v4
    var v55 uint8 = 128
    v20 = v20
    var v56 int64 = 123
    switch "foo bar" {
    case "foo bar":
        var v57 uint16 = v20
        _ = v57
    default:
        if (0 != 32767) {
            for fi58 := 0; fi58 < 17; fi58++ {
                var v59 int16 = genFn1(false, int64(v4))
                switch (v5 + (v5 - 42)) {
                case 32767:
                    switch v5 {
                    case 32767:
                        for fi60 := 0; fi60 < 3; fi60++ {
                            v55 -= (v55 % (((^(v55))) | 1))
                            if true {
                                var v61 bool = (-128 > 100)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v62 := GenS5{F0: (-(v56)), F1: 413, F2: "hello"}
                                _ = v61
                                _ = v62
                            } else {
                                if true {
                                    var v63 int32 = int32(v55)
                                    v64 := GenS4{F0: uint32(v55)}
                                    _ = v63
                                    _ = v64
                                }
                                var v65 int16 = genFn1(false, (v56 ^ v56))
                                _ = v65
                            }
                            _ = fi60
                        }
                    default:
                        v66 := GenS5{F0: -1, F1: 547, F2: "foo bar"}
                        _ = v66
                    }
                case 42:
                    var v67 uint8 = v55
                    _ = v67
                default:
                    v68 := GenS5{F0: v56, F1: (v20 ^ 32767), F2: ""}
                    _ = v68
                }
                delete(v1, 841)
                _ = fi58
                _ = v59
            }
        }
    }
    for fi69 := 0; fi69 < 11; fi69++ {
        fmt.Println(v56, v55, fi69)
        _ = fi69
    }
    if false {
        var v70 uint16 = v20
        v71 := GenS3{F0: (v55 | 2)}
        _ = v70
        _ = v71
    }
    var v72 float64 = (0.1 - ((v2 * (0.0 / (1.5 - v2))) - v2))
    var v73 uint64 = 2
    var v74 uint16 = v20
    for fi75 := 0; fi75 < 6; fi75++ {
        if len(v3) > 0 {
            v3[0] = ("hello" + "foo bar")
        }
        var v76 int64 = 1
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 20; fi77++ {
        var v78 uint64 = (v73 >> ((345) & 63))
        v74 = 32767
        if (!true) {
            if true {
                if false {
                    fmt.Println(fi77, v56)
                }
            }
            fmt.Println(v2)
        } else {
            switch ("\n\t" + ((("foo bar" + "foo bar") + ("foo bar" + ("\n\t" + ("\n\t" + ("" + ("a" + (((("" + (("" + "\n\t") + ("" + ""))) + "") + "\n\t") + "foo bar"))))))) + (((((("\n\t" + (("a" + (("" + "\n\t") + "a")) + ("foo bar" + ("" + (((("" + "\n\t") + "hello") + "") + ("foo bar" + "a")))))) + "a") + "a") + "hello") + "\n\t") + ""))) {
            case "hello":
                v73 -= v73
            case "":
                v79 := GenS3{F0: uint8(v20)}
                _ = v79
            default:
                fmt.Println(v2, v5)
            }
        }
        _ = fi77
        _ = v78
    }
    if (3.14159 < (0.0 * 1e10)) {
        switch "foo bar" {
        case "a":
            var v80 uint64 = v73
            _ = v80
        default:
            for fi81 := 0; fi81 < 20; fi81++ {
                var v82 float64 = v72
                var v83 uint8 = (v55 ^ 35)
                if ((873 < 0) && true) {
                    if false {
                        for fi84 := 0; fi84 < 11; fi84++ {
                            _ = len(v3)
                            var v85 int16 = -32768
                            _ = fi84
                            _ = v85
                        }
                        if false {
                            var v86 int = v5
                            _ = v86
                        } else {
                            v2 = v82
                        }
                        v55 += (^(v83))
                    } else {
                        v83 -= v55
                    }
                    v87 := GenS3{F0: v55}
                    for fi88 := 0; fi88 < 10; fi88++ {
                        v56 -= (-(v56))
                        v89 := GenS6{F0: -32768, F1: 302, F2: (v73 - (v73 << ((18446744073709551615) & 63)))}
                        _ = fi88
                        _ = v89
                    }
                    _ = v87
                }
                _ = fi81
                _ = v82
                _ = v83
            }
        }
        var v90 int = (v5 | -9223372036854775808)
        switch ("a" + "\n\t") {
        case "":
            if false {
                if false {
                    if true {
                        var v91 uint = uint(v55)
                        var v92 uint = uint(v20)
                        _ = v91
                        _ = v92
                    } else {
                        switch v5 {
                        case -1:
                            v55 -= v55
                        case 3:
                            switch 32768 {
                            case 145:
                                _ = len(v1)
                            case 1:
                                switch ("a" + "hello") {
                                case "foo bar":
                                    var v93 uint32 = 65535
                                    _ = v93
                                case "hello":
                                    v72 += v2
                                case "a":
                                    switch v5 {
                                    case -1:
                                        if (false && (-660 > (v90 >> ((42) & 63)))) {
                                            fmt.Println(v5)
                                        } else {
                                            var v94 float64 = (0.0 - v72)
                                            for fi95 := 0; fi95 < 5; fi95++ {
                                                var v96 int8 = int8(v90)
                                                _ = fi95
                                                _ = v96
                                            }
                                            var v97 int64 = v56
                                            _ = v94
                                            _ = v97
                                        }
                                    }
                                default:
                                    v74 = (v74 + 253)
                                }
                            case -1:
                                _ = len(v1)
                            default:
                                var v98 int8 = 127
                                _ = v98
                            }
                        case 639:
                            for fi99 := 0; fi99 < 1; fi99++ {
                                fmt.Println(v5)
                                _ = fi99
                            }
                        default:
                            v72 = (v2 + 1e10)
                        }
                    }
                    var v100 uint64 = uint64(v74)
                    var v101 int64 = (v56 | v56)
                    _ = v100
                    _ = v101
                } else {
                    if (-9223372036854775808 != -1) {
                        v5 -= v90
                        v74 -= (v20 & 1)
                    } else {
                        if (!true) {
                            v4 = (^(v4))
                            for fi102 := 0; fi102 < 16; fi102++ {
                                v103 := GenS6{F0: 986, F1: 31, F2: v73}
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                                _ = fi102
                                _ = v103
                            }
                            switch v5 {
                            case 1:
                                v104 := GenS7{F0: -2147483648}
                                _ = v104
                            case -2147483648:
                                v105 := GenS8{F0: 231, F1: (^(v20)), F2: -128}
                                _ = v105
                            }
                        }
                        if true {
                            v106 := GenS2{F0: 153, F1: (v74 % ((65535) | 1)), F2: (0.0 < (1.5 / 3.14159))}
                            var v107 string = (((("foo bar" + ("a" + "a")) + "hello") + ("a" + "\n\t")) + "hello")
                            var v108 int32 = 0
                            _ = v106
                            _ = v107
                            _ = v108
                        }
                    }
                    if (-238 <= 0) {
                        for fi109 := 0; fi109 < 14; fi109++ {
                            var v110 int8 = 1
                            var v111 int = 4294967295
                            if false {
                                v72 -= v2
                                for fi112 := 0; fi112 < 16; fi112++ {
                                    if false {
                                        v2 += v2
                                        v110 -= (v110 | v110)
                                        var v113 uint8 = (v55 * 5)
                                        _ = v113
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi112
                                }
                                for fi114 := 0; fi114 < 14; fi114++ {
                                    delete(v1, 2147483647)
                                    _ = fi114
                                }
                            }
                            _ = fi109
                            _ = v110
                            _ = v111
                        }
                        if false {
                            v115 := GenS2{F0: 91, F1: v74, F2: true}
                            _ = v115
                        } else {
                            v116 := GenS9{F0: 708, F1: v73, F2: (v4 / ((0) | 1))}
                            v73 = v73
                            _ = v116
                        }
                        v20 += v20
                    } else {
                        for fi117 := 0; fi117 < 15; fi117++ {
                            if (2147483648 > uint32(v73)) {
                                fmt.Println(v73, v20)
                                v118 := GenS10{F0: int(v56)}
                                _ = v118
                            }
                            switch "foo bar" {
                            case "a":
                                var v119 uint16 = (v20 + v20)
                                _ = v119
                            default:
                                var v120 int16 = 487
                                _ = v120
                            }
                            var v121 int8 = -47
                            _ = fi117
                            _ = v121
                        }
                        var v122 uint8 = 128
                        _ = v122
                    }
                    switch v5 {
                    case -9223372036854775808:
                        if (1 >= 397) {
                            for fi123 := 0; fi123 < 7; fi123++ {
                                fmt.Println(v55, v73)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi123
                            }
                            if (1e10 > (-2.25 + 3.14159)) {
                                v55 += (v55 % ((1) | 1))
                                var v124 uint64 = (v73 + v73)
                                if (("" + ("hello" + "\n\t")) <= "") {
                                    var v125 uint = 1
                                    if (0 >= v56) {
                                        for fi126 := 0; fi126 < 13; fi126++ {
                                            v2 = (v2 + v2)
                                            v2 -= 3.14159
                                            if (!(true && ((1.5 * (0.0 / (float32(1e10) * 0.1))) > (1.5 + 1.5)))) {
                                                switch 32768 {
                                                case 2147483647:
                                                    fmt.Println(fi126, v5, fi126)
                                                case 9223372036854775807:
                                                    v4 = -2147483648
                                                case 42:
                                                    var v127 uint8 = v55
                                                    _ = v127
                                                }
                                                if false {
                                                    v128 := GenS5{F0: v56, F1: 2, F2: ""}
                                                    v74 = v74
                                                    _ = v128
                                                }
                                            } else {
                                                var v129 string = "hello"
                                                if false {
                                                    v2 = (float64(float32(v2)) * ((((0.1 / ((1e10 * v72) / -2.25)) - 1.5) - float64((0.1 / 1.5))) * v2))
                                                    v130 := GenS5{F0: v56, F1: v74, F2: "a"}
                                                    _ = v130
                                                }
                                                v1[0] = 42
                                                _ = v129
                                            }
                                            _ = fi126
                                        }
                                        if false {
                                            if true {
                                                var v131 bool = true
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v131
                                            } else {
                                                if (v2 < v2) {
                                                    v90 = -1
                                                    if (42 != 2147483648) {
                                                        if true {
                                                            if (false && true) {
                                                                var v132 float64 = (v72 / 3.14159)
                                                                _ = v132
                                                            } else {
                                                                fmt.Println(v20, v20)
                                                                for fi133 := 0; fi133 < 20; fi133++ {
                                                                    fmt.Println(v90)
                                                                    v134 := GenS11{F0: false, F1: (^(v74))}
                                                                    for fi135 := 0; fi135 < 8; fi135++ {
                                                                        v136 := GenS12{F0: 725, F1: int32(v124), F2: 83}
                                                                        v124 = (v73 ^ v73)
                                                                        var v137 int64 = int64(fi135)
                                                                        _ = fi135
                                                                        _ = v136
                                                                        _ = v137
                                                                    }
                                                                    _ = fi133
                                                                    _ = v134
                                                                }
                                                            }
                                                            v5 = v90
                                                        } else {
                                                            v74 += v20
                                                        }
                                                        v138 := GenS13{F0: 230, F1: ("\n\t" + "foo bar")}
                                                        _ = v138
                                                    } else {
                                                        v139 := GenS6{F0: 32767, F1: 401, F2: (v124 << ((v125) & 63))}
                                                        for fi140 := 0; fi140 < 9; fi140++ {
                                                            var v141 bool = true
                                                            v142 := GenS13{F0: v55, F1: "foo bar"}
                                                            v143 := GenS14{F0: (v73 % (((^(v124))) | 1)), F1: 100, F2: "a"}
                                                            _ = fi140
                                                            _ = v141
                                                            _ = v142
                                                            _ = v143
                                                        }
                                                        v144 := GenS2{F0: (^(v55)), F1: v74, F2: true}
                                                        _ = v139
                                                        _ = v144
                                                    }
                                                    fmt.Println(v124, v124, v90)
                                                } else {
                                                    v145 := GenS14{F0: v73, F1: 1, F2: "\n\t"}
                                                    if ((1e10 * 0.0) != 1e10) {
                                                        for fi146 := 0; fi146 < 12; fi146++ {
                                                            if (4294967295 < 65535) {
                                                                for fi147 := 0; fi147 < 16; fi147++ {
                                                                    v55 = v55
                                                                    _ = fi147
                                                                }
                                                                for fi148 := 0; fi148 < 6; fi148++ {
                                                                    v149 := v1[820]
                                                                    var v150 float32 = -2.25
                                                                    v4 -= (-(v4))
                                                                    _ = fi148
                                                                    _ = v149
                                                                    _ = v150
                                                                }
                                                            }
                                                            fmt.Println(v124)
                                                            _ = fi146
                                                        }
                                                        delete(v1, 2147483647)
                                                        for fi151 := 0; fi151 < 2; fi151++ {
                                                            if true {
                                                                if false {
                                                                    var v152 uint = v125
                                                                    switch fi151 {
                                                                    case 42:
                                                                        var v153 float64 = float64(0.0)
                                                                        _ = v153
                                                                    }
                                                                    _ = v152
                                                                } else {
                                                                    v1[2147483648] = uint32(v20)
                                                                    if (0 > (v55 ^ (v55 % ((204) | 1)))) {
                                                                        var v154 uint16 = v20
                                                                        _ = v154
                                                                    }
                                                                }
                                                                var v155 uint16 = 42
                                                                var v156 uint16 = (^(v20))
                                                                _ = v155
                                                                _ = v156
                                                            } else {
                                                                for fi157 := 0; fi157 < 8; fi157++ {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v158 int = 128
                                                                        _ = v158
                                                                    }
                                                                    var v159 string = "foo bar"
                                                                    _ = fi157
                                                                    _ = v159
                                                                }
                                                            }
                                                            v5 -= (-(v5))
                                                            _ = fi151
                                                        }
                                                    } else {
                                                        if ("hello" == "foo bar") {
                                                            v160 := GenS6{F0: 1000, F1: 4294967295, F2: v124}
                                                            fmt.Println(v2)
                                                            var v161 uint16 = (v74 / ((32767) | 1))
                                                            _ = v160
                                                            _ = v161
                                                        }
                                                        delete(v1, 0)
                                                        for fi162 := 0; fi162 < 13; fi162++ {
                                                            var v163 uint8 = v55
                                                            var v164 uint64 = (v73 | v124)
                                                            switch ("" + (("\n\t" + ("foo bar" + "hello")) + ("a" + "hello"))) {
                                                            case "\n\t":
                                                                var v165 uint32 = 42
                                                                _ = v165
                                                            default:
                                                                fmt.Println(v56)
                                                            }
                                                            _ = fi162
                                                            _ = v163
                                                            _ = v164
                                                        }
                                                    }
                                                    for fi166 := 0; fi166 < 17; fi166++ {
                                                        for fi167 := 0; fi167 < 8; fi167++ {
                                                            v168 := GenS9{F0: v20, F1: (^(v124)), F2: int32(v125)}
                                                            v72 = v72
                                                            fmt.Println(v20, v72)
                                                            _ = fi167
                                                            _ = v168
                                                        }
                                                        v90 -= (fi166 - int(v125))
                                                        _ = fi166
                                                    }
                                                    _ = v145
                                                }
                                                var v169 uint = (^(v125))
                                                _ = v169
                                            }
                                            if false {
                                                switch (v90 & 32767) {
                                                case 42:
                                                    for fi170 := 0; fi170 < 1; fi170++ {
                                                        var v171 uint8 = (^(v55))
                                                        _ = fi170
                                                        _ = v171
                                                    }
                                                case 2147483647:
                                                    var v172 string = (("" + "a") + "")
                                                    _ = v172
                                                }
                                                if false {
                                                    v90 = (v90 + v90)
                                                    for fi173 := 0; fi173 < 15; fi173++ {
                                                        if len(v3) > 0 {
                                                            v3[0] = ("foo bar" + "foo bar")
                                                        }
                                                        v125 = (v125 << ((919) & 63))
                                                        _ = fi173
                                                    }
                                                    var v174 float32 = ((0.1 + 1.5) / 0.1)
                                                    _ = v174
                                                }
                                            }
                                        }
                                        var v175 uint64 = 0
                                        _ = v175
                                    } else {
                                        fmt.Println(v4, v4, v5)
                                    }
                                    var v176 bool = (v4 < (-(v4)))
                                    _ = v125
                                    _ = v176
                                }
                                _ = v124
                            } else {
                                v2 += 1e10
                                delete(v1, 2147483647)
                            }
                        }
                    default:
                        _ = len(v1)
                    }
                }
                v55 += 2
            }
        }
        _ = v90
    } else {
        for fi177 := 0; fi177 < 3; fi177++ {
            for fi178 := 0; fi178 < 11; fi178++ {
                var v179 int32 = 108
                _ = fi178
                _ = v179
            }
            switch 889 {
            case -2147483648:
                v180 := GenS15{F0: 301, F1: v56, F2: v56}
                _ = v180
            case 32767:
                var v181 int16 = 1
                _ = v181
            case -843:
                v56 = (v56 % ((v56) | 1))
            }
            if len(v3) > 0 {
                v3[0] = ""
            }
            _ = fi177
        }
        if true {
            var v182 uint8 = (v55 ^ v55)
            var v183 uint8 = v182
            _ = v182
            _ = v183
        }
        var v184 int = (v5 - (v5 & 2147483647))
        _ = v184
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v72
    _ = v73
    _ = v74
}
