package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint64; F2 int }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 int8; F2 string }
type GenS11 struct { F0 int16 }
func genFn0() bool {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 bool = ((0.1 > (float64((3.14159 * 1e10)) + 3.14159)) || false)
        _ = fi5
        _ = v6
    }
    v7 := []int{637, -283, -1, 0}
    v8 := GenS1{F0: -251, F1: 9223372036854775807, F2: 42}
    var v9 int16 = 32767
    v10 := GenS1{F0: (v9 << ((65535) & 63)), F1: 235, F2: 42}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 uint8, p1 bool) int8 {
    v11 := GenS2{F0: -287}
    switch (("foo bar" + ((("a" + ("hello" + "\n\t")) + ((("foo bar" + (("a" + "\n\t") + "foo bar")) + ("" + "hello")) + "")) + ("\n\t" + ""))) + ("a" + (("" + (("hello" + ("\n\t" + ("a" + "a"))) + ("\n\t" + "hello"))) + ""))) {
    case "a":
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS2{F0: 42}
                _ = fi13
                _ = v14
            }
            if (42 != 1000) {
                var v15 int16 = 998
                v15 -= 637
                _ = v15
            }
            var v16 uint16 = 65535
            _ = fi12
            _ = v16
        }
    case "\n\t":
        switch "\n\t" {
        case "hello":
            var v17 uint = 517
            _ = v17
        case "":
            if p1 {
                fmt.Println(p1, p0)
                var v18 uint = 207
                _ = v18
            }
        }
    }
    v19 := GenS1{F0: 1, F1: 0, F2: 1}
    if p1 {
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                p1 = p1
                var v22 uint32 = 4294967295
                if (65535 != 42) {
                    p1 = p1
                    if ((((-2.25 * -2.25) * 1e10) + 1.5) == 3.14159) {
                        v22 = (^(v22))
                        v23 := []int{fi21, (fi21 & 0), fi21, 4294967295}
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        _ = v23
                    } else {
                        var v24 bool = p1
                        v25 := []int{(fi20 / ((1) | 1)), (fi20 | (fi20 ^ fi21)), fi21}
                        if v24 {
                            if len(v25) > 0 {
                                v25[0] = 32768
                            }
                        } else {
                            v22 -= 0
                        }
                        _ = v24
                        _ = v25
                    }
                } else {
                    v26 := []int{}
                    _ = v26
                }
                switch "" {
                case "hello":
                    var v27 bool = p1
                    _ = v27
                default:
                    switch (fi21 | fi21) {
                    case 128:
                        p0 += (p0 << ((0) & 63))
                    case 32768:
                        fmt.Println(fi20)
                        fmt.Println(p1, fi20)
                    default:
                        v28 := []int{}
                        fmt.Println(fi20, p0, v22)
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                        _ = v28
                    }
                }
                v29 := []int{(-(fi21)), 32768, fi21}
                _ = fi21
                _ = v22
                _ = v29
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                fmt.Println(p0, p0)
                _ = fi30
            }
            p1 = ((false && p1) && (!p1))
            _ = fi20
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v11
    _ = v19
    return -1
}

func main() {
    v1 := []int{}
    var v2 uint64 = 18446744073709551615
    var v3 uint = 42
    var v4 uint16 = 42
    v1 = append(v1, 9223372036854775807)
    if len(v1) > 0 {
        v1[0] = -489
    }
    if true {
        fmt.Println(v2, v3, v3)
    } else {
        fmt.Println(v4, v4, v4)
        var v31 uint64 = 635
        var v32 uint32 = 0
        _ = v31
        _ = v32
    }
    v1 = append(v1, -829)
    v33 := GenS1{F0: -1, F1: 0, F2: 981}
    if (65535 > 922) {
        if (63 < 231) {
            var v34 string = "hello"
            switch v34 {
            case "\n\t":
                _ = len(v1)
            case "foo bar":
                fmt.Println(v4, v4, v4)
            }
            for fi35 := 0; fi35 < 15; fi35++ {
                for fi36 := 0; fi36 < 5; fi36++ {
                    fmt.Println(v34)
                    _ = fi36
                }
                _ = fi35
            }
            _ = v34
        }
        v4 += v4
    } else {
        var v37 int16 = -416
        _ = v37
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        v39 := GenS1{F0: 494, F1: -647, F2: v4}
        _ = len(v1)
        for fi40 := 0; fi40 < 4; fi40++ {
            v4 = 1
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        v2 = v2
        _ = fi41
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if (110 == 583) {
        v3 = 65535
    } else {
        if (false || ((v2 % ((0) | 1)) == v2)) {
            switch int(v2) {
            case 1:
                if true {
                    var v42 uint64 = v2
                    _ = v42
                }
            }
            if true {
                v1 = append(v1, 4294967295)
            }
            for fi43 := 0; fi43 < 16; fi43++ {
                fmt.Println(v4)
                v44 := GenS3{F0: (v4 + v4), F1: 378, F2: v3}
                fmt.Println(v4, v2)
                _ = fi43
                _ = v44
            }
        }
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        switch "a" {
        case "\n\t":
            for fi46 := 0; fi46 < 14; fi46++ {
                if true {
                    switch ("" + ("" + "hello")) {
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "\n\t":
                        v47 := GenS1{F0: int16(v2), F1: -1, F2: v4}
                        _ = v47
                    case "foo bar":
                        switch fi46 {
                        case 4294967295:
                            v3 += v3
                        }
                    default:
                        var v48 int64 = 42
                        _ = v48
                    }
                    fmt.Println(v3)
                    v2 = v2
                } else {
                    if false {
                        fmt.Println(fi46)
                    } else {
                        switch "a" {
                        case "hello":
                            fmt.Println(fi46, v4)
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "foo bar":
                            v4 -= v4
                        default:
                            var v49 int64 = 9223372036854775807
                            var v50 int8 = 127
                            v2 -= (v2 % ((v2) | 1))
                            _ = v49
                            _ = v50
                        }
                        v3 += (v3 << ((uint(v4)) & 63))
                        var v51 uint32 = 0
                        if false {
                            switch "a" {
                            case "hello":
                                switch "\n\t" {
                                case "\n\t":
                                    var v52 int = 128
                                    _ = v52
                                case "":
                                    if false {
                                        v1 = append(v1, fi45)
                                    } else {
                                        var v53 bool = false
                                        var v54 uint32 = uint32(v3)
                                        _ = v53
                                        _ = v54
                                    }
                                }
                            default:
                                switch fi45 {
                                case -9223372036854775808:
                                    v2 += (v2 - uint64(v4))
                                case 128:
                                    v51 += (^(v51))
                                case 9223372036854775807:
                                    v3 = 255
                                default:
                                    v4 = v4
                                }
                            }
                            switch "hello" {
                            case "hello":
                                if true {
                                    switch -1 {
                                    case -1:
                                        fmt.Println(v51)
                                    case 32768:
                                        var v55 int16 = 43
                                        v56 := GenS4{F0: fi46}
                                        _ = v55
                                        _ = v56
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v57 float32 = 3.14159
                                            _ = v57
                                        }
                                    }
                                    v51 = 568
                                    switch fi45 {
                                    case 831:
                                        v1 = append(v1, 2147483647)
                                    default:
                                        switch 934 {
                                        case -7:
                                            fmt.Println(v2, fi46)
                                        case 9223372036854775807:
                                            v4 += uint16(v2)
                                        case 42:
                                            switch (fi45 ^ 32768) {
                                            case 168:
                                                v58 := GenS1{F0: 0, F1: 544, F2: (v4 ^ 255)}
                                                _ = v58
                                            case -557:
                                                var v59 uint8 = uint8(v51)
                                                _ = v59
                                            }
                                        default:
                                            fmt.Println(v4, v2)
                                        }
                                    }
                                    v2 += v2
                                }
                            case "a":
                                switch (fi45 | fi45) {
                                case -9223372036854775808:
                                    if (!true) {
                                        v60 := GenS5{F0: fi45}
                                        _ = v60
                                    } else {
                                        fmt.Println(v4, v51)
                                        if (-2147483648 < -1) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v61 := GenS6{F0: 496, F1: v2, F2: (fi45 >> ((v3) & 63))}
                                            var v62 int16 = -809
                                            _ = v61
                                            _ = v62
                                        } else {
                                            v3 = (v3 ^ v3)
                                            v4 -= v4
                                        }
                                    }
                                case 4294967295:
                                    if ((0.1 - (1.5 - 0.1)) >= 0.1) {
                                        var v63 int32 = 2147483647
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v64 := GenS7{F0: v63, F1: 255}
                                        _ = v63
                                        _ = v64
                                    } else {
                                        if false {
                                            v65 := GenS7{F0: 0, F1: 26}
                                            _ = v65
                                        }
                                        v66 := GenS4{F0: fi46}
                                        v67 := GenS8{F0: (("foo bar" + "") + "a")}
                                        _ = v66
                                        _ = v67
                                    }
                                default:
                                    if false {
                                        var v68 int8 = 1
                                        switch "hello" {
                                        case "":
                                            v4 += (v4 % ((v4) | 1))
                                            if (!false) {
                                                if (!false) {
                                                    var v69 int8 = v68
                                                    if true {
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v70 uint32 = 1
                                                            _ = v70
                                                        case "hello":
                                                            if true {
                                                                fmt.Println(v4, v69, v2)
                                                                v68 -= v69
                                                                var v71 string = "\n\t"
                                                                _ = v71
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 161
                                                                }
                                                                v2 -= uint64(v51)
                                                            }
                                                        default:
                                                            if ((1 != 0) && ((!true) && false)) {
                                                                if (false || false) {
                                                                    if (v4 != v4) {
                                                                        if (3.14159 == (-2.25 + 3.14159)) {
                                                                            fmt.Println(v4, fi46)
                                                                        }
                                                                        v4 += v4
                                                                    } else {
                                                                        switch int(v69) {
                                                                        case 4294967295:
                                                                            v51 = (v51 % ((464) | 1))
                                                                        case 42:
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                if true {
                                                                                    v69 = genFn1(0, false)
                                                                                    v51 += 647
                                                                                    fmt.Println(v69, v68, v51)
                                                                                }
                                                                            case "a":
                                                                                if false {
                                                                                    v4 = (v4 % ((v4) | 1))
                                                                                }
                                                                            case "hello":
                                                                                if (-1 <= -2147483648) {
                                                                                    v51 = (v51 | v51)
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    v68 -= int8(v51)
                                                                                    fmt.Println(v69, fi45)
                                                                                    v68 -= v68
                                                                                    fmt.Println(v69, fi45)
                                                                                    var v72 int32 = -1
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                        case 472:
                                                                            v69 += (v68 * v68)
                                                                        }
                                                                        var v73 int64 = -790
                                                                        v1 = append(v1, fi45)
                                                                        _ = v73
                                                                    }
                                                                    v68 -= genFn1(255, false)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (fi46 >> (((v3 >> (((v3 ^ (v3 / ((v3) | 1)))) & 63))) & 63))
                                                                    }
                                                                }
                                                                v74 := GenS9{F0: v3}
                                                                _ = v74
                                                            }
                                                        }
                                                        if (("\n\t" > ("foo bar" + "foo bar")) && (true && (42 > -169))) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = len(v1)
                                                        }
                                                    } else {
                                                        v68 -= v68
                                                    }
                                                    v1 = append(v1, fi45)
                                                    _ = v69
                                                } else {
                                                    v75 := GenS8{F0: "\n\t"}
                                                    _ = v75
                                                }
                                                var v76 bool = true
                                                _ = v76
                                            }
                                        case "foo bar":
                                            fmt.Println(v68, v51)
                                        default:
                                            switch fi46 {
                                            case 128:
                                                var v77 uint32 = v51
                                                _ = v77
                                            default:
                                                if (1e10 <= 1.5) {
                                                    v1 = append(v1, (-(fi46)))
                                                    v2 += v2
                                                    var v78 bool = true
                                                    _ = v78
                                                }
                                            }
                                        }
                                        _ = v68
                                    } else {
                                        var v79 int = (fi45 << ((1) & 63))
                                        if (true && ((fi45 > v79) || false)) {
                                            fmt.Println(v79, v4, v79)
                                            if len(v1) > 0 {
                                                v1[0] = int(v51)
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                var v80 float32 = 1e10
                                                _ = v80
                                            }
                                        }
                                        _ = v79
                                    }
                                }
                            default:
                                var v81 int64 = 0
                                _ = v81
                            }
                            fmt.Println(v2, v4, fi46)
                        } else {
                            v82 := GenS8{F0: ("\n\t" + "")}
                            if len(v1) > 0 {
                                v1[0] = (-(fi45))
                            }
                            v83 := GenS7{F0: -1, F1: 220}
                            _ = v82
                            _ = v83
                        }
                        var v84 int16 = -219
                        _ = v51
                        _ = v84
                    }
                    switch -2147483648 {
                    case -2147483648:
                        switch (^(fi46)) {
                        case -2147483648:
                            v2 = uint64(v3)
                        case -409:
                            var v85 int16 = 32767
                            _ = v85
                        case 1:
                            var v86 float32 = (0.0 - (float32(((1.5 + 3.14159) / (0.0 + 1.5))) - float32(float64((0.0 - ((0.0 / -2.25) / 1e10))))))
                            _ = v86
                        }
                    case 32768:
                        var v87 float64 = 0.0
                        _ = v87
                    }
                }
                switch "foo bar" {
                case "\n\t":
                    v1 = append(v1, (fi45 % ((fi46) | 1)))
                case "hello":
                    v88 := GenS10{F0: (!true), F1: -1, F2: "\n\t"}
                    _ = v88
                }
                if true {
                    v2 = v2
                    var v89 int32 = int32(v4)
                    var v90 float32 = 0.0
                    _ = v89
                    _ = v90
                } else {
                    v2 = v2
                    v91 := GenS11{F0: 1}
                    _ = v91
                }
                _ = fi46
            }
        case "":
            var v92 uint16 = uint16(v3)
            _ = v92
        }
        _ = fi45
    }
    var v93 uint8 = 2
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v93
}
