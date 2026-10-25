package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64; F1 string; F2 int64 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int8; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0(p0 uint64, p1 int16, p2 uint16) string {
    v7 := []int{}
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 2147483647, F1: p1, F2: fi8}
        v10 := GenS1{F0: 2, F1: (p1 >> ((42) & 63)), F2: -965}
        switch fi8 {
        case 408:
            if ("hello" <= (("hello" + "\n\t") + "hello")) {
                p2 += 976
                var v11 uint32 = 4294967295
                var v12 int32 = 1
                _ = v11
                _ = v12
            }
        case 128:
            if ("\n\t" >= "foo bar") {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 int = (^(fi13))
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        if false {
                            var v15 float64 = ((-2.25 / (0.0 - 3.14159)) * ((0.1 * 1e10) + float64(1.5)))
                            _ = v15
                        } else {
                            v14 = (fi8 & (v14 << ((4294967295) & 63)))
                        }
                    case "hello":
                        if true {
                            fmt.Println(fi8)
                        }
                    default:
                        var v16 uint8 = 119
                        _ = v16
                    }
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v14 += (v14 / (((fi8 / ((v14) | 1))) | 1))
                    fmt.Println(p2)
                    _ = fi13
                    _ = v14
                }
                fmt.Println(p2, fi8)
            }
        default:
            p2 = (p2 ^ p2)
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 4294967295 {
    case 4294967295:
        p2 -= 0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("a" + ("\n\t" + ("\n\t" + ("foo bar" + ""))))
}
func genFn1(p0 uint) bool {
    for fi17 := 0; fi17 < 1; fi17++ {
        if false {
            p0 = 65535
            p0 -= p0
        } else {
            fmt.Println(fi17, fi17)
            if (53 == 244) {
                var v18 string = ("foo bar" + "")
                _ = v18
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    switch (fi19 * fi19) {
                    case 32768:
                        p0 = 73
                    }
                    _ = fi19
                }
                v20 := []int{647}
                for fi21 := 0; fi21 < 14; fi21++ {
                    v20 = append(v20, fi21)
                    if true {
                        v22 := GenS2{F0: 32768, F1: 4294967295, F2: (4294967295 <= 42)}
                        _ = v22
                    } else {
                        switch (^(fi21)) {
                        case 32768:
                            var v23 int32 = -2147483648
                            _ = v23
                        case 128:
                            v24 := GenS1{F0: 0, F1: -1, F2: (fi21 & fi21)}
                            _ = v24
                        case 9223372036854775807:
                            if (!false) {
                                if len(v20) > 0 {
                                    v20[0] = 128
                                }
                                _ = len(v20)
                            }
                        default:
                            switch "" {
                            case "\n\t":
                                var v25 int8 = -47
                                _ = v25
                            }
                        }
                        p0 = (p0 - (^(p0)))
                        if (0 == 0) {
                            fmt.Println(p0, fi21)
                        } else {
                            v26 := GenS3{F0: 1, F1: ((("a" + ((("a" + "foo bar") + ("" + "\n\t")) + "hello")) + ("a" + "a")) + "foo bar"), F2: 42}
                            _ = v26
                        }
                    }
                    if false {
                        var v27 int8 = 100
                        if (("hello" + "") <= "foo bar") {
                            var v28 float32 = 0.0
                            var v29 int8 = (v27 & v27)
                            v27 = (v27 & -1)
                            v29 += (^(v27))
                            _ = v28
                            _ = v29
                        } else {
                            var v30 float32 = 1.5
                            _ = v30
                        }
                        _ = v27
                    } else {
                        v20 = append(v20, fi21)
                        var v31 int64 = 683
                        _ = v31
                    }
                    v32 := GenS2{F0: (fi17 + (fi17 - (^(fi17)))), F1: 0, F2: false}
                    p0 = (p0 << ((4294967295) & 63))
                    if ("hello" >= "foo bar") {
                        v33 := GenS1{F0: 215, F1: 681, F2: fi21}
                        _ = v33
                    }
                    p0 = p0
                    _ = fi21
                    _ = v32
                }
                _ = v20
            }
            if ((true || ((fi17 + (fi17 / ((fi17) | 1))) == (fi17 >> (((^(p0))) & 63)))) || false) {
                p0 = p0
                var v34 float32 = 1.5
                p0 -= (p0 ^ (^(p0)))
                _ = v34
            } else {
                var v35 uint16 = 1
                switch ("" + ("a" + "hello")) {
                case "":
                    fmt.Println(fi17, p0)
                }
                v36 := []int{}
                _ = v35
                _ = v36
            }
        }
        _ = fi17
    }
    if ((1.5 + -2.25) < 0.1) {
        if ("" < ("" + ("a" + "\n\t"))) {
            p0 = p0
            for fi37 := 0; fi37 < 2; fi37++ {
                v38 := GenS3{F0: 511, F1: ("hello" + ""), F2: -9223372036854775808}
                if false {
                    var v39 int64 = 42
                    v39 = (v39 - v39)
                    _ = v39
                } else {
                    v40 := GenS4{F0: 0, F1: -1}
                    _ = v40
                }
                _ = fi37
                _ = v38
            }
        } else {
            var v41 uint64 = 2
            _ = v41
        }
    }
    p0 += p0
    for fi42 := 0; fi42 < 7; fi42++ {
        for fi43 := 0; fi43 < 17; fi43++ {
            v44 := []int{fi43, (fi43 << ((p0) & 63)), (fi43 / ((2147483647) | 1))}
            _ = fi43
            _ = v44
        }
        switch fi42 {
        case 42:
            if true {
                fmt.Println(fi42, fi42)
                switch "" {
                case "\n\t":
                    for fi45 := 0; fi45 < 9; fi45++ {
                        var v46 int32 = -2147483648
                        v47 := GenS4{F0: 2, F1: 2147483647}
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                case "foo bar":
                    var v48 string = ""
                    _ = v48
                default:
                    var v49 int16 = 1000
                    _ = v49
                }
            } else {
                switch (fi42 % ((9223372036854775807) | 1)) {
                case 69:
                    v50 := GenS3{F0: 42, F1: "\n\t", F2: -908}
                    _ = v50
                case 2147483647:
                    p0 -= 255
                case 9223372036854775807:
                    p0 = (p0 << ((836) & 63))
                }
                var v51 int64 = 1
                _ = v51
            }
        case 127:
            var v52 string = "hello"
            _ = v52
        default:
            var v53 int32 = 768
            _ = v53
        }
        _ = fi42
    }
    var v54 float32 = ((1e10 - 1e10) / (1e10 + (1e10 + (1.5 * 0.0))))
    _ = p0
    _ = v54
    return true
}

func main() {
    var v1 uint8 = 255
    var v2 uint64 = uint64(v1)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 int = 128
    var v6 uint = 0
    v55 := GenS1{F0: 65535, F1: 327, F2: v5}
    v56 := GenS5{F0: -1, F1: v2, F2: 974}
    switch (v5 / ((v5) | 1)) {
    case -651:
        v57 := GenS2{F0: v5, F1: (v2 << (((v6 << ((42) & 63))) & 63)), F2: (false || false)}
        _ = v57
    case 128:
        v58 := GenS2{F0: v5, F1: v2, F2: (true || (-1 <= -9223372036854775808))}
        _ = v58
    case 0:
        for fi59 := 0; fi59 < 14; fi59++ {
            v60 := []int{fi59, (fi59 << (((v6 / ((711) | 1))) & 63)), fi59, v5}
            switch "hello" {
            case "foo bar":
                v1 = (v1 << (((v6 * 2)) & 63))
            case "a":
                v61 := GenS4{F0: 221, F1: -713}
                _ = v61
            default:
                for fi62 := 0; fi62 < 7; fi62++ {
                    var v63 int64 = 273
                    if false {
                        if (1e10 > 0.0) {
                            switch (fi62 << (((v6 / (((v6 & v6)) | 1))) & 63)) {
                            case -9223372036854775808:
                                var v64 bool = (32767 <= -32768)
                                _ = v64
                            case 2147483647:
                                var v65 uint = v6
                                _ = v65
                            }
                            switch ("foo bar" + "") {
                            case "\n\t":
                                v2 = 712
                                var v66 int8 = 0
                                _ = v66
                            case "hello":
                                v60 = append(v60, fi59)
                            }
                            if len(v60) > 0 {
                                v60[0] = fi62
                            }
                            v2 = 1
                            if (v63 < v63) {
                                if (((1e10 / -2.25) - 1.5) != 1.5) {
                                    v6 = (v6 >> ((uint(v1)) & 63))
                                    v4[2] = (v5 | int(v2))
                                    if (2147483648 <= 42) {
                                        if false {
                                            v2 -= (^(v2))
                                        }
                                        if true {
                                            v67 := GenS2{F0: fi59, F1: 42, F2: true}
                                            switch "a" {
                                            case "a":
                                                v5 += 1
                                            case "\n\t":
                                                v68 := GenS2{F0: int(v6), F1: 18446744073709551615, F2: ((fi62 == 827) && (true || true))}
                                                _ = v68
                                            default:
                                                var v69 uint = uint(v1)
                                                _ = v69
                                            }
                                            _ = v67
                                        } else {
                                            var v70 uint16 = 2
                                            _ = v70
                                        }
                                    } else {
                                        fmt.Println(v6)
                                    }
                                    var v71 int8 = 1
                                    fmt.Println(v1, v2, v2)
                                    v3[(v5 + (-(fi62)))] = false
                                    var v72 int = (^(v5))
                                    var v73 int64 = (v63 | (v63 | v63))
                                    v4[65535] = v5
                                    var v74 int = (fi59 * (v72 - -512))
                                    var v75 uint32 = 2147483648
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                } else {
                                    switch (fi62 << ((42) & 63)) {
                                    case 422:
                                        _ = len(v60)
                                    }
                                }
                            } else {
                                v60 = append(v60, (v5 + (-(fi62))))
                            }
                        } else {
                            var v76 int16 = 113
                            var v77 bool = (-518 != int32(v6))
                            var v78 int32 = 42
                            _ = v76
                            _ = v77
                            _ = v78
                        }
                        var v79 uint8 = v1
                        v79 += 34
                        v80 := GenS6{F0: (v6 >> (((v6 & 42)) & 63)), F1: 65535, F2: v63}
                        v81 := GenS1{F0: 2147483647, F1: 1000, F2: v5}
                        _ = v79
                        _ = v80
                        _ = v81
                    } else {
                        var v82 bool = true
                        if len(v60) >= 2 {
                            v60 = v60[1:len(v60)]
                        }
                        _ = v82
                    }
                    if true {
                        var v83 int64 = (v63 << (((v6 - v6)) & 63))
                        var v84 uint8 = 42
                        _ = v83
                        _ = v84
                    }
                    _ = fi62
                    _ = v63
                }
            }
            var v85 uint16 = 230
            _ = fi59
            _ = v60
            _ = v85
        }
    }
    fmt.Println(v2)
    v86 := v3[(v5 >> (((v6 + v6)) & 63))]
    var v87 bool = (!false)
    var v88 string = "a"
    var v89 int16 = -797
    for fi90 := 0; fi90 < 3; fi90++ {
        v91 := []int{}
        _ = fi90
        _ = v91
    }
    if (-32768 < -1) {
        var v92 int = (v5 / (((v5 - 34)) | 1))
        fmt.Println(v2, v86, v89)
        switch "" {
        case "foo bar":
            v93 := GenS1{F0: uint32(v2), F1: v89, F2: 42}
            _ = v93
        case "hello":
            if true {
                var v94 int = 128
                _ = v94
            }
        }
        _ = v92
    }
    v5 = 4294967295
    fmt.Println(v89, v6)
    v1 -= v1
    v1 -= v1
    v1 -= v1
    v95 := []int{760, (v5 % ((32768) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("slicelen v95:", len(v95), v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v6
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v95
}
