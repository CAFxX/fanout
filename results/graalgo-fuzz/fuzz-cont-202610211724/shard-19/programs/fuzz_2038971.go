package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint8; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int8; F1 int8; F2 uint32 }
type GenS8 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint; F1 int16 }
type GenS10 struct { F0 int; F1 int64; F2 int32 }
type GenS11 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int = 0
    var v2 int = 42
    v3 := make(map[int]uint32)
    var v4 bool = (-525 <= 0)
    for fi5 := 0; fi5 < 13; fi5++ {
        switch -272 {
        case 42:
            if v4 {
                delete(v3, (v2 | v1))
                v6 := GenS1{F0: 2}
                _ = v6
            } else {
                for fi7 := 0; fi7 < 11; fi7++ {
                    v1 = -246
                    v3[(fi5 - fi7)] = 1
                    _ = fi7
                }
            }
        case -9223372036854775808:
            v8 := []int{(-(fi5)), 1, (v2 & (fi5 ^ (v1 << ((46) & 63))))}
            _ = v8
        default:
            v4 = false
        }
        _ = fi5
    }
    v9 := GenS1{F0: 808}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 float32 = 1e10
        _ = fi10
        _ = v11
    }
    var v12 int32 = 968
    var v13 uint64 = 232
    switch int(v13) {
    case -1:
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: 43}
            if (3.14159 != 0.1) {
                var v16 float64 = 0.1
                _ = v16
            }
            _ = fi14
            _ = v15
        }
    case 42:
        if v4 {
            if v4 {
                var v17 int64 = 0
                _ = v17
            } else {
                var v18 int = 527
                _ = v18
            }
        }
    case 1:
        if false {
            v19 := GenS2{F0: 882}
            v20 := GenS3{F0: 65535, F1: 255, F2: (v12 - 1)}
            v21 := []int{v2, (v2 & v1), (v1 - v2), (-(v1))}
            _ = v19
            _ = v20
            _ = v21
        } else {
            fmt.Println(v12)
        }
    }
    if v4 {
        var v22 uint8 = 221
        if (v22 < v22) {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 uint8 = (v22 << ((374) & 63))
                var v25 int16 = -324
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 uint16 = 492
            if v4 {
                for fi27 := 0; fi27 < 2; fi27++ {
                    v28 := []int{}
                    _ = fi27
                    _ = v28
                }
                var v29 float32 = -2.25
                for fi30 := 0; fi30 < 14; fi30++ {
                    v31 := []int{fi30, (v2 * 32767), (v2 ^ fi30), v1}
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            _ = v26
        }
        _ = v22
    }
    var v32 uint32 = 65535
    delete(v3, v1)
    v33 := GenS2{F0: 42}
    v34 := GenS3{F0: 2, F1: 127, F2: 0}
    v1 = 42
    var v35 uint16 = 1
    var v36 string = (("" + "\n\t") + "foo bar")
    if true {
        for fi37 := 0; fi37 < 9; fi37++ {
            if (42 < 1000) {
                var v38 bool = (v32 <= v32)
                if (false && (42 == 106)) {
                    v39 := GenS2{F0: -1}
                    v40 := []int{}
                    _ = v39
                    _ = v40
                }
                _ = v38
            } else {
                var v41 string = (v36 + (("hello" + "hello") + "foo bar"))
                var v42 int8 = -121
                for fi43 := 0; fi43 < 17; fi43++ {
                    v44 := GenS3{F0: 2, F1: 146, F2: (v12 / ((v12) | 1))}
                    _ = fi43
                    _ = v44
                }
                _ = v41
                _ = v42
            }
            _ = fi37
        }
    } else {
        var v45 int64 = 1
        delete(v3, (v2 | v2))
        v46 := []int{v2, v2, -9223372036854775808, 32768}
        _ = v45
        _ = v46
    }
    if (("" >= ((("hello" + "hello") + (v36 + "hello")) + (v36 + v36))) && (((3.14159 * 1e10) * ((0.1 + ((-2.25 * 3.14159) / 1.5)) / -2.25)) >= 0.0)) {
        if v4 {
            var v47 uint = 4294967295
            if v4 {
                for fi48 := 0; fi48 < 1; fi48++ {
                    var v49 int64 = 0
                    switch v36 {
                    case "\n\t":
                        var v50 float32 = (((-2.25 - (((0.1 / 1.5) + ((0.0 + 1.5) * 0.0)) * 0.1)) - (0.1 / ((0.1 - -2.25) + -2.25))) + 3.14159)
                        _ = v50
                    case "a":
                        fmt.Println(v4)
                    case "hello":
                        var v51 int = (fi48 & int(v47))
                        _ = v51
                    default:
                        v52 := GenS1{F0: v47}
                        _ = v52
                    }
                    v1 = (-(fi48))
                    _ = fi48
                    _ = v49
                }
            }
            _ = v47
        } else {
            v3[int(v12)] = 120
            _ = len(v3)
            if (255 >= v35) {
                v53 := GenS2{F0: -1}
                _ = v53
            } else {
                switch ((v36 + "\n\t") + v36) {
                case "":
                    if true {
                        var v54 float32 = 0.1
                        if v4 {
                            var v55 int32 = -1
                            var v56 float64 = -2.25
                            _ = v55
                            _ = v56
                        } else {
                            for fi57 := 0; fi57 < 17; fi57++ {
                                if false {
                                    v58 := GenS4{F0: "foo bar"}
                                    v59 := GenS5{F0: 42}
                                    v60 := GenS6{F0: (v35 / ((v35) | 1)), F1: v32}
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                } else {
                                    var v61 string = (((v36 + v36) + "") + (v36 + v36))
                                    _ = v61
                                }
                                v62 := GenS2{F0: 847}
                                _ = fi57
                                _ = v62
                            }
                            v12 -= (-(v12))
                            v63 := GenS3{F0: 52, F1: 57, F2: v12}
                            _ = v63
                        }
                        fmt.Println(v32, v1)
                        _ = v54
                    }
                case "\n\t":
                    var v64 int = 0
                    _ = v64
                }
                switch "\n\t" {
                case "":
                    v13 = 217
                case "foo bar":
                    switch -1 {
                    case 127:
                        v13 -= v13
                    case -568:
                        for fi65 := 0; fi65 < 12; fi65++ {
                            if (-1 != 80) {
                                for fi66 := 0; fi66 < 8; fi66++ {
                                    if ((v32 | (v32 | (v32 - v32))) == (v32 >> ((2) & 63))) {
                                        if true {
                                            var v67 uint16 = v35
                                            _ = v67
                                        } else {
                                            switch 32768 {
                                            case 2147483647:
                                                v68 := GenS2{F0: 0}
                                                v1 = -1
                                                v3[32768] = v32
                                                _ = v68
                                            default:
                                                v12 = 42
                                            }
                                        }
                                    } else {
                                        var v69 string = v36
                                        if v4 {
                                            var v70 string = ""
                                            var v71 int32 = -2147483648
                                            var v72 float64 = float64(1e10)
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v69
                                    }
                                    if false {
                                        fmt.Println(v12)
                                        var v73 uint8 = 42
                                        var v74 int16 = int16(v35)
                                        _ = v73
                                        _ = v74
                                    } else {
                                        if true {
                                            var v75 string = v36
                                            if v4 {
                                                v12 = v12
                                                switch v1 {
                                                case 32768:
                                                    v76 := GenS3{F0: 255, F1: 62, F2: v12}
                                                    if v4 {
                                                        var v77 string = ""
                                                        v78 := GenS7{F0: -86, F1: 40, F2: v32}
                                                        if (!true) {
                                                            v12 = 151
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    switch v75 {
                                                    case "hello":
                                                        fmt.Println(v75, v32)
                                                    default:
                                                        var v79 int16 = int16(v35)
                                                        _ = v79
                                                    }
                                                    _ = v76
                                                case 2147483647:
                                                    v80 := GenS8{F0: (v32 % ((v32) | 1)), F1: 65535, F2: 42}
                                                    _ = v80
                                                }
                                                fmt.Println(v4, v2)
                                            } else {
                                                var v81 int8 = 1
                                                _ = v81
                                            }
                                            var v82 int = (v1 % ((int(v12)) | 1))
                                            _ = v75
                                            _ = v82
                                        }
                                        if true {
                                            var v83 int8 = 127
                                            fmt.Println(v4, v13)
                                            var v84 int16 = -215
                                            _ = v83
                                            _ = v84
                                        } else {
                                            if v4 {
                                                var v85 uint8 = 42
                                                _ = v85
                                            }
                                            var v86 string = v36
                                            _ = v86
                                        }
                                        fmt.Println(fi65, v12)
                                    }
                                    _ = fi66
                                }
                                for fi87 := 0; fi87 < 18; fi87++ {
                                    if (false && true) {
                                        if v4 {
                                            switch v36 {
                                            case "":
                                                switch v36 {
                                                case "a":
                                                    if true {
                                                        fmt.Println(fi65)
                                                        v88 := GenS8{F0: (v32 - (^(v32))), F1: 0, F2: v35}
                                                        _ = v88
                                                    }
                                                case "hello":
                                                    v89 := GenS3{F0: 229, F1: 1, F2: (v12 << ((899) & 63))}
                                                    v1 += int(v32)
                                                    v4 = (uint(v35) == 42)
                                                    var v90 int8 = 42
                                                    switch fi87 {
                                                    case 2147483647:
                                                        v13 -= 2
                                                    case -116:
                                                        fmt.Println(v4, v4, v90)
                                                    case 83:
                                                        var v91 int = 9223372036854775807
                                                        _ = v91
                                                    default:
                                                        var v92 int16 = 32767
                                                        _ = v92
                                                    }
                                                    v4 = v4
                                                    var v93 int32 = v12
                                                    v94 := GenS9{F0: 2, F1: 1}
                                                    _ = v89
                                                    _ = v90
                                                    _ = v93
                                                    _ = v94
                                                }
                                            }
                                            fmt.Println(v1, v35)
                                            var v95 int8 = -21
                                            _ = v95
                                        } else {
                                            _ = len(v3)
                                            var v96 int16 = 1
                                            v13 -= (v13 | 4294967295)
                                            _ = v96
                                        }
                                        switch (v2 - v1) {
                                        case -660:
                                            if false {
                                                v2 = -1
                                                if (v12 <= (v12 << ((810) & 63))) {
                                                    v36 = v36
                                                    var v97 uint16 = (v35 - (v35 | uint16(v13)))
                                                    v3[275] = 2147483648
                                                    _ = v97
                                                } else {
                                                    if (!("\n\t" < "\n\t")) {
                                                        if (!(!(16 <= 18))) {
                                                            if v4 {
                                                                v4 = ((((0.1 - 0.0) * ((((0.1 + 1.5) / 1e10) * (((1e10 / 0.1) - float64(-2.25)) * 1e10)) + 1.5)) <= (0.0 / float64(((-2.25 + 0.1) - (float32(0.1) / 1.5))))) && ((^(v13)) < (^(v13))))
                                                            }
                                                            v98 := []int{(v2 - (v2 % ((int(v35)) | 1))), 128}
                                                            var v99 int = v1
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            var v100 int = (fi87 % ((-9223372036854775808) | 1))
                                                            v100 = fi65
                                                            v101 := GenS1{F0: 18446744073709551615}
                                                            _ = v100
                                                            _ = v101
                                                        }
                                                        var v102 int = (v1 & (-(fi65)))
                                                        v36 = (v36 + (v36 + ""))
                                                        _ = v102
                                                    } else {
                                                        switch ((v36 + v36) + "a") {
                                                        case "a":
                                                            v32 -= (v32 + 2147483648)
                                                            v103 := GenS10{F0: fi65, F1: 42, F2: (^(v12))}
                                                            var v104 float32 = 0.1
                                                            _ = v103
                                                            _ = v104
                                                        default:
                                                            var v105 uint = 758
                                                            _ = v105
                                                        }
                                                        fmt.Println(v1, v4, fi65)
                                                    }
                                                }
                                                v2 = 2147483647
                                            } else {
                                                v106 := []int{-1, fi87, 42}
                                                v107 := GenS2{F0: 42}
                                                _ = v106
                                                _ = v107
                                            }
                                            _ = len(v3)
                                        case -61:
                                            switch (^(v1)) {
                                            case 2147483647:
                                                switch 0 {
                                                case 1:
                                                    v13 = 2
                                                default:
                                                    var v108 int16 = 1
                                                    _ = v108
                                                }
                                            case 32768:
                                                if v4 {
                                                    v13 += (v13 / ((18446744073709551615) | 1))
                                                    var v109 bool = false
                                                    _ = v109
                                                }
                                            }
                                        }
                                        v110 := GenS3{F0: uint(v32), F1: 1, F2: v12}
                                        _ = v110
                                    }
                                    _ = fi87
                                }
                                _ = len(v3)
                            } else {
                                if true {
                                    if v4 {
                                        v111 := []int{fi65}
                                        var v112 uint64 = uint64(v35)
                                        v113 := GenS2{F0: 42}
                                        _ = v111
                                        _ = v112
                                        _ = v113
                                    }
                                } else {
                                    for fi114 := 0; fi114 < 15; fi114++ {
                                        v36 = v36
                                        _ = fi114
                                    }
                                    fmt.Println(v13, v12)
                                    for fi115 := 0; fi115 < 7; fi115++ {
                                        v116 := GenS8{F0: v32, F1: 696, F2: (^(v35))}
                                        if (uint(v35) > 42) {
                                            if v4 {
                                                v117 := GenS3{F0: 65535, F1: 1, F2: (v12 - (-(v12)))}
                                                v2 -= int(v32)
                                                var v118 uint32 = (v32 & v32)
                                                _ = v117
                                                _ = v118
                                            } else {
                                                switch v36 {
                                                case "hello":
                                                    var v119 float64 = -2.25
                                                    _ = v119
                                                default:
                                                    v3[(v2 ^ v1)] = (v32 / ((v32) | 1))
                                                }
                                            }
                                        } else {
                                            v120 := GenS2{F0: -525}
                                            var v121 uint32 = (v32 * v32)
                                            var v122 float32 = -2.25
                                            v3[v2] = v32
                                            var v123 int8 = 59
                                            switch v36 {
                                            case "foo bar":
                                                v13 -= v13
                                            }
                                            fmt.Println(v12, v12)
                                            _ = v120
                                            _ = v121
                                            _ = v122
                                            _ = v123
                                        }
                                        v13 = (v13 ^ v13)
                                        _ = fi115
                                        _ = v116
                                    }
                                }
                            }
                            v124 := v3[(^(v2))]
                            _ = fi65
                            _ = v124
                        }
                    case 4294967295:
                        v35 = v35
                    default:
                        v12 = v12
                    }
                case "a":
                    if ((1.5 * (1.5 * (3.14159 * (float64(-2.25) - 0.1)))) != 0.1) {
                        fmt.Println(v12, v1, v2)
                    } else {
                        if true {
                            fmt.Println(v4, v2, v36)
                            switch (v36 + "") {
                            case "":
                                var v125 int64 = 9223372036854775807
                                _ = v125
                            }
                        } else {
                            v36 = v36
                        }
                        var v126 uint16 = uint16(v13)
                        v127 := GenS11{F0: 42, F1: 251}
                        _ = v126
                        _ = v127
                    }
                }
            }
        }
        var v128 int = -100
        v129 := GenS2{F0: 42}
        _ = v128
        _ = v129
    }
    var v130 int32 = int32(v2)
    delete(v3, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v130
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v9
}
