package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 bool; F1 uint8; F2 int }
type GenS5 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int64; F1 bool; F2 int8 }
type GenS8 struct { F0 int32; F1 uint; F2 int8 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS11 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int16; F1 uint32; F2 int64 }
type GenS16 struct { F0 uint8; F1 int16; F2 int32 }
type GenS17 struct { F0 uint8; F1 bool }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int; F1 uint16; F2 int16 }
type GenS20 struct { F0 int32; F1 int32 }
func genFn0(p0 string, p1 bool) int {
    fmt.Println(p0)
    p0 = "hello"
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: 42, F1: -1, F2: false}
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint32, p1 uint64, p2 int16) uint16 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 3; fi6++ {
        p1 -= p1
        var v7 uint8 = 42
        if (p2 == (-(p2))) {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 float32 = (0.1 / (3.14159 / (1e10 / 0.1)))
                switch "" {
                case "foo bar":
                    var v10 uint64 = p1
                    _ = v10
                }
                p1 += (p1 >> ((178) & 63))
                v11 := GenS2{F0: 703}
                fmt.Println(v9)
                _ = fi8
                _ = v11
                _ = v9
            }
            var v12 int8 = 100
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    var v13 uint64 = (p1 & p1)
    v13 += (v13 | (p1 & (p1 ^ p1)))
    if false {
        if (1 < 128) {
            var v14 int32 = -1
            v15 := GenS3{F0: v14, F1: 65535}
            _ = v14
            _ = v15
        } else {
            p1 = (p1 - (^(p1)))
            v13 += 885
            if true {
                var v16 float64 = 1e10
                if false {
                    switch ("" + ("foo bar" + "foo bar")) {
                    case "hello":
                        switch genFn0("foo bar", (825 <= -937)) {
                        case 32768:
                            v17 := GenS2{F0: 2}
                            _ = v17
                        case -115:
                            v13 += v13
                        default:
                            v13 = p1
                        }
                    case "":
                        v18 := []int{128, -9223372036854775808, -2147483648}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 4; fi19++ {
                        if ("foo bar" == "foo bar") {
                            v20 := []int{(fi19 & (fi19 ^ genFn0("foo bar", (false && false)))), fi19}
                            _ = v20
                        } else {
                            fmt.Println(v13, fi19, p1)
                        }
                        _ = fi19
                    }
                    switch (("hello" + "\n\t") + "") {
                    case "\n\t":
                        if true {
                            switch "" {
                            case "":
                                switch genFn0("", true) {
                                case -680:
                                    v21 := GenS3{F0: 42, F1: 18446744073709551615}
                                    _ = v21
                                case -244:
                                    var v22 int = 0
                                    _ = v22
                                case 999:
                                    v23 := GenS4{F0: false, F1: 2, F2: 755}
                                    _ = v23
                                default:
                                    for fi24 := 0; fi24 < 11; fi24++ {
                                        var v25 uint8 = 42
                                        v26 := []int{(fi24 & (^(fi24))), (fi24 - genFn0("", true)), fi24, int(p2)}
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                }
                            case "\n\t":
                                var v27 uint = 0
                                _ = v27
                            default:
                                if (1 >= -2147483648) {
                                    for fi28 := 0; fi28 < 7; fi28++ {
                                        if false {
                                            if true {
                                                v29 := []int{(fi28 | (fi28 ^ int(p0))), genFn0("", false)}
                                                if true {
                                                    var v30 uint64 = p1
                                                    var v31 string = ""
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    if (true || (42 == -2147483648)) {
                                                        if (true || false) {
                                                            if true {
                                                                v13 += (^(v13))
                                                                for fi32 := 0; fi32 < 12; fi32++ {
                                                                    var v33 float64 = v16
                                                                    v29 = append(v29, genFn0("foo bar", false))
                                                                    if true {
                                                                        if false {
                                                                            v29 = append(v29, fi28)
                                                                            if (-715 < (fi32 + (fi28 >> ((4294967295) & 63)))) {
                                                                                fmt.Println(fi32, v33)
                                                                                var v34 uint64 = (^(v13))
                                                                                _ = v34
                                                                            }
                                                                        } else {
                                                                            if false {
                                                                                v35 := GenS5{F0: -9223372036854775808, F1: 272, F2: 2147483648}
                                                                                v29 = append(v29, (fi32 * (^(fi32))))
                                                                                v36 := GenS2{F0: uint(p0)}
                                                                                _ = v35
                                                                                _ = v36
                                                                            }
                                                                            if (fi28 < int(p2)) {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        v37 := GenS3{F0: 418, F1: 0}
                                                                                        v38 := GenS1{F0: p2, F1: -1, F2: false}
                                                                                        var v39 int = (fi32 / ((fi32) | 1))
                                                                                        _ = v37
                                                                                        _ = v38
                                                                                        _ = v39
                                                                                    }
                                                                                case "":
                                                                                    p0 -= 813
                                                                                case "a":
                                                                                    v16 += v33
                                                                                }
                                                                                fmt.Println(p2)
                                                                            } else {
                                                                                v13 = v13
                                                                            }
                                                                        }
                                                                        if false {
                                                                            p0 += (p0 >> ((42) & 63))
                                                                        } else {
                                                                            v40 := GenS6{F0: (^(p2)), F1: 18446744073709551615}
                                                                            _ = len(v29)
                                                                            if len(v29) >= 2 {
                                                                                v29 = v29[1:len(v29)]
                                                                            }
                                                                            _ = v40
                                                                        }
                                                                    } else {
                                                                        v41 := GenS3{F0: 875, F1: 775}
                                                                        _ = v41
                                                                    }
                                                                    _ = fi32
                                                                    _ = v33
                                                                }
                                                                var v42 int32 = -841
                                                                _ = v42
                                                            } else {
                                                                if len(v29) > 0 {
                                                                    v29[0] = (fi28 + fi28)
                                                                }
                                                            }
                                                            var v43 bool = true
                                                            fmt.Println(p2, p0, v13)
                                                            _ = v43
                                                        } else {
                                                            if ((fi28 | (fi28 / ((genFn0("\n\t", false)) | 1))) > -2147483648) {
                                                                var v44 int32 = -1
                                                                switch (fi28 ^ genFn0(("a" + ("a" + "foo bar")), (p2 < p2))) {
                                                                case 414:
                                                                    if false {
                                                                        if true {
                                                                            v45 := GenS7{F0: -1, F1: ((v44 << ((uint(p1)) & 63)) < v44), F2: 38}
                                                                            _ = v45
                                                                        }
                                                                        var v46 uint8 = 0
                                                                        _ = v46
                                                                    }
                                                                }
                                                                _ = v44
                                                            } else {
                                                                var v47 int64 = 42
                                                                for fi48 := 0; fi48 < 17; fi48++ {
                                                                    switch "" {
                                                                    case "":
                                                                        var v49 uint8 = 42
                                                                        _ = v49
                                                                    case "hello":
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            if len(v29) > 0 {
                                                                                v29[0] = (fi28 - genFn0("a", true))
                                                                            }
                                                                        case "":
                                                                            var v50 int16 = -832
                                                                            _ = v50
                                                                        default:
                                                                            v13 = 934
                                                                        }
                                                                    case "a":
                                                                        v51 := GenS8{F0: 42, F1: 42, F2: -18}
                                                                        _ = v51
                                                                    default:
                                                                        v29 = append(v29, (fi28 & (fi48 + (-(fi28)))))
                                                                    }
                                                                    var v52 uint16 = 288
                                                                    _ = fi48
                                                                    _ = v52
                                                                }
                                                                _ = v47
                                                            }
                                                        }
                                                        for fi53 := 0; fi53 < 10; fi53++ {
                                                            if len(v29) > 0 {
                                                                v29[0] = fi28
                                                            }
                                                            v54 := GenS1{F0: 32767, F1: -9223372036854775808, F2: (("\n\t" == "\n\t") || false)}
                                                            v55 := GenS9{F0: 65535, F1: p0}
                                                            _ = fi53
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    }
                                                    var v56 int32 = 1
                                                    fmt.Println(fi28, v13)
                                                    _ = v56
                                                }
                                                _ = v29
                                            } else {
                                                v16 += float64(3.14159)
                                                fmt.Println(p2)
                                            }
                                            for fi57 := 0; fi57 < 15; fi57++ {
                                                p0 = 65535
                                                _ = fi57
                                            }
                                            v58 := GenS10{F0: true, F1: 691, F2: 1}
                                            _ = v58
                                        } else {
                                            p1 = v13
                                            v16 = v16
                                            v59 := GenS11{F0: p2, F1: 42, F2: p0}
                                            _ = v59
                                        }
                                        var v60 uint16 = 9
                                        _ = fi28
                                        _ = v60
                                    }
                                    var v61 float32 = float32(0.1)
                                    v62 := []int{-786, -676}
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := GenS12{F0: (p2 | p2)}
                                    _ = v63
                                }
                            }
                        } else {
                            p2 += (^(p2))
                            for fi64 := 0; fi64 < 1; fi64++ {
                                v65 := GenS13{F0: "foo bar"}
                                _ = fi64
                                _ = v65
                            }
                            var v66 uint32 = 4294967295
                            _ = v66
                        }
                    case "a":
                        var v67 uint64 = (^(p1))
                        _ = v67
                    case "hello":
                        for fi68 := 0; fi68 < 17; fi68++ {
                            v69 := []int{(fi68 >> ((42) & 63)), genFn0("a", false), fi68}
                            switch ("hello" + "hello") {
                            case "hello":
                                switch fi68 {
                                case 975:
                                    for fi70 := 0; fi70 < 7; fi70++ {
                                        var v71 uint64 = (v13 << ((1) & 63))
                                        var v72 uint32 = (^(p0))
                                        if false {
                                            p1 = p1
                                            var v73 int32 = 745
                                            _ = v73
                                        }
                                        _ = fi70
                                        _ = v71
                                        _ = v72
                                    }
                                case 42:
                                    v16 = v16
                                default:
                                    v74 := GenS14{F0: (false || true)}
                                    _ = v74
                                }
                            }
                            _ = fi68
                            _ = v69
                        }
                    }
                } else {
                    switch "foo bar" {
                    case "":
                        v75 := GenS6{F0: -489, F1: 59}
                        _ = v75
                    case "\n\t":
                        v76 := GenS3{F0: 42, F1: 42}
                        _ = v76
                    }
                    var v77 int8 = 100
                    _ = v77
                }
                _ = v16
            }
        }
        fmt.Println(v13, p2)
        var v78 float32 = 0.0
        _ = v78
    } else {
        var v79 string = "foo bar"
        _ = v79
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 278
}

func main() {
    var v1 string = (("" + "hello") + ("foo bar" + "\n\t"))
    var v2 float32 = (0.0 + 1.5)
    var v3 string = "foo bar"
    var v80 uint = 1
    if ((!true) && (false || true)) {
        v80 = 2
        v3 = v3
        v3 = ""
    } else {
        for fi81 := 0; fi81 < 13; fi81++ {
            v82 := GenS15{F0: -1, F1: 585, F2: -1}
            if true {
                switch 9223372036854775807 {
                case -708:
                    v80 -= 18446744073709551615
                default:
                    var v83 uint64 = 279
                    _ = v83
                }
                fmt.Println(fi81)
            } else {
                v84 := GenS16{F0: 2, F1: 1, F2: 0}
                if true {
                    for fi85 := 0; fi85 < 7; fi85++ {
                        v3 = "hello"
                        _ = fi85
                    }
                }
                if true {
                    v86 := []int{fi81, fi81, (^(fi81)), (fi81 & fi81)}
                    if true {
                        fmt.Println(v80, v1, v2)
                    } else {
                        v80 -= 688
                        fmt.Println(v3, v1, v80)
                    }
                    v87 := GenS1{F0: -1, F1: 9223372036854775807, F2: (1e10 != v2)}
                    _ = v86
                    _ = v87
                }
                _ = v84
            }
            var v88 uint16 = 822
            _ = fi81
            _ = v82
            _ = v88
        }
        fmt.Println(v80, v2, v3)
        var v89 float32 = 3.14159
        _ = v89
    }
    for fi90 := 0; fi90 < 10; fi90++ {
        v91 := []int{fi90, (^(fi90))}
        if true {
            _ = len(v91)
            var v92 int64 = -460
            _ = v92
        } else {
            fmt.Println(v1)
            v2 += v2
            var v93 uint8 = 2
            _ = v93
        }
        _ = fi90
        _ = v91
    }
    for fi94 := 0; fi94 < 8; fi94++ {
        v2 = (-2.25 - (3.14159 / (float32(1.5) - float32(0.0))))
        v95 := GenS8{F0: 1, F1: v80, F2: 0}
        var v96 int = (fi94 << (((v80 << ((v80) & 63))) & 63))
        _ = fi94
        _ = v95
        _ = v96
    }
    v3 = "foo bar"
    v97 := []int{genFn0((v1 + ((v1 + "a") + "hello")), true)}
    fmt.Println(v80)
    if (false && (0 > 180)) {
        var v98 uint32 = 911
        v80 = v80
        if true {
            _ = len(v97)
            var v99 string = v1
            _ = v99
        }
        _ = v98
    } else {
        var v100 float32 = 1.5
        _ = v100
    }
    var v101 bool = true
    v102 := GenS12{F0: -617}
    v1 = v3
    for fi103 := 0; fi103 < 5; fi103++ {
        switch 32768 {
        case 9223372036854775807:
            var v104 int16 = -1
            _ = v104
        case -945:
            var v105 int64 = -791
            _ = v105
        case 2147483647:
            var v106 uint8 = 255
            _ = v106
        default:
            if v101 {
                if (837 <= 32767) {
                    v3 = "foo bar"
                    v80 -= (v80 >> (((v80 ^ v80)) & 63))
                    var v107 uint16 = 32767
                    _ = v107
                }
                if ((!true) && ((((0.1 + 1e10) + 3.14159) / 3.14159) >= (3.14159 / 0.1))) {
                    v108 := GenS16{F0: 255, F1: 1000, F2: 0}
                    v109 := GenS10{F0: true, F1: 994, F2: 148}
                    v110 := GenS17{F0: 42, F1: v101}
                    _ = v108
                    _ = v109
                    _ = v110
                } else {
                    for fi111 := 0; fi111 < 10; fi111++ {
                        switch (("\n\t" + (v1 + ("\n\t" + (("a" + "hello") + v1)))) + ((v1 + (v1 + ((("a" + v1) + "foo bar") + ((v1 + v1) + ((v3 + v3) + ""))))) + ((v1 + v3) + v1))) {
                        case "a":
                            v101 = true
                        case "foo bar":
                            if len(v97) >= 2 {
                                v97 = v97[1:len(v97)]
                            }
                        case "hello":
                            v1 = v3
                        default:
                            fmt.Println(v3)
                        }
                        _ = fi111
                    }
                    if v101 {
                        _ = len(v97)
                        v3 = (("a" + v3) + "\n\t")
                        v1 = (v1 + v3)
                    }
                    if (!(-742 <= 154)) {
                        if (v1 != v1) {
                            switch "hello" {
                            case "\n\t":
                                var v112 int32 = -2147483648
                                _ = v112
                            case "foo bar":
                                if v101 {
                                    v113 := GenS18{F0: 32767}
                                    if v101 {
                                        v114 := GenS19{F0: (fi103 & (fi103 ^ fi103)), F1: 0, F2: 42}
                                        _ = v114
                                    }
                                    v2 += 3.14159
                                    _ = v113
                                } else {
                                    if v101 {
                                        v115 := GenS9{F0: 153, F1: 42}
                                        _ = v115
                                    } else {
                                        var v116 int = (fi103 * fi103)
                                        v80 = v80
                                        var v117 bool = (!v101)
                                        _ = v116
                                        _ = v117
                                    }
                                }
                            }
                        }
                        for fi118 := 0; fi118 < 20; fi118++ {
                            var v119 string = v1
                            v120 := GenS20{F0: -322, F1: 42}
                            var v121 int32 = -1
                            _ = fi118
                            _ = v119
                            _ = v120
                            _ = v121
                        }
                    } else {
                        var v122 string = ((v3 + v3) + "hello")
                        _ = v122
                    }
                }
            } else {
                v2 = (v2 - v2)
                if len(v97) >= 2 {
                    v97 = v97[1:len(v97)]
                }
            }
        }
        _ = fi103
    }
    v123 := GenS5{F0: 1, F1: 255, F2: 697}
    if len(v97) > 0 {
        v97[0] = 808
    }
    v97 = append(v97, 4294967295)
    var v124 float32 = 0.0
    _ = len(v97)
    if (640 >= 914) {
        if v101 {
            var v125 uint64 = 18446744073709551615
            if (-365 >= 9223372036854775807) {
                switch 0 {
                case 1:
                    var v126 bool = (v125 <= (v125 + 1))
                    _ = v126
                case -832:
                    if (0.0 <= -2.25) {
                        fmt.Println(v80, v124, v125)
                        for fi127 := 0; fi127 < 3; fi127++ {
                            if v101 {
                                v3 = v3
                                v1 = (v1 + ("foo bar" + v3))
                                var v128 uint64 = (v125 >> (((^(v80))) & 63))
                                _ = v128
                            } else {
                                for fi129 := 0; fi129 < 15; fi129++ {
                                    if v101 {
                                        if (!v101) {
                                            v130 := GenS1{F0: 358, F1: 1, F2: (((!true) && ((v80 + (v80 % ((929) | 1))) != v80)) && (127 > 100))}
                                            v124 -= (1e10 * float32((float64((v2 / 1e10)) - 0.0)))
                                            v3 = v3
                                            _ = v130
                                        }
                                    }
                                    fmt.Println(v2)
                                    _ = fi129
                                }
                            }
                            _ = fi127
                        }
                        v131 := GenS3{F0: 42, F1: v80}
                        _ = v131
                    } else {
                        fmt.Println(v125)
                        for fi132 := 0; fi132 < 13; fi132++ {
                            for fi133 := 0; fi133 < 4; fi133++ {
                                v134 := GenS15{F0: -32768, F1: 2147483648, F2: -1}
                                _ = fi133
                                _ = v134
                            }
                            if v101 {
                                v3 = v3
                                v135 := GenS10{F0: v101, F1: 4294967295, F2: 65535}
                                v101 = (!((!(-1 < 42)) || (1.5 >= -2.25)))
                                _ = v135
                            }
                            _ = fi132
                        }
                    }
                default:
                    for fi136 := 0; fi136 < 16; fi136++ {
                        fmt.Println(v3, v1, fi136)
                        var v137 uint64 = 0
                        _ = fi136
                        _ = v137
                    }
                }
                if ((4294967295 == 1) || v101) {
                    var v138 int32 = -649
                    _ = v138
                } else {
                    v1 = "foo bar"
                }
                var v139 int8 = 1
                _ = v139
            }
            _ = v125
        } else {
            for fi140 := 0; fi140 < 3; fi140++ {
                for fi141 := 0; fi141 < 12; fi141++ {
                    var v142 int16 = -521
                    var v143 int64 = 9223372036854775807
                    _ = fi141
                    _ = v142
                    _ = v143
                }
                v144 := GenS5{F0: 50, F1: 262, F2: 91}
                _ = fi140
                _ = v144
            }
        }
        for fi145 := 0; fi145 < 8; fi145++ {
            v146 := GenS6{F0: 42, F1: 936}
            _ = fi145
            _ = v146
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v97:", len(v97), v97)
    _ = v1
    _ = v101
    _ = v102
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v80
    _ = v97
}
