package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint32; F2 int }

func main() {
    v1 := []int{}
    var v2 int8 = 42
    var v3 int = -1
    var v4 int16 = 571
    v3 += 0
    var v5 int8 = v2
    v6 := GenS1{F0: -947, F1: 1, F2: v4}
    v5 = 100
    fmt.Println(v3, v5, v2)
    if true {
        var v7 int8 = (v2 - v5)
        for fi8 := 0; fi8 < 7; fi8++ {
            v4 -= 270
            v9 := GenS2{F0: 1, F1: 42, F2: v7}
            switch "a" {
            case "":
                var v10 int64 = 42
                _ = v10
            case "a":
                v7 -= v2
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    if false {
        var v11 uint8 = 0
        _ = v11
    }
    var v12 bool = false
    v4 += v4
    var v13 float64 = (1e10 / 3.14159)
    v2 = 127
    v13 = v13
    switch (v3 << ((255) & 63)) {
    case 9223372036854775807:
        if (int64(v4) >= 1) {
            var v14 uint8 = 13
            fmt.Println(v12, v12)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v14
        }
    case 524:
        switch (("foo bar" + "hello") + ("" + "")) {
        case "foo bar":
            if true {
                for fi15 := 0; fi15 < 18; fi15++ {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 int8 = v2
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        switch fi18 {
                        case 2147483647:
                            v19 := GenS2{F0: 670, F1: 0, F2: v5}
                            _ = v19
                        case 32767:
                            var v20 bool = (v4 == v4)
                            _ = v20
                        case 9223372036854775807:
                            var v21 float64 = (0.1 * v13)
                            _ = v21
                        default:
                            v3 = 2147483647
                        }
                        _ = fi18
                    }
                    _ = fi15
                }
                switch -1 {
                case 128:
                    v1 = append(v1, (v3 << ((877) & 63)))
                case 494:
                    for fi22 := 0; fi22 < 7; fi22++ {
                        v12 = (false && ((v4 + (v4 | (v4 / ((v4) | 1)))) > v4))
                        for fi23 := 0; fi23 < 10; fi23++ {
                            if true {
                                if ((v13 - v13) == v13) {
                                    if ((v12 && (22 > 255)) || (float32(v13) < (((((1.5 + 3.14159) + 1e10) * ((((((3.14159 * 0.0) - (1.5 * 0.1)) + -2.25) / ((-2.25 / 1.0) - (3.14159 / ((0.1 - 0.0) / 1.0)))) * 1e10) + (3.14159 / -2.25))) / float32((v13 - v13))) - (float32(v13) / 1.5)))) {
                                        var v24 int = fi23
                                        v5 -= (v2 | v5)
                                        var v25 uint16 = 2
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = -41
                                        }
                                        switch ((("a" + "hello") + "") + "hello") {
                                        case "foo bar":
                                            var v26 int16 = (v4 % (((v4 + (v4 / (((^(v4))) | 1)))) | 1))
                                            _ = v26
                                        default:
                                            v12 = (!(v12 || false))
                                        }
                                    }
                                    if (!v12) {
                                        v4 = (v4 / (((v4 / ((-121) | 1))) | 1))
                                        switch "\n\t" {
                                        case "hello":
                                            if false {
                                                var v27 uint16 = 510
                                                v28 := GenS1{F0: 1, F1: 18446744073709551615, F2: (^(v4))}
                                                switch fi22 {
                                                case 339:
                                                    if false {
                                                        v27 += v27
                                                        if (!false) {
                                                            v3 = v3
                                                            var v29 int = fi23
                                                            _ = v29
                                                        }
                                                    }
                                                default:
                                                    var v30 float32 = 3.14159
                                                    _ = v30
                                                }
                                                _ = v27
                                                _ = v28
                                            }
                                        case "foo bar":
                                            if (772 == 42) {
                                                if (((0.1 * (-2.25 - 1.5)) + 1.5) > (0.0 - 1.5)) {
                                                    if v12 {
                                                        v31 := GenS1{F0: 0, F1: 4294967295, F2: 646}
                                                        var v32 int16 = v4
                                                        v3 = fi22
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        fmt.Println(v2)
                                                        v5 -= 42
                                                        v13 = v13
                                                    }
                                                }
                                                v33 := GenS2{F0: int64(v4), F1: 42, F2: (v2 & v2)}
                                                v34 := GenS1{F0: -118, F1: 149, F2: v4}
                                                v35 := GenS1{F0: 42, F1: 1, F2: v4}
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            }
                                        default:
                                            var v36 int = (^(fi22))
                                            _ = v36
                                        }
                                    } else {
                                        if v12 {
                                            v1 = append(v1, fi23)
                                            if v12 {
                                                var v37 int32 = 371
                                                var v38 int = v3
                                                _ = v37
                                                _ = v38
                                            }
                                            v3 = (v3 ^ -2147483648)
                                        }
                                        v3 = (fi22 | -294)
                                        var v39 uint = 936
                                        _ = v39
                                    }
                                    switch 127 {
                                    case 127:
                                        v4 = v4
                                    }
                                }
                                v1 = append(v1, (fi23 ^ (v3 & (^(fi22)))))
                                v40 := GenS1{F0: -204, F1: 18446744073709551615, F2: 1}
                                _ = v40
                            }
                            var v41 uint32 = 42
                            _ = fi23
                            _ = v41
                        }
                        _ = fi22
                    }
                case -245:
                    if len(v1) > 0 {
                        v1[0] = (v3 & (v3 + (v3 % ((-75) | 1))))
                    }
                default:
                    switch 32767 {
                    case -9223372036854775808:
                        if len(v1) > 0 {
                            v1[0] = v3
                        }
                    case -962:
                        switch "hello" {
                        case "":
                            if (-1 >= v2) {
                                v42 := GenS2{F0: -93, F1: 180, F2: v2}
                                _ = v42
                            } else {
                                if ((v4 ^ v4) != (v4 & v4)) {
                                    var v43 int8 = v2
                                    _ = v43
                                } else {
                                    v12 = v12
                                }
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    var v45 string = "a"
                                    _ = fi44
                                    _ = v45
                                }
                                v46 := GenS1{F0: 1, F1: 4294967295, F2: (v4 ^ -125)}
                                _ = v46
                            }
                        default:
                            for fi47 := 0; fi47 < 20; fi47++ {
                                var v48 int32 = -1
                                _ = fi47
                                _ = v48
                            }
                        }
                    case 830:
                        var v49 int32 = -1
                        _ = v49
                    default:
                        switch (v3 >> ((895) & 63)) {
                        case 32768:
                            fmt.Println(v12, v4, v5)
                        }
                    }
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v3, v13)
                v50 := GenS3{F0: 388, F1: 95}
                _ = v50
            }
        case "hello":
            var v51 int16 = -32768
            _ = v51
        default:
            if (!v12) {
                var v52 float64 = v13
                for fi53 := 0; fi53 < 20; fi53++ {
                    var v54 float32 = (1.5 / 1.0)
                    for fi55 := 0; fi55 < 15; fi55++ {
                        fmt.Println(v13, v54)
                        _ = fi55
                    }
                    v54 -= ((3.14159 * float32(((v13 + (v13 * v52)) - v52))) - v54)
                    _ = fi53
                    _ = v54
                }
                _ = v52
            }
        }
    case 127:
        switch ("\n\t" + "a") {
        case "":
            v2 += (v2 + (^(v5)))
        case "\n\t":
            switch (v3 >> ((928) & 63)) {
            case 0:
                v1 = append(v1, (v3 >> ((2) & 63)))
            case -1:
                switch v3 {
                case -243:
                    v3 = v3
                default:
                    if true {
                        if (!v12) {
                            v56 := GenS4{F0: (v4 * 42), F1: 1, F2: v3}
                            _ = v56
                        }
                        switch (v3 + v3) {
                        case 951:
                            fmt.Println(v4)
                        }
                        v12 = (47 > 65535)
                    }
                }
            case 32767:
                for fi57 := 0; fi57 < 10; fi57++ {
                    var v58 uint8 = 2
                    for fi59 := 0; fi59 < 17; fi59++ {
                        var v60 int16 = v4
                        var v61 uint16 = 255
                        if v12 {
                            v1 = append(v1, (v3 / ((fi59) | 1)))
                            fmt.Println(fi57)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    _ = fi57
                    _ = v58
                }
            default:
                v12 = v12
            }
        default:
            switch "foo bar" {
            case "\n\t":
                var v62 uint8 = 135
                _ = v62
            case "foo bar":
                var v63 int8 = v2
                _ = v63
            case "":
                v64 := GenS2{F0: 648, F1: 2, F2: (v2 << ((4294967295) & 63))}
                _ = v64
            default:
                v5 = (v2 / ((106) | 1))
            }
        }
    default:
        var v65 int32 = int32(v5)
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
