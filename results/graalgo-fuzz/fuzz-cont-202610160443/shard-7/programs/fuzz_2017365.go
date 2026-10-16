package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int64; F2 int32 }
type GenS7 struct { F0 int8; F1 string; F2 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 bool; F2 int16 }
type GenS10 struct { F0 string; F1 int64; F2 int64 }
type GenS11 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 1e10
    var v3 int16 = -32768
    if (3.14159 != ((((float64(float32(0.0)) * (0.0 * 3.14159)) * 0.0) + 0.0) + (1.5 * (-2.25 * 0.1)))) {
        var v4 int16 = (-(v3))
        v4 = v4
        if false {
            v3 -= (^(v3))
            var v5 string = "a"
            _ = v5
        }
        _ = v4
    }
    v2 = (v2 - (float32(-2.25) + float32(float64(((((v2 / v2) * v2) * v2) - 1.5)))))
    v2 += float32(-2.25)
    switch "" {
    case "a":
        var v6 int8 = 77
        _ = v6
    case "":
        if (!(65535 == 42)) {
            var v7 uint8 = 42
            _ = v7
        } else {
            if (true && (-1 == 0)) {
                v8 := GenS1{F0: 72, F1: -1, F2: 236}
                if (42 == 42) {
                    var v9 uint = 4294967295
                    var v10 int8 = 42
                    _ = v10
                    _ = v9
                } else {
                    var v11 int8 = 17
                    _ = v11
                }
                var v12 uint32 = 4294967295
                _ = v12
                _ = v8
            }
            if (2 < 688) {
                if false {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 float32 = 0.1
                        v3 -= v3
                        _ = fi13
                        _ = v14
                    }
                    v1[-1] = ("\n\t" + "foo bar")
                    v2 += float32(1.5)
                }
            } else {
                v3 += v3
                delete(v1, -302)
            }
        }
    default:
        if true {
            v3 = v3
            switch -353 {
            case -2147483648:
                v15 := GenS2{F0: 1}
                _ = v15
            default:
                var v16 uint8 = 1
                _ = v16
            }
            var v17 float32 = (v2 * (0.0 + ((((v2 / v2) / ((((v2 - float32((1.5 * 1.5))) * v2) - (v2 * (-2.25 + 3.14159))) * 1.5)) / (-2.25 - 0.1)) * 1e10)))
            _ = v17
        } else {
            v18 := GenS3{F0: ("" + "hello"), F1: (2147483648 >= 4294967295)}
            v3 -= (^(v3))
            _ = v18
        }
    }
    var v19 uint16 = 255
    if true {
        v3 += v3
        switch -2147483648 {
        case 1:
            v1[1] = "a"
        case 32768:
            var v20 uint32 = 2
            _ = v20
        default:
            var v21 int64 = 9223372036854775807
            _ = v21
        }
    } else {
        v22 := GenS3{F0: "a", F1: ((342 < 995) && (v19 <= (v19 - v19)))}
        var v23 float32 = v2
        _ = v22
        _ = v23
    }
    if false {
        v24 := []int{-789, 32768, -7}
        if true {
            for fi25 := 0; fi25 < 17; fi25++ {
                if true {
                    _ = len(v24)
                    _ = len(v1)
                } else {
                    var v26 uint8 = 0
                    switch ("" + "\n\t") {
                    case "\n\t":
                        v19 -= v19
                    case "hello":
                        if true {
                            var v27 int = (fi25 & (fi25 - fi25))
                            _ = v27
                        } else {
                            var v28 int = fi25
                            fmt.Println(v26, v28, v28)
                            switch "hello" {
                            case "":
                                _ = len(v24)
                            case "a":
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v26 = v26
                                    _ = len(v24)
                                    if false {
                                        v1[(fi25 << ((1) & 63))] = ("foo bar" + ("foo bar" + "a"))
                                    } else {
                                        var v30 bool = (!false)
                                        if (float64(-2.25) >= float64((v2 * v2))) {
                                            v31 := GenS2{F0: v26}
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                            if v30 {
                                                switch (("\n\t" + "hello") + ("\n\t" + "")) {
                                                case "":
                                                    if (v19 >= (v19 - 42)) {
                                                        delete(v1, v28)
                                                        v3 -= v3
                                                        if ((556 < 2) || (!(!(27 >= 42)))) {
                                                            if v30 {
                                                                var v32 int32 = 1
                                                                switch (fi25 - v28) {
                                                                case -2147483648:
                                                                    v33 := GenS3{F0: "foo bar", F1: v30}
                                                                    _ = v33
                                                                case 9223372036854775807:
                                                                    v3 = 1
                                                                case 220:
                                                                    var v34 float64 = 1.5
                                                                    _ = v34
                                                                default:
                                                                    v32 += 1
                                                                }
                                                                _ = v32
                                                            }
                                                            if v30 {
                                                                v35 := GenS1{F0: 42, F1: -525, F2: (v26 + v26)}
                                                                _ = v35
                                                            }
                                                        } else {
                                                            v3 = 774
                                                            switch fi25 {
                                                            case 32768:
                                                                v28 += fi29
                                                                v2 = (1.5 + v2)
                                                            case -9223372036854775808:
                                                                if (255 <= v19) {
                                                                    var v36 uint16 = v19
                                                                    v2 -= ((0.0 * v2) + ((3.14159 / v2) * ((-2.25 * v2) / v2)))
                                                                    fmt.Println(v28)
                                                                    fmt.Println(v26, fi29)
                                                                    _ = v36
                                                                }
                                                            case 127:
                                                                var v37 float32 = (v2 / v2)
                                                                _ = v37
                                                            }
                                                            var v38 int8 = 42
                                                            v26 = (v26 | 42)
                                                            v30 = ((v30 || ((v3 / ((-1) | 1)) == 1)) && (60 != v26))
                                                            _ = v38
                                                        }
                                                        var v39 uint32 = 2147483647
                                                        if v30 {
                                                            if len(v24) > 0 {
                                                                v24[0] = 0
                                                            }
                                                            delete(v1, fi29)
                                                            if v30 {
                                                                var v40 uint8 = 236
                                                                _ = v40
                                                            }
                                                            var v41 int8 = -99
                                                            v42 := GenS2{F0: v26}
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        _ = v39
                                                    } else {
                                                        switch "" {
                                                        case "\n\t":
                                                            var v43 uint64 = 18446744073709551615
                                                            _ = v43
                                                        case "a":
                                                            v26 = 2
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                        var v44 bool = true
                                                        var v45 int = fi29
                                                        v26 = v26
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                case "foo bar":
                                                    switch fi29 {
                                                    case 42:
                                                        v46 := GenS4{F0: 201, F1: 4294967295}
                                                        _ = v46
                                                    case 128:
                                                        var v47 int8 = -1
                                                        _ = v47
                                                    case 127:
                                                        var v48 uint8 = (v26 - (v26 % ((128) | 1)))
                                                        _ = v48
                                                    }
                                                }
                                                var v49 uint16 = (^(v19))
                                                var v50 string = "\n\t"
                                                _ = v49
                                                _ = v50
                                            } else {
                                                if true {
                                                    v51 := GenS5{F0: 0, F1: 133, F2: 2}
                                                    if v30 {
                                                        if true {
                                                            var v52 int64 = -1
                                                            _ = v52
                                                        } else {
                                                            v53 := GenS6{F0: 4294967295, F1: -9223372036854775808, F2: -986}
                                                            _ = v53
                                                        }
                                                    } else {
                                                        if (fi25 > 127) {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v54 int32 = -1
                                                                _ = v54
                                                            }
                                                            fmt.Println(v19)
                                                        }
                                                    }
                                                    _ = v51
                                                } else {
                                                    v3 = v3
                                                }
                                                fmt.Println(v28)
                                            }
                                            fmt.Println(fi25)
                                            _ = v31
                                        } else {
                                            fmt.Println(v28, fi29)
                                            if v30 {
                                                if v30 {
                                                    v30 = v30
                                                    if len(v24) > 0 {
                                                        v24[0] = v28
                                                    }
                                                    v55 := v1[fi25]
                                                    _ = v55
                                                } else {
                                                    v56 := GenS3{F0: (("" + "foo bar") + "hello"), F1: (v28 < int(v26))}
                                                    var v57 uint16 = 394
                                                    _ = v56
                                                    _ = v57
                                                }
                                                v2 += v2
                                            }
                                            v24 = append(v24, (-(fi29)))
                                            if false {
                                                v1[fi29] = ("" + "a")
                                                v2 = (1e10 + (float32((0.0 / float64(-2.25))) * 0.0))
                                            }
                                            _ = len(v24)
                                        }
                                        _ = v30
                                    }
                                    _ = fi29
                                }
                            case "foo bar":
                                v58 := GenS7{F0: 1, F1: "a", F2: -953}
                                _ = v58
                            default:
                                v59 := v1[v28]
                                _ = v59
                            }
                            _ = v28
                        }
                    default:
                        for fi60 := 0; fi60 < 13; fi60++ {
                            var v61 int32 = -2147483648
                            v19 -= (v19 & 255)
                            _ = fi60
                            _ = v61
                        }
                    }
                    _ = v26
                }
                _ = fi25
            }
        }
        _ = v24
    } else {
        fmt.Println(v2)
        v1[128] = (("\n\t" + ("\n\t" + "a")) + "\n\t")
        v62 := GenS8{F0: (v19 * (v19 ^ v19))}
        _ = v62
    }
    switch "foo bar" {
    case "\n\t":
        var v63 int8 = 100
        _ = v63
    case "":
        v19 -= v19
    }
    v64 := GenS8{F0: 32767}
    if (42 < 929) {
        var v65 uint32 = 332
        _ = v65
    } else {
        var v66 int = 32768
        v19 = 791
        for fi67 := 0; fi67 < 18; fi67++ {
            var v68 bool = ((false && false) && (!(false || true)))
            fmt.Println(v19, v68, v19)
            _ = fi67
            _ = v68
        }
        _ = v66
    }
    for fi69 := 0; fi69 < 19; fi69++ {
        v70 := GenS9{F0: (18446744073709551615 != 2), F1: (((false || (24 != 42)) || (!(!false))) || true), F2: (v3 << ((868) & 63))}
        for fi71 := 0; fi71 < 20; fi71++ {
            switch ("a" + "\n\t") {
            case "a":
                v72 := []int{}
                _ = v72
            case "":
                v19 = 2
            case "foo bar":
                v73 := GenS4{F0: -2147483648, F1: 2147483647}
                _ = v73
            default:
                v2 = float32(float64(v2))
            }
            if true {
                var v74 uint16 = v19
                var v75 float64 = 1.5
                _ = v74
                _ = v75
            }
            var v76 float64 = 1.5
            _ = fi71
            _ = v76
        }
        _ = fi69
        _ = v70
    }
    var v77 uint16 = (v19 >> ((65535) & 63))
    v2 += ((v2 + (float32((1.5 / 1.5)) + v2)) * (v2 - (v2 + float32((0.1 + ((((1e10 - 1.5) + (0.1 / 1.0)) - 0.1) / 1.5))))))
    if (!(620 <= 18446744073709551615)) {
        fmt.Println(v77)
        for fi78 := 0; fi78 < 9; fi78++ {
            v79 := GenS10{F0: "a", F1: -9223372036854775808, F2: 9223372036854775807}
            _ = fi78
            _ = v79
        }
        for fi80 := 0; fi80 < 14; fi80++ {
            fmt.Println(v19, v19, v2)
            for fi81 := 0; fi81 < 18; fi81++ {
                fmt.Println(v19)
                _ = fi81
            }
            if true {
                var v82 float32 = 3.14159
                v83 := GenS3{F0: "a", F1: (-30 == 34)}
                var v84 float32 = (v82 / 3.14159)
                _ = v82
                _ = v83
                _ = v84
            }
            _ = fi80
        }
    }
    _ = len(v1)
    if true {
        fmt.Println(v3, v2, v3)
        v85 := GenS11{F0: 42, F1: v3}
        var v86 int = 363
        _ = v85
        _ = v86
    }
    var v87 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v64
    _ = v77
    _ = v87
}
