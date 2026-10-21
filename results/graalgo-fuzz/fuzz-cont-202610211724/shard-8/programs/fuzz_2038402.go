package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 bool; F2 uint8 }
type GenS7 struct { F0 bool; F1 string; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS9 struct { F0 string; F1 uint32; F2 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int8; F1 bool; F2 bool }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint; F1 string; F2 uint64 }
type GenS14 struct { F0 uint64; F1 bool }
type GenS15 struct { F0 uint; F1 int8 }
type GenS16 struct { F0 uint16; F1 int }
type GenS17 struct { F0 uint32; F1 bool }
type GenS18 struct { F0 uint8; F1 uint8 }
type GenS19 struct { F0 int16; F1 uint; F2 int }
type GenS20 struct { F0 int32; F1 bool; F2 uint32 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 bool; F1 bool }
type GenS24 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 65535
    v3 := make(map[string]string)
    var v4 uint8 = 0
    v5 := []uint32{v2, 337, 65535}
    v6 := make(map[int]int)
    fmt.Println(v4, v2)
    fmt.Println(v2, v2)
    v4 -= 255
    _ = len(v1)
    fmt.Println(v2)
    v7 := GenS1{F0: (3.14159 > 0.1)}
    switch "hello" {
    case "a":
        v4 = (v4 & v4)
    case "foo bar":
        var v8 uint64 = 4294967295
        _ = v8
    case "hello":
        var v9 int32 = 361
        _ = v9
    default:
        switch ("\n\t" + "\n\t") {
        case "a":
            var v10 int = 2147483647
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 7; fi11++ {
                v2 -= 65535
                switch 1 {
                case 42:
                    for fi12 := 0; fi12 < 14; fi12++ {
                        fmt.Println(v2, fi11)
                        v13 := v3["hello"]
                        _ = fi12
                        _ = v13
                    }
                case -1:
                    var v14 int64 = 0
                    _ = v14
                default:
                    switch fi11 {
                    case 837:
                        var v15 uint64 = 502
                        _ = v15
                    case -9223372036854775808:
                        if (v4 != (v4 ^ (v4 >> ((237) & 63)))) {
                            if false {
                                for fi16 := 0; fi16 < 19; fi16++ {
                                    v17 := GenS1{F0: (false && true)}
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            v2 = (v2 & 0)
                            for fi18 := 0; fi18 < 9; fi18++ {
                                var v19 int16 = 619
                                fmt.Println(fi18, fi11)
                                _ = fi18
                                _ = v19
                            }
                        } else {
                            fmt.Println(v4, v4)
                            for fi20 := 0; fi20 < 4; fi20++ {
                                var v21 bool = false
                                if true {
                                    delete(v6, (fi20 * fi20))
                                    if v21 {
                                        switch "a" {
                                        case "foo bar":
                                            v22 := GenS1{F0: v21}
                                            _ = v22
                                        case "a":
                                            v21 = (float32(1.5) <= float32(0.0))
                                        case "":
                                            var v23 uint = 1
                                            _ = v23
                                        }
                                        if (v21 || false) {
                                            v24 := v1[uint32(v4)]
                                            _ = v24
                                        } else {
                                            fmt.Println(v4, v2)
                                        }
                                        v21 = v21
                                    } else {
                                        v4 = (v4 | (v4 + v4))
                                        _ = len(v5)
                                        fmt.Println(fi11, v4, fi11)
                                        if v21 {
                                            v21 = v21
                                            var v25 uint16 = 65535
                                            _ = v25
                                        }
                                        if (v4 == v4) {
                                            switch "" {
                                            case "hello":
                                                if v21 {
                                                    var v26 int32 = 1
                                                    if (v21 && v21) {
                                                        _ = len(v5)
                                                    } else {
                                                        var v27 string = "a"
                                                        var v28 float32 = ((0.0 + 1e10) / float32(0.1))
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                    _ = v26
                                                }
                                            }
                                            if v21 {
                                                switch fi11 {
                                                case 32768:
                                                    if ((32767 >= 1) || (112 > fi20)) {
                                                        if (!v21) {
                                                            v2 -= 782
                                                            var v29 int16 = int16(fi20)
                                                            var v30 int64 = 42
                                                            v4 += v4
                                                            if v21 {
                                                                fmt.Println(v4, v21)
                                                                var v31 bool = true
                                                                _ = len(v3)
                                                                if (true || v21) {
                                                                    fmt.Println(v2, v4, v2)
                                                                    var v32 string = "foo bar"
                                                                    switch v32 {
                                                                    case "foo bar":
                                                                        v33 := GenS2{F0: uint16(v2), F1: 1, F2: 883}
                                                                        _ = v33
                                                                    case "hello":
                                                                        v32 = v32
                                                                    case "a":
                                                                        switch -1 {
                                                                        case 128:
                                                                            v34 := GenS3{F0: v21}
                                                                            _ = v34
                                                                        }
                                                                        if len(v5) > 0 {
                                                                            v5[0] = uint32(v4)
                                                                        }
                                                                    }
                                                                    _ = v32
                                                                }
                                                                _ = v31
                                                            }
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    } else {
                                                        if false {
                                                            if (true || v21) {
                                                                var v35 string = "foo bar"
                                                                v36 := GenS2{F0: 823, F1: 2, F2: -577}
                                                                v21 = (316 <= 18446744073709551615)
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                            if (v21 || true) {
                                                                v2 = 4294967295
                                                                v37 := GenS3{F0: (!v21)}
                                                                _ = v37
                                                            }
                                                            var v38 float32 = (3.14159 - 1e10)
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                fmt.Println(v38, v4)
                                                                v39 := GenS2{F0: 2, F1: 2, F2: 0}
                                                                if v21 {
                                                                    v3["\n\t"] = "a"
                                                                } else {
                                                                    fmt.Println(v4)
                                                                }
                                                                _ = v39
                                                            }
                                                            _ = len(v1)
                                                            _ = v38
                                                        } else {
                                                            v21 = (v4 >= 128)
                                                            v4 -= v4
                                                            fmt.Println(v2)
                                                        }
                                                    }
                                                default:
                                                    _ = len(v5)
                                                    fmt.Println(fi11)
                                                    v40 := GenS3{F0: true}
                                                    v2 += (v2 % ((65535) | 1))
                                                    fmt.Println(v21, v2, fi11)
                                                    _ = v40
                                                }
                                                var v41 float64 = 0.0
                                                v42 := GenS4{F0: v21, F1: fi20, F2: (-126 != -1)}
                                                _ = v41
                                                _ = v42
                                            } else {
                                                switch "a" {
                                                case "\n\t":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v43 float32 = (((1.5 * (1e10 - 0.0)) * -2.25) - 3.14159)
                                                        _ = v43
                                                    case "hello":
                                                        v44 := GenS1{F0: v21}
                                                        _ = v44
                                                    }
                                                case "foo bar":
                                                    v4 = v4
                                                case "a":
                                                    var v45 int16 = -646
                                                    _ = v45
                                                default:
                                                    _ = len(v6)
                                                }
                                                if v21 {
                                                    switch "a" {
                                                    case "":
                                                        v46 := GenS4{F0: v21, F1: fi11, F2: v21}
                                                        _ = v46
                                                    default:
                                                        v21 = v21
                                                    }
                                                    switch 726 {
                                                    case 340:
                                                        var v47 int32 = 0
                                                        _ = v47
                                                    case 32767:
                                                        if len(v5) > 0 {
                                                            v5[0] = 2147483648
                                                        }
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                    if v21 {
                                                        var v48 int8 = 42
                                                        v21 = ("foo bar" != ((("a" + "\n\t") + "hello") + "hello"))
                                                        var v49 float32 = (0.0 - 0.1)
                                                        _ = len(v5)
                                                        if (v21 || false) {
                                                            v50 := v3[""]
                                                            _ = v50
                                                        }
                                                        if v21 {
                                                            v51 := GenS4{F0: v21, F1: fi20, F2: true}
                                                            if false {
                                                                fmt.Println(fi11, v2)
                                                                if len(v5) > 0 {
                                                                    v5[0] = v2
                                                                }
                                                                v52 := GenS5{F0: 747}
                                                                _ = v52
                                                            } else {
                                                                fmt.Println(v2, v2, v2)
                                                            }
                                                            _ = v51
                                                        }
                                                        var v53 bool = ("hello" > "\n\t")
                                                        _ = len(v5)
                                                        if v53 {
                                                            v49 = ((0.0 - ((1.5 - 3.14159) / (v49 - ((float32(1e10) / 1.5) + 0.0)))) - (v49 + (v49 + float32(-2.25))))
                                                            if (v21 || true) {
                                                                fmt.Println(v48, v49)
                                                                fmt.Println(v2, v49)
                                                            }
                                                        }
                                                        v48 -= 127
                                                        _ = v48
                                                        _ = v49
                                                        _ = v53
                                                    } else {
                                                        v21 = (!((fi20 / (((fi11 << ((255) & 63))) | 1)) <= (fi11 ^ 127)))
                                                        var v54 int16 = -283
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v55 uint8 = v4
                                                            _ = v55
                                                        case "\n\t":
                                                            v4 += (v4 ^ (v4 * (v4 * 2)))
                                                        case "":
                                                            v4 -= 1
                                                        default:
                                                            v21 = false
                                                        }
                                                        _ = v54
                                                    }
                                                }
                                                if (258 <= 9223372036854775807) {
                                                    fmt.Println(fi11, v2, fi11)
                                                    if false {
                                                        v4 += v4
                                                    } else {
                                                        fmt.Println(fi11, fi11)
                                                    }
                                                } else {
                                                    v4 = (^(v4))
                                                    v56 := v3[("\n\t" + "foo bar")]
                                                    if v21 {
                                                        if (!true) {
                                                            switch fi20 {
                                                            case 0:
                                                                var v57 int8 = 71
                                                                _ = v57
                                                            case -2147483648:
                                                                v2 += v2
                                                            case 154:
                                                                var v58 string = (v56 + v56)
                                                                _ = v58
                                                            default:
                                                                v21 = true
                                                            }
                                                            v56 = v56
                                                            delete(v6, (fi20 + (fi11 | 128)))
                                                            v4 = (v4 / ((v4) | 1))
                                                            v59 := GenS3{F0: true}
                                                            switch (fi11 * fi20) {
                                                            case 127:
                                                                v60 := GenS1{F0: ((0 >= (v4 / (((v4 << ((uint(v4)) & 63))) | 1))) && (!(!(!(!true)))))}
                                                                _ = v60
                                                            }
                                                            _ = len(v6)
                                                            _ = v59
                                                        }
                                                    }
                                                    _ = v56
                                                }
                                            }
                                        }
                                    }
                                    v61 := GenS4{F0: v21, F1: 42, F2: v21}
                                    switch -857 {
                                    case 0:
                                        v4 += 128
                                    }
                                    var v62 int32 = -540
                                    _ = v61
                                    _ = v62
                                }
                                _ = fi20
                                _ = v21
                            }
                            fmt.Println(v2, v2)
                        }
                    default:
                        if (!(!true)) {
                            fmt.Println(v2, v4, fi11)
                            switch 4294967295 {
                            case -889:
                                if (false && true) {
                                    switch fi11 {
                                    case 947:
                                        v3["a"] = ((("hello" + "\n\t") + "a") + "foo bar")
                                    case 2147483647:
                                        v63 := GenS6{F0: 0, F1: (0 > 0), F2: 2}
                                        _ = v63
                                    case -2147483648:
                                        var v64 bool = (-76 == 1)
                                        _ = v64
                                    default:
                                        if true {
                                            v4 -= v4
                                        }
                                    }
                                }
                            default:
                                if (("foo bar" + ("foo bar" + "a")) == "\n\t") {
                                    switch 874 {
                                    case -805:
                                        for fi65 := 0; fi65 < 5; fi65++ {
                                            v4 = (^(v4))
                                            _ = fi65
                                        }
                                    }
                                    for fi66 := 0; fi66 < 7; fi66++ {
                                        var v67 int = (fi66 - (-(fi11)))
                                        if (v67 != fi66) {
                                            var v68 float64 = 3.14159
                                            _ = v68
                                        }
                                        _ = fi66
                                        _ = v67
                                    }
                                } else {
                                    v69 := GenS4{F0: (!true), F1: fi11, F2: (-2.25 != ((0.1 * 0.1) * 3.14159))}
                                    if false {
                                        if (-1 == -38) {
                                            if false {
                                                for fi70 := 0; fi70 < 7; fi70++ {
                                                    var v71 string = "a"
                                                    v71 = v71
                                                    _ = fi70
                                                    _ = v71
                                                }
                                            }
                                            v2 = (v2 | v2)
                                            var v72 float32 = (0.1 * (float32(((float64(((-2.25 / 3.14159) * (0.0 * ((-2.25 + 1e10) * (0.1 / ((1e10 + 1.5) / (3.14159 * (0.1 + -2.25)))))))) / 1.5) * float64(0.1))) - 1.5))
                                            _ = v72
                                        } else {
                                            fmt.Println(v4)
                                        }
                                    }
                                    _ = v69
                                }
                            }
                        } else {
                            if (128 > (v4 / ((v4) | 1))) {
                                v73 := GenS2{F0: 655, F1: 412, F2: -212}
                                if (-1 != -655) {
                                    if false {
                                        v74 := GenS1{F0: true}
                                        v75 := GenS7{F0: (!(((fi11 >> ((1) & 63)) == fi11) || (255 <= 539))), F1: ("a" + "\n\t"), F2: -739}
                                        _ = v74
                                        _ = v75
                                    } else {
                                        if true {
                                            fmt.Println(v2, v4)
                                            v76 := GenS8{F0: 255, F1: v2, F2: v4}
                                            _ = v76
                                        }
                                        switch (fi11 / ((-2147483648) | 1)) {
                                        case -9223372036854775808:
                                            v3["hello"] = "a"
                                        case 9223372036854775807:
                                            var v77 bool = (18446744073709551615 == 2)
                                            _ = v77
                                        case 32767:
                                            v78 := v6[-1]
                                            _ = v78
                                        }
                                        var v79 uint16 = 32767
                                        _ = v79
                                    }
                                    v4 = (^(v4))
                                    if false {
                                        _ = len(v5)
                                        v4 = 182
                                    }
                                }
                                _ = v73
                            }
                            var v80 int16 = -32768
                            for fi81 := 0; fi81 < 6; fi81++ {
                                if (-814 >= -2147483648) {
                                    var v82 float32 = (((3.14159 + ((0.1 * 0.1) + 1e10)) / 1e10) + (1e10 / (1e10 + 0.1)))
                                    var v83 bool = (-1 != -1)
                                    var v84 float64 = 1.5
                                    _ = v82
                                    _ = v83
                                    _ = v84
                                } else {
                                    v80 += -297
                                    v4 += v4
                                    switch fi81 {
                                    case -830:
                                        if ((v80 - 1) != v80) {
                                            switch (("foo bar" + "hello") + "foo bar") {
                                            case "\n\t":
                                                if (v80 > 0) {
                                                    v2 += v2
                                                    fmt.Println(fi81, v2)
                                                }
                                            case "":
                                                switch "hello" {
                                                case "":
                                                    var v85 int = -1
                                                    _ = v85
                                                case "hello":
                                                    if (42 < 1) {
                                                        var v86 uint16 = uint16(v4)
                                                        fmt.Println(fi11, v80, v80)
                                                        _ = v86
                                                    } else {
                                                        var v87 uint8 = (v4 * (v4 - v4))
                                                        switch fi11 {
                                                        case 127:
                                                            if (false || false) {
                                                                v4 = (^(v4))
                                                                var v88 uint8 = (v87 ^ 46)
                                                                if (4294967295 != v2) {
                                                                    fmt.Println(v2, fi81)
                                                                    var v89 string = "a"
                                                                    var v90 float32 = float32(float64((-2.25 / ((1e10 * ((-2.25 - -2.25) * float32(1e10))) + ((1.5 - float32((1e10 + float64(1e10)))) + 0.1)))))
                                                                    _ = v89
                                                                    _ = v90
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v80 = (v80 >> ((65535) & 63))
                                                                    default:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                }
                                                                _ = v88
                                                            } else {
                                                                var v91 string = "\n\t"
                                                                v5 = append(v5, 65535)
                                                                _ = v91
                                                            }
                                                        case 32768:
                                                            v87 += (v87 * v4)
                                                        }
                                                        v2 = (v2 % (((v2 / ((v2) | 1))) | 1))
                                                        _ = v87
                                                    }
                                                }
                                            default:
                                                if true {
                                                    if false {
                                                        switch -9223372036854775808 {
                                                        case -346:
                                                            var v92 float64 = -2.25
                                                            _ = v92
                                                        default:
                                                            switch -9223372036854775808 {
                                                            case 0:
                                                                v93 := GenS9{F0: "", F1: (v2 + 2), F2: 889}
                                                                _ = v93
                                                            default:
                                                                fmt.Println(v80)
                                                                var v94 float64 = 0.1
                                                                var v95 int16 = (-(v80))
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if false {
                                                                var v96 int16 = v80
                                                                _ = v96
                                                            }
                                                        default:
                                                            v80 -= v80
                                                        }
                                                    } else {
                                                        v97 := GenS10{F0: -2147483648}
                                                        v98 := v1[v2]
                                                        var v99 uint8 = (v4 - (v4 / ((54) | 1)))
                                                        switch ("\n\t" + (("" + (v98 + "foo bar")) + ((v98 + (v98 + v98)) + v98))) {
                                                        case "a":
                                                            v5 = append(v5, v2)
                                                        case "hello":
                                                            switch fi81 {
                                                            case -2147483648:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            case 1:
                                                                v4 = v99
                                                                if false {
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (v2 ^ (^(v2)))
                                                                    }
                                                                    fmt.Println(fi11, v99, v99)
                                                                } else {
                                                                    if (1 >= uint64(v4)) {
                                                                        switch 127 {
                                                                        case -9223372036854775808:
                                                                            if false {
                                                                                var v100 string = ((v98 + (v98 + ("\n\t" + (("\n\t" + ("foo bar" + "")) + v98)))) + v98)
                                                                                if (!(612 < uint64(v2))) {
                                                                                    var v101 int = 2147483647
                                                                                    _ = v101
                                                                                } else {
                                                                                    if false {
                                                                                        v2 -= v2
                                                                                        switch (fi81 * -675) {
                                                                                        case 9223372036854775807:
                                                                                            v100 = (("a" + v100) + v100)
                                                                                        }
                                                                                    }
                                                                                    switch (fi81 * 1) {
                                                                                    case 401:
                                                                                        var v102 int64 = 838
                                                                                        _ = v102
                                                                                    case 4294967295:
                                                                                        v99 = (v4 + (v4 % ((v4) | 1)))
                                                                                        switch v98 {
                                                                                        case "foo bar":
                                                                                            v98 = "hello"
                                                                                        case "a":
                                                                                            var v103 uint = 926
                                                                                            _ = v103
                                                                                        }
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    default:
                                                                                        v100 = ("\n\t" + v98)
                                                                                    }
                                                                                    v104 := GenS8{F0: (^(v4)), F1: (v2 - (v2 << ((0) & 63))), F2: (v4 % ((v99) | 1))}
                                                                                    v105 := GenS11{F0: 127, F1: true, F2: (!true)}
                                                                                    _ = v104
                                                                                    _ = v105
                                                                                }
                                                                                v80 = v80
                                                                                _ = v100
                                                                            } else {
                                                                                if ("hello" != v98) {
                                                                                    v106 := GenS3{F0: false}
                                                                                    if (v99 <= v4) {
                                                                                        if (!(0.1 < 3.14159)) {
                                                                                            switch v98 {
                                                                                            case "hello":
                                                                                                v98 = v98
                                                                                            default:
                                                                                                var v107 uint = 255
                                                                                                _ = v107
                                                                                            }
                                                                                            v98 = "hello"
                                                                                        } else {
                                                                                            v108 := GenS8{F0: v4, F1: v2, F2: 196}
                                                                                            _ = v108
                                                                                        }
                                                                                        _ = len(v5)
                                                                                        v109 := GenS6{F0: 42, F1: true, F2: (^(v4))}
                                                                                        _ = v109
                                                                                    } else {
                                                                                        delete(v3, v98)
                                                                                        v80 = v80
                                                                                        v110 := GenS11{F0: 42, F1: true, F2: false}
                                                                                        _ = v110
                                                                                    }
                                                                                    v2 = 1
                                                                                    switch (v98 + "\n\t") {
                                                                                    case "foo bar":
                                                                                        switch fi11 {
                                                                                        case 1:
                                                                                            fmt.Println(fi81, fi81, fi11)
                                                                                        case 128:
                                                                                            var v111 float64 = 0.0
                                                                                            _ = v111
                                                                                        case 32767:
                                                                                            var v112 uint8 = 127
                                                                                            _ = v112
                                                                                        }
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi11)
                                                                                    default:
                                                                                        var v113 float64 = 0.0
                                                                                        _ = v113
                                                                                    }
                                                                                    var v114 string = v98
                                                                                    _ = v106
                                                                                    _ = v114
                                                                                }
                                                                            }
                                                                        case 128:
                                                                            switch (fi11 % (((^(fi11))) | 1)) {
                                                                            case -2147483648:
                                                                                if ((^(v2)) == (v2 ^ v2)) {
                                                                                    v115 := GenS8{F0: 128, F1: (^(v2)), F2: v4}
                                                                                    v5 = append(v5, v2)
                                                                                    _ = v115
                                                                                } else {
                                                                                    fmt.Println(fi81)
                                                                                    var v116 string = "\n\t"
                                                                                    _ = v116
                                                                                }
                                                                            case 1:
                                                                                v99 += v99
                                                                            case 252:
                                                                                v5 = append(v5, v2)
                                                                            default:
                                                                                if false {
                                                                                    v117 := GenS9{F0: "a", F1: 2147483647, F2: 2147483647}
                                                                                    if true {
                                                                                        v98 = v98
                                                                                        if false {
                                                                                            var v118 uint32 = v2
                                                                                            delete(v1, v2)
                                                                                            var v119 uint8 = 255
                                                                                            _ = v118
                                                                                            _ = v119
                                                                                        } else {
                                                                                            switch v98 {
                                                                                            case "a":
                                                                                                switch (v98 + v98) {
                                                                                                case "a":
                                                                                                    v4 += (^(v99))
                                                                                                case "\n\t":
                                                                                                    switch v98 {
                                                                                                    case "\n\t":
                                                                                                        v120 := GenS5{F0: (v80 / (((^(v80))) | 1))}
                                                                                                        _ = v120
                                                                                                    case "hello":
                                                                                                        if (!(true && true)) {
                                                                                                            if (v99 == (v99 + v99)) {
                                                                                                                fmt.Println(v80)
                                                                                                                var v121 int8 = int8(v4)
                                                                                                                _ = v121
                                                                                                            } else {
                                                                                                                var v122 int64 = 9223372036854775807
                                                                                                                var v123 int = 9223372036854775807
                                                                                                                if len(v5) > 0 {
                                                                                                                    v5[0] = v2
                                                                                                                }
                                                                                                                _ = v122
                                                                                                                _ = v123
                                                                                                            }
                                                                                                            var v124 uint8 = (v4 % (((v4 >> ((56) & 63))) | 1))
                                                                                                            _ = v124
                                                                                                        } else {
                                                                                                            var v125 int64 = 0
                                                                                                            v126 := GenS11{F0: -1, F1: (1.5 < float64(float32(0.1))), F2: ((true || false) && true)}
                                                                                                            var v127 uint16 = 32767
                                                                                                            var v128 bool = false
                                                                                                            _ = v125
                                                                                                            _ = v126
                                                                                                            _ = v127
                                                                                                            _ = v128
                                                                                                        }
                                                                                                    default:
                                                                                                        if (true && (!(true && true))) {
                                                                                                            delete(v1, (^(v2)))
                                                                                                        } else {
                                                                                                            fmt.Println(v4, v80)
                                                                                                            fmt.Println(v4, fi81, fi11)
                                                                                                        }
                                                                                                    }
                                                                                                case "hello":
                                                                                                    if true {
                                                                                                        v129 := GenS12{F0: 767}
                                                                                                        var v130 uint = 997
                                                                                                        var v131 string = "a"
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = 350
                                                                                                        }
                                                                                                        var v132 int = 32768
                                                                                                        var v133 uint8 = v4
                                                                                                        _ = v129
                                                                                                        _ = v130
                                                                                                        _ = v131
                                                                                                        _ = v132
                                                                                                        _ = v133
                                                                                                    }
                                                                                                default:
                                                                                                    v134 := GenS8{F0: uint8(v2), F1: v2, F2: v4}
                                                                                                    _ = v134
                                                                                                }
                                                                                                fmt.Println(fi11)
                                                                                                if true {
                                                                                                    if true {
                                                                                                        var v135 uint32 = v2
                                                                                                        var v136 uint = 733
                                                                                                        _ = v135
                                                                                                        _ = v136
                                                                                                    }
                                                                                                    var v137 float32 = float32((float64(1e10) - 0.0))
                                                                                                    fmt.Println(v137)
                                                                                                    v138 := GenS9{F0: ("" + v98), F1: v2, F2: 42}
                                                                                                    fmt.Println(v80)
                                                                                                    if true {
                                                                                                        v139 := GenS4{F0: (105 == 100), F1: fi81, F2: true}
                                                                                                        var v140 float64 = 1e10
                                                                                                        fmt.Println(v4, v2, v99)
                                                                                                        _ = v139
                                                                                                        _ = v140
                                                                                                    } else {
                                                                                                        var v141 int64 = 42
                                                                                                        _ = v141
                                                                                                    }
                                                                                                    _ = v137
                                                                                                    _ = v138
                                                                                                } else {
                                                                                                    switch "hello" {
                                                                                                    case "a":
                                                                                                        var v142 float32 = -2.25
                                                                                                        _ = v142
                                                                                                    }
                                                                                                    v143 := GenS1{F0: false}
                                                                                                    _ = v143
                                                                                                }
                                                                                                if (((!true) && false) && false) {
                                                                                                    var v144 float64 = (-2.25 + 1.5)
                                                                                                    if ((fi81 - fi11) > 1) {
                                                                                                        if (true && ((^(fi81)) != (fi11 / ((fi81) | 1)))) {
                                                                                                            fmt.Println(v2, v98, v2)
                                                                                                            v145 := GenS3{F0: false}
                                                                                                            _ = v145
                                                                                                        }
                                                                                                    }
                                                                                                    v146 := GenS13{F0: 900, F1: "foo bar", F2: 4294967295}
                                                                                                    _ = v144
                                                                                                    _ = v146
                                                                                                }
                                                                                                var v147 int8 = 1
                                                                                                _ = v147
                                                                                            }
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            v5 = append(v5, 984)
                                                                                        case "a":
                                                                                            var v148 int16 = -505
                                                                                            _ = v148
                                                                                        case "hello":
                                                                                            var v149 string = (v98 + ((v98 + (v98 + "")) + ("hello" + v98)))
                                                                                            _ = v149
                                                                                        default:
                                                                                            if true {
                                                                                                if (!(true && ((fi11 ^ (^(fi11))) > fi11))) {
                                                                                                    switch v98 {
                                                                                                    case "a":
                                                                                                        var v150 int32 = 1
                                                                                                        _ = v150
                                                                                                    case "":
                                                                                                        if (447 <= -1) {
                                                                                                            v98 = "\n\t"
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        switch fi81 {
                                                                                                        case 857:
                                                                                                            var v151 int = (fi81 | (^(fi81)))
                                                                                                            _ = v151
                                                                                                        case 441:
                                                                                                            v2 = (^(v2))
                                                                                                        default:
                                                                                                            fmt.Println(fi11, v99)
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    if ((2 <= 255) || false) {
                                                                                                        v80 = (v80 | (v80 % ((v80) | 1)))
                                                                                                        var v152 uint32 = (^(v2))
                                                                                                        _ = len(v6)
                                                                                                        v153 := GenS1{F0: (fi11 >= 1)}
                                                                                                        switch (v98 + v98) {
                                                                                                        case "\n\t":
                                                                                                            var v154 uint32 = v2
                                                                                                            _ = v154
                                                                                                        }
                                                                                                        _ = v152
                                                                                                        _ = v153
                                                                                                    }
                                                                                                }
                                                                                                v155 := GenS2{F0: 65535, F1: 1, F2: -941}
                                                                                                var v156 int8 = -128
                                                                                                _ = v155
                                                                                                _ = v156
                                                                                            } else {
                                                                                                var v157 int32 = 42
                                                                                                _ = v157
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        if true {
                                                                                            v5 = append(v5, 1)
                                                                                            switch (fi81 - fi11) {
                                                                                            case 42:
                                                                                                v4 = (v4 >> ((18446744073709551615) & 63))
                                                                                            case 26:
                                                                                                var v158 int8 = 0
                                                                                                _ = v158
                                                                                            case -858:
                                                                                                fmt.Println(v4, fi81, fi81)
                                                                                            }
                                                                                            var v159 bool = (!(((1.5 + 1e10) + (float32(-2.25) * 3.14159)) <= ((0.1 / 0.1) / ((1e10 - 3.14159) - 3.14159))))
                                                                                            _ = v159
                                                                                        }
                                                                                        v80 -= -233
                                                                                        var v160 uint16 = 265
                                                                                        v161 := GenS2{F0: (v160 - 65535), F1: 1, F2: -160}
                                                                                        _ = v160
                                                                                        _ = v161
                                                                                    } else {
                                                                                        fmt.Println(v2)
                                                                                        switch (fi81 >> ((18446744073709551615) & 63)) {
                                                                                        case 128:
                                                                                            v5 = append(v5, (v2 | v2))
                                                                                        case -9223372036854775808:
                                                                                            v162 := GenS8{F0: v4, F1: 2147483647, F2: 0}
                                                                                            _ = v162
                                                                                        default:
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            v99 -= v4
                                                                                            var v163 int32 = 352
                                                                                            v163 = v163
                                                                                            _ = v163
                                                                                        } else {
                                                                                            var v164 uint = 65535
                                                                                            v99 = v99
                                                                                            _ = v164
                                                                                        }
                                                                                    }
                                                                                    _ = v117
                                                                                } else {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (v2 << ((255) & 63))
                                                                                    }
                                                                                    switch v98 {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v98, v4, v4)
                                                                                    default:
                                                                                        if (-76 >= 42) {
                                                                                            v4 = uint8(v2)
                                                                                            switch -9223372036854775808 {
                                                                                            case -161:
                                                                                                v4 = (v99 | 42)
                                                                                            case 2147483647:
                                                                                                var v165 uint = 255
                                                                                                _ = v165
                                                                                            }
                                                                                            v166 := GenS14{F0: 0, F1: true}
                                                                                            _ = v166
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        case 829:
                                                                            v98 = v98
                                                                        }
                                                                    }
                                                                    v4 += 255
                                                                    var v167 bool = (v98 < v98)
                                                                    _ = v167
                                                                }
                                                                var v168 string = v98
                                                                _ = v168
                                                            default:
                                                                v2 = v2
                                                                if false {
                                                                    if true {
                                                                        v80 = (v80 * -32768)
                                                                        switch v98 {
                                                                        case "":
                                                                            v80 = -453
                                                                        case "\n\t":
                                                                            v169 := GenS13{F0: 0, F1: v98, F2: 1}
                                                                            _ = v169
                                                                        case "foo bar":
                                                                            if ((fi11 > 128) && true) {
                                                                                v170 := GenS10{F0: 1}
                                                                                _ = v170
                                                                            } else {
                                                                                v98 = v98
                                                                                fmt.Println(v2, fi11, v2)
                                                                                if false {
                                                                                    fmt.Println(v80, v4)
                                                                                    if false {
                                                                                        switch (v98 + v98) {
                                                                                        case "\n\t":
                                                                                            v5 = append(v5, 65535)
                                                                                        }
                                                                                        switch v98 {
                                                                                        case "\n\t":
                                                                                            _ = len(v1)
                                                                                        }
                                                                                        switch (fi81 - fi11) {
                                                                                        case -416:
                                                                                            if true {
                                                                                                v80 = 32767
                                                                                                v99 -= (v99 + 0)
                                                                                                fmt.Println(v99)
                                                                                                if false {
                                                                                                    v171 := GenS15{F0: uint(v2), F1: 100}
                                                                                                    v99 = 255
                                                                                                    _ = v171
                                                                                                } else {
                                                                                                    if (v80 == (-(v80))) {
                                                                                                        v172 := GenS16{F0: 2, F1: fi81}
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            v173 := GenS17{F0: (v2 + 2), F1: false}
                                                                                                            if false {
                                                                                                                var v174 bool = false
                                                                                                                _ = v174
                                                                                                            } else {
                                                                                                                switch (-(fi11)) {
                                                                                                                case 32767:
                                                                                                                    v175 := GenS7{F0: false, F1: (v98 + ((v98 + ((v98 + "hello") + "a")) + v98)), F2: 42}
                                                                                                                    _ = v175
                                                                                                                case 42:
                                                                                                                    v80 -= v80
                                                                                                                }
                                                                                                                var v176 float32 = (0.1 + (-2.25 / 1.0))
                                                                                                                _ = v176
                                                                                                            }
                                                                                                            _ = v173
                                                                                                        default:
                                                                                                            v177 := GenS12{F0: (fi81 / ((-2147483648) | 1))}
                                                                                                            _ = v177
                                                                                                        }
                                                                                                        _ = v172
                                                                                                    } else {
                                                                                                        v2 = (v2 << ((255) & 63))
                                                                                                        var v178 int32 = -300
                                                                                                        if ((v178 ^ 204) != -1) {
                                                                                                            _ = len(v5)
                                                                                                            v99 = uint8(v2)
                                                                                                        }
                                                                                                        _ = v178
                                                                                                    }
                                                                                                    var v179 string = v98
                                                                                                    var v180 int16 = 402
                                                                                                    v4 += v4
                                                                                                    _ = v179
                                                                                                    _ = v180
                                                                                                }
                                                                                            }
                                                                                        case 130:
                                                                                            v5 = append(v5, (v2 / ((v2) | 1)))
                                                                                        case -914:
                                                                                            v98 = ((v98 + v98) + ("foo bar" + "foo bar"))
                                                                                        default:
                                                                                            if true {
                                                                                                fmt.Println(v4, v80)
                                                                                                v2 += 50
                                                                                                var v181 int = (fi81 % (((fi81 & fi11)) | 1))
                                                                                                _ = v181
                                                                                            } else {
                                                                                                fmt.Println(v4, fi81, fi81)
                                                                                                v182 := GenS18{F0: (^(v4)), F1: v99}
                                                                                                var v183 float64 = -2.25
                                                                                                _ = v182
                                                                                                _ = v183
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v184 int = (fi11 >> ((2) & 63))
                                                                        _ = v184
                                                                    } else {
                                                                        if ((v80 / ((v80) | 1)) >= v80) {
                                                                            var v185 uint8 = v99
                                                                            if (!(0.1 > 0.0)) {
                                                                                switch (v98 + (v98 + (v98 + "foo bar"))) {
                                                                                case "":
                                                                                    var v186 uint64 = 2
                                                                                    _ = v186
                                                                                case "foo bar":
                                                                                    v99 = (v185 ^ (v4 ^ (^(v99))))
                                                                                case "a":
                                                                                    v3["foo bar"] = v98
                                                                                default:
                                                                                    _ = len(v5)
                                                                                }
                                                                                v80 += (v80 >> ((2) & 63))
                                                                                v80 += v80
                                                                                var v187 int = 997
                                                                                v188 := GenS19{F0: -1, F1: 100, F2: 32767}
                                                                                _ = v187
                                                                                _ = v188
                                                                            } else {
                                                                                var v189 int64 = -1
                                                                                _ = v189
                                                                            }
                                                                            _ = v185
                                                                        } else {
                                                                            var v190 uint16 = 32767
                                                                            v191 := GenS12{F0: (fi81 - (fi81 / ((866) | 1)))}
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            _ = v190
                                                                            _ = v191
                                                                        }
                                                                        v192 := GenS20{F0: 413, F1: (0.0 >= 0.0), F2: v2}
                                                                        v4 += 0
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        _ = v192
                                                                    }
                                                                    var v193 int64 = 1
                                                                    _ = v193
                                                                }
                                                                if len(v5) > 0 {
                                                                    v5[0] = v2
                                                                }
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            v194 := GenS10{F0: -1}
                                                            _ = v194
                                                        }
                                                        v195 := GenS21{F0: 115}
                                                        _ = v195
                                                        _ = v97
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                    if true {
                                                        v80 += 1
                                                    }
                                                }
                                            }
                                        } else {
                                            var v196 bool = true
                                            v197 := GenS22{F0: 4294967295}
                                            _ = v196
                                            _ = v197
                                        }
                                    case 32767:
                                        var v198 int16 = 0
                                        _ = v198
                                    default:
                                        v199 := v6[(fi11 ^ fi81)]
                                        _ = v199
                                    }
                                }
                                fmt.Println(fi81, v80, fi11)
                                fmt.Println(fi11, fi81)
                                var v200 bool = true
                                if true {
                                    if ((!true) || (true || (!((1.5 - float32(1e10)) != (3.14159 / float32(1.5)))))) {
                                        var v201 int32 = 0
                                        v5 = append(v5, (v2 & (v2 ^ (^(v2)))))
                                        v4 = v4
                                        _ = v201
                                    }
                                    v80 -= v80
                                    if false {
                                        switch fi81 {
                                        case 172:
                                            v200 = true
                                        }
                                        v200 = false
                                    } else {
                                        v4 += 165
                                        _ = len(v5)
                                    }
                                }
                                fmt.Println(v200, v4)
                                _ = fi81
                                _ = v200
                            }
                            _ = v80
                        }
                    }
                }
                _ = fi11
            }
        case "\n\t":
            for fi202 := 0; fi202 < 10; fi202++ {
                _ = len(v1)
                _ = fi202
            }
        }
    }
    v203 := GenS23{F0: true, F1: (1.5 != 0.1)}
    switch 0 {
    case 112:
        v204 := GenS24{F0: 446}
        _ = v204
    case 32768:
        v4 -= v4
    }
    if false {
        var v205 int16 = 0
        switch "\n\t" {
        case "":
            fmt.Println(v205)
        case "a":
            var v206 uint8 = (v4 * (v4 / (((v4 ^ 2)) | 1)))
            _ = v206
        }
        var v207 uint16 = 1
        _ = v205
        _ = v207
    }
    if true {
        switch "hello" {
        case "\n\t":
            if true {
                v4 = (^(v4))
            }
        case "hello":
            v4 -= 0
        case "a":
            v4 = (^(v4))
        }
        var v208 uint16 = uint16(v2)
        _ = v208
    }
    v209 := GenS19{F0: 42, F1: 18446744073709551615, F2: 9223372036854775807}
    v4 = v4
    var v210 bool = ((v2 - (^(v2))) < (^(v2)))
    var v211 int16 = -32768
    fmt.Println(v211, v2, v210)
    for fi212 := 0; fi212 < 13; fi212++ {
        if v210 {
            v2 = 2147483648
            v4 = v4
        }
        var v213 uint = 2
        if (!v210) {
            if ((v2 ^ (v2 << ((v213) & 63))) >= 65535) {
                v213 = (v213 * (v213 / ((v213) | 1)))
            }
            if len(v5) > 0 {
                v5[0] = 1
            }
        }
        _ = fi212
        _ = v213
    }
    switch "\n\t" {
    case "a":
        if v210 {
            v4 = (v4 * (^(v4)))
        } else {
            if false {
                if (!((!false) && v210)) {
                    v3[(("\n\t" + "a") + (((((("" + "\n\t") + ("" + ("" + ""))) + "a") + "foo bar") + "foo bar") + ("hello" + "\n\t")))] = "foo bar"
                    var v214 int8 = 59
                    switch 239 {
                    case 2147483647:
                        switch "hello" {
                        case "foo bar":
                            if (((v4 * v4) <= v4) && v210) {
                                if len(v5) > 0 {
                                    v5[0] = 2147483647
                                }
                                fmt.Println(v2)
                            }
                        }
                    }
                    _ = v214
                } else {
                    switch "hello" {
                    case "foo bar":
                        var v215 uint32 = v2
                        _ = v215
                    case "":
                        var v216 int = -9223372036854775808
                        _ = v216
                    case "hello":
                        v211 = v211
                    default:
                        var v217 uint64 = 4294967295
                        _ = v217
                    }
                    var v218 uint8 = v4
                    var v219 int8 = 116
                    _ = v218
                    _ = v219
                }
                switch 42 {
                case 2147483647:
                    if ((v210 && false) && v210) {
                        v4 = v4
                        v210 = v210
                    }
                case 603:
                    if (v211 >= 0) {
                        fmt.Println(v2, v2, v4)
                    } else {
                        switch 795 {
                        case -764:
                            switch 4294967295 {
                            case 32768:
                                var v220 int32 = 0
                                _ = v220
                            }
                        }
                        for fi221 := 0; fi221 < 1; fi221++ {
                            v222 := GenS9{F0: "a", F1: (v2 - (v2 + (v2 ^ v2))), F2: int32(v2)}
                            _ = len(v5)
                            switch fi221 {
                            case 890:
                                v223 := GenS1{F0: (255 < 255)}
                                _ = v223
                            case -242:
                                var v224 int32 = 87
                                _ = v224
                            }
                            _ = fi221
                            _ = v222
                        }
                    }
                case 32768:
                    var v225 int32 = -789
                    _ = v225
                default:
                    v210 = false
                }
                v226 := v6[32768]
                _ = v226
            } else {
                v210 = v210
                v227 := GenS9{F0: "foo bar", F1: 1, F2: 1}
                var v228 uint8 = (v4 & 18)
                _ = v227
                _ = v228
            }
            _ = len(v5)
        }
    case "hello":
        var v229 uint32 = v2
        _ = v229
    default:
        fmt.Println(v2, v210)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v210:", v210)
    fmt.Println("v211:", v211)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v203
    _ = v209
    _ = v210
    _ = v211
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
