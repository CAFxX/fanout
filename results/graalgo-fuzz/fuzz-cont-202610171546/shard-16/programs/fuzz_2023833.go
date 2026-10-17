package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 uint32; F1 uint }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 string; F1 string; F2 int }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint32; F1 uint }
type GenS16 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS17 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS18 struct { F0 bool; F1 int64 }
type GenS19 struct { F0 uint; F1 bool }
type GenS20 struct { F0 int16; F1 int }
type GenS21 struct { F0 bool; F1 bool }
type GenS22 struct { F0 string; F1 uint }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 uint }
type GenS25 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 bool; F1 uint32; F2 string }
type GenS28 struct { F0 int32; F1 bool }
type GenS29 struct { F0 uint; F1 uint; F2 int }
type GenS30 struct { F0 int; F1 uint }
type GenS31 struct { F0 bool; F1 uint; F2 int16 }
type GenS32 struct { F0 uint; F1 uint64 }
type GenS33 struct { F0 int32; F1 int16; F2 int }
type GenS34 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS35 struct { F0 bool; F1 uint }
type GenS36 struct { F0 int32; F1 int; F2 int32 }

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 128
    var v3 uint8 = (v2 / (((v2 % ((42) | 1))) | 1))
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]string)
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 2; fi7++ {
        delete(v5, 730)
        var v8 int16 = -1
        _ = fi7
        _ = v8
    }
    var v9 float32 = (float32(((1.5 - 0.0) / 3.14159)) + 0.1)
    v10 := v5[32768]
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            var v12 uint64 = 347
            v6[1] = 23
            var v13 uint32 = 2147483647
            _ = v12
            _ = v13
        } else {
            v10 = (v10 + ((v10 + v10) + v10))
            v3 += (v3 % ((v3) | 1))
        }
        v14 := []int{fi11, fi11, (fi11 * (^(fi11)))}
        v15 := GenS1{F0: 42, F1: "\n\t", F2: -1}
        _ = fi11
        _ = v14
        _ = v15
    }
    if false {
        if (v10 >= ((("foo bar" + "a") + v10) + ("" + (v10 + ((((v10 + (v10 + v10)) + (v10 + v10)) + "") + "foo bar"))))) {
            fmt.Println(v3, v4, v9)
            delete(v6, 2147483647)
            v9 = v9
        } else {
            var v16 string = v10
            switch "hello" {
            case "a":
                v3 = uint8(v1)
            case "":
                for fi17 := 0; fi17 < 19; fi17++ {
                    for fi18 := 0; fi18 < 12; fi18++ {
                        if (!true) {
                            v19 := GenS1{F0: 436, F1: (((("" + v16) + ((v10 + ("a" + ("a" + v16))) + v10)) + (("\n\t" + ((v16 + v16) + (v16 + v16))) + v16)) + v10), F2: -2147483648}
                            v10 = v16
                            _ = v19
                        } else {
                            v16 = "foo bar"
                            switch (fi17 % (((fi18 << ((0) & 63))) | 1)) {
                            case -747:
                                if false {
                                    _ = len(v5)
                                    switch (("foo bar" + (v16 + v10)) + (v10 + (("hello" + v16) + "hello"))) {
                                    case "foo bar":
                                        v6[2] = 0
                                    case "":
                                        v20 := []int{(-(fi17)), -432, fi17, (fi17 % (((fi18 + (fi17 | (^(fi18))))) | 1))}
                                        _ = v20
                                    case "hello":
                                        if (-1 >= 32767) {
                                            if false {
                                                v10 = "\n\t"
                                                v4 = v4
                                            }
                                            switch ("a" + (((v16 + "foo bar") + ((v16 + v10) + (v16 + (v16 + "foo bar")))) + v10)) {
                                            case "\n\t":
                                                if false {
                                                    switch (((("\n\t" + (("hello" + "a") + v10)) + "hello") + v16) + (((v16 + "hello") + (v16 + "")) + v10)) {
                                                    case "":
                                                        if (true || (0 >= 2147483647)) {
                                                            if true {
                                                                var v21 int32 = int32(v2)
                                                                if (!false) {
                                                                    switch (fi17 ^ fi17) {
                                                                    case 0:
                                                                        v22 := GenS2{F0: v4, F1: 32767}
                                                                        if (!true) {
                                                                            var v23 int64 = (v4 >> ((65535) & 63))
                                                                            v9 = 1e10
                                                                            _ = v23
                                                                        }
                                                                        switch v16 {
                                                                        case "a":
                                                                            fmt.Println(fi18, v9)
                                                                            v24 := GenS3{F0: fi18, F1: 42}
                                                                            fmt.Println(v21, v1, fi18)
                                                                            _ = v24
                                                                        case "foo bar":
                                                                            switch int(v1) {
                                                                            case 42:
                                                                                var v25 bool = false
                                                                                _ = v25
                                                                            case -2147483648:
                                                                                if true {
                                                                                    v26 := GenS2{F0: v4, F1: -32768}
                                                                                    _ = v26
                                                                                } else {
                                                                                    var v27 float32 = 0.1
                                                                                    v28 := GenS4{F0: v21, F1: -1}
                                                                                    v1 = (v1 ^ 1)
                                                                                    v10 = "foo bar"
                                                                                    var v29 uint16 = v1
                                                                                    _ = v27
                                                                                    _ = v28
                                                                                    _ = v29
                                                                                }
                                                                            case 524:
                                                                                if (v2 <= v3) {
                                                                                    var v30 float32 = v9
                                                                                    v21 = (v21 - 201)
                                                                                    _ = v30
                                                                                }
                                                                                fmt.Println(v21)
                                                                                if true {
                                                                                    switch (v10 + v16) {
                                                                                    case "":
                                                                                        delete(v5, fi18)
                                                                                    }
                                                                                    v3 = v2
                                                                                }
                                                                            default:
                                                                                var v31 int64 = v4
                                                                                _ = v31
                                                                            }
                                                                            var v32 uint = 42
                                                                            if (2147483648 != 1) {
                                                                                v2 = v3
                                                                                v33 := GenS5{F0: 0, F1: 968}
                                                                                if (!true) {
                                                                                    var v34 float32 = v9
                                                                                    _ = v34
                                                                                } else {
                                                                                    fmt.Println(v3, fi18, v21)
                                                                                    v9 += (-2.25 / 1e10)
                                                                                    fmt.Println(v32)
                                                                                    v35 := GenS6{F0: -48, F1: v16}
                                                                                    _ = v35
                                                                                }
                                                                                _ = v33
                                                                            } else {
                                                                                switch int(v32) {
                                                                                case 0:
                                                                                    v36 := GenS7{F0: (-(v4)), F1: 1, F2: 0}
                                                                                    _ = v36
                                                                                case 32768:
                                                                                    switch -1 {
                                                                                    case 32768:
                                                                                        v37 := GenS1{F0: 0, F1: "hello", F2: int32(fi18)}
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            if true {
                                                                                                v2 -= (v2 % (((v2 ^ (v2 % (((v2 >> (((v32 ^ 0)) & 63))) | 1)))) | 1))
                                                                                                v2 = (v3 & v2)
                                                                                            } else {
                                                                                                var v38 int = (fi18 - fi18)
                                                                                                if (v1 == v1) {
                                                                                                    v3 = v2
                                                                                                    fmt.Println(v38, fi17, fi17)
                                                                                                } else {
                                                                                                    v4 -= v4
                                                                                                }
                                                                                                _ = v38
                                                                                            }
                                                                                        case "":
                                                                                            v4 = 0
                                                                                        case "hello":
                                                                                            fmt.Println(v4, fi17)
                                                                                        default:
                                                                                            v39 := GenS8{F0: 22, F1: 255}
                                                                                            _ = v39
                                                                                        }
                                                                                        v5[-1] = v16
                                                                                        _ = v37
                                                                                    case 1:
                                                                                        var v40 int8 = 112
                                                                                        v41 := GenS1{F0: 917, F1: v16, F2: -1}
                                                                                        v42 := GenS2{F0: v4, F1: int16(fi17)}
                                                                                        _ = v40
                                                                                        _ = v41
                                                                                        _ = v42
                                                                                    case 19:
                                                                                        fmt.Println(fi17, v32, v2)
                                                                                    default:
                                                                                        fmt.Println(fi18)
                                                                                    }
                                                                                    var v43 bool = (!(!true))
                                                                                    _ = v43
                                                                                case 491:
                                                                                    if true {
                                                                                        if (!false) {
                                                                                            switch (-(fi18)) {
                                                                                            case 32767:
                                                                                                v32 -= v32
                                                                                            case 59:
                                                                                                v1 = v1
                                                                                            case 4294967295:
                                                                                                var v44 float64 = ((1e10 / 0.1) * 3.14159)
                                                                                                if false {
                                                                                                    var v45 uint = 65535
                                                                                                    v46 := GenS2{F0: (v4 / (((v4 >> ((0) & 63))) | 1)), F1: 42}
                                                                                                    _ = v45
                                                                                                    _ = v46
                                                                                                }
                                                                                                v9 += v9
                                                                                                v47 := GenS2{F0: v4, F1: -683}
                                                                                                if (!true) {
                                                                                                    v48 := GenS9{F0: 2, F1: -203, F2: (^(v1))}
                                                                                                    fmt.Println(v10, v10)
                                                                                                    if false {
                                                                                                        if false {
                                                                                                            var v49 int = (fi18 / ((int(v3)) | 1))
                                                                                                            v50 := GenS1{F0: 4294967295, F1: v16, F2: 2147483647}
                                                                                                            _ = v49
                                                                                                            _ = v50
                                                                                                        } else {
                                                                                                            switch "a" {
                                                                                                            case "\n\t":
                                                                                                                v51 := []int{(fi17 ^ fi17)}
                                                                                                                var v52 uint8 = 51
                                                                                                                v16 = (((("" + v10) + v16) + v10) + "hello")
                                                                                                                _ = v51
                                                                                                                _ = v52
                                                                                                            default:
                                                                                                                delete(v6, 550)
                                                                                                            }
                                                                                                            var v53 float32 = (v9 - 1.5)
                                                                                                            var v54 int = (^(fi17))
                                                                                                            _ = v53
                                                                                                            _ = v54
                                                                                                        }
                                                                                                        var v55 uint32 = 2147483647
                                                                                                        _ = v55
                                                                                                    } else {
                                                                                                        v56 := v5[(-(fi18))]
                                                                                                        if ((v32 >> (((v32 + v32)) & 63)) >= (v32 - (v32 / ((v32) | 1)))) {
                                                                                                            v10 = (v16 + v10)
                                                                                                        }
                                                                                                        _ = v56
                                                                                                    }
                                                                                                    _ = len(v6)
                                                                                                    _ = v48
                                                                                                }
                                                                                                _ = v44
                                                                                                _ = v47
                                                                                            default:
                                                                                                var v57 uint64 = 0
                                                                                                _ = v57
                                                                                            }
                                                                                            switch fi17 {
                                                                                            case 9223372036854775807:
                                                                                                v58 := []int{9223372036854775807}
                                                                                                _ = v58
                                                                                            case 128:
                                                                                                v59 := GenS10{F0: v1, F1: 1}
                                                                                                _ = v59
                                                                                            case 32767:
                                                                                                v10 = (v10 + (v10 + "hello"))
                                                                                            }
                                                                                            if false {
                                                                                                var v60 float32 = float32(((0.0 / (float64(v9) + -2.25)) - ((0.0 / 1e10) * 0.0)))
                                                                                                if ((v21 & 42) >= 2147483647) {
                                                                                                    var v61 uint64 = 620
                                                                                                    _ = v61
                                                                                                }
                                                                                                var v62 int16 = 40
                                                                                                _ = v60
                                                                                                _ = v62
                                                                                            }
                                                                                            v4 = v4
                                                                                        } else {
                                                                                            fmt.Println(fi18, v1, fi18)
                                                                                            if false {
                                                                                                var v63 int = (fi18 << (((v32 >> ((707) & 63))) & 63))
                                                                                                v64 := GenS11{F0: 728, F1: (v32 ^ (v32 ^ v32))}
                                                                                                _ = v63
                                                                                                _ = v64
                                                                                            } else {
                                                                                                v5[(fi17 * (fi17 | 127))] = ((v10 + ((v16 + "\n\t") + v10)) + v10)
                                                                                                if true {
                                                                                                    fmt.Println(v4, fi18, v3)
                                                                                                    if ((^(v1)) > 826) {
                                                                                                        var v65 uint8 = (v2 ^ v2)
                                                                                                        var v66 int8 = 127
                                                                                                        v67 := []int{}
                                                                                                        _ = v65
                                                                                                        _ = v66
                                                                                                        _ = v67
                                                                                                    } else {
                                                                                                        v68 := GenS4{F0: (v21 << ((18446744073709551615) & 63)), F1: (v21 >> ((255) & 63))}
                                                                                                        switch 1 {
                                                                                                        case 886:
                                                                                                            v69 := GenS12{F0: 98, F1: 18446744073709551615}
                                                                                                            _ = v69
                                                                                                        case -796:
                                                                                                            switch v10 {
                                                                                                            case "":
                                                                                                                fmt.Println(v2, v16)
                                                                                                            case "hello":
                                                                                                                var v70 uint64 = 1
                                                                                                                _ = v70
                                                                                                            case "\n\t":
                                                                                                                var v71 uint32 = 2147483648
                                                                                                                _ = v71
                                                                                                            default:
                                                                                                                switch v10 {
                                                                                                                case "\n\t":
                                                                                                                    v72 := GenS13{F0: (v10 + v16), F1: (("" + ((v10 + "") + v16)) + v10), F2: 127}
                                                                                                                    _ = v72
                                                                                                                case "a":
                                                                                                                    var v73 int32 = (v21 | (v21 & 1))
                                                                                                                    _ = v73
                                                                                                                case "":
                                                                                                                    v74 := GenS13{F0: (v16 + v10), F1: v16, F2: fi17}
                                                                                                                    _ = v74
                                                                                                                default:
                                                                                                                    switch "a" {
                                                                                                                    case "a":
                                                                                                                        var v75 bool = true
                                                                                                                        _ = v75
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        case 127:
                                                                                                            v10 = v16
                                                                                                        default:
                                                                                                            var v76 float64 = 1e10
                                                                                                            _ = v76
                                                                                                        }
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    var v77 string = ("\n\t" + v10)
                                                                                                    if true {
                                                                                                        var v78 int8 = 21
                                                                                                        _ = len(v6)
                                                                                                        switch v16 {
                                                                                                        case "":
                                                                                                            var v79 int64 = 0
                                                                                                            _ = v79
                                                                                                        case "a":
                                                                                                            v21 = int32(v78)
                                                                                                        case "foo bar":
                                                                                                            v16 = "hello"
                                                                                                        default:
                                                                                                            switch v16 {
                                                                                                            case "a":
                                                                                                                fmt.Println(v32, fi17, v32)
                                                                                                            case "":
                                                                                                                var v80 uint32 = 42
                                                                                                                _ = v80
                                                                                                            case "hello":
                                                                                                                var v81 int64 = v4
                                                                                                                _ = v81
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v78
                                                                                                    } else {
                                                                                                        switch fi18 {
                                                                                                        case 9223372036854775807:
                                                                                                            v82 := GenS14{F0: (v1 ^ v1)}
                                                                                                            _ = v82
                                                                                                        }
                                                                                                        fmt.Println(v32, v1)
                                                                                                    }
                                                                                                    v83 := []int{(fi17 & fi18), int(v32), (fi18 >> (((^(v32))) & 63)), fi17}
                                                                                                    _ = v77
                                                                                                    _ = v83
                                                                                                }
                                                                                                _ = len(v5)
                                                                                            }
                                                                                            var v84 int8 = 42
                                                                                            _ = v84
                                                                                        }
                                                                                        delete(v5, fi17)
                                                                                    } else {
                                                                                        v21 -= v21
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v32
                                                                        default:
                                                                            fmt.Println(v21, fi18, v21)
                                                                        }
                                                                        v2 -= (v3 - (^(v2)))
                                                                        var v85 uint32 = 65535
                                                                        _ = v22
                                                                        _ = v85
                                                                    default:
                                                                        switch 42 {
                                                                        case 4294967295:
                                                                            if (!(0 > 227)) {
                                                                                fmt.Println(v21)
                                                                                v1 = v1
                                                                                var v86 int8 = 116
                                                                                if false {
                                                                                    delete(v6, 1)
                                                                                    fmt.Println(v2, v1, v9)
                                                                                } else {
                                                                                    switch ("foo bar" + (("hello" + ((("\n\t" + v10) + "") + (v16 + v16))) + ("hello" + v16))) {
                                                                                    case "a":
                                                                                        var v87 uint16 = v1
                                                                                        _ = v87
                                                                                    default:
                                                                                        var v88 float64 = 0.0
                                                                                        _ = v88
                                                                                    }
                                                                                    if (!((0 != 895) && true)) {
                                                                                        v1 += 2
                                                                                        var v89 int = (fi17 & 42)
                                                                                        _ = v89
                                                                                    } else {
                                                                                        var v90 string = ("\n\t" + v16)
                                                                                        _ = v90
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        var v91 uint8 = v3
                                                                                        _ = v91
                                                                                    case "hello":
                                                                                        if true {
                                                                                            delete(v5, 844)
                                                                                        }
                                                                                    case "a":
                                                                                        if false {
                                                                                            switch v16 {
                                                                                            case "a":
                                                                                                if (v1 >= v1) {
                                                                                                    if ((328 == v1) || true) {
                                                                                                        v4 = v4
                                                                                                        v92 := GenS4{F0: (^(v21)), F1: (-(v21))}
                                                                                                        _ = v92
                                                                                                    } else {
                                                                                                        v93 := GenS15{F0: 70, F1: 0}
                                                                                                        fmt.Println(v2, fi18, v10)
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    fmt.Println(v16, v9)
                                                                                                    v94 := GenS2{F0: (v4 % (((v4 | (v4 - (v4 - 1)))) | 1)), F1: 42}
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    v95 := GenS14{F0: 255}
                                                                                                    v96 := GenS9{F0: 0, F1: -32768, F2: v1}
                                                                                                    if (false && (int16(v21) != 242)) {
                                                                                                        v1 += 255
                                                                                                        fmt.Println(v9)
                                                                                                    } else {
                                                                                                        var v97 uint = 4294967295
                                                                                                        fmt.Println(v3, v10)
                                                                                                        if true {
                                                                                                            v98 := []int{(fi17 - (fi17 << ((255) & 63))), fi18, (fi18 << ((4294967295) & 63)), fi17}
                                                                                                            var v99 uint16 = v1
                                                                                                            var v100 uint16 = 2
                                                                                                            if true {
                                                                                                                if ((true && false) || (-1 == v86)) {
                                                                                                                    v2 = (v3 % ((122) | 1))
                                                                                                                    fmt.Println(fi17, v100, fi18)
                                                                                                                } else {
                                                                                                                    if (3.14159 > 3.14159) {
                                                                                                                        var v101 bool = false
                                                                                                                        var v102 int = (fi18 - -2147483648)
                                                                                                                        switch (v102 >> ((259) & 63)) {
                                                                                                                        case 1:
                                                                                                                            var v103 uint = v97
                                                                                                                            _ = v103
                                                                                                                        case 128:
                                                                                                                            v3 = (v3 >> (((v97 ^ 2)) & 63))
                                                                                                                        case 42:
                                                                                                                            fmt.Println(v2)
                                                                                                                        default:
                                                                                                                            v5[v102] = (("a" + "foo bar") + "foo bar")
                                                                                                                        }
                                                                                                                        _ = v101
                                                                                                                        _ = v102
                                                                                                                    } else {
                                                                                                                        v16 = "hello"
                                                                                                                        var v104 int8 = v86
                                                                                                                        _ = v104
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        var v105 bool = (1.5 <= (v9 - (float32(float64(-2.25)) * v9)))
                                                                                                                        var v106 int8 = (^(v86))
                                                                                                                        v107 := GenS16{F0: (^(v3)), F1: v3, F2: -1}
                                                                                                                        _ = v105
                                                                                                                        _ = v106
                                                                                                                        _ = v107
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            if true {
                                                                                                                switch (fi18 % ((2147483647) | 1)) {
                                                                                                                case -2147483648:
                                                                                                                    v99 = v1
                                                                                                                    var v108 uint32 = 966
                                                                                                                    v109 := GenS17{F0: 38, F1: (v21 + v21), F2: (v100 * v1)}
                                                                                                                    _ = v108
                                                                                                                    _ = v109
                                                                                                                default:
                                                                                                                    v110 := GenS8{F0: 1, F1: v100}
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                                v111 := GenS18{F0: (537 == 2147483648), F1: v4}
                                                                                                                v9 = 1e10
                                                                                                                v21 += 0
                                                                                                                switch v10 {
                                                                                                                case "hello":
                                                                                                                    v16 = v10
                                                                                                                case "":
                                                                                                                    fmt.Println(v86)
                                                                                                                case "\n\t":
                                                                                                                    v97 = 214
                                                                                                                default:
                                                                                                                    if true {
                                                                                                                        if ((v86 > (v86 << ((138) & 63))) && true) {
                                                                                                                            v9 = ((((v9 / (((v9 + v9) / 1.0) - 1.5)) + v9) - 1.5) + v9)
                                                                                                                            v99 = v99
                                                                                                                            var v112 float64 = (0.0 / 1e10)
                                                                                                                            _ = v112
                                                                                                                        } else {
                                                                                                                            v4 += v4
                                                                                                                            _ = len(v98)
                                                                                                                        }
                                                                                                                        v100 += (v99 - (v1 | 899))
                                                                                                                        var v113 uint = (v97 * uint(v3))
                                                                                                                        v98 = append(v98, (^(fi17)))
                                                                                                                        v114 := GenS19{F0: 18446744073709551615, F1: (42 == 42)}
                                                                                                                        switch (fi17 ^ fi18) {
                                                                                                                        case 2147483647:
                                                                                                                            var v115 int16 = -32768
                                                                                                                            _ = v115
                                                                                                                        case 32767:
                                                                                                                            if (v4 > v4) {
                                                                                                                                v98 = append(v98, 398)
                                                                                                                            } else {
                                                                                                                                var v116 uint64 = 42
                                                                                                                                _ = v116
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v21 = v21
                                                                                                                        _ = v113
                                                                                                                        _ = v114
                                                                                                                    } else {
                                                                                                                        v117 := GenS14{F0: (v1 % (((v100 - 255)) | 1))}
                                                                                                                        _ = v117
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v111
                                                                                                            }
                                                                                                            _ = v100
                                                                                                            _ = v98
                                                                                                            _ = v99
                                                                                                        }
                                                                                                        v118 := GenS1{F0: 42, F1: v10, F2: (v21 << ((0) & 63))}
                                                                                                        _ = v118
                                                                                                        _ = v97
                                                                                                    }
                                                                                                    _ = v95
                                                                                                    _ = v96
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v119 := []int{(-(fi18)), (^(fi18))}
                                                                                                _ = v119
                                                                                            case "hello":
                                                                                                if false {
                                                                                                    var v120 uint8 = v2
                                                                                                    if false {
                                                                                                        v121 := []int{(fi18 ^ (fi18 - (fi17 % ((fi17) | 1))))}
                                                                                                        _ = v121
                                                                                                    } else {
                                                                                                        var v122 uint64 = 2
                                                                                                        _ = v122
                                                                                                    }
                                                                                                    _ = v120
                                                                                                } else {
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        v21 += -2147483648
                                                                                                    }
                                                                                                    v10 = ("hello" + v10)
                                                                                                    v123 := GenS20{F0: -1, F1: (fi18 >> ((42) & 63))}
                                                                                                    _ = v123
                                                                                                }
                                                                                            }
                                                                                            if false {
                                                                                                fmt.Println(v86, v1)
                                                                                            } else {
                                                                                                if true {
                                                                                                    if true {
                                                                                                        v124 := GenS21{F0: (true || false), F1: false}
                                                                                                        v125 := GenS22{F0: ("a" + ((v10 + v10) + v10)), F1: uint(v3)}
                                                                                                        _ = v124
                                                                                                        _ = v125
                                                                                                    } else {
                                                                                                        v126 := GenS23{F0: 4294967295}
                                                                                                        if true {
                                                                                                            if ((v86 > v86) || (100 < int8(fi17))) {
                                                                                                                var v127 string = (v16 + v16)
                                                                                                                var v128 float32 = v9
                                                                                                                _ = v127
                                                                                                                _ = v128
                                                                                                            } else {
                                                                                                                if (!false) {
                                                                                                                    var v129 int32 = (v21 >> ((838) & 63))
                                                                                                                    v21 += -377
                                                                                                                    _ = v129
                                                                                                                }
                                                                                                                v1 = v1
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch (v16 + v10) {
                                                                                                            case "foo bar":
                                                                                                                v130 := GenS12{F0: v86, F1: 42}
                                                                                                                _ = v130
                                                                                                            default:
                                                                                                                switch fi18 {
                                                                                                                case 128:
                                                                                                                    fmt.Println(fi17, v16)
                                                                                                                case 479:
                                                                                                                    v86 += -1
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        if false {
                                                                                                            v4 = (v4 / (((v4 + (v4 / ((42) | 1)))) | 1))
                                                                                                            if (fi18 <= fi17) {
                                                                                                                v131 := GenS18{F0: (1e10 <= 3.14159), F1: (-(v4))}
                                                                                                                v132 := GenS8{F0: 1000, F1: v1}
                                                                                                                _ = v131
                                                                                                                _ = v132
                                                                                                            } else {
                                                                                                                var v133 float32 = v9
                                                                                                                v10 = ""
                                                                                                                switch fi17 {
                                                                                                                case 32768:
                                                                                                                    if true {
                                                                                                                        if (1 < 2) {
                                                                                                                            var v134 int64 = (v4 ^ v4)
                                                                                                                            if (!true) {
                                                                                                                                if (0 >= 1) {
                                                                                                                                    v2 += v3
                                                                                                                                }
                                                                                                                                if (((((0.1 / 1.5) - float64(v9)) / 0.1) + (0.1 * -2.25)) <= (float64((((v133 + v9) - v133) / (0.1 - v133))) / 1.5)) {
                                                                                                                                    if true {
                                                                                                                                        var v135 int32 = 42
                                                                                                                                        switch ("\n\t" + (v10 + (v16 + "a"))) {
                                                                                                                                        case "a":
                                                                                                                                            var v136 float32 = (3.14159 + v133)
                                                                                                                                            _ = v136
                                                                                                                                        default:
                                                                                                                                            v137 := GenS10{F0: (v1 & 255), F1: 1}
                                                                                                                                            _ = v137
                                                                                                                                        }
                                                                                                                                        var v138 int16 = -1
                                                                                                                                        _ = v135
                                                                                                                                        _ = v138
                                                                                                                                    }
                                                                                                                                    if (42 > v4) {
                                                                                                                                        switch "a" {
                                                                                                                                        case "hello":
                                                                                                                                            delete(v5, (fi17 ^ fi18))
                                                                                                                                        case "":
                                                                                                                                            var v139 string = v16
                                                                                                                                            _ = v139
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v16, v134, fi17)
                                                                                                                                    var v140 uint64 = 0
                                                                                                                                    if true {
                                                                                                                                        v10 = v16
                                                                                                                                        switch int(v1) {
                                                                                                                                        case 127:
                                                                                                                                            var v141 uint16 = 886
                                                                                                                                            fmt.Println(v2, v133, v86)
                                                                                                                                            var v142 int16 = 0
                                                                                                                                            _ = v141
                                                                                                                                            _ = v142
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v143 := []int{}
                                                                                                                                    switch (fi18 - fi17) {
                                                                                                                                    case 42:
                                                                                                                                        v21 = v21
                                                                                                                                    case 9223372036854775807:
                                                                                                                                        v144 := GenS24{F0: 481}
                                                                                                                                        _ = v144
                                                                                                                                    }
                                                                                                                                    _ = v140
                                                                                                                                    _ = v143
                                                                                                                                }
                                                                                                                                var v145 bool = ((v134 ^ v134) > (v134 * v134))
                                                                                                                                _ = v145
                                                                                                                            }
                                                                                                                            _ = v134
                                                                                                                        } else {
                                                                                                                            switch v16 {
                                                                                                                            case "\n\t":
                                                                                                                                v1 -= (v1 & (v1 * v1))
                                                                                                                            default:
                                                                                                                                v146 := []int{fi17}
                                                                                                                                _ = v146
                                                                                                                            }
                                                                                                                            if (false && true) {
                                                                                                                                v16 = (((v16 + "hello") + v16) + v10)
                                                                                                                                if false {
                                                                                                                                    var v147 uint8 = v2
                                                                                                                                    var v148 float32 = -2.25
                                                                                                                                    var v149 bool = false
                                                                                                                                    v9 -= 1.5
                                                                                                                                    _ = v147
                                                                                                                                    _ = v148
                                                                                                                                    _ = v149
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if false {
                                                                                                                            v150 := GenS25{F0: 645, F1: 2, F2: (v1 - (v1 / ((v1) | 1)))}
                                                                                                                            var v151 float32 = v9
                                                                                                                            _ = v150
                                                                                                                            _ = v151
                                                                                                                        } else {
                                                                                                                            v86 = v86
                                                                                                                            switch (fi17 + fi18) {
                                                                                                                            case -8:
                                                                                                                                fmt.Println(fi18)
                                                                                                                            default:
                                                                                                                                v152 := []int{}
                                                                                                                                _ = v152
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                case -884:
                                                                                                                    fmt.Println(v21, v1)
                                                                                                                    var v153 int8 = v86
                                                                                                                    if true {
                                                                                                                        v10 = v16
                                                                                                                        v154 := GenS10{F0: v1, F1: 18446744073709551615}
                                                                                                                        v155 := GenS13{F0: ((v10 + v10) + (v10 + v16)), F1: v16, F2: fi17}
                                                                                                                        switch v16 {
                                                                                                                        case "":
                                                                                                                            var v156 int8 = (v86 + v86)
                                                                                                                            _ = v156
                                                                                                                        case "foo bar":
                                                                                                                            v157 := GenS12{F0: int8(v1), F1: 4294967295}
                                                                                                                            _ = v157
                                                                                                                        }
                                                                                                                        var v158 uint64 = 2
                                                                                                                        _ = v154
                                                                                                                        _ = v155
                                                                                                                        _ = v158
                                                                                                                    } else {
                                                                                                                        var v159 float64 = float64(1e10)
                                                                                                                        _ = v159
                                                                                                                    }
                                                                                                                    _ = v153
                                                                                                                }
                                                                                                                _ = v133
                                                                                                            }
                                                                                                        } else {
                                                                                                            v86 += 100
                                                                                                        }
                                                                                                        _ = v126
                                                                                                    }
                                                                                                    v160 := GenS6{F0: -1, F1: (v10 + v16)}
                                                                                                    _ = v160
                                                                                                } else {
                                                                                                    if ((v21 ^ (v21 * 525)) < int32(v1)) {
                                                                                                        var v161 int16 = 1
                                                                                                        if false {
                                                                                                            if true {
                                                                                                                var v162 int64 = (v4 | v4)
                                                                                                                _ = v162
                                                                                                            }
                                                                                                            switch (v16 + v16) {
                                                                                                            case "\n\t":
                                                                                                                var v163 uint16 = uint16(v2)
                                                                                                                _ = v163
                                                                                                            }
                                                                                                            v86 = (v86 / ((v86) | 1))
                                                                                                        } else {
                                                                                                            fmt.Println(v2, v86, v1)
                                                                                                            if (!false) {
                                                                                                                switch 42 {
                                                                                                                case 0:
                                                                                                                    if true {
                                                                                                                        v164 := GenS26{F0: ((-2.25 / 1.0) == -2.25)}
                                                                                                                        _ = v164
                                                                                                                    } else {
                                                                                                                        v165 := GenS25{F0: 0, F1: 279, F2: (v1 & v1)}
                                                                                                                        v166 := []int{(-(fi17)), (fi18 + (fi18 % (((fi18 % ((952) | 1))) | 1))), (fi18 >> ((707) & 63)), fi17}
                                                                                                                        v10 = "hello"
                                                                                                                        v16 = (v16 + ((v16 + v16) + v16))
                                                                                                                        _ = v165
                                                                                                                        _ = v166
                                                                                                                    }
                                                                                                                case 9223372036854775807:
                                                                                                                    if true {
                                                                                                                        v9 = 0.1
                                                                                                                        var v167 string = "hello"
                                                                                                                        switch fi17 {
                                                                                                                        case -589:
                                                                                                                            if true {
                                                                                                                                var v168 uint8 = (v3 % ((v2) | 1))
                                                                                                                                fmt.Println(v86)
                                                                                                                                _ = v168
                                                                                                                            } else {
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v1)
                                                                                                                                    fmt.Println(v16, v167, v16)
                                                                                                                                    if (65535 <= 2147483647) {
                                                                                                                                        switch v16 {
                                                                                                                                        case "":
                                                                                                                                            if true {
                                                                                                                                                var v169 float64 = 1.5
                                                                                                                                                _ = v169
                                                                                                                                            }
                                                                                                                                        case "foo bar":
                                                                                                                                            if ((v86 + v86) < 0) {
                                                                                                                                                delete(v6, 2147483647)
                                                                                                                                            } else {
                                                                                                                                                var v170 uint64 = 4294967295
                                                                                                                                                var v171 string = v16
                                                                                                                                                _ = v170
                                                                                                                                                _ = v171
                                                                                                                                            }
                                                                                                                                        case "\n\t":
                                                                                                                                            v172 := GenS6{F0: (v86 + 0), F1: (v16 + (v167 + v167))}
                                                                                                                                            _ = v172
                                                                                                                                        }
                                                                                                                                        var v173 uint = 18446744073709551615
                                                                                                                                        _ = v173
                                                                                                                                    }
                                                                                                                                    delete(v6, 2147483648)
                                                                                                                                    if true {
                                                                                                                                        v174 := []int{}
                                                                                                                                        var v175 int64 = (-(v4))
                                                                                                                                        var v176 int64 = int64(v161)
                                                                                                                                        var v177 uint8 = v3
                                                                                                                                        v10 = "hello"
                                                                                                                                        v175 = int64(fi18)
                                                                                                                                        _ = v174
                                                                                                                                        _ = v175
                                                                                                                                        _ = v176
                                                                                                                                        _ = v177
                                                                                                                                    }
                                                                                                                                    fmt.Println(v1)
                                                                                                                                    v1 += 962
                                                                                                                                    v178 := GenS25{F0: 2147483647, F1: 65535, F2: (v1 - v1)}
                                                                                                                                    _ = v178
                                                                                                                                }
                                                                                                                                delete(v6, 42)
                                                                                                                                v179 := []int{int(v2), fi18}
                                                                                                                                if (1 == 895) {
                                                                                                                                    fmt.Println(v3)
                                                                                                                                } else {
                                                                                                                                    v180 := v6[612]
                                                                                                                                    v179 = append(v179, fi18)
                                                                                                                                    if len(v179) > 0 {
                                                                                                                                        v179[0] = int(v21)
                                                                                                                                    }
                                                                                                                                    v181 := GenS27{F0: ((!false) || (v16 <= ((v16 + "\n\t") + "\n\t"))), F1: v180, F2: v10}
                                                                                                                                    _ = v180
                                                                                                                                    _ = v181
                                                                                                                                }
                                                                                                                                _ = v179
                                                                                                                            }
                                                                                                                        case -9223372036854775808:
                                                                                                                            fmt.Println(v16)
                                                                                                                        case 0:
                                                                                                                            var v182 float32 = 0.0
                                                                                                                            _ = v182
                                                                                                                        }
                                                                                                                        _ = v167
                                                                                                                    }
                                                                                                                case 270:
                                                                                                                    var v183 float32 = (3.14159 * (v9 - 0.1))
                                                                                                                    _ = v183
                                                                                                                default:
                                                                                                                    var v184 uint8 = (v2 & (v3 + (^(v2))))
                                                                                                                    _ = v184
                                                                                                                }
                                                                                                            }
                                                                                                            v185 := []int{fi18, fi18}
                                                                                                            _ = v185
                                                                                                        }
                                                                                                        v1 = 1
                                                                                                        _ = v161
                                                                                                    }
                                                                                                    v186 := GenS4{F0: v21, F1: (-(v21))}
                                                                                                    var v187 uint8 = (v3 << ((250) & 63))
                                                                                                    _ = v186
                                                                                                    _ = v187
                                                                                                }
                                                                                            }
                                                                                            v1 = v1
                                                                                        } else {
                                                                                            var v188 float32 = v9
                                                                                            _ = v188
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v189 int8 = v86
                                                                                _ = v189
                                                                                _ = v86
                                                                            } else {
                                                                                var v190 string = v16
                                                                                v191 := GenS8{F0: 434, F1: (v1 << ((693) & 63))}
                                                                                fmt.Println(v21)
                                                                                _ = v190
                                                                                _ = v191
                                                                            }
                                                                        case 2147483647:
                                                                            if (true && (fi18 != (fi17 << ((4294967295) & 63)))) {
                                                                                v192 := v5[(fi18 / (((^(fi17))) | 1))]
                                                                                var v193 uint8 = 45
                                                                                fmt.Println(v21, fi18)
                                                                                _ = v192
                                                                                _ = v193
                                                                            }
                                                                        }
                                                                        switch (v16 + v16) {
                                                                        case "hello":
                                                                            v194 := []int{fi18, (fi17 + fi18), 36}
                                                                            v195 := GenS9{F0: 611, F1: 32767, F2: (v1 * v1)}
                                                                            _ = v194
                                                                            _ = v195
                                                                        case "\n\t":
                                                                            fmt.Println(v4)
                                                                        case "a":
                                                                            var v196 uint16 = uint16(v3)
                                                                            _ = v196
                                                                        default:
                                                                            v197 := GenS28{F0: v21, F1: true}
                                                                            v198 := v5[(fi17 / (((fi18 - 32768)) | 1))]
                                                                            v199 := []int{-963}
                                                                            _ = v197
                                                                            _ = v198
                                                                            _ = v199
                                                                        }
                                                                        fmt.Println(fi18, fi18)
                                                                    }
                                                                }
                                                                v200 := GenS11{F0: 4294967295, F1: 255}
                                                                fmt.Println(v16, fi18, v10)
                                                                v201 := GenS22{F0: v16, F1: 4294967295}
                                                                v202 := GenS29{F0: 65535, F1: uint(v1), F2: (fi18 | fi18)}
                                                                var v203 uint8 = 47
                                                                _ = v200
                                                                _ = v201
                                                                _ = v202
                                                                _ = v203
                                                                _ = v21
                                                            } else {
                                                                var v204 float32 = (v9 + (float32(3.14159) * v9))
                                                                _ = v204
                                                            }
                                                            var v205 uint = 4294967295
                                                            if (2 > 0) {
                                                                v9 -= 1e10
                                                                if true {
                                                                    fmt.Println(v205)
                                                                    var v206 uint8 = (v2 | (v2 % (((v2 << ((v205) & 63))) | 1)))
                                                                    fmt.Println(v2)
                                                                    _ = v206
                                                                }
                                                                fmt.Println(v9)
                                                            } else {
                                                                var v207 int8 = -92
                                                                v208 := v5[(^(fi18))]
                                                                var v209 uint16 = 1
                                                                _ = v207
                                                                _ = v208
                                                                _ = v209
                                                            }
                                                            _ = v205
                                                        }
                                                    case "\n\t":
                                                        v210 := GenS30{F0: (fi18 >> ((0) & 63)), F1: 255}
                                                        var v211 uint16 = v1
                                                        _ = v210
                                                        _ = v211
                                                    }
                                                    var v212 bool = false
                                                    v3 -= (v3 * v3)
                                                    _ = v212
                                                } else {
                                                    if (!true) {
                                                        v16 = ("hello" + v16)
                                                        v2 = (^(v3))
                                                    } else {
                                                        fmt.Println(v10, v2)
                                                    }
                                                    v213 := v5[int(v1)]
                                                    _ = v213
                                                }
                                            case "hello":
                                                v214 := []int{fi18, fi17, (fi17 % ((fi18) | 1)), (-(fi18))}
                                                _ = v214
                                            default:
                                                var v215 uint = 65535
                                                _ = v215
                                            }
                                            v216 := GenS31{F0: (true || false), F1: 1, F2: 1000}
                                            if ((fi18 ^ (fi17 * (fi18 * (fi17 + (fi18 / ((127) | 1)))))) > fi18) {
                                                v16 = ("\n\t" + v16)
                                            } else {
                                                v6[4294967295] = 273
                                                switch fi17 {
                                                case -263:
                                                    switch fi17 {
                                                    case 2147483647:
                                                        var v217 string = "\n\t"
                                                        _ = v217
                                                    case 0:
                                                        var v218 int64 = 87
                                                        _ = v218
                                                    }
                                                case 9223372036854775807:
                                                    if false {
                                                        switch fi17 {
                                                        case 9223372036854775807:
                                                            v10 = v10
                                                        case 42:
                                                            _ = len(v6)
                                                        }
                                                        v2 = v3
                                                        v4 = v4
                                                    }
                                                default:
                                                    if false {
                                                        switch fi17 {
                                                        case 42:
                                                            switch v16 {
                                                            case "foo bar":
                                                                if true {
                                                                    switch fi18 {
                                                                    case 2147483647:
                                                                        v219 := GenS4{F0: 42, F1: -395}
                                                                        _ = v219
                                                                    case 32768:
                                                                        fmt.Println(v9, fi17, fi18)
                                                                    }
                                                                    var v220 int8 = 42
                                                                    _ = v220
                                                                } else {
                                                                    switch fi18 {
                                                                    case 667:
                                                                        v2 += (v3 | v3)
                                                                    }
                                                                    fmt.Println(v10, v16, v4)
                                                                    if (60 != -42) {
                                                                        if (658 == 42) {
                                                                            v221 := []int{(fi18 | (fi18 % (((fi18 ^ -2147483648)) | 1))), fi17}
                                                                            if (610 > 42) {
                                                                                switch fi17 {
                                                                                case 9223372036854775807:
                                                                                    if true {
                                                                                        v3 += (v2 + (v2 << ((465) & 63)))
                                                                                    }
                                                                                case -1:
                                                                                    v3 = 135
                                                                                case 0:
                                                                                    delete(v6, 751)
                                                                                }
                                                                                switch (v10 + "") {
                                                                                case "":
                                                                                    var v222 int64 = (v4 % ((9223372036854775807) | 1))
                                                                                    _ = v222
                                                                                case "a":
                                                                                    v16 = v16
                                                                                case "\n\t":
                                                                                    v16 = "foo bar"
                                                                                }
                                                                                fmt.Println(v2, v16, fi18)
                                                                            } else {
                                                                                var v223 uint64 = 2
                                                                                v224 := GenS2{F0: (v4 / (((-(v4))) | 1)), F1: 1000}
                                                                                _ = v223
                                                                                _ = v224
                                                                            }
                                                                            v221 = append(v221, (^(fi17)))
                                                                            if true {
                                                                                v9 = v9
                                                                                v9 = (float32(((1.5 - 1.5) + ((float64(v9) - -2.25) * 0.1))) / 1.0)
                                                                            } else {
                                                                                if (v1 <= v1) {
                                                                                    if true {
                                                                                        if (0 > 4294967295) {
                                                                                            fmt.Println(v9, v16)
                                                                                            var v225 int = fi18
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v3, v9)
                                                                                            case "":
                                                                                                if (2147483648 == uint32(v1)) {
                                                                                                    v10 = v10
                                                                                                    v4 = (v4 << ((0) & 63))
                                                                                                }
                                                                                                v9 = v9
                                                                                                var v226 float32 = v9
                                                                                                _ = v226
                                                                                            }
                                                                                            _ = v225
                                                                                        }
                                                                                    }
                                                                                    v16 = v10
                                                                                }
                                                                            }
                                                                            switch (fi17 >> ((255) & 63)) {
                                                                            case -1:
                                                                                if (!true) {
                                                                                    _ = len(v6)
                                                                                    if (988 >= 1) {
                                                                                        switch -9223372036854775808 {
                                                                                        case 626:
                                                                                            delete(v5, 2147483647)
                                                                                        case 468:
                                                                                            var v227 uint32 = 4294967295
                                                                                            v9 += float32(float64(v9))
                                                                                            _ = v227
                                                                                        case -801:
                                                                                            fmt.Println(fi17)
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            switch (v16 + ("" + (v10 + v16))) {
                                                                                            case "hello":
                                                                                                v2 = 42
                                                                                            case "\n\t":
                                                                                                var v228 float32 = v9
                                                                                                _ = v228
                                                                                            case "foo bar":
                                                                                                if len(v221) > 0 {
                                                                                                    v221[0] = -659
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            _ = len(v221)
                                                                                        case "\n\t":
                                                                                            if (1 < 4294967295) {
                                                                                                delete(v6, 42)
                                                                                                switch v16 {
                                                                                                case "hello":
                                                                                                    var v229 uint16 = v1
                                                                                                    _ = v229
                                                                                                case "a":
                                                                                                    if ("foo bar" <= v10) {
                                                                                                        delete(v6, 607)
                                                                                                    } else {
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            if true {
                                                                                                                var v230 int16 = 1000
                                                                                                                v231 := GenS6{F0: 127, F1: ("hello" + v16)}
                                                                                                                _ = v230
                                                                                                                _ = v231
                                                                                                            }
                                                                                                        case "":
                                                                                                            fmt.Println(fi17, v2)
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v10, fi17)
                                                                                                        }
                                                                                                        var v232 uint = 785
                                                                                                        switch fi18 {
                                                                                                        case 127:
                                                                                                            var v233 int = fi17
                                                                                                            _ = v233
                                                                                                        default:
                                                                                                            switch "hello" {
                                                                                                            case "foo bar":
                                                                                                                var v234 float64 = 1.5
                                                                                                                _ = v234
                                                                                                            case "a":
                                                                                                                switch fi18 {
                                                                                                                case -9223372036854775808:
                                                                                                                    v235 := v6[641]
                                                                                                                    _ = v235
                                                                                                                default:
                                                                                                                    var v236 uint8 = 1
                                                                                                                    _ = v236
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                if true {
                                                                                                                    if (-2147483648 > 0) {
                                                                                                                        if len(v221) > 0 {
                                                                                                                            v221[0] = fi18
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v237 int32 = 318
                                                                                                                    _ = v237
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        if true {
                                                                                                                            v2 = 133
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v238 := GenS10{F0: v1, F1: (v232 % ((793) | 1))}
                                                                                                                    _ = v238
                                                                                                                }
                                                                                                            default:
                                                                                                                var v239 float32 = (1.5 * ((v9 * 1.5) * v9))
                                                                                                                _ = v239
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v232
                                                                                                    }
                                                                                                default:
                                                                                                    v3 += 0
                                                                                                }
                                                                                            } else {
                                                                                                var v240 uint64 = 1
                                                                                                v241 := GenS25{F0: 4294967295, F1: 1, F2: (^(v1))}
                                                                                                v2 = (v3 % ((1) | 1))
                                                                                                _ = v240
                                                                                                _ = v241
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v242 uint8 = (^(v3))
                                                                                        var v243 bool = (65535 != 42)
                                                                                        _ = v242
                                                                                        _ = v243
                                                                                    }
                                                                                    var v244 int32 = 925
                                                                                    _ = v244
                                                                                } else {
                                                                                    fmt.Println(v9)
                                                                                    var v245 uint64 = 42
                                                                                    if true {
                                                                                        v3 += v2
                                                                                    } else {
                                                                                        fmt.Println(v1, v1)
                                                                                        if false {
                                                                                            v1 -= v1
                                                                                        } else {
                                                                                            _ = len(v6)
                                                                                            v1 = (v1 * (v1 ^ v1))
                                                                                            if (v4 >= -1) {
                                                                                                switch v10 {
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        v5[4294967295] = (((v10 + v10) + ((v10 + (("\n\t" + v16) + "a")) + (v16 + ("" + v10)))) + (v16 + v16))
                                                                                                    } else {
                                                                                                        v246 := GenS2{F0: (-(v4)), F1: 0}
                                                                                                        if len(v221) > 0 {
                                                                                                            v221[0] = 127
                                                                                                        }
                                                                                                        _ = v246
                                                                                                    }
                                                                                                }
                                                                                                var v247 uint8 = (v3 / ((v3) | 1))
                                                                                                v248 := GenS25{F0: 2147483648, F1: 697, F2: v1}
                                                                                                _ = v247
                                                                                                _ = v248
                                                                                            }
                                                                                        }
                                                                                        switch (-(fi17)) {
                                                                                        case 0:
                                                                                            var v249 int64 = v4
                                                                                            _ = v249
                                                                                        default:
                                                                                            _ = len(v221)
                                                                                        }
                                                                                    }
                                                                                    switch fi18 {
                                                                                    case 63:
                                                                                        v2 -= v3
                                                                                    case 32768:
                                                                                        v4 += v4
                                                                                    case 128:
                                                                                        fmt.Println(fi18, fi18, v9)
                                                                                    default:
                                                                                        v245 -= v245
                                                                                    }
                                                                                    fmt.Println(v3)
                                                                                    _ = v245
                                                                                }
                                                                            case 536:
                                                                                _ = len(v5)
                                                                            case 757:
                                                                                if len(v221) > 0 {
                                                                                    v221[0] = fi18
                                                                                }
                                                                                v9 -= (v9 + (0.0 - 1e10))
                                                                                fmt.Println(v16)
                                                                            default:
                                                                                var v250 int8 = 62
                                                                                _ = v250
                                                                            }
                                                                            _ = v221
                                                                        } else {
                                                                            v251 := []int{fi18, 885, 0}
                                                                            _ = v251
                                                                        }
                                                                        var v252 int64 = v4
                                                                        _ = v252
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v253 := v5[42]
                                                        v16 = "foo bar"
                                                        fmt.Println(v9, v3)
                                                        _ = v253
                                                    } else {
                                                        var v254 uint32 = 1
                                                        _ = v254
                                                    }
                                                    v9 += (v9 * ((((3.14159 + (0.0 * 0.1)) - 3.14159) - v9) - 1.5))
                                                    v255 := []int{fi17, fi18, 4294967295}
                                                    _ = v255
                                                }
                                                v256 := []int{}
                                                _ = v256
                                            }
                                            if (!true) {
                                                if (!true) {
                                                    v5[fi17] = v10
                                                    if (false && ((v2 >> ((657) & 63)) > (^(v3)))) {
                                                        v2 += 228
                                                        var v257 int = fi17
                                                        v258 := GenS26{F0: (9223372036854775807 == (v4 % ((-1) | 1)))}
                                                        _ = v257
                                                        _ = v258
                                                    } else {
                                                        var v259 float64 = ((0.1 - 3.14159) / 1.5)
                                                        _ = v259
                                                    }
                                                    v5[-2147483648] = v16
                                                }
                                                switch (fi17 % ((fi17) | 1)) {
                                                case 42:
                                                    v260 := []int{fi17}
                                                    _ = v260
                                                case 113:
                                                    var v261 float32 = v9
                                                    _ = v261
                                                default:
                                                    if false {
                                                        v262 := []int{(fi17 * (fi17 ^ fi17))}
                                                        var v263 bool = (false && (uint(v3) > 18446744073709551615))
                                                        _ = v262
                                                        _ = v263
                                                    }
                                                }
                                                v16 = v16
                                                v264 := GenS32{F0: 65535, F1: uint64(v2)}
                                                switch ("a" + v16) {
                                                case "a":
                                                    switch (v16 + "hello") {
                                                    case "hello":
                                                        var v265 int64 = (v4 ^ -613)
                                                        _ = v265
                                                    case "":
                                                        var v266 uint32 = 616
                                                        v5[(-(fi17))] = v10
                                                        _ = v266
                                                    case "a":
                                                        fmt.Println(v1, v16)
                                                    default:
                                                        v6[2147483647] = uint32(v2)
                                                    }
                                                }
                                                _ = v264
                                            }
                                            _ = v216
                                        }
                                    }
                                    delete(v5, (fi17 + fi18))
                                    if (0 > 811) {
                                        v267 := GenS33{F0: -2147483648, F1: 0, F2: 363}
                                        v268 := []int{(fi17 >> ((1) & 63)), -2147483648, fi18, (-(fi17))}
                                        fmt.Println(fi17)
                                        _ = v267
                                        _ = v268
                                    }
                                    if false {
                                        fmt.Println(v3, v10, v9)
                                        v269 := GenS34{F0: 1, F1: 0, F2: v2}
                                        v270 := GenS13{F0: ((v16 + v16) + v16), F1: "\n\t", F2: fi18}
                                        v10 = ("foo bar" + (v16 + "foo bar"))
                                        fmt.Println(v3, v2, v16)
                                        _ = v269
                                        _ = v270
                                    } else {
                                        var v271 uint = 42
                                        _ = v271
                                    }
                                } else {
                                    var v272 uint64 = 4294967295
                                    var v273 bool = false
                                    var v274 string = "foo bar"
                                    _ = v272
                                    _ = v273
                                    _ = v274
                                }
                            case 329:
                                _ = len(v6)
                            case 128:
                                var v275 int16 = 705
                                v4 += 1
                                _ = v275
                            default:
                                if (!((-2147483648 <= 612) || (!((false && false) || (fi17 >= fi17))))) {
                                    v9 = (0.1 - ((float32(0.0) - float32(3.14159)) * (0.1 - (float32((3.14159 + 1e10)) - 0.0))))
                                } else {
                                    fmt.Println(v9, v3, v2)
                                    v276 := GenS20{F0: -1, F1: fi18}
                                    if (!((^(fi17)) < fi17)) {
                                        var v277 int32 = 0
                                        _ = v277
                                    } else {
                                        var v278 int64 = (v4 + (v4 % ((1) | 1)))
                                        v3 -= (v2 ^ 2)
                                        _ = v278
                                    }
                                    _ = v276
                                }
                            }
                            if (127 < 0) {
                                switch (("a" + "") + "hello") {
                                case "hello":
                                    var v279 bool = false
                                    _ = v279
                                case "foo bar":
                                    v5[9223372036854775807] = "foo bar"
                                case "":
                                    var v280 float64 = (3.14159 / 1.5)
                                    _ = v280
                                default:
                                    delete(v5, fi17)
                                }
                                var v281 bool = (true && true)
                                var v282 uint8 = v2
                                _ = v281
                                _ = v282
                            } else {
                                if (("a" + ((v10 + v16) + v10)) >= "") {
                                    fmt.Println(v2, v4, v2)
                                    var v283 uint8 = uint8(v1)
                                    var v284 int8 = 127
                                    _ = v283
                                    _ = v284
                                }
                                switch (v16 + v16) {
                                case "hello":
                                    switch ((("" + ((v10 + "hello") + ((v16 + v10) + v10))) + ("\n\t" + "")) + v10) {
                                    case "hello":
                                        _ = len(v6)
                                    case "foo bar":
                                        v285 := GenS7{F0: v4, F1: 0, F2: 2}
                                        _ = v285
                                    default:
                                        switch ("hello" + v16) {
                                        case "a":
                                            fmt.Println(fi18, fi18)
                                        default:
                                            var v286 uint8 = 127
                                            _ = v286
                                        }
                                    }
                                case "a":
                                    if (!true) {
                                        v3 += 163
                                    } else {
                                        var v287 string = v10
                                        _ = v287
                                    }
                                    v288 := []int{89, fi18, -1, (fi18 / (((fi17 >> ((1) & 63))) | 1))}
                                    var v289 string = ""
                                    var v290 uint16 = v1
                                    v16 = (v10 + (v10 + (v16 + ((v10 + (((v289 + "a") + v16) + v289)) + v16))))
                                    if ((-(v4)) == (v4 + (v4 ^ v4))) {
                                        var v291 bool = true
                                        if v291 {
                                            var v292 uint32 = 4294967295
                                            var v293 bool = v291
                                            _ = v292
                                            _ = v293
                                        }
                                        _ = v291
                                    }
                                    if false {
                                        if (((!false) || (-728 > 42)) && (false || false)) {
                                            var v294 bool = true
                                            fmt.Println(v294)
                                            switch v10 {
                                            case "":
                                                v3 = v2
                                            case "\n\t":
                                                if ((!(!(true && v294))) || (v2 <= v3)) {
                                                    var v295 int32 = -1
                                                    var v296 uint16 = v290
                                                    var v297 int32 = v295
                                                    _ = v295
                                                    _ = v296
                                                    _ = v297
                                                }
                                            }
                                            _ = v294
                                        } else {
                                            v298 := GenS35{F0: true, F1: 2}
                                            if false {
                                                var v299 uint = 42
                                                _ = v299
                                            }
                                            var v300 float64 = (1e10 - float64((0.1 + v9)))
                                            _ = v298
                                            _ = v300
                                        }
                                        fmt.Println(v10)
                                        v301 := GenS31{F0: (0.0 < ((-2.25 + ((1.5 - 0.0) + (float32(((0.0 * 0.0) - (-2.25 - 3.14159))) - v9))) - ((v9 + v9) / -2.25))), F1: 42, F2: int16(v4)}
                                        _ = v301
                                    }
                                    v302 := GenS36{F0: 0, F1: 2147483647, F2: 2147483647}
                                    _ = v288
                                    _ = v289
                                    _ = v290
                                    _ = v302
                                default:
                                    var v303 int64 = (^(v4))
                                    _ = v303
                                }
                            }
                        }
                        _ = fi18
                    }
                    for fi304 := 0; fi304 < 13; fi304++ {
                        v9 -= ((v9 - v9) - (float32(1.5) + (v9 - float32(float64(v9)))))
                        _ = fi304
                    }
                    _ = fi17
                }
            default:
                v3 = (v3 << ((255) & 63))
            }
            if ((42 < 4294967295) && true) {
                v3 -= v2
            }
            _ = v16
        }
        switch 128 {
        case -2147483648:
            if (!false) {
                v10 = v10
                v305 := []int{106, 127, -813}
                _ = v305
            }
        case 2147483647:
            var v306 bool = true
            _ = v306
        default:
            var v307 float32 = v9
            _ = v307
        }
    } else {
        if true {
            fmt.Println(v10)
        } else {
            v308 := []int{279, 219}
            _ = v308
        }
    }
    var v309 bool = false
    var v310 uint64 = 4294967295
    switch (v10 + v10) {
    case "hello":
        v311 := v6[65535]
        _ = v311
    }
    var v312 int32 = -1
    for fi313 := 0; fi313 < 18; fi313++ {
        if true {
            switch "a" {
            case "\n\t":
                fmt.Println(v1)
            default:
                v1 = v1
            }
        }
        for fi314 := 0; fi314 < 12; fi314++ {
            _ = len(v6)
            v315 := GenS8{F0: 1, F1: 0}
            _ = fi314
            _ = v315
        }
        fmt.Println(v9)
        _ = fi313
    }
    for fi316 := 0; fi316 < 17; fi316++ {
        var v317 int = (fi316 & fi316)
        v318 := GenS20{F0: 1, F1: v317}
        v6[1] = 65535
        _ = fi316
        _ = v317
        _ = v318
    }
    v2 += 2
    v6[4294967295] = 1
    v3 = uint8(v1)
    if true {
        if v309 {
            v10 = v10
        }
        var v319 string = v10
        _ = v319
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v309:", v309)
    fmt.Println("v310:", v310)
    fmt.Println("v312:", v312)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v309
    _ = v310
    _ = v312
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
