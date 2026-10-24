package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 int8; F1 uint16; F2 int64 }
type GenS4 struct { F0 int16; F1 int16; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 uint; F2 int8 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int64; F1 bool; F2 int64 }
func genFn0() uint64 {
    var v7 uint64 = 326
    var v8 uint = 18446744073709551615
    fmt.Println(v8)
    if false {
        if false {
            var v9 int16 = -32768
            _ = v9
        } else {
            switch 997 {
            case 32768:
                v7 = v7
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := []int{}
                    v13 := GenS1{F0: (true || (false || false)), F1: ""}
                    var v14 int16 = -1
                    fmt.Println(v14, v14)
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
            var v15 uint = (v8 & v8)
            _ = v15
        }
        v16 := GenS1{F0: true, F1: ""}
        v7 += 42
        _ = v16
    }
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 330
    var v3 uint64 = 1
    var v4 float64 = 1.5
    var v5 int8 = 50
    v6 := []uint32{v2, v2, v2, 4294967295}
    for fi17 := 0; fi17 < 1; fi17++ {
        if false {
            if ((v2 * 2147483647) <= (v2 % ((v2) | 1))) {
                v18 := v1[fi17]
                _ = len(v1)
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 int8 = -128
                    var v21 int16 = 42
                    v2 += v18
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            } else {
                v22 := GenS1{F0: (42 >= 995), F1: "foo bar"}
                switch fi17 {
                case 128:
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v24 := GenS2{F0: fi23, F1: 127, F2: "a"}
                        _ = fi23
                        _ = v24
                    }
                case 9223372036854775807:
                    if (false && true) {
                        fmt.Println(v2, v3, v5)
                        v3 += v3
                    }
                case 2147483647:
                    fmt.Println(v4, fi17, fi17)
                default:
                    fmt.Println(v5, v5)
                }
                _ = v22
            }
            if true {
                v25 := GenS2{F0: fi17, F1: 1, F2: ""}
                _ = v25
            }
            if (true || (42 <= 432)) {
                v6 = append(v6, 42)
                if (v3 != (v3 ^ (v3 >> ((255) & 63)))) {
                    switch "\n\t" {
                    case "a":
                        var v26 float64 = 0.1
                        _ = v26
                    case "foo bar":
                        fmt.Println(v4, v4)
                    }
                    for fi27 := 0; fi27 < 8; fi27++ {
                        if len(v6) > 0 {
                            v6[0] = v2
                        }
                        _ = fi27
                    }
                    var v28 int8 = (v5 / (((v5 / (((v5 / ((int8(v3)) | 1))) | 1))) | 1))
                    _ = v28
                }
                v3 -= v3
            } else {
                v1[fi17] = (^(v2))
                var v29 int64 = 0
                switch (fi17 * (fi17 & fi17)) {
                case 127:
                    var v30 float32 = float32(v4)
                    _ = v30
                case 265:
                    v29 = 24
                }
                _ = v29
            }
        } else {
            fmt.Println(v5)
            v3 = 2
        }
        _ = len(v6)
        v31 := GenS3{F0: v5, F1: 287, F2: -9223372036854775808}
        _ = fi17
        _ = v31
    }
    var v32 uint16 = 0
    var v33 string = "foo bar"
    if ((v4 - 0.1) < (v4 / v4)) {
        fmt.Println(v33, v2)
    } else {
        for fi34 := 0; fi34 < 18; fi34++ {
            for fi35 := 0; fi35 < 15; fi35++ {
                v5 -= -1
                _ = fi35
            }
            fmt.Println(v5)
            _ = fi34
        }
        v36 := GenS4{F0: -684, F1: 1, F2: 32767}
        _ = v36
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        if ((!(false && true)) && false) {
            v38 := GenS5{F0: (-(fi37))}
            for fi39 := 0; fi39 < 13; fi39++ {
                v3 = (v3 + uint64(v2))
                if (-2.25 != 1.5) {
                    switch (fi39 & 9223372036854775807) {
                    case -636:
                        v33 = (v33 + v33)
                    case 32768:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case 42:
                        switch v33 {
                        case "foo bar":
                            v3 = 842
                        default:
                            switch v33 {
                            case "":
                                v40 := GenS6{F0: 128, F1: -50}
                                v32 -= 0
                                _ = v40
                            case "hello":
                                v3 -= 2
                            case "a":
                                var v41 string = v33
                                _ = v41
                            }
                        }
                    default:
                        fmt.Println(v32, fi39)
                        var v42 uint16 = uint16(v3)
                        _ = v42
                    }
                    switch "\n\t" {
                    case "":
                        v3 = (v3 ^ (v3 * genFn0()))
                    case "hello":
                        switch 0 {
                        case -9223372036854775808:
                            v2 -= v2
                        default:
                            var v43 uint8 = 42
                            _ = v43
                        }
                        var v44 string = v33
                        switch 589 {
                        case 42:
                            v2 = v2
                            var v45 uint = 65535
                            fmt.Println(v45, v2, v45)
                            if (!(42 < 641)) {
                                if (false || false) {
                                    var v46 uint16 = 32767
                                    v2 = v2
                                    if false {
                                        fmt.Println(v44, fi39, v4)
                                    }
                                    _ = v46
                                }
                            } else {
                                v4 = ((v4 * float64((3.14159 * 1e10))) * v4)
                                switch (v44 + v44) {
                                case "hello":
                                    v3 = (v3 << ((18446744073709551615) & 63))
                                case "\n\t":
                                    var v47 int64 = -807
                                    v3 = (v3 | v3)
                                    var v48 int = fi37
                                    _ = v47
                                    _ = v48
                                case "foo bar":
                                    switch (fi39 << ((v45) & 63)) {
                                    case -592:
                                        if (v4 != float64(0.0)) {
                                            var v49 uint = v45
                                            if true {
                                                var v50 uint16 = v32
                                                var v51 uint32 = (v2 * v2)
                                                _ = v50
                                                _ = v51
                                            }
                                            v3 = genFn0()
                                            var v52 int16 = int16(v2)
                                            _ = v49
                                            _ = v52
                                        } else {
                                            var v53 float64 = float64(1e10)
                                            v3 = (v3 / (((v3 / ((4294967295) | 1))) | 1))
                                            var v54 uint8 = 127
                                            _ = len(v1)
                                            v2 = v2
                                            _ = v53
                                            _ = v54
                                        }
                                    case 128:
                                        fmt.Println(v44)
                                    }
                                }
                            }
                            _ = v45
                        case 478:
                            fmt.Println(v32, fi37)
                        case -507:
                            if true {
                                v55 := GenS4{F0: 0, F1: -1, F2: fi39}
                                var v56 uint32 = (v2 + (v2 - v2))
                                _ = v55
                                _ = v56
                            } else {
                                var v57 int16 = 809
                                _ = v57
                            }
                        default:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        _ = v44
                    case "\n\t":
                        var v58 int32 = -2147483648
                        _ = v58
                    }
                } else {
                    if len(v6) > 0 {
                        v6[0] = v2
                    }
                    v33 = v33
                }
                var v59 uint8 = 127
                fmt.Println(v32)
                _ = fi39
                _ = v59
            }
            v60 := GenS6{F0: 128, F1: v5}
            _ = v38
            _ = v60
        } else {
            switch v33 {
            case "foo bar":
                for fi61 := 0; fi61 < 4; fi61++ {
                    switch fi61 {
                    case 437:
                        var v62 int32 = -2147483648
                        _ = v62
                    case 4294967295:
                        if true {
                            switch fi37 {
                            case 0:
                                var v63 int64 = -9223372036854775808
                                _ = v63
                            case -588:
                                switch "a" {
                                case "hello":
                                    if (!(!(!(true || false)))) {
                                        if len(v6) > 0 {
                                            v6[0] = v2
                                        }
                                        var v64 bool = false
                                        _ = v64
                                    }
                                }
                                if true {
                                    if false {
                                        v4 = ((v4 * 1.5) + 0.1)
                                    } else {
                                        v1[128] = (v2 ^ 47)
                                        v4 -= float64((1e10 + 1e10))
                                    }
                                    var v65 int16 = 7
                                    _ = v65
                                } else {
                                    switch (-(fi37)) {
                                    case -2147483648:
                                        var v66 uint8 = 152
                                        _ = v66
                                    case 127:
                                        v67 := GenS7{F0: 127}
                                        _ = v67
                                    case -797:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    default:
                                        v5 = v5
                                    }
                                    switch ("a" + v33) {
                                    case "foo bar":
                                        v68 := GenS7{F0: 241}
                                        _ = v68
                                    case "hello":
                                        fmt.Println(v32, v3, v3)
                                    case "\n\t":
                                        var v69 uint64 = v3
                                        _ = v69
                                    }
                                    fmt.Println(v2, fi37)
                                    v70 := GenS6{F0: 98, F1: v5}
                                    var v71 bool = true
                                    fmt.Println(v33, v4)
                                    _ = v70
                                    _ = v71
                                }
                                var v72 uint8 = 2
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v73 int16 = 42
                                _ = v72
                                _ = v73
                            default:
                                v4 += 1.5
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                switch (fi61 >> ((18446744073709551615) & 63)) {
                                case 864:
                                    var v74 int16 = 42
                                    _ = v74
                                case -2147483648:
                                    if false {
                                        if (!(fi37 == fi37)) {
                                            v33 = ""
                                            v5 += v5
                                        }
                                        v75 := v1[int(v3)]
                                        var v76 int = 626
                                        if ((!false) && true) {
                                            switch "\n\t" {
                                            case "\n\t":
                                                v32 += v32
                                            default:
                                                switch (fi37 >> ((42) & 63)) {
                                                case -350:
                                                    v77 := GenS1{F0: (int32(v5) <= 2147483647), F1: ""}
                                                    _ = v77
                                                case 4294967295:
                                                    v78 := GenS8{F0: 0, F1: 2, F2: v5}
                                                    _ = v78
                                                default:
                                                    v32 = (v32 | (v32 & (v32 % ((32767) | 1))))
                                                }
                                            }
                                            if false {
                                                fmt.Println(v3, fi61, v75)
                                                v33 = ""
                                                v79 := GenS9{F0: 179, F1: 9223372036854775807}
                                                _ = v79
                                            }
                                            v80 := GenS9{F0: 0, F1: 0}
                                            if len(v6) > 0 {
                                                v6[0] = 2147483648
                                            }
                                            _ = v80
                                        } else {
                                            v6 = append(v6, 221)
                                            v81 := GenS5{F0: fi61}
                                            if len(v6) > 0 {
                                                v6[0] = v2
                                            }
                                            _ = v81
                                        }
                                        _ = v75
                                        _ = v76
                                    }
                                default:
                                    if len(v6) > 0 {
                                        v6[0] = v2
                                    }
                                }
                                if ((^(v32)) == uint16(v3)) {
                                    switch v33 {
                                    case "":
                                        if ((v3 >> ((1) & 63)) == genFn0()) {
                                            if len(v6) > 0 {
                                                v6[0] = v2
                                            }
                                            _ = len(v1)
                                            v4 = (v4 * 3.14159)
                                        }
                                    case "hello":
                                        fmt.Println(v4)
                                        delete(v1, (^(fi61)))
                                    }
                                    var v82 int = (fi61 & 0)
                                    _ = v82
                                }
                                v83 := GenS3{F0: int8(v32), F1: v32, F2: 0}
                                v84 := GenS10{F0: -1, F1: false, F2: int64(v32)}
                                _ = v83
                                _ = v84
                            }
                            v6 = append(v6, 2)
                        } else {
                            var v85 int = 640
                            var v86 int8 = v5
                            var v87 float32 = 0.1
                            _ = v85
                            _ = v86
                            _ = v87
                        }
                    default:
                        var v88 uint64 = 42
                        _ = v88
                    }
                    fmt.Println(v33, v33)
                    _ = fi61
                }
            default:
                for fi89 := 0; fi89 < 4; fi89++ {
                    if false {
                        delete(v1, fi37)
                        switch v33 {
                        case "":
                            fmt.Println(v2, fi89, v32)
                        case "a":
                            v33 = ((((v33 + ("" + v33)) + ("a" + v33)) + v33) + "foo bar")
                        }
                    }
                    v33 = (v33 + v33)
                    var v90 uint32 = v2
                    _ = fi89
                    _ = v90
                }
            }
            _ = len(v1)
        }
        v91 := v1[(^(fi37))]
        _ = fi37
        _ = v91
    }
    var v92 uint64 = v3
    var v93 uint64 = 575
    v1[127] = v2
    var v94 string = v33
    var v95 uint = 844
    for fi96 := 0; fi96 < 10; fi96++ {
        for fi97 := 0; fi97 < 14; fi97++ {
            var v98 uint16 = v32
            if true {
                v92 = v92
            } else {
                var v99 uint32 = v2
                v6 = append(v6, v2)
                var v100 float32 = 0.0
                _ = v100
                _ = v99
            }
            _ = fi97
            _ = v98
        }
        v101 := GenS8{F0: (v95 * (v95 - (v95 << ((uint(v2)) & 63)))), F1: 65535, F2: -61}
        v102 := GenS8{F0: (v95 - v95), F1: 255, F2: (v5 & 0)}
        _ = fi96
        _ = v101
        _ = v102
    }
    v4 += (v4 * (v4 / v4))
    for fi103 := 0; fi103 < 14; fi103++ {
        var v104 uint8 = 2
        switch (-(fi103)) {
        case 1:
            for fi105 := 0; fi105 < 4; fi105++ {
                v106 := v1[fi105]
                var v107 uint = 1
                v95 = uint(v3)
                _ = fi105
                _ = v106
                _ = v107
            }
        }
        v1[128] = v2
        _ = fi103
        _ = v104
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v92
    _ = v93
    _ = v94
    _ = v95
}
