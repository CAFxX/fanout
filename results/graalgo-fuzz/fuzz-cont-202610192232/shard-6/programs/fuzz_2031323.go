package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 int64; F1 int16; F2 int }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
type GenS4 struct { F0 int32; F1 int32; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 int; F1 uint32; F2 int16 }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS13 struct { F0 uint16; F1 int8; F2 bool }
type GenS14 struct { F0 uint64; F1 bool; F2 int32 }
type GenS15 struct { F0 int64; F1 int16; F2 uint32 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int16; F1 uint; F2 int16 }
type GenS18 struct { F0 uint; F1 int64 }
func genFn0(p0 int32, p1 int64, p2 uint) string {
    v6 := []int{128, 554, -733}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 14; fi8++ {
        p2 += 0
        v6 = append(v6, (^(fi8)))
        switch fi8 {
        case 127:
            var v9 uint = p2
            _ = v9
        case -768:
            for fi10 := 0; fi10 < 16; fi10++ {
                v7 = ""
                _ = fi10
            }
        }
        _ = fi8
    }
    v11 := GenS1{F0: 42, F1: p0, F2: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int8, p1 uint32, p2 uint64) int16 {
    if false {
        var v12 uint8 = 2
        _ = v12
    }
    v13 := GenS2{F0: 42, F1: 42, F2: 1}
    v14 := GenS3{F0: (42 != 2), F1: -1, F2: p0}
    v15 := GenS4{F0: 323, F1: 1, F2: 127}
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    var v2 string = ""
    var v3 float32 = (1.5 / -2.25)
    var v4 uint16 = 65535
    var v5 float64 = ((0.0 / (0.1 + ((1e10 - 1.5) - float64(float32(float64(v3)))))) / float64(-2.25))
    if true {
        switch 117 {
        case 9223372036854775807:
            v16 := []int{42, 1, -9223372036854775808}
            _ = v16
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint64 = uint64(v4)
            switch fi17 {
            case 448:
                switch 1 {
                case -2147483648:
                    switch v2 {
                    case "":
                        var v19 int = fi17
                        _ = v19
                    }
                case 771:
                    var v20 int = fi17
                    _ = v20
                }
            }
            fmt.Println(fi17, v3, v1)
            _ = fi17
            _ = v18
        }
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            if false {
                v22 := GenS2{F0: 42, F1: 32767, F2: 32768}
                v4 -= 424
                fmt.Println(v2, v4)
                _ = v22
            } else {
                var v23 float32 = (v3 + (v3 - v3))
                _ = v23
            }
            if (3.14159 > v3) {
                v24 := GenS5{F0: int16(fi21)}
                if (777 <= uint(v4)) {
                    v25 := []int{fi21, (-(fi21)), (fi21 | 9223372036854775807)}
                    v26 := GenS1{F0: (v1 % (((^(v1))) | 1)), F1: 1, F2: "\n\t"}
                    _ = v25
                    _ = v26
                }
                switch fi21 {
                case -2147483648:
                    switch 4294967295 {
                    case 32768:
                        fmt.Println(v2, v4)
                    case 9223372036854775807:
                        if (!true) {
                            if true {
                                v27 := GenS6{F0: 2}
                                _ = v27
                            } else {
                                v28 := GenS7{F0: 4294967295, F1: v1}
                                _ = v28
                            }
                            var v29 int64 = -386
                            v30 := []int{fi21, (fi21 / ((fi21) | 1))}
                            _ = v29
                            _ = v30
                        }
                    case 127:
                        var v31 uint8 = 127
                        _ = v31
                    }
                case 9223372036854775807:
                    v5 = v5
                case 0:
                    var v32 int32 = 0
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 2; fi33++ {
                        switch (v2 + "foo bar") {
                        case "a":
                            if false {
                                if true {
                                    v5 = ((0.1 * v5) + v5)
                                }
                                var v34 int8 = 42
                                _ = v34
                            }
                        case "\n\t":
                            v2 = ("\n\t" + v2)
                        case "":
                            v35 := GenS8{F0: 101, F1: 2147483647}
                            _ = v35
                        default:
                            var v36 int8 = 100
                            _ = v36
                        }
                        _ = fi33
                    }
                }
                _ = v24
            }
            _ = fi21
        }
        if false {
            var v37 int = 9223372036854775807
            v38 := []int{(-(v37))}
            _ = v37
            _ = v38
        }
        if false {
            var v39 uint64 = 343
            if (!true) {
                var v40 bool = true
                _ = v40
            }
            _ = v39
        } else {
            var v41 uint = 18446744073709551615
            _ = v41
        }
    }
    v42 := []int{-9223372036854775808, 1, 128, -9223372036854775808}
    var v43 bool = true
    v5 = -2.25
    for fi44 := 0; fi44 < 13; fi44++ {
        var v45 float64 = ((-2.25 + float64(0.1)) / v5)
        v43 = (99 < 42)
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 20; fi46++ {
        v47 := GenS9{F0: 1, F1: 65535, F2: 42}
        _ = fi46
        _ = v47
    }
    if v43 {
        v48 := GenS10{F0: 1000, F1: v1}
        for fi49 := 0; fi49 < 9; fi49++ {
            switch v2 {
            case "":
                if (v43 || (!((^(fi49)) > 4294967295))) {
                    var v50 int8 = 1
                    if false {
                        v51 := GenS11{F0: v1}
                        _ = v51
                    }
                    if (!((v2 + "foo bar") < v2)) {
                        var v52 int8 = v50
                        v53 := GenS1{F0: v1, F1: 2147483647, F2: "hello"}
                        _ = v52
                        _ = v53
                    } else {
                        var v54 bool = v43
                        _ = v54
                    }
                    _ = v50
                }
            default:
                var v55 int16 = 32767
                _ = v55
            }
            var v56 uint16 = (v4 % ((v4) | 1))
            var v57 uint32 = (v1 << ((255) & 63))
            _ = fi49
            _ = v56
            _ = v57
        }
        _ = v48
    } else {
        v58 := GenS12{F0: 450, F1: 255, F2: -32768}
        if (-1 >= -206) {
            switch "foo bar" {
            case "hello":
                for fi59 := 0; fi59 < 18; fi59++ {
                    if v43 {
                        fmt.Println(v2)
                        v60 := GenS13{F0: (^(v4)), F1: 0, F2: false}
                        _ = v60
                    }
                    _ = fi59
                }
            case "":
                v5 = 3.14159
            }
            v5 = v5
            v5 = ((v5 / 1.0) * 3.14159)
        }
        for fi61 := 0; fi61 < 15; fi61++ {
            switch (fi61 & fi61) {
            case -2147483648:
                v43 = (true && v43)
            case 32768:
                if false {
                    if v43 {
                        if (v43 || true) {
                            for fi62 := 0; fi62 < 11; fi62++ {
                                if (false && ((1 < -128) && v43)) {
                                    v63 := GenS8{F0: 127, F1: -2147483648}
                                    var v64 bool = (2147483647 < -1)
                                    fmt.Println(v43)
                                    _ = v63
                                    _ = v64
                                } else {
                                    switch (^(fi61)) {
                                    case 32767:
                                        var v65 float32 = v3
                                        _ = v65
                                    case 9223372036854775807:
                                        v1 -= v1
                                    case 627:
                                        v66 := GenS8{F0: 128, F1: 0}
                                        _ = v66
                                    }
                                    v4 -= 32767
                                }
                                _ = fi62
                            }
                            var v67 float32 = ((v3 + 3.14159) - (v3 * v3))
                            v68 := GenS10{F0: 62, F1: 65535}
                            _ = v67
                            _ = v68
                        }
                        v69 := GenS1{F0: v1, F1: 0, F2: "foo bar"}
                        _ = v69
                    } else {
                        if len(v42) > 0 {
                            v42[0] = fi61
                        }
                        v70 := GenS14{F0: 353, F1: (127 != 127), F2: -2147483648}
                        var v71 float64 = (1.5 / (v5 * ((v5 / (((float64(v3) / float64(0.1)) / v5) - -2.25)) - v5)))
                        _ = v70
                        _ = v71
                    }
                    if len(v42) > 0 {
                        v42[0] = 127
                    }
                    if (-754 >= 1) {
                        fmt.Println(v1, v4)
                        switch "a" {
                        case "\n\t":
                            for fi72 := 0; fi72 < 19; fi72++ {
                                if false {
                                    var v73 float32 = ((v3 - (float32(v5) + ((0.1 / v3) * -2.25))) * 0.1)
                                    var v74 float64 = v5
                                    var v75 float32 = 1e10
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                }
                                v43 = ("" <= v2)
                                _ = fi72
                            }
                        case "a":
                            var v76 int8 = -25
                            _ = v76
                        case "":
                            v42 = append(v42, fi61)
                        }
                        var v77 int32 = 42
                        _ = v77
                    }
                }
            }
            v3 -= v3
            if (true && v43) {
                for fi78 := 0; fi78 < 5; fi78++ {
                    var v79 uint16 = v4
                    if v43 {
                        switch "a" {
                        case "hello":
                            if (!(!true)) {
                                v80 := GenS7{F0: 0, F1: (v1 | 2147483647)}
                                _ = v80
                            } else {
                                var v81 uint16 = (v4 + (v79 << ((224) & 63)))
                                var v82 uint = 880
                                v1 = (v1 & 2)
                                _ = v81
                                _ = v82
                            }
                        case "":
                            fmt.Println(v4)
                        case "foo bar":
                            v5 = ((-2.25 - 0.1) - 1e10)
                        }
                    } else {
                        var v83 uint = uint(v1)
                        _ = v83
                    }
                    var v84 uint64 = 42
                    _ = fi78
                    _ = v79
                    _ = v84
                }
                var v85 int16 = -1
                _ = v85
            }
            _ = fi61
        }
        _ = v58
    }
    switch int(v1) {
    case 42:
        switch 42 {
        case 36:
            if (42 < 923) {
                v3 += v3
            }
        case 4294967295:
            v1 -= v1
        case 42:
            if (!(4294967295 == v1)) {
                if ((!v43) || (v3 == v3)) {
                    for fi86 := 0; fi86 < 8; fi86++ {
                        if (((v2 + v2) + ((v2 + v2) + ("hello" + (v2 + "\n\t")))) < "hello") {
                            switch (fi86 - fi86) {
                            case 32767:
                                var v87 int16 = genFn1(0, (^(v1)), 443)
                                _ = v87
                            default:
                                var v88 uint32 = (v1 << ((1) & 63))
                                _ = v88
                            }
                            fmt.Println(v5, v5, v3)
                            if (463 > 1) {
                                var v89 uint32 = v1
                                _ = v89
                            } else {
                                var v90 uint32 = 374
                                if v43 {
                                    var v91 uint = 1
                                    _ = v91
                                } else {
                                    _ = len(v42)
                                    var v92 int64 = -1
                                    _ = v92
                                }
                                _ = v90
                            }
                        }
                        if (0 >= v4) {
                            if true {
                                if v43 {
                                    fmt.Println(v2, v43, v43)
                                    _ = len(v42)
                                } else {
                                    fmt.Println(v5, v4)
                                    v43 = ((v2 + "") < "")
                                    var v93 uint16 = (v4 - 32767)
                                    _ = v93
                                }
                            }
                        }
                        _ = len(v42)
                        _ = fi86
                    }
                    var v94 float64 = v5
                    _ = v94
                }
                v4 = v4
                for fi95 := 0; fi95 < 19; fi95++ {
                    var v96 int16 = 42
                    switch (v2 + (v2 + "hello")) {
                    case "":
                        v97 := GenS15{F0: 0, F1: genFn1(118, v1, 4294967295), F2: v1}
                        _ = v97
                    case "a":
                        v98 := GenS16{F0: ""}
                        _ = v98
                    case "\n\t":
                        v99 := GenS17{F0: (v96 * (v96 ^ v96)), F1: 255, F2: (v96 - 32767)}
                        _ = v99
                    default:
                        var v100 uint16 = (v4 | 2)
                        _ = v100
                    }
                    if len(v42) >= 2 {
                        v42 = v42[1:len(v42)]
                    }
                    _ = fi95
                    _ = v96
                }
            } else {
                var v101 uint16 = v4
                v102 := GenS5{F0: 42}
                v103 := GenS18{F0: 503, F1: 1}
                _ = v101
                _ = v102
                _ = v103
            }
        }
    case 48:
        switch ("\n\t" + v2) {
        case "":
            switch "hello" {
            case "hello":
                for fi104 := 0; fi104 < 3; fi104++ {
                    var v105 int16 = 1000
                    _ = fi104
                    _ = v105
                }
            case "foo bar":
                var v106 bool = (1000 == -83)
                _ = v106
            case "":
                for fi107 := 0; fi107 < 16; fi107++ {
                    for fi108 := 0; fi108 < 15; fi108++ {
                        var v109 int32 = 1
                        v1 = (v1 / ((4294967295) | 1))
                        _ = fi108
                        _ = v109
                    }
                    _ = fi107
                }
            default:
                v4 = v4
            }
        }
    }
    v4 -= 42
    v5 -= (v5 + 1e10)
    fmt.Println(v1)
    if v43 {
        fmt.Println(v2, v3)
    }
    var v110 uint = 18446744073709551615
    if len(v42) > 0 {
        v42[0] = 801
    }
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    fmt.Println(v1, v43)
    v2 = "a"
    var v111 uint32 = (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
