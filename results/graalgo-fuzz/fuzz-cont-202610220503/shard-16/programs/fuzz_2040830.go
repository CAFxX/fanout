package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint8) uint32 {
    if true {
        v7 := GenS1{F0: 1}
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: (fi8 << ((255) & 63))}
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        v10 := GenS1{F0: 741}
        p0 -= 102
        _ = v10
    }
    if (-657 < 0) {
        if (!true) {
            p0 -= (p0 & 0)
        }
    } else {
        if (!true) {
            var v11 uint64 = 18446744073709551615
            p0 = (^(p0))
            _ = v11
        } else {
            switch 276 {
            case 0:
                p0 = p0
            case 9223372036854775807:
                v12 := []int{-1}
                _ = v12
            case 831:
                p0 = p0
            }
            v13 := []int{32768, 4294967295}
            _ = v13
        }
        fmt.Println(p0)
    }
    fmt.Println(p0)
    if (((!false) && (false && (-128 >= -128))) && (false && (1 <= -352))) {
        v14 := []int{}
        if false {
            if false {
                for fi15 := 0; fi15 < 13; fi15++ {
                    if true {
                        if true {
                            if ((fi15 ^ -1) <= (fi15 ^ (^(fi15)))) {
                                fmt.Println(p0, p0)
                            }
                            var v16 int8 = 46
                            p0 += p0
                            _ = v16
                        } else {
                            var v17 float32 = (3.14159 + -2.25)
                            _ = v17
                        }
                        p0 -= p0
                    }
                    p0 = p0
                    _ = fi15
                }
                var v18 bool = ((p0 * (p0 << ((565) & 63))) == (^(p0)))
                p0 -= 1
                _ = v18
            } else {
                fmt.Println(p0)
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 float32 = float32(1e10)
                    var v21 int64 = 42
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                p0 += 2
                _ = fi22
            }
        }
        v23 := GenS2{F0: 42, F1: 9223372036854775807}
        _ = v14
        _ = v23
    }
    _ = p0
    return 4294967295
}
func genFn1(p0 uint16, p1 int32) int32 {
    var v24 bool = true
    p0 = 0
    v25 := []int{32768}
    switch "foo bar" {
    case "":
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS1{F0: fi26}
            if len(v25) > 0 {
                v25[0] = -2147483648
            }
            p1 = p1
            _ = fi26
            _ = v27
        }
    case "foo bar":
        switch -1 {
        case -1:
            switch 1 {
            case -1:
                if (-343 < (^(p1))) {
                    v28 := GenS1{F0: 560}
                    _ = v28
                }
            default:
                var v29 uint16 = 200
                _ = v29
            }
        case -387:
            var v30 bool = (-1 == 127)
            _ = v30
        case -2147483648:
            for fi31 := 0; fi31 < 10; fi31++ {
                v25 = append(v25, -640)
                var v32 int16 = -1
                _ = fi31
                _ = v32
            }
        default:
            if (((-2.25 - 1.5) - 0.1) < ((0.0 - 1e10) + ((float32(1e10) / float32(-2.25)) * float32(3.14159)))) {
                p1 = p1
                fmt.Println(v24)
                if len(v25) > 0 {
                    v25[0] = -2147483648
                }
            }
        }
    default:
        var v33 int8 = -128
        _ = v33
    }
    v34 := GenS3{F0: p1, F1: p0}
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v34
    return 1
}

func main() {
    var v1 uint16 = 42
    v2 := []string{"hello", ("a" + "foo bar")}
    v3 := []uint32{331, 2}
    var v4 int64 = 538
    var v5 uint = 65535
    v6 := make(map[int]int)
    switch "foo bar" {
    case "":
        var v35 int16 = 0
        _ = v35
    default:
        if true {
            _ = len(v3)
        }
    }
    v6[32767] = 1
    v36 := GenS2{F0: (v1 << ((v5) & 63)), F1: -547}
    v5 = (^(v5))
    var v37 uint16 = v1
    for fi38 := 0; fi38 < 19; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 uint = (^(v5))
            v6[(fi39 & -920)] = 1
            _ = fi39
            _ = v40
        }
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        for fi41 := 0; fi41 < 12; fi41++ {
            if false {
                fmt.Println(fi38, fi38)
                fmt.Println(v1, fi38, v37)
                v37 -= (v37 ^ uint16(v5))
            } else {
                switch 9223372036854775807 {
                case 0:
                    fmt.Println(v1, v37, v1)
                default:
                    v37 += 2
                }
            }
            if false {
                var v42 uint16 = (v37 << ((18446744073709551615) & 63))
                if (v42 >= v42) {
                    v43 := v6[fi41]
                    v4 = v4
                    var v44 string = ((("" + "") + ("\n\t" + "hello")) + "a")
                    if false {
                        if true {
                            v45 := GenS4{F0: int(v1), F1: 714, F2: (^(v37))}
                            var v46 int64 = (v4 + -1)
                            _ = v45
                            _ = v46
                        } else {
                            var v47 int64 = v4
                            _ = v47
                        }
                        v37 -= v42
                        var v48 uint16 = v42
                        _ = v48
                    } else {
                        v49 := GenS2{F0: v42, F1: (fi41 * (v43 - (fi41 * v43)))}
                        _ = v49
                    }
                    v3 = append(v3, 0)
                    var v50 int64 = 1
                    if true {
                        var v51 float64 = float64(1e10)
                        _ = v51
                    }
                    _ = v43
                    _ = v44
                    _ = v50
                } else {
                    v5 = (^(v5))
                }
                var v52 uint32 = 42
                _ = v42
                _ = v52
            } else {
                v4 = (v4 / ((v4) | 1))
                switch "a" {
                case "":
                    if (false && (-670 <= 1)) {
                        switch "hello" {
                        case "hello":
                            v3 = append(v3, 915)
                        case "foo bar":
                            if (344 < genFn0(56)) {
                                var v53 float32 = 0.0
                                if false {
                                    fmt.Println(v4)
                                    fmt.Println(v53, v5, fi38)
                                    if false {
                                        v37 += (v37 / ((249) | 1))
                                        switch fi41 {
                                        case 1:
                                            var v54 int = (fi38 << ((v5) & 63))
                                            _ = v54
                                        }
                                    }
                                } else {
                                    v55 := GenS2{F0: v1, F1: fi41}
                                    _ = v55
                                }
                                var v56 uint64 = 1
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v53
                                _ = v56
                            } else {
                                if false {
                                    switch (fi41 ^ fi38) {
                                    case 1:
                                        v57 := GenS1{F0: fi38}
                                        _ = v57
                                    case 2147483647:
                                        if true {
                                            v6[(fi41 * fi41)] = (fi38 * (fi38 + int(v5)))
                                            switch "a" {
                                            case "foo bar":
                                                fmt.Println(fi38)
                                            case "\n\t":
                                                var v58 int64 = 1
                                                _ = v58
                                            default:
                                                switch (fi38 & fi38) {
                                                case 1:
                                                    switch "a" {
                                                    case "foo bar":
                                                        v1 = 65535
                                                    case "a":
                                                        fmt.Println(v5)
                                                    case "hello":
                                                        v1 = v37
                                                        if true {
                                                            fmt.Println(fi38, v1, v37)
                                                            if false {
                                                                var v59 bool = (!true)
                                                                v60 := GenS2{F0: (v1 - v37), F1: int(v5)}
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                v61 := GenS5{F0: 287, F1: v4}
                                                                _ = v61
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                var v62 uint8 = 42
                                                                _ = v62
                                                            case "\n\t":
                                                                v63 := GenS4{F0: 32768, F1: 2147483647, F2: (v1 - v1)}
                                                                _ = v63
                                                            case "foo bar":
                                                                switch (fi41 % ((0) | 1)) {
                                                                case -725:
                                                                    if false {
                                                                        var v64 float64 = 0.1
                                                                        if ((2147483647 <= 450) || true) {
                                                                            if (v37 > (^(v1))) {
                                                                                var v65 int8 = -112
                                                                                _ = v65
                                                                            } else {
                                                                                var v66 float32 = 0.0
                                                                                var v67 float64 = ((v64 + 0.1) * v64)
                                                                                var v68 uint16 = (v37 % ((v1) | 1))
                                                                                _ = v66
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        } else {
                                                                            v69 := GenS1{F0: (fi41 & 127)}
                                                                            if false {
                                                                                switch (-(fi38)) {
                                                                                case 519:
                                                                                    v70 := v6[fi41]
                                                                                    _ = v70
                                                                                case -24:
                                                                                    v6[fi38] = 32767
                                                                                case -262:
                                                                                    v5 -= v5
                                                                                default:
                                                                                    v64 -= (v64 + 0.0)
                                                                                }
                                                                                var v71 uint16 = (v37 >> ((42) & 63))
                                                                                v64 = 1.5
                                                                                _ = v71
                                                                            } else {
                                                                                var v72 uint32 = 185
                                                                                v73 := GenS6{F0: ("" + "\n\t"), F1: 110}
                                                                                v37 += v1
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                            if false {
                                                                                v64 = (0.1 - 0.0)
                                                                                if false {
                                                                                    var v74 float32 = (0.1 + 0.0)
                                                                                    var v75 int16 = 1000
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                }
                                                                            }
                                                                            _ = v69
                                                                        }
                                                                        _ = v64
                                                                    }
                                                                case -2147483648:
                                                                    _ = len(v3)
                                                                case 9223372036854775807:
                                                                    switch (fi38 & fi41) {
                                                                    case -9223372036854775808:
                                                                        var v76 bool = ((uint32(v5) == genFn0(229)) || (1 != 604))
                                                                        _ = v76
                                                                    }
                                                                default:
                                                                    v2 = append(v2, "hello")
                                                                }
                                                            default:
                                                                switch (((("" + "") + "hello") + "hello") + "hello") {
                                                                case "foo bar":
                                                                    if false {
                                                                        var v77 uint = (^(v5))
                                                                        _ = v77
                                                                    }
                                                                case "":
                                                                    if false {
                                                                        var v78 bool = (((2 == 2) || (230 != 42)) && ((-(v4)) == (v4 & 9223372036854775807)))
                                                                        _ = v78
                                                                    }
                                                                case "hello":
                                                                    var v79 float64 = 1.5
                                                                    _ = v79
                                                                }
                                                            }
                                                        }
                                                        switch "" {
                                                        case "hello":
                                                            if true {
                                                                var v80 float64 = (float64(1.5) - 0.1)
                                                                _ = v80
                                                            }
                                                        case "":
                                                            v81 := GenS7{F0: fi41}
                                                            _ = v81
                                                        }
                                                    default:
                                                        v37 -= v37
                                                    }
                                                case 128:
                                                    var v82 string = "\n\t"
                                                    _ = v82
                                                case 9223372036854775807:
                                                    v83 := GenS8{F0: -32768}
                                                    _ = v83
                                                default:
                                                    v4 = (v4 & v4)
                                                }
                                            }
                                            if true {
                                                v37 = (^(v37))
                                                v5 = 2
                                            }
                                        } else {
                                            switch fi38 {
                                            case 0:
                                                v84 := GenS8{F0: -866}
                                                _ = v84
                                            case 42:
                                                v85 := GenS9{F0: 0, F1: uint64(v37), F2: 9223372036854775807}
                                                _ = v85
                                            case 9223372036854775807:
                                                var v86 int = fi41
                                                _ = v86
                                            default:
                                                if false {
                                                    var v87 float32 = 3.14159
                                                    switch (fi38 - 4294967295) {
                                                    case -1:
                                                        v88 := GenS5{F0: -1, F1: (v4 / (((v4 / (((v4 % ((v4) | 1))) | 1))) | 1))}
                                                        _ = v88
                                                    }
                                                    _ = v87
                                                }
                                            }
                                            if ((false || (42 < 87)) || ((1 == 128) && true)) {
                                                v37 += (v37 >> ((v5) & 63))
                                                v37 = v1
                                            }
                                            _ = len(v6)
                                        }
                                    default:
                                        v4 -= (v4 - v4)
                                    }
                                    var v89 float64 = 1.5
                                    var v90 uint = 2
                                    _ = v89
                                    _ = v90
                                } else {
                                    var v91 int64 = 0
                                    _ = v91
                                }
                                if false {
                                    v92 := GenS10{F0: fi38}
                                    _ = v92
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    v93 := GenS11{F0: 2, F1: 42}
                                    _ = v93
                                }
                            }
                        case "\n\t":
                            var v94 int64 = (v4 | -724)
                            _ = v94
                        default:
                            switch (fi41 - (fi41 ^ (-(fi38)))) {
                            case 0:
                                v5 = v5
                            default:
                                fmt.Println(v1, v4, v37)
                            }
                        }
                    } else {
                        switch (fi38 - (fi41 - (fi41 / ((253) | 1)))) {
                        case 0:
                            v1 += v37
                        }
                        if ((v5 % ((uint(v37)) | 1)) != (v5 ^ (v5 & (^(v5))))) {
                            v3 = append(v3, genFn0(uint8(v5)))
                            var v95 uint16 = 255
                            _ = v95
                        }
                    }
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        v6[fi38] = int(v5)
                    case "foo bar":
                        v4 -= v4
                    default:
                        fmt.Println(fi38, fi41)
                    }
                case "hello":
                    v96 := GenS1{F0: fi41}
                    _ = v96
                default:
                    v1 += (^(v37))
                }
            }
            v97 := GenS11{F0: 194, F1: 0}
            _ = fi41
            _ = v97
        }
        _ = fi38
    }
    for fi98 := 0; fi98 < 9; fi98++ {
        v5 = v5
        _ = fi98
    }
    for fi99 := 0; fi99 < 1; fi99++ {
        var v100 int64 = -1
        _ = fi99
        _ = v100
    }
    var v101 int32 = int32(v4)
    var v102 int = 203
    v102 = v102
    switch ("\n\t" + (("" + ("foo bar" + ("\n\t" + (("" + ("a" + "foo bar")) + ((((("foo bar" + (("" + (("foo bar" + "a") + ("hello" + "hello"))) + "")) + "hello") + "\n\t") + "a") + "a"))))) + "\n\t")) {
    case "":
        v102 = v102
    case "\n\t":
        fmt.Println(v5, v37)
    case "a":
        v102 = int(v5)
    }
    fmt.Println(v4, v102)
    v37 = 42
    var v103 float64 = 3.14159
    switch (("" + "\n\t") + (("hello" + "\n\t") + "hello")) {
    case "a":
        v104 := GenS10{F0: (-(v102))}
        _ = v104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
