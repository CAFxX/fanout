package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint16) int8 {
    v4 := GenS1{F0: -9223372036854775808}
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS2{F0: 163, F1: 127}
            v8 := GenS2{F0: 2, F1: 128}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return p0
}
func genFn1(p0 int16, p1 int, p2 uint16) bool {
    p0 = p0
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 string = ""
        p0 -= p0
        v11 := []int{42}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(fi13, p2)
            var v14 string = ("" + "hello")
            var v15 uint64 = 2
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    var v16 int32 = -1
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return true
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := []string{"hello"}
    var v17 uint16 = 42
    v18 := GenS2{F0: 100, F1: 1}
    v19 := GenS3{F0: 2147483647}
    v17 += (v17 | v17)
    if true {
        v20 := GenS3{F0: 2147483648}
        v21 := GenS4{F0: (!(2 < 0))}
        var v22 int64 = 811
        _ = v20
        _ = v21
        _ = v22
    } else {
        if (119 == 2) {
            fmt.Println(v17)
            v17 = 666
            for fi23 := 0; fi23 < 2; fi23++ {
                if (((float32(1e10) > (1.5 / 3.14159)) || false) || (!false)) {
                    switch "" {
                    case "":
                        switch "\n\t" {
                        case "\n\t":
                            _ = len(v3)
                        }
                    }
                } else {
                    v24 := GenS5{F0: 2147483647, F1: 2147483647}
                    _ = v24
                }
                _ = fi23
            }
        }
        if true {
            v17 += 65535
        } else {
            var v25 int64 = -1
            _ = v25
        }
    }
    var v26 int64 = -122
    if false {
        v17 = 0
        v26 = v26
        v17 -= (v17 & 32767)
    } else {
        var v27 bool = true
        switch 128 {
        case 32767:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            var v28 bool = v27
            _ = v28
        }
        var v29 uint32 = 2
        _ = v27
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        switch fi30 {
        case -729:
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        case -2147483648:
            if ((fi30 << ((138) & 63)) > (-(fi30))) {
                switch (fi30 - fi30) {
                case -158:
                    v17 = (v17 >> ((2) & 63))
                case 32767:
                    var v31 string = ("\n\t" + "hello")
                    _ = v31
                case -1:
                    var v32 bool = (((26 >= 42) && true) || false)
                    _ = v32
                default:
                    if (false && (42 <= 2147483647)) {
                        v33 := GenS1{F0: 525}
                        v17 -= v17
                        for fi34 := 0; fi34 < 9; fi34++ {
                            v35 := GenS5{F0: 76, F1: 2}
                            v26 -= (v26 * v26)
                            _ = fi34
                            _ = v35
                        }
                        _ = v33
                    } else {
                        v26 += -9223372036854775808
                    }
                }
                v17 -= v17
                fmt.Println(v26, fi30)
            } else {
                var v36 int32 = -2147483648
                _ = v36
            }
        case 2147483647:
            v37 := GenS4{F0: true}
            _ = v37
        default:
            var v38 float32 = float32((3.14159 / 3.14159))
            _ = v38
        }
        var v39 int8 = genFn0(100, 940)
        v39 = (v39 | (v39 << ((42) & 63)))
        _ = fi30
        _ = v39
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        if false {
            if len(v3) > 0 {
                v3[0] = ""
            }
            if ((!false) && (1000 <= int16(v17))) {
                v17 -= 66
            } else {
                if false {
                    var v41 int64 = int64(v17)
                    _ = len(v1)
                    _ = v41
                } else {
                    v42 := GenS6{F0: 1, F1: (^(v17))}
                    v17 += 65535
                    v17 -= (v17 + (v17 * 32767))
                    _ = v42
                }
                var v43 int64 = v26
                fmt.Println(v26, v43)
                _ = v43
            }
            if (((((1e10 + (0.0 - 1e10)) * 0.0) * 3.14159) / 0.1) >= (1e10 * (1.5 + 0.0))) {
                for fi44 := 0; fi44 < 20; fi44++ {
                    if (false || false) {
                        v26 = (v26 + v26)
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v26)
                        default:
                            fmt.Println(v17, v26)
                            switch fi40 {
                            case -242:
                                fmt.Println(v26, fi44)
                            case -768:
                                var v45 int32 = -2147483648
                                _ = v45
                            case -1:
                                switch (("\n\t" + "\n\t") + "a") {
                                case "\n\t":
                                    v26 = (v26 ^ (v26 ^ v26))
                                case "a":
                                    v46 := GenS4{F0: ((!((-255 >= 42) && true)) || (2147483648 == 2))}
                                    _ = v46
                                case "foo bar":
                                    if (true || (!(255 >= 439))) {
                                        switch fi44 {
                                        case 862:
                                            switch fi44 {
                                            case 601:
                                                switch ("hello" + "hello") {
                                                case "\n\t":
                                                    delete(v1, "")
                                                case "foo bar":
                                                    var v47 uint8 = 2
                                                    _ = v47
                                                case "a":
                                                    fmt.Println(v17, v26, fi44)
                                                default:
                                                    if true {
                                                        fmt.Println(fi40, fi44, fi40)
                                                        var v48 bool = (!(false && true))
                                                        _ = v48
                                                    }
                                                }
                                            case -2147483648:
                                                v26 += 829
                                            case -507:
                                                if false {
                                                    if false {
                                                        v49 := v2[(fi44 ^ (-(fi44)))]
                                                        if false {
                                                            fmt.Println(fi44)
                                                            v49 = v49
                                                            v50 := GenS5{F0: -2147483648, F1: 522}
                                                            var v51 uint = 65535
                                                            _ = v50
                                                            _ = v51
                                                        } else {
                                                            fmt.Println(fi44)
                                                        }
                                                        _ = v49
                                                    }
                                                    fmt.Println(v26, v17, fi40)
                                                } else {
                                                    fmt.Println(fi40, v17, fi44)
                                                    v17 += v17
                                                    fmt.Println(v17)
                                                }
                                            default:
                                                switch "foo bar" {
                                                case "":
                                                    var v52 float64 = 1e10
                                                    _ = v52
                                                }
                                            }
                                        case 928:
                                            var v53 int16 = 799
                                            _ = v53
                                        }
                                        var v54 int16 = -1
                                        if true {
                                            var v55 uint16 = v17
                                            _ = v55
                                        } else {
                                            var v56 int32 = -550
                                            v57 := GenS3{F0: 2}
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = v54
                                    } else {
                                        if ((!(!(42 != 18446744073709551615))) && (!true)) {
                                            var v58 float64 = float64(1.5)
                                            _ = v58
                                        }
                                    }
                                default:
                                    var v59 int8 = -128
                                    _ = v59
                                }
                            }
                        }
                        if (!(true || true)) {
                            var v60 float64 = ((((0.1 * float64(-2.25)) - 0.0) / -2.25) / ((-2.25 - 1e10) / float64(3.14159)))
                            v61 := GenS7{F0: 1, F1: 0, F2: 189}
                            v62 := GenS7{F0: 0, F1: 1000, F2: 223}
                            _ = v60
                            _ = v61
                            _ = v62
                        } else {
                            _ = len(v1)
                            if (!true) {
                                var v63 uint64 = 18446744073709551615
                                fmt.Println(fi40, fi44)
                                _ = v63
                            }
                            if false {
                                var v64 int16 = 891
                                fmt.Println(fi44, fi44)
                                var v65 float32 = (1e10 * (1.5 + 0.0))
                                if false {
                                    fmt.Println(fi44)
                                }
                                if true {
                                    _ = len(v2)
                                    v66 := GenS1{F0: fi44}
                                    _ = v66
                                } else {
                                    var v67 int64 = 9223372036854775807
                                    var v68 uint64 = 404
                                    _ = v67
                                    _ = v68
                                }
                                var v69 int64 = -9223372036854775808
                                v26 = (-(v26))
                                _ = v64
                                _ = v65
                                _ = v69
                            }
                            var v70 int16 = -1
                            _ = v70
                        }
                    } else {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v26, v26)
                        case "\n\t":
                            var v71 int64 = (v26 & (v26 >> ((0) & 63)))
                            _ = v71
                        }
                        switch (fi40 ^ (fi40 ^ fi40)) {
                        case 4294967295:
                            switch fi40 {
                            case -1:
                                _ = len(v3)
                                v72 := GenS6{F0: (v17 >> ((0) & 63)), F1: v17}
                                _ = v72
                            }
                        case -1:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            var v73 uint64 = 42
                            _ = v73
                        }
                        v74 := v2[(fi40 * (-(fi40)))]
                        if false {
                            if (1e10 == ((0.0 + 1.5) / float64(0.1))) {
                                v26 = 28
                                fmt.Println(fi44, v74)
                            } else {
                                var v75 float64 = 1e10
                                _ = len(v2)
                                delete(v2, (fi40 & 148))
                                _ = v75
                            }
                            fmt.Println(v26)
                            v74 = v74
                        }
                        _ = len(v2)
                        _ = v74
                    }
                    var v76 int = (fi44 % ((fi44) | 1))
                    _ = fi44
                    _ = v76
                }
            } else {
                var v77 string = ("foo bar" + "\n\t")
                delete(v1, "\n\t")
                v78 := GenS2{F0: 255, F1: 229}
                _ = v77
                _ = v78
            }
        } else {
            var v79 int8 = -41
            fmt.Println(v26, v17)
            v79 = 1
            _ = v79
        }
        v80 := GenS8{F0: 1}
        _ = fi40
        _ = v80
    }
    for fi81 := 0; fi81 < 16; fi81++ {
        var v82 uint8 = 42
        _ = fi81
        _ = v82
    }
    v26 = v26
    if false {
        var v83 int8 = -127
        v17 = 0
        if ("foo bar" != "a") {
            v17 = v17
        }
        _ = v83
    } else {
        for fi84 := 0; fi84 < 20; fi84++ {
            for fi85 := 0; fi85 < 16; fi85++ {
                fmt.Println(fi84, fi85, fi84)
                var v86 uint = 255
                var v87 bool = true
                var v88 string = "a"
                var v89 uint32 = 42
                _ = fi85
                _ = v86
                _ = v87
                _ = v88
                _ = v89
            }
            _ = fi84
        }
        switch -496 {
        case 42:
            var v90 uint16 = v17
            _ = v90
        case -9223372036854775808:
            v26 = v26
        default:
            fmt.Println(v17, v26)
        }
    }
    var v91 int64 = (v26 | (v26 << ((4294967295) & 63)))
    v92 := GenS8{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v91
    _ = v92
}
