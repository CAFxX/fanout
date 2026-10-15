package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 int32; F2 int32 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 15; fi4++ {
        if true {
            var v5 int32 = 0
            v6 := []int{517, fi4}
            if len(v6) > 0 {
                v6[0] = (fi4 - 32767)
            }
            _ = v5
            _ = v6
        } else {
            if true {
                if (100 != 127) {
                    for fi7 := 0; fi7 < 15; fi7++ {
                        var v8 int8 = -19
                        v8 = v8
                        v9 := GenS1{F0: true, F1: "hello"}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    var v10 uint64 = 0
                    v11 := GenS1{F0: (false || false), F1: ("a" + "a")}
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint16 = 431
                    _ = v12
                }
                var v13 int8 = 42
                v13 += (v13 * v13)
                _ = v13
            }
            var v14 string = ("a" + "\n\t")
            if (!(0 <= -1)) {
                switch (("hello" + (("a" + v14) + v14)) + (v14 + "\n\t")) {
                case "foo bar":
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v16 := GenS1{F0: false, F1: v14}
                        var v17 int64 = -759
                        var v18 float32 = (0.1 * 3.14159)
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                default:
                    var v19 float32 = 1.5
                    _ = v19
                }
            }
            _ = v14
        }
        v20 := GenS1{F0: true, F1: ("" + ("\n\t" + (("a" + "a") + "foo bar")))}
        var v21 float64 = float64(-2.25)
        _ = fi4
        _ = v20
        _ = v21
    }
    var v22 uint = 4294967295
    if false {
        var v23 uint32 = 0
        fmt.Println(v23, v22)
        var v24 bool = ((127 != 2) && false)
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 string = ("foo bar" + "foo bar")
            var v27 float32 = (0.0 * (1.5 / 1.0))
            v26 = v26
            _ = fi25
            _ = v26
            _ = v27
        }
        switch "foo bar" {
        case "\n\t":
            if ((1.5 - float64(float32(1e10))) == 3.14159) {
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 uint8 = 1
                    _ = fi28
                    _ = v29
                }
                if (-592 == 938) {
                    v30 := GenS1{F0: true, F1: "a"}
                    v22 = 661
                    _ = v30
                } else {
                    v22 = (v22 * v22)
                    var v31 int64 = 157
                    fmt.Println(v22, v22)
                    _ = v31
                }
                if true {
                    var v32 int16 = 1
                    v33 := []int{-2147483648, 0, 852, 128}
                    v22 += (v22 % (((v22 + 806)) | 1))
                    _ = v32
                    _ = v33
                } else {
                    var v34 string = ""
                    v22 += 255
                    _ = v34
                }
            } else {
                var v35 int = 127
                v36 := []int{(v35 / ((v35) | 1))}
                _ = v35
                _ = v36
            }
        case "":
            v37 := []int{2147483647, 127, -557, -473}
            _ = v37
        case "hello":
            if (true || true) {
                v22 -= (^(v22))
                v38 := GenS1{F0: (3.14159 == -2.25), F1: "foo bar"}
                var v39 bool = true
                _ = v38
                _ = v39
            }
        }
        v22 -= 2
    }
    fmt.Println(v22)
    for fi40 := 0; fi40 < 6; fi40++ {
        v22 = v22
        var v41 int32 = 0
        var v42 int32 = (v41 % (((v41 * (v41 | (v41 >> ((v22) & 63))))) | 1))
        _ = fi40
        _ = v41
        _ = v42
    }
    _ = v22
    return 220
}
func genFn1(p0 int8, p1 uint16, p2 int) uint64 {
    v43 := []int{(p2 - (-(p2))), 127, p2}
    for fi44 := 0; fi44 < 11; fi44++ {
        for fi45 := 0; fi45 < 12; fi45++ {
            p1 = p1
            v46 := GenS2{F0: (p0 - (p0 & p0)), F1: true, F2: (p1 & (p1 >> ((1) & 63)))}
            var v47 string = (("hello" + ("" + "\n\t")) + "foo bar")
            var v48 int8 = -104
            v49 := GenS2{F0: 100, F1: true, F2: p1}
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
            _ = v49
        }
        var v50 int64 = 977
        switch fi44 {
        case 127:
            p1 += p1
        case 32767:
            v43 = append(v43, 128)
        case -433:
            if len(v43) > 0 {
                v43[0] = -9223372036854775808
            }
        default:
            switch (p2 - p2) {
            case 0:
                fmt.Println(fi44)
            case 42:
                for fi51 := 0; fi51 < 6; fi51++ {
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    var v52 uint8 = 253
                    switch "a" {
                    case "\n\t":
                        switch int(v52) {
                        case 128:
                            if false {
                                if false {
                                    p1 = p1
                                    var v53 int64 = v50
                                    if len(v43) >= 2 {
                                        v43 = v43[1:len(v43)]
                                    }
                                    var v54 int8 = (p0 - int8(v52))
                                    v43 = append(v43, fi51)
                                    var v55 int32 = -947
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                } else {
                                    switch (fi44 % ((4294967295) | 1)) {
                                    case 2147483647:
                                        fmt.Println(fi51, p2, v52)
                                        var v56 uint8 = (v52 / (((v52 * 0)) | 1))
                                        _ = v56
                                    case -898:
                                        if len(v43) > 0 {
                                            v43[0] = fi44
                                        }
                                    }
                                }
                                _ = len(v43)
                            } else {
                                var v57 int16 = -32768
                                if len(v43) >= 2 {
                                    v43 = v43[1:len(v43)]
                                }
                                _ = v57
                            }
                        default:
                            var v58 int64 = (v50 >> ((4294967295) & 63))
                            fmt.Println(p1, fi44)
                            _ = v58
                        }
                    case "hello":
                        v59 := GenS2{F0: p0, F1: (-1 < 42), F2: (p1 ^ p1)}
                        _ = v59
                    case "foo bar":
                        p1 += 2
                        v52 += v52
                    }
                    v50 = (v50 / ((0) | 1))
                    _ = fi51
                    _ = v52
                }
            case 849:
                var v60 uint32 = 4294967295
                _ = v60
            }
        }
        _ = fi44
        _ = v50
    }
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    if false {
        var v61 uint64 = 42
        _ = v61
    } else {
        p2 -= (p2 & -9223372036854775808)
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{1, 42, 2147483648}
    var v3 float64 = (0.0 + 1e10)
    v62 := GenS3{F0: 2147483648, F1: 0}
    if (1 != 32767) {
        v63 := GenS4{F0: -2147483648}
        for fi64 := 0; fi64 < 2; fi64++ {
            v3 = 0.0
            _ = fi64
        }
        _ = v63
    } else {
        switch 42 {
        case 641:
            v3 = 1.5
        case 127:
            v65 := GenS5{F0: 660, F1: 2147483647, F2: -411}
            _ = v65
        }
        var v66 int = 127
        _ = v66
    }
    v67 := GenS1{F0: (!true), F1: "\n\t"}
    switch -157 {
    case 229:
        var v68 bool = true
        _ = v68
    }
    var v69 string = "hello"
    v69 = (v69 + v69)
    if len(v2) > 0 {
        v2[0] = 0
    }
    v3 += (v3 + ((-2.25 * v3) / v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v67
    _ = v69
}
