package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 int64; F1 int8 }
func genFn0(p0 uint, p1 uint8, p2 int32) uint64 {
    v7 := []int{42}
    p2 = p2
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int8 = -90
    v2 := make(map[string]uint32)
    var v3 float32 = (0.0 + -2.25)
    var v4 uint8 = 85
    var v5 uint64 = 42
    var v6 uint8 = (v4 >> ((255) & 63))
    v5 = v5
    v5 = v5
    if (false && (145 >= 2)) {
        if (-2.25 == 0.0) {
            v4 = (v4 - v6)
        }
        v1 = (v1 & v1)
    }
    v8 := GenS1{F0: -1, F1: 2}
    fmt.Println(v6, v6)
    if false {
        var v9 int32 = 754
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint = 4294967295
            v12 := []int{(fi10 + (^(fi10))), (fi10 - fi10), (fi10 * fi10), fi10}
            var v13 uint16 = 42
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint64 = v5
                var v17 float64 = (0.1 - 1e10)
                v5 -= (v5 * v5)
                _ = fi15
                _ = v16
                _ = v17
            }
            v3 = ((1.5 * 0.0) - ((0.0 * v3) * v3))
            _ = fi14
        }
        _ = v9
    } else {
        switch 157 {
        case 1:
            v18 := GenS2{F0: 2, F1: 0, F2: 0}
            _ = v18
        case 4294967295:
            v19 := []int{}
            _ = v19
        }
        v5 = (v5 | v5)
        fmt.Println(v6, v5, v1)
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS1{F0: -32768, F1: 0}
        if true {
            v22 := []int{-9223372036854775808, fi20, 127, 42}
            _ = v22
        } else {
            v23 := GenS3{F0: v5}
            _ = v23
        }
        for fi24 := 0; fi24 < 16; fi24++ {
            if ((v5 | 4) < 2) {
                v25 := v2["a"]
                v26 := []int{(fi20 & (fi24 % ((fi24) | 1))), fi24, 1, (^(fi24))}
                var v27 int = (fi20 % (((fi20 & (fi24 + fi24))) | 1))
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi24
        }
        _ = fi20
        _ = v21
    }
    if true {
        for fi28 := 0; fi28 < 14; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                switch 9223372036854775807 {
                case 1:
                    var v30 int32 = 341
                    switch "foo bar" {
                    case "\n\t":
                        var v31 uint = 0
                        _ = v31
                    default:
                        var v32 uint64 = (v5 - (v5 - 18446744073709551615))
                        _ = v32
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v33 int64 = -1
                        _ = v33
                    case "a":
                        v4 = 1
                    default:
                        switch "foo bar" {
                        case "a":
                            var v34 uint16 = 42
                            _ = v34
                        case "":
                            v6 = v4
                            if false {
                                v6 = (v6 / ((v6) | 1))
                            }
                        case "\n\t":
                            var v35 string = ("foo bar" + "foo bar")
                            _ = v35
                        }
                    }
                    _ = v30
                default:
                    if (!false) {
                        if true {
                            v36 := []int{fi28, int(v5), (^(fi29))}
                            _ = v36
                        } else {
                            var v37 uint8 = 1
                            _ = v37
                        }
                        if true {
                            if (42 >= 0) {
                                if true {
                                    v38 := GenS1{F0: 32767, F1: 255}
                                    fmt.Println(v3, v6)
                                    v39 := GenS4{F0: 1, F1: 1, F2: int32(fi29)}
                                    _ = v38
                                    _ = v39
                                } else {
                                    v4 += v4
                                    switch ("a" + ("hello" + ("a" + "hello"))) {
                                    case "hello":
                                        var v40 string = ((("\n\t" + "\n\t") + (("foo bar" + "foo bar") + ("\n\t" + "a"))) + (("\n\t" + ("a" + "hello")) + "hello"))
                                        _ = v40
                                    case "\n\t":
                                        switch (-(fi28)) {
                                        case 42:
                                            v5 = 4294967295
                                        case -2147483648:
                                            fmt.Println(v6, fi29)
                                        default:
                                            if true {
                                                v41 := GenS3{F0: genFn0(0, (v6 + (^(v6))), 42)}
                                                _ = v41
                                            }
                                        }
                                    default:
                                        v42 := GenS1{F0: 42, F1: 65535}
                                        _ = v42
                                    }
                                    var v43 float32 = float32((float64(float32(-2.25)) / (1e10 / 1.0)))
                                    var v44 uint8 = 42
                                    _ = v43
                                    _ = v44
                                }
                                v45 := []int{}
                                _ = v45
                            } else {
                                v3 = float32(1.5)
                                switch "\n\t" {
                                case "":
                                    v3 = float32(1e10)
                                case "foo bar":
                                    if (1 <= -25) {
                                        var v46 int16 = 1000
                                        v47 := v2["\n\t"]
                                        _ = v46
                                        _ = v47
                                    } else {
                                        fmt.Println(v5, fi28)
                                    }
                                case "\n\t":
                                    if true {
                                        var v48 uint64 = (^(v5))
                                        _ = v48
                                    } else {
                                        v5 = (v5 >> ((42) & 63))
                                    }
                                }
                            }
                            var v49 string = (("\n\t" + "foo bar") + "")
                            if (1 <= 708) {
                                v50 := GenS5{F0: 161, F1: -272}
                                var v51 float32 = (((((v3 + (float32(((0.1 * 0.1) + 3.14159)) / ((v3 - (v3 / 1.0)) + (v3 + v3)))) - v3) + 0.0) / v3) - v3)
                                _ = v50
                                _ = v51
                            }
                            var v52 float64 = 0.1
                            _ = v49
                            _ = v52
                        } else {
                            v1 = 1
                            if false {
                                fmt.Println(v5, v5)
                            } else {
                                if (!false) {
                                    var v53 bool = (-930 > 1)
                                    fmt.Println(v1)
                                    _ = v53
                                }
                            }
                            v54 := GenS3{F0: v5}
                            _ = v54
                        }
                        if false {
                            v55 := GenS5{F0: 255, F1: -1}
                            if (v6 <= v4) {
                                if false {
                                    if ((!(!(int64(v5) <= 345))) || true) {
                                        v56 := GenS6{F0: 656, F1: 794}
                                        fmt.Println(v5, v4)
                                        if true {
                                            fmt.Println(v5)
                                            v57 := GenS7{F0: 0}
                                            _ = v57
                                        } else {
                                            var v58 int64 = 0
                                            v3 -= ((v3 / v3) - v3)
                                            switch ("\n\t" + (((((("\n\t" + ("\n\t" + "a")) + "hello") + "foo bar") + "hello") + "\n\t") + ("" + "hello"))) {
                                            case "foo bar":
                                                var v59 int8 = (^(v1))
                                                _ = v59
                                            case "hello":
                                                v1 -= (v1 * v1)
                                            case "":
                                                v5 -= v5
                                            }
                                            _ = v58
                                        }
                                        _ = v56
                                    }
                                    switch (fi28 ^ (fi29 / ((fi28) | 1))) {
                                    case 9223372036854775807:
                                        switch "\n\t" {
                                        case "\n\t":
                                            v4 += 33
                                        case "hello":
                                            v60 := []int{}
                                            _ = v60
                                        case "":
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v61 int8 = (v1 << ((517) & 63))
                                                _ = v61
                                            case "hello":
                                                var v62 int = fi28
                                                v6 = 211
                                                var v63 int8 = v1
                                                _ = v62
                                                _ = v63
                                            default:
                                                fmt.Println(v6, fi29)
                                            }
                                        }
                                    case 128:
                                        v6 += v6
                                    }
                                    v4 = (v6 * (v6 << ((42) & 63)))
                                } else {
                                    if (198 < 494) {
                                        v64 := GenS2{F0: 4294967295, F1: 1, F2: 1}
                                        fmt.Println(fi28, fi28)
                                        _ = v64
                                    } else {
                                        v5 += 370
                                    }
                                    if ("" <= ("" + "")) {
                                        if (true || (v3 < v3)) {
                                            fmt.Println(fi28, v4, v3)
                                            var v65 uint16 = 0
                                            _ = v65
                                        }
                                    } else {
                                        v4 -= (v4 % ((v4) | 1))
                                        v6 = v6
                                    }
                                }
                                v1 = (v1 << ((4294967295) & 63))
                                v66 := GenS8{F0: -2147483648, F1: 2147483647}
                                v67 := GenS2{F0: 2147483648, F1: 1, F2: 189}
                                var v68 int32 = 0
                                if ((float64(3.14159) > 3.14159) || false) {
                                    var v69 int16 = 1000
                                    switch (fi28 >> ((65535) & 63)) {
                                    case 724:
                                        if (4294967295 != 2) {
                                            var v70 int = 127
                                            _ = v70
                                        }
                                    case 42:
                                        var v71 string = (("a" + "foo bar") + ("hello" + ((("" + ("hello" + "")) + "\n\t") + "")))
                                        _ = v71
                                    case 4294967295:
                                        var v72 bool = (fi29 > (fi28 & (-(fi28))))
                                        _ = v72
                                    default:
                                        switch int(v5) {
                                        case 9223372036854775807:
                                            var v73 uint = 2
                                            if (false || false) {
                                                if (uint(v4) >= 682) {
                                                    v74 := GenS3{F0: (v5 / (((^(v5))) | 1))}
                                                    _ = v74
                                                } else {
                                                    v75 := GenS4{F0: 42, F1: 42, F2: v68}
                                                    var v76 uint64 = (^(v5))
                                                    v77 := GenS3{F0: (^(v76))}
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                v78 := GenS9{F0: (fi28 / ((fi28) | 1)), F1: 716}
                                                switch fi29 {
                                                case -1:
                                                    fmt.Println(fi28)
                                                case 32767:
                                                    v79 := []int{}
                                                    _ = v79
                                                case -9223372036854775808:
                                                    var v80 uint8 = v4
                                                    _ = v80
                                                }
                                                _ = v78
                                            } else {
                                                v81 := []int{2147483647}
                                                if false {
                                                    var v82 float32 = (((v3 - v3) - (1e10 + v3)) / v3)
                                                    fmt.Println(v4, v4, v73)
                                                    switch "" {
                                                    case "foo bar":
                                                        v83 := GenS9{F0: (fi29 + 127), F1: 42}
                                                        switch "" {
                                                        case "foo bar":
                                                            v84 := GenS10{F0: -1, F1: 42}
                                                            _ = v84
                                                        case "":
                                                            var v85 uint64 = (v5 << ((uint(v4)) & 63))
                                                            _ = v85
                                                        }
                                                        _ = v83
                                                    }
                                                    _ = v82
                                                }
                                                _ = v81
                                            }
                                            _ = v73
                                        }
                                    }
                                    _ = v69
                                } else {
                                    v68 += v68
                                    v86 := []int{(fi29 * (^(fi28))), -1, (fi28 | fi29), (fi28 + -1)}
                                    _ = v86
                                }
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                            _ = v55
                        }
                    }
                }
                _ = fi29
            }
            _ = len(v2)
            v5 = (v5 | 0)
            _ = fi28
        }
        var v87 int64 = -616
        v88 := GenS8{F0: 575, F1: 65535}
        _ = v87
        _ = v88
    }
    var v89 uint16 = 255
    var v90 float32 = float32((1e10 - 0.0))
    var v91 float32 = (0.0 + v3)
    var v92 int32 = 395
    switch "foo bar" {
    case "a":
        var v93 float32 = 1.5
        _ = v93
    default:
        for fi94 := 0; fi94 < 4; fi94++ {
            var v95 bool = false
            switch (fi94 * 671) {
            case 2147483647:
                for fi96 := 0; fi96 < 11; fi96++ {
                    v3 = (float32(0.1) - (v3 * (v3 * v91)))
                    switch "hello" {
                    case "\n\t":
                        v91 = float32(0.0)
                    case "":
                        v1 = int8(v89)
                    default:
                        var v97 int = (fi94 % (((fi94 >> ((2) & 63))) | 1))
                        _ = v97
                    }
                    _ = fi96
                }
            }
            _ = fi94
            _ = v95
        }
    }
    if ((v89 / (((v89 - (^(v89)))) | 1)) < 32767) {
        if false {
            for fi98 := 0; fi98 < 20; fi98++ {
                var v99 uint64 = 350
                for fi100 := 0; fi100 < 7; fi100++ {
                    if (32767 >= -32768) {
                        var v101 float32 = 1e10
                        v6 = v6
                        _ = v101
                    } else {
                        v3 -= v91
                        v4 += v6
                        v6 -= v6
                    }
                    fmt.Println(fi100)
                    _ = fi100
                }
                v3 -= -2.25
                _ = fi98
                _ = v99
            }
            var v102 int8 = v1
            _ = v102
        }
        v103 := []int{372}
        _ = v103
    } else {
        if (v5 >= 0) {
            var v104 uint16 = (v89 >> ((15) & 63))
            for fi105 := 0; fi105 < 5; fi105++ {
                if (0.0 >= v91) {
                    if false {
                        v106 := GenS10{F0: 0, F1: (-(v1))}
                        _ = v106
                    } else {
                        fmt.Println(v92, v5)
                    }
                    if true {
                        for fi107 := 0; fi107 < 1; fi107++ {
                            delete(v2, (((("a" + ("" + "\n\t")) + ("hello" + "a")) + ("" + "\n\t")) + (("\n\t" + "\n\t") + "a")))
                            _ = fi107
                        }
                        for fi108 := 0; fi108 < 14; fi108++ {
                            v3 = (0.1 * (v3 * v3))
                            v109 := GenS4{F0: 329, F1: 1, F2: v92}
                            _ = fi108
                            _ = v109
                        }
                    }
                    var v110 float64 = 0.0
                    _ = v110
                }
                _ = fi105
            }
            v1 += v1
            _ = v104
        }
        var v111 float64 = 3.14159
        for fi112 := 0; fi112 < 14; fi112++ {
            v3 += v3
            v2["hello"] = 826
            var v113 string = ""
            _ = fi112
            _ = v113
        }
        _ = v111
    }
    for fi114 := 0; fi114 < 14; fi114++ {
        var v115 int32 = (v92 * 2147483647)
        _ = fi114
        _ = v115
    }
    v90 += (1e10 - v90)
    switch "" {
    case "a":
        v4 = uint8(v89)
    case "hello":
        for fi116 := 0; fi116 < 5; fi116++ {
            var v117 int8 = v1
            _ = len(v2)
            _ = fi116
            _ = v117
        }
    default:
        v118 := []int{}
        _ = v118
    }
    for fi119 := 0; fi119 < 18; fi119++ {
        fmt.Println(v1, v5)
        _ = fi119
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v89
    _ = v90
    _ = v91
    _ = v92
}
