package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string; F1 bool; F2 int16 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int32; F1 int16; F2 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 string; F2 int32 }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int16; F1 uint; F2 uint8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int32; F1 uint16 }
type GenS15 struct { F0 int16; F1 int; F2 uint }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint; F1 int8; F2 int16 }
type GenS18 struct { F0 uint8; F1 int }
type GenS19 struct { F0 int8; F1 int }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 uint8; F1 int16; F2 bool }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 string }
type GenS24 struct { F0 uint32; F1 uint; F2 bool }
type GenS25 struct { F0 uint; F1 int32; F2 int }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 int16; F1 int64; F2 uint }
type GenS28 struct { F0 uint32; F1 int16 }
func genFn0(p0 int8, p1 int) int64 {
    switch p1 {
    case -131:
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: 0}
            v7 := []int{(-(fi5)), (fi5 % ((fi5) | 1))}
            _ = fi5
            _ = v6
            _ = v7
        }
    case 711:
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(p1, p0, p1)
            if false {
                v9 := []int{p1, p1}
                p1 = 488
                _ = v9
            }
            _ = fi8
        }
    case -572:
        if (("foo bar" + ("a" + "foo bar")) < "") {
            fmt.Println(p1, p0)
        } else {
            v10 := GenS2{F0: ("hello" + "hello"), F1: 757}
            v11 := []int{128}
            v12 := GenS2{F0: "hello", F1: 94}
            _ = v10
            _ = v11
            _ = v12
        }
    }
    p1 -= p1
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int = (fi13 & fi13)
        switch "" {
        case "a":
            for fi15 := 0; fi15 < 18; fi15++ {
                if (934 != 1) {
                    var v16 uint16 = 1
                    fmt.Println(fi13)
                    _ = v16
                } else {
                    var v17 int = 128
                    var v18 float32 = -2.25
                    _ = v17
                    _ = v18
                }
                var v19 bool = true
                var v20 string = "a"
                _ = fi15
                _ = v19
                _ = v20
            }
        case "foo bar":
            if ((true || (65535 == 65535)) && true) {
                p0 -= 93
                for fi21 := 0; fi21 < 8; fi21++ {
                    v22 := GenS1{F0: -666}
                    _ = fi21
                    _ = v22
                }
                v23 := []int{(p1 & fi13), v14, p1, v14}
                _ = v23
            }
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint64 = 27
    v2 := []string{("hello" + "\n\t"), (("a" + ("\n\t" + ("a" + ("hello" + ("a" + (("hello" + "hello") + "foo bar")))))) + ""), ((("" + ("hello" + "")) + "a") + "foo bar"), ("hello" + (("foo bar" + "\n\t") + ("" + "foo bar")))}
    var v3 uint8 = 42
    var v4 uint32 = 469
    if true {
        v24 := GenS2{F0: "\n\t", F1: 9}
        _ = v24
    } else {
        if (42 > 0) {
            v1 = (v1 * 2)
        }
    }
    v2 = append(v2, "")
    var v25 int32 = 206
    for fi26 := 0; fi26 < 15; fi26++ {
        v3 += (^(v3))
        _ = fi26
    }
    v1 = 484
    var v27 int8 = -1
    fmt.Println(v27, v25, v25)
    for fi28 := 0; fi28 < 4; fi28++ {
        v3 = v3
        v27 += (v27 / ((100) | 1))
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = fi28
    }
    v2 = append(v2, (("a" + ("a" + "a")) + ("foo bar" + "")))
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS1{F0: -116}
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 9223372036854775807}
    if true {
        switch "foo bar" {
        case "hello":
            if true {
                fmt.Println(v3)
                v32 := GenS3{F0: "a", F1: false, F2: 1}
                var v33 int32 = (v25 / ((int32(v4)) | 1))
                _ = v32
                _ = v33
            }
        }
        var v34 string = ("\n\t" + "a")
        _ = v34
    } else {
        switch 0 {
        case -1:
            v2 = append(v2, "a")
        case 42:
            for fi35 := 0; fi35 < 10; fi35++ {
                if true {
                    if ((^(v1)) <= (v1 + v1)) {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            if false {
                                _ = len(v2)
                                var v37 uint16 = 877
                                _ = v37
                            } else {
                                var v38 uint16 = 0
                                _ = v38
                            }
                            _ = fi36
                        }
                        v39 := GenS4{F0: true, F1: (v4 >> ((282) & 63))}
                        var v40 int16 = 1
                        _ = v39
                        _ = v40
                    } else {
                        var v41 uint16 = 144
                        _ = v41
                    }
                    v25 = v25
                } else {
                    for fi42 := 0; fi42 < 6; fi42++ {
                        switch (-(fi42)) {
                        case 32768:
                            switch "\n\t" {
                            case "":
                                if false {
                                    var v43 int8 = 42
                                    _ = v43
                                } else {
                                    var v44 int = fi42
                                    switch fi42 {
                                    case 32768:
                                        v45 := GenS3{F0: "a", F1: false, F2: int16(v4)}
                                        _ = v45
                                    case 32767:
                                        switch fi42 {
                                        case 1:
                                            switch ("a" + ("" + "a")) {
                                            case "foo bar":
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(v25, v44, v44)
                                                }
                                            }
                                        case 515:
                                            fmt.Println(v4, fi42)
                                        }
                                    case 275:
                                        v2 = append(v2, "")
                                    }
                                    _ = v44
                                }
                            case "foo bar":
                                if (!(!(!(true && (1e10 != (0.1 - float64(-2.25))))))) {
                                    switch ("\n\t" + (((("foo bar" + "foo bar") + "\n\t") + "hello") + "")) {
                                    case "a":
                                        if false {
                                            switch "foo bar" {
                                            case "hello":
                                                v1 = (v1 / (((v1 % (((v1 + v1)) | 1))) | 1))
                                            case "\n\t":
                                                v1 -= 604
                                            default:
                                                if (!(0.1 > 0.1)) {
                                                    v3 -= 0
                                                    v3 += v3
                                                } else {
                                                    if false {
                                                        var v46 int16 = 1
                                                        fmt.Println(v4, v4, fi35)
                                                        _ = v46
                                                    } else {
                                                        v3 = v3
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v47 := GenS5{F0: uint(v3), F1: uint32(v1)}
                                                    if false {
                                                        var v48 int16 = -32768
                                                        _ = v48
                                                    } else {
                                                        v1 += v1
                                                        var v49 bool = (!((v27 / ((v27) | 1)) >= (v27 & (v27 | (-(v27))))))
                                                        v27 -= (v27 + v27)
                                                        var v50 int8 = v27
                                                        v51 := GenS6{F0: 2147483647, F1: -32768, F2: (fi35 + 0)}
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    _ = v47
                                                }
                                            }
                                            switch (^(fi42)) {
                                            case 127:
                                                v4 = (v4 << ((4294967295) & 63))
                                            default:
                                                switch fi35 {
                                                case 127:
                                                    var v52 float64 = (1e10 - -2.25)
                                                    switch (^(fi42)) {
                                                    case 32767:
                                                        var v53 int16 = 1000
                                                        _ = v53
                                                    }
                                                    v1 = (v1 ^ (v1 % ((4294967295) | 1)))
                                                    if (2 >= 142) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v54 uint32 = (^(v4))
                                                        _ = v54
                                                    }
                                                    v3 = (v3 | v3)
                                                    v4 = (v4 ^ v4)
                                                    _ = v52
                                                default:
                                                    var v55 string = "a"
                                                    if (fi35 == (fi42 / (((fi35 * (fi42 & 2147483647))) | 1))) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v55 = v55
                                                        if true {
                                                            if (128 == v3) {
                                                                switch v55 {
                                                                case "hello":
                                                                    var v56 uint = 868
                                                                    _ = v56
                                                                case "a":
                                                                    var v57 uint64 = uint64(v4)
                                                                    _ = v57
                                                                default:
                                                                    switch v55 {
                                                                    case "\n\t":
                                                                        switch int(v25) {
                                                                        case 0:
                                                                            if true {
                                                                                if (v4 < 2147483648) {
                                                                                    switch (fi42 << ((18446744073709551615) & 63)) {
                                                                                    case -9223372036854775808:
                                                                                        v58 := GenS7{F0: v3}
                                                                                        _ = v58
                                                                                    case -755:
                                                                                        v55 = (v55 + "a")
                                                                                    default:
                                                                                        v59 := GenS8{F0: genFn0(100, (fi42 + -9223372036854775808)), F1: "\n\t", F2: (^(v25))}
                                                                                        _ = v59
                                                                                    }
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                } else {
                                                                                    v27 = 127
                                                                                }
                                                                                switch fi35 {
                                                                                case 4294967295:
                                                                                    if (!((((1.5 + 0.0) <= 1e10) || false) || false)) {
                                                                                        switch v55 {
                                                                                        case "hello":
                                                                                            _ = len(v2)
                                                                                        case "\n\t":
                                                                                            if ((65535 >= 0) || false) {
                                                                                                if false {
                                                                                                    v55 = v55
                                                                                                    var v60 int32 = -2147483648
                                                                                                    var v61 int8 = v27
                                                                                                    _ = v60
                                                                                                    _ = v61
                                                                                                } else {
                                                                                                    fmt.Println(v55, v3, v3)
                                                                                                    var v62 uint32 = v4
                                                                                                    v63 := GenS9{F0: v1, F1: v25}
                                                                                                    switch v55 {
                                                                                                    case "hello":
                                                                                                        if false {
                                                                                                            if (!true) {
                                                                                                                v27 += v27
                                                                                                                v2 = append(v2, v55)
                                                                                                            }
                                                                                                            v4 -= (v62 ^ (v4 ^ 655))
                                                                                                            v3 = v3
                                                                                                        }
                                                                                                    default:
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    }
                                                                                                    var v64 int = (fi42 + fi35)
                                                                                                    _ = len(v2)
                                                                                                    _ = v62
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                                var v65 string = ((v55 + v55) + "\n\t")
                                                                                                _ = v65
                                                                                            } else {
                                                                                                v1 = uint64(v4)
                                                                                                if (0 == 631) {
                                                                                                    _ = len(v2)
                                                                                                    if ((v25 / ((v25) | 1)) != v25) {
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = v55
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v66 float64 = 1.5
                                                                                                            var v67 uint16 = 42
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = (v55 + (v55 + ("foo bar" + (("\n\t" + ((("foo bar" + (v55 + (v55 + v55))) + v55) + "a")) + v55))))
                                                                                                            }
                                                                                                            _ = v66
                                                                                                            _ = v67
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v68 uint16 = 42
                                                                                                _ = v68
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    switch 0 {
                                                                                    case -2147483648:
                                                                                        v69 := GenS7{F0: v3}
                                                                                        _ = v69
                                                                                    case 4294967295:
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = v55
                                                                                        }
                                                                                    default:
                                                                                        var v70 bool = true
                                                                                        _ = v70
                                                                                    }
                                                                                }
                                                                            }
                                                                            v71 := GenS6{F0: (v25 % ((v25) | 1)), F1: -289, F2: fi35}
                                                                            _ = v71
                                                                        case 42:
                                                                            var v72 int = (fi35 >> ((530) & 63))
                                                                            _ = v72
                                                                        case 4294967295:
                                                                            fmt.Println(v4, v1, fi42)
                                                                        }
                                                                    case "hello":
                                                                        v73 := GenS10{F0: 810}
                                                                        _ = v73
                                                                    case "":
                                                                        v74 := GenS10{F0: 350}
                                                                        _ = v74
                                                                    default:
                                                                        v55 = "hello"
                                                                    }
                                                                }
                                                                var v75 float32 = (0.1 + -2.25)
                                                                _ = v75
                                                            } else {
                                                                if (!((598 < 65535) || false)) {
                                                                    var v76 int64 = int64(v1)
                                                                    var v77 uint16 = 255
                                                                    v2 = append(v2, "a")
                                                                    _ = v76
                                                                    _ = v77
                                                                } else {
                                                                    if false {
                                                                        if (((v25 / ((v25) | 1)) < (v25 / (((v25 + 157)) | 1))) && true) {
                                                                            fmt.Println(v3)
                                                                            fmt.Println(v25, v4)
                                                                            v1 += v1
                                                                        } else {
                                                                            v4 = (^(v4))
                                                                        }
                                                                    } else {
                                                                        v78 := GenS7{F0: v3}
                                                                        _ = v78
                                                                    }
                                                                    v1 = (v1 | 700)
                                                                    v1 = (^(v1))
                                                                }
                                                                var v79 bool = (1 == v3)
                                                                _ = v79
                                                            }
                                                        } else {
                                                            v27 = v27
                                                            var v80 uint = 561
                                                            if (!true) {
                                                                v81 := GenS11{F0: 260}
                                                                if len(v2) > 0 {
                                                                    v2[0] = v55
                                                                }
                                                                _ = v81
                                                            } else {
                                                                if (2 != 144) {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v2 = append(v2, v55)
                                                                    case "hello":
                                                                        if true {
                                                                            v80 = 1
                                                                            if true {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                switch v55 {
                                                                                case "\n\t":
                                                                                    var v82 bool = false
                                                                                    _ = v82
                                                                                case "hello":
                                                                                    var v83 string = v55
                                                                                    _ = v83
                                                                                }
                                                                                v84 := GenS12{F0: 0, F1: 255, F2: (v3 * v3)}
                                                                                switch ((v55 + "hello") + "") {
                                                                                case "\n\t":
                                                                                    v55 = v55
                                                                                default:
                                                                                    var v85 uint8 = (v3 % ((v3) | 1))
                                                                                    _ = v85
                                                                                }
                                                                                v3 = 128
                                                                                _ = v84
                                                                            } else {
                                                                                if (3.14159 == 0.1) {
                                                                                    var v86 int64 = genFn0(int8(v25), (-(fi42)))
                                                                                    _ = v86
                                                                                } else {
                                                                                    switch v55 {
                                                                                    case "a":
                                                                                        v87 := GenS2{F0: (v55 + v55), F1: 42}
                                                                                        _ = v87
                                                                                    case "foo bar":
                                                                                        fmt.Println(v80, v80, v4)
                                                                                        fmt.Println(v3, fi42, v25)
                                                                                        if (-1 <= 385) {
                                                                                            if (v25 >= v25) {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = v55
                                                                                                }
                                                                                            }
                                                                                            switch v55 {
                                                                                            case "hello":
                                                                                                switch fi35 {
                                                                                                case 32767:
                                                                                                    v25 = v25
                                                                                                case -226:
                                                                                                    fmt.Println(v55, v1)
                                                                                                case 481:
                                                                                                    if true {
                                                                                                        v88 := GenS6{F0: int32(v3), F1: 361, F2: fi35}
                                                                                                        switch fi35 {
                                                                                                        case 128:
                                                                                                            v3 += v3
                                                                                                        case 4294967295:
                                                                                                            v25 += (^(v25))
                                                                                                        case -1:
                                                                                                            var v89 uint8 = v3
                                                                                                            fmt.Println(v89)
                                                                                                            v25 += (v25 + v25)
                                                                                                            switch "a" {
                                                                                                            case "hello":
                                                                                                                fmt.Println(fi42, fi42, v25)
                                                                                                            case "\n\t":
                                                                                                                if false {
                                                                                                                    switch (fi42 * 127) {
                                                                                                                    case -9223372036854775808:
                                                                                                                        v4 += v4
                                                                                                                    }
                                                                                                                    v2 = append(v2, (v55 + v55))
                                                                                                                    v55 = (v55 + v55)
                                                                                                                }
                                                                                                            case "":
                                                                                                                fmt.Println(v89, v4, v89)
                                                                                                            }
                                                                                                            v4 = 301
                                                                                                            _ = v89
                                                                                                        }
                                                                                                        _ = v88
                                                                                                    } else {
                                                                                                        var v90 uint16 = 32767
                                                                                                        v3 += (v3 & 93)
                                                                                                        v91 := GenS13{F0: v80}
                                                                                                        _ = v90
                                                                                                        _ = v91
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                switch v55 {
                                                                                                case "":
                                                                                                    v25 += v25
                                                                                                case "foo bar":
                                                                                                    v4 += (v4 + 65535)
                                                                                                case "hello":
                                                                                                    switch v55 {
                                                                                                    case "foo bar":
                                                                                                        if ((v25 >> (((v80 - v80)) & 63)) > (v25 + (-(v25)))) {
                                                                                                            var v92 uint16 = 55
                                                                                                            switch -8 {
                                                                                                            case 9223372036854775807:
                                                                                                                if ((v25 * (v25 - v25)) <= v25) {
                                                                                                                    var v93 string = v55
                                                                                                                    var v94 int16 = 32767
                                                                                                                    _ = v93
                                                                                                                    _ = v94
                                                                                                                } else {
                                                                                                                    if ((!((true && ((true && (false || true)) && false)) && true)) || (v92 <= (v92 << ((0) & 63)))) {
                                                                                                                        v95 := GenS14{F0: (v25 | v25), F1: 2}
                                                                                                                        _ = v95
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = v55
                                                                                                            }
                                                                                                            _ = v92
                                                                                                        }
                                                                                                    case "":
                                                                                                        if false {
                                                                                                            v96 := GenS2{F0: ("hello" + v55), F1: 65535}
                                                                                                            _ = v96
                                                                                                        }
                                                                                                    case "a":
                                                                                                        if (-2.25 < -2.25) {
                                                                                                            var v97 int16 = 714
                                                                                                            var v98 bool = true
                                                                                                            switch v55 {
                                                                                                            case "hello":
                                                                                                                switch "" {
                                                                                                                case "a":
                                                                                                                    var v99 int32 = v25
                                                                                                                    _ = v99
                                                                                                                default:
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v2 = append(v2, ((v55 + v55) + (("" + v55) + ("\n\t" + "a"))))
                                                                                                                var v100 int16 = v97
                                                                                                                v3 = (v3 / ((v3) | 1))
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = (v55 + "hello")
                                                                                                                }
                                                                                                                if len(v2) >= 2 {
                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                }
                                                                                                                fmt.Println(fi35, v25)
                                                                                                                v101 := GenS15{F0: 1000, F1: 32767, F2: v80}
                                                                                                                _ = v100
                                                                                                                _ = v101
                                                                                                            default:
                                                                                                                var v102 int64 = int64(v97)
                                                                                                                _ = v102
                                                                                                            }
                                                                                                            _ = v97
                                                                                                            _ = v98
                                                                                                        } else {
                                                                                                            var v103 int32 = (v25 - 42)
                                                                                                            fmt.Println(fi42, v4, v4)
                                                                                                            _ = v103
                                                                                                        }
                                                                                                    default:
                                                                                                        fmt.Println(v80, fi35, v1)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if ((true && ((v1 / ((v1) | 1)) != v1)) && (255 < 32767)) {
                                                                                                fmt.Println(fi35)
                                                                                                if (!(("\n\t" + "hello") >= "\n\t")) {
                                                                                                    v3 -= 209
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(fi35, v1, v80)
                                                                                                switch v55 {
                                                                                                case "\n\t":
                                                                                                    _ = len(v2)
                                                                                                case "foo bar":
                                                                                                    v104 := GenS16{F0: 42}
                                                                                                    _ = v104
                                                                                                case "":
                                                                                                    v2 = append(v2, "hello")
                                                                                                }
                                                                                            }
                                                                                            if (!((^(v3)) < v3)) {
                                                                                                var v105 uint32 = (v4 | v4)
                                                                                                v106 := GenS16{F0: 0}
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            }
                                                                                            var v107 int16 = -324
                                                                                            v108 := GenS15{F0: 32767, F1: 4294967295, F2: v80}
                                                                                            switch v55 {
                                                                                            case "hello":
                                                                                                var v109 string = v55
                                                                                                _ = v109
                                                                                            case "":
                                                                                                fmt.Println(v80)
                                                                                            case "\n\t":
                                                                                                switch (v55 + "a") {
                                                                                                case "\n\t":
                                                                                                    if (((false || false) && true) || true) {
                                                                                                        var v110 int16 = int16(v3)
                                                                                                        _ = v110
                                                                                                    }
                                                                                                case "a":
                                                                                                    switch (v55 + ((v55 + v55) + v55)) {
                                                                                                    case "foo bar":
                                                                                                        if (true && true) {
                                                                                                            var v111 uint8 = (v3 << (((v80 | 382)) & 63))
                                                                                                            var v112 uint64 = v1
                                                                                                            _ = v111
                                                                                                            _ = v112
                                                                                                        }
                                                                                                    default:
                                                                                                        switch fi42 {
                                                                                                        case -1:
                                                                                                            v2 = append(v2, "foo bar")
                                                                                                        case 32768:
                                                                                                            if (v27 >= (v27 | 1)) {
                                                                                                                var v113 int8 = 42
                                                                                                                var v114 float32 = float32(3.14159)
                                                                                                                fmt.Println(v25, v80, v25)
                                                                                                                _ = v113
                                                                                                                _ = v114
                                                                                                            } else {
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = v55
                                                                                                                }
                                                                                                                if ("\n\t" < "a") {
                                                                                                                    var v115 int16 = int16(v80)
                                                                                                                    v2 = append(v2, "hello")
                                                                                                                    _ = v115
                                                                                                                }
                                                                                                                var v116 int32 = (v25 ^ (v25 << ((v80) & 63)))
                                                                                                                v2 = append(v2, v55)
                                                                                                                v117 := GenS12{F0: (v107 * 32767), F1: v80, F2: (^(v3))}
                                                                                                                _ = v116
                                                                                                                _ = v117
                                                                                                            }
                                                                                                        case 679:
                                                                                                            v118 := GenS1{F0: 0}
                                                                                                            _ = v118
                                                                                                        }
                                                                                                    }
                                                                                                case "":
                                                                                                    var v119 uint16 = 42
                                                                                                    _ = v119
                                                                                                default:
                                                                                                    fmt.Println(fi42)
                                                                                                }
                                                                                            default:
                                                                                                var v120 int16 = v107
                                                                                                _ = v120
                                                                                            }
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        }
                                                                                    default:
                                                                                        switch fi42 {
                                                                                        case 127:
                                                                                            var v121 int8 = v27
                                                                                            _ = v121
                                                                                        case -537:
                                                                                            fmt.Println(v55, v55, v55)
                                                                                        case -846:
                                                                                            if (false || false) {
                                                                                                var v122 uint = (v80 - v80)
                                                                                                fmt.Println(v80, fi35)
                                                                                                _ = v122
                                                                                            }
                                                                                        default:
                                                                                            v4 -= 948
                                                                                        }
                                                                                    }
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v123 uint16 = 255
                                                                            _ = v123
                                                                        }
                                                                    default:
                                                                        if (350 <= v4) {
                                                                            v27 += (v27 / ((int8(v25)) | 1))
                                                                        } else {
                                                                            var v124 int32 = v25
                                                                            var v125 float32 = 1e10
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v124
                                                                            _ = v125
                                                                        }
                                                                    }
                                                                }
                                                                if (!true) {
                                                                    _ = len(v2)
                                                                }
                                                                v126 := GenS12{F0: -1, F1: (v80 >> (((^(v80))) & 63)), F2: v3}
                                                                _ = v126
                                                            }
                                                            _ = v80
                                                        }
                                                    } else {
                                                        var v127 int8 = v27
                                                        var v128 uint8 = 0
                                                        switch "" {
                                                        case "\n\t":
                                                            v129 := GenS17{F0: 2, F1: (v27 % ((1) | 1)), F2: 42}
                                                            _ = v129
                                                        case "":
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v130 float32 = (0.1 * float32((3.14159 * -2.25)))
                                                                _ = v130
                                                            case "foo bar":
                                                                fmt.Println(v55, v55, fi42)
                                                                if ((v25 > v25) || ((v4 - 4294967295) < 312)) {
                                                                    v2 = append(v2, (v55 + v55))
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                } else {
                                                                    switch (fi42 & (fi35 | -637)) {
                                                                    case 128:
                                                                        v131 := GenS18{F0: (v3 & (^(v3))), F1: fi35}
                                                                        _ = v131
                                                                    default:
                                                                        switch (^(fi42)) {
                                                                        case -9223372036854775808:
                                                                            v25 += (^(v25))
                                                                        case -669:
                                                                            v132 := GenS1{F0: 83}
                                                                            _ = v132
                                                                        case 4294967295:
                                                                            v27 -= -1
                                                                            var v133 int64 = 0
                                                                            _ = v133
                                                                        }
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            var v134 float64 = (0.1 / (-2.25 + float64((0.1 + 0.1))))
                                                                            _ = v134
                                                                        default:
                                                                            fmt.Println(v3)
                                                                        }
                                                                        if (!false) {
                                                                            v135 := GenS18{F0: (^(v128)), F1: -68}
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v135
                                                                        }
                                                                    }
                                                                }
                                                            case "":
                                                                var v136 int = (fi35 | (fi42 * (fi35 / ((0) | 1))))
                                                                _ = v136
                                                            default:
                                                                v55 = v55
                                                            }
                                                        default:
                                                            var v137 int32 = 1
                                                            _ = v137
                                                        }
                                                        _ = v127
                                                        _ = v128
                                                    }
                                                    v1 += v1
                                                    _ = v55
                                                }
                                            }
                                            var v138 uint64 = 4294967295
                                            switch (("foo bar" + "\n\t") + "foo bar") {
                                            case "foo bar":
                                                if (v27 == v27) {
                                                    if false {
                                                        var v139 bool = false
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v25 = -2147483648
                                                        _ = v139
                                                    } else {
                                                        switch fi42 {
                                                        case -80:
                                                            var v140 int8 = v27
                                                            _ = v140
                                                        }
                                                        switch (fi42 >> ((1) & 63)) {
                                                        case -1:
                                                            v3 = (^(v3))
                                                            v27 = -128
                                                            if true {
                                                                var v141 int8 = (v27 | (v27 - (v27 >> ((96) & 63))))
                                                                if len(v2) > 0 {
                                                                    v2[0] = "a"
                                                                }
                                                                v141 -= v27
                                                                if (!(366 == 42)) {
                                                                    if false {
                                                                        if ((!(18446744073709551615 > 323)) || (!false)) {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ((("" + ("a" + "hello")) + "") + "\n\t")
                                                                            }
                                                                            v138 = 2
                                                                            v141 = 42
                                                                        }
                                                                        v141 += 42
                                                                        switch 36 {
                                                                        case 1:
                                                                            v142 := GenS19{F0: 42, F1: fi42}
                                                                            _ = v142
                                                                        case 637:
                                                                            fmt.Println(v141, fi42)
                                                                            var v143 int = fi35
                                                                            _ = v143
                                                                        default:
                                                                            v141 += v27
                                                                        }
                                                                    }
                                                                    if ((v1 - (v138 % ((2) | 1))) >= 2) {
                                                                        var v144 int16 = -321
                                                                        if (!(false || false)) {
                                                                            fmt.Println(v138, v144, v138)
                                                                            if false {
                                                                                var v145 uint64 = v1
                                                                                v146 := GenS20{F0: 9223372036854775807}
                                                                                v147 := GenS7{F0: v3}
                                                                                _ = v145
                                                                                _ = v146
                                                                                _ = v147
                                                                            } else {
                                                                                fmt.Println(v1, v138)
                                                                            }
                                                                        } else {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            fmt.Println(v1, v4, v4)
                                                                            v148 := GenS21{F0: v3, F1: (v144 / ((int16(v27)) | 1)), F2: (!true)}
                                                                            v149 := GenS22{F0: 700}
                                                                            _ = v148
                                                                            _ = v149
                                                                        }
                                                                        switch (fi42 << ((0) & 63)) {
                                                                        case 42:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = "a"
                                                                            }
                                                                        }
                                                                        _ = v144
                                                                    }
                                                                }
                                                                var v150 uint8 = (v3 ^ 255)
                                                                _ = v141
                                                                _ = v150
                                                            } else {
                                                                switch fi42 {
                                                                case 9223372036854775807:
                                                                    var v151 int8 = (v27 - 42)
                                                                    _ = v151
                                                                case -9223372036854775808:
                                                                    v1 -= 42
                                                                }
                                                                v4 = uint32(v3)
                                                            }
                                                            switch fi35 {
                                                            case -9223372036854775808:
                                                                v152 := GenS6{F0: -1, F1: -32768, F2: (fi42 | (fi35 ^ 42))}
                                                                _ = v152
                                                            }
                                                        case -9223372036854775808:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            fmt.Println(v25)
                                                        case 1:
                                                            var v153 int = int(v25)
                                                            _ = v153
                                                        }
                                                    }
                                                    v1 -= (v138 / ((55) | 1))
                                                }
                                            }
                                            _ = v138
                                        }
                                    default:
                                        var v154 int64 = 0
                                        _ = v154
                                    }
                                    if true {
                                        v155 := GenS18{F0: v3, F1: fi42}
                                        fmt.Println(v3, v1, fi42)
                                        v156 := GenS20{F0: int64(v27)}
                                        _ = v155
                                        _ = v156
                                    } else {
                                        if (0 >= 0) {
                                            if (false && ((fi42 / (((-(fi35))) | 1)) <= fi42)) {
                                                v1 += v1
                                            } else {
                                                if true {
                                                    v2 = append(v2, "foo bar")
                                                    v3 -= v3
                                                    switch "foo bar" {
                                                    case "a":
                                                        if true {
                                                            var v157 uint16 = 763
                                                            v4 = (v4 * (v4 + v4))
                                                            _ = v157
                                                        } else {
                                                            v158 := GenS6{F0: (v25 + 42), F1: -32768, F2: fi42}
                                                            fmt.Println(v25)
                                                            v27 -= 127
                                                            _ = v158
                                                        }
                                                    case "\n\t":
                                                        switch -978 {
                                                        case -778:
                                                            var v159 uint8 = 0
                                                            _ = v159
                                                        case 181:
                                                            fmt.Println(v27, v1)
                                                        case -383:
                                                            var v160 uint32 = (v4 / ((664) | 1))
                                                            _ = v160
                                                        default:
                                                            if true {
                                                                var v161 int32 = (v25 - v25)
                                                                _ = v161
                                                            } else {
                                                                v1 += (v1 ^ 2)
                                                                var v162 uint8 = (v3 + 2)
                                                                _ = v162
                                                            }
                                                        }
                                                        fmt.Println(v3, v3)
                                                        _ = len(v2)
                                                    default:
                                                        if ((9223372036854775807 >= 42) || true) {
                                                            var v163 string = "hello"
                                                            if (880 >= 32767) {
                                                                var v164 uint8 = 1
                                                                v165 := GenS7{F0: (v164 - (v3 - (v164 % ((v164) | 1))))}
                                                                _ = v164
                                                                _ = v165
                                                            } else {
                                                                v27 += 42
                                                                v166 := GenS23{F0: v163}
                                                                _ = v166
                                                            }
                                                            var v167 int8 = v27
                                                            _ = v163
                                                            _ = v167
                                                        }
                                                    }
                                                }
                                            }
                                            var v168 int = (fi42 * (^(fi35)))
                                            _ = v168
                                        } else {
                                            var v169 int16 = 0
                                            _ = v169
                                        }
                                    }
                                }
                            case "hello":
                                var v170 uint32 = v4
                                v170 = (v4 & (v4 * (v170 ^ (^(v4)))))
                                _ = v170
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        case 9223372036854775807:
                            v3 = 0
                        default:
                            var v171 float32 = 0.0
                            _ = v171
                        }
                        _ = fi42
                    }
                    var v172 float32 = 1e10
                    switch fi35 {
                    case 797:
                        for fi173 := 0; fi173 < 11; fi173++ {
                            var v174 uint8 = 255
                            _ = fi173
                            _ = v174
                        }
                    case -22:
                        switch ("" + ("hello" + "a")) {
                        case "a":
                            var v175 int = fi35
                            _ = v175
                        case "hello":
                            var v176 int64 = genFn0(v27, fi35)
                            _ = v176
                        }
                    case 9223372036854775807:
                        v27 -= v27
                    }
                    _ = v172
                }
                var v177 int = (fi35 & (-(fi35)))
                for fi178 := 0; fi178 < 17; fi178++ {
                    var v179 int = 2147483647
                    var v180 uint16 = 42
                    var v181 bool = (!false)
                    switch "hello" {
                    case "a":
                        switch 1 {
                        case 2147483647:
                            var v182 uint32 = 42
                            _ = v182
                        case 9223372036854775807:
                            switch v177 {
                            case 2147483647:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        case 973:
                            v183 := GenS4{F0: v181, F1: 2147483648}
                            _ = v183
                        default:
                            v180 = (v180 << ((738) & 63))
                        }
                        v184 := GenS14{F0: -667, F1: (v180 + 65535)}
                        v1 = uint64(v3)
                        v181 = false
                        switch 127 {
                        case -819:
                            fmt.Println(v3, v27, fi178)
                        }
                        _ = v184
                    case "":
                        v185 := GenS12{F0: 1000, F1: 0, F2: (v3 + v3)}
                        _ = v185
                    default:
                        switch "hello" {
                        case "foo bar":
                            var v186 float32 = (1e10 / 1e10)
                            _ = v186
                        }
                    }
                    _ = fi178
                    _ = v179
                    _ = v180
                    _ = v181
                }
                _ = fi35
                _ = v177
            }
        case -675:
            switch "" {
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                if true {
                    var v187 uint16 = 677
                    for fi188 := 0; fi188 < 10; fi188++ {
                        v187 = v187
                        switch "a" {
                        case "hello":
                            switch "" {
                            case "foo bar":
                                if true {
                                    v187 = v187
                                    var v189 uint64 = uint64(v3)
                                    _ = v189
                                }
                            case "":
                                v187 += v187
                            default:
                                var v190 float64 = 0.1
                                _ = v190
                            }
                        case "foo bar":
                            var v191 int8 = (v27 & v27)
                            _ = v191
                        case "a":
                            v25 -= int32(v27)
                        }
                        _ = fi188
                    }
                    if false {
                        if false {
                            v192 := GenS24{F0: (^(v4)), F1: 4294967295, F2: (-32768 >= 474)}
                            var v193 float32 = float32(((1.5 - (((float64(0.0) / 1.5) / (-2.25 + 1e10)) * 1e10)) + 1e10))
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            _ = v192
                            _ = v193
                        }
                        if ((v1 * uint64(v187)) >= 486) {
                            if (821 != -240) {
                                var v194 uint16 = v187
                                _ = v194
                            } else {
                                v195 := GenS25{F0: 2, F1: (v25 * v25), F2: int(v25)}
                                if false {
                                    v196 := GenS16{F0: -1}
                                    fmt.Println(v4, v1, v187)
                                    _ = v196
                                } else {
                                    v4 -= (v4 * (v4 % (((v4 | uint32(v187))) | 1)))
                                }
                                v197 := GenS2{F0: "", F1: 2}
                                _ = v195
                                _ = v197
                            }
                        } else {
                            _ = len(v2)
                            for fi198 := 0; fi198 < 15; fi198++ {
                                var v199 string = ("hello" + "a")
                                if false {
                                    var v200 uint = 803
                                    var v201 uint64 = uint64(v187)
                                    _ = v200
                                    _ = v201
                                } else {
                                    var v202 int32 = -1
                                    for fi203 := 0; fi203 < 14; fi203++ {
                                        var v204 float32 = (3.14159 - 1.5)
                                        v27 += 63
                                        var v205 uint = 706
                                        _ = fi203
                                        _ = v204
                                        _ = v205
                                    }
                                    v202 += (v25 * -1)
                                    _ = v202
                                }
                                for fi206 := 0; fi206 < 17; fi206++ {
                                    switch (fi206 << ((917) & 63)) {
                                    case 128:
                                        if (!true) {
                                            v207 := GenS5{F0: 988, F1: (v4 * (v4 / (((v4 ^ uint32(v187))) | 1)))}
                                            v208 := GenS12{F0: 272, F1: 18446744073709551615, F2: 42}
                                            _ = v207
                                            _ = v208
                                        }
                                    case 32767:
                                        v3 = v3
                                    case -9223372036854775808:
                                        var v209 float64 = (0.1 - float64(0.1))
                                        _ = v209
                                    default:
                                        v1 -= (v1 * uint64(v3))
                                    }
                                    switch fi198 {
                                    case -618:
                                        if (0 != 946) {
                                            if true {
                                                v210 := GenS21{F0: v3, F1: -32768, F2: false}
                                                var v211 int64 = int64(v3)
                                                _ = v210
                                                _ = v211
                                            } else {
                                                if ((!true) && true) {
                                                    v212 := GenS16{F0: 668}
                                                    _ = v212
                                                } else {
                                                    var v213 int = fi206
                                                    v27 += int8(v3)
                                                    v2 = append(v2, ("" + v199))
                                                    _ = v213
                                                }
                                            }
                                            if (((v4 & 4294967295) != v4) || (true || ((^(v187)) >= v187))) {
                                                switch v199 {
                                                case "hello":
                                                    v2 = append(v2, v199)
                                                case "a":
                                                    v187 += v187
                                                case "":
                                                    if ((false || false) || true) {
                                                        fmt.Println(v3, fi198, v27)
                                                        v25 = -74
                                                        v1 = 1
                                                        switch (v199 + v199) {
                                                        case "hello":
                                                            if (v4 != (v4 / (((^(v4))) | 1))) {
                                                                fmt.Println(v25, v4, v3)
                                                            } else {
                                                                fmt.Println(v25, fi198, v4)
                                                                switch (v199 + v199) {
                                                                case "":
                                                                    v214 := GenS15{F0: 32767, F1: 9223372036854775807, F2: 2}
                                                                    _ = v214
                                                                case "\n\t":
                                                                    v3 += (^(v3))
                                                                }
                                                                if false {
                                                                    var v215 uint8 = (v3 | v3)
                                                                    _ = v215
                                                                }
                                                            }
                                                        default:
                                                            var v216 float64 = -2.25
                                                            _ = v216
                                                        }
                                                        v3 = 1
                                                    } else {
                                                        switch (fi206 ^ 0) {
                                                        case -288:
                                                            var v217 uint16 = uint16(v1)
                                                            _ = v217
                                                        case 128:
                                                            var v218 uint16 = (v187 & 32767)
                                                            _ = v218
                                                        case 32767:
                                                            v2 = append(v2, (v199 + (((v199 + ((v199 + v199) + "")) + (("foo bar" + v199) + "a")) + v199)))
                                                        default:
                                                            if len(v2) > 0 {
                                                                v2[0] = "foo bar"
                                                            }
                                                        }
                                                        v25 -= int32(v187)
                                                        if (false || (!true)) {
                                                            var v219 int32 = v25
                                                            var v220 uint = 727
                                                            _ = v219
                                                            _ = v220
                                                        } else {
                                                            v187 -= uint16(v3)
                                                            _ = len(v2)
                                                            fmt.Println(fi206, v27)
                                                        }
                                                        switch v199 {
                                                        case "":
                                                            var v221 int32 = (v25 - 2147483647)
                                                            _ = v221
                                                        case "\n\t":
                                                            v25 = (v25 | (v25 + v25))
                                                        case "hello":
                                                            if len(v2) > 0 {
                                                                v2[0] = "a"
                                                            }
                                                        default:
                                                            switch v199 {
                                                            case "a":
                                                                if len(v2) > 0 {
                                                                    v2[0] = v199
                                                                }
                                                            case "hello":
                                                                var v222 uint64 = (v1 / ((v1) | 1))
                                                                _ = v222
                                                            case "\n\t":
                                                                fmt.Println(fi206)
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v1 += (v1 << ((0) & 63))
                                                }
                                                var v223 float32 = 1.5
                                                _ = v223
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = ("\n\t" + v199)
                                            }
                                        } else {
                                            var v224 uint16 = 197
                                            fmt.Println(fi198, v3, fi198)
                                            var v225 uint32 = 2
                                            _ = v224
                                            _ = v225
                                        }
                                        v199 = v199
                                        if false {
                                            var v226 uint8 = 0
                                            v227 := GenS21{F0: (v226 % (((^(v3))) | 1)), F1: 994, F2: (!true)}
                                            if true {
                                                v228 := GenS26{F0: -94}
                                                if (!true) {
                                                    var v229 int16 = 773
                                                    fmt.Println(v25, v27)
                                                    _ = v229
                                                }
                                                _ = v228
                                            } else {
                                                switch v199 {
                                                case "\n\t":
                                                    fmt.Println(v187, v226, fi198)
                                                case "a":
                                                    v2 = append(v2, (v199 + "a"))
                                                }
                                                v1 = 934
                                            }
                                            _ = v226
                                            _ = v227
                                        } else {
                                            var v230 int32 = (v25 & v25)
                                            v231 := GenS17{F0: 0, F1: 0, F2: 32767}
                                            v27 = 127
                                            v3 = uint8(v187)
                                            v230 = 42
                                            v232 := GenS27{F0: 42, F1: 1, F2: 42}
                                            var v233 int64 = 1
                                            v234 := GenS2{F0: (((("hello" + (v199 + (v199 + ("\n\t" + ((v199 + v199) + ""))))) + "\n\t") + "a") + v199), F1: v187}
                                            _ = v230
                                            _ = v231
                                            _ = v232
                                            _ = v233
                                            _ = v234
                                        }
                                    case -1:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    v3 -= 128
                                    if (!(false && (1 < -220))) {
                                        v2 = append(v2, (v199 + "hello"))
                                    }
                                    v4 -= v4
                                    _ = fi206
                                }
                                _ = fi198
                                _ = v199
                            }
                        }
                        var v235 uint = 173
                        _ = v235
                    }
                    _ = v187
                } else {
                    var v236 int32 = v25
                    _ = v236
                }
            case "hello":
                var v237 string = "\n\t"
                _ = v237
            }
        }
        if (438 == uint(v3)) {
            if true {
                if (!((((0.0 + float64(0.0)) * (1e10 + (((((0.1 + 1e10) - -2.25) * 1e10) - 0.0) + -2.25))) - 3.14159) == (3.14159 / 3.14159))) {
                    v1 -= v1
                } else {
                    fmt.Println(v4, v3, v4)
                }
            } else {
                for fi238 := 0; fi238 < 1; fi238++ {
                    for fi239 := 0; fi239 < 13; fi239++ {
                        v3 = v3
                        if true {
                            v4 += 65535
                            if ((3.14159 + float32(-2.25)) > (0.0 * 0.1)) {
                                var v240 bool = ((-1 == v27) && (!false))
                                _ = v240
                            } else {
                                fmt.Println(v25, v25, v3)
                                switch fi239 {
                                case 128:
                                    var v241 float64 = (0.0 * 3.14159)
                                    _ = v241
                                }
                                v242 := GenS1{F0: 1}
                                _ = v242
                            }
                            var v243 int32 = v25
                            _ = v243
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    fmt.Println(fi238)
                                    fmt.Println(v4, v4)
                                } else {
                                    var v244 uint = 666
                                    _ = v244
                                }
                            }
                            var v245 uint64 = 15
                            _ = v245
                        }
                        v3 += 128
                        _ = fi239
                    }
                    if (v25 > v25) {
                        fmt.Println(v3, v27)
                        var v246 int32 = 42
                        _ = v246
                    }
                    _ = fi238
                }
            }
            fmt.Println(v3, v3, v25)
        } else {
            v2 = append(v2, "hello")
        }
    }
    for fi247 := 0; fi247 < 9; fi247++ {
        v248 := GenS28{F0: v4, F1: 0}
        fmt.Println(v1)
        _ = fi247
        _ = v248
    }
    switch 2147483647 {
    case 0:
        v25 = v25
    case 14:
        v4 -= (v4 % (((v4 - v4)) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
