package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint; F1 uint64; F2 string }
type GenS12 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS13 struct { F0 uint64; F1 uint16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS17 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int) int64 {
    v7 := []int{}
    if false {
        p0 = 32767
        fmt.Println(p0)
    }
    if (!(0 == 2147483647)) {
        if true {
            switch -540 {
            case 32767:
                if true {
                    for fi8 := 0; fi8 < 12; fi8++ {
                        var v9 uint = 42
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v10 := GenS1{F0: -2147483648, F1: 1}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    switch "hello" {
                    case "a":
                        switch p0 {
                        case 4294967295:
                            var v11 uint8 = 127
                            _ = v11
                        case -2147483648:
                            p0 = (p0 - (p0 | (^(p0))))
                        case 32768:
                            p0 = (p0 ^ (p0 % (((p0 ^ (p0 - 128))) | 1)))
                        default:
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                fmt.Println(p0)
                            case "\n\t":
                                fmt.Println(p0)
                            case "hello":
                                var v12 float32 = -2.25
                                _ = v12
                            default:
                                p0 -= p0
                            }
                        }
                    }
                } else {
                    v13 := GenS2{F0: (p0 + (^(p0))), F1: 693, F2: 2}
                    _ = v13
                }
            case 897:
                fmt.Println(p0)
            case 1:
                fmt.Println(p0)
            }
            switch 1 {
            case 2147483647:
                v7 = append(v7, (p0 * (p0 / ((127) | 1))))
            default:
                fmt.Println(p0)
            }
            v14 := GenS3{F0: 49, F1: 694, F2: (^(p0))}
            _ = v14
        }
        if false {
            switch (("\n\t" + "hello") + (("\n\t" + "\n\t") + "foo bar")) {
            case "hello":
                v15 := GenS3{F0: 0, F1: 65535, F2: (^(p0))}
                _ = v15
            default:
                var v16 float64 = -2.25
                _ = v16
            }
            if true {
                var v17 string = ""
                switch (p0 + p0) {
                case 4294967295:
                    fmt.Println(v17, p0)
                case 42:
                    var v18 int64 = -238
                    _ = v18
                default:
                    if false {
                        p0 = 32767
                        v19 := GenS2{F0: (p0 >> ((2) & 63)), F1: -9223372036854775808, F2: 2}
                        _ = v19
                    }
                }
                _ = v17
            }
        }
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 int64) int8 {
    p1 = p1
    p0 = p0
    _ = p0
    _ = p1
    return (p0 << ((42) & 63))
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = -9223372036854775808
    var v3 int32 = 2147483647
    var v4 uint64 = 18446744073709551615
    var v5 int = 1
    var v6 int64 = (v2 >> ((65535) & 63))
    v20 := []int{v5}
    var v21 bool = true
    v22 := GenS4{F0: genFn1(1, (v6 << ((18446744073709551615) & 63)))}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    switch "\n\t" {
    case "":
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 float32 = -2.25
            if (-7 >= -128) {
                v20 = append(v20, (fi23 | v5))
            } else {
                var v25 float32 = v24
                v20 = append(v20, v5)
                _ = v25
            }
            _ = fi23
            _ = v24
        }
    default:
        v20 = append(v20, (v5 & 2147483647))
    }
    fmt.Println(v1, v5, v3)
    fmt.Println(v21)
    switch "" {
    case "\n\t":
        _ = len(v20)
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        _ = len(v20)
        var v27 uint32 = 42
        _ = fi26
        _ = v27
    }
    v28 := GenS5{F0: (v5 - (-(v5))), F1: -629, F2: 21}
    switch (("foo bar" + "") + ("foo bar" + "hello")) {
    case "hello":
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 float32 = 3.14159
            _ = fi29
            _ = v30
        }
    case "foo bar":
        switch (v5 + v5) {
        case -657:
            v6 = genFn0(v5)
        case 9223372036854775807:
            if ("foo bar" == (("" + "") + "hello")) {
                v2 = genFn0(2147483647)
                v20 = append(v20, int(v6))
                v31 := GenS6{F0: 42}
                _ = v31
            } else {
                var v32 int16 = 42
                _ = v32
            }
        case 628:
            if v21 {
                v3 = (^(v3))
                v33 := GenS2{F0: 9223372036854775807, F1: (^(v2)), F2: 799}
                for fi34 := 0; fi34 < 12; fi34++ {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    var v35 float32 = float32(0.1)
                    switch (fi34 / ((v5) | 1)) {
                    case -976:
                        if true {
                            for fi36 := 0; fi36 < 12; fi36++ {
                                if v21 {
                                    v37 := GenS1{F0: v3, F1: 148}
                                    _ = v37
                                } else {
                                    if (!v21) {
                                        var v38 float32 = v35
                                        v39 := GenS5{F0: 4294967295, F1: 32767, F2: v4}
                                        fmt.Println(v3)
                                        _ = v38
                                        _ = v39
                                    }
                                    if ((1e10 + float32(float64((v35 * -2.25)))) >= float32(-2.25)) {
                                        var v40 string = "hello"
                                        var v41 float64 = 1.5
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 bool = (!v21)
                                    _ = v42
                                }
                                v35 += v35
                                switch (v5 % ((v5) | 1)) {
                                case 696:
                                    switch "" {
                                    case "":
                                        var v43 uint64 = 0
                                        _ = v43
                                    }
                                case 612:
                                    if v21 {
                                        _ = len(v20)
                                        if v21 {
                                            v5 += (v5 >> ((382) & 63))
                                            v21 = (v21 || (101 == 2))
                                        } else {
                                            var v44 int8 = -32
                                            _ = v44
                                        }
                                    }
                                case 9223372036854775807:
                                    if (825 <= 32767) {
                                        switch ("hello" + "\n\t") {
                                        case "\n\t":
                                            if (!v21) {
                                                switch v5 {
                                                case 368:
                                                    if (float64(0.1) == -2.25) {
                                                        if len(v20) > 0 {
                                                            v20[0] = fi34
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v45 bool = v21
                                                            _ = v45
                                                        case "\n\t":
                                                            switch -1 {
                                                            case 9223372036854775807:
                                                                switch "" {
                                                                case "\n\t":
                                                                    v35 = v35
                                                                case "foo bar":
                                                                    var v46 float32 = float32(0.1)
                                                                    _ = v46
                                                                default:
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                }
                                                                var v47 int32 = (v3 << ((4294967295) & 63))
                                                                if ((-(v2)) < genFn0(1)) {
                                                                    fmt.Println(v47, fi34)
                                                                    var v48 float64 = 1e10
                                                                    _ = v48
                                                                } else {
                                                                    if (v21 || true) {
                                                                        if len(v20) > 0 {
                                                                            v20[0] = 927
                                                                        }
                                                                    } else {
                                                                        if (((^(v3)) < -74) || ((1.5 / v35) == v35)) {
                                                                            if v21 {
                                                                                v5 += -2147483648
                                                                            }
                                                                        }
                                                                        v2 += -9223372036854775808
                                                                    }
                                                                }
                                                                _ = v47
                                                            case -9223372036854775808:
                                                                switch fi36 {
                                                                case -892:
                                                                    fmt.Println(v35, v1, v1)
                                                                case 77:
                                                                    if (fi36 == v5) {
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                        switch 1 {
                                                                        case 700:
                                                                            fmt.Println(v1)
                                                                        }
                                                                    } else {
                                                                        v3 += (v3 / ((v3) | 1))
                                                                        switch fi36 {
                                                                        case -319:
                                                                            v49 := GenS7{F0: 1, F1: 1000}
                                                                            _ = v49
                                                                        case -2147483648:
                                                                            var v50 float32 = 1e10
                                                                            _ = v50
                                                                        case 537:
                                                                            v4 -= (^(v4))
                                                                        }
                                                                    }
                                                                case -2147483648:
                                                                    _ = len(v20)
                                                                default:
                                                                    switch "hello" {
                                                                    case "":
                                                                        v51 := GenS8{F0: v5}
                                                                        _ = v51
                                                                    case "a":
                                                                        v52 := GenS3{F0: 65535, F1: 538, F2: fi36}
                                                                        _ = v52
                                                                    case "foo bar":
                                                                        if len(v20) > 0 {
                                                                            v20[0] = v5
                                                                        }
                                                                    }
                                                                }
                                                            case -372:
                                                                v53 := GenS9{F0: -1, F1: fi34}
                                                                _ = v53
                                                            }
                                                        case "":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v54 int = (^(fi34))
                                                                _ = v54
                                                            case "\n\t":
                                                                v2 = 42
                                                                var v55 float32 = v35
                                                                fmt.Println(v21)
                                                                _ = v55
                                                            default:
                                                                switch (fi36 * v5) {
                                                                case 127:
                                                                    fmt.Println(v6, v1, fi36)
                                                                case -634:
                                                                    var v56 float64 = ((0.1 + ((3.14159 / ((1e10 * -2.25) - (3.14159 + 3.14159))) - ((float64(v35) / (1e10 + 3.14159)) - 1.5))) - (-2.25 - -2.25))
                                                                    _ = v56
                                                                case 0:
                                                                    var v57 uint32 = v1
                                                                    _ = v57
                                                                }
                                                            }
                                                        }
                                                        v58 := GenS9{F0: 1, F1: fi34}
                                                        _ = v58
                                                    } else {
                                                        v59 := GenS10{F0: 0}
                                                        _ = v59
                                                    }
                                                case 712:
                                                    v60 := GenS3{F0: 0, F1: 150, F2: (^(fi34))}
                                                    _ = v60
                                                }
                                                var v61 bool = (int32(v4) <= (v3 / ((v3) | 1)))
                                                if true {
                                                    if len(v20) >= 2 {
                                                        v20 = v20[1:len(v20)]
                                                    }
                                                    v62 := GenS11{F0: 2, F1: (v4 >> ((676) & 63)), F2: ""}
                                                    if (!(v61 || v21)) {
                                                        if v61 {
                                                            var v63 uint16 = 89
                                                            if true {
                                                                if v61 {
                                                                    var v64 float32 = ((v35 / 0.1) / float32(1e10))
                                                                    v4 += (v4 % ((978) | 1))
                                                                    v65 := GenS9{F0: 0, F1: fi34}
                                                                    v20 = append(v20, fi34)
                                                                    v4 = v4
                                                                    if ((v61 && (0 == 445)) && v61) {
                                                                        var v66 float32 = float32(3.14159)
                                                                        _ = v66
                                                                    }
                                                                    v64 = float32((1.5 / (0.1 / (0.1 + float64(v64)))))
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                                if ((^(v4)) < (^(v4))) {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        if false {
                                                                            if false {
                                                                                var v67 uint16 = (v63 + v63)
                                                                                fmt.Println(v5, v63, v67)
                                                                                if true {
                                                                                    fmt.Println(v3)
                                                                                    var v68 uint8 = 0
                                                                                    v5 -= (fi34 - int(v3))
                                                                                    var v69 int8 = 79
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        if v21 {
                                                                                            var v70 int64 = 0
                                                                                            if ((v63 << ((65535) & 63)) < v67) {
                                                                                                v71 := GenS6{F0: 665}
                                                                                                _ = v71
                                                                                            } else {
                                                                                                v4 = 1
                                                                                            }
                                                                                            _ = v70
                                                                                        }
                                                                                    case "":
                                                                                        v3 = (v3 >> ((4294967295) & 63))
                                                                                    default:
                                                                                        _ = len(v20)
                                                                                    }
                                                                                    var v72 float64 = 0.0
                                                                                    var v73 uint = 255
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                }
                                                                                _ = v67
                                                                            } else {
                                                                                v6 = 42
                                                                            }
                                                                            var v74 bool = ((v4 + (v4 + v4)) != (v4 / ((v4) | 1)))
                                                                            _ = v74
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                if v21 {
                                                                                    if len(v20) > 0 {
                                                                                        v20[0] = 1
                                                                                    }
                                                                                    v35 -= (((1e10 / v35) - v35) * (v35 * v35))
                                                                                    var v75 int16 = 42
                                                                                    _ = v75
                                                                                } else {
                                                                                    fmt.Println(v5)
                                                                                    if (v21 && (!true)) {
                                                                                        _ = len(v20)
                                                                                        v61 = v21
                                                                                    } else {
                                                                                        v2 += (v2 - int64(fi36))
                                                                                        if (32767 == 32767) {
                                                                                            var v76 uint8 = 19
                                                                                            v20 = append(v20, (v5 >> ((626) & 63)))
                                                                                            var v77 uint8 = v76
                                                                                            var v78 int16 = -32768
                                                                                            if (false && (v1 >= (v1 & v1))) {
                                                                                                fmt.Println(v77, fi34, v5)
                                                                                                v79 := GenS12{F0: 338, F1: (v76 % ((0) | 1)), F2: (v3 | (v3 ^ v3))}
                                                                                                _ = v79
                                                                                            } else {
                                                                                                var v80 int16 = 42
                                                                                                fmt.Println(v4, v1)
                                                                                                if false {
                                                                                                    var v81 int32 = v3
                                                                                                    fmt.Println(fi34)
                                                                                                    if v61 {
                                                                                                        switch "foo bar" {
                                                                                                        case "a":
                                                                                                            v6 += (-(v2))
                                                                                                        case "foo bar":
                                                                                                            v35 = v35
                                                                                                        default:
                                                                                                            var v82 int32 = 907
                                                                                                            _ = v82
                                                                                                        }
                                                                                                        fmt.Println(v80, v4, fi34)
                                                                                                    }
                                                                                                    _ = v81
                                                                                                }
                                                                                                _ = v80
                                                                                            }
                                                                                            var v83 int8 = 118
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                            _ = v83
                                                                                        }
                                                                                        if len(v20) >= 2 {
                                                                                            v20 = v20[1:len(v20)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                if (v21 || v61) {
                                                                                    switch ((("\n\t" + "a") + "foo bar") + ("\n\t" + (("foo bar" + ("hello" + "hello")) + "\n\t"))) {
                                                                                    case "a":
                                                                                        var v84 int = (fi34 >> ((4294967295) & 63))
                                                                                        _ = v84
                                                                                    default:
                                                                                        v21 = false
                                                                                    }
                                                                                    v85 := GenS13{F0: v4, F1: 610}
                                                                                    v86 := GenS3{F0: 0, F1: (v63 << ((351) & 63)), F2: (v5 & fi34)}
                                                                                    v87 := GenS10{F0: -42}
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                } else {
                                                                                    switch -1 {
                                                                                    case 4294967295:
                                                                                        fmt.Println(fi36)
                                                                                        switch (fi36 & 558) {
                                                                                        case 32767:
                                                                                            fmt.Println(v6, v5)
                                                                                        default:
                                                                                            v4 = v4
                                                                                        }
                                                                                    case 127:
                                                                                        var v88 int32 = int32(fi36)
                                                                                        _ = v88
                                                                                    case 673:
                                                                                        fmt.Println(v61, v3)
                                                                                    }
                                                                                    v20 = append(v20, (-(fi34)))
                                                                                    var v89 int32 = 135
                                                                                    _ = v89
                                                                                }
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v3, v3)
                                                                    case "foo bar":
                                                                        v90 := GenS1{F0: 42, F1: -1}
                                                                        _ = v90
                                                                    default:
                                                                        _ = len(v20)
                                                                    }
                                                                    v91 := GenS14{F0: 42}
                                                                    var v92 int8 = 127
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    var v93 float64 = float64(-2.25)
                                                                    var v94 int8 = -1
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                                v5 = 4294967295
                                                            }
                                                            _ = v63
                                                        } else {
                                                            v95 := GenS15{F0: 4294967295}
                                                            if false {
                                                                var v96 uint64 = (v4 - 42)
                                                                var v97 uint8 = 6
                                                                v98 := GenS16{F0: 42, F1: 65535, F2: uint16(v4)}
                                                                _ = v96
                                                                _ = v97
                                                                _ = v98
                                                            } else {
                                                                switch (v5 | (v5 | fi36)) {
                                                                case 717:
                                                                    if v61 {
                                                                        if (!((((((v5 > v5) || v61) || (v35 < v35)) && v21) || true) || true)) {
                                                                            switch 2147483647 {
                                                                            case 2147483647:
                                                                                v3 = v3
                                                                            }
                                                                        }
                                                                        switch fi34 {
                                                                        case 901:
                                                                            switch v5 {
                                                                            case 128:
                                                                                switch (fi34 | -9223372036854775808) {
                                                                                case 42:
                                                                                    fmt.Println(fi36, v1, v2)
                                                                                default:
                                                                                    var v99 uint32 = (^(v1))
                                                                                    _ = v99
                                                                                }
                                                                            case -299:
                                                                                v21 = false
                                                                            case -460:
                                                                                var v100 uint8 = 127
                                                                                _ = v100
                                                                            default:
                                                                                if len(v20) >= 2 {
                                                                                    v20 = v20[1:len(v20)]
                                                                                }
                                                                            }
                                                                        case 159:
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "\n\t":
                                                                                v2 = int64(v4)
                                                                            case "":
                                                                                var v101 string = (("" + "") + ("\n\t" + ("a" + (("hello" + ("\n\t" + "")) + "foo bar"))))
                                                                                _ = v101
                                                                            case "foo bar":
                                                                                v5 = (fi36 * v5)
                                                                            default:
                                                                                v102 := GenS17{F0: 1, F1: v4}
                                                                                _ = v102
                                                                            }
                                                                        default:
                                                                            var v103 float32 = v35
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                case 720:
                                                                    v20 = append(v20, 127)
                                                                }
                                                            }
                                                            _ = v95
                                                        }
                                                    }
                                                    _ = v62
                                                } else {
                                                    if true {
                                                        var v104 uint8 = 168
                                                        var v105 int32 = v3
                                                        v1 = v1
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                    var v106 uint8 = 0
                                                    _ = v106
                                                }
                                                _ = v61
                                            }
                                        default:
                                            v20 = append(v20, v5)
                                        }
                                    }
                                }
                                v2 -= int64(v1)
                                fmt.Println(v2)
                                _ = fi36
                            }
                        }
                    }
                    _ = fi34
                    _ = v35
                }
                _ = v33
            } else {
                _ = len(v20)
                if v21 {
                    fmt.Println(v4)
                    v107 := GenS16{F0: 127, F1: 540, F2: 1}
                    v5 -= v5
                    _ = v107
                } else {
                    if ((!v21) || (v5 != v5)) {
                        if (!true) {
                            for fi108 := 0; fi108 < 12; fi108++ {
                                var v109 uint64 = (v4 % (((v4 ^ v4)) | 1))
                                _ = fi108
                                _ = v109
                            }
                            v2 = -1
                        } else {
                            v110 := GenS5{F0: v5, F1: 473, F2: v4}
                            if false {
                                v21 = true
                            }
                            _ = v110
                        }
                        v6 = 365
                    }
                    if false {
                        if len(v20) > 0 {
                            v20[0] = 128
                        }
                        var v111 uint8 = 33
                        v1 -= (v1 + v1)
                        _ = v111
                    } else {
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        fmt.Println(v3, v3)
                    }
                    v4 = 4294967295
                }
            }
        default:
            fmt.Println(v3)
        }
    case "":
        v5 = v5
    default:
        var v112 int64 = (^(v2))
        _ = v112
    }
    if ((v6 * -1) <= v2) {
        var v113 int64 = (v2 * genFn0(v5))
        _ = v113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
