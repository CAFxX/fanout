package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 uint; F1 int32 }
func genFn0(p0 int32, p1 uint) uint16 {
    var v6 float32 = ((float32((((3.14159 / ((-2.25 + 3.14159) - 1e10)) * float64(0.0)) - 1.5)) * 1e10) + 3.14159)
    p0 += (p0 % ((p0) | 1))
    v6 -= (v6 + v6)
    v6 += -2.25
    v7 := GenS1{F0: -1, F1: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint64) uint16 {
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint16 = 557
        v10 := GenS1{F0: -962, F1: 1000}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (0 >= 255) {
        var v11 uint64 = (p0 >> ((65535) & 63))
        _ = v11
    }
    var v12 float32 = ((((3.14159 + 1e10) / float32(float64((float32(-2.25) * 3.14159)))) + 0.0) + (0.0 - 0.0))
    v13 := GenS1{F0: 616, F1: 1}
    _ = p0
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 762
    var v4 float32 = 0.1
    v5 := []uint32{2, 183}
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = len(v5)
        delete(v2, 592)
    }
    v14 := GenS1{F0: -1, F1: 0}
    fmt.Println(v4, v1, v1)
    var v15 uint64 = (v1 & v1)
    v1 = (^(v15))
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch ("foo bar" + ("hello" + ("" + (("a" + ("hello" + ((("a" + "a") + "") + (("" + "a") + ("a" + ""))))) + "\n\t")))) {
        case "\n\t":
            v3 += 695
        case "":
            _ = len(v2)
        default:
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 float32 = float32(1.5)
                if false {
                    switch int(v1) {
                    case 2147483647:
                        var v18 int8 = 42
                        _ = v18
                    case -508:
                        var v19 uint = 4294967295
                        _ = v19
                    case 431:
                        delete(v2, 2)
                    default:
                        v20 := GenS2{F0: -1, F1: 2}
                        _ = v20
                    }
                    v3 -= (v3 * 0)
                } else {
                    if (2 <= 2) {
                        var v21 int32 = -1
                        v3 = v3
                        v22 := GenS2{F0: (v21 / ((v21) | 1)), F1: 110}
                        _ = v21
                        _ = v22
                    } else {
                        if len(v5) > 0 {
                            v5[0] = 2147483647
                        }
                        var v23 int16 = 32767
                        v17 -= (v17 / v4)
                        _ = v23
                    }
                    v1 -= v1
                }
                for fi24 := 0; fi24 < 17; fi24++ {
                    v17 = v4
                    _ = fi24
                }
                _ = fi16
                _ = v17
            }
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            if (!(181 > 1)) {
                for fi26 := 0; fi26 < 19; fi26++ {
                    if (1000 != 1) {
                        var v27 bool = false
                        if v27 {
                            switch "a" {
                            case "foo bar":
                                var v28 float32 = (v4 / 0.1)
                                _ = v28
                            case "":
                                var v29 int8 = 1
                                _ = v29
                            case "\n\t":
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                            }
                            v15 = v15
                            if (!(v27 && true)) {
                                v3 += v3
                                v30 := GenS3{F0: -65, F1: 436}
                                if ((1e10 / (-2.25 - 0.1)) == float64(v4)) {
                                    v3 = 42
                                    if v27 {
                                        if len(v5) > 0 {
                                            v5[0] = 273
                                        }
                                        v5 = append(v5, 2147483647)
                                    } else {
                                        v31 := GenS4{F0: 0, F1: 65535, F2: 65535}
                                        v1 = v1
                                        _ = v31
                                    }
                                }
                                _ = v30
                            } else {
                                v32 := GenS3{F0: 104, F1: uint32(v15)}
                                var v33 int8 = 98
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v5)
                                case "a":
                                    switch "\n\t" {
                                    case "":
                                        var v34 float32 = v4
                                        _ = v34
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "a":
                                            switch (^(fi25)) {
                                            case -928:
                                                v4 -= float32(-2.25)
                                            case 42:
                                                v4 += v4
                                            case -960:
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(v1)
                                                case "foo bar":
                                                    if true {
                                                        if (v27 && (v27 || true)) {
                                                            if false {
                                                                switch (fi25 + (fi25 - (^(fi26)))) {
                                                                case 318:
                                                                    if (v1 < v1) {
                                                                        v33 -= v33
                                                                    }
                                                                default:
                                                                    fmt.Println(fi25)
                                                                }
                                                                switch 9223372036854775807 {
                                                                case -9223372036854775808:
                                                                    switch fi25 {
                                                                    case 32767:
                                                                        switch (fi25 << ((982) & 63)) {
                                                                        case 9223372036854775807:
                                                                            if v27 {
                                                                                v33 = v33
                                                                                v35 := GenS3{F0: 42, F1: 0}
                                                                                v15 += 2
                                                                                if ("a" > "a") {
                                                                                    v36 := GenS5{F0: uint8(v3)}
                                                                                    v37 := GenS5{F0: 230}
                                                                                    _ = v36
                                                                                    _ = v37
                                                                                } else {
                                                                                    if (150 >= 2) {
                                                                                        var v38 int16 = 42
                                                                                        if v27 {
                                                                                            _ = len(v5)
                                                                                            v15 += (v15 / (((^(v1))) | 1))
                                                                                        } else {
                                                                                            var v39 int32 = 2147483647
                                                                                            _ = v39
                                                                                        }
                                                                                        var v40 int16 = 1
                                                                                        _ = v38
                                                                                        _ = v40
                                                                                    }
                                                                                    v27 = v27
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = 42
                                                                                    }
                                                                                }
                                                                                v15 -= v15
                                                                                if true {
                                                                                    if ((-2.25 <= (float32(0.1) / ((v4 * (v4 + v4)) * v4))) || v27) {
                                                                                        if true {
                                                                                            var v41 float32 = 0.1
                                                                                            _ = v41
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v35
                                                                            }
                                                                        case 130:
                                                                            fmt.Println(v3, v4, fi26)
                                                                        default:
                                                                            _ = len(v5)
                                                                        }
                                                                    case -155:
                                                                        v42 := GenS5{F0: 141}
                                                                        _ = v42
                                                                    case -2147483648:
                                                                        v3 = genFn1((v1 + v1))
                                                                    }
                                                                    fmt.Println(v1, fi26)
                                                                }
                                                            } else {
                                                                v5 = append(v5, 2)
                                                            }
                                                            fmt.Println(v1)
                                                            var v43 int64 = 0
                                                            _ = v43
                                                        } else {
                                                            var v44 uint8 = 1
                                                            if false {
                                                                var v45 int16 = int16(fi26)
                                                                switch "\n\t" {
                                                                case "":
                                                                    _ = len(v5)
                                                                }
                                                                var v46 bool = ((v33 ^ v33) <= -1)
                                                                _ = v45
                                                                _ = v46
                                                            } else {
                                                                if false {
                                                                    v15 = v1
                                                                    v47 := GenS3{F0: 0, F1: 728}
                                                                    v48 := GenS6{F0: (v3 % (((v3 << ((0) & 63))) | 1)), F1: (v44 / (((v44 + (v44 - v44))) | 1))}
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                            }
                                                            var v49 int16 = -324
                                                            _ = v44
                                                            _ = v49
                                                        }
                                                        if len(v5) > 0 {
                                                            v5[0] = 2
                                                        }
                                                        _ = len(v5)
                                                    }
                                                    fmt.Println(v15, v1)
                                                    v33 = 1
                                                default:
                                                    v15 += v15
                                                }
                                            default:
                                                var v50 string = "\n\t"
                                                _ = v50
                                            }
                                        }
                                    default:
                                        if false {
                                            if (2 < 127) {
                                                if (!(!((1e10 - -2.25) > 0.0))) {
                                                    v27 = false
                                                } else {
                                                    v15 -= 306
                                                    var v51 uint64 = 1
                                                    if true {
                                                        var v52 uint = 42
                                                        _ = len(v5)
                                                        switch (fi25 - fi25) {
                                                        case 128:
                                                            var v53 int = (fi26 - (fi25 % (((fi26 & fi26)) | 1)))
                                                            _ = v53
                                                        }
                                                        _ = v52
                                                    }
                                                    _ = v51
                                                }
                                            } else {
                                                switch fi25 {
                                                case 32767:
                                                    if (2147483647 != 4294967295) {
                                                        v3 += 1
                                                        switch "a" {
                                                        case "a":
                                                            if ((^(fi25)) >= fi26) {
                                                                if (0.0 >= -2.25) {
                                                                    if false {
                                                                        v54 := GenS7{F0: v1, F1: 1}
                                                                        _ = v54
                                                                    }
                                                                    v55 := GenS6{F0: uint16(v15), F1: 0}
                                                                    if (42 != 378) {
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        case "\n\t":
                                                                            v3 = genFn1(v15)
                                                                        case "":
                                                                            v3 = 32767
                                                                        }
                                                                        fmt.Println(v33)
                                                                        v1 = (v15 % ((v15) | 1))
                                                                    }
                                                                    fmt.Println(v33)
                                                                    v2[579] = 626
                                                                    _ = v55
                                                                } else {
                                                                    if (-32768 != 1000) {
                                                                        var v56 uint16 = (v3 / ((v3) | 1))
                                                                        v5 = append(v5, 2147483647)
                                                                        _ = v56
                                                                    } else {
                                                                        if v27 {
                                                                            if v27 {
                                                                                v33 -= (v33 >> ((uint(v3)) & 63))
                                                                                if (!v27) {
                                                                                    var v57 float32 = (v4 - v4)
                                                                                    var v58 uint = 943
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(v15)
                                                                fmt.Println(v27, fi25)
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                fmt.Println(v15, v15, v4)
                                                            }
                                                        case "foo bar":
                                                            var v59 uint64 = v1
                                                            _ = v59
                                                        case "\n\t":
                                                            var v60 string = ("\n\t" + "")
                                                            _ = v60
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            v61 := GenS8{F0: 65535, F1: (fi26 ^ fi25)}
                                                            _ = v61
                                                        default:
                                                            v15 = 4294967295
                                                        }
                                                    }
                                                case 127:
                                                    fmt.Println(fi26)
                                                default:
                                                    v15 = v1
                                                }
                                                v62 := GenS9{F0: 18446744073709551615, F1: 2147483647}
                                                _ = v62
                                            }
                                            v3 += 0
                                            v27 = false
                                            var v63 int16 = 32767
                                            _ = v63
                                        }
                                    }
                                default:
                                    fmt.Println(fi25, v4, fi25)
                                }
                                _ = v32
                                _ = v33
                            }
                            v64 := GenS3{F0: 0, F1: 4294967295}
                            v3 += 583
                            if v27 {
                                var v65 int32 = -2147483648
                                var v66 uint = uint(v1)
                                _ = v65
                                _ = v66
                            } else {
                                v27 = (!true)
                                v15 = v1
                                switch 108 {
                                case 0:
                                    if v27 {
                                        if (-1 == 1000) {
                                            v4 -= (v4 - -2.25)
                                        } else {
                                            var v67 uint64 = v1
                                            v68 := GenS3{F0: 32767, F1: 2147483648}
                                            if ((0.0 - 0.1) == ((1e10 / ((0.1 / (1.5 - 3.14159)) + (0.1 + (3.14159 - float64(v4))))) * 0.1)) {
                                                v67 = 12
                                                if (v27 && ((!v27) && v27)) {
                                                    var v69 uint64 = 42
                                                    fmt.Println(v15, v69)
                                                    v27 = false
                                                    _ = v69
                                                } else {
                                                    v5 = append(v5, 1)
                                                    delete(v2, 2)
                                                    switch fi25 {
                                                    case -9223372036854775808:
                                                        var v70 int = fi25
                                                        _ = v70
                                                    case -2147483648:
                                                        v3 -= 224
                                                    case 127:
                                                        var v71 uint64 = 742
                                                        _ = v71
                                                    default:
                                                        var v72 int8 = 100
                                                        _ = v72
                                                    }
                                                }
                                            } else {
                                                if (fi26 <= -2147483648) {
                                                    fmt.Println(v27)
                                                }
                                                v73 := GenS8{F0: 65535, F1: (-(fi25))}
                                                v67 = v15
                                                _ = v73
                                            }
                                            _ = v67
                                            _ = v68
                                        }
                                        var v74 bool = v27
                                        _ = v74
                                    } else {
                                        v1 += v1
                                        if true {
                                            if ((4294967295 > 255) && false) {
                                                if v27 {
                                                    v2[978] = 65535
                                                }
                                            } else {
                                                var v75 uint = 255
                                                fmt.Println(v1, v4, v27)
                                                _ = v75
                                            }
                                            var v76 float64 = float64(v4)
                                            _ = v76
                                        } else {
                                            var v77 float64 = 1e10
                                            if false {
                                                var v78 string = ("a" + ("" + "foo bar"))
                                                _ = v78
                                            }
                                            _ = v77
                                        }
                                    }
                                }
                                switch ("\n\t" + ("foo bar" + ("" + ""))) {
                                case "foo bar":
                                    v5 = append(v5, 1)
                                }
                                switch 1 {
                                case 9223372036854775807:
                                    switch "a" {
                                    case "hello":
                                        var v79 string = "a"
                                        _ = v79
                                    case "foo bar":
                                        if v27 {
                                            var v80 int8 = 42
                                            _ = v80
                                        } else {
                                            fmt.Println(v15, v1)
                                        }
                                        var v81 int16 = -179
                                        v27 = true
                                        _ = v81
                                    default:
                                        v1 -= (v15 & (v1 * (^(v15))))
                                    }
                                case 0:
                                    var v82 int64 = -9223372036854775808
                                    _ = v82
                                default:
                                    v83 := GenS7{F0: 4294967295, F1: 42}
                                    _ = v83
                                }
                            }
                            _ = v64
                        }
                        var v84 string = ("\n\t" + "hello")
                        _ = v27
                        _ = v84
                    } else {
                        var v85 uint16 = (v3 << ((532) & 63))
                        _ = v85
                    }
                    v86 := v2[801]
                    var v87 uint64 = uint64(v86)
                    var v88 string = ""
                    var v89 uint = 42
                    _ = fi26
                    _ = v86
                    _ = v87
                    _ = v88
                    _ = v89
                }
                v3 = v3
                var v90 bool = false
                _ = v90
            }
            _ = fi25
        }
    }
    v1 -= (v15 - (v15 / ((v15) | 1)))
    delete(v2, 396)
    _ = len(v5)
    var v91 int16 = 700
    v91 = v91
    v91 += (v91 / ((v91) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
}
