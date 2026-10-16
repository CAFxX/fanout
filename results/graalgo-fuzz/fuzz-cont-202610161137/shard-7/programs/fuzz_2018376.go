package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint; F2 bool }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 int32; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint32; F1 int16; F2 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint64; F1 uint64 }
type GenS12 struct { F0 uint32; F1 int; F2 uint }

func main() {
    var v1 uint64 = 42
    var v2 bool = true
    var v3 uint8 = 1
    v4 := make(map[uint32]int)
    var v5 float64 = 0.0
    var v6 string = "a"
    v5 = (1e10 / v5)
    if v2 {
        for fi7 := 0; fi7 < 18; fi7++ {
            switch v6 {
            case "\n\t":
                fmt.Println(v2, v2)
            }
            _ = fi7
        }
        var v8 int = -805
        _ = v8
    } else {
        fmt.Println(v5)
        v9 := GenS1{F0: "\n\t", F1: 1, F2: 4294967295}
        var v10 int8 = -128
        _ = v10
        _ = v9
    }
    var v11 string = ""
    v12 := GenS1{F0: "", F1: 127, F2: 4294967295}
    v13 := GenS1{F0: v11, F1: -736, F2: 273}
    var v14 int8 = -39
    v15 := GenS1{F0: (v6 + v6), F1: 651, F2: 105}
    var v16 int32 = 42
    var v17 uint16 = 110
    var v18 uint8 = (v3 & (v3 + (v3 % ((187) | 1))))
    if false {
        var v19 float64 = -2.25
        _ = len(v4)
        _ = v19
    }
    if (!true) {
        var v20 uint64 = v1
        _ = v20
    }
    switch (v11 + v11) {
    case "\n\t":
        fmt.Println(v11, v18)
    case "hello":
        v21 := []int{2147483647, -9223372036854775808}
        _ = v21
    case "foo bar":
        if (("a" + (((v11 + v11) + "\n\t") + v11)) != v11) {
            switch 42 {
            case -279:
                delete(v4, 42)
            }
        }
    default:
        v22 := v4[uint32(v17)]
        _ = v22
    }
    if v2 {
        switch 42 {
        case -585:
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(fi23, v6)
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 float64 = float64((-2.25 / 0.1))
                    v26 := []int{}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        case -768:
            v27 := GenS1{F0: (v6 + "\n\t"), F1: 962, F2: uint(v17)}
            _ = v27
        default:
            v6 = (v11 + v6)
        }
        switch 2147483647 {
        case 127:
            v2 = ((!false) && (v2 && false))
        case 1:
            if false {
                for fi28 := 0; fi28 < 4; fi28++ {
                    v18 += 59
                    v17 = (v17 * (v17 << ((0) & 63)))
                    var v29 float64 = (v5 + v5)
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS2{F0: 0, F1: 42, F2: (^(v1))}
                    switch (-(fi30)) {
                    case 490:
                        var v32 float64 = v5
                        _ = v32
                    case -9223372036854775808:
                        v17 = v17
                    case 71:
                        var v33 string = ((v6 + "foo bar") + v6)
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
                var v34 int32 = v16
                _ = v34
            } else {
                v35 := []int{32768}
                if false {
                    if len(v35) > 0 {
                        v35[0] = 42
                    }
                    v5 += v5
                }
                _ = v35
            }
        default:
            fmt.Println(v1, v11)
        }
    } else {
        for fi36 := 0; fi36 < 19; fi36++ {
            if (v17 <= 1) {
                for fi37 := 0; fi37 < 10; fi37++ {
                    var v38 uint16 = 32767
                    v39 := GenS3{F0: (v6 == (v6 + ((v11 + "foo bar") + v11)))}
                    v40 := GenS4{F0: (v1 << ((4294967295) & 63))}
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
            } else {
                delete(v4, 42)
                for fi41 := 0; fi41 < 16; fi41++ {
                    if (1 < (fi41 % ((fi36) | 1))) {
                        fmt.Println(fi36, v1, fi41)
                        v6 = (v11 + ("a" + v6))
                    }
                    var v42 float64 = float64(-2.25)
                    _ = fi41
                    _ = v42
                }
                var v43 int32 = v16
                _ = v43
            }
            var v44 bool = false
            switch (v6 + (v11 + v6)) {
            case "hello":
                fmt.Println(v6)
            case "":
                var v45 int = fi36
                _ = v45
            }
            _ = fi36
            _ = v44
        }
        for fi46 := 0; fi46 < 16; fi46++ {
            switch v11 {
            case "a":
                switch fi46 {
                case 9223372036854775807:
                    if (v1 != (v1 - 286)) {
                        var v47 int8 = 85
                        if false {
                            if v2 {
                                for fi48 := 0; fi48 < 18; fi48++ {
                                    var v49 uint64 = 2
                                    v50 := []int{(fi46 ^ fi48)}
                                    if v2 {
                                        _ = len(v50)
                                        if len(v50) >= 2 {
                                            v50 = v50[1:len(v50)]
                                        }
                                        if v2 {
                                            if (v2 && true) {
                                                fmt.Println(v3)
                                                fmt.Println(v14)
                                            }
                                            var v51 int = fi48
                                            _ = v51
                                        }
                                    } else {
                                        switch v6 {
                                        case "":
                                            switch -298 {
                                            case 289:
                                                fmt.Println(v3)
                                            default:
                                                var v52 uint64 = uint64(v17)
                                                _ = v52
                                            }
                                        case "foo bar":
                                            if v2 {
                                                fmt.Println(v5, v2)
                                                v53 := GenS5{F0: 127, F1: 18446744073709551615, F2: (-904 != -1)}
                                                _ = v53
                                            }
                                        default:
                                            v54 := GenS2{F0: 42, F1: 0, F2: 1}
                                            _ = v54
                                        }
                                        var v55 int8 = v14
                                        if ((42 > -1) || ((v18 - v3) == (^(v18)))) {
                                            var v56 uint8 = 255
                                            fmt.Println(fi46, fi46)
                                            _ = v56
                                        }
                                        if (0.1 <= float64(1.5)) {
                                            v18 += (^(v18))
                                        }
                                        var v57 int16 = -63
                                        _ = v55
                                        _ = v57
                                    }
                                    v2 = (0 != 1)
                                    var v58 uint8 = (v18 >> ((4294967295) & 63))
                                    v59 := GenS4{F0: 4294967295}
                                    if v2 {
                                        switch v6 {
                                        case "\n\t":
                                            var v60 float32 = 0.0
                                            _ = v60
                                        }
                                        var v61 int32 = (v16 | 2147483647)
                                        _ = v61
                                    } else {
                                        v62 := GenS3{F0: (65535 != 2)}
                                        _ = v62
                                    }
                                    switch v6 {
                                    case "foo bar":
                                        if true {
                                            var v63 uint8 = (v3 + uint8(v49))
                                            v64 := v4[576]
                                            v65 := v4[4294967295]
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                        if false {
                                            var v66 int8 = v47
                                            v66 = v66
                                            var v67 uint = 42
                                            var v68 bool = (!(v2 || (!true)))
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        } else {
                                            v3 = v18
                                        }
                                        if len(v50) > 0 {
                                            v50[0] = fi46
                                        }
                                        var v69 int16 = 1
                                        v70 := v4[0]
                                        _ = len(v4)
                                        _ = v69
                                        _ = v70
                                    case "":
                                        var v71 uint32 = 2
                                        _ = v71
                                    default:
                                        var v72 uint64 = v49
                                        _ = v72
                                    }
                                    v11 = "\n\t"
                                    var v73 float32 = float32(float64(0.1))
                                    var v74 int32 = v16
                                    _ = fi48
                                    _ = v49
                                    _ = v50
                                    _ = v58
                                    _ = v59
                                    _ = v73
                                    _ = v74
                                }
                            } else {
                                v4[255] = fi46
                            }
                        } else {
                            var v75 int8 = v14
                            _ = v75
                        }
                        _ = v47
                    }
                case -1:
                    v6 = v6
                case -9223372036854775808:
                    v18 -= 0
                default:
                    v76 := GenS6{F0: "foo bar", F1: (fi46 - int(v17))}
                    _ = v76
                }
            case "foo bar":
                if false {
                    var v77 float64 = v5
                    var v78 float32 = (0.0 / (float32((0.1 * float64(3.14159))) + float32(v77)))
                    _ = v77
                    _ = v78
                }
            case "hello":
                fmt.Println(v6, v14)
            default:
                if v2 {
                    var v79 uint32 = 706
                    switch "foo bar" {
                    case "foo bar":
                        if v2 {
                            switch v11 {
                            case "\n\t":
                                delete(v4, v79)
                            case "a":
                                if v2 {
                                    v80 := GenS2{F0: 2, F1: 1, F2: v1}
                                    for fi81 := 0; fi81 < 11; fi81++ {
                                        var v82 uint8 = (^(v3))
                                        v83 := []int{fi46, fi46, fi81, -795}
                                        var v84 int64 = int64(v14)
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                    _ = v80
                                }
                            default:
                                _ = len(v4)
                            }
                        } else {
                            var v85 float32 = 1.5
                            v86 := []int{758, fi46}
                            v3 = 64
                            _ = v85
                            _ = v86
                        }
                    default:
                        if (v6 <= v11) {
                            v87 := []int{}
                            _ = v87
                        }
                    }
                    v18 = 200
                    _ = v79
                } else {
                    for fi88 := 0; fi88 < 13; fi88++ {
                        var v89 int64 = 0
                        v89 = (v89 - v89)
                        fmt.Println(v6, v11)
                        _ = fi88
                        _ = v89
                    }
                }
            }
            _ = fi46
        }
        switch 0 {
        case -2147483648:
            if true {
                var v90 int = -750
                _ = v90
            } else {
                v91 := GenS6{F0: ("\n\t" + v6), F1: 127}
                v92 := GenS7{F0: 0, F1: 0}
                v93 := v4[198]
                _ = v91
                _ = v92
                _ = v93
            }
        case 9223372036854775807:
            switch v6 {
            case "\n\t":
                switch v11 {
                case "hello":
                    var v94 float64 = ((v5 - (v5 - (-2.25 + float64(float32(float64(-2.25)))))) + -2.25)
                    _ = v94
                case "foo bar":
                    switch v6 {
                    case "":
                        switch "" {
                        case "foo bar":
                            v95 := GenS8{F0: int32(v18), F1: (-(v14)), F2: 4294967295}
                            _ = v95
                        case "a":
                            switch 145 {
                            case 1:
                                v96 := GenS9{F0: 4294967295, F1: -32768, F2: v2}
                                _ = v96
                            case 291:
                                for fi97 := 0; fi97 < 6; fi97++ {
                                    if v2 {
                                        var v98 bool = ((v3 / ((v3) | 1)) == v18)
                                        _ = v98
                                    } else {
                                        for fi99 := 0; fi99 < 2; fi99++ {
                                            if v2 {
                                                if (!v2) {
                                                    fmt.Println(v18, v1, v16)
                                                } else {
                                                    fmt.Println(fi97, v2, v6)
                                                }
                                                v5 = (0.1 - v5)
                                                v100 := GenS9{F0: 42, F1: 0, F2: true}
                                                if true {
                                                    var v101 float32 = (0.0 + 1e10)
                                                    _ = v101
                                                } else {
                                                    if true {
                                                        v102 := GenS3{F0: (((1 > 2) && v2) && ((v17 % ((uint16(v1)) | 1)) > 1))}
                                                        fmt.Println(v3, v14, v14)
                                                        v103 := []int{}
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                }
                                                _ = v100
                                            }
                                            v104 := GenS3{F0: (v2 && false)}
                                            if (fi99 != (-(fi99))) {
                                                var v105 float32 = (((0.0 * 0.0) * 1e10) + 0.1)
                                                if (639 > 2147483647) {
                                                    var v106 uint32 = 3
                                                    v1 += 4294967295
                                                    _ = v106
                                                } else {
                                                    if false {
                                                        v107 := GenS10{F0: 65535}
                                                        _ = v107
                                                    } else {
                                                        var v108 uint64 = v1
                                                        var v109 int = (fi99 >> ((42) & 63))
                                                        v110 := []int{(-(v109)), v109}
                                                        v111 := GenS5{F0: 45, F1: 255, F2: ((^(v17)) == 0)}
                                                        var v112 uint = 289
                                                        _ = v108
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                        _ = v112
                                                    }
                                                }
                                                v105 = (float32(3.14159) / v105)
                                                var v113 string = v11
                                                _ = v105
                                                _ = v113
                                            }
                                            _ = fi99
                                            _ = v104
                                        }
                                        v114 := GenS10{F0: 267}
                                        _ = v114
                                    }
                                    if (v17 < v17) {
                                        v115 := v4[2147483648]
                                        fmt.Println(v3, v17, v3)
                                        _ = v115
                                    }
                                    var v116 bool = false
                                    _ = fi97
                                    _ = v116
                                }
                            case 127:
                                v3 += 56
                            default:
                                var v117 float32 = 0.1
                                _ = v117
                            }
                        case "\n\t":
                            v118 := GenS10{F0: 1}
                            _ = v118
                        default:
                            var v119 float32 = -2.25
                            _ = v119
                        }
                    case "foo bar":
                        switch -93 {
                        case 128:
                            switch int(v3) {
                            case -2147483648:
                                var v120 uint32 = 0
                                _ = v120
                            case -69:
                                var v121 float32 = (0.1 / float32(1e10))
                                _ = v121
                            case -1:
                                fmt.Println(v5, v14)
                            default:
                                var v122 uint16 = v17
                                _ = v122
                            }
                        case -2147483648:
                            v5 = (v5 / 1.5)
                        default:
                            for fi123 := 0; fi123 < 4; fi123++ {
                                if ((v3 == 206) || (!(!(!false)))) {
                                    v124 := GenS1{F0: ("a" + v11), F1: fi123, F2: uint(v18)}
                                    v3 -= v18
                                    _ = v124
                                }
                                for fi125 := 0; fi125 < 6; fi125++ {
                                    fmt.Println(fi125, v2)
                                    _ = fi125
                                }
                                if true {
                                    var v126 int64 = -483
                                    v4[0] = fi123
                                    var v127 float64 = 1.5
                                    _ = v126
                                    _ = v127
                                } else {
                                    v11 = v11
                                    v128 := GenS11{F0: 18446744073709551615, F1: 205}
                                    fmt.Println(fi123, v6, v2)
                                    _ = v128
                                }
                                _ = fi123
                            }
                        }
                    }
                case "":
                    switch v11 {
                    case "\n\t":
                        for fi129 := 0; fi129 < 11; fi129++ {
                            v11 = (v11 + "hello")
                            _ = fi129
                        }
                    case "":
                        v11 = ""
                    default:
                        switch "foo bar" {
                        case "":
                            for fi130 := 0; fi130 < 14; fi130++ {
                                var v131 float64 = 3.14159
                                _ = fi130
                                _ = v131
                            }
                        case "a":
                            switch -782 {
                            case 42:
                                switch "hello" {
                                case "":
                                    fmt.Println(v5, v16)
                                case "foo bar":
                                    v18 = (v18 | 172)
                                default:
                                    v132 := GenS11{F0: v1, F1: (v1 * v1)}
                                    _ = v132
                                }
                            case 128:
                                v14 = 42
                            case -559:
                                v133 := GenS12{F0: 2147483647, F1: -583, F2: 211}
                                _ = v133
                            default:
                                v134 := GenS2{F0: uint32(v17), F1: -32768, F2: 42}
                                _ = v134
                            }
                        }
                    }
                }
            default:
                v4[2] = 415
            }
        case -907:
            v11 = v6
        }
    }
    v135 := []int{-9223372036854775808, -881, -848}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v135:", len(v135), v135)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v135
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
