package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint8; F1 int }
func genFn0(p0 uint16) uint16 {
    v5 := GenS1{F0: 127, F1: 81, F2: 18446744073709551615}
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6, fi6)
        v7 := GenS1{F0: 255, F1: 971, F2: 1}
        _ = fi6
        _ = v7
    }
    switch 128 {
    case 487:
        var v8 float32 = ((float32((((((0.1 + ((0.1 + (0.1 - 1e10)) / 1.0)) - (3.14159 * 0.1)) * 0.1) / (float64((1.5 / 0.1)) / (1e10 / 1.0))) / 1.5)) / 0.1) * -2.25)
        _ = v8
    case 128:
        v9 := GenS2{F0: -938, F1: 96}
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            if (!false) {
                var v12 bool = true
                _ = v12
            } else {
                p0 = 255
                fmt.Println(fi11, p0)
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            switch fi10 {
            case 677:
                v14 := GenS2{F0: 2147483647, F1: 2}
                _ = v14
            case 386:
                v15 := []int{42, 0, fi13}
                _ = v15
            case 4294967295:
                v16 := []int{fi13, (-(fi13)), (fi13 * fi13), 2147483647}
                _ = v16
            }
            var v17 uint = 828
            if true {
                if (1e10 <= 0.1) {
                    if false {
                        var v18 int16 = 1000
                        var v19 int = (fi10 << (((v17 - v17)) & 63))
                        var v20 int = -218
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint32 = 4294967295
                        var v22 uint = 18446744073709551615
                        _ = v21
                        _ = v22
                    }
                } else {
                    fmt.Println(v17, p0, p0)
                }
            }
            fmt.Println(fi13, p0, p0)
            _ = fi13
            _ = v17
        }
        _ = fi10
    }
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    _ = v5
    return (p0 & 18)
}
func genFn1() int64 {
    var v23 int32 = 874
    v24 := GenS2{F0: (-(v23)), F1: 810}
    fmt.Println(v23)
    v23 += v23
    fmt.Println(v23)
    _ = v23
    _ = v24
    return 0
}

func main() {
    var v1 int16 = -1
    var v2 uint64 = 4294967295
    var v3 int64 = 946
    var v4 int16 = v1
    var v25 int32 = 934
    var v26 int32 = v25
    fmt.Println(v1)
    var v27 int = 4294967295
    var v28 bool = (0 > 64)
    v29 := []int{v27}
    v28 = ((v3 | 0) <= genFn1())
    v30 := GenS3{F0: "\n\t"}
    var v31 uint = 998
    if v28 {
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint16 = 871
            if false {
                v34 := GenS4{F0: uint8(v31), F1: (v31 >> ((v31) & 63))}
                _ = v34
            } else {
                v33 -= 593
                var v35 uint8 = 127
                _ = v35
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                v37 := GenS4{F0: 127, F1: 255}
                _ = fi36
                _ = v37
            }
            _ = fi32
            _ = v33
        }
        for fi38 := 0; fi38 < 2; fi38++ {
            v3 = v3
            _ = fi38
        }
    }
    var v39 uint64 = v2
    if v28 {
        if false {
            fmt.Println(v3, v2, v25)
            var v40 float32 = (float32(3.14159) * (1.5 / (float32((0.0 * (3.14159 + 0.0))) + 3.14159)))
            if v28 {
                if len(v29) > 0 {
                    v29[0] = (v27 - (v27 / ((v27) | 1)))
                }
            } else {
                v27 -= 508
                var v41 int8 = 42
                if v28 {
                    switch "hello" {
                    case "hello":
                        switch "a" {
                        case "a":
                            switch "hello" {
                            case "a":
                                var v42 string = "\n\t"
                                _ = v42
                            }
                        }
                    case "a":
                        v39 -= v2
                    case "foo bar":
                        v3 += (v3 | genFn1())
                    }
                    for fi43 := 0; fi43 < 1; fi43++ {
                        for fi44 := 0; fi44 < 17; fi44++ {
                            var v45 uint32 = 447
                            v2 = (v2 << (((v31 * v31)) & 63))
                            var v46 bool = v28
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        _ = fi43
                    }
                } else {
                    switch ("" + ("\n\t" + "a")) {
                    case "hello":
                        if len(v29) > 0 {
                            v29[0] = v27
                        }
                    case "\n\t":
                        if false {
                            var v47 int = (v27 - 1)
                            _ = v47
                        } else {
                            v48 := GenS5{F0: 1, F1: (v27 >> (((v31 / ((1) | 1))) & 63))}
                            var v49 float64 = ((1.5 + -2.25) / 1e10)
                            _ = v48
                            _ = v49
                        }
                    case "":
                        for fi50 := 0; fi50 < 12; fi50++ {
                            switch (v27 / ((0) | 1)) {
                            case 0:
                                if true {
                                    for fi51 := 0; fi51 < 9; fi51++ {
                                        v41 = (^(v41))
                                        v29 = append(v29, v27)
                                        _ = fi51
                                    }
                                } else {
                                    v52 := GenS4{F0: 0, F1: 4294967295}
                                    fmt.Println(v26, fi50, v1)
                                    _ = v52
                                }
                            case -9223372036854775808:
                                var v53 int64 = int64(v41)
                                _ = v53
                            }
                            var v54 int = 2147483647
                            _ = fi50
                            _ = v54
                        }
                    default:
                        for fi55 := 0; fi55 < 19; fi55++ {
                            v1 -= -32768
                            if v28 {
                                switch "\n\t" {
                                case "hello":
                                    v28 = v28
                                }
                            } else {
                                for fi56 := 0; fi56 < 9; fi56++ {
                                    v57 := GenS1{F0: 128, F1: 584, F2: 0}
                                    var v58 int32 = (v26 + v26)
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                }
                                if ((v4 & (v4 + (^(v4)))) < -1) {
                                    var v59 int32 = v25
                                    _ = v59
                                }
                            }
                            _ = fi55
                        }
                    }
                }
                _ = v41
            }
            _ = v40
        }
        v60 := GenS4{F0: 10, F1: v31}
        if len(v29) > 0 {
            v29[0] = (-(v27))
        }
        _ = v60
    }
    var v61 int = v27
    var v62 uint16 = 628
    var v63 float32 = (1e10 * ((3.14159 + -2.25) + 1.5))
    if len(v29) > 0 {
        v29[0] = (v61 ^ int(v31))
    }
    for fi64 := 0; fi64 < 19; fi64++ {
        var v65 int = 32768
        fmt.Println(v65, v26)
        switch ("foo bar" + "foo bar") {
        case "hello":
            if false {
                for fi66 := 0; fi66 < 17; fi66++ {
                    v26 = v26
                    if (!v28) {
                        if v28 {
                            var v67 bool = ((v26 + v25) >= (v26 & (v25 * (v26 * (-(v25))))))
                            var v68 uint16 = 255
                            _ = v67
                            _ = v68
                        } else {
                            fmt.Println(v3, fi66)
                            switch "foo bar" {
                            case "foo bar":
                                var v69 uint = (v31 - 18446744073709551615)
                                _ = v69
                            case "":
                                var v70 uint16 = v62
                                _ = v70
                            case "hello":
                                if len(v29) > 0 {
                                    v29[0] = (v65 * v61)
                                }
                            default:
                                var v71 string = "hello"
                                _ = v71
                            }
                        }
                        var v72 uint32 = 1
                        v65 += v61
                        v73 := GenS2{F0: v25, F1: (v31 + v31)}
                        var v74 uint16 = (v62 ^ 2)
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                    var v75 float64 = -2.25
                    _ = fi66
                    _ = v75
                }
            }
        default:
            var v76 float64 = 3.14159
            _ = v76
        }
        _ = fi64
        _ = v65
    }
    _ = len(v29)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v61
    _ = v62
    _ = v63
}
