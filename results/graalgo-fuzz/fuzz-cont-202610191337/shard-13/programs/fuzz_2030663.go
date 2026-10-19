package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 uint; F2 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int16; F1 bool; F2 int8 }
type GenS12 struct { F0 uint; F1 int; F2 uint16 }
type GenS13 struct { F0 bool; F1 uint; F2 uint64 }
type GenS14 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS15 struct { F0 bool; F1 int16; F2 int8 }
type GenS16 struct { F0 int; F1 int; F2 uint8 }
type GenS17 struct { F0 int64; F1 uint32 }
type GenS18 struct { F0 int16; F1 uint64 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint32; F1 uint64 }
func genFn0() uint16 {
    if false {
        v5 := GenS1{F0: 561, F1: 2147483647, F2: -797}
        var v6 int64 = 0
        v7 := []int{}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 32767, F1: 42}
    _ = v8
    return 65535
}
func genFn1(p0 uint64, p1 int) string {
    var v9 bool = false
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 string = "\n\t"
        var v12 int16 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        p0 -= 4294967295
        _ = fi13
    }
    switch 32768 {
    case 4294967295:
        v14 := []int{}
        _ = v14
    default:
        if (-9223372036854775808 > 42) {
            v15 := GenS1{F0: 2147483647, F1: 2147483647, F2: (p1 << ((232) & 63))}
            v9 = (v9 && v9)
            fmt.Println(p0, p0)
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int32 = 0
        p1 = (^(fi16))
        v17 = (v17 % ((0) | 1))
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 uint64 = 42
    var v4 int16 = 765
    var v18 string = ("a" + "\n\t")
    var v19 uint32 = 2147483647
    v20 := []int{-660, -483, 42}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    var v21 int32 = 1
    switch "foo bar" {
    case "foo bar":
        v4 = (v4 << ((18446744073709551615) & 63))
    default:
        for fi22 := 0; fi22 < 13; fi22++ {
            if false {
                switch (fi22 / ((fi22) | 1)) {
                case -2147483648:
                    v1[(^(v19))] = ("hello" + "a")
                default:
                    if (false || (!false)) {
                        v19 = v19
                        if false {
                            for fi23 := 0; fi23 < 15; fi23++ {
                                var v24 uint64 = 1
                                if (false || true) {
                                    if true {
                                        switch ((v18 + v18) + "") {
                                        case "foo bar":
                                            if true {
                                                fmt.Println(v3, v3)
                                            }
                                        default:
                                            var v25 int8 = -47
                                            _ = v25
                                        }
                                    } else {
                                        if (18446744073709551615 <= v3) {
                                            v26 := GenS3{F0: 18446744073709551615}
                                            switch v18 {
                                            case "":
                                                v27 := GenS2{F0: (v4 >> ((352) & 63)), F1: 852}
                                                _ = v27
                                            }
                                            var v28 int16 = v4
                                            switch (-(fi22)) {
                                            case 1:
                                                v29 := GenS4{F0: 96, F1: -1, F2: 1}
                                                _ = v29
                                            case 2147483647:
                                                var v30 int = fi22
                                                _ = v30
                                            default:
                                                fmt.Println(v3, v4)
                                            }
                                            v21 = (^(v21))
                                            var v31 uint8 = 2
                                            switch 127 {
                                            case 519:
                                                delete(v2, "")
                                            case 42:
                                                if ((^(v28)) < -1) {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v2[("a" + "hello")] = (("hello" + v18) + ((("hello" + "hello") + (v18 + v18)) + v18))
                                                    default:
                                                        fmt.Println(v3)
                                                    }
                                                } else {
                                                    v32 := GenS2{F0: (^(v28)), F1: int16(v31)}
                                                    if true {
                                                        if true {
                                                            v31 = (v31 % ((v31) | 1))
                                                            v31 += v31
                                                            fmt.Println(fi23)
                                                        }
                                                        v33 := GenS5{F0: 548, F1: 2}
                                                        var v34 uint64 = (v24 >> ((18446744073709551615) & 63))
                                                        fmt.Println(v34, v21)
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        if true {
                                                            var v35 bool = true
                                                            v20 = append(v20, 2147483647)
                                                            _ = v35
                                                        } else {
                                                            if false {
                                                                if true {
                                                                    if true {
                                                                        fmt.Println(v19, v21, v4)
                                                                        v28 = v4
                                                                        v36 := GenS1{F0: (v19 | v19), F1: -2147483648, F2: fi22}
                                                                        var v37 uint = 272
                                                                        fmt.Println(v3, fi22)
                                                                        _ = v36
                                                                        _ = v37
                                                                    } else {
                                                                        v28 = v4
                                                                    }
                                                                } else {
                                                                    if (!(true && true)) {
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                        v19 = v19
                                                                        v4 = (-(v28))
                                                                    } else {
                                                                        var v38 float64 = 3.14159
                                                                        switch fi22 {
                                                                        case -518:
                                                                            v39 := v1[v19]
                                                                            _ = v39
                                                                        case -2147483648:
                                                                            var v40 bool = (false && false)
                                                                            _ = v40
                                                                        }
                                                                        v41 := GenS6{F0: 904}
                                                                        _ = v38
                                                                        _ = v41
                                                                    }
                                                                    _ = len(v2)
                                                                    v2[v18] = v18
                                                                }
                                                                v18 = v18
                                                                v42 := GenS7{F0: 255, F1: 42, F2: 65535}
                                                                _ = v42
                                                            } else {
                                                                switch (fi23 / (((fi23 | fi23)) | 1)) {
                                                                case -660:
                                                                    v28 = v4
                                                                case 127:
                                                                    var v43 int8 = int8(fi23)
                                                                    _ = v43
                                                                case 0:
                                                                    if true {
                                                                        v44 := GenS8{F0: genFn0()}
                                                                        v45 := GenS9{F0: -9223372036854775808, F1: 255, F2: v19}
                                                                        _ = v44
                                                                        _ = v45
                                                                    } else {
                                                                        var v46 int64 = -1
                                                                        var v47 int64 = v46
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                default:
                                                                    var v48 uint = 880
                                                                    _ = v48
                                                                }
                                                                fmt.Println(v19, v19)
                                                                v49 := v2[v18]
                                                                _ = v49
                                                            }
                                                            v4 += (v4 ^ v4)
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                            case 718:
                                                fmt.Println(fi23, v31)
                                            default:
                                                switch v18 {
                                                case "foo bar":
                                                    fmt.Println(v28, v28, v18)
                                                case "\n\t":
                                                    v4 -= (v28 | (v4 * (-(v4))))
                                                }
                                            }
                                            if (!((!(104 == 0)) || (!(0.1 <= 0.1)))) {
                                                var v50 int32 = -793
                                                v3 = 405
                                                _ = v50
                                            } else {
                                                var v51 uint8 = 1
                                                if (!true) {
                                                    if ((v4 & (v4 % (((v4 << ((255) & 63))) | 1))) >= v28) {
                                                        v52 := GenS10{F0: v21}
                                                        _ = v52
                                                    }
                                                    var v53 uint = 65535
                                                    var v54 uint8 = 0
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    var v55 uint = 18446744073709551615
                                                    var v56 uint32 = (v19 | v19)
                                                    v28 = (v4 * v28)
                                                    _ = v55
                                                    _ = v56
                                                }
                                                v31 = 0
                                                _ = v51
                                            }
                                            _ = v26
                                            _ = v28
                                            _ = v31
                                        } else {
                                            if true {
                                                if false {
                                                    v21 = (v21 & int32(v19))
                                                    v57 := GenS3{F0: 255}
                                                    _ = v57
                                                }
                                            } else {
                                                fmt.Println(fi22)
                                                if (42 >= 42) {
                                                    if false {
                                                        var v58 float64 = 0.1
                                                        v20 = append(v20, 467)
                                                        var v59 float64 = (1e10 / float64(0.1))
                                                        v60 := GenS11{F0: -1, F1: (1000 <= v4), F2: 127}
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        v4 -= (v4 + (v4 >> ((4294967295) & 63)))
                                                    }
                                                    switch v18 {
                                                    case "":
                                                        v61 := GenS12{F0: 0, F1: 1, F2: 65535}
                                                        _ = v61
                                                    case "\n\t":
                                                        v62 := GenS1{F0: 2147483647, F1: (v21 ^ v21), F2: int(v24)}
                                                        _ = v62
                                                    case "foo bar":
                                                        var v63 string = ("hello" + v18)
                                                        _ = v63
                                                    default:
                                                        fmt.Println(v18)
                                                    }
                                                    delete(v1, (v19 / (((v19 - v19)) | 1)))
                                                }
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                if true {
                                                    v21 -= v21
                                                    var v64 int16 = int16(v21)
                                                    _ = v64
                                                }
                                            case "hello":
                                                v4 -= 781
                                            case "\n\t":
                                                v65 := GenS13{F0: true, F1: 4294967295, F2: (v24 & (v24 >> ((993) & 63)))}
                                                _ = v65
                                            }
                                            v3 += (v3 / (((v3 ^ v24)) | 1))
                                        }
                                        var v66 bool = ("" >= (v18 + v18))
                                        _ = v66
                                    }
                                    v21 = v21
                                    var v67 uint64 = v24
                                    _ = v67
                                }
                                switch (fi22 & 579) {
                                case -9223372036854775808:
                                    v68 := GenS9{F0: 673, F1: uint(v3), F2: v19}
                                    _ = v68
                                case 1:
                                    v4 = (v4 | 603)
                                }
                                _ = fi23
                                _ = v24
                            }
                        }
                        v18 = v18
                    } else {
                        var v69 int64 = 0
                        _ = v69
                    }
                }
                var v70 float64 = 1e10
                _ = v70
            }
            _ = fi22
        }
    }
    if (!true) {
        var v71 float32 = 1.5
        _ = v71
    } else {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        v72 := GenS14{F0: -1, F1: 128, F2: (v19 + 2)}
        _ = v72
    }
    v73 := GenS6{F0: 255}
    var v74 int8 = int8(v4)
    if ((v18 + v18) <= v18) {
        var v75 int16 = 1
        if true {
            for fi76 := 0; fi76 < 18; fi76++ {
                if (74 <= 127) {
                    v77 := GenS9{F0: 0, F1: 42, F2: v19}
                    for fi78 := 0; fi78 < 9; fi78++ {
                        if (((^(v74)) >= 61) || false) {
                            if (248 > 914) {
                                switch fi76 {
                                case -9223372036854775808:
                                    switch (("a" + v18) + v18) {
                                    case "hello":
                                        if (2 < 255) {
                                            fmt.Println(v18, fi76, v18)
                                        } else {
                                            switch fi78 {
                                            case 4294967295:
                                                if false {
                                                    if true {
                                                        var v79 bool = (v4 > 42)
                                                        var v80 uint8 = uint8(v19)
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        fmt.Println(v19)
                                                        fmt.Println(v19, fi76, fi78)
                                                        if (65535 < 1) {
                                                            fmt.Println(v18, fi78, v75)
                                                        }
                                                        v18 = ("a" + "")
                                                        fmt.Println(fi76, fi76, fi78)
                                                    }
                                                }
                                            case -1:
                                                v81 := GenS6{F0: 65535}
                                                _ = v81
                                            }
                                        }
                                    case "\n\t":
                                        v3 += v3
                                    }
                                default:
                                    v18 = ""
                                }
                                var v82 uint64 = 250
                                _ = v82
                            } else {
                                switch fi78 {
                                case 9223372036854775807:
                                    v83 := GenS15{F0: false, F1: (v4 >> ((733) & 63)), F2: 1}
                                    _ = v83
                                case 1:
                                    var v84 float32 = 1e10
                                    _ = v84
                                default:
                                    if len(v20) > 0 {
                                        v20[0] = fi76
                                    }
                                }
                                v21 = 2147483647
                            }
                            var v85 uint8 = 0
                            if (!((true || true) || false)) {
                                var v86 uint = 0
                                if (!(false || false)) {
                                    v86 -= (v86 << (((v86 / (((v86 & v86)) | 1))) & 63))
                                    v19 += 2
                                    switch ("" + v18) {
                                    case "hello":
                                        var v87 float32 = (float32(3.14159) * (((float32(0.1) * 3.14159) * (1e10 * 1e10)) - 3.14159))
                                        _ = v87
                                    case "a":
                                        var v88 uint32 = 4294967295
                                        _ = v88
                                    }
                                } else {
                                    v89 := GenS6{F0: (v86 % (((^(v86))) | 1))}
                                    if false {
                                        v86 = (v86 - v86)
                                    } else {
                                        v18 = ("a" + "a")
                                        delete(v2, v18)
                                    }
                                    _ = v89
                                }
                                _ = v86
                            } else {
                                v20 = append(v20, (^(fi78)))
                            }
                            v18 = ""
                            _ = v85
                        } else {
                            if ("a" == v18) {
                                v90 := GenS16{F0: (fi78 / ((-428) | 1)), F1: (fi76 << ((677) & 63)), F2: 143}
                                _ = v90
                            }
                            v91 := GenS2{F0: v75, F1: -1}
                            v92 := GenS5{F0: uint(v3), F1: 255}
                            v93 := GenS17{F0: 9223372036854775807, F1: 691}
                            if (!false) {
                                v94 := GenS9{F0: 9223372036854775807, F1: 763, F2: 2147483647}
                                _ = v94
                            } else {
                                var v95 uint16 = 65535
                                v96 := GenS15{F0: false, F1: v4, F2: v74}
                                v97 := GenS17{F0: -1, F1: 2147483648}
                                _ = v95
                                _ = v96
                                _ = v97
                            }
                            _ = v91
                            _ = v92
                            _ = v93
                        }
                        v18 = (("" + "\n\t") + v18)
                        var v98 uint8 = 2
                        var v99 uint32 = 1
                        v100 := GenS11{F0: v75, F1: (!(false && false)), F2: (^(v74))}
                        var v101 uint32 = (^(v99))
                        fmt.Println(v98, v21)
                        _ = fi78
                        _ = v100
                        _ = v101
                        _ = v98
                        _ = v99
                    }
                    _ = v77
                }
                _ = fi76
            }
            v102 := GenS18{F0: v4, F1: v3}
            _ = v102
        }
        v18 = ((v18 + (v18 + ("hello" + ("" + "a")))) + (((("a" + v18) + ("\n\t" + (v18 + (("a" + v18) + v18)))) + (v18 + ((v18 + v18) + v18))) + "\n\t"))
        _ = v75
    } else {
        for fi103 := 0; fi103 < 15; fi103++ {
            v74 = 100
            if false {
                switch (("\n\t" + v18) + "a") {
                case "a":
                    if false {
                        for fi104 := 0; fi104 < 20; fi104++ {
                            if (v74 == v74) {
                                fmt.Println(v19)
                                var v105 uint64 = v3
                                if (!(!(true || (42 >= 18446744073709551615)))) {
                                    _ = len(v20)
                                } else {
                                    switch (v18 + ("foo bar" + (v18 + v18))) {
                                    case "foo bar":
                                        var v106 uint32 = (v19 - 0)
                                        _ = v106
                                    case "\n\t":
                                        var v107 int8 = 42
                                        _ = v107
                                    default:
                                        v2[v18] = v18
                                    }
                                }
                                _ = v105
                            }
                            _ = fi104
                        }
                        var v108 uint = 65535
                        _ = v108
                    } else {
                        v21 = (-(v21))
                        for fi109 := 0; fi109 < 8; fi109++ {
                            if true {
                                fmt.Println(v3)
                                if (!(2147483647 >= -2147483648)) {
                                    var v110 float32 = float32((1e10 / 3.14159))
                                    v111 := GenS19{F0: 42}
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    v21 = v21
                                    _ = v110
                                    _ = v111
                                } else {
                                    var v112 int64 = -1
                                    _ = v112
                                }
                            } else {
                                var v113 float64 = float64(1.5)
                                switch 509 {
                                case 32767:
                                    var v114 string = "\n\t"
                                    _ = v114
                                }
                                _ = v113
                            }
                            fmt.Println(v74, v18)
                            v19 = 809
                            _ = fi109
                        }
                    }
                }
                v74 -= (v74 >> ((586) & 63))
            }
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi103
        }
        var v115 float32 = 1.5
        _ = v115
    }
    _ = len(v1)
    fmt.Println(v21, v18, v3)
    var v116 string = (v18 + "foo bar")
    fmt.Println(v21, v116)
    v3 = (^(v3))
    v20 = append(v20, 42)
    if ((2 > 255) || false) {
        for fi117 := 0; fi117 < 4; fi117++ {
            var v118 uint = 255
            _ = len(v1)
            _ = fi117
            _ = v118
        }
        if false {
            if (v3 <= 902) {
                v119 := GenS20{F0: v19, F1: 42}
                _ = v119
            } else {
                v120 := v1[2]
                for fi121 := 0; fi121 < 11; fi121++ {
                    var v122 string = v18
                    _ = len(v1)
                    v123 := v1[2]
                    _ = fi121
                    _ = v122
                    _ = v123
                }
                fmt.Println(v3)
                _ = v120
            }
        }
    } else {
        switch 0 {
        case 128:
            var v124 int16 = (v4 * (v4 >> ((7) & 63)))
            _ = v124
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v116:", v116)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v116
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v73
    _ = v74
}
