package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 string; F1 int8; F2 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16; F1 bool }
type GenS13 struct { F0 uint; F1 string; F2 bool }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int32; F1 int8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint8; F1 int16; F2 int16 }
type GenS18 struct { F0 uint16; F1 uint; F2 int8 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint64; F1 int; F2 int32 }
type GenS21 struct { F0 uint; F1 uint16; F2 string }
type GenS22 struct { F0 uint64; F1 uint64 }
type GenS23 struct { F0 uint; F1 int8 }
type GenS24 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 bool; F1 bool; F2 int32 }
type GenS27 struct { F0 bool }
type GenS28 struct { F0 int8; F1 int32 }
type GenS29 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS30 struct { F0 int64; F1 int32 }
type GenS31 struct { F0 int16 }
type GenS32 struct { F0 int }
type GenS33 struct { F0 bool; F1 int32 }
func genFn0(p0 int, p1 int, p2 bool) uint32 {
    var v6 int64 = -9223372036854775808
    switch (-(p0)) {
    case 9223372036854775807:
        v7 := GenS1{F0: (p0 * (-(p1)))}
        _ = v7
    case -2147483648:
        v8 := []int{(^(p1))}
        _ = v8
    case 247:
        var v9 int32 = 209
        _ = v9
    default:
        p0 = 32768
    }
    p1 = (p1 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 29
}

func main() {
    v1 := []string{}
    var v2 uint8 = 128
    var v3 uint16 = 32767
    var v4 int8 = -1
    var v5 int8 = v4
    switch "hello" {
    case "hello":
        switch 127 {
        case 42:
            if false {
                switch -2147483648 {
                case -2147483648:
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v11 := GenS2{F0: v3, F1: 93, F2: 42}
                        var v12 float32 = -2.25
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                default:
                    if (uint(v2) > 1) {
                        _ = len(v1)
                        for fi13 := 0; fi13 < 4; fi13++ {
                            for fi14 := 0; fi14 < 10; fi14++ {
                                v5 = -128
                                fmt.Println(v5)
                                if (false || true) {
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                                _ = fi14
                            }
                            v3 = 42
                            var v15 bool = true
                            _ = fi13
                            _ = v15
                        }
                    }
                }
            }
        case 9223372036854775807:
            if (-224 >= -212) {
                v16 := GenS2{F0: (v3 + v3), F1: v3, F2: 9223372036854775807}
                if (v3 <= v3) {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        fmt.Println(fi17, fi17, v5)
                        _ = fi17
                    }
                } else {
                    v5 -= (v5 & -128)
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v19 := GenS3{F0: (v4 + 42), F1: fi18, F2: 65535}
                        v5 -= (v5 ^ v5)
                        v4 = v5
                        _ = fi18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 uint32 = genFn0(fi20, fi20, (216 == v2))
                        v2 += v2
                        _ = fi20
                        _ = v21
                    }
                }
                var v22 int = 579
                _ = v16
                _ = v22
            } else {
                v3 += 42
                v3 = (v3 * (v3 % (((v3 + v3)) | 1)))
                if false {
                    v1 = append(v1, "a")
                } else {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        if ((fi23 % ((fi23) | 1)) < (fi23 | fi23)) {
                            v5 = int8(v2)
                            switch "\n\t" {
                            case "":
                                if false {
                                    v4 = (^(v4))
                                } else {
                                    v4 -= v4
                                    v3 = v3
                                }
                            }
                        }
                        if (!true) {
                            v1 = append(v1, "foo bar")
                            if (false && false) {
                                var v24 uint32 = 554
                                v4 = 42
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 9; fi25++ {
                                v2 += (v2 | 106)
                                if true {
                                    switch 32768 {
                                    case -613:
                                        var v26 int32 = 2147483647
                                        _ = v26
                                    case -936:
                                        if true {
                                            if true {
                                                if (!true) {
                                                    v27 := GenS1{F0: -258}
                                                    fmt.Println(v2, v2, v4)
                                                    if (2 < 4294967295) {
                                                        v1 = append(v1, ("a" + "hello"))
                                                        v4 += 1
                                                    } else {
                                                        if (-128 >= v5) {
                                                            v5 -= v5
                                                            var v28 int = (fi25 + (fi23 ^ (fi23 / (((fi23 ^ (fi23 >> ((214) & 63)))) | 1))))
                                                            _ = v28
                                                        }
                                                        fmt.Println(v3, fi25)
                                                        fmt.Println(v5, v2)
                                                        if ("a" > ("a" + (("\n\t" + "\n\t") + ""))) {
                                                            v3 += (v3 % (((v3 / ((2) | 1))) | 1))
                                                            fmt.Println(v2, v5, v2)
                                                            v1 = append(v1, ("hello" + "foo bar"))
                                                        } else {
                                                            v2 -= 2
                                                            v29 := GenS4{F0: 419, F1: (v2 ^ uint8(v3))}
                                                            _ = v29
                                                        }
                                                        if false {
                                                            switch ("" + (("a" + "foo bar") + ((("hello" + ("a" + "\n\t")) + "hello") + ("a" + "")))) {
                                                            case "a":
                                                                var v30 bool = false
                                                                _ = v30
                                                            case "hello":
                                                                v4 = (v4 & (v5 / ((100) | 1)))
                                                                var v31 int = fi25
                                                                _ = v31
                                                            case "":
                                                                var v32 uint32 = 2147483648
                                                                _ = v32
                                                            default:
                                                                var v33 uint8 = 2
                                                                _ = v33
                                                            }
                                                            switch fi25 {
                                                            case -2147483648:
                                                                var v34 int = (fi23 % (((fi23 / (((fi25 % ((-9223372036854775808) | 1))) | 1))) | 1))
                                                                _ = v34
                                                            default:
                                                                if true {
                                                                    switch (fi23 + fi25) {
                                                                    case 128:
                                                                        v35 := GenS5{F0: 128, F1: 0}
                                                                        _ = v35
                                                                    case 996:
                                                                        if (42 > 65535) {
                                                                            v1 = append(v1, "foo bar")
                                                                            fmt.Println(v3)
                                                                            v5 += (v4 ^ v4)
                                                                            v4 = v5
                                                                            if false {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = "\n\t"
                                                                                }
                                                                            }
                                                                            var v36 bool = (1 < -908)
                                                                            _ = v36
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v37 uint16 = v3
                                                                        v38 := GenS1{F0: fi23}
                                                                        _ = v37
                                                                        _ = v38
                                                                    } else {
                                                                        fmt.Println(v4, v4)
                                                                        fmt.Println(v2)
                                                                        var v39 uint32 = 65535
                                                                        _ = v39
                                                                    }
                                                                    switch fi23 {
                                                                    case 1:
                                                                        fmt.Println(fi25, v5)
                                                                    case 2147483647:
                                                                        var v40 float64 = 1e10
                                                                        _ = v40
                                                                    case 924:
                                                                        if (65535 != 18446744073709551615) {
                                                                            switch -594 {
                                                                            case 2147483647:
                                                                                if true {
                                                                                    v3 -= 255
                                                                                    if (!false) {
                                                                                        if false {
                                                                                            if ((0.0 / -2.25) >= (0.0 - 1.5)) {
                                                                                                var v41 int = -1
                                                                                                _ = v41
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v42 int8 = (^(v4))
                                                                                                    v4 -= v4
                                                                                                    _ = v42
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v5 -= v4
                                                                                                    } else {
                                                                                                        switch ("\n\t" + "hello") {
                                                                                                        case "\n\t":
                                                                                                            if (false && (!(-306 >= 90))) {
                                                                                                                _ = len(v1)
                                                                                                                var v43 bool = (((!(false || true)) || ((!(969 <= (^(v3)))) && false)) || (true || (!(!(1 > uint(v2))))))
                                                                                                                _ = v43
                                                                                                            }
                                                                                                        default:
                                                                                                            if len(v1) > 0 {
                                                                                                                v1[0] = "\n\t"
                                                                                                            }
                                                                                                        }
                                                                                                        var v44 bool = (3.14159 < 0.1)
                                                                                                        v2 += 2
                                                                                                        _ = v44
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v5 = v4
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                var v45 int16 = 32767
                                                                                                _ = v45
                                                                                            case "foo bar":
                                                                                                var v46 float32 = (0.0 * 0.0)
                                                                                                _ = v46
                                                                                            case "a":
                                                                                                if false {
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                    var v47 float64 = (((1e10 - -2.25) / 1.0) * 0.0)
                                                                                                    switch "\n\t" {
                                                                                                    case "a":
                                                                                                        var v48 uint16 = (v3 << ((42) & 63))
                                                                                                        _ = v48
                                                                                                    case "foo bar":
                                                                                                        var v49 int32 = 2147483647
                                                                                                        _ = v49
                                                                                                    default:
                                                                                                        var v50 float32 = ((float32(v47) * ((0.1 - float32((float64((((float32(v47) * 0.0) / 0.1) * 0.0)) * float64(float32(v47))))) / (-2.25 - (-2.25 - -2.25)))) * 1e10)
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    _ = v47
                                                                                                }
                                                                                            }
                                                                                            if (-1 > 1) {
                                                                                                if (!true) {
                                                                                                    if false {
                                                                                                        var v51 string = "a"
                                                                                                        if (796 == 4294967295) {
                                                                                                            v4 -= (-(v5))
                                                                                                        } else {
                                                                                                            v52 := GenS4{F0: -9223372036854775808, F1: (v2 << ((18446744073709551615) & 63))}
                                                                                                            _ = v52
                                                                                                        }
                                                                                                        _ = v51
                                                                                                    } else {
                                                                                                        switch fi25 {
                                                                                                        case 32767:
                                                                                                            fmt.Println(fi25)
                                                                                                        case 1:
                                                                                                            var v53 string = (("hello" + "\n\t") + "")
                                                                                                            _ = v53
                                                                                                        case 127:
                                                                                                            if len(v1) > 0 {
                                                                                                                v1[0] = "\n\t"
                                                                                                            }
                                                                                                        default:
                                                                                                            v54 := GenS6{F0: "foo bar"}
                                                                                                            _ = v54
                                                                                                        }
                                                                                                        if (v3 >= (v3 ^ (^(v3)))) {
                                                                                                            v5 += (^(v5))
                                                                                                            _ = len(v1)
                                                                                                            v55 := GenS4{F0: -511, F1: (v2 & 15)}
                                                                                                            v56 := GenS7{F0: -251, F1: 235}
                                                                                                            _ = v55
                                                                                                            _ = v56
                                                                                                        }
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            if (false && false) {
                                                                                                                switch fi25 {
                                                                                                                case 572:
                                                                                                                    v57 := GenS8{F0: 2147483648, F1: fi25}
                                                                                                                    _ = v57
                                                                                                                case 32768:
                                                                                                                    var v58 uint8 = uint8(v3)
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            _ = len(v1)
                                                                                                        default:
                                                                                                            v1 = append(v1, "foo bar")
                                                                                                            var v59 int64 = 0
                                                                                                            var v60 string = ((("hello" + "") + (("a" + "hello") + ("\n\t" + ("a" + "hello")))) + "a")
                                                                                                            _ = v59
                                                                                                            _ = v60
                                                                                                        }
                                                                                                        v61 := GenS1{F0: (fi25 * fi25)}
                                                                                                        _ = v61
                                                                                                    }
                                                                                                    var v62 int32 = -1
                                                                                                    _ = v62
                                                                                                } else {
                                                                                                    v63 := GenS5{F0: 0, F1: -1}
                                                                                                    if false {
                                                                                                        v4 = int8(v2)
                                                                                                    }
                                                                                                    if true {
                                                                                                        v64 := GenS9{F0: "foo bar", F1: int8(v2), F2: (v2 << ((0) & 63))}
                                                                                                        switch ("a" + "hello") {
                                                                                                        case "\n\t":
                                                                                                            var v65 int = (fi23 - 9223372036854775807)
                                                                                                            _ = v65
                                                                                                        case "":
                                                                                                            v2 = v2
                                                                                                        case "hello":
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v64
                                                                                                    } else {
                                                                                                        var v66 bool = (32 < 396)
                                                                                                        _ = v66
                                                                                                    }
                                                                                                    _ = v63
                                                                                                }
                                                                                                var v67 float32 = (0.1 - 3.14159)
                                                                                                var v68 float64 = 0.1
                                                                                                _ = v67
                                                                                                _ = v68
                                                                                            } else {
                                                                                                var v69 uint32 = 0
                                                                                                _ = v69
                                                                                            }
                                                                                            v4 -= v4
                                                                                        } else {
                                                                                            var v70 float64 = 1e10
                                                                                            var v71 int8 = v5
                                                                                            if true {
                                                                                                var v72 uint64 = 1
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                _ = v72
                                                                                            }
                                                                                            switch ("foo bar" + ("\n\t" + "foo bar")) {
                                                                                            case "a":
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            case "foo bar":
                                                                                                fmt.Println(v71, v71, v5)
                                                                                            }
                                                                                            _ = len(v1)
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        var v73 float64 = (1.5 * 0.0)
                                                                                        if true {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = ("\n\t" + ("hello" + "a"))
                                                                                            }
                                                                                            if ((-(fi25)) <= fi25) {
                                                                                                fmt.Println(v2, fi23, v2)
                                                                                                var v74 float32 = 1.5
                                                                                                _ = v74
                                                                                            } else {
                                                                                                if (1 > 659) {
                                                                                                    if (false && ("foo bar" == ("\n\t" + ("a" + "foo bar")))) {
                                                                                                        v3 -= v3
                                                                                                        switch fi25 {
                                                                                                        case 127:
                                                                                                            var v75 uint16 = v3
                                                                                                            _ = v75
                                                                                                        case -9223372036854775808:
                                                                                                            v1 = append(v1, "foo bar")
                                                                                                        }
                                                                                                        var v76 int32 = int32(v3)
                                                                                                        _ = v76
                                                                                                    } else {
                                                                                                        if (v5 >= (-(v5))) {
                                                                                                            v1 = append(v1, "\n\t")
                                                                                                        } else {
                                                                                                            var v77 int = (fi25 * (^(fi23)))
                                                                                                            _ = len(v1)
                                                                                                            if true {
                                                                                                                v77 = -646
                                                                                                                v78 := GenS10{F0: 9223372036854775807}
                                                                                                                if (((v73 / 1.0) != v73) || (false || false)) {
                                                                                                                    v79 := GenS11{F0: 1000}
                                                                                                                    v2 += 127
                                                                                                                    _ = v79
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v73 -= v73
                                                                                                                        switch (fi25 % (((fi25 - v77)) | 1)) {
                                                                                                                        case 2147483647:
                                                                                                                            v4 = v4
                                                                                                                        case -427:
                                                                                                                            switch ("\n\t" + "\n\t") {
                                                                                                                            case "a":
                                                                                                                                var v80 string = ""
                                                                                                                                _ = v80
                                                                                                                            }
                                                                                                                        case 32767:
                                                                                                                            switch (("hello" + ("hello" + "a")) + "hello") {
                                                                                                                            case "a":
                                                                                                                                v81 := GenS3{F0: -96, F1: v77, F2: v3}
                                                                                                                                _ = v81
                                                                                                                            case "hello":
                                                                                                                                v3 = 518
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v78
                                                                                                            }
                                                                                                            _ = v77
                                                                                                        }
                                                                                                        var v82 int8 = v5
                                                                                                        var v83 float32 = 1e10
                                                                                                        _ = v82
                                                                                                        _ = v83
                                                                                                    }
                                                                                                    fmt.Println(v4, v3)
                                                                                                    v84 := GenS7{F0: (fi25 / (((fi23 ^ fi25)) | 1)), F1: 0}
                                                                                                    _ = v84
                                                                                                } else {
                                                                                                    if false {
                                                                                                        _ = len(v1)
                                                                                                    } else {
                                                                                                        v1 = append(v1, "hello")
                                                                                                        v85 := GenS3{F0: v5, F1: (fi23 | (fi23 / ((fi23) | 1))), F2: (v3 % ((255) | 1))}
                                                                                                        var v86 uint = 65535
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v87 uint = 18446744073709551615
                                                                                            _ = v87
                                                                                        } else {
                                                                                            v88 := GenS3{F0: -1, F1: fi25, F2: (v3 + v3)}
                                                                                            switch fi25 {
                                                                                            case -242:
                                                                                                v89 := GenS12{F0: 65535, F1: (0 <= 2)}
                                                                                                _ = v89
                                                                                            case 4294967295:
                                                                                                var v90 uint = 2
                                                                                                _ = v90
                                                                                            case -2147483648:
                                                                                                v91 := GenS2{F0: uint16(v2), F1: (v3 + v3), F2: 215}
                                                                                                _ = v91
                                                                                            }
                                                                                            switch 4294967295 {
                                                                                            case -9223372036854775808:
                                                                                                if false {
                                                                                                    var v92 int32 = 0
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    var v93 float32 = -2.25
                                                                                                    _ = v93
                                                                                                }
                                                                                            case 4294967295:
                                                                                                v73 = v73
                                                                                            case -2147483648:
                                                                                                if false {
                                                                                                    v94 := GenS4{F0: 599, F1: 2}
                                                                                                    v3 += (v3 % (((v3 << ((42) & 63))) | 1))
                                                                                                    v73 = v73
                                                                                                    _ = v94
                                                                                                }
                                                                                            default:
                                                                                                if (0 <= 7) {
                                                                                                    v5 = v5
                                                                                                }
                                                                                            }
                                                                                            _ = v88
                                                                                        }
                                                                                        _ = v73
                                                                                    }
                                                                                    v3 -= 709
                                                                                    var v95 bool = true
                                                                                    v95 = (-63 == v4)
                                                                                    _ = v95
                                                                                } else {
                                                                                    v5 = (v4 & v5)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            switch (fi25 ^ (fi25 ^ fi25)) {
                                                                            case 664:
                                                                                if (false && (419 > 0)) {
                                                                                    fmt.Println(v4, v5)
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = "hello"
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (3.14159 <= (0.0 - -2.25)) {
                                                                                fmt.Println(v2, v2, v2)
                                                                                v96 := GenS4{F0: 1, F1: v2}
                                                                                _ = v96
                                                                            } else {
                                                                                var v97 float64 = 1e10
                                                                                _ = len(v1)
                                                                                var v98 uint16 = (v3 >> ((18446744073709551615) & 63))
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                            v5 += (v5 / ((-51) | 1))
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            switch fi25 {
                                                                            case 0:
                                                                                v2 -= (v2 * (v2 ^ v2))
                                                                            case 501:
                                                                                v1 = append(v1, "hello")
                                                                            case 2147483647:
                                                                                if false {
                                                                                    var v99 uint16 = v3
                                                                                    fmt.Println(v4)
                                                                                    _ = v99
                                                                                } else {
                                                                                    switch ("" + ("hello" + "a")) {
                                                                                    case "hello":
                                                                                        _ = len(v1)
                                                                                    case "foo bar":
                                                                                        v4 -= (v5 & (-(v5)))
                                                                                    case "":
                                                                                        if true {
                                                                                            if false {
                                                                                                v100 := GenS13{F0: 1, F1: "hello", F2: (!(("a" + "hello") >= "foo bar"))}
                                                                                                _ = v100
                                                                                            }
                                                                                            v4 -= (^(v5))
                                                                                        } else {
                                                                                            v1 = append(v1, (("hello" + "\n\t") + "\n\t"))
                                                                                            v2 -= v2
                                                                                        }
                                                                                    default:
                                                                                        v101 := GenS10{F0: 920}
                                                                                        _ = v101
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if (false || true) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = ("a" + ("\n\t" + "foo bar"))
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        var v102 float32 = 3.14159
                                                                                        _ = v102
                                                                                    case "\n\t":
                                                                                        v1 = append(v1, "\n\t")
                                                                                    case "hello":
                                                                                        fmt.Println(v5, v3, v4)
                                                                                    default:
                                                                                        if false {
                                                                                            v103 := GenS14{F0: 0}
                                                                                            var v104 uint64 = 711
                                                                                            var v105 string = (("\n\t" + ("" + "foo bar")) + ((("" + "hello") + "") + "hello"))
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                            _ = v105
                                                                                        } else {
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                switch fi23 {
                                                                                                case 2147483647:
                                                                                                    v106 := GenS13{F0: 0, F1: "hello", F2: true}
                                                                                                    _ = v106
                                                                                                case 0:
                                                                                                    v5 += (^(v5))
                                                                                                case -346:
                                                                                                    fmt.Println(v3, fi23)
                                                                                                }
                                                                                            case "":
                                                                                                v107 := GenS15{F0: 0, F1: (v4 | v4)}
                                                                                                if true {
                                                                                                    fmt.Println(fi23, v2, fi23)
                                                                                                    v3 = v3
                                                                                                    if (-458 > 42) {
                                                                                                        v108 := GenS14{F0: 18446744073709551615}
                                                                                                        switch (-(fi23)) {
                                                                                                        case 9223372036854775807:
                                                                                                            if true {
                                                                                                                var v109 int64 = -1
                                                                                                                var v110 uint16 = 32767
                                                                                                                _ = v109
                                                                                                                _ = v110
                                                                                                            } else {
                                                                                                                if len(v1) >= 2 {
                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                }
                                                                                                            }
                                                                                                        case 128:
                                                                                                            switch 615 {
                                                                                                            case 4294967295:
                                                                                                                v5 += (v5 & (v4 / ((v5) | 1)))
                                                                                                            case 0:
                                                                                                                fmt.Println(fi23, v3)
                                                                                                            default:
                                                                                                                v111 := GenS12{F0: (v3 ^ 566), F1: (4294967295 != 18446744073709551615)}
                                                                                                                _ = v111
                                                                                                            }
                                                                                                        default:
                                                                                                            if false {
                                                                                                                var v112 float32 = (float32(0.0) + float32(1.5))
                                                                                                                v4 = v4
                                                                                                                switch "\n\t" {
                                                                                                                case "hello":
                                                                                                                    v113 := GenS16{F0: 1}
                                                                                                                    _ = v113
                                                                                                                case "foo bar":
                                                                                                                    if false {
                                                                                                                        fmt.Println(fi23, fi25)
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    var v114 int = -228
                                                                                                                    _ = v114
                                                                                                                }
                                                                                                                _ = v112
                                                                                                            }
                                                                                                            if (-940 < -951) {
                                                                                                                if len(v1) >= 2 {
                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                }
                                                                                                                if (2147483647 < 1) {
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        switch fi23 {
                                                                                                                        case 127:
                                                                                                                            var v115 uint16 = (v3 ^ v3)
                                                                                                                            _ = v115
                                                                                                                        case -9223372036854775808:
                                                                                                                            switch (("a" + "a") + "foo bar") {
                                                                                                                            case "\n\t":
                                                                                                                                var v116 uint8 = 255
                                                                                                                                _ = v116
                                                                                                                            case "a":
                                                                                                                                fmt.Println(v2)
                                                                                                                            default:
                                                                                                                                v117 := GenS7{F0: fi25, F1: 18446744073709551615}
                                                                                                                                _ = v117
                                                                                                                            }
                                                                                                                        case 935:
                                                                                                                            v118 := GenS7{F0: fi25, F1: 2}
                                                                                                                            _ = v118
                                                                                                                        }
                                                                                                                        v5 = 1
                                                                                                                        if (!true) {
                                                                                                                            fmt.Println(fi25)
                                                                                                                            v3 = (^(v3))
                                                                                                                            fmt.Println(v2, v4)
                                                                                                                        } else {
                                                                                                                            var v119 int32 = 0
                                                                                                                            _ = v119
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        fmt.Println(v5, v5)
                                                                                                                        fmt.Println(fi23, fi23, fi23)
                                                                                                                        v120 := GenS17{F0: (v2 + v2), F1: 42, F2: 650}
                                                                                                                        v1 = append(v1, "a")
                                                                                                                        var v121 uint32 = 189
                                                                                                                        _ = v120
                                                                                                                        _ = v121
                                                                                                                    }
                                                                                                                }
                                                                                                                if true {
                                                                                                                    switch (("foo bar" + ("a" + (("a" + "a") + (("hello" + "\n\t") + "hello")))) + ("a" + "foo bar")) {
                                                                                                                    case "a":
                                                                                                                        switch (fi25 | (fi25 << ((415) & 63))) {
                                                                                                                        case -1:
                                                                                                                            if (("\n\t" >= "hello") || (21 == v2)) {
                                                                                                                                fmt.Println(v4, v4, v4)
                                                                                                                            }
                                                                                                                        case 32768:
                                                                                                                            var v122 int = fi25
                                                                                                                            _ = v122
                                                                                                                        case 4294967295:
                                                                                                                            var v123 int = fi25
                                                                                                                            _ = v123
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        _ = len(v1)
                                                                                                                    default:
                                                                                                                        var v124 uint64 = 1
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                    var v125 uint16 = (^(v3))
                                                                                                                    _ = v125
                                                                                                                }
                                                                                                            }
                                                                                                            _ = len(v1)
                                                                                                        }
                                                                                                        var v126 string = (("a" + "") + "a")
                                                                                                        _ = v108
                                                                                                        _ = v126
                                                                                                    } else {
                                                                                                        v5 += 100
                                                                                                        v127 := GenS18{F0: (v3 ^ 701), F1: 951, F2: (v4 * int8(v3))}
                                                                                                        v3 = 310
                                                                                                        _ = v127
                                                                                                    }
                                                                                                }
                                                                                                switch fi25 {
                                                                                                case -1:
                                                                                                    if false {
                                                                                                        v128 := GenS19{F0: (v2 - 255)}
                                                                                                        if false {
                                                                                                            _ = len(v1)
                                                                                                            _ = len(v1)
                                                                                                        }
                                                                                                        v3 = v3
                                                                                                        _ = v128
                                                                                                    } else {
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            v129 := GenS3{F0: (v5 / ((127) | 1)), F1: (fi25 ^ (fi25 ^ (fi23 - (fi25 ^ (fi25 & fi25))))), F2: (v3 | 32767)}
                                                                                                            var v130 float32 = float32((1.5 + 3.14159))
                                                                                                            var v131 uint16 = v3
                                                                                                            _ = v129
                                                                                                            _ = v130
                                                                                                            _ = v131
                                                                                                        }
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = ""
                                                                                                        }
                                                                                                        var v132 int32 = -1
                                                                                                        var v133 float64 = 0.0
                                                                                                        _ = v132
                                                                                                        _ = v133
                                                                                                    }
                                                                                                case 42:
                                                                                                    v2 -= v2
                                                                                                }
                                                                                                fmt.Println(fi25)
                                                                                                if (1.5 > (0.1 * 1e10)) {
                                                                                                    var v134 string = (("foo bar" + ("foo bar" + ("\n\t" + "hello"))) + "hello")
                                                                                                    v3 = v3
                                                                                                    _ = v134
                                                                                                } else {
                                                                                                    v4 = 1
                                                                                                    v135 := GenS20{F0: 18446744073709551615, F1: (fi25 | int(v5)), F2: -1}
                                                                                                    var v136 float64 = -2.25
                                                                                                    _ = v135
                                                                                                    _ = v136
                                                                                                }
                                                                                                switch ("foo bar" + ("foo bar" + "hello")) {
                                                                                                case "a":
                                                                                                    var v137 uint = 820
                                                                                                    _ = v137
                                                                                                case "\n\t":
                                                                                                    v138 := GenS21{F0: 42, F1: (v3 * v3), F2: "foo bar"}
                                                                                                    _ = v138
                                                                                                case "foo bar":
                                                                                                    if (-9223372036854775808 <= 9223372036854775807) {
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                switch (^(fi25)) {
                                                                                                case -19:
                                                                                                    v5 -= (v5 ^ -1)
                                                                                                case 1:
                                                                                                    v2 = (v2 % ((255) | 1))
                                                                                                    v2 -= v2
                                                                                                    switch int(v5) {
                                                                                                    case 0:
                                                                                                        var v139 uint64 = 784
                                                                                                        _ = v139
                                                                                                    }
                                                                                                case 32767:
                                                                                                    var v140 uint64 = 1
                                                                                                    _ = v140
                                                                                                default:
                                                                                                    v141 := GenS22{F0: 1, F1: 42}
                                                                                                    _ = v141
                                                                                                }
                                                                                                _ = v107
                                                                                            case "\n\t":
                                                                                                v5 -= v4
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v142 int8 = v5
                                                                                    switch (-(fi25)) {
                                                                                    case 693:
                                                                                        v2 -= 1
                                                                                    case 848:
                                                                                        if (true || (false && (v2 == (v2 - (v2 - v2))))) {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = "hello"
                                                                                            }
                                                                                            var v143 int16 = 1
                                                                                            switch (fi23 * fi25) {
                                                                                            case 9223372036854775807:
                                                                                                var v144 bool = true
                                                                                                fmt.Println(fi25, v142, v4)
                                                                                                _ = v144
                                                                                            }
                                                                                            _ = v143
                                                                                        } else {
                                                                                            var v145 float64 = 1e10
                                                                                            if ((!true) && false) {
                                                                                                _ = len(v1)
                                                                                                if (true || (!false)) {
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = ""
                                                                                                    }
                                                                                                    var v146 int = (fi23 ^ (fi23 | fi23))
                                                                                                    _ = v146
                                                                                                } else {
                                                                                                    if (!(("a" + "a") <= "foo bar")) {
                                                                                                        fmt.Println(v2)
                                                                                                        var v147 string = "a"
                                                                                                        _ = v147
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                switch fi25 {
                                                                                                case -2147483648:
                                                                                                    var v148 uint8 = (v2 & (v2 ^ (v2 + (v2 * (v2 & (^(v2)))))))
                                                                                                    _ = v148
                                                                                                default:
                                                                                                    switch 208 {
                                                                                                    case -2147483648:
                                                                                                        var v149 int16 = int16(v3)
                                                                                                        _ = v149
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v145
                                                                                        }
                                                                                    case 258:
                                                                                        var v150 uint64 = 18446744073709551615
                                                                                        _ = v150
                                                                                    default:
                                                                                        v4 = 100
                                                                                    }
                                                                                    _ = v142
                                                                                }
                                                                            }
                                                                            v2 -= v2
                                                                        } else {
                                                                            var v151 bool = ((!(!((v3 % ((1) | 1)) <= v3))) && (((fi23 % ((670) | 1)) > fi25) || true))
                                                                            switch -892 {
                                                                            case 128:
                                                                                v3 += 1
                                                                            default:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            if v151 {
                                                                                fmt.Println(fi23, v2)
                                                                            }
                                                                            fmt.Println(v5)
                                                                            _ = v151
                                                                        }
                                                                    }
                                                                } else {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    var v152 int8 = (v4 + v4)
                                                                    v5 -= v152
                                                                    v4 = (v4 % (((v5 << ((697) & 63))) | 1))
                                                                    _ = v152
                                                                }
                                                            }
                                                        }
                                                        var v153 uint8 = 1
                                                        var v154 int32 = -540
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                    _ = v27
                                                } else {
                                                    var v155 int16 = 0
                                                    v156 := GenS23{F0: 4294967295, F1: v4}
                                                    switch "" {
                                                    case "":
                                                        v3 += 637
                                                    case "hello":
                                                        v157 := GenS24{F0: v2, F1: uint64(v3), F2: 54}
                                                        _ = v157
                                                    case "\n\t":
                                                        if ((-15 > (-(fi25))) || (false && ("\n\t" == "a"))) {
                                                            v3 -= (v3 & v3)
                                                            if (-1 >= -2147483648) {
                                                                fmt.Println(fi23)
                                                            } else {
                                                                if (!((fi25 >> ((65535) & 63)) <= (fi23 % (((^(fi25))) | 1)))) {
                                                                    var v158 uint16 = v3
                                                                    v1 = append(v1, ("foo bar" + ""))
                                                                    if true {
                                                                        switch (fi23 << ((4294967295) & 63)) {
                                                                        case -9223372036854775808:
                                                                            var v159 int64 = -971
                                                                            _ = v159
                                                                        case 32767:
                                                                            v4 -= v5
                                                                            v160 := GenS25{F0: -1}
                                                                            _ = v160
                                                                        case -478:
                                                                            var v161 bool = false
                                                                            _ = v161
                                                                        }
                                                                        v4 = (v4 % ((v5) | 1))
                                                                        v5 = 100
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            switch -2147483648 {
                                                                            case 128:
                                                                                v162 := GenS3{F0: -39, F1: (^(fi23)), F2: (v3 + v3)}
                                                                                _ = v162
                                                                            case -349:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = "foo bar"
                                                                                }
                                                                            case 1:
                                                                                v1 = append(v1, "foo bar")
                                                                            default:
                                                                                var v163 string = ("foo bar" + ("foo bar" + "\n\t"))
                                                                                _ = v163
                                                                            }
                                                                            v2 += 128
                                                                            switch fi23 {
                                                                            case 1:
                                                                                var v164 int8 = (^(v5))
                                                                                _ = v164
                                                                            case -788:
                                                                                _ = len(v1)
                                                                            default:
                                                                                if (2 > uint(v2)) {
                                                                                    v165 := GenS8{F0: genFn0((fi25 % ((42) | 1)), (fi25 & int(v155)), (!true)), F1: fi25}
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = "\n\t"
                                                                                    }
                                                                                    if (!((v2 % ((2) | 1)) == (v2 ^ (v2 & 128)))) {
                                                                                        _ = len(v1)
                                                                                        if (1e10 >= (float64(float32(0.0)) * float64(0.0))) {
                                                                                            if true {
                                                                                                if true {
                                                                                                    v5 = (^(v5))
                                                                                                    v1 = append(v1, "")
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = "foo bar"
                                                                                                    }
                                                                                                    var v166 int8 = v5
                                                                                                    var v167 int16 = 1
                                                                                                    _ = v166
                                                                                                    _ = v167
                                                                                                } else {
                                                                                                    switch 22 {
                                                                                                    case 773:
                                                                                                        v4 = v4
                                                                                                    case 128:
                                                                                                        v5 = (v4 / ((v4) | 1))
                                                                                                    default:
                                                                                                        var v168 uint64 = 2
                                                                                                        _ = v168
                                                                                                    }
                                                                                                    v155 += v155
                                                                                                    v1 = append(v1, "hello")
                                                                                                }
                                                                                                fmt.Println(v3)
                                                                                                _ = len(v1)
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v3, v2)
                                                                                    }
                                                                                    _ = v165
                                                                                } else {
                                                                                    var v169 uint16 = v158
                                                                                    _ = v169
                                                                                }
                                                                            }
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ""
                                                                        }
                                                                        var v170 uint8 = (v2 - v2)
                                                                        _ = v170
                                                                    }
                                                                    _ = v158
                                                                }
                                                                v5 += v5
                                                            }
                                                            if true {
                                                                if (false && (int64(fi25) != -786)) {
                                                                    if (int64(fi25) > 0) {
                                                                        var v171 int64 = -9223372036854775808
                                                                        _ = v171
                                                                    }
                                                                    _ = len(v1)
                                                                    var v172 uint16 = v3
                                                                    _ = v172
                                                                }
                                                                _ = len(v1)
                                                                v173 := GenS5{F0: v2, F1: -126}
                                                                if (!("\n\t" > "")) {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        v174 := GenS13{F0: 18446744073709551615, F1: "", F2: ((true || (!false)) && false)}
                                                                        _ = v174
                                                                    case "":
                                                                        if true {
                                                                            var v175 int32 = 419
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v176 float32 = 0.0
                                                                            var v177 bool = false
                                                                            switch (fi25 | (fi25 << ((4294967295) & 63))) {
                                                                            case 0:
                                                                                v178 := GenS15{F0: int32(v3), F1: (v4 & -95)}
                                                                                _ = v178
                                                                            case -1:
                                                                                var v179 int64 = 0
                                                                                _ = v179
                                                                            }
                                                                            if false {
                                                                                v4 = (-(v4))
                                                                                fmt.Println(v3, v177)
                                                                            } else {
                                                                                v180 := GenS26{F0: v177, F1: (-671 != 1), F2: int32(v5)}
                                                                                v155 = (v155 * (v155 >> ((1) & 63)))
                                                                                if v177 {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = ("hello" + "a")
                                                                                    }
                                                                                } else {
                                                                                    v1 = append(v1, "")
                                                                                    fmt.Println(v3)
                                                                                    v181 := GenS21{F0: 2, F1: v3, F2: ((("foo bar" + "foo bar") + "hello") + "a")}
                                                                                    _ = v181
                                                                                }
                                                                                _ = v180
                                                                            }
                                                                            switch (fi23 << ((4294967295) & 63)) {
                                                                            case 128:
                                                                                v155 = v155
                                                                            case -719:
                                                                                var v182 string = ("\n\t" + "foo bar")
                                                                                _ = v182
                                                                            case 4294967295:
                                                                                v1 = append(v1, "a")
                                                                            default:
                                                                                var v183 uint8 = 128
                                                                                _ = v183
                                                                            }
                                                                            _ = v175
                                                                            _ = v176
                                                                            _ = v177
                                                                        } else {
                                                                            var v184 uint = 255
                                                                            _ = v184
                                                                        }
                                                                        var v185 uint64 = 4294967295
                                                                        switch "" {
                                                                        case "":
                                                                            v186 := GenS14{F0: v185}
                                                                            _ = v186
                                                                        case "\n\t":
                                                                            var v187 uint8 = (v2 | uint8(v3))
                                                                            _ = v187
                                                                        case "a":
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                var v188 float32 = 1e10
                                                                                _ = v188
                                                                            case "":
                                                                                fmt.Println(v4, fi25, v4)
                                                                                v185 -= v185
                                                                            }
                                                                        }
                                                                        _ = v185
                                                                    case "foo bar":
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v155 -= v155
                                                                        }
                                                                    }
                                                                    var v189 uint32 = 4294967295
                                                                    var v190 int64 = 9223372036854775807
                                                                    _ = v189
                                                                    _ = v190
                                                                }
                                                                switch 32767 {
                                                                case 2147483647:
                                                                    v1 = append(v1, "")
                                                                    v3 = 614
                                                                default:
                                                                    switch fi23 {
                                                                    case 110:
                                                                        _ = len(v1)
                                                                    default:
                                                                        v191 := GenS13{F0: 65535, F1: "a", F2: (!false)}
                                                                        _ = v191
                                                                    }
                                                                }
                                                                v4 = (v4 + (^(v5)))
                                                                _ = v173
                                                            }
                                                        } else {
                                                            if true {
                                                                v192 := GenS3{F0: (v5 - (v5 - v4)), F1: (fi25 / ((fi25) | 1)), F2: v3}
                                                                v155 = (v155 << ((255) & 63))
                                                                fmt.Println(fi25, fi23)
                                                                switch (-(fi23)) {
                                                                case -2147483648:
                                                                    var v193 uint8 = (v2 + v2)
                                                                    _ = v193
                                                                case 128:
                                                                    fmt.Println(v2, v2)
                                                                default:
                                                                    v194 := GenS27{F0: (!(!(true && (v5 > int8(v2)))))}
                                                                    _ = v194
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = "\n\t"
                                                                }
                                                                var v195 bool = true
                                                                var v196 int8 = v5
                                                                v197 := GenS13{F0: 485, F1: "hello", F2: false}
                                                                _ = v192
                                                                _ = v195
                                                                _ = v196
                                                                _ = v197
                                                            }
                                                            v1 = append(v1, "")
                                                        }
                                                    default:
                                                        v198 := GenS14{F0: 1}
                                                        _ = v198
                                                    }
                                                    _ = v155
                                                    _ = v156
                                                }
                                                switch ("foo bar" + "a") {
                                                case "a":
                                                    v5 += (^(v5))
                                                default:
                                                    var v199 uint64 = uint64(v2)
                                                    _ = v199
                                                }
                                                fmt.Println(v5, v4)
                                            } else {
                                                var v200 float32 = 0.1
                                                _ = v200
                                            }
                                        } else {
                                            fmt.Println(v2, v2)
                                            v201 := GenS1{F0: fi23}
                                            var v202 int32 = -1
                                            if (!(true || (2 <= 1))) {
                                                if false {
                                                    var v203 int16 = -32768
                                                    v5 = (v5 + v5)
                                                    var v204 string = ("a" + "a")
                                                    _ = v203
                                                    _ = v204
                                                } else {
                                                    v205 := GenS22{F0: uint64(v3), F1: 2}
                                                    switch fi25 {
                                                    case 128:
                                                        v4 = 74
                                                    case 1:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if (!false) {
                                                            fmt.Println(v5, v4)
                                                        } else {
                                                            var v206 int16 = 1
                                                            _ = len(v1)
                                                            fmt.Println(v2, v2, v206)
                                                            _ = v206
                                                        }
                                                    case -892:
                                                        v207 := GenS13{F0: uint(v3), F1: "foo bar", F2: true}
                                                        _ = v207
                                                    default:
                                                        v4 = 42
                                                    }
                                                    if false {
                                                        v202 -= v202
                                                        v1 = append(v1, "")
                                                        if true {
                                                            switch ("\n\t" + ("foo bar" + "hello")) {
                                                            case "\n\t":
                                                                v3 = v3
                                                            case "hello":
                                                                if true {
                                                                    switch fi25 {
                                                                    case 1:
                                                                        switch (fi23 ^ fi25) {
                                                                        case 32768:
                                                                            var v208 int32 = int32(fi25)
                                                                            var v209 string = "foo bar"
                                                                            _ = v208
                                                                            _ = v209
                                                                        case 4294967295:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = "a"
                                                                            }
                                                                        }
                                                                    case -747:
                                                                        var v210 float64 = float64(float32(3.14159))
                                                                        _ = v210
                                                                    case -2147483648:
                                                                        if true {
                                                                            v3 += v3
                                                                        } else {
                                                                            fmt.Println(v2, v3, fi23)
                                                                            var v211 uint32 = 42
                                                                            switch ((("hello" + "") + "hello") + "a") {
                                                                            case "a":
                                                                                _ = len(v1)
                                                                            case "foo bar":
                                                                                if (fi25 == (-(fi23))) {
                                                                                    v202 -= v202
                                                                                }
                                                                            case "":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ""
                                                                                }
                                                                            default:
                                                                                fmt.Println(v202, fi23)
                                                                            }
                                                                            _ = v211
                                                                        }
                                                                    default:
                                                                        switch fi25 {
                                                                        case -1:
                                                                            v212 := GenS12{F0: (v3 - (v3 + (v3 % ((917) | 1)))), F1: true}
                                                                            var v213 int = (fi23 - 895)
                                                                            fmt.Println(v5, v3)
                                                                            _ = v212
                                                                            _ = v213
                                                                        case 32767:
                                                                            v4 += v5
                                                                        case 42:
                                                                            v5 += v5
                                                                        }
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v3, v3)
                                                            default:
                                                                var v214 int32 = -2147483648
                                                                _ = v214
                                                            }
                                                            var v215 int32 = (v202 | v202)
                                                            _ = v215
                                                        } else {
                                                            _ = len(v1)
                                                            v4 = (v5 ^ v5)
                                                        }
                                                    } else {
                                                        var v216 float32 = 1.5
                                                        if false {
                                                            if len(v1) > 0 {
                                                                v1[0] = ("\n\t" + "")
                                                            }
                                                            v2 -= v2
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = ("\n\t" + "a")
                                                        }
                                                        _ = v216
                                                    }
                                                    _ = v205
                                                }
                                            } else {
                                                if (0 <= -1) {
                                                    var v217 int = (fi25 << ((300) & 63))
                                                    switch (("a" + "hello") + ("a" + "foo bar")) {
                                                    case "":
                                                        if true {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                switch v217 {
                                                                case 162:
                                                                    if false {
                                                                        v4 += (v5 & (v4 / (((^(v5))) | 1)))
                                                                    }
                                                                case 417:
                                                                    if ((!(739 > 275)) && ((1.5 - (0.0 - 0.0)) < 0.0)) {
                                                                        var v218 uint32 = 2147483648
                                                                        v219 := GenS28{F0: -1, F1: (v202 + 44)}
                                                                        v220 := GenS11{F0: 1000}
                                                                        var v221 int = v217
                                                                        _ = v218
                                                                        _ = v219
                                                                        _ = v220
                                                                        _ = v221
                                                                    } else {
                                                                        v222 := GenS29{F0: 2, F1: -74, F2: 437}
                                                                        var v223 uint64 = 42
                                                                        _ = v222
                                                                        _ = v223
                                                                    }
                                                                default:
                                                                    if ((-198 == 42) && false) {
                                                                        v3 = (v3 + (^(v3)))
                                                                        v224 := GenS24{F0: (v2 << ((144) & 63)), F1: 2, F2: (^(v2))}
                                                                        _ = v224
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v202 = (v202 << ((123) & 63))
                                                                switch (fi25 - 2147483647) {
                                                                case 9223372036854775807:
                                                                    fmt.Println(v3, v5, fi23)
                                                                default:
                                                                    if (0.1 == float32(((0.0 / 1.5) - 0.0))) {
                                                                        var v225 bool = (-1 <= -553)
                                                                        v3 -= 435
                                                                        var v226 uint16 = v3
                                                                        _ = v225
                                                                        _ = v226
                                                                    } else {
                                                                        var v227 float64 = 3.14159
                                                                        v228 := GenS8{F0: genFn0(fi25, 42, true), F1: 2147483647}
                                                                        _ = v227
                                                                        _ = v228
                                                                    }
                                                                }
                                                                v229 := GenS30{F0: -1, F1: (v202 / (((v202 ^ (v202 - v202))) | 1))}
                                                                _ = v229
                                                            case "":
                                                                fmt.Println(v2)
                                                            default:
                                                                v230 := GenS31{F0: 0}
                                                                _ = v230
                                                            }
                                                        } else {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if true {
                                                                    v3 -= v3
                                                                } else {
                                                                    switch 42 {
                                                                    case 32768:
                                                                        var v231 int8 = 100
                                                                        _ = v231
                                                                    case 42:
                                                                        var v232 int32 = (v202 % (((v202 ^ 2147483647)) | 1))
                                                                        _ = v232
                                                                    default:
                                                                        v202 -= v202
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        fmt.Println(v2)
                                                                    }
                                                                }
                                                            case "a":
                                                                switch "a" {
                                                                case "hello":
                                                                    if (v202 >= (v202 - (v202 | (v202 % (((v202 - -1)) | 1))))) {
                                                                        v5 += 24
                                                                    }
                                                                case "foo bar":
                                                                    switch ("hello" + "a") {
                                                                    case "a":
                                                                        v2 += v2
                                                                    case "foo bar":
                                                                        var v233 int16 = 32767
                                                                        _ = v233
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            if true {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (("a" + "a") + "")
                                                                                }
                                                                                _ = len(v1)
                                                                                v217 += (fi23 * (v217 & fi23))
                                                                                v202 += 596
                                                                            }
                                                                            v5 += (^(v4))
                                                                        default:
                                                                            var v234 int16 = 412
                                                                            _ = v234
                                                                        }
                                                                    }
                                                                case "":
                                                                    v5 += v4
                                                                default:
                                                                    switch ("foo bar" + (("" + ("foo bar" + ("hello" + "\n\t"))) + (((("hello" + ("hello" + "hello")) + (("" + "hello") + (("hello" + (("\n\t" + "\n\t") + "foo bar")) + "hello"))) + (("a" + "") + "a")) + "foo bar"))) {
                                                                    case "":
                                                                        v1 = append(v1, (("hello" + "a") + "foo bar"))
                                                                    default:
                                                                        switch (fi23 % ((fi25) | 1)) {
                                                                        case 32768:
                                                                            v1 = append(v1, "a")
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                v235 := GenS21{F0: 65535, F1: 1, F2: ""}
                                                                _ = v235
                                                            }
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = ("foo bar" + "hello")
                                                        }
                                                    case "hello":
                                                        v236 := GenS32{F0: 42}
                                                        _ = v236
                                                    }
                                                    switch 0 {
                                                    case 4294967295:
                                                        var v237 uint = 737
                                                        _ = v237
                                                    case -1:
                                                        fmt.Println(fi23, fi25, v217)
                                                    }
                                                    fmt.Println(v5, v4, fi23)
                                                    v238 := GenS4{F0: -89, F1: (v2 >> ((0) & 63))}
                                                    _ = v217
                                                    _ = v238
                                                } else {
                                                    var v239 float64 = 0.0
                                                    _ = v239
                                                }
                                                v240 := GenS31{F0: 32767}
                                                var v241 float64 = 3.14159
                                                _ = v240
                                                _ = v241
                                            }
                                            _ = v201
                                            _ = v202
                                        }
                                    }
                                    var v242 string = "\n\t"
                                    _ = v242
                                }
                                fmt.Println(v4, v2)
                                var v243 int32 = -2147483648
                                _ = fi25
                                _ = v243
                            }
                        }
                        _ = fi23
                    }
                    v1 = append(v1, ("\n\t" + "hello"))
                }
            }
        }
    case "\n\t":
        v3 += 235
    default:
        var v244 int = 906
        _ = v244
    }
    var v245 int32 = -292
    switch "a" {
    case "hello":
        switch "hello" {
        case "a":
            v2 += (v2 + (^(v2)))
        }
    case "":
        v245 -= 2147483647
    case "a":
        var v246 float32 = 1e10
        _ = v246
    }
    for fi247 := 0; fi247 < 19; fi247++ {
        var v248 float32 = 0.1
        _ = fi247
        _ = v248
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    switch 9223372036854775807 {
    case 4294967295:
        _ = len(v1)
    case 128:
        fmt.Println(v5, v5, v3)
    }
    fmt.Println(v2, v4)
    v2 -= v2
    if (!false) {
        v3 -= uint16(v2)
        if (-9223372036854775808 <= int64(v3)) {
            v249 := GenS6{F0: "hello"}
            var v250 int16 = 1
            _ = v249
            _ = v250
        }
    } else {
        v251 := GenS23{F0: 121, F1: (v4 / ((v5) | 1))}
        v252 := GenS21{F0: 4294967295, F1: 32767, F2: "\n\t"}
        _ = v251
        _ = v252
    }
    v2 = v2
    if true {
        if true {
            v253 := GenS33{F0: false, F1: v245}
            for fi254 := 0; fi254 < 14; fi254++ {
                v2 += (^(v2))
                _ = fi254
            }
            _ = v253
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "hello" {
        case "a":
            fmt.Println(v4, v2)
        case "hello":
            var v255 string = (("a" + "a") + "")
            _ = v255
        default:
            for fi256 := 0; fi256 < 18; fi256++ {
                if false {
                    fmt.Println(v4, v2)
                    v4 = v4
                    var v257 uint8 = 64
                    _ = v257
                } else {
                    var v258 int16 = 0
                    _ = v258
                }
                var v259 float32 = 3.14159
                _ = fi256
                _ = v259
            }
        }
    }
    fmt.Println(v3, v245)
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v245:", v245)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v245
    _ = v3
    _ = v4
    _ = v5
}
