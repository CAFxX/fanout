package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 int64 }
func genFn0() uint64 {
    v4 := GenS1{F0: 1000, F1: 90, F2: -128}
    var v5 int64 = -947
    v6 := GenS2{F0: 1, F1: 1, F2: (("a" + "a") + "")}
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1() uint8 {
    switch "\n\t" {
    case "a":
        if true {
            if (3.14159 <= 1.5) {
                v7 := GenS2{F0: 127, F1: 492, F2: ""}
                _ = v7
            } else {
                if false {
                    var v8 uint64 = 515
                    _ = v8
                }
            }
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int32 = 2147483647
                var v11 float32 = 1e10
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    case "hello":
        for fi12 := 0; fi12 < 12; fi12++ {
            if (1 == 42) {
                switch fi12 {
                case -2147483648:
                    var v13 float32 = 0.0
                    _ = v13
                case 4294967295:
                    var v14 int16 = 32767
                    _ = v14
                case 128:
                    v15 := []int{(fi12 - (fi12 << ((18446744073709551615) & 63))), (fi12 ^ fi12)}
                    _ = v15
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    var v17 float32 = 3.14159
                    v17 -= float32((0.1 - -2.25))
                    _ = fi16
                    _ = v17
                }
                var v18 uint = 876
                _ = v18
            } else {
                v19 := GenS2{F0: 42, F1: 2, F2: ""}
                var v20 int = fi12
                _ = v19
                _ = v20
            }
            _ = fi12
        }
    case "foo bar":
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := []int{(fi21 % ((fi21) | 1)), fi21, fi21}
            var v23 uint16 = 42
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v24 := []int{-867, 386}
    if len(v24) > 0 {
        v24[0] = 127
    }
    var v25 int16 = 42
    _ = v24
    _ = v25
    return 1
}

func main() {
    var v1 int = -625
    var v2 uint16 = 1
    var v3 int = v1
    if (255 > 1) {
        if (!true) {
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := GenS3{F0: (v1 * -214), F1: ("hello" + "\n\t"), F2: -542}
                for fi28 := 0; fi28 < 3; fi28++ {
                    v2 += 780
                    var v29 uint16 = 409
                    fmt.Println(v2)
                    _ = fi28
                    _ = v29
                }
                v30 := []int{(v3 / (((fi26 >> ((255) & 63))) | 1))}
                _ = fi26
                _ = v27
                _ = v30
            }
            if (-105 >= 127) {
                v31 := GenS3{F0: -947, F1: "hello", F2: 1}
                _ = v31
            }
            switch "a" {
            case "foo bar":
                for fi32 := 0; fi32 < 11; fi32++ {
                    v33 := []int{}
                    _ = fi32
                    _ = v33
                }
            case "\n\t":
                fmt.Println(v3, v2)
            case "hello":
                switch v1 {
                case -2147483648:
                    switch "a" {
                    case "":
                        v1 -= (^(v1))
                    case "a":
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                v2 = v2
                            } else {
                                fmt.Println(v2)
                                v34 := GenS1{F0: 42, F1: 94, F2: -1}
                                _ = v34
                            }
                        case "a":
                            switch (v3 - v3) {
                            case 0:
                                switch v1 {
                                case 32768:
                                    v35 := []int{(v3 ^ v1), (v3 % ((v3) | 1)), 927, (v1 | v1)}
                                    _ = v35
                                case 259:
                                    v2 -= v2
                                case 32767:
                                    if ((1.5 / float64(0.1)) == -2.25) {
                                        v3 = (v3 % ((4294967295) | 1))
                                        var v36 uint8 = 2
                                        v37 := []int{-1, 783}
                                        _ = v36
                                        _ = v37
                                    }
                                default:
                                    fmt.Println(v1, v1, v1)
                                }
                            case 322:
                                for fi38 := 0; fi38 < 15; fi38++ {
                                    if (750 == 2) {
                                        var v39 int64 = -1
                                        var v40 uint = 1
                                        if (42 != v2) {
                                            v39 = 9223372036854775807
                                            if true {
                                                var v41 uint8 = uint8(v40)
                                                switch v3 {
                                                case 608:
                                                    v3 = v3
                                                case -579:
                                                    v41 += v41
                                                }
                                                if true {
                                                    for fi42 := 0; fi42 < 19; fi42++ {
                                                        if true {
                                                            v43 := []int{fi38, (v3 * int(v41)), v3, (v3 % (((v1 << (((v40 << (((v40 ^ (v40 + (v40 * (v40 - v40))))) & 63))) & 63))) | 1))}
                                                            _ = v43
                                                        } else {
                                                            v44 := []int{fi38, 1}
                                                            _ = v44
                                                        }
                                                        var v45 string = "\n\t"
                                                        fmt.Println(fi38, v45, v39)
                                                        if true {
                                                            if false {
                                                                v46 := []int{-2147483648, (fi42 >> ((v40) & 63))}
                                                                if (1.5 <= 1.5) {
                                                                    var v47 int32 = 1
                                                                    _ = v47
                                                                } else {
                                                                    var v48 int16 = -907
                                                                    switch 9223372036854775807 {
                                                                    case 32767:
                                                                        v49 := GenS2{F0: (v41 % (((v41 + v41)) | 1)), F1: (v2 | 42), F2: ((v45 + "hello") + v45)}
                                                                        _ = v49
                                                                    case 280:
                                                                        v45 = v45
                                                                    default:
                                                                        _ = len(v46)
                                                                    }
                                                                    var v50 int32 = -1
                                                                    if len(v46) >= 2 {
                                                                        v46 = v46[1:len(v46)]
                                                                    }
                                                                    v48 = v48
                                                                    _ = v48
                                                                    _ = v50
                                                                }
                                                                v2 -= 65535
                                                                if ((v41 * genFn1()) > uint8(v40)) {
                                                                    if len(v46) > 0 {
                                                                        v46[0] = (fi38 + (v3 / ((861) | 1)))
                                                                    }
                                                                }
                                                                _ = v46
                                                            } else {
                                                                if ((v45 + v45) <= "a") {
                                                                    switch 9223372036854775807 {
                                                                    case 32768:
                                                                        var v51 uint16 = (^(v2))
                                                                        _ = v51
                                                                    default:
                                                                        v2 += 278
                                                                    }
                                                                    var v52 int64 = v39
                                                                    v53 := []int{(fi38 & v3), fi38}
                                                                    var v54 float64 = -2.25
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    v55 := GenS3{F0: (^(v1)), F1: v45, F2: int16(v41)}
                                                                    _ = v55
                                                                }
                                                                var v56 int64 = v39
                                                                switch "" {
                                                                case "":
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v57 := []int{v1, fi38}
                                                                        _ = v57
                                                                    default:
                                                                        if (false || (!((false || (false || false)) || (v41 < v41)))) {
                                                                            v3 = fi42
                                                                        }
                                                                    }
                                                                default:
                                                                    fmt.Println(v41, v45, v1)
                                                                }
                                                                _ = v56
                                                            }
                                                        } else {
                                                            v39 += (-(v39))
                                                        }
                                                        v58 := GenS2{F0: (v41 * v41), F1: 65535, F2: "foo bar"}
                                                        _ = fi42
                                                        _ = v45
                                                        _ = v58
                                                    }
                                                }
                                                _ = v41
                                            } else {
                                                for fi59 := 0; fi59 < 3; fi59++ {
                                                    var v60 int32 = -1
                                                    if true {
                                                        switch 42 {
                                                        case 32767:
                                                            v2 = (^(v2))
                                                        case 9223372036854775807:
                                                            v61 := []int{127, fi59, int(v2)}
                                                            _ = v61
                                                        default:
                                                            v62 := GenS1{F0: 89, F1: 36, F2: 127}
                                                            _ = v62
                                                        }
                                                        var v63 float64 = float64(((1e10 * 1.5) * 1e10))
                                                        var v64 uint32 = 728
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v65 := GenS3{F0: (fi59 ^ v3), F1: "\n\t", F2: 396}
                                                        _ = v65
                                                    }
                                                    _ = fi59
                                                    _ = v60
                                                }
                                                var v66 int16 = -32768
                                                for fi67 := 0; fi67 < 10; fi67++ {
                                                    switch v1 {
                                                    case 0:
                                                        var v68 uint64 = 18446744073709551615
                                                        _ = v68
                                                    case 2147483647:
                                                        v69 := []int{-818}
                                                        _ = v69
                                                    case 128:
                                                        v70 := []int{(fi38 << ((2) & 63)), 32768}
                                                        _ = v70
                                                    }
                                                    if (((^(v40)) < v40) && false) {
                                                        switch 127 {
                                                        case -9223372036854775808:
                                                            v71 := []int{}
                                                            _ = v71
                                                        default:
                                                            var v72 string = "a"
                                                            _ = v72
                                                        }
                                                        v1 = -1
                                                        v73 := GenS1{F0: 752, F1: 100, F2: int8(v2)}
                                                        _ = v73
                                                    } else {
                                                        var v74 int16 = -1
                                                        var v75 int16 = (-(v66))
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    fmt.Println(fi38, fi38, v40)
                                                    _ = fi67
                                                }
                                                _ = v66
                                            }
                                            if ((-389 > 42) || true) {
                                                for fi76 := 0; fi76 < 2; fi76++ {
                                                    var v77 uint32 = 0
                                                    fmt.Println(v3, v2)
                                                    if (!(v39 != (v39 % ((v39) | 1)))) {
                                                        v78 := []int{(^(v3))}
                                                        v39 += v39
                                                        _ = v78
                                                    } else {
                                                        var v79 int64 = 42
                                                        _ = v79
                                                    }
                                                    _ = fi76
                                                    _ = v77
                                                }
                                                if (((false || (1.5 != (-2.25 * float64(1e10)))) && ("foo bar" >= "\n\t")) && true) {
                                                    v80 := GenS3{F0: v1, F1: "a", F2: 42}
                                                    fmt.Println(v39, fi38)
                                                    switch "hello" {
                                                    case "hello":
                                                        v81 := []int{v1, -9223372036854775808, v1}
                                                        _ = v81
                                                    case "foo bar":
                                                        v2 -= (v2 << (((v40 / ((297) | 1))) & 63))
                                                    case "":
                                                        v82 := GenS1{F0: 1000, F1: 0, F2: 127}
                                                        _ = v82
                                                    default:
                                                        var v83 bool = true
                                                        _ = v83
                                                    }
                                                    _ = v80
                                                }
                                            }
                                        } else {
                                            if true {
                                                v39 += 9223372036854775807
                                                for fi84 := 0; fi84 < 15; fi84++ {
                                                    v1 = (v3 + -9223372036854775808)
                                                    v3 += 9
                                                    _ = fi84
                                                }
                                                if false {
                                                    v85 := GenS3{F0: v1, F1: ((("" + "foo bar") + "hello") + (("a" + (("foo bar" + "foo bar") + "\n\t")) + "foo bar")), F2: 1000}
                                                    fmt.Println(v40)
                                                    var v86 uint64 = 2
                                                    _ = v85
                                                    _ = v86
                                                }
                                            }
                                            var v87 int32 = -564
                                            var v88 uint = (v40 + v40)
                                            _ = v87
                                            _ = v88
                                        }
                                        _ = v39
                                        _ = v40
                                    }
                                    for fi89 := 0; fi89 < 3; fi89++ {
                                        v3 = (fi38 * fi38)
                                        _ = fi89
                                    }
                                    _ = fi38
                                }
                            default:
                                var v90 uint16 = v2
                                _ = v90
                            }
                        }
                    default:
                        v91 := []int{}
                        _ = v91
                    }
                case 128:
                    switch (((("a" + "hello") + "") + "hello") + ("foo bar" + "hello")) {
                    case "a":
                        fmt.Println(v1, v1)
                    default:
                        v3 = v1
                    }
                default:
                    v92 := GenS1{F0: 42, F1: 100, F2: 89}
                    _ = v92
                }
            default:
                v1 = v1
            }
        }
    }
    v93 := []int{9223372036854775807, v1}
    switch (("\n\t" + (("\n\t" + (("\n\t" + "") + "\n\t")) + "hello")) + "") {
    case "\n\t":
        var v94 int64 = -9223372036854775808
        _ = v94
    case "hello":
        if false {
            if (0 >= 0) {
                var v95 uint32 = 2147483647
                for fi96 := 0; fi96 < 12; fi96++ {
                    var v97 uint32 = 2
                    _ = fi96
                    _ = v97
                }
                for fi98 := 0; fi98 < 6; fi98++ {
                    fmt.Println(v3)
                    _ = fi98
                }
                _ = v95
            }
        } else {
            fmt.Println(v1)
        }
    case "foo bar":
        var v99 uint8 = 127
        _ = v99
    }
    v3 += (v3 + 42)
    if (399 > 42) {
        fmt.Println(v2)
    } else {
        var v100 uint16 = (v2 / ((v2) | 1))
        var v101 uint64 = 1
        fmt.Println(v100)
        _ = v100
        _ = v101
    }
    _ = len(v93)
    var v102 int = (v1 >> ((4294967295) & 63))
    if len(v93) >= 2 {
        v93 = v93[1:len(v93)]
    }
    var v103 uint = 42
    var v104 float32 = 0.1
    for fi105 := 0; fi105 < 14; fi105++ {
        v3 += -42
        _ = fi105
    }
    var v106 uint = (v103 & 579)
    if true {
        v3 = 185
        if true {
            var v107 uint8 = 49
            var v108 uint32 = 2147483647
            _ = v107
            _ = v108
        } else {
            fmt.Println(v2, v104)
        }
        v109 := GenS4{F0: 0}
        _ = v109
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v106
    _ = v2
    _ = v3
    _ = v93
}
