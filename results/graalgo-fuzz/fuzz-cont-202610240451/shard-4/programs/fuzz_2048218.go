package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 int8; F1 int16; F2 uint16 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint64; F1 string }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool; F1 int; F2 bool }
type GenS13 struct { F0 uint32; F1 uint16 }
type GenS14 struct { F0 uint64; F1 uint32 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 string; F1 uint8 }
type GenS19 struct { F0 bool; F1 int }
type GenS20 struct { F0 int; F1 int8 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS23 struct { F0 string }
type GenS24 struct { F0 uint32; F1 uint64; F2 string }
type GenS25 struct { F0 bool }
type GenS26 struct { F0 int; F1 uint8; F2 int64 }
type GenS27 struct { F0 uint8 }
type GenS28 struct { F0 int16; F1 int8; F2 uint }
type GenS29 struct { F0 string }
type GenS30 struct { F0 uint32 }
type GenS31 struct { F0 bool; F1 uint64 }
type GenS32 struct { F0 int8 }
type GenS33 struct { F0 int32; F1 int; F2 int }
type GenS34 struct { F0 uint8; F1 int64 }
type GenS35 struct { F0 int16 }
type GenS36 struct { F0 uint; F1 uint }
type GenS37 struct { F0 int64 }
type GenS38 struct { F0 int8 }
type GenS39 struct { F0 uint32; F1 int8; F2 int8 }
type GenS40 struct { F0 int; F1 int32 }
type GenS41 struct { F0 bool; F1 uint8 }
type GenS42 struct { F0 uint; F1 uint64 }
type GenS43 struct { F0 uint32 }
type GenS44 struct { F0 uint32 }
type GenS45 struct { F0 uint }

func main() {
    var v1 int64 = 225
    var v2 uint64 = 696
    v3 := []uint32{1, 4294967295, 2147483648}
    v4 := make(map[int]bool)
    var v5 uint32 = 42
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6)
        var v7 uint32 = v5
        if (1e10 < (-2.25 + -2.25)) {
            if ((!(((v1 & v1) < (^(v1))) || (true || false))) && true) {
                _ = len(v3)
            }
            v8 := GenS1{F0: 954}
            var v9 uint = 819
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                v1 += v1
                if (0 == int16(v5)) {
                    switch "" {
                    case "hello":
                        v11 := GenS2{F0: 44, F1: 42, F2: 255}
                        _ = v11
                    default:
                        if true {
                            var v12 int = (-(fi10))
                            _ = v12
                        } else {
                            v13 := GenS1{F0: 2}
                            switch "foo bar" {
                            case "foo bar":
                                v14 := GenS1{F0: 1}
                                _ = v14
                            }
                            if true {
                                switch int(v7) {
                                case -1:
                                    var v15 int = fi10
                                    _ = v15
                                case 127:
                                    if (!false) {
                                        fmt.Println(v1, v5, v5)
                                    }
                                }
                            } else {
                                delete(v4, (fi6 << ((65535) & 63)))
                                v16 := GenS3{F0: 4294967295, F1: 8}
                                var v17 uint32 = v5
                                var v18 string = (("\n\t" + (("" + "hello") + ("" + "hello"))) + ((("foo bar" + "\n\t") + ("a" + ("a" + ((((((("foo bar" + "\n\t") + ((("hello" + ("\n\t" + ("" + "hello"))) + ("hello" + ((((("a" + (("foo bar" + "a") + "a")) + "hello") + "hello") + ("hello" + "hello")) + ("foo bar" + ((("hello" + ("\n\t" + "foo bar")) + "foo bar") + (("a" + "\n\t") + "")))))) + "foo bar")) + "\n\t") + ("foo bar" + "\n\t")) + (("a" + "") + (("a" + ("\n\t" + "\n\t")) + "a"))) + "a") + "")))) + "a"))
                                switch ("a" + (v18 + v18)) {
                                case "\n\t":
                                    var v19 string = (v18 + v18)
                                    _ = v19
                                case "a":
                                    v2 = (v2 + (^(v2)))
                                case "":
                                    switch (fi6 % (((fi10 - fi10)) | 1)) {
                                    case 2147483647:
                                        v7 = v7
                                    case 4294967295:
                                        v1 = (v1 | v1)
                                        v2 = (^(v2))
                                        var v20 float64 = 0.1
                                        _ = v20
                                    case 32768:
                                        v21 := GenS1{F0: v7}
                                        v2 -= v2
                                        v1 -= (^(v1))
                                        _ = v21
                                    }
                                }
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            _ = v13
                        }
                        v2 += uint64(v5)
                    }
                    fmt.Println(fi10, v2, fi10)
                } else {
                    fmt.Println(v2, v7, v2)
                    var v22 uint8 = 127
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v22
                }
                _ = fi10
            }
            v23 := GenS1{F0: uint32(v2)}
            _ = v23
        }
        _ = fi6
        _ = v7
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v1 = v1
        for fi25 := 0; fi25 < 6; fi25++ {
            switch 666 {
            case 928:
                v3 = append(v3, 0)
            default:
                var v26 uint = 42
                _ = v26
            }
            switch fi24 {
            case -527:
                if (!(float32(1e10) >= (0.0 - 1.5))) {
                    if false {
                        fmt.Println(v2, fi25, v5)
                        var v27 uint16 = 1
                        var v28 uint8 = 42
                        v27 = 42
                        v3 = append(v3, v5)
                        v1 = -9223372036854775808
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(v1)
                    if len(v3) > 0 {
                        v3[0] = uint32(v2)
                    }
                }
            case 0:
                var v29 uint32 = v5
                _ = v29
            case 15:
                var v30 int8 = -76
                _ = v30
            default:
                v31 := GenS1{F0: 2147483648}
                _ = v31
            }
            _ = len(v4)
            _ = fi25
        }
        _ = fi24
    }
    var v32 int8 = -55
    var v33 uint64 = (v2 >> ((65535) & 63))
    var v34 uint64 = 468
    if (false && (!(true && true))) {
        switch 594 {
        case 987:
            var v35 int8 = v32
            _ = v35
        }
        if ("hello" == ("" + "\n\t")) {
            if true {
                v36 := v4[0]
                var v37 uint64 = 987
                _ = v36
                _ = v37
            }
        } else {
            var v38 uint8 = 35
            v39 := GenS1{F0: 4294967295}
            if true {
                v40 := GenS3{F0: 42, F1: 18446744073709551615}
                if false {
                    if (int64(v38) == (v1 & 9223372036854775807)) {
                        if (!(!(!true))) {
                            for fi41 := 0; fi41 < 6; fi41++ {
                                var v42 int32 = 623
                                _ = fi41
                                _ = v42
                            }
                            v38 = 127
                            delete(v4, 911)
                        } else {
                            for fi43 := 0; fi43 < 6; fi43++ {
                                for fi44 := 0; fi44 < 2; fi44++ {
                                    v34 = (v2 >> ((uint(v38)) & 63))
                                    fmt.Println(v32, v33, fi44)
                                    v45 := GenS3{F0: 255, F1: 42}
                                    v33 -= (v2 | v34)
                                    _ = fi44
                                    _ = v45
                                }
                                var v46 uint = 2
                                if true {
                                    if false {
                                        fmt.Println(v33)
                                    } else {
                                        v47 := GenS4{F0: v33, F1: 875, F2: 4294967295}
                                        v46 = v46
                                        _ = v47
                                    }
                                    var v48 int = fi43
                                    v49 := GenS1{F0: (v5 / ((921) | 1))}
                                    _ = v48
                                    _ = v49
                                }
                                _ = fi43
                                _ = v46
                            }
                            switch "" {
                            case "a":
                                fmt.Println(v32, v38)
                            case "":
                                if (true || true) {
                                    if true {
                                        if false {
                                            if (false || true) {
                                                v38 -= (v38 << ((65535) & 63))
                                            } else {
                                                if true {
                                                    if false {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if (0.1 == 3.14159) {
                                                            v50 := v4[32768]
                                                            var v51 int32 = -1
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                    v33 = 2
                                                } else {
                                                    v52 := GenS5{F0: 1, F1: v5}
                                                    if false {
                                                        for fi53 := 0; fi53 < 10; fi53++ {
                                                            v34 = (^(v33))
                                                            fmt.Println(v5)
                                                            fmt.Println(v1)
                                                            _ = fi53
                                                        }
                                                        var v54 int32 = 42
                                                        switch ("foo bar" + "foo bar") {
                                                        case "\n\t":
                                                            var v55 string = ""
                                                            _ = v55
                                                        case "hello":
                                                            v56 := v4[int(v38)]
                                                            _ = v56
                                                        case "":
                                                            for fi57 := 0; fi57 < 10; fi57++ {
                                                                if (v32 == 18) {
                                                                    for fi58 := 0; fi58 < 1; fi58++ {
                                                                        if (false && (v38 > (^(v38)))) {
                                                                            v3 = append(v3, (v5 - (v5 + 2147483648)))
                                                                            var v59 uint32 = (v5 + v5)
                                                                            v60 := GenS6{F0: 0}
                                                                            _ = v59
                                                                            _ = v60
                                                                        } else {
                                                                            v61 := GenS7{F0: uint16(v38), F1: ""}
                                                                            _ = v61
                                                                        }
                                                                        _ = fi58
                                                                    }
                                                                }
                                                                v62 := GenS4{F0: (v34 + 0), F1: 0, F2: 42}
                                                                for fi63 := 0; fi63 < 6; fi63++ {
                                                                    _ = len(v3)
                                                                    v64 := GenS8{F0: v32, F1: 1, F2: 32767}
                                                                    _ = fi63
                                                                    _ = v64
                                                                }
                                                                _ = fi57
                                                                _ = v62
                                                            }
                                                        }
                                                        _ = v54
                                                    } else {
                                                        v1 -= (-(v1))
                                                        if ((1e10 / 0.1) != -2.25) {
                                                            if (!false) {
                                                                v38 = v38
                                                                for fi65 := 0; fi65 < 11; fi65++ {
                                                                    switch (fi65 >> ((332) & 63)) {
                                                                    case -399:
                                                                        for fi66 := 0; fi66 < 6; fi66++ {
                                                                            var v67 int16 = 139
                                                                            if (!(false && true)) {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                v68 := GenS9{F0: v1, F1: false}
                                                                                v33 = 42
                                                                                _ = v68
                                                                            } else {
                                                                                v67 = 1
                                                                                var v69 int32 = -1
                                                                                if false {
                                                                                    if (!(8 >= v32)) {
                                                                                        v4[fi65] = (false || (false || false))
                                                                                        v33 = (^(v34))
                                                                                    }
                                                                                    fmt.Println(v38)
                                                                                    var v70 int = -598
                                                                                    _ = v70
                                                                                } else {
                                                                                    v38 = v38
                                                                                    switch fi66 {
                                                                                    case 9223372036854775807:
                                                                                        var v71 int = fi66
                                                                                        _ = v71
                                                                                    case -737:
                                                                                        v72 := GenS9{F0: (v1 & -74), F1: ((^(v5)) == (^(v5)))}
                                                                                        _ = v72
                                                                                    case 68:
                                                                                        v34 -= v34
                                                                                    default:
                                                                                        delete(v4, 308)
                                                                                    }
                                                                                }
                                                                                _ = v69
                                                                            }
                                                                            _ = fi66
                                                                            _ = v67
                                                                        }
                                                                    default:
                                                                        v38 -= (v38 % ((uint8(v34)) | 1))
                                                                    }
                                                                    v5 -= 65535
                                                                    v3 = append(v3, (v5 ^ (^(v5))))
                                                                    _ = fi65
                                                                }
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v73 string = ""
                                                                    _ = v73
                                                                case "foo bar":
                                                                    v38 = v38
                                                                }
                                                            }
                                                            v33 += (v2 / ((v33) | 1))
                                                        }
                                                        v74 := GenS4{F0: v33, F1: 1000, F2: 65535}
                                                        _ = v74
                                                    }
                                                    v32 = int8(v38)
                                                    _ = v52
                                                }
                                                switch -9223372036854775808 {
                                                case 128:
                                                    for fi75 := 0; fi75 < 6; fi75++ {
                                                        var v76 uint8 = (^(v38))
                                                        for fi77 := 0; fi77 < 9; fi77++ {
                                                            var v78 int64 = (v1 | (v1 >> ((4294967295) & 63)))
                                                            v79 := v4[(fi75 * (fi75 % ((fi77) | 1)))]
                                                            var v80 uint8 = 2
                                                            _ = fi77
                                                            _ = v78
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        v32 = v32
                                                        _ = fi75
                                                        _ = v76
                                                    }
                                                default:
                                                    var v81 uint8 = 255
                                                    _ = v81
                                                }
                                            }
                                            v82 := GenS10{F0: 1, F1: ("\n\t" + "a")}
                                            _ = v82
                                        } else {
                                            v83 := GenS4{F0: (v2 << ((65535) & 63)), F1: -32768, F2: (^(v5))}
                                            fmt.Println(v1, v33)
                                            var v84 string = "foo bar"
                                            _ = v83
                                            _ = v84
                                        }
                                        for fi85 := 0; fi85 < 5; fi85++ {
                                            switch fi85 {
                                            case -9223372036854775808:
                                                var v86 bool = false
                                                _ = v86
                                            case 127:
                                                v87 := GenS11{F0: v1}
                                                _ = v87
                                            case 476:
                                                v4[0] = true
                                            default:
                                                for fi88 := 0; fi88 < 16; fi88++ {
                                                    v33 = 18446744073709551615
                                                    switch fi88 {
                                                    case 9223372036854775807:
                                                        v38 = (^(v38))
                                                    default:
                                                        v38 = 124
                                                    }
                                                    var v89 float64 = -2.25
                                                    _ = fi88
                                                    _ = v89
                                                }
                                            }
                                            switch ("" + ("" + "a")) {
                                            case "a":
                                                switch fi85 {
                                                case 32768:
                                                    v33 -= 249
                                                }
                                            case "":
                                                var v90 int32 = 42
                                                _ = v90
                                            case "\n\t":
                                                if true {
                                                    for fi91 := 0; fi91 < 15; fi91++ {
                                                        v33 = uint64(v5)
                                                        v5 += 2147483647
                                                        if false {
                                                            v34 += 2
                                                        }
                                                        _ = fi91
                                                    }
                                                    v92 := GenS8{F0: int8(v5), F1: 697, F2: 807}
                                                    _ = v92
                                                } else {
                                                    var v93 float64 = (-2.25 + 1.5)
                                                    switch fi85 {
                                                    case 32768:
                                                        fmt.Println(v5, v32, v38)
                                                    default:
                                                        if (v1 < -253) {
                                                            v34 = v33
                                                            if ((42 >= -2147483648) && (0 >= 2)) {
                                                                if (((0 != 215) || (false || false)) && false) {
                                                                    var v94 float32 = ((1e10 * 3.14159) / 0.1)
                                                                    _ = v94
                                                                }
                                                                v5 = 2
                                                                for fi95 := 0; fi95 < 15; fi95++ {
                                                                    var v96 uint64 = uint64(v38)
                                                                    _ = fi95
                                                                    _ = v96
                                                                }
                                                            } else {
                                                                var v97 float32 = (3.14159 / float32((float64(3.14159) * (-2.25 - float64(0.1)))))
                                                                var v98 int8 = (v32 / ((v32) | 1))
                                                                if true {
                                                                    v5 -= 362
                                                                    for fi99 := 0; fi99 < 7; fi99++ {
                                                                        switch fi99 {
                                                                        case 0:
                                                                            v98 -= -128
                                                                        case 9223372036854775807:
                                                                            v100 := GenS7{F0: 32767, F1: ("foo bar" + (((("\n\t" + "\n\t") + ("hello" + "")) + ("a" + "")) + (("\n\t" + "foo bar") + ("a" + "hello"))))}
                                                                            var v101 string = "foo bar"
                                                                            v102 := GenS12{F0: false, F1: (fi99 % (((fi85 - (fi85 + 32768))) | 1)), F2: true}
                                                                            _ = v100
                                                                            _ = v101
                                                                            _ = v102
                                                                        default:
                                                                            var v103 int16 = -1
                                                                            _ = v103
                                                                        }
                                                                        if (-632 >= 2147483647) {
                                                                            fmt.Println(v38, v1, v38)
                                                                            var v104 int64 = (^(v1))
                                                                            _ = len(v3)
                                                                            _ = v104
                                                                        } else {
                                                                            var v105 bool = true
                                                                            _ = v105
                                                                        }
                                                                        v106 := GenS13{F0: (v5 ^ (v5 - v5)), F1: 42}
                                                                        _ = fi99
                                                                        _ = v106
                                                                    }
                                                                    var v107 uint32 = (v5 - (v5 * v5))
                                                                    _ = v107
                                                                }
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                            switch "hello" {
                                                            case "":
                                                                if (v32 < v32) {
                                                                    var v108 int64 = (v1 * (v1 + (v1 * (v1 - (v1 & (-(v1)))))))
                                                                    v108 = (v1 & (^(v108)))
                                                                    if (127 >= (fi85 - fi85)) {
                                                                        if true {
                                                                            for fi109 := 0; fi109 < 20; fi109++ {
                                                                                if true {
                                                                                    if (true || (42 == 1)) {
                                                                                        switch fi109 {
                                                                                        case 4294967295:
                                                                                            if true {
                                                                                                var v110 float32 = 1.5
                                                                                                v32 += (^(v32))
                                                                                                fmt.Println(v5, v38)
                                                                                                var v111 string = "foo bar"
                                                                                                if (1000 != -807) {
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    v112 := GenS5{F0: 0, F1: v5}
                                                                                                    var v113 int32 = 1
                                                                                                    var v114 int16 = 585
                                                                                                    var v115 uint16 = 0
                                                                                                    v116 := GenS9{F0: (v1 | v1), F1: false}
                                                                                                    _ = v112
                                                                                                    _ = v113
                                                                                                    _ = v114
                                                                                                    _ = v115
                                                                                                    _ = v116
                                                                                                } else {
                                                                                                    if (((v110 - v110) <= v110) || ((v93 / (((((0.0 - 0.1) + v93) - v93) / 1.0) * float64(v110))) < v93)) {
                                                                                                        v5 = v5
                                                                                                        if (!(((-(v32)) != (v32 - 1)) && true)) {
                                                                                                            switch (fi85 % ((fi85) | 1)) {
                                                                                                            case 0:
                                                                                                                fmt.Println(v1, v32, v33)
                                                                                                                v111 = v111
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v32, v108)
                                                                                                }
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                        case 2147483647:
                                                                                            var v117 uint = 0
                                                                                            _ = v117
                                                                                        }
                                                                                    }
                                                                                    _ = len(v4)
                                                                                    var v118 int8 = (v32 >> ((uint(v5)) & 63))
                                                                                    _ = v118
                                                                                } else {
                                                                                    var v119 int32 = 0
                                                                                    _ = v119
                                                                                }
                                                                                if (false && true) {
                                                                                    var v120 int32 = 42
                                                                                    v4[(fi109 >> ((255) & 63))] = (1.5 != v93)
                                                                                    switch fi109 {
                                                                                    case 903:
                                                                                        v93 = 3.14159
                                                                                    case 32768:
                                                                                        var v121 uint64 = 710
                                                                                        _ = v121
                                                                                    default:
                                                                                        var v122 bool = false
                                                                                        _ = v122
                                                                                    }
                                                                                    _ = v120
                                                                                }
                                                                                _ = fi109
                                                                            }
                                                                        } else {
                                                                            var v123 bool = true
                                                                            var v124 uint = 42
                                                                            _ = v123
                                                                            _ = v124
                                                                        }
                                                                        v108 -= v108
                                                                    }
                                                                    _ = v108
                                                                }
                                                            default:
                                                                for fi125 := 0; fi125 < 20; fi125++ {
                                                                    v33 = (v2 ^ v34)
                                                                    if (0 == 32767) {
                                                                        switch "a" {
                                                                        case "a":
                                                                            v38 = v38
                                                                        }
                                                                        fmt.Println(v93)
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 65535
                                                                        }
                                                                        if false {
                                                                            var v126 int = fi85
                                                                            if false {
                                                                                if true {
                                                                                    v5 -= (^(v5))
                                                                                } else {
                                                                                    fmt.Println(v33)
                                                                                    fmt.Println(v1, v93, v93)
                                                                                    switch ("foo bar" + "a") {
                                                                                    case "foo bar":
                                                                                        v34 -= 18446744073709551615
                                                                                    case "a":
                                                                                        if false {
                                                                                            switch ("" + ("" + "\n\t")) {
                                                                                            case "\n\t":
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        v34 = v2
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            v1 = v1
                                                                                                            var v127 int = int(v1)
                                                                                                            _ = v127
                                                                                                        } else {
                                                                                                            var v128 uint32 = uint32(v33)
                                                                                                            _ = v128
                                                                                                        }
                                                                                                    default:
                                                                                                        v93 = -2.25
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    var v129 float64 = float64((float32((v93 * v93)) - 3.14159))
                                                                                                    _ = v129
                                                                                                case "foo bar":
                                                                                                    var v130 int = (fi125 / (((v126 + (fi125 & (fi85 / ((-134) | 1))))) | 1))
                                                                                                    _ = v130
                                                                                                default:
                                                                                                    var v131 int = (v126 ^ fi125)
                                                                                                    _ = v131
                                                                                                }
                                                                                            case "foo bar":
                                                                                                if true {
                                                                                                    var v132 uint32 = (v5 & 42)
                                                                                                    v132 += (v5 % (((^(v5))) | 1))
                                                                                                    var v133 bool = (143 >= 140)
                                                                                                    _ = v132
                                                                                                    _ = v133
                                                                                                }
                                                                                            }
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                v134 := GenS14{F0: (v33 & v2), F1: (v5 & 0)}
                                                                                                _ = v134
                                                                                            case "\n\t":
                                                                                                if true {
                                                                                                    switch fi125 {
                                                                                                    case 0:
                                                                                                        switch ("a" + "foo bar") {
                                                                                                        case "foo bar":
                                                                                                            switch -484 {
                                                                                                            case -108:
                                                                                                                if (42 <= int8(fi125)) {
                                                                                                                    fmt.Println(v32, v38)
                                                                                                                    var v135 uint64 = v2
                                                                                                                    _ = v135
                                                                                                                }
                                                                                                                v136 := GenS15{F0: 1000}
                                                                                                                var v137 int64 = v1
                                                                                                                _ = v136
                                                                                                                _ = v137
                                                                                                            case 42:
                                                                                                                v138 := GenS16{F0: 1}
                                                                                                                _ = v138
                                                                                                            default:
                                                                                                                if (42 >= 2) {
                                                                                                                    switch "hello" {
                                                                                                                    case "foo bar":
                                                                                                                        var v139 uint64 = 224
                                                                                                                        _ = v139
                                                                                                                    default:
                                                                                                                        if ((("" + "foo bar") + "hello") != ("a" + "foo bar")) {
                                                                                                                            v140 := GenS6{F0: 18446744073709551615}
                                                                                                                            v93 = -2.25
                                                                                                                            fmt.Println(v34, v33, v126)
                                                                                                                            _ = v140
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v38 = v38
                                                                                                                    v141 := v4[1]
                                                                                                                    v142 := GenS17{F0: 0}
                                                                                                                    if (fi125 < (fi85 / (((v126 + (fi125 ^ (v126 & (fi125 ^ v126))))) | 1))) {
                                                                                                                        switch "hello" {
                                                                                                                        case "hello":
                                                                                                                            var v143 bool = (37 >= 0)
                                                                                                                            _ = v143
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        _ = len(v4)
                                                                                                                    }
                                                                                                                    _ = v141
                                                                                                                    _ = v142
                                                                                                                }
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            if (("\n\t" < "foo bar") && (105 > 1)) {
                                                                                                                delete(v4, 484)
                                                                                                                v93 = v93
                                                                                                                fmt.Println(v93, v34)
                                                                                                                if false {
                                                                                                                    fmt.Println(v93, v32)
                                                                                                                    var v144 uint64 = 0
                                                                                                                    v145 := GenS18{F0: "", F1: v38}
                                                                                                                    _ = v144
                                                                                                                    _ = v145
                                                                                                                }
                                                                                                                switch (v126 << ((65535) & 63)) {
                                                                                                                case -2147483648:
                                                                                                                    v2 -= v2
                                                                                                                case -9223372036854775808:
                                                                                                                    if (42 <= 2147483647) {
                                                                                                                        switch (^(fi85)) {
                                                                                                                        case -2147483648:
                                                                                                                            v1 -= v1
                                                                                                                        case 0:
                                                                                                                            v3 = append(v3, (v5 & v5))
                                                                                                                        case -251:
                                                                                                                            v38 = (^(v38))
                                                                                                                        }
                                                                                                                        fmt.Println(fi125, v34)
                                                                                                                        v146 := GenS19{F0: true, F1: v126}
                                                                                                                        if false {
                                                                                                                            v147 := GenS20{F0: (-(fi85)), F1: v32}
                                                                                                                            _ = v147
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            v148 := GenS13{F0: 42, F1: 1}
                                                                                                                            _ = v148
                                                                                                                        }
                                                                                                                        _ = v146
                                                                                                                    } else {
                                                                                                                        var v149 int8 = (v32 ^ 100)
                                                                                                                        _ = v149
                                                                                                                    }
                                                                                                                case 32767:
                                                                                                                    _ = len(v4)
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(v38)
                                                                                                        }
                                                                                                    case -620:
                                                                                                        v150 := GenS21{F0: 56}
                                                                                                        _ = v150
                                                                                                    case 1:
                                                                                                        var v151 int64 = -9223372036854775808
                                                                                                        if (0.1 > 3.14159) {
                                                                                                            var v152 uint16 = 255
                                                                                                            v1 = -451
                                                                                                            _ = v152
                                                                                                        }
                                                                                                        fmt.Println(v33)
                                                                                                        _ = v151
                                                                                                    }
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = 2147483648
                                                                                                    }
                                                                                                }
                                                                                            case "a":
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    if (!(-574 <= 0)) {
                                                                                                        v153 := GenS4{F0: 57, F1: 555, F2: (v5 % (((v5 >> ((0) & 63))) | 1))}
                                                                                                        if false {
                                                                                                            v5 = (v5 - (v5 * (^(v5))))
                                                                                                            v154 := GenS22{F0: v32, F1: (v5 + v5), F2: v5}
                                                                                                            _ = v154
                                                                                                        }
                                                                                                        if false {
                                                                                                            switch "" {
                                                                                                            case "a":
                                                                                                                v155 := GenS5{F0: -140, F1: v5}
                                                                                                                _ = v155
                                                                                                            case "\n\t":
                                                                                                                v1 = (v1 - (v1 >> ((997) & 63)))
                                                                                                            case "foo bar":
                                                                                                                if (-306 >= (v126 ^ fi85)) {
                                                                                                                    switch fi85 {
                                                                                                                    case 32768:
                                                                                                                        if len(v3) >= 2 {
                                                                                                                            v3 = v3[1:len(v3)]
                                                                                                                        }
                                                                                                                    case 511:
                                                                                                                        var v156 uint8 = (v38 | 2)
                                                                                                                        _ = v156
                                                                                                                    case -9223372036854775808:
                                                                                                                        if (v5 > 42) {
                                                                                                                            if len(v3) > 0 {
                                                                                                                                v3[0] = 4294967295
                                                                                                                            }
                                                                                                                            v157 := GenS10{F0: v2, F1: ""}
                                                                                                                            var v158 float32 = 0.1
                                                                                                                            _ = v157
                                                                                                                            _ = v158
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        fmt.Println(v32)
                                                                                                                        var v159 uint = 4294967295
                                                                                                                        _ = v159
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            fmt.Println(v38, v32)
                                                                                                        }
                                                                                                        _ = v153
                                                                                                    } else {
                                                                                                        v33 = v34
                                                                                                        var v160 float32 = 0.0
                                                                                                        _ = v160
                                                                                                    }
                                                                                                case "a":
                                                                                                    v93 = v93
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v126, v126, v32)
                                                                                fmt.Println(v2, v38, v1)
                                                                            } else {
                                                                                v161 := GenS14{F0: v2, F1: 65535}
                                                                                v162 := v4[-984]
                                                                                switch 50 {
                                                                                case 32767:
                                                                                    var v163 int32 = -877
                                                                                    _ = v163
                                                                                case 657:
                                                                                    var v164 int32 = -1
                                                                                    _ = v164
                                                                                default:
                                                                                    v165 := GenS22{F0: 100, F1: (^(v5)), F2: (v5 & (v5 - (v5 ^ v5)))}
                                                                                    _ = v165
                                                                                }
                                                                                _ = len(v4)
                                                                                _ = v161
                                                                                _ = v162
                                                                            }
                                                                            _ = v126
                                                                        } else {
                                                                            if ((v32 & (v32 & (-(v32)))) < (v32 - (-(v32)))) {
                                                                                fmt.Println(v93, v33)
                                                                            } else {
                                                                                if ((false && ((v1 * v1) <= (^(v1)))) || (true && (("a" + "hello") <= "a"))) {
                                                                                    fmt.Println(fi85)
                                                                                    var v166 string = "a"
                                                                                    _ = v166
                                                                                }
                                                                                if ((("a" < "foo bar") && false) && true) {
                                                                                    var v167 float64 = (v93 / v93)
                                                                                    var v168 uint = 647
                                                                                    if true {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = (v5 | 1)
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v2, v167)
                                                                                        switch (("a" + "foo bar") + "\n\t") {
                                                                                        case "":
                                                                                            switch -339 {
                                                                                            case -2147483648:
                                                                                                var v169 int = (^(fi125))
                                                                                                _ = v169
                                                                                            case 4294967295:
                                                                                                v93 += 1e10
                                                                                            default:
                                                                                                fmt.Println(fi85)
                                                                                            }
                                                                                        case "foo bar":
                                                                                            fmt.Println(v1, v167)
                                                                                        default:
                                                                                            v2 = v34
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v168, v33)
                                                                                    _ = v167
                                                                                    _ = v168
                                                                                } else {
                                                                                    var v170 uint = 0
                                                                                    if false {
                                                                                        var v171 int = 42
                                                                                        var v172 uint = (v170 | (v170 | v170))
                                                                                        switch fi125 {
                                                                                        case 801:
                                                                                            var v173 float32 = 0.0
                                                                                            _ = v173
                                                                                        default:
                                                                                            switch fi125 {
                                                                                            case 2147483647:
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = v5
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v171
                                                                                        _ = v172
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            if (1 > -373) {
                                                                                                v3 = append(v3, (^(v5)))
                                                                                                if false {
                                                                                                    fmt.Println(v38, v1)
                                                                                                } else {
                                                                                                    v3 = append(v3, 2)
                                                                                                    var v174 uint8 = 205
                                                                                                    _ = v174
                                                                                                }
                                                                                            } else {
                                                                                                switch (fi85 - fi85) {
                                                                                                case 2147483647:
                                                                                                    v175 := GenS19{F0: (!(1 == (v5 | (v5 / ((v5) | 1))))), F1: fi85}
                                                                                                    _ = v175
                                                                                                case 32767:
                                                                                                    var v176 uint = (v170 | (v170 >> (((v170 >> ((28) & 63))) & 63)))
                                                                                                    _ = v176
                                                                                                case 42:
                                                                                                    var v177 int = fi85
                                                                                                    _ = v177
                                                                                                default:
                                                                                                    var v178 uint8 = (v38 ^ (v38 ^ (v38 / ((v38) | 1))))
                                                                                                    _ = v178
                                                                                                }
                                                                                                if false {
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        fmt.Println(v33, v170, v38)
                                                                                                    case "\n\t":
                                                                                                        v4[fi125] = (!(0.0 != 1.5))
                                                                                                    }
                                                                                                    v179 := GenS23{F0: ""}
                                                                                                    switch (fi85 + int(v5)) {
                                                                                                    case 920:
                                                                                                        v34 -= 2
                                                                                                    default:
                                                                                                        if true {
                                                                                                            v32 -= (-(v32))
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v179
                                                                                                } else {
                                                                                                    if (3 != int16(v1)) {
                                                                                                        fmt.Println(v170)
                                                                                                    }
                                                                                                    v180 := GenS21{F0: 439}
                                                                                                    var v181 uint = 542
                                                                                                    _ = v180
                                                                                                    _ = v181
                                                                                                }
                                                                                            }
                                                                                        case "hello":
                                                                                            v93 = (v93 * float64((1.5 + -2.25)))
                                                                                        case "foo bar":
                                                                                            v182 := GenS24{F0: v5, F1: 18446744073709551615, F2: "a"}
                                                                                            _ = v182
                                                                                        }
                                                                                        var v183 uint8 = v38
                                                                                        _ = v183
                                                                                    }
                                                                                    _ = v170
                                                                                }
                                                                                switch -9223372036854775808 {
                                                                                case -9223372036854775808:
                                                                                    if true {
                                                                                        v3 = append(v3, (v5 - (v5 ^ 4294967295)))
                                                                                    }
                                                                                case 1:
                                                                                    var v184 uint = 18446744073709551615
                                                                                    if (false && ((-2.25 / (3.14159 * 3.14159)) >= (1.5 / ((1.5 - float32(1.5)) - (1.5 - 0.0))))) {
                                                                                        v185 := v4[833]
                                                                                        v93 -= (float64(-2.25) / v93)
                                                                                        _ = v185
                                                                                    }
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = (^(v5))
                                                                                    }
                                                                                    _ = v184
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v186 string = "hello"
                                                                    _ = fi125
                                                                    _ = v186
                                                                }
                                                            }
                                                        } else {
                                                            if true {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v187 := GenS21{F0: 1}
                                                                    _ = v187
                                                                case "a":
                                                                    v32 += (v32 << ((2) & 63))
                                                                case "\n\t":
                                                                    if (!false) {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (^(v5))
                                                                        }
                                                                        for fi188 := 0; fi188 < 18; fi188++ {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 2
                                                                            }
                                                                            fmt.Println(v1, v93)
                                                                            var v189 uint16 = 626
                                                                            _ = fi188
                                                                            _ = v189
                                                                        }
                                                                    } else {
                                                                        v5 = (v5 ^ (v5 * (v5 - (v5 >> ((622) & 63)))))
                                                                        fmt.Println(v34)
                                                                        var v190 uint = 65535
                                                                        _ = v190
                                                                    }
                                                                default:
                                                                    for fi191 := 0; fi191 < 11; fi191++ {
                                                                        fmt.Println(v1, fi85)
                                                                        var v192 uint8 = 10
                                                                        _ = fi191
                                                                        _ = v192
                                                                    }
                                                                }
                                                                if ("a" > "hello") {
                                                                    fmt.Println(fi85, v33, v2)
                                                                    if (!true) {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            switch fi85 {
                                                                            case -1:
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            case 689:
                                                                                _ = len(v4)
                                                                            }
                                                                        case "\n\t":
                                                                            v1 = v1
                                                                        case "":
                                                                            var v193 int8 = (-(v32))
                                                                            _ = v193
                                                                        default:
                                                                            var v194 uint64 = v2
                                                                            _ = v194
                                                                        }
                                                                        v32 = (v32 / (((v32 | -116)) | 1))
                                                                        for fi195 := 0; fi195 < 9; fi195++ {
                                                                            if (true || false) {
                                                                                if (v93 < 3.14159) {
                                                                                    v196 := GenS25{F0: true}
                                                                                    var v197 int = (fi85 << ((97) & 63))
                                                                                    _ = v196
                                                                                    _ = v197
                                                                                } else {
                                                                                    _ = len(v4)
                                                                                    if true {
                                                                                        v32 += (v32 >> ((2) & 63))
                                                                                    } else {
                                                                                        v198 := GenS26{F0: (fi85 | (fi195 << ((1) & 63))), F1: uint8(v33), F2: (v1 % (((v1 & (v1 / ((v1) | 1)))) | 1))}
                                                                                        _ = v198
                                                                                    }
                                                                                    fmt.Println(v32, v33)
                                                                                }
                                                                                if true {
                                                                                    if (false && (1e10 >= 3.14159)) {
                                                                                        var v199 uint64 = v34
                                                                                        v200 := GenS9{F0: -839, F1: (616 >= uint16(v38))}
                                                                                        if (2 < 567) {
                                                                                            var v201 int = 127
                                                                                            v202 := GenS1{F0: (v5 + (v5 / ((v5) | 1)))}
                                                                                            if true {
                                                                                                v203 := GenS1{F0: (^(v5))}
                                                                                                _ = v203
                                                                                            } else {
                                                                                                v204 := GenS24{F0: v5, F1: v199, F2: "foo bar"}
                                                                                                var v205 uint32 = v5
                                                                                                v206 := GenS2{F0: (v32 ^ -1), F1: 565, F2: uint(v5)}
                                                                                                var v207 int16 = 32767
                                                                                                _ = v204
                                                                                                _ = v205
                                                                                                _ = v206
                                                                                                _ = v207
                                                                                            }
                                                                                            _ = v201
                                                                                            _ = v202
                                                                                        } else {
                                                                                            v208 := GenS27{F0: 120}
                                                                                            v93 -= (v93 + v93)
                                                                                            _ = v208
                                                                                        }
                                                                                        _ = v199
                                                                                        _ = v200
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v209 uint64 = v33
                                                                                _ = v209
                                                                            }
                                                                            _ = fi195
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(v5, v93, v2)
                                                                if true {
                                                                    if true {
                                                                        v3 = append(v3, v5)
                                                                        v210 := GenS28{F0: 32767, F1: (v32 % (((v32 & v32)) | 1)), F2: 369}
                                                                        v211 := GenS29{F0: ("hello" + "a")}
                                                                        _ = v210
                                                                        _ = v211
                                                                    }
                                                                }
                                                            }
                                                            if (((v93 / float64(1.5)) - float64(float32((((v93 - 1.5) - float64(1.5)) - float64((3.14159 - 1.5)))))) > (v93 - v93)) {
                                                                _ = len(v4)
                                                            }
                                                        }
                                                    }
                                                    v1 = v1
                                                    _ = v93
                                                }
                                            }
                                            for fi212 := 0; fi212 < 9; fi212++ {
                                                var v213 int = (fi85 + (^(fi85)))
                                                if (801 >= 982) {
                                                    v1 = v1
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v214 := GenS27{F0: 127}
                                                        _ = v214
                                                    }
                                                }
                                                v5 = (v5 ^ (^(v5)))
                                                fmt.Println(v1)
                                                if true {
                                                    var v215 float64 = (-2.25 / -2.25)
                                                    _ = v215
                                                }
                                                _ = fi212
                                                _ = v213
                                            }
                                            _ = fi85
                                        }
                                        v1 += v1
                                    }
                                } else {
                                    for fi216 := 0; fi216 < 17; fi216++ {
                                        var v217 float32 = (1e10 / 1e10)
                                        if len(v3) > 0 {
                                            v3[0] = v5
                                        }
                                        fmt.Println(v32, fi216, fi216)
                                        _ = fi216
                                        _ = v217
                                    }
                                }
                            case "\n\t":
                                v32 = (-(v32))
                            }
                        }
                        var v218 uint32 = v5
                        v219 := GenS3{F0: 479, F1: 876}
                        _ = v218
                        _ = v219
                    }
                    switch 469 {
                    case 1:
                        v5 += v5
                    case 128:
                        switch -972 {
                        case 0:
                            var v220 float32 = float32(1e10)
                            _ = v220
                        case 888:
                            if true {
                                v221 := GenS30{F0: 65535}
                                _ = v221
                            } else {
                                var v222 int64 = v1
                                v223 := GenS13{F0: 42, F1: 255}
                                _ = v222
                                _ = v223
                            }
                        default:
                            v3 = append(v3, (v5 - v5))
                        }
                    case 32768:
                        var v224 int32 = -1
                        _ = v224
                    default:
                        var v225 uint64 = (v33 & 18446744073709551615)
                        _ = v225
                    }
                }
                v33 -= v34
                _ = v40
            }
            _ = v38
            _ = v39
        }
    } else {
        v5 = (v5 ^ (v5 * v5))
    }
    for fi226 := 0; fi226 < 1; fi226++ {
        v34 -= v33
        fmt.Println(fi226)
        switch (fi226 >> ((0) & 63)) {
        case -779:
            if false {
                if (0 >= (v1 | (v1 ^ 9223372036854775807))) {
                    switch "" {
                    case "a":
                        var v227 int16 = 1000
                        _ = v227
                    }
                    var v228 bool = true
                    _ = v228
                } else {
                    if false {
                        v5 -= v5
                        for fi229 := 0; fi229 < 12; fi229++ {
                            v4[(fi229 - (-(fi229)))] = false
                            var v230 int64 = -331
                            _ = fi229
                            _ = v230
                        }
                        fmt.Println(v5)
                    } else {
                        v33 = v33
                        for fi231 := 0; fi231 < 12; fi231++ {
                            var v232 int8 = (v32 ^ v32)
                            _ = fi231
                            _ = v232
                        }
                        for fi233 := 0; fi233 < 3; fi233++ {
                            _ = len(v4)
                            var v234 uint8 = 0
                            _ = fi233
                            _ = v234
                        }
                    }
                    if ((2147483648 != v5) || true) {
                        for fi235 := 0; fi235 < 15; fi235++ {
                            if true {
                                var v236 float32 = -2.25
                                if false {
                                    var v237 int8 = v32
                                    v5 = (v5 ^ 4294967295)
                                    v237 = (v237 % ((-1) | 1))
                                    _ = v237
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = v5
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v238 := GenS15{F0: 0}
                                    var v239 uint32 = (v5 & 2)
                                    _ = v238
                                    _ = v239
                                }
                                _ = v236
                            } else {
                                var v240 uint8 = 42
                                if (((float64((-2.25 - -2.25)) * 0.0) / 1e10) != ((0.0 * (0.0 - 0.0)) * 0.0)) {
                                    v241 := GenS9{F0: (v1 + 572), F1: ((v5 >= (v5 * (^(v5)))) || false)}
                                    _ = v241
                                }
                                switch fi235 {
                                case -610:
                                    v2 = (v34 << ((255) & 63))
                                    fmt.Println(fi226, fi235, v5)
                                    fmt.Println(v33)
                                case -658:
                                    var v242 int = (-(fi235))
                                    if (!(false || (v240 >= (v240 & (^(v240)))))) {
                                        if true {
                                            _ = len(v3)
                                        }
                                        var v243 uint16 = 42
                                        if (((0 <= (v34 * (v33 % (((v34 << ((747) & 63))) | 1)))) && true) && true) {
                                            if ((float32((1e10 / (float64(1.5) * (0.0 - 3.14159)))) + 0.1) != (1e10 * 3.14159)) {
                                                v244 := GenS14{F0: 572, F1: uint32(v240)}
                                                if ((v32 * 100) != int8(v1)) {
                                                    switch ("hello" + ("" + "\n\t")) {
                                                    case "hello":
                                                        switch (fi226 | fi226) {
                                                        case 127:
                                                            v245 := GenS31{F0: (3.14159 >= (1.5 / (((float32(0.1) / ((1.5 * 1.5) * 1e10)) + (float32(-2.25) * -2.25)) / 0.1))), F1: v2}
                                                            _ = v245
                                                        case 1:
                                                            v240 = (v240 << ((18446744073709551615) & 63))
                                                        }
                                                    default:
                                                        if true {
                                                            if false {
                                                                var v246 float32 = 0.0
                                                                var v247 string = "hello"
                                                                v248 := GenS32{F0: 42}
                                                                _ = v246
                                                                _ = v247
                                                                _ = v248
                                                            } else {
                                                                fmt.Println(v34)
                                                                var v249 bool = (true && true)
                                                                var v250 uint32 = 65535
                                                                _ = v249
                                                                _ = v250
                                                            }
                                                            var v251 string = ""
                                                            if (!true) {
                                                                var v252 int8 = int8(v240)
                                                                _ = v252
                                                            }
                                                            _ = v251
                                                        } else {
                                                            var v253 bool = (false || (true || true))
                                                            _ = v253
                                                        }
                                                    }
                                                    _ = len(v4)
                                                    v240 += (v240 % ((v240) | 1))
                                                } else {
                                                    v254 := GenS5{F0: -1, F1: (v5 >> ((42) & 63))}
                                                    if true {
                                                        v34 += 42
                                                    }
                                                    _ = v254
                                                }
                                                v255 := GenS9{F0: v1, F1: (1 != 0)}
                                                _ = v244
                                                _ = v255
                                            }
                                            var v256 int16 = 1000
                                            _ = v256
                                        } else {
                                            fmt.Println(v2, fi235, v240)
                                        }
                                        _ = v243
                                    } else {
                                        v1 = (v1 & 1)
                                        if false {
                                            v32 += -128
                                            if true {
                                                v257 := GenS14{F0: v33, F1: (v5 - v5)}
                                                _ = v257
                                            }
                                            var v258 uint16 = 65535
                                            _ = v258
                                        } else {
                                            v2 = 2
                                            v259 := GenS22{F0: (v32 % ((v32) | 1)), F1: (v5 * (v5 / ((v5) | 1))), F2: v5}
                                            var v260 uint = uint(v240)
                                            _ = v259
                                            _ = v260
                                        }
                                    }
                                    v261 := GenS5{F0: 0, F1: (v5 & (v5 - 528))}
                                    _ = v242
                                    _ = v261
                                default:
                                    v32 = v32
                                }
                                _ = v240
                            }
                            var v262 uint16 = 0
                            v263 := GenS12{F0: ((1 >= 1) || true), F1: fi235, F2: false}
                            _ = fi235
                            _ = v262
                            _ = v263
                        }
                        v264 := GenS33{F0: 2147483647, F1: fi226, F2: fi226}
                        v265 := GenS18{F0: ("\n\t" + ((("hello" + "a") + "a") + "foo bar")), F1: 128}
                        _ = v264
                        _ = v265
                    } else {
                        fmt.Println(fi226, v33)
                        for fi266 := 0; fi266 < 3; fi266++ {
                            if true {
                                if true {
                                    var v267 bool = true
                                    v3 = append(v3, (v5 * v5))
                                    v268 := GenS1{F0: v5}
                                    _ = v267
                                    _ = v268
                                } else {
                                    fmt.Println(v32)
                                    if (!(true && (1e10 > (1e10 / 0.1)))) {
                                        switch (fi266 + (^(fi266))) {
                                        case 2147483647:
                                            if (true || (false || (32767 != 32767))) {
                                                var v269 float64 = (float64((1e10 + 0.0)) + 1.5)
                                                fmt.Println(v2)
                                                _ = v269
                                            } else {
                                                if (!(false && (false && (!(!false))))) {
                                                    if (!false) {
                                                        switch 128 {
                                                        case 42:
                                                            v270 := GenS34{F0: 0, F1: 1}
                                                            _ = v270
                                                        default:
                                                            v34 = (v34 % ((v33) | 1))
                                                        }
                                                        fmt.Println(v34, v32, v33)
                                                    } else {
                                                        v1 = 751
                                                        if len(v3) > 0 {
                                                            v3[0] = 2
                                                        }
                                                    }
                                                    var v271 int = (fi226 << ((2) & 63))
                                                    var v272 float64 = -2.25
                                                    _ = v271
                                                    _ = v272
                                                } else {
                                                    if (true || (0.1 == -2.25)) {
                                                        v32 = -1
                                                        switch fi226 {
                                                        case 42:
                                                            if (("hello" + ("hello" + "a")) <= ((("hello" + "a") + "foo bar") + "")) {
                                                                var v273 float32 = 1.5
                                                                v2 = (v34 * (v2 & (v34 % ((v33) | 1))))
                                                                _ = v273
                                                            }
                                                        case -238:
                                                            var v274 int32 = 2147483647
                                                            _ = v274
                                                        default:
                                                            v275 := GenS35{F0: 0}
                                                            _ = v275
                                                        }
                                                        var v276 uint32 = 267
                                                        fmt.Println(fi226, v1)
                                                        _ = v276
                                                    } else {
                                                        fmt.Println(v1, v33, fi266)
                                                    }
                                                    switch 42 {
                                                    case 2147483647:
                                                        delete(v4, (fi266 * fi266))
                                                    default:
                                                        var v277 float64 = 3.14159
                                                        _ = v277
                                                    }
                                                    v33 = v34
                                                }
                                            }
                                        case 32767:
                                            v278 := GenS36{F0: 440, F1: 42}
                                            _ = v278
                                        }
                                    } else {
                                        v1 -= -1
                                        if (!(false || (!((v5 ^ v5) != (v5 % (((v5 | (^(v5)))) | 1)))))) {
                                            if true {
                                                if false {
                                                    switch 986 {
                                                    case 4294967295:
                                                        v33 = v34
                                                    case -937:
                                                        v279 := GenS17{F0: 828}
                                                        _ = v279
                                                    default:
                                                        v5 = v5
                                                    }
                                                    switch "a" {
                                                    case "\n\t":
                                                        if (-2147483648 != 2147483647) {
                                                            v5 -= (^(v5))
                                                        } else {
                                                            var v280 float64 = (((3.14159 / (((-2.25 + 0.0) * float64(3.14159)) - (0.1 / (3.14159 / 1.5)))) - (1e10 + (-2.25 * (1e10 + -2.25)))) - -2.25)
                                                            fmt.Println(v32, fi226)
                                                            switch 599 {
                                                            case 32767:
                                                                v281 := GenS37{F0: (v1 >> ((39) & 63))}
                                                                _ = v281
                                                            case 2147483647:
                                                                switch "" {
                                                                case "a":
                                                                    fmt.Println(v5, v34)
                                                                case "hello":
                                                                    if false {
                                                                        v5 -= (v5 & v5)
                                                                        fmt.Println(v1)
                                                                        v5 += v5
                                                                        if (-2147483648 >= -751) {
                                                                            var v282 uint8 = 172
                                                                            var v283 bool = (18446744073709551615 > 65535)
                                                                            v282 = 2
                                                                            fmt.Println(v282, v280)
                                                                            _ = v282
                                                                            _ = v283
                                                                        }
                                                                    } else {
                                                                        var v284 bool = (-2.25 > 0.1)
                                                                        _ = len(v4)
                                                                        switch (fi266 * (-(fi266))) {
                                                                        case 2147483647:
                                                                            var v285 string = "\n\t"
                                                                            _ = v285
                                                                        case 127:
                                                                            if (42 != 448) {
                                                                                fmt.Println(fi266, v5)
                                                                                v2 = 0
                                                                                if ((2 > 128) || v284) {
                                                                                    v286 := GenS38{F0: 22}
                                                                                    _ = v286
                                                                                } else {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    _ = len(v4)
                                                                                    if (691 > 925) {
                                                                                        if ((v33 - (v2 >> ((65535) & 63))) <= (v2 | 814)) {
                                                                                            v32 = (v32 % ((0) | 1))
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (v5 + (v5 % ((v5) | 1)))
                                                                                            }
                                                                                        }
                                                                                        if v284 {
                                                                                            var v287 uint8 = 173
                                                                                            v288 := GenS18{F0: "hello", F1: uint8(v5)}
                                                                                            _ = v287
                                                                                            _ = v288
                                                                                        } else {
                                                                                            fmt.Println(v1, v2)
                                                                                            fmt.Println(fi266)
                                                                                            if (!v284) {
                                                                                                v284 = v284
                                                                                            }
                                                                                            v1 = v1
                                                                                            var v289 uint32 = v5
                                                                                            _ = v289
                                                                                        }
                                                                                        var v290 string = ("\n\t" + (("a" + "foo bar") + "a"))
                                                                                        _ = v290
                                                                                    } else {
                                                                                        v33 += (v34 / (((v33 + v34)) | 1))
                                                                                        v291 := GenS26{F0: (fi226 / ((32767) | 1)), F1: 128, F2: -1000}
                                                                                        v292 := GenS10{F0: v34, F1: ""}
                                                                                        _ = v291
                                                                                        _ = v292
                                                                                    }
                                                                                }
                                                                            }
                                                                        case 764:
                                                                            switch 417 {
                                                                            case 32767:
                                                                                v284 = v284
                                                                            case 127:
                                                                                if (int8(v5) > (^(v32))) {
                                                                                    fmt.Println(fi226, fi226)
                                                                                    fmt.Println(v280, v32, fi266)
                                                                                    var v293 uint16 = 32767
                                                                                    if v284 {
                                                                                        var v294 int64 = v1
                                                                                        _ = v294
                                                                                    } else {
                                                                                        if v284 {
                                                                                            var v295 uint = 1
                                                                                            v284 = (0 != 145)
                                                                                            if v284 {
                                                                                                if true {
                                                                                                    var v296 int32 = -1
                                                                                                    _ = v296
                                                                                                }
                                                                                                v297 := v4[-2147483648]
                                                                                                var v298 uint64 = v34
                                                                                                var v299 bool = (!(!((-(v32)) <= v32)))
                                                                                                v5 += (v5 | v5)
                                                                                                _ = v297
                                                                                                _ = v298
                                                                                                _ = v299
                                                                                            }
                                                                                            _ = v295
                                                                                        }
                                                                                        fmt.Println(v280, v34, v293)
                                                                                    }
                                                                                    var v300 int = fi226
                                                                                    fmt.Println(v300, v293)
                                                                                    switch ((("\n\t" + ("foo bar" + "a")) + (("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) + "")) + ("hello" + ("a" + "hello"))) {
                                                                                    case "hello":
                                                                                        var v301 int = -1
                                                                                        _ = v301
                                                                                    case "\n\t":
                                                                                        v5 -= (v5 ^ 0)
                                                                                    }
                                                                                    _ = v293
                                                                                    _ = v300
                                                                                }
                                                                            case 372:
                                                                                v1 = v1
                                                                            }
                                                                        default:
                                                                            v280 = (v280 + float64(-2.25))
                                                                            delete(v4, (fi226 % ((846) | 1)))
                                                                            var v302 float64 = (v280 - v280)
                                                                            _ = v302
                                                                        }
                                                                        _ = v284
                                                                    }
                                                                default:
                                                                    v2 += v34
                                                                }
                                                            case -1:
                                                                v1 += -1
                                                            }
                                                            _ = v280
                                                        }
                                                    case "":
                                                        var v303 bool = (!true)
                                                        _ = v303
                                                    case "a":
                                                        v1 -= (v1 * 0)
                                                    }
                                                } else {
                                                    v304 := GenS39{F0: v5, F1: (^(v32)), F2: (^(v32))}
                                                    if true {
                                                        v3 = append(v3, (v5 * v5))
                                                        v305 := GenS40{F0: fi226, F1: 0}
                                                        switch (fi266 & (fi226 - (fi266 + fi226))) {
                                                        case 2147483647:
                                                            v5 -= v5
                                                        case -2147483648:
                                                            if false {
                                                                if false {
                                                                    var v306 uint64 = 4294967295
                                                                    delete(v4, (fi266 >> ((255) & 63)))
                                                                    v307 := GenS41{F0: (4294967295 == 0), F1: 42}
                                                                    _ = v306
                                                                    _ = v307
                                                                } else {
                                                                    v308 := GenS42{F0: 65535, F1: uint64(v5)}
                                                                    _ = v308
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v5 >> ((304) & 63))
                                                                }
                                                            }
                                                        }
                                                        _ = v305
                                                    } else {
                                                        fmt.Println(v2, v2, v32)
                                                    }
                                                    _ = v304
                                                }
                                                var v309 string = "\n\t"
                                                var v310 int8 = (^(v32))
                                                _ = v309
                                                _ = v310
                                            }
                                        } else {
                                            v311 := GenS43{F0: v5}
                                            var v312 uint16 = 32767
                                            v312 = v312
                                            _ = v311
                                            _ = v312
                                        }
                                    }
                                    if (!(true || false)) {
                                        var v313 uint64 = 42
                                        _ = v313
                                    } else {
                                        fmt.Println(v5, v33)
                                        switch "hello" {
                                        case "\n\t":
                                            v34 = (v2 / ((v34) | 1))
                                        default:
                                            v314 := GenS21{F0: 0}
                                            _ = v314
                                        }
                                        fmt.Println(fi226, v32, v2)
                                    }
                                }
                                fmt.Println(v32, v5, v33)
                            } else {
                                if (float64(-2.25) <= -2.25) {
                                    var v315 int16 = 32767
                                    fmt.Println(v1, v32)
                                    fmt.Println(v5, v1, v315)
                                    _ = v315
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = v5
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = (v5 & v5)
                                    }
                                    v316 := v4[800]
                                    _ = v316
                                }
                                var v317 uint8 = uint8(v34)
                                v1 += 1
                                _ = v317
                            }
                            var v318 float32 = (0.0 + (((0.0 * -2.25) - 3.14159) * 3.14159))
                            v33 = v34
                            _ = fi266
                            _ = v318
                        }
                        v319 := v4[fi226]
                        _ = v319
                    }
                    v32 = v32
                }
                for fi320 := 0; fi320 < 6; fi320++ {
                    var v321 uint32 = 2
                    v1 = 0
                    var v322 int32 = 42
                    _ = fi320
                    _ = v321
                    _ = v322
                }
            }
        }
        _ = fi226
    }
    var v323 uint32 = 0
    for fi324 := 0; fi324 < 6; fi324++ {
        for fi325 := 0; fi325 < 10; fi325++ {
            v33 += v33
            var v326 int32 = 0
            delete(v4, (fi324 % ((-402) | 1)))
            v327 := GenS43{F0: v323}
            if true {
                var v328 float64 = 1.5
                v329 := GenS44{F0: (v5 & v323)}
                _ = v328
                _ = v329
            }
            fmt.Println(fi325, v2, v34)
            _ = fi325
            _ = v326
            _ = v327
        }
        _ = fi324
    }
    switch 292 {
    case 127:
        v4[471] = false
    case 4294967295:
        v323 -= v323
    case 980:
        fmt.Println(v5)
    }
    v5 -= v323
    v330 := v4[1]
    v331 := GenS16{F0: 255}
    v332 := GenS42{F0: 65535, F1: (v34 - (v2 * v2))}
    if (float64((1.5 / 1.0)) < float64(-2.25)) {
        switch -94 {
        case 0:
            var v333 uint32 = v323
            _ = v333
        }
    } else {
        if (2 >= 2) {
            _ = len(v3)
        } else {
            v334 := GenS45{F0: 4294967295}
            _ = v334
        }
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v335 uint64 = (v33 | 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v323:", v323)
    fmt.Println("v33:", v33)
    fmt.Println("v330:", v330)
    fmt.Println("v335:", v335)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v323
    _ = v33
    _ = v330
    _ = v331
    _ = v332
    _ = v335
    _ = v34
    _ = v4
    _ = v5
}
