package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int64; F1 int16; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string; F1 uint64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 uint64; F2 int8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 uint32; F2 uint }
type GenS15 struct { F0 bool; F1 int; F2 uint8 }
type GenS16 struct { F0 string; F1 int64 }
type GenS17 struct { F0 int8; F1 int32; F2 uint8 }
type GenS18 struct { F0 uint8; F1 string }
type GenS19 struct { F0 bool; F1 int16 }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    var v7 uint32 = 2
    var v8 int32 = 2147483647
    p0 += 648
    _ = p0
    _ = v7
    _ = v8
    return (v8 ^ (v8 << ((304) & 63)))
}
func genFn1(p0 uint8) uint32 {
    switch 19 {
    case 9223372036854775807:
        switch 127 {
        case 2147483647:
            if true {
                v9 := []int{-1}
                v10 := GenS1{F0: p0}
                _ = v10
                _ = v9
            }
        case 596:
            p0 += (^(p0))
        default:
            fmt.Println(p0)
        }
    case 2147483647:
        fmt.Println(p0)
    case -1:
        switch "hello" {
        case "foo bar":
            for fi11 := 0; fi11 < 2; fi11++ {
                p0 = (p0 * p0)
                if false {
                    fmt.Println(p0, p0)
                } else {
                    p0 -= 2
                    var v12 uint8 = 121
                    _ = v12
                }
                v13 := []int{(fi11 + (fi11 & -2147483648))}
                _ = fi11
                _ = v13
            }
        case "":
            switch "" {
            case "\n\t":
                p0 -= p0
            case "foo bar":
                p0 -= p0
            }
        case "\n\t":
            fmt.Println(p0)
        default:
            v14 := GenS1{F0: (p0 & p0)}
            _ = v14
        }
    default:
        var v15 bool = (p0 < p0)
        _ = v15
    }
    v16 := GenS2{F0: 1, F1: 18446744073709551615}
    v17 := []int{-477, 580}
    v18 := GenS1{F0: p0}
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return 2147483648
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{(("foo bar" + "hello") + "a"), ""}
    var v3 float64 = 3.14159
    var v4 float32 = (0.1 * 0.0)
    var v5 int64 = -334
    var v6 int64 = v5
    v3 -= v3
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS3{F0: 18446744073709551615, F1: v5}
        for fi21 := 0; fi21 < 12; fi21++ {
            v1[4294967295] = "\n\t"
            var v22 uint32 = 2147483648
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v5 = v5
    if (false || false) {
        if false {
            v5 = (v5 + (v5 << ((433) & 63)))
        } else {
            fmt.Println(v6, v4, v6)
            var v23 int32 = -980
            _ = v23
        }
        if false {
            if false {
                var v24 uint32 = 42
                var v25 uint32 = v24
                _ = v24
                _ = v25
            } else {
                if (889 <= 18446744073709551615) {
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                }
                if true {
                    v26 := GenS4{F0: 0, F1: false}
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        switch (("hello" + "a") + (((("foo bar" + ((("" + "hello") + "hello") + "")) + "hello") + "hello") + "a")) {
                        case "":
                            v28 := GenS5{F0: -619, F1: 42, F2: 141}
                            _ = v28
                        case "foo bar":
                            if (true && ((((-(v5)) == -817) && false) && true)) {
                                if (2 != 18446744073709551615) {
                                    v3 += (1e10 * (3.14159 / (v3 - ((v3 + v3) + float64(0.1)))))
                                    v29 := GenS6{F0: genFn1(128)}
                                    _ = v29
                                }
                            } else {
                                fmt.Println(v5, v3, fi27)
                                if false {
                                    v4 = (v4 + v4)
                                    if (2 == 18) {
                                        var v30 uint64 = 2
                                        v30 = v30
                                        _ = v30
                                    }
                                } else {
                                    var v31 uint32 = 2147483647
                                    var v32 uint16 = 0
                                    fmt.Println(v32)
                                    _ = v31
                                    _ = v32
                                }
                                if false {
                                    for fi33 := 0; fi33 < 18; fi33++ {
                                        var v34 float32 = float32(v3)
                                        _ = fi33
                                        _ = v34
                                    }
                                    for fi35 := 0; fi35 < 20; fi35++ {
                                        v36 := GenS3{F0: 992, F1: v5}
                                        var v37 int = (fi27 % (((fi35 >> ((0) & 63))) | 1))
                                        _ = len(v2)
                                        if true {
                                            var v38 uint64 = 328
                                            if (!((^(v38)) != v38)) {
                                                if true {
                                                    v37 += fi27
                                                    var v39 float64 = 0.1
                                                    switch (("a" + "a") + ("foo bar" + "\n\t")) {
                                                    case "hello":
                                                        _ = len(v1)
                                                    case "a":
                                                        v37 += fi35
                                                    default:
                                                        fmt.Println(v3, v6)
                                                    }
                                                    _ = v39
                                                }
                                                if false {
                                                    var v40 uint64 = v38
                                                    _ = v40
                                                }
                                                if true {
                                                    v6 = v5
                                                    switch "" {
                                                    case "hello":
                                                        if (!true) {
                                                            if true {
                                                                var v41 uint32 = 42
                                                                switch "\n\t" {
                                                                case "":
                                                                    if true {
                                                                        var v42 uint = 1
                                                                        var v43 int32 = 1
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        _ = v42
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 float32 = v4
                                                                        var v45 int8 = -128
                                                                        switch fi35 {
                                                                        case -712:
                                                                            v6 += v5
                                                                        case 642:
                                                                            var v46 uint16 = 137
                                                                            _ = v46
                                                                        default:
                                                                            v41 -= (v41 - v41)
                                                                        }
                                                                        _ = v44
                                                                        _ = v45
                                                                    }
                                                                case "a":
                                                                    var v47 int = (fi27 % ((2147483647) | 1))
                                                                    _ = v47
                                                                case "hello":
                                                                    if (2147483648 == genFn1(2)) {
                                                                        v48 := GenS7{F0: 2, F1: 745}
                                                                        v3 -= ((((v3 - (0.0 / v3)) + v3) - ((float64(v4) * 0.0) + -2.25)) - float64((v4 - v4)))
                                                                        switch (fi27 & 1) {
                                                                        case 32767:
                                                                            if (!true) {
                                                                                v41 = v41
                                                                            } else {
                                                                                var v49 bool = (318 == 726)
                                                                                _ = v49
                                                                            }
                                                                        default:
                                                                            v50 := GenS8{F0: 255, F1: v38}
                                                                            _ = v50
                                                                        }
                                                                        _ = v48
                                                                    }
                                                                default:
                                                                    var v51 float32 = 1.5
                                                                    _ = v51
                                                                }
                                                                v4 = v4
                                                                if ("" == "hello") {
                                                                    if (2 == 1) {
                                                                        _ = len(v2)
                                                                    } else {
                                                                        v37 = v37
                                                                        switch "a" {
                                                                        case "":
                                                                            if (!false) {
                                                                                var v52 int64 = -983
                                                                                _ = v52
                                                                            } else {
                                                                                if ((fi27 >> ((65535) & 63)) == v37) {
                                                                                    var v53 int = (v37 & fi35)
                                                                                    _ = v53
                                                                                } else {
                                                                                    v6 = (v5 >> ((42) & 63))
                                                                                    v54 := GenS2{F0: 1000, F1: (v38 / ((v38) | 1))}
                                                                                    var v55 uint64 = (v38 / ((4294967295) | 1))
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            fmt.Println(v5, v4, v5)
                                                                            var v56 float32 = v4
                                                                            var v57 int16 = 1
                                                                            _ = v56
                                                                            _ = v57
                                                                        case "a":
                                                                            var v58 int8 = 1
                                                                            _ = v58
                                                                        default:
                                                                            var v59 uint16 = 2
                                                                            v37 = (fi27 / (((fi35 / (((fi35 & v37)) | 1))) | 1))
                                                                            _ = v59
                                                                        }
                                                                    }
                                                                }
                                                                _ = v41
                                                            }
                                                        } else {
                                                            var v60 uint8 = 42
                                                            if false {
                                                                switch (fi35 - (v37 ^ int(v5))) {
                                                                case 32767:
                                                                    fmt.Println(v6)
                                                                case -1:
                                                                    fmt.Println(v4)
                                                                }
                                                            }
                                                            _ = v60
                                                        }
                                                    case "foo bar":
                                                        switch (v37 / (((fi35 << ((764) & 63))) | 1)) {
                                                        case -954:
                                                            var v61 float64 = (0.0 - 0.1)
                                                            _ = v61
                                                        }
                                                    case "":
                                                        fmt.Println(fi27)
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = "\n\t"
                                                    }
                                                } else {
                                                    v1[0] = "\n\t"
                                                }
                                                if (3.14159 == 1e10) {
                                                    v3 -= ((v3 / 1.0) - ((0.1 + v3) / (v3 / (v3 * 1e10))))
                                                    v37 += fi35
                                                    v37 = fi35
                                                    v5 = -640
                                                }
                                            }
                                            _ = v38
                                        } else {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v6 -= 1
                                            case "hello":
                                                var v62 float32 = v4
                                                _ = v62
                                            }
                                        }
                                        var v63 float32 = v4
                                        v3 -= 1e10
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                        _ = v63
                                    }
                                    var v64 uint32 = 2147483647
                                    _ = v64
                                }
                            }
                        case "\n\t":
                            v4 += 3.14159
                        default:
                            if true {
                                v2 = append(v2, "a")
                                if (69 <= 580) {
                                    _ = len(v2)
                                }
                                v3 += -2.25
                            } else {
                                switch 9223372036854775807 {
                                case 128:
                                    v3 += -2.25
                                default:
                                    switch (^(fi27)) {
                                    case 128:
                                        fmt.Println(v5, v5, v6)
                                    }
                                }
                                switch (("\n\t" + "") + (("foo bar" + ("a" + "foo bar")) + "")) {
                                case "foo bar":
                                    v3 = 1.5
                                }
                                v65 := GenS4{F0: 42, F1: false}
                                _ = v65
                            }
                        }
                        fmt.Println(v6, v4)
                        _ = fi27
                    }
                    if false {
                        var v66 uint32 = 2
                        _ = v66
                    } else {
                        for fi67 := 0; fi67 < 6; fi67++ {
                            v6 += (v5 / (((v5 * 313)) | 1))
                            if (0 < 85) {
                                if (!true) {
                                    var v68 string = ("\n\t" + ("a" + ""))
                                    v3 = v3
                                    if (802 >= 4294967295) {
                                        fmt.Println(v4, v5, v3)
                                        for fi69 := 0; fi69 < 6; fi69++ {
                                            var v70 uint64 = 42
                                            v5 -= v5
                                            var v71 bool = (0.1 >= (v3 / v3))
                                            switch fi69 {
                                            case 1:
                                                switch ("\n\t" + "foo bar") {
                                                case "a":
                                                    switch v68 {
                                                    case "":
                                                        v72 := GenS3{F0: 1, F1: 9223372036854775807}
                                                        _ = v72
                                                    }
                                                    v73 := v1[2147483647]
                                                    _ = v73
                                                case "\n\t":
                                                    v70 += v70
                                                default:
                                                    switch fi69 {
                                                    case 42:
                                                        var v74 bool = false
                                                        if (!v74) {
                                                            var v75 float32 = -2.25
                                                            v76 := GenS4{F0: 1, F1: ((!v74) && v74)}
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        var v77 int8 = 127
                                                        _ = v74
                                                        _ = v77
                                                    default:
                                                        var v78 uint64 = (v70 + v70)
                                                        _ = v78
                                                    }
                                                }
                                            default:
                                                var v79 string = v68
                                                _ = v79
                                            }
                                            v5 += (v6 + v5)
                                            _ = fi69
                                            _ = v70
                                            _ = v71
                                        }
                                        v4 += float32((v3 - v3))
                                    } else {
                                        switch "hello" {
                                        case "":
                                            v80 := GenS9{F0: 2}
                                            _ = v80
                                        case "\n\t":
                                            var v81 uint = 2
                                            _ = v81
                                        default:
                                            delete(v1, genFn1(0))
                                        }
                                        v82 := GenS10{F0: v68, F1: 2}
                                        _ = v82
                                    }
                                    _ = v68
                                }
                                for fi83 := 0; fi83 < 16; fi83++ {
                                    switch "" {
                                    case "foo bar":
                                        switch (^(fi67)) {
                                        case -498:
                                            v84 := GenS11{F0: (fi83 & fi67)}
                                            _ = v84
                                        case 4294967295:
                                            var v85 uint16 = 2
                                            _ = v85
                                        default:
                                            var v86 int64 = (v5 * (v6 ^ 92))
                                            _ = v86
                                        }
                                    }
                                    _ = fi83
                                }
                                fmt.Println(v4, v3)
                            } else {
                                v87 := GenS12{F0: (^(fi67)), F1: 18446744073709551615, F2: 93}
                                _ = v87
                            }
                            _ = fi67
                        }
                        delete(v1, 42)
                    }
                    v88 := GenS7{F0: 4294967295, F1: 340}
                    _ = v88
                }
                for fi89 := 0; fi89 < 20; fi89++ {
                    if false {
                        v3 = ((v3 * -2.25) - float64((1e10 + (0.1 / (v4 / (v4 / (1e10 / ((((((v4 + v4) + float32(0.1)) - (((v4 / v4) * v4) - 1.5)) - 1.5) - v4) * (0.1 * (((float32(v3) - v4) + (v4 / v4)) * float32(((-2.25 - v3) / float64(v4)))))))))))))
                        var v90 bool = true
                        switch "foo bar" {
                        case "":
                            switch 32768 {
                            case -471:
                                switch ("" + "") {
                                case "hello":
                                    if (65535 <= 2) {
                                        v5 -= v5
                                        fmt.Println(v4, fi89)
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = "\n\t"
                                        }
                                    }
                                case "foo bar":
                                    if v90 {
                                        if v90 {
                                            var v91 uint8 = 255
                                            var v92 uint = 18446744073709551615
                                            if false {
                                                if (3.14159 == (v4 / 3.14159)) {
                                                    if ((0.0 * 1e10) > (3.14159 - v3)) {
                                                        switch fi89 {
                                                        case -9223372036854775808:
                                                            v6 = (v6 % (((v6 | (v6 % ((v5) | 1)))) | 1))
                                                        case -918:
                                                            if (false && (v90 && v90)) {
                                                                v92 -= (v92 << ((594) & 63))
                                                                v2 = append(v2, "")
                                                            }
                                                        case -2147483648:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    }
                                                }
                                                v6 = v5
                                                switch (fi89 | (fi89 & int(v91))) {
                                                case 127:
                                                    var v93 uint = uint(v91)
                                                    _ = v93
                                                case 4294967295:
                                                    fmt.Println(v6, v4, v90)
                                                default:
                                                    v3 += ((v3 + v3) * v3)
                                                }
                                            } else {
                                                v91 += (v91 << (((^(v92))) & 63))
                                                var v94 int32 = genFn0(v5)
                                                _ = v94
                                            }
                                            _ = v91
                                            _ = v92
                                        }
                                        var v95 string = (("foo bar" + (("hello" + "") + ("hello" + ""))) + (("\n\t" + (("" + "") + "\n\t")) + ("" + "foo bar")))
                                        _ = v95
                                    } else {
                                        if v90 {
                                            if (1.5 < ((v4 + float32((v3 * -2.25))) - (1e10 * 0.0))) {
                                                if v90 {
                                                    for fi96 := 0; fi96 < 4; fi96++ {
                                                        if v90 {
                                                            if (0 < 42) {
                                                                var v97 int64 = v6
                                                                switch fi89 {
                                                                case 9223372036854775807:
                                                                    var v98 uint32 = 4294967295
                                                                    _ = v98
                                                                case 598:
                                                                    v4 = v4
                                                                case 2147483647:
                                                                    var v99 uint8 = 2
                                                                    _ = v99
                                                                }
                                                                if ((!v90) && v90) {
                                                                    if false {
                                                                        if ((float64(v4) - v3) <= (float64(v4) * v3)) {
                                                                            if (((255 <= 255) || true) && ((fi89 + fi89) <= (fi89 / (((fi89 / ((128) | 1))) | 1)))) {
                                                                                if false {
                                                                                    if (!v90) {
                                                                                        switch -86 {
                                                                                        case 0:
                                                                                            delete(v1, 823)
                                                                                        case -2147483648:
                                                                                            v3 -= 0.0
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            v100 := GenS13{F0: 255}
                                                                                            _ = v100
                                                                                        case "":
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                var v101 float32 = (v4 * 0.0)
                                                                                                _ = v101
                                                                                            }
                                                                                        default:
                                                                                            delete(v1, 4294967295)
                                                                                        }
                                                                                        if v90 {
                                                                                            v102 := GenS13{F0: 255}
                                                                                            var v103 uint64 = 787
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                v104 := GenS7{F0: v103, F1: 914}
                                                                                                _ = v104
                                                                                            default:
                                                                                                if false {
                                                                                                    if (-2147483648 != genFn0((v5 << ((65535) & 63)))) {
                                                                                                        if v90 {
                                                                                                            v6 = (v97 * v97)
                                                                                                        }
                                                                                                    } else {
                                                                                                        v2 = append(v2, ("a" + ""))
                                                                                                        v2 = append(v2, (("hello" + "hello") + ("\n\t" + ("a" + (("a" + ((("hello" + "\n\t") + "\n\t") + "")) + "\n\t")))))
                                                                                                        v105 := GenS13{F0: 2}
                                                                                                        v3 -= v3
                                                                                                        _ = v105
                                                                                                    }
                                                                                                    v5 = (v97 / (((v97 ^ 1000)) | 1))
                                                                                                } else {
                                                                                                    var v106 uint = 18446744073709551615
                                                                                                    _ = v106
                                                                                                }
                                                                                            }
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                        } else {
                                                                                            v90 = v90
                                                                                            v107 := GenS14{F0: fi96, F1: 2147483648, F2: 18446744073709551615}
                                                                                            v108 := v1[65535]
                                                                                            v109 := GenS15{F0: (v90 && v90), F1: (fi89 % (((fi96 - -460)) | 1)), F2: 242}
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                            _ = v109
                                                                                        }
                                                                                    }
                                                                                    var v110 uint = 913
                                                                                    _ = v110
                                                                                } else {
                                                                                    v5 = v5
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if v90 {
                                                                                v111 := v1[genFn1(127)]
                                                                                v112 := GenS12{F0: (fi89 | (fi96 / (((fi89 * fi89)) | 1))), F1: 0, F2: -1}
                                                                                fmt.Println(v111, v6)
                                                                                _ = v111
                                                                                _ = v112
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v3 = 1e10
                                                                        var v113 int64 = v5
                                                                        _ = v113
                                                                    }
                                                                    switch fi96 {
                                                                    case -652:
                                                                        if (!(42 <= 128)) {
                                                                            v5 = (v97 >> ((1) & 63))
                                                                            v90 = false
                                                                        } else {
                                                                            switch (^(fi96)) {
                                                                            case 128:
                                                                                if (!v90) {
                                                                                    if (true && (v90 && v90)) {
                                                                                        v97 += v6
                                                                                        var v114 uint32 = 2147483647
                                                                                        _ = v114
                                                                                    } else {
                                                                                        _ = len(v1)
                                                                                        fmt.Println(v4, v3)
                                                                                    }
                                                                                    fmt.Println(v3)
                                                                                    v4 = float32((-2.25 + ((v3 / v3) + v3)))
                                                                                } else {
                                                                                    v4 = 3.14159
                                                                                    v3 -= 1.5
                                                                                    v115 := v1[529]
                                                                                    _ = v115
                                                                                }
                                                                            default:
                                                                                fmt.Println(v6, v90)
                                                                            }
                                                                            v116 := GenS6{F0: 199}
                                                                            _ = v116
                                                                        }
                                                                    case 912:
                                                                        var v117 int64 = (v97 << ((42) & 63))
                                                                        _ = v117
                                                                    case 128:
                                                                        v3 += v3
                                                                    }
                                                                    delete(v1, 65535)
                                                                } else {
                                                                    if true {
                                                                        var v118 uint8 = 1
                                                                        _ = v118
                                                                    }
                                                                }
                                                                _ = v97
                                                            } else {
                                                                switch fi96 {
                                                                case -38:
                                                                    if v90 {
                                                                        if false {
                                                                            v119 := GenS1{F0: 0}
                                                                            _ = v119
                                                                        }
                                                                    } else {
                                                                        delete(v1, 42)
                                                                        v5 += (v5 & v5)
                                                                        v120 := GenS1{F0: 212}
                                                                        v4 -= (v4 - float32((((0.1 + v3) * ((((v3 / v3) - v3) + ((-2.25 - float64((3.14159 / ((v4 + 0.0) / v4)))) / (0.0 + 1e10))) - (1e10 + v3))) * (v3 + ((v3 * v3) / (0.1 / (v3 / v3)))))))
                                                                        v121 := GenS4{F0: 2, F1: (!(v6 >= 9223372036854775807))}
                                                                        _ = v120
                                                                        _ = v121
                                                                    }
                                                                case -310:
                                                                    v122 := GenS14{F0: 0, F1: genFn1(207), F2: 65535}
                                                                    _ = v122
                                                                case 1:
                                                                    switch ("\n\t" + "a") {
                                                                    case "a":
                                                                        v123 := GenS11{F0: fi96}
                                                                        _ = v123
                                                                    case "foo bar":
                                                                        fmt.Println(v4, v4, v5)
                                                                    case "\n\t":
                                                                        v124 := GenS16{F0: "hello", F1: (v5 - v6)}
                                                                        _ = v124
                                                                    default:
                                                                        if true {
                                                                            var v125 uint16 = 1
                                                                            v125 = (^(v125))
                                                                            if true {
                                                                                if (949 <= 2147483647) {
                                                                                    var v126 string = ("foo bar" + "foo bar")
                                                                                    fmt.Println(v4, v5, v126)
                                                                                    v5 += v5
                                                                                    v5 = (v5 << ((uint(v125)) & 63))
                                                                                    _ = v126
                                                                                }
                                                                            } else {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                v6 += (v6 << ((2) & 63))
                                                                                var v127 uint16 = v125
                                                                                _ = v127
                                                                            }
                                                                            var v128 float64 = v3
                                                                            _ = v125
                                                                            _ = v128
                                                                        } else {
                                                                            fmt.Println(v90, v90)
                                                                        }
                                                                    }
                                                                    v6 += -9223372036854775808
                                                                }
                                                                v3 = (float64(v4) * ((v3 + v3) / (v3 * (1e10 - (v3 / (v3 - 0.1))))))
                                                            }
                                                            var v129 bool = ((!v90) && (v90 || false))
                                                            _ = v129
                                                        }
                                                        _ = fi96
                                                    }
                                                    var v130 uint16 = 1
                                                    _ = v130
                                                }
                                                v131 := GenS8{F0: 219, F1: 2}
                                                var v132 uint16 = 499
                                                _ = v131
                                                _ = v132
                                            }
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "hello"
                                        }
                                    }
                                }
                            case 498:
                                v5 = v5
                            case 710:
                                v133 := GenS16{F0: (("foo bar" + (("" + ("hello" + "foo bar")) + (("" + ("a" + "")) + ("hello" + (("foo bar" + "foo bar") + "foo bar"))))) + "hello"), F1: 42}
                                _ = v133
                            }
                        case "\n\t":
                            fmt.Println(v90, v90, v90)
                        }
                        _ = v90
                    } else {
                        switch int(v6) {
                        case 1:
                            switch fi89 {
                            case -9223372036854775808:
                                v3 -= ((1.5 - ((((v3 + 0.0) / 3.14159) - (v3 + v3)) * (v3 * (v3 / (1.5 / float64(float32((v3 + (v3 + v3))))))))) - 0.1)
                            }
                        case 128:
                            v134 := GenS13{F0: 220}
                            _ = v134
                        case 805:
                            for fi135 := 0; fi135 < 18; fi135++ {
                                v136 := GenS16{F0: (("foo bar" + "foo bar") + (("a" + (("a" + ("a" + "")) + ("\n\t" + "a"))) + (((("" + "a") + "foo bar") + "a") + ("foo bar" + "hello")))), F1: (v5 | v5)}
                                if (1 == 0) {
                                    fmt.Println(fi89)
                                } else {
                                    v6 = (^(v5))
                                    v1[genFn1(7)] = "\n\t"
                                    var v137 int = (fi89 & 128)
                                    _ = v137
                                }
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                v5 = v6
                                _ = fi135
                                _ = v136
                            }
                        }
                        fmt.Println(fi89)
                    }
                    var v138 float32 = -2.25
                    _ = fi89
                    _ = v138
                }
            }
        }
    } else {
        if (v5 != v6) {
            v4 -= v4
            v139 := v1[genFn1(127)]
            v139 = "hello"
            _ = v139
        }
        v4 += 1.5
    }
    if (733 == -492) {
        v140 := GenS7{F0: 18446744073709551615, F1: 255}
        v3 -= ((-2.25 / v3) - v3)
        _ = v140
    }
    if true {
        if false {
            v2 = append(v2, "hello")
            _ = len(v2)
            v4 = 3.14159
        } else {
            fmt.Println(v3, v5)
            var v141 float64 = (((v3 * v3) / ((-2.25 / v3) - -2.25)) / v3)
            switch "hello" {
            case "hello":
                v1[0] = "a"
            default:
                v142 := GenS9{F0: 65535}
                _ = v142
            }
            _ = v141
        }
    } else {
        if true {
            v143 := GenS4{F0: 128, F1: true}
            var v144 bool = false
            _ = v143
            _ = v144
        }
        if len(v2) > 0 {
            v2[0] = ("a" + "")
        }
        var v145 string = "a"
        _ = v145
    }
    if true {
        v4 -= 3.14159
        for fi146 := 0; fi146 < 15; fi146++ {
            if (2 == 65535) {
                if (65535 != 42) {
                    v3 = 0.0
                    switch fi146 {
                    case 42:
                        if ((!true) && false) {
                            if (false || true) {
                                if true {
                                    if false {
                                        fmt.Println(v6)
                                        var v147 uint32 = 65535
                                        var v148 float64 = 3.14159
                                        _ = v147
                                        _ = v148
                                    } else {
                                        var v149 int32 = int32(fi146)
                                        _ = v149
                                    }
                                }
                            }
                            switch "foo bar" {
                            case "a":
                                v3 = v3
                            case "\n\t":
                                switch (((("\n\t" + "a") + "") + "foo bar") + ("" + "\n\t")) {
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "hello":
                                    var v150 string = "hello"
                                    _ = v150
                                default:
                                    var v151 bool = false
                                    _ = v151
                                }
                            default:
                                v5 -= (v6 << ((0) & 63))
                            }
                        }
                    case -411:
                        var v152 uint = 0
                        _ = v152
                    default:
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v3, v5)
                        default:
                            v153 := GenS17{F0: -15, F1: 2147483647, F2: 255}
                            _ = v153
                        }
                    }
                }
                var v154 bool = false
                _ = v154
            } else {
                if false {
                    if true {
                        fmt.Println(v4)
                        var v155 uint32 = 2
                        _ = v155
                    }
                } else {
                    var v156 int64 = v5
                    fmt.Println(v4, v5, fi146)
                    var v157 float64 = 0.0
                    _ = v156
                    _ = v157
                }
                if (1000 == 1) {
                    v6 += (v6 % ((0) | 1))
                    if len(v2) > 0 {
                        v2[0] = ("a" + "a")
                    }
                    switch fi146 {
                    case 42:
                        if (1e10 > v3) {
                            v3 = (v3 * v3)
                            fmt.Println(v4, v6, v6)
                            fmt.Println(fi146, fi146, v5)
                        }
                    default:
                        fmt.Println(v4, v6, v4)
                    }
                } else {
                    var v158 int64 = (v6 ^ (^(v5)))
                    _ = v158
                }
                if (true || false) {
                    fmt.Println(fi146, v5)
                    fmt.Println(v6)
                }
            }
            _ = fi146
        }
        v159 := GenS15{F0: true, F1: 32768, F2: 1}
        _ = v159
    }
    var v160 int = 9223372036854775807
    switch ("" + ("hello" + ("\n\t" + (("hello" + "hello") + ((("a" + "\n\t") + ("foo bar" + "foo bar")) + ""))))) {
    case "":
        v161 := GenS3{F0: 18446744073709551615, F1: v6}
        _ = v161
    default:
        v3 = (v3 - v3)
    }
    v3 = v3
    fmt.Println(v5, v4)
    if (!("hello" < "\n\t")) {
        for fi162 := 0; fi162 < 10; fi162++ {
            v4 = v4
            _ = fi162
        }
        switch (v160 + v160) {
        case 43:
            var v163 uint32 = 2147483648
            _ = v163
        }
        fmt.Println(v160, v5, v5)
    } else {
        fmt.Println(v5)
    }
    if (32767 < 32767) {
        v160 -= 388
        v164 := GenS4{F0: 115, F1: false}
        for fi165 := 0; fi165 < 12; fi165++ {
            v166 := GenS18{F0: 23, F1: "foo bar"}
            _ = fi165
            _ = v166
        }
        _ = v164
    } else {
        if false {
            v167 := GenS5{F0: (v6 % ((9223372036854775807) | 1)), F1: 1, F2: 186}
            _ = v167
        } else {
            v168 := GenS19{F0: (!true), F1: 42}
            _ = v168
        }
        var v169 bool = false
        var v170 int64 = 363
        _ = v169
        _ = v170
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v160:", v160)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
