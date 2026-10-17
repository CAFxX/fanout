package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }
type GenS5 struct { F0 int8; F1 int32; F2 uint8 }
type GenS6 struct { F0 int8; F1 int16; F2 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 bool; F1 int16; F2 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string; F1 uint; F2 int8 }
type GenS13 struct { F0 int32; F1 int8; F2 uint32 }
type GenS14 struct { F0 int32; F1 int32; F2 int32 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{1, 4294967295, 2147483647}
    v3 := make(map[int]bool)
    v4 := []int{111, 32768, 32768}
    v5 := []uint32{}
    fmt.Println(v1)
    var v6 string = ("foo bar" + v1)
    switch 0 {
    case 32768:
        v7 := GenS1{F0: 128, F1: -424, F2: 380}
        _ = v7
    case -1:
        if (2 > 42) {
            if (((3.14159 - (1e10 / 1.0)) > float64(-2.25)) && true) {
                v1 = "foo bar"
            } else {
                var v8 int8 = 59
                v9 := GenS2{F0: 65535}
                _ = v8
                _ = v9
            }
            v6 = ("hello" + v1)
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := GenS3{F0: 127, F1: 1}
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(v1, v1)
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                _ = len(v3)
                if true {
                    fmt.Println(fi13)
                }
                _ = fi13
            }
            v14 := GenS4{F0: 100, F1: 1, F2: -1}
            var v15 uint32 = 0
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v4 = append(v4, 938)
        _ = fi16
    }
    switch 2147483647 {
    case 128:
        fmt.Println(v6, v6)
    case 2147483647:
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint8 = 0
            v18 = (v18 >> ((227) & 63))
            for fi19 := 0; fi19 < 14; fi19++ {
                delete(v3, -34)
                if true {
                    fmt.Println(fi19, v6)
                    if false {
                        var v20 uint16 = 65535
                        if (false && false) {
                            v20 -= (v20 >> ((65535) & 63))
                            var v21 float32 = 1.5
                            _ = v21
                        } else {
                            var v22 int32 = -2147483648
                            _ = v22
                        }
                        _ = v20
                    } else {
                        if (false || (false && (false || (4294967295 == 65535)))) {
                            switch "a" {
                            case "hello":
                                switch "a" {
                                case "\n\t":
                                    switch (^(fi17)) {
                                    case 236:
                                        var v23 int8 = -1
                                        _ = v23
                                    case 32768:
                                        var v24 int32 = 0
                                        _ = v24
                                    }
                                }
                            }
                            if true {
                                var v25 uint64 = 57
                                v26 := GenS4{F0: 1, F1: 1, F2: -39}
                                v27 := GenS5{F0: 0, F1: -2147483648, F2: v18}
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                if false {
                                    if true {
                                        var v28 int16 = -343
                                        switch -2147483648 {
                                        case -9223372036854775808:
                                            v18 = (^(v18))
                                        default:
                                            v18 = (v18 - 255)
                                        }
                                        _ = v28
                                    } else {
                                        switch v6 {
                                        case "\n\t":
                                            var v29 int16 = -868
                                            _ = v29
                                        case "":
                                            v1 = ("hello" + v6)
                                        default:
                                            var v30 float64 = -2.25
                                            _ = v30
                                        }
                                        switch ((v1 + (v1 + ((v1 + "a") + v1))) + (v1 + "foo bar")) {
                                        case "hello":
                                            switch ("hello" + v1) {
                                            case "hello":
                                                v31 := GenS4{F0: -106, F1: 0, F2: -112}
                                                _ = v31
                                            case "foo bar":
                                                switch v1 {
                                                case "\n\t":
                                                    v1 = ("" + "foo bar")
                                                default:
                                                    var v32 uint16 = 0
                                                    _ = v32
                                                }
                                                if (!(!(-1 != (fi19 & 32768)))) {
                                                    v3[(^(fi17))] = false
                                                    v18 -= (v18 + (^(v18)))
                                                    if false {
                                                        v6 = ((v6 + "hello") + ("a" + ("" + (v1 + (v1 + ((v6 + v1) + (v1 + ("\n\t" + v6))))))))
                                                        var v33 uint16 = 69
                                                        v34 := GenS6{F0: 1, F1: -32768, F2: "a"}
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        var v35 uint16 = 65535
                                                        fmt.Println(fi19, fi17)
                                                        v18 = (v18 >> ((4294967295) & 63))
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v35
                                                    }
                                                } else {
                                                    if ((583 < 42) || false) {
                                                        switch "foo bar" {
                                                        case "a":
                                                            switch "hello" {
                                                            case "hello":
                                                                var v36 int32 = 539
                                                                _ = v36
                                                            default:
                                                                v6 = ("\n\t" + ("foo bar" + ""))
                                                            }
                                                        default:
                                                            v18 = v18
                                                        }
                                                    }
                                                    v18 += (v18 % ((128) | 1))
                                                }
                                            case "":
                                                switch 32768 {
                                                case 32768:
                                                    v18 -= (v18 - 214)
                                                case 987:
                                                    v1 = "foo bar"
                                                case 0:
                                                    v1 = "a"
                                                default:
                                                    switch (fi19 | fi19) {
                                                    case -397:
                                                        var v37 uint8 = (v18 % ((2) | 1))
                                                        _ = v37
                                                    case -1:
                                                        _ = len(v4)
                                                    }
                                                }
                                            default:
                                                if (true || ((v18 > (v18 >> ((65535) & 63))) && (((0.1 + 0.1) == 1e10) || true))) {
                                                    if (true && false) {
                                                        var v38 int = (^(fi17))
                                                        _ = v38
                                                    }
                                                } else {
                                                    var v39 uint32 = 65535
                                                    _ = v39
                                                }
                                            }
                                        case "foo bar":
                                            if false {
                                                if true {
                                                    delete(v3, fi19)
                                                    fmt.Println(fi17, fi17, v1)
                                                    v6 = v6
                                                } else {
                                                    var v40 int8 = 0
                                                    v1 = (("foo bar" + v1) + "\n\t")
                                                    _ = v40
                                                }
                                                v41 := v3[(fi17 >> ((18446744073709551615) & 63))]
                                                _ = v41
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = (fi17 ^ (^(fi17)))
                                                }
                                                v42 := GenS7{F0: 529}
                                                v1 = "a"
                                                v43 := GenS1{F0: 2, F1: -32768, F2: 2147483647}
                                                if (-780 == 757) {
                                                    fmt.Println(v6, v18, v18)
                                                    if (129 < 432) {
                                                        fmt.Println(v6)
                                                        var v44 string = v6
                                                        _ = v44
                                                    } else {
                                                        fmt.Println(fi17, fi19)
                                                        if false {
                                                            fmt.Println(v18)
                                                            delete(v3, 9223372036854775807)
                                                            switch (^(fi17)) {
                                                            case 9223372036854775807:
                                                                if (!(false && true)) {
                                                                    v1 = (v1 + v6)
                                                                    var v45 uint = 188
                                                                    _ = v45
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v1 = v6
                                                    fmt.Println(fi19, fi19)
                                                    _ = len(v3)
                                                    if (-1 < 9223372036854775807) {
                                                        v6 = (v6 + "a")
                                                        var v46 int8 = -128
                                                        v46 -= -124
                                                        _ = v46
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = 4294967295
                                                        }
                                                        v6 = "a"
                                                    }
                                                } else {
                                                    if (32767 > 255) {
                                                        var v47 uint64 = 33
                                                        v6 = (v6 + (((("hello" + "foo bar") + "foo bar") + (v1 + "hello")) + v1))
                                                        _ = v47
                                                    } else {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            if true {
                                                                v18 = 127
                                                                var v48 float32 = 1.5
                                                                fmt.Println(v48, fi19)
                                                                _ = v48
                                                            }
                                                        case "a":
                                                            v49 := GenS8{F0: (fi19 | fi19), F1: 1000}
                                                            _ = v49
                                                        case "hello":
                                                            switch "foo bar" {
                                                            case "":
                                                                if true {
                                                                    v1 = (v6 + "\n\t")
                                                                }
                                                            case "\n\t":
                                                                var v50 uint8 = v18
                                                                _ = v50
                                                            case "a":
                                                                var v51 uint8 = v18
                                                                _ = v51
                                                            default:
                                                                fmt.Println(fi19, fi17, v6)
                                                            }
                                                        }
                                                    }
                                                    v6 = ((v1 + ("\n\t" + (v1 + (v6 + "hello")))) + v6)
                                                    _ = len(v2)
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                        }
                                        v52 := GenS6{F0: 100, F1: -450, F2: ""}
                                        _ = v52
                                    }
                                    v53 := v3[-204]
                                    v54 := GenS9{F0: v18}
                                    _ = v53
                                    _ = v54
                                } else {
                                    fmt.Println(v18, v18, v18)
                                }
                                v55 := GenS10{F0: true, F1: -709, F2: 4294967295}
                                if false {
                                    if true {
                                        v1 = v1
                                    }
                                    switch "hello" {
                                    case "hello":
                                        v1 = (("\n\t" + v6) + ((("\n\t" + v6) + v6) + v1))
                                    }
                                }
                                _ = v55
                            }
                        } else {
                            switch fi17 {
                            case 127:
                                var v56 uint64 = 18446744073709551615
                                _ = v56
                            }
                            v6 = (("hello" + v6) + v1)
                            v57 := GenS7{F0: 42}
                            if (1 != 2147483647) {
                                var v58 uint = 0
                                v3[fi19] = true
                                _ = v58
                            }
                            v59 := GenS9{F0: 127}
                            v6 = v6
                            _ = v57
                            _ = v59
                        }
                        var v60 uint16 = 0
                        v60 = (v60 >> ((65535) & 63))
                        _ = v60
                    }
                    var v61 int32 = 42
                    v62 := GenS2{F0: 255}
                    _ = v61
                    _ = v62
                }
                if (100 <= 15) {
                    delete(v3, fi17)
                } else {
                    delete(v3, (fi19 - fi17))
                    v63 := GenS4{F0: 100, F1: -1, F2: 42}
                    _ = v63
                }
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
    case 9223372036854775807:
        v64 := GenS11{F0: 9223372036854775807}
        _ = v64
    }
    fmt.Println(v6)
    v6 = "\n\t"
    if (!((115 != 0) || true)) {
        v6 = v1
    }
    var v65 int16 = -328
    fmt.Println(v65)
    for fi66 := 0; fi66 < 12; fi66++ {
        v67 := GenS4{F0: -128, F1: -9223372036854775808, F2: -409}
        _ = fi66
        _ = v67
    }
    var v68 float64 = (3.14159 * (0.0 + (0.1 - 1.5)))
    v69 := GenS6{F0: 42, F1: (v65 * v65), F2: v6}
    if true {
        for fi70 := 0; fi70 < 14; fi70++ {
            if (755 != 95) {
                for fi71 := 0; fi71 < 7; fi71++ {
                    fmt.Println(v65, fi70, v6)
                    v5 = append(v5, 1)
                    _ = fi71
                }
                v68 = v68
            }
            _ = fi70
        }
        var v72 int16 = (^(v65))
        _ = v72
    } else {
        var v73 string = ""
        var v74 int32 = -753
        _ = v73
        _ = v74
    }
    switch -2147483648 {
    case 4294967295:
        if false {
            v75 := GenS12{F0: v1, F1: 255, F2: 100}
            if true {
                v76 := GenS8{F0: 0, F1: (v65 / ((48) | 1))}
                v77 := GenS13{F0: -898, F1: 0, F2: 888}
                _ = v76
                _ = v77
            } else {
                v78 := v3[-611]
                for fi79 := 0; fi79 < 12; fi79++ {
                    v5 = append(v5, 42)
                    if v78 {
                        fmt.Println(fi79)
                        v80 := GenS9{F0: 42}
                        v1 = v6
                        _ = v80
                    }
                    _ = fi79
                }
                _ = v78
            }
            _ = v75
        } else {
            var v81 string = ("foo bar" + (v1 + ("\n\t" + v1)))
            if true {
                var v82 int16 = (-(v65))
                v83 := GenS14{F0: 385, F1: 1, F2: -156}
                v65 += v65
                _ = v82
                _ = v83
            }
            _ = v81
        }
    case 42:
        v84 := GenS13{F0: -331, F1: 118, F2: 1}
        _ = v84
    case 1:
        var v85 int64 = 1
        _ = v85
    default:
        v6 = "foo bar"
    }
    v65 += v65
    var v86 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v68
    _ = v69
    _ = v86
}
