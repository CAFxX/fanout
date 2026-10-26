package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 bool }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 uint32; F1 int32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 int32; F2 bool }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 string; F1 int; F2 int32 }
type GenS15 struct { F0 int; F1 uint32; F2 int64 }
type GenS16 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 bool; F1 int }
type GenS19 struct { F0 uint8; F1 bool; F2 int64 }
func genFn0() bool {
    if (0.1 >= ((1.5 + ((float32((3.14159 + (3.14159 / 1.0))) / 1.5) - 1e10)) * -2.25)) {
        v7 := []int{4294967295, 128, 4294967295}
        _ = v7
    }
    if true {
        if (1.5 > 1e10) {
            v8 := []int{-9223372036854775808, 127, -883}
            var v9 uint32 = 2147483647
            var v10 uint64 = 18446744073709551615
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 int32 = -2147483648
            _ = v11
        }
    } else {
        if false {
            var v12 string = "hello"
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := []int{(fi13 >> ((830) & 63)), fi13}
                _ = fi13
                _ = v14
            }
            v12 = (("a" + v12) + "hello")
            _ = v12
        } else {
            if (false && (false && true)) {
                var v15 int = 736
                _ = v15
            } else {
                v16 := GenS1{F0: -9223372036854775808, F1: "hello", F2: (-9223372036854775808 != 0)}
                var v17 int8 = 0
                v18 := GenS1{F0: -1, F1: "foo bar", F2: (9223372036854775807 > -9223372036854775808)}
                _ = v16
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: 1, F1: ("hello" + ""), F2: (42 <= 0)}
            _ = v19
        }
        v20 := []int{128, 32768, -9223372036854775808, 224}
        var v21 uint8 = 127
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 9223372036854775807, F1: "", F2: (42 > 2147483647)}
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: 42, F1: (("\n\t" + "") + "foo bar")}
        _ = fi23
        _ = v24
    }
    _ = v22
    return true
}
func genFn1() uint {
    var v25 bool = true
    v26 := GenS3{F0: -2147483648}
    _ = v25
    _ = v26
    return 97
}

func main() {
    var v1 uint32 = 42
    var v2 uint64 = 73
    v3 := make(map[uint32]uint32)
    var v4 uint = 806
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 0
    var v27 uint8 = 135
    switch 128 {
    case 4294967295:
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 uint16 = 42
            for fi30 := 0; fi30 < 19; fi30++ {
                switch "a" {
                case "\n\t":
                    v31 := GenS1{F0: 0, F1: ("a" + ("a" + "\n\t")), F2: true}
                    _ = v31
                default:
                    v5[v1] = (^(v1))
                }
                v2 -= v2
                fmt.Println(fi30)
                _ = fi30
            }
            for fi32 := 0; fi32 < 19; fi32++ {
                _ = len(v3)
                v33 := []int{}
                v34 := GenS4{F0: v27, F1: 61}
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi28
            _ = v29
        }
    case -9223372036854775808:
        var v35 float32 = (1e10 - ((-2.25 + -2.25) * 1e10))
        _ = v35
    case 909:
        if false {
            fmt.Println(v6, v1)
            switch -1 {
            case 2:
                v36 := GenS4{F0: (v27 | v27), F1: -128}
                _ = v36
            case -588:
                v1 = (v1 | 8)
            case 4294967295:
                switch int(v4) {
                case 32768:
                    v37 := []int{-610, 228, 32768}
                    _ = v37
                case 4294967295:
                    v27 = (v27 & (v27 << ((genFn1()) & 63)))
                case 9223372036854775807:
                    fmt.Println(v27)
                }
            default:
                var v38 float64 = 3.14159
                _ = v38
            }
            v4 -= genFn1()
        }
    default:
        var v39 int = 127
        _ = v39
    }
    v2 -= v2
    fmt.Println(v1, v4)
    v40 := GenS1{F0: 549, F1: "", F2: false}
    fmt.Println(v4, v6)
    v41 := []int{924, -2147483648}
    switch ("foo bar" + ("foo bar" + "hello")) {
    case "hello":
        if false {
            for fi42 := 0; fi42 < 3; fi42++ {
                v27 = v27
                _ = fi42
            }
            var v43 float32 = (1e10 / 3.14159)
            var v44 uint32 = (v1 * (v1 ^ (^(v1))))
            _ = v43
            _ = v44
        }
    case "":
        v45 := GenS4{F0: (v27 % (((v27 | (^(v27)))) | 1)), F1: 42}
        _ = v45
    case "foo bar":
        var v46 float32 = (float32(0.1) * 3.14159)
        _ = v46
    }
    v47 := GenS5{F0: 128}
    if true {
        v48 := GenS2{F0: 1000, F1: "hello"}
        if (-2.25 == 3.14159) {
            v27 -= v27
            if false {
                for fi49 := 0; fi49 < 4; fi49++ {
                    v50 := GenS2{F0: int16(v2), F1: "foo bar"}
                    v1 += v1
                    _ = fi49
                    _ = v50
                }
                if (uint32(v27) >= (v1 % (((^(v1))) | 1))) {
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                    for fi51 := 0; fi51 < 16; fi51++ {
                        for fi52 := 0; fi52 < 19; fi52++ {
                            switch int(v1) {
                            case 4294967295:
                                var v53 uint16 = v6
                                _ = v53
                            case 894:
                                switch "a" {
                                case "foo bar":
                                    var v54 int8 = 1
                                    _ = v54
                                case "\n\t":
                                    if len(v41) > 0 {
                                        v41[0] = fi51
                                    }
                                default:
                                    if (617 == v2) {
                                        v4 = v4
                                    }
                                    var v55 int64 = 9223372036854775807
                                    _ = v55
                                }
                                var v56 int = (fi52 + 2147483647)
                                if false {
                                    v1 += (v1 & v1)
                                    v57 := GenS6{F0: (v2 + 646), F1: 90}
                                    _ = v57
                                }
                                if (genFn1() >= 883) {
                                    var v58 uint64 = (^(v2))
                                    var v59 int32 = 0
                                    v59 = (v59 - 0)
                                    _ = v58
                                    _ = v59
                                }
                                _ = v56
                            case 127:
                                var v60 uint8 = v27
                                _ = v60
                            default:
                                v1 -= (^(v1))
                            }
                            if false {
                                if true {
                                    v61 := GenS7{F0: (v27 ^ (v27 >> (((v4 % ((v4) | 1))) & 63)))}
                                    _ = v61
                                } else {
                                    if (305 != -418) {
                                        v62 := GenS3{F0: (fi51 / ((fi52) | 1))}
                                        v63 := GenS5{F0: -9223372036854775808}
                                        if (false && false) {
                                            switch int(v6) {
                                            case -2147483648:
                                                var v64 bool = (48 < 100)
                                                _ = v64
                                            case 4294967295:
                                                if false {
                                                    var v65 int = 1
                                                    if len(v41) >= 2 {
                                                        v41 = v41[1:len(v41)]
                                                    }
                                                    v66 := GenS1{F0: -444, F1: ("\n\t" + "foo bar"), F2: false}
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    v67 := v5[v1]
                                                    v68 := GenS5{F0: (fi52 ^ (-(fi52)))}
                                                    _ = v67
                                                    _ = v68
                                                }
                                            }
                                            v69 := GenS8{F0: 32767, F1: (^(v1))}
                                            _ = v69
                                        } else {
                                            if (uint32(v2) >= (v1 % ((v1) | 1))) {
                                                var v70 uint8 = 1
                                                switch (fi51 | fi52) {
                                                case 233:
                                                    _ = len(v3)
                                                    v71 := v3[(v1 | v1)]
                                                    v1 -= v71
                                                    v1 = 26
                                                    if ("foo bar" >= "hello") {
                                                        v72 := GenS8{F0: v6, F1: (v71 >> ((uint(v2)) & 63))}
                                                        switch ("hello" + ("a" + "")) {
                                                        case "foo bar":
                                                            var v73 int = fi51
                                                            _ = v73
                                                        case "a":
                                                            v74 := GenS6{F0: v2, F1: 1}
                                                            _ = v74
                                                        }
                                                        _ = v72
                                                    } else {
                                                        v71 += (v1 * 2)
                                                        v1 += uint32(v4)
                                                        if (0.0 > (0.0 + float32((1.5 * 0.0)))) {
                                                            var v75 uint64 = uint64(v71)
                                                            switch 2147483647 {
                                                            case -1:
                                                                switch fi52 {
                                                                case -9223372036854775808:
                                                                    v76 := GenS8{F0: v6, F1: v1}
                                                                    _ = v76
                                                                }
                                                                switch (fi52 << (((v4 >> ((49) & 63))) & 63)) {
                                                                case -9223372036854775808:
                                                                    var v77 float64 = float64((0.0 * 0.0))
                                                                    _ = v77
                                                                case 1:
                                                                    v71 += v1
                                                                case 42:
                                                                    v70 -= v27
                                                                }
                                                                switch int(v4) {
                                                                case 128:
                                                                    var v78 uint64 = v2
                                                                    _ = v78
                                                                case 77:
                                                                    if (((v4 != uint(v2)) && true) && false) {
                                                                        fmt.Println(v70, v71, v75)
                                                                    } else {
                                                                        var v79 uint = v4
                                                                        fmt.Println(v4, v27, v70)
                                                                        v80 := GenS2{F0: 32767, F1: "\n\t"}
                                                                        var v81 string = ((((((("foo bar" + "a") + "\n\t") + ("\n\t" + ("a" + ("a" + ("\n\t" + (("" + "a") + "\n\t")))))) + ("foo bar" + (("" + (("" + ("a" + ((("" + "hello") + ("a" + ("foo bar" + ("\n\t" + "hello")))) + "a"))) + "hello")) + "foo bar"))) + ("a" + ("a" + "foo bar"))) + "foo bar") + ("" + "hello"))
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                    }
                                                                case -9223372036854775808:
                                                                    switch ((("foo bar" + "\n\t") + "foo bar") + "") {
                                                                    case "":
                                                                        fmt.Println(v4, v4)
                                                                    case "\n\t":
                                                                        v4 = 18446744073709551615
                                                                    case "foo bar":
                                                                        v5[(v1 >> ((genFn1()) & 63))] = (v1 << ((uint(v27)) & 63))
                                                                    }
                                                                }
                                                                var v82 uint64 = (v75 & v75)
                                                                var v83 uint8 = v27
                                                                _ = v82
                                                                _ = v83
                                                            case 32768:
                                                                v27 = (v27 * 34)
                                                            case -212:
                                                                v3[(v1 | v1)] = (v71 % (((v1 | v71)) | 1))
                                                            default:
                                                                switch (("" + "foo bar") + "hello") {
                                                                case "a":
                                                                    v71 -= uint32(v75)
                                                                }
                                                            }
                                                            v3[v71] = v1
                                                            _ = v75
                                                        } else {
                                                            if false {
                                                                if (!(v1 != uint32(v70))) {
                                                                    var v84 string = "hello"
                                                                    _ = v84
                                                                } else {
                                                                    var v85 int64 = -641
                                                                    _ = v85
                                                                }
                                                                if (1 == -128) {
                                                                    v86 := GenS4{F0: (v27 >> (((^(v4))) & 63)), F1: int8(v1)}
                                                                    _ = v86
                                                                } else {
                                                                    var v87 uint = (^(v4))
                                                                    v4 -= genFn1()
                                                                    _ = v87
                                                                }
                                                                v71 += v1
                                                                var v88 uint64 = (^(v2))
                                                                v71 += 2
                                                                _ = v88
                                                            } else {
                                                                v4 = v4
                                                                v6 = 320
                                                                v6 += v6
                                                            }
                                                            v71 = (v1 / ((2147483648) | 1))
                                                            var v89 float64 = (-2.25 - float64(float32(1.5)))
                                                            _ = v89
                                                        }
                                                    }
                                                    v90 := GenS9{F0: (^(v2)), F1: 0}
                                                    _ = v71
                                                    _ = v90
                                                case 4294967295:
                                                    _ = len(v41)
                                                case -861:
                                                    fmt.Println(v4)
                                                default:
                                                    if (false && true) {
                                                        if len(v41) > 0 {
                                                            v41[0] = -1
                                                        }
                                                    } else {
                                                        if true {
                                                            var v91 int8 = 100
                                                            fmt.Println(v70, v27, v4)
                                                            fmt.Println(v2, v6, fi51)
                                                            _ = v91
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    if ((v6 & (^(v6))) <= (^(v6))) {
                                                                        v70 -= v70
                                                                    } else {
                                                                        if false {
                                                                            if (-1 >= -32768) {
                                                                                if len(v41) >= 2 {
                                                                                    v41 = v41[1:len(v41)]
                                                                                }
                                                                                v92 := GenS10{F0: (v1 - uint32(v70)), F1: 2147483647}
                                                                                _ = v92
                                                                            } else {
                                                                                var v93 int = 32768
                                                                                var v94 uint = uint(v2)
                                                                                _ = v93
                                                                                _ = v94
                                                                            }
                                                                            v95 := GenS11{F0: uint64(v70)}
                                                                            v96 := GenS6{F0: (v2 | v2), F1: -2147483648}
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                        fmt.Println(v70, fi51, fi52)
                                                                    }
                                                                    var v97 int32 = 268
                                                                    _ = v97
                                                                }
                                                                var v98 int32 = 971
                                                                _ = v98
                                                            } else {
                                                                var v99 uint = (v4 + (v4 % ((genFn1()) | 1)))
                                                                v1 -= (v1 - v1)
                                                                _ = v99
                                                            }
                                                            v1 -= (v1 ^ (v1 >> ((42) & 63)))
                                                            switch fi52 {
                                                            case 1:
                                                                switch (fi51 ^ (-(fi52))) {
                                                                case 42:
                                                                    var v100 float64 = 1e10
                                                                    _ = v100
                                                                case 4294967295:
                                                                    v101 := GenS12{F0: -2147483648, F1: 42, F2: (false || false)}
                                                                    _ = v101
                                                                case -188:
                                                                    if true {
                                                                        var v102 int64 = 0
                                                                        _ = v102
                                                                    } else {
                                                                        var v103 uint16 = v6
                                                                        v27 -= 229
                                                                        _ = v103
                                                                    }
                                                                default:
                                                                    fmt.Println(fi52, v4, v27)
                                                                }
                                                            default:
                                                                var v104 float32 = ((1.5 / float32(-2.25)) * (((1e10 + float32((3.14159 - (0.0 / (1.5 + ((0.1 / (-2.25 + 1e10)) * (-2.25 / 1.5))))))) / (1.5 / (3.14159 / 1.0))) - 0.0))
                                                                _ = v104
                                                            }
                                                        }
                                                        var v105 int32 = 42
                                                        v41 = append(v41, -9223372036854775808)
                                                        switch ("\n\t" + ("foo bar" + "a")) {
                                                        case "":
                                                            v106 := GenS13{F0: (v4 >> ((v4) & 63))}
                                                            _ = v106
                                                        case "foo bar":
                                                            v107 := GenS14{F0: ("\n\t" + "a"), F1: fi52, F2: (-(v105))}
                                                            _ = v107
                                                        }
                                                        if (!true) {
                                                            v108 := GenS15{F0: fi52, F1: uint32(v6), F2: int64(v4)}
                                                            v2 -= (v2 >> (((v4 & (v4 >> ((v4) & 63)))) & 63))
                                                            _ = v108
                                                        } else {
                                                            v1 -= v1
                                                            if (1.5 > 0.0) {
                                                                switch -2147483648 {
                                                                case 119:
                                                                    v109 := GenS16{F0: 1, F1: uint16(v70), F2: (v27 + 225)}
                                                                    _ = v109
                                                                case 9223372036854775807:
                                                                    v5[65535] = v1
                                                                case -1:
                                                                    if ("" != ("" + "foo bar")) {
                                                                        var v110 float64 = (float64(float32(0.1)) + 3.14159)
                                                                        var v111 float64 = (0.1 - v110)
                                                                        _ = v110
                                                                        _ = v111
                                                                    }
                                                                }
                                                                v70 = v70
                                                                v1 += v1
                                                            } else {
                                                                var v112 uint16 = v6
                                                                v113 := GenS7{F0: v70}
                                                                if (!false) {
                                                                    v114 := GenS17{F0: -819}
                                                                    _ = v114
                                                                } else {
                                                                    if (42 != fi52) {
                                                                        var v115 uint8 = (^(v27))
                                                                        var v116 int32 = 0
                                                                        if false {
                                                                            if false {
                                                                                fmt.Println(v4, v115, v1)
                                                                            }
                                                                            var v117 int8 = 0
                                                                            _ = v117
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(v4, v6)
                                                                                v2 += uint64(v6)
                                                                                if true {
                                                                                    var v118 int8 = int8(fi52)
                                                                                    delete(v5, v1)
                                                                                    v118 = v118
                                                                                    _ = v118
                                                                                }
                                                                            } else {
                                                                                v1 = (v1 & (v1 ^ 2))
                                                                            }
                                                                            if (true || (v2 >= (v2 & v2))) {
                                                                                if true {
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        _ = len(v41)
                                                                                    case "\n\t":
                                                                                        v41 = append(v41, (fi51 % (((fi51 << ((v4) & 63))) | 1)))
                                                                                    case "a":
                                                                                        if (0 > -9223372036854775808) {
                                                                                            v119 := GenS5{F0: 128}
                                                                                            _ = v119
                                                                                        } else {
                                                                                            var v120 bool = true
                                                                                            var v121 uint8 = (v70 >> (((v4 + (v4 * v4))) & 63))
                                                                                            _ = v120
                                                                                            _ = v121
                                                                                        }
                                                                                    }
                                                                                    var v122 uint16 = uint16(v2)
                                                                                    _ = v122
                                                                                }
                                                                                v115 += (v27 << ((genFn1()) & 63))
                                                                                v123 := GenS17{F0: -317}
                                                                                _ = len(v41)
                                                                                v2 = uint64(v27)
                                                                                _ = v123
                                                                            }
                                                                        }
                                                                        _ = v115
                                                                        _ = v116
                                                                    } else {
                                                                        fmt.Println(v6)
                                                                        if len(v41) > 0 {
                                                                            v41[0] = 1
                                                                        }
                                                                    }
                                                                    v6 -= (v112 & 175)
                                                                    fmt.Println(v112, v105, fi51)
                                                                }
                                                                v5[v1] = (v1 & v1)
                                                                _ = v112
                                                                _ = v113
                                                            }
                                                        }
                                                        _ = v105
                                                    }
                                                }
                                                if true {
                                                    v1 = (v1 >> ((143) & 63))
                                                    v6 += (v6 ^ v6)
                                                    fmt.Println(v2, v2)
                                                }
                                                _ = v70
                                            } else {
                                                switch 2147483647 {
                                                case 9223372036854775807:
                                                    _ = len(v3)
                                                case 147:
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v124 int16 = -497
                                                        _ = v124
                                                    case "\n\t":
                                                        fmt.Println(fi51)
                                                    case "":
                                                        fmt.Println(fi51, v1, fi52)
                                                    }
                                                }
                                            }
                                            var v125 float64 = -2.25
                                            var v126 float32 = 0.0
                                            var v127 float64 = 0.0
                                            _ = v125
                                            _ = v126
                                            _ = v127
                                        }
                                        _ = v62
                                        _ = v63
                                    } else {
                                        if false {
                                            if true {
                                                var v128 bool = ((!true) || false)
                                                var v129 uint32 = 65535
                                                v4 += (v4 & uint(v1))
                                                _ = v128
                                                _ = v129
                                            }
                                            v2 += (v2 >> (((^(v4))) & 63))
                                        }
                                    }
                                    fmt.Println(v27)
                                }
                            } else {
                                v4 = 1
                                var v130 uint = uint(v6)
                                _ = v130
                            }
                            v131 := v3[(v1 + 65535)]
                            _ = fi52
                            _ = v131
                        }
                        _ = fi51
                    }
                }
            } else {
                v4 = v4
            }
            v1 += v1
        } else {
            v4 = (v4 >> ((genFn1()) & 63))
            v41 = append(v41, -9223372036854775808)
            var v132 uint = 255
            _ = v132
        }
        _ = v48
    } else {
        v3[2147483647] = (v1 << ((uint(v6)) & 63))
        fmt.Println(v2)
        v1 += (^(v1))
    }
    if (-2147483648 != 432) {
        v27 -= v27
        v6 -= v6
    } else {
        var v133 uint16 = v6
        switch 519 {
        case -326:
            _ = len(v3)
        case -505:
            v134 := GenS11{F0: (v2 + 1)}
            _ = v134
        case 667:
            v135 := GenS18{F0: (!(true && (false || true))), F1: -271}
            _ = v135
        default:
            if (((!(-418 <= -544)) || (-424 <= int(v27))) && true) {
                v136 := GenS19{F0: (v27 % ((v27) | 1)), F1: (147 < (v1 | uint32(v4))), F2: 0}
                v133 -= (v133 | uint16(v4))
                v41 = append(v41, 128)
                _ = v136
            }
        }
        _ = v133
    }
    v137 := GenS10{F0: (v1 % ((uint32(v6)) | 1)), F1: -1}
    v4 += genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v137
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v47
    _ = v5
    _ = v6
}
