package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int64; F1 uint32; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int64; F1 uint8; F2 bool }
type GenS8 struct { F0 int8; F1 int8; F2 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 uint16 }
func genFn0() int64 {
    var v4 int = -178
    if true {
        v5 := GenS1{F0: 86}
        fmt.Println(v4)
        _ = v5
    } else {
        v4 += (v4 | (v4 << ((2) & 63)))
        var v6 int64 = 947
        if (false && (42 == 115)) {
            var v7 uint = 388
            if false {
                if ((("a" + "foo bar") != "") && (148 <= 42)) {
                    for fi8 := 0; fi8 < 20; fi8++ {
                        fmt.Println(v6)
                        var v9 int8 = -52
                        fmt.Println(v4, v6, v9)
                        _ = fi8
                        _ = v9
                    }
                    if (!true) {
                        v10 := []int{(^(v4)), (^(v4))}
                        v6 += (v6 & v6)
                        _ = v10
                    } else {
                        v7 -= 4294967295
                        var v11 int64 = (v6 & 9223372036854775807)
                        _ = v11
                    }
                }
                v7 = 4294967295
            } else {
                var v12 bool = false
                var v13 string = (("foo bar" + ("" + "foo bar")) + "a")
                v7 = (v7 - v7)
                _ = v12
                _ = v13
            }
            _ = v7
        } else {
            switch v4 {
            case 4294967295:
                v6 -= -640
            case 1:
                var v14 int16 = 1000
                _ = v14
            default:
                v4 = v4
            }
        }
        _ = v6
    }
    if (9223372036854775807 != -9223372036854775808) {
        if (3.14159 > ((((3.14159 - (0.0 - 3.14159)) * 1e10) * 1e10) * (1.5 - 0.0))) {
            fmt.Println(v4)
        } else {
            v15 := GenS1{F0: 42}
            v4 = (^(v4))
            _ = v15
        }
    }
    v16 := []int{-9223372036854775808}
    v4 = (v4 | v4)
    _ = v16
    _ = v4
    return 42
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 0
    v17 := GenS1{F0: 255}
    switch 2147483647 {
    case -234:
        v18 := GenS2{F0: -1, F1: (!true)}
        _ = v18
    case 186:
        delete(v2, 2147483647)
    case 333:
        v19 := GenS2{F0: 1000, F1: ((1.5 * 1.5) > 1e10)}
        _ = v19
    default:
        var v20 int64 = 1
        _ = v20
    }
    fmt.Println(v3)
    if (1 > 0) {
        switch "" {
        case "foo bar":
            v21 := GenS2{F0: 441, F1: (!false)}
            _ = v21
        }
        if false {
            if false {
                v3 = (v3 | v3)
                v3 += v3
            }
            fmt.Println(v3)
            var v22 uint8 = 0
            _ = v22
        }
        if false {
            v23 := GenS1{F0: 248}
            _ = v23
        } else {
            if true {
                if (false && false) {
                    v3 = v3
                }
            }
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        fmt.Println(v3, v3)
        v25 := GenS1{F0: 88}
        _ = fi24
        _ = v25
    }
    var v26 uint8 = 127
    v3 = 18446744073709551615
    if true {
        fmt.Println(v26, v26)
    }
    if (true && false) {
        if false {
            _ = len(v1)
        } else {
            v3 -= v3
            switch "" {
            case "hello":
                fmt.Println(v3, v3)
            case "\n\t":
                v1 = append(v1, "a")
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            for fi27 := 0; fi27 < 13; fi27++ {
                for fi28 := 0; fi28 < 9; fi28++ {
                    v3 -= 1
                    var v29 float64 = 0.1
                    var v30 float32 = float32((-2.25 + 0.0))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
        }
        fmt.Println(v3, v26)
    } else {
        var v31 uint8 = 79
        fmt.Println(v3, v31)
        _ = v31
    }
    if ((v26 ^ 128) >= v26) {
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "")
        }
    }
    var v32 int64 = genFn0()
    for fi33 := 0; fi33 < 2; fi33++ {
        _ = len(v2)
        fmt.Println(fi33, v3, v32)
        v32 = v32
        _ = fi33
    }
    switch -923 {
    case 441:
        v26 -= v26
    case -2147483648:
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 float32 = 0.0
            v36 := GenS1{F0: 1}
            _ = fi34
            _ = v35
            _ = v36
        }
    default:
        v37 := GenS2{F0: -485, F1: false}
        _ = v37
    }
    _ = len(v1)
    switch "" {
    case "":
        switch -2147483648 {
        case -298:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        case 0:
            var v38 uint16 = 42
            _ = v38
        case 9223372036854775807:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    _ = len(v2)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v32)
    case "foo bar":
        v39 := GenS1{F0: (v26 ^ 128)}
        _ = v39
    case "a":
        switch "foo bar" {
        case "\n\t":
            v40 := GenS1{F0: v26}
            _ = v40
        case "":
            v32 += (v32 & -9223372036854775808)
        case "hello":
            for fi41 := 0; fi41 < 17; fi41++ {
                if (-1 == 30) {
                    v42 := GenS3{F0: (v32 << ((2) & 63)), F1: 4294967295, F2: "\n\t"}
                    var v43 int16 = 1000
                    var v44 bool = (-443 == -32768)
                    _ = v42
                    _ = v43
                    _ = v44
                }
                var v45 string = "foo bar"
                if (true && true) {
                    for fi46 := 0; fi46 < 18; fi46++ {
                        if false {
                            v26 += 210
                        } else {
                            var v47 int32 = -682
                            v47 = int32(v26)
                            _ = v47
                        }
                        if true {
                            if false {
                                var v48 uint32 = 65535
                                fmt.Println(fi41, v45, v45)
                                var v49 uint = 65535
                                _ = v48
                                _ = v49
                            }
                        }
                        _ = fi46
                    }
                    switch (-(fi41)) {
                    case 2147483647:
                        for fi50 := 0; fi50 < 8; fi50++ {
                            var v51 uint8 = v26
                            _ = fi50
                            _ = v51
                        }
                    case 128:
                        v52 := GenS3{F0: v32, F1: 2147483648, F2: "foo bar"}
                        _ = v52
                    case -263:
                        if true {
                            var v53 string = v45
                            switch (fi41 & fi41) {
                            case -838:
                                fmt.Println(fi41, fi41)
                            case -672:
                                switch (fi41 ^ fi41) {
                                case 1:
                                    if true {
                                        var v54 uint16 = 462
                                        if ((v26 | uint8(v3)) >= 127) {
                                            var v55 string = ""
                                            var v56 int8 = 100
                                            if (v54 < 1) {
                                                delete(v2, 2147483647)
                                                for fi57 := 0; fi57 < 1; fi57++ {
                                                    v2[uint32(v26)] = 2
                                                    delete(v2, 1)
                                                    _ = fi57
                                                }
                                            }
                                            _ = v55
                                            _ = v56
                                        }
                                        _ = v54
                                    }
                                case 32768:
                                    v26 = v26
                                default:
                                    var v58 int64 = (v32 ^ (v32 << ((0) & 63)))
                                    _ = v58
                                }
                            default:
                                var v59 string = v45
                                _ = v59
                            }
                            for fi60 := 0; fi60 < 14; fi60++ {
                                switch (v45 + "a") {
                                case "hello":
                                    v26 = 2
                                case "a":
                                    var v61 uint32 = 198
                                    _ = v61
                                case "foo bar":
                                    v32 += (v32 * (v32 >> ((65535) & 63)))
                                }
                                _ = len(v2)
                                var v62 uint32 = 1
                                v1 = append(v1, (v45 + "a"))
                                if (42 < 1000) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        v63 := GenS1{F0: (v26 & 185)}
                                        _ = v63
                                    case "a":
                                        v64 := v2[v62]
                                        _ = v64
                                    default:
                                        v65 := GenS1{F0: v26}
                                        _ = v65
                                    }
                                } else {
                                    v62 = (^(v62))
                                    fmt.Println(v45, v53)
                                    if (!((-1 > 42) && ((v45 + v45) == v53))) {
                                        switch ("hello" + "\n\t") {
                                        case "hello":
                                            var v66 float64 = 1e10
                                            v67 := GenS4{F0: 18446744073709551615}
                                            _ = len(v1)
                                            var v68 string = "hello"
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        case "a":
                                            fmt.Println(v45)
                                            if (145 == 664) {
                                                if true {
                                                    v2[(^(v62))] = 4294967295
                                                    v69 := v2[65535]
                                                    v2[v69] = (v62 & (v69 & 2))
                                                    v70 := GenS3{F0: 142, F1: v62, F2: v45}
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    switch (fi41 << ((413) & 63)) {
                                                    case 1:
                                                        switch (fi41 - (fi41 + fi41)) {
                                                        case 2147483647:
                                                            if ((65535 > 579) || (0.1 >= (float64((-2.25 / (-2.25 * ((float32(1.5) + 3.14159) + 1.5)))) / (0.1 / 1.0)))) {
                                                                v71 := GenS5{F0: -2147483648, F1: 606}
                                                                v72 := v2[(v62 << ((1) & 63))]
                                                                switch ((("foo bar" + "a") + v45) + v45) {
                                                                case "":
                                                                    fmt.Println(v26)
                                                                case "a":
                                                                    fmt.Println(v45)
                                                                case "hello":
                                                                    v26 -= (v26 * uint8(v3))
                                                                }
                                                                fmt.Println(v32)
                                                                v26 = (v26 << ((65535) & 63))
                                                                if true {
                                                                    switch v45 {
                                                                    case "a":
                                                                        v73 := GenS2{F0: 1000, F1: (v26 < v26)}
                                                                        _ = v73
                                                                    default:
                                                                        var v74 float64 = float64(-2.25)
                                                                        _ = v74
                                                                    }
                                                                    if len(v1) > 0 {
                                                                        v1[0] = ("" + v45)
                                                                    }
                                                                    if (0 != uint16(v3)) {
                                                                        var v75 bool = true
                                                                        if v75 {
                                                                            v72 += (v72 - v72)
                                                                            v3 = (v3 / ((v3) | 1))
                                                                        } else {
                                                                            v45 = ((v45 + "") + (v45 + v45))
                                                                            fmt.Println(fi41)
                                                                            switch (fi60 % (((-(fi60))) | 1)) {
                                                                            case -2147483648:
                                                                                v76 := GenS1{F0: (v26 << ((322) & 63))}
                                                                                _ = v76
                                                                            case 4294967295:
                                                                                var v77 int = (^(fi60))
                                                                                _ = v77
                                                                            }
                                                                            var v78 uint16 = 879
                                                                            _ = v78
                                                                        }
                                                                        var v79 int32 = 0
                                                                        var v80 int32 = (v79 - 1)
                                                                        v72 += (^(v72))
                                                                        _ = v75
                                                                        _ = v79
                                                                        _ = v80
                                                                    } else {
                                                                        if false {
                                                                            v3 -= v3
                                                                        } else {
                                                                            var v81 int32 = 0
                                                                            fmt.Println(fi60, v45, fi60)
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                } else {
                                                                    delete(v2, 718)
                                                                }
                                                                fmt.Println(v3, v62)
                                                                v72 = 1
                                                                fmt.Println(v53, v3, v45)
                                                                _ = v71
                                                                _ = v72
                                                            } else {
                                                                var v82 uint8 = (v26 >> ((65535) & 63))
                                                                v83 := GenS2{F0: 1000, F1: true}
                                                                var v84 uint64 = 1
                                                                _ = v82
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                        case -9223372036854775808:
                                                            v85 := GenS6{F0: 32767, F1: 819}
                                                            _ = v85
                                                        default:
                                                            v86 := GenS7{F0: 0, F1: 118, F2: (0 >= -32768)}
                                                            fmt.Println(fi41, fi41)
                                                            var v87 uint32 = v62
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                    case 128:
                                                        var v88 int = (^(fi41))
                                                        _ = v88
                                                    case -195:
                                                        if (float64(1e10) == ((((((-2.25 * 3.14159) * -2.25) * ((1e10 * 1.5) / 0.1)) / ((1e10 * float64(0.1)) * -2.25)) - 0.1) / ((1e10 + ((0.1 / (0.1 / (3.14159 + 3.14159))) / -2.25)) - 3.14159))) {
                                                            _ = len(v1)
                                                        }
                                                    }
                                                    v89 := GenS8{F0: -1, F1: 100, F2: false}
                                                    switch (^(fi60)) {
                                                    case -9223372036854775808:
                                                        fmt.Println(v3, fi60)
                                                    }
                                                    _ = v89
                                                }
                                                v3 = (v3 & 42)
                                            } else {
                                                var v90 int32 = 0
                                                v91 := GenS9{F0: 127}
                                                _ = v90
                                                _ = v91
                                            }
                                            var v92 int16 = 172
                                            v45 = v53
                                            v93 := GenS10{F0: v45, F1: 32767}
                                            _ = v92
                                            _ = v93
                                        default:
                                            fmt.Println(v32, fi60)
                                        }
                                        v62 = (^(v62))
                                        var v94 int = (^(fi41))
                                        _ = v94
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = (v53 + ("\n\t" + v45))
                                }
                                switch (v53 + v53) {
                                case "":
                                    if ((0 <= 0) || true) {
                                        v2[(v62 ^ v62)] = v62
                                        _ = len(v2)
                                        v95 := GenS4{F0: 116}
                                        _ = v95
                                    }
                                default:
                                    var v96 float64 = 3.14159
                                    _ = v96
                                }
                                _ = fi60
                                _ = v62
                            }
                            _ = v53
                        }
                    default:
                        switch (fi41 & fi41) {
                        case 54:
                            v26 = (v26 << ((18446744073709551615) & 63))
                        default:
                            v45 = (v45 + v45)
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi41
                _ = v45
            }
        default:
            v3 -= 354
        }
    default:
        for fi97 := 0; fi97 < 19; fi97++ {
            var v98 string = "foo bar"
            v98 = ("" + "\n\t")
            _ = fi97
            _ = v98
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v32
}
