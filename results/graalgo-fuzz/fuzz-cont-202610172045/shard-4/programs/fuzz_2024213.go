package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint; F1 bool }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 bool; F1 int64; F2 int8 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 int16; F1 int32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16; F1 uint; F2 string }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 uint32; F1 string }
type GenS19 struct { F0 string }
type GenS20 struct { F0 string; F1 int32 }
type GenS21 struct { F0 int8; F1 uint8 }
type GenS22 struct { F0 uint; F1 uint32; F2 int }
type GenS23 struct { F0 bool; F1 uint; F2 uint32 }
type GenS24 struct { F0 bool; F1 uint; F2 uint32 }
type GenS25 struct { F0 int8; F1 uint32 }
type GenS26 struct { F0 int; F1 int8 }
func genFn0() bool {
    for fi6 := 0; fi6 < 6; fi6++ {
        if (false && (1 > 0)) {
            var v7 string = "hello"
            _ = v7
        } else {
            var v8 float32 = 1.5
            v8 = v8
            _ = v8
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint32 = 2147483648
            var v11 uint64 = 42
            v11 = v11
            if ((v10 ^ (v10 >> ((568) & 63))) > v10) {
                v12 := GenS1{F0: 1, F1: false, F2: v11}
                if false {
                    switch "a" {
                    case "":
                        var v13 bool = (1 < 920)
                        _ = v13
                    default:
                        v14 := GenS2{F0: 1, F1: 84}
                        _ = v14
                    }
                    var v15 float32 = (1.5 + -2.25)
                    _ = v15
                }
                _ = v12
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi6
    }
    v16 := []int{-745}
    switch 216 {
    case -221:
        if true {
            v17 := GenS2{F0: 1, F1: 100}
            var v18 int32 = 0
            v18 = -2147483648
            _ = v17
            _ = v18
        }
    case -66:
        var v19 uint32 = 0
        _ = v19
    }
    var v20 uint16 = 104
    fmt.Println(v20)
    _ = v16
    _ = v20
    return (!false)
}

func main() {
    v1 := []uint32{2147483648, 42}
    var v2 int32 = 663
    var v3 int64 = 750
    v4 := make(map[uint32]int)
    var v5 uint = 186
    switch "foo bar" {
    case "":
        for fi21 := 0; fi21 < 20; fi21++ {
            switch -9223372036854775808 {
            case 0:
                v3 -= -1
            case 700:
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 bool = ((true || (0.1 > 1e10)) || true)
                    if (v23 || v23) {
                        if false {
                            v3 = (v3 << (((v5 % (((^(v5))) | 1))) & 63))
                        } else {
                            var v24 uint = (v5 >> ((v5) & 63))
                            _ = v24
                        }
                        v25 := GenS2{F0: 1, F1: 0}
                        _ = v25
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi22
                    _ = v23
                }
            case -9223372036854775808:
                v26 := GenS3{F0: 773, F1: "foo bar"}
                _ = v26
            }
            _ = fi21
        }
    case "foo bar":
        switch 2147483647 {
        case 32768:
            v27 := GenS4{F0: ("hello" + "\n\t")}
            _ = v27
        case -1:
            for fi28 := 0; fi28 < 9; fi28++ {
                fmt.Println(v5)
                var v29 int32 = v2
                v30 := GenS4{F0: ("hello" + "hello")}
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            if len(v1) > 0 {
                v1[0] = 858
            }
        }
    case "hello":
        for fi31 := 0; fi31 < 14; fi31++ {
            v32 := GenS5{F0: 2, F1: 985}
            _ = fi31
            _ = v32
        }
    default:
        var v33 float64 = (0.1 + 0.1)
        _ = v33
    }
    fmt.Println(v3, v2)
    var v34 string = "hello"
    var v35 bool = (!true)
    var v36 uint8 = 2
    _ = len(v4)
    switch "a" {
    case "":
        if (18 != 42) {
            var v37 int = -9223372036854775808
            v3 += v3
            switch v37 {
            case 4294967295:
                v3 = v3
            default:
                if (!v35) {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        v39 := GenS6{F0: (v34 + ("\n\t" + v34)), F1: v34}
                        for fi40 := 0; fi40 < 2; fi40++ {
                            v34 = "\n\t"
                            if ((1e10 + ((0.0 - 0.0) + 1e10)) >= (3.14159 / (((-2.25 * 3.14159) * 0.0) + (float64(1e10) - 0.0)))) {
                                if (-1 != 1) {
                                    v5 = (^(v5))
                                }
                            }
                            _ = fi40
                        }
                        _ = fi38
                        _ = v39
                    }
                    var v41 uint32 = 2147483647
                    if (!false) {
                        switch v34 {
                        case "\n\t":
                            v42 := GenS6{F0: (v34 + v34), F1: v34}
                            _ = v42
                        case "":
                            var v43 uint16 = 522
                            _ = v43
                        case "a":
                            var v44 float64 = 0.1
                            _ = v44
                        }
                        v41 = 4294967295
                        delete(v4, (v41 + v41))
                    } else {
                        if v35 {
                            if false {
                                var v45 int16 = 247
                                _ = v45
                            } else {
                                if (731 != v37) {
                                    v41 = v41
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (^(v41))
                                    }
                                    switch v34 {
                                    case "a":
                                        var v46 int32 = (^(v2))
                                        _ = v46
                                    case "\n\t":
                                        v37 = -2147483648
                                    case "foo bar":
                                        v37 = v37
                                    default:
                                        var v47 uint64 = 18446744073709551615
                                        _ = v47
                                    }
                                    switch "" {
                                    case "hello":
                                        v48 := v4[v41]
                                        _ = v48
                                    case "":
                                        for fi49 := 0; fi49 < 10; fi49++ {
                                            fmt.Println(v37, v41)
                                            v50 := GenS5{F0: 4294967295, F1: uint64(v5)}
                                            _ = fi49
                                            _ = v50
                                        }
                                    case "a":
                                        v34 = ("a" + "foo bar")
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = v41
                                }
                                v37 -= (v37 - v37)
                            }
                            v36 -= (v36 ^ (v36 + 128))
                            switch (v34 + v34) {
                            case "\n\t":
                                for fi51 := 0; fi51 < 4; fi51++ {
                                    var v52 uint32 = (v41 ^ 2)
                                    switch "foo bar" {
                                    case "":
                                        v53 := GenS5{F0: 4294967295, F1: 67}
                                        _ = v53
                                    }
                                    switch (^(fi51)) {
                                    case 127:
                                        if (fi51 < (fi51 / ((v37) | 1))) {
                                            var v54 int32 = v2
                                            if (1 != 42) {
                                                if (255 > 32767) {
                                                    if (!v35) {
                                                        v55 := GenS7{F0: 0, F1: 255}
                                                        _ = v55
                                                    } else {
                                                        for fi56 := 0; fi56 < 8; fi56++ {
                                                            var v57 uint16 = 760
                                                            _ = fi56
                                                            _ = v57
                                                        }
                                                        if false {
                                                            var v58 int64 = v3
                                                            var v59 int64 = (v3 + v3)
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            v4[(^(v52))] = (fi51 + int(v2))
                                                            v54 = (v2 >> ((v5) & 63))
                                                            var v60 float32 = 1e10
                                                            _ = v60
                                                        }
                                                        if (!v35) {
                                                            for fi61 := 0; fi61 < 7; fi61++ {
                                                                v62 := GenS7{F0: 0, F1: 0}
                                                                _ = len(v4)
                                                                _ = fi61
                                                                _ = v62
                                                            }
                                                            var v63 uint32 = (v52 ^ 2147483648)
                                                            for fi64 := 0; fi64 < 11; fi64++ {
                                                                _ = len(v1)
                                                                var v65 int32 = v2
                                                                v65 = int32(v5)
                                                                _ = fi64
                                                                _ = v65
                                                            }
                                                            _ = v63
                                                        } else {
                                                            v66 := GenS8{F0: (v5 + (v5 / ((uint(v36)) | 1))), F1: 1}
                                                            var v67 uint = v5
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    }
                                                } else {
                                                    v4[798] = -931
                                                    if v35 {
                                                        if false {
                                                            fmt.Println(v34, v5, v34)
                                                            fmt.Println(v36, v5)
                                                        }
                                                        v5 = v5
                                                    } else {
                                                        if (!(!(!true))) {
                                                            for fi68 := 0; fi68 < 2; fi68++ {
                                                                if v35 {
                                                                    var v69 int = (^(fi68))
                                                                    switch v34 {
                                                                    case "\n\t":
                                                                        if (2147483648 >= v41) {
                                                                            if false {
                                                                                if (1e10 != 0.0) {
                                                                                    delete(v4, 971)
                                                                                }
                                                                                if v35 {
                                                                                    switch fi51 {
                                                                                    case 304:
                                                                                        fmt.Println(fi68, v35)
                                                                                    default:
                                                                                        v36 -= 42
                                                                                    }
                                                                                }
                                                                                var v70 string = (("" + v34) + v34)
                                                                                _ = v70
                                                                            } else {
                                                                                v4[1] = (v69 & (^(v37)))
                                                                                var v71 float32 = float32(3.14159)
                                                                                _ = v71
                                                                            }
                                                                            v72 := GenS9{F0: v5, F1: ((173 > v69) || (int32(v3) >= (v54 + int32(v3))))}
                                                                            var v73 int64 = v3
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            v74 := GenS3{F0: v52, F1: v34}
                                                                            _ = v74
                                                                        }
                                                                    case "":
                                                                        var v75 int64 = v3
                                                                        if v35 {
                                                                            var v76 float32 = 0.1
                                                                            if (false || (0 <= 127)) {
                                                                                switch v34 {
                                                                                case "\n\t":
                                                                                    if (-2.25 > v76) {
                                                                                        var v77 int64 = (v3 - v75)
                                                                                        v36 = (v36 << (((v5 - 578)) & 63))
                                                                                        _ = v77
                                                                                    }
                                                                                case "foo bar":
                                                                                    v76 = (float32((3.14159 + -2.25)) / 1e10)
                                                                                case "hello":
                                                                                    v75 = (v75 - (^(v3)))
                                                                                    var v78 bool = false
                                                                                    var v79 int32 = (v2 / ((0) | 1))
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                default:
                                                                                    v1 = append(v1, (v41 + (v52 ^ 2)))
                                                                                }
                                                                                v80 := GenS1{F0: v52, F1: (v35 || v35), F2: 470}
                                                                                var v81 uint64 = 18446744073709551615
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                        _ = v75
                                                                    default:
                                                                        switch (v69 * (^(fi51))) {
                                                                        case -2147483648:
                                                                            v82 := GenS4{F0: v34}
                                                                            _ = v82
                                                                        default:
                                                                            fmt.Println(v36, v34)
                                                                        }
                                                                    }
                                                                    if ((v69 & fi51) >= (v37 ^ fi51)) {
                                                                        if ((v34 + v34) != ((v34 + "") + v34)) {
                                                                            if (v52 > uint32(v5)) {
                                                                                v83 := GenS10{F0: (^(v41)), F1: 4294967295}
                                                                                v5 += 4294967295
                                                                                switch int(v36) {
                                                                                case 32767:
                                                                                    if v35 {
                                                                                        v36 = v36
                                                                                    } else {
                                                                                        if v35 {
                                                                                            v2 = (v54 - (v2 << (((^(v5))) & 63)))
                                                                                        }
                                                                                        var v84 bool = ((^(v36)) == v36)
                                                                                        _ = v84
                                                                                    }
                                                                                default:
                                                                                    if v35 {
                                                                                        v2 += v2
                                                                                        if (!(!v35)) {
                                                                                            v54 = int32(v36)
                                                                                            v1 = append(v1, (v52 & 175))
                                                                                            fmt.Println(v54, v69, v36)
                                                                                        } else {
                                                                                            v85 := GenS11{F0: v35, F1: 9223372036854775807, F2: int8(v5)}
                                                                                            var v86 int32 = 0
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                        }
                                                                                        v54 += 42
                                                                                        var v87 string = (v34 + v34)
                                                                                        _ = v87
                                                                                    } else {
                                                                                        var v88 int = 128
                                                                                        _ = v88
                                                                                    }
                                                                                }
                                                                                var v89 float32 = (1.5 / (0.1 / 0.1))
                                                                                _ = v83
                                                                                _ = v89
                                                                            }
                                                                            v90 := GenS12{F0: (v5 % ((v5) | 1)), F1: (v52 >> (((^(v5))) & 63))}
                                                                            _ = v90
                                                                        } else {
                                                                            var v91 uint16 = 65535
                                                                            v34 = "\n\t"
                                                                            v69 += (fi68 | (fi51 / ((fi68) | 1)))
                                                                            _ = v91
                                                                        }
                                                                        var v92 int16 = 32767
                                                                        _ = v92
                                                                    }
                                                                    v54 += (^(v54))
                                                                    fmt.Println(v3)
                                                                    v35 = v35
                                                                    _ = v69
                                                                } else {
                                                                    v36 = v36
                                                                    v52 += (^(v41))
                                                                    v35 = false
                                                                }
                                                                var v93 uint64 = 2
                                                                if ((("" + v34) + "a") >= v34) {
                                                                    if (v5 == (v5 - uint(v36))) {
                                                                        v94 := GenS5{F0: (v93 + v93), F1: (v93 >> ((v5) & 63))}
                                                                        fmt.Println(v5, v5, v2)
                                                                        var v95 int8 = int8(v2)
                                                                        _ = v94
                                                                        _ = v95
                                                                    } else {
                                                                        fmt.Println(v37, v37, v54)
                                                                        v96 := GenS6{F0: "\n\t", F1: v34}
                                                                        switch (v37 >> ((309) & 63)) {
                                                                        case 985:
                                                                            var v97 bool = (!true)
                                                                            _ = v97
                                                                        case 0:
                                                                            var v98 float32 = 0.1
                                                                            _ = v98
                                                                        case 9223372036854775807:
                                                                            v4[(v41 ^ 412)] = v37
                                                                        default:
                                                                            var v99 float64 = float64((0.0 + 3.14159))
                                                                            v100 := GenS13{F0: -1, F1: 1}
                                                                            _ = v100
                                                                            _ = v99
                                                                        }
                                                                        _ = v96
                                                                    }
                                                                    if (v35 || true) {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v101 := GenS11{F0: true, F1: v3, F2: -79}
                                                                        var v102 uint8 = uint8(v93)
                                                                        var v103 bool = false
                                                                        _ = v101
                                                                        _ = v102
                                                                        _ = v103
                                                                    } else {
                                                                        v104 := GenS4{F0: v34}
                                                                        _ = v104
                                                                    }
                                                                    var v105 string = (v34 + "foo bar")
                                                                    _ = v105
                                                                }
                                                                _ = fi68
                                                                _ = v93
                                                            }
                                                            if v35 {
                                                                v54 -= (^(v54))
                                                            } else {
                                                                fmt.Println(v52, v52, v36)
                                                                v2 -= -281
                                                                var v106 uint32 = v52
                                                                _ = v106
                                                            }
                                                        } else {
                                                            fmt.Println(v5)
                                                        }
                                                        if v35 {
                                                            for fi107 := 0; fi107 < 10; fi107++ {
                                                                if (((float32((-2.25 + float64(-2.25))) - 1e10) + 1.5) == -2.25) {
                                                                    switch (v37 * (v37 + (v37 - (-(v37))))) {
                                                                    case 4294967295:
                                                                        var v108 uint32 = 2
                                                                        _ = v108
                                                                    case 2147483647:
                                                                        fmt.Println(fi107)
                                                                    case 32767:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = v41
                                                                        }
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            if v35 {
                                                                                var v109 float64 = ((-2.25 * 1.5) + 3.14159)
                                                                                v54 = 2147483647
                                                                                _ = len(v4)
                                                                                _ = v109
                                                                            }
                                                                        } else {
                                                                            v34 = "foo bar"
                                                                            switch (v37 * (v37 / (((^(v37))) | 1))) {
                                                                            case -9223372036854775808:
                                                                                if (v35 && true) {
                                                                                    if (v2 >= v2) {
                                                                                        if (((0.0 * 0.1) - 1.5) <= 3.14159) {
                                                                                            _ = len(v4)
                                                                                            v41 = uint32(v36)
                                                                                        } else {
                                                                                            if v35 {
                                                                                                fmt.Println(v2, v54, v35)
                                                                                                v110 := GenS7{F0: 1, F1: 1}
                                                                                                _ = v110
                                                                                            }
                                                                                        }
                                                                                        var v111 int = 777
                                                                                        fmt.Println(v37)
                                                                                        _ = v111
                                                                                    } else {
                                                                                        var v112 int64 = -1
                                                                                        _ = v112
                                                                                    }
                                                                                }
                                                                            case 2147483647:
                                                                                v41 = 65535
                                                                            case 292:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (v52 ^ v41)
                                                                                }
                                                                            }
                                                                        }
                                                                        var v113 bool = v35
                                                                        v34 = (v34 + ("hello" + v34))
                                                                        switch v37 {
                                                                        case 979:
                                                                            v114 := GenS14{F0: "a"}
                                                                            _ = v114
                                                                        case 32767:
                                                                            var v115 bool = (!v35)
                                                                            _ = v115
                                                                        default:
                                                                            v116 := GenS5{F0: 42, F1: uint64(v5)}
                                                                            _ = v116
                                                                        }
                                                                        v117 := GenS15{F0: -61, F1: (v5 * (v5 % ((v5) | 1))), F2: ""}
                                                                        if len(v1) > 0 {
                                                                            v1[0] = v41
                                                                        }
                                                                        v118 := v4[287]
                                                                        v118 -= fi51
                                                                        if (((v34 + v34) + "a") > ("a" + "hello")) {
                                                                            fmt.Println(v54, fi51, v118)
                                                                            v35 = ((42 != 2) || (uint(v52) >= (v5 / ((v5) | 1))))
                                                                        }
                                                                        _ = v113
                                                                        _ = v117
                                                                        _ = v118
                                                                    } else {
                                                                        v5 -= 65535
                                                                        var v119 int8 = 46
                                                                        _ = v119
                                                                    }
                                                                    if v35 {
                                                                        v120 := GenS1{F0: 0, F1: (v35 && (100 != 100)), F2: 2}
                                                                        v121 := GenS16{F0: v3, F1: 0}
                                                                        fmt.Println(v36)
                                                                        if (!(!(!v35))) {
                                                                            v37 -= 128
                                                                            var v122 string = "a"
                                                                            v4[v41] = fi51
                                                                            v123 := GenS4{F0: (("" + v34) + "hello")}
                                                                            v54 = -1
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                var v124 uint8 = v36
                                                                                _ = v124
                                                                            case "\n\t":
                                                                                v41 = 42
                                                                            case "":
                                                                                var v125 string = v122
                                                                                _ = v125
                                                                            default:
                                                                                v122 = "\n\t"
                                                                            }
                                                                            _ = v122
                                                                            _ = v123
                                                                        } else {
                                                                            if (v35 || v35) {
                                                                                v35 = false
                                                                                _ = len(v4)
                                                                                if v35 {
                                                                                    if (127 <= -62) {
                                                                                        var v126 int32 = (v54 - (v2 % ((v54) | 1)))
                                                                                        _ = v126
                                                                                    }
                                                                                    v127 := v4[v41]
                                                                                    _ = v127
                                                                                } else {
                                                                                    if (1.5 >= (0.0 * 0.0)) {
                                                                                        fmt.Println(v35)
                                                                                        v36 += 0
                                                                                    }
                                                                                }
                                                                            }
                                                                            switch v34 {
                                                                            case "hello":
                                                                                var v128 float64 = 0.1
                                                                                _ = v128
                                                                            case "\n\t":
                                                                                if (42 > 1) {
                                                                                    v2 += (v54 | -2147483648)
                                                                                    fmt.Println(v3)
                                                                                    var v129 uint = 0
                                                                                    _ = v129
                                                                                } else {
                                                                                    var v130 float32 = 1e10
                                                                                    fmt.Println(v130, v5, v130)
                                                                                    if (("a" + v34) != ("foo bar" + (v34 + v34))) {
                                                                                        var v131 uint8 = (v36 / (((v36 * v36)) | 1))
                                                                                        var v132 uint8 = (v36 + (v36 - (v131 | v131)))
                                                                                        switch 32767 {
                                                                                        case 127:
                                                                                            if true {
                                                                                                v3 -= int64(v37)
                                                                                                fmt.Println(v34, v35, v132)
                                                                                                v133 := GenS12{F0: (v5 / (((^(v5))) | 1)), F1: v52}
                                                                                                _ = v133
                                                                                            } else {
                                                                                                var v134 uint64 = 0
                                                                                                _ = v134
                                                                                            }
                                                                                        }
                                                                                        v3 += v3
                                                                                        _ = v131
                                                                                        _ = v132
                                                                                    }
                                                                                    _ = v130
                                                                                }
                                                                            default:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 673
                                                                                }
                                                                            }
                                                                            v135 := GenS17{F0: 4294967295}
                                                                            _ = v135
                                                                        }
                                                                        var v136 string = v34
                                                                        _ = v120
                                                                        _ = v121
                                                                        _ = v136
                                                                    }
                                                                }
                                                                fmt.Println(v3, v2, fi107)
                                                                _ = fi107
                                                            }
                                                            v137 := GenS18{F0: v52, F1: v34}
                                                            _ = v137
                                                        } else {
                                                            switch -1 {
                                                            case 32767:
                                                                var v138 uint32 = 1
                                                                _ = v138
                                                            case -858:
                                                                for fi139 := 0; fi139 < 14; fi139++ {
                                                                    var v140 int64 = (v3 + v3)
                                                                    switch (fi51 ^ v37) {
                                                                    case -94:
                                                                        var v141 uint64 = 1
                                                                        _ = v141
                                                                    case -1:
                                                                        fmt.Println(v34, v37, v3)
                                                                    default:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2147483647
                                                                        }
                                                                    }
                                                                    v140 -= v3
                                                                    _ = fi139
                                                                    _ = v140
                                                                }
                                                            case 128:
                                                                v36 += v36
                                                            default:
                                                                switch fi51 {
                                                                case 2147483647:
                                                                    v142 := GenS12{F0: (v5 | (^(v5))), F1: (v41 | 4294967295)}
                                                                    _ = v142
                                                                case 32768:
                                                                    switch int(v3) {
                                                                    case 127:
                                                                        if (fi51 == (fi51 / ((32768) | 1))) {
                                                                            switch -1 {
                                                                            case 128:
                                                                                if v35 {
                                                                                    switch (v34 + ("a" + "\n\t")) {
                                                                                    case "":
                                                                                        v2 = v54
                                                                                    case "foo bar":
                                                                                        var v143 float32 = 1e10
                                                                                        _ = v143
                                                                                    default:
                                                                                        if (!true) {
                                                                                            fmt.Println(v52)
                                                                                            switch v34 {
                                                                                            case "a":
                                                                                                var v144 uint16 = 0
                                                                                                _ = v144
                                                                                            case "foo bar":
                                                                                                for fi145 := 0; fi145 < 1; fi145++ {
                                                                                                    var v146 uint8 = v36
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        v146 += v146
                                                                                                    case "foo bar":
                                                                                                        switch (v34 + (v34 + "foo bar")) {
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v37, fi51)
                                                                                                            switch int(v5) {
                                                                                                            case 42:
                                                                                                                v35 = v35
                                                                                                            case 4294967295:
                                                                                                                var v147 int8 = 42
                                                                                                                _ = v147
                                                                                                            case -9223372036854775808:
                                                                                                                v148 := GenS19{F0: "hello"}
                                                                                                                _ = v148
                                                                                                            default:
                                                                                                                _ = len(v1)
                                                                                                            }
                                                                                                            var v149 uint8 = (v36 % ((v146) | 1))
                                                                                                            _ = v149
                                                                                                        default:
                                                                                                            fmt.Println(v35, v35, fi145)
                                                                                                        }
                                                                                                    case "a":
                                                                                                        _ = len(v1)
                                                                                                    default:
                                                                                                        v54 += v2
                                                                                                        if (v35 && v35) {
                                                                                                            fmt.Println(v52)
                                                                                                            switch "hello" {
                                                                                                            case "":
                                                                                                                v150 := GenS1{F0: (v52 % ((v52) | 1)), F1: (1 > 0), F2: 2}
                                                                                                                _ = v150
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    var v151 uint64 = 0
                                                                                                                    v152 := v4[v52]
                                                                                                                    v153 := GenS20{F0: v34, F1: v54}
                                                                                                                    _ = v151
                                                                                                                    _ = v152
                                                                                                                    _ = v153
                                                                                                                } else {
                                                                                                                    switch ((v34 + (("a" + "hello") + "foo bar")) + v34) {
                                                                                                                    case "a":
                                                                                                                        if (v35 && (255 == 295)) {
                                                                                                                            _ = len(v1)
                                                                                                                            var v154 int16 = 32767
                                                                                                                            _ = v154
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch v34 {
                                                                                                                        case "\n\t":
                                                                                                                            _ = len(v4)
                                                                                                                        case "foo bar":
                                                                                                                            fmt.Println(v54)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v155 float64 = 3.14159
                                                                                                                    v156 := GenS8{F0: (v5 << ((v5) & 63)), F1: -27}
                                                                                                                    _ = len(v1)
                                                                                                                    var v157 uint16 = 143
                                                                                                                    v36 = v36
                                                                                                                    _ = v155
                                                                                                                    _ = v156
                                                                                                                    _ = v157
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v158 := GenS2{F0: 18446744073709551615, F1: 1}
                                                                                                        if v35 {
                                                                                                            if (!v35) {
                                                                                                                var v159 float32 = 3.14159
                                                                                                                _ = v159
                                                                                                            }
                                                                                                            delete(v4, 42)
                                                                                                        } else {
                                                                                                            if v35 {
                                                                                                                var v160 uint8 = 2
                                                                                                                v161 := v4[(v41 + v52)]
                                                                                                                v162 := GenS21{F0: -110, F1: (^(v146))}
                                                                                                                var v163 float32 = float32(((-2.25 * (1.5 + 0.1)) - -2.25))
                                                                                                                _ = v160
                                                                                                                _ = v161
                                                                                                                _ = v162
                                                                                                                _ = v163
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v158
                                                                                                    }
                                                                                                    var v164 int64 = v3
                                                                                                    _ = fi145
                                                                                                    _ = v146
                                                                                                    _ = v164
                                                                                                }
                                                                                            default:
                                                                                                for fi165 := 0; fi165 < 12; fi165++ {
                                                                                                    fmt.Println(fi51)
                                                                                                    _ = fi165
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v34, v3)
                                                                                        }
                                                                                    }
                                                                                    var v166 float64 = (-2.25 - float64(3.14159))
                                                                                    _ = v166
                                                                                } else {
                                                                                    fmt.Println(v3, v54, v36)
                                                                                    v167 := GenS10{F0: (v41 % ((v52) | 1)), F1: 4294967295}
                                                                                    if v35 {
                                                                                        for fi168 := 0; fi168 < 1; fi168++ {
                                                                                            var v169 string = v34
                                                                                            _ = fi168
                                                                                            _ = v169
                                                                                        }
                                                                                    } else {
                                                                                        v52 = 1
                                                                                        var v170 string = (v34 + "hello")
                                                                                        _ = v170
                                                                                    }
                                                                                    _ = v167
                                                                                }
                                                                            default:
                                                                                if v35 {
                                                                                    if v35 {
                                                                                        v171 := GenS22{F0: 509, F1: (v41 / ((v41) | 1)), F2: (-(v37))}
                                                                                        switch ("\n\t" + v34) {
                                                                                        case "foo bar":
                                                                                            var v172 uint32 = 989
                                                                                            _ = v172
                                                                                        case "":
                                                                                            v36 -= (v36 >> (((v5 << ((v5) & 63))) & 63))
                                                                                        default:
                                                                                            v41 -= v41
                                                                                        }
                                                                                        for fi173 := 0; fi173 < 16; fi173++ {
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                switch v34 {
                                                                                                case "\n\t":
                                                                                                    v5 += v5
                                                                                                }
                                                                                            case "":
                                                                                                fmt.Println(v36, v2, v54)
                                                                                            case "hello":
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    fmt.Println(fi173, fi173)
                                                                                                default:
                                                                                                    v174 := GenS23{F0: (v35 && (true && false)), F1: (v5 << (((v5 / ((v5) | 1))) & 63)), F2: 2}
                                                                                                    _ = v174
                                                                                                }
                                                                                            default:
                                                                                                v2 = 42
                                                                                            }
                                                                                            v175 := v4[(v41 >> (((v5 / (((v5 >> (((v5 >> ((v5) & 63))) & 63))) | 1))) & 63))]
                                                                                            _ = fi173
                                                                                            _ = v175
                                                                                        }
                                                                                        _ = v171
                                                                                    }
                                                                                } else {
                                                                                    if (-128 <= -128) {
                                                                                        v34 = ""
                                                                                        v3 = 1
                                                                                        for fi176 := 0; fi176 < 7; fi176++ {
                                                                                            v35 = v35
                                                                                            _ = fi176
                                                                                        }
                                                                                    } else {
                                                                                        v177 := GenS2{F0: (v5 + 65535), F1: -128}
                                                                                        v178 := v4[930]
                                                                                        _ = v177
                                                                                        _ = v178
                                                                                    }
                                                                                    switch v34 {
                                                                                    case "":
                                                                                        v36 -= v36
                                                                                    case "foo bar":
                                                                                        v179 := GenS21{F0: 1, F1: (v36 << (((v5 & (v5 ^ (v5 >> ((793) & 63))))) & 63))}
                                                                                        _ = v179
                                                                                    default:
                                                                                        for fi180 := 0; fi180 < 17; fi180++ {
                                                                                            var v181 int32 = (v54 | v2)
                                                                                            v37 = 850
                                                                                            _ = fi180
                                                                                            _ = v181
                                                                                        }
                                                                                    }
                                                                                    var v182 uint8 = v36
                                                                                    _ = v182
                                                                                }
                                                                            }
                                                                            v41 += 2147483647
                                                                        }
                                                                    default:
                                                                        var v183 uint = (^(v5))
                                                                        _ = v183
                                                                    }
                                                                }
                                                            }
                                                            v184 := GenS24{F0: v35, F1: v5, F2: uint32(v5)}
                                                            for fi185 := 0; fi185 < 11; fi185++ {
                                                                var v186 int64 = v3
                                                                v34 = "hello"
                                                                _ = fi185
                                                                _ = v186
                                                            }
                                                            _ = v184
                                                        }
                                                    }
                                                }
                                            } else {
                                                v35 = ((v2 % (((^(v54))) | 1)) == -607)
                                                switch (v34 + v34) {
                                                case "":
                                                    var v187 float32 = 0.1
                                                    _ = v187
                                                case "foo bar":
                                                    for fi188 := 0; fi188 < 12; fi188++ {
                                                        switch -751 {
                                                        case 32767:
                                                            var v189 uint32 = 2147483648
                                                            _ = v189
                                                        case -2147483648:
                                                            v54 += v2
                                                            switch (fi51 & -171) {
                                                            case 2147483647:
                                                                _ = len(v4)
                                                            case -2147483648:
                                                                delete(v4, (v52 | 492))
                                                            case 32768:
                                                                if (!((v36 * (v36 >> (((^(v5))) & 63))) == (v36 | uint8(v41)))) {
                                                                    _ = len(v4)
                                                                    if v35 {
                                                                        if v35 {
                                                                            _ = len(v1)
                                                                        } else {
                                                                            var v190 string = "a"
                                                                            var v191 uint16 = uint16(v36)
                                                                            v192 := GenS25{F0: 15, F1: 42}
                                                                            _ = v190
                                                                            _ = v191
                                                                            _ = v192
                                                                        }
                                                                        switch fi51 {
                                                                        case 127:
                                                                            v193 := GenS12{F0: (v5 >> ((v5) & 63)), F1: v41}
                                                                            if ((127 == (fi51 ^ (-(fi51)))) && v35) {
                                                                                v3 = 0
                                                                                delete(v4, (^(v41)))
                                                                            } else {
                                                                                var v194 uint8 = 255
                                                                                if (true || ((((v34 + v34) + v34) + "hello") != v34)) {
                                                                                    v36 += (^(v194))
                                                                                }
                                                                                _ = v194
                                                                            }
                                                                            v195 := GenS22{F0: (v5 + 65535), F1: (v41 << ((0) & 63)), F2: -1}
                                                                            _ = v193
                                                                            _ = v195
                                                                        case 32767:
                                                                            fmt.Println(v2)
                                                                        case 42:
                                                                            if (127 > 100) {
                                                                                v1 = append(v1, v52)
                                                                                v1 = append(v1, 65535)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v196 := GenS19{F0: "a"}
                                                                        if true {
                                                                            if v35 {
                                                                                v197 := GenS9{F0: (v5 >> ((uint(v36)) & 63)), F1: (1 != (v3 ^ (v3 << (((v5 / ((65535) | 1))) & 63))))}
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v1 = append(v1, (v41 * 768))
                                                                                _ = v197
                                                                            }
                                                                            var v198 int32 = (v54 - v2)
                                                                            _ = v198
                                                                        }
                                                                        v5 += v5
                                                                        _ = v196
                                                                    }
                                                                    v36 = v36
                                                                }
                                                            default:
                                                                switch v37 {
                                                                case -1:
                                                                    var v199 string = (v34 + v34)
                                                                    _ = v199
                                                                case 2147483647:
                                                                    v2 += (^(v54))
                                                                default:
                                                                    delete(v4, v52)
                                                                }
                                                            }
                                                            v4[v41] = v37
                                                        }
                                                        v2 += (v2 << (((^(v5))) & 63))
                                                        v200 := GenS12{F0: (v5 % (((v5 % ((0) | 1))) | 1)), F1: v41}
                                                        _ = fi188
                                                        _ = v200
                                                    }
                                                default:
                                                    var v201 int64 = (v3 ^ v3)
                                                    _ = v201
                                                }
                                            }
                                            v2 = v54
                                            _ = v54
                                        }
                                    case -2147483648:
                                        for fi202 := 0; fi202 < 16; fi202++ {
                                            v35 = v35
                                            var v203 uint16 = 609
                                            if v35 {
                                                var v204 uint16 = (v203 | v203)
                                                var v205 uint8 = (^(v36))
                                                _ = v204
                                                _ = v205
                                            }
                                            _ = fi202
                                            _ = v203
                                        }
                                    case 973:
                                        switch v34 {
                                        case "a":
                                            fmt.Println(v52, fi51)
                                        case "hello":
                                            v5 += v5
                                        }
                                    }
                                    _ = fi51
                                    _ = v52
                                }
                            case "":
                                for fi206 := 0; fi206 < 18; fi206++ {
                                    v2 = v2
                                    v207 := GenS13{F0: int16(v41), F1: (v2 % ((v2) | 1))}
                                    _ = fi206
                                    _ = v207
                                }
                            case "a":
                                v208 := GenS26{F0: (v37 >> (((v5 % ((v5) | 1))) & 63)), F1: 1}
                                _ = v208
                            }
                        } else {
                            for fi209 := 0; fi209 < 10; fi209++ {
                                for fi210 := 0; fi210 < 1; fi210++ {
                                    fmt.Println(fi209)
                                    var v211 int32 = (v2 ^ int32(fi209))
                                    if len(v1) > 0 {
                                        v1[0] = (^(v41))
                                    }
                                    _ = fi210
                                    _ = v211
                                }
                                _ = fi209
                            }
                            var v212 float64 = ((float64(3.14159) - ((-2.25 * 0.1) + float64(float32(((((3.14159 + 0.0) - 0.0) * 0.1) / (-2.25 - ((0.1 / 3.14159) + (0.0 / 1.5)))))))) + (1.5 / float64(1e10)))
                            v213 := GenS2{F0: v5, F1: 1}
                            _ = v212
                            _ = v213
                        }
                    }
                    _ = v41
                } else {
                    _ = len(v1)
                    v35 = (!(!(v35 && (float64((((float32(0.0) / 1.5) - 3.14159) - ((0.0 + (0.1 * -2.25)) - float32(0.1)))) > 1.5))))
                    fmt.Println(v36, v34, v5)
                }
            }
            _ = v37
        }
    }
    var v214 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v214:", v214)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v214
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
