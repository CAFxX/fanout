package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 uint32; F2 string }
type GenS6 struct { F0 uint; F1 int8; F2 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 int64; F1 bool; F2 int64 }
type GenS11 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 int }
func genFn0() uint8 {
    v4 := GenS1{F0: 1}
    v5 := GenS2{F0: -1, F1: 395, F2: -1}
    var v6 string = ""
    fmt.Println(v6)
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(fi7 + 724), 42}
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 bool = (false || (65535 != 42))
    var v2 uint8 = 86
    var v3 int = 127
    v1 = (42 == 1)
    v9 := GenS2{F0: 0, F1: 439, F2: -9223372036854775808}
    if (v1 || (0.1 > (0.0 - 1e10))) {
        v10 := GenS1{F0: -108}
        _ = v10
    }
    var v11 float32 = (0.1 * 0.0)
    var v12 uint32 = 0
    v12 += v12
    v13 := []int{}
    if (1 <= 1) {
        if (!(42 >= 255)) {
            for fi14 := 0; fi14 < 20; fi14++ {
                if len(v13) > 0 {
                    v13[0] = v3
                }
                if (1000 == 32767) {
                    _ = len(v13)
                    v15 := GenS1{F0: 100}
                    _ = v15
                }
                _ = fi14
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                v2 = (v2 >> ((1) & 63))
                fmt.Println(fi16, v1, v1)
                var v17 int32 = -613
                _ = fi16
                _ = v17
            }
            switch (("" + (("" + "a") + "\n\t")) + "") {
            case "hello":
                var v18 uint16 = 161
                _ = v18
            }
        }
        var v19 int64 = 9223372036854775807
        var v20 float64 = -2.25
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 float64 = 1e10
            switch (v3 * (fi21 + (^(v3)))) {
            case 4294967295:
                var v23 uint32 = 724
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        var v24 int16 = 42
        for fi25 := 0; fi25 < 9; fi25++ {
            v2 += v2
            if v1 {
                v26 := GenS2{F0: int16(v3), F1: 255, F2: 9223372036854775807}
                _ = v26
            } else {
                switch "foo bar" {
                case "hello":
                    v11 += (((v11 / (1e10 * 0.1)) - float32((1e10 * (-2.25 * ((0.1 / 1.5) + 0.0))))) * 1e10)
                case "foo bar":
                    v24 += (v24 & v24)
                default:
                    _ = len(v13)
                }
                v27 := GenS3{F0: v12}
                _ = v27
            }
            v2 = (v2 ^ v2)
            _ = fi25
        }
        _ = v24
    }
    v1 = (((3.14159 + (0.0 / float64(0.1))) + 1.5) == 0.1)
    if false {
        v28 := GenS2{F0: 1, F1: 255, F2: 9223372036854775807}
        _ = v28
    } else {
        fmt.Println(v1, v11, v2)
        v29 := GenS3{F0: 853}
        switch (v3 * (v3 - v3)) {
        case -2147483648:
            v3 = (v3 + v3)
        case 32768:
            var v30 uint32 = 398
            _ = v30
        case 9223372036854775807:
            if true {
                var v31 int = v3
                switch 2147483647 {
                case 28:
                    if (-861 == 1000) {
                        fmt.Println(v11, v12, v3)
                    }
                case -1:
                    v11 += (0.1 / (((v11 * (float32(float64((0.1 + (0.1 + 0.1)))) - v11)) + v11) * v11))
                case 1:
                    if v1 {
                        for fi32 := 0; fi32 < 11; fi32++ {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            if ((0 < 42) || v1) {
                                v2 = (v2 / ((v2) | 1))
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    var v34 bool = v1
                                    if v34 {
                                        v35 := GenS4{F0: 44}
                                        v36 := GenS5{F0: -1, F1: (^(v12)), F2: "\n\t"}
                                        v37 := GenS6{F0: 65535, F1: -118, F2: 1}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            switch (v31 ^ int(v12)) {
                            case 32767:
                                fmt.Println(v12)
                            default:
                                var v38 bool = ((v31 + v31) >= (v3 / ((fi32) | 1)))
                                _ = v38
                            }
                            _ = fi32
                        }
                        v39 := GenS6{F0: 18446744073709551615, F1: -128, F2: 726}
                        _ = v39
                    } else {
                        switch "a" {
                        case "foo bar":
                            if v1 {
                                fmt.Println(v1, v31, v12)
                            }
                        default:
                            var v40 uint64 = 472
                            _ = v40
                        }
                        v13 = append(v13, v31)
                        for fi41 := 0; fi41 < 13; fi41++ {
                            if true {
                                var v42 uint16 = 255
                                _ = v42
                            } else {
                                for fi43 := 0; fi43 < 3; fi43++ {
                                    fmt.Println(v3, v12)
                                    fmt.Println(v31, v1)
                                    v44 := GenS4{F0: -2147483648}
                                    if len(v13) > 0 {
                                        v13[0] = 1
                                    }
                                    _ = fi43
                                    _ = v44
                                }
                                if (uint32(v2) <= v12) {
                                    _ = len(v13)
                                    var v45 float32 = ((float32((0.1 * (-2.25 - 0.0))) - v11) - v11)
                                    var v46 float32 = 0.0
                                    _ = v45
                                    _ = v46
                                } else {
                                    fmt.Println(v1, v3, v1)
                                    v47 := GenS7{F0: 35}
                                    _ = v47
                                }
                                if (!v1) {
                                    if (v2 != (v2 * 166)) {
                                        v12 += uint32(v2)
                                        v31 += (fi41 % ((fi41) | 1))
                                    }
                                } else {
                                    v48 := GenS8{F0: 18446744073709551615, F1: "\n\t"}
                                    for fi49 := 0; fi49 < 5; fi49++ {
                                        if v1 {
                                            v50 := GenS5{F0: -128, F1: 412, F2: "a"}
                                            var v51 bool = ("" <= "a")
                                            _ = v50
                                            _ = v51
                                        }
                                        var v52 uint = 186
                                        if v1 {
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(v2, fi49)
                                            default:
                                                if (v2 <= v2) {
                                                    _ = len(v13)
                                                    if v1 {
                                                        var v53 float64 = (1.5 / ((1e10 + (1e10 + float64(float32(((0.0 - 0.1) + 0.1))))) + 0.0))
                                                        v54 := GenS9{F0: 700, F1: 354}
                                                        v12 = (v12 + 207)
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        v52 -= 2
                                                        v13 = append(v13, v31)
                                                        v3 = fi49
                                                    }
                                                    v1 = false
                                                }
                                                switch 32767 {
                                                case 42:
                                                    var v55 uint = v52
                                                    _ = v55
                                                case 1:
                                                    switch "" {
                                                    case "hello":
                                                        v56 := GenS5{F0: -1, F1: uint32(v2), F2: "hello"}
                                                        _ = v56
                                                    case "foo bar":
                                                        if (true || v1) {
                                                            var v57 uint8 = (v2 ^ v2)
                                                            v58 := GenS9{F0: -714, F1: -1}
                                                            v52 -= (v52 & 65535)
                                                            _ = v57
                                                            _ = v58
                                                        } else {
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                            v12 = v12
                                                            if (9223372036854775807 > 0) {
                                                                v59 := GenS4{F0: 2147483647}
                                                                if len(v13) > 0 {
                                                                    v13[0] = (v31 + fi41)
                                                                }
                                                                _ = v59
                                                            }
                                                            var v60 uint8 = v2
                                                            _ = v60
                                                        }
                                                    case "\n\t":
                                                        v12 = (v12 | (v12 + (v12 >> ((v52) & 63))))
                                                    default:
                                                        v2 -= (v2 - (v2 % (((^(v2))) | 1)))
                                                    }
                                                default:
                                                    fmt.Println(fi41, v2)
                                                }
                                                fmt.Println(v11)
                                                if true {
                                                    if ((fi49 % ((fi49) | 1)) == 4294967295) {
                                                        v61 := GenS10{F0: 0, F1: (2147483647 > 715), F2: 0}
                                                        _ = len(v13)
                                                        var v62 uint = (v52 & v52)
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    var v63 bool = (v2 != (^(v2)))
                                                    _ = v63
                                                }
                                            }
                                            var v64 int = -1
                                            _ = v64
                                        }
                                        fmt.Println(v2)
                                        var v65 int8 = 13
                                        switch "" {
                                        case "foo bar":
                                            if ((!v1) || (v1 && v1)) {
                                                if (42 < 0) {
                                                    if false {
                                                        if ((^(v12)) == (^(v12))) {
                                                            if len(v13) > 0 {
                                                                v13[0] = fi41
                                                            }
                                                            if v1 {
                                                                v52 += (v52 ^ v52)
                                                            } else {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v13 = append(v13, (v3 >> (((v52 - 1)) & 63)))
                                                                case "a":
                                                                    if (!v1) {
                                                                        fmt.Println(v12)
                                                                        var v66 int32 = int32(v2)
                                                                        _ = v66
                                                                    }
                                                                }
                                                                if len(v13) > 0 {
                                                                    v13[0] = (v31 % ((int(v65)) | 1))
                                                                }
                                                                fmt.Println(v2, v3)
                                                                v67 := GenS11{F0: 150, F1: v2, F2: 1}
                                                                _ = v67
                                                            }
                                                            v2 = v2
                                                        }
                                                    } else {
                                                        fmt.Println(fi41)
                                                        if (42 < uint16(v2)) {
                                                            _ = len(v13)
                                                            v68 := GenS12{F0: (!false), F1: (fi49 + -203)}
                                                            if v1 {
                                                                fmt.Println(v12)
                                                                fmt.Println(v12, v1, fi41)
                                                                switch v31 {
                                                                case -902:
                                                                    if v1 {
                                                                        var v69 uint64 = 148
                                                                        _ = v69
                                                                    }
                                                                default:
                                                                    switch v3 {
                                                                    case 355:
                                                                        var v70 int = fi41
                                                                        _ = v70
                                                                    case 128:
                                                                        _ = len(v13)
                                                                    default:
                                                                        switch fi41 {
                                                                        case -9223372036854775808:
                                                                            var v71 int = v31
                                                                            _ = v71
                                                                        case 9223372036854775807:
                                                                            v52 += (v52 ^ (v52 >> ((154) & 63)))
                                                                        case 2147483647:
                                                                            var v72 bool = v1
                                                                            _ = v72
                                                                        default:
                                                                            if len(v13) > 0 {
                                                                                v13[0] = fi49
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v68
                                                        }
                                                        fmt.Println(v11)
                                                    }
                                                }
                                                v11 = 0.0
                                            } else {
                                                if (v12 <= 906) {
                                                    fmt.Println(fi49)
                                                    switch ("" + ("" + "")) {
                                                    case "\n\t":
                                                        var v73 int16 = -751
                                                        _ = v73
                                                    case "a":
                                                        _ = len(v13)
                                                    default:
                                                        v13 = append(v13, fi49)
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(v31, v2)
                                        default:
                                            var v74 uint8 = (v2 << ((42) & 63))
                                            _ = v74
                                        }
                                        _ = fi49
                                        _ = v52
                                        _ = v65
                                    }
                                    v75 := GenS13{F0: v3}
                                    _ = v48
                                    _ = v75
                                }
                            }
                            var v76 uint = 65535
                            v31 = v31
                            _ = fi41
                            _ = v76
                        }
                    }
                default:
                    var v77 float32 = v11
                    _ = v77
                }
                _ = v31
            } else {
                var v78 uint = 776
                var v79 int64 = -9223372036854775808
                for fi80 := 0; fi80 < 2; fi80++ {
                    var v81 uint = 0
                    _ = fi80
                    _ = v81
                }
                _ = v78
                _ = v79
            }
        }
        _ = v29
    }
    if len(v13) > 0 {
        v13[0] = (v3 << ((42) & 63))
    }
    if false {
        for fi82 := 0; fi82 < 14; fi82++ {
            var v83 int8 = 100
            _ = fi82
            _ = v83
        }
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
