package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint8; F2 string }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint8; F1 int32; F2 uint }
type GenS9 struct { F0 int; F1 uint; F2 uint16 }
type GenS10 struct { F0 int32; F1 uint; F2 int }
type GenS11 struct { F0 bool; F1 uint; F2 uint64 }
type GenS12 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    var v3 uint8 = 1
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    var v6 float64 = 1e10
    if false {
        var v7 uint8 = v3
        _ = v7
    }
    var v8 string = "foo bar"
    var v9 float64 = v6
    v5 = append(v5, 873)
    delete(v1, 536)
    v6 = float64(-2.25)
    if true {
        var v10 bool = (1 != 450)
        v10 = false
        _ = v10
    } else {
        fmt.Println(v2)
        v6 -= (v6 * v6)
    }
    fmt.Println(v8, v9, v2)
    var v11 uint32 = 85
    switch 32768 {
    case 111:
        var v12 uint16 = v2
        _ = v12
    default:
        switch 9223372036854775807 {
        case -9223372036854775808:
            v11 -= (v11 | v11)
        case 32767:
            var v13 uint64 = 0
            _ = v13
        case -142:
            var v14 int = 128
            _ = v14
        default:
            if false {
                for fi15 := 0; fi15 < 6; fi15++ {
                    if false {
                        v16 := GenS1{F0: 0}
                        var v17 uint8 = v3
                        _ = v16
                        _ = v17
                    }
                    _ = len(v5)
                    _ = fi15
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    fmt.Println(v9, v9)
                    if (true || ((3.14159 / 1.5) >= (float32(v9) + 3.14159))) {
                        v19 := GenS1{F0: 349}
                        v20 := GenS2{F0: 1, F1: uint32(v3), F2: ((v8 > ("" + (v8 + v8))) && (v3 <= 121))}
                        _ = v19
                        _ = v20
                    } else {
                        if (!(94 < 42)) {
                            for fi21 := 0; fi21 < 9; fi21++ {
                                fmt.Println(v9, v3)
                                v22 := GenS2{F0: 42, F1: v11, F2: (!(!(-703 > 2147483647)))}
                                v11 = v11
                                _ = fi21
                                _ = v22
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            v11 -= v11
                            fmt.Println(v3, fi18, fi18)
                        }
                        if len(v5) > 0 {
                            v5[0] = fi18
                        }
                        if (!(v2 > v2)) {
                            fmt.Println(v2)
                            v5 = append(v5, fi18)
                            if (!(!false)) {
                                for fi23 := 0; fi23 < 2; fi23++ {
                                    fmt.Println(v2, v2, v3)
                                    _ = fi23
                                }
                            }
                        } else {
                            if (0 <= -32768) {
                                switch fi18 {
                                case 36:
                                    var v24 uint32 = v11
                                    _ = v24
                                }
                            }
                            var v25 int64 = -1
                            var v26 string = ((v8 + "foo bar") + v8)
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi18
                }
                switch 32768 {
                case 924:
                    var v27 int32 = 0
                    _ = v27
                case 2147483647:
                    for fi28 := 0; fi28 < 5; fi28++ {
                        if (-2147483648 < -251) {
                            fmt.Println(v8, v11)
                        }
                        _ = fi28
                    }
                case 4294967295:
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v30 := GenS1{F0: 816}
                        var v31 uint8 = v3
                        for fi32 := 0; fi32 < 12; fi32++ {
                            _ = len(v5)
                            if ((v6 / 0.1) > ((((v6 * (v6 + 3.14159)) - 0.0) - float64(1e10)) * (0.0 + v9))) {
                                if false {
                                    v6 = (((1e10 / (1.5 / v6)) + 1e10) / v6)
                                    v9 += v6
                                }
                                if true {
                                    var v33 uint8 = v3
                                    v33 = (v33 >> ((65535) & 63))
                                    v5 = append(v5, fi32)
                                    _ = v33
                                } else {
                                    var v34 uint8 = v31
                                    _ = v34
                                }
                                v35 := GenS3{F0: v11, F1: 499}
                                if false {
                                    switch ((v8 + "") + v8) {
                                    case "hello":
                                        v1[725] = fi29
                                    case "\n\t":
                                        v31 = (v3 + v31)
                                    case "":
                                        v11 += (v11 << ((65535) & 63))
                                        v36 := GenS4{F0: -2147483648}
                                        var v37 uint16 = (^(v2))
                                        v5 = append(v5, (fi29 & -176))
                                        _ = v36
                                        _ = v37
                                    default:
                                        v1[2147483647] = fi32
                                    }
                                    v11 += 4294967295
                                } else {
                                    if true {
                                        if false {
                                            v3 = (v31 >> ((18446744073709551615) & 63))
                                        } else {
                                            if (!true) {
                                                if true {
                                                    var v38 int16 = 1000
                                                    if (72 == 2147483647) {
                                                        fmt.Println(fi29)
                                                        var v39 uint8 = (^(v3))
                                                        var v40 int32 = -1
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    v41 := GenS2{F0: 127, F1: 504, F2: true}
                                                    switch (v8 + "") {
                                                    case "\n\t":
                                                        switch fi32 {
                                                        case 0:
                                                            v42 := GenS5{F0: (false || true)}
                                                            _ = v42
                                                        }
                                                    case "":
                                                        v43 := v4[(v11 ^ v11)]
                                                        _ = v43
                                                    default:
                                                        v3 += v31
                                                    }
                                                    if false {
                                                        fmt.Println(fi32, v31, v11)
                                                    }
                                                    _ = v38
                                                    _ = v41
                                                }
                                                v44 := GenS4{F0: -522}
                                                _ = v44
                                            } else {
                                                var v45 string = v8
                                                if (true && true) {
                                                    switch fi29 {
                                                    case 42:
                                                        v8 = ""
                                                    default:
                                                        var v46 int16 = 463
                                                        _ = v46
                                                    }
                                                }
                                                _ = v45
                                            }
                                            v47 := GenS2{F0: -112, F1: (v11 + (^(v11))), F2: (18446744073709551615 > 490)}
                                            switch 585 {
                                            case 32767:
                                                if (!true) {
                                                    switch ("\n\t" + "a") {
                                                    case "a":
                                                        v6 += v6
                                                    case "foo bar":
                                                        var v48 uint32 = v11
                                                        _ = v48
                                                    case "\n\t":
                                                        if (-4 == -105) {
                                                            var v49 string = v8
                                                            v50 := GenS5{F0: true}
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            v51 := GenS1{F0: 2}
                                                            v6 = (0.0 + 0.0)
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                    var v52 int16 = 42
                                                    var v53 string = v8
                                                    _ = v52
                                                    _ = v53
                                                }
                                            case 128:
                                                switch "a" {
                                                case "\n\t":
                                                    v54 := GenS6{F0: v8, F1: (v31 | uint8(v2)), F2: ((v8 + "hello") + v8)}
                                                    _ = v54
                                                case "foo bar":
                                                    switch "" {
                                                    case "foo bar":
                                                        var v55 uint16 = (v2 * 65535)
                                                        _ = v55
                                                    default:
                                                        fmt.Println(fi32, v9, v11)
                                                    }
                                                default:
                                                    var v56 uint64 = 858
                                                    _ = v56
                                                }
                                            case -2147483648:
                                                var v57 string = (v8 + "hello")
                                                _ = v57
                                            }
                                            _ = v47
                                        }
                                        v9 += float64(1e10)
                                        v58 := GenS7{F0: 938, F1: v8}
                                        switch ((v8 + v8) + "foo bar") {
                                        case "foo bar":
                                            v59 := GenS1{F0: 1}
                                            switch fi29 {
                                            case 4294967295:
                                                if false {
                                                    v3 = uint8(v2)
                                                }
                                            }
                                            if false {
                                                fmt.Println(fi32, v3)
                                            } else {
                                                switch v8 {
                                                case "\n\t":
                                                    v60 := GenS8{F0: v31, F1: -839, F2: 2}
                                                    _ = v60
                                                case "foo bar":
                                                    _ = len(v5)
                                                case "a":
                                                    switch (fi29 + fi32) {
                                                    case -2147483648:
                                                        if false {
                                                            if false {
                                                                fmt.Println(fi32, v3)
                                                                if false {
                                                                    if true {
                                                                        v6 -= 1e10
                                                                    } else {
                                                                        v6 += v9
                                                                    }
                                                                    v61 := v4[v11]
                                                                    _ = v61
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v9, v6)
                                                                        var v62 float32 = 0.0
                                                                        switch int(v3) {
                                                                        case 272:
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                        _ = v62
                                                                    } else {
                                                                        if true {
                                                                            switch (v8 + (((v8 + ((("foo bar" + "foo bar") + v8) + v8)) + "foo bar") + v8)) {
                                                                            case "foo bar":
                                                                                var v63 uint8 = 0
                                                                                _ = v63
                                                                            default:
                                                                                if (!(!(fi32 > fi32))) {
                                                                                    v64 := GenS9{F0: fi32, F1: 948, F2: 2}
                                                                                    fmt.Println(v31, v2, v6)
                                                                                    v65 := GenS6{F0: v8, F1: 0, F2: ("" + v8)}
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                }
                                                                            }
                                                                            fmt.Println(v3, v11, fi32)
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            var v66 uint64 = 936
                                                                            var v67 float32 = ((((0.0 * 3.14159) * ((float32(((((1e10 + v9) + float64((0.1 / 0.1))) - (((1e10 / ((0.0 - v6) * 1.5)) * (v6 + 1e10)) - float64((((-2.25 + float32(float64((1e10 - 3.14159)))) * (0.1 - 1.5)) + 0.1)))) - 0.1)) / 1e10) / -2.25)) - 3.14159) + (1e10 * 0.1))
                                                                            if true {
                                                                                switch -2147483648 {
                                                                                case -9223372036854775808:
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = 32767
                                                                                    }
                                                                                case 127:
                                                                                    var v68 uint = 42
                                                                                    _ = v68
                                                                                case 9223372036854775807:
                                                                                    v2 += (^(v2))
                                                                                }
                                                                            }
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        v11 = v11
                                                                    }
                                                                    var v69 int8 = 0
                                                                    _ = v69
                                                                }
                                                                var v70 uint8 = 0
                                                                _ = v70
                                                            } else {
                                                                if len(v5) > 0 {
                                                                    v5[0] = fi32
                                                                }
                                                                v11 = (v11 | 2)
                                                                if (0.0 < (0.1 + ((((1e10 + float32(v9)) + 0.1) + -2.25) + ((0.1 / 1e10) - float32(((float64(((0.0 * (1e10 + float32(v9))) * 1.5)) + v6) / (v9 + float64(0.1)))))))) {
                                                                    v9 += 3.14159
                                                                    if false {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            if (0 > v2) {
                                                                                v8 = v8
                                                                                if (835 == 4294967295) {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = fi32
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    switch (-(fi29)) {
                                                                    case 128:
                                                                        v71 := GenS10{F0: 42, F1: 392, F2: 351}
                                                                        _ = v71
                                                                    case 748:
                                                                        v72 := GenS4{F0: 218}
                                                                        _ = v72
                                                                    case 1:
                                                                        var v73 uint64 = 2
                                                                        _ = v73
                                                                    }
                                                                    fmt.Println(v3)
                                                                }
                                                            }
                                                        }
                                                    case 408:
                                                        switch v8 {
                                                        case "foo bar":
                                                            v2 += v2
                                                        default:
                                                            if false {
                                                                v11 += v11
                                                                v11 -= (v11 - (v11 + v11))
                                                            } else {
                                                                _ = len(v5)
                                                                v74 := v1[v11]
                                                                _ = v74
                                                            }
                                                        }
                                                        var v75 int8 = -17
                                                        v11 = 65535
                                                        v75 = v75
                                                        _ = v75
                                                    case 221:
                                                        if false {
                                                            var v76 bool = false
                                                            _ = v76
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        if false {
                                                            v77 := GenS4{F0: 42}
                                                            switch v8 {
                                                            case "":
                                                                v9 = v9
                                                            case "foo bar":
                                                                var v78 int64 = -9223372036854775808
                                                                _ = v78
                                                            case "hello":
                                                                v79 := GenS11{F0: (true || true), F1: 65535, F2: 1}
                                                                _ = v79
                                                            }
                                                            var v80 int8 = 100
                                                            _ = v77
                                                            _ = v80
                                                        }
                                                        delete(v4, v11)
                                                    }
                                                }
                                                v3 += (v31 | (v3 ^ (v3 ^ (v31 - 1))))
                                                var v81 int64 = 0
                                                _ = v81
                                            }
                                            _ = v59
                                        case "a":
                                            if (!true) {
                                                fmt.Println(v11, v2)
                                                fmt.Println(v6, fi29, v9)
                                            } else {
                                                v82 := GenS12{F0: 825}
                                                if false {
                                                    if true {
                                                        var v83 int8 = 1
                                                        if true {
                                                            var v84 uint32 = 4294967295
                                                            v11 -= 65535
                                                            _ = v84
                                                        }
                                                        if false {
                                                            v85 := GenS4{F0: -453}
                                                            _ = v85
                                                        }
                                                        _ = v83
                                                    }
                                                    v9 = v9
                                                    var v86 string = (v8 + "foo bar")
                                                    v87 := GenS7{F0: 0, F1: ((v86 + ((((v86 + "") + ((v8 + (v8 + v86)) + "foo bar")) + v8) + v86)) + ((v86 + ((v8 + "hello") + v8)) + v8))}
                                                    if len(v5) > 0 {
                                                        v5[0] = 2147483647
                                                    }
                                                    switch v8 {
                                                    case "a":
                                                        v88 := GenS9{F0: fi32, F1: 42, F2: (v2 - (v2 | (v2 % ((v2) | 1))))}
                                                        _ = v88
                                                    default:
                                                        fmt.Println(v11)
                                                    }
                                                    _ = v86
                                                    _ = v87
                                                }
                                                var v89 uint = 0
                                                _ = v82
                                                _ = v89
                                            }
                                        }
                                        var v90 uint16 = (^(v2))
                                        _ = v58
                                        _ = v90
                                    }
                                    fmt.Println(v11, v9, v6)
                                }
                                _ = len(v5)
                                var v91 float32 = -2.25
                                _ = v35
                                _ = v91
                            }
                            _ = fi32
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
            } else {
                var v92 uint64 = 4294967295
                v6 -= 1.5
                _ = v92
            }
        }
    }
    for fi93 := 0; fi93 < 8; fi93++ {
        if (1 != 65535) {
            var v94 bool = true
            if true {
                v9 = (1.5 / (1.5 * ((v9 * v6) - ((0.1 - (((float64((float32(float64((0.0 + ((0.0 + 1e10) + 1.5)))) - (float32(3.14159) / -2.25))) / 3.14159) - v9) * 0.0)) - float64((((1.5 * 1.5) + -2.25) * 1.5))))))
                fmt.Println(v6)
            } else {
                for fi95 := 0; fi95 < 9; fi95++ {
                    var v96 uint = 911
                    if v94 {
                        var v97 uint16 = (v2 + (v2 + uint16(v11)))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v97
                    }
                    _ = fi95
                    _ = v96
                }
            }
            _ = v94
        }
        _ = fi93
    }
    switch 348 {
    case 587:
        fmt.Println(v2, v3)
    case 127:
        var v98 int16 = 0
        _ = v98
    }
    v2 -= v2
    var v99 uint32 = (v11 + 42)
    var v100 uint32 = (v99 << ((4294967295) & 63))
    v6 -= 1.5
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
    _ = v99
}
