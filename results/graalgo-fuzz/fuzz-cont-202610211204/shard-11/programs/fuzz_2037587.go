package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint32; F1 int8; F2 string }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 int64; F1 bool }

func main() {
    var v1 float32 = 3.14159
    var v2 int32 = 2147483647
    var v3 uint64 = 1
    var v4 bool = false
    v5 := []int{-979, -519, 761, 4294967295}
    var v6 uint64 = v3
    v1 -= (v1 / v1)
    v4 = (!(1 != 75))
    var v7 int16 = -797
    switch "foo bar" {
    case "\n\t":
        v8 := GenS1{F0: 590}
        _ = v8
    case "":
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS2{F0: v7, F1: (fi9 * (fi9 - fi9)), F2: "a"}
            var v11 string = "foo bar"
            v12 := GenS2{F0: 1, F1: 849, F2: (v11 + v11)}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v6, v6)
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS3{F0: 1}
        if len(v5) > 0 {
            v5[0] = fi13
        }
        _ = fi13
        _ = v14
    }
    if v4 {
        var v15 string = "\n\t"
        _ = v15
    } else {
        var v16 uint64 = (v6 << ((1) & 63))
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        switch ("hello" + "\n\t") {
        case "":
            for fi18 := 0; fi18 < 7; fi18++ {
                if (v4 && ((255 != 513) && (1 < 9223372036854775807))) {
                    var v19 uint32 = 4294967295
                    var v20 int = 776
                    _ = v19
                    _ = v20
                } else {
                    var v21 float32 = (v1 + ((0.1 / 1.0) - float32(3.14159)))
                    v22 := GenS2{F0: v7, F1: -21, F2: ("\n\t" + "")}
                    fmt.Println(v7)
                    _ = v21
                    _ = v22
                }
                switch (("hello" + "foo bar") + ("a" + (("\n\t" + "foo bar") + (("a" + "hello") + ("a" + "a"))))) {
                case "":
                    switch fi17 {
                    case 250:
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            v6 = v3
                        case "foo bar":
                            if v4 {
                                fmt.Println(v2, v2, v3)
                            } else {
                                var v23 int = fi17
                                v2 = v2
                                v24 := GenS3{F0: 785}
                                _ = v23
                                _ = v24
                            }
                        case "":
                            if v4 {
                                v2 += (v2 | v2)
                                if len(v5) > 0 {
                                    v5[0] = -1
                                }
                                switch fi17 {
                                case 42:
                                    fmt.Println(v2, v3)
                                default:
                                    var v25 int32 = (v2 | -1)
                                    _ = v25
                                }
                            }
                        }
                    case 1:
                        if (v2 > v2) {
                            if v4 {
                                if (!(!(v4 || (false && false)))) {
                                    var v26 int8 = int8(fi17)
                                    v2 -= v2
                                    v3 = 2
                                    _ = v26
                                } else {
                                    v3 = v6
                                    v1 = (-2.25 * (v1 - ((1e10 * v1) - v1)))
                                    v2 += v2
                                }
                                var v27 int8 = -128
                                v28 := GenS3{F0: 4294967295}
                                var v29 int8 = v27
                                v6 += v3
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int16 = -178
                                fmt.Println(v4, v2, v1)
                                if true {
                                    var v31 int16 = (v7 >> ((78) & 63))
                                    var v32 uint64 = v6
                                    var v33 uint32 = 707
                                    switch "foo bar" {
                                    case "foo bar":
                                        _ = len(v5)
                                    case "hello":
                                        v34 := GenS2{F0: int16(v3), F1: fi17, F2: "\n\t"}
                                        _ = v34
                                    default:
                                        v32 -= v6
                                    }
                                    v35 := GenS4{F0: 21}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                    _ = v35
                                } else {
                                    if v4 {
                                        var v36 uint16 = 845
                                        v37 := GenS5{F0: 62}
                                        _ = v36
                                        _ = v37
                                    }
                                }
                                _ = v30
                            }
                            var v38 int32 = (v2 * v2)
                            _ = v38
                        } else {
                            v3 = (v3 & 2)
                        }
                    }
                case "hello":
                    v39 := GenS1{F0: 0}
                    _ = v39
                case "\n\t":
                    var v40 uint8 = 1
                    _ = v40
                }
                _ = fi18
            }
        case "hello":
            v41 := GenS6{F0: "\n\t", F1: 2147483647, F2: 42}
            _ = v41
        }
        _ = fi17
    }
    var v42 int = 1
    for fi43 := 0; fi43 < 20; fi43++ {
        v42 -= (fi43 >> ((1) & 63))
        _ = fi43
    }
    switch (("" + ("hello" + "hello")) + ("\n\t" + ("foo bar" + ("a" + (("" + "foo bar") + ("\n\t" + "")))))) {
    case "a":
        switch v42 {
        case -9223372036854775808:
            for fi44 := 0; fi44 < 18; fi44++ {
                var v45 uint = 18446744073709551615
                _ = len(v5)
                switch "a" {
                case "foo bar":
                    for fi46 := 0; fi46 < 1; fi46++ {
                        var v47 int64 = -9223372036854775808
                        var v48 float32 = (((1e10 * ((float32((0.0 - float64(1e10))) - v1) - 0.0)) / v1) / (0.1 / ((v1 / v1) * v1)))
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                }
                _ = fi44
                _ = v45
            }
        case 4294967295:
            if true {
                v49 := GenS7{F0: 0, F1: 42, F2: ("" + "foo bar")}
                switch "" {
                case "":
                    if len(v5) > 0 {
                        v5[0] = v42
                    }
                case "\n\t":
                    var v50 int64 = 406
                    _ = v50
                }
                var v51 int8 = -1
                _ = v49
                _ = v51
            } else {
                v42 -= (v42 * (v42 | v42))
                var v52 int64 = 948
                for fi53 := 0; fi53 < 18; fi53++ {
                    if ((v7 / ((v7) | 1)) <= 32767) {
                        v54 := GenS8{F0: 392, F1: 662}
                        _ = v54
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v55 := GenS2{F0: v7, F1: (fi53 << ((0) & 63)), F2: "hello"}
                        v6 -= (v3 - 0)
                        _ = v55
                    }
                    v56 := GenS3{F0: 2147483647}
                    var v57 int32 = (^(v2))
                    _ = fi53
                    _ = v56
                    _ = v57
                }
                _ = v52
            }
        default:
            if v4 {
                switch ("hello" + ("hello" + (((((("hello" + ("hello" + ("\n\t" + ("a" + "a")))) + "hello") + "") + "foo bar") + ("\n\t" + ("a" + ("a" + "\n\t")))) + "a"))) {
                case "hello":
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v42, v2)
                    }
                }
                v58 := GenS1{F0: 42}
                var v59 bool = ((v42 + v42) == (v42 >> ((18446744073709551615) & 63)))
                _ = v58
                _ = v59
            } else {
                if (-105 < 127) {
                    v60 := GenS3{F0: 166}
                    switch (("" + "") + ("hello" + "a")) {
                    case "a":
                        if v4 {
                            fmt.Println(v7, v3)
                        } else {
                            var v61 int32 = (-(v2))
                            switch (v42 | (v42 / ((v42) | 1))) {
                            case -1:
                                var v62 float32 = v1
                                _ = v62
                            case 0:
                                switch 42 {
                                case -9223372036854775808:
                                    for fi63 := 0; fi63 < 20; fi63++ {
                                        var v64 int8 = -75
                                        var v65 int8 = (v64 >> ((42) & 63))
                                        _ = fi63
                                        _ = v64
                                        _ = v65
                                    }
                                case 128:
                                    for fi66 := 0; fi66 < 16; fi66++ {
                                        var v67 int32 = (v61 % ((-2147483648) | 1))
                                        v6 += v6
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = fi66
                                        _ = v67
                                    }
                                }
                            case 128:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = v61
                        }
                    case "foo bar":
                        if (-9223372036854775808 <= 42) {
                            fmt.Println(v4, v6)
                            v68 := GenS5{F0: 128}
                            _ = v68
                        } else {
                            if len(v5) > 0 {
                                v5[0] = -552
                            }
                            switch "" {
                            case "hello":
                                if false {
                                    fmt.Println(v1)
                                }
                            case "a":
                                switch "hello" {
                                case "a":
                                    for fi69 := 0; fi69 < 18; fi69++ {
                                        var v70 uint = 42
                                        v4 = (!(!(!false)))
                                        _ = fi69
                                        _ = v70
                                    }
                                }
                            }
                        }
                    case "\n\t":
                        var v71 int16 = -32768
                        _ = v71
                    default:
                        v72 := GenS1{F0: 2}
                        _ = v72
                    }
                    _ = v60
                }
            }
        }
    default:
        if true {
            v73 := GenS8{F0: 244, F1: 0}
            _ = v73
        } else {
            var v74 uint = 0
            for fi75 := 0; fi75 < 18; fi75++ {
                v76 := GenS3{F0: 92}
                v77 := GenS9{F0: 255, F1: 1}
                _ = fi75
                _ = v76
                _ = v77
            }
            var v78 uint = v74
            _ = v74
            _ = v78
        }
    }
    if ((v3 | (v3 ^ (v3 * (v6 + 18446744073709551615)))) <= 18446744073709551615) {
        fmt.Println(v3, v3)
        v79 := GenS10{F0: 1, F1: (1 < (v2 + v2))}
        _ = len(v5)
        _ = v79
    }
    var v80 uint16 = 255
    var v81 int = v42
    var v82 uint64 = (v6 >> ((65535) & 63))
    var v83 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
