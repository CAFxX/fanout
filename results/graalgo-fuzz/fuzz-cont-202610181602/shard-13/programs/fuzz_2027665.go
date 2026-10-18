package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 string }
type GenS3 struct { F0 int8; F1 uint8; F2 uint }
type GenS4 struct { F0 uint64; F1 bool; F2 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 bool; F2 uint64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS15 struct { F0 uint16; F1 uint8 }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 string; F1 uint32 }
type GenS18 struct { F0 bool; F1 int16; F2 int8 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 uint32; F1 string }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 int }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 int; F1 int; F2 int }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 string }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 uint8 }
type GenS29 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int = -285
    var v2 float32 = 1e10
    v3 := make(map[uint32]int)
    var v4 uint64 = 18446744073709551615
    var v5 int16 = -32768
    v6 := []string{("a" + "\n\t")}
    switch "hello" {
    case "a":
        var v7 float64 = 0.0
        _ = v7
    case "foo bar":
        var v8 uint64 = 4294967295
        _ = v8
    default:
        if false {
            switch v1 {
            case -120:
                if false {
                    if false {
                        var v9 float64 = (1.5 * (3.14159 / float64(v2)))
                        for fi10 := 0; fi10 < 20; fi10++ {
                            for fi11 := 0; fi11 < 10; fi11++ {
                                v2 = (v2 * 0.1)
                                _ = fi11
                            }
                            v6 = append(v6, "")
                            _ = fi10
                        }
                        var v12 string = ("a" + "foo bar")
                        _ = v12
                        _ = v9
                    } else {
                        var v13 string = ("" + "")
                        if (!false) {
                            fmt.Println(v13, v2, v4)
                        } else {
                            for fi14 := 0; fi14 < 17; fi14++ {
                                v6 = append(v6, v13)
                                if true {
                                    if true {
                                        if ((true && (255 <= 127)) && true) {
                                            var v15 uint8 = 255
                                            if ((v2 / 1.0) < 0.1) {
                                                for fi16 := 0; fi16 < 13; fi16++ {
                                                    if false {
                                                        switch fi16 {
                                                        case 324:
                                                            switch v13 {
                                                            case "hello":
                                                                v13 = (("" + (v13 + v13)) + v13)
                                                            case "":
                                                                if (fi16 > (v1 >> ((506) & 63))) {
                                                                    var v17 string = ""
                                                                    v17 = (((("hello" + ("" + "")) + "hello") + "\n\t") + v17)
                                                                    _ = v17
                                                                }
                                                            default:
                                                                var v18 int32 = 0
                                                                _ = v18
                                                            }
                                                        case -9223372036854775808:
                                                            if false {
                                                                if (65535 < 1) {
                                                                    fmt.Println(fi14, v1, v15)
                                                                    var v19 int = fi16
                                                                    switch (fi16 - (fi14 - (v1 ^ fi16))) {
                                                                    case -9223372036854775808:
                                                                        switch v19 {
                                                                        case -1:
                                                                            if false {
                                                                                if false {
                                                                                    v2 = float32(1.5)
                                                                                    if true {
                                                                                        if true {
                                                                                            v2 = (v2 / -2.25)
                                                                                            if false {
                                                                                                if (!(!true)) {
                                                                                                    if (!false) {
                                                                                                        var v20 uint16 = 42
                                                                                                        _ = v20
                                                                                                    } else {
                                                                                                        v19 = v1
                                                                                                        _ = len(v6)
                                                                                                        v13 = v13
                                                                                                        if ((-1 < 100) || true) {
                                                                                                            fmt.Println(fi14, v15)
                                                                                                            fmt.Println(v19, v2, v15)
                                                                                                            var v21 float64 = 1e10
                                                                                                            _ = len(v3)
                                                                                                            var v22 int64 = 1
                                                                                                            _ = v21
                                                                                                            _ = v22
                                                                                                        }
                                                                                                        fmt.Println(v2, v2)
                                                                                                        var v23 uint64 = (v4 << ((255) & 63))
                                                                                                        if false {
                                                                                                            v24 := GenS1{F0: 640, F1: 0, F2: v15}
                                                                                                            var v25 uint16 = 2
                                                                                                            v26 := GenS2{F0: -9223372036854775808, F1: v25, F2: "a"}
                                                                                                            var v27 bool = (((v2 / v2) + -2.25) == ((1e10 * v2) + 1e10))
                                                                                                            switch v1 {
                                                                                                            case 0:
                                                                                                                v28 := v3[2147483648]
                                                                                                                _ = v28
                                                                                                            default:
                                                                                                                v19 -= (fi16 - fi14)
                                                                                                            }
                                                                                                            _ = v24
                                                                                                            _ = v25
                                                                                                            _ = v26
                                                                                                            _ = v27
                                                                                                        } else {
                                                                                                            fmt.Println(v5)
                                                                                                            fmt.Println(v19, v1)
                                                                                                            fmt.Println(v13, fi14)
                                                                                                        }
                                                                                                        v1 -= (v19 - 1)
                                                                                                        v13 = "hello"
                                                                                                        _ = v23
                                                                                                    }
                                                                                                    var v29 int = v1
                                                                                                    fmt.Println(fi14, v4, v13)
                                                                                                    fmt.Println(v15)
                                                                                                    _ = v29
                                                                                                } else {
                                                                                                    v30 := GenS3{F0: 42, F1: (^(v15)), F2: 958}
                                                                                                    _ = v30
                                                                                                }
                                                                                                v2 = ((((float32(3.14159) / (v2 + (0.0 * v2))) * v2) - (v2 + float32((1e10 * (1.5 / ((float64((v2 / v2)) * 3.14159) * float64(v2))))))) + 0.1)
                                                                                            }
                                                                                        } else {
                                                                                            v4 = uint64(v15)
                                                                                            if len(v6) >= 2 {
                                                                                                v6 = v6[1:len(v6)]
                                                                                            }
                                                                                            v1 += (fi16 & (v19 << ((1) & 63)))
                                                                                        }
                                                                                        var v31 int16 = 42
                                                                                        _ = v31
                                                                                    }
                                                                                    v2 = 3.14159
                                                                                } else {
                                                                                    v6 = append(v6, (v13 + "a"))
                                                                                }
                                                                                v32 := GenS4{F0: v4, F1: false, F2: true}
                                                                                var v33 string = v13
                                                                                _ = v32
                                                                                _ = v33
                                                                            } else {
                                                                                if false {
                                                                                    fmt.Println(v1, v13)
                                                                                    v3[507] = (v1 & v19)
                                                                                } else {
                                                                                    var v34 bool = (false || (0 <= -763))
                                                                                    if v34 {
                                                                                        v35 := GenS1{F0: -265, F1: -9223372036854775808, F2: (^(v15))}
                                                                                        _ = v35
                                                                                    }
                                                                                    _ = v34
                                                                                }
                                                                            }
                                                                        case 128:
                                                                            v36 := GenS3{F0: -1, F1: v15, F2: 889}
                                                                            if true {
                                                                                if false {
                                                                                    v37 := GenS2{F0: 175, F1: 255, F2: (v13 + (v13 + v13))}
                                                                                    v38 := GenS3{F0: 0, F1: 130, F2: 264}
                                                                                    _ = v37
                                                                                    _ = v38
                                                                                }
                                                                            }
                                                                            _ = v36
                                                                        case -2147483648:
                                                                            v2 = float32((1e10 + float64(v2)))
                                                                        }
                                                                        if false {
                                                                            v4 = (v4 ^ 0)
                                                                            var v39 uint16 = 578
                                                                            _ = v39
                                                                        } else {
                                                                            fmt.Println(v13, v13, fi14)
                                                                            switch v13 {
                                                                            case "\n\t":
                                                                                switch -1 {
                                                                                case 2147483647:
                                                                                    fmt.Println(v2, v13)
                                                                                case 32768:
                                                                                    fmt.Println(v4, v19, v5)
                                                                                }
                                                                            case "":
                                                                                delete(v3, 4294967295)
                                                                                v40 := GenS2{F0: -1, F1: 65535, F2: ((v13 + "hello") + "a")}
                                                                                if true {
                                                                                    var v41 float32 = (3.14159 / 1.0)
                                                                                    switch v13 {
                                                                                    case "\n\t":
                                                                                        if true {
                                                                                            v2 = ((float32(0.0) * (v2 - float32((3.14159 * -2.25)))) - ((float32(1e10) / (v41 / v41)) * 0.1))
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                var v42 int64 = 11
                                                                                                _ = v42
                                                                                            default:
                                                                                                v43 := GenS5{F0: true}
                                                                                                _ = v43
                                                                                            }
                                                                                        } else {
                                                                                            v44 := GenS2{F0: -126, F1: 32767, F2: (v13 + v13)}
                                                                                            v3[4294967295] = fi16
                                                                                            var v45 uint16 = 445
                                                                                            _ = v44
                                                                                            _ = v45
                                                                                        }
                                                                                        var v46 int = fi14
                                                                                        v13 = v13
                                                                                        _ = v46
                                                                                    }
                                                                                    if false {
                                                                                        v5 += (v5 | -969)
                                                                                        switch v1 {
                                                                                        case 127:
                                                                                            if (false || false) {
                                                                                                fmt.Println(fi14, v5)
                                                                                                if ((v5 + v5) > (v5 * (v5 / (((v5 - v5)) | 1)))) {
                                                                                                    v47 := GenS6{F0: 657, F1: 32767, F2: 255}
                                                                                                    var v48 int = fi14
                                                                                                    _ = v47
                                                                                                    _ = v48
                                                                                                } else {
                                                                                                    if false {
                                                                                                        if (!true) {
                                                                                                            v4 = v4
                                                                                                        }
                                                                                                        if true {
                                                                                                            var v49 uint16 = 65535
                                                                                                            _ = v49
                                                                                                        }
                                                                                                        v50 := GenS7{F0: v1}
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    fmt.Println(v41, v1, v19)
                                                                                                    var v51 int32 = 1
                                                                                                    _ = v51
                                                                                                }
                                                                                                if true {
                                                                                                    switch (fi14 | 32768) {
                                                                                                    case -240:
                                                                                                        if false {
                                                                                                            var v52 int = -9223372036854775808
                                                                                                            v13 = (v13 + (((v13 + "") + v13) + "\n\t"))
                                                                                                            _ = v52
                                                                                                        } else {
                                                                                                            var v53 int64 = -9223372036854775808
                                                                                                            switch (fi14 - (fi16 + (v1 / ((v1) | 1)))) {
                                                                                                            case -1:
                                                                                                                var v54 uint32 = 42
                                                                                                                _ = v54
                                                                                                            }
                                                                                                            _ = v53
                                                                                                        }
                                                                                                    case 0:
                                                                                                        v55 := v3[517]
                                                                                                        _ = v55
                                                                                                    }
                                                                                                } else {
                                                                                                    v56 := GenS8{F0: v19, F1: 42}
                                                                                                    _ = len(v6)
                                                                                                    _ = v56
                                                                                                }
                                                                                                if false {
                                                                                                    var v57 int8 = 1
                                                                                                    _ = v57
                                                                                                }
                                                                                                v58 := GenS6{F0: 0, F1: 1, F2: 65535}
                                                                                                _ = v58
                                                                                            }
                                                                                        default:
                                                                                            if true {
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = ((v13 + v13) + v13)
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    v1 = v1
                                                                                                    v3[4294967295] = (v1 % (((-(fi14))) | 1))
                                                                                                    var v59 float64 = 0.1
                                                                                                    _ = v59
                                                                                                } else {
                                                                                                    v19 = 4294967295
                                                                                                    _ = len(v6)
                                                                                                    v5 = 1
                                                                                                    v15 = (v15 << ((4294967295) & 63))
                                                                                                }
                                                                                                fmt.Println(v19, v15)
                                                                                                var v60 uint32 = 0
                                                                                                v61 := GenS4{F0: (v4 % ((18446744073709551615) | 1)), F1: false, F2: true}
                                                                                                _ = v60
                                                                                                _ = v61
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            if (1 < v5) {
                                                                                                if (false || false) {
                                                                                                    if false {
                                                                                                        v62 := GenS9{F0: 570, F1: 65535, F2: 1}
                                                                                                        _ = v62
                                                                                                    } else {
                                                                                                        if len(v6) >= 2 {
                                                                                                            v6 = v6[1:len(v6)]
                                                                                                        }
                                                                                                    }
                                                                                                    var v63 bool = true
                                                                                                    fmt.Println(v19, v13, v63)
                                                                                                    v4 -= 4294967295
                                                                                                    if len(v6) > 0 {
                                                                                                        v6[0] = (v13 + ("" + "foo bar"))
                                                                                                    }
                                                                                                    _ = v63
                                                                                                }
                                                                                                v19 = fi16
                                                                                            }
                                                                                            if true {
                                                                                                if true {
                                                                                                    v15 = (v15 - (^(v15)))
                                                                                                }
                                                                                                if false {
                                                                                                    v3[35] = fi14
                                                                                                    var v64 int64 = -1
                                                                                                    _ = v64
                                                                                                }
                                                                                            }
                                                                                            var v65 uint8 = v15
                                                                                            switch v1 {
                                                                                            case 1:
                                                                                                v65 += 2
                                                                                            case 618:
                                                                                                v66 := GenS2{F0: 648, F1: 255, F2: v13}
                                                                                                _ = v66
                                                                                            default:
                                                                                                v6 = append(v6, v13)
                                                                                            }
                                                                                            fmt.Println(v15)
                                                                                            _ = v65
                                                                                        }
                                                                                    } else {
                                                                                        v67 := GenS9{F0: 2147483648, F1: 885, F2: 544}
                                                                                        var v68 bool = true
                                                                                        v69 := GenS4{F0: 4294967295, F1: false, F2: v68}
                                                                                        switch (v13 + v13) {
                                                                                        case "\n\t":
                                                                                            var v70 int32 = -595
                                                                                            _ = v70
                                                                                        case "foo bar":
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = v13
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(fi16, v13, v41)
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    }
                                                                                    _ = v41
                                                                                }
                                                                                _ = v40
                                                                            }
                                                                            var v71 uint8 = v15
                                                                            v19 -= 9223372036854775807
                                                                            _ = v71
                                                                        }
                                                                    case 32767:
                                                                        v72 := GenS7{F0: (fi14 >> ((4294967295) & 63))}
                                                                        _ = v72
                                                                    case 0:
                                                                        var v73 uint = 65535
                                                                        _ = v73
                                                                    }
                                                                    if len(v6) > 0 {
                                                                        v6[0] = (v13 + ("\n\t" + ""))
                                                                    }
                                                                    if true {
                                                                        switch "hello" {
                                                                        case "":
                                                                            fmt.Println(fi16, fi16, v4)
                                                                        default:
                                                                            v2 = float32((-2.25 - (0.0 - ((-2.25 + 1.5) * 3.14159))))
                                                                        }
                                                                    }
                                                                    fmt.Println(v19, v1)
                                                                    _ = v19
                                                                } else {
                                                                    v4 -= v4
                                                                }
                                                            } else {
                                                                v3[2147483647] = (fi16 ^ fi14)
                                                                fmt.Println(v15, v4, v15)
                                                                if (false && false) {
                                                                    if (42 > 0) {
                                                                        if (!true) {
                                                                            fmt.Println(v4, v4, v4)
                                                                            switch "a" {
                                                                            case "":
                                                                                fmt.Println(fi16)
                                                                                v74 := GenS10{F0: (v5 | (^(v5)))}
                                                                                if false {
                                                                                    v4 -= (v4 / ((v4) | 1))
                                                                                }
                                                                                _ = v74
                                                                            case "hello":
                                                                                switch (v1 * (fi14 >> ((18446744073709551615) & 63))) {
                                                                                case -490:
                                                                                    var v75 float32 = v2
                                                                                    _ = v75
                                                                                case -530:
                                                                                    var v76 uint16 = 65535
                                                                                    _ = v76
                                                                                case 9223372036854775807:
                                                                                    var v77 bool = false
                                                                                    _ = v77
                                                                                default:
                                                                                    var v78 uint64 = v4
                                                                                    v5 -= (v5 * v5)
                                                                                    _ = v78
                                                                                }
                                                                            case "foo bar":
                                                                                v79 := GenS11{F0: v5, F1: (100 >= -92), F2: v4}
                                                                                _ = v79
                                                                            default:
                                                                                if false {
                                                                                    var v80 uint16 = 2
                                                                                    _ = v80
                                                                                } else {
                                                                                    switch fi16 {
                                                                                    case 32768:
                                                                                        var v81 uint32 = 2147483648
                                                                                        if (true && false) {
                                                                                            v2 = 0.1
                                                                                            fmt.Println(v5)
                                                                                            v82 := GenS12{F0: ((42 != 255) || true)}
                                                                                            _ = v82
                                                                                        }
                                                                                        v2 = v2
                                                                                        _ = v81
                                                                                    default:
                                                                                        v83 := GenS10{F0: (v5 * v5)}
                                                                                        _ = v83
                                                                                    }
                                                                                    var v84 float32 = float32(1e10)
                                                                                    fmt.Println(fi16, v5, v5)
                                                                                    _ = v84
                                                                                }
                                                                                v6 = append(v6, v13)
                                                                            }
                                                                        } else {
                                                                            var v85 bool = (false && false)
                                                                            if v85 {
                                                                                v86 := GenS2{F0: -1, F1: 2, F2: ""}
                                                                                v87 := GenS13{F0: 42}
                                                                                var v88 float64 = float64(v2)
                                                                                _ = v86
                                                                                _ = v87
                                                                                _ = v88
                                                                            } else {
                                                                                if (!(42 == 0)) {
                                                                                    if true {
                                                                                        v3[42] = fi14
                                                                                        switch v13 {
                                                                                        case "a":
                                                                                            v2 = v2
                                                                                        default:
                                                                                            switch ("foo bar" + (((v13 + v13) + v13) + (v13 + v13))) {
                                                                                            case "hello":
                                                                                                v2 = ((v2 - -2.25) / v2)
                                                                                                if v85 {
                                                                                                    v1 -= (v1 & (fi16 & (v1 % ((fi16) | 1))))
                                                                                                } else {
                                                                                                    if len(v6) >= 2 {
                                                                                                        v6 = v6[1:len(v6)]
                                                                                                    }
                                                                                                    v4 += (v4 ^ (v4 >> ((42) & 63)))
                                                                                                    fmt.Println(v4, v4, v4)
                                                                                                }
                                                                                                v2 += v2
                                                                                                var v89 uint16 = 518
                                                                                                v2 += v2
                                                                                                switch (v13 + v13) {
                                                                                                case "a":
                                                                                                    var v90 int8 = -35
                                                                                                    _ = v90
                                                                                                case "foo bar":
                                                                                                    if v85 {
                                                                                                        var v91 int64 = -1
                                                                                                        v92 := GenS3{F0: -1, F1: (v15 >> ((0) & 63)), F2: 2}
                                                                                                        fmt.Println(v2, v4, v15)
                                                                                                        _ = v91
                                                                                                        _ = v92
                                                                                                    } else {
                                                                                                        switch (v1 - (^(v1))) {
                                                                                                        case 2147483647:
                                                                                                            _ = len(v6)
                                                                                                        case 42:
                                                                                                            v6 = append(v6, v13)
                                                                                                            var v93 bool = v85
                                                                                                            var v94 int16 = 167
                                                                                                            _ = v93
                                                                                                            _ = v94
                                                                                                        case 4294967295:
                                                                                                            var v95 uint64 = (v4 * (v4 | v4))
                                                                                                            _ = v95
                                                                                                        }
                                                                                                        fmt.Println(fi16, fi14, v1)
                                                                                                    }
                                                                                                case "hello":
                                                                                                    var v96 uint64 = (^(v4))
                                                                                                    _ = v96
                                                                                                }
                                                                                                _ = v89
                                                                                            case "foo bar":
                                                                                                v97 := GenS13{F0: 0}
                                                                                                _ = v97
                                                                                            case "a":
                                                                                                var v98 int = fi14
                                                                                                _ = v98
                                                                                            default:
                                                                                                if v85 {
                                                                                                    fmt.Println(v4, v13, v5)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v15)
                                                                                    } else {
                                                                                        v99 := GenS2{F0: -9223372036854775808, F1: 255, F2: v13}
                                                                                        v15 = 255
                                                                                        _ = v99
                                                                                    }
                                                                                }
                                                                                var v100 float64 = (1.5 + 3.14159)
                                                                                v100 = float64((v2 / (1.5 + 0.1)))
                                                                                _ = v100
                                                                            }
                                                                            var v101 int16 = (v5 * (v5 ^ (v5 * (v5 * -1))))
                                                                            _ = v101
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        v13 = v13
                                                                    }
                                                                    v102 := GenS6{F0: 265, F1: 255, F2: 0}
                                                                    fmt.Println(fi16, v1, fi16)
                                                                    if (((v4 ^ v4) == (v4 % (((^(v4))) | 1))) && ((!false) || (2147483647 == 424))) {
                                                                        switch v13 {
                                                                        case "foo bar":
                                                                            if ((v5 - (^(v5))) >= -1) {
                                                                                v103 := GenS2{F0: 0, F1: 678, F2: ("" + "\n\t")}
                                                                                _ = v103
                                                                            } else {
                                                                                fmt.Println(v2, v15, v5)
                                                                                var v104 float64 = (float64(((0.0 - v2) + (float32((0.1 - ((((-2.25 / 0.1) / 1.0) * float64(v2)) / float64(float32(-2.25))))) * v2))) * 3.14159)
                                                                                fmt.Println(v15, fi16)
                                                                                _ = v104
                                                                            }
                                                                        case "":
                                                                            fmt.Println(v2, fi16)
                                                                        default:
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            var v105 int16 = v5
                                                                            var v106 int = v1
                                                                            _ = v105
                                                                            _ = v106
                                                                        }
                                                                        v15 = (v15 * v15)
                                                                    }
                                                                    _ = v102
                                                                } else {
                                                                    var v107 uint32 = 515
                                                                    _ = v107
                                                                }
                                                            }
                                                        }
                                                        var v108 uint8 = v15
                                                        if len(v6) > 0 {
                                                            v6[0] = ((("foo bar" + ((("\n\t" + v13) + v13) + "foo bar")) + "\n\t") + ("foo bar" + "foo bar"))
                                                        }
                                                        _ = v108
                                                    } else {
                                                        v6 = append(v6, ("\n\t" + "foo bar"))
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v109 int64 = -1
                                                    _ = fi16
                                                    _ = v109
                                                }
                                                v110 := GenS4{F0: (v4 % ((v4) | 1)), F1: true, F2: (65535 < 18446744073709551615)}
                                                for fi111 := 0; fi111 < 14; fi111++ {
                                                    switch (v1 & fi111) {
                                                    case 128:
                                                        fmt.Println(fi111, fi14)
                                                    case 1:
                                                        switch (fi14 & (^(fi14))) {
                                                        case 1:
                                                            v3[550] = v1
                                                        case 0:
                                                            var v112 uint8 = (v15 / (((v15 ^ v15)) | 1))
                                                            _ = v112
                                                        case 32767:
                                                            var v113 uint = 18446744073709551615
                                                            _ = v113
                                                        }
                                                        v3[2147483647] = (v1 | (v1 ^ (v1 * (fi111 >> ((uint(v15)) & 63)))))
                                                        var v114 uint32 = 2
                                                        v15 -= v15
                                                        if len(v6) > 0 {
                                                            v6[0] = (v13 + "a")
                                                        }
                                                        if (!true) {
                                                            v115 := GenS13{F0: int32(v15)}
                                                            _ = v115
                                                        } else {
                                                            v1 = fi14
                                                            if (1 >= (v5 - 1000)) {
                                                                if len(v6) > 0 {
                                                                    v6[0] = v13
                                                                }
                                                            }
                                                        }
                                                        _ = v114
                                                    }
                                                    if (!(!((v4 >> ((413) & 63)) <= (v4 / ((v4) | 1))))) {
                                                        v116 := GenS13{F0: -2147483648}
                                                        v5 = (-(v5))
                                                        v4 -= 741
                                                        _ = v116
                                                    }
                                                    _ = fi111
                                                }
                                                _ = v110
                                            } else {
                                                for fi117 := 0; fi117 < 5; fi117++ {
                                                    v15 = v15
                                                    if (0 != 1) {
                                                        v1 = fi117
                                                        switch (fi14 / (((v1 + fi14)) | 1)) {
                                                        case -77:
                                                            delete(v3, 613)
                                                        case 127:
                                                            var v118 bool = (!(!false))
                                                            _ = v118
                                                        case 1:
                                                            switch v13 {
                                                            case "a":
                                                                v13 = v13
                                                            case "hello":
                                                                var v119 uint64 = v4
                                                                _ = v119
                                                            }
                                                        default:
                                                            if true {
                                                                switch fi14 {
                                                                case 127:
                                                                    if true {
                                                                        v1 -= (^(fi117))
                                                                        v15 += 1
                                                                    } else {
                                                                        var v120 int = -137
                                                                        _ = v120
                                                                    }
                                                                case 32767:
                                                                    var v121 float32 = float32(float64(1e10))
                                                                    _ = v121
                                                                case -961:
                                                                    var v122 int64 = -1
                                                                    _ = v122
                                                                default:
                                                                    v15 = 127
                                                                }
                                                                v13 = "a"
                                                            }
                                                        }
                                                    } else {
                                                        v1 = (-(v1))
                                                        switch v13 {
                                                        case "a":
                                                            v2 += v2
                                                        default:
                                                            var v123 int = (v1 / ((fi14) | 1))
                                                            _ = v123
                                                        }
                                                    }
                                                    v124 := v3[65535]
                                                    _ = fi117
                                                    _ = v124
                                                }
                                                var v125 uint32 = 372
                                                _ = v125
                                            }
                                            _ = v15
                                        } else {
                                            for fi126 := 0; fi126 < 19; fi126++ {
                                                if false {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                                _ = fi126
                                            }
                                        }
                                        v2 -= ((((v2 + ((v2 - v2) / (((0.0 + (v2 * ((v2 * v2) + float32(0.1)))) / -2.25) * (3.14159 - v2)))) / float32((1.5 - 0.1))) / 1.5) - -2.25)
                                    }
                                    v127 := GenS14{F0: 4294967295, F1: -128, F2: 2}
                                    switch (fi14 & v1) {
                                    case -923:
                                        if (983 < 65535) {
                                            switch v1 {
                                            case 42:
                                                var v128 int32 = 2147483647
                                                _ = v128
                                            case 127:
                                                v13 = "hello"
                                            case 2147483647:
                                                if true {
                                                    var v129 int8 = -95
                                                    v1 = fi14
                                                    _ = v129
                                                }
                                            default:
                                                _ = len(v6)
                                            }
                                        }
                                    case 32768:
                                        v4 += (v4 & v4)
                                    case 127:
                                        var v130 int32 = -2147483648
                                        _ = v130
                                    default:
                                        if (!false) {
                                            var v131 int8 = -128
                                            _ = v131
                                        } else {
                                            switch v13 {
                                            case "foo bar":
                                                _ = len(v6)
                                            case "":
                                                v13 = (("" + (v13 + v13)) + ("a" + v13))
                                            }
                                            var v132 float64 = float64(v2)
                                            v133 := GenS1{F0: -9223372036854775808, F1: -23, F2: 42}
                                            _ = v132
                                            _ = v133
                                        }
                                    }
                                    _ = v127
                                } else {
                                    if (!true) {
                                        if true {
                                            v5 += (v5 >> ((1) & 63))
                                        } else {
                                            v134 := GenS6{F0: 42, F1: 598, F2: 935}
                                            for fi135 := 0; fi135 < 17; fi135++ {
                                                switch (v1 / ((42) | 1)) {
                                                case -374:
                                                    var v136 uint64 = v4
                                                    _ = v136
                                                default:
                                                    if false {
                                                        v5 -= (v5 / ((42) | 1))
                                                        var v137 bool = (!false)
                                                        _ = v137
                                                    } else {
                                                        var v138 int32 = -2147483648
                                                        _ = v138
                                                    }
                                                }
                                                switch ("" + v13) {
                                                case "":
                                                    v139 := GenS15{F0: 2, F1: 82}
                                                    _ = v139
                                                }
                                                var v140 int8 = 0
                                                _ = fi135
                                                _ = v140
                                            }
                                            _ = v134
                                        }
                                        switch v1 {
                                        case -197:
                                            if (0 != 1) {
                                                for fi141 := 0; fi141 < 19; fi141++ {
                                                    var v142 uint8 = 2
                                                    _ = fi141
                                                    _ = v142
                                                }
                                                v1 += 128
                                            }
                                        }
                                        switch v13 {
                                        case "hello":
                                            v13 = (v13 + v13)
                                        case "\n\t":
                                            delete(v3, 0)
                                        case "":
                                            v143 := GenS16{F0: -9223372036854775808, F1: 2147483647}
                                            _ = v143
                                        default:
                                            var v144 int32 = -2147483648
                                            _ = v144
                                        }
                                    }
                                    var v145 int16 = (v5 << ((42) & 63))
                                    _ = v145
                                }
                                _ = fi14
                            }
                            if len(v6) > 0 {
                                v6[0] = (v13 + v13)
                            }
                        }
                        _ = v13
                    }
                    v2 -= v2
                } else {
                    var v146 int = v1
                    _ = v146
                }
            case 4294967295:
                var v147 uint16 = 42
                _ = v147
            case 106:
                v2 = 1.5
            }
            v5 -= v5
            if ((v4 << ((4294967295) & 63)) > v4) {
                var v148 float32 = (((v2 - (1e10 / 1.0)) - v2) - 1e10)
                _ = v148
            }
        }
    }
    var v149 int8 = 100
    if false {
        for fi150 := 0; fi150 < 18; fi150++ {
            if true {
                var v151 uint = 427
                v152 := v3[4294967295]
                if true {
                    switch v1 {
                    case 32767:
                        v2 += v2
                    default:
                        var v153 int8 = (v149 - v149)
                        _ = v153
                    }
                    for fi154 := 0; fi154 < 15; fi154++ {
                        delete(v3, 2147483647)
                        if false {
                            var v155 uint16 = 2
                            v155 += (v155 + (v155 << (((v151 - 18446744073709551615)) & 63)))
                            if (int16(v151) <= 0) {
                                if false {
                                    var v156 uint8 = 128
                                    v157 := v3[13]
                                    if (!(!false)) {
                                        switch v157 {
                                        case 42:
                                            if true {
                                                var v158 int8 = (v149 >> ((v151) & 63))
                                                _ = v158
                                            }
                                        case -278:
                                            switch ("\n\t" + "foo bar") {
                                            case "\n\t":
                                                v156 -= v156
                                            case "foo bar":
                                                switch v157 {
                                                case -2147483648:
                                                    switch (v1 & (v157 << (((v151 ^ 4294967295)) & 63))) {
                                                    case -2147483648:
                                                        v159 := GenS13{F0: 869}
                                                        _ = v159
                                                    case 42:
                                                        switch (("foo bar" + "hello") + "") {
                                                        case "a":
                                                            var v160 uint8 = (v156 * v156)
                                                            _ = v160
                                                        default:
                                                            var v161 int8 = 100
                                                            _ = v161
                                                        }
                                                    }
                                                case -752:
                                                    if (-2.25 >= -2.25) {
                                                        if false {
                                                            var v162 string = "a"
                                                            var v163 uint8 = (v156 + v156)
                                                            var v164 float64 = 0.1
                                                            _ = v162
                                                            _ = v163
                                                            _ = v164
                                                        }
                                                        v156 += (v156 << ((v151) & 63))
                                                    }
                                                case -1:
                                                    fmt.Println(v152, v4, v157)
                                                default:
                                                    delete(v3, 2147483648)
                                                }
                                            }
                                        case 2147483647:
                                            switch fi150 {
                                            case -9223372036854775808:
                                                v165 := GenS14{F0: 42, F1: 100, F2: v156}
                                                _ = v165
                                            }
                                        default:
                                            v166 := GenS17{F0: "hello", F1: 2147483648}
                                            _ = v166
                                        }
                                    } else {
                                        switch 966 {
                                        case 32767:
                                            v152 = int(v4)
                                        }
                                        switch (v152 >> (((v151 / (((v151 << (((v151 * v151)) & 63))) | 1))) & 63)) {
                                        case -2147483648:
                                            var v167 int16 = v5
                                            _ = v167
                                        case -445:
                                            switch v157 {
                                            case 128:
                                                if ((v4 >= v4) && (true && ((v156 - 1) == (v156 << (((v151 * (v151 + v151))) & 63))))) {
                                                    v4 = v4
                                                    if true {
                                                        switch ("a" + "\n\t") {
                                                        case "hello":
                                                            v157 -= -2147483648
                                                        case "foo bar":
                                                            var v168 uint = v151
                                                            _ = v168
                                                        case "":
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            if (!true) {
                                                                v169 := GenS7{F0: (v1 << ((v151) & 63))}
                                                                if true {
                                                                    var v170 uint8 = 128
                                                                    _ = v170
                                                                } else {
                                                                    var v171 int32 = 540
                                                                    _ = v171
                                                                }
                                                                v172 := GenS9{F0: 4294967295, F1: (v155 | 1), F2: 42}
                                                                _ = v169
                                                                _ = v172
                                                            } else {
                                                                if true {
                                                                    var v173 uint16 = uint16(v151)
                                                                    _ = v173
                                                                }
                                                                if (float32(float64((v2 - (v2 - (v2 + (-2.25 - (3.14159 / ((0.1 + 3.14159) * ((((v2 - v2) - -2.25) * (v2 * ((v2 - v2) + (v2 + v2)))) - v2))))))))) > float32((3.14159 + 0.1))) {
                                                                    var v174 uint = v151
                                                                    _ = v174
                                                                } else {
                                                                    var v175 int64 = 1
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    v176 := GenS18{F0: (v155 >= v155), F1: (v5 * 673), F2: int8(v5)}
                                                                    _ = v175
                                                                    _ = v176
                                                                }
                                                            }
                                                            var v177 int = 2147483647
                                                            _ = v177
                                                        default:
                                                            if false {
                                                                v5 += v5
                                                                v178 := GenS4{F0: 42, F1: true, F2: false}
                                                                v4 = (v4 - (v4 | (v4 % (((v4 / ((v4) | 1))) | 1))))
                                                                _ = v178
                                                            }
                                                        }
                                                        fmt.Println(v155)
                                                        switch (v1 | (-(fi150))) {
                                                        case -9223372036854775808:
                                                            v5 = -976
                                                        case 127:
                                                            v179 := GenS19{F0: true}
                                                            _ = v179
                                                        }
                                                    }
                                                    if (417 == -1) {
                                                        var v180 uint8 = 99
                                                        _ = v180
                                                    }
                                                    var v181 uint = v151
                                                    _ = v181
                                                }
                                            case -1:
                                                switch "foo bar" {
                                                case "hello":
                                                    v182 := GenS20{F0: 65535, F1: "foo bar"}
                                                    _ = v182
                                                default:
                                                    v156 = 2
                                                }
                                            }
                                        case 32768:
                                            if true {
                                                if (!true) {
                                                    v156 = v156
                                                    v183 := GenS21{F0: 651}
                                                    _ = v183
                                                } else {
                                                    switch v152 {
                                                    case -567:
                                                        delete(v3, 2147483648)
                                                    case -2147483648:
                                                        v155 += v155
                                                    default:
                                                        fmt.Println(v155)
                                                    }
                                                    v151 -= v151
                                                    fmt.Println(v155)
                                                    var v184 int16 = (^(v5))
                                                    v4 += v4
                                                    if false {
                                                        if true {
                                                            v185 := GenS11{F0: v184, F1: true, F2: v4}
                                                            _ = v185
                                                        }
                                                        var v186 int = (fi150 & (v1 - fi154))
                                                        _ = v186
                                                    }
                                                    if (!false) {
                                                        if (true && (1 > 181)) {
                                                            if false {
                                                                if (!(true || (("foo bar" + "a") >= ("hello" + "")))) {
                                                                    switch fi150 {
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v151)
                                                                    case 585:
                                                                        v156 = (v156 - (v156 << ((v151) & 63)))
                                                                    default:
                                                                        if (v184 != (v184 - v5)) {
                                                                            v187 := GenS7{F0: (v1 | (v1 << ((uint(v156)) & 63)))}
                                                                            v188 := GenS22{F0: (v157 * (fi150 ^ (fi154 | -2147483648)))}
                                                                            _ = v187
                                                                            _ = v188
                                                                        } else {
                                                                            var v189 uint8 = v156
                                                                            var v190 bool = false
                                                                            _ = v189
                                                                            _ = v190
                                                                        }
                                                                    }
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                } else {
                                                                    v191 := GenS21{F0: 178}
                                                                    v152 = -194
                                                                    var v192 uint32 = 4294967295
                                                                    _ = v191
                                                                    _ = v192
                                                                }
                                                                v193 := v3[540]
                                                                var v194 uint16 = (v155 ^ (^(v155)))
                                                                _ = v193
                                                                _ = v194
                                                            } else {
                                                                v149 = (v149 << (((v151 & v151)) & 63))
                                                            }
                                                        } else {
                                                            var v195 int8 = (v149 - (v149 << ((4294967295) & 63)))
                                                            v149 = 1
                                                            _ = v195
                                                        }
                                                    }
                                                    _ = v184
                                                }
                                                var v196 int = fi154
                                                _ = v196
                                            } else {
                                                _ = len(v6)
                                            }
                                        }
                                        v197 := GenS19{F0: false}
                                        if true {
                                            var v198 uint64 = 2
                                            switch fi154 {
                                            case 44:
                                                switch (v152 >> ((4294967295) & 63)) {
                                                case 0:
                                                    var v199 int8 = (v149 % (((v149 ^ (v149 * int8(v156)))) | 1))
                                                    _ = v199
                                                case 42:
                                                    var v200 int64 = -613
                                                    _ = v200
                                                }
                                            }
                                            var v201 float64 = -2.25
                                            _ = v198
                                            _ = v201
                                        }
                                        var v202 int8 = (v149 - -1)
                                        v203 := GenS9{F0: 65535, F1: (^(v155)), F2: 1}
                                        _ = v197
                                        _ = v202
                                        _ = v203
                                    }
                                    _ = v156
                                    _ = v157
                                } else {
                                    fmt.Println(v151, fi150)
                                    v204 := GenS5{F0: (!true)}
                                    _ = v204
                                }
                            } else {
                                if ((^(v155)) > v155) {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if false {
                                        v205 := GenS23{F0: 187}
                                        var v206 uint16 = (^(v155))
                                        var v207 float32 = 1.5
                                        _ = v205
                                        _ = v206
                                        _ = v207
                                    } else {
                                        v208 := GenS24{F0: int(v4), F1: -9223372036854775808, F2: (fi150 ^ v1)}
                                        v209 := GenS12{F0: (176 >= 127)}
                                        _ = v208
                                        _ = v209
                                    }
                                    switch "a" {
                                    case "hello":
                                        v6 = append(v6, "")
                                    case "a":
                                        if false {
                                            fmt.Println(fi154)
                                            if ((!true) && false) {
                                                fmt.Println(v5, fi154)
                                                _ = len(v6)
                                            } else {
                                                v2 += (float32(-2.25) / 3.14159)
                                                switch "foo bar" {
                                                case "hello":
                                                    var v210 int8 = v149
                                                    _ = v210
                                                case "a":
                                                    v211 := GenS25{F0: v149}
                                                    _ = v211
                                                case "\n\t":
                                                    fmt.Println(fi154, v151)
                                                }
                                                v1 = -472
                                            }
                                            var v212 float64 = 1e10
                                            _ = v212
                                        }
                                    case "\n\t":
                                        var v213 uint16 = 826
                                        _ = v213
                                    }
                                    var v214 int = 127
                                    _ = v214
                                }
                                v215 := GenS8{F0: -710, F1: -32768}
                                _ = v215
                            }
                            _ = v155
                        }
                        _ = fi154
                    }
                }
                _ = v151
                _ = v152
            }
            v216 := GenS20{F0: 454, F1: (("hello" + ("hello" + "foo bar")) + ("hello" + ((("foo bar" + "\n\t") + "foo bar") + (("" + (("\n\t" + "\n\t") + "hello")) + "foo bar"))))}
            _ = fi150
            _ = v216
        }
        switch v1 {
        case -312:
            for fi217 := 0; fi217 < 17; fi217++ {
                v1 = (v1 + (fi217 >> ((0) & 63)))
                var v218 float32 = float32(-2.25)
                _ = fi217
                _ = v218
            }
        case 42:
            v5 = v5
        }
        if (v4 >= 883) {
            for fi219 := 0; fi219 < 13; fi219++ {
                var v220 float64 = 3.14159
                fmt.Println(v1)
                _ = fi219
                _ = v220
            }
        }
    }
    v221 := GenS11{F0: (^(v5)), F1: (!true), F2: v4}
    v222 := GenS26{F0: ""}
    v2 -= (v2 - (1e10 - 3.14159))
    v4 = v4
    var v223 uint8 = 0
    _ = len(v3)
    if (false && (42 < 542)) {
        v224 := GenS27{F0: 0}
        v1 = (v1 << ((4294967295) & 63))
        _ = v224
    } else {
        var v225 uint = 1
        var v226 uint8 = 2
        if true {
            if (v1 == v1) {
                if (!true) {
                    v227 := v3[65535]
                    v149 += 100
                    _ = len(v6)
                    _ = v227
                } else {
                    v4 = 53
                }
                for fi228 := 0; fi228 < 20; fi228++ {
                    v5 = (v5 ^ v5)
                    fmt.Println(v5, v149)
                    switch "hello" {
                    case "hello":
                        v229 := v3[0]
                        _ = v229
                    case "foo bar":
                        switch int(v5) {
                        case 2147483647:
                            for fi230 := 0; fi230 < 13; fi230++ {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v231 := v3[0]
                                _ = fi230
                                _ = v231
                            }
                        case -190:
                            v232 := GenS27{F0: uint16(v4)}
                            _ = v232
                        }
                    default:
                        switch v1 {
                        case 1:
                            v233 := GenS28{F0: 255}
                            _ = v233
                        case -2147483648:
                            v225 += (v225 >> ((v225) & 63))
                        default:
                            var v234 uint = v225
                            _ = v234
                        }
                    }
                    _ = fi228
                }
            } else {
                if len(v6) > 0 {
                    v6[0] = "foo bar"
                }
            }
            var v235 int16 = (v5 & int16(v223))
            switch "\n\t" {
            case "":
                v3[4294967295] = 1
            }
            _ = v235
        } else {
            for fi236 := 0; fi236 < 15; fi236++ {
                v2 = 3.14159
                _ = fi236
            }
            for fi237 := 0; fi237 < 20; fi237++ {
                switch "foo bar" {
                case "\n\t":
                    v223 = 255
                case "a":
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        for fi238 := 0; fi238 < 19; fi238++ {
                            var v239 int64 = 850
                            v2 = v2
                            switch 0 {
                            case 42:
                                v240 := GenS7{F0: (^(fi238))}
                                _ = v240
                            }
                            v241 := GenS29{F0: 4294967295, F1: 307}
                            _ = fi238
                            _ = v239
                            _ = v241
                        }
                    case 843:
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                    }
                }
                v2 -= v2
                _ = fi237
            }
            v242 := GenS23{F0: 430}
            _ = v242
        }
        _ = v225
        _ = v226
    }
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("v223:", v223)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v149
    _ = v2
    _ = v221
    _ = v222
    _ = v223
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
