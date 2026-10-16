package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS12 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0() uint32 {
    v5 := GenS1{F0: 50}
    var v6 uint8 = 140
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 int16 = 250
    v2 := []string{("\n\t" + ("foo bar" + (("hello" + "") + ("foo bar" + "hello")))), "a"}
    v3 := make(map[int]int)
    var v4 int32 = -832
    fmt.Println(v1)
    switch 128 {
    case -9223372036854775808:
        var v7 float64 = 0.0
        _ = v7
    case -907:
        v8 := GenS2{F0: 18446744073709551615, F1: 0}
        _ = v8
    case -66:
        v1 += v1
    }
    switch -596 {
    case 2147483647:
        _ = len(v2)
    case 4294967295:
        v9 := GenS2{F0: 755, F1: 712}
        _ = v9
    case 128:
        switch 32768 {
        case -9223372036854775808:
            var v10 uint16 = 375
            _ = v10
        }
    }
    v11 := GenS3{F0: 48, F1: 2, F2: (^(v4))}
    if (-1 == 1) {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 float32 = float32(1.5)
            v14 := GenS2{F0: 42, F1: 867}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v1, v4)
    switch "a" {
    case "":
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 int32 = v4
                _ = fi16
                _ = v17
            }
            v18 := GenS3{F0: 1, F1: 583, F2: v4}
            _ = fi15
            _ = v18
        }
    case "hello":
        fmt.Println(v4, v4)
    case "\n\t":
        var v19 uint64 = 18446744073709551615
        _ = v19
    default:
        for fi20 := 0; fi20 < 15; fi20++ {
            if true {
                if (1.5 < float32((1e10 / 1.0))) {
                    switch "foo bar" {
                    case "a":
                        v1 -= v1
                    case "":
                        if false {
                            var v21 int64 = 42
                            _ = v21
                        }
                    }
                    var v22 float32 = (0.0 + 0.0)
                    v4 += -761
                    _ = v22
                } else {
                    v23 := GenS4{F0: (v4 - (v4 | v4))}
                    delete(v3, fi20)
                    v24 := GenS4{F0: v4}
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 18; fi25++ {
                    v1 += v1
                    _ = fi25
                }
            }
            if false {
                if true {
                    v26 := GenS5{F0: "hello"}
                    v1 = v1
                    v27 := GenS4{F0: v4}
                    _ = v26
                    _ = v27
                } else {
                    if true {
                        var v28 int32 = 1
                        fmt.Println(v1, fi20)
                        var v29 float32 = -2.25
                        _ = v28
                        _ = v29
                    }
                    if true {
                        v2 = append(v2, ("hello" + ("hello" + "a")))
                    } else {
                        var v30 uint32 = 0
                        var v31 uint8 = 2
                        delete(v3, 32767)
                        _ = v30
                        _ = v31
                    }
                    if (fi20 > fi20) {
                        v1 -= -1
                    }
                }
                var v32 float64 = float64(0.0)
                _ = v32
            } else {
                var v33 uint32 = 0
                if false {
                    var v34 uint8 = 128
                    v3[fi20] = int(v33)
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v33 += (v33 + v33)
                        var v36 float32 = float32(-2.25)
                        if ((v1 - 1) <= v1) {
                            fmt.Println(fi35, v1, v34)
                        }
                        var v37 uint64 = 18446744073709551615
                        var v38 int16 = v1
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = v34
                } else {
                    v1 -= v1
                }
                _ = v33
            }
            _ = fi20
        }
    }
    if (!(3.14159 == 1e10)) {
        v4 = v4
        var v39 int32 = v4
        switch ("" + "hello") {
        case "":
            switch "a" {
            case "foo bar":
                switch "a" {
                case "\n\t":
                    var v40 float64 = 3.14159
                    _ = v40
                }
            case "":
                if (false || (false || (1 >= 255))) {
                    switch -364 {
                    case 2147483647:
                        if (65535 >= 316) {
                            var v41 bool = true
                            var v42 int32 = v4
                            _ = v41
                            _ = v42
                        }
                    case 544:
                        v4 -= v39
                    }
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v44 := GenS5{F0: "\n\t"}
                        _ = fi43
                        _ = v44
                    }
                } else {
                    for fi45 := 0; fi45 < 6; fi45++ {
                        switch (fi45 ^ 68) {
                        case -2147483648:
                            if false {
                                if true {
                                    var v46 int64 = -282
                                    var v47 int64 = v46
                                    v46 -= (v46 % (((v47 * 1)) | 1))
                                    _ = v46
                                    _ = v47
                                }
                                for fi48 := 0; fi48 < 17; fi48++ {
                                    _ = len(v2)
                                    if (!(!(true && (0.1 >= 1.5)))) {
                                        v4 = int32(fi48)
                                        fmt.Println(v1, v1)
                                    }
                                    if true {
                                        var v49 int64 = 42
                                        var v50 uint32 = 2147483648
                                        switch "\n\t" {
                                        case "hello":
                                            if (v1 >= (v1 & 22)) {
                                                var v51 float32 = float32(-2.25)
                                                _ = v51
                                            }
                                        default:
                                            if (!(((("a" + "hello") == ("foo bar" + "")) && true) && (283 < 18446744073709551615))) {
                                                var v52 bool = (2 < 2)
                                                v53 := GenS6{F0: v52}
                                                v54 := GenS5{F0: ""}
                                                var v55 float32 = 3.14159
                                                var v56 int = (fi48 >> ((960) & 63))
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            } else {
                                                v57 := GenS7{F0: 27, F1: v50}
                                                v58 := GenS1{F0: 100}
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                        switch fi45 {
                                        case 1:
                                            if (1 > 2147483647) {
                                                switch ("foo bar" + (("hello" + ("a" + (("a" + "") + "\n\t"))) + "\n\t")) {
                                                case "hello":
                                                    v39 = (v39 >> ((65535) & 63))
                                                case "a":
                                                    v49 = v49
                                                case "foo bar":
                                                    if false {
                                                        if len(v2) > 0 {
                                                            v2[0] = "foo bar"
                                                        }
                                                        var v59 uint = 42
                                                        if (!(255 <= 80)) {
                                                            v50 += uint32(v59)
                                                            delete(v3, fi45)
                                                        } else {
                                                            v4 = -1
                                                        }
                                                        _ = v59
                                                    }
                                                default:
                                                    fmt.Println(v50, v1, fi48)
                                                }
                                                var v60 float32 = 1.5
                                                switch fi45 {
                                                case -2147483648:
                                                    if false {
                                                        v61 := GenS1{F0: 42}
                                                        _ = v61
                                                    } else {
                                                        v1 -= v1
                                                    }
                                                default:
                                                    switch fi45 {
                                                    case -2147483648:
                                                        v62 := GenS3{F0: 1, F1: 42, F2: (v4 * v39)}
                                                        _ = v62
                                                    }
                                                }
                                                _ = v60
                                            } else {
                                                v63 := GenS6{F0: (float64(float32((0.0 / 1e10))) < 1e10)}
                                                v64 := GenS8{F0: 42, F1: 2, F2: 2}
                                                v65 := GenS3{F0: 32767, F1: 0, F2: v4}
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            }
                                        default:
                                            v66 := GenS9{F0: 1, F1: 351}
                                            _ = v66
                                        }
                                        if (1 <= 1) {
                                            v67 := v3[(^(fi45))]
                                            v49 += (v49 | 328)
                                            v3[(fi48 - (v67 + fi45))] = (fi45 + -1)
                                            _ = v67
                                        }
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v68 := GenS5{F0: ("a" + "foo bar")}
                                        if true {
                                            var v69 int32 = 868
                                            if true {
                                                var v70 int32 = 2147483647
                                                _ = len(v2)
                                                if true {
                                                    v1 = (v1 | (v1 >> ((132) & 63)))
                                                    v70 += (v39 / (((^(v70))) | 1))
                                                }
                                                _ = v70
                                            }
                                            _ = v69
                                        }
                                        v39 = v39
                                        if false {
                                            v4 -= (v4 | v39)
                                        }
                                        _ = v68
                                    }
                                    var v71 float64 = 1e10
                                    _ = fi48
                                    _ = v71
                                }
                            } else {
                                if (fi45 > (fi45 >> ((1) & 63))) {
                                    if true {
                                        v72 := GenS10{F0: -1, F1: "foo bar"}
                                        fmt.Println(fi45, v39)
                                        fmt.Println(v1, fi45)
                                        _ = v72
                                    } else {
                                        if false {
                                            var v73 string = ("" + "hello")
                                            v74 := GenS10{F0: -318, F1: v73}
                                            for fi75 := 0; fi75 < 18; fi75++ {
                                                v73 = ("\n\t" + v73)
                                                var v76 uint16 = 702
                                                var v77 int16 = (v1 + v1)
                                                _ = fi75
                                                _ = v76
                                                _ = v77
                                            }
                                            _ = v73
                                            _ = v74
                                        }
                                        var v78 float64 = 1e10
                                        fmt.Println(v39, v4)
                                        _ = v78
                                    }
                                    v79 := GenS11{F0: 1, F1: 1, F2: 1}
                                    for fi80 := 0; fi80 < 9; fi80++ {
                                        v81 := GenS3{F0: 825, F1: 0, F2: v4}
                                        _ = fi80
                                        _ = v81
                                    }
                                    _ = v79
                                }
                            }
                        case 9223372036854775807:
                            v82 := GenS12{F0: 0, F1: 0, F2: 0}
                            _ = v82
                        case -9223372036854775808:
                            _ = len(v3)
                        default:
                            var v83 uint64 = 4294967295
                            _ = v83
                        }
                        if true {
                            var v84 float64 = 0.1
                            _ = v84
                        } else {
                            for fi85 := 0; fi85 < 16; fi85++ {
                                switch (fi45 ^ fi45) {
                                case 42:
                                    var v86 uint16 = 65535
                                    _ = v86
                                case 32768:
                                    v87 := GenS6{F0: ((4294967295 > 42) || false)}
                                    _ = v87
                                default:
                                    var v88 int64 = -1
                                    _ = v88
                                }
                                var v89 uint32 = 1
                                _ = fi85
                                _ = v89
                            }
                            v39 -= (v4 / ((v39) | 1))
                        }
                        _ = fi45
                    }
                    var v90 uint32 = 42
                    _ = v90
                }
            default:
                if (1 > 255) {
                    v4 = (v39 / ((v39) | 1))
                }
            }
        case "\n\t":
            v4 += v39
        default:
            var v91 uint8 = 42
            _ = v91
        }
        _ = v39
    } else {
        v4 += -788
    }
    var v92 uint16 = 911
    v2 = append(v2, "foo bar")
    var v93 int16 = (-(v1))
    v3[127] = 42
    _ = len(v2)
    v92 += v92
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v92
    _ = v93
}
