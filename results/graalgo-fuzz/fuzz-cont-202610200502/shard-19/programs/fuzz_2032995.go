package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int32; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint }
type GenS6 struct { F0 uint8; F1 int8; F2 string }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool; F1 bool }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16; F1 int }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint8; F1 uint16 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int64; F1 string; F2 uint8 }
type GenS19 struct { F0 uint8; F1 string }
type GenS20 struct { F0 bool; F1 bool }
type GenS21 struct { F0 int64; F1 bool }
type GenS22 struct { F0 bool; F1 int32 }
type GenS23 struct { F0 uint; F1 int16 }
type GenS24 struct { F0 int32; F1 uint64 }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 int32 }
type GenS28 struct { F0 uint32 }
type GenS29 struct { F0 string; F1 uint8 }
type GenS30 struct { F0 uint16 }
type GenS31 struct { F0 int8 }
type GenS32 struct { F0 int8 }
type GenS33 struct { F0 int8 }
type GenS34 struct { F0 int8; F1 uint64; F2 string }
type GenS35 struct { F0 string; F1 int8; F2 string }
type GenS36 struct { F0 int }
func genFn0(p0 int8, p1 int8, p2 uint) uint8 {
    p2 -= (^(p2))
    fmt.Println(p1, p2)
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    return 181
}
func genFn1(p0 int64) int32 {
    var v5 uint = 275
    p0 = p0
    if false {
        v6 := GenS1{F0: ("a" + ("a" + ("\n\t" + "a")))}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            if true {
                v5 += 0
                v5 = (v5 ^ 577)
            }
            _ = fi7
        }
        switch "\n\t" {
        case "foo bar":
            p0 = (p0 >> (((v5 & (v5 << (((v5 & (v5 | v5))) & 63)))) & 63))
        case "\n\t":
            fmt.Println(v5)
        case "a":
            for fi8 := 0; fi8 < 19; fi8++ {
                p0 += p0
                v5 = (^(v5))
                _ = fi8
            }
        default:
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int16 = 0
                v5 = (v5 & v5)
                _ = fi9
                _ = v10
            }
        }
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            var v12 uint32 = 29
            if ((fi11 & (fi11 ^ (-(fi11)))) > fi11) {
                v13 := GenS2{F0: genFn0(42, 127, uint(v12))}
                _ = v13
            } else {
                v14 := []int{(-(fi11)), fi11, fi11, int(v5)}
                _ = v14
            }
            _ = v12
        } else {
            v15 := []int{fi11, -214}
            if false {
                if (4294967295 <= 148) {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v15 = append(v15, int(v5))
                        if true {
                            switch (fi16 | (^(fi11))) {
                            case 2147483647:
                                if (!(!(float32(((float64(1.5) * 3.14159) + ((1e10 / float64(1e10)) + (float64(3.14159) * (1.5 * 0.1))))) < (0.0 / 0.1)))) {
                                    var v17 uint16 = 0
                                    var v18 uint64 = 18446744073709551615
                                    v17 = (v17 | v17)
                                    fmt.Println(v18, v18, v18)
                                    p0 = p0
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    var v19 int16 = -1
                                    if len(v15) > 0 {
                                        v15[0] = (-(fi16))
                                    }
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                } else {
                                    if false {
                                        switch "a" {
                                        case "a":
                                            _ = len(v15)
                                        case "":
                                            var v20 string = ""
                                            _ = v20
                                        case "hello":
                                            var v21 int8 = 1
                                            _ = v21
                                        }
                                        if false {
                                            var v22 uint = v5
                                            fmt.Println(v22, v5)
                                            v23 := GenS1{F0: ("a" + "\n\t")}
                                            v15 = append(v15, fi16)
                                            _ = v22
                                            _ = v23
                                        }
                                    } else {
                                        v24 := GenS2{F0: 128}
                                        _ = v24
                                    }
                                }
                            case -1:
                                v25 := GenS3{F0: 255}
                                _ = v25
                            }
                            v26 := GenS4{F0: int(p0), F1: -2147483648, F2: -163}
                            _ = v26
                        } else {
                            var v27 bool = (((0.0 + (float64(0.0) + 3.14159)) * 0.1) >= 0.1)
                            _ = v27
                        }
                        _ = fi16
                    }
                    if false {
                        v15 = append(v15, fi11)
                        if (0.0 <= (0.1 + (1e10 * (float64(float32(1e10)) + 1.5)))) {
                            switch "\n\t" {
                            case "a":
                                if false {
                                    v5 = v5
                                    var v28 int = (fi11 % ((fi11) | 1))
                                    _ = v28
                                }
                            }
                            var v29 string = "hello"
                            _ = v29
                        }
                    } else {
                        var v30 uint16 = 1
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 9; fi31++ {
                        if true {
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    var v32 int32 = 2147483647
                                    if true {
                                        if false {
                                            switch "foo bar" {
                                            case "":
                                                v32 += (^(v32))
                                            }
                                            _ = len(v15)
                                            v33 := GenS5{F0: 2147483647, F1: 171, F2: (v5 | (v5 % (((v5 / ((v5) | 1))) | 1)))}
                                            _ = v33
                                        } else {
                                            if (true || ((-104 > 100) && (1 < 726))) {
                                                if false {
                                                    if true {
                                                        var v34 string = ""
                                                        if false {
                                                            if (false || (p0 <= (p0 | (p0 + -932)))) {
                                                                if true {
                                                                    var v35 float64 = 0.0
                                                                    _ = v35
                                                                }
                                                            }
                                                            var v36 int64 = 0
                                                            switch 127 {
                                                            case -511:
                                                                var v37 int = 4294967295
                                                                _ = v37
                                                            }
                                                            _ = v36
                                                        }
                                                        var v38 uint8 = 255
                                                        _ = v34
                                                        _ = v38
                                                    } else {
                                                        v32 += 42
                                                        if ("hello" == (("hello" + "\n\t") + "\n\t")) {
                                                            v32 = -2147483648
                                                            if (255 < genFn0(int8(fi11), 42, v5)) {
                                                                var v39 uint32 = 664
                                                                if true {
                                                                    var v40 int = -2147483648
                                                                    var v41 int16 = 1000
                                                                    _ = len(v15)
                                                                    v42 := GenS6{F0: 150, F1: 1, F2: (("\n\t" + ("a" + ("a" + "\n\t"))) + (("" + "a") + "hello"))}
                                                                    var v43 uint = (v5 ^ v5)
                                                                    _ = v40
                                                                    _ = v41
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                                _ = v39
                                                            } else {
                                                                if (true || false) {
                                                                    fmt.Println(fi11, fi11)
                                                                    var v44 uint32 = 2147483648
                                                                    _ = len(v15)
                                                                    _ = v44
                                                                } else {
                                                                    v5 += (v5 << (((v5 << (((v5 >> ((v5) & 63))) & 63))) & 63))
                                                                    var v45 float64 = ((0.1 + 1e10) / 1e10)
                                                                    v32 -= 42
                                                                    _ = v45
                                                                }
                                                                if ((17 >= 255) && true) {
                                                                    fmt.Println(p0, p0, v32)
                                                                    if true {
                                                                        p0 = p0
                                                                        fmt.Println(v32)
                                                                    } else {
                                                                        _ = len(v15)
                                                                        fmt.Println(fi11, v5, v5)
                                                                        v46 := GenS7{F0: -32768, F1: false}
                                                                        _ = v46
                                                                    }
                                                                    if false {
                                                                        if ((-(v32)) > v32) {
                                                                            p0 -= p0
                                                                        } else {
                                                                            var v47 float64 = 1.5
                                                                            _ = v47
                                                                        }
                                                                        v5 -= (v5 - 2)
                                                                        var v48 int64 = p0
                                                                        _ = v48
                                                                    } else {
                                                                        p0 = (-(p0))
                                                                        v49 := GenS8{F0: 65535}
                                                                        _ = v49
                                                                    }
                                                                }
                                                                p0 = p0
                                                                if true {
                                                                    if false {
                                                                        var v50 uint32 = 488
                                                                        _ = v50
                                                                    } else {
                                                                        if true {
                                                                            if (911 != p0) {
                                                                                v5 += (v5 >> (((^(v5))) & 63))
                                                                            }
                                                                            p0 += (p0 | p0)
                                                                            if len(v15) >= 2 {
                                                                                v15 = v15[1:len(v15)]
                                                                            }
                                                                        } else {
                                                                            var v51 int64 = (-(p0))
                                                                            v52 := GenS4{F0: fi31, F1: (-(v32)), F2: v51}
                                                                            _ = v51
                                                                            _ = v52
                                                                        }
                                                                        v15 = append(v15, 449)
                                                                    }
                                                                    switch fi31 {
                                                                    case 44:
                                                                        if len(v15) >= 2 {
                                                                            v15 = v15[1:len(v15)]
                                                                        }
                                                                    case 816:
                                                                        var v53 int64 = (p0 / ((p0) | 1))
                                                                        _ = v53
                                                                    default:
                                                                        var v54 int = fi11
                                                                        _ = v54
                                                                    }
                                                                } else {
                                                                    var v55 uint8 = genFn0(-48, -1, 706)
                                                                    v5 -= 321
                                                                    if (!false) {
                                                                        var v56 uint = v5
                                                                        if ((!true) || ((^(fi11)) != fi31)) {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                if (p0 <= p0) {
                                                                                    var v57 uint32 = 42
                                                                                    fmt.Println(v55, p0)
                                                                                    var v58 uint32 = v57
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                } else {
                                                                                    var v59 uint32 = 800
                                                                                    var v60 bool = ((!(323 != 255)) && false)
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                        }
                                                                        var v61 float32 = float32(float64((0.0 / 1.0)))
                                                                        var v62 uint8 = (v55 * (v55 % ((v55) | 1)))
                                                                        v32 = (v32 + 0)
                                                                        p0 = (-(p0))
                                                                        var v63 int32 = -618
                                                                        switch (-(fi11)) {
                                                                        case 32768:
                                                                            v56 -= v5
                                                                        case 0:
                                                                            if (false || (!(!true))) {
                                                                                if true {
                                                                                    if len(v15) > 0 {
                                                                                        v15[0] = fi11
                                                                                    }
                                                                                    if true {
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v5 -= v56
                                                                                            if (-1 >= -332) {
                                                                                                if len(v15) >= 2 {
                                                                                                    v15 = v15[1:len(v15)]
                                                                                                }
                                                                                            }
                                                                                        case "foo bar":
                                                                                            if false {
                                                                                                switch ("hello" + "") {
                                                                                                case "\n\t":
                                                                                                    if ("a" > "\n\t") {
                                                                                                        switch int(v63) {
                                                                                                        case -1:
                                                                                                            v63 -= -1
                                                                                                        }
                                                                                                    } else {
                                                                                                        if ((p0 + (p0 / (((p0 % ((p0) | 1))) | 1))) == 762) {
                                                                                                            _ = len(v15)
                                                                                                            switch fi31 {
                                                                                                            case 32768:
                                                                                                                switch "" {
                                                                                                                case "":
                                                                                                                    v64 := GenS2{F0: v55}
                                                                                                                    _ = v64
                                                                                                                default:
                                                                                                                    v63 += v63
                                                                                                                }
                                                                                                            case -94:
                                                                                                                var v65 uint16 = 685
                                                                                                                _ = v65
                                                                                                            case 127:
                                                                                                                var v66 float32 = v61
                                                                                                                _ = v66
                                                                                                            default:
                                                                                                                v67 := GenS1{F0: ""}
                                                                                                                _ = v67
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v15) >= 2 {
                                                                                                            v15 = v15[1:len(v15)]
                                                                                                        }
                                                                                                        v56 = (^(v56))
                                                                                                        if ((fi31 ^ -930) < fi11) {
                                                                                                            if len(v15) >= 2 {
                                                                                                                v15 = v15[1:len(v15)]
                                                                                                            }
                                                                                                            var v68 uint = v5
                                                                                                            _ = v68
                                                                                                        }
                                                                                                        v69 := GenS7{F0: 32767, F1: (false || true)}
                                                                                                        _ = v69
                                                                                                    }
                                                                                                default:
                                                                                                    var v70 string = "foo bar"
                                                                                                    _ = v70
                                                                                                }
                                                                                            } else {
                                                                                                if ((v55 | v55) < genFn0(-110, 1, v5)) {
                                                                                                    v71 := GenS3{F0: (v62 & (v55 / ((v62) | 1)))}
                                                                                                    var v72 int64 = p0
                                                                                                    fmt.Println(v63, v56)
                                                                                                    _ = v71
                                                                                                    _ = v72
                                                                                                }
                                                                                                v62 -= 2
                                                                                            }
                                                                                        case "a":
                                                                                            if len(v15) >= 2 {
                                                                                                v15 = v15[1:len(v15)]
                                                                                            }
                                                                                        }
                                                                                        var v73 uint = 1
                                                                                        v74 := GenS7{F0: -855, F1: true}
                                                                                        _ = v73
                                                                                        _ = v74
                                                                                    } else {
                                                                                        var v75 string = "\n\t"
                                                                                        if true {
                                                                                            if len(v15) > 0 {
                                                                                                v15[0] = (fi11 - (fi11 << (((v56 << ((v56) & 63))) & 63)))
                                                                                            }
                                                                                            v75 = v75
                                                                                            var v76 float64 = 1.5
                                                                                            p0 += (p0 | p0)
                                                                                            switch ("foo bar" + "") {
                                                                                            case "":
                                                                                                v32 += (v32 / (((v32 & 361)) | 1))
                                                                                            case "hello":
                                                                                                var v77 string = v75
                                                                                                _ = v77
                                                                                            default:
                                                                                                var v78 int32 = 42
                                                                                                _ = v78
                                                                                            }
                                                                                            var v79 int8 = 1
                                                                                            _ = v76
                                                                                            _ = v79
                                                                                        } else {
                                                                                            v80 := GenS3{F0: v55}
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                if ((v5 * 529) >= (v5 ^ v5)) {
                                                                                                    if (32767 != int(v5)) {
                                                                                                        switch ("\n\t" + v75) {
                                                                                                        case "\n\t":
                                                                                                            v81 := GenS9{F0: int16(p0), F1: 931}
                                                                                                            _ = v81
                                                                                                        default:
                                                                                                            if ((^(v56)) >= v5) {
                                                                                                                if (int8(fi11) <= 42) {
                                                                                                                    if false {
                                                                                                                        v5 = v5
                                                                                                                    } else {
                                                                                                                        v82 := GenS10{F0: (v55 * genFn0(1, int8(fi31), 433))}
                                                                                                                        switch ("foo bar" + v75) {
                                                                                                                        case "\n\t":
                                                                                                                            switch v75 {
                                                                                                                            case "foo bar":
                                                                                                                                if (0 > v55) {
                                                                                                                                    v83 := GenS5{F0: 42, F1: (v55 % ((v62) | 1)), F2: (v5 + (v5 - v5))}
                                                                                                                                    _ = v83
                                                                                                                                } else {
                                                                                                                                    if false {
                                                                                                                                        v32 -= (v63 - (v32 * 793))
                                                                                                                                        if len(v15) > 0 {
                                                                                                                                            v15[0] = fi11
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        var v84 uint32 = 434
                                                                                                                                        _ = len(v15)
                                                                                                                                        v84 += v84
                                                                                                                                        _ = v84
                                                                                                                                    }
                                                                                                                                    v5 += (^(v56))
                                                                                                                                    switch (v75 + (v75 + ((v75 + "\n\t") + "\n\t"))) {
                                                                                                                                    case "":
                                                                                                                                        p0 += p0
                                                                                                                                    case "hello":
                                                                                                                                        var v85 uint8 = v62
                                                                                                                                        _ = v85
                                                                                                                                    default:
                                                                                                                                        v56 = 42
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                if ((fi31 / ((fi11) | 1)) != 4294967295) {
                                                                                                                                    v86 := GenS4{F0: 0, F1: (-(v32)), F2: p0}
                                                                                                                                    switch (v75 + "\n\t") {
                                                                                                                                    case "\n\t":
                                                                                                                                        v15 = append(v15, 2147483647)
                                                                                                                                    }
                                                                                                                                    _ = v86
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v87 uint8 = v55
                                                                                                                            _ = v87
                                                                                                                        default:
                                                                                                                            fmt.Println(v62, v63, v5)
                                                                                                                        }
                                                                                                                        fmt.Println(v56, v75)
                                                                                                                        _ = v82
                                                                                                                    }
                                                                                                                    var v88 uint32 = 2147483647
                                                                                                                    v63 += v63
                                                                                                                    _ = v88
                                                                                                                } else {
                                                                                                                    v56 = v5
                                                                                                                    if len(v15) >= 2 {
                                                                                                                        v15 = v15[1:len(v15)]
                                                                                                                    }
                                                                                                                    if ((v61 >= v61) || false) {
                                                                                                                        v75 = v75
                                                                                                                        switch 32768 {
                                                                                                                        case 42:
                                                                                                                            var v89 float32 = v61
                                                                                                                            _ = v89
                                                                                                                        case 744:
                                                                                                                            if false {
                                                                                                                                if (-19 == 0) {
                                                                                                                                    v56 = (v56 ^ 65535)
                                                                                                                                }
                                                                                                                                var v90 uint16 = 32767
                                                                                                                                switch v75 {
                                                                                                                                case "":
                                                                                                                                    v61 = (v61 / float32(1.5))
                                                                                                                                }
                                                                                                                                _ = v90
                                                                                                                            } else {
                                                                                                                                var v91 float32 = float32(3.14159)
                                                                                                                                _ = v91
                                                                                                                            }
                                                                                                                        case 32768:
                                                                                                                            var v92 uint32 = 925
                                                                                                                            _ = v92
                                                                                                                        }
                                                                                                                        v32 = 42
                                                                                                                        if true {
                                                                                                                            var v93 float64 = 0.1
                                                                                                                            v93 -= 1.5
                                                                                                                            var v94 int8 = 42
                                                                                                                            _ = v93
                                                                                                                            _ = v94
                                                                                                                        }
                                                                                                                        switch (fi31 << ((v56) & 63)) {
                                                                                                                        case 128:
                                                                                                                            switch v75 {
                                                                                                                            case "hello":
                                                                                                                                v95 := GenS11{F0: (0.0 >= float64((((v61 + float32((1.5 + (0.0 - 0.0)))) / float32(3.14159)) + float32(-2.25)))), F1: (!true)}
                                                                                                                                _ = v95
                                                                                                                            default:
                                                                                                                                var v96 float32 = (float32(1.5) - 3.14159)
                                                                                                                                _ = v96
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                var v97 string = v75
                                                                                                                switch 32768 {
                                                                                                                case 77:
                                                                                                                    if true {
                                                                                                                        v63 -= -1
                                                                                                                        var v98 int16 = 1000
                                                                                                                        switch v75 {
                                                                                                                        case "hello":
                                                                                                                            v99 := GenS7{F0: (v98 | (v98 ^ (v98 - v98))), F1: true}
                                                                                                                            _ = v99
                                                                                                                        }
                                                                                                                        _ = v98
                                                                                                                    }
                                                                                                                case -1:
                                                                                                                    if false {
                                                                                                                        if true {
                                                                                                                            switch "" {
                                                                                                                            case "\n\t":
                                                                                                                                var v100 int32 = v63
                                                                                                                                _ = v100
                                                                                                                            case "a":
                                                                                                                                if (((fi11 ^ fi31) > fi31) && false) {
                                                                                                                                    var v101 int32 = v32
                                                                                                                                    switch v97 {
                                                                                                                                    case "hello":
                                                                                                                                        v102 := GenS3{F0: 45}
                                                                                                                                        _ = v102
                                                                                                                                    case "a":
                                                                                                                                        if ("hello" != v75) {
                                                                                                                                            v61 += (v61 / ((v61 * (1e10 / (v61 + (float32(float64(v61)) * ((float32(0.1) - v61) - (1.5 * v61)))))) / (1e10 - (-2.25 * 3.14159))))
                                                                                                                                        }
                                                                                                                                    case "":
                                                                                                                                        var v103 int = fi31
                                                                                                                                        _ = v103
                                                                                                                                    }
                                                                                                                                    var v104 int = (fi11 * fi31)
                                                                                                                                    _ = v101
                                                                                                                                    _ = v104
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v97)
                                                                                                                                    var v105 string = v75
                                                                                                                                    var v106 uint = 42
                                                                                                                                    _ = v105
                                                                                                                                    _ = v106
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                v107 := GenS5{F0: 2, F1: (v62 ^ v62), F2: 170}
                                                                                                                                _ = v107
                                                                                                                            default:
                                                                                                                                v108 := GenS11{F0: (!false), F1: true}
                                                                                                                                _ = v108
                                                                                                                            }
                                                                                                                            switch ("hello" + (("foo bar" + ("a" + v97)) + v75)) {
                                                                                                                            case "hello":
                                                                                                                                if (!true) {
                                                                                                                                    v109 := GenS12{F0: -967}
                                                                                                                                    _ = v109
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                if false {
                                                                                                                                    var v110 uint64 = 18446744073709551615
                                                                                                                                    _ = v110
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v111 string = ((v75 + v97) + "\n\t")
                                                                                                                            _ = v111
                                                                                                                        } else {
                                                                                                                            var v112 int32 = v63
                                                                                                                            _ = v112
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v62 -= (v55 - v62)
                                                                                                                        v61 = v61
                                                                                                                        v62 = uint8(v5)
                                                                                                                    }
                                                                                                                case 128:
                                                                                                                    v113 := GenS9{F0: int16(v32), F1: 255}
                                                                                                                    _ = v113
                                                                                                                }
                                                                                                                _ = v97
                                                                                                            } else {
                                                                                                                _ = len(v15)
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        if len(v15) >= 2 {
                                                                                                            v15 = v15[1:len(v15)]
                                                                                                        }
                                                                                                        v15 = append(v15, 4294967295)
                                                                                                        v32 = -2147483648
                                                                                                    }
                                                                                                    v56 = (v56 ^ (v56 ^ 2))
                                                                                                } else {
                                                                                                    v114 := GenS3{F0: genFn0(-1, 96, 679)}
                                                                                                    var v115 int16 = 542
                                                                                                    _ = v114
                                                                                                    _ = v115
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v116 := GenS9{F0: -32768, F1: 593}
                                                                                                _ = v116
                                                                                            default:
                                                                                                v5 = 42
                                                                                            }
                                                                                            v62 = (v62 | (v55 ^ v55))
                                                                                            _ = v80
                                                                                        }
                                                                                        if len(v15) >= 2 {
                                                                                            v15 = v15[1:len(v15)]
                                                                                        }
                                                                                        _ = v75
                                                                                    }
                                                                                    if (false && false) {
                                                                                        v15 = append(v15, -2147483648)
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            if (genFn0(100, 1, v56) != (v62 / ((v55) | 1))) {
                                                                                                v117 := GenS13{F0: int(v32)}
                                                                                                _ = v117
                                                                                            }
                                                                                        case "hello":
                                                                                            v62 = (v62 << ((v5) & 63))
                                                                                        case "\n\t":
                                                                                            v118 := GenS8{F0: 705}
                                                                                            _ = v118
                                                                                        }
                                                                                        var v119 uint64 = 42
                                                                                        _ = v119
                                                                                    }
                                                                                    v120 := GenS2{F0: v55}
                                                                                    v56 = v5
                                                                                    p0 = -544
                                                                                    _ = v120
                                                                                } else {
                                                                                    v63 -= (-(v32))
                                                                                }
                                                                                var v121 int16 = 1
                                                                                switch ("hello" + "hello") {
                                                                                case "a":
                                                                                    v122 := GenS9{F0: (^(v121)), F1: 255}
                                                                                    _ = v122
                                                                                case "hello":
                                                                                    var v123 int8 = 112
                                                                                    _ = v123
                                                                                case "foo bar":
                                                                                    v62 = v55
                                                                                    var v124 int16 = -32768
                                                                                    var v125 bool = ((v63 < (v32 ^ v63)) || (!true))
                                                                                    _ = v124
                                                                                    _ = v125
                                                                                }
                                                                                switch 919 {
                                                                                case -1:
                                                                                    v126 := GenS7{F0: int16(p0), F1: true}
                                                                                    _ = v126
                                                                                case 2147483647:
                                                                                    if false {
                                                                                        if ((v5 << ((v56) & 63)) < (v5 % ((948) | 1))) {
                                                                                            fmt.Println(v61)
                                                                                            p0 += p0
                                                                                            v127 := GenS12{F0: v121}
                                                                                            _ = v127
                                                                                        } else {
                                                                                            v128 := GenS13{F0: fi31}
                                                                                            v56 = 255
                                                                                            v129 := GenS14{F0: 1, F1: (fi11 | (fi11 | (fi11 % ((1) | 1))))}
                                                                                            if true {
                                                                                                var v130 uint32 = 0
                                                                                                fmt.Println(v55, v61, fi11)
                                                                                                if (fi31 < fi11) {
                                                                                                    if (false && ((v5 + v5) <= v5)) {
                                                                                                        var v131 int8 = 0
                                                                                                        var v132 bool = ("foo bar" >= "hello")
                                                                                                        var v133 bool = v132
                                                                                                        switch (fi11 << (((v5 ^ (v56 ^ 0))) & 63)) {
                                                                                                        case -41:
                                                                                                            var v134 int16 = (v121 + v121)
                                                                                                            _ = v134
                                                                                                        case -230:
                                                                                                            switch ("hello" + "foo bar") {
                                                                                                            case "a":
                                                                                                                v132 = true
                                                                                                            case "":
                                                                                                                v121 += 1
                                                                                                            case "\n\t":
                                                                                                                var v135 bool = (p0 > -9223372036854775808)
                                                                                                                _ = v135
                                                                                                            }
                                                                                                            if (411 >= 32767) {
                                                                                                                switch (fi31 % ((fi31) | 1)) {
                                                                                                                case -2147483648:
                                                                                                                    var v136 uint16 = 112
                                                                                                                    _ = v136
                                                                                                                case 2147483647:
                                                                                                                    var v137 uint64 = 18446744073709551615
                                                                                                                    _ = v137
                                                                                                                }
                                                                                                                fmt.Println(v56)
                                                                                                                v138 := GenS15{F0: -626}
                                                                                                                v139 := GenS3{F0: v62}
                                                                                                                var v140 uint8 = (v62 - (v62 + (v62 & 155)))
                                                                                                                _ = v138
                                                                                                                _ = v139
                                                                                                                _ = v140
                                                                                                            }
                                                                                                            v121 -= v121
                                                                                                        case -893:
                                                                                                            v141 := GenS8{F0: v130}
                                                                                                            _ = v141
                                                                                                        default:
                                                                                                            var v142 float32 = v61
                                                                                                            _ = v142
                                                                                                        }
                                                                                                        _ = v131
                                                                                                        _ = v132
                                                                                                        _ = v133
                                                                                                    }
                                                                                                    switch 172 {
                                                                                                    case 1:
                                                                                                        v121 = v121
                                                                                                    case -997:
                                                                                                        var v143 uint8 = (v62 | 127)
                                                                                                        _ = v143
                                                                                                    case -419:
                                                                                                        switch ("foo bar" + ("hello" + "\n\t")) {
                                                                                                        case "foo bar":
                                                                                                            v130 = 40
                                                                                                        case "hello":
                                                                                                            v130 = (v130 & v130)
                                                                                                        default:
                                                                                                            switch ("" + "foo bar") {
                                                                                                            case "":
                                                                                                                v55 = (^(v62))
                                                                                                                switch ("foo bar" + ("foo bar" + "\n\t")) {
                                                                                                                case "foo bar":
                                                                                                                    var v144 uint32 = v130
                                                                                                                    _ = v144
                                                                                                                case "":
                                                                                                                    v32 -= (^(v32))
                                                                                                                default:
                                                                                                                    v63 = int32(v5)
                                                                                                                }
                                                                                                                v61 += v61
                                                                                                            default:
                                                                                                                var v145 int = fi11
                                                                                                                _ = v145
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v55 -= (v55 + v62)
                                                                                                    }
                                                                                                    fmt.Println(fi31)
                                                                                                }
                                                                                                _ = v130
                                                                                            }
                                                                                            if (v62 == v62) {
                                                                                                switch 32768 {
                                                                                                case -749:
                                                                                                    if true {
                                                                                                        v56 = 4294967295
                                                                                                        if (((2 != 42) || false) || (v63 < (v63 | int32(v5)))) {
                                                                                                            v146 := GenS12{F0: 1}
                                                                                                            var v147 uint8 = (v62 | 6)
                                                                                                            _ = v146
                                                                                                            _ = v147
                                                                                                        } else {
                                                                                                            fmt.Println(v32)
                                                                                                            fmt.Println(v61, v121, fi11)
                                                                                                        }
                                                                                                    }
                                                                                                case 38:
                                                                                                    fmt.Println(p0, v61, v63)
                                                                                                }
                                                                                            }
                                                                                            v148 := GenS10{F0: 128}
                                                                                            if false {
                                                                                                switch (("hello" + "") + "a") {
                                                                                                case "":
                                                                                                    switch (fi11 & fi31) {
                                                                                                    case 9223372036854775807:
                                                                                                        fmt.Println(v32, v121, v63)
                                                                                                    default:
                                                                                                        if (2147483648 == 65535) {
                                                                                                            v63 -= v32
                                                                                                            var v149 int16 = v121
                                                                                                            if (false && false) {
                                                                                                                switch (fi11 & -595) {
                                                                                                                case 9223372036854775807:
                                                                                                                    if (((v61 + (0.0 * v61)) * float32(1.5)) <= (0.1 - (((3.14159 - (v61 * float32(1.5))) / v61) * (v61 + 1e10)))) {
                                                                                                                        _ = len(v15)
                                                                                                                    } else {
                                                                                                                        v150 := GenS7{F0: v121, F1: (((v55 ^ genFn0(127, 100, 65535)) < v62) && (128 > (fi11 * (fi11 ^ fi31))))}
                                                                                                                        if len(v15) >= 2 {
                                                                                                                            v15 = v15[1:len(v15)]
                                                                                                                        }
                                                                                                                        _ = v150
                                                                                                                    }
                                                                                                                case 210:
                                                                                                                    switch (-(fi11)) {
                                                                                                                    case -9223372036854775808:
                                                                                                                        switch "a" {
                                                                                                                        case "\n\t":
                                                                                                                            fmt.Println(v55)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v121 -= v121
                                                                                                                    }
                                                                                                                }
                                                                                                                var v151 int64 = int64(v5)
                                                                                                                if true {
                                                                                                                    v152 := GenS15{F0: -631}
                                                                                                                    switch -9223372036854775808 {
                                                                                                                    case 807:
                                                                                                                        if true {
                                                                                                                            v62 += 2
                                                                                                                            _ = len(v15)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v56 = v56
                                                                                                                    }
                                                                                                                    _ = len(v15)
                                                                                                                    _ = v152
                                                                                                                }
                                                                                                                v63 = 1
                                                                                                                _ = v151
                                                                                                            } else {
                                                                                                                var v153 uint32 = uint32(v5)
                                                                                                                v154 := GenS14{F0: v121, F1: fi31}
                                                                                                                _ = v153
                                                                                                                _ = v154
                                                                                                            }
                                                                                                            _ = v149
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v61 = (1.5 - 0.0)
                                                                                                case "foo bar":
                                                                                                    v56 = (v5 - 42)
                                                                                                }
                                                                                                var v155 int = (fi11 ^ (fi31 / ((fi11) | 1)))
                                                                                                fmt.Println(v63, v55, v63)
                                                                                                _ = v155
                                                                                            } else {
                                                                                                v62 -= (v55 & v55)
                                                                                                var v156 int = 42
                                                                                                _ = v156
                                                                                            }
                                                                                            v56 -= (v56 + (^(v5)))
                                                                                            v55 = v55
                                                                                            _ = v128
                                                                                            _ = v129
                                                                                            _ = v148
                                                                                        }
                                                                                        p0 = (p0 % ((p0) | 1))
                                                                                    }
                                                                                }
                                                                                _ = v121
                                                                            }
                                                                        }
                                                                        _ = v56
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v63
                                                                    } else {
                                                                        v157 := GenS12{F0: 1000}
                                                                        _ = v157
                                                                    }
                                                                    _ = v55
                                                                }
                                                                fmt.Println(p0, fi31)
                                                            }
                                                        } else {
                                                            var v158 float64 = -2.25
                                                            if (fi11 <= (-(fi31))) {
                                                                if len(v15) > 0 {
                                                                    v15[0] = (fi31 % ((fi11) | 1))
                                                                }
                                                                v158 = 0.0
                                                                v159 := GenS7{F0: -32768, F1: (float64((0.1 * 3.14159)) != v158)}
                                                                _ = v159
                                                            } else {
                                                                p0 = p0
                                                            }
                                                            switch (fi31 * 0) {
                                                            case 4294967295:
                                                                if true {
                                                                    var v160 int = -966
                                                                    _ = v160
                                                                }
                                                            }
                                                            _ = v158
                                                        }
                                                    }
                                                    _ = len(v15)
                                                    var v161 int32 = v32
                                                    _ = v161
                                                } else {
                                                    fmt.Println(p0, v32)
                                                    fmt.Println(fi11, fi31, p0)
                                                    var v162 float32 = 0.1
                                                    _ = v162
                                                }
                                                fmt.Println(fi31)
                                            }
                                        }
                                        v163 := GenS13{F0: (fi11 | fi11)}
                                        if true {
                                            if true {
                                                v32 += v32
                                                var v164 uint8 = 255
                                                _ = v164
                                            } else {
                                                var v165 int = fi31
                                                _ = v165
                                            }
                                            fmt.Println(fi31, fi31, fi11)
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                        }
                                        var v166 uint64 = 0
                                        _ = v163
                                        _ = v166
                                    }
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v5, fi31, v5)
                                    case "\n\t":
                                        switch 32768 {
                                        case 9223372036854775807:
                                            fmt.Println(p0)
                                        case 2147483647:
                                            var v167 bool = (!(0.0 > float32(0.1)))
                                            _ = v167
                                        }
                                    }
                                    v15 = append(v15, -1)
                                    _ = v32
                                } else {
                                    p0 = 9223372036854775807
                                    v5 -= (v5 % ((v5) | 1))
                                    var v168 int16 = 1000
                                    if (1 > -1) {
                                        switch fi31 {
                                        case 32767:
                                            var v169 uint64 = 381
                                            _ = v169
                                        case 32768:
                                            v170 := GenS16{F0: 150, F1: 255}
                                            _ = v170
                                        }
                                    }
                                    fmt.Println(fi31, fi31)
                                    _ = v168
                                }
                            case "a":
                                v5 = v5
                            default:
                                v171 := GenS5{F0: 338, F1: 255, F2: (v5 >> (((v5 ^ v5)) & 63))}
                                _ = v171
                            }
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            var v172 int32 = 2147483647
                            _ = v172
                        }
                        var v173 int32 = 0
                        v173 = v173
                        v174 := GenS17{F0: ""}
                        if (42 != 661) {
                            fmt.Println(fi11)
                        } else {
                            v175 := GenS18{F0: p0, F1: (("" + "a") + ""), F2: 42}
                            var v176 int32 = (v173 % ((v173) | 1))
                            _ = v175
                            _ = v176
                        }
                        _ = fi31
                        _ = v173
                        _ = v174
                    }
                } else {
                    fmt.Println(v5, fi11, fi11)
                }
                v5 = v5
            }
            _ = v15
        }
        v177 := []int{128, 0, int(p0), (fi11 | fi11)}
        _ = fi11
        _ = v177
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := []uint32{959}
    v2 := []string{"foo bar"}
    var v3 bool = true
    v4 := []uint32{2147483648, 191}
    switch -59 {
    case 14:
        if false {
            v3 = ((!(v3 || v3)) || ((0 > 113) && false))
            fmt.Println(v3)
        } else {
            for fi178 := 0; fi178 < 5; fi178++ {
                var v179 int64 = 42
                switch (fi178 >> ((2) & 63)) {
                case 2147483647:
                    switch fi178 {
                    case 630:
                        var v180 string = ("" + "\n\t")
                        _ = v180
                    case 32768:
                        var v181 uint32 = 4294967295
                        _ = v181
                    case 42:
                        fmt.Println(v179, fi178)
                    default:
                        v179 = (v179 - v179)
                    }
                default:
                    v182 := GenS19{F0: 0, F1: ""}
                    _ = v182
                }
                _ = fi178
                _ = v179
            }
            for fi183 := 0; fi183 < 17; fi183++ {
                if v3 {
                    v3 = true
                    var v184 uint32 = 372
                    fmt.Println(v184, v3)
                    _ = v184
                }
                _ = fi183
            }
        }
    case 124:
        v185 := GenS3{F0: 1}
        _ = v185
    case 322:
        switch "foo bar" {
        case "a":
            switch ("a" + "foo bar") {
            case "a":
                var v186 uint16 = 65535
                _ = v186
            case "hello":
                switch "\n\t" {
                case "\n\t":
                    v3 = v3
                case "hello":
                    switch "a" {
                    case "a":
                        if (!(904 > 18446744073709551615)) {
                            for fi187 := 0; fi187 < 2; fi187++ {
                                for fi188 := 0; fi188 < 5; fi188++ {
                                    var v189 int = (fi187 * fi188)
                                    v3 = v3
                                    var v190 string = "\n\t"
                                    _ = fi188
                                    _ = v189
                                    _ = v190
                                }
                                if true {
                                    var v191 int32 = genFn1(9223372036854775807)
                                    _ = v191
                                } else {
                                    if v3 {
                                        for fi192 := 0; fi192 < 15; fi192++ {
                                            var v193 string = "\n\t"
                                            fmt.Println(fi192, v193)
                                            _ = fi192
                                            _ = v193
                                        }
                                        v194 := GenS12{F0: 32767}
                                        _ = v194
                                    }
                                    if v3 {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v195 uint16 = 32767
                                            _ = v195
                                        }
                                        v3 = true
                                        v196 := GenS12{F0: 0}
                                        _ = v196
                                    } else {
                                        var v197 int = 32767
                                        var v198 float32 = (1.5 * 0.0)
                                        _ = v197
                                        _ = v198
                                    }
                                    switch fi187 {
                                    case 376:
                                        v3 = (2147483648 <= 220)
                                    }
                                }
                                var v199 int16 = 32767
                                _ = fi187
                                _ = v199
                            }
                            v3 = (1 != genFn1(42))
                            if (1000 > 42) {
                                v2 = append(v2, "foo bar")
                                var v200 int32 = 2147483647
                                v200 -= v200
                                _ = v200
                            } else {
                                _ = len(v2)
                                if v3 {
                                    var v201 uint16 = 984
                                    _ = v201
                                } else {
                                    fmt.Println(v3)
                                    v3 = (2 <= 128)
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                }
                            }
                        } else {
                            for fi202 := 0; fi202 < 1; fi202++ {
                                for fi203 := 0; fi203 < 13; fi203++ {
                                    fmt.Println(v3)
                                    v3 = v3
                                    _ = fi203
                                }
                                _ = len(v2)
                                var v204 float64 = ((1e10 + -2.25) / 0.1)
                                _ = fi202
                                _ = v204
                            }
                            if (!(!v3)) {
                                for fi205 := 0; fi205 < 14; fi205++ {
                                    v3 = true
                                    _ = fi205
                                }
                                _ = len(v2)
                                v3 = false
                            } else {
                                switch 128 {
                                case -1:
                                    var v206 uint64 = 42
                                    _ = v206
                                }
                            }
                            for fi207 := 0; fi207 < 3; fi207++ {
                                if false {
                                    fmt.Println(fi207, v3)
                                    var v208 int16 = -1
                                    if (((-2.25 / (-2.25 / (3.14159 - 0.1))) - float64(1e10)) > (float64((1e10 + 0.1)) * 1e10)) {
                                        for fi209 := 0; fi209 < 13; fi209++ {
                                            var v210 uint32 = 42
                                            if true {
                                                v208 += v208
                                                switch "foo bar" {
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "a":
                                                        v3 = v3
                                                    case "\n\t":
                                                        if (v3 || true) {
                                                            var v211 float32 = ((float32((-2.25 + -2.25)) + float32(0.1)) / 1.0)
                                                            v212 := GenS16{F0: 1, F1: 748}
                                                            _ = v211
                                                            _ = v212
                                                        } else {
                                                            var v213 float32 = 0.0
                                                            _ = v213
                                                        }
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                case "foo bar":
                                                    _ = len(v1)
                                                case "":
                                                    fmt.Println(v208, fi207)
                                                }
                                            }
                                            _ = fi209
                                            _ = v210
                                        }
                                        var v214 uint8 = 42
                                        _ = v214
                                    } else {
                                        var v215 bool = (4294967295 <= 2147483648)
                                        switch "foo bar" {
                                        case "\n\t":
                                            if false {
                                                var v216 uint8 = 0
                                                if (494 > v208) {
                                                    var v217 int8 = 42
                                                    v208 -= v208
                                                    if v3 {
                                                        if (((0.0 - -2.25) + 0.0) != (((0.0 + 1.5) - 0.1) * (((-2.25 - (-2.25 * (((0.1 * (0.1 / 1.0)) + float64((float32(0.1) / float32(1e10)))) * 0.0))) + -2.25) - 0.1))) {
                                                            for fi218 := 0; fi218 < 2; fi218++ {
                                                                var v219 int16 = (v208 * 68)
                                                                var v220 uint64 = 451
                                                                if v3 {
                                                                    var v221 float64 = 1.5
                                                                    var v222 int = 32768
                                                                    _ = v221
                                                                    _ = v222
                                                                }
                                                                _ = fi218
                                                                _ = v219
                                                                _ = v220
                                                            }
                                                            fmt.Println(v208)
                                                            for fi223 := 0; fi223 < 19; fi223++ {
                                                                v224 := GenS20{F0: (v217 == (v217 | (v217 ^ (^(v217))))), F1: true}
                                                                v215 = (v215 && ((float32((((0.1 * 1e10) / 3.14159) - (((1e10 * ((1e10 / (-2.25 * 1e10)) / (float64(0.1) + (0.1 / 3.14159)))) * (float64(float32((1.5 - 0.0))) + float64(0.1))) - 3.14159))) != 0.1) || (938 == 978)))
                                                                var v225 string = ""
                                                                _ = fi223
                                                                _ = v224
                                                                _ = v225
                                                            }
                                                        }
                                                        var v226 string = "foo bar"
                                                        v3 = true
                                                        _ = v226
                                                    } else {
                                                        v216 -= v216
                                                    }
                                                    _ = v217
                                                } else {
                                                    var v227 bool = v215
                                                    _ = v227
                                                }
                                                v228 := GenS11{F0: ((!(!(!true))) && true), F1: v3}
                                                _ = v216
                                                _ = v228
                                            }
                                        case "a":
                                            for fi229 := 0; fi229 < 6; fi229++ {
                                                var v230 int = 9223372036854775807
                                                fmt.Println(fi229)
                                                v230 += 513
                                                _ = fi229
                                                _ = v230
                                            }
                                        case "":
                                            var v231 float32 = -2.25
                                            _ = v231
                                        }
                                        _ = v215
                                    }
                                    _ = v208
                                }
                                if (!("foo bar" > "a")) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    for fi232 := 0; fi232 < 15; fi232++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v3 = (false || v3)
                                        v233 := GenS21{F0: -670, F1: true}
                                        _ = fi232
                                        _ = v233
                                    }
                                    var v234 string = ("foo bar" + (("" + (("foo bar" + "\n\t") + "\n\t")) + "\n\t"))
                                    _ = v234
                                } else {
                                    v235 := GenS1{F0: ""}
                                    v3 = v3
                                    _ = v235
                                }
                                for fi236 := 0; fi236 < 6; fi236++ {
                                    v3 = (true || (0 == 42))
                                    var v237 bool = v3
                                    if (1 >= 0) {
                                        v3 = true
                                        fmt.Println(v3, v3)
                                        fmt.Println(v3)
                                    }
                                    if (1 > 127) {
                                        _ = len(v4)
                                    } else {
                                        if v3 {
                                            v2 = append(v2, "a")
                                            if true {
                                                if (!(618 >= 2)) {
                                                    v237 = v3
                                                    if len(v2) > 0 {
                                                        v2[0] = ""
                                                    }
                                                    var v238 string = "a"
                                                    _ = v238
                                                } else {
                                                    v2 = append(v2, ("foo bar" + ("a" + "")))
                                                    if len(v4) > 0 {
                                                        v4[0] = 835
                                                    }
                                                }
                                                v237 = false
                                                var v239 string = ""
                                                _ = v239
                                            }
                                        }
                                        v3 = (-9223372036854775808 == -1)
                                        var v240 uint64 = 1
                                        v237 = ("foo bar" != "a")
                                        _ = v240
                                    }
                                    _ = fi236
                                    _ = v237
                                }
                                _ = fi207
                            }
                        }
                    case "":
                        for fi241 := 0; fi241 < 8; fi241++ {
                            var v242 int64 = 42
                            v242 = v242
                            switch 601 {
                            case -494:
                                if v3 {
                                    var v243 bool = (v3 && ((!(!v3)) && (v242 != int64(fi241))))
                                    _ = v243
                                } else {
                                    for fi244 := 0; fi244 < 9; fi244++ {
                                        if v3 {
                                            v3 = (v3 || v3)
                                        } else {
                                            v242 = v242
                                        }
                                        _ = fi244
                                    }
                                    if ((-392 < 0) || (!true)) {
                                        for fi245 := 0; fi245 < 19; fi245++ {
                                            var v246 uint = 2
                                            var v247 int8 = -1
                                            _ = fi245
                                            _ = v246
                                            _ = v247
                                        }
                                    }
                                    v248 := GenS22{F0: (67 != 867), F1: genFn1(v242)}
                                    _ = v248
                                }
                            case 4294967295:
                                v249 := GenS21{F0: v242, F1: ((fi241 < (fi241 & (fi241 / ((-185) | 1)))) && (v3 && false))}
                                _ = v249
                            default:
                                _ = len(v2)
                            }
                            _ = fi241
                            _ = v242
                        }
                    default:
                        fmt.Println(v3)
                    }
                case "a":
                    var v250 float64 = ((((0.0 / 3.14159) / 1.0) + (1.5 * 0.1)) * (1.5 + -2.25))
                    _ = v250
                }
            case "\n\t":
                var v251 bool = true
                _ = v251
            }
        case "\n\t":
            v252 := GenS14{F0: 1, F1: -1}
            _ = v252
        default:
            v3 = (false || v3)
        }
    default:
        switch 199 {
        case 633:
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        case 127:
            var v253 uint = 1
            _ = v253
        case -157:
            v254 := GenS3{F0: 42}
            _ = v254
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "" {
    case "foo bar":
        var v255 uint8 = 2
        _ = v255
    case "hello":
        var v256 string = "\n\t"
        _ = v256
    case "a":
        v257 := GenS3{F0: 226}
        _ = v257
    default:
        var v258 int16 = -94
        _ = v258
    }
    v259 := GenS23{F0: 648, F1: 1000}
    for fi260 := 0; fi260 < 2; fi260++ {
        switch fi260 {
        case 127:
            v261 := GenS4{F0: fi260, F1: 1, F2: 1}
            _ = v261
        case 73:
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        case 128:
            _ = len(v2)
        default:
            switch 42 {
            case 128:
                var v262 int16 = -705
                _ = v262
            case -405:
                for fi263 := 0; fi263 < 17; fi263++ {
                    v264 := GenS13{F0: fi263}
                    v3 = v3
                    _ = fi263
                    _ = v264
                }
            case -2147483648:
                v3 = v3
            default:
                var v265 int8 = -1
                _ = v265
            }
        }
        v266 := GenS24{F0: 0, F1: 0}
        fmt.Println(v3, fi260)
        _ = fi260
        _ = v266
    }
    if true {
        v3 = (42 <= 18446744073709551615)
    } else {
        if (295 <= 1000) {
            v267 := GenS18{F0: -1, F1: "hello", F2: 2}
            v3 = v3
            if (!(577 < 42)) {
                if v3 {
                    var v268 int64 = 42
                    if (42 >= -520) {
                        for fi269 := 0; fi269 < 2; fi269++ {
                            v270 := GenS23{F0: 274, F1: -1}
                            _ = fi269
                            _ = v270
                        }
                        v271 := GenS25{F0: -1}
                        if (!("foo bar" != "foo bar")) {
                            v272 := GenS26{F0: v3}
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                            _ = v272
                        }
                        _ = v271
                    } else {
                        v273 := GenS7{F0: 32767, F1: (608 > 1)}
                        var v274 string = ("a" + "a")
                        _ = v273
                        _ = v274
                    }
                    _ = v268
                }
            } else {
                v3 = (32767 > 0)
                if v3 {
                    v3 = ((127 < 100) || v3)
                    if len(v1) > 0 {
                        v1[0] = 117
                    }
                }
                if (!(-1 > 230)) {
                    var v275 float64 = (3.14159 * 3.14159)
                    v275 -= (0.0 * v275)
                    _ = v275
                } else {
                    if (false || (!v3)) {
                        var v276 int64 = 1
                        v277 := GenS27{F0: genFn1((v276 - v276))}
                        for fi278 := 0; fi278 < 14; fi278++ {
                            var v279 int16 = -143
                            v276 = -1
                            _ = fi278
                            _ = v279
                        }
                        _ = v276
                        _ = v277
                    } else {
                        if v3 {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                        }
                        var v280 int32 = 275
                        _ = v280
                    }
                    var v281 uint = 223
                    var v282 uint = 506
                    _ = v281
                    _ = v282
                }
            }
            _ = v267
        }
        if (519 >= 4294967295) {
            switch 32768 {
            case 0:
                for fi283 := 0; fi283 < 13; fi283++ {
                    v3 = false
                    v284 := GenS1{F0: ""}
                    v3 = (204 > 255)
                    _ = fi283
                    _ = v284
                }
            case 42:
                for fi285 := 0; fi285 < 2; fi285++ {
                    var v286 uint = 1
                    for fi287 := 0; fi287 < 6; fi287++ {
                        var v288 uint64 = 1
                        _ = fi287
                        _ = v288
                    }
                    v289 := GenS28{F0: 830}
                    _ = fi285
                    _ = v286
                    _ = v289
                }
            case 9223372036854775807:
                v3 = false
            default:
                for fi290 := 0; fi290 < 15; fi290++ {
                    var v291 int = fi290
                    v292 := GenS5{F0: 42, F1: 103, F2: 65535}
                    v3 = v3
                    _ = fi290
                    _ = v291
                    _ = v292
                }
            }
            switch "\n\t" {
            case "\n\t":
                for fi293 := 0; fi293 < 16; fi293++ {
                    fmt.Println(v3)
                    v294 := GenS6{F0: 0, F1: -22, F2: ("a" + ("\n\t" + (("" + "\n\t") + ("hello" + ("" + ("foo bar" + "a"))))))}
                    _ = fi293
                    _ = v294
                }
            case "hello":
                if (!(float64(1.5) == float64(1e10))) {
                    if v3 {
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                        if v3 {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v295 := GenS7{F0: 160, F1: false}
                            _ = v295
                        } else {
                            v3 = (1e10 < float64(-2.25))
                        }
                        switch 9223372036854775807 {
                        case -205:
                            if v3 {
                                switch -9223372036854775808 {
                                case 4294967295:
                                    if v3 {
                                        for fi296 := 0; fi296 < 6; fi296++ {
                                            for fi297 := 0; fi297 < 5; fi297++ {
                                                fmt.Println(fi296)
                                                _ = fi297
                                            }
                                            _ = fi296
                                        }
                                        if (9223372036854775807 <= 42) {
                                            if v3 {
                                                fmt.Println(v3)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v298 := GenS15{F0: -1}
                                                _ = v298
                                            }
                                            var v299 uint64 = 590
                                            v300 := GenS29{F0: "", F1: genFn0(-128, 1, 2)}
                                            _ = v299
                                            _ = v300
                                        }
                                    } else {
                                        var v301 float64 = 1e10
                                        _ = len(v1)
                                        for fi302 := 0; fi302 < 4; fi302++ {
                                            var v303 uint16 = 932
                                            _ = fi302
                                            _ = v303
                                        }
                                        _ = v301
                                    }
                                case 32768:
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch 4294967295 {
                                        case 0:
                                            var v304 uint = 4294967295
                                            _ = v304
                                        case 42:
                                            var v305 int16 = 65
                                            _ = v305
                                        case 2147483647:
                                            switch 441 {
                                            case -876:
                                                fmt.Println(v3)
                                            case 32767:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 32768:
                                                if true {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v306 uint8 = 127
                                                            _ = v306
                                                        case "\n\t":
                                                            for fi307 := 0; fi307 < 10; fi307++ {
                                                                fmt.Println(fi307)
                                                                if (v3 || true) {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                                _ = fi307
                                                            }
                                                        default:
                                                            for fi308 := 0; fi308 < 12; fi308++ {
                                                                var v309 bool = (true || (-1 > 1))
                                                                _ = fi308
                                                                _ = v309
                                                            }
                                                        }
                                                    default:
                                                        var v310 int16 = 164
                                                        _ = v310
                                                    }
                                                }
                                            default:
                                                v311 := GenS3{F0: 128}
                                                _ = v311
                                            }
                                        default:
                                            v3 = true
                                        }
                                    }
                                case -2147483648:
                                    v3 = (true || v3)
                                default:
                                    for fi312 := 0; fi312 < 17; fi312++ {
                                        switch ("a" + "a") {
                                        case "":
                                            for fi313 := 0; fi313 < 18; fi313++ {
                                                v3 = ("a" == ("" + "hello"))
                                                fmt.Println(fi313, v3, fi312)
                                                var v314 int64 = 42
                                                _ = fi313
                                                _ = v314
                                            }
                                        }
                                        v2 = append(v2, ("a" + "foo bar"))
                                        switch "" {
                                        case "\n\t":
                                            v315 := GenS25{F0: 343}
                                            _ = v315
                                        }
                                        _ = fi312
                                    }
                                }
                                var v316 string = "foo bar"
                                _ = v316
                            } else {
                                var v317 float32 = 0.1
                                if true {
                                    v318 := GenS8{F0: 4294967295}
                                    var v319 uint8 = 2
                                    fmt.Println(v319, v3, v317)
                                    _ = v318
                                    _ = v319
                                } else {
                                    for fi320 := 0; fi320 < 14; fi320++ {
                                        var v321 uint64 = 1
                                        v321 = (v321 << ((0) & 63))
                                        _ = fi320
                                        _ = v321
                                    }
                                    for fi322 := 0; fi322 < 18; fi322++ {
                                        v323 := GenS30{F0: 279}
                                        _ = fi322
                                        _ = v323
                                    }
                                }
                                v324 := GenS15{F0: genFn1(0)}
                                _ = v317
                                _ = v324
                            }
                        case 32768:
                            switch (("foo bar" + "foo bar") + "") {
                            case "\n\t":
                                switch ("\n\t" + ((("a" + "hello") + "") + "\n\t")) {
                                case "foo bar":
                                    var v325 float32 = ((float32(0.0) - 3.14159) / 0.1)
                                    _ = v325
                                default:
                                    for fi326 := 0; fi326 < 20; fi326++ {
                                        switch (fi326 - -725) {
                                        case 32767:
                                            var v327 string = ("\n\t" + ("hello" + "\n\t"))
                                            _ = v327
                                        case -497:
                                            v1 = append(v1, 2147483647)
                                        case 259:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v3 = true
                                        _ = fi326
                                    }
                                }
                            default:
                                var v328 int = 32767
                                _ = v328
                            }
                        default:
                            for fi329 := 0; fi329 < 20; fi329++ {
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                                fmt.Println(fi329, fi329)
                                _ = fi329
                            }
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 223
                        }
                        var v330 string = "\n\t"
                        switch v330 {
                        case "foo bar":
                            if (0 <= 633) {
                                var v331 uint = 304
                                if len(v4) > 0 {
                                    v4[0] = 2
                                }
                                fmt.Println(v330, v3, v331)
                                _ = v331
                            } else {
                                v332 := GenS31{F0: 100}
                                _ = v332
                            }
                        case "\n\t":
                            if (v3 && v3) {
                                var v333 uint64 = 281
                                switch 128 {
                                case -877:
                                    var v334 int16 = 0
                                    _ = v334
                                case 128:
                                    if (!v3) {
                                        var v335 int = 728
                                        v336 := GenS32{F0: 114}
                                        _ = v335
                                        _ = v336
                                    }
                                case 1:
                                    v337 := GenS12{F0: -32768}
                                    _ = v337
                                }
                                _ = v333
                            }
                        }
                        _ = v330
                    }
                    if len(v1) > 0 {
                        v1[0] = 119
                    }
                    v1 = append(v1, 305)
                } else {
                    for fi338 := 0; fi338 < 6; fi338++ {
                        for fi339 := 0; fi339 < 12; fi339++ {
                            v3 = true
                            _ = fi339
                        }
                        var v340 int16 = -32768
                        v341 := GenS33{F0: 0}
                        _ = fi338
                        _ = v340
                        _ = v341
                    }
                    for fi342 := 0; fi342 < 20; fi342++ {
                        if v3 {
                            v343 := GenS4{F0: fi342, F1: 1, F2: 1}
                            for fi344 := 0; fi344 < 3; fi344++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v345 bool = v3
                                _ = fi344
                                _ = v345
                            }
                            var v346 bool = v3
                            _ = v343
                            _ = v346
                        }
                        var v347 float32 = 0.1
                        _ = fi342
                        _ = v347
                    }
                    v348 := GenS1{F0: "\n\t"}
                    _ = v348
                }
            default:
                for fi349 := 0; fi349 < 12; fi349++ {
                    v350 := GenS34{F0: -128, F1: 4294967295, F2: "foo bar"}
                    v351 := GenS18{F0: 42, F1: "hello", F2: 1}
                    var v352 uint16 = 32767
                    _ = fi349
                    _ = v350
                    _ = v351
                    _ = v352
                }
            }
            if (!v3) {
                if v3 {
                    for fi353 := 0; fi353 < 14; fi353++ {
                        var v354 uint64 = 956
                        _ = fi353
                        _ = v354
                    }
                    var v355 uint = 732
                    _ = v355
                } else {
                    fmt.Println(v3)
                }
            }
        }
        switch "\n\t" {
        case "\n\t":
            v3 = (v3 || false)
        case "":
            fmt.Println(v3)
        default:
            for fi356 := 0; fi356 < 11; fi356++ {
                for fi357 := 0; fi357 < 13; fi357++ {
                    v2 = append(v2, "\n\t")
                    if v3 {
                        if v3 {
                            fmt.Println(fi357, v3, v3)
                            var v358 uint64 = 249
                            _ = v358
                        }
                        if true {
                            v3 = (0.0 >= ((float32(0.1) / ((0.0 - 1.5) + 0.0)) - 0.0))
                            v3 = (1 == 99)
                            switch "foo bar" {
                            case "hello":
                                v2 = append(v2, "")
                            case "\n\t":
                                var v359 uint64 = 18446744073709551615
                                _ = v359
                            case "":
                                var v360 uint8 = 255
                                _ = v360
                            default:
                                if v3 {
                                    v3 = v3
                                    var v361 uint64 = 42
                                    v4 = append(v4, 1)
                                    v361 -= 1
                                    switch 9223372036854775807 {
                                    case 648:
                                        if v3 {
                                            switch 0 {
                                            case -311:
                                                var v362 string = "\n\t"
                                                _ = v362
                                            case -390:
                                                _ = len(v1)
                                            }
                                            v361 = (^(v361))
                                        }
                                    }
                                    fmt.Println(v3, v3)
                                    _ = v361
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = 450
                                    }
                                }
                            }
                        } else {
                            v1 = append(v1, 65535)
                            var v363 uint8 = 36
                            _ = v363
                        }
                        v3 = (1 == 1)
                    }
                    if (806 != 1) {
                        var v364 uint = 1
                        var v365 float64 = 0.1
                        var v366 bool = v3
                        var v367 uint32 = 1
                        _ = v364
                        _ = v365
                        _ = v366
                        _ = v367
                    }
                    _ = fi357
                }
                for fi368 := 0; fi368 < 11; fi368++ {
                    v3 = true
                    _ = fi368
                }
                var v369 int16 = -32768
                _ = fi356
                _ = v369
            }
        }
    }
    v370 := GenS35{F0: "a", F1: -93, F2: ""}
    for fi371 := 0; fi371 < 5; fi371++ {
        var v372 int16 = 1000
        var v373 float32 = (1.5 * 1.5)
        _ = fi371
        _ = v372
        _ = v373
    }
    if (((!true) && v3) || (1 == 4294967295)) {
        if (42 < 0) {
            v3 = (("\n\t" + "foo bar") == ("\n\t" + "a"))
        }
        v3 = v3
    }
    switch "\n\t" {
    case "foo bar":
        switch 4294967295 {
        case -202:
            v3 = v3
        case 127:
            switch ("hello" + "a") {
            case "\n\t":
                for fi374 := 0; fi374 < 13; fi374++ {
                    for fi375 := 0; fi375 < 8; fi375++ {
                        v3 = false
                        _ = fi375
                    }
                    var v376 bool = v3
                    v377 := GenS36{F0: -1}
                    _ = fi374
                    _ = v376
                    _ = v377
                }
            case "":
                for fi378 := 0; fi378 < 3; fi378++ {
                    var v379 float32 = (0.0 / 1e10)
                    _ = fi378
                    _ = v379
                }
            case "foo bar":
                var v380 uint64 = 363
                _ = v380
            default:
                fmt.Println(v3)
            }
        case -1:
            var v381 int8 = 1
            _ = v381
        default:
            v3 = true
        }
    }
    var v382 uint64 = 42
    v382 -= v382
    if v3 {
        var v383 int = -436
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v383
    } else {
        fmt.Println(v3)
        fmt.Println(v3, v382)
        var v384 uint64 = 449
        _ = v384
    }
    for fi385 := 0; fi385 < 7; fi385++ {
        switch fi385 {
        case 0:
            v386 := GenS20{F0: v3, F1: v3}
            _ = v386
        case 4294967295:
            var v387 int = (fi385 / ((-152) | 1))
            _ = v387
        case -2147483648:
            var v388 uint = 541
            _ = v388
        default:
            var v389 uint64 = (v382 - (^(v382)))
            _ = v389
        }
        fmt.Println(v3, v3, fi385)
        _ = fi385
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v382:", v382)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v259
    _ = v3
    _ = v370
    _ = v382
    _ = v4
}
