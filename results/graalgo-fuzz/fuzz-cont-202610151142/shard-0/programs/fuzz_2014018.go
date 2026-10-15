package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint16; F1 uint16; F2 bool }
type GenS6 struct { F0 int; F1 int32; F2 int32 }

func main() {
    v1 := []string{("foo bar" + "foo bar"), "hello", "", "foo bar"}
    var v2 uint8 = 2
    var v3 int = 4294967295
    v4 := make(map[uint32]uint32)
    if len(v1) > 0 {
        v1[0] = (("hello" + ("\n\t" + "foo bar")) + "hello")
    }
    var v5 bool = false
    delete(v4, 128)
    var v6 int = (v3 & (v3 >> ((1) & 63)))
    fmt.Println(v5, v5, v2)
    fmt.Println(v6, v6)
    v6 -= 0
    var v7 uint8 = (^(v2))
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint8 = v2
        switch (v3 & v3) {
        case 32767:
            if ("a" < "") {
                v2 = 0
            } else {
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 uint8 = v7
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v12 float32 = (0.0 / 1e10)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: v7, F1: -2147483648, F2: (v9 >> ((1) & 63))}
                _ = v13
            }
        case -771:
            v6 = -83
        case -2147483648:
            if false {
                var v14 uint = 255
                if v5 {
                    v15 := GenS2{F0: -987, F1: (v14 & 1), F2: (v2 >> ((42) & 63))}
                    var v16 int16 = -74
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 int8 = -48
                        switch 32768 {
                        case 32768:
                            var v19 int = 9223372036854775807
                            _ = v19
                        case -2147483648:
                            if v5 {
                                v7 = (v9 ^ 42)
                                v20 := GenS2{F0: v16, F1: 0, F2: (v2 ^ v2)}
                                v16 -= -1
                                _ = v20
                            } else {
                                if v5 {
                                    var v21 float64 = 3.14159
                                    v9 = 255
                                    v7 = v9
                                    _ = v21
                                } else {
                                    if false {
                                        v2 = 42
                                        switch (-(v6)) {
                                        case -165:
                                            if (2 == 91) {
                                                var v22 bool = v5
                                                var v23 float32 = -2.25
                                                v18 = (-(v18))
                                                _ = v22
                                                _ = v23
                                            } else {
                                                var v24 int16 = v16
                                                _ = len(v1)
                                                v25 := GenS1{F0: v2, F1: 560, F2: 42}
                                                _ = v24
                                                _ = v25
                                            }
                                        }
                                    }
                                }
                                _ = len(v4)
                            }
                        case -1:
                            v16 = v16
                            v16 = 32767
                            var v26 int32 = -1
                            _ = v26
                        }
                        v27 := GenS2{F0: (-(v16)), F1: 558, F2: 42}
                        _ = fi17
                        _ = v18
                        _ = v27
                    }
                    _ = v15
                    _ = v16
                } else {
                    v28 := GenS2{F0: int16(v2), F1: v14, F2: 1}
                    switch ((("\n\t" + "a") + ((("\n\t" + "a") + ("hello" + ("foo bar" + ("foo bar" + "foo bar")))) + ("foo bar" + "foo bar"))) + "a") {
                    case "":
                        v5 = (-1 <= 0)
                    case "\n\t":
                        fmt.Println(v2)
                    case "foo bar":
                        v3 = fi8
                    }
                    if true {
                        v4[828] = uint32(v9)
                        if ((v7 % ((v7) | 1)) <= v7) {
                            v29 := v4[4294967295]
                            _ = v29
                        }
                    }
                    _ = v28
                }
                var v30 string = ""
                _ = v14
                _ = v30
            }
        }
        v5 = (false || true)
        _ = fi8
        _ = v9
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        switch ("hello" + "hello") {
        case "":
            v6 -= 4294967295
        case "foo bar":
            fmt.Println(v6, fi31)
        }
        var v32 uint16 = 423
        for fi33 := 0; fi33 < 9; fi33++ {
            if v5 {
                var v34 bool = (2 <= 4294967295)
                switch 0 {
                case 127:
                    v35 := GenS3{F0: 295, F1: (fi33 % (((fi31 % ((fi31) | 1))) | 1)), F2: 2147483648}
                    _ = v35
                case 9223372036854775807:
                    v6 -= (fi33 / ((2147483647) | 1))
                case -9223372036854775808:
                    if false {
                        if (v34 && v5) {
                            if (uint32(v7) <= 2147483648) {
                                v36 := GenS4{F0: 0, F1: fi33}
                                v32 += 42
                                switch ("a" + (("" + "") + "a")) {
                                case "":
                                    fmt.Println(fi31)
                                }
                                _ = v36
                            } else {
                                v2 = v2
                                var v37 string = ("a" + ("" + "\n\t"))
                                switch 4294967295 {
                                case 9223372036854775807:
                                    v38 := v4[65535]
                                    _ = v38
                                case -2147483648:
                                    var v39 string = "foo bar"
                                    _ = v39
                                case 4294967295:
                                    v40 := GenS5{F0: (v32 & (v32 << ((uint(v7)) & 63))), F1: v32, F2: ((((1.5 - (1.5 * 0.1)) - (float32(0.1) + 0.0)) / -2.25) == -2.25)}
                                    _ = v40
                                default:
                                    v5 = (-32768 > -1)
                                }
                                switch "a" {
                                case "foo bar":
                                    if (4294967295 >= 562) {
                                        v41 := GenS6{F0: v6, F1: -9, F2: -1}
                                        fmt.Println(v2, v2, v2)
                                        var v42 int64 = 9223372036854775807
                                        var v43 uint8 = v2
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                default:
                                    var v44 float32 = (((-2.25 - (1e10 * 3.14159)) / 1e10) + (0.1 / 1.0))
                                    _ = v44
                                }
                                var v45 int = 32768
                                _ = v37
                                _ = v45
                            }
                            switch (v3 / ((v6) | 1)) {
                            case 4294967295:
                                var v46 int16 = 311
                                _ = v46
                            case 653:
                                fmt.Println(v7, v6)
                            case 32767:
                                var v47 uint32 = 1
                                _ = v47
                            }
                            if v34 {
                                v3 = (v3 % (((v3 & 1)) | 1))
                            }
                        }
                        if v5 {
                            fmt.Println(v2, v5)
                            v7 += (^(v2))
                        } else {
                            var v48 float64 = (0.0 / float64(3.14159))
                            switch "a" {
                            case "hello":
                                v34 = (true || (true || false))
                                var v49 float64 = v48
                                var v50 int = (fi31 + (-(fi31)))
                                _ = v49
                                _ = v50
                            case "\n\t":
                                v5 = v5
                            case "foo bar":
                                if (false && (!v5)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    var v51 uint64 = 1
                                    _ = v51
                                }
                            default:
                                var v52 int64 = -9223372036854775808
                                _ = v52
                            }
                            _ = v48
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        v34 = (1 == 42)
                        delete(v4, 42)
                    }
                default:
                    fmt.Println(v34, fi33)
                }
                v6 = (-(v3))
                _ = v34
            } else {
                var v53 int32 = 2147483647
                _ = v53
            }
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    if v5 {
        v54 := v4[2]
        _ = v54
    } else {
        for fi55 := 0; fi55 < 13; fi55++ {
            switch "hello" {
            case "hello":
                fmt.Println(v6, v7)
            case "foo bar":
                v2 -= v7
            case "":
                for fi56 := 0; fi56 < 4; fi56++ {
                    v7 += (v2 | v7)
                    if (false || ("\n\t" >= ((("" + ("hello" + "a")) + "\n\t") + ""))) {
                        v3 = fi55
                        if true {
                            if len(v1) > 0 {
                                v1[0] = (("a" + "foo bar") + "hello")
                            }
                            delete(v4, 549)
                            v57 := GenS5{F0: uint16(v7), F1: 637, F2: (60 >= -1)}
                            _ = v57
                        }
                    }
                    v58 := GenS6{F0: fi56, F1: 0, F2: int32(fi55)}
                    v7 = (^(v2))
                    v2 = 127
                    _ = fi56
                    _ = v58
                }
            default:
                var v59 uint32 = 2
                _ = v59
            }
            v60 := GenS1{F0: v7, F1: 0, F2: (v2 << ((1) & 63))}
            _ = fi55
            _ = v60
        }
    }
    switch "\n\t" {
    case "a":
        var v61 int8 = 42
        _ = v61
    case "hello":
        v5 = v5
    case "foo bar":
        var v62 int = (v3 & (v6 ^ (^(v3))))
        _ = v62
    default:
        var v63 int64 = -1
        _ = v63
    }
    fmt.Println(v2, v7, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
