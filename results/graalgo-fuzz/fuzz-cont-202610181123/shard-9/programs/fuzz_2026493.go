package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint; F1 int16; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 uint32; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 string; F2 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 int8; F2 bool }
type GenS11 struct { F0 string }
func genFn0(p0 uint, p1 int8, p2 uint) uint8 {
    p1 = p1
    var v6 float64 = ((1.5 - 3.14159) / 0.1)
    var v7 int64 = 546
    switch 128 {
    case 127:
        switch "foo bar" {
        case "\n\t":
            v8 := []int{-1, 4294967295, 0}
            _ = v8
        }
    default:
        var v9 uint8 = 255
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}
func genFn1() uint32 {
    var v10 int8 = 42
    fmt.Println(v10)
    _ = v10
    return 915
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    var v3 uint8 = 2
    var v4 float64 = 3.14159
    var v5 uint32 = 2147483648
    v11 := v2[("a" + "\n\t")]
    if false {
        switch "foo bar" {
        case "foo bar":
            v3 += 2
        }
    }
    if false {
        var v12 int8 = v1
        v1 -= 1
        v13 := GenS1{F0: 42}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS2{F0: 18446744073709551615, F1: v3, F2: 722}
            if false {
                for fi16 := 0; fi16 < 4; fi16++ {
                    switch ("foo bar" + ("a" + ("" + ""))) {
                    case "foo bar":
                        if (float32(float64(1.5)) <= 0.0) {
                            var v17 uint = 255
                            v3 += 2
                            _ = v17
                        }
                    }
                    v1 -= (v1 ^ v1)
                    fmt.Println(fi16, v11)
                    _ = fi16
                }
                v18 := []int{-762, (fi14 << ((4294967295) & 63)), (fi14 & fi14), 42}
                var v19 uint32 = uint32(v3)
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 1; fi20++ {
                    v21 := GenS3{F0: 255, F1: 32767, F2: 568}
                    fmt.Println(v4, fi20)
                    if (((false || true) || true) && (!true)) {
                        if true {
                            switch fi14 {
                            case 1:
                                var v22 int32 = -757
                                _ = v22
                            case 42:
                                fmt.Println(v1)
                            case -129:
                                v3 = (^(v3))
                            default:
                                v1 = (v1 & (v1 * v1))
                            }
                            if false {
                                v2["\n\t"] = genFn1()
                                var v23 int32 = 42
                                _ = v23
                            }
                        }
                    }
                    _ = fi20
                    _ = v21
                }
                var v24 float32 = 3.14159
                for fi25 := 0; fi25 < 16; fi25++ {
                    if true {
                        var v26 uint64 = 818
                        _ = len(v2)
                        v27 := GenS4{F0: -386}
                        _ = v26
                        _ = v27
                    }
                    _ = fi25
                }
                _ = v24
            }
            _ = fi14
            _ = v15
        }
        v28 := GenS2{F0: 4294967295, F1: (v3 & (v3 * v3)), F2: 65535}
        if ((!(1 < 98)) || (!true)) {
            v3 -= (^(v3))
            fmt.Println(v3, v5, v3)
            if true {
                for fi29 := 0; fi29 < 14; fi29++ {
                    switch (-(fi29)) {
                    case -2147483648:
                        var v30 uint = 0
                        _ = v30
                    case -9223372036854775808:
                        for fi31 := 0; fi31 < 11; fi31++ {
                            if (((v4 / ((1.5 * v4) + v4)) - float64((3.14159 - ((0.0 + -2.25) + 1e10)))) <= 1.5) {
                                v32 := []int{(fi31 * fi29), fi29, fi31}
                                _ = v32
                            }
                            if (false && (false && false)) {
                                v4 -= 0.0
                                v33 := GenS2{F0: 1, F1: (v3 / ((v3) | 1)), F2: 532}
                                if ((-(v1)) < v1) {
                                    if true {
                                        var v34 int = (fi29 << ((0) & 63))
                                        v1 = v1
                                        v35 := v2["\n\t"]
                                        _ = v34
                                        _ = v35
                                    } else {
                                        fmt.Println(v4)
                                    }
                                } else {
                                    v1 += 100
                                }
                                _ = v33
                            }
                            if (106 <= 0) {
                                v1 += -21
                                v36 := GenS2{F0: 4294967295, F1: 42, F2: 275}
                                _ = v36
                            }
                            _ = fi31
                        }
                    default:
                        if false {
                            v4 += v4
                            v4 = 0.0
                            var v37 int32 = 2147483647
                            _ = v37
                        }
                    }
                    v38 := []int{9223372036854775807, (fi29 >> ((65535) & 63))}
                    _ = fi29
                    _ = v38
                }
                for fi39 := 0; fi39 < 10; fi39++ {
                    v1 = (-(v1))
                    if ((0.1 / (1e10 - 1.5)) == ((0.1 - (-2.25 - ((1.5 + ((float32(v4) * 1.5) - 1.5)) - (1e10 - 0.0)))) - 0.0)) {
                        if ((false && (42 == 724)) || true) {
                            var v40 uint64 = 969
                            if (!(false || ((v3 < (v3 - genFn0(65535, 127, 2))) || ((^(v3)) > uint8(v40))))) {
                                var v41 int = (fi39 / ((fi39) | 1))
                                v42 := v2[""]
                                for fi43 := 0; fi43 < 6; fi43++ {
                                    v44 := v2["\n\t"]
                                    v45 := GenS5{F0: v40}
                                    v46 := v2[""]
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                                _ = v41
                                _ = v42
                            } else {
                                v3 += uint8(v40)
                                v3 -= 2
                                v2["\n\t"] = v11
                            }
                            _ = v40
                        } else {
                            var v47 int16 = 42
                            v11 = (v11 | genFn1())
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                v48 := []int{}
                                _ = v48
                            case "a":
                                var v49 uint16 = 255
                                _ = v49
                            default:
                                v50 := GenS2{F0: 0, F1: genFn0(540, v1, 65535), F2: 338}
                                _ = v50
                            }
                            _ = v47
                        }
                        v4 += v4
                        if (26 < 4294967295) {
                            v51 := GenS1{F0: 426}
                            _ = v51
                        } else {
                            v52 := v2["\n\t"]
                            fmt.Println(fi39, v11)
                            for fi53 := 0; fi53 < 6; fi53++ {
                                v1 = 0
                                _ = fi53
                            }
                            _ = v52
                        }
                    } else {
                        v54 := GenS3{F0: 741, F1: -1, F2: 1}
                        v55 := GenS5{F0: 115}
                        _ = v54
                        _ = v55
                    }
                    _ = fi39
                }
            } else {
                var v56 int8 = (v1 | -128)
                v57 := []int{}
                var v58 int16 = -226
                _ = v56
                _ = v57
                _ = v58
            }
        }
        _ = v28
    }
    var v59 int32 = 469
    v60 := GenS3{F0: 830, F1: 738, F2: 0}
    fmt.Println(v3, v1)
    v4 += ((0.1 * v4) - v4)
    v1 = v1
    if true {
        if ((!false) || ((!(18446744073709551615 <= 65535)) && false)) {
            v5 = 411
            switch (("a" + ("foo bar" + "hello")) + (("foo bar" + "a") + "")) {
            case "a":
                if false {
                    v61 := GenS6{F0: 532, F1: v11, F2: -78}
                    if false {
                        var v62 int64 = -9223372036854775808
                        _ = v62
                    }
                    _ = v61
                } else {
                    switch ("" + "hello") {
                    case "foo bar":
                        var v63 int8 = v1
                        _ = v63
                    default:
                        for fi64 := 0; fi64 < 1; fi64++ {
                            var v65 uint8 = v3
                            v4 += (1.5 + v4)
                            _ = fi64
                            _ = v65
                        }
                    }
                    if false {
                        v66 := GenS4{F0: -9223372036854775808}
                        v67 := GenS7{F0: 42}
                        _ = v66
                        _ = v67
                    } else {
                        v68 := GenS8{F0: 65535, F1: "", F2: (v5 / (((v11 | (^(v11)))) | 1))}
                        switch 0 {
                        case 775:
                            var v69 uint32 = 65535
                            _ = v69
                        case 37:
                            for fi70 := 0; fi70 < 7; fi70++ {
                                var v71 int16 = 352
                                var v72 uint16 = 1
                                v73 := GenS4{F0: -806}
                                _ = fi70
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                        }
                        var v74 float32 = 1.5
                        _ = v68
                        _ = v74
                    }
                }
            case "\n\t":
                var v75 int16 = -1
                _ = v75
            }
        } else {
            if true {
                if (!false) {
                    v76 := []int{156, -399, 32767}
                    v5 = uint32(v3)
                    fmt.Println(v3, v4)
                    _ = v76
                } else {
                    if ((v59 << ((408) & 63)) >= v59) {
                        var v77 int = 128
                        fmt.Println(v3, v59)
                        if false {
                            fmt.Println(v5, v59, v5)
                            for fi78 := 0; fi78 < 1; fi78++ {
                                var v79 int64 = -965
                                _ = fi78
                                _ = v79
                            }
                            fmt.Println(v5)
                        } else {
                            var v80 uint32 = v5
                            for fi81 := 0; fi81 < 8; fi81++ {
                                v11 -= v5
                                for fi82 := 0; fi82 < 6; fi82++ {
                                    v4 -= 0.0
                                    _ = fi82
                                }
                                if true {
                                    switch fi81 {
                                    case 4294967295:
                                        v59 = v59
                                    default:
                                        var v83 uint16 = 32767
                                        _ = v83
                                    }
                                    var v84 float64 = 1.5
                                    _ = v84
                                }
                                _ = fi81
                            }
                            var v85 int16 = 1000
                            _ = v80
                            _ = v85
                        }
                        _ = v77
                    } else {
                        var v86 int32 = 42
                        if false {
                            v87 := GenS9{F0: (v3 * 11)}
                            _ = v87
                        } else {
                            switch (("" + "foo bar") + ("" + "hello")) {
                            case "hello":
                                v88 := GenS3{F0: 4294967295, F1: 0, F2: -447}
                                _ = v88
                            case "foo bar":
                                for fi89 := 0; fi89 < 1; fi89++ {
                                    v86 = (v59 % ((v86) | 1))
                                    if (!(3.14159 >= 3.14159)) {
                                        v5 = (v5 * v11)
                                    }
                                    var v90 string = ""
                                    _ = fi89
                                    _ = v90
                                }
                            case "a":
                                v4 = (float64(-2.25) / (((1e10 / v4) * (v4 / 1.0)) + v4))
                            default:
                                switch -2147483648 {
                                case -132:
                                    for fi91 := 0; fi91 < 8; fi91++ {
                                        if ((v5 << ((55) & 63)) < genFn1()) {
                                            if true {
                                                v3 = 0
                                                v92 := GenS8{F0: 433, F1: ("" + ("foo bar" + "hello")), F2: (v11 << ((4294967295) & 63))}
                                                if false {
                                                    if (3.14159 == ((-2.25 * v4) + (0.1 / (v4 + (1e10 / (v4 + 3.14159)))))) {
                                                        fmt.Println(v1, v4, v59)
                                                        var v93 float64 = (((-2.25 / (3.14159 / v4)) + (v4 * (v4 * ((v4 * 0.0) + v4)))) + -2.25)
                                                        var v94 int8 = -113
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    v4 -= (v4 * -2.25)
                                                } else {
                                                    v4 = (1e10 - float64(0.0))
                                                }
                                                _ = v92
                                            }
                                            fmt.Println(v5, v3, v3)
                                            var v95 float64 = v4
                                            _ = v95
                                        }
                                        var v96 bool = false
                                        _ = fi91
                                        _ = v96
                                    }
                                }
                            }
                            v4 = -2.25
                        }
                        _ = v86
                    }
                    v59 = -53
                }
                if (2147483647 >= genFn1()) {
                    fmt.Println(v11, v5, v1)
                } else {
                    v97 := []int{713, -162, -1}
                    var v98 int64 = -782
                    _ = v97
                    _ = v98
                }
                var v99 uint16 = 438
                _ = v99
            }
        }
        var v100 uint64 = 18446744073709551615
        _ = v100
    } else {
        if true {
            v101 := []int{0, -9223372036854775808, -9223372036854775808}
            var v102 int8 = v1
            v11 -= v5
            _ = v101
            _ = v102
        } else {
            switch "hello" {
            case "\n\t":
                v103 := []int{1, -1}
                _ = v103
            case "a":
                if false {
                    if (((0.0 + v4) * (0.0 / v4)) == v4) {
                        var v104 uint16 = 2
                        v105 := GenS10{F0: -798, F1: (^(v1)), F2: true}
                        v3 -= (v3 | (v3 / ((v3) | 1)))
                        _ = v104
                        _ = v105
                    } else {
                        var v106 string = ("\n\t" + "foo bar")
                        v107 := GenS3{F0: 42, F1: -32768, F2: (v59 | v59)}
                        fmt.Println(v3, v5, v4)
                        _ = v106
                        _ = v107
                    }
                } else {
                    v108 := []int{32767, 2147483647, -755, 1}
                    delete(v2, "\n\t")
                    var v109 uint8 = (v3 >> ((700) & 63))
                    _ = v108
                    _ = v109
                }
            default:
                v3 -= v3
            }
            v5 = v5
            v59 = 1
        }
        var v110 int16 = 527
        fmt.Println(v5, v11, v110)
        _ = v110
    }
    v111 := v2[((("hello" + "hello") + "hello") + "a")]
    fmt.Println(v4, v1)
    var v112 uint32 = v5
    v113 := GenS11{F0: ("\n\t" + "foo bar")}
    delete(v2, "\n\t")
    v11 = genFn1()
    for fi114 := 0; fi114 < 16; fi114++ {
        v115 := []int{(fi114 >> ((42) & 63)), (fi114 - 32767)}
        _ = fi114
        _ = v115
    }
    var v116 uint8 = 2
    v117 := []int{9223372036854775807, 128, 128, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v117:", len(v117), v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v11
    _ = v111
    _ = v112
    _ = v113
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
