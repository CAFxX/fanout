package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint16; F2 int64 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int; F2 uint8 }
type GenS8 struct { F0 int8; F1 uint; F2 uint16 }

func main() {
    v1 := []uint32{42}
    var v2 float64 = float64((((-2.25 + 3.14159) * (-2.25 - 1e10)) * 1.5))
    var v3 uint32 = 2147483647
    var v4 int16 = 0
    var v5 int = -152
    v6 := make(map[uint32]int)
    _ = len(v1)
    v5 += 9223372036854775807
    var v7 bool = (255 <= 127)
    if true {
        v5 -= (^(v5))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 = v3
    var v8 uint8 = 0
    v1 = append(v1, v3)
    v5 = -535
    switch "hello" {
    case "hello":
        v2 -= v2
    case "foo bar":
        for fi9 := 0; fi9 < 9; fi9++ {
            if (!v7) {
                _ = len(v1)
                v10 := GenS1{F0: (263 > (v3 | (v3 % (((v3 / ((v3) | 1))) | 1))))}
                _ = v10
            } else {
                var v11 uint16 = 255
                var v12 int32 = 5
                for fi13 := 0; fi13 < 4; fi13++ {
                    switch "foo bar" {
                    case "\n\t":
                        switch "" {
                        case "hello":
                            var v14 uint8 = (^(v8))
                            _ = v14
                        case "a":
                            var v15 float32 = float32(float64(0.0))
                            _ = v15
                        case "foo bar":
                            var v16 bool = (!v7)
                            switch ("hello" + "a") {
                            case "":
                                v11 -= v11
                            case "foo bar":
                                v17 := GenS1{F0: (v16 && (2147483647 < v12))}
                                _ = v17
                            default:
                                _ = len(v1)
                            }
                            if true {
                                if v16 {
                                    var v18 int64 = 1
                                    v5 -= (fi13 >> ((1) & 63))
                                    if false {
                                        var v19 int64 = v18
                                        switch "hello" {
                                        case "\n\t":
                                            var v20 int16 = 32767
                                            _ = v20
                                        case "foo bar":
                                            v16 = false
                                        case "":
                                            if (!(382 != 739)) {
                                                if (v2 < v2) {
                                                    v18 += (v18 ^ -9223372036854775808)
                                                }
                                            } else {
                                                fmt.Println(v2, v18, v7)
                                                fmt.Println(v7, v16, v2)
                                            }
                                        default:
                                            v4 = v4
                                            v12 = (v12 + (v12 | 0))
                                            switch (fi13 & (fi13 >> ((1) & 63))) {
                                            case 32768:
                                                if (-16 <= -1) {
                                                    switch (fi13 & (fi9 & fi13)) {
                                                    case -9223372036854775808:
                                                        var v21 uint32 = 2
                                                        _ = v21
                                                    case 127:
                                                        v1 = append(v1, 42)
                                                    case 2147483647:
                                                        var v22 string = "\n\t"
                                                        _ = v22
                                                    default:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v23 := GenS1{F0: true}
                                                        if (!v16) {
                                                            v4 = (^(v4))
                                                            if false {
                                                                switch ((((("\n\t" + "a") + "hello") + "") + "foo bar") + ("" + "hello")) {
                                                                case "a":
                                                                    var v24 int16 = v4
                                                                    _ = v24
                                                                }
                                                                var v25 uint8 = (v8 / ((255) | 1))
                                                                if v16 {
                                                                    v26 := GenS2{F0: v16}
                                                                    _ = v26
                                                                } else {
                                                                    fmt.Println(v19, v2)
                                                                }
                                                                switch "hello" {
                                                                case "":
                                                                    var v27 float32 = 1e10
                                                                    _ = v27
                                                                case "hello":
                                                                    fmt.Println(v11)
                                                                case "\n\t":
                                                                    v16 = ("" > "foo bar")
                                                                default:
                                                                    v25 = 0
                                                                }
                                                                switch (v5 | fi9) {
                                                                case 127:
                                                                    var v28 int = v5
                                                                    _ = v28
                                                                default:
                                                                    v16 = (!false)
                                                                    v29 := GenS3{F0: 93}
                                                                    v30 := GenS1{F0: (2 < v11)}
                                                                    _ = v29
                                                                    _ = v30
                                                                }
                                                                _ = v25
                                                            }
                                                            var v31 uint64 = 343
                                                            _ = v31
                                                        }
                                                        _ = v23
                                                    }
                                                }
                                            case -9223372036854775808:
                                                fmt.Println(v5, v16, v4)
                                            case 9223372036854775807:
                                                if (-424 <= v12) {
                                                    switch (v5 >> ((255) & 63)) {
                                                    case 32767:
                                                        var v32 int16 = (v4 - v4)
                                                        _ = v32
                                                    case 42:
                                                        v33 := GenS1{F0: v7}
                                                        _ = v33
                                                    }
                                                    if (!(v3 >= (v3 & (v3 % ((2147483647) | 1))))) {
                                                        if false {
                                                            if (1 >= 759) {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v34 := GenS4{F0: (v8 + 128), F1: 971}
                                                                _ = v34
                                                            } else {
                                                                var v35 float32 = 0.0
                                                                _ = v35
                                                            }
                                                            if true {
                                                                if true {
                                                                    if v7 {
                                                                        v1 = append(v1, (^(v3)))
                                                                        var v36 uint = 65535
                                                                        _ = v36
                                                                    }
                                                                } else {
                                                                    fmt.Println(v19, fi13, v11)
                                                                    var v37 uint32 = uint32(v8)
                                                                    _ = v37
                                                                }
                                                            } else {
                                                                v19 -= v19
                                                                var v38 float64 = v2
                                                                var v39 int = fi9
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            v18 = -1
                                                        }
                                                        if v7 {
                                                            v40 := GenS2{F0: ((v12 & v12) > (v12 & (v12 * v12)))}
                                                            _ = v40
                                                        } else {
                                                            v19 -= (v19 << ((18446744073709551615) & 63))
                                                        }
                                                        v19 += (v18 / (((v19 - int64(v4))) | 1))
                                                    } else {
                                                        switch "\n\t" {
                                                        case "":
                                                            v8 = (^(v8))
                                                        case "foo bar":
                                                            switch ("" + "\n\t") {
                                                            case "\n\t":
                                                                fmt.Println(fi13)
                                                            case "hello":
                                                                var v41 float64 = 3.14159
                                                                _ = v41
                                                            default:
                                                                switch fi9 {
                                                                case -151:
                                                                    v42 := GenS5{F0: v12, F1: 701, F2: v19}
                                                                    _ = v42
                                                                case -9223372036854775808:
                                                                    if (v11 >= 59) {
                                                                        v43 := GenS6{F0: v16, F1: (v11 / ((568) | 1))}
                                                                        v44 := GenS7{F0: v3, F1: v5, F2: (v8 / ((128) | 1))}
                                                                        if (v12 == 0) {
                                                                            var v45 uint8 = (v8 & 127)
                                                                            if ((v11 & v11) != (^(v11))) {
                                                                                var v46 float64 = 1.5
                                                                                v19 = 1
                                                                                fmt.Println(v7)
                                                                                var v47 uint16 = (v11 ^ v11)
                                                                                var v48 int32 = 2147483647
                                                                                _ = v46
                                                                                _ = v47
                                                                                _ = v48
                                                                            }
                                                                            fmt.Println(v3, v16, v45)
                                                                            _ = v45
                                                                        } else {
                                                                            switch fi13 {
                                                                            case -196:
                                                                                v16 = (160 < 32767)
                                                                            default:
                                                                                v6[630] = v5
                                                                                fmt.Println(v3, v12, v4)
                                                                            }
                                                                        }
                                                                        _ = v43
                                                                        _ = v44
                                                                    } else {
                                                                        v7 = ("hello" >= "")
                                                                        v49 := GenS8{F0: 0, F1: 18446744073709551615, F2: (v11 / (((v11 * v11)) | 1))}
                                                                        var v50 int32 = 42
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                case 307:
                                                                    v18 = int64(v8)
                                                                }
                                                            }
                                                        default:
                                                            v5 = v5
                                                            v51 := GenS2{F0: true}
                                                            _ = v51
                                                        }
                                                    }
                                                }
                                            }
                                            var v52 float32 = 0.1
                                            _ = v52
                                        }
                                        _ = v19
                                    }
                                    delete(v6, (v3 & (v3 - (^(v3)))))
                                    _ = v18
                                }
                                v4 += v4
                                v8 = v8
                            } else {
                                var v53 uint8 = (^(v8))
                                var v54 uint = 2
                                var v55 int64 = 9223372036854775807
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            _ = v16
                        default:
                            if ((!v7) || (!((1.5 + (0.0 - 3.14159)) < ((1.5 - 0.0) + 3.14159)))) {
                                if v7 {
                                    if v7 {
                                        if v7 {
                                            if true {
                                                if (!v7) {
                                                    _ = len(v6)
                                                    v12 = (v12 + (^(v12)))
                                                } else {
                                                    switch ("a" + "") {
                                                    case "":
                                                        fmt.Println(v5, fi13)
                                                    default:
                                                        v1 = append(v1, 65535)
                                                    }
                                                    v4 += (v4 % ((v4) | 1))
                                                    v56 := v6[v3]
                                                    _ = v56
                                                }
                                                fmt.Println(v2, v4, v3)
                                            } else {
                                                if (2 > 18446744073709551615) {
                                                    _ = len(v6)
                                                } else {
                                                    v57 := v6[v3]
                                                    _ = v57
                                                }
                                                if v7 {
                                                    if (((1e10 - -2.25) == float32(1e10)) && (851 < 18446744073709551615)) {
                                                        fmt.Println(fi13, v3, v12)
                                                        if len(v1) > 0 {
                                                            v1[0] = v3
                                                        }
                                                    } else {
                                                        v12 -= 0
                                                    }
                                                } else {
                                                    fmt.Println(fi13, fi9)
                                                    if false {
                                                        v11 = v11
                                                    }
                                                }
                                                fmt.Println(v3, fi9)
                                            }
                                            var v58 uint8 = v8
                                            _ = v58
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = v3
                                            }
                                        }
                                    }
                                    var v59 int16 = -1
                                    _ = v59
                                }
                            } else {
                                var v60 int32 = (^(v12))
                                fmt.Println(v3, v8)
                                var v61 float64 = (v2 / (((v2 * v2) + v2) / (v2 / v2)))
                                fmt.Println(v12)
                                v62 := GenS6{F0: ((v12 % ((42) | 1)) == v12), F1: (v11 & (v11 * 680))}
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                        }
                    case "hello":
                        fmt.Println(v4, v4, v5)
                        var v63 bool = true
                        var v64 uint64 = 42
                        _ = v63
                        _ = v64
                    case "foo bar":
                        var v65 float64 = float64((0.1 * 1e10))
                        _ = v65
                    default:
                        v3 = v3
                    }
                    _ = fi13
                }
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    case "\n\t":
        switch "a" {
        case "\n\t":
            for fi66 := 0; fi66 < 20; fi66++ {
                v67 := v6[v3]
                for fi68 := 0; fi68 < 6; fi68++ {
                    var v69 uint16 = 1
                    _ = fi68
                    _ = v69
                }
                var v70 uint = 0
                _ = fi66
                _ = v67
                _ = v70
            }
        }
    }
    for fi71 := 0; fi71 < 4; fi71++ {
        v3 = v3
        v72 := GenS6{F0: (2 == 2), F1: 32767}
        switch v5 {
        case 0:
            var v73 uint32 = (v3 | 42)
            _ = v73
        case 452:
            v5 -= (v5 | (fi71 + fi71))
        default:
            var v74 uint32 = v3
            _ = v74
        }
        _ = fi71
        _ = v72
    }
    switch 128 {
    case 128:
        fmt.Println(v8, v2)
    case 722:
        if ((0 <= 145) && false) {
            v7 = ((v7 || v7) && (698 != 737))
        }
    case -949:
        v75 := v6[v3]
        _ = v75
    }
    v2 += v2
    var v76 float64 = (1e10 - (v2 * float64(-2.25)))
    v7 = (v7 && true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v76:", v76)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v8
}
