package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 string; F2 uint16 }
type GenS6 struct { F0 uint; F1 bool; F2 uint8 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint; F1 int8; F2 uint64 }
type GenS13 struct { F0 string; F1 uint; F2 bool }
type GenS14 struct { F0 int16; F1 bool; F2 uint8 }
type GenS15 struct { F0 uint16; F1 uint; F2 uint16 }
func genFn0(p0 string, p1 int32) int8 {
    v6 := []int{2147483647}
    var v7 int64 = -1
    p1 = int32(v7)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 := GenS1{F0: false, F1: (v7 + (v7 << ((240) & 63)))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 bool = ((-9223372036854775808 == -9223372036854775808) || (false && (!true)))
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := []int{}
    var v5 int16 = 1
    v9 := GenS1{F0: v1, F1: 9223372036854775807}
    for fi10 := 0; fi10 < 10; fi10++ {
        v5 = (v5 * v5)
        _ = fi10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v3[(fi11 / ((fi11) | 1))] = (-(fi11))
        var v12 int64 = -529
        v12 += (v12 % (((v12 - (v12 % ((v12) | 1)))) | 1))
        _ = fi11
        _ = v12
    }
    if (v5 >= -32768) {
        v13 := GenS1{F0: v1, F1: -9223372036854775808}
        if (-128 < -128) {
            v1 = v1
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 int32 = -1
                v1 = (42 != 42)
                _ = fi14
                _ = v15
            }
            var v16 uint32 = 4294967295
            _ = v16
        }
        _ = v13
    } else {
        if v1 {
            var v17 int16 = (v5 - v5)
            v18 := GenS2{F0: ((42 < 127) || v1)}
            _ = v17
            _ = v18
        } else {
            var v19 int16 = v5
            var v20 int = 0
            if v1 {
                if (780 == 2147483647) {
                    fmt.Println(v1, v20, v5)
                    var v21 string = "\n\t"
                    _ = v21
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v19 = int16(v20)
                }
                var v22 uint = 18446744073709551615
                v5 -= (v19 / (((v19 | v5)) | 1))
                _ = v22
            } else {
                for fi23 := 0; fi23 < 7; fi23++ {
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v1 = (!v1)
                        v4 = append(v4, fi23)
                        v1 = (!true)
                        _ = fi24
                    }
                    _ = fi23
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    _ = len(v3)
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v27 := GenS3{F0: 4294967295, F1: 42}
                        v28 := GenS2{F0: false}
                        _ = len(v2)
                        v29 := GenS4{F0: 108}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    if (!v1) {
                        v30 := GenS3{F0: 2147483648, F1: 1}
                        switch ("" + "a") {
                        case "hello":
                            _ = len(v4)
                        case "foo bar":
                            v31 := GenS3{F0: 244, F1: 0}
                            _ = v31
                        case "":
                            for fi32 := 0; fi32 < 19; fi32++ {
                                switch v20 {
                                case 32768:
                                    if ((!true) && ((-(v19)) > v5)) {
                                        var v33 uint8 = 6
                                        delete(v2, 2)
                                        _ = v33
                                    }
                                case 127:
                                    v34 := GenS2{F0: false}
                                    _ = v34
                                default:
                                    v35 := GenS3{F0: 2147483648, F1: 2147483647}
                                    _ = v35
                                }
                                if v1 {
                                    if len(v4) > 0 {
                                        v4[0] = (fi25 * 9223372036854775807)
                                    }
                                    v4 = append(v4, fi25)
                                    var v36 int32 = 2147483647
                                    _ = v36
                                }
                                _ = fi32
                            }
                        }
                        _ = v30
                    } else {
                        var v37 int32 = 42
                        _ = v37
                    }
                    _ = fi25
                }
                fmt.Println(v1)
            }
            _ = v19
            _ = v20
        }
        var v38 float64 = 1e10
        v39 := GenS4{F0: 127}
        _ = v38
        _ = v39
    }
    var v40 int = 9223372036854775807
    if (42 >= -1) {
        if (487 != 0) {
            var v41 string = ""
            switch 4294967295 {
            case 1:
                if v1 {
                    var v42 int64 = -1
                    v43 := GenS5{F0: -103, F1: v41, F2: 42}
                    switch -907 {
                    case -957:
                        var v44 uint16 = 1
                        _ = v44
                    default:
                        if v1 {
                            v5 += 1000
                            var v45 int16 = (^(v5))
                            v46 := v3[(^(v40))]
                            _ = v45
                            _ = v46
                        }
                    }
                    _ = v42
                    _ = v43
                }
            case 886:
                for fi47 := 0; fi47 < 18; fi47++ {
                    v5 = (v5 & (v5 >> ((255) & 63)))
                    _ = fi47
                }
            case 128:
                v41 = v41
            default:
                v40 -= v40
            }
            _ = v41
        } else {
            var v48 int16 = (v5 ^ v5)
            _ = len(v4)
            _ = v48
        }
        switch (v40 % ((9223372036854775807) | 1)) {
        case -9223372036854775808:
            if v1 {
                for fi49 := 0; fi49 < 11; fi49++ {
                    v50 := GenS3{F0: 4294967295, F1: 4294967295}
                    v51 := GenS6{F0: 186, F1: v1, F2: 15}
                    if false {
                        v40 -= 665
                        v5 = 32767
                    }
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                for fi52 := 0; fi52 < 6; fi52++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if (255 >= 64) {
                        v53 := GenS4{F0: genFn0("\n\t", 42)}
                        var v54 uint32 = 282
                        for fi55 := 0; fi55 < 1; fi55++ {
                            fmt.Println(fi52)
                            if (0 < 18446744073709551615) {
                                switch ("a" + "a") {
                                case "a":
                                    v40 = 0
                                case "foo bar":
                                    v54 = 0
                                default:
                                    var v56 uint64 = 690
                                    _ = v56
                                }
                                var v57 float32 = 1e10
                                v58 := GenS7{F0: (v1 || true), F1: 683}
                                _ = v57
                                _ = v58
                            } else {
                                var v59 int64 = 0
                                _ = v59
                            }
                            _ = fi55
                        }
                        _ = v53
                        _ = v54
                    }
                    if false {
                        switch -617 {
                        case -1:
                            v60 := GenS3{F0: 2147483647, F1: 4294967295}
                            _ = v60
                        case -385:
                            v40 = -410
                        default:
                            for fi61 := 0; fi61 < 6; fi61++ {
                                if (2 != 2147483647) {
                                    var v62 int32 = 1
                                    var v63 int = 4294967295
                                    v64 := v3[(^(v40))]
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                switch 128 {
                                case 9223372036854775807:
                                    if true {
                                        fmt.Println(fi52, fi52)
                                        v65 := GenS8{F0: 106, F1: 1, F2: 42}
                                        fmt.Println(fi61, v40)
                                        _ = v65
                                    } else {
                                        v66 := GenS2{F0: (18446744073709551615 <= 0)}
                                        _ = v66
                                    }
                                case 2147483647:
                                    if len(v4) > 0 {
                                        v4[0] = fi61
                                    }
                                }
                                fmt.Println(fi52)
                                _ = fi61
                            }
                        }
                        v1 = false
                    } else {
                        fmt.Println(v5, fi52)
                        if false {
                            if len(v4) > 0 {
                                v4[0] = (fi52 / ((247) | 1))
                            }
                            for fi67 := 0; fi67 < 3; fi67++ {
                                var v68 string = (("foo bar" + (("\n\t" + "foo bar") + "hello")) + "\n\t")
                                v69 := GenS9{F0: -413}
                                _ = fi67
                                _ = v68
                                _ = v69
                            }
                            for fi70 := 0; fi70 < 14; fi70++ {
                                fmt.Println(v1)
                                _ = fi70
                            }
                        } else {
                            v71 := GenS7{F0: (((("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) + (("foo bar" + "foo bar") + "\n\t")) + "hello") >= "\n\t"), F1: 1}
                            v72 := GenS10{F0: -243}
                            _ = v71
                            _ = v72
                        }
                        for fi73 := 0; fi73 < 3; fi73++ {
                            switch "foo bar" {
                            case "hello":
                                var v74 int32 = 2147483647
                                _ = v74
                            default:
                                v3[fi52] = 961
                            }
                            v75 := GenS10{F0: 0}
                            v1 = v1
                            _ = fi73
                            _ = v75
                        }
                    }
                    _ = fi52
                }
            } else {
                switch 42 {
                case -831:
                    if (-1 <= -128) {
                        v5 = (-(v5))
                        if false {
                            v5 = (-(v5))
                            _ = len(v4)
                        }
                    }
                }
            }
        default:
            v76 := v2[65535]
            _ = v76
        }
        for fi77 := 0; fi77 < 15; fi77++ {
            v40 = v40
            v78 := GenS4{F0: 67}
            _ = fi77
            _ = v78
        }
    }
    fmt.Println(v40, v40, v40)
    v79 := GenS11{F0: 255}
    for fi80 := 0; fi80 < 6; fi80++ {
        switch -9223372036854775808 {
        case 127:
            v5 -= v5
        default:
            var v81 int8 = genFn0((("\n\t" + "hello") + "a"), 42)
            _ = v81
        }
        switch (^(v40)) {
        case 328:
            if v1 {
                fmt.Println(v1)
            }
        default:
            for fi82 := 0; fi82 < 5; fi82++ {
                var v83 int16 = 42
                _ = fi82
                _ = v83
            }
        }
        if false {
            if (v1 || false) {
                v5 = (v5 / ((0) | 1))
                v1 = v1
            }
            switch "\n\t" {
            case "\n\t":
                for fi84 := 0; fi84 < 20; fi84++ {
                    v2[2] = (!(42 >= -9223372036854775808))
                    v85 := GenS12{F0: 265, F1: -1, F2: 2}
                    _ = fi84
                    _ = v85
                }
            case "":
                v1 = true
            case "a":
                v86 := GenS7{F0: v1, F1: 255}
                _ = v86
            }
            for fi87 := 0; fi87 < 19; fi87++ {
                fmt.Println(v1)
                fmt.Println(fi87)
                fmt.Println(fi80, v1)
                _ = fi87
            }
        } else {
            v88 := GenS13{F0: "hello", F1: 1, F2: true}
            v1 = ((1e10 + (-2.25 / 1.5)) > -2.25)
            v5 += (v5 + (v5 >> ((0) & 63)))
            _ = v88
        }
        _ = fi80
    }
    v89 := GenS13{F0: "hello", F1: 1, F2: v1}
    var v90 uint8 = 255
    v91 := GenS5{F0: -116, F1: "foo bar", F2: 14}
    v92 := GenS14{F0: v5, F1: (v1 || v1), F2: (v90 / ((v90) | 1))}
    fmt.Println(v5, v1, v5)
    v93 := GenS15{F0: 255, F1: 197, F2: 0}
    for fi94 := 0; fi94 < 5; fi94++ {
        if true {
            var v95 uint8 = 71
            _ = v95
        }
        v96 := GenS15{F0: 0, F1: 18446744073709551615, F2: 571}
        _ = fi94
        _ = v96
    }
    if true {
        v97 := GenS12{F0: 604, F1: -128, F2: 42}
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v79
    _ = v89
    _ = v9
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
