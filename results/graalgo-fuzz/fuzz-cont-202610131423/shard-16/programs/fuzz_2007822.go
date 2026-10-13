package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS4 struct { F0 int64; F1 uint; F2 int8 }
type GenS5 struct { F0 int16; F1 int16 }

func main() {
    v1 := []uint32{42, 1}
    v2 := make(map[uint32]int)
    var v3 int = -760
    var v4 int32 = 0
    var v5 uint64 = 698
    fmt.Println(v5)
    var v6 uint32 = 42
    if ("" > "foo bar") {
        var v7 uint = 861
        if true {
            var v8 uint8 = 42
            var v9 uint16 = 42
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: v5, F1: 42, F2: ((false && (!false)) && true)}
            v11 := GenS1{F0: v5, F1: 1, F2: (65535 == 0)}
            for fi12 := 0; fi12 < 10; fi12++ {
                v4 = v4
                var v13 int8 = 100
                v14 := GenS1{F0: 802, F1: int16(v4), F2: ((!(v13 <= (v13 << (((v7 * (v7 | v7))) & 63)))) || (0.0 == 0.1))}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v10
            _ = v11
        }
        _ = v7
    }
    v3 = (v3 - 128)
    if (!true) {
        if (0 > -9223372036854775808) {
            if false {
                var v15 int64 = int64(v5)
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 string = "foo bar"
                    _ = fi16
                    _ = v17
                }
                var v18 float64 = (1.5 + 1e10)
                _ = v15
                _ = v18
            }
            if (true || (127 >= 0)) {
                v4 = v4
            }
            switch "hello" {
            case "foo bar":
                v19 := GenS1{F0: 18446744073709551615, F1: 42, F2: true}
                _ = v19
            case "\n\t":
                v6 += 4294967295
            case "hello":
                _ = len(v2)
            default:
                var v20 int8 = 0
                _ = v20
            }
        }
        if true {
            if (128 <= 0) {
                var v21 int16 = -32768
                v3 += (v3 | v3)
                for fi22 := 0; fi22 < 17; fi22++ {
                    _ = len(v2)
                    _ = fi22
                }
                _ = v21
            }
            if (1 < 0) {
                v5 -= (v5 + v5)
            } else {
                fmt.Println(v4, v3, v5)
                fmt.Println(v6)
                if false {
                    _ = len(v2)
                    switch "hello" {
                    case "":
                        var v23 float32 = float32((1e10 + 1.5))
                        _ = v23
                    case "a":
                        var v24 string = "hello"
                        _ = v24
                    default:
                        for fi25 := 0; fi25 < 18; fi25++ {
                            var v26 uint8 = 2
                            v3 = (v3 % (((^(fi25))) | 1))
                            if false {
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    if ((false && false) && false) {
                                        var v28 int8 = 42
                                        _ = v28
                                    }
                                    if ((v26 + v26) != uint8(v5)) {
                                        if true {
                                            if (65535 == 2) {
                                                fmt.Println(fi27, fi27)
                                                v4 = v4
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                        var v29 int64 = -401
                                        v26 = 42
                                        v4 -= 42
                                        _ = v29
                                    }
                                    _ = fi27
                                }
                                switch int(v5) {
                                case -234:
                                    fmt.Println(v4)
                                case 32768:
                                    v30 := GenS2{F0: v26, F1: 0}
                                    _ = v30
                                case -2147483648:
                                    _ = len(v2)
                                }
                                if (!(!true)) {
                                    switch fi25 {
                                    case 23:
                                        for fi31 := 0; fi31 < 13; fi31++ {
                                            v32 := v2[(v6 - (v6 % (((v6 % (((v6 << ((1) & 63))) | 1))) | 1)))]
                                            var v33 uint64 = v5
                                            _ = fi31
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                    var v34 int = (v3 << ((2) & 63))
                                    v26 = 208
                                    _ = v34
                                } else {
                                    var v35 int8 = -102
                                    _ = v35
                                }
                            } else {
                                v36 := GenS1{F0: (v5 << ((1) & 63)), F1: 1, F2: true}
                                switch "" {
                                case "foo bar":
                                    var v37 float32 = float32(1e10)
                                    _ = v37
                                case "hello":
                                    switch "a" {
                                    case "\n\t":
                                        var v38 int16 = -32768
                                        _ = v38
                                    case "hello":
                                        var v39 float32 = 1e10
                                        _ = v39
                                    case "foo bar":
                                        v3 += v3
                                    default:
                                        var v40 bool = (v3 < (fi25 ^ (v3 * 4294967295)))
                                        _ = v40
                                    }
                                }
                                _ = v36
                            }
                            _ = fi25
                            _ = v26
                        }
                    }
                    for fi41 := 0; fi41 < 5; fi41++ {
                        for fi42 := 0; fi42 < 12; fi42++ {
                            var v43 uint16 = 45
                            var v44 float32 = 0.1
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 11; fi45++ {
                            fmt.Println(fi41, fi45, fi45)
                            if true {
                                var v46 float32 = (-2.25 * 0.1)
                                v1 = append(v1, v6)
                                var v47 uint64 = (^(v5))
                                _ = v46
                                _ = v47
                            } else {
                                var v48 int32 = (-(v4))
                                if ((v4 / ((v48) | 1)) >= 42) {
                                    v49 := GenS3{F0: (v5 / (((^(v5))) | 1)), F1: v6, F2: -128}
                                    var v50 bool = false
                                    var v51 int = -9223372036854775808
                                    var v52 uint64 = v5
                                    v53 := v2[1]
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                if len(v1) > 0 {
                                    v1[0] = v6
                                }
                                _ = v48
                            }
                            if true {
                                v54 := GenS3{F0: 1, F1: 0, F2: -1}
                                var v55 uint8 = 127
                                var v56 bool = false
                                v6 -= v6
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                            _ = fi45
                        }
                        _ = fi41
                    }
                } else {
                    v3 += v3
                    for fi57 := 0; fi57 < 4; fi57++ {
                        v4 -= (v4 - (v4 % ((-1) | 1)))
                        var v58 uint64 = (^(v5))
                        if (99 <= 0) {
                            var v59 float64 = float64(0.1)
                            _ = v59
                        } else {
                            for fi60 := 0; fi60 < 14; fi60++ {
                                v61 := GenS1{F0: (v58 >> ((662) & 63)), F1: 32767, F2: ((301 < 955) && (953 != 255))}
                                v58 += v5
                                v3 = fi57
                                _ = fi60
                                _ = v61
                            }
                        }
                        _ = fi57
                        _ = v58
                    }
                }
            }
            var v62 int = -141
            _ = v62
        }
        if true {
            _ = len(v1)
        } else {
            switch ("" + "\n\t") {
            case "foo bar":
                if true {
                    fmt.Println(v3, v4)
                }
            case "a":
                fmt.Println(v3)
            }
            for fi63 := 0; fi63 < 16; fi63++ {
                if false {
                    var v64 uint = 0
                    _ = v64
                }
                for fi65 := 0; fi65 < 18; fi65++ {
                    switch (^(fi65)) {
                    case 9223372036854775807:
                        if true {
                            if true {
                                switch fi65 {
                                case -9223372036854775808:
                                    var v66 int64 = 42
                                    _ = v66
                                default:
                                    if true {
                                        var v67 int = (fi65 >> ((624) & 63))
                                        if len(v1) > 0 {
                                            v1[0] = (v6 ^ (v6 >> ((2) & 63)))
                                        }
                                        var v68 float32 = 0.0
                                        if false {
                                            if false {
                                                fmt.Println(v67)
                                            } else {
                                                var v69 bool = true
                                                v70 := GenS2{F0: 127, F1: -119}
                                                if v69 {
                                                    var v71 float32 = (v68 * v68)
                                                    _ = v71
                                                } else {
                                                    var v72 float32 = (1e10 * ((1.5 / v68) * v68))
                                                    if (false || (0 == 100)) {
                                                        v72 += (v68 + v72)
                                                        switch "foo bar" {
                                                        case "":
                                                            var v73 int8 = 127
                                                            _ = v73
                                                        case "hello":
                                                            var v74 uint32 = (^(v6))
                                                            _ = v74
                                                        case "\n\t":
                                                            switch "a" {
                                                            case "\n\t":
                                                                v68 = v72
                                                            case "hello":
                                                                switch fi65 {
                                                                case 0:
                                                                    switch (("foo bar" + (("a" + "foo bar") + "a")) + "\n\t") {
                                                                    case "hello":
                                                                        _ = len(v1)
                                                                    case "a":
                                                                        _ = len(v1)
                                                                    }
                                                                default:
                                                                    fmt.Println(v72)
                                                                }
                                                            case "":
                                                                var v75 int32 = (v4 % ((-980) | 1))
                                                                _ = v75
                                                            default:
                                                                v76 := GenS4{F0: 0, F1: 255, F2: 34}
                                                                _ = v76
                                                            }
                                                        default:
                                                            v72 = float32(float64(((1.5 / (((v72 + (((((1.5 - 1e10) * (1.5 * 1e10)) * v72) / 1.0) - ((v68 + (((v72 / v72) * v68) * (v72 * v72))) / (float32(float64(-2.25)) + v68)))) + 0.0) * -2.25)) / (v72 - float32(0.0)))))
                                                        }
                                                    }
                                                    if ((v69 || (653 < 4294967295)) || v69) {
                                                        fmt.Println(v68, v67)
                                                        v77 := GenS5{F0: int16(v4), F1: -761}
                                                        _ = v77
                                                    }
                                                    _ = v72
                                                }
                                                _ = v69
                                                _ = v70
                                            }
                                        } else {
                                            v68 += ((3.14159 - 1e10) - v68)
                                        }
                                        _ = v67
                                        _ = v68
                                    }
                                }
                                var v78 uint64 = (v5 * 4294967295)
                                _ = v78
                            } else {
                                var v79 bool = true
                                switch (fi65 * (v3 - (v3 | int(v4)))) {
                                case 128:
                                    v80 := GenS5{F0: -1, F1: -1}
                                    _ = v80
                                case 127:
                                    if (!(1e10 >= 1e10)) {
                                        if (42 > 0) {
                                            v6 = (v6 ^ v6)
                                            v3 += -2147483648
                                            v3 = fi65
                                        } else {
                                            var v81 float32 = ((((1.5 - 1.5) - (0.0 * 3.14159)) / (-2.25 + (1e10 + (0.1 - (0.1 + (1e10 * 1e10)))))) + 1e10)
                                            fmt.Println(fi63, v4, fi63)
                                            var v82 int8 = -51
                                            _ = v81
                                            _ = v82
                                        }
                                        var v83 uint32 = v6
                                        delete(v2, v6)
                                        _ = v83
                                    } else {
                                        if ((v6 % ((v6) | 1)) == (^(v6))) {
                                            switch 0 {
                                            case 547:
                                                v5 = (v5 - v5)
                                            default:
                                                v3 += (-(v3))
                                            }
                                            _ = len(v1)
                                        } else {
                                            var v84 uint64 = 902
                                            var v85 uint8 = 169
                                            _ = v84
                                            _ = v85
                                        }
                                    }
                                case -357:
                                    var v86 int16 = 1000
                                    _ = v86
                                default:
                                    v6 += v6
                                }
                                fmt.Println(v4, fi63, v79)
                                _ = v79
                            }
                            if true {
                                switch (-(fi63)) {
                                case -2147483648:
                                    v3 = 2147483647
                                case 32767:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    v87 := GenS3{F0: 18446744073709551615, F1: 42, F2: -1}
                                    _ = v87
                                }
                                switch fi63 {
                                case 2147483647:
                                    v88 := v2[65535]
                                    _ = v88
                                default:
                                    v89 := v2[v6]
                                    _ = v89
                                }
                                var v90 uint = 4294967295
                                _ = v90
                            } else {
                                v5 -= (v5 / ((v5) | 1))
                                var v91 int64 = -1
                                var v92 uint32 = 4294967295
                                _ = v91
                                _ = v92
                            }
                            v3 = fi63
                        }
                    }
                    var v93 uint32 = 522
                    if (int32(v93) != (v4 % ((42) | 1))) {
                        if (32767 > 255) {
                            v6 -= (v6 | 4294967295)
                        }
                        if (128 == (fi65 << ((2) & 63))) {
                            var v94 int = (-(fi65))
                            var v95 int64 = 1
                            switch ("" + "\n\t") {
                            case "hello":
                                v96 := GenS2{F0: 145, F1: -1}
                                _ = v96
                            case "a":
                                if false {
                                    var v97 uint8 = 142
                                    _ = v97
                                }
                            }
                            var v98 bool = true
                            var v99 uint = 18446744073709551615
                            _ = v94
                            _ = v95
                            _ = v98
                            _ = v99
                        }
                        v4 += (v4 & v4)
                    } else {
                        fmt.Println(v4, v3)
                        if (false || false) {
                            v2[(v6 ^ v93)] = (fi65 - v3)
                            v5 -= 0
                            if true {
                                fmt.Println(fi63)
                            } else {
                                fmt.Println(fi63, v6, v3)
                            }
                            v3 += -1
                        }
                        v93 = (^(v6))
                    }
                    fmt.Println(v4, v5)
                    _ = fi65
                    _ = v93
                }
                _ = fi63
            }
            for fi100 := 0; fi100 < 10; fi100++ {
                var v101 float32 = 0.0
                var v102 uint64 = (v5 & 2)
                _ = fi100
                _ = v101
                _ = v102
            }
        }
    } else {
        v3 = v3
    }
    fmt.Println(v5, v5, v3)
    var v103 uint32 = 2147483647
    if len(v1) > 0 {
        v1[0] = (v103 >> ((1) & 63))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
