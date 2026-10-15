package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 int16 }
type GenS6 struct { F0 int8; F1 bool; F2 int }
func genFn0() bool {
    v5 := GenS1{F0: 9223372036854775807, F1: -2147483648, F2: 5}
    var v6 uint32 = 296
    v7 := GenS1{F0: 9223372036854775807, F1: -2147483648, F2: 127}
    _ = v5
    _ = v6
    _ = v7
    return (false && true)
}

func main() {
    var v1 int32 = 493
    var v2 uint = 1
    var v3 uint64 = 0
    v4 := []string{}
    v1 = v1
    v1 = v1
    switch "a" {
    case "foo bar":
        if (false || ((v2 != 4294967295) || true)) {
            var v8 int64 = 1
            _ = v8
        }
    }
    var v9 string = (("" + "hello") + ("\n\t" + "\n\t"))
    if (4294967295 != (v3 ^ v3)) {
        switch "" {
        case "hello":
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 uint = (^(v2))
                v12 := GenS1{F0: fi10, F1: fi10, F2: 42}
                if true {
                    var v13 string = v9
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        fmt.Println(v9, v2)
                        _ = fi14
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        case "\n\t":
            v1 -= v1
        default:
            v2 -= (v2 + 42)
        }
        var v15 uint16 = 133
        _ = v15
    } else {
        v3 += 954
        fmt.Println(v2, v9)
        var v16 int32 = v1
        _ = v16
    }
    v1 -= (v1 / (((-(v1))) | 1))
    switch -533 {
    case 32767:
        for fi17 := 0; fi17 < 14; fi17++ {
            if (1.5 <= float64(float32((0.1 * -2.25)))) {
                var v18 float64 = (1e10 / (-2.25 - (0.1 - -2.25)))
                _ = v18
            } else {
                v2 += v2
                for fi19 := 0; fi19 < 15; fi19++ {
                    if false {
                        fmt.Println(v2, v9)
                    } else {
                        switch v9 {
                        case "\n\t":
                            fmt.Println(v3, fi17, v3)
                        }
                        v2 = (^(v2))
                        if (false || true) {
                            switch ("foo bar" + (v9 + "\n\t")) {
                            case "\n\t":
                                fmt.Println(v3, fi17, v9)
                            case "foo bar":
                                v4 = append(v4, ("hello" + ((v9 + ((("foo bar" + ((v9 + v9) + (("hello" + ("a" + v9)) + v9))) + "a") + v9)) + "\n\t")))
                            case "hello":
                                var v20 uint16 = 0
                                _ = v20
                            default:
                                if (!(-2.25 == float64(1e10))) {
                                    if ((!false) || (!true)) {
                                        var v21 uint32 = 4294967295
                                        _ = v21
                                    } else {
                                        v3 = (^(v3))
                                        switch (fi17 * fi17) {
                                        case 4294967295:
                                            var v22 string = (v9 + ("a" + v9))
                                            _ = v22
                                        case -9223372036854775808:
                                            v1 += v1
                                        case 127:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    }
                                } else {
                                    v23 := GenS1{F0: -1, F1: 2147483647, F2: 127}
                                    var v24 uint16 = 65535
                                    v1 = 1
                                    _ = v23
                                    _ = v24
                                }
                            }
                        }
                        v25 := GenS2{F0: -466}
                        var v26 uint32 = 0
                        if true {
                            if (!false) {
                                v3 += uint64(v2)
                                switch int(v2) {
                                case 9223372036854775807:
                                    v26 += (v26 & 2147483648)
                                default:
                                    var v27 float64 = 1e10
                                    _ = v27
                                }
                                if len(v4) > 0 {
                                    v4[0] = v9
                                }
                            }
                            var v28 int8 = 100
                            v29 := GenS3{F0: (65535 != 42), F1: 153, F2: (v28 - (v28 ^ (v28 >> ((v2) & 63))))}
                            _ = v28
                            _ = v29
                        }
                        v26 = v26
                        var v30 bool = (true && (v26 >= (v26 | 342)))
                        _ = v25
                        _ = v26
                        _ = v30
                    }
                    if (((1.5 / (1e10 / 1.5)) + (3.14159 * -2.25)) == 3.14159) {
                        if true {
                            switch (fi17 % (((fi17 | 127)) | 1)) {
                            case 128:
                                v9 = ((v9 + "hello") + v9)
                            case 0:
                                fmt.Println(fi19, v9)
                            case 4294967295:
                                if true {
                                    v9 = "foo bar"
                                } else {
                                    v1 = v1
                                    var v31 float64 = -2.25
                                    _ = v31
                                }
                            default:
                                v4 = append(v4, "foo bar")
                            }
                            if true {
                                if false {
                                    var v32 int64 = 42
                                    var v33 uint32 = 2147483647
                                    fmt.Println(fi19, fi19)
                                    var v34 int64 = v32
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    fmt.Println(fi17, fi19)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (1 <= -1) {
                                        var v35 string = (v9 + v9)
                                        fmt.Println(fi19, v2)
                                        _ = v35
                                    } else {
                                        v2 = (v2 / (((v2 * (v2 & v2))) | 1))
                                        v36 := GenS2{F0: v1}
                                        _ = v36
                                    }
                                }
                            } else {
                                v4 = append(v4, v9)
                                _ = len(v4)
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v37 bool = (fi17 <= (fi19 - fi19))
                                    if (107 != 121) {
                                        v2 = v2
                                    } else {
                                        v3 += (v3 % (((v3 | (v3 | (v3 | (v3 ^ (v3 - 1)))))) | 1))
                                        switch v9 {
                                        case "hello":
                                            var v38 uint8 = 0
                                            _ = v38
                                        case "":
                                            if ((true && (409 >= 679)) || (32767 < 65535)) {
                                                switch (v9 + "\n\t") {
                                                case "foo bar":
                                                    switch "a" {
                                                    case "hello":
                                                        v39 := GenS4{F0: v3, F1: 1, F2: (v3 ^ v3)}
                                                        _ = v39
                                                    case "foo bar":
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                case "\n\t":
                                                    v1 = v1
                                                case "":
                                                    switch v9 {
                                                    case "foo bar":
                                                        fmt.Println(v37, v2)
                                                    case "a":
                                                        v9 = (v9 + (v9 + "a"))
                                                    default:
                                                        switch v9 {
                                                        case "a":
                                                            var v40 int64 = 9223372036854775807
                                                            _ = v40
                                                        case "hello":
                                                            if (0.0 > 1.5) {
                                                                v9 = v9
                                                                if v37 {
                                                                    v2 = uint(v3)
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v41 uint = 18446744073709551615
                                                                        _ = v41
                                                                    case "":
                                                                        v42 := GenS5{F0: v3, F1: ((^(v3)) <= (v3 << (((v2 | uint(v3))) & 63))), F2: 1000}
                                                                        _ = v42
                                                                    case "\n\t":
                                                                        fmt.Println(fi17)
                                                                    default:
                                                                        var v43 bool = v37
                                                                        _ = v43
                                                                    }
                                                                    fmt.Println(v37, v1, fi17)
                                                                }
                                                            } else {
                                                                var v44 uint32 = 255
                                                                switch v9 {
                                                                case "":
                                                                    var v45 float64 = -2.25
                                                                    _ = v45
                                                                case "foo bar":
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                if len(v4) > 0 {
                                                                    v4[0] = v9
                                                                }
                                                                _ = v44
                                                            }
                                                        default:
                                                            if len(v4) > 0 {
                                                                v4[0] = "foo bar"
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            v3 = (v3 % ((v3) | 1))
                                        default:
                                            var v46 uint8 = 255
                                            _ = v46
                                        }
                                        var v47 string = v9
                                        _ = v47
                                    }
                                    _ = v37
                                }
                                switch (fi19 / ((fi17) | 1)) {
                                case -1:
                                    switch fi17 {
                                    case 523:
                                        fmt.Println(v2, fi19, fi17)
                                    case 205:
                                        v48 := GenS2{F0: v1}
                                        _ = v48
                                    default:
                                        var v49 uint32 = 81
                                        _ = v49
                                    }
                                }
                            }
                        } else {
                            v9 = "foo bar"
                        }
                        if false {
                            var v50 string = v9
                            v51 := GenS5{F0: (^(v3)), F1: false, F2: -695}
                            switch (fi19 / ((797) | 1)) {
                            case -2147483648:
                                if false {
                                    var v52 int64 = 42
                                    _ = v52
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch fi17 {
                                    case 128:
                                        if true {
                                            v4 = append(v4, ("hello" + ((("\n\t" + (v50 + v9)) + "") + v9)))
                                            _ = len(v4)
                                            _ = len(v4)
                                        } else {
                                            _ = len(v4)
                                        }
                                    case -9223372036854775808:
                                        var v53 bool = (-9223372036854775808 != 1)
                                        _ = v53
                                    case 32768:
                                        var v54 int32 = (^(v1))
                                        _ = v54
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = (v50 + v50)
                                        }
                                    }
                                    var v55 float64 = -2.25
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        v50 = v50
                                    default:
                                        v55 -= -2.25
                                    }
                                    if (3.14159 != float32(v55)) {
                                        v9 = ("a" + "")
                                        var v56 uint64 = (^(v3))
                                        fmt.Println(v2)
                                        _ = v56
                                    }
                                    v2 = (v2 | 255)
                                    _ = v55
                                }
                            case 42:
                                var v57 int64 = 410
                                _ = v57
                            case 127:
                                if len(v4) > 0 {
                                    v4[0] = "\n\t"
                                }
                            }
                            _ = v50
                            _ = v51
                        }
                        if true {
                            v2 = (v2 * v2)
                        }
                    }
                    _ = fi19
                }
                if (-17 <= -748) {
                    for fi58 := 0; fi58 < 13; fi58++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v59 float64 = 1.5
                        switch int(v3) {
                        case 9223372036854775807:
                            v9 = (v9 + v9)
                        case 984:
                            var v60 uint32 = 727
                            _ = v60
                        case 128:
                            var v61 int16 = 492
                            _ = v61
                        }
                        _ = fi58
                        _ = v59
                    }
                    for fi62 := 0; fi62 < 4; fi62++ {
                        var v63 uint8 = 0
                        var v64 int = 42
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                    _ = len(v4)
                }
            }
            _ = fi17
        }
    default:
        fmt.Println(v9)
    }
    for fi65 := 0; fi65 < 20; fi65++ {
        if false {
            var v66 uint64 = (v3 / (((v3 / (((v3 ^ 18446744073709551615)) | 1))) | 1))
            _ = v66
        }
        var v67 string = v9
        _ = fi65
        _ = v67
    }
    v9 = (v9 + "foo bar")
    if (-839 != 42) {
        var v68 int32 = v1
        for fi69 := 0; fi69 < 15; fi69++ {
            v70 := GenS6{F0: -128, F1: (!(-1 != v1)), F2: -2147483648}
            _ = fi69
            _ = v70
        }
        var v71 bool = true
        _ = v68
        _ = v71
    } else {
        v72 := GenS1{F0: 9223372036854775807, F1: 2147483647, F2: 127}
        v9 = ""
        _ = len(v4)
        _ = v72
    }
    var v73 float64 = 0.1
    for fi74 := 0; fi74 < 17; fi74++ {
        var v75 string = v9
        fmt.Println(v9, v75, v1)
        fmt.Println(v3)
        _ = fi74
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v73:", v73)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v9
}
