package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int32; F1 uint8; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int8; F2 bool }
type GenS5 struct { F0 bool; F1 int16; F2 uint16 }
type GenS6 struct { F0 bool; F1 int64; F2 int64 }
func genFn0(p0 int32, p1 uint, p2 int) bool {
    v5 := []int{(p2 ^ p2)}
    var v6 bool = true
    for fi7 := 0; fi7 < 17; fi7++ {
        p2 = p2
        if true {
            var v8 string = ("\n\t" + "\n\t")
            switch fi7 {
            case 2147483647:
                var v9 string = ("foo bar" + v8)
                _ = v9
            case -9223372036854775808:
                var v10 uint32 = 4294967295
                _ = v10
            case 901:
                v11 := GenS1{F0: uint8(p1), F1: v8}
                _ = v11
            default:
                switch (p2 + (fi7 + (fi7 - (p2 & (p2 / ((127) | 1)))))) {
                case -780:
                    v12 := GenS1{F0: 0, F1: v8}
                    _ = v12
                case -755:
                    if ((("" + v8) + (v8 + v8)) <= v8) {
                        v5 = append(v5, (fi7 / (((fi7 / (((fi7 | (fi7 * (fi7 ^ int(p0))))) | 1))) | 1)))
                        v8 = ("" + v8)
                        var v13 bool = (42 <= -22)
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 47, F1: (("foo bar" + v8) + v8)}
                        fmt.Println(p2, p1)
                        fmt.Println(p2, v8, p0)
                        _ = v14
                    }
                case 2147483647:
                    var v15 uint64 = 18446744073709551615
                    _ = v15
                default:
                    var v16 int16 = 0
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 5; fi17++ {
                switch fi17 {
                case 2147483647:
                    if true {
                        p1 += (p1 + p1)
                    } else {
                        var v18 uint = p1
                        var v19 int = p2
                        p2 -= 127
                        _ = v18
                        _ = v19
                    }
                case 32767:
                    switch v8 {
                    case "\n\t":
                        if (593 != -1) {
                            var v20 uint16 = 934
                            _ = v20
                        } else {
                            p0 = -2147483648
                            if (!true) {
                                switch "\n\t" {
                                case "":
                                    if v6 {
                                        var v21 int32 = 434
                                        var v22 uint = 4294967295
                                        v5 = append(v5, fi17)
                                        fmt.Println(fi17)
                                        if (v6 || (fi17 > 2147483647)) {
                                            if v6 {
                                                var v23 int = int(v21)
                                                if (v6 && (v21 <= v21)) {
                                                    fmt.Println(v6)
                                                    v24 := GenS2{F0: (p0 ^ v21), F1: 42, F2: p0}
                                                    _ = v24
                                                } else {
                                                    v25 := GenS3{F0: ""}
                                                    switch v23 {
                                                    case 499:
                                                        v26 := GenS3{F0: "foo bar"}
                                                        _ = v26
                                                    }
                                                    switch fi7 {
                                                    case 810:
                                                        switch "a" {
                                                        case "foo bar":
                                                            if v6 {
                                                                var v27 bool = v6
                                                                p2 -= (p2 >> ((v22) & 63))
                                                                _ = v27
                                                            }
                                                        case "\n\t":
                                                            if v6 {
                                                                p1 -= p1
                                                                v8 = ("hello" + "\n\t")
                                                                switch -523 {
                                                                case -9223372036854775808:
                                                                    v22 += 42
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(p0, fi17)
                                                    }
                                                    _ = v25
                                                }
                                                v21 -= (p0 * (v21 << (((^(p1))) & 63)))
                                                _ = v23
                                            }
                                        }
                                        _ = v21
                                        _ = v22
                                    } else {
                                        if v6 {
                                            v28 := GenS3{F0: v8}
                                            if v6 {
                                                p1 += p1
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = (p2 ^ (fi17 ^ (fi17 << (((p1 * 42)) & 63))))
                                            }
                                            _ = v28
                                        }
                                        v5 = append(v5, (fi17 + fi17))
                                        var v29 int = fi7
                                        _ = v29
                                    }
                                default:
                                    if v6 {
                                        switch "hello" {
                                        case "foo bar":
                                            if (18446744073709551615 == 1) {
                                                if len(v5) > 0 {
                                                    v5[0] = (fi17 >> (((^(p1))) & 63))
                                                }
                                            } else {
                                                var v30 string = "\n\t"
                                                _ = v30
                                            }
                                        case "":
                                            switch fi17 {
                                            case 128:
                                                var v31 uint64 = 303
                                                _ = v31
                                            default:
                                                fmt.Println(p2, fi7, v6)
                                            }
                                        case "a":
                                            v32 := GenS4{F0: 9223372036854775807, F1: 127, F2: v6}
                                            _ = v32
                                        default:
                                            var v33 uint16 = 32767
                                            _ = v33
                                        }
                                    }
                                }
                                switch (fi17 + -9223372036854775808) {
                                case 127:
                                    if (2 != 32767) {
                                        fmt.Println(v6, fi17)
                                        p0 -= 0
                                    }
                                case 9223372036854775807:
                                    var v34 string = v8
                                    _ = v34
                                }
                            }
                            v35 := GenS4{F0: -935, F1: -76, F2: v6}
                            _ = v35
                        }
                    }
                    if false {
                        _ = len(v5)
                    }
                    v8 = (v8 + "\n\t")
                case -2147483648:
                    var v36 uint16 = 42
                    _ = v36
                default:
                    var v37 int = -743
                    _ = v37
                }
                fmt.Println(fi17)
                if (((956 != 1) && false) && (!(35 > 318))) {
                    switch ("hello" + ("\n\t" + "foo bar")) {
                    case "hello":
                        if len(v5) > 0 {
                            v5[0] = 9223372036854775807
                        }
                    }
                }
                _ = fi17
            }
            _ = v8
        } else {
            p0 -= p0
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    var v3 uint16 = 322
    v4 := make(map[string]int)
    switch 32768 {
    case 42:
        var v38 uint = 42
        _ = v38
    case -9223372036854775808:
        var v39 int16 = 42
        _ = v39
    case 9223372036854775807:
        if ((1 == 1) && false) {
            _ = len(v2)
        } else {
            var v40 int64 = -9223372036854775808
            var v41 float64 = 0.0
            _ = v40
            _ = v41
        }
    default:
        for fi42 := 0; fi42 < 3; fi42++ {
            v1[fi42] = (128 != 127)
            _ = fi42
        }
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        fmt.Println(v3, v3)
        v3 = v3
        for fi44 := 0; fi44 < 12; fi44++ {
            if false {
                var v45 uint64 = 1
                switch (((("hello" + "\n\t") + ("a" + (("a" + ("hello" + "")) + "a"))) + ((("\n\t" + "\n\t") + "foo bar") + "hello")) + ("" + ("a" + "hello"))) {
                case "foo bar":
                    v3 -= (v3 | 65535)
                }
                _ = v45
            } else {
                if false {
                    var v46 float32 = (-2.25 - 1.5)
                    v47 := GenS3{F0: "hello"}
                    _ = v46
                    _ = v47
                }
            }
            if (!false) {
                v3 = (v3 % ((v3) | 1))
            } else {
                v3 = v3
                fmt.Println(fi43)
            }
            _ = fi44
        }
        _ = fi43
    }
    fmt.Println(v3)
    v3 += 0
    if true {
        if true {
            v3 = (v3 & (v3 | (v3 | (v3 | v3))))
            v4[((((("" + "\n\t") + "") + "a") + "\n\t") + "\n\t")] = 1
            if true {
                for fi48 := 0; fi48 < 13; fi48++ {
                    v3 += 55
                    _ = fi48
                }
                fmt.Println(v3)
                v3 -= 830
            } else {
                var v49 int16 = -965
                fmt.Println(v49)
                _ = v49
            }
        }
        var v50 int = 32767
        _ = v50
    } else {
        switch 2147483647 {
        case 789:
            if false {
                if (18446744073709551615 < 42) {
                    v3 -= (v3 | (v3 / (((v3 / (((v3 * v3)) | 1))) | 1)))
                    var v51 string = ""
                    var v52 uint8 = 127
                    _ = v51
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 8; fi53++ {
                        switch (fi53 * fi53) {
                        case 9223372036854775807:
                            v54 := GenS5{F0: (!(true || (float32(0.0) != 1.5))), F1: 32767, F2: v3}
                            _ = v54
                        }
                        for fi55 := 0; fi55 < 18; fi55++ {
                            var v56 bool = ((true && (-2.25 >= 0.1)) || true)
                            v3 = (v3 * v3)
                            _ = fi55
                            _ = v56
                        }
                        _ = fi53
                    }
                }
                v57 := GenS2{F0: -1, F1: 255, F2: 0}
                for fi58 := 0; fi58 < 13; fi58++ {
                    if ((1.5 + float32(0.1)) != 0.1) {
                        if ((0.0 > 1e10) && true) {
                            v59 := GenS1{F0: 128, F1: "hello"}
                            _ = v59
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v3 = (^(v3))
                        default:
                            var v60 uint16 = v3
                            _ = v60
                        }
                        switch (-(fi58)) {
                        case -1:
                            for fi61 := 0; fi61 < 19; fi61++ {
                                if true {
                                    v62 := GenS4{F0: 1, F1: 127, F2: (4294967295 > 0)}
                                    _ = v62
                                } else {
                                    fmt.Println(fi58, v3)
                                    v3 = (v3 & 32767)
                                    var v63 bool = false
                                    _ = v63
                                }
                                var v64 bool = (!true)
                                var v65 uint16 = v3
                                _ = fi61
                                _ = v64
                                _ = v65
                            }
                        }
                    }
                    if false {
                        if (true && false) {
                            for fi66 := 0; fi66 < 16; fi66++ {
                                v3 += 208
                                fmt.Println(v3)
                                _ = fi66
                            }
                            _ = len(v1)
                            var v67 int = (fi58 >> ((65535) & 63))
                            _ = v67
                        }
                        _ = len(v4)
                        v68 := GenS6{F0: false, F1: 9223372036854775807, F2: 42}
                        _ = v68
                    } else {
                        fmt.Println(v3)
                        v3 -= (v3 & v3)
                        v3 -= 65535
                    }
                    for fi69 := 0; fi69 < 12; fi69++ {
                        fmt.Println(v3, fi69)
                        v3 += v3
                        v3 += (v3 / ((v3) | 1))
                        switch "foo bar" {
                        case "\n\t":
                            var v70 int64 = 42
                            _ = v70
                        case "foo bar":
                            if ((fi58 ^ 128) > (fi69 ^ (fi69 * (fi69 >> ((255) & 63))))) {
                                fmt.Println(fi69, fi58)
                                v3 -= (^(v3))
                                var v71 uint16 = v3
                                _ = v71
                            }
                        case "":
                            if len(v2) > 0 {
                                v2[0] = fi58
                            }
                        }
                        _ = fi69
                    }
                    _ = fi58
                }
                _ = v57
            }
        default:
            for fi72 := 0; fi72 < 18; fi72++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi73 := 0; fi73 < 19; fi73++ {
                    v3 += (v3 << ((784) & 63))
                    delete(v1, 127)
                    _ = fi73
                }
                _ = fi72
            }
        }
        v74 := GenS6{F0: true, F1: -1, F2: -1}
        if false {
            delete(v4, "a")
            switch 127 {
            case -2147483648:
                _ = len(v2)
            case -601:
                _ = len(v1)
            case -240:
                v3 += 0
            default:
                switch "\n\t" {
                case "a":
                    v75 := GenS3{F0: "a"}
                    _ = v75
                case "hello":
                    fmt.Println(v3)
                }
            }
            for fi76 := 0; fi76 < 3; fi76++ {
                if (255 < 600) {
                    var v77 int = -9223372036854775808
                    _ = v77
                }
                fmt.Println(fi76, fi76)
                _ = fi76
            }
        }
        _ = v74
    }
    v3 = v3
    fmt.Println(v3)
    fmt.Println(v3)
    var v78 uint64 = 1
    delete(v1, 9223372036854775807)
    for fi79 := 0; fi79 < 12; fi79++ {
        var v80 bool = true
        _ = fi79
        _ = v80
    }
    var v81 uint32 = 2
    v81 -= (v81 | (^(v81)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v81
}
