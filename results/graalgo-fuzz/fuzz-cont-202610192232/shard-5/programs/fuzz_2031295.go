package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int; F1 string; F2 int64 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 string; F1 uint8; F2 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16 }

func main() {
    var v1 string = ("" + "")
    var v2 int64 = 42
    var v3 int16 = 804
    var v4 int16 = -508
    v5 := []string{}
    if false {
        v6 := GenS1{F0: v2, F1: 255, F2: 1}
        _ = len(v5)
        var v7 uint16 = 255
        _ = v6
        _ = v7
    }
    fmt.Println(v3)
    v8 := GenS1{F0: (v2 & 1), F1: 14, F2: 0}
    if false {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS2{F0: 0, F1: (fi9 / (((fi9 * fi9)) | 1)), F2: 0}
            switch (v1 + (v1 + ((v1 + "a") + v1))) {
            case "":
                for fi11 := 0; fi11 < 4; fi11++ {
                    fmt.Println(v1, fi9, v4)
                    var v12 uint64 = 495
                    switch v1 {
                    case "foo bar":
                        var v13 string = ("\n\t" + v1)
                        _ = v13
                    case "hello":
                        switch (v1 + ("a" + ("hello" + "\n\t"))) {
                        case "a":
                            v12 -= v12
                        case "\n\t":
                            v4 -= v3
                        }
                        switch (fi9 + fi9) {
                        case 0:
                            var v14 uint = 4294967295
                            _ = v14
                        case 42:
                            v15 := GenS2{F0: 1, F1: (fi11 << ((4294967295) & 63)), F2: 751}
                            _ = v15
                        default:
                            switch fi11 {
                            case 32768:
                                v3 += (-(v3))
                            default:
                                v5 = append(v5, v1)
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "a":
                        if true {
                            v2 += v2
                            if false {
                                if (!true) {
                                    v12 = v12
                                    var v16 int32 = 835
                                    _ = v16
                                } else {
                                    var v17 int16 = v3
                                    var v18 uint8 = 42
                                    if false {
                                        var v19 string = "foo bar"
                                        v5 = append(v5, ("" + (v19 + v19)))
                                        var v20 float32 = 1.5
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                var v21 int64 = (v2 >> ((0) & 63))
                                _ = v21
                            } else {
                                switch ("\n\t" + ((v1 + v1) + v1)) {
                                case "hello":
                                    fmt.Println(v1)
                                }
                                v2 = (v2 ^ 1)
                            }
                            var v22 float32 = (3.14159 / 1e10)
                            _ = v22
                        } else {
                            v3 = (v3 % (((v3 % ((v3) | 1))) | 1))
                            switch ("\n\t" + v1) {
                            case "":
                                fmt.Println(v3, v3)
                            case "a":
                                fmt.Println(v2, v2, fi11)
                            case "hello":
                                var v23 uint8 = 173
                                _ = v23
                            default:
                                if true {
                                    _ = len(v5)
                                    if ((v4 <= -942) || true) {
                                        _ = len(v5)
                                        fmt.Println(v3)
                                        if (!(((-2.25 * (0.0 + (0.0 - 0.1))) < -2.25) && ((false && true) || true))) {
                                            fmt.Println(v4, fi9, v2)
                                            v2 = (v2 - (v2 + (v2 + v2)))
                                        }
                                        v24 := GenS2{F0: 1, F1: -507, F2: 798}
                                        v25 := GenS1{F0: -433, F1: 255, F2: 42}
                                        _ = v24
                                        _ = v25
                                    }
                                    v3 -= (v3 % (((v4 - v4)) | 1))
                                } else {
                                    var v26 int64 = v2
                                    var v27 uint8 = 127
                                    _ = v26
                                    _ = v27
                                }
                            }
                            switch fi11 {
                            case 9223372036854775807:
                                var v28 int8 = 127
                                _ = v28
                            }
                            v4 += v3
                        }
                        var v29 string = "foo bar"
                        _ = len(v5)
                        _ = v29
                    }
                    if (116 >= 1) {
                        var v30 int16 = (^(v4))
                        v5 = append(v5, "foo bar")
                        _ = v30
                    }
                    fmt.Println(v4)
                    _ = fi11
                    _ = v12
                }
            case "hello":
                if (!true) {
                    for fi31 := 0; fi31 < 15; fi31++ {
                        var v32 int16 = (v3 - v4)
                        _ = fi31
                        _ = v32
                    }
                    v4 += -1
                    v1 = (v1 + v1)
                } else {
                    v1 = (v1 + "foo bar")
                    switch fi9 {
                    case -1:
                        v3 += v3
                    default:
                        _ = len(v5)
                    }
                }
            case "\n\t":
                if (!true) {
                    var v33 int8 = 127
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v33
                } else {
                    v34 := GenS3{F0: (42 <= 255), F1: -1}
                    _ = v34
                }
            }
            _ = fi9
            _ = v10
        }
        if (2 == 2) {
            v4 = (v4 << ((663) & 63))
        }
    } else {
        var v35 bool = ((false || false) || false)
        if (3.14159 == (1e10 - (0.1 + ((((1e10 / (0.0 - 3.14159)) * 1e10) + float32((-2.25 * float64(0.0)))) / ((float32(((float64(float32(float64(float32((3.14159 - ((1.5 * float64(1e10)) * 1.5)))))) * 0.1) * 0.0)) + (0.1 - 0.1)) - 1e10))))) {
            v3 = (v3 & v3)
            v36 := GenS1{F0: (v2 * v2), F1: 1, F2: 2}
            _ = v36
        } else {
            if false {
                for fi37 := 0; fi37 < 13; fi37++ {
                    fmt.Println(v4, v4, v35)
                    v3 -= v3
                    _ = fi37
                }
                for fi38 := 0; fi38 < 3; fi38++ {
                    var v39 uint = 0
                    _ = fi38
                    _ = v39
                }
            } else {
                for fi40 := 0; fi40 < 17; fi40++ {
                    _ = len(v5)
                    v41 := GenS4{F0: 32768, F1: "hello", F2: v2}
                    for fi42 := 0; fi42 < 11; fi42++ {
                        v1 = v1
                        v1 = v1
                        fmt.Println(v4, fi42, v3)
                        _ = fi42
                    }
                    _ = fi40
                    _ = v41
                }
                switch (("foo bar" + v1) + v1) {
                case "":
                    var v43 int = 85
                    _ = v43
                }
                switch 4294967295 {
                case 133:
                    if (80 != 128) {
                        fmt.Println(v4, v35, v4)
                        v44 := GenS2{F0: -128, F1: -2147483648, F2: 709}
                        _ = v44
                    } else {
                        var v45 uint64 = 18446744073709551615
                        var v46 int = 0
                        _ = v45
                        _ = v46
                    }
                case -620:
                    v47 := GenS5{F0: v1, F1: 119}
                    _ = v47
                case 287:
                    v4 = (v4 >> ((18446744073709551615) & 63))
                default:
                    var v48 float32 = 3.14159
                    _ = v48
                }
            }
            var v49 uint16 = 65535
            for fi50 := 0; fi50 < 7; fi50++ {
                v2 = v2
                for fi51 := 0; fi51 < 19; fi51++ {
                    if (!false) {
                        if v35 {
                            v52 := GenS6{F0: "hello", F1: 2, F2: false}
                            var v53 int8 = 100
                            v5 = append(v5, v1)
                            _ = v52
                            _ = v53
                        } else {
                            var v54 uint32 = 4294967295
                            switch ("hello" + v1) {
                            case "":
                                v55 := GenS7{F0: (v54 >> ((0) & 63))}
                                _ = v55
                            case "hello":
                                if len(v5) > 0 {
                                    v5[0] = (v1 + v1)
                                }
                            case "foo bar":
                                fmt.Println(v35)
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            var v56 int64 = 58
                            _ = v54
                            _ = v56
                        }
                        v4 -= (^(v3))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        var v57 uint32 = 42
                        _ = v57
                    }
                    var v58 int16 = (v3 ^ (v4 - 42))
                    v4 += 32767
                    _ = fi51
                    _ = v58
                }
                _ = fi50
            }
            _ = v49
        }
        _ = v35
    }
    for fi59 := 0; fi59 < 19; fi59++ {
        var v60 string = ((v1 + "") + "foo bar")
        for fi61 := 0; fi61 < 12; fi61++ {
            v1 = v60
            if true {
                if (!true) {
                    if (1 <= 733) {
                        v4 = (v4 | v3)
                    }
                    var v62 bool = false
                    _ = v62
                }
                if (744 != 436) {
                    if (!true) {
                        fmt.Println(v3, v3)
                    }
                    var v63 float32 = float32(float64(1.5))
                    var v64 uint16 = 32767
                    _ = v63
                    _ = v64
                } else {
                    switch fi59 {
                    case 42:
                        v65 := GenS8{F0: (v4 | v3)}
                        _ = v65
                    case 0:
                        v2 = v2
                    }
                    if (true || false) {
                        var v66 string = v60
                        if true {
                            var v67 uint = 42
                            if len(v5) > 0 {
                                v5[0] = "foo bar"
                            }
                            _ = v67
                        } else {
                            if false {
                                switch 32767 {
                                case 159:
                                    fmt.Println(v1, v66)
                                case 9223372036854775807:
                                    var v68 float64 = 1e10
                                    _ = v68
                                case 463:
                                    var v69 int8 = -60
                                    v2 -= 9223372036854775807
                                    var v70 float64 = 0.1
                                    v4 -= 1
                                    _ = v69
                                    _ = v70
                                }
                                if false {
                                    var v71 bool = true
                                    _ = v71
                                }
                            } else {
                                v2 = v2
                            }
                            v4 = v4
                            fmt.Println(v1, v3, v66)
                        }
                        _ = v66
                    }
                    fmt.Println(v1, fi61)
                }
                v4 = (v3 | v4)
                var v72 bool = true
                v3 -= v3
                var v73 bool = true
                _ = v72
                _ = v73
            } else {
                v74 := GenS8{F0: v4}
                _ = v74
            }
            _ = fi61
        }
        if (431 >= -2147483648) {
            var v75 uint32 = 0
            _ = v75
        }
        _ = fi59
        _ = v60
    }
    switch v1 {
    case "a":
        v4 += int16(v2)
    case "hello":
        for fi76 := 0; fi76 < 19; fi76++ {
            v77 := GenS7{F0: 187}
            _ = fi76
            _ = v77
        }
    }
    switch v1 {
    case "":
        if (32768 >= 85) {
            for fi78 := 0; fi78 < 17; fi78++ {
                var v79 float32 = 3.14159
                v3 = v4
                _ = len(v5)
                _ = fi78
                _ = v79
            }
        }
    }
    switch v1 {
    case "\n\t":
        v80 := GenS6{F0: (v1 + v1), F1: 96, F2: true}
        _ = v80
    case "hello":
        fmt.Println(v3, v3)
    }
    var v81 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v81
}
