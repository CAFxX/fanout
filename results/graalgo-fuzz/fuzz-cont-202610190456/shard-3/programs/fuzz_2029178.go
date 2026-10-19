package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint32 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint32, p1 int32) uint8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            if (127 > 1) {
                if true {
                    v9 := GenS1{F0: -9223372036854775808}
                    var v10 bool = (9223372036854775807 < -796)
                    p1 -= p1
                    _ = v10
                    _ = v9
                } else {
                    if ((0.1 - (1e10 - (((3.14159 * 0.0) + (3.14159 / (0.1 / 0.1))) * float64(1.5)))) == 1.5) {
                        var v11 int = int(p1)
                        _ = v11
                    }
                    v12 := GenS1{F0: 464}
                    if (-2.25 < 3.14159) {
                        var v13 float32 = -2.25
                        var v14 bool = false
                        _ = v13
                        _ = v14
                    } else {
                        if false {
                            v15 := []int{532}
                            v15 = append(v15, 128)
                            fmt.Println(p0, fi7, fi8)
                            if len(v15) > 0 {
                                v15[0] = (fi7 - 4294967295)
                            }
                            _ = v15
                        }
                        v16 := []int{fi7}
                        if (((-2.25 * (-2.25 + 3.14159)) / 0.1) <= (((0.0 * float64(1e10)) / 1.5) / float64(float32(-2.25)))) {
                            fmt.Println(p0, fi7, fi8)
                        }
                        switch fi7 {
                        case 2147483647:
                            p1 = p1
                        case 42:
                            var v17 bool = false
                            _ = v17
                        }
                        var v18 int = (fi7 + -2147483648)
                        _ = v16
                        _ = v18
                    }
                    v19 := GenS2{F0: 0, F1: int16(p1), F2: 2}
                    var v20 int16 = 254
                    _ = v12
                    _ = v19
                    _ = v20
                }
                switch (^(fi8)) {
                case 2147483647:
                    v21 := GenS3{F0: 1, F1: 42, F2: p1}
                    _ = v21
                case 42:
                    p0 -= p0
                default:
                    fmt.Println(p0)
                    if true {
                        v22 := []int{}
                        p0 = (p0 ^ 1)
                        v23 := GenS4{F0: 0, F1: (-(fi8)), F2: p0}
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS2{F0: 1, F1: 0, F2: 127}
                    _ = v24
                }
            } else {
                v25 := []int{(fi8 & (^(fi8))), (fi8 << ((1) & 63)), int(p1), 0}
                if (!(!((2 <= 240) && ((100 >= 0) || false)))) {
                    fmt.Println(p1)
                    v25 = append(v25, (fi8 | (fi8 % ((fi7) | 1))))
                    var v26 uint64 = 2
                    _ = v26
                } else {
                    p1 = p1
                }
                _ = v25
            }
            var v27 int64 = 0
            v28 := GenS5{F0: v27, F1: 32767}
            _ = fi8
            _ = v27
            _ = v28
        }
        _ = fi7
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 int8 = 127
        _ = fi29
        _ = v30
    }
    var v31 float64 = (float64(0.0) - 1.5)
    _ = p0
    _ = p1
    _ = v31
    return 237
}

func main() {
    var v1 uint8 = 92
    var v2 int64 = 1
    var v3 string = "\n\t"
    v4 := []int{-9223372036854775808, 4294967295}
    var v5 uint16 = 837
    v6 := []string{(((v3 + v3) + "hello") + v3)}
    fmt.Println(v1, v3, v5)
    for fi32 := 0; fi32 < 7; fi32++ {
        for fi33 := 0; fi33 < 12; fi33++ {
            v5 = (v5 * v5)
            if true {
                if (((0.1 * 1e10) - 0.1) <= 0.1) {
                    switch int(v5) {
                    case -1:
                        v34 := GenS5{F0: -1, F1: -1}
                        _ = v34
                    case -2147483648:
                        v1 = v1
                        switch "a" {
                        case "hello":
                            v5 = 42
                        case "\n\t":
                            if (65535 <= v5) {
                                var v35 int8 = 127
                                _ = v35
                            } else {
                                v1 -= uint8(v5)
                            }
                        case "a":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        switch fi32 {
                        case -316:
                            v5 = (v5 + v5)
                        case -2147483648:
                            _ = len(v6)
                        case -775:
                            v36 := GenS1{F0: v2}
                            _ = v36
                        }
                    default:
                        if len(v6) > 0 {
                            v6[0] = ("" + ("hello" + v3))
                        }
                    }
                    v1 += v1
                }
            }
            _ = fi33
        }
        v37 := GenS6{F0: 42, F1: 1000}
        _ = fi32
        _ = v37
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        if true {
            v39 := GenS7{F0: 42}
            for fi40 := 0; fi40 < 4; fi40++ {
                v6 = append(v6, v3)
                _ = fi40
            }
            _ = v39
        } else {
            fmt.Println(fi38)
        }
        _ = fi38
    }
    fmt.Println(v5, v2, v3)
    v1 += 128
    var v41 uint64 = uint64(v1)
    for fi42 := 0; fi42 < 11; fi42++ {
        if (!false) {
            v2 = (v2 + (v2 << ((65535) & 63)))
        } else {
            for fi43 := 0; fi43 < 15; fi43++ {
                v44 := GenS5{F0: (v2 & (^(v2))), F1: 32767}
                if (-66 > -1) {
                    var v45 int8 = 71
                    v45 = 42
                    switch 4294967295 {
                    case 108:
                        fmt.Println(fi42, fi43, fi42)
                    case 32767:
                        var v46 int32 = 42
                        _ = v46
                    case 1:
                        v1 += v1
                    }
                    _ = v45
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v2 = v2
                    v1 = uint8(v5)
                }
                var v47 int = (fi43 | (fi42 % (((fi43 & 0)) | 1)))
                _ = fi43
                _ = v44
                _ = v47
            }
            var v48 int8 = 46
            _ = v48
        }
        for fi49 := 0; fi49 < 15; fi49++ {
            v50 := GenS8{F0: 0, F1: v41}
            var v51 float64 = -2.25
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi42
    }
    v3 = "hello"
    if (!true) {
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        switch "foo bar" {
        case "foo bar":
            v3 = ""
        case "a":
            v1 -= v1
        }
    }
    var v52 int16 = 154
    var v53 int = -1
    fmt.Println(v53)
    if false {
        for fi54 := 0; fi54 < 1; fi54++ {
            var v55 int32 = -1
            switch v3 {
            case "":
                v56 := GenS9{F0: 2147483647, F1: genFn0(2147483647, 612)}
                _ = v56
            case "\n\t":
                v57 := GenS8{F0: -1, F1: 33}
                _ = v57
            default:
                var v58 uint8 = v1
                _ = v58
            }
            v3 = "\n\t"
            _ = fi54
            _ = v55
        }
    }
    if (!true) {
        var v59 uint8 = 1
        v1 -= (v59 / ((2) | 1))
        _ = v59
    } else {
        for fi60 := 0; fi60 < 16; fi60++ {
            var v61 bool = ((v41 - (v41 / (((v41 << ((255) & 63))) | 1))) <= (v41 << ((2) & 63)))
            _ = fi60
            _ = v61
        }
        v5 = 0
        switch (v53 << ((18446744073709551615) & 63)) {
        case 2147483647:
            for fi62 := 0; fi62 < 4; fi62++ {
                v63 := GenS3{F0: (v52 >> ((18446744073709551615) & 63)), F1: v41, F2: -1}
                _ = len(v6)
                if true {
                    switch 619 {
                    case 32768:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 2147483647:
                        for fi64 := 0; fi64 < 15; fi64++ {
                            v65 := GenS2{F0: 166, F1: int16(v5), F2: (v1 % ((v1) | 1))}
                            var v66 int8 = -76
                            if (!(!(1e10 < 1.5))) {
                                if len(v6) > 0 {
                                    v6[0] = v3
                                }
                                switch 32767 {
                                case 130:
                                    if ((false || true) && false) {
                                        v5 = (^(v5))
                                        v5 = (v5 << ((4294967295) & 63))
                                    } else {
                                        v67 := GenS10{F0: v53}
                                        switch v3 {
                                        case "hello":
                                            fmt.Println(fi64, v2)
                                        case "a":
                                            fmt.Println(v2, v3)
                                        default:
                                            v53 += fi64
                                        }
                                        _ = v67
                                    }
                                }
                                v68 := GenS4{F0: uint64(v1), F1: 42, F2: 2147483647}
                                _ = v68
                            } else {
                                v4 = append(v4, (-(fi62)))
                                var v69 int8 = v66
                                if (!false) {
                                    if true {
                                        if false {
                                            var v70 uint64 = 681
                                            _ = v70
                                        }
                                        var v71 float32 = -2.25
                                        var v72 bool = false
                                        _ = v71
                                        _ = v72
                                    } else {
                                        switch -45 {
                                        case -9223372036854775808:
                                            v73 := GenS8{F0: -1, F1: 510}
                                            if len(v6) > 0 {
                                                v6[0] = ("hello" + v3)
                                            }
                                            _ = v73
                                        case -2147483648:
                                            fmt.Println(fi62, v3)
                                        case 1:
                                            var v74 int32 = 2147483647
                                            _ = v74
                                        default:
                                            var v75 int32 = 262
                                            _ = v75
                                        }
                                        v5 = v5
                                        v41 = uint64(v1)
                                        if true {
                                            var v76 uint16 = v5
                                            _ = v76
                                        }
                                        var v77 int = (fi64 / ((-2147483648) | 1))
                                        switch v3 {
                                        case "":
                                            switch (v3 + "\n\t") {
                                            case "":
                                                if (0 != 18446744073709551615) {
                                                    v2 += (v2 & (^(v2)))
                                                    v77 += (fi64 % ((973) | 1))
                                                }
                                            case "hello":
                                                v5 += v5
                                            case "a":
                                                var v78 int32 = -2147483648
                                                _ = v78
                                            default:
                                                v79 := GenS4{F0: v41, F1: fi62, F2: 1}
                                                _ = v79
                                            }
                                        case "foo bar":
                                            v66 = v69
                                        case "\n\t":
                                            var v80 uint16 = (v5 << ((42) & 63))
                                            fmt.Println(v3)
                                            _ = v80
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = 0
                                            }
                                        }
                                        switch (v77 % (((-(v53))) | 1)) {
                                        case 1:
                                            var v81 uint8 = (v1 - v1)
                                            _ = v81
                                        case -1:
                                            v69 -= -80
                                        }
                                        if ((3.14159 * 1.5) < (((0.0 * -2.25) * 0.1) * 0.1)) {
                                            v82 := GenS11{F0: 42, F1: v3}
                                            if false {
                                                if (!true) {
                                                    v66 = 127
                                                }
                                                fmt.Println(v3)
                                            } else {
                                                if false {
                                                    var v83 int8 = v66
                                                    var v84 uint8 = v1
                                                    var v85 int = (fi62 << ((308) & 63))
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    v86 := GenS4{F0: (v41 & v41), F1: 32767, F2: 877}
                                                    var v87 int32 = -173
                                                    _ = v86
                                                    _ = v87
                                                }
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v82
                                        } else {
                                            var v88 int16 = -32768
                                            _ = v88
                                        }
                                        _ = v77
                                    }
                                    v3 = "a"
                                }
                                _ = v69
                            }
                            if len(v6) > 0 {
                                v6[0] = v3
                            }
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                    case 128:
                        if (!(42 != int32(v5))) {
                            _ = len(v4)
                            switch fi62 {
                            case 32767:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v89 := GenS12{F0: 4294967295, F1: (v1 & (v1 / (((v1 & 76)) | 1)))}
                            _ = v89
                        } else {
                            for fi90 := 0; fi90 < 5; fi90++ {
                                fmt.Println(v53)
                                v91 := GenS6{F0: 958, F1: (v52 - v52)}
                                _ = fi90
                                _ = v91
                            }
                            var v92 float64 = 1e10
                            _ = v92
                        }
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                } else {
                    v3 = v3
                }
                _ = fi62
                _ = v63
            }
        case 128:
            fmt.Println(v52)
        case -202:
            switch "a" {
            case "\n\t":
                var v93 float64 = ((3.14159 + 0.1) / 1.0)
                _ = v93
            }
        }
    }
    var v94 int32 = 238
    v95 := GenS4{F0: (^(v41)), F1: v53, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v94
    _ = v95
}
