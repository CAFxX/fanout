package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int; F1 int8; F2 string }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint8; F1 bool; F2 int8 }
type GenS6 struct { F0 int8; F1 bool }
func genFn0(p0 int16, p1 uint16) int32 {
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 6; fi5++ {
        switch "" {
        case "\n\t":
            for fi6 := 0; fi6 < 17; fi6++ {
                p0 += p0
                v7 := GenS1{F0: -1}
                v8 := GenS1{F0: 83}
                switch "\n\t" {
                case "foo bar":
                    if (!(18446744073709551615 < 856)) {
                        switch fi5 {
                        case 42:
                            var v9 int8 = 127
                            _ = v9
                        case 128:
                            var v10 float32 = 3.14159
                            _ = v10
                        }
                        v11 := GenS1{F0: -101}
                        _ = v11
                    } else {
                        v12 := GenS2{F0: (fi6 / ((-528) | 1)), F1: (-(fi5))}
                        switch "foo bar" {
                        case "a":
                            v13 := []int{fi5, 0, fi5, fi5}
                            _ = v13
                        default:
                            v14 := GenS3{F0: -573, F1: 100, F2: "foo bar"}
                            _ = v14
                        }
                        _ = v12
                    }
                case "":
                    if false {
                        v15 := []int{(fi6 - (fi5 * fi6)), -538, (^(fi5)), fi6}
                        p0 += p0
                        var v16 int16 = -1
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        case "a":
            var v17 uint64 = 42
            _ = v17
        case "hello":
            if false {
                if false {
                    var v18 string = ("" + ((("\n\t" + ("foo bar" + (("a" + "") + "hello"))) + "") + ("a" + "foo bar")))
                    v19 := []int{-550, (fi5 & fi5), (^(fi5)), 4294967295}
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 int8 = 127
                    v22 := []int{fi5, fi5, (fi20 >> ((583) & 63)), -9223372036854775808}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                if true {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        p1 += (^(p1))
                        var v24 float32 = (((-2.25 - float32(float64(3.14159))) / 1.5) / 0.1)
                        _ = fi23
                        _ = v24
                    }
                    v25 := []int{186, (-(fi5)), (fi5 + fi5), 1}
                    var v26 float64 = -2.25
                    _ = v25
                    _ = v26
                }
            } else {
                if (0 <= 42) {
                    if (false || true) {
                        var v27 uint32 = 1
                        _ = v27
                    }
                }
                for fi28 := 0; fi28 < 3; fi28++ {
                    p0 -= p0
                    fmt.Println(fi28)
                    switch (fi28 & fi28) {
                    case 128:
                        if true {
                            var v29 uint = 0
                            _ = v29
                        } else {
                            var v30 int64 = -82
                            v4 = v4
                            _ = v30
                        }
                    case 9223372036854775807:
                        switch (fi28 - 0) {
                        case -263:
                            var v31 int16 = 1
                            _ = v31
                        default:
                            p1 -= (p1 | (^(p1)))
                        }
                    default:
                        v32 := []int{(fi28 | (fi5 << ((4294967295) & 63)))}
                        _ = v32
                    }
                    _ = fi28
                }
            }
        }
        _ = fi5
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 int16 = -1
        if true {
            for fi35 := 0; fi35 < 13; fi35++ {
                if true {
                    if (((((0.1 - -2.25) * (1e10 - (0.1 / (-2.25 / (float32(v4) - 1e10))))) + (0.1 / ((-2.25 - (1.5 / (1.5 + 1.5))) / 1.5))) * ((3.14159 + -2.25) + -2.25)) != -2.25) {
                        var v36 bool = false
                        v37 := []int{(fi33 ^ 1), fi33, fi35}
                        fmt.Println(v36)
                        _ = v36
                        _ = v37
                    }
                }
                _ = fi35
            }
            for fi38 := 0; fi38 < 13; fi38++ {
                v34 = v34
                var v39 string = "\n\t"
                v40 := GenS1{F0: -128}
                _ = fi38
                _ = v39
                _ = v40
            }
        } else {
            switch "" {
            case "foo bar":
                var v41 uint32 = 1
                _ = v41
            case "\n\t":
                v42 := []int{fi33}
                _ = v42
            case "hello":
                fmt.Println(p1, v34, v4)
            }
            var v43 uint32 = 0
            for fi44 := 0; fi44 < 12; fi44++ {
                var v45 uint16 = 42
                var v46 uint = 651
                if ((fi33 & fi44) != (fi44 + fi33)) {
                    v47 := []int{(fi33 * 42), int(p1), fi33}
                    v47 = append(v47, fi33)
                    v4 = v4
                    _ = v47
                } else {
                    var v48 uint8 = 116
                    _ = v48
                }
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = v43
        }
        _ = fi33
        _ = v34
    }
    p0 -= int16(p1)
    _ = p0
    _ = p1
    _ = v4
    return 410
}
func genFn1(p0 uint, p1 uint, p2 string) uint16 {
    p2 = (p2 + "\n\t")
    v49 := GenS2{F0: 32768, F1: 128}
    var v50 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v49
    _ = v50
    return 624
}

func main() {
    var v1 int16 = 42
    var v2 int = 42
    v3 := make(map[string]int)
    v3["a"] = (v2 ^ (v2 + (v2 & (v2 ^ -772))))
    for fi51 := 0; fi51 < 11; fi51++ {
        switch 127 {
        case 1:
            v2 = (v2 / (((v2 - 128)) | 1))
        case 2147483647:
            for fi52 := 0; fi52 < 20; fi52++ {
                var v53 uint16 = 42
                var v54 uint = 523
                _ = fi52
                _ = v53
                _ = v54
            }
        case -9223372036854775808:
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v2)
            }
        default:
            for fi55 := 0; fi55 < 2; fi55++ {
                v2 -= (fi51 / ((int(v1)) | 1))
                v2 = fi55
                _ = fi55
            }
        }
        _ = fi51
    }
    switch (v2 / (((-(v2))) | 1)) {
    case -76:
        if true {
            var v56 int = (v2 | v2)
            _ = v56
        } else {
            var v57 uint64 = 2
            _ = v57
        }
    case 4294967295:
        for fi58 := 0; fi58 < 5; fi58++ {
            v1 -= v1
            if false {
                fmt.Println(v2, v1)
                if true {
                    for fi59 := 0; fi59 < 16; fi59++ {
                        switch (-(v2)) {
                        case 2147483647:
                            v1 -= (v1 ^ v1)
                            v3["hello"] = fi59
                            v1 = v1
                        case 128:
                            fmt.Println(v1)
                        case -9223372036854775808:
                            v2 = (v2 / ((v2) | 1))
                        }
                        var v60 int64 = 9223372036854775807
                        switch "\n\t" {
                        case "\n\t":
                            v61 := GenS2{F0: (v2 ^ (fi58 % ((32767) | 1))), F1: (v2 + 4294967295)}
                            if (fi58 != fi59) {
                                var v62 uint32 = 0
                                v2 -= fi59
                                _ = v62
                            }
                            var v63 uint32 = 502
                            fmt.Println(v2, v63, v2)
                            fmt.Println(v60, v60)
                            _ = v61
                            _ = v63
                        case "":
                            if false {
                                var v64 int64 = 553
                                switch "foo bar" {
                                case "foo bar":
                                    v65 := []int{fi58, (v2 << ((65535) & 63)), fi58}
                                    _ = v65
                                }
                                _ = v64
                            }
                        case "hello":
                            if (-577 == -1) {
                                v1 = 42
                                v66 := GenS4{F0: 819, F1: 1}
                                switch v2 {
                                case -336:
                                    v60 = -1
                                case 4294967295:
                                    v60 = v60
                                default:
                                    switch "" {
                                    case "a":
                                        fmt.Println(v60, fi59)
                                    default:
                                        v67 := GenS5{F0: 0, F1: true, F2: 49}
                                        _ = v67
                                    }
                                }
                                var v68 int8 = -7
                                _ = v66
                                _ = v68
                            } else {
                                var v69 uint = 4294967295
                                var v70 bool = true
                                _ = v69
                                _ = v70
                            }
                        default:
                            if (!(true && (18446744073709551615 < 65535))) {
                                v71 := GenS6{F0: -128, F1: (42 > -1)}
                                _ = v71
                            }
                        }
                        v72 := GenS6{F0: 0, F1: (!false)}
                        _ = fi59
                        _ = v60
                        _ = v72
                    }
                    fmt.Println(v2, v1, fi58)
                } else {
                    switch v2 {
                    case -755:
                        switch ("" + (("\n\t" + "") + ("a" + "hello"))) {
                        case "foo bar":
                            if false {
                                v1 = (v1 / ((32767) | 1))
                            } else {
                                fmt.Println(v1)
                                var v73 uint16 = genFn1(1, 2, "")
                                _ = v73
                            }
                        }
                    case 128:
                        v1 = (v1 - (v1 | v1))
                    case -1:
                        for fi74 := 0; fi74 < 20; fi74++ {
                            var v75 int = 127
                            if (2147483648 == 775) {
                                v2 -= (fi58 - (v2 / ((0) | 1)))
                            }
                            delete(v3, "foo bar")
                            _ = fi74
                            _ = v75
                        }
                    default:
                        var v76 int8 = 0
                        _ = v76
                    }
                }
            }
            _ = fi58
        }
    case 9223372036854775807:
        _ = len(v3)
    default:
        var v77 string = "\n\t"
        _ = v77
    }
    v1 = 42
    v78 := GenS4{F0: 0, F1: 4294967295}
    v79 := []int{}
    v1 -= v1
    v1 = v1
    var v80 int32 = 0
    var v81 uint8 = 152
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v79:", len(v79), v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v78
    _ = v79
    _ = v80
    _ = v81
}
