package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS6 struct { F0 int; F1 bool; F2 int16 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 uint64; F1 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64; F1 uint; F2 int32 }
type GenS12 struct { F0 string; F1 uint64 }
type GenS13 struct { F0 string; F1 string }
type GenS14 struct { F0 int32; F1 uint8; F2 int }

func main() {
    var v1 string = "a"
    var v2 int32 = 2147483647
    var v3 int32 = v2
    var v4 uint8 = 42
    fmt.Println(v4, v4)
    if (0 >= 94) {
        if true {
            v5 := GenS1{F0: 575, F1: 2147483647}
            if (1000 < 1) {
                for fi6 := 0; fi6 < 2; fi6++ {
                    v7 := GenS1{F0: 42, F1: -938}
                    v2 -= v2
                    v4 = v4
                    _ = fi6
                    _ = v7
                }
                v8 := GenS2{F0: 255}
                var v9 float32 = -2.25
                _ = v8
                _ = v9
            }
            _ = v5
        }
        switch -9223372036854775808 {
        case 0:
            v10 := GenS2{F0: v4}
            _ = v10
        case -2147483648:
            var v11 int8 = -1
            _ = v11
        case 1:
            v1 = v1
        default:
            if ((uint64(v4) >= 1) || (!false)) {
                v3 = 2147483647
                v12 := []int{32767, 32767}
                _ = v12
            } else {
                var v13 string = ("foo bar" + v1)
                v14 := GenS3{F0: -661, F1: 620}
                if true {
                    var v15 float64 = float64(1e10)
                    _ = v15
                } else {
                    var v16 int16 = -1
                    for fi17 := 0; fi17 < 4; fi17++ {
                        for fi18 := 0; fi18 < 12; fi18++ {
                            if true {
                                if true {
                                    if (3.14159 > 0.1) {
                                        v19 := GenS4{F0: (v3 + (v2 ^ (v3 & (v3 / (((v3 + v2)) | 1)))))}
                                        var v20 uint = 4294967295
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v21 := []int{-459, fi18}
                                        switch (v1 + "") {
                                        case "a":
                                            v4 += 74
                                        case "":
                                            if false {
                                                fmt.Println(fi17)
                                            }
                                        default:
                                            v1 = v1
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v4)
                                                if len(v21) > 0 {
                                                    v21[0] = fi17
                                                }
                                                if len(v21) > 0 {
                                                    v21[0] = 128
                                                }
                                            case "hello":
                                                if false {
                                                    v1 = (v1 + v1)
                                                }
                                                v4 += (v4 / (((^(v4))) | 1))
                                                var v22 int16 = (v16 << ((42) & 63))
                                                _ = v22
                                            case "\n\t":
                                                var v23 string = ("hello" + "\n\t")
                                                _ = v23
                                            }
                                        }
                                        v16 = (v16 ^ 180)
                                        _ = v21
                                    }
                                    v2 -= v3
                                    if true {
                                        var v24 float32 = float32((1.5 + (((0.1 / ((((0.0 - 1e10) + 3.14159) * float64(float32(float64(0.1)))) - 1e10)) - (-2.25 + 3.14159)) + -2.25)))
                                        _ = v24
                                    }
                                }
                            }
                            switch 42 {
                            case 128:
                                if (true && (((v4 & 1) >= 9) || (float32(3.14159) < 1e10))) {
                                    v25 := GenS5{F0: 65535, F1: 42, F2: 564}
                                    _ = v25
                                }
                            case -1:
                                v26 := GenS6{F0: (fi17 - fi18), F1: (((3.14159 - 1e10) + float64((1.5 + 0.0))) == 0.1), F2: v16}
                                _ = v26
                            default:
                                var v27 uint64 = 0
                                _ = v27
                            }
                            if (!false) {
                                v16 = (v16 | (^(v16)))
                                v1 = v13
                                switch -215 {
                                case -702:
                                    var v28 uint64 = 18446744073709551615
                                    _ = v28
                                case 9223372036854775807:
                                    switch 4294967295 {
                                    case 1:
                                        v16 -= 0
                                    case -1:
                                        var v29 int32 = (v3 + v2)
                                        _ = v29
                                    case 32768:
                                        fmt.Println(v3)
                                    }
                                case 32768:
                                    if true {
                                        v2 = v2
                                    } else {
                                        v30 := []int{(fi17 | 127), fi17}
                                        v3 += 2147483647
                                        _ = v30
                                    }
                                default:
                                    var v31 bool = false
                                    _ = v31
                                }
                            } else {
                                var v32 int16 = (-(v16))
                                v33 := []int{(fi17 & int(v4)), fi17, 42, fi18}
                                _ = v32
                                _ = v33
                            }
                            _ = fi18
                        }
                        _ = fi17
                    }
                    _ = v16
                }
                _ = v13
                _ = v14
            }
        }
    }
    var v34 int32 = v3
    if (1 < 42) {
        if (!(false || (false || (v1 <= v1)))) {
            var v35 int32 = 1
            fmt.Println(v35)
            _ = v35
        } else {
            v36 := []int{}
            _ = v36
        }
        var v37 int = 399
        _ = v37
    } else {
        v38 := GenS7{F0: -1, F1: 42}
        _ = v38
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 14; fi40++ {
            v41 := GenS4{F0: v2}
            v42 := GenS1{F0: 42, F1: (fi39 - (fi39 >> ((255) & 63)))}
            var v43 float32 = 1e10
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
        }
        var v44 bool = (2 > 2)
        for fi45 := 0; fi45 < 13; fi45++ {
            switch "" {
            case "a":
                var v46 float32 = float32(1e10)
                v47 := GenS6{F0: (-(fi45)), F1: (v44 && (!((!v44) || (!v44)))), F2: 0}
                if (!true) {
                    v48 := GenS7{F0: -531, F1: -1}
                    v49 := GenS8{F0: 0, F1: v1}
                    _ = v48
                    _ = v49
                } else {
                    switch fi45 {
                    case 32767:
                        fmt.Println(fi45, v3, fi39)
                    case 42:
                        v44 = (32767 != -1)
                    default:
                        v44 = v44
                    }
                    fmt.Println(v46)
                    if (65535 == 2) {
                        var v50 string = v1
                        _ = v50
                    } else {
                        var v51 uint8 = v4
                        var v52 int64 = 0
                        var v53 uint = 1
                        v54 := []int{}
                        var v55 int64 = 80
                        fmt.Println(v53)
                        var v56 float32 = 1.5
                        _ = v51
                        _ = v52
                        _ = v53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                    v57 := GenS9{F0: 0}
                    v58 := []int{}
                    switch v1 {
                    case "":
                        if v44 {
                            var v59 uint8 = v4
                            fmt.Println(v3, v1)
                            v2 -= 2147483647
                            _ = v59
                        } else {
                            fmt.Println(v4, v3)
                            switch (fi39 + (fi45 | 32767)) {
                            case 475:
                                if (1.5 > 1.5) {
                                    var v60 uint64 = 18446744073709551615
                                    switch (("a" + v1) + v1) {
                                    case "\n\t":
                                        var v61 uint16 = 65535
                                        _ = v61
                                    case "a":
                                        if len(v58) >= 2 {
                                            v58 = v58[1:len(v58)]
                                        }
                                    case "":
                                        if (v44 || true) {
                                            var v62 int8 = int8(fi45)
                                            switch "foo bar" {
                                            case "":
                                                if (((^(v62)) >= (v62 / (((v62 << ((42) & 63))) | 1))) || v44) {
                                                    var v63 int8 = v62
                                                    v46 += (v46 - 0.0)
                                                    v64 := GenS10{F0: ((^(fi39)) != fi39)}
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    if v44 {
                                                        fmt.Println(v60, v2, fi45)
                                                        switch (fi39 / ((-87) | 1)) {
                                                        case 128:
                                                            v34 -= v3
                                                        case -9223372036854775808:
                                                            v1 = ""
                                                        case -1:
                                                            v34 = v34
                                                        default:
                                                            var v65 uint16 = 630
                                                            _ = v65
                                                        }
                                                        if v44 {
                                                            fmt.Println(v3, v3, v44)
                                                            if ("hello" < "a") {
                                                                v60 += (v60 * v60)
                                                            }
                                                            v62 -= (v62 / ((-1) | 1))
                                                            var v66 int64 = -1
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                            _ = v66
                                                        } else {
                                                            v67 := GenS9{F0: fi39}
                                                            v68 := GenS7{F0: (v2 - (v2 % ((42) | 1))), F1: 966}
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                if v44 {
                                                    if (!v44) {
                                                        switch (-(fi45)) {
                                                        case 4294967295:
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                        default:
                                                            v69 := GenS11{F0: 0, F1: 1, F2: (v2 - 2147483647)}
                                                            _ = v69
                                                        }
                                                        var v70 int32 = -2147483648
                                                        _ = v70
                                                    } else {
                                                        _ = len(v58)
                                                    }
                                                    v71 := GenS9{F0: (fi39 >> ((255) & 63))}
                                                    _ = v71
                                                } else {
                                                    var v72 bool = (v44 && (uint64(v4) != (v60 % ((1) | 1))))
                                                    v60 = v60
                                                    _ = v72
                                                }
                                            case "a":
                                                var v73 string = "\n\t"
                                                _ = v73
                                            default:
                                                v34 = -634
                                            }
                                            _ = v62
                                        }
                                    }
                                    _ = v60
                                }
                            default:
                                var v74 float32 = ((v46 - (0.0 / 1.5)) + ((float32(1.5) * ((v46 - (float32(1e10) - v46)) - (1e10 + ((v46 - (v46 * (float32((((1e10 / 1.5) * (float64(v46) + 0.0)) * 3.14159)) / -2.25))) / (-2.25 + ((float32(float64(float32(0.1))) * 0.1) - v46)))))) + v46))
                                var v75 uint64 = 18446744073709551615
                                _ = v74
                                _ = v75
                            }
                            if true {
                                v76 := GenS12{F0: ((((((("foo bar" + v1) + (v1 + ("a" + v1))) + (v1 + (v1 + "\n\t"))) + (v1 + v1)) + "\n\t") + "") + "a"), F1: 4294967295}
                                v77 := GenS7{F0: v2, F1: 1}
                                _ = len(v58)
                                _ = v76
                                _ = v77
                            }
                            if (v46 != 0.1) {
                                switch v1 {
                                case "foo bar":
                                    var v78 string = ("" + v1)
                                    _ = v78
                                case "\n\t":
                                    v58 = append(v58, -343)
                                case "":
                                    var v79 int8 = 0
                                    _ = v79
                                }
                                var v80 uint16 = 147
                                _ = v80
                            }
                            var v81 bool = true
                            _ = v81
                        }
                    case "foo bar":
                        if len(v58) >= 2 {
                            v58 = v58[1:len(v58)]
                        }
                    }
                    v82 := GenS10{F0: (1 <= 0)}
                    _ = v57
                    _ = v58
                    _ = v82
                }
                _ = v46
                _ = v47
            case "foo bar":
                var v83 uint32 = 42
                _ = v83
            case "hello":
                if (v44 && v44) {
                    v84 := []int{(fi39 % ((fi45) | 1)), fi45, 9223372036854775807, int(v34)}
                    var v85 uint8 = v4
                    fmt.Println(v4, fi39)
                    _ = v84
                    _ = v85
                }
            default:
                v86 := GenS13{F0: "\n\t", F1: ""}
                _ = v86
            }
            _ = fi45
        }
        _ = fi39
        _ = v44
    }
    v87 := GenS9{F0: 42}
    var v88 float32 = 0.0
    switch "hello" {
    case "a":
        var v89 string = (v1 + ("" + "foo bar"))
        _ = v89
    case "foo bar":
        for fi90 := 0; fi90 < 9; fi90++ {
            v91 := []int{(fi90 >> ((4294967295) & 63)), (fi90 << ((4294967295) & 63))}
            _ = fi90
            _ = v91
        }
    case "hello":
        switch 2147483647 {
        case 2147483647:
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v3, v2, v2)
            case "a":
                for fi92 := 0; fi92 < 2; fi92++ {
                    v93 := []int{32767, (fi92 >> ((0) & 63))}
                    v94 := GenS4{F0: 581}
                    _ = fi92
                    _ = v93
                    _ = v94
                }
            case "":
                v95 := GenS14{F0: v34, F1: v4, F2: 127}
                _ = v95
            default:
                if true {
                    if (-1 > -1) {
                        for fi96 := 0; fi96 < 18; fi96++ {
                            var v97 int16 = 1
                            _ = fi96
                            _ = v97
                        }
                    } else {
                        if false {
                            var v98 uint64 = 0
                            fmt.Println(v88, v88, v4)
                            _ = v98
                        } else {
                            var v99 bool = false
                            for fi100 := 0; fi100 < 4; fi100++ {
                                var v101 uint16 = 0
                                _ = fi100
                                _ = v101
                            }
                            var v102 string = v1
                            _ = v102
                            _ = v99
                        }
                    }
                    v103 := GenS8{F0: 306, F1: "hello"}
                    v104 := GenS8{F0: 630, F1: v1}
                    _ = v103
                    _ = v104
                } else {
                    switch (v1 + "") {
                    case "hello":
                        v88 += v88
                    case "\n\t":
                        v88 = float32(((float64(v88) - (-2.25 / float64(v88))) * 3.14159))
                    }
                    v2 = int32(v4)
                    if (!false) {
                        fmt.Println(v3, v2, v34)
                    }
                }
            }
        case 127:
            fmt.Println(v2, v4, v34)
        case -943:
            switch 467 {
            case -780:
                var v105 int8 = 1
                _ = v105
            case -380:
                v88 = v88
            case 954:
                var v106 int8 = -1
                _ = v106
            }
        default:
            var v107 bool = true
            _ = v107
        }
    default:
        var v108 uint32 = 65535
        _ = v108
    }
    var v109 int32 = (v2 * (^(v2)))
    for fi110 := 0; fi110 < 19; fi110++ {
        fmt.Println(v109, v4, v109)
        for fi111 := 0; fi111 < 19; fi111++ {
            var v112 int32 = v3
            var v113 uint32 = 0
            var v114 uint32 = v113
            v3 = v34
            _ = fi111
            _ = v112
            _ = v113
            _ = v114
        }
        _ = fi110
    }
    fmt.Println(v88, v34)
    var v115 int64 = -588
    v115 -= v115
    v88 = ((float32((1.5 + 0.1)) + v88) + v88)
    v116 := []int{-2147483648, 42}
    var v117 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v115:", v115)
    fmt.Println("slicelen v116:", len(v116), v116)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v109
    _ = v115
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v87
    _ = v88
}
