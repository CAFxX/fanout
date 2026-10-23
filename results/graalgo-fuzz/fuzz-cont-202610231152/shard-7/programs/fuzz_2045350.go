package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 int; F2 int32 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 int32; F2 uint8 }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 int64; F1 uint64 }
type GenS14 struct { F0 int16; F1 string; F2 string }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int16, p2 bool) uint64 {
    if p2 {
        p0 = 341
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint64 = 1
            _ = fi7
            _ = v8
        }
        v9 := []int{-9223372036854775808, 1, 2147483647, 32768}
        _ = v9
    } else {
        var v10 uint32 = 2
        if p2 {
            p1 -= (^(p1))
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 float64 = 1e10
                _ = fi11
                _ = v12
            }
        }
        if (32767 > (p0 >> ((1) & 63))) {
            p2 = ((-148 == 128) || true)
            v10 += (v10 | (v10 + 2147483647))
        }
        _ = v10
    }
    p0 += p0
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := []int{}
        for fi15 := 0; fi15 < 6; fi15++ {
            if p2 {
                p0 += (p0 ^ (p0 >> ((140) & 63)))
                v14 = append(v14, fi13)
            }
            p1 = (p1 << ((907) & 63))
            fmt.Println(fi13)
            _ = fi15
        }
        if p2 {
            switch 421 {
            case -2147483648:
                switch ("hello" + (("" + "foo bar") + "hello")) {
                case "":
                    p1 -= p1
                case "\n\t":
                    v16 := GenS1{F0: 487, F1: 42, F2: "foo bar"}
                    _ = v16
                default:
                    p1 = -71
                }
            }
            p1 = (p1 + p1)
        }
        _ = fi13
        _ = v14
    }
    p2 = (255 > 255)
    _ = p0
    _ = p1
    _ = p2
    return 77
}

func main() {
    var v1 float32 = 0.1
    var v2 int32 = -1
    var v3 int = -987
    var v4 int32 = v2
    var v5 uint32 = 65535
    var v6 uint16 = 195
    v2 += 1
    switch (^(v3)) {
    case -58:
        switch v3 {
        case -2147483648:
            switch 128 {
            case -1:
                if true {
                    v4 += 0
                } else {
                    v17 := GenS2{F0: 42, F1: 127}
                    v18 := []int{-2147483648}
                    var v19 uint = 0
                    _ = v17
                    _ = v18
                    _ = v19
                }
            case -9223372036854775808:
                var v20 uint8 = 35
                _ = v20
            case 476:
                v21 := []int{(v3 + (v3 / ((-9223372036854775808) | 1))), 128, v3, v3}
                _ = v21
            default:
                var v22 int16 = 32767
                _ = v22
            }
        case 2147483647:
            if ((true || false) && (1 > -127)) {
                v5 += (v5 - v5)
                v1 = 0.1
                var v23 uint = 885
                _ = v23
            }
        case 42:
            var v24 uint8 = 0
            _ = v24
        }
    case 32768:
        fmt.Println(v1, v3)
    case 313:
        var v25 int32 = v2
        _ = v25
    default:
        fmt.Println(v5, v1, v3)
    }
    v26 := GenS2{F0: 42, F1: 42}
    if false {
        var v27 int8 = int8(v3)
        v3 += 0
        v28 := GenS3{F0: 0}
        _ = v27
        _ = v28
    }
    switch "hello" {
    case "hello":
        if true {
            v29 := GenS4{F0: 42, F1: v3, F2: -1}
            for fi30 := 0; fi30 < 12; fi30++ {
                for fi31 := 0; fi31 < 14; fi31++ {
                    if true {
                        if (!true) {
                            v32 := []int{fi30, (fi31 | int(v5)), (fi31 - (^(v3))), fi30}
                            v6 = (v6 % (((v6 & v6)) | 1))
                            v32 = append(v32, (^(fi30)))
                            _ = v32
                        } else {
                            v6 = v6
                            switch -934 {
                            case 127:
                                var v33 float32 = (v1 - ((1e10 + v1) + v1))
                                _ = v33
                            default:
                                if true {
                                    switch "" {
                                    case "hello":
                                        switch (fi30 << ((4294967295) & 63)) {
                                        case -931:
                                            fmt.Println(v5, fi30)
                                        default:
                                            v34 := []int{fi30, (^(v3))}
                                            _ = v34
                                        }
                                    default:
                                        fmt.Println(fi30)
                                    }
                                } else {
                                    switch (^(v3)) {
                                    case 32768:
                                        var v35 float64 = 0.0
                                        _ = v35
                                    case -1:
                                        var v36 int8 = -102
                                        _ = v36
                                    }
                                }
                            }
                            fmt.Println(v6, v4, fi31)
                        }
                        v4 -= v4
                    }
                    if (4294967295 == 65535) {
                        if ((9223372036854775807 >= 0) && false) {
                            switch ("a" + "foo bar") {
                            case "":
                                var v37 uint8 = 60
                                _ = v37
                            case "a":
                                fmt.Println(v1, v3, v3)
                            default:
                                v38 := GenS5{F0: genFn0((v6 & (v6 / (((v6 - v6)) | 1))), -384, (-9223372036854775808 == (fi31 << ((255) & 63)))), F1: true, F2: 127}
                                _ = v38
                            }
                            if (("" + "foo bar") >= "hello") {
                                v39 := []int{}
                                fmt.Println(v5, fi30, v4)
                                v40 := GenS6{F0: ((2 >= 128) && false), F1: (v6 >> ((4294967295) & 63))}
                                _ = v39
                                _ = v40
                            } else {
                                v5 += (v5 % (((v5 / (((^(v5))) | 1))) | 1))
                            }
                            v5 -= v5
                            v41 := GenS7{F0: (v6 ^ v6)}
                            switch v3 {
                            case -440:
                                v42 := GenS8{F0: (true || false), F1: v6}
                                _ = v42
                            case 0:
                                v43 := GenS6{F0: true, F1: (v6 / ((32767) | 1))}
                                _ = v43
                            }
                            _ = v41
                        } else {
                            v44 := GenS6{F0: ((42 != v6) || false), F1: 42}
                            _ = v44
                        }
                        v3 = (^(v3))
                        v3 = (^(v3))
                    }
                    if true {
                        fmt.Println(v3)
                        v45 := GenS9{F0: 66}
                        v5 += (v5 & 898)
                        _ = v45
                    } else {
                        switch "" {
                        case "foo bar":
                            fmt.Println(v1, v3)
                        case "\n\t":
                            if false {
                                v4 = v4
                            }
                        case "":
                            v3 = fi30
                            var v46 bool = (4294967295 >= 0)
                            var v47 int8 = 42
                            _ = v46
                            _ = v47
                        default:
                            v2 = (v4 * 42)
                        }
                        if false {
                            if (false && true) {
                                v1 = 0.0
                                switch "hello" {
                                case "":
                                    v3 = (v3 - fi30)
                                }
                            } else {
                                fmt.Println(v4, v5)
                            }
                            v5 = (v5 - (v5 - 154))
                        }
                        v48 := []int{(v3 << ((0) & 63)), 1, 4294967295}
                        _ = v48
                    }
                    _ = fi31
                }
                fmt.Println(v1)
                v4 += v4
                _ = fi30
            }
            var v49 uint32 = (v5 >> ((1) & 63))
            _ = v29
            _ = v49
        } else {
            var v50 int64 = -1
            v51 := GenS10{F0: (^(v6))}
            var v52 uint64 = 42
            _ = v50
            _ = v51
            _ = v52
        }
    case "foo bar":
        var v53 uint64 = 0
        _ = v53
    case "":
        switch -294 {
        case 4294967295:
            for fi54 := 0; fi54 < 15; fi54++ {
                for fi55 := 0; fi55 < 17; fi55++ {
                    var v56 int32 = (v4 << ((148) & 63))
                    switch (("a" + "\n\t") + "\n\t") {
                    case "":
                        if ((!true) && (((!true) && true) && (0.0 <= v1))) {
                            var v57 uint32 = (v5 - uint32(v6))
                            switch (fi54 & 291) {
                            case -1:
                                var v58 uint64 = 4294967295
                                _ = v58
                            case 32768:
                                if (0 >= -9223372036854775808) {
                                    v59 := []int{(-(v3))}
                                    var v60 int = (v3 - -1)
                                    v59 = append(v59, 1)
                                    _ = v59
                                    _ = v60
                                }
                            case 43:
                                v56 -= (^(v4))
                            default:
                                v61 := GenS11{F0: (v56 % ((v2) | 1)), F1: -2147483648, F2: 84}
                                _ = v61
                            }
                            _ = v57
                        } else {
                            var v62 int = v3
                            v63 := GenS12{F0: v6, F1: fi55}
                            _ = v62
                            _ = v63
                        }
                    case "a":
                        v4 += (v4 * int32(v5))
                    }
                    _ = fi55
                    _ = v56
                }
                _ = fi54
            }
        case 2147483647:
            fmt.Println(v1, v4, v5)
        }
    }
    var v64 uint8 = 1
    v64 = 1
    var v65 string = ""
    v66 := GenS12{F0: (v6 + v6), F1: v3}
    for fi67 := 0; fi67 < 10; fi67++ {
        switch fi67 {
        case 42:
            v68 := GenS5{F0: 951, F1: true, F2: v64}
            _ = v68
        default:
            if true {
                for fi69 := 0; fi69 < 2; fi69++ {
                    if true {
                        if (1000 == 0) {
                            if true {
                                v70 := GenS13{F0: -9223372036854775808, F1: 18446744073709551615}
                                v71 := GenS8{F0: (!true), F1: (v6 * uint16(v5))}
                                var v72 string = v65
                                var v73 int8 = -128
                                _ = v70
                                _ = v71
                                _ = v72
                                _ = v73
                            } else {
                                switch (v3 & fi69) {
                                case 127:
                                    var v74 bool = (v65 == v65)
                                    _ = v74
                                case 9223372036854775807:
                                    var v75 uint8 = 1
                                    _ = v75
                                case 32767:
                                    var v76 float32 = float32(1.5)
                                    _ = v76
                                default:
                                    fmt.Println(v5)
                                }
                                var v77 bool = true
                                v5 = (v5 ^ v5)
                                _ = v77
                            }
                        } else {
                            v78 := GenS14{F0: 1000, F1: ((v65 + (v65 + v65)) + (("a" + v65) + v65)), F2: (v65 + v65)}
                            v1 = (v1 * v1)
                            var v79 int = (-(fi69))
                            _ = v78
                            _ = v79
                        }
                    }
                    if (!(int32(v64) >= 0)) {
                        if ((v6 <= v6) || false) {
                            fmt.Println(fi67)
                            var v80 bool = false
                            var v81 int8 = 34
                            _ = v80
                            _ = v81
                        } else {
                            var v82 int64 = -327
                            if (255 != v64) {
                                v65 = "a"
                            }
                            var v83 uint32 = 4294967295
                            _ = v82
                            _ = v83
                        }
                    } else {
                        v84 := GenS2{F0: uint64(v6), F1: -128}
                        _ = v84
                    }
                    _ = fi69
                }
            } else {
                var v85 float64 = 1e10
                var v86 float64 = v85
                var v87 float64 = (0.0 + v85)
                _ = v85
                _ = v86
                _ = v87
            }
        }
        fmt.Println(v1)
        switch (v65 + "a") {
        case "\n\t":
            v5 = (v5 ^ v5)
        case "a":
            if ((70 == -84) && (207 != 122)) {
                switch fi67 {
                case -635:
                    for fi88 := 0; fi88 < 19; fi88++ {
                        var v89 int64 = -1
                        v90 := GenS15{F0: -495, F1: 1}
                        var v91 uint = 255
                        _ = fi88
                        _ = v89
                        _ = v90
                        _ = v91
                    }
                case 32767:
                    if false {
                        for fi92 := 0; fi92 < 12; fi92++ {
                            switch v65 {
                            case "foo bar":
                                v1 = 1.5
                            case "a":
                                var v93 int = (^(fi67))
                                _ = v93
                            case "hello":
                                v94 := GenS3{F0: -9223372036854775808}
                                _ = v94
                            default:
                                fmt.Println(fi92, fi92)
                            }
                            var v95 int8 = 7
                            _ = fi92
                            _ = v95
                        }
                        var v96 int16 = 32767
                        for fi97 := 0; fi97 < 7; fi97++ {
                            switch (v3 << ((2) & 63)) {
                            case 127:
                                var v98 int = (v3 << ((42) & 63))
                                _ = v98
                            default:
                                fmt.Println(v5, v1)
                            }
                            var v99 float32 = (1.5 - float32((0.0 * 0.1)))
                            var v100 int64 = 42
                            _ = fi97
                            _ = v100
                            _ = v99
                        }
                        _ = v96
                    }
                case 127:
                    v101 := []int{int(v64)}
                    _ = v101
                }
            }
        default:
            v102 := []int{int(v5), -461, v3}
            _ = v102
        }
        _ = fi67
    }
    var v103 uint = 444
    var v104 int64 = 247
    fmt.Println(v1, v103)
    switch (v3 + v3) {
    case 32767:
        switch v65 {
        case "foo bar":
            for fi105 := 0; fi105 < 8; fi105++ {
                v106 := GenS16{F0: 1}
                v64 -= uint8(v5)
                v64 = (^(v64))
                _ = fi105
                _ = v106
            }
        }
    }
    switch "hello" {
    case "foo bar":
        v107 := []int{v3, v3}
        _ = v107
    }
    var v108 float32 = -2.25
    v65 = v65
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v103
    _ = v104
    _ = v108
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
