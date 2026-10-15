package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int32; F1 int16; F2 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 uint16 }
func genFn0() uint32 {
    var v7 uint8 = 255
    v8 := []int{}
    for fi9 := 0; fi9 < 17; fi9++ {
        v7 = (v7 % (((v7 - v7)) | 1))
        _ = len(v8)
        v10 := GenS1{F0: 2147483647, F1: 1}
        _ = fi9
        _ = v10
    }
    v7 = (^(v7))
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint32, p1 uint, p2 int) string {
    var v11 uint16 = uint16(p1)
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(p0)
        _ = fi12
    }
    var v13 uint64 = 2
    v14 := []int{32767, p2}
    var v15 uint16 = (v11 / ((v11) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    return "a"
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 42
    var v3 int16 = -1
    var v4 bool = false
    v5 := make(map[int]string)
    var v6 int8 = 127
    v16 := []int{-513, -105, -91, 4294967295}
    var v17 int32 = 42
    for fi18 := 0; fi18 < 6; fi18++ {
        if v4 {
            var v19 uint64 = (v2 ^ v2)
            fmt.Println(v6, v2)
            var v20 int16 = -1
            _ = v19
            _ = v20
        }
        _ = fi18
    }
    if v4 {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 bool = ((0 == 1) || (true && v4))
            for fi23 := 0; fi23 < 19; fi23++ {
                switch "\n\t" {
                case "\n\t":
                    v3 -= v3
                case "foo bar":
                    if v22 {
                        fmt.Println(v3)
                    }
                case "hello":
                    var v24 int64 = -9223372036854775808
                    _ = v24
                }
                v25 := GenS1{F0: (v17 ^ int32(v2)), F1: (-(v17))}
                _ = fi23
                _ = v25
            }
            if (!(969 != genFn0())) {
                var v26 float32 = float32(float64((1.5 - (1e10 + 1.5))))
                var v27 uint = 255
                for fi28 := 0; fi28 < 2; fi28++ {
                    if (3.14159 != 1.5) {
                        if v4 {
                            fmt.Println(fi28, v4, v22)
                            v6 = v6
                        } else {
                            var v29 float32 = ((((1e10 / (((float32((float64((1.5 / v26)) - 0.0)) * (v26 + -2.25)) - v26) - (float32(float64((v26 / 1.0))) + (-2.25 - v26)))) * (v26 / -2.25)) * -2.25) * (((1e10 + 1e10) + v26) / -2.25))
                            var v30 float32 = ((-2.25 - v26) + 1e10)
                            _ = v29
                            _ = v30
                        }
                        if true {
                            fmt.Println(v27)
                            v27 -= 0
                        } else {
                            fmt.Println(v22, v6, v4)
                            delete(v1, fi21)
                            v6 = v6
                        }
                        fmt.Println(v22)
                    } else {
                        v16 = append(v16, (fi21 * (fi28 - fi21)))
                        switch ("" + (((("" + "hello") + "") + "") + "hello")) {
                        case "hello":
                            if false {
                                if len(v16) > 0 {
                                    v16[0] = fi28
                                }
                                fmt.Println(v4, v2, v4)
                                if v4 {
                                    fmt.Println(v4, v26)
                                    fmt.Println(v6, v26, v2)
                                } else {
                                    var v31 string = ""
                                    if len(v16) > 0 {
                                        v16[0] = 9223372036854775807
                                    }
                                    _ = v31
                                }
                            }
                        case "\n\t":
                            if v4 {
                                v3 = v3
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            } else {
                                if ((643 >= 32767) || true) {
                                    v17 -= v17
                                    v22 = (v26 != v26)
                                    var v32 uint8 = 29
                                    var v33 uint64 = v2
                                    _ = v32
                                    _ = v33
                                }
                            }
                        case "a":
                            var v34 bool = (42 >= 255)
                            _ = v34
                        }
                    }
                    fmt.Println(v4, fi21, v2)
                    fmt.Println(v6, fi21)
                    v35 := GenS1{F0: v17, F1: 42}
                    var v36 uint = v27
                    if v4 {
                        fmt.Println(v26, v2, v26)
                        switch "hello" {
                        case "":
                            delete(v1, fi21)
                        case "a":
                            if (((fi28 / (((^(fi21))) | 1)) > (fi21 << (((^(v36))) & 63))) || false) {
                                if false {
                                    var v37 uint16 = 0
                                    _ = v37
                                }
                                var v38 int32 = (v17 ^ -2147483648)
                                v16 = append(v16, (-(fi28)))
                                var v39 uint8 = 0
                                _ = v38
                                _ = v39
                            }
                        default:
                            if v4 {
                                var v40 float64 = -2.25
                                _ = v40
                            } else {
                                if v22 {
                                    v2 += 18446744073709551615
                                    var v41 uint16 = 838
                                    if ((v2 + v2) < (v2 + 18446744073709551615)) {
                                        v3 -= -32768
                                        v36 = (v36 * (v36 - (^(v27))))
                                        v42 := GenS2{F0: 1, F1: (v2 >> ((v27) & 63))}
                                        _ = v42
                                    }
                                    _ = v41
                                }
                            }
                        }
                    }
                    delete(v1, fi21)
                    var v43 float64 = float64((v26 + (v26 - 1.5)))
                    _ = fi28
                    _ = v35
                    _ = v36
                    _ = v43
                }
                _ = v26
                _ = v27
            } else {
                var v44 uint64 = (v2 - (v2 ^ (v2 / (((v2 % ((v2) | 1))) | 1))))
                for fi45 := 0; fi45 < 14; fi45++ {
                    var v46 int32 = v17
                    _ = fi45
                    _ = v46
                }
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = v44
            }
            _ = fi21
            _ = v22
        }
        if v4 {
            if (0 > 2147483647) {
                var v47 int8 = (v6 - v6)
                for fi48 := 0; fi48 < 5; fi48++ {
                    if (!("foo bar" != "hello")) {
                        fmt.Println(v6)
                        switch "a" {
                        case "":
                            for fi49 := 0; fi49 < 2; fi49++ {
                                var v50 int16 = (^(v3))
                                var v51 float64 = 1.5
                                v6 += v47
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                        case "hello":
                            _ = len(v1)
                        case "a":
                            var v52 int32 = v17
                            _ = v52
                        default:
                            _ = len(v1)
                        }
                    }
                    var v53 uint = 484
                    var v54 int64 = 9223372036854775807
                    _ = fi48
                    _ = v53
                    _ = v54
                }
                _ = v47
            } else {
                v55 := GenS3{F0: 135, F1: 2147483648}
                _ = v55
            }
            v2 = (v2 / (((v2 + 1)) | 1))
            v2 += v2
        }
        v56 := GenS4{F0: 0, F1: "hello"}
        _ = v56
    } else {
        _ = len(v16)
        var v57 int32 = (^(v17))
        for fi58 := 0; fi58 < 11; fi58++ {
            v2 = v2
            v6 += v6
            v3 += v3
            _ = fi58
        }
        _ = v57
    }
    if (-963 <= 0) {
        switch ("\n\t" + "a") {
        case "foo bar":
            var v59 uint32 = 65535
            _ = v59
        case "a":
            switch -2147483648 {
            case 9223372036854775807:
                var v60 uint64 = 2
                _ = v60
            }
        }
        v61 := GenS3{F0: 2, F1: 573}
        if v4 {
            v3 = (v3 + 1)
            var v62 int32 = (v17 >> ((2) & 63))
            v63 := GenS2{F0: 42, F1: (v2 & (v2 / (((^(v2))) | 1)))}
            _ = v62
            _ = v63
        } else {
            switch ("\n\t" + "a") {
            case "hello":
                if (v4 || (!(false && v4))) {
                    var v64 float32 = (float32(float64(0.0)) + 0.1)
                    v1[-959] = "foo bar"
                    _ = v64
                }
            case "":
                var v65 string = "foo bar"
                _ = v65
            case "foo bar":
                for fi66 := 0; fi66 < 8; fi66++ {
                    var v67 uint16 = 2
                    v68 := GenS4{F0: (v67 | (^(v67))), F1: ""}
                    _ = fi66
                    _ = v67
                    _ = v68
                }
            default:
                v3 = (v3 | (-(v3)))
            }
        }
        _ = v61
    } else {
        if v4 {
            switch -2147483648 {
            case 32767:
                fmt.Println(v2, v4)
            }
            var v69 bool = true
            _ = v69
        }
        v6 -= 0
        v6 += (^(v6))
    }
    var v70 int8 = v6
    switch "hello" {
    case "\n\t":
        if true {
            switch ("\n\t" + "a") {
            case "hello":
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            default:
                if (!((v17 % ((v17) | 1)) < int32(v3))) {
                    if v4 {
                        v6 = (-(v6))
                    }
                } else {
                    switch "hello" {
                    case "\n\t":
                        if v4 {
                            delete(v5, -75)
                            fmt.Println(v17)
                        } else {
                            v71 := GenS5{F0: v17, F1: 255}
                            _ = v71
                        }
                    case "a":
                        v72 := GenS6{F0: v17, F1: v3, F2: 0}
                        _ = v72
                    case "":
                        v6 = (v6 + (-(v70)))
                    default:
                        var v73 uint32 = genFn0()
                        _ = v73
                    }
                    for fi74 := 0; fi74 < 2; fi74++ {
                        if (-1 == v3) {
                            v70 -= -128
                            var v75 float64 = 0.0
                            switch "a" {
                            case "\n\t":
                                v4 = v4
                            }
                            _ = v75
                        }
                        _ = fi74
                    }
                }
            }
            for fi76 := 0; fi76 < 4; fi76++ {
                if (v17 == v17) {
                    v16 = append(v16, (-(fi76)))
                }
                _ = len(v5)
                if true {
                    if v4 {
                        fmt.Println(v2, v3)
                    } else {
                        switch "hello" {
                        case "":
                            if v4 {
                                for fi77 := 0; fi77 < 14; fi77++ {
                                    var v78 uint16 = 42
                                    _ = fi77
                                    _ = v78
                                }
                            } else {
                                v79 := GenS4{F0: 0, F1: ""}
                                if v4 {
                                    if true {
                                        for fi80 := 0; fi80 < 4; fi80++ {
                                            if (fi76 != 128) {
                                                var v81 int32 = 625
                                                var v82 int16 = (v3 * v3)
                                                v83 := GenS5{F0: -2147483648, F1: 834}
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            }
                                            if v4 {
                                                fmt.Println(v70)
                                            }
                                            _ = fi80
                                        }
                                    }
                                } else {
                                    fmt.Println(v4, fi76)
                                    v84 := GenS7{F0: 42}
                                    _ = v84
                                }
                                _ = v79
                            }
                        default:
                            fmt.Println(v70, v70)
                        }
                        v3 = 1000
                    }
                    switch "a" {
                    case "":
                        var v85 uint = 150
                        _ = v85
                    case "a":
                        v86 := GenS5{F0: v17, F1: 1}
                        _ = v86
                    case "\n\t":
                        if ((!(v4 && true)) || v4) {
                            for fi87 := 0; fi87 < 17; fi87++ {
                                v88 := v1[(fi76 / ((fi87) | 1))]
                                var v89 uint32 = 2
                                _ = fi87
                                _ = v88
                                _ = v89
                            }
                        } else {
                            delete(v1, int(v17))
                            v4 = (18446744073709551615 == 938)
                            v3 = v3
                        }
                    }
                    v70 += (^(v70))
                }
                _ = fi76
            }
            v1[2147483647] = ("hello" + ("hello" + "a"))
        } else {
            v3 = 631
            var v90 uint = 42
            _ = v90
        }
    case "a":
        v3 = (v3 + (v3 / ((v3) | 1)))
    default:
        _ = len(v1)
    }
    v70 = v6
    v6 += (v6 + (v70 * v70))
    v6 = v6
    switch ("hello" + ("" + "\n\t")) {
    case "":
        v17 = v17
    default:
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    v91 := GenS8{F0: 4294967295, F1: 65535}
    v17 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v91
}
