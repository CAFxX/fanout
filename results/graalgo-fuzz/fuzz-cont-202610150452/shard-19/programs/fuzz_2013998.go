package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
type GenS4 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 uint; F1 int8; F2 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int }
type GenS13 struct { F0 bool; F1 string }
type GenS14 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS15 struct { F0 int64; F1 uint16; F2 uint }
type GenS16 struct { F0 bool; F1 int32 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 bool; F1 int32 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 uint32; F1 int; F2 uint32 }
type GenS22 struct { F0 string; F1 uint64; F2 uint8 }
type GenS23 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ""
    var v3 bool = true
    var v4 uint64 = 0
    var v5 uint64 = (v4 / ((v4) | 1))
    var v6 uint8 = 0
    var v7 int8 = 42
    if false {
        v8 := GenS1{F0: 777, F1: ""}
        var v9 uint = 883
        v6 = (v6 | (v6 ^ (v6 - 196)))
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: 474, F1: 32767, F2: 255}
        fmt.Println(v5, v3)
        _ = v10
    }
    switch ("" + v2) {
    case "hello":
        if v3 {
            _ = len(v1)
            switch ("a" + "") {
            case "foo bar":
                switch 32767 {
                case 128:
                    fmt.Println(v4, v7)
                case 9223372036854775807:
                    var v11 uint = 4294967295
                    _ = v11
                case -2147483648:
                    v12 := GenS3{F0: 1, F1: 1, F2: (v2 + (v2 + (("foo bar" + "") + v2)))}
                    _ = v12
                }
            case "hello":
                v1[v2] = 2147483648
            case "\n\t":
                v13 := []int{0, -9223372036854775808, 42, 2147483647}
                _ = v13
            default:
                fmt.Println(v2)
            }
            v4 = v5
        }
    case "\n\t":
        for fi14 := 0; fi14 < 8; fi14++ {
            if v3 {
                v2 = v2
                v15 := GenS1{F0: 2, F1: "hello"}
                _ = v15
            }
            delete(v1, v2)
            _ = fi14
        }
    default:
        v6 = (v6 - v6)
    }
    if (((42 != 0) || (false || false)) || v3) {
        v16 := GenS2{F0: 0, F1: -136, F2: 65535}
        if (v3 && (882 != 255)) {
            var v17 string = ""
            for fi18 := 0; fi18 < 4; fi18++ {
                if v3 {
                    if (v7 > (v7 >> ((1) & 63))) {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v20 := GenS4{F0: 0, F1: 0, F2: -188}
                            _ = fi19
                            _ = v20
                        }
                        var v21 uint64 = v4
                        fmt.Println(v21, v3)
                        _ = v21
                    } else {
                        if (!v3) {
                            v22 := GenS1{F0: uint32(v5), F1: "a"}
                            if (!(32767 == 1000)) {
                                if (!v3) {
                                    var v23 float64 = (0.0 - (3.14159 - (((3.14159 - 0.1) / 1.0) - 1.5)))
                                    fmt.Println(v7)
                                    for fi24 := 0; fi24 < 20; fi24++ {
                                        fmt.Println(v23, v6)
                                        if v3 {
                                            switch 0 {
                                            case -2147483648:
                                                var v25 uint16 = 42
                                                _ = v25
                                            }
                                            if (!v3) {
                                                v26 := GenS3{F0: 1, F1: 42, F2: (((v2 + ("a" + ((((v17 + (("" + v2) + "a")) + "hello") + "hello") + v2))) + ((v2 + ((v2 + "hello") + v2)) + v17)) + v17)}
                                                v27 := GenS2{F0: 951, F1: -1, F2: 2}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v23 = v23
                                            }
                                            var v28 int32 = 695
                                            _ = v28
                                        }
                                        _ = fi24
                                    }
                                    _ = v23
                                } else {
                                    v29 := GenS5{F0: 18446744073709551615, F1: 827}
                                    _ = v29
                                }
                                var v30 string = "a"
                                for fi31 := 0; fi31 < 10; fi31++ {
                                    v7 = v7
                                    _ = fi31
                                }
                                _ = v30
                            } else {
                                v32 := GenS6{F0: v5, F1: v7}
                                if (fi18 == (fi18 << ((255) & 63))) {
                                    if (42 < 42) {
                                        var v33 float32 = float32((float64((float32(3.14159) * (1.5 * (-2.25 + (3.14159 / 3.14159))))) * -2.25))
                                        switch fi18 {
                                        case 9223372036854775807:
                                            delete(v1, ((v2 + v17) + (v17 + v17)))
                                        case -9223372036854775808:
                                            v4 += v5
                                        case 128:
                                            fmt.Println(v2, fi18, v4)
                                        default:
                                            switch 0 {
                                            case 127:
                                                switch 9223372036854775807 {
                                                case -148:
                                                    if false {
                                                        if v3 {
                                                            for fi34 := 0; fi34 < 15; fi34++ {
                                                                v33 = ((1e10 / 1.0) + 0.1)
                                                                var v35 uint8 = 0
                                                                _ = fi34
                                                                _ = v35
                                                            }
                                                            v36 := []int{fi18, (fi18 | fi18)}
                                                            switch (((v2 + v17) + v2) + v2) {
                                                            case "":
                                                                var v37 uint8 = (v6 << ((984) & 63))
                                                                _ = v37
                                                            }
                                                            _ = v36
                                                        }
                                                    } else {
                                                        v17 = v2
                                                        fmt.Println(fi18, v17)
                                                        var v38 uint8 = 0
                                                        _ = v38
                                                    }
                                                case -1:
                                                    fmt.Println(v4)
                                                default:
                                                    for fi39 := 0; fi39 < 17; fi39++ {
                                                        switch fi18 {
                                                        case -1:
                                                            v17 = ("\n\t" + v17)
                                                        case 383:
                                                            v40 := v1[v2]
                                                            _ = v40
                                                        case 234:
                                                            v41 := GenS7{F0: v2}
                                                            _ = v41
                                                        default:
                                                            switch ("foo bar" + "foo bar") {
                                                            case "foo bar":
                                                                v33 -= ((v33 / float32((1e10 * 1.5))) / float32(((float64(float32(0.1)) - 3.14159) + 0.1)))
                                                            case "":
                                                                var v42 uint8 = 6
                                                                _ = v42
                                                            default:
                                                                fmt.Println(v4, v17, v7)
                                                            }
                                                            switch v2 {
                                                            case "\n\t":
                                                                fmt.Println(v4)
                                                            case "hello":
                                                                if (((float32((0.1 / 1.0)) + v33) / v33) < v33) {
                                                                    var v43 int32 = 0
                                                                    if ((!((0 == 513) || true)) || false) {
                                                                        fmt.Println(v3, fi18)
                                                                    } else {
                                                                        v44 := GenS8{F0: 4294967295, F1: 32767, F2: v6}
                                                                        v45 := GenS2{F0: 1, F1: 1, F2: 2}
                                                                        v43 -= (v43 ^ -803)
                                                                        _ = v44
                                                                        _ = v45
                                                                    }
                                                                    fmt.Println(fi39, v4)
                                                                    _ = v43
                                                                }
                                                            case "foo bar":
                                                                if (v6 >= (^(v6))) {
                                                                    var v46 int32 = 42
                                                                    _ = v46
                                                                } else {
                                                                    switch fi18 {
                                                                    case 128:
                                                                        v5 = v4
                                                                    case 32767:
                                                                        var v47 float32 = float32((3.14159 + 1e10))
                                                                        _ = v47
                                                                    }
                                                                }
                                                                fmt.Println(v4)
                                                            }
                                                            v6 -= 2
                                                        }
                                                        v48 := GenS7{F0: v17}
                                                        v6 -= (v6 % (((v6 / (((v6 & (^(v6)))) | 1))) | 1))
                                                        _ = fi39
                                                        _ = v48
                                                    }
                                                }
                                            case 32768:
                                                v49 := []int{(fi18 + fi18), 1, 42}
                                                _ = v49
                                            case -872:
                                                v50 := GenS9{F0: 42, F1: 255}
                                                _ = v50
                                            }
                                        }
                                        _ = v33
                                    } else {
                                        for fi51 := 0; fi51 < 7; fi51++ {
                                            v52 := []int{(fi18 * (fi51 * 4294967295)), fi51, fi18}
                                            v52 = append(v52, fi18)
                                            v53 := GenS10{F0: 2, F1: -128, F2: 4294967295}
                                            switch 127 {
                                            case 0:
                                                if false {
                                                    fmt.Println(v2, fi51)
                                                } else {
                                                    v7 = int8(v5)
                                                }
                                            case -9223372036854775808:
                                                var v54 int16 = 0
                                                _ = v54
                                            }
                                            var v55 uint = 886
                                            _ = fi51
                                            _ = v52
                                            _ = v53
                                            _ = v55
                                        }
                                        v3 = v3
                                    }
                                }
                                _ = v32
                            }
                            var v56 uint8 = (^(v6))
                            _ = v22
                            _ = v56
                        }
                        fmt.Println(v7, v5)
                        for fi57 := 0; fi57 < 14; fi57++ {
                            if ((true || v3) && (true && ((v6 ^ 237) != (v6 * (v6 * (^(v6))))))) {
                                v4 = v4
                                var v58 uint32 = 2147483648
                                _ = v58
                            }
                            switch "hello" {
                            case "a":
                                var v59 int64 = -9223372036854775808
                                _ = v59
                            case "\n\t":
                                v60 := []int{}
                                fmt.Println(v4)
                                _ = v60
                            }
                            v61 := GenS11{F0: 140}
                            _ = fi57
                            _ = v61
                        }
                    }
                    fmt.Println(v6)
                }
                var v62 uint8 = (v6 | 255)
                _ = fi18
                _ = v62
            }
            var v63 int8 = (v7 << ((967) & 63))
            _ = v17
            _ = v63
        } else {
            for fi64 := 0; fi64 < 19; fi64++ {
                v65 := []int{fi64, (fi64 & -9223372036854775808), fi64, fi64}
                for fi66 := 0; fi66 < 20; fi66++ {
                    v67 := v1["foo bar"]
                    switch (fi66 + (-(fi66))) {
                    case -581:
                        v67 = (v67 >> ((18446744073709551615) & 63))
                        if v3 {
                            var v68 uint8 = v6
                            _ = v68
                        } else {
                            if ((!false) || false) {
                                var v69 int64 = 0
                                _ = v69
                            } else {
                                var v70 int64 = 42
                                v7 = 1
                                _ = v70
                            }
                        }
                        if false {
                            fmt.Println(v7, v7, v6)
                            v71 := GenS12{F0: 803}
                            var v72 uint8 = v6
                            _ = len(v1)
                            var v73 uint8 = v72
                            v73 = v73
                            _ = v71
                            _ = v72
                            _ = v73
                        } else {
                            v74 := GenS8{F0: 2, F1: 32767, F2: 42}
                            v75 := v1[((v2 + v2) + v2)]
                            var v76 uint = uint(v6)
                            _ = v74
                            _ = v75
                            _ = v76
                        }
                    case 593:
                        v5 += v5
                    case 9223372036854775807:
                        fmt.Println(v67, v7, v67)
                        v3 = (v2 == v2)
                        if v3 {
                            switch "foo bar" {
                            case "\n\t":
                                var v77 uint = 65535
                                _ = v77
                            case "a":
                                v78 := GenS13{F0: false, F1: (("foo bar" + (v2 + (v2 + ("foo bar" + (v2 + v2))))) + (v2 + ""))}
                                _ = v78
                            default:
                                if len(v65) > 0 {
                                    v65[0] = (fi66 & (fi66 * fi66))
                                }
                            }
                            fmt.Println(v4)
                        }
                    default:
                        var v79 int32 = 1
                        v80 := GenS14{F0: -32768, F1: 143, F2: 1}
                        if v3 {
                            var v81 int = -9223372036854775808
                            switch "\n\t" {
                            case "a":
                                v6 += (^(v6))
                            }
                            v81 = (fi64 * fi66)
                            _ = v81
                        } else {
                            if (1.5 != 1.5) {
                                var v82 uint = 42
                                _ = v82
                            }
                            v83 := GenS3{F0: v7, F1: -9223372036854775808, F2: v2}
                            v67 += v67
                            var v84 string = (v2 + "foo bar")
                            _ = v83
                            _ = v84
                        }
                        _ = v79
                        _ = v80
                    }
                    var v85 bool = true
                    _ = fi66
                    _ = v67
                    _ = v85
                }
                _ = fi64
                _ = v65
            }
            var v86 int8 = -92
            _ = v86
        }
        v6 = uint8(v4)
        _ = v16
    } else {
        v87 := GenS2{F0: -9223372036854775808, F1: -32768, F2: 1}
        var v88 int8 = v7
        _ = v87
        _ = v88
    }
    if (v2 == "foo bar") {
        v7 = (v7 | v7)
        for fi89 := 0; fi89 < 14; fi89++ {
            if (!(!v3)) {
                if (v5 <= 42) {
                    if v3 {
                        v6 -= 179
                        v4 = uint64(v6)
                    }
                    if v3 {
                        v5 = (v5 >> ((649) & 63))
                        for fi90 := 0; fi90 < 5; fi90++ {
                            v1[""] = 241
                            _ = fi90
                        }
                        delete(v1, "\n\t")
                    }
                    if (!v3) {
                        for fi91 := 0; fi91 < 9; fi91++ {
                            v92 := []int{}
                            _ = fi91
                            _ = v92
                        }
                        var v93 int = fi89
                        v94 := GenS14{F0: 32767, F1: (v6 * 0), F2: 1}
                        _ = v93
                        _ = v94
                    }
                } else {
                    fmt.Println(v2, v5, v3)
                }
                var v95 float32 = 0.0
                _ = v95
            } else {
                v96 := GenS15{F0: 9223372036854775807, F1: 435, F2: 932}
                _ = v96
            }
            v97 := []int{fi89}
            if (659 >= 65535) {
                if v3 {
                    var v98 uint32 = 4294967295
                    if v3 {
                        v99 := GenS16{F0: v3, F1: 42}
                        _ = v99
                    } else {
                        v2 = "hello"
                        v100 := GenS17{F0: (^(v98))}
                        _ = v100
                    }
                    for fi101 := 0; fi101 < 4; fi101++ {
                        if v3 {
                            v98 += (^(v98))
                            var v102 uint = 4294967295
                            _ = v102
                        } else {
                            if (!((0.1 - (0.0 - ((3.14159 * 0.1) - float64(1e10)))) <= float64(0.1))) {
                                v103 := GenS18{F0: true}
                                fmt.Println(v3, fi89, v98)
                                var v104 float32 = float32(0.0)
                                var v105 int = fi89
                                _ = v103
                                _ = v104
                                _ = v105
                            } else {
                                v106 := GenS1{F0: v98, F1: "a"}
                                _ = v106
                            }
                            _ = len(v97)
                            v107 := GenS19{F0: (v3 && (v2 != (v2 + "foo bar"))), F1: 2147483647}
                            _ = v107
                        }
                        v7 = -40
                        var v108 uint8 = uint8(v98)
                        v109 := GenS20{F0: 127}
                        _ = fi101
                        _ = v108
                        _ = v109
                    }
                    _ = v98
                } else {
                    for fi110 := 0; fi110 < 12; fi110++ {
                        fmt.Println(v7)
                        var v111 int32 = 344
                        _ = fi110
                        _ = v111
                    }
                    var v112 float64 = 1e10
                    _ = v112
                }
            }
            _ = fi89
            _ = v97
        }
    } else {
        fmt.Println(v6, v5)
        v7 = 100
        switch -713 {
        case 32767:
            v113 := v1["\n\t"]
            _ = v113
        case -284:
            for fi114 := 0; fi114 < 12; fi114++ {
                var v115 string = ("\n\t" + (v2 + v2))
                switch (v115 + "foo bar") {
                case "foo bar":
                    for fi116 := 0; fi116 < 18; fi116++ {
                        v115 = v2
                        switch v115 {
                        case "hello":
                            v7 -= v7
                            var v117 uint16 = 32767
                            fmt.Println(fi116)
                            _ = v117
                        case "":
                            v3 = v3
                            v118 := []int{}
                            v5 = 1
                            _ = v118
                        case "foo bar":
                            v7 = (^(v7))
                        }
                        var v119 uint = 42
                        var v120 uint32 = uint32(v119)
                        _ = fi116
                        _ = v119
                        _ = v120
                    }
                }
                _ = fi114
                _ = v115
            }
        default:
            var v121 int16 = 398
            _ = v121
        }
    }
    v122 := []int{-9223372036854775808, -1, 127}
    v4 -= v4
    var v123 float32 = (0.1 * (0.1 / float32(3.14159)))
    v2 = (v2 + ("a" + "\n\t"))
    v124 := GenS8{F0: (v4 + (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))), F1: 0, F2: v6}
    if (18446744073709551615 < 65535) {
        v1[(((((v2 + v2) + (v2 + v2)) + "hello") + ((v2 + v2) + v2)) + (("\n\t" + v2) + ((v2 + "") + "\n\t")))] = 2147483648
        delete(v1, (("a" + "foo bar") + ""))
    } else {
        if len(v122) >= 2 {
            v122 = v122[1:len(v122)]
        }
        v2 = ("foo bar" + v2)
        delete(v1, "foo bar")
    }
    v125 := GenS21{F0: 65535, F1: 9223372036854775807, F2: 65535}
    v126 := GenS13{F0: v3, F1: (v2 + v2)}
    for fi127 := 0; fi127 < 15; fi127++ {
        if true {
            var v128 int16 = 1
            switch 127 {
            case -252:
                if ((293 != 42) && false) {
                    if len(v122) >= 2 {
                        v122 = v122[1:len(v122)]
                    }
                    _ = len(v122)
                }
            case 2147483647:
                if false {
                    switch fi127 {
                    case 1:
                        var v129 bool = (1e10 >= 0.1)
                        _ = v129
                    }
                    switch fi127 {
                    case 2147483647:
                        for fi130 := 0; fi130 < 6; fi130++ {
                            fmt.Println(v3, v4)
                            _ = fi130
                        }
                    case -614:
                        v6 = (v6 / ((140) | 1))
                    case 42:
                        switch fi127 {
                        case 2147483647:
                            var v131 uint = 305
                            _ = v131
                        default:
                            v128 -= 32767
                        }
                    default:
                        var v132 string = "\n\t"
                        _ = v132
                    }
                }
            case 440:
                if true {
                    var v133 int8 = 42
                    _ = len(v122)
                    if v3 {
                        var v134 bool = ("\n\t" >= v2)
                        v3 = (((v5 % ((v5) | 1)) <= (v4 ^ v5)) || v3)
                        _ = v134
                    }
                    _ = v133
                }
            }
            _ = v128
        } else {
            v135 := GenS22{F0: (("a" + v2) + v2), F1: (v4 - 42), F2: uint8(v5)}
            v123 += 0.1
            v136 := GenS22{F0: v2, F1: (v4 * (^(v4))), F2: (^(v6))}
            _ = v135
            _ = v136
        }
        _ = fi127
    }
    var v137 int8 = (v7 & v7)
    switch v2 {
    case "":
        for fi138 := 0; fi138 < 1; fi138++ {
            for fi139 := 0; fi139 < 7; fi139++ {
                v1["foo bar"] = 284
                _ = fi139
            }
            v123 += (v123 * v123)
            _ = fi138
        }
    case "hello":
        var v140 uint64 = (^(v4))
        _ = v140
    }
    var v141 uint8 = 2
    v142 := GenS23{F0: int16(v7), F1: v7}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v122:", len(v122), v122)
    fmt.Println("v123:", v123)
    fmt.Println("v137:", v137)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v122
    _ = v123
    _ = v124
    _ = v125
    _ = v126
    _ = v137
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
