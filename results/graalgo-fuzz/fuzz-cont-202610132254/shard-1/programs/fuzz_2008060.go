package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int64; F2 int32 }
type GenS7 struct { F0 int64; F1 int8; F2 int32 }

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{}
    var v3 int64 = 0
    var v4 float64 = 1.5
    v5 := make(map[uint32]int)
    var v6 uint = 4294967295
    var v7 int32 = -1
    if true {
        for fi8 := 0; fi8 < 9; fi8++ {
            v1 = 435
            fmt.Println(v1, v3)
            _ = fi8
        }
        v4 = v4
    } else {
        v7 = v7
        v2 = append(v2, 341)
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 2147483648, F1: 0, F2: (^(v1))}
            var v11 int = 1
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = len(v5)
    for fi12 := 0; fi12 < 17; fi12++ {
        if len(v2) > 0 {
            v2[0] = 188
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v13 := GenS1{F0: 305, F1: 42, F2: v1}
        _ = fi12
        _ = v13
    }
    switch 0 {
    case -9223372036854775808:
        var v14 int16 = 0
        _ = v14
    case 9223372036854775807:
        var v15 int64 = v3
        _ = v15
    case 1:
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 float64 = ((3.14159 - 0.0) + 3.14159)
            if (!(1 != 2)) {
                if false {
                    v1 = v1
                    var v18 uint8 = 255
                    _ = v18
                } else {
                    switch (fi16 ^ (fi16 - fi16)) {
                    case -9223372036854775808:
                        if ((3.14159 - float32(-2.25)) == (float32(v4) + float32(v17))) {
                            v19 := v5[uint32(v1)]
                            fmt.Println(v1)
                            v20 := GenS1{F0: 2147483647, F1: 4294967295, F2: (v1 | (v1 >> (((v6 / ((v6) | 1))) & 63)))}
                            _ = v19
                            _ = v20
                        } else {
                            var v21 float64 = (float64(1e10) - v17)
                            if true {
                                if (v1 <= 42) {
                                    var v22 int8 = 106
                                    var v23 uint = (^(v6))
                                    v6 = v6
                                    _ = v22
                                    _ = v23
                                }
                            }
                            _ = v21
                        }
                    }
                    var v24 float32 = 1e10
                    for fi25 := 0; fi25 < 10; fi25++ {
                        if (false || (!(0 > 487))) {
                            v26 := GenS1{F0: 2147483648, F1: 886, F2: (v1 & (v1 % (((v1 * (v1 * (v1 * v1)))) | 1)))}
                            v7 += v7
                            v17 = 1e10
                            _ = v26
                        }
                        v3 = v3
                        v27 := GenS2{F0: fi16}
                        _ = fi25
                        _ = v27
                    }
                    _ = v24
                }
                for fi28 := 0; fi28 < 16; fi28++ {
                    v5[2] = (fi16 * (fi28 * int(v1)))
                    v29 := GenS3{F0: (^(v6)), F1: (v6 * (v6 >> (((v6 - v6)) & 63))), F2: int8(v3)}
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi16
            _ = v17
        }
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 uint64 = 0
        v32 := GenS2{F0: -816}
        _ = fi30
        _ = v31
        _ = v32
    }
    if true {
        fmt.Println(v1, v4, v7)
        _ = len(v2)
        switch ("a" + "") {
        case "hello":
            for fi33 := 0; fi33 < 12; fi33++ {
                if ((4294967295 >= 275) || true) {
                    if true {
                        var v34 uint = (v6 - v6)
                        v2 = append(v2, 1)
                        _ = v34
                    }
                    var v35 uint64 = 18446744073709551615
                    if true {
                        if ("" > ("a" + "a")) {
                            var v36 int8 = -69
                            if false {
                                var v37 int16 = 1
                                _ = v37
                            } else {
                                _ = len(v2)
                                v38 := GenS3{F0: 0, F1: (v6 << (((v6 - v6)) & 63)), F2: -100}
                                var v39 bool = (int(v35) >= fi33)
                                _ = v38
                                _ = v39
                            }
                            fmt.Println(v6, v6)
                            _ = v36
                        }
                        v40 := GenS4{F0: 1, F1: (^(v35)), F2: int32(fi33)}
                        var v41 int8 = 100
                        _ = v40
                        _ = v41
                    } else {
                        var v42 float64 = v4
                        switch "\n\t" {
                        case "\n\t":
                            var v43 uint64 = (v35 << (((v6 % ((174) | 1))) & 63))
                            _ = v43
                        case "foo bar":
                            v44 := GenS5{F0: 255}
                            _ = v44
                        default:
                            switch int(v1) {
                            case 32767:
                                v7 = v7
                            case -2147483648:
                                if len(v2) > 0 {
                                    v2[0] = 381
                                }
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        _ = v42
                    }
                    _ = v35
                } else {
                    switch fi33 {
                    case 697:
                        switch fi33 {
                        case -9223372036854775808:
                            for fi45 := 0; fi45 < 4; fi45++ {
                                delete(v5, 1)
                                _ = fi45
                            }
                        case -804:
                            if true {
                                v46 := GenS1{F0: 2147483648, F1: 1, F2: (v1 - v1)}
                                for fi47 := 0; fi47 < 6; fi47++ {
                                    var v48 int8 = 127
                                    fmt.Println(v6)
                                    _ = fi47
                                    _ = v48
                                }
                                var v49 uint64 = 18446744073709551615
                                _ = v46
                                _ = v49
                            }
                        case 363:
                            v4 = 1e10
                        }
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v4)
                        case "":
                            v7 -= (v7 >> (((v6 - 1)) & 63))
                        }
                    }
                    v7 = -1
                }
                _ = fi33
            }
        case "\n\t":
            for fi50 := 0; fi50 < 20; fi50++ {
                v6 = (v6 / ((1) | 1))
                fmt.Println(v4, fi50, v7)
                _ = fi50
            }
        case "":
            v2 = append(v2, 0)
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v51 := GenS2{F0: 127}
        if (true && false) {
            v4 -= float64(-2.25)
            var v52 uint64 = 0
            _ = v52
        }
        _ = v51
    }
    switch "\n\t" {
    case "hello":
        _ = len(v5)
    default:
        v7 -= v7
    }
    if false {
        var v53 int8 = -56
        switch -2147483648 {
        case 42:
            switch 32768 {
            case 9223372036854775807:
                if (int16(v3) <= 42) {
                    v54 := GenS6{F0: 1, F1: (v3 | -451), F2: (-(v7))}
                    _ = v54
                } else {
                    v7 = -659
                    if (false || (false && (true && false))) {
                        var v55 uint64 = 2
                        _ = v55
                    }
                    v4 = v4
                }
            case 128:
                v53 -= v53
            case -2147483648:
                if true {
                    v1 = 65535
                    v56 := GenS6{F0: 699, F1: (v3 % (((v3 | (v3 % (((-(v3))) | 1)))) | 1)), F2: v7}
                    _ = v56
                } else {
                    switch 9223372036854775807 {
                    case 2147483647:
                        var v57 int64 = (v3 % ((v3) | 1))
                        _ = v57
                    default:
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                    }
                    v6 = v6
                }
            }
        }
        _ = v53
    } else {
        v58 := GenS3{F0: 4294967295, F1: (v6 >> ((2) & 63)), F2: 32}
        _ = v58
    }
    var v59 int64 = v3
    v60 := GenS7{F0: v59, F1: 42, F2: (v7 << (((^(v6))) & 63))}
    var v61 uint16 = v1
    var v62 int8 = -22
    var v63 string = "a"
    switch ("a" + (v63 + v63)) {
    case "a":
        for fi64 := 0; fi64 < 18; fi64++ {
            var v65 bool = true
            if ((float64(3.14159) - 1.5) >= ((v4 * v4) + (v4 - (v4 / ((v4 * (((v4 * float64(-2.25)) - (float64((0.1 * 3.14159)) * (v4 - ((float64(((3.14159 * 3.14159) * 3.14159)) - 1e10) * ((-2.25 - (1.5 / v4)) * 0.0))))) - v4)) / v4))))) {
                for fi66 := 0; fi66 < 15; fi66++ {
                    var v67 int = fi66
                    _ = fi66
                    _ = v67
                }
            }
            _ = fi64
            _ = v65
        }
    default:
        v68 := GenS4{F0: 127, F1: 223, F2: 2147483647}
        _ = v68
    }
    fmt.Println(v61, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v7
}
