package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
type GenS4 struct { F0 int; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint32; F1 string; F2 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 int64; F2 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64; F1 uint16; F2 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool; F1 int64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint32; F1 int16 }
type GenS15 struct { F0 int; F1 int8 }
type GenS16 struct { F0 int; F1 int8; F2 int32 }
type GenS17 struct { F0 int32; F1 int64 }
type GenS18 struct { F0 string; F1 string; F2 int64 }
type GenS19 struct { F0 uint8; F1 int }

func main() {
    var v1 uint64 = 2
    var v2 bool = false
    var v3 int = -737
    if (false && v2) {
        for fi4 := 0; fi4 < 4; fi4++ {
            fmt.Println(v3, v1, v2)
            var v5 float32 = 0.1
            _ = fi4
            _ = v5
        }
        fmt.Println(v3, v2, v2)
        fmt.Println(v3, v2)
    } else {
        if (21 == -32768) {
            var v6 float64 = 1e10
            var v7 float32 = (1.5 + 0.1)
            _ = v6
            _ = v7
        } else {
            switch v3 {
            case 415:
                switch (v3 >> ((1) & 63)) {
                case 668:
                    var v8 int8 = -128
                    _ = v8
                case 0:
                    if v2 {
                        v3 -= v3
                        var v9 int = v3
                        v2 = true
                        _ = v9
                    }
                case -2147483648:
                    var v10 bool = true
                    _ = v10
                default:
                    if false {
                        v1 = v1
                    }
                }
            case 32767:
                var v11 float64 = 3.14159
                _ = v11
            case -845:
                for fi12 := 0; fi12 < 1; fi12++ {
                    if ((-668 > fi12) || (v2 || v2)) {
                        var v13 int64 = -18
                        _ = v13
                    } else {
                        v2 = (-2147483648 != 1)
                    }
                    var v14 bool = true
                    _ = fi12
                    _ = v14
                }
            }
            v1 = 989
            switch v3 {
            case -2147483648:
                if (1 == -1) {
                    v1 += 0
                    for fi15 := 0; fi15 < 15; fi15++ {
                        if v2 {
                            v3 = (fi15 ^ fi15)
                        }
                        if v2 {
                            v3 = -9223372036854775808
                        }
                        v16 := GenS1{F0: 553, F1: 127}
                        _ = fi15
                        _ = v16
                    }
                    fmt.Println(v2, v3, v3)
                } else {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        if (v3 == (fi17 + (^(v3)))) {
                            for fi18 := 0; fi18 < 4; fi18++ {
                                v19 := GenS2{F0: 194, F1: uint(v1), F2: 128}
                                var v20 float64 = ((1e10 + 3.14159) * 0.0)
                                if v2 {
                                    v21 := GenS3{F0: "foo bar", F1: 32767, F2: -102}
                                    v1 -= (v1 | v1)
                                    _ = v21
                                }
                                v22 := []int{v3, fi18}
                                v1 += (v1 / ((v1) | 1))
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v22
                            }
                            if false {
                                v3 = (v3 / ((2147483647) | 1))
                            } else {
                                v23 := GenS2{F0: 42, F1: 363, F2: 127}
                                v24 := GenS4{F0: fi17, F1: 128, F2: 127}
                                var v25 int16 = 42
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            fmt.Println(v3, v1, v2)
                        }
                        v26 := []int{(fi17 >> ((42) & 63)), -1000, v3, (fi17 * -822)}
                        _ = fi17
                        _ = v26
                    }
                    v1 += (v1 >> ((0) & 63))
                    if (706 == 0) {
                        v3 -= 42
                        var v27 string = ("a" + "a")
                        _ = v27
                    } else {
                        if true {
                            v28 := GenS5{F0: -154, F1: 0}
                            v29 := []int{(v3 & v3), v3, 1, (v3 / ((-770) | 1))}
                            for fi30 := 0; fi30 < 17; fi30++ {
                                if (v2 || true) {
                                    var v31 int32 = 588
                                    var v32 float32 = (0.1 / 1.0)
                                    v1 = v1
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := GenS3{F0: ("hello" + "a"), F1: -32768, F2: 100}
                                    var v34 uint = 1
                                    _ = v33
                                    _ = v34
                                }
                                _ = fi30
                            }
                            _ = v28
                            _ = v29
                        }
                        v35 := GenS3{F0: "foo bar", F1: 1000, F2: -86}
                        _ = v35
                    }
                }
            case 1:
                if (v3 > v3) {
                    v1 = v1
                }
            }
        }
        for fi36 := 0; fi36 < 19; fi36++ {
            if ((v2 || v2) || false) {
                v37 := []int{v3, (v3 >> ((384) & 63)), fi36, v3}
                _ = v37
            } else {
                v2 = v2
            }
            v2 = (18446744073709551615 != 733)
            v38 := GenS6{F0: 0, F1: "foo bar", F2: (1 > v1)}
            _ = fi36
            _ = v38
        }
    }
    if v2 {
        v39 := GenS7{F0: -1}
        _ = v39
    } else {
        if ((v3 << ((1) & 63)) > 128) {
            var v40 uint16 = 596
            var v41 uint64 = v1
            _ = v40
            _ = v41
        }
        switch v3 {
        case -2147483648:
            v42 := []int{v3, (v3 / ((v3) | 1))}
            _ = v42
        }
    }
    v1 += (^(v1))
    if v2 {
        var v43 int16 = 929
        _ = v43
    } else {
        v44 := GenS7{F0: 147}
        _ = v44
    }
    v1 = 18446744073709551615
    for fi45 := 0; fi45 < 8; fi45++ {
        if v2 {
            var v46 float64 = 1e10
            var v47 string = "hello"
            if (false || (65535 < uint32(v1))) {
                for fi48 := 0; fi48 < 11; fi48++ {
                    if (true && true) {
                        v47 = (v47 + "foo bar")
                        var v49 uint32 = 42
                        var v50 int8 = -40
                        _ = v49
                        _ = v50
                    } else {
                        fmt.Println(fi45, v3)
                    }
                    _ = fi48
                }
                var v51 uint16 = 2
                if (-61 >= 100) {
                    v46 = float64(0.1)
                    for fi52 := 0; fi52 < 7; fi52++ {
                        fmt.Println(fi52, fi45, v51)
                        _ = fi52
                    }
                    v53 := []int{}
                    _ = v53
                }
                _ = v51
            } else {
                switch v47 {
                case "foo bar":
                    var v54 uint8 = 2
                    _ = v54
                case "a":
                    if v2 {
                        if v2 {
                            for fi55 := 0; fi55 < 20; fi55++ {
                                var v56 float32 = 3.14159
                                _ = fi55
                                _ = v56
                            }
                            v57 := []int{(fi45 % ((fi45) | 1)), (fi45 & (fi45 * 32768)), v3, fi45}
                            _ = v57
                        } else {
                            v58 := GenS4{F0: (fi45 << ((18446744073709551615) & 63)), F1: 255, F2: 128}
                            if (!(v47 < ((v47 + (v47 + (v47 + (v47 + ((("\n\t" + v47) + v47) + ((v47 + v47) + v47)))))) + (v47 + v47)))) {
                                v59 := []int{}
                                var v60 int = (fi45 | (v3 / (((fi45 * (fi45 & fi45))) | 1)))
                                _ = v59
                                _ = v60
                            } else {
                                v46 = float64(-2.25)
                                var v61 uint32 = 130
                                var v62 uint = 94
                                _ = v61
                                _ = v62
                            }
                            v3 = v3
                            _ = v58
                        }
                    } else {
                        if ((v2 && (255 >= 129)) && ((false || v2) || (42 <= 255))) {
                            for fi63 := 0; fi63 < 20; fi63++ {
                                switch "a" {
                                case "":
                                    if v2 {
                                        switch v47 {
                                        case "hello":
                                            var v64 uint16 = 1
                                            _ = v64
                                        case "a":
                                            v65 := GenS5{F0: 32767, F1: 65535}
                                            _ = v65
                                        case "":
                                            switch ("foo bar" + (v47 + "a")) {
                                            case "":
                                                v66 := GenS8{F0: (true || (0 != 2147483647)), F1: 42, F2: 1000}
                                                v1 += 4294967295
                                                switch (^(v3)) {
                                                case 4294967295:
                                                    var v67 int8 = 1
                                                    _ = v67
                                                }
                                                if (false || (v2 && (-312 <= 1000))) {
                                                    if true {
                                                        v3 = -2147483648
                                                        var v68 int16 = 1
                                                        _ = v68
                                                    }
                                                    v1 += v1
                                                }
                                                v69 := GenS4{F0: (v3 ^ fi45), F1: 90, F2: 42}
                                                if ((v3 | 42) <= 32768) {
                                                    v70 := GenS7{F0: -779}
                                                    switch (fi45 | 0) {
                                                    case 42:
                                                        v1 = 4294967295
                                                    case -957:
                                                        fmt.Println(fi63, v47)
                                                    case 32768:
                                                        fmt.Println(fi45, v2)
                                                    default:
                                                        if (fi45 > (fi63 - fi45)) {
                                                            v2 = true
                                                        } else {
                                                            fmt.Println(v3, v47)
                                                            if (1 >= 769) {
                                                                v71 := []int{(fi45 + v3), v3}
                                                                if (true || v2) {
                                                                    v72 := GenS9{F0: 34}
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        switch (v47 + v47) {
                                                                        case "\n\t":
                                                                            var v73 float32 = 1.5
                                                                            _ = v73
                                                                        case "hello":
                                                                            if len(v71) > 0 {
                                                                                v71[0] = (v3 << ((1) & 63))
                                                                            }
                                                                        case "":
                                                                            if len(v71) > 0 {
                                                                                v71[0] = (fi63 + v3)
                                                                            }
                                                                        default:
                                                                            v47 = (v47 + v47)
                                                                        }
                                                                    default:
                                                                        switch (v3 >> ((716) & 63)) {
                                                                        case 128:
                                                                            v3 += (v3 << ((979) & 63))
                                                                        default:
                                                                            var v74 float32 = ((((3.14159 * -2.25) + (float32(v46) * 3.14159)) * (3.14159 / 1.0)) / 1e10)
                                                                            _ = v74
                                                                        }
                                                                    }
                                                                    if (((("hello" + v47) + (("\n\t" + ("" + v47)) + v47)) + v47) <= v47) {
                                                                        var v75 int32 = -2147483648
                                                                        v76 := GenS1{F0: 18446744073709551615, F1: 42}
                                                                        v71 = append(v71, 2147483647)
                                                                        fmt.Println(v47, v2)
                                                                        v75 = v75
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                    var v77 int = -543
                                                                    switch v47 {
                                                                    case "foo bar":
                                                                        v78 := GenS10{F0: (v1 % ((2) | 1)), F1: 32767, F2: 0}
                                                                        _ = v78
                                                                    case "":
                                                                        var v79 uint64 = v1
                                                                        _ = v79
                                                                    default:
                                                                        _ = len(v71)
                                                                    }
                                                                    _ = v72
                                                                    _ = v77
                                                                } else {
                                                                    if v2 {
                                                                        var v80 float32 = 0.0
                                                                        var v81 uint8 = 255
                                                                        _ = v80
                                                                        _ = v81
                                                                    } else {
                                                                        if ((v47 + v47) == v47) {
                                                                            var v82 int8 = -41
                                                                            if len(v71) > 0 {
                                                                                v71[0] = fi63
                                                                            }
                                                                            var v83 int32 = 1
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                        v47 = (v47 + "")
                                                                        switch (-(fi63)) {
                                                                        case 767:
                                                                            switch ("a" + v47) {
                                                                            case "a":
                                                                                if v2 {
                                                                                    if len(v71) >= 2 {
                                                                                        v71 = v71[1:len(v71)]
                                                                                    }
                                                                                    v46 = 3.14159
                                                                                } else {
                                                                                    if (42 < 42) {
                                                                                        v84 := GenS2{F0: 2, F1: 255, F2: 214}
                                                                                        if (!v2) {
                                                                                            var v85 float64 = v46
                                                                                            _ = v85
                                                                                        }
                                                                                        v86 := GenS2{F0: 194, F1: 42, F2: 0}
                                                                                        if (-1 != -1) {
                                                                                            v47 = (v47 + v47)
                                                                                            v87 := GenS9{F0: 42}
                                                                                            _ = v87
                                                                                        }
                                                                                        _ = len(v71)
                                                                                        var v88 int16 = 917
                                                                                        var v89 int64 = 9223372036854775807
                                                                                        _ = v84
                                                                                        _ = v86
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                    } else {
                                                                                        if ((fi45 / ((fi45) | 1)) <= (fi63 - fi45)) {
                                                                                            var v90 uint8 = 174
                                                                                            var v91 uint32 = 2147483648
                                                                                            _ = v90
                                                                                            _ = v91
                                                                                        } else {
                                                                                            switch (fi45 / (((-(v3))) | 1)) {
                                                                                            case 128:
                                                                                                var v92 uint = 29
                                                                                                _ = v92
                                                                                            default:
                                                                                                var v93 int8 = -99
                                                                                                _ = v93
                                                                                            }
                                                                                            switch v47 {
                                                                                            case "hello":
                                                                                                fmt.Println(v2, v1)
                                                                                            }
                                                                                            var v94 uint32 = 891
                                                                                            _ = v94
                                                                                        }
                                                                                        v2 = (!(((((-2.25 / 1.5) * -2.25) / 3.14159) * 0.0) <= (float32(float64(0.0)) + (0.1 - 1e10))))
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                var v95 uint64 = v1
                                                                                _ = v95
                                                                            case "foo bar":
                                                                                var v96 int64 = 245
                                                                                v2 = true
                                                                                var v97 bool = true
                                                                                _ = v96
                                                                                _ = v97
                                                                            default:
                                                                                var v98 int32 = 161
                                                                                _ = v98
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi45, v2, fi63)
                                                                        }
                                                                        var v99 int64 = -47
                                                                        _ = v99
                                                                    }
                                                                    v47 = "hello"
                                                                }
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v2 = (432 >= 0)
                                                                case "":
                                                                    var v100 int8 = -1
                                                                    fmt.Println(v3)
                                                                    var v101 uint8 = 240
                                                                    _ = v100
                                                                    _ = v101
                                                                case "\n\t":
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        var v102 int16 = -32768
                                                                        _ = v102
                                                                    default:
                                                                        var v103 float64 = v46
                                                                        _ = v103
                                                                    }
                                                                }
                                                                _ = v71
                                                            }
                                                            v104 := GenS7{F0: -1}
                                                            _ = v104
                                                        }
                                                    }
                                                    v46 = v46
                                                    _ = v70
                                                }
                                                _ = v66
                                                _ = v69
                                            }
                                        default:
                                            var v105 float64 = -2.25
                                            _ = v105
                                        }
                                        if v2 {
                                            v106 := GenS7{F0: -681}
                                            v47 = ""
                                            if (!v2) {
                                                v107 := []int{fi63, 406}
                                                var v108 int16 = -763
                                                if (!v2) {
                                                    v46 += 3.14159
                                                    _ = len(v107)
                                                    var v109 uint8 = 127
                                                    var v110 int = fi63
                                                    v46 = v46
                                                    var v111 int64 = -1
                                                    if len(v107) >= 2 {
                                                        v107 = v107[1:len(v107)]
                                                    }
                                                    _ = v109
                                                    _ = v110
                                                    _ = v111
                                                }
                                                _ = v107
                                                _ = v108
                                            } else {
                                                if (!(!((2147483647 == v3) && ((!(!(1 != 759))) || (v46 <= v46))))) {
                                                    var v112 uint32 = 65535
                                                    v113 := []int{(fi45 % ((9223372036854775807) | 1)), 32768, int(v1)}
                                                    var v114 int = (fi63 | (fi45 << ((255) & 63)))
                                                    _ = v112
                                                    _ = v113
                                                    _ = v114
                                                } else {
                                                    if v2 {
                                                        v115 := []int{(v3 << ((uint(v1)) & 63)), (fi45 >> ((817) & 63)), int(v1)}
                                                        var v116 int32 = 2147483647
                                                        _ = v115
                                                        _ = v116
                                                    } else {
                                                        fmt.Println(v46)
                                                        fmt.Println(v1, v46)
                                                        var v117 uint32 = 4294967295
                                                        var v118 int32 = -1
                                                        var v119 int16 = -32768
                                                        if v2 {
                                                            var v120 string = (v47 + (v47 + (v47 + ("" + ("hello" + ("" + "\n\t"))))))
                                                            fmt.Println(v2)
                                                            var v121 int32 = (-(v118))
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        v122 := []int{fi63, fi63}
                                                        _ = v117
                                                        _ = v118
                                                        _ = v119
                                                        _ = v122
                                                    }
                                                    var v123 uint = 18446744073709551615
                                                    fmt.Println(v46, v2, v3)
                                                    _ = v123
                                                }
                                                switch int(v1) {
                                                case 127:
                                                    v124 := GenS8{F0: (-2147483648 <= -1), F1: -711, F2: -997}
                                                    _ = v124
                                                case 42:
                                                    var v125 int8 = 1
                                                    _ = v125
                                                }
                                                switch (v3 * (v3 << ((357) & 63))) {
                                                case -1:
                                                    if v2 {
                                                        v126 := GenS11{F0: -422}
                                                        if v2 {
                                                            var v127 float64 = v46
                                                            v128 := GenS12{F0: ((fi63 / (((fi63 << ((570) & 63))) | 1)) != (v3 & (v3 << ((2) & 63)))), F1: -1}
                                                            v3 -= (fi63 / (((v3 << ((4294967295) & 63))) | 1))
                                                            _ = v127
                                                            _ = v128
                                                        }
                                                        if (0 >= 127) {
                                                            var v129 int8 = 127
                                                            var v130 uint8 = 88
                                                            _ = v129
                                                            _ = v130
                                                        }
                                                        fmt.Println(fi63, v47, fi45)
                                                        _ = v126
                                                    } else {
                                                        v131 := []int{(fi63 << ((4294967295) & 63)), (fi45 - fi63), fi63}
                                                        if (32767 == 240) {
                                                            var v132 int16 = -32768
                                                            v133 := GenS4{F0: (v3 % ((fi63) | 1)), F1: 42, F2: 42}
                                                            fmt.Println(fi63, fi45, v47)
                                                            v132 = 214
                                                            v134 := GenS4{F0: fi45, F1: uint8(v1), F2: 2}
                                                            _ = v132
                                                            _ = v133
                                                            _ = v134
                                                        } else {
                                                            v135 := GenS13{F0: (-2.25 < float32(1e10))}
                                                            if len(v131) > 0 {
                                                                v131[0] = fi45
                                                            }
                                                            _ = v135
                                                        }
                                                        _ = v131
                                                    }
                                                case -720:
                                                    var v136 int = (fi63 ^ (v3 / ((-2147483648) | 1)))
                                                    _ = v136
                                                case 1:
                                                    if (!v2) {
                                                        var v137 string = v47
                                                        var v138 int32 = 42
                                                        _ = v137
                                                        _ = v138
                                                    } else {
                                                        var v139 int8 = -128
                                                        if v2 {
                                                            v140 := GenS14{F0: 718, F1: -838}
                                                            _ = v140
                                                        }
                                                        var v141 uint = 2
                                                        _ = v139
                                                        _ = v141
                                                    }
                                                default:
                                                    v142 := GenS15{F0: 32767, F1: -74}
                                                    _ = v142
                                                }
                                                v143 := GenS16{F0: v3, F1: 127, F2: -190}
                                                switch ((v47 + v47) + v47) {
                                                case "hello":
                                                    v46 += float64((0.0 / 1.5))
                                                }
                                                _ = v143
                                            }
                                            _ = v106
                                        }
                                        var v144 uint8 = 234
                                        _ = v144
                                    }
                                    var v145 int16 = 32767
                                    _ = v145
                                default:
                                    var v146 bool = (42 < 0)
                                    _ = v146
                                }
                                v147 := GenS6{F0: 4294967295, F1: v47, F2: (v2 || v2)}
                                _ = fi63
                                _ = v147
                            }
                            if (-218 > -1) {
                                v46 -= (3.14159 - -2.25)
                            } else {
                                v148 := GenS17{F0: 941, F1: 42}
                                for fi149 := 0; fi149 < 3; fi149++ {
                                    switch (v3 & (v3 + -831)) {
                                    case 127:
                                        var v150 uint32 = 244
                                        _ = v150
                                    case -953:
                                        v151 := []int{(v3 << ((1) & 63)), 1}
                                        _ = v151
                                    case 412:
                                        v46 += (v46 / (3.14159 + 0.0))
                                    }
                                    var v152 int64 = 9223372036854775807
                                    _ = fi149
                                    _ = v152
                                }
                                v3 = 2147483647
                                _ = v148
                            }
                        } else {
                            var v153 int64 = -723
                            _ = v153
                        }
                        for fi154 := 0; fi154 < 6; fi154++ {
                            fmt.Println(fi154, fi154)
                            if false {
                                v155 := GenS18{F0: (v47 + "\n\t"), F1: (v47 + v47), F2: -330}
                                if ((((v47 + "") + v47) >= v47) || v2) {
                                    if v2 {
                                        v156 := []int{fi154, (v3 & v3), (-(fi154))}
                                        v47 = ((((v47 + v47) + "hello") + v47) + (v47 + v47))
                                        var v157 uint = 4294967295
                                        _ = v156
                                        _ = v157
                                    }
                                    v3 -= (v3 % (((-(fi45))) | 1))
                                    switch "" {
                                    case "\n\t":
                                        v2 = (-128 >= 46)
                                    case "foo bar":
                                        v158 := []int{127, (fi154 % ((fi154) | 1)), fi154, 126}
                                        _ = v158
                                    }
                                    var v159 uint32 = 1
                                    v160 := GenS13{F0: v2}
                                    _ = v159
                                    _ = v160
                                }
                                _ = v155
                            } else {
                                v161 := GenS16{F0: (-(v3)), F1: 84, F2: 2147483647}
                                v47 = (v47 + "a")
                                _ = v161
                            }
                            _ = fi154
                        }
                        v1 = (v1 << ((733) & 63))
                    }
                case "hello":
                    for fi162 := 0; fi162 < 7; fi162++ {
                        v3 = (fi45 - v3)
                        if v2 {
                            v163 := []int{v3, -604, (-(v3)), v3}
                            if len(v163) > 0 {
                                v163[0] = (^(fi45))
                            }
                            _ = v163
                        } else {
                            var v164 uint16 = 255
                            v165 := []int{216, fi45, 539}
                            _ = v164
                            _ = v165
                        }
                        fmt.Println(v3, v1, v1)
                        _ = fi162
                    }
                }
                if (!v2) {
                    v46 -= (0.1 * 1e10)
                }
            }
            _ = v46
            _ = v47
        }
        if v2 {
            fmt.Println(v1)
            fmt.Println(v1)
            var v166 int32 = -2147483648
            _ = v166
        } else {
            if true {
                v167 := GenS1{F0: 1, F1: 2}
                fmt.Println(v3, v3, v1)
                v168 := GenS2{F0: 32767, F1: 42, F2: 255}
                _ = v167
                _ = v168
            }
        }
        var v169 float32 = 1e10
        _ = fi45
        _ = v169
    }
    if (!(203 == 178)) {
        fmt.Println(v2, v2, v1)
        v170 := GenS19{F0: 2, F1: 128}
        v1 = v1
        _ = v170
    }
    for fi171 := 0; fi171 < 13; fi171++ {
        if true {
            var v172 bool = false
            v2 = true
            v1 += (^(v1))
            _ = v172
        } else {
            for fi173 := 0; fi173 < 9; fi173++ {
                v3 = fi173
                v174 := GenS6{F0: 2, F1: "foo bar", F2: false}
                v2 = v2
                v2 = v2
                var v175 uint64 = (^(v1))
                _ = fi173
                _ = v174
                _ = v175
            }
            v176 := []int{(fi171 << ((2) & 63)), fi171, v3}
            v177 := GenS6{F0: 42, F1: "", F2: (42 < -578)}
            _ = v176
            _ = v177
        }
        v2 = v2
        _ = fi171
    }
    for fi178 := 0; fi178 < 9; fi178++ {
        v179 := []int{(^(v3)), 0, (v3 << ((65535) & 63)), 130}
        v180 := GenS18{F0: ("\n\t" + "\n\t"), F1: "", F2: 1}
        _ = fi178
        _ = v179
        _ = v180
    }
    switch "hello" {
    case "foo bar":
        for fi181 := 0; fi181 < 20; fi181++ {
            v2 = v2
            v182 := []int{}
            _ = fi181
            _ = v182
        }
    default:
        var v183 int16 = 0
        _ = v183
    }
    var v184 string = "a"
    var v185 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v184:", v184)
    fmt.Println("v185:", v185)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v184
    _ = v185
    _ = v2
    _ = v3
}
