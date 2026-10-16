package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint8; F1 uint8; F2 bool }
type GenS5 struct { F0 int; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint64; F1 uint; F2 string }
type GenS8 struct { F0 int8; F1 int8; F2 uint32 }
type GenS9 struct { F0 string; F1 uint; F2 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint; F1 uint; F2 int32 }
func genFn0(p0 uint64, p1 int8, p2 int64) string {
    var v5 int = 0
    var v6 uint8 = 112
    v7 := GenS1{F0: 42, F1: 0, F2: 0}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: v6, F1: -1, F2: 2147483647}
            _ = fi9
            _ = v10
        }
        var v11 bool = ((0.0 - 0.1) != 0.1)
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        switch ("a" + ("\n\t" + (("\n\t" + ("\n\t" + "")) + ((((("hello" + ("" + ("\n\t" + "hello"))) + "") + ("a" + "a")) + (("hello" + (("hello" + "foo bar") + ("\n\t" + ""))) + "")) + "hello")))) {
        case "hello":
            v13 := GenS1{F0: (v6 | (^(v6))), F1: -812, F2: 2147483647}
            _ = v13
        case "":
            if (!((^(p0)) > (p0 + p0))) {
                v14 := []int{v5, v5, 689}
                _ = v14
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if (!(65535 == 0)) {
                        var v16 uint16 = 2
                        _ = v16
                    } else {
                        if true {
                            switch (-(fi12)) {
                            case 384:
                                if (0 >= (p0 >> ((255) & 63))) {
                                    switch (fi12 << ((2) & 63)) {
                                    case 32767:
                                        if (!(("" + "\n\t") < "a")) {
                                            v17 := []int{fi15, (v5 & 9223372036854775807)}
                                            _ = len(v17)
                                            fmt.Println(fi12)
                                            p0 -= 42
                                            var v18 uint8 = v6
                                            _ = v17
                                            _ = v18
                                        } else {
                                            v6 += (v6 + 42)
                                            switch "" {
                                            case "":
                                                var v19 bool = true
                                                _ = v19
                                            }
                                            v20 := []int{(-(v5)), (fi15 % ((fi12) | 1)), (^(v5))}
                                            if len(v20) >= 2 {
                                                v20 = v20[1:len(v20)]
                                            }
                                            if false {
                                                if (42 == 4294967295) {
                                                    fmt.Println(v5)
                                                    var v21 uint32 = 42
                                                    _ = v21
                                                } else {
                                                    switch fi15 {
                                                    case 4294967295:
                                                        fmt.Println(p0, p2)
                                                    case 128:
                                                        var v22 float32 = ((1.5 + 0.0) * ((float32((-2.25 - (0.1 - ((1.5 + 0.1) + (((0.0 - 0.0) + (((1e10 - float64(0.1)) - (-2.25 * 3.14159)) * 1e10)) - 1.5))))) + (((1e10 + 1.5) + 3.14159) / ((3.14159 + 1e10) + float32(0.0)))) / 1.0))
                                                        switch ("" + "foo bar") {
                                                        case "foo bar":
                                                            v20 = append(v20, fi15)
                                                        case "hello":
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                        }
                                                        v23 := GenS1{F0: (v6 % ((231) | 1)), F1: 42, F2: 83}
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                    if (!(("" + (("\n\t" + ((("a" + "") + "hello") + "hello")) + "\n\t")) != "foo bar")) {
                                                        if (v6 <= 2) {
                                                            switch 128 {
                                                            case 128:
                                                                var v24 int32 = 42
                                                                _ = v24
                                                            }
                                                            fmt.Println(fi15, v6, fi15)
                                                        } else {
                                                            if (!((true && (471 >= 1)) || (!((p1 >> ((4294967295) & 63)) < int8(v5))))) {
                                                                p2 = (p2 ^ p2)
                                                                var v25 int32 = int32(p1)
                                                                if (((^(v6)) >= (v6 + 127)) || true) {
                                                                    switch 0 {
                                                                    case 127:
                                                                        v5 += int(v6)
                                                                    }
                                                                    switch v5 {
                                                                    case 127:
                                                                        var v26 int32 = -286
                                                                        _ = v26
                                                                    case -117:
                                                                        var v27 string = (("a" + ("foo bar" + (("a" + "\n\t") + (("hello" + "foo bar") + ("\n\t" + ""))))) + "\n\t")
                                                                        _ = v27
                                                                    default:
                                                                        p2 = p2
                                                                    }
                                                                    p2 = 0
                                                                }
                                                                p2 -= p2
                                                                _ = v25
                                                            } else {
                                                                v28 := GenS1{F0: (^(v6)), F1: -2147483648, F2: 0}
                                                                if (((p2 * 0) >= 9223372036854775807) || (p0 < p0)) {
                                                                    p0 += p0
                                                                    v29 := GenS1{F0: (v6 ^ v6), F1: 42, F2: 1}
                                                                    var v30 int = 32767
                                                                    _ = v29
                                                                    _ = v30
                                                                }
                                                                if true {
                                                                    var v31 string = "foo bar"
                                                                    _ = v31
                                                                }
                                                                var v32 int32 = 789
                                                                v33 := GenS1{F0: (v6 - v6), F1: v32, F2: 0}
                                                                _ = v28
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                        }
                                                        if true {
                                                            fmt.Println(p1, fi12)
                                                            v34 := GenS1{F0: v6, F1: 2147483647, F2: 2}
                                                            _ = v34
                                                        }
                                                        v6 -= 128
                                                    } else {
                                                        p1 += 63
                                                        switch v5 {
                                                        case -1:
                                                            if true {
                                                                v35 := GenS2{F0: "foo bar", F1: -2147483648, F2: 1}
                                                                _ = v35
                                                            } else {
                                                                var v36 int64 = (p2 & p2)
                                                                var v37 int8 = 75
                                                                v38 := GenS2{F0: ("\n\t" + ((("a" + "a") + ("" + "")) + "foo bar")), F1: 42, F2: (^(v37))}
                                                                _ = v36
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        case 127:
                                                            v39 := GenS3{F0: 0, F1: 97}
                                                            _ = v39
                                                        case 42:
                                                            switch (^(v5)) {
                                                            case -442:
                                                                var v40 string = "foo bar"
                                                                _ = v40
                                                            }
                                                        }
                                                        var v41 int = (^(fi12))
                                                        v42 := GenS1{F0: (^(v6)), F1: 0, F2: 363}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                }
                                                fmt.Println(p1, v6)
                                                v43 := GenS4{F0: v6, F1: v6, F2: true}
                                                _ = v43
                                            } else {
                                                if (980 >= p2) {
                                                    _ = len(v20)
                                                } else {
                                                    v44 := GenS3{F0: 2, F1: 280}
                                                    var v45 int16 = 42
                                                    _ = v44
                                                    _ = v45
                                                }
                                                v46 := GenS1{F0: v6, F1: 0, F2: 65535}
                                                _ = v46
                                            }
                                            _ = v20
                                        }
                                    case -1:
                                        v47 := []int{(fi12 | v5), (fi15 - v5), fi15, (^(fi15))}
                                        _ = v47
                                    }
                                    v48 := []int{4294967295, fi12}
                                    _ = v48
                                } else {
                                    v49 := []int{(-(fi15)), fi12}
                                    if ((42 >= 181) || (!true)) {
                                        if (true || (9223372036854775807 <= 42)) {
                                            var v50 int16 = -949
                                            v51 := GenS5{F0: (v5 << ((42) & 63)), F1: 42, F2: 2}
                                            _ = v50
                                            _ = v51
                                        } else {
                                            p1 -= (p1 << ((255) & 63))
                                            var v52 bool = true
                                            _ = v52
                                        }
                                    }
                                    v53 := GenS6{F0: 2, F1: 32767, F2: 42}
                                    _ = v49
                                    _ = v53
                                }
                            case 1:
                                v54 := GenS7{F0: (^(p0)), F1: 2, F2: "a"}
                                _ = v54
                            case -600:
                                v55 := []int{(fi12 ^ v5), (^(fi12))}
                                _ = v55
                            default:
                                v56 := []int{(fi15 - (v5 + 0)), (fi12 * (fi15 * (v5 | v5))), fi15, v5}
                                _ = v56
                            }
                            fmt.Println(fi15, p1, v6)
                            if (false || (-9223372036854775808 < 4294967295)) {
                                var v57 string = "hello"
                                v58 := GenS4{F0: 0, F1: (^(v6)), F2: ((fi12 - 32767) >= (fi12 * fi12))}
                                _ = v57
                                _ = v58
                            }
                        }
                    }
                    if (!(true || false)) {
                        var v59 float32 = (1e10 - 1.5)
                        _ = v59
                    } else {
                        if ((-2.25 < 1e10) || false) {
                            switch "" {
                            case "\n\t":
                                if (-1 <= -2147483648) {
                                    var v60 float32 = 1e10
                                    var v61 int8 = (p1 / (((p1 % ((100) | 1))) | 1))
                                    var v62 int = fi12
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                } else {
                                    p1 = 127
                                }
                            case "hello":
                                v63 := []int{}
                                _ = v63
                            default:
                                if (42 >= (v6 + uint8(p0))) {
                                    v64 := GenS2{F0: ("a" + ""), F1: 1, F2: p1}
                                    var v65 float64 = -2.25
                                    v66 := []int{(fi12 ^ fi12)}
                                    v5 = (v5 & 0)
                                    switch (-(v5)) {
                                    case -2147483648:
                                        if len(v66) >= 2 {
                                            v66 = v66[1:len(v66)]
                                        }
                                    case -922:
                                        if false {
                                            v67 := GenS8{F0: (p1 ^ (p1 + p1)), F1: (p1 - p1), F2: 2147483648}
                                            var v68 string = ("" + (("a" + "hello") + ("foo bar" + "\n\t")))
                                            var v69 int = (fi12 - fi15)
                                            var v70 uint16 = 2
                                            v71 := GenS5{F0: -643, F1: 0, F2: 42}
                                            fmt.Println(v65)
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                        } else {
                                            if false {
                                                v65 -= v65
                                            }
                                        }
                                        var v72 string = (("hello" + ("foo bar" + "a")) + "foo bar")
                                        _ = v72
                                    }
                                    v6 -= (^(v6))
                                    v6 = (v6 - (v6 % ((v6) | 1)))
                                    var v73 uint16 = 201
                                    fmt.Println(v73, v73, v73)
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                    _ = v73
                                } else {
                                    v74 := []int{42, -9223372036854775808, (v5 % ((v5) | 1))}
                                    v5 -= (fi12 << ((65535) & 63))
                                    var v75 int = fi15
                                    _ = v74
                                    _ = v75
                                }
                            }
                        } else {
                            var v76 uint = 65535
                            if (int16(v6) == -1) {
                                v77 := []int{fi15, fi12, -400, fi12}
                                switch (fi15 >> (((^(v76))) & 63)) {
                                case 1:
                                    v78 := GenS7{F0: (p0 ^ (p0 << ((v76) & 63))), F1: (^(v76)), F2: "a"}
                                    _ = v78
                                case 4294967295:
                                    if false {
                                        if len(v77) >= 2 {
                                            v77 = v77[1:len(v77)]
                                        }
                                        v6 -= (v6 | 121)
                                        var v79 uint64 = p0
                                        p1 -= p1
                                        _ = v79
                                    }
                                case -959:
                                    fmt.Println(p1, p1, p2)
                                default:
                                    var v80 string = ("a" + "a")
                                    _ = v80
                                }
                                _ = v77
                            } else {
                                v81 := []int{(-(v5)), (fi15 - (v5 & (fi12 + 32768)))}
                                _ = v81
                            }
                            v82 := GenS9{F0: "\n\t", F1: v76, F2: 1}
                            fmt.Println(p1, p2)
                            _ = v76
                            _ = v82
                        }
                        var v83 float64 = -2.25
                        fmt.Println(p1, v5)
                        _ = v83
                    }
                    fmt.Println(v5)
                    _ = fi15
                }
                for fi84 := 0; fi84 < 20; fi84++ {
                    if true {
                        fmt.Println(fi12, p1)
                        v85 := []int{4294967295, v5, 2147483647, -230}
                        _ = v85
                    }
                    _ = fi84
                }
                var v86 int8 = 107
                _ = v86
            }
        case "\n\t":
            switch "foo bar" {
            case "":
                p2 = p2
            }
        default:
            switch -1 {
            case 9223372036854775807:
                var v87 float32 = (1.5 - (1e10 * 3.14159))
                _ = v87
            }
        }
        v88 := GenS6{F0: 42, F1: 809, F2: 366}
        v89 := GenS1{F0: 67, F1: 42, F2: 892}
        _ = fi12
        _ = v88
        _ = v89
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (("a" + "hello") + "")
}

func main() {
    v1 := []int{377, 312}
    var v2 int32 = 42
    var v3 uint16 = 1
    var v4 bool = false
    var v90 string = "hello"
    switch -212 {
    case 88:
        v2 = (-(v2))
    case 266:
        v91 := GenS10{F0: 937}
        _ = v91
    default:
        if true {
            var v92 bool = v4
            v4 = v92
            v92 = false
            _ = v92
        }
    }
    if (42 >= 128) {
        var v93 int32 = (v2 - 0)
        _ = v93
    } else {
        if len(v1) > 0 {
            v1[0] = 985
        }
        var v94 float64 = -2.25
        fmt.Println(v4, v3, v3)
        _ = v94
    }
    var v95 int8 = 100
    var v96 uint16 = 376
    var v97 int = 128
    v98 := GenS3{F0: v96, F1: 65535}
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = v97
    }
    if ((false && v4) || (uint64(v3) >= 246)) {
        v2 = (v2 + (v2 - v2))
        var v99 int32 = v2
        if (!v4) {
            v100 := GenS11{F0: v90}
            v101 := GenS12{F0: 954, F1: 4294967295, F2: v2}
            v2 = (v99 - (v99 - v99))
            _ = v100
            _ = v101
        }
        _ = v99
    }
    var v102 string = v90
    var v103 int32 = v2
    switch ("hello" + ("foo bar" + v90)) {
    case "foo bar":
        fmt.Println(v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v90:", v90)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v90
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
