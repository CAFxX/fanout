package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int8; F2 string }
type GenS5 struct { F0 int8; F1 int8; F2 uint }
type GenS6 struct { F0 string; F1 uint; F2 int8 }
type GenS7 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 string; F1 uint64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint8; F1 bool }
type GenS13 struct { F0 int16; F1 int64; F2 int }
type GenS14 struct { F0 uint32; F1 int }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int8; F1 int16 }

func main() {
    var v1 int16 = 42
    var v2 int8 = 0
    var v3 uint32 = 2
    v4 := []uint32{865}
    var v5 uint = 0
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: -474, F1: -9223372036854775808, F2: -1}
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int32 = -2147483648
                v10 := GenS1{F0: 0, F1: -9223372036854775808, F2: (-(v2))}
                var v11 int8 = (^(v2))
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS1{F0: 9223372036854775807, F1: -1, F2: -124}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) > 0 {
                v4[0] = v3
            }
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS2{F0: (v1 / (((v1 ^ v1)) | 1)), F1: 2147483647, F2: 42}
    for fi15 := 0; fi15 < 17; fi15++ {
        if true {
            v5 += v5
            v3 = v3
            v16 := GenS3{F0: v2}
            _ = v16
        } else {
            var v17 int16 = 0
            switch (fi15 - fi15) {
            case 758:
                switch "a" {
                case "foo bar":
                    var v18 int32 = 2147483647
                    _ = v18
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        if true {
                            fmt.Println(fi15, v3)
                        } else {
                            var v19 uint = 65535
                            var v20 int32 = -284
                            v20 = (v20 ^ v20)
                            _ = v19
                            _ = v20
                        }
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = v3
                        }
                    case "a":
                        v2 += v2
                    default:
                        for fi21 := 0; fi21 < 11; fi21++ {
                            fmt.Println(v5, v17)
                            v3 += v3
                            _ = fi21
                        }
                    }
                case "a":
                    switch (fi15 * fi15) {
                    case 823:
                        v1 = (v17 + v17)
                    default:
                        v4 = append(v4, 203)
                    }
                default:
                    if true {
                        var v22 string = "foo bar"
                        fmt.Println(v1)
                        v5 = v5
                        _ = v22
                    }
                }
            case 42:
                if (v5 > v5) {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        fmt.Println(v1, v3)
                        v2 = -124
                        if false {
                            v5 = (v5 ^ 2)
                            v24 := GenS4{F0: (-(v2)), F1: (^(v2)), F2: "foo bar"}
                            var v25 float32 = ((1.5 * 1e10) - 0.0)
                            _ = v24
                            _ = v25
                        } else {
                            var v26 int32 = 2147483647
                            var v27 int = 42
                            _ = v26
                            _ = v27
                        }
                        _ = fi23
                    }
                    if false {
                        v2 = v2
                    }
                } else {
                    v28 := GenS5{F0: int8(v5), F1: (v2 ^ (-(v2))), F2: 4294967295}
                    v29 := GenS6{F0: "", F1: (^(v5)), F2: int8(v5)}
                    _ = v28
                    _ = v29
                }
            case 9223372036854775807:
                fmt.Println(v1, v17, v17)
            }
            _ = v17
        }
        var v30 uint8 = 128
        var v31 int = (fi15 >> (((v5 << ((619) & 63))) & 63))
        _ = fi15
        _ = v30
        _ = v31
    }
    switch ("a" + "foo bar") {
    case "\n\t":
        if true {
            switch -852 {
            case -9223372036854775808:
                var v32 bool = true
                _ = v32
            case -398:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 108:
                fmt.Println(v1, v1, v5)
            default:
                switch 0 {
                case -2147483648:
                    v3 = (v3 - 4294967295)
                case 0:
                    if (!true) {
                        switch "a" {
                        case "foo bar":
                            v33 := GenS5{F0: (v2 | (v2 ^ -1)), F1: (v2 + -42), F2: 954}
                            _ = v33
                        default:
                            fmt.Println(v1, v2)
                        }
                        if false {
                            v5 = 2
                        }
                    } else {
                        if (1.5 != 1e10) {
                            var v34 uint8 = 0
                            v3 += v3
                            fmt.Println(v5, v34)
                            _ = v34
                        } else {
                            switch int(v2) {
                            case 402:
                                v3 = v3
                            case 42:
                                fmt.Println(v3, v1)
                            case 32768:
                                if len(v4) > 0 {
                                    v4[0] = (v3 >> (((v5 & 42)) & 63))
                                }
                            default:
                                for fi35 := 0; fi35 < 1; fi35++ {
                                    _ = len(v4)
                                    _ = fi35
                                }
                            }
                            var v36 uint = (v5 >> ((v5) & 63))
                            _ = v36
                        }
                        switch 32767 {
                        case -9223372036854775808:
                            if false {
                                switch -1 {
                                case 9223372036854775807:
                                    for fi37 := 0; fi37 < 4; fi37++ {
                                        switch ("foo bar" + (("\n\t" + "") + "a")) {
                                        case "\n\t":
                                            v38 := GenS7{F0: 2, F1: 989, F2: v3}
                                            _ = v38
                                        case "hello":
                                            var v39 bool = ((179 <= (^(v3))) || (1e10 == -2.25))
                                            _ = v39
                                        case "":
                                            if (!false) {
                                                v4 = append(v4, v3)
                                            }
                                        default:
                                            switch ((((("\n\t" + "") + "a") + ("\n\t" + "")) + "hello") + ("\n\t" + "foo bar")) {
                                            case "\n\t":
                                                v3 -= 65535
                                            }
                                        }
                                        var v40 uint8 = 1
                                        _ = fi37
                                        _ = v40
                                    }
                                case -2147483648:
                                    v4 = append(v4, (^(v3)))
                                default:
                                    for fi41 := 0; fi41 < 13; fi41++ {
                                        _ = len(v4)
                                        if true {
                                            switch (fi41 + fi41) {
                                            case 42:
                                                switch (fi41 | fi41) {
                                                case 32768:
                                                    fmt.Println(v2, v5)
                                                case 2147483647:
                                                    fmt.Println(v5)
                                                default:
                                                    var v42 int64 = 1
                                                    _ = v42
                                                }
                                            case 168:
                                                var v43 float32 = 0.1
                                                _ = v43
                                            case 128:
                                                v5 = 18446744073709551615
                                            }
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = 42
                                            }
                                            for fi44 := 0; fi44 < 18; fi44++ {
                                                var v45 uint = 917
                                                if false {
                                                    switch (fi41 + 983) {
                                                    case 127:
                                                        switch -806 {
                                                        case 127:
                                                            v46 := GenS8{F0: v45}
                                                            _ = v46
                                                        }
                                                    case -2147483648:
                                                        v3 = v3
                                                    case -231:
                                                        var v47 uint32 = (v3 - (v3 * (v3 + (v3 | v3))))
                                                        _ = v47
                                                    default:
                                                        v48 := GenS9{F0: (v45 ^ (^(v5))), F1: (v2 - (v2 * (v2 + (v2 - v2))))}
                                                        _ = v48
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v49 int64 = 9223372036854775807
                                                    _ = v49
                                                }
                                                _ = fi44
                                                _ = v45
                                            }
                                        }
                                        for fi50 := 0; fi50 < 15; fi50++ {
                                            var v51 uint = (v5 ^ 938)
                                            v52 := GenS10{F0: "foo bar", F1: 703}
                                            v1 = -1
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = fi41
                                    }
                                }
                            } else {
                                if true {
                                    for fi53 := 0; fi53 < 6; fi53++ {
                                        if false {
                                            if (!(!false)) {
                                                v3 = v3
                                                var v54 string = "a"
                                                v54 = "\n\t"
                                                _ = v54
                                            }
                                            switch int(v5) {
                                            case 2147483647:
                                                for fi55 := 0; fi55 < 9; fi55++ {
                                                    var v56 int64 = -504
                                                    fmt.Println(v5, v2, v56)
                                                    var v57 uint = v5
                                                    _ = fi55
                                                    _ = v56
                                                    _ = v57
                                                }
                                            case -288:
                                                for fi58 := 0; fi58 < 3; fi58++ {
                                                    switch "hello" {
                                                    case "":
                                                        switch "a" {
                                                        case "a":
                                                            var v59 int = (^(fi58))
                                                            _ = v59
                                                        }
                                                    case "foo bar":
                                                        v60 := GenS9{F0: v5, F1: (v2 << (((v5 / (((v5 / ((v5) | 1))) | 1))) & 63))}
                                                        _ = v60
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    var v61 string = "foo bar"
                                                    _ = fi58
                                                    _ = v61
                                                }
                                            case 127:
                                                if true {
                                                    fmt.Println(v5)
                                                    _ = len(v4)
                                                    for fi62 := 0; fi62 < 3; fi62++ {
                                                        var v63 float64 = (1e10 + float64(0.0))
                                                        var v64 int8 = (v2 << ((v5) & 63))
                                                        v65 := GenS11{F0: v2}
                                                        _ = fi62
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                }
                                            }
                                            _ = len(v4)
                                        } else {
                                            for fi66 := 0; fi66 < 10; fi66++ {
                                                var v67 uint32 = 759
                                                var v68 int32 = 42
                                                _ = fi66
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                        _ = fi53
                                    }
                                } else {
                                    v1 = (v1 + 988)
                                    v4 = append(v4, (v3 / ((v3) | 1)))
                                    if (true && (!false)) {
                                        if false {
                                            v69 := GenS8{F0: uint(v3)}
                                            _ = v69
                                        }
                                        v5 -= 0
                                    } else {
                                        v70 := GenS7{F0: 255, F1: 17, F2: (^(v3))}
                                        _ = v70
                                    }
                                }
                                fmt.Println(v1, v1)
                                if false {
                                    var v71 uint8 = 255
                                    if len(v4) > 0 {
                                        v4[0] = (v3 / ((v3) | 1))
                                    }
                                    _ = v71
                                } else {
                                    switch 659 {
                                    case -875:
                                        fmt.Println(v3, v3, v3)
                                    case -486:
                                        v5 -= 1
                                    }
                                    if (743 <= 4294967295) {
                                        v1 = (v1 << ((0) & 63))
                                        v72 := GenS12{F0: 0, F1: true}
                                        if (("\n\t" + "hello") != "foo bar") {
                                            var v73 int8 = v2
                                            if true {
                                                v74 := GenS13{F0: v1, F1: -9223372036854775808, F2: -792}
                                                var v75 int16 = v1
                                                _ = v74
                                                _ = v75
                                            }
                                            _ = v73
                                        }
                                        _ = v72
                                    } else {
                                        var v76 uint32 = (v3 / (((v3 << (((^(v5))) & 63))) | 1))
                                        switch int(v5) {
                                        case -9223372036854775808:
                                            if (!(18446744073709551615 >= 18446744073709551615)) {
                                                _ = len(v4)
                                            }
                                        case 361:
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(v3)
                                            case "hello":
                                                v76 = v76
                                            }
                                        default:
                                            if true {
                                                var v77 uint16 = 65535
                                                _ = v77
                                            }
                                        }
                                        _ = v76
                                    }
                                    for fi78 := 0; fi78 < 16; fi78++ {
                                        var v79 int = -976
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi78
                                        _ = v79
                                    }
                                }
                            }
                        case 127:
                            var v80 string = (("a" + "hello") + "")
                            _ = v80
                        case 599:
                            v81 := GenS7{F0: 127, F1: -1, F2: (^(v3))}
                            _ = v81
                        }
                        fmt.Println(v1)
                    }
                case 354:
                    var v82 uint16 = 1
                    _ = v82
                default:
                    var v83 uint32 = 2
                    _ = v83
                }
            }
            fmt.Println(v5, v3, v1)
        } else {
            var v84 int64 = -1
            var v85 string = ("hello" + "foo bar")
            _ = v84
            _ = v85
        }
    default:
        for fi86 := 0; fi86 < 10; fi86++ {
            var v87 uint = (^(v5))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi86
            _ = v87
        }
    }
    var v88 float32 = (-2.25 / 0.1)
    var v89 int = 944
    var v90 string = "a"
    fmt.Println(v88)
    var v91 float64 = float64(v88)
    v2 += int8(v89)
    if false {
        var v92 string = "hello"
        _ = v92
    } else {
        if true {
            v90 = v90
            var v93 int64 = 9223372036854775807
            for fi94 := 0; fi94 < 10; fi94++ {
                var v95 int = int(v93)
                _ = fi94
                _ = v95
            }
            _ = v93
        }
    }
    if (173 >= 1) {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    } else {
        for fi96 := 0; fi96 < 12; fi96++ {
            v3 = v3
            _ = fi96
        }
        for fi97 := 0; fi97 < 1; fi97++ {
            var v98 int16 = (v1 % ((v1) | 1))
            if ((v2 <= (-(v2))) && (v88 <= 0.1)) {
                var v99 uint16 = 752
                v100 := GenS4{F0: v2, F1: (v2 << ((v5) & 63)), F2: v90}
                _ = v100
                _ = v99
            } else {
                v2 += v2
            }
            if (1.5 != v91) {
                v101 := GenS14{F0: v3, F1: (fi97 - fi97)}
                var v102 int32 = -2147483648
                _ = v101
                _ = v102
            } else {
                _ = len(v4)
                if false {
                    v103 := GenS15{F0: (v2 / ((v2) | 1))}
                    _ = v103
                } else {
                    fmt.Println(v90)
                }
                _ = len(v4)
            }
            _ = fi97
            _ = v98
        }
        v4 = append(v4, 2147483647)
    }
    fmt.Println(v89, v91)
    if true {
        for fi104 := 0; fi104 < 7; fi104++ {
            v105 := GenS16{F0: (v2 - (v2 | v2)), F1: (v1 - (v1 >> (((v5 / ((v5) | 1))) & 63)))}
            fmt.Println(v2, v5, v2)
            v3 += (v3 * v3)
            _ = fi104
            _ = v105
        }
        fmt.Println(v2)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
