package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32; F1 int64; F2 uint32 }
type GenS11 struct { F0 int16; F1 int }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 uint8; F1 int8 }
type GenS14 struct { F0 uint; F1 bool }
type GenS15 struct { F0 int64 }

func main() {
    var v1 uint16 = 242
    var v2 uint16 = (^(v1))
    v3 := make(map[string]int)
    var v4 bool = false
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int64 = 278
        _ = fi5
        _ = v6
    }
    v3["a"] = -404
    if (!v4) {
        v7 := GenS1{F0: (v1 << ((384) & 63))}
        delete(v3, "foo bar")
        _ = v7
    } else {
        if v4 {
            fmt.Println(v1, v1)
            for fi8 := 0; fi8 < 3; fi8++ {
                v2 = 42
                _ = fi8
            }
            var v9 uint64 = 42
            _ = v9
        } else {
            if v4 {
                var v10 float64 = ((0.0 + 3.14159) * (-2.25 * (1e10 / -2.25)))
                _ = v10
            } else {
                switch 32768 {
                case -543:
                    for fi11 := 0; fi11 < 7; fi11++ {
                        for fi12 := 0; fi12 < 12; fi12++ {
                            switch (("" + "foo bar") + "") {
                            case "foo bar":
                                var v13 uint16 = v2
                                _ = v13
                            default:
                                var v14 int8 = 1
                                _ = v14
                            }
                            var v15 int16 = 42
                            v1 = (v2 ^ v1)
                            _ = fi12
                            _ = v15
                        }
                        _ = fi11
                    }
                case 4294967295:
                    v4 = v4
                default:
                    var v16 uint = 779
                    _ = v16
                }
                for fi17 := 0; fi17 < 8; fi17++ {
                    fmt.Println(fi17, fi17)
                    _ = fi17
                }
            }
            v4 = (!v4)
        }
        var v18 uint8 = 139
        v2 = 1
        _ = v18
    }
    var v19 float64 = 1e10
    if v4 {
        for fi20 := 0; fi20 < 9; fi20++ {
            if false {
                var v21 uint16 = v2
                _ = v21
            } else {
                switch (fi20 - fi20) {
                case 455:
                    if v4 {
                        if true {
                            v4 = (-1 != (fi20 << ((1) & 63)))
                            var v22 float64 = v19
                            var v23 bool = true
                            _ = v22
                            _ = v23
                        }
                        if (409 >= 42) {
                            switch fi20 {
                            case 128:
                                var v24 float32 = 0.0
                                _ = v24
                            case 32767:
                                v4 = (!v4)
                            case -274:
                                for fi25 := 0; fi25 < 6; fi25++ {
                                    v26 := GenS1{F0: (^(v2))}
                                    if false {
                                        v2 -= 2
                                    } else {
                                        v19 = (v19 - (v19 + v19))
                                        v19 = v19
                                        v1 = v2
                                    }
                                    var v27 bool = true
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            }
                            delete(v3, "foo bar")
                            v28 := GenS2{F0: -1}
                            _ = v28
                        }
                        if false {
                            var v29 int = (^(fi20))
                            switch "foo bar" {
                            case "hello":
                                v19 += 0.1
                            case "\n\t":
                                fmt.Println(v2, v1)
                            default:
                                if (!v4) {
                                    if v4 {
                                        if (!v4) {
                                            v1 = 42
                                        } else {
                                            fmt.Println(v4, v1)
                                            for fi30 := 0; fi30 < 4; fi30++ {
                                                fmt.Println(v19, v1)
                                                fmt.Println(v2)
                                                v4 = (v4 && (1 != 255))
                                                _ = fi30
                                            }
                                        }
                                    }
                                    if (fi20 == (fi20 ^ -2147483648)) {
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v19, fi20, v1)
                                        case "":
                                            var v31 uint8 = 2
                                            _ = v31
                                        }
                                        for fi32 := 0; fi32 < 14; fi32++ {
                                            v19 = 0.0
                                            switch fi20 {
                                            case 128:
                                                switch ((("foo bar" + "") + "a") + "hello") {
                                                case "hello":
                                                    v33 := GenS3{F0: 2, F1: 1, F2: 2}
                                                    _ = v33
                                                case "\n\t":
                                                    switch -2147483648 {
                                                    case -1:
                                                        v34 := GenS4{F0: 2, F1: 2}
                                                        _ = v34
                                                    case 956:
                                                        var v35 uint64 = 4294967295
                                                        _ = v35
                                                    case 4294967295:
                                                        delete(v3, "")
                                                    }
                                                default:
                                                    var v36 uint64 = 4294967295
                                                    _ = v36
                                                }
                                            case 32768:
                                                if v4 {
                                                    var v37 int32 = -1
                                                    v38 := GenS1{F0: (v1 >> ((1) & 63))}
                                                    fmt.Println(v19, fi20, v4)
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    switch (fi20 >> ((18446744073709551615) & 63)) {
                                                    case -1:
                                                        v19 -= (((((v19 + v19) / 1.5) - v19) / -2.25) / v19)
                                                    case 9223372036854775807:
                                                        v1 = v1
                                                    case 4294967295:
                                                        v39 := GenS5{F0: -200}
                                                        _ = v39
                                                    default:
                                                        v2 += 61
                                                    }
                                                    switch v29 {
                                                    case 127:
                                                        if v4 {
                                                            fmt.Println(fi32, v1, v1)
                                                            if ((v19 - (v19 * -2.25)) > v19) {
                                                                v19 = v19
                                                                v40 := []int{}
                                                                _ = v40
                                                            } else {
                                                                var v41 int32 = 317
                                                                v1 -= (v2 / ((v2) | 1))
                                                                _ = v41
                                                            }
                                                            if (2147483648 != 2147483648) {
                                                                var v42 uint = 42
                                                                _ = v42
                                                            }
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v4 = false
                                                            case "\n\t":
                                                                v29 -= (fi20 % (((fi32 * (fi20 % ((v29) | 1)))) | 1))
                                                            case "hello":
                                                                switch fi32 {
                                                                case 128:
                                                                    if v4 {
                                                                        v43 := GenS2{F0: 127}
                                                                        v1 = v1
                                                                        v2 = v1
                                                                        v44 := GenS6{F0: v2, F1: 4294967295, F2: 282}
                                                                        _ = v43
                                                                        _ = v44
                                                                    } else {
                                                                        v45 := GenS1{F0: v1}
                                                                        _ = v45
                                                                    }
                                                                }
                                                            }
                                                            v19 -= v19
                                                        }
                                                    case 32767:
                                                        var v46 uint32 = 1
                                                        _ = v46
                                                    case 0:
                                                        v2 = v2
                                                        v47 := GenS6{F0: v2, F1: 2147483647, F2: 376}
                                                        _ = v47
                                                    default:
                                                        v48 := GenS6{F0: (v2 % (((v1 | (v1 ^ v1))) | 1)), F1: 561, F2: 18446744073709551615}
                                                        v1 += v2
                                                        _ = v48
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v49 int16 = 585
                                                        _ = v49
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v50 float32 = (float32(float64(3.14159)) + 1.5)
                                                            _ = v50
                                                        case "hello":
                                                            fmt.Println(v4, v4)
                                                        case "foo bar":
                                                            var v51 uint16 = 32767
                                                            _ = v51
                                                        }
                                                    case "foo bar":
                                                        v4 = (((-(fi32)) != fi20) && false)
                                                    default:
                                                        var v52 int = fi32
                                                        _ = v52
                                                    }
                                                    fmt.Println(fi32, v2, v2)
                                                }
                                            }
                                            _ = fi32
                                        }
                                        v53 := GenS7{F0: "\n\t", F1: (fi20 * (fi20 / ((fi20) | 1)))}
                                        _ = v53
                                    } else {
                                        v4 = (true && (true || (-1 < 1)))
                                    }
                                    v4 = (0.0 == (3.14159 + ((float32(float64((1.5 / 1e10))) - 1e10) * (1.5 / 1.5))))
                                }
                            }
                            if true {
                                v1 = (v1 | 255)
                                switch fi20 {
                                case 1:
                                    for fi54 := 0; fi54 < 18; fi54++ {
                                        v2 = v2
                                        _ = fi54
                                    }
                                default:
                                    var v55 uint16 = 972
                                    _ = v55
                                }
                            }
                            _ = v29
                        }
                    }
                case -1:
                    for fi56 := 0; fi56 < 5; fi56++ {
                        v57 := GenS8{F0: 1, F1: 1}
                        _ = fi56
                        _ = v57
                    }
                default:
                    switch ((("\n\t" + "foo bar") + ("" + "a")) + "a") {
                    case "\n\t":
                        v1 = 174
                    case "foo bar":
                        var v58 int64 = -9223372036854775808
                        _ = v58
                    case "a":
                        var v59 float32 = -2.25
                        _ = v59
                    default:
                        if (-2.25 != 0.1) {
                            v60 := GenS9{F0: 2}
                            var v61 uint8 = 22
                            _ = v60
                            _ = v61
                        } else {
                            for fi62 := 0; fi62 < 10; fi62++ {
                                var v63 float32 = ((float32(v19) - float32(v19)) - -2.25)
                                fmt.Println(v63)
                                var v64 bool = (v4 && (139 >= 41))
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                            for fi65 := 0; fi65 < 5; fi65++ {
                                v66 := []int{fi20, fi20}
                                v4 = v4
                                v2 -= 626
                                _ = fi65
                                _ = v66
                            }
                            for fi67 := 0; fi67 < 8; fi67++ {
                                fmt.Println(fi20, v1)
                                _ = fi67
                            }
                        }
                    }
                }
            }
            if (((3.14159 + 0.0) >= 1e10) && (42 != 0)) {
                if false {
                    var v68 uint32 = 42
                    for fi69 := 0; fi69 < 1; fi69++ {
                        v68 += v68
                        _ = fi69
                    }
                    _ = v68
                }
            }
            _ = fi20
        }
    }
    var v70 int64 = -253
    v70 = (v70 / (((-(v70))) | 1))
    v71 := []int{-156, -216}
    v72 := GenS10{F0: -2147483648, F1: v70, F2: 1}
    if false {
        for fi73 := 0; fi73 < 16; fi73++ {
            delete(v3, "a")
            v74 := GenS11{F0: 32767, F1: fi73}
            _ = fi73
            _ = v74
        }
        v75 := GenS2{F0: -1}
        _ = v75
    } else {
        _ = len(v71)
        switch "\n\t" {
        case "\n\t":
            var v76 bool = (76 != 255)
            _ = v76
        case "a":
            switch "a" {
            case "":
                var v77 bool = (!v4)
                _ = v77
            }
        }
    }
    if false {
        if (65535 == 32767) {
            for fi78 := 0; fi78 < 3; fi78++ {
                var v79 uint64 = 2
                _ = fi78
                _ = v79
            }
        }
        var v80 uint64 = uint64(v1)
        for fi81 := 0; fi81 < 1; fi81++ {
            v1 -= v2
            var v82 float32 = 0.0
            _ = fi81
            _ = v82
        }
        _ = v80
    } else {
        v2 -= (v2 ^ v1)
        var v83 uint16 = (v1 * 32767)
        for fi84 := 0; fi84 < 11; fi84++ {
            v19 = v19
            for fi85 := 0; fi85 < 18; fi85++ {
                if false {
                    var v86 uint = 4294967295
                    _ = v86
                } else {
                    v1 += v2
                    switch "hello" {
                    case "foo bar":
                        var v87 uint32 = 4294967295
                        _ = v87
                    case "hello":
                        if len(v71) > 0 {
                            v71[0] = fi84
                        }
                    case "":
                        _ = len(v71)
                    default:
                        v88 := GenS3{F0: 398, F1: 2147483647, F2: 127}
                        _ = v88
                    }
                }
                v2 = v83
                _ = fi85
            }
            if (((v4 && true) || (4294967295 > 2147483648)) || ((fi84 - 42) != fi84)) {
                var v89 int16 = 1000
                v83 = (v2 & 65535)
                _ = v89
            } else {
                switch -168 {
                case -819:
                    var v90 uint16 = (v1 / ((881) | 1))
                    _ = v90
                case 32768:
                    if len(v71) > 0 {
                        v71[0] = (^(fi84))
                    }
                case -749:
                    v83 = (v1 << ((65535) & 63))
                }
                v83 = 2
            }
            _ = fi84
        }
        _ = v83
    }
    if (1 < -576) {
        switch 0 {
        case 376:
            for fi91 := 0; fi91 < 6; fi91++ {
                v92 := GenS12{F0: v70, F1: 100}
                if v4 {
                    v93 := GenS12{F0: (^(v70)), F1: 123}
                    _ = v93
                } else {
                    var v94 int32 = -1
                    _ = len(v3)
                    _ = len(v3)
                    _ = v94
                }
                var v95 int = (fi91 << ((715) & 63))
                _ = fi91
                _ = v92
                _ = v95
            }
        default:
            if v4 {
                v19 = (v19 - v19)
                var v96 int8 = 55
                _ = v96
            } else {
                var v97 uint64 = 4294967295
                v98 := GenS2{F0: -2}
                _ = v97
                _ = v98
            }
        }
    } else {
        var v99 int16 = int16(v2)
        if false {
            if false {
                for fi100 := 0; fi100 < 14; fi100++ {
                    switch fi100 {
                    case 32767:
                        switch "foo bar" {
                        case "hello":
                            for fi101 := 0; fi101 < 3; fi101++ {
                                var v102 float64 = float64(0.0)
                                delete(v3, "")
                                var v103 int = (fi100 / (((fi100 + (^(fi101)))) | 1))
                                _ = fi101
                                _ = v102
                                _ = v103
                            }
                        case "":
                            var v104 uint = 255
                            _ = v104
                        case "\n\t":
                            fmt.Println(v4, v4, v70)
                        }
                    case -239:
                        var v105 float32 = 1e10
                        _ = v105
                    }
                    v2 -= (v2 & v2)
                    var v106 uint = 65535
                    _ = fi100
                    _ = v106
                }
            } else {
                if v4 {
                    v19 = v19
                } else {
                    fmt.Println(v4)
                    v107 := GenS13{F0: 208, F1: -128}
                    _ = v107
                }
                _ = len(v71)
                if len(v71) > 0 {
                    v71[0] = 910
                }
            }
            v108 := v3["hello"]
            if false {
                switch (("foo bar" + ("foo bar" + "\n\t")) + "hello") {
                case "foo bar":
                    v109 := GenS13{F0: 2, F1: 42}
                    _ = v109
                case "\n\t":
                    switch (("foo bar" + "hello") + ("\n\t" + ("a" + "hello"))) {
                    case "a":
                        v19 += v19
                    case "":
                        if len(v71) >= 2 {
                            v71 = v71[1:len(v71)]
                        }
                    case "\n\t":
                        if v4 {
                            switch -853 {
                            case -9223372036854775808:
                                v99 += v99
                            case 32768:
                                for fi110 := 0; fi110 < 18; fi110++ {
                                    v111 := GenS14{F0: 630, F1: ("foo bar" == ("" + ("\n\t" + "a")))}
                                    v108 = (fi110 % ((-853) | 1))
                                    _ = fi110
                                    _ = v111
                                }
                            case 0:
                                v1 += v1
                            default:
                                if v4 {
                                    var v112 uint8 = 162
                                    _ = v112
                                } else {
                                    if len(v71) >= 2 {
                                        v71 = v71[1:len(v71)]
                                    }
                                    v108 = v108
                                }
                            }
                            var v113 uint32 = 65535
                            _ = v113
                        }
                    }
                case "a":
                    if v4 {
                        fmt.Println(v70, v2)
                    }
                default:
                    if (-2.25 != ((1.5 / 1.5) + 0.0)) {
                        if v4 {
                            if (float64(0.0) <= 1e10) {
                                _ = len(v71)
                            }
                        } else {
                            var v114 uint8 = 1
                            _ = v114
                        }
                        for fi115 := 0; fi115 < 13; fi115++ {
                            v116 := GenS11{F0: v99, F1: int(v70)}
                            _ = fi115
                            _ = v116
                        }
                        v117 := GenS13{F0: 1, F1: 1}
                        _ = v117
                    }
                }
            } else {
                v4 = ((v19 / v19) > v19)
                var v118 int8 = 1
                _ = v118
            }
            _ = v108
        } else {
            if len(v71) > 0 {
                v71[0] = 0
            }
            v2 -= v2
            for fi119 := 0; fi119 < 10; fi119++ {
                var v120 bool = (0.1 == ((float32((1.5 + 3.14159)) * 1.5) * -2.25))
                var v121 int32 = 423
                _ = fi119
                _ = v120
                _ = v121
            }
        }
        _ = v99
    }
    if (!((^(v70)) == v70)) {
        var v122 float32 = ((3.14159 / (float32(1e10) + 1.5)) + 3.14159)
        _ = v122
    }
    if (574 >= 42) {
        fmt.Println(v4, v70)
        v123 := GenS1{F0: (v2 - (v1 >> ((18446744073709551615) & 63)))}
        _ = v123
    } else {
        v124 := GenS3{F0: 1, F1: 2147483648, F2: 2}
        if v4 {
            if v4 {
                var v125 uint = 2
                _ = v125
            }
            v19 -= (v19 * (((v19 * -2.25) + (v19 + (((v19 * float64((0.0 * (0.1 + 0.0)))) - v19) + 0.0))) - (v19 * 1.5)))
            v4 = (!false)
        } else {
            v126 := v3[((("\n\t" + "foo bar") + "") + "hello")]
            var v127 uint = 4294967295
            _ = v126
            _ = v127
        }
        v128 := GenS15{F0: 42}
        _ = v124
        _ = v128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v72
}
