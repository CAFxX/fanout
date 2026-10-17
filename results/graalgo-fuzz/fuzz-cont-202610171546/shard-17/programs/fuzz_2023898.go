package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int32; F1 bool; F2 string }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 32767
    var v3 uint64 = 18446744073709551615
    var v4 int32 = -191
    v2 -= v2
    v5 := GenS1{F0: (v4 / ((v4) | 1)), F1: v3, F2: 4294967295}
    var v6 bool = false
    _ = len(v1)
    v7 := []int{}
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: v4, F1: (v3 + 42), F2: (fi8 | 128)}
        switch fi8 {
        case 127:
            v10 := GenS2{F0: 0}
            _ = v10
        default:
            v11 := GenS3{F0: 4294967295, F1: ((v3 ^ 18446744073709551615) >= (v3 * 42))}
            _ = v11
        }
        if len(v7) > 0 {
            v7[0] = (^(fi8))
        }
        _ = fi8
        _ = v9
    }
    switch "foo bar" {
    case "":
        var v12 uint = 2
        _ = v12
    }
    var v13 bool = (false || (633 < 4294967295))
    var v14 int32 = (v4 << ((18446744073709551615) & 63))
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint32 = 42
        for fi17 := 0; fi17 < 6; fi17++ {
            if len(v7) > 0 {
                v7[0] = fi15
            }
            switch "" {
            case "a":
                var v18 uint16 = v2
                if v13 {
                    var v19 bool = (((0.1 + ((0.1 + 0.0) + ((0.1 - 0.1) / float64(0.1)))) + 1e10) <= 0.0)
                    v3 -= v3
                    _ = v19
                } else {
                    if v13 {
                        if v6 {
                            switch (fi15 >> ((18446744073709551615) & 63)) {
                            case 4294967295:
                                if (v6 && true) {
                                    v20 := GenS2{F0: -102}
                                    if (v4 == (v14 >> ((2) & 63))) {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    }
                                    _ = v20
                                } else {
                                    var v21 uint = 2
                                    if (1 > 42) {
                                        v4 -= v14
                                        var v22 int16 = 32767
                                        fmt.Println(v21)
                                        _ = v22
                                    } else {
                                        var v23 uint8 = uint8(v2)
                                        _ = v23
                                    }
                                    _ = v21
                                }
                            }
                        } else {
                            if false {
                                v4 = (v4 ^ (v4 * v14))
                                if (1 <= 1) {
                                    v24 := v1[641]
                                    v24 += fi17
                                    fmt.Println(fi17, v4)
                                    switch "" {
                                    case "a":
                                        var v25 bool = v6
                                        _ = v25
                                    }
                                    fmt.Println(v4)
                                    _ = v24
                                }
                                if len(v7) > 0 {
                                    v7[0] = 9223372036854775807
                                }
                            } else {
                                fmt.Println(v16, fi17)
                                var v26 float32 = (-2.25 * 1e10)
                                switch fi15 {
                                case -1:
                                    var v27 uint16 = v2
                                    _ = v27
                                case 32768:
                                    fmt.Println(v26)
                                case -9223372036854775808:
                                    if v13 {
                                        _ = len(v1)
                                        v1[(fi15 | fi15)] = fi15
                                    } else {
                                        var v28 int32 = v14
                                        v16 = v16
                                        _ = v28
                                    }
                                }
                                v1[(fi15 << ((382) & 63))] = (-(fi15))
                                var v29 uint = 18446744073709551615
                                _ = v26
                                _ = v29
                            }
                            if len(v7) > 0 {
                                v7[0] = 127
                            }
                            var v30 uint = uint(v16)
                            v4 -= (v4 | -1)
                            if (false && (v13 && v13)) {
                                v31 := GenS2{F0: 70}
                                switch (fi15 % ((fi17) | 1)) {
                                case 1:
                                    v6 = false
                                case 4294967295:
                                    var v32 uint = (v30 & v30)
                                    _ = v32
                                default:
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                                if v13 {
                                    var v33 float64 = (3.14159 * 0.0)
                                    v7 = append(v7, fi17)
                                    fmt.Println(v4, v18)
                                    _ = v33
                                }
                                _ = v31
                            } else {
                                delete(v1, -50)
                                switch fi17 {
                                case 32767:
                                    _ = len(v1)
                                case 42:
                                    if true {
                                        var v34 float32 = float32(1.5)
                                        _ = v34
                                    } else {
                                        var v35 bool = true
                                        fmt.Println(fi17)
                                        _ = v35
                                    }
                                }
                                if v13 {
                                    var v36 int32 = v14
                                    _ = v36
                                }
                            }
                            _ = v30
                        }
                        if v6 {
                            switch fi15 {
                            case -2147483648:
                                if (!v6) {
                                    var v37 bool = v6
                                    switch (fi15 % ((fi17) | 1)) {
                                    case -597:
                                        v38 := GenS1{F0: -2147483648, F1: (v3 << ((738) & 63)), F2: 2147483647}
                                        _ = v38
                                    case 32768:
                                        var v39 int64 = -9223372036854775808
                                        _ = v39
                                    default:
                                        var v40 int8 = 127
                                        _ = v40
                                    }
                                    var v41 uint16 = (v18 << ((1) & 63))
                                    v42 := GenS4{F0: 0, F1: (float64((0.0 / -2.25)) <= 1e10), F2: "hello"}
                                    _ = v37
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 float32 = 1e10
                                    _ = v43
                                }
                            case 32768:
                                var v44 int = (fi15 << ((74) & 63))
                                _ = v44
                            case 128:
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v2, fi17)
                                case "":
                                    if (uint16(v16) != v2) {
                                        if ((0 <= 42) && (0.1 != (-2.25 * 0.1))) {
                                            fmt.Println(v2)
                                            v45 := GenS2{F0: 42}
                                            v46 := GenS1{F0: v14, F1: v3, F2: fi17}
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                case "hello":
                                    var v47 uint = 1
                                    _ = v47
                                }
                            default:
                                var v48 float64 = float64(1.5)
                                _ = v48
                            }
                            if (true && (572 <= 1)) {
                                v49 := GenS5{F0: 713, F1: -165}
                                if v13 {
                                    v2 -= 2
                                }
                                _ = v49
                            } else {
                                var v50 string = "hello"
                                v51 := GenS2{F0: 27}
                                v4 -= 680
                                if false {
                                    var v52 int64 = -687
                                    switch ("foo bar" + (v50 + ((v50 + v50) + (v50 + v50)))) {
                                    case "\n\t":
                                        var v53 uint64 = v3
                                        _ = v53
                                    }
                                    _ = v52
                                } else {
                                    fmt.Println(v16, v6)
                                    switch (fi15 / ((4294967295) | 1)) {
                                    case 2147483647:
                                        v6 = ((v3 ^ v3) >= (v3 & 471))
                                    case -506:
                                        delete(v1, (fi15 * 340))
                                    default:
                                        if true {
                                            v54 := GenS6{F0: 1000, F1: -925}
                                            if (!v13) {
                                                v50 = "\n\t"
                                                if (32767 >= -32768) {
                                                    if v6 {
                                                        fmt.Println(v18, fi15, v6)
                                                        v55 := GenS7{F0: v4, F1: ((fi15 * fi15) <= (fi17 * (fi17 | (fi15 - (fi17 << ((1) & 63)))))), F2: v50}
                                                        _ = v55
                                                    }
                                                }
                                                switch v50 {
                                                case "hello":
                                                    v56 := GenS8{F0: 303, F1: -49}
                                                    _ = v56
                                                case "\n\t":
                                                    v7 = append(v7, (-(fi15)))
                                                default:
                                                    v57 := v1[(fi17 ^ (^(fi15)))]
                                                    v7 = append(v7, (fi15 / ((-9223372036854775808) | 1)))
                                                    _ = v57
                                                }
                                                var v58 int8 = 42
                                                var v59 uint8 = 127
                                                _ = v58
                                                _ = v59
                                            }
                                            v2 = (v2 << ((1) & 63))
                                            _ = v54
                                        }
                                        var v60 int = fi17
                                        v61 := GenS9{F0: (^(v16))}
                                        _ = v60
                                        _ = v61
                                    }
                                }
                                v62 := v1[(fi17 - 413)]
                                switch v50 {
                                case "a":
                                    v63 := GenS10{F0: -1, F1: "foo bar"}
                                    _ = v63
                                case "hello":
                                    if ((^(v16)) > 2147483647) {
                                        switch v62 {
                                        case 579:
                                            switch fi15 {
                                            case 31:
                                                fmt.Println(v14, v50, fi17)
                                            default:
                                                var v64 int8 = -128
                                                _ = v64
                                            }
                                        default:
                                            switch (v62 >> ((42) & 63)) {
                                            case 97:
                                                fmt.Println(fi17, v14)
                                            case 9223372036854775807:
                                                v65 := GenS11{F0: uint16(v16), F1: ""}
                                                _ = v65
                                            case 128:
                                                v66 := GenS11{F0: 65535, F1: ""}
                                                _ = v66
                                            default:
                                                if (!(-1 >= v14)) {
                                                    v67 := GenS10{F0: -131, F1: v50}
                                                    _ = v67
                                                }
                                            }
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = v62
                                        }
                                        fmt.Println(v6, v3)
                                    }
                                default:
                                    var v68 uint16 = 32767
                                    _ = v68
                                }
                                v69 := GenS5{F0: (v16 << ((4294967295) & 63)), F1: -9223372036854775808}
                                v70 := GenS5{F0: 997, F1: 0}
                                _ = v50
                                _ = v51
                                _ = v62
                                _ = v69
                                _ = v70
                            }
                            v18 -= v18
                        } else {
                            v3 = 653
                        }
                        switch ("a" + "a") {
                        case "hello":
                            fmt.Println(v4)
                        case "":
                            var v71 uint16 = 1
                            _ = v71
                        }
                    }
                }
                _ = v18
            case "":
                v72 := GenS2{F0: 42}
                _ = v72
            }
            fmt.Println(v3, fi15, v16)
            v2 -= (v2 % (((v2 ^ (v2 ^ (v2 / (((^(v2))) | 1))))) | 1))
            fmt.Println(v13)
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    switch -9223372036854775808 {
    case 168:
        var v73 int = -2147483648
        _ = v73
    case 0:
        var v74 int8 = 127
        _ = v74
    case 1:
        var v75 uint = 0
        _ = v75
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
