package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
type GenS3 struct { F0 uint32; F1 int64; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 bool }
func genFn0(p0 uint) uint8 {
    switch "foo bar" {
    case "":
        p0 = p0
    default:
        if false {
            for fi5 := 0; fi5 < 19; fi5++ {
                if true {
                    for fi6 := 0; fi6 < 3; fi6++ {
                        if true {
                            if (42 != int32(fi5)) {
                                if ((p0 << ((4294967295) & 63)) < 1) {
                                    v7 := GenS1{F0: -9, F1: (p0 / ((42) | 1)), F2: (p0 - 0)}
                                    var v8 uint8 = 1
                                    _ = v7
                                    _ = v8
                                } else {
                                    v9 := []int{fi6, 1, -2147483648, 32767}
                                    v9 = append(v9, (-(fi5)))
                                    _ = v9
                                }
                                if (0 <= 18446744073709551615) {
                                    var v10 int64 = -9223372036854775808
                                    _ = v10
                                } else {
                                    if ((22 != 498) || (!(-1 >= -9223372036854775808))) {
                                        p0 += 42
                                        v11 := GenS1{F0: -1, F1: p0, F2: (p0 - p0)}
                                        _ = v11
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            p0 += (^(p0))
                                        case "a":
                                            fmt.Println(fi5, p0, fi5)
                                        }
                                        v12 := []int{fi6, (-(fi6))}
                                        v13 := GenS1{F0: -104, F1: 255, F2: p0}
                                        _ = v12
                                        _ = v13
                                    }
                                    p0 -= 65535
                                    p0 = (p0 & 18446744073709551615)
                                }
                            }
                            var v14 string = ("foo bar" + "foo bar")
                            var v15 uint32 = 65535
                            _ = v14
                            _ = v15
                        } else {
                            var v16 uint64 = 4294967295
                            if false {
                                v16 += v16
                                v17 := GenS2{F0: 631, F1: (false && false), F2: -529}
                                _ = v17
                            }
                            v18 := []int{fi5, fi6, (fi6 + fi5), int(p0)}
                            _ = v16
                            _ = v18
                        }
                        p0 += p0
                        v19 := GenS1{F0: -89, F1: p0, F2: (p0 ^ p0)}
                        _ = fi6
                        _ = v19
                    }
                    switch fi5 {
                    case -907:
                        switch "" {
                        case "\n\t":
                            var v20 int32 = 903
                            _ = v20
                        case "hello":
                            var v21 int = (fi5 >> ((p0) & 63))
                            _ = v21
                        case "":
                            v22 := GenS3{F0: 342, F1: -684, F2: -223}
                            _ = v22
                        }
                    }
                    var v23 float32 = 0.1
                    _ = v23
                } else {
                    var v24 int16 = 42
                    _ = v24
                }
                _ = fi5
            }
            if false {
                switch ("\n\t" + ("foo bar" + "foo bar")) {
                case "hello":
                    v25 := []int{0}
                    _ = v25
                case "a":
                    for fi26 := 0; fi26 < 16; fi26++ {
                        var v27 uint = 802
                        var v28 float32 = 0.0
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                default:
                    v29 := GenS4{F0: 313}
                    _ = v29
                }
                v30 := GenS5{F0: 0, F1: 2, F2: 42}
                _ = v30
            } else {
                var v31 int = 618
                p0 -= 212
                fmt.Println(p0, p0)
                _ = v31
            }
        } else {
            if false {
                var v32 float32 = 1.5
                v33 := []int{32767}
                _ = v32
                _ = v33
            }
            v34 := GenS1{F0: -64, F1: 255, F2: (p0 * p0)}
            p0 = (p0 % (((p0 ^ (p0 | p0))) | 1))
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        switch "hello" {
        case "":
            for fi36 := 0; fi36 < 15; fi36++ {
                v37 := GenS4{F0: 32768}
                p0 = (p0 / (((p0 % ((p0) | 1))) | 1))
                _ = fi36
                _ = v37
            }
        case "\n\t":
            for fi38 := 0; fi38 < 5; fi38++ {
                v39 := GenS3{F0: 1, F1: 42, F2: fi38}
                var v40 uint64 = 1
                _ = fi38
                _ = v39
                _ = v40
            }
        case "foo bar":
            fmt.Println(p0, p0)
        default:
            switch 9223372036854775807 {
            case 2147483647:
                var v41 uint32 = 808
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 uint = (p0 % (((p0 ^ 4294967295)) | 1))
            var v44 uint = 255
            _ = fi42
            _ = v43
            _ = v44
        }
        v45 := []int{fi35, fi35, (fi35 & 4294967295)}
        _ = fi35
        _ = v45
    }
    v46 := GenS3{F0: 2147483648, F1: 657, F2: -9223372036854775808}
    _ = p0
    _ = v46
    return 198
}
func genFn1() int32 {
    var v47 uint32 = 0
    fmt.Println(v47)
    _ = v47
    return 1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int = 127
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v48 int64 = v1
    v49 := GenS2{F0: 42, F1: false, F2: 2147483647}
    switch ("foo bar" + ("foo bar" + ("hello" + "a"))) {
    case "":
        v48 -= v1
    case "hello":
        for fi50 := 0; fi50 < 18; fi50++ {
            switch (v2 | v2) {
            case 307:
                var v51 uint32 = 1
                _ = v51
            case 4294967295:
                var v52 int8 = -1
                _ = v52
            case 32768:
                var v53 string = "foo bar"
                _ = v53
            default:
                v54 := []int{v2, v2, -128}
                _ = v54
            }
            _ = fi50
        }
    case "foo bar":
        for fi55 := 0; fi55 < 20; fi55++ {
            for fi56 := 0; fi56 < 10; fi56++ {
                if (!(genFn1() == -763)) {
                    v57 := GenS1{F0: -128, F1: 506, F2: 214}
                    v58 := GenS6{F0: 732}
                    switch (^(v2)) {
                    case 42:
                        v1 -= (v1 / ((v48) | 1))
                    case 9223372036854775807:
                        var v59 bool = (0 > 127)
                        _ = v59
                    case 2147483647:
                        v60 := []int{int(v1), fi56, v2, (fi56 << ((42) & 63))}
                        _ = v60
                    default:
                        fmt.Println(v1)
                    }
                    _ = v57
                    _ = v58
                } else {
                    v2 = v2
                    v1 += v1
                    if false {
                        var v61 uint64 = 42
                        var v62 uint8 = 255
                        fmt.Println(v2)
                        _ = v61
                        _ = v62
                    } else {
                        if ((0.0 / 3.14159) == 3.14159) {
                            fmt.Println(fi55, v48, v48)
                            v2 = fi55
                            v63 := []int{int(v48), v2, 32768, fi56}
                            _ = v63
                        }
                        var v64 float64 = (1e10 / 0.1)
                        var v65 int = fi55
                        _ = v64
                        _ = v65
                    }
                }
                var v66 uint = 18446744073709551615
                var v67 uint64 = 1
                var v68 uint = (v66 * (v66 ^ v66))
                fmt.Println(fi55)
                _ = fi56
                _ = v66
                _ = v67
                _ = v68
            }
            delete(v4, 65535)
            var v69 int8 = -1
            _ = fi55
            _ = v69
        }
    }
    if (float32(((-2.25 + 0.1) * (3.14159 + (1e10 - (-2.25 - 0.1))))) >= (float32(0.0) - 3.14159)) {
        v70 := GenS5{F0: 18446744073709551615, F1: 1, F2: 42}
        v1 = -9223372036854775808
        fmt.Println(v1)
        _ = v70
    } else {
        var v71 uint8 = 41
        _ = v71
    }
    if false {
        var v72 int32 = 531
        if true {
            if (!((!((genFn1() >= (v72 - (v72 & v72))) || (false && (!false)))) && (!(0 < 273)))) {
                v1 = (v48 - (v1 >> ((853) & 63)))
                v72 -= genFn1()
            } else {
                var v73 int64 = (v1 ^ (v48 ^ v1))
                v74 := GenS6{F0: 533}
                if true {
                    if ("foo bar" >= ("hello" + "hello")) {
                        if (!false) {
                            var v75 int = (^(v2))
                            v76 := []int{(v2 ^ v2), v75, (v75 << ((255) & 63))}
                            fmt.Println(v48, v75)
                            _ = v75
                            _ = v76
                        } else {
                            fmt.Println(v72)
                        }
                        v1 = v48
                    } else {
                        v77 := GenS3{F0: 2, F1: v1, F2: -995}
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 3; fi78++ {
                        v72 -= (-(v72))
                        if true {
                            switch (v2 * -1) {
                            case -2147483648:
                                v79 := GenS1{F0: 0, F1: 138, F2: 593}
                                _ = v79
                            case 1:
                                for fi80 := 0; fi80 < 11; fi80++ {
                                    var v81 int = 610
                                    v2 -= fi78
                                    v82 := GenS7{F0: 1}
                                    _ = fi80
                                    _ = v81
                                    _ = v82
                                }
                            default:
                                for fi83 := 0; fi83 < 6; fi83++ {
                                    v84 := GenS7{F0: -817}
                                    v85 := GenS8{F0: (v72 - (v72 - (-(v72)))), F1: (0 >= 255)}
                                    switch v2 {
                                    case -607:
                                        switch -9223372036854775808 {
                                        case 9223372036854775807:
                                            switch v2 {
                                            case 127:
                                                v72 -= 42
                                            }
                                        case 32768:
                                            v86 := v4[238]
                                            _ = v86
                                        }
                                        v87 := GenS5{F0: 2, F1: 4294967295, F2: 105}
                                        _ = v87
                                    case 0:
                                        var v88 uint64 = 42
                                        _ = v88
                                    case 127:
                                        fmt.Println(v1)
                                    }
                                    _ = fi83
                                    _ = v84
                                    _ = v85
                                }
                            }
                        } else {
                            var v89 int8 = -1
                            v90 := GenS1{F0: (v89 % ((-64) | 1)), F1: 347, F2: 4294967295}
                            _ = v89
                            _ = v90
                        }
                        _ = fi78
                    }
                    v48 = (v48 >> ((848) & 63))
                }
                _ = v73
                _ = v74
            }
        }
        _ = v72
    } else {
        var v91 string = "a"
        v92 := GenS5{F0: 65, F1: 4294967295, F2: 17}
        fmt.Println(v91, v1)
        _ = v91
        _ = v92
    }
    fmt.Println(v1, v1)
    for fi93 := 0; fi93 < 2; fi93++ {
        v1 += v1
        _ = fi93
    }
    fmt.Println(v2, v1, v48)
    v1 = (v48 % ((v1) | 1))
    var v94 uint16 = 1
    switch "" {
    case "":
        if false {
            delete(v3, v2)
            v95 := []int{}
            _ = v95
        } else {
            v1 += (v1 << ((4294967295) & 63))
            v96 := GenS2{F0: 1, F1: (!true), F2: (v2 ^ (v2 & (v2 + 0)))}
            for fi97 := 0; fi97 < 7; fi97++ {
                v2 += v2
                _ = fi97
            }
            _ = v96
        }
    }
    var v98 uint8 = 255
    var v99 int64 = (v1 ^ (v48 >> ((18446744073709551615) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v94:", v94)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v94
    _ = v98
    _ = v99
}
