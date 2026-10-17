package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int8, p1 bool, p2 int) bool {
    p2 = (p2 * int(p0))
    fmt.Println(p1)
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1() int64 {
    var v7 bool = true
    fmt.Println(v7)
    switch ("foo bar" + ("foo bar" + ("a" + "\n\t"))) {
    case "":
        v8 := []int{673, -1}
        _ = v8
    case "\n\t":
        var v9 uint16 = 478
        _ = v9
    case "foo bar":
        var v10 float32 = ((-2.25 - ((-2.25 - 0.1) - (0.1 / (0.0 - -2.25)))) * 1.5)
        _ = v10
    default:
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int8 = -128
            switch "" {
            case "foo bar":
                if (2147483647 == fi11) {
                    fmt.Println(v7, fi11)
                    fmt.Println(v12, v12)
                    v13 := []int{fi11, fi11, 269}
                    _ = v13
                } else {
                    v12 += (v12 >> ((979) & 63))
                    switch fi11 {
                    case 987:
                        for fi14 := 0; fi14 < 20; fi14++ {
                            switch "a" {
                            case "foo bar":
                                v15 := []int{fi14, fi11, -2147483648}
                                _ = v15
                            case "hello":
                                var v16 int8 = (v12 ^ (v12 & (v12 << ((337) & 63))))
                                _ = v16
                            case "a":
                                if (!false) {
                                    v17 := GenS1{F0: (v12 >> ((65535) & 63))}
                                    v18 := GenS2{F0: v7, F1: 401}
                                    v19 := GenS3{F0: 42, F1: (v12 | (v12 - (v12 | v12)))}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            default:
                                var v20 float32 = 0.0
                                _ = v20
                            }
                            var v21 int = fi11
                            if (!false) {
                                if v7 {
                                    var v22 int8 = 42
                                    v21 = (fi11 | (^(fi14)))
                                    _ = v22
                                }
                                v23 := []int{}
                                _ = v23
                            }
                            _ = fi14
                            _ = v21
                        }
                    case 128:
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v25 := GenS4{F0: int16(v12), F1: -857}
                            if ((!false) && v7) {
                                if ((fi24 | -2147483648) >= (fi11 % ((434) | 1))) {
                                    var v26 uint16 = 0
                                    _ = v26
                                } else {
                                    var v27 int8 = 1
                                    switch (^(fi24)) {
                                    case 0:
                                        if ((!true) || (2147483648 <= 500)) {
                                            v28 := GenS4{F0: 538, F1: 9223372036854775807}
                                            var v29 float32 = -2.25
                                            _ = v28
                                            _ = v29
                                        }
                                    case 128:
                                        v30 := []int{(fi11 << ((2) & 63)), fi11, fi24, (^(fi24))}
                                        _ = v30
                                    default:
                                        var v31 int = (fi11 + (fi24 + fi24))
                                        _ = v31
                                    }
                                    _ = v27
                                }
                                fmt.Println(v12, v12, v12)
                                if true {
                                    var v32 string = ""
                                    _ = v32
                                } else {
                                    var v33 uint = 42
                                    _ = v33
                                }
                                if true {
                                    var v34 int16 = -676
                                    v12 = -118
                                    if false {
                                        var v35 int8 = 127
                                        fmt.Println(v34, v12, fi11)
                                        _ = v35
                                    }
                                    _ = v34
                                } else {
                                    v36 := []int{(fi24 & (fi24 % ((0) | 1)))}
                                    var v37 string = "hello"
                                    _ = v36
                                    _ = v37
                                }
                                switch 4294967295 {
                                case 128:
                                    v12 -= (-(v12))
                                }
                                var v38 bool = (v7 || v7)
                                var v39 uint32 = 4294967295
                                _ = v38
                                _ = v39
                            } else {
                                var v40 int8 = v12
                                switch (fi24 / ((fi11) | 1)) {
                                case -2147483648:
                                    v41 := GenS5{F0: 87, F1: int64(fi24)}
                                    _ = v41
                                case 32768:
                                    v7 = true
                                default:
                                    var v42 int8 = v40
                                    _ = v42
                                }
                                _ = v40
                            }
                            _ = fi24
                            _ = v25
                        }
                    case 0:
                        v7 = false
                    default:
                        if v7 {
                            switch ("a" + "\n\t") {
                            case "":
                                switch "hello" {
                                case "\n\t":
                                    v43 := []int{(fi11 * (-(fi11))), fi11}
                                    _ = v43
                                default:
                                    switch ("\n\t" + ("a" + "\n\t")) {
                                    case "\n\t":
                                        v12 -= (-(v12))
                                    case "hello":
                                        for fi44 := 0; fi44 < 6; fi44++ {
                                            v7 = ((255 <= 61) || ((98 <= 255) || false))
                                            _ = fi44
                                        }
                                    }
                                }
                            default:
                                for fi45 := 0; fi45 < 13; fi45++ {
                                    v12 = (v12 % ((v12) | 1))
                                    var v46 int8 = (^(v12))
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            for fi47 := 0; fi47 < 6; fi47++ {
                                switch ("foo bar" + "foo bar") {
                                case "":
                                    v48 := []int{(-(fi11)), -514, 2147483647}
                                    _ = v48
                                case "hello":
                                    v7 = true
                                    v7 = v7
                                default:
                                    switch fi47 {
                                    case 9223372036854775807:
                                        var v49 int8 = -1
                                        _ = v49
                                    case -922:
                                        v50 := GenS6{F0: 552}
                                        _ = v50
                                    case -497:
                                        var v51 uint64 = 657
                                        _ = v51
                                    default:
                                        var v52 bool = true
                                        _ = v52
                                    }
                                }
                                _ = fi47
                            }
                            var v53 float64 = (float64(3.14159) + (0.0 * ((1.5 / 0.1) + -2.25)))
                            _ = v53
                        } else {
                            v54 := GenS4{F0: 1, F1: 42}
                            v12 += (v12 << ((255) & 63))
                            _ = v54
                        }
                    }
                    fmt.Println(v7)
                }
            default:
                v55 := GenS1{F0: v12}
                _ = v55
            }
            var v56 bool = v7
            _ = fi11
            _ = v12
            _ = v56
        }
    }
    _ = v7
    return 460
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{}
    var v3 float64 = 1.5
    var v4 uint64 = 360
    v5 := []string{("hello" + "hello"), ("\n\t" + "hello")}
    v6 := make(map[string]uint32)
    var v57 bool = (-1 <= 1)
    v1 -= v1
    for fi58 := 0; fi58 < 5; fi58++ {
        v3 -= v3
        _ = fi58
    }
    if true {
        if (0 != -1) {
            if v57 {
                var v59 int8 = 127
                var v60 uint8 = 128
                switch "a" {
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v59
                _ = v60
            }
        } else {
            var v61 float32 = 0.1
            switch "a" {
            case "a":
                for fi62 := 0; fi62 < 14; fi62++ {
                    fmt.Println(v3, fi62)
                    _ = fi62
                }
            }
            var v63 uint16 = 42
            _ = v61
            _ = v63
        }
    } else {
        v6[("a" + "")] = 399
        v4 = (^(v4))
    }
    var v64 float64 = v3
    for fi65 := 0; fi65 < 14; fi65++ {
        fmt.Println(fi65, v64)
        switch -2147483648 {
        case -904:
            switch "\n\t" {
            case "":
                v57 = (-2.25 == (v64 / 1.5))
            case "hello":
                if (!(-863 < -1)) {
                    var v66 uint32 = 0
                    v67 := GenS6{F0: (v4 << ((1) & 63))}
                    _ = v66
                    _ = v67
                }
            case "foo bar":
                v1 = 9223372036854775807
            }
        case 32767:
            if (2147483647 != 151) {
                v68 := GenS6{F0: (v4 >> ((103) & 63))}
                _ = v68
            } else {
                var v69 uint32 = 1
                switch (fi65 | (fi65 >> ((255) & 63))) {
                case 32768:
                    fmt.Println(v57, v4, fi65)
                default:
                    var v70 float64 = float64(((1e10 / 0.1) / 1.5))
                    _ = v70
                }
                _ = v69
            }
        case -9223372036854775808:
            delete(v6, (("foo bar" + "hello") + "a"))
        }
        var v71 int32 = -757
        _ = fi65
        _ = v71
    }
    var v72 int8 = 91
    for fi73 := 0; fi73 < 8; fi73++ {
        v64 += (float64(0.1) / 3.14159)
        for fi74 := 0; fi74 < 19; fi74++ {
            switch fi73 {
            case -9223372036854775808:
                var v75 float32 = 1.5
                _ = v75
            default:
                var v76 uint16 = 482
                _ = v76
            }
            _ = len(v2)
            _ = fi74
        }
        _ = fi73
    }
    var v77 float32 = float32((v3 * v3))
    fmt.Println(v3, v72, v72)
    for fi78 := 0; fi78 < 12; fi78++ {
        switch "hello" {
        case "foo bar":
            switch (fi78 >> ((255) & 63)) {
            case 127:
                v3 -= v64
            }
        case "hello":
            v3 -= -2.25
        case "\n\t":
            if false {
                var v79 uint = 2
                v77 = v77
                var v80 int32 = 42
                _ = v79
                _ = v80
            } else {
                var v81 float32 = v77
                var v82 int16 = -1
                var v83 int16 = v82
                _ = v81
                _ = v82
                _ = v83
            }
        }
        for fi84 := 0; fi84 < 12; fi84++ {
            var v85 uint32 = 412
            _ = fi84
            _ = v85
        }
        if true {
            if true {
                switch fi78 {
                case -601:
                    _ = len(v2)
                case 9223372036854775807:
                    var v86 int64 = 282
                    _ = v86
                case -2147483648:
                    var v87 uint = 18446744073709551615
                    _ = v87
                default:
                    switch 2147483647 {
                    case 4294967295:
                        var v88 string = "foo bar"
                        _ = v88
                    case -9223372036854775808:
                        switch fi78 {
                        case 878:
                            var v89 int = 42
                            _ = v89
                        case 89:
                            switch 0 {
                            case 128:
                                var v90 uint8 = 207
                                _ = v90
                            case 9223372036854775807:
                                v6[((("\n\t" + (("\n\t" + ("" + "\n\t")) + "\n\t")) + "hello") + "a")] = 4294967295
                            case 32767:
                                switch "a" {
                                case "hello":
                                    for fi91 := 0; fi91 < 11; fi91++ {
                                        fmt.Println(fi91, v72, v64)
                                        _ = fi91
                                    }
                                case "\n\t":
                                    v92 := GenS7{F0: (v4 | v4), F1: 1}
                                    _ = v92
                                case "":
                                    var v93 int32 = 475
                                    _ = v93
                                default:
                                    v2 = append(v2, ("hello" + "foo bar"))
                                }
                            }
                        }
                    case 32768:
                        v3 = v64
                    default:
                        v72 += v72
                    }
                }
            }
            v6["foo bar"] = 2147483648
        } else {
            v64 = v64
            fmt.Println(v1, v64, v64)
        }
        _ = fi78
    }
    var v94 int32 = 1
    fmt.Println(v94)
    for fi95 := 0; fi95 < 15; fi95++ {
        _ = len(v2)
        for fi96 := 0; fi96 < 20; fi96++ {
            var v97 bool = true
            fmt.Println(v4, v77)
            var v98 uint16 = 2
            _ = fi96
            _ = v97
            _ = v98
        }
        v1 = int64(v94)
        _ = fi95
    }
    switch "hello" {
    case "\n\t":
        var v99 uint32 = 303
        _ = v99
    case "":
        if true {
            if false {
                var v100 uint64 = (v4 - (v4 << ((2) & 63)))
                fmt.Println(v94)
                if (!v57) {
                    fmt.Println(v3)
                    v64 += v64
                }
                _ = v100
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v77 += v77
            var v101 int16 = 983
            switch 128 {
            case 32767:
                var v102 uint = 845
                _ = v102
            case -470:
                v2 = append(v2, "foo bar")
            default:
                v103 := GenS4{F0: int16(v72), F1: 42}
                _ = v103
            }
            _ = v101
        }
    default:
        if v57 {
            v104 := v6["foo bar"]
            fmt.Println(v1, v104, v64)
            _ = v104
        }
    }
    if v57 {
        v6["a"] = 1
        var v105 uint32 = 991
        _ = v105
    }
    v4 -= (v4 & v4)
    v3 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v72:", v72)
    fmt.Println("v77:", v77)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v64
    _ = v72
    _ = v77
    _ = v94
}
