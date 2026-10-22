package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int16; F2 uint32 }
type GenS6 struct { F0 int16; F1 int32; F2 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 uint32; F2 int }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 uint32; F1 uint8 }
type GenS11 struct { F0 bool }

func main() {
    var v1 int32 = -19
    var v2 float32 = (float32((3.14159 + 1e10)) * (1.5 - (0.0 / 1.0)))
    v3 := make(map[uint32]uint32)
    switch 0 {
    case 2147483647:
        for fi4 := 0; fi4 < 20; fi4++ {
            if false {
                for fi5 := 0; fi5 < 7; fi5++ {
                    var v6 int16 = int16(v1)
                    _ = fi5
                    _ = v6
                }
                v7 := GenS1{F0: -9223372036854775808, F1: 255}
                v1 -= v1
                _ = v7
            } else {
                for fi8 := 0; fi8 < 1; fi8++ {
                    if (-109 == 1) {
                        v2 = v2
                        v1 += v1
                    }
                    _ = fi8
                }
                for fi9 := 0; fi9 < 20; fi9++ {
                    v1 = -597
                    v3[2] = 570
                    _ = fi9
                }
                var v10 uint8 = 1
                _ = v10
            }
            _ = fi4
        }
    case 32767:
        v11 := GenS2{F0: 42}
        _ = v11
    }
    var v12 uint = 432
    v12 = (^(v12))
    for fi13 := 0; fi13 < 19; fi13++ {
        if false {
            switch 629 {
            case 1:
                for fi14 := 0; fi14 < 8; fi14++ {
                    v12 = v12
                    if false {
                        fmt.Println(fi13, v12, v12)
                    }
                    v15 := GenS3{F0: (v1 >> ((18446744073709551615) & 63)), F1: -63}
                    _ = fi14
                    _ = v15
                }
            case 690:
                v16 := GenS2{F0: 2}
                _ = v16
            }
            v17 := []int{(fi13 | (fi13 / (((fi13 >> ((42) & 63))) | 1))), (fi13 | 32767), fi13}
            var v18 int64 = 1
            _ = v17
            _ = v18
        }
        v19 := GenS4{F0: 9223372036854775807}
        var v20 uint64 = 18446744073709551615
        _ = fi13
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        if true {
            v2 = v2
        } else {
            var v22 float64 = 1e10
            var v23 float64 = 3.14159
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 string = ("foo bar" + ("\n\t" + "\n\t"))
            v26 := []int{fi24}
            var v27 bool = false
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        switch fi21 {
        case 371:
            v28 := GenS5{F0: (fi21 & fi21), F1: 42, F2: 42}
            _ = v28
        case 0:
            v12 -= 216
        }
        _ = fi21
    }
    v29 := []int{859}
    if true {
        if ((0 >= 673) && (false && (false && (false || (640 <= 2147483647))))) {
            v30 := GenS3{F0: (v1 | -2147483648), F1: 42}
            _ = v30
        }
        delete(v3, 2147483648)
        fmt.Println(v1, v12, v2)
    } else {
        if true {
            if false {
                switch "hello" {
                case "a":
                    for fi31 := 0; fi31 < 2; fi31++ {
                        v12 = (v12 % ((v12) | 1))
                        v32 := GenS5{F0: -157, F1: -769, F2: 65535}
                        var v33 int64 = -945
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                case "hello":
                    var v34 string = (((("\n\t" + ("hello" + ("a" + ""))) + "hello") + "") + (("hello" + "hello") + "hello"))
                    _ = v34
                case "\n\t":
                    for fi35 := 0; fi35 < 1; fi35++ {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            var v37 uint16 = 868
                            var v38 uint32 = 472
                            v39 := GenS1{F0: 1, F1: 42}
                            var v40 bool = ((v37 + (v37 * v37)) >= (v37 | v37))
                            v3[(v38 % ((2147483647) | 1))] = v38
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 8; fi41++ {
                            var v42 int16 = 32767
                            _ = fi41
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 6; fi43++ {
                            if true {
                                if false {
                                    switch ("\n\t" + "") {
                                    case "foo bar":
                                        v12 -= v12
                                    case "":
                                        var v44 int8 = 42
                                        _ = v44
                                    case "a":
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v1)
                                        case "foo bar":
                                            var v45 float32 = v2
                                            _ = v45
                                        case "":
                                            v2 = v2
                                        }
                                    default:
                                        v1 = -1
                                    }
                                    if ((v2 * v2) <= float32(0.0)) {
                                        var v46 int = (-(fi35))
                                        _ = v46
                                    } else {
                                        if false {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v47 := GenS6{F0: -32768, F1: v1, F2: 0}
                                                _ = v47
                                            case "":
                                                if false {
                                                    if len(v29) >= 2 {
                                                        v29 = v29[1:len(v29)]
                                                    }
                                                    if (!(fi43 == 128)) {
                                                        switch "foo bar" {
                                                        case "":
                                                            fmt.Println(v1, fi43, fi43)
                                                        case "a":
                                                            fmt.Println(v12, v12, fi43)
                                                        case "hello":
                                                            v48 := GenS3{F0: v1, F1: 0}
                                                            _ = v48
                                                        }
                                                    }
                                                    v1 -= 2147483647
                                                }
                                            case "hello":
                                                if true {
                                                    if false {
                                                        v2 = (v2 / v2)
                                                        var v49 int32 = v1
                                                        var v50 uint = 255
                                                        v50 += v50
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        v51 := GenS4{F0: -1}
                                                        v12 = 51
                                                        if false {
                                                            var v52 int32 = (v1 / (((v1 % (((v1 << ((v12) & 63))) | 1))) | 1))
                                                            var v53 int8 = -64
                                                            var v54 uint32 = 859
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                        } else {
                                                            v12 = v12
                                                        }
                                                        _ = v51
                                                    }
                                                } else {
                                                    if len(v29) > 0 {
                                                        v29[0] = 127
                                                    }
                                                }
                                            }
                                            fmt.Println(fi43)
                                        }
                                    }
                                    v55 := GenS7{F0: 127}
                                    _ = v55
                                }
                                var v56 uint8 = 1
                                _ = v56
                            }
                            var v57 float32 = 3.14159
                            fmt.Println(v57, v12)
                            v1 -= 0
                            _ = fi43
                            _ = v57
                        }
                        _ = fi35
                    }
                }
                fmt.Println(v12, v1)
                fmt.Println(v2, v1)
            } else {
                for fi58 := 0; fi58 < 17; fi58++ {
                    if false {
                        v59 := GenS6{F0: 42, F1: v1, F2: int64(fi58)}
                        fmt.Println(v1, v12, v2)
                        v1 = v1
                        _ = v59
                    }
                    var v60 uint16 = 42
                    _ = fi58
                    _ = v60
                }
            }
            if (1 != 0) {
                v2 = -2.25
            } else {
                v1 -= 355
                for fi61 := 0; fi61 < 11; fi61++ {
                    if (!true) {
                        switch -575 {
                        case 1:
                            v1 = v1
                        }
                        var v62 uint = v12
                        _ = v62
                    } else {
                        if (float32(-2.25) == 1.5) {
                            fmt.Println(v2, v2, v12)
                            var v63 float32 = ((v2 / v2) * float32(1.5))
                            _ = v63
                        }
                        if (!false) {
                            v64 := GenS8{F0: uint8(v12), F1: 42, F2: (fi61 & fi61)}
                            _ = v64
                        } else {
                            for fi65 := 0; fi65 < 14; fi65++ {
                                v12 += 1
                                _ = len(v29)
                                var v66 int8 = -128
                                var v67 string = ""
                                fmt.Println(fi65, fi65, v66)
                                v12 -= (v12 >> ((v12) & 63))
                                _ = fi65
                                _ = v66
                                _ = v67
                            }
                        }
                        var v68 uint = v12
                        _ = v68
                    }
                    if false {
                        switch ("foo bar" + (("" + "foo bar") + "foo bar")) {
                        case "\n\t":
                            if ((-128 >= -54) || true) {
                                if false {
                                    if false {
                                        var v69 string = ("foo bar" + (((("foo bar" + "foo bar") + "hello") + "a") + "hello"))
                                        fmt.Println(fi61)
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        _ = v69
                                    } else {
                                        fmt.Println(fi61)
                                    }
                                    for fi70 := 0; fi70 < 19; fi70++ {
                                        switch "" {
                                        case "foo bar":
                                            v71 := GenS3{F0: 1, F1: -22}
                                            v1 = v1
                                            _ = v71
                                        default:
                                            if (true || false) {
                                                _ = len(v3)
                                            }
                                        }
                                        v72 := GenS8{F0: 1, F1: 485, F2: (fi70 / ((4294967295) | 1))}
                                        if (!(!(459 != 716))) {
                                            v12 -= (v12 >> ((v12) & 63))
                                            v73 := v3[2]
                                            v3[(v73 & (^(v73)))] = 102
                                            _ = v73
                                        }
                                        _ = fi70
                                        _ = v72
                                    }
                                    for fi74 := 0; fi74 < 18; fi74++ {
                                        v75 := GenS9{F0: 26, F1: "\n\t"}
                                        fmt.Println(v1)
                                        var v76 uint = (v12 * (v12 + (^(v12))))
                                        _ = fi74
                                        _ = v75
                                        _ = v76
                                    }
                                }
                            }
                        }
                        var v77 bool = (!true)
                        _ = v77
                    } else {
                        for fi78 := 0; fi78 < 10; fi78++ {
                            var v79 uint32 = 490
                            _ = fi78
                            _ = v79
                        }
                        var v80 string = "hello"
                        _ = v80
                    }
                    if false {
                        v81 := GenS10{F0: 4294967295, F1: 3}
                        switch 0 {
                        case 0:
                            var v82 string = ("hello" + ("a" + "foo bar"))
                            _ = v82
                        case 127:
                            v1 -= v1
                        case 32768:
                            fmt.Println(v12, v2)
                        }
                        if ((true && false) && false) {
                            v1 = 2147483647
                        } else {
                            for fi83 := 0; fi83 < 16; fi83++ {
                                fmt.Println(v1, v1)
                                v84 := GenS1{F0: 0, F1: 0}
                                var v85 uint32 = 4294967295
                                _ = fi83
                                _ = v84
                                _ = v85
                            }
                        }
                        _ = v81
                    }
                    _ = fi61
                }
            }
        } else {
            v86 := GenS5{F0: -332, F1: -778, F2: 42}
            v12 -= (v12 << ((1) & 63))
            v1 += v1
            _ = v86
        }
    }
    switch "" {
    case "foo bar":
        v2 -= 1e10
    case "":
        v2 = 1.5
    default:
        if false {
            v3[4294967295] = 0
        }
    }
    v12 = v12
    v12 -= (v12 * v12)
    for fi87 := 0; fi87 < 13; fi87++ {
        v88 := GenS11{F0: false}
        fmt.Println(v2)
        _ = fi87
        _ = v88
    }
    fmt.Println(v2)
    var v89 uint = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v89:", v89)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v89
}
