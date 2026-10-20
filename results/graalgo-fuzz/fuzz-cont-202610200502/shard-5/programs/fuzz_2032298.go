package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int; F2 int }
type GenS4 struct { F0 bool; F1 int16; F2 uint64 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 bool = false
    var v2 int = 978
    var v3 int32 = -2147483648
    var v4 float32 = (-2.25 + 0.1)
    var v5 int64 = 42
    v6 := []int{v2, v2, (v2 + (-(v2))), -9223372036854775808}
    if false {
        var v7 float64 = (3.14159 - (1e10 + (1.5 * 1.5)))
        _ = v7
    } else {
        if true {
            if v1 {
                v8 := GenS1{F0: v2, F1: 47}
                v5 = (v5 * (v5 << ((255) & 63)))
                _ = v8
            } else {
                v1 = (("a" + (("a" + "") + "\n\t")) >= "")
                for fi9 := 0; fi9 < 9; fi9++ {
                    if (((v1 || v1) || true) || (0 <= 949)) {
                        v3 -= -2147483648
                        _ = len(v6)
                    }
                    _ = fi9
                }
                var v10 uint8 = 2
                _ = v10
            }
        }
    }
    if (((0 >= 1) || (4294967295 < 639)) || v1) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(v5)
        v11 := GenS2{F0: 42}
        _ = v11
    }
    if v1 {
        var v12 uint8 = 0
        var v13 uint64 = 2
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        if false {
            var v15 uint = 4294967295
            if true {
                v4 += (1e10 / 1.0)
                fmt.Println(v3)
            }
            if len(v6) > 0 {
                v6[0] = (^(v2))
            }
            _ = v15
        } else {
            var v16 int32 = v3
            v16 = (v3 * v16)
            v1 = true
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            v3 = (v3 / (((v3 & -2147483648)) | 1))
            v2 = fi17
            v1 = false
            _ = fi17
        }
        fmt.Println(v4)
        _ = fi14
    }
    var v18 uint8 = 255
    v19 := GenS2{F0: (v18 << ((2) & 63))}
    fmt.Println(v1, v5, v1)
    v20 := GenS3{F0: 42, F1: v2, F2: v2}
    var v21 int32 = (v3 - v3)
    for fi22 := 0; fi22 < 11; fi22++ {
        if v1 {
            if v1 {
                fmt.Println(fi22)
                v4 -= v4
                if len(v6) > 0 {
                    v6[0] = (fi22 / (((v2 | (v2 - -2147483648))) | 1))
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v4 -= 0.1
            }
            v21 = v21
        }
        fmt.Println(v1, v18, fi22)
        _ = fi22
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v21 -= v21
    for fi23 := 0; fi23 < 12; fi23++ {
        switch "foo bar" {
        case "foo bar":
            if false {
                for fi24 := 0; fi24 < 16; fi24++ {
                    if (-1 > (-(v5))) {
                        v2 = fi24
                        var v25 int = fi24
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v26 bool = true
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(fi24, v18)
                    switch (fi24 - v2) {
                    case -786:
                        if v1 {
                            var v27 int16 = -58
                            if v1 {
                                v21 = int32(v27)
                                switch ((("hello" + "\n\t") + "") + "a") {
                                case "\n\t":
                                    var v28 float32 = ((v4 - v4) / (float32(1.5) * v4))
                                    _ = v28
                                default:
                                    var v29 int64 = (-(v5))
                                    _ = v29
                                }
                                var v30 bool = true
                                _ = v30
                            }
                            var v31 int = fi24
                            _ = v27
                            _ = v31
                        }
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v1)
                        }
                    }
                    v4 = float32(-2.25)
                    if v1 {
                        if (true || (((3.14159 * 1.5) + 1e10) >= 1e10)) {
                            switch "" {
                            case "":
                                fmt.Println(v4, v5, fi24)
                            }
                        }
                    }
                    var v32 uint = 42
                    _ = fi24
                    _ = v32
                }
            }
        case "hello":
            if (1 != int8(v21)) {
                if v1 {
                    for fi33 := 0; fi33 < 9; fi33++ {
                        v34 := GenS3{F0: 1000, F1: fi33, F2: fi33}
                        _ = fi33
                        _ = v34
                    }
                } else {
                    switch ("\n\t" + "a") {
                    case "hello":
                        for fi35 := 0; fi35 < 4; fi35++ {
                            var v36 int32 = -2147483648
                            v1 = true
                            _ = fi35
                            _ = v36
                        }
                    case "":
                        v1 = v1
                    case "foo bar":
                        if (!v1) {
                            for fi37 := 0; fi37 < 2; fi37++ {
                                if (2 > 2) {
                                    if v1 {
                                        var v38 int8 = 42
                                        _ = v38
                                    }
                                    v6 = append(v6, (fi23 ^ 0))
                                    switch v2 {
                                    case 32767:
                                        fmt.Println(v2, fi23, v4)
                                    case 655:
                                        if (!v1) {
                                            switch v2 {
                                            case 9223372036854775807:
                                                var v39 int8 = -128
                                                _ = v39
                                            case 32768:
                                                if v1 {
                                                    if v1 {
                                                        var v40 int = fi37
                                                        if ((65535 >= 209) || false) {
                                                            if (!v1) {
                                                                var v41 float32 = v4
                                                                var v42 string = "hello"
                                                                fmt.Println(v40, v2, v3)
                                                                var v43 bool = false
                                                                _ = v41
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                var v44 uint16 = 345
                                                                v40 = v2
                                                                if (v44 > (v44 << ((587) & 63))) {
                                                                    switch "" {
                                                                    case "hello":
                                                                        var v45 uint = 520
                                                                        _ = v45
                                                                    }
                                                                    var v46 bool = v1
                                                                    v5 += (v5 << ((65535) & 63))
                                                                    v40 -= int(v18)
                                                                    _ = v46
                                                                } else {
                                                                    v3 -= 2147483647
                                                                    var v47 int = 657
                                                                    v48 := GenS3{F0: 1, F1: v40, F2: (fi37 >> ((0) & 63))}
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                                _ = v44
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = (fi23 << ((2) & 63))
                                                            }
                                                            var v49 bool = (v1 || true)
                                                            _ = v49
                                                        }
                                                        _ = v40
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                            case -2147483648:
                                                v6 = append(v6, (-(fi23)))
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v21 = (v21 & v21)
                                                case "\n\t":
                                                    if v1 {
                                                        if (18446744073709551615 >= 255) {
                                                            v50 := GenS2{F0: 128}
                                                            v3 = v3
                                                            v51 := GenS4{F0: v1, F1: 37, F2: 300}
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    } else {
                                                        var v52 bool = ((v1 && false) && true)
                                                        switch -9223372036854775808 {
                                                        case -154:
                                                            v53 := GenS3{F0: 118, F1: (^(v2)), F2: -840}
                                                            _ = v53
                                                        case 42:
                                                            fmt.Println(v3)
                                                        default:
                                                            if (255 != 32767) {
                                                                v2 = (^(fi23))
                                                            }
                                                        }
                                                        v6 = append(v6, (fi37 << ((216) & 63)))
                                                        var v54 float32 = float32((float64(float32((float64(((v4 - ((v4 + (v4 - v4)) / 1.0)) * ((0.0 * v4) + 3.14159))) / 0.1))) * 0.1))
                                                        var v55 string = ("a" + "\n\t")
                                                        _ = v52
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                }
                                            }
                                            fmt.Println(v21, v21)
                                        } else {
                                            if v1 {
                                                switch "foo bar" {
                                                case "hello":
                                                    var v56 int16 = 0
                                                    _ = v56
                                                case "":
                                                    v1 = v1
                                                    if true {
                                                        v6 = append(v6, v2)
                                                        if (v21 != v3) {
                                                            _ = len(v6)
                                                            v2 += (fi37 % ((1) | 1))
                                                            v57 := GenS2{F0: 224}
                                                            _ = v57
                                                        } else {
                                                            if (v1 || v1) {
                                                                v58 := GenS3{F0: -32768, F1: (fi37 + 32768), F2: 0}
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    var v59 string = ""
                                                                    _ = v59
                                                                }
                                                                _ = v58
                                                            }
                                                            _ = len(v6)
                                                            var v60 uint32 = 233
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                            }
                                            v61 := GenS2{F0: v18}
                                            _ = v61
                                        }
                                    case 2147483647:
                                        v62 := GenS1{F0: (fi23 + v2), F1: 4294967295}
                                        _ = v62
                                    }
                                } else {
                                    v5 = (v5 % ((-1) | 1))
                                    fmt.Println(v1, v18, v4)
                                    if (("foo bar" + "") >= "\n\t") {
                                        v4 = 3.14159
                                        var v63 uint8 = v18
                                        v64 := GenS1{F0: (fi23 % (((fi37 * (^(fi37)))) | 1)), F1: 1}
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                if len(v6) > 0 {
                                    v6[0] = fi23
                                }
                                v65 := GenS3{F0: -32768, F1: (v2 - (-(fi23))), F2: (^(fi23))}
                                _ = fi37
                                _ = v65
                            }
                            v18 -= 2
                        } else {
                            var v66 uint64 = 759
                            v5 = v5
                            _ = v66
                        }
                    default:
                        var v67 bool = (v2 == fi23)
                        _ = v67
                    }
                    v6 = append(v6, -358)
                    switch ("a" + ("" + "\n\t")) {
                    case "":
                        _ = len(v6)
                    case "hello":
                        switch (fi23 & 9223372036854775807) {
                        case 32767:
                            var v68 int64 = 1
                            _ = v68
                        case -1:
                            if true {
                                if ("hello" <= "hello") {
                                    for fi69 := 0; fi69 < 14; fi69++ {
                                        switch 9223372036854775807 {
                                        case 32767:
                                            _ = len(v6)
                                        case -2147483648:
                                            switch ("" + ("hello" + ("foo bar" + ("hello" + "\n\t")))) {
                                            case "foo bar":
                                                if (v1 && false) {
                                                    v70 := GenS1{F0: fi23, F1: 18446744073709551615}
                                                    var v71 string = "foo bar"
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    var v72 bool = ((fi23 < (v2 & (v2 >> ((191) & 63)))) || (false || false))
                                                    fmt.Println(v3, v72)
                                                    _ = v72
                                                }
                                            default:
                                                if v1 {
                                                    var v73 int8 = -13
                                                    fmt.Println(v18)
                                                    switch -1 {
                                                    case 32768:
                                                        if (true && true) {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            switch (-(v2)) {
                                                            case 128:
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            case 967:
                                                                if v1 {
                                                                    v3 -= -49
                                                                } else {
                                                                    fmt.Println(v4, fi23, v3)
                                                                }
                                                            }
                                                            v74 := GenS3{F0: 0, F1: (^(v2)), F2: (-(fi23))}
                                                            _ = v74
                                                        } else {
                                                            fmt.Println(v18)
                                                            v75 := GenS1{F0: (^(v2)), F1: 1}
                                                            if len(v6) > 0 {
                                                                v6[0] = (fi69 + (v2 - (-(v2))))
                                                            }
                                                            if v1 {
                                                                v76 := GenS4{F0: (v1 && ((true || v1) && (v1 || v1))), F1: 525, F2: 237}
                                                                v1 = false
                                                                _ = v76
                                                            } else {
                                                                var v77 int8 = v73
                                                                if v1 {
                                                                    v6 = append(v6, v2)
                                                                    v78 := GenS1{F0: fi69, F1: 2}
                                                                    v6 = append(v6, 9223372036854775807)
                                                                    _ = v78
                                                                } else {
                                                                    var v79 int32 = (v21 / (((v21 / ((1) | 1))) | 1))
                                                                    fmt.Println(v79, v2, v1)
                                                                    _ = v79
                                                                }
                                                                _ = v77
                                                            }
                                                            fmt.Println(v2, v18, v21)
                                                            _ = v75
                                                        }
                                                    case 741:
                                                        v80 := GenS3{F0: 1000, F1: 9223372036854775807, F2: fi23}
                                                        _ = v80
                                                    }
                                                    _ = v73
                                                }
                                            }
                                        default:
                                            v4 = 0.0
                                        }
                                        _ = fi69
                                    }
                                    for fi81 := 0; fi81 < 3; fi81++ {
                                        v6 = append(v6, v2)
                                        fmt.Println(v5, fi23, v2)
                                        if v1 {
                                            if true {
                                                if (v4 != (v4 * v4)) {
                                                    v82 := GenS2{F0: v18}
                                                    var v83 uint = 1
                                                    _ = v82
                                                    _ = v83
                                                }
                                            } else {
                                                v3 -= v3
                                                var v84 int32 = (v3 / ((v21) | 1))
                                                _ = v84
                                            }
                                            var v85 float64 = (3.14159 + (1e10 - -2.25))
                                            if v1 {
                                                fmt.Println(v21)
                                                if v1 {
                                                    v5 = 42
                                                    v4 = float32((float64(3.14159) / (v85 - 3.14159)))
                                                    fmt.Println(v2)
                                                }
                                                v85 += ((1.5 / 0.1) * 0.1)
                                            } else {
                                                fmt.Println(v21, v18)
                                                v86 := GenS3{F0: 1000, F1: (fi81 ^ -1), F2: fi81}
                                                v6 = append(v6, fi81)
                                                _ = v86
                                            }
                                            _ = v85
                                        }
                                        _ = fi81
                                    }
                                }
                                v18 += v18
                            }
                        case -857:
                            var v87 uint = 255
                            _ = v87
                        }
                    case "\n\t":
                        switch "" {
                        case "foo bar":
                            if v1 {
                                v4 = v4
                                switch 1 {
                                case 42:
                                    v88 := GenS4{F0: (-1 <= (v5 - v5)), F1: int16(v3), F2: 0}
                                    _ = v88
                                default:
                                    if v1 {
                                        v6 = append(v6, int(v5))
                                    }
                                }
                                v21 = v3
                            } else {
                                switch (^(fi23)) {
                                case -113:
                                    v89 := GenS4{F0: ((v5 == v5) && v1), F1: 0, F2: 0}
                                    _ = v89
                                }
                                v1 = (0 < 127)
                            }
                        case "hello":
                            v3 = int32(v18)
                        }
                    default:
                        var v90 uint16 = 255
                        _ = v90
                    }
                }
                switch -457 {
                case 9223372036854775807:
                    var v91 int64 = v5
                    _ = v91
                }
                v18 = (^(v18))
            }
        case "a":
            v2 = (fi23 | (fi23 & (fi23 % ((fi23) | 1))))
        }
        switch v2 {
        case 0:
            v21 = (v3 >> ((0) & 63))
        case 32768:
            for fi92 := 0; fi92 < 15; fi92++ {
                if len(v6) > 0 {
                    v6[0] = fi23
                }
                var v93 int16 = 42
                v2 += v2
                var v94 int = (fi92 / ((v2) | 1))
                _ = fi92
                _ = v93
                _ = v94
            }
        case -9223372036854775808:
            switch ("a" + "hello") {
            case "\n\t":
                for fi95 := 0; fi95 < 14; fi95++ {
                    var v96 bool = (42 < -39)
                    _ = fi95
                    _ = v96
                }
            case "hello":
                if len(v6) > 0 {
                    v6[0] = (v2 ^ (v2 + (v2 / ((fi23) | 1))))
                }
            case "foo bar":
                var v97 int = -964
                _ = v97
            default:
                var v98 uint32 = 277
                _ = v98
            }
        }
        for fi99 := 0; fi99 < 19; fi99++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi99
        }
        _ = fi23
    }
    if false {
        v1 = false
        var v100 int64 = v5
        _ = v100
    }
    v101 := GenS5{F0: -1}
    if (v1 || (v2 < 42)) {
        var v102 uint16 = 2
        v4 -= float32(3.14159)
        _ = v102
    } else {
        v2 = 128
    }
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v101
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
