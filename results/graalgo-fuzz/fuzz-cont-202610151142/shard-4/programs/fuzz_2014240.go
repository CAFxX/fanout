package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 string }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int16 }
func genFn0(p0 int8, p1 uint, p2 int16) uint64 {
    if (("hello" + "a") >= ("\n\t" + "\n\t")) {
        p0 = p0
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 42}
            var v8 int32 = 42
            fmt.Println(p0, p2, p2)
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        var v9 bool = (false && ((1 > 418) && ((false || true) && (696 < 0))))
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS2{F0: 127, F1: 424, F2: "a"}
            var v12 uint8 = 255
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            if (!(1 < 0)) {
                if v9 {
                    v14 := GenS1{F0: 2147483647}
                    var v15 float64 = -2.25
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        p2 = (p2 % ((int16(p0)) | 1))
                        p1 = (p1 >> ((p1) & 63))
                        fmt.Println(fi16)
                        switch "foo bar" {
                        case "foo bar":
                            switch "\n\t" {
                            case "foo bar":
                                p1 -= p1
                                if ((p2 % ((1000) | 1)) == p2) {
                                    fmt.Println(p2, fi13, p0)
                                    var v17 string = ""
                                    v18 := []int{}
                                    _ = v17
                                    _ = v18
                                }
                            case "a":
                                var v19 int16 = 0
                                _ = v19
                            case "":
                                var v20 float64 = (0.1 / (float64(0.1) + 1.5))
                                var v21 string = "a"
                                v9 = (v9 || false)
                                _ = v20
                                _ = v21
                            default:
                                p2 -= p2
                            }
                        case "a":
                            var v22 uint = p1
                            _ = v22
                        case "hello":
                            fmt.Println(fi13, v9, p1)
                        }
                        _ = fi16
                    }
                }
                if false {
                    p2 = (p2 * (-(p2)))
                    var v23 uint32 = 1
                    _ = v23
                }
            }
            _ = fi13
        }
        _ = v9
    }
    var v24 uint64 = 776
    if (1.5 < 1.5) {
        if (!((!false) || (-9223372036854775808 < -9223372036854775808))) {
            if false {
                var v25 float64 = (float64((1.5 * 0.1)) - (-2.25 * float64(0.1)))
                if true {
                    var v26 string = ""
                    _ = v26
                } else {
                    v27 := []int{128}
                    if ((1 > (p2 << ((677) & 63))) && ("hello" >= "\n\t")) {
                        if true {
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            var v28 int64 = -217
                            var v29 uint = p1
                            _ = v28
                            _ = v29
                        } else {
                            p1 = p1
                        }
                        v30 := GenS2{F0: 26, F1: -2147483648, F2: "\n\t"}
                        v25 = (1.5 + (v25 / v25))
                        _ = v30
                    }
                    _ = v27
                }
                _ = v25
            }
            if false {
                switch 480 {
                case 1:
                    for fi31 := 0; fi31 < 13; fi31++ {
                        if (1 <= 128) {
                            var v32 int16 = (p2 | (p2 / (((p2 - (p2 + -413))) | 1)))
                            _ = v32
                        } else {
                            v33 := []int{fi31, 32768, 0, (^(fi31))}
                            _ = v33
                        }
                        _ = fi31
                    }
                case 42:
                    for fi34 := 0; fi34 < 14; fi34++ {
                        v35 := []int{}
                        _ = fi34
                        _ = v35
                    }
                }
                v36 := []int{714, 4294967295}
                _ = v36
            }
            for fi37 := 0; fi37 < 10; fi37++ {
                if (fi37 <= -284) {
                    v38 := []int{fi37, -32, -1, (fi37 ^ -9223372036854775808)}
                    _ = v38
                }
                v39 := GenS2{F0: 2, F1: 1, F2: "\n\t"}
                _ = fi37
                _ = v39
            }
        }
    } else {
        v40 := GenS3{F0: -9223372036854775808, F1: 0}
        switch (("\n\t" + "a") + "foo bar") {
        case "a":
            p2 = -32768
        case "hello":
            var v41 int = 9223372036854775807
            _ = v41
        default:
            for fi42 := 0; fi42 < 15; fi42++ {
                var v43 uint32 = 2147483647
                var v44 uint = uint(v43)
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        var v45 bool = (!(!true))
        _ = v40
        _ = v45
    }
    var v46 uint = p1
    if true {
        for fi47 := 0; fi47 < 5; fi47++ {
            for fi48 := 0; fi48 < 2; fi48++ {
                p2 = (p2 << (((v46 % ((4294967295) | 1))) & 63))
                _ = fi48
            }
            fmt.Println(p0, p2)
            _ = fi47
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v46
    return (v24 - (v24 ^ v24))
}
func genFn1(p0 int32, p1 uint) int16 {
    fmt.Println(p1, p1)
    p1 = p1
    var v49 float32 = (0.0 * -2.25)
    v49 -= (v49 - ((v49 / (((-2.25 * (v49 - (0.1 / v49))) / ((((float32(1e10) * (v49 + v49)) * (0.0 + v49)) * float32(1e10)) / (v49 - 3.14159))) / (0.1 - (1.5 * v49)))) * v49))
    _ = p0
    _ = p1
    _ = v49
    return -842
}

func main() {
    var v1 float32 = (1.5 + float32((3.14159 / (float64(0.0) + (0.1 + 0.0)))))
    var v2 int8 = 9
    var v3 int16 = 1000
    var v4 uint64 = 445
    var v5 bool = false
    v50 := GenS3{F0: -669, F1: (v4 * genFn0((v2 ^ (^(v2))), 18446744073709551615, (v3 ^ genFn1(1, 643))))}
    v51 := []int{-9223372036854775808, -1, 32768}
    switch "hello" {
    case "hello":
        v51 = append(v51, 1)
    case "\n\t":
        v2 = v2
    case "a":
        if ((0.0 > -2.25) && (int64(v3) >= 42)) {
            if false {
                v1 -= -2.25
                for fi52 := 0; fi52 < 14; fi52++ {
                    switch ("hello" + "foo bar") {
                    case "hello":
                        for fi53 := 0; fi53 < 14; fi53++ {
                            if (!true) {
                                v51 = append(v51, (fi53 - 4294967295))
                                v1 = v1
                                if true {
                                    if (v5 || v5) {
                                        fmt.Println(fi53)
                                        if len(v51) > 0 {
                                            v51[0] = fi52
                                        }
                                        var v54 int32 = -1
                                        _ = v54
                                    }
                                    if (0.1 <= (-2.25 + 1e10)) {
                                        var v55 uint32 = uint32(v4)
                                        _ = v55
                                    }
                                } else {
                                    v1 -= (((-2.25 - v1) + v1) - v1)
                                }
                            } else {
                                v5 = (v4 != genFn0(v2, 0, 42))
                                _ = len(v51)
                                var v56 int16 = (v3 << ((255) & 63))
                                _ = v56
                            }
                            _ = fi53
                        }
                    default:
                        if v5 {
                            fmt.Println(v4, v5, v1)
                            v3 -= (^(v3))
                            v4 = v4
                        } else {
                            v5 = false
                            if v5 {
                                v57 := GenS4{F0: int64(v4), F1: 2}
                                if (1 >= 1) {
                                    for fi58 := 0; fi58 < 18; fi58++ {
                                        var v59 int16 = (v3 - genFn1(-1, 71))
                                        _ = fi58
                                        _ = v59
                                    }
                                    for fi60 := 0; fi60 < 6; fi60++ {
                                        v4 = (v4 | v4)
                                        _ = fi60
                                    }
                                }
                                _ = v57
                            } else {
                                for fi61 := 0; fi61 < 16; fi61++ {
                                    if (!v5) {
                                        v2 += 124
                                        v2 += -43
                                    } else {
                                        var v62 uint8 = 171
                                        if (!true) {
                                            v1 = 3.14159
                                            if v5 {
                                                if true {
                                                    v1 += float32(0.1)
                                                    var v63 int32 = int32(v4)
                                                    _ = v63
                                                }
                                            } else {
                                                v5 = v5
                                                v64 := GenS2{F0: v62, F1: (-(fi61)), F2: "foo bar"}
                                                v1 = -2.25
                                                _ = v64
                                            }
                                        } else {
                                            v2 -= -128
                                        }
                                        if ("\n\t" < "foo bar") {
                                            switch "hello" {
                                            case "\n\t":
                                                v51 = append(v51, (fi52 - (fi52 << ((635) & 63))))
                                            case "":
                                                var v65 uint16 = 2
                                                _ = v65
                                            case "a":
                                                v5 = (("" >= "foo bar") || v5)
                                            default:
                                                var v66 string = ("" + "")
                                                _ = v66
                                            }
                                            var v67 uint64 = v4
                                            fmt.Println(v3, v4)
                                            var v68 float64 = 0.0
                                            fmt.Println(v3)
                                            v51 = append(v51, fi61)
                                            _ = v67
                                            _ = v68
                                        } else {
                                            switch ("" + ("foo bar" + "\n\t")) {
                                            case "":
                                                var v69 float64 = 1.5
                                                _ = v69
                                            case "a":
                                                var v70 uint32 = uint32(v4)
                                                _ = v70
                                            case "hello":
                                                var v71 uint16 = 0
                                                _ = v71
                                            default:
                                                v1 = (v1 + v1)
                                            }
                                            _ = len(v51)
                                            v1 += v1
                                        }
                                        if (!false) {
                                            var v72 int32 = int32(v4)
                                            v72 = (^(v72))
                                            if v5 {
                                                if len(v51) > 0 {
                                                    v51[0] = -684
                                                }
                                                v5 = (fi52 < fi61)
                                                switch (("a" + "a") + "\n\t") {
                                                case "foo bar":
                                                    _ = len(v51)
                                                    v4 += v4
                                                default:
                                                    if v5 {
                                                        v2 = (v2 % ((100) | 1))
                                                        v73 := GenS5{F0: v3, F1: fi52}
                                                        v2 = (v2 ^ -40)
                                                        v4 = v4
                                                        _ = len(v51)
                                                        _ = v73
                                                    } else {
                                                        switch fi61 {
                                                        case 127:
                                                            if v5 {
                                                                var v74 int32 = -121
                                                                _ = v74
                                                            }
                                                        }
                                                    }
                                                    v3 += int16(fi61)
                                                    v72 = (v72 - -1)
                                                }
                                            } else {
                                                v2 += -4
                                            }
                                            _ = v72
                                        }
                                        _ = v62
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        v1 = v1
                                    case "foo bar":
                                        fmt.Println(v4)
                                    }
                                    var v75 float32 = v1
                                    _ = fi61
                                    _ = v75
                                }
                                var v76 uint64 = v4
                                v5 = true
                                _ = v76
                            }
                        }
                    }
                    v3 = (v3 & (v3 ^ genFn1(-2147483648, 0)))
                    v2 -= (-(v2))
                    _ = fi52
                }
            }
        } else {
            v77 := GenS1{F0: 42}
            v2 = v2
            _ = v77
        }
    }
    var v78 uint16 = 0
    if (2147483647 != 2) {
        v79 := GenS3{F0: -1, F1: (^(v4))}
        var v80 int = 32767
        _ = v79
        _ = v80
    }
    fmt.Println(v1)
    var v81 uint16 = v78
    var v82 float64 = 1.5
    v83 := GenS2{F0: 29, F1: 916, F2: ("foo bar" + (("foo bar" + "foo bar") + "hello"))}
    v3 -= -497
    for fi84 := 0; fi84 < 3; fi84++ {
        for fi85 := 0; fi85 < 3; fi85++ {
            var v86 uint = 270
            fmt.Println(v81, fi84)
            var v87 int8 = (^(v2))
            _ = fi85
            _ = v86
            _ = v87
        }
        switch (^(fi84)) {
        case 1:
            v3 = (-(v3))
        case 4294967295:
            v81 = (v81 + (v78 ^ (v78 | 732)))
        case -881:
            if v5 {
                if (255 > uint8(v81)) {
                    v88 := GenS2{F0: 255, F1: fi84, F2: "a"}
                    if false {
                        v2 = v2
                    }
                    _ = v88
                } else {
                    v89 := GenS6{F0: v3}
                    v90 := GenS3{F0: 42, F1: (v4 >> ((395) & 63))}
                    _ = v89
                    _ = v90
                }
                v82 -= (float64(3.14159) * v82)
            } else {
                if (-2147483648 <= 1) {
                    fmt.Println(v3)
                    for fi91 := 0; fi91 < 5; fi91++ {
                        if v5 {
                            _ = len(v51)
                        }
                        var v92 int8 = (v2 - v2)
                        var v93 float64 = float64((-2.25 / (v1 - (v1 * ((((1.5 / 1.0) + v1) * float32(v82)) - (v1 + (-2.25 / (v1 + float32(0.0)))))))))
                        _ = fi91
                        _ = v92
                        _ = v93
                    }
                    v82 = v82
                }
            }
        default:
            if v5 {
                for fi94 := 0; fi94 < 5; fi94++ {
                    var v95 int32 = 0
                    v1 -= (v1 / float32(v82))
                    _ = fi94
                    _ = v95
                }
                if v5 {
                    v96 := GenS5{F0: v3, F1: (^(fi84))}
                    _ = v96
                }
                for fi97 := 0; fi97 < 19; fi97++ {
                    v98 := GenS4{F0: -227, F1: (v78 + 336)}
                    var v99 int64 = -762
                    _ = fi97
                    _ = v98
                    _ = v99
                }
            } else {
                v100 := GenS4{F0: 9223372036854775807, F1: v81}
                _ = v100
            }
        }
        _ = fi84
    }
    var v101 float64 = ((0.1 - v82) * v82)
    v1 += 1.5
    v102 := GenS3{F0: 0, F1: (v4 % (((v4 << ((734) & 63))) | 1))}
    var v103 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v78
    _ = v81
    _ = v82
    _ = v83
}
