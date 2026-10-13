package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint32; F1 string; F2 int8 }
type GenS7 struct { F0 bool; F1 uint32; F2 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v4 := []int{357, 42, 2147483647, 32767}
    v5 := GenS1{F0: "", F1: -32768}
    switch 9223372036854775807 {
    case 128:
        v1 = 83
    default:
        v6 := GenS1{F0: ("foo bar" + ("hello" + "a")), F1: -32768}
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v3[("foo bar" + "\n\t")] = fi7
        var v8 int32 = -897
        if (((v8 | (v8 | (-(v8)))) < (-(v8))) && (101 != 100)) {
            var v9 int64 = 314
            _ = v9
        } else {
            if true {
                var v10 uint32 = uint32(v1)
                _ = v10
            }
            var v11 int8 = 127
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    var v12 uint32 = 0
    fmt.Println(v1)
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS2{F0: "a", F1: ("" + "a")}
            if (!false) {
                v12 = (v12 & (v12 / ((v12) | 1)))
                v16 := GenS3{F0: 1}
                _ = v16
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi14
            _ = v15
        }
        var v17 uint8 = 177
        _ = fi13
        _ = v17
    }
    if (false || (-1 >= 1)) {
        v1 -= 128
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 int = fi18
            v4 = append(v4, 309)
            _ = fi18
            _ = v19
        }
        v1 = v1
    }
    if false {
        v1 -= (^(v1))
        fmt.Println(v1, v12)
    } else {
        switch ("hello" + "hello") {
        case "hello":
            v20 := GenS4{F0: 0, F1: v12}
            _ = v20
        default:
            var v21 uint = 0
            _ = v21
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 float32 = 1e10
            v24 := v2["a"]
            _ = fi22
            _ = v23
            _ = v24
        }
        if len(v4) > 0 {
            v4[0] = 652
        }
    }
    fmt.Println(v1, v1)
    if (!true) {
        if false {
            if (!(!("hello" == "foo bar"))) {
                fmt.Println(v12, v12)
            }
            for fi25 := 0; fi25 < 14; fi25++ {
                var v26 int64 = 240
                _ = fi25
                _ = v26
            }
            fmt.Println(v1)
        } else {
            v1 = 42
        }
        var v27 int8 = -98
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 float32 = 3.14159
            v12 += 42
            v4 = append(v4, 258)
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    var v30 float32 = -2.25
    v30 -= float32((float64(v30) + ((float64(0.1) - 0.1) * 1e10)))
    v30 += (v30 / (((v30 - v30) / v30) + v30))
    fmt.Println(v1, v1, v1)
    switch "" {
    case "foo bar":
        v1 -= v1
    case "\n\t":
        for fi31 := 0; fi31 < 20; fi31++ {
            if true {
                var v32 uint8 = (v1 << ((1) & 63))
                if (((-1 == 119) && false) && true) {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        if true {
                            var v34 int64 = -9223372036854775808
                            v30 -= v30
                            _ = v34
                        }
                        v30 += ((v30 - (1.5 + v30)) / 0.1)
                        switch 32767 {
                        case 633:
                            v35 := GenS4{F0: 644, F1: uint32(v32)}
                            _ = v35
                        case 42:
                            switch (("\n\t" + "a") + "hello") {
                            case "foo bar":
                                switch (("foo bar" + "") + ("" + "hello")) {
                                case "\n\t":
                                    _ = len(v2)
                                case "hello":
                                    v36 := GenS5{F0: (v12 * 2147483647), F1: -32768, F2: 0}
                                    _ = v36
                                case "foo bar":
                                    v37 := GenS4{F0: 130, F1: v12}
                                    _ = v37
                                default:
                                    if false {
                                        v12 = 0
                                        if ("hello" > (("a" + "\n\t") + "\n\t")) {
                                            var v38 int64 = 9223372036854775807
                                            v39 := GenS2{F0: (("hello" + (("\n\t" + ("" + "hello")) + "")) + ""), F1: "foo bar"}
                                            if ((1 < -2147483648) && false) {
                                                var v40 int8 = 0
                                                _ = v40
                                            } else {
                                                v32 = (v1 - (^(v32)))
                                                var v41 uint = 255
                                                _ = v41
                                            }
                                            fmt.Println(v32, v30, v12)
                                            switch fi33 {
                                            case -1:
                                                v38 -= 42
                                            case 4294967295:
                                                if ((1e10 + -2.25) > float32(((0.0 + ((((-2.25 - -2.25) / ((3.14159 / 0.1) + 0.1)) - float64(v30)) / -2.25)) - -2.25))) {
                                                    v1 = 1
                                                    v42 := v3[("" + ("\n\t" + ("a" + "foo bar")))]
                                                    _ = v42
                                                }
                                            default:
                                                v38 += (v38 | 9223372036854775807)
                                            }
                                            var v43 uint = 418
                                            _ = v38
                                            _ = v39
                                            _ = v43
                                        } else {
                                            v1 = v1
                                            switch fi33 {
                                            case 2147483647:
                                                v32 -= (^(v32))
                                            default:
                                                var v44 bool = true
                                                _ = v44
                                            }
                                            var v45 float32 = v30
                                            _ = v45
                                        }
                                        var v46 int = fi31
                                        v47 := GenS3{F0: -1}
                                        _ = v46
                                        _ = v47
                                    } else {
                                        v48 := GenS4{F0: 255, F1: (v12 - 68)}
                                        var v49 int32 = -1
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            default:
                                var v50 uint8 = (^(v1))
                                var v51 uint64 = 1
                                var v52 int = (fi33 * 4294967295)
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        }
                        var v53 uint32 = v12
                        _ = fi33
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 14; fi54++ {
                        if true {
                            v32 -= 0
                            if true {
                                if true {
                                    var v55 int16 = -32768
                                    v56 := GenS6{F0: 911, F1: "a", F2: 100}
                                    _ = v55
                                    _ = v56
                                }
                                v57 := GenS7{F0: false, F1: 2, F2: true}
                                var v58 string = ("\n\t" + "foo bar")
                                _ = v57
                                _ = v58
                            } else {
                                var v59 int8 = 127
                                if true {
                                    v60 := GenS8{F0: (^(fi54))}
                                    v61 := GenS4{F0: 42, F1: (^(v12))}
                                    var v62 string = ""
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := v3["a"]
                                    _ = v63
                                }
                                if true {
                                    v64 := GenS4{F0: 42, F1: v12}
                                    var v65 bool = false
                                    _ = v64
                                    _ = v65
                                }
                                _ = v59
                            }
                        }
                        if false {
                            var v66 int32 = -650
                            v67 := GenS9{F0: 1}
                            _ = v66
                            _ = v67
                        } else {
                            v30 = (0.0 * v30)
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v68 int64 = 316
                        _ = fi54
                        _ = v68
                    }
                    for fi69 := 0; fi69 < 11; fi69++ {
                        if true {
                            switch "a" {
                            case "hello":
                                v12 = v12
                            case "a":
                                fmt.Println(v30, v30)
                            default:
                                if false {
                                    v70 := GenS6{F0: 2147483648, F1: "a", F2: 127}
                                    if (true || (true || (65535 <= 1))) {
                                        v1 -= 0
                                        v30 -= (((float32(3.14159) / v30) + v30) * -2.25)
                                        v71 := GenS4{F0: 2, F1: 2}
                                        _ = v71
                                    } else {
                                        if (42 > 1) {
                                            if false {
                                                fmt.Println(v30)
                                                if true {
                                                    if true {
                                                        if (v30 == v30) {
                                                            v32 = (v1 + v1)
                                                        }
                                                    } else {
                                                        fmt.Println(v30, v30)
                                                        v30 = ((v30 / 1e10) / (0.1 + ((((v30 + v30) / v30) - 1.5) + v30)))
                                                        switch ("" + "") {
                                                        case "a":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v72 int32 = 1
                                                                _ = v72
                                                            case "\n\t":
                                                                switch (("hello" + "a") + "a") {
                                                                case "hello":
                                                                    _ = len(v4)
                                                                case "":
                                                                    var v73 uint32 = v12
                                                                    _ = v73
                                                                case "a":
                                                                    fmt.Println(v32, fi31)
                                                                }
                                                            case "":
                                                                if true {
                                                                    var v74 uint8 = 2
                                                                    fmt.Println(v30)
                                                                    _ = v74
                                                                }
                                                            default:
                                                                fmt.Println(v12, fi69, v1)
                                                            }
                                                        }
                                                    }
                                                }
                                                var v75 int = int(v32)
                                                _ = v75
                                            } else {
                                                var v76 float64 = float64(1e10)
                                                if true {
                                                    var v77 int = fi69
                                                    v1 = (v1 << ((18446744073709551615) & 63))
                                                    var v78 int32 = 2147483647
                                                    _ = v77
                                                    _ = v78
                                                }
                                                if false {
                                                    switch "hello" {
                                                    case "hello":
                                                        var v79 uint32 = (v12 | v12)
                                                        _ = v79
                                                    case "foo bar":
                                                        var v80 int32 = 0
                                                        _ = v80
                                                    default:
                                                        v76 = 1e10
                                                    }
                                                }
                                                _ = v76
                                            }
                                            fmt.Println(fi31)
                                            var v81 uint32 = v12
                                            _ = v81
                                        }
                                        v32 = (v1 % (((^(v32))) | 1))
                                        v32 -= v1
                                    }
                                    var v82 int = -983
                                    _ = v70
                                    _ = v82
                                } else {
                                    v4 = append(v4, 128)
                                }
                            }
                            switch fi31 {
                            case 487:
                                if (!(!false)) {
                                    v1 = (^(v1))
                                } else {
                                    v83 := GenS9{F0: 2}
                                    v84 := GenS4{F0: 42, F1: (v12 % ((v12) | 1))}
                                    var v85 float32 = 3.14159
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                }
                            default:
                                var v86 string = "a"
                                _ = v86
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = len(v3)
                        if (!true) {
                            v1 -= 39
                        } else {
                            var v87 int32 = -1
                            _ = v87
                        }
                        _ = fi69
                    }
                }
                for fi88 := 0; fi88 < 10; fi88++ {
                    var v89 int64 = 9223372036854775807
                    fmt.Println(v89, v1, v32)
                    var v90 int16 = -32768
                    if true {
                        v89 += (v89 % ((0) | 1))
                    } else {
                        v91 := GenS8{F0: fi88}
                        var v92 uint64 = 919
                        v12 = (v12 - 4294967295)
                        _ = v91
                        _ = v92
                    }
                    _ = fi88
                    _ = v89
                    _ = v90
                }
                _ = v32
            } else {
                for fi93 := 0; fi93 < 5; fi93++ {
                    var v94 int16 = 32767
                    switch "a" {
                    case "":
                        v94 -= -605
                    case "hello":
                        v2[(("hello" + "hello") + "hello")] = (v12 / (((v12 >> ((763) & 63))) | 1))
                    default:
                        v95 := GenS10{F0: true, F1: v94}
                        if true {
                            v96 := GenS1{F0: "", F1: (-(v94))}
                            var v97 int8 = 42
                            _ = v96
                            _ = v97
                        }
                        _ = v95
                    }
                    _ = len(v4)
                    _ = fi93
                    _ = v94
                }
                if false {
                    if ("hello" <= "a") {
                        var v98 int = (-(fi31))
                        var v99 uint32 = v12
                        for fi100 := 0; fi100 < 5; fi100++ {
                            v101 := v2["a"]
                            if (float32(0.0) != 0.0) {
                                var v102 uint = 2
                                _ = v102
                            } else {
                                if ((false || (!false)) && false) {
                                    v4 = append(v4, fi100)
                                    v4 = append(v4, (^(v98)))
                                    v99 = v99
                                } else {
                                    if ((v98 - (fi100 + (fi31 & (^(fi100))))) == fi100) {
                                        var v103 uint32 = uint32(v1)
                                        var v104 int64 = 0
                                        if ((v1 - v1) > (v1 >> ((18446744073709551615) & 63))) {
                                            var v105 uint32 = (v103 - (v99 ^ (v12 << ((2) & 63))))
                                            var v106 int16 = 1000
                                            _ = v105
                                            _ = v106
                                        }
                                        _ = v103
                                        _ = v104
                                    } else {
                                        if (!(!(-1 == 78))) {
                                            v101 -= (v12 & (v101 / ((v101) | 1)))
                                            var v107 uint32 = v101
                                            _ = len(v3)
                                            _ = v107
                                        }
                                        var v108 int32 = 0
                                        fmt.Println(v108, fi100)
                                        _ = v108
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v109 int64 = 1
                                    _ = v109
                                }
                                var v110 float64 = 3.14159
                                var v111 int64 = -1
                                if len(v4) > 0 {
                                    v4[0] = fi31
                                }
                                var v112 string = ("\n\t" + "\n\t")
                                _ = v110
                                _ = v111
                                _ = v112
                            }
                            _ = fi100
                            _ = v101
                        }
                        _ = v98
                        _ = v99
                    }
                    _ = len(v4)
                } else {
                    if (74 >= 42) {
                        if (255 >= 456) {
                            var v113 uint = 702
                            _ = v113
                        } else {
                            switch fi31 {
                            case -585:
                                fmt.Println(fi31, v12)
                            case 302:
                                var v114 uint64 = 815
                                _ = v114
                            default:
                                var v115 float32 = (v30 * (((v30 - v30) * 0.0) * v30))
                                _ = v115
                            }
                            _ = len(v2)
                        }
                        var v116 float64 = 1.5
                        var v117 float64 = -2.25
                        _ = v116
                        _ = v117
                    }
                    if len(v4) > 0 {
                        v4[0] = (fi31 | 2147483647)
                    }
                    for fi118 := 0; fi118 < 7; fi118++ {
                        v12 = v12
                        fmt.Println(fi31, v1, fi31)
                        _ = fi118
                    }
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v119 := v3[(((("\n\t" + "") + (("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))) + "foo bar")) + "hello") + "hello")]
            _ = fi31
            _ = v119
        }
    case "hello":
        v12 -= (v12 << ((65535) & 63))
    default:
        v30 += float32((float64(v30) - 1.5))
    }
    for fi120 := 0; fi120 < 19; fi120++ {
        for fi121 := 0; fi121 < 17; fi121++ {
            v122 := GenS8{F0: (fi120 % ((42) | 1))}
            if (v30 >= v30) {
                var v123 int64 = -1
                if ((float64(0.1) / 1.0) != float64((0.1 * v30))) {
                    var v124 uint = 255
                    if (32767 != int16(v123)) {
                        v125 := GenS3{F0: 862}
                        _ = v125
                    }
                    var v126 int64 = 823
                    _ = v124
                    _ = v126
                } else {
                    var v127 int16 = 32767
                    fmt.Println(fi120)
                    _ = v127
                }
                _ = v123
            }
            if false {
                fmt.Println(fi120, fi120, v30)
            } else {
                var v128 int32 = 1
                _ = len(v4)
                _ = v128
            }
            var v129 int8 = -61
            _ = fi121
            _ = v122
            _ = v129
        }
        _ = fi120
    }
    var v130 int64 = -9223372036854775808
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v130:", v130)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v130
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
