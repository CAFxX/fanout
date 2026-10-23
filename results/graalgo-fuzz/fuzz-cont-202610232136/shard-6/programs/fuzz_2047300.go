package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 string; F1 uint32; F2 uint32 }
type GenS10 struct { F0 bool; F1 bool; F2 int64 }
type GenS11 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 uint64; F1 uint64 }
func genFn0() int32 {
    v7 := []int{2147483647, 707, 407}
    if true {
        var v8 int = 128
        v8 += (v8 | 4294967295)
        _ = v8
    } else {
        var v9 int64 = -9223372036854775808
        var v10 int32 = 424
        var v11 int8 = 1
        _ = v10
        _ = v11
        _ = v9
    }
    _ = v7
    return 554
}
func genFn1(p0 int16, p1 int) int16 {
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            switch "a" {
            case "a":
                switch fi13 {
                case 2147483647:
                    switch (fi13 << ((0) & 63)) {
                    case -909:
                        switch 128 {
                        case -203:
                            fmt.Println(fi12, p1)
                        default:
                            switch "a" {
                            case "hello":
                                if (fi13 >= fi13) {
                                    fmt.Println(p0, fi13, fi13)
                                    v14 := GenS1{F0: "foo bar"}
                                    if true {
                                        switch "" {
                                        case "a":
                                            v15 := GenS2{F0: 255, F1: 2}
                                            _ = v15
                                        case "hello":
                                            var v16 int8 = -30
                                            v16 += v16
                                            var v17 int8 = (v16 / (((^(v16))) | 1))
                                            _ = v16
                                            _ = v17
                                        default:
                                            switch "\n\t" {
                                            case "foo bar":
                                                p1 = fi13
                                            case "\n\t":
                                                var v18 int8 = -106
                                                _ = v18
                                            }
                                        }
                                        var v19 int64 = 1
                                        v20 := GenS3{F0: 128}
                                        _ = v19
                                        _ = v20
                                    } else {
                                        if (!true) {
                                            v21 := []int{}
                                            v22 := GenS1{F0: ("a" + "a")}
                                            _ = v21
                                            _ = v22
                                        } else {
                                            if (0 <= int64(fi13)) {
                                                var v23 uint64 = 2
                                                _ = v23
                                            } else {
                                                p1 += (-(fi13))
                                                v24 := []int{fi12, p1, p1}
                                                p0 += p0
                                                _ = v24
                                            }
                                            var v25 uint = 65535
                                            v26 := []int{p1, (fi12 >> ((42) & 63)), fi13, fi12}
                                            _ = v25
                                            _ = v26
                                        }
                                        v27 := GenS4{F0: 42}
                                        _ = v27
                                    }
                                    _ = v14
                                } else {
                                    p0 = p0
                                }
                            }
                        }
                    case 127:
                        var v28 string = "foo bar"
                        _ = v28
                    case -820:
                        v29 := []int{}
                        var v30 string = ""
                        v29 = append(v29, (fi13 * fi13))
                        switch "\n\t" {
                        case "hello":
                            var v31 uint64 = 645
                            _ = v31
                        case "a":
                            var v32 uint8 = 92
                            _ = v32
                        case "":
                            p1 += fi13
                        default:
                            switch "" {
                            case "":
                                p1 -= 922
                            case "foo bar":
                                if ((fi12 + -578) != fi12) {
                                    v29 = append(v29, 0)
                                } else {
                                    fmt.Println(p1, p0)
                                }
                            case "a":
                                var v33 uint64 = 18446744073709551615
                                _ = v33
                            }
                        }
                        v34 := GenS2{F0: 2, F1: 65535}
                        _ = v29
                        _ = v30
                        _ = v34
                    }
                case -2147483648:
                    v35 := []int{}
                    _ = v35
                case -764:
                    fmt.Println(p1, p0)
                }
            case "\n\t":
                v36 := []int{fi13}
                _ = v36
            case "":
                p0 -= 920
            default:
                fmt.Println(p0)
            }
            var v37 float32 = -2.25
            v38 := GenS5{F0: true}
            v39 := []int{p1}
            var v40 float32 = ((v37 / v37) + 1e10)
            p1 = (^(fi13))
            v37 += v37
            _ = fi13
            _ = v37
            _ = v38
            _ = v39
            _ = v40
        }
        var v41 bool = (65535 > 0)
        _ = fi12
        _ = v41
    }
    if false {
        for fi42 := 0; fi42 < 11; fi42++ {
            for fi43 := 0; fi43 < 18; fi43++ {
                var v44 uint8 = 128
                _ = fi43
                _ = v44
            }
            p1 -= p1
            _ = fi42
        }
        p0 = p0
        p1 += p1
    }
    _ = p0
    _ = p1
    return (^(p0))
}

func main() {
    var v1 uint = 0
    var v2 int64 = -1
    var v3 int64 = v2
    var v4 int16 = 1
    v5 := make(map[string]int)
    v6 := make(map[uint32]bool)
    if true {
        for fi45 := 0; fi45 < 2; fi45++ {
            for fi46 := 0; fi46 < 19; fi46++ {
                var v47 bool = (float32(((3.14159 * (1e10 * (1.5 / (0.1 * (3.14159 + 0.1))))) * 0.1)) == 1e10)
                switch fi46 {
                case -907:
                    if (v3 == (v3 / ((-9223372036854775808) | 1))) {
                        if (uint64(v1) != 916) {
                            var v48 int8 = 110
                            var v49 uint32 = 4294967295
                            var v50 int = fi46
                            v50 -= (fi46 * (fi46 + 128))
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    } else {
                        v51 := GenS6{F0: (true || ((!false) && v47)), F1: -32, F2: 4294967295}
                        v1 = v1
                        _ = v51
                    }
                case 42:
                    var v52 int = (-(fi46))
                    _ = v52
                }
                var v53 int = -1
                if (v47 || v47) {
                    v2 += v3
                }
                _ = fi46
                _ = v47
                _ = v53
            }
            var v54 int = fi45
            _ = fi45
            _ = v54
        }
        v1 = (v1 ^ v1)
    } else {
        if false {
            var v55 int32 = -470
            _ = v55
        } else {
            for fi56 := 0; fi56 < 5; fi56++ {
                v57 := GenS2{F0: 802, F1: 18446744073709551615}
                _ = fi56
                _ = v57
            }
            v2 = -9223372036854775808
        }
        v58 := GenS2{F0: (v1 - (v1 - v1)), F1: v1}
        _ = v58
    }
    for fi59 := 0; fi59 < 16; fi59++ {
        var v60 uint8 = 79
        if (0 < 1) {
            if true {
                for fi61 := 0; fi61 < 11; fi61++ {
                    switch 0 {
                    case 128:
                        fmt.Println(v60)
                        v62 := GenS2{F0: (v1 ^ (v1 - (v1 & (v1 - v1)))), F1: 65535}
                        switch fi61 {
                        case 32767:
                            v63 := GenS5{F0: (-2.25 > 3.14159)}
                            _ = v63
                        case -2147483648:
                            v64 := v6[65535]
                            _ = v64
                        case 127:
                            fmt.Println(fi59, fi61, v60)
                        }
                        var v65 uint8 = 127
                        fmt.Println(v65, fi59, v65)
                        _ = v62
                        _ = v65
                    case -1:
                        v66 := []int{(fi59 >> (((v1 - v1)) & 63)), 2147483647}
                        _ = v66
                    default:
                        v67 := v6[227]
                        _ = v67
                    }
                    fmt.Println(v4)
                    _ = fi61
                }
                var v68 int64 = v3
                _ = v68
            } else {
                v69 := GenS7{F0: 42, F1: 425}
                _ = v69
            }
        }
        switch fi59 {
        case 711:
            v70 := v5["\n\t"]
            _ = v70
        case -645:
            var v71 int8 = 100
            _ = v71
        case -9223372036854775808:
            if (!false) {
                var v72 int16 = (v4 + v4)
                _ = v72
            } else {
                var v73 uint = (^(v1))
                v74 := GenS8{F0: (v3 | v2), F1: 1}
                _ = v73
                _ = v74
            }
        }
        _ = fi59
        _ = v60
    }
    var v75 string = "\n\t"
    v75 = v75
    v76 := GenS1{F0: (((((("" + v75) + (v75 + v75)) + "\n\t") + ("a" + v75)) + v75) + (((("\n\t" + (("foo bar" + v75) + v75)) + v75) + (("a" + "foo bar") + ((v75 + v75) + v75))) + (((v75 + "") + v75) + "foo bar")))}
    var v77 float32 = (1.5 * (0.0 / (((((0.0 / (0.1 * 0.1)) + 1.5) * (((-2.25 - (3.14159 * 1e10)) - (0.1 - float32(float64(3.14159)))) + ((0.0 * (0.0 / -2.25)) + 0.0))) * (1e10 * -2.25)) / 1.0)))
    v75 = ((("" + "") + (v75 + v75)) + "hello")
    var v78 uint16 = 65535
    for fi79 := 0; fi79 < 2; fi79++ {
        v80 := GenS9{F0: (((("a" + "foo bar") + v75) + (v75 + "")) + "foo bar"), F1: 2147483648, F2: 23}
        _ = fi79
        _ = v80
    }
    if (2 < 65535) {
        _ = len(v6)
        v81 := []int{}
        if ((177 >= 351) && false) {
            switch "a" {
            case "hello":
                if len(v81) > 0 {
                    v81[0] = -9223372036854775808
                }
            case "":
                v5[v75] = -996
            case "a":
                switch v75 {
                case "a":
                    v77 -= 1.5
                case "hello":
                    switch (v75 + v75) {
                    case "\n\t":
                        v82 := v6[42]
                        _ = v82
                    case "":
                        fmt.Println(v3, v4)
                    case "hello":
                        for fi83 := 0; fi83 < 1; fi83++ {
                            v84 := GenS10{F0: true, F1: false, F2: (-(v3))}
                            _ = fi83
                            _ = v84
                        }
                    }
                case "\n\t":
                    for fi85 := 0; fi85 < 12; fi85++ {
                        switch v75 {
                        case "foo bar":
                            v78 += (v78 ^ 14)
                        case "":
                            v86 := GenS11{F0: (v4 >> ((v1) & 63)), F1: 65535, F2: 106}
                            _ = v86
                        default:
                            var v87 string = (v75 + ((v75 + v75) + v75))
                            _ = v87
                        }
                        if (-138 != genFn0()) {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v3)
                            case "hello":
                                delete(v5, ("foo bar" + ("hello" + (((((("a" + (v75 + ("" + ("" + ((v75 + v75) + (v75 + ("a" + (v75 + v75)))))))) + v75) + "hello") + "hello") + "a") + v75))))
                            case "foo bar":
                                for fi88 := 0; fi88 < 1; fi88++ {
                                    v4 += (v4 | genFn1(v4, 128))
                                    v89 := GenS9{F0: (("a" + ("hello" + v75)) + ((((v75 + v75) + "a") + "") + v75)), F1: 42, F2: 0}
                                    delete(v6, 65535)
                                    _ = fi88
                                    _ = v89
                                }
                            }
                            var v90 uint64 = 18446744073709551615
                            _ = v90
                        } else {
                            v91 := GenS12{F0: true}
                            if (42 > -1) {
                                v77 += float32(1e10)
                            }
                            _ = v91
                        }
                        if (1 < 2147483648) {
                            for fi92 := 0; fi92 < 17; fi92++ {
                                v81 = append(v81, int(v2))
                                _ = fi92
                            }
                        }
                        _ = fi85
                    }
                }
            }
        } else {
            v77 = v77
            if len(v81) >= 2 {
                v81 = v81[1:len(v81)]
            }
            var v93 uint8 = 235
            _ = v93
        }
        _ = v81
    } else {
        v4 -= 1000
        for fi94 := 0; fi94 < 19; fi94++ {
            v95 := []int{-290, fi94, fi94}
            for fi96 := 0; fi96 < 20; fi96++ {
                switch v75 {
                case "":
                    v97 := GenS13{F0: "hello", F1: 209}
                    _ = v97
                case "a":
                    v98 := GenS4{F0: v2}
                    _ = v98
                default:
                    delete(v5, (("" + "foo bar") + v75))
                }
                if (1e10 >= (3.14159 / 1.0)) {
                    _ = len(v5)
                    v3 = (v3 / (((^(v2))) | 1))
                    switch (fi94 - 1) {
                    case -367:
                        var v99 int32 = -217
                        _ = v99
                    case 127:
                        var v100 uint32 = 519
                        v101 := GenS1{F0: v75}
                        var v102 uint8 = 255
                        _ = v100
                        _ = v101
                        _ = v102
                    default:
                        fmt.Println(v2, v77, v78)
                    }
                }
                v5[""] = (-(fi94))
                _ = fi96
            }
            v6[574] = false
            _ = fi94
            _ = v95
        }
        for fi103 := 0; fi103 < 10; fi103++ {
            var v104 int = 495
            var v105 int32 = int32(fi103)
            v106 := GenS14{F0: 106, F1: 42}
            _ = fi103
            _ = v104
            _ = v105
            _ = v106
        }
    }
    v107 := v5[(v75 + "hello")]
    v77 = float32(0.0)
    v108 := v6[uint32(v1)]
    _ = len(v5)
    var v109 uint64 = 347
    var v110 uint16 = v78
    v2 -= v2
    v108 = v108
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
