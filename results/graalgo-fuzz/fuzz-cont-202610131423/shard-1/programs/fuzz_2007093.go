package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int16) int8 {
    v7 := []int{}
    var v8 uint = 18446744073709551615
    switch "\n\t" {
    case "hello":
        if (true || (!false)) {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: (("" + ("\n\t" + "\n\t")) + ("foo bar" + "hello")), F1: 42, F2: 631}
                _ = fi9
                _ = v10
            }
            var v11 uint = 255
            switch 1 {
            case -231:
                for fi12 := 0; fi12 < 9; fi12++ {
                    v11 += 42
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v7 = append(v7, (fi13 << ((v11) & 63)))
                        var v14 int32 = 1
                        v7 = append(v7, fi12)
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
            default:
                for fi15 := 0; fi15 < 7; fi15++ {
                    p0 = (p0 - -32768)
                    v11 = 65535
                    _ = fi15
                }
            }
            _ = v11
        }
    case "a":
        p0 = (^(p0))
    default:
        p0 -= (p0 | (p0 / ((1000) | 1)))
    }
    v8 = v8
    _ = p0
    _ = v7
    _ = v8
    return -128
}
func genFn1() bool {
    if (!true) {
        v16 := GenS1{F0: "\n\t", F1: 32767, F2: -617}
        v17 := GenS1{F0: "\n\t", F1: 388, F2: -544}
        _ = v16
        _ = v17
    }
    var v18 bool = (false && true)
    _ = v18
    return (((3.14159 * 3.14159) < (0.0 * 1.5)) && true)
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 int16 = 32767
    var v5 int16 = 32767
    v6 := make(map[int]string)
    var v19 int32 = 0
    v20 := []int{}
    if false {
        switch 4294967295 {
        case 2147483647:
            if false {
                v5 = (v5 / ((v5) | 1))
            }
        case -1:
            v5 -= 42
        case 127:
            v19 = (v19 << ((18446744073709551615) & 63))
        default:
            v21 := GenS1{F0: "\n\t", F1: 362, F2: 643}
            _ = v21
        }
        var v22 int8 = 42
        if true {
            if false {
                switch -9223372036854775808 {
                case 430:
                    if false {
                        v22 = (v22 ^ (v22 & 127))
                    }
                case -1:
                    v4 = (v4 - v4)
                case 9223372036854775807:
                    var v23 uint32 = 4294967295
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 int8 = 42
                        if true {
                            var v26 float64 = (3.14159 - -2.25)
                            v19 = (v19 | -132)
                            v25 -= v22
                            _ = v26
                        } else {
                            v22 -= v22
                            v5 = v4
                            v27 := GenS2{F0: 42, F1: 42}
                            _ = v27
                        }
                        if (!false) {
                            var v28 int = fi24
                            _ = v28
                        } else {
                            var v29 int8 = (v25 | v22)
                            switch (("\n\t" + "a") + "") {
                            case "hello":
                                fmt.Println(v22, v25)
                            case "\n\t":
                                for fi30 := 0; fi30 < 3; fi30++ {
                                    var v31 uint64 = 2
                                    _ = fi30
                                    _ = v31
                                }
                            case "foo bar":
                                v4 = v5
                            }
                            if true {
                                v32 := GenS1{F0: ("\n\t" + "a"), F1: 255, F2: v19}
                                var v33 uint16 = 244
                                _ = v32
                                _ = v33
                            } else {
                                if false {
                                    if (false || (false && ((fi24 ^ 0) != -662))) {
                                        for fi34 := 0; fi34 < 5; fi34++ {
                                            if ((false || true) && true) {
                                                var v35 int32 = (v19 & (v19 & v19))
                                                var v36 float64 = -2.25
                                                v25 = v22
                                                v36 += float64(float32(float64(-2.25)))
                                                if false {
                                                    var v37 uint16 = 474
                                                    v25 += v29
                                                    v5 = -1
                                                    _ = v37
                                                } else {
                                                    var v38 uint = 383
                                                    v39 := GenS2{F0: 32767, F1: 130}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                _ = v35
                                                _ = v36
                                            }
                                            if (("\n\t" + "a") != "") {
                                                v19 = (v19 | v19)
                                            }
                                            v40 := GenS1{F0: "\n\t", F1: 32767, F2: -2147483648}
                                            _ = fi34
                                            _ = v40
                                        }
                                        var v41 int8 = (v25 << ((42) & 63))
                                        v42 := GenS2{F0: 189, F1: 2}
                                        _ = v41
                                        _ = v42
                                    }
                                    v19 -= (v19 ^ (v19 / (((v19 % (((v19 / ((-879) | 1))) | 1))) | 1)))
                                }
                                if (true && true) {
                                    if true {
                                        fmt.Println(v19, v22, v22)
                                        var v43 float64 = (0.0 / float64(0.1))
                                        var v44 float64 = (v43 + float64((((0.0 / 1.0) / 1e10) - 1e10)))
                                        _ = v43
                                        _ = v44
                                    }
                                    v45 := GenS3{F0: v19}
                                    _ = v45
                                }
                                fmt.Println(v19, v25, v5)
                            }
                            _ = v29
                        }
                        _ = fi24
                        _ = v25
                    }
                }
                if true {
                    fmt.Println(v22, v22, v4)
                    v19 += -2147483648
                    if false {
                        var v46 bool = (v22 < (v22 | (v22 % (((v22 / ((v22) | 1))) | 1))))
                        var v47 float32 = 1e10
                        _ = v46
                        _ = v47
                    } else {
                        if len(v20) > 0 {
                            v20[0] = -627
                        }
                        v19 -= -2147483648
                    }
                } else {
                    fmt.Println(v4, v22, v22)
                    switch ("" + "") {
                    case "":
                        v4 += (v4 % (((v5 + v4)) | 1))
                    case "foo bar":
                        var v48 int32 = v19
                        _ = v48
                    case "a":
                        for fi49 := 0; fi49 < 1; fi49++ {
                            v50 := GenS4{F0: "\n\t", F1: 65535}
                            v19 = v19
                            v4 += v4
                            _ = fi49
                            _ = v50
                        }
                    }
                    fmt.Println(v19)
                }
                if true {
                    if (-2.25 > -2.25) {
                        var v51 bool = (42 >= 4294967295)
                        _ = v51
                    }
                    var v52 float64 = 0.1
                    for fi53 := 0; fi53 < 2; fi53++ {
                        var v54 uint32 = 2
                        if false {
                            v55 := GenS3{F0: v19}
                            _ = v55
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                var v56 uint8 = 127
                                _ = v56
                            case "\n\t":
                                v52 = float64(1.5)
                            case "a":
                                v54 = v54
                            default:
                                switch "hello" {
                                case "hello":
                                    _ = len(v20)
                                case "\n\t":
                                    var v57 int32 = v19
                                    _ = v57
                                case "":
                                    var v58 uint64 = 42
                                    _ = v58
                                }
                            }
                            if true {
                                v59 := GenS5{F0: (v19 % ((v19) | 1))}
                                _ = v59
                            } else {
                                var v60 uint64 = 444
                                for fi61 := 0; fi61 < 18; fi61++ {
                                    _ = len(v20)
                                    v62 := GenS3{F0: -527}
                                    _ = fi61
                                    _ = v62
                                }
                                var v63 int32 = v19
                                _ = v60
                                _ = v63
                            }
                        }
                        _ = fi53
                        _ = v54
                    }
                    _ = v52
                } else {
                    v64 := GenS3{F0: (v19 * (v19 >> ((1) & 63)))}
                    v2[913] = false
                    var v65 float64 = 3.14159
                    _ = v64
                    _ = v65
                }
            } else {
                var v66 int64 = 42
                for fi67 := 0; fi67 < 9; fi67++ {
                    _ = len(v20)
                    fmt.Println(v66, v66)
                    _ = fi67
                }
                _ = v66
            }
            var v68 string = (("hello" + "hello") + "")
            v68 = v68
            _ = v68
        }
        _ = v22
    }
    v69 := GenS2{F0: 255, F1: 0}
    for fi70 := 0; fi70 < 12; fi70++ {
        switch 32767 {
        case 42:
            fmt.Println(v5)
        case 128:
            fmt.Println(v19, fi70, v4)
        case 4294967295:
            v5 += v4
        default:
            fmt.Println(v4, fi70, v19)
        }
        _ = fi70
    }
    if true {
        var v71 float64 = 1e10
        if (255 > 42) {
            v72 := GenS6{F0: (v5 / ((v5) | 1)), F1: 0}
            v73 := GenS7{F0: v19}
            var v74 float64 = 0.1
            _ = v72
            _ = v73
            _ = v74
        } else {
            v75 := GenS1{F0: "", F1: 255, F2: v19}
            v71 = 0.1
            switch 218 {
            case -317:
                v20 = append(v20, 170)
            case 32768:
                fmt.Println(v4, v4)
            }
            _ = v75
        }
        _ = v71
    }
    v76 := v3[1]
    if (false && (3.14159 < (((3.14159 / 1.5) / 1e10) / 1.0))) {
        if true {
            v4 = (-(v4))
        }
        var v77 uint16 = 1
        _ = v77
    }
    var v78 uint64 = 0
    if (("\n\t" == "a") && (127 < 42)) {
        for fi79 := 0; fi79 < 4; fi79++ {
            switch fi79 {
            case -9223372036854775808:
                v4 += v5
            case -1:
                v80 := GenS8{F0: "hello"}
                _ = v80
            }
            v76 = (v76 >> ((1) & 63))
            v81 := GenS2{F0: 0, F1: 2}
            _ = fi79
            _ = v81
        }
        v82 := GenS6{F0: 544, F1: 190}
        _ = v82
    }
    v4 -= (^(v5))
    var v83 uint = 26
    if (((false || (!(v76 != v76))) || true) && true) {
        switch "a" {
        case "":
            var v84 int16 = (v5 + (v5 / (((v4 + 0)) | 1)))
            _ = v84
        case "foo bar":
            var v85 float32 = (0.1 / 3.14159)
            _ = v85
        default:
            var v86 float64 = (0.1 * 3.14159)
            _ = v86
        }
        v4 = 0
        var v87 uint = v83
        _ = v87
    }
    for fi88 := 0; fi88 < 14; fi88++ {
        v89 := GenS9{F0: 127, F1: 502}
        for fi90 := 0; fi90 < 13; fi90++ {
            v78 += v78
            v19 = (-(v19))
            _ = fi90
        }
        for fi91 := 0; fi91 < 6; fi91++ {
            var v92 int8 = 50
            var v93 string = "foo bar"
            _ = fi91
            _ = v92
            _ = v93
        }
        _ = fi88
        _ = v89
    }
    for fi94 := 0; fi94 < 9; fi94++ {
        for fi95 := 0; fi95 < 13; fi95++ {
            var v96 float32 = 0.0
            v19 -= (^(v19))
            if ((v5 < (v5 - 0)) || (!(!((fi94 | fi95) >= (fi94 << (((v83 - (v83 + v83))) & 63)))))) {
                switch "\n\t" {
                case "a":
                    fmt.Println(v96, v96, v5)
                default:
                    v5 += -32768
                    if false {
                        var v97 uint64 = uint64(v76)
                        _ = v97
                    } else {
                        v4 += 32767
                        v76 = 2
                    }
                }
            }
            _ = fi95
            _ = v96
        }
        if false {
            fmt.Println(v5, v5)
        }
        _ = fi94
    }
    for fi98 := 0; fi98 < 18; fi98++ {
        for fi99 := 0; fi99 < 17; fi99++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            fmt.Println(fi98, v5)
            var v100 uint32 = 42
            _ = fi99
            _ = v100
        }
        fmt.Println(v19, v78)
        _ = fi98
    }
    var v101 int16 = v4
    fmt.Println(v78)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v101
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v76
    _ = v78
    _ = v83
}
