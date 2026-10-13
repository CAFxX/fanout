package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool; F1 string; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int; F1 int8; F2 int32 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int; F1 uint32; F2 bool }
type GenS11 struct { F0 int }
func genFn0(p0 uint64) uint64 {
    var v5 uint64 = (p0 >> ((2) & 63))
    if (p0 >= 1) {
        v6 := []int{-1, -9223372036854775808, 42, 32768}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            v5 = 49
            switch ("foo bar" + ("" + "a")) {
            case "":
                fmt.Println(v5)
            }
            v8 := []int{}
            _ = fi7
            _ = v8
        }
    }
    v9 := []int{4294967295, 307, 1, 674}
    _ = p0
    _ = v5
    _ = v9
    return p0
}
func genFn1(p0 uint64, p1 int, p2 int32) int16 {
    switch "hello" {
    case "a":
        var v10 uint64 = (p0 << ((255) & 63))
        _ = v10
    case "\n\t":
        if (!(true && false)) {
            if true {
                v11 := []int{-1}
                _ = v11
            } else {
                if (false || false) {
                    p0 += (p0 + (p0 * genFn0((^(p0)))))
                } else {
                    p1 = 127
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 int32 = p2
                        v14 := []int{}
                        for fi15 := 0; fi15 < 15; fi15++ {
                            if (true && (0.1 < (-2.25 / -2.25))) {
                                if false {
                                    switch "foo bar" {
                                    case "a":
                                        var v16 uint = 65535
                                        _ = v16
                                    case "\n\t":
                                        v17 := GenS1{F0: 9223372036854775807, F1: genFn0((p0 >> ((856) & 63))), F2: 127}
                                        _ = v17
                                    case "":
                                        var v18 int = -2147483648
                                        _ = v18
                                    default:
                                        if (123 >= 0) {
                                            if (!false) {
                                                var v19 uint16 = 2
                                                if false {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v13 = (p2 << ((818) & 63))
                                                    case "hello":
                                                        v20 := GenS2{F0: false, F1: "\n\t", F2: (fi15 * fi12)}
                                                        _ = v20
                                                    case "":
                                                        var v21 float32 = 1.5
                                                        _ = v21
                                                    }
                                                    if len(v14) > 0 {
                                                        v14[0] = p1
                                                    }
                                                    v22 := GenS3{F0: (fi12 * -893)}
                                                    _ = v22
                                                }
                                                _ = len(v14)
                                                _ = v19
                                            } else {
                                                p0 -= 1
                                                var v23 uint8 = 231
                                                _ = v23
                                            }
                                        } else {
                                            switch (((("a" + "\n\t") + ("hello" + "\n\t")) + "\n\t") + "") {
                                            case "\n\t":
                                                if true {
                                                    var v24 uint = 255
                                                    v24 = 255
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                    p1 += fi15
                                                    _ = v24
                                                } else {
                                                    switch (("hello" + "") + ((("foo bar" + ("" + "a")) + "foo bar") + ("foo bar" + "a"))) {
                                                    case "a":
                                                        if len(v14) >= 2 {
                                                            v14 = v14[1:len(v14)]
                                                        }
                                                    }
                                                    p1 -= (p1 + (fi12 + 128))
                                                    var v25 int16 = 42
                                                    _ = v25
                                                }
                                            }
                                            v26 := GenS3{F0: (fi15 ^ (fi12 / ((32767) | 1)))}
                                            var v27 int64 = 9223372036854775807
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                }
                                v13 -= p2
                                if (!(!("foo bar" <= "\n\t"))) {
                                    var v28 float32 = float32(1.5)
                                    v29 := GenS4{F0: 214}
                                    var v30 bool = ("\n\t" > (((("foo bar" + "a") + "hello") + ("a" + "foo bar")) + ""))
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(fi12, v13, p1)
                                        _ = len(v14)
                                        v31 := GenS4{F0: 1}
                                        _ = v31
                                    default:
                                        v32 := GenS2{F0: ((p1 % ((32768) | 1)) != (fi12 / (((p1 / ((fi15) | 1))) | 1))), F1: (("a" + ("foo bar" + ("\n\t" + "hello"))) + "foo bar"), F2: -757}
                                        _ = v32
                                    }
                                    v14 = append(v14, fi15)
                                    fmt.Println(fi15, p2)
                                }
                            } else {
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                v33 := GenS5{F0: (v13 ^ v13)}
                                if true {
                                    p2 -= p2
                                    fmt.Println(p1)
                                    if (!false) {
                                        switch 32767 {
                                        case -1:
                                            v34 := GenS5{F0: v13}
                                            _ = v34
                                        default:
                                            fmt.Println(fi12, p1)
                                        }
                                        var v35 int64 = -1
                                        if false {
                                            if false {
                                                var v36 int32 = (v13 % ((v13) | 1))
                                                var v37 float32 = (3.14159 - 0.1)
                                                _ = v36
                                                _ = v37
                                            } else {
                                                var v38 bool = (0 != 0)
                                                var v39 uint = 833
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                        _ = v35
                                    } else {
                                        fmt.Println(p0)
                                        if (fi12 < fi15) {
                                            var v40 int64 = 42
                                            if (42 >= 4294967295) {
                                                switch fi12 {
                                                case -9223372036854775808:
                                                    v13 -= v13
                                                }
                                            } else {
                                                v13 -= (p2 * p2)
                                                fmt.Println(v13, v13)
                                                if false {
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                    var v41 int64 = (^(v40))
                                                    p2 = (p2 % ((p2) | 1))
                                                    _ = v41
                                                }
                                                if (!false) {
                                                    var v42 uint64 = 4294967295
                                                    p0 += p0
                                                    p2 += v13
                                                    _ = v42
                                                }
                                            }
                                            _ = v40
                                        }
                                        if (uint8(p0) <= 0) {
                                            p1 += 2147483647
                                        }
                                        v43 := GenS3{F0: -9223372036854775808}
                                        _ = v43
                                    }
                                } else {
                                    var v44 int = (fi15 ^ -2147483648)
                                    v14 = append(v14, 128)
                                    v44 = 128
                                    v14 = append(v14, p1)
                                    p0 = p0
                                    switch "foo bar" {
                                    case "foo bar":
                                        v45 := GenS6{F0: 127, F1: -48}
                                        _ = v45
                                    case "":
                                        if false {
                                            var v46 int8 = -25
                                            var v47 uint16 = 2
                                            v47 -= 255
                                            _ = v46
                                            _ = v47
                                        }
                                    default:
                                        v48 := GenS7{F0: 2147483647, F1: 42, F2: 0}
                                        _ = v48
                                    }
                                    v49 := GenS3{F0: (-(p1))}
                                    _ = v44
                                    _ = v49
                                }
                                _ = v33
                            }
                            _ = fi15
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    if false {
                        p2 = (^(p2))
                    }
                }
            }
        } else {
            p1 += p1
            fmt.Println(p2, p1)
            v50 := []int{p1, (p1 | (p1 % (((p1 >> ((2) & 63))) | 1))), (p1 | (p1 ^ (p1 - p1))), -949}
            _ = v50
        }
    }
    v51 := GenS2{F0: (1 <= 32767), F1: "foo bar", F2: 42}
    if false {
        switch "\n\t" {
        case "\n\t":
            var v52 int16 = 32767
            _ = v52
        case "hello":
            fmt.Println(p2, p0)
        }
    } else {
        v53 := []int{(p1 >> ((332) & 63)), (p1 % ((p1) | 1)), p1, p1}
        fmt.Println(p0)
        _ = v53
    }
    if true {
        switch "hello" {
        case "":
            fmt.Println(p2, p0)
        case "hello":
            var v54 uint = 0
            _ = v54
        case "foo bar":
            for fi55 := 0; fi55 < 3; fi55++ {
                switch (p1 | fi55) {
                case 127:
                    p1 = fi55
                case 887:
                    var v56 int = (p1 >> ((1) & 63))
                    _ = v56
                case 9223372036854775807:
                    v57 := GenS7{F0: (-(p1)), F1: 42, F2: 0}
                    _ = v57
                }
                switch 32768 {
                case -52:
                    switch (((("" + "a") + ("hello" + "\n\t")) + "hello") + "") {
                    case "":
                        var v58 int8 = 35
                        _ = v58
                    case "hello":
                        if (4294967295 != 18446744073709551615) {
                            if true {
                                switch int(p2) {
                                case 4294967295:
                                    if ((p2 & (p2 >> ((677) & 63))) > (p2 * -2147483648)) {
                                        var v59 float64 = (-2.25 / (1e10 + 1.5))
                                        var v60 uint8 = 1
                                        _ = v59
                                        _ = v60
                                    }
                                case 0:
                                    p1 = (-(p1))
                                default:
                                    fmt.Println(p2)
                                }
                            }
                            v61 := GenS8{F0: -664, F1: 127}
                            v62 := []int{(^(p1)), (p1 / ((128) | 1)), p1}
                            _ = v61
                            _ = v62
                        } else {
                            v63 := GenS8{F0: 0, F1: 69}
                            p2 -= (p2 ^ 42)
                            switch "\n\t" {
                            case "foo bar":
                                v64 := []int{fi55, (p1 + p1), p1}
                                _ = v64
                            default:
                                v65 := GenS6{F0: 1, F1: 127}
                                _ = v65
                            }
                            _ = v63
                        }
                    default:
                        var v66 int16 = -734
                        _ = v66
                    }
                }
                for fi67 := 0; fi67 < 19; fi67++ {
                    switch -1 {
                    case 9223372036854775807:
                        p1 = 0
                    case -972:
                        v68 := []int{}
                        _ = v68
                    case -2147483648:
                        p0 = (^(p0))
                    }
                    var v69 int64 = 0
                    fmt.Println(v69)
                    _ = fi67
                    _ = v69
                }
                _ = fi55
            }
        }
        p0 = (p0 | 1)
    }
    if true {
        switch ("hello" + "") {
        case "\n\t":
            var v70 float32 = float32((-2.25 + ((((float64((((0.1 - (((-2.25 + (3.14159 / 1.0)) / (((1.5 / (3.14159 + ((1.5 + 0.0) / 1.0))) + 0.0) - 1.5)) / 3.14159)) / 0.1) + 3.14159)) + 1e10) * 0.0) / (3.14159 * (0.1 * 0.1))) - 1e10)))
            _ = v70
        default:
            switch "\n\t" {
            case "a":
                fmt.Println(p0, p2, p0)
            case "":
                if (false && true) {
                    switch (("foo bar" + "") + "") {
                    case "hello":
                        v71 := GenS8{F0: -1, F1: 42}
                        _ = v71
                    case "a":
                        p0 = (p0 >> ((uint(p0)) & 63))
                    }
                }
            case "\n\t":
                v72 := GenS9{F0: "a"}
                _ = v72
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v51
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{688, 551}
    v3 := []string{}
    var v4 int = -2147483648
    for fi73 := 0; fi73 < 20; fi73++ {
        fmt.Println(fi73, fi73)
        if true {
            var v74 int16 = 1000
            _ = v74
        } else {
            fmt.Println(fi73, fi73)
            if (!false) {
                for fi75 := 0; fi75 < 15; fi75++ {
                    fmt.Println(fi75, fi75, fi73)
                    v4 -= -911
                    v3 = append(v3, "")
                    var v76 string = ("" + "\n\t")
                    var v77 int16 = 0
                    switch (v76 + ((v76 + v76) + v76)) {
                    case "hello":
                        v76 = (v76 + v76)
                    case "foo bar":
                        var v78 uint64 = genFn0(831)
                        _ = v78
                    default:
                        v77 -= (v77 % (((v77 % ((genFn1(301, 9223372036854775807, 1)) | 1))) | 1))
                    }
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                _ = len(v1)
            }
        }
        _ = fi73
    }
    if (1e10 >= float32((1.5 - 3.14159))) {
        if false {
            var v79 int64 = -9223372036854775808
            _ = v79
        }
    }
    v80 := GenS5{F0: -2147483648}
    fmt.Println(v4)
    if true {
        _ = len(v1)
        v81 := v1[2147483648]
        _ = v81
    }
    switch "hello" {
    case "a":
        v4 -= v4
    default:
        var v82 int8 = 77
        _ = v82
    }
    fmt.Println(v4)
    var v83 uint = 4294967295
    v84 := GenS10{F0: (v4 ^ v4), F1: 4294967295, F2: (true || false)}
    var v85 int32 = -1
    if true {
        switch v4 {
        case 127:
            for fi86 := 0; fi86 < 11; fi86++ {
                switch v4 {
                case -2147483648:
                    var v87 uint16 = 134
                    _ = v87
                case 128:
                    _ = len(v1)
                case 607:
                    v85 = v85
                }
                v88 := GenS7{F0: (v4 >> (((^(v83))) & 63)), F1: 0, F2: v85}
                _ = fi86
                _ = v88
            }
        case 0:
            v89 := GenS11{F0: v4}
            _ = v89
        default:
            for fi90 := 0; fi90 < 17; fi90++ {
                for fi91 := 0; fi91 < 2; fi91++ {
                    v85 = (v85 / (((-(v85))) | 1))
                    _ = fi91
                }
                var v92 uint8 = 66
                fmt.Println(v4, fi90, fi90)
                _ = fi90
                _ = v92
            }
        }
        v93 := GenS11{F0: v4}
        fmt.Println(v83, v4, v85)
        _ = v93
    } else {
        var v94 float64 = 0.1
        _ = v94
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v80
    _ = v83
    _ = v84
    _ = v85
}
