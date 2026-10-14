package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8; F1 int8; F2 int32 }
type GenS12 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int8; F1 uint16; F2 int8 }
type GenS15 struct { F0 int32; F1 int64; F2 int }
type GenS16 struct { F0 string; F1 int16; F2 uint }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint32; F1 bool; F2 int32 }
type GenS20 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS21 struct { F0 int; F1 uint32 }
func genFn0(p0 int, p1 uint64, p2 int64) int {
    v5 := []int{(p0 - p0), p0, (p0 * -2147483648), p0}
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            if len(v5) > 0 {
                v5[0] = (fi6 & fi6)
            }
            _ = fi6
        }
    } else {
        v7 := GenS1{F0: 1000, F1: (p1 - (p1 * p1))}
        for fi8 := 0; fi8 < 9; fi8++ {
            p1 += 448
            _ = fi8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    var v1 uint16 = 480
    var v2 int = 9223372036854775807
    var v3 int64 = 0
    var v4 bool = (((0.0 + (-2.25 + 1e10)) - 0.0) == 0.0)
    if ((76 != 619) && (!(-32768 < 182))) {
        v9 := []int{(v2 >> ((42) & 63)), (v2 * (v2 * (^(v2))))}
        var v10 int16 = 1
        _ = v10
        _ = v9
    } else {
        v1 -= 223
    }
    v3 = (v3 ^ v3)
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v2 = (v2 >> ((2) & 63))
            v13 := []int{(fi11 >> ((255) & 63))}
            v14 := GenS2{F0: 1, F1: 533}
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 uint64 = 42
            var v17 int32 = -98
            v18 := []int{(fi15 / ((fi15) | 1))}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi11
    }
    v3 -= (v3 | -180)
    v19 := GenS3{F0: -2147483648}
    v20 := []int{(v2 - genFn0(713, 833, v3)), 1, 32767, -9223372036854775808}
    var v21 uint8 = 127
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            for fi23 := 0; fi23 < 2; fi23++ {
                switch -875 {
                case -2147483648:
                    v20 = append(v20, genFn0((fi23 >> ((605) & 63)), 2, 493))
                case 2147483647:
                    if ((^(v21)) < (^(v21))) {
                        var v24 bool = v4
                        _ = v24
                    } else {
                        v25 := GenS2{F0: 33, F1: 88}
                        switch fi22 {
                        case 894:
                            v26 := GenS3{F0: 42}
                            _ = v26
                        case 7:
                            switch -2147483648 {
                            case 127:
                                var v27 uint8 = v21
                                _ = v27
                            case -9223372036854775808:
                                var v28 int32 = 1
                                _ = v28
                            }
                        case 127:
                            var v29 uint = 4294967295
                            _ = v29
                        }
                        _ = v25
                    }
                default:
                    v3 += v3
                }
                v21 += v21
                if (("" + "foo bar") > ((("\n\t" + "hello") + (("hello" + "foo bar") + "foo bar")) + "")) {
                    if ("foo bar" == "\n\t") {
                        switch (("a" + "hello") + ((("a" + ("\n\t" + "hello")) + "a") + "a")) {
                        case "\n\t":
                            fmt.Println(v3)
                            fmt.Println(v21, v2, v4)
                            if false {
                                var v30 uint32 = 4294967295
                                _ = v30
                            } else {
                                switch genFn0(genFn0(128, 4294967295, -9223372036854775808), 449, 9223372036854775807) {
                                case 230:
                                    v31 := GenS4{F0: 50, F1: ("\n\t" + "")}
                                    _ = v31
                                }
                                if (!v4) {
                                    var v32 int16 = 937
                                    v33 := GenS5{F0: fi22, F1: 42}
                                    _ = v32
                                    _ = v33
                                } else {
                                    if v4 {
                                        var v34 uint8 = v21
                                        _ = v34
                                    } else {
                                        if v4 {
                                            fmt.Println(v4)
                                        } else {
                                            v2 = genFn0(-616, 230, (v3 + -1))
                                        }
                                        v21 = 255
                                    }
                                    v35 := GenS3{F0: 1}
                                    v21 -= v21
                                    _ = v35
                                }
                                var v36 int32 = 843
                                _ = v36
                            }
                        default:
                            var v37 float32 = float32(1e10)
                            _ = v37
                        }
                    } else {
                        if ((42 >= 463) && (("\n\t" + ("a" + ("foo bar" + ""))) < (("a" + "\n\t") + "foo bar"))) {
                            var v38 bool = false
                            var v39 uint64 = 42
                            v4 = false
                            _ = v38
                            _ = v39
                        } else {
                            switch "hello" {
                            case "\n\t":
                                v40 := GenS6{F0: (v21 - v21), F1: 325}
                                v21 = v21
                                _ = v40
                            default:
                                fmt.Println(v21, v1, fi22)
                            }
                            fmt.Println(fi22, fi22)
                            var v41 int8 = 127
                            _ = v41
                        }
                        if v4 {
                            if (((v21 << ((108) & 63)) < (v21 ^ (^(v21)))) && true) {
                                v2 = (v2 << ((2) & 63))
                                v1 = (^(v1))
                            } else {
                                var v42 int64 = 134
                                v3 += v42
                                _ = v42
                            }
                            switch (fi22 & (v2 & (fi23 / ((genFn0((fi22 - v2), 18446744073709551615, (-(v3)))) | 1)))) {
                            case -1:
                                var v43 int = (fi23 - fi23)
                                _ = v43
                            default:
                                switch "hello" {
                                case "hello":
                                    var v44 string = "a"
                                    _ = v44
                                }
                            }
                            v45 := GenS7{F0: "hello", F1: v3}
                            _ = v45
                        } else {
                            if ("hello" < "a") {
                                v4 = v4
                                var v46 int64 = v3
                                if (true && false) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v47 uint64 = uint64(v21)
                                        _ = v47
                                    case "":
                                        if (4294967295 == 4294967295) {
                                            var v48 uint32 = 362
                                            v20 = append(v20, fi22)
                                            switch (v2 & (fi23 / ((fi23) | 1))) {
                                            case 32768:
                                                v49 := GenS8{F0: fi23}
                                                _ = v49
                                            default:
                                                var v50 uint = 109
                                                _ = v50
                                            }
                                            _ = v48
                                        } else {
                                            switch genFn0(fi23, 4294967295, (^(v3))) {
                                            case 0:
                                                if v4 {
                                                    v3 = (v3 & v3)
                                                    v51 := GenS5{F0: genFn0(v2, 0, (v46 ^ (^(v46)))), F1: (v2 >> ((484) & 63))}
                                                    v3 -= v46
                                                    _ = v51
                                                } else {
                                                    v52 := GenS8{F0: (fi23 + (v2 ^ (fi23 * (^(v2)))))}
                                                    if (!v4) {
                                                        if (v4 || v4) {
                                                            if v4 {
                                                                switch (("a" + "a") + ((("\n\t" + (("\n\t" + "") + (("" + ("\n\t" + "hello")) + "\n\t"))) + "hello") + ("\n\t" + ""))) {
                                                                case "foo bar":
                                                                    switch "" {
                                                                    case "a":
                                                                        if (v4 || v4) {
                                                                            var v53 float64 = ((3.14159 / -2.25) / 1.0)
                                                                            _ = len(v20)
                                                                            v2 = fi22
                                                                            var v54 uint = 76
                                                                            var v55 uint = v54
                                                                            var v56 uint16 = (v1 & v1)
                                                                            _ = v53
                                                                            _ = v54
                                                                            _ = v55
                                                                            _ = v56
                                                                        } else {
                                                                            fmt.Println(v2, fi23, v46)
                                                                            fmt.Println(v21, v21)
                                                                            v20 = append(v20, fi22)
                                                                        }
                                                                    case "hello":
                                                                        if true {
                                                                            if (0.1 > (1.5 * 1e10)) {
                                                                                v57 := GenS9{F0: (v1 % ((v1) | 1))}
                                                                                if true {
                                                                                    fmt.Println(v2, v21, v1)
                                                                                    if len(v20) > 0 {
                                                                                        v20[0] = (fi22 - (fi22 - fi23))
                                                                                    }
                                                                                    v21 -= 42
                                                                                }
                                                                                v4 = (v4 && false)
                                                                                fmt.Println(fi22, v2, v3)
                                                                                var v58 uint32 = 0
                                                                                _ = v57
                                                                                _ = v58
                                                                            } else {
                                                                                switch fi23 {
                                                                                case 9223372036854775807:
                                                                                    if len(v20) >= 2 {
                                                                                        v20 = v20[1:len(v20)]
                                                                                    }
                                                                                case 127:
                                                                                    switch 1 {
                                                                                    case -2147483648:
                                                                                        if v4 {
                                                                                            var v59 int64 = 462
                                                                                            switch 32768 {
                                                                                            case 2147483647:
                                                                                                v1 = v1
                                                                                            case 0:
                                                                                                switch fi22 {
                                                                                                case 42:
                                                                                                    var v60 uint8 = (v21 / ((42) | 1))
                                                                                                    _ = v60
                                                                                                case -998:
                                                                                                    var v61 bool = (42 >= 499)
                                                                                                    _ = v61
                                                                                                case -882:
                                                                                                    v59 += v46
                                                                                                default:
                                                                                                    v3 -= int64(v21)
                                                                                                }
                                                                                            default:
                                                                                                v59 = 9223372036854775807
                                                                                            }
                                                                                            _ = v59
                                                                                        }
                                                                                    default:
                                                                                        if len(v20) > 0 {
                                                                                            v20[0] = (v2 % ((fi22) | 1))
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v62 float32 = float32(1.5)
                                                                            v63 := GenS6{F0: (^(v21)), F1: 42}
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            var v64 int64 = v46
                                                                            var v65 int64 = 385
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                var v66 int32 = -2147483648
                                                                if v4 {
                                                                    v20 = append(v20, 372)
                                                                    if true {
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                    }
                                                                    if v4 {
                                                                        v67 := GenS9{F0: (^(v1))}
                                                                        v46 = v46
                                                                        if len(v20) > 0 {
                                                                            v20[0] = (v2 >> ((255) & 63))
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                }
                                                                v46 -= 42
                                                                _ = v66
                                                            }
                                                            if true {
                                                                if v4 {
                                                                    v2 += (fi22 - (fi22 - fi23))
                                                                    if false {
                                                                        var v68 float64 = -2.25
                                                                        fmt.Println(v1, fi22)
                                                                        var v69 int64 = (^(v46))
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                    if ("hello" != "a") {
                                                                        if v4 {
                                                                            v21 = (v21 << ((261) & 63))
                                                                            v70 := GenS10{F0: -1}
                                                                            fmt.Println(v4, v3, v4)
                                                                            v71 := GenS4{F0: v21, F1: (("" + "") + ("\n\t" + ""))}
                                                                            var v72 uint16 = (v1 % ((v1) | 1))
                                                                            if true {
                                                                                var v73 float32 = 1e10
                                                                                _ = v73
                                                                            } else {
                                                                                _ = len(v20)
                                                                            }
                                                                            _ = v70
                                                                            _ = v71
                                                                            _ = v72
                                                                        } else {
                                                                            if len(v20) > 0 {
                                                                                v20[0] = (fi22 << ((372) & 63))
                                                                            }
                                                                            var v74 string = "foo bar"
                                                                            v75 := GenS11{F0: v21, F1: 102, F2: -1}
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        var v76 uint16 = 427
                                                                        _ = v76
                                                                    }
                                                                    if (!v4) {
                                                                        v77 := GenS3{F0: -1}
                                                                        v21 -= v21
                                                                        _ = v77
                                                                    }
                                                                }
                                                                if v4 {
                                                                    var v78 uint32 = 760
                                                                    _ = v78
                                                                }
                                                            }
                                                        }
                                                        v2 = fi23
                                                        if false {
                                                            if true {
                                                                if false {
                                                                    if len(v20) > 0 {
                                                                        v20[0] = -743
                                                                    }
                                                                }
                                                                var v79 int32 = int32(v2)
                                                                var v80 float64 = -2.25
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                            v1 = (^(v1))
                                                        }
                                                        var v81 uint16 = v1
                                                        v1 += (^(v1))
                                                        _ = v81
                                                    }
                                                    v3 = int64(v1)
                                                    _ = v52
                                                }
                                            case -776:
                                                var v82 float32 = (float32(0.1) + float32(1e10))
                                                _ = v82
                                            }
                                            v2 -= (v2 & (fi23 * 42))
                                            var v83 string = (((("foo bar" + ("" + "")) + "a") + ("hello" + "hello")) + "")
                                            _ = v83
                                        }
                                    case "a":
                                        v46 += v3
                                    }
                                    var v84 uint16 = (^(v1))
                                    _ = v84
                                }
                                _ = v46
                            } else {
                                var v85 string = "\n\t"
                                var v86 int32 = -149
                                _ = v85
                                _ = v86
                            }
                            var v87 uint64 = 1
                            v3 = (^(v3))
                            _ = v87
                        }
                        var v88 float64 = (1e10 * ((float64(-2.25) + ((1.5 - ((1e10 - ((0.0 * 0.1) * -2.25)) + (float64(3.14159) / 0.1))) - ((0.0 - (-2.25 / 1.5)) / 1.0))) + 1.5))
                        v2 += fi22
                        v89 := GenS12{F0: (^(v3)), F1: 1, F2: 42}
                        _ = v88
                        _ = v89
                    }
                    v4 = (42 != 9223372036854775807)
                    v20 = append(v20, 843)
                    v1 = (v1 / (((^(v1))) | 1))
                    if (0 >= v1) {
                        var v90 int64 = (v3 | 42)
                        _ = len(v20)
                        _ = v90
                    }
                }
                var v91 float64 = 3.14159
                _ = fi23
                _ = v91
            }
            var v92 string = ((((("\n\t" + ("\n\t" + "\n\t")) + "foo bar") + "foo bar") + "foo bar") + ("\n\t" + "\n\t"))
            _ = v92
        }
        switch -673 {
        case 345:
            for fi93 := 0; fi93 < 18; fi93++ {
                v3 -= (v3 ^ 1)
                v1 = (^(v1))
                if v4 {
                    v4 = v4
                    var v94 float32 = 1.5
                    v95 := GenS1{F0: 42, F1: 0}
                    _ = v94
                    _ = v95
                } else {
                    if (genFn0((v2 ^ genFn0((fi93 - (^(fi22))), 0, 514)), 928, int64(fi22)) >= genFn0((fi22 * v2), 1, v3)) {
                        if v4 {
                            switch (("a" + "hello") + "foo bar") {
                            case "\n\t":
                                v2 = (fi22 % ((128) | 1))
                                v3 += v3
                            case "":
                                switch "" {
                                case "\n\t":
                                    _ = len(v20)
                                case "foo bar":
                                    if v4 {
                                        _ = len(v20)
                                        v2 -= fi22
                                        var v96 float32 = (float32((float64(1e10) * -2.25)) + ((1.5 / 1.5) / -2.25))
                                        _ = v96
                                    } else {
                                        if v4 {
                                            fmt.Println(fi93, v2, v21)
                                            if (false && (42 > 1)) {
                                                var v97 string = "\n\t"
                                                if true {
                                                    v3 = v3
                                                }
                                                v4 = (!(96 != v21))
                                                _ = v97
                                            } else {
                                                var v98 uint64 = 18446744073709551615
                                                _ = v98
                                            }
                                            v99 := GenS12{F0: v3, F1: 2, F2: 42}
                                            v21 += v21
                                            _ = v99
                                        }
                                        v3 = v3
                                    }
                                }
                            }
                            v21 -= (v21 & 1)
                        } else {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            var v100 float64 = 3.14159
                            var v101 bool = false
                            _ = v100
                            _ = v101
                        }
                    }
                    var v102 int8 = 1
                    if (!(v2 <= int(v102))) {
                        switch "" {
                        case "a":
                            var v103 string = "a"
                            _ = v103
                        case "hello":
                            v1 -= v1
                        case "":
                            var v104 int8 = 0
                            _ = v104
                        default:
                            v105 := GenS13{F0: 1000}
                            _ = v105
                        }
                        if (fi93 != genFn0(fi22, 827, -139)) {
                            v3 += (v3 % ((v3) | 1))
                        } else {
                            if len(v20) > 0 {
                                v20[0] = int(v102)
                            }
                            _ = len(v20)
                        }
                        v1 -= v1
                        switch fi22 {
                        case 1:
                            if ((v4 && v4) || ((2 <= 42) || v4)) {
                                if false {
                                    var v106 float32 = 3.14159
                                    var v107 int8 = v102
                                    v4 = true
                                    _ = v106
                                    _ = v107
                                }
                                switch (^(fi93)) {
                                case 128:
                                    if (((v1 & v1) > v1) || true) {
                                        v4 = (!(false || false))
                                        fmt.Println(v3, v3, fi93)
                                        switch 456 {
                                        case -913:
                                            if v4 {
                                                if len(v20) > 0 {
                                                    v20[0] = (fi93 >> ((65535) & 63))
                                                }
                                                var v108 string = ""
                                                _ = v108
                                            } else {
                                                v102 = (v102 % ((0) | 1))
                                                v2 += v2
                                            }
                                        default:
                                            v109 := GenS8{F0: 2147483647}
                                            _ = v109
                                        }
                                        switch genFn0(fi22, 4294967295, v3) {
                                        case -233:
                                            if false {
                                                switch (fi93 ^ (fi93 << ((4294967295) & 63))) {
                                                case 128:
                                                    switch v2 {
                                                    case -9223372036854775808:
                                                        if v4 {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                switch (fi22 + v2) {
                                                                case 1:
                                                                    v110 := GenS2{F0: int8(fi22), F1: 65535}
                                                                    _ = v110
                                                                case 0:
                                                                    v21 = (v21 + 255)
                                                                case 127:
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                    }
                                                                    v111 := GenS11{F0: (v21 & v21), F1: 64, F2: 0}
                                                                    _ = v111
                                                                default:
                                                                    v3 = -1
                                                                }
                                                            case "":
                                                                if len(v20) > 0 {
                                                                    v20[0] = (fi93 & (^(v2)))
                                                                }
                                                            case "\n\t":
                                                                var v112 uint32 = uint32(v1)
                                                                switch genFn0(-825, 42, (v3 ^ (v3 | v3))) {
                                                                case -9223372036854775808:
                                                                    var v113 float32 = 1e10
                                                                    _ = v113
                                                                case 9223372036854775807:
                                                                    fmt.Println(v21, v3)
                                                                case -103:
                                                                    v21 = v21
                                                                }
                                                                if v4 {
                                                                    if (float32((3.14159 + float64(0.0))) <= (-2.25 / 1.5)) {
                                                                        v4 = v4
                                                                        switch (fi93 + genFn0(1, 71, (v3 << ((2) & 63)))) {
                                                                        case -748:
                                                                            if v4 {
                                                                                v4 = v4
                                                                            }
                                                                        case 128:
                                                                            v114 := GenS14{F0: v102, F1: (v1 * (v1 * (v1 & v1))), F2: (v102 & 127)}
                                                                            _ = v114
                                                                        default:
                                                                            var v115 uint32 = v112
                                                                            if len(v20) > 0 {
                                                                                v20[0] = int(v3)
                                                                            }
                                                                            v116 := GenS15{F0: 0, F1: (v3 % ((v3) | 1)), F2: fi22}
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                    if false {
                                                                        if v4 {
                                                                            _ = len(v20)
                                                                            _ = len(v20)
                                                                        } else {
                                                                            if ((v3 ^ 1) >= v3) {
                                                                                var v117 string = "foo bar"
                                                                                var v118 int16 = 32767
                                                                                v20 = append(v20, (fi93 | (v2 * (-(fi22)))))
                                                                                var v119 float32 = (0.0 / 1.5)
                                                                                v2 = (v2 % ((genFn0(fi93, 18446744073709551615, v3)) | 1))
                                                                                _ = v117
                                                                                _ = v118
                                                                                _ = v119
                                                                            } else {
                                                                                v120 := GenS6{F0: v21, F1: 1}
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                        var v121 string = "foo bar"
                                                                        var v122 uint32 = (v112 >> ((2) & 63))
                                                                        _ = v121
                                                                        _ = v122
                                                                    } else {
                                                                        var v123 float32 = 1e10
                                                                        _ = v123
                                                                    }
                                                                    _ = len(v20)
                                                                }
                                                                _ = v112
                                                            default:
                                                                v1 = (v1 >> ((0) & 63))
                                                            }
                                                            var v124 int = v2
                                                            if (false && true) {
                                                                var v125 int8 = -128
                                                                _ = v125
                                                            }
                                                            _ = v124
                                                        }
                                                    case 32767:
                                                        if (18446744073709551615 <= 420) {
                                                            var v126 int32 = 0
                                                            var v127 float32 = float32(3.14159)
                                                            fmt.Println(v102)
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            switch "foo bar" {
                                                            case "a":
                                                                if (!false) {
                                                                    v20 = append(v20, fi93)
                                                                    if false {
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                        v128 := GenS7{F0: "hello", F1: -9223372036854775808}
                                                                        _ = v128
                                                                    }
                                                                } else {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v129 uint32 = 1
                                                                        _ = v129
                                                                    case "a":
                                                                        v2 -= (v2 * (^(fi22)))
                                                                    default:
                                                                        var v130 string = "a"
                                                                        _ = v130
                                                                    }
                                                                    var v131 int64 = (v3 * v3)
                                                                    _ = v131
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v1)
                                                            case "hello":
                                                                v132 := GenS3{F0: 0}
                                                                _ = v132
                                                            default:
                                                                v21 += v21
                                                            }
                                                        }
                                                    case 4294967295:
                                                        var v133 uint8 = v21
                                                        _ = v133
                                                    }
                                                case -116:
                                                    if false {
                                                        if (!(v3 == 9223372036854775807)) {
                                                            _ = len(v20)
                                                            var v134 int64 = v3
                                                            if len(v20) > 0 {
                                                                v20[0] = fi22
                                                            }
                                                            if (1 == 1) {
                                                                var v135 string = ""
                                                                _ = v135
                                                            }
                                                            if len(v20) > 0 {
                                                                v20[0] = (fi22 % (((v2 ^ (^(fi93)))) | 1))
                                                            }
                                                            v136 := GenS16{F0: "a", F1: 1000, F2: 824}
                                                            _ = v134
                                                            _ = v136
                                                        } else {
                                                            switch genFn0(v2, 1, v3) {
                                                            case -1:
                                                                v137 := GenS15{F0: 42, F1: (v3 % ((0) | 1)), F2: (fi93 ^ fi22)}
                                                                _ = v137
                                                            }
                                                            v102 += -128
                                                        }
                                                        var v138 int = fi93
                                                        v1 = 65535
                                                        _ = v138
                                                    } else {
                                                        v4 = v4
                                                        v139 := GenS17{F0: ("hello" + "\n\t")}
                                                        v2 += (fi93 + (fi22 >> ((4294967295) & 63)))
                                                        fmt.Println(v2)
                                                        switch 738 {
                                                        case 2147483647:
                                                            v102 = v102
                                                        case -984:
                                                            _ = len(v20)
                                                        case -875:
                                                            var v140 uint16 = (v1 | v1)
                                                            _ = v140
                                                        default:
                                                            var v141 string = ""
                                                            _ = v141
                                                        }
                                                        v4 = false
                                                        var v142 int32 = 1
                                                        v102 -= v102
                                                        var v143 uint16 = 65535
                                                        _ = v139
                                                        _ = v142
                                                        _ = v143
                                                    }
                                                case 2147483647:
                                                    v1 += (v1 * (v1 & 0))
                                                }
                                                v102 = (^(v102))
                                            } else {
                                                var v144 uint = 65535
                                                if (1.5 == float64(((0.1 * (1e10 / (((0.1 / (-2.25 + -2.25)) * (0.0 - 1e10)) - float32(1e10)))) - -2.25))) {
                                                    v21 -= (v21 - v21)
                                                } else {
                                                    if (!(!(("hello" + "\n\t") < ""))) {
                                                        _ = len(v20)
                                                        fmt.Println(v21, fi93, fi22)
                                                        v20 = append(v20, v2)
                                                    } else {
                                                        if (!(int64(v21) <= (v3 - (v3 + v3)))) {
                                                            var v145 uint64 = 18446744073709551615
                                                            _ = v145
                                                        } else {
                                                            v102 -= (v102 - (v102 >> ((255) & 63)))
                                                            var v146 bool = true
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                            _ = v146
                                                        }
                                                        var v147 int8 = -20
                                                        fmt.Println(fi22)
                                                        v2 = (fi22 % (((v2 - (fi22 | (^(v2))))) | 1))
                                                        var v148 float64 = 3.14159
                                                        _ = v147
                                                        _ = v148
                                                    }
                                                    v2 = 32767
                                                }
                                                v149 := GenS18{F0: 2147483647}
                                                _ = v144
                                                _ = v149
                                            }
                                        default:
                                            var v150 uint64 = 4294967295
                                            _ = v150
                                        }
                                    } else {
                                        fmt.Println(v2)
                                        v2 += int(v102)
                                        switch (-(v2)) {
                                        case 42:
                                            v102 = (v102 | (v102 & 100))
                                        }
                                        if len(v20) > 0 {
                                            v20[0] = fi22
                                        }
                                    }
                                case 882:
                                    var v151 int = fi22
                                    _ = v151
                                case 32768:
                                    if ((^(v3)) != -775) {
                                        var v152 uint = 65535
                                        _ = v152
                                    }
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        v2 = 55
                                    case "a":
                                        v153 := GenS19{F0: 2147483647, F1: false, F2: 0}
                                        _ = v153
                                    }
                                    if (-1 != genFn0(0, 2, v3)) {
                                        fmt.Println(fi22, v102, v4)
                                    }
                                    var v154 int8 = 1
                                    _ = v154
                                }
                                var v155 uint16 = uint16(v21)
                                v3 = v3
                                _ = v155
                            }
                            switch -218 {
                            case 32767:
                                v102 = (-(v102))
                            default:
                                v156 := GenS6{F0: (v21 | (v21 - v21)), F1: 2}
                                _ = v156
                            }
                        default:
                            if v4 {
                                v157 := GenS8{F0: 524}
                                fmt.Println(v3, fi93)
                                _ = v157
                            }
                        }
                    }
                    _ = v102
                }
                _ = fi93
            }
        case 0:
            var v158 float32 = 3.14159
            _ = v158
        case 9223372036854775807:
            switch (("\n\t" + ((("foo bar" + "") + "") + ("foo bar" + ""))) + (("\n\t" + "") + "a")) {
            case "":
                v159 := GenS20{F0: int32(v1), F1: 18446744073709551615, F2: 403}
                _ = v159
            case "\n\t":
                v3 += v3
            case "hello":
                var v160 int = (v2 ^ fi22)
                _ = v160
            }
        }
        fmt.Println(fi22)
        _ = fi22
    }
    v2 += v2
    var v161 uint = 18446744073709551615
    for fi162 := 0; fi162 < 5; fi162++ {
        if v4 {
            fmt.Println(v4)
            fmt.Println(v3, v4, v4)
            switch (fi162 | genFn0((-(fi162)), 1, 101)) {
            case 32768:
                for fi163 := 0; fi163 < 9; fi163++ {
                    var v164 float32 = 1.5
                    v4 = (42 != 624)
                    v161 = v161
                    if (v1 <= v1) {
                        fmt.Println(v21)
                        switch fi163 {
                        case -441:
                            if v4 {
                                v165 := GenS5{F0: (fi162 >> ((v161) & 63)), F1: (fi163 % ((fi163) | 1))}
                                switch (fi162 + v2) {
                                case -615:
                                    var v166 uint64 = 4294967295
                                    _ = v166
                                default:
                                    v167 := GenS21{F0: (fi163 - int(v21)), F1: 56}
                                    _ = v167
                                }
                                var v168 float32 = ((v164 * v164) - (3.14159 * 1e10))
                                _ = v165
                                _ = v168
                            } else {
                                if (((0.1 - float64(v164)) == 1.5) || (100 <= 1)) {
                                    v164 += ((1e10 + 1e10) - (-2.25 - v164))
                                }
                            }
                        }
                    }
                    _ = fi163
                    _ = v164
                }
            }
        } else {
            fmt.Println(v4, v2)
        }
        _ = fi162
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v161:", v161)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v161
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
