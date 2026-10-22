package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS7 struct { F0 int8; F1 int8; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int }

func main() {
    var v1 uint64 = 1
    var v2 float32 = (-2.25 - (0.1 * 0.1))
    var v3 int16 = int16(v1)
    var v4 float32 = 1e10
    v5 := make(map[string]uint32)
    v2 -= v2
    switch ((("hello" + "hello") + (("\n\t" + "") + "hello")) + "") {
    case "foo bar":
        v5[""] = 0
    case "":
        _ = len(v5)
    }
    fmt.Println(v2, v3, v2)
    var v6 float32 = v2
    v4 = 1e10
    fmt.Println(v2)
    var v7 float32 = float32(1.5)
    v8 := []int{}
    if false {
        var v9 int64 = 0
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: v9}
            v12 := v5["\n\t"]
            _ = fi10
            _ = v11
            _ = v12
        }
        switch ("hello" + ("\n\t" + "a")) {
        case "a":
            if ((3.14159 + 0.0) != -2.25) {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 bool = false
                    _ = fi13
                    _ = v14
                }
                switch 128 {
                case -2147483648:
                    _ = len(v5)
                }
                v7 = float32(float64(float32((3.14159 * (0.1 - 0.1)))))
            } else {
                v7 = (1e10 * (1e10 * ((-2.25 * float32(0.0)) - 0.0)))
                v15 := GenS1{F0: v9}
                for fi16 := 0; fi16 < 19; fi16++ {
                    switch (fi16 + 32767) {
                    case 9223372036854775807:
                        v2 = -2.25
                    case 671:
                        fmt.Println(v9, v9, fi16)
                    }
                    for fi17 := 0; fi17 < 9; fi17++ {
                        fmt.Println(fi17, v3)
                        var v18 uint16 = 32767
                        _ = fi17
                        _ = v18
                    }
                    if (-2.25 >= float64(v4)) {
                        v6 = (v6 - v6)
                        delete(v5, "foo bar")
                        var v19 float32 = ((v2 - 0.1) * v7)
                        _ = v19
                    } else {
                        v4 = (v4 / (((((0.1 / v4) - (float32(1.5) + 0.0)) - ((-2.25 + (1.5 * v6)) + v4)) + v7) + v7))
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v21 := GenS1{F0: v9}
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v22 := GenS2{F0: (v3 / ((437) | 1)), F1: "hello"}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi16
                }
                _ = v15
            }
        case "foo bar":
            v9 = (v9 * v9)
        default:
            fmt.Println(v7, v9, v2)
        }
        _ = v9
    }
    if true {
        fmt.Println(v3)
        if len(v8) > 0 {
            v8[0] = -9223372036854775808
        }
        v1 -= v1
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        switch ("hello" + "hello") {
        case "hello":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = fi23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        if (v3 <= 32767) {
            switch "\n\t" {
            case "":
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 uint32 = 2147483648
                    fmt.Println(v4)
                    _ = fi25
                    _ = v26
                }
            case "\n\t":
                if (42 <= 73) {
                    var v27 float32 = v6
                    _ = v27
                } else {
                    v1 = v1
                    v4 -= v2
                    var v28 string = "a"
                    _ = v28
                }
            case "foo bar":
                switch fi24 {
                case 128:
                    v29 := GenS1{F0: 434}
                    _ = v29
                case -456:
                    switch -713 {
                    case 9223372036854775807:
                        v5["a"] = 4294967295
                    case -2147483648:
                        v2 = (((3.14159 / (1e10 * 3.14159)) - ((float32(3.14159) - v7) + float32(0.1))) * (v4 + -2.25))
                    case -34:
                        if false {
                            v30 := GenS1{F0: 1}
                            _ = v30
                        }
                    default:
                        if (true && true) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            if (352 <= 749) {
                                switch fi24 {
                                case -195:
                                    v31 := v5["\n\t"]
                                    _ = v31
                                case 1:
                                    if (0 >= 179) {
                                        v32 := GenS1{F0: 42}
                                        _ = v32
                                    }
                                default:
                                    var v33 uint32 = 2
                                    _ = v33
                                }
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v5)
                                default:
                                    for fi34 := 0; fi34 < 17; fi34++ {
                                        if (v1 < v1) {
                                            v35 := v5["a"]
                                            v36 := GenS3{F0: (v3 ^ v3)}
                                            v37 := GenS4{F0: 70}
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        } else {
                                            v1 += 18446744073709551615
                                            var v38 bool = true
                                            var v39 int = (fi24 | fi34)
                                            var v40 int = (-(fi34))
                                            if (!((v3 * (v3 >> ((255) & 63))) > -1)) {
                                                var v41 int = fi34
                                                _ = v41
                                            }
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                        _ = len(v5)
                                        _ = fi34
                                    }
                                }
                                v4 -= -2.25
                                var v42 float64 = ((0.0 * ((0.1 + 0.1) - 0.0)) * 1e10)
                                _ = v42
                            }
                            v4 += 1.5
                        } else {
                            if false {
                                v6 = v2
                            }
                        }
                    }
                default:
                    var v43 uint8 = 2
                    _ = v43
                }
            default:
                if true {
                    fmt.Println(fi24, v1)
                }
            }
            switch -2147483648 {
            case 32767:
                for fi44 := 0; fi44 < 3; fi44++ {
                    var v45 float64 = 0.0
                    _ = fi44
                    _ = v45
                }
            case -2147483648:
                for fi46 := 0; fi46 < 10; fi46++ {
                    v7 = ((0.1 + (v4 + ((1e10 + (float32(((0.1 + 3.14159) - 1.5)) * ((((float32((float64(float32(0.0)) + -2.25)) - (v4 + v6)) - v2) / v7) / ((v7 / 1.5) + v2)))) / float32(1e10)))) - 0.1)
                    if false {
                        fmt.Println(v7, fi24, v3)
                    } else {
                        var v47 uint = 2
                        _ = len(v5)
                        fmt.Println(fi46, v47, v47)
                        var v48 uint16 = 32767
                        _ = v47
                        _ = v48
                    }
                    switch (fi24 >> ((0) & 63)) {
                    case -2147483648:
                        v8 = append(v8, fi46)
                    case -697:
                        v3 -= v3
                    case -551:
                        if false {
                            v2 = 0.1
                        } else {
                            var v49 uint64 = 92
                            if (32767 == 0) {
                                var v50 uint = 42
                                _ = v50
                            }
                            if true {
                                if (false && true) {
                                    fmt.Println(v49)
                                } else {
                                    delete(v5, "a")
                                }
                                v3 = v3
                                var v51 int32 = 428
                                _ = v51
                            }
                            _ = v49
                        }
                    default:
                        switch (fi46 >> ((2) & 63)) {
                        case 4294967295:
                            var v52 float32 = float32(1.5)
                            _ = v52
                        case -144:
                            v53 := GenS5{F0: (!true), F1: v3}
                            _ = v53
                        case 32767:
                            if ((((0 != -9223372036854775808) && false) || false) || true) {
                                if (!true) {
                                    if (((((float32(-2.25) / v7) + v4) + float32((0.0 - 1e10))) - float32(1.5)) == ((1.5 + v7) + v6)) {
                                        if false {
                                            switch (-(fi24)) {
                                            case -2147483648:
                                                v1 -= v1
                                            case -711:
                                                if ((((1000 >= 32767) && (!true)) && false) && (!(-2.25 > 0.1))) {
                                                    v3 += v3
                                                } else {
                                                    var v54 float32 = float32((1e10 / 1.0))
                                                    var v55 int32 = -1
                                                    _ = v54
                                                    _ = v55
                                                }
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    switch "" {
                                                    case "hello":
                                                        var v56 uint32 = 2147483647
                                                        _ = v56
                                                    case "a":
                                                        var v57 uint8 = 127
                                                        _ = v57
                                                    case "":
                                                        if false {
                                                            v58 := GenS6{F0: 23, F1: ((fi24 >> ((42) & 63)) <= (fi46 % ((-2147483648) | 1))), F2: v1}
                                                            var v59 uint8 = 2
                                                            var v60 bool = (41 > -5)
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            if true {
                                                                var v61 int = 128
                                                                switch (fi24 + (^(v61))) {
                                                                case 462:
                                                                    v61 = (fi46 + (fi24 | -519))
                                                                }
                                                                _ = len(v5)
                                                                var v62 int64 = 0
                                                                _ = v61
                                                                _ = v62
                                                            }
                                                            switch (fi46 + -585) {
                                                            case 32768:
                                                                if true {
                                                                    v63 := GenS1{F0: -9223372036854775808}
                                                                    var v64 uint = 71
                                                                    _ = v63
                                                                    _ = v64
                                                                } else {
                                                                    if (!(((false && ((255 >= 128) || false)) || ((^(v1)) == (v1 & v1))) && (true && true))) {
                                                                        var v65 string = ("a" + "hello")
                                                                        _ = v65
                                                                    } else {
                                                                        if ((!false) && true) {
                                                                            if false {
                                                                                v6 += float32(-2.25)
                                                                                fmt.Println(fi24, v1)
                                                                            } else {
                                                                                if false {
                                                                                    if len(v8) > 0 {
                                                                                        v8[0] = fi46
                                                                                    }
                                                                                    var v66 int64 = 1
                                                                                    v67 := GenS2{F0: (v3 % (((v3 + v3)) | 1)), F1: ("a" + (("a" + "hello") + ""))}
                                                                                    if ("\n\t" >= "\n\t") {
                                                                                        v66 -= (v66 + v66)
                                                                                        v3 = (-(v3))
                                                                                        v68 := GenS7{F0: -95, F1: 42, F2: 32767}
                                                                                        _ = v68
                                                                                    } else {
                                                                                        v6 = 3.14159
                                                                                    }
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                }
                                                                                switch ("" + "\n\t") {
                                                                                case "hello":
                                                                                    v4 -= v2
                                                                                default:
                                                                                    v69 := GenS1{F0: -9223372036854775808}
                                                                                    _ = v69
                                                                                }
                                                                                var v70 int16 = (v3 ^ -459)
                                                                                _ = v70
                                                                            }
                                                                            fmt.Println(v3)
                                                                        } else {
                                                                            var v71 float64 = -2.25
                                                                            v72 := GenS2{F0: v3, F1: ("hello" + "")}
                                                                            v7 = (-2.25 + v2)
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        v2 = (v6 * v7)
                                                                        v7 = v7
                                                                    }
                                                                }
                                                            }
                                                            v73 := GenS8{F0: 2, F1: 65535, F2: fi46}
                                                            _ = v73
                                                        }
                                                    }
                                                }
                                            }
                                            var v74 int8 = -1
                                            var v75 uint16 = 2
                                            v76 := GenS4{F0: 879}
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                        v77 := v5[(("" + "") + ("hello" + "a"))]
                                        var v78 string = (("a" + ("a" + "")) + "foo bar")
                                        _ = v77
                                        _ = v78
                                    }
                                    fmt.Println(v1, v3)
                                }
                            }
                        }
                    }
                    _ = fi46
                }
            case 127:
                _ = len(v5)
            default:
                if (v4 < -2.25) {
                    v4 = v2
                    v1 -= 1
                }
            }
            for fi79 := 0; fi79 < 3; fi79++ {
                v2 -= v2
                fmt.Println(v1)
                fmt.Println(v7)
                v5["a"] = 289
                _ = fi79
            }
        } else {
            fmt.Println(v2)
            if (true || false) {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch fi24 {
                case -2147483648:
                    var v80 float64 = (3.14159 + 1.5)
                    _ = v80
                default:
                    var v81 int32 = 2147483647
                    _ = v81
                }
            } else {
                if (4294967295 <= 156) {
                    v82 := GenS6{F0: 41, F1: false, F2: v1}
                    _ = v82
                } else {
                    if (((("hello" + (("hello" + "a") + "hello")) + ((("\n\t" + "foo bar") + ("foo bar" + ("\n\t" + ("\n\t" + "a")))) + "\n\t")) + "") < (("foo bar" + "a") + ("hello" + ""))) {
                        v1 = 4294967295
                        for fi83 := 0; fi83 < 8; fi83++ {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            fmt.Println(v7, fi83)
                            _ = fi83
                        }
                        var v84 int8 = 100
                        _ = v84
                    } else {
                        v3 = (v3 ^ v3)
                        var v85 uint32 = 42
                        var v86 bool = (false && (2 > 1))
                        _ = v85
                        _ = v86
                    }
                    _ = len(v5)
                    if (-499 <= -1) {
                        v87 := GenS1{F0: 42}
                        if false {
                            v88 := GenS1{F0: 0}
                            v3 += (-(v3))
                            v1 += 18446744073709551615
                            _ = v88
                        } else {
                            v89 := v5["foo bar"]
                            delete(v5, "a")
                            _ = v89
                        }
                        _ = len(v8)
                        _ = v87
                    } else {
                        var v90 int16 = 15
                        _ = v90
                    }
                }
                if (42 == 127) {
                    var v91 float32 = float32(1e10)
                    v6 = 3.14159
                    _ = v91
                } else {
                    v1 = v1
                    v7 -= 1e10
                    for fi92 := 0; fi92 < 16; fi92++ {
                        v8 = append(v8, 9223372036854775807)
                        fmt.Println(v2, v2, fi92)
                        _ = fi92
                    }
                }
                _ = len(v5)
            }
            v3 += (v3 | (v3 / ((v3) | 1)))
        }
        _ = fi24
    }
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
