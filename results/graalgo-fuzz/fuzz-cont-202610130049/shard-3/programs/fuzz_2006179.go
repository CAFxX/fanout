package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 0
    v3 := []string{"foo bar", ("a" + ""), "\n\t"}
    var v4 int16 = 235
    if true {
        var v5 float64 = (0.1 + 0.1)
        if false {
            var v6 uint8 = v2
            if true {
                fmt.Println(v6, v2, v1)
                v4 = v4
            }
            fmt.Println(v4)
            _ = v6
        }
        if ((!((!true) || true)) && (false && false)) {
            for fi7 := 0; fi7 < 14; fi7++ {
                switch (fi7 % ((127) | 1)) {
                case 4294967295:
                    for fi8 := 0; fi8 < 12; fi8++ {
                        var v9 uint = 1
                        _ = fi8
                        _ = v9
                    }
                }
                switch ("a" + (("a" + "\n\t") + "\n\t")) {
                case "a":
                    var v10 uint32 = 42
                    _ = v10
                case "":
                    v3 = append(v3, "\n\t")
                case "hello":
                    v11 := GenS1{F0: 0, F1: 1, F2: (v2 + uint8(v1))}
                    _ = v11
                }
                if (!(false || (true || (255 > 1)))) {
                    switch (-(fi7)) {
                    case 0:
                        v4 += v4
                    case -1:
                        v4 = (v4 ^ -472)
                    case 32768:
                        v3 = append(v3, ("" + "foo bar"))
                    default:
                        v5 += 1.5
                    }
                    var v12 uint = 4294967295
                    _ = v12
                }
                _ = fi7
            }
            if true {
                v4 = v4
                v4 = v4
            } else {
                fmt.Println(v4, v2)
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                if (!((!((0 == 0) || (!(-1 < 1)))) && false)) {
                    fmt.Println(v5, v4, v1)
                    for fi14 := 0; fi14 < 2; fi14++ {
                        _ = len(v3)
                        _ = fi14
                    }
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                } else {
                    if (false || true) {
                        fmt.Println(v4, v1)
                    }
                    for fi15 := 0; fi15 < 11; fi15++ {
                        fmt.Println(fi15, v5)
                        _ = fi15
                    }
                }
                switch (^(fi13)) {
                case 9223372036854775807:
                    var v16 float64 = v5
                    _ = v16
                case 985:
                    v2 = (v2 << ((18446744073709551615) & 63))
                case 4294967295:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        var v18 bool = false
                        _ = fi17
                        _ = v18
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi13
            }
        } else {
            _ = len(v3)
        }
        _ = v5
    }
    if false {
        v19 := GenS1{F0: 772, F1: 100, F2: (v2 | v2)}
        fmt.Println(v1, v2, v4)
        _ = v19
    } else {
        v4 += v4
        fmt.Println(v4)
    }
    v20 := GenS1{F0: 290, F1: 118, F2: (^(v2))}
    for fi21 := 0; fi21 < 20; fi21++ {
        fmt.Println(v2, fi21, v2)
        var v22 uint16 = v1
        if false {
            switch "foo bar" {
            case "a":
                if true {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v2 = (v2 + 255)
                        switch (fi21 % (((-(fi23))) | 1)) {
                        case 2147483647:
                            v1 = (v22 + v22)
                        default:
                            switch (fi23 / (((fi21 * (fi21 << ((526) & 63)))) | 1)) {
                            case 127:
                                if false {
                                    v22 += (^(v22))
                                    v1 += 255
                                }
                                var v24 uint = 947
                                _ = v24
                            default:
                                var v25 int16 = (-(v4))
                                _ = v25
                            }
                        }
                        _ = fi23
                    }
                    v3 = append(v3, ((("foo bar" + "foo bar") + "hello") + "hello"))
                    for fi26 := 0; fi26 < 9; fi26++ {
                        if (!true) {
                            _ = len(v3)
                            if (-495 >= 12) {
                                if ((50 > 0) || (!false)) {
                                    fmt.Println(fi26)
                                    switch "\n\t" {
                                    case "":
                                        if (!true) {
                                            v22 = (v22 ^ 585)
                                            v1 = v22
                                            v27 := GenS1{F0: 42, F1: 0, F2: (v2 - (v2 - 0))}
                                            _ = v27
                                        }
                                    case "\n\t":
                                        var v28 uint = 1
                                        _ = v28
                                    case "hello":
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                    }
                                    if false {
                                        if true {
                                            v29 := GenS2{F0: 642}
                                            var v30 int8 = -1
                                            switch int(v22) {
                                            case 1:
                                                v31 := GenS2{F0: 554}
                                                _ = v31
                                            default:
                                                var v32 uint16 = 2
                                                _ = v32
                                            }
                                            _ = v29
                                            _ = v30
                                        } else {
                                            v33 := GenS1{F0: 42, F1: 42, F2: 128}
                                            _ = v33
                                        }
                                    }
                                } else {
                                    v34 := GenS3{F0: 42, F1: -1, F2: 18446744073709551615}
                                    v35 := GenS4{F0: (v4 | v4)}
                                    if false {
                                        v2 = (v2 - v2)
                                        v36 := GenS5{F0: ("" + (("a" + "a") + "hello")), F1: 317}
                                        _ = v36
                                    }
                                    _ = v34
                                    _ = v35
                                }
                                v4 -= (v4 * 1000)
                                v22 = (^(v22))
                                switch fi21 {
                                case 0:
                                    v37 := GenS6{F0: 850}
                                    _ = v37
                                case -1:
                                    switch (fi26 % ((fi26) | 1)) {
                                    case -2147483648:
                                        v22 = v22
                                    case -1:
                                        var v38 float64 = float64((0.1 - 0.0))
                                        _ = v38
                                    case 127:
                                        if (-128 > -100) {
                                            var v39 uint8 = 180
                                            _ = v39
                                        }
                                    }
                                case 42:
                                    if false {
                                        v3 = append(v3, ("\n\t" + ("a" + ("foo bar" + ("" + ("a" + (("foo bar" + ("a" + "\n\t")) + "a")))))))
                                        var v40 uint32 = 306
                                        v2 += (v2 >> ((18446744073709551615) & 63))
                                        _ = v40
                                    } else {
                                        v41 := GenS1{F0: 65535, F1: 127, F2: (v2 % (((v2 | v2)) | 1))}
                                        v42 := GenS4{F0: (v4 & -327)}
                                        fmt.Println(v22, fi26, v1)
                                        var v43 uint64 = 18446744073709551615
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                }
                            } else {
                                var v44 int = 953
                                v45 := GenS4{F0: (^(v4))}
                                v46 := GenS2{F0: 469}
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        } else {
                            var v47 int64 = 1
                            var v48 int64 = -9223372036854775808
                            v1 = (v1 + v1)
                            _ = v47
                            _ = v48
                        }
                        _ = fi26
                    }
                }
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = ("" + "foo bar")
                }
            default:
                if (3.14159 > (1e10 - -2.25)) {
                    v1 -= (v22 - (v22 ^ v1))
                }
            }
            v1 = (v22 + (v22 / ((0) | 1)))
            v4 = v4
        } else {
            var v49 int32 = 0
            var v50 uint64 = 42
            _ = v49
            _ = v50
        }
        _ = fi21
        _ = v22
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v51 int64 = 470
        _ = v51
    case -158:
        v4 += (-(v4))
    case -9223372036854775808:
        if false {
            v52 := GenS5{F0: ("\n\t" + "\n\t"), F1: 42}
            if (128 >= -277) {
                fmt.Println(v4, v4, v4)
                v53 := GenS5{F0: "", F1: 705}
                _ = v53
            }
            _ = v52
        } else {
            var v54 bool = (!false)
            if v54 {
                var v55 uint8 = (v2 / ((1) | 1))
                fmt.Println(v55)
                _ = v55
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v54
        }
    }
    if (("\n\t" + ("a" + "\n\t")) == "foo bar") {
        if (((1e10 - 3.14159) * 1.5) == 0.0) {
            var v56 uint = 65535
            _ = v56
        } else {
            var v57 int8 = -47
            _ = v57
        }
        fmt.Println(v2, v1)
    }
    switch "a" {
    case "a":
        if true {
            v1 = 880
            var v58 int64 = -1
            if ((-832 > 9223372036854775807) || (42 == 18446744073709551615)) {
                for fi59 := 0; fi59 < 11; fi59++ {
                    fmt.Println(v58)
                    _ = fi59
                }
            } else {
                v60 := GenS6{F0: v1}
                if ((v1 % (((v1 * v1)) | 1)) > uint16(v2)) {
                    fmt.Println(v4, v58)
                    v61 := GenS1{F0: 42, F1: -128, F2: v2}
                    var v62 uint8 = v2
                    _ = v61
                    _ = v62
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v63 float32 = 1.5
                    _ = v63
                }
                _ = v60
            }
            _ = v58
        }
    case "\n\t":
        var v64 uint64 = 1
        _ = v64
    default:
        v65 := GenS3{F0: 527, F1: 0, F2: 0}
        _ = v65
    }
    for fi66 := 0; fi66 < 20; fi66++ {
        fmt.Println(v2, v4)
        _ = fi66
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
