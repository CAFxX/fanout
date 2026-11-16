package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int64; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int64; F1 int8; F2 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 string; F1 string; F2 int }
type GenS10 struct { F0 int32; F1 int16; F2 int }
type GenS11 struct { F0 int32; F1 string }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int; F1 uint32 }
func genFn0() int64 {
    switch "foo bar" {
    case "\n\t":
        for fi4 := 0; fi4 < 8; fi4++ {
            var v5 uint8 = 1
            _ = fi4
            _ = v5
        }
    case "hello":
        var v6 int8 = 42
        _ = v6
    }
    if true {
        var v7 bool = true
        var v8 uint64 = 18446744073709551615
        _ = v7
        _ = v8
    } else {
        var v9 int16 = -32768
        v10 := GenS1{F0: v9}
        _ = v10
        _ = v9
    }
    return -9223372036854775808
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{}
    var v3 uint = 385
    v11 := GenS2{F0: "a", F1: genFn0(), F2: ("" < (("a" + ("" + "foo bar")) + "hello"))}
    v12 := GenS1{F0: v1}
    v13 := GenS1{F0: (-(v1))}
    for fi14 := 0; fi14 < 10; fi14++ {
        if (false || (!false)) {
            switch fi14 {
            case -567:
                v15 := GenS2{F0: "\n\t", F1: 9223372036854775807, F2: (128 >= 1)}
                _ = v15
            case 32768:
                switch fi14 {
                case 4294967295:
                    switch "" {
                    case "a":
                        var v16 int32 = -1
                        _ = v16
                    default:
                        var v17 uint64 = 250
                        _ = v17
                    }
                case 127:
                    for fi18 := 0; fi18 < 7; fi18++ {
                        if false {
                            switch "\n\t" {
                            case "foo bar":
                                if false {
                                    var v19 float32 = 1.5
                                    fmt.Println(fi14)
                                    fmt.Println(fi14)
                                    _ = v19
                                } else {
                                    v20 := GenS2{F0: "\n\t", F1: 9223372036854775807, F2: true}
                                    switch -166 {
                                    case 127:
                                        v2 = append(v2, 878)
                                    case -2147483648:
                                        if false {
                                            if ((!(!(false || false))) && ("foo bar" <= ((("hello" + ("foo bar" + "foo bar")) + "hello") + "\n\t"))) {
                                                v3 -= (v3 ^ (^(v3)))
                                                var v21 int16 = (-(v1))
                                                _ = v21
                                            }
                                            v1 -= v1
                                        }
                                    }
                                    switch fi18 {
                                    case -370:
                                        var v22 string = "hello"
                                        _ = v22
                                    case 128:
                                        fmt.Println(fi18)
                                    case 42:
                                        if false {
                                            fmt.Println(fi18)
                                            if true {
                                                v23 := GenS3{F0: -1}
                                                fmt.Println(fi14, fi18, v1)
                                                if (true || (false && true)) {
                                                    var v24 int32 = 2147483647
                                                    _ = v24
                                                } else {
                                                    v25 := GenS1{F0: v1}
                                                    v2 = append(v2, 2)
                                                    if false {
                                                        var v26 float32 = 3.14159
                                                        _ = v26
                                                    } else {
                                                        fmt.Println(fi18, fi18)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if ((v3 % (((v3 >> ((1) & 63))) | 1)) == v3) {
                                                            fmt.Println(v1, v1)
                                                            switch (((("" + "\n\t") + ("" + "hello")) + "") + "") {
                                                            case "hello":
                                                                if false {
                                                                    if false {
                                                                        _ = len(v2)
                                                                        if false {
                                                                            fmt.Println(v1)
                                                                        } else {
                                                                            var v27 int = (fi14 ^ fi18)
                                                                            switch (fi14 << (((^(v3))) & 63)) {
                                                                            case 32767:
                                                                                var v28 int32 = 600
                                                                                _ = v28
                                                                            case -2147483648:
                                                                                if (65535 == 32767) {
                                                                                    v29 := GenS4{F0: (^(v3))}
                                                                                    _ = v29
                                                                                } else {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                }
                                                                            case 4294967295:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                            var v30 float64 = float64(3.14159)
                                                                            _ = v27
                                                                            _ = v30
                                                                        }
                                                                        v3 = (v3 ^ v3)
                                                                    } else {
                                                                        if (!false) {
                                                                            v31 := GenS5{F0: 286, F1: -9223372036854775808}
                                                                            v1 += int16(v3)
                                                                            _ = v31
                                                                        } else {
                                                                            var v32 uint64 = 18446744073709551615
                                                                            v3 -= v3
                                                                            v32 = 42
                                                                            _ = v32
                                                                        }
                                                                        var v33 int = fi14
                                                                        if false {
                                                                            switch v33 {
                                                                            case 2147483647:
                                                                                fmt.Println(v1)
                                                                            case 127:
                                                                                v1 += (^(v1))
                                                                            case 183:
                                                                                switch ("" + ("" + "\n\t")) {
                                                                                case "":
                                                                                    v34 := GenS3{F0: -2147483648}
                                                                                    _ = v34
                                                                                default:
                                                                                    fmt.Println(fi14)
                                                                                }
                                                                            default:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = uint32(v3)
                                                                            }
                                                                        }
                                                                        if true {
                                                                            var v35 uint8 = 128
                                                                            _ = v35
                                                                        }
                                                                        var v36 int8 = -110
                                                                        v33 -= 32767
                                                                        _ = v33
                                                                        _ = v36
                                                                    }
                                                                }
                                                            case "a":
                                                                switch (fi14 & (fi18 - (fi14 + fi18))) {
                                                                case 32768:
                                                                    _ = len(v2)
                                                                default:
                                                                    v2 = append(v2, 940)
                                                                }
                                                            }
                                                            v1 += (v1 | (v1 + (v1 / (((^(v1))) | 1))))
                                                        }
                                                    }
                                                    _ = v25
                                                }
                                                _ = v23
                                            }
                                        } else {
                                            if false {
                                                var v37 uint64 = 2
                                                fmt.Println(fi14, v37, v37)
                                                var v38 int32 = 2147483647
                                                var v39 int8 = 1
                                                v37 = 131
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            } else {
                                                v40 := GenS6{F0: (fi18 / ((2147483647) | 1)), F1: 128}
                                                if (!true) {
                                                    if false {
                                                        v1 -= (v1 << ((v3) & 63))
                                                        var v41 int64 = 42
                                                        if (true || (51 != (v1 | (v1 >> ((1) & 63))))) {
                                                            var v42 uint32 = 2
                                                            fmt.Println(v41, v1, v3)
                                                            fmt.Println(v41, fi14)
                                                            _ = v42
                                                        } else {
                                                            fmt.Println(fi14, fi18)
                                                        }
                                                        _ = v41
                                                    }
                                                } else {
                                                    if true {
                                                        fmt.Println(v1, fi14)
                                                        if (0.1 != float64(float32((0.1 + -2.25)))) {
                                                            if false {
                                                                if false {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v1 -= (v1 << ((v3) & 63))
                                                                }
                                                                var v43 bool = true
                                                                _ = v43
                                                            } else {
                                                                switch (fi14 / ((fi18) | 1)) {
                                                                case 32767:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 42
                                                                    }
                                                                case 4294967295:
                                                                    var v44 float32 = 0.0
                                                                    _ = v44
                                                                }
                                                                var v45 uint64 = uint64(v3)
                                                                _ = v45
                                                            }
                                                            if false {
                                                                v46 := GenS1{F0: -637}
                                                                if (true && ("\n\t" != "")) {
                                                                    fmt.Println(fi14, fi14)
                                                                    v3 += (v3 % (((v3 + 4294967295)) | 1))
                                                                }
                                                                _ = v46
                                                            }
                                                        } else {
                                                            var v47 int32 = 0
                                                            var v48 uint = v3
                                                            v49 := GenS3{F0: -738}
                                                            var v50 int8 = int8(fi18)
                                                            _ = v47
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    } else {
                                                        v1 -= (v1 | (-(v1)))
                                                        _ = len(v2)
                                                        if (((!((0.0 != 1.5) || true)) || true) && true) {
                                                            var v51 uint8 = 128
                                                            _ = v51
                                                        } else {
                                                            v52 := GenS7{F0: 9223372036854775807, F1: 120, F2: (^(v3))}
                                                            v1 = (v1 / (((v1 + v1)) | 1))
                                                            _ = v52
                                                        }
                                                    }
                                                    var v53 string = "hello"
                                                    v1 = v1
                                                    _ = v53
                                                }
                                                var v54 int64 = 9223372036854775807
                                                if (genFn0() != -831) {
                                                    var v55 uint16 = 793
                                                    var v56 int8 = 96
                                                    var v57 string = "foo bar"
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    switch fi18 {
                                                    case -916:
                                                        fmt.Println(v3)
                                                    case 128:
                                                        switch (^(fi14)) {
                                                        case 130:
                                                            var v58 uint32 = 1
                                                            _ = v58
                                                        case 128:
                                                            v54 += -1
                                                        case 32767:
                                                            v54 = (v54 | genFn0())
                                                        default:
                                                            v59 := GenS1{F0: 1}
                                                            _ = v59
                                                        }
                                                    }
                                                }
                                                v54 -= v54
                                                _ = v40
                                                _ = v54
                                            }
                                        }
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = 247
                                        }
                                    }
                                    var v60 string = ((("\n\t" + "a") + "hello") + "hello")
                                    if (true || ((!(!false)) && true)) {
                                        fmt.Println(fi14, fi18)
                                    } else {
                                        v60 = "hello"
                                    }
                                    _ = v20
                                    _ = v60
                                }
                            case "":
                                fmt.Println(fi18, fi14)
                            case "a":
                                switch ("hello" + ("a" + "\n\t")) {
                                case "hello":
                                    v61 := GenS8{F0: 2147483647}
                                    _ = v61
                                default:
                                    v3 += (v3 << (((v3 / ((0) | 1))) & 63))
                                }
                            default:
                                var v62 uint = 255
                                _ = v62
                            }
                            v3 = 2
                            v2 = append(v2, 2)
                        } else {
                            if true {
                                switch fi14 {
                                case -1:
                                    if (1 >= -1) {
                                        var v63 uint8 = 79
                                        _ = v63
                                    } else {
                                        v64 := GenS9{F0: ("" + "a"), F1: ("a" + ""), F2: fi14}
                                        v3 = 18446744073709551615
                                        _ = v64
                                    }
                                case -114:
                                    v3 -= 42
                                default:
                                    var v65 bool = ((^(v3)) <= v3)
                                    v66 := GenS6{F0: (fi18 * fi18), F1: 2}
                                    _ = v65
                                    _ = v66
                                }
                            }
                            v1 -= v1
                        }
                        _ = fi18
                    }
                default:
                    if (-75 != 1) {
                        v1 -= v1
                    }
                }
            case 0:
                v1 = (v1 << ((188) & 63))
            default:
                _ = len(v2)
            }
            if (-1 >= 520) {
                var v67 float64 = (0.1 - (0.0 + 0.1))
                v3 -= 1
                v1 = v1
                _ = v67
            } else {
                v2 = append(v2, 568)
            }
            if (1e10 < -2.25) {
                for fi68 := 0; fi68 < 20; fi68++ {
                    if true {
                        switch "" {
                        case "":
                            var v69 uint64 = 4294967295
                            _ = v69
                        case "a":
                            v1 = (^(v1))
                        default:
                            switch 4294967295 {
                            case 2147483647:
                                switch "hello" {
                                case "":
                                    var v70 float64 = (((1.5 + (float64(-2.25) * 1.5)) - 3.14159) / ((0.1 / -2.25) * (((0.1 + 1.5) / 0.1) / ((1.5 / -2.25) / (0.1 + ((float64(1.5) / (0.1 * 0.1)) * 1e10))))))
                                    _ = v70
                                case "hello":
                                    _ = len(v2)
                                default:
                                    v1 = (^(v1))
                                }
                            }
                        }
                    }
                    var v71 uint64 = 42
                    _ = fi68
                    _ = v71
                }
                var v72 int64 = int64(v3)
                v73 := GenS10{F0: 2147483647, F1: v1, F2: (fi14 / ((fi14) | 1))}
                _ = v72
                _ = v73
            } else {
                v3 = v3
                for fi74 := 0; fi74 < 17; fi74++ {
                    v75 := GenS11{F0: 254, F1: ""}
                    v76 := GenS11{F0: 42, F1: "\n\t"}
                    _ = fi74
                    _ = v75
                    _ = v76
                }
            }
        } else {
            if (true && (v1 != v1)) {
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v77 int32 = -2147483648
                    _ = v77
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi78 := 0; fi78 < 7; fi78++ {
                        var v79 string = "\n\t"
                        _ = fi78
                        _ = v79
                    }
                }
                if true {
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    switch (fi14 >> ((255) & 63)) {
                    case 0:
                        if false {
                            v1 = v1
                            var v80 string = "hello"
                            _ = v80
                        } else {
                            v3 += (v3 << ((v3) & 63))
                            v1 = 1
                        }
                    }
                }
                for fi81 := 0; fi81 < 15; fi81++ {
                    v82 := GenS12{F0: 1}
                    v1 = (v1 >> (((v3 * 407)) & 63))
                    _ = fi81
                    _ = v82
                }
            } else {
                v1 -= v1
                v83 := GenS10{F0: 42, F1: -32768, F2: int(v1)}
                v84 := GenS13{F0: (fi14 | (fi14 & fi14)), F1: 1}
                _ = v83
                _ = v84
            }
        }
        _ = fi14
    }
    v1 -= 1000
    v3 = v3
    _ = len(v2)
    var v85 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v85
}
