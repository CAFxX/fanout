package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 int8; F2 bool }
type GenS10 struct { F0 uint8; F1 int16 }
func genFn0(p0 int64, p1 int32) bool {
    p0 = p0
    v6 := GenS1{F0: 4294967295, F1: 1, F2: true}
    var v7 int32 = (p1 - 848)
    if false {
        var v8 int8 = 100
        _ = v8
    }
    v9 := []int{807, 0, -537, 127}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v9
    return (p0 >= 0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 2
    var v3 float64 = ((0.1 / 1e10) / 1.0)
    var v4 int32 = 111
    v5 := make(map[string]uint32)
    v2 += 128
    v10 := v5["foo bar"]
    var v11 uint16 = 480
    var v12 int32 = (v4 >> ((1) & 63))
    switch 48 {
    case 797:
        if (false && (1 != 355)) {
            var v13 float32 = (float32(((v3 + 1.5) - v3)) - 0.0)
            var v14 int16 = 32767
            var v15 uint = 2
            _ = v13
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: v2}
            _ = v16
        }
    case 159:
        if false {
            var v17 int32 = v12
            var v18 int8 = 100
            if true {
                switch "" {
                case "foo bar":
                    switch "hello" {
                    case "hello":
                        v19 := GenS1{F0: 42, F1: -2147483648, F2: ((v18 < -1) && false)}
                        _ = v19
                    case "foo bar":
                        if (v3 <= (0.1 + float64(1e10))) {
                            fmt.Println(v3, v11)
                            var v20 int8 = int8(v4)
                            var v21 int8 = v18
                            _ = v20
                            _ = v21
                        }
                    case "a":
                        v12 -= (v12 & 2147483647)
                    }
                case "hello":
                    var v22 int8 = (v18 + (v18 | -37))
                    _ = v22
                }
                switch 42 {
                case -9223372036854775808:
                    var v23 uint32 = (v10 / ((v10) | 1))
                    _ = v23
                case -1:
                    v24 := []int{-536, 1}
                    _ = v24
                case 127:
                    v17 -= (v4 % ((v4) | 1))
                default:
                    v18 -= v18
                }
                fmt.Println(v2)
            }
            _ = v17
            _ = v18
        }
    }
    switch "a" {
    case "\n\t":
        v25 := []int{9223372036854775807, -9223372036854775808}
        _ = v25
    }
    _ = len(v1)
    delete(v1, (v10 / ((v10) | 1)))
    v26 := GenS2{F0: 127}
    var v27 uint8 = 111
    v28 := []int{32767, 90}
    switch ("foo bar" + "") {
    case "a":
        switch 2147483647 {
        case 32768:
            if false {
                var v29 string = (("a" + "a") + "a")
                for fi30 := 0; fi30 < 7; fi30++ {
                    if (true || true) {
                        v31 := GenS2{F0: v27}
                        _ = v31
                    } else {
                        if (false && true) {
                            if false {
                                fmt.Println(v4)
                                if (false && true) {
                                    switch fi30 {
                                    case -1:
                                        var v32 bool = true
                                        _ = v32
                                    default:
                                        for fi33 := 0; fi33 < 20; fi33++ {
                                            v11 -= 2
                                            var v34 int64 = 0
                                            _ = fi33
                                            _ = v34
                                        }
                                    }
                                }
                            } else {
                                v28 = append(v28, fi30)
                                var v35 uint16 = (v11 >> ((0) & 63))
                                if len(v28) > 0 {
                                    v28[0] = 1
                                }
                                _ = v35
                            }
                            if false {
                                fmt.Println(v12)
                                for fi36 := 0; fi36 < 13; fi36++ {
                                    v37 := GenS2{F0: v27}
                                    if false {
                                        v38 := GenS1{F0: fi36, F1: fi30, F2: ((v10 - 4294967295) < v10)}
                                        var v39 float64 = v3
                                        _ = v38
                                        _ = v39
                                    } else {
                                        v40 := v5[("hello" + (v29 + ("" + v29)))]
                                        _ = v40
                                    }
                                    v11 = v11
                                    _ = fi36
                                    _ = v37
                                }
                                var v41 uint32 = 521
                                _ = v41
                            } else {
                                if true {
                                    v27 += v2
                                    if (int8(fi30) >= int8(v27)) {
                                        v4 -= v12
                                        var v42 uint64 = 4294967295
                                        v43 := GenS3{F0: v12, F1: 255}
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                fmt.Println(v29, v27)
                                fmt.Println(v10, v29)
                            }
                            switch (^(fi30)) {
                            case 42:
                                for fi44 := 0; fi44 < 1; fi44++ {
                                    var v45 int = 376
                                    _ = fi44
                                    _ = v45
                                }
                            case 128:
                                v46 := GenS1{F0: fi30, F1: (fi30 >> ((264) & 63)), F2: (false || true)}
                                _ = v46
                            case -9223372036854775808:
                                var v47 bool = (v12 <= (-(v12)))
                                _ = v47
                            default:
                                fmt.Println(v3)
                            }
                        }
                    }
                    _ = fi30
                }
                for fi48 := 0; fi48 < 10; fi48++ {
                    switch fi48 {
                    case 32768:
                        var v49 uint16 = 0
                        _ = v49
                    }
                    var v50 float32 = 1.5
                    for fi51 := 0; fi51 < 2; fi51++ {
                        var v52 bool = (-118 < 100)
                        _ = fi51
                        _ = v52
                    }
                    _ = fi48
                    _ = v50
                }
                _ = v29
            }
        default:
            for fi53 := 0; fi53 < 17; fi53++ {
                if (v10 >= 485) {
                    if (42 < 42) {
                        _ = len(v28)
                    }
                    if (true || ("" > "")) {
                        for fi54 := 0; fi54 < 9; fi54++ {
                            if false {
                                v55 := GenS1{F0: 2147483647, F1: fi54, F2: (74 >= -590)}
                                var v56 uint = 1
                                if (!false) {
                                    fmt.Println(fi54)
                                    if (true && (!((0.0 - 0.1) >= v3))) {
                                        var v57 int8 = int8(fi54)
                                        _ = v57
                                    } else {
                                        v27 += (v27 * 127)
                                        var v58 uint32 = v10
                                        if false {
                                            v59 := GenS1{F0: fi53, F1: (-(fi54)), F2: (2 < 2)}
                                            if ((0 != 0) && true) {
                                                var v60 uint64 = 1
                                                if (9223372036854775807 <= 1) {
                                                    var v61 string = ("" + "a")
                                                    v12 -= v4
                                                    var v62 float32 = float32(3.14159)
                                                    _ = v61
                                                    _ = v62
                                                }
                                                _ = v60
                                            }
                                            switch fi54 {
                                            case 213:
                                                if (true && (v27 != 127)) {
                                                    v27 = v27
                                                }
                                            case 0:
                                                switch ((("" + ((("foo bar" + ("" + "\n\t")) + "a") + ("hello" + ""))) + "\n\t") + "foo bar") {
                                                case "hello":
                                                    var v63 string = "a"
                                                    _ = v63
                                                case "\n\t":
                                                    v28 = append(v28, (fi53 % (((fi53 & 2147483647)) | 1)))
                                                    var v64 uint64 = 42
                                                    v3 = 0.1
                                                    _ = v64
                                                case "a":
                                                    v65 := GenS3{F0: int32(v11), F1: (v56 & v56)}
                                                    _ = v65
                                                default:
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch fi53 {
                                                        case 32767:
                                                            if (252 >= 42) {
                                                                fmt.Println(v12)
                                                            } else {
                                                                v12 = -523
                                                            }
                                                        case -370:
                                                            switch fi54 {
                                                            case 128:
                                                                v66 := GenS1{F0: 127, F1: 0, F2: false}
                                                                var v67 uint = v56
                                                                var v68 float32 = 0.0
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            default:
                                                                switch fi53 {
                                                                case 479:
                                                                    if false {
                                                                        var v69 string = ("hello" + "foo bar")
                                                                        fmt.Println(fi53, v12)
                                                                        _ = v69
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        var v70 bool = true
                                                                        v28 = append(v28, (-(fi53)))
                                                                        _ = v70
                                                                    }
                                                                case -9223372036854775808:
                                                                    switch fi53 {
                                                                    case 32768:
                                                                        v71 := GenS2{F0: 2}
                                                                        _ = v71
                                                                    }
                                                                default:
                                                                    var v72 int64 = -9223372036854775808
                                                                    _ = v72
                                                                }
                                                            }
                                                        default:
                                                            if ((((1e10 - v3) - (float64(((3.14159 - ((1.5 * 1.5) + 1.5)) - 0.0)) - v3)) - float64(-2.25)) == v3) {
                                                                if ((v11 != v11) && (!true)) {
                                                                    switch fi54 {
                                                                    case -534:
                                                                        v73 := GenS4{F0: 42}
                                                                        _ = v73
                                                                    case 0:
                                                                        var v74 float32 = float32((1.5 / v3))
                                                                        var v75 string = "hello"
                                                                        var v76 int8 = 65
                                                                        _ = v74
                                                                        _ = v75
                                                                        _ = v76
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            if (0 == v4) {
                                                                                fmt.Println(fi53, v11)
                                                                                v56 -= v56
                                                                            }
                                                                        }
                                                                    }
                                                                    var v77 uint16 = v11
                                                                    _ = v77
                                                                } else {
                                                                    v78 := GenS1{F0: (fi53 << ((2) & 63)), F1: 9223372036854775807, F2: (!(!((0 >= 18446744073709551615) && (false && true))))}
                                                                    _ = v78
                                                                }
                                                                v79 := GenS5{F0: (^(v2))}
                                                                if (1000 > 42) {
                                                                    _ = len(v28)
                                                                    var v80 uint8 = v27
                                                                    if (v11 > v11) {
                                                                        var v81 int64 = 1
                                                                        if len(v28) > 0 {
                                                                            v28[0] = fi54
                                                                        }
                                                                        _ = v81
                                                                    }
                                                                    _ = v80
                                                                } else {
                                                                    var v82 int64 = int64(v4)
                                                                    v11 = (v11 << ((v56) & 63))
                                                                    if ((v4 & 0) != v12) {
                                                                        fmt.Println(fi54)
                                                                        _ = len(v1)
                                                                        v83 := GenS5{F0: 128}
                                                                        _ = v83
                                                                    } else {
                                                                        v84 := GenS6{F0: 2147483648}
                                                                        v28 = append(v28, -692)
                                                                        v28 = append(v28, fi53)
                                                                        _ = v84
                                                                    }
                                                                    _ = v82
                                                                }
                                                                _ = v79
                                                            }
                                                        }
                                                    case "a":
                                                        if true {
                                                            delete(v5, ("a" + ("a" + "a")))
                                                            v85 := GenS7{F0: (false || (((^(v11)) <= 377) && ((((0.0 - 0.0) / (float64((0.0 / 1.0)) * v3)) / ((((((v3 * float64(-2.25)) * v3) - 0.0) / v3) + v3) + float64(-2.25))) == v3))), F1: (^(v10))}
                                                            if false {
                                                                switch ("" + "foo bar") {
                                                                case "a":
                                                                    var v86 uint16 = v11
                                                                    _ = v86
                                                                case "hello":
                                                                    _ = len(v28)
                                                                }
                                                                switch 2147483647 {
                                                                case 9223372036854775807:
                                                                    switch fi53 {
                                                                    case 32767:
                                                                        v4 -= v12
                                                                    case 4294967295:
                                                                        var v87 uint64 = 42
                                                                        _ = v87
                                                                    }
                                                                default:
                                                                    v88 := GenS1{F0: fi54, F1: fi53, F2: ((true || true) || true)}
                                                                    _ = v88
                                                                }
                                                            } else {
                                                                var v89 int32 = v12
                                                                _ = v89
                                                            }
                                                            _ = v85
                                                        }
                                                    }
                                                }
                                            case 768:
                                                var v90 int32 = v4
                                                _ = v90
                                            default:
                                                v11 = v11
                                            }
                                            _ = v59
                                        } else {
                                            var v91 uint = v56
                                            v92 := GenS8{F0: fi53}
                                            v10 -= v10
                                            _ = v91
                                            _ = v92
                                        }
                                        _ = v58
                                    }
                                }
                                _ = v55
                                _ = v56
                            }
                            v10 = 2
                            _ = fi54
                        }
                        fmt.Println(v10)
                        v93 := GenS9{F0: 29, F1: 1, F2: ((("hello" + "") + "\n\t") <= ("a" + "\n\t"))}
                        _ = v93
                    } else {
                        v94 := GenS10{F0: v27, F1: int16(v2)}
                        var v95 float64 = v3
                        _ = v94
                        _ = v95
                    }
                }
                _ = fi53
            }
        }
    case "hello":
        delete(v1, 42)
    default:
        var v96 bool = ((((0.0 * -2.25) / (-2.25 / -2.25)) > (0.1 * 1.5)) && true)
        _ = v96
    }
    for fi97 := 0; fi97 < 14; fi97++ {
        for fi98 := 0; fi98 < 6; fi98++ {
            var v99 uint8 = (v2 % ((v2) | 1))
            var v100 int64 = -672
            var v101 uint32 = v10
            _ = fi98
            _ = v100
            _ = v101
            _ = v99
        }
        v10 += 466
        _ = fi97
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
