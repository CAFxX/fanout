package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint64; F2 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 int64; F2 int8 }
type GenS8 struct { F0 uint; F1 int32; F2 int8 }
type GenS9 struct { F0 int64; F1 bool; F2 uint }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 int16; F1 uint64; F2 string }
type GenS12 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 uint32) int {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{}
        var v7 float32 = 0.0
        if true {
            var v8 float64 = -2.25
            for fi9 := 0; fi9 < 15; fi9++ {
                switch int(p1) {
                case 0:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    switch fi5 {
                    case -1:
                        p1 += 251
                    default:
                        if false {
                            if (-128 <= 127) {
                                v10 := GenS1{F0: 42}
                                var v11 bool = true
                                if false {
                                    v12 := GenS2{F0: 413, F1: p1}
                                    _ = v12
                                } else {
                                    if (2 != 777) {
                                        v8 = ((3.14159 * (v8 / v8)) * (1.5 / ((v8 * float64((1.5 + float32(v8)))) * 0.0)))
                                        var v13 string = "foo bar"
                                        _ = v13
                                    }
                                }
                                var v14 bool = ((42 == 0) || v11)
                                var v15 int16 = 32767
                                _ = v10
                                _ = v11
                                _ = v14
                                _ = v15
                            }
                            v7 = v7
                        }
                    }
                }
                _ = fi9
            }
            for fi16 := 0; fi16 < 4; fi16++ {
                v8 -= v8
                _ = fi16
            }
            _ = v8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if (610 > 0) {
            var v18 bool = true
            if (p0 == (p0 & p0)) {
                switch (^(fi17)) {
                case -477:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        if (1000 <= 815) {
                            var v20 float64 = (((3.14159 * 0.0) + 1e10) / -2.25)
                            switch "" {
                            case "hello":
                                if (2 > 65535) {
                                    if (!false) {
                                        v21 := GenS2{F0: 255, F1: 63}
                                        var v22 uint8 = 76
                                        _ = v21
                                        _ = v22
                                    } else {
                                        p0 = 2
                                        var v23 int64 = -1
                                        v24 := GenS3{F0: 42, F1: uint64(p1), F2: v18}
                                        var v25 string = ("a" + ("foo bar" + ("foo bar" + ("a" + (("foo bar" + "a") + ("hello" + ""))))))
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    fmt.Println(p1, p1)
                                    var v26 string = (((("" + "foo bar") + "") + "hello") + "hello")
                                    _ = v26
                                }
                            default:
                                if (42 >= 84) {
                                    switch (fi17 | (fi17 + (fi19 * 32767))) {
                                    case 1:
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v27 uint32 = p1
                                            _ = v27
                                        case "a":
                                            var v28 uint32 = p1
                                            _ = v28
                                        case "hello":
                                            p1 += p1
                                        default:
                                            var v29 uint16 = 0
                                            _ = v29
                                        }
                                    case 2147483647:
                                        var v30 uint = 0
                                        _ = v30
                                    default:
                                        p0 -= p0
                                    }
                                    v31 := GenS4{F0: p0, F1: (p0 / ((p0) | 1))}
                                    _ = v31
                                } else {
                                    if v18 {
                                        v32 := []int{fi17, fi19}
                                        switch 9223372036854775807 {
                                        case -812:
                                            var v33 float64 = v20
                                            _ = v33
                                        case 128:
                                            v20 = (v20 + ((v20 / float64((-2.25 / -2.25))) - 1e10))
                                        default:
                                            var v34 uint64 = p0
                                            _ = v34
                                        }
                                        _ = v32
                                    } else {
                                        v35 := GenS4{F0: (p0 | (p0 - 18446744073709551615)), F1: (p0 ^ (p0 << ((1) & 63)))}
                                        _ = v35
                                    }
                                    v36 := []int{1}
                                    switch ("" + "a") {
                                    case "a":
                                        v37 := GenS1{F0: -2147483648}
                                        _ = v37
                                    }
                                    _ = v36
                                }
                            }
                            if v18 {
                                fmt.Println(fi19, p0, p1)
                            }
                            _ = v20
                        } else {
                            if ((645 > uint(p1)) || v18) {
                                v38 := GenS5{F0: -533, F1: 5}
                                _ = v38
                            } else {
                                var v39 int8 = 42
                                var v40 uint32 = p1
                                if ((65535 <= 65535) && true) {
                                    fmt.Println(fi19, fi19, p0)
                                } else {
                                    p0 -= 2
                                    v41 := GenS6{F0: v18}
                                    v42 := []int{(fi19 % (((fi19 / (((fi19 << ((4294967295) & 63))) | 1))) | 1)), (fi17 - 4294967295)}
                                    _ = v41
                                    _ = v42
                                }
                                var v43 int16 = -51
                                _ = v39
                                _ = v40
                                _ = v43
                            }
                            v44 := []int{(fi19 % ((-9223372036854775808) | 1)), (fi19 & (fi17 * (-(fi19))))}
                            fmt.Println(p1, fi17, fi19)
                            _ = v44
                        }
                        v45 := GenS7{F0: (fi17 % ((-2147483648) | 1)), F1: -913, F2: 127}
                        switch "hello" {
                        case "\n\t":
                            v46 := []int{(-(fi19)), -2147483648, (fi17 + (fi19 % ((1) | 1))), -1}
                            _ = v46
                        case "foo bar":
                            p0 += p0
                        case "":
                            v47 := []int{(fi17 >> ((uint(p1)) & 63)), fi17, fi17}
                            _ = v47
                        }
                        _ = fi19
                        _ = v45
                    }
                default:
                    p0 = (p0 & p0)
                }
                if (((fi17 == (-(fi17))) && (v18 && v18)) || (-2.25 < (((1e10 * (1e10 + (1.5 * 3.14159))) * 1e10) + (3.14159 + (1e10 - 3.14159))))) {
                    for fi48 := 0; fi48 < 1; fi48++ {
                        if true {
                            p1 -= 319
                            p1 += p1
                        }
                        v49 := GenS8{F0: 65535, F1: -118, F2: -128}
                        _ = fi48
                        _ = v49
                    }
                    v50 := []int{}
                    _ = v50
                }
                v51 := GenS9{F0: -9223372036854775808, F1: v18, F2: uint(p0)}
                _ = v51
            } else {
                if v18 {
                    v52 := GenS5{F0: 273, F1: 0}
                    var v53 uint64 = (p0 + (p0 | (p0 * 42)))
                    _ = v52
                    _ = v53
                } else {
                    v54 := GenS10{F0: 1, F1: 100}
                    for fi55 := 0; fi55 < 16; fi55++ {
                        var v56 int32 = -2147483648
                        if (0 <= 0) {
                            p1 = 105
                            switch ("a" + "\n\t") {
                            case "foo bar":
                                var v57 uint = 512
                                _ = v57
                            case "\n\t":
                                var v58 uint64 = p0
                                _ = v58
                            case "a":
                                if (v18 && v18) {
                                    if (42 != 255) {
                                        switch fi17 {
                                        case 0:
                                            p0 = p0
                                        case -468:
                                            var v59 uint8 = 55
                                            _ = v59
                                        case 686:
                                            p0 += p0
                                        default:
                                            var v60 uint32 = (p1 + 2147483648)
                                            _ = v60
                                        }
                                        p0 -= (^(p0))
                                        v61 := []int{(fi17 * (fi55 ^ (-(fi55)))), 198}
                                        _ = v61
                                    } else {
                                        fmt.Println(fi17, fi55)
                                    }
                                    var v62 uint32 = p1
                                    var v63 uint8 = 209
                                    if true {
                                        var v64 int32 = int32(fi55)
                                        switch (("foo bar" + "a") + ("hello" + ("a" + ("foo bar" + "\n\t")))) {
                                        case "hello":
                                            if (v18 || (((-(fi17)) <= fi17) || (v18 || true))) {
                                                var v65 uint8 = (v63 / ((v63) | 1))
                                                var v66 uint16 = 42
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 uint8 = v63
                                                _ = v67
                                            }
                                        default:
                                            v68 := []int{fi55}
                                            _ = v68
                                        }
                                        if ((v62 - v62) == (p1 ^ (v62 + (p1 | (^(v62)))))) {
                                            switch (fi55 - fi17) {
                                            case 9223372036854775807:
                                                var v69 uint32 = p1
                                                _ = v69
                                            case 4294967295:
                                                v70 := []int{-874, -884, (fi17 ^ 32767), fi55}
                                                _ = v70
                                            case 32768:
                                                v71 := GenS7{F0: -2147483648, F1: 421, F2: 127}
                                                _ = v71
                                            default:
                                                v72 := GenS5{F0: v56, F1: 1}
                                                _ = v72
                                            }
                                            var v73 string = "foo bar"
                                            _ = v73
                                        } else {
                                            fmt.Println(fi17)
                                            var v74 bool = v18
                                            _ = v74
                                        }
                                        switch "foo bar" {
                                        case "":
                                            v75 := GenS1{F0: v56}
                                            _ = v75
                                        case "\n\t":
                                            v62 = (^(p1))
                                            v76 := GenS9{F0: -1, F1: (v18 || v18), F2: 579}
                                            _ = v76
                                        case "foo bar":
                                            v77 := GenS6{F0: (v18 && v18)}
                                            var v78 uint16 = 1
                                            switch ((("hello" + "foo bar") + "") + "foo bar") {
                                            case "a":
                                                if (false || (p0 != p0)) {
                                                    fmt.Println(v62, fi55, v56)
                                                    var v79 string = (("hello" + "hello") + "a")
                                                    var v80 bool = ((42 <= -83) && false)
                                                    _ = v79
                                                    _ = v80
                                                }
                                            case "foo bar":
                                                var v81 float64 = 0.1
                                                _ = v81
                                            case "\n\t":
                                                var v82 int8 = 26
                                                _ = v82
                                            }
                                            _ = v77
                                            _ = v78
                                        }
                                        _ = v64
                                    }
                                    _ = v62
                                    _ = v63
                                } else {
                                    fmt.Println(v18)
                                    v83 := []int{(fi17 % ((341) | 1))}
                                    _ = v83
                                }
                            }
                            var v84 bool = v18
                            _ = v84
                        } else {
                            v85 := []int{127, (fi17 - 9223372036854775807), 32767, int(p1)}
                            if (!(42 == 0)) {
                                if (65535 > 255) {
                                    if (v18 && v18) {
                                        if ((v18 || (1e10 > 1e10)) && v18) {
                                            var v86 string = "\n\t"
                                            _ = v86
                                        }
                                        switch -15 {
                                        case 28:
                                            if (!v18) {
                                                var v87 float32 = -2.25
                                                _ = v87
                                            }
                                        case 128:
                                            fmt.Println(p1)
                                        default:
                                            if v18 {
                                                v18 = true
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                var v88 uint64 = p0
                                                _ = v88
                                            case "a":
                                                var v89 bool = v18
                                                _ = v89
                                            case "":
                                                p1 += 443
                                            default:
                                                v56 = (v56 + 251)
                                            }
                                            v85 = append(v85, fi17)
                                            fmt.Println(p1)
                                        }
                                    } else {
                                        if v18 {
                                            p0 += p0
                                            if len(v85) >= 2 {
                                                v85 = v85[1:len(v85)]
                                            }
                                            p0 -= 457
                                        }
                                        v90 := GenS11{F0: 32767, F1: p0, F2: (("" + "\n\t") + ("hello" + ("" + "foo bar")))}
                                        _ = v90
                                    }
                                }
                                if (fi55 != -30) {
                                    p1 = (p1 >> ((2) & 63))
                                    var v91 float32 = float32((0.1 * 1e10))
                                    v18 = (846 > 2)
                                    _ = v91
                                }
                                fmt.Println(v18, v56, fi55)
                            }
                            v56 = v56
                            _ = v85
                        }
                        _ = fi55
                        _ = v56
                    }
                    var v92 float64 = (3.14159 * (-2.25 + 1.5))
                    _ = v54
                    _ = v92
                }
            }
            p1 = 2147483648
            _ = v18
        }
        switch "hello" {
        case "a":
            p1 = p1
        case "\n\t":
            fmt.Println(p0, p0)
        }
        _ = fi17
    }
    p1 += (^(p1))
    _ = p0
    _ = p1
    return 938
}

func main() {
    var v1 uint64 = 1
    var v2 uint64 = 603
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v93 int = 32767
    v93 -= v93
    if ((35 == 217) || (9223372036854775807 <= 0)) {
        var v94 int = 387
        _ = v94
    } else {
        if (("" > "hello") && (42 >= -2147483648)) {
            v1 += 1
            for fi95 := 0; fi95 < 14; fi95++ {
                v3["\n\t"] = fi95
                _ = fi95
            }
            var v96 uint32 = 65535
            _ = v96
        } else {
            for fi97 := 0; fi97 < 5; fi97++ {
                v93 -= v93
                for fi98 := 0; fi98 < 2; fi98++ {
                    v99 := GenS12{F0: 2, F1: 42, F2: 872}
                    _ = fi98
                    _ = v99
                }
                _ = fi97
            }
        }
        delete(v4, 2147483647)
    }
    v100 := []int{(v93 * -840), v93}
    v2 = v2
    var v101 float32 = float32(1e10)
    for fi102 := 0; fi102 < 13; fi102++ {
        v103 := v4[fi102]
        _ = fi102
        _ = v103
    }
    var v104 int = v93
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v101:", v101)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v93:", v93)
    _ = v1
    _ = v100
    _ = v101
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v93
}
