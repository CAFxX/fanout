package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 string; F1 int16; F2 uint64 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 int64 }
func genFn0(p0 int64, p1 uint32, p2 bool) int64 {
    p0 = (p0 * p0)
    v6 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 uint8 = 172
    v4 := []int{-278, 128, 0}
    v5 := make(map[uint32]uint32)
    _ = len(v2)
    switch "\n\t" {
    case "":
        v1 = v1
    case "a":
        switch "a" {
        case "foo bar":
            if (0.1 == (1.5 - 0.1)) {
                fmt.Println(v3, v1)
            }
        case "a":
            if (!(0 < 18446744073709551615)) {
                var v7 bool = false
                if (!(3.14159 != -2.25)) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v8 := GenS1{F0: 2}
                    if (!v1) {
                        fmt.Println(v7)
                    } else {
                        v9 := v5[0]
                        fmt.Println(v9, v3)
                        var v10 uint16 = 435
                        _ = v10
                        _ = v9
                    }
                    _ = v8
                }
                v1 = (!(v7 && ((v3 | v3) > (v3 >> ((2) & 63)))))
                _ = v7
            } else {
                v11 := GenS2{F0: 9223372036854775807, F1: -73}
                switch "" {
                case "a":
                    for fi12 := 0; fi12 < 3; fi12++ {
                        switch "\n\t" {
                        case "hello":
                            var v13 float64 = 3.14159
                            _ = v13
                        case "\n\t":
                            for fi14 := 0; fi14 < 9; fi14++ {
                                v2[850] = "a"
                                switch 1 {
                                case 999:
                                    switch (fi14 / (((fi14 << ((255) & 63))) | 1)) {
                                    case 2147483647:
                                        if v1 {
                                            switch (fi12 ^ fi14) {
                                            case 87:
                                                v3 = (v3 * v3)
                                                if v1 {
                                                    var v15 int64 = -352
                                                    _ = v15
                                                } else {
                                                    var v16 uint8 = 2
                                                    fmt.Println(fi12)
                                                    var v17 int32 = -1
                                                    _ = v16
                                                    _ = v17
                                                }
                                                switch ("foo bar" + "foo bar") {
                                                case "":
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi12 ^ (fi14 ^ fi12))
                                                    }
                                                case "\n\t":
                                                    v3 -= (v3 % ((v3) | 1))
                                                default:
                                                    v1 = (((!v1) || v1) && (32767 > -706))
                                                }
                                            case -1:
                                                if (42 == 32767) {
                                                    var v18 string = "foo bar"
                                                    _ = v18
                                                }
                                            case -917:
                                                fmt.Println(fi14, v3)
                                            default:
                                                if (!(!v1)) {
                                                    switch "\n\t" {
                                                    case "":
                                                        v1 = (556 > 0)
                                                    case "hello":
                                                        var v19 bool = (!(42 != 65535))
                                                        _ = v19
                                                    case "a":
                                                        v3 -= v3
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    v1 = (!v1)
                                                } else {
                                                    var v20 int64 = 0
                                                    v20 = (v20 | v20)
                                                    _ = v20
                                                }
                                            }
                                            if (42 <= 662) {
                                                v21 := GenS2{F0: -1, F1: -802}
                                                _ = v21
                                            } else {
                                                v3 -= 2
                                                if true {
                                                    if v1 {
                                                        v1 = (-358 == 2147483647)
                                                    } else {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v22 uint16 = 32767
                                                            _ = v22
                                                        case "foo bar":
                                                            if v1 {
                                                                var v23 int8 = 127
                                                                v24 := GenS3{F0: (("a" + "a") + "")}
                                                                _ = v23
                                                                _ = v24
                                                            }
                                                        case "":
                                                            var v25 uint = 1
                                                            _ = v25
                                                        }
                                                    }
                                                    v3 += v3
                                                    if (!false) {
                                                        v1 = (2 == 465)
                                                        v26 := GenS2{F0: 1, F1: 0}
                                                        _ = v26
                                                    }
                                                }
                                                fmt.Println(v3, v1, fi12)
                                            }
                                            v1 = (1 < v3)
                                        }
                                    case 313:
                                        if true {
                                            switch "hello" {
                                            case "foo bar":
                                                switch (fi14 << ((65535) & 63)) {
                                                case 9223372036854775807:
                                                    v1 = (((65535 <= 65535) && false) || v1)
                                                case 128:
                                                    if v1 {
                                                        v27 := GenS4{F0: ("a" + ("" + "a")), F1: "foo bar"}
                                                        if ((42 >= 35) && (255 == 18446744073709551615)) {
                                                            fmt.Println(v3)
                                                            v28 := GenS1{F0: 424}
                                                            _ = v28
                                                        } else {
                                                            var v29 uint8 = v3
                                                            var v30 int64 = 9223372036854775807
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                        _ = v27
                                                    } else {
                                                        var v31 uint32 = 4294967295
                                                        _ = v31
                                                    }
                                                case -4:
                                                    v32 := GenS3{F0: "foo bar"}
                                                    _ = v32
                                                }
                                            case "":
                                                v2[65535] = "a"
                                            default:
                                                switch "\n\t" {
                                                case "hello":
                                                    v33 := GenS5{F0: v3, F1: (500 != 809)}
                                                    _ = v33
                                                default:
                                                    if ((!(false || v1)) && v1) {
                                                        v1 = true
                                                        v3 = (v3 << ((1) & 63))
                                                        v34 := GenS6{F0: -330, F1: genFn0(-9223372036854775808, 2, (true && ((1e10 == -2.25) || false)))}
                                                        if v1 {
                                                            v35 := GenS1{F0: 42}
                                                            _ = len(v2)
                                                            var v36 int8 = 127
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        _ = v34
                                                    } else {
                                                        var v37 int = 128
                                                        _ = v37
                                                    }
                                                }
                                            }
                                        }
                                    }
                                case -2147483648:
                                    switch ("foo bar" + "a") {
                                    case "foo bar":
                                        if v1 {
                                            v3 -= v3
                                            v38 := GenS4{F0: ("hello" + "foo bar"), F1: "foo bar"}
                                            var v39 int8 = -4
                                            _ = v38
                                            _ = v39
                                        } else {
                                            var v40 int = (fi12 / (((fi12 - 32767)) | 1))
                                            v4 = append(v4, (v40 - (v40 >> ((243) & 63))))
                                            _ = v40
                                        }
                                    }
                                    _ = len(v2)
                                    if v1 {
                                        if v1 {
                                            if v1 {
                                                switch fi12 {
                                                case 9223372036854775807:
                                                    if (65535 != 32767) {
                                                        v41 := GenS5{F0: v3, F1: v1}
                                                        fmt.Println(fi12, fi12, v1)
                                                        _ = v41
                                                    }
                                                case 32767:
                                                    var v42 float64 = -2.25
                                                    _ = v42
                                                }
                                                v1 = true
                                                var v43 int8 = 0
                                                _ = len(v2)
                                                _ = v43
                                            } else {
                                                v3 += 42
                                                fmt.Println(v3, fi14)
                                                if v1 {
                                                    v1 = ((18446744073709551615 < 42) && false)
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v44 float32 = 0.1
                                                    var v45 int = fi14
                                                    _ = v44
                                                    _ = v45
                                                }
                                                var v46 uint = 65535
                                                v46 = 18446744073709551615
                                                _ = v46
                                            }
                                        }
                                        v47 := GenS4{F0: "foo bar", F1: "a"}
                                        v3 = v3
                                        v3 += 150
                                        var v48 uint = 65535
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v49 := GenS2{F0: 1, F1: 883}
                                        var v50 int16 = 438
                                        _ = v49
                                        _ = v50
                                    }
                                case 32767:
                                    if ((!v1) && (-9223372036854775808 < -185)) {
                                        fmt.Println(fi14)
                                        v3 -= v3
                                    } else {
                                        var v51 uint32 = 42
                                        switch "a" {
                                        case "hello":
                                            v3 = v3
                                        case "\n\t":
                                            v52 := GenS7{F0: "hello", F1: 0, F2: 681}
                                            if true {
                                                v3 += v3
                                            } else {
                                                fmt.Println(v51)
                                                var v53 int = -9223372036854775808
                                                _ = v53
                                            }
                                            v54 := GenS8{F0: v51, F1: (v3 & (v3 & (v3 - (v3 / (((v3 / (((^(v3))) | 1))) | 1)))))}
                                            _ = v52
                                            _ = v54
                                        }
                                        if (42 > 370) {
                                            switch "hello" {
                                            case "":
                                                var v55 bool = (v1 && ((!(!(((((((("" + "hello") + "foo bar") + "a") + ("foo bar" + "\n\t")) + ("foo bar" + "a")) + "\n\t") + "a") == ("" + "\n\t")))) || (127 != 0)))
                                                _ = v55
                                            }
                                            v1 = false
                                        } else {
                                            fmt.Println(v51, v1)
                                            fmt.Println(v1)
                                        }
                                        _ = v51
                                    }
                                }
                                _ = fi14
                            }
                        case "a":
                            v56 := GenS2{F0: genFn0(-268, 926, (100 < -19)), F1: -904}
                            _ = v56
                        default:
                            if (-32768 > 42) {
                                for fi57 := 0; fi57 < 11; fi57++ {
                                    v1 = (v1 && (1e10 == -2.25))
                                    var v58 uint32 = 2
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    fmt.Println(fi57, v58, v1)
                                    if v1 {
                                        var v59 int8 = 127
                                        _ = v59
                                    } else {
                                        v4 = append(v4, (fi57 | (fi12 >> ((0) & 63))))
                                        v58 = v58
                                    }
                                    _ = fi57
                                    _ = v58
                                }
                            }
                        }
                        var v60 string = ("" + "foo bar")
                        var v61 uint8 = v3
                        _ = fi12
                        _ = v60
                        _ = v61
                    }
                case "\n\t":
                    switch 1 {
                    case 32768:
                        if true {
                            var v62 int = 127
                            if false {
                                for fi63 := 0; fi63 < 17; fi63++ {
                                    for fi64 := 0; fi64 < 16; fi64++ {
                                        fmt.Println(fi63, v1)
                                        var v65 int32 = 2147483647
                                        var v66 int16 = 0
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                    _ = len(v2)
                                    _ = fi63
                                }
                            }
                            fmt.Println(v62, v1, v3)
                            _ = v62
                        }
                    default:
                        switch "hello" {
                        case "hello":
                            if false {
                                if false {
                                    for fi67 := 0; fi67 < 4; fi67++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        fmt.Println(v1, fi67, fi67)
                                        v3 = (v3 / ((127) | 1))
                                        _ = fi67
                                    }
                                }
                                var v68 string = "hello"
                                v69 := GenS5{F0: (v3 + 1), F1: (!true)}
                                _ = v68
                                _ = v69
                            }
                        case "":
                            for fi70 := 0; fi70 < 15; fi70++ {
                                v3 -= (v3 % ((128) | 1))
                                switch int(v3) {
                                case 127:
                                    if v1 {
                                        var v71 bool = v1
                                        for fi72 := 0; fi72 < 7; fi72++ {
                                            var v73 uint32 = 2147483647
                                            switch fi70 {
                                            case 128:
                                                if v71 {
                                                    v74 := GenS9{F0: 42}
                                                    v3 = (v3 >> ((uint(v3)) & 63))
                                                    v75 := GenS4{F0: "foo bar", F1: "\n\t"}
                                                    _ = v74
                                                    _ = v75
                                                }
                                            default:
                                                switch (("" + "") + "") {
                                                case "":
                                                    v1 = ((!(true && (v73 > (v73 % (((v73 >> ((65535) & 63))) | 1))))) && true)
                                                default:
                                                    var v76 int32 = 607
                                                    _ = v76
                                                }
                                            }
                                            delete(v5, v73)
                                            v77 := GenS9{F0: -470}
                                            _ = fi72
                                            _ = v73
                                            _ = v77
                                        }
                                        _ = v71
                                    }
                                }
                                _ = fi70
                            }
                        case "\n\t":
                            var v78 uint16 = 940
                            _ = v78
                        }
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v11
            }
        case "hello":
            for fi79 := 0; fi79 < 11; fi79++ {
                var v80 float64 = 1.5
                v81 := GenS10{F0: 2147483647, F1: -114}
                v80 -= (-2.25 + 1.5)
                _ = fi79
                _ = v80
                _ = v81
            }
        default:
            if (!(!v1)) {
                if true {
                    var v82 int16 = -32768
                    _ = v82
                }
            }
        }
    default:
        v2[65535] = "a"
    }
    v1 = true
    var v83 uint32 = 1
    v1 = (!v1)
    var v84 bool = (v1 || v1)
    var v85 uint16 = 2
    for fi86 := 0; fi86 < 20; fi86++ {
        var v87 uint64 = uint64(v83)
        switch (^(fi86)) {
        case 127:
            _ = len(v4)
        default:
            var v88 int32 = -234
            _ = v88
        }
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v89 int32 = -1
            _ = v89
        } else {
            var v90 bool = true
            _ = v90
        }
        _ = fi86
        _ = v87
    }
    for fi91 := 0; fi91 < 7; fi91++ {
        var v92 float32 = (float32(((3.14159 * ((0.0 + 0.1) / -2.25)) - -2.25)) * (-2.25 + 0.0))
        _ = fi91
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v85
}
