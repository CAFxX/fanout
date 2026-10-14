package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint; F1 int16; F2 bool }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint8; F1 int; F2 int32 }
type GenS6 struct { F0 int16; F1 string; F2 uint }
type GenS7 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int }
func genFn0(p0 int, p1 int32, p2 uint16) int16 {
    v5 := GenS1{F0: p2, F1: (p1 >> ((uint(p2)) & 63))}
    var v6 uint64 = 358
    switch (p0 / (((p0 & (p0 - (^(p0))))) | 1)) {
    case -1000:
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(p2, p2)
            p2 -= p2
            v8 := GenS2{F0: ("foo bar" + "hello"), F1: 1}
            _ = fi7
            _ = v8
        }
    case -536:
        var v9 uint8 = 255
        _ = v9
    }
    p1 = 42
    if (-1 != -32768) {
        if (!(true && false)) {
            var v10 uint64 = 7
            _ = v10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            if (0.0 > (0.0 - float64(((0.1 - 0.1) * (((0.1 + 1.5) + -2.25) + float32(-2.25)))))) {
                if ("hello" < "") {
                    fmt.Println(v6, p1)
                    v12 := GenS2{F0: (("hello" + "a") + "\n\t"), F1: 603}
                    if (!((-592 > 32768) && (-935 < -32768))) {
                        if (127 == -66) {
                            var v13 uint32 = 1
                            p2 = (p2 + p2)
                            if true {
                                v14 := GenS2{F0: "a", F1: -625}
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    v16 := GenS3{F0: 12, F1: -1, F2: ((v13 >> ((65535) & 63)) >= v13)}
                                    _ = fi15
                                    _ = v16
                                }
                                _ = v14
                            } else {
                                if true {
                                    fmt.Println(p2, p1)
                                }
                                p0 = -1
                            }
                            _ = v13
                        } else {
                            v17 := GenS3{F0: 761, F1: 0, F2: false}
                            v18 := GenS1{F0: p2, F1: p1}
                            _ = v17
                            _ = v18
                        }
                        var v19 int16 = 922
                        if false {
                            for fi20 := 0; fi20 < 6; fi20++ {
                                var v21 uint16 = (^(p2))
                                v6 -= v6
                                switch (^(fi11)) {
                                case -107:
                                    v22 := []int{p0, (fi20 >> ((779) & 63)), p0, (fi11 << ((42) & 63))}
                                    _ = v22
                                }
                                _ = fi20
                                _ = v21
                            }
                            for fi23 := 0; fi23 < 4; fi23++ {
                                v19 = (-(v19))
                                _ = fi23
                            }
                        } else {
                            switch -573 {
                            case 32768:
                                if false {
                                    p0 -= fi11
                                    var v24 uint16 = 0
                                    for fi25 := 0; fi25 < 20; fi25++ {
                                        switch "foo bar" {
                                        case "hello":
                                            if (!false) {
                                                var v26 float64 = 0.1
                                                _ = v26
                                            } else {
                                                p0 = (-(p0))
                                                fmt.Println(v19, fi25, v24)
                                            }
                                        default:
                                            v24 = p2
                                        }
                                        _ = fi25
                                    }
                                    _ = v24
                                } else {
                                    v27 := GenS4{F0: (p1 / (((p1 >> ((65535) & 63))) | 1)), F1: ""}
                                    v28 := []int{(p0 * (-(p0))), fi11}
                                    _ = v27
                                    _ = v28
                                }
                            case 128:
                                for fi29 := 0; fi29 < 17; fi29++ {
                                    var v30 int64 = 1
                                    var v31 uint8 = 255
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            case -1:
                                p0 -= (p0 / (((^(p0))) | 1))
                            }
                        }
                        _ = v19
                    } else {
                        v32 := []int{p0}
                        _ = v32
                    }
                    _ = v12
                }
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -533
}
func genFn1(p0 uint, p1 int64) int16 {
    p0 += p0
    v33 := GenS3{F0: 255, F1: 1000, F2: (!false)}
    _ = p0
    _ = p1
    _ = v33
    return 32767
}

func main() {
    var v1 bool = (true || (0 < -2147483648))
    v2 := make(map[int]uint32)
    var v3 int16 = -468
    var v4 float32 = 1.5
    var v34 uint64 = 4294967295
    for fi35 := 0; fi35 < 10; fi35++ {
        v36 := GenS5{F0: 93, F1: fi35, F2: 0}
        v37 := GenS3{F0: 0, F1: 0, F2: v1}
        if (!false) {
            var v38 int32 = 1
            v39 := []int{(fi35 << ((18446744073709551615) & 63)), (fi35 * -9223372036854775808)}
            var v40 int64 = -858
            _ = v38
            _ = v39
            _ = v40
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    var v41 uint = 255
    if (!true) {
        v42 := GenS3{F0: 0, F1: (v3 / ((genFn0(4294967295, 501, 278)) | 1)), F2: (32767 > 0)}
        _ = v42
    }
    v43 := v2[966]
    v44 := []int{-733, 127, 32768}
    for fi45 := 0; fi45 < 18; fi45++ {
        var v46 uint16 = 709
        _ = fi45
        _ = v46
    }
    fmt.Println(v1, v41)
    fmt.Println(v1, v43, v1)
    for fi47 := 0; fi47 < 3; fi47++ {
        v48 := GenS3{F0: (v41 % ((65535) | 1)), F1: (v3 - 32767), F2: false}
        _ = fi47
        _ = v48
    }
    var v49 uint32 = (v43 ^ uint32(v34))
    v50 := GenS3{F0: 0, F1: 32767, F2: false}
    fmt.Println(v43, v4)
    fmt.Println(v34, v3, v1)
    v4 = 1e10
    switch "\n\t" {
    case "":
        v3 += v3
    case "hello":
        for fi51 := 0; fi51 < 4; fi51++ {
            v34 = (^(v34))
            for fi52 := 0; fi52 < 2; fi52++ {
                switch 128 {
                case -726:
                    fmt.Println(v34, v1)
                case 1:
                    if ((v4 + float32(float64((1.5 / v4)))) == ((0.1 * 3.14159) + v4)) {
                        delete(v2, (^(fi51)))
                        var v53 int = 377
                        v54 := GenS6{F0: (v3 ^ v3), F1: ("hello" + ("" + ("a" + "a"))), F2: v41}
                        _ = v53
                        _ = v54
                    }
                case -602:
                    v55 := GenS5{F0: 1, F1: -563, F2: 42}
                    _ = v55
                }
                _ = fi52
            }
            _ = fi51
        }
    case "a":
        switch -1 {
        case -9223372036854775808:
            var v56 int64 = 9223372036854775807
            _ = v56
        case 9223372036854775807:
            v57 := GenS6{F0: v3, F1: "hello", F2: 255}
            _ = v57
        }
    default:
        switch -1 {
        case 42:
            for fi58 := 0; fi58 < 1; fi58++ {
                if ((v34 * v34) > (^(v34))) {
                    fmt.Println(v49, v1)
                } else {
                    var v59 uint = v41
                    _ = v59
                }
                var v60 uint = v41
                if (!v1) {
                    for fi61 := 0; fi61 < 1; fi61++ {
                        fmt.Println(fi58, fi61)
                        v4 = (v4 - -2.25)
                        if (v1 && v1) {
                            var v62 uint8 = 2
                            var v63 uint = (v60 ^ 646)
                            var v64 string = "foo bar"
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                        _ = fi61
                    }
                    fmt.Println(v49, v1, v4)
                } else {
                    var v65 float64 = 0.0
                    v41 -= (^(v41))
                    _ = v65
                }
                _ = fi58
                _ = v60
            }
        case -9223372036854775808:
            var v66 string = "foo bar"
            _ = v66
        case 2147483647:
            fmt.Println(v1, v43)
        default:
            switch ("foo bar" + "hello") {
            case "a":
                v44 = append(v44, 32767)
            default:
                v67 := GenS7{F0: 42, F1: 32767, F2: 65535}
                _ = v67
            }
        }
    }
    if true {
        v68 := GenS8{F0: -9223372036854775808}
        v2[1] = 406
        _ = v68
    } else {
        var v69 uint = v41
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v43
    _ = v44
    _ = v49
    _ = v50
}
