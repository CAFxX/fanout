package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint; F1 uint32; F2 string }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint; F1 int; F2 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 uint, p1 int, p2 int16) uint64 {
    p2 -= (p2 | 0)
    v5 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}
func genFn1(p0 uint16, p1 bool) uint64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        p1 = p1
        v7 := []int{1, fi6, fi6, -514}
        switch -1 {
        case 32768:
            v8 := GenS1{F0: p0}
            _ = v8
        case 127:
            p0 = 32767
        case 589:
            v9 := GenS1{F0: 42}
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if ((!(((p1 || (("" + ("foo bar" + "\n\t")) == "a")) && true) && p1)) || p1) {
        p0 = 1
        var v10 uint64 = 843
        if (-1 > 0) {
            v10 += v10
        } else {
            var v11 string = "\n\t"
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 int = -1
                v14 := GenS1{F0: p0}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        _ = v10
    }
    var v15 int8 = 0
    _ = p0
    _ = p1
    _ = v15
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 int32 = 42
    v4 := []string{"hello", (("" + "") + ("foo bar" + "a")), "foo bar", ("foo bar" + "\n\t")}
    switch "\n\t" {
    case "\n\t":
        switch -9223372036854775808 {
        case -721:
            v3 += v3
        case -1:
            var v16 int16 = 1
            _ = v16
        }
    case "hello":
        var v17 int = -746
        _ = v17
    default:
        switch 694 {
        case 128:
            delete(v2, 32768)
        case 4294967295:
            var v18 bool = false
            _ = v18
        case 42:
            if false {
                switch 1 {
                case -1:
                    if (float64(((-2.25 * float32(0.0)) - 1.5)) < 0.0) {
                        v3 += (v3 + (v3 ^ -2147483648))
                        switch 32768 {
                        case -9223372036854775808:
                            for fi19 := 0; fi19 < 9; fi19++ {
                                v20 := GenS1{F0: 2}
                                var v21 bool = (true || true)
                                var v22 uint = 1
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        case 2147483647:
                            if len(v4) > 0 {
                                v4[0] = ("" + ((("" + "hello") + "hello") + ""))
                            }
                        }
                    } else {
                        fmt.Println(v3)
                        if ((127 <= -9223372036854775808) || (true && (0 < 2))) {
                            v3 += v3
                            if (!(v3 >= (v3 >> ((696) & 63)))) {
                                var v23 uint32 = 2147483647
                                _ = v23
                            } else {
                                for fi24 := 0; fi24 < 11; fi24++ {
                                    for fi25 := 0; fi25 < 10; fi25++ {
                                        switch ("hello" + ("\n\t" + "foo bar")) {
                                        case "\n\t":
                                            switch "foo bar" {
                                            case "foo bar":
                                                v26 := GenS2{F0: ""}
                                                _ = v26
                                            }
                                        case "foo bar":
                                            v27 := GenS3{F0: 65535}
                                            _ = v27
                                        case "a":
                                            v4 = append(v4, ("" + "a"))
                                        default:
                                            _ = len(v1)
                                        }
                                        var v28 float64 = 1e10
                                        _ = fi25
                                        _ = v28
                                    }
                                    v4 = append(v4, "")
                                    v3 += (^(v3))
                                    _ = fi24
                                }
                                if (!true) {
                                    v29 := GenS2{F0: ("\n\t" + "foo bar")}
                                    _ = v29
                                } else {
                                    v3 = v3
                                }
                                var v30 int32 = v3
                                _ = v30
                            }
                        } else {
                            v3 -= v3
                            if (0 != -507) {
                                var v31 int32 = v3
                                for fi32 := 0; fi32 < 7; fi32++ {
                                    if (false || (!true)) {
                                        v3 = (^(v3))
                                        switch fi32 {
                                        case 1:
                                            v31 = (v31 * (v31 - v31))
                                        case 42:
                                            var v33 int32 = v3
                                            _ = v33
                                        case 9223372036854775807:
                                            v3 -= 318
                                        default:
                                            fmt.Println(v3)
                                        }
                                        for fi34 := 0; fi34 < 6; fi34++ {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v35 float32 = ((0.0 * 1e10) * 0.1)
                                            _ = fi34
                                            _ = v35
                                        }
                                    } else {
                                        var v36 int = int(v3)
                                        var v37 float64 = ((1e10 - float64((float32((-2.25 + -2.25)) / 0.1))) * 3.14159)
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS2{F0: "a"}
                                    _ = fi32
                                    _ = v38
                                }
                                _ = v31
                            }
                            v3 = -105
                        }
                        switch -834 {
                        case 9223372036854775807:
                            for fi39 := 0; fi39 < 8; fi39++ {
                                v2[-2147483648] = 474
                                v3 -= v3
                                switch ("foo bar" + "hello") {
                                case "":
                                    v3 = 42
                                default:
                                    if true {
                                        if false {
                                            var v40 string = (("" + "foo bar") + "a")
                                            fmt.Println(v3, fi39)
                                            _ = v40
                                        } else {
                                            var v41 int = (fi39 / (((-(fi39))) | 1))
                                            _ = v41
                                        }
                                        v2[32768] = 348
                                    }
                                }
                                _ = fi39
                            }
                        case -2147483648:
                            var v42 int16 = 711
                            _ = v42
                        }
                    }
                case 127:
                    var v43 float32 = ((1.5 * float32(float64(0.0))) + 1e10)
                    _ = v43
                default:
                    switch -697 {
                    case -2147483648:
                        var v44 float64 = (1.5 + 3.14159)
                        _ = v44
                    default:
                        if ((2 < 918) && (!true)) {
                            for fi45 := 0; fi45 < 9; fi45++ {
                                _ = len(v1)
                                _ = fi45
                            }
                        }
                    }
                }
            } else {
                var v46 bool = ((v3 ^ v3) > v3)
                for fi47 := 0; fi47 < 17; fi47++ {
                    if (false || false) {
                        var v48 float64 = -2.25
                        fmt.Println(v3)
                        _ = len(v1)
                        _ = v48
                    }
                    _ = fi47
                }
                v49 := GenS4{F0: -1}
                _ = v46
                _ = v49
            }
        }
    }
    var v50 uint64 = 0
    v3 = v3
    v50 += v50
    if (!false) {
        for fi51 := 0; fi51 < 10; fi51++ {
            switch "foo bar" {
            case "":
                var v52 uint8 = 42
                _ = v52
            }
            _ = fi51
        }
    }
    if false {
        v4 = append(v4, "\n\t")
        if ((v3 > (-(v3))) || (789 >= 9223372036854775807)) {
            v50 = v50
            var v53 int64 = -9223372036854775808
            _ = v53
        }
        v2[1] = 395
    } else {
        v54 := GenS5{F0: -894, F1: 42, F2: genFn1(2, ((42 > 100) || (!true)))}
        switch 386 {
        case 933:
            if false {
                fmt.Println(v3)
                if (!false) {
                    if (true && false) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v55 := GenS1{F0: 42}
                            _ = v55
                        default:
                            v56 := GenS3{F0: 521}
                            _ = v56
                        }
                        if (1 < 1) {
                            for fi57 := 0; fi57 < 12; fi57++ {
                                for fi58 := 0; fi58 < 20; fi58++ {
                                    if len(v4) > 0 {
                                        v4[0] = ("a" + "a")
                                    }
                                    var v59 uint32 = 2147483647
                                    fmt.Println(fi58, fi58)
                                    var v60 uint16 = 32767
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                                switch "foo bar" {
                                case "hello":
                                    var v61 uint16 = 1
                                    _ = v61
                                case "":
                                    v50 = v50
                                }
                                fmt.Println(v3, v3)
                                _ = fi57
                            }
                            fmt.Println(v50, v3)
                            var v62 float32 = float32(0.1)
                            _ = v62
                        } else {
                            fmt.Println(v50, v50)
                            fmt.Println(v50, v3)
                        }
                    }
                    switch 620 {
                    case -9223372036854775808:
                        _ = len(v4)
                    case 1:
                        fmt.Println(v50, v50)
                    }
                }
            } else {
                for fi63 := 0; fi63 < 20; fi63++ {
                    switch "hello" {
                    case "a":
                        switch fi63 {
                        case 42:
                            if ("hello" > "foo bar") {
                                if true {
                                    v64 := GenS6{F0: 1, F1: 593, F2: "\n\t"}
                                    v50 += (v50 % ((v50) | 1))
                                    fmt.Println(v3, v3)
                                    _ = v64
                                }
                                if (0 < 61) {
                                    switch ((("foo bar" + "a") + "foo bar") + "hello") {
                                    case "\n\t":
                                        if false {
                                            for fi65 := 0; fi65 < 7; fi65++ {
                                                v50 = v50
                                                switch (fi65 << ((0) & 63)) {
                                                case -2147483648:
                                                    if true {
                                                        var v66 uint32 = 1
                                                        _ = v66
                                                    } else {
                                                        switch (fi65 ^ fi63) {
                                                        case 2147483647:
                                                            v3 = 1
                                                        case 176:
                                                            v4 = append(v4, ((("a" + "hello") + "foo bar") + "\n\t"))
                                                        case -1:
                                                            v67 := GenS7{F0: (v3 & (^(v3))), F1: 127}
                                                            _ = v67
                                                        default:
                                                            v4 = append(v4, "foo bar")
                                                        }
                                                        v68 := v2[fi63]
                                                        var v69 float32 = 0.0
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    v70 := GenS6{F0: 2, F1: 2147483647, F2: ""}
                                                    switch 32767 {
                                                    case -782:
                                                        v3 = v3
                                                    }
                                                    _ = v70
                                                case 42:
                                                    fmt.Println(fi65, v50)
                                                case 0:
                                                    var v71 uint32 = 297
                                                    _ = v71
                                                default:
                                                    v1[495] = true
                                                }
                                                _ = fi65
                                            }
                                            var v72 uint32 = uint32(v50)
                                            if ("" >= ("a" + ("a" + ((("" + ("\n\t" + "\n\t")) + "hello") + "foo bar")))) {
                                                v73 := GenS8{F0: 66, F1: fi63, F2: -36}
                                                _ = v73
                                            } else {
                                                v74 := v2[(fi63 - fi63)]
                                                for fi75 := 0; fi75 < 20; fi75++ {
                                                    if (((0.0 + 0.1) + float64(3.14159)) == float64(((((-2.25 - 0.1) * -2.25) * -2.25) / 1.0))) {
                                                        switch (("hello" + "\n\t") + ("hello" + ("foo bar" + ("" + "a")))) {
                                                        case "hello":
                                                            if ((true || ((v3 - (v3 % ((v3) | 1))) > v3)) && (false || (1 < 1))) {
                                                                v74 = 160
                                                            }
                                                        }
                                                        var v76 string = ""
                                                        _ = v76
                                                    } else {
                                                        v74 -= (v72 % ((v72) | 1))
                                                        if (true && ((v72 & v74) != (v72 | v74))) {
                                                            v1[(v72 - v74)] = true
                                                            switch (fi75 << ((4294967295) & 63)) {
                                                            case 944:
                                                                v74 = (^(v74))
                                                            case 711:
                                                                if false {
                                                                    v77 := GenS9{F0: uint(v50)}
                                                                    _ = v77
                                                                } else {
                                                                    v78 := GenS9{F0: 2}
                                                                    delete(v2, 849)
                                                                    v3 = v3
                                                                    _ = v78
                                                                }
                                                                var v79 uint64 = 0
                                                                _ = v79
                                                            case 9223372036854775807:
                                                                switch (-(fi75)) {
                                                                case -2147483648:
                                                                    fmt.Println(fi63)
                                                                    v3 += (-(v3))
                                                                case 4294967295:
                                                                    fmt.Println(v74)
                                                                }
                                                                if (-9223372036854775808 != 9223372036854775807) {
                                                                    v80 := GenS10{F0: fi63}
                                                                    _ = v80
                                                                }
                                                                v72 += 1
                                                            }
                                                        }
                                                        v72 = v72
                                                        v3 -= (v3 + -1)
                                                    }
                                                    v50 = (v50 ^ genFn1(255, true))
                                                    _ = len(v4)
                                                    _ = fi75
                                                }
                                                v81 := GenS11{F0: 40}
                                                _ = v74
                                                _ = v81
                                            }
                                            _ = v72
                                        }
                                    }
                                    if (329 == 920) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    for fi82 := 0; fi82 < 7; fi82++ {
                                        switch "foo bar" {
                                        case "hello":
                                            v3 += v3
                                        case "a":
                                            if (false || (-32768 > 42)) {
                                                var v83 int = fi82
                                                _ = v83
                                            }
                                        }
                                        var v84 uint = 2
                                        fmt.Println(fi82, fi82, fi82)
                                        _ = fi82
                                        _ = v84
                                    }
                                }
                                switch (fi63 << ((0) & 63)) {
                                case -380:
                                    var v85 uint8 = 128
                                    _ = v85
                                default:
                                    v3 = (v3 | 42)
                                }
                            }
                        }
                    case "\n\t":
                        _ = len(v4)
                    case "foo bar":
                        fmt.Println(v50)
                    }
                    if true {
                        var v86 int = fi63
                        for fi87 := 0; fi87 < 13; fi87++ {
                            var v88 float64 = -2.25
                            _ = fi87
                            _ = v88
                        }
                        switch 32767 {
                        case 128:
                            var v89 int = (v86 + (fi63 - (-(fi63))))
                            _ = v89
                        case 1:
                            v86 = v86
                        }
                        _ = v86
                    }
                    _ = fi63
                }
                if (true && (!false)) {
                    if ("" <= ("\n\t" + "")) {
                        if true {
                            fmt.Println(v3)
                            v90 := GenS4{F0: -1}
                            for fi91 := 0; fi91 < 10; fi91++ {
                                if ((^(v50)) > v50) {
                                    if (1 > (v3 ^ -2147483648)) {
                                        v92 := GenS6{F0: 42, F1: 42, F2: "a"}
                                        if true {
                                            for fi93 := 0; fi93 < 13; fi93++ {
                                                delete(v1, 2147483647)
                                                switch ("" + "foo bar") {
                                                case "":
                                                    var v94 int64 = 0
                                                    _ = v94
                                                default:
                                                    fmt.Println(v50, fi93, fi93)
                                                }
                                                v95 := GenS5{F0: fi93, F1: 0, F2: (v50 + v50)}
                                                _ = fi93
                                                _ = v95
                                            }
                                            v50 -= (v50 % ((v50) | 1))
                                        }
                                        _ = v92
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = ("foo bar" + "\n\t")
                                    }
                                }
                                _ = len(v4)
                                _ = fi91
                            }
                            _ = v90
                        } else {
                            fmt.Println(v50)
                            v50 = (v50 | genFn0(4294967295, 32768, -1))
                            var v96 int32 = v3
                            _ = v96
                        }
                    }
                }
            }
        case 906:
            var v97 int64 = -325
            _ = v97
        }
        _ = v54
    }
    for fi98 := 0; fi98 < 4; fi98++ {
        fmt.Println(v3, fi98)
        _ = fi98
    }
    var v99 int64 = int64(v50)
    v100 := GenS12{F0: 1, F1: 713, F2: 55}
    v4 = append(v4, "foo bar")
    v3 -= (v3 - (v3 & (v3 << ((0) & 63))))
    var v101 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v99
}
