package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint }
type GenS4 struct { F0 int32; F1 bool; F2 int32 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 int16; F1 uint8; F2 int8 }
type GenS11 struct { F0 bool; F1 uint32; F2 int32 }
type GenS12 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64; F1 int64 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint8; F1 uint16; F2 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 uint32; F1 int; F2 uint16 }
type GenS23 struct { F0 uint16; F1 uint }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 int32; F1 uint16 }
type GenS26 struct { F0 int16 }
type GenS27 struct { F0 int }
type GenS28 struct { F0 int32; F1 int16 }
type GenS29 struct { F0 int; F1 int8 }
type GenS30 struct { F0 uint16; F1 int32 }
type GenS31 struct { F0 uint; F1 int8 }
type GenS32 struct { F0 int16; F1 int64 }
type GenS33 struct { F0 int64 }
type GenS34 struct { F0 int64; F1 uint; F2 int }

func main() {
    var v1 float32 = 1e10
    v2 := []int{-366, 42, 4294967295}
    var v3 int8 = -25
    fmt.Println(v1)
    for fi4 := 0; fi4 < 16; fi4++ {
        v2 = append(v2, fi4)
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 int16 = -1
            switch (("\n\t" + "") + ("" + ("a" + (("hello" + "a") + ((("a" + "a") + "foo bar") + ""))))) {
            case "hello":
                v6 += (v6 & (v6 | v6))
            case "":
                var v7 bool = true
                _ = v7
            case "foo bar":
                v8 := GenS1{F0: 1}
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    if (2 <= 23) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (0 <= 101) {
            v9 := GenS2{F0: 255, F1: -1}
            v2 = append(v2, 128)
            _ = v9
        }
        fmt.Println(v3)
    } else {
        _ = len(v2)
        v3 -= 127
        if (true && false) {
            if (-9223372036854775808 <= 128) {
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS3{F0: 810, F1: 42, F2: 0}
                    _ = fi10
                    _ = v11
                }
                _ = len(v2)
                for fi12 := 0; fi12 < 11; fi12++ {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v1 = 3.14159
                        v14 := GenS4{F0: -1, F1: (201 > 2147483648), F2: -1}
                        switch 9223372036854775807 {
                        case 1:
                            if ((0.1 > (1e10 / 1.0)) || false) {
                                v3 = v3
                                v15 := GenS4{F0: 375, F1: (0 < 2), F2: 1}
                                v16 := GenS5{F0: 1, F1: 1}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v17 uint32 = 708
                                switch -2147483648 {
                                case -9223372036854775808:
                                    v18 := GenS6{F0: 0, F1: 2147483647}
                                    _ = v18
                                case 4294967295:
                                    v1 = v1
                                case 42:
                                    if len(v2) > 0 {
                                        v2[0] = fi13
                                    }
                                    v19 := GenS7{F0: 42, F1: v3}
                                    if true {
                                        v20 := GenS8{F0: -471, F1: v17}
                                        if false {
                                            if false {
                                                fmt.Println(fi13)
                                            } else {
                                                var v21 uint16 = 751
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v17 = v17
                                                _ = v21
                                            }
                                            fmt.Println(fi13, fi12)
                                        }
                                        switch "" {
                                        case "a":
                                            v3 += (v3 + (v3 | (v3 >> ((1) & 63))))
                                        case "":
                                            var v22 uint8 = 42
                                            v17 -= (^(v17))
                                            v23 := GenS9{F0: (^(v3)), F1: "a"}
                                            _ = v22
                                            _ = v23
                                        case "foo bar":
                                            v24 := GenS10{F0: 102, F1: 88, F2: -1}
                                            _ = v24
                                        }
                                        if (42 <= 846) {
                                            if false {
                                                if (1000 >= 1) {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        switch ("foo bar" + "") {
                                                        case "hello":
                                                            var v25 bool = false
                                                            if false {
                                                                if false {
                                                                    if (4294967295 < 18446744073709551615) {
                                                                        v26 := GenS5{F0: 91, F1: (fi13 << ((2) & 63))}
                                                                        v3 -= v3
                                                                        if (!v25) {
                                                                            var v27 int8 = 0
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v27
                                                                        } else {
                                                                            v28 := GenS6{F0: -787, F1: (fi12 - (-(fi12)))}
                                                                            _ = v28
                                                                        }
                                                                        _ = v26
                                                                    }
                                                                    v1 += v1
                                                                }
                                                                if len(v2) > 0 {
                                                                    v2[0] = -585
                                                                }
                                                                var v29 string = ("a" + "")
                                                                v30 := GenS11{F0: (!(v1 != v1)), F1: (v17 % (((v17 + (v17 - v17))) | 1)), F2: -1}
                                                                _ = v29
                                                                _ = v30
                                                            } else {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            if v25 {
                                                                var v31 uint32 = (^(v17))
                                                                _ = v31
                                                            } else {
                                                                v1 += v1
                                                            }
                                                            _ = v25
                                                        }
                                                    case "":
                                                        v17 += (^(v17))
                                                        var v32 float64 = 0.1
                                                        switch "a" {
                                                        case "":
                                                            v32 += (0.1 / v32)
                                                        case "\n\t":
                                                            var v33 int8 = v3
                                                            _ = v33
                                                        case "foo bar":
                                                            var v34 float32 = ((((-2.25 * (-2.25 * v1)) / v1) * float32(((((v32 + 1e10) / ((v32 - 0.0) * float64((1e10 + v1)))) * v32) * (-2.25 / float64(v1))))) / -2.25)
                                                            _ = v34
                                                        default:
                                                            if (!true) {
                                                                if true {
                                                                    var v35 uint32 = 42
                                                                    _ = v35
                                                                } else {
                                                                    var v36 int64 = 0
                                                                    _ = v36
                                                                }
                                                                var v37 string = ("" + "foo bar")
                                                                fmt.Println(v32, v3, v1)
                                                                _ = v37
                                                            }
                                                        }
                                                        _ = v32
                                                    }
                                                    if true {
                                                        v38 := GenS12{F0: 32767, F1: -349, F2: 1}
                                                        fmt.Println(v17)
                                                        v17 = v17
                                                        v39 := GenS13{F0: 0}
                                                        fmt.Println(v1)
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                } else {
                                                    v40 := GenS5{F0: 0, F1: (^(fi12))}
                                                    var v41 int16 = -11
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi13 >> ((1) & 63))
                                                    }
                                                    _ = v40
                                                    _ = v41
                                                }
                                                switch ("\n\t" + "\n\t") {
                                                case "\n\t":
                                                    var v42 uint8 = 116
                                                    _ = v42
                                                case "foo bar":
                                                    v2 = append(v2, -2147483648)
                                                case "a":
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v43 uint16 = 42
                                                        if (true && (-105 >= (v3 >> ((0) & 63)))) {
                                                            v17 -= 2147483647
                                                            v44 := GenS4{F0: -2147483648, F1: false, F2: -2147483648}
                                                            var v45 int8 = v3
                                                            if (("foo bar" + "\n\t") != ("\n\t" + ("a" + "hello"))) {
                                                                if (true && true) {
                                                                    var v46 int = (fi12 - (^(fi13)))
                                                                    v47 := GenS2{F0: 127, F1: 1}
                                                                    v48 := GenS14{F0: int(v17)}
                                                                    _ = v46
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                            }
                                                            switch fi12 {
                                                            case 4294967295:
                                                                if false {
                                                                    switch fi13 {
                                                                    case -94:
                                                                        fmt.Println(fi13)
                                                                    case -498:
                                                                        switch 912 {
                                                                        case 127:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi13
                                                                            }
                                                                        default:
                                                                            v49 := GenS4{F0: 2147483647, F1: true, F2: -2147483648}
                                                                            _ = v49
                                                                        }
                                                                    }
                                                                    var v50 int64 = -719
                                                                    fmt.Println(v17, fi12)
                                                                    _ = v50
                                                                }
                                                            case 32768:
                                                                if false {
                                                                    var v51 int8 = (^(v3))
                                                                    if (1000 < 485) {
                                                                        var v52 int64 = int64(v43)
                                                                        v43 = 584
                                                                        _ = v52
                                                                    } else {
                                                                        var v53 int64 = 1
                                                                        v1 = v1
                                                                        _ = v53
                                                                    }
                                                                    var v54 bool = true
                                                                    _ = v51
                                                                    _ = v54
                                                                }
                                                            }
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            switch (-(fi12)) {
                                                            case -28:
                                                                v55 := GenS15{F0: 251, F1: 765}
                                                                _ = v55
                                                            case 42:
                                                                v3 = v3
                                                            }
                                                            fmt.Println(v43)
                                                            _ = len(v2)
                                                        }
                                                        var v56 int8 = v3
                                                        v56 = (^(v56))
                                                        var v57 bool = false
                                                        if true {
                                                            var v58 float32 = 1.5
                                                            _ = v58
                                                        }
                                                        _ = v43
                                                        _ = v56
                                                        _ = v57
                                                    case "hello":
                                                        var v59 string = ("" + "\n\t")
                                                        _ = v59
                                                    }
                                                default:
                                                    fmt.Println(v1)
                                                }
                                                fmt.Println(v1, fi12, v3)
                                                _ = len(v2)
                                                var v60 uint = 18446744073709551615
                                                var v61 int16 = 42
                                                var v62 int64 = 328
                                                _ = v60
                                                _ = v61
                                                _ = v62
                                            } else {
                                                if false {
                                                    if (false && ((!((v3 + 127) < (v3 ^ v3))) || true)) {
                                                        var v63 uint64 = 31
                                                        _ = len(v2)
                                                        var v64 string = "foo bar"
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v17 -= (v17 - 1)
                                                        switch "a" {
                                                        case "\n\t":
                                                            fmt.Println(v1)
                                                        default:
                                                            var v65 int64 = int64(v3)
                                                            _ = v65
                                                        }
                                                        switch 9223372036854775807 {
                                                        case 0:
                                                            var v66 uint8 = 127
                                                            _ = v66
                                                        case 131:
                                                            var v67 int = (fi13 & int(v3))
                                                            _ = v67
                                                        case 2147483647:
                                                            v68 := GenS15{F0: 2, F1: 0}
                                                            _ = v68
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    if true {
                                                        fmt.Println(v17, fi13, v17)
                                                    } else {
                                                        v1 = 1.5
                                                        var v69 uint64 = uint64(v17)
                                                        switch (("hello" + "foo bar") + "") {
                                                        case "a":
                                                            if (!(!false)) {
                                                                if (!false) {
                                                                    v70 := GenS4{F0: 0, F1: (!(false || false)), F2: 42}
                                                                    _ = v70
                                                                }
                                                                fmt.Println(v17)
                                                                if false {
                                                                    if (true && true) {
                                                                        v1 = v1
                                                                        switch fi13 {
                                                                        case -528:
                                                                            v69 -= (^(v69))
                                                                        }
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = fi13
                                                                        }
                                                                        var v71 bool = false
                                                                        if v71 {
                                                                            v17 -= 626
                                                                            var v72 uint8 = 255
                                                                            fmt.Println(v69, v72)
                                                                            v72 = 53
                                                                            var v73 int8 = 127
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        _ = v71
                                                                    }
                                                                } else {
                                                                    v1 -= (v1 / v1)
                                                                    var v74 int = -469
                                                                    _ = v74
                                                                }
                                                            }
                                                        default:
                                                            var v75 uint32 = v17
                                                            _ = v75
                                                        }
                                                        _ = v69
                                                    }
                                                    _ = len(v2)
                                                    v76 := GenS16{F0: 210}
                                                    _ = v76
                                                } else {
                                                    _ = len(v2)
                                                    if true {
                                                        _ = len(v2)
                                                    }
                                                }
                                            }
                                        }
                                        fmt.Println(fi12, v17)
                                        _ = v20
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = v19
                                default:
                                    v3 += v3
                                    v77 := GenS5{F0: 776, F1: fi13}
                                    v78 := GenS17{F0: uint8(v17), F1: 2, F2: (!false)}
                                    v1 += v1
                                    if (!true) {
                                        var v79 bool = (!(127 <= 128))
                                        var v80 bool = v79
                                        _ = v79
                                        _ = v80
                                    } else {
                                        var v81 float32 = v1
                                        _ = v81
                                    }
                                    _ = v77
                                    _ = v78
                                }
                                switch -1 {
                                case -155:
                                    var v82 int16 = -430
                                    _ = v82
                                case 970:
                                    v83 := GenS18{F0: false}
                                    _ = v83
                                default:
                                    switch fi13 {
                                    case 4294967295:
                                        var v84 int16 = -582
                                        _ = v84
                                    case -263:
                                        switch (fi12 + -9223372036854775808) {
                                        case -23:
                                            var v85 uint16 = 42
                                            _ = v85
                                        default:
                                            fmt.Println(v3, fi12)
                                        }
                                        var v86 int16 = 1
                                        _ = v86
                                    default:
                                        v87 := GenS19{F0: 32767}
                                        _ = v87
                                    }
                                }
                                if ((v3 << ((22) & 63)) >= -9) {
                                    v3 = 127
                                    var v88 string = "\n\t"
                                    _ = v88
                                }
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                v89 := GenS3{F0: 1, F1: 42, F2: 2}
                                v90 := GenS15{F0: 18446744073709551615, F1: 0}
                                v1 = 3.14159
                                switch 312 {
                                case -711:
                                    v91 := GenS14{F0: fi12}
                                    _ = v91
                                case 4294967295:
                                    var v92 int32 = 43
                                    _ = v92
                                case 1:
                                    v3 = v3
                                default:
                                    var v93 uint32 = 42
                                    _ = v93
                                }
                                _ = len(v2)
                                _ = v89
                                _ = v90
                            }
                        case 4294967295:
                            if true {
                                if false {
                                    var v94 int16 = -1
                                    _ = v94
                                } else {
                                    v3 -= v3
                                }
                            } else {
                                if (-32768 != -1) {
                                    switch int(v3) {
                                    case 32767:
                                        v95 := GenS9{F0: v3, F1: "foo bar"}
                                        if true {
                                            v96 := GenS20{F0: 0}
                                            v97 := GenS21{F0: (false || (!(false || (9223372036854775807 >= 42))))}
                                            _ = v96
                                            _ = v97
                                        } else {
                                            if true {
                                                v1 -= 3.14159
                                                if ((((false || false) || (false && (!(9223372036854775807 <= -1)))) && ((200 != 25) && (!(false || (42 >= 2))))) || true) {
                                                    if (42 == 2) {
                                                        if false {
                                                            var v98 uint16 = 32767
                                                            v2 = append(v2, 0)
                                                            if (2 > 0) {
                                                                v99 := GenS19{F0: 32767}
                                                                var v100 string = "hello"
                                                                if false {
                                                                    fmt.Println(fi12, v98, v1)
                                                                } else {
                                                                    if (v1 > ((((0.0 / float32(-2.25)) / -2.25) / ((0.0 / -2.25) / ((v1 - v1) / (float32(3.14159) + 1.5)))) - 1e10)) {
                                                                        if false {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                        var v101 float64 = (1e10 / float64(v1))
                                                                        _ = v101
                                                                    }
                                                                    var v102 uint16 = (v98 * (v98 - (^(v98))))
                                                                    fmt.Println(v3, fi13)
                                                                    _ = v102
                                                                }
                                                                _ = v100
                                                                _ = v99
                                                            }
                                                            _ = v98
                                                        }
                                                        var v103 string = ""
                                                        fmt.Println(fi13, v1)
                                                        _ = v103
                                                    } else {
                                                        v104 := GenS22{F0: 598, F1: (fi12 * (-(fi13))), F2: 0}
                                                        _ = v104
                                                    }
                                                } else {
                                                    v3 -= 100
                                                    var v105 int64 = 539
                                                    _ = v105
                                                }
                                            }
                                            if (true || (0 < 104)) {
                                                var v106 string = ((("\n\t" + "") + "") + "\n\t")
                                                var v107 uint = 2
                                                _ = v106
                                                _ = v107
                                            } else {
                                                if true {
                                                    if false {
                                                        switch "foo bar" {
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = fi12
                                                            }
                                                        }
                                                        v2 = append(v2, 2147483647)
                                                        var v108 int = (fi13 % ((fi12) | 1))
                                                        v1 = (float32(1.5) + v1)
                                                        var v109 uint = 2
                                                        _ = v108
                                                        _ = v109
                                                    } else {
                                                        switch "foo bar" {
                                                        case "a":
                                                            v110 := GenS23{F0: 762, F1: 255}
                                                            _ = v110
                                                        case "\n\t":
                                                            v111 := GenS14{F0: fi13}
                                                            _ = v111
                                                        case "":
                                                            v112 := GenS24{F0: 0}
                                                            _ = v112
                                                        }
                                                    }
                                                    var v113 int64 = -9223372036854775808
                                                    if (18446744073709551615 > 146) {
                                                        v1 = float32(-2.25)
                                                        fmt.Println(v3)
                                                    } else {
                                                        v3 -= 0
                                                        if (2 != 212) {
                                                            v3 = -27
                                                            v3 -= v3
                                                            var v114 uint64 = 962
                                                            v115 := GenS18{F0: false}
                                                            _ = len(v2)
                                                            _ = v114
                                                            _ = v115
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    _ = v113
                                                }
                                                v1 = 3.14159
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v116 float32 = v1
                                                _ = v116
                                            case "a":
                                                var v117 uint32 = 697
                                                _ = v117
                                            }
                                        }
                                        _ = v95
                                    case 1:
                                        v3 -= 100
                                    case 127:
                                        if false {
                                            v2 = append(v2, (fi12 - (^(fi13))))
                                            _ = len(v2)
                                        } else {
                                            if (2 != 654) {
                                                var v118 int16 = -1
                                                _ = v118
                                            } else {
                                                v3 -= v3
                                                _ = len(v2)
                                                fmt.Println(v3, v3, fi13)
                                                v3 = v3
                                                v1 = (1e10 + (3.14159 / v1))
                                            }
                                            fmt.Println(v3)
                                            fmt.Println(v1)
                                        }
                                    default:
                                        fmt.Println(v1, fi13)
                                    }
                                    var v119 bool = (!false)
                                    _ = v119
                                }
                                var v120 float64 = ((-2.25 * (0.0 / 1e10)) - ((-2.25 * 3.14159) + (-2.25 + float64((1.5 / (v1 * (v1 + ((v1 * 0.0) / (3.14159 / ((0.0 - (v1 + (((((v1 + -2.25) / 1e10) + 0.1) - (v1 * v1)) / -2.25))) + 3.14159))))))))))
                                var v121 float64 = (3.14159 / float64((1.5 / 1.5)))
                                switch ((("a" + ("\n\t" + "a")) + "") + "a") {
                                case "foo bar":
                                    switch "foo bar" {
                                    case "foo bar":
                                        v122 := GenS25{F0: 2147483647, F1: 2}
                                        _ = v122
                                    case "\n\t":
                                        var v123 int = -161
                                        _ = v123
                                    }
                                case "hello":
                                    if true {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v124 uint64 = 866
                                            _ = v124
                                        case "a":
                                            v125 := GenS25{F0: 698, F1: 42}
                                            _ = v125
                                        }
                                    } else {
                                        if false {
                                            if false {
                                                switch (fi12 & (fi12 * -2147483648)) {
                                                case 2147483647:
                                                    var v126 float32 = 1.5
                                                    _ = v126
                                                case 809:
                                                    v127 := GenS7{F0: -799, F1: v3}
                                                    _ = v127
                                                }
                                                v121 -= (float64(v1) + v121)
                                            } else {
                                                v120 += (v120 + (v120 + v121))
                                                fmt.Println(v120, v120)
                                                var v128 uint32 = 489
                                                _ = v128
                                            }
                                            v2 = append(v2, 32767)
                                        } else {
                                            var v129 int32 = 42
                                            _ = v129
                                        }
                                        v121 += 3.14159
                                    }
                                }
                                if len(v2) > 0 {
                                    v2[0] = 128
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v2) > 0 {
                                        v2[0] = (fi13 | fi12)
                                    }
                                case "\n\t":
                                    switch "\n\t" {
                                    case "hello":
                                        var v130 bool = true
                                        _ = v130
                                    default:
                                        v1 = (float32(((v120 * v120) + v120)) * v1)
                                    }
                                case "a":
                                    if len(v2) > 0 {
                                        v2[0] = 9223372036854775807
                                    }
                                default:
                                    if false {
                                        if false {
                                            v120 = 0.1
                                            v2 = append(v2, fi12)
                                        } else {
                                            v131 := GenS6{F0: 0, F1: (fi13 ^ fi13)}
                                            _ = v131
                                        }
                                    }
                                }
                                fmt.Println(v3, v120, fi12)
                                var v132 string = ""
                                v2 = append(v2, fi12)
                                var v133 uint16 = 1
                                v132 = ("a" + "\n\t")
                                _ = v120
                                _ = v121
                                _ = v132
                                _ = v133
                            }
                        case -306:
                            var v134 uint = 96
                            _ = v134
                        }
                        fmt.Println(fi13, v3)
                        var v135 uint64 = 2
                        _ = fi13
                        _ = v135
                        _ = v14
                    }
                    if (191 >= 911) {
                        v1 = float32(1.5)
                    }
                    if true {
                        for fi136 := 0; fi136 < 11; fi136++ {
                            var v137 float64 = 1.5
                            _ = fi136
                            _ = v137
                        }
                        for fi138 := 0; fi138 < 15; fi138++ {
                            if ((1 >= 0) || false) {
                                v2 = append(v2, fi138)
                                if false {
                                    var v139 string = "a"
                                    _ = v139
                                } else {
                                    var v140 uint8 = 42
                                    _ = v140
                                }
                            }
                            var v141 int = fi12
                            v142 := GenS26{F0: 705}
                            _ = fi138
                            _ = v141
                            _ = v142
                        }
                    } else {
                        fmt.Println(fi12)
                        v2 = append(v2, (fi12 << ((4294967295) & 63)))
                        if false {
                            if true {
                                fmt.Println(v3, fi12, fi12)
                                if false {
                                    var v143 bool = ((true || (2 != 65535)) && (!((fi12 / ((fi12) | 1)) > (fi12 - (fi12 + (fi12 + fi12))))))
                                    _ = v143
                                } else {
                                    var v144 int8 = v3
                                    _ = v144
                                }
                            }
                        } else {
                            v145 := GenS13{F0: 2}
                            _ = v145
                        }
                    }
                    _ = fi12
                }
            }
            v2 = append(v2, 32767)
            var v146 bool = (((true || false) || (-2147483648 != -838)) && (40 <= 4294967295))
            _ = v146
        }
    }
    if len(v2) > 0 {
        v2[0] = -944
    }
    v3 = v3
    switch "hello" {
    case "hello":
        switch -997 {
        case 127:
            if (1 < 42) {
                for fi147 := 0; fi147 < 20; fi147++ {
                    var v148 int = int(v3)
                    var v149 int = fi147
                    _ = fi147
                    _ = v148
                    _ = v149
                }
                for fi150 := 0; fi150 < 8; fi150++ {
                    if (32767 <= 626) {
                        v151 := GenS27{F0: fi150}
                        _ = v151
                    } else {
                        for fi152 := 0; fi152 < 20; fi152++ {
                            var v153 int32 = 0
                            var v154 int = 1
                            _ = fi152
                            _ = v153
                            _ = v154
                        }
                        v1 += ((v1 / (v1 - (((((v1 + v1) + (v1 + -2.25)) * v1) + v1) + 0.1))) + v1)
                    }
                    _ = fi150
                }
                v1 = v1
            }
        case 9223372036854775807:
            if (-1 < -1) {
                for fi155 := 0; fi155 < 5; fi155++ {
                    fmt.Println(fi155, v3)
                    if false {
                        v1 -= float32(0.0)
                    }
                    _ = fi155
                }
                if (153 == 65535) {
                    _ = len(v2)
                    fmt.Println(v1)
                } else {
                    var v156 uint16 = 0
                    _ = v156
                }
            }
        case 128:
            v157 := GenS15{F0: 11, F1: -1}
            _ = v157
        default:
            var v158 int8 = 45
            _ = v158
        }
    case "foo bar":
        v159 := GenS28{F0: -1, F1: 32767}
        _ = v159
    }
    if true {
        var v160 int = 1
        if true {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v161 := GenS5{F0: 255, F1: v160}
                _ = v161
            } else {
                switch 32768 {
                case 2147483647:
                    if (((float64(1e10) / ((1.5 * -2.25) * 1e10)) < (float64(v1) + 0.1)) && false) {
                        for fi162 := 0; fi162 < 2; fi162++ {
                            var v163 int8 = (v3 / (((-(v3))) | 1))
                            v164 := GenS29{F0: (fi162 - v160), F1: v3}
                            var v165 int64 = -9223372036854775808
                            _ = fi162
                            _ = v163
                            _ = v164
                            _ = v165
                        }
                        v2 = append(v2, 4294967295)
                        v166 := GenS8{F0: -494, F1: 2}
                        _ = v166
                    }
                case 1:
                    for fi167 := 0; fi167 < 7; fi167++ {
                        v3 = 1
                        for fi168 := 0; fi168 < 14; fi168++ {
                            v169 := GenS30{F0: 32767, F1: 42}
                            v160 = (fi167 / ((0) | 1))
                            _ = fi168
                            _ = v169
                        }
                        var v170 uint32 = 0
                        _ = fi167
                        _ = v170
                    }
                case -272:
                    for fi171 := 0; fi171 < 9; fi171++ {
                        var v172 int64 = -989
                        _ = fi171
                        _ = v172
                    }
                default:
                    v1 -= -2.25
                }
                for fi173 := 0; fi173 < 6; fi173++ {
                    if (((-2.25 * 1e10) >= float64(float32(float64(0.1)))) || false) {
                        v174 := GenS18{F0: false}
                        _ = v174
                    } else {
                        v160 -= -1
                        v160 -= v160
                    }
                    v175 := GenS31{F0: 255, F1: 42}
                    _ = fi173
                    _ = v175
                }
            }
        } else {
            for fi176 := 0; fi176 < 2; fi176++ {
                if (!false) {
                    _ = len(v2)
                    if (v3 <= v3) {
                        switch "" {
                        case "hello":
                            _ = len(v2)
                        case "foo bar":
                            v160 -= (fi176 ^ -649)
                        }
                        if true {
                            v2 = append(v2, (fi176 ^ fi176))
                        } else {
                            v3 += (v3 * v3)
                            if true {
                                var v177 int32 = -2147483648
                                if (true || ((1 == -402) || (fi176 <= fi176))) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v2 = append(v2, fi176)
                                } else {
                                    v2 = append(v2, 137)
                                }
                                if true {
                                    switch "foo bar" {
                                    case "":
                                        switch -9223372036854775808 {
                                        case -827:
                                            v178 := GenS1{F0: 0}
                                            _ = v178
                                        case 1:
                                            switch (fi176 / ((2147483647) | 1)) {
                                            case 128:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                v179 := GenS32{F0: 1, F1: 9223372036854775807}
                                                _ = v179
                                            }
                                        case 9223372036854775807:
                                            if len(v2) > 0 {
                                                v2[0] = 345
                                            }
                                        }
                                    case "hello":
                                        for fi180 := 0; fi180 < 8; fi180++ {
                                            v181 := GenS29{F0: (fi180 - 361), F1: -61}
                                            _ = fi180
                                            _ = v181
                                        }
                                    default:
                                        v177 -= v177
                                    }
                                    var v182 int = -41
                                    v183 := GenS29{F0: 2147483647, F1: v3}
                                    _ = v182
                                    _ = v183
                                } else {
                                    if true {
                                        v184 := GenS33{F0: -1}
                                        if (true && (42 != (^(v177)))) {
                                            v177 = -2147483648
                                            for fi185 := 0; fi185 < 6; fi185++ {
                                                v186 := GenS34{F0: 9223372036854775807, F1: 65535, F2: (fi185 >> ((42) & 63))}
                                                if (true && false) {
                                                    switch (("hello" + "") + ("a" + ("foo bar" + (((("\n\t" + "") + "a") + "foo bar") + ("hello" + (("" + ("foo bar" + "\n\t")) + ("" + "hello"))))))) {
                                                    case "hello":
                                                        var v187 int16 = 32767
                                                        _ = v187
                                                    case "foo bar":
                                                        v177 += 210
                                                    default:
                                                        v3 = (v3 * (v3 >> ((255) & 63)))
                                                    }
                                                } else {
                                                    v188 := GenS13{F0: 128}
                                                    _ = v188
                                                }
                                                _ = fi185
                                                _ = v186
                                            }
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v184
                                    }
                                    fmt.Println(v160, v160)
                                    if (2147483647 != 4294967295) {
                                        if false {
                                            var v189 bool = true
                                            _ = v189
                                        }
                                        var v190 int32 = v177
                                        for fi191 := 0; fi191 < 11; fi191++ {
                                            fmt.Println(v190, v177, v3)
                                            v1 = float32(3.14159)
                                            _ = fi191
                                        }
                                        _ = v190
                                    } else {
                                        v160 = (fi176 * 0)
                                    }
                                }
                                _ = v177
                            } else {
                                if false {
                                    v3 = -1
                                    var v192 int16 = 1
                                    var v193 int32 = 2147483647
                                    _ = v192
                                    _ = v193
                                } else {
                                    var v194 uint64 = 397
                                    v3 = v3
                                    _ = v194
                                }
                            }
                        }
                        if true {
                            var v195 int32 = 598
                            _ = v195
                        } else {
                            var v196 uint64 = 18446744073709551615
                            _ = v196
                        }
                    }
                } else {
                    v1 -= v1
                    var v197 float64 = ((0.1 / (float64(1e10) * 1e10)) - 1e10)
                    var v198 int8 = 127
                    _ = v197
                    _ = v198
                }
                _ = fi176
            }
        }
        v1 -= ((v1 * (1.5 * v1)) - -2.25)
        _ = v160
    }
    v1 += (float32(((1.5 - 3.14159) / 1.0)) + 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
