package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int; F1 int64; F2 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 uint; F2 int8 }
type GenS8 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 string, p1 uint) uint32 {
    var v7 int32 = 460
    for fi8 := 0; fi8 < 1; fi8++ {
        p1 += (p1 | (p1 & (p1 * (p1 << ((p1) & 63)))))
        fmt.Println(p1, p1, p1)
        var v9 int32 = -2147483648
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 2147483648
}

func main() {
    var v1 uint16 = 703
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v4 int64 = -9223372036854775808
    var v5 int8 = 127
    var v6 bool = false
    fmt.Println(v1, v1, v6)
    _ = len(v2)
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: 255}
        var v12 uint16 = v1
        v5 = v5
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint16 = (v1 / (((v1 >> ((1) & 63))) | 1))
    v6 = v6
    v13 += 299
    v4 = (v4 & -1)
    switch "" {
    case "":
        switch ((("" + "hello") + "") + (("a" + "") + "")) {
        case "hello":
            for fi14 := 0; fi14 < 6; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if true {
                        v4 = 42
                        var v16 int = fi14
                        v16 -= int(v1)
                        v4 -= v4
                        _ = v16
                    }
                    _ = fi15
                }
                var v17 bool = true
                v1 = (v13 - v13)
                _ = fi14
                _ = v17
            }
        case "":
            var v18 uint8 = 255
            _ = v18
        case "\n\t":
            for fi19 := 0; fi19 < 9; fi19++ {
                v3[int(v5)] = uint32(v1)
                v13 -= 255
                var v20 uint16 = (v1 / (((^(v13))) | 1))
                _ = fi19
                _ = v20
            }
        }
    case "a":
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS1{F0: 42}
            if (float32(0.0) < 0.0) {
                _ = len(v2)
                fmt.Println(v13, v1, v13)
            }
            _ = fi21
            _ = v22
        }
    case "\n\t":
        fmt.Println(v13, v1)
    default:
        var v23 float64 = (-2.25 / 0.1)
        _ = v23
    }
    v24 := GenS1{F0: 191}
    fmt.Println(v1, v1, v5)
    v25 := GenS1{F0: 78}
    if true {
        if true {
            for fi26 := 0; fi26 < 8; fi26++ {
                v2[genFn0("\n\t", 958)] = fi26
                v27 := GenS1{F0: 0}
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                if (!v6) {
                    switch fi28 {
                    case 42:
                        var v29 uint32 = genFn0("\n\t", 964)
                        _ = v29
                    case 334:
                        var v30 bool = v6
                        _ = v30
                    }
                    var v31 uint = 4294967295
                    _ = v31
                } else {
                    v4 = (v4 ^ 9223372036854775807)
                }
                var v32 uint8 = 0
                _ = fi28
                _ = v32
            }
            _ = len(v2)
        }
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        switch "\n\t" {
        case "a":
            v5 = 38
        }
        switch "\n\t" {
        case "foo bar":
            if false {
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 uint32 = 763
                    switch (-(fi34)) {
                    case 0:
                        if (-1 >= -2147483648) {
                            var v36 float64 = 1.5
                            v36 = v36
                            _ = v36
                        }
                        if ((^(fi34)) == (fi34 % (((fi34 >> ((0) & 63))) | 1))) {
                            var v37 uint16 = 151
                            v38 := GenS1{F0: 50}
                            _ = v37
                            _ = v38
                        } else {
                            v5 -= 19
                            if (!(320 >= 903)) {
                                v13 = (^(v1))
                                if v6 {
                                    if (v6 && v6) {
                                        var v39 int64 = -1
                                        fmt.Println(v1, fi34)
                                        _ = v39
                                    }
                                } else {
                                    v6 = ((0.1 <= -2.25) || v6)
                                    var v40 uint32 = v35
                                    fmt.Println(v35, fi34, v1)
                                    _ = v40
                                }
                                if ((^(v5)) > -96) {
                                    if ((!(v6 && (((((("" + ("a" + "hello")) + ("\n\t" + "\n\t")) + ("" + ("\n\t" + ("hello" + "foo bar")))) + "") + "foo bar") >= "foo bar"))) || ((v5 | v5) < (^(v5)))) {
                                        v41 := GenS1{F0: 128}
                                        if v6 {
                                            if v6 {
                                                v42 := GenS2{F0: v35}
                                                v35 = 42
                                                fmt.Println(fi33, v4, fi34)
                                                _ = v42
                                            }
                                            v13 = (v1 + v13)
                                            var v43 string = "\n\t"
                                            _ = v43
                                        }
                                        _ = v41
                                    } else {
                                        if false {
                                            delete(v2, v35)
                                            fmt.Println(v1, v5)
                                        }
                                        var v44 string = "a"
                                        v35 -= (v35 * (v35 << ((309) & 63)))
                                        switch (fi33 & fi34) {
                                        case 817:
                                            var v45 int = int(v35)
                                            _ = v45
                                        case 2147483647:
                                            v6 = v6
                                        }
                                        _ = v44
                                    }
                                    v6 = v6
                                    switch "a" {
                                    case "foo bar":
                                        if (2 == 18446744073709551615) {
                                            v13 += (^(v13))
                                            var v46 float64 = (1e10 * (float64(-2.25) - 3.14159))
                                            _ = v46
                                        }
                                    case "hello":
                                        var v47 int64 = int64(v1)
                                        _ = v47
                                    default:
                                        v35 = (v35 + 63)
                                    }
                                    if false {
                                        v48 := []int{}
                                        v5 = (v5 << ((255) & 63))
                                        _ = len(v48)
                                        v13 = (v13 / (((v1 / ((v1) | 1))) | 1))
                                        v48 = append(v48, fi34)
                                        _ = v48
                                    } else {
                                        v6 = ((127 <= 232) && v6)
                                    }
                                    switch "hello" {
                                    case "":
                                        var v49 int32 = -2147483648
                                        _ = v49
                                    case "foo bar":
                                        var v50 uint32 = (v35 >> ((196) & 63))
                                        _ = v50
                                    }
                                    if false {
                                        var v51 int16 = 1000
                                        _ = v51
                                    }
                                    switch -118 {
                                    case -9223372036854775808:
                                        v52 := v2[2]
                                        _ = v52
                                    case 128:
                                        v4 = 0
                                    }
                                } else {
                                    v6 = (96 == 2)
                                    switch int(v35) {
                                    case -997:
                                        switch (("" + "") + "a") {
                                        case "":
                                            _ = len(v3)
                                        case "\n\t":
                                            switch (("a" + ("" + "foo bar")) + "") {
                                            case "hello":
                                                fmt.Println(v4)
                                            case "a":
                                                if true {
                                                    if true {
                                                        fmt.Println(v13, v13)
                                                        if (2 < 2) {
                                                            v5 = v5
                                                            v53 := v3[int(v1)]
                                                            _ = v53
                                                        } else {
                                                            switch 966 {
                                                            case -160:
                                                                if v6 {
                                                                    var v54 bool = (1 < 1)
                                                                    _ = v54
                                                                } else {
                                                                    delete(v2, (v35 * (v35 - v35)))
                                                                    switch fi34 {
                                                                    case -183:
                                                                        if (!(!v6)) {
                                                                            var v55 float64 = float64(((float32(((0.0 / 3.14159) - (0.1 / -2.25))) - 1e10) + 1.5))
                                                                            v56 := v3[(fi34 / (((fi34 | (fi34 ^ fi33))) | 1))]
                                                                            var v57 int = fi33
                                                                            _ = v55
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                    case -2147483648:
                                                                        var v58 uint16 = 255
                                                                        _ = v58
                                                                    default:
                                                                        v1 = v13
                                                                    }
                                                                    var v59 int32 = 149
                                                                    if (0 != -32768) {
                                                                        var v60 string = ((("foo bar" + "\n\t") + ("a" + ("\n\t" + "hello"))) + (("a" + "a") + "foo bar"))
                                                                        if (!(v6 && (501 <= -480))) {
                                                                            v61 := v2[(v35 / (((v35 | 2147483647)) | 1))]
                                                                            _ = v61
                                                                        } else {
                                                                            var v62 int16 = 1
                                                                            v2[genFn0("\n\t", 4294967295)] = fi34
                                                                            _ = v62
                                                                        }
                                                                        _ = v60
                                                                    } else {
                                                                        v13 = 42
                                                                    }
                                                                    var v63 uint8 = 128
                                                                    _ = v59
                                                                    _ = v63
                                                                }
                                                            case 128:
                                                                var v64 int = (^(fi34))
                                                                _ = v64
                                                            }
                                                        }
                                                        fmt.Println(fi33)
                                                    } else {
                                                        var v65 uint8 = uint8(v1)
                                                        fmt.Println(v13, fi34, v4)
                                                        var v66 uint8 = 2
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    _ = len(v3)
                                                } else {
                                                    v1 -= v13
                                                }
                                            case "":
                                                switch 460 {
                                                case 2147483647:
                                                    var v67 int = (^(fi34))
                                                    _ = v67
                                                }
                                            default:
                                                switch ("foo bar" + "") {
                                                case "hello":
                                                    var v68 float32 = (((0.0 / 1.0) / 1.0) / (((0.1 + (((3.14159 / (3.14159 * 3.14159)) + -2.25) - -2.25)) + (1e10 - (0.0 / 1.0))) - float32(float64((3.14159 * float32(float64(1.5)))))))
                                                    _ = v68
                                                }
                                            }
                                        default:
                                            v69 := GenS3{F0: "\n\t", F1: 23}
                                            _ = v69
                                        }
                                    case 1:
                                        fmt.Println(fi33, v35, v35)
                                    case -9223372036854775808:
                                        if false {
                                            v70 := GenS3{F0: "hello", F1: v35}
                                            var v71 uint64 = 1
                                            _ = v70
                                            _ = v71
                                        } else {
                                            switch fi33 {
                                            case 128:
                                                v72 := GenS4{F0: 584, F1: v4, F2: v5}
                                                _ = v72
                                            default:
                                                v13 -= (v1 / (((v1 | v13)) | 1))
                                            }
                                            switch (fi34 + (fi34 + fi33)) {
                                            case 42:
                                                var v73 float32 = 3.14159
                                                _ = v73
                                            }
                                        }
                                        v74 := GenS5{F0: v6, F1: 42}
                                        _ = v74
                                    default:
                                        var v75 uint16 = 0
                                        _ = v75
                                    }
                                    fmt.Println(v35, v6, v35)
                                }
                            }
                        }
                    }
                    v6 = v6
                    v6 = (!((0.0 >= 1.5) && true))
                    _ = fi34
                    _ = v35
                }
                switch ("" + ("hello" + ("a" + ("hello" + ((("foo bar" + "hello") + "a") + "foo bar"))))) {
                case "foo bar":
                    v76 := GenS3{F0: ("" + ""), F1: 2}
                    _ = v76
                }
            }
        case "":
            if ((-2.25 + 1.5) < ((0.1 - 3.14159) * -2.25)) {
                var v77 int16 = -1
                _ = v77
            } else {
                var v78 int = (fi33 * fi33)
                _ = v78
            }
        case "\n\t":
            v5 = (^(v5))
        }
        if false {
            if (!v6) {
                for fi79 := 0; fi79 < 17; fi79++ {
                    fmt.Println(v13, v1, fi33)
                    v80 := GenS6{F0: 1}
                    v81 := GenS6{F0: 42}
                    _ = fi79
                    _ = v80
                    _ = v81
                }
            } else {
                var v82 uint64 = 288
                _ = v82
            }
        } else {
            v83 := GenS7{F0: "foo bar", F1: 42, F2: (v5 | v5)}
            v1 = v1
            for fi84 := 0; fi84 < 20; fi84++ {
                v85 := GenS8{F0: 2, F1: (2 > 2), F2: v1}
                _ = fi84
                _ = v85
            }
            _ = v83
        }
        _ = fi33
    }
    var v86 uint = 2
    switch -1 {
    case 2147483647:
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            v87 := GenS7{F0: "\n\t", F1: v86, F2: -14}
            _ = v87
        }
    default:
        if ((v4 ^ v4) == (v4 + v4)) {
            var v88 float32 = 3.14159
            v4 = 9223372036854775807
            _ = v88
        }
    }
    for fi89 := 0; fi89 < 4; fi89++ {
        v86 = v86
        for fi90 := 0; fi90 < 20; fi90++ {
            switch "a" {
            case "foo bar":
                var v91 bool = v6
                _ = v91
            case "a":
                fmt.Println(fi90)
            }
            var v92 uint64 = 1
            _ = fi90
            _ = v92
        }
        _ = fi89
    }
    for fi93 := 0; fi93 < 5; fi93++ {
        for fi94 := 0; fi94 < 5; fi94++ {
            if (2147483647 > 42) {
                if (v6 && true) {
                    v86 -= v86
                } else {
                    switch 864 {
                    case -225:
                        v4 = v4
                    case -984:
                        var v95 float64 = 0.0
                        _ = v95
                    case 127:
                        if true {
                            v96 := []int{int(v4), (^(fi94))}
                            _ = v96
                        }
                    }
                    var v97 uint8 = 1
                    _ = v97
                }
                v98 := v2[42]
                _ = v98
            }
            var v99 int32 = -2147483648
            v1 += v13
            fmt.Println(v1, v1)
            _ = fi94
            _ = v99
        }
        var v100 float32 = 3.14159
        _ = fi93
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v86
}
