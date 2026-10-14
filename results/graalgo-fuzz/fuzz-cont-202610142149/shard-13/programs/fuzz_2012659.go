package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 int8; F1 int8; F2 int }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int; F1 uint8; F2 uint32 }
type GenS10 struct { F0 uint; F1 int8; F2 int64 }
func genFn0() int32 {
    var v7 uint64 = 4294967295
    switch "foo bar" {
    case "a":
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint8 = 209
            v10 := []int{fi8, -626, -245}
            var v11 int64 = 367
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case "hello":
        v12 := []int{9223372036854775807, -572, 2147483647}
        _ = v12
    }
    var v13 uint = 42
    _ = v13
    _ = v7
    return 2147483647
}
func genFn1(p0 int64, p1 int32) bool {
    var v14 string = "foo bar"
    if (false && (185 <= 32767)) {
        if (!false) {
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(p0, p1)
                v16 := []int{fi15, -674}
                _ = fi15
                _ = v16
            }
            fmt.Println(v14)
            var v17 int = -1
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 16; fi18++ {
            fmt.Println(v14, fi18)
            if true {
                switch (("foo bar" + v14) + v14) {
                case "":
                    p1 -= p1
                case "hello":
                    p1 -= (p1 | genFn0())
                default:
                    v19 := GenS1{F0: 42}
                    _ = v19
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 float64 = 1.5
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 12; fi22++ {
                    v23 := GenS1{F0: 255}
                    var v24 float64 = 1e10
                    if false {
                        var v25 uint64 = 4294967295
                        if ((true && ((p0 + p0) != p0)) && false) {
                            if true {
                                var v26 int32 = genFn0()
                                _ = v26
                            } else {
                                fmt.Println(fi22, v25)
                                var v27 string = (v14 + v14)
                                _ = v27
                            }
                            v28 := []int{(fi18 >> ((4294967295) & 63)), -3, 4294967295, 127}
                            _ = v28
                        }
                        p0 = (^(p0))
                        _ = v25
                    } else {
                        v29 := []int{}
                        _ = v29
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            v30 := GenS2{F0: 559}
            _ = fi18
            _ = v30
        }
    }
    _ = p0
    _ = p1
    _ = v14
    return (true || ((!false) || (!(!(2147483647 >= 722)))))
}

func main() {
    var v1 uint16 = 52
    var v2 string = "\n\t"
    var v3 int32 = -143
    var v4 float64 = 0.0
    v5 := []uint32{}
    var v6 string = v2
    var v31 int8 = 1
    switch v2 {
    case "":
        v4 = (v4 / (v4 - v4))
    case "foo bar":
        v32 := GenS2{F0: 2147483647}
        _ = v32
    }
    var v33 float64 = (1.5 - v4)
    v34 := GenS3{F0: 1, F1: 563}
    var v35 string = ("" + (v2 + v2))
    v36 := GenS4{F0: true, F1: v1, F2: v6}
    v3 = v3
    v37 := GenS5{F0: (^(v31)), F1: (v31 + v31), F2: -2147483648}
    var v38 int32 = (v3 | -176)
    v39 := GenS2{F0: 2147483647}
    _ = len(v5)
    v3 = int32(v31)
    v5 = append(v5, 4294967295)
    if (!true) {
        var v40 float64 = 1e10
        switch -297 {
        case 9223372036854775807:
            for fi41 := 0; fi41 < 16; fi41++ {
                for fi42 := 0; fi42 < 17; fi42++ {
                    v3 += -789
                    v40 = v33
                    var v43 uint16 = (v1 - v1)
                    _ = fi42
                    _ = v43
                }
                if (((v2 + ("foo bar" + "a")) + v35) > "a") {
                    var v44 int16 = 32
                    switch fi41 {
                    case -2147483648:
                        var v45 bool = false
                        _ = v45
                    default:
                        if true {
                            switch (-(fi41)) {
                            case 9223372036854775807:
                                if false {
                                    for fi46 := 0; fi46 < 9; fi46++ {
                                        switch (fi46 ^ (fi41 / (((fi41 / ((32768) | 1))) | 1))) {
                                        case 457:
                                            if true {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                fmt.Println(v2, v31)
                                            }
                                        case -9223372036854775808:
                                            v47 := GenS5{F0: v31, F1: v31, F2: -9223372036854775808}
                                            _ = v47
                                        case -74:
                                            v4 += float64(3.14159)
                                        }
                                        fmt.Println(v38)
                                        v33 = v33
                                        v48 := GenS6{F0: -275, F1: 302}
                                        _ = fi46
                                        _ = v48
                                    }
                                }
                            case -25:
                                v2 = v35
                            }
                            for fi49 := 0; fi49 < 3; fi49++ {
                                if false {
                                    switch v2 {
                                    case "\n\t":
                                        switch fi49 {
                                        case 1:
                                            fmt.Println(v44, fi49)
                                        case -752:
                                            var v50 uint16 = v1
                                            if true {
                                                v51 := GenS7{F0: false}
                                                fmt.Println(fi41)
                                                _ = v51
                                            } else {
                                                if true {
                                                    if ((!true) || false) {
                                                        var v52 uint = 65535
                                                        var v53 int = 127
                                                        switch (fi49 - fi41) {
                                                        case -2147483648:
                                                            v54 := GenS3{F0: 42, F1: -1}
                                                            _ = v54
                                                        case 42:
                                                            fmt.Println(v31)
                                                        default:
                                                            v50 = v50
                                                        }
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        fmt.Println(v44, v2, fi41)
                                                    }
                                                    if true {
                                                        if (!(((!(false && (!false))) || true) || true)) {
                                                            switch (((v35 + (v2 + "hello")) + v35) + "hello") {
                                                            case "":
                                                                var v55 uint64 = 160
                                                                _ = v55
                                                            case "a":
                                                                v56 := GenS3{F0: 244, F1: -9223372036854775808}
                                                                _ = v56
                                                            }
                                                        }
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            }
                                            var v57 float32 = float32(float64((3.14159 * -2.25)))
                                            _ = v50
                                            _ = v57
                                        case 127:
                                            if true {
                                                if (((v33 >= (1.5 + (0.0 + float64(float32(v40))))) && true) && false) {
                                                    switch ("a" + v2) {
                                                    case "hello":
                                                        var v58 bool = true
                                                        _ = v58
                                                    }
                                                    if true {
                                                        v40 = (((v4 - v40) + ((((0.0 / 0.1) * (float64(1.5) * 3.14159)) + v33) / float64((((-2.25 + 3.14159) / 3.14159) + 1e10)))) / (float64(3.14159) / 1e10))
                                                        v5 = append(v5, 2147483648)
                                                        v4 = (v40 - v4)
                                                        var v59 uint32 = 435
                                                        _ = v59
                                                    }
                                                    switch fi41 {
                                                    case 873:
                                                        v33 += -2.25
                                                    case 1:
                                                        v44 = v44
                                                    case 874:
                                                        v35 = "foo bar"
                                                    }
                                                    if true {
                                                        if false {
                                                            v4 = (v4 * v4)
                                                            v60 := GenS8{F0: 42, F1: 255, F2: v1}
                                                            switch fi41 {
                                                            case 32768:
                                                                if true {
                                                                    fmt.Println(v33, v1)
                                                                    var v61 uint16 = (v1 & (v1 & 82))
                                                                    var v62 float64 = ((v33 + float64(3.14159)) - (v4 + v40))
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    var v63 int = 42
                                                                    _ = v63
                                                                }
                                                                v35 = ("\n\t" + "foo bar")
                                                            case 352:
                                                                switch fi49 {
                                                                case -183:
                                                                    v5 = append(v5, 1)
                                                                case 127:
                                                                    v33 -= (3.14159 * (v40 * ((float64((1.5 - 1e10)) / v33) / (0.0 * v40))))
                                                                }
                                                            }
                                                            _ = v60
                                                        } else {
                                                            var v64 int16 = -752
                                                            var v65 int = fi41
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    } else {
                                                        if true {
                                                            var v66 int8 = 42
                                                            _ = v66
                                                        } else {
                                                            if true {
                                                                v35 = v6
                                                                var v67 float32 = float32(v40)
                                                                _ = v67
                                                            }
                                                        }
                                                        v38 += (v38 << ((1) & 63))
                                                    }
                                                } else {
                                                    var v68 uint8 = 1
                                                    if (v31 > v31) {
                                                        var v69 int32 = (^(v38))
                                                        v35 = "a"
                                                        _ = v69
                                                    } else {
                                                        var v70 int8 = v31
                                                        v31 += (v70 * 0)
                                                        _ = v70
                                                    }
                                                    _ = v68
                                                }
                                                v71 := GenS9{F0: (fi49 ^ (fi49 / (((fi49 + (^(fi41)))) | 1))), F1: 1, F2: 482}
                                                if false {
                                                    v3 += 434
                                                    var v72 int = fi41
                                                    _ = v72
                                                } else {
                                                    v2 = "foo bar"
                                                }
                                                _ = v71
                                            } else {
                                                switch 667 {
                                                case 695:
                                                    v4 -= v4
                                                case 4294967295:
                                                    var v73 bool = (744 != 65535)
                                                    _ = v73
                                                default:
                                                    if ((v38 + v38) < v38) {
                                                        v38 = 42
                                                        var v74 uint16 = (v1 << ((4294967295) & 63))
                                                        _ = v74
                                                    }
                                                }
                                                switch -2147483648 {
                                                case 0:
                                                    v3 = (v38 << ((1) & 63))
                                                default:
                                                    fmt.Println(v40, v4)
                                                }
                                            }
                                        }
                                    case "":
                                        if (v3 >= -1) {
                                            v2 = (v6 + v35)
                                        }
                                    }
                                } else {
                                    fmt.Println(v4, v33, v4)
                                }
                                if (!false) {
                                    v35 = "hello"
                                }
                                fmt.Println(v31, v6, v2)
                                _ = fi49
                            }
                        }
                    }
                    _ = v44
                } else {
                    switch fi41 {
                    case 4294967295:
                        if (v1 > v1) {
                            var v75 uint = 255
                            v6 = (((v2 + v35) + v35) + v6)
                            for fi76 := 0; fi76 < 1; fi76++ {
                                if ((!false) && (-32768 == 1)) {
                                    switch "" {
                                    case "a":
                                        fmt.Println(v6, v33)
                                    case "hello":
                                        var v77 uint8 = 42
                                        v38 -= (v3 & (v38 / ((genFn0()) | 1)))
                                        var v78 bool = ("a" == v6)
                                        _ = v77
                                        _ = v78
                                    case "\n\t":
                                        v3 = genFn0()
                                    default:
                                        v1 = (v1 & 2)
                                    }
                                    if false {
                                        switch (fi76 / (((^(fi76))) | 1)) {
                                        case 636:
                                            v31 = v31
                                        }
                                        var v79 uint64 = 18446744073709551615
                                        fmt.Println(v79)
                                        _ = v79
                                    } else {
                                        if (2147483648 == 65535) {
                                            var v80 int = (fi41 >> (((v75 + (v75 * (v75 ^ (v75 % (((v75 ^ 1)) | 1)))))) & 63))
                                            switch v80 {
                                            case 4294967295:
                                                var v81 int32 = v38
                                                var v82 float64 = ((v4 - ((float64(0.1) + float64(-2.25)) - (float64(1.5) + (-2.25 / v33)))) / v40)
                                                _ = v81
                                                _ = v82
                                            case 42:
                                                v33 = float64(((-2.25 + 1e10) / (0.0 + (-2.25 / ((1e10 - -2.25) / 0.1)))))
                                            default:
                                                if (v1 >= v1) {
                                                    v83 := GenS2{F0: 286}
                                                    v84 := GenS5{F0: (v31 - v31), F1: v31, F2: 32767}
                                                    v75 = (v75 + v75)
                                                    _ = v83
                                                    _ = v84
                                                }
                                            }
                                            _ = v80
                                        } else {
                                            fmt.Println(v75)
                                            switch "hello" {
                                            case "\n\t":
                                                _ = len(v5)
                                            default:
                                                v85 := GenS1{F0: (v1 - (v1 & v1))}
                                                _ = v85
                                            }
                                            var v86 uint32 = 2
                                            _ = v86
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = 114
                                        }
                                        if ((0.0 - 0.1) <= 0.0) {
                                            v87 := GenS8{F0: 0, F1: 1, F2: (v1 / ((v1) | 1))}
                                            var v88 int64 = 42
                                            _ = v87
                                            _ = v88
                                        }
                                    }
                                } else {
                                    var v89 string = "a"
                                    v90 := GenS10{F0: v75, F1: v31, F2: -9223372036854775808}
                                    _ = v89
                                    _ = v90
                                }
                                switch fi41 {
                                case 12:
                                    v91 := GenS7{F0: ((fi41 & (fi76 << ((400) & 63))) != fi76)}
                                    _ = v91
                                }
                                var v92 uint = 255
                                _ = fi76
                                _ = v92
                            }
                            _ = v75
                        }
                    case -887:
                        var v93 string = "foo bar"
                        _ = v93
                    }
                    v1 = v1
                }
                v5 = append(v5, 0)
                _ = fi41
            }
        case 42:
            var v94 string = v6
            _ = v94
        }
        var v95 string = "foo bar"
        _ = v40
        _ = v95
    }
    v3 = 0
    var v96 float64 = v33
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v96
}
