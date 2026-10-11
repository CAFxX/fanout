package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int }
type GenS12 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 850
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    v6 := []string{("hello" + "\n\t"), "foo bar", "\n\t", ((("" + "foo bar") + ("a" + (("hello" + "\n\t") + ((((("foo bar" + "\n\t") + ("" + "a")) + "") + "") + "foo bar")))) + ("foo bar" + ("" + "")))}
    var v7 uint = 255
    v2 = (v2 >> ((2) & 63))
    v7 = v7
    fmt.Println(v2, v2)
    var v8 uint16 = 633
    v7 += v7
    for fi9 := 0; fi9 < 7; fi9++ {
        v7 += v7
        if (false && false) {
            if (false || true) {
                fmt.Println(v8, fi9, v7)
                var v10 uint16 = (v8 << (((v7 * v7)) & 63))
                _ = v10
            }
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int8 = 0
                var v13 uint16 = v8
                v8 -= (v8 / ((v13) | 1))
                _ = fi11
                _ = v12
                _ = v13
            }
            if (!false) {
                if true {
                    var v14 uint64 = 729
                    v14 = 134
                    var v15 int64 = 42
                    _ = v14
                    _ = v15
                }
                var v16 int16 = 32767
                _ = v16
            } else {
                v17 := GenS1{F0: uint(v8)}
                if (false || (false || (-2 != 1))) {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        if true {
                            if false {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            } else {
                                fmt.Println(fi18, fi18)
                                v19 := v3[-834]
                                _ = v19
                            }
                        } else {
                            if false {
                                if false {
                                    switch fi18 {
                                    case 127:
                                        var v20 bool = false
                                        _ = v20
                                    case 2147483647:
                                        var v21 int64 = 42
                                        _ = v21
                                    case 32768:
                                        v2 += v2
                                    default:
                                        var v22 int16 = 32767
                                        _ = v22
                                    }
                                    var v23 uint8 = 255
                                    _ = v23
                                } else {
                                    if ((!true) || (!false)) {
                                        v24 := GenS2{F0: 0, F1: -32768}
                                        delete(v5, 246)
                                        _ = v24
                                    } else {
                                        v25 := GenS2{F0: 2147483648, F1: -32768}
                                        var v26 uint16 = v8
                                        v27 := GenS3{F0: -128, F1: 32767}
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                v28 := GenS2{F0: 65535, F1: 0}
                                _ = v28
                            } else {
                                if (42 == -397) {
                                    if false {
                                        if len(v6) > 0 {
                                            v6[0] = ""
                                        }
                                        var v29 uint = (^(v7))
                                        var v30 bool = (95 > 128)
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v2 = v2
                                        if (true || false) {
                                            v7 -= v7
                                            switch "" {
                                            case "foo bar":
                                                if true {
                                                    var v31 int16 = 32767
                                                    v31 += -350
                                                    var v32 bool = false
                                                    _ = v31
                                                    _ = v32
                                                }
                                            }
                                        }
                                        delete(v4, "a")
                                        v33 := v3[(fi18 | 4294967295)]
                                        v6 = append(v6, "a")
                                        _ = v33
                                    }
                                } else {
                                    var v34 int16 = -1
                                    var v35 int = (fi9 + fi9)
                                    var v36 string = "a"
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                v2 = (v2 << ((0) & 63))
                                v2 -= v2
                                _ = len(v6)
                                v7 -= (v7 | (v7 >> ((v7) & 63)))
                            }
                            var v37 int64 = int64(v7)
                            _ = v37
                        }
                        fmt.Println(fi9)
                        switch ("a" + "hello") {
                        case "\n\t":
                            v8 = v8
                        case "":
                            v3[fi18] = fi18
                        }
                        var v38 int64 = 9223372036854775807
                        switch (fi18 % (((fi9 + -1)) | 1)) {
                        case 42:
                            v39 := GenS4{F0: -2147483648}
                            _ = v39
                        }
                        var v40 bool = true
                        _ = fi18
                        _ = v38
                        _ = v40
                    }
                    switch "\n\t" {
                    case "":
                        if true {
                            var v41 bool = false
                            switch "foo bar" {
                            case "hello":
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v4["hello"] = 2
                                    if (42 != 42) {
                                        if v41 {
                                            var v43 int32 = -1
                                            _ = v43
                                        }
                                        if false {
                                            if v41 {
                                                var v44 int8 = 100
                                                _ = v44
                                            } else {
                                                v45 := GenS4{F0: (^(fi42))}
                                                v46 := GenS3{F0: 125, F1: -504}
                                                _ = v45
                                                _ = v46
                                            }
                                            var v47 uint32 = 1
                                            _ = v47
                                        } else {
                                            if (((true || true) || false) && v41) {
                                                v8 -= (^(v8))
                                                v48 := GenS3{F0: 1, F1: 1000}
                                                v2 += (v2 & (v2 + (v2 ^ 4294967295)))
                                                if len(v6) > 0 {
                                                    v6[0] = ("" + "a")
                                                }
                                                _ = v48
                                            } else {
                                                var v49 int64 = 1
                                                var v50 uint = 0
                                                var v51 int64 = (v49 % ((v49) | 1))
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            }
                                            if (-2147483648 != 1) {
                                                v52 := GenS5{F0: 1000}
                                                var v53 int8 = 42
                                                var v54 string = "foo bar"
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            } else {
                                                var v55 string = "\n\t"
                                                _ = v55
                                            }
                                            v56 := GenS2{F0: 187, F1: 990}
                                            _ = v56
                                        }
                                        var v57 float64 = (0.0 + (0.1 + -2.25))
                                        _ = v57
                                    } else {
                                        fmt.Println(fi42)
                                        if (-71 <= -74) {
                                            if len(v6) > 0 {
                                                v6[0] = ("a" + "a")
                                            }
                                            v58 := GenS1{F0: v7}
                                            v8 += v8
                                            _ = v58
                                        }
                                    }
                                    var v59 int8 = -1
                                    _ = fi42
                                    _ = v59
                                }
                            case "\n\t":
                                v41 = false
                            }
                            _ = v41
                        } else {
                            v60 := GenS2{F0: 65535, F1: 1}
                            if true {
                                for fi61 := 0; fi61 < 19; fi61++ {
                                    if false {
                                        fmt.Println(fi9, v2)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                    v7 += v7
                                    var v62 int64 = -609
                                    _ = fi61
                                    _ = v62
                                }
                                v63 := GenS4{F0: fi9}
                                _ = v63
                            } else {
                                fmt.Println(fi9)
                                var v64 bool = (-940 == 943)
                                _ = v64
                            }
                            var v65 bool = (v8 == 2)
                            _ = v60
                            _ = v65
                        }
                    }
                    for fi66 := 0; fi66 < 12; fi66++ {
                        if false {
                            var v67 int8 = 18
                            switch 846 {
                            case -777:
                                if (0.1 < (0.1 / (float32(0.1) + -2.25))) {
                                    if (float64(0.0) >= 0.0) {
                                        fmt.Println(fi9)
                                    } else {
                                        v7 = (v7 & v7)
                                    }
                                }
                            case -2147483648:
                                fmt.Println(fi9, fi66, v67)
                            default:
                                var v68 string = ((((("\n\t" + ((("\n\t" + "\n\t") + (("" + "\n\t") + "")) + (("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))) + ("foo bar" + "a")))) + "hello") + ("a" + "\n\t")) + "hello") + "")
                                _ = v68
                            }
                            v2 += (^(v2))
                            _ = v67
                        } else {
                            v2 = v2
                            if true {
                                var v69 float32 = 0.1
                                var v70 uint16 = v8
                                fmt.Println(v70, fi66, v8)
                                _ = v69
                                _ = v70
                            } else {
                                if false {
                                    if (v2 >= v2) {
                                        var v71 int32 = 0
                                        v72 := GenS6{F0: (v71 % (((v71 % ((v71) | 1))) | 1))}
                                        _ = len(v3)
                                        fmt.Println(v71)
                                        _ = v71
                                        _ = v72
                                    } else {
                                        if len(v6) > 0 {
                                            v6[0] = "\n\t"
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v73 bool = true
                                        fmt.Println(v2, v73, fi66)
                                        _ = v73
                                    }
                                    v8 = (^(v8))
                                    v74 := GenS7{F0: -1, F1: true}
                                    _ = v74
                                } else {
                                    switch "" {
                                    case "a":
                                        _ = len(v6)
                                    case "\n\t":
                                        var v75 uint8 = 0
                                        _ = v75
                                    }
                                }
                            }
                        }
                        v8 -= 255
                        _ = fi66
                    }
                } else {
                    if false {
                        var v76 float32 = 0.1
                        _ = len(v6)
                        var v77 uint8 = 127
                        _ = v76
                        _ = v77
                    } else {
                        v78 := GenS3{F0: -1, F1: 203}
                        for fi79 := 0; fi79 < 17; fi79++ {
                            if false {
                                v80 := GenS5{F0: -32768}
                                if (!false) {
                                    fmt.Println(fi9, v7, v2)
                                    if (2147483647 < 633) {
                                        delete(v4, "foo bar")
                                    }
                                    if true {
                                        if false {
                                            v81 := v3[(-(fi79))]
                                            _ = v81
                                        } else {
                                            v82 := GenS7{F0: 1, F1: (42 <= 1)}
                                            fmt.Println(fi9, fi79)
                                            _ = v82
                                        }
                                        switch (fi9 * fi79) {
                                        case 127:
                                            if true {
                                                v83 := GenS8{F0: 196}
                                                _ = v83
                                            } else {
                                                v8 = v8
                                                v2 = (v2 & (v2 | 0))
                                            }
                                        case -398:
                                            v84 := GenS5{F0: 1}
                                            _ = v84
                                        case -16:
                                            var v85 int32 = -1
                                            _ = v85
                                        }
                                        v86 := GenS7{F0: 0, F1: (true && false)}
                                        _ = v86
                                    }
                                    fmt.Println(fi9)
                                    var v87 uint = 65535
                                    _ = v87
                                }
                                var v88 float64 = -2.25
                                _ = v80
                                _ = v88
                            }
                            v89 := GenS6{F0: 1}
                            _ = fi79
                            _ = v89
                        }
                        switch fi9 {
                        case -9223372036854775808:
                            var v90 string = "foo bar"
                            _ = v90
                        case 2147483647:
                            v91 := GenS2{F0: 4294967295, F1: -1}
                            _ = v91
                        case 1:
                            v92 := GenS5{F0: 0}
                            _ = v92
                        default:
                            if false {
                                for fi93 := 0; fi93 < 10; fi93++ {
                                    switch ("" + "") {
                                    case "a":
                                        v7 = 4294967295
                                    }
                                    _ = fi93
                                }
                            }
                        }
                        _ = v78
                    }
                    for fi94 := 0; fi94 < 20; fi94++ {
                        v95 := GenS5{F0: 935}
                        _ = fi94
                        _ = v95
                    }
                    for fi96 := 0; fi96 < 5; fi96++ {
                        var v97 float64 = 0.1
                        v2 += v2
                        _ = fi96
                        _ = v97
                    }
                }
                v3[(-(fi9))] = (fi9 | fi9)
                _ = v17
            }
        }
        _ = fi9
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!((v2 ^ (v2 + (v2 - (v2 - (^(v2)))))) != v2)) {
        var v98 int = 9223372036854775807
        for fi99 := 0; fi99 < 12; fi99++ {
            for fi100 := 0; fi100 < 7; fi100++ {
                v101 := GenS9{F0: true, F1: (^(v8))}
                v102 := GenS5{F0: 0}
                if false {
                    v103 := GenS10{F0: "a"}
                    v104 := GenS5{F0: -1}
                    _ = v103
                    _ = v104
                } else {
                    var v105 uint8 = 111
                    var v106 int8 = -128
                    var v107 uint8 = 58
                    switch "\n\t" {
                    case "foo bar":
                        v108 := GenS1{F0: 1}
                        _ = v108
                    case "\n\t":
                        delete(v3, -1)
                    }
                    _ = v105
                    _ = v106
                    _ = v107
                }
                v7 = (v7 | (v7 + v7))
                _ = fi100
                _ = v101
                _ = v102
            }
            v8 -= 31
            if (v7 >= 65535) {
                v98 = fi99
                for fi109 := 0; fi109 < 5; fi109++ {
                    v110 := GenS3{F0: -47, F1: 1}
                    _ = fi109
                    _ = v110
                }
                v8 = (v8 << ((18446744073709551615) & 63))
            } else {
                var v111 int32 = -316
                for fi112 := 0; fi112 < 7; fi112++ {
                    v6 = append(v6, ("" + ""))
                    _ = fi112
                }
                var v113 int32 = (v111 | v111)
                _ = v111
                _ = v113
            }
            _ = fi99
        }
        _ = v98
    } else {
        var v114 float32 = (0.1 / -2.25)
        _ = v114
    }
    if true {
        _ = len(v1)
        fmt.Println(v7)
    }
    switch -1 {
    case 32768:
        v8 = (v8 << ((v7) & 63))
    case 128:
        for fi115 := 0; fi115 < 9; fi115++ {
            var v116 int = (fi115 >> (((v7 & (v7 - (v7 | v7)))) & 63))
            if (!true) {
                v116 = (fi115 | v116)
            } else {
                delete(v3, (v116 & -923))
            }
            if (!true) {
                if true {
                    switch "hello" {
                    case "hello":
                        v7 += (^(v7))
                    default:
                        if false {
                            if ("" != "a") {
                                if (0.1 > (1e10 - 0.1)) {
                                    v117 := GenS11{F0: (v116 & 32768)}
                                    if ((v8 + v8) != (^(v8))) {
                                        for fi118 := 0; fi118 < 6; fi118++ {
                                            var v119 string = "\n\t"
                                            var v120 uint16 = uint16(v7)
                                            _ = fi118
                                            _ = v119
                                            _ = v120
                                        }
                                        v121 := GenS3{F0: -1, F1: 0}
                                        _ = v121
                                    }
                                    _ = v117
                                } else {
                                    v122 := GenS3{F0: 1, F1: -32768}
                                    for fi123 := 0; fi123 < 9; fi123++ {
                                        v116 = 9223372036854775807
                                        _ = fi123
                                    }
                                    v2 = v2
                                    _ = v122
                                }
                            } else {
                                v124 := GenS12{F0: false, F1: 42}
                                _ = v124
                            }
                        }
                    }
                    v116 = -2147483648
                }
            }
            _ = fi115
            _ = v116
        }
    default:
        v7 = (v7 | v7)
    }
    switch -2147483648 {
    case 4294967295:
        var v125 float32 = 0.0
        _ = v125
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
