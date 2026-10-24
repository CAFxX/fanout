package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int; F1 uint; F2 uint8 }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 int64; F1 uint64; F2 int16 }
type GenS10 struct { F0 int; F1 int }
type GenS11 struct { F0 string; F1 uint8; F2 string }
type GenS12 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint16; F1 string; F2 uint }
type GenS17 struct { F0 int32; F1 uint }

func main() {
    var v1 uint8 = 225
    var v2 uint8 = v1
    v3 := make(map[int]bool)
    v4 := []int{9223372036854775807, -690, -2147483648, 1}
    var v5 uint = 124
    var v6 uint64 = 2
    if false {
        switch "hello" {
        case "\n\t":
            if (false && (697 > 0)) {
                delete(v3, -951)
                v6 -= 18446744073709551615
            }
        case "hello":
            switch "\n\t" {
            case "":
                v1 = (v2 | v1)
            }
        default:
            var v7 uint8 = 1
            _ = v7
        }
        var v8 bool = false
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 float32 = 1.5
            switch ((("\n\t" + "\n\t") + ("hello" + ("\n\t" + "a"))) + "a") {
            case "foo bar":
                if true {
                    var v12 int64 = 752
                    fmt.Println(v2, v11, fi10)
                    switch (fi10 << ((369) & 63)) {
                    case -1:
                        var v13 uint8 = (v1 + v1)
                        _ = v13
                    case 669:
                        fmt.Println(fi9)
                        v12 = 1
                        if (fi9 != (fi9 << (((v5 + v5)) & 63))) {
                            v14 := GenS1{F0: fi9}
                            if true {
                                if (!(!true)) {
                                    if (((false || true) && false) && true) {
                                        v15 := GenS2{F0: (int(v5) <= fi10)}
                                        fmt.Println(v2, v2, v11)
                                        v16 := GenS3{F0: 42, F1: int8(v2)}
                                        _ = v15
                                        _ = v16
                                    } else {
                                        var v17 int8 = 1
                                        v2 += v1
                                        if (true && false) {
                                            if false {
                                                fmt.Println(v1, fi9, fi9)
                                                var v18 int16 = 81
                                                _ = v18
                                            } else {
                                                var v19 bool = (9223372036854775807 > 9223372036854775807)
                                                fmt.Println(v5, fi10)
                                                var v20 float64 = 1.5
                                                _ = v19
                                                _ = v20
                                            }
                                            if (false && (359 == 680)) {
                                                switch 315 {
                                                case 32767:
                                                    v2 = uint8(v6)
                                                case 128:
                                                    switch 32768 {
                                                    case 2147483647:
                                                        fmt.Println(v1)
                                                    case 127:
                                                        v21 := GenS1{F0: (fi9 * 32768)}
                                                        _ = v21
                                                    }
                                                default:
                                                    var v22 uint8 = v2
                                                    _ = v22
                                                }
                                            } else {
                                                var v23 uint8 = v2
                                                fmt.Println(fi9, fi10)
                                                switch "hello" {
                                                case "a":
                                                    v23 = (v2 | v2)
                                                }
                                                _ = v23
                                            }
                                            fmt.Println(v2, fi9)
                                        }
                                        _ = v17
                                    }
                                    var v24 string = (("a" + "a") + "a")
                                    _ = v24
                                } else {
                                    v2 = (v2 >> (((v5 + 255)) & 63))
                                    fmt.Println(v5, v2, v2)
                                    v25 := GenS2{F0: true}
                                    _ = v25
                                }
                                v6 -= (v6 ^ (v6 & 776))
                                v26 := v3[32768]
                                _ = v26
                            } else {
                                var v27 float32 = (((0.1 * -2.25) + 0.1) * 0.1)
                                switch (fi10 % ((-1) | 1)) {
                                case 32767:
                                    switch "hello" {
                                    case "a":
                                        var v28 int8 = 1
                                        _ = v28
                                    case "hello":
                                        v1 = v2
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    v11 -= -2.25
                                default:
                                    switch int(v2) {
                                    case 128:
                                        var v29 uint8 = (v1 | 127)
                                        _ = v29
                                    }
                                }
                                _ = v27
                            }
                            v12 = v12
                            _ = v14
                        } else {
                            var v30 float32 = 3.14159
                            v3[fi9] = (!(false || true))
                            _ = v30
                        }
                    default:
                        if (!(!(("\n\t" + "a") < ("\n\t" + "a")))) {
                            v31 := GenS1{F0: fi10}
                            _ = v31
                        } else {
                            v32 := GenS1{F0: (fi9 & (-(fi10)))}
                            _ = v32
                        }
                    }
                    v4 = append(v4, fi9)
                    v33 := GenS1{F0: (fi9 ^ int(v6))}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v12
                    _ = v33
                } else {
                    if false {
                        if false {
                            if len(v4) > 0 {
                                v4[0] = fi9
                            }
                            v11 = v11
                            var v34 int16 = -1
                            _ = v34
                        } else {
                            if (false || false) {
                                v3[int(v5)] = true
                            } else {
                                if (v5 <= (^(v5))) {
                                    v35 := GenS4{F0: (int64(v5) >= int64(fi9)), F1: fi10}
                                    if len(v4) > 0 {
                                        v4[0] = (fi9 << ((v5) & 63))
                                    }
                                    v2 += (^(v2))
                                    _ = v35
                                }
                                var v36 uint16 = 42
                                var v37 uint = 42
                                v38 := v3[fi10]
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        }
                        var v39 float32 = (v11 / ((v11 + v11) / v11))
                        fmt.Println(v2)
                        _ = v39
                    } else {
                        if false {
                            v40 := GenS2{F0: (0 >= 2147483648)}
                            if false {
                                switch ("hello" + "hello") {
                                case "\n\t":
                                    v41 := GenS5{F0: 255, F1: v5}
                                    _ = v41
                                case "foo bar":
                                    if true {
                                        fmt.Println(v2, v2)
                                        var v42 int64 = -1
                                        v1 = v2
                                        _ = v42
                                    }
                                    switch 1 {
                                    case -9223372036854775808:
                                        var v43 uint32 = uint32(v6)
                                        _ = v43
                                    case -739:
                                        if false {
                                            fmt.Println(v2, v6)
                                            if (false || (!true)) {
                                                v44 := v3[int(v2)]
                                                if (9223372036854775807 == -9223372036854775808) {
                                                    var v45 uint16 = 1
                                                    _ = v45
                                                }
                                                _ = v44
                                            }
                                            v5 = (v5 * (^(v5)))
                                            if (float64((float32(0.0) - v11)) == (0.1 * 3.14159)) {
                                                var v46 bool = true
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v46
                                            }
                                            var v47 int64 = 42
                                            v3[629] = false
                                            v1 -= v1
                                            var v48 float64 = -2.25
                                            v49 := GenS6{F0: (v5 - 1000), F1: -1}
                                            v11 -= ((-2.25 * (v11 / v11)) - (3.14159 - (v11 - -2.25)))
                                            if false {
                                                v11 += (((1.5 / v11) + float32(0.0)) - (v11 * v11))
                                            } else {
                                                v50 := GenS2{F0: false}
                                                v51 := GenS7{F0: 42, F1: (v5 | v5), F2: 2}
                                                if ((("hello" + "") < "foo bar") || (true && (1000 > 1))) {
                                                    if (3.14159 < (-2.25 / v48)) {
                                                        if true {
                                                            v48 -= (1e10 / v48)
                                                            var v52 uint8 = (v1 >> ((v5) & 63))
                                                            var v53 int = fi9
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        fmt.Println(v2, v2, fi10)
                                                    }
                                                    var v54 int = 0
                                                    _ = len(v4)
                                                    _ = v54
                                                } else {
                                                    if true {
                                                        if ((^(v1)) >= v1) {
                                                            v55 := GenS7{F0: (fi9 + (fi10 % (((fi10 % ((fi9) | 1))) | 1))), F1: v5, F2: v2}
                                                            switch ("" + "a") {
                                                            case "foo bar":
                                                                if (0.1 >= 1.5) {
                                                                    v1 = v1
                                                                    var v56 bool = (!true)
                                                                    _ = v56
                                                                }
                                                            case "":
                                                                switch "" {
                                                                case "":
                                                                    v57 := GenS7{F0: (fi9 & fi10), F1: v5, F2: v1}
                                                                    _ = v57
                                                                }
                                                            case "hello":
                                                                v58 := GenS7{F0: (fi9 | -1), F1: 255, F2: 40}
                                                                _ = v58
                                                            default:
                                                                switch fi10 {
                                                                case 2147483647:
                                                                    fmt.Println(v48)
                                                                case 127:
                                                                    var v59 int64 = v47
                                                                    _ = v59
                                                                case -1:
                                                                    if (-32768 > -582) {
                                                                        v3[fi10] = true
                                                                    }
                                                                }
                                                            }
                                                            if (1 < 0) {
                                                                delete(v3, (^(fi9)))
                                                                v60 := GenS8{F0: -115, F1: (v5 ^ (v5 | (v5 * 0)))}
                                                                v48 += v48
                                                                _ = v60
                                                            }
                                                            fmt.Println(fi10, v6, v11)
                                                            var v61 int64 = v47
                                                            v1 -= 42
                                                            if false {
                                                                switch (fi10 % ((fi10) | 1)) {
                                                                case -450:
                                                                    switch (("foo bar" + ("\n\t" + "foo bar")) + "a") {
                                                                    case "":
                                                                        fmt.Println(fi9)
                                                                    }
                                                                default:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v62 := GenS9{F0: (v61 * (v47 & v61)), F1: v6, F2: 556}
                                                                        _ = v62
                                                                    }
                                                                }
                                                                var v63 uint32 = 1
                                                                var v64 uint8 = v2
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                            _ = v55
                                                            _ = v61
                                                        } else {
                                                            v6 = (^(v6))
                                                            var v65 float32 = v11
                                                            _ = v65
                                                        }
                                                        v1 -= v1
                                                        v1 -= 41
                                                        var v66 int = fi10
                                                        if (false && false) {
                                                            fmt.Println(fi9, v6)
                                                            v67 := GenS8{F0: 1, F1: (v5 + 255)}
                                                            var v68 float32 = 0.0
                                                            v66 = (fi9 & (^(fi9)))
                                                            v47 -= v47
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        _ = v66
                                                    } else {
                                                        var v69 uint8 = (v1 + 96)
                                                        _ = v69
                                                    }
                                                    v70 := GenS10{F0: (fi9 | fi10), F1: fi9}
                                                    switch (fi9 - (fi9 & (fi9 << ((v5) & 63)))) {
                                                    case 127:
                                                        v71 := GenS8{F0: 100, F1: (v5 & v5)}
                                                        _ = v71
                                                    case 9223372036854775807:
                                                        var v72 int16 = 32767
                                                        v73 := GenS9{F0: v47, F1: (^(v6)), F2: v72}
                                                        _ = v72
                                                        _ = v73
                                                    case 42:
                                                        v1 = (v1 | 165)
                                                    default:
                                                        if (0 <= 401) {
                                                            v5 = (v5 >> ((18446744073709551615) & 63))
                                                        }
                                                    }
                                                    _ = v70
                                                }
                                                _ = v50
                                                _ = v51
                                            }
                                            var v74 float32 = 3.14159
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                            _ = v74
                                        }
                                    default:
                                        v1 = (v2 - 255)
                                    }
                                case "":
                                    var v75 bool = true
                                    var v76 uint64 = v6
                                    v77 := GenS1{F0: fi9}
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                                v6 = v6
                                switch (("foo bar" + ("" + ("foo bar" + "\n\t"))) + ("foo bar" + ("" + "\n\t"))) {
                                case "a":
                                    var v78 string = "foo bar"
                                    _ = v78
                                case "foo bar":
                                    var v79 int16 = 42
                                    _ = v79
                                case "":
                                    v80 := GenS11{F0: "", F1: v2, F2: ("a" + "\n\t")}
                                    _ = v80
                                default:
                                    if false {
                                        fmt.Println(v1)
                                    } else {
                                        fmt.Println(v11, fi9, v1)
                                    }
                                }
                            }
                            var v81 float32 = v11
                            switch ("a" + ("a" + "\n\t")) {
                            case "":
                                if len(v4) > 0 {
                                    v4[0] = fi9
                                }
                            case "\n\t":
                                var v82 float64 = 3.14159
                                switch fi9 {
                                case 9223372036854775807:
                                    fmt.Println(v82, v11, v1)
                                default:
                                    if true {
                                        fmt.Println(v81, v2, v81)
                                        var v83 int8 = -58
                                        _ = v83
                                    } else {
                                        var v84 uint16 = 32767
                                        _ = v84
                                    }
                                }
                                _ = v82
                            case "a":
                                var v85 string = "foo bar"
                                _ = v85
                            }
                            v86 := GenS12{F0: 113, F1: v1, F2: (v2 * 127)}
                            v87 := GenS13{F0: fi9}
                            var v88 string = ((("a" + "hello") + ("hello" + "hello")) + ((("a" + ("a" + (("\n\t" + "\n\t") + "hello"))) + "hello") + ""))
                            _ = v40
                            _ = v81
                            _ = v86
                            _ = v87
                            _ = v88
                        }
                        v89 := GenS7{F0: (^(fi9)), F1: v5, F2: v2}
                        _ = v89
                    }
                    var v90 uint16 = 531
                    _ = v90
                }
            case "":
                if ((v6 >> ((210) & 63)) > (v6 >> (((v5 & 4294967295)) & 63))) {
                    if (!(32768 > fi10)) {
                        v91 := GenS14{F0: 635}
                        var v92 int = (fi10 - fi9)
                        v3[-2147483648] = false
                        _ = v91
                        _ = v92
                    } else {
                        var v93 int64 = 1
                        v4 = append(v4, (fi10 & (^(fi10))))
                        if len(v4) > 0 {
                            v4[0] = fi10
                        }
                        var v94 int8 = 127
                        _ = v93
                        _ = v94
                    }
                    var v95 uint16 = 32767
                    _ = v95
                } else {
                    var v96 float64 = 1.5
                    _ = v96
                }
            case "a":
                if false {
                    v6 = uint64(v5)
                } else {
                    var v97 bool = false
                    fmt.Println(v2)
                    _ = v97
                }
            }
            if true {
                v3[fi10] = true
                var v98 bool = true
                _ = v98
            }
            _ = fi10
            _ = v11
        }
        var v99 uint = v5
        for fi100 := 0; fi100 < 6; fi100++ {
            var v101 int8 = int8(v2)
            _ = fi100
            _ = v101
        }
        _ = fi9
        _ = v99
    }
    v102 := GenS4{F0: true, F1: 32768}
    v103 := v3[782]
    switch 79 {
    case 2147483647:
        var v104 uint64 = (v6 + v6)
        _ = v104
    }
    v105 := GenS15{F0: uint32(v1)}
    v2 += v1
    var v106 uint8 = v2
    v107 := GenS5{F0: 0, F1: (v5 - uint(v6))}
    var v108 float32 = -2.25
    var v109 int = 127
    var v110 uint32 = 2147483648
    v111 := GenS2{F0: ((v2 ^ (^(v1))) <= v2)}
    if v103 {
        v112 := GenS14{F0: v110}
        v113 := GenS16{F0: 2, F1: "hello", F2: v5}
        var v114 uint = (^(v5))
        _ = v112
        _ = v113
        _ = v114
    }
    v115 := GenS2{F0: true}
    if (!v103) {
        _ = len(v3)
        v3[(v109 | -1)] = (((0.0 / -2.25) * ((3.14159 / (((-2.25 - (1e10 + 0.0)) - 3.14159) / -2.25)) - 1e10)) > 0.0)
    } else {
        v4 = append(v4, v109)
        for fi116 := 0; fi116 < 7; fi116++ {
            v5 -= v5
            _ = fi116
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi117 := 0; fi117 < 17; fi117++ {
        v110 = (v110 << ((18446744073709551615) & 63))
        for fi118 := 0; fi118 < 4; fi118++ {
            if true {
                v119 := GenS2{F0: (-629 == -98)}
                fmt.Println(v1, v2)
                var v120 int64 = 638
                v109 = int(v110)
                _ = v119
                _ = v120
            } else {
                if v103 {
                    v121 := GenS17{F0: -1, F1: (v5 * (v5 | 0))}
                    switch ((("\n\t" + "") + "hello") + ("foo bar" + (("foo bar" + ("\n\t" + "a")) + "\n\t"))) {
                    case "\n\t":
                        v122 := GenS12{F0: 255, F1: 127, F2: 128}
                        _ = v122
                    case "hello":
                        switch ("hello" + ("" + "\n\t")) {
                        case "foo bar":
                            v106 -= 42
                        }
                    case "a":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v108 -= (v108 / 3.14159)
                    var v123 uint32 = v110
                    _ = v121
                    _ = v123
                }
            }
            v106 += 128
            _ = fi118
        }
        _ = fi117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v106:", v106)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v103
    _ = v105
    _ = v106
    _ = v107
    _ = v108
    _ = v109
    _ = v110
    _ = v111
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
