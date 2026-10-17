package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8; F1 int8; F2 uint16 }
type GenS10 struct { F0 int64; F1 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint; F1 uint32; F2 int }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 bool; F2 uint16 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint32; F1 int64 }
type GenS18 struct { F0 uint; F1 bool }
type GenS19 struct { F0 uint32; F1 int }
type GenS20 struct { F0 int32; F1 string }
type GenS21 struct { F0 uint; F1 uint }
type GenS22 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS23 struct { F0 string; F1 uint16 }
func genFn0(p0 string, p1 uint) int16 {
    var v4 int16 = 0
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int8 = 1
        switch p0 {
        case "hello":
            p1 = p1
        default:
            var v7 float32 = 0.1
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    var v8 int16 = v4
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return (v8 + v8)
}
func genFn1(p0 uint32, p1 uint64) uint64 {
    for fi9 := 0; fi9 < 10; fi9++ {
        if false {
            p0 += 0
            var v10 uint16 = 65535
            _ = v10
        } else {
            switch "a" {
            case "a":
                fmt.Println(fi9)
            case "hello":
                v11 := GenS1{F0: 796}
                _ = v11
            default:
                v12 := []int{-2147483648, (fi9 % (((fi9 | fi9)) | 1)), -397}
                _ = v12
            }
            var v13 uint64 = 0
            switch fi9 {
            case 32768:
                if (0 <= 2147483647) {
                    v14 := GenS1{F0: 1000}
                    var v15 float32 = 3.14159
                    _ = v14
                    _ = v15
                }
            default:
                var v16 bool = ((false && (!(128 < 1))) || true)
                _ = v16
            }
            _ = v13
        }
        var v17 uint = 278
        switch 9223372036854775807 {
        case 2147483647:
            var v18 int = (fi9 - fi9)
            _ = v18
        default:
            var v19 int16 = 1000
            _ = v19
        }
        _ = fi9
        _ = v17
    }
    fmt.Println(p1)
    fmt.Println(p1)
    p0 += (p0 / ((797) | 1))
    v20 := []int{}
    _ = p0
    _ = p1
    _ = v20
    return p1
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]bool)
    var v3 string = "hello"
    fmt.Println(v3, v1)
    var v21 int64 = (v1 + v1)
    v22 := GenS2{F0: (v21 & v1)}
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS3{F0: fi23, F1: 42, F2: 2}
        if false {
            if false {
                _ = len(v2)
                var v25 float32 = -2.25
                _ = v25
            }
            var v26 bool = false
            if v26 {
                v27 := GenS1{F0: -1}
                switch v3 {
                case "hello":
                    var v28 uint64 = 0
                    _ = v28
                case "\n\t":
                    v29 := GenS4{F0: 32767, F1: 685}
                    _ = v29
                case "":
                    var v30 int32 = 261
                    _ = v30
                default:
                    fmt.Println(v1, fi23, v21)
                }
                v26 = true
                _ = v27
            } else {
                switch fi23 {
                case 9223372036854775807:
                    fmt.Println(v1, v26, fi23)
                case 660:
                    v31 := GenS5{F0: v3}
                    _ = v31
                case 565:
                    if v26 {
                        v32 := GenS3{F0: fi23, F1: 0, F2: 1}
                        var v33 string = v3
                        switch (("foo bar" + v3) + v33) {
                        case "foo bar":
                            for fi34 := 0; fi34 < 15; fi34++ {
                                if false {
                                    fmt.Println(fi23)
                                    v33 = (v3 + "")
                                }
                                _ = fi34
                            }
                        default:
                            if v26 {
                                var v35 uint16 = 255
                                var v36 int32 = -2147483648
                                var v37 string = (((("\n\t" + ("\n\t" + (v3 + "hello"))) + v33) + ((v33 + v33) + (("a" + v3) + v3))) + ("a" + "a"))
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                v38 := v2[("a" + (v3 + "\n\t"))]
                                v39 := GenS4{F0: 177, F1: 746}
                                switch (fi23 + (fi23 * 763)) {
                                case 9223372036854775807:
                                    var v40 bool = v38
                                    _ = v40
                                case -499:
                                    if (100 < 127) {
                                        v41 := GenS5{F0: ("a" + (v33 + v33))}
                                        v38 = v26
                                        if v26 {
                                            if true {
                                                if (42 <= -1) {
                                                    switch fi23 {
                                                    case 9223372036854775807:
                                                        var v42 uint8 = 62
                                                        _ = v42
                                                    case -93:
                                                        for fi43 := 0; fi43 < 19; fi43++ {
                                                            if v38 {
                                                                var v44 float32 = float32((-2.25 + float64(0.1)))
                                                                v44 += float32((0.1 + (1.5 / -2.25)))
                                                                switch v3 {
                                                                case "a":
                                                                    var v45 string = v3
                                                                    _ = v45
                                                                case "foo bar":
                                                                    var v46 uint8 = 128
                                                                    _ = v46
                                                                case "":
                                                                    if (v26 || (18446744073709551615 == 2)) {
                                                                        var v47 uint64 = genFn1(171, 18446744073709551615)
                                                                        switch ((v3 + "a") + v33) {
                                                                        case "hello":
                                                                            v26 = v38
                                                                        }
                                                                        if (!v38) {
                                                                            var v48 int8 = -128
                                                                            v26 = ((v48 > v48) && (!(v38 && v26)))
                                                                            if v38 {
                                                                                if v26 {
                                                                                    v33 = ((v33 + "foo bar") + "hello")
                                                                                }
                                                                                v49 := GenS6{F0: (!(2147483647 == 42))}
                                                                                v48 = v48
                                                                                _ = v49
                                                                            } else {
                                                                                if (((v44 * (v44 + (3.14159 - v44))) / v44) <= (v44 / 1.0)) {
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        if false {
                                                                                            v50 := []int{}
                                                                                            _ = v50
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v51 uint8 = 68
                                                                                        _ = v51
                                                                                    case "foo bar":
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            v3 = "foo bar"
                                                                                        case "a":
                                                                                            v3 = v3
                                                                                        case "foo bar":
                                                                                            var v52 string = ("\n\t" + (v3 + ""))
                                                                                            _ = v52
                                                                                        default:
                                                                                            switch fi23 {
                                                                                            case -299:
                                                                                                var v53 uint64 = (v47 | 406)
                                                                                                _ = v53
                                                                                            case -859:
                                                                                                if v38 {
                                                                                                    v54 := []int{9223372036854775807, fi23, fi23, fi43}
                                                                                                    if len(v54) >= 2 {
                                                                                                        v54 = v54[1:len(v54)]
                                                                                                    }
                                                                                                    var v55 int8 = v48
                                                                                                    _ = v54
                                                                                                    _ = v55
                                                                                                } else {
                                                                                                    var v56 int16 = 220
                                                                                                    _ = v56
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                v47 -= (v47 + v47)
                                                                                            default:
                                                                                                v44 -= 0.0
                                                                                            }
                                                                                            if v38 {
                                                                                                v1 = int64(v48)
                                                                                                v48 = v48
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    v57 := GenS3{F0: (fi23 % ((873) | 1)), F1: int32(v48), F2: 1}
                                                                                                    _ = v57
                                                                                                case "\n\t":
                                                                                                    var v58 uint8 = 255
                                                                                                    _ = v58
                                                                                                case "":
                                                                                                    if v26 {
                                                                                                        if v38 {
                                                                                                            var v59 int32 = 2147483647
                                                                                                            v60 := GenS2{F0: (v21 % ((-1) | 1))}
                                                                                                            _ = v59
                                                                                                            _ = v60
                                                                                                        }
                                                                                                        var v61 uint16 = 32767
                                                                                                        delete(v2, ((v33 + v33) + "hello"))
                                                                                                        v62 := GenS5{F0: "a"}
                                                                                                        _ = v61
                                                                                                        _ = v62
                                                                                                    } else {
                                                                                                        v63 := []int{fi23, fi23, (-(fi23)), fi23}
                                                                                                        fmt.Println(fi43, v1, v38)
                                                                                                        v64 := GenS7{F0: -32768, F1: 0}
                                                                                                        _ = v63
                                                                                                        _ = v64
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v33 = v33
                                                                                    }
                                                                                    fmt.Println(fi23, v33, fi43)
                                                                                    v65 := GenS4{F0: 42, F1: 2147483647}
                                                                                    v66 := GenS2{F0: int64(v47)}
                                                                                    v67 := GenS3{F0: 127, F1: int32(v1), F2: 32767}
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                }
                                                                                fmt.Println(v47)
                                                                            }
                                                                            v68 := []int{fi23, fi23, (-(fi23))}
                                                                            delete(v2, ("" + "hello"))
                                                                            _ = v48
                                                                            _ = v68
                                                                        } else {
                                                                            fmt.Println(v38, v26)
                                                                        }
                                                                        _ = v47
                                                                    }
                                                                default:
                                                                    v3 = v33
                                                                }
                                                                _ = v44
                                                            } else {
                                                                var v69 int32 = 42
                                                                v33 = v33
                                                                _ = v69
                                                            }
                                                            v70 := GenS8{F0: 2}
                                                            _ = fi43
                                                            _ = v70
                                                        }
                                                    case 32767:
                                                        v71 := GenS9{F0: 100, F1: -1, F2: 134}
                                                        _ = v71
                                                    }
                                                    var v72 int = (fi23 ^ -1)
                                                    _ = v72
                                                } else {
                                                    if (-1 > 0) {
                                                        v73 := GenS8{F0: 1}
                                                        if (!v38) {
                                                            v26 = v26
                                                        } else {
                                                            if v26 {
                                                                switch (fi23 - fi23) {
                                                                case 9223372036854775807:
                                                                    var v74 int32 = -616
                                                                    _ = v74
                                                                case 127:
                                                                    v1 = (v21 ^ v21)
                                                                case -823:
                                                                    var v75 uint16 = 42
                                                                    _ = v75
                                                                }
                                                            }
                                                        }
                                                        v76 := GenS6{F0: (genFn1(948, genFn1(2147483648, 712)) != 707)}
                                                        _ = v73
                                                        _ = v76
                                                    }
                                                    if (!v38) {
                                                        v77 := GenS7{F0: 0, F1: 276}
                                                        var v78 string = ("\n\t" + "hello")
                                                        var v79 uint64 = 18446744073709551615
                                                        _ = v77
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    v26 = v38
                                                }
                                                if v26 {
                                                    v80 := GenS10{F0: -270, F1: 2}
                                                    for fi81 := 0; fi81 < 9; fi81++ {
                                                        v82 := GenS2{F0: (v1 << ((678) & 63))}
                                                        var v83 float64 = 0.1
                                                        v33 = "foo bar"
                                                        v38 = ((v38 || (!v38)) || (((v3 + (v3 + v33)) + (v33 + ((v3 + "foo bar") + v3))) > ("\n\t" + ("hello" + "hello"))))
                                                        v84 := GenS11{F0: 2}
                                                        _ = fi81
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    if (0.0 >= 0.1) {
                                                        fmt.Println(v3, v26, v1)
                                                        switch v3 {
                                                        case "":
                                                            if (false && v26) {
                                                                v3 = v33
                                                            }
                                                        case "\n\t":
                                                            for fi85 := 0; fi85 < 3; fi85++ {
                                                                if v38 {
                                                                    v86 := GenS12{F0: 4294967295, F1: 676, F2: (^(fi85))}
                                                                    _ = v86
                                                                }
                                                                var v87 float32 = 1.5
                                                                _ = fi85
                                                                _ = v87
                                                            }
                                                        }
                                                        v88 := []int{fi23, (fi23 & int(v1))}
                                                        _ = v88
                                                    } else {
                                                        switch (-(fi23)) {
                                                        case -48:
                                                            var v89 bool = (v38 || false)
                                                            _ = v89
                                                        default:
                                                            v3 = "\n\t"
                                                        }
                                                    }
                                                    _ = v80
                                                }
                                                var v90 uint64 = 18446744073709551615
                                                _ = v90
                                            }
                                            var v91 float64 = (3.14159 - 3.14159)
                                            _ = v91
                                        } else {
                                            v26 = (v33 <= ((v3 + v33) + v33))
                                            v3 = "a"
                                            var v92 float32 = (0.1 - (1.5 - (1e10 * float32((3.14159 + 1.5)))))
                                            _ = v92
                                        }
                                        _ = v41
                                    }
                                }
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = v32
                        _ = v33
                    } else {
                        v93 := GenS5{F0: ("hello" + v3)}
                        _ = v93
                    }
                }
            }
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    v94 := GenS11{F0: genFn1(2147483648, 322)}
    _ = len(v2)
    if true {
        for fi95 := 0; fi95 < 11; fi95++ {
            if (4294967295 == 309) {
                v96 := []int{(fi95 >> ((0) & 63)), (fi95 / ((4294967295) | 1))}
                for fi97 := 0; fi97 < 18; fi97++ {
                    switch v3 {
                    case "a":
                        v3 = "foo bar"
                    }
                    switch (v3 + (v3 + ("a" + v3))) {
                    case "hello":
                        var v98 string = v3
                        _ = v98
                    case "foo bar":
                        v3 = "hello"
                    }
                    v99 := GenS13{F0: 831}
                    _ = fi97
                    _ = v99
                }
                _ = v96
            }
            v1 -= v1
            _ = fi95
        }
        _ = len(v2)
        v100 := []int{-2147483648}
        _ = v100
    } else {
        for fi101 := 0; fi101 < 7; fi101++ {
            v21 = v21
            v102 := GenS10{F0: (^(v1)), F1: 128}
            v103 := GenS6{F0: (460 != 2147483647)}
            _ = fi101
            _ = v102
            _ = v103
        }
    }
    for fi104 := 0; fi104 < 12; fi104++ {
        if false {
            var v105 int8 = -128
            v106 := GenS14{F0: 0, F1: false, F2: 468}
            for fi107 := 0; fi107 < 10; fi107++ {
                switch v3 {
                case "\n\t":
                    switch (^(fi107)) {
                    case 123:
                        if (133 != 0) {
                            fmt.Println(v105)
                            switch v3 {
                            case "hello":
                                if (false || (true || true)) {
                                    v108 := []int{fi107, 128, fi104}
                                    _ = v108
                                } else {
                                    fmt.Println(fi107, v3)
                                    var v109 float32 = ((0.1 + ((-2.25 / 3.14159) / 3.14159)) / (1.5 * (-2.25 - (1.5 + 0.1))))
                                    _ = v109
                                }
                                switch fi107 {
                                case 0:
                                    v110 := GenS10{F0: v21, F1: 195}
                                    _ = v110
                                case 9223372036854775807:
                                    fmt.Println(v3, fi107)
                                case 110:
                                    fmt.Println(fi104, fi104)
                                default:
                                    if (-2.25 > -2.25) {
                                        if true {
                                            v111 := GenS9{F0: (v105 + (v105 | (v105 << ((18446744073709551615) & 63)))), F1: 100, F2: 255}
                                            fmt.Println(fi107)
                                            _ = v111
                                        } else {
                                            v3 = v3
                                            if (true || true) {
                                                var v112 float32 = 0.0
                                                _ = v112
                                            } else {
                                                fmt.Println(v105, v1, v1)
                                                v1 = v21
                                            }
                                            v21 = v21
                                        }
                                        v113 := []int{-432, (fi104 << ((42) & 63))}
                                        _ = v113
                                    } else {
                                        v114 := GenS15{F0: 2}
                                        v115 := []int{32768, (fi107 ^ (fi104 >> ((0) & 63))), (fi104 & fi104), (-(fi104))}
                                        switch fi107 {
                                        case 42:
                                            var v116 uint8 = 2
                                            _ = v116
                                        case 32768:
                                            fmt.Println(fi104, v105)
                                        }
                                        var v117 int16 = 239
                                        fmt.Println(v21, fi107, v105)
                                        _ = v114
                                        _ = v115
                                        _ = v117
                                    }
                                }
                            case "\n\t":
                                if (false && false) {
                                    v1 = v21
                                    var v118 int16 = 32767
                                    v119 := GenS16{F0: (v118 * v118)}
                                    _ = v118
                                    _ = v119
                                }
                            case "foo bar":
                                var v120 string = ""
                                _ = v120
                            }
                            v21 = v21
                        } else {
                            switch fi104 {
                            case 543:
                                v1 += (v1 << ((18446744073709551615) & 63))
                            case 129:
                                v121 := GenS17{F0: 2147483648, F1: (v21 * (^(v1)))}
                                _ = v121
                            case 42:
                                v105 = v105
                            default:
                                v21 = v1
                            }
                            v122 := []int{}
                            if (-730 < -330) {
                                if len(v122) > 0 {
                                    v122[0] = (fi104 ^ (fi107 ^ 32768))
                                }
                            }
                            _ = v122
                        }
                    case 1:
                        var v123 int8 = v105
                        _ = v123
                    case -531:
                        if (fi107 >= (fi104 / ((fi104) | 1))) {
                            if (-2147483648 > -1) {
                                v124 := GenS18{F0: 42, F1: true}
                                if false {
                                    v105 -= (v105 >> ((4294967295) & 63))
                                    fmt.Println(v105, v21)
                                }
                                if false {
                                    v125 := GenS9{F0: (^(v105)), F1: (v105 ^ -128), F2: 255}
                                    _ = v125
                                } else {
                                    v126 := []int{}
                                    if (((1.5 / float32(-2.25)) / 1.0) < float32(float64(float32(-2.25)))) {
                                        v127 := GenS19{F0: 2, F1: (fi104 >> ((1) & 63))}
                                        v128 := GenS20{F0: 2147483647, F1: ""}
                                        fmt.Println(v1, fi104, v3)
                                        _ = v127
                                        _ = v128
                                    } else {
                                        if len(v126) >= 2 {
                                            v126 = v126[1:len(v126)]
                                        }
                                        var v129 uint = 1
                                        var v130 uint16 = 470
                                        v1 = (v21 | v21)
                                        _ = v129
                                        _ = v130
                                    }
                                    if len(v126) >= 2 {
                                        v126 = v126[1:len(v126)]
                                    }
                                    _ = v126
                                }
                                _ = v124
                            } else {
                                fmt.Println(v3)
                            }
                            v21 -= 9223372036854775807
                            v21 = -1
                            switch v3 {
                            case "":
                                switch (fi104 | (fi104 % ((-438) | 1))) {
                                case 0:
                                    delete(v2, "\n\t")
                                    v131 := GenS15{F0: 65535}
                                    _ = v131
                                case 32768:
                                    if true {
                                        var v132 int = fi107
                                        v21 = 440
                                        var v133 int32 = 42
                                        if (539 != 42) {
                                            v133 = -1
                                        }
                                        v134 := GenS21{F0: 1, F1: 1}
                                        var v135 float64 = float64(1e10)
                                        _ = v132
                                        _ = v133
                                        _ = v134
                                        _ = v135
                                    }
                                    var v136 int8 = v105
                                    v105 = v136
                                    _ = v136
                                default:
                                    var v137 uint = 18446744073709551615
                                    _ = v137
                                }
                            }
                            delete(v2, v3)
                            var v138 int32 = 834
                            var v139 uint = 42
                            switch (-(fi104)) {
                            case 2147483647:
                                v140 := v2[v3]
                                _ = v140
                            case 4294967295:
                                fmt.Println(fi107, v138)
                            case 128:
                                v139 += (v139 / ((v139) | 1))
                            }
                            _ = v138
                            _ = v139
                        }
                    default:
                        var v141 bool = (4294967295 == 630)
                        _ = v141
                    }
                case "":
                    var v142 float64 = ((((-2.25 / ((0.0 / ((0.0 - float64(-2.25)) - -2.25)) + 3.14159)) * -2.25) * float64(0.0)) * 0.1)
                    _ = v142
                default:
                    v21 = (v21 ^ v1)
                }
                _ = fi107
            }
            _ = v105
            _ = v106
        } else {
            if false {
                if ((v3 + v3) >= v3) {
                    var v143 float32 = ((0.0 - 3.14159) * (0.1 * 1.5))
                    _ = v143
                }
                v144 := GenS22{F0: 255, F1: 232, F2: 2}
                v145 := GenS13{F0: 0}
                _ = v144
                _ = v145
            } else {
                for fi146 := 0; fi146 < 12; fi146++ {
                    var v147 int8 = -60
                    v2[v3] = false
                    if (true && (!(457 > -32768))) {
                        v148 := GenS23{F0: v3, F1: 217}
                        _ = v148
                    } else {
                        var v149 bool = (0 != fi146)
                        var v150 string = v3
                        _ = v149
                        _ = v150
                    }
                    _ = fi146
                    _ = v147
                }
                v151 := []int{(fi104 / ((fi104) | 1)), fi104, (fi104 % (((^(fi104))) | 1))}
                _ = v151
            }
            fmt.Println(v3, v21)
            v152 := GenS7{F0: 1000, F1: 4294967295}
            _ = v152
        }
        _ = fi104
    }
    switch 114 {
    case -745:
        var v153 uint = 1
        _ = v153
    case 1:
        fmt.Println(v1, v1, v3)
    default:
        for fi154 := 0; fi154 < 15; fi154++ {
            switch fi154 {
            case 684:
                if true {
                    v1 -= (v1 >> ((1) & 63))
                    var v155 bool = (513 < 178)
                    var v156 float64 = 1.5
                    _ = v155
                    _ = v156
                }
            case 9223372036854775807:
                var v157 uint16 = 2
                _ = v157
            default:
                var v158 int64 = 1
                _ = v158
            }
            v2["a"] = false
            fmt.Println(fi154, fi154, v1)
            _ = fi154
        }
    }
    v3 = "hello"
    if (120 > -128) {
        if (757 != 42) {
            v1 = v21
            var v159 int64 = (v21 | (v1 ^ (v21 >> ((0) & 63))))
            _ = v159
        }
        v21 = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v94
}
