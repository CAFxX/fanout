package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint; F1 int; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint32; F2 int }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 uint16; F1 int }

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]string)
    var v3 uint8 = 255
    v4 := []string{"\n\t", (("a" + "") + "\n\t"), "foo bar"}
    v3 -= 247
    v5 := GenS1{F0: 32767}
    switch "\n\t" {
    case "hello":
        if true {
            var v6 uint64 = 1
            _ = v6
        } else {
            v3 += 0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            if (false && (-9223372036854775808 <= 9223372036854775807)) {
                var v8 uint = (v1 * (^(v1)))
                v9 := GenS1{F0: 255}
                _ = v8
                _ = v9
            } else {
                var v10 uint = v1
                _ = v10
            }
            fmt.Println(v3, v1)
            delete(v2, 65535)
        } else {
            var v11 uint = (v1 ^ v1)
            _ = v11
        }
        _ = fi7
    }
    switch "hello" {
    case "":
        var v12 uint16 = 650
        _ = v12
    default:
        v3 -= v3
    }
    v13 := GenS2{F0: 32767}
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 string = ""
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: 32767}
            v1 += (v1 % (((^(v1))) | 1))
            v1 = v1
            v1 = (v1 ^ v1)
            var v18 int16 = 1
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            var v19 float64 = (float64((float32(1e10) / 1.0)) - (((-2.25 + 0.1) * 1.5) + 0.0))
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    var v20 int64 = -9223372036854775808
    var v21 int32 = 1
    var v22 float32 = 0.1
    var v23 uint16 = 2
    for fi24 := 0; fi24 < 18; fi24++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            v22 = float32(1.5)
            if (-128 >= 42) {
                if true {
                    var v26 uint64 = 4294967295
                    var v27 int8 = 73
                    v27 -= v27
                    _ = v26
                    _ = v27
                } else {
                    var v28 int16 = -886
                    fmt.Println(v28)
                    v1 += 18446744073709551615
                    switch fi24 {
                    case 127:
                        v21 = v21
                    case -2147483648:
                        v29 := GenS3{F0: v28, F1: ""}
                        _ = v29
                    case -1:
                        v3 = 127
                    default:
                        if ((v23 ^ (v23 - (^(v23)))) >= (v23 << ((v1) & 63))) {
                            if true {
                                if true {
                                    fmt.Println(fi25)
                                    var v30 int32 = (v21 / ((42) | 1))
                                    if false {
                                        v21 -= 1
                                        v31 := GenS4{F0: 4294967295, F1: (^(fi25)), F2: 100}
                                        var v32 int32 = v30
                                        var v33 uint64 = 810
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v30
                                } else {
                                    v34 := GenS2{F0: v28}
                                    if len(v4) > 0 {
                                        v4[0] = (("" + "a") + "a")
                                    }
                                    if (!((float64((((1.5 / v22) * ((v22 + (0.0 / (1.5 * 1e10))) + v22)) + float32(1e10))) - float64((((((v22 + 0.1) - v22) + ((3.14159 + v22) + 0.1)) - v22) + float32(0.0)))) >= (float64(float32((3.14159 - float64(v22)))) - 0.1))) {
                                        v35 := GenS4{F0: v1, F1: fi24, F2: -128}
                                        _ = v35
                                    }
                                    _ = v34
                                }
                                var v36 float32 = -2.25
                                switch -9223372036854775808 {
                                case 466:
                                    if false {
                                        v36 = 3.14159
                                        var v37 float32 = 0.0
                                        fmt.Println(v37, v28, v23)
                                        _ = v37
                                    } else {
                                        delete(v2, 32)
                                    }
                                case -1:
                                    var v38 int32 = v21
                                    _ = v38
                                }
                                _ = v36
                            }
                            var v39 int = (fi25 / ((fi25) | 1))
                            if (190 == 2147483648) {
                                switch "" {
                                case "foo bar":
                                    if false {
                                        if false {
                                            _ = len(v4)
                                            if false {
                                                var v40 int8 = 19
                                                if (v20 < v20) {
                                                    if (true || false) {
                                                        v23 = (v23 | v23)
                                                        switch "a" {
                                                        case "hello":
                                                            v22 += v22
                                                        default:
                                                            var v41 uint64 = 1
                                                            _ = v41
                                                        }
                                                        v42 := GenS5{F0: 4294967295, F1: 4294967295}
                                                        _ = v42
                                                    }
                                                    var v43 float32 = (v22 + v22)
                                                    _ = v43
                                                }
                                                _ = v40
                                            }
                                        } else {
                                            var v44 uint32 = 4294967295
                                            switch (fi25 ^ (fi24 - fi25)) {
                                            case 884:
                                                _ = len(v2)
                                            case 4294967295:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch fi24 {
                                                    case -47:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v45 := GenS6{F0: 232, F1: 42, F2: v3}
                                                            _ = v45
                                                        case "hello":
                                                            switch ("" + "a") {
                                                            case "":
                                                                _ = len(v2)
                                                            default:
                                                                if true {
                                                                    if true {
                                                                        var v46 int16 = (v28 * 0)
                                                                        _ = v46
                                                                    }
                                                                    if (!true) {
                                                                        v21 = 0
                                                                    } else {
                                                                        v21 = v21
                                                                        v28 += v28
                                                                    }
                                                                }
                                                                v28 -= (-(v28))
                                                                var v47 float64 = 0.0
                                                                v2[v44] = ("\n\t" + "")
                                                                v4 = append(v4, "foo bar")
                                                                if true {
                                                                    switch (fi25 ^ 9223372036854775807) {
                                                                    case 223:
                                                                        var v48 uint16 = 1
                                                                        _ = v48
                                                                    }
                                                                    if true {
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            fmt.Println(v44, v44)
                                                                        default:
                                                                            if true {
                                                                                v47 += -2.25
                                                                            }
                                                                        }
                                                                        fmt.Println(v39)
                                                                    } else {
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                                _ = v47
                                                            }
                                                        case "foo bar":
                                                            if (true && ((fi24 & fi25) < (fi24 / ((-2147483648) | 1)))) {
                                                                switch fi25 {
                                                                case 165:
                                                                    v49 := GenS6{F0: 0, F1: 18446744073709551615, F2: v3}
                                                                    _ = v49
                                                                case -828:
                                                                    v4 = append(v4, "a")
                                                                case 930:
                                                                    _ = len(v2)
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v3 -= v3
                                                                    default:
                                                                        if false {
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v50 bool = (!((fi24 | 32768) > v39))
                                                                                _ = v50
                                                                            case "hello":
                                                                                v3 = (v3 >> ((v1) & 63))
                                                                            }
                                                                        } else {
                                                                            v51 := GenS1{F0: (v23 - 42)}
                                                                            fmt.Println(v20)
                                                                            if (v20 != v20) {
                                                                                v44 += v44
                                                                                v4 = append(v4, ("foo bar" + ""))
                                                                            } else {
                                                                                var v52 int16 = v28
                                                                                _ = v52
                                                                            }
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                    switch ("hello" + "hello") {
                                                                    case "":
                                                                        v20 = v20
                                                                    }
                                                                default:
                                                                    _ = len(v2)
                                                                    var v53 uint64 = 2
                                                                    _ = v53
                                                                }
                                                                var v54 uint = v1
                                                                switch 42 {
                                                                case -9223372036854775808:
                                                                    if (!false) {
                                                                        v55 := GenS3{F0: (v28 - v28), F1: ("\n\t" + "hello")}
                                                                        _ = v55
                                                                    }
                                                                default:
                                                                    fmt.Println(v22, v1, v39)
                                                                }
                                                                _ = v54
                                                            }
                                                        default:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    case -309:
                                                        var v56 int16 = -647
                                                        _ = v56
                                                    case 9223372036854775807:
                                                        switch "foo bar" {
                                                        case "a":
                                                            v39 -= 128
                                                        default:
                                                            if ((!false) || (!(2147483647 != fi25))) {
                                                                v23 = 264
                                                            }
                                                        }
                                                    default:
                                                        var v57 float64 = (((((1.5 * (0.0 * ((0.1 + (1e10 / 1.0)) / 0.1))) * (1.5 * float64(1e10))) / 0.1) / 1e10) + (((0.1 + 0.0) * 0.0) - 3.14159))
                                                        _ = v57
                                                    }
                                                case "hello":
                                                    var v58 uint = v1
                                                    _ = v58
                                                }
                                            case -2147483648:
                                                v59 := GenS4{F0: v1, F1: 32768, F2: -22}
                                                _ = v59
                                            }
                                            _ = v44
                                        }
                                    } else {
                                        if (v28 < (v28 - (v28 | 42))) {
                                            fmt.Println(v39)
                                        }
                                    }
                                default:
                                    v60 := GenS7{F0: 390, F1: 113, F2: (v39 + (fi25 >> ((v1) & 63)))}
                                    _ = v60
                                }
                            }
                            if (false || false) {
                                fmt.Println(v39, v23, fi24)
                                v61 := GenS8{F0: (v3 << ((v1) & 63)), F1: 2}
                                v4 = append(v4, "hello")
                                _ = v61
                            } else {
                                if (!((^(v20)) == 65)) {
                                    fmt.Println(v23, v28, v21)
                                    if (!(true && (uint16(v3) >= (v23 >> ((v1) & 63))))) {
                                        v62 := GenS8{F0: v3, F1: 4294967295}
                                        v63 := GenS3{F0: (v28 >> (((v1 & v1)) & 63)), F1: "a"}
                                        _ = v62
                                        _ = v63
                                    } else {
                                        if false {
                                            var v64 int32 = (v21 ^ 320)
                                            fmt.Println(v39)
                                            _ = v64
                                        }
                                        var v65 float64 = (3.14159 / 1.0)
                                        var v66 uint64 = 18446744073709551615
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 uint = 455
                                    if false {
                                        fmt.Println(v3)
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = ("a" + ("a" + ("hello" + "")))
                                    }
                                    _ = v67
                                }
                            }
                            _ = v39
                        }
                    }
                    switch 4294967295 {
                    case 32767:
                        if true {
                            switch "foo bar" {
                            case "foo bar":
                                v68 := GenS9{F0: v23, F1: fi25}
                                _ = v68
                            case "a":
                                v28 += v28
                            case "hello":
                                var v69 uint64 = 746
                                _ = v69
                            }
                            var v70 uint = v1
                            var v71 int32 = v21
                            _ = v70
                            _ = v71
                        } else {
                            var v72 int16 = v28
                            var v73 int = -9223372036854775808
                            _ = v72
                            _ = v73
                        }
                    case -9223372036854775808:
                        fmt.Println(v28, v28)
                    default:
                        v3 = v3
                    }
                    _ = v28
                }
            }
            var v74 uint64 = 533
            _ = fi25
            _ = v74
        }
        _ = len(v4)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
