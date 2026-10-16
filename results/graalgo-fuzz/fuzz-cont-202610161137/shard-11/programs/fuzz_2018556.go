package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 int8; F1 int; F2 int32 }
type GenS4 struct { F0 int8; F1 int64; F2 uint16 }
type GenS5 struct { F0 int32; F1 int; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int16; F1 string }

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 4294967295
    var v3 uint8 = 128
    for fi4 := 0; fi4 < 14; fi4++ {
        v2 = v2
        _ = fi4
    }
    v5 := GenS1{F0: 1, F1: (v2 - 347)}
    var v6 int64 = 188
    switch ("a" + "") {
    case "":
        v7 := GenS1{F0: -483, F1: 0}
        _ = v7
    }
    switch "\n\t" {
    case "":
        if (-587 <= -978) {
            for fi8 := 0; fi8 < 13; fi8++ {
                v2 = 18446744073709551615
                v3 = (v3 >> ((18446744073709551615) & 63))
                _ = fi8
            }
        } else {
            var v9 uint32 = 788
            v10 := GenS1{F0: 32767, F1: v2}
            v11 := GenS1{F0: -171, F1: (v2 | (v2 + 18446744073709551615))}
            _ = v10
            _ = v11
            _ = v9
        }
    case "a":
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(v6)
            _ = fi12
        }
    case "\n\t":
        fmt.Println(v6, v6, v2)
    }
    switch -257 {
    case 1:
        v2 += 18446744073709551615
    case -956:
        fmt.Println(v6, v6, v2)
    default:
        v1 -= v1
    }
    v13 := []int{1, 2147483647, 2147483647, 127}
    for fi14 := 0; fi14 < 15; fi14++ {
        switch fi14 {
        case -2147483648:
            switch "hello" {
            case "foo bar":
                fmt.Println(v2, v1, v1)
            }
        case 127:
            if false {
                v15 := GenS2{F0: (v1 ^ -704), F1: (("" + "\n\t") + "\n\t"), F2: 42}
                for fi16 := 0; fi16 < 3; fi16++ {
                    fmt.Println(v6)
                    if false {
                        var v17 float64 = ((float64(1.5) / 1.5) + (0.1 * ((3.14159 / -2.25) / (0.1 + (3.14159 * ((1.5 + (0.1 * 0.0)) + (3.14159 + float64((1e10 + 3.14159)))))))))
                        switch "hello" {
                        case "\n\t":
                            v17 = (((((v17 / ((v17 / v17) / 0.1)) / (-2.25 - 0.0)) * v17) / v17) - v17)
                        default:
                            v13 = append(v13, fi14)
                        }
                        var v18 uint32 = 2147483647
                        _ = v17
                        _ = v18
                    }
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        var v19 int8 = 115
                        _ = v19
                    case "hello":
                        var v20 int8 = 53
                        _ = v20
                    }
                    _ = fi16
                }
                var v21 uint16 = 1
                _ = v15
                _ = v21
            } else {
                switch fi14 {
                case 1:
                    var v22 int8 = 126
                    _ = v22
                case 32767:
                    fmt.Println(fi14, v3)
                case 127:
                    v6 = v6
                }
            }
        default:
            v23 := GenS1{F0: 0, F1: 42}
            _ = v23
        }
        _ = fi14
    }
    if (!false) {
        fmt.Println(v2)
        v24 := GenS2{F0: (v1 / ((-459) | 1)), F1: "hello", F2: 255}
        for fi25 := 0; fi25 < 4; fi25++ {
            for fi26 := 0; fi26 < 18; fi26++ {
                switch fi25 {
                case 585:
                    v27 := GenS3{F0: 56, F1: fi25, F2: 0}
                    _ = v27
                }
                var v28 uint32 = 65535
                var v29 int8 = -1
                if false {
                    var v30 int32 = (-(v1))
                    switch "\n\t" {
                    case "\n\t":
                        v3 = 127
                    }
                    _ = v30
                }
                _ = fi26
                _ = v28
                _ = v29
            }
            _ = fi25
        }
        _ = v24
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 uint32 = 42
        _ = fi31
        _ = v32
    }
    if (((!false) && (true && false)) && ((1e10 - 1e10) < 3.14159)) {
        v3 = (v3 / (((^(v3))) | 1))
    } else {
        switch "hello" {
        case "\n\t":
            v33 := GenS4{F0: -49, F1: (^(v6)), F2: 65535}
            _ = v33
        case "":
            fmt.Println(v3, v1)
        default:
            for fi34 := 0; fi34 < 19; fi34++ {
                switch fi34 {
                case -503:
                    if false {
                        v3 = v3
                    }
                case -1:
                    for fi35 := 0; fi35 < 6; fi35++ {
                        v6 -= (-(v6))
                        v36 := GenS5{F0: v1, F1: 42, F2: (v6 ^ -9223372036854775808)}
                        _ = fi35
                        _ = v36
                    }
                default:
                    var v37 uint = uint(v3)
                    _ = v37
                }
                for fi38 := 0; fi38 < 6; fi38++ {
                    v6 = 0
                    v39 := GenS5{F0: v1, F1: fi34, F2: v6}
                    v1 = -126
                    _ = fi38
                    _ = v39
                }
                var v40 string = "a"
                _ = fi34
                _ = v40
            }
        }
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        v1 -= (-(v1))
        if ((v1 | (v1 % ((-2147483648) | 1))) >= v1) {
            v13 = append(v13, (fi41 & fi41))
            var v42 uint32 = 0
            if (903 > int16(v1)) {
                for fi43 := 0; fi43 < 20; fi43++ {
                    if false {
                        v3 = 2
                    }
                    _ = fi43
                }
                for fi44 := 0; fi44 < 7; fi44++ {
                    switch "\n\t" {
                    case "foo bar":
                        var v45 float32 = 1.5
                        _ = v45
                    case "a":
                        v46 := GenS6{F0: ("" + "hello")}
                        _ = v46
                    default:
                        var v47 string = "\n\t"
                        _ = v47
                    }
                    var v48 int = fi44
                    var v49 uint32 = 2
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = fi44
                    _ = v48
                    _ = v49
                }
            } else {
                for fi50 := 0; fi50 < 15; fi50++ {
                    if false {
                        v3 = 1
                        v2 = 1
                    } else {
                        fmt.Println(v3, v6, fi50)
                    }
                    if (-2.25 <= 3.14159) {
                        if (!false) {
                            var v51 int16 = 0
                            fmt.Println(v2)
                            _ = v51
                        }
                    } else {
                        v3 = (v3 >> ((4294967295) & 63))
                        if true {
                            var v52 uint16 = uint16(v42)
                            var v53 uint16 = 0
                            if (false && false) {
                                switch "hello" {
                                case "":
                                    if true {
                                        v6 = v6
                                        if ((v6 ^ 0) < v6) {
                                            var v54 string = ""
                                            _ = v54
                                        } else {
                                            if (v3 <= (v3 ^ v3)) {
                                                v42 = 4294967295
                                                if ((v1 < (v1 - int32(v3))) && ("\n\t" >= ("hello" + "\n\t"))) {
                                                    v6 = v6
                                                    _ = len(v13)
                                                    v42 += v42
                                                } else {
                                                    switch "" {
                                                    case "hello":
                                                        var v55 uint16 = 1
                                                        _ = v55
                                                    }
                                                }
                                                fmt.Println(v1, fi41)
                                            }
                                        }
                                        if len(v13) > 0 {
                                            v13[0] = 26
                                        }
                                        v3 = v3
                                    } else {
                                        var v56 string = (("hello" + "foo bar") + "")
                                        fmt.Println(v1, fi50, v53)
                                        if (v42 != (v42 + (v42 | (v42 * v42)))) {
                                            if len(v13) > 0 {
                                                v13[0] = (fi50 + int(v3))
                                            }
                                            if ((v1 + v1) == v1) {
                                                var v57 int32 = int32(fi41)
                                                switch "a" {
                                                case "hello":
                                                    var v58 int8 = 1
                                                    _ = v58
                                                case "\n\t":
                                                    var v59 int16 = -410
                                                    v60 := GenS5{F0: v57, F1: 4294967295, F2: v6}
                                                    fmt.Println(fi50)
                                                    _ = v59
                                                    _ = v60
                                                case "a":
                                                    var v61 int32 = 0
                                                    _ = v61
                                                default:
                                                    if false {
                                                        var v62 uint8 = 255
                                                        var v63 uint16 = 32767
                                                        if len(v13) > 0 {
                                                            v13[0] = fi50
                                                        }
                                                        if len(v13) > 0 {
                                                            v13[0] = -1
                                                        }
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                                _ = v57
                                            }
                                        } else {
                                            if true {
                                                if true {
                                                    switch (v56 + v56) {
                                                    case "\n\t":
                                                        _ = len(v13)
                                                    case "hello":
                                                        v52 -= 32767
                                                    default:
                                                        var v64 int64 = -215
                                                        _ = v64
                                                    }
                                                    if (v42 < (v42 + (v42 + 336))) {
                                                        var v65 bool = true
                                                        _ = v65
                                                    }
                                                    v56 = ("\n\t" + "hello")
                                                }
                                                if len(v13) > 0 {
                                                    v13[0] = fi41
                                                }
                                                var v66 string = "\n\t"
                                                _ = v66
                                            }
                                            var v67 uint64 = 930
                                            _ = v67
                                        }
                                        _ = v56
                                    }
                                case "hello":
                                    v2 = 18446744073709551615
                                }
                                if (fi50 > 777) {
                                    if (!(127 > v3)) {
                                        if ((-(v1)) < (v1 * 42)) {
                                            fmt.Println(fi41)
                                        } else {
                                            v6 = v6
                                        }
                                        var v68 uint8 = 42
                                        var v69 uint32 = 1
                                        if (true || (0 >= -32768)) {
                                            v53 -= (v53 + v53)
                                            var v70 uint64 = uint64(v69)
                                            var v71 int16 = -1
                                            _ = v70
                                            _ = v71
                                        }
                                        _ = v68
                                        _ = v69
                                    }
                                }
                                v72 := GenS7{F0: v42, F1: fi41}
                                _ = v72
                            } else {
                                var v73 bool = (false || false)
                                v6 = (v6 % ((42) | 1))
                                _ = v73
                            }
                            _ = v52
                            _ = v53
                        } else {
                            v74 := GenS8{F0: int16(v1), F1: ("a" + "hello")}
                            v6 += (v6 + (v6 & v6))
                            v2 -= (^(v2))
                            _ = v74
                        }
                    }
                    _ = fi50
                }
                fmt.Println(v6)
            }
            _ = v42
        }
        if false {
            var v75 int16 = 154
            _ = v75
        }
        _ = fi41
    }
    if true {
        v6 = (v6 / ((-1) | 1))
        v76 := GenS1{F0: -68, F1: v2}
        v6 -= int64(v3)
        _ = v76
    }
    v3 += (v3 >> ((42) & 63))
    v2 = (^(v2))
    for fi77 := 0; fi77 < 11; fi77++ {
        v3 += 88
        _ = fi77
    }
    fmt.Println(v3, v1, v3)
    v78 := GenS7{F0: 65535, F1: 32768}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v78
}
