package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint16; F1 string; F2 uint }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0(p0 int, p1 int16) int16 {
    p1 += -1
    p0 += (p0 << ((255) & 63))
    v5 := GenS1{F0: 18446744073709551615, F1: true}
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1() string {
    var v6 uint16 = 0
    var v7 float32 = 1e10
    _ = v6
    _ = v7
    return ""
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 682
    var v3 bool = true
    var v4 uint = 396
    v3 = false
    var v8 int16 = -32768
    v4 = 1
    var v9 uint8 = 0
    if true {
        v10 := []int{126, 9223372036854775807, 4294967295}
        if false {
            if v3 {
                switch 0 {
                case 325:
                    v2 = 437
                default:
                    v3 = (v4 < (v4 + 1))
                }
                if (true && v3) {
                    if (0 > 42) {
                        v11 := GenS1{F0: (v4 >> (((v4 / (((v4 | 18446744073709551615)) | 1))) & 63)), F1: v3}
                        switch "hello" {
                        case "a":
                            v9 = v9
                        default:
                            switch ("a" + "\n\t") {
                            case "foo bar":
                                for fi12 := 0; fi12 < 16; fi12++ {
                                    _ = len(v1)
                                    v4 -= v4
                                    _ = fi12
                                }
                            case "a":
                                v13 := GenS1{F0: 65535, F1: v3}
                                _ = v13
                            case "hello":
                                delete(v1, 1)
                            }
                        }
                        fmt.Println(v2)
                        _ = v11
                    }
                    fmt.Println(v9, v4, v4)
                } else {
                    var v14 string = (("a" + ("foo bar" + ("foo bar" + ("" + ("foo bar" + ("foo bar" + "foo bar")))))) + ("foo bar" + "a"))
                    _ = v14
                }
                for fi15 := 0; fi15 < 10; fi15++ {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v8 += v8
                        delete(v1, fi16)
                        _ = fi16
                    }
                    v4 = v4
                    _ = fi15
                }
            } else {
                var v17 float64 = 0.1
                var v18 uint8 = (v9 + (v9 + 42))
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int = 2147483647
                    v21 := GenS1{F0: (v4 / (((v4 / ((65535) | 1))) | 1)), F1: ((400 >= (v2 ^ v2)) || (false || true))}
                    if true {
                        v3 = v3
                        v8 = (-(v8))
                    } else {
                        fmt.Println(v18, v4)
                        var v22 int8 = 100
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v17
                _ = v18
            }
            var v23 float32 = (1.5 * 0.0)
            if v3 {
                switch "a" {
                case "foo bar":
                    v8 = (v8 ^ 1000)
                case "\n\t":
                    switch "foo bar" {
                    case "foo bar":
                        var v24 bool = (!true)
                        _ = v24
                    case "hello":
                        var v25 float64 = 1.5
                        _ = v25
                    case "\n\t":
                        if true {
                            v3 = (42 != 18446744073709551615)
                            if len(v10) > 0 {
                                v10[0] = 32768
                            }
                            _ = len(v1)
                        }
                    }
                default:
                    if v3 {
                        v23 += (v23 - 1.5)
                        if true {
                            for fi26 := 0; fi26 < 8; fi26++ {
                                _ = len(v10)
                                _ = fi26
                            }
                            var v27 uint8 = v9
                            var v28 string = "hello"
                            _ = v27
                            _ = v28
                        }
                    } else {
                        switch "" {
                        case "":
                            if v3 {
                                for fi29 := 0; fi29 < 7; fi29++ {
                                    delete(v1, (fi29 >> (((v4 % ((v4) | 1))) & 63)))
                                    _ = fi29
                                }
                            }
                        case "a":
                            for fi30 := 0; fi30 < 15; fi30++ {
                                var v31 uint16 = 32767
                                var v32 int = fi30
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        for fi33 := 0; fi33 < 4; fi33++ {
                            v34 := GenS1{F0: v4, F1: true}
                            switch fi33 {
                            case 420:
                                v4 += (v4 | v4)
                            case 127:
                                v9 = v9
                            case 42:
                                var v35 uint64 = 4294967295
                                _ = v35
                            default:
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                            _ = fi33
                            _ = v34
                        }
                    }
                }
                v36 := GenS1{F0: (v4 % (((v4 + v4)) | 1)), F1: (v3 || (v3 && v3))}
                _ = v36
            } else {
                fmt.Println(v8, v3, v23)
            }
            _ = v23
        }
        v37 := GenS2{F0: 106, F1: -2147483648}
        _ = v10
        _ = v37
    } else {
        v4 += 291
        if v3 {
            var v38 string = (("\n\t" + ("a" + ("\n\t" + "hello"))) + "hello")
            var v39 int16 = -933
            fmt.Println(v9)
            _ = v38
            _ = v39
        } else {
            if false {
                delete(v1, 417)
                var v40 int16 = v8
                _ = v40
            }
            var v41 uint64 = 4294967295
            fmt.Println(v4)
            _ = v41
        }
        if (!false) {
            var v42 uint = v4
            fmt.Println(v4, v3)
            var v43 int16 = v8
            _ = v42
            _ = v43
        } else {
            switch int(v2) {
            case 32768:
                if v3 {
                    var v44 int32 = 71
                    _ = v44
                } else {
                    var v45 float32 = -2.25
                    v46 := GenS3{F0: 1, F1: v4, F2: 42}
                    var v47 string = "hello"
                    _ = v45
                    _ = v46
                    _ = v47
                }
            case 128:
                v2 = v2
            default:
                _ = len(v1)
            }
        }
    }
    for fi48 := 0; fi48 < 20; fi48++ {
        v9 = (v9 ^ (v9 + v9))
        var v49 int64 = (v2 | (v2 ^ v2))
        _ = fi48
        _ = v49
    }
    v9 -= v9
    v50 := v1[4294967295]
    for fi51 := 0; fi51 < 18; fi51++ {
        switch v50 {
        case "":
            if false {
                v52 := GenS4{F0: 4294967295, F1: (v4 ^ v4)}
                var v53 int16 = (v8 * (v8 << (((v4 | 109)) & 63)))
                _ = v52
                _ = v53
            } else {
                v50 = v50
                for fi54 := 0; fi54 < 17; fi54++ {
                    fmt.Println(v8, v4)
                    var v55 uint8 = v9
                    if ((float64(0.0) > 0.1) && ((0.0 - (3.14159 + float64((0.1 + 1.5)))) != 3.14159)) {
                        fmt.Println(v2, v2)
                        v56 := GenS3{F0: 42, F1: (v4 ^ (v4 | (v4 + (v4 - v4)))), F2: 4294967295}
                        fmt.Println(v8, v55)
                        var v57 uint16 = 926
                        _ = v56
                        _ = v57
                    } else {
                        var v58 int32 = 1
                        _ = v58
                    }
                    _ = fi54
                    _ = v55
                }
            }
        case "a":
            var v59 int16 = (v8 % (((-(v8))) | 1))
            _ = v59
        }
        if v3 {
            v60 := GenS5{F0: 32767, F1: v50, F2: v4}
            var v61 uint32 = 0
            _ = v60
            _ = v61
        } else {
            var v62 int16 = v8
            var v63 bool = ((0 <= 2) || (v4 != 4294967295))
            switch ("a" + "foo bar") {
            case "hello":
                var v64 uint8 = 255
                _ = v64
            case "a":
                for fi65 := 0; fi65 < 14; fi65++ {
                    if v3 {
                        v9 = v9
                        v63 = true
                        var v66 uint16 = 255
                        _ = v66
                    }
                    _ = fi65
                }
            }
            _ = v62
            _ = v63
        }
        _ = fi51
    }
    if (0 != 117) {
        v2 += (v2 >> ((v4) & 63))
        switch 127 {
        case 2147483647:
            _ = len(v1)
        case 4294967295:
            if true {
                v67 := GenS3{F0: v2, F1: v4, F2: 0}
                v68 := GenS1{F0: v4, F1: (127 <= 4294967295)}
                var v69 int16 = int16(v9)
                _ = v67
                _ = v68
                _ = v69
            } else {
                v70 := GenS3{F0: (v2 >> (((v4 & v4)) & 63)), F1: (v4 | (v4 | v4)), F2: 2}
                if (("hello" + v50) < "\n\t") {
                    switch v50 {
                    case "\n\t":
                        if true {
                            if v3 {
                                v9 += 255
                                v4 -= (v4 & (^(v4)))
                                if ((v9 * v9) <= (v9 - (v9 + (v9 * (v9 ^ (v9 << (((v4 * v4)) & 63))))))) {
                                    v71 := []int{32767}
                                    v9 = (v9 + 207)
                                    for fi72 := 0; fi72 < 17; fi72++ {
                                        var v73 float32 = -2.25
                                        for fi74 := 0; fi74 < 16; fi74++ {
                                            v2 = 0
                                            _ = fi74
                                        }
                                        v50 = v50
                                        _ = fi72
                                        _ = v73
                                    }
                                    _ = v71
                                }
                            }
                            var v75 uint64 = 42
                            v2 += (v2 ^ (v2 / ((v2) | 1)))
                            _ = v75
                        }
                    case "hello":
                        v76 := GenS3{F0: v2, F1: v4, F2: 18446744073709551615}
                        _ = v76
                    case "a":
                        if false {
                            v8 = (v8 * 0)
                            v77 := []int{128, 32768}
                            _ = v77
                        }
                    }
                    v8 -= v8
                } else {
                    for fi78 := 0; fi78 < 17; fi78++ {
                        if v3 {
                            if (3.14159 != 1e10) {
                                var v79 int = fi78
                                v1[int(v8)] = v50
                                _ = v79
                            }
                            switch fi78 {
                            case 42:
                                v2 -= v2
                            }
                            v80 := GenS6{F0: 42, F1: (v9 * 0)}
                            _ = v80
                        } else {
                            v81 := GenS7{F0: 1, F1: 1}
                            _ = v81
                        }
                        v82 := GenS4{F0: 655, F1: (v4 / ((v4) | 1))}
                        switch ("foo bar" + ("a" + v50)) {
                        case "foo bar":
                            v50 = ("a" + v50)
                        default:
                            fmt.Println(v4, v3)
                        }
                        _ = fi78
                        _ = v82
                    }
                    switch -11 {
                    case -339:
                        if v3 {
                            v83 := GenS6{F0: 0, F1: (^(v9))}
                            _ = v83
                        }
                    case 9223372036854775807:
                        fmt.Println(v50)
                    case 1:
                        v84 := GenS6{F0: -31, F1: 255}
                        _ = v84
                    }
                    for fi85 := 0; fi85 < 14; fi85++ {
                        for fi86 := 0; fi86 < 7; fi86++ {
                            if (!(v3 || true)) {
                                var v87 uint8 = (v9 >> (((v4 * (v4 >> ((590) & 63)))) & 63))
                                v87 += v87
                                switch (fi86 ^ fi85) {
                                case -9223372036854775808:
                                    v87 += (v87 >> (((^(v4))) & 63))
                                case 32767:
                                    v88 := GenS4{F0: 677, F1: (^(v4))}
                                    _ = v88
                                case -2147483648:
                                    v3 = (0 >= 42)
                                }
                                _ = v87
                            } else {
                                v89 := GenS8{F0: (^(v9)), F1: 42}
                                v9 -= 75
                                _ = v89
                            }
                            var v90 float32 = (float32(-2.25) * 0.0)
                            _ = fi86
                            _ = v90
                        }
                        v91 := []int{(-(fi85)), (fi85 % ((fi85) | 1)), int(v4), fi85}
                        v2 = -506
                        _ = fi85
                        _ = v91
                    }
                }
                _ = v70
            }
        default:
            v92 := v1[-86]
            _ = v92
        }
        v3 = v3
    } else {
        v93 := []int{-9223372036854775808, 2147483647}
        v50 = "\n\t"
        _ = v93
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v8
    _ = v9
}
