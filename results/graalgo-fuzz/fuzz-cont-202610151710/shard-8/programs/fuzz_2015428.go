package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 int32; F1 uint16; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint; F2 int16 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 string = "a"
    v3 := make(map[int]string)
    fmt.Println(v2)
    v1[v2] = false
    var v4 uint16 = 32767
    v5 := GenS1{F0: 14, F1: 127, F2: (634 > 2147483648)}
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: -1, F1: (fi6 | (fi6 & 9223372036854775807)), F2: (true || true)}
        if (!true) {
            var v8 uint16 = (^(v4))
            switch (v2 + "a") {
            case "a":
                v2 = (v2 + "")
            case "foo bar":
                v8 = (v8 % (((^(v8))) | 1))
            }
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 bool = true
                    switch fi6 {
                    case 2147483647:
                        fmt.Println(v10, v8, v10)
                    case 127:
                        switch v2 {
                        case "foo bar":
                            v11 := GenS1{F0: (^(fi6)), F1: (fi9 << ((255) & 63)), F2: ((-1 == 0) && true)}
                            _ = v11
                        case "":
                            var v12 float64 = 0.0
                            _ = v12
                        case "a":
                            var v13 float32 = 0.1
                            _ = v13
                        default:
                            v14 := []int{-1, fi9, (fi9 / ((312) | 1))}
                            _ = v14
                        }
                    case -755:
                        fmt.Println(v4)
                    }
                    v15 := GenS1{F0: 1, F1: (-(fi6)), F2: false}
                    var v16 uint64 = 1
                    var v17 float64 = (0.0 / 1.5)
                    _ = len(v1)
                    var v18 int32 = -2147483648
                    if (v10 && true) {
                        v19 := GenS2{F0: (v18 - 677), F1: 255, F2: -128}
                        _ = v19
                    }
                    _ = fi9
                    _ = v10
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                fmt.Println(v8, fi6, v2)
            }
            _ = v8
        }
        if false {
            if false {
                var v20 float32 = (1.5 - 1e10)
                for fi21 := 0; fi21 < 15; fi21++ {
                    v2 = v2
                    _ = fi21
                }
                _ = v20
            }
        } else {
            v4 += v4
            v4 = (v4 | 65535)
            var v22 float64 = (1.5 - 1e10)
            _ = v22
        }
        _ = fi6
        _ = v7
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v4, v2, v2)
        var v24 int = 128
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 float32 = 0.1
        fmt.Println(v4, v26, v4)
        for fi27 := 0; fi27 < 11; fi27++ {
            switch 1 {
            case 0:
                var v28 int = (fi27 - fi27)
                _ = v28
            case 1:
                switch (^(fi25)) {
                case 9223372036854775807:
                    fmt.Println(v26, v4, fi25)
                default:
                    v4 += (v4 | (v4 >> ((0) & 63)))
                    v29 := GenS1{F0: (fi25 ^ (fi25 * fi25)), F1: (fi27 & fi25), F2: true}
                    _ = v29
                }
            }
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            if (true || (255 < 1)) {
                var v32 int = (-(fi31))
                var v33 int = fi30
                if (("foo bar" + v2) == (v2 + (v2 + "hello"))) {
                    if true {
                        fmt.Println(v32, fi30, v32)
                    } else {
                        switch fi30 {
                        case -1:
                            v34 := GenS3{F0: false}
                            _ = v34
                        }
                    }
                    v32 = 1
                    v35 := GenS1{F0: (v33 & (^(v32))), F1: (v33 + 32767), F2: (0.0 <= 0.1)}
                    _ = v35
                } else {
                    delete(v3, -1)
                    fmt.Println(fi31, fi31)
                    var v36 int = v33
                    _ = v36
                }
                _ = v32
                _ = v33
            }
            _ = fi31
        }
        if false {
            for fi37 := 0; fi37 < 11; fi37++ {
                if true {
                    var v38 float64 = (-2.25 * 3.14159)
                    _ = v38
                }
                _ = fi37
            }
            var v39 int8 = 26
            _ = v39
        }
        _ = fi30
    }
    v40 := []int{}
    if (false && false) {
        for fi41 := 0; fi41 < 5; fi41++ {
            v42 := GenS1{F0: fi41, F1: fi41, F2: (fi41 < -289)}
            v3[(fi41 / ((733) | 1))] = "hello"
            var v43 float64 = -2.25
            _ = fi41
            _ = v42
            _ = v43
        }
    } else {
        if true {
            v44 := GenS2{F0: 0, F1: (v4 | (v4 ^ 83)), F2: 858}
            for fi45 := 0; fi45 < 20; fi45++ {
                delete(v1, "")
                if false {
                    switch (v2 + v2) {
                    case "foo bar":
                        if (0.1 >= float32(1.5)) {
                            switch (fi45 * 32767) {
                            case 822:
                                if len(v40) >= 2 {
                                    v40 = v40[1:len(v40)]
                                }
                            case 127:
                                v46 := GenS3{F0: true}
                                _ = v46
                            }
                            v47 := GenS4{F0: v4, F1: 65535, F2: -32768}
                            _ = v47
                        } else {
                            for fi48 := 0; fi48 < 8; fi48++ {
                                fmt.Println(v4, v4, v4)
                                switch (^(fi48)) {
                                case 32767:
                                    v49 := GenS4{F0: (^(v4)), F1: 309, F2: 32767}
                                    _ = v49
                                }
                                if false {
                                    var v50 uint = 1
                                    v50 = 4294967295
                                    _ = v50
                                }
                                _ = fi48
                            }
                        }
                    case "\n\t":
                        switch -105 {
                        case 32767:
                            for fi51 := 0; fi51 < 16; fi51++ {
                                v4 = 656
                                _ = fi51
                            }
                        case 128:
                            v2 = v2
                        case -1:
                            for fi52 := 0; fi52 < 1; fi52++ {
                                fmt.Println(v4)
                                _ = fi52
                            }
                        }
                    case "":
                        v2 = ""
                    }
                } else {
                    delete(v3, (fi45 - (fi45 % (((-(fi45))) | 1))))
                    v4 += (v4 + v4)
                }
                if true {
                    fmt.Println(fi45, fi45)
                    fmt.Println(v2, v2, v2)
                }
                _ = fi45
            }
            _ = v44
        } else {
            if (2 != 122) {
                fmt.Println(v4)
                v4 = 488
                v4 = (v4 >> ((276) & 63))
            }
        }
    }
    for fi53 := 0; fi53 < 9; fi53++ {
        if (!true) {
            v4 += v4
            var v54 string = (("hello" + v2) + ((v2 + "") + v2))
            _ = v54
        } else {
            if len(v40) > 0 {
                v40[0] = -1
            }
        }
        for fi55 := 0; fi55 < 20; fi55++ {
            var v56 bool = false
            _ = fi55
            _ = v56
        }
        v2 = (v2 + ("a" + v2))
        _ = fi53
    }
    var v57 uint64 = 4294967295
    v3[int(v57)] = ("\n\t" + "hello")
    switch 32768 {
    case -2147483648:
        var v58 uint8 = 1
        _ = v58
    case 1:
        v40 = append(v40, 4294967295)
    }
    v4 += v4
    if (0 >= 362) {
        if false {
            switch ("a" + v2) {
            case "":
                fmt.Println(v4, v2)
            default:
                for fi59 := 0; fi59 < 4; fi59++ {
                    if ((0.0 / 1e10) <= (1e10 / -2.25)) {
                        if true {
                            if true {
                                for fi60 := 0; fi60 < 1; fi60++ {
                                    v61 := GenS4{F0: 473, F1: 0, F2: -434}
                                    var v62 uint16 = (v4 - uint16(v57))
                                    var v63 uint32 = 114
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                                v2 = "a"
                                if false {
                                    v2 = v2
                                } else {
                                    if false {
                                        fmt.Println(v4, fi59, v4)
                                        var v64 int8 = 100
                                        _ = v64
                                    } else {
                                        var v65 uint32 = 2147483648
                                        _ = v65
                                    }
                                }
                            }
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                            fmt.Println(v2, v2)
                        }
                        v57 -= (v57 << ((18446744073709551615) & 63))
                        switch (-(fi59)) {
                        case -2147483648:
                            v66 := v3[fi59]
                            _ = v66
                        }
                    } else {
                        for fi67 := 0; fi67 < 18; fi67++ {
                            v4 += (v4 >> ((781) & 63))
                            _ = fi67
                        }
                        fmt.Println(v57)
                        switch (-(fi59)) {
                        case -1:
                            v2 = v2
                        default:
                            v68 := GenS5{F0: 1, F1: 1}
                            _ = v68
                        }
                    }
                    _ = fi59
                }
            }
        } else {
            switch 1 {
            case -410:
                v2 = v2
            case 128:
                for fi69 := 0; fi69 < 18; fi69++ {
                    for fi70 := 0; fi70 < 2; fi70++ {
                        if true {
                            v2 = v2
                            v71 := GenS1{F0: -2147483648, F1: -2147483648, F2: false}
                            switch fi70 {
                            case -9223372036854775808:
                                var v72 float64 = (float64(0.0) + (3.14159 - 0.1))
                                _ = v72
                            case 32767:
                                v2 = "a"
                                v73 := GenS3{F0: false}
                                var v74 float32 = (3.14159 + ((float32((-2.25 / -2.25)) / ((-2.25 * 0.1) + -2.25)) / 1e10))
                                _ = v73
                                _ = v74
                            case 32768:
                                if true {
                                    v75 := GenS5{F0: -32768, F1: -763}
                                    var v76 int16 = -637
                                    _ = v75
                                    _ = v76
                                }
                            default:
                                v2 = "a"
                                var v77 uint8 = 237
                                _ = v77
                            }
                            _ = v71
                        }
                        v40 = append(v40, (fi70 ^ 2147483647))
                        _ = fi70
                    }
                    _ = fi69
                }
            default:
                var v78 uint = 1
                _ = v78
            }
            if (42 != -1) {
                if true {
                    var v79 string = v2
                    delete(v3, -987)
                    _ = v79
                }
            }
            for fi80 := 0; fi80 < 15; fi80++ {
                for fi81 := 0; fi81 < 17; fi81++ {
                    if true {
                        if true {
                            var v82 int = (fi81 | (fi80 / (((fi81 | fi81)) | 1)))
                            _ = v82
                        }
                        if (1 < 127) {
                            v40 = append(v40, (fi81 * 128))
                            v57 = v57
                        }
                        if (v4 == (^(v4))) {
                            v83 := v3[fi80]
                            _ = v83
                        }
                    } else {
                        var v84 uint8 = 0
                        v85 := GenS6{F0: -1}
                        v86 := GenS2{F0: 2147483647, F1: 121, F2: -843}
                        v57 -= (v57 % ((v57) | 1))
                        _ = v84
                        _ = v85
                        _ = v86
                    }
                    var v87 string = v2
                    var v88 int = (fi80 - (fi80 + (fi80 & fi80)))
                    var v89 uint = 4294967295
                    _ = fi81
                    _ = v87
                    _ = v88
                    _ = v89
                }
                var v90 float64 = (-2.25 - 0.0)
                v90 = (v90 / ((1.5 - 1e10) / ((v90 + v90) - 1e10)))
                _ = fi80
                _ = v90
            }
        }
        switch v2 {
        case "\n\t":
            var v91 int32 = 42
            _ = v91
        }
        v4 -= 0
    } else {
        var v92 uint64 = 4294967295
        _ = v92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v57
}
