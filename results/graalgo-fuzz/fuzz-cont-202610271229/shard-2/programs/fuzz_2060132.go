package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 uint16 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint; F1 int }
type GenS11 struct { F0 int32; F1 uint32 }
type GenS12 struct { F0 int32; F1 uint16; F2 int8 }
type GenS13 struct { F0 uint32; F1 int8 }
type GenS14 struct { F0 uint32; F1 uint8; F2 uint }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 string }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]bool)
    var v3 float64 = 3.14159
    var v4 int64 = -1
    v5 := make(map[int]int)
    v6 := v5[1]
    var v7 bool = true
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 0
        v7 = (!(((!(0.0 >= v3)) || false) && ((v4 / ((0) | 1)) < (v4 - (v4 - -1)))))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 42, F2: 545}
    var v11 float32 = 0.0
    var v12 uint16 = 1
    if ((326 < (^(v4))) || v7) {
        fmt.Println(v7, v6)
    }
    fmt.Println(v7)
    v6 -= -988
    v13 := []int{}
    switch "a" {
    case "":
        if v7 {
            v14 := GenS2{F0: 2147483647, F1: 484, F2: (^(v12))}
            _ = v14
        } else {
            v13 = append(v13, v6)
            var v15 int8 = 127
            if (v7 && true) {
                for fi16 := 0; fi16 < 1; fi16++ {
                    v13 = append(v13, 32767)
                    v1 = v12
                    v7 = true
                    _ = fi16
                }
                if (!true) {
                    if ((0 >= 12) && v7) {
                        if false {
                            switch (-(v6)) {
                            case -9223372036854775808:
                                switch (v6 >> ((362) & 63)) {
                                case -2147483648:
                                    var v17 uint = 2
                                    _ = v17
                                default:
                                    _ = len(v2)
                                }
                            }
                            v15 = v15
                            if false {
                                v11 = ((3.14159 / (((((v11 / -2.25) - (v11 * v11)) - 3.14159) + 3.14159) - v11)) / ((v11 / v11) + ((v11 * v11) + (v11 - v11))))
                            } else {
                                switch ("a" + ("hello" + "\n\t")) {
                                case "foo bar":
                                    for fi18 := 0; fi18 < 3; fi18++ {
                                        if false {
                                            v19 := GenS3{F0: 42, F1: 1}
                                            _ = v19
                                        } else {
                                            fmt.Println(fi18, v1, v11)
                                            switch -254 {
                                            case 4294967295:
                                                for fi20 := 0; fi20 < 7; fi20++ {
                                                    switch fi18 {
                                                    case 1:
                                                        if (!true) {
                                                            var v21 int16 = 0
                                                            v1 = 255
                                                            _ = v21
                                                        }
                                                    case 127:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            fmt.Println(v3, v12)
                                                            v22 := GenS4{F0: "foo bar", F1: v4}
                                                            _ = v22
                                                        case "a":
                                                            fmt.Println(v1, v3)
                                                        case "\n\t":
                                                            if v7 {
                                                                v23 := GenS4{F0: ("\n\t" + "foo bar"), F1: (v4 - v4)}
                                                                _ = v23
                                                            }
                                                        default:
                                                            var v24 float32 = v11
                                                            v6 += (v6 % ((v6) | 1))
                                                            var v25 string = ("foo bar" + "a")
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                    case 9223372036854775807:
                                                        if (!v7) {
                                                            v26 := GenS4{F0: ("\n\t" + "a"), F1: 0}
                                                            fmt.Println(fi18)
                                                            _ = v26
                                                        } else {
                                                            var v27 uint = 42
                                                            _ = v27
                                                        }
                                                    default:
                                                        v15 = (v15 | v15)
                                                    }
                                                    _ = fi20
                                                }
                                            case 2147483647:
                                                v6 -= v6
                                            default:
                                                for fi28 := 0; fi28 < 12; fi28++ {
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    _ = fi28
                                                }
                                            }
                                        }
                                        var v29 int32 = 2147483647
                                        var v30 float32 = v11
                                        _ = fi18
                                        _ = v29
                                        _ = v30
                                    }
                                case "hello":
                                    _ = len(v2)
                                default:
                                    var v31 uint8 = 127
                                    _ = v31
                                }
                                v7 = false
                            }
                        }
                        v12 += (v12 & 32767)
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v3 = v3
                            _ = fi32
                        }
                    } else {
                        var v33 int8 = (v15 * 0)
                        switch "foo bar" {
                        case "foo bar":
                            switch "" {
                            case "hello":
                                v12 -= 2
                            case "a":
                                var v34 string = "\n\t"
                                _ = v34
                            case "\n\t":
                                v1 -= 65535
                            }
                        case "hello":
                            switch ((("" + "\n\t") + "a") + (("a" + "a") + ("hello" + "foo bar"))) {
                            case "a":
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    var v36 int = -2147483648
                                    for fi37 := 0; fi37 < 14; fi37++ {
                                        var v38 uint8 = 109
                                        _ = fi37
                                        _ = v38
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                            case "hello":
                                v11 += (-2.25 + (v11 * v11))
                            case "foo bar":
                                v39 := GenS5{F0: v6}
                                _ = v39
                            }
                        default:
                            var v40 string = "\n\t"
                            _ = v40
                        }
                        _ = v33
                    }
                    v5[v6] = v6
                    var v41 uint64 = 307
                    _ = v41
                } else {
                    var v42 float32 = ((1.5 - (v11 * ((-2.25 - -2.25) * 0.0))) - 1e10)
                    v1 -= v12
                    _ = v42
                }
                var v43 bool = (v7 && (v3 == v3))
                _ = v43
            } else {
                var v44 uint64 = 1
                for fi45 := 0; fi45 < 19; fi45++ {
                    if v7 {
                        for fi46 := 0; fi46 < 1; fi46++ {
                            v1 += 2
                            if (v6 < v6) {
                                if (v1 < 65535) {
                                    switch (("" + "hello") + "a") {
                                    case "foo bar":
                                        if false {
                                            v3 -= float64(v11)
                                            v13 = append(v13, fi45)
                                        } else {
                                            v1 = 255
                                            v11 = -2.25
                                            switch "hello" {
                                            case "":
                                                var v47 int16 = 0
                                                _ = v47
                                            case "hello":
                                                switch (fi46 & (v6 - (fi46 / (((^(v6))) | 1)))) {
                                                case 777:
                                                    v12 = (v1 << ((573) & 63))
                                                case 4294967295:
                                                    v48 := GenS6{F0: 65535}
                                                    v5[(^(fi45))] = fi45
                                                    var v49 int8 = -1
                                                    _ = v48
                                                    _ = v49
                                                case 9223372036854775807:
                                                    var v50 string = ""
                                                    _ = v50
                                                }
                                            case "\n\t":
                                                var v51 bool = true
                                                _ = v51
                                            default:
                                                v11 -= v11
                                            }
                                        }
                                    case "a":
                                        fmt.Println(v11, v4)
                                    }
                                    v44 += (^(v44))
                                } else {
                                    var v52 uint64 = (^(v44))
                                    _ = v52
                                }
                            } else {
                                if (599 < int16(fi46)) {
                                    v15 = (^(v15))
                                    v53 := GenS7{F0: 0}
                                    var v54 uint = 4294967295
                                    _ = v53
                                    _ = v54
                                } else {
                                    var v55 int8 = (v15 + 59)
                                    _ = v55
                                }
                                var v56 int = 128
                                _ = v56
                            }
                            var v57 string = ""
                            _ = fi46
                            _ = v57
                        }
                        switch (v6 & (fi45 / ((4294967295) | 1))) {
                        case 85:
                            switch ("a" + (("foo bar" + "hello") + ("hello" + "foo bar"))) {
                            case "foo bar":
                                if v7 {
                                    v58 := GenS8{F0: 1, F1: v12}
                                    _ = v58
                                } else {
                                    if (((1 > -32768) || (v7 || v7)) && v7) {
                                        switch "hello" {
                                        case "foo bar":
                                            for fi59 := 0; fi59 < 20; fi59++ {
                                                if true {
                                                    v11 -= (1.5 / v11)
                                                    v3 = v3
                                                    if (v15 >= (v15 / ((42) | 1))) {
                                                        var v60 uint = 0
                                                        var v61 uint8 = 2
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        v5[(fi59 & (fi59 % ((fi45) | 1)))] = int(v4)
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                        if false {
                                                            var v62 uint16 = 2
                                                            _ = v62
                                                        }
                                                    }
                                                } else {
                                                    switch v6 {
                                                    case 1:
                                                        if (42 != 1) {
                                                            v63 := GenS5{F0: 71}
                                                            var v64 string = ("a" + "a")
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            v44 -= 18446744073709551615
                                                        }
                                                    }
                                                    switch v6 {
                                                    case 9223372036854775807:
                                                        v1 -= (v12 << ((1) & 63))
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if true {
                                                        var v65 uint64 = (v44 << ((1000) & 63))
                                                        v13 = append(v13, fi45)
                                                        var v66 uint64 = (v65 / ((v44) | 1))
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                default:
                                                    if (42 == 592) {
                                                        var v67 int16 = 32767
                                                        v44 = 42
                                                        _ = v67
                                                    } else {
                                                        fmt.Println(v12)
                                                    }
                                                }
                                                v15 -= (v15 - (v15 << ((115) & 63)))
                                                _ = fi59
                                            }
                                        case "\n\t":
                                            if (v44 == 974) {
                                                var v68 uint16 = v12
                                                if (v4 > 42) {
                                                    var v69 uint16 = v1
                                                    _ = v69
                                                }
                                                _ = v68
                                            } else {
                                                var v70 int8 = (v15 + v15)
                                                var v71 uint = 4294967295
                                                _ = v70
                                                _ = v71
                                            }
                                        case "":
                                            var v72 bool = v7
                                            _ = v72
                                        default:
                                            switch (-(fi45)) {
                                            case 127:
                                                var v73 uint16 = v12
                                                _ = v73
                                            default:
                                                for fi74 := 0; fi74 < 7; fi74++ {
                                                    if false {
                                                        if (2147483648 > uint32(v44)) {
                                                            fmt.Println(v44, v11)
                                                        } else {
                                                            v12 += (^(v12))
                                                        }
                                                        var v75 int32 = int32(v44)
                                                        _ = v75
                                                    } else {
                                                        var v76 uint = 42
                                                        if true {
                                                            if (v11 > 1.5) {
                                                                var v77 uint64 = (v44 | 1)
                                                                _ = v77
                                                            }
                                                        } else {
                                                            var v78 uint8 = 153
                                                            if (float32(v3) != v11) {
                                                                fmt.Println(v76, v7, v44)
                                                            } else {
                                                                if false {
                                                                    v12 += v12
                                                                    fmt.Println(v76, v3)
                                                                } else {
                                                                    v79 := GenS1{F0: 352, F1: 2, F2: 2147483648}
                                                                    _ = v79
                                                                }
                                                            }
                                                            if (891 > 500) {
                                                                if v7 {
                                                                    var v80 uint = (^(v76))
                                                                    v81 := GenS7{F0: 2147483647}
                                                                    v82 := GenS9{F0: 0, F1: 279}
                                                                    _ = v80
                                                                    _ = v81
                                                                    _ = v82
                                                                } else {
                                                                    v4 = -1
                                                                    fmt.Println(v12)
                                                                    v11 = float32(v3)
                                                                }
                                                            } else {
                                                                v83 := GenS10{F0: (^(v76)), F1: (fi45 << ((v76) & 63))}
                                                                v12 = v12
                                                                var v84 uint = (v76 % ((v76) | 1))
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                            _ = v78
                                                        }
                                                        v85 := GenS5{F0: (v6 >> (((v76 ^ (v76 + 255))) & 63))}
                                                        _ = v76
                                                        _ = v85
                                                    }
                                                    v12 = 565
                                                    v86 := GenS6{F0: 867}
                                                    _ = fi74
                                                    _ = v86
                                                }
                                            }
                                        }
                                        if v7 {
                                            v4 = v4
                                            var v87 int16 = -1
                                            var v88 bool = ((v4 % ((9223372036854775807) | 1)) < 9223372036854775807)
                                            _ = v87
                                            _ = v88
                                        } else {
                                            v89 := GenS3{F0: int(v4), F1: 127}
                                            v7 = v7
                                            if false {
                                                _ = len(v13)
                                                var v90 uint = 878
                                                if true {
                                                    if ((-415 != v6) || v7) {
                                                        fmt.Println(fi45, v90, fi45)
                                                        v91 := GenS11{F0: -1, F1: 177}
                                                        for fi92 := 0; fi92 < 14; fi92++ {
                                                            if (63 <= 2) {
                                                                v13 = append(v13, (^(v6)))
                                                            }
                                                            var v93 int = (^(fi45))
                                                            switch "hello" {
                                                            case "hello":
                                                                if v7 {
                                                                    if (!v7) {
                                                                        var v94 float32 = v11
                                                                        v93 -= fi45
                                                                        _ = v94
                                                                    } else {
                                                                        var v95 uint64 = v44
                                                                        _ = v95
                                                                    }
                                                                } else {
                                                                    if len(v13) >= 2 {
                                                                        v13 = v13[1:len(v13)]
                                                                    }
                                                                    if true {
                                                                        if v7 {
                                                                            if len(v13) >= 2 {
                                                                                v13 = v13[1:len(v13)]
                                                                            }
                                                                            delete(v5, 4294967295)
                                                                            fmt.Println(v44, v15)
                                                                        }
                                                                        var v96 uint64 = v44
                                                                        _ = v96
                                                                    }
                                                                }
                                                            case "a":
                                                                switch fi45 {
                                                                case 9223372036854775807:
                                                                    v97 := GenS3{F0: 727, F1: 42}
                                                                    _ = v97
                                                                case -1:
                                                                    fmt.Println(v90, fi45, v3)
                                                                    switch (fi92 % (((fi45 | 644)) | 1)) {
                                                                    case 0:
                                                                        if v7 {
                                                                            v1 = (v12 & uint16(v90))
                                                                            v15 = (^(v15))
                                                                            v98 := GenS6{F0: 169}
                                                                            _ = v98
                                                                        } else {
                                                                            v4 = v4
                                                                            fmt.Println(fi92, v7, v6)
                                                                            v11 = v11
                                                                        }
                                                                    case 42:
                                                                        v99 := GenS12{F0: 2147483647, F1: v1, F2: (v15 - (^(v15)))}
                                                                        _ = v99
                                                                    case -842:
                                                                        var v100 float64 = 1.5
                                                                        fmt.Println(v6, v6, v11)
                                                                        _ = v100
                                                                    default:
                                                                        if v7 {
                                                                            if ((v15 == (v15 - (v15 % (((v15 + (v15 / ((1) | 1)))) | 1)))) && v7) {
                                                                                if true {
                                                                                    if v7 {
                                                                                        if true {
                                                                                            var v101 uint64 = v44
                                                                                            v102 := GenS13{F0: 2147483648, F1: (v15 >> ((65535) & 63))}
                                                                                            v103 := v5[2147483647]
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                        }
                                                                                        v90 = v90
                                                                                        if true {
                                                                                            v104 := GenS14{F0: 0, F1: 0, F2: v90}
                                                                                            _ = v104
                                                                                        } else {
                                                                                            if ((v15 % (((v15 >> (((v90 << ((v90) & 63))) & 63))) | 1)) < -128) {
                                                                                                fmt.Println(v15)
                                                                                                var v105 int = int(v15)
                                                                                                var v106 string = ("a" + (((("\n\t" + "foo bar") + "foo bar") + "") + "a"))
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            }
                                                                                            v107 := GenS14{F0: 21, F1: 1, F2: v90}
                                                                                            var v108 uint32 = 572
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        }
                                                                                    } else {
                                                                                        v44 = (v44 & v44)
                                                                                        var v109 int8 = (v15 << (((v90 ^ (v90 ^ (v90 ^ v90)))) & 63))
                                                                                        _ = v109
                                                                                    }
                                                                                    v110 := GenS15{F0: v7}
                                                                                    var v111 string = "a"
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                                fmt.Println(fi92, fi45, v7)
                                                                            }
                                                                            v112 := GenS5{F0: int(v1)}
                                                                            v7 = (2 <= 1)
                                                                            _ = v112
                                                                        } else {
                                                                            fmt.Println(v44, fi92, v90)
                                                                        }
                                                                    }
                                                                    v6 = (^(fi45))
                                                                    if len(v13) > 0 {
                                                                        v13[0] = (fi92 | (v93 | (-(v6))))
                                                                    }
                                                                    if v7 {
                                                                        var v113 bool = false
                                                                        _ = len(v2)
                                                                        if false {
                                                                            var v114 string = ((("" + "") + ("" + "a")) + "")
                                                                            v115 := GenS13{F0: 266, F1: (v15 >> ((v90) & 63))}
                                                                            v15 = (v15 / ((-1) | 1))
                                                                            _ = v114
                                                                            _ = v115
                                                                        }
                                                                        _ = v113
                                                                    }
                                                                    delete(v2, ("" + ((("hello" + "a") + "") + "a")))
                                                                case 4294967295:
                                                                    v90 = 136
                                                                default:
                                                                    switch "a" {
                                                                    case "a":
                                                                        v90 += v90
                                                                    case "\n\t":
                                                                        v116 := GenS11{F0: 42, F1: 2147483647}
                                                                        _ = v116
                                                                    default:
                                                                        _ = len(v5)
                                                                        v117 := GenS16{F0: "foo bar"}
                                                                        fmt.Println(v1, v12, v1)
                                                                        _ = v117
                                                                    }
                                                                }
                                                            default:
                                                                delete(v5, fi45)
                                                            }
                                                            _ = fi92
                                                            _ = v93
                                                        }
                                                        _ = v91
                                                    } else {
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                    }
                                                    switch ("\n\t" + "") {
                                                    case "":
                                                        var v118 uint32 = 2147483648
                                                        _ = v118
                                                    case "\n\t":
                                                        var v119 uint8 = uint8(v12)
                                                        _ = v119
                                                    case "foo bar":
                                                        var v120 int = (^(fi45))
                                                        _ = v120
                                                    }
                                                    var v121 uint = (^(v90))
                                                    _ = v121
                                                } else {
                                                    var v122 float64 = v3
                                                    for fi123 := 0; fi123 < 4; fi123++ {
                                                        v124 := GenS10{F0: 384, F1: (-(v6))}
                                                        var v125 int = (^(v6))
                                                        if ((((v11 - (0.1 * v11)) * (v11 + ((((1.5 * v11) - v11) / float32(v3)) / v11))) + (v11 * 3.14159)) == v11) {
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                            var v126 int64 = 1
                                                            v44 += 885
                                                            _ = v126
                                                        }
                                                        var v127 float64 = v122
                                                        _ = fi123
                                                        _ = v124
                                                        _ = v125
                                                        _ = v127
                                                    }
                                                    _ = v122
                                                }
                                                _ = v90
                                            } else {
                                                v128 := GenS9{F0: 204, F1: (v4 + v4)}
                                                var v129 bool = false
                                                var v130 bool = v129
                                                _ = v128
                                                _ = v129
                                                _ = v130
                                            }
                                            _ = v89
                                        }
                                        var v131 uint = 998
                                        _ = v131
                                    }
                                    if (!v7) {
                                        for fi132 := 0; fi132 < 11; fi132++ {
                                            var v133 uint8 = 255
                                            fmt.Println(v15, v133, v44)
                                            _ = fi132
                                            _ = v133
                                        }
                                        for fi134 := 0; fi134 < 3; fi134++ {
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            _ = fi134
                                        }
                                        v135 := GenS8{F0: 4294967295, F1: (v1 >> ((42) & 63))}
                                        _ = v135
                                    }
                                    v6 += fi45
                                }
                            case "\n\t":
                                v12 = 1
                            }
                        case 32767:
                            v4 += (v4 << ((42) & 63))
                        case 1:
                            fmt.Println(v6, v1, v44)
                        default:
                            var v136 uint64 = v44
                            _ = v136
                        }
                        for fi137 := 0; fi137 < 9; fi137++ {
                            if true {
                                v12 -= 0
                                fmt.Println(v4)
                                var v138 float32 = 3.14159
                                v12 += (^(v1))
                                var v139 int64 = (v4 << ((4294967295) & 63))
                                _ = v138
                                _ = v139
                            }
                            _ = fi137
                        }
                    }
                    _ = fi45
                }
                _ = v44
            }
            _ = v15
        }
    case "hello":
        for fi140 := 0; fi140 < 2; fi140++ {
            for fi141 := 0; fi141 < 15; fi141++ {
                v142 := GenS2{F0: 1, F1: -2147483648, F2: v1}
                v3 = ((((1.5 * float64((1.5 + float32(v3)))) * (float64(1.5) / 1.5)) * v3) - v3)
                var v143 int = v6
                v144 := GenS13{F0: 4294967295, F1: 60}
                _ = fi141
                _ = v142
                _ = v143
                _ = v144
            }
            _ = fi140
        }
    }
    var v145 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v145:", v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
