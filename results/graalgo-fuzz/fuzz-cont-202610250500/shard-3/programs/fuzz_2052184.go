package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS6 struct { F0 uint32; F1 int; F2 int16 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 uint16; F1 int; F2 string }
type GenS9 struct { F0 int8 }

func main() {
    var v1 float64 = float64(((1.5 + (3.14159 - 3.14159)) + 0.1))
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 uint8 = 205
    v4 = 0
    v5 := v2[-578]
    fmt.Println(v1, v1)
    var v6 bool = true
    var v7 uint8 = 255
    if (1 > 1) {
        if (!v6) {
            var v8 uint64 = 2
            _ = v8
        } else {
            var v9 float32 = 0.1
            v6 = v6
            _ = v9
        }
        v10 := GenS1{F0: -1}
        _ = v10
    }
    var v11 uint = 0
    if v6 {
        v12 := []int{32768, 127, 1}
        _ = len(v12)
        _ = v12
    } else {
        if v6 {
            fmt.Println(v6, v6, v7)
            switch "a" {
            case "hello":
                for fi13 := 0; fi13 < 16; fi13++ {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v11 += (v11 >> (((v11 + (^(v11)))) & 63))
                        if ((((^(v4)) == (v7 ^ 127)) && (false || (((v4 / (((^(v7))) | 1)) >= 1) && ((v11 - v11) == (v11 & 255))))) || v6) {
                            if (false || (v1 < v1)) {
                                switch (fi14 & 205) {
                                case -1:
                                    switch "\n\t" {
                                    case "a":
                                        var v15 uint64 = 42
                                        _ = v15
                                    case "hello":
                                        v16 := v2[(-(fi14))]
                                        _ = v16
                                    default:
                                        switch "a" {
                                        case "hello":
                                            switch fi13 {
                                            case 4294967295:
                                                v2[(-(fi14))] = v5
                                            case -9223372036854775808:
                                                var v17 float32 = float32(0.0)
                                                _ = v17
                                            case 32768:
                                                if v6 {
                                                    if (-1 == 0) {
                                                        v18 := GenS1{F0: fi14}
                                                        v19 := []int{fi13, 173}
                                                        var v20 int = (fi14 + 734)
                                                        _ = v18
                                                        _ = v19
                                                        _ = v20
                                                    } else {
                                                        fmt.Println(v11, v5)
                                                        var v21 string = (("\n\t" + ("" + (("" + "a") + (("a" + "a") + "")))) + ("foo bar" + (("foo bar" + "foo bar") + ("foo bar" + ("a" + "\n\t")))))
                                                        if v6 {
                                                            if v6 {
                                                                v22 := GenS2{F0: 100, F1: 42, F2: 128}
                                                                fmt.Println(v6)
                                                                v23 := GenS3{F0: 100, F1: (^(v5))}
                                                                _ = v22
                                                                _ = v23
                                                            }
                                                            var v24 uint32 = (v5 | 42)
                                                            if v6 {
                                                                if v6 {
                                                                    v6 = v6
                                                                    var v25 uint32 = v5
                                                                    _ = v25
                                                                } else {
                                                                    v26 := []int{}
                                                                    v4 -= 127
                                                                    _ = v26
                                                                }
                                                            }
                                                            v21 = ((v21 + (v21 + v21)) + v21)
                                                            if true {
                                                                v3[(fi14 + fi13)] = v5
                                                                switch (fi13 ^ fi14) {
                                                                case 127:
                                                                    v27 := GenS4{F0: ""}
                                                                    _ = v27
                                                                default:
                                                                    switch ("foo bar" + ((v21 + v21) + "hello")) {
                                                                    case "\n\t":
                                                                        delete(v2, fi13)
                                                                    default:
                                                                        v1 = 1e10
                                                                    }
                                                                }
                                                                var v28 int64 = 9223372036854775807
                                                                _ = v28
                                                            }
                                                            _ = v24
                                                        } else {
                                                            var v29 int32 = 277
                                                            switch (v21 + "") {
                                                            case "hello":
                                                                switch fi13 {
                                                                case -415:
                                                                    v7 += (v4 >> (((v11 - v11)) & 63))
                                                                default:
                                                                    v3[fi13] = (v5 - 42)
                                                                    v30 := GenS1{F0: (fi14 / (((-(fi13))) | 1))}
                                                                    v31 := GenS5{F0: 42, F1: 18446744073709551615, F2: -478}
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                            case "a":
                                                                var v32 int8 = 127
                                                                _ = v32
                                                            case "\n\t":
                                                                var v33 uint16 = 255
                                                                _ = v33
                                                            default:
                                                                switch fi14 {
                                                                case -857:
                                                                    fmt.Println(v1, v7)
                                                                case 128:
                                                                    var v34 float64 = (((-2.25 / v1) * 1e10) - v1)
                                                                    var v35 int64 = 0
                                                                    _ = v34
                                                                    _ = v35
                                                                case -2147483648:
                                                                    v3[fi13] = v5
                                                                default:
                                                                    var v36 bool = v6
                                                                    _ = v36
                                                                }
                                                            }
                                                            v4 -= v4
                                                            var v37 bool = true
                                                            _ = v29
                                                            _ = v37
                                                        }
                                                        fmt.Println(v7)
                                                        if false {
                                                            fmt.Println(v6, fi14)
                                                        }
                                                        _ = v21
                                                    }
                                                    if (fi14 > 9223372036854775807) {
                                                        switch ((("" + "a") + "hello") + "a") {
                                                        case "hello":
                                                            v7 -= v4
                                                        case "foo bar":
                                                            v38 := GenS3{F0: -128, F1: (^(v5))}
                                                            _ = v38
                                                        case "a":
                                                            var v39 float64 = v1
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if v6 {
                                                                    if (42 < 2) {
                                                                        fmt.Println(v7)
                                                                        var v40 int64 = 988
                                                                        _ = v40
                                                                    } else {
                                                                        v4 += v4
                                                                    }
                                                                    var v41 int = (fi14 - -2147483648)
                                                                    v41 = (^(fi13))
                                                                    _ = v41
                                                                } else {
                                                                    var v42 int8 = 106
                                                                    v11 = 334
                                                                    _ = v42
                                                                }
                                                            case "\n\t":
                                                                v43 := GenS6{F0: uint32(v7), F1: fi14, F2: 635}
                                                                _ = v43
                                                            }
                                                            var v44 int64 = 1
                                                            _ = v39
                                                            _ = v44
                                                        }
                                                    } else {
                                                        var v45 int = (fi14 / (((fi14 % (((fi14 >> ((329) & 63))) | 1))) | 1))
                                                        var v46 uint16 = 42
                                                        v46 -= v46
                                                        fmt.Println(v45, fi13, v4)
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    var v47 uint32 = (v5 % ((v5) | 1))
                                                    _ = v47
                                                }
                                            default:
                                                fmt.Println(fi13, v7)
                                            }
                                        default:
                                            if false {
                                                var v48 int16 = int16(v4)
                                                var v49 float32 = float32((v1 - v1))
                                                v50 := GenS7{F0: (^(v7)), F1: (fi14 << ((uint(v7)) & 63))}
                                                switch 127 {
                                                case 32768:
                                                    fmt.Println(fi14, v7, fi13)
                                                case -1:
                                                    if (2147483647 >= -1) {
                                                        v11 = (v11 % (((v11 - v11)) | 1))
                                                        if v6 {
                                                            fmt.Println(v4, v5, v48)
                                                            v51 := GenS8{F0: 2, F1: (fi14 | (fi14 + (fi14 >> ((v11) & 63)))), F2: "foo bar"}
                                                            _ = v51
                                                        } else {
                                                            var v52 uint16 = 2
                                                            _ = v52
                                                        }
                                                    }
                                                case 4294967295:
                                                    v1 = (v1 / 1.5)
                                                default:
                                                    delete(v2, fi13)
                                                }
                                                v53 := GenS9{F0: 0}
                                                if (("a" != "") && v6) {
                                                    fmt.Println(fi14)
                                                    v7 -= (v4 & v7)
                                                } else {
                                                    v7 += (v4 | (v4 / ((v4) | 1)))
                                                }
                                                v48 = int16(fi13)
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                                _ = v53
                                            }
                                        }
                                    }
                                default:
                                    var v54 int8 = 0
                                    _ = v54
                                }
                                v55 := []int{(fi14 - fi14), fi14, (fi13 >> (((v11 - (v11 % ((v11) | 1)))) & 63)), 1}
                                _ = v55
                            }
                            v56 := GenS6{F0: v5, F1: (fi14 - fi13), F2: -186}
                            _ = v56
                        } else {
                            var v57 int8 = -22
                            _ = v57
                        }
                        var v58 float64 = 3.14159
                        _ = fi14
                        _ = v58
                    }
                    fmt.Println(fi13, v1, v1)
                    _ = fi13
                }
            case "a":
                v59 := GenS3{F0: -128, F1: 0}
                _ = v59
            }
        }
        v60 := GenS5{F0: 34, F1: 4294967295, F2: -9223372036854775808}
        for fi61 := 0; fi61 < 7; fi61++ {
            v4 = (v4 - (^(v4)))
            v1 = (v1 / v1)
            if true {
                switch "\n\t" {
                case "a":
                    var v62 int64 = int64(v11)
                    _ = v62
                case "\n\t":
                    if (((1.5 / 1e10) * 3.14159) == (1e10 - ((0.1 * ((3.14159 + 1e10) + 1e10)) - (0.1 - -2.25)))) {
                        var v63 bool = (768 > 538)
                        v2[(-(fi61))] = v5
                        switch "" {
                        case "":
                            v3[fi61] = 2147483648
                        case "a":
                            var v64 int8 = 40
                            _ = v64
                        }
                        _ = v63
                    } else {
                        v1 += 0.1
                        fmt.Println(v7, v6)
                        for fi65 := 0; fi65 < 11; fi65++ {
                            var v66 float32 = float32(v1)
                            v66 = ((v66 - (3.14159 + (((((v66 * ((((v66 * v66) / v66) + (v66 + (((v66 / (v66 - 1e10)) * v66) * (v66 * float32(v1))))) - v66)) * (v66 * (1e10 * ((1e10 * -2.25) + 1e10)))) + 0.0) / -2.25) + float32((((-2.25 - 1e10) + v1) + (3.14159 + (v1 - v1))))))) + v66)
                            _ = fi65
                            _ = v66
                        }
                    }
                }
                var v67 int64 = -122
                v68 := GenS9{F0: 42}
                _ = v67
                _ = v68
            } else {
                fmt.Println(fi61)
            }
            _ = fi61
        }
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
