package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
func genFn0() uint8 {
    if (2 >= 0) {
        v6 := []int{128}
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(fi7)
            _ = fi8
        }
        switch "hello" {
        case "a":
            if (42 >= 4294967295) {
                var v9 uint32 = 65535
                _ = v9
            } else {
                v10 := []int{32768, (fi7 | 4294967295), (-(fi7))}
                var v11 uint = 2
                _ = v10
                _ = v11
            }
        case "foo bar":
            var v12 string = ""
            _ = v12
        case "hello":
            var v13 float32 = (-2.25 - (float32((0.1 * ((0.0 - 1e10) + ((float64(3.14159) + 1e10) * (0.0 / -2.25))))) / 1e10))
            _ = v13
        default:
            v14 := []int{fi7, 42, (^(fi7))}
            _ = v14
        }
        _ = fi7
    }
    v15 := GenS1{F0: -1}
    for fi16 := 0; fi16 < 13; fi16++ {
        switch fi16 {
        case 127:
            v17 := []int{(-(fi16)), 1}
            _ = v17
        case 2147483647:
            var v18 bool = false
            _ = v18
        case 42:
            if (true && (42 >= 32767)) {
                fmt.Println(fi16)
            }
        }
        var v19 string = (("hello" + "foo bar") + "a")
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int16 = 894
        if true {
            v21 = v21
        } else {
            v22 := []int{(fi20 * (fi20 ^ fi20)), (fi20 - fi20)}
            var v23 bool = (1 <= 0)
            _ = v22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    _ = v15
    return 42
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    var v3 uint8 = 255
    var v4 int32 = int32(v3)
    v5 := []int{882, -356, 127, -22}
    v24 := GenS2{F0: -404, F1: 42, F2: 591}
    var v25 int = -840
    var v26 uint32 = 42
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 bool = true
        v26 += uint32(v3)
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: 640, F1: v26, F2: 100}
    if true {
        if (!(true && false)) {
            if ("hello" < ("\n\t" + "a")) {
                v26 = v26
                switch "" {
                case "hello":
                    var v30 string = "a"
                    _ = v30
                case "\n\t":
                    v31 := GenS2{F0: int64(v3), F1: v26, F2: 255}
                    _ = v31
                default:
                    v5 = append(v5, (v25 * v25))
                }
                fmt.Println(v3)
            } else {
                v25 -= v25
                v32 := GenS1{F0: (v4 << ((65535) & 63))}
                _ = v32
            }
            var v33 uint16 = 1
            if (true || (true && (v4 <= v4))) {
                v1 += 1e10
                var v34 int16 = -930
                _ = v34
            } else {
                if (149 <= v26) {
                    for fi35 := 0; fi35 < 13; fi35++ {
                        var v36 float32 = 0.1
                        fmt.Println(v4)
                        _ = fi35
                        _ = v36
                    }
                    var v37 uint32 = v26
                    switch (v25 ^ (v25 >> ((1) & 63))) {
                    case -962:
                        v3 += 171
                    case 2147483647:
                        v26 += 871
                    case 835:
                        var v38 uint16 = (^(v33))
                        _ = v38
                    }
                    _ = v37
                } else {
                    for fi39 := 0; fi39 < 15; fi39++ {
                        switch (fi39 - 128) {
                        case 128:
                            if (!(!(false || (true && ((^(v26)) == (v26 / ((v26) | 1))))))) {
                                if (true && (1 < 42)) {
                                    v1 += 0.0
                                } else {
                                    for fi40 := 0; fi40 < 8; fi40++ {
                                        switch "" {
                                        case "":
                                            v41 := GenS3{F0: 1, F1: 18446744073709551615, F2: (fi40 * (fi39 - (fi40 ^ (fi40 * 32768))))}
                                            _ = v41
                                        }
                                        _ = len(v5)
                                        if (false && true) {
                                            v33 += (^(v33))
                                            v26 -= (v26 % ((v26) | 1))
                                        }
                                        _ = fi40
                                    }
                                }
                                fmt.Println(fi39)
                            } else {
                                v42 := GenS3{F0: 643, F1: 42, F2: (-(v25))}
                                _ = v42
                            }
                        case 2147483647:
                            var v43 int8 = 111
                            _ = v43
                        case 127:
                            var v44 bool = true
                            _ = v44
                        default:
                            v1 = 0.1
                        }
                        _ = fi39
                    }
                }
            }
            _ = v33
        } else {
            fmt.Println(v4, v26)
            if (v1 >= v1) {
                var v45 int32 = (v4 ^ (v4 - (v4 | 2147483647)))
                switch "" {
                case "hello":
                    for fi46 := 0; fi46 < 15; fi46++ {
                        if false {
                            v1 = 3.14159
                            switch ("foo bar" + (("" + "") + (("foo bar" + ("foo bar" + "\n\t")) + ("" + ("hello" + "a"))))) {
                            case "hello":
                                var v47 uint16 = 42
                                _ = v47
                            case "":
                                v26 += 1
                            default:
                                var v48 uint8 = v3
                                _ = v48
                            }
                            if false {
                                for fi49 := 0; fi49 < 8; fi49++ {
                                    v45 -= (v4 << ((18446744073709551615) & 63))
                                    fmt.Println(v26)
                                    var v50 uint8 = (v3 << ((4294967295) & 63))
                                    _ = fi49
                                    _ = v50
                                }
                            } else {
                                v4 = (v4 / (((v4 ^ v4)) | 1))
                            }
                        }
                        for fi51 := 0; fi51 < 15; fi51++ {
                            v3 += (v3 - (v3 / ((genFn0()) | 1)))
                            var v52 float32 = -2.25
                            _ = fi51
                            _ = v52
                        }
                        _ = fi46
                    }
                case "foo bar":
                    v3 = genFn0()
                default:
                    var v53 uint8 = (v3 >> ((413) & 63))
                    _ = v53
                }
                v45 = (^(v4))
                _ = v45
            }
            v5 = append(v5, (v25 - 456))
        }
    } else {
        v1 += float64(0.0)
        switch "a" {
        case "":
            v4 -= (v4 & -511)
        case "hello":
            if false {
                v54 := GenS1{F0: v4}
                if (-2.25 < 3.14159) {
                    fmt.Println(v1, v25, v26)
                    fmt.Println(v26, v1, v1)
                } else {
                    var v55 float32 = 1.5
                    _ = v55
                }
                _ = v54
            } else {
                v26 += v26
            }
        default:
            switch "\n\t" {
            case "\n\t":
                var v56 uint8 = 228
                _ = v56
            case "hello":
                v25 -= (v25 >> ((42) & 63))
            default:
                v25 = (v25 | 2147483647)
            }
        }
        var v57 int32 = -2147483648
        _ = v57
    }
    if false {
        fmt.Println(v3)
        v4 -= (v4 << ((849) & 63))
    }
    if true {
        v3 = (v3 & genFn0())
        if ("foo bar" > "a") {
            _ = len(v5)
            v3 += (v3 << ((18446744073709551615) & 63))
            fmt.Println(v1)
        } else {
            v58 := GenS2{F0: -1, F1: (v26 * v26), F2: 4294967295}
            v1 = v1
            for fi59 := 0; fi59 < 7; fi59++ {
                var v60 uint32 = (^(v26))
                v25 = fi59
                if true {
                    v61 := GenS2{F0: 0, F1: v60, F2: 907}
                    if true {
                        v4 += (v4 >> ((42) & 63))
                        v3 = v3
                        v62 := GenS1{F0: (v4 - v4)}
                        _ = v62
                    } else {
                        fmt.Println(v1, v25)
                    }
                    _ = v61
                } else {
                    v63 := GenS3{F0: 934, F1: 42, F2: -9223372036854775808}
                    if ((-(v4)) >= v4) {
                        switch "hello" {
                        case "\n\t":
                            var v64 bool = ((("\n\t" + "foo bar") + ("\n\t" + "foo bar")) < "hello")
                            _ = v64
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = -9223372036854775808
                            }
                        default:
                            var v65 int = v25
                            _ = v65
                        }
                        switch "foo bar" {
                        case "hello":
                            v66 := GenS4{F0: (^(fi59)), F1: ("foo bar" + ("" + "hello")), F2: 65535}
                            _ = v66
                        }
                        fmt.Println(v26, v25, v3)
                    } else {
                        v67 := GenS4{F0: 4294967295, F1: "foo bar", F2: 713}
                        _ = v67
                    }
                    _ = v63
                }
                _ = fi59
                _ = v60
            }
            _ = v58
        }
    } else {
        switch (v25 * v25) {
        case 4294967295:
            var v68 string = ("foo bar" + "\n\t")
            _ = v68
        default:
            if ((^(v3)) == v3) {
                var v69 int64 = -460
                var v70 uint8 = 42
                if (42 == v69) {
                    v25 += v25
                    if false {
                        switch "foo bar" {
                        case "a":
                            v4 += int32(v69)
                        case "foo bar":
                            var v71 uint16 = 2
                            _ = v71
                        case "hello":
                            v69 += v69
                        }
                    }
                    v72 := GenS1{F0: v4}
                    _ = v72
                }
                _ = v69
                _ = v70
            } else {
                var v73 int64 = 0
                if ((float64(((0.0 * 1.5) + 0.0)) - 0.0) > (v1 - v1)) {
                    var v74 int8 = 100
                    switch 127 {
                    case -1:
                        var v75 uint32 = (v26 * 1)
                        _ = v75
                    case -2147483648:
                        v73 = (v73 >> ((0) & 63))
                    }
                    _ = v74
                }
                switch "" {
                case "a":
                    v76 := GenS3{F0: 0, F1: 885, F2: v25}
                    _ = v76
                case "":
                    var v77 string = (("hello" + ("a" + ("foo bar" + ("foo bar" + ("hello" + ((((("hello" + "a") + ("a" + "\n\t")) + "hello") + "hello") + "a")))))) + "")
                    _ = v77
                case "hello":
                    v1 = ((((v1 / v1) + 1e10) / float64(3.14159)) * 0.1)
                default:
                    if true {
                        v3 -= (v3 << ((266) & 63))
                    } else {
                        fmt.Println(v26, v4, v4)
                        switch 127 {
                        case 757:
                            if ((^(v26)) > v26) {
                                v78 := GenS4{F0: v25, F1: ("" + ""), F2: 891}
                                _ = v78
                            }
                        case -9223372036854775808:
                            switch v25 {
                            case 9223372036854775807:
                                v3 -= 127
                            }
                        default:
                            var v79 bool = false
                            _ = v79
                        }
                        if len(v5) > 0 {
                            v5[0] = 32767
                        }
                    }
                }
                _ = v73
            }
        }
    }
    if false {
        v80 := GenS2{F0: 42, F1: v26, F2: 515}
        var v81 uint8 = 127
        _ = v80
        _ = v81
    } else {
        delete(v2, 9223372036854775807)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
