package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 int; F1 int; F2 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int32; F2 uint64 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 uint16; F1 uint32 }
type GenS12 struct { F0 bool; F1 int64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int; F1 int }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint; F1 uint64; F2 int16 }
func genFn0(p0 uint) uint64 {
    var v5 uint16 = 2
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint16 = v5
        var v8 int8 = 85
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return 617
}

func main() {
    var v1 uint16 = 0
    var v2 bool = true
    var v3 float64 = -2.25
    v4 := make(map[uint32]bool)
    delete(v4, 0)
    v2 = true
    switch "" {
    case "":
        v3 = (3.14159 + 0.0)
    case "foo bar":
        v2 = v2
    }
    v3 = v3
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v1)
        _ = fi9
    }
    fmt.Println(v2)
    v10 := GenS1{F0: -578, F1: 32767, F2: ""}
    fmt.Println(v3, v3, v1)
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS2{F0: 2, F1: ""}
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 float64 = v3
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            if true {
                var v16 string = "a"
                switch "hello" {
                case "foo bar":
                    var v17 uint64 = 42
                    _ = v17
                case "a":
                    var v18 string = ("foo bar" + "")
                    _ = v18
                default:
                    if true {
                        v19 := GenS3{F0: 1, F1: -282, F2: -128}
                        var v20 uint16 = v1
                        _ = v19
                        _ = v20
                    } else {
                        v21 := []int{-9223372036854775808, fi15, fi11}
                        v22 := GenS3{F0: 799, F1: 16, F2: 12}
                        var v23 int16 = -32768
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                }
                v24 := GenS4{F0: fi11, F1: fi15, F2: -1}
                _ = v16
                _ = v24
            }
            _ = fi15
        }
        _ = fi11
        _ = v12
    }
    v2 = ((v2 && (v2 && false)) && (32767 < 128))
    var v25 float64 = 0.0
    v26 := v4[137]
    v2 = (1 <= 133)
    switch -1 {
    case -684:
        switch (("a" + ("a" + "")) + "") {
        case "a":
            v25 = (v25 * v3)
        }
    case 32768:
        delete(v4, 23)
    case -2147483648:
        v27 := GenS5{F0: 127}
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := GenS1{F0: -9223372036854775808, F1: 640, F2: "hello"}
        v3 = (0.0 / -2.25)
        _ = fi28
        _ = v29
    }
    if false {
        if v2 {
            if v26 {
                if (!(((0.1 / 3.14159) * 1.5) >= -2.25)) {
                    switch int(v1) {
                    case 9223372036854775807:
                        v26 = (1 <= -1)
                    case -1:
                        for fi30 := 0; fi30 < 18; fi30++ {
                            var v31 int = fi30
                            fmt.Println(v26, fi30, v31)
                            fmt.Println(v3, v31)
                            _ = fi30
                            _ = v31
                        }
                    }
                    if v26 {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            var v33 float64 = v3
                            if (!(59 < 114)) {
                                v33 += 1e10
                                var v34 uint = 18446744073709551615
                                var v35 float64 = float64(3.14159)
                                _ = v34
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 13; fi36++ {
                                    v25 += (v3 - v3)
                                    v1 = v1
                                    switch (fi36 * (fi36 ^ fi36)) {
                                    case -2147483648:
                                        switch "" {
                                        case "a":
                                            switch fi36 {
                                            case 127:
                                                if v2 {
                                                    if (false && (!(2147483647 < -2147483648))) {
                                                        switch ("" + (((("a" + "\n\t") + "foo bar") + ("hello" + "hello")) + "")) {
                                                        case "a":
                                                            if v26 {
                                                                fmt.Println(v25)
                                                                if ((float64(0.0) - v33) >= 0.0) {
                                                                    v37 := []int{fi32, fi36, (^(fi36)), fi32}
                                                                    var v38 uint32 = 248
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                                if v26 {
                                                                    v39 := v4[2147483647]
                                                                    var v40 uint32 = 188
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                var v41 float64 = v3
                                                                var v42 bool = true
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                        case "":
                                                            if false {
                                                                if (0.1 != (1e10 * ((-2.25 / ((0.1 + 1e10) * 1e10)) * 1e10))) {
                                                                    v2 = true
                                                                    if (1000 == 251) {
                                                                        var v43 string = (("a" + "") + ("hello" + "a"))
                                                                        if (v26 && v2) {
                                                                            v43 = v43
                                                                            v33 -= 3.14159
                                                                        }
                                                                        _ = v43
                                                                    } else {
                                                                        v33 -= (1.5 * (1e10 - v25))
                                                                        var v44 uint32 = 687
                                                                        _ = v44
                                                                    }
                                                                    switch "" {
                                                                    case "a":
                                                                        if (!false) {
                                                                            var v45 int8 = 42
                                                                            _ = v45
                                                                        } else {
                                                                            v46 := GenS6{F0: 674, F1: -2147483648, F2: 4294967295}
                                                                            delete(v4, 65535)
                                                                            _ = v46
                                                                        }
                                                                    case "foo bar":
                                                                        var v47 uint64 = 4294967295
                                                                        _ = v47
                                                                    case "\n\t":
                                                                        switch "a" {
                                                                        case "":
                                                                            if (42 <= 255) {
                                                                                var v48 int32 = -1
                                                                                _ = v48
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(v1)
                                                                        case "a":
                                                                            var v49 bool = ((((42 == 1) && (!(v26 && v2))) || false) || false)
                                                                            _ = v49
                                                                        default:
                                                                            v50 := v4[65535]
                                                                            _ = v50
                                                                        }
                                                                    default:
                                                                        var v51 int32 = 42
                                                                        fmt.Println(v3, fi32, v51)
                                                                        _ = v51
                                                                    }
                                                                }
                                                                v3 = v3
                                                                switch 32768 {
                                                                case -9223372036854775808:
                                                                    var v52 int16 = 0
                                                                    _ = v52
                                                                case 725:
                                                                    v25 += -2.25
                                                                default:
                                                                    v53 := GenS1{F0: 42, F1: 1000, F2: ""}
                                                                    _ = v53
                                                                }
                                                                var v54 uint32 = 42
                                                                v55 := GenS4{F0: fi36, F1: fi36, F2: 100}
                                                                _ = v54
                                                                _ = v55
                                                            } else {
                                                                v56 := GenS7{F0: 0, F1: (v1 & (^(v1))), F2: 42}
                                                                switch fi36 {
                                                                case 4294967295:
                                                                    var v57 int8 = -89
                                                                    _ = v57
                                                                default:
                                                                    v58 := GenS5{F0: -1}
                                                                    _ = v58
                                                                }
                                                                v59 := v4[65535]
                                                                v60 := GenS3{F0: 42, F1: 2147483647, F2: 63}
                                                                v61 := GenS5{F0: 0}
                                                                _ = v56
                                                                _ = v59
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        default:
                                                            v62 := []int{}
                                                            _ = v62
                                                        }
                                                        _ = len(v4)
                                                    }
                                                    if false {
                                                        v1 += 957
                                                    } else {
                                                        v63 := GenS8{F0: 546, F1: v26}
                                                        _ = v63
                                                    }
                                                    switch fi32 {
                                                    case -1:
                                                        switch ("hello" + "a") {
                                                        case "a":
                                                            v64 := GenS9{F0: 1, F1: fi36}
                                                            _ = v64
                                                        }
                                                    case 4294967295:
                                                        v65 := GenS3{F0: -1, F1: 1, F2: 1}
                                                        _ = v65
                                                    default:
                                                        fmt.Println(fi36, v3)
                                                    }
                                                }
                                            default:
                                                v25 = v25
                                            }
                                        case "foo bar":
                                            v3 += (v3 / v25)
                                        case "\n\t":
                                            v66 := GenS1{F0: 42, F1: 42, F2: "hello"}
                                            _ = v66
                                        }
                                        _ = len(v4)
                                        fmt.Println(v33, v26)
                                    case -27:
                                        if v26 {
                                            v67 := GenS10{F0: fi32, F1: 53}
                                            var v68 string = ""
                                            var v69 uint64 = 18446744073709551615
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        }
                                    default:
                                        fmt.Println(v1, v3)
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        if v2 {
                                            v70 := GenS11{F0: v1, F1: 2147483647}
                                            _ = v70
                                        } else {
                                            var v71 uint32 = 2147483648
                                            v2 = true
                                            var v72 uint32 = (v71 | (v71 % (((v71 << ((1) & 63))) | 1)))
                                            _ = v71
                                            _ = v72
                                        }
                                    case "a":
                                        var v73 int16 = 0
                                        _ = v73
                                    }
                                    _ = fi36
                                }
                            }
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        if v26 {
                            var v74 int64 = -1
                            v75 := GenS2{F0: 2, F1: ("a" + "\n\t")}
                            _ = v74
                            _ = v75
                        }
                        v1 = v1
                        if v26 {
                            if v26 {
                                var v76 int8 = -94
                                for fi77 := 0; fi77 < 20; fi77++ {
                                    v78 := GenS12{F0: (370 > 948), F1: -9223372036854775808}
                                    fmt.Println(v3)
                                    _ = fi77
                                    _ = v78
                                }
                                _ = v76
                            }
                            v2 = (true || v26)
                        } else {
                            if v26 {
                                fmt.Println(v26, v1)
                                v2 = v26
                            } else {
                                v79 := GenS2{F0: 0, F1: (("hello" + "\n\t") + "")}
                                v80 := v4[4294967295]
                                _ = v79
                                _ = v80
                            }
                            for fi81 := 0; fi81 < 19; fi81++ {
                                fmt.Println(v2)
                                for fi82 := 0; fi82 < 4; fi82++ {
                                    fmt.Println(v25)
                                    _ = fi82
                                }
                                var v83 float32 = (1e10 * (-2.25 + 3.14159))
                                _ = fi81
                                _ = v83
                            }
                        }
                    }
                } else {
                    var v84 int = 45
                    _ = v84
                }
                for fi85 := 0; fi85 < 6; fi85++ {
                    fmt.Println(v25, fi85)
                    v86 := GenS13{F0: -1}
                    for fi87 := 0; fi87 < 11; fi87++ {
                        var v88 float32 = float32(v25)
                        _ = fi87
                        _ = v88
                    }
                    _ = fi85
                    _ = v86
                }
            }
            if true {
                v1 = (v1 | 110)
            } else {
                v26 = (2 == genFn0(601))
                for fi89 := 0; fi89 < 2; fi89++ {
                    v90 := v4[4294967295]
                    v91 := GenS14{F0: (fi89 >> ((255) & 63)), F1: (fi89 ^ fi89)}
                    if (((true && v2) || v2) || v26) {
                        if (false || ((!false) || (false || (9223372036854775807 != 0)))) {
                            if (345 >= 42) {
                                v92 := []int{}
                                var v93 float32 = (3.14159 - 1e10)
                                if ((!(v93 > v93)) || ((!v2) || true)) {
                                    v94 := GenS13{F0: -970}
                                    _ = v94
                                }
                                _ = v92
                                _ = v93
                            } else {
                                for fi95 := 0; fi95 < 20; fi95++ {
                                    var v96 uint16 = v1
                                    if v90 {
                                        var v97 uint16 = (^(v96))
                                        _ = v97
                                    }
                                    v98 := []int{(fi89 << ((255) & 63)), fi89, (fi95 >> ((uint(v96)) & 63)), 4294967295}
                                    _ = fi95
                                    _ = v96
                                    _ = v98
                                }
                            }
                        }
                    }
                    _ = fi89
                    _ = v90
                    _ = v91
                }
            }
            for fi99 := 0; fi99 < 4; fi99++ {
                if (((("foo bar" + "a") + "a") + "hello") == "hello") {
                    var v100 int = (fi99 << ((255) & 63))
                    for fi101 := 0; fi101 < 7; fi101++ {
                        var v102 uint8 = 1
                        _ = fi101
                        _ = v102
                    }
                    _ = v100
                } else {
                    if (4294967295 >= genFn0(4294967295)) {
                        v103 := GenS15{F0: 18446744073709551615}
                        var v104 uint = 65535
                        _ = v103
                        _ = v104
                    } else {
                        v26 = (32767 < 42)
                        for fi105 := 0; fi105 < 11; fi105++ {
                            var v106 bool = (127 < -1)
                            v25 += (0.1 * 1e10)
                            switch (fi99 | (fi99 & fi99)) {
                            case 32767:
                                switch ("a" + "") {
                                case "foo bar":
                                    v107 := GenS9{F0: 900, F1: (fi105 / ((fi99) | 1))}
                                    _ = v107
                                case "hello":
                                    if (127 != int8(fi99)) {
                                        var v108 int64 = -700
                                        var v109 float32 = (1e10 * 3.14159)
                                        _ = v108
                                        _ = v109
                                    }
                                }
                            case -9223372036854775808:
                                _ = len(v4)
                                fmt.Println(v3)
                            case 128:
                                var v110 int32 = int32(v1)
                                _ = v110
                            default:
                                v2 = (!(fi99 > (fi99 / (((fi99 - (fi105 * fi105))) | 1))))
                            }
                            _ = fi105
                            _ = v106
                        }
                        var v111 int8 = -128
                        _ = v111
                    }
                    v112 := GenS16{F0: 1, F1: 1, F2: 42}
                    var v113 string = "hello"
                    _ = v112
                    _ = v113
                }
                _ = fi99
            }
        }
    } else {
        if ((!false) && v2) {
            v114 := []int{0, -492, 447}
            _ = len(v114)
            fmt.Println(v2, v3)
            _ = v114
        }
        switch 128 {
        case 151:
            var v115 string = ("a" + "foo bar")
            _ = v115
        case -2147483648:
            fmt.Println(v2, v26, v25)
        default:
            if ((((0.0 + 1.5) - 0.0) * float32(-2.25)) <= 1.5) {
                v116 := []int{2147483647}
                _ = len(v4)
                _ = v116
            } else {
                var v117 uint32 = 65535
                if false {
                    v25 = v3
                }
                var v118 bool = (!(217 == 42))
                _ = v117
                _ = v118
            }
        }
    }
    fmt.Println(v3, v25, v2)
    v2 = (18446744073709551615 < 255)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
