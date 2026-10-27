package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int64; F1 uint64; F2 uint }
type GenS7 struct { F0 int16; F1 int; F2 int64 }
type GenS8 struct { F0 uint8; F1 int16; F2 int8 }

func main() {
    var v1 uint = 859
    var v2 uint = v1
    v3 := make(map[string]bool)
    var v4 int64 = 415
    v5 := []uint32{445, uint32(v2), 1}
    if false {
        v3[("a" + ((((("foo bar" + "a") + "") + "") + ("\n\t" + "\n\t")) + (("a" + ("foo bar" + "a")) + ("\n\t" + "foo bar"))))] = false
        switch "foo bar" {
        case "\n\t":
            var v6 uint32 = 0
            _ = v6
        case "":
            if true {
                v1 += (v2 << ((v1) & 63))
                var v7 float64 = 1.5
                _ = len(v3)
                _ = v7
            }
        default:
            v4 += (^(v4))
        }
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            if true {
                if true {
                    if (-508 != -1) {
                        fmt.Println(fi9)
                        v3["\n\t"] = (32767 > 0)
                        v10 := GenS1{F0: 1, F1: true, F2: 128}
                        _ = v10
                    } else {
                        var v11 uint32 = 0
                        fmt.Println(v11, v2, v1)
                        switch ("a" + (("hello" + ("\n\t" + "foo bar")) + "foo bar")) {
                        case "foo bar":
                            v12 := GenS2{F0: 1000, F1: 42}
                            _ = v12
                        case "a":
                            if true {
                                var v13 int32 = -333
                                _ = v13
                            } else {
                                if (0 <= (v4 >> ((v1) & 63))) {
                                    v14 := GenS2{F0: 42, F1: 255}
                                    _ = v14
                                } else {
                                    v11 = v11
                                    var v15 uint16 = 943
                                    var v16 uint64 = 25
                                    _ = v15
                                    _ = v16
                                }
                                var v17 float64 = (0.1 + 1.5)
                                var v18 int32 = 397
                                _ = v17
                                _ = v18
                            }
                            v4 += (v4 + (v4 >> ((4294967295) & 63)))
                            v4 += v4
                        case "\n\t":
                            switch "hello" {
                            case "foo bar":
                                var v19 float64 = 0.0
                                _ = v19
                            case "a":
                                delete(v3, "foo bar")
                            case "hello":
                                v2 += 0
                            default:
                                if true {
                                    fmt.Println(v4, v1, fi8)
                                    var v20 uint32 = (v11 % ((v11) | 1))
                                    _ = v20
                                }
                            }
                        }
                        _ = v11
                    }
                    v21 := GenS1{F0: 231, F1: (!false), F2: 0}
                    v1 = (v1 % (((v2 * v2)) | 1))
                    _ = v21
                }
                switch (-(fi9)) {
                case 127:
                    v2 -= v2
                case 128:
                    v1 -= (v2 / ((v1) | 1))
                default:
                    v2 -= (v2 - 2)
                }
                var v22 float64 = (((1.5 / float64(1e10)) + 0.1) + (float64(float32(0.0)) / float64(3.14159)))
                _ = v22
            } else {
                var v23 int16 = 1
                var v24 int32 = -1
                v23 -= (v23 + v23)
                _ = v23
                _ = v24
            }
            v4 += (v4 / ((v4) | 1))
            v1 += (^(v2))
            v25 := GenS3{F0: 0}
            v26 := GenS2{F0: 1000, F1: 65535}
            var v27 uint64 = 843
            _ = fi9
            _ = v25
            _ = v26
            _ = v27
        }
        var v28 uint32 = 950
        _ = fi8
        _ = v28
    }
    v29 := GenS2{F0: 1, F1: 2}
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            _ = len(v5)
            v32 := v3["hello"]
            v33 := v3[("hello" + "foo bar")]
            _ = fi31
            _ = v32
            _ = v33
        }
        v1 += (v1 + v2)
        v2 += (v1 ^ 698)
        _ = fi30
    }
    var v34 uint8 = 2
    var v35 int64 = int64(v1)
    switch "hello" {
    case "a":
        fmt.Println(v34, v1)
    case "hello":
        var v36 uint8 = v34
        _ = v36
    case "foo bar":
        for fi37 := 0; fi37 < 2; fi37++ {
            fmt.Println(v1, v4, fi37)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if ((-128 < 120) && (360 < 42)) {
                for fi38 := 0; fi38 < 6; fi38++ {
                    var v39 int16 = -1
                    if true {
                        fmt.Println(v4, v4, fi37)
                        v40 := GenS2{F0: (v39 * v39), F1: 42}
                        v41 := GenS4{F0: (v34 % ((93) | 1)), F1: v2, F2: v34}
                        _ = len(v3)
                        _ = v40
                        _ = v41
                    } else {
                        var v42 int64 = (v35 - v4)
                        switch "foo bar" {
                        case "foo bar":
                            v3[("\n\t" + "\n\t")] = false
                        default:
                            v2 = v1
                        }
                        if true {
                            if len(v5) > 0 {
                                v5[0] = 407
                            }
                        }
                        var v43 uint64 = 75
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        switch "a" {
                        case "":
                            switch (fi38 | fi38) {
                            case -234:
                                if (float64(0.1) == (0.0 + 3.14159)) {
                                    v44 := GenS3{F0: 42}
                                    if true {
                                        v1 = (v1 % (((v2 + (^(v2)))) | 1))
                                        switch (fi38 << ((v2) & 63)) {
                                        case 32767:
                                            v5 = append(v5, 1)
                                        case 4294967295:
                                            v34 = v34
                                        default:
                                            fmt.Println(v2)
                                        }
                                        if (true && (42 <= 35)) {
                                            if (false && true) {
                                                var v45 uint = 2
                                                v42 -= v4
                                                v35 = int64(v39)
                                                _ = v45
                                            }
                                            fmt.Println(v4)
                                        }
                                    } else {
                                        var v46 uint64 = (v43 & (v43 / (((v43 + 470)) | 1)))
                                        var v47 int = fi37
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = v44
                                } else {
                                    var v48 int32 = -2147483648
                                    delete(v3, "")
                                    _ = v48
                                }
                            default:
                                fmt.Println(v42)
                            }
                        }
                        v35 = v42
                        var v49 float64 = float64((float32(3.14159) * 0.1))
                        _ = v42
                        _ = v43
                        _ = v49
                    }
                    _ = fi38
                    _ = v39
                }
                v50 := GenS5{F0: (fi37 << ((v2) & 63)), F1: v2}
                _ = v50
            }
            _ = fi37
        }
    default:
        if (!(float32(0.0) > float32((float64(0.1) + -2.25)))) {
            if true {
                var v51 float32 = (0.1 / 1.0)
                if false {
                    for fi52 := 0; fi52 < 6; fi52++ {
                        v53 := GenS6{F0: 42, F1: 0, F2: 85}
                        if (2147483648 == 2147483647) {
                            _ = len(v5)
                            var v54 uint16 = 689
                            fmt.Println(v34)
                            _ = v54
                        } else {
                            var v55 float32 = (v51 - (v51 + v51))
                            _ = len(v5)
                            v56 := GenS1{F0: 2, F1: false, F2: v34}
                            _ = v55
                            _ = v56
                        }
                        _ = fi52
                        _ = v53
                    }
                    switch 128 {
                    case -354:
                        v57 := GenS6{F0: v4, F1: 1, F2: (v2 + (v2 * (v2 - 956)))}
                        _ = v57
                    }
                    v58 := GenS7{F0: 1, F1: 1, F2: int64(v2)}
                    _ = v58
                }
                var v59 int16 = -309
                _ = v51
                _ = v59
            }
            var v60 int8 = 0
            _ = v60
        } else {
            for fi61 := 0; fi61 < 2; fi61++ {
                var v62 uint32 = 314
                var v63 float64 = float64((1e10 / 1.5))
                v64 := v3[(("foo bar" + ("hello" + "a")) + "foo bar")]
                _ = fi61
                _ = v62
                _ = v63
                _ = v64
            }
        }
    }
    var v65 int32 = -1
    v65 = (v65 / ((v65) | 1))
    for fi66 := 0; fi66 < 13; fi66++ {
        if len(v5) > 0 {
            v5[0] = 1
        }
        var v67 string = "\n\t"
        switch fi66 {
        case -968:
            for fi68 := 0; fi68 < 6; fi68++ {
                _ = len(v3)
                v69 := GenS5{F0: (fi66 | fi66), F1: v2}
                _ = fi68
                _ = v69
            }
        default:
            for fi70 := 0; fi70 < 1; fi70++ {
                if false {
                    if true {
                        v67 = v67
                        v2 -= 18446744073709551615
                    } else {
                        var v71 uint8 = (^(v34))
                        v35 += (v35 * v35)
                        v72 := GenS1{F0: 2, F1: (v34 == 239), F2: (v71 * 127)}
                        _ = v71
                        _ = v72
                    }
                    v67 = v67
                }
                var v73 int64 = 9223372036854775807
                v74 := GenS8{F0: 18, F1: -156, F2: -128}
                _ = fi70
                _ = v73
                _ = v74
            }
        }
        _ = fi66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v65
}
