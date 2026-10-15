package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 string; F1 bool; F2 int }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint64; F2 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int }
func genFn0(p0 uint) int64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 350, F1: 1}
        v7 := GenS1{F0: 65535, F1: 494}
        var v8 int16 = 32767
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        p0 = 0
        v10 := []int{(fi9 / ((fi9) | 1)), fi9, 959}
        var v11 int64 = -9223372036854775808
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return -359
}

func main() {
    var v1 uint8 = 127
    v2 := []uint32{1, 2147483647}
    var v3 float32 = 1e10
    var v4 uint8 = (v1 << ((65535) & 63))
    fmt.Println(v1, v4, v1)
    if (42 > 65535) {
        switch (("foo bar" + "a") + "hello") {
        case "a":
            v1 = (v4 ^ v4)
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v12 int = 9223372036854775807
        if (!(1 <= 42)) {
            if ((!(!((4294967295 == 1) || (!true)))) && true) {
                v13 := GenS1{F0: 1, F1: -2147483648}
                var v14 bool = (79 <= -32768)
                switch v12 {
                case 944:
                    var v15 uint = 130
                    _ = v15
                case -36:
                    var v16 int32 = -1
                    _ = v16
                default:
                    var v17 int = 0
                    _ = v17
                }
                _ = v13
                _ = v14
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := GenS1{F0: 374, F1: (fi19 / (((^(v12))) | 1))}
                    if (!(((v4 != 42) && true) && true)) {
                        var v21 float32 = (v3 - v3)
                        if true {
                            v12 = 32767
                            if true {
                                switch fi19 {
                                case 0:
                                    v21 = (v21 * v21)
                                }
                            }
                        } else {
                            fmt.Println(fi19, v1)
                        }
                        v21 += v21
                        _ = v21
                    } else {
                        v1 -= 166
                        v4 -= 0
                        var v22 uint32 = 0
                        switch ("a" + "foo bar") {
                        case "a":
                            v23 := GenS2{F0: 762, F1: 255}
                            var v24 int16 = 1000
                            v25 := GenS3{F0: 1000, F1: 748}
                            _ = v23
                            _ = v24
                            _ = v25
                        case "foo bar":
                            var v26 uint = 4294967295
                            _ = v26
                        default:
                            v12 -= v12
                        }
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v12
    } else {
        if true {
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS3{F0: -35, F1: 1}
                v3 = float32(1.5)
                _ = fi27
                _ = v28
            }
            if ((false && ((((v3 - v3) + 1.5) - v3) > 0.1)) || false) {
                switch "\n\t" {
                case "foo bar":
                    var v29 int32 = -926
                    _ = v29
                case "a":
                    v30 := GenS1{F0: 918, F1: 841}
                    _ = v30
                default:
                    switch "\n\t" {
                    case "\n\t":
                        v3 = v3
                    default:
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                    }
                }
            }
        }
        v4 = (v4 / ((v1) | 1))
    }
    switch -457 {
    case -2147483648:
        v3 -= ((0.0 + (v3 / 1.0)) / v3)
    default:
        v31 := GenS4{F0: "foo bar", F1: true, F2: 2147483647}
        _ = v31
    }
    fmt.Println(v1, v3)
    switch "foo bar" {
    case "foo bar":
        if (2147483647 < 4294967295) {
            fmt.Println(v4, v3, v3)
            var v32 float64 = ((float64((1e10 / v3)) / 0.1) - 1.5)
            v3 = float32(v32)
            _ = v32
        } else {
            v33 := GenS5{F0: 930, F1: uint(v4)}
            v4 -= (v1 << ((4294967295) & 63))
            _ = v33
        }
    case "":
        if true {
            var v34 uint = 2
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v35 := GenS3{F0: 1000, F1: 301}
            _ = v34
            _ = v35
        } else {
            v1 -= (^(v4))
            _ = len(v2)
        }
    }
    var v36 float32 = v3
    v37 := GenS6{F0: 2147483647}
    switch -9223372036854775808 {
    case 32768:
        for fi38 := 0; fi38 < 12; fi38++ {
            v36 += (v3 + float32(float64(float32(1.5))))
            v3 = ((v36 * v3) * (0.0 - 0.0))
            if (1e10 <= 1.5) {
                for fi39 := 0; fi39 < 14; fi39++ {
                    fmt.Println(v1, fi39, v4)
                    v2 = append(v2, uint32(v4))
                    if true {
                        if (814 == 2) {
                            fmt.Println(v3)
                            fmt.Println(v1, fi38)
                        } else {
                            switch "\n\t" {
                            case "hello":
                                var v40 uint64 = 253
                                _ = v40
                            case "\n\t":
                                var v41 int8 = int8(fi39)
                                _ = v41
                            }
                            v3 += (v3 / (v3 - v3))
                        }
                        v4 += (v1 & (v4 / (((^(v1))) | 1)))
                        if (true && true) {
                            v36 -= (1.5 + v3)
                            v36 += (v36 * v36)
                        } else {
                            var v42 int = fi39
                            _ = v42
                        }
                        switch fi39 {
                        case 42:
                            var v43 bool = ((!true) || false)
                            _ = v43
                        case -2147483648:
                            _ = len(v2)
                        case -9223372036854775808:
                            var v44 int16 = 1
                            _ = v44
                        default:
                            var v45 uint64 = 1
                            _ = v45
                        }
                    }
                    _ = fi39
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    var v46 int32 = -2147483648
                    var v47 bool = ((!(1 >= 0)) || true)
                    if false {
                        for fi48 := 0; fi48 < 12; fi48++ {
                            v49 := GenS7{F0: 349, F1: 72}
                            var v50 int16 = 902
                            fmt.Println(v3, v3)
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                    } else {
                        if true {
                            fmt.Println(v4, v36)
                            if (!(65535 > 2)) {
                                v47 = v47
                                var v51 uint64 = 835
                                _ = v51
                            } else {
                                switch -654 {
                                case 32767:
                                    switch "\n\t" {
                                    case "":
                                        for fi52 := 0; fi52 < 17; fi52++ {
                                            if (!((1e10 / float32(((0.0 + 1e10) / 1.0))) > v3)) {
                                                var v53 int32 = (v46 << ((18446744073709551615) & 63))
                                                _ = v53
                                            } else {
                                                v54 := GenS8{F0: genFn0(2), F1: 42, F2: fi38}
                                                v4 += 119
                                                fmt.Println(fi52)
                                                v47 = v47
                                                switch (^(fi38)) {
                                                case 42:
                                                    v55 := GenS9{F0: "\n\t"}
                                                    _ = v55
                                                case -9223372036854775808:
                                                    fmt.Println(v46, fi38, v1)
                                                default:
                                                    v56 := GenS7{F0: 2147483648, F1: 2}
                                                    _ = v56
                                                }
                                                if v47 {
                                                    _ = len(v2)
                                                    _ = len(v2)
                                                    if len(v2) > 0 {
                                                        v2[0] = 0
                                                    }
                                                } else {
                                                    v2 = append(v2, 65535)
                                                    if len(v2) > 0 {
                                                        v2[0] = 2147483647
                                                    }
                                                    var v57 float64 = float64(1e10)
                                                    _ = v57
                                                }
                                                fmt.Println(fi52, v36)
                                                _ = v54
                                            }
                                            v58 := GenS10{F0: fi38}
                                            _ = fi52
                                            _ = v58
                                        }
                                    case "hello":
                                        var v59 uint64 = 1
                                        _ = v59
                                    case "\n\t":
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                    }
                                case 42:
                                    for fi60 := 0; fi60 < 15; fi60++ {
                                        switch fi38 {
                                        case 0:
                                            v46 += v46
                                        case -9223372036854775808:
                                            v47 = (2 >= 65535)
                                        }
                                        _ = fi60
                                    }
                                }
                                for fi61 := 0; fi61 < 17; fi61++ {
                                    switch (("hello" + ("a" + ("\n\t" + ("hello" + ("\n\t" + ("foo bar" + "a")))))) + (("\n\t" + "hello") + ("a" + "hello"))) {
                                    case "":
                                        if (true && true) {
                                            v4 = v1
                                            fmt.Println(v4, v46, v47)
                                            var v62 uint32 = 2
                                            _ = v62
                                        } else {
                                            switch fi61 {
                                            case 378:
                                                var v63 float64 = (0.0 - (-2.25 + ((1e10 + -2.25) * 3.14159)))
                                                _ = v63
                                            default:
                                                if (!(!true)) {
                                                    switch (fi38 % (((fi61 ^ fi61)) | 1)) {
                                                    case -468:
                                                        var v64 uint64 = 239
                                                        _ = v64
                                                    default:
                                                        v1 += v1
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = 2
                                                    }
                                                } else {
                                                    var v65 uint64 = 42
                                                    var v66 float32 = (((float32(1e10) * 1e10) / v3) - (v3 * v3))
                                                    if (v47 && v47) {
                                                        if (1 == 642) {
                                                            v3 += float32((float64(((0.0 + (3.14159 * v36)) + (3.14159 - (((float32(3.14159) / 1.0) / (v66 + 1e10)) / float32(float64(v66)))))) / (0.0 + 3.14159)))
                                                        }
                                                        fmt.Println(v4, v46)
                                                        switch (fi38 % ((-49) | 1)) {
                                                        case 42:
                                                            if (42 >= 65535) {
                                                                if v47 {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 4294967295
                                                                    }
                                                                }
                                                                var v67 int16 = -289
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v67
                                                            }
                                                            fmt.Println(v1, v46)
                                                        case 32767:
                                                            fmt.Println(fi61, v46)
                                                        case 2147483647:
                                                            var v68 float64 = (float64(float32((0.1 * (3.14159 + (((0.0 * -2.25) * float64(v66)) / 3.14159))))) + 3.14159)
                                                            _ = v68
                                                        default:
                                                            fmt.Println(fi38, v1, v4)
                                                        }
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                }
                                            }
                                        }
                                    }
                                    _ = fi61
                                }
                                v4 = v4
                            }
                        }
                    }
                    _ = v46
                    _ = v47
                }
            } else {
                var v69 float32 = 0.1
                fmt.Println(v1, fi38, v4)
                var v70 float64 = 1e10
                _ = v69
                _ = v70
            }
            _ = fi38
        }
    case 2147483647:
        var v71 int = 32768
        _ = v71
    case 0:
        v36 += 3.14159
    default:
        if false {
            var v72 uint32 = 2
            _ = v72
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
