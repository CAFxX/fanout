package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint8; F1 uint16; F2 string }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint32; F1 int8 }
type GenS16 struct { F0 uint; F1 uint64 }
type GenS17 struct { F0 uint16; F1 uint16; F2 bool }
type GenS18 struct { F0 int16; F1 uint8; F2 string }
type GenS19 struct { F0 string; F1 int8 }

func main() {
    var v1 float32 = (((3.14159 / 3.14159) * (1e10 - 3.14159)) / (0.0 - (((1e10 - float32(1.5)) / -2.25) + (0.1 * 1e10))))
    var v2 bool = (true || true)
    var v3 bool = v2
    v4 := []uint32{4294967295, 1, 722}
    var v5 bool = (548 >= -504)
    var v6 float64 = ((float64((v1 / float32(1e10))) * 1.5) - 0.0)
    v3 = false
    v2 = (1 >= 42)
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (fi7 % ((fi7) | 1)) {
        case 14:
            var v8 bool = (!false)
            _ = v8
        case 1:
            v2 = v5
        default:
            v9 := GenS1{F0: -1, F1: 583, F2: 489}
            _ = v9
        }
        if v2 {
            v3 = v3
            fmt.Println(v1, v3, v5)
            var v10 int8 = 1
            _ = v10
        }
        switch "hello" {
        case "\n\t":
            switch "a" {
            case "foo bar":
                v6 = (v6 / v6)
            case "":
                if ((!(-904 < 9223372036854775807)) || v3) {
                    switch "a" {
                    case "a":
                        v5 = (963 == 2)
                    case "hello":
                        switch 42 {
                        case 42:
                            if (127 > 1) {
                                if false {
                                    for fi11 := 0; fi11 < 20; fi11++ {
                                        var v12 int16 = 1
                                        _ = fi11
                                        _ = v12
                                    }
                                    v6 = v6
                                } else {
                                    v13 := GenS2{F0: 127}
                                    if v3 {
                                        switch (^(fi7)) {
                                        case -9223372036854775808:
                                            fmt.Println(v1)
                                        case -246:
                                            for fi14 := 0; fi14 < 10; fi14++ {
                                                var v15 bool = v2
                                                var v16 uint = 42
                                                v17 := GenS3{F0: 42}
                                                _ = fi14
                                                _ = v15
                                                _ = v16
                                                _ = v17
                                            }
                                        case -511:
                                            v18 := GenS4{F0: 562, F1: fi7}
                                            _ = v18
                                        default:
                                            switch fi7 {
                                            case -9223372036854775808:
                                                for fi19 := 0; fi19 < 19; fi19++ {
                                                    if v5 {
                                                        if (206 == 14) {
                                                            switch (fi19 % ((675) | 1)) {
                                                            case -86:
                                                                var v20 int16 = 1000
                                                                _ = v20
                                                            case 4294967295:
                                                                switch fi7 {
                                                                case 32767:
                                                                    v2 = false
                                                                case 855:
                                                                    var v21 float64 = (v6 + (0.0 / 1.5))
                                                                    _ = v21
                                                                case 4294967295:
                                                                    var v22 float32 = (v1 - float32(1.5))
                                                                    _ = v22
                                                                }
                                                            default:
                                                                v3 = (true || (((!v5) || (42 <= 65535)) && true))
                                                            }
                                                            _ = len(v4)
                                                        }
                                                        if (-1 == 453) {
                                                            v6 -= 0.1
                                                            if (!(v5 && v5)) {
                                                                fmt.Println(fi7, v2)
                                                            } else {
                                                                fmt.Println(fi19, v6)
                                                            }
                                                            switch fi19 {
                                                            case -288:
                                                                if (!(!false)) {
                                                                    if (1 == 1) {
                                                                        v23 := GenS2{F0: 42}
                                                                        switch (-(fi7)) {
                                                                        case 966:
                                                                            v5 = true
                                                                        case 32768:
                                                                            if v2 {
                                                                                if v5 {
                                                                                    var v24 int16 = -1
                                                                                    _ = v24
                                                                                }
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    fmt.Println(fi19, fi19)
                                                                                case "hello":
                                                                                    v3 = v3
                                                                                case "":
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 217
                                                                                    }
                                                                                }
                                                                                v1 = float32((((((v6 * (((v6 + float64(v1)) - (((-2.25 - (v6 + (-2.25 * v6))) - (1.5 * (float64(float32(3.14159)) + 3.14159))) * v6)) + 1e10)) * 0.1) / float64(1e10)) / float64((v1 / float32(v6)))) * 0.0))
                                                                                v3 = false
                                                                            }
                                                                        default:
                                                                            var v25 int16 = 775
                                                                            _ = v25
                                                                        }
                                                                        var v26 float32 = 3.14159
                                                                        _ = v23
                                                                        _ = v26
                                                                    }
                                                                } else {
                                                                    switch (("hello" + "hello") + "foo bar") {
                                                                    case "hello":
                                                                        fmt.Println(v6)
                                                                    case "":
                                                                        switch ("\n\t" + "hello") {
                                                                        case "":
                                                                            var v27 int8 = 65
                                                                            _ = v27
                                                                        default:
                                                                            v28 := GenS3{F0: 127}
                                                                            _ = v28
                                                                        }
                                                                    case "foo bar":
                                                                        fmt.Println(fi7, fi19, v3)
                                                                    }
                                                                }
                                                            case 0:
                                                                fmt.Println(fi7, v2, fi19)
                                                            case 128:
                                                                v29 := GenS1{F0: -216, F1: 65535, F2: 4294967295}
                                                                _ = v29
                                                            }
                                                            v6 = float64(float32(((((3.14159 * v6) * (((v6 * v6) + v6) / float64((0.1 * float32(float64(v1)))))) - (v6 - v6)) / v6)))
                                                        } else {
                                                            var v30 uint16 = 42
                                                            _ = v30
                                                        }
                                                        var v31 uint64 = 0
                                                        v6 = 1.5
                                                        _ = v31
                                                    }
                                                    var v32 int = (fi19 * fi7)
                                                    if true {
                                                        fmt.Println(v2, v32)
                                                        _ = len(v4)
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = 2147483647
                                                    }
                                                    v32 -= 32767
                                                    _ = fi19
                                                    _ = v32
                                                }
                                            case 128:
                                                v5 = true
                                            case 4294967295:
                                                var v33 int = fi7
                                                _ = v33
                                            default:
                                                var v34 bool = true
                                                _ = v34
                                            }
                                        }
                                        v5 = (v3 || false)
                                    }
                                    switch 0 {
                                    case 9223372036854775807:
                                        v6 -= v6
                                    case 128:
                                        _ = len(v4)
                                    case -336:
                                        if len(v4) > 0 {
                                            v4[0] = 697
                                        }
                                    default:
                                        var v35 int8 = 1
                                        _ = v35
                                    }
                                    _ = v13
                                }
                                v4 = append(v4, 2147483648)
                            }
                        case -1:
                            v36 := GenS5{F0: fi7, F1: 1}
                            _ = v36
                        default:
                            if ((v3 && v5) && v2) {
                                switch fi7 {
                                case 2147483647:
                                    var v37 int32 = -718
                                    _ = v37
                                case 9223372036854775807:
                                    fmt.Println(v1, fi7)
                                case 127:
                                    _ = len(v4)
                                default:
                                    v4 = append(v4, 340)
                                }
                                var v38 float32 = (v1 + v1)
                                _ = v38
                            }
                        }
                    default:
                        if false {
                            v4 = append(v4, 1)
                            var v39 float32 = v1
                            switch (^(fi7)) {
                            case 788:
                                v5 = v3
                            }
                            _ = v39
                        } else {
                            if v2 {
                                for fi40 := 0; fi40 < 15; fi40++ {
                                    if (-32768 > 32767) {
                                        v3 = (v5 || v2)
                                        if (-32768 > -32768) {
                                            switch (fi40 - 128) {
                                            case 2147483647:
                                                v41 := GenS2{F0: -105}
                                                _ = v41
                                            default:
                                                if v2 {
                                                    switch (fi7 << ((4294967295) & 63)) {
                                                    case 32768:
                                                        if v5 {
                                                            v42 := GenS4{F0: 255, F1: (fi7 << ((461) & 63))}
                                                            v5 = (66 < 0)
                                                            switch "foo bar" {
                                                            case "":
                                                                if len(v4) > 0 {
                                                                    v4[0] = 276
                                                                }
                                                            case "a":
                                                                if (!((9223372036854775807 >= 475) && true)) {
                                                                    v1 -= v1
                                                                    var v43 int = (fi40 << ((510) & 63))
                                                                    _ = v43
                                                                } else {
                                                                    v5 = ((fi7 & 4294967295) <= (fi7 + (fi7 + (-(fi40)))))
                                                                    v44 := GenS5{F0: (fi40 << ((65535) & 63)), F1: 405}
                                                                    _ = v44
                                                                }
                                                            }
                                                            _ = v42
                                                        }
                                                    case 427:
                                                        switch 0 {
                                                        case 42:
                                                            v2 = v2
                                                        case -513:
                                                            fmt.Println(v2, fi40)
                                                        case 9223372036854775807:
                                                            if (996 == 1) {
                                                                var v45 int16 = 0
                                                                var v46 int32 = 2147483647
                                                                v45 = v45
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        }
                                                    case 239:
                                                        v2 = (-1 != -128)
                                                    default:
                                                        var v47 int16 = 32767
                                                        _ = v47
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    v4 = append(v4, 829)
                                                    switch fi40 {
                                                    case 2147483647:
                                                        var v48 int16 = -382
                                                        _ = v48
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = 2147483648
                                                    }
                                                }
                                            }
                                            v6 += v6
                                        }
                                        var v49 string = ("foo bar" + ("a" + (("\n\t" + "a") + "\n\t")))
                                        var v50 bool = true
                                        v1 = 3.14159
                                        switch v49 {
                                        case "":
                                            switch v49 {
                                            case "\n\t":
                                                var v51 uint32 = 0
                                                _ = v51
                                            case "foo bar":
                                                v52 := GenS4{F0: 255, F1: (fi7 >> ((65535) & 63))}
                                                _ = v52
                                            case "hello":
                                                switch (fi7 << ((65535) & 63)) {
                                                case -2147483648:
                                                    var v53 int32 = 2147483647
                                                    _ = v53
                                                default:
                                                    v54 := GenS6{F0: 18446744073709551615, F1: 65535, F2: 2}
                                                    if true {
                                                        if (!((!v5) && true)) {
                                                            var v55 int16 = -32768
                                                            v6 -= 1e10
                                                            _ = v55
                                                        }
                                                    } else {
                                                        if (-1 >= 79) {
                                                            if (18446744073709551615 < 1) {
                                                                v56 := GenS7{F0: 2}
                                                                v4 = append(v4, 1)
                                                                switch v49 {
                                                                case "\n\t":
                                                                    v50 = v2
                                                                case "a":
                                                                    switch fi7 {
                                                                    case 42:
                                                                        fmt.Println(fi40)
                                                                    default:
                                                                        fmt.Println(fi7, fi7)
                                                                    }
                                                                case "foo bar":
                                                                    var v57 float64 = (((float64(v1) - v6) - v6) - ((v6 - v6) * v6))
                                                                    _ = v57
                                                                }
                                                                _ = v56
                                                            }
                                                            v58 := GenS7{F0: 70}
                                                            var v59 uint64 = 0
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            if v5 {
                                                                v5 = v2
                                                                v60 := GenS3{F0: 701}
                                                                _ = v60
                                                            } else {
                                                                var v61 int16 = 0
                                                                switch "a" {
                                                                case "a":
                                                                    v5 = true
                                                                case "\n\t":
                                                                    switch v49 {
                                                                    case "\n\t":
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    default:
                                                                        v2 = ((fi40 ^ fi7) == fi40)
                                                                    }
                                                                }
                                                                _ = v61
                                                            }
                                                        }
                                                        var v62 string = v49
                                                        var v63 float32 = (v1 + float32(1.5))
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    if true {
                                                        if (!(-123 >= 127)) {
                                                            var v64 bool = v50
                                                            _ = v64
                                                        }
                                                    } else {
                                                        if false {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            if true {
                                                                v6 -= 0.0
                                                                fmt.Println(fi40, fi7, v49)
                                                            }
                                                            v3 = v3
                                                            if (!v2) {
                                                                if v3 {
                                                                    fmt.Println(fi40, v2)
                                                                    v65 := GenS8{F0: 42, F1: 65535}
                                                                    _ = v65
                                                                }
                                                                v1 = (v1 + float32((v6 * -2.25)))
                                                            } else {
                                                                fmt.Println(v49)
                                                                var v66 string = "\n\t"
                                                                v1 = float32(0.0)
                                                                _ = v66
                                                            }
                                                            var v67 int32 = 0
                                                            _ = v67
                                                        } else {
                                                            v50 = v5
                                                            switch v49 {
                                                            case "":
                                                                if ((-389 > 409) && v3) {
                                                                    fmt.Println(v3, v50, v2)
                                                                    var v68 int64 = 9223372036854775807
                                                                    _ = v68
                                                                } else {
                                                                    _ = len(v4)
                                                                }
                                                            default:
                                                                v50 = v2
                                                            }
                                                            if v2 {
                                                                fmt.Println(v1, v49)
                                                            }
                                                        }
                                                    }
                                                    if (0 <= 42) {
                                                        var v69 int = fi7
                                                        var v70 int64 = 845
                                                        v50 = v3
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    var v71 int = fi40
                                                    _ = v54
                                                    _ = v71
                                                }
                                            }
                                            if v5 {
                                                if v3 {
                                                    var v72 float64 = -2.25
                                                    v73 := GenS9{F0: 0, F1: 84, F2: "a"}
                                                    v1 = v1
                                                    _ = v72
                                                    _ = v73
                                                }
                                            } else {
                                                switch v49 {
                                                case "foo bar":
                                                    if (-1 == 0) {
                                                        if (false && v5) {
                                                            var v74 uint64 = 2
                                                            if true {
                                                                fmt.Println(v50)
                                                            }
                                                            if v2 {
                                                                if ((v74 % ((v74) | 1)) > 18446744073709551615) {
                                                                    var v75 uint32 = 4294967295
                                                                    if (1 == 250) {
                                                                        var v76 int16 = 42
                                                                        _ = v76
                                                                    }
                                                                    var v77 int8 = 127
                                                                    _ = v75
                                                                    _ = v77
                                                                }
                                                            }
                                                            _ = v74
                                                        } else {
                                                            switch v49 {
                                                            case "":
                                                                _ = len(v4)
                                                            case "a":
                                                                var v78 int16 = -330
                                                                _ = v78
                                                            default:
                                                                v2 = v3
                                                            }
                                                            switch fi40 {
                                                            case 128:
                                                                var v79 uint = 0
                                                                _ = v79
                                                            case 843:
                                                                v4 = append(v4, 2147483648)
                                                            case -9223372036854775808:
                                                                if ((v49 < v49) || v50) {
                                                                    switch "a" {
                                                                    case "":
                                                                        v2 = ((^(fi7)) == 1)
                                                                    case "a":
                                                                        var v80 uint64 = 566
                                                                        _ = v80
                                                                    case "foo bar":
                                                                        if v2 {
                                                                            var v81 string = (v49 + (v49 + v49))
                                                                            v2 = v5
                                                                            v82 := GenS2{F0: 42}
                                                                            v83 := GenS10{F0: -825, F1: 42}
                                                                            switch "a" {
                                                                            case "":
                                                                                v84 := GenS11{F0: 1}
                                                                                var v85 string = (((v81 + ("hello" + (v49 + v49))) + v49) + v49)
                                                                                var v86 float64 = float64(float32((0.0 / v6)))
                                                                                _ = v84
                                                                                _ = v85
                                                                                _ = v86
                                                                            }
                                                                            v3 = (224 >= -2147483648)
                                                                            _ = v81
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                } else {
                                                                    if v2 {
                                                                        v87 := GenS9{F0: 255, F1: 108, F2: ("\n\t" + v49)}
                                                                        _ = v87
                                                                    }
                                                                    fmt.Println(fi40, v49)
                                                                }
                                                            }
                                                            var v88 int64 = 42
                                                            _ = v88
                                                        }
                                                    }
                                                default:
                                                    v50 = true
                                                }
                                                fmt.Println(v50, v1, v50)
                                                if len(v4) > 0 {
                                                    v4[0] = 2147483648
                                                }
                                                if (65535 >= 1) {
                                                    if ((0.0 + ((((0.1 - (float32(v6) / 0.1)) - (v1 - 0.1)) - v1) * (v1 * v1))) > v1) {
                                                        var v89 string = (v49 + (v49 + (v49 + v49)))
                                                        v90 := GenS12{F0: 2, F1: -9223372036854775808}
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                    v5 = false
                                                    v3 = ((-2147483648 == 2147483647) && false)
                                                    v1 = (v1 - v1)
                                                    v91 := GenS11{F0: 0}
                                                    _ = v91
                                                }
                                                var v92 float64 = v6
                                                var v93 float64 = float64(3.14159)
                                                var v94 int16 = 1
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            }
                                        case "hello":
                                            switch v49 {
                                            case "a":
                                                if (42 < 4294967295) {
                                                    if true {
                                                        var v95 uint32 = 458
                                                        switch fi7 {
                                                        case 928:
                                                            fmt.Println(fi7)
                                                        default:
                                                            if (-1 < 42) {
                                                                switch fi7 {
                                                                case 128:
                                                                    fmt.Println(v5)
                                                                case 933:
                                                                    fmt.Println(v3, v6)
                                                                }
                                                            } else {
                                                                v5 = true
                                                            }
                                                        }
                                                        fmt.Println(v49, fi40)
                                                        v96 := GenS13{F0: 65535, F1: 473, F2: 613}
                                                        v5 = ((fi7 / (((^(fi7))) | 1)) <= (^(fi40)))
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    if (132 > 616) {
                                                        var v97 uint64 = 18446744073709551615
                                                        if false {
                                                            switch (fi40 - fi40) {
                                                            case -1:
                                                                v4 = append(v4, 2147483647)
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                            v4 = append(v4, 507)
                                                            var v98 int64 = 0
                                                            v50 = (1 <= 0)
                                                            if len(v4) > 0 {
                                                                v4[0] = 4294967295
                                                            }
                                                            _ = len(v4)
                                                            fmt.Println(v1)
                                                            _ = v98
                                                        } else {
                                                            var v99 int64 = 9223372036854775807
                                                            if len(v4) > 0 {
                                                                v4[0] = 442
                                                            }
                                                            _ = v99
                                                        }
                                                        _ = v97
                                                    }
                                                    var v100 uint16 = 42
                                                    v101 := GenS7{F0: 84}
                                                    _ = v100
                                                    _ = v101
                                                } else {
                                                    switch 2147483647 {
                                                    case -2147483648:
                                                        v102 := GenS4{F0: 1, F1: (fi7 & 855)}
                                                        _ = v102
                                                    case 0:
                                                        _ = len(v4)
                                                    default:
                                                        switch fi40 {
                                                        case 128:
                                                            if (v3 || v3) {
                                                                v103 := GenS14{F0: 0}
                                                                var v104 int32 = 440
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                if false {
                                                                    var v105 uint16 = 0
                                                                    _ = v105
                                                                }
                                                            }
                                                        case 965:
                                                            switch v49 {
                                                            case "":
                                                                if v5 {
                                                                    v5 = false
                                                                    var v106 int64 = 0
                                                                    _ = v106
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = 4294967295
                                                    }
                                                }
                                            }
                                        case "a":
                                            switch (fi7 * (fi40 % (((fi7 & 2147483647)) | 1))) {
                                            case 4294967295:
                                                if v50 {
                                                    if v50 {
                                                        var v107 int8 = -128
                                                        v49 = (v49 + v49)
                                                        _ = v107
                                                    }
                                                }
                                            case -1:
                                                switch (fi7 ^ (^(fi7))) {
                                                case 0:
                                                    v2 = (!(42 >= 2))
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v6 -= float64(float32(float64(v1)))
                                                    var v108 uint8 = 127
                                                    _ = v108
                                                case -346:
                                                    var v109 uint = 824
                                                    _ = v109
                                                case 2147483647:
                                                    v110 := GenS1{F0: -744, F1: 255, F2: 172}
                                                    var v111 int8 = 8
                                                    v2 = v50
                                                    if (v5 && true) {
                                                        if (356 != 18446744073709551615) {
                                                            var v112 int = 127
                                                            _ = v112
                                                        } else {
                                                            switch fi7 {
                                                            case 376:
                                                                fmt.Println(v49, v3)
                                                            }
                                                            v113 := GenS2{F0: v111}
                                                            _ = v113
                                                        }
                                                        var v114 bool = true
                                                        fmt.Println(v49, v5, v114)
                                                        _ = v114
                                                    }
                                                    _ = v110
                                                    _ = v111
                                                default:
                                                    var v115 int = (fi7 - (-(fi40)))
                                                    _ = v115
                                                }
                                            case 32768:
                                                v2 = v3
                                            default:
                                                switch fi7 {
                                                case 9223372036854775807:
                                                    switch (fi7 / (((-(fi40))) | 1)) {
                                                    case 2147483647:
                                                        if v50 {
                                                            if len(v4) > 0 {
                                                                v4[0] = 4294967295
                                                            }
                                                            if ((503 != 42) && true) {
                                                                v116 := GenS4{F0: 42, F1: fi7}
                                                                _ = len(v4)
                                                                fmt.Println(fi7)
                                                                _ = v116
                                                            }
                                                            v117 := GenS3{F0: 18446744073709551615}
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v118 uint16 = 0
                                                            if len(v4) > 0 {
                                                                v4[0] = 979
                                                            }
                                                            var v119 int8 = 42
                                                            _ = v117
                                                            _ = v118
                                                            _ = v119
                                                        }
                                                    }
                                                case 2147483647:
                                                    v6 = 1.5
                                                }
                                            }
                                        default:
                                            fmt.Println(fi7, v49, v3)
                                        }
                                        if ("hello" >= (v49 + ("" + v49))) {
                                            switch ("a" + v49) {
                                            case "foo bar":
                                                if ((42 >= 127) && v3) {
                                                    if (true && (882 < -2147483648)) {
                                                        v120 := GenS10{F0: 1, F1: 961}
                                                        if false {
                                                            var v121 bool = true
                                                            _ = v121
                                                        } else {
                                                            var v122 bool = (0.1 != ((float32(1e10) * (v1 - v1)) + 1.5))
                                                            if v3 {
                                                                if ((fi40 << ((42) & 63)) <= (fi40 & 1)) {
                                                                    switch fi40 {
                                                                    case -1:
                                                                        v49 = "\n\t"
                                                                    case 9223372036854775807:
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    default:
                                                                        if (255 <= 169) {
                                                                            v123 := GenS5{F0: -9223372036854775808, F1: 67}
                                                                            v2 = v122
                                                                            _ = v123
                                                                        }
                                                                    }
                                                                }
                                                                if (-1 != -1) {
                                                                    var v124 int = 2147483647
                                                                    v6 = 1.5
                                                                    fmt.Println(v1)
                                                                    switch ("\n\t" + "hello") {
                                                                    case "hello":
                                                                        if (float32(v6) > v1) {
                                                                            v125 := GenS15{F0: 4294967295, F1: -128}
                                                                            var v126 uint = 805
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                    case "foo bar":
                                                                        fmt.Println(v122, fi7, fi7)
                                                                    case "a":
                                                                        v1 += v1
                                                                        fmt.Println(v2, v124)
                                                                    default:
                                                                        var v127 int8 = -1
                                                                        _ = v127
                                                                    }
                                                                    _ = v124
                                                                }
                                                                v122 = true
                                                            }
                                                            v1 = v1
                                                            _ = v122
                                                        }
                                                        _ = v120
                                                    } else {
                                                        var v128 string = v49
                                                        var v129 int16 = 0
                                                        switch "a" {
                                                        case "\n\t":
                                                            fmt.Println(v128, v5, v129)
                                                        case "a":
                                                            switch (v49 + (("\n\t" + ((v49 + v128) + v49)) + v49)) {
                                                            case "\n\t":
                                                                var v130 int = (fi7 - (fi40 * fi7))
                                                                _ = v130
                                                            }
                                                        default:
                                                            if (v2 || (false && (134 >= 1))) {
                                                                v4 = append(v4, 65535)
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                fmt.Println(v1, v49)
                                                            } else {
                                                                v131 := GenS16{F0: 1, F1: 1}
                                                                v6 += 0.1
                                                                _ = v131
                                                            }
                                                        }
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                }
                                                var v132 bool = (9223372036854775807 > -52)
                                                _ = v132
                                            case "hello":
                                                if true {
                                                    v5 = (595 != -9223372036854775808)
                                                    v133 := GenS7{F0: 1}
                                                    _ = v133
                                                }
                                            case "a":
                                                if len(v4) > 0 {
                                                    v4[0] = 1
                                                }
                                            }
                                            var v134 bool = (!v3)
                                            _ = v134
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = 0
                                            }
                                        }
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v3 = v3
                                        v2 = (!(v5 || (fi7 < (fi40 / ((-9223372036854775808) | 1)))))
                                    }
                                    v135 := GenS17{F0: 1, F1: 1, F2: ((!(1 >= 32767)) || (!(1 >= 255)))}
                                    if (-16 >= 42) {
                                        _ = len(v4)
                                        switch (fi7 / ((fi7) | 1)) {
                                        case 431:
                                            switch "foo bar" {
                                            case "a":
                                                v6 = 1.5
                                            }
                                            var v136 int16 = -32768
                                            v136 += v136
                                            _ = v136
                                        case -392:
                                            v137 := GenS18{F0: 410, F1: 255, F2: "hello"}
                                            _ = v137
                                        }
                                        v138 := GenS9{F0: 42, F1: 255, F2: ("a" + ("hello" + ""))}
                                        _ = v138
                                    } else {
                                        _ = len(v4)
                                    }
                                    _ = fi40
                                    _ = v135
                                }
                                switch fi7 {
                                case -9223372036854775808:
                                    v2 = (42 < 918)
                                }
                            }
                        }
                    }
                    v3 = (236 != 128)
                    switch fi7 {
                    case 495:
                        v139 := GenS19{F0: "hello", F1: 1}
                        _ = v139
                    default:
                        v5 = (v3 && (v2 && true))
                    }
                }
            case "hello":
                var v140 int32 = 1
                _ = v140
            default:
                v141 := GenS10{F0: 225, F1: 4294967295}
                _ = v141
            }
        case "a":
            v142 := GenS18{F0: 0, F1: 1, F2: "\n\t"}
            _ = v142
        }
        _ = fi7
    }
    var v143 int16 = -263
    v144 := GenS8{F0: 0, F1: 1}
    switch "hello" {
    case "\n\t":
        var v145 uint16 = 113
        _ = v145
    case "hello":
        v6 -= float64((0.1 - v1))
    default:
        _ = len(v4)
    }
    var v146 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v143
    _ = v144
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
