package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 int64; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 uint64; F1 int8; F2 int16 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 uint64; F1 uint16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int64; F1 bool }
type GenS14 struct { F0 string; F1 int32 }
type GenS15 struct { F0 int16; F1 string }
type GenS16 struct { F0 int64; F1 uint }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS19 struct { F0 uint16; F1 uint32 }
type GenS20 struct { F0 int32; F1 int16; F2 int64 }
type GenS21 struct { F0 uint8; F1 int8 }

func main() {
    var v1 float32 = -2.25
    v2 := []string{}
    v3 := make(map[string]uint32)
    v4 := []string{}
    v5 := make(map[int]string)
    v6 := []int{-2147483648, 32767, 665, 905}
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    default:
        v7 := GenS1{F0: 96}
        _ = v7
    }
    switch "hello" {
    case "hello":
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
    case "a":
        v1 -= v1
    case "":
        switch -1 {
        case 247:
            var v10 bool = (false || false)
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        if (2 <= 137) {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 uint32 = 3
                if (-2.25 < v1) {
                    if (v13 < (v13 * (v13 - 858))) {
                        var v14 float64 = 1e10
                        if true {
                            if true {
                                v1 = v1
                                v14 = 3.14159
                            }
                            v13 -= (v13 * v13)
                            var v15 int = fi12
                            _ = v15
                        } else {
                            v16 := GenS2{F0: 255, F1: 222, F2: 2147483647}
                            var v17 string = "hello"
                            var v18 float32 = float32((v14 - v14))
                            v13 = (v13 >> ((10) & 63))
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        v13 = v13
                        v19 := GenS1{F0: 784}
                        v13 = 409
                        var v20 float64 = float64(v1)
                        fmt.Println(v20, v14)
                        v20 += float64(3.14159)
                        _ = v14
                        _ = v19
                        _ = v20
                    }
                    v1 = v1
                    fmt.Println(fi11, fi12, fi11)
                } else {
                    v13 -= 737
                    v13 = v13
                    var v21 int16 = 106
                    _ = v21
                }
                v1 -= ((0.1 * 3.14159) / v1)
                _ = fi12
                _ = v13
            }
            if ((-2147483648 > 2147483647) || true) {
                if false {
                    v1 = v1
                    if (-1 >= (fi11 % (((^(fi11))) | 1))) {
                        v1 += (3.14159 * v1)
                    } else {
                        delete(v3, "hello")
                    }
                }
                if (!(fi11 >= fi11)) {
                    v1 = ((-2.25 * float32(float64(-2.25))) + (v1 + ((1.5 + (float32(0.0) * ((v1 - v1) / v1))) / 1.0)))
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 string = (("\n\t" + "foo bar") + "foo bar")
                        if true {
                            switch fi22 {
                            case 32767:
                                var v24 uint64 = 4294967295
                                _ = v24
                            case 128:
                                if (3.14159 != 1e10) {
                                    if len(v6) > 0 {
                                        v6[0] = (fi11 >> ((2) & 63))
                                    }
                                } else {
                                    v25 := v3["foo bar"]
                                    _ = v25
                                }
                            default:
                                v23 = v23
                            }
                            var v26 int = (fi22 / ((4294967295) | 1))
                            v23 = "hello"
                            _ = v26
                        } else {
                            v4 = append(v4, "")
                            v1 += (v1 + float32(float64(float32(1e10))))
                            switch 968 {
                            case 9223372036854775807:
                                switch 128 {
                                case 32767:
                                    switch (fi22 | 32767) {
                                    case 2147483647:
                                        if true {
                                            switch ((v23 + v23) + v23) {
                                            case "foo bar":
                                                fmt.Println(fi11)
                                            case "a":
                                                var v27 int16 = 1
                                                _ = v27
                                            default:
                                                v28 := GenS3{F0: 42, F1: 114, F2: 42}
                                                _ = v28
                                            }
                                        }
                                    case 9223372036854775807:
                                        switch 4294967295 {
                                        case 636:
                                            fmt.Println(fi11, v1)
                                        case -1:
                                            v29 := GenS4{F0: 65535}
                                            _ = v29
                                        case 128:
                                            v30 := GenS2{F0: 1, F1: -761, F2: 4294967295}
                                            _ = v30
                                        }
                                    case -2147483648:
                                        fmt.Println(v1)
                                    }
                                case -1:
                                    switch fi11 {
                                    case 444:
                                        if len(v2) > 0 {
                                            v2[0] = ("a" + "\n\t")
                                        }
                                    case 9223372036854775807:
                                        if false {
                                            var v31 int64 = 170
                                            var v32 int = 127
                                            _ = v31
                                            _ = v32
                                        }
                                        fmt.Println(v1, fi11, fi11)
                                    case -499:
                                        var v33 int32 = 1
                                        _ = v33
                                    default:
                                        fmt.Println(fi22)
                                    }
                                case 42:
                                    var v34 int16 = -1
                                    if (4294967295 == 243) {
                                        v35 := GenS3{F0: 812, F1: 35, F2: 2}
                                        v36 := GenS2{F0: 851, F1: 2147483647, F2: 1}
                                        if (v23 <= v23) {
                                            switch (v23 + "foo bar") {
                                            case "":
                                                if true {
                                                    v37 := GenS5{F0: -1, F1: v34}
                                                    fmt.Println(fi22, fi22, v1)
                                                    _ = v37
                                                } else {
                                                    if true {
                                                        var v38 uint8 = 2
                                                        v38 -= v38
                                                        fmt.Println(v23)
                                                        _ = v38
                                                    }
                                                }
                                            case "a":
                                                switch "a" {
                                                case "hello":
                                                    if false {
                                                        switch "foo bar" {
                                                        case "a":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        case "foo bar":
                                                            if false {
                                                                if true {
                                                                    var v39 string = v23
                                                                    v34 -= v34
                                                                    var v40 int16 = (v34 << ((2) & 63))
                                                                    v41 := GenS6{F0: 113, F1: ("foo bar" + "foo bar")}
                                                                    if false {
                                                                        switch v39 {
                                                                        case "a":
                                                                            if true {
                                                                                var v42 uint32 = 443
                                                                                if (1 == 922) {
                                                                                    fmt.Println(v23, v42, v1)
                                                                                }
                                                                                _ = v42
                                                                            }
                                                                        case "":
                                                                            if (true && true) {
                                                                                switch v23 {
                                                                                case "":
                                                                                    var v43 string = "foo bar"
                                                                                    var v44 string = "a"
                                                                                    v39 = "hello"
                                                                                    if (false || (128 == fi11)) {
                                                                                        var v45 int = (^(fi22))
                                                                                        v46 := GenS7{F0: 1, F1: 33, F2: (v34 & (v40 * v34))}
                                                                                        var v47 int16 = v40
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    } else {
                                                                                        v43 = v44
                                                                                        switch fi11 {
                                                                                        case 442:
                                                                                            if true {
                                                                                                var v48 int32 = 42
                                                                                                v49 := GenS4{F0: 1}
                                                                                                v50 := GenS8{F0: fi11, F1: 2}
                                                                                                _ = v48
                                                                                                _ = v49
                                                                                                _ = v50
                                                                                            }
                                                                                        case 127:
                                                                                            v23 = (v39 + (v43 + ((v43 + "a") + (v44 + "a"))))
                                                                                        case 0:
                                                                                            v34 += v34
                                                                                        default:
                                                                                            var v51 string = (v44 + ("\n\t" + v23))
                                                                                            _ = v51
                                                                                        }
                                                                                    }
                                                                                    _ = v43
                                                                                    _ = v44
                                                                                }
                                                                                v52 := GenS9{F0: -124}
                                                                                if (v39 <= v39) {
                                                                                    v23 = (v23 + ("\n\t" + (v23 + "hello")))
                                                                                    switch -9223372036854775808 {
                                                                                    case 32768:
                                                                                        v53 := GenS5{F0: 9223372036854775807, F1: v40}
                                                                                        _ = v53
                                                                                    case 1:
                                                                                        v54 := GenS10{F0: 127, F1: 115}
                                                                                        _ = v54
                                                                                    default:
                                                                                        switch (-(fi11)) {
                                                                                        case -2147483648:
                                                                                            if (0 > 191) {
                                                                                                v40 -= v40
                                                                                            } else {
                                                                                                _ = len(v5)
                                                                                                v55 := GenS11{F0: 4294967295, F1: 229}
                                                                                                _ = v55
                                                                                            }
                                                                                        case 2147483647:
                                                                                            v34 = (v34 << ((42) & 63))
                                                                                            switch (v39 + (v39 + (v23 + v23))) {
                                                                                            case "foo bar":
                                                                                                var v56 int8 = 98
                                                                                                _ = v56
                                                                                            case "":
                                                                                                v57 := GenS4{F0: 318}
                                                                                                _ = v57
                                                                                            default:
                                                                                                v4 = append(v4, v23)
                                                                                            }
                                                                                        case 4294967295:
                                                                                            v23 = v23
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v40, v40)
                                                                                }
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                        if false {
                                                                            var v58 uint64 = 552
                                                                            fmt.Println(v39, fi22, v58)
                                                                            fmt.Println(v39, fi22, v23)
                                                                            _ = v58
                                                                        } else {
                                                                            if true {
                                                                                v59 := GenS12{F0: -598}
                                                                                v60 := GenS7{F0: 42, F1: -25, F2: (v34 - v34)}
                                                                                fmt.Println(fi11, v23, v40)
                                                                                switch (fi22 - fi22) {
                                                                                case -191:
                                                                                    var v61 int = (fi22 - fi11)
                                                                                    _ = v61
                                                                                default:
                                                                                    _ = len(v5)
                                                                                }
                                                                                _ = v59
                                                                                _ = v60
                                                                            } else {
                                                                                if (true && true) {
                                                                                    var v62 float32 = 3.14159
                                                                                    var v63 int16 = (v40 | 390)
                                                                                    _ = len(v3)
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                            if (false || (!false)) {
                                                                                var v64 int32 = 1
                                                                                _ = v64
                                                                            }
                                                                            switch fi11 {
                                                                            case 127:
                                                                                if ("" == "") {
                                                                                    if ((!true) && (-641 <= -1)) {
                                                                                        switch 9223372036854775807 {
                                                                                        case 2147483647:
                                                                                            _ = len(v6)
                                                                                            var v65 int = (fi11 % ((fi22) | 1))
                                                                                            _ = v65
                                                                                        case 9223372036854775807:
                                                                                            if (!(-1 > -9223372036854775808)) {
                                                                                                v40 += (v34 / (((v34 - v34)) | 1))
                                                                                            } else {
                                                                                                fmt.Println(v1)
                                                                                                v66 := GenS13{F0: 688, F1: (((v1 < (v1 - v1)) && (!(!true))) && false)}
                                                                                                _ = v66
                                                                                            }
                                                                                        case 128:
                                                                                            v40 = v40
                                                                                        default:
                                                                                            switch (fi11 ^ fi11) {
                                                                                            case -2147483648:
                                                                                                delete(v3, "hello")
                                                                                            case 490:
                                                                                                v23 = (("a" + (v23 + "\n\t")) + v39)
                                                                                            case 412:
                                                                                                v34 = v34
                                                                                            }
                                                                                        }
                                                                                        var v67 uint8 = 1
                                                                                        v34 = 0
                                                                                        switch 128 {
                                                                                        case 32767:
                                                                                            _ = len(v5)
                                                                                        case -9223372036854775808:
                                                                                            var v68 uint64 = 18446744073709551615
                                                                                            _ = v68
                                                                                        }
                                                                                        var v69 uint32 = 4294967295
                                                                                        _ = v67
                                                                                        _ = v69
                                                                                    }
                                                                                }
                                                                            case -9223372036854775808:
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    v70 := GenS14{F0: ((v23 + (v39 + v39)) + (v23 + ("hello" + "hello"))), F1: 1}
                                                                                    _ = v70
                                                                                }
                                                                            default:
                                                                                v71 := GenS7{F0: 42, F1: 100, F2: (v34 / ((v34) | 1))}
                                                                                v34 -= -32768
                                                                                v40 = v40
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                        var v72 uint16 = 0
                                                                        if false {
                                                                            if (0 <= 88) {
                                                                                if false {
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = (fi22 | fi11)
                                                                                    }
                                                                                    delete(v3, v39)
                                                                                    switch 127 {
                                                                                    case 895:
                                                                                        var v73 uint = 855
                                                                                        _ = v73
                                                                                    case 32767:
                                                                                        fmt.Println(v72)
                                                                                    case 127:
                                                                                        v39 = "a"
                                                                                    default:
                                                                                        if (true || (!((!true) || (int64(v34) > 47)))) {
                                                                                            if true {
                                                                                                v5[(fi11 - (fi11 + fi22))] = "a"
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            if (-128 < -1) {
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = v39
                                                                                                }
                                                                                                var v74 uint64 = uint64(v72)
                                                                                                _ = v74
                                                                                            } else {
                                                                                                var v75 float64 = -2.25
                                                                                                _ = v75
                                                                                            }
                                                                                            v39 = "hello"
                                                                                            v39 = "foo bar"
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v76 := GenS12{F0: -1}
                                                                                    v39 = v39
                                                                                    _ = v76
                                                                                }
                                                                                v39 = "a"
                                                                                v34 += v40
                                                                            }
                                                                            if true {
                                                                                var v77 float64 = 0.1
                                                                                if ((255 > 128) || (true || true)) {
                                                                                    v34 -= (v40 % (((v34 * (-(v34)))) | 1))
                                                                                    v77 = v77
                                                                                } else {
                                                                                    if false {
                                                                                        fmt.Println(v23, v40)
                                                                                        var v78 int = fi11
                                                                                        v72 = (v72 - 1)
                                                                                        _ = v78
                                                                                    } else {
                                                                                        v34 += (v34 + v40)
                                                                                    }
                                                                                    var v79 float64 = v77
                                                                                    _ = v79
                                                                                }
                                                                                switch v23 {
                                                                                case "a":
                                                                                    if false {
                                                                                        v72 -= (v72 & v72)
                                                                                    } else {
                                                                                        v80 := GenS15{F0: v40, F1: (v23 + ("a" + v39))}
                                                                                        _ = v80
                                                                                    }
                                                                                case "\n\t":
                                                                                    v81 := GenS12{F0: -2147483648}
                                                                                    _ = v81
                                                                                case "foo bar":
                                                                                    v82 := GenS9{F0: -128}
                                                                                    _ = v82
                                                                                }
                                                                                _ = v77
                                                                            }
                                                                            var v83 uint16 = (^(v72))
                                                                            _ = v83
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                    _ = v39
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                                v84 := GenS8{F0: fi22, F1: 0}
                                                                _ = v84
                                                            }
                                                        case "hello":
                                                            v23 = v23
                                                        default:
                                                            var v85 float32 = v1
                                                            _ = v85
                                                        }
                                                        fmt.Println(fi22, fi22)
                                                        v86 := GenS6{F0: -90, F1: (((v23 + "foo bar") + v23) + v23)}
                                                        _ = v86
                                                    } else {
                                                        v3[v23] = 699
                                                    }
                                                case "a":
                                                    var v87 int64 = 42
                                                    _ = v87
                                                case "foo bar":
                                                    switch v23 {
                                                    case "a":
                                                        var v88 int64 = 42
                                                        _ = v88
                                                    case "\n\t":
                                                        v23 = v23
                                                        _ = len(v2)
                                                        if true {
                                                            var v89 uint16 = 65535
                                                            v90 := GenS12{F0: 1}
                                                            if true {
                                                                v34 -= v34
                                                                if (18446744073709551615 > 18446744073709551615) {
                                                                    if false {
                                                                        delete(v5, (fi11 | -900))
                                                                        v4 = append(v4, ("foo bar" + v23))
                                                                    } else {
                                                                        v91 := GenS3{F0: 42, F1: 42, F2: v89}
                                                                        _ = v91
                                                                    }
                                                                    var v92 string = "foo bar"
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = v92
                                                                } else {
                                                                    if (v89 < v89) {
                                                                        var v93 int32 = 42
                                                                        if (true || false) {
                                                                            fmt.Println(v1)
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            var v94 uint32 = 4294967295
                                                                            _ = v94
                                                                        }
                                                                        var v95 int8 = 97
                                                                        _ = v93
                                                                        _ = v95
                                                                    } else {
                                                                        v89 += (^(v89))
                                                                        _ = len(v2)
                                                                        var v96 int8 = -128
                                                                        _ = v96
                                                                    }
                                                                    v34 -= v34
                                                                }
                                                            } else {
                                                                fmt.Println(v34, fi22)
                                                                if (43 <= 44) {
                                                                    var v97 int = 273
                                                                    switch -2147483648 {
                                                                    case -135:
                                                                        var v98 bool = (2 <= 293)
                                                                        _ = v98
                                                                    default:
                                                                        var v99 float32 = float32(0.1)
                                                                        _ = v99
                                                                    }
                                                                    _ = v97
                                                                } else {
                                                                    v100 := v3[v23]
                                                                    var v101 int8 = 1
                                                                    v34 -= int16(v101)
                                                                    _ = v100
                                                                    _ = v101
                                                                }
                                                                switch -486 {
                                                                case 32767:
                                                                    var v102 int16 = 42
                                                                    _ = v102
                                                                }
                                                            }
                                                            v103 := GenS4{F0: v89}
                                                            v104 := GenS16{F0: 1, F1: 65535}
                                                            _ = v103
                                                            _ = v104
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                switch ((v23 + (v23 + "hello")) + "hello") {
                                                case "":
                                                    v1 += v1
                                                case "hello":
                                                    v34 += (v34 * v34)
                                                case "foo bar":
                                                    if false {
                                                        v1 = v1
                                                        if false {
                                                            switch (v23 + "hello") {
                                                            case "foo bar":
                                                                switch (v23 + v23) {
                                                                case "a":
                                                                    v105 := GenS15{F0: 0, F1: v23}
                                                                    _ = v105
                                                                case "hello":
                                                                    v3[("hello" + (v23 + (v23 + "")))] = 138
                                                                default:
                                                                    v106 := GenS17{F0: 108}
                                                                    _ = v106
                                                                }
                                                            case "":
                                                                switch fi22 {
                                                                case 0:
                                                                    var v107 int8 = 42
                                                                    _ = v107
                                                                default:
                                                                    fmt.Println(fi22, fi11, v1)
                                                                }
                                                            case "\n\t":
                                                                v23 = ((v23 + v23) + (v23 + "foo bar"))
                                                            }
                                                        } else {
                                                            v108 := GenS16{F0: -9223372036854775808, F1: 18446744073709551615}
                                                            _ = v108
                                                        }
                                                        var v109 uint8 = 200
                                                        _ = v109
                                                    } else {
                                                        var v110 uint16 = 1
                                                        v23 = "foo bar"
                                                        v111 := GenS18{F0: 1, F1: 206, F2: 128}
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                    fmt.Println(fi11)
                                                }
                                            default:
                                                v4 = append(v4, ("\n\t" + v23))
                                            }
                                            var v112 uint16 = 1
                                            fmt.Println(v23, fi22, v1)
                                            _ = v112
                                        }
                                        _ = v35
                                        _ = v36
                                    } else {
                                        switch v23 {
                                        case "":
                                            v23 = v23
                                        default:
                                            v1 -= (v1 * (v1 + v1))
                                        }
                                    }
                                    _ = v34
                                default:
                                    fmt.Println(fi22, v23, fi11)
                                    v113 := GenS19{F0: 42, F1: 2147483647}
                                    if len(v4) > 0 {
                                        v4[0] = v23
                                    }
                                    var v114 uint64 = 1
                                    var v115 uint16 = 2
                                    _ = v113
                                    _ = v114
                                    _ = v115
                                }
                            case -2147483648:
                                v116 := GenS20{F0: 42, F1: 612, F2: 42}
                                _ = v116
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = len(v6)
                }
                v117 := GenS11{F0: 1, F1: 32767}
                _ = v117
            }
            v1 -= (v1 * 1e10)
        }
        _ = fi11
    }
    if (342 != 867) {
        if (float32(3.14159) > 3.14159) {
            var v118 bool = false
            v1 -= (0.1 - float32(3.14159))
            _ = v118
        } else {
            for fi119 := 0; fi119 < 2; fi119++ {
                v1 += (((0.0 + v1) * v1) + v1)
                _ = fi119
            }
            v120 := GenS3{F0: -910, F1: 42, F2: 25}
            for fi121 := 0; fi121 < 11; fi121++ {
                delete(v5, (fi121 - 42))
                fmt.Println(fi121, v1)
                _ = fi121
            }
            _ = v120
        }
    } else {
        v122 := v5[128]
        if (!(1e10 == 1e10)) {
            v123 := GenS21{F0: 128, F1: 117}
            _ = v123
        } else {
            v124 := GenS15{F0: -773, F1: v122}
            fmt.Println(v122, v1)
            v122 = v122
            _ = v124
        }
        _ = v122
    }
    fmt.Println(v1)
    v1 -= v1
    var v125 int64 = 9223372036854775807
    v125 -= (v125 + (v125 ^ v125))
    _ = len(v5)
    var v126 string = "\n\t"
    v127 := GenS3{F0: v125, F1: 1, F2: 42}
    delete(v3, "")
    _ = len(v3)
    var v128 uint = 42
    for fi129 := 0; fi129 < 13; fi129++ {
        var v130 string = "\n\t"
        _ = fi129
        _ = v130
    }
    v128 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v128:", v128)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v125
    _ = v126
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
