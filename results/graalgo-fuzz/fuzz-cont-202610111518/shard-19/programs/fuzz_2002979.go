package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int16; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 int8; F2 int32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 string }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 bool; F1 uint64 }
type GenS14 struct { F0 uint64; F1 string; F2 int32 }
func genFn0(p0 uint16, p1 uint16) bool {
    if false {
        var v6 uint32 = 163
        var v7 uint = 18446744073709551615
        _ = v6
        _ = v7
    } else {
        fmt.Println(p0, p1)
        p1 -= (p0 + p1)
        var v8 int8 = 0
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 uint64) string {
    p0 += (^(p0))
    var v9 uint16 = 947
    var v10 uint64 = (p0 >> ((550) & 63))
    _ = p0
    _ = v10
    _ = v9
    return "foo bar"
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = 0
    v3 := []string{"\n\t", "\n\t", v1}
    v4 := make(map[int]bool)
    var v5 float32 = 1e10
    if false {
        switch 42 {
        case -1:
            switch 823 {
            case 9223372036854775807:
                v2 += v2
            case 205:
                v5 = ((float32((0.0 - 0.1)) - v5) * 0.0)
            default:
                var v11 int = 311
                _ = v11
            }
        case -9223372036854775808:
            if true {
                if len(v3) > 0 {
                    v3[0] = ("hello" + v1)
                }
            } else {
                v12 := GenS1{F0: 2, F1: v2, F2: -1}
                for fi13 := 0; fi13 < 13; fi13++ {
                    v14 := GenS2{F0: fi13}
                    for fi15 := 0; fi15 < 14; fi15++ {
                        if true {
                            fmt.Println(v2, v5)
                            var v16 int8 = 5
                            v1 = v1
                            _ = v16
                        }
                        var v17 int32 = (v2 - 0)
                        _ = fi15
                        _ = v17
                    }
                    if true {
                        fmt.Println(v2)
                        v18 := GenS3{F0: 32767, F1: 32767, F2: v2}
                        var v19 uint = 0
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint8 = 127
                        if false {
                            if (18446744073709551615 >= 4294967295) {
                                if false {
                                    var v21 uint32 = 721
                                    if (!false) {
                                        if (!(("hello" != v1) && false)) {
                                            var v22 float32 = v5
                                            _ = v22
                                        }
                                        v3 = append(v3, (("foo bar" + (((((v1 + v1) + ("hello" + v1)) + "") + v1) + v1)) + "a"))
                                    } else {
                                        v2 -= (v2 / ((v2) | 1))
                                    }
                                    var v23 float64 = 0.1
                                    _ = v21
                                    _ = v23
                                }
                                v5 -= (-2.25 - (1e10 * 0.0))
                            } else {
                                v24 := GenS4{F0: false}
                                _ = v24
                            }
                            v2 -= v2
                            if (true && false) {
                                switch "hello" {
                                case "hello":
                                    for fi25 := 0; fi25 < 17; fi25++ {
                                        var v26 float64 = 1e10
                                        _ = fi25
                                        _ = v26
                                    }
                                case "foo bar":
                                    var v27 int16 = -32768
                                    _ = v27
                                case "":
                                    v28 := GenS5{F0: -5, F1: 4294967295, F2: 406}
                                    _ = v28
                                }
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    v20 -= v20
                                    _ = fi29
                                }
                            } else {
                                for fi30 := 0; fi30 < 4; fi30++ {
                                    v31 := GenS6{F0: 2, F1: 0}
                                    _ = fi30
                                    _ = v31
                                }
                                if (true || false) {
                                    v1 = v1
                                } else {
                                    if (false && false) {
                                        v32 := GenS3{F0: -90, F1: -1, F2: v2}
                                        v33 := GenS3{F0: 610, F1: 32767, F2: (v2 & v2)}
                                        for fi34 := 0; fi34 < 2; fi34++ {
                                            var v35 int = fi34
                                            _ = fi34
                                            _ = v35
                                        }
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v1 = (("a" + "a") + v1)
                                        var v36 float64 = -2.25
                                        _ = v36
                                    }
                                }
                                v37 := GenS7{F0: 255}
                                _ = v37
                            }
                        } else {
                            switch v1 {
                            case "\n\t":
                                v20 += (^(v20))
                            default:
                                fmt.Println(v2, v20)
                            }
                            v20 = v20
                        }
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = v20
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v12
            }
        case -643:
            for fi38 := 0; fi38 < 2; fi38++ {
                for fi39 := 0; fi39 < 7; fi39++ {
                    v1 = "foo bar"
                    _ = fi39
                }
                switch (fi38 & fi38) {
                case 127:
                    delete(v4, fi38)
                default:
                    var v40 int64 = -1
                    _ = v40
                }
                for fi41 := 0; fi41 < 9; fi41++ {
                    if true {
                        v5 += 0.1
                    }
                    switch (fi41 * -9223372036854775808) {
                    case -1:
                        switch fi41 {
                        case 2147483647:
                            if false {
                                v1 = v1
                            } else {
                                fmt.Println(fi41, v1, fi38)
                                var v42 uint16 = 42
                                _ = v42
                            }
                        case -1:
                            _ = len(v4)
                        }
                    case 127:
                        var v43 float64 = (0.1 + 1.5)
                        _ = v43
                    default:
                        if false {
                            if false {
                                delete(v4, 9223372036854775807)
                                v1 = "hello"
                                var v44 uint16 = 2
                                _ = v44
                            } else {
                                fmt.Println(v1, v1)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            var v45 float64 = 3.14159
                            if ((!((1 > 0) && ((false || false) && true))) || (592 > 1)) {
                                fmt.Println(fi38)
                            }
                            v45 = v45
                            _ = v45
                        }
                    }
                    _ = fi41
                }
                _ = fi38
            }
        }
        var v46 int16 = 0
        fmt.Println(v5, v46)
        _ = v46
    } else {
        if (-464 > 0) {
            switch -1 {
            case -743:
                var v47 bool = (!(true || false))
                _ = v47
            default:
                v48 := GenS6{F0: 929, F1: 1}
                _ = v48
            }
            var v49 int8 = 1
            fmt.Println(v49, v2)
            _ = v49
        }
        var v50 int = -622
        for fi51 := 0; fi51 < 10; fi51++ {
            v50 = 32768
            _ = fi51
        }
        _ = v50
    }
    for fi52 := 0; fi52 < 14; fi52++ {
        v53 := GenS8{F0: 1, F1: 100, F2: v2}
        switch v1 {
        case "":
            switch (v1 + ((v1 + v1) + v1)) {
            case "foo bar":
                for fi54 := 0; fi54 < 12; fi54++ {
                    var v55 int16 = 1000
                    var v56 uint = 0
                    _ = fi54
                    _ = v55
                    _ = v56
                }
            }
        case "foo bar":
            if false {
                _ = len(v4)
                for fi57 := 0; fi57 < 6; fi57++ {
                    v3 = append(v3, ("\n\t" + (((v1 + v1) + ((("foo bar" + "") + ((("hello" + v1) + ((v1 + v1) + v1)) + v1)) + v1)) + "\n\t")))
                    switch v1 {
                    case "\n\t":
                        if (-733 > -866) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            var v58 string = v1
                            if (-128 != -1) {
                                fmt.Println(fi57, fi52, v5)
                                v59 := GenS6{F0: 2, F1: 1}
                                v60 := GenS9{F0: 2147483647}
                                _ = v59
                                _ = v60
                            } else {
                                if false {
                                    var v61 int = -1
                                    if true {
                                        if (((1.5 / v5) * float32(float64(0.1))) > v5) {
                                            if len(v3) > 0 {
                                                v3[0] = v58
                                            }
                                            v2 = v2
                                        }
                                    } else {
                                        v62 := GenS3{F0: -32768, F1: -381, F2: -92}
                                        _ = v62
                                    }
                                    v58 = v1
                                    _ = v61
                                } else {
                                    if false {
                                        v63 := GenS2{F0: fi57}
                                        v3 = append(v3, v1)
                                        _ = v63
                                    } else {
                                        switch (fi57 % (((fi57 / (((fi57 & fi57)) | 1))) | 1)) {
                                        case 0:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case 4294967295:
                                            var v64 float32 = v5
                                            _ = v64
                                        case 944:
                                            var v65 bool = false
                                            _ = v65
                                        }
                                        var v66 float32 = (1e10 - float32(((1e10 * (float64(-2.25) / (-2.25 + 1e10))) * 3.14159)))
                                        v58 = v58
                                        _ = v66
                                    }
                                    var v67 uint64 = 2
                                    _ = v67
                                }
                                v68 := GenS4{F0: true}
                                var v69 uint16 = 2
                                _ = v68
                                _ = v69
                            }
                            var v70 uint16 = 2
                            v70 += (v70 % ((65535) | 1))
                            _ = v58
                            _ = v70
                        }
                    case "hello":
                        var v71 uint32 = 2
                        _ = v71
                    default:
                        if (((v1 + "hello") > ((("\n\t" + v1) + "a") + "\n\t")) && false) {
                            fmt.Println(fi52)
                            v5 = (v5 - 0.1)
                            fmt.Println(fi52)
                        }
                    }
                    _ = fi57
                }
            } else {
                var v72 uint32 = 2147483647
                _ = v72
            }
        case "hello":
            var v73 int32 = 507
            _ = v73
        }
        v74 := GenS8{F0: 4294967295, F1: -77, F2: v2}
        _ = fi52
        _ = v53
        _ = v74
    }
    var v75 float32 = 1.5
    var v76 uint = 4294967295
    v75 = 0.1
    v1 = v1
    var v77 bool = true
    for fi78 := 0; fi78 < 2; fi78++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 = (v2 % ((v2) | 1))
        var v79 uint8 = 1
        _ = fi78
        _ = v79
    }
    if (0 <= 0) {
        fmt.Println(v5, v76)
        switch 32767 {
        case -1:
            _ = len(v4)
        default:
            if true {
                v5 -= (v5 / (v5 - 0.0))
            } else {
                fmt.Println(v75, v75)
            }
        }
    } else {
        if (v2 == 42) {
            if (!(0.1 < 0.1)) {
                if (v77 && (!false)) {
                    var v80 int16 = 42
                    switch ((((v1 + "a") + v1) + v1) + (v1 + "hello")) {
                    case "foo bar":
                        if true {
                            for fi81 := 0; fi81 < 15; fi81++ {
                                var v82 int = fi81
                                _ = fi81
                                _ = v82
                            }
                            if (1 == -9223372036854775808) {
                                for fi83 := 0; fi83 < 8; fi83++ {
                                    _ = len(v4)
                                    _ = fi83
                                }
                                switch ("\n\t" + "hello") {
                                case "a":
                                    _ = len(v4)
                                }
                                switch ((v1 + "a") + "hello") {
                                case "a":
                                    if false {
                                        var v84 float32 = (v75 / v75)
                                        _ = v84
                                    } else {
                                        v80 = (^(v80))
                                        for fi85 := 0; fi85 < 15; fi85++ {
                                            if true {
                                                v80 += (v80 ^ v80)
                                                fmt.Println(v75, fi85)
                                            } else {
                                                switch v1 {
                                                case "foo bar":
                                                    var v86 bool = (!v77)
                                                    _ = v86
                                                case "":
                                                    if len(v3) > 0 {
                                                        v3[0] = ("" + "a")
                                                    }
                                                }
                                                v77 = (2147483647 > 70)
                                                var v87 int64 = -9223372036854775808
                                                _ = v87
                                            }
                                            _ = fi85
                                        }
                                        var v88 string = v1
                                        _ = v88
                                    }
                                case "\n\t":
                                    switch -9223372036854775808 {
                                    case 781:
                                        v89 := GenS1{F0: v76, F1: (v2 / (((v2 * v2)) | 1)), F2: 9223372036854775807}
                                        _ = v89
                                    case 0:
                                        v90 := GenS3{F0: v80, F1: (v80 | (^(v80))), F2: v2}
                                        _ = v90
                                    default:
                                        var v91 uint64 = 12
                                        _ = v91
                                    }
                                case "foo bar":
                                    if (1 < 317) {
                                        switch ("" + "\n\t") {
                                        case "hello":
                                            v3 = append(v3, "hello")
                                        default:
                                            var v92 bool = v77
                                            _ = v92
                                        }
                                        var v93 int32 = (v2 ^ 523)
                                        _ = v93
                                    } else {
                                        if ((-978 <= -2147483648) || ("hello" > v1)) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        v94 := GenS9{F0: 439}
                                        v1 = "a"
                                        _ = v94
                                    }
                                }
                            }
                            v95 := v4[-9223372036854775808]
                            _ = v95
                        }
                    case "hello":
                        var v96 uint = v76
                        _ = v96
                    case "\n\t":
                        var v97 uint32 = 65535
                        _ = v97
                    }
                    var v98 uint = 784
                    _ = v80
                    _ = v98
                } else {
                    v99 := GenS10{F0: (v76 | (v76 << (((v76 << ((133) & 63))) & 63)))}
                    v77 = true
                    _ = v99
                }
            } else {
                v76 -= v76
                fmt.Println(v77, v75)
                var v100 int = 4294967295
                _ = v100
            }
            var v101 float32 = 0.1
            v102 := GenS11{F0: -100, F1: v1}
            _ = v101
            _ = v102
        } else {
            switch 878 {
            case 0:
                var v103 int16 = 1000
                _ = v103
            case -166:
                v5 = -2.25
            default:
                v1 = v1
            }
        }
    }
    if true {
        v5 -= v75
        v2 -= v2
        if true {
            var v104 int32 = v2
            v77 = ((32767 != -455) || v77)
            _ = v104
        } else {
            v2 = (v2 / ((-2147483648) | 1))
            if (32767 != -32768) {
                if ((v76 >> ((v76) & 63)) != (v76 / ((v76) | 1))) {
                    fmt.Println(v2, v2, v76)
                    var v105 float32 = (v75 * v5)
                    var v106 uint32 = 65535
                    _ = v105
                    _ = v106
                } else {
                    var v107 float64 = (0.1 * -2.25)
                    _ = v107
                }
                if v77 {
                    v76 = (^(v76))
                }
                var v108 bool = v77
                _ = v108
            }
        }
    }
    delete(v4, 1)
    for fi109 := 0; fi109 < 20; fi109++ {
        v76 += v76
        var v110 int32 = (v2 - v2)
        _ = fi109
        _ = v110
    }
    if ((9223372036854775807 == -9223372036854775808) && ((v2 % ((v2) | 1)) < (v2 % ((v2) | 1)))) {
        for fi111 := 0; fi111 < 18; fi111++ {
            var v112 int32 = v2
            if (true || v77) {
                switch (("foo bar" + ((v1 + v1) + v1)) + (v1 + v1)) {
                case "hello":
                    v76 += (v76 + v76)
                default:
                    v76 = v76
                }
                if v77 {
                    v113 := GenS12{F0: v112}
                    var v114 uint32 = 2147483648
                    v76 = v76
                    _ = v113
                    _ = v114
                } else {
                    v115 := v4[(^(fi111))]
                    v116 := GenS5{F0: -115, F1: 42, F2: 127}
                    if (1 <= 0) {
                        for fi117 := 0; fi117 < 1; fi117++ {
                            if v115 {
                                var v118 bool = (-128 < 127)
                                var v119 uint64 = 2
                                _ = v118
                                _ = v119
                            } else {
                                v5 = v75
                                fmt.Println(v5, fi111)
                                v115 = (0 < 0)
                            }
                            v112 += v112
                            _ = fi117
                        }
                    }
                    _ = v115
                    _ = v116
                }
            }
            _ = fi111
            _ = v112
        }
        var v120 uint16 = 179
        for fi121 := 0; fi121 < 11; fi121++ {
            var v122 uint = v76
            for fi123 := 0; fi123 < 15; fi123++ {
                v124 := GenS2{F0: -677}
                var v125 int = 1
                if (!v77) {
                    var v126 int64 = 42
                    if (!v77) {
                        var v127 uint64 = uint64(v122)
                        if (!v77) {
                            if false {
                                v76 -= (^(v122))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v128 float32 = v5
                                v129 := GenS3{F0: 119, F1: -1, F2: v2}
                                if (42 > -1) {
                                    var v130 int32 = v2
                                    v131 := v4[128]
                                    v131 = (42 <= -11)
                                    _ = v130
                                    _ = v131
                                }
                                _ = v128
                                _ = v129
                            } else {
                                v132 := GenS2{F0: -616}
                                var v133 int32 = v2
                                _ = v132
                                _ = v133
                            }
                        } else {
                            fmt.Println(v120, v125)
                            var v134 uint64 = uint64(v122)
                            _ = len(v4)
                            _ = v134
                        }
                        _ = v127
                    }
                    _ = v126
                } else {
                    var v135 int16 = 403
                    if true {
                        var v136 int16 = v135
                        var v137 float64 = 1.5
                        _ = v136
                        _ = v137
                    } else {
                        v138 := GenS5{F0: 0, F1: 25, F2: (v120 / ((21) | 1))}
                        _ = v138
                    }
                    _ = v135
                }
                v139 := GenS13{F0: v77, F1: 1}
                var v140 uint = v76
                _ = fi123
                _ = v124
                _ = v125
                _ = v139
                _ = v140
            }
            for fi141 := 0; fi141 < 7; fi141++ {
                if true {
                    v120 -= v120
                    var v142 uint = v76
                    _ = v142
                } else {
                    switch "foo bar" {
                    case "":
                        switch "" {
                        case "":
                            var v143 int64 = 212
                            _ = v143
                        case "hello":
                            fmt.Println(fi141)
                        default:
                            var v144 uint8 = 2
                            _ = v144
                        }
                    default:
                        var v145 string = v1
                        _ = v145
                    }
                }
                if false {
                    v146 := GenS11{F0: -1, F1: v1}
                    _ = v146
                }
                var v147 bool = v77
                _ = fi141
                _ = v147
            }
            _ = fi121
            _ = v122
        }
        _ = v120
    } else {
        if (!v77) {
            v76 += 4294967295
            fmt.Println(v75, v1)
            switch -2147483648 {
            case 127:
                v148 := GenS14{F0: 2, F1: v1, F2: v2}
                _ = v148
            case 607:
                v5 -= 3.14159
            }
        } else {
            for fi149 := 0; fi149 < 19; fi149++ {
                if (fi149 >= fi149) {
                    v1 = v1
                    if (4294967295 > 0) {
                        v77 = (32767 == 1000)
                        var v150 float32 = (1e10 * v5)
                        _ = v150
                    } else {
                        v151 := GenS5{F0: -1, F1: 866, F2: 2}
                        fmt.Println(v2)
                        _ = v151
                    }
                }
                _ = fi149
            }
            v1 = ("foo bar" + v1)
        }
        if v77 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    fmt.Println(v75, v5)
    fmt.Println(v77)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v76
    _ = v77
}
