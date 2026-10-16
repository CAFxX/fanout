package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint64; F2 int16 }
type GenS4 struct { F0 int8; F1 int64; F2 uint }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint) int {
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = (p0 % ((p0) | 1))
        _ = fi7
    }
    var v8 int = 127
    v9 := []int{int(p0), (v8 + v8)}
    _ = p0
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 int32) int {
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := []int{7, 127}
        for fi12 := 0; fi12 < 19; fi12++ {
            switch (fi10 / ((fi12) | 1)) {
            case 146:
                p0 = p0
            case 2147483647:
                p0 = (p0 - (p0 | 669))
            case 0:
                var v13 int16 = 1000
                _ = v13
            }
            var v14 float64 = (1e10 - 3.14159)
            _ = fi12
            _ = v14
        }
        p0 = p0
        _ = fi10
        _ = v11
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint16 = 327
        v17 := GenS1{F0: 256, F1: (p0 | 2147483647), F2: -2147483648}
        _ = fi15
        _ = v16
        _ = v17
    }
    p0 -= (p0 - p0)
    fmt.Println(p0)
    _ = p0
    return genFn0(0)
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    v3 := []uint32{103, 2}
    var v4 uint64 = 0
    var v5 string = ("foo bar" + "hello")
    v6 := []int{42}
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int64 = 645
            v1 = 1e10
            if (32767 > 1) {
                var v21 int = fi18
                if (((((0.0 * 0.0) + 3.14159) / 1.0) + 1e10) <= 1.5) {
                    if (v4 != (v4 << ((722) & 63))) {
                        _ = len(v2)
                        v4 -= v4
                        if true {
                            var v22 uint = 4294967295
                            var v23 float32 = -2.25
                            v24 := GenS1{F0: (v4 >> ((4294967295) & 63)), F1: 0, F2: 2147483647}
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    } else {
                        if true {
                            if false {
                                v1 += 0.0
                                if (!true) {
                                    v1 += v1
                                    var v25 uint16 = 0
                                    if (false && (!(v5 <= "\n\t"))) {
                                        if false {
                                            var v26 int = (v21 ^ genFn1(-1))
                                            fmt.Println(v5, v21)
                                            _ = v26
                                        } else {
                                            v21 = fi19
                                            v20 = -1
                                            var v27 uint = 2
                                            _ = v27
                                        }
                                        var v28 int64 = 0
                                        if true {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v1 -= 3.14159
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        } else {
                                            if true {
                                                v21 = (fi18 << ((18446744073709551615) & 63))
                                                if false {
                                                    v25 -= v25
                                                    v29 := GenS2{F0: 2}
                                                    _ = v29
                                                } else {
                                                    switch (fi18 * (fi18 + genFn1(-1))) {
                                                    case 9223372036854775807:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    default:
                                                        v30 := GenS1{F0: v4, F1: -2147483648, F2: 413}
                                                        _ = v30
                                                    }
                                                    fmt.Println(fi18)
                                                    var v31 int64 = v20
                                                    v1 = (v1 + 0.0)
                                                    _ = v31
                                                }
                                                var v32 uint8 = 2
                                                _ = v32
                                            }
                                        }
                                        if true {
                                            v33 := GenS1{F0: (v4 - 0), F1: 32, F2: -839}
                                            var v34 int16 = -221
                                            _ = v33
                                            _ = v34
                                        }
                                        _ = v28
                                    } else {
                                        v21 = (v21 | (^(fi19)))
                                        if true {
                                            fmt.Println(fi18, v20)
                                            v20 = v20
                                            v35 := GenS2{F0: 42}
                                            v36 := GenS3{F0: true, F1: 18446744073709551615, F2: 32767}
                                            switch fi18 {
                                            case 42:
                                                v21 += genFn1(2147483647)
                                            case 2147483647:
                                                if (false || (v20 != v20)) {
                                                    var v37 int32 = 1
                                                    fmt.Println(v20, v20, fi19)
                                                    switch genFn0(125) {
                                                    case -2147483648:
                                                        switch v5 {
                                                        case "":
                                                            fmt.Println(fi18)
                                                        }
                                                    case 127:
                                                        if (!(!(9223372036854775807 == (v20 ^ 9223372036854775807)))) {
                                                            v38 := v2[(v21 | (fi18 % (((fi18 | 32767)) | 1)))]
                                                            v4 = uint64(v25)
                                                            _ = v38
                                                        } else {
                                                            var v39 int16 = 32767
                                                            v5 = (v5 + "foo bar")
                                                            _ = v39
                                                        }
                                                    case 4294967295:
                                                        v20 = -1
                                                    default:
                                                        var v40 int16 = 42
                                                        _ = v40
                                                    }
                                                    _ = v37
                                                }
                                            case -637:
                                                fmt.Println(v1, fi19, fi19)
                                            default:
                                                if true {
                                                    fmt.Println(fi18, fi18)
                                                    fmt.Println(v4, v5)
                                                }
                                            }
                                            var v41 uint8 = 2
                                            _ = v35
                                            _ = v36
                                            _ = v41
                                        }
                                    }
                                    _ = v25
                                } else {
                                    if (255 < 1) {
                                        v42 := GenS4{F0: 100, F1: v20, F2: 184}
                                        if len(v3) > 0 {
                                            v3[0] = 2147483647
                                        }
                                        _ = len(v3)
                                        _ = v42
                                    } else {
                                        var v43 int8 = 1
                                        v4 -= 1
                                        _ = v43
                                    }
                                    if (42 < 2) {
                                        if false {
                                            v44 := GenS2{F0: 127}
                                            fmt.Println(v4, v5, v5)
                                            _ = v44
                                        }
                                        switch v21 {
                                        case -885:
                                            v21 = (^(fi18))
                                        case -2147483648:
                                            v5 = "\n\t"
                                        case 32768:
                                            v45 := GenS3{F0: (127 > -11), F1: (^(v4)), F2: -350}
                                            _ = v45
                                        default:
                                            var v46 int64 = (v20 % ((-181) | 1))
                                            if true {
                                                v47 := GenS4{F0: 49, F1: v46, F2: 65535}
                                                _ = v47
                                            }
                                            _ = v46
                                        }
                                    } else {
                                        var v48 int16 = -264
                                        _ = v48
                                    }
                                    fmt.Println(fi18)
                                }
                                v5 = ("a" + v5)
                                v49 := v2[v21]
                                v50 := GenS5{F0: v4, F1: 1}
                                _ = v49
                                _ = v50
                            }
                            v4 = (v4 / ((122) | 1))
                        }
                        var v51 uint64 = (v4 - v4)
                        _ = v51
                    }
                    v5 = ((("" + v5) + (v5 + ((v5 + (v5 + (v5 + (((v5 + "\n\t") + v5) + ("" + ("a" + "a")))))) + v5))) + v5)
                } else {
                    if (!true) {
                        v20 = 9223372036854775807
                        v5 = "foo bar"
                        var v52 int8 = -128
                        _ = v52
                    } else {
                        v53 := GenS2{F0: 255}
                        v54 := GenS2{F0: 42}
                        v1 = 3.14159
                        _ = len(v2)
                        var v55 int32 = 1
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    var v56 int8 = -1
                    var v57 uint8 = 128
                    _ = v56
                    _ = v57
                }
                v58 := GenS5{F0: (v4 + (^(v4))), F1: -615}
                _ = v21
                _ = v58
            } else {
                v59 := GenS3{F0: (true && true), F1: v4, F2: 0}
                var v60 int16 = 1
                fmt.Println(v5)
                _ = v59
                _ = v60
            }
            _ = fi19
            _ = v20
        }
        v4 = 847
        _ = fi18
    }
    for fi61 := 0; fi61 < 20; fi61++ {
        for fi62 := 0; fi62 < 5; fi62++ {
            v3 = append(v3, 42)
            v4 += 1
            v63 := GenS2{F0: 255}
            _ = fi62
            _ = v63
        }
        _ = fi61
    }
    fmt.Println(v1, v1, v4)
    v64 := GenS2{F0: 42}
    switch ((v5 + "") + v5) {
    case "a":
        var v65 string = v5
        _ = v65
    case "hello":
        if false {
            v66 := GenS3{F0: false, F1: (v4 | (v4 / ((v4) | 1))), F2: -853}
            _ = v66
        } else {
            v1 += 0.1
            v5 = v5
            v67 := GenS6{F0: 764}
            _ = v67
        }
    }
    var v68 uint8 = 127
    delete(v2, 273)
    if true {
        for fi69 := 0; fi69 < 11; fi69++ {
            v70 := GenS6{F0: (fi69 / (((fi69 - fi69)) | 1))}
            delete(v2, fi69)
            switch v5 {
            case "hello":
                v68 -= 2
            case "":
                for fi71 := 0; fi71 < 5; fi71++ {
                    v4 = 42
                    _ = fi71
                }
            case "\n\t":
                v72 := GenS1{F0: (v4 | (^(v4))), F1: -767, F2: 2147483647}
                _ = v72
            default:
                var v73 uint = 245
                _ = v73
            }
            _ = fi69
            _ = v70
        }
        var v74 uint16 = 42
        for fi75 := 0; fi75 < 5; fi75++ {
            fmt.Println(v1)
            switch (fi75 % ((0) | 1)) {
            case -9223372036854775808:
                switch 32768 {
                case 32767:
                    var v76 string = v5
                    _ = v76
                case 4294967295:
                    v5 = ""
                }
            case 1:
                for fi77 := 0; fi77 < 7; fi77++ {
                    var v78 uint16 = v74
                    if (!false) {
                        var v79 int16 = 42
                        if false {
                            if true {
                                v4 += (v4 | 956)
                                _ = len(v3)
                                var v80 uint = 65535
                                var v81 uint8 = (v68 + 0)
                                _ = v80
                                _ = v81
                            } else {
                                var v82 float32 = (((0.0 / (0.1 + ((float32(v1) - (1.5 / 1.0)) + (3.14159 - float32(-2.25))))) + 1.5) + float32((v1 * (v1 * float64(0.0)))))
                                if true {
                                    v4 -= 0
                                    fmt.Println(v4)
                                } else {
                                    var v83 float32 = (v82 - 0.1)
                                    _ = len(v2)
                                    var v84 uint32 = 522
                                    var v85 uint16 = (v78 & (^(v74)))
                                    v86 := GenS7{F0: 100, F1: v68}
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                }
                                v87 := GenS6{F0: (fi75 - 1)}
                                _ = v82
                                _ = v87
                            }
                            v5 = v5
                        }
                        _ = v79
                    }
                    v4 -= v4
                    v4 = (^(v4))
                    if (1 < 463) {
                        var v88 int = genFn0(18446744073709551615)
                        v4 += v4
                        _ = v88
                    } else {
                        v3 = append(v3, 2)
                    }
                    _ = fi77
                    _ = v78
                }
            case 4294967295:
                for fi89 := 0; fi89 < 12; fi89++ {
                    fmt.Println(fi89, v74)
                    var v90 float64 = ((v1 * v1) + ((float64(1e10) / 1.5) + (float64(3.14159) / (v1 * 0.1))))
                    v5 = (v5 + v5)
                    _ = fi89
                    _ = v90
                }
            default:
                for fi91 := 0; fi91 < 12; fi91++ {
                    var v92 uint32 = 4294967295
                    v5 = v5
                    v93 := GenS8{F0: v4}
                    var v94 string = v5
                    switch fi75 {
                    case -772:
                        var v95 float32 = 0.0
                        _ = v95
                    }
                    _ = fi91
                    _ = v92
                    _ = v93
                    _ = v94
                }
            }
            if false {
                if true {
                    if true {
                        fmt.Println(v4, v4, v4)
                        fmt.Println(v1, v4, v68)
                        if false {
                            v96 := GenS7{F0: 127, F1: v68}
                            _ = v96
                        }
                    }
                }
                v74 += 32767
                v97 := GenS4{F0: 127, F1: -9223372036854775808, F2: 0}
                _ = v97
            } else {
                fmt.Println(v74, v68)
            }
            _ = fi75
        }
        _ = v74
    }
    v1 -= v1
    for fi98 := 0; fi98 < 17; fi98++ {
        var v99 float32 = (1.5 * 0.1)
        _ = fi98
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v68
}
