package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 int16 }
type GenS5 struct { F0 int32; F1 int64; F2 string }
type GenS6 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 343
    var v3 int16 = 1
    v4 := make(map[string]int)
    var v5 float32 = (1e10 / ((0.0 * -2.25) - float32(3.14159)))
    var v6 uint8 = 255
    fmt.Println(v3)
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(fi7, v2, v2)
            switch (fi7 % (((fi8 & 2147483647)) | 1)) {
            case 32768:
                switch (v2 + fi7) {
                case 970:
                    var v9 uint64 = 0
                    _ = v9
                case -194:
                    v3 -= v3
                case 0:
                    var v10 uint16 = 65535
                    v11 := []int{(fi7 - (fi8 + fi7))}
                    _ = v10
                    _ = v11
                default:
                    fmt.Println(v6, v3)
                }
            case 1:
                if true {
                    v12 := GenS1{F0: -2147483648, F1: 359, F2: 42}
                    switch 4294967295 {
                    case -1:
                        var v13 float32 = v5
                        switch "hello" {
                        case "hello":
                            v5 = v5
                        default:
                            var v14 bool = true
                            _ = v14
                        }
                        _ = v13
                    case -394:
                        fmt.Println(v3)
                    }
                    var v15 uint16 = 32767
                    _ = v12
                    _ = v15
                } else {
                    v16 := GenS2{F0: (v2 & (fi8 & (fi7 & -2147483648))), F1: 1, F2: 42}
                    var v17 int = (^(v2))
                    if (703 > 4294967295) {
                        fmt.Println(v3)
                        v18 := []int{v2}
                        v2 = v17
                        if true {
                            switch 677 {
                            case 594:
                                _ = len(v4)
                            default:
                                v5 += v5
                            }
                            v19 := GenS2{F0: (v2 - v2), F1: 42, F2: 0}
                            _ = len(v18)
                            _ = v19
                        }
                        _ = v18
                    }
                    _ = v16
                    _ = v17
                }
            default:
                var v20 int64 = 0
                _ = v20
            }
            _ = fi8
        }
        v5 = -2.25
        for fi21 := 0; fi21 < 13; fi21++ {
            fmt.Println(v5, v3, v6)
            _ = fi21
        }
        _ = fi7
    }
    if true {
        var v22 int8 = 127
        for fi23 := 0; fi23 < 5; fi23++ {
            v6 = v6
            v22 += (v22 / ((v22) | 1))
            v5 -= 0.0
            _ = fi23
        }
        _ = v22
    } else {
        v6 = 2
        v24 := GenS2{F0: v2, F1: -40, F2: 18446744073709551615}
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 2; fi26++ {
            v3 -= v3
            var v27 uint32 = 65535
            _ = fi26
            _ = v27
        }
        fmt.Println(v5)
        v28 := GenS3{F0: (v3 & 32767)}
        _ = fi25
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 20; fi30++ {
            v31 := GenS4{F0: 652, F1: (v3 & (v3 & v3)), F2: -32768}
            v3 = 1000
            switch -277 {
            case 127:
                var v32 uint32 = 42
                _ = v32
            case 128:
                var v33 int32 = 1
                _ = v33
            case -1:
                v34 := GenS5{F0: int32(v2), F1: 0, F2: "\n\t"}
                _ = v34
            }
            switch fi30 {
            case 1:
                var v35 int64 = -9223372036854775808
                _ = v35
            }
            v36 := GenS6{F0: (255 == 978), F1: 391, F2: 42}
            _ = fi30
            _ = v31
            _ = v36
        }
        for fi37 := 0; fi37 < 19; fi37++ {
            var v38 int32 = -91
            switch ("foo bar" + "hello") {
            case "hello":
                var v39 uint = 492
                _ = v39
            case "a":
                switch (v2 * fi37) {
                case 117:
                    var v40 int = 32768
                    _ = v40
                case 732:
                    var v41 int64 = 519
                    _ = v41
                case -928:
                    v38 += (v38 * v38)
                default:
                    v38 = v38
                    if (((v2 % ((32768) | 1)) < 32767) && true) {
                        v42 := v4[""]
                        var v43 uint16 = 255
                        if (1 != 499) {
                            var v44 uint64 = 4294967295
                            _ = v44
                        } else {
                            fmt.Println(v5, v3, v42)
                            var v45 int16 = v3
                            switch (v42 / (((fi29 % (((^(fi29))) | 1))) | 1)) {
                            case 546:
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v2, v43)
                                case "hello":
                                    v38 = (-(v38))
                                case "":
                                    v46 := GenS7{F0: ""}
                                    _ = v46
                                }
                            case 342:
                                v3 = -1
                            }
                            _ = v45
                        }
                        _ = v42
                        _ = v43
                    } else {
                        v47 := GenS8{F0: 32767, F1: 322}
                        v6 += (v6 | (^(v6)))
                        _ = v47
                    }
                    fmt.Println(v38, v6, fi29)
                }
            }
            if false {
                v38 += (v38 + v38)
                v5 += ((v5 - 1.5) * v5)
            }
            var v48 uint16 = 255
            var v49 uint16 = v48
            switch fi29 {
            case -749:
                v48 = 100
            case 9223372036854775807:
                var v50 uint8 = (v6 | 42)
                _ = v50
            default:
                switch fi29 {
                case 612:
                    v6 = 2
                }
            }
            _ = fi37
            _ = v38
            _ = v48
            _ = v49
        }
        _ = fi29
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        v52 := GenS6{F0: (v3 < (v3 - v3)), F1: 951, F2: 828}
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 20; fi53++ {
        switch "\n\t" {
        case "a":
            v5 -= v5
        }
        _ = fi53
    }
    fmt.Println(v6, v3, v5)
    var v54 int32 = 183
    for fi55 := 0; fi55 < 15; fi55++ {
        if (!(true && false)) {
            v56 := GenS9{F0: 9223372036854775807, F1: v6}
            if false {
                v54 -= v54
            } else {
                switch 685 {
                case 425:
                    for fi57 := 0; fi57 < 6; fi57++ {
                        v2 = fi57
                        var v58 float32 = 1.5
                        _ = fi57
                        _ = v58
                    }
                default:
                    v3 = -285
                }
                fmt.Println(v6, v5)
            }
            _ = v56
        } else {
            var v59 float64 = ((float64(v5) - 0.1) + (0.1 - 1e10))
            v59 += float64(v5)
            v2 -= v2
            _ = v59
        }
        v2 = v2
        _ = fi55
    }
    var v60 bool = (128 > (v6 % (((v6 ^ (v6 >> ((722) & 63)))) | 1)))
    if false {
        var v61 float32 = v5
        switch "hello" {
        case "":
            fmt.Println(v60, v2)
        default:
            v54 -= v54
        }
        var v62 int64 = 0
        _ = v61
        _ = v62
    }
    var v63 uint16 = 392
    for fi64 := 0; fi64 < 16; fi64++ {
        var v65 bool = false
        for fi66 := 0; fi66 < 12; fi66++ {
            switch "foo bar" {
            case "":
                v67 := []int{(v2 * fi64), fi64, fi64, v2}
                _ = v67
            case "foo bar":
                var v68 float32 = float32(float64(v5))
                _ = v68
            }
            _ = fi66
        }
        _ = fi64
        _ = v65
    }
    if (v60 && (v60 && (901 == 719))) {
        v60 = (!v60)
        fmt.Println(v3, v60, v54)
    }
    var v69 int8 = -128
    fmt.Println(v5, v2, v69)
    for fi70 := 0; fi70 < 3; fi70++ {
        v69 -= v69
        for fi71 := 0; fi71 < 14; fi71++ {
            v4["\n\t"] = fi71
            var v72 string = ""
            if (v60 || v60) {
                var v73 int64 = -462
                _ = v73
            } else {
                var v74 float32 = float32(0.0)
                _ = v74
            }
            _ = fi71
            _ = v72
        }
        _ = fi70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v60
    _ = v63
    _ = v69
}
