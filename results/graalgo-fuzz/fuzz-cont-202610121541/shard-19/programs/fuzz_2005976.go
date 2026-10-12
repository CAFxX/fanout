package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
type GenS5 struct { F0 string; F1 int16; F2 uint8 }
type GenS6 struct { F0 bool; F1 int64; F2 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 int64; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 int; F1 int16 }
type GenS13 struct { F0 int32; F1 bool; F2 uint }
type GenS14 struct { F0 int32; F1 uint16; F2 string }
type GenS15 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint16; F1 int32; F2 int32 }
type GenS18 struct { F0 uint; F1 uint32 }
type GenS19 struct { F0 int64; F1 uint64 }
type GenS20 struct { F0 uint16; F1 string }
func genFn0(p0 uint16, p1 uint8) uint64 {
    v4 := []int{}
    if false {
        v4 = append(v4, 4294967295)
        if ("" <= ("hello" + "hello")) {
            if false {
                switch "a" {
                case "hello":
                    p0 -= p0
                case "foo bar":
                    if (((("a" + ("" + ("hello" + "a"))) + ("" + ("\n\t" + ("a" + ("" + "foo bar"))))) + "foo bar") < "") {
                        var v5 uint8 = 42
                        if (uint64(p0) != 524) {
                            for fi6 := 0; fi6 < 13; fi6++ {
                                var v7 float64 = (-2.25 / 1.5)
                                fmt.Println(fi6)
                                _ = fi6
                                _ = v7
                            }
                        } else {
                            for fi8 := 0; fi8 < 11; fi8++ {
                                for fi9 := 0; fi9 < 12; fi9++ {
                                    var v10 uint32 = 2
                                    v10 = v10
                                    _ = fi9
                                    _ = v10
                                }
                                v11 := GenS1{F0: -648, F1: fi8}
                                v5 -= p1
                                _ = fi8
                                _ = v11
                            }
                            if true {
                                var v12 int16 = -32768
                                _ = v12
                            }
                            var v13 string = ("a" + "a")
                            _ = v13
                        }
                        _ = v5
                    }
                }
                var v14 uint = 2
                _ = v14
            } else {
                var v15 uint32 = 65535
                if false {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        p1 = 192
                        switch (fi16 << ((18446744073709551615) & 63)) {
                        case 2147483647:
                            if false {
                                if true {
                                    v15 = uint32(p1)
                                    v17 := GenS1{F0: 1, F1: 128}
                                    if false {
                                        for fi18 := 0; fi18 < 13; fi18++ {
                                            fmt.Println(p1)
                                            p1 += 128
                                            p1 = (p1 << ((0) & 63))
                                            _ = fi18
                                        }
                                        p0 += p0
                                        var v19 uint16 = 490
                                        _ = v19
                                    }
                                    _ = v17
                                }
                                v20 := GenS1{F0: 1000, F1: 32768}
                                _ = v20
                            } else {
                                v21 := GenS2{F0: -260, F1: 0}
                                var v22 uint = 1
                                _ = v21
                                _ = v22
                            }
                        default:
                            v23 := GenS3{F0: 255}
                            _ = v23
                        }
                        _ = fi16
                    }
                } else {
                    var v24 float32 = float32(0.1)
                    v24 -= ((-2.25 + v24) / v24)
                    v25 := GenS4{F0: 434, F1: 4294967295, F2: -243}
                    _ = v24
                    _ = v25
                }
                _ = v15
            }
        } else {
            if true {
                p0 -= (p0 >> ((18446744073709551615) & 63))
                for fi26 := 0; fi26 < 15; fi26++ {
                    if len(v4) > 0 {
                        v4[0] = (fi26 * fi26)
                    }
                    _ = fi26
                }
            }
            fmt.Println(p0, p1)
        }
        var v27 int16 = 985
        _ = v27
    } else {
        p0 = p0
        if true {
            for fi28 := 0; fi28 < 6; fi28++ {
                p1 = p1
                _ = fi28
            }
            v29 := GenS4{F0: uint64(p0), F1: 255, F2: 128}
            var v30 int16 = 42
            _ = v29
            _ = v30
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 969
}
func genFn1(p0 int16) uint64 {
    for fi31 := 0; fi31 < 3; fi31++ {
        p0 += p0
        v32 := GenS4{F0: 18446744073709551615, F1: 264, F2: (fi31 * (fi31 & (fi31 & (fi31 >> ((42) & 63)))))}
        _ = fi31
        _ = v32
    }
    p0 += (p0 * 42)
    v33 := GenS4{F0: 4294967295, F1: 18446744073709551615, F2: -1}
    fmt.Println(p0)
    var v34 int8 = 127
    _ = p0
    _ = v33
    _ = v34
    return 2
}

func main() {
    var v1 uint32 = 86
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    v35 := []int{}
    var v36 int = 9223372036854775807
    v37 := GenS5{F0: "\n\t", F1: 32767, F2: 140}
    var v38 uint32 = 65535
    for fi39 := 0; fi39 < 1; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            v36 = v36
            _ = fi40
        }
        for fi41 := 0; fi41 < 15; fi41++ {
            var v42 uint = 691
            var v43 uint32 = v38
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 uint = 4294967295
        _ = fi39
        _ = v44
    }
    if (!true) {
        fmt.Println(v1, v1)
        if false {
            var v45 uint16 = 32767
            fmt.Println(v36, v1)
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        if false {
            if true {
                for fi47 := 0; fi47 < 18; fi47++ {
                    switch "foo bar" {
                    case "\n\t":
                        v38 -= (v1 ^ (^(v1)))
                    }
                    v38 += (v1 >> ((4294967295) & 63))
                    if true {
                        if ((true && (!false)) || (true || (968 > 0))) {
                            var v48 uint16 = 104
                            _ = v48
                        } else {
                            v36 = (v36 ^ -1)
                            var v49 uint64 = uint64(v38)
                            _ = v49
                        }
                    } else {
                        v1 -= (v1 | v38)
                        switch 0 {
                        case 9223372036854775807:
                            v50 := GenS6{F0: (!false), F1: 46, F2: ""}
                            _ = v50
                        case 128:
                            v38 = v1
                        case -900:
                            v51 := GenS1{F0: 0, F1: (v36 / ((fi46) | 1))}
                            _ = v51
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                v1 += 65535
                            default:
                                var v52 uint64 = 803
                                _ = v52
                            }
                        }
                        if len(v35) > 0 {
                            v35[0] = 32768
                        }
                    }
                    _ = fi47
                }
                switch "" {
                case "a":
                    var v53 uint16 = 0
                    _ = v53
                case "\n\t":
                    _ = len(v35)
                }
            } else {
                v54 := GenS7{F0: 42}
                _ = v54
            }
            if (true && (74 < 151)) {
                delete(v2, ("" + "a"))
                switch (fi46 - (^(fi46))) {
                case 42:
                    var v55 string = ("a" + "")
                    _ = v55
                case 128:
                    v38 -= (v1 >> ((42) & 63))
                default:
                    v56 := v2["foo bar"]
                    _ = v56
                }
                v57 := GenS8{F0: (v38 * 699), F1: 1, F2: -177}
                _ = v57
            }
            var v58 uint8 = 0
            _ = v58
        }
        switch 32767 {
        case -926:
            if false {
                v36 -= v36
                v38 = v38
                var v59 uint32 = v1
                _ = v59
            } else {
                switch "hello" {
                case "foo bar":
                    for fi60 := 0; fi60 < 8; fi60++ {
                        v61 := GenS5{F0: ("\n\t" + ""), F1: 653, F2: 70}
                        var v62 int16 = 78
                        v63 := GenS5{F0: ("\n\t" + "\n\t"), F1: (v62 * (-(v62))), F2: 128}
                        _ = fi60
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                case "\n\t":
                    for fi64 := 0; fi64 < 20; fi64++ {
                        fmt.Println(v1, v36)
                        v65 := GenS2{F0: 708, F1: -1}
                        v36 = (fi64 * (v36 / ((9223372036854775807) | 1)))
                        _ = fi64
                        _ = v65
                    }
                default:
                    for fi66 := 0; fi66 < 4; fi66++ {
                        _ = len(v2)
                        v38 += (v38 | v38)
                        switch (-(v36)) {
                        case 4294967295:
                            v38 = (v38 | 4294967295)
                        case -1:
                            var v67 int16 = 1000
                            _ = v67
                        }
                        _ = fi66
                    }
                }
                if false {
                    v38 = v38
                } else {
                    var v68 string = "a"
                    _ = v68
                }
                v69 := GenS2{F0: -1, F1: -1}
                _ = v69
            }
        }
        for fi70 := 0; fi70 < 6; fi70++ {
            var v71 int16 = 1
            v72 := GenS9{F0: "\n\t"}
            _ = fi70
            _ = v71
            _ = v72
        }
        _ = fi46
    }
    v1 = (v1 | 5)
    for fi73 := 0; fi73 < 17; fi73++ {
        switch "" {
        case "hello":
            fmt.Println(fi73)
        case "":
            if len(v35) > 0 {
                v35[0] = fi73
            }
        case "\n\t":
            switch 32768 {
            case 125:
                v74 := GenS10{F0: 1, F1: 215}
                _ = v74
            case -2147483648:
                _ = len(v2)
            }
        }
        var v75 uint16 = 487
        _ = fi73
        _ = v75
    }
    v35 = append(v35, v36)
    for fi76 := 0; fi76 < 4; fi76++ {
        if false {
            for fi77 := 0; fi77 < 17; fi77++ {
                if (!true) {
                    v78 := v3[""]
                    var v79 uint8 = 255
                    _ = v78
                    _ = v79
                } else {
                    if false {
                        var v80 float64 = 1e10
                        v81 := GenS8{F0: v38, F1: 42, F2: 9223372036854775807}
                        _ = v80
                        _ = v81
                    } else {
                        var v82 uint64 = genFn0(42, 242)
                        v36 = fi76
                        _ = v82
                    }
                    if (92 == -128) {
                        v1 = v1
                        v83 := GenS11{F0: 554, F1: v38}
                        _ = v83
                    } else {
                        if false {
                            v3[""] = (v38 ^ 42)
                            fmt.Println(v1)
                            v1 -= v1
                            var v84 string = "a"
                            var v85 bool = (!(30 < 127))
                            _ = v84
                            _ = v85
                        } else {
                            v35 = append(v35, (v36 & 32767))
                        }
                        var v86 uint32 = v38
                        var v87 uint = 1
                        var v88 uint32 = (v86 & 469)
                        _ = v86
                        _ = v87
                        _ = v88
                    }
                }
                v1 += 0
                v36 = v36
                v89 := GenS12{F0: (fi77 | (fi76 ^ (fi77 & v36))), F1: -526}
                v90 := GenS13{F0: 1, F1: true, F2: 11}
                _ = fi77
                _ = v89
                _ = v90
            }
            var v91 bool = (18446744073709551615 < 65535)
            _ = v91
        }
        _ = fi76
    }
    for fi92 := 0; fi92 < 20; fi92++ {
        v93 := GenS14{F0: -1, F1: 65535, F2: "hello"}
        for fi94 := 0; fi94 < 14; fi94++ {
            var v95 uint16 = 32767
            switch "" {
            case "":
                if false {
                    v96 := GenS6{F0: true, F1: -746, F2: "a"}
                    fmt.Println(v95, v1, fi92)
                    _ = v96
                } else {
                    v38 += 4294967295
                    _ = len(v35)
                }
            case "foo bar":
                var v97 uint16 = v95
                _ = v97
            }
            v98 := GenS9{F0: "\n\t"}
            v99 := GenS15{F0: 1, F1: 1, F2: 18446744073709551615}
            v38 = v1
            if false {
                v3[("\n\t" + "")] = v1
            }
            if true {
                v1 -= v38
                v100 := GenS10{F0: 464, F1: -9223372036854775808}
                _ = v100
            }
            if (false && (!(!(-128 < -95)))) {
                v2["a"] = (fi92 ^ (fi94 * (v36 & fi92)))
                v101 := GenS16{F0: false}
                var v102 int64 = 0
                _ = v101
                _ = v102
            } else {
                v103 := GenS17{F0: 2, F1: -214, F2: -1}
                v104 := GenS18{F0: 18446744073709551615, F1: (^(v1))}
                var v105 float32 = ((3.14159 - 0.0) / (0.1 + float32(0.1)))
                _ = v103
                _ = v104
                _ = v105
            }
            _ = fi94
            _ = v95
            _ = v98
            _ = v99
        }
        var v106 uint8 = 232
        _ = fi92
        _ = v106
        _ = v93
    }
    var v107 float64 = 0.0
    if false {
        v38 -= (v38 | v38)
        var v108 int16 = int16(v38)
        if ((0 <= -1) && (v107 != v107)) {
            var v109 uint8 = 80
            for fi110 := 0; fi110 < 20; fi110++ {
                v107 += (v107 + 3.14159)
                if (2 < 67) {
                    v111 := GenS3{F0: 42}
                    if false {
                        if false {
                            fmt.Println(v108)
                            v36 += 4294967295
                            var v112 int16 = v108
                            _ = v112
                        }
                        if true {
                            var v113 int8 = -99
                            _ = v113
                        } else {
                            fmt.Println(fi110)
                        }
                        if (-2147483648 < 0) {
                            for fi114 := 0; fi114 < 17; fi114++ {
                                v109 = (v109 - (v109 & (v109 << ((204) & 63))))
                                _ = fi114
                            }
                        }
                    }
                    _ = v111
                }
                _ = fi110
            }
            switch ("hello" + "a") {
            case "\n\t":
                var v115 int64 = 0
                _ = v115
            default:
                delete(v2, "foo bar")
            }
            _ = v109
        } else {
            v116 := GenS6{F0: (1 <= 979), F1: 0, F2: "\n\t"}
            var v117 int = v36
            _ = v116
            _ = v117
        }
        _ = v108
    } else {
        var v118 int8 = 111
        _ = v118
    }
    fmt.Println(v1, v38)
    if (2 != 0) {
        v119 := GenS14{F0: 42, F1: 834, F2: (("" + "hello") + "\n\t")}
        if ((42 <= uint(v1)) || true) {
            for fi120 := 0; fi120 < 16; fi120++ {
                switch ("a" + (("" + ("foo bar" + (("foo bar" + "hello") + (("\n\t" + "") + "foo bar")))) + ((("foo bar" + "") + "a") + "a"))) {
                case "foo bar":
                    v1 = 2
                case "a":
                    for fi121 := 0; fi121 < 16; fi121++ {
                        var v122 int8 = 0
                        if true {
                            v123 := GenS19{F0: -688, F1: 42}
                            _ = v123
                        }
                        v124 := v3[("\n\t" + "")]
                        _ = fi121
                        _ = v122
                        _ = v124
                    }
                default:
                    if (3.14159 > 0.1) {
                        switch "a" {
                        case "foo bar":
                            var v125 float32 = 0.1
                            _ = v125
                        }
                        var v126 float32 = (-2.25 + -2.25)
                        _ = v126
                    } else {
                        if true {
                            var v127 uint = 885
                            v1 = v38
                            _ = v127
                        }
                    }
                }
                v3[""] = (v1 - v38)
                _ = fi120
            }
            switch (v36 << ((113) & 63)) {
            case 2147483647:
                v2["a"] = (-(v36))
            case 42:
                for fi128 := 0; fi128 < 13; fi128++ {
                    if true {
                        if true {
                            var v129 uint8 = 2
                            _ = v129
                        } else {
                            fmt.Println(fi128, v107)
                        }
                        v1 += 13
                    }
                    v1 -= 0
                    _ = fi128
                }
            default:
                v130 := GenS20{F0: 0, F1: "\n\t"}
                _ = v130
            }
        } else {
            if true {
                fmt.Println(v36, v107)
                if (0.0 <= ((v107 / v107) + v107)) {
                    v107 = v107
                }
            }
            var v131 int8 = -128
            _ = len(v35)
            _ = v131
        }
        _ = v119
    }
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
