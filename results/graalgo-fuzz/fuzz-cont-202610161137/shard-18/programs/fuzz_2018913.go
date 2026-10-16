package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int16 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0() string {
    if (!(672 < 18446744073709551615)) {
        var v5 uint64 = 1
        _ = v5
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint8 = 0
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := []int{(^(fi8)), fi8, fi6, (fi6 >> ((931) & 63))}
                var v10 uint16 = 781
                _ = fi8
                _ = v10
                _ = v9
            }
            v7 = v7
            _ = fi6
            _ = v7
        }
        var v11 uint32 = 249
        _ = v11
    }
    var v12 int16 = 1
    var v13 int64 = -9223372036854775808
    var v14 int = -752
    for fi15 := 0; fi15 < 6; fi15++ {
        if false {
            v12 -= 365
            for fi16 := 0; fi16 < 16; fi16++ {
                if (float64((3.14159 + 1e10)) != (1e10 / 1.0)) {
                    var v17 uint16 = 45
                    if (0.1 != 3.14159) {
                        if false {
                            v18 := GenS1{F0: 0, F1: 42}
                            v12 -= -726
                            switch (fi16 - (^(fi15))) {
                            case 127:
                                v19 := GenS1{F0: 1, F1: -128}
                                _ = v19
                            case 2147483647:
                                var v20 int = fi16
                                _ = v20
                            default:
                                v21 := []int{}
                                _ = v21
                            }
                            _ = v18
                        } else {
                            switch (fi15 / ((fi15) | 1)) {
                            case 32767:
                                var v22 int32 = 0
                                _ = v22
                            }
                            var v23 uint = 255
                            _ = v23
                        }
                        var v24 bool = true
                        if ((362 >= 943) || (-2.25 >= 0.1)) {
                            switch "\n\t" {
                            case "\n\t":
                                var v25 uint32 = 2147483648
                                _ = v25
                            default:
                                var v26 int = -929
                                _ = v26
                            }
                            fmt.Println(v14, fi15)
                        } else {
                            v12 += 1000
                            if (!v24) {
                                v17 = (v17 * v17)
                                var v27 int64 = (v13 + v13)
                                _ = v27
                            } else {
                                fmt.Println(v17)
                                v13 = v13
                            }
                            switch "\n\t" {
                            case "hello":
                                var v28 int32 = int32(v17)
                                _ = v28
                            case "foo bar":
                                v29 := GenS1{F0: 2, F1: 100}
                                _ = v29
                            }
                        }
                        v24 = v24
                        _ = v24
                    }
                    v30 := []int{v14}
                    _ = v17
                    _ = v30
                }
                var v31 int16 = (v12 & (v12 & (-(v12))))
                switch fi15 {
                case 0:
                    fmt.Println(fi16, v31, v14)
                case 9223372036854775807:
                    fmt.Println(fi15, v12)
                case 127:
                    v13 = 803
                    fmt.Println(v14, v14, fi15)
                    v14 = (v14 - int(v12))
                default:
                    if ((!true) && (!false)) {
                        var v32 float64 = -2.25
                        if (955 >= 42) {
                            var v33 uint = 18446744073709551615
                            v33 = v33
                            switch "a" {
                            case "":
                                var v34 int64 = 0
                                _ = v34
                            }
                            _ = v33
                        }
                        _ = v32
                    } else {
                        v35 := []int{723}
                        if len(v35) > 0 {
                            v35[0] = (v14 * fi15)
                        }
                        v36 := GenS1{F0: 0, F1: -62}
                        _ = v35
                        _ = v36
                    }
                }
                var v37 int8 = -128
                v37 = v37
                if false {
                    v38 := []int{(fi16 + (fi15 / ((v14) | 1))), (fi16 + (fi16 * fi16))}
                    _ = v38
                } else {
                    if ((0.1 > 3.14159) || ((true && (621 > 576)) && true)) {
                        if ((42 > 1) || true) {
                            if ((657 >= 65535) || true) {
                                if false {
                                    fmt.Println(fi15)
                                    var v39 float64 = 0.0
                                    _ = v39
                                }
                                v13 += v13
                            }
                            v13 = 42
                            var v40 int8 = v37
                            _ = v40
                        }
                        v41 := GenS2{F0: v31, F1: 232, F2: 103}
                        v12 = (-(v31))
                        _ = v41
                    } else {
                        v14 -= (-(fi15))
                        var v42 string = "hello"
                        v14 = -848
                        _ = v42
                    }
                    v43 := []int{v14}
                    _ = v43
                }
                switch ("foo bar" + ((("a" + "hello") + (("" + ((("a" + (("\n\t" + (("" + ("a" + ("foo bar" + ("" + "hello")))) + "foo bar")) + "hello")) + "hello") + "")) + "")) + "\n\t")) {
                case "":
                    v44 := GenS3{F0: 155, F1: v13, F2: (v31 % (((v12 - v12)) | 1))}
                    _ = v44
                default:
                    switch fi16 {
                    case 32768:
                        fmt.Println(v14)
                    case 9223372036854775807:
                        var v45 bool = false
                        _ = v45
                    }
                }
                var v46 int8 = v37
                _ = fi16
                _ = v31
                _ = v37
                _ = v46
            }
            v47 := GenS4{F0: 0, F1: 1}
            _ = v47
        } else {
            if (true && (fi15 >= 73)) {
                fmt.Println(fi15)
            }
            switch "hello" {
            case "":
                fmt.Println(v14)
            default:
                var v48 int8 = 71
                _ = v48
            }
        }
        v14 += 32768
        _ = fi15
    }
    _ = v12
    _ = v13
    _ = v14
    return ""
}
func genFn1(p0 int, p1 uint8) int64 {
    if true {
        p0 = (p0 ^ p0)
        if (43 == uint(p1)) {
            p0 += 1
            p1 = 42
            v49 := []int{}
            _ = v49
        }
    }
    var v50 uint32 = 1
    v50 = v50
    if false {
        var v51 uint64 = 42
        _ = v51
    }
    if (true && false) {
        p1 -= p1
        p1 = p1
        if (false && (743 > 2147483647)) {
            p1 -= 0
            p1 -= 128
            if (9223372036854775807 != (^(p0))) {
                v50 -= v50
                p1 -= 171
                p0 = (p0 + (p0 * 32768))
            }
        }
    } else {
        v52 := []int{128}
        v53 := GenS1{F0: 0, F1: 87}
        var v54 int32 = 1
        _ = v52
        _ = v53
        _ = v54
    }
    _ = p0
    _ = p1
    _ = v50
    return 1
}

func main() {
    var v1 int = 32768
    var v2 int16 = 0
    v3 := []uint32{851, 0, 2147483648, 2147483648}
    var v4 float64 = float64(0.0)
    v4 += -2.25
    switch "\n\t" {
    case "a":
        fmt.Println(v4, v2, v4)
    case "foo bar":
        fmt.Println(v2, v4)
    case "hello":
        for fi55 := 0; fi55 < 7; fi55++ {
            v56 := GenS1{F0: 1, F1: 100}
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            _ = fi55
            _ = v56
        }
    default:
        _ = len(v3)
    }
    var v57 uint32 = 1
    fmt.Println(v4, v1)
    v1 = v1
    var v58 string = "\n\t"
    if (false && true) {
        v1 = v1
        v57 = (v57 & (v57 * (^(v57))))
    } else {
        var v59 uint16 = uint16(v57)
        v1 = 1
        _ = v59
    }
    var v60 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v60
}
