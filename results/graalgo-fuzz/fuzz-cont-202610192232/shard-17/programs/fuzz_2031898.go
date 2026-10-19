package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int8; F2 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 string; F1 bool; F2 uint8 }
type GenS7 struct { F0 int64 }
func genFn0() uint32 {
    var v4 uint = 2
    var v5 uint32 = 2147483647
    v6 := []int{0, 9223372036854775807, 9223372036854775807}
    v5 += v5
    _ = v4
    _ = v5
    _ = v6
    return 2147483648
}
func genFn1(p0 bool, p1 uint32) int8 {
    var v7 float32 = ((1e10 + ((0.0 / 1.0) + (-2.25 - (((3.14159 + (float32(float64((float32(1e10) / 1e10))) * 1e10)) + 1.5) - float32(1.5))))) - 3.14159)
    v8 := []int{}
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 2; fi10++ {
        p1 += uint32(v9)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v11 := GenS1{F0: 2, F1: 0, F2: 0}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{523, 539}
    v2 := make(map[int]bool)
    var v3 float32 = 0.0
    v2[0] = ((32768 <= 480) && true)
    v3 += 0.0
    v3 -= v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v12 := GenS1{F0: 493, F1: 1, F2: -1}
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: 2}
        v3 = 1.5
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint64 = 744
            if false {
                var v17 uint16 = 15
                v16 = v16
                if len(v1) > 0 {
                    v1[0] = uint32(v17)
                }
                _ = v17
            } else {
                switch (fi15 + fi15) {
                case 4294967295:
                    var v18 bool = false
                    _ = v18
                }
                fmt.Println(v16)
                switch "" {
                case "a":
                    v19 := GenS2{F0: 1}
                    _ = v19
                case "":
                    switch -9223372036854775808 {
                    case 1:
                        if (("foo bar" + "a") <= "hello") {
                            if true {
                                v20 := GenS3{F0: 9223372036854775807, F1: 11, F2: 127}
                                v16 += v16
                                var v21 uint16 = 42
                                _ = v20
                                _ = v21
                            }
                            v3 -= 0.1
                            v22 := GenS4{F0: 4294967295}
                            var v23 bool = true
                            _ = v22
                            _ = v23
                        } else {
                            fmt.Println(v16)
                        }
                    case 127:
                        if (2 != 65535) {
                            v1 = append(v1, 922)
                        } else {
                            fmt.Println(v16, v3)
                            var v24 int16 = -299
                            v16 = (v16 >> ((2) & 63))
                            switch "" {
                            case "hello":
                                var v25 uint = 255
                                _ = v25
                            default:
                                v26 := GenS5{F0: v16, F1: 42}
                                _ = v26
                            }
                            _ = v24
                        }
                    default:
                        if true {
                            if (4294967295 > 65535) {
                                v16 = v16
                                var v27 int16 = 1000
                                _ = v27
                            }
                        } else {
                            v16 -= (v16 >> ((18446744073709551615) & 63))
                        }
                    }
                }
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v28 := v2[4294967295]
    if (((1e10 / v3) + v3) < v3) {
        v3 = (((float32(0.0) / (3.14159 + 1e10)) - 1.5) - (((float32(0.0) - float32(-2.25)) - v3) - 3.14159))
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            if (-1 > -128) {
                switch fi30 {
                case 42:
                    switch "foo bar" {
                    case "\n\t":
                        v3 = v3
                    case "foo bar":
                        v28 = (0.0 > 3.14159)
                        var v31 string = ("a" + "")
                        _ = v31
                    case "a":
                        var v32 uint64 = 2
                        _ = v32
                    }
                case 281:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v33 uint64 = 1
                var v34 uint32 = 186
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v3, v34, v34)
                case "hello":
                    v28 = ((655 <= uint(v34)) && v28)
                }
                v33 = 239
                var v35 uint8 = 255
                _ = v33
                _ = v34
                _ = v35
            } else {
                if false {
                    v28 = false
                } else {
                    fmt.Println(v28, fi29)
                    v28 = v28
                    switch (fi29 + (fi30 + 256)) {
                    case 128:
                        var v36 int64 = 9223372036854775807
                        var v37 string = ("\n\t" + "foo bar")
                        v28 = v28
                        _ = v36
                        _ = v37
                    case 390:
                        v38 := GenS3{F0: 504, F1: 127, F2: genFn1((v28 && v28), 65535)}
                        _ = v38
                    case -1:
                        switch (fi30 | fi29) {
                        case -2147483648:
                            var v39 uint16 = 255
                            _ = v39
                        case 1:
                            if (true && false) {
                                fmt.Println(fi30, v3, fi30)
                                if v28 {
                                    switch "" {
                                    case "\n\t":
                                        v40 := GenS3{F0: -9223372036854775808, F1: 0, F2: -128}
                                        _ = v40
                                    default:
                                        if (!v28) {
                                            v3 -= float32(-2.25)
                                            v28 = v28
                                            var v41 float64 = 1e10
                                            _ = v41
                                        }
                                    }
                                } else {
                                    _ = len(v1)
                                }
                            } else {
                                var v42 uint32 = 2
                                _ = len(v2)
                                _ = v42
                            }
                        case -103:
                            v3 = (((((0.1 + 3.14159) + v3) * 3.14159) / 1.0) - float32((float64((3.14159 * v3)) * 3.14159)))
                        }
                    }
                }
            }
            fmt.Println(v3, v3)
            _ = fi30
        }
        v3 = v3
        _ = fi29
    }
    var v43 uint = 42
    if (32767 >= 717) {
        switch ("\n\t" + "hello") {
        case "":
            if ((3.14159 < 1.5) || (9223372036854775807 == 32767)) {
                fmt.Println(v3, v3)
                if (((v3 + float32((0.0 / float64(v3)))) / 1e10) == v3) {
                    v28 = v28
                    v3 += v3
                    _ = len(v2)
                }
                if ((!((!(!v28)) || ("foo bar" <= "a"))) || v28) {
                    v44 := GenS6{F0: "", F1: true, F2: 0}
                    var v45 int32 = -726
                    var v46 uint64 = 4294967295
                    _ = v44
                    _ = v45
                    _ = v46
                }
            } else {
                if v28 {
                    var v47 string = "hello"
                    v48 := GenS6{F0: (v47 + "foo bar"), F1: (255 >= 127), F2: 100}
                    _ = v47
                    _ = v48
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        v49 := GenS3{F0: -9223372036854775808, F1: -1, F2: 42}
        var v50 int8 = 100
        _ = v49
        _ = v50
    } else {
        var v51 uint16 = 0
        _ = v51
    }
    fmt.Println(v43, v43)
    var v52 uint32 = 2147483648
    for fi53 := 0; fi53 < 12; fi53++ {
        if (-2147483648 <= -2147483648) {
            if (uint16(v52) != 255) {
                var v54 string = "foo bar"
                _ = v54
            }
            for fi55 := 0; fi55 < 14; fi55++ {
                if true {
                    v3 -= (0.0 + (v3 / (0.0 + (((((((1e10 / v3) + (3.14159 + float32(-2.25))) - ((3.14159 + ((v3 + v3) + 0.0)) / v3)) - 0.1) - 0.0) / v3) - v3))))
                    var v56 float64 = 1e10
                    _ = v56
                } else {
                    v57 := GenS1{F0: 42, F1: 0, F2: 9223372036854775807}
                    var v58 uint32 = (v52 + v52)
                    v43 = (v43 << (((v43 | (v43 - 1))) & 63))
                    _ = v57
                    _ = v58
                }
                var v59 uint8 = 42
                var v60 int64 = int64(v43)
                switch "\n\t" {
                case "hello":
                    v2[(fi53 << (((v43 / (((v43 & (^(v43)))) | 1))) & 63))] = v28
                case "a":
                    switch ("hello" + "a") {
                    case "foo bar":
                        fmt.Println(v59)
                    case "a":
                        var v61 int64 = v60
                        _ = v61
                    }
                case "\n\t":
                    v1 = append(v1, (v52 + v52))
                default:
                    if (!(895 == 1000)) {
                        switch fi55 {
                        case 62:
                            v28 = true
                        case 0:
                            if v28 {
                                fmt.Println(v28)
                                var v62 uint8 = v59
                                v3 -= 0.0
                                _ = v62
                            }
                        default:
                            delete(v2, 32768)
                        }
                        v1 = append(v1, (v52 / ((genFn0()) | 1)))
                        var v63 string = ("" + ("a" + ("foo bar" + "hello")))
                        _ = v63
                    }
                }
                v64 := GenS5{F0: 1, F1: 713}
                v65 := GenS7{F0: v60}
                if v28 {
                    var v66 bool = v28
                    _ = v66
                }
                _ = fi55
                _ = v59
                _ = v60
                _ = v64
                _ = v65
            }
        }
        fmt.Println(fi53, v28, v43)
        _ = fi53
    }
    v67 := GenS2{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v43
    _ = v52
    _ = v67
}
