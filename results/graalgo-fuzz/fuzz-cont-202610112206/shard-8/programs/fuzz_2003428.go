package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
type GenS2 struct { F0 int64; F1 bool; F2 uint32 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int; F1 uint; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 uint; F2 uint16 }
func genFn0(p0 bool) int64 {
    p0 = true
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 float32 = 1.5
        p0 = p0
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(p0, v5, fi6)
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    switch ("" + "foo bar") {
    case "hello":
        p0 = true
    }
    _ = p0
    return 843
}

func main() {
    v1 := []string{}
    var v2 uint32 = 175
    var v3 float32 = 3.14159
    v3 += (1e10 + v3)
    var v7 bool = false
    var v8 float32 = v3
    var v9 float32 = (v8 / v8)
    v8 += v9
    var v10 float32 = (v9 * v8)
    if false {
        var v11 int16 = 1000
        for fi12 := 0; fi12 < 16; fi12++ {
            v2 += (v2 * v2)
            _ = fi12
        }
        _ = v11
    } else {
        if v7 {
            v7 = ((42 == 1000) && v7)
            v9 += float32(0.0)
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 float32 = 3.14159
                    _ = fi14
                    _ = v15
                }
                switch "a" {
                case "hello":
                    for fi16 := 0; fi16 < 15; fi16++ {
                        var v17 uint = 1
                        _ = fi16
                        _ = v17
                    }
                case "":
                    switch ("hello" + "foo bar") {
                    case "":
                        v1 = append(v1, "\n\t")
                    case "foo bar":
                        var v18 uint32 = (v2 % ((v2) | 1))
                        _ = v18
                    case "a":
                        v19 := GenS1{F0: 128, F1: "", F2: -915}
                        _ = v19
                    default:
                        for fi20 := 0; fi20 < 16; fi20++ {
                            var v21 string = "hello"
                            _ = fi20
                            _ = v21
                        }
                    }
                default:
                    switch (fi13 * (fi13 / ((fi13) | 1))) {
                    case -1:
                        if (v2 != 811) {
                            fmt.Println(v7, v8, v9)
                        }
                    case 32768:
                        v22 := GenS1{F0: (fi13 * fi13), F1: "a", F2: -203}
                        _ = v22
                    case 32767:
                        switch (^(fi13)) {
                        case -641:
                            switch fi13 {
                            case 127:
                                fmt.Println(v8, v8)
                            case 843:
                                for fi23 := 0; fi23 < 15; fi23++ {
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                    _ = fi23
                                }
                            default:
                                for fi24 := 0; fi24 < 19; fi24++ {
                                    fmt.Println(v3, fi13)
                                    _ = fi24
                                }
                            }
                        case 0:
                            v7 = v7
                        case 128:
                            _ = len(v1)
                        }
                    }
                }
                v25 := GenS1{F0: fi13, F1: "foo bar", F2: -32768}
                _ = fi13
                _ = v25
            }
        }
        var v26 bool = (1 > 42)
        _ = v26
    }
    v10 = 0.1
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint = 18446744073709551615
        fmt.Println(v28)
        v29 := GenS2{F0: genFn0(v7), F1: (v7 || (!false)), F2: v2}
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        if ((528 < 42) && v7) {
            var v31 int16 = 346
            _ = v31
        } else {
            if v7 {
                for fi32 := 0; fi32 < 8; fi32++ {
                    if (1 > 42) {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                    fmt.Println(v7, v10)
                    _ = fi32
                }
                for fi33 := 0; fi33 < 19; fi33++ {
                    if true {
                        if v7 {
                            fmt.Println(v9, v3, v3)
                        }
                        if (255 == 0) {
                            switch ((("a" + ("\n\t" + "foo bar")) + ("" + (("" + ("a" + ("a" + "\n\t"))) + "foo bar"))) + "a") {
                            case "foo bar":
                                if (v9 == v9) {
                                    v10 -= v8
                                }
                            case "":
                                v3 += (0.1 - (3.14159 - -2.25))
                            default:
                                if true {
                                    _ = len(v1)
                                    if (65535 != 255) {
                                        v10 = 0.1
                                        v7 = v7
                                    }
                                    v7 = ((1000 >= -544) && v7)
                                    if v7 {
                                        if (!v7) {
                                            v10 -= v3
                                        }
                                        if false {
                                            var v34 int8 = 127
                                            if true {
                                                var v35 bool = v7
                                                if (fi30 <= fi30) {
                                                    var v36 int = (fi33 / ((fi30) | 1))
                                                    _ = v36
                                                } else {
                                                    switch (-(fi33)) {
                                                    case 2147483647:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if v7 {
                                                                var v37 float64 = 0.0
                                                                switch (fi30 >> ((18446744073709551615) & 63)) {
                                                                case 42:
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v1 = append(v1, ("a" + "foo bar"))
                                                                case 0:
                                                                    var v38 int64 = -121
                                                                    _ = v38
                                                                default:
                                                                    var v39 string = "\n\t"
                                                                    v3 += v9
                                                                    v3 -= v10
                                                                    _ = v39
                                                                }
                                                                var v40 int16 = 32767
                                                                _ = v37
                                                                _ = v40
                                                            } else {
                                                                fmt.Println(fi30, v35)
                                                                v41 := GenS3{F0: v35, F1: "hello"}
                                                                fmt.Println(fi30, v10)
                                                                _ = v41
                                                            }
                                                        default:
                                                            var v42 int16 = 578
                                                            _ = v42
                                                        }
                                                    case 714:
                                                        fmt.Println(v3)
                                                    case 1:
                                                        switch fi30 {
                                                        case 0:
                                                            var v43 int8 = -128
                                                            _ = v43
                                                        }
                                                    default:
                                                        v44 := GenS4{F0: fi30, F1: 60, F2: 2}
                                                        _ = v44
                                                    }
                                                    var v45 int = -848
                                                    v46 := GenS2{F0: -1, F1: v7, F2: v2}
                                                    _ = v45
                                                    _ = v46
                                                }
                                                _ = v35
                                            }
                                            _ = v34
                                        }
                                    }
                                    v47 := GenS3{F0: true, F1: "hello"}
                                    if (1 == 127) {
                                        var v48 uint64 = 510
                                        if ((2147483647 >= -2147483648) || (v48 <= v48)) {
                                            v10 -= 1e10
                                            var v49 int = fi30
                                            var v50 uint32 = (^(v2))
                                            fmt.Println(fi33, v3)
                                            _ = v49
                                            _ = v50
                                        } else {
                                            var v51 string = "\n\t"
                                            _ = v51
                                        }
                                        var v52 int32 = 2147483647
                                        _ = v48
                                        _ = v52
                                    }
                                    v53 := GenS4{F0: fi30, F1: 18446744073709551615, F2: v2}
                                    _ = v47
                                    _ = v53
                                }
                            }
                        } else {
                            fmt.Println(v3, v2)
                            var v54 uint16 = 32767
                            if ((v2 == v2) || v7) {
                                v55 := GenS1{F0: fi30, F1: "", F2: 1}
                                _ = v55
                            } else {
                                var v56 int64 = -999
                                fmt.Println(v9, v10)
                                _ = v56
                            }
                            v9 = 0.0
                            _ = v54
                        }
                        v57 := GenS1{F0: (fi33 >> ((2) & 63)), F1: ("foo bar" + "\n\t"), F2: -327}
                        fmt.Println(v10, v8, v2)
                        _ = v57
                    }
                    _ = fi33
                }
            } else {
                v58 := GenS4{F0: (fi30 + fi30), F1: 42, F2: 2147483647}
                v59 := GenS3{F0: (!v7), F1: ("foo bar" + "foo bar")}
                fmt.Println(v8, v8, v10)
                _ = v58
                _ = v59
            }
        }
        v2 = v2
        _ = fi30
    }
    v60 := GenS1{F0: 127, F1: "foo bar", F2: -498}
    if false {
        fmt.Println(v3)
        v2 -= v2
    } else {
        var v61 float32 = float32(0.1)
        var v62 uint8 = 2
        v1 = append(v1, "\n\t")
        _ = v61
        _ = v62
    }
    var v63 int64 = -176
    for fi64 := 0; fi64 < 3; fi64++ {
        v10 -= v3
        v1 = append(v1, "a")
        v3 += v9
        _ = fi64
    }
    if v7 {
        switch ("\n\t" + "hello") {
        case "":
            _ = len(v1)
        case "hello":
            for fi65 := 0; fi65 < 14; fi65++ {
                switch (^(fi65)) {
                case -9223372036854775808:
                    var v66 uint8 = 128
                    _ = v66
                case 42:
                    for fi67 := 0; fi67 < 19; fi67++ {
                        var v68 int8 = 42
                        v3 = 0.0
                        fmt.Println(v2, v63)
                        _ = fi67
                        _ = v68
                    }
                case 471:
                    v9 = v3
                default:
                    var v69 int = 592
                    _ = v69
                }
                for fi70 := 0; fi70 < 18; fi70++ {
                    var v71 uint = 18446744073709551615
                    switch "a" {
                    case "\n\t":
                        if v7 {
                            fmt.Println(v3, v8, v10)
                            v7 = (!((0.1 / 1.5) <= 1.5))
                        }
                    case "hello":
                        fmt.Println(v8)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (1 <= 1) {
                        _ = len(v1)
                        v9 = 0.0
                        var v72 float64 = 0.1
                        _ = v72
                    } else {
                        var v73 uint = (v71 + v71)
                        v1 = append(v1, "foo bar")
                        v63 = genFn0(false)
                        _ = v73
                    }
                    _ = fi70
                    _ = v71
                }
                v7 = false
                _ = fi65
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi74 := 0; fi74 < 11; fi74++ {
            var v75 string = "hello"
            v76 := GenS3{F0: (!("a" > v75)), F1: v75}
            v77 := GenS5{F0: -1, F1: (v2 ^ (v2 % ((v2) | 1)))}
            _ = fi74
            _ = v75
            _ = v76
            _ = v77
        }
    } else {
        switch "hello" {
        case "hello":
            _ = len(v1)
        case "":
            if v7 {
                v78 := GenS6{F0: (2 != 0)}
                _ = v78
            }
        case "\n\t":
            fmt.Println(v7)
        default:
            var v79 float32 = v8
            _ = v79
        }
        for fi80 := 0; fi80 < 10; fi80++ {
            v10 = v9
            var v81 bool = v7
            fmt.Println(v9, fi80, v81)
            _ = fi80
            _ = v81
        }
    }
    if (-940 < -920) {
        var v82 float32 = ((v3 / 1.5) - v10)
        v8 += v10
        _ = v82
    } else {
        if (1 <= 213) {
            var v83 uint32 = v2
            if v7 {
                var v84 int64 = (v63 >> ((42) & 63))
                v85 := GenS7{F0: -1, F1: 255, F2: 2}
                _ = v84
                _ = v85
            } else {
                var v86 string = ("hello" + "a")
                _ = v86
            }
            _ = v83
        } else {
            v1 = append(v1, ("hello" + ""))
            _ = len(v1)
        }
        switch -771 {
        case -9223372036854775808:
            var v87 int64 = v63
            _ = v87
        case -893:
            switch 42 {
            case 0:
                for fi88 := 0; fi88 < 11; fi88++ {
                    v7 = v7
                    v89 := GenS1{F0: 972, F1: "\n\t", F2: 589}
                    _ = fi88
                    _ = v89
                }
            case -1:
                for fi90 := 0; fi90 < 10; fi90++ {
                    var v91 uint16 = 65535
                    _ = fi90
                    _ = v91
                }
            default:
                if ((3.14159 / 3.14159) == float64((3.14159 / (0.1 - v3)))) {
                    fmt.Println(v3)
                } else {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                }
            }
        case 32768:
            fmt.Println(v2, v9)
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    for fi92 := 0; fi92 < 3; fi92++ {
        v10 += (v8 + v9)
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        for fi93 := 0; fi93 < 5; fi93++ {
            var v94 int32 = 42
            _ = fi93
            _ = v94
        }
        _ = fi92
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v60
    _ = v63
    _ = v7
    _ = v8
    _ = v9
}
