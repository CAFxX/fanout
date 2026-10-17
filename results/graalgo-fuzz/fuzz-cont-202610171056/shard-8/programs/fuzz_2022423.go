package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 string; F1 uint32; F2 int16 }
type GenS10 struct { F0 int64; F1 int16; F2 uint64 }
type GenS11 struct { F0 int32; F1 uint8; F2 int32 }
type GenS12 struct { F0 uint16; F1 bool }
type GenS13 struct { F0 uint64; F1 int; F2 int8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int8; F1 int }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int; F1 int16 }
type GenS18 struct { F0 string; F1 int8; F2 int }

func main() {
    var v1 float64 = 1.5
    var v2 float64 = float64(3.14159)
    var v3 int32 = 1
    v4 := []string{}
    var v5 int64 = -430
    switch "\n\t" {
    case "a":
        var v6 int64 = v5
        _ = v6
    case "\n\t":
        fmt.Println(v3)
    case "foo bar":
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                if false {
                    v2 += 1.5
                } else {
                    if false {
                        if true {
                            if false {
                                if false {
                                    v3 -= v3
                                    _ = len(v4)
                                    if false {
                                        var v9 int64 = -1
                                        v4 = append(v4, ("a" + ""))
                                        _ = v9
                                    } else {
                                        var v10 int16 = -448
                                        _ = v10
                                    }
                                } else {
                                    v11 := GenS1{F0: fi7, F1: -9223372036854775808, F2: -80}
                                    if (!((!true) || true)) {
                                        var v12 uint8 = 128
                                        if (42 > 2) {
                                            switch ("a" + "a") {
                                            case "\n\t":
                                                v12 = 255
                                            case "a":
                                                var v13 uint8 = v12
                                                _ = v13
                                            case "":
                                                var v14 uint64 = 2
                                                _ = v14
                                            }
                                            var v15 float32 = float32(((float64(-2.25) * ((3.14159 / (((float64(0.1) + ((v1 / (((v2 / float64(-2.25)) + v2) / ((v2 / 1e10) * ((0.0 - v2) - (((v1 * ((v1 - float64(3.14159)) - v1)) - v2) - float64((3.14159 / 1.5))))))) - ((((v1 + v1) - ((v2 * (v1 - v2)) + ((v2 + (v2 + v2)) - (((v2 / v1) * 3.14159) - v2)))) / v1) + (v2 * 1e10)))) - ((float64(3.14159) + (v1 * (v2 + 0.0))) / (-2.25 + v1))) / (v2 / 0.1))) + 0.0)) + v2))
                                            _ = v15
                                        }
                                        _ = v12
                                    } else {
                                        v1 = (float64((0.1 / 0.1)) / 0.1)
                                        v1 += v1
                                        if (!(82 != -91)) {
                                            v4 = append(v4, "hello")
                                            v5 -= (v5 / (((v5 | 0)) | 1))
                                        }
                                        if false {
                                            v1 -= v1
                                            var v16 int8 = 76
                                            _ = v16
                                        }
                                    }
                                    _ = v11
                                }
                                _ = len(v4)
                                v2 = -2.25
                            } else {
                                v3 = (-(v3))
                            }
                            if true {
                                var v17 float64 = v2
                                if true {
                                    var v18 uint32 = 4294967295
                                    var v19 bool = false
                                    v18 -= (v18 / ((455) | 1))
                                    _ = v18
                                    _ = v19
                                } else {
                                    fmt.Println(fi8, v5)
                                    _ = len(v4)
                                }
                                _ = v17
                            }
                        } else {
                            v3 -= v3
                            var v20 int16 = -1
                            fmt.Println(v5, v1)
                            _ = v20
                        }
                        if (!((true || (988 != 0)) || ((!(!(65535 == 65535))) && true))) {
                            fmt.Println(fi8, fi8, v1)
                            v21 := GenS1{F0: fi7, F1: (v5 >> ((4294967295) & 63)), F2: 1}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v21
                        }
                        switch "hello" {
                        case "a":
                            v2 += (v2 + (v1 * (-2.25 * v2)))
                        case "foo bar":
                            v3 = (v3 % (((-(v3))) | 1))
                        }
                        switch 42 {
                        case -887:
                            var v22 uint16 = 255
                            _ = v22
                        default:
                            var v23 bool = true
                            _ = v23
                        }
                        var v24 uint = 18446744073709551615
                        _ = v24
                    } else {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        var v25 uint16 = 32767
                        v25 = 1
                        _ = v25
                    }
                    fmt.Println(fi8, fi7, v1)
                }
                _ = fi8
            }
            var v26 int16 = 1
            _ = fi7
            _ = v26
        }
    default:
        if (("\n\t" == ("hello" + "foo bar")) || false) {
            if true {
                var v27 int = 0
                var v28 uint32 = 1
                v5 -= 9223372036854775807
                _ = v27
                _ = v28
            }
        } else {
            var v29 float32 = (-2.25 + 1.5)
            v2 = (v1 + ((v1 / (-2.25 * v1)) - v2))
            _ = v29
        }
    }
    v4 = append(v4, "foo bar")
    if (2147483648 == 381) {
        if (true || false) {
            var v30 uint32 = 803
            fmt.Println(v3, v5)
            _ = v30
        }
    }
    v31 := GenS2{F0: v3, F1: 0, F2: 210}
    var v32 uint = 18446744073709551615
    var v33 int16 = -1
    if (!false) {
        fmt.Println(v5)
        v2 = float64(0.0)
        v34 := GenS2{F0: (v3 / ((v3) | 1)), F1: 18446744073709551615, F2: 2147483648}
        _ = v34
    } else {
        v4 = append(v4, "")
        v3 = (v3 + v3)
        v3 += (v3 >> ((255) & 63))
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            switch "a" {
            case "":
                v32 = v32
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch ("" + "a") {
            case "foo bar":
                v2 = 1.5
            case "a":
                var v37 uint = 627
                if true {
                    v38 := GenS2{F0: 1, F1: 2, F2: 4294967295}
                    var v39 string = "hello"
                    v4 = append(v4, v39)
                    _ = v38
                    _ = v39
                }
                if ("foo bar" == "") {
                    v37 += 63
                    _ = len(v4)
                    v37 += (v32 >> ((v37) & 63))
                    var v40 string = "a"
                    if (v5 >= v5) {
                        v41 := GenS1{F0: fi35, F1: (v5 - (v5 - int64(v3))), F2: 111}
                        v40 = "\n\t"
                        v37 -= v37
                        _ = v41
                    }
                    _ = v40
                } else {
                    v42 := GenS1{F0: fi35, F1: v5, F2: 100}
                    _ = v42
                }
                _ = v37
            case "hello":
                var v43 uint = (v32 & v32)
                _ = v43
            default:
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                v3 -= v3
                var v44 float64 = (v2 - (1.5 * v2))
                var v45 uint = v32
                fmt.Println(v5)
                _ = v44
                _ = v45
            }
            v33 = (v33 * 42)
            v4 = append(v4, ("a" + "\n\t"))
            v32 += 367
            _ = len(v4)
            _ = fi36
        }
        for fi46 := 0; fi46 < 5; fi46++ {
            if false {
                v2 = v2
                if true {
                    v2 = (((v2 * ((v2 / -2.25) + v1)) - 0.1) * 0.1)
                    var v47 int64 = (v5 % ((v5) | 1))
                    v48 := GenS3{F0: true}
                    v2 -= v2
                    fmt.Println(v33, v5, v5)
                    _ = v47
                    _ = v48
                } else {
                    var v49 string = ("a" + ("foo bar" + "foo bar"))
                    var v50 float64 = 1e10
                    if true {
                        v51 := GenS4{F0: 1, F1: 65535, F2: 345}
                        _ = v51
                    } else {
                        v52 := GenS3{F0: true}
                        switch (v49 + "hello") {
                        case "hello":
                            if false {
                                if true {
                                    var v53 string = "foo bar"
                                    switch fi35 {
                                    case 305:
                                        v2 += (v50 * (3.14159 - ((0.1 - -2.25) + (v50 / (float64(((0.0 - -2.25) / -2.25)) + (float64((float32(((v50 - v1) / (float64(0.0) - (3.14159 - (v1 - v2))))) + 3.14159)) / (-2.25 / v2)))))))
                                    }
                                    v32 -= (v32 - v32)
                                    _ = v53
                                } else {
                                    v54 := GenS1{F0: (fi35 & (-(fi35))), F1: (^(v5)), F2: 0}
                                    switch v49 {
                                    case "\n\t":
                                        fmt.Println(v5, v49, v49)
                                    }
                                    v3 = (v3 << ((0) & 63))
                                    v50 += (v2 + v1)
                                    v4 = append(v4, v49)
                                    _ = v54
                                }
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = ((v49 + ((v49 + v49) + (v49 + "hello"))) + "a")
                        }
                        var v55 uint = (v32 * v32)
                        _ = v52
                        _ = v55
                    }
                    var v56 int16 = 1000
                    _ = v49
                    _ = v50
                    _ = v56
                }
                if ((v3 + v3) >= v3) {
                    if (false && ((v3 - v3) == -1)) {
                        var v57 uint8 = 2
                        var v58 uint8 = v57
                        _ = v57
                        _ = v58
                    } else {
                        fmt.Println(v5)
                        _ = len(v4)
                        v2 -= 1e10
                    }
                    v59 := GenS5{F0: (v33 % ((v33) | 1))}
                    _ = v59
                } else {
                    if true {
                        var v60 uint16 = 255
                        var v61 int8 = -128
                        if false {
                            var v62 int64 = (-(v5))
                            v63 := GenS4{F0: (v61 | (v61 - v61)), F1: v60, F2: 1}
                            var v64 int = (fi35 * (fi46 & fi35))
                            v3 = (v3 << ((v32) & 63))
                            switch "a" {
                            case "foo bar":
                                v65 := GenS6{F0: 42}
                                _ = v65
                            case "hello":
                                v4 = append(v4, ("hello" + "hello"))
                            case "\n\t":
                                v66 := GenS4{F0: 42, F1: v60, F2: 42}
                                v64 = (fi35 << ((337) & 63))
                                _ = v66
                            }
                            v2 = v2
                            fmt.Println(v61)
                            _ = v62
                            _ = v63
                            _ = v64
                        } else {
                            if true {
                                var v67 int = (fi35 * (fi46 * fi46))
                                v32 = (v32 >> (((v32 & v32)) & 63))
                                if (false || false) {
                                    v68 := GenS7{F0: -9223372036854775808, F1: 2147483648}
                                    v69 := GenS8{F0: 377, F1: (v60 - v60)}
                                    _ = v68
                                    _ = v69
                                } else {
                                    if false {
                                        var v70 int8 = 1
                                        if false {
                                            v61 = v70
                                            if (false && true) {
                                                if true {
                                                    if (1 == (-(v33))) {
                                                        if (v33 != (v33 | (v33 >> (((v32 >> (((v32 / (((^(v32))) | 1))) & 63))) & 63)))) {
                                                            switch -292 {
                                                            case 42:
                                                                switch (^(v67)) {
                                                                case 32768:
                                                                    v70 = 126
                                                                }
                                                            default:
                                                                if false {
                                                                    v33 = 684
                                                                    if (false || false) {
                                                                        var v71 int64 = v5
                                                                        if false {
                                                                            var v72 int32 = (v3 / ((v3) | 1))
                                                                            fmt.Println(v1, v2, v32)
                                                                            switch (fi35 * (v67 - -225)) {
                                                                            case 32767:
                                                                                if false {
                                                                                    v32 += uint(v60)
                                                                                    v73 := GenS9{F0: "", F1: 0, F2: v33}
                                                                                    v1 = v1
                                                                                    fmt.Println(v70)
                                                                                    var v74 int64 = (^(v71))
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                } else {
                                                                                    v75 := GenS1{F0: (fi35 * (fi46 ^ fi35)), F1: v71, F2: (v70 ^ (v70 | v61))}
                                                                                    v76 := GenS6{F0: 703}
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                            _ = v72
                                                                        }
                                                                        fmt.Println(v60, v67)
                                                                        _ = v71
                                                                    }
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            }
                                                            v77 := GenS10{F0: v5, F1: -32768, F2: uint64(v32)}
                                                            v3 = -247
                                                            _ = v77
                                                        } else {
                                                            v67 -= 32768
                                                            var v78 uint8 = 2
                                                            if true {
                                                                if len(v4) > 0 {
                                                                    v4[0] = "a"
                                                                }
                                                            } else {
                                                                _ = len(v4)
                                                                var v79 int64 = v5
                                                                v79 -= 107
                                                                _ = v79
                                                            }
                                                            _ = v78
                                                        }
                                                        var v80 int16 = int16(v70)
                                                        v81 := GenS8{F0: 0, F1: (v60 >> ((42) & 63))}
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                }
                                                if true {
                                                    if ((42 == 0) || ((float32(v1) - 3.14159) != (1e10 * 3.14159))) {
                                                        var v82 uint64 = 344
                                                        var v83 int64 = v5
                                                        var v84 float64 = (1.5 - 1e10)
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        v85 := GenS10{F0: (v5 % (((v5 * (v5 >> (((v32 * v32)) & 63)))) | 1)), F1: int16(v67), F2: 42}
                                                        _ = v85
                                                    }
                                                    if (1000 > -517) {
                                                        v5 -= (v5 / ((v5) | 1))
                                                        var v86 int64 = -9223372036854775808
                                                        v32 = v32
                                                        _ = v86
                                                    } else {
                                                        v87 := GenS5{F0: (-(v33))}
                                                        var v88 bool = ((!(true || (!false))) && (((!(true && (-1 > v70))) || (2147483648 < 2147483647)) && false))
                                                        if false {
                                                            v32 += (v32 * (v32 << ((18446744073709551615) & 63)))
                                                            fmt.Println(v2, v5)
                                                            v89 := GenS5{F0: -1}
                                                            switch "a" {
                                                            case "a":
                                                                v5 = v5
                                                            case "foo bar":
                                                                var v90 int = fi46
                                                                _ = v90
                                                            }
                                                            fmt.Println(v61)
                                                            var v91 string = ("hello" + "foo bar")
                                                            v4 = append(v4, v91)
                                                            if v88 {
                                                                switch v91 {
                                                                case "foo bar":
                                                                    v92 := GenS4{F0: (v70 / (((v70 >> ((v32) & 63))) | 1)), F1: (^(v60)), F2: 339}
                                                                    _ = v92
                                                                default:
                                                                    v1 -= float64(1e10)
                                                                }
                                                            } else {
                                                                var v93 uint16 = 0
                                                                var v94 uint8 = 255
                                                                if ((^(v32)) < (v32 % ((921) | 1))) {
                                                                    v95 := GenS8{F0: 519, F1: v93}
                                                                    switch int(v5) {
                                                                    case 32768:
                                                                        var v96 uint = v32
                                                                        switch v91 {
                                                                        case "\n\t":
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        case "a":
                                                                            v97 := GenS7{F0: int(v96), F1: 2147483648}
                                                                            _ = v97
                                                                        default:
                                                                            if v88 {
                                                                                var v98 uint16 = (^(v60))
                                                                                v99 := GenS11{F0: -2147483648, F1: v94, F2: (^(v3))}
                                                                                if ((v91 >= (v91 + v91)) || false) {
                                                                                    fmt.Println(v32)
                                                                                    fmt.Println(v94, v70, fi35)
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v91
                                                                                    }
                                                                                } else {
                                                                                    var v100 int64 = (v5 * v5)
                                                                                    switch v91 {
                                                                                    case "\n\t":
                                                                                        var v101 int = fi46
                                                                                        _ = v101
                                                                                    default:
                                                                                        fmt.Println(v96, v3, v88)
                                                                                    }
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v91
                                                                                    }
                                                                                    _ = v100
                                                                                }
                                                                                _ = v98
                                                                                _ = v99
                                                                            } else {
                                                                                v102 := GenS4{F0: (-(v70)), F1: (v93 | (v93 >> ((v96) & 63))), F2: 0}
                                                                                var v103 int64 = v5
                                                                                v104 := GenS12{F0: v93, F1: v88}
                                                                                var v105 float64 = (v2 * float64((1.5 + (((float32((v2 * ((v2 * v1) / 1.0))) / float32((0.0 - v1))) + float32((v1 * v2))) - (0.1 / 0.1)))))
                                                                                _ = v102
                                                                                _ = v103
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                        }
                                                                        v3 = -170
                                                                        switch v91 {
                                                                        case "":
                                                                            var v106 uint8 = v94
                                                                            _ = v106
                                                                        }
                                                                        var v107 float32 = 0.0
                                                                        _ = v107
                                                                        _ = v96
                                                                    case 584:
                                                                        var v108 bool = ((v91 < v91) || true)
                                                                        _ = v108
                                                                    case 32767:
                                                                        var v109 uint16 = (^(v60))
                                                                        _ = v109
                                                                    default:
                                                                        v110 := GenS4{F0: -1, F1: (v93 & (v60 + (^(v60)))), F2: 2}
                                                                        _ = v110
                                                                    }
                                                                    var v111 uint32 = 1
                                                                    v61 = -1
                                                                    v112 := GenS13{F0: 4294967295, F1: (-(fi35)), F2: (v61 ^ v61)}
                                                                    fmt.Println(v67, v70, v111)
                                                                    var v113 uint64 = 4294967295
                                                                    v114 := GenS14{F0: v3}
                                                                    v91 = v91
                                                                    _ = v111
                                                                    _ = v112
                                                                    _ = v113
                                                                    _ = v114
                                                                    _ = v95
                                                                } else {
                                                                    if v88 {
                                                                        if false {
                                                                            if (true || ((1e10 / ((0.1 + 1.5) / 1e10)) <= (float32(float64(float32(1.5))) * -2.25))) {
                                                                                var v115 int8 = v70
                                                                                v33 -= (v33 >> (((^(v32))) & 63))
                                                                                switch (v67 | (v67 | (fi35 * fi35))) {
                                                                                case 544:
                                                                                    v5 = (v5 | (v5 + 42))
                                                                                }
                                                                                _ = v115
                                                                            } else {
                                                                                var v116 float64 = (v1 + ((v2 / v2) - v1))
                                                                                v94 += (^(v94))
                                                                                _ = v116
                                                                            }
                                                                            var v117 bool = v88
                                                                            _ = v117
                                                                        }
                                                                    } else {
                                                                        var v118 float64 = (float64(1e10) - (((0.1 / (v1 / -2.25)) - (-2.25 * 1e10)) - float64(0.0)))
                                                                        if v88 {
                                                                            v4 = append(v4, "a")
                                                                            var v119 int = (fi35 / (((v67 << ((469) & 63))) | 1))
                                                                            _ = v119
                                                                        } else {
                                                                            _ = len(v4)
                                                                            v93 += (v93 * (v93 << (((v32 ^ v32)) & 63)))
                                                                        }
                                                                        if len(v4) > 0 {
                                                                            v4[0] = "\n\t"
                                                                        }
                                                                        _ = v118
                                                                    }
                                                                    if ("a" > (v91 + ((v91 + (((v91 + v91) + (v91 + v91)) + v91)) + v91))) {
                                                                        switch 934 {
                                                                        case 0:
                                                                            if len(v4) > 0 {
                                                                                v4[0] = v91
                                                                            }
                                                                        default:
                                                                            if false {
                                                                                var v120 int64 = v5
                                                                                switch (v91 + v91) {
                                                                                case "foo bar":
                                                                                    var v121 float64 = 0.1
                                                                                    _ = v121
                                                                                case "hello":
                                                                                    var v122 int = 0
                                                                                    _ = v122
                                                                                case "\n\t":
                                                                                    v123 := GenS15{F0: 1, F1: (v67 | -963)}
                                                                                    _ = v123
                                                                                }
                                                                                v70 += (v70 & v70)
                                                                                _ = v120
                                                                            } else {
                                                                                v91 = v91
                                                                                var v124 int64 = (v5 >> (((^(v32))) & 63))
                                                                                _ = v124
                                                                            }
                                                                        }
                                                                        fmt.Println(v61, v93)
                                                                        v91 = "a"
                                                                    } else {
                                                                        fmt.Println(v33, v1)
                                                                    }
                                                                    var v125 uint16 = (v60 % (((^(v60))) | 1))
                                                                    _ = v125
                                                                }
                                                                v126 := GenS4{F0: 42, F1: v60, F2: uint32(v60)}
                                                                if (true || v88) {
                                                                    v127 := GenS16{F0: 1}
                                                                    v1 -= (1e10 / 1e10)
                                                                    _ = v127
                                                                }
                                                                _ = v126
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            switch 4294967295 {
                                                            case -923:
                                                                v128 := GenS7{F0: fi35, F1: 65535}
                                                                _ = v128
                                                            }
                                                            _ = v89
                                                            _ = v91
                                                        } else {
                                                            if v88 {
                                                                v129 := GenS7{F0: 42, F1: 1}
                                                                _ = v129
                                                            }
                                                        }
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    v4 = append(v4, "")
                                                } else {
                                                    v61 += (v61 * 127)
                                                    if (false && true) {
                                                        var v130 int32 = int32(v32)
                                                        var v131 uint16 = v60
                                                        if true {
                                                            if (fi46 <= (fi46 & fi35)) {
                                                                v4 = append(v4, "foo bar")
                                                                var v132 int16 = v33
                                                                if false {
                                                                    v5 += v5
                                                                    v32 -= v32
                                                                } else {
                                                                    v133 := GenS17{F0: fi46, F1: (v33 - v132)}
                                                                    _ = v133
                                                                }
                                                                _ = v132
                                                            } else {
                                                                var v134 uint8 = 255
                                                                v4 = append(v4, ("a" + ("foo bar" + ("hello" + "hello"))))
                                                                v134 -= v134
                                                                _ = v134
                                                            }
                                                        } else {
                                                            v33 += v33
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v130
                                                        _ = v131
                                                    } else {
                                                        var v135 int32 = v3
                                                        v136 := GenS2{F0: 0, F1: 732, F2: 2147483648}
                                                        v137 := GenS17{F0: 1, F1: (v33 >> ((v32) & 63))}
                                                        _ = v135
                                                        _ = v136
                                                        _ = v137
                                                    }
                                                }
                                            }
                                        }
                                        _ = v70
                                    } else {
                                        v5 = v5
                                        var v138 string = "hello"
                                        if (!(!true)) {
                                            v32 -= v32
                                            fmt.Println(v138)
                                        } else {
                                            _ = len(v4)
                                        }
                                        _ = v138
                                    }
                                    switch "a" {
                                    case "hello":
                                        v67 = fi46
                                    }
                                    if ((331 < 2) || true) {
                                        v33 += (-(v33))
                                        _ = len(v4)
                                    }
                                }
                                _ = v67
                            } else {
                                v4 = append(v4, "")
                                v5 += v5
                            }
                            v4 = append(v4, ("" + ""))
                            if len(v4) > 0 {
                                v4[0] = (((("a" + "") + ("foo bar" + ("hello" + "hello"))) + ("foo bar" + "")) + ("" + ("" + "a")))
                            }
                        }
                        _ = v60
                        _ = v61
                    } else {
                        v5 -= (v5 >> ((v32) & 63))
                        var v139 float64 = 0.0
                        _ = v139
                    }
                    var v140 string = "hello"
                    _ = v140
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v141 := GenS7{F0: int(v3), F1: 42}
                _ = len(v4)
                _ = v141
            }
            switch (fi35 * (-(fi46))) {
            case 128:
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            case 150:
                if false {
                    if (1 > 1) {
                        var v142 int64 = v5
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            v5 -= v142
                        case "":
                            var v143 uint32 = 1
                            _ = v143
                        case "a":
                            if (1e10 != float32(1e10)) {
                                v1 -= (v2 * ((v2 + (1e10 * v2)) / v1))
                            }
                        default:
                            var v144 uint16 = 1
                            _ = v144
                        }
                        _ = v142
                    }
                } else {
                    if (false && true) {
                        switch "a" {
                        case "":
                            var v145 string = "foo bar"
                            _ = v145
                        case "\n\t":
                            var v146 uint8 = 61
                            _ = v146
                        case "hello":
                            v147 := GenS18{F0: "\n\t", F1: -128, F2: (fi35 - 2147483647)}
                            _ = v147
                        }
                    } else {
                        var v148 uint8 = 1
                        var v149 int64 = v5
                        v33 -= v33
                        _ = v148
                        _ = v149
                    }
                }
            default:
                if true {
                    switch ("hello" + ("a" + "\n\t")) {
                    case "\n\t":
                        v2 -= v2
                    case "a":
                        if len(v4) > 0 {
                            v4[0] = (("foo bar" + "") + "foo bar")
                        }
                        var v150 uint8 = 167
                        var v151 int = (fi46 + fi46)
                        _ = v150
                        _ = v151
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi35, v3)
                }
            }
            var v152 uint32 = 2
            _ = fi46
            _ = v152
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi35
    }
    for fi153 := 0; fi153 < 5; fi153++ {
        v2 -= v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v154 int32 = (^(v3))
        _ = fi153
        _ = v154
    }
    var v155 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v155:", v155)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v155
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
