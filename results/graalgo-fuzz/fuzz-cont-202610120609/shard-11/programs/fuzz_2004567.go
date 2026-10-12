package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 int; F2 uint16 }
type GenS8 struct { F0 uint; F1 uint16; F2 bool }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 string }
type GenS12 struct { F0 int; F1 int8; F2 int8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS15 struct { F0 uint16; F1 int16 }
func genFn0() int16 {
    v6 := GenS1{F0: 42, F1: 908, F2: 843}
    var v7 float64 = 1.5
    v7 += v7
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8, (fi8 + -9223372036854775808), fi8, (^(fi8))}
        var v10 int = (-(fi8))
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 += -2.25
    _ = v6
    _ = v7
    return -273
}
func genFn1() uint32 {
    var v11 uint64 = 1
    v12 := []int{4294967295, 127, 9223372036854775807, 4294967295}
    var v13 string = "\n\t"
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 int8 = 127
    var v4 int16 = 1000
    var v5 string = (((((("hello" + "") + "foo bar") + "foo bar") + ("a" + "")) + "hello") + "foo bar")
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(v3, v4)
        v3 = v3
        var v15 int8 = (v3 - (v3 >> ((0) & 63)))
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 32767, F1: 42, F2: 152}
    fmt.Println(v3)
    fmt.Println(v4, v5, v4)
    if (1 >= 779) {
        var v17 int8 = (v3 << ((255) & 63))
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 uint32 = 0
            v3 = v17
            fmt.Println(fi18)
            _ = fi18
            _ = v19
        }
        var v20 float32 = 3.14159
        _ = v17
        _ = v20
    } else {
        var v21 uint16 = 42
        if true {
            switch -1 {
            case 2147483647:
                for fi22 := 0; fi22 < 11; fi22++ {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v24 := GenS2{F0: (v21 + v21), F1: fi23}
                        v3 -= v3
                        if (255 >= 128) {
                            if false {
                                v25 := GenS1{F0: v21, F1: 4294967295, F2: (v21 ^ v21)}
                                _ = v25
                            }
                            v26 := GenS2{F0: v21, F1: (^(fi23))}
                            _ = v26
                        } else {
                            v21 += v21
                        }
                        _ = fi23
                        _ = v24
                    }
                    _ = len(v2)
                    _ = fi22
                }
            case 128:
                for fi27 := 0; fi27 < 10; fi27++ {
                    if len(v1) > 0 {
                        v1[0] = v5
                    }
                    _ = fi27
                }
            case 0:
                if false {
                    if true {
                        if true {
                            v21 -= v21
                        } else {
                            v28 := GenS1{F0: (v21 + v21), F1: 380, F2: 65535}
                            _ = v28
                        }
                        v29 := GenS3{F0: 1}
                        for fi30 := 0; fi30 < 13; fi30++ {
                            var v31 uint = 61
                            var v32 string = (v5 + ((v5 + "\n\t") + v5))
                            v33 := GenS2{F0: (v21 % (((^(v21))) | 1)), F1: (fi30 << ((v31) & 63))}
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        _ = v29
                    }
                    for fi34 := 0; fi34 < 3; fi34++ {
                        for fi35 := 0; fi35 < 10; fi35++ {
                            var v36 int8 = v3
                            switch v5 {
                            case "hello":
                                fmt.Println(v36, fi35)
                            case "":
                                var v37 float32 = 1e10
                                _ = v37
                            case "foo bar":
                                switch v5 {
                                case "hello":
                                    var v38 int16 = 32767
                                    _ = v38
                                case "\n\t":
                                    _ = len(v2)
                                    if true {
                                        switch "foo bar" {
                                        case "":
                                            switch (fi34 & 4294967295) {
                                            case -696:
                                                var v39 uint32 = 0
                                                _ = v39
                                            default:
                                                v36 = (v36 ^ -1)
                                            }
                                        case "a":
                                            delete(v2, (v5 + v5))
                                        }
                                    } else {
                                        v40 := GenS1{F0: v21, F1: 972, F2: (v21 >> ((2) & 63))}
                                        _ = v40
                                    }
                                }
                            default:
                                v41 := GenS3{F0: 299}
                                _ = v41
                            }
                            _ = fi35
                            _ = v36
                        }
                        var v42 int32 = int32(v4)
                        _ = fi34
                        _ = v42
                    }
                    v2[v5] = true
                }
            default:
                v43 := GenS4{F0: 65535}
                _ = v43
            }
            if (((-507 >= 816) || true) || true) {
                v21 = (v21 / ((65535) | 1))
            }
            var v44 float32 = (((1.5 - -2.25) + 3.14159) + 3.14159)
            _ = v44
        }
        if ((v3 ^ v3) == (v3 * (v3 | v3))) {
            v4 += 0
        } else {
            v45 := GenS3{F0: 0}
            _ = v45
        }
        _ = v21
    }
    if ((("foo bar" + (v5 + ("foo bar" + "foo bar"))) + "a") <= (v5 + (v5 + v5))) {
        v46 := GenS1{F0: 65535, F1: 4294967295, F2: 1}
        v47 := GenS5{F0: (v3 * (v3 * 1)), F1: 2}
        if (0 > -147) {
            for fi48 := 0; fi48 < 11; fi48++ {
                if (!false) {
                    v49 := GenS4{F0: 821}
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 6; fi50++ {
                        v5 = v5
                        v51 := GenS2{F0: 1, F1: fi50}
                        if (4294967295 <= 321) {
                            var v52 float32 = -2.25
                            _ = v52
                        } else {
                            var v53 uint8 = 1
                            _ = v53
                        }
                        _ = fi50
                        _ = v51
                    }
                    v5 = (v5 + "\n\t")
                    fmt.Println(v3, v4)
                }
                _ = len(v2)
                if (!(false && true)) {
                    v4 += 1000
                    v54 := GenS6{F0: v5}
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 13; fi55++ {
                        v2[v5] = false
                        v56 := GenS7{F0: 32767, F1: fi55, F2: 2}
                        _ = fi55
                        _ = v56
                    }
                    v57 := GenS8{F0: 816, F1: 255, F2: (v5 <= (v5 + v5))}
                    var v58 int8 = (-(v3))
                    _ = v57
                    _ = v58
                }
                _ = fi48
            }
        } else {
            fmt.Println(v3, v3)
            v59 := GenS9{F0: 0, F1: v5}
            var v60 uint = 0
            _ = v59
            _ = v60
        }
        _ = v46
        _ = v47
    } else {
        fmt.Println(v5, v4)
        v61 := GenS1{F0: 708, F1: 4294967295, F2: 1}
        _ = v61
    }
    switch (("foo bar" + (("hello" + ((((v5 + "a") + v5) + (((v5 + v5) + v5) + (v5 + v5))) + "")) + (v5 + "a"))) + "hello") {
    case "":
        var v62 string = v5
        _ = v62
    case "a":
        var v63 uint32 = 475
        _ = v63
    default:
        if false {
            if (!(128 != 191)) {
                var v64 float32 = 3.14159
                if len(v1) > 0 {
                    v1[0] = (((v5 + ((((((v5 + ((v5 + v5) + ((("a" + v5) + v5) + "\n\t"))) + v5) + ("" + ("a" + v5))) + v5) + "a") + v5)) + "\n\t") + (v5 + "\n\t"))
                }
                _ = v64
            } else {
                v5 = "\n\t"
                v5 = (v5 + "\n\t")
                if (int8(v4) != v3) {
                    var v65 uint = 42
                    var v66 int32 = 44
                    _ = v65
                    _ = v66
                } else {
                    v1 = append(v1, v5)
                    v67 := GenS5{F0: 118, F1: 42}
                    var v68 uint8 = 3
                    _ = v67
                    _ = v68
                }
            }
            if (v3 == v3) {
                for fi69 := 0; fi69 < 16; fi69++ {
                    switch (fi69 ^ 32767) {
                    case -1:
                        v1 = append(v1, v5)
                    case -895:
                        v70 := GenS9{F0: (v3 / ((v3) | 1)), F1: ("hello" + ("hello" + v5))}
                        _ = v70
                    default:
                        v71 := GenS10{F0: 42}
                        _ = v71
                    }
                    _ = fi69
                }
                for fi72 := 0; fi72 < 3; fi72++ {
                    v73 := v2[v5]
                    if false {
                        if (genFn0() <= -254) {
                            for fi74 := 0; fi74 < 5; fi74++ {
                                _ = len(v1)
                                switch v5 {
                                case "":
                                    fmt.Println(fi72, v4)
                                case "\n\t":
                                    var v75 int8 = (v3 ^ -128)
                                    v76 := GenS9{F0: 0, F1: v5}
                                    v3 -= v3
                                    _ = v75
                                    _ = v76
                                default:
                                    v5 = v5
                                }
                                _ = fi74
                            }
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi77 := 0; fi77 < 8; fi77++ {
                                    switch (fi72 | fi77) {
                                    case 32768:
                                        v78 := GenS11{F0: 1, F1: v5}
                                        _ = v78
                                    case -9223372036854775808:
                                        var v79 uint = 42
                                        _ = v79
                                    }
                                    fmt.Println(fi72, fi77, v3)
                                    _ = fi77
                                }
                            }
                            for fi80 := 0; fi80 < 8; fi80++ {
                                if (v5 != (v5 + v5)) {
                                    v3 += (v3 * (-(v3)))
                                }
                                _ = fi80
                            }
                        }
                        var v81 float64 = 1.5
                        var v82 int8 = v3
                        _ = v81
                        _ = v82
                    } else {
                        for fi83 := 0; fi83 < 4; fi83++ {
                            if (float64(0.1) == 0.0) {
                                fmt.Println(v3)
                                v4 += 42
                                if (!v73) {
                                    v84 := GenS8{F0: 65535, F1: 0, F2: (v73 || v73)}
                                    v73 = true
                                    _ = v84
                                } else {
                                    if v73 {
                                        v85 := GenS10{F0: 42}
                                        v86 := GenS12{F0: (^(fi83)), F1: -48, F2: (v3 / ((v3) | 1))}
                                        v87 := GenS3{F0: -1}
                                        switch (fi72 - (fi72 ^ fi83)) {
                                        case 128:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            if v73 {
                                                fmt.Println(fi83, fi83, v4)
                                                v2[("hello" + v5)] = (true && false)
                                                var v88 string = "a"
                                                switch -9223372036854775808 {
                                                case -9223372036854775808:
                                                    v88 = v5
                                                    v89 := GenS5{F0: v3, F1: 621}
                                                    _ = v89
                                                case 2147483647:
                                                    var v90 uint8 = 1
                                                    _ = v90
                                                }
                                                _ = v88
                                            } else {
                                                v3 = -1
                                            }
                                        }
                                        var v91 uint16 = 42
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                        _ = v91
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v3 = 42
                                    v92 := v2[(v5 + v5)]
                                    v93 := v2["foo bar"]
                                    _ = v92
                                    _ = v93
                                }
                            } else {
                                switch 1 {
                                case -1:
                                    if (!(1 < 0)) {
                                        fmt.Println(v5, fi83, v5)
                                    } else {
                                        if (-1 != (fi72 | 359)) {
                                            v3 = v3
                                        }
                                        var v94 uint64 = 1
                                        fmt.Println(v4, v4)
                                        _ = v94
                                    }
                                }
                                v5 = ("" + v5)
                                fmt.Println(v5, v5, v73)
                            }
                            v73 = v73
                            _ = fi83
                        }
                        for fi95 := 0; fi95 < 15; fi95++ {
                            v4 -= v4
                            _ = fi95
                        }
                        var v96 uint16 = 2
                        _ = v96
                    }
                    if (-9223372036854775808 < 9223372036854775807) {
                        v73 = false
                        var v97 int = 288
                        v1 = append(v1, "\n\t")
                        _ = v97
                    }
                    _ = fi72
                    _ = v73
                }
            }
            for fi98 := 0; fi98 < 10; fi98++ {
                switch (("foo bar" + (v5 + ((v5 + v5) + (v5 + ("" + "hello"))))) + ((("\n\t" + "") + "foo bar") + (v5 + ("hello" + v5)))) {
                case "hello":
                    var v99 int8 = (-(v3))
                    _ = v99
                case "a":
                    if (!true) {
                        v100 := GenS4{F0: 42}
                        v3 = (^(v3))
                        if (3.14159 < (float64(float32(0.1)) / 0.1)) {
                            var v101 float64 = 1.5
                            _ = v101
                        } else {
                            if (0.1 <= 0.0) {
                                for fi102 := 0; fi102 < 8; fi102++ {
                                    v103 := GenS13{F0: 65535}
                                    _ = fi102
                                    _ = v103
                                }
                                for fi104 := 0; fi104 < 9; fi104++ {
                                    switch (fi98 % ((fi98) | 1)) {
                                    case -17:
                                        v5 = v5
                                    }
                                    var v105 uint32 = 2147483648
                                    _ = fi104
                                    _ = v105
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v106 int32 = 1
                            _ = v106
                        }
                        _ = v100
                    }
                default:
                    v2[(v5 + ((v5 + v5) + ""))] = (51 > 96)
                }
                v107 := GenS7{F0: 65535, F1: 412, F2: 2}
                for fi108 := 0; fi108 < 20; fi108++ {
                    var v109 string = v5
                    if ((!(255 >= 175)) && (true || (true && false))) {
                        var v110 int32 = 42
                        _ = v110
                    } else {
                        v111 := GenS14{F0: 4294967295, F1: (v4 & genFn0()), F2: 0}
                        delete(v2, (v109 + v5))
                        var v112 uint16 = 420
                        var v113 float64 = float64(0.0)
                        v3 = (v3 - v3)
                        _ = v111
                        _ = v112
                        _ = v113
                    }
                    _ = fi108
                    _ = v109
                }
                _ = fi98
                _ = v107
            }
        }
    }
    v5 = ("\n\t" + "a")
    v3 -= (v3 | (v3 | v3))
    var v114 int16 = 42
    v115 := GenS11{F0: -309, F1: (v5 + "hello")}
    for fi116 := 0; fi116 < 1; fi116++ {
        v4 = genFn0()
        var v117 int = fi116
        _ = fi116
        _ = v117
    }
    var v118 int16 = (v4 - 0)
    v119 := GenS15{F0: 255, F1: (v118 >> ((2) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v118:", v118)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v114
    _ = v115
    _ = v118
    _ = v119
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
