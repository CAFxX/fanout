package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int8; F1 int; F2 int }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int16; F1 int8; F2 uint32 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 int32; F1 int8; F2 int8 }
type GenS11 struct { F0 int16; F1 uint8; F2 int64 }
type GenS12 struct { F0 int; F1 int8; F2 int16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 string; F1 int16 }
type GenS15 struct { F0 int64; F1 int32 }
type GenS16 struct { F0 int16; F1 int64 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS19 struct { F0 int64 }
func genFn0(p0 uint, p1 int8, p2 uint) uint64 {
    p1 = (p1 / ((p1) | 1))
    fmt.Println(p0, p1)
    p0 = (p0 + (p2 | (p0 << (((p0 >> (((p2 & p2)) & 63))) & 63))))
    v4 := []int{-446, -526, 579}
    var v5 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 621
}
func genFn1(p0 uint16) string {
    var v6 int64 = 550
    fmt.Println(v6)
    p0 = p0
    if false {
        v6 = v6
        switch 508 {
        case -1:
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: 0, F1: -1, F2: 508}
                _ = fi7
                _ = v8
            }
        case 32768:
            v9 := []int{4294967295}
            _ = v9
        default:
            for fi10 := 0; fi10 < 20; fi10++ {
                fmt.Println(fi10, p0, fi10)
                fmt.Println(v6)
                switch ("\n\t" + "a") {
                case "hello":
                    if ((4294967295 >= 4294967295) && (65535 != 255)) {
                        var v11 int32 = 0
                        _ = v11
                    } else {
                        fmt.Println(p0, v6)
                        v6 = 0
                        for fi12 := 0; fi12 < 3; fi12++ {
                            p0 -= p0
                            if true {
                                if false {
                                    switch (fi10 << ((0) & 63)) {
                                    case 4294967295:
                                        if false {
                                            v13 := []int{(fi10 / ((fi10) | 1)), fi12, fi10, fi10}
                                            var v14 float64 = -2.25
                                            _ = v13
                                            _ = v14
                                        } else {
                                            var v15 int64 = -528
                                            v6 += v6
                                            switch ("hello" + "foo bar") {
                                            case "":
                                                if (false && (true && true)) {
                                                    var v16 float64 = 1e10
                                                    _ = v16
                                                }
                                            case "a":
                                                v17 := GenS2{F0: (524 >= (v6 * (v15 & v15))), F1: (fi12 / ((fi12) | 1))}
                                                _ = v17
                                            }
                                            _ = v15
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(fi12, p0, v6)
                                    }
                                } else {
                                    if false {
                                        v18 := []int{(^(fi12)), fi10, fi10, 842}
                                        switch (fi10 % (((fi12 - (^(fi10)))) | 1)) {
                                        case 9223372036854775807:
                                            v6 = v6
                                        case -9223372036854775808:
                                            switch (((("hello" + ("" + "a")) + (((("\n\t" + ("a" + "a")) + "a") + "foo bar") + "hello")) + (("hello" + "hello") + "")) + ("foo bar" + "a")) {
                                            case "hello":
                                                var v19 uint16 = p0
                                                _ = v19
                                            case "a":
                                                v20 := GenS2{F0: false, F1: (fi10 & 512)}
                                                _ = v20
                                            }
                                        case -1:
                                            v6 = (v6 & v6)
                                        }
                                        _ = v18
                                    } else {
                                        v21 := []int{fi12, fi10, 710}
                                        if true {
                                            switch 0 {
                                            case 32767:
                                                if len(v21) > 0 {
                                                    v21[0] = (fi10 % (((fi10 % ((fi10) | 1))) | 1))
                                                }
                                            default:
                                                var v22 uint16 = p0
                                                _ = v22
                                            }
                                            switch (fi10 >> ((255) & 63)) {
                                            case 880:
                                                var v23 float64 = 1e10
                                                _ = v23
                                            case -440:
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                            }
                                        }
                                        var v24 bool = (!false)
                                        _ = v21
                                        _ = v24
                                    }
                                }
                                v25 := GenS3{F0: -128, F1: fi10, F2: (-(fi10))}
                                p0 = (^(p0))
                                _ = v25
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    v6 = v6
                                case "":
                                    v26 := GenS4{F0: fi12, F1: 192}
                                    _ = v26
                                case "foo bar":
                                    var v27 int8 = -123
                                    var v28 uint = 4294967295
                                    v27 = v27
                                    _ = v27
                                    _ = v28
                                default:
                                    p0 -= 547
                                }
                                if (!(!(!true))) {
                                    if false {
                                        var v29 int32 = 2147483647
                                        var v30 uint8 = 1
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                fmt.Println(v6)
                            }
                            var v31 bool = (true || true)
                            var v32 int8 = 100
                            fmt.Println(fi10)
                            v32 += v32
                            v33 := GenS1{F0: 425, F1: -180, F2: 42}
                            _ = fi12
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                }
                _ = fi10
            }
        }
    } else {
        p0 += (^(p0))
        v34 := []int{0, 9223372036854775807, 42, -9223372036854775808}
        _ = v34
    }
    _ = p0
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    var v3 float64 = (1.5 + 1e10)
    v3 = v3
    v2 = v2
    var v35 int8 = 127
    var v36 uint32 = 90
    v3 = v3
    v37 := []int{-2147483648}
    v38 := GenS5{F0: (v2 ^ (^(v2))), F1: 255}
    if false {
        switch "foo bar" {
        case "hello":
            switch "\n\t" {
            case "hello":
                v39 := v1[("foo bar" + ("foo bar" + ""))]
                _ = v39
            }
        case "":
            for fi40 := 0; fi40 < 11; fi40++ {
                if (((((v3 / 0.1) / (3.14159 + 1e10)) * v3) - 3.14159) != ((v3 + v3) + 3.14159)) {
                    fmt.Println(v3, v2, v35)
                    fmt.Println(v3, fi40)
                    fmt.Println(v3, v3, v36)
                }
                if (!(!(v3 < v3))) {
                    fmt.Println(v35)
                    switch ((("foo bar" + "a") + ("a" + "hello")) + "a") {
                    case "hello":
                        v41 := GenS2{F0: false, F1: 4294967295}
                        _ = v41
                    case "\n\t":
                        switch (^(fi40)) {
                        case -134:
                            v3 += v3
                        case 519:
                            if (!true) {
                                if len(v37) >= 2 {
                                    v37 = v37[1:len(v37)]
                                }
                                v2 = (^(v2))
                            }
                        case 128:
                            var v42 uint32 = (v36 / ((4294967295) | 1))
                            _ = v42
                        default:
                            switch ((("foo bar" + "hello") + "foo bar") + "foo bar") {
                            case "hello":
                                var v43 bool = (((false || true) || (((v36 + v36) < 498) && ((v2 >> ((255) & 63)) == (v2 + (v2 + v2))))) || false)
                                _ = v43
                            }
                        }
                    case "foo bar":
                        v44 := GenS4{F0: fi40, F1: 0}
                        _ = v44
                    default:
                        v35 += (v35 ^ (-(v35)))
                    }
                    fmt.Println(v36, fi40)
                } else {
                    if ((float32((v3 + v3)) <= -2.25) || (int(v35) < fi40)) {
                        for fi45 := 0; fi45 < 13; fi45++ {
                            v2 = v2
                            _ = fi45
                        }
                        _ = len(v37)
                    }
                    if len(v37) > 0 {
                        v37[0] = fi40
                    }
                    for fi46 := 0; fi46 < 10; fi46++ {
                        v47 := GenS5{F0: 0, F1: 544}
                        _ = fi46
                        _ = v47
                    }
                }
                v35 = (^(v35))
                _ = fi40
            }
        default:
            if (!("foo bar" >= "foo bar")) {
                _ = len(v37)
                for fi48 := 0; fi48 < 10; fi48++ {
                    _ = len(v1)
                    v49 := GenS5{F0: -1, F1: 218}
                    _ = fi48
                    _ = v49
                }
            }
        }
    }
    switch ("a" + "a") {
    case "a":
        _ = len(v1)
    case "\n\t":
        var v50 bool = true
        _ = v50
    }
    fmt.Println(v35, v35)
    v2 += 0
    var v51 float32 = float32((((v3 / (v3 / v3)) / 1.0) - v3))
    for fi52 := 0; fi52 < 9; fi52++ {
        if len(v37) > 0 {
            v37[0] = (fi52 % (((-(fi52))) | 1))
        }
        var v53 int64 = -1
        for fi54 := 0; fi54 < 4; fi54++ {
            var v55 uint8 = 42
            if (fi52 != fi54) {
                v56 := GenS5{F0: -1, F1: 4294967295}
                v2 += 0
                var v57 uint32 = v36
                _ = v56
                _ = v57
            }
            _ = fi54
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    v3 += float64(((v51 + (0.1 / v51)) + v51))
    v58 := GenS6{F0: false, F1: "\n\t"}
    v3 -= 1e10
    v3 = float64(v51)
    for fi59 := 0; fi59 < 2; fi59++ {
        for fi60 := 0; fi60 < 11; fi60++ {
            var v61 int64 = -706
            if false {
                fmt.Println(v35, fi59, v35)
                var v62 int64 = 42
                if true {
                    v62 = v61
                    var v63 float32 = v51
                    _ = v63
                } else {
                    v61 = 35
                    fmt.Println(v61, v2, fi59)
                    if (50 < 32767) {
                        if len(v37) > 0 {
                            v37[0] = (-(fi59))
                        }
                    } else {
                        switch 226 {
                        case -396:
                            var v64 uint64 = 849
                            _ = v64
                        case 0:
                            switch (^(fi59)) {
                            case -880:
                                if len(v37) > 0 {
                                    v37[0] = 32768
                                }
                            case 32767:
                                v65 := GenS5{F0: v2, F1: uint(v36)}
                                _ = v65
                            case -828:
                                var v66 uint32 = v36
                                _ = v66
                            }
                        case 498:
                            v67 := GenS7{F0: int16(v35), F1: v35, F2: (^(v36))}
                            _ = v67
                        default:
                            v68 := GenS7{F0: -32768, F1: v35, F2: 991}
                            _ = v68
                        }
                    }
                }
                _ = v62
            }
            v35 = -1
            switch fi59 {
            case 2147483647:
                if true {
                    var v69 string = (("hello" + "foo bar") + "\n\t")
                    var v70 int = (fi59 - -2147483648)
                    _ = v69
                    _ = v70
                }
            case 1:
                var v71 int64 = -9223372036854775808
                if true {
                    v72 := GenS3{F0: v35, F1: -1, F2: (-(fi59))}
                    var v73 bool = true
                    _ = v72
                    _ = v73
                } else {
                    var v74 uint16 = 0
                    _ = v74
                }
                v75 := GenS1{F0: 762, F1: -1, F2: v36}
                if true {
                    if (true && (!false)) {
                        if (true && (1 > genFn0(65535, (v35 - (v35 % (((v35 >> ((150) & 63))) | 1))), 4294967295))) {
                            fmt.Println(v61, fi60)
                            v35 -= v35
                        } else {
                            var v76 float32 = 0.1
                            _ = v76
                        }
                        if len(v37) >= 2 {
                            v37 = v37[1:len(v37)]
                        }
                    } else {
                        fmt.Println(fi59, fi60, v35)
                        v2 = v2
                    }
                    var v77 int64 = (v61 ^ v71)
                    _ = v77
                }
                _ = v71
                _ = v75
            case 4294967295:
                fmt.Println(v61)
            default:
                if true {
                    v36 -= v36
                } else {
                    switch ("hello" + "\n\t") {
                    case "":
                        if (v36 > v36) {
                            if true {
                                if true {
                                    v78 := GenS8{F0: 1, F1: (fi59 + fi60)}
                                    _ = v78
                                } else {
                                    var v79 string = "\n\t"
                                    _ = v79
                                }
                                if false {
                                    var v80 int16 = -598
                                    if len(v37) >= 2 {
                                        v37 = v37[1:len(v37)]
                                    }
                                    _ = v80
                                }
                                switch (fi60 | fi59) {
                                case 0:
                                    v81 := GenS9{F0: 415, F1: v36}
                                    _ = v81
                                case 4294967295:
                                    var v82 float64 = 0.1
                                    _ = v82
                                case 127:
                                    if (true && true) {
                                        var v83 int32 = v2
                                        if false {
                                            fmt.Println(v3, v51)
                                            if false {
                                                v35 += (v35 * 42)
                                                if (v35 < (-(v35))) {
                                                    v51 += v51
                                                } else {
                                                    if true {
                                                        var v84 bool = (fi60 == fi60)
                                                        _ = v84
                                                    } else {
                                                        if ((fi60 | int(v2)) <= fi60) {
                                                            var v85 int8 = 1
                                                            _ = v85
                                                        } else {
                                                            _ = len(v37)
                                                        }
                                                    }
                                                    v86 := GenS8{F0: 455, F1: (fi59 - (fi60 % (((^(fi60))) | 1)))}
                                                    var v87 bool = true
                                                    _ = v86
                                                    _ = v87
                                                }
                                                if len(v37) >= 2 {
                                                    v37 = v37[1:len(v37)]
                                                }
                                                if true {
                                                    switch (fi60 ^ 4294967295) {
                                                    case 32767:
                                                        var v88 uint = 627
                                                        _ = v88
                                                    default:
                                                        var v89 uint8 = 125
                                                        _ = v89
                                                    }
                                                } else {
                                                    var v90 int16 = 42
                                                    _ = len(v37)
                                                    _ = v90
                                                }
                                                if (!false) {
                                                    v36 += 826
                                                    if ((("" + ("a" + ("foo bar" + ((("hello" + "a") + "") + "hello")))) + "\n\t") != "foo bar") {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if true {
                                                                v51 = -2.25
                                                                if (75 == 959) {
                                                                    var v91 int64 = (v61 - 42)
                                                                    if false {
                                                                        if (v61 == 0) {
                                                                            var v92 uint16 = 945
                                                                            if len(v37) >= 2 {
                                                                                v37 = v37[1:len(v37)]
                                                                            }
                                                                            var v93 uint16 = v92
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                    _ = v91
                                                                }
                                                                v2 = -2147483648
                                                            }
                                                        }
                                                    } else {
                                                        switch ((("a" + "hello") + (((("\n\t" + (("\n\t" + ("" + "foo bar")) + "a")) + (("foo bar" + ("a" + ("\n\t" + ("a" + ("a" + "foo bar"))))) + ("foo bar" + ("hello" + (("\n\t" + "foo bar") + "foo bar"))))) + ("" + "foo bar")) + "foo bar")) + ((("foo bar" + ("\n\t" + ("" + "foo bar"))) + ("hello" + "a")) + "hello")) {
                                                        case "\n\t":
                                                            if ("" == "hello") {
                                                                v94 := GenS10{F0: v83, F1: -128, F2: v35}
                                                                switch "a" {
                                                                case "a":
                                                                    v51 += v51
                                                                default:
                                                                    switch (fi60 ^ 46) {
                                                                    case 77:
                                                                        v2 = (v2 - (v2 % ((-45) | 1)))
                                                                    case -18:
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            v95 := GenS10{F0: (v2 << ((0) & 63)), F1: -1, F2: 100}
                                                                            _ = v95
                                                                        default:
                                                                            v96 := GenS5{F0: v2, F1: 4294967295}
                                                                            _ = v96
                                                                        }
                                                                    default:
                                                                        fmt.Println(v51, v2, v51)
                                                                    }
                                                                }
                                                                _ = v94
                                                            } else {
                                                                fmt.Println(v35)
                                                            }
                                                        case "":
                                                            fmt.Println(v51)
                                                        case "foo bar":
                                                            delete(v1, ("foo bar" + "foo bar"))
                                                            v51 -= (-2.25 / v51)
                                                        }
                                                        if (true && (((0 < 110) && (!(false || false))) && (!(true && (!true))))) {
                                                            switch "" {
                                                            case "":
                                                                v97 := GenS11{F0: -552, F1: 128, F2: (v61 % (((-(v61))) | 1))}
                                                                _ = v97
                                                            case "hello":
                                                                v51 -= v51
                                                            default:
                                                                switch (fi60 ^ fi60) {
                                                                case 1:
                                                                    var v98 float64 = v3
                                                                    _ = v98
                                                                case 9223372036854775807:
                                                                    var v99 bool = true
                                                                    _ = v99
                                                                case 2147483647:
                                                                    v83 = (v83 - v83)
                                                                default:
                                                                    v35 += v35
                                                                }
                                                            }
                                                            var v100 uint16 = 979
                                                            _ = v100
                                                        } else {
                                                            v101 := GenS12{F0: (-(fi59)), F1: v35, F2: -546}
                                                            _ = len(v1)
                                                            fmt.Println(v3)
                                                            _ = v101
                                                        }
                                                        fmt.Println(v35, v3, v51)
                                                        v102 := GenS13{F0: 2}
                                                        var v103 int16 = -32768
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                } else {
                                                    v1["\n\t"] = 42
                                                }
                                                var v104 int8 = (v35 / (((-(v35))) | 1))
                                                v61 -= (v61 | 0)
                                                if false {
                                                    v105 := GenS14{F0: "hello", F1: -67}
                                                    if false {
                                                        v3 = v3
                                                        v37 = append(v37, (fi60 * int(v83)))
                                                        v51 -= v51
                                                    } else {
                                                        switch fi59 {
                                                        case -742:
                                                            v106 := GenS1{F0: 67, F1: -721, F2: (v36 | v36)}
                                                            _ = v106
                                                        default:
                                                            var v107 uint64 = 4294967295
                                                            _ = v107
                                                        }
                                                        v104 += (v104 % (((v35 + (v104 ^ (v35 | (v104 << ((599) & 63)))))) | 1))
                                                    }
                                                    _ = v105
                                                } else {
                                                    var v108 uint = 255
                                                    _ = v108
                                                }
                                                _ = v104
                                            }
                                        } else {
                                            fmt.Println(v36)
                                            switch fi59 {
                                            case -496:
                                                fmt.Println(fi60, v51, v83)
                                            default:
                                                if (!((!true) || ((v51 * 0.1) != v51))) {
                                                    if false {
                                                        if (!(1 >= v36)) {
                                                            switch ("a" + ("foo bar" + "a")) {
                                                            case "a":
                                                                var v109 float32 = float32(float64(float32(1e10)))
                                                                _ = v109
                                                            case "foo bar":
                                                                if (v35 <= (v35 << ((140) & 63))) {
                                                                    fmt.Println(v51, v83, v2)
                                                                    v110 := GenS9{F0: 0, F1: v36}
                                                                    if false {
                                                                        fmt.Println(v61, v83)
                                                                        var v111 int8 = 0
                                                                        if (0 < v2) {
                                                                            v112 := GenS8{F0: -32768, F1: 0}
                                                                            if true {
                                                                                if len(v37) > 0 {
                                                                                    v37[0] = 380
                                                                                }
                                                                                if (1e10 < (1.5 * 1e10)) {
                                                                                    v51 = v51
                                                                                    v111 = (v35 | v35)
                                                                                    v113 := GenS15{F0: (^(v61)), F1: v83}
                                                                                    _ = v113
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(v51, fi59, v83)
                                                                                    v36 = v36
                                                                                    if false {
                                                                                        switch ("a" + ("\n\t" + ("hello" + "a"))) {
                                                                                        case "foo bar":
                                                                                            v114 := GenS1{F0: 595, F1: -1, F2: v36}
                                                                                            _ = v114
                                                                                        case "\n\t":
                                                                                            if (fi59 > fi60) {
                                                                                                var v115 int32 = 445
                                                                                                _ = v115
                                                                                            } else {
                                                                                                v51 += (v51 * v51)
                                                                                                v116 := v1[""]
                                                                                                if true {
                                                                                                    fmt.Println(v35, fi60)
                                                                                                }
                                                                                                _ = v116
                                                                                            }
                                                                                        default:
                                                                                            switch int(v36) {
                                                                                            case -9223372036854775808:
                                                                                                v1[(("hello" + ("a" + "a")) + ("hello" + "foo bar"))] = v36
                                                                                            case 9223372036854775807:
                                                                                                var v117 float64 = 1e10
                                                                                                _ = v117
                                                                                            default:
                                                                                                var v118 uint = 4294967295
                                                                                                _ = v118
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v119 bool = (!(0 != 458))
                                                                                    var v120 string = "foo bar"
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                }
                                                                                fmt.Println(fi59, fi59)
                                                                                var v121 int64 = 899
                                                                                _ = v121
                                                                            }
                                                                            _ = len(v1)
                                                                            _ = v112
                                                                        } else {
                                                                            var v122 uint64 = 2
                                                                            var v123 uint64 = 2
                                                                            _ = v122
                                                                            _ = v123
                                                                        }
                                                                        _ = v111
                                                                    } else {
                                                                        v36 += 65535
                                                                        if len(v37) > 0 {
                                                                            v37[0] = int(v61)
                                                                        }
                                                                    }
                                                                    _ = v110
                                                                } else {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v124 := GenS5{F0: (v2 ^ (v83 ^ (v83 / ((v2) | 1)))), F1: 2}
                                                                        _ = v124
                                                                    case "":
                                                                        fmt.Println(v83, fi60, v51)
                                                                    }
                                                                    v125 := v1[(("\n\t" + "a") + ("foo bar" + ""))]
                                                                    _ = v125
                                                                }
                                                            case "\n\t":
                                                                v126 := GenS10{F0: (v2 >> ((538) & 63)), F1: v35, F2: v35}
                                                                _ = v126
                                                            default:
                                                                var v127 uint16 = 255
                                                                _ = v127
                                                            }
                                                        }
                                                        if (1 > 0) {
                                                            if len(v37) >= 2 {
                                                                v37 = v37[1:len(v37)]
                                                            }
                                                            var v128 float32 = float32(v3)
                                                            _ = v128
                                                        }
                                                        var v129 int16 = -431
                                                        v130 := v1["hello"]
                                                        v131 := GenS16{F0: (-(v129)), F1: (v61 >> ((166) & 63))}
                                                        v129 = v129
                                                        _ = v129
                                                        _ = v130
                                                        _ = v131
                                                    }
                                                    var v132 string = "a"
                                                    var v133 string = v132
                                                    fmt.Println(v36)
                                                    _ = v132
                                                    _ = v133
                                                } else {
                                                    v3 = (1e10 / (((v3 / ((1.5 * 0.0) - v3)) + v3) * 1e10))
                                                    v1[("foo bar" + "foo bar")] = (v36 - v36)
                                                    _ = len(v37)
                                                    v35 -= -75
                                                }
                                            }
                                        }
                                        if true {
                                            if (!true) {
                                                switch fi59 {
                                                case 9223372036854775807:
                                                    v134 := v1["hello"]
                                                    _ = v134
                                                case 718:
                                                    var v135 uint64 = 18446744073709551615
                                                    _ = v135
                                                }
                                            } else {
                                                var v136 uint32 = 675
                                                _ = v136
                                            }
                                            var v137 string = ((((("\n\t" + "\n\t") + "a") + ((("" + (((("\n\t" + "a") + "a") + "foo bar") + "")) + "hello") + "a")) + "\n\t") + "a")
                                            v138 := GenS17{F0: int64(v36)}
                                            _ = v137
                                            _ = v138
                                        }
                                        _ = v83
                                    } else {
                                        v139 := GenS7{F0: -32768, F1: -1, F2: v36}
                                        var v140 uint16 = 32767
                                        _ = v139
                                        _ = v140
                                    }
                                }
                                switch (fi59 % (((fi59 ^ 42)) | 1)) {
                                case 1:
                                    var v141 uint32 = v36
                                    _ = v141
                                case -292:
                                    v142 := GenS18{F0: (v36 | 1), F1: 255, F2: 662}
                                    _ = v142
                                case 42:
                                    v1["foo bar"] = v36
                                }
                            } else {
                                v3 = 1.5
                                v143 := GenS19{F0: v61}
                                _ = v143
                            }
                        }
                    case "\n\t":
                        var v144 uint16 = 65535
                        _ = v144
                    case "foo bar":
                        fmt.Println(v51, v36)
                    default:
                        var v145 int8 = int8(v2)
                        _ = v145
                    }
                }
            }
            v3 = v3
            var v146 string = "hello"
            _ = fi60
            _ = v146
            _ = v61
        }
        var v147 uint64 = 445
        v36 = v36
        _ = fi59
        _ = v147
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v51
    _ = v58
}
