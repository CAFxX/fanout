package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int; F1 int16; F2 bool }
type GenS4 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int32; F1 uint32; F2 int64 }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint64, p1 int8) uint32 {
    var v5 float64 = 1.5
    p1 = (p1 * 100)
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{(-(fi6)), fi6, fi6}
        if false {
            p1 = p1
            fmt.Println(v5, p0)
        } else {
            var v8 uint = 1
            switch "" {
            case "\n\t":
                switch (fi6 / (((fi6 % ((fi6) | 1))) | 1)) {
                case 0:
                    p0 -= p0
                case 42:
                    v8 += (v8 << ((v8) & 63))
                case -1:
                    v8 += v8
                }
            case "a":
                p1 = -1
            case "foo bar":
                var v9 int32 = 0
                _ = v9
            default:
                v10 := GenS1{F0: p1}
                _ = v10
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 int = (fi6 / ((fi6) | 1))
                v13 := GenS2{F0: ("\n\t" + "hello"), F1: (("foo bar" + "\n\t") + "\n\t")}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    p1 -= (^(p1))
    var v14 int32 = -877
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return 412
}

func main() {
    var v1 uint32 = 620
    var v2 uint16 = 143
    var v3 float32 = (1e10 * 3.14159)
    var v4 uint16 = 511
    var v15 string = (("a" + "\n\t") + "foo bar")
    switch v15 {
    case "foo bar":
        if (346 == 42) {
            v16 := []int{2147483647, 128, 1, 127}
            _ = v16
        } else {
            v1 -= v1
            v4 -= v2
        }
    case "\n\t":
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v2, v4)
            var v18 int = 0
            _ = fi17
            _ = v18
        }
    case "hello":
        fmt.Println(v4)
    default:
        var v19 int = -1
        _ = v19
    }
    switch v15 {
    case "foo bar":
        v20 := GenS1{F0: 42}
        _ = v20
    case "hello":
        v21 := []int{32768, -799, 0}
        _ = v21
    case "a":
        var v22 int8 = 80
        _ = v22
    default:
        v23 := []int{0, -880, 32767}
        _ = v23
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        v24 := []int{-445}
        _ = v24
    case 127:
        if false {
            v25 := GenS3{F0: 32767, F1: 42, F2: (!(-270 != 0))}
            v26 := GenS3{F0: int(v1), F1: -1, F2: false}
            _ = v25
            _ = v26
        }
    case 278:
        v27 := GenS4{F0: 127, F1: (!true), F2: (v1 >> ((0) & 63))}
        _ = v27
    }
    switch "foo bar" {
    case "\n\t":
        var v28 uint16 = v2
        _ = v28
    case "hello":
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 10; fi30++ {
                if ((0 != 9223372036854775807) || (!false)) {
                    var v31 int64 = 42
                    v32 := []int{-953, fi30, int(v2), 1}
                    _ = v31
                    _ = v32
                }
                v33 := []int{(fi29 + 9223372036854775807), fi29, -1}
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                _ = fi30
                _ = v33
            }
            if ((!true) && true) {
                if true {
                    if false {
                        for fi34 := 0; fi34 < 3; fi34++ {
                            var v35 int = (fi34 % ((fi29) | 1))
                            if true {
                                v36 := []int{(^(fi29))}
                                if len(v36) > 0 {
                                    v36[0] = fi29
                                }
                                _ = len(v36)
                                _ = v36
                            }
                            v37 := GenS1{F0: 107}
                            _ = fi34
                            _ = v35
                            _ = v37
                        }
                        v38 := []int{(fi29 % (((fi29 << ((255) & 63))) | 1)), (^(fi29))}
                        v38 = append(v38, fi29)
                        _ = v38
                    } else {
                        if (!(false && (-32768 < -140))) {
                            var v39 int64 = 1
                            for fi40 := 0; fi40 < 17; fi40++ {
                                var v41 float64 = ((1.5 + 1e10) - -2.25)
                                v42 := GenS2{F0: "hello", F1: v15}
                                fmt.Println(v15)
                                var v43 int = -113
                                v44 := GenS5{F0: (fi29 / ((v43) | 1)), F1: -1}
                                _ = fi40
                                _ = v41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            if true {
                                var v45 int = fi29
                                var v46 int8 = 1
                                _ = v45
                                _ = v46
                            } else {
                                if true {
                                    switch 42 {
                                    case 4294967295:
                                        var v47 int64 = v39
                                        _ = v47
                                    case 1:
                                        v48 := GenS6{F0: 1, F1: (v1 / (((v1 | (v1 / ((2147483648) | 1)))) | 1)), F2: (v39 * (^(v39)))}
                                        _ = v48
                                    default:
                                        if true {
                                            if false {
                                                if (!(true && true)) {
                                                    var v49 int64 = v39
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch ((v15 + ("\n\t" + v15)) + "foo bar") {
                                                        case "foo bar":
                                                            if false {
                                                                v50 := GenS5{F0: -9223372036854775808, F1: 42}
                                                                switch ("\n\t" + "") {
                                                                case "":
                                                                    v51 := []int{(fi29 / (((fi29 - -1)) | 1)), (fi29 + (fi29 << ((2) & 63)))}
                                                                    _ = v51
                                                                }
                                                                v4 = v4
                                                                _ = v50
                                                            } else {
                                                                fmt.Println(v49, v39)
                                                                fmt.Println(v15)
                                                                if true {
                                                                    v49 = (^(v49))
                                                                }
                                                            }
                                                        case "a":
                                                            if true {
                                                                fmt.Println(fi29, v4, fi29)
                                                                var v52 uint16 = 42
                                                                _ = v52
                                                            }
                                                        case "\n\t":
                                                            if false {
                                                                v53 := GenS7{F0: -128}
                                                                _ = v53
                                                            } else {
                                                                var v54 string = v15
                                                                var v55 uint8 = 127
                                                                v56 := GenS6{F0: -732, F1: v1, F2: (v49 + (v49 - (^(v39))))}
                                                                _ = v54
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        default:
                                                            var v57 int8 = 42
                                                            _ = v57
                                                        }
                                                    default:
                                                        switch (fi29 / ((fi29) | 1)) {
                                                        case -9223372036854775808:
                                                            v58 := []int{(fi29 * -2), (-(fi29)), 501, fi29}
                                                            _ = v58
                                                        case -282:
                                                            v4 -= (v2 + (v2 * (v4 % ((774) | 1))))
                                                        case -1:
                                                            for fi59 := 0; fi59 < 7; fi59++ {
                                                                var v60 int64 = v39
                                                                _ = fi59
                                                                _ = v60
                                                            }
                                                        default:
                                                            v15 = "\n\t"
                                                        }
                                                    }
                                                    _ = v49
                                                }
                                                var v61 uint8 = 42
                                                if (76 == 4294967295) {
                                                    if (false || true) {
                                                        var v62 uint16 = uint16(v1)
                                                        _ = v62
                                                    } else {
                                                        v39 = (v39 | v39)
                                                        fmt.Println(v61, v15, fi29)
                                                        for fi63 := 0; fi63 < 17; fi63++ {
                                                            fmt.Println(v61, v61)
                                                            v64 := []int{(fi29 ^ (fi29 * (fi63 * fi29)))}
                                                            _ = fi63
                                                            _ = v64
                                                        }
                                                    }
                                                    var v65 int32 = 0
                                                    v2 = v2
                                                    _ = v65
                                                } else {
                                                    var v66 string = ("\n\t" + (v15 + "a"))
                                                    v67 := []int{(-(fi29)), (fi29 % (((^(fi29))) | 1)), (^(fi29)), (fi29 + fi29)}
                                                    var v68 float64 = 1.5
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                }
                                                _ = v61
                                            } else {
                                                v69 := []int{}
                                                _ = v69
                                            }
                                            v1 += v1
                                        }
                                    }
                                    var v70 int16 = 971
                                    _ = v70
                                }
                            }
                            _ = v39
                        }
                        v2 += (v4 / (((v4 % (((v4 % ((1) | 1))) | 1))) | 1))
                    }
                }
                if true {
                    for fi71 := 0; fi71 < 16; fi71++ {
                        v72 := []int{(fi71 & 9223372036854775807), 111, 128}
                        _ = fi71
                        _ = v72
                    }
                }
            }
            for fi73 := 0; fi73 < 11; fi73++ {
                v74 := []int{}
                if len(v74) >= 2 {
                    v74 = v74[1:len(v74)]
                }
                _ = fi73
                _ = v74
            }
            _ = fi29
        }
    case "foo bar":
        v1 += (v1 << ((970) & 63))
    }
    var v75 uint32 = v1
    switch 0 {
    case 337:
        if true {
            fmt.Println(v2, v75)
            for fi76 := 0; fi76 < 8; fi76++ {
                if false {
                    var v77 uint64 = 42
                    _ = v77
                } else {
                    switch v15 {
                    case "":
                        v75 -= 42
                    }
                    v78 := []int{(^(fi76)), (fi76 | -1)}
                    var v79 uint = 2
                    _ = v78
                    _ = v79
                }
                for fi80 := 0; fi80 < 16; fi80++ {
                    var v81 bool = true
                    v75 += (^(v1))
                    if true {
                        var v82 uint64 = 0
                        v75 += (^(v75))
                        var v83 uint32 = 2147483648
                        v83 += v83
                        var v84 uint16 = v4
                        _ = v82
                        _ = v83
                        _ = v84
                    }
                    _ = fi80
                    _ = v81
                }
                _ = fi76
            }
            v85 := GenS2{F0: ("a" + ""), F1: ""}
            _ = v85
        } else {
            if true {
                v86 := []int{147}
                _ = v86
            } else {
                switch 0 {
                case 42:
                    var v87 float32 = ((v3 / v3) * ((v3 * 3.14159) - (0.1 - -2.25)))
                    _ = v87
                case 192:
                    v1 -= 43
                }
                v2 -= (v2 | 39)
            }
            v88 := []int{32768}
            switch 42 {
            case 4294967295:
                if true {
                    var v89 uint32 = 372
                    for fi90 := 0; fi90 < 7; fi90++ {
                        var v91 string = (v15 + (v15 + v15))
                        if ((1e10 / 1e10) != 0.0) {
                            var v92 float64 = (float64(float32(3.14159)) / 1.5)
                            _ = v92
                        } else {
                            if true {
                                var v93 uint32 = (v89 >> ((1) & 63))
                                if len(v88) > 0 {
                                    v88[0] = (fi90 % (((fi90 / ((fi90) | 1))) | 1))
                                }
                                fmt.Println(fi90)
                                _ = v93
                            } else {
                                for fi94 := 0; fi94 < 9; fi94++ {
                                    v1 = v89
                                    _ = fi94
                                }
                            }
                        }
                        _ = fi90
                        _ = v91
                    }
                    _ = v89
                }
            }
            _ = v88
        }
    case 846:
        if false {
            var v95 float64 = 1.5
            fmt.Println(v15, v75, v15)
            for fi96 := 0; fi96 < 8; fi96++ {
                v4 = v4
                var v97 uint = 4294967295
                fmt.Println(fi96, v3, fi96)
                _ = fi96
                _ = v97
            }
            _ = v95
        } else {
            var v98 int8 = 62
            _ = v98
        }
    default:
        var v99 uint = 255
        _ = v99
    }
    v15 = v15
    v2 -= v2
    var v100 uint = 18446744073709551615
    var v101 int = 9223372036854775807
    if false {
        var v102 float32 = 0.0
        _ = v102
    }
    switch ((("hello" + "\n\t") + "") + (v15 + ((v15 + v15) + "a"))) {
    case "hello":
        if false {
            for fi103 := 0; fi103 < 20; fi103++ {
                v104 := GenS7{F0: 42}
                if ((!(v15 == (v15 + v15))) && (v4 > uint16(v1))) {
                    if (!(false && (v3 == v3))) {
                        v1 = v1
                    }
                    var v105 string = v15
                    _ = v105
                } else {
                    var v106 int8 = -128
                    var v107 uint32 = 959
                    _ = v106
                    _ = v107
                }
                if (0 == 128) {
                    v15 = v15
                } else {
                    var v108 int64 = 0
                    _ = v108
                }
                _ = fi103
                _ = v104
            }
            v109 := []int{(v101 << ((v100) & 63)), v101}
            fmt.Println(v3, v15)
            _ = v109
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v100
    _ = v101
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v75
}
