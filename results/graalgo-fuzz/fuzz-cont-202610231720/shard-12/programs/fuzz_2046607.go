package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 string; F1 uint8; F2 int16 }
type GenS4 struct { F0 bool; F1 int16; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 uint; F2 string }
type GenS7 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 string }
func genFn0(p0 string) string {
    for fi5 := 0; fi5 < 3; fi5++ {
        switch ((p0 + "hello") + "foo bar") {
        case "hello":
            if ((fi5 & 128) > 2147483647) {
                v6 := []int{(fi5 & 2147483647), 42, (fi5 + (fi5 >> ((4294967295) & 63))), (fi5 * fi5)}
                switch p0 {
                case "a":
                    for fi7 := 0; fi7 < 19; fi7++ {
                        if (-128 <= -128) {
                            v8 := GenS1{F0: fi5}
                            _ = v8
                        }
                        v9 := GenS2{F0: 1, F1: 4294967295}
                        _ = fi7
                        _ = v9
                    }
                case "\n\t":
                    v10 := GenS2{F0: 32767, F1: 2}
                    _ = v10
                case "":
                    if ((198 < 42) || true) {
                        p0 = "foo bar"
                        p0 = ""
                    }
                default:
                    switch (^(fi5)) {
                    case -2147483648:
                        if (!false) {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                if len(v6) > 0 {
                                    v6[0] = 9223372036854775807
                                }
                                switch (fi11 + -2147483648) {
                                case 42:
                                    p0 = p0
                                case -578:
                                    var v12 float64 = 3.14159
                                    _ = v12
                                case 2147483647:
                                    if (0 < 112) {
                                        if (65535 != 42) {
                                            var v13 int8 = 0
                                            fmt.Println(fi11, p0, p0)
                                            _ = v13
                                        } else {
                                            if (0 > 705) {
                                                var v14 string = ""
                                                _ = v14
                                            }
                                            var v15 uint = 369
                                            _ = v15
                                        }
                                    }
                                default:
                                    _ = len(v6)
                                }
                                switch (("" + "a") + (p0 + p0)) {
                                case "a":
                                    p0 = p0
                                default:
                                    v16 := GenS1{F0: (fi5 + fi5)}
                                    _ = v16
                                }
                                v17 := GenS3{F0: p0, F1: 0, F2: -387}
                                if true {
                                    var v18 uint16 = 255
                                    v6 = append(v6, 0)
                                    var v19 string = (p0 + (p0 + p0))
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 float32 = 3.14159
                                    v6 = append(v6, -2147483648)
                                    _ = v20
                                }
                                _ = fi11
                                _ = v17
                            }
                            if false {
                                fmt.Println(p0, p0)
                                v21 := GenS4{F0: (!true), F1: -32768, F2: (((("foo bar" + p0) + p0) + "\n\t") + p0)}
                                if (1 <= 0) {
                                    for fi22 := 0; fi22 < 4; fi22++ {
                                        switch fi22 {
                                        case 128:
                                            p0 = p0
                                        default:
                                            if len(v6) > 0 {
                                                v6[0] = (fi22 >> ((273) & 63))
                                            }
                                        }
                                        if true {
                                            v6 = append(v6, (fi5 ^ fi22))
                                        }
                                        _ = fi22
                                    }
                                    v6 = append(v6, 352)
                                    var v23 uint8 = 2
                                    _ = v23
                                } else {
                                    if (fi5 >= fi5) {
                                        var v24 int8 = 100
                                        v6 = append(v6, fi5)
                                        if ((584 < 255) || false) {
                                            switch (^(fi5)) {
                                            case 32768:
                                                var v25 bool = (!true)
                                                _ = v25
                                            }
                                            var v26 int32 = -2147483648
                                            _ = v26
                                        } else {
                                            v27 := GenS4{F0: (928 >= 65535), F1: 42, F2: ("\n\t" + (p0 + p0))}
                                            _ = v27
                                        }
                                        _ = v24
                                    }
                                    switch 0 {
                                    case 9223372036854775807:
                                        switch 9223372036854775807 {
                                        case -954:
                                            v28 := GenS4{F0: (989 == 728), F1: 32767, F2: p0}
                                            _ = v28
                                        case 4294967295:
                                            v6 = append(v6, fi5)
                                        }
                                    }
                                }
                                _ = v21
                            } else {
                                v29 := GenS1{F0: 157}
                                _ = v29
                            }
                        }
                    case 42:
                        p0 = p0
                    case -881:
                        for fi30 := 0; fi30 < 8; fi30++ {
                            if ((482 <= 2) || false) {
                                var v31 int64 = 9223372036854775807
                                p0 = p0
                                _ = v31
                            }
                            _ = fi30
                        }
                    }
                }
                _ = v6
            }
        case "a":
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 uint8 = 0
                v34 := GenS3{F0: (p0 + p0), F1: v33, F2: 42}
                var v35 bool = (255 < 42)
                v36 := GenS4{F0: v35, F1: 1, F2: (p0 + p0)}
                v33 = v33
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
                _ = v36
            }
        default:
            if (915 < -335) {
                switch -677 {
                case 42:
                    v37 := GenS3{F0: ("a" + ("a" + p0)), F1: 127, F2: 32767}
                    _ = v37
                default:
                    var v38 int16 = -1
                    _ = v38
                }
                p0 = (p0 + p0)
                var v39 int16 = 1000
                _ = v39
            } else {
                for fi40 := 0; fi40 < 16; fi40++ {
                    switch (fi40 >> ((906) & 63)) {
                    case -9223372036854775808:
                        if false {
                            fmt.Println(fi5, fi40, fi40)
                            p0 = "foo bar"
                            var v41 float32 = 1e10
                            _ = v41
                        } else {
                            var v42 uint8 = 0
                            _ = v42
                        }
                    case 2147483647:
                        if false {
                            p0 = ""
                            if false {
                                if ((fi40 | fi40) < fi5) {
                                    fmt.Println(fi40)
                                    p0 = p0
                                    if false {
                                        v43 := GenS5{F0: 18446744073709551615}
                                        v44 := GenS3{F0: ("\n\t" + ((p0 + "foo bar") + (p0 + (p0 + (p0 + (p0 + p0)))))), F1: 7, F2: -77}
                                        p0 = p0
                                        _ = v43
                                        _ = v44
                                    }
                                } else {
                                    var v45 int = -573
                                    p0 = p0
                                    _ = v45
                                }
                            }
                        }
                    case -1:
                        fmt.Println(fi40, fi40, fi40)
                    }
                    if ((255 < 93) || true) {
                        p0 = "\n\t"
                    }
                    _ = fi40
                }
            }
        }
        v46 := GenS6{F0: 4294967295, F1: 42, F2: "\n\t"}
        _ = fi5
        _ = v46
    }
    v47 := []int{647, -864}
    _ = len(v47)
    _ = p0
    _ = v47
    return p0
}
func genFn1(p0 int, p1 string, p2 int32) uint32 {
    if (-32768 != -32768) {
        v48 := GenS5{F0: 45}
        _ = v48
    } else {
        v49 := []int{p0, p0, (p0 ^ (p0 << ((18446744073709551615) & 63)))}
        p2 = p2
        _ = v49
    }
    v50 := GenS1{F0: 32767}
    var v51 string = "hello"
    var v52 uint8 = 202
    v53 := GenS7{F0: 0, F1: 4294967295, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    return 231
}

func main() {
    var v1 uint8 = 232
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v4 uint64 = 790
    fmt.Println(v4, v1)
    switch "" {
    case "a":
        if (0 > -620) {
            if (false && false) {
                v4 = v4
                v54 := v3[1]
                _ = v54
            }
            if (!((363 < (^(v4))) && ((722 < 32767) || false))) {
                if (4294967295 != 2147483647) {
                    _ = len(v3)
                    v55 := []int{661, 32768}
                    _ = v55
                }
                for fi56 := 0; fi56 < 20; fi56++ {
                    v57 := GenS3{F0: "a", F1: v1, F2: 42}
                    v4 += 596
                    _ = fi56
                    _ = v57
                }
                switch ("foo bar" + "") {
                case "hello":
                    v1 += (v1 >> ((25) & 63))
                default:
                    switch ("hello" + "foo bar") {
                    case "hello":
                        v4 = 18446744073709551615
                    case "foo bar":
                        for fi58 := 0; fi58 < 1; fi58++ {
                            fmt.Println(fi58, fi58, v4)
                            _ = fi58
                        }
                    default:
                        for fi59 := 0; fi59 < 8; fi59++ {
                            for fi60 := 0; fi60 < 14; fi60++ {
                                _ = len(v2)
                                var v61 int16 = 674
                                _ = fi60
                                _ = v61
                            }
                            _ = fi59
                        }
                    }
                }
            }
            v1 = (v1 - (^(v1)))
        }
    default:
        if (-2147483648 != -1) {
            v62 := GenS8{F0: 271, F1: 0}
            v63 := []int{-912, 9223372036854775807, -2147483648, 946}
            _ = len(v63)
            _ = v62
            _ = v63
        }
    }
    v64 := GenS4{F0: true, F1: -32768, F2: "foo bar"}
    delete(v2, "\n\t")
    v4 -= (v4 % ((v4) | 1))
    v65 := GenS7{F0: 2, F1: 0, F2: 0}
    switch 9223372036854775807 {
    case 4294967295:
        switch -1 {
        case -253:
            if (255 <= 0) {
                for fi66 := 0; fi66 < 8; fi66++ {
                    var v67 float64 = (((3.14159 * 3.14159) / float64(float32(3.14159))) - (0.0 + (1e10 - 1e10)))
                    _ = fi66
                    _ = v67
                }
                v68 := []int{2147483647, 2147483647, 42}
                for fi69 := 0; fi69 < 3; fi69++ {
                    switch fi69 {
                    case 32768:
                        fmt.Println(v1, fi69, fi69)
                    case 42:
                        if true {
                            v1 = v1
                            v2["foo bar"] = "a"
                        }
                    }
                    v2[""] = "foo bar"
                    _ = fi69
                }
                _ = v68
            } else {
                if (1e10 == (0.0 * -2.25)) {
                    var v70 int8 = 42
                    for fi71 := 0; fi71 < 18; fi71++ {
                        var v72 int64 = 42
                        _ = fi71
                        _ = v72
                    }
                    _ = v70
                } else {
                    switch 897 {
                    case 127:
                        v1 = 255
                    case -22:
                        for fi73 := 0; fi73 < 5; fi73++ {
                            switch "foo bar" {
                            case "foo bar":
                                v74 := GenS3{F0: (("" + "hello") + ("" + (("" + "") + "a"))), F1: v1, F2: 538}
                                _ = v74
                            }
                            _ = fi73
                        }
                    case -36:
                        for fi75 := 0; fi75 < 14; fi75++ {
                            for fi76 := 0; fi76 < 6; fi76++ {
                                var v77 uint16 = 255
                                fmt.Println(fi75, v4, v4)
                                _ = fi76
                                _ = v77
                            }
                            _ = fi75
                        }
                    default:
                        v4 = v4
                    }
                    v4 = 646
                }
                v78 := GenS9{F0: ""}
                _ = v78
            }
        }
    case 0:
        var v79 int64 = -1
        _ = v79
    }
    var v80 string = "a"
    var v81 int64 = -9223372036854775808
    var v82 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v80
    _ = v81
    _ = v82
}
