package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
type GenS4 struct { F0 int8; F1 bool; F2 bool }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint32; F1 string; F2 int32 }

func main() {
    var v1 uint64 = 2
    var v2 int32 = 578
    var v3 float64 = (0.1 - 0.1)
    var v4 int8 = -24
    v1 = v1
    v5 := []int{-9223372036854775808, 9223372036854775807, 4294967295, -625}
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float32 = (float32(v3) - (3.14159 - 1.5))
        var v8 int16 = 1
        v9 := GenS1{F0: 256}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    if (!(!(2147483647 >= 2147483647))) {
        v10 := GenS1{F0: 619}
        v1 = v1
        if true {
            switch 1 {
            case 42:
                v11 := GenS1{F0: uint32(v1)}
                _ = v11
            case -9223372036854775808:
                var v12 bool = true
                _ = v12
            case 1:
                for fi13 := 0; fi13 < 1; fi13++ {
                    if false {
                        if (false || false) {
                            fmt.Println(v1)
                            v14 := GenS1{F0: 2}
                            _ = v14
                        }
                    } else {
                        if true {
                            fmt.Println(v1, v2, fi13)
                        }
                        if (((504 < 255) || (!(v4 > (v4 | v4)))) && (858 >= 42)) {
                            switch (-(fi13)) {
                            case -342:
                                v15 := GenS1{F0: 713}
                                _ = v15
                            case 42:
                                if (((v3 / float64(3.14159)) + v3) != v3) {
                                    fmt.Println(v1)
                                } else {
                                    switch fi13 {
                                    case 128:
                                        for fi16 := 0; fi16 < 5; fi16++ {
                                            if (!true) {
                                                v17 := GenS1{F0: 0}
                                                v18 := GenS2{F0: (((2 >= 65535) || (!(true && (-32768 == -1)))) && ((1.5 >= (0.0 * (3.14159 + -2.25))) || false))}
                                                if false {
                                                    if true {
                                                        switch fi13 {
                                                        case 842:
                                                            if false {
                                                                switch "a" {
                                                                case "hello":
                                                                    var v19 uint32 = 4294967295
                                                                    _ = v19
                                                                case "a":
                                                                    var v20 uint8 = 180
                                                                    _ = v20
                                                                default:
                                                                    var v21 bool = (fi16 == 32768)
                                                                    _ = v21
                                                                }
                                                                v2 -= (v2 * v2)
                                                                switch "a" {
                                                                case "hello":
                                                                    v4 = (v4 + v4)
                                                                case "foo bar":
                                                                    switch fi16 {
                                                                    case -598:
                                                                        var v22 uint16 = 2
                                                                        _ = v22
                                                                    case -9223372036854775808:
                                                                        _ = len(v5)
                                                                    case 32767:
                                                                        switch (-(fi16)) {
                                                                        case 32767:
                                                                            v3 += v3
                                                                            switch "" {
                                                                            case "a":
                                                                                if true {
                                                                                    if (((fi13 < (fi13 >> ((0) & 63))) && false) && ((((v2 * (v2 ^ (v2 % (((v2 >> ((2) & 63))) | 1)))) != v2) || true) && true)) {
                                                                                        fmt.Println(fi16, v3, fi16)
                                                                                        fmt.Println(v3, fi16)
                                                                                    } else {
                                                                                        switch fi13 {
                                                                                        case 9223372036854775807:
                                                                                            if false {
                                                                                                if ((!true) || (false || false)) {
                                                                                                    if (false || true) {
                                                                                                        v2 = (v2 - 2147483647)
                                                                                                        var v23 uint8 = 2
                                                                                                        v24 := GenS3{F0: -643, F1: v1, F2: 0}
                                                                                                        _ = v23
                                                                                                        _ = v24
                                                                                                    }
                                                                                                    if (4294967295 < 18446744073709551615) {
                                                                                                        v4 += v4
                                                                                                        var v25 int = 0
                                                                                                        v26 := GenS2{F0: (128 <= 10)}
                                                                                                        _ = v25
                                                                                                        _ = v26
                                                                                                    }
                                                                                                    var v27 float64 = (v3 + float64(float32(1e10)))
                                                                                                    _ = v27
                                                                                                } else {
                                                                                                    v5 = append(v5, fi16)
                                                                                                }
                                                                                                var v28 int32 = v2
                                                                                                v1 += (v1 & (v1 ^ (^(v1))))
                                                                                                _ = v28
                                                                                            }
                                                                                        case -693:
                                                                                            if (true || true) {
                                                                                                if len(v5) >= 2 {
                                                                                                    v5 = v5[1:len(v5)]
                                                                                                }
                                                                                                var v29 string = ("a" + (((("a" + "hello") + "a") + "\n\t") + ""))
                                                                                                v30 := GenS2{F0: true}
                                                                                                _ = v29
                                                                                                _ = v30
                                                                                            } else {
                                                                                                _ = len(v5)
                                                                                                v31 := GenS4{F0: 127, F1: true, F2: (65535 == 2)}
                                                                                                v3 -= v3
                                                                                                _ = v31
                                                                                            }
                                                                                        case 479:
                                                                                            v3 -= 1e10
                                                                                        default:
                                                                                            if (0 >= -573) {
                                                                                                var v32 uint = 950
                                                                                                switch (fi16 + fi16) {
                                                                                                case -800:
                                                                                                    if (true && false) {
                                                                                                        switch ("a" + "foo bar") {
                                                                                                        case "\n\t":
                                                                                                            v4 += (v4 & 100)
                                                                                                        }
                                                                                                        if (-180 == 0) {
                                                                                                            var v33 uint = (v32 / (((v32 | 255)) | 1))
                                                                                                            v34 := GenS2{F0: false}
                                                                                                            _ = v33
                                                                                                            _ = v34
                                                                                                        }
                                                                                                        if len(v5) >= 2 {
                                                                                                            v5 = v5[1:len(v5)]
                                                                                                        }
                                                                                                        v2 -= -288
                                                                                                        v4 = 79
                                                                                                    }
                                                                                                default:
                                                                                                    var v35 uint16 = uint16(v32)
                                                                                                    _ = v35
                                                                                                }
                                                                                                if (!(279 >= -32768)) {
                                                                                                    v36 := GenS2{F0: true}
                                                                                                    fmt.Println(fi16)
                                                                                                    _ = v36
                                                                                                }
                                                                                                _ = v32
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v2, v2)
                                                                                    }
                                                                                    var v37 int16 = 32767
                                                                                    _ = v37
                                                                                }
                                                                            default:
                                                                                switch fi16 {
                                                                                case 128:
                                                                                    var v38 uint16 = 40
                                                                                    _ = v38
                                                                                case 32767:
                                                                                    fmt.Println(fi13, v2)
                                                                                default:
                                                                                    v3 = v3
                                                                                    switch ("hello" + "\n\t") {
                                                                                    case "":
                                                                                        if (-193 < -1) {
                                                                                            if true {
                                                                                                v2 += v2
                                                                                                v1 -= (v1 % ((v1) | 1))
                                                                                                v4 += v4
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    v1 -= 18446744073709551615
                                                                                                case "a":
                                                                                                    var v39 int8 = v4
                                                                                                    _ = v39
                                                                                                default:
                                                                                                    var v40 int8 = (v4 + v4)
                                                                                                    _ = v40
                                                                                                }
                                                                                            }
                                                                                            var v41 int16 = -1
                                                                                            var v42 int64 = 601
                                                                                            _ = v41
                                                                                            _ = v42
                                                                                        } else {
                                                                                            v1 += (v1 & v1)
                                                                                            fmt.Println(v3, fi13, v2)
                                                                                        }
                                                                                    default:
                                                                                        _ = len(v5)
                                                                                    }
                                                                                }
                                                                            }
                                                                            v43 := GenS4{F0: v4, F1: (433 > 1), F2: false}
                                                                            _ = v43
                                                                        default:
                                                                            v44 := GenS3{F0: 400, F1: v1, F2: 0}
                                                                            _ = v44
                                                                        }
                                                                    default:
                                                                        if len(v5) > 0 {
                                                                            v5[0] = fi16
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            var v45 int32 = v2
                                                            _ = v45
                                                        }
                                                        var v46 uint16 = 1
                                                        v4 -= v4
                                                        _ = v46
                                                    }
                                                    var v47 uint16 = 255
                                                    v3 -= float64(0.1)
                                                    _ = v47
                                                }
                                                _ = v17
                                                _ = v18
                                            } else {
                                                var v48 float32 = 0.0
                                                fmt.Println(v2, fi13, v48)
                                                _ = v48
                                            }
                                            _ = fi16
                                        }
                                    case -1:
                                        v1 -= (v1 / (((v1 >> ((814) & 63))) | 1))
                                    case 32767:
                                        v5 = append(v5, fi13)
                                    }
                                    fmt.Println(v4)
                                    v3 = (((1.5 + ((3.14159 / v3) - float64(-2.25))) - (float64(0.0) * v3)) + float64(0.0))
                                }
                            case 4294967295:
                                var v49 string = ("\n\t" + "a")
                                _ = v49
                            }
                        }
                        fmt.Println(fi13, v2)
                    }
                    if false {
                        fmt.Println(v1, v4, v2)
                        var v50 int32 = v2
                        _ = v50
                    } else {
                        var v51 bool = (-9223372036854775808 != -9223372036854775808)
                        _ = v51
                    }
                    _ = fi13
                }
            }
            for fi52 := 0; fi52 < 13; fi52++ {
                fmt.Println(fi52, v4, fi52)
                _ = fi52
            }
            var v53 uint64 = (v1 / (((v1 % (((v1 >> ((970) & 63))) | 1))) | 1))
            _ = v53
        } else {
            fmt.Println(v2, v1)
            v54 := GenS5{F0: v2, F1: 127, F2: 65535}
            var v55 bool = false
            _ = v54
            _ = v55
        }
        _ = v10
    } else {
        var v56 int16 = -32768
        var v57 bool = false
        _ = v56
        _ = v57
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    if false {
        for fi58 := 0; fi58 < 6; fi58++ {
            var v59 uint32 = 568
            var v60 uint16 = 42
            _ = fi58
            _ = v59
            _ = v60
        }
        for fi61 := 0; fi61 < 20; fi61++ {
            fmt.Println(v1)
            _ = fi61
        }
    } else {
        var v62 int8 = (v4 % (((v4 & (v4 & (^(v4))))) | 1))
        v63 := GenS6{F0: 2147483648, F1: "a", F2: (v2 / ((42) | 1))}
        _ = v62
        _ = v63
    }
    v64 := GenS1{F0: 2147483647}
    for fi65 := 0; fi65 < 2; fi65++ {
        var v66 bool = (1 >= 9223372036854775807)
        _ = fi65
        _ = v66
    }
    v3 += 3.14159
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi67 := 0; fi67 < 16; fi67++ {
        fmt.Println(v4, v2, v3)
        var v68 float64 = v3
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi67
        _ = v68
    }
    for fi69 := 0; fi69 < 17; fi69++ {
        v5 = append(v5, fi69)
        if len(v5) > 0 {
            v5[0] = 0
        }
        v2 = v2
        _ = fi69
    }
    if true {
        fmt.Println(v3, v1)
    } else {
        fmt.Println(v4)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 235 {
    case 642:
        v4 = (v4 + -128)
    case 593:
        var v70 int8 = v4
        _ = v70
    }
    switch 127 {
    case -568:
        for fi71 := 0; fi71 < 15; fi71++ {
            var v72 uint16 = 972
            switch fi71 {
            case 109:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 9223372036854775807:
                var v73 uint32 = 634
                _ = v73
            }
            v72 -= 65535
            _ = fi71
            _ = v72
        }
    case 1:
        fmt.Println(v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
}
