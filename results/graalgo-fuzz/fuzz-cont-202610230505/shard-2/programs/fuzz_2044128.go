package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint; F2 bool }
type GenS6 struct { F0 int64; F1 uint; F2 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int; F1 uint64; F2 int8 }

func main() {
    var v1 uint = 255
    var v2 uint32 = 780
    v3 := []string{"hello", "\n\t"}
    v4 := make(map[int]string)
    fmt.Println(v1, v2)
    v5 := GenS1{F0: 32768, F1: 42, F2: v2}
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: (fi6 & fi6), F1: 93, F2: (v2 + v2)}
        var v8 int16 = -32768
        var v9 int8 = -1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 2, F2: (v2 % ((0) | 1))}
    switch "hello" {
    case "hello":
        fmt.Println(v1, v2)
    default:
        v11 := GenS1{F0: 832, F1: 696, F2: (v2 << ((v1) & 63))}
        _ = v11
    }
    var v12 bool = true
    if v12 {
        _ = len(v3)
        if v12 {
            if (("foo bar" + ("foo bar" + (("a" + "\n\t") + ("hello" + "a")))) != "\n\t") {
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 int = fi13
                    _ = fi13
                    _ = v14
                }
                switch 42 {
                case 116:
                    v1 -= v1
                case 32768:
                    var v15 float32 = 0.1
                    _ = v15
                case 1:
                    v16 := v4[-1]
                    _ = v16
                }
            } else {
                v1 += v1
                fmt.Println(v2, v1, v12)
            }
            if ((3.14159 - -2.25) != (1.5 / (1.5 * 1e10))) {
                v2 += (v2 >> ((255) & 63))
            } else {
                if v12 {
                    if v12 {
                        v1 -= v1
                    }
                    var v17 uint32 = (^(v2))
                    for fi18 := 0; fi18 < 11; fi18++ {
                        fmt.Println(v17, v17, fi18)
                        v1 = v1
                        _ = fi18
                    }
                    _ = v17
                } else {
                    switch (("\n\t" + "") + ("a" + "")) {
                    case "hello":
                        var v19 uint8 = 10
                        _ = v19
                    case "a":
                        switch 338 {
                        case 42:
                            var v20 int16 = 0
                            _ = v20
                        case 128:
                            if v12 {
                                v1 -= v1
                                var v21 int16 = 0
                                v1 += v1
                                _ = v21
                            } else {
                                var v22 uint8 = 1
                                _ = v22
                            }
                        default:
                            v12 = v12
                        }
                    case "foo bar":
                        if (true || v12) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    default:
                        for fi23 := 0; fi23 < 9; fi23++ {
                            v24 := v4[(fi23 - fi23)]
                            _ = fi23
                            _ = v24
                        }
                    }
                    fmt.Println(v2, v12)
                    v25 := GenS1{F0: 0, F1: 32767, F2: 214}
                    _ = v25
                }
                var v26 uint64 = 66
                _ = v26
            }
            var v27 uint = v1
            _ = v27
        } else {
            var v28 uint32 = v2
            _ = v28
        }
    } else {
        if (0 >= 1) {
            if (0 <= 597) {
                for fi29 := 0; fi29 < 2; fi29++ {
                    if (42 == 882) {
                        v30 := GenS2{F0: 9223372036854775807, F1: 0, F2: 851}
                        var v31 int64 = -9223372036854775808
                        _ = v30
                        _ = v31
                    } else {
                        var v32 int = (fi29 << ((v1) & 63))
                        _ = v32
                    }
                    _ = fi29
                }
                switch "" {
                case "hello":
                    if v12 {
                        switch (("hello" + "\n\t") + ("foo bar" + "a")) {
                        case "\n\t":
                            v1 = (^(v1))
                        default:
                            var v33 float32 = -2.25
                            _ = v33
                        }
                    } else {
                        _ = len(v4)
                        fmt.Println(v2, v12, v1)
                    }
                }
                if (-2147483648 < 2147483647) {
                    v34 := GenS3{F0: v12}
                    _ = v34
                }
            } else {
                v35 := GenS2{F0: 0, F1: -32768, F2: 221}
                var v36 string = "hello"
                v2 = (^(v2))
                _ = v35
                _ = v36
            }
            if false {
                if (float32(((((-2.25 - 1.5) / 1.0) / 1.0) - (1e10 * 0.0))) < 1.5) {
                    if false {
                        if v12 {
                            v12 = v12
                            if ((("hello" + (("hello" + ("foo bar" + ((("a" + "") + ("foo bar" + "a")) + "foo bar"))) + ("foo bar" + ("a" + "foo bar")))) == "a") || v12) {
                                var v37 int16 = 229
                                v38 := GenS4{F0: 76, F1: 1}
                                if (0 >= 1) {
                                    var v39 bool = (!true)
                                    _ = v39
                                }
                                _ = v37
                                _ = v38
                            } else {
                                for fi40 := 0; fi40 < 7; fi40++ {
                                    for fi41 := 0; fi41 < 20; fi41++ {
                                        v42 := GenS4{F0: 42, F1: 42}
                                        v12 = v12
                                        _ = fi41
                                        _ = v42
                                    }
                                    _ = fi40
                                }
                            }
                            var v43 uint16 = 0
                            _ = v43
                        }
                        switch "hello" {
                        case "\n\t":
                            v44 := v4[-770]
                            _ = v44
                        case "foo bar":
                            fmt.Println(v2)
                        case "a":
                            for fi45 := 0; fi45 < 3; fi45++ {
                                v46 := GenS4{F0: 2147483647, F1: 4294967295}
                                for fi47 := 0; fi47 < 17; fi47++ {
                                    var v48 int32 = 42
                                    v1 = v1
                                    if (!false) {
                                        switch ("" + (("a" + "\n\t") + ("a" + "hello"))) {
                                        case "foo bar":
                                            var v49 uint64 = 2
                                            _ = v49
                                        case "":
                                            var v50 uint16 = 265
                                            _ = v50
                                        }
                                        fmt.Println(fi45, v1, v12)
                                        if (false || (-9223372036854775808 <= 9223372036854775807)) {
                                            _ = len(v3)
                                            switch fi47 {
                                            case 32768:
                                                v12 = (uint32(v1) >= 237)
                                                switch (((("hello" + "foo bar") + (("hello" + "a") + ((("" + ("a" + "a")) + "foo bar") + "a"))) + "foo bar") + ("foo bar" + "foo bar")) {
                                                case "a":
                                                    var v51 float32 = (0.1 + -2.25)
                                                    _ = v51
                                                case "hello":
                                                    var v52 uint16 = 1
                                                    _ = v52
                                                default:
                                                    fmt.Println(fi45, v12)
                                                }
                                            default:
                                                if (1 > v48) {
                                                    v12 = (v12 && (0 >= 1))
                                                    _ = len(v4)
                                                    _ = len(v4)
                                                } else {
                                                    v48 = (v48 - v48)
                                                    v48 = 2147483647
                                                    v53 := v4[fi47]
                                                    _ = v53
                                                }
                                            }
                                        } else {
                                            v12 = (!(!(!(2 != 429))))
                                            switch "foo bar" {
                                            case "a":
                                                v54 := GenS4{F0: (v48 | v48), F1: 42}
                                                _ = v54
                                            case "foo bar":
                                                fmt.Println(v48, v12)
                                            case "hello":
                                                fmt.Println(fi45, v2)
                                            }
                                            var v55 uint8 = 255
                                            _ = v55
                                        }
                                    } else {
                                        v56 := GenS2{F0: 1, F1: 32767, F2: v2}
                                        _ = v56
                                    }
                                    if ((-2.25 + float64(-2.25)) >= -2.25) {
                                        if (4294967295 < 1) {
                                            v12 = (fi45 <= fi45)
                                            v57 := v4[(-(fi47))]
                                            fmt.Println(v1, fi45, v48)
                                            _ = v57
                                        } else {
                                            if (v2 <= (^(v2))) {
                                                v1 = (v1 & (v1 + v1))
                                            } else {
                                                if (1.5 >= 1.5) {
                                                    v12 = v12
                                                    var v58 uint64 = 18446744073709551615
                                                    _ = v58
                                                }
                                            }
                                            fmt.Println(v12, fi45, fi47)
                                            var v59 int16 = -32768
                                            _ = v59
                                        }
                                    }
                                    delete(v4, 32767)
                                    _ = fi47
                                    _ = v48
                                }
                                var v60 int = fi45
                                _ = fi45
                                _ = v46
                                _ = v60
                            }
                        default:
                            var v61 uint32 = v2
                            _ = v61
                        }
                    }
                    v2 += (v2 * (^(v2)))
                    for fi62 := 0; fi62 < 2; fi62++ {
                        if v12 {
                            switch fi62 {
                            case -399:
                                if (true || (!true)) {
                                    for fi63 := 0; fi63 < 8; fi63++ {
                                        fmt.Println(v2, v2, v12)
                                        var v64 uint32 = uint32(v1)
                                        switch "" {
                                        case "\n\t":
                                            v65 := GenS3{F0: v12}
                                            _ = v65
                                        case "":
                                            v66 := GenS5{F0: (-748 > 32767), F1: 18446744073709551615, F2: true}
                                            _ = v66
                                        case "a":
                                            v1 = (v1 + (v1 + 42))
                                        }
                                        _ = fi63
                                        _ = v64
                                    }
                                }
                            default:
                                v4[fi62] = ""
                            }
                            if (v12 || (!false)) {
                                v67 := GenS3{F0: (-128 != 1)}
                                v12 = (true || false)
                                _ = v67
                            } else {
                                v2 = 1
                                v2 += (^(v2))
                            }
                        }
                        v68 := GenS6{F0: 968, F1: v1, F2: uint(v2)}
                        v69 := GenS7{F0: 451}
                        _ = fi62
                        _ = v68
                        _ = v69
                    }
                } else {
                    if (v12 && false) {
                        var v70 int16 = 32767
                        var v71 bool = v12
                        v72 := GenS8{F0: v71}
                        _ = v70
                        _ = v71
                        _ = v72
                    }
                    v73 := GenS7{F0: 374}
                    v74 := GenS9{F0: 1, F1: (^(v1))}
                    _ = v73
                    _ = v74
                }
            }
        } else {
            var v75 int64 = -1
            if v12 {
                v12 = v12
                fmt.Println(v12, v1, v1)
                switch 9223372036854775807 {
                case 128:
                    var v76 bool = (false && ((v2 - v2) == (v2 / ((uint32(v1)) | 1))))
                    _ = v76
                default:
                    var v77 int16 = 32767
                    _ = v77
                }
            }
            if v12 {
                v12 = (false && (("" + "") < ""))
                v1 -= 306
                v78 := GenS8{F0: v12}
                _ = v78
            }
            _ = v75
        }
        v79 := GenS1{F0: -9223372036854775808, F1: 255, F2: (v2 ^ (v2 ^ v2))}
        _ = v79
    }
    var v80 float32 = (((0.1 * 0.1) - (((1e10 - (1e10 - 3.14159)) * 3.14159) - 0.1)) * (0.1 + (float32(float64(0.1)) / 0.1)))
    v3 = append(v3, ((("\n\t" + "") + "a") + "\n\t"))
    v12 = (-2147483648 >= 42)
    v80 = ((v80 - -2.25) + 0.1)
    v80 += v80
    v81 := GenS10{F0: 128, F1: 4294967295, F2: -34}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v80:", v80)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v81
}
