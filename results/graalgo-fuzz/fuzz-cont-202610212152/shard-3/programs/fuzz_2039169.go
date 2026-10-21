package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint16; F2 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 uint8; F1 uint; F2 int64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int16; F1 uint32 }
type GenS14 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint32, p2 uint32) bool {
    switch 914 {
    case -1:
        v6 := GenS1{F0: -1}
        _ = v6
    }
    var v7 uint64 = 42
    var v8 string = ("\n\t" + ("a" + ("foo bar" + "a")))
    v9 := GenS1{F0: -32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 100
    var v3 uint16 = 1
    var v4 uint = 2
    var v5 int8 = v2
    if false {
        v4 += 18446744073709551615
    } else {
        v10 := []int{9223372036854775807, 556, 2147483647, -2147483648}
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS2{F0: ("" + ""), F1: 1}
            v13 := GenS3{F0: false}
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(v4, v5, v5)
        _ = v10
    }
    v14 := []int{-2147483648, 128, 128, -9223372036854775808}
    if true {
        switch -9223372036854775808 {
        case 32768:
            v15 := GenS2{F0: "\n\t", F1: 1}
            _ = v15
        default:
            v4 = v4
        }
        var v16 int16 = 42
        _ = v16
    } else {
        v3 = 190
    }
    if (761 >= v4) {
        v17 := GenS3{F0: false}
        fmt.Println(v3, v3)
        var v18 uint16 = v3
        _ = v17
        _ = v18
    } else {
        fmt.Println(v2, v3, v2)
        fmt.Println(v5, v2)
    }
    var v19 uint16 = (^(v3))
    if (!true) {
        if (521 < 876) {
            var v20 uint8 = 42
            v1[32768] = 910
            if true {
                if len(v14) > 0 {
                    v14[0] = -2147483648
                }
                v1[-9223372036854775808] = 1
            } else {
                var v21 int8 = (v2 | v5)
                _ = v21
            }
            _ = v20
        } else {
            v19 -= uint16(v4)
            if false {
                v22 := GenS1{F0: -943}
                for fi23 := 0; fi23 < 4; fi23++ {
                    v3 = (v3 | 255)
                    switch fi23 {
                    case -1:
                        var v24 int8 = v5
                        _ = v24
                    case 42:
                        if len(v14) > 0 {
                            v14[0] = -796
                        }
                    case -9223372036854775808:
                        switch fi23 {
                        case -2147483648:
                            v25 := GenS4{F0: 65535, F1: -2147483648}
                            _ = v25
                        }
                    }
                    if false {
                        fmt.Println(v4, v19, fi23)
                        if false {
                            v2 = v5
                            for fi26 := 0; fi26 < 12; fi26++ {
                                switch "\n\t" {
                                case "foo bar":
                                    var v27 int16 = -32768
                                    _ = v27
                                }
                                v5 += 100
                                switch "foo bar" {
                                case "a":
                                    v4 = 0
                                case "":
                                    v4 = v4
                                    v28 := GenS1{F0: 536}
                                    _ = v28
                                case "foo bar":
                                    v29 := GenS5{F0: v19}
                                    _ = v29
                                default:
                                    var v30 float64 = 1e10
                                    _ = v30
                                }
                                switch ("hello" + "foo bar") {
                                case "a":
                                    switch fi26 {
                                    case 32768:
                                        var v31 uint = 857
                                        _ = v31
                                    }
                                default:
                                    v4 = v4
                                }
                                _ = fi26
                            }
                            v32 := GenS2{F0: "\n\t", F1: -1}
                            _ = v32
                        } else {
                            var v33 float32 = (-2.25 + 3.14159)
                            switch "foo bar" {
                            case "a":
                                v14 = append(v14, (fi23 + 2147483647))
                            default:
                                for fi34 := 0; fi34 < 17; fi34++ {
                                    v19 = 0
                                    v1[(^(fi34))] = fi23
                                    v35 := GenS5{F0: 389}
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            _ = v33
                        }
                    } else {
                        if len(v14) > 0 {
                            v14[0] = (^(fi23))
                        }
                        if (18446744073709551615 > 4294967295) {
                            v36 := GenS6{F0: 998, F1: -66}
                            v37 := GenS3{F0: true}
                            _ = v36
                            _ = v37
                        }
                    }
                    _ = fi23
                }
                for fi38 := 0; fi38 < 1; fi38++ {
                    switch (("a" + "hello") + "hello") {
                    case "":
                        if ((1 < 1) || false) {
                            v39 := GenS4{F0: 2147483648, F1: (fi38 ^ fi38)}
                            var v40 uint8 = 255
                            _ = v39
                            _ = v40
                        } else {
                            var v41 string = "hello"
                            var v42 string = v41
                            v43 := GenS7{F0: v19, F1: (v19 ^ v3), F2: fi38}
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                    default:
                        if (((1.5 < 1e10) && true) && false) {
                            switch fi38 {
                            case 809:
                                v14 = append(v14, fi38)
                            case 513:
                                v3 = (v3 & (v3 | v19))
                            case 42:
                                v3 = (^(v3))
                            default:
                                if (((1.5 * -2.25) < 3.14159) && false) {
                                    for fi44 := 0; fi44 < 10; fi44++ {
                                        var v45 string = "\n\t"
                                        if false {
                                            if true {
                                                var v46 uint = 0
                                                v2 += v5
                                                if true {
                                                    if false {
                                                        v47 := GenS2{F0: (("\n\t" + (v45 + (v45 + v45))) + "foo bar"), F1: 42}
                                                        var v48 string = "foo bar"
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        if (1e10 < 1.5) {
                                                            if false {
                                                                v5 -= 109
                                                                v49 := GenS8{F0: v4}
                                                                _ = v49
                                                            }
                                                        } else {
                                                            delete(v1, 1)
                                                        }
                                                        var v50 float32 = 1e10
                                                        _ = v50
                                                    }
                                                    fmt.Println(v19)
                                                } else {
                                                    v3 = (v19 & 580)
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v45)
                                                    case "\n\t":
                                                        v51 := GenS9{F0: -795, F1: (fi38 & fi38)}
                                                        _ = v51
                                                    }
                                                }
                                                _ = v46
                                            } else {
                                                v52 := GenS10{F0: 2147483647, F1: (fi38 >> (((v4 / (((v4 >> (((v4 | v4)) & 63))) | 1))) & 63))}
                                                var v53 string = "\n\t"
                                                v19 = (v3 ^ (^(v19)))
                                                _ = v52
                                                _ = v53
                                            }
                                            var v54 string = (v45 + v45)
                                            var v55 int16 = 65
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v56 := GenS6{F0: 18446744073709551615, F1: (v2 - (v2 & (v5 / ((42) | 1))))}
                                            v57 := GenS11{F0: 189, F1: (v4 | (v4 * 276)), F2: 9223372036854775807}
                                            var v58 string = "\n\t"
                                            var v59 float32 = 3.14159
                                            var v60 uint8 = 1
                                            fmt.Println(v58, v19, fi44)
                                            _ = v56
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        }
                                        v14 = append(v14, fi38)
                                        _ = fi44
                                        _ = v45
                                    }
                                    if true {
                                        switch (fi38 << (((^(v4))) & 63)) {
                                        case 32767:
                                            for fi61 := 0; fi61 < 13; fi61++ {
                                                switch "a" {
                                                case "\n\t":
                                                    v62 := GenS12{F0: v3}
                                                    _ = v62
                                                case "":
                                                    var v63 uint64 = 1
                                                    _ = v63
                                                case "foo bar":
                                                    v5 += v5
                                                }
                                                switch 128 {
                                                case 32768:
                                                    if false {
                                                        v2 += v5
                                                        var v64 uint = 42
                                                        _ = v64
                                                    }
                                                case 932:
                                                    var v65 float32 = 3.14159
                                                    _ = v65
                                                case -2147483648:
                                                    if (((false || ((v3 | v19) <= (v3 << (((^(v4))) & 63)))) || true) || (0 != 18446744073709551615)) {
                                                        v4 += (^(v4))
                                                    } else {
                                                        switch fi61 {
                                                        case -479:
                                                            if len(v14) > 0 {
                                                                v14[0] = 0
                                                            }
                                                        case 331:
                                                            v66 := GenS1{F0: 32767}
                                                            _ = v66
                                                        }
                                                    }
                                                }
                                                var v67 bool = (!false)
                                                _ = fi61
                                                _ = v67
                                            }
                                        case -245:
                                            switch (-(fi38)) {
                                            case 1:
                                                for fi68 := 0; fi68 < 13; fi68++ {
                                                    fmt.Println(v4, fi38)
                                                    _ = fi68
                                                }
                                            default:
                                                v5 = v5
                                            }
                                        default:
                                            delete(v1, (-(fi38)))
                                        }
                                    } else {
                                        for fi69 := 0; fi69 < 1; fi69++ {
                                            v70 := GenS13{F0: 42, F1: 2147483648}
                                            if ((false || (!false)) && false) {
                                                fmt.Println(v2, v4)
                                                v19 += 2
                                            }
                                            fmt.Println(v2, v5, v19)
                                            _ = fi69
                                            _ = v70
                                        }
                                    }
                                    var v71 uint32 = 1
                                    _ = v71
                                } else {
                                    v19 = (^(v19))
                                    if len(v14) > 0 {
                                        v14[0] = int(v3)
                                    }
                                }
                            }
                            v3 = v3
                        } else {
                            v5 += v2
                        }
                    }
                    if true {
                        v72 := GenS14{F0: 2}
                        var v73 uint8 = 255
                        _ = v72
                        _ = v73
                    } else {
                        v19 = (v19 >> (((v4 * 2)) & 63))
                        var v74 bool = true
                        fmt.Println(fi38)
                        _ = v74
                    }
                    _ = fi38
                }
                _ = v22
            }
            v75 := GenS10{F0: 249, F1: 127}
            _ = v75
        }
        var v76 int = 42
        _ = v76
    } else {
        var v77 uint8 = 127
        v3 += v19
        v3 = v19
        _ = v77
    }
    var v78 string = "hello"
    var v79 int32 = 0
    switch "a" {
    case "a":
        v3 -= v19
    case "\n\t":
        if len(v14) > 0 {
            v14[0] = 127
        }
    case "foo bar":
        switch ((("hello" + "\n\t") + v78) + (v78 + ("foo bar" + "a"))) {
        case "hello":
            for fi80 := 0; fi80 < 9; fi80++ {
                fmt.Println(v3)
                fmt.Println(v5)
                v14 = append(v14, (-(fi80)))
                _ = fi80
            }
        case "foo bar":
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
    }
    var v81 uint32 = 65535
    fmt.Println(v4, v5)
    v82 := GenS5{F0: v19}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v79
    _ = v81
    _ = v82
}
