package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 uint64 }
type GenS3 struct { F0 int32; F1 int32; F2 uint }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int16 }

func main() {
    v1 := []int{}
    v2 := make(map[int]int)
    var v3 float32 = (3.14159 - ((3.14159 + -2.25) - float32(0.1)))
    var v4 string = ""
    var v5 float64 = -2.25
    v6 := GenS1{F0: 42, F1: 41, F2: 62}
    fmt.Println(v5, v4)
    switch "" {
    case "\n\t":
        if false {
            fmt.Println(v4, v3)
            fmt.Println(v3)
        } else {
            v7 := GenS1{F0: -2147483648, F1: 100, F2: 255}
            v2[727] = -1
            _ = v7
        }
    default:
        v3 += 1e10
    }
    v8 := GenS1{F0: 32767, F1: 100, F2: 11}
    if true {
        v3 = float32(1.5)
    } else {
        v9 := GenS1{F0: 1, F1: 1, F2: 2}
        if (false && (!true)) {
            var v10 int = 32767
            _ = v10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            v1 = append(v1, (fi11 ^ (fi11 * -879)))
            v12 := GenS1{F0: 32767, F1: 122, F2: 136}
            _ = fi11
            _ = v12
        }
        _ = v9
    }
    var v13 int16 = 1000
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: v13, F1: (v3 > (1.5 / -2.25)), F2: 0}
            v4 = (v4 + "")
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: v13, F1: true, F2: 18446744073709551615}
        v3 += v3
        _ = fi14
        _ = v17
    }
    var v18 uint16 = 657
    switch v4 {
    case "":
        switch 1 {
        case -142:
            if (false && false) {
                fmt.Println(v5)
            }
        case 32768:
            if false {
                _ = len(v1)
                var v19 int16 = (v13 << ((65535) & 63))
                _ = v19
            }
        case -621:
            switch -2147483648 {
            case 127:
                v20 := GenS2{F0: v13, F1: false, F2: 4294967295}
                _ = v20
            case -1:
                fmt.Println(v18)
            }
        default:
            if true {
                fmt.Println(v4)
                var v21 uint8 = 1
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 float32 = v3
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := GenS3{F0: 42, F1: -895, F2: 0}
                    _ = fi24
                    _ = v25
                }
                if (128 < 0) {
                    var v26 int64 = -9223372036854775808
                    for fi27 := 0; fi27 < 16; fi27++ {
                        fmt.Println(v18, v13)
                        if true {
                            v28 := v2[(fi27 + (fi27 - fi27))]
                            switch "a" {
                            case "hello":
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    var v30 uint = 42
                                    _ = fi29
                                    _ = v30
                                }
                            case "foo bar":
                                if (!(float64(-2.25) <= ((v5 - 0.0) + v5))) {
                                    var v31 string = "foo bar"
                                    v31 = "foo bar"
                                    _ = v31
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (fi27 >> ((255) & 63))
                                    }
                                }
                            }
                            var v32 int32 = -42
                            _ = v28
                            _ = v32
                        }
                        _ = fi27
                    }
                    if false {
                        if true {
                            fmt.Println(v13, v5)
                            var v33 float32 = 0.1
                            for fi34 := 0; fi34 < 6; fi34++ {
                                switch 1 {
                                case 117:
                                    if true {
                                        _ = len(v2)
                                        var v35 int8 = -1
                                        v36 := GenS1{F0: 128, F1: 1, F2: 255}
                                        _ = v35
                                        _ = v36
                                    }
                                case -884:
                                    v37 := GenS4{F0: 0, F1: 2}
                                    _ = v37
                                case -1:
                                    for fi38 := 0; fi38 < 6; fi38++ {
                                        v33 = 3.14159
                                        fmt.Println(v18)
                                        v5 = ((v5 * 0.0) / ((3.14159 - (v5 - float64((v33 + float32((v5 / 1.5)))))) / float64(v33)))
                                        switch v4 {
                                        case "\n\t":
                                            v13 -= v13
                                        case "foo bar":
                                            var v39 float32 = v3
                                            _ = v39
                                        case "hello":
                                            v40 := GenS1{F0: fi34, F1: 123, F2: 128}
                                            _ = v40
                                        }
                                        _ = fi38
                                    }
                                }
                                switch v4 {
                                case "a":
                                    for fi41 := 0; fi41 < 10; fi41++ {
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v33)
                                        }
                                        switch -2147483648 {
                                        case 127:
                                            v13 += v13
                                        case 32768:
                                            fmt.Println(fi41, fi41, fi41)
                                        }
                                        _ = fi41
                                    }
                                case "\n\t":
                                    v42 := GenS1{F0: (fi34 / (((fi34 / ((int(v13)) | 1))) | 1)), F1: int8(v26), F2: 127}
                                    _ = v42
                                }
                                var v43 float32 = (3.14159 / (float32(float64((v3 / v33))) - 0.1))
                                _ = fi34
                                _ = v43
                            }
                            _ = v33
                        } else {
                            v18 += (v18 >> ((18446744073709551615) & 63))
                        }
                        if (18446744073709551615 <= 42) {
                            if false {
                                for fi44 := 0; fi44 < 7; fi44++ {
                                    switch "hello" {
                                    case "\n\t":
                                        v13 -= (-(v13))
                                    default:
                                        if false {
                                            for fi45 := 0; fi45 < 8; fi45++ {
                                                v46 := GenS5{F0: 4294967295}
                                                var v47 int16 = (v13 >> ((2) & 63))
                                                v1 = append(v1, (fi45 << ((159) & 63)))
                                                _ = fi45
                                                _ = v46
                                                _ = v47
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                    var v48 string = v4
                                    var v49 string = v48
                                    _ = fi44
                                    _ = v48
                                    _ = v49
                                }
                                if true {
                                    if true {
                                        v5 += v5
                                    }
                                    if (!false) {
                                        v50 := GenS1{F0: 42, F1: 117, F2: 2}
                                        _ = v50
                                    }
                                    switch "" {
                                    case "\n\t":
                                        for fi51 := 0; fi51 < 19; fi51++ {
                                            v52 := v2[fi51]
                                            _ = fi51
                                            _ = v52
                                        }
                                    case "a":
                                        for fi53 := 0; fi53 < 8; fi53++ {
                                            switch fi53 {
                                            case -9223372036854775808:
                                                v4 = v4
                                            default:
                                                switch 778 {
                                                case 2147483647:
                                                    for fi54 := 0; fi54 < 16; fi54++ {
                                                        var v55 uint32 = 2147483648
                                                        var v56 int8 = -31
                                                        _ = fi54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                case 800:
                                                    if true {
                                                        for fi57 := 0; fi57 < 12; fi57++ {
                                                            v26 += v26
                                                            _ = fi57
                                                        }
                                                        v18 += 32767
                                                    }
                                                }
                                            }
                                            v13 += (v13 >> ((0) & 63))
                                            _ = fi53
                                        }
                                    }
                                }
                            } else {
                                switch -276 {
                                case 128:
                                    var v58 float32 = (v3 * (float32(v5) - (v3 / v3)))
                                    _ = v58
                                case 256:
                                    var v59 int8 = 42
                                    _ = v59
                                default:
                                    v26 += -1
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v60 := GenS6{F0: 1, F1: -1}
                            _ = v60
                        }
                        if true {
                            var v61 uint8 = 255
                            var v62 uint32 = 550
                            _ = v61
                            _ = v62
                        } else {
                            fmt.Println(v5, v4, v5)
                            for fi63 := 0; fi63 < 1; fi63++ {
                                for fi64 := 0; fi64 < 14; fi64++ {
                                    switch fi64 {
                                    case 0:
                                        if (((v4 + v4) >= ((v4 + v4) + v4)) || false) {
                                            fmt.Println(v26)
                                            var v65 uint8 = 128
                                            var v66 bool = false
                                            _ = v65
                                            _ = v66
                                        }
                                    case 4294967295:
                                        fmt.Println(v26, v26)
                                    }
                                    _ = fi64
                                }
                                fmt.Println(v26)
                                _ = fi63
                            }
                            fmt.Println(v26, v18, v5)
                        }
                    } else {
                        delete(v2, 9223372036854775807)
                        switch 4294967295 {
                        case 32767:
                            switch -2147483648 {
                            case 307:
                                switch 9223372036854775807 {
                                case 9223372036854775807:
                                    _ = len(v1)
                                case -2147483648:
                                    for fi67 := 0; fi67 < 18; fi67++ {
                                        v4 = ("foo bar" + v4)
                                        _ = fi67
                                    }
                                default:
                                    var v68 float32 = float32(((0.1 / v5) / float64(v3)))
                                    _ = v68
                                }
                            case -575:
                                var v69 int64 = v26
                                _ = v69
                            }
                        case 509:
                            v5 += v5
                        }
                    }
                    _ = v26
                }
            }
        }
    case "foo bar":
        v1 = append(v1, 2147483647)
    }
    v5 = ((0.1 * 0.1) * v5)
    v18 -= v18
    switch ("foo bar" + "a") {
    case "a":
        if false {
            fmt.Println(v5, v18, v13)
        }
    case "":
        fmt.Println(v5, v3, v18)
    }
    for fi70 := 0; fi70 < 19; fi70++ {
        var v71 int8 = 28
        v3 += v3
        v13 += v13
        _ = fi70
        _ = v71
    }
    fmt.Println(v18, v3)
    if true {
        v72 := GenS2{F0: 1, F1: false, F2: 4294967295}
        _ = v72
    } else {
        v4 = v4
        v73 := GenS3{F0: 1, F1: 2147483647, F2: 65535}
        _ = v73
    }
    var v74 float64 = v5
    switch 42 {
    case 32768:
        var v75 uint32 = 42
        _ = v75
    default:
        fmt.Println(v74, v74)
    }
    var v76 int8 = 62
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v76
    _ = v8
}
