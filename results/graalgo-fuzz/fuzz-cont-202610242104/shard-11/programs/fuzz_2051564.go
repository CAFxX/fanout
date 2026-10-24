package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint16; F1 string; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 bool; F1 int16; F2 uint16 }
type GenS15 struct { F0 int16; F1 int32; F2 uint64 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int64 }
func genFn0(p0 int32) string {
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            if true {
                v9 := GenS1{F0: 4294967295, F1: fi7, F2: -1}
                _ = v9
            }
            switch 42 {
            case 0:
                p0 -= (p0 | 1)
            case 128:
                v10 := GenS1{F0: 42, F1: fi8, F2: 32767}
                _ = v10
            case 596:
                fmt.Println(fi7, fi8)
            default:
                v11 := []int{}
                _ = v11
            }
            v12 := GenS1{F0: 158, F1: fi8, F2: -32768}
            p0 = (p0 ^ p0)
            _ = fi8
            _ = v12
        }
        if false {
            fmt.Println(fi7, fi7)
            switch (fi7 | (fi7 | (fi7 * (fi7 << ((771) & 63))))) {
            case 761:
                var v13 bool = true
                _ = v13
            case -2147483648:
                var v14 bool = true
                _ = v14
            case -271:
                var v15 float32 = -2.25
                _ = v15
            default:
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 int16 = 1000
                    _ = fi16
                    _ = v17
                }
            }
        }
        _ = fi7
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS2{F0: 1, F1: p0}
        _ = fi18
        _ = v19
    }
    if (!true) {
        v20 := []int{32767, -2147483648}
        if false {
            for fi21 := 0; fi21 < 13; fi21++ {
                if true {
                    if (-386 != 0) {
                        fmt.Println(p0, fi21)
                    } else {
                        if false {
                            var v22 float32 = (float32((3.14159 + float64((1e10 - (float32(-2.25) / 1e10))))) * float32((3.14159 / 1.0)))
                            var v23 uint16 = 32767
                            _ = v22
                            _ = v23
                        }
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                    p0 = -400
                    var v24 float64 = -2.25
                    _ = v24
                }
                fmt.Println(p0, fi21)
                _ = fi21
            }
            var v25 int = -9223372036854775808
            p0 = p0
            _ = v25
        }
        _ = v20
    } else {
        if ((1.5 / 1.0) < 0.1) {
            var v26 bool = (18446744073709551615 == 926)
            _ = v26
        }
        if (false && (!("" < ("\n\t" + "a")))) {
            if (!(true && (!(("foo bar" != ("" + "hello")) || false)))) {
                fmt.Println(p0)
                for fi27 := 0; fi27 < 13; fi27++ {
                    if ((fi27 < fi27) || (true || (("foo bar" + "") >= "a"))) {
                        v28 := []int{fi27, (fi27 * fi27)}
                        var v29 float64 = (-2.25 * float64(1e10))
                        v30 := GenS3{F0: p0, F1: 2}
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    if true {
                        if true {
                            var v31 bool = true
                            fmt.Println(fi27, fi27)
                            v32 := GenS3{F0: (p0 >> ((42) & 63)), F1: 255}
                            _ = v31
                            _ = v32
                        }
                        v33 := []int{}
                        _ = v33
                    } else {
                        v34 := GenS4{F0: 721}
                        var v35 int = 522
                        v36 := GenS5{F0: 42}
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    _ = fi27
                }
                v37 := GenS2{F0: -1, F1: (p0 & (-(p0)))}
                _ = v37
            }
            v38 := []int{-570, 9223372036854775807, 793}
            switch 42 {
            case -2147483648:
                for fi39 := 0; fi39 < 13; fi39++ {
                    if (("" + "") < "a") {
                        _ = len(v38)
                    } else {
                        p0 = (p0 + 0)
                        var v40 float64 = ((((1.5 * -2.25) - (1.5 * (1.5 - ((0.0 / 0.1) + 3.14159)))) - (3.14159 - float64(1.5))) / (1e10 * (((1.5 / float64((float32(1e10) * -2.25))) / 1.0) + 1e10)))
                        _ = v40
                    }
                    _ = fi39
                }
            default:
                var v41 float64 = 0.1
                _ = v41
            }
            _ = v38
        } else {
            v42 := GenS6{F0: 9223372036854775807, F1: 127}
            _ = v42
        }
        p0 += -1
    }
    v43 := GenS7{F0: 240, F1: "hello", F2: 1000}
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 int32 = p0
        var v46 int8 = 104
        v47 := GenS8{F0: 1, F1: 221}
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    _ = p0
    _ = v43
    return "hello"
}
func genFn1() int8 {
    switch "\n\t" {
    case "\n\t":
        v48 := []int{1, -262, 32767, 32768}
        _ = v48
    default:
        var v49 int64 = 1
        _ = v49
    }
    var v50 uint = 42
    var v51 bool = (!(750 >= 2147483647))
    var v52 int16 = 32767
    if (!v51) {
        if (744 >= 4294967295) {
            v53 := GenS9{F0: 223}
            _ = v53
        }
        v54 := []int{127, -514}
        _ = v54
    }
    _ = v50
    _ = v51
    _ = v52
    return 100
}

func main() {
    var v1 int = 20
    var v2 uint16 = 858
    var v3 uint32 = 57
    var v4 uint16 = 42
    var v5 float32 = 0.0
    v6 := make(map[string]uint32)
    var v55 uint8 = 42
    if (!false) {
        if false {
            _ = len(v6)
            var v56 bool = true
            if (37 == 0) {
                v6[("hello" + "a")] = (^(v3))
                for fi57 := 0; fi57 < 15; fi57++ {
                    switch 128 {
                    case 127:
                        v5 -= ((v5 - ((3.14159 * v5) * v5)) / (v5 * ((v5 / v5) / (((3.14159 + v5) * v5) * (1e10 + (((v5 / 3.14159) / v5) - v5))))))
                    }
                    v58 := GenS1{F0: v3, F1: (v1 - v1), F2: -1}
                    switch "\n\t" {
                    case "a":
                        switch ("hello" + "a") {
                        case "hello":
                            switch "foo bar" {
                            case "\n\t":
                                switch "a" {
                                case "":
                                    v59 := GenS7{F0: v2, F1: "\n\t", F2: 0}
                                    _ = v59
                                case "hello":
                                    if true {
                                        switch (v1 ^ int(v55)) {
                                        case 579:
                                            var v60 string = ("\n\t" + "hello")
                                            _ = v60
                                        case -1:
                                            for fi61 := 0; fi61 < 4; fi61++ {
                                                v62 := []int{v1, fi57, -9223372036854775808}
                                                switch (fi57 >> ((1) & 63)) {
                                                case 9223372036854775807:
                                                    v62 = append(v62, fi57)
                                                }
                                                _ = fi61
                                                _ = v62
                                            }
                                        default:
                                            v5 -= 0.0
                                        }
                                        delete(v6, "\n\t")
                                    }
                                case "a":
                                    if (1 <= 65535) {
                                        var v63 float64 = 1.5
                                        v64 := GenS5{F0: 2}
                                        fmt.Println(fi57)
                                        _ = v63
                                        _ = v64
                                    } else {
                                        v65 := GenS2{F0: 1, F1: 0}
                                        v5 = (1e10 - v5)
                                        v66 := GenS10{F0: 1}
                                        _ = v65
                                        _ = v66
                                    }
                                }
                            case "hello":
                                var v67 float32 = (((float32(3.14159) + (1.5 * v5)) / 3.14159) * (v5 - 0.0))
                                _ = v67
                            }
                        case "":
                            for fi68 := 0; fi68 < 10; fi68++ {
                                if v56 {
                                    if (1 != 1000) {
                                        var v69 uint16 = v4
                                        var v70 float64 = 1e10
                                        switch v1 {
                                        case -701:
                                            fmt.Println(v55, fi57)
                                        }
                                        _ = v69
                                        _ = v70
                                    } else {
                                        v56 = (!false)
                                        var v71 uint32 = v3
                                        v2 = (v4 ^ (v4 / ((v2) | 1)))
                                        switch (v1 & 32768) {
                                        case -712:
                                            v72 := GenS11{F0: "a", F1: 1}
                                            _ = v72
                                        case 4294967295:
                                            var v73 string = (("a" + "foo bar") + "a")
                                            _ = v73
                                        }
                                        v1 = (fi68 | (fi68 - fi68))
                                        _ = v71
                                    }
                                    if (585 < 4294967295) {
                                        v74 := GenS12{F0: 0}
                                        _ = v74
                                    }
                                    if (!v56) {
                                        v75 := GenS13{F0: 583}
                                        _ = v75
                                    }
                                }
                                var v76 int32 = 649
                                switch fi68 {
                                case 630:
                                    v77 := GenS9{F0: v4}
                                    _ = v77
                                }
                                var v78 uint = 1
                                switch "" {
                                case "hello":
                                    if (65535 == v78) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v76 -= (^(v76))
                                        case "\n\t":
                                            var v79 int16 = 860
                                            _ = v79
                                        case "":
                                            if false {
                                                fmt.Println(v4, v78, v56)
                                            }
                                            v5 += (v5 - (((((v5 - (v5 * 3.14159)) - 0.0) * (v5 - (float32(1.5) * 3.14159))) * 1e10) + 1e10))
                                        default:
                                            _ = len(v6)
                                        }
                                    }
                                case "foo bar":
                                    v1 -= (v1 + -9223372036854775808)
                                default:
                                    v1 = (fi68 % ((v1) | 1))
                                }
                                _ = fi68
                                _ = v76
                                _ = v78
                            }
                        }
                    case "":
                        if v56 {
                            v4 = uint16(v55)
                        }
                    case "foo bar":
                        v80 := GenS1{F0: (v3 % (((v3 / ((v3) | 1))) | 1)), F1: v1, F2: 0}
                        _ = v80
                    }
                    _ = fi57
                    _ = v58
                }
                v81 := v6[""]
                _ = v81
            } else {
                for fi82 := 0; fi82 < 7; fi82++ {
                    v55 += 0
                    _ = fi82
                }
            }
            _ = v56
        } else {
            v4 = (v2 - (v4 / ((v2) | 1)))
        }
        var v83 float64 = float64(v5)
        var v84 uint64 = 1
        _ = v83
        _ = v84
    } else {
        v1 = (v1 + -2147483648)
        v85 := GenS5{F0: v55}
        _ = v85
    }
    if (1 >= 9223372036854775807) {
        var v86 int16 = 42
        _ = v86
    }
    v87 := GenS14{F0: false, F1: 1, F2: 974}
    var v88 float32 = v5
    _ = len(v6)
    v89 := GenS2{F0: 1000, F1: 42}
    v2 -= 32767
    v90 := []int{(v1 / (((^(v1))) | 1)), 9223372036854775807}
    if false {
        var v91 float64 = 1.5
        v4 = v2
        _ = v91
    }
    if (true && (!true)) {
        v92 := GenS15{F0: 824, F1: 42, F2: 2}
        v93 := GenS16{F0: 1}
        var v94 int8 = 127
        _ = v92
        _ = v93
        _ = v94
    } else {
        fmt.Println(v2, v1, v1)
    }
    var v95 string = "\n\t"
    if true {
        if (0 > 42) {
            var v96 int32 = 42
            fmt.Println(v3)
            v4 = (^(v2))
            _ = v96
        } else {
            for fi97 := 0; fi97 < 12; fi97++ {
                var v98 uint16 = v2
                delete(v6, v95)
                _ = fi97
                _ = v98
            }
        }
        if true {
            v90 = append(v90, v1)
            v99 := GenS17{F0: -9223372036854775808}
            _ = v99
        } else {
            var v100 int8 = -27
            if len(v90) >= 2 {
                v90 = v90[1:len(v90)]
            }
            switch 9223372036854775807 {
            case 4294967295:
                if true {
                    if len(v90) >= 2 {
                        v90 = v90[1:len(v90)]
                    }
                    v101 := GenS14{F0: false, F1: int16(v55), F2: (v4 % ((0) | 1))}
                    _ = v101
                }
            case -2147483648:
                v100 = -1
            case 42:
                v102 := GenS6{F0: -859, F1: v100}
                _ = v102
            }
            _ = v100
        }
        v95 = "\n\t"
    } else {
        switch ((("\n\t" + (v95 + ("" + v95))) + "\n\t") + (v95 + v95)) {
        case "foo bar":
            v103 := GenS8{F0: 128, F1: 337}
            _ = v103
        default:
            v104 := GenS2{F0: 0, F1: 490}
            _ = v104
        }
        var v105 int8 = -98
        _ = v105
    }
    for fi106 := 0; fi106 < 18; fi106++ {
        for fi107 := 0; fi107 < 18; fi107++ {
            _ = len(v6)
            _ = fi107
        }
        var v108 int8 = -128
        for fi109 := 0; fi109 < 5; fi109++ {
            fmt.Println(v88, fi109, v88)
            _ = fi109
        }
        _ = fi106
        _ = v108
    }
    switch "\n\t" {
    case "a":
        v110 := GenS16{F0: 65535}
        _ = v110
    default:
        if (745 == v3) {
            var v111 int64 = 0
            fmt.Println(v55, v5)
            _ = v111
        }
    }
    if len(v90) > 0 {
        v90[0] = v1
    }
    v1 -= (^(v1))
    fmt.Println(v1, v95, v95)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v88:", v88)
    fmt.Println("slicelen v90:", len(v90), v90)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v87
    _ = v88
    _ = v89
    _ = v90
    _ = v95
}
