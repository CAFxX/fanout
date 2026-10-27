package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int64; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS8 struct { F0 int32; F1 string; F2 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -17
    var v3 string = "foo bar"
    v2 = (-(v2))
    var v4 uint = 255
    var v5 float32 = (((1e10 * -2.25) / ((((((0.1 / 3.14159) / 3.14159) * -2.25) * (1.5 - (-2.25 + float32((1.5 - 0.1))))) - (float32(((0.1 - 0.1) - float64((0.1 + 0.1)))) + 3.14159)) + 0.0)) - 1.5)
    if true {
        var v6 uint32 = 360
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 1}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: -2147483648}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := []int{42, (-(fi11)), (fi11 >> ((v4) & 63))}
        _ = fi11
        _ = v12
    }
    switch v3 {
    case "a":
        var v13 int32 = 1
        _ = v13
    case "\n\t":
        switch -997 {
        case 2147483647:
            var v14 int16 = -332
            _ = v14
        case 239:
            _ = len(v1)
        default:
            v15 := GenS1{F0: 2147483647}
            _ = v15
        }
    case "foo bar":
        var v16 int64 = -9223372036854775808
        _ = v16
    }
    v17 := GenS1{F0: -899}
    fmt.Println(v5)
    v18 := GenS1{F0: -1}
    for fi19 := 0; fi19 < 5; fi19++ {
        if true {
            v3 = (v3 + (((v3 + (v3 + "\n\t")) + ("foo bar" + ("\n\t" + v3))) + v3))
            if true {
                v20 := GenS1{F0: 0}
                var v21 float64 = 0.1
                if false {
                    if (1 <= 2) {
                        var v22 uint32 = 2147483648
                        var v23 int16 = 32767
                        v24 := GenS2{F0: 1}
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        if false {
                            switch 32767 {
                            case 4294967295:
                                if (!true) {
                                    for fi25 := 0; fi25 < 15; fi25++ {
                                        if false {
                                            if ((v4 ^ v4) == 1) {
                                                var v26 uint = (v4 ^ v4)
                                                _ = v26
                                            }
                                            switch "\n\t" {
                                            case "":
                                                if (false && true) {
                                                    if false {
                                                        v5 = (v5 - ((v5 + (v5 + 0.0)) + (0.1 + 1e10)))
                                                        if (fi19 == 0) {
                                                            v27 := GenS1{F0: 2147483647}
                                                            _ = v27
                                                        } else {
                                                            if true {
                                                                var v28 uint32 = 720
                                                                fmt.Println(v21, v4, fi25)
                                                                var v29 uint64 = 4294967295
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                            v5 += -2.25
                                                        }
                                                    } else {
                                                        v21 = float64(3.14159)
                                                        var v30 int = fi25
                                                        var v31 float64 = v21
                                                        v5 -= ((v5 / float32(3.14159)) + v5)
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    var v32 int64 = 1
                                                    _ = v32
                                                } else {
                                                    if false {
                                                        var v33 bool = true
                                                        v34 := []int{fi25}
                                                        fmt.Println(v2, v5, v21)
                                                        v21 += float64(v5)
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    fmt.Println(v5)
                                                    v1[("foo bar" + (((v3 + "a") + "foo bar") + (("foo bar" + ((v3 + (v3 + v3)) + (((v3 + v3) + "\n\t") + v3))) + v3)))] = (fi25 + fi25)
                                                }
                                            case "foo bar":
                                                var v35 int32 = 1
                                                _ = v35
                                            case "a":
                                                var v36 bool = true
                                                _ = v36
                                            }
                                            var v37 uint8 = 4
                                            _ = v37
                                        } else {
                                            fmt.Println(v21, fi25)
                                            fmt.Println(v5, v3, v2)
                                            var v38 uint16 = 2
                                            _ = v38
                                        }
                                        var v39 int64 = -490
                                        if (!(-378 > 2147483647)) {
                                            if false {
                                                var v40 uint64 = 18446744073709551615
                                                if (!false) {
                                                    v1[v3] = fi19
                                                    switch fi19 {
                                                    case -840:
                                                        if (false && (!(145 < -1))) {
                                                            switch 9223372036854775807 {
                                                            case -9223372036854775808:
                                                                var v41 int64 = -9223372036854775808
                                                                if false {
                                                                    v42 := GenS3{F0: (float32(v21) < -2.25), F1: 1}
                                                                    _ = v42
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(fi19)
                                                                    } else {
                                                                        switch (v3 + "") {
                                                                        case "hello":
                                                                            v39 += (v41 << ((v4) & 63))
                                                                        }
                                                                    }
                                                                    fmt.Println(v4)
                                                                    if (false && (36 != 2147483648)) {
                                                                        var v43 uint16 = 65535
                                                                        v3 = v3
                                                                        v41 -= int64(v4)
                                                                        _ = v43
                                                                    }
                                                                }
                                                                _ = v41
                                                            case 128:
                                                                var v44 int32 = -2147483648
                                                                _ = v44
                                                            case -1:
                                                                v45 := GenS4{F0: fi19}
                                                                _ = v45
                                                            }
                                                        } else {
                                                            v40 = v40
                                                            _ = len(v1)
                                                            fmt.Println(fi25, v2)
                                                        }
                                                        var v46 uint64 = (v40 >> ((4294967295) & 63))
                                                        if false {
                                                            var v47 uint8 = 42
                                                            _ = v47
                                                        }
                                                        var v48 int8 = v2
                                                        _ = v46
                                                        _ = v48
                                                    case 4294967295:
                                                        var v49 int = (fi19 / (((-(fi19))) | 1))
                                                        _ = v49
                                                    }
                                                }
                                                if true {
                                                    v50 := GenS4{F0: 32768}
                                                    v40 = 0
                                                    var v51 uint16 = 325
                                                    var v52 int64 = -1
                                                    v53 := GenS3{F0: true, F1: (v40 / (((^(v40))) | 1))}
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    delete(v1, (v3 + v3))
                                                    v4 = (v4 / ((v4) | 1))
                                                    var v54 int64 = (-(v39))
                                                    _ = v54
                                                }
                                                _ = v40
                                            }
                                        } else {
                                            fmt.Println(v3)
                                        }
                                        _ = fi25
                                        _ = v39
                                    }
                                    var v55 int = 343
                                    var v56 int8 = -128
                                    _ = v55
                                    _ = v56
                                }
                            case 2147483647:
                                v57 := GenS3{F0: false, F1: 951}
                                _ = v57
                            default:
                                v21 -= (float64(v5) - ((1e10 + (v21 * v21)) / float64(v5)))
                            }
                            v58 := []int{(^(fi19)), 888}
                            _ = v58
                        }
                        fmt.Println(v2, v2)
                        v3 = (("hello" + v3) + "foo bar")
                    }
                    for fi59 := 0; fi59 < 3; fi59++ {
                        v4 = (v4 << ((4294967295) & 63))
                        fmt.Println(v3)
                        _ = fi59
                    }
                    for fi60 := 0; fi60 < 3; fi60++ {
                        switch 2147483647 {
                        case -9223372036854775808:
                            v3 = v3
                        case 32767:
                            v21 += v21
                        }
                        v61 := []int{}
                        v62 := GenS2{F0: 770}
                        v4 = 349
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                } else {
                    for fi63 := 0; fi63 < 4; fi63++ {
                        v64 := []int{fi19, -191, (fi19 + (^(fi19))), 2147483647}
                        v5 += (-2.25 + v5)
                        var v65 int32 = -564
                        var v66 int64 = 263
                        if false {
                            if (!false) {
                                if len(v64) > 0 {
                                    v64[0] = 2147483647
                                }
                                switch (v3 + v3) {
                                case "":
                                    if len(v64) > 0 {
                                        v64[0] = fi19
                                    }
                                }
                                if false {
                                    switch int(v65) {
                                    case 1:
                                        var v67 uint8 = 255
                                        _ = v67
                                    case -513:
                                        switch (fi63 | 468) {
                                        case -798:
                                            v68 := GenS5{F0: 42, F1: fi63}
                                            v69 := GenS1{F0: (v65 * (v65 & -1))}
                                            v70 := v1[(v3 + v3)]
                                            _ = v68
                                            _ = v69
                                            _ = v70
                                        case 128:
                                            var v71 float32 = 3.14159
                                            _ = v71
                                        default:
                                            switch fi19 {
                                            case 493:
                                                var v72 uint32 = 4294967295
                                                _ = v72
                                            case 1:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v73 int8 = 0
                                                    _ = v73
                                                case "\n\t":
                                                    if ((fi63 ^ (fi19 - -1)) != 128) {
                                                        v4 = v4
                                                    }
                                                case "a":
                                                    v74 := GenS4{F0: fi63}
                                                    _ = v74
                                                default:
                                                    fmt.Println(fi19, v65, fi63)
                                                }
                                            default:
                                                v75 := GenS2{F0: 2}
                                                _ = v75
                                            }
                                        }
                                    default:
                                        v2 += v2
                                        v76 := GenS6{F0: -970, F1: (!true), F2: 2}
                                        _ = v76
                                    }
                                    fmt.Println(v5, v21)
                                }
                            } else {
                                var v77 string = v3
                                _ = v77
                            }
                            if (-9223372036854775808 < 32768) {
                                v21 = v21
                                v78 := GenS7{F0: 41, F1: 128, F2: 2}
                                fmt.Println(v65, v3)
                                _ = v78
                            } else {
                                v5 += ((float32(((v21 / v21) * v21)) - (v5 + (v5 + v5))) * v5)
                                var v79 uint64 = 42
                                v80 := GenS4{F0: 127}
                                v65 -= (^(v65))
                                _ = v79
                                _ = v80
                            }
                        }
                        _ = fi63
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                }
                _ = v20
                _ = v21
            }
            var v81 int16 = -1
            _ = v81
        } else {
            v82 := GenS8{F0: -1, F1: "\n\t", F2: 154}
            _ = v82
        }
        _ = fi19
    }
    var v83 uint8 = 127
    for fi84 := 0; fi84 < 5; fi84++ {
        switch (v3 + "\n\t") {
        case "hello":
            v83 = (v83 % ((uint8(v4)) | 1))
        default:
            switch v3 {
            case "hello":
                var v85 int8 = -128
                _ = v85
            case "\n\t":
                v86 := GenS5{F0: 342, F1: (fi84 & fi84)}
                _ = v86
            default:
                var v87 string = v3
                _ = v87
            }
        }
        _ = fi84
    }
    v2 -= 42
    if (0 < int64(v2)) {
        var v88 int8 = (v2 + 37)
        _ = v88
    } else {
        if (!(!false)) {
            switch (((((v3 + ("a" + v3)) + v3) + ("\n\t" + v3)) + (v3 + (v3 + "hello"))) + "a") {
            case "":
                v89 := v1[(v3 + (("" + "a") + "foo bar"))]
                _ = v89
            default:
                v3 = (v3 + (v3 + v3))
            }
            v90 := v1["a"]
            v3 = v3
            _ = v90
        }
    }
    v91 := GenS4{F0: 42}
    if false {
        v92 := GenS9{F0: 95}
        var v93 string = v3
        switch 32767 {
        case 1:
            switch (v3 + "hello") {
            case "\n\t":
                v94 := GenS10{F0: 273}
                _ = v94
            }
        case -56:
            fmt.Println(v83, v5, v93)
        }
        _ = v92
        _ = v93
    } else {
        fmt.Println(v83, v5, v4)
        v83 -= v83
        v4 += (v4 * (^(v4)))
    }
    var v95 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v91
    _ = v95
}
