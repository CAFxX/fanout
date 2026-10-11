package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 int; F2 uint }
type GenS10 struct { F0 uint64; F1 string; F2 int16 }
type GenS11 struct { F0 string; F1 int32 }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 uint; F1 uint64 }
func genFn0(p0 int8) int8 {
    for fi5 := 0; fi5 < 10; fi5++ {
        if true {
            if false {
                var v6 uint = 862
                _ = v6
            }
            var v7 uint = 42
            _ = v7
        } else {
            fmt.Println(p0)
            for fi8 := 0; fi8 < 4; fi8++ {
                p0 = 21
                _ = fi8
            }
        }
        p0 = -128
        if (true && false) {
            switch (fi5 / (((fi5 / ((fi5) | 1))) | 1)) {
            case 128:
                if false {
                    var v9 int64 = 12
                    p0 += p0
                    _ = v9
                }
            default:
                fmt.Println(fi5, p0)
            }
            v10 := []int{fi5, (fi5 >> ((42) & 63)), (fi5 << ((4294967295) & 63)), fi5}
            fmt.Println(p0, fi5)
            _ = v10
        }
        _ = fi5
    }
    p0 = p0
    _ = p0
    return p0
}
func genFn1(p0 uint64, p1 int16) bool {
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := []int{(fi11 | 680), (fi11 % (((fi11 * fi11)) | 1)), 2147483647, -9223372036854775808}
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS1{F0: 2}
            if true {
                v12 = append(v12, (fi11 ^ fi13))
            } else {
                p0 += p0
                if len(v12) > 0 {
                    v12[0] = fi11
                }
                switch fi11 {
                case 19:
                    if (0.0 == 0.1) {
                        if true {
                            v15 := GenS2{F0: (("\n\t" + "a") + "hello")}
                            switch "foo bar" {
                            case "\n\t":
                                v16 := GenS2{F0: "\n\t"}
                                _ = v16
                            }
                            _ = v15
                        } else {
                            var v17 uint = 42
                            var v18 int8 = -58
                            _ = v17
                            _ = v18
                        }
                        switch "hello" {
                        case "\n\t":
                            v19 := GenS2{F0: (("\n\t" + ("hello" + (("a" + ("\n\t" + ("foo bar" + "\n\t"))) + "hello"))) + "a")}
                            _ = v19
                        case "hello":
                            var v20 bool = true
                            _ = v20
                        default:
                            if (true || (!true)) {
                                p0 += p0
                            }
                        }
                    } else {
                        v21 := GenS1{F0: 4294967295}
                        p1 += p1
                        switch fi13 {
                        case -774:
                            v22 := GenS2{F0: ("\n\t" + "hello")}
                            _ = v22
                        case -9223372036854775808:
                            if (255 <= 65) {
                                fmt.Println(p1, p1, fi13)
                                var v23 int = fi13
                                _ = v23
                            }
                        case 9223372036854775807:
                            switch "\n\t" {
                            case "":
                                v24 := GenS1{F0: 127}
                                _ = v24
                            }
                        default:
                            fmt.Println(fi13, p1, p1)
                            p1 = (p1 ^ p1)
                            p1 -= p1
                        }
                        _ = v21
                    }
                default:
                    fmt.Println(fi11, p0, p0)
                }
            }
            p1 -= (p1 * (p1 | (^(p1))))
            v25 := GenS3{F0: (p0 ^ 0), F1: true, F2: p1}
            _ = fi13
            _ = v14
            _ = v25
        }
        _ = fi11
        _ = v12
    }
    v26 := []int{0, 93, 776}
    var v27 uint = 18446744073709551615
    var v28 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v28
    return true
}

func main() {
    var v1 int8 = 1
    var v2 float32 = (0.0 - -2.25)
    var v3 uint16 = 2
    var v4 int = 146
    switch (v4 / ((v4) | 1)) {
    case 308:
        v2 = float32((-2.25 / 1.0))
    case 312:
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := GenS4{F0: v1, F1: 0}
            _ = fi29
            _ = v30
        }
    }
    v31 := []int{v4, (v4 >> ((0) & 63)), (-(v4))}
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    var v32 string = ""
    var v33 float64 = 0.0
    var v34 float32 = ((v2 - (float32(((v33 + (v33 / 1.0)) - (float64(v2) + -2.25))) * ((1e10 / v2) * 0.1))) / v2)
    var v35 int = v4
    if (!false) {
        if false {
            fmt.Println(v4, v2, v4)
        } else {
            _ = len(v31)
            var v36 uint16 = v3
            v31 = append(v31, v35)
            _ = v36
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            if (fi37 < (v4 & v35)) {
                var v38 float32 = (1e10 / (v34 + float32(0.1)))
                _ = v38
            }
            _ = fi37
        }
        var v39 uint8 = 0
        _ = v39
    }
    var v40 bool = true
    v3 -= (^(v3))
    switch v32 {
    case "":
        for fi41 := 0; fi41 < 13; fi41++ {
            v40 = (2 < v3)
            if ((2 >= 97) || (!(!(!false)))) {
                for fi42 := 0; fi42 < 14; fi42++ {
                    var v43 int64 = -1
                    v31 = append(v31, -987)
                    if (v33 < v33) {
                        switch v32 {
                        case "foo bar":
                            v44 := GenS5{F0: 0}
                            _ = v44
                        case "a":
                            if len(v31) >= 2 {
                                v31 = v31[1:len(v31)]
                            }
                        case "hello":
                            v45 := GenS6{F0: 898}
                            _ = v45
                        default:
                            var v46 string = ("\n\t" + "foo bar")
                            _ = v46
                        }
                        v1 += (-(v1))
                        fmt.Println(v33, fi42)
                        v31 = append(v31, (fi42 * fi42))
                        if true {
                            var v47 int64 = (v43 / ((v43) | 1))
                            var v48 uint64 = 42
                            if len(v31) > 0 {
                                v31[0] = (^(fi42))
                            }
                            _ = v47
                            _ = v48
                        } else {
                            fmt.Println(v35, v34, v32)
                            _ = len(v31)
                        }
                        var v49 int8 = v1
                        if ((fi42 << ((494) & 63)) >= (v4 / ((v35) | 1))) {
                            _ = len(v31)
                        } else {
                            v1 = v1
                        }
                        if len(v31) > 0 {
                            v31[0] = fi41
                        }
                        v50 := GenS2{F0: v32}
                        if len(v31) > 0 {
                            v31[0] = (-(v4))
                        }
                        v51 := GenS3{F0: 18446744073709551615, F1: true, F2: 155}
                        if (v2 != float32(v33)) {
                            v43 += v43
                            if (v40 && (2147483647 < 0)) {
                                v52 := GenS2{F0: ("hello" + v32)}
                                _ = v52
                            } else {
                                v34 = float32(0.1)
                                fmt.Println(v43, v43)
                                if (!((65535 <= 665) && v40)) {
                                    switch (v32 + "hello") {
                                    case "hello":
                                        v40 = ((v1 - (v1 | v1)) == (v49 ^ 0))
                                    case "\n\t":
                                        var v53 uint8 = 235
                                        _ = v53
                                    default:
                                        if (v43 <= v43) {
                                            if ((-(v43)) == -1) {
                                                v54 := GenS6{F0: 994}
                                                v35 = 0
                                                fmt.Println(v35, v33)
                                                _ = v54
                                            } else {
                                                _ = len(v31)
                                            }
                                        }
                                    }
                                    v1 = genFn0(v49)
                                    if (42 == v3) {
                                        var v55 int8 = (v49 * 127)
                                        _ = v55
                                    }
                                } else {
                                    var v56 int = fi42
                                    _ = v56
                                }
                            }
                        } else {
                            if ((32767 < 1) || true) {
                                v57 := GenS1{F0: 4294967295}
                                var v58 float32 = (v2 + 1.5)
                                _ = v57
                                _ = v58
                            }
                            switch v32 {
                            case "hello":
                                switch -130 {
                                case 4294967295:
                                    v31 = append(v31, (v35 % ((v35) | 1)))
                                case 128:
                                    v31 = append(v31, fi42)
                                }
                            }
                            var v59 uint64 = 0
                            _ = v59
                        }
                        v43 += (v43 | (^(v43)))
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    if (!v40) {
                        var v60 int32 = 1
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        var v61 bool = false
                        _ = v60
                        _ = v61
                    } else {
                        v62 := GenS7{F0: 0, F1: 2}
                        if v40 {
                            v40 = ((int8(v3) <= v1) || (((v32 + ((v32 + v32) + v32)) + ((v32 + (v32 + v32)) + "")) <= (v32 + v32)))
                        }
                        _ = v62
                    }
                    _ = fi42
                    _ = v43
                }
            } else {
                if (!v40) {
                    if (1 != 0) {
                        for fi63 := 0; fi63 < 3; fi63++ {
                            v3 = 42
                            _ = fi63
                        }
                        switch -1 {
                        case 128:
                            for fi64 := 0; fi64 < 2; fi64++ {
                                v33 -= float64((1e10 - 1e10))
                                v34 = float32(0.1)
                                _ = fi64
                            }
                        case 42:
                            var v65 uint32 = 0
                            _ = v65
                        case 824:
                            v35 = (v4 >> ((133) & 63))
                        default:
                            v34 -= v34
                        }
                        var v66 uint = 2
                        _ = v66
                    }
                    if ((fi41 << ((421) & 63)) >= (v35 + -9223372036854775808)) {
                        v32 = v32
                    } else {
                        switch (v35 - v4) {
                        case 32767:
                            if v40 {
                                for fi67 := 0; fi67 < 18; fi67++ {
                                    if (!false) {
                                        var v68 uint16 = (v3 - v3)
                                        switch fi67 {
                                        case -581:
                                            var v69 int8 = (v1 << ((0) & 63))
                                            _ = v69
                                        case 42:
                                            v68 += (v68 & v68)
                                        case -2147483648:
                                            if len(v31) > 0 {
                                                v31[0] = 128
                                            }
                                        }
                                        if true {
                                            fmt.Println(v1, v4, v2)
                                            var v70 uint32 = 861
                                            var v71 int16 = 0
                                            v72 := GenS8{F0: (v71 - v71)}
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v68
                                    }
                                    switch v32 {
                                    case "a":
                                        v2 = v34
                                    case "hello":
                                        v32 = v32
                                    case "\n\t":
                                        v73 := GenS9{F0: 128, F1: (fi67 * (^(v4))), F2: 748}
                                        _ = v73
                                    default:
                                        var v74 uint64 = 355
                                        _ = v74
                                    }
                                    var v75 uint16 = (v3 % ((v3) | 1))
                                    _ = fi67
                                    _ = v75
                                }
                                v76 := GenS1{F0: 2}
                                var v77 float32 = (v2 + (3.14159 * 0.1))
                                _ = v76
                                _ = v77
                            }
                        case 32768:
                            v78 := GenS10{F0: 18446744073709551615, F1: v32, F2: 0}
                            _ = v78
                        default:
                            fmt.Println(v35, v35, v32)
                        }
                        v34 = v34
                        fmt.Println(v1, fi41, v34)
                    }
                } else {
                    var v79 uint64 = 442
                    _ = v79
                }
                if true {
                    v33 -= v33
                    v32 = (v32 + (v32 + "\n\t"))
                    switch v4 {
                    case -2147483648:
                        var v80 uint32 = 712
                        _ = v80
                    case -756:
                        var v81 uint16 = (^(v3))
                        _ = v81
                    }
                }
                for fi82 := 0; fi82 < 7; fi82++ {
                    v83 := GenS11{F0: v32, F1: 42}
                    var v84 uint8 = 1
                    v35 += fi82
                    var v85 uint = 1
                    _ = fi82
                    _ = v83
                    _ = v84
                    _ = v85
                }
            }
            _ = fi41
        }
    case "hello":
        var v86 float32 = (((v2 + v2) - ((float32(3.14159) + v34) * v2)) - v34)
        _ = v86
    case "a":
        v33 -= 0.1
    default:
        v87 := GenS12{F0: v3, F1: 0}
        _ = v87
    }
    if ((-32768 >= -64) || ((v3 & v3) < (v3 * 0))) {
        var v88 uint = 4294967295
        v40 = v40
        _ = v88
    }
    if ((^(v1)) >= v1) {
        if v40 {
            fmt.Println(v34, v33, v1)
            fmt.Println(v1, v32, v2)
            v34 = (-2.25 / 1.0)
        } else {
            for fi89 := 0; fi89 < 9; fi89++ {
                fmt.Println(v32, v35)
                fmt.Println(v33, v4, v3)
                _ = fi89
            }
            if (v35 > 522) {
                v90 := GenS13{F0: 587, F1: 106}
                var v91 uint16 = (v3 >> ((764) & 63))
                _ = v90
                _ = v91
            } else {
                var v92 uint16 = 2
                _ = v92
            }
            var v93 uint8 = 127
            _ = v93
        }
        v3 -= 122
    } else {
        var v94 uint64 = 2
        _ = v94
    }
    for fi95 := 0; fi95 < 2; fi95++ {
        v33 -= (((v33 * (float64((v34 / -2.25)) * v33)) * (3.14159 + float64(((float32((float64(1.5) / (v33 + 1e10))) * 1e10) * ((0.0 / 1.5) - 1e10))))) / v33)
        _ = fi95
    }
    v33 -= v33
    switch (v35 % ((v4) | 1)) {
    case 32768:
        var v96 string = "foo bar"
        _ = v96
    }
    v4 += v4
    v40 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
}
