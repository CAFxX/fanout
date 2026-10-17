package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 bool; F1 bool; F2 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 int }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint32; F1 uint32; F2 int32 }
func genFn0(p0 int32) uint32 {
    switch "hello" {
    case "hello":
        var v5 uint8 = 255
        _ = v5
    default:
        v6 := GenS1{F0: "a", F1: 1}
        _ = v6
    }
    fmt.Println(p0)
    var v7 uint16 = 2
    _ = p0
    _ = v7
    return 783
}
func genFn1(p0 uint8, p1 int8, p2 uint64) uint16 {
    if true {
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 = p0
            var v9 int16 = -32768
            v10 := []int{}
            _ = fi8
            _ = v10
            _ = v9
        }
        p1 -= 87
        p1 = 42
    } else {
        var v11 uint = 611
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint16 = 65535
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 bool = (("a" + "hello") != "foo bar")
    var v2 uint16 = 199
    v3 := make(map[string]int)
    var v4 bool = (0 <= 42)
    v14 := []int{32767, 147, 128, 127}
    switch -652 {
    case -458:
        fmt.Println(v4, v1)
    case -9223372036854775808:
        v1 = true
    case 860:
        var v15 float32 = 1.5
        _ = v15
    default:
        v1 = (1 > genFn0(0))
    }
    var v16 int8 = -1
    fmt.Println(v2, v4, v1)
    if false {
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                if v4 {
                    var v19 uint32 = genFn0(1)
                    _ = v19
                }
                switch -10 {
                case 1:
                    if v4 {
                        if v1 {
                            fmt.Println(v16, fi17, fi17)
                            v1 = false
                            if (v4 || v1) {
                                v20 := GenS2{F0: -32768}
                                if (((v1 || (v1 || (!v1))) && (v4 || false)) && false) {
                                    var v21 int64 = -1
                                    _ = v21
                                } else {
                                    if true {
                                        v22 := GenS1{F0: "a", F1: 159}
                                        v1 = ((false && (345 != -847)) || v1)
                                        _ = v22
                                    } else {
                                        v23 := GenS3{F0: 921}
                                        v24 := GenS4{F0: 490, F1: 2}
                                        if (true || (!v4)) {
                                            var v25 uint8 = 42
                                            if true {
                                                var v26 int16 = 42
                                                v4 = false
                                                _ = v26
                                            } else {
                                                var v27 uint32 = 4294967295
                                                v1 = (!false)
                                                v28 := GenS5{F0: (fi18 - int(v16))}
                                                v2 -= v2
                                                _ = v27
                                                _ = v28
                                            }
                                            _ = v25
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                    var v29 int = (fi17 >> ((uint(v2)) & 63))
                                    _ = v29
                                }
                                var v30 float64 = float64(0.0)
                                if (v1 && v4) {
                                    v31 := GenS2{F0: 0}
                                    _ = v31
                                }
                                _ = v20
                                _ = v30
                            } else {
                                switch ("" + "foo bar") {
                                case "":
                                    v4 = true
                                    var v32 bool = v4
                                    _ = v32
                                case "\n\t":
                                    _ = len(v3)
                                default:
                                    if v1 {
                                        switch "hello" {
                                        case "a":
                                            var v33 string = "foo bar"
                                            _ = v33
                                        case "\n\t":
                                            switch (fi18 * 2147483647) {
                                            case 9223372036854775807:
                                                v4 = ((true || (true || (3.14159 != ((0.1 * -2.25) + (3.14159 / 1e10))))) && false)
                                            case -1:
                                                var v34 bool = (!(v1 && v4))
                                                _ = v34
                                            case 32767:
                                                if true {
                                                    var v35 int32 = -994
                                                    v35 -= (v35 << ((0) & 63))
                                                    _ = v35
                                                } else {
                                                    var v36 bool = v1
                                                    _ = v36
                                                }
                                            }
                                        default:
                                            fmt.Println(v4, v4)
                                        }
                                        _ = len(v14)
                                    } else {
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                }
                                v37 := GenS6{F0: -969, F1: v2}
                                v38 := v3["\n\t"]
                                if (v1 && (!v1)) {
                                    var v39 bool = ("a" > "hello")
                                    _ = v39
                                }
                                _ = v37
                                _ = v38
                            }
                        }
                    } else {
                        if false {
                            switch (^(fi18)) {
                            case -1:
                                if v4 {
                                    if (!(v16 < v16)) {
                                        delete(v3, ("foo bar" + ("foo bar" + "a")))
                                    } else {
                                        var v40 int32 = 761
                                        _ = v40
                                    }
                                    if len(v14) > 0 {
                                        v14[0] = (fi17 ^ (fi18 * 42))
                                    }
                                    fmt.Println(fi17, fi18, fi18)
                                } else {
                                    v41 := GenS4{F0: 39, F1: 719}
                                    _ = v41
                                }
                            case 680:
                                v42 := GenS3{F0: 42}
                                _ = v42
                            case -872:
                                v3[""] = fi17
                            }
                        }
                        fmt.Println(v4)
                        var v43 bool = ("\n\t" >= ("a" + "hello"))
                        _ = v43
                    }
                case 127:
                    v44 := v3["hello"]
                    _ = v44
                case -1:
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                if len(v14) > 0 {
                    v14[0] = fi17
                }
                _ = fi18
            }
            _ = fi17
        }
    }
    switch 194 {
    case 128:
        switch "" {
        case "\n\t":
            v2 += (^(v2))
        case "a":
            var v45 int32 = -782
            _ = v45
        case "hello":
            if ((3.14159 - -2.25) == 0.1) {
                for fi46 := 0; fi46 < 19; fi46++ {
                    switch (fi46 << ((2) & 63)) {
                    case 147:
                        _ = len(v14)
                    case 460:
                        if true {
                            v47 := GenS7{F0: (-1 >= 2147483647), F1: (128 <= 42), F2: ((v2 > (v2 & 255)) || true)}
                            _ = v47
                        } else {
                            v48 := GenS8{F0: 18446744073709551615}
                            _ = v48
                        }
                    default:
                        var v49 uint64 = 1
                        _ = v49
                    }
                    fmt.Println(v16)
                    if (42 <= 560) {
                        for fi50 := 0; fi50 < 18; fi50++ {
                            v51 := GenS2{F0: 42}
                            v14 = append(v14, fi46)
                            _ = fi50
                            _ = v51
                        }
                    }
                    _ = fi46
                }
                for fi52 := 0; fi52 < 12; fi52++ {
                    for fi53 := 0; fi53 < 20; fi53++ {
                        var v54 uint8 = 2
                        switch 32768 {
                        case 32767:
                            if (uint16(v54) != 357) {
                                var v55 uint16 = (v2 * (v2 % ((genFn1((v54 / (((v54 - 255)) | 1)), (v16 - v16), 18446744073709551615)) | 1)))
                                var v56 float32 = -2.25
                                v1 = (570 != 1)
                                _ = v55
                                _ = v56
                            } else {
                                var v57 float32 = (float32(3.14159) / 0.1)
                                switch "foo bar" {
                                case "a":
                                    v57 += (v57 - v57)
                                case "":
                                    fmt.Println(v4, fi53)
                                case "\n\t":
                                    switch fi52 {
                                    case 127:
                                        var v58 float32 = v57
                                        _ = v58
                                    default:
                                        v4 = (!v1)
                                    }
                                }
                                var v59 uint64 = 18446744073709551615
                                _ = v57
                                _ = v59
                            }
                            var v60 uint64 = 2
                            v1 = false
                            v2 = v2
                            v60 = 300
                            _ = v60
                        case -1:
                            v54 = (v54 / (((v54 >> ((65535) & 63))) | 1))
                        case 1:
                            v2 += v2
                        default:
                            var v61 int32 = 2147483647
                            _ = v61
                        }
                        var v62 uint8 = (v54 ^ 25)
                        _ = fi53
                        _ = v54
                        _ = v62
                    }
                    _ = fi52
                }
            }
        }
    case 4294967295:
        v4 = false
    case 0:
        var v63 uint8 = 2
        _ = v63
    default:
        v64 := GenS5{F0: 42}
        _ = v64
    }
    fmt.Println(v2, v2)
    for fi65 := 0; fi65 < 16; fi65++ {
        var v66 uint32 = uint32(v2)
        _ = fi65
        _ = v66
    }
    _ = len(v14)
    for fi67 := 0; fi67 < 17; fi67++ {
        v3["a"] = (^(fi67))
        switch "hello" {
        case "a":
            _ = len(v14)
        case "\n\t":
            var v68 uint64 = 0
            _ = v68
        default:
            var v69 int = fi67
            _ = v69
        }
        _ = fi67
    }
    if v4 {
        v2 = 65535
        v70 := GenS2{F0: 1}
        _ = v70
    } else {
        if true {
            for fi71 := 0; fi71 < 7; fi71++ {
                delete(v3, "a")
                switch "hello" {
                case "a":
                    for fi72 := 0; fi72 < 14; fi72++ {
                        var v73 float32 = 3.14159
                        _ = fi72
                        _ = v73
                    }
                case "foo bar":
                    var v74 string = "hello"
                    _ = v74
                case "hello":
                    if len(v14) > 0 {
                        v14[0] = -979
                    }
                }
                var v75 uint8 = 128
                _ = fi71
                _ = v75
            }
        } else {
            v16 -= 0
        }
        v3["\n\t"] = -694
    }
    v76 := GenS9{F0: "\n\t", F1: 128}
    var v77 uint = 303
    if ((!(v4 && v4)) && (1e10 != (0.1 / (1e10 + (0.1 - 0.0))))) {
        if v1 {
            v16 = v16
            v78 := GenS10{F0: v77, F1: 0}
            for fi79 := 0; fi79 < 11; fi79++ {
                for fi80 := 0; fi80 < 1; fi80++ {
                    v81 := GenS6{F0: 9, F1: 12}
                    fmt.Println(v77, v1)
                    if (!(1e10 <= (float32(0.1) + 3.14159))) {
                        var v82 uint64 = 0
                        fmt.Println(fi79)
                        _ = v82
                    }
                    _ = fi80
                    _ = v81
                }
                var v83 string = "a"
                _ = fi79
                _ = v83
            }
            _ = v78
        }
        if (v16 == v16) {
            if (42 == 2) {
                var v84 int16 = -32768
                if (!(!true)) {
                    fmt.Println(v1, v84)
                    for fi85 := 0; fi85 < 17; fi85++ {
                        _ = len(v3)
                        var v86 string = ("\n\t" + ("hello" + "hello"))
                        var v87 string = ((((v86 + ("hello" + ((v86 + "\n\t") + (v86 + v86)))) + v86) + "") + v86)
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                    v1 = (false || true)
                } else {
                    for fi88 := 0; fi88 < 19; fi88++ {
                        _ = len(v14)
                        _ = fi88
                    }
                    for fi89 := 0; fi89 < 6; fi89++ {
                        v90 := GenS3{F0: 0}
                        var v91 float32 = 3.14159
                        for fi92 := 0; fi92 < 14; fi92++ {
                            v3["hello"] = (^(fi89))
                            var v93 int32 = 42
                            _ = len(v14)
                            v91 = 0.0
                            _ = fi92
                            _ = v93
                        }
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                }
                _ = v84
            }
        }
    } else {
        v94 := v3["foo bar"]
        var v95 uint = v77
        _ = v94
        _ = v95
    }
    fmt.Println(v2, v4, v2)
    var v96 float32 = 1e10
    for fi97 := 0; fi97 < 16; fi97++ {
        v98 := GenS5{F0: fi97}
        switch "hello" {
        case "\n\t":
            for fi99 := 0; fi99 < 9; fi99++ {
                _ = len(v14)
                if (true && true) {
                    switch "hello" {
                    case "hello":
                        if v4 {
                            v100 := GenS11{F0: -1}
                            _ = v100
                        } else {
                            var v101 bool = (1e10 <= -2.25)
                            _ = v101
                        }
                    }
                    v102 := GenS6{F0: -477, F1: v2}
                    v14 = append(v14, (-(fi99)))
                    _ = v102
                } else {
                    var v103 uint64 = 42
                    v104 := GenS2{F0: 0}
                    if v4 {
                        var v105 bool = v1
                        _ = v105
                    }
                    _ = v103
                    _ = v104
                }
                if v4 {
                    if (!false) {
                        switch "" {
                        case "a":
                            v4 = v1
                        case "":
                            switch "a" {
                            case "hello":
                                fmt.Println(fi99)
                            }
                        case "\n\t":
                            var v106 uint8 = 108
                            _ = v106
                        default:
                            switch "foo bar" {
                            case "a":
                                v1 = (42 == 32767)
                            case "\n\t":
                                var v107 float64 = 1.5
                                _ = v107
                            default:
                                fmt.Println(fi97, fi99)
                            }
                        }
                        var v108 float32 = 1.5
                        if v1 {
                            _ = len(v3)
                        } else {
                            var v109 int32 = 0
                            v110 := GenS12{F0: 42, F1: 65535, F2: (v109 >> ((v77) & 63))}
                            var v111 int64 = 869
                            _ = v109
                            _ = v110
                            _ = v111
                        }
                        _ = v108
                    } else {
                        var v112 int32 = 2147483647
                        v113 := GenS1{F0: "a", F1: (v77 % ((v77) | 1))}
                        _ = v112
                        _ = v113
                    }
                    fmt.Println(fi97, v77, v2)
                    switch ("a" + ("hello" + "")) {
                    case "\n\t":
                        var v114 uint64 = 4294967295
                        delete(v3, "a")
                        _ = v114
                    default:
                        if v4 {
                            var v115 int32 = -1
                            if v1 {
                                var v116 uint8 = 164
                                var v117 uint16 = (v2 >> ((v77) & 63))
                                if ((v116 % ((v116) | 1)) <= v116) {
                                    var v118 int = (fi99 >> (((^(v77))) & 63))
                                    _ = v118
                                }
                                fmt.Println(v1, v115, fi97)
                                v77 = (v77 & 45)
                                _ = v116
                                _ = v117
                            }
                            fmt.Println(fi99, v77)
                            _ = v115
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                switch fi97 {
                                case 127:
                                    if (("a" + "hello") == "") {
                                        var v119 bool = v1
                                        v16 = v16
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        _ = v119
                                    }
                                case -858:
                                    switch fi97 {
                                    case 4294967295:
                                        v1 = false
                                    case 128:
                                        var v120 uint8 = 127
                                        _ = v120
                                    case 840:
                                        if v1 {
                                            if v4 {
                                                if v4 {
                                                    var v121 uint16 = v2
                                                    _ = v121
                                                }
                                                v16 = v16
                                            } else {
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                                v122 := GenS5{F0: 0}
                                                var v123 int32 = 1
                                                _ = v122
                                                _ = v123
                                            }
                                            _ = len(v3)
                                        } else {
                                            var v124 uint16 = genFn1(67, v16, 18446744073709551615)
                                            var v125 float64 = 0.1
                                            _ = v124
                                            _ = v125
                                        }
                                    }
                                case 9223372036854775807:
                                    fmt.Println(fi97, v1)
                                }
                            case "hello":
                                if (!(2 == 65535)) {
                                    v96 += v96
                                }
                            case "":
                                v2 += 307
                            }
                            var v126 int64 = -9223372036854775808
                            v1 = (genFn0(1) < genFn0(-607))
                            _ = v126
                        }
                    }
                }
                v1 = ((-1 < 42) && true)
                fmt.Println(v96, fi97, v2)
                v127 := GenS4{F0: 777, F1: 2147483647}
                _ = fi99
                _ = v127
            }
        }
        if v4 {
            var v128 int8 = v16
            var v129 string = "hello"
            _ = v128
            _ = v129
        } else {
            v16 -= v16
        }
        _ = fi97
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v77:", v77)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v76
    _ = v77
    _ = v96
}
