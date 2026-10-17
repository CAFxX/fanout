package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 string; F1 bool; F2 int64 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int32; F1 uint; F2 int }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 uint; F1 uint; F2 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8; F1 int64; F2 int16 }
type GenS13 struct { F0 uint32 }
func genFn0(p0 uint16, p1 bool, p2 int32) int16 {
    if (p2 >= (-(p2))) {
        v6 := GenS1{F0: 39, F1: p1}
        var v7 int16 = -32768
        _ = v6
        _ = v7
    }
    v8 := []int{2147483647, 32768}
    var v9 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []int{}
    v2 := []string{}
    var v3 int16 = -942
    v4 := make(map[string]uint32)
    var v5 uint = 243
    if (!true) {
        v3 -= 373
        v3 -= v3
    }
    switch ("\n\t" + "") {
    case "hello":
        if false {
            v3 -= v3
            for fi10 := 0; fi10 < 15; fi10++ {
                switch 0 {
                case 127:
                    v5 = v5
                case 9223372036854775807:
                    v11 := GenS1{F0: 1, F1: true}
                    _ = v11
                }
                _ = fi10
            }
            var v12 uint16 = 2
            _ = v12
        }
    case "a":
        switch 4294967295 {
        case 4294967295:
            v5 += v5
        case -743:
            if false {
                v13 := GenS2{F0: 90, F1: 4294967295, F2: 19}
                _ = v13
            }
        case -2147483648:
            v3 += (v3 >> ((v5) & 63))
        }
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v5 = (^(v5))
    var v14 string = "hello"
    v14 = (("" + "hello") + "foo bar")
    if false {
        if true {
            var v15 string = v14
            switch "hello" {
            case "a":
                if (-1 != -472) {
                    switch -898 {
                    case -2147483648:
                        var v16 int16 = 631
                        _ = v16
                    case 9223372036854775807:
                        v17 := v4["a"]
                        _ = v17
                    case 857:
                        for fi18 := 0; fi18 < 8; fi18++ {
                            if (!(681 > 0)) {
                                _ = len(v1)
                            }
                            _ = fi18
                        }
                    }
                    var v19 uint64 = 171
                    v1 = append(v1, -9223372036854775808)
                    _ = v19
                } else {
                    var v20 int = -1
                    var v21 uint = v5
                    v5 = (v5 % ((0) | 1))
                    _ = v20
                    _ = v21
                }
            case "\n\t":
                fmt.Println(v5)
            }
            _ = v15
        } else {
            var v22 uint32 = 65535
            v5 += (v5 * (v5 / (((^(v5))) | 1)))
            v23 := GenS3{F0: v14, F1: 0, F2: v5}
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 13; fi24++ {
            switch 852 {
            case 282:
                v25 := GenS2{F0: 105, F1: (^(v5)), F2: 441}
                _ = v25
            case -48:
                v3 = -196
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                if (!false) {
                    if false {
                        if true {
                            v4["foo bar"] = 2147483647
                        }
                        var v27 int64 = 0
                        _ = v27
                    } else {
                        v28 := GenS4{F0: "foo bar", F1: false, F2: -9223372036854775808}
                        v29 := GenS1{F0: 0, F1: true}
                        fmt.Println(v3, v3, v14)
                        _ = v28
                        _ = v29
                    }
                    var v30 bool = true
                    _ = v30
                } else {
                    v5 += (v5 / ((2) | 1))
                }
                v31 := GenS4{F0: v14, F1: false, F2: 9223372036854775807}
                fmt.Println(fi26, v5, v3)
                _ = fi26
                _ = v31
            }
            _ = fi24
        }
        v32 := GenS5{F0: 128, F1: -487}
        _ = v32
    } else {
        var v33 bool = (!(-570 > -9223372036854775808))
        switch int(v3) {
        case 128:
            if len(v2) > 0 {
                v2[0] = v14
            }
        case 739:
            for fi34 := 0; fi34 < 8; fi34++ {
                switch (fi34 | 32768) {
                case 4294967295:
                    v35 := GenS1{F0: 917, F1: v33}
                    _ = v35
                }
                _ = fi34
            }
        case 911:
            delete(v4, "a")
        default:
            v36 := GenS6{F0: -462, F1: v5, F2: 9223372036854775807}
            _ = v36
        }
        if false {
            if v33 {
                delete(v4, v14)
            }
            var v37 int64 = -245
            _ = v37
        } else {
            if (1 == 2) {
                v38 := GenS7{F0: 672, F1: v5}
                _ = v38
            } else {
                var v39 uint16 = 498
                v40 := v4["a"]
                var v41 int8 = -128
                _ = v39
                _ = v40
                _ = v41
            }
        }
        _ = v33
    }
    if true {
        if true {
            var v42 int32 = 42
            v43 := GenS8{F0: 640, F1: 1}
            _ = v42
            _ = v43
        } else {
            if true {
                for fi44 := 0; fi44 < 5; fi44++ {
                    v5 = 4294967295
                    for fi45 := 0; fi45 < 15; fi45++ {
                        _ = len(v1)
                        _ = fi45
                    }
                    _ = len(v2)
                    _ = fi44
                }
            } else {
                v5 += v5
                fmt.Println(v3, v14)
                if (false && true) {
                    v46 := v4[""]
                    _ = v46
                }
            }
            _ = len(v4)
            if false {
                var v47 int64 = 9223372036854775807
                if (!false) {
                    var v48 uint16 = 32767
                    _ = v48
                }
                _ = v47
            } else {
                v2 = append(v2, ((v14 + v14) + ""))
                switch 127 {
                case 367:
                    v14 = "\n\t"
                }
            }
        }
        switch (v14 + v14) {
        case "hello":
            switch "a" {
            case "hello":
                if (!((false && (true || (65535 != v5))) && true)) {
                    v14 = v14
                } else {
                    v49 := GenS4{F0: "\n\t", F1: (!(255 == 42)), F2: -9223372036854775808}
                    _ = v49
                }
            case "foo bar":
                if true {
                    switch "" {
                    case "\n\t":
                        fmt.Println(v14, v5, v5)
                    case "hello":
                        switch v14 {
                        case "\n\t":
                            v50 := GenS9{F0: 254, F1: 970}
                            _ = v50
                        case "hello":
                            fmt.Println(v3, v5, v14)
                        case "a":
                            if (506 == 753) {
                                for fi51 := 0; fi51 < 5; fi51++ {
                                    switch (fi51 >> (((v5 << ((v5) & 63))) & 63)) {
                                    case 42:
                                        v3 = (v3 << ((v5) & 63))
                                    }
                                    var v52 string = "foo bar"
                                    _ = fi51
                                    _ = v52
                                }
                            }
                        }
                    default:
                        switch "hello" {
                        case "\n\t":
                            v5 -= (v5 >> ((647) & 63))
                        }
                    }
                    v14 = v14
                } else {
                    if true {
                        var v53 uint8 = 206
                        v5 += v5
                        _ = v53
                    } else {
                        fmt.Println(v5, v14)
                        switch v14 {
                        case "\n\t":
                            for fi54 := 0; fi54 < 18; fi54++ {
                                if (!false) {
                                    var v55 uint64 = 18446744073709551615
                                    _ = v55
                                } else {
                                    for fi56 := 0; fi56 < 20; fi56++ {
                                        if true {
                                            _ = len(v2)
                                            if (!false) {
                                                delete(v4, v14)
                                                fmt.Println(v5, fi54)
                                                var v57 uint8 = 0
                                                v5 = 352
                                                _ = v57
                                            } else {
                                                if false {
                                                    fmt.Println(fi54, v14)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v58 := GenS10{F0: 0, F1: 4294967295, F2: v3}
                                                _ = v58
                                            }
                                            switch (-(fi56)) {
                                            case 2147483647:
                                                if false {
                                                    if true {
                                                        v5 += v5
                                                        var v59 bool = false
                                                        fmt.Println(fi54, v5, v5)
                                                        _ = v59
                                                    } else {
                                                        fmt.Println(v14, fi54, v3)
                                                    }
                                                    v4[("foo bar" + v14)] = 2
                                                } else {
                                                    v5 += 0
                                                }
                                                fmt.Println(v14, v14)
                                                if (42 == 42) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v14 = "a"
                                                    var v60 uint = (^(v5))
                                                    _ = v60
                                                } else {
                                                    switch ((("\n\t" + v14) + (v14 + v14)) + v14) {
                                                    case "":
                                                        switch 324 {
                                                        case 4294967295:
                                                            v5 = (v5 | v5)
                                                        case 820:
                                                            _ = len(v1)
                                                        case 9223372036854775807:
                                                            switch v14 {
                                                            case "hello":
                                                                fmt.Println(v14)
                                                            }
                                                        default:
                                                            v61 := GenS5{F0: 127, F1: 0}
                                                            _ = v61
                                                        }
                                                    case "\n\t":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    v14 = ("" + ("a" + "hello"))
                                                    v62 := GenS11{F0: 1}
                                                    fmt.Println(v3, v3)
                                                    _ = v62
                                                }
                                            case -9223372036854775808:
                                                v3 -= -32768
                                            case -68:
                                                switch 42 {
                                                case 105:
                                                    fmt.Println(fi54)
                                                case -1:
                                                    var v63 uint8 = 218
                                                    _ = v63
                                                default:
                                                    v5 += v5
                                                }
                                                var v64 uint64 = 1
                                                if (false && true) {
                                                    v65 := GenS1{F0: (v64 % ((42) | 1)), F1: true}
                                                    v64 = (v64 - 0)
                                                    switch (((v14 + v14) + "hello") + (v14 + ((("hello" + (("foo bar" + v14) + v14)) + v14) + "\n\t"))) {
                                                    case "foo bar":
                                                        v64 = (v64 << (((^(v5))) & 63))
                                                    case "\n\t":
                                                        v3 = v3
                                                    }
                                                    v66 := v4[v14]
                                                    v5 = (v5 / ((v5) | 1))
                                                    v5 = (v5 + 0)
                                                    fmt.Println(v14, v3, v3)
                                                    fmt.Println(v5, fi54, v64)
                                                    if true {
                                                        var v67 int = (fi54 & fi56)
                                                        delete(v4, v14)
                                                        _ = v67
                                                    } else {
                                                        v66 = v66
                                                        switch v14 {
                                                        case "\n\t":
                                                            var v68 uint16 = 255
                                                            _ = v68
                                                        case "hello":
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v14 = ""
                                                            }
                                                            v69 := v4[(v14 + v14)]
                                                            _ = v69
                                                        case "a":
                                                            v70 := GenS12{F0: 153, F1: 0, F2: genFn0(117, true, 657)}
                                                            v71 := GenS10{F0: v5, F1: (v5 * v5), F2: (v3 - (v3 * v3))}
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                    }
                                                    v72 := GenS13{F0: (v66 * (v66 & v66))}
                                                    fmt.Println(v5, v14, v3)
                                                    _ = v65
                                                    _ = v66
                                                    _ = v72
                                                }
                                                _ = v64
                                            default:
                                                switch "a" {
                                                case "hello":
                                                    var v73 uint8 = 208
                                                    _ = v73
                                                default:
                                                    var v74 int8 = 0
                                                    _ = v74
                                                }
                                            }
                                        } else {
                                            v14 = v14
                                        }
                                        var v75 int64 = -9223372036854775808
                                        var v76 uint32 = 1
                                        _ = fi56
                                        _ = v75
                                        _ = v76
                                    }
                                    var v77 int16 = 1000
                                    _ = v77
                                }
                                if (("foo bar" + "") >= v14) {
                                    var v78 float32 = 1.5
                                    _ = v78
                                } else {
                                    if (!false) {
                                        v4[""] = 0
                                        for fi79 := 0; fi79 < 17; fi79++ {
                                            v5 = (v5 | v5)
                                            _ = fi79
                                        }
                                    } else {
                                        v5 -= v5
                                    }
                                    v14 = "hello"
                                    if ((((0.1 + 1.5) < ((1e10 - ((float32(float64(-2.25)) + 1e10) + 1e10)) / ((1.5 / (-2.25 / (-2.25 * 0.1))) + (float32(0.0) + 3.14159)))) && (false && false)) || (32767 == 2)) {
                                        for fi80 := 0; fi80 < 6; fi80++ {
                                            if (0.0 == float64(0.1)) {
                                                var v81 bool = (0.1 <= 0.1)
                                                if (("hello" > v14) || ((v3 * v3) < v3)) {
                                                    var v82 uint8 = 86
                                                    _ = v82
                                                } else {
                                                    if (((("hello" + "\n\t") + "hello") + v14) <= v14) {
                                                        var v83 int16 = 0
                                                        _ = v83
                                                    } else {
                                                        fmt.Println(fi80, v5, v5)
                                                        if true {
                                                            fmt.Println(fi54)
                                                            v84 := v4[(v14 + "a")]
                                                            fmt.Println(v5)
                                                            _ = v84
                                                        }
                                                    }
                                                    switch v14 {
                                                    case "foo bar":
                                                        if true {
                                                            v81 = v81
                                                            var v85 int16 = v3
                                                            _ = v85
                                                        }
                                                    default:
                                                        var v86 uint8 = 184
                                                        _ = len(v2)
                                                        _ = v86
                                                    }
                                                }
                                                _ = v81
                                            } else {
                                                var v87 uint8 = 127
                                                if (false && true) {
                                                    v87 += v87
                                                    v3 -= 42
                                                } else {
                                                    v3 += v3
                                                }
                                                var v88 uint16 = 0
                                                _ = v87
                                                _ = v88
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = (^(fi54))
                                            }
                                            _ = fi80
                                        }
                                    }
                                }
                                v5 = 2
                                _ = fi54
                            }
                        case "a":
                            var v89 int32 = 2147483647
                            _ = v89
                        }
                        v3 = v3
                    }
                    for fi90 := 0; fi90 < 3; fi90++ {
                        var v91 uint = v5
                        var v92 int16 = v3
                        fmt.Println(v5, fi90, fi90)
                        _ = fi90
                        _ = v91
                        _ = v92
                    }
                }
            case "":
                v5 = v5
            }
        case "\n\t":
            for fi93 := 0; fi93 < 17; fi93++ {
                var v94 int16 = (v3 / ((0) | 1))
                _ = fi93
                _ = v94
            }
        case "":
            v14 = ("" + ("hello" + ("hello" + v14)))
        default:
            var v95 uint64 = 4294967295
            _ = v95
        }
        if false {
            _ = len(v2)
        }
    } else {
        if (false || (!(0 == 4294967295))) {
            var v96 uint16 = 234
            _ = v96
        } else {
            var v97 int32 = -809
            v98 := GenS2{F0: 113, F1: 65535, F2: -1}
            _ = v97
            _ = v98
        }
    }
    v4[("\n\t" + (((v14 + "hello") + v14) + (v14 + v14)))] = 65535
    v99 := v4[v14]
    v99 = v99
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v99 = (^(v99))
    v100 := GenS11{F0: 0}
    var v101 int16 = v3
    switch 42 {
    case 2147483647:
        v102 := GenS6{F0: 42, F1: (v5 | 18446744073709551615), F2: 1}
        _ = v102
    default:
        switch "\n\t" {
        case "\n\t":
            v3 = v3
        case "hello":
            v101 += 1000
        default:
            v103 := GenS2{F0: 127, F1: (v5 * v5), F2: 9223372036854775807}
            _ = v103
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
