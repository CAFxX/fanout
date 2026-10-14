package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint; F1 int; F2 int16 }
type GenS6 struct { F0 bool; F1 int16; F2 uint16 }
type GenS7 struct { F0 int8; F1 string; F2 int }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS10 struct { F0 int16; F1 string; F2 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := []int{206, -1}
    v3 := make(map[string]int)
    var v4 int = -691
    v5 := make(map[string]bool)
    var v6 uint16 = v1
    v6 = (^(v1))
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 bool = (false && true)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        delete(v3, "")
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        if false {
            v1 += (v6 >> ((42) & 63))
            v1 -= 255
            switch 42 {
            case -686:
                for fi10 := 0; fi10 < 4; fi10++ {
                    v1 -= (^(v6))
                    v4 = -2147483648
                    _ = len(v2)
                    _ = fi10
                }
            default:
                v11 := GenS1{F0: 1}
                _ = v11
            }
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 bool = (true && (215 >= 2))
        v4 = (^(fi12))
        var v14 int8 = 91
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int = (v4 & int(v1))
    v15 -= v15
    v16 := GenS2{F0: v6, F1: 42, F2: 793}
    fmt.Println(v1, v15)
    fmt.Println(v1, v4)
    var v17 int8 = -25
    if (!((true || false) || (!(65535 <= 2147483647)))) {
        v3["a"] = 128
        v18 := GenS3{F0: 701, F1: (v6 & (v6 ^ (v6 % (((v6 >> ((65535) & 63))) | 1))))}
        _ = v18
    } else {
        v1 = 42
        v6 -= v6
    }
    if (-9223372036854775808 != 1) {
        if (2 != 0) {
            switch "a" {
            case "hello":
                for fi19 := 0; fi19 < 8; fi19++ {
                    _ = len(v3)
                    var v20 int = v4
                    v21 := v5[("hello" + "foo bar")]
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            if false {
                switch "foo bar" {
                case "foo bar":
                    var v22 int64 = 9223372036854775807
                    _ = v22
                case "":
                    var v23 uint16 = (^(v6))
                    _ = v23
                case "hello":
                    for fi24 := 0; fi24 < 9; fi24++ {
                        if true {
                            var v25 int64 = 42
                            _ = v25
                        } else {
                            v6 = (v6 & v6)
                            var v26 uint8 = 165
                            v27 := v3[("a" + ("foo bar" + "foo bar"))]
                            _ = v26
                            _ = v27
                        }
                        _ = fi24
                    }
                default:
                    if false {
                        fmt.Println(v4, v15)
                        var v28 int32 = -853
                        fmt.Println(v15)
                        _ = v28
                    } else {
                        if false {
                            var v29 int16 = 0
                            _ = v29
                        } else {
                            v17 = (-(v17))
                            for fi30 := 0; fi30 < 16; fi30++ {
                                v31 := GenS1{F0: -1}
                                v6 -= v6
                                var v32 float64 = 0.0
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        var v33 float64 = (float64(((0.1 * (((1e10 + 0.0) - float32(3.14159)) / (3.14159 / -2.25))) / 1.0)) + 1e10)
                        for fi34 := 0; fi34 < 19; fi34++ {
                            var v35 int32 = 1
                            v15 = (v4 % ((v4) | 1))
                            for fi36 := 0; fi36 < 13; fi36++ {
                                if ((fi34 >> ((370) & 63)) != v4) {
                                    fmt.Println(fi34, fi34, fi36)
                                } else {
                                    if false {
                                        fmt.Println(fi34, v35)
                                    } else {
                                        if false {
                                            var v37 int = -122
                                            _ = v37
                                        }
                                        var v38 int32 = (v35 & (v35 ^ (^(v35))))
                                        var v39 int8 = 127
                                        _ = v38
                                        _ = v39
                                    }
                                    var v40 string = ("" + "\n\t")
                                    switch v40 {
                                    case "a":
                                        v41 := GenS1{F0: (v17 ^ 1)}
                                        _ = v41
                                    case "foo bar":
                                        if (2147483647 >= (v35 | (v35 << ((42) & 63)))) {
                                            var v42 uint16 = v6
                                            var v43 uint64 = 409
                                            switch ((v40 + v40) + "") {
                                            case "foo bar":
                                                if (v17 == (^(v17))) {
                                                    fmt.Println(v40)
                                                    v42 = (v6 * (v6 ^ v6))
                                                } else {
                                                    var v44 uint8 = 0
                                                    _ = v44
                                                }
                                            }
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v15 -= fi34
                                            if true {
                                                _ = len(v2)
                                                fmt.Println(fi34)
                                            } else {
                                                var v45 int = v4
                                                v46 := GenS3{F0: 453, F1: v6}
                                                switch ("\n\t" + "a") {
                                                case "":
                                                    var v47 bool = (!false)
                                                    _ = v47
                                                case "foo bar":
                                                    v4 = int(v35)
                                                case "\n\t":
                                                    var v48 uint32 = uint32(v1)
                                                    _ = v48
                                                }
                                                v49 := GenS4{F0: 259, F1: 42}
                                                v15 = -529
                                                if ((42 != (v35 + v35)) || false) {
                                                    v50 := GenS5{F0: 871, F1: fi34, F2: 42}
                                                    if true {
                                                        var v51 bool = false
                                                        if v51 {
                                                            switch v40 {
                                                            case "a":
                                                                var v52 uint32 = 0
                                                                _ = v52
                                                            }
                                                            delete(v5, (v40 + v40))
                                                            var v53 float64 = v33
                                                            _ = v53
                                                        }
                                                        if (!((0.1 / -2.25) >= 1e10)) {
                                                            v54 := GenS2{F0: v6, F1: 18446744073709551615, F2: 978}
                                                            v55 := GenS1{F0: v17}
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        _ = v51
                                                    } else {
                                                        v6 += v6
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v56 uint32 = 2147483648
                                                        _ = v56
                                                    }
                                                    _ = v50
                                                }
                                                _ = v45
                                                _ = v46
                                                _ = v49
                                            }
                                        }
                                    case "hello":
                                        fmt.Println(v6, v4)
                                    default:
                                        var v57 uint = 4294967295
                                        _ = v57
                                    }
                                    _ = v40
                                }
                                fmt.Println(v33, fi36)
                                fmt.Println(v33, v1, v15)
                                fmt.Println(fi34, v4, v35)
                                _ = fi36
                            }
                            _ = fi34
                            _ = v35
                        }
                        _ = v33
                    }
                }
            }
        } else {
            v1 = (v1 | 32767)
            v6 = (v1 << ((18446744073709551615) & 63))
        }
        v58 := GenS6{F0: (!true), F1: 32767, F2: v6}
        _ = v58
    } else {
        fmt.Println(v15)
        var v59 int32 = -1
        _ = v59
    }
    fmt.Println(v17, v1)
    if false {
        if len(v2) > 0 {
            v2[0] = v4
        }
        if (false || (2 < 2147483648)) {
            for fi60 := 0; fi60 < 5; fi60++ {
                for fi61 := 0; fi61 < 10; fi61++ {
                    var v62 uint64 = 0
                    if false {
                        var v63 int8 = 0
                        switch "\n\t" {
                        case "":
                            _ = len(v2)
                        case "a":
                            var v64 int8 = v63
                            var v65 uint64 = v62
                            _ = v64
                            _ = v65
                        default:
                            var v66 uint32 = 2147483648
                            _ = v66
                        }
                        var v67 uint = 65535
                        _ = v63
                        _ = v67
                    }
                    v15 += 32767
                    v68 := GenS1{F0: (v17 / ((v17) | 1))}
                    _ = fi61
                    _ = v62
                    _ = v68
                }
                if ((v17 << ((65535) & 63)) != v17) {
                    v17 += int8(fi60)
                    var v69 uint = 255
                    v17 = v17
                    _ = v69
                } else {
                    var v70 float32 = (0.1 * float32(0.0))
                    var v71 float64 = (float64(float32((1e10 + (0.1 * -2.25)))) / 0.1)
                    _ = v70
                    _ = v71
                }
                var v72 int64 = 268
                _ = fi60
                _ = v72
            }
        } else {
            for fi73 := 0; fi73 < 11; fi73++ {
                if (true && false) {
                    v6 -= (v6 | (v1 % ((v6) | 1)))
                }
                _ = fi73
            }
        }
    } else {
        switch "foo bar" {
        case "a":
            for fi74 := 0; fi74 < 11; fi74++ {
                fmt.Println(fi74)
                _ = fi74
            }
        default:
            var v75 int64 = 0
            _ = v75
        }
        if true {
            v15 = v15
            v4 = 128
            switch 32768 {
            case -161:
                v6 = (^(v6))
            case 903:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                v1 = (v6 ^ (v6 >> ((65535) & 63)))
            }
        } else {
            fmt.Println(v17)
            if (65535 < 2) {
                for fi76 := 0; fi76 < 6; fi76++ {
                    if false {
                        if (((float64(-2.25) / (3.14159 * 1e10)) * 0.1) > (1.5 - 3.14159)) {
                            switch (v4 | (v15 * v4)) {
                            case 32767:
                                if (true || (1000 > 0)) {
                                    for fi77 := 0; fi77 < 4; fi77++ {
                                        v78 := GenS5{F0: 18446744073709551615, F1: 127, F2: 967}
                                        if true {
                                            fmt.Println(v6)
                                        }
                                        var v79 uint16 = (v6 << ((915) & 63))
                                        _ = fi77
                                        _ = v78
                                        _ = v79
                                    }
                                }
                            default:
                                switch ("" + "hello") {
                                case "hello":
                                    if (2147483647 >= 0) {
                                        v1 = (v6 << ((18446744073709551615) & 63))
                                        var v80 int = (-(fi76))
                                        switch "foo bar" {
                                        case "hello":
                                            var v81 int32 = 0
                                            _ = v81
                                        case "foo bar":
                                            for fi82 := 0; fi82 < 10; fi82++ {
                                                if false {
                                                    v15 = fi76
                                                }
                                                v83 := GenS1{F0: v17}
                                                if false {
                                                    v2 = append(v2, fi76)
                                                    var v84 int64 = -1
                                                    v85 := GenS4{F0: 4294967295, F1: 778}
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    fmt.Println(fi82, v80)
                                                }
                                                fmt.Println(v17)
                                                _ = fi82
                                                _ = v83
                                            }
                                        }
                                        _ = v80
                                    }
                                case "":
                                    var v86 float64 = -2.25
                                    _ = v86
                                case "a":
                                    fmt.Println(v4)
                                default:
                                    _ = len(v5)
                                }
                            }
                            v87 := GenS2{F0: (v1 % ((v6) | 1)), F1: 18446744073709551615, F2: 0}
                            _ = v87
                        }
                        var v88 string = "a"
                        if false {
                            if ((v17 <= v17) || false) {
                                for fi89 := 0; fi89 < 18; fi89++ {
                                    if len(v2) > 0 {
                                        v2[0] = v4
                                    }
                                    _ = fi89
                                }
                                v90 := GenS7{F0: (^(v17)), F1: ("\n\t" + ("\n\t" + (("hello" + v88) + v88))), F2: v15}
                                _ = v90
                            }
                        } else {
                            var v91 uint64 = 308
                            v92 := GenS1{F0: v17}
                            v6 = (v6 % ((v6) | 1))
                            _ = v91
                            _ = v92
                        }
                        _ = v88
                    }
                    _ = fi76
                }
                v93 := GenS8{F0: 137, F1: ("a" + ("hello" + (("\n\t" + "foo bar") + "hello")))}
                _ = v93
            } else {
                v4 += v4
                if true {
                    fmt.Println(v4, v4)
                    var v94 int64 = 1
                    switch (-(v4)) {
                    case 802:
                        var v95 uint16 = (v1 * (v1 * (^(v1))))
                        _ = v95
                    case 1:
                        var v96 string = "\n\t"
                        _ = v96
                    default:
                        for fi97 := 0; fi97 < 11; fi97++ {
                            v15 = v4
                            for fi98 := 0; fi98 < 17; fi98++ {
                                var v99 int16 = int16(v94)
                                v15 = fi98
                                _ = fi98
                                _ = v99
                            }
                            v17 = -1
                            _ = fi97
                        }
                    }
                    _ = v94
                } else {
                    v6 -= v1
                    v6 -= (v1 >> ((65535) & 63))
                    for fi100 := 0; fi100 < 18; fi100++ {
                        v6 = v6
                        for fi101 := 0; fi101 < 11; fi101++ {
                            switch (v4 >> ((65535) & 63)) {
                            case -9223372036854775808:
                                if (0 <= 42) {
                                    switch fi101 {
                                    case 4294967295:
                                        v15 = int(v1)
                                    case 0:
                                        _ = len(v2)
                                        var v102 uint64 = 395
                                        _ = v102
                                    }
                                    if (false || (((1e10 + (1e10 + (((1.5 * 0.0) - -2.25) / 1.5))) - 1e10) > (-2.25 * -2.25))) {
                                        var v103 uint = 42
                                        if (true && true) {
                                            switch "" {
                                            case "\n\t":
                                                _ = len(v5)
                                            }
                                        }
                                        var v104 int16 = 1000
                                        _ = v103
                                        _ = v104
                                    }
                                    v2 = append(v2, (fi101 >> ((648) & 63)))
                                } else {
                                    v105 := GenS9{F0: 124, F1: (155 == 0), F2: 128}
                                    _ = v105
                                }
                            case 743:
                                var v106 uint16 = 0
                                _ = v106
                            }
                            _ = fi101
                        }
                        v107 := GenS5{F0: uint(v1), F1: v15, F2: -812}
                        _ = fi100
                        _ = v107
                    }
                }
                fmt.Println(v17, v4, v4)
            }
            v108 := GenS6{F0: false, F1: -91, F2: 42}
            _ = v108
        }
        switch (v15 % ((v15) | 1)) {
        case 52:
            var v109 int32 = 1
            _ = v109
        case -982:
            var v110 uint32 = 2147483647
            _ = v110
        case 42:
            fmt.Println(v4, v1)
        }
    }
    v111 := GenS10{F0: -711, F1: "", F2: 2147483647}
    v2 = append(v2, (v15 % (((v15 / ((v4) | 1))) | 1)))
    if ((!true) && (0 <= 42)) {
        if true {
            v4 -= (-(v4))
            v15 += (v4 + -2147483648)
            var v112 int16 = -573
            _ = v112
        }
        v1 = 0
    } else {
        v113 := GenS9{F0: 0, F1: ("\n\t" != "\n\t"), F2: 74}
        _ = v113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v111
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
