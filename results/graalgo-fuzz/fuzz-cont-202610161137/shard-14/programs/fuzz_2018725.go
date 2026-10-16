package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 uint8; F1 int; F2 int32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int32; F1 int16 }
type GenS13 struct { F0 bool; F1 uint8; F2 int64 }
type GenS14 struct { F0 uint8; F1 int64 }

func main() {
    var v1 uint = 433
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 0.0
    v4 := []uint32{847}
    var v5 int = 0
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 5; fi8++ {
        if (!false) {
            v2 = v2
            v5 = 0
        } else {
            fmt.Println(v5, v2)
        }
        _ = fi8
    }
    if true {
        var v9 uint64 = v2
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                if (true || true) {
                    if (127 >= 1) {
                        v9 = (v2 << ((v1) & 63))
                    }
                    fmt.Println(fi11, v3)
                }
                _ = fi11
            }
            _ = fi10
        }
        _ = v9
    }
    if ((v5 & -283) > v5) {
        var v12 bool = true
        _ = v12
    }
    fmt.Println(v1, v1)
    var v13 bool = true
    v5 = (v5 & v5)
    for fi14 := 0; fi14 < 10; fi14++ {
        v3 = v3
        _ = fi14
    }
    v6[(v5 * (-(v5)))] = 42
    v15 := v6[int(v2)]
    for fi16 := 0; fi16 < 18; fi16++ {
        switch ("foo bar" + (("" + "") + "\n\t")) {
        case "a":
            var v17 uint = v1
            _ = v17
        case "hello":
            v18 := GenS2{F0: 100, F1: -128, F2: (fi16 ^ (^(fi16)))}
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v13)
    if v13 {
        v3 = (0.1 * v3)
        var v19 bool = (true && (!false))
        v20 := GenS2{F0: 127, F1: 16, F2: v5}
        _ = v19
        _ = v20
    } else {
        var v21 string = "hello"
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS3{F0: -1}
            v5 = (v5 / ((fi22) | 1))
            _ = fi22
            _ = v23
        }
        switch v21 {
        case "foo bar":
            v15 = 341
        case "":
            if v13 {
                switch ((v21 + v21) + "\n\t") {
                case "hello":
                    fmt.Println(v15)
                case "a":
                    var v24 float64 = float64(3.14159)
                    _ = v24
                case "":
                    v25 := GenS4{F0: (v15 / (((v15 >> (((v1 & v1)) & 63))) | 1)), F1: 2147483647, F2: 1}
                    _ = v25
                default:
                    fmt.Println(v15)
                }
                for fi26 := 0; fi26 < 15; fi26++ {
                    if false {
                        v21 = v21
                        v27 := GenS4{F0: v15, F1: -2147483648, F2: 100}
                        _ = v27
                    } else {
                        if v13 {
                            switch ("a" + v21) {
                            case "":
                                v6[-640] = 110
                            case "foo bar":
                                var v28 uint8 = 255
                                _ = v28
                            case "a":
                                switch (((v21 + v21) + v21) + v21) {
                                case "a":
                                    var v29 int8 = -32
                                    _ = v29
                                case "\n\t":
                                    if v13 {
                                        v2 += (^(v2))
                                        for fi30 := 0; fi30 < 8; fi30++ {
                                            v21 = (v21 + (v21 + "hello"))
                                            fmt.Println(fi30, v3)
                                            var v31 uint32 = (^(v15))
                                            _ = fi30
                                            _ = v31
                                        }
                                    } else {
                                        var v32 uint16 = 32767
                                        _ = v32
                                    }
                                default:
                                    for fi33 := 0; fi33 < 4; fi33++ {
                                        if (!v13) {
                                            if v13 {
                                                v13 = (1 != 0)
                                            } else {
                                                v15 += (v15 * (v15 - 628))
                                            }
                                            v34 := GenS3{F0: 0}
                                            v1 += uint(v15)
                                            _ = v34
                                        } else {
                                            var v35 float64 = 3.14159
                                            v5 = int(v1)
                                            _ = v35
                                        }
                                        fmt.Println(v21, v3)
                                        v2 = uint64(v15)
                                        _ = fi33
                                    }
                                }
                            }
                            switch v21 {
                            case "\n\t":
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    v1 = 0
                                    _ = fi36
                                }
                            case "hello":
                                v37 := GenS3{F0: 2147483647}
                                _ = v37
                            case "foo bar":
                                switch 127 {
                                case -1:
                                    if (v3 != float64((1e10 / float32(v3)))) {
                                        v15 = 2147483647
                                    }
                                case 127:
                                    delete(v6, fi26)
                                default:
                                    if v13 {
                                        v21 = "\n\t"
                                        if (!((^(v5)) <= (v5 << (((^(v1))) & 63)))) {
                                            if ((2 <= 191) || v13) {
                                                v15 = v15
                                            }
                                            var v38 string = v21
                                            if (!v13) {
                                                var v39 bool = (!v13)
                                                v2 += 2
                                                var v40 int8 = -10
                                                _ = v39
                                                _ = v40
                                            } else {
                                                switch 32767 {
                                                case -2147483648:
                                                    v41 := v6[(fi26 >> ((uint(v15)) & 63))]
                                                    _ = v41
                                                case 1:
                                                    var v42 int8 = 1
                                                    _ = v42
                                                case 42:
                                                    switch v21 {
                                                    case "a":
                                                        for fi43 := 0; fi43 < 17; fi43++ {
                                                            v5 += 32768
                                                            _ = fi43
                                                        }
                                                    }
                                                default:
                                                    v4 = append(v4, (v15 << ((v1) & 63)))
                                                }
                                                switch 127 {
                                                case 805:
                                                    switch ("a" + v21) {
                                                    case "a":
                                                        for fi44 := 0; fi44 < 6; fi44++ {
                                                            v45 := GenS4{F0: (v15 / ((uint32(v2)) | 1)), F1: -980, F2: 48}
                                                            var v46 uint32 = (v15 << (((v1 + uint(v15))) & 63))
                                                            _ = fi44
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    case "\n\t":
                                                        var v47 bool = v13
                                                        _ = v47
                                                    }
                                                case 32767:
                                                    var v48 uint16 = 719
                                                    _ = v48
                                                default:
                                                    if v13 {
                                                        if v13 {
                                                            v15 = 42
                                                            if (-32768 < -32768) {
                                                                var v49 uint64 = v2
                                                                if v13 {
                                                                    fmt.Println(v38, v13, v3)
                                                                    for fi50 := 0; fi50 < 16; fi50++ {
                                                                        var v51 bool = v13
                                                                        if (!v51) {
                                                                            switch v21 {
                                                                            case "":
                                                                                v49 -= (v2 >> (((v1 ^ v1)) & 63))
                                                                            }
                                                                        }
                                                                        var v52 float32 = -2.25
                                                                        _ = fi50
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                } else {
                                                                    v6[(fi26 | v5)] = v15
                                                                    v38 = v21
                                                                }
                                                                v4 = append(v4, v15)
                                                                _ = v49
                                                            } else {
                                                                for fi53 := 0; fi53 < 1; fi53++ {
                                                                    fmt.Println(v21, v15)
                                                                    _ = fi53
                                                                }
                                                                if (v13 || true) {
                                                                    fmt.Println(v38)
                                                                    var v54 float32 = 0.1
                                                                    var v55 bool = true
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                                v56 := GenS4{F0: (v15 | 72), F1: 564, F2: 1}
                                                                _ = v56
                                                            }
                                                        }
                                                        switch (((v38 + "hello") + ((v21 + "hello") + ("" + v38))) + "\n\t") {
                                                        case "":
                                                            for fi57 := 0; fi57 < 1; fi57++ {
                                                                v15 = uint32(v2)
                                                                var v58 int = 42
                                                                _ = fi57
                                                                _ = v58
                                                            }
                                                        case "hello":
                                                            if v13 {
                                                                v59 := GenS1{F0: 1}
                                                                _ = v59
                                                            } else {
                                                                v5 = 1
                                                            }
                                                        }
                                                    } else {
                                                        v60 := GenS5{F0: (v21 + "a"), F1: "hello"}
                                                        for fi61 := 0; fi61 < 18; fi61++ {
                                                            v6[128] = (v15 - (v15 >> (((^(v1))) & 63)))
                                                            var v62 int8 = 1
                                                            _ = fi61
                                                            _ = v62
                                                        }
                                                        var v63 uint8 = 126
                                                        _ = v60
                                                        _ = v63
                                                    }
                                                }
                                                v4 = append(v4, (v15 & 65535))
                                            }
                                            _ = v38
                                        }
                                        var v64 int64 = 9223372036854775807
                                        _ = v64
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch v5 {
                                            case -247:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        case "foo bar":
                                            for fi65 := 0; fi65 < 12; fi65++ {
                                                if v13 {
                                                    v21 = ((v21 + "foo bar") + "foo bar")
                                                    var v66 uint32 = v15
                                                    v1 -= v1
                                                    _ = v66
                                                } else {
                                                    v15 = 0
                                                    v3 = v3
                                                    _ = len(v4)
                                                }
                                                _ = fi65
                                            }
                                        case "a":
                                            if v13 {
                                                v13 = false
                                                v67 := GenS4{F0: uint32(v2), F1: 0, F2: -1}
                                                for fi68 := 0; fi68 < 15; fi68++ {
                                                    if v13 {
                                                        var v69 uint = (v1 >> (((v1 & v1)) & 63))
                                                        _ = v69
                                                    } else {
                                                        var v70 bool = (v13 || (v15 < (v15 + (v15 >> (((^(v1))) & 63)))))
                                                        if v70 {
                                                            var v71 string = "hello"
                                                            switch 127 {
                                                            case 128:
                                                                v72 := GenS6{F0: 2147483647, F1: 65535}
                                                                _ = v72
                                                            case 576:
                                                                var v73 int32 = 0
                                                                _ = v73
                                                            case 32767:
                                                                var v74 int32 = 1
                                                                v71 = v21
                                                                v74 -= 2147483647
                                                                v4 = append(v4, v15)
                                                                _ = v74
                                                            default:
                                                                if (!true) {
                                                                    if (65535 >= (v1 ^ (v1 ^ (v1 << (((^(v1))) & 63))))) {
                                                                        fmt.Println(v3, fi68)
                                                                        switch fi68 {
                                                                        case -2147483648:
                                                                            switch "" {
                                                                            case "a":
                                                                                if v70 {
                                                                                    v75 := GenS6{F0: 1, F1: 32767}
                                                                                    var v76 uint16 = 255
                                                                                    _ = len(v6)
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                } else {
                                                                                    v77 := v6[int(v15)]
                                                                                    fmt.Println(v71, fi26, v71)
                                                                                    _ = v77
                                                                                }
                                                                            case "":
                                                                                v78 := GenS4{F0: (^(v15)), F1: 2147483647, F2: 72}
                                                                                _ = v78
                                                                            case "\n\t":
                                                                                switch v71 {
                                                                                case "a":
                                                                                    _ = len(v4)
                                                                                case "foo bar":
                                                                                    v79 := GenS7{F0: (v15 + 42), F1: -32768}
                                                                                    _ = v79
                                                                                }
                                                                            default:
                                                                                v15 = v15
                                                                            }
                                                                            var v80 float32 = float32(v3)
                                                                            _ = v80
                                                                        case -457:
                                                                            v13 = (v13 || false)
                                                                        }
                                                                        delete(v6, v5)
                                                                    } else {
                                                                        v81 := GenS5{F0: v71, F1: v21}
                                                                        var v82 uint64 = v2
                                                                        v83 := GenS8{F0: 2147483647}
                                                                        v21 = v71
                                                                        _ = v81
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                } else {
                                                                    if (false || v70) {
                                                                        if v70 {
                                                                            v84 := v6[fi26]
                                                                            _ = v84
                                                                        }
                                                                    }
                                                                    v6[v5] = (v15 + 42)
                                                                    v5 -= (fi26 * (fi26 | (^(fi26))))
                                                                    v85 := GenS9{F0: (^(v2)), F1: (!true)}
                                                                    fmt.Println(v71, v3, v2)
                                                                    v6[(-(fi26))] = uint32(v2)
                                                                    _ = v85
                                                                }
                                                            }
                                                            _ = v71
                                                        } else {
                                                            v6[fi26] = v15
                                                        }
                                                        v15 += 2147483647
                                                        if v13 {
                                                            v13 = true
                                                            var v86 uint64 = (v2 | v2)
                                                            _ = v86
                                                        } else {
                                                            fmt.Println(v15, fi26, v15)
                                                        }
                                                        _ = v70
                                                    }
                                                    _ = fi68
                                                }
                                                _ = v67
                                            }
                                        }
                                        v21 = ((v21 + "a") + v21)
                                        var v87 uint16 = 255
                                        _ = v87
                                    }
                                }
                            default:
                                v3 -= 0.0
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        for fi88 := 0; fi88 < 9; fi88++ {
                            v89 := GenS7{F0: (v15 - v15), F1: 1}
                            var v90 int16 = 0
                            if (v13 || (246 != 0)) {
                                var v91 uint8 = 148
                                switch fi26 {
                                case 1:
                                    switch (v21 + ((v21 + ((v21 + ("\n\t" + v21)) + v21)) + (v21 + (((v21 + "a") + (("a" + (v21 + "a")) + "\n\t")) + v21)))) {
                                    case "hello":
                                        var v92 uint8 = (v91 + uint8(v1))
                                        _ = v92
                                    default:
                                        v6[32768] = (v15 + (v15 * v15))
                                    }
                                    if true {
                                        v93 := GenS9{F0: v2, F1: v13}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if v13 {
                                            var v94 int = int(v1)
                                            _ = v94
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = v15
                                            }
                                        }
                                        _ = v93
                                    }
                                case 424:
                                    v3 = 3.14159
                                }
                                var v95 uint64 = 0
                                _ = v91
                                _ = v95
                            } else {
                                if v13 {
                                    if v13 {
                                        switch fi88 {
                                        case 128:
                                            v96 := v6[fi88]
                                            _ = v96
                                        case 289:
                                            switch ("" + v21) {
                                            case "\n\t":
                                                fmt.Println(fi88, v15)
                                            case "foo bar":
                                                switch fi26 {
                                                case 127:
                                                    fmt.Println(v13, v5)
                                                case 32767:
                                                    v2 += 4294967295
                                                case 1:
                                                    var v97 uint8 = 1
                                                    v2 += (v2 & 4294967295)
                                                    if false {
                                                        if (v3 > v3) {
                                                            if v13 {
                                                                if v13 {
                                                                    fmt.Println(fi88)
                                                                }
                                                                var v98 uint32 = 1
                                                                _ = v98
                                                            } else {
                                                                v99 := GenS9{F0: v2, F1: v13}
                                                                var v100 float32 = 1.5
                                                                _ = v100
                                                                _ = v99
                                                            }
                                                            switch (fi26 << (((v1 << ((42) & 63))) & 63)) {
                                                            case -578:
                                                                var v101 bool = (v15 > v15)
                                                                _ = v101
                                                            case 1:
                                                                var v102 int8 = -1
                                                                _ = v102
                                                            case 9223372036854775807:
                                                                v103 := GenS6{F0: 0, F1: uint16(v15)}
                                                                _ = v103
                                                            }
                                                            delete(v6, (fi26 | 32767))
                                                        }
                                                        fmt.Println(v3)
                                                    } else {
                                                        v3 += (0.1 + (float64(1e10) / v3))
                                                        fmt.Println(v2, v15)
                                                    }
                                                    var v104 uint16 = 65535
                                                    _ = v104
                                                    _ = v97
                                                }
                                            default:
                                                var v105 float32 = -2.25
                                                _ = v105
                                            }
                                        case 9223372036854775807:
                                            v106 := GenS10{F0: 4, F1: (v5 + fi88), F2: 2147483647}
                                            v21 = ""
                                            _ = len(v6)
                                            _ = v106
                                        default:
                                            v3 -= v3
                                        }
                                        _ = len(v6)
                                        v107 := GenS11{F0: int16(fi26)}
                                        _ = v107
                                    }
                                }
                            }
                            _ = fi88
                            _ = v89
                            _ = v90
                        }
                        switch v5 {
                        case 1:
                            var v108 int = 1
                            _ = v108
                        case -922:
                            if ((^(v2)) <= 86) {
                                if (!true) {
                                    v6[(^(fi26))] = v15
                                    var v109 bool = v13
                                    switch (^(v5)) {
                                    case 1:
                                        v15 += (v15 << ((v1) & 63))
                                    default:
                                        for fi110 := 0; fi110 < 11; fi110++ {
                                            v2 += 18446744073709551615
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v111 int = (^(fi26))
                                                _ = v111
                                            default:
                                                var v112 uint8 = 0
                                                var v113 int64 = 926
                                                if (1 < 0) {
                                                    var v114 int8 = 80
                                                    _ = v114
                                                } else {
                                                    v1 = (v1 & (v1 * (v1 % ((18446744073709551615) | 1))))
                                                }
                                                _ = v112
                                                _ = v113
                                            }
                                            _ = len(v4)
                                            var v115 int16 = 1
                                            v5 = (v5 << (((^(v1))) & 63))
                                            _ = fi110
                                            _ = v115
                                        }
                                    }
                                    _ = v109
                                }
                            } else {
                                v3 = ((((v3 + float64((1e10 - float32((v3 + 0.1))))) * v3) + v3) * (v3 + 1.5))
                            }
                        case 4294967295:
                            v2 = (v2 / ((v2) | 1))
                        default:
                            switch v21 {
                            case "\n\t":
                                var v116 string = "hello"
                                _ = v116
                            case "a":
                                var v117 int64 = -32
                                _ = v117
                            case "foo bar":
                                fmt.Println(v5, v13)
                            default:
                                if v13 {
                                    for fi118 := 0; fi118 < 2; fi118++ {
                                        fmt.Println(v15, v3)
                                        switch 635 {
                                        case 128:
                                            if v13 {
                                                v119 := GenS12{F0: 2147483647, F1: -1}
                                                _ = v119
                                            } else {
                                                fmt.Println(v15)
                                                v120 := v6[fi26]
                                                v121 := GenS3{F0: 0}
                                                _ = v120
                                                _ = v121
                                            }
                                        case 32768:
                                            var v122 float32 = (((1e10 / float32(float64((1.5 / 1e10)))) / float32(float64(-2.25))) / 1.0)
                                            v122 -= v122
                                            if v13 {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v122 -= (((((v122 * (v122 - 0.1)) / 0.1) + (0.1 / ((((v122 + (1.5 - v122)) - ((float32(float64((v122 * (v122 * (-2.25 + 0.1))))) / 1.5) / (0.0 / v122))) / v122) / (v122 * -2.25)))) + v122) + v122)
                                            }
                                            _ = v122
                                        }
                                        _ = fi118
                                    }
                                    if v13 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v123 int8 = -128
                                        delete(v6, (fi26 / (((^(v5))) | 1)))
                                        _ = v123
                                    } else {
                                        if v13 {
                                            fmt.Println(v15, v15, v2)
                                            var v124 uint = 1
                                            _ = v124
                                        }
                                    }
                                }
                            }
                        }
                    }
                    v125 := GenS13{F0: ((127 == 100) && true), F1: 0, F2: -933}
                    _ = fi26
                    _ = v125
                }
            } else {
                for fi126 := 0; fi126 < 16; fi126++ {
                    v127 := GenS14{F0: 159, F1: 1}
                    var v128 int32 = 600
                    v6[32768] = v15
                    _ = fi126
                    _ = v127
                    _ = v128
                }
            }
        default:
            v13 = (9223372036854775807 > -807)
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
