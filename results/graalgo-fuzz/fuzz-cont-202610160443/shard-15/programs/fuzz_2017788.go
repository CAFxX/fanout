package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 string; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 int16; F2 int16 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 int8; F1 string; F2 string }
type GenS11 struct { F0 int64; F1 bool }
func genFn0(p0 int32, p1 string) uint16 {
    v7 := []int{}
    if true {
        var v8 uint = 346
        v7 = append(v7, 857)
        p1 = p1
        _ = v8
    } else {
        if true {
            var v9 int = 127
            p1 = p1
            switch (v9 ^ 395) {
            case 1:
                switch v9 {
                case 32767:
                    v10 := GenS1{F0: 1}
                    _ = v10
                }
            case 105:
                for fi11 := 0; fi11 < 16; fi11++ {
                    if (65535 >= 2) {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            var v13 int16 = -1
                            _ = fi12
                            _ = v13
                        }
                        v9 = fi11
                        p0 = (p0 >> ((2) & 63))
                    } else {
                        var v14 float64 = (-2.25 * -2.25)
                        v15 := GenS1{F0: 1}
                        _ = v14
                        _ = v15
                    }
                    _ = fi11
                }
            case 498:
                v16 := GenS2{F0: p1}
                _ = v16
            }
            _ = v9
        } else {
            var v17 int8 = 127
            _ = v17
        }
    }
    var v18 bool = (false || (511 <= 1))
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 float64 = 3.14159
        _ = fi19
        _ = v20
    }
    var v21 uint32 = 716
    _ = p0
    _ = p1
    _ = v18
    _ = v21
    _ = v7
    return 555
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]string)
    var v5 uint = 227
    v6 := make(map[uint32]int)
    v22 := GenS1{F0: -9223372036854775808}
    v23 := v2[2147483647]
    v3 = 1
    var v24 int16 = 1
    fmt.Println(v24, v3)
    if v23 {
        var v25 uint16 = 65535
        v26 := []int{-551, 9223372036854775807}
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 42}
    var v28 uint16 = 42
    v29 := []int{}
    v1 = (v1 + v1)
    var v30 int64 = -9223372036854775808
    var v31 int16 = (v24 >> (((v5 / (((^(v5))) | 1))) & 63))
    for fi32 := 0; fi32 < 4; fi32++ {
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS2{F0: v1}
            _ = fi33
            _ = v34
        }
        if (1 == 98) {
            for fi35 := 0; fi35 < 1; fi35++ {
                v36 := GenS1{F0: (v30 + 790)}
                var v37 uint64 = 4294967295
                fmt.Println(v37, v24)
                _ = fi35
                _ = v36
                _ = v37
            }
            switch fi32 {
            case 127:
                fmt.Println(v31, v1)
            case 32767:
                var v38 int = (fi32 >> (((v5 * v5)) & 63))
                _ = v38
            default:
                if v23 {
                    switch v1 {
                    case "a":
                        v39 := GenS3{F0: (fi32 & (fi32 >> ((v5) & 63))), F1: -732}
                        _ = v39
                    case "hello":
                        v40 := GenS4{F0: "hello", F1: false, F2: (v28 ^ (v28 | v28))}
                        _ = v40
                    case "foo bar":
                        var v41 uint8 = 255
                        _ = v41
                    }
                    var v42 int64 = int64(v3)
                    _ = v42
                }
            }
            _ = len(v29)
        } else {
            _ = len(v29)
            if (!v23) {
                if (false && v23) {
                    v31 += v31
                    for fi43 := 0; fi43 < 10; fi43++ {
                        if (2 != 42) {
                            var v44 int32 = -889
                            v45 := GenS2{F0: ((v1 + v1) + (v1 + v1))}
                            var v46 int8 = -40
                            v31 += (v24 | (-(v31)))
                            if v23 {
                                v6[2147483647] = fi32
                                v29 = append(v29, fi43)
                                v2[4294967295] = true
                            }
                            if v23 {
                                switch (-(fi43)) {
                                case -1:
                                    v2[2] = (1e10 <= (float64(0.1) + float64(float32(-2.25))))
                                case -765:
                                    v28 = uint16(v3)
                                case -9223372036854775808:
                                    v47 := GenS5{F0: (^(v28)), F1: v28}
                                    _ = v47
                                }
                                v1 = ((((v1 + (v1 + v1)) + "foo bar") + "foo bar") + ("\n\t" + ((v1 + v1) + v1)))
                                var v48 int8 = int8(v3)
                                var v49 uint32 = 2
                                var v50 int = (fi32 & fi32)
                                _ = v48
                                _ = v49
                                _ = v50
                            } else {
                                var v51 float32 = 0.1
                                _ = v51
                            }
                            v31 = v24
                            if len(v29) > 0 {
                                v29[0] = (fi32 * (fi32 % ((2147483647) | 1)))
                            }
                            if v23 {
                                if (!v23) {
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    if v23 {
                                        delete(v2, 2147483648)
                                        v52 := GenS2{F0: v1}
                                        v30 = v30
                                        _ = v52
                                    }
                                    v6[4294967295] = fi43
                                } else {
                                    if false {
                                        v23 = (v44 >= v44)
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        v53 := v4[fi43]
                                        _ = v53
                                    } else {
                                        var v54 int32 = v44
                                        _ = v54
                                    }
                                }
                                switch "" {
                                case "":
                                    v30 -= v30
                                case "hello":
                                    v23 = ((0 != uint8(v5)) && (1.5 <= 1e10))
                                case "a":
                                    var v55 int64 = -9223372036854775808
                                    _ = v55
                                default:
                                    switch ((v1 + "hello") + v1) {
                                    case "hello":
                                        var v56 uint32 = 2147483647
                                        _ = v56
                                    }
                                }
                            } else {
                                if len(v29) > 0 {
                                    v29[0] = fi43
                                }
                                var v57 int = 864
                                _ = v57
                            }
                            v31 -= 311
                            _ = v44
                            _ = v45
                            _ = v46
                        } else {
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            var v58 int8 = 100
                            _ = v58
                        }
                        fmt.Println(fi43)
                        var v59 int16 = v24
                        _ = fi43
                        _ = v59
                    }
                } else {
                    fmt.Println(v31, v5, v31)
                    v1 = v1
                }
                delete(v2, 4294967295)
            } else {
                var v60 int = fi32
                _ = v60
            }
        }
        _ = fi32
    }
    if false {
        fmt.Println(v28)
        switch 323 {
        case -2147483648:
            fmt.Println(v3)
        case 4294967295:
            if true {
                switch 1 {
                case 1:
                    if v23 {
                        v61 := GenS4{F0: "\n\t", F1: (!(3.14159 > (3.14159 - (float32(-2.25) * 1e10)))), F2: v28}
                        v1 = v1
                        _ = v61
                    }
                case 127:
                    var v62 int16 = -991
                    _ = v62
                case 32768:
                    switch "\n\t" {
                    case "foo bar":
                        if (2147483647 < -2147483648) {
                            v28 = 255
                            v24 = (v24 / ((1) | 1))
                            if (774 <= v3) {
                                switch v1 {
                                case "\n\t":
                                    switch -353 {
                                    case 32767:
                                        for fi63 := 0; fi63 < 3; fi63++ {
                                            switch v1 {
                                            case "hello":
                                                for fi64 := 0; fi64 < 19; fi64++ {
                                                    var v65 uint8 = 42
                                                    delete(v6, 1)
                                                    var v66 int64 = v30
                                                    _ = fi64
                                                    _ = v65
                                                    _ = v66
                                                }
                                            case "":
                                                v67 := GenS6{F0: (v28 - genFn0(42, "\n\t"))}
                                                _ = v67
                                            default:
                                                fmt.Println(v24)
                                            }
                                            _ = fi63
                                        }
                                    case 32768:
                                        v24 += v31
                                    case 0:
                                        switch v1 {
                                        case "\n\t":
                                            switch "" {
                                            case "a":
                                                v68 := GenS1{F0: -564}
                                                _ = v68
                                            case "hello":
                                                for fi69 := 0; fi69 < 15; fi69++ {
                                                    for fi70 := 0; fi70 < 12; fi70++ {
                                                        fmt.Println(v5, v3)
                                                        var v71 string = v1
                                                        var v72 string = "hello"
                                                        var v73 int8 = 42
                                                        v4[fi70] = v71
                                                        if v23 {
                                                            v1 = (v71 + (("\n\t" + ((((v72 + v71) + v72) + "hello") + v1)) + "hello"))
                                                            v23 = (v30 < (v30 >> ((42) & 63)))
                                                            var v74 uint = v5
                                                            _ = v74
                                                        }
                                                        v73 += (v73 ^ v73)
                                                        var v75 uint = v5
                                                        _ = fi70
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                        _ = v75
                                                    }
                                                    var v76 float64 = 3.14159
                                                    for fi77 := 0; fi77 < 7; fi77++ {
                                                        var v78 uint64 = v3
                                                        var v79 int8 = 0
                                                        _ = fi77
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    _ = fi69
                                                    _ = v76
                                                }
                                            case "foo bar":
                                                for fi80 := 0; fi80 < 6; fi80++ {
                                                    v81 := GenS7{F0: 0, F1: v31, F2: 1000}
                                                    fmt.Println(v24)
                                                    _ = fi80
                                                    _ = v81
                                                }
                                            default:
                                                var v82 uint16 = (v28 << (((v5 ^ 1)) & 63))
                                                _ = v82
                                            }
                                        case "":
                                            v24 += (v31 & (v31 | (v24 * 75)))
                                        case "a":
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v5, v23)
                                            case "hello":
                                                var v83 int16 = 1000
                                                _ = v83
                                            case "":
                                                v84 := GenS1{F0: v30}
                                                _ = v84
                                            default:
                                                v31 -= v24
                                            }
                                        }
                                    default:
                                        v30 = v30
                                    }
                                case "hello":
                                    for fi85 := 0; fi85 < 11; fi85++ {
                                        delete(v4, fi85)
                                        _ = fi85
                                    }
                                }
                                var v86 int16 = (v31 ^ (v31 / ((v31) | 1)))
                                if (!v23) {
                                    if false {
                                        var v87 uint32 = 2147483647
                                        if false {
                                            var v88 int64 = (-(v30))
                                            v30 -= (v30 >> ((v5) & 63))
                                            var v89 int = 0
                                            _ = v88
                                            _ = v89
                                        } else {
                                            for fi90 := 0; fi90 < 3; fi90++ {
                                                v91 := GenS4{F0: (v1 + ""), F1: v23, F2: genFn0(42, "foo bar")}
                                                v92 := GenS5{F0: v28, F1: (v28 ^ (v28 << (((v5 ^ 42)) & 63)))}
                                                _ = fi90
                                                _ = v91
                                                _ = v92
                                            }
                                            if ((^(v28)) != genFn0(-2147483648, ((("\n\t" + ((v1 + ((("a" + v1) + ((("" + v1) + v1) + (v1 + v1))) + v1)) + (v1 + ""))) + "") + "foo bar"))) {
                                                v3 = (v3 >> ((v5) & 63))
                                                var v93 string = v1
                                                v94 := GenS3{F0: 127, F1: 1}
                                                _ = v93
                                                _ = v94
                                            }
                                            v28 = v28
                                        }
                                        v95 := GenS8{F0: (32767 != 4294967295), F1: v1}
                                        _ = v87
                                        _ = v95
                                    } else {
                                        var v96 uint32 = 4294967295
                                        _ = v96
                                    }
                                    fmt.Println(v24)
                                    v6[1] = -307
                                }
                                _ = v86
                            }
                        } else {
                            if len(v29) > 0 {
                                v29[0] = -720
                            }
                            var v97 int8 = 42
                            _ = v97
                        }
                    case "hello":
                        var v98 float64 = (0.0 + ((0.0 * 1.5) / 1.0))
                        _ = v98
                    case "a":
                        switch 743 {
                        case 32767:
                            if (v30 > -418) {
                                for fi99 := 0; fi99 < 20; fi99++ {
                                    var v100 int16 = (v31 & (v24 | v24))
                                    v101 := GenS9{F0: (v3 + uint64(v5)), F1: (v30 / ((v30) | 1))}
                                    _ = fi99
                                    _ = v100
                                    _ = v101
                                }
                                for fi102 := 0; fi102 < 19; fi102++ {
                                    v3 -= (^(v3))
                                    switch "hello" {
                                    case "hello":
                                        v31 += v31
                                    }
                                    _ = fi102
                                }
                                for fi103 := 0; fi103 < 12; fi103++ {
                                    var v104 uint32 = 2147483647
                                    var v105 uint8 = 169
                                    var v106 float64 = float64((0.1 + (0.1 - -2.25)))
                                    _ = fi103
                                    _ = v104
                                    _ = v105
                                    _ = v106
                                }
                            } else {
                                for fi107 := 0; fi107 < 5; fi107++ {
                                    if len(v29) > 0 {
                                        v29[0] = int(v5)
                                    }
                                    _ = fi107
                                }
                            }
                        }
                    }
                default:
                    if false {
                        var v108 int16 = (v24 * v31)
                        for fi109 := 0; fi109 < 17; fi109++ {
                            for fi110 := 0; fi110 < 19; fi110++ {
                                if v23 {
                                    switch (fi109 * 0) {
                                    case 9223372036854775807:
                                        var v111 string = v1
                                        var v112 uint64 = (v3 & v3)
                                        _ = v111
                                        _ = v112
                                    case 128:
                                        v3 += 18446744073709551615
                                    case 127:
                                        v113 := GenS2{F0: v1}
                                        _ = v113
                                    }
                                    fmt.Println(v1)
                                } else {
                                    if v23 {
                                        _ = len(v29)
                                        var v114 bool = v23
                                        v115 := GenS1{F0: v30}
                                        _ = v114
                                        _ = v115
                                    } else {
                                        var v116 string = v1
                                        var v117 bool = (v23 || v23)
                                        if (250 == 127) {
                                            _ = len(v6)
                                            fmt.Println(v30, v30)
                                        }
                                        _ = v116
                                        _ = v117
                                    }
                                    v24 += v24
                                    switch "\n\t" {
                                    case "hello":
                                        v108 += v31
                                    case "foo bar":
                                        _ = len(v4)
                                    case "":
                                        var v118 float32 = 0.0
                                        _ = v118
                                    }
                                }
                                _ = fi110
                            }
                            var v119 int32 = 42
                            if (v31 > (v31 * (^(v31)))) {
                                v120 := GenS3{F0: (fi109 - (fi109 / ((4294967295) | 1))), F1: v119}
                                _ = v120
                            }
                            _ = fi109
                            _ = v119
                        }
                        _ = v108
                    }
                }
                if (v1 == "foo bar") {
                    var v121 uint = v5
                    _ = v121
                } else {
                    var v122 uint32 = 477
                    var v123 int8 = -94
                    _ = v122
                    _ = v123
                }
                switch (("" + "hello") + "\n\t") {
                case "":
                    if (!v23) {
                        fmt.Println(v31, v28, v1)
                    } else {
                        v24 = 32767
                        for fi124 := 0; fi124 < 19; fi124++ {
                            for fi125 := 0; fi125 < 12; fi125++ {
                                if (0 > -27) {
                                    v28 = uint16(v3)
                                    v126 := GenS10{F0: 42, F1: "\n\t", F2: "a"}
                                    _ = v126
                                } else {
                                    v127 := GenS11{F0: (v30 / ((42) | 1)), F1: ((^(v3)) <= v3)}
                                    var v128 uint32 = 956
                                    v129 := GenS11{F0: -1, F1: (((v3 <= (v3 / (((v3 & (v3 * 972))) | 1))) && true) || true)}
                                    _ = v127
                                    _ = v128
                                    _ = v129
                                }
                                _ = fi125
                            }
                            var v130 int8 = 127
                            _ = fi124
                            _ = v130
                        }
                        var v131 int32 = 42
                        _ = v131
                    }
                case "hello":
                    v24 += (v24 / ((v31) | 1))
                case "a":
                    v3 = (v3 << (((^(v5))) & 63))
                default:
                    var v132 string = ((v1 + v1) + "hello")
                    _ = v132
                }
            }
        default:
            var v133 int8 = 42
            _ = v133
        }
    } else {
        if len(v29) > 0 {
            v29[0] = 127
        }
        v29 = append(v29, 4294967295)
        for fi134 := 0; fi134 < 7; fi134++ {
            for fi135 := 0; fi135 < 3; fi135++ {
                v136 := GenS10{F0: 1, F1: v1, F2: v1}
                v31 = int16(fi135)
                _ = fi135
                _ = v136
            }
            var v137 uint8 = uint8(v3)
            _ = fi134
            _ = v137
        }
    }
    _ = len(v29)
    var v138 int64 = v30
    fmt.Println("v1:", v1)
    fmt.Println("v138:", v138)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v138
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
