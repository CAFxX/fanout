package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 string; F1 uint64; F2 string }
type GenS3 struct { F0 int16; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint) int {
    var v5 int = 364
    if false {
        var v6 int64 = 0
        _ = v6
    }
    p0 += (p0 * (p0 * p0))
    _ = p0
    _ = v5
    return v5
}
func genFn1() string {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{(fi7 ^ genFn0(65535)), (fi7 ^ (fi7 ^ (fi7 + (fi7 << ((4294967295) & 63))))), (fi7 & fi7), (fi7 << ((2) & 63))}
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint8 = 24
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 bool = (4294967295 >= 246)
    var v12 int8 = 1
    _ = v11
    _ = v12
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 int = 0
    var v4 uint = 18446744073709551615
    var v13 float64 = 1e10
    var v14 string = ""
    var v15 uint8 = 44
    switch (v3 >> ((2) & 63)) {
    case 1:
        v4 += v4
    case 369:
        switch int(v4) {
        case 1:
            switch v14 {
            case "":
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 float32 = 0.1
                    v17 = v17
                    fmt.Println(v13)
                    _ = fi16
                    _ = v17
                }
            case "hello":
                var v18 int = genFn0(v4)
                _ = v18
            }
        case 0:
            v19 := []int{v3, (v3 << ((1) & 63)), (v3 << (((^(v4))) & 63))}
            _ = v19
        case 128:
            if false {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint16 = 2
                    if false {
                        var v22 uint64 = 4294967295
                        switch (v14 + v14) {
                        case "foo bar":
                            v2[v3] = (!true)
                        case "\n\t":
                            fmt.Println(v15, v13, v22)
                        }
                        for fi23 := 0; fi23 < 4; fi23++ {
                            var v24 uint = (v4 | 42)
                            var v25 uint16 = (v21 | (v21 | 1))
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    } else {
                        v1[0] = 954
                        if true {
                            if false {
                                v26 := GenS1{F0: 0, F1: (^(v15)), F2: 255}
                                switch (fi20 | genFn0(2)) {
                                case 1:
                                    v13 = v13
                                case -56:
                                    fmt.Println(v21)
                                default:
                                    v3 = (fi20 ^ (-(fi20)))
                                }
                                _ = v26
                            } else {
                                for fi27 := 0; fi27 < 6; fi27++ {
                                    v21 -= 472
                                    v28 := []int{fi20, fi20, (^(v3))}
                                    _ = fi27
                                    _ = v28
                                }
                                fmt.Println(v4)
                                v14 = (((("\n\t" + v14) + "a") + "a") + "foo bar")
                            }
                        } else {
                            v29 := v1[0]
                            _ = v29
                        }
                        v13 -= v13
                    }
                    fmt.Println(fi20, v15)
                    _ = fi20
                    _ = v21
                }
                if false {
                    var v30 float32 = (0.0 / (3.14159 / -2.25))
                    for fi31 := 0; fi31 < 7; fi31++ {
                        v14 = ("\n\t" + v14)
                        _ = fi31
                    }
                    v32 := []int{(v3 + (v3 * (v3 >> ((701) & 63)))), (-(v3)), v3}
                    _ = v30
                    _ = v32
                } else {
                    var v33 int8 = -36
                    _ = v33
                }
            } else {
                switch (v14 + "a") {
                case "\n\t":
                    switch int(v15) {
                    case 878:
                        switch "a" {
                        case "":
                            switch ("" + "hello") {
                            case "foo bar":
                                if true {
                                    for fi34 := 0; fi34 < 20; fi34++ {
                                        if false {
                                            if (!false) {
                                                v3 = genFn0((v4 % (((v4 | 18446744073709551615)) | 1)))
                                            } else {
                                                if true {
                                                    switch ((v14 + v14) + v14) {
                                                    case "\n\t":
                                                        if (false && (!false)) {
                                                            var v35 string = (v14 + v14)
                                                            _ = v35
                                                        }
                                                    case "":
                                                        if true {
                                                            v15 += (v15 & v15)
                                                            v3 = v3
                                                        }
                                                    default:
                                                        if true {
                                                            var v36 int32 = -456
                                                            _ = v36
                                                        }
                                                    }
                                                } else {
                                                    v15 = v15
                                                }
                                            }
                                            for fi37 := 0; fi37 < 20; fi37++ {
                                                v3 = (fi37 * (fi37 ^ genFn0(uint(v15))))
                                                var v38 uint32 = 212
                                                _ = fi37
                                                _ = v38
                                            }
                                            if false {
                                                v2[32767] = (false || true)
                                                v3 -= (fi34 % ((v3) | 1))
                                                for fi39 := 0; fi39 < 5; fi39++ {
                                                    v14 = ""
                                                    var v40 float64 = (v13 * (v13 - v13))
                                                    _ = fi39
                                                    _ = v40
                                                }
                                            } else {
                                                var v41 float32 = (1.5 + 0.0)
                                                for fi42 := 0; fi42 < 13; fi42++ {
                                                    switch v3 {
                                                    case 4294967295:
                                                        fmt.Println(v13, v13)
                                                    default:
                                                        switch (fi42 + v3) {
                                                        case 2147483647:
                                                            if true {
                                                                v43 := GenS2{F0: ((v14 + "hello") + "hello"), F1: 4294967295, F2: ""}
                                                                v14 = "foo bar"
                                                                v44 := GenS1{F0: -1, F1: (^(v15)), F2: 282}
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                        default:
                                                            switch int(v15) {
                                                            case 127:
                                                                if false {
                                                                    var v45 uint16 = 255
                                                                    switch genFn0((v4 % ((v4) | 1))) {
                                                                    case 346:
                                                                        v41 -= (float32((v13 * float64(v41))) * (v41 / float32((3.14159 / v13))))
                                                                    case 350:
                                                                        var v46 uint64 = 2
                                                                        _ = v46
                                                                    case -1:
                                                                        var v47 int32 = -2147483648
                                                                        _ = v47
                                                                    }
                                                                    v45 = 65535
                                                                    _ = v45
                                                                } else {
                                                                    var v48 int = -617
                                                                    v48 += (fi42 / ((genFn0(1)) | 1))
                                                                    var v49 uint32 = 0
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                            case 1:
                                                                v50 := GenS2{F0: "\n\t", F1: 357, F2: ("foo bar" + v14)}
                                                                _ = v50
                                                            }
                                                        }
                                                    }
                                                    switch v14 {
                                                    case "a":
                                                        v51 := GenS3{F0: 585, F1: -1, F2: (v15 - (v15 >> ((v4) & 63)))}
                                                        _ = v51
                                                    case "foo bar":
                                                        fmt.Println(v14, v4)
                                                    }
                                                    if ((717 != int32(v4)) || true) {
                                                        var v52 int16 = 198
                                                        if ((v3 >= (fi42 & fi34)) || ((!("hello" <= v14)) && false)) {
                                                            var v53 float64 = (v13 - (v13 + v13))
                                                            v41 = v41
                                                            if true {
                                                                switch (v14 + v14) {
                                                                case "\n\t":
                                                                    fmt.Println(v53)
                                                                case "hello":
                                                                    var v54 int8 = -128
                                                                    _ = v54
                                                                case "":
                                                                    v55 := GenS4{F0: v52, F1: 214, F2: 42}
                                                                    _ = v55
                                                                default:
                                                                    if true {
                                                                        fmt.Println(v41, v4, v15)
                                                                    }
                                                                }
                                                                v56 := []int{}
                                                                _ = v56
                                                            }
                                                            var v57 int = (fi42 >> (((v4 + (v4 & 591))) & 63))
                                                            _ = v53
                                                            _ = v57
                                                        }
                                                        v58 := v2[1]
                                                        _ = v52
                                                        _ = v58
                                                    } else {
                                                        var v59 string = ("" + v14)
                                                        _ = v59
                                                    }
                                                    _ = fi42
                                                }
                                                _ = v41
                                            }
                                        }
                                        v14 = v14
                                        _ = fi34
                                    }
                                    if (0 == 853) {
                                        var v60 uint32 = 0
                                        _ = v60
                                    } else {
                                        v61 := GenS4{F0: -268, F1: 65535, F2: 1}
                                        _ = v61
                                    }
                                    v4 = 2
                                }
                            case "":
                                for fi62 := 0; fi62 < 13; fi62++ {
                                    if (1.5 <= ((float64(float32(((v13 - (((((0.0 * v13) / 3.14159) + float64(3.14159)) / (((v13 - -2.25) / float64(1e10)) - v13)) / (v13 * (0.0 * v13)))) + 0.0))) / (1e10 / (v13 + v13))) / (3.14159 / v13))) {
                                        v13 -= (v13 - v13)
                                        v63 := GenS3{F0: 929, F1: 9223372036854775807, F2: (v15 << ((v4) & 63))}
                                        _ = v63
                                    }
                                    for fi64 := 0; fi64 < 13; fi64++ {
                                        v13 = (float64(((1e10 + (1.5 + 1.5)) / 1e10)) - (((((0.0 / v13) * float64((float32(v13) + 1e10))) / v13) + v13) - v13))
                                        _ = fi64
                                    }
                                    var v65 string = v14
                                    _ = fi62
                                    _ = v65
                                }
                            default:
                                v13 += 0.1
                            }
                        case "hello":
                            for fi66 := 0; fi66 < 19; fi66++ {
                                for fi67 := 0; fi67 < 15; fi67++ {
                                    v13 -= v13
                                    v68 := GenS5{F0: 2147483648, F1: ((!(!((^(v15)) <= v15))) && ((1 < -2147483648) && (!((true || true) || (true && false)))))}
                                    var v69 uint32 = 493
                                    _ = fi67
                                    _ = v68
                                    _ = v69
                                }
                                for fi70 := 0; fi70 < 15; fi70++ {
                                    var v71 int32 = -1
                                    v72 := GenS1{F0: -1, F1: (v15 ^ 208), F2: 458}
                                    v73 := GenS6{F0: 100}
                                    _ = fi70
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                                if ((float32((1e10 + -2.25)) + 0.1) != -2.25) {
                                    var v74 uint16 = 256
                                    v15 = (v15 + (^(v15)))
                                    _ = v74
                                } else {
                                    v75 := GenS7{F0: 531, F1: -542}
                                    v76 := []int{fi66}
                                    var v77 uint64 = 42
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                                _ = fi66
                            }
                        case "a":
                            v13 += v13
                        }
                    case 4294967295:
                        v3 += v3
                    case 128:
                        _ = len(v2)
                    default:
                        fmt.Println(v3, v13, v15)
                    }
                case "foo bar":
                    switch (v3 << (((v4 + 336)) & 63)) {
                    case 128:
                        var v78 int32 = 0
                        _ = v78
                    case -95:
                        fmt.Println(v3, v13)
                    case -1:
                        v15 = 35
                    }
                }
                var v79 string = (v14 + "\n\t")
                _ = v79
            }
        default:
            v80 := GenS6{F0: 1}
            _ = v80
        }
    }
    v81 := []int{(v3 >> ((7) & 63)), (^(v3)), 128}
    if false {
        var v82 uint = 1
        _ = v82
    } else {
        fmt.Println(v15, v4)
    }
    if (2147483647 != -991) {
        for fi83 := 0; fi83 < 20; fi83++ {
            var v84 uint64 = 230
            switch ((v14 + v14) + v14) {
            case "":
                delete(v2, genFn0(v4))
            case "hello":
                fmt.Println(v13)
            case "a":
                v85 := GenS3{F0: int16(v15), F1: 0, F2: (v15 % (((v15 ^ (v15 ^ 127))) | 1))}
                _ = v85
            }
            _ = fi83
            _ = v84
        }
        switch v14 {
        case "\n\t":
            if len(v81) > 0 {
                v81[0] = 4294967295
            }
        case "":
            delete(v2, v3)
        }
    }
    var v86 uint = v4
    var v87 uint32 = 2147483648
    v88 := GenS8{F0: 42}
    v87 -= v87
    switch (v14 + (("a" + ((v14 + v14) + "\n\t")) + v14)) {
    case "\n\t":
        v13 += (v13 - 0.1)
    default:
        v89 := GenS5{F0: (v87 << ((v86) & 63)), F1: true}
        _ = v89
    }
    if ((v4 + v86) < (v4 / (((v4 - v4)) | 1))) {
        for fi90 := 0; fi90 < 16; fi90++ {
            for fi91 := 0; fi91 < 17; fi91++ {
                var v92 bool = (0.0 != 0.1)
                if ((v14 + (v14 + v14)) == v14) {
                    fmt.Println(fi91)
                    var v93 bool = v92
                    if (v13 <= ((((v13 - v13) / v13) * float64((-2.25 / (float32((v13 - (((v13 - 1e10) + 0.0) / v13))) + 0.0)))) - ((((float64(1e10) + ((v13 + ((v13 / 0.1) / (1e10 * v13))) - v13)) + ((v13 * float64((((-2.25 + 1.5) * 3.14159) * (3.14159 - (((0.1 + (3.14159 * 3.14159)) / (float32(v13) - float32((v13 + (0.0 / (((((v13 / float64(((3.14159 / (-2.25 - 0.0)) / float32(float64(1e10))))) - v13) * v13) * (v13 / ((v13 - -2.25) / (v13 / v13)))) * float64(3.14159))))))) / 1.0))))) * ((v13 - (-2.25 / -2.25)) / v13))) / 1.0) - (v13 / float64(((0.1 / 1.0) / 1e10)))))) {
                        var v94 uint32 = (^(v87))
                        _ = v94
                    }
                    _ = v93
                } else {
                    v95 := GenS2{F0: v14, F1: 18446744073709551615, F2: v14}
                    _ = v95
                }
                _ = fi91
                _ = v92
            }
            _ = fi90
        }
    }
    var v96 int32 = -2147483648
    var v97 uint16 = 42
    fmt.Println(v3, v97)
    var v98 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v81
    _ = v86
    _ = v87
    _ = v88
    _ = v96
    _ = v97
    _ = v98
}
