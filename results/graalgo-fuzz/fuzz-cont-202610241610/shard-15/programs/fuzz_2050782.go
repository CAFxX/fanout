package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int8; F2 int16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint32; F1 int64; F2 int64 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 string; F1 int; F2 int32 }
type GenS9 struct { F0 int; F1 int64; F2 bool }
func genFn0(p0 uint16, p1 uint32, p2 uint64) uint8 {
    var v7 int = int(p1)
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int16, p1 bool, p2 uint32) uint16 {
    fmt.Println(p1, p0, p0)
    p2 += p2
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(p1, p1, p0)
            fmt.Println(fi9, p0)
            _ = fi10
        }
        switch "\n\t" {
        case "":
            p1 = p1
        case "\n\t":
            var v11 int32 = -2147483648
            _ = v11
        default:
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := []int{fi9, (fi12 & fi9), (fi12 / ((875) | 1)), 904}
                var v14 bool = p1
                var v15 int64 = 0
                switch (fi12 | -177) {
                case 4294967295:
                    var v16 int32 = 0
                    _ = v16
                case 127:
                    switch 651 {
                    case 42:
                        if p1 {
                            p1 = p1
                            v15 -= v15
                        } else {
                            switch "hello" {
                            case "":
                                p2 = p2
                            case "\n\t":
                                var v17 bool = (!(v14 || p1))
                                _ = v17
                            default:
                                if false {
                                    if len(v13) > 0 {
                                        v13[0] = (fi9 * (fi12 ^ (fi12 + 4294967295)))
                                    }
                                } else {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                }
                            }
                            var v18 float64 = 0.1
                            _ = v18
                        }
                    case -303:
                        var v19 uint8 = 128
                        _ = v19
                    case -9223372036854775808:
                        fmt.Println(p2)
                        if p1 {
                            v14 = v14
                            v20 := GenS1{F0: 42}
                            _ = v20
                        } else {
                            switch "" {
                            case "hello":
                                var v21 int16 = p0
                                _ = v21
                            }
                            var v22 uint = 0
                            v23 := GenS2{F0: (v22 + (v22 | v22))}
                            _ = v22
                            _ = v23
                        }
                        if (18446744073709551615 == 2) {
                            switch ("foo bar" + ("\n\t" + "hello")) {
                            case "hello":
                                var v24 int32 = 42
                                _ = v24
                            case "foo bar":
                                v15 = v15
                                var v25 int64 = v15
                                if v14 {
                                    fmt.Println(v25, p0, v25)
                                    v13 = append(v13, fi9)
                                    var v26 uint16 = 32767
                                    _ = v26
                                } else {
                                    if true {
                                        fmt.Println(v14, v14, v25)
                                        if p1 {
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            var v27 uint8 = 2
                                            v14 = (!((0 <= (v25 - (v25 ^ (v15 + 233)))) || v14))
                                            _ = v27
                                        }
                                        switch fi12 {
                                        case 640:
                                            v28 := GenS2{F0: 255}
                                            _ = v28
                                        case 32767:
                                            fmt.Println(v14, fi9, v15)
                                        case 42:
                                            p1 = ((!(!(42 > 128))) && (824 == 998))
                                        default:
                                            v13 = append(v13, fi9)
                                        }
                                    }
                                }
                                _ = v25
                            }
                            p1 = true
                            if v14 {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                v29 := GenS3{F0: int(p0), F1: -13, F2: p0}
                                _ = v29
                            }
                        } else {
                            v14 = (0.1 == float64(((0.1 - 0.1) + ((1.5 * ((-2.25 * ((((((0.0 - 1.5) + 0.0) + (3.14159 / 1e10)) + 1e10) - ((1.5 - 0.1) + (((1e10 - 1.5) * 1.5) - 1.5))) * 1e10)) + 1.5)) * (1.5 * 1.5)))))
                        }
                    default:
                        if (false || p1) {
                            p0 = (p0 % (((p0 >> ((255) & 63))) | 1))
                            p1 = v14
                            v30 := GenS4{F0: 1, F1: 2}
                            _ = v30
                        } else {
                            fmt.Println(v14, fi9, fi12)
                            if (v14 && v14) {
                                if p1 {
                                    switch -1 {
                                    case -728:
                                        switch fi9 {
                                        case 448:
                                            v31 := GenS1{F0: 128}
                                            _ = v31
                                        case 4294967295:
                                            var v32 int64 = v15
                                            _ = v32
                                        }
                                    }
                                    v14 = v14
                                }
                            } else {
                                var v33 int = fi9
                                _ = v33
                            }
                        }
                    }
                case 0:
                    v15 = (v15 << ((uint(p2)) & 63))
                }
                var v34 int64 = (v15 % ((375) | 1))
                if p1 {
                    var v35 uint8 = 2
                    var v36 uint64 = 26
                    _ = v35
                    _ = v36
                } else {
                    var v37 string = ("" + "foo bar")
                    fmt.Println(v14, v37, v14)
                    _ = len(v13)
                    _ = v37
                }
                var v38 uint64 = 76
                if p1 {
                    v13 = append(v13, 9223372036854775807)
                } else {
                    p2 += 266
                    var v39 int32 = 1
                    _ = v39
                }
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
                _ = v34
                _ = v38
            }
        }
        p0 += -32768
        _ = fi9
    }
    var v40 uint32 = (p2 * (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    return 42
}

func main() {
    var v1 uint8 = 189
    v2 := make(map[uint32]string)
    var v3 bool = ((820 < 1) || (!(-2147483648 <= 0)))
    v4 := []int{4294967295, int(v1)}
    var v5 float32 = 3.14159
    var v6 float32 = v5
    var v41 float32 = 1e10
    v6 -= 0.1
    v1 += (v1 * (v1 ^ v1))
    for fi42 := 0; fi42 < 12; fi42++ {
        fmt.Println(fi42, v6)
        _ = fi42
    }
    v3 = v3
    if v3 {
        fmt.Println(v5, v41, v41)
        for fi43 := 0; fi43 < 14; fi43++ {
            var v44 float64 = (1.5 - (1e10 + -2.25))
            v44 = ((v44 + v44) / 1.5)
            _ = fi43
            _ = v44
        }
        if (!(0 <= 100)) {
            var v45 int = -148
            _ = v45
        } else {
            for fi46 := 0; fi46 < 15; fi46++ {
                if len(v4) > 0 {
                    v4[0] = (fi46 % (((fi46 >> ((255) & 63))) | 1))
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi46
            }
        }
    }
    var v47 int = 0
    var v48 uint8 = (v1 / (((v1 | (v1 - (v1 ^ v1)))) | 1))
    v48 -= v1
    for fi49 := 0; fi49 < 10; fi49++ {
        var v50 int64 = 0
        var v51 uint32 = 42
        if (v3 || (("a" + ("a" + "hello")) > "\n\t")) {
            for fi52 := 0; fi52 < 6; fi52++ {
                if (-1 <= -2147483648) {
                    _ = len(v2)
                    v50 = 9223372036854775807
                    fmt.Println(v41, v3)
                    switch (^(v47)) {
                    case 1:
                        var v53 int64 = 598
                        _ = v53
                    case 128:
                        if (((0.0 / (float64(v6) + 1.5)) > (float64(v41) * 0.1)) && (!false)) {
                            v2[v51] = "a"
                        }
                    case 0:
                        v54 := GenS2{F0: 201}
                        _ = v54
                    }
                }
                v55 := GenS3{F0: v47, F1: 127, F2: 0}
                _ = len(v4)
                _ = fi52
                _ = v55
            }
            var v56 int16 = -32768
            _ = v56
        } else {
            var v57 float64 = 0.1
            if (1 > 0) {
                var v58 bool = ("" != "foo bar")
                v59 := GenS2{F0: 18446744073709551615}
                if v3 {
                    switch "\n\t" {
                    case "foo bar":
                        v57 -= (((0.0 - 1e10) / v57) * 3.14159)
                    case "\n\t":
                        var v60 int8 = 100
                        _ = v60
                    case "a":
                        if v58 {
                            var v61 string = "\n\t"
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v61
                        }
                    default:
                        var v62 int16 = 479
                        _ = v62
                    }
                    switch 777 {
                    case 42:
                        v41 = (0.1 / ((v5 + (0.1 / (float32(v57) / ((((0.0 / (v41 * v5)) + v5) - 1.5) + float32((float64(v6) + (v57 / -2.25))))))) / 1.5))
                    case 2147483647:
                        if v3 {
                            for fi63 := 0; fi63 < 2; fi63++ {
                                v64 := v2[v51]
                                var v65 int16 = 798
                                fmt.Println(fi63, v65)
                                if true {
                                    v66 := GenS3{F0: fi63, F1: 100, F2: (-(v65))}
                                    _ = v66
                                } else {
                                    v67 := GenS5{F0: v65, F1: 65535}
                                    v51 = v51
                                    var v68 uint = 255
                                    _ = v67
                                    _ = v68
                                }
                                switch v64 {
                                case "hello":
                                    var v69 uint8 = 0
                                    _ = v69
                                case "\n\t":
                                    v70 := GenS6{F0: v51, F1: v50, F2: -1}
                                    _ = v70
                                case "foo bar":
                                    switch fi49 {
                                    case -1:
                                        v48 -= 128
                                    case 32767:
                                        var v71 uint16 = 1
                                        _ = v71
                                    case -9223372036854775808:
                                        if (v3 || ((int16(v50) > v65) && v3)) {
                                            if len(v4) > 0 {
                                                v4[0] = fi63
                                            }
                                            v65 += (v65 ^ -32768)
                                        }
                                    default:
                                        switch v64 {
                                        case "hello":
                                            var v72 uint32 = 986
                                            _ = v72
                                        case "":
                                            switch 42 {
                                            case 32767:
                                                v65 = v65
                                            default:
                                                switch v64 {
                                                case "\n\t":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case "a":
                                                    if v58 {
                                                        fmt.Println(v47, v64, v64)
                                                    }
                                                default:
                                                    if (1 <= 117) {
                                                        delete(v2, (v51 >> ((960) & 63)))
                                                        fmt.Println(v65, v48)
                                                        var v73 int64 = (v50 % (((^(v50))) | 1))
                                                        _ = v73
                                                    } else {
                                                        if true {
                                                            var v74 float64 = (1e10 + (0.1 / v57))
                                                            var v75 uint64 = 185
                                                            _ = len(v2)
                                                            if v58 {
                                                                v51 = (v51 & uint32(v75))
                                                            } else {
                                                                fmt.Println(v50, v50, v65)
                                                            }
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        v76 := GenS4{F0: (v50 % ((v50) | 1)), F1: (v48 >> ((716) & 63))}
                                                        switch "a" {
                                                        case "a":
                                                            switch 1 {
                                                            case 908:
                                                                if v58 {
                                                                    fmt.Println(v1, v57)
                                                                    v2[843] = v64
                                                                } else {
                                                                    _ = len(v4)
                                                                    v3 = true
                                                                }
                                                            case 9223372036854775807:
                                                                fmt.Println(v41, v51, v48)
                                                            case -619:
                                                                var v77 bool = (fi63 <= v47)
                                                                _ = v77
                                                            }
                                                            v41 -= ((v6 + float32(v57)) / (-2.25 - 1.5))
                                                            switch ("a" + (v64 + (v64 + (v64 + v64)))) {
                                                            case "\n\t":
                                                                v57 = v57
                                                            case "foo bar":
                                                                switch (fi49 << ((747) & 63)) {
                                                                case -391:
                                                                    if (42 < 0) {
                                                                        v1 += 0
                                                                        v78 := GenS3{F0: (fi63 * (fi63 & fi49)), F1: -1, F2: (-(v65))}
                                                                        _ = v78
                                                                    }
                                                                default:
                                                                    var v79 uint8 = v48
                                                                    _ = v79
                                                                }
                                                            }
                                                        default:
                                                            v80 := GenS5{F0: v65, F1: 840}
                                                            _ = v80
                                                        }
                                                        _ = v76
                                                    }
                                                }
                                            }
                                        default:
                                            var v81 uint = 42
                                            _ = v81
                                        }
                                        var v82 uint = 1
                                        var v83 float32 = v5
                                        _ = v82
                                        _ = v83
                                    }
                                default:
                                    v47 -= (-(fi63))
                                }
                                switch (fi49 >> ((65535) & 63)) {
                                case 42:
                                    v65 = v65
                                default:
                                    v84 := GenS7{F0: -40, F1: 68}
                                    _ = v84
                                }
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            v6 = v41
                        }
                    case -664:
                        v85 := GenS3{F0: (fi49 / (((fi49 % (((v47 & fi49)) | 1))) | 1)), F1: -117, F2: 0}
                        _ = v85
                    }
                } else {
                    fmt.Println(v1, v41)
                    switch ("\n\t" + "") {
                    case "\n\t":
                        v51 += (v51 << ((207) & 63))
                    case "foo bar":
                        v2[(v51 / ((v51) | 1))] = "hello"
                    }
                    var v86 float64 = (-2.25 - v57)
                    _ = v86
                }
                _ = v58
                _ = v59
            } else {
                v87 := GenS8{F0: "hello", F1: -1, F2: int32(v50)}
                _ = v87
            }
            _ = v57
        }
        _ = fi49
        _ = v50
        _ = v51
    }
    delete(v2, 0)
    _ = len(v4)
    switch v47 {
    case 32768:
        fmt.Println(v41, v3, v5)
    case 980:
        v6 = (1e10 - (-2.25 + v6))
    case 32767:
        var v88 uint8 = genFn0(662, 2, 42)
        _ = v88
    default:
        switch "\n\t" {
        case "foo bar":
            if v3 {
                _ = len(v2)
            }
        }
    }
    v89 := GenS9{F0: (v47 ^ v47), F1: -115, F2: (115 < 52)}
    var v90 uint8 = 1
    v48 = v1
    v91 := GenS1{F0: v90}
    for fi92 := 0; fi92 < 7; fi92++ {
        for fi93 := 0; fi93 < 20; fi93++ {
            var v94 string = "\n\t"
            var v95 int64 = 0
            if false {
                v90 += 193
                if true {
                    var v96 int64 = (v95 | (v95 & -9223372036854775808))
                    if ((!(v95 > 1)) && (!v3)) {
                        v2[2147483647] = v94
                        if v3 {
                            var v97 int8 = -1
                            _ = len(v2)
                            _ = v97
                        }
                        fmt.Println(v6, v95)
                    }
                    _ = v96
                }
            }
            _ = fi93
            _ = v94
            _ = v95
        }
        _ = fi92
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
