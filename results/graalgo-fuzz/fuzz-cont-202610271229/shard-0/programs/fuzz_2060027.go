package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 bool; F1 int16; F2 uint16 }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 bool; F1 uint64 }
type GenS12 struct { F0 int64; F1 int64 }
func genFn0(p0 string) int {
    v5 := GenS1{F0: "hello"}
    v6 := []int{2147483647, -527}
    v7 := GenS1{F0: p0}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -343
}
func genFn1(p0 uint32, p1 bool) int64 {
    if (4294967295 != 9223372036854775807) {
        switch 9223372036854775807 {
        case 32767:
            p1 = true
        case 2147483647:
            var v8 int32 = 401
            _ = v8
        }
        var v9 uint32 = p0
        p1 = p1
        _ = v9
    }
    p1 = p1
    var v10 uint64 = 42
    for fi11 := 0; fi11 < 18; fi11++ {
        v10 += (^(v10))
        p0 = p0
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    if false {
        var v13 uint16 = 65535
        fmt.Println(p1, v13, p0)
        _ = v13
    } else {
        var v14 uint16 = 255
        var v15 float64 = float64(1.5)
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    return 875
}

func main() {
    v1 := []uint32{65535, 1, 828, 41}
    var v2 int16 = -1
    var v3 uint16 = 65535
    var v4 int16 = (v2 + v2)
    v1 = append(v1, 2147483648)
    switch "\n\t" {
    case "a":
        var v16 uint16 = 42
        _ = v16
    case "hello":
        for fi17 := 0; fi17 < 17; fi17++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                if (!true) {
                    v1 = append(v1, 95)
                } else {
                    switch (fi17 & fi17) {
                    case 32767:
                        v18 := GenS1{F0: "a"}
                        _ = v18
                    default:
                        var v19 string = ""
                        _ = v19
                    }
                    v2 = v4
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            } else {
                v3 -= (v3 | (v3 - v3))
                switch fi17 {
                case 4294967295:
                    if true {
                        if false {
                            v3 = v3
                            v20 := GenS1{F0: (("foo bar" + "hello") + ("\n\t" + (("\n\t" + "a") + "foo bar")))}
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 2; fi21++ {
                                fmt.Println(fi17, v4, fi21)
                                if false {
                                    if false {
                                        if true {
                                            v22 := GenS1{F0: "\n\t"}
                                            if ((("\n\t" + "") == ("foo bar" + "hello")) || (-9223372036854775808 >= 0)) {
                                                v4 -= v4
                                                v1 = append(v1, 314)
                                            }
                                            _ = v22
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = 211
                                            }
                                            var v23 float64 = -2.25
                                            _ = v23
                                        }
                                    }
                                } else {
                                    v24 := GenS2{F0: 127, F1: (v3 + v3), F2: false}
                                    _ = v24
                                }
                                v3 = (^(v3))
                                var v25 int32 = -299
                                _ = fi21
                                _ = v25
                            }
                            var v26 int8 = -73
                            v27 := GenS2{F0: 128, F1: (v3 >> ((255) & 63)), F2: true}
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS1{F0: "a"}
                        _ = v28
                    } else {
                        if (!false) {
                            switch (((("" + "") + "") + "") + "a") {
                            case "":
                                fmt.Println(fi17, fi17, v3)
                            case "foo bar":
                                v2 += v4
                            }
                            if false {
                                fmt.Println(fi17)
                                v2 -= (v2 % ((-317) | 1))
                                var v29 string = ("" + "")
                                _ = v29
                            } else {
                                v30 := GenS3{F0: false, F1: 1}
                                v2 = -32768
                                if (222 < 128) {
                                    var v31 uint32 = 4294967295
                                    _ = v31
                                }
                                _ = v30
                            }
                        } else {
                            v32 := GenS1{F0: ""}
                            var v33 bool = false
                            _ = v32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 7; fi34++ {
                            v35 := GenS3{F0: true, F1: 1}
                            v36 := GenS3{F0: ((539 != (v3 & (^(v3)))) || true), F1: 42}
                            if (32767 != (v3 | (^(v3)))) {
                                v4 = v2
                            } else {
                                v1 = append(v1, 836)
                                var v37 int64 = int64(fi17)
                                var v38 uint64 = 834
                                _ = v37
                                _ = v38
                            }
                            var v39 uint = uint(v3)
                            var v40 uint8 = 127
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v39
                            _ = v40
                        }
                        v41 := GenS1{F0: ("\n\t" + "foo bar")}
                        _ = v41
                    }
                case 9223372036854775807:
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                default:
                    var v42 float32 = 0.1
                    _ = v42
                }
                if false {
                    var v43 float64 = float64(float32(0.0))
                    _ = v43
                } else {
                    if false {
                        var v44 string = "hello"
                        var v45 uint32 = 65535
                        v1 = append(v1, v45)
                        _ = v44
                        _ = v45
                    } else {
                        v3 = (v3 / (((v3 >> ((65535) & 63))) | 1))
                    }
                    v3 = v3
                    fmt.Println(v2, fi17, v4)
                }
            }
            _ = len(v1)
            _ = fi17
        }
    case "foo bar":
        v3 = (v3 % ((0) | 1))
    default:
        if (-9223372036854775808 > 523) {
            v4 = (^(v4))
            var v46 uint32 = 178
            _ = v46
        }
    }
    _ = len(v1)
    var v47 uint16 = 255
    for fi48 := 0; fi48 < 5; fi48++ {
        switch ("a" + ("a" + "hello")) {
        case "hello":
            v2 += 42
        }
        v49 := GenS2{F0: 151, F1: v3, F2: false}
        _ = fi48
        _ = v49
    }
    switch 32768 {
    case -215:
        v50 := GenS4{F0: 4294967295}
        _ = v50
    case 1:
        for fi51 := 0; fi51 < 7; fi51++ {
            if false {
                switch 1 {
                case 0:
                    var v52 uint64 = 4294967295
                    _ = v52
                case 9223372036854775807:
                    v3 = (^(v3))
                case 4294967295:
                    for fi53 := 0; fi53 < 4; fi53++ {
                        if (v3 == v3) {
                            v54 := GenS3{F0: true, F1: 0}
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                            _ = v54
                        }
                        v55 := GenS4{F0: 870}
                        var v56 string = "hello"
                        _ = fi53
                        _ = v55
                        _ = v56
                    }
                default:
                    v2 += v2
                }
            } else {
                v57 := GenS1{F0: ((("foo bar" + ("\n\t" + "a")) + "") + "a")}
                if false {
                    for fi58 := 0; fi58 < 19; fi58++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi58
                    }
                    var v59 float32 = 1e10
                    if (!true) {
                        fmt.Println(v3, v3, v59)
                        switch "\n\t" {
                        case "foo bar":
                            v59 -= (v59 + (v59 / 1.5))
                        default:
                            var v60 int32 = 1
                            _ = v60
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = v59
                } else {
                    v61 := GenS3{F0: true, F1: 127}
                    if true {
                        for fi62 := 0; fi62 < 4; fi62++ {
                            v63 := GenS2{F0: 127, F1: (v47 * (v47 | 65535)), F2: false}
                            if true {
                                v64 := GenS5{F0: (fi62 + (fi51 - -1))}
                                v47 = v47
                                v1 = append(v1, 2147483648)
                                _ = v64
                            } else {
                                v65 := GenS2{F0: 255, F1: 447, F2: (!(-2147483648 > 0))}
                                if (32767 == (v3 * (v3 % (((^(v3))) | 1)))) {
                                    v47 = v3
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v3 += v47
                                    v3 = (v47 * v3)
                                    switch (fi62 % (((fi62 * genFn0("a"))) | 1)) {
                                    case 4294967295:
                                        v4 = v4
                                    case 635:
                                        if (!false) {
                                            if true {
                                                if (!false) {
                                                    if false {
                                                        v66 := GenS6{F0: (fi62 * fi51), F1: 9223372036854775807}
                                                        switch (((("hello" + "a") + "a") + "\n\t") + "") {
                                                        case "\n\t":
                                                            fmt.Println(v3)
                                                        case "foo bar":
                                                            v47 -= v3
                                                            if ((!false) || false) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 4294967295
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = 1
                                                                }
                                                                var v67 int16 = (v4 / ((-1) | 1))
                                                                _ = v67
                                                            }
                                                            switch "" {
                                                            case "a":
                                                                v3 -= v47
                                                            case "":
                                                                var v68 int32 = -681
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v68
                                                            }
                                                        case "":
                                                            fmt.Println(v2, v4, v3)
                                                        default:
                                                            v2 += (v4 >> ((4294967295) & 63))
                                                        }
                                                        _ = v66
                                                    } else {
                                                        var v69 float64 = (-2.25 - 1.5)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v69 -= float64(1.5)
                                                        _ = v69
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = 42
                                                    }
                                                }
                                                v4 = v2
                                                var v70 int = genFn0((("foo bar" + ("" + "hello")) + "foo bar"))
                                                _ = v70
                                            }
                                        }
                                    }
                                    var v71 uint16 = v47
                                    switch (fi51 >> ((196) & 63)) {
                                    case 32767:
                                        v72 := GenS1{F0: "hello"}
                                        _ = v72
                                    default:
                                        if true {
                                            if false {
                                                var v73 int64 = 499
                                                v1 = append(v1, 338)
                                                _ = v73
                                            }
                                            _ = len(v1)
                                        } else {
                                            switch genFn0("\n\t") {
                                            case 2147483647:
                                                var v74 bool = true
                                                fmt.Println(v74, fi51, v74)
                                                _ = v74
                                            case 127:
                                                v1 = append(v1, 1)
                                            case -9223372036854775808:
                                                switch (("hello" + "a") + "hello") {
                                                case "":
                                                    v71 -= (v47 & v71)
                                                default:
                                                    v71 = (^(v3))
                                                }
                                            default:
                                                var v75 int16 = -933
                                                _ = v75
                                            }
                                        }
                                    }
                                    _ = v71
                                }
                                if false {
                                    switch genFn0("a") {
                                    case -9223372036854775808:
                                        v3 = v3
                                    case 9223372036854775807:
                                        v4 -= v4
                                    case -718:
                                        v3 += v47
                                    default:
                                        v3 = v3
                                    }
                                    var v76 uint64 = 0
                                    _ = v76
                                } else {
                                    v3 = (v47 << ((4294967295) & 63))
                                }
                                _ = v65
                            }
                            v77 := GenS4{F0: 42}
                            _ = fi62
                            _ = v63
                            _ = v77
                        }
                        for fi78 := 0; fi78 < 2; fi78++ {
                            if (1.5 <= 1e10) {
                                if false {
                                    switch ("hello" + "") {
                                    case "\n\t":
                                        var v79 float32 = 1.5
                                        _ = v79
                                    case "hello":
                                        if ((3.14159 / 1.5) != -2.25) {
                                            _ = len(v1)
                                            if (!true) {
                                                if len(v1) > 0 {
                                                    v1[0] = 907
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    v4 = v2
                                                case "":
                                                    v47 = 987
                                                default:
                                                    var v80 float32 = 0.1
                                                    _ = v80
                                                }
                                                if true {
                                                    var v81 int = (fi78 % ((fi51) | 1))
                                                    _ = v81
                                                } else {
                                                    var v82 float64 = 1e10
                                                    v3 += (v47 & v47)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v82
                                                }
                                            } else {
                                                var v83 float32 = (-2.25 + 0.0)
                                                switch "a" {
                                                case "hello":
                                                    if false {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        var v84 uint8 = 42
                                                        _ = v84
                                                    }
                                                default:
                                                    v85 := GenS7{F0: 695, F1: genFn1(2147483648, false)}
                                                    _ = v85
                                                }
                                                switch ("\n\t" + ("foo bar" + ("" + ("foo bar" + "")))) {
                                                case "foo bar":
                                                    var v86 int32 = 57
                                                    _ = v86
                                                default:
                                                    fmt.Println(fi78, fi51, fi78)
                                                }
                                                if (689 <= 0) {
                                                    fmt.Println(v47)
                                                    var v87 int16 = v2
                                                    _ = v87
                                                }
                                                _ = v83
                                            }
                                            v88 := GenS7{F0: 0, F1: 42}
                                            _ = v88
                                        } else {
                                            var v89 string = ("" + "a")
                                            _ = v89
                                        }
                                    }
                                    var v90 float64 = 1e10
                                    var v91 string = "\n\t"
                                    var v92 int32 = -671
                                    _ = v90
                                    _ = v91
                                    _ = v92
                                }
                                v1 = append(v1, 65535)
                            }
                            v2 = -32768
                            var v93 int64 = -157
                            _ = fi78
                            _ = v93
                        }
                        fmt.Println(v3, v3)
                    } else {
                        v94 := GenS5{F0: fi51}
                        _ = v94
                    }
                    for fi95 := 0; fi95 < 5; fi95++ {
                        if true {
                            v96 := GenS8{F0: 363, F1: 42}
                            if (2 >= 42) {
                                if false {
                                    v2 += (v2 | v2)
                                }
                                fmt.Println(fi95)
                                v2 = v4
                            } else {
                                v97 := GenS6{F0: fi95, F1: 7}
                                _ = v97
                            }
                            _ = v96
                        }
                        _ = fi95
                    }
                    _ = v61
                }
                _ = v57
            }
            var v98 uint16 = v47
            if false {
                _ = len(v1)
                var v99 int8 = 1
                v4 = v2
                _ = v99
            } else {
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v98 = v47
                    var v100 int32 = -1
                    _ = v100
                }
            }
            _ = fi51
            _ = v98
        }
    case 9223372036854775807:
        v2 += (v4 & 32767)
    }
    for fi101 := 0; fi101 < 7; fi101++ {
        switch "foo bar" {
        case "hello":
            v102 := GenS9{F0: true, F1: v2, F2: 2}
            _ = v102
        case "\n\t":
            var v103 int = fi101
            _ = v103
        case "":
            v104 := GenS10{F0: 42, F1: true}
            _ = v104
        }
        var v105 int32 = 1
        _ = fi101
        _ = v105
    }
    v106 := GenS9{F0: ((2 > 18446744073709551615) && (((float32(0.1) + float32(3.14159)) != 1.5) || true)), F1: (v2 * v4), F2: 2}
    if false {
        fmt.Println(v3, v2, v47)
        v107 := GenS11{F0: (!((1 >= -37) || true)), F1: 42}
        if len(v1) > 0 {
            v1[0] = 2
        }
        _ = v107
    }
    v4 += (v2 / (((v2 >> ((781) & 63))) | 1))
    fmt.Println(v4, v4, v4)
    v2 += v4
    v3 += v47
    switch "\n\t" {
    case "":
        fmt.Println(v4)
    case "foo bar":
        var v108 int8 = -128
        _ = v108
    default:
        v109 := GenS1{F0: "a"}
        _ = v109
    }
    for fi110 := 0; fi110 < 2; fi110++ {
        var v111 int32 = 1
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        v112 := GenS12{F0: 768, F1: -1}
        _ = fi110
        _ = v111
        _ = v112
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v47
}
