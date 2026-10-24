package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 string; F2 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 int8; F2 uint64 }
type GenS10 struct { F0 int; F1 uint8; F2 uint }
type GenS11 struct { F0 int; F1 uint16; F2 string }
type GenS12 struct { F0 int32; F1 string; F2 uint8 }
type GenS13 struct { F0 int16; F1 uint64 }
type GenS14 struct { F0 uint64 }

func main() {
    var v1 bool = true
    var v2 int8 = -89
    v3 := make(map[int]int)
    var v4 int16 = 1
    v5 := []int{32767, 4294967295}
    if true {
        if (1 > 160) {
            v2 = (v2 * (^(v2)))
            var v6 int8 = -1
            fmt.Println(v2, v1)
            _ = v6
        } else {
            if true {
                switch "foo bar" {
                case "\n\t":
                    for fi7 := 0; fi7 < 14; fi7++ {
                        if v1 {
                            var v8 uint = 1
                            var v9 uint64 = 0
                            fmt.Println(v1, v8)
                            _ = v8
                            _ = v9
                        }
                        v2 += (-(v2))
                        _ = fi7
                    }
                default:
                    if v1 {
                        if (0 < 2) {
                            v10 := GenS1{F0: v4, F1: 18446744073709551615, F2: 865}
                            v11 := GenS2{F0: 36}
                            _ = v10
                            _ = v11
                        }
                    }
                }
                v2 = 1
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v2 = (v2 ^ (v2 << ((65535) & 63)))
                switch 128 {
                case -2147483648:
                    var v12 uint = 2
                    _ = v12
                default:
                    v2 -= v2
                }
            }
        }
        v13 := GenS3{F0: 42}
        delete(v3, 32768)
        _ = v13
    } else {
        switch -2147483648 {
        case 128:
            if (v1 || (42 != -2147483648)) {
                for fi14 := 0; fi14 < 11; fi14++ {
                    if v1 {
                        var v15 uint = 352
                        fmt.Println(v15, v4, v15)
                        _ = v15
                    }
                    _ = fi14
                }
                v2 = -128
            } else {
                v4 -= (v4 + v4)
                v16 := GenS1{F0: v4, F1: 932, F2: 2147483647}
                var v17 uint32 = 0
                _ = v16
                _ = v17
            }
        case 710:
            switch 750 {
            case -685:
                for fi18 := 0; fi18 < 19; fi18++ {
                    if (2147483647 != -321) {
                        _ = len(v3)
                        if v1 {
                            if (true || (!(!((-2.25 - 0.1) >= -2.25)))) {
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    if false {
                                        if true {
                                            if (((fi18 / (((fi18 - 128)) | 1)) < fi19) && v1) {
                                                var v20 uint16 = 0
                                                v2 -= (v2 / (((v2 << ((2) & 63))) | 1))
                                                v21 := GenS4{F0: 1}
                                                var v22 float32 = 0.0
                                                v4 = (v4 - v4)
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            } else {
                                                if false {
                                                    if v1 {
                                                        switch (^(fi18)) {
                                                        case -659:
                                                            if true {
                                                                fmt.Println(fi18, v1, v4)
                                                                if (false && true) {
                                                                    v23 := v3[fi19]
                                                                    v24 := GenS1{F0: v4, F1: 0, F2: 0}
                                                                    v23 -= v23
                                                                    _ = v23
                                                                    _ = v24
                                                                }
                                                            } else {
                                                                if v1 {
                                                                    fmt.Println(fi18, v1)
                                                                    if ("foo bar" != "a") {
                                                                        v5 = append(v5, fi18)
                                                                        fmt.Println(fi19, fi18, fi19)
                                                                    }
                                                                }
                                                                var v25 int8 = v2
                                                                _ = v25
                                                            }
                                                            fmt.Println(v4)
                                                        case 127:
                                                            switch 128 {
                                                            case 2147483647:
                                                                var v26 uint16 = 0
                                                                _ = v26
                                                            }
                                                        case 128:
                                                            fmt.Println(fi18, fi18)
                                                        }
                                                    } else {
                                                        v27 := v3[(fi18 << ((20) & 63))]
                                                        if v1 {
                                                            v28 := GenS3{F0: 42}
                                                            _ = v28
                                                        } else {
                                                            switch (^(fi19)) {
                                                            case 2147483647:
                                                                v29 := GenS4{F0: 0}
                                                                _ = v29
                                                            case 42:
                                                                fmt.Println(v1, fi19, v2)
                                                            case 9223372036854775807:
                                                                v27 -= (fi19 + 128)
                                                            }
                                                            if v1 {
                                                                v30 := GenS1{F0: v4, F1: 0, F2: v27}
                                                                _ = v30
                                                            } else {
                                                                v31 := GenS5{F0: 231, F1: "", F2: 127}
                                                                _ = v31
                                                            }
                                                            v3[32768] = 128
                                                        }
                                                        delete(v3, (^(fi19)))
                                                        _ = v27
                                                    }
                                                    if v1 {
                                                        switch fi18 {
                                                        case 4294967295:
                                                            if v1 {
                                                                v2 = v2
                                                            }
                                                        case -781:
                                                            var v32 int64 = 42
                                                            _ = v32
                                                        case -929:
                                                            v5 = append(v5, fi18)
                                                        default:
                                                            var v33 uint = 658
                                                            _ = v33
                                                        }
                                                        fmt.Println(v2, fi18)
                                                    } else {
                                                        v34 := GenS5{F0: 42, F1: ("\n\t" + "\n\t"), F2: 209}
                                                        _ = v34
                                                    }
                                                    v35 := GenS4{F0: 4294967295}
                                                    _ = v35
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v36 float64 = float64((((1.5 + 0.1) / -2.25) + (-2.25 + 3.14159)))
                                                        switch (fi19 / ((fi18) | 1)) {
                                                        case -1:
                                                            switch "a" {
                                                            case "foo bar":
                                                                v37 := GenS1{F0: 0, F1: 2, F2: 128}
                                                                _ = v37
                                                            default:
                                                                fmt.Println(v36)
                                                            }
                                                        case 2147483647:
                                                            v38 := GenS4{F0: 18446744073709551615}
                                                            _ = v38
                                                        }
                                                        _ = v36
                                                    case "foo bar":
                                                        var v39 uint8 = 2
                                                        _ = v39
                                                    case "a":
                                                        fmt.Println(fi19, fi18)
                                                    default:
                                                        v3[fi19] = fi18
                                                    }
                                                    fmt.Println(v1, v2, fi18)
                                                    v1 = v1
                                                }
                                                var v40 int = 15
                                                fmt.Println(v4, fi18)
                                                _ = v40
                                            }
                                            var v41 uint32 = 965
                                            var v42 uint = 65535
                                            _ = v41
                                            _ = v42
                                        }
                                        switch (fi18 + (fi19 ^ (fi19 - 4294967295))) {
                                        case 2147483647:
                                            if v1 {
                                                v4 -= (v4 & v4)
                                                var v43 uint16 = 69
                                                _ = v43
                                            }
                                        case -2147483648:
                                            switch ("" + ("" + "\n\t")) {
                                            case "a":
                                                v4 -= (v4 & 629)
                                            case "hello":
                                                v44 := GenS4{F0: 18446744073709551615}
                                                _ = v44
                                            }
                                        }
                                        v4 -= v4
                                    }
                                    v45 := GenS6{F0: 127}
                                    _ = len(v3)
                                    v46 := GenS7{F0: false}
                                    v47 := GenS6{F0: v2}
                                    _ = fi19
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                                v5 = append(v5, (fi18 ^ (^(fi18))))
                            }
                        } else {
                            var v48 uint = 1
                            _ = v48
                        }
                        var v49 int8 = (^(v2))
                        _ = v49
                    } else {
                        fmt.Println(v2)
                        v1 = v1
                    }
                    switch (fi18 << ((4294967295) & 63)) {
                    case -913:
                        v50 := GenS8{F0: 115}
                        _ = v50
                    case 698:
                        v51 := GenS5{F0: 18446744073709551615, F1: "hello", F2: 0}
                        _ = v51
                    }
                    v52 := GenS5{F0: 1, F1: ("a" + "a"), F2: 0}
                    _ = fi18
                    _ = v52
                }
            case 28:
                v2 -= 42
            default:
                var v53 int = -9223372036854775808
                _ = v53
            }
        }
        v1 = v1
    }
    switch 546 {
    case -2147483648:
        if (-9223372036854775808 != -1) {
            v54 := GenS9{F0: 9223372036854775807, F1: (-(v2)), F2: 0}
            v2 += (^(v2))
            _ = v54
        } else {
            for fi55 := 0; fi55 < 20; fi55++ {
                var v56 uint8 = 42
                v1 = true
                _ = fi55
                _ = v56
            }
        }
    }
    switch 9223372036854775807 {
    case 32768:
        switch -108 {
        case -9223372036854775808:
            switch "" {
            case "":
                v2 -= 1
            case "a":
                v57 := GenS1{F0: (v4 / ((v4) | 1)), F1: 495, F2: 32767}
                _ = v57
            case "\n\t":
                switch 757 {
                case 32768:
                    v1 = v1
                case 74:
                    v58 := GenS5{F0: 0, F1: "hello", F2: 255}
                    _ = v58
                }
            }
        }
    case 0:
        for fi59 := 0; fi59 < 5; fi59++ {
            var v60 float32 = -2.25
            var v61 string = "hello"
            _ = fi59
            _ = v60
            _ = v61
        }
    }
    if v1 {
        var v62 int = 2147483647
        _ = v62
    } else {
        v4 = -32768
    }
    var v63 string = ""
    v2 = v2
    var v64 int = 42
    v2 = (v2 & 1)
    switch v64 {
    case 450:
        v3[0] = v64
    case 4294967295:
        v1 = ((0 == 0) || (!true))
    case 9223372036854775807:
        for fi65 := 0; fi65 < 6; fi65++ {
            v2 = v2
            v66 := GenS10{F0: (fi65 ^ v64), F1: 1, F2: 513}
            _ = fi65
            _ = v66
        }
    default:
        if (v64 > (v64 | v64)) {
            for fi67 := 0; fi67 < 10; fi67++ {
                fmt.Println(v4, v64, v64)
                for fi68 := 0; fi68 < 3; fi68++ {
                    v69 := GenS8{F0: 195}
                    if len(v5) > 0 {
                        v5[0] = (-(fi68))
                    }
                    var v70 uint16 = 65535
                    if (0 < 551) {
                        v71 := GenS1{F0: 1000, F1: 2, F2: 456}
                        v72 := GenS11{F0: fi67, F1: (v70 % ((379) | 1)), F2: ("" + ("" + v63))}
                        v63 = "\n\t"
                        _ = v71
                        _ = v72
                    }
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                _ = fi67
            }
            v63 = ""
            for fi73 := 0; fi73 < 2; fi73++ {
                for fi74 := 0; fi74 < 12; fi74++ {
                    v1 = true
                    switch 1 {
                    case 32768:
                        v63 = (("foo bar" + v63) + "\n\t")
                    case 173:
                        fmt.Println(fi73, v1, v63)
                    }
                    v75 := GenS3{F0: 42}
                    var v76 uint64 = 4294967295
                    var v77 uint16 = 255
                    _ = fi74
                    _ = v75
                    _ = v76
                    _ = v77
                }
                v5 = append(v5, -284)
                _ = fi73
            }
        } else {
            switch v63 {
            case "":
                v78 := v3[int(v4)]
                _ = v78
            case "\n\t":
                v5 = append(v5, 4294967295)
            default:
                v2 = 100
            }
        }
    }
    for fi79 := 0; fi79 < 3; fi79++ {
        if v1 {
            v63 = v63
            v3[(fi79 - (fi79 | fi79))] = fi79
            for fi80 := 0; fi80 < 19; fi80++ {
                if v1 {
                    v81 := GenS4{F0: 4294967295}
                    _ = len(v3)
                    _ = v81
                } else {
                    fmt.Println(v1)
                    delete(v3, 2147483647)
                }
                v2 -= v2
                _ = fi80
            }
        }
        _ = fi79
    }
    if (v1 || false) {
        v3[-906] = (v64 | v64)
    } else {
        var v82 int32 = 1
        if v1 {
            v83 := GenS12{F0: (v82 >> ((255) & 63)), F1: v63, F2: 127}
            var v84 bool = (1 <= 44)
            for fi85 := 0; fi85 < 1; fi85++ {
                var v86 uint32 = 42
                v87 := GenS13{F0: 208, F1: 4294967295}
                var v88 int32 = v82
                _ = fi85
                _ = v86
                _ = v87
                _ = v88
            }
            _ = v83
            _ = v84
        } else {
            if v1 {
                var v89 uint8 = 1
                var v90 uint = 0
                fmt.Println(v82)
                _ = v89
                _ = v90
            } else {
                for fi91 := 0; fi91 < 4; fi91++ {
                    if v1 {
                        for fi92 := 0; fi92 < 2; fi92++ {
                            switch fi92 {
                            case 32767:
                                var v93 int = fi92
                                v94 := GenS14{F0: 4294967295}
                                _ = v93
                                _ = v94
                            case 346:
                                var v95 uint16 = 0
                                _ = v95
                            case 4294967295:
                                switch 4294967295 {
                                case 289:
                                    var v96 uint = 255
                                    _ = v96
                                case 548:
                                    v2 -= v2
                                    var v97 uint64 = 1
                                    _ = v97
                                default:
                                    v98 := v3[v64]
                                    _ = v98
                                }
                            default:
                                v63 = "foo bar"
                            }
                            _ = fi92
                        }
                        _ = len(v3)
                        _ = len(v3)
                    }
                    _ = fi91
                }
            }
            fmt.Println(v63, v1)
            v1 = (!(1.5 != (0.0 + 0.0)))
        }
        _ = v82
    }
    fmt.Println(v2, v64, v4)
    fmt.Println(v2, v64, v64)
    var v99 int32 = -1
    v2 = v2
    var v100 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v99
}
