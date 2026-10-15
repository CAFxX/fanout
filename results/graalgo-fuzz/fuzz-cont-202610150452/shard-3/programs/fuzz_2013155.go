package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 string; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS8 struct { F0 uint; F1 bool; F2 int8 }

func main() {
    v1 := []uint32{4294967295, 2147483647, 42}
    var v2 int32 = -1
    var v3 int16 = -785
    var v4 int64 = 334
    var v5 string = ((("hello" + "foo bar") + (("foo bar" + ("hello" + "foo bar")) + "foo bar")) + "")
    _ = len(v1)
    var v6 string = v5
    var v7 int = 9223372036854775807
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch v6 {
    case "":
        var v8 uint64 = 0
        _ = v8
    }
    var v9 uint8 = 184
    var v10 uint8 = (v9 & (^(v9)))
    v3 += v3
    var v11 int8 = -1
    for fi12 := 0; fi12 < 8; fi12++ {
        switch "a" {
        case "hello":
            switch v6 {
            case "":
                if true {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v5 = v6
                        _ = fi13
                    }
                    var v14 uint16 = 42
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        if (false && false) {
                            v4 = v4
                        }
                        _ = fi15
                    }
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v17 := GenS1{F0: (v11 % ((v11) | 1)), F1: 479}
                        _ = fi16
                        _ = v17
                    }
                }
            }
        case "a":
            if (false || true) {
                v1 = append(v1, 813)
                switch (v7 * (v7 / (((fi12 - 0)) | 1))) {
                case 4294967295:
                    switch "a" {
                    case "a":
                        v10 = v9
                    case "":
                        if (!true) {
                            var v18 int16 = v3
                            switch v5 {
                            case "":
                                v19 := GenS1{F0: 42, F1: 255}
                                _ = v19
                            case "foo bar":
                                if (v11 <= -19) {
                                    _ = len(v1)
                                    var v20 uint8 = v9
                                    var v21 float64 = -2.25
                                    _ = v20
                                    _ = v21
                                } else {
                                    var v22 float32 = -2.25
                                    for fi23 := 0; fi23 < 14; fi23++ {
                                        v3 = -1
                                        if false {
                                            v24 := GenS2{F0: 255}
                                            _ = v24
                                        }
                                        _ = fi23
                                    }
                                    _ = v22
                                }
                            default:
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    switch int(v3) {
                                    case 447:
                                        v26 := GenS3{F0: 0, F1: v5, F2: 852}
                                        _ = v26
                                    case -9223372036854775808:
                                        v7 = 2147483647
                                    case 127:
                                        if len(v1) > 0 {
                                            v1[0] = 42
                                        }
                                    }
                                    var v27 float32 = (-2.25 / (0.1 - ((-2.25 + (0.0 + float32(-2.25))) + 0.0)))
                                    _ = fi25
                                    _ = v27
                                }
                            }
                            _ = v18
                        }
                    case "\n\t":
                        for fi28 := 0; fi28 < 14; fi28++ {
                            v29 := GenS2{F0: 910}
                            var v30 int = v7
                            v2 = 0
                            switch "" {
                            case "foo bar":
                                v10 += v9
                            }
                            v31 := GenS2{F0: 2}
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                            if false {
                                v2 = v2
                                fmt.Println(v3, v30)
                                var v32 int64 = 107
                                _ = v32
                            } else {
                                switch 32768 {
                                case -489:
                                    var v33 int32 = v2
                                    _ = v33
                                default:
                                    if ((574 == 0) && ((v5 != v5) || false)) {
                                        var v34 int8 = int8(v4)
                                        switch v30 {
                                        case 2147483647:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case -452:
                                            var v35 bool = ((((v4 >> ((18446744073709551615) & 63)) != (v4 / (((v4 % (((v4 + 42)) | 1))) | 1))) || true) || false)
                                            _ = v35
                                        case -1:
                                            if true {
                                                var v36 int16 = (^(v3))
                                                switch v5 {
                                                case "\n\t":
                                                    if ((float64(1e10) == (1e10 * -2.25)) && (704 > 2147483647)) {
                                                        v37 := GenS4{F0: 2}
                                                        if (!false) {
                                                            if (0 > -1) {
                                                                v38 := GenS1{F0: v11, F1: 0}
                                                                if false {
                                                                    if (!(false || false)) {
                                                                        v36 += v3
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 338
                                                                        }
                                                                        v39 := GenS1{F0: (v11 ^ v34), F1: 255}
                                                                        _ = v39
                                                                    }
                                                                    v1 = append(v1, 42)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 452
                                                                    }
                                                                    switch (fi28 * v7) {
                                                                    case 875:
                                                                        v40 := GenS2{F0: 65535}
                                                                        _ = v40
                                                                    case 127:
                                                                        var v41 float32 = (0.0 - ((((((0.1 - float32((float64(float32(0.1)) / 1.5))) + (1e10 + 0.0)) - (0.1 * ((3.14159 - 1e10) + (3.14159 / (1e10 / (3.14159 * 3.14159)))))) / -2.25) * 1e10) * 0.1))
                                                                        _ = v41
                                                                    default:
                                                                        switch ("" + v5) {
                                                                        case "\n\t":
                                                                            if (false && (false && true)) {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 65535
                                                                                }
                                                                                v10 -= v9
                                                                                switch v5 {
                                                                                case "\n\t":
                                                                                    var v42 bool = (!(((758 >= 1) || (!false)) || true))
                                                                                    _ = v42
                                                                                }
                                                                            }
                                                                        default:
                                                                            v3 += 32767
                                                                        }
                                                                    }
                                                                    var v43 int = (v7 << ((4294967295) & 63))
                                                                    _ = v43
                                                                }
                                                                if false {
                                                                    v7 = (-(fi12))
                                                                    v44 := GenS4{F0: 2}
                                                                    _ = len(v1)
                                                                    _ = v44
                                                                }
                                                                _ = v38
                                                            } else {
                                                                v9 = (v10 % ((v9) | 1))
                                                                var v45 uint16 = 2
                                                                v46 := GenS5{F0: v7}
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                            v47 := GenS6{F0: uint64(v9)}
                                                            v30 -= (v30 & fi12)
                                                            _ = v47
                                                        }
                                                        _ = v37
                                                    } else {
                                                        _ = len(v1)
                                                        if (!(2147483647 < 152)) {
                                                            var v48 int = int(v11)
                                                            fmt.Println(v5, v11)
                                                            switch v5 {
                                                            case "hello":
                                                                var v49 uint16 = 1
                                                                _ = v49
                                                            case "foo bar":
                                                                switch 888 {
                                                                case 42:
                                                                    var v50 float32 = 3.14159
                                                                    _ = v50
                                                                case -2147483648:
                                                                    var v51 uint64 = 400
                                                                    _ = v51
                                                                case 1:
                                                                    if true {
                                                                        var v52 int16 = v3
                                                                        _ = v52
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                switch v5 {
                                                                case "":
                                                                    v53 := GenS7{F0: v10, F1: v9, F2: uint64(v9)}
                                                                    _ = v53
                                                                case "hello":
                                                                    v1 = append(v1, 2147483648)
                                                                default:
                                                                    v30 -= v48
                                                                }
                                                            }
                                                            _ = v48
                                                        }
                                                    }
                                                }
                                                fmt.Println(v36, v9, fi28)
                                                _ = v36
                                            } else {
                                                v10 = v10
                                            }
                                        default:
                                            v10 -= (v10 + 2)
                                        }
                                        _ = v34
                                    } else {
                                        v54 := GenS2{F0: 2}
                                        _ = v54
                                    }
                                }
                                fmt.Println(v10)
                                if false {
                                    var v55 uint = 42
                                    v1 = append(v1, 42)
                                    var v56 int32 = (v2 / ((1) | 1))
                                    var v57 uint32 = 82
                                    v55 = (v55 & uint(v57))
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                } else {
                                    if (!(!false)) {
                                        if (305 == 0) {
                                            v9 = (v10 >> ((uint(v9)) & 63))
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = 1
                                            }
                                            var v58 bool = ((3.14159 + 1e10) == -2.25)
                                            if v58 {
                                                fmt.Println(v58, v6, v4)
                                                if v58 {
                                                    if v58 {
                                                        var v59 int16 = 0
                                                        v60 := GenS6{F0: 42}
                                                        v2 = int32(v11)
                                                        v1 = append(v1, 2147483647)
                                                        v1 = append(v1, 956)
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        var v61 string = ("" + v5)
                                                        _ = v61
                                                    }
                                                    v1 = append(v1, 65535)
                                                    switch v7 {
                                                    case 0:
                                                        v1 = append(v1, 499)
                                                    default:
                                                        if (!((0.0 * 1e10) <= 3.14159)) {
                                                            switch v30 {
                                                            case 4294967295:
                                                                v7 = 4294967295
                                                            case -978:
                                                                v62 := GenS2{F0: 291}
                                                                _ = v62
                                                            case -9223372036854775808:
                                                                fmt.Println(v3, v4)
                                                            }
                                                            v1 = append(v1, 608)
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v58
                                        }
                                        v1 = append(v1, 2)
                                    } else {
                                        fmt.Println(fi12, v4, fi28)
                                        if (v3 > (v3 ^ (-(v3)))) {
                                            var v63 float64 = 3.14159
                                            _ = v63
                                        }
                                        if false {
                                            v4 += 209
                                            if len(v1) > 0 {
                                                v1[0] = 4294967295
                                            }
                                            v64 := GenS8{F0: 1, F1: ((^(v10)) < (^(v10))), F2: (v11 ^ v11)}
                                            switch v6 {
                                            case "\n\t":
                                                switch v6 {
                                                case "":
                                                    var v65 int = (v7 - fi28)
                                                    _ = v65
                                                default:
                                                    var v66 uint64 = 4294967295
                                                    v10 -= 128
                                                    var v67 bool = false
                                                    _ = v66
                                                    _ = v67
                                                }
                                            }
                                            _ = v64
                                        }
                                    }
                                }
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    default:
                        fmt.Println(v2, v2, v7)
                    }
                case 9223372036854775807:
                    var v68 uint = 65535
                    _ = v68
                case 127:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v69 uint32 = 65535
                    _ = v69
                }
                switch (((("hello" + v6) + "hello") + v5) + v6) {
                case "":
                    for fi70 := 0; fi70 < 18; fi70++ {
                        v11 -= (v11 ^ v11)
                        _ = fi70
                    }
                case "foo bar":
                    fmt.Println(v4, v5)
                }
            }
        case "":
            fmt.Println(v10)
        }
        v10 = 127
        switch (^(v7)) {
        case -533:
            switch v5 {
            case "\n\t":
                v71 := GenS3{F0: uint16(v9), F1: ("hello" + v6), F2: 706}
                _ = v71
            }
        default:
            if true {
                fmt.Println(fi12, fi12, v10)
                switch v7 {
                case 128:
                    for fi72 := 0; fi72 < 17; fi72++ {
                        v6 = v6
                        fmt.Println(v7, v10, v11)
                        _ = fi72
                    }
                case 4294967295:
                    v6 = "a"
                default:
                    v9 = 41
                }
            } else {
                var v73 int64 = v4
                var v74 int32 = int32(v11)
                _ = v73
                _ = v74
            }
        }
        _ = fi12
    }
    var v75 int = (v7 ^ (v7 ^ (v7 ^ (v7 >> ((255) & 63)))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v75:", v75)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v75
    _ = v9
}
