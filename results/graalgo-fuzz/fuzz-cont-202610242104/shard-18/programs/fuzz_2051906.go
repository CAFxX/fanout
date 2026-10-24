package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 bool; F2 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint8; F1 uint64; F2 bool }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint16 = 493
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    var v4 float32 = 1.5
    var v5 string = "hello"
    fmt.Println(v1, v5)
    for fi6 := 0; fi6 < 9; fi6++ {
        if (1000 == -1) {
            for fi7 := 0; fi7 < 17; fi7++ {
                v5 = v5
                _ = fi7
            }
            v4 = v4
        }
        _ = fi6
    }
    v5 = v5
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint = 18446744073709551615
        if ((!(!("\n\t" == "hello"))) || true) {
            fmt.Println(v4)
            fmt.Println(v3)
        } else {
            v10 := []int{fi8}
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    if ((v1 * v1) != v1) {
        fmt.Println(v5, v5)
        var v11 float64 = 0.1
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            v1 += (v1 | 255)
            if (v4 <= 1e10) {
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint16 = v1
                    _ = fi13
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS1{F0: 100, F1: 1}
                    if false {
                        var v17 int32 = v3
                        switch "a" {
                        case "":
                            v17 -= (v17 & v17)
                        case "\n\t":
                            v4 += -2.25
                        case "foo bar":
                            fmt.Println(v5, v4)
                        }
                        v3 = v17
                        if false {
                            v18 := GenS2{F0: true, F1: v1, F2: 2147483648}
                            fmt.Println(v4, v17, fi15)
                            _ = v18
                        } else {
                            v4 -= (v4 + v4)
                        }
                        _ = v17
                    }
                    v19 := v2[fi12]
                    _ = fi15
                    _ = v16
                    _ = v19
                }
                switch fi12 {
                case -1:
                    var v20 string = "hello"
                    _ = v20
                default:
                    switch v5 {
                    case "a":
                        v21 := GenS3{F0: 1}
                        _ = v21
                    case "":
                        v1 = v1
                    case "foo bar":
                        switch ("" + "") {
                        case "a":
                            fmt.Println(v3, v4)
                        case "\n\t":
                            v4 = float32((1.5 + 1e10))
                        case "":
                            var v22 int32 = -604
                            _ = v22
                        default:
                            for fi23 := 0; fi23 < 12; fi23++ {
                                if ((v3 ^ v3) > -297) {
                                    var v24 float32 = (1.5 + v4)
                                    v25 := GenS1{F0: int8(v1), F1: uint32(v1)}
                                    _ = v24
                                    _ = v25
                                } else {
                                    if (false || true) {
                                        var v26 string = v5
                                        v26 = "hello"
                                        _ = v26
                                    }
                                }
                                fmt.Println(v5, v5, v3)
                                v1 += 32767
                                var v27 bool = true
                                v27 = v27
                                _ = fi23
                                _ = v27
                            }
                        }
                    default:
                        if false {
                            switch v5 {
                            case "foo bar":
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    delete(v2, (fi12 - 480))
                                    v29 := GenS3{F0: -59}
                                    v30 := GenS4{F0: 1}
                                    switch (^(fi28)) {
                                    case -247:
                                        fmt.Println(fi12, fi28)
                                        v31 := GenS5{F0: (v3 / (((v3 | (v3 - v3))) | 1)), F1: 0, F2: v1}
                                        _ = v31
                                    case 9223372036854775807:
                                        v4 += 0.1
                                    }
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            case "":
                                fmt.Println(v1, v4, v4)
                            }
                        }
                    }
                }
            }
            _ = fi12
        }
    }
    if true {
        for fi32 := 0; fi32 < 7; fi32++ {
            delete(v2, fi32)
            for fi33 := 0; fi33 < 10; fi33++ {
                switch ("a" + (v5 + (("foo bar" + "\n\t") + "a"))) {
                case "hello":
                    switch "hello" {
                    case "":
                        v34 := GenS6{F0: 42}
                        _ = v34
                    default:
                        switch (fi33 & (fi32 - fi33)) {
                        case 14:
                            v4 += -2.25
                        case 1:
                            fmt.Println(v1)
                        case -2147483648:
                            var v35 int8 = -128
                            _ = v35
                        default:
                            var v36 int8 = -36
                            _ = v36
                        }
                    }
                case "":
                    var v37 int32 = (-(v3))
                    _ = v37
                default:
                    v4 -= v4
                }
                _ = fi33
            }
            _ = fi32
        }
        var v38 int64 = 42
        _ = v38
    } else {
        for fi39 := 0; fi39 < 9; fi39++ {
            if false {
                for fi40 := 0; fi40 < 16; fi40++ {
                    v5 = ""
                    v1 -= v1
                    var v41 int64 = -1
                    _ = fi40
                    _ = v41
                }
                for fi42 := 0; fi42 < 3; fi42++ {
                    if true {
                        switch (fi39 % (((fi39 / (((fi42 * (fi42 >> ((0) & 63)))) | 1))) | 1)) {
                        case -96:
                            v43 := []int{(fi39 - (fi39 % ((fi42) | 1)))}
                            _ = v43
                        case -9:
                            v44 := GenS5{F0: (v3 | -1), F1: 741, F2: 32767}
                            _ = v44
                        case 0:
                            switch v5 {
                            case "a":
                                if (-9223372036854775808 >= (fi42 & (-(fi42)))) {
                                    if (v3 == int32(fi42)) {
                                        v1 -= v1
                                        v5 = v5
                                    }
                                    var v45 uint64 = 1
                                    _ = v45
                                } else {
                                    var v46 int = 32767
                                    var v47 int32 = v3
                                    _ = v46
                                    _ = v47
                                }
                            case "foo bar":
                                if (true || false) {
                                    if (v5 <= ((v5 + v5) + "hello")) {
                                        v3 -= v3
                                    } else {
                                        if false {
                                            v48 := GenS7{F0: 42, F1: false, F2: fi39}
                                            v49 := GenS1{F0: -1, F1: 478}
                                            _ = v48
                                            _ = v49
                                        } else {
                                            switch "a" {
                                            case "":
                                                fmt.Println(fi39, v3, fi39)
                                            default:
                                                switch (fi39 - (^(fi39))) {
                                                case -2147483648:
                                                    var v50 int8 = 100
                                                    _ = v50
                                                case 147:
                                                    var v51 float64 = -2.25
                                                    _ = v51
                                                default:
                                                    switch fi39 {
                                                    case 5:
                                                        v52 := GenS1{F0: 107, F1: 98}
                                                        _ = v52
                                                    case -2147483648:
                                                        var v53 uint32 = 65535
                                                        _ = v53
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v3, v5)
                                }
                            }
                        default:
                            if (712 > 42) {
                                var v54 bool = true
                                var v55 string = "hello"
                                v3 += (v3 ^ 0)
                                if (((v4 / (0.1 / v4)) * float32(1e10)) != ((v4 * 1e10) / v4)) {
                                    var v56 uint8 = 7
                                    fmt.Println(v1)
                                    v3 += v3
                                    _ = v56
                                }
                                var v57 uint8 = 2
                                v57 = 1
                                _ = v54
                                _ = v55
                                _ = v57
                            } else {
                                v58 := []int{(fi42 * -2147483648), (-(fi42)), fi39, (fi39 >> ((203) & 63))}
                                switch (fi39 % ((fi39) | 1)) {
                                case -2147483648:
                                    v59 := GenS8{F0: 2147483648}
                                    _ = v59
                                case 4294967295:
                                    switch fi39 {
                                    case 2147483647:
                                        switch (fi42 * fi39) {
                                        case 455:
                                            v4 -= v4
                                        case 32767:
                                            v60 := GenS9{F0: -32768}
                                            _ = v60
                                        case 4294967295:
                                            fmt.Println(fi42)
                                        }
                                    default:
                                        var v61 uint32 = 1
                                        _ = v61
                                    }
                                default:
                                    if len(v58) > 0 {
                                        v58[0] = fi42
                                    }
                                }
                                var v62 uint32 = 245
                                _ = v58
                                _ = v62
                            }
                        }
                        _ = len(v2)
                        v4 = 3.14159
                        if true {
                            var v63 uint = 4294967295
                            _ = v63
                        }
                        v1 -= 2
                        v64 := GenS10{F0: 1, F1: 126, F2: false}
                        _ = v64
                    }
                    switch v5 {
                    case "a":
                        v1 = (v1 - (^(v1)))
                    case "":
                        var v65 int16 = 1000
                        _ = v65
                    }
                    var v66 int16 = -1
                    _ = fi42
                    _ = v66
                }
            } else {
                if false {
                    var v67 int64 = 924
                    _ = len(v2)
                    fmt.Println(v4, v67, v3)
                    _ = v67
                }
                if (((!(42 > -9223372036854775808)) && true) || true) {
                    if true {
                        var v68 int32 = (-(v3))
                        v1 -= v1
                        switch 32767 {
                        case -199:
                            fmt.Println(v4, v5, fi39)
                        default:
                            var v69 int8 = 1
                            _ = v69
                        }
                        _ = v68
                    }
                } else {
                    delete(v2, fi39)
                    if false {
                        v70 := GenS11{F0: 233, F1: "a"}
                        _ = v70
                    }
                    v71 := []int{fi39}
                    _ = v71
                }
                if false {
                    delete(v2, (fi39 ^ (-(fi39))))
                } else {
                    if (true && (!false)) {
                        if true {
                            v72 := GenS2{F0: (42 < 88), F1: (^(v1)), F2: 42}
                            for fi73 := 0; fi73 < 10; fi73++ {
                                var v74 uint8 = 2
                                if false {
                                    v4 -= (0.0 / float32(-2.25))
                                    var v75 uint16 = 292
                                    _ = v75
                                }
                                _ = fi73
                                _ = v74
                            }
                            _ = v72
                        } else {
                            fmt.Println(v4)
                            if true {
                                for fi76 := 0; fi76 < 5; fi76++ {
                                    switch fi39 {
                                    case 772:
                                        switch ("\n\t" + ("a" + "a")) {
                                        case "foo bar":
                                            delete(v2, 32768)
                                        case "a":
                                            var v77 float32 = (1e10 - float32(0.0))
                                            v78 := GenS6{F0: 2}
                                            _ = v77
                                            _ = v78
                                        case "hello":
                                            _ = len(v2)
                                            switch v5 {
                                            case "":
                                                var v79 bool = false
                                                _ = v79
                                            default:
                                                fmt.Println(v4, fi76)
                                            }
                                            fmt.Println(v3)
                                        default:
                                            v80 := []int{4294967295}
                                            _ = v80
                                        }
                                    case -35:
                                        var v81 uint32 = 22
                                        _ = v81
                                    case -800:
                                        v5 = "foo bar"
                                    default:
                                        switch v5 {
                                        case "":
                                            fmt.Println(v4, fi39, fi39)
                                        }
                                    }
                                    _ = fi76
                                }
                            } else {
                                switch v5 {
                                case "foo bar":
                                    for fi82 := 0; fi82 < 19; fi82++ {
                                        var v83 uint8 = 128
                                        _ = fi82
                                        _ = v83
                                    }
                                case "":
                                    for fi84 := 0; fi84 < 13; fi84++ {
                                        if (float32(((1e10 - 0.1) + -2.25)) <= v4) {
                                            var v85 int8 = -1
                                            v1 += v1
                                            var v86 float32 = -2.25
                                            _ = v85
                                            _ = v86
                                        }
                                        _ = fi84
                                    }
                                }
                                for fi87 := 0; fi87 < 2; fi87++ {
                                    if ((1.5 * 1.5) < 0.0) {
                                        v4 = (((1e10 / v4) * v4) * float32(-2.25))
                                        var v88 int = 127
                                        if true {
                                            if (0 != 615) {
                                                var v89 int16 = -1
                                                switch (^(fi87)) {
                                                case 127:
                                                    var v90 int64 = 1
                                                    _ = v90
                                                case 380:
                                                    var v91 uint32 = 0
                                                    _ = v91
                                                default:
                                                    if true {
                                                        var v92 uint32 = 65535
                                                        v5 = v5
                                                        _ = v92
                                                    }
                                                }
                                                _ = v89
                                            } else {
                                                if (((3.14159 * v4) + 0.1) >= 1.5) {
                                                    v4 += v4
                                                    switch (v5 + (v5 + v5)) {
                                                    case "\n\t":
                                                        v93 := GenS5{F0: v3, F1: 18446744073709551615, F2: (v1 * (^(v1)))}
                                                        if false {
                                                            fmt.Println(fi87)
                                                        }
                                                        _ = v93
                                                    case "a":
                                                        var v94 uint = 0
                                                        _ = v94
                                                    }
                                                }
                                            }
                                            var v95 int = 9223372036854775807
                                            v95 -= (fi39 * fi87)
                                            _ = v95
                                        } else {
                                            if (193 == 1) {
                                                var v96 string = ("" + "a")
                                                _ = v96
                                            }
                                            v97 := GenS12{F0: 18446744073709551615, F1: 0}
                                            var v98 uint16 = (v1 << ((65535) & 63))
                                            _ = v97
                                            _ = v98
                                        }
                                        _ = v88
                                    }
                                    _ = fi87
                                }
                            }
                            v3 += -2147483648
                        }
                    }
                    v1 += 256
                    switch fi39 {
                    case -1:
                        switch (v5 + ((v5 + (("foo bar" + "a") + "a")) + "\n\t")) {
                        case "a":
                            v99 := GenS6{F0: 0}
                            _ = v99
                        case "foo bar":
                            var v100 int64 = 42
                            _ = v100
                        case "\n\t":
                            for fi101 := 0; fi101 < 12; fi101++ {
                                switch 4294967295 {
                                case 128:
                                    v102 := []int{32767}
                                    _ = v102
                                }
                                _ = fi101
                            }
                        default:
                            v103 := GenS13{F0: (^(v1)), F1: -1}
                            _ = v103
                        }
                    }
                }
            }
            v104 := GenS14{F0: (-(v3)), F1: 1}
            _ = fi39
            _ = v104
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
