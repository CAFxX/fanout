package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }
type GenS5 struct { F0 int16; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-1, 9223372036854775807, -281}
    v3 := make(map[string]string)
    var v4 uint16 = 65535
    var v5 string = ((("hello" + "a") + "a") + "")
    delete(v1, "a")
    v6 := GenS1{F0: 255}
    var v7 string = (((v5 + (v5 + "a")) + "\n\t") + ((("a" + "hello") + "") + v5))
    for fi8 := 0; fi8 < 6; fi8++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = (v5 + (((((v7 + (v7 + ((v5 + v5) + ((v7 + v5) + v7)))) + v5) + "a") + "foo bar") + v7))
        _ = fi8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint16 = 42
        fmt.Println(v5)
        v11 := v1[v5]
        _ = fi9
        _ = v10
        _ = v11
    }
    if (-32768 <= 1) {
        if false {
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := GenS2{F0: 1}
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 float32 = ((3.14159 * -2.25) * (float32(1.5) * 1e10))
                    var v16 uint16 = v4
                    v7 = "\n\t"
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 uint8 = 128
                _ = fi12
                _ = v13
                _ = v17
            }
            v18 := v1[v7]
            _ = v18
        } else {
            fmt.Println(v5)
            v4 -= 896
            var v19 int64 = 42
            _ = v19
        }
        v4 = (v4 + (v4 / (((v4 | (v4 >> ((42) & 63)))) | 1)))
        switch 128 {
        case 1:
            if true {
                var v20 int8 = -1
                if false {
                    var v21 uint16 = (v4 << ((42) & 63))
                    _ = v21
                } else {
                    switch 4294967295 {
                    case 128:
                        _ = len(v2)
                    case 40:
                        v20 = v20
                    default:
                        v7 = ("" + v5)
                    }
                    v4 += 95
                }
                switch 127 {
                case -9223372036854775808:
                    if (!(1 <= v4)) {
                        v7 = v5
                        fmt.Println(v7, v4)
                    } else {
                        fmt.Println(v5)
                        _ = len(v1)
                    }
                }
                _ = v20
            } else {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v7 = "\n\t"
                    var v23 int16 = 0
                    if (!(!(!(0 <= 1)))) {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            if ((!false) || true) {
                                var v25 int8 = -128
                                v25 = v25
                                v25 -= 1
                                _ = v25
                            } else {
                                v23 = v23
                            }
                            if (v4 != v4) {
                                var v26 string = v7
                                if false {
                                    if (((0 < -128) || false) && (true || ((((false || false) && true) && true) && (false && (false || ((535 == 255) || (fi24 <= -230))))))) {
                                        if ((1 == 131) && (0.0 == 1e10)) {
                                            switch v5 {
                                            case "a":
                                                v27 := GenS3{F0: 18446744073709551615}
                                                _ = v27
                                            case "\n\t":
                                                v28 := GenS3{F0: 116}
                                                if true {
                                                    var v29 uint16 = v4
                                                    if true {
                                                        var v30 int16 = v23
                                                        _ = v30
                                                    }
                                                    _ = len(v1)
                                                    _ = v29
                                                }
                                                if ((-1 == 294) || (3.14159 <= (1e10 / 1.5))) {
                                                    if (-813 <= 1000) {
                                                        v5 = ("a" + ("\n\t" + v5))
                                                        var v31 uint8 = 0
                                                        switch fi24 {
                                                        case 9223372036854775807:
                                                            v32 := GenS1{F0: v31}
                                                            _ = v32
                                                        }
                                                        _ = v31
                                                    }
                                                }
                                                _ = v28
                                            case "hello":
                                                if (float64((0.1 / float32((0.1 - -2.25)))) > 0.1) {
                                                    var v33 string = v26
                                                    if false {
                                                        delete(v1, "hello")
                                                        fmt.Println(v5, v33)
                                                    } else {
                                                        if (true || true) {
                                                            switch v26 {
                                                            case "foo bar":
                                                                var v34 uint = 42
                                                                _ = v34
                                                            case "":
                                                                v35 := GenS1{F0: 127}
                                                                _ = v35
                                                            case "hello":
                                                                var v36 uint8 = 255
                                                                _ = v36
                                                            default:
                                                                if true {
                                                                    delete(v1, v26)
                                                                    var v37 int64 = 0
                                                                    switch ("" + "hello") {
                                                                    case "\n\t":
                                                                        fmt.Println(v4, fi22, v26)
                                                                    case "hello":
                                                                        v37 -= v37
                                                                    default:
                                                                        v7 = v5
                                                                    }
                                                                    _ = v37
                                                                }
                                                            }
                                                            var v38 string = ("a" + ("\n\t" + v26))
                                                            v26 = ("foo bar" + "a")
                                                            var v39 uint = 42
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                    }
                                                    var v40 int = fi24
                                                    switch fi22 {
                                                    case -270:
                                                        if false {
                                                            if (!false) {
                                                                v23 += v23
                                                            } else {
                                                                v40 = fi22
                                                                var v41 uint16 = 1
                                                                var v42 float64 = (0.1 + 3.14159)
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            var v43 float64 = float64(float32(0.0))
                                                            _ = v43
                                                        }
                                                    case -2147483648:
                                                        switch v40 {
                                                        case 9223372036854775807:
                                                            if false {
                                                                var v44 int16 = (^(v23))
                                                                if true {
                                                                    v7 = "a"
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (fi24 % ((v40) | 1))
                                                                    }
                                                                }
                                                                var v45 int = 180
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                var v46 int16 = -601
                                                                _ = v46
                                                            }
                                                        case 127:
                                                            switch v40 {
                                                            case -2147483648:
                                                                if ((520 >= 2) && (!(false && (1 > -1)))) {
                                                                    var v47 uint8 = 128
                                                                    v26 = ((v26 + (v5 + v33)) + "\n\t")
                                                                    v33 = v33
                                                                    fmt.Println(v26, v23)
                                                                    _ = v47
                                                                } else {
                                                                    _ = len(v3)
                                                                    v48 := GenS4{F0: v4, F1: -2147483648, F2: 1}
                                                                    v40 = (fi24 + 127)
                                                                    _ = v48
                                                                }
                                                            case 32767:
                                                                var v49 uint64 = 0
                                                                _ = v49
                                                            case 9223372036854775807:
                                                                if len(v2) > 0 {
                                                                    v2[0] = (v40 | v40)
                                                                }
                                                            default:
                                                                var v50 uint = 697
                                                                _ = v50
                                                            }
                                                        case 128:
                                                            v40 = (fi24 - -554)
                                                        }
                                                        if (2 <= 68) {
                                                            v5 = v5
                                                            v40 += 4294967295
                                                            if (v7 < "hello") {
                                                                v51 := GenS5{F0: (-(v23)), F1: 90, F2: 4294967295}
                                                                v4 = v4
                                                                _ = v51
                                                            }
                                                        } else {
                                                            if (!false) {
                                                                if (127 == 100) {
                                                                    var v52 bool = ((1 > 690) || false)
                                                                    var v53 bool = v52
                                                                    v26 = (((v5 + (("\n\t" + ("" + ((("a" + (v7 + "\n\t")) + "") + ((("foo bar" + ("foo bar" + v26)) + ((("foo bar" + v7) + v26) + "\n\t")) + "")))) + (v5 + "foo bar"))) + ((v7 + (v5 + "hello")) + "a")) + ("\n\t" + ("foo bar" + v7)))
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                                v7 = "a"
                                                            }
                                                            var v54 bool = (v23 > (v23 << ((65535) & 63)))
                                                            fmt.Println(v26)
                                                            _ = v54
                                                        }
                                                        var v55 uint64 = 18446744073709551615
                                                        _ = v55
                                                    }
                                                    if (!(((4294967295 == 18446744073709551615) && false) || (v4 == v4))) {
                                                        if true {
                                                            switch fi24 {
                                                            case -519:
                                                                var v56 int32 = 802
                                                                _ = v56
                                                            case -15:
                                                                v26 = "a"
                                                            }
                                                            v4 += (v4 & (^(v4)))
                                                            var v57 uint32 = 65535
                                                            _ = len(v2)
                                                            if false {
                                                                if false {
                                                                    var v58 string = v26
                                                                    v2 = append(v2, fi22)
                                                                    _ = v58
                                                                } else {
                                                                    v1[v5] = (^(v57))
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        _ = len(v2)
                                                                        var v59 int = (v40 / ((v40) | 1))
                                                                        _ = v59
                                                                    }
                                                                    v33 = v26
                                                                }
                                                                var v60 bool = false
                                                                _ = v60
                                                            }
                                                            _ = v57
                                                        }
                                                    }
                                                    v4 -= v4
                                                    _ = v33
                                                    _ = v40
                                                }
                                            }
                                            _ = len(v2)
                                        }
                                        var v61 string = ("a" + "hello")
                                        v61 = v5
                                        _ = v61
                                    }
                                }
                                _ = len(v1)
                                var v62 string = (v5 + ((v7 + (((("foo bar" + v5) + ("a" + "a")) + v26) + "a")) + v26))
                                v4 = v4
                                var v63 uint = 0
                                switch v62 {
                                case "foo bar":
                                    v62 = (v62 + "a")
                                    var v64 int16 = (^(v23))
                                    v65 := GenS5{F0: v64, F1: -550, F2: 185}
                                    _ = v64
                                    _ = v65
                                default:
                                    v23 -= (v23 - (v23 << (((^(v63))) & 63)))
                                }
                                fmt.Println(fi22, v4)
                                _ = v26
                                _ = v62
                                _ = v63
                            } else {
                                var v66 float64 = float64(-2.25)
                                v67 := GenS1{F0: 32}
                                if (-2.25 <= 3.14159) {
                                    v1["a"] = 2147483647
                                }
                                var v68 float64 = v66
                                var v69 uint = 65535
                                _ = v66
                                _ = v67
                                _ = v68
                                _ = v69
                            }
                            if true {
                                var v70 uint16 = (v4 | v4)
                                _ = v70
                            } else {
                                var v71 uint16 = (v4 * (v4 >> ((42) & 63)))
                                var v72 bool = ((1e10 * (3.14159 * 1e10)) < 0.1)
                                v73 := GenS4{F0: 2, F1: 482, F2: 2147483647}
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            _ = fi24
                        }
                    } else {
                        if (v23 != -220) {
                            var v74 bool = (128 == fi22)
                            if true {
                                var v75 float32 = (((-2.25 / 3.14159) - 0.0) + float32(((0.0 - 0.0) - 1e10)))
                                if (255 == 146) {
                                    v75 += float32((3.14159 / 1e10))
                                    var v76 int = fi22
                                    var v77 int32 = 0
                                    _ = v76
                                    _ = v77
                                } else {
                                    v7 = v7
                                }
                                _ = v75
                            } else {
                                switch fi22 {
                                case 127:
                                    var v78 float32 = float32((0.0 + 1e10))
                                    _ = v78
                                }
                                v79 := GenS6{F0: 891}
                                v1[v7] = 2
                                _ = v79
                            }
                            _ = v74
                        }
                    }
                    _ = fi22
                    _ = v23
                }
            }
        case -9223372036854775808:
            fmt.Println(v7, v7)
        case -2147483648:
            if (v4 >= v4) {
                v4 -= (v4 & 32767)
                v80 := GenS7{F0: -99, F1: 128}
                _ = v80
            }
        }
    } else {
        v4 = 42
        var v81 uint = 0
        _ = v81
    }
    for fi82 := 0; fi82 < 6; fi82++ {
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
        _ = fi82
    }
    var v83 uint8 = 42
    fmt.Println(v5, v4)
    v5 = ((("foo bar" + (v5 + "hello")) + v5) + v7)
    for fi84 := 0; fi84 < 11; fi84++ {
        var v85 int = (-(fi84))
        _ = fi84
        _ = v85
    }
    fmt.Println(v4, v7, v4)
    v86 := GenS7{F0: 69, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v83
    _ = v86
}
