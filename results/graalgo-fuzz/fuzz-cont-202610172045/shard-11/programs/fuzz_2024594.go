package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 uint8; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 74
    for fi4 := 0; fi4 < 15; fi4++ {
        v3 = (v3 - v3)
        fmt.Println(fi4, v3)
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: 32767, F1: -9223372036854775808, F2: false}
            v3 += (v3 | 103)
            fmt.Println(v3)
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v3 = v3
    var v7 int8 = 127
    var v8 uint32 = 65535
    if true {
        var v9 int64 = 42
        fmt.Println(v9, v7)
        switch "a" {
        case "hello":
            switch "hello" {
            case "hello":
                var v10 uint16 = 915
                _ = v10
            default:
                v8 -= v8
            }
        case "foo bar":
            for fi11 := 0; fi11 < 6; fi11++ {
                v9 = (v9 << ((255) & 63))
                var v12 bool = ((v7 - v7) != (v7 & v7))
                _ = fi11
                _ = v12
            }
        case "":
            v13 := []int{int(v3)}
            _ = v13
        default:
            var v14 uint64 = 420
            _ = v14
        }
        _ = v9
    } else {
        if true {
            if true {
                var v15 float32 = (0.1 - 0.0)
                v3 = (v3 & (v3 + v3))
                _ = v15
            } else {
                v3 += v3
                switch "foo bar" {
                case "foo bar":
                    for fi16 := 0; fi16 < 16; fi16++ {
                        if (1 > 2) {
                            v3 -= v3
                        }
                        for fi17 := 0; fi17 < 20; fi17++ {
                            var v18 uint8 = (v3 ^ v3)
                            if (2147483647 > 1) {
                                fmt.Println(v8, fi16)
                                fmt.Println(v7, fi17, v3)
                            } else {
                                v19 := GenS2{F0: 1, F1: -9223372036854775808, F2: -158}
                                delete(v2, (v8 | v8))
                                _ = v19
                            }
                            if false {
                                if ((!(1.5 != -2.25)) || (9223372036854775807 != -1)) {
                                    fmt.Println(v7)
                                    v3 += v3
                                }
                                switch ("\n\t" + "hello") {
                                case "":
                                    if (909 > 42) {
                                        if true {
                                            if (true || false) {
                                                if false {
                                                    delete(v2, v8)
                                                } else {
                                                    var v20 int8 = (v7 * v7)
                                                    var v21 int8 = v7
                                                    _ = v20
                                                    _ = v21
                                                }
                                                if (!(!true)) {
                                                    if (!((1 <= 18446744073709551615) && (v7 >= (v7 >> ((42) & 63))))) {
                                                        v22 := []int{(fi16 / (((fi16 ^ (fi16 / ((int(v8)) | 1)))) | 1)), (fi16 - (fi17 ^ (fi17 - fi16))), fi16}
                                                        var v23 uint32 = v8
                                                        v22 = append(v22, (fi17 - fi16))
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                }
                                            } else {
                                                if ((true || false) || false) {
                                                    var v24 float64 = 1e10
                                                    _ = v24
                                                } else {
                                                    v25 := GenS2{F0: v18, F1: 964, F2: -1}
                                                    v7 += v7
                                                    switch 936 {
                                                    case 4294967295:
                                                        switch fi16 {
                                                        case -9223372036854775808:
                                                            if false {
                                                                if (true && true) {
                                                                    var v26 uint = 2
                                                                    if (!(472 >= -1)) {
                                                                        fmt.Println(fi16)
                                                                        if (v18 >= v3) {
                                                                            var v27 int32 = 42
                                                                            var v28 int32 = v27
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                var v29 uint32 = (v8 & (v8 & 512))
                                                                                _ = v29
                                                                            }
                                                                            _ = v27
                                                                            _ = v28
                                                                        }
                                                                        var v30 int16 = 42
                                                                        _ = v30
                                                                    } else {
                                                                        v26 = uint(v8)
                                                                    }
                                                                    _ = v26
                                                                }
                                                            }
                                                        case -130:
                                                            v31 := GenS2{F0: v3, F1: 42, F2: -32768}
                                                            v32 := v1[v8]
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                    case -848:
                                                        var v33 uint32 = v8
                                                        _ = v33
                                                    case -973:
                                                        var v34 int8 = v7
                                                        _ = v34
                                                    }
                                                    _ = v25
                                                }
                                                v18 += (^(v18))
                                            }
                                        }
                                        switch fi16 {
                                        case 486:
                                            var v35 uint32 = (v8 >> ((1) & 63))
                                            _ = v35
                                        }
                                    } else {
                                        v18 -= (v18 | (v18 * 255))
                                        var v36 uint8 = v18
                                        if (716 != 255) {
                                            var v37 bool = false
                                            v3 -= 128
                                            if (55 >= (v3 | (v18 % (((v36 / ((v36) | 1))) | 1)))) {
                                                var v38 bool = false
                                                _ = v38
                                            }
                                            _ = v37
                                        } else {
                                            var v39 uint64 = 42
                                            switch (fi17 - fi17) {
                                            case 580:
                                                if true {
                                                    v7 += (-(v7))
                                                    v39 = 2
                                                }
                                            case -412:
                                                v8 = v8
                                                fmt.Println(v18, v7, v3)
                                                if (!(!true)) {
                                                    v40 := GenS1{F0: 42, F1: -1, F2: false}
                                                    var v41 int64 = -778
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    v42 := GenS2{F0: v3, F1: int64(v7), F2: 48}
                                                    v43 := []int{723}
                                                    _ = v42
                                                    _ = v43
                                                }
                                            default:
                                                v44 := GenS2{F0: 60, F1: 798, F2: -32768}
                                                _ = v44
                                            }
                                            var v45 uint64 = v39
                                            _ = v39
                                            _ = v45
                                        }
                                        v46 := GenS2{F0: v18, F1: 932, F2: int16(v7)}
                                        if (!false) {
                                            switch ("\n\t" + ((("" + "") + "foo bar") + "foo bar")) {
                                            case "":
                                                v47 := []int{}
                                                _ = v47
                                            default:
                                                fmt.Println(v8, v8, v18)
                                            }
                                            v18 = v3
                                            v18 += v36
                                        } else {
                                            var v48 float32 = 1.5
                                            var v49 uint16 = 1
                                            _ = v48
                                            _ = v49
                                        }
                                        if (1 <= 0) {
                                            fmt.Println(v7)
                                        } else {
                                            var v50 uint64 = 42
                                            var v51 uint64 = 86
                                            if true {
                                                v52 := GenS2{F0: 128, F1: 535, F2: 42}
                                                _ = v52
                                            }
                                            _ = v50
                                            _ = v51
                                        }
                                        _ = v36
                                        _ = v46
                                    }
                                case "foo bar":
                                    if (!true) {
                                        v3 -= (v18 % ((v3) | 1))
                                        delete(v1, v8)
                                        v53 := []int{(fi17 % (((^(fi17))) | 1)), (fi16 ^ (fi16 & (fi16 | (^(fi17)))))}
                                        var v54 float32 = (0.0 - (-2.25 / float32(float64(0.1))))
                                        _ = v53
                                        _ = v54
                                    }
                                default:
                                    v55 := GenS1{F0: 42, F1: -248, F2: false}
                                    _ = v55
                                }
                            } else {
                                if true {
                                    _ = len(v2)
                                    _ = len(v1)
                                }
                                v7 += (^(v7))
                                v56 := []int{fi16, -2147483648}
                                _ = v56
                            }
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                case "":
                    for fi57 := 0; fi57 < 15; fi57++ {
                        v3 += v3
                        _ = fi57
                    }
                default:
                    fmt.Println(v7, v3)
                }
                for fi58 := 0; fi58 < 20; fi58++ {
                    switch fi58 {
                    case 32767:
                        if (((v3 << ((uint(v3)) & 63)) != v3) || true) {
                            v8 += (v8 + v8)
                            v7 += 100
                        }
                    case 1:
                        for fi59 := 0; fi59 < 4; fi59++ {
                            fmt.Println(fi58, v7)
                            var v60 int8 = -107
                            _ = fi59
                            _ = v60
                        }
                    case 4294967295:
                        if (65535 < 0) {
                            fmt.Println(v3, v8, fi58)
                            if ((true && true) && false) {
                                v61 := v2[v8]
                                v8 = (v8 << ((1) & 63))
                                v8 = v61
                                _ = v61
                            } else {
                                var v62 uint64 = 854
                                for fi63 := 0; fi63 < 15; fi63++ {
                                    if true {
                                        var v64 bool = true
                                        _ = v64
                                    }
                                    var v65 int8 = v7
                                    switch (fi58 << ((918) & 63)) {
                                    case -324:
                                        if (!(!false)) {
                                            if (-999 == int16(v62)) {
                                                fmt.Println(v8, fi58, v65)
                                            } else {
                                                v66 := []int{128, 32767, 128, 32767}
                                                var v67 float64 = 0.0
                                                if true {
                                                    switch ("foo bar" + "") {
                                                    case "":
                                                        switch (("foo bar" + "\n\t") + ("foo bar" + "a")) {
                                                        case "hello":
                                                            var v68 float32 = float32(v67)
                                                            _ = v68
                                                        case "a":
                                                            switch ("foo bar" + "hello") {
                                                            case "\n\t":
                                                                v66 = append(v66, fi58)
                                                            case "a":
                                                                var v69 uint64 = 1
                                                                _ = v69
                                                            }
                                                        default:
                                                            v3 += v3
                                                        }
                                                    case "hello":
                                                        v70 := GenS2{F0: v3, F1: 0, F2: -32768}
                                                        _ = v70
                                                    default:
                                                        delete(v2, (v8 - v8))
                                                    }
                                                    var v71 bool = (65535 < 32767)
                                                    _ = v71
                                                } else {
                                                    fmt.Println(v65, fi58)
                                                    v72 := GenS2{F0: (^(v3)), F1: 9223372036854775807, F2: -865}
                                                    var v73 uint = 18446744073709551615
                                                    var v74 uint64 = (^(v62))
                                                    fmt.Println(fi58, v65, v7)
                                                    v75 := GenS3{F0: 127, F1: true}
                                                    v76 := GenS4{F0: ((v8 + v8) <= v8), F1: (v74 ^ 1)}
                                                    switch (^(fi58)) {
                                                    case 349:
                                                        v62 -= (v74 - v62)
                                                    case 751:
                                                        v67 -= v67
                                                    case 9223372036854775807:
                                                        v77 := GenS4{F0: true, F1: 43}
                                                        _ = v77
                                                    default:
                                                        if true {
                                                            if true {
                                                                var v78 bool = false
                                                                _ = v78
                                                            }
                                                        } else {
                                                            var v79 int32 = 42
                                                            var v80 int64 = 9223372036854775807
                                                            var v81 float64 = v67
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                    }
                                                    var v82 int8 = (v7 / ((v7) | 1))
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                    _ = v76
                                                    _ = v82
                                                }
                                                if true {
                                                    v65 += v7
                                                    fmt.Println(v62, v67)
                                                    if false {
                                                        if false {
                                                            delete(v2, 305)
                                                            v1[65535] = 527
                                                            var v83 float64 = (v67 / (v67 + 0.0))
                                                            _ = v83
                                                        } else {
                                                            var v84 float64 = (float64(3.14159) * (1.5 * v67))
                                                            v85 := GenS5{F0: 1, F1: -679, F2: (v62 - v62)}
                                                            var v86 int64 = -1
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                    }
                                                } else {
                                                    var v87 int16 = -379
                                                    _ = v87
                                                }
                                                _ = v66
                                                _ = v67
                                            }
                                            v3 += (v3 + (v3 | (^(v3))))
                                            fmt.Println(v7, v8, v8)
                                        } else {
                                            v62 += 571
                                            v88 := v2[65]
                                            v89 := GenS6{F0: 0}
                                            _ = v88
                                            _ = v89
                                        }
                                    }
                                    _ = fi63
                                    _ = v65
                                }
                                _ = v62
                            }
                        }
                    }
                    _ = fi58
                }
            }
        } else {
            var v90 int8 = (v7 ^ (^(v7)))
            _ = v90
        }
    }
    v8 -= (v8 | v8)
    v91 := GenS2{F0: 128, F1: 0, F2: 1}
    fmt.Println(v8, v7, v3)
    for fi92 := 0; fi92 < 6; fi92++ {
        if false {
            for fi93 := 0; fi93 < 8; fi93++ {
                v8 -= 42
                if (100 != 100) {
                    var v94 int16 = 32767
                    _ = v94
                }
                _ = fi93
            }
            var v95 int16 = 0
            v7 = (v7 - v7)
            _ = v95
        }
        _ = len(v1)
        _ = fi92
    }
    for fi96 := 0; fi96 < 6; fi96++ {
        v97 := GenS7{F0: "foo bar", F1: v3, F2: 4294967295}
        var v98 uint = 18446744073709551615
        _ = fi96
        _ = v97
        _ = v98
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v91
}
