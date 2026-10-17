package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 int64; F1 int64 }

func main() {
    var v1 int8 = 100
    var v2 uint = 65535
    v3 := []string{"foo bar"}
    if true {
        var v4 float64 = 3.14159
        _ = v4
    }
    v2 = v2
    switch "hello" {
    case "foo bar":
        v1 += (v1 * (^(v1)))
    default:
        v1 += (v1 & (-(v1)))
    }
    v5 := GenS1{F0: ("hello" + "a")}
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint64 = 0
        _ = fi6
        _ = v7
    }
    switch ("foo bar" + "") {
    case "":
        fmt.Println(v1, v1)
    case "hello":
        switch "" {
        case "foo bar":
            var v8 uint = v2
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 float64 = 0.0
                v11 := GenS1{F0: "hello"}
                var v12 bool = (v10 < (-2.25 / 1.0))
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        default:
            if false {
                v3 = append(v3, ("" + "foo bar"))
            } else {
                v13 := GenS2{F0: "a"}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v13
            }
        }
    }
    if (false && (1 > -325)) {
        switch ("hello" + ("foo bar" + (("\n\t" + ("" + "\n\t")) + "hello"))) {
        case "foo bar":
            v1 = (v1 * -1)
        case "a":
            v14 := GenS3{F0: (v2 | (v2 >> ((18446744073709551615) & 63))), F1: -488}
            _ = v14
        case "\n\t":
            switch -304 {
            case 757:
                v3 = append(v3, (("foo bar" + ("hello" + "")) + "hello"))
            case 854:
                for fi15 := 0; fi15 < 17; fi15++ {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(fi15)
                    case "a":
                        if false {
                            if true {
                                switch "" {
                                case "foo bar":
                                    if true {
                                        switch 42 {
                                        case 32767:
                                            v16 := GenS1{F0: "\n\t"}
                                            _ = v16
                                        case 32768:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case 128:
                                            fmt.Println(fi15, v1, v1)
                                        }
                                        v1 = (-(v1))
                                    }
                                case "a":
                                    for fi17 := 0; fi17 < 20; fi17++ {
                                        v18 := GenS2{F0: "\n\t"}
                                        switch ("\n\t" + "a") {
                                        case "\n\t":
                                            v2 = (v2 % (((v2 & v2)) | 1))
                                            if false {
                                                v1 = -52
                                                if false {
                                                    var v19 int8 = v1
                                                    _ = v19
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v2 += (v2 & v2)
                                                    var v20 string = ("\n\t" + ("a" + "foo bar"))
                                                    _ = v20
                                                }
                                            }
                                            var v21 float32 = (1.5 - ((((0.0 / 1.5) / 1.5) + 0.0) * 3.14159))
                                            _ = v21
                                        default:
                                            if (!(0 != 67)) {
                                                switch "hello" {
                                                case "\n\t":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                default:
                                                    v22 := GenS4{F0: v1}
                                                    _ = v22
                                                }
                                                if false {
                                                    if true {
                                                        v23 := GenS4{F0: -128}
                                                        var v24 uint = (^(v2))
                                                        v25 := GenS4{F0: v1}
                                                        switch "a" {
                                                        case "hello":
                                                            var v26 uint16 = 42
                                                            _ = v26
                                                        case "a":
                                                            switch (("a" + "\n\t") + "\n\t") {
                                                            case "":
                                                                v24 = (v2 >> (((v24 ^ (v2 / ((65535) | 1)))) & 63))
                                                            case "a":
                                                                if true {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v1 += v1
                                                                    }
                                                                    v24 -= v24
                                                                }
                                                            case "foo bar":
                                                                v27 := GenS2{F0: ""}
                                                                _ = v27
                                                            }
                                                        case "\n\t":
                                                            if true {
                                                                var v28 float64 = 1e10
                                                                _ = v28
                                                            }
                                                        default:
                                                            if (0 > 0) {
                                                                v3 = append(v3, "hello")
                                                            }
                                                        }
                                                        _ = v23
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        fmt.Println(v1)
                                                        var v29 bool = true
                                                        if (!(!(!(0 <= 2)))) {
                                                            if (976 != 695) {
                                                                if (v29 || v29) {
                                                                    var v30 int = -449
                                                                    _ = v30
                                                                } else {
                                                                    fmt.Println(v1, fi17)
                                                                }
                                                            }
                                                            v31 := GenS5{F0: (v1 >> ((v2) & 63))}
                                                            _ = v31
                                                        }
                                                        _ = v29
                                                    }
                                                    v32 := GenS6{F0: -2147483648, F1: 609}
                                                    _ = len(v3)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    fmt.Println(v2)
                                                    _ = v32
                                                }
                                                fmt.Println(fi17, fi17, v2)
                                                v1 = (v1 >> (((v2 & v2)) & 63))
                                            }
                                        }
                                        v1 += (v1 / ((1) | 1))
                                        v33 := GenS4{F0: (^(v1))}
                                        _ = fi17
                                        _ = v18
                                        _ = v33
                                    }
                                default:
                                    var v34 int32 = 0
                                    _ = v34
                                }
                                v2 -= v2
                                v35 := GenS2{F0: "hello"}
                                _ = v35
                            }
                            v2 += 65535
                            if ((v2 != (^(v2))) || true) {
                                if (0 > -207) {
                                    fmt.Println(v2, v2)
                                }
                                fmt.Println(v2)
                                var v36 float32 = 1.5
                                _ = v36
                            }
                        }
                    case "":
                        v1 = v1
                    }
                    for fi37 := 0; fi37 < 18; fi37++ {
                        v2 -= (v2 * 4294967295)
                        _ = fi37
                    }
                    var v38 int32 = -2147483648
                    _ = fi15
                    _ = v38
                }
            }
        }
        fmt.Println(v1, v1)
        for fi39 := 0; fi39 < 16; fi39++ {
            if (230 < 143) {
                var v40 uint16 = 18
                switch fi39 {
                case 1:
                    v41 := GenS7{F0: 42, F1: 94}
                    _ = v41
                }
                _ = v40
            }
            _ = fi39
        }
    } else {
        var v42 bool = (789 != 985)
        v2 -= v2
        _ = v42
    }
    v43 := GenS3{F0: v2, F1: -9223372036854775808}
    v1 = -11
    if true {
        for fi44 := 0; fi44 < 1; fi44++ {
            switch "\n\t" {
            case "foo bar":
                v45 := GenS8{F0: 255, F1: 1, F2: -1}
                _ = v45
            case "hello":
                switch "" {
                case "":
                    if ((v1 << ((v2) & 63)) <= v1) {
                        v3 = append(v3, ("hello" + "hello"))
                    }
                case "a":
                    var v46 uint16 = 255
                    _ = v46
                }
            case "\n\t":
                switch 9223372036854775807 {
                case 42:
                    var v47 float32 = ((-2.25 * 0.0) * (3.14159 * 3.14159))
                    _ = v47
                case 307:
                    var v48 int32 = 2147483647
                    _ = v48
                case -2147483648:
                    var v49 bool = (0 == 1)
                    _ = v49
                }
            }
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v50 := GenS1{F0: ""}
            _ = fi44
            _ = v50
        }
        var v51 uint32 = uint32(v2)
        _ = v51
    } else {
        v2 = (v2 % ((v2) | 1))
        var v52 int16 = -32768
        _ = v52
    }
    for fi53 := 0; fi53 < 9; fi53++ {
        var v54 uint8 = 51
        switch "foo bar" {
        case "a":
            for fi55 := 0; fi55 < 18; fi55++ {
                if ((fi55 << ((v2) & 63)) >= fi53) {
                    v54 = v54
                    v1 = (-(v1))
                }
                var v56 uint16 = uint16(v54)
                if true {
                    switch "\n\t" {
                    case "hello":
                        var v57 uint64 = 4294967295
                        _ = v57
                    case "":
                        v54 = 2
                    case "a":
                        v3 = append(v3, "foo bar")
                    default:
                        var v58 string = ("hello" + "\n\t")
                        _ = v58
                    }
                }
                _ = fi55
                _ = v56
            }
        }
        _ = fi53
        _ = v54
    }
    for fi59 := 0; fi59 < 4; fi59++ {
        if false {
            var v60 int = -439
            if false {
                v61 := GenS9{F0: -128, F1: ("foo bar" + ((("foo bar" + ((((("hello" + ("" + "\n\t")) + ("foo bar" + "hello")) + "hello") + "") + (("\n\t" + ((("" + ("\n\t" + ("hello" + ""))) + ("\n\t" + "")) + "\n\t")) + ""))) + "") + "\n\t"))}
                _ = v61
            } else {
                if (!true) {
                    fmt.Println(v60, fi59)
                } else {
                    if false {
                        v60 += 9223372036854775807
                        for fi62 := 0; fi62 < 3; fi62++ {
                            fmt.Println(fi62)
                            v63 := GenS2{F0: ("\n\t" + ("hello" + "foo bar"))}
                            var v64 uint64 = 0
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    }
                    v60 += fi59
                    if true {
                        if false {
                            for fi65 := 0; fi65 < 7; fi65++ {
                                switch "hello" {
                                case "a":
                                    v1 = 42
                                case "hello":
                                    switch "" {
                                    case "a":
                                        fmt.Println(v60)
                                    }
                                default:
                                    fmt.Println(fi65, v60)
                                }
                                _ = fi65
                            }
                        } else {
                            v66 := GenS10{F0: 1, F1: 0}
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 11; fi67++ {
                            switch "" {
                            case "hello":
                                var v68 float32 = -2.25
                                _ = v68
                            case "":
                                v2 = v2
                            case "\n\t":
                                v1 -= v1
                            }
                            _ = fi67
                        }
                    }
                }
                for fi69 := 0; fi69 < 19; fi69++ {
                    var v70 int = v60
                    _ = fi69
                    _ = v70
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v60 = v60
            _ = v60
        }
        _ = fi59
    }
    if true {
        var v71 float32 = ((float32(1.5) * (3.14159 - (1e10 * -2.25))) + ((1e10 - (1e10 * 0.0)) * -2.25))
        fmt.Println(v1)
        v1 = (v1 | 127)
        _ = v71
    }
    v3 = append(v3, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v5
}
