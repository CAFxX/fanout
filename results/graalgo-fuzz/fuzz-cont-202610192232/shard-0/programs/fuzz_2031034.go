package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 uint32; F2 uint64 }

func main() {
    v1 := []uint32{42, 558}
    var v2 uint8 = 42
    var v3 string = "\n\t"
    var v4 uint8 = (^(v2))
    v5 := make(map[string]string)
    v6 := GenS1{F0: "", F1: 42}
    v3 = "a"
    var v7 float32 = 1.5
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS2{F0: fi8, F1: 128, F2: "\n\t"}
        var v10 int64 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 -= (v2 | v4)
    switch 0 {
    case 42:
        if (301 == 1000) {
            v11 := GenS2{F0: 9223372036854775807, F1: 255, F2: v3}
            var v12 uint = 569
            if true {
                switch v3 {
                case "hello":
                    v12 = v12
                }
            } else {
                var v13 int64 = 0
                _ = v13
            }
            _ = v11
            _ = v12
        } else {
            v3 = ("hello" + ((v3 + v3) + "foo bar"))
        }
    case 2147483647:
        var v14 int16 = -258
        _ = v14
    case 903:
        var v15 int32 = 1
        _ = v15
    default:
        v5[((v3 + ("foo bar" + ("foo bar" + ((((("" + (v3 + v3)) + v3) + "a") + v3) + "a")))) + v3)] = v3
    }
    if len(v1) > 0 {
        v1[0] = 2
    }
    var v16 uint32 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v5)
    v2 = (^(v4))
    if false {
        _ = len(v1)
        switch 4294967295 {
        case 592:
            for fi17 := 0; fi17 < 6; fi17++ {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                fmt.Println(v16)
                _ = fi17
            }
        case 0:
            v7 = v7
        }
    } else {
        if (!false) {
            v18 := GenS3{F0: 100, F1: 1, F2: v4}
            if true {
                var v19 int16 = -1
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 uint64 = 4294967295
                    v7 += ((v7 + v7) + v7)
                    _ = len(v1)
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                v16 = 71
                v23 := GenS4{F0: fi22}
                var v24 int = 32768
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = v18
        }
        v7 -= v7
        v3 = v3
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            v4 += (v4 << ((42) & 63))
            switch (fi26 % (((fi26 | -51)) | 1)) {
            case 128:
                fmt.Println(v2, fi26, v2)
            case 2147483647:
                if (-2147483648 <= -1) {
                    v27 := GenS1{F0: v3, F1: 2147483647}
                    _ = v27
                }
            default:
                v5[("a" + ("foo bar" + v3))] = "\n\t"
            }
            _ = fi26
        }
        var v28 uint16 = 391
        for fi29 := 0; fi29 < 17; fi29++ {
            fmt.Println(fi29)
            v30 := GenS3{F0: 0, F1: 2, F2: v4}
            switch (fi29 | 42) {
            case 876:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            delete(v5, (v3 + "hello"))
            _ = fi29
            _ = v30
        }
        _ = fi25
        _ = v28
    }
    if false {
        if false {
            fmt.Println(v16, v4, v2)
        }
        for fi31 := 0; fi31 < 7; fi31++ {
            v16 = uint32(v4)
            if ((float64(-2.25) != ((-2.25 + float64(v7)) - (-2.25 + 3.14159))) && (true && true)) {
                for fi32 := 0; fi32 < 18; fi32++ {
                    if false {
                        var v33 uint32 = (^(v16))
                        var v34 string = v3
                        var v35 uint64 = 0
                        fmt.Println(v2, v3)
                        var v36 uint16 = 42
                        var v37 int32 = 1
                        if false {
                            switch "\n\t" {
                            case "foo bar":
                                var v38 int8 = -86
                                _ = v38
                            default:
                                var v39 string = (((v3 + "foo bar") + v3) + v3)
                                _ = v39
                            }
                            v4 = v4
                            if (-1 != (v37 >> ((42) & 63))) {
                                v40 := GenS5{F0: v2}
                                delete(v5, "\n\t")
                                _ = v40
                            } else {
                                var v41 int16 = int16(v33)
                                fmt.Println(v33, fi32, v36)
                                v34 = "a"
                                _ = v41
                            }
                        }
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    var v42 int32 = -2147483648
                    v43 := GenS6{F0: -505, F1: v42}
                    switch fi32 {
                    case 32767:
                        v4 = 126
                    default:
                        v42 = v42
                    }
                    v44 := GenS7{F0: (v3 + ((v3 + "foo bar") + v3))}
                    var v45 bool = (!true)
                    _ = fi32
                    _ = v42
                    _ = v43
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 11; fi46++ {
                    if false {
                        v1 = append(v1, 2147483648)
                    }
                    v5[(v3 + (v3 + v3))] = ("a" + "hello")
                    _ = fi46
                }
            } else {
                fmt.Println(v2)
            }
            _ = fi31
        }
    }
    if (-1 < 42) {
        for fi47 := 0; fi47 < 20; fi47++ {
            delete(v5, v3)
            for fi48 := 0; fi48 < 7; fi48++ {
                switch fi48 {
                case 734:
                    switch v3 {
                    case "":
                        v49 := v5[""]
                        _ = v49
                    case "hello":
                        v50 := GenS1{F0: ((v3 + v3) + v3), F1: 292}
                        if true {
                            var v51 float64 = 1.5
                            if (-1 >= -205) {
                                if true {
                                    var v52 uint = 18446744073709551615
                                    _ = v52
                                }
                                var v53 string = v3
                                _ = v53
                            }
                            if (!true) {
                                v7 = (3.14159 - 1.5)
                                v54 := GenS6{F0: 0, F1: 0}
                                var v55 float32 = float32(1.5)
                                _ = v54
                                _ = v55
                            }
                            _ = v51
                        }
                        v7 += ((1e10 + (v7 - (v7 + 1.5))) - v7)
                        _ = v50
                    case "foo bar":
                        if false {
                            var v56 bool = true
                            switch v3 {
                            case "foo bar":
                                var v57 uint8 = (v4 + 255)
                                _ = v57
                            }
                            _ = v56
                        } else {
                            var v58 float32 = 0.0
                            var v59 uint64 = 239
                            fmt.Println(v7, v7)
                            _ = v58
                            _ = v59
                        }
                    default:
                        if true {
                            v7 = v7
                        } else {
                            var v60 bool = (!false)
                            v16 -= (v16 + 42)
                            var v61 int8 = -43
                            _ = v60
                            _ = v61
                        }
                    }
                case 0:
                    if ((!true) && (((!(!false)) || (-660 <= -691)) || true)) {
                        var v62 int16 = 42
                        var v63 bool = (false && false)
                        var v64 uint32 = v16
                        switch 784 {
                        case 9223372036854775807:
                            v2 -= v4
                        default:
                            fmt.Println(v3)
                        }
                        if false {
                            var v65 bool = (2147483647 >= 1)
                            fmt.Println(v64, fi48, fi47)
                            var v66 int16 = (v62 % ((v62) | 1))
                            _ = v65
                            _ = v66
                        } else {
                            v4 = v4
                            var v67 float64 = 1e10
                            _ = v67
                        }
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    fmt.Println(fi47, v4, v4)
                    _ = len(v1)
                    delete(v5, v3)
                case 440:
                    v16 = v16
                default:
                    v5["foo bar"] = "hello"
                }
                var v68 uint = 255
                _ = fi48
                _ = v68
            }
            _ = fi47
        }
        for fi69 := 0; fi69 < 10; fi69++ {
            for fi70 := 0; fi70 < 17; fi70++ {
                var v71 float64 = (1e10 + 3.14159)
                _ = fi70
                _ = v71
            }
            delete(v5, v3)
            fmt.Println(v2, v7, v2)
            _ = fi69
        }
        switch "\n\t" {
        case "a":
            if false {
                v72 := GenS8{F0: 255}
                for fi73 := 0; fi73 < 1; fi73++ {
                    fmt.Println(v2, fi73, v16)
                    if false {
                        fmt.Println(v2)
                    }
                    _ = fi73
                }
                if (!((!false) && true)) {
                    if (true || ((0 > 18446744073709551615) || ((!false) && true))) {
                        for fi74 := 0; fi74 < 14; fi74++ {
                            v4 += (v2 / ((v4) | 1))
                            fmt.Println(v16)
                            if (-752 != 42) {
                                fmt.Println(v4, fi74)
                                if true {
                                    v16 -= (^(v16))
                                    v75 := GenS9{F0: false, F1: 2, F2: 18446744073709551615}
                                    for fi76 := 0; fi76 < 3; fi76++ {
                                        var v77 string = v3
                                        var v78 uint16 = uint16(v2)
                                        _ = fi76
                                        _ = v77
                                        _ = v78
                                    }
                                    _ = v75
                                } else {
                                    var v79 uint8 = (v4 | (^(v2)))
                                    var v80 uint32 = v16
                                    fmt.Println(v79)
                                    _ = v79
                                    _ = v80
                                }
                                fmt.Println(v2, v2)
                            } else {
                                fmt.Println(v3, fi74)
                            }
                            _ = fi74
                        }
                    } else {
                        var v81 uint64 = 5
                        var v82 int8 = 100
                        v81 = v81
                        _ = v81
                        _ = v82
                    }
                    v4 = (v4 % ((v4) | 1))
                }
                _ = v72
            }
        case "foo bar":
            var v83 int64 = -9223372036854775808
            _ = v83
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
