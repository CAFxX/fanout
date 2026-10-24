package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint32; F1 int8; F2 string }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint16; F2 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64; F1 int }
func genFn0(p0 string, p1 string) int64 {
    v5 := GenS1{F0: 2, F1: 255, F2: 18446744073709551615}
    var v6 int64 = -235
    switch p0 {
    case "":
        switch "hello" {
        case "a":
            switch 2147483647 {
            case -9223372036854775808:
                if (((float32(0.1) + (0.0 / (1.5 + 1e10))) < 1e10) || (65535 >= 0)) {
                    if (!(((v6 - v6) < 816) || true)) {
                        if false {
                            if true {
                                var v7 int32 = 9
                                v8 := GenS1{F0: 42, F1: 42, F2: 2}
                                _ = v7
                                _ = v8
                            } else {
                                var v9 string = p1
                                for fi10 := 0; fi10 < 10; fi10++ {
                                    v6 = 1
                                    if true {
                                        switch (fi10 & fi10) {
                                        case 270:
                                            fmt.Println(p1)
                                        case 128:
                                            var v11 uint = 797
                                            _ = v11
                                        }
                                        p1 = "a"
                                        p0 = p1
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            if true {
                                                v12 := GenS1{F0: 2, F1: 1, F2: 2}
                                                var v13 uint16 = 638
                                                fmt.Println(p0)
                                                _ = v12
                                                _ = v13
                                            }
                                        }
                                        var v14 uint16 = 42
                                        var v15 string = p0
                                        _ = v14
                                        _ = v15
                                    }
                                    if false {
                                        v16 := GenS1{F0: 42, F1: 875, F2: 472}
                                        _ = v16
                                    }
                                    _ = fi10
                                }
                                _ = v9
                            }
                            if (false || false) {
                                var v17 uint64 = 283
                                var v18 float32 = -2.25
                                p1 = ("hello" + "")
                                _ = v17
                                _ = v18
                            } else {
                                fmt.Println(p1, v6)
                                var v19 float64 = (0.0 * 0.1)
                                _ = v19
                            }
                        }
                        p0 = ((p1 + p1) + ("foo bar" + p0))
                    } else {
                        v20 := []int{9223372036854775807, 128, 460}
                        var v21 uint = 42
                        for fi22 := 0; fi22 < 7; fi22++ {
                            for fi23 := 0; fi23 < 18; fi23++ {
                                if (!false) {
                                    fmt.Println(fi22, v21)
                                    var v24 float64 = 3.14159
                                    switch p1 {
                                    case "hello":
                                        var v25 uint32 = 2
                                        var v26 float64 = v24
                                        var v27 uint64 = 18446744073709551615
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v24
                                }
                                if (fi23 < (fi22 - -187)) {
                                    if true {
                                        v21 += (v21 & (v21 * 2))
                                        var v28 uint8 = 140
                                        v6 = v6
                                        _ = v28
                                    } else {
                                        fmt.Println(p0, fi22, v21)
                                        var v29 uint8 = 2
                                        _ = v29
                                    }
                                    _ = len(v20)
                                    p1 = p1
                                }
                                _ = fi23
                            }
                            v30 := GenS2{F0: 0, F1: v21}
                            switch (^(fi22)) {
                            case 128:
                                v31 := GenS3{F0: 0, F1: 0, F2: (p0 + (p0 + p0))}
                                _ = v31
                            case 42:
                                fmt.Println(fi22, v6, v6)
                            }
                            _ = fi22
                            _ = v30
                        }
                        _ = v20
                        _ = v21
                    }
                    var v32 uint64 = 42
                    _ = v32
                }
            case 2147483647:
                v33 := GenS4{F0: 128, F1: 255, F2: 0}
                _ = v33
            case 572:
                if true {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v36 := GenS2{F0: -1, F1: 226}
                            var v37 float32 = (float32(0.0) + 0.1)
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 11; fi38++ {
                            var v39 uint16 = 32767
                            _ = fi38
                            _ = v39
                        }
                        _ = fi34
                    }
                    v40 := []int{4294967295, 9223372036854775807}
                    _ = v40
                }
            }
        }
    case "\n\t":
        p1 = "a"
    case "a":
        var v41 float64 = 0.0
        _ = v41
    default:
        fmt.Println(p1, v6, v6)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (v6 % (((v6 >> ((65535) & 63))) | 1))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    if (!true) {
        fmt.Println(v2)
        if true {
            fmt.Println(v2)
            fmt.Println(v2)
            var v42 bool = true
            _ = v42
        } else {
            var v43 int64 = 42
            var v44 int64 = v43
            var v45 int32 = 1
            _ = v43
            _ = v44
            _ = v45
        }
    } else {
        if true {
            if true {
                _ = len(v4)
            } else {
                v46 := GenS2{F0: 1, F1: 18446744073709551615}
                v2 -= 1
                _ = v46
            }
        }
        v2 = v2
        var v47 int64 = -663
        _ = v47
    }
    v48 := GenS2{F0: 1000, F1: 42}
    v2 -= v2
    if (!(((1.5 + (3.14159 + 0.0)) - 0.0) != (0.1 / (0.1 / (3.14159 + float64(1e10)))))) {
        v2 -= v2
        v49 := GenS5{F0: -950}
        _ = v49
    }
    if false {
        v2 = (v2 * 2)
        fmt.Println(v2)
    } else {
        if (!true) {
            for fi50 := 0; fi50 < 7; fi50++ {
                v2 -= 4294967295
                _ = fi50
            }
        }
    }
    if false {
        switch (("foo bar" + "\n\t") + "foo bar") {
        case "hello":
            v2 += 18446744073709551615
        }
        if (128 == 42) {
            if true {
                v51 := GenS1{F0: 0, F1: 65535, F2: (v2 + (v2 % (((^(v2))) | 1)))}
                for fi52 := 0; fi52 < 14; fi52++ {
                    delete(v3, uint32(v2))
                    _ = fi52
                }
                _ = v51
            }
            v2 = (v2 * v2)
        }
        v2 = v2
    }
    for fi53 := 0; fi53 < 1; fi53++ {
        _ = len(v1)
        fmt.Println(v2, fi53)
        for fi54 := 0; fi54 < 3; fi54++ {
            switch ("hello" + "hello") {
            case "hello":
                var v55 float32 = 1e10
                _ = v55
            case "foo bar":
                var v56 int64 = 0
                v57 := GenS1{F0: 351, F1: 65535, F2: (v2 & v2)}
                _ = v56
                _ = v57
            case "a":
                if len(v4) > 0 {
                    v4[0] = (((("\n\t" + ("\n\t" + "")) + ("" + "")) + "a") + ("foo bar" + "hello"))
                }
            }
            v58 := GenS4{F0: 211, F1: 42, F2: 0}
            if true {
                switch "\n\t" {
                case "hello":
                    fmt.Println(fi54)
                case "":
                    v59 := GenS2{F0: 42, F1: 0}
                    _ = v59
                case "foo bar":
                    fmt.Println(fi54, v2, fi54)
                default:
                    if false {
                        var v60 int = fi54
                        var v61 int8 = -1
                        switch ("" + "") {
                        case "a":
                            switch (("\n\t" + "foo bar") + ("a" + "\n\t")) {
                            case "":
                                switch "a" {
                                case "":
                                    if false {
                                        var v62 bool = ("" >= "")
                                        if (4294967295 == 4294967295) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v63 uint32 = 445
                                            if len(v4) > 0 {
                                                v4[0] = "\n\t"
                                            }
                                            _ = v63
                                        }
                                        v60 = fi53
                                        _ = v62
                                    }
                                case "hello":
                                    v64 := GenS5{F0: 1000}
                                    _ = v64
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                }
                            }
                        case "foo bar":
                            v65 := GenS2{F0: 1000, F1: 1}
                            _ = v65
                        default:
                            _ = len(v3)
                        }
                        _ = v60
                        _ = v61
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 = (v2 - 723)
            } else {
                var v66 uint32 = 1
                v66 = v66
                if (31 == 241) {
                    v67 := GenS2{F0: 1, F1: 901}
                    if false {
                        if false {
                            _ = len(v3)
                            if (!((int(v2) < (fi53 << ((1) & 63))) || true)) {
                                v68 := GenS6{F0: 4294967295, F1: 1, F2: 2}
                                var v69 float64 = 1.5
                                _ = v68
                                _ = v69
                            }
                        }
                    } else {
                        var v70 bool = true
                        if (-1 <= -1) {
                            var v71 float64 = float64(0.0)
                            v71 -= v71
                            if v70 {
                                var v72 float32 = 1.5
                                if ((v2 >> ((1) & 63)) == v2) {
                                    v2 -= v2
                                    v73 := GenS1{F0: 1, F1: 255, F2: 621}
                                    _ = v73
                                }
                                _ = v72
                            } else {
                                v71 = (0.0 * float64(0.0))
                                switch (fi54 & (^(fi54))) {
                                case -938:
                                    v2 += v2
                                }
                                var v74 string = "foo bar"
                                _ = v74
                            }
                            _ = v71
                        }
                        _ = v70
                    }
                    _ = v67
                } else {
                    var v75 int8 = 1
                    if false {
                        v66 += (^(v66))
                        fmt.Println(v2, v2, v66)
                    } else {
                        if false {
                            v75 = v75
                            var v76 uint32 = (v66 - v66)
                            if true {
                                var v77 uint64 = v2
                                if ((^(fi53)) > fi54) {
                                    v78 := GenS7{F0: 406}
                                    _ = v78
                                } else {
                                    v1["hello"] = (v66 + v76)
                                }
                                fmt.Println(v76, v76)
                                v76 += v66
                                switch 32768 {
                                case -9223372036854775808:
                                    _ = len(v3)
                                case 9223372036854775807:
                                    v75 = 100
                                }
                                _ = v77
                            } else {
                                v79 := GenS8{F0: false}
                                _ = len(v3)
                                _ = len(v3)
                                v80 := GenS3{F0: 226, F1: (v75 * (v75 | (v75 << ((1) & 63)))), F2: "foo bar"}
                                var v81 float32 = -2.25
                                var v82 int = (fi54 % ((fi54) | 1))
                                _ = v79
                                _ = v80
                                _ = v81
                                _ = v82
                            }
                            _ = v76
                        } else {
                            v83 := GenS1{F0: 0, F1: 1, F2: (v2 & (^(v2)))}
                            _ = v83
                        }
                    }
                    _ = v75
                }
                _ = v66
            }
            _ = fi54
            _ = v58
        }
        _ = fi53
    }
    v84 := v3[0]
    if (768 == v2) {
        v85 := GenS9{F0: 1}
        _ = v85
    } else {
        v86 := GenS10{F0: v2, F1: 127}
        for fi87 := 0; fi87 < 6; fi87++ {
            for fi88 := 0; fi88 < 17; fi88++ {
                fmt.Println(v2, v84, v2)
                _ = fi88
            }
            _ = fi87
        }
        switch "hello" {
        case "a":
            var v89 int8 = -1
            _ = v89
        case "hello":
            fmt.Println(v2)
        }
        _ = v86
    }
    fmt.Println(v84)
    var v90 float64 = 0.0
    var v91 int = -9223372036854775808
    switch "foo bar" {
    case "foo bar":
        v2 = v2
    case "\n\t":
        if ((1 <= -509) && false) {
            fmt.Println(v90, v91, v90)
            fmt.Println(v90)
        }
    case "a":
        v91 -= v91
    default:
        var v92 string = "hello"
        _ = v92
    }
    fmt.Println(v90, v91)
    v93 := GenS2{F0: 0, F1: 255}
    v2 = (v2 ^ (v2 >> ((uint(v2)) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v84:", v84)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v84
    _ = v90
    _ = v91
    _ = v93
}
