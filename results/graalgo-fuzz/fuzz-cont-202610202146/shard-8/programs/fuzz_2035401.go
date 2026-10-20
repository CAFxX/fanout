package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 int8; F1 int8; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 uint16; F2 uint }
type GenS9 struct { F0 uint8; F1 int16; F2 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 string }

func main() {
    var v1 bool = true
    var v2 string = "foo bar"
    v3 := make(map[int]uint32)
    var v4 uint64 = 4294967295
    var v5 uint8 = 201
    v6 := []int{667}
    v7 := GenS1{F0: 42, F1: 42, F2: 100}
    var v8 string = v2
    var v9 uint16 = 0
    v8 = (v8 + "foo bar")
    if (((("\n\t" + ((v2 + "hello") + (v2 + ("a" + v8)))) == ("a" + "foo bar")) && ((v5 > v5) || v1)) || v1) {
        v4 -= v4
    } else {
        var v10 bool = (!(((v2 + v2) + (((v2 + ((v2 + ((v2 + (v8 + v2)) + v2)) + (v2 + ((v8 + v8) + v2)))) + v8) + "a")) < ((v8 + ("hello" + v8)) + v2)))
        var v11 uint32 = 2
        _ = v10
        _ = v11
    }
    if v1 {
        _ = len(v3)
        if false {
            _ = len(v6)
            var v12 int32 = int32(v4)
            var v13 string = (v8 + ("foo bar" + v8))
            _ = v12
            _ = v13
        } else {
            v8 = v8
            if (!(v1 || (32767 <= 128))) {
                v14 := GenS2{F0: -32768, F1: v4, F2: ((4294967295 <= 32768) || ((-931 <= 561) && (32767 > 0)))}
                _ = v14
            } else {
                var v15 uint = 713
                _ = v15
            }
            switch v2 {
            case "a":
                var v16 uint = 42
                _ = v16
            case "\n\t":
                var v17 uint64 = 4294967295
                _ = v17
            case "hello":
                for fi18 := 0; fi18 < 15; fi18++ {
                    if v1 {
                        v9 = 976
                        if v1 {
                            switch (fi18 ^ (fi18 | 42)) {
                            case 42:
                                switch (-(fi18)) {
                                case 747:
                                    v4 += (^(v4))
                                case 127:
                                    if (false && false) {
                                        switch (v2 + "hello") {
                                        case "a":
                                            var v19 int32 = -1
                                            _ = v19
                                        case "":
                                            if (-32768 >= 0) {
                                                for fi20 := 0; fi20 < 14; fi20++ {
                                                    v9 += v9
                                                    if len(v6) > 0 {
                                                        v6[0] = fi18
                                                    }
                                                    _ = fi20
                                                }
                                                if false {
                                                    switch v8 {
                                                    case "foo bar":
                                                        var v21 uint16 = 32767
                                                        _ = v21
                                                    case "a":
                                                        fmt.Println(v9)
                                                    case "":
                                                        for fi22 := 0; fi22 < 1; fi22++ {
                                                            v4 -= 42
                                                            _ = fi22
                                                        }
                                                    }
                                                } else {
                                                    if v1 {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v3[fi18] = 911
                                                        case "":
                                                            if (1 <= 127) {
                                                                var v23 float64 = 0.0
                                                                _ = v23
                                                            } else {
                                                                if len(v6) > 0 {
                                                                    v6[0] = -1
                                                                }
                                                                v24 := GenS2{F0: -859, F1: (v4 | v4), F2: ((false || v1) || true)}
                                                                v9 = v9
                                                                _ = v24
                                                            }
                                                        }
                                                        if ((v5 & 63) > (v5 >> ((255) & 63))) {
                                                            var v25 uint = 0
                                                            _ = v25
                                                        } else {
                                                            fmt.Println(v5, v1, v9)
                                                            switch fi18 {
                                                            case 127:
                                                                if len(v6) > 0 {
                                                                    v6[0] = (^(fi18))
                                                                }
                                                            case -561:
                                                                fmt.Println(v2)
                                                            case 1:
                                                                _ = len(v3)
                                                            }
                                                        }
                                                    } else {
                                                        v9 -= 0
                                                        for fi26 := 0; fi26 < 16; fi26++ {
                                                            v27 := GenS1{F0: -1, F1: 9223372036854775807, F2: -128}
                                                            if (!false) {
                                                                delete(v3, (^(fi18)))
                                                                var v28 int32 = -51
                                                                _ = v28
                                                            }
                                                            _ = fi26
                                                            _ = v27
                                                        }
                                                    }
                                                }
                                                if false {
                                                    for fi29 := 0; fi29 < 17; fi29++ {
                                                        _ = len(v3)
                                                        var v30 int = 1
                                                        var v31 int8 = 127
                                                        _ = fi29
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    switch ((v2 + v2) + v2) {
                                                    case "\n\t":
                                                        var v32 int64 = 1
                                                        _ = v32
                                                    }
                                                    v4 = v4
                                                } else {
                                                    v33 := GenS2{F0: -576, F1: (v4 / ((v4) | 1)), F2: ((v9 ^ v9) != 989)}
                                                    for fi34 := 0; fi34 < 15; fi34++ {
                                                        var v35 float32 = 1e10
                                                        v36 := GenS3{F0: (^(v4)), F1: (v9 & (v9 << ((1) & 63)))}
                                                        v3[(fi34 << ((65535) & 63))] = 65535
                                                        var v37 int64 = -9223372036854775808
                                                        if v1 {
                                                            var v38 float32 = (v35 * -2.25)
                                                            v37 += -26
                                                            fmt.Println(v5)
                                                            v39 := GenS2{F0: 32767, F1: v4, F2: ((v38 - (0.1 - (v35 * 0.1))) == float32(-2.25))}
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        _ = fi34
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    v8 = "hello"
                                                    _ = v33
                                                }
                                            } else {
                                                var v40 int = 128
                                                for fi41 := 0; fi41 < 7; fi41++ {
                                                    var v42 uint32 = 2
                                                    var v43 float32 = 0.0
                                                    _ = fi41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = v40
                                            }
                                        case "hello":
                                            for fi44 := 0; fi44 < 20; fi44++ {
                                                if v1 {
                                                    v6 = append(v6, 938)
                                                    v45 := GenS4{F0: fi44, F1: (((v2 + "a") + ((v8 + "a") + v8)) + (v2 + (v8 + v2))), F2: (fi44 | (fi18 % (((fi44 << ((1) & 63))) | 1)))}
                                                    var v46 int16 = 1000
                                                    _ = v45
                                                    _ = v46
                                                }
                                                var v47 uint8 = (v5 | 2)
                                                v2 = "\n\t"
                                                v2 = v2
                                                v48 := v3[(fi18 >> ((18446744073709551615) & 63))]
                                                switch fi18 {
                                                case -2147483648:
                                                    v47 = (v5 - 0)
                                                case -9223372036854775808:
                                                    if (!(3.14159 != 3.14159)) {
                                                        v49 := GenS5{F0: -128, F1: -60, F2: -9223372036854775808}
                                                        var v50 uint16 = 1
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                case 346:
                                                    switch "hello" {
                                                    case "a":
                                                        switch 746 {
                                                        case 277:
                                                            fmt.Println(fi44, v48, v48)
                                                        case 1:
                                                            if (74 != 127) {
                                                                if (true || (127 == -128)) {
                                                                    switch v8 {
                                                                    case "foo bar":
                                                                        fmt.Println(fi18, v8)
                                                                    case "":
                                                                        fmt.Println(v2, v8, fi44)
                                                                    case "hello":
                                                                        v2 = v2
                                                                    default:
                                                                        var v51 bool = false
                                                                        _ = v51
                                                                    }
                                                                    if v1 {
                                                                        switch "hello" {
                                                                        case "a":
                                                                            if ((true || (!v1)) || true) {
                                                                                v5 -= v47
                                                                                if v1 {
                                                                                    var v52 bool = true
                                                                                    _ = v52
                                                                                }
                                                                            } else {
                                                                                var v53 uint8 = v5
                                                                                v8 = "foo bar"
                                                                                _ = v53
                                                                            }
                                                                        case "hello":
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v54 uint32 = 129
                                                                        fmt.Println(v5)
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            fmt.Println(fi44)
                                                                        case "a":
                                                                            switch (v2 + (v2 + "\n\t")) {
                                                                            case "a":
                                                                                var v55 uint = 42
                                                                                _ = v55
                                                                            case "hello":
                                                                                switch fi18 {
                                                                                case -366:
                                                                                    var v56 uint64 = (v4 ^ (v4 / (((v4 % ((4294967295) | 1))) | 1)))
                                                                                    _ = v56
                                                                                case 316:
                                                                                    if v1 {
                                                                                        var v57 bool = (!(2147483647 == (fi18 ^ (fi18 - 180))))
                                                                                        _ = v57
                                                                                    }
                                                                                case 0:
                                                                                    v6 = append(v6, (fi44 / ((-2147483648) | 1)))
                                                                                default:
                                                                                    var v58 float32 = 3.14159
                                                                                    _ = v58
                                                                                }
                                                                            case "foo bar":
                                                                                var v59 uint32 = v48
                                                                                _ = v59
                                                                            default:
                                                                                var v60 string = (v2 + v8)
                                                                                _ = v60
                                                                            }
                                                                        case "hello":
                                                                            if true {
                                                                                var v61 uint = 528
                                                                                _ = v61
                                                                            } else {
                                                                                if ((!false) && ((v1 && (-620 < -9223372036854775808)) && false)) {
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                } else {
                                                                                    v62 := GenS3{F0: v4, F1: v9}
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v63 uint = 102
                                                                            _ = v63
                                                                        }
                                                                        _ = v54
                                                                    }
                                                                    if len(v6) > 0 {
                                                                        v6[0] = 2147483647
                                                                    }
                                                                } else {
                                                                    switch fi44 {
                                                                    case 886:
                                                                        var v64 float32 = ((0.1 * (float32(3.14159) + 1.5)) + 1e10)
                                                                        _ = v64
                                                                    case 128:
                                                                        v5 -= (v5 - 255)
                                                                    }
                                                                }
                                                                if ((!(!v1)) && v1) {
                                                                    v65 := GenS6{F0: 183}
                                                                    if v1 {
                                                                        var v66 int16 = 807
                                                                        v5 = 1
                                                                        _ = v66
                                                                    } else {
                                                                        var v67 bool = (v9 != (v9 * (v9 << ((42) & 63))))
                                                                        _ = v67
                                                                    }
                                                                    _ = v65
                                                                } else {
                                                                    v2 = (v2 + v8)
                                                                    if v1 {
                                                                        if ((v4 % ((679) | 1)) < v4) {
                                                                            v4 -= 466
                                                                            if len(v6) > 0 {
                                                                                v6[0] = fi44
                                                                            }
                                                                        } else {
                                                                            v68 := GenS7{F0: ((-(fi44)) != (^(fi44)))}
                                                                            v69 := v3[(fi44 | fi18)]
                                                                            var v70 float32 = float32(3.14159)
                                                                            var v71 bool = true
                                                                            v72 := GenS1{F0: -1, F1: 643, F2: -128}
                                                                            _ = v68
                                                                            _ = v69
                                                                            _ = v70
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                }
                                                                v5 = (v47 * v47)
                                                            }
                                                        }
                                                    }
                                                }
                                                var v73 uint16 = v9
                                                v1 = (!(-2.25 == 3.14159))
                                                _ = fi44
                                                _ = v47
                                                _ = v48
                                                _ = v73
                                            }
                                        default:
                                            for fi74 := 0; fi74 < 15; fi74++ {
                                                v75 := GenS3{F0: (^(v4)), F1: (v9 * (v9 | (v9 >> ((42) & 63))))}
                                                if v1 {
                                                    switch (v8 + (v8 + "a")) {
                                                    case "a":
                                                        if (0 == v5) {
                                                            var v76 float32 = 0.1
                                                            var v77 int16 = 586
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            var v78 float32 = 1.5
                                                            if v1 {
                                                                delete(v3, (fi74 + (fi18 & fi18)))
                                                            } else {
                                                                v5 = v5
                                                                v4 += v4
                                                                switch (^(fi18)) {
                                                                case 32767:
                                                                    var v79 uint16 = 192
                                                                    _ = v79
                                                                case -881:
                                                                    var v80 int = 32767
                                                                    _ = v80
                                                                }
                                                            }
                                                            v78 -= 0.0
                                                            _ = v78
                                                        }
                                                    case "hello":
                                                        if (2147483647 == 1) {
                                                            if true {
                                                                v4 = 274
                                                                fmt.Println(v9, v1)
                                                            } else {
                                                                _ = len(v6)
                                                                delete(v3, (fi74 >> ((65535) & 63)))
                                                                v3[(fi74 + -849)] = 2147483648
                                                            }
                                                            fmt.Println(v8)
                                                            if len(v6) > 0 {
                                                                v6[0] = fi74
                                                            }
                                                        } else {
                                                            var v81 float32 = (1e10 + 1e10)
                                                            switch 1 {
                                                            case 988:
                                                                var v82 int64 = 0
                                                                _ = v82
                                                            default:
                                                                v6 = append(v6, fi18)
                                                            }
                                                            v5 = (v5 + v5)
                                                            _ = v81
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v8)
                                                    }
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v83 int16 = -32768
                                                        _ = v83
                                                    case "a":
                                                        fmt.Println(fi18, v5)
                                                    default:
                                                        fmt.Println(v1, v4)
                                                    }
                                                }
                                                v84 := GenS8{F0: 18446744073709551615, F1: v9, F2: 65535}
                                                var v85 int32 = 53
                                                var v86 int64 = 9223372036854775807
                                                if (v1 || true) {
                                                    var v87 string = v8
                                                    _ = v87
                                                }
                                                var v88 uint32 = 780
                                                v4 += v4
                                                fmt.Println(fi74, v86, v8)
                                                _ = fi74
                                                _ = v75
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                                _ = v88
                                            }
                                        }
                                        if (!(!v1)) {
                                            var v89 int16 = 32767
                                            if false {
                                                v89 = (^(v89))
                                                for fi90 := 0; fi90 < 17; fi90++ {
                                                    v2 = (("\n\t" + ("foo bar" + v2)) + v2)
                                                    _ = fi90
                                                }
                                                var v91 uint64 = (^(v4))
                                                _ = v91
                                            } else {
                                                for fi92 := 0; fi92 < 13; fi92++ {
                                                    v93 := GenS6{F0: int16(v4)}
                                                    _ = fi92
                                                    _ = v93
                                                }
                                                var v94 uint64 = 18446744073709551615
                                                _ = v94
                                            }
                                            _ = v89
                                        }
                                    } else {
                                        v5 -= (v5 + (^(v5)))
                                    }
                                }
                            case 32767:
                                switch (-(fi18)) {
                                case 0:
                                    if v1 {
                                        v95 := GenS9{F0: (v5 / ((1) | 1)), F1: int16(v4), F2: (((v8 + v8) + v8) + "hello")}
                                        _ = v95
                                    }
                                default:
                                    v96 := GenS10{F0: 850}
                                    _ = v96
                                }
                            default:
                                v8 = ((v2 + "") + v8)
                            }
                            var v97 float32 = 0.0
                            if true {
                                for fi98 := 0; fi98 < 19; fi98++ {
                                    v2 = ("hello" + v8)
                                    switch "foo bar" {
                                    case "a":
                                        v99 := GenS10{F0: 32767}
                                        _ = v99
                                    case "\n\t":
                                        v9 -= 65535
                                    }
                                    _ = len(v6)
                                    _ = fi98
                                }
                                var v100 int8 = -32
                                _ = v100
                            }
                            _ = v97
                        }
                        v2 = v8
                    } else {
                        var v101 int16 = 32767
                        var v102 int16 = (v101 & v101)
                        _ = v101
                        _ = v102
                    }
                    var v103 uint = 836
                    if (127 != 127) {
                        if v1 {
                            v104 := v3[fi18]
                            fmt.Println(v9)
                            _ = v104
                        }
                    }
                    _ = fi18
                    _ = v103
                }
            default:
                var v105 float64 = 0.0
                _ = v105
            }
        }
        var v106 uint8 = 18
        _ = v106
    }
    v1 = (v1 && v1)
    for fi107 := 0; fi107 < 4; fi107++ {
        if v1 {
            for fi108 := 0; fi108 < 5; fi108++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if ((((1.5 - 0.1) / -2.25) - -2.25) == 0.0) {
                    var v109 uint16 = 42
                    v110 := GenS8{F0: v4, F1: (v9 | v109), F2: 42}
                    _ = v109
                    _ = v110
                }
                _ = fi108
            }
        } else {
            if v1 {
                for fi111 := 0; fi111 < 11; fi111++ {
                    if false {
                        v5 = v5
                    } else {
                        _ = len(v6)
                    }
                    _ = len(v6)
                    fmt.Println(fi111)
                    _ = fi111
                }
            }
            switch v8 {
            case "foo bar":
                var v112 int8 = 127
                _ = v112
            case "\n\t":
                for fi113 := 0; fi113 < 12; fi113++ {
                    fmt.Println(v8)
                    if v1 {
                        if (42 != -1) {
                            v9 = v9
                            _ = len(v3)
                        }
                        v114 := GenS2{F0: 1, F1: (^(v4)), F2: (0 != -9223372036854775808)}
                        _ = v114
                    }
                    var v115 uint8 = 173
                    if (!v1) {
                        v4 = (v4 >> ((18446744073709551615) & 63))
                        v116 := GenS5{F0: -116, F1: 1, F2: -1}
                        v8 = ((v8 + v2) + "a")
                        _ = v116
                    }
                    v117 := GenS11{F0: ((v2 + ((v8 + (v8 + ("a" + "hello"))) + v8)) + "hello")}
                    _ = fi113
                    _ = v115
                    _ = v117
                }
            default:
                v9 = 925
            }
        }
        var v118 int = -364
        var v119 float32 = 1e10
        _ = fi107
        _ = v118
        _ = v119
    }
    var v120 uint8 = uint8(v9)
    v6 = append(v6, 32767)
    var v121 string = v2
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
