package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 int16; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 int64 }
type GenS11 struct { F0 int64; F1 uint; F2 int32 }
func genFn0(p0 int32, p1 uint64, p2 uint) int16 {
    v6 := []int{-533}
    for fi7 := 0; fi7 < 14; fi7++ {
        if (113 != 255) {
            if (fi7 != fi7) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v8 uint32 = 1
                _ = v8
            }
        }
        p1 += p1
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1000
}
func genFn1(p0 int32) uint16 {
    v9 := []int{-2147483648, 2147483647}
    if true {
        p0 -= p0
    }
    if false {
        switch "\n\t" {
        case "":
            fmt.Println(p0)
        }
    }
    switch "foo bar" {
    case "foo bar":
        var v10 int16 = -812
        _ = v10
    case "\n\t":
        v11 := GenS1{F0: 2, F1: false}
        _ = v11
    default:
        var v12 uint16 = 42
        _ = v12
    }
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 324
    var v2 int32 = 943
    var v3 int16 = 553
    v4 := []uint32{65535, 2, 2}
    var v5 uint32 = 42
    var v13 int = 32767
    switch -617 {
    case 845:
        var v14 int64 = -1
        _ = v14
    case 128:
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS2{F0: v5}
            v17 := GenS2{F0: 42}
            _ = fi15
            _ = v16
            _ = v17
        }
    case 0:
        var v18 int8 = 127
        _ = v18
    }
    switch (v13 ^ 32768) {
    case 238:
        v19 := GenS2{F0: v5}
        _ = v19
    case 32767:
        var v20 uint = 65535
        _ = v20
    case -736:
        var v21 uint = 272
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        if true {
            v13 += 839
            if false {
                if (0.1 == 1.5) {
                    v4 = append(v4, (^(v5)))
                } else {
                    if (127 > 99) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v3 += (v3 + v3)
                            v3 -= genFn0((v2 >> ((216) & 63)), (v1 ^ (v1 + (v1 >> ((2) & 63)))), 65535)
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = (v5 - (v5 ^ 2))
                                }
                                v1 = uint64(v5)
                                v24 := GenS1{F0: (v1 ^ (v1 % ((v1) | 1))), F1: false}
                                if (true && false) {
                                    var v25 uint64 = (^(v1))
                                    if false {
                                        v26 := GenS1{F0: (^(v1)), F1: false}
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v27 uint16 = 1
                                            _ = v27
                                        case "hello":
                                            if ((1.5 + float64(1.5)) >= (((0.1 / 0.1) / 1e10) + 1e10)) {
                                                if true {
                                                    if false {
                                                        var v28 int = 689
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if (!true) {
                                                            v2 -= 2147483647
                                                            v29 := GenS2{F0: (^(v5))}
                                                            if ((("hello" < "foo bar") && true) || false) {
                                                                var v30 float64 = -2.25
                                                                _ = v30
                                                            }
                                                            _ = v29
                                                        } else {
                                                            if true {
                                                                v31 := GenS2{F0: v5}
                                                                _ = v31
                                                            }
                                                            v32 := GenS2{F0: v5}
                                                            switch (fi22 & -6) {
                                                            case -2147483648:
                                                                v33 := GenS3{F0: 127, F1: (v5 >> ((65535) & 63)), F2: (19 >= 108)}
                                                                _ = v33
                                                            case -954:
                                                                v5 = (v5 / ((v5) | 1))
                                                            default:
                                                                fmt.Println(fi23, v28)
                                                                var v34 int8 = 127
                                                                if len(v4) > 0 {
                                                                    v4[0] = 2147483647
                                                                }
                                                                var v35 float64 = 0.0
                                                                if true {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                } else {
                                                                    v35 += v35
                                                                    v36 := GenS3{F0: 127, F1: 65535, F2: true}
                                                                    _ = v36
                                                                }
                                                                fmt.Println(v5)
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                            _ = v32
                                                        }
                                                        fmt.Println(fi22)
                                                        v3 = v3
                                                        v37 := GenS2{F0: (v5 & (^(v5)))}
                                                        v28 = v28
                                                        fmt.Println(v5, v3, v5)
                                                        var v38 float32 = (1.5 * 3.14159)
                                                        _ = v28
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        _ = len(v4)
                                                        var v39 uint16 = 706
                                                        var v40 float32 = (0.1 + -2.25)
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                } else {
                                                    if true {
                                                        v4 = append(v4, v5)
                                                    }
                                                    v3 += 1000
                                                    switch (fi22 << ((270) & 63)) {
                                                    case 32768:
                                                        var v41 float32 = 1e10
                                                        _ = v41
                                                    case -1:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    case -9223372036854775808:
                                                        v3 = (v3 ^ (v3 / ((genFn0((v2 % (((v2 * (v2 * v2))) | 1)), 42, 2)) | 1)))
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v25, v1)
                                                v4 = append(v4, v5)
                                                var v42 int64 = 42
                                                if (((("hello" < "foo bar") || false) && (42 <= (v13 >> ((42) & 63)))) || false) {
                                                    fmt.Println(v25)
                                                    v43 := GenS4{F0: 1, F1: (v25 / ((133) | 1))}
                                                    if false {
                                                        fmt.Println(v1, v42, v42)
                                                        if true {
                                                            var v44 uint32 = 2
                                                            if true {
                                                                v45 := GenS5{F0: 255, F1: v3}
                                                                if len(v4) > 0 {
                                                                    v4[0] = (v44 / ((v5) | 1))
                                                                }
                                                                if false {
                                                                    v46 := GenS6{F0: "foo bar"}
                                                                    v4 = append(v4, v5)
                                                                    var v47 uint16 = 1
                                                                    _ = v46
                                                                    _ = v47
                                                                } else {
                                                                    v13 = (v13 + v13)
                                                                    if false {
                                                                        fmt.Println(v1, v2)
                                                                        fmt.Println(v2)
                                                                        v48 := GenS3{F0: 2, F1: (v44 + (v44 % ((406) | 1))), F2: ((65535 != 65535) || false)}
                                                                        _ = v48
                                                                    }
                                                                }
                                                                if false {
                                                                    if ((!((v42 / ((v42) | 1)) <= (v42 - -204))) || ((1 < 127) || ((1e10 == 3.14159) && true))) {
                                                                        var v49 uint = 42
                                                                        var v50 float32 = (3.14159 / float32(0.1))
                                                                        if false {
                                                                            var v51 int8 = 100
                                                                            _ = v51
                                                                        }
                                                                        _ = v49
                                                                        _ = v50
                                                                    } else {
                                                                        _ = len(v4)
                                                                        var v52 bool = (606 >= -930)
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (v44 * (v44 / ((2147483647) | 1)))
                                                                        }
                                                                        _ = v52
                                                                    }
                                                                    var v53 uint16 = 255
                                                                    var v54 float32 = (1.5 - 0.0)
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    _ = len(v4)
                                                                    v25 = v25
                                                                    if (!false) {
                                                                        fmt.Println(v5)
                                                                        var v55 bool = false
                                                                        _ = v55
                                                                    } else {
                                                                        var v56 bool = ((((1.5 - (float64(float32(3.14159)) + float64((float32(-2.25) + 0.1)))) * float64(0.1)) + (1.5 / float64((1.5 * (((0.0 - (1e10 - ((0.0 / 1.5) * 0.0))) * -2.25) + ((0.1 + -2.25) * (((3.14159 - (-2.25 * 3.14159)) * float32((1e10 - 1.5))) / 1.5))))))) == 0.0)
                                                                        _ = v56
                                                                    }
                                                                }
                                                                _ = v45
                                                            }
                                                            var v57 int = (^(fi23))
                                                            _ = v44
                                                            _ = v57
                                                        } else {
                                                            var v58 int64 = v42
                                                            v42 += v58
                                                            v59 := GenS2{F0: 0}
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        v25 -= 42
                                                    } else {
                                                        var v60 uint64 = 1
                                                        _ = v60
                                                    }
                                                    _ = v43
                                                } else {
                                                    var v61 int = fi22
                                                    _ = v61
                                                }
                                                v4 = append(v4, v5)
                                                _ = v42
                                            }
                                        }
                                        _ = v26
                                    }
                                    var v62 uint16 = 657
                                    _ = v25
                                    _ = v62
                                }
                                v63 := GenS6{F0: "a"}
                                _ = v24
                                _ = v63
                            }
                            if ((1e10 == (float32(float64((0.0 - (0.1 + float32(-2.25))))) - 3.14159)) && (42 <= 127)) {
                                fmt.Println(v13)
                                v64 := GenS2{F0: (v5 & v5)}
                                _ = v64
                            } else {
                                v3 = int16(v2)
                                var v65 int32 = v2
                                if len(v4) > 0 {
                                    v4[0] = 146
                                }
                                _ = v65
                            }
                            _ = fi23
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = (v5 ^ v5)
                        }
                        v13 -= (fi22 & (-(fi22)))
                    }
                    if (((^(fi22)) > -16) && false) {
                        var v66 int64 = -1
                        switch "" {
                        case "a":
                            switch v13 {
                            case -1:
                                v67 := GenS7{F0: (v2 * (v2 / (((-(v2))) | 1))), F1: v3, F2: ((("foo bar" + "a") + "\n\t") + "a")}
                                _ = v67
                            }
                        default:
                            v4 = append(v4, 636)
                        }
                        switch -9223372036854775808 {
                        case -527:
                            for fi68 := 0; fi68 < 14; fi68++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v69 := GenS5{F0: 1, F1: (v3 * v3)}
                                var v70 bool = (v13 > (fi22 * 128))
                                v71 := GenS8{F0: 100}
                                if v70 {
                                    fmt.Println(v3, v66)
                                } else {
                                    v66 += v66
                                }
                                _ = fi68
                                _ = v69
                                _ = v70
                                _ = v71
                            }
                        case 128:
                            switch "foo bar" {
                            case "\n\t":
                                if true {
                                    if false {
                                        _ = len(v4)
                                    } else {
                                        v3 += genFn0(v2, v1, 255)
                                        v1 = v1
                                        v72 := GenS9{F0: -333}
                                        _ = v72
                                    }
                                    var v73 int16 = v3
                                    var v74 uint64 = 42
                                    _ = v73
                                    _ = v74
                                }
                            case "a":
                                var v75 uint64 = (v1 | 0)
                                _ = v75
                            }
                        case 2147483647:
                            var v76 uint32 = (v5 << ((42) & 63))
                            _ = v76
                        }
                        _ = v66
                    } else {
                        if (0 != 270) {
                            v4 = append(v4, v5)
                            v4 = append(v4, v5)
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                for fi77 := 0; fi77 < 1; fi77++ {
                    fmt.Println(v1)
                    _ = fi77
                }
            }
            fmt.Println(v2, v2, v1)
        } else {
            if len(v4) > 0 {
                v4[0] = (v5 << ((0) & 63))
            }
            var v78 uint32 = v5
            var v79 int16 = v3
            _ = v78
            _ = v79
        }
        if len(v4) > 0 {
            v4[0] = v5
        }
        switch (v13 ^ fi22) {
        case 127:
            switch 0 {
            case -1:
                _ = len(v4)
            case 32767:
                var v80 uint16 = 155
                _ = v80
            }
        case -199:
            v1 = 4294967295
        default:
            if (!true) {
                v81 := GenS8{F0: -128}
                var v82 int = fi22
                _ = v81
                _ = v82
            } else {
                v83 := GenS10{F0: "foo bar", F1: 0}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v83
            }
        }
        _ = fi22
    }
    var v84 bool = (!(v13 > v13))
    var v85 int64 = 42
    v86 := GenS1{F0: v1, F1: (true && v84)}
    v87 := GenS11{F0: 9223372036854775807, F1: 0, F2: (v2 + v2)}
    var v88 int32 = v2
    if (((((float32((3.14159 * 0.0)) + ((3.14159 * (((0.1 - 1e10) * 1.5) + float32(1e10))) * float32(3.14159))) / (1e10 - 1.5)) / -2.25) * 0.0) == (1.5 / 0.1)) {
        var v89 int32 = (v2 % ((v2) | 1))
        _ = v89
    } else {
        v90 := GenS10{F0: ("\n\t" + "hello"), F1: v85}
        _ = v90
    }
    if true {
        if (!(v5 >= v5)) {
            fmt.Println(v5, v88)
        }
    }
    var v91 int = 9223372036854775807
    for fi92 := 0; fi92 < 13; fi92++ {
        var v93 uint64 = (v1 * v1)
        var v94 uint32 = (v5 ^ (v5 % ((v5) | 1)))
        _ = fi92
        _ = v93
        _ = v94
    }
    var v95 int8 = int8(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v88:", v88)
    fmt.Println("v91:", v91)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    _ = v88
    _ = v91
    _ = v95
}
