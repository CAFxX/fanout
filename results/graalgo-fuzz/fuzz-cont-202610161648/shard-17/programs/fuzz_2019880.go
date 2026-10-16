package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 string; F2 int }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 bool; F1 int16; F2 int }
type GenS10 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS11 struct { F0 bool; F1 uint64; F2 int32 }
type GenS12 struct { F0 bool; F1 uint16; F2 int16 }
type GenS13 struct { F0 int8; F1 int64; F2 int8 }
type GenS14 struct { F0 int64; F1 string; F2 string }
type GenS15 struct { F0 int64; F1 uint16 }
type GenS16 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS17 struct { F0 uint8; F1 int16 }

func main() {
    v1 := []uint32{0, 0, 65535, 2147483647}
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: -623, F1: (54 == 82)}
    v5 := GenS1{F0: 1, F1: true}
    if (1 > 2) {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = -1
            v2 = (v2 & -315)
            v8 := GenS1{F0: (^(v7)), F1: ((v2 * (v2 / ((1) | 1))) == -1)}
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: -32768, F1: (140 >= 18446744073709551615)}
        _ = v9
    } else {
        if (("" + "") >= ("\n\t" + ("foo bar" + ("a" + "\n\t")))) {
            v2 += v2
            fmt.Println(v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v1 = append(v1, 2)
        fmt.Println(v2)
    }
    v2 -= (v2 | (v2 << ((42) & 63)))
    switch -902 {
    case 523:
        var v10 int64 = v2
        _ = v10
    }
    var v11 uint16 = 255
    v12 := GenS2{F0: 499}
    fmt.Println(v11)
    v13 := GenS3{F0: v2, F1: "foo bar", F2: -950}
    for fi14 := 0; fi14 < 2; fi14++ {
        if (0.1 < -2.25) {
            fmt.Println(v11)
            v15 := GenS4{F0: -590, F1: 1}
            for fi16 := 0; fi16 < 12; fi16++ {
                v1 = append(v1, uint32(v11))
                v2 = v2
                var v17 int64 = 42
                switch ("a" + ("hello" + "a")) {
                case "foo bar":
                    fmt.Println(fi14, v2)
                case "a":
                    if (42 > 70) {
                        v11 -= (^(v11))
                        var v18 int8 = -128
                        _ = v18
                    }
                }
                v11 = 388
                var v19 bool = true
                v20 := GenS5{F0: 0, F1: 65535}
                v17 -= (-(v2))
                if v19 {
                    _ = len(v1)
                    switch int(v11) {
                    case 115:
                        var v21 uint64 = 1
                        _ = v21
                    case 9223372036854775807:
                        if false {
                            if (42 >= v11) {
                                if true {
                                    v19 = v19
                                } else {
                                    v19 = (-1 == int8(v11))
                                    if v19 {
                                        var v22 int32 = 0
                                        v23 := GenS4{F0: -739, F1: 0}
                                        _ = v22
                                        _ = v23
                                    } else {
                                        if (uint8(v11) >= 82) {
                                            var v24 string = ("" + (("foo bar" + "\n\t") + ((("" + ("\n\t" + ("a" + ("foo bar" + ("\n\t" + (("foo bar" + "hello") + ("" + ""))))))) + (("foo bar" + ("a" + (((("\n\t" + "") + ("" + "")) + "foo bar") + ""))) + "\n\t")) + ("\n\t" + "hello"))))
                                            delete(v3, 65535)
                                            _ = v24
                                        }
                                        var v25 bool = (!((^(v2)) != (v2 << ((429) & 63))))
                                        fmt.Println(fi16, fi14, fi16)
                                        _ = v25
                                    }
                                }
                                v26 := GenS6{F0: (v11 * (v11 - v11))}
                                _ = v26
                            } else {
                                v27 := GenS1{F0: 42, F1: true}
                                if (!(-1 >= -1)) {
                                    var v28 int64 = v17
                                    var v29 uint = 193
                                    var v30 uint8 = 118
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 int64 = v2
                                    if ("hello" > "foo bar") {
                                        v3[4294967295] = fi16
                                        v31 += v31
                                    }
                                    if v19 {
                                        var v32 uint32 = 2147483648
                                        _ = v32
                                    } else {
                                        _ = len(v1)
                                    }
                                    _ = v31
                                }
                                _ = v27
                            }
                        }
                    case 144:
                        switch (-(fi16)) {
                        case 4294967295:
                            v2 -= v2
                        }
                        var v33 float64 = 1e10
                        _ = v33
                    default:
                        if (42 < 128) {
                            var v34 float32 = 1e10
                            var v35 bool = false
                            _ = v34
                            _ = v35
                        }
                    }
                    v2 += -1
                } else {
                    switch 42 {
                    case 32768:
                        var v36 float64 = (3.14159 * ((0.0 - 1.5) - -2.25))
                        _ = v36
                    case -9223372036854775808:
                        if ((2147483647 < 892) && (v17 > 455)) {
                            var v37 float32 = (((-2.25 + (3.14159 + 0.1)) / 1e10) + 3.14159)
                            if false {
                                var v38 int16 = 0
                                if (((v17 << ((762) & 63)) >= (v2 & 42)) && v19) {
                                    v39 := GenS7{F0: 1, F1: 1}
                                    _ = v39
                                } else {
                                    switch "hello" {
                                    case "":
                                        v40 := GenS2{F0: -183}
                                        _ = v40
                                    }
                                    if v19 {
                                        if v19 {
                                            var v41 int16 = v38
                                            fmt.Println(v19, fi14, v41)
                                            _ = v41
                                        } else {
                                            fmt.Println(fi14)
                                            var v42 float64 = 1e10
                                            _ = v42
                                        }
                                        var v43 bool = (v19 || v19)
                                        _ = v43
                                    } else {
                                        var v44 uint64 = 445
                                        var v45 uint16 = v11
                                        if v19 {
                                            var v46 int = 2147483647
                                            _ = v46
                                        }
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                switch fi14 {
                                case 128:
                                    v47 := GenS8{F0: ("a" + ((("a" + (("\n\t" + "a") + "foo bar")) + "hello") + (("a" + ("hello" + "\n\t")) + ""))), F1: 2147483648}
                                    _ = v47
                                case 127:
                                    if v19 {
                                        switch fi16 {
                                        case 0:
                                            var v48 bool = ((v38 * (v38 - (v38 / ((532) | 1)))) >= v38)
                                            _ = v48
                                        case 42:
                                            if ((v19 && false) || (!false)) {
                                                v2 = (^(v17))
                                                v49 := GenS9{F0: false, F1: v38, F2: 42}
                                                _ = v49
                                            } else {
                                                switch fi14 {
                                                case 42:
                                                    v50 := GenS9{F0: (true && (!false)), F1: v38, F2: (^(fi16))}
                                                    _ = v50
                                                case 807:
                                                    if v19 {
                                                        _ = len(v3)
                                                    }
                                                case -877:
                                                    fmt.Println(v11, v37, fi16)
                                                }
                                            }
                                            v37 = v37
                                            v11 = v11
                                        case 32767:
                                            fmt.Println(v17)
                                        default:
                                            fmt.Println(v19, v37, v19)
                                        }
                                    }
                                default:
                                    fmt.Println(v17, fi14, v2)
                                }
                                _ = v38
                            }
                            _ = v37
                        }
                    case 730:
                        v1 = append(v1, 42)
                    default:
                        fmt.Println(v17, v2)
                    }
                    v51 := GenS4{F0: 570, F1: 36}
                    if (3.14159 > 0.0) {
                        v1 = append(v1, 0)
                        v11 = v11
                    } else {
                        var v52 int8 = 42
                        if (!(true && ((v11 * (v11 % ((v11) | 1))) > v11))) {
                            v19 = v19
                            switch "" {
                            case "foo bar":
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v52, v52)
                                case "":
                                    v1 = append(v1, 2)
                                }
                            }
                        }
                        _ = v52
                    }
                    _ = v51
                }
                _ = fi16
                _ = v17
                _ = v19
                _ = v20
            }
            _ = v15
        } else {
            switch 127 {
            case 42:
                if (3.14159 < 0.1) {
                    if true {
                        var v53 int64 = (v2 << ((358) & 63))
                        for fi54 := 0; fi54 < 20; fi54++ {
                            v55 := GenS3{F0: v2, F1: (("hello" + ("hello" + "a")) + "hello"), F2: 127}
                            _ = fi54
                            _ = v55
                        }
                        _ = v53
                    } else {
                        fmt.Println(v2, v2, fi14)
                    }
                    v2 += int64(fi14)
                    for fi56 := 0; fi56 < 9; fi56++ {
                        switch "foo bar" {
                        case "a":
                            var v57 float32 = (1.5 * 3.14159)
                            _ = v57
                        case "":
                            switch fi56 {
                            case 1:
                                if true {
                                    switch -751 {
                                    case 127:
                                        switch ("foo bar" + "\n\t") {
                                        case "\n\t":
                                            switch fi56 {
                                            case 4294967295:
                                                switch ("a" + (("" + "") + "hello")) {
                                                case "":
                                                    delete(v3, 369)
                                                case "a":
                                                    var v58 int64 = (v2 + (-(v2)))
                                                    _ = v58
                                                case "\n\t":
                                                    v2 = v2
                                                default:
                                                    if false {
                                                        var v59 uint16 = v11
                                                        if true {
                                                            if ((-110 <= -1) || false) {
                                                                var v60 bool = ("foo bar" >= "a")
                                                                _ = v60
                                                            } else {
                                                                if true {
                                                                    var v61 uint64 = 391
                                                                    fmt.Println(fi56, v61, v2)
                                                                    if (v11 < 42) {
                                                                        var v62 uint = 42
                                                                        if (!((fi56 > 551) || false)) {
                                                                            var v63 uint8 = 1
                                                                            var v64 uint16 = v59
                                                                            v61 = (v61 | v61)
                                                                            if true {
                                                                                switch fi14 {
                                                                                case 42:
                                                                                    v65 := GenS4{F0: -152, F1: 96}
                                                                                    _ = v65
                                                                                case -1:
                                                                                    _ = len(v1)
                                                                                case -584:
                                                                                    var v66 int8 = -1
                                                                                    _ = v66
                                                                                }
                                                                            } else {
                                                                                v67 := GenS10{F0: 2, F1: 4294967295, F2: 938}
                                                                                delete(v3, 4294967295)
                                                                                v2 += 0
                                                                                _ = v67
                                                                            }
                                                                            _ = v63
                                                                            _ = v64
                                                                        } else {
                                                                            v59 = v11
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 4294967295
                                                                            }
                                                                            var v68 bool = true
                                                                            _ = v68
                                                                        }
                                                                        var v69 int8 = 42
                                                                        var v70 uint32 = 273
                                                                        _ = v62
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                    _ = v61
                                                                } else {
                                                                    _ = len(v1)
                                                                    v71 := GenS11{F0: (!(66 == 2)), F1: 994, F2: 783}
                                                                    if (727 >= 2) {
                                                                        v11 = v11
                                                                        v2 = (-(v2))
                                                                    }
                                                                    _ = v71
                                                                }
                                                            }
                                                            v3[833] = (-(fi56))
                                                            v72 := GenS12{F0: true, F1: v11, F2: 1}
                                                            _ = v72
                                                        }
                                                        v2 = (v2 + (v2 - 42))
                                                        _ = v59
                                                    }
                                                }
                                            case -719:
                                                var v73 string = "hello"
                                                _ = v73
                                            case -539:
                                                var v74 uint32 = uint32(v11)
                                                _ = v74
                                            default:
                                                var v75 int32 = -260
                                                _ = v75
                                            }
                                            _ = len(v3)
                                            if false {
                                                _ = len(v3)
                                                if false {
                                                    v3[2147483648] = (-(fi14))
                                                    fmt.Println(v11)
                                                    if true {
                                                        v76 := GenS2{F0: 632}
                                                        v2 += (v2 * (v2 | v2))
                                                        _ = v76
                                                    } else {
                                                        v77 := GenS13{F0: int8(fi56), F1: v2, F2: 0}
                                                        switch "" {
                                                        case "":
                                                            if (!(true && true)) {
                                                                var v78 float64 = -2.25
                                                                v2 += -9223372036854775808
                                                                switch 0 {
                                                                case 271:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        if ((true || (-693 == 0)) || (-508 <= -2147483648)) {
                                                                            if (2 > 65535) {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 1
                                                                                }
                                                                                v1 = append(v1, 2147483647)
                                                                                if true {
                                                                                    v2 = 311
                                                                                    v78 = 3.14159
                                                                                }
                                                                            } else {
                                                                                v11 = 32767
                                                                                switch fi14 {
                                                                                case 42:
                                                                                    switch fi14 {
                                                                                    case 42:
                                                                                        var v79 int8 = 127
                                                                                        _ = v79
                                                                                    case 0:
                                                                                        v11 = v11
                                                                                    default:
                                                                                        v78 += (((1e10 - 0.0) * float64(((-2.25 + -2.25) / (1.5 + (-2.25 * 3.14159))))) / float64(float32(1e10)))
                                                                                    }
                                                                                case 97:
                                                                                    if true {
                                                                                        var v80 float32 = (1e10 + (0.0 * 0.0))
                                                                                        var v81 int16 = -1
                                                                                        var v82 int64 = 42
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    } else {
                                                                                        v11 = (v11 ^ (v11 << ((2) & 63)))
                                                                                        v78 = float64(1e10)
                                                                                    }
                                                                                case -2147483648:
                                                                                    v78 -= v78
                                                                                }
                                                                                fmt.Println(v11, v11)
                                                                                var v83 int32 = -2147483648
                                                                                _ = v83
                                                                            }
                                                                            v2 += v2
                                                                        }
                                                                    case "":
                                                                        var v84 uint8 = 128
                                                                        _ = v84
                                                                    }
                                                                case -74:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 2147483647
                                                                    }
                                                                case 32768:
                                                                    v85 := v3[4294967295]
                                                                    _ = v85
                                                                }
                                                                _ = v78
                                                            } else {
                                                                v11 = v11
                                                                v86 := GenS4{F0: 42, F1: 45}
                                                                var v87 int64 = 757
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                        case "\n\t":
                                                            var v88 uint = 18446744073709551615
                                                            _ = v88
                                                        default:
                                                            var v89 float32 = (1e10 - 0.1)
                                                            _ = v89
                                                        }
                                                        v2 += v2
                                                        if (!true) {
                                                            var v90 bool = true
                                                            switch 337 {
                                                            case -2147483648:
                                                                v91 := GenS14{F0: int64(v11), F1: (("hello" + "") + ("\n\t" + "")), F2: ("" + "")}
                                                                _ = v91
                                                            case 9223372036854775807:
                                                                var v92 int64 = v2
                                                                _ = v92
                                                            case 0:
                                                                var v93 uint64 = 2
                                                                _ = v93
                                                            default:
                                                                var v94 uint8 = 42
                                                                _ = v94
                                                            }
                                                            var v95 int8 = -44
                                                            v11 = (^(v11))
                                                            _ = v90
                                                            _ = v95
                                                        }
                                                        var v96 bool = (false || (2 < 221))
                                                        _ = v77
                                                        _ = v96
                                                    }
                                                    var v97 int = fi56
                                                    _ = v97
                                                } else {
                                                    var v98 uint16 = 32767
                                                    var v99 uint8 = 1
                                                    _ = len(v3)
                                                    _ = v98
                                                    _ = v99
                                                }
                                            }
                                        case "":
                                            var v100 string = ""
                                            _ = v100
                                        }
                                    default:
                                        switch (fi14 * (fi14 << ((2) & 63))) {
                                        case 166:
                                            v2 = -1
                                        default:
                                            v101 := GenS15{F0: v2, F1: v11}
                                            _ = v101
                                        }
                                    }
                                    var v102 string = ((("a" + "") + "") + "a")
                                    v103 := GenS12{F0: ((0 >= 349) && (128 <= 128)), F1: v11, F2: int16(fi56)}
                                    var v104 uint = uint(v11)
                                    v2 -= (v2 % (((-(v2))) | 1))
                                    v2 = 9223372036854775807
                                    fmt.Println(v104, v11, v11)
                                    _ = v102
                                    _ = v103
                                    _ = v104
                                }
                            case 2147483647:
                                v2 -= -687
                            case -1:
                                if false {
                                    var v105 uint32 = 287
                                    var v106 uint8 = 128
                                    v11 = v11
                                    _ = v105
                                    _ = v106
                                }
                            }
                        }
                        var v107 string = (("" + "foo bar") + "\n\t")
                        var v108 float64 = (0.0 / 1.5)
                        _ = fi56
                        _ = v107
                        _ = v108
                    }
                }
            case 9223372036854775807:
                v109 := GenS1{F0: -627, F1: false}
                _ = v109
            case 1:
                v3[1] = 2147483647
            }
            var v110 bool = false
            if (1e10 >= (0.1 + 0.0)) {
                v111 := GenS12{F0: (!v110), F1: (v11 << ((42) & 63)), F2: 111}
                var v112 int64 = (v2 >> ((587) & 63))
                _ = len(v1)
                _ = v111
                _ = v112
            } else {
                v113 := v3[1]
                _ = v113
            }
            _ = v110
        }
        _ = fi14
    }
    _ = len(v1)
    v2 = v2
    v114 := GenS3{F0: (^(v2)), F1: ("foo bar" + (("\n\t" + "\n\t") + ("hello" + ""))), F2: 42}
    if true {
        if ("a" >= "a") {
            v115 := GenS3{F0: v2, F1: "a", F2: 42}
            _ = v115
        }
        v116 := GenS16{F0: 1, F1: false, F2: (^(v11))}
        if (true && false) {
            var v117 uint16 = v11
            switch 775 {
            case -1:
                var v118 float64 = 0.1
                _ = v118
            case 127:
                fmt.Println(v2, v117)
            default:
                var v119 uint = 413
                _ = v119
            }
            _ = v117
        } else {
            var v120 float64 = float64(1e10)
            delete(v3, 34)
            _ = v120
        }
        _ = v116
    } else {
        var v121 int64 = 9223372036854775807
        for fi122 := 0; fi122 < 12; fi122++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                v123 := GenS17{F0: 127, F1: -900}
                switch ("foo bar" + "foo bar") {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "foo bar":
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                }
                v121 = v2
                _ = v123
            }
            _ = fi122
        }
        _ = v121
    }
    var v124 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v114
    _ = v12
    _ = v124
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
