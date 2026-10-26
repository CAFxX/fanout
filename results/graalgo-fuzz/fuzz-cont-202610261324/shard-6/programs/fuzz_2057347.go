package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint32; F1 string; F2 int }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 int64; F2 uint16 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int32 }
func genFn0() int {
    var v6 uint = 255
    fmt.Println(v6)
    v7 := GenS1{F0: 42}
    _ = v6
    _ = v7
    return -95
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = 1e10
    var v3 bool = (127 == 448)
    var v4 uint8 = 183
    v5 := []uint32{}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint8 = (v4 | (^(v4)))
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v3 = v3
        fmt.Println(v2, v1)
        if (false || true) {
            v11 := GenS1{F0: 42}
            v2 = v2
            if (genFn0() < genFn0()) {
                var v12 uint64 = (v1 + 824)
                v2 += (v2 + v2)
                if v3 {
                    for fi13 := 0; fi13 < 6; fi13++ {
                        var v14 int16 = 1000
                        v15 := GenS1{F0: 738}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    fmt.Println(v3, v2, fi10)
                }
                _ = v12
            } else {
                v4 = (v4 - 128)
            }
            _ = v11
        }
        _ = fi10
    }
    if ((((v4 % ((120) | 1)) >= (^(v4))) || false) && true) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            v1 -= v1
            switch fi16 {
            case 9223372036854775807:
                v4 += (v4 % ((v4) | 1))
            }
            _ = fi16
        }
        v5 = append(v5, 65535)
    } else {
        if ((32768 != -834) || false) {
            v4 = (v4 | (v4 & 204))
        } else {
            v3 = ("foo bar" <= ("hello" + "hello"))
            fmt.Println(v2, v1)
            v3 = v3
        }
        v2 = 0.1
        for fi17 := 0; fi17 < 11; fi17++ {
            switch (fi17 * (fi17 >> ((18446744073709551615) & 63))) {
            case 4294967295:
                fmt.Println(v4)
            case 0:
                v18 := GenS1{F0: 1}
                _ = v18
            case -2147483648:
                var v19 string = "hello"
                _ = v19
            }
            _ = fi17
        }
    }
    switch 32768 {
    case 1:
        v3 = v3
    case 321:
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS1{F0: 32767}
            v1 = (v1 % ((v1) | 1))
            _ = fi20
            _ = v21
        }
    default:
        fmt.Println(v2, v4, v4)
    }
    if v3 {
        var v22 int16 = 42
        _ = v22
    } else {
        if v3 {
            v23 := GenS1{F0: 207}
            _ = v23
        } else {
            if ((v3 && (-948 != 418)) || v3) {
                var v24 uint16 = 65535
                v2 -= v2
                v25 := GenS2{F0: 0, F1: int(v4), F2: uint32(v4)}
                _ = v24
                _ = v25
            }
            _ = len(v5)
            var v26 uint = 18446744073709551615
            _ = v26
        }
        switch 128 {
        case 2147483647:
            v1 += 2
        case 32768:
            fmt.Println(v4, v3)
        case -1:
            switch "foo bar" {
            case "hello":
                var v27 int16 = 1
                _ = v27
            case "\n\t":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "":
                for fi28 := 0; fi28 < 9; fi28++ {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if v3 {
                            if (!v3) {
                                var v30 uint64 = v1
                                _ = v30
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = 2147483647
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if (!false) {
                                    var v31 bool = true
                                    _ = v31
                                }
                                _ = len(v5)
                                if true {
                                    var v32 int16 = 0
                                    _ = v32
                                } else {
                                    v4 += (^(v4))
                                    switch ("foo bar" + "hello") {
                                    case "":
                                        v3 = v3
                                    case "a":
                                        fmt.Println(fi28, v2)
                                    }
                                    v33 := GenS1{F0: 2}
                                    _ = v33
                                }
                                v5 = append(v5, 1)
                            }
                            v3 = true
                            if v3 {
                                v34 := GenS1{F0: 985}
                                if (!((v1 >> ((uint(v4)) & 63)) >= (^(v1)))) {
                                    var v35 int = (fi28 ^ (^(fi28)))
                                    v1 -= 1
                                    if v3 {
                                        v36 := GenS3{F0: 0}
                                        fmt.Println(v4)
                                        switch "a" {
                                        case "foo bar":
                                            var v37 bool = (v1 != v1)
                                            var v38 uint32 = 983
                                            _ = v37
                                            _ = v38
                                        case "\n\t":
                                            if v3 {
                                                v39 := GenS4{F0: 43, F1: v3}
                                                var v40 int = fi29
                                                v40 = fi28
                                                var v41 uint8 = v4
                                                var v42 bool = (!((!(v3 || (false && (false && (v3 && false))))) && true))
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                        case "":
                                            v1 -= (v1 * (^(v1)))
                                        default:
                                            if (!(1e10 >= (3.14159 / float64(1e10)))) {
                                                v1 = (v1 % (((v1 + (v1 * (v1 + v1)))) | 1))
                                            }
                                        }
                                        if v3 {
                                            switch (fi29 - 128) {
                                            case 2147483647:
                                                fmt.Println(fi28, fi28)
                                            case 32767:
                                                if v3 {
                                                    if ((!(!true)) || (("foo bar" <= "\n\t") && (22 != int32(fi28)))) {
                                                        if (v3 && false) {
                                                            v43 := GenS1{F0: 32767}
                                                            var v44 uint8 = 116
                                                            switch "hello" {
                                                            case "":
                                                                v45 := GenS5{F0: 256, F1: 0}
                                                                _ = v45
                                                            default:
                                                                v3 = (!v3)
                                                            }
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            if v3 {
                                                                var v46 int16 = -869
                                                                fmt.Println(v1)
                                                                if v3 {
                                                                    switch v35 {
                                                                    case -18:
                                                                        v2 += -2.25
                                                                    case 32767:
                                                                        v47 := GenS6{F0: 2147483647, F1: ("foo bar" + "a"), F2: (^(fi28))}
                                                                        _ = v47
                                                                    }
                                                                    if false {
                                                                        var v48 uint16 = 506
                                                                        fmt.Println(v4, v46)
                                                                        var v49 int32 = 7
                                                                        var v50 int8 = 98
                                                                        _ = v48
                                                                        _ = v49
                                                                        _ = v50
                                                                    } else {
                                                                        v3 = v3
                                                                        var v51 int64 = 0
                                                                        v2 += 0.0
                                                                        _ = v51
                                                                    }
                                                                    switch (-(v35)) {
                                                                    case 1:
                                                                        _ = len(v5)
                                                                    default:
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 339
                                                                        }
                                                                    }
                                                                }
                                                                _ = v46
                                                            } else {
                                                                v3 = (984 != -32768)
                                                            }
                                                        }
                                                        var v52 int8 = 98
                                                        var v53 float64 = 1e10
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        if ((1 != 1) || false) {
                                                            v54 := GenS7{F0: (^(v1)), F1: fi28}
                                                            _ = v54
                                                        } else {
                                                            if (v4 >= (v4 << ((65535) & 63))) {
                                                                var v55 int16 = 1000
                                                                v3 = ((v3 || v3) || ((fi28 << ((42) & 63)) >= fi29))
                                                                switch "a" {
                                                                case "a":
                                                                    var v56 uint64 = (v1 - v1)
                                                                    _ = v56
                                                                case "\n\t":
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                                _ = len(v5)
                                                                _ = v55
                                                            } else {
                                                                v1 -= 2
                                                            }
                                                            var v57 int16 = -32768
                                                            v2 += ((1e10 * 3.14159) / 3.14159)
                                                            switch "" {
                                                            case "\n\t":
                                                                var v58 int = fi29
                                                                _ = v58
                                                            default:
                                                                if v3 {
                                                                    v59 := GenS7{F0: v1, F1: fi29}
                                                                    fmt.Println(v3)
                                                                    fmt.Println(v4, fi29, v4)
                                                                    _ = v59
                                                                } else {
                                                                    v5 = append(v5, 2)
                                                                    var v60 uint8 = (v4 / ((v4) | 1))
                                                                    var v61 int64 = 575
                                                                    switch 32768 {
                                                                    case -9223372036854775808:
                                                                        _ = len(v5)
                                                                    default:
                                                                        var v62 int8 = -1
                                                                        _ = v62
                                                                    }
                                                                    _ = v60
                                                                    _ = v61
                                                                }
                                                            }
                                                            _ = v57
                                                        }
                                                    }
                                                }
                                            case 677:
                                                v2 -= (v2 / 1.0)
                                            }
                                            v63 := GenS2{F0: 42, F1: (fi29 * (v35 + 32767)), F2: 4294967295}
                                            switch 937 {
                                            case -674:
                                                var v64 float32 = (-2.25 + 1e10)
                                                _ = v64
                                            case 9223372036854775807:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            _ = v63
                                        } else {
                                            var v65 bool = (true && (0 >= 42))
                                            _ = v65
                                        }
                                        v3 = (!true)
                                        _ = v36
                                    }
                                    _ = v35
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v66 uint = 749
                                    var v67 string = ""
                                    v3 = (42 >= 1)
                                    _ = v66
                                    _ = v67
                                }
                                _ = v34
                            } else {
                                switch ("\n\t" + "") {
                                case "a":
                                    v68 := GenS2{F0: 65535, F1: fi29, F2: 42}
                                    _ = v68
                                case "foo bar":
                                    v69 := GenS8{F0: (v1 >> ((4294967295) & 63)), F1: 1}
                                    _ = v69
                                default:
                                    var v70 float32 = 1.5
                                    _ = v70
                                }
                                switch "" {
                                case "a":
                                    switch (("hello" + ("\n\t" + (((("foo bar" + ("\n\t" + ((("" + "hello") + "hello") + ""))) + "foo bar") + ("a" + "foo bar")) + "foo bar"))) + "\n\t") {
                                    case "foo bar":
                                        v4 -= v4
                                        var v71 float32 = (v2 / float32((1.5 * -2.25)))
                                        _ = v71
                                    case "a":
                                        v2 -= (v2 + (0.1 - 0.0))
                                    case "":
                                        var v72 int = (fi29 - fi28)
                                        _ = v72
                                    }
                                case "\n\t":
                                    var v73 uint32 = 42
                                    _ = v73
                                case "":
                                    fmt.Println(fi29, v1)
                                default:
                                    switch ("foo bar" + "a") {
                                    case "a":
                                        var v74 float64 = 1e10
                                        _ = v74
                                    case "\n\t":
                                        var v75 string = ("hello" + "\n\t")
                                        _ = v75
                                    }
                                }
                                var v76 int64 = -1
                                v2 -= -2.25
                                var v77 uint64 = (v1 ^ v1)
                                _ = v76
                                _ = v77
                            }
                        }
                        v2 -= 0.1
                        v78 := GenS9{F0: -78}
                        _ = fi29
                        _ = v78
                    }
                    _ = fi28
                }
            default:
                var v79 bool = (((v2 * (0.0 - ((v2 / float32((0.1 - 1.5))) / (((0.1 * (((0.0 / 3.14159) * 1.5) - float32((0.1 * 0.1)))) + ((v2 / ((v2 - v2) / 1e10)) / (1.5 / v2))) / v2)))) - (1.5 / float32(((1e10 - (-2.25 / 3.14159)) * ((((0.1 - 1.5) / float64(v2)) - (0.0 * -2.25)) + (3.14159 / -2.25)))))) == 3.14159)
                _ = v79
            }
        }
    }
    v4 -= v4
    var v80 float64 = (-2.25 + (1e10 + ((1e10 + (float64(v2) / 1.5)) + 3.14159)))
    if len(v5) > 0 {
        v5[0] = 42
    }
    var v81 float64 = v80
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    for fi82 := 0; fi82 < 5; fi82++ {
        for fi83 := 0; fi83 < 3; fi83++ {
            if (v3 && (127 < -128)) {
                v80 += v81
                switch (fi83 / ((genFn0()) | 1)) {
                case 9223372036854775807:
                    var v84 uint16 = 0
                    _ = v84
                case -2147483648:
                    v5 = append(v5, 2)
                case -823:
                    var v85 int16 = 32767
                    _ = v85
                }
                fmt.Println(v1)
            } else {
                v4 = (v4 & (v4 + v4))
                v5 = append(v5, 988)
            }
            v86 := GenS1{F0: 0}
            _ = fi83
            _ = v86
        }
        _ = fi82
    }
    if (!(883 == 2147483647)) {
        for fi87 := 0; fi87 < 10; fi87++ {
            v1 = 89
            _ = fi87
        }
        v1 += v1
    }
    var v88 uint8 = (^(v4))
    if true {
        v89 := GenS10{F0: -1, F1: int64(v1), F2: 1}
        for fi90 := 0; fi90 < 15; fi90++ {
            var v91 int32 = 2147483647
            var v92 uint64 = v1
            v4 += 34
            _ = fi90
            _ = v91
            _ = v92
        }
        _ = v89
    } else {
        var v93 string = "\n\t"
        v4 = 42
        _ = v93
    }
    if (32768 >= 2147483647) {
        if (100 <= 1) {
            v80 = (0.1 - v80)
        } else {
            _ = len(v5)
            _ = len(v5)
        }
    }
    switch "a" {
    case "a":
        for fi94 := 0; fi94 < 8; fi94++ {
            if v3 {
                fmt.Println(v2)
            } else {
                v2 = float32(v80)
                for fi95 := 0; fi95 < 16; fi95++ {
                    switch "\n\t" {
                    case "\n\t":
                        switch (fi94 >> ((255) & 63)) {
                        case 311:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 948:
                            var v96 uint32 = 695
                            _ = v96
                        default:
                            if v3 {
                                var v97 int16 = -913
                                var v98 bool = false
                                v98 = (!(256 == 1))
                                v99 := GenS11{F0: 195, F1: 130}
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                        }
                    }
                    _ = fi95
                }
                v100 := GenS12{F0: 519}
                _ = v100
            }
            _ = fi94
        }
    case "hello":
        if (2 >= 65535) {
            var v101 int16 = 256
            v2 = -2.25
            _ = v101
        } else {
            var v102 uint = 4294967295
            if ("" > "") {
                v102 = uint(v1)
                var v103 float32 = v2
                _ = v103
            } else {
                if len(v5) > 0 {
                    v5[0] = 2
                }
                v104 := GenS13{F0: 542}
                _ = v104
            }
            if v3 {
                fmt.Println(v102, v2, v102)
                v81 = (v80 * ((v81 / v80) + (float64(1.5) + float64(v2))))
                v88 -= 2
            } else {
                for fi105 := 0; fi105 < 5; fi105++ {
                    v106 := GenS14{F0: 1}
                    v80 += v80
                    _ = fi105
                    _ = v106
                }
            }
            _ = v102
        }
    default:
        v4 = 153
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v88
}
