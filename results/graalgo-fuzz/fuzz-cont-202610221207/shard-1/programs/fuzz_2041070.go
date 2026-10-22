package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 uint }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 int16; F1 int16; F2 bool }
type GenS11 struct { F0 int32; F1 uint32 }
type GenS12 struct { F0 int32; F1 int16; F2 bool }
type GenS13 struct { F0 string }
func genFn0(p0 uint) uint64 {
    fmt.Println(p0)
    switch -2147483648 {
    case 128:
        var v7 int16 = 434
        _ = v7
    case -1:
        p0 += 0
    case 691:
        var v8 uint32 = 65535
        _ = v8
    default:
        switch "" {
        case "hello":
            switch "foo bar" {
            case "\n\t":
                fmt.Println(p0)
            case "foo bar":
                fmt.Println(p0)
            default:
                var v9 uint32 = 0
                _ = v9
            }
        case "foo bar":
            p0 = (p0 / ((p0) | 1))
        default:
            var v10 uint32 = 42
            _ = v10
        }
    }
    var v11 uint64 = 0
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 int8 = 85
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v11
    return v11
}

func main() {
    v1 := []string{"hello", "a", ("" + ("hello" + ("hello" + "foo bar")))}
    var v2 bool = (772 == 0)
    var v3 uint64 = 697
    var v4 int = 339
    var v5 float64 = 1.5
    v6 := make(map[string]bool)
    v14 := GenS1{F0: 42}
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v3, v5, v2)
        var v16 uint16 = 198
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 0}
    _ = len(v6)
    var v18 uint64 = genFn0(0)
    if v2 {
        v3 = genFn0(255)
    }
    v18 = 18446744073709551615
    switch "a" {
    case "hello":
        if v2 {
            if v2 {
                v19 := GenS1{F0: 0}
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 uint8 = 241
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v18 += genFn0(42)
                        v23 := GenS2{F0: fi22, F1: 42}
                        _ = fi22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v6[(("a" + ("foo bar" + "\n\t")) + "\n\t")] = v2
                        v2 = v2
                        v2 = (!v2)
                        _ = fi24
                    }
                    _ = fi20
                    _ = v21
                }
                var v25 int32 = 1
                _ = v19
                _ = v25
            }
            v26 := GenS2{F0: v4, F1: -86}
            if false {
                v5 = (v5 * 0.1)
                for fi27 := 0; fi27 < 2; fi27++ {
                    if false {
                        if (!v2) {
                            v2 = ((v4 << ((42) & 63)) == 4294967295)
                        } else {
                            if true {
                                var v28 int = fi27
                                if (127 == 0) {
                                    v29 := GenS1{F0: -32768}
                                    fmt.Println(v3, v2)
                                    var v30 int64 = 136
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            } else {
                                var v31 uint16 = 2
                                _ = v31
                            }
                            _ = len(v1)
                            var v32 uint32 = 984
                            _ = v32
                        }
                        v5 = v5
                    } else {
                        switch (-(v4)) {
                        case -2147483648:
                            v1 = append(v1, "a")
                        case 42:
                            for fi33 := 0; fi33 < 20; fi33++ {
                                var v34 uint16 = 65535
                                var v35 uint16 = v34
                                switch (-(fi33)) {
                                case 1:
                                    var v36 uint16 = (v35 % ((v35) | 1))
                                    var v37 int = int(v3)
                                    var v38 int8 = 127
                                    switch "hello" {
                                    case "foo bar":
                                        var v39 string = "a"
                                        _ = v39
                                    case "hello":
                                        var v40 uint = 255
                                        _ = len(v6)
                                        switch "a" {
                                        case "\n\t":
                                            if v2 {
                                                var v41 int16 = -902
                                                fmt.Println(v5)
                                                _ = v41
                                            } else {
                                                v37 -= (^(v37))
                                                var v42 float32 = (-2.25 * 0.0)
                                                v3 += v3
                                                _ = v42
                                            }
                                        case "":
                                            switch "foo bar" {
                                            case "hello":
                                                fmt.Println(v2)
                                            }
                                        case "a":
                                            switch (^(fi27)) {
                                            case -1:
                                                var v43 int = (v4 + (v37 | (fi27 >> ((v40) & 63))))
                                                _ = v43
                                            case 128:
                                                var v44 int64 = 936
                                                _ = v44
                                            }
                                        default:
                                            delete(v6, (("\n\t" + "") + ((("" + "hello") + "a") + ("hello" + "hello"))))
                                        }
                                        _ = v40
                                    }
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                case -1:
                                    _ = len(v1)
                                default:
                                    switch "hello" {
                                    case "":
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    case "foo bar":
                                        var v45 bool = (65535 > 18446744073709551615)
                                        _ = v45
                                    case "hello":
                                        if v2 {
                                            if (v34 != (v35 << ((4294967295) & 63))) {
                                                switch (fi33 - fi27) {
                                                case 9223372036854775807:
                                                    v46 := GenS1{F0: 1}
                                                    _ = v46
                                                case 4294967295:
                                                    fmt.Println(v35)
                                                }
                                                _ = len(v6)
                                                if v2 {
                                                    var v47 int = (v4 / ((v4) | 1))
                                                    delete(v6, "\n\t")
                                                    _ = v47
                                                } else {
                                                    switch v4 {
                                                    case 869:
                                                        v48 := GenS1{F0: -1}
                                                        v1 = append(v1, "")
                                                        var v49 int = (fi27 + (v4 + fi27))
                                                        var v50 int32 = -95
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    case 32767:
                                                        v5 = v5
                                                    case -9223372036854775808:
                                                        v1 = append(v1, "a")
                                                    default:
                                                        v5 -= 1e10
                                                    }
                                                    var v51 float32 = 0.1
                                                    var v52 bool = v2
                                                    _ = v51
                                                    _ = v52
                                                }
                                            } else {
                                                v34 = 963
                                                var v53 int32 = -606
                                                switch (fi27 ^ (v4 % ((int(v18)) | 1))) {
                                                case 32768:
                                                    v18 = genFn0(4294967295)
                                                case 1:
                                                    if (9223372036854775807 == -9223372036854775808) {
                                                        var v54 bool = (18446744073709551615 <= 65535)
                                                        if v54 {
                                                            v2 = ((!(-128 != -128)) || ((((("" + "\n\t") + "\n\t") + "hello") + "hello") < "a"))
                                                            if ((42 < 40) || v54) {
                                                                v55 := GenS1{F0: -32768}
                                                                if v54 {
                                                                    var v56 uint8 = 127
                                                                    _ = v56
                                                                } else {
                                                                    if v2 {
                                                                        _ = len(v1)
                                                                        fmt.Println(v2, fi27, v34)
                                                                        var v57 string = (("\n\t" + ((("" + (("a" + "\n\t") + ("\n\t" + ("" + "hello")))) + "") + "foo bar")) + "\n\t")
                                                                        _ = v57
                                                                    } else {
                                                                        if false {
                                                                            var v58 uint = 0
                                                                            _ = v58
                                                                        } else {
                                                                            v59 := GenS3{F0: 2147483647}
                                                                            _ = v59
                                                                        }
                                                                    }
                                                                    v53 -= (v53 | 0)
                                                                }
                                                                _ = v55
                                                            }
                                                            var v60 uint16 = v34
                                                            v4 -= 128
                                                            var v61 uint = 1
                                                            switch 1 {
                                                            case 128:
                                                                var v62 uint = uint(v35)
                                                                _ = v62
                                                            }
                                                            v34 = v60
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            fmt.Println(v18, v2, v5)
                                                            switch int(v3) {
                                                            case 32768:
                                                                v1 = append(v1, "a")
                                                            case -1:
                                                                if v54 {
                                                                    v4 = v4
                                                                } else {
                                                                    fmt.Println(fi33, v5, v3)
                                                                    v6[("hello" + "foo bar")] = v2
                                                                    v3 -= 4294967295
                                                                    fmt.Println(v34, v5, v3)
                                                                    v18 -= genFn0(586)
                                                                }
                                                            case 0:
                                                                if (!v54) {
                                                                    v63 := v6[("foo bar" + ((("foo bar" + (("" + "a") + "foo bar")) + "a") + "\n\t"))]
                                                                    if true {
                                                                        var v64 string = ("foo bar" + "foo bar")
                                                                        if v2 {
                                                                            v5 += ((v5 + (((-2.25 / float64(1.5)) * ((1.5 / 3.14159) * 0.0)) + 0.1)) + 0.0)
                                                                        } else {
                                                                            v35 += 987
                                                                        }
                                                                        _ = v64
                                                                    } else {
                                                                        v65 := GenS1{F0: 0}
                                                                        if v63 {
                                                                            v35 = (v35 - v34)
                                                                            if v2 {
                                                                                var v66 uint8 = 2
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    var v67 int32 = v53
                                                                                    _ = v67
                                                                                }
                                                                                _ = v66
                                                                            } else {
                                                                                var v68 float32 = 1.5
                                                                                v69 := GenS2{F0: fi33, F1: -128}
                                                                                _ = v68
                                                                                _ = v69
                                                                            }
                                                                            v70 := GenS1{F0: 42}
                                                                            _ = v70
                                                                        } else {
                                                                            if true {
                                                                                if (42 != v53) {
                                                                                    v3 += v18
                                                                                    if v54 {
                                                                                        if false {
                                                                                            v2 = v2
                                                                                            switch (-(v4)) {
                                                                                            case 32767:
                                                                                                v71 := GenS2{F0: fi27, F1: 42}
                                                                                                _ = v71
                                                                                            default:
                                                                                                switch (-(fi27)) {
                                                                                                case -2147483648:
                                                                                                    switch (-(fi27)) {
                                                                                                    case 385:
                                                                                                        v34 += v34
                                                                                                    case 32768:
                                                                                                        if (!v63) {
                                                                                                            var v72 int = fi27
                                                                                                            _ = v72
                                                                                                        } else {
                                                                                                            var v73 int16 = 0
                                                                                                            v2 = v63
                                                                                                            _ = v73
                                                                                                        }
                                                                                                    case 0:
                                                                                                        v6["\n\t"] = ((1 != 2) || v63)
                                                                                                    default:
                                                                                                        v35 -= v35
                                                                                                    }
                                                                                                case 172:
                                                                                                    var v74 bool = (!v63)
                                                                                                    _ = v74
                                                                                                case 171:
                                                                                                    v54 = (3.14159 > 0.1)
                                                                                                default:
                                                                                                    if (65535 != (v34 / ((868) | 1))) {
                                                                                                        var v75 uint64 = genFn0(255)
                                                                                                        var v76 int32 = (v53 / (((v53 << ((255) & 63))) | 1))
                                                                                                        v5 -= ((v5 / (float64(-2.25) / float64(0.1))) * v5)
                                                                                                        _ = v75
                                                                                                        _ = v76
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            switch (fi27 & fi33) {
                                                                                            case 540:
                                                                                                fmt.Println(fi27)
                                                                                            case 4294967295:
                                                                                                if (!(true || v54)) {
                                                                                                    switch ("\n\t" + "\n\t") {
                                                                                                    case "a":
                                                                                                        var v77 int8 = -91
                                                                                                        _ = v77
                                                                                                    case "":
                                                                                                        var v78 float32 = 0.1
                                                                                                        _ = v78
                                                                                                    default:
                                                                                                        fmt.Println(v2, v53)
                                                                                                    }
                                                                                                    switch (fi33 | -680) {
                                                                                                    case 42:
                                                                                                        v53 = v53
                                                                                                        var v79 float32 = float32((v5 / v5))
                                                                                                        _ = v79
                                                                                                    case 1:
                                                                                                        if (-128 <= 42) {
                                                                                                            fmt.Println(v34, v2, v3)
                                                                                                            v18 = (v3 * (v3 / (((v3 << ((577) & 63))) | 1)))
                                                                                                            if true {
                                                                                                                fmt.Println(v54)
                                                                                                                fmt.Println(fi33)
                                                                                                            } else {
                                                                                                                switch ("foo bar" + "\n\t") {
                                                                                                                case "hello":
                                                                                                                    if (191 == 912) {
                                                                                                                        switch "foo bar" {
                                                                                                                        case "foo bar":
                                                                                                                            v3 += v3
                                                                                                                        case "\n\t":
                                                                                                                            _ = len(v6)
                                                                                                                        case "hello":
                                                                                                                            v80 := GenS3{F0: 1}
                                                                                                                            _ = v80
                                                                                                                        default:
                                                                                                                            var v81 int16 = -883
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                        var v82 int8 = 55
                                                                                                                        _ = v82
                                                                                                                    } else {
                                                                                                                        var v83 uint64 = (v3 - 2)
                                                                                                                        if v54 {
                                                                                                                            var v84 string = ("foo bar" + "hello")
                                                                                                                            if (1 <= 100) {
                                                                                                                                var v85 uint32 = 65535
                                                                                                                                if v54 {
                                                                                                                                    if v2 {
                                                                                                                                        v3 = uint64(v35)
                                                                                                                                        v86 := GenS4{F0: 1, F1: -1, F2: 833}
                                                                                                                                        _ = v86
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v85
                                                                                                                            } else {
                                                                                                                                v87 := GenS5{F0: uint8(v35), F1: (^(v18))}
                                                                                                                                var v88 uint = uint(v83)
                                                                                                                                _ = v87
                                                                                                                                _ = v88
                                                                                                                            }
                                                                                                                            if v54 {
                                                                                                                                if v63 {
                                                                                                                                    var v89 float32 = float32(v5)
                                                                                                                                    _ = v89
                                                                                                                                }
                                                                                                                                var v90 string = ((((((("foo bar" + "") + (v84 + ((v84 + v84) + ((v84 + v84) + v84)))) + (v84 + ("" + v84))) + ((("foo bar" + "foo bar") + v84) + (v84 + v84))) + (v84 + (v84 + (v84 + (v84 + v84))))) + (v84 + v84)) + v84)
                                                                                                                                switch (v84 + (v84 + "hello")) {
                                                                                                                                case "hello":
                                                                                                                                    v6["foo bar"] = (v90 > v90)
                                                                                                                                case "a":
                                                                                                                                    v84 = (("" + v90) + v84)
                                                                                                                                case "":
                                                                                                                                    v91 := GenS6{F0: "foo bar", F1: v90}
                                                                                                                                    _ = v91
                                                                                                                                }
                                                                                                                                v53 = v53
                                                                                                                                if true {
                                                                                                                                    if v54 {
                                                                                                                                        v84 = "\n\t"
                                                                                                                                        var v92 uint16 = v35
                                                                                                                                        var v93 int32 = v53
                                                                                                                                        _ = v92
                                                                                                                                        _ = v93
                                                                                                                                    }
                                                                                                                                    var v94 float64 = 1.5
                                                                                                                                    if v2 {
                                                                                                                                        if v54 {
                                                                                                                                            v95 := GenS7{F0: (4294967295 != 18446744073709551615), F1: 1}
                                                                                                                                            var v96 int64 = 0
                                                                                                                                            if (4294967295 != 2) {
                                                                                                                                                if (0 <= 42) {
                                                                                                                                                    if v2 {
                                                                                                                                                        v96 += v96
                                                                                                                                                        var v97 int8 = 6
                                                                                                                                                        _ = v97
                                                                                                                                                    }
                                                                                                                                                    v84 = v90
                                                                                                                                                    v84 = (v84 + "hello")
                                                                                                                                                }
                                                                                                                                                fmt.Println(v35, v53, fi33)
                                                                                                                                            } else {
                                                                                                                                                switch v90 {
                                                                                                                                                case "hello":
                                                                                                                                                    v98 := GenS3{F0: 2147483647}
                                                                                                                                                    fmt.Println(v90, v53)
                                                                                                                                                    v2 = (32767 >= 877)
                                                                                                                                                    v99 := GenS7{F0: true, F1: (v53 * (v53 | -689))}
                                                                                                                                                    _ = v98
                                                                                                                                                    _ = v99
                                                                                                                                                case "":
                                                                                                                                                    if true {
                                                                                                                                                        if true {
                                                                                                                                                            if len(v1) > 0 {
                                                                                                                                                                v1[0] = (v90 + "hello")
                                                                                                                                                            }
                                                                                                                                                            var v100 int16 = -218
                                                                                                                                                            v34 += (v34 - v35)
                                                                                                                                                            _ = v100
                                                                                                                                                        } else {
                                                                                                                                                            fmt.Println(v3, v4)
                                                                                                                                                            _ = len(v1)
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                default:
                                                                                                                                                    v90 = "foo bar"
                                                                                                                                                }
                                                                                                                                                fmt.Println(fi27)
                                                                                                                                            }
                                                                                                                                            _ = v95
                                                                                                                                            _ = v96
                                                                                                                                        } else {
                                                                                                                                            if (!(2 != genFn0(42))) {
                                                                                                                                                v101 := GenS8{F0: v90, F1: v2}
                                                                                                                                                _ = v101
                                                                                                                                            }
                                                                                                                                            var v102 bool = v63
                                                                                                                                            v103 := GenS9{F0: 255, F1: 0}
                                                                                                                                            _ = v102
                                                                                                                                            _ = v103
                                                                                                                                        }
                                                                                                                                        v18 += (v18 % (((v3 ^ (v18 << ((18446744073709551615) & 63)))) | 1))
                                                                                                                                    } else {
                                                                                                                                        v53 = (-(v53))
                                                                                                                                    }
                                                                                                                                    if (v63 || v54) {
                                                                                                                                        v104 := GenS7{F0: true, F1: v53}
                                                                                                                                        var v105 float32 = 0.1
                                                                                                                                        if (v34 >= (v35 % ((v35) | 1))) {
                                                                                                                                            var v106 int32 = v53
                                                                                                                                            if false {
                                                                                                                                                var v107 int64 = 969
                                                                                                                                                v5 -= float64((v105 * 1.5))
                                                                                                                                                if (v34 != (^(v35))) {
                                                                                                                                                    switch (fi33 | v4) {
                                                                                                                                                    case -1:
                                                                                                                                                        _ = len(v1)
                                                                                                                                                    case 127:
                                                                                                                                                        var v108 string = v84
                                                                                                                                                        _ = v108
                                                                                                                                                    case -2147483648:
                                                                                                                                                        var v109 uint16 = 42
                                                                                                                                                        _ = v109
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                _ = v107
                                                                                                                                            }
                                                                                                                                            v110 := GenS2{F0: fi27, F1: -122}
                                                                                                                                            _ = v106
                                                                                                                                            _ = v110
                                                                                                                                        }
                                                                                                                                        _ = v104
                                                                                                                                        _ = v105
                                                                                                                                    } else {
                                                                                                                                        if ((v53 & -970) > v53) {
                                                                                                                                            delete(v6, ((("\n\t" + v84) + v90) + v90))
                                                                                                                                            v2 = v2
                                                                                                                                            v83 = 1
                                                                                                                                            if (v2 || (42 >= 42)) {
                                                                                                                                                v90 = "\n\t"
                                                                                                                                                _ = len(v6)
                                                                                                                                                var v111 uint16 = (^(v34))
                                                                                                                                                _ = v111
                                                                                                                                            } else {
                                                                                                                                                _ = len(v6)
                                                                                                                                                var v112 float32 = -2.25
                                                                                                                                                _ = v112
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v94
                                                                                                                                }
                                                                                                                                v113 := GenS4{F0: 32767, F1: 42, F2: 4294967295}
                                                                                                                                _ = v113
                                                                                                                                _ = v90
                                                                                                                            } else {
                                                                                                                                var v114 uint32 = 2147483648
                                                                                                                                if v2 {
                                                                                                                                    var v115 uint64 = (v83 ^ (v83 << ((841) & 63)))
                                                                                                                                    _ = v115
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v63)
                                                                                                                                }
                                                                                                                                var v116 float32 = (float32(float64((1e10 * (0.0 / float32((v5 + (float64(1e10) - v5))))))) / float32((float64(0.1) / 1.5)))
                                                                                                                                _ = v114
                                                                                                                                _ = v116
                                                                                                                            }
                                                                                                                            _ = v84
                                                                                                                        }
                                                                                                                        _ = v83
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    v3 -= (v3 & 1)
                                                                                                                }
                                                                                                            }
                                                                                                            v2 = (0.1 != float32(float64(3.14159)))
                                                                                                        }
                                                                                                    default:
                                                                                                        v117 := GenS10{F0: 712, F1: 1, F2: ((100 < 52) && false)}
                                                                                                        _ = v117
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v118 := GenS5{F0: 0, F1: (v3 & v3)}
                                                                                            _ = v118
                                                                                        } else {
                                                                                            if true {
                                                                                                v53 = -1
                                                                                            }
                                                                                            var v119 uint8 = 255
                                                                                            _ = v119
                                                                                        }
                                                                                        if v63 {
                                                                                            var v120 uint32 = 2147483647
                                                                                            var v121 float64 = v5
                                                                                            if (!v54) {
                                                                                                switch ((("\n\t" + "") + "foo bar") + ("a" + "hello")) {
                                                                                                case "foo bar":
                                                                                                    v3 += (v18 * 877)
                                                                                                case "\n\t":
                                                                                                    var v122 float32 = (((3.14159 * 3.14159) - (1.5 * (3.14159 - -2.25))) * -2.25)
                                                                                                    var v123 uint8 = 1
                                                                                                    if true {
                                                                                                        v6[("a" + "a")] = (("" + (("hello" + "") + (("" + "\n\t") + (("hello" + "a") + ("a" + (("\n\t" + "") + (((("a" + (("a" + "foo bar") + (("\n\t" + "\n\t") + "foo bar"))) + "") + "") + ("" + "a")))))))) == "foo bar")
                                                                                                        var v124 float64 = ((v5 - 3.14159) / v121)
                                                                                                        if v63 {
                                                                                                            v6[""] = ((v123 + 42) > (v123 ^ (v123 % (((v123 | (v123 >> ((4294967295) & 63)))) | 1))))
                                                                                                            v3 += v3
                                                                                                            v6[("hello" + "a")] = v63
                                                                                                            fmt.Println(v123, v123, v63)
                                                                                                        } else {
                                                                                                            var v125 int16 = -986
                                                                                                            v5 = v124
                                                                                                            var v126 int64 = 0
                                                                                                            _ = v125
                                                                                                            _ = v126
                                                                                                        }
                                                                                                        _ = v124
                                                                                                    } else {
                                                                                                        _ = len(v6)
                                                                                                    }
                                                                                                    _ = v122
                                                                                                    _ = v123
                                                                                                case "hello":
                                                                                                    if v2 {
                                                                                                        v127 := GenS8{F0: ("a" + "hello"), F1: true}
                                                                                                        _ = v127
                                                                                                    } else {
                                                                                                        v3 += v3
                                                                                                        v54 = v2
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v120
                                                                                            _ = v121
                                                                                        }
                                                                                    } else {
                                                                                        var v128 bool = (-233 >= -200)
                                                                                        delete(v6, "\n\t")
                                                                                        _ = v128
                                                                                    }
                                                                                    v63 = true
                                                                                } else {
                                                                                    if false {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            var v129 uint = 209
                                                                                            _ = v129
                                                                                        case "a":
                                                                                            if (true || v2) {
                                                                                                var v130 uint64 = (v18 << ((65535) & 63))
                                                                                                _ = v130
                                                                                            }
                                                                                        }
                                                                                    case "a":
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            v131 := GenS6{F0: "foo bar", F1: "\n\t"}
                                                                                            _ = v131
                                                                                        case "a":
                                                                                            v132 := GenS6{F0: "hello", F1: (("hello" + "a") + (("foo bar" + "") + ""))}
                                                                                            _ = v132
                                                                                        case "\n\t":
                                                                                            var v133 int = 168
                                                                                            switch fi33 {
                                                                                            case -1:
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = ((("" + "a") + "\n\t") + "")
                                                                                                }
                                                                                            case 127:
                                                                                                fmt.Println(v35, v133, v53)
                                                                                            case 0:
                                                                                                fmt.Println(v63, v35)
                                                                                            }
                                                                                            _ = v133
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v134 float64 = 1.5
                                                                                        _ = v134
                                                                                    }
                                                                                }
                                                                            }
                                                                            v54 = (42 >= 57)
                                                                            v4 -= (-(fi27))
                                                                        }
                                                                        var v135 int32 = 2147483647
                                                                        _ = v135
                                                                        _ = v65
                                                                    }
                                                                    var v136 int16 = 1
                                                                    v137 := GenS11{F0: v53, F1: 1}
                                                                    _ = v136
                                                                    _ = v137
                                                                    _ = v63
                                                                }
                                                            default:
                                                                var v138 int = -1
                                                                _ = v138
                                                            }
                                                            v139 := GenS12{F0: v53, F1: 675, F2: v54}
                                                            _ = v139
                                                        }
                                                        switch -9223372036854775808 {
                                                        case -826:
                                                            if v2 {
                                                                if (42 <= 42) {
                                                                    v2 = true
                                                                    v6["\n\t"] = (((v35 >= uint16(v18)) || (42 > v18)) && true)
                                                                }
                                                            }
                                                        }
                                                        _ = v54
                                                    }
                                                default:
                                                    var v140 bool = ("hello" > "foo bar")
                                                    _ = v140
                                                }
                                                _ = v53
                                            }
                                        } else {
                                            var v141 float64 = v5
                                            if (!v2) {
                                                var v142 float32 = 3.14159
                                                v5 = -2.25
                                                _ = v142
                                            }
                                            v141 = float64(float32(v5))
                                            _ = v141
                                        }
                                    default:
                                        fmt.Println(v5, v35, v18)
                                    }
                                }
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        case 165:
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        default:
                            if v2 {
                                switch 127 {
                                case 32767:
                                    fmt.Println(v3, v5, fi27)
                                default:
                                    var v143 int32 = 0
                                    _ = v143
                                }
                                var v144 int8 = 0
                                _ = v144
                            }
                        }
                        v18 = v3
                        var v145 uint32 = 65535
                        _ = v145
                    }
                    _ = fi27
                }
            } else {
                v2 = v2
            }
            _ = v26
        }
    case "foo bar":
        v4 = v4
    case "":
        fmt.Println(v3, v5)
    }
    var v146 bool = ((-486 >= -1) || v2)
    _ = len(v1)
    if false {
        if v146 {
            for fi147 := 0; fi147 < 2; fi147++ {
                if (v2 || (65535 > 535)) {
                    for fi148 := 0; fi148 < 19; fi148++ {
                        _ = len(v1)
                        var v149 uint = 472
                        _ = fi148
                        _ = v149
                    }
                    v5 += v5
                    v4 -= 1
                } else {
                    var v150 uint32 = 554
                    if (!(!v2)) {
                        v4 = -699
                        v151 := GenS6{F0: "a", F1: "foo bar"}
                        _ = v151
                    }
                    v2 = v2
                    _ = v150
                }
                if (!true) {
                    v152 := GenS12{F0: 0, F1: -1, F2: (((v5 / v5) <= (float64(0.1) / ((v5 * 0.0) / 0.1))) || (false && (!(-9223372036854775808 != -1))))}
                    var v153 bool = true
                    _ = v152
                    _ = v153
                }
                _ = fi147
            }
            var v154 float64 = ((v5 / v5) + v5)
            _ = v154
        }
        _ = len(v1)
        var v155 uint64 = v3
        _ = v155
    } else {
        v18 -= (v18 | v3)
        var v156 int32 = -2147483648
        var v157 uint16 = 42
        _ = v156
        _ = v157
    }
    v4 = (v4 >> ((68) & 63))
    v158 := GenS13{F0: "\n\t"}
    var v159 uint32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v146:", v146)
    fmt.Println("v159:", v159)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v146
    _ = v158
    _ = v159
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
