package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int32; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int64 = -514
    v2 := make(map[int]int)
    var v3 uint16 = 255
    var v4 uint = 925
    var v5 uint16 = v3
    v6 := GenS1{F0: v3}
    var v7 uint = v4
    if (65535 == 832) {
        if ((v5 * (v3 - (^(v5)))) >= v5) {
            var v8 string = "\n\t"
            _ = v8
        } else {
            v3 = (v5 * v3)
            var v9 int32 = 0
            var v10 int64 = v1
            _ = v10
            _ = v9
        }
        if (!(!(1e10 >= 3.14159))) {
            var v11 int16 = -97
            v7 = (^(v4))
            _ = v11
        }
    } else {
        var v12 uint8 = 2
        v7 = v4
        v13 := GenS2{F0: v12, F1: false, F2: 1}
        _ = v12
        _ = v13
    }
    switch 2147483647 {
    case 4294967295:
        switch "\n\t" {
        case "hello":
            v1 -= 9223372036854775807
        case "\n\t":
            switch (("hello" + (("" + ("foo bar" + "foo bar")) + "hello")) + "") {
            case "a":
                fmt.Println(v3)
            case "\n\t":
                switch 716 {
                case 42:
                    v2[32767] = 42
                case -9223372036854775808:
                    v1 += v1
                default:
                    var v14 uint16 = (v5 - 32767)
                    _ = v14
                }
            default:
                var v15 uint64 = 18446744073709551615
                _ = v15
            }
        case "a":
            for fi16 := 0; fi16 < 18; fi16++ {
                _ = len(v2)
                _ = fi16
            }
        default:
            v2[-939] = 0
        }
    case 64:
        v17 := GenS1{F0: 2}
        _ = v17
    default:
        switch -2147483648 {
        case 711:
            v4 = uint(v5)
        case 384:
            v5 += 727
        case -1:
            v18 := []int{}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v7 = (v7 + (v7 * (v7 * (v7 & (v7 % (((v7 << (((v4 | (v7 ^ (v7 | (v7 - 1))))) & 63))) | 1))))))
        v20 := []int{fi19}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 uint64 = 42
        _ = fi21
        _ = v22
    }
    var v23 uint16 = 255
    fmt.Println(v3, v1)
    delete(v2, 127)
    if ((!(((1e10 + 1.5) + (1.5 + 0.0)) <= 0.0)) && true) {
        v4 -= v7
        var v24 uint32 = 65535
        for fi25 := 0; fi25 < 11; fi25++ {
            switch fi25 {
            case 0:
                switch (fi25 >> (((v7 - 389)) & 63)) {
                case -525:
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v24 = v24
                        var v27 int32 = 42
                        if (true || (0.0 > ((3.14159 + 3.14159) + 1e10))) {
                            v28 := GenS1{F0: v23}
                            _ = v28
                        }
                        _ = fi26
                        _ = v27
                    }
                case 32768:
                    v29 := GenS3{F0: fi25}
                    _ = v29
                }
            case 42:
                fmt.Println(v7, v24)
            default:
                for fi30 := 0; fi30 < 16; fi30++ {
                    v3 -= v3
                    var v31 uint16 = (^(v3))
                    _ = fi30
                    _ = v31
                }
            }
            var v32 uint8 = 2
            var v33 uint64 = uint64(v24)
            _ = fi25
            _ = v32
            _ = v33
        }
        _ = v24
    }
    v34 := []int{579}
    switch 228 {
    case 4294967295:
        if len(v34) > 0 {
            v34[0] = 128
        }
    case -1:
        v35 := GenS3{F0: 32768}
        _ = v35
    case 9223372036854775807:
        v36 := GenS4{F0: (v1 + v1)}
        _ = v36
    default:
        switch "foo bar" {
        case "hello":
            var v37 string = "hello"
            _ = v37
        case "\n\t":
            var v38 uint16 = v23
            _ = v38
        default:
            v3 = v23
        }
    }
    if false {
        v7 += (^(v7))
        switch "" {
        case "foo bar":
            var v39 int8 = -128
            _ = v39
        case "a":
            v40 := GenS5{F0: -2147483648, F1: 896, F2: -119}
            _ = v40
        default:
            v41 := v2[2147483647]
            _ = v41
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            v43 := GenS4{F0: 843}
            fmt.Println(v3, v23)
            if true {
                for fi44 := 0; fi44 < 7; fi44++ {
                    v45 := GenS6{F0: 367}
                    v46 := GenS3{F0: (fi44 * (fi42 / ((fi42) | 1)))}
                    if (!true) {
                        if false {
                            if len(v34) > 0 {
                                v34[0] = (fi44 >> (((v7 * 1)) & 63))
                            }
                            var v47 uint = 65535
                            switch "hello" {
                            case "\n\t":
                                if len(v34) >= 2 {
                                    v34 = v34[1:len(v34)]
                                }
                            case "foo bar":
                                switch (-(fi42)) {
                                case -275:
                                    if (true || (false && false)) {
                                        _ = len(v34)
                                        var v48 uint8 = 2
                                        if ("" == (("\n\t" + "a") + (("" + "hello") + (("foo bar" + (("\n\t" + "a") + "hello")) + "")))) {
                                            if false {
                                                var v49 int = 42
                                                _ = v49
                                            }
                                            var v50 uint32 = 4294967295
                                            if false {
                                                switch -658 {
                                                case 773:
                                                    v47 = (v47 >> (((v7 + 255)) & 63))
                                                case 1:
                                                    switch -662 {
                                                    case 0:
                                                        var v51 uint16 = v3
                                                        if false {
                                                            var v52 bool = (fi42 <= (fi42 | 127))
                                                            var v53 int16 = 784
                                                            if (!true) {
                                                                fmt.Println(v4, v5, v1)
                                                                var v54 int8 = int8(v23)
                                                                if len(v34) > 0 {
                                                                    v34[0] = (^(fi42))
                                                                }
                                                                if len(v34) > 0 {
                                                                    v34[0] = (fi42 >> ((uint(v50)) & 63))
                                                                }
                                                                var v55 float64 = 1e10
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        v7 = v7
                                                        _ = v51
                                                    case -994:
                                                        var v56 uint16 = (v23 + 2)
                                                        _ = v56
                                                    }
                                                case 9223372036854775807:
                                                    var v57 uint32 = 2
                                                    _ = v57
                                                default:
                                                    _ = len(v34)
                                                    var v58 int = int(v47)
                                                    _ = v58
                                                }
                                            }
                                            _ = v50
                                        }
                                        _ = v48
                                    } else {
                                        var v59 int32 = 42
                                        delete(v2, -297)
                                        _ = v59
                                    }
                                }
                            case "a":
                                if false {
                                    v1 = (v1 << ((42) & 63))
                                } else {
                                    v7 += v47
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v60 uint = v7
                                        _ = v60
                                    case "":
                                        v4 = (v47 << ((v7) & 63))
                                    case "hello":
                                        _ = len(v2)
                                    }
                                    v61 := GenS7{F0: int64(v47)}
                                    v34 = append(v34, fi42)
                                    _ = v61
                                }
                            }
                            _ = v47
                        }
                        var v62 int32 = 42
                        _ = v62
                    } else {
                        var v63 uint = (v7 << ((v4) & 63))
                        v2[(fi42 - (fi42 % (((fi44 ^ 4294967295)) | 1)))] = (fi42 << (((v63 | v63)) & 63))
                        if true {
                            var v64 uint16 = (v5 + 776)
                            v1 = -997
                            fmt.Println(fi44, v4)
                            _ = v64
                        }
                        _ = v63
                    }
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            _ = fi42
            _ = v43
        }
    }
    for fi65 := 0; fi65 < 8; fi65++ {
        var v66 uint = (v7 * 4294967295)
        _ = fi65
        _ = v66
    }
    if true {
        _ = len(v2)
        switch -2147483648 {
        case 191:
            v67 := GenS7{F0: 9223372036854775807}
            _ = v67
        }
        for fi68 := 0; fi68 < 16; fi68++ {
            v69 := GenS8{F0: (false || (-1 != -2147483648)), F1: uint32(v7)}
            _ = fi68
            _ = v69
        }
    } else {
        var v70 float64 = (1e10 / (3.14159 / 1e10))
        if false {
            v71 := v2[538]
            _ = v71
        } else {
            switch -2147483648 {
            case 42:
                fmt.Println(v1)
            case -575:
                v23 -= v5
            case 2147483647:
                var v72 int64 = v1
                _ = v72
            default:
                if ((v3 + v23) != (v5 + v23)) {
                    if ((-257 < -32768) || (!false)) {
                        var v73 float64 = 0.1
                        if false {
                            v70 = v70
                        } else {
                            v4 += v7
                            var v74 float32 = 1.5
                            _ = v74
                        }
                        _ = v73
                    }
                    var v75 int32 = 0
                    if true {
                        v7 = v4
                        for fi76 := 0; fi76 < 5; fi76++ {
                            switch 128 {
                            case 32767:
                                if (fi76 == fi76) {
                                    var v77 int16 = 42
                                    _ = v77
                                }
                            }
                            switch (fi76 / ((fi76) | 1)) {
                            case -1:
                                v23 = v3
                            case 4294967295:
                                v70 = v70
                            default:
                                v23 = v23
                            }
                            v5 = (v5 / (((v23 << (((v4 << (((v7 - v4)) & 63))) & 63))) | 1))
                            _ = fi76
                        }
                        for fi78 := 0; fi78 < 13; fi78++ {
                            fmt.Println(fi78, v3)
                            _ = fi78
                        }
                    }
                    _ = v75
                }
            }
            var v79 uint = uint(v3)
            if len(v34) > 0 {
                v34[0] = 127
            }
            _ = v79
        }
        v3 = (^(v3))
        _ = v70
    }
    v80 := v2[127]
    v7 = (v4 + 255)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v80
}
