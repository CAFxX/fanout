package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16; F1 int16; F2 uint64 }
type GenS10 struct { F0 int; F1 int16 }
type GenS11 struct { F0 uint64; F1 bool; F2 bool }
type GenS12 struct { F0 string; F1 uint64; F2 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 float32 = (1e10 / 1.5)
    v4 := make(map[int]string)
    var v5 int64 = -395
    var v6 string = "hello"
    fmt.Println(v6, v3, v3)
    if true {
        var v7 int = -748
        _ = v7
    } else {
        if true {
            switch (v6 + "") {
            case "":
                v8 := GenS1{F0: 53}
                _ = v8
            case "foo bar":
                var v9 uint = 286
                _ = v9
            case "\n\t":
                v3 = (float32(3.14159) - float32(1e10))
            }
            if true {
                switch ((v6 + v6) + v6) {
                case "":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v11 := GenS1{F0: 4294967295}
                        _ = fi10
                        _ = v11
                    }
                }
                fmt.Println(v5, v3, v3)
            }
            if (true || false) {
                if true {
                    fmt.Println(v5, v3, v6)
                }
                v12 := v4[-236]
                v3 = v3
                _ = v12
            } else {
                if false {
                    if (2147483647 < 4294967295) {
                        if (0 < 136) {
                            fmt.Println(v3, v3, v5)
                            v5 -= (v5 - -209)
                        } else {
                            if (160 == 128) {
                                switch 42 {
                                case 2147483647:
                                    v13 := v2[31]
                                    _ = v13
                                case 1:
                                    switch v6 {
                                    case "foo bar":
                                        v3 = (0.0 + v3)
                                    case "":
                                        var v14 uint8 = 0
                                        _ = v14
                                    }
                                default:
                                    v15 := GenS1{F0: 715}
                                    _ = v15
                                }
                                v16 := []int{610, 42}
                                _ = v16
                            }
                            v3 = 1e10
                            var v17 uint16 = 1
                            _ = v17
                        }
                        v18 := []int{-27, 658, 32768}
                        var v19 int64 = (v5 | (v5 ^ 42))
                        _ = v18
                        _ = v19
                    }
                    v20 := []int{305, 127, 735}
                    v21 := GenS2{F0: true}
                    _ = v20
                    _ = v21
                }
            }
        } else {
            v5 -= (v5 >> ((1) & 63))
        }
        v22 := []int{32767}
        if len(v22) > 0 {
            v22[0] = -916
        }
        _ = v22
    }
    v5 = v5
    v3 = 0.0
    fmt.Println(v6, v5)
    if false {
        v2[2147483648] = ""
    }
    var v23 int8 = 10
    for fi24 := 0; fi24 < 1; fi24++ {
        if (1 > 65535) {
            if true {
                var v25 int64 = (^(v5))
                _ = v25
            }
            if (2147483648 != 146) {
                fmt.Println(v5, v23)
                for fi26 := 0; fi26 < 10; fi26++ {
                    var v27 int8 = v23
                    v28 := GenS2{F0: (false && true)}
                    delete(v4, fi26)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                switch "hello" {
                case "a":
                    if false {
                        v3 = ((-2.25 / 3.14159) * (float32(1e10) - v3))
                        v29 := GenS2{F0: (42 < -32768)}
                        _ = v29
                    } else {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            if false {
                                v31 := GenS3{F0: v23}
                                if (128 > 2) {
                                    var v32 float64 = 1.5
                                    fmt.Println(v3, fi30, v5)
                                    _ = v32
                                }
                                if true {
                                    var v33 string = "a"
                                    if (2 > 674) {
                                        if (!((-2.25 - float64(float32((1.5 / (0.0 - (1.5 * (0.0 - (3.14159 * 1.5)))))))) <= 0.1)) {
                                            var v34 int8 = v23
                                            _ = v34
                                        }
                                        switch v6 {
                                        case "\n\t":
                                            v5 = (v5 - v5)
                                        case "":
                                            if false {
                                                if true {
                                                    if (1 == 255) {
                                                        v35 := GenS2{F0: true}
                                                        _ = v35
                                                    }
                                                    v36 := GenS2{F0: false}
                                                    _ = v36
                                                } else {
                                                    v23 -= 100
                                                    switch v33 {
                                                    case "foo bar":
                                                        v5 = v5
                                                    case "a":
                                                        switch (fi24 & fi30) {
                                                        case 127:
                                                            fmt.Println(v6, v5, v33)
                                                        case 32767:
                                                            if (!true) {
                                                                v3 = v3
                                                            } else {
                                                                v37 := GenS4{F0: v5, F1: ""}
                                                                _ = v37
                                                            }
                                                        default:
                                                            switch "hello" {
                                                            case "":
                                                                fmt.Println(v23, v6, fi24)
                                                            }
                                                        }
                                                        if true {
                                                            v33 = v33
                                                        }
                                                    case "":
                                                        v38 := GenS5{F0: 42, F1: (v23 / (((v23 & v23)) | 1))}
                                                        _ = v38
                                                    default:
                                                        v39 := GenS6{F0: 618, F1: 255, F2: (^(v5))}
                                                        var v40 uint64 = 2
                                                        v41 := GenS6{F0: 65535, F1: 255, F2: int64(v23)}
                                                        _ = v39
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    var v42 string = "\n\t"
                                                    _ = v42
                                                }
                                            } else {
                                                v43 := GenS7{F0: -398, F1: 0}
                                                fmt.Println(fi24, fi24)
                                                _ = v43
                                            }
                                        default:
                                            var v44 int = (fi30 & 32768)
                                            _ = v44
                                        }
                                        var v45 uint16 = 854
                                        _ = v45
                                    } else {
                                        if true {
                                            v46 := GenS5{F0: 741, F1: (v23 << ((1) & 63))}
                                            var v47 uint16 = 2
                                            var v48 uint64 = 1
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        } else {
                                            v49 := []int{-146, (fi30 + fi24)}
                                            if ((false && (42 >= fi30)) || (true || (65535 > 2))) {
                                                v4[1] = v33
                                            } else {
                                                var v50 int8 = (v23 & 127)
                                                _ = v50
                                            }
                                            _ = v49
                                        }
                                        var v51 int32 = 27
                                        var v52 float32 = v3
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 float64 = 0.0
                                    var v54 uint = 1
                                    _ = v33
                                    _ = v53
                                    _ = v54
                                } else {
                                    var v55 uint16 = 0
                                    v6 = ((("hello" + v6) + "\n\t") + "")
                                    _ = v55
                                }
                                _ = v31
                            }
                            v4[(fi24 / (((fi24 - (^(fi24)))) | 1))] = (v6 + (v6 + v6))
                            var v56 uint8 = 79
                            _ = fi30
                            _ = v56
                        }
                        if ((!false) && false) {
                            delete(v2, 2147483647)
                            if true {
                                v6 = v6
                                var v57 int = -1
                                _ = v57
                            }
                        }
                        var v58 uint = 496
                        _ = v58
                    }
                case "hello":
                    v23 -= 30
                }
            } else {
                switch (fi24 >> ((4294967295) & 63)) {
                case 32768:
                    var v59 uint = 327
                    _ = v59
                case 2147483647:
                    var v60 float32 = v3
                    _ = v60
                case -2147483648:
                    for fi61 := 0; fi61 < 2; fi61++ {
                        v6 = v6
                        v5 += v5
                        fmt.Println(v6)
                        _ = fi61
                    }
                default:
                    for fi62 := 0; fi62 < 4; fi62++ {
                        if true {
                            if (!(!true)) {
                                var v63 string = ""
                                _ = v63
                            }
                            v64 := []int{}
                            v65 := GenS8{F0: 42}
                            _ = v64
                            _ = v65
                        }
                        var v66 uint64 = 1
                        v67 := GenS3{F0: (v23 ^ -1)}
                        _ = fi62
                        _ = v66
                        _ = v67
                    }
                }
                var v68 uint16 = 617
                _ = v68
            }
            fmt.Println(fi24, fi24, v23)
        }
        for fi69 := 0; fi69 < 10; fi69++ {
            switch fi24 {
            case 32768:
                v6 = ""
            case 471:
                v70 := GenS2{F0: (true && (!true))}
                _ = v70
            default:
                v23 += (-(v23))
            }
            if false {
                if true {
                    v3 += 1e10
                    v71 := GenS4{F0: v5, F1: "hello"}
                    if (1 != 18446744073709551615) {
                        var v72 uint32 = 2
                        if (!false) {
                            switch (fi24 >> ((0) & 63)) {
                            case 0:
                                var v73 int = 42
                                _ = v73
                            case 42:
                                if false {
                                    if (2147483648 == v72) {
                                        v74 := GenS8{F0: v5}
                                        fmt.Println(v3, fi69, v72)
                                        var v75 bool = false
                                        _ = v74
                                        _ = v75
                                    } else {
                                        if true {
                                            var v76 int8 = -34
                                            fmt.Println(fi24)
                                            fmt.Println(v76)
                                            _ = v76
                                        } else {
                                            var v77 int8 = v23
                                            var v78 uint32 = v72
                                            fmt.Println(v77)
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                } else {
                                    v79 := GenS4{F0: -9223372036854775808, F1: v6}
                                    _ = v79
                                }
                            case 501:
                                v80 := GenS9{F0: 42, F1: 1, F2: 441}
                                _ = v80
                            default:
                                v3 += (float32(1.5) / float32((0.0 - float64(((v3 * v3) + float32(1.5))))))
                            }
                            v81 := GenS8{F0: (v5 ^ 9223372036854775807)}
                            v6 = "a"
                            _ = v81
                        }
                        _ = v72
                    } else {
                        v82 := GenS7{F0: -1, F1: 2147483647}
                        v83 := []int{0}
                        _ = v82
                        _ = v83
                    }
                    v84 := []int{fi24, 1, fi69}
                    v85 := GenS10{F0: -1, F1: 63}
                    switch (fi24 / ((fi24) | 1)) {
                    case -1:
                        v86 := GenS11{F0: 0, F1: ((false || ((fi24 >> ((0) & 63)) >= fi24)) && true), F2: (!(false && (0.0 < 1.5)))}
                        _ = v86
                    }
                    if true {
                        if false {
                            var v87 int16 = 368
                            _ = v87
                        } else {
                            var v88 float64 = (-2.25 / (float64(3.14159) * 0.1))
                            _ = v88
                        }
                        fmt.Println(v3)
                    } else {
                        v89 := GenS12{F0: "a", F1: 4294967295, F2: (255 >= 255)}
                        var v90 int8 = 42
                        _ = v89
                        _ = v90
                    }
                    var v91 int32 = 42
                    fmt.Println(fi24)
                    var v92 float32 = (float32((float64((v3 + v3)) + 0.0)) / (-2.25 * v3))
                    _ = v71
                    _ = v84
                    _ = v85
                    _ = v91
                    _ = v92
                } else {
                    if (false || (0 == -2147483648)) {
                        v93 := v4[(fi69 / ((fi69) | 1))]
                        switch "" {
                        case "foo bar":
                            v94 := v4[127]
                            _ = v94
                        default:
                            v95 := []int{fi69, 768}
                            _ = v95
                        }
                        delete(v2, 1)
                        _ = v93
                    }
                }
                var v96 string = (v6 + v6)
                _ = v96
            } else {
                fmt.Println(fi24, v3)
                v97 := GenS11{F0: 1, F1: (!(false || true)), F2: (0.0 < v3)}
                _ = v97
            }
            _ = fi69
        }
        switch (("foo bar" + ("a" + v6)) + "hello") {
        case "foo bar":
            v5 -= (v5 * v5)
        }
        _ = fi24
    }
    var v98 float64 = 3.14159
    for fi99 := 0; fi99 < 14; fi99++ {
        if (v98 <= 1.5) {
            var v100 float64 = v98
            _ = v100
        } else {
            v3 += (v3 * (v3 * 1e10))
        }
        _ = fi99
    }
    var v101 int8 = (v23 ^ v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v101
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
