package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 bool; F1 string; F2 uint64 }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool; F1 int }

func main() {
    var v1 int16 = 1000
    var v2 int64 = -1
    var v3 float32 = 1.5
    var v4 string = "a"
    v5 := make(map[string]uint32)
    var v6 float64 = 0.1
    v7 := []int{-9223372036854775808, -277, -2147483648}
    for fi8 := 0; fi8 < 12; fi8++ {
        if false {
            fmt.Println(v6, fi8, fi8)
            fmt.Println(v2, v4)
            if true {
                v9 := GenS1{F0: 42, F1: v1, F2: -298}
                var v10 int64 = v2
                v11 := GenS1{F0: (v10 | -9223372036854775808), F1: (v1 + 32767), F2: -2147483648}
                _ = v10
                _ = v11
                _ = v9
            }
        }
        _ = fi8
    }
    v3 -= v3
    var v12 bool = (!(!false))
    v3 = 3.14159
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            var v14 uint64 = 42
            fmt.Println(v3, v12)
            for fi15 := 0; fi15 < 13; fi15++ {
                v12 = v12
                if false {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v16 uint32 = 1
                    v1 = -1
                    var v17 float64 = float64((v3 * (0.1 + (1e10 - (float32(v6) - ((1.5 - float32((v6 / (v6 * v6)))) / v3))))))
                    if v12 {
                        var v18 uint = 42
                        _ = v18
                    } else {
                        v2 = (-(v2))
                        v1 = -1
                    }
                    _ = v16
                    _ = v17
                } else {
                    v19 := GenS1{F0: -489, F1: 1000, F2: 42}
                    if true {
                        switch (^(fi13)) {
                        case -354:
                            if (-2147483648 == 0) {
                                v14 -= 0
                            } else {
                                if v12 {
                                    switch fi15 {
                                    case -558:
                                        var v20 string = (v4 + v4)
                                        _ = v20
                                    case 607:
                                        var v21 int64 = -1
                                        _ = v21
                                    default:
                                        switch fi13 {
                                        case 127:
                                            var v22 bool = v12
                                            _ = v22
                                        case -1:
                                            v2 = (v2 | (^(v2)))
                                        }
                                    }
                                    switch (v4 + ("foo bar" + v4)) {
                                    case "foo bar":
                                        switch fi13 {
                                        case 1:
                                            switch fi13 {
                                            case 2147483647:
                                                v7 = append(v7, (fi13 ^ fi15))
                                            case -2147483648:
                                                var v23 uint16 = 255
                                                _ = v23
                                            }
                                        case -216:
                                            v24 := GenS1{F0: (v2 / (((v2 + -338)) | 1)), F1: (-(v1)), F2: 2147483647}
                                            _ = v24
                                        default:
                                            fmt.Println(v4, v2)
                                        }
                                    case "\n\t":
                                        if false {
                                            fmt.Println(v12, v6)
                                            switch fi15 {
                                            case 127:
                                                if v12 {
                                                    v4 = (((v4 + ("a" + ((v4 + ("\n\t" + v4)) + v4))) + ((v4 + v4) + (v4 + "a"))) + "\n\t")
                                                } else {
                                                    switch fi13 {
                                                    case -793:
                                                        v25 := GenS1{F0: 9223372036854775807, F1: v1, F2: 947}
                                                        _ = v25
                                                    }
                                                    _ = len(v5)
                                                    switch 32767 {
                                                    case 1:
                                                        v5[v4] = 2147483647
                                                    case -234:
                                                        switch fi15 {
                                                        case 127:
                                                            if ((-(v1)) < (v1 + v1)) {
                                                                v6 = (v6 / v6)
                                                                var v26 float64 = ((((1e10 + v6) * (float64((v3 / 0.1)) * 1.5)) + float64(v3)) * v6)
                                                                _ = v26
                                                            }
                                                        case 42:
                                                            if v12 {
                                                                if v12 {
                                                                    if (true && (0 == 127)) {
                                                                        v3 -= 0.0
                                                                        v1 = (v1 | (v1 * (^(v1))))
                                                                        v27 := GenS2{F0: v12}
                                                                        _ = v27
                                                                    } else {
                                                                        switch v4 {
                                                                        case "hello":
                                                                            var v28 string = (v4 + "hello")
                                                                            _ = v28
                                                                        case "a":
                                                                            var v29 uint = 2
                                                                            _ = v29
                                                                        case "foo bar":
                                                                            if ((4294967295 >= 65535) || (42 < 128)) {
                                                                                v14 = v14
                                                                                v2 += v2
                                                                                switch 4294967295 {
                                                                                case 32767:
                                                                                    fmt.Println(v1)
                                                                                default:
                                                                                    v30 := GenS3{F0: 100, F1: false}
                                                                                    _ = v30
                                                                                }
                                                                            } else {
                                                                                switch v4 {
                                                                                case "foo bar":
                                                                                    v12 = false
                                                                                case "hello":
                                                                                    v3 -= float32(float64(v3))
                                                                                case "a":
                                                                                    var v31 uint64 = v14
                                                                                    _ = v31
                                                                                default:
                                                                                    delete(v5, v4)
                                                                                }
                                                                            }
                                                                        default:
                                                                            v2 = -1
                                                                            switch v4 {
                                                                            case "hello":
                                                                                var v32 int64 = (v2 - (-(v2)))
                                                                                _ = v32
                                                                            case "":
                                                                                var v33 bool = (402 > v1)
                                                                                _ = v33
                                                                            case "\n\t":
                                                                                v12 = ((v1 ^ v1) > v1)
                                                                            default:
                                                                                v34 := GenS4{F0: v12}
                                                                                _ = v34
                                                                            }
                                                                            if len(v7) > 0 {
                                                                                v7[0] = fi15
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v3, fi13)
                                                                }
                                                                v1 += (-(v1))
                                                            } else {
                                                                fmt.Println(v6, v14, v14)
                                                                var v35 uint16 = 1
                                                                var v36 int16 = -32768
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                        }
                                                    case 42:
                                                        if ((v2 * v2) <= (v2 ^ (v2 ^ v2))) {
                                                            v3 = 0.1
                                                            v4 = ("\n\t" + v4)
                                                            v4 = v4
                                                        } else {
                                                            if false {
                                                                var v37 int = fi15
                                                                _ = v37
                                                            } else {
                                                                switch fi15 {
                                                                case -1:
                                                                    v38 := GenS5{F0: (fi13 ^ (fi15 % ((fi15) | 1))), F1: 42}
                                                                    _ = v38
                                                                default:
                                                                    fmt.Println(v12, v2)
                                                                }
                                                                fmt.Println(fi13, v1)
                                                            }
                                                            fmt.Println(v4, fi15, v12)
                                                        }
                                                    }
                                                }
                                            case 4294967295:
                                                var v39 uint32 = 71
                                                _ = v39
                                            case 306:
                                                if (1 >= 100) {
                                                    delete(v5, v4)
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v1 -= 42
                                    v40 := GenS6{F0: v2, F1: (v2 >> ((255) & 63))}
                                    _ = v40
                                }
                            }
                        case -1:
                            var v41 uint16 = 1
                            _ = v41
                        case 127:
                            fmt.Println(v4, v6)
                        default:
                            v14 = 150
                        }
                    } else {
                        var v42 int32 = 42
                        if (v4 >= "") {
                            switch (fi13 * fi15) {
                            case 0:
                                switch (-(fi13)) {
                                case 9223372036854775807:
                                    var v43 uint32 = 2
                                    if ((!true) || ((fi13 - fi15) != 128)) {
                                        switch fi15 {
                                        case 9223372036854775807:
                                            var v44 int32 = (v42 | v42)
                                            _ = v44
                                        case 234:
                                            v45 := GenS7{F0: int32(v1), F1: uint(v43)}
                                            _ = v45
                                        case 1:
                                            var v46 uint8 = 243
                                            _ = v46
                                        }
                                    }
                                    v1 = 0
                                    _ = v43
                                }
                            }
                            fmt.Println(v6, v42)
                            if v12 {
                                if (v12 && v12) {
                                    if v12 {
                                        var v47 int8 = 107
                                        v14 = 322
                                        _ = v47
                                    } else {
                                        if v12 {
                                            var v48 float64 = (float64((float32(1.5) / v3)) - (-2.25 + float64(((0.1 / v3) / (v3 / v3)))))
                                            _ = v48
                                        }
                                        var v49 uint8 = 0
                                        fmt.Println(v3, v3)
                                        _ = v49
                                    }
                                }
                                v50 := GenS1{F0: -1, F1: (-(v1)), F2: v42}
                                _ = v50
                            }
                        } else {
                            var v51 int32 = (v42 & (v42 % (((v42 % ((int32(fi13)) | 1))) | 1)))
                            if ((!v12) && v12) {
                                switch v4 {
                                case "\n\t":
                                    if true {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        v6 = (v6 - v6)
                                        fmt.Println(v14, v1)
                                    } else {
                                        if true {
                                            fmt.Println(fi15)
                                            v51 = v42
                                        } else {
                                            switch (((((v4 + ((("a" + v4) + (v4 + "a")) + (v4 + (v4 + "foo bar")))) + "foo bar") + v4) + v4) + (v4 + v4)) {
                                            case "":
                                                v5[""] = 4
                                            case "hello":
                                                fmt.Println(v4, fi13, v6)
                                            }
                                            var v52 uint16 = 1
                                            _ = v52
                                        }
                                    }
                                }
                                var v53 uint32 = 65535
                                if (!true) {
                                    delete(v5, "hello")
                                } else {
                                    v54 := GenS8{F0: 128, F1: v53}
                                    _ = v54
                                }
                                _ = v53
                            } else {
                                fmt.Println(v1, v4)
                                switch "hello" {
                                case "foo bar":
                                    if false {
                                        if (int8(v51) < 42) {
                                            v12 = (false || v12)
                                        }
                                    }
                                case "":
                                    _ = len(v7)
                                }
                                fmt.Println(v2)
                            }
                            _ = v51
                        }
                        _ = v42
                    }
                    _ = v19
                }
                if v12 {
                    v55 := GenS1{F0: v2, F1: 32767, F2: 0}
                    _ = len(v5)
                    if true {
                        v7 = append(v7, fi13)
                    } else {
                        v1 += (^(v1))
                        if false {
                            if len(v7) > 0 {
                                v7[0] = (fi15 + fi15)
                            }
                            v56 := GenS9{F0: 128, F1: 42}
                            _ = v56
                        }
                        var v57 int = fi13
                        _ = v57
                    }
                    _ = v55
                } else {
                    v2 = v2
                    switch (v4 + "") {
                    case "a":
                        fmt.Println(v1, v1, v4)
                    case "\n\t":
                        v58 := GenS10{F0: true, F1: "foo bar", F2: v14}
                        _ = v58
                    default:
                        var v59 uint64 = (v14 % (((v14 << ((42) & 63))) | 1))
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        _ = v59
                    }
                }
                var v60 uint32 = 311
                if v12 {
                    v3 = (((float32((float64(1.5) - 1.5)) / (((v3 - (v3 - v3)) + 1e10) / (v3 * v3))) * v3) / 1.0)
                } else {
                    switch v4 {
                    case "a":
                        v61 := v5[v4]
                        _ = v61
                    case "hello":
                        _ = len(v5)
                    default:
                        _ = len(v7)
                    }
                }
                switch v4 {
                case "hello":
                    switch 2147483647 {
                    case 347:
                        v14 -= v14
                    case 128:
                        switch (v4 + "") {
                        case "hello":
                            if (!((uint32(v14) > v60) || (v12 || v12))) {
                                var v62 int32 = 42
                                if v12 {
                                    fmt.Println(fi13, v4, v2)
                                    var v63 int64 = v2
                                    _ = v63
                                }
                                if len(v7) > 0 {
                                    v7[0] = 0
                                }
                                _ = v62
                            }
                        default:
                            var v64 float32 = v3
                            _ = v64
                        }
                    default:
                        v65 := GenS8{F0: int(v1), F1: (v60 * v60)}
                        _ = v65
                    }
                case "\n\t":
                    fmt.Println(fi15, v3, v2)
                }
                var v66 uint = 712
                if (!("" == "hello")) {
                    v67 := GenS7{F0: -137, F1: 42}
                    fmt.Println(v6, v3, v1)
                    switch v4 {
                    case "foo bar":
                        if (-1 > -200) {
                            v68 := GenS11{F0: 19, F1: (^(fi15))}
                            _ = v68
                        } else {
                            v14 += v14
                            v7 = append(v7, fi15)
                            v5["\n\t"] = v60
                        }
                    case "a":
                        var v69 uint8 = 164
                        _ = v69
                    }
                    _ = v67
                } else {
                    var v70 float64 = v6
                    switch 2147483647 {
                    case 9223372036854775807:
                        var v71 float32 = v3
                        _ = v71
                    case -283:
                        if (!v12) {
                            var v72 int64 = (v2 << ((65535) & 63))
                            _ = v72
                        }
                    case 257:
                        var v73 int = 2147483647
                        _ = v73
                    }
                    _ = v70
                }
                _ = fi15
                _ = v60
                _ = v66
            }
            _ = v14
        } else {
            var v74 float32 = ((v3 / 0.1) + v3)
            fmt.Println(v3, v74, v3)
            var v75 bool = false
            _ = v74
            _ = v75
        }
        var v76 int8 = 1
        _ = fi13
        _ = v76
    }
    for fi77 := 0; fi77 < 4; fi77++ {
        for fi78 := 0; fi78 < 17; fi78++ {
            if ((587 == 1) || (false || (-1 >= 638))) {
                var v79 bool = (!true)
                _ = v79
            }
            _ = fi78
        }
        if true {
            if (-1 >= -887) {
                v2 = (v2 | v2)
            }
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi77
    }
    if (!(883 != 2)) {
        v80 := GenS12{F0: "foo bar"}
        var v81 uint16 = 4
        _ = v80
        _ = v81
    } else {
        var v82 int16 = (v1 | -984)
        var v83 int8 = 100
        if true {
            for fi84 := 0; fi84 < 8; fi84++ {
                switch ("a" + v4) {
                case "foo bar":
                    v82 = v1
                case "":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                case "hello":
                    var v85 uint8 = 2
                    _ = v85
                }
                v2 = 42
                v86 := GenS13{F0: (42 == 128), F1: (fi84 * (fi84 & (-(fi84))))}
                _ = fi84
                _ = v86
            }
            fmt.Println(v83)
        }
        _ = v82
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
