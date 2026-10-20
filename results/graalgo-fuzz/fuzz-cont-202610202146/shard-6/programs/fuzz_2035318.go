package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 bool; F1 int16; F2 uint16 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 string; F1 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-289, -480}
    v3 := []int{42, -9223372036854775808}
    var v4 uint16 = 830
    v5 := []uint32{2147483648}
    if (((!(-616 != -1)) && (250 < 1)) || false) {
        v6 := GenS1{F0: (v4 + (^(v4)))}
        fmt.Println(v4)
        _ = v6
    } else {
        v4 = (v4 * 65535)
        v4 += (^(v4))
    }
    switch "foo bar" {
    case "foo bar":
        v7 := GenS1{F0: 255}
        _ = v7
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    delete(v1, 0)
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint64 = 18446744073709551615
        for fi10 := 0; fi10 < 7; fi10++ {
            if false {
                v9 += (v9 << ((58) & 63))
                switch -9223372036854775808 {
                case 2147483647:
                    v9 += 2
                case -503:
                    fmt.Println(fi10, fi10, v4)
                case 9223372036854775807:
                    v9 += v9
                    switch fi10 {
                    case 32767:
                        v11 := GenS1{F0: uint16(v9)}
                        _ = v11
                    case -803:
                        var v12 uint16 = v4
                        _ = v12
                    case -9223372036854775808:
                        if true {
                            switch ("" + "hello") {
                            case "hello":
                                var v13 int = 614
                                _ = v13
                            }
                            fmt.Println(v9, fi10, fi8)
                            if len(v3) > 0 {
                                v3[0] = fi8
                            }
                        }
                    default:
                        v9 = v9
                    }
                    _ = len(v1)
                default:
                    if (42 <= 42) {
                        if false {
                            v4 = 983
                        } else {
                            if (3.14159 >= 1.5) {
                                v9 += 42
                                var v14 uint8 = 2
                                v4 = v4
                                _ = v14
                            }
                            if true {
                                v9 = 42
                            } else {
                                var v15 uint64 = (v9 % ((0) | 1))
                                _ = v15
                            }
                            if true {
                                switch fi10 {
                                case 32768:
                                    var v16 int32 = 42
                                    _ = v16
                                case -585:
                                    v9 += 0
                                case 1:
                                    delete(v1, 776)
                                }
                                var v17 uint16 = (v4 % ((v4) | 1))
                                _ = v17
                            } else {
                                v9 = 0
                                if true {
                                    if false {
                                        v4 -= 42
                                    } else {
                                        var v18 uint32 = 4294967295
                                        switch fi10 {
                                        case 958:
                                            var v19 uint8 = 42
                                            _ = v19
                                        case 9223372036854775807:
                                            switch fi10 {
                                            case -258:
                                                v18 -= (v18 + 618)
                                            case 0:
                                                var v20 int32 = -1
                                                _ = v20
                                            default:
                                                var v21 int16 = 1
                                                _ = v21
                                            }
                                        case 735:
                                            if ((true || (1 == uint(v18))) && true) {
                                                switch (fi8 >> ((uint(v4)) & 63)) {
                                                case 42:
                                                    v18 = (v18 << ((255) & 63))
                                                case -643:
                                                    switch ("" + "foo bar") {
                                                    case "":
                                                        v9 += 0
                                                    case "a":
                                                        var v22 int32 = 42
                                                        _ = v22
                                                    }
                                                default:
                                                    switch ("hello" + "foo bar") {
                                                    case "foo bar":
                                                        if ((-128 > 1) && (!false)) {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                switch (fi10 | (^(fi8))) {
                                                                case -276:
                                                                    var v23 uint8 = 207
                                                                    _ = v23
                                                                case 282:
                                                                    var v24 string = "\n\t"
                                                                    _ = v24
                                                                case 706:
                                                                    v25 := GenS1{F0: (v4 + (v4 >> ((42) & 63)))}
                                                                    _ = v25
                                                                default:
                                                                    var v26 int8 = -128
                                                                    _ = v26
                                                                }
                                                            case "foo bar":
                                                                var v27 int = int(v4)
                                                                _ = v27
                                                            default:
                                                                var v28 int = fi10
                                                                _ = v28
                                                            }
                                                        } else {
                                                            if ((0.0 * ((((0.0 - (1e10 - -2.25)) + float32(3.14159)) - (1.5 + ((0.1 / ((1e10 * 1.5) - float32(float64(-2.25)))) * 0.1))) - 0.0)) > (0.1 - -2.25)) {
                                                                v3 = append(v3, fi8)
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                fmt.Println(v18)
                                                            }
                                                            v29 := GenS1{F0: (v4 >> ((42) & 63))}
                                                            fmt.Println(v18)
                                                            _ = v29
                                                        }
                                                    case "hello":
                                                        fmt.Println(v9, v4, fi8)
                                                    case "\n\t":
                                                        _ = len(v5)
                                                    }
                                                }
                                                var v30 uint16 = (v4 & (v4 - v4))
                                                _ = v30
                                            }
                                        default:
                                            switch (fi8 << ((1) & 63)) {
                                            case -9223372036854775808:
                                                var v31 int32 = 42
                                                _ = v31
                                            case -153:
                                                fmt.Println(fi8)
                                            case 0:
                                                var v32 string = ""
                                                var v33 int64 = 319
                                                _ = len(v5)
                                                v33 -= 859
                                                v1[(v18 - v18)] = ""
                                                _ = v32
                                                _ = v33
                                            }
                                            var v34 uint8 = 255
                                            v35 := GenS2{F0: v18, F1: 65535, F2: 494}
                                            _ = len(v1)
                                            _ = v34
                                            _ = v35
                                        }
                                        v4 = (v4 | 1)
                                        _ = v18
                                    }
                                    if (!true) {
                                        v36 := v1[2147483648]
                                        switch fi10 {
                                        case 127:
                                            v36 = (v36 + "\n\t")
                                        case 4294967295:
                                            v9 -= v9
                                        case 70:
                                            var v37 float32 = 1.5
                                            _ = v37
                                        }
                                        _ = v36
                                    }
                                    v4 = v4
                                } else {
                                    v4 -= v4
                                    if len(v2) > 0 {
                                        v2[0] = fi8
                                    }
                                }
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = -564
                        }
                        if (-2147483648 == -1) {
                            _ = len(v5)
                            v38 := v1[65535]
                            v4 -= v4
                            _ = v38
                        }
                    } else {
                        if (v9 < (v9 >> ((255) & 63))) {
                            v39 := GenS3{F0: 255}
                            if (true && ((1e10 / 1.0) >= 1e10)) {
                                switch "a" {
                                case "hello":
                                    v40 := GenS4{F0: (v4 - v4), F1: 2147483647, F2: (fi8 + fi10)}
                                    _ = v40
                                }
                                if false {
                                    delete(v1, 2147483648)
                                    var v41 uint64 = v9
                                    _ = v41
                                }
                                v42 := GenS3{F0: 0}
                                _ = v42
                            }
                            _ = v39
                        }
                    }
                }
            }
            var v43 int8 = 1
            _ = fi10
            _ = v43
        }
        var v44 int8 = 100
        _ = fi8
        _ = v44
        _ = v9
    }
    v45 := GenS5{F0: 186, F1: 415, F2: v4}
    for fi46 := 0; fi46 < 11; fi46++ {
        if (0 > 320) {
            for fi47 := 0; fi47 < 2; fi47++ {
                var v48 float32 = -2.25
                _ = fi47
                _ = v48
            }
            v4 = v4
        } else {
            v49 := GenS6{F0: true, F1: 1, F2: (v4 * (v4 | v4))}
            _ = v49
        }
        for fi50 := 0; fi50 < 19; fi50++ {
            v51 := GenS6{F0: false, F1: -693, F2: 32767}
            if (!((2147483647 != -2147483648) || (!false))) {
                var v52 int8 = -128
                var v53 int8 = v52
                v4 = (v4 * 724)
                _ = v52
                _ = v53
            } else {
                var v54 bool = true
                v2 = append(v2, (fi50 & -887))
                var v55 float32 = (1.5 / 1.0)
                var v56 uint64 = 18446744073709551615
                _ = v54
                _ = v55
                _ = v56
            }
            v3 = append(v3, fi46)
            v57 := GenS6{F0: ((!(42 <= -195)) || false), F1: 32767, F2: 65535}
            _ = fi50
            _ = v51
            _ = v57
        }
        _ = fi46
    }
    var v58 float64 = ((1.5 * 0.1) - 3.14159)
    var v59 string = "\n\t"
    for fi60 := 0; fi60 < 8; fi60++ {
        for fi61 := 0; fi61 < 17; fi61++ {
            if (42 > 100) {
                v62 := GenS2{F0: 37, F1: 0, F2: 500}
                var v63 uint16 = 0
                _ = v62
                _ = v63
            }
            fmt.Println(v59, fi60, fi60)
            fmt.Println(v4, v58)
            _ = fi61
        }
        _ = fi60
    }
    for fi64 := 0; fi64 < 5; fi64++ {
        v4 += v4
        if (1 != v4) {
            if (42 == 18446744073709551615) {
                v58 -= v58
                v58 = v58
                _ = len(v1)
            } else {
                for fi65 := 0; fi65 < 9; fi65++ {
                    switch v59 {
                    case "\n\t":
                        v66 := GenS7{F0: 0, F1: 2}
                        _ = v66
                    case "":
                        if true {
                            var v67 int = fi65
                            _ = v67
                        } else {
                            if true {
                                var v68 int32 = 1
                                _ = v68
                            } else {
                                delete(v1, 4294967295)
                                fmt.Println(v58, v58, v4)
                                v59 = "foo bar"
                            }
                        }
                        v4 -= (^(v4))
                        v69 := GenS3{F0: 4294967295}
                        _ = v69
                    case "hello":
                        v70 := GenS8{F0: (v59 + "\n\t"), F1: "\n\t"}
                        _ = v70
                    }
                    var v71 string = (("hello" + v59) + v59)
                    _ = fi65
                    _ = v71
                }
                var v72 uint16 = 874
                if true {
                    if false {
                        for fi73 := 0; fi73 < 17; fi73++ {
                            var v74 string = "a"
                            if ((true && false) || false) {
                                var v75 int64 = -834
                                v72 = v72
                                var v76 uint32 = 42
                                _ = v75
                                _ = v76
                            } else {
                                if (2147483648 != 42) {
                                    if true {
                                        v77 := v1[982]
                                        fmt.Println(fi73, fi73, v58)
                                        v77 = v59
                                        _ = v77
                                    } else {
                                        if (!((!true) || true)) {
                                            fmt.Println(v58)
                                            if false {
                                                fmt.Println(v72, v74, v74)
                                            } else {
                                                fmt.Println(v59, v74)
                                                v2 = append(v2, 128)
                                            }
                                        }
                                    }
                                } else {
                                    switch ("a" + v59) {
                                    case "":
                                        if (504 < 300) {
                                            fmt.Println(v4, v72, fi73)
                                            v58 = (((3.14159 * v58) - (float64(1e10) + v58)) / (((3.14159 / 3.14159) - v58) - v58))
                                            v72 -= v72
                                            v72 += 42
                                        } else {
                                            v78 := GenS3{F0: 18446744073709551615}
                                            switch v59 {
                                            case "":
                                                var v79 float32 = 1.5
                                                _ = v79
                                            }
                                            v58 = 3.14159
                                            _ = v78
                                        }
                                    case "hello":
                                        delete(v1, 42)
                                    default:
                                        v80 := v1[1]
                                        _ = v80
                                    }
                                    v72 -= v72
                                    v58 = v58
                                    var v81 int16 = 1
                                    var v82 uint64 = 113
                                    _ = v81
                                    _ = v82
                                }
                            }
                            _ = fi73
                            _ = v74
                        }
                        v59 = v59
                        var v83 uint = 778
                        _ = v83
                    }
                }
                _ = v72
            }
            v58 = v58
        } else {
            if len(v2) > 0 {
                v2[0] = fi64
            }
            v59 = (v59 + v59)
            fmt.Println(fi64)
        }
        delete(v1, 2147483648)
        _ = fi64
    }
    v84 := GenS4{F0: (v4 / (((v4 % ((0) | 1))) | 1)), F1: 0, F2: 1}
    v3 = append(v3, -2147483648)
    var v85 uint16 = (^(v4))
    var v86 int = 9223372036854775807
    v85 += (v4 % (((v4 | 1)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v58
    _ = v59
    _ = v84
    _ = v85
    _ = v86
}
