package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int32; F1 int8; F2 uint64 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 string; F1 uint16; F2 uint16 }
type GenS12 struct { F0 bool; F1 bool; F2 string }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 bool; F1 uint8 }
type GenS15 struct { F0 uint8; F1 string }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int16; F1 int }
type GenS18 struct { F0 int64; F1 uint16 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 string; F1 bool; F2 uint64 }
type GenS24 struct { F0 int16 }
type GenS25 struct { F0 bool; F1 bool; F2 uint16 }
type GenS26 struct { F0 uint32 }
type GenS27 struct { F0 int32; F1 int }
type GenS28 struct { F0 uint16; F1 uint }
type GenS29 struct { F0 uint16 }
type GenS30 struct { F0 int; F1 uint64; F2 int64 }
func genFn0() int {
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            if false {
                var v6 int8 = -128
                switch (fi5 << ((4294967295) & 63)) {
                case 4294967295:
                    switch -2147483648 {
                    case -687:
                        var v7 int32 = 0
                        _ = v7
                    case -388:
                        v6 -= (v6 - 0)
                    case -9223372036854775808:
                        if true {
                            fmt.Println(v6)
                            if (!(true || (!(87 < 42)))) {
                                fmt.Println(v6, v6, v6)
                            }
                        }
                    }
                case 127:
                    fmt.Println(v6)
                default:
                    v6 = (v6 ^ v6)
                }
                _ = v6
            }
            v8 := []int{}
            _ = fi5
            _ = v8
        }
        fmt.Println(fi4)
        _ = fi4
    }
    var v9 uint64 = 0
    if false {
        if true {
            for fi10 := 0; fi10 < 18; fi10++ {
                v9 = (v9 % ((v9) | 1))
                for fi11 := 0; fi11 < 9; fi11++ {
                    if true {
                        switch ("" + ("a" + "\n\t")) {
                        case "a":
                            if true {
                                var v12 uint = 18446744073709551615
                                _ = v12
                            } else {
                                fmt.Println(fi10, v9)
                            }
                        case "foo bar":
                            var v13 bool = false
                            _ = v13
                        default:
                            var v14 uint = 255
                            _ = v14
                        }
                        var v15 int8 = 100
                        _ = v15
                    }
                    v16 := GenS1{F0: 0}
                    var v17 int8 = 2
                    _ = fi11
                    _ = v16
                    _ = v17
                }
                fmt.Println(fi10, v9)
                _ = fi10
            }
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float32 = float32(((0.0 + 1e10) / 1e10))
                v19 += v19
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint8 = 42
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
            v22 := []int{1, 4294967295}
            _ = v22
        }
        if true {
            if false {
                if ((v9 | (^(v9))) <= v9) {
                    fmt.Println(v9)
                }
                v23 := []int{9223372036854775807, 2147483647, -9223372036854775808}
                for fi24 := 0; fi24 < 19; fi24++ {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        fmt.Println(fi24)
                        v26 := GenS1{F0: 1}
                        var v27 uint16 = 1
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    var v28 int8 = -128
                    if (!(!false)) {
                        if ((270 <= 312) || false) {
                            if true {
                                for fi29 := 0; fi29 < 9; fi29++ {
                                    var v30 bool = false
                                    if (!(!false)) {
                                        fmt.Println(v9, fi24)
                                        if v30 {
                                            switch fi29 {
                                            case -9223372036854775808:
                                                v30 = true
                                            case 4294967295:
                                                v31 := GenS1{F0: 1}
                                                fmt.Println(fi24, v9, v28)
                                                if (!v30) {
                                                    v23 = append(v23, fi29)
                                                } else {
                                                    var v32 uint16 = 707
                                                    if v30 {
                                                        var v33 int32 = -2147483648
                                                        _ = v33
                                                    }
                                                    if v30 {
                                                        var v34 int64 = 1
                                                        _ = v34
                                                    }
                                                    switch (fi29 & int(v28)) {
                                                    case 127:
                                                        if (42 == 486) {
                                                            if v30 {
                                                                fmt.Println(fi29, v32, fi29)
                                                                v30 = v30
                                                            } else {
                                                                if (!v30) {
                                                                    if v30 {
                                                                        v9 = (^(v9))
                                                                    } else {
                                                                        v9 += (v9 | (v9 & 2))
                                                                        if ("hello" <= "a") {
                                                                            switch (fi24 << ((835) & 63)) {
                                                                            case -221:
                                                                                fmt.Println(v28)
                                                                            case -9223372036854775808:
                                                                                v28 += v28
                                                                            case 1:
                                                                                var v35 int8 = (v28 * v28)
                                                                                _ = v35
                                                                            }
                                                                        } else {
                                                                            var v36 uint64 = v9
                                                                            switch -788 {
                                                                            case -730:
                                                                                if (0 != 0) {
                                                                                    var v37 int32 = 2147483647
                                                                                    if (!(!(0.1 >= -2.25))) {
                                                                                        v28 += (v28 * (^(v28)))
                                                                                        var v38 uint = 2
                                                                                        _ = v38
                                                                                    } else {
                                                                                        if ((32767 < -32768) || (v28 < v28)) {
                                                                                            var v39 int64 = 0
                                                                                            if len(v23) >= 2 {
                                                                                                v23 = v23[1:len(v23)]
                                                                                            }
                                                                                            if v30 {
                                                                                                if true {
                                                                                                    v9 = 4294967295
                                                                                                    v36 -= (^(v9))
                                                                                                    switch fi24 {
                                                                                                    case 32768:
                                                                                                        if false {
                                                                                                            v40 := GenS1{F0: v39}
                                                                                                            if ((!(true && v30)) && (true || false)) {
                                                                                                                var v41 uint64 = (v9 << ((uint(v9)) & 63))
                                                                                                                _ = v41
                                                                                                            } else {
                                                                                                                var v42 string = "\n\t"
                                                                                                                var v43 uint32 = 485
                                                                                                                var v44 float32 = 0.0
                                                                                                                if v30 {
                                                                                                                    v45 := GenS2{F0: 15}
                                                                                                                    var v46 bool = (uint(v43) >= uint(v9))
                                                                                                                    _ = v45
                                                                                                                    _ = v46
                                                                                                                } else {
                                                                                                                    if v30 {
                                                                                                                        v28 += v28
                                                                                                                        if (!(int32(v28) > 50)) {
                                                                                                                            v23 = append(v23, (fi24 << ((4294967295) & 63)))
                                                                                                                            if (!false) {
                                                                                                                                v30 = v30
                                                                                                                            }
                                                                                                                        }
                                                                                                                        switch fi29 {
                                                                                                                        case -9223372036854775808:
                                                                                                                            switch ("hello" + v42) {
                                                                                                                            case "\n\t":
                                                                                                                                if (uint(v9) == 383) {
                                                                                                                                    v47 := GenS3{F0: v39, F1: v43}
                                                                                                                                    if (!false) {
                                                                                                                                        if ((v39 + (v39 + v39)) < int64(v37)) {
                                                                                                                                            fmt.Println(v30)
                                                                                                                                            v43 += (v43 ^ (^(v43)))
                                                                                                                                        } else {
                                                                                                                                            v42 = "\n\t"
                                                                                                                                            v48 := GenS2{F0: 127}
                                                                                                                                            var v49 bool = (2 >= (v32 / (((v32 ^ 255)) | 1)))
                                                                                                                                            _ = v48
                                                                                                                                            _ = v49
                                                                                                                                        }
                                                                                                                                        v42 = ("hello" + (v42 + v42))
                                                                                                                                        if (false && ((206 != 2) || ((v28 & (v28 & (v28 ^ (v28 * v28)))) == v28))) {
                                                                                                                                            var v50 uint16 = 32767
                                                                                                                                            v36 = v36
                                                                                                                                            _ = v50
                                                                                                                                        } else {
                                                                                                                                            var v51 bool = true
                                                                                                                                            _ = v51
                                                                                                                                        }
                                                                                                                                        if false {
                                                                                                                                            var v52 uint = 788
                                                                                                                                            _ = v52
                                                                                                                                        } else {
                                                                                                                                            v53 := GenS4{F0: (v37 % ((v37) | 1)), F1: v32}
                                                                                                                                            _ = v53
                                                                                                                                        }
                                                                                                                                        switch v42 {
                                                                                                                                        case "foo bar":
                                                                                                                                            var v54 uint32 = (v43 & v43)
                                                                                                                                            _ = v54
                                                                                                                                        default:
                                                                                                                                            fmt.Println(fi24, v42)
                                                                                                                                        }
                                                                                                                                        v39 -= 42
                                                                                                                                    } else {
                                                                                                                                        switch (v42 + (v42 + v42)) {
                                                                                                                                        case "foo bar":
                                                                                                                                            var v55 int64 = (v39 >> ((2) & 63))
                                                                                                                                            _ = v55
                                                                                                                                        }
                                                                                                                                        if len(v23) > 0 {
                                                                                                                                            v23[0] = 271
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if v30 {
                                                                                                                                        v39 += (v39 % (((-(v39))) | 1))
                                                                                                                                        var v56 int = (fi29 / ((fi29) | 1))
                                                                                                                                        _ = v56
                                                                                                                                    }
                                                                                                                                    v36 += 922
                                                                                                                                    v57 := GenS5{F0: (v37 | v37), F1: -1, F2: 4294967295}
                                                                                                                                    if (!(!(v30 || (false || v30)))) {
                                                                                                                                        v58 := GenS1{F0: (v39 & v39)}
                                                                                                                                        if (v37 <= v37) {
                                                                                                                                            var v59 int64 = 9223372036854775807
                                                                                                                                            if (!v30) {
                                                                                                                                                if true {
                                                                                                                                                    switch "hello" {
                                                                                                                                                    case "a":
                                                                                                                                                        if true {
                                                                                                                                                            fmt.Println(v28, fi29)
                                                                                                                                                            if (-32768 == -1) {
                                                                                                                                                                if (v59 == 56) {
                                                                                                                                                                    v23 = append(v23, (fi24 - -380))
                                                                                                                                                                    var v60 uint8 = 134
                                                                                                                                                                    _ = v60
                                                                                                                                                                }
                                                                                                                                                                v30 = (-1 > (^(v28)))
                                                                                                                                                                v61 := GenS6{F0: int64(v32), F1: v37}
                                                                                                                                                                _ = v61
                                                                                                                                                            } else {
                                                                                                                                                                if v30 {
                                                                                                                                                                    var v62 float64 = 0.1
                                                                                                                                                                    if len(v23) > 0 {
                                                                                                                                                                        v23[0] = -1
                                                                                                                                                                    }
                                                                                                                                                                    _ = v62
                                                                                                                                                                }
                                                                                                                                                                if len(v23) >= 2 {
                                                                                                                                                                    v23 = v23[1:len(v23)]
                                                                                                                                                                }
                                                                                                                                                                var v63 uint64 = v36
                                                                                                                                                                _ = v63
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        var v64 float32 = v44
                                                                                                                                                        if len(v23) >= 2 {
                                                                                                                                                            v23 = v23[1:len(v23)]
                                                                                                                                                        }
                                                                                                                                                        _ = v64
                                                                                                                                                    }
                                                                                                                                                    var v65 string = v42
                                                                                                                                                    _ = v65
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            _ = v59
                                                                                                                                        } else {
                                                                                                                                            var v66 string = v42
                                                                                                                                            switch ("foo bar" + "foo bar") {
                                                                                                                                            case "\n\t":
                                                                                                                                                var v67 uint16 = (^(v32))
                                                                                                                                                _ = v67
                                                                                                                                            case "":
                                                                                                                                                fmt.Println(v32, v32)
                                                                                                                                            case "a":
                                                                                                                                                var v68 uint = 1
                                                                                                                                                _ = v68
                                                                                                                                            }
                                                                                                                                            v69 := GenS5{F0: (-(v37)), F1: v28, F2: (^(v9))}
                                                                                                                                            fmt.Println(v43)
                                                                                                                                            var v70 uint8 = 42
                                                                                                                                            var v71 int16 = int16(v32)
                                                                                                                                            if true {
                                                                                                                                                if len(v23) > 0 {
                                                                                                                                                    v23[0] = (fi29 & 32768)
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                v32 = v32
                                                                                                                                                fmt.Println(v28, v28, v9)
                                                                                                                                            }
                                                                                                                                            _ = v66
                                                                                                                                            _ = v69
                                                                                                                                            _ = v70
                                                                                                                                            _ = v71
                                                                                                                                        }
                                                                                                                                        _ = v58
                                                                                                                                    }
                                                                                                                                    if true {
                                                                                                                                        switch ("" + "") {
                                                                                                                                        case "a":
                                                                                                                                            v32 = v32
                                                                                                                                        case "":
                                                                                                                                            if true {
                                                                                                                                                var v72 int = (fi24 + fi24)
                                                                                                                                                var v73 int16 = -1
                                                                                                                                                if (v73 != -1) {
                                                                                                                                                    v72 = (^(fi24))
                                                                                                                                                    v32 = v32
                                                                                                                                                    if v30 {
                                                                                                                                                        if v30 {
                                                                                                                                                            v39 -= v39
                                                                                                                                                            switch v42 {
                                                                                                                                                            case "a":
                                                                                                                                                                switch v72 {
                                                                                                                                                                case 0:
                                                                                                                                                                    var v74 uint8 = 2
                                                                                                                                                                    _ = v74
                                                                                                                                                                case -925:
                                                                                                                                                                    v73 = 32767
                                                                                                                                                                case -2147483648:
                                                                                                                                                                    var v75 float64 = -2.25
                                                                                                                                                                    _ = v75
                                                                                                                                                                }
                                                                                                                                                            case "foo bar":
                                                                                                                                                                var v76 uint64 = (v36 / ((v9) | 1))
                                                                                                                                                                _ = v76
                                                                                                                                                            case "hello":
                                                                                                                                                                var v77 int8 = v28
                                                                                                                                                                _ = v77
                                                                                                                                                            }
                                                                                                                                                            var v78 float32 = (v44 + 0.0)
                                                                                                                                                            _ = v78
                                                                                                                                                        }
                                                                                                                                                        _ = len(v23)
                                                                                                                                                        v30 = ((v37 == -287) || (1 < 2))
                                                                                                                                                    }
                                                                                                                                                    if (!true) {
                                                                                                                                                        v39 = v39
                                                                                                                                                        v72 -= v72
                                                                                                                                                        v30 = ((v73 >> ((42) & 63)) != v73)
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    switch 32767 {
                                                                                                                                                    case 18:
                                                                                                                                                        v79 := GenS7{F0: (v39 - (^(v39)))}
                                                                                                                                                        _ = v79
                                                                                                                                                    case -766:
                                                                                                                                                        v80 := GenS8{F0: 1, F1: v28}
                                                                                                                                                        _ = v80
                                                                                                                                                    default:
                                                                                                                                                        v43 += uint32(v32)
                                                                                                                                                    }
                                                                                                                                                    if len(v23) > 0 {
                                                                                                                                                        v23[0] = fi24
                                                                                                                                                    }
                                                                                                                                                    if v30 {
                                                                                                                                                        switch (^(fi29)) {
                                                                                                                                                        case 186:
                                                                                                                                                            v30 = true
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        v30 = v30
                                                                                                                                                        v44 = float32(((3.14159 - 1e10) + (0.0 + ((0.0 + -2.25) - (0.1 + 0.0)))))
                                                                                                                                                        fmt.Println(fi24)
                                                                                                                                                    }
                                                                                                                                                    if v30 {
                                                                                                                                                        if v30 {
                                                                                                                                                            if false {
                                                                                                                                                                fmt.Println(v36)
                                                                                                                                                            } else {
                                                                                                                                                                if (v37 != -1) {
                                                                                                                                                                    if (v43 > (^(v43))) {
                                                                                                                                                                        switch v72 {
                                                                                                                                                                        case 32767:
                                                                                                                                                                            if len(v23) >= 2 {
                                                                                                                                                                                v23 = v23[1:len(v23)]
                                                                                                                                                                            }
                                                                                                                                                                        case 9223372036854775807:
                                                                                                                                                                            v37 += v37
                                                                                                                                                                        case -490:
                                                                                                                                                                            var v81 bool = (("\n\t" + "hello") <= "foo bar")
                                                                                                                                                                            _ = v81
                                                                                                                                                                        default:
                                                                                                                                                                            v23 = append(v23, (fi29 ^ (fi29 & fi29)))
                                                                                                                                                                        }
                                                                                                                                                                        v82 := GenS9{F0: 127, F1: (v36 >> ((4294967295) & 63))}
                                                                                                                                                                        fmt.Println(v39, fi24)
                                                                                                                                                                        _ = v82
                                                                                                                                                                    } else {
                                                                                                                                                                        fmt.Println(v28, v72)
                                                                                                                                                                        if (!(78 != 46)) {
                                                                                                                                                                            v83 := GenS10{F0: (^(v9))}
                                                                                                                                                                            _ = v83
                                                                                                                                                                        }
                                                                                                                                                                        var v84 float64 = float64((v44 * float32(0.1)))
                                                                                                                                                                        v32 += (v32 + v32)
                                                                                                                                                                        _ = v84
                                                                                                                                                                    }
                                                                                                                                                                    if v30 {
                                                                                                                                                                        v30 = false
                                                                                                                                                                        v73 = v73
                                                                                                                                                                        _ = len(v23)
                                                                                                                                                                        v85 := GenS11{F0: (v42 + (v42 + v42)), F1: v32, F2: v32}
                                                                                                                                                                        var v86 int8 = v28
                                                                                                                                                                        var v87 bool = v30
                                                                                                                                                                        var v88 float32 = 0.0
                                                                                                                                                                        _ = v85
                                                                                                                                                                        _ = v86
                                                                                                                                                                        _ = v87
                                                                                                                                                                        _ = v88
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            var v89 int16 = int16(v43)
                                                                                                                                                            v89 = (v89 >> ((415) & 63))
                                                                                                                                                            _ = v89
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                _ = v72
                                                                                                                                                _ = v73
                                                                                                                                            }
                                                                                                                                        case "foo bar":
                                                                                                                                            if (int16(fi29) != -723) {
                                                                                                                                                v39 -= (^(v39))
                                                                                                                                                if (-2.25 <= float64(v44)) {
                                                                                                                                                    v90 := GenS12{F0: ((v37 | (^(v37))) >= (-(v37))), F1: v30, F2: (v42 + v42)}
                                                                                                                                                    _ = v90
                                                                                                                                                }
                                                                                                                                                fmt.Println(v39)
                                                                                                                                                v39 = v39
                                                                                                                                            } else {
                                                                                                                                                fmt.Println(v37, fi29)
                                                                                                                                                if v30 {
                                                                                                                                                    switch (fi24 + fi24) {
                                                                                                                                                    case -879:
                                                                                                                                                        var v91 int16 = 1000
                                                                                                                                                        _ = v91
                                                                                                                                                    case 9223372036854775807:
                                                                                                                                                        if v30 {
                                                                                                                                                            var v92 bool = false
                                                                                                                                                            fmt.Println(v9)
                                                                                                                                                            var v93 int = (-(fi29))
                                                                                                                                                            _ = v92
                                                                                                                                                            _ = v93
                                                                                                                                                        }
                                                                                                                                                    case 18:
                                                                                                                                                        if len(v23) >= 2 {
                                                                                                                                                            v23 = v23[1:len(v23)]
                                                                                                                                                        }
                                                                                                                                                    default:
                                                                                                                                                        var v94 uint32 = v43
                                                                                                                                                        _ = v94
                                                                                                                                                    }
                                                                                                                                                    if ((((^(v36)) <= 941) || v30) || v30) {
                                                                                                                                                        v36 += (v9 ^ v9)
                                                                                                                                                        if (!v30) {
                                                                                                                                                            fmt.Println(v30, v37, v32)
                                                                                                                                                            v28 = 127
                                                                                                                                                            v28 -= v28
                                                                                                                                                        }
                                                                                                                                                        v23 = append(v23, fi24)
                                                                                                                                                    }
                                                                                                                                                    v39 = v39
                                                                                                                                                }
                                                                                                                                                _ = len(v23)
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            switch fi29 {
                                                                                                                                            case -1:
                                                                                                                                                v23 = append(v23, (-(fi24)))
                                                                                                                                            case 635:
                                                                                                                                                v9 = 1
                                                                                                                                            case 9223372036854775807:
                                                                                                                                                switch v42 {
                                                                                                                                                case "hello":
                                                                                                                                                    fmt.Println(v9)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v47
                                                                                                                                    _ = v57
                                                                                                                                }
                                                                                                                                var v95 bool = v30
                                                                                                                                _ = v95
                                                                                                                            default:
                                                                                                                                v23 = append(v23, (-(fi24)))
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            if len(v23) > 0 {
                                                                                                                                v23[0] = 1
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v96 float32 = (float32(3.14159) - (v44 - 0.1))
                                                                                                                        v97 := GenS13{F0: 323}
                                                                                                                        _ = v96
                                                                                                                        _ = v97
                                                                                                                    }
                                                                                                                    fmt.Println(v30)
                                                                                                                    v44 = 0.1
                                                                                                                }
                                                                                                                _ = v42
                                                                                                                _ = v43
                                                                                                                _ = v44
                                                                                                            }
                                                                                                            _ = v40
                                                                                                        } else {
                                                                                                            switch (fi24 << ((42) & 63)) {
                                                                                                            case 32768:
                                                                                                                switch "" {
                                                                                                                case "hello":
                                                                                                                    v98 := GenS14{F0: (!v30), F1: 2}
                                                                                                                    _ = v98
                                                                                                                case "a":
                                                                                                                    if v30 {
                                                                                                                        v36 += 42
                                                                                                                        fmt.Println(v30, fi24)
                                                                                                                    } else {
                                                                                                                        v28 = v28
                                                                                                                        var v99 int8 = 42
                                                                                                                        _ = v99
                                                                                                                    }
                                                                                                                }
                                                                                                                var v100 float64 = (1e10 + 1e10)
                                                                                                                if (v37 != 42) {
                                                                                                                    if len(v23) > 0 {
                                                                                                                        v23[0] = int(v37)
                                                                                                                    }
                                                                                                                    if (4294967295 >= 127) {
                                                                                                                        var v101 float64 = (v100 - float64((-2.25 * 3.14159)))
                                                                                                                        v28 += (v28 << ((1) & 63))
                                                                                                                        _ = v101
                                                                                                                    }
                                                                                                                    v36 -= (v9 - (v36 | (v9 / ((42) | 1))))
                                                                                                                }
                                                                                                                _ = v100
                                                                                                            case 128:
                                                                                                                switch fi24 {
                                                                                                                case 0:
                                                                                                                    _ = len(v23)
                                                                                                                    var v102 int = (^(fi29))
                                                                                                                    var v103 uint = 1
                                                                                                                    _ = len(v23)
                                                                                                                    _ = len(v23)
                                                                                                                    _ = v102
                                                                                                                    _ = v103
                                                                                                                case 231:
                                                                                                                    var v104 int64 = (v39 | 1)
                                                                                                                    _ = v104
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    case -9223372036854775808:
                                                                                                        var v105 int = fi29
                                                                                                        _ = v105
                                                                                                    case 0:
                                                                                                        v36 = (^(v9))
                                                                                                    default:
                                                                                                        v106 := GenS15{F0: 127, F1: "\n\t"}
                                                                                                        _ = v106
                                                                                                    }
                                                                                                    fmt.Println(fi29)
                                                                                                    fmt.Println(v39, v28, v30)
                                                                                                    var v107 uint64 = 1
                                                                                                    _ = v107
                                                                                                }
                                                                                                if ((127 == (v28 * (-(v28)))) || (v30 || true)) {
                                                                                                    var v108 uint16 = (v32 << ((0) & 63))
                                                                                                    v9 = (v9 & (v36 % ((v9) | 1)))
                                                                                                    _ = v108
                                                                                                } else {
                                                                                                    if len(v23) > 0 {
                                                                                                        v23[0] = fi24
                                                                                                    }
                                                                                                    if v30 {
                                                                                                        if true {
                                                                                                            v109 := GenS5{F0: v37, F1: (v28 % ((127) | 1)), F2: 4294967295}
                                                                                                            var v110 int16 = -32768
                                                                                                            _ = v109
                                                                                                            _ = v110
                                                                                                        } else {
                                                                                                            v23 = append(v23, int(v28))
                                                                                                            v111 := GenS16{F0: 42}
                                                                                                            _ = v111
                                                                                                        }
                                                                                                    }
                                                                                                    if v30 {
                                                                                                        var v112 uint8 = 2
                                                                                                        var v113 int32 = -1
                                                                                                        _ = v112
                                                                                                        _ = v113
                                                                                                    }
                                                                                                    fmt.Println(v32)
                                                                                                    switch (fi24 / ((fi29) | 1)) {
                                                                                                    case -138:
                                                                                                        if len(v23) >= 2 {
                                                                                                            v23 = v23[1:len(v23)]
                                                                                                        }
                                                                                                    case -546:
                                                                                                        v36 += 724
                                                                                                        var v114 string = "foo bar"
                                                                                                        v30 = v30
                                                                                                        _ = v114
                                                                                                    case 32768:
                                                                                                        v115 := GenS14{F0: false, F1: 42}
                                                                                                        _ = v115
                                                                                                    default:
                                                                                                        fmt.Println(v37, v28, v32)
                                                                                                    }
                                                                                                }
                                                                                                if len(v23) >= 2 {
                                                                                                    v23 = v23[1:len(v23)]
                                                                                                }
                                                                                            } else {
                                                                                                var v116 float64 = 0.0
                                                                                                v39 = int64(fi24)
                                                                                                v117 := GenS17{F0: -107, F1: (fi29 + fi24)}
                                                                                                _ = v116
                                                                                                _ = v117
                                                                                            }
                                                                                            _ = v39
                                                                                        } else {
                                                                                            if true {
                                                                                                switch ("\n\t" + "") {
                                                                                                case "":
                                                                                                    if ((^(v9)) <= (v9 / ((v9) | 1))) {
                                                                                                        if (-1 != -9223372036854775808) {
                                                                                                            if true {
                                                                                                                v118 := GenS15{F0: 42, F1: "hello"}
                                                                                                                v119 := GenS18{F0: -9223372036854775808, F1: 760}
                                                                                                                _ = v118
                                                                                                                _ = v119
                                                                                                            }
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                v32 += (v32 & v32)
                                                                                                                v23 = append(v23, -279)
                                                                                                                var v120 int = fi29
                                                                                                                _ = v120
                                                                                                            } else {
                                                                                                                var v121 uint8 = 127
                                                                                                                _ = v121
                                                                                                            }
                                                                                                            if len(v23) >= 2 {
                                                                                                                v23 = v23[1:len(v23)]
                                                                                                            }
                                                                                                            if len(v23) >= 2 {
                                                                                                                v23 = v23[1:len(v23)]
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v122 := GenS11{F0: ("" + "a"), F1: v32, F2: v32}
                                                                                                    var v123 int16 = 32767
                                                                                                    _ = v122
                                                                                                    _ = v123
                                                                                                case "a":
                                                                                                    fmt.Println(fi29, fi29, v9)
                                                                                                    switch (fi24 + (-(fi29))) {
                                                                                                    case 1:
                                                                                                        v36 += v9
                                                                                                    default:
                                                                                                        var v124 int64 = 412
                                                                                                        _ = v124
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    if (199 <= 1) {
                                                                                                        v125 := GenS17{F0: -48, F1: -2147483648}
                                                                                                        v23 = append(v23, (fi29 + (fi24 << ((65535) & 63))))
                                                                                                        _ = v125
                                                                                                    } else {
                                                                                                        if len(v23) >= 2 {
                                                                                                            v23 = v23[1:len(v23)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v126 uint16 = v32
                                                                                                if true {
                                                                                                    v126 = (v32 & 255)
                                                                                                }
                                                                                                _ = v126
                                                                                            } else {
                                                                                                var v127 uint64 = (v36 * (v9 - v9))
                                                                                                _ = v127
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if ((v9 / (((v36 + (^(v9)))) | 1)) != (v9 + 540)) {
                                                                                        v23 = append(v23, (-(fi24)))
                                                                                        var v128 uint = 376
                                                                                        fmt.Println(v30)
                                                                                        _ = v128
                                                                                    }
                                                                                    _ = v37
                                                                                } else {
                                                                                    fmt.Println(v28, v32, fi29)
                                                                                }
                                                                            case 42:
                                                                                var v129 int = fi24
                                                                                _ = v129
                                                                            default:
                                                                                var v130 uint8 = 43
                                                                                _ = v130
                                                                            }
                                                                            _ = v36
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(v28, fi29)
                                                                }
                                                                if (int32(v9) < -1) {
                                                                    var v131 uint = 84
                                                                    _ = v131
                                                                } else {
                                                                    if v30 {
                                                                        var v132 uint32 = 2147483648
                                                                        if len(v23) >= 2 {
                                                                            v23 = v23[1:len(v23)]
                                                                        }
                                                                        _ = v132
                                                                    } else {
                                                                        v23 = append(v23, 809)
                                                                        var v133 float64 = 0.1
                                                                        _ = v133
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case -2147483648:
                                                        if v30 {
                                                            v134 := GenS19{F0: v9}
                                                            _ = v134
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                                v30 = ((-2.25 / 1.5) == 3.14159)
                                                _ = v31
                                            default:
                                                switch (-(fi29)) {
                                                case 1:
                                                    v30 = v30
                                                default:
                                                    v135 := GenS20{F0: (^(v9)), F1: 786, F2: (v9 ^ v9)}
                                                    _ = v135
                                                }
                                            }
                                        }
                                        if (-1 < 1000) {
                                            if (!(!v30)) {
                                                if v30 {
                                                    v9 -= (v9 | v9)
                                                }
                                                if (!v30) {
                                                    _ = len(v23)
                                                    fmt.Println(v30, v28)
                                                }
                                                v23 = append(v23, fi24)
                                            }
                                        } else {
                                            var v136 int8 = v28
                                            fmt.Println(v28)
                                            _ = v136
                                        }
                                    } else {
                                        _ = len(v23)
                                        if v30 {
                                            var v137 uint32 = 42
                                            _ = v137
                                        }
                                        var v138 int64 = -479
                                        var v139 uint = 42
                                        _ = v138
                                        _ = v139
                                    }
                                    v140 := GenS2{F0: 2}
                                    _ = fi29
                                    _ = v140
                                    _ = v30
                                }
                                for fi141 := 0; fi141 < 2; fi141++ {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    v28 -= (v28 << ((1) & 63))
                                    switch fi141 {
                                    case 42:
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                    case -475:
                                        if ((3.14159 <= 0.0) && (!(!false))) {
                                            v28 = v28
                                        }
                                    case -706:
                                        if len(v23) > 0 {
                                            v23[0] = (fi24 | (fi141 << ((98) & 63)))
                                        }
                                    }
                                    _ = fi141
                                }
                            }
                        } else {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            if true {
                                var v142 int32 = -2147483648
                                v28 += 94
                                _ = v142
                            } else {
                                var v143 uint = 4294967295
                                _ = v143
                            }
                        }
                    } else {
                        var v144 int8 = -128
                        var v145 uint8 = 128
                        v28 += v144
                        _ = v144
                        _ = v145
                    }
                    _ = fi24
                    _ = v28
                }
                _ = v23
            }
            v9 = (v9 & 0)
            v146 := []int{4294967295, -312, -9223372036854775808, -19}
            _ = v146
        } else {
            v147 := GenS18{F0: 0, F1: 65535}
            _ = v147
        }
    } else {
        switch int(v9) {
        case -2147483648:
            var v148 uint64 = (^(v9))
            _ = v148
        case 756:
            if false {
                if (!(((float64(1.5) + 1.5) * 3.14159) <= 0.0)) {
                    for fi149 := 0; fi149 < 3; fi149++ {
                        if (-2147483648 <= -920) {
                            if ((("" + "\n\t") + "\n\t") == "hello") {
                                if (127 < 98) {
                                    var v150 bool = false
                                    v151 := []int{}
                                    for fi152 := 0; fi152 < 16; fi152++ {
                                        var v153 uint32 = 2147483647
                                        v9 -= (v9 | (^(v9)))
                                        var v154 float64 = float64((3.14159 * 0.1))
                                        _ = fi152
                                        _ = v153
                                        _ = v154
                                    }
                                    _ = v150
                                    _ = v151
                                }
                                var v155 float64 = 3.14159
                                for fi156 := 0; fi156 < 10; fi156++ {
                                    v157 := GenS21{F0: (!false)}
                                    v158 := []int{fi156, fi156, fi156}
                                    _ = fi156
                                    _ = v157
                                    _ = v158
                                }
                                _ = v155
                            } else {
                                for fi159 := 0; fi159 < 11; fi159++ {
                                    var v160 int32 = 1
                                    var v161 int16 = 1
                                    var v162 string = ("foo bar" + "")
                                    _ = fi159
                                    _ = v160
                                    _ = v161
                                    _ = v162
                                }
                            }
                        }
                        _ = fi149
                    }
                    v9 = (v9 ^ 42)
                }
            } else {
                if false {
                    if false {
                        var v163 int64 = -1
                        fmt.Println(v163, v9)
                        _ = v163
                    } else {
                        var v164 float64 = 0.1
                        _ = v164
                    }
                    fmt.Println(v9)
                }
            }
        case -652:
            v165 := []int{}
            _ = v165
        }
        if false {
            switch -534 {
            case -2147483648:
                v166 := GenS19{F0: v9}
                _ = v166
            case -1:
                if false {
                    var v167 int = 9223372036854775807
                    for fi168 := 0; fi168 < 15; fi168++ {
                        if false {
                            if (-1 < 0) {
                                var v169 float64 = ((1e10 * 0.1) - ((1e10 + 1.5) + 3.14159))
                                v170 := GenS11{F0: "hello", F1: 709, F2: 0}
                                v171 := []int{v167, (v167 + v167)}
                                _ = v169
                                _ = v170
                                _ = v171
                            }
                            v167 -= fi168
                            v172 := GenS9{F0: 42, F1: (v9 << ((18446744073709551615) & 63))}
                            _ = v172
                        } else {
                            fmt.Println(v9, v9, fi168)
                            v173 := []int{v167}
                            var v174 uint8 = 127
                            _ = v173
                            _ = v174
                        }
                        switch fi168 {
                        case 1:
                            v9 += 18446744073709551615
                        case 32767:
                            for fi175 := 0; fi175 < 7; fi175++ {
                                var v176 uint = 18446744073709551615
                                fmt.Println(v167, v9, v167)
                                _ = fi175
                                _ = v176
                            }
                        }
                        _ = fi168
                    }
                    v177 := GenS22{F0: 420}
                    _ = v167
                    _ = v177
                } else {
                    v178 := GenS21{F0: (true || false)}
                    var v179 uint64 = (v9 >> ((18446744073709551615) & 63))
                    v9 = v179
                    _ = v178
                    _ = v179
                }
            }
            if (!(false && false)) {
                fmt.Println(v9)
                v180 := []int{324}
                _ = v180
            }
            var v181 bool = true
            _ = v181
        } else {
            var v182 uint16 = 0
            if (!(false || (-1 > -2147483648))) {
                if (-788 <= -2147483648) {
                    fmt.Println(v9)
                } else {
                    v9 -= (^(v9))
                }
                v183 := GenS23{F0: "", F1: (!true), F2: (^(v9))}
                _ = v183
            } else {
                for fi184 := 0; fi184 < 16; fi184++ {
                    v185 := []int{(fi184 >> ((255) & 63)), -942, (fi184 ^ fi184), 32767}
                    _ = len(v185)
                    _ = fi184
                    _ = v185
                }
                fmt.Println(v9)
                for fi186 := 0; fi186 < 14; fi186++ {
                    v187 := GenS17{F0: -296, F1: fi186}
                    _ = fi186
                    _ = v187
                }
            }
            v9 = (v9 - 0)
            _ = v182
        }
        var v188 int64 = 42
        _ = v188
    }
    v189 := GenS18{F0: 9223372036854775807, F1: 42}
    _ = v189
    _ = v9
    return 128
}

func main() {
    v1 := []string{"hello", "foo bar"}
    var v2 bool = false
    var v3 uint64 = 18446744073709551615
    var v190 int16 = 618
    fmt.Println(v2, v190)
    switch (("a" + "foo bar") + "") {
    case "foo bar":
        if ((((v2 || v2) && v2) || (v2 && (!(!v2)))) || v2) {
            for fi191 := 0; fi191 < 6; fi191++ {
                switch (fi191 << ((255) & 63)) {
                case -1:
                    if (v2 && (!(!(!v2)))) {
                        if v2 {
                            for fi192 := 0; fi192 < 6; fi192++ {
                                if (!v2) {
                                    switch 32768 {
                                    case 0:
                                        if (2 > 546) {
                                            fmt.Println(v2, fi192, fi192)
                                            v190 -= (v190 - (^(v190)))
                                            var v193 uint8 = 127
                                            _ = v193
                                        } else {
                                            var v194 int64 = 1
                                            var v195 uint16 = 0
                                            switch "\n\t" {
                                            case "foo bar":
                                                v2 = (183 >= uint8(v3))
                                            case "a":
                                                _ = len(v1)
                                            case "hello":
                                                v196 := GenS21{F0: ((v195 / (((v195 / ((v195) | 1))) | 1)) == 760)}
                                                _ = v196
                                            default:
                                                fmt.Println(v3, fi191, v3)
                                            }
                                            _ = v194
                                            _ = v195
                                        }
                                    case 2147483647:
                                        v2 = v2
                                    case 33:
                                        fmt.Println(fi192)
                                        fmt.Println(fi192, fi191, v3)
                                        if false {
                                            v3 -= (v3 | (v3 | 984))
                                        } else {
                                            v197 := GenS7{F0: -625}
                                            var v198 int8 = 127
                                            if false {
                                                if v2 {
                                                    v199 := GenS24{F0: int16(fi192)}
                                                    var v200 int32 = 42
                                                    if true {
                                                        var v201 float64 = (-2.25 - -2.25)
                                                        switch (fi192 * (fi191 + genFn0())) {
                                                        case 4294967295:
                                                            fmt.Println(v3)
                                                        case -1:
                                                            if ((v2 || v2) && ((!(v2 && v2)) && v2)) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = "hello"
                                                                }
                                                                v202 := GenS9{F0: 128, F1: (^(v3))}
                                                                if true {
                                                                    v203 := GenS25{F0: (((^(v3)) >= v3) || ((v198 | v198) <= (v198 - v198))), F1: v2, F2: uint16(v3)}
                                                                    switch (fi191 >> ((65535) & 63)) {
                                                                    case 1:
                                                                        if (3.14159 != 3.14159) {
                                                                            var v204 int16 = int16(v200)
                                                                            v190 = v190
                                                                            var v205 float64 = -2.25
                                                                            _ = v204
                                                                            _ = v205
                                                                        }
                                                                    case 127:
                                                                        v190 -= (v190 + v190)
                                                                    }
                                                                    _ = v203
                                                                }
                                                                _ = v202
                                                            } else {
                                                                v200 -= -473
                                                            }
                                                        default:
                                                            fmt.Println(v190)
                                                            v201 = (v201 / (v201 - 0.0))
                                                            v206 := GenS16{F0: (v190 / ((v190) | 1))}
                                                            _ = v206
                                                        }
                                                        if false {
                                                            v207 := GenS1{F0: -1}
                                                            _ = v207
                                                        }
                                                        _ = v201
                                                    } else {
                                                        v1 = append(v1, "foo bar")
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                    var v208 uint = 2
                                                    _ = v199
                                                    _ = v200
                                                    _ = v208
                                                } else {
                                                    fmt.Println(v190)
                                                    v209 := GenS22{F0: v190}
                                                    fmt.Println(fi192)
                                                    _ = v209
                                                }
                                                v3 -= v3
                                            }
                                            var v210 uint16 = 0
                                            _ = v197
                                            _ = v198
                                            _ = v210
                                        }
                                        switch genFn0() {
                                        case 12:
                                            switch "" {
                                            case "a":
                                                if v2 {
                                                    v1 = append(v1, ("foo bar" + ""))
                                                    fmt.Println(v190)
                                                    switch fi192 {
                                                    case 32768:
                                                        if len(v1) > 0 {
                                                            v1[0] = "hello"
                                                        }
                                                    default:
                                                        switch (^(fi192)) {
                                                        case -308:
                                                            v190 = v190
                                                        case 42:
                                                            var v211 bool = v2
                                                            _ = v211
                                                        default:
                                                            var v212 int32 = -2147483648
                                                            _ = v212
                                                        }
                                                    }
                                                    v213 := GenS8{F0: 4294967295, F1: int8(v3)}
                                                    _ = v213
                                                }
                                            case "hello":
                                                v3 = (v3 % ((v3) | 1))
                                            default:
                                                switch ("\n\t" + "") {
                                                case "foo bar":
                                                    var v214 int16 = v190
                                                    _ = v214
                                                }
                                            }
                                        case 32768:
                                            var v215 int = genFn0()
                                            _ = v215
                                        }
                                    default:
                                        var v216 uint32 = 2147483647
                                        _ = v216
                                    }
                                }
                                v217 := GenS26{F0: 601}
                                fmt.Println(v2, v2, v190)
                                if true {
                                    v1 = append(v1, "foo bar")
                                    fmt.Println(fi191, v2)
                                    var v218 float32 = 1e10
                                    _ = v218
                                }
                                _ = fi192
                                _ = v217
                            }
                        }
                    }
                }
                fmt.Println(v2, fi191)
                _ = fi191
            }
            v2 = v2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            for fi219 := 0; fi219 < 15; fi219++ {
                for fi220 := 0; fi220 < 14; fi220++ {
                    var v221 float64 = (1e10 / 0.1)
                    var v222 int8 = -1
                    _ = fi220
                    _ = v221
                    _ = v222
                }
                _ = fi219
            }
        }
    case "hello":
        var v223 int = 4294967295
        _ = v223
    case "\n\t":
        var v224 int64 = 0
        _ = v224
    default:
        fmt.Println(v3, v190)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v2, v2)
    var v225 int8 = 42
    for fi226 := 0; fi226 < 17; fi226++ {
        for fi227 := 0; fi227 < 1; fi227++ {
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            v228 := GenS24{F0: int16(fi227)}
            var v229 int16 = v190
            _ = fi227
            _ = v228
            _ = v229
        }
        v2 = (!(v190 < int16(fi226)))
        _ = fi226
    }
    for fi230 := 0; fi230 < 4; fi230++ {
        var v231 uint = 65535
        v232 := GenS8{F0: (v231 - 285), F1: (v225 >> ((v231) & 63))}
        _ = fi230
        _ = v231
        _ = v232
    }
    if v2 {
        v1 = append(v1, "")
    } else {
        switch 1 {
        case 42:
            switch "" {
            case "\n\t":
                v190 = (v190 ^ v190)
            case "hello":
                var v233 int16 = -745
                _ = v233
            default:
                v2 = (true || (v2 && ((!false) || (v2 || (!v2)))))
            }
        case -896:
            var v234 bool = (255 == 255)
            _ = v234
        case 0:
            v235 := GenS23{F0: "\n\t", F1: ((true || (!(0 < (v190 & 1)))) || (v225 < (v225 & -1))), F2: v3}
            _ = v235
        default:
            for fi236 := 0; fi236 < 20; fi236++ {
                if v2 {
                    v225 = v225
                    if ((((955 != 42) || v2) && v2) || (!(0 < -32768))) {
                        switch (^(fi236)) {
                        case 4294967295:
                            var v237 int32 = 42
                            _ = v237
                        }
                        v225 = 1
                        v238 := GenS13{F0: -1}
                        _ = v238
                    } else {
                        var v239 uint16 = 0
                        _ = v239
                    }
                    for fi240 := 0; fi240 < 17; fi240++ {
                        var v241 int64 = -9223372036854775808
                        if v2 {
                            v241 = (v241 * v241)
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            var v242 int16 = (v190 * (v190 % (((v190 & 32767)) | 1)))
                            switch "a" {
                            case "\n\t":
                                v1 = append(v1, "hello")
                            }
                            _ = v242
                        }
                        v243 := GenS19{F0: v3}
                        v190 = v190
                        _ = fi240
                        _ = v241
                        _ = v243
                    }
                } else {
                    switch (fi236 - (-(fi236))) {
                    case 138:
                        var v244 uint = 302
                        _ = v244
                    case -762:
                        v245 := GenS27{F0: -2147483648, F1: (^(fi236))}
                        _ = v245
                    case 880:
                        v246 := GenS10{F0: (v3 - (v3 | (v3 ^ v3)))}
                        _ = v246
                    }
                }
                v190 = 1000
                v247 := GenS28{F0: 471, F1: 2}
                _ = fi236
                _ = v247
            }
        }
    }
    for fi248 := 0; fi248 < 15; fi248++ {
        v2 = v2
        var v249 int16 = -1
        v190 += (v249 + v190)
        _ = fi248
        _ = v249
    }
    if (-2147483648 == -561) {
        var v250 uint64 = (^(v3))
        var v251 uint8 = 124
        v252 := GenS29{F0: 65535}
        _ = v250
        _ = v251
        _ = v252
    }
    v190 -= (v190 >> ((18446744073709551615) & 63))
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if v2 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v253 uint8 = 129
        for fi254 := 0; fi254 < 11; fi254++ {
            var v255 int = fi254
            for fi256 := 0; fi256 < 3; fi256++ {
                var v257 uint8 = v253
                if (("" + "") <= "") {
                    if (v2 && v2) {
                        if len(v1) > 0 {
                            v1[0] = ("hello" + "")
                        }
                    } else {
                        fmt.Println(v253)
                    }
                    if (-79 <= 0) {
                        v258 := GenS27{F0: 2147483647, F1: v255}
                        v2 = (32767 < genFn0())
                        v255 = 2147483647
                        _ = v258
                    } else {
                        var v259 uint = 2
                        if false {
                            v260 := GenS4{F0: 0, F1: 255}
                            var v261 float64 = 3.14159
                            v225 = v225
                            _ = v260
                            _ = v261
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v3 = uint64(v253)
                        _ = v259
                    }
                } else {
                    var v262 int8 = v225
                    if v2 {
                        switch "hello" {
                        case "hello":
                            switch "foo bar" {
                            case "foo bar":
                                var v263 uint16 = 0
                                _ = v263
                            }
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "":
                            var v264 uint16 = 1
                            _ = v264
                        default:
                            if ((1 <= 725) && (v2 && v2)) {
                                var v265 float32 = 1.5
                                _ = v265
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                        v1 = append(v1, ("a" + ("\n\t" + (((("foo bar" + "") + "hello") + "foo bar") + "foo bar"))))
                        var v266 int32 = 331
                        _ = v266
                    } else {
                        if (-2147483648 > -2147483648) {
                            var v267 uint32 = 434
                            v3 -= 76
                            fmt.Println(v257, v257)
                            switch (fi254 & fi256) {
                            case -2147483648:
                                var v268 float32 = 0.0
                                _ = v268
                            }
                            v269 := GenS10{F0: 4294967295}
                            _ = v267
                            _ = v269
                        } else {
                            if false {
                                var v270 int16 = v190
                                v257 -= v253
                                _ = v270
                            }
                        }
                        fmt.Println(v253)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if v2 {
                        var v271 float64 = (0.0 / 1.0)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v262)
                        _ = v271
                    } else {
                        fmt.Println(v253, v262, v255)
                        _ = len(v1)
                    }
                    _ = v262
                }
                if ((-(v225)) >= v225) {
                    v190 += v190
                    v272 := GenS25{F0: (1e10 >= 3.14159), F1: v2, F2: 2}
                    _ = v272
                }
                _ = fi256
                _ = v257
            }
            _ = fi254
            _ = v255
        }
        _ = v253
    } else {
        v225 = v225
        for fi273 := 0; fi273 < 20; fi273++ {
            v190 = (v190 ^ v190)
            v274 := GenS2{F0: 0}
            _ = fi273
            _ = v274
        }
    }
    var v275 bool = v2
    fmt.Println(v3, v225, v225)
    if v2 {
        var v276 string = "foo bar"
        v3 = 2
        for fi277 := 0; fi277 < 10; fi277++ {
            var v278 int16 = (v190 << ((2) & 63))
            _ = fi277
            _ = v278
        }
        _ = v276
    } else {
        var v279 float64 = ((0.0 / 1e10) + ((1e10 * 0.0) + 0.1))
        for fi280 := 0; fi280 < 4; fi280++ {
            for fi281 := 0; fi281 < 10; fi281++ {
                if v275 {
                    v1 = append(v1, (("hello" + (("foo bar" + "\n\t") + "foo bar")) + "foo bar"))
                    var v282 uint64 = (v3 & 764)
                    v3 -= v3
                    v1 = append(v1, (("foo bar" + "foo bar") + "a"))
                    _ = v282
                }
                _ = fi281
            }
            v283 := GenS30{F0: 128, F1: (^(v3)), F2: 42}
            var v284 string = "hello"
            _ = fi280
            _ = v283
            _ = v284
        }
        _ = v279
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v190:", v190)
    fmt.Println("v2:", v2)
    fmt.Println("v225:", v225)
    fmt.Println("v275:", v275)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v190
    _ = v2
    _ = v225
    _ = v275
    _ = v3
}
