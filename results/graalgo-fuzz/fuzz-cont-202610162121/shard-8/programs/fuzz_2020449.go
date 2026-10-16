package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 int32; F1 uint64 }
type GenS11 struct { F0 uint8; F1 string }
type GenS12 struct { F0 int32; F1 int }
type GenS13 struct { F0 int32; F1 int8; F2 uint16 }
type GenS14 struct { F0 string; F1 int8; F2 string }
type GenS15 struct { F0 int8; F1 int8; F2 int }
type GenS16 struct { F0 int32; F1 uint64 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint64; F1 string; F2 uint64 }
type GenS19 struct { F0 string; F1 bool; F2 string }
type GenS20 struct { F0 int16; F1 uint32; F2 int16 }
type GenS21 struct { F0 int32; F1 uint8; F2 uint64 }
func genFn0(p0 uint8) int32 {
    if (18446744073709551615 > 65535) {
        switch -656 {
        case 0:
            p0 = (^(p0))
        }
    }
    var v4 uint32 = 42
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = (p0 % (((p0 - 128)) | 1))
        v7 := []int{(fi5 << ((0) & 63)), fi5, fi5, (^(fi5))}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 string) int8 {
    v8 := GenS1{F0: 1}
    v9 := []int{2147483647, 2147483647}
    p0 = p0
    _ = p0
    _ = v8
    _ = v9
    return -128
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    var v3 string = (v2 + v2)
    v3 = "hello"
    v10 := GenS1{F0: 4294967295}
    switch "\n\t" {
    case "\n\t":
        switch ("a" + v2) {
        case "\n\t":
            v2 = "hello"
        case "a":
            var v11 uint32 = 0
            _ = v11
        case "foo bar":
            if (((1e10 / (0.1 + 3.14159)) < (3.14159 - float32(0.0))) && (-2147483648 < -1)) {
                v3 = (v2 + v2)
                v3 = (v3 + (("\n\t" + v2) + ("a" + ("hello" + "a"))))
                switch 187 {
                case 128:
                    var v12 uint = 1
                    _ = v12
                default:
                    if ((!(128 >= 1)) || true) {
                        var v13 uint64 = 270
                        _ = len(v1)
                        _ = v13
                    }
                }
            } else {
                var v14 int8 = 116
                var v15 int64 = -1
                _ = v14
                _ = v15
            }
        default:
            if true {
                switch "a" {
                case "":
                    v16 := GenS2{F0: 0}
                    _ = v16
                }
                switch 971 {
                case 32767:
                    fmt.Println(v2, v2)
                case 4294967295:
                    if (-1 >= genFn0(127)) {
                        v17 := GenS2{F0: 435}
                        v3 = "a"
                        switch 2147483647 {
                        case -1:
                            fmt.Println(v2, v3)
                        case 9223372036854775807:
                            for fi18 := 0; fi18 < 17; fi18++ {
                                v2 = ("\n\t" + (v2 + v2))
                                var v19 int32 = -724
                                _ = fi18
                                _ = v19
                            }
                        case 683:
                            var v20 int8 = 42
                            _ = v20
                        default:
                            v21 := v1[2147483647]
                            _ = v21
                        }
                        _ = v17
                    }
                }
            }
        }
    case "":
        for fi22 := 0; fi22 < 19; fi22++ {
            if false {
                if ((!(!(true && true))) && true) {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v3 = v3
                        _ = len(v1)
                        var v24 int64 = -1
                        _ = fi23
                        _ = v24
                    }
                    v25 := []int{}
                    if true {
                        var v26 uint32 = 968
                        switch fi22 {
                        case 32767:
                            if false {
                                fmt.Println(v3)
                                if false {
                                    if (fi22 >= (fi22 % ((fi22) | 1))) {
                                        for fi27 := 0; fi27 < 10; fi27++ {
                                            switch v3 {
                                            case "a":
                                                if (!false) {
                                                    switch v2 {
                                                    case "foo bar":
                                                        v26 -= v26
                                                        fmt.Println(v2)
                                                        v28 := GenS2{F0: 42}
                                                        _ = v28
                                                    case "hello":
                                                        v3 = v3
                                                    case "\n\t":
                                                        if false {
                                                            var v29 string = "a"
                                                            v29 = v29
                                                            _ = v29
                                                        }
                                                    default:
                                                        if false {
                                                            fmt.Println(fi27)
                                                            v3 = v3
                                                        }
                                                    }
                                                    v26 = (v26 | 4294967295)
                                                } else {
                                                    v2 = (((v3 + v2) + v2) + ((("foo bar" + "a") + ("" + "a")) + v2))
                                                }
                                            case "hello":
                                                v30 := GenS2{F0: 0}
                                                _ = v30
                                            default:
                                                v2 = (v3 + ("hello" + v3))
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                v3 = "foo bar"
                                            case "foo bar":
                                                var v31 uint32 = (v26 ^ v26)
                                                _ = v31
                                            }
                                            switch "hello" {
                                            case "":
                                                var v32 bool = false
                                                _ = v32
                                            case "a":
                                                var v33 float64 = (-2.25 * 1.5)
                                                _ = v33
                                            case "hello":
                                                switch 600 {
                                                case 32768:
                                                    v3 = v3
                                                case 42:
                                                    v2 = v2
                                                }
                                            default:
                                                v2 = ("\n\t" + "a")
                                            }
                                            _ = fi27
                                        }
                                        fmt.Println(v2, v3)
                                        v34 := GenS3{F0: v26}
                                        _ = v34
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        switch v3 {
                                        case "hello":
                                            var v35 int8 = genFn1(v3)
                                            _ = v35
                                        default:
                                            delete(v1, 32767)
                                        }
                                    case "":
                                        for fi36 := 0; fi36 < 2; fi36++ {
                                            switch "a" {
                                            case "\n\t":
                                                if (false || ((v26 + v26) < 0)) {
                                                    if (0.1 == 3.14159) {
                                                        v25 = append(v25, (fi22 >> ((255) & 63)))
                                                        var v37 bool = true
                                                        _ = v37
                                                    } else {
                                                        v25 = append(v25, 128)
                                                        v3 = "foo bar"
                                                        v2 = v3
                                                    }
                                                } else {
                                                    v3 = v2
                                                    if true {
                                                        _ = len(v25)
                                                        var v38 uint64 = 0
                                                        _ = v38
                                                    } else {
                                                        if (fi36 > (fi22 % (((fi36 >> ((18446744073709551615) & 63))) | 1))) {
                                                            var v39 int16 = 1000
                                                            switch 42 {
                                                            case 9223372036854775807:
                                                                if (fi36 != (-(fi22))) {
                                                                    var v40 bool = (0.1 > (0.1 - 1e10))
                                                                    fmt.Println(fi22, v3)
                                                                    _ = v40
                                                                } else {
                                                                    var v41 string = "\n\t"
                                                                    var v42 int = fi36
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                            case 128:
                                                                var v43 int32 = 279
                                                                _ = v43
                                                            case -1:
                                                                if false {
                                                                    v26 = v26
                                                                    var v44 float32 = 1e10
                                                                    _ = v44
                                                                } else {
                                                                    v26 = 707
                                                                    v45 := v1[(^(fi22))]
                                                                    var v46 uint = 1
                                                                    switch ((("hello" + v3) + "") + v2) {
                                                                    case "hello":
                                                                        v3 = "foo bar"
                                                                    case "a":
                                                                        if len(v25) > 0 {
                                                                            v25[0] = fi22
                                                                        }
                                                                        if len(v25) > 0 {
                                                                            v25[0] = fi22
                                                                        }
                                                                        var v47 int32 = 2147483647
                                                                        _ = v47
                                                                    case "":
                                                                        v48 := GenS1{F0: (^(v46))}
                                                                        _ = v48
                                                                    }
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                            default:
                                                                v3 = v2
                                                                fmt.Println(fi36)
                                                                switch v2 {
                                                                case "\n\t":
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v49 string = (v3 + v2)
                                                                        _ = v49
                                                                    case "a":
                                                                        if false {
                                                                            var v50 uint16 = 852
                                                                            switch (fi36 * fi36) {
                                                                            case 825:
                                                                                fmt.Println(fi36, v50)
                                                                            case 42:
                                                                                v26 += v26
                                                                            default:
                                                                                v50 = 42
                                                                            }
                                                                            var v51 int = (-(fi22))
                                                                            _ = v50
                                                                            _ = v51
                                                                        } else {
                                                                            var v52 int64 = 1
                                                                            v39 -= v39
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v53 uint16 = 1
                                                                    v53 = (v53 & v53)
                                                                    _ = v53
                                                                case "":
                                                                    var v54 int64 = 0
                                                                    _ = v54
                                                                default:
                                                                    switch v2 {
                                                                    case "foo bar":
                                                                        if ((255 >= 255) && (356 <= 9223372036854775807)) {
                                                                            if false {
                                                                                v2 = (v3 + (("\n\t" + "foo bar") + (("foo bar" + v2) + ("foo bar" + v2))))
                                                                                v55 := GenS4{F0: 0}
                                                                                if len(v25) >= 2 {
                                                                                    v25 = v25[1:len(v25)]
                                                                                }
                                                                                _ = v55
                                                                            } else {
                                                                                v56 := GenS5{F0: 209, F1: v26}
                                                                                if (!(65535 < 852)) {
                                                                                    fmt.Println(v39, fi36)
                                                                                    v57 := GenS3{F0: (v26 - (v26 * 2147483648))}
                                                                                    var v58 int64 = 9223372036854775807
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                                _ = v56
                                                                            }
                                                                            v59 := GenS2{F0: 42}
                                                                            var v60 uint16 = 65535
                                                                            _ = v59
                                                                            _ = v60
                                                                        } else {
                                                                            v61 := GenS6{F0: 42}
                                                                            v62 := GenS4{F0: 40}
                                                                            var v63 uint64 = 42
                                                                            _ = v61
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                        var v64 int16 = (v39 - int16(v26))
                                                                        v65 := GenS4{F0: 65535}
                                                                        _ = v64
                                                                        _ = v65
                                                                    case "hello":
                                                                        if false {
                                                                            switch fi22 {
                                                                            case 9223372036854775807:
                                                                                if false {
                                                                                    if false {
                                                                                        v26 += 4294967295
                                                                                        v2 = v2
                                                                                        switch (("\n\t" + "hello") + ("a" + v3)) {
                                                                                        case "\n\t":
                                                                                            if (!false) {
                                                                                                v2 = (((v2 + ("hello" + v3)) + "") + ((v3 + v3) + (v3 + "")))
                                                                                                if (1e10 == 1e10) {
                                                                                                    v3 = ""
                                                                                                    v66 := GenS2{F0: -2147483648}
                                                                                                    _ = v66
                                                                                                }
                                                                                                v2 = v2
                                                                                            }
                                                                                        default:
                                                                                            if (v26 > 632) {
                                                                                                var v67 float32 = (0.1 / 1.0)
                                                                                                _ = v67
                                                                                            } else {
                                                                                                var v68 uint32 = (^(v26))
                                                                                                _ = v68
                                                                                            }
                                                                                            if (127 > 42) {
                                                                                                v3 = "hello"
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 413:
                                                                                if ((false || false) || ((0.1 > (1.5 + ((1e10 - float32((0.0 - 1e10))) + (0.1 * (0.1 - ((1e10 / -2.25) - float32(1.5))))))) || ((127 != -78) || true))) {
                                                                                    v2 = (v3 + "\n\t")
                                                                                    v69 := GenS1{F0: 1}
                                                                                    var v70 int16 = 32767
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                } else {
                                                                                    v3 = ("hello" + v2)
                                                                                    if false {
                                                                                        v26 = 42
                                                                                    }
                                                                                    var v71 int8 = 1
                                                                                    _ = v71
                                                                                }
                                                                            case -2147483648:
                                                                                switch v3 {
                                                                                case "":
                                                                                    var v72 string = v2
                                                                                    _ = v72
                                                                                case "\n\t":
                                                                                    switch fi22 {
                                                                                    case 127:
                                                                                        fmt.Println(v2, v26)
                                                                                    }
                                                                                default:
                                                                                    var v73 float32 = ((-2.25 - 1e10) / 0.1)
                                                                                    _ = v73
                                                                                }
                                                                            default:
                                                                                switch v3 {
                                                                                case "hello":
                                                                                    v3 = "\n\t"
                                                                                case "\n\t":
                                                                                    v74 := GenS3{F0: (^(v26))}
                                                                                    _ = v74
                                                                                default:
                                                                                    v3 = ("hello" + ("hello" + v2))
                                                                                }
                                                                            }
                                                                            var v75 int8 = -128
                                                                            _ = v75
                                                                        }
                                                                    case "a":
                                                                        if (float32(float64((3.14159 - ((1e10 / (((0.0 * 0.1) - (1.5 - 1e10)) * -2.25)) * (3.14159 + 1e10))))) == 1.5) {
                                                                            v3 = v3
                                                                            var v76 uint16 = 65535
                                                                            v25 = append(v25, fi36)
                                                                            _ = v76
                                                                        } else {
                                                                            switch (v3 + (("\n\t" + v3) + v3)) {
                                                                            case "hello":
                                                                                switch (fi22 / ((fi22) | 1)) {
                                                                                case -687:
                                                                                    var v77 int8 = genFn1("hello")
                                                                                    _ = v77
                                                                                case 2147483647:
                                                                                    if (true || ((("hello" + ("\n\t" + "hello")) + v2) == (v3 + v3))) {
                                                                                        v78 := GenS4{F0: 65535}
                                                                                        var v79 int32 = -620
                                                                                        var v80 int8 = int8(fi36)
                                                                                        _ = v78
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    }
                                                                                default:
                                                                                    switch 0 {
                                                                                    case 2147483647:
                                                                                        _ = len(v25)
                                                                                    case -902:
                                                                                        var v81 int32 = 2147483647
                                                                                        _ = v81
                                                                                    default:
                                                                                        if true {
                                                                                            if ((0.1 + (float64(float32(1.5)) - -2.25)) >= (3.14159 + 3.14159)) {
                                                                                                v2 = v2
                                                                                            } else {
                                                                                                var v82 int32 = genFn0(0)
                                                                                                _ = v82
                                                                                            }
                                                                                            v2 = (v3 + v3)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case "":
                                                                                if true {
                                                                                    if true {
                                                                                        if (9223372036854775807 >= -1) {
                                                                                            v83 := GenS5{F0: 337, F1: (v26 % (((^(v26))) | 1))}
                                                                                            _ = v83
                                                                                        } else {
                                                                                            v2 = (("hello" + "foo bar") + v2)
                                                                                            v84 := GenS2{F0: genFn0(164)}
                                                                                            var v85 uint64 = 2
                                                                                            v86 := v1[fi36]
                                                                                            v1[(fi22 + (^(fi36)))] = (v86 + ("foo bar" + "hello"))
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        v87 := GenS3{F0: 65535}
                                                                                        if ((1 >= 798) || true) {
                                                                                            if (-1 < 2147483647) {
                                                                                                if true {
                                                                                                    var v88 uint32 = (v26 & v26)
                                                                                                    v89 := GenS7{F0: (v39 | v39), F1: 273, F2: 4294967295}
                                                                                                    if (571 > 42) {
                                                                                                        switch fi36 {
                                                                                                        case -9223372036854775808:
                                                                                                            var v90 uint32 = (^(v88))
                                                                                                            _ = v90
                                                                                                        case -1:
                                                                                                            v26 = (v88 << ((18446744073709551615) & 63))
                                                                                                        case 397:
                                                                                                            v39 = v39
                                                                                                        }
                                                                                                    }
                                                                                                    var v91 int16 = (v39 % ((v39) | 1))
                                                                                                    var v92 uint32 = (v88 | v26)
                                                                                                    if (-120 != 100) {
                                                                                                        fmt.Println(v91)
                                                                                                        var v93 int16 = (v39 | (v39 * v39))
                                                                                                        _ = v93
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            fmt.Println(fi36, v2, v91)
                                                                                                        }
                                                                                                        var v94 int64 = 938
                                                                                                        var v95 int64 = v94
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            var v96 uint32 = v26
                                                                                                            _ = v96
                                                                                                        case "hello":
                                                                                                            fmt.Println(v3, fi22, v95)
                                                                                                        case "\n\t":
                                                                                                            var v97 int32 = 2147483647
                                                                                                            _ = v97
                                                                                                        default:
                                                                                                            v3 = v3
                                                                                                        }
                                                                                                        v39 -= v39
                                                                                                        var v98 int = fi22
                                                                                                        _ = v94
                                                                                                        _ = v95
                                                                                                        _ = v98
                                                                                                    }
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                    _ = v91
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    switch ("\n\t" + v2) {
                                                                                                    case "\n\t":
                                                                                                        v26 = v26
                                                                                                    case "":
                                                                                                        var v99 float32 = 1e10
                                                                                                        fmt.Println(v99, v39)
                                                                                                        if false {
                                                                                                            v100 := v1[53]
                                                                                                            var v101 uint = 1
                                                                                                            v102 := GenS4{F0: 221}
                                                                                                            _ = v100
                                                                                                            _ = v101
                                                                                                            _ = v102
                                                                                                        } else {
                                                                                                            if (-2147483648 >= genFn0(0)) {
                                                                                                                var v103 uint32 = 2147483648
                                                                                                                fmt.Println(fi22, v103)
                                                                                                                _ = v103
                                                                                                            } else {
                                                                                                                var v104 string = v2
                                                                                                                v25 = append(v25, 606)
                                                                                                                _ = v104
                                                                                                            }
                                                                                                            var v105 bool = true
                                                                                                            _ = v105
                                                                                                        }
                                                                                                        _ = v99
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(fi22, v39)
                                                                                                    default:
                                                                                                        var v106 uint64 = 1
                                                                                                        var v107 bool = (565 >= int32(v26))
                                                                                                        if len(v25) >= 2 {
                                                                                                            v25 = v25[1:len(v25)]
                                                                                                        }
                                                                                                        _ = v106
                                                                                                        _ = v107
                                                                                                    }
                                                                                                    v26 -= 355
                                                                                                    v25 = append(v25, 42)
                                                                                                }
                                                                                                if (false && (v39 < v39)) {
                                                                                                    var v108 float64 = 0.1
                                                                                                    if true {
                                                                                                        if (int32(v39) != -112) {
                                                                                                            var v109 int = fi22
                                                                                                            var v110 uint64 = 1
                                                                                                            delete(v1, 0)
                                                                                                            _ = v109
                                                                                                            _ = v110
                                                                                                        }
                                                                                                        v39 = v39
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                v2 = (v2 + v3)
                                                                                                                v2 = v3
                                                                                                                switch fi36 {
                                                                                                                case 32767:
                                                                                                                    v2 = ("\n\t" + v3)
                                                                                                                case 128:
                                                                                                                    fmt.Println(fi36)
                                                                                                                case -9223372036854775808:
                                                                                                                    var v111 uint = 0
                                                                                                                    _ = v111
                                                                                                                default:
                                                                                                                    var v112 string = v3
                                                                                                                    _ = v112
                                                                                                                }
                                                                                                            }
                                                                                                            if false {
                                                                                                                if false {
                                                                                                                    switch "a" {
                                                                                                                    case "hello":
                                                                                                                        _ = len(v25)
                                                                                                                    case "foo bar":
                                                                                                                        v39 -= 42
                                                                                                                    case "\n\t":
                                                                                                                        if (false || false) {
                                                                                                                            v3 = ""
                                                                                                                            fmt.Println(v2, fi36)
                                                                                                                            var v113 float32 = (3.14159 - 0.0)
                                                                                                                            _ = v113
                                                                                                                        } else {
                                                                                                                            _ = len(v1)
                                                                                                                            v114 := GenS7{F0: 42, F1: 0, F2: 246}
                                                                                                                            _ = v114
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v115 := GenS3{F0: v26}
                                                                                                                        _ = v115
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v116 := GenS8{F0: 32767}
                                                                                                                    v117 := GenS7{F0: (v39 % ((v39) | 1)), F1: 65535, F2: 469}
                                                                                                                    v108 -= v108
                                                                                                                    _ = v116
                                                                                                                    _ = v117
                                                                                                                }
                                                                                                                v26 -= (v26 << ((112) & 63))
                                                                                                            } else {
                                                                                                                _ = len(v1)
                                                                                                            }
                                                                                                            if (true || true) {
                                                                                                                var v118 int32 = genFn0(128)
                                                                                                                _ = v118
                                                                                                            } else {
                                                                                                                v119 := GenS3{F0: 567}
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            if false {
                                                                                                                v3 = (((("" + v3) + v2) + v2) + v3)
                                                                                                                var v120 uint32 = v26
                                                                                                                _ = v120
                                                                                                            } else {
                                                                                                                v121 := GenS1{F0: uint(v26)}
                                                                                                                _ = v121
                                                                                                            }
                                                                                                            v2 = "a"
                                                                                                            if len(v25) > 0 {
                                                                                                                v25[0] = (fi36 | 42)
                                                                                                            }
                                                                                                            v108 -= 1.5
                                                                                                        }
                                                                                                    }
                                                                                                    var v122 int16 = v39
                                                                                                    _ = v108
                                                                                                    _ = v122
                                                                                                } else {
                                                                                                    v123 := GenS6{F0: 160}
                                                                                                    _ = v123
                                                                                                }
                                                                                            } else {
                                                                                                var v124 float32 = (((0.1 - 1e10) - (float32(3.14159) / 1.5)) + float32(1e10))
                                                                                                v125 := GenS7{F0: 42, F1: 710, F2: 0}
                                                                                                _ = v124
                                                                                                _ = v125
                                                                                            }
                                                                                            v126 := GenS9{F0: "", F1: 2}
                                                                                            var v127 string = (((v2 + v2) + v3) + v2)
                                                                                            _ = v126
                                                                                            _ = v127
                                                                                        }
                                                                                        v26 = 65535
                                                                                        if (!(42 >= 443)) {
                                                                                            v3 = (("" + ("hello" + ((v3 + v2) + "foo bar"))) + "")
                                                                                        }
                                                                                        _ = v87
                                                                                    } else {
                                                                                        v39 = v39
                                                                                        var v128 int = fi22
                                                                                        _ = v128
                                                                                    }
                                                                                    if false {
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            v2 = (v3 + ("hello" + v3))
                                                                                            var v129 int = (^(fi36))
                                                                                            v130 := GenS9{F0: "a", F1: 65535}
                                                                                            _ = v129
                                                                                            _ = v130
                                                                                        case "\n\t":
                                                                                            switch fi36 {
                                                                                            case -52:
                                                                                                var v131 int16 = int16(v26)
                                                                                                _ = v131
                                                                                            case 1:
                                                                                                var v132 string = v3
                                                                                                _ = v132
                                                                                            default:
                                                                                                if (float64((float32((1e10 / (1e10 - ((0.1 / 1.0) + float64(0.0))))) + float32(((0.1 / ((-2.25 + 1e10) - 0.0)) - 0.1)))) == (0.1 * 0.0)) {
                                                                                                    var v133 uint = 1
                                                                                                    v134 := GenS10{F0: -1, F1: 4294967295}
                                                                                                    if len(v25) > 0 {
                                                                                                        v25[0] = fi36
                                                                                                    }
                                                                                                    _ = v133
                                                                                                    _ = v134
                                                                                                } else {
                                                                                                    if len(v25) >= 2 {
                                                                                                        v25 = v25[1:len(v25)]
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            v135 := GenS10{F0: 42, F1: 42}
                                                                                            _ = v135
                                                                                        default:
                                                                                            v136 := GenS9{F0: v3, F1: 65535}
                                                                                            _ = v136
                                                                                        }
                                                                                        switch (fi22 << ((0) & 63)) {
                                                                                        case -936:
                                                                                            v137 := GenS9{F0: "hello", F1: 2}
                                                                                            _ = v137
                                                                                        default:
                                                                                            fmt.Println(fi22, v26)
                                                                                        }
                                                                                    } else {
                                                                                        v2 = v3
                                                                                    }
                                                                                    fmt.Println(v3, fi22)
                                                                                }
                                                                            case "foo bar":
                                                                                switch v2 {
                                                                                case "a":
                                                                                    v3 = v2
                                                                                case "hello":
                                                                                    fmt.Println(v2, v3, v39)
                                                                                case "foo bar":
                                                                                    v138 := GenS11{F0: 127, F1: v3}
                                                                                    _ = v138
                                                                                default:
                                                                                    if (!false) {
                                                                                        switch fi22 {
                                                                                        case -807:
                                                                                            var v139 int32 = -978
                                                                                            _ = v139
                                                                                        default:
                                                                                            var v140 int16 = (v39 << ((763) & 63))
                                                                                            _ = v140
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    if false {
                                                                                        v141 := GenS12{F0: 2147483647, F1: (-(fi36))}
                                                                                        var v142 int32 = -2147483648
                                                                                        _ = v141
                                                                                        _ = v142
                                                                                    } else {
                                                                                        if (!true) {
                                                                                            var v143 uint8 = 128
                                                                                            fmt.Println(v39, v2, fi36)
                                                                                            v39 += 32767
                                                                                            _ = v143
                                                                                        }
                                                                                        switch v3 {
                                                                                        case "a":
                                                                                            var v144 string = v3
                                                                                            _ = v144
                                                                                        }
                                                                                    }
                                                                                    v145 := v1[fi36]
                                                                                    var v146 uint64 = uint64(v26)
                                                                                    _ = v145
                                                                                    _ = v146
                                                                                } else {
                                                                                    switch fi22 {
                                                                                    case 595:
                                                                                        if (42 == 42) {
                                                                                            var v147 int8 = 0
                                                                                            var v148 float64 = 0.0
                                                                                            v2 = v3
                                                                                            _ = v147
                                                                                            _ = v148
                                                                                        } else {
                                                                                            var v149 uint16 = 0
                                                                                            v26 -= (v26 + v26)
                                                                                            v149 = v149
                                                                                            _ = v149
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(v2, v26, v2)
                                                                                    default:
                                                                                        v150 := GenS13{F0: genFn0(127), F1: -128, F2: 1}
                                                                                        _ = v150
                                                                                    }
                                                                                }
                                                                                _ = len(v1)
                                                                                v26 += v26
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v151 uint32 = v26
                                                            _ = v151
                                                            _ = v39
                                                        } else {
                                                            _ = len(v25)
                                                            switch 32768 {
                                                            case -406:
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    switch fi22 {
                                                                    case 0:
                                                                        v3 = "foo bar"
                                                                    case -660:
                                                                        var v152 uint64 = 213
                                                                        _ = v152
                                                                    case -1:
                                                                        var v153 int16 = -32768
                                                                        _ = v153
                                                                    default:
                                                                        var v154 float64 = (((1e10 + 1.5) / 1.0) + 1e10)
                                                                        if false {
                                                                            if len(v25) > 0 {
                                                                                v25[0] = (fi22 % ((fi22) | 1))
                                                                            }
                                                                            fmt.Println(v154)
                                                                            if (-460 > -49) {
                                                                                v2 = ((v3 + (v3 + v3)) + (v2 + v3))
                                                                            } else {
                                                                                v3 = v2
                                                                            }
                                                                        }
                                                                        _ = v154
                                                                    }
                                                                case "a":
                                                                    var v155 uint64 = 460
                                                                    _ = v155
                                                                }
                                                            case 4294967295:
                                                                v156 := GenS14{F0: (("hello" + v3) + v2), F1: 127, F2: "foo bar"}
                                                                _ = v156
                                                            }
                                                            v157 := GenS7{F0: 748, F1: uint16(v26), F2: 988}
                                                            delete(v1, (fi36 | fi22))
                                                            var v158 int32 = 0
                                                            if ((float64(((0.1 - (((0.0 - (3.14159 - ((float32(3.14159) + 0.1) * 0.1))) / 1e10) - 1e10)) / (0.1 * (0.1 - float32(1e10))))) + (0.0 + 1.5)) == ((float64(0.0) - float64(-2.25)) + 0.1)) {
                                                                var v159 uint32 = 2147483648
                                                                _ = v159
                                                            }
                                                            _ = v157
                                                            _ = v158
                                                        }
                                                        if ((v26 - (v26 ^ 0)) == 545) {
                                                            fmt.Println(v3)
                                                            if true {
                                                                if ((((v2 + ("a" + (("\n\t" + (("\n\t" + ((v3 + (v3 + (("foo bar" + "") + ""))) + v3)) + "\n\t")) + (v2 + ((v3 + v2) + "hello"))))) + "\n\t") + ("" + "foo bar")) <= (v3 + "hello")) {
                                                                    fmt.Println(v3)
                                                                    v160 := GenS12{F0: 42, F1: fi22}
                                                                    _ = v160
                                                                } else {
                                                                    var v161 uint32 = 2
                                                                    _ = v161
                                                                }
                                                                fmt.Println(fi36)
                                                            } else {
                                                                v26 = (v26 << ((388) & 63))
                                                            }
                                                            v162 := GenS15{F0: int8(v26), F1: 0, F2: fi22}
                                                            _ = v162
                                                        }
                                                        if len(v25) >= 2 {
                                                            v25 = v25[1:len(v25)]
                                                        }
                                                    }
                                                    fmt.Println(fi36, v3)
                                                }
                                            case "hello":
                                                v2 = v3
                                            case "a":
                                                v2 = v2
                                            }
                                            v163 := GenS1{F0: 42}
                                            fmt.Println(v26, fi36)
                                            _ = fi36
                                            _ = v163
                                        }
                                    case "hello":
                                        var v164 uint16 = 32767
                                        _ = v164
                                    default:
                                        if true {
                                            switch (fi22 & -2147483648) {
                                            case 2147483647:
                                                v165 := v1[155]
                                                _ = v165
                                            case -912:
                                                if false {
                                                    for fi166 := 0; fi166 < 20; fi166++ {
                                                        var v167 uint8 = 49
                                                        var v168 uint = 255
                                                        _ = fi166
                                                        _ = v167
                                                        _ = v168
                                                    }
                                                }
                                            default:
                                                var v169 int16 = -1
                                                _ = v169
                                            }
                                            var v170 uint32 = v26
                                            _ = v170
                                        }
                                    }
                                } else {
                                    for fi171 := 0; fi171 < 8; fi171++ {
                                        v172 := GenS2{F0: -417}
                                        _ = fi171
                                        _ = v172
                                    }
                                }
                                for fi173 := 0; fi173 < 13; fi173++ {
                                    if false {
                                        v26 += (^(v26))
                                    } else {
                                        switch 520 {
                                        case 32768:
                                            v3 = ((v3 + v3) + ((("a" + "hello") + v2) + "foo bar"))
                                            if ((true || false) && (1e10 == 1.5)) {
                                                v174 := GenS16{F0: -65, F1: 18446744073709551615}
                                                var v175 string = v3
                                                if false {
                                                    if (!true) {
                                                        var v176 uint64 = 404
                                                        _ = v176
                                                    }
                                                    v177 := GenS9{F0: "hello", F1: 1}
                                                    _ = len(v25)
                                                    if (true && true) {
                                                        switch v3 {
                                                        case "":
                                                            var v178 float64 = 1.5
                                                            _ = v178
                                                        case "a":
                                                            switch (v3 + "") {
                                                            case "hello":
                                                                _ = len(v25)
                                                            case "a":
                                                                if false {
                                                                    if true {
                                                                        if len(v25) > 0 {
                                                                            v25[0] = fi173
                                                                        }
                                                                        fmt.Println(v3, fi173)
                                                                        if (!((v26 | v26) <= 2147483647)) {
                                                                            v3 = ((v3 + v3) + v175)
                                                                        }
                                                                    }
                                                                    v26 = (v26 % (((v26 + v26)) | 1))
                                                                    v179 := GenS16{F0: 1, F1: 676}
                                                                    var v180 int16 = 1000
                                                                    _ = v179
                                                                    _ = v180
                                                                }
                                                            case "foo bar":
                                                                if false {
                                                                    switch (fi22 + 32767) {
                                                                    case -2147483648:
                                                                        fmt.Println(fi22, fi22)
                                                                    case -9223372036854775808:
                                                                        v181 := v1[(^(fi173))]
                                                                        _ = v181
                                                                    case 32767:
                                                                        if len(v25) > 0 {
                                                                            v25[0] = fi22
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v175 = "a"
                                                                        var v182 int8 = -1
                                                                        v183 := GenS17{F0: -9223372036854775808}
                                                                        _ = v182
                                                                        _ = v183
                                                                    }
                                                                    switch v175 {
                                                                    case "a":
                                                                        v184 := GenS8{F0: (fi22 >> ((2) & 63))}
                                                                        _ = v184
                                                                    case "\n\t":
                                                                        delete(v1, fi22)
                                                                    case "hello":
                                                                        fmt.Println(v26, v2, v2)
                                                                    default:
                                                                        v185 := GenS7{F0: 0, F1: 483, F2: 1}
                                                                        _ = v185
                                                                    }
                                                                } else {
                                                                    fmt.Println(v26)
                                                                    switch (fi173 * (fi22 - (fi22 % ((fi22) | 1)))) {
                                                                    case 1:
                                                                        v1[(fi22 >> ((65535) & 63))] = v2
                                                                    case 32768:
                                                                        switch -9223372036854775808 {
                                                                        case 2147483647:
                                                                            switch 4294967295 {
                                                                            case 98:
                                                                                fmt.Println(v3)
                                                                            }
                                                                        case 4294967295:
                                                                            var v186 float64 = 1e10
                                                                            _ = v186
                                                                        case 42:
                                                                            if (((!true) || false) || false) {
                                                                                if len(v25) >= 2 {
                                                                                    v25 = v25[1:len(v25)]
                                                                                }
                                                                                v187 := GenS12{F0: -1, F1: 808}
                                                                                if (1 > 69) {
                                                                                    if true {
                                                                                        v188 := v1[2147483647]
                                                                                        fmt.Println(v175, v3, v188)
                                                                                        _ = v188
                                                                                    }
                                                                                }
                                                                                _ = v187
                                                                            } else {
                                                                                var v189 uint32 = (^(v26))
                                                                                _ = v189
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (!true) {
                                                                            var v190 string = v175
                                                                            switch (fi22 * fi22) {
                                                                            case 127:
                                                                                if (1000 == 0) {
                                                                                    v191 := GenS11{F0: 42, F1: "a"}
                                                                                    _ = v191
                                                                                }
                                                                            case -1:
                                                                                var v192 uint = 18446744073709551615
                                                                                _ = v192
                                                                            default:
                                                                                _ = len(v25)
                                                                            }
                                                                            _ = v190
                                                                        }
                                                                    }
                                                                    var v193 float64 = -2.25
                                                                    _ = v193
                                                                }
                                                            default:
                                                                var v194 int8 = -1
                                                                _ = v194
                                                            }
                                                        case "\n\t":
                                                            switch (fi22 * (fi22 ^ (fi173 / ((4294967295) | 1)))) {
                                                            case -230:
                                                                v3 = ((v3 + (v175 + (v2 + ((v3 + "hello") + v2)))) + ((((v3 + v3) + v3) + v175) + v2))
                                                            case 128:
                                                                var v195 uint = 65535
                                                                v196 := GenS14{F0: "a", F1: 1, F2: (v2 + "a")}
                                                                _ = v195
                                                                _ = v196
                                                            default:
                                                                _ = len(v25)
                                                            }
                                                        default:
                                                            var v197 uint8 = 42
                                                            _ = v197
                                                        }
                                                    } else {
                                                        v2 = "hello"
                                                        if (!(2 <= 255)) {
                                                            fmt.Println(v26, v175)
                                                        }
                                                        if (42 >= 4294967295) {
                                                            v26 -= v26
                                                            var v198 int64 = -9223372036854775808
                                                            var v199 string = "\n\t"
                                                            _ = v198
                                                            _ = v199
                                                        } else {
                                                            var v200 uint16 = 753
                                                            var v201 uint = 2
                                                            _ = v200
                                                            _ = v201
                                                        }
                                                        var v202 bool = (!(!true))
                                                        var v203 uint64 = 0
                                                        _ = v202
                                                        _ = v203
                                                    }
                                                    _ = v177
                                                } else {
                                                    v204 := GenS18{F0: 2, F1: "a", F2: 42}
                                                    _ = v204
                                                }
                                                delete(v1, (fi173 ^ (-(fi22))))
                                                var v205 float32 = 0.1
                                                _ = v174
                                                _ = v175
                                                _ = v205
                                            }
                                        default:
                                            if (1 == -1) {
                                                switch v2 {
                                                case "":
                                                    v206 := v1[127]
                                                    _ = v206
                                                case "hello":
                                                    v25 = append(v25, (fi22 | (fi173 + 4294967295)))
                                                case "a":
                                                    if ((9223372036854775807 > -1) || ((fi22 / (((fi22 / ((fi22) | 1))) | 1)) < (-(fi22)))) {
                                                        if len(v25) > 0 {
                                                            v25[0] = (^(fi173))
                                                        }
                                                        var v207 string = v2
                                                        _ = v207
                                                    } else {
                                                        v2 = "foo bar"
                                                        if (-904 <= int16(fi22)) {
                                                            fmt.Println(fi22, v2, v2)
                                                            fmt.Println(v3)
                                                        } else {
                                                            v2 = ""
                                                            v26 += (v26 | (v26 + v26))
                                                        }
                                                    }
                                                }
                                                var v208 uint64 = 298
                                                v26 += v26
                                                _ = v208
                                            } else {
                                                var v209 int32 = -2147483648
                                                var v210 int = (^(fi22))
                                                v211 := GenS14{F0: "foo bar", F1: genFn1("a"), F2: "hello"}
                                                _ = v209
                                                _ = v210
                                                _ = v211
                                            }
                                        }
                                        if len(v25) > 0 {
                                            v25[0] = (^(fi22))
                                        }
                                        var v212 int64 = 0
                                        _ = v212
                                    }
                                    v213 := v1[(fi173 - fi22)]
                                    if (-32768 > 1) {
                                        var v214 uint = 2
                                        switch v213 {
                                        case "\n\t":
                                            v26 -= (^(v26))
                                        }
                                        fmt.Println(fi22, fi22, v3)
                                        v214 = v214
                                        _ = v214
                                    }
                                    _ = fi173
                                    _ = v213
                                }
                            } else {
                                if len(v25) > 0 {
                                    v25[0] = 1
                                }
                                v2 = v3
                            }
                        case -2147483648:
                            switch fi22 {
                            case 127:
                                var v215 int8 = 84
                                _ = v215
                            default:
                                if false {
                                    for fi216 := 0; fi216 < 5; fi216++ {
                                        fmt.Println(v2, fi22, v2)
                                        if false {
                                            switch v3 {
                                            case "\n\t":
                                                var v217 uint = 4294967295
                                                _ = v217
                                            case "foo bar":
                                                var v218 uint8 = 2
                                                _ = v218
                                            default:
                                                if true {
                                                    var v219 float32 = -2.25
                                                    v3 = (v2 + (v2 + v2))
                                                    v219 = v219
                                                    fmt.Println(fi22, v2)
                                                    v26 = v26
                                                    var v220 float64 = -2.25
                                                    if false {
                                                        var v221 int32 = genFn0(255)
                                                        var v222 int64 = 3
                                                        if (!true) {
                                                            if (((v222 ^ (v222 - (v222 / ((int64(v221)) | 1)))) >= v222) || true) {
                                                                var v223 int64 = (v222 % (((^(v222))) | 1))
                                                                _ = v223
                                                            } else {
                                                                fmt.Println(v219, v2, v221)
                                                                v2 = v2
                                                            }
                                                            switch -106 {
                                                            case 42:
                                                                var v224 int32 = (^(v221))
                                                                _ = v224
                                                            case 9223372036854775807:
                                                                if false {
                                                                    var v225 int32 = genFn0(109)
                                                                    v225 = genFn0(15)
                                                                    _ = v225
                                                                } else {
                                                                    var v226 uint16 = 523
                                                                    var v227 string = ""
                                                                    v220 = (1e10 - 1.5)
                                                                    _ = v226
                                                                    _ = v227
                                                                }
                                                            }
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(v219, v2, v3)
                                                            case "":
                                                                v25 = append(v25, (fi22 % ((fi22) | 1)))
                                                            default:
                                                                fmt.Println(v3)
                                                                v1[(fi216 - fi216)] = v3
                                                            }
                                                        } else {
                                                            var v228 uint32 = v26
                                                            fmt.Println(v222)
                                                            _ = v228
                                                        }
                                                        _ = v221
                                                        _ = v222
                                                    }
                                                    _ = v219
                                                    _ = v220
                                                } else {
                                                    if true {
                                                        v229 := GenS3{F0: (^(v26))}
                                                        if false {
                                                            v1[32767] = "foo bar"
                                                            var v230 float32 = float32(1e10)
                                                            _ = v230
                                                        } else {
                                                            if (!true) {
                                                                var v231 float64 = 0.0
                                                                if (-128 == -128) {
                                                                    var v232 int64 = 0
                                                                    var v233 bool = ((genFn0(110) <= -135) && true)
                                                                    fmt.Println(v2, v2)
                                                                    _ = len(v1)
                                                                    _ = v232
                                                                    _ = v233
                                                                } else {
                                                                    if (186 > 60) {
                                                                        v234 := GenS10{F0: 0, F1: 18446744073709551615}
                                                                        v235 := GenS7{F0: 1, F1: 255, F2: 52}
                                                                        var v236 uint = 0
                                                                        _ = v234
                                                                        _ = v235
                                                                        _ = v236
                                                                    }
                                                                    var v237 float64 = -2.25
                                                                    v26 -= v26
                                                                    _ = v237
                                                                }
                                                                var v238 int64 = -9223372036854775808
                                                                switch (v2 + "") {
                                                                case "":
                                                                    v2 = "\n\t"
                                                                    fmt.Println(fi22, fi22)
                                                                default:
                                                                    v3 = ""
                                                                }
                                                                if len(v25) >= 2 {
                                                                    v25 = v25[1:len(v25)]
                                                                }
                                                                _ = v231
                                                                _ = v238
                                                            } else {
                                                                fmt.Println(v26)
                                                            }
                                                            v2 = "hello"
                                                            var v239 int32 = genFn0(129)
                                                            _ = v239
                                                        }
                                                        v2 = (("hello" + v3) + v3)
                                                        fmt.Println(v3, v26, v2)
                                                        _ = v229
                                                    } else {
                                                        v2 = v2
                                                        v240 := GenS9{F0: v3, F1: 1}
                                                        fmt.Println(v3, fi216)
                                                        switch (fi22 ^ fi22) {
                                                        case -764:
                                                            fmt.Println(fi216, v2, v26)
                                                        case 32768:
                                                            v3 = ""
                                                            var v241 int8 = 1
                                                            v242 := v1[0]
                                                            var v243 int16 = 42
                                                            _ = v241
                                                            _ = v242
                                                            _ = v243
                                                        case -3:
                                                            v244 := GenS7{F0: -937, F1: 327, F2: 42}
                                                            _ = v244
                                                        default:
                                                            v26 = v26
                                                        }
                                                        _ = v240
                                                    }
                                                }
                                            }
                                            v245 := GenS16{F0: 776, F1: 965}
                                            fmt.Println(fi22)
                                            _ = v245
                                        }
                                        _ = fi216
                                    }
                                }
                            }
                        case 42:
                            fmt.Println(v26, v3, fi22)
                        default:
                            var v246 uint64 = 42
                            _ = v246
                        }
                        var v247 uint64 = 609
                        _ = v247
                        _ = v26
                    } else {
                        v248 := GenS16{F0: 1, F1: 4294967295}
                        switch v2 {
                        case "hello":
                            for fi249 := 0; fi249 < 1; fi249++ {
                                v2 = (v3 + v2)
                                var v250 int64 = -230
                                delete(v1, (^(fi22)))
                                _ = fi249
                                _ = v250
                            }
                        case "a":
                            if (18446744073709551615 >= 4294967295) {
                                v2 = v2
                                var v251 uint8 = 122
                                var v252 float64 = (float64(0.0) + float64(-2.25))
                                _ = v251
                                _ = v252
                            }
                        case "":
                            var v253 float32 = 1e10
                            _ = v253
                        default:
                            if true {
                                v1[-9223372036854775808] = ("" + "a")
                                var v254 uint32 = 2147483647
                                fmt.Println(v3)
                                _ = v254
                            }
                        }
                        if false {
                            v255 := GenS14{F0: "\n\t", F1: -128, F2: v2}
                            _ = v255
                        } else {
                            var v256 int = fi22
                            for fi257 := 0; fi257 < 5; fi257++ {
                                if false {
                                    fmt.Println(v3)
                                    v25 = append(v25, fi257)
                                    v3 = "\n\t"
                                } else {
                                    fmt.Println(v2)
                                    var v258 bool = false
                                    _ = v258
                                }
                                _ = fi257
                            }
                            switch (v256 ^ fi22) {
                            case 128:
                                var v259 uint8 = 42
                                _ = v259
                            }
                            _ = v256
                        }
                        _ = v248
                    }
                    _ = v25
                } else {
                    v260 := []int{}
                    switch v2 {
                    case "\n\t":
                        if true {
                            var v261 int16 = 42
                            for fi262 := 0; fi262 < 2; fi262++ {
                                fmt.Println(fi262)
                                v3 = v3
                                if false {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v263 bool = (!(0.1 < 0.1))
                                        _ = v263
                                    }
                                    if true {
                                        v1[fi262] = "hello"
                                        if (0.0 != 1e10) {
                                            var v264 float32 = (0.0 / (3.14159 / 1.0))
                                            switch v3 {
                                            case "foo bar":
                                                var v265 int32 = -2147483648
                                                _ = v265
                                            }
                                            _ = v264
                                        }
                                        v261 = (v261 / ((1000) | 1))
                                    } else {
                                        if (65535 == 2) {
                                            if false {
                                                v266 := GenS19{F0: ("foo bar" + ("hello" + "hello")), F1: true, F2: "hello"}
                                                _ = v266
                                            } else {
                                                switch fi262 {
                                                case 127:
                                                    v3 = ("foo bar" + "")
                                                case 4294967295:
                                                    if true {
                                                        v2 = "foo bar"
                                                        v267 := GenS17{F0: fi262}
                                                        _ = v267
                                                    }
                                                default:
                                                    v261 -= (^(v261))
                                                }
                                                v268 := GenS20{F0: v261, F1: 981, F2: v261}
                                                delete(v1, (fi262 % (((^(fi22))) | 1)))
                                                _ = v268
                                            }
                                        }
                                        _ = len(v260)
                                    }
                                } else {
                                    var v269 float32 = 0.1
                                    _ = v269
                                }
                                _ = fi262
                            }
                            _ = v261
                        }
                    case "hello":
                        v3 = v3
                    }
                    var v270 uint64 = 2
                    _ = v260
                    _ = v270
                }
                v1[(fi22 + (fi22 | (fi22 & fi22)))] = "foo bar"
                for fi271 := 0; fi271 < 19; fi271++ {
                    switch v3 {
                    case "foo bar":
                        v3 = v2
                    case "":
                        var v272 float32 = float32(0.0)
                        _ = v272
                    default:
                        if false {
                            if (!true) {
                                v273 := []int{(fi271 * 128)}
                                _ = v273
                            } else {
                                v274 := GenS12{F0: 0, F1: fi22}
                                v2 = v3
                                _ = v274
                            }
                            if (65535 >= 2147483648) {
                                v275 := GenS4{F0: 863}
                                var v276 string = v3
                                _ = v275
                                _ = v276
                            } else {
                                v277 := GenS18{F0: 262, F1: v2, F2: 1}
                                v278 := GenS21{F0: 1, F1: 127, F2: 42}
                                _ = v277
                                _ = v278
                            }
                        } else {
                            var v279 int8 = 42
                            var v280 uint8 = 0
                            _ = v279
                            _ = v280
                        }
                    }
                    _ = fi271
                }
            }
            _ = fi22
        }
    }
    v3 = "a"
    v3 = v3
    v281 := GenS18{F0: 1, F1: (((v2 + v3) + "\n\t") + (v2 + v3)), F2: 106}
    for fi282 := 0; fi282 < 10; fi282++ {
        for fi283 := 0; fi283 < 17; fi283++ {
            var v284 uint = 1
            _ = fi283
            _ = v284
        }
        var v285 int32 = 891
        for fi286 := 0; fi286 < 11; fi286++ {
            var v287 uint32 = 829
            _ = fi286
            _ = v287
        }
        _ = fi282
        _ = v285
    }
    v288 := GenS7{F0: -701, F1: 211, F2: 42}
    v2 = v2
    fmt.Println(v2)
    v3 = v3
    v289 := []int{-48}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v289:", len(v289), v289)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v281
    _ = v288
    _ = v289
    _ = v3
}
