package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint; F2 int32 }
type GenS4 struct { F0 int; F1 int8; F2 int32 }

func main() {
    var v1 int16 = 439
    v2 := []uint32{65535}
    var v3 int16 = v1
    var v4 uint64 = 18446744073709551615
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 float32 = 3.14159
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 int = fi5
                switch ("hello" + "") {
                case "":
                    v1 += int16(v8)
                case "foo bar":
                    var v9 string = "a"
                    _ = v9
                }
                var v10 int16 = v3
                if ("foo bar" >= "hello") {
                    var v11 bool = false
                    fmt.Println(fi7, v8)
                    _ = v11
                }
                _ = fi7
                _ = v10
                _ = v8
            }
            if (!false) {
                if len(v2) > 0 {
                    v2[0] = 110
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        } else {
            v2 = append(v2, 2147483647)
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(fi5, v4, fi12)
                _ = fi12
            }
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(v3)
    if (!false) {
        v13 := GenS1{F0: 403}
        _ = len(v2)
        v14 := GenS1{F0: 42}
        _ = v13
        _ = v14
    }
    if false {
        if false {
            var v15 int16 = 1000
            var v16 int = 9223372036854775807
            var v17 float32 = 1.5
            _ = v15
            _ = v16
            _ = v17
        } else {
            fmt.Println(v3, v4, v4)
            var v18 float32 = 1e10
            v3 -= (v1 >> ((1) & 63))
            _ = v18
        }
    } else {
        if false {
            v3 += 1
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int16 = (v1 + v1)
                switch 0 {
                case 42:
                    if (!false) {
                        v4 += (v4 >> ((2) & 63))
                        v21 := GenS1{F0: 42}
                        _ = len(v2)
                        _ = v21
                    } else {
                        v4 -= v4
                        _ = len(v2)
                    }
                case 1:
                    for fi22 := 0; fi22 < 2; fi22++ {
                        v23 := GenS1{F0: 42}
                        fmt.Println(v4, v3, fi22)
                        if true {
                            if true {
                                var v24 int = (fi22 ^ fi19)
                                _ = v24
                            }
                            var v25 int32 = -931
                            _ = v25
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v20 -= -999
                                    v3 -= (v20 / ((1000) | 1))
                                }
                            case "":
                                var v26 int8 = 42
                                _ = v26
                            case "hello":
                                switch "\n\t" {
                                case "hello":
                                    switch fi22 {
                                    case 42:
                                        v27 := GenS1{F0: 795}
                                        _ = v27
                                    case 381:
                                        switch "foo bar" {
                                        case "":
                                            v1 = v20
                                        case "hello":
                                            fmt.Println(v1)
                                        case "foo bar":
                                            v1 += (v3 + (v1 + 0))
                                        default:
                                            v1 += v1
                                        }
                                    case -1:
                                        v28 := GenS1{F0: 4294967295}
                                        _ = v28
                                    }
                                case "":
                                    if true {
                                        var v29 uint16 = 0
                                        switch -2147483648 {
                                        case 128:
                                            var v30 int32 = 2147483647
                                            _ = v30
                                        case 835:
                                            v31 := GenS2{F0: 9223372036854775807}
                                            _ = v31
                                        default:
                                            var v32 int32 = int32(v4)
                                            _ = v32
                                        }
                                        _ = v29
                                    }
                                default:
                                    v33 := GenS2{F0: 96}
                                    _ = v33
                                }
                            default:
                                if len(v2) > 0 {
                                    v2[0] = 217
                                }
                            }
                            fmt.Println(v20, v1, fi19)
                        }
                        v4 += (v4 >> ((444) & 63))
                        switch -379 {
                        case -2147483648:
                            var v34 string = ("" + "foo bar")
                            _ = v34
                        case 32767:
                            switch "" {
                            case "a":
                                var v35 uint = 4294967295
                                _ = v35
                            case "foo bar":
                                var v36 uint32 = 2147483648
                                _ = v36
                            default:
                                var v37 uint8 = 128
                                _ = v37
                            }
                        case 328:
                            switch fi22 {
                            case 2147483647:
                                v2 = append(v2, 1)
                            }
                        default:
                            if (!true) {
                                var v38 int8 = -1
                                _ = v38
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                case 32768:
                    var v39 float64 = 0.0
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 14; fi40++ {
                        switch fi19 {
                        case 32767:
                            v4 = 512
                        case 205:
                            switch (fi19 << ((2) & 63)) {
                            case 32768:
                                v4 = v4
                                v1 -= (v20 & v1)
                                v2 = append(v2, uint32(v4))
                            }
                        case -36:
                            if (false || false) {
                                v4 += (v4 >> ((42) & 63))
                            }
                        default:
                            v1 += -32768
                        }
                        v1 = -414
                        v4 = (^(v4))
                        _ = fi40
                    }
                }
                _ = fi19
                _ = v20
            }
            for fi41 := 0; fi41 < 11; fi41++ {
                v42 := GenS1{F0: 2}
                v43 := GenS3{F0: 1, F1: 2, F2: -277}
                _ = fi41
                _ = v42
                _ = v43
            }
            var v44 uint32 = 0
            _ = v44
        }
        if true {
            v1 = (v1 | (v1 ^ -32768))
            v4 += 32
        } else {
            v3 = v3
            var v45 uint8 = 42
            _ = v45
        }
    }
    v1 += v3
    var v46 float32 = ((((0.1 + 0.0) + (-2.25 * (3.14159 - 1.5))) / 1.0) + 3.14159)
    var v47 uint16 = 42
    fmt.Println(v46, v1, v46)
    if (!(42 <= 100)) {
        var v48 float32 = (v46 / v46)
        _ = v48
    }
    v49 := GenS3{F0: 18446744073709551615, F1: 4294967295, F2: -1}
    v50 := GenS3{F0: 0, F1: 2, F2: 42}
    for fi51 := 0; fi51 < 2; fi51++ {
        for fi52 := 0; fi52 < 16; fi52++ {
            var v53 bool = (!(9223372036854775807 > 42))
            switch 128 {
            case -1:
                var v54 uint = 2
                _ = v54
            default:
                if v53 {
                    var v55 int32 = 42
                    v55 += v55
                    _ = v55
                } else {
                    if (!(!(v1 < v1))) {
                        if (true && (42 > -54)) {
                            v1 = (v1 | (^(v1)))
                            v53 = false
                        } else {
                            var v56 float32 = 3.14159
                            _ = v56
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v57 uint32 = 589
                        fmt.Println(fi52)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v57
                    }
                    switch "" {
                    case "a":
                        if v53 {
                            var v58 int16 = (v3 + (v1 | 371))
                            _ = v58
                        }
                    default:
                        var v59 int8 = -128
                        _ = v59
                    }
                    var v60 uint64 = 2
                    v2 = append(v2, 571)
                    if ((true && (!(-314 != 1))) || (0 < 2147483647)) {
                        var v61 float64 = ((0.1 - (0.1 * ((0.0 * (((0.1 + 1e10) / -2.25) / (1e10 + ((-2.25 * (0.1 + -2.25)) / 1e10)))) + (1e10 - 1.5)))) / 1.0)
                        _ = v61
                    }
                    if (42 < 0) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v62 int = (fi51 << ((175) & 63))
                        v4 -= (^(v4))
                        _ = len(v2)
                        _ = v62
                    } else {
                        var v63 float32 = v46
                        fmt.Println(fi51, v63, v60)
                        if false {
                            _ = len(v2)
                            v4 += 691
                            var v64 uint64 = 4294967295
                            v65 := GenS3{F0: 65535, F1: 0, F2: -1}
                            var v66 float32 = 0.0
                            var v67 uint8 = 127
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v3 -= 1000
                            _ = v64
                            _ = v65
                            _ = v66
                            _ = v67
                        }
                        if true {
                            if ((0.0 * float32(3.14159)) == v46) {
                                var v68 string = "\n\t"
                                _ = v68
                            }
                            v47 -= (v47 ^ 60)
                        } else {
                            fmt.Println(v3)
                        }
                        _ = v63
                    }
                    var v69 uint16 = v47
                    switch "a" {
                    case "hello":
                        switch "\n\t" {
                        case "foo bar":
                            if v53 {
                                switch "hello" {
                                case "":
                                    var v70 int64 = 965
                                    _ = v70
                                case "foo bar":
                                    _ = len(v2)
                                }
                                if v53 {
                                    v46 += (v46 * (v46 - (-2.25 / -2.25)))
                                } else {
                                    var v71 uint64 = 953
                                    var v72 float64 = -2.25
                                    _ = v71
                                    _ = v72
                                }
                                switch fi52 {
                                case 2147483647:
                                    fmt.Println(fi52, v3, v69)
                                case -99:
                                    fmt.Println(fi52)
                                case 128:
                                    v2 = append(v2, 897)
                                }
                            }
                        case "\n\t":
                            v73 := GenS2{F0: -9223372036854775808}
                            _ = v73
                        case "":
                            var v74 int16 = v3
                            _ = v74
                        }
                    case "":
                        v69 = 32767
                    default:
                        var v75 int64 = -474
                        _ = v75
                    }
                    var v76 float64 = (0.0 * 1.5)
                    _ = v60
                    _ = v69
                    _ = v76
                }
            }
            v2 = append(v2, 836)
            _ = fi52
            _ = v53
        }
        for fi77 := 0; fi77 < 15; fi77++ {
            var v78 bool = (42 == 128)
            _ = fi77
            _ = v78
        }
        v1 = 438
        _ = fi51
    }
    for fi79 := 0; fi79 < 3; fi79++ {
        if true {
            v2 = append(v2, 653)
        }
        var v80 bool = (!(false && true))
        v1 -= v1
        _ = fi79
        _ = v80
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v3)
    case "a":
        if false {
            var v81 string = ("a" + (("" + ("a" + "\n\t")) + (("a" + "") + "")))
            _ = v81
        }
    }
    if (((0.0 - (1e10 * 1.5)) * 0.1) > 0.1) {
        if true {
            v82 := GenS3{F0: 0, F1: 255, F2: 2147483647}
            if true {
                for fi83 := 0; fi83 < 15; fi83++ {
                    switch int(v4) {
                    case 127:
                        if true {
                            for fi84 := 0; fi84 < 17; fi84++ {
                                v46 = v46
                                var v85 bool = false
                                fmt.Println(v1)
                                _ = fi84
                                _ = v85
                            }
                            var v86 float32 = v46
                            var v87 uint = 65535
                            _ = v86
                            _ = v87
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                    }
                    _ = fi83
                }
                v4 = v4
                if len(v2) > 0 {
                    v2[0] = 939
                }
            } else {
                v47 += v47
            }
            switch "a" {
            case "hello":
                v4 = 18446744073709551615
            case "a":
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
            case "\n\t":
                v88 := GenS4{F0: 1, F1: 49, F2: -1}
                _ = v88
            }
            _ = v82
        } else {
            v4 -= (v4 ^ v4)
            var v89 bool = false
            _ = v89
        }
        v1 = 32767
    } else {
        _ = len(v2)
    }
    if ((-2147483648 < -2147483648) && (uint16(v4) > v47)) {
        fmt.Println(v47)
        if false {
            v46 -= (0.1 * v46)
            var v90 uint32 = 694
            _ = v90
        }
    } else {
        v2 = append(v2, 2)
    }
    if (1 >= 1) {
        switch 247 {
        case -1:
            for fi91 := 0; fi91 < 18; fi91++ {
                v2 = append(v2, 80)
                var v92 uint = 515
                _ = fi91
                _ = v92
            }
        case 757:
            var v93 int32 = 710
            _ = v93
        case -768:
            v3 = 1000
        }
        v2 = append(v2, 2147483648)
    } else {
        v2 = append(v2, 4294967295)
        v47 += v47
        var v94 uint16 = (v47 / (((^(v47))) | 1))
        _ = v94
    }
    v4 -= (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v50
}
