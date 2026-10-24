package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64; F2 int8 }
type GenS4 struct { F0 int; F1 bool; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 int64; F2 string }
type GenS8 struct { F0 uint16; F1 int16; F2 int }
type GenS9 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint = 1
    var v2 int = 32767
    var v3 uint = 0
    var v4 int64 = 284
    var v5 uint16 = 877
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 18; fi7++ {
        v5 = v5
        var v8 int8 = -128
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint16 = uint16(v3)
        _ = fi9
        _ = v10
    }
    var v11 int8 = 1
    v12 := GenS1{F0: (v11 >> ((uint(v5)) & 63))}
    var v13 int64 = -815
    v14 := GenS2{F0: v3}
    _ = len(v6)
    var v15 float64 = 0.1
    switch v2 {
    case -9223372036854775808:
        v4 -= (v13 | v13)
    case -783:
        if false {
            var v16 uint = (^(v3))
            var v17 uint32 = 2147483648
            if true {
                fmt.Println(v3, v13)
                var v18 int16 = 1
                _ = v18
            } else {
                switch v2 {
                case 9223372036854775807:
                    v19 := GenS1{F0: v11}
                    _ = v19
                }
                switch "hello" {
                case "\n\t":
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v21 := []int{(v2 - -562), 0}
                        fmt.Println(v5)
                        _ = fi20
                        _ = v21
                    }
                case "a":
                    v22 := []int{int(v3)}
                    _ = v22
                case "hello":
                    v23 := GenS2{F0: (v1 % ((847) | 1))}
                    _ = v23
                }
                var v24 int64 = v13
                _ = v24
            }
            _ = v16
            _ = v17
        }
    case 347:
        v25 := []int{-1, 934}
        _ = v25
    default:
        if (("a" >= "\n\t") && (!true)) {
            v26 := GenS2{F0: v1}
            v27 := []int{}
            fmt.Println(v5, v11)
            _ = v26
            _ = v27
        } else {
            if false {
                v28 := GenS3{F0: (42 >= 0), F1: 588, F2: v11}
                _ = v28
            }
            if true {
                v1 = (v1 >> (((v1 / ((uint(v5)) | 1))) & 63))
            } else {
                switch (v2 - v2) {
                case 128:
                    if false {
                        switch (v2 - (v2 ^ (^(v2)))) {
                        case 128:
                            v6[4294967295] = v2
                        default:
                            switch ("" + "hello") {
                            case "":
                                _ = len(v6)
                            case "a":
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    var v30 float64 = ((((v15 + v15) + (v15 / v15)) * v15) - (v15 * v15))
                                    _ = fi29
                                    _ = v30
                                }
                            default:
                                v1 -= 1
                            }
                        }
                        v6[1] = (v2 & (^(v2)))
                        for fi31 := 0; fi31 < 17; fi31++ {
                            var v32 int8 = (^(v11))
                            v33 := GenS4{F0: -1, F1: false, F2: 1}
                            v2 = (^(v2))
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                case -255:
                    var v34 int32 = 1
                    _ = v34
                case 32767:
                    v3 -= 2
                }
                fmt.Println(v3)
            }
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 int32 = 0
                if true {
                    v4 += (v13 / (((v13 / ((v13) | 1))) | 1))
                }
                v37 := []int{}
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    }
    fmt.Println(v11, v2)
    if false {
        v15 = v15
        switch "foo bar" {
        case "a":
            if true {
                if true {
                    if true {
                        v3 = v3
                        switch (("\n\t" + (("\n\t" + "foo bar") + "hello")) + ("" + "a")) {
                        case "\n\t":
                            v15 = ((-2.25 * (float64(0.0) + float64(float32(((float64(((0.0 - float32((v15 / (v15 * v15)))) / 1.0)) - (float64((0.0 / 3.14159)) - 0.0)) - v15))))) + 1.5)
                        case "hello":
                            for fi38 := 0; fi38 < 2; fi38++ {
                                v2 -= (v2 ^ v2)
                                v11 = v11
                                if (995 == (^(v1))) {
                                    if (549 < 532) {
                                        if (false || true) {
                                            v39 := v6[0]
                                            _ = v39
                                        }
                                        v40 := []int{-2147483648, v2, 0, (v2 & 4294967295)}
                                        _ = v40
                                    }
                                    for fi41 := 0; fi41 < 11; fi41++ {
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v2)
                                        case "foo bar":
                                            var v42 float32 = float32((v15 + ((((float64((0.1 * 0.1)) + ((v15 + v15) - float64((3.14159 - 1.5)))) - (1.5 + v15)) / -2.25) * v15)))
                                            _ = v42
                                        case "":
                                            var v43 uint32 = 677
                                            _ = v43
                                        }
                                        _ = fi41
                                    }
                                    v44 := GenS3{F0: false, F1: 18446744073709551615, F2: (v11 >> ((v1) & 63))}
                                    _ = v44
                                } else {
                                    v6[42] = (v2 % ((1) | 1))
                                    v45 := GenS4{F0: -684, F1: ((v5 % ((42) | 1)) >= (v5 | (v5 >> ((v1) & 63)))), F2: 42}
                                    _ = v45
                                }
                                _ = fi38
                            }
                        }
                        v3 = (v3 - (v3 ^ v3))
                    } else {
                        fmt.Println(v13, v5, v11)
                    }
                } else {
                    switch (v2 & v2) {
                    case -1:
                        v46 := GenS5{F0: 0}
                        _ = v46
                    case -9223372036854775808:
                        v3 = (v3 / (((v3 / (((v1 - (^(v1)))) | 1))) | 1))
                    case -997:
                        v47 := GenS6{F0: 200}
                        _ = v47
                    default:
                        fmt.Println(v13, v15)
                    }
                    if true {
                        var v48 uint64 = 4294967295
                        _ = v48
                    }
                    if true {
                        fmt.Println(v13, v11, v5)
                    } else {
                        for fi49 := 0; fi49 < 10; fi49++ {
                            v50 := GenS7{F0: v2, F1: v13, F2: "foo bar"}
                            _ = fi49
                            _ = v50
                        }
                        fmt.Println(v1)
                    }
                }
            } else {
                if true {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        v52 := GenS3{F0: true, F1: 253, F2: (v11 + v11)}
                        if false {
                            v3 += v3
                            switch -675 {
                            case -9223372036854775808:
                                v2 -= (-(fi51))
                            case -2147483648:
                                v3 = v1
                            default:
                                var v53 uint8 = 159
                                _ = v53
                            }
                            v2 = fi51
                        }
                        v4 -= int64(v11)
                        _ = fi51
                        _ = v52
                    }
                }
                v54 := v6[0]
                var v55 uint16 = v5
                _ = v54
                _ = v55
            }
        case "":
            for fi56 := 0; fi56 < 19; fi56++ {
                v57 := GenS8{F0: 2, F1: int16(v2), F2: fi56}
                var v58 string = ("a" + "foo bar")
                var v59 uint8 = 178
                _ = fi56
                _ = v57
                _ = v58
                _ = v59
            }
        case "\n\t":
            if (!(true || (887 == 0))) {
                var v60 uint = (v1 | (^(v3)))
                delete(v6, 4294967295)
                _ = v60
            }
        }
        if true {
            v2 = v2
            for fi61 := 0; fi61 < 9; fi61++ {
                v62 := GenS9{F0: 0, F1: 2, F2: v5}
                for fi63 := 0; fi63 < 10; fi63++ {
                    v64 := GenS5{F0: -1}
                    v65 := GenS10{F0: 0, F1: v11}
                    switch (fi63 / ((32768) | 1)) {
                    case 32768:
                        v11 = v11
                    case -9223372036854775808:
                        if (true && (!true)) {
                            switch fi61 {
                            case 32767:
                                fmt.Println(fi63)
                            default:
                                v66 := GenS1{F0: (v11 ^ 42)}
                                _ = v66
                            }
                            if ((1.5 * 1e10) >= (v15 / v15)) {
                                switch (v2 | 556) {
                                case -2147483648:
                                    if (42 < int32(v11)) {
                                        switch "a" {
                                        case "a":
                                            var v67 uint32 = 0
                                            _ = v67
                                        case "\n\t":
                                            v4 += (v4 + v13)
                                        case "foo bar":
                                            v68 := GenS8{F0: v5, F1: -32768, F2: fi63}
                                            _ = v68
                                        default:
                                            var v69 uint64 = 0
                                            v70 := GenS11{F0: 65535}
                                            _ = v69
                                            _ = v70
                                        }
                                        var v71 int = (fi63 % ((fi61) | 1))
                                        if (((^(v3)) < (v3 / (((v3 >> ((769) & 63))) | 1))) && (((v4 << (((^(v3))) & 63)) > (v4 / ((v13) | 1))) && (!((v11 & (v11 << ((v3) & 63))) <= v11)))) {
                                            var v72 uint64 = 371
                                            v73 := GenS12{F0: 21, F1: 255}
                                            _ = v72
                                            _ = v73
                                        }
                                        _ = v71
                                    } else {
                                        _ = len(v6)
                                        var v74 uint16 = (v5 & v5)
                                        v75 := GenS12{F0: 0, F1: 1}
                                        _ = v74
                                        _ = v75
                                    }
                                default:
                                    var v76 int32 = 0
                                    _ = v76
                                }
                            }
                            v1 = (v1 + v1)
                            v1 -= v3
                        } else {
                            fmt.Println(fi63)
                            fmt.Println(v4)
                            var v77 int32 = -1
                            _ = v77
                        }
                    case 128:
                        var v78 uint32 = 647
                        _ = v78
                    default:
                        v79 := GenS13{F0: 1, F1: (true || false)}
                        _ = v79
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                _ = fi61
                _ = v62
            }
            v2 = (v2 / ((v2) | 1))
        }
    } else {
        if ((0.1 >= -2.25) && false) {
            for fi80 := 0; fi80 < 7; fi80++ {
                v4 += (v13 ^ -1)
                var v81 uint32 = 584
                _ = fi80
                _ = v81
            }
        } else {
            fmt.Println(v1)
            v82 := v6[4294967295]
            var v83 uint64 = 18446744073709551615
            _ = v82
            _ = v83
        }
        var v84 float32 = (0.1 + (1.5 - -2.25))
        v11 += -81
        _ = v84
    }
    v85 := []int{(v2 ^ (v2 % (((v2 - v2)) | 1))), (v2 - 1)}
    delete(v6, 117)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v85:", len(v85), v85)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
}
