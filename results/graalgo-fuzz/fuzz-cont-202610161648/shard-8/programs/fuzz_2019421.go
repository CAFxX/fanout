package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int16; F1 string; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 string; F1 uint32; F2 string }
type GenS9 struct { F0 int; F1 bool; F2 uint16 }

func main() {
    v1 := []uint32{2147483647}
    var v2 uint32 = 2147483648
    var v3 float64 = float64(1e10)
    v4 := make(map[string]uint32)
    var v5 int8 = 1
    var v6 uint64 = 165
    switch "foo bar" {
    case "foo bar":
        switch -71 {
        case 695:
            if false {
                v3 -= v3
                if false {
                    for fi7 := 0; fi7 < 13; fi7++ {
                        v3 += 3.14159
                        _ = fi7
                    }
                    v5 = v5
                    v8 := GenS1{F0: 679, F1: 0}
                    _ = v8
                }
            } else {
                v5 += (v5 * (-(v5)))
                v9 := v4["foo bar"]
                _ = v9
            }
        case -9223372036854775808:
            v5 -= 1
        }
    case "hello":
        switch -9223372036854775808 {
        case 359:
            _ = len(v1)
        }
    case "":
        switch 32768 {
        case 2147483647:
            v3 = 3.14159
        case -9223372036854775808:
            var v10 int16 = 182
            _ = v10
        case 32767:
            v1 = append(v1, (v2 % (((^(v2))) | 1)))
        default:
            var v11 bool = (false || false)
            _ = v11
        }
    default:
        if false {
            v6 = v6
            var v12 uint8 = 158
            for fi13 := 0; fi13 < 16; fi13++ {
                _ = len(v4)
                v5 = (-(v5))
                _ = fi13
            }
            _ = v12
        }
    }
    var v14 uint32 = v2
    switch "hello" {
    case "a":
        switch 32767 {
        case -569:
            var v15 int8 = (v5 + 0)
            _ = v15
        case 42:
            v5 -= 1
        case -1:
            switch -605 {
            case 42:
                v16 := GenS2{F0: v6, F1: v2, F2: 222}
                _ = v16
            case -2147483648:
                for fi17 := 0; fi17 < 6; fi17++ {
                    v5 += (-(v5))
                    _ = fi17
                }
            case 127:
                if (true && ("foo bar" != "foo bar")) {
                    if (-2.25 > float32((1.5 - float64((((float32((v3 - 0.0)) - 1.5) / ((1e10 - -2.25) / 1.0)) / 1.0))))) {
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v19 := GenS3{F0: 856, F1: 138, F2: v6}
                            _ = len(v1)
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        var v20 uint16 = 1
                        fmt.Println(v20, v14, v3)
                        _ = v20
                    }
                    v5 -= v5
                    v2 = (v14 & (^(v14)))
                }
            default:
                v21 := v4[("\n\t" + "hello")]
                _ = v21
            }
        default:
            v4["a"] = v14
        }
    case "hello":
        var v22 int = 42
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 3; fi23++ {
            fmt.Println(v6, v3, v5)
            _ = fi23
        }
    default:
        if (0 < -32768) {
            fmt.Println(v6, v2)
        }
    }
    v24 := GenS4{F0: 2147483647, F1: 1}
    var v25 int = -179
    var v26 uint8 = 0
    v26 -= 2
    switch "a" {
    case "":
        v27 := GenS4{F0: 42, F1: v6}
        _ = v27
    case "a":
        if (int32(v14) == -484) {
            _ = len(v1)
            for fi28 := 0; fi28 < 4; fi28++ {
                for fi29 := 0; fi29 < 3; fi29++ {
                    v30 := GenS2{F0: 1, F1: v14, F2: (v26 & v26)}
                    v5 -= v5
                    if (false || (-716 > 59)) {
                        var v31 uint16 = 255
                        v2 += v2
                        var v32 uint16 = (^(v31))
                        _ = v31
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
            v3 = v3
        } else {
            v6 = v6
            if true {
                var v33 uint64 = v6
                switch "" {
                case "foo bar":
                    var v34 float32 = float32(float64(1.5))
                    _ = v34
                case "a":
                    v25 -= 9223372036854775807
                }
                _ = v33
            }
            switch v25 {
            case 1:
                if false {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 uint = 707
                        v1 = append(v1, 2)
                        if true {
                            var v37 uint64 = uint64(v36)
                            fmt.Println(fi35)
                            switch "hello" {
                            case "\n\t":
                                v38 := GenS3{F0: 1, F1: (v26 | v26), F2: 588}
                                _ = v38
                            case "foo bar":
                                v14 += 1
                            default:
                                fmt.Println(v3, v2, v6)
                            }
                            _ = v37
                        } else {
                            _ = len(v4)
                            var v39 string = ("\n\t" + "")
                            _ = v39
                        }
                        _ = fi35
                        _ = v36
                    }
                    v40 := GenS3{F0: 65535, F1: (v26 % (((v26 | v26)) | 1)), F2: v6}
                    _ = v40
                }
            case -2147483648:
                v1 = append(v1, v2)
            case 47:
                var v41 string = ""
                _ = v41
            }
        }
    }
    var v42 uint64 = v6
    v26 -= 183
    if false {
        for fi43 := 0; fi43 < 13; fi43++ {
            for fi44 := 0; fi44 < 9; fi44++ {
                var v45 uint64 = v42
                _ = fi44
                _ = v45
            }
            if true {
                for fi46 := 0; fi46 < 5; fi46++ {
                    v42 = uint64(v26)
                    if (false || true) {
                        if ((v3 / v3) < (((v3 / ((v3 * (3.14159 + (v3 - 1.5))) + (-2.25 + float64(0.1)))) / float64((1e10 / (float32(1e10) / ((((-2.25 - 0.0) / (-2.25 + (0.1 * (((0.1 / 1.0) + (1.5 * -2.25)) / 1.0)))) * -2.25) + 0.1))))) + float64(0.1))) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v47 uint32 = (^(v14))
                            var v48 int8 = 1
                            _ = v47
                            _ = v48
                        } else {
                            var v49 uint16 = 445
                            v50 := GenS5{F0: 338, F1: (("foo bar" + ((("a" + "a") + "\n\t") + "foo bar")) + ""), F2: ((v5 - v5) != v5)}
                            v3 = ((v3 - 0.1) * 0.0)
                            _ = v49
                            _ = v50
                        }
                        v14 -= 65535
                        switch "\n\t" {
                        case "":
                            var v51 uint64 = (v42 + 319)
                            _ = v51
                        case "hello":
                            fmt.Println(fi46, v5, fi46)
                        default:
                            v6 -= (^(v42))
                        }
                    } else {
                        v52 := GenS1{F0: 65535, F1: 9223372036854775807}
                        v53 := GenS6{F0: 1000}
                        _ = v52
                        _ = v53
                    }
                    v25 -= (fi46 ^ v25)
                    _ = fi46
                }
            } else {
                if (0 <= v14) {
                    v2 = (v2 + (v2 - v14))
                    fmt.Println(v26)
                    if true {
                        v2 -= v2
                    } else {
                        for fi54 := 0; fi54 < 14; fi54++ {
                            fmt.Println(v14, v2, v25)
                            if false {
                                v55 := v4["\n\t"]
                                v56 := GenS7{F0: 159, F1: 1}
                                switch "\n\t" {
                                case "":
                                    v1 = append(v1, v55)
                                case "foo bar":
                                    if (((v26 + 1) > v26) || ("foo bar" == "")) {
                                        v4["a"] = (^(v2))
                                    } else {
                                        var v57 uint32 = (v2 & v2)
                                        _ = v57
                                    }
                                case "hello":
                                    v5 += 127
                                default:
                                    v55 = (v55 - (^(v14)))
                                }
                                _ = v55
                                _ = v56
                            } else {
                                v3 -= v3
                                switch "" {
                                case "a":
                                    switch "a" {
                                    case "\n\t":
                                        _ = len(v1)
                                    default:
                                        switch fi43 {
                                        case -1:
                                            fmt.Println(v26)
                                        default:
                                            fmt.Println(v3)
                                        }
                                    }
                                case "hello":
                                    var v58 uint8 = v26
                                    _ = v58
                                case "":
                                    var v59 float64 = (v3 * ((v3 * v3) * v3))
                                    _ = v59
                                }
                                v60 := GenS7{F0: 2, F1: 914}
                                _ = v60
                            }
                            var v61 uint64 = (v42 << ((255) & 63))
                            if false {
                                v62 := GenS2{F0: 994, F1: v14, F2: (v26 << ((333) & 63))}
                                if ((v61 | (v6 >> ((831) & 63))) != (v61 % (((v61 >> ((1) & 63))) | 1))) {
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                } else {
                                    var v63 uint16 = 65535
                                    v3 += float64(float32((v3 - v3)))
                                    v1 = append(v1, (v14 - v2))
                                    _ = v63
                                }
                                _ = v62
                            } else {
                                if (v26 <= v26) {
                                    var v64 uint64 = (v42 >> ((2) & 63))
                                    fmt.Println(v2, v64, v14)
                                    v14 -= (v14 & (v14 - v2))
                                    _ = v64
                                }
                            }
                            _ = fi54
                            _ = v61
                        }
                    }
                } else {
                    fmt.Println(v42)
                }
                var v65 int16 = -32768
                _ = v65
            }
            fmt.Println(fi43)
            _ = fi43
        }
        v25 -= v25
        v66 := GenS4{F0: -2147483648, F1: v6}
        _ = v66
    }
    var v67 float64 = v3
    v6 = (v42 >> ((147) & 63))
    var v68 int = (v25 >> ((0) & 63))
    if (!(0 > -1)) {
        if true {
            fmt.Println(v67, v25)
            if true {
                switch v68 {
                case -9223372036854775808:
                    for fi69 := 0; fi69 < 6; fi69++ {
                        for fi70 := 0; fi70 < 2; fi70++ {
                            if false {
                                var v71 bool = (890 != -1)
                                _ = v71
                            } else {
                                if ((((v6 ^ (v6 ^ (^(v6)))) < v6) && (((^(v26)) < (v26 + v26)) && (v26 != v26))) && (!false)) {
                                    if (true || (42 <= v2)) {
                                        var v72 string = "a"
                                        var v73 uint32 = v2
                                        v68 -= fi70
                                        v4[v72] = v2
                                        _ = v72
                                        _ = v73
                                    } else {
                                        var v74 uint16 = 255
                                        fmt.Println(v6, v3, fi69)
                                        _ = v74
                                    }
                                    v75 := GenS8{F0: ("\n\t" + "a"), F1: v14, F2: "hello"}
                                    _ = v75
                                } else {
                                    v76 := GenS5{F0: 32767, F1: "foo bar", F2: (1 != int16(v2))}
                                    v2 -= (v2 >> ((65535) & 63))
                                    _ = v76
                                }
                                switch ("a" + "a") {
                                case "foo bar":
                                    var v77 int16 = int16(v2)
                                    v78 := v4["a"]
                                    _ = v77
                                    _ = v78
                                case "\n\t":
                                    var v79 int32 = 1
                                    _ = v79
                                case "a":
                                    var v80 int64 = 0
                                    _ = v80
                                default:
                                    v26 -= v26
                                }
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        var v81 uint64 = (v42 >> ((2) & 63))
                                        _ = v81
                                    }
                                    switch "a" {
                                    case "":
                                        var v82 int64 = 42
                                        var v83 float32 = -2.25
                                        v1 = append(v1, 0)
                                        _ = v82
                                        _ = v83
                                    case "hello":
                                        v84 := GenS8{F0: ("" + "\n\t"), F1: v14, F2: "a"}
                                        _ = v84
                                    }
                                } else {
                                    var v85 int8 = 42
                                    _ = v85
                                }
                            }
                            _ = fi70
                        }
                        switch "hello" {
                        case "":
                            v6 += (v42 & 0)
                        }
                        v86 := GenS4{F0: -1, F1: v42}
                        _ = fi69
                        _ = v86
                    }
                }
                if ((v68 >> ((0) & 63)) >= -774) {
                    v42 = 0
                }
                for fi87 := 0; fi87 < 18; fi87++ {
                    for fi88 := 0; fi88 < 5; fi88++ {
                        fmt.Println(v42, v67)
                        _ = fi88
                    }
                    _ = fi87
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    var v89 float64 = v67
                    _ = v89
                default:
                    v5 -= (v5 << ((825) & 63))
                }
            }
        }
    } else {
        for fi90 := 0; fi90 < 8; fi90++ {
            switch (v68 * (v68 * fi90)) {
            case -9223372036854775808:
                var v91 float64 = (v67 / 1.0)
                _ = v91
            case -830:
                if ((2147483647 <= -913) || false) {
                    for fi92 := 0; fi92 < 8; fi92++ {
                        delete(v4, "foo bar")
                        var v93 float64 = 1.5
                        _ = fi92
                        _ = v93
                    }
                    v68 = fi90
                    fmt.Println(v42)
                }
            case 286:
                fmt.Println(v67)
            default:
                var v94 int32 = 0
                _ = v94
            }
            _ = fi90
        }
        var v95 uint16 = 1
        _ = v95
    }
    v96 := GenS9{F0: (v25 >> ((2) & 63)), F1: true, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v96
}
