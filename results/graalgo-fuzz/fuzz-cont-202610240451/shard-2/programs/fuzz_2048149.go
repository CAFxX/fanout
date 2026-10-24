package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 uint }
type GenS3 struct { F0 bool; F1 string; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 bool }
type GenS5 struct { F0 bool; F1 uint32; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 int; F2 uint32 }
type GenS10 struct { F0 string; F1 string; F2 uint16 }
type GenS11 struct { F0 int8; F1 uint }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0() int64 {
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int32 = -1
        v7 := GenS1{F0: 2}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int16 = 1000
    _ = v4
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    v3 := []string{"foo bar", ("foo bar" + "a")}
    if (!(!false)) {
        v9 := GenS2{F0: "a", F1: -419, F2: 65535}
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        var v10 uint64 = 940
        _ = v10
        _ = v9
    }
    var v11 string = (("\n\t" + "a") + "hello")
    if (!false) {
        switch -827 {
        case 127:
            v12 := GenS2{F0: "\n\t", F1: 475, F2: 255}
            _ = v12
        case 32767:
            var v13 int8 = 95
            _ = v13
        case 714:
            var v14 uint = 65535
            _ = v14
        }
        v15 := GenS3{F0: (((float64(float32(1.5)) / 3.14159) + -2.25) > 3.14159), F1: v11, F2: -9223372036854775808}
        _ = v15
    } else {
        if false {
            var v16 uint16 = 255
            for fi17 := 0; fi17 < 16; fi17++ {
                if ((42 >= 100) || (true && (true && true))) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        v16 = v16
                        var v19 int8 = -128
                        if (v11 >= "foo bar") {
                            if (fi18 < fi18) {
                                v16 = v16
                            } else {
                                if (!true) {
                                    v16 = v16
                                    v16 = v16
                                }
                                var v20 uint16 = v16
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch 4294967295 {
                                case 42:
                                    v21 := GenS2{F0: v11, F1: 42, F2: 124}
                                    _ = v21
                                case -9223372036854775808:
                                    var v22 int8 = (v19 * (^(v19)))
                                    _ = v22
                                case 4294967295:
                                    v20 = v20
                                default:
                                    fmt.Println(v16, v20, fi18)
                                }
                                switch (v11 + "a") {
                                case "":
                                    fmt.Println(v11)
                                case "hello":
                                    var v23 uint8 = 0
                                    if ((-(v19)) >= v19) {
                                        switch ("" + ("a" + (v11 + (v11 + ("" + ""))))) {
                                        case "foo bar":
                                            fmt.Println(fi18, v16)
                                        default:
                                            if true {
                                                fmt.Println(v23, v20, fi17)
                                                if (v20 == (v16 ^ v20)) {
                                                    if false {
                                                        v24 := GenS4{F0: 18446744073709551615, F1: 0, F2: true}
                                                        switch -2147483648 {
                                                        case -2147483648:
                                                            if true {
                                                                var v25 uint8 = v23
                                                                _ = v25
                                                            }
                                                        case 32767:
                                                            var v26 uint64 = 729
                                                            _ = v26
                                                        case 4294967295:
                                                            var v27 string = v11
                                                            if (0.1 != 0.1) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = ("hello" + "a")
                                                                }
                                                                v28 := v2[v11]
                                                                v3 = append(v3, ("\n\t" + v11))
                                                                v29 := GenS1{F0: (v16 + (v20 >> ((42) & 63)))}
                                                                var v30 int32 = int32(v23)
                                                                _ = len(v3)
                                                                _ = v28
                                                                _ = v29
                                                                _ = v30
                                                            } else {
                                                                v16 += (v20 - (v16 | (v16 * v20)))
                                                                if len(v3) > 0 {
                                                                    v3[0] = "hello"
                                                                }
                                                            }
                                                            v11 = "\n\t"
                                                            _ = v27
                                                        default:
                                                            if ("hello" <= (((v11 + (v11 + ("hello" + (((v11 + v11) + "a") + v11)))) + v11) + v11)) {
                                                                var v31 uint = 255
                                                                _ = v31
                                                            }
                                                        }
                                                        var v32 bool = (2 != 90)
                                                        _ = v24
                                                        _ = v32
                                                    }
                                                    switch v11 {
                                                    case "foo bar":
                                                        switch v11 {
                                                        case "\n\t":
                                                            switch v11 {
                                                            case "hello":
                                                                switch fi18 {
                                                                case -9223372036854775808:
                                                                    v20 += (v16 % ((v20) | 1))
                                                                case -393:
                                                                    if (!false) {
                                                                        v33 := GenS5{F0: false, F1: 65535, F2: 0}
                                                                        fmt.Println(v23, fi18, fi17)
                                                                        _ = v33
                                                                    }
                                                                default:
                                                                    v16 = (v20 % (((^(v20))) | 1))
                                                                }
                                                            default:
                                                                v19 = v19
                                                            }
                                                        default:
                                                            fmt.Println(fi17, v20, v16)
                                                        }
                                                    case "":
                                                        v34 := GenS6{F0: int32(v20)}
                                                        _ = v34
                                                    case "\n\t":
                                                        var v35 uint = 2
                                                        _ = v35
                                                    default:
                                                        v11 = v11
                                                        var v36 uint64 = 0
                                                        v20 = (v16 | 1)
                                                        _ = v36
                                                    }
                                                    fmt.Println(v16, v11)
                                                }
                                                fmt.Println(v11)
                                                switch v11 {
                                                case "foo bar":
                                                    v37 := GenS7{F0: v19}
                                                    _ = v37
                                                }
                                                v20 = 1000
                                                v16 = (v20 | (v20 + uint16(v23)))
                                                fmt.Println(v11)
                                            } else {
                                                v38 := GenS1{F0: v20}
                                                switch v11 {
                                                case "":
                                                    var v39 int16 = -660
                                                    _ = v39
                                                case "foo bar":
                                                    v3 = append(v3, v11)
                                                case "hello":
                                                    var v40 int32 = -2147483648
                                                    _ = v40
                                                }
                                                var v41 float64 = -2.25
                                                var v42 bool = (((float32(1e10) != (0.1 + (-2.25 * (((0.1 / 0.1) / float32((0.1 + (1.5 * ((0.1 * (((v41 / ((1e10 - ((((-2.25 * 1.5) / (v41 / v41)) + ((1.5 * float64(0.1)) - 1.5)) * v41)) / 1e10)) * ((((v41 * -2.25) * float64(0.0)) / 0.1) / float64((0.1 / -2.25)))) * (1e10 * v41))) * ((v41 * 1.5) * 3.14159)))))) + float32(v41))))) || (0 < 2)) && false)
                                                _ = v38
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                        v16 += 65535
                                        fmt.Println(v19)
                                        v23 = 2
                                    }
                                    _ = v23
                                }
                                _ = v20
                            }
                            if true {
                                v3 = append(v3, "\n\t")
                                var v43 uint64 = 982
                                v44 := GenS8{F0: false}
                                _ = v43
                                _ = v44
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                    if (!(!(1 != 747))) {
                        v16 += v16
                        for fi45 := 0; fi45 < 6; fi45++ {
                            v16 -= (v16 & (^(v16)))
                            _ = fi45
                        }
                    }
                }
                _ = fi17
            }
            var v46 uint8 = 200
            _ = v16
            _ = v46
        } else {
            v11 = ""
            if true {
                if false {
                    switch 901 {
                    case -104:
                        if (((3.14159 - 0.1) * 1e10) >= 1e10) {
                            var v47 int32 = -780
                            if (!(1 == 1)) {
                                v48 := GenS8{F0: (false || true)}
                                var v49 float32 = 1.5
                                v47 = v47
                                _ = v48
                                _ = v49
                            }
                            _ = v47
                        } else {
                            v11 = ""
                        }
                    case 9223372036854775807:
                        fmt.Println(v11)
                    }
                } else {
                    if (!true) {
                        fmt.Println(v11)
                    }
                    var v50 uint16 = 304
                    for fi51 := 0; fi51 < 6; fi51++ {
                        switch (v11 + v11) {
                        case "\n\t":
                            switch (fi51 % ((fi51) | 1)) {
                            case 1:
                                switch ("hello" + v11) {
                                case "hello":
                                    for fi52 := 0; fi52 < 6; fi52++ {
                                        switch (v11 + (v11 + "")) {
                                        case "":
                                            fmt.Println(v50, fi51)
                                            var v53 float64 = 1e10
                                            _ = v53
                                        case "hello":
                                            _ = len(v2)
                                        case "a":
                                            fmt.Println(v11, v11, v50)
                                        }
                                        v50 -= v50
                                        _ = fi52
                                    }
                                case "":
                                    v50 = v50
                                }
                            case 9223372036854775807:
                                for fi54 := 0; fi54 < 20; fi54++ {
                                    v11 = v11
                                    var v55 int16 = -1
                                    _ = fi54
                                    _ = v55
                                }
                            case 127:
                                var v56 int32 = 0
                                _ = v56
                            }
                        case "foo bar":
                            v50 = 32767
                        default:
                            switch v11 {
                            case "a":
                                var v57 int64 = genFn0()
                                _ = v57
                            case "\n\t":
                                v50 = (v50 / ((v50) | 1))
                            default:
                                v1[v11] = fi51
                            }
                        }
                        v11 = "hello"
                        v11 = "a"
                        _ = fi51
                    }
                    _ = v50
                }
            } else {
                for fi58 := 0; fi58 < 10; fi58++ {
                    var v59 float64 = (-2.25 + -2.25)
                    var v60 int8 = 100
                    _ = len(v2)
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                v61 := GenS9{F0: 643, F1: -345, F2: 2147483648}
                v62 := GenS10{F0: "foo bar", F1: "\n\t", F2: 65535}
                _ = v61
                _ = v62
            }
        }
        fmt.Println(v11)
    }
    for fi63 := 0; fi63 < 9; fi63++ {
        fmt.Println(fi63)
        _ = fi63
    }
    v11 = (((v11 + v11) + v11) + v11)
    var v64 float32 = ((1.5 * ((0.1 * ((1.5 * 0.0) - (((1.5 / 1e10) / 1e10) + -2.25))) + (1e10 / 0.1))) / (0.1 / (0.1 * -2.25)))
    switch "" {
    case "a":
        for fi65 := 0; fi65 < 20; fi65++ {
            if false {
                if true {
                    switch 687 {
                    case -792:
                        v64 = (0.0 + 0.0)
                    default:
                        v66 := GenS7{F0: -1}
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 20; fi67++ {
                        v68 := GenS2{F0: "foo bar", F1: -649, F2: 719}
                        switch v11 {
                        case "a":
                            var v69 uint32 = 2147483648
                            _ = v69
                        default:
                            v11 = "\n\t"
                        }
                        _ = fi67
                        _ = v68
                    }
                } else {
                    var v70 uint = 0
                    for fi71 := 0; fi71 < 7; fi71++ {
                        var v72 int16 = -551
                        var v73 uint32 = 104
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    var v74 uint16 = 32767
                    _ = v70
                    _ = v74
                }
            } else {
                var v75 int8 = -128
                if (v75 == (v75 >> ((42) & 63))) {
                    if ((!false) || false) {
                        if true {
                            fmt.Println(v75, v64, v11)
                            v75 = v75
                            switch 2147483647 {
                            case 267:
                                v76 := GenS11{F0: (v75 & (^(v75))), F1: 4294967295}
                                _ = v76
                            case -616:
                                fmt.Println(fi65, v75, v64)
                            case 759:
                                var v77 uint64 = 0
                                _ = v77
                            }
                        }
                        v64 = 3.14159
                    } else {
                        var v78 int64 = genFn0()
                        _ = v78
                    }
                    fmt.Println(v64)
                }
                var v79 int8 = (v75 / ((-1) | 1))
                _ = v75
                _ = v79
            }
            _ = fi65
        }
    case "\n\t":
        if (false || true) {
            v64 -= v64
            for fi80 := 0; fi80 < 17; fi80++ {
                fmt.Println(fi80)
                _ = fi80
            }
        } else {
            if false {
                v81 := GenS12{F0: 0}
                _ = v81
            }
            fmt.Println(v64, v11)
        }
    }
    if (1 > -1) {
        var v82 float32 = (v64 - (float32(0.0) + (v64 - (v64 * v64))))
        if false {
            v83 := v2["foo bar"]
            _ = v83
        } else {
            v84 := GenS7{F0: 3}
            _ = v84
        }
        _ = len(v3)
        _ = v82
    } else {
        v11 = v11
        if (2 < 0) {
            var v85 uint8 = 128
            v86 := GenS13{F0: (v85 + (v85 + v85)), F1: 872, F2: v85}
            var v87 int16 = -1
            _ = v85
            _ = v86
            _ = v87
        }
        var v88 uint8 = 127
        _ = v88
    }
    v64 = (v64 - float32(float64(v64)))
    v64 = float32(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v64
}
