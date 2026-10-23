package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 int64; F1 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint; F1 string }
func genFn0(p0 int16) int16 {
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: v7}
        _ = fi8
        _ = v9
    }
    var v10 int32 = 291
    _ = p0
    _ = v10
    _ = v7
    return (p0 >> ((2) & 63))
}
func genFn1(p0 uint32, p1 uint16) uint {
    var v11 string = "foo bar"
    var v12 int = 32767
    v13 := []int{121, (-(v12)), (-(v12)), (v12 - (v12 & 32767))}
    if (!(p1 == 47)) {
        p1 = (^(p1))
        for fi14 := 0; fi14 < 8; fi14++ {
            switch fi14 {
            case 2147483647:
                for fi15 := 0; fi15 < 14; fi15++ {
                    p1 -= p1
                    _ = fi15
                }
            }
            v12 = (fi14 % ((489) | 1))
            v11 = v11
            _ = fi14
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            if (1 >= 1) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint = 0
                    var v19 uint16 = (^(p1))
                    v12 -= 594
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            v12 = fi16
            _ = fi16
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    var v1 bool = true
    var v2 bool = true
    v3 := make(map[string]string)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 222
    v6 := []int{-1, 32767, 128}
    if v1 {
        v20 := GenS2{F0: 906}
        for fi21 := 0; fi21 < 8; fi21++ {
            for fi22 := 0; fi22 < 19; fi22++ {
                if v1 {
                    if v1 {
                        switch fi21 {
                        case 24:
                            fmt.Println(v1, v2, v5)
                        default:
                            if v1 {
                                v23 := GenS1{F0: 2}
                                fmt.Println(v5)
                                var v24 int16 = -1
                                _ = v23
                                _ = v24
                            } else {
                                v2 = (true && v2)
                                v25 := GenS2{F0: 1}
                                v2 = (true && (!(false || false)))
                                _ = v25
                            }
                        }
                    } else {
                        switch "foo bar" {
                        case "":
                            _ = len(v3)
                        case "foo bar":
                            v1 = (1 == 2)
                        case "a":
                            _ = len(v3)
                        }
                        switch "a" {
                        case "":
                            var v26 uint32 = 2147483647
                            _ = v26
                        default:
                            var v27 int64 = -1
                            _ = v27
                        }
                    }
                    fmt.Println(v2, fi21, fi22)
                    if v2 {
                        var v28 int16 = 42
                        if (true && v2) {
                            v2 = ("a" != ("\n\t" + ""))
                        }
                        _ = v28
                    } else {
                        if v2 {
                            _ = len(v4)
                        } else {
                            v5 -= 119
                            if (1.5 >= 0.1) {
                                if v1 {
                                    if v2 {
                                        if len(v6) > 0 {
                                            v6[0] = 42
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v1 = v1
                                    var v29 float32 = 0.0
                                    _ = v29
                                }
                            } else {
                                v1 = v1
                                var v30 uint = 4294967295
                                var v31 uint32 = 4294967295
                                _ = v30
                                _ = v31
                            }
                            var v32 int = -263
                            _ = v32
                        }
                        var v33 int32 = 950
                        v34 := GenS2{F0: 149}
                        _ = v33
                        _ = v34
                    }
                }
                var v35 float32 = float32(0.0)
                _ = fi22
                _ = v35
            }
            var v36 bool = true
            for fi37 := 0; fi37 < 16; fi37++ {
                v36 = v1
                v38 := GenS3{F0: "\n\t"}
                v3[("hello" + "hello")] = "a"
                _ = fi37
                _ = v38
            }
            _ = fi21
            _ = v36
        }
        fmt.Println(v2, v1)
        _ = v20
    }
    v5 = (^(v5))
    if (18446744073709551615 != 42) {
        v1 = ((float32(float64(((((0.1 + 1.5) + 0.1) * 0.0) * 3.14159))) * (float32(((3.14159 / 1.0) / 1.0)) - (((3.14159 - 0.1) + 0.0) * -2.25))) < (((1e10 / ((1.5 / 0.1) * 0.1)) + (((0.0 + ((float32(((3.14159 - 0.1) + 1.5)) + 0.1) / 1.5)) / 1.5) - ((0.1 - (float32((3.14159 + ((-2.25 / 1.5) / (-2.25 + (1.5 + 1e10))))) + 1e10)) + float32((0.1 / -2.25))))) + (0.1 - 1.5)))
        if (3.14159 == -2.25) {
            v1 = (-9223372036854775808 != 9223372036854775807)
            v39 := GenS4{F0: v1, F1: 128, F2: 42}
            if (-2.25 == (0.0 + (3.14159 - 1e10))) {
                v5 += 127
            } else {
                var v40 uint64 = 805
                fmt.Println(v1, v5, v5)
                _ = v40
            }
            _ = v39
        } else {
            v41 := GenS5{F0: 987}
            for fi42 := 0; fi42 < 8; fi42++ {
                fmt.Println(fi42, fi42)
                _ = fi42
            }
            _ = v41
        }
        v2 = (genFn1(4294967295, 1) < 494)
    }
    var v43 bool = ((v2 || (-46 >= 0)) && v2)
    var v44 uint = 42
    v3[""] = ("a" + ("" + ""))
    fmt.Println(v5, v1, v43)
    switch 32768 {
    case -9223372036854775808:
        if (int64(v44) >= 1) {
            v2 = v43
            v45 := GenS2{F0: (v44 | genFn1(2, 65535))}
            _ = v45
        } else {
            v5 += v5
            switch "a" {
            case "\n\t":
                var v46 uint64 = uint64(v5)
                _ = v46
            case "hello":
                if v2 {
                    v44 += (v44 / ((136) | 1))
                } else {
                    var v47 uint8 = 2
                    _ = v47
                }
            case "a":
                var v48 string = ((("\n\t" + "hello") + ("\n\t" + "\n\t")) + ("foo bar" + ("hello" + ((((("hello" + "foo bar") + "a") + "") + "a") + "\n\t"))))
                _ = v48
            }
        }
    case -33:
        var v49 float32 = float32(-2.25)
        _ = v49
    case -220:
        if (518 > 42) {
            if ((0.0 + 0.0) != float32(0.1)) {
                v1 = (-721 != 574)
                switch 32767 {
                case -153:
                    var v50 int32 = 356
                    _ = v50
                case 128:
                    if (-556 >= genFn0(-32768)) {
                        for fi51 := 0; fi51 < 4; fi51++ {
                            v3[("\n\t" + "\n\t")] = ("\n\t" + "\n\t")
                            _ = fi51
                        }
                        if v1 {
                            switch (("hello" + "hello") + ("foo bar" + ("\n\t" + ("" + "hello")))) {
                            case "":
                                if (v43 && (!v1)) {
                                    var v52 int64 = 957
                                    _ = v52
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        var v53 uint32 = 2
                                        _ = v53
                                    default:
                                        v44 -= (^(v44))
                                    }
                                    switch -3 {
                                    case 32768:
                                        var v54 int = 1
                                        _ = v54
                                    default:
                                        var v55 uint64 = uint64(v5)
                                        _ = v55
                                    }
                                }
                            case "\n\t":
                                var v56 uint32 = 4294967295
                                _ = v56
                            }
                            fmt.Println(v43, v5, v1)
                        } else {
                            if false {
                                v57 := GenS4{F0: false, F1: 1, F2: 255}
                                switch 32768 {
                                case -427:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                var v58 uint = (^(v44))
                                _ = v57
                                _ = v58
                            }
                            v59 := GenS1{F0: 0}
                            delete(v3, ("hello" + "\n\t"))
                            _ = v59
                        }
                        switch "" {
                        case "hello":
                            switch ("\n\t" + ("foo bar" + "foo bar")) {
                            case "\n\t":
                                _ = len(v6)
                            case "a":
                                for fi60 := 0; fi60 < 1; fi60++ {
                                    v2 = ((4294967295 <= 0) || false)
                                    if (-2.25 <= float32(3.14159)) {
                                        switch ("a" + ("foo bar" + "hello")) {
                                        case "foo bar":
                                            delete(v4, 0)
                                        default:
                                            v61 := GenS4{F0: v1, F1: 65535, F2: 42}
                                            _ = v61
                                        }
                                    } else {
                                        v62 := GenS4{F0: v2, F1: 255, F2: 255}
                                        v63 := GenS6{F0: (fi60 % (((-(fi60))) | 1)), F1: 32767}
                                        var v64 int64 = -1
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    }
                                    _ = fi60
                                }
                            case "hello":
                                var v65 uint32 = 65535
                                _ = v65
                            default:
                                v6 = append(v6, -2147483648)
                            }
                        case "":
                            switch "foo bar" {
                            case "":
                                var v66 int64 = -1
                                _ = v66
                            case "\n\t":
                                v67 := GenS7{F0: true}
                                _ = v67
                            }
                        case "\n\t":
                            v5 += (v5 - 5)
                        default:
                            fmt.Println(v43)
                        }
                    }
                }
                if (-2.25 <= ((0.0 - (float64(-2.25) * (1e10 - 0.0))) / 1.0)) {
                    for fi68 := 0; fi68 < 9; fi68++ {
                        v69 := GenS8{F0: 0, F1: -9223372036854775808}
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            var v70 float64 = -2.25
                            _ = v70
                        case "\n\t":
                            v71 := GenS6{F0: fi68, F1: 768}
                            _ = v71
                        case "":
                            for fi72 := 0; fi72 < 13; fi72++ {
                                v44 += genFn1(4294967295, 42)
                                _ = fi72
                            }
                        }
                        for fi73 := 0; fi73 < 20; fi73++ {
                            switch ("hello" + "\n\t") {
                            case "foo bar":
                                v43 = ("\n\t" >= ((("foo bar" + "hello") + "foo bar") + ("hello" + "\n\t")))
                            case "\n\t":
                                switch -9223372036854775808 {
                                case -41:
                                    if len(v6) > 0 {
                                        v6[0] = fi68
                                    }
                                default:
                                    if v1 {
                                        var v74 uint = (v44 % ((v44) | 1))
                                        v75 := GenS3{F0: "hello"}
                                        _ = v74
                                        _ = v75
                                    } else {
                                        if v1 {
                                            delete(v4, 2147483647)
                                            v76 := GenS1{F0: 32767}
                                            _ = v76
                                        } else {
                                            if v2 {
                                                if (!v2) {
                                                    v44 = genFn1(42, 231)
                                                    v77 := GenS9{F0: (v5 << (((v44 - 65535)) & 63)), F1: 42}
                                                    _ = v77
                                                } else {
                                                    if ((42 == 42) || v2) {
                                                        v5 -= (v5 + (^(v5)))
                                                        fmt.Println(v44, fi68, fi68)
                                                        var v78 int32 = 2147483647
                                                        _ = v78
                                                    } else {
                                                        var v79 int16 = -32768
                                                        var v80 string = "hello"
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                    v5 = (^(v5))
                                                    if (fi68 >= fi68) {
                                                        v81 := GenS10{F0: -9223372036854775808, F1: 255}
                                                        if v43 {
                                                            var v82 int8 = 42
                                                            var v83 float32 = float32(0.1)
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        _ = v81
                                                    } else {
                                                        v84 := GenS10{F0: 42, F1: (v5 + 1)}
                                                        _ = v84
                                                    }
                                                }
                                            }
                                        }
                                        if (612 > 1) {
                                            v85 := v3["a"]
                                            fmt.Println(v43, v43, v44)
                                            v6 = append(v6, (fi68 | fi73))
                                            var v86 int8 = 126
                                            switch (v85 + (v85 + ("foo bar" + "foo bar"))) {
                                            case "":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            case "a":
                                                if v1 {
                                                    switch "a" {
                                                    case "hello":
                                                        var v87 int = (^(fi73))
                                                        _ = v87
                                                    case "":
                                                        if v1 {
                                                            v86 = v86
                                                            v2 = (genFn0(-796) > -32768)
                                                        }
                                                    default:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                    v88 := v3["foo bar"]
                                                    _ = v88
                                                }
                                            case "hello":
                                                v89 := GenS2{F0: 65535}
                                                _ = v89
                                            }
                                            v86 -= v86
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                }
                            }
                            var v90 uint = (v44 << ((v44) & 63))
                            v3[("" + "foo bar")] = ("" + "\n\t")
                            v1 = (-760 <= 642)
                            _ = fi73
                            _ = v90
                        }
                        _ = fi68
                        _ = v69
                    }
                    var v91 int8 = int8(v5)
                    _ = v91
                } else {
                    if v1 {
                        var v92 bool = (v44 > v44)
                        _ = v92
                    } else {
                        v44 -= genFn1(65535, 555)
                    }
                    if v2 {
                        fmt.Println(v43)
                        var v93 int64 = 177
                        v93 -= (v93 - v93)
                        _ = v93
                    } else {
                        var v94 uint8 = (v5 ^ v5)
                        for fi95 := 0; fi95 < 7; fi95++ {
                            v94 -= v94
                            var v96 int = fi95
                            fmt.Println(v94, v43, v96)
                            _ = fi95
                            _ = v96
                        }
                        for fi97 := 0; fi97 < 14; fi97++ {
                            switch 0 {
                            case 789:
                                if ("\n\t" == ("\n\t" + "")) {
                                    var v98 uint8 = (v94 & v5)
                                    _ = v98
                                }
                            case -2147483648:
                                for fi99 := 0; fi99 < 15; fi99++ {
                                    if (v2 && v2) {
                                        v1 = v43
                                        v94 = (v5 + 127)
                                        switch (("a" + "hello") + "foo bar") {
                                        case "a":
                                            fmt.Println(fi97)
                                        case "hello":
                                            v100 := GenS11{F0: (0.0 >= (3.14159 - (0.0 / 3.14159)))}
                                            _ = v100
                                        default:
                                            _ = len(v4)
                                        }
                                        var v101 uint64 = 42
                                        if (0 == 0) {
                                            v102 := GenS12{F0: v5, F1: (v5 + (^(v94)))}
                                            v2 = (v2 || v1)
                                            switch "hello" {
                                            case "a":
                                                v103 := GenS7{F0: true}
                                                _ = v103
                                            case "":
                                                var v104 bool = v43
                                                _ = v104
                                            default:
                                                v6 = append(v6, (fi97 >> (((v44 - v44)) & 63)))
                                            }
                                            if (32767 == -1) {
                                                if v43 {
                                                    fmt.Println(v1, fi99, v101)
                                                    var v105 uint32 = 2147483647
                                                    if false {
                                                        v106 := GenS8{F0: -2147483648, F1: (-(fi99))}
                                                        v5 = v5
                                                        v107 := GenS1{F0: 0}
                                                        var v108 string = "\n\t"
                                                        _ = v106
                                                        _ = v107
                                                        _ = v108
                                                    } else {
                                                        var v109 int16 = 0
                                                        v110 := GenS13{F0: v105}
                                                        v5 -= (v94 ^ uint8(v101))
                                                        _ = v109
                                                        _ = v110
                                                    }
                                                    _ = v105
                                                }
                                                var v111 string = (("foo bar" + "\n\t") + ("" + "\n\t"))
                                                v5 = uint8(v44)
                                                _ = v111
                                            } else {
                                                switch (fi99 & 9223372036854775807) {
                                                case 2147483647:
                                                    v101 = (v101 % (((^(v101))) | 1))
                                                case -9223372036854775808:
                                                    var v112 uint16 = 602
                                                    _ = v112
                                                default:
                                                    var v113 uint32 = 2147483648
                                                    _ = v113
                                                }
                                            }
                                            v114 := GenS14{F0: 0, F1: "hello"}
                                            _ = v102
                                            _ = v114
                                        }
                                        _ = v101
                                    }
                                    _ = fi99
                                }
                            case 32768:
                                v4[2147483647] = (false && v43)
                            }
                            fmt.Println(v2, v1)
                            var v115 uint8 = 128
                            _ = fi97
                            _ = v115
                        }
                        _ = v94
                    }
                }
            } else {
                v2 = (!v43)
                fmt.Println(v5, v43)
            }
        } else {
            v116 := GenS10{F0: 9223372036854775807, F1: 42}
            v3["foo bar"] = "\n\t"
            _ = v116
        }
    }
    v117 := GenS14{F0: genFn1(844, 140), F1: "foo bar"}
    v43 = (1e10 > (0.0 / 1.5))
    var v118 int8 = 42
    fmt.Println(v2, v118, v2)
    var v119 string = ("" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
