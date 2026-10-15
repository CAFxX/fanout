package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 string; F2 int32 }
type GenS6 struct { F0 uint; F1 bool; F2 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 uint32; F1 int64; F2 int16 }
type GenS11 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS12 struct { F0 uint8; F1 int16 }
type GenS13 struct { F0 uint8; F1 string }
func genFn0(p0 int, p1 int8, p2 uint8) int8 {
    var v5 uint8 = p2
    v5 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 / ((p1) | 1))
}
func genFn1() int16 {
    var v6 int = -1
    v7 := GenS1{F0: v6, F1: 2147483647}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS2{F0: 0}
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS3{F0: -368}
            fmt.Println(fi10, v6, fi10)
            var v12 uint16 = 0
            fmt.Println(fi8, v12)
            var v13 int8 = 42
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        switch (^(v6)) {
        case -1:
            switch "a" {
            case "\n\t":
                var v14 uint16 = 0
                _ = v14
            case "a":
                if ((genFn0((fi8 ^ (-(fi8))), -128, 127) != -35) && true) {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 uint = 685
                        var v17 bool = (841 >= (fi8 / ((1) | 1)))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
            }
        case 9223372036854775807:
            if true {
                if true {
                    v18 := GenS4{F0: 583}
                    v19 := []int{-9223372036854775808, (v6 - 1), v6, (-(fi8))}
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS1{F0: fi8, F1: -2147483648}
                    _ = v20
                }
            } else {
                var v21 int16 = 3
                v21 = 32767
                switch -1 {
                case 2147483647:
                    for fi22 := 0; fi22 < 19; fi22++ {
                        v6 += (fi22 & (v6 & (v6 & (^(fi22)))))
                        fmt.Println(v21, fi22)
                        _ = fi22
                    }
                }
                _ = v21
            }
        }
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 string = ("foo bar" + "")
    var v2 string = "foo bar"
    var v3 uint = 1
    v4 := []int{42, 32767}
    var v23 int32 = -297
    if false {
        for fi24 := 0; fi24 < 11; fi24++ {
            switch "foo bar" {
            case "hello":
                v1 = (((v1 + (v2 + "a")) + "hello") + v1)
            }
            _ = fi24
        }
        var v25 uint8 = 1
        _ = v25
    }
    if (1 != 42) {
        var v26 uint16 = 32767
        if false {
            var v27 int64 = -963
            _ = len(v4)
            _ = v27
        }
        _ = v26
    } else {
        switch v1 {
        case "a":
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS3{F0: -569}
                if true {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        fmt.Println(v3, fi28)
                        _ = fi30
                    }
                    switch v2 {
                    case "hello":
                        var v31 float32 = 1e10
                        _ = v31
                    default:
                        v23 += 0
                    }
                }
                _ = fi28
                _ = v29
            }
        case "foo bar":
            if (3.14159 > 0.1) {
                if (false && false) {
                    if true {
                        if false {
                            v32 := GenS4{F0: 1}
                            v33 := GenS5{F0: -2147483648, F1: v2, F2: v23}
                            var v34 int16 = 849
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            var v35 int16 = 155
                            switch v1 {
                            case "":
                                v3 = v3
                            case "foo bar":
                                v4 = append(v4, 621)
                            case "\n\t":
                                if false {
                                    v23 = (v23 - v23)
                                    v1 = v1
                                } else {
                                    if (!(false && false)) {
                                        switch -9223372036854775808 {
                                        case 128:
                                            v36 := GenS1{F0: 127, F1: -1}
                                            _ = v36
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if false {
                                            v23 += v23
                                        } else {
                                            switch (("a" + v1) + ((((v1 + (v1 + "")) + v2) + (((("\n\t" + v2) + "\n\t") + "a") + (v2 + "\n\t"))) + v1)) {
                                            case "a":
                                                for fi37 := 0; fi37 < 9; fi37++ {
                                                    v38 := GenS6{F0: (v3 - 4294967295), F1: false, F2: 42}
                                                    _ = fi37
                                                    _ = v38
                                                }
                                            }
                                            for fi39 := 0; fi39 < 6; fi39++ {
                                                for fi40 := 0; fi40 < 4; fi40++ {
                                                    if true {
                                                        switch (v2 + "a") {
                                                        case "a":
                                                            _ = len(v4)
                                                        case "\n\t":
                                                            v2 = v1
                                                            var v41 uint = 300
                                                            v3 += v41
                                                            _ = v41
                                                        }
                                                        v1 = v1
                                                        var v42 int16 = genFn1()
                                                        _ = v42
                                                    } else {
                                                        if false {
                                                            var v43 uint = (^(v3))
                                                            v23 += v23
                                                            if (true && true) {
                                                                if (100 <= -128) {
                                                                    if false {
                                                                        switch ((v2 + "hello") + "foo bar") {
                                                                        case "foo bar":
                                                                            v23 = 42
                                                                        case "":
                                                                            var v44 int32 = (v23 >> (((^(v43))) & 63))
                                                                            _ = v44
                                                                        case "a":
                                                                            var v45 uint = (v43 | v43)
                                                                            _ = v45
                                                                        }
                                                                        if (1 <= -9223372036854775808) {
                                                                            switch v1 {
                                                                            case "":
                                                                                v1 = ""
                                                                            case "hello":
                                                                                v46 := GenS7{F0: 42}
                                                                                _ = v46
                                                                            default:
                                                                                if false {
                                                                                    v1 = "hello"
                                                                                    var v47 uint16 = 0
                                                                                    v4 = append(v4, fi39)
                                                                                    _ = v47
                                                                                }
                                                                                v48 := GenS5{F0: (fi40 - (^(fi40))), F1: v1, F2: (^(v23))}
                                                                                if (42 > 0) {
                                                                                    var v49 uint32 = 119
                                                                                    _ = v49
                                                                                } else {
                                                                                    switch -884 {
                                                                                    case 32768:
                                                                                        if (true && (128 > 0)) {
                                                                                            var v50 uint16 = 65535
                                                                                            _ = v50
                                                                                        }
                                                                                    case 587:
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                    default:
                                                                                        v51 := GenS8{F0: (^(v23)), F1: ("foo bar" + "hello")}
                                                                                        _ = v51
                                                                                    }
                                                                                    var v52 string = v1
                                                                                    _ = v52
                                                                                }
                                                                                _ = len(v4)
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                _ = v48
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v35 -= (v35 + v35)
                                                                    }
                                                                    switch (fi40 << (((^(v3))) & 63)) {
                                                                    case -9223372036854775808:
                                                                        if true {
                                                                            var v53 uint = 4294967295
                                                                            var v54 float32 = 1.5
                                                                            _ = v53
                                                                            _ = v54
                                                                        }
                                                                    case 42:
                                                                        v23 += (^(v23))
                                                                    default:
                                                                        v35 -= (-(v35))
                                                                    }
                                                                    v4 = append(v4, fi40)
                                                                }
                                                                v4 = append(v4, (fi40 + (fi40 & int(v35))))
                                                            }
                                                            _ = v43
                                                        }
                                                        v2 = v1
                                                    }
                                                    v23 = v23
                                                    _ = fi40
                                                }
                                                v35 = genFn1()
                                                v55 := GenS6{F0: (v3 % ((v3) | 1)), F1: ((v35 << ((381) & 63)) > v35), F2: 2}
                                                _ = fi39
                                                _ = v55
                                            }
                                        }
                                    }
                                    switch 9223372036854775807 {
                                    case -998:
                                        for fi56 := 0; fi56 < 5; fi56++ {
                                            var v57 int16 = int16(v3)
                                            if true {
                                                if true {
                                                    switch 977 {
                                                    case 42:
                                                        switch v2 {
                                                        case "\n\t":
                                                            var v58 string = ((v2 + ((v1 + (("\n\t" + "a") + v2)) + (v2 + ((v2 + "foo bar") + "hello")))) + v1)
                                                            _ = v58
                                                        case "hello":
                                                            for fi59 := 0; fi59 < 3; fi59++ {
                                                                var v60 string = ("foo bar" + ("" + "hello"))
                                                                _ = fi59
                                                                _ = v60
                                                            }
                                                        default:
                                                            var v61 int32 = v23
                                                            _ = v61
                                                        }
                                                    }
                                                    v62 := GenS6{F0: (v3 | (v3 - (v3 & 1))), F1: false, F2: 18446744073709551615}
                                                    v3 = v3
                                                    _ = v62
                                                }
                                                v63 := GenS3{F0: 42}
                                                _ = v63
                                            }
                                            _ = fi56
                                            _ = v57
                                        }
                                    }
                                    for fi64 := 0; fi64 < 17; fi64++ {
                                        fmt.Println(v35, v3)
                                        for fi65 := 0; fi65 < 19; fi65++ {
                                            var v66 float64 = 0.0
                                            if true {
                                                fmt.Println(v23, fi65)
                                                v1 = v2
                                                fmt.Println(v35, v23)
                                            } else {
                                                switch (fi64 - fi64) {
                                                case -1:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if ((v23 ^ (^(v23))) > -1) {
                                                            switch fi65 {
                                                            case 3:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            case -9223372036854775808:
                                                                switch (v1 + "a") {
                                                                case "":
                                                                    switch (fi65 - -2147483648) {
                                                                    case 128:
                                                                        _ = len(v4)
                                                                    }
                                                                default:
                                                                    fmt.Println(fi64, v2, v2)
                                                                }
                                                                fmt.Println(v2, v23, v66)
                                                            case 32768:
                                                                v67 := GenS8{F0: 269, F1: (("" + v2) + (v2 + (((v1 + "hello") + v1) + "foo bar")))}
                                                                _ = v67
                                                            default:
                                                                if true {
                                                                    if (1 <= 4294967295) {
                                                                        switch v2 {
                                                                        case "\n\t":
                                                                            var v68 int16 = v35
                                                                            _ = v68
                                                                        default:
                                                                            v3 = (v3 + (v3 - v3))
                                                                        }
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            v69 := GenS4{F0: -1}
                                                                            fmt.Println(fi64)
                                                                            _ = v69
                                                                        case "a":
                                                                            var v70 int64 = 101
                                                                            _ = v70
                                                                        case "hello":
                                                                            v1 = v1
                                                                        }
                                                                    }
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            }
                                                            v66 = v66
                                                        } else {
                                                            v4 = append(v4, 128)
                                                        }
                                                    default:
                                                        v4 = append(v4, (^(fi65)))
                                                    }
                                                case 4294967295:
                                                    fmt.Println(fi64, v1)
                                                case 2147483647:
                                                    if (v35 > v35) {
                                                        v1 = ""
                                                        v4 = append(v4, 1)
                                                        v71 := GenS9{F0: 42, F1: 0}
                                                        _ = v71
                                                    } else {
                                                        fmt.Println(fi64)
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v72 int32 = (v23 ^ v23)
                                                        _ = v72
                                                    }
                                                }
                                                if false {
                                                    v73 := GenS7{F0: v35}
                                                    _ = v73
                                                } else {
                                                    var v74 int16 = genFn1()
                                                    _ = v74
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            _ = fi65
                                            _ = v66
                                        }
                                        _ = fi64
                                    }
                                }
                            }
                            v35 = genFn1()
                            _ = v35
                        }
                    }
                }
                switch ("hello" + v2) {
                case "\n\t":
                    fmt.Println(v1, v2)
                case "":
                    v4 = append(v4, 676)
                }
            }
        }
    }
    for fi75 := 0; fi75 < 6; fi75++ {
        var v76 uint16 = 1
        switch fi75 {
        case 0:
            v23 = -2147483648
        case 502:
            v76 += 1
        default:
            v4 = append(v4, -9223372036854775808)
        }
        _ = fi75
        _ = v76
    }
    v2 = v2
    fmt.Println(v2)
    var v77 uint16 = 825
    if true {
        v4 = append(v4, 689)
    } else {
        v78 := GenS10{F0: 4294967295, F1: 42, F2: 32767}
        switch ("" + (v1 + v1)) {
        case "hello":
            v77 += v77
        case "a":
            v79 := GenS10{F0: 4294967295, F1: 0, F2: -1}
            _ = v79
        default:
            fmt.Println(v77)
        }
        v4 = append(v4, -540)
        _ = v78
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v80 uint64 = 696
    fmt.Println(v1, v77)
    for fi81 := 0; fi81 < 4; fi81++ {
        if (v1 > v2) {
            _ = len(v4)
        } else {
            switch v1 {
            case "foo bar":
                fmt.Println(v2, v1)
            case "":
                fmt.Println(v2, v2)
            default:
                var v82 float32 = float32((3.14159 / 1.0))
                _ = v82
            }
            var v83 int64 = -1
            _ = v83
        }
        if true {
            switch "hello" {
            case "a":
                v1 = v2
            }
            if true {
                v80 = v80
                if true {
                    v2 = v1
                    _ = len(v4)
                } else {
                    for fi84 := 0; fi84 < 17; fi84++ {
                        var v85 bool = ((false || false) || true)
                        _ = fi84
                        _ = v85
                    }
                    if ((fi81 <= 0) || true) {
                        var v86 float32 = (0.1 - -2.25)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi87 := 0; fi87 < 17; fi87++ {
                            v4 = append(v4, (fi87 | 32768))
                            v1 = ("hello" + ("" + v1))
                            if (true && false) {
                                var v88 int16 = 1000
                                _ = v88
                            } else {
                                v89 := GenS1{F0: -547, F1: (-(v23))}
                                v4 = append(v4, int(v3))
                                _ = v89
                            }
                            _ = fi87
                        }
                        _ = v86
                    } else {
                        for fi90 := 0; fi90 < 6; fi90++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v91 int8 = 0
                            _ = fi90
                            _ = v91
                        }
                        if (0 < 1) {
                            var v92 int = fi81
                            _ = v92
                        } else {
                            v4 = append(v4, (-(fi81)))
                        }
                        var v93 int8 = 100
                        _ = v93
                    }
                    for fi94 := 0; fi94 < 1; fi94++ {
                        v2 = "foo bar"
                        fmt.Println(v23, v77)
                        _ = fi94
                    }
                }
                v23 = v23
            }
            v95 := GenS11{F0: -32768, F1: 699, F2: 0}
            _ = v95
        }
        _ = fi81
    }
    switch 32767 {
    case 4:
        v3 += v3
    case 700:
        if (!(true && false)) {
            v23 -= (v23 / (((v23 | (v23 - v23))) | 1))
            if false {
                var v96 int = 267
                v77 -= 32767
                for fi97 := 0; fi97 < 7; fi97++ {
                    v98 := GenS12{F0: 97, F1: 32767}
                    switch v96 {
                    case -650:
                        var v99 float64 = 0.1
                        _ = v99
                    case 859:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 0:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        var v100 int64 = 42
                        _ = v100
                    }
                    if ((!true) || false) {
                        fmt.Println(v23, v96, v23)
                        var v101 uint16 = 76
                        _ = v101
                    }
                    _ = fi97
                    _ = v98
                }
                _ = v96
            } else {
                _ = len(v4)
            }
        }
    case 9223372036854775807:
        if ((false && (!true)) && true) {
            v102 := GenS13{F0: 127, F1: "foo bar"}
            _ = v102
        } else {
            for fi103 := 0; fi103 < 12; fi103++ {
                _ = len(v4)
                switch v2 {
                case "foo bar":
                    v2 = v1
                case "\n\t":
                    var v104 float32 = (float32(1e10) + (1.5 / (1.5 / ((((0.0 + (1e10 - 3.14159)) * 3.14159) * (-2.25 * 3.14159)) * 0.1))))
                    _ = v104
                }
                if len(v4) > 0 {
                    v4[0] = fi103
                }
                _ = fi103
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v77:", v77)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v77
    _ = v80
}
