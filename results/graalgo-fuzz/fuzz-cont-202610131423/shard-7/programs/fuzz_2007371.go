package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 bool; F1 int8; F2 uint16 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int16; F1 bool; F2 int32 }
type GenS7 struct { F0 int32; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 uint; F1 int8 }

func main() {
    v1 := []uint32{1}
    var v2 float64 = 0.1
    var v3 int = -9223372036854775808
    v4 := []int{(v3 << ((790) & 63)), v3}
    var v5 uint64 = 1
    v6 := GenS1{F0: 9223372036854775807, F1: 34}
    var v7 int64 = -342
    v8 := GenS1{F0: v7, F1: 0}
    v9 := GenS2{F0: (!true), F1: 1}
    var v10 uint16 = 570
    fmt.Println(v2)
    v3 = 127
    if false {
        switch "" {
        case "\n\t":
            if true {
                if ((v10 / ((v10) | 1)) < v10) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        if false {
                            if true {
                                if (false || true) {
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                    v10 -= v10
                                    switch ("\n\t" + (("" + "foo bar") + "hello")) {
                                    case "":
                                        v11 := GenS2{F0: (18446744073709551615 >= 42), F1: 100}
                                        _ = v11
                                    case "hello":
                                        v12 := GenS3{F0: "", F1: "\n\t"}
                                        _ = v12
                                    case "foo bar":
                                        var v13 uint32 = 42
                                        _ = v13
                                    }
                                } else {
                                    v14 := GenS2{F0: (0.0 > 3.14159), F1: 1}
                                    _ = v14
                                }
                                fmt.Println(v2, v10)
                                for fi15 := 0; fi15 < 5; fi15++ {
                                    if ((-2.25 - v2) < v2) {
                                        for fi16 := 0; fi16 < 16; fi16++ {
                                            if true {
                                                v17 := GenS3{F0: "hello", F1: "foo bar"}
                                                v10 -= (v10 << ((342) & 63))
                                                var v18 uint8 = 146
                                                _ = v17
                                                _ = v18
                                            } else {
                                                var v19 string = (("" + "\n\t") + "")
                                                _ = v19
                                            }
                                            _ = fi16
                                        }
                                        switch 32768 {
                                        case 260:
                                            for fi20 := 0; fi20 < 11; fi20++ {
                                                v21 := GenS2{F0: false, F1: -26}
                                                _ = fi20
                                                _ = v21
                                            }
                                        case 4294967295:
                                            v7 = v7
                                        case 42:
                                            fmt.Println(v7, v5, v7)
                                        default:
                                            var v22 float64 = 1e10
                                            _ = v22
                                        }
                                        switch 4294967295 {
                                        case -2147483648:
                                            if false {
                                                var v23 int = -9223372036854775808
                                                if (1 == 42) {
                                                    for fi24 := 0; fi24 < 18; fi24++ {
                                                        if (1 >= 749) {
                                                            v25 := GenS4{F0: (65535 <= 704), F1: -128, F2: (v10 + (v10 % (((^(v10))) | 1)))}
                                                            v4 = append(v4, v23)
                                                            v3 += fi15
                                                            if false {
                                                                v7 += -1
                                                            }
                                                            _ = v25
                                                        } else {
                                                            var v26 string = "foo bar"
                                                            _ = v26
                                                        }
                                                        _ = fi24
                                                    }
                                                    v10 = 751
                                                }
                                                var v27 int32 = 0
                                                _ = v23
                                                _ = v27
                                            } else {
                                                switch int(v7) {
                                                case 765:
                                                    if (9 < 0) {
                                                        fmt.Println(v7, fi15)
                                                        for fi28 := 0; fi28 < 6; fi28++ {
                                                            if (4294967295 <= v5) {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                            _ = fi28
                                                        }
                                                        for fi29 := 0; fi29 < 4; fi29++ {
                                                            if false {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 1
                                                                }
                                                                if false {
                                                                    var v30 int = -9223372036854775808
                                                                    _ = v30
                                                                }
                                                            } else {
                                                                var v31 int16 = 1000
                                                                v4 = append(v4, (-(fi15)))
                                                                v32 := GenS5{F0: true, F1: (("foo bar" + "\n\t") + "a")}
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            _ = fi29
                                                        }
                                                    } else {
                                                        fmt.Println(v5, v5, v5)
                                                        if (!(false && false)) {
                                                            var v33 uint64 = v5
                                                            v33 += (v5 / ((v33) | 1))
                                                            _ = v33
                                                        }
                                                    }
                                                case -266:
                                                    var v34 int64 = 1
                                                    _ = v34
                                                case 127:
                                                    if true {
                                                        v3 -= fi15
                                                    } else {
                                                        var v35 uint8 = 128
                                                        _ = v35
                                                    }
                                                default:
                                                    v2 = v2
                                                }
                                                for fi36 := 0; fi36 < 9; fi36++ {
                                                    if (42 > 4294967295) {
                                                        v10 -= 336
                                                        v37 := GenS6{F0: -183, F1: false, F2: 1}
                                                        var v38 uint16 = v10
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    _ = fi36
                                                }
                                                v10 -= v10
                                            }
                                        }
                                    }
                                    v3 -= (v3 << ((65535) & 63))
                                    v39 := GenS5{F0: (!(((v10 + v10) > v10) || true)), F1: "foo bar"}
                                    _ = fi15
                                    _ = v39
                                }
                            } else {
                                _ = len(v1)
                            }
                            var v40 int32 = -1
                            v41 := GenS6{F0: -1, F1: (v7 >= (v7 / ((-623) | 1))), F2: 0}
                            _ = v40
                            _ = v41
                        }
                    } else {
                        var v42 int8 = -1
                        switch (v3 << ((65535) & 63)) {
                        case 0:
                            v3 += (v3 | v3)
                        case -34:
                            for fi43 := 0; fi43 < 8; fi43++ {
                                v5 -= (^(v5))
                                for fi44 := 0; fi44 < 2; fi44++ {
                                    v45 := GenS3{F0: ("hello" + ""), F1: ("" + ("" + ("\n\t" + "\n\t")))}
                                    var v46 float64 = (3.14159 - v2)
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                v5 = 42
                                _ = fi43
                            }
                        }
                        _ = v42
                    }
                }
                fmt.Println(v3, v7, v3)
                switch (("a" + "a") + (("" + "a") + (("a" + "foo bar") + ("foo bar" + ("hello" + ("hello" + (("a" + "a") + "a"))))))) {
                case "a":
                    v4 = append(v4, (v3 * v3))
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    var v47 float32 = 0.0
                    _ = v47
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = 163
                }
                v10 -= (v10 / ((201) | 1))
                if (1e10 < -2.25) {
                    var v48 uint32 = 2147483648
                    v2 = ((-2.25 / v2) * ((-2.25 / 1e10) + -2.25))
                    _ = v48
                } else {
                    if (v10 != (v10 | v10)) {
                        v10 += v10
                        for fi49 := 0; fi49 < 17; fi49++ {
                            if (v2 < float64(1.5)) {
                                var v50 int8 = 0
                                fmt.Println(v7, v5)
                                v51 := GenS7{F0: 1, F1: 1, F2: 132}
                                _ = v50
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 5; fi52++ {
                                switch (-(v3)) {
                                case 4294967295:
                                    fmt.Println(v3)
                                }
                                if false {
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = fi49
                                        }
                                        if false {
                                            var v53 uint32 = 1
                                            switch 128 {
                                            case 819:
                                                if (626 < 0) {
                                                    switch "a" {
                                                    case "a":
                                                        v54 := GenS4{F0: true, F1: 62, F2: v10}
                                                        _ = v54
                                                    case "\n\t":
                                                        _ = len(v4)
                                                    default:
                                                        switch (fi52 >> ((1) & 63)) {
                                                        case 9223372036854775807:
                                                            if len(v4) > 0 {
                                                                v4[0] = v3
                                                            }
                                                        case 161:
                                                            var v55 int32 = 229
                                                            _ = v55
                                                        case 1:
                                                            v3 = 32767
                                                        default:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    }
                                                    var v56 uint64 = v5
                                                    _ = v56
                                                } else {
                                                    if len(v4) > 0 {
                                                        v4[0] = 9223372036854775807
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v57 uint16 = v10
                                                    if (99 == (^(v5))) {
                                                        v10 = 65535
                                                        if (!false) {
                                                            v58 := GenS8{F0: (v57 << ((18446744073709551615) & 63))}
                                                            var v59 int = -9223372036854775808
                                                            v10 += (v57 >> ((2) & 63))
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            var v60 uint = 255
                                                            var v61 uint32 = (^(v53))
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                    } else {
                                                        if true {
                                                            var v62 int32 = -1
                                                            v62 += (-(v62))
                                                            v53 += v53
                                                            v3 -= fi49
                                                            _ = v62
                                                        }
                                                        if true {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        fmt.Println(v7, v57)
                                                    }
                                                    if ((v5 - (v5 % (((v5 + v5)) | 1))) < uint64(v53)) {
                                                        fmt.Println(v3)
                                                    } else {
                                                        var v63 int = fi52
                                                        if false {
                                                            var v64 uint8 = 128
                                                            v57 = uint16(v64)
                                                            _ = len(v1)
                                                            _ = v64
                                                        }
                                                        _ = v63
                                                    }
                                                    switch "hello" {
                                                    case "":
                                                        var v65 string = "\n\t"
                                                        _ = v65
                                                    default:
                                                        var v66 bool = (!false)
                                                        _ = v66
                                                    }
                                                    _ = v57
                                                }
                                            default:
                                                if false {
                                                    v53 += v53
                                                    switch int(v10) {
                                                    case 0:
                                                        var v67 int16 = -1
                                                        _ = v67
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                v2 -= v2
                                                fmt.Println(v53)
                                            }
                                            if false {
                                                fmt.Println(v10, v7, v2)
                                                _ = len(v4)
                                            }
                                            switch "" {
                                            case "hello":
                                                var v68 int16 = -1
                                                _ = v68
                                            case "foo bar":
                                                var v69 bool = (!(fi52 == fi52))
                                                _ = v69
                                            case "a":
                                                if (true || (32767 >= 32767)) {
                                                    if false {
                                                        var v70 uint8 = 255
                                                        v70 += uint8(v10)
                                                        _ = v70
                                                    } else {
                                                        switch "a" {
                                                        case "hello":
                                                            switch fi52 {
                                                            case 0:
                                                                v3 -= (-(fi49))
                                                                var v71 uint = uint(v10)
                                                                _ = v71
                                                            case 1:
                                                                var v72 int = (fi52 ^ 127)
                                                                _ = v72
                                                            case 128:
                                                                v73 := GenS8{F0: (v10 >> ((648) & 63))}
                                                                _ = v73
                                                            default:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        case "\n\t":
                                                            _ = len(v4)
                                                        }
                                                        switch ("foo bar" + "a") {
                                                        case "foo bar":
                                                            if false {
                                                                if false {
                                                                    var v74 int8 = 15
                                                                    v4 = append(v4, int(v10))
                                                                    fmt.Println(v10, v53, v5)
                                                                    _ = v74
                                                                } else {
                                                                    var v75 uint = 42
                                                                    var v76 bool = (1.5 >= 3.14159)
                                                                    v7 = 0
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                                switch ("foo bar" + "") {
                                                                case "\n\t":
                                                                    fmt.Println(v3, v5)
                                                                default:
                                                                    if (2 >= 1) {
                                                                        v77 := GenS5{F0: ((true || (0 < 1)) || (v7 != (v7 ^ -101))), F1: "hello"}
                                                                        _ = v77
                                                                    } else {
                                                                        _ = len(v1)
                                                                        fmt.Println(v53, v2)
                                                                    }
                                                                }
                                                            } else {
                                                                v3 = v3
                                                                v5 = 1
                                                                if false {
                                                                    var v78 uint = 4294967295
                                                                    _ = v78
                                                                }
                                                            }
                                                        case "a":
                                                            var v79 int64 = (-(v7))
                                                            _ = v79
                                                        case "hello":
                                                            var v80 string = "\n\t"
                                                            _ = v80
                                                        }
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = uint32(v10)
                                                    }
                                                    fmt.Println(v53, v3, v53)
                                                } else {
                                                    var v81 int8 = 1
                                                    switch "\n\t" {
                                                    case "":
                                                        v3 += (fi52 / ((957) | 1))
                                                    }
                                                    _ = v81
                                                }
                                            default:
                                                v53 += v53
                                            }
                                            v82 := GenS4{F0: (true && false), F1: 0, F2: v10}
                                            if len(v1) > 0 {
                                                v1[0] = (v53 / ((v53) | 1))
                                            }
                                            _ = v53
                                            _ = v82
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = 65535
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if true {
                                                if (4294967295 <= 996) {
                                                    v5 = (^(v5))
                                                    v10 = (v10 + v10)
                                                    fmt.Println(fi49, v7, fi52)
                                                }
                                                v83 := GenS9{F0: 65535}
                                                v10 = (v10 >> ((0) & 63))
                                                var v84 uint8 = 1
                                                var v85 uint64 = v5
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                        var v86 int16 = -1
                                        var v87 uint64 = 2
                                        var v88 float32 = -2.25
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                }
                                _ = fi52
                            }
                            _ = fi49
                        }
                        if false {
                            if false {
                                for fi89 := 0; fi89 < 4; fi89++ {
                                    v7 -= v7
                                    v3 += 32768
                                    v3 -= v3
                                    _ = fi89
                                }
                            } else {
                                for fi90 := 0; fi90 < 1; fi90++ {
                                    switch (v3 << ((61) & 63)) {
                                    case 128:
                                        var v91 uint8 = 105
                                        _ = v91
                                    case 553:
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                    default:
                                        v5 = v5
                                    }
                                    _ = fi90
                                }
                                v2 += (1e10 / ((v2 - v2) / (((1e10 / v2) + 3.14159) - 1e10)))
                                if ((1.5 / float32(v2)) > 0.1) {
                                    var v92 int = v3
                                    var v93 bool = (float64((3.14159 + (0.0 + (3.14159 / (0.0 * (float32(v2) / float32(float64((float32((v2 / (v2 + (3.14159 * (1.5 - (v2 + 0.0)))))) * (3.14159 + -2.25)))))))))) <= v2)
                                    if (!false) {
                                        fmt.Println(v93, v3, v92)
                                        fmt.Println(v2, v92)
                                        v5 -= 4294967295
                                    }
                                    _ = v92
                                    _ = v93
                                }
                            }
                        } else {
                            v1 = append(v1, 2147483647)
                        }
                    } else {
                        switch ("a" + "a") {
                        case "hello":
                            var v94 int16 = 559
                            _ = v94
                        case "foo bar":
                            v10 = v10
                        case "\n\t":
                            if (!true) {
                                var v95 bool = ((1e10 - (0.0 + 0.1)) != -2.25)
                                _ = v95
                            }
                        default:
                            v96 := GenS10{F0: 2, F1: 42}
                            _ = v96
                        }
                    }
                    var v97 uint32 = 1
                    fmt.Println(v7, v2)
                    _ = v97
                }
            }
        case "":
            var v98 float32 = (0.1 - (float32(v2) / (0.1 + -2.25)))
            _ = v98
        default:
            v3 = v3
        }
        fmt.Println(v5)
        v99 := GenS11{F0: 42, F1: 71}
        _ = v99
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
