package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 int; F2 int8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint8; F1 int }
func genFn0(p0 uint8) uint16 {
    var v7 uint64 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(p0, fi8)
        fmt.Println(v7, v7)
        _ = fi8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{(fi9 << ((65535) & 63)), -9223372036854775808, (fi9 << ((255) & 63))}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    return uint16(v7)
}
func genFn1(p0 uint64, p1 int8) int8 {
    v11 := []int{32767, int(p1), 32767, -506}
    if false {
        if false {
            p0 -= (^(p0))
            var v12 string = (("\n\t" + "foo bar") + "\n\t")
            p0 += p0
            _ = v12
        }
        var v13 string = "\n\t"
        p0 += 2
        _ = v13
    }
    var v14 int = 32768
    v14 = 128
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return -128
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 string = ("hello" + ("hello" + ("foo bar" + "\n\t")))
    var v4 int64 = -9223372036854775808
    v5 := []int{9223372036854775807}
    v6 := []uint32{61}
    var v15 int8 = -128
    var v16 int = 42
    switch int(v15) {
    case -9223372036854775808:
        v17 := GenS1{F0: 32767, F1: 42, F2: -774}
        _ = v17
    }
    switch (v16 * (-(v16))) {
    case 128:
        v18 := GenS2{F0: (!(-702 < 0)), F1: 127}
        _ = v18
    default:
        if (((1 != 34) && (!((0.1 >= 0.1) || true))) && true) {
            var v19 uint32 = 4294967295
            for fi20 := 0; fi20 < 18; fi20++ {
                v3 = (("" + "\n\t") + v3)
                _ = fi20
            }
            _ = len(v2)
            _ = v19
        }
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v4 -= v4
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int64 = (v4 << ((363) & 63))
            switch (v16 - (fi21 & fi21)) {
            case 784:
                v24 := GenS3{F0: 128, F1: 0}
                _ = v24
            }
            v16 -= 32767
            _ = fi22
            _ = v23
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 int32 = -2147483648
            _ = fi25
            _ = v26
        }
        _ = fi21
    }
    switch v3 {
    case "\n\t":
        v2["foo bar"] = 2
    case "a":
        var v27 int64 = 9223372036854775807
        _ = v27
    case "":
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := v1[v3]
            v4 = (^(v4))
            switch ("hello" + (v3 + v3)) {
            case "a":
                var v30 uint = 1
                _ = v30
            case "foo bar":
                fmt.Println(v15)
            default:
                for fi31 := 0; fi31 < 13; fi31++ {
                    if (1 <= 1) {
                        v32 := GenS4{F0: true, F1: 42}
                        var v33 uint = 0
                        _ = v32
                        _ = v33
                    }
                    _ = fi31
                }
            }
            _ = fi28
            _ = v29
        }
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v1["a"] = 146
        var v35 bool = true
        _ = fi34
        _ = v35
    }
    if ((-39 < (v15 >> ((508) & 63))) || ((^(v15)) <= 127)) {
        v3 = v3
    }
    var v36 int64 = (v4 | -9223372036854775808)
    switch (v3 + v3) {
    case "foo bar":
        var v37 float64 = 1e10
        _ = v37
    case "\n\t":
        switch v3 {
        case "hello":
            var v38 int = v16
            _ = v38
        case "":
            var v39 string = v3
            _ = v39
        case "foo bar":
            fmt.Println(v36)
        }
    default:
        v15 = v15
    }
    switch -143 {
    case 0:
        v40 := GenS4{F0: true, F1: 42}
        _ = v40
    case -626:
        v6 = append(v6, 64)
    case 882:
        fmt.Println(v3, v15, v4)
    }
    v3 = (("\n\t" + (v3 + "foo bar")) + "")
    switch -9223372036854775808 {
    case 4294967295:
        v4 = v4
    case 2147483647:
        v36 = v4
    default:
        _ = len(v6)
    }
    switch 2147483647 {
    case -628:
        fmt.Println(v36, v16, v4)
    case 1:
        switch (v16 >> ((2) & 63)) {
        case 32768:
            switch v3 {
            case "hello":
                if (127 != 40) {
                    var v41 int16 = 42
                    var v42 int64 = (v4 ^ 0)
                    _ = v41
                    _ = v42
                }
            case "":
                v3 = "\n\t"
            default:
                var v43 uint64 = 0
                _ = v43
            }
        case 111:
            if false {
                for fi44 := 0; fi44 < 1; fi44++ {
                    if (false || false) {
                        for fi45 := 0; fi45 < 9; fi45++ {
                            switch (^(v16)) {
                            case 1:
                                v46 := GenS5{F0: 94, F1: 236, F2: fi44}
                                v36 = -1
                                var v47 float64 = (0.0 / (-2.25 * 1e10))
                                _ = len(v1)
                                if true {
                                    if (false || true) {
                                        var v48 uint8 = 222
                                        _ = v48
                                    } else {
                                        var v49 uint16 = 1
                                        v50 := GenS6{F0: int(v15), F1: v36}
                                        switch "a" {
                                        case "\n\t":
                                            var v51 int32 = 1
                                            _ = v51
                                        case "":
                                            if true {
                                                if ((673 >= 779) && false) {
                                                    v36 = (v36 ^ v36)
                                                } else {
                                                    v52 := GenS7{F0: 1, F1: 42}
                                                    if (("hello" + "a") <= (v3 + v3)) {
                                                        _ = len(v1)
                                                        v49 += 430
                                                    } else {
                                                        v53 := GenS5{F0: 165, F1: 0, F2: (-(v16))}
                                                        var v54 uint16 = v49
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    v55 := GenS6{F0: (fi44 % (((v16 | fi45)) | 1)), F1: -9223372036854775808}
                                                    switch 0 {
                                                    case 127:
                                                        var v56 float32 = (0.1 / 1.0)
                                                        fmt.Println(v49, v36, v4)
                                                        v57 := GenS8{F0: 4294967295}
                                                        _ = v56
                                                        _ = v57
                                                    case 4294967295:
                                                        var v58 string = "a"
                                                        _ = v58
                                                    }
                                                    var v59 int8 = (^(v15))
                                                    fmt.Println(v4)
                                                    if true {
                                                        if false {
                                                            delete(v1, v3)
                                                            var v60 uint16 = (^(v49))
                                                            fmt.Println(v3)
                                                            v4 += -128
                                                            v61 := GenS9{F0: 42, F1: (fi44 & v16), F2: genFn1(18446744073709551615, (v15 - genFn1(42, v59)))}
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        v62 := GenS10{F0: 321}
                                                        v47 -= (float64((3.14159 / 3.14159)) - v47)
                                                        _ = v62
                                                    }
                                                    _ = v52
                                                    _ = v55
                                                    _ = v59
                                                }
                                                var v63 int = (fi44 + -834)
                                                _ = len(v5)
                                                v64 := GenS1{F0: -649, F1: 0, F2: (v63 ^ v63)}
                                                _ = v63
                                                _ = v64
                                            }
                                        case "foo bar":
                                            if true {
                                                var v65 uint64 = 2
                                                _ = v65
                                            }
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                    if false {
                                        if ((true && ((v4 << ((232) & 63)) == v36)) && (!true)) {
                                            v47 -= 1e10
                                        } else {
                                            var v66 uint16 = 215
                                            _ = v66
                                        }
                                    }
                                    _ = len(v5)
                                } else {
                                    v47 = (float64(0.0) * (v47 + float64(1.5)))
                                    if ((v15 >> ((2) & 63)) < (v15 + (v15 | (v15 - genFn1(29, (v15 & v15)))))) {
                                        v5 = append(v5, fi45)
                                        fmt.Println(v4, v3)
                                    }
                                }
                                if (65535 < 255) {
                                    var v67 int32 = -2147483648
                                    var v68 uint32 = 2147483648
                                    fmt.Println(v16, v67, v67)
                                    _ = v67
                                    _ = v68
                                }
                                _ = v46
                                _ = v47
                            case 32768:
                                var v69 float64 = float64(-2.25)
                                _ = v69
                            }
                            if false {
                                v70 := v1[v3]
                                _ = v70
                            }
                            v15 += -113
                            _ = fi45
                        }
                        if (true && (((v15 % ((-25) | 1)) > v15) || (!(1000 < 42)))) {
                            v5 = append(v5, (fi44 | v16))
                            if true {
                                v71 := GenS6{F0: (v16 * fi44), F1: v4}
                                v72 := v2[((((v3 + ((("hello" + v3) + "foo bar") + (v3 + v3))) + "foo bar") + "a") + (((v3 + "\n\t") + v3) + (v3 + ("foo bar" + "\n\t"))))]
                                for fi73 := 0; fi73 < 19; fi73++ {
                                    _ = len(v2)
                                    _ = fi73
                                }
                                _ = v71
                                _ = v72
                            } else {
                                v16 = fi44
                            }
                            var v74 int32 = 552
                            _ = v74
                        }
                        v75 := GenS10{F0: 546}
                        _ = v75
                    } else {
                        v76 := GenS10{F0: 2}
                        v15 -= 0
                        _ = v76
                    }
                    var v77 int16 = 32767
                    _ = fi44
                    _ = v77
                }
                for fi78 := 0; fi78 < 19; fi78++ {
                    for fi79 := 0; fi79 < 9; fi79++ {
                        v80 := GenS6{F0: 576, F1: v4}
                        _ = fi79
                        _ = v80
                    }
                    for fi81 := 0; fi81 < 8; fi81++ {
                        v82 := GenS11{F0: 2}
                        _ = fi81
                        _ = v82
                    }
                    _ = fi78
                }
            }
        case 1:
            v36 += v4
        }
    default:
        switch (v3 + v3) {
        case "hello":
            switch ("\n\t" + "foo bar") {
            case "a":
                var v83 uint = 0
                _ = v83
            case "hello":
                v4 = v36
            case "":
                if (((v15 ^ v15) > (v15 % ((v15) | 1))) || true) {
                    for fi84 := 0; fi84 < 16; fi84++ {
                        v2[v3] = 0
                        fmt.Println(fi84, v15, v36)
                        _ = fi84
                    }
                    switch (v16 & (v16 + (v16 - int(v15)))) {
                    case -499:
                        var v85 string = (((v3 + v3) + ((("" + ((v3 + ("" + "")) + "")) + ("a" + v3)) + "")) + v3)
                        _ = v85
                    default:
                        if true {
                            v86 := GenS7{F0: 0, F1: -1}
                            switch (v16 * v16) {
                            case 1:
                                if ((v16 >= (v16 % ((v16) | 1))) && ("" >= (v3 + "\n\t"))) {
                                    v87 := GenS12{F0: 128, F1: (v16 - -2147483648)}
                                    _ = v87
                                }
                            default:
                                fmt.Println(v4, v16)
                            }
                            _ = v86
                        } else {
                            var v88 uint64 = 18446744073709551615
                            _ = v88
                        }
                    }
                }
            }
        case "foo bar":
            v15 -= v15
        case "":
            _ = len(v2)
        }
    }
    if false {
        v89 := GenS12{F0: 2, F1: (v16 >> ((258) & 63))}
        _ = v89
    }
    if true {
        var v90 uint32 = 2147483647
        _ = v90
    }
    var v91 int = -169
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v91
}
