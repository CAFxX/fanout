package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint16; F1 string; F2 uint16 }
type GenS5 struct { F0 int8; F1 int64; F2 int64 }
func genFn0() bool {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    if (true || (!((false || false) && false))) {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int16 = -1
            fmt.Println(fi9, v10)
            _ = fi9
            _ = v10
        }
        var v11 int64 = 0
        var v12 uint16 = 255
        _ = v11
        _ = v12
    }
    v13 := []int{}
    switch 2147483647 {
    case 358:
        var v14 bool = false
        _ = v14
    case 32767:
        if true {
            v15 := GenS1{F0: 2147483647, F1: 939}
            var v16 int = 88
            v16 += -75
            _ = v15
            _ = v16
        }
    case 9223372036854775807:
        var v17 uint8 = 112
        _ = v17
    }
    v18 := GenS1{F0: 461, F1: -595}
    _ = v13
    _ = v18
    return (18446744073709551615 < 2)
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{v1, 2147483648, v1, v1}
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v5 := []string{"", "\n\t", "\n\t"}
    v6 := make(map[string]string)
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 float32 = 3.14159
        v20 -= float32((1.5 * 3.14159))
        _ = fi19
        _ = v20
    }
    var v21 int64 = -1
    v22 := GenS1{F0: (v1 + v1), F1: 1}
    for fi23 := 0; fi23 < 20; fi23++ {
        _ = len(v5)
        v24 := GenS1{F0: (^(v1)), F1: -1}
        if true {
            var v25 float32 = float32(0.0)
            if (0.1 < 3.14159) {
                v25 += v25
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 uint8 = 255
                    var v28 string = "\n\t"
                    _ = len(v3)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 20; fi29++ {
                    v30 := GenS1{F0: 2147483648, F1: 42}
                    _ = fi29
                    _ = v30
                }
            }
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    var v31 float32 = ((-2.25 * 0.0) + (((0.0 * 1e10) / 1.0) / 1e10))
    switch "\n\t" {
    case "a":
        var v32 uint16 = 2
        _ = v32
    default:
        var v33 float32 = (v31 + (float32(1e10) * float32((float64(((float32(((float64(-2.25) + 1e10) * -2.25)) / v31) / float32(-2.25))) + (3.14159 / 0.1)))))
        _ = v33
    }
    switch "a" {
    case "":
        if false {
            fmt.Println(v21, v21)
            for fi34 := 0; fi34 < 1; fi34++ {
                if false {
                    v31 += v31
                    for fi35 := 0; fi35 < 8; fi35++ {
                        v1 += (v1 / ((1) | 1))
                        var v36 float64 = (1e10 - (0.1 / (float64(float32(float64(1.5))) - (1.5 - ((float64(v31) - (((0.0 * 1.5) + ((((3.14159 * (1.5 + -2.25)) / ((-2.25 + 1e10) - 0.0)) / (1.5 * 0.1)) - 1.5)) - 1e10)) + (float64(1.5) - 1.5))))))
                        switch (^(fi35)) {
                        case 0:
                            var v37 bool = true
                            _ = v37
                        default:
                            v21 = v21
                        }
                        _ = fi35
                        _ = v36
                    }
                }
                for fi38 := 0; fi38 < 2; fi38++ {
                    var v39 float32 = -2.25
                    v1 = (^(v1))
                    var v40 int = (fi34 / ((fi34) | 1))
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                _ = fi34
            }
            v21 = (v21 / (((v21 / ((-509) | 1))) | 1))
        } else {
            switch ("foo bar" + "\n\t") {
            case "a":
                for fi41 := 0; fi41 < 5; fi41++ {
                    v1 = 42
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                    _ = fi41
                }
            case "foo bar":
                if false {
                    var v42 int16 = -32768
                    var v43 int64 = -9223372036854775808
                    _ = v42
                    _ = v43
                }
            }
            fmt.Println(v31, v21, v1)
            if len(v5) > 0 {
                v5[0] = "a"
            }
        }
    case "foo bar":
        var v44 int = 0
        _ = v44
    case "\n\t":
        switch 0 {
        case 2147483647:
            switch "" {
            case "a":
                for fi45 := 0; fi45 < 14; fi45++ {
                    if true {
                        v21 -= v21
                        v21 = (v21 << ((4294967295) & 63))
                        var v46 string = "\n\t"
                        _ = v46
                    }
                    _ = fi45
                }
            }
        case 9223372036854775807:
            if false {
                v21 = 42
            } else {
                var v47 float64 = ((-2.25 - 3.14159) * 0.0)
                v31 = float32(v47)
                v48 := GenS2{F0: 215, F1: true}
                _ = v47
                _ = v48
            }
        case -2147483648:
            for fi49 := 0; fi49 < 13; fi49++ {
                if false {
                    var v50 uint16 = 1
                    v50 = v50
                    _ = v50
                } else {
                    if (!((-2.25 + (0.1 / 3.14159)) >= (1.5 * -2.25))) {
                        for fi51 := 0; fi51 < 18; fi51++ {
                            var v52 float64 = float64(v31)
                            var v53 int8 = -1
                            if false {
                                if (!false) {
                                    v31 = v31
                                }
                                var v54 uint16 = 65535
                                var v55 float64 = -2.25
                                _ = v54
                                _ = v55
                            } else {
                                v31 = v31
                                v56 := GenS1{F0: (^(v1)), F1: 352}
                                _ = v56
                            }
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        if (0 == -531) {
                            v1 = 658
                        } else {
                            var v57 float64 = 1.5
                            var v58 uint8 = 127
                            _ = v57
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 19; fi59++ {
                            v21 = v21
                            var v60 uint = 6
                            _ = fi59
                            _ = v60
                        }
                    }
                    fmt.Println(v1, fi49, v31)
                }
                _ = fi49
            }
        default:
            fmt.Println(v21, v31)
        }
    }
    v61 := GenS3{F0: 0, F1: -1}
    for fi62 := 0; fi62 < 16; fi62++ {
        var v63 float64 = 1.5
        if (false && true) {
            if (881 != 104) {
                if true {
                    v31 = v31
                } else {
                    for fi64 := 0; fi64 < 19; fi64++ {
                        if true {
                            var v65 int32 = -2147483648
                            v66 := GenS1{F0: v1, F1: 0}
                            v1 -= (v1 & 2)
                            if true {
                                v67 := GenS1{F0: 547, F1: 1000}
                                v63 -= (v63 / (-2.25 * v63))
                                v68 := v3[42]
                                _ = v67
                                _ = v68
                            }
                            _ = v65
                            _ = v66
                        } else {
                            switch fi62 {
                            case 1:
                                v69 := GenS1{F0: v1, F1: -32768}
                                _ = v69
                            case -928:
                                delete(v3, -294)
                            }
                            if true {
                                v2 = append(v2, v1)
                                v63 = -2.25
                            }
                            var v70 int = (fi62 * (fi64 + (fi64 / ((127) | 1))))
                            _ = v70
                        }
                        if false {
                            v63 = (float64((v31 + v31)) - (v63 / (0.1 + v63)))
                            var v71 float32 = float32(float64((0.1 / -2.25)))
                            var v72 uint64 = 1
                            var v73 bool = false
                            v5 = append(v5, "a")
                            _ = v71
                            _ = v72
                            _ = v73
                        } else {
                            switch "foo bar" {
                            case "hello":
                                v21 -= 0
                            case "":
                                if true {
                                    if (255 < 142) {
                                        v74 := GenS2{F0: 128, F1: true}
                                        _ = v74
                                    } else {
                                        var v75 int16 = -381
                                        _ = v75
                                    }
                                    var v76 int32 = int32(v1)
                                    if false {
                                        v5 = append(v5, "foo bar")
                                    } else {
                                        if true {
                                            v77 := GenS3{F0: 255, F1: 1}
                                            v31 = (-2.25 - 1e10)
                                            _ = v77
                                        }
                                        fmt.Println(v63)
                                        if ((true || (168 != (v1 | (v1 | 2147483648)))) || ((("hello" + "foo bar") == ("a" + (("hello" + "") + ""))) || ((18446744073709551615 != 255) && true))) {
                                            if true {
                                                switch 628 {
                                                case 32768:
                                                    v21 = (v21 - v21)
                                                }
                                            }
                                            var v78 int8 = -1
                                            v79 := GenS2{F0: 1, F1: false}
                                            _ = v78
                                            _ = v79
                                        }
                                        v80 := GenS4{F0: 42, F1: "a", F2: 42}
                                        _ = v80
                                    }
                                    if true {
                                        var v81 int16 = 42
                                        var v82 int8 = -118
                                        fmt.Println(v1, v21, v76)
                                        _ = v81
                                        _ = v82
                                    }
                                    _ = v76
                                } else {
                                    switch (fi62 | 63) {
                                    case 809:
                                        if len(v5) > 0 {
                                            v5[0] = ("" + "foo bar")
                                        }
                                    }
                                }
                                v83 := GenS3{F0: 255, F1: 100}
                                var v84 bool = true
                                _ = v83
                                _ = v84
                            default:
                                var v85 bool = (42 < 100)
                                switch "a" {
                                case "hello":
                                    v31 += (v31 - v31)
                                case "\n\t":
                                    v86 := GenS5{F0: -1, F1: -161, F2: -1}
                                    _ = v86
                                default:
                                    if (696 <= 18446744073709551615) {
                                        v1 += v1
                                        _ = len(v6)
                                        var v87 int = fi64
                                        _ = v87
                                    }
                                }
                                v31 = (float32(((v63 / (1.5 / (-2.25 - (v63 * -2.25)))) * ((v63 - ((((1.5 / v63) - v63) - v63) + 0.1)) * (((v63 / -2.25) + 1e10) / float64((1e10 / 1.0)))))) / v31)
                                _ = v85
                            }
                        }
                        _ = fi64
                    }
                }
            }
            var v88 int32 = 821
            fmt.Println(v31, v88)
            _ = v88
        }
        _ = fi62
        _ = v63
    }
    switch 941 {
    case -9223372036854775808:
        fmt.Println(v31, v21, v1)
    default:
        var v89 uint64 = 42
        _ = v89
    }
    v90 := GenS1{F0: v1, F1: 42}
    v1 += (^(v1))
    v31 -= v31
    if true {
        var v91 uint64 = 9
        var v92 uint8 = 69
        v93 := GenS1{F0: 4294967295, F1: 1}
        _ = v91
        _ = v92
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v90
}
