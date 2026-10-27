package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 int64; F1 uint; F2 uint16 }
type GenS11 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -13
    var v3 uint32 = 4294967295
    v4 := []string{("\n\t" + "hello"), ("hello" + "hello")}
    var v5 int16 = 600
    var v6 uint = 4294967295
    v6 = (v6 | (v6 + 774))
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v3, v5)
        var v7 bool = (42 == -705)
        _ = v7
    }
    if ("" == "a") {
        v3 = (v3 ^ v3)
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 float32 = 1.5
            fmt.Println(v9, fi8, v2)
            _ = fi8
            _ = v9
        }
    } else {
        v2 = v2
        v6 -= 42
        switch ("" + ("foo bar" + ("foo bar" + "hello"))) {
        case "a":
            v6 = v6
        case "foo bar":
            if (262 != 42) {
                _ = len(v4)
                if false {
                    switch 4294967295 {
                    case 9223372036854775807:
                        switch ("\n\t" + ("foo bar" + "hello")) {
                        case "hello":
                            var v10 uint64 = 4294967295
                            _ = v10
                        case "\n\t":
                            if false {
                                v11 := GenS1{F0: int64(v6), F1: v5}
                                _ = v11
                            } else {
                                for fi12 := 0; fi12 < 6; fi12++ {
                                    switch "foo bar" {
                                    case "":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "hello":
                                        fmt.Println(fi12)
                                    case "a":
                                        var v13 uint = v6
                                        _ = v13
                                    }
                                    var v14 int32 = 1
                                    v15 := GenS2{F0: 42, F1: 655}
                                    _ = fi12
                                    _ = v14
                                    _ = v15
                                }
                                fmt.Println(v3, v6, v2)
                                var v16 int32 = 1
                                _ = v16
                            }
                        case "foo bar":
                            var v17 bool = false
                            _ = v17
                        default:
                            if (true || (false || ((false || (!((0.1 - 3.14159) <= 0.1))) && (1000 >= (v5 ^ v5))))) {
                                if false {
                                    var v18 uint16 = 408
                                    _ = v18
                                }
                                v6 = v6
                                var v19 bool = (true && true)
                                _ = v19
                            } else {
                                switch "hello" {
                                case "":
                                    fmt.Println(v6, v5, v6)
                                case "\n\t":
                                    v6 -= 4294967295
                                default:
                                    v5 += -1
                                }
                                var v20 float32 = -2.25
                                _ = v20
                            }
                        }
                    case -402:
                        var v21 int8 = (v2 - v2)
                        _ = v21
                    case -9223372036854775808:
                        for fi22 := 0; fi22 < 12; fi22++ {
                            v23 := GenS1{F0: 9223372036854775807, F1: v5}
                            for fi24 := 0; fi24 < 11; fi24++ {
                                switch "hello" {
                                case "a":
                                    v25 := GenS3{F0: 65535, F1: 42}
                                    _ = v25
                                case "foo bar":
                                    fmt.Println(v5, v3)
                                case "":
                                    var v26 int16 = v5
                                    _ = v26
                                }
                                switch "\n\t" {
                                case "":
                                    var v27 float64 = -2.25
                                    _ = v27
                                }
                                _ = fi24
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                    v5 = v5
                    switch "\n\t" {
                    case "a":
                        v28 := GenS3{F0: 2, F1: uint16(v3)}
                        _ = v28
                    case "\n\t":
                        v29 := GenS2{F0: 255, F1: v3}
                        _ = v29
                    case "":
                        if true {
                            v6 = 42
                            var v30 uint8 = 44
                            if false {
                                if true {
                                    v5 = v5
                                    var v31 uint64 = 2
                                    v32 := v1[16]
                                    _ = v31
                                    _ = v32
                                }
                            } else {
                                for fi33 := 0; fi33 < 2; fi33++ {
                                    _ = len(v1)
                                    _ = len(v1)
                                    v2 = 127
                                    _ = fi33
                                }
                                if true {
                                    for fi34 := 0; fi34 < 10; fi34++ {
                                        v30 -= v30
                                        for fi35 := 0; fi35 < 15; fi35++ {
                                            if ((!("\n\t" >= "a")) || (35 != 32767)) {
                                                var v36 uint64 = 1
                                                _ = v36
                                            }
                                            _ = fi35
                                        }
                                        v5 -= v5
                                        _ = fi34
                                    }
                                    for fi37 := 0; fi37 < 1; fi37++ {
                                        var v38 int8 = -128
                                        if (("" + (("foo bar" + "\n\t") + ("a" + "a"))) > "\n\t") {
                                            var v39 uint32 = (^(v3))
                                            for fi40 := 0; fi40 < 5; fi40++ {
                                                switch fi37 {
                                                case -194:
                                                    v41 := GenS4{F0: 255}
                                                    _ = v41
                                                }
                                                if true {
                                                    delete(v1, (-(fi40)))
                                                    fmt.Println(v38)
                                                    if true {
                                                        var v42 int16 = (-(v5))
                                                        switch (("hello" + "hello") + ("" + "hello")) {
                                                        case "foo bar":
                                                            v43 := GenS5{F0: "", F1: v6, F2: int64(v2)}
                                                            _ = v43
                                                        case "a":
                                                            if false {
                                                                switch (fi37 | fi37) {
                                                                case 128:
                                                                    var v44 uint = 65535
                                                                    _ = v44
                                                                case -23:
                                                                    var v45 uint32 = (v3 - v3)
                                                                    _ = v45
                                                                case 42:
                                                                    if (((true && false) && (v30 == (v30 + (v30 % ((237) | 1))))) && true) {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v46 := GenS6{F0: 1}
                                                                        v47 := GenS6{F0: (v6 ^ v6)}
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                default:
                                                                    if false {
                                                                        var v48 uint = v6
                                                                        _ = v48
                                                                    } else {
                                                                        fmt.Println(fi40, v2)
                                                                        switch (^(fi40)) {
                                                                        case -1:
                                                                            switch fi40 {
                                                                            case -110:
                                                                                v49 := GenS3{F0: uint16(v30), F1: 2}
                                                                                _ = v49
                                                                            }
                                                                        case 1:
                                                                            v39 += (^(v39))
                                                                        case 127:
                                                                            v42 = v5
                                                                        }
                                                                    }
                                                                }
                                                                var v50 uint32 = 42
                                                                _ = v50
                                                            }
                                                        case "hello":
                                                            var v51 uint64 = 0
                                                            fmt.Println(fi37, v3, v2)
                                                            v52 := v1[fi37]
                                                            _ = v51
                                                            _ = v52
                                                        default:
                                                            var v53 int32 = -2147483648
                                                            fmt.Println(fi40)
                                                            v42 = 42
                                                            _ = v53
                                                        }
                                                        if ((fi37 % ((-9223372036854775808) | 1)) < (fi40 ^ fi37)) {
                                                            var v54 uint = (^(v6))
                                                            _ = v54
                                                        }
                                                        v2 = (v38 + (v38 | v2))
                                                        _ = v42
                                                    }
                                                    fmt.Println(v5, v39)
                                                } else {
                                                    v5 += 0
                                                    fmt.Println(v30, v30, v38)
                                                }
                                                v1[(fi40 << ((v6) & 63))] = (v3 / ((uint32(v6)) | 1))
                                                _ = fi40
                                            }
                                            for fi55 := 0; fi55 < 17; fi55++ {
                                                var v56 int16 = v5
                                                fmt.Println(v30)
                                                if (true && true) {
                                                    var v57 uint32 = (v39 * v39)
                                                    v1[-2147483648] = v57
                                                    var v58 float64 = (float64(1e10) / -2.25)
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    if true {
                                                        v3 = (^(v3))
                                                    }
                                                    fmt.Println(v39, fi55)
                                                    v59 := GenS7{F0: -592}
                                                    v39 -= v3
                                                    if true {
                                                        if false {
                                                            v60 := GenS4{F0: (v6 << ((v6) & 63))}
                                                            _ = v60
                                                        } else {
                                                            var v61 uint8 = uint8(v3)
                                                            _ = v61
                                                        }
                                                        fmt.Println(fi37)
                                                        var v62 int8 = v38
                                                        _ = v62
                                                    }
                                                    _ = v59
                                                }
                                                v30 -= 2
                                                v2 -= v2
                                                _ = fi55
                                                _ = v56
                                            }
                                            _ = v39
                                        }
                                        _ = fi37
                                        _ = v38
                                    }
                                    if (v5 <= v5) {
                                        switch int(v5) {
                                        case 42:
                                            if false {
                                                v63 := GenS8{F0: -103}
                                                v64 := GenS7{F0: 0}
                                                _ = v63
                                                _ = v64
                                            } else {
                                                fmt.Println(v5, v3)
                                                if ((v5 << ((v6) & 63)) > (v5 >> ((v6) & 63))) {
                                                    v2 = (v2 ^ 22)
                                                    if (42 > 32767) {
                                                        fmt.Println(v2, v3, v3)
                                                        for fi65 := 0; fi65 < 19; fi65++ {
                                                            v3 = (^(v3))
                                                            v66 := GenS9{F0: 2, F1: "hello"}
                                                            var v67 int = fi65
                                                            _ = fi65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    }
                                                    v3 -= (v3 ^ (v3 * 42))
                                                }
                                                var v68 uint8 = v30
                                                _ = v68
                                            }
                                        }
                                    } else {
                                        var v69 float64 = -2.25
                                        v30 += v30
                                        _ = v69
                                    }
                                } else {
                                    var v70 float32 = 1e10
                                    v3 += 722
                                    var v71 uint32 = (v3 % ((v3) | 1))
                                    _ = v70
                                    _ = v71
                                }
                            }
                            _ = v30
                        }
                    }
                }
            }
        default:
            fmt.Println(v6, v3, v6)
        }
    }
    for fi72 := 0; fi72 < 14; fi72++ {
        var v73 int16 = v5
        _ = fi72
        _ = v73
    }
    v2 -= (v2 & (v2 >> (((v6 | 42)) & 63)))
    switch "\n\t" {
    case "":
        fmt.Println(v3, v2, v5)
    case "\n\t":
        v74 := GenS10{F0: -9223372036854775808, F1: 2, F2: 2}
        _ = v74
    }
    v1[4294967295] = 1
    var v75 uint8 = 224
    if (true || true) {
        switch -1 {
        case -9223372036854775808:
            if ((!(-796 >= -2147483648)) && (float64(3.14159) >= 3.14159)) {
                v76 := GenS11{F0: 108, F1: v2}
                _ = v76
            } else {
                v1[0] = v3
            }
        case -927:
            fmt.Println(v6, v2)
        }
        var v77 int = -429
        v4 = append(v4, ((("hello" + "\n\t") + "") + "hello"))
        _ = v77
    } else {
        v78 := GenS9{F0: 42, F1: "a"}
        _ = v78
    }
    var v79 uint8 = v75
    switch -1 {
    case -794:
        fmt.Println(v75, v5)
    }
    v80 := GenS5{F0: ("\n\t" + "hello"), F1: (^(v6)), F2: 9223372036854775807}
    if true {
        v81 := GenS10{F0: 568, F1: 655, F2: 2}
        if (42 < 0) {
            v3 = 478
            switch 128 {
            case 724:
                v5 -= v5
            case -9223372036854775808:
                var v82 int32 = 0
                _ = v82
            default:
                var v83 int64 = 672
                _ = v83
            }
            switch ("a" + ("" + (("" + "hello") + "foo bar"))) {
            case "foo bar":
                if (127 != 100) {
                    var v84 uint32 = 2147483648
                    for fi85 := 0; fi85 < 9; fi85++ {
                        var v86 bool = true
                        v75 = v75
                        var v87 int = fi85
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                    _ = v84
                }
            }
        }
        _ = v81
    }
    var v88 uint64 = 4294967295
    v79 = 2
    v88 = v88
    fmt.Println(v5, v88)
    if (false || false) {
        fmt.Println(v3)
        v89 := GenS8{F0: 79}
        v90 := GenS9{F0: v88, F1: (("foo bar" + "") + "a")}
        _ = v89
        _ = v90
    } else {
        v3 = (^(v3))
        var v91 float32 = (1e10 / 3.14159)
        _ = v91
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    fmt.Println("v79:", v79)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v79
    _ = v80
    _ = v88
}
