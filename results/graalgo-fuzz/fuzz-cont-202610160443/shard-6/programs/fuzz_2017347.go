package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64; F1 uint; F2 bool }
func genFn0(p0 int16, p1 uint8) uint32 {
    var v5 bool = false
    var v6 uint32 = 65535
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (v6 >> ((1) & 63))
}
func genFn1() bool {
    var v7 uint = 18446744073709551615
    v8 := []int{}
    var v9 float32 = -2.25
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := []int{32767, 444, 32768, 988}
    v2 := []string{"\n\t", ("a" + (("hello" + "\n\t") + "")), "", "\n\t"}
    var v3 int8 = 0
    var v4 uint32 = 2147483647
    for fi10 := 0; fi10 < 12; fi10++ {
        if (-2.25 <= 0.0) {
            if ((fi10 / ((fi10) | 1)) == -2147483648) {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                switch "a" {
                case "a":
                    for fi11 := 0; fi11 < 4; fi11++ {
                        switch "" {
                        case "\n\t":
                            v2 = append(v2, "\n\t")
                        case "":
                            v1 = append(v1, fi11)
                        case "a":
                            v4 -= 42
                        default:
                            fmt.Println(fi10, v4)
                        }
                        if true {
                            if (!(-355 >= -1)) {
                                switch (fi11 ^ fi10) {
                                case 504:
                                    if true {
                                        switch -9223372036854775808 {
                                        case 42:
                                            var v12 int16 = -587
                                            _ = v12
                                        case -820:
                                            v3 = (v3 << ((18446744073709551615) & 63))
                                            var v13 uint8 = 0
                                            _ = v13
                                        case 60:
                                            v3 = (v3 + v3)
                                        }
                                        if (false && (21 < 97)) {
                                            var v14 int64 = 42
                                            v3 = v3
                                            _ = v14
                                        }
                                        v15 := GenS1{F0: 37}
                                        if (9223372036854775807 == (fi11 & int(v3))) {
                                            fmt.Println(v3)
                                            v16 := GenS1{F0: 127}
                                            _ = v16
                                        }
                                        var v17 uint16 = 408
                                        if false {
                                            if (1e10 == 3.14159) {
                                                fmt.Println(fi10)
                                                var v18 uint8 = 1
                                                _ = v18
                                            } else {
                                                v4 += genFn0(1, 131)
                                                var v19 uint16 = (v17 % (((v17 ^ v17)) | 1))
                                                var v20 bool = false
                                                _ = v19
                                                _ = v20
                                            }
                                        }
                                        _ = v15
                                        _ = v17
                                    }
                                }
                                if (42 == 442) {
                                    var v21 float32 = 0.0
                                    if true {
                                        _ = len(v1)
                                    } else {
                                        var v22 uint8 = 255
                                        if (!true) {
                                            v23 := GenS2{F0: 222, F1: (v4 & 0), F2: v3}
                                            v22 = v22
                                            v24 := GenS1{F0: v22}
                                            _ = v23
                                            _ = v24
                                        } else {
                                            var v25 int64 = 42
                                            if (false && false) {
                                                v21 = (1e10 - float32(-2.25))
                                                v22 = 2
                                                v22 += (v22 | 255)
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v26 uint32 = (v4 >> ((864) & 63))
                                            _ = v25
                                            _ = v26
                                        }
                                        v4 -= (v4 << ((917) & 63))
                                        v1 = append(v1, 621)
                                        var v27 int = (fi11 * (fi11 | -444))
                                        _ = v22
                                        _ = v27
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v21
                                }
                            }
                            v28 := GenS3{F0: 703}
                            _ = v28
                        } else {
                            _ = len(v1)
                        }
                        _ = fi11
                    }
                case "hello":
                    v3 = -1
                default:
                    for fi29 := 0; fi29 < 4; fi29++ {
                        v1 = append(v1, fi29)
                        v4 = v4
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi29
                    }
                }
                fmt.Println(v3, v3, fi10)
            }
            var v30 uint64 = 18446744073709551615
            if (((float64(-2.25) / 0.1) + -2.25) > 0.0) {
                var v31 bool = (42 > 1)
                _ = v31
            } else {
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 int = (fi32 * (-(fi10)))
                    v1 = append(v1, (fi32 ^ (^(fi32))))
                    if true {
                        var v34 uint32 = 58
                        v35 := GenS1{F0: 127}
                        _ = v34
                        _ = v35
                    }
                    if false {
                        if (!true) {
                            v36 := GenS2{F0: 220, F1: (v4 | v4), F2: v3}
                            switch (v33 % (((v33 << ((2) & 63))) | 1)) {
                            case 9223372036854775807:
                                switch "\n\t" {
                                case "":
                                    v37 := GenS4{F0: 42, F1: 94}
                                    _ = v37
                                }
                            default:
                                v30 += v30
                            }
                            _ = v36
                        } else {
                            v38 := GenS5{F0: v33, F1: fi10}
                            _ = v38
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v1 = append(v1, (^(fi32)))
                        v33 = (fi32 ^ (fi10 * (fi10 * v33)))
                    }
                    v3 -= 127
                    _ = fi32
                    _ = v33
                }
                for fi39 := 0; fi39 < 8; fi39++ {
                    if (true || (-382 > int32(v4))) {
                        switch 0 {
                        case 4294967295:
                            v40 := GenS6{F0: (fi10 & (fi39 | fi39))}
                            _ = v40
                        default:
                            v1 = append(v1, (fi10 << ((1) & 63)))
                        }
                        v41 := GenS7{F0: (false || (!(false || (!(((float32(0.1) / -2.25) >= 1.5) && false)))))}
                        var v42 int32 = 19
                        _ = v41
                        _ = v42
                    }
                    fmt.Println(v30)
                    _ = fi39
                }
            }
            _ = v30
        } else {
            var v43 float64 = ((0.1 + 3.14159) / 1.0)
            var v44 int64 = 1
            _ = v43
            _ = v44
        }
        v3 -= v3
        _ = fi10
    }
    switch 667 {
    case 4294967295:
        if true {
            if (float32((1e10 * (1.5 - (3.14159 / 1.5)))) > 0.1) {
                if false {
                    if (!true) {
                        for fi45 := 0; fi45 < 1; fi45++ {
                            fmt.Println(v4, v3, v3)
                            _ = fi45
                        }
                        for fi46 := 0; fi46 < 13; fi46++ {
                            if true {
                                var v47 float64 = (1e10 + -2.25)
                                var v48 float32 = (1.5 * (3.14159 - ((1.5 / (1.5 - 0.0)) * float32((((1.5 * float64(((float32(v47) - -2.25) - 1.5))) * float64((0.0 - float32(v47)))) + float64((((-2.25 + 1e10) * -2.25) / 1e10)))))))
                                _ = v47
                                _ = v48
                            }
                            _ = fi46
                        }
                        v1 = append(v1, -104)
                    }
                    if (42 != 42) {
                        v2 = append(v2, (((("foo bar" + ("a" + "\n\t")) + "hello") + ("a" + "foo bar")) + ""))
                    }
                }
                var v49 int64 = -620
                _ = v49
            } else {
                switch (("a" + "\n\t") + ("" + "a")) {
                case "a":
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                default:
                    v50 := GenS4{F0: -2147483648, F1: int64(v4)}
                    _ = v50
                }
                if true {
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            v4 = (v4 << ((0) & 63))
                        case "\n\t":
                            if true {
                                var v51 bool = (true || (255 > 2))
                                var v52 bool = (!(false || v51))
                                v53 := GenS8{F0: 748, F1: 151, F2: 42}
                                _ = v51
                                _ = v52
                                _ = v53
                            } else {
                                switch 9223372036854775807 {
                                case -255:
                                    var v54 uint = 0
                                    _ = v54
                                case -1:
                                    fmt.Println(v4)
                                case -419:
                                    switch -563 {
                                    case 2147483647:
                                        var v55 int64 = 559
                                        _ = v55
                                    case 32767:
                                        v4 = (v4 % (((^(v4))) | 1))
                                    case -2147483648:
                                        v4 = (v4 >> ((255) & 63))
                                    default:
                                        v1 = append(v1, 0)
                                    }
                                }
                                v56 := GenS9{F0: 591, F1: 1000}
                                _ = v56
                            }
                        case "a":
                            for fi57 := 0; fi57 < 9; fi57++ {
                                if len(v1) > 0 {
                                    v1[0] = fi57
                                }
                                switch (fi57 + (fi57 / (((fi57 | (fi57 / (((fi57 * fi57)) | 1)))) | 1))) {
                                case 0:
                                    v58 := GenS10{F0: 32767}
                                    _ = v58
                                case 2147483647:
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                case 128:
                                    for fi59 := 0; fi59 < 6; fi59++ {
                                        _ = len(v2)
                                        var v60 float32 = ((0.1 + 1e10) + ((0.1 - 1e10) + 0.1))
                                        _ = fi59
                                        _ = v60
                                    }
                                default:
                                    if (-2.25 == (0.0 * 3.14159)) {
                                        var v61 uint = 4294967295
                                        _ = v61
                                    }
                                }
                                var v62 uint32 = v4
                                _ = fi57
                                _ = v62
                            }
                        }
                        for fi63 := 0; fi63 < 4; fi63++ {
                            v3 += v3
                            _ = fi63
                        }
                    }
                }
            }
        }
    case 0:
        v1 = append(v1, -773)
    }
    switch "a" {
    case "":
        v4 = (v4 ^ genFn0(42, 2))
    case "\n\t":
        if true {
            if true {
                v3 -= (-(v3))
                fmt.Println(v3, v4)
                var v64 int8 = (v3 + (-(v3)))
                _ = v64
            }
            v4 = (v4 % ((v4) | 1))
            var v65 int = 127
            _ = v65
        }
    }
    v66 := GenS4{F0: 42, F1: 897}
    var v67 uint8 = 255
    var v68 uint16 = uint16(v67)
    var v69 bool = true
    for fi70 := 0; fi70 < 20; fi70++ {
        fmt.Println(v3, v68)
        _ = fi70
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v71 uint = 65535
    v72 := GenS5{F0: 4294967295, F1: -10}
    var v73 float32 = float32(3.14159)
    if (-2.25 <= float64(1e10)) {
        v71 += (v71 - v71)
        var v74 string = ("\n\t" + "\n\t")
        _ = v74
    } else {
        for fi75 := 0; fi75 < 17; fi75++ {
            var v76 int32 = 0
            fmt.Println(v71, fi75)
            _ = fi75
            _ = v76
        }
        fmt.Println(v67, v3)
        _ = len(v1)
    }
    v3 += (-(v3))
    fmt.Println(v69, v68)
    v67 += (v67 & v67)
    if v69 {
        var v77 int32 = 1
        fmt.Println(v4)
        _ = v77
    } else {
        for fi78 := 0; fi78 < 14; fi78++ {
            switch (fi78 | (fi78 % (((fi78 % ((fi78) | 1))) | 1))) {
            case -1:
                for fi79 := 0; fi79 < 9; fi79++ {
                    v1 = append(v1, fi79)
                    var v80 int = (^(fi79))
                    var v81 int8 = v3
                    v69 = (false || true)
                    _ = fi79
                    _ = v80
                    _ = v81
                }
            default:
                if v69 {
                    switch ("\n\t" + ("\n\t" + ((("hello" + "hello") + "a") + "a"))) {
                    case "":
                        fmt.Println(v71, v3)
                    case "a":
                        switch ((("" + ("" + "hello")) + "") + "foo bar") {
                        case "":
                            var v82 uint16 = 699
                            _ = v82
                        case "foo bar":
                            _ = len(v1)
                        }
                    case "\n\t":
                        switch "foo bar" {
                        case "hello":
                            var v83 int64 = 43
                            _ = v83
                        case "\n\t":
                            switch (-(fi78)) {
                            case 2147483647:
                                v84 := GenS11{F0: 4294967295, F1: (v71 | (v71 << ((v71) & 63))), F2: v69}
                                _ = v84
                            }
                        default:
                            v71 += (v71 * v71)
                        }
                    default:
                        for fi85 := 0; fi85 < 11; fi85++ {
                            var v86 uint16 = (^(v68))
                            var v87 bool = false
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                    }
                    fmt.Println(fi78, v69, v68)
                    var v88 uint16 = v68
                    _ = v88
                } else {
                    var v89 bool = true
                    _ = v89
                }
            }
            v90 := GenS11{F0: 42, F1: uint(v68), F2: v69}
            v73 -= 0.1
            _ = fi78
            _ = v90
        }
        var v91 uint16 = v68
        _ = v91
    }
    if v69 {
        fmt.Println(v68)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v71
    _ = v72
    _ = v73
}
