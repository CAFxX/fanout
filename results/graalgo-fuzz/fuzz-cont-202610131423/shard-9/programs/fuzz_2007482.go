package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32; F1 int16; F2 int16 }
type GenS11 struct { F0 uint64; F1 uint8 }
type GenS12 struct { F0 uint16; F1 int16 }
type GenS13 struct { F0 int8; F1 int32; F2 string }
type GenS14 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS15 struct { F0 int64; F1 string; F2 int32 }
type GenS16 struct { F0 int; F1 uint8; F2 uint32 }
type GenS17 struct { F0 int32; F1 string; F2 uint }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 int8; F1 uint16 }
type GenS20 struct { F0 string; F1 bool }
type GenS21 struct { F0 bool }
func genFn0(p0 uint64, p1 string) uint32 {
    p1 = "hello"
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 2
    v3 := []int{441, -37, 32767}
    var v4 uint = 13
    var v5 float32 = float32((-2.25 + 3.14159))
    v6 := GenS1{F0: 11, F1: ("" + "foo bar")}
    fmt.Println(v5)
    switch 127 {
    case 32768:
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: genFn0((v2 | (v2 >> ((v4) & 63))), "a"), F1: ""}
        fmt.Println(v1)
        var v10 float32 = (0.0 - v5)
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 4294967295 {
    case 4294967295:
        switch -2147483648 {
        case 1:
            var v11 uint = 65535
            _ = v11
        case 935:
            if (207 >= -2147483648) {
                v12 := GenS2{F0: 436}
                var v13 uint16 = 913
                _ = v12
                _ = v13
            }
        }
    case 769:
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int64 = 792
            fmt.Println(v4, v2)
            v5 = 1.5
            _ = fi14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int16 = 32767
            switch fi16 {
            case 4294967295:
                for fi18 := 0; fi18 < 18; fi18++ {
                    switch "a" {
                    case "hello":
                        switch (fi18 ^ (fi16 << ((v4) & 63))) {
                        case -1:
                            if true {
                                var v19 uint = 2
                                fmt.Println(fi16)
                                var v20 uint16 = (v1 & (v1 - 255))
                                _ = v19
                                _ = v20
                            } else {
                                if true {
                                    if (2 != 249) {
                                        fmt.Println(v17, fi16)
                                        v3 = append(v3, (fi16 + 1))
                                    }
                                    if ((!false) && false) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if true {
                                            switch 0 {
                                            case 935:
                                                fmt.Println(fi18)
                                            case -520:
                                                if len(v3) > 0 {
                                                    v3[0] = fi16
                                                }
                                            case 149:
                                                switch 2147483647 {
                                                case -9223372036854775808:
                                                    var v21 uint16 = (v1 * v1)
                                                    _ = v21
                                                case 678:
                                                    if true {
                                                        v22 := GenS2{F0: 164}
                                                        var v23 int16 = v17
                                                        _ = v22
                                                        _ = v23
                                                    } else {
                                                        switch (-(fi16)) {
                                                        case 2147483647:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v24 := GenS3{F0: (v17 << (((^(v4))) & 63))}
                                                                _ = v24
                                                            case "a":
                                                                switch 1 {
                                                                case 128:
                                                                    v25 := GenS3{F0: int16(v1)}
                                                                    _ = v25
                                                                default:
                                                                    var v26 int8 = -92
                                                                    _ = v26
                                                                }
                                                            default:
                                                                switch "a" {
                                                                case "\n\t":
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v27 int8 = 1
                                                                        _ = v27
                                                                    case "hello":
                                                                        fmt.Println(fi18, v2, v4)
                                                                    case "\n\t":
                                                                        switch (fi16 & fi16) {
                                                                        case 9223372036854775807:
                                                                            v1 += v1
                                                                            var v28 int16 = (v17 / ((v17) | 1))
                                                                            _ = v28
                                                                        case 4294967295:
                                                                            fmt.Println(v4, v17)
                                                                        case -2147483648:
                                                                            var v29 uint64 = (v2 * (^(v2)))
                                                                            _ = v29
                                                                        }
                                                                    }
                                                                case "a":
                                                                    v3 = append(v3, fi16)
                                                                case "hello":
                                                                    if (229 <= 255) {
                                                                        if ("a" != (("\n\t" + "hello") + "\n\t")) {
                                                                            var v30 bool = false
                                                                            _ = v30
                                                                        }
                                                                        if false {
                                                                            var v31 int64 = 42
                                                                            var v32 float32 = (v5 / (((v5 - (v5 * float32((0.0 * 1.5)))) / (float32((3.14159 / ((0.0 - (3.14159 / 3.14159)) - (3.14159 * 3.14159)))) + ((((1.5 + (v5 - (v5 * 1.5))) - v5) * 0.1) / v5))) / v5))
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                        var v33 float32 = 0.0
                                                                        _ = v33
                                                                    } else {
                                                                        if false {
                                                                            if false {
                                                                                v2 = (v2 & v2)
                                                                                v17 = (^(v17))
                                                                                switch (((("a" + "a") + "a") + "hello") + ("foo bar" + "hello")) {
                                                                                case "foo bar":
                                                                                    v3 = append(v3, int(v1))
                                                                                    var v34 uint32 = 42
                                                                                    _ = v34
                                                                                }
                                                                            } else {
                                                                                if (v17 >= v17) {
                                                                                    fmt.Println(v5, v1, v1)
                                                                                    v4 += (v4 | 65535)
                                                                                }
                                                                                v35 := GenS2{F0: 637}
                                                                                v36 := GenS1{F0: 2, F1: "foo bar"}
                                                                                if (!true) {
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            v17 = (-(v17))
                                                                                        default:
                                                                                            var v37 uint64 = v2
                                                                                            _ = v37
                                                                                        }
                                                                                    case "\n\t":
                                                                                        fmt.Println(v2, v17)
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            v2 = v2
                                                                                            var v38 int = (fi16 % (((-(fi18))) | 1))
                                                                                            if (!false) {
                                                                                                if (false && false) {
                                                                                                    switch v38 {
                                                                                                    case -1:
                                                                                                        v4 += (^(v4))
                                                                                                    default:
                                                                                                        var v39 uint16 = (v1 % ((166) | 1))
                                                                                                        _ = v39
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v40 bool = (uint16(v2) >= (v1 * (v1 * 406)))
                                                                                                        if (!(!(0 <= -574))) {
                                                                                                            switch v38 {
                                                                                                            case -1:
                                                                                                                if true {
                                                                                                                    v2 = uint64(v1)
                                                                                                                    switch "foo bar" {
                                                                                                                    case "\n\t":
                                                                                                                        v4 = (^(v4))
                                                                                                                    default:
                                                                                                                        v1 = v1
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch (v38 % ((fi18) | 1)) {
                                                                                                                    case 695:
                                                                                                                        switch "foo bar" {
                                                                                                                        case "foo bar":
                                                                                                                            if len(v3) >= 2 {
                                                                                                                                v3 = v3[1:len(v3)]
                                                                                                                            }
                                                                                                                        case "\n\t":
                                                                                                                            fmt.Println(fi18, v1, v40)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v41 := GenS1{F0: genFn0((v2 * v2), "\n\t"), F1: "foo bar"}
                                                                                                                        _ = v41
                                                                                                                    }
                                                                                                                    fmt.Println(v1, v38)
                                                                                                                }
                                                                                                            case 0:
                                                                                                                if v40 {
                                                                                                                    v38 = v38
                                                                                                                } else {
                                                                                                                    fmt.Println(v40, v17)
                                                                                                                    var v42 float32 = v5
                                                                                                                    if ((v1 | v1) > uint16(v2)) {
                                                                                                                        fmt.Println(v2, v42, fi18)
                                                                                                                        var v43 string = "a"
                                                                                                                        _ = len(v3)
                                                                                                                        _ = v43
                                                                                                                    }
                                                                                                                    _ = v42
                                                                                                                }
                                                                                                            }
                                                                                                            v4 = (v4 / (((v4 | v4)) | 1))
                                                                                                        } else {
                                                                                                            v4 = 111
                                                                                                            fmt.Println(v4, v1, v1)
                                                                                                        }
                                                                                                        v1 += (v1 % ((v1) | 1))
                                                                                                        _ = v40
                                                                                                    } else {
                                                                                                        switch ("\n\t" + ("foo bar" + "foo bar")) {
                                                                                                        case "a":
                                                                                                            v5 -= (((v5 + float32(1e10)) - -2.25) / 0.1)
                                                                                                        default:
                                                                                                            switch fi18 {
                                                                                                            case 128:
                                                                                                                var v44 float64 = ((1e10 + float64(float32(1e10))) / (float64(3.14159) + -2.25))
                                                                                                                _ = v44
                                                                                                            case -2147483648:
                                                                                                                var v45 string = ""
                                                                                                                _ = v45
                                                                                                            case -641:
                                                                                                                v46 := GenS1{F0: 1, F1: "foo bar"}
                                                                                                                _ = v46
                                                                                                            }
                                                                                                        }
                                                                                                        switch fi18 {
                                                                                                        case 9223372036854775807:
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = 127
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            var v47 int = -69
                                                                                                            _ = v47
                                                                                                        case 42:
                                                                                                            fmt.Println(v17, v5, v38)
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = (fi16 / ((v38) | 1))
                                                                                                            }
                                                                                                            var v48 bool = ((!true) || false)
                                                                                                            _ = v48
                                                                                                        }
                                                                                                        v49 := GenS4{F0: 9223372036854775807}
                                                                                                        _ = v49
                                                                                                    }
                                                                                                } else {
                                                                                                    var v50 int16 = 32767
                                                                                                    _ = v50
                                                                                                }
                                                                                                v2 -= 2
                                                                                                fmt.Println(fi16, v1, v38)
                                                                                            } else {
                                                                                                switch ("foo bar" + (("\n\t" + ("foo bar" + "")) + ("hello" + ("a" + "foo bar")))) {
                                                                                                case "\n\t":
                                                                                                    v2 = 42
                                                                                                case "":
                                                                                                    switch "hello" {
                                                                                                    case "foo bar":
                                                                                                        var v51 int8 = 0
                                                                                                        _ = v51
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v17, v2, v1)
                                                                                                    case "hello":
                                                                                                        var v52 uint16 = v1
                                                                                                        if false {
                                                                                                            switch fi18 {
                                                                                                            case -9223372036854775808:
                                                                                                                var v53 uint8 = 128
                                                                                                                _ = v53
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v54 float32 = float32(((float64((1e10 * 1e10)) * 3.14159) * 0.1))
                                                                                                                v55 := GenS4{F0: v38}
                                                                                                                _ = v54
                                                                                                                _ = v55
                                                                                                            } else {
                                                                                                                v5 = ((v5 * v5) - v5)
                                                                                                                v56 := GenS3{F0: (v17 + -1)}
                                                                                                                _ = v56
                                                                                                            }
                                                                                                            v57 := GenS2{F0: 65535}
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            var v58 uint32 = 65535
                                                                                                            fmt.Println(v52, v52, v1)
                                                                                                            fmt.Println(v17, v52, v4)
                                                                                                            _ = v57
                                                                                                            _ = v58
                                                                                                        } else {
                                                                                                            v3 = append(v3, (fi18 & (fi18 & fi18)))
                                                                                                        }
                                                                                                        _ = v52
                                                                                                    default:
                                                                                                        v4 = v4
                                                                                                        v17 += (v17 + (v17 * v17))
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    var v59 int64 = 42
                                                                                                    _ = v59
                                                                                                }
                                                                                                v2 += uint64(v4)
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            var v60 float32 = v5
                                                                                            _ = len(v3)
                                                                                            _ = v38
                                                                                            _ = v60
                                                                                        } else {
                                                                                            _ = len(v3)
                                                                                            if (false || false) {
                                                                                                v61 := GenS5{F0: true}
                                                                                                _ = v61
                                                                                            } else {
                                                                                                v2 = (^(v2))
                                                                                                v2 -= v2
                                                                                                v62 := GenS6{F0: int16(v1), F1: 120}
                                                                                                _ = v62
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v63 uint16 = 65535
                                                                                    if (true && true) {
                                                                                        if true {
                                                                                            if true {
                                                                                                if (42 > -984) {
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = (fi18 % ((0) | 1))
                                                                                                    }
                                                                                                    var v64 uint16 = v63
                                                                                                    if (float64((float32((0.1 * (1.5 + ((0.0 + 3.14159) + (float64(((v5 + (3.14159 + 0.0)) * v5)) + -2.25))))) - (0.0 * 1e10))) > (float64((v5 + v5)) - (0.1 + float64(1.5)))) {
                                                                                                        v63 -= (^(v64))
                                                                                                    }
                                                                                                    _ = v64
                                                                                                }
                                                                                                var v65 uint64 = (v2 << (((^(v4))) & 63))
                                                                                                _ = v65
                                                                                            }
                                                                                        } else {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = -296
                                                                                        }
                                                                                        if (255 == v4) {
                                                                                            switch ("" + ("\n\t" + "")) {
                                                                                            case "foo bar":
                                                                                                var v66 float64 = (3.14159 / 0.1)
                                                                                                _ = v66
                                                                                            default:
                                                                                                var v67 string = "foo bar"
                                                                                                _ = v67
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v68 := GenS7{F0: 87, F1: "\n\t"}
                                                                                        fmt.Println(v5)
                                                                                        if false {
                                                                                            fmt.Println(v1)
                                                                                        } else {
                                                                                            v69 := GenS8{F0: -1, F1: ""}
                                                                                            v4 += (v4 >> (((v4 / (((v4 | uint(v2))) | 1))) & 63))
                                                                                            var v70 uint32 = 418
                                                                                            var v71 string = ""
                                                                                            _ = v69
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                        }
                                                                                        _ = v68
                                                                                    }
                                                                                    _ = v63
                                                                                } else {
                                                                                    var v72 int16 = 849
                                                                                    if true {
                                                                                        var v73 uint64 = 4294967295
                                                                                        _ = len(v3)
                                                                                        v72 -= int16(v2)
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = (fi18 / ((fi18) | 1))
                                                                                        }
                                                                                        _ = v73
                                                                                    } else {
                                                                                        if false {
                                                                                            v74 := GenS8{F0: 52, F1: "a"}
                                                                                            _ = v74
                                                                                        }
                                                                                        fmt.Println(fi18, fi16, v1)
                                                                                        var v75 bool = ((-9223372036854775808 == 0) || false)
                                                                                        _ = v75
                                                                                    }
                                                                                    var v76 int32 = 42
                                                                                    _ = v72
                                                                                    _ = v76
                                                                                }
                                                                                _ = len(v3)
                                                                                _ = v35
                                                                                _ = v36
                                                                            }
                                                                            _ = len(v3)
                                                                            var v77 string = "foo bar"
                                                                            v1 = v1
                                                                            if (true && false) {
                                                                                v1 = v1
                                                                                if (!false) {
                                                                                    var v78 uint = (v4 * v4)
                                                                                    fmt.Println(v77, v77, v4)
                                                                                    if ((^(v2)) >= (^(v2))) {
                                                                                        v79 := GenS2{F0: 4294967295}
                                                                                        fmt.Println(v5, fi18)
                                                                                        _ = v79
                                                                                    } else {
                                                                                        v77 = "hello"
                                                                                        var v80 uint8 = 2
                                                                                        v17 = v17
                                                                                        _ = v80
                                                                                    }
                                                                                    _ = v78
                                                                                }
                                                                            }
                                                                            v81 := GenS3{F0: v17}
                                                                            v4 = 2
                                                                            if false {
                                                                                var v82 float32 = (v5 / v5)
                                                                                var v83 float32 = v5
                                                                                var v84 int32 = 0
                                                                                _ = v82
                                                                                _ = v83
                                                                                _ = v84
                                                                            } else {
                                                                                v77 = v77
                                                                            }
                                                                            var v85 float32 = (v5 / 1.5)
                                                                            v2 = v2
                                                                            if false {
                                                                                v1 += v1
                                                                            } else {
                                                                                if (int16(fi18) > -1) {
                                                                                    var v86 float64 = 1.5
                                                                                    _ = v86
                                                                                }
                                                                                var v87 int32 = -840
                                                                                _ = v87
                                                                            }
                                                                            var v88 uint16 = v1
                                                                            v89 := GenS9{F0: (42 < 42)}
                                                                            var v90 bool = (!((57 == 127) || true))
                                                                            _ = v77
                                                                            _ = v81
                                                                            _ = v85
                                                                            _ = v88
                                                                            _ = v89
                                                                            _ = v90
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                v17 -= 32767
                                                                            case "hello":
                                                                                var v91 uint32 = 42
                                                                                _ = v91
                                                                            case "a":
                                                                                var v92 uint16 = (v1 >> ((65535) & 63))
                                                                                _ = v92
                                                                            default:
                                                                                v93 := GenS10{F0: 237, F1: -13, F2: v17}
                                                                                _ = v93
                                                                            }
                                                                            if (true && (v2 >= uint64(v4))) {
                                                                                v94 := GenS4{F0: int(v4)}
                                                                                fmt.Println(v1)
                                                                                switch (fi16 | 626) {
                                                                                case 1:
                                                                                    v95 := GenS11{F0: v2, F1: 166}
                                                                                    _ = v95
                                                                                default:
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        var v96 bool = true
                                                                                        _ = v96
                                                                                    case "\n\t":
                                                                                        switch (fi18 / ((fi18) | 1)) {
                                                                                        case 42:
                                                                                            var v97 int64 = 42
                                                                                            _ = v97
                                                                                        case 817:
                                                                                            v17 = v17
                                                                                        case -818:
                                                                                            switch ("\n\t" + "foo bar") {
                                                                                            case "a":
                                                                                                if false {
                                                                                                    switch (("a" + "a") + ("\n\t" + ("hello" + "a"))) {
                                                                                                    case "foo bar":
                                                                                                        var v98 uint = (v4 * uint(v1))
                                                                                                        _ = v98
                                                                                                    case "":
                                                                                                        v4 = v4
                                                                                                    }
                                                                                                }
                                                                                            case "\n\t":
                                                                                                var v99 string = ("a" + ("hello" + ("a" + ("\n\t" + ("hello" + "a")))))
                                                                                                _ = v99
                                                                                            case "foo bar":
                                                                                                var v100 int = (fi16 & fi18)
                                                                                                _ = v100
                                                                                            default:
                                                                                                var v101 float64 = (1.5 * ((0.1 / (float64(((3.14159 / ((v5 - (((float32(3.14159) - 1.5) / v5) * 1.5)) + v5)) * (0.0 * (-2.25 / (float32(0.0) - v5))))) / ((float64(1.5) + ((float64((float32(3.14159) + (v5 * ((1e10 + v5) + v5)))) - ((1.5 * 3.14159) + (-2.25 / 1e10))) + 1e10)) - (1e10 - float64(-2.25))))) / 3.14159))
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            v4 = 18446744073709551615
                                                                                            v102 := GenS7{F0: 42, F1: ("\n\t" + "hello")}
                                                                                            _ = v102
                                                                                        case "hello":
                                                                                            var v103 int8 = -1
                                                                                            _ = v103
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = fi18
                                                                                            }
                                                                                            if false {
                                                                                                if (!false) {
                                                                                                    switch (^(fi16)) {
                                                                                                    case 9223372036854775807:
                                                                                                        var v104 uint16 = v1
                                                                                                        _ = v104
                                                                                                    case -2147483648:
                                                                                                        if true {
                                                                                                            v4 += 65535
                                                                                                            var v105 int16 = 32767
                                                                                                            var v106 uint16 = (^(v1))
                                                                                                            switch "\n\t" {
                                                                                                            case "a":
                                                                                                                switch (fi16 << (((v4 >> ((v4) & 63))) & 63)) {
                                                                                                                case -659:
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = fi18
                                                                                                                    }
                                                                                                                case -990:
                                                                                                                    switch "a" {
                                                                                                                    case "a":
                                                                                                                        var v107 string = ""
                                                                                                                        _ = v107
                                                                                                                    case "":
                                                                                                                        v108 := GenS12{F0: v1, F1: int16(v106)}
                                                                                                                        _ = v108
                                                                                                                    case "foo bar":
                                                                                                                        if (0 < (v106 ^ (v106 % ((v1) | 1)))) {
                                                                                                                            v109 := GenS8{F0: -107, F1: "a"}
                                                                                                                            if true {
                                                                                                                                switch (fi18 + int(v1)) {
                                                                                                                                case 32768:
                                                                                                                                    var v110 uint32 = genFn0(v2, "foo bar")
                                                                                                                                    _ = v110
                                                                                                                                default:
                                                                                                                                    if (!false) {
                                                                                                                                        v111 := GenS12{F0: (v106 ^ 32767), F1: -32768}
                                                                                                                                        if false {
                                                                                                                                            if false {
                                                                                                                                                v5 -= (-2.25 / (v5 * (((((1e10 * 1.5) + ((v5 / v5) + float32((1e10 * (0.0 / -2.25))))) + 1.5) * ((v5 / float32(1e10)) - float32(0.1))) - 0.0)))
                                                                                                                                                v4 = 18446744073709551615
                                                                                                                                            } else {
                                                                                                                                                v112 := GenS6{F0: (v17 & (-(v17))), F1: 42}
                                                                                                                                                v113 := GenS13{F0: 1, F1: 42, F2: "a"}
                                                                                                                                                _ = v112
                                                                                                                                                _ = v113
                                                                                                                                            }
                                                                                                                                            v105 = v105
                                                                                                                                            switch (-(fi18)) {
                                                                                                                                            case -1:
                                                                                                                                                v114 := GenS13{F0: 89, F1: 42, F2: ("" + "foo bar")}
                                                                                                                                                _ = v114
                                                                                                                                            default:
                                                                                                                                                if true {
                                                                                                                                                    v2 = v2
                                                                                                                                                    v106 = (v1 ^ 801)
                                                                                                                                                    v115 := GenS14{F0: (v17 * v17), F1: 255, F2: 42}
                                                                                                                                                    _ = v115
                                                                                                                                                } else {
                                                                                                                                                    fmt.Println(v105, v2, v17)
                                                                                                                                                    v116 := GenS12{F0: (^(v1)), F1: (v17 * 32767)}
                                                                                                                                                    _ = v116
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            v4 = (v4 + 65535)
                                                                                                                                        } else {
                                                                                                                                            var v117 int8 = 1
                                                                                                                                            _ = v117
                                                                                                                                        }
                                                                                                                                        _ = v111
                                                                                                                                    } else {
                                                                                                                                        v5 += v5
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v118 string = ""
                                                                                                                            _ = v109
                                                                                                                            _ = v118
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch "\n\t" {
                                                                                                                        case "hello":
                                                                                                                            if true {
                                                                                                                                v119 := GenS11{F0: v2, F1: 0}
                                                                                                                                var v120 uint64 = (v2 / ((4294967295) | 1))
                                                                                                                                v121 := GenS13{F0: 0, F1: 148, F2: "a"}
                                                                                                                                _ = v119
                                                                                                                                _ = v120
                                                                                                                                _ = v121
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            var v122 string = "\n\t"
                                                                                                                            _ = v122
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case "":
                                                                                                                var v123 int = fi16
                                                                                                                _ = v123
                                                                                                            case "foo bar":
                                                                                                                var v124 uint32 = 2147483647
                                                                                                                _ = v124
                                                                                                            }
                                                                                                            _ = v105
                                                                                                            _ = v106
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                v5 = v5
                                                                                                                v5 = 1e10
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = fi18
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v125 int8 = -1
                                                                                                                fmt.Println(v1)
                                                                                                                _ = v125
                                                                                                            }
                                                                                                            v17 -= (v17 + v17)
                                                                                                            v1 += v1
                                                                                                        }
                                                                                                    case 32768:
                                                                                                        var v126 int32 = -2147483648
                                                                                                        _ = v126
                                                                                                    default:
                                                                                                        if false {
                                                                                                            var v127 string = "a"
                                                                                                            _ = v127
                                                                                                        } else {
                                                                                                            v128 := GenS15{F0: -1, F1: "\n\t", F2: 53}
                                                                                                            var v129 uint = (^(v4))
                                                                                                            var v130 uint = v129
                                                                                                            _ = v128
                                                                                                            _ = v129
                                                                                                            _ = v130
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                v4 = (^(v4))
                                                                                                if true {
                                                                                                    var v131 int = fi18
                                                                                                    _ = len(v3)
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = (v131 % ((v131) | 1))
                                                                                                    }
                                                                                                    v2 += v2
                                                                                                    _ = v131
                                                                                                }
                                                                                                if ((v2 >> ((v4) & 63)) == v2) {
                                                                                                    var v132 int16 = 1000
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = fi18
                                                                                                    }
                                                                                                    v133 := GenS16{F0: fi16, F1: 0, F2: 952}
                                                                                                    _ = v132
                                                                                                    _ = v133
                                                                                                }
                                                                                            }
                                                                                        case "hello":
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                v1 = 803
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        switch "" {
                                                                                        case "a":
                                                                                            v5 += v5
                                                                                        }
                                                                                    default:
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 1
                                                                                        }
                                                                                    }
                                                                                    v4 = 65535
                                                                                }
                                                                                _ = v94
                                                                            } else {
                                                                                v5 -= v5
                                                                                fmt.Println(fi16, v4, v5)
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    v2 -= (v2 >> (((v4 ^ (^(v4)))) & 63))
                                                                                case "":
                                                                                    if ((v4 << ((794) & 63)) > (v4 / ((v4) | 1))) {
                                                                                        v5 = v5
                                                                                    } else {
                                                                                        v1 = v1
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi18, v1)
                                                                            }
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    v4 = v4
                                                                }
                                                            }
                                                        default:
                                                            var v134 uint64 = (v2 >> ((v4) & 63))
                                                            _ = v134
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(fi18)
                                } else {
                                    var v135 int8 = 127
                                    v4 = (v4 | (v4 * (v4 - (^(v4)))))
                                    _ = v135
                                }
                            }
                            if (42 > 42) {
                                _ = len(v3)
                                if true {
                                    v136 := GenS17{F0: -2147483648, F1: "hello", F2: v4}
                                    if (!(((((!false) || true) && true) && (true || (94 <= 0))) && false)) {
                                        switch (^(fi16)) {
                                        case 216:
                                            var v137 float32 = (v5 * v5)
                                            _ = v137
                                        case 139:
                                            if false {
                                                v3 = append(v3, (fi16 << ((255) & 63)))
                                            }
                                            v5 = 1.5
                                            var v138 bool = false
                                            _ = v138
                                        case -1:
                                            switch 2147483647 {
                                            case 841:
                                                var v139 uint64 = (v2 & 4294967295)
                                                _ = v139
                                            default:
                                                var v140 uint32 = 808
                                                _ = v140
                                            }
                                        default:
                                            fmt.Println(fi16, fi16, fi18)
                                        }
                                        switch (fi18 % (((fi18 + int(v17))) | 1)) {
                                        case 32767:
                                            if len(v3) > 0 {
                                                v3[0] = (fi16 ^ fi16)
                                            }
                                        default:
                                            v141 := GenS18{F0: v4}
                                            _ = v141
                                        }
                                        var v142 bool = true
                                        v143 := GenS19{F0: 115, F1: v1}
                                        _ = v142
                                        _ = v143
                                    }
                                    _ = v136
                                } else {
                                    var v144 uint = uint(v2)
                                    _ = v144
                                }
                            } else {
                                var v145 uint64 = (v2 + 4294967295)
                                _ = v145
                            }
                            v146 := GenS20{F0: "a", F1: false}
                            _ = v146
                        case 1:
                            v2 -= (v2 >> ((uint(v1)) & 63))
                        default:
                            var v147 uint32 = 2147483648
                            _ = v147
                        }
                    case "\n\t":
                        v148 := GenS4{F0: fi16}
                        _ = v148
                    case "foo bar":
                        if true {
                            v149 := GenS21{F0: (!true)}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v149
                        }
                    default:
                        v4 -= v4
                    }
                    v4 += (^(v4))
                    _ = fi18
                }
            case 128:
                var v150 uint64 = (^(v2))
                _ = v150
            }
            _ = fi16
            _ = v17
        }
    }
    var v151 float32 = 0.0
    if (0 < -728) {
        if (2 != genFn0((v2 + (^(v2))), "hello")) {
            fmt.Println(v1, v4, v1)
            var v152 uint8 = 0
            _ = v152
        } else {
            v151 = v5
            var v153 int32 = -112
            switch 127 {
            case 32767:
                switch -9223372036854775808 {
                case 1:
                    v153 = v153
                }
            case 2147483647:
                _ = len(v3)
            }
            _ = v153
        }
        if false {
            var v154 string = "hello"
            _ = v154
        } else {
            v5 += v151
        }
    } else {
        var v155 int8 = -1
        var v156 int32 = -362
        fmt.Println(v2, v155)
        _ = v155
        _ = v156
    }
    switch "a" {
    case "a":
        if (32767 != (v1 ^ (v1 % (((v1 - 583)) | 1)))) {
            v157 := GenS13{F0: 0, F1: -1, F2: ("foo bar" + "hello")}
            v4 += v4
            var v158 uint = (v4 | v4)
            _ = v157
            _ = v158
        }
    case "foo bar":
        v4 -= v4
    }
    v2 += (v2 >> ((v4) & 63))
    for fi159 := 0; fi159 < 14; fi159++ {
        switch "" {
        case "foo bar":
            var v160 int64 = -9223372036854775808
            _ = v160
        case "hello":
            var v161 float32 = v151
            _ = v161
        case "\n\t":
            for fi162 := 0; fi162 < 17; fi162++ {
                v163 := GenS19{F0: 0, F1: v1}
                var v164 uint16 = (v1 % ((32767) | 1))
                _ = fi162
                _ = v163
                _ = v164
            }
        }
        var v165 int32 = -288
        _ = fi159
        _ = v165
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v151:", v151)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v151
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
