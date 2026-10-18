package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 string; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int16; F1 int64; F2 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16; F1 bool }

func main() {
    var v1 float64 = -2.25
    var v2 int = 42
    var v3 uint32 = 42
    var v4 float64 = v1
    var v5 uint64 = 4294967295
    v6 := []int{(v2 | -2147483648), v2}
    v7 := GenS1{F0: false}
    var v8 uint64 = 1
    v2 -= (v2 - v2)
    v3 = v3
    v1 -= float64(1.5)
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "foo bar" {
        case "a":
            v10 := GenS2{F0: 65535}
            _ = v10
        case "foo bar":
            var v11 float32 = (0.1 + 3.14159)
            _ = v11
        default:
            var v12 uint64 = v8
            _ = v12
        }
        var v13 int32 = -2147483648
        _ = fi9
        _ = v13
    }
    fmt.Println(v2)
    v2 = 225
    v6 = append(v6, (-(v2)))
    if ((((3.14159 + v4) >= v1) && false) || true) {
        v14 := GenS2{F0: 2}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch "a" {
        case "a":
            var v15 uint32 = v3
            _ = v15
        case "foo bar":
            var v16 uint64 = v8
            _ = v16
        case "":
            v6 = append(v6, v2)
        }
        _ = v14
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS2{F0: 42}
            var v19 int16 = -32768
            if (-152 < 9223372036854775807) {
                v6 = append(v6, fi17)
            } else {
                if true {
                    fmt.Println(v2)
                } else {
                    switch ("foo bar" + "hello") {
                    case "a":
                        switch "hello" {
                        case "a":
                            var v20 uint16 = 694
                            _ = v20
                        case "hello":
                            switch fi17 {
                            case 32767:
                                if false {
                                    var v21 int = v2
                                    _ = v21
                                } else {
                                    if true {
                                        v3 -= 2147483648
                                        if true {
                                            if (1e10 < float64((float32((v4 / v4)) * 3.14159))) {
                                                var v22 string = "a"
                                                _ = v22
                                            }
                                            fmt.Println(v19, v19)
                                            v19 = 0
                                        }
                                        fmt.Println(v5, v5, v19)
                                    } else {
                                        switch v2 {
                                        case 41:
                                            for fi23 := 0; fi23 < 20; fi23++ {
                                                v19 -= v19
                                                v5 = v5
                                                v24 := GenS3{F0: 0, F1: "\n\t", F2: 107}
                                                _ = fi23
                                                _ = v24
                                            }
                                        }
                                        v4 = float64((3.14159 * 3.14159))
                                        if true {
                                            if false {
                                                if (false || (!false)) {
                                                    for fi25 := 0; fi25 < 8; fi25++ {
                                                        switch (fi17 * fi25) {
                                                        case -9223372036854775808:
                                                            if true {
                                                                fmt.Println(v19, fi25)
                                                                fmt.Println(fi17, v8)
                                                                if true {
                                                                    v5 -= 0
                                                                } else {
                                                                    v26 := GenS3{F0: 128, F1: "hello", F2: 834}
                                                                    fmt.Println(v2, v1, fi17)
                                                                    _ = v26
                                                                }
                                                                v3 = v3
                                                                var v27 uint = 65535
                                                                _ = v27
                                                            } else {
                                                                fmt.Println(v2, v5)
                                                            }
                                                        case 204:
                                                            var v28 int8 = -5
                                                            _ = v28
                                                        default:
                                                            switch "a" {
                                                            case "\n\t":
                                                                fmt.Println(v19)
                                                            }
                                                        }
                                                        _ = fi25
                                                    }
                                                }
                                                var v29 uint32 = (v3 * 1)
                                                _ = v29
                                            }
                                        }
                                    }
                                    if (((v3 & v3) != 65535) || true) {
                                        for fi30 := 0; fi30 < 5; fi30++ {
                                            var v31 int64 = 400
                                            _ = fi30
                                            _ = v31
                                        }
                                    }
                                }
                            case 189:
                                var v32 int = v2
                                _ = v32
                            case 595:
                                var v33 uint8 = 127
                                _ = v33
                            default:
                                v34 := GenS3{F0: 42, F1: "foo bar", F2: 255}
                                _ = v34
                            }
                        default:
                            v6 = append(v6, 936)
                        }
                    case "foo bar":
                        for fi35 := 0; fi35 < 15; fi35++ {
                            v8 -= v8
                            if len(v6) > 0 {
                                v6[0] = 128
                            }
                            _ = fi35
                        }
                    default:
                        _ = len(v6)
                    }
                    for fi36 := 0; fi36 < 17; fi36++ {
                        if true {
                            v19 += 32767
                        } else {
                            switch "hello" {
                            case "":
                                var v37 uint64 = (^(v8))
                                _ = v37
                            case "a":
                                switch "foo bar" {
                                case "":
                                    if false {
                                        var v38 uint16 = 777
                                        _ = v38
                                    } else {
                                        if (true && true) {
                                            if (!(1 >= 255)) {
                                                var v39 string = "hello"
                                                _ = v39
                                            }
                                        }
                                        v40 := GenS2{F0: 108}
                                        var v41 uint16 = 32767
                                        if (("" + (("foo bar" + (("hello" + ("hello" + "a")) + "foo bar")) + "")) >= ((("hello" + "hello") + "hello") + "")) {
                                            if ((v41 - (v41 << ((854) & 63))) > (^(v41))) {
                                                var v42 int32 = 0
                                                _ = v42
                                            } else {
                                                v19 = (-(v19))
                                                _ = len(v6)
                                                var v43 int8 = 100
                                                _ = v43
                                            }
                                        }
                                        _ = v40
                                        _ = v41
                                    }
                                    v1 = 1e10
                                    v2 = (fi17 + 4294967295)
                                    v1 -= ((1e10 * (v4 / (float64(float32(1e10)) - v4))) + ((3.14159 / 0.1) + (float64((0.1 + -2.25)) + v4)))
                                    fmt.Println(v19, v5, v3)
                                case "foo bar":
                                    fmt.Println(v8, fi36)
                                case "hello":
                                    v19 -= (^(v19))
                                default:
                                    v44 := GenS4{F0: ((("\n\t" + ("foo bar" + "\n\t")) + "\n\t") + "\n\t")}
                                    _ = v44
                                }
                            case "\n\t":
                                var v45 uint8 = uint8(v5)
                                _ = v45
                            default:
                                if (768 <= 9223372036854775807) {
                                    v19 -= (v19 << ((4294967295) & 63))
                                    var v46 float32 = (3.14159 + 1.5)
                                    _ = v46
                                } else {
                                    v2 -= (^(fi17))
                                    if ((v5 | (v8 * (v5 / ((347) | 1)))) < 0) {
                                        fmt.Println(fi17, v3)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v5 = (v8 % ((298) | 1))
                                        switch "" {
                                        case "a":
                                            v1 = ((-2.25 * v4) * v1)
                                        case "hello":
                                            if (42 == v19) {
                                                v47 := GenS5{F0: v3, F1: (v5 >> ((65) & 63))}
                                                _ = v47
                                            }
                                        case "":
                                            var v48 uint8 = 42
                                            _ = v48
                                        }
                                        v4 -= v4
                                    }
                                    var v49 int64 = -9223372036854775808
                                    _ = v49
                                }
                            }
                            v4 -= v4
                            switch ("" + ("" + "a")) {
                            case "a":
                                _ = len(v6)
                                if (("foo bar" == ("\n\t" + "a")) || false) {
                                    v3 -= v3
                                    _ = len(v6)
                                    v8 = (v8 << ((1) & 63))
                                } else {
                                    var v50 float64 = v1
                                    if true {
                                        switch "" {
                                        case "":
                                            fmt.Println(fi17)
                                        case "hello":
                                            v6 = append(v6, v2)
                                        case "foo bar":
                                            if true {
                                                var v51 int = -182
                                                _ = v51
                                            }
                                        default:
                                            v19 += 1000
                                        }
                                        var v52 float64 = (v50 + v1)
                                        var v53 uint64 = 349
                                        v5 += uint64(v3)
                                        fmt.Println(v50, v53, v52)
                                        _ = v52
                                        _ = v53
                                    }
                                    v4 -= float64(0.1)
                                    v54 := GenS6{F0: v19, F1: (-(fi17))}
                                    _ = v50
                                    _ = v54
                                }
                                v55 := GenS7{F0: v19, F1: -323, F2: v8}
                                _ = v55
                            case "foo bar":
                                v56 := GenS8{F0: "\n\t"}
                                _ = v56
                            }
                        }
                        fmt.Println(fi36, v3)
                        _ = fi36
                    }
                    v4 += (float64(1.5) * 0.0)
                }
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        if (!(((-23 == 1) || false) || false)) {
            switch -580 {
            case 2147483647:
                if true {
                    var v57 int64 = 9223372036854775807
                    _ = v57
                } else {
                    var v58 uint = 683
                    var v59 int32 = -745
                    _ = v58
                    _ = v59
                }
            case 32767:
                for fi60 := 0; fi60 < 16; fi60++ {
                    var v61 uint8 = 200
                    if len(v6) > 0 {
                        v6[0] = (^(v2))
                    }
                    fmt.Println(v4, v4, fi60)
                    _ = fi60
                    _ = v61
                }
            case -1:
                for fi62 := 0; fi62 < 11; fi62++ {
                    for fi63 := 0; fi63 < 5; fi63++ {
                        var v64 uint16 = 65535
                        _ = fi63
                        _ = v64
                    }
                    v65 := GenS8{F0: ("a" + "foo bar")}
                    _ = fi62
                    _ = v65
                }
            default:
                v66 := GenS9{F0: 32767, F1: (2 >= 891)}
                _ = v66
            }
        } else {
            for fi67 := 0; fi67 < 2; fi67++ {
                if true {
                    v1 = v4
                    var v68 uint32 = v3
                    var v69 string = ("a" + ("hello" + "a"))
                    _ = v68
                    _ = v69
                }
                _ = fi67
            }
            v2 = (v2 << ((42) & 63))
            for fi70 := 0; fi70 < 4; fi70++ {
                var v71 float64 = v4
                _ = fi70
                _ = v71
            }
        }
        v72 := GenS7{F0: 526, F1: 267, F2: v8}
        _ = v72
    }
    v73 := GenS7{F0: 0, F1: 940, F2: 1}
    for fi74 := 0; fi74 < 11; fi74++ {
        if false {
            v3 -= (v3 & v3)
            v2 = (fi74 & v2)
        }
        _ = fi74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v73
    _ = v8
}
