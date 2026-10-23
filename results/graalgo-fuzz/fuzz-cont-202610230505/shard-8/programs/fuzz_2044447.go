package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint16; F1 int8; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 bool; F1 int32; F2 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS14 struct { F0 int64; F1 int16; F2 int8 }
type GenS15 struct { F0 string; F1 int16 }
type GenS16 struct { F0 uint8; F1 uint32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int32; F1 uint16; F2 int16 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 int8; F1 int32; F2 int8 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 int8 = -1
    var v5 float64 = 0.0
    var v6 int = 4294967295
    v7 := []int{}
    if false {
        var v8 uint8 = v1
        var v9 uint64 = 42
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint64 = 2
            _ = fi10
            _ = v11
        }
        v4 += (v4 - v4)
    }
    if true {
        v12 := GenS1{F0: 449, F1: 0, F2: -9223372036854775808}
        if ("a" <= "hello") {
            var v13 uint32 = 2147483647
            _ = v13
        }
        _ = v12
    } else {
        v14 := GenS2{F0: 42, F1: 984}
        for fi15 := 0; fi15 < 9; fi15++ {
            switch 32767 {
            case 42:
                v5 -= ((v5 - (((v5 + 0.0) - (float64(-2.25) + float64(3.14159))) + float64((-2.25 - 0.1)))) - v5)
            }
            if true {
                var v16 uint8 = v1
                _ = v16
            } else {
                v17 := v2[int(v4)]
                v5 -= float64(1.5)
                v18 := GenS3{F0: (^(v17))}
                _ = v17
                _ = v18
            }
            fmt.Println(v1, v5)
            _ = fi15
        }
        v4 = (v4 & (v4 * -1))
        _ = v14
    }
    v19 := GenS4{F0: v4, F1: "a"}
    switch 42 {
    case 32767:
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS5{F0: 32767, F1: v4, F2: 925}
            if (!false) {
                var v22 int32 = 1
                if false {
                    v23 := GenS6{F0: 653}
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if (false || (v1 == (v1 >> ((42) & 63)))) {
                            v25 := GenS7{F0: (v1 & (v1 << ((838) & 63))), F1: v1}
                            fmt.Println(v22, fi20)
                            _ = v25
                        } else {
                            _ = len(v7)
                            v26 := GenS8{F0: true, F1: v22, F2: 1}
                            v5 += v5
                            _ = v26
                        }
                        v27 := GenS1{F0: 478, F1: uint32(v1), F2: 9223372036854775807}
                        _ = fi24
                        _ = v27
                    }
                    v4 = v4
                    _ = v23
                } else {
                    var v28 uint16 = 65535
                    switch (fi20 - (fi20 * 128)) {
                    case 9223372036854775807:
                        if false {
                            switch "hello" {
                            case "foo bar":
                                if (((-2.25 / 1.5) * 1.5) < 3.14159) {
                                    switch (fi20 << ((65535) & 63)) {
                                    case -2147483648:
                                        var v29 int64 = 897
                                        _ = v29
                                    case -930:
                                        for fi30 := 0; fi30 < 9; fi30++ {
                                            fmt.Println(v22)
                                            if true {
                                                var v31 bool = ("a" >= ("foo bar" + ""))
                                                var v32 string = "hello"
                                                _ = v31
                                                _ = v32
                                            } else {
                                                fmt.Println(v1)
                                                if len(v7) > 0 {
                                                    v7[0] = (fi30 / (((v6 ^ v6)) | 1))
                                                }
                                                if len(v7) > 0 {
                                                    v7[0] = (v6 | fi30)
                                                }
                                            }
                                            _ = fi30
                                        }
                                    }
                                    fmt.Println(v1)
                                }
                            case "a":
                                if (("hello" + "foo bar") > "a") {
                                    var v33 bool = (598 != 0)
                                    var v34 int32 = v22
                                    _ = v33
                                    _ = v34
                                } else {
                                    if true {
                                        var v35 uint64 = 0
                                        _ = v35
                                    } else {
                                        var v36 uint = 175
                                        switch "hello" {
                                        case "a":
                                            var v37 uint8 = 128
                                            _ = v37
                                        case "":
                                            v3[((("\n\t" + "\n\t") + (("" + "") + ((("a" + "hello") + ("hello" + "foo bar")) + ""))) + ("foo bar" + "\n\t"))] = (-(v6))
                                        }
                                        _ = v36
                                    }
                                    for fi38 := 0; fi38 < 13; fi38++ {
                                        v39 := GenS9{F0: 1}
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                            }
                            _ = len(v7)
                            switch "" {
                            case "foo bar":
                                v28 = v28
                            case "\n\t":
                                switch (v6 % ((-1) | 1)) {
                                case 4294967295:
                                    fmt.Println(v6)
                                case 1:
                                    if (!true) {
                                        if true {
                                            var v40 float32 = 0.1
                                            _ = v40
                                        }
                                    } else {
                                        v22 = -1
                                        var v41 uint64 = 1
                                        _ = v41
                                    }
                                case 9223372036854775807:
                                    if ((v6 << ((762) & 63)) > (fi20 << ((255) & 63))) {
                                        v7 = append(v7, v6)
                                        v22 += 0
                                    } else {
                                        var v42 float32 = -2.25
                                        var v43 uint = 1
                                        v6 += (-(v6))
                                        _ = v42
                                        _ = v43
                                    }
                                default:
                                    switch 574 {
                                    case 4294967295:
                                        if (!(!(3.14159 <= (1e10 + -2.25)))) {
                                            var v44 uint64 = 1
                                            v6 = 4294967295
                                            _ = v44
                                        } else {
                                            if false {
                                                var v45 string = "a"
                                                var v46 int64 = 408
                                                _ = v45
                                                _ = v46
                                            } else {
                                                var v47 uint64 = 18446744073709551615
                                                v48 := GenS10{F0: 255}
                                                var v49 uint32 = 800
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                    case 32767:
                                        fmt.Println(v28, v22)
                                    default:
                                        var v50 uint32 = 29
                                        _ = v50
                                    }
                                }
                            case "a":
                                var v51 int16 = 0
                                _ = v51
                            }
                        } else {
                            switch "foo bar" {
                            case "hello":
                                var v52 int64 = -1
                                _ = v52
                            case "":
                                var v53 uint = 18446744073709551615
                                _ = v53
                            case "\n\t":
                                delete(v2, v6)
                            }
                            if (415 < 537) {
                                switch -138 {
                                case 32768:
                                    var v54 int32 = (-(v22))
                                    _ = v54
                                case -160:
                                    v55 := GenS11{F0: true}
                                    _ = v55
                                case 2147483647:
                                    v1 -= (v1 + 0)
                                default:
                                    v4 -= v4
                                }
                            } else {
                                var v56 float64 = ((v5 - (0.1 - float64(float32(float64((((1.5 * -2.25) / 1.5) - 3.14159)))))) + v5)
                                v28 = 32767
                                _ = v56
                            }
                            v57 := GenS12{F0: -32768}
                            _ = v57
                        }
                    case 42:
                        var v58 int16 = 0
                        _ = v58
                    default:
                        v5 = v5
                    }
                    _ = v28
                }
                if (int16(fi20) <= 1000) {
                    if (((0.0 * (float32(v5) - 0.1)) - (float32(v5) + (1.5 - 1.5))) == ((3.14159 - 0.0) - float32((v5 * (v5 + v5))))) {
                        var v59 uint8 = v1
                        _ = v59
                    } else {
                        v60 := GenS13{F0: 0, F1: (!true), F2: 704}
                        _ = v60
                    }
                    var v61 uint = 1
                    _ = v61
                }
                _ = v22
            } else {
                var v62 string = (((("hello" + "\n\t") + "foo bar") + ("a" + "foo bar")) + "")
                if false {
                    _ = len(v3)
                    for fi63 := 0; fi63 < 1; fi63++ {
                        var v64 uint16 = 255
                        _ = fi63
                        _ = v64
                    }
                } else {
                    var v65 uint8 = (v1 << ((255) & 63))
                    var v66 float64 = v5
                    var v67 uint8 = 2
                    _ = v65
                    _ = v66
                    _ = v67
                }
                for fi68 := 0; fi68 < 1; fi68++ {
                    v7 = append(v7, fi20)
                    v62 = v62
                    _ = fi68
                }
                _ = v62
            }
            _ = fi20
            _ = v21
        }
    case 0:
        var v69 uint64 = 1
        _ = v69
    case 32768:
        if false {
            var v70 float64 = float64(-2.25)
            delete(v2, v6)
            v71 := GenS11{F0: false}
            _ = v70
            _ = v71
        }
    }
    v72 := GenS13{F0: 65535, F1: true, F2: 1}
    if ((!(true || true)) || (42 > -795)) {
        if ((^(v1)) < v1) {
            for fi73 := 0; fi73 < 5; fi73++ {
                for fi74 := 0; fi74 < 18; fi74++ {
                    var v75 int = v6
                    _ = fi74
                    _ = v75
                }
                if ((fi73 ^ (fi73 % ((v6) | 1))) != int(v1)) {
                    v76 := GenS8{F0: false, F1: -994, F2: 189}
                    v77 := GenS14{F0: -846, F1: -466, F2: (v4 & v4)}
                    _ = v76
                    _ = v77
                } else {
                    switch (-(v6)) {
                    case 9223372036854775807:
                        if (65535 <= 255) {
                            for fi78 := 0; fi78 < 2; fi78++ {
                                var v79 uint = 456
                                if false {
                                    v5 -= (v5 - (v5 / v5))
                                    if ((^(v6)) == -1) {
                                        if true {
                                            var v80 int16 = -253
                                            _ = v80
                                        } else {
                                            if (!(false && true)) {
                                                var v81 uint = v79
                                                fmt.Println(v79, v6, v5)
                                                _ = v81
                                            }
                                            if false {
                                                fmt.Println(v5, v6)
                                                var v82 string = ("a" + "\n\t")
                                                if true {
                                                    if true {
                                                        v83 := GenS15{F0: ("" + v82), F1: -1}
                                                        if false {
                                                            switch v6 {
                                                            case -2147483648:
                                                                var v84 uint = (^(v79))
                                                                _ = v84
                                                            case -28:
                                                                v85 := GenS11{F0: (!(v5 <= v5))}
                                                                _ = v85
                                                            case 127:
                                                                var v86 uint8 = (v1 & (v1 >> (((v79 << ((81) & 63))) & 63)))
                                                                _ = v86
                                                            default:
                                                                if (((float32(0.1) + (0.0 * -2.25)) / (-2.25 - 0.1)) < 0.0) {
                                                                    var v87 int = v6
                                                                    var v88 uint8 = (v1 * v1)
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            }
                                                            v2[(^(fi73))] = 648
                                                        }
                                                        _ = v83
                                                    }
                                                    fmt.Println(v5, fi78)
                                                    v89 := GenS7{F0: v1, F1: 89}
                                                    _ = v89
                                                } else {
                                                    v90 := GenS5{F0: 42, F1: 1, F2: (v79 - v79)}
                                                    v82 = "\n\t"
                                                    fmt.Println(fi73, v82)
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    fmt.Println(fi78)
                                                    _ = v90
                                                }
                                                _ = v82
                                            }
                                            switch (fi78 - (fi73 & (fi78 / (((fi73 << ((v79) & 63))) | 1)))) {
                                            case 4294967295:
                                                var v91 string = (("hello" + "\n\t") + "foo bar")
                                                _ = v91
                                            case 127:
                                                if ((0.0 < (3.14159 - 0.1)) || true) {
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    v92 := v3["hello"]
                                                    v93 := GenS16{F0: v1, F1: 2147483648}
                                                    switch (fi78 >> ((v79) & 63)) {
                                                    case 32767:
                                                        v5 = (float64(float32(1e10)) * float64(((((float32(v5) + float32(float64(0.0))) * 0.0) * 3.14159) - 0.0)))
                                                    case -799:
                                                        v7 = append(v7, (-(fi78)))
                                                    default:
                                                        var v94 uint16 = 32767
                                                        _ = v94
                                                    }
                                                    _ = v92
                                                    _ = v93
                                                }
                                            }
                                            if (true || ("hello" >= "foo bar")) {
                                                var v95 bool = true
                                                var v96 int16 = 249
                                                _ = v95
                                                _ = v96
                                            }
                                            fmt.Println(v4, v79, v4)
                                        }
                                        if (-9223372036854775808 != (fi73 - fi78)) {
                                            if (608 == 0) {
                                                var v97 uint16 = 42
                                                v97 = 65535
                                                if true {
                                                    if (1 <= (v97 % ((v97) | 1))) {
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        v6 += (-(fi73))
                                                        v98 := GenS1{F0: v79, F1: 2147483647, F2: 1}
                                                        _ = v98
                                                    } else {
                                                        v3["foo bar"] = (-(fi78))
                                                        var v99 bool = (false && (((255 < v97) || (!false)) || (true || (!(2147483647 >= 1)))))
                                                        v100 := GenS17{F0: v79}
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                    if (1e10 > v5) {
                                                        var v101 string = "foo bar"
                                                        switch (fi73 - fi73) {
                                                        case 1:
                                                            v3[v101] = (fi73 & int(v4))
                                                        case 32767:
                                                            var v102 int64 = 521
                                                            _ = v102
                                                        default:
                                                            v6 -= fi78
                                                        }
                                                        v103 := GenS4{F0: (v4 ^ (v4 >> ((v79) & 63))), F1: "a"}
                                                        v104 := v2[-9223372036854775808]
                                                        switch v101 {
                                                        case "":
                                                            v4 += 0
                                                        }
                                                        _ = v101
                                                        _ = v103
                                                        _ = v104
                                                    } else {
                                                        v3[""] = -1
                                                        if true {
                                                            v97 -= 42
                                                        } else {
                                                            v97 = 778
                                                            v105 := GenS15{F0: ("hello" + "foo bar"), F1: 0}
                                                            _ = v105
                                                        }
                                                    }
                                                    v106 := GenS18{F0: (false && ((602 != 42) && (false || true)))}
                                                    _ = v106
                                                } else {
                                                    v7 = append(v7, fi78)
                                                }
                                                _ = v97
                                            } else {
                                                v107 := GenS18{F0: (v4 <= (v4 ^ (^(v4))))}
                                                v2[(fi73 | (v6 % ((32768) | 1)))] = (fi78 * 11)
                                                _ = v107
                                            }
                                            switch (fi78 - fi78) {
                                            case 9223372036854775807:
                                                v6 = fi78
                                                var v108 int16 = 581
                                                fmt.Println(v6, v1, v1)
                                                var v109 float64 = 3.14159
                                                _ = v108
                                                _ = v109
                                            case -61:
                                                v110 := GenS10{F0: (v1 * (v1 * (v1 << (((v79 ^ 65535)) & 63))))}
                                                _ = v110
                                            }
                                            var v111 uint32 = 2147483648
                                            if false {
                                                fmt.Println(v4, v1)
                                            } else {
                                                v112 := GenS8{F0: false, F1: 129, F2: -963}
                                                var v113 uint = (v79 - v79)
                                                _ = v112
                                                _ = v113
                                            }
                                            var v114 uint32 = (v111 ^ (^(v111)))
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                if (false && (!true)) {
                                                    fmt.Println(fi78, v5, v4)
                                                } else {
                                                    var v115 float64 = v5
                                                    v5 = (v5 * float64(3.14159))
                                                    v116 := GenS19{F0: 42, F1: 961, F2: 985}
                                                    _ = v115
                                                    _ = v116
                                                }
                                            case "\n\t":
                                                var v117 uint16 = 0
                                                _ = v117
                                            case "a":
                                                var v118 int32 = 0
                                                _ = v118
                                            }
                                            _ = v111
                                            _ = v114
                                        }
                                    } else {
                                        var v119 uint64 = 42
                                        var v120 int16 = 257
                                        _ = v119
                                        _ = v120
                                    }
                                } else {
                                    v121 := GenS20{F0: uint16(v1)}
                                    fmt.Println(fi73, fi73)
                                    var v122 int32 = 971
                                    if (true && (v79 != (^(v79)))) {
                                        fmt.Println(v4)
                                    }
                                    var v123 int16 = 1000
                                    _ = v121
                                    _ = v122
                                    _ = v123
                                }
                                v1 = 255
                                _ = fi78
                                _ = v79
                            }
                            switch -922 {
                            case -989:
                                v124 := GenS8{F0: false, F1: 2147483647, F2: 2147483647}
                                _ = v124
                            case 1:
                                for fi125 := 0; fi125 < 7; fi125++ {
                                    switch 587 {
                                    case 1:
                                        var v126 float64 = v5
                                        _ = v126
                                    case -1:
                                        v4 += v4
                                    default:
                                        var v127 uint64 = 912
                                        _ = v127
                                    }
                                    v128 := GenS16{F0: 0, F1: 4294967295}
                                    switch "a" {
                                    case "a":
                                        v129 := GenS21{F0: (v4 / ((v4) | 1)), F1: -2147483648, F2: int8(fi125)}
                                        _ = v129
                                    case "\n\t":
                                        var v130 uint = 0
                                        _ = v130
                                    }
                                    _ = fi125
                                    _ = v128
                                }
                            }
                        } else {
                            fmt.Println(fi73, v6)
                            v5 += ((v5 * v5) * ((1e10 / ((v5 + float64(float32((float64(((3.14159 / (1.5 * -2.25)) - 3.14159)) / 0.1)))) + v5)) + (v5 + v5)))
                        }
                    case -676:
                        fmt.Println(v6, fi73, v6)
                    case -2147483648:
                        var v131 string = "a"
                        _ = v131
                    default:
                        v132 := GenS20{F0: 32767}
                        _ = v132
                    }
                }
                _ = fi73
            }
            switch "a" {
            case "\n\t":
                var v133 uint64 = 523
                _ = v133
            case "hello":
                v6 += -9223372036854775808
            }
            v6 -= v6
        }
        v1 += v1
        if (!(375 != v6)) {
            var v134 int8 = (v4 | (v4 * (v4 + 75)))
            _ = v134
        }
    }
    if true {
        v5 = v5
        if true {
            v7 = append(v7, (v6 + (v6 % (((v6 + 600)) | 1))))
        }
    } else {
        var v135 float64 = float64(1e10)
        fmt.Println(v6, v1)
        _ = v135
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v72
}
