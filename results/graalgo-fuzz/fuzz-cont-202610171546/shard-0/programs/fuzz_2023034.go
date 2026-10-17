package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int32; F1 int16; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 int16; F2 int64 }
type GenS8 struct { F0 bool; F1 uint8; F2 string }
type GenS9 struct { F0 uint16; F1 uint; F2 uint }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int32; F1 uint16; F2 bool }
type GenS12 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS13 struct { F0 int32; F1 uint64 }
type GenS14 struct { F0 uint8; F1 int64 }
type GenS15 struct { F0 uint; F1 int; F2 uint8 }
type GenS16 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int8, p2 int16) int16 {
    v6 := GenS1{F0: -194, F1: 9223372036854775807, F2: ""}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float64 = (1e10 * (((1e10 * 0.0) * 3.14159) - 0.1))
        v9 := GenS2{F0: p1, F1: p0, F2: 2}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            p2 = int16(p0)
            if (false && (32767 >= 65535)) {
                p1 = int8(p2)
            } else {
                if false {
                    switch -485 {
                    case 360:
                        fmt.Println(fi11)
                    case 32768:
                        v12 := []int{}
                        _ = v12
                    case 128:
                        v13 := GenS2{F0: (-(p1)), F1: (p0 + 127), F2: 1}
                        _ = v13
                    }
                    var v14 string = "hello"
                    var v15 int = 32768
                    _ = v14
                    _ = v15
                } else {
                    if ((!false) || (!(!true))) {
                        v16 := GenS2{F0: p1, F1: p0, F2: 4294967295}
                        v17 := GenS3{F0: -31, F1: true}
                        _ = v16
                        _ = v17
                    } else {
                        if false {
                            p0 = (p0 * (^(p0)))
                        }
                    }
                    var v18 string = "\n\t"
                    v19 := []int{-542, -9223372036854775808, fi10, fi10}
                    _ = v18
                    _ = v19
                }
            }
            v20 := []int{(fi11 - fi11)}
            _ = fi11
            _ = v20
        }
        _ = fi10
    }
    switch -722 {
    case 2147483647:
        v21 := GenS2{F0: p1, F1: (p0 << ((4294967295) & 63)), F2: 0}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1() uint64 {
    var v22 string = "\n\t"
    switch 686 {
    case 499:
        switch "a" {
        case "hello":
            v23 := []int{986, 127, 1}
            _ = v23
        case "a":
            v24 := GenS4{F0: "", F1: 1}
            _ = v24
        default:
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 bool = false
                _ = fi25
                _ = v26
            }
        }
    case 2147483647:
        v27 := GenS4{F0: "foo bar", F1: 150}
        _ = v27
    }
    var v28 bool = true
    var v29 uint32 = 403
    _ = v22
    _ = v28
    _ = v29
    return 380
}

func main() {
    v1 := []string{((("" + "\n\t") + "\n\t") + ("foo bar" + (("hello" + "foo bar") + "\n\t"))), "foo bar", ((("a" + "a") + (("" + ("hello" + "\n\t")) + "hello")) + "a"), ("a" + "")}
    var v2 float32 = -2.25
    var v3 uint16 = 369
    var v4 uint64 = 4294967295
    var v5 float64 = (3.14159 * (1e10 * 1e10))
    var v30 bool = ((4294967295 == 65535) || (!true))
    v31 := GenS1{F0: 32767, F1: -259, F2: "foo bar"}
    if len(v1) > 0 {
        v1[0] = ("" + "foo bar")
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        switch (fi32 >> ((0) & 63)) {
        case -573:
            for fi33 := 0; fi33 < 7; fi33++ {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                v30 = (2147483647 > 42)
                _ = fi33
            }
        case 42:
            if (137 > v3) {
                var v34 bool = (("" >= "foo bar") && ((!true) && v30))
                _ = v34
            } else {
                for fi35 := 0; fi35 < 2; fi35++ {
                    v2 -= ((float32(v5) + v2) * (v2 + (v2 + v2)))
                    v36 := GenS3{F0: -128, F1: true}
                    v4 += (v4 + (v4 >> ((42) & 63)))
                    _ = fi35
                    _ = v36
                }
                var v37 int8 = -128
                _ = v37
            }
        case 891:
            for fi38 := 0; fi38 < 1; fi38++ {
                var v39 int64 = 0
                if (32767 < 255) {
                    v30 = v30
                    var v40 int16 = 32767
                    _ = v40
                } else {
                    v4 = (v4 * (v4 >> ((914) & 63)))
                    v30 = false
                    if ((2147483647 == 69) || v30) {
                        var v41 uint16 = (^(v3))
                        if (!(v2 > (3.14159 * v2))) {
                            v42 := GenS5{F0: 680, F1: 1000, F2: 1}
                            v43 := GenS6{F0: 1}
                            v44 := GenS7{F0: 0, F1: 0, F2: (v39 % ((int64(v41)) | 1))}
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        var v45 uint16 = v41
                        v46 := GenS8{F0: ("a" != ""), F1: 127, F2: ((("" + (("\n\t" + "") + (("\n\t" + ("" + "\n\t")) + ("a" + "a")))) + "foo bar") + "hello")}
                        switch "hello" {
                        case "foo bar":
                            var v47 uint64 = (v4 / (((v4 >> ((2) & 63))) | 1))
                            _ = v47
                        case "a":
                            switch "foo bar" {
                            case "foo bar":
                                var v48 float32 = v2
                                _ = v48
                            case "a":
                                if v30 {
                                    v4 -= (v4 + (^(v4)))
                                } else {
                                    _ = len(v1)
                                    var v49 int64 = (^(v39))
                                    var v50 bool = (!(!true))
                                    _ = v49
                                    _ = v50
                                }
                            case "hello":
                                var v51 uint32 = 157
                                _ = v51
                            }
                        case "":
                            _ = len(v1)
                            if len(v1) > 0 {
                                v1[0] = ("hello" + ("a" + (("foo bar" + "foo bar") + "hello")))
                            }
                        default:
                            v52 := GenS7{F0: -1, F1: 32767, F2: (v39 >> ((1) & 63))}
                            _ = v52
                        }
                        _ = v41
                        _ = v45
                        _ = v46
                    }
                }
                v53 := GenS2{F0: 113, F1: 1, F2: genFn1()}
                v54 := GenS5{F0: 0, F1: 434, F2: 219}
                fmt.Println(fi38)
                _ = fi38
                _ = v39
                _ = v53
                _ = v54
            }
        default:
            for fi55 := 0; fi55 < 5; fi55++ {
                v56 := GenS8{F0: v30, F1: 128, F2: ("" + "\n\t")}
                fmt.Println(v2, v2)
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                if v30 {
                    switch "a" {
                    case "":
                        fmt.Println(v4)
                    case "foo bar":
                        v30 = v30
                    case "hello":
                        if (false && v30) {
                            v57 := GenS2{F0: 0, F1: 67, F2: (v4 * (v4 << ((uint(v3)) & 63)))}
                            if v30 {
                                fmt.Println(v3, v4)
                                switch (fi32 - fi32) {
                                case 887:
                                    if true {
                                        v58 := GenS6{F0: 42}
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        _ = v58
                                    } else {
                                        v5 = ((float64(0.0) + v5) - ((3.14159 - 1e10) + float64(((v2 * (v2 * v2)) + ((1.5 * 1e10) * v2)))))
                                        v5 = v5
                                        var v59 int = (fi32 + (fi55 - fi32))
                                        _ = v59
                                    }
                                case -9223372036854775808:
                                    v60 := GenS9{F0: 61, F1: 0, F2: 255}
                                    _ = v60
                                case 886:
                                    var v61 int8 = 127
                                    _ = v61
                                }
                            }
                            _ = v57
                        }
                    }
                } else {
                    v3 = 1
                }
                _ = fi55
                _ = v56
            }
        }
        v3 -= v3
        _ = fi32
    }
    if true {
        if v30 {
            var v62 int = 32768
            _ = v62
        }
        for fi63 := 0; fi63 < 15; fi63++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi63
        }
        if v30 {
            v5 += (v5 + 1e10)
        }
    } else {
        var v64 float32 = 3.14159
        var v65 int8 = 127
        _ = v64
        _ = v65
    }
    v66 := GenS10{F0: -740}
    v5 -= (v5 - (((v5 * v5) + v5) / v5))
    var v67 uint32 = 2
    var v68 uint64 = (v4 << ((0) & 63))
    switch "\n\t" {
    case "\n\t":
        v69 := GenS1{F0: -177, F1: 564, F2: ("" + "hello")}
        _ = v69
    default:
        v70 := GenS6{F0: 42}
        _ = v70
    }
    for fi71 := 0; fi71 < 18; fi71++ {
        switch 995 {
        case -2147483648:
            for fi72 := 0; fi72 < 9; fi72++ {
                v30 = true
                v73 := GenS11{F0: -2147483648, F1: v3, F2: v30}
                v30 = (!(!v30))
                v1 = append(v1, "hello")
                var v74 uint32 = uint32(v3)
                _ = fi72
                _ = v73
                _ = v74
            }
        case -694:
            v75 := GenS12{F0: 255, F1: v68, F2: (v68 % (((v68 << ((0) & 63))) | 1))}
            _ = v75
        case 82:
            fmt.Println(v2, v68, v2)
        default:
            v76 := GenS8{F0: (true && true), F1: 37, F2: "hello"}
            _ = v76
        }
        _ = fi71
    }
    var v77 bool = v30
    if v30 {
        switch "foo bar" {
        case "a":
            for fi78 := 0; fi78 < 14; fi78++ {
                var v79 int64 = -1
                fmt.Println(v30, v5, v3)
                _ = fi78
                _ = v79
            }
        default:
            switch ("foo bar" + "\n\t") {
            case "":
                _ = len(v1)
            case "hello":
                v30 = v30
            case "a":
                v3 += (v3 - 65535)
            }
        }
        if len(v1) > 0 {
            v1[0] = ("a" + ((("hello" + "a") + "foo bar") + "\n\t"))
        }
    } else {
        v4 = (v4 & (^(v4)))
    }
    for fi80 := 0; fi80 < 7; fi80++ {
        if (0.0 >= (v2 * v2)) {
            switch fi80 {
            case 42:
                var v81 uint64 = 929
                _ = v81
            case -9223372036854775808:
                var v82 int16 = -1
                _ = v82
            default:
                switch -1 {
                case 32767:
                    var v83 uint64 = (v68 - v68)
                    _ = v83
                }
            }
        }
        v2 = float32(float64(((v2 + (1.5 * float32(v5))) * 0.1)))
        _ = fi80
    }
    for fi84 := 0; fi84 < 17; fi84++ {
        var v85 int16 = -32768
        if (2 >= 128) {
            v67 -= (^(v67))
            switch "" {
            case "\n\t":
                for fi86 := 0; fi86 < 20; fi86++ {
                    v87 := GenS7{F0: -2147483648, F1: (v85 << ((uint(v68)) & 63)), F2: 0}
                    var v88 float32 = v2
                    v3 -= uint16(v68)
                    _ = fi86
                    _ = v87
                    _ = v88
                }
            case "":
                switch "a" {
                case "foo bar":
                    for fi89 := 0; fi89 < 17; fi89++ {
                        v90 := GenS8{F0: (!true), F1: 128, F2: "\n\t"}
                        var v91 int8 = 100
                        var v92 uint32 = (v67 / (((v67 / ((v67) | 1))) | 1))
                        switch 0 {
                        case -15:
                            var v93 int16 = v85
                            _ = v93
                        case 1:
                            if v77 {
                                if v30 {
                                    switch fi89 {
                                    case 42:
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    case -733:
                                        v94 := GenS13{F0: -677, F1: uint64(v3)}
                                        _ = v94
                                    default:
                                        var v95 uint = uint(v3)
                                        _ = v95
                                    }
                                    if false {
                                        v96 := GenS5{F0: 0, F1: 305, F2: v85}
                                        v92 = (v92 - 42)
                                        v1 = append(v1, (("\n\t" + "foo bar") + "\n\t"))
                                        _ = v96
                                    }
                                    var v97 bool = (v2 > v2)
                                    _ = v97
                                } else {
                                    if (uint64(v3) != v4) {
                                        _ = len(v1)
                                        if (fi84 < -9223372036854775808) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v98 float64 = float64(v2)
                                            _ = v98
                                        } else {
                                            var v99 uint64 = (v68 << ((2) & 63))
                                            var v100 uint16 = v3
                                            fmt.Println(v99, v67, v91)
                                            _ = v100
                                            _ = v99
                                        }
                                        v2 = v2
                                        v101 := GenS9{F0: v3, F1: 546, F2: 14}
                                        _ = v101
                                    }
                                    v102 := GenS10{F0: 418}
                                    if false {
                                        if ((!(fi84 == 1)) || (!v30)) {
                                            fmt.Println(fi84)
                                            if (!(!((!((int32(v92) > 2147483647) && v77)) || (((^(v4)) == (v4 & (v68 ^ 18446744073709551615))) && (1 < 1))))) {
                                                switch 11 {
                                                case -1:
                                                    if (!v30) {
                                                        v103 := GenS12{F0: 334, F1: genFn1(), F2: v68}
                                                        var v104 uint8 = 128
                                                        _ = v103
                                                        _ = v104
                                                    } else {
                                                        v92 -= v92
                                                        v105 := GenS14{F0: 213, F1: 42}
                                                        v1 = append(v1, "\n\t")
                                                        _ = v105
                                                    }
                                                default:
                                                    v1 = append(v1, ("hello" + "a"))
                                                }
                                                v68 = v68
                                                v68 -= 4294967295
                                            }
                                            var v106 uint32 = 0
                                            _ = v106
                                        } else {
                                            v2 += 1e10
                                            v92 = (v67 * (^(v92)))
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                    }
                                    if (!v77) {
                                        switch fi84 {
                                        case -1:
                                            fmt.Println(v68, v5, v3)
                                        }
                                        var v107 int16 = (v85 ^ int16(v3))
                                        _ = v107
                                    }
                                    _ = v102
                                }
                                var v108 uint = 1
                                v109 := GenS7{F0: 42, F1: int16(v68), F2: -9223372036854775808}
                                _ = v108
                                _ = v109
                            }
                        default:
                            v5 -= 0.1
                        }
                        if v77 {
                            if len(v1) > 0 {
                                v1[0] = ("" + "a")
                            }
                            switch "\n\t" {
                            case "foo bar":
                                var v110 uint8 = 42
                                _ = v110
                            default:
                                v77 = false
                            }
                            var v111 uint = 42
                            _ = v111
                        }
                        _ = fi89
                        _ = v90
                        _ = v91
                        _ = v92
                    }
                case "a":
                    switch fi84 {
                    case 0:
                        var v112 uint16 = v3
                        _ = v112
                    case -799:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 613:
                        if v30 {
                            if (!(v77 && true)) {
                                v2 = 0.0
                                switch (fi84 * fi84) {
                                case -1:
                                    var v113 uint32 = (v67 ^ v67)
                                    _ = v113
                                case 128:
                                    v114 := GenS15{F0: 65535, F1: fi84, F2: 128}
                                    _ = v114
                                case 1:
                                    var v115 float32 = ((((v2 - (v2 + v2)) * ((v2 - v2) * 0.1)) * float32(-2.25)) * 3.14159)
                                    _ = v115
                                }
                                if true {
                                    switch "a" {
                                    case "a":
                                        var v116 string = "\n\t"
                                        _ = v116
                                    default:
                                        var v117 int8 = 100
                                        _ = v117
                                    }
                                } else {
                                    var v118 int16 = 144
                                    _ = v118
                                }
                            }
                        }
                    }
                }
            }
            v119 := GenS3{F0: 0, F1: v77}
            _ = v119
        } else {
            for fi120 := 0; fi120 < 6; fi120++ {
                v30 = (-775 < 0)
                switch -553 {
                case 32767:
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                case 2147483647:
                    var v121 int8 = 87
                    _ = v121
                default:
                    v122 := GenS16{F0: 2}
                    _ = v122
                }
                _ = fi120
            }
            v85 += (v85 % ((v85) | 1))
            fmt.Println(v3)
        }
        var v123 uint8 = 42
        _ = fi84
        _ = v123
        _ = v85
    }
    v124 := GenS6{F0: -285}
    fmt.Println(v3, v3, v30)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v124
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v68
    _ = v77
}
