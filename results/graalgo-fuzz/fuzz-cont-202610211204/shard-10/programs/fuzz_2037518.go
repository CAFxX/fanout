package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint8; F1 int32; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 int8 }
func genFn0(p0 uint32) int8 {
    v5 := []int{42}
    var v6 uint = 1
    v6 += (v6 ^ (v6 ^ 440))
    _ = p0
    _ = v5
    _ = v6
    return -42
}
func genFn1(p0 int, p1 uint, p2 uint) uint8 {
    var v7 float32 = 1.5
    switch "foo bar" {
    case "":
        p0 -= p0
    case "a":
        v8 := GenS1{F0: 100}
        _ = v8
    case "foo bar":
        if true {
            if (false || true) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    fmt.Println(p0)
                    v10 := GenS1{F0: 0}
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 17; fi11++ {
                    if false {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            v13 := []int{(^(fi12)), (fi11 + (fi12 & (p0 % ((fi12) | 1)))), (fi11 - 9223372036854775807), -744}
                            _ = fi12
                            _ = v13
                        }
                        if false {
                            v14 := GenS1{F0: 127}
                            _ = v14
                        } else {
                            var v15 string = ("a" + "")
                            switch v15 {
                            case "\n\t":
                                v16 := GenS2{F0: -128}
                                _ = v16
                            case "hello":
                                v17 := []int{fi11}
                                _ = v17
                            case "foo bar":
                                var v18 bool = (true && true)
                                _ = v18
                            }
                            _ = v15
                        }
                        v19 := []int{fi11, 4294967295, fi11}
                        _ = v19
                    }
                    if (-824 == 42) {
                        if true {
                            var v20 uint = (p1 << (((p1 + (^(p1)))) & 63))
                            _ = v20
                        } else {
                            var v21 uint32 = uint32(p1)
                            _ = v21
                        }
                    } else {
                        p2 -= (^(p1))
                    }
                    _ = fi11
                }
            } else {
                switch (p0 & (^(p0))) {
                case 412:
                    switch (p0 % ((p0) | 1)) {
                    case 32767:
                        p0 -= (p0 + p0)
                    case 620:
                        p1 += p1
                    case 1:
                        var v22 uint64 = 822
                        _ = v22
                    }
                }
            }
            var v23 float64 = (3.14159 / 0.1)
            _ = v23
        } else {
            p1 -= 18446744073709551615
            var v24 uint64 = 4294967295
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 int8 = 9
                v27 := GenS1{F0: (v26 << (((p2 % (((p2 + 2)) | 1))) & 63))}
                if true {
                    switch (-(p0)) {
                    case 127:
                        var v28 string = ""
                        _ = v28
                    case 995:
                        if false {
                            p1 += p2
                            v7 += v7
                        }
                    case 0:
                        var v29 int32 = 0
                        _ = v29
                    }
                    fmt.Println(v7, p0, p2)
                    v7 = ((v7 * 1e10) / v7)
                } else {
                    if false {
                        if true {
                            v30 := GenS3{F0: -128}
                            fmt.Println(v26, p1)
                            v31 := GenS4{F0: p1, F1: ("" + "")}
                            _ = v30
                            _ = v31
                        }
                    } else {
                        if true {
                            v26 -= (v26 ^ (v26 / (((v26 << (((^(p2))) & 63))) | 1)))
                            v32 := []int{127, (p0 | 707)}
                            v33 := GenS1{F0: (^(v26))}
                            _ = v32
                            _ = v33
                        } else {
                            if (255 > 734) {
                                v34 := []int{(fi25 << ((65535) & 63))}
                                switch (-(fi25)) {
                                case 127:
                                    fmt.Println(fi25)
                                case 4294967295:
                                    p0 = (fi25 / (((p0 ^ fi25)) | 1))
                                case 0:
                                    p1 = p2
                                default:
                                    fmt.Println(fi25, p1, v7)
                                }
                                var v35 int16 = 1000
                                _ = v34
                                _ = v35
                            } else {
                                switch p0 {
                                case 128:
                                    v36 := GenS5{F0: 127, F1: -350, F2: -864}
                                    _ = v36
                                }
                                v37 := GenS6{F0: 42}
                                var v38 int = (fi25 << (((p2 >> (((p2 * p2)) & 63))) & 63))
                                _ = v37
                                _ = v38
                            }
                        }
                    }
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = v24
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 235
}

func main() {
    v1 := []int{1, -2147483648, 761}
    v2 := []string{((("" + ("a" + "hello")) + "a") + "foo bar"), ("" + "a"), "", ""}
    var v3 float32 = 1.5
    var v4 uint16 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v39 int8 = -128
    if (128 >= 180) {
        var v40 uint = 2
        switch -380 {
        case -671:
            v1 = append(v1, 235)
        case 128:
            v39 = (v39 & (v39 & (v39 % (((v39 ^ -1)) | 1))))
        default:
            for fi41 := 0; fi41 < 18; fi41++ {
                switch (^(fi41)) {
                case -9223372036854775808:
                    if true {
                        if true {
                            v4 = v4
                        } else {
                            v42 := GenS7{F0: (fi41 ^ (fi41 + fi41))}
                            v40 = (v40 | 774)
                            _ = v42
                        }
                    }
                }
                _ = fi41
            }
        }
        _ = v40
    } else {
        v39 += v39
        var v43 uint = 18446744073709551615
        for fi44 := 0; fi44 < 10; fi44++ {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch -513 {
                case 32768:
                    var v45 bool = false
                    _ = v45
                default:
                    if false {
                        switch "" {
                        case "foo bar":
                            switch "hello" {
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = 713
                                }
                            case "hello":
                                var v46 uint8 = 255
                                _ = v46
                            case "\n\t":
                                if (false && true) {
                                    var v47 int8 = (-(v39))
                                    var v48 int32 = -2147483648
                                    _ = v47
                                    _ = v48
                                }
                            default:
                                if true {
                                    for fi49 := 0; fi49 < 5; fi49++ {
                                        var v50 float64 = 0.0
                                        v51 := GenS2{F0: (v39 >> (((v43 % (((v43 + v43)) | 1))) & 63))}
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                    v39 = genFn0(2147483647)
                                    _ = len(v2)
                                }
                            }
                        case "hello":
                            for fi52 := 0; fi52 < 13; fi52++ {
                                if len(v2) > 0 {
                                    v2[0] = ("" + "\n\t")
                                }
                                _ = fi52
                            }
                        }
                    } else {
                        var v53 int64 = -9223372036854775808
                        if true {
                            var v54 uint = 4294967295
                            v53 = (v53 ^ 9223372036854775807)
                            switch (fi44 - fi44) {
                            case 2147483647:
                                _ = len(v1)
                            default:
                                for fi55 := 0; fi55 < 16; fi55++ {
                                    if true {
                                        fmt.Println(v43, v54, v53)
                                        var v56 int = fi55
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v56
                                    }
                                    var v57 uint32 = 175
                                    _ = fi55
                                    _ = v57
                                }
                            }
                            _ = v54
                        }
                        var v58 float32 = float32(-2.25)
                        _ = v53
                        _ = v58
                    }
                }
                if (("a" + "a") > "hello") {
                    fmt.Println(v43, v39, v3)
                    v1 = append(v1, fi44)
                    var v59 int32 = 0
                    _ = v59
                } else {
                    v60 := GenS8{F0: -1, F1: (v39 / ((0) | 1))}
                    _ = v60
                }
            }
            var v61 uint = v43
            _ = fi44
            _ = v61
        }
        _ = v43
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v62 bool = true
    v63 := GenS9{F0: (v39 % ((-11) | 1))}
    var v64 int16 = 0
    v64 -= 1000
    fmt.Println(v3, v39, v64)
    if len(v1) > 0 {
        v1[0] = 128
    }
    if true {
        var v65 uint16 = 1000
        if true {
            for fi66 := 0; fi66 < 19; fi66++ {
                if ((!(v62 && v62)) && true) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi66
            }
        } else {
            _ = len(v2)
        }
        _ = v65
    }
    v3 = (0.1 - v3)
    for fi67 := 0; fi67 < 14; fi67++ {
        v68 := GenS6{F0: (^(v4))}
        _ = fi67
        _ = v68
    }
    var v69 string = "\n\t"
    _ = len(v1)
    v64 -= int16(v4)
    v1 = append(v1, 32767)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v62
    _ = v63
    _ = v64
    _ = v69
}
