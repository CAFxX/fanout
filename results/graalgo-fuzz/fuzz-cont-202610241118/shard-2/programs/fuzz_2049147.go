package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 bool; F1 bool; F2 bool }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint16; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 int; F2 int16 }
type GenS8 struct { F0 int8; F1 int16; F2 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 string; F1 uint; F2 uint64 }

func main() {
    var v1 int16 = 163
    var v2 int16 = v1
    v3 := make(map[uint32]string)
    var v4 uint = 0
    var v5 int8 = -1
    switch "" {
    case "a":
        fmt.Println(v1)
    case "foo bar":
        if true {
            v6 := []int{32768, -9223372036854775808}
            switch 42 {
            case 128:
                v4 = 65535
            }
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: -383}
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    case "hello":
        if false {
            v9 := []int{-9223372036854775808, 32768, -9223372036854775808, 965}
            _ = v9
        }
    }
    v10 := GenS1{F0: -438}
    for fi11 := 0; fi11 < 8; fi11++ {
        v1 += (v2 - (v2 % (((^(v1))) | 1)))
        var v12 float64 = 0.0
        _ = fi11
        _ = v12
    }
    v13 := []int{32767, -869, 32767}
    fmt.Println(v5, v1, v5)
    v14 := v3[65535]
    v2 = -32768
    var v15 int16 = 42
    if true {
        var v16 int = 32768
        _ = v16
    }
    v17 := GenS2{F0: (v14 + "foo bar"), F1: 141}
    v4 = 65535
    fmt.Println(v1, v5, v2)
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int16 = (v15 / (((v2 ^ v15)) | 1))
        if true {
            for fi20 := 0; fi20 < 10; fi20++ {
                if true {
                    var v21 uint64 = 2
                    v22 := GenS3{F0: (128 == 255), F1: (!((42 > 0) || (true && false))), F2: false}
                    _ = v21
                    _ = v22
                } else {
                    if false {
                        v19 -= 604
                        v5 = v5
                    }
                    var v23 uint32 = 698
                    _ = v23
                }
                if (!(794 <= 317)) {
                    delete(v3, 2147483647)
                    _ = len(v13)
                    v24 := GenS4{F0: (fi18 << ((v4) & 63)), F1: 248}
                    _ = v24
                } else {
                    var v25 uint64 = 18446744073709551615
                    v15 = v19
                    _ = v25
                }
                _ = fi20
            }
            var v26 float32 = ((0.0 / (float32(-2.25) / 3.14159)) - (0.1 / (float32(((float64(((((((-2.25 - 0.1) + 1.5) - 0.0) - 0.1) + 3.14159) / 1.0)) + float64(0.1)) / 1.5)) / 3.14159)))
            for fi27 := 0; fi27 < 11; fi27++ {
                var v28 float64 = 0.1
                var v29 uint = v4
                if ((!((v2 / (((v1 % ((-1) | 1))) | 1)) > (v2 * int16(v29)))) || false) {
                    switch 128 {
                    case 127:
                        v19 -= (v1 / ((v2) | 1))
                    }
                    fmt.Println(v28, v26)
                    v30 := GenS4{F0: fi18, F1: 127}
                    v31 := v3[4294967295]
                    _ = v30
                    _ = v31
                }
                _ = len(v13)
                switch -627 {
                case 127:
                    fmt.Println(v4)
                case 32767:
                    var v32 bool = (fi18 != fi27)
                    _ = v32
                case 128:
                    var v33 int = fi27
                    _ = v33
                default:
                    switch (fi18 << ((v4) & 63)) {
                    case 217:
                        v19 += -1
                    case -124:
                        if (v29 >= (v4 & (v4 >> ((65535) & 63)))) {
                            if (float32(((v28 / (0.0 + 1e10)) - (v28 - (v28 / (1.5 / 1.0))))) != v26) {
                                switch (("\n\t" + v14) + (v14 + ("\n\t" + "\n\t"))) {
                                case "\n\t":
                                    v29 = 227
                                case "hello":
                                    v13 = append(v13, fi27)
                                }
                                v15 = -1
                                if (0 >= 411) {
                                    if false {
                                        switch fi27 {
                                        case 818:
                                            v34 := GenS5{F0: 0, F1: 2, F2: 1}
                                            _ = v34
                                        case -563:
                                            v2 -= v19
                                        case -1:
                                            v35 := GenS1{F0: -1}
                                            _ = v35
                                        }
                                        if false {
                                            v26 -= 3.14159
                                            v4 = 0
                                            v15 -= v19
                                        }
                                        v3[2147483648] = v14
                                    }
                                    v36 := GenS6{F0: -56}
                                    if ((!true) || false) {
                                        v26 -= v26
                                        var v37 uint32 = 2
                                        v1 += v15
                                        _ = v37
                                    } else {
                                        if ((42 > 0) && false) {
                                            if (1 < 1) {
                                                var v38 int16 = 1000
                                                v26 -= float32(0.0)
                                                var v39 bool = false
                                                _ = v38
                                                _ = v39
                                            } else {
                                                var v40 uint16 = 1
                                                fmt.Println(v1, v5, v28)
                                                switch fi27 {
                                                case 485:
                                                    v3[1] = v14
                                                }
                                                _ = v40
                                            }
                                            var v41 int32 = int32(v15)
                                            var v42 string = ""
                                            if false {
                                                var v43 uint = 0
                                                v2 = -1
                                                _ = v43
                                            }
                                            _ = v41
                                            _ = v42
                                        } else {
                                            var v44 uint16 = 32767
                                            _ = len(v13)
                                            var v45 uint16 = 119
                                            v1 += (-(v19))
                                            _ = v44
                                            _ = v45
                                        }
                                        switch (fi18 / ((fi18) | 1)) {
                                        case 32768:
                                            fmt.Println(fi18, v14)
                                        case -2147483648:
                                            v5 = v5
                                        }
                                        v4 += (v4 + v4)
                                    }
                                    v15 = 34
                                    _ = v36
                                }
                            }
                        } else {
                            v46 := GenS7{F0: 4294967295, F1: (fi27 ^ (fi27 % ((fi27) | 1))), F2: (^(v19))}
                            if true {
                                v5 -= v5
                                fmt.Println(v19, v26, v5)
                                v47 := v3[0]
                                v4 += v29
                                _ = v47
                            } else {
                                fmt.Println(v4)
                                v48 := GenS8{F0: (v5 - v5), F1: v15, F2: -1}
                                if ((v4 >= (v4 % ((0) | 1))) || (!(2 != 255))) {
                                    if ((v26 >= v26) || (v4 != (v29 >> ((v29) & 63)))) {
                                        var v49 int64 = 0
                                        _ = v49
                                    }
                                    switch fi18 {
                                    case -9223372036854775808:
                                        v50 := GenS9{F0: v14}
                                        _ = v50
                                    case 2147483647:
                                        v29 -= v29
                                        var v51 int = fi18
                                        _ = v51
                                    }
                                    v29 = (^(v29))
                                } else {
                                    v28 += (v28 + v28)
                                    switch "foo bar" {
                                    case "a":
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                    case "\n\t":
                                        var v52 int8 = v5
                                        _ = v52
                                    case "foo bar":
                                        v53 := GenS2{F0: v14, F1: 307}
                                        _ = v53
                                    }
                                }
                                if (false && false) {
                                    if true {
                                        fmt.Println(v15, v14)
                                    } else {
                                        v54 := GenS10{F0: true}
                                        v19 = (v15 + v19)
                                        var v55 int64 = -1
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 uint8 = 1
                                    v15 = v19
                                    var v57 int64 = -1
                                    _ = v56
                                    _ = v57
                                } else {
                                    var v58 int = fi18
                                    v4 = v4
                                    if len(v13) > 0 {
                                        v13[0] = fi18
                                    }
                                    _ = v58
                                }
                                _ = v48
                            }
                            _ = v46
                        }
                    default:
                        var v59 uint16 = 42
                        _ = v59
                    }
                }
                fmt.Println(v4)
                v5 -= 42
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v26
        } else {
            for fi60 := 0; fi60 < 6; fi60++ {
                v13 = append(v13, (fi60 ^ fi18))
                _ = fi60
            }
            for fi61 := 0; fi61 < 11; fi61++ {
                if false {
                    var v62 uint8 = 0
                    v63 := GenS5{F0: 0, F1: 65535, F2: 42}
                    _ = v62
                    _ = v63
                }
                _ = fi61
            }
        }
        if false {
            if false {
                v4 = (^(v4))
                v64 := GenS1{F0: 0}
                _ = v64
            }
        } else {
            for fi65 := 0; fi65 < 8; fi65++ {
                fmt.Println(fi65)
                if (!(-98 < 2147483647)) {
                    switch v14 {
                    case "hello":
                        v3[513] = ((v14 + v14) + "a")
                    }
                    v15 -= (v1 & -835)
                    v19 += 1000
                }
                switch ("a" + ("foo bar" + v14)) {
                case "foo bar":
                    var v66 string = "\n\t"
                    _ = v66
                default:
                    v19 = v1
                }
                _ = fi65
            }
            switch v14 {
            case "hello":
                fmt.Println(v14)
            default:
                for fi67 := 0; fi67 < 17; fi67++ {
                    if false {
                        switch fi18 {
                        case 42:
                            fmt.Println(v14, v5, v5)
                        }
                    } else {
                        if (!true) {
                            switch v14 {
                            case "a":
                                fmt.Println(v15, v14, fi18)
                            }
                            if ((v14 + ("\n\t" + v14)) == v14) {
                                v1 = (v19 % (((-(v2))) | 1))
                                fmt.Println(v4, v1, v19)
                                var v68 int32 = 0
                                _ = v68
                            } else {
                                delete(v3, 2147483648)
                                v15 -= v19
                                v69 := GenS11{F0: fi18, F1: 329}
                                _ = v69
                            }
                            v1 = (v1 | v2)
                            fmt.Println(v14, v14, v1)
                            switch v14 {
                            case "":
                                if len(v13) > 0 {
                                    v13[0] = 32768
                                }
                                v5 -= 0
                                var v70 uint64 = 2
                                fmt.Println(v19, v4, v70)
                                v5 += 100
                                v71 := v3[2147483648]
                                _ = v70
                                _ = v71
                            case "hello":
                                v72 := GenS1{F0: 2147483647}
                                if false {
                                    fmt.Println(v19, fi18, v5)
                                    fmt.Println(v19, v5)
                                } else {
                                    _ = len(v3)
                                    var v73 float32 = -2.25
                                    _ = v73
                                }
                                if (-2147483648 != 714) {
                                    v74 := GenS11{F0: fi67, F1: 42}
                                    v75 := GenS5{F0: -2147483648, F1: uint16(v4), F2: -357}
                                    var v76 int8 = v5
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                }
                                var v77 int32 = -2147483648
                                v78 := GenS12{F0: "a", F1: 255, F2: 2}
                                if (((fi18 << (((v4 ^ v4)) & 63)) >= (fi18 ^ int(v15))) || false) {
                                    v2 = (v19 ^ (v19 & (v15 + (v1 - (v2 / ((-84) | 1))))))
                                } else {
                                    v79 := v3[1]
                                    _ = v79
                                }
                                _ = len(v13)
                                _ = v72
                                _ = v77
                                _ = v78
                            }
                            switch "" {
                            case "a":
                                fmt.Println(v19)
                            case "foo bar":
                                var v80 uint32 = 2147483647
                                _ = v80
                            }
                        } else {
                            fmt.Println(fi18, v2, v1)
                            var v81 uint16 = 1
                            if (1 != (v81 << (((v4 | (v4 * 144))) & 63))) {
                                var v82 int = fi67
                                _ = v82
                            }
                            _ = v81
                        }
                        switch v14 {
                        case "hello":
                            _ = len(v13)
                        case "\n\t":
                            v2 -= (v1 ^ (v1 % (((^(v2))) | 1)))
                        case "":
                            v4 -= v4
                        default:
                            var v83 bool = true
                            _ = v83
                        }
                        v84 := GenS9{F0: v14}
                        switch (((("hello" + v14) + v14) + v14) + v14) {
                        case "a":
                            _ = len(v3)
                        case "foo bar":
                            var v85 uint16 = 217
                            _ = v85
                        }
                        if (38 > 42) {
                            v13 = append(v13, fi67)
                            var v86 int = (fi67 << (((v4 + (v4 | 0))) & 63))
                            var v87 float64 = ((3.14159 + 1.5) / float64(0.1))
                            var v88 uint = v4
                            v1 += v15
                            _ = v86
                            _ = v87
                            _ = v88
                        }
                        _ = v84
                    }
                    if true {
                        v1 = v1
                    } else {
                        v89 := GenS6{F0: 1}
                        _ = v89
                    }
                    _ = fi67
                }
            }
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
