package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }
type GenS5 struct { F0 bool; F1 string }
func genFn0(p0 int8, p1 uint) uint16 {
    fmt.Println(p1, p0)
    fmt.Println(p0, p0)
    var v4 int8 = p0
    _ = p0
    _ = p1
    _ = v4
    return 32767
}
func genFn1(p0 bool, p1 int16) int8 {
    v5 := []int{2147483647, 128}
    for fi6 := 0; fi6 < 9; fi6++ {
        p1 = p1
        v7 := GenS1{F0: (p1 >> ((65535) & 63))}
        _ = len(v5)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -1}
    p1 -= 97
    p1 = -1
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 71
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"", "\n\t", "foo bar"}
    v3 := make(map[uint32]uint32)
    var v9 int = -995
    v10 := GenS1{F0: 0}
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint16 = 65535
            switch "a" {
            case "a":
                if (!(!(18 != 255))) {
                    v12 -= 0
                    switch fi11 {
                    case 128:
                        fmt.Println(fi11, v9)
                    case 32768:
                        _ = len(v3)
                    default:
                        if (65535 != 4294967295) {
                            switch (fi11 << ((0) & 63)) {
                            case 9223372036854775807:
                                v13 := GenS1{F0: 0}
                                _ = v13
                            case -2147483648:
                                v9 -= 4294967295
                            default:
                                for fi14 := 0; fi14 < 6; fi14++ {
                                    if true {
                                        v12 += 590
                                    } else {
                                        v3[2147483648] = 2147483647
                                    }
                                    if false {
                                        v15 := GenS1{F0: 32767}
                                        v2 = append(v2, "\n\t")
                                        v16 := GenS2{F0: 110}
                                        v12 += (v12 * (v12 & 0))
                                        if (float64(-2.25) >= float64(3.14159)) {
                                            var v17 float64 = (1.5 * float64(0.0))
                                            var v18 int64 = 1
                                            _ = v17
                                            _ = v18
                                        }
                                        _ = v15
                                        _ = v16
                                    } else {
                                        v2 = append(v2, ("a" + "hello"))
                                        var v19 int8 = 77
                                        fmt.Println(v12, v12)
                                        _ = v19
                                    }
                                    var v20 int = (fi11 & fi14)
                                    _ = fi14
                                    _ = v20
                                }
                            }
                        }
                    }
                    if true {
                        if true {
                            if false {
                                var v21 uint16 = (^(v12))
                                v12 -= v12
                                _ = v21
                            }
                            fmt.Println(v9, fi11)
                            _ = len(v2)
                        } else {
                            v22 := GenS1{F0: -32768}
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 18; fi23++ {
                            v24 := GenS1{F0: -1}
                            switch (v9 + fi23) {
                            case 1:
                                var v25 int32 = 42
                                _ = v25
                            case 2147483647:
                                if ((!false) && (826 <= 954)) {
                                    v9 += (-(fi23))
                                    fmt.Println(v9, fi11)
                                } else {
                                    var v26 string = ("foo bar" + ("\n\t" + ""))
                                    delete(v1, (fi23 * 127))
                                    v9 = v9
                                    _ = v26
                                }
                            case 32768:
                                fmt.Println(fi23)
                                v12 = (^(v12))
                                v9 = fi23
                            }
                            _ = fi23
                            _ = v24
                        }
                        fmt.Println(fi11, fi11)
                    }
                } else {
                    var v27 uint64 = 1
                    _ = v27
                }
            case "hello":
                v3[2147483647] = 543
            }
            if true {
                for fi28 := 0; fi28 < 12; fi28++ {
                    var v29 uint = 2
                    _ = len(v2)
                    fmt.Println(v12)
                    v30 := GenS1{F0: 0}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            } else {
                switch (v9 + (fi11 << ((999) & 63))) {
                case 32768:
                    var v31 bool = (!true)
                    _ = v31
                case -2147483648:
                    v32 := GenS3{F0: 42, F1: 9223372036854775807}
                    _ = v32
                case 0:
                    switch (-(v9)) {
                    case 42:
                        if (!true) {
                            if (0 < 127) {
                                var v33 uint32 = 638
                                switch v9 {
                                case 1:
                                    for fi34 := 0; fi34 < 6; fi34++ {
                                        delete(v3, 0)
                                        _ = fi34
                                    }
                                case 0:
                                    switch "" {
                                    case "":
                                        var v35 int = v9
                                        _ = v35
                                    case "a":
                                        switch (^(v9)) {
                                        case -1:
                                            delete(v3, v33)
                                        default:
                                            var v36 bool = (-127 != genFn1(false, -1))
                                            _ = v36
                                        }
                                    default:
                                        v12 = (v12 << ((2) & 63))
                                    }
                                }
                                if (0 < 0) {
                                    v12 += v12
                                }
                                _ = v33
                            } else {
                                v37 := GenS2{F0: 0}
                                var v38 int16 = 389
                                delete(v1, v9)
                                _ = v37
                                _ = v38
                            }
                            var v39 uint32 = 1
                            _ = v39
                        }
                    default:
                        v40 := GenS1{F0: -32768}
                        _ = v40
                    }
                }
            }
            _ = fi11
            _ = v12
        }
        switch (v9 % ((v9) | 1)) {
        case 4294967295:
            for fi41 := 0; fi41 < 14; fi41++ {
                v9 = (fi41 << ((65535) & 63))
                v2 = append(v2, ("hello" + ""))
                if false {
                    fmt.Println(fi41, fi41)
                } else {
                    var v42 uint = 452
                    var v43 int = 127
                    v43 = 128
                    _ = v42
                    _ = v43
                }
                _ = fi41
            }
        default:
            for fi44 := 0; fi44 < 2; fi44++ {
                switch (v9 & 4294967295) {
                case 42:
                    v45 := GenS1{F0: 356}
                    _ = v45
                case 32768:
                    fmt.Println(fi44, v9)
                case 127:
                    var v46 int = (fi44 | (v9 & (fi44 | 32767)))
                    _ = v46
                default:
                    if (!((2 >= 4294967295) || (-1 < 1))) {
                        v9 = v9
                        var v47 float32 = float32(1.5)
                        switch fi44 {
                        case 948:
                            for fi48 := 0; fi48 < 13; fi48++ {
                                var v49 int64 = 919
                                v2 = append(v2, ("\n\t" + "\n\t"))
                                _ = fi48
                                _ = v49
                            }
                        default:
                            fmt.Println(v9)
                        }
                        _ = v47
                    }
                }
                v9 += 623
                if (!true) {
                    var v50 float64 = 3.14159
                    _ = v50
                } else {
                    var v51 uint8 = 22
                    v9 -= fi44
                    v52 := GenS4{F0: 42, F1: genFn1((true && (1 > -2147483648)), -598), F2: -2147483648}
                    _ = v51
                    _ = v52
                }
                _ = fi44
            }
        }
        for fi53 := 0; fi53 < 4; fi53++ {
            v9 = v9
            var v54 uint32 = 65535
            for fi55 := 0; fi55 < 8; fi55++ {
                switch (fi55 << ((42) & 63)) {
                case 9223372036854775807:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    fmt.Println(v54, v54, v9)
                }
                if true {
                    _ = len(v1)
                    v9 = fi53
                    _ = len(v1)
                    switch (fi53 >> ((uint(v54)) & 63)) {
                    case 32768:
                        v9 = -114
                    case -9223372036854775808:
                        var v56 int8 = 0
                        _ = v56
                    default:
                        if ((v9 % ((128) | 1)) <= 400) {
                            var v57 float32 = (3.14159 - 3.14159)
                            switch "\n\t" {
                            case "hello":
                                v57 -= (1.5 / ((v57 - v57) - (v57 + v57)))
                            case "\n\t":
                                fmt.Println(v54, v9)
                            }
                            v58 := GenS4{F0: -1, F1: int8(v54), F2: -1}
                            delete(v1, (fi55 / (((fi55 & (fi53 + (fi55 - fi55)))) | 1)))
                            var v59 uint16 = 255
                            _ = v57
                            _ = v58
                            _ = v59
                        } else {
                            v54 = v54
                            var v60 uint32 = (v54 >> ((857) & 63))
                            v61 := GenS4{F0: -128, F1: -48, F2: 2147483647}
                            v9 += v9
                            _ = v60
                            _ = v61
                        }
                    }
                }
                _ = fi55
            }
            _ = fi53
            _ = v54
        }
    }
    fmt.Println(v9)
    var v62 uint = 255
    v63 := GenS5{F0: true, F1: (("a" + "foo bar") + (("\n\t" + ("a" + "\n\t")) + ""))}
    v64 := GenS4{F0: genFn1(false, -539), F1: int8(v62), F2: -2147483648}
    _ = len(v2)
    fmt.Println(v9, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v62:", v62)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v64
    _ = v9
}
