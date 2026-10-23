package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 string; F2 uint8 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint16; F1 int64; F2 int8 }
type GenS10 struct { F0 uint; F1 string }
func genFn0(p0 uint64, p1 int32, p2 int8) bool {
    v7 := GenS1{F0: p2, F1: int(p2), F2: 977}
    v8 := []int{32768, 127, 42, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := []string{}
    v2 := []int{}
    v3 := []int{}
    v4 := []string{("foo bar" + "a"), "a", ("a" + "foo bar"), ((("hello" + "\n\t") + "\n\t") + "foo bar")}
    var v5 float32 = 3.14159
    v6 := make(map[uint32]uint32)
    v5 -= ((float32(3.14159) + (1.5 - 0.1)) / (-2.25 / 0.1))
    var v9 uint = 1000
    v5 += v5
    switch "foo bar" {
    case "a":
        if ((v5 - v5) == (0.0 * 0.0)) {
            if (-340 <= 250) {
                fmt.Println(v5, v9)
                v9 = (v9 - (v9 & v9))
                v5 = v5
            }
            switch 127 {
            case 937:
                v5 += float32(0.1)
            default:
                switch "foo bar" {
                case "foo bar":
                    var v10 float64 = -2.25
                    _ = v10
                case "\n\t":
                    var v11 float64 = 1e10
                    _ = v11
                }
            }
        } else {
            fmt.Println(v5)
        }
    case "":
        v9 = v9
    default:
        v12 := GenS1{F0: 42, F1: -1, F2: 184}
        _ = v12
    }
    v13 := v6[900]
    if false {
        var v14 float32 = 1e10
        switch "hello" {
        case "":
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 uint16 = 551
                for fi17 := 0; fi17 < 1; fi17++ {
                    if false {
                        fmt.Println(v5)
                        if (0 > 839) {
                            _ = len(v2)
                            var v18 uint64 = 1
                            if (32767 <= 42) {
                                var v19 uint32 = 1
                                switch (-(fi17)) {
                                case 433:
                                    v19 += 0
                                case 187:
                                    var v20 int8 = -1
                                    _ = v20
                                case 60:
                                    var v21 int64 = 806
                                    _ = v21
                                default:
                                    delete(v6, v13)
                                }
                                _ = v19
                            } else {
                                v13 -= v13
                                if false {
                                    _ = len(v6)
                                    var v22 uint8 = 0
                                    _ = v22
                                }
                                if (false && false) {
                                    switch (fi17 / ((fi15) | 1)) {
                                    case 32768:
                                        if (96 == 0) {
                                            switch "" {
                                            case "a":
                                                v14 = 1.5
                                            default:
                                                v13 -= v13
                                            }
                                            switch "a" {
                                            case "a":
                                                var v23 float64 = 1e10
                                                _ = v23
                                            case "foo bar":
                                                if false {
                                                    v5 = v5
                                                    var v24 int = -9223372036854775808
                                                    _ = v24
                                                } else {
                                                    var v25 int8 = -128
                                                    v9 = v9
                                                    v5 -= v14
                                                    _ = v25
                                                }
                                            case "hello":
                                                switch 630 {
                                                case 42:
                                                    v26 := GenS1{F0: 42, F1: (fi17 << ((42) & 63)), F2: v16}
                                                    _ = v26
                                                }
                                            }
                                        }
                                    case 1:
                                        v14 = ((v14 * (0.0 * (((v14 - 0.1) + v14) - (v14 - 0.0)))) * v5)
                                    case 32767:
                                        switch "foo bar" {
                                        case "":
                                            v27 := GenS2{F0: 107, F1: int(v16), F2: 127}
                                            _ = v27
                                        }
                                    }
                                    var v28 int16 = 0
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(fi17)
                                    }
                                    _ = v28
                                }
                            }
                            _ = v18
                        } else {
                            var v29 int = -1
                            var v30 bool = ((v29 << ((42) & 63)) < 9223372036854775807)
                            switch (fi15 % ((v29) | 1)) {
                            case 128:
                                v14 = ((0.1 - 0.1) + 0.1)
                                if true {
                                    if false {
                                        v31 := GenS3{F0: 0}
                                        v32 := GenS1{F0: 60, F1: (fi17 * fi15), F2: (v16 >> (((v9 & (^(v9)))) & 63))}
                                        v9 -= (^(v9))
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    var v33 float64 = (-2.25 - 3.14159)
                                    _ = v33
                                }
                            case 0:
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = int(v9)
                                    }
                                } else {
                                    var v34 uint32 = (^(v13))
                                    var v35 uint8 = 42
                                    var v36 int64 = 42
                                    v6[(^(v13))] = 787
                                    fmt.Println(v14, v5, v16)
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                            case -979:
                                v37 := v6[v13]
                                _ = v37
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                v38 := GenS4{F0: 775}
                                v39 := GenS5{F0: (v16 >> ((66) & 63)), F1: (("a" + "\n\t") + "a"), F2: 2}
                                _ = v38
                                _ = v39
                            } else {
                                v40 := v6[2]
                                v6[(v13 % (((v13 * v13)) | 1))] = uint32(v9)
                                v9 = v9
                                var v41 float32 = v5
                                v5 -= (v41 / v14)
                                v40 = (v40 | 461)
                                if v30 {
                                    v40 -= v13
                                }
                                _ = v40
                                _ = v41
                            }
                            _ = len(v6)
                            _ = v29
                            _ = v30
                        }
                        _ = len(v6)
                    } else {
                        var v42 float32 = v5
                        _ = v42
                    }
                    v13 += (v13 - (v13 / (((^(v13))) | 1)))
                    _ = fi17
                }
                v43 := GenS6{F0: 858, F1: 1}
                _ = fi15
                _ = v16
                _ = v43
            }
        default:
            v44 := GenS6{F0: 42, F1: 0}
            _ = v44
        }
        for fi45 := 0; fi45 < 13; fi45++ {
            v46 := GenS7{F0: v9}
            _ = fi45
            _ = v46
        }
        _ = v14
    }
    fmt.Println(v13)
    switch "hello" {
    case "\n\t":
        v5 = ((v5 * v5) / (v5 * (((v5 + ((-2.25 * v5) / float32((-2.25 / (3.14159 - 1e10))))) - -2.25) / v5)))
    default:
        switch "a" {
        case "hello":
            v13 = (v13 >> ((v9) & 63))
        case "":
            if false {
                fmt.Println(v5)
                var v47 uint = (v9 - (v9 << ((v9) & 63)))
                var v48 int = 0
                _ = v47
                _ = v48
            }
        case "a":
            delete(v6, v13)
        default:
            var v49 uint32 = v13
            _ = v49
        }
    }
    v5 += (v5 - (0.0 + (v5 + (0.0 / (float32(0.0) * (((v5 * v5) * (v5 - v5)) / (((3.14159 / v5) * v5) * v5)))))))
    if (695 != (^(v9))) {
        switch int(v9) {
        case 581:
            for fi50 := 0; fi50 < 6; fi50++ {
                if (!false) {
                    var v51 uint32 = v13
                    if false {
                        if true {
                            v13 -= v13
                            for fi52 := 0; fi52 < 3; fi52++ {
                                switch ("a" + ("" + "foo bar")) {
                                case "":
                                    var v53 uint8 = 0
                                    _ = v53
                                case "a":
                                    var v54 uint = (v9 | 280)
                                    _ = v54
                                case "foo bar":
                                    v13 = (^(v51))
                                default:
                                    switch "" {
                                    case "\n\t":
                                        var v55 int8 = 1
                                        _ = v55
                                    case "a":
                                        if (true && false) {
                                            v56 := GenS8{F0: v9, F1: -470}
                                            if false {
                                                v13 += 4294967295
                                            } else {
                                                if (true && true) {
                                                    var v57 float64 = (0.1 + (float64(v5) * 0.1))
                                                    _ = v57
                                                }
                                            }
                                            _ = len(v3)
                                            _ = v56
                                        } else {
                                            v58 := GenS7{F0: v9}
                                            _ = v58
                                        }
                                    case "":
                                        var v59 uint = (v9 << ((42) & 63))
                                        _ = v59
                                    default:
                                        v9 += v9
                                    }
                                }
                                _ = len(v4)
                                fmt.Println(v5, v5)
                                v5 = (1.5 * v5)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi52
                            }
                        } else {
                            v60 := v6[v13]
                            fmt.Println(fi50)
                            _ = v60
                        }
                    }
                    _ = v51
                }
                v61 := GenS9{F0: 901, F1: -9223372036854775808, F2: 127}
                _ = fi50
                _ = v61
            }
        case 32767:
            v13 -= (v13 ^ 791)
        default:
            if false {
                var v62 uint32 = v13
                _ = v62
            }
        }
        var v63 uint16 = 0
        var v64 int64 = 9223372036854775807
        _ = v63
        _ = v64
    }
    v5 = (1.5 / float32(3.14159))
    if (190 <= 0) {
        v13 += (v13 >> (((v9 >> ((255) & 63))) & 63))
    } else {
        if (!(v13 >= (v13 << ((v9) & 63)))) {
            v65 := GenS3{F0: -964}
            v66 := GenS5{F0: 2, F1: "foo bar", F2: 0}
            fmt.Println(v5)
            _ = v65
            _ = v66
        } else {
            for fi67 := 0; fi67 < 14; fi67++ {
                var v68 uint16 = 65535
                v68 += v68
                _ = fi67
                _ = v68
            }
            fmt.Println(v5, v13)
        }
        var v69 int64 = 9223372036854775807
        var v70 bool = true
        _ = v69
        _ = v70
    }
    var v71 string = ""
    var v72 int = 128
    if (true && true) {
        v73 := GenS8{F0: 920, F1: 32767}
        if (true || false) {
            if false {
                var v74 int8 = -1
                v72 = 1
                var v75 uint8 = 234
                _ = v74
                _ = v75
            }
            v76 := GenS10{F0: 4294967295, F1: (((v71 + (("\n\t" + "foo bar") + v71)) + (v71 + "hello")) + (v71 + ("hello" + v71)))}
            _ = v76
        } else {
            if (3.14159 >= 0.1) {
                v77 := GenS3{F0: 778}
                if true {
                    if true {
                        var v78 uint = 516
                        _ = v78
                    }
                    v79 := v6[(v13 / (((v13 >> ((v9) & 63))) | 1))]
                    v80 := GenS10{F0: 255, F1: v71}
                    _ = v79
                    _ = v80
                } else {
                    v3 = append(v3, 32768)
                    switch (((("hello" + "hello") + (v71 + "a")) + (v71 + v71)) + ("\n\t" + v71)) {
                    case "foo bar":
                        var v81 uint = v9
                        _ = v81
                    default:
                        if false {
                            if true {
                                for fi82 := 0; fi82 < 8; fi82++ {
                                    for fi83 := 0; fi83 < 12; fi83++ {
                                        var v84 int64 = -1
                                        if false {
                                            var v85 int8 = -1
                                            var v86 float32 = -2.25
                                            if len(v4) > 0 {
                                                v4[0] = "hello"
                                            }
                                            _ = v85
                                            _ = v86
                                        }
                                        v84 -= (v84 / (((v84 + v84)) | 1))
                                        fmt.Println(v71, v71, v13)
                                        _ = fi83
                                        _ = v84
                                    }
                                    v72 = v72
                                    _ = fi82
                                }
                                switch 127 {
                                case 42:
                                    v9 = v9
                                }
                            }
                            var v87 int8 = 50
                            _ = v87
                        }
                    }
                    if true {
                        v72 -= v72
                        var v88 int = (v72 % (((v72 - v72)) | 1))
                        _ = v88
                    } else {
                        v71 = v71
                    }
                }
                _ = v77
            }
        }
        _ = v73
    } else {
        var v89 int16 = 411
        _ = v89
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v90 int16 = 1000
    var v91 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v9:", v9)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
    _ = v9
    _ = v90
    _ = v91
}
