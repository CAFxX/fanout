package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 int16; F2 int16 }
type GenS8 struct { F0 bool }
func genFn0(p0 bool, p1 string, p2 uint64) string {
    p1 = p1
    if true {
        p2 = 104
        var v6 uint = 1
        _ = v6
    } else {
        var v7 string = ("hello" + p1)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := []uint32{739, 2, 2, 4294967295}
    var v2 uint8 = 33
    var v3 int64 = 42
    var v4 uint16 = 255
    v5 := make(map[string]uint32)
    delete(v5, "")
    _ = len(v1)
    v3 = v3
    var v8 uint32 = 2147483648
    switch 139 {
    case 128:
        if ((false && true) || true) {
            switch ("foo bar" + "foo bar") {
            case "a":
                if (!true) {
                    v1 = append(v1, (v8 + 2))
                }
            case "foo bar":
                if false {
                    v9 := GenS1{F0: 0, F1: v8, F2: v2}
                    var v10 uint8 = v2
                    if (false || (-2.25 > (((1e10 - 1.5) - 0.0) + (1e10 * 1e10)))) {
                        fmt.Println(v8)
                        v11 := GenS1{F0: -128, F1: v8, F2: 1}
                        _ = v11
                    }
                    _ = v10
                    _ = v9
                }
            case "\n\t":
                for fi12 := 0; fi12 < 15; fi12++ {
                    if false {
                        var v13 uint64 = 0
                        v13 = (v13 % ((v13) | 1))
                        _ = v13
                    } else {
                        v4 = v4
                    }
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 int64 = v3
                        if true {
                            v4 = 1
                        }
                        v15 = (v15 - v3)
                        v16 := v5[("a" + "")]
                        fmt.Println(v8, v8, v3)
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = fi12
                }
            default:
                var v17 float32 = 0.0
                _ = v17
            }
            if (float64((0.1 / float32(-2.25))) == 0.1) {
                if true {
                    v3 = (v3 & v3)
                }
                switch ("hello" + "a") {
                case "\n\t":
                    v18 := GenS2{F0: 42, F1: 42}
                    _ = v18
                case "foo bar":
                    v19 := GenS2{F0: v3, F1: 42}
                    _ = v19
                case "a":
                    v20 := GenS1{F0: 107, F1: 0, F2: 5}
                    _ = v20
                }
            }
        }
    }
    var v21 bool = false
    var v22 uint64 = 0
    switch ("\n\t" + ("hello" + "\n\t")) {
    case "hello":
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 uint = 42
            switch fi23 {
            case 32768:
                if false {
                    _ = len(v1)
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v26 := GenS1{F0: 42, F1: (v8 | 700), F2: (v2 & (v2 << (((v24 | v24)) & 63)))}
                        if (!((1e10 / (3.14159 + 0.0)) >= (0.0 + -2.25))) {
                            v27 := GenS2{F0: -297, F1: 42}
                            var v28 uint64 = 42
                            var v29 float32 = -2.25
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            v24 = (v24 ^ (v24 | 206))
                            v30 := GenS1{F0: -128, F1: (v8 / (((v8 - 421)) | 1)), F2: (v2 - (^(v2)))}
                            var v31 uint8 = (v2 + v2)
                            _ = v30
                            _ = v31
                        }
                        _ = fi25
                        _ = v26
                    }
                } else {
                    v21 = v21
                }
            case -534:
                v32 := GenS3{F0: v24}
                _ = v32
            case -2147483648:
                var v33 float64 = 3.14159
                _ = v33
            }
            fmt.Println(v21, fi23, v8)
            _ = fi23
            _ = v24
        }
    case "\n\t":
        v4 -= (v4 | (v4 / ((v4) | 1)))
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        v35 := GenS1{F0: 100, F1: v8, F2: 195}
        _ = fi34
        _ = v35
    }
    if v21 {
        if false {
            if v21 {
                v2 = v2
                v2 -= v2
            } else {
                if len(v1) > 0 {
                    v1[0] = uint32(v4)
                }
            }
            for fi36 := 0; fi36 < 5; fi36++ {
                fmt.Println(fi36)
                _ = fi36
            }
            v1 = append(v1, v8)
        }
    }
    v22 += 18446744073709551615
    v3 = -803
    for fi37 := 0; fi37 < 15; fi37++ {
        for fi38 := 0; fi38 < 4; fi38++ {
            var v39 string = "hello"
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            if len(v1) > 0 {
                v1[0] = (v8 | v8)
            }
            var v41 float32 = -2.25
            if v21 {
                v3 = v3
                delete(v5, "")
            }
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 1; fi42++ {
            fmt.Println(v22)
            v2 = (v2 & (v2 >> ((2) & 63)))
            v3 = v3
            _ = fi42
        }
        _ = fi37
    }
    var v43 float32 = -2.25
    for fi44 := 0; fi44 < 6; fi44++ {
        fmt.Println(v2, v2, v43)
        v45 := GenS4{F0: ("a" + ("hello" + "\n\t")), F1: v3, F2: v22}
        if ("hello" == "\n\t") {
            var v46 uint = 4294967295
            fmt.Println(v4, v46)
            _ = v46
        } else {
            if false {
                if ((v21 || v21) || ((v4 == 0) || (0 != (v3 | v3)))) {
                    if false {
                        v4 = v4
                    } else {
                        v2 += v2
                        switch ("a" + "a") {
                        case "hello":
                            if (true && ("a" == "\n\t")) {
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    var v48 string = "a"
                                    if v21 {
                                        var v49 uint16 = uint16(v8)
                                        v50 := v5[v48]
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = fi47
                                    _ = v48
                                }
                            }
                        }
                    }
                    fmt.Println(v4, v8, v3)
                    _ = len(v1)
                } else {
                    if true {
                        if (!(!true)) {
                            var v51 int32 = 0
                            var v52 uint32 = 4294967295
                            _ = v51
                            _ = v52
                        }
                    }
                    delete(v5, ("foo bar" + ("hello" + ("foo bar" + "a"))))
                }
                fmt.Println(v3, v21, fi44)
                switch (("" + "a") + "\n\t") {
                case "hello":
                    var v53 uint32 = v8
                    _ = v53
                case "\n\t":
                    if (v21 && (183 == 2147483647)) {
                        v8 += (^(v8))
                        fmt.Println(fi44, v4, v22)
                        if ((v8 / ((265) | 1)) != v8) {
                            switch "hello" {
                            case "hello":
                                if ((fi44 * fi44) > fi44) {
                                    if v21 {
                                        if ((v2 ^ (v2 << ((4294967295) & 63))) >= 183) {
                                            switch 1 {
                                            case 1:
                                                if (!true) {
                                                    v54 := v5[("a" + ("a" + "foo bar"))]
                                                    switch "\n\t" {
                                                    case "":
                                                        var v55 float32 = v43
                                                        _ = v55
                                                    case "a":
                                                        var v56 float32 = v43
                                                        _ = v56
                                                    case "foo bar":
                                                        v57 := v5[("\n\t" + ("" + ("\n\t" + ("\n\t" + ("hello" + ((("a" + ("hello" + "\n\t")) + (("hello" + "\n\t") + "foo bar")) + "foo bar"))))))]
                                                        _ = v57
                                                    default:
                                                        var v58 uint16 = 2
                                                        _ = v58
                                                    }
                                                    _ = v54
                                                } else {
                                                    v4 = v4
                                                    for fi59 := 0; fi59 < 2; fi59++ {
                                                        if (v21 && (v3 != 42)) {
                                                            v60 := GenS5{F0: 0}
                                                            if v21 {
                                                                v61 := GenS2{F0: 1, F1: 0}
                                                                _ = v61
                                                            } else {
                                                                fmt.Println(fi59, v22, fi44)
                                                            }
                                                            _ = v60
                                                        }
                                                        _ = fi59
                                                    }
                                                    v4 += v4
                                                }
                                            }
                                            var v62 string = "a"
                                            _ = v62
                                        } else {
                                            v4 += (^(v4))
                                            var v63 string = (("\n\t" + "a") + ("foo bar" + ("\n\t" + "")))
                                            v1 = append(v1, (v8 & v8))
                                            _ = v63
                                        }
                                        for fi64 := 0; fi64 < 1; fi64++ {
                                            switch fi44 {
                                            case -406:
                                                fmt.Println(v8, v4, fi44)
                                            case -9223372036854775808:
                                                switch (fi44 + (fi44 * int(v22))) {
                                                case -155:
                                                    v65 := GenS6{F0: v3}
                                                    _ = v65
                                                case 128:
                                                    delete(v5, ("hello" + "foo bar"))
                                                default:
                                                    fmt.Println(v22, v8, v8)
                                                }
                                            default:
                                                if ((v4 << ((uint(v4)) & 63)) == v4) {
                                                    var v66 uint64 = (v22 + 164)
                                                    _ = v66
                                                }
                                            }
                                            _ = fi64
                                        }
                                        fmt.Println(v8, v3)
                                    } else {
                                        switch (fi44 % (((fi44 >> ((1) & 63))) | 1)) {
                                        case -889:
                                            v67 := GenS4{F0: "a", F1: v3, F2: (v22 % ((42) | 1))}
                                            _ = v67
                                        }
                                        fmt.Println(v4)
                                    }
                                    for fi68 := 0; fi68 < 7; fi68++ {
                                        v22 += (v22 << ((42) & 63))
                                        _ = fi68
                                    }
                                } else {
                                    switch fi44 {
                                    case 128:
                                        v8 = (v8 << ((18446744073709551615) & 63))
                                    default:
                                        var v69 uint16 = (v4 + v4)
                                        _ = v69
                                    }
                                    var v70 float32 = (v43 / v43)
                                    _ = v70
                                }
                            }
                            _ = len(v1)
                        }
                    } else {
                        v71 := GenS7{F0: "a", F1: 838, F2: 1000}
                        fmt.Println(v3, v21, v43)
                        _ = v71
                    }
                default:
                    if ((v21 || v21) || ((v2 >> ((1) & 63)) >= v2)) {
                        for fi72 := 0; fi72 < 10; fi72++ {
                            v73 := GenS6{F0: (v3 & (v3 / ((539) | 1)))}
                            _ = fi72
                            _ = v73
                        }
                        v8 -= (v8 + v8)
                    }
                }
            } else {
                if ((v4 >= v4) || (v21 || (float32(float64(v43)) < v43))) {
                    if v21 {
                        if v21 {
                            var v74 float32 = -2.25
                            _ = v74
                        } else {
                            v4 -= (v4 | v4)
                            _ = len(v1)
                        }
                        v75 := GenS1{F0: 0, F1: v8, F2: (^(v2))}
                        v2 -= (v2 + v2)
                        _ = v75
                    }
                    fmt.Println(v21, v21)
                    for fi76 := 0; fi76 < 20; fi76++ {
                        v77 := GenS3{F0: uint(v4)}
                        v78 := GenS6{F0: v3}
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                }
                fmt.Println(v21, v43)
            }
            _ = len(v1)
        }
        _ = fi44
        _ = v45
    }
    if true {
        for fi79 := 0; fi79 < 15; fi79++ {
            switch ("hello" + "foo bar") {
            case "hello":
                v22 = v22
            case "foo bar":
                var v80 int = 277
                _ = v80
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 953
                }
            default:
                v81 := GenS2{F0: -9223372036854775808, F1: 0}
                _ = v81
            }
            var v82 float32 = 0.0
            _ = fi79
            _ = v82
        }
        v22 = 4294967295
    }
    var v83 string = ("a" + (("a" + ("foo bar" + "\n\t")) + ("hello" + "foo bar")))
    if (!((v8 - v8) > 961)) {
        for fi84 := 0; fi84 < 12; fi84++ {
            switch (fi84 & fi84) {
            case 128:
                var v85 uint16 = 255
                _ = v85
            }
            v86 := GenS5{F0: 1}
            fmt.Println(v21)
            _ = fi84
            _ = v86
        }
        v87 := GenS8{F0: (v21 || (1 < -32768))}
        var v88 uint32 = v8
        _ = v87
        _ = v88
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v8
    _ = v83
}
