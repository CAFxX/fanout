package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 bool; F1 string; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32; F1 uint16; F2 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 int64 }

func main() {
    var v1 int = 32768
    var v2 int8 = -128
    v3 := make(map[string]int)
    v4 := []uint32{1}
    v5 := []string{}
    v1 -= 2147483647
    v3["foo bar"] = (v1 - v1)
    if (!false) {
        var v6 uint8 = 0
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v6
    } else {
        var v7 int = v1
        var v8 int = v7
        var v9 string = "\n\t"
        _ = v7
        _ = v8
        _ = v9
    }
    if (2147483647 >= 155) {
        for fi10 := 0; fi10 < 3; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: 270, F1: -194}
                v13 := GenS2{F0: 2147483647, F1: 127, F2: 2}
                var v14 uint64 = 4294967295
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi10
        }
        v2 = v2
        v15 := GenS3{F0: 118, F1: 0}
        _ = v15
    }
    v1 = v1
    var v16 uint8 = 196
    var v17 uint32 = 2
    switch "a" {
    case "\n\t":
        var v18 int16 = 1
        _ = v18
    case "a":
        switch "a" {
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = (^(v17))
            }
        case "a":
            _ = len(v4)
        }
    case "hello":
        if (!(85 != 82)) {
            v2 += v2
            fmt.Println(v17)
        } else {
            for fi19 := 0; fi19 < 12; fi19++ {
                v17 += 2147483647
                var v20 int = fi19
                v21 := v3[("foo bar" + ("a" + "a"))]
                _ = fi19
                _ = v20
                _ = v21
            }
            if (uint(v16) == 18446744073709551615) {
                v22 := GenS4{F0: (v17 > 42), F1: (("hello" + ("foo bar" + "\n\t")) + "\n\t"), F2: false}
                switch (v1 * v1) {
                case 32767:
                    switch (v1 | -488) {
                    case 1000:
                        v23 := GenS1{F0: 4294967295, F1: 1000}
                        _ = v23
                    default:
                        v3[("foo bar" + "")] = v1
                    }
                case 198:
                    v24 := GenS5{F0: 1}
                    _ = v24
                default:
                    var v25 int8 = 106
                    _ = v25
                }
                _ = v22
            }
        }
    default:
        v2 = 42
    }
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if false {
            var v26 int64 = -1
            var v27 int16 = 709
            var v28 uint32 = (v17 ^ v17)
            _ = v26
            _ = v27
            _ = v28
        }
    } else {
        v1 = (-(v1))
    }
    v29 := GenS6{F0: (^(v16)), F1: 816, F2: -9223372036854775808}
    if true {
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 string = "\n\t"
            var v32 uint16 = 5
            switch fi30 {
            case -919:
                v33 := GenS2{F0: (v17 ^ v17), F1: v16, F2: (v32 % ((32767) | 1))}
                _ = v33
            case 4294967295:
                for fi34 := 0; fi34 < 13; fi34++ {
                    if ((^(v32)) >= (^(v32))) {
                        if false {
                            if (1e10 < 3.14159) {
                                v16 = 1
                                if (!true) {
                                    var v35 int64 = 0
                                    _ = v35
                                } else {
                                    switch v31 {
                                    case "\n\t":
                                        v36 := GenS4{F0: false, F1: v31, F2: false}
                                        _ = v36
                                    default:
                                        switch v31 {
                                        case "\n\t":
                                            var v37 string = v31
                                            _ = v37
                                        case "foo bar":
                                            if true {
                                                v17 += (v17 ^ (^(v17)))
                                                v38 := GenS7{F0: true}
                                                switch 1 {
                                                case 128:
                                                    if false {
                                                        _ = len(v3)
                                                    } else {
                                                        v39 := GenS2{F0: 1, F1: 123, F2: (v32 ^ (^(v32)))}
                                                        var v40 uint64 = 2
                                                        v41 := GenS8{F0: 592}
                                                        _ = v39
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                case 127:
                                                    var v42 float64 = (((3.14159 + 1e10) / -2.25) - 0.0)
                                                    _ = v42
                                                case 0:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    if (-1 == -1) {
                                                        _ = len(v3)
                                                    } else {
                                                        var v43 float32 = float32(1e10)
                                                        _ = v43
                                                    }
                                                }
                                                if true {
                                                    var v44 uint8 = v16
                                                    _ = v44
                                                } else {
                                                    var v45 float32 = -2.25
                                                    _ = v45
                                                }
                                                _ = v38
                                            }
                                        }
                                    }
                                    fmt.Println(v2)
                                    delete(v3, v31)
                                }
                                fmt.Println(fi30, fi34, v2)
                            } else {
                                switch ((v31 + ("\n\t" + (v31 + v31))) + (v31 + v31)) {
                                case "":
                                    v46 := GenS9{F0: -1, F1: v32, F2: v2}
                                    _ = v46
                                case "hello":
                                    var v47 uint64 = 0
                                    v48 := GenS10{F0: 1}
                                    switch v31 {
                                    case "a":
                                        v49 := GenS11{F0: v17, F1: fi30}
                                        _ = v49
                                    case "hello":
                                        var v50 int64 = -522
                                        _ = v50
                                    case "foo bar":
                                        v1 -= fi30
                                        v2 -= 79
                                    default:
                                        var v51 bool = true
                                        _ = v51
                                    }
                                    _ = v47
                                    _ = v48
                                case "\n\t":
                                    _ = len(v3)
                                default:
                                    v31 = v31
                                }
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = len(v5)
                        } else {
                            v31 = (v31 + (v31 + (("" + (v31 + "foo bar")) + "\n\t")))
                            if (((v17 >> ((0) & 63)) >= v17) && (-2147483648 == -930)) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if true {
                                    switch fi30 {
                                    case 2147483647:
                                        var v52 float32 = 0.1
                                        _ = v52
                                    case 128:
                                        if len(v4) > 0 {
                                            v4[0] = (v17 * 65535)
                                        }
                                        v3[((((v31 + (v31 + ("\n\t" + (v31 + v31)))) + v31) + v31) + v31)] = 42
                                        if ((v17 | 827) != v17) {
                                            v1 += 53
                                            v53 := GenS5{F0: v16}
                                            var v54 float32 = 1e10
                                            _ = v53
                                            _ = v54
                                        } else {
                                            var v55 int32 = -1
                                            _ = v55
                                        }
                                    default:
                                        var v56 string = v31
                                        v16 = 235
                                        v17 -= v17
                                        _ = v56
                                    }
                                    if false {
                                        v16 += (v16 * (v16 << ((374) & 63)))
                                    } else {
                                        v17 = (v17 * (v17 | 4294967295))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    var v57 float32 = 3.14159
                                    var v58 int32 = 2147483647
                                    _ = v57
                                    _ = v58
                                }
                            } else {
                                v3[v31] = fi34
                                v59 := GenS4{F0: (1 == 550), F1: (v31 + (v31 + ("" + ("foo bar" + v31)))), F2: true}
                                _ = v59
                            }
                            v32 -= 948
                            switch "\n\t" {
                            case "foo bar":
                                if (42 <= 127) {
                                    var v60 int32 = -1
                                    if (4294967295 != 525) {
                                        switch "a" {
                                        case "hello":
                                            v32 = 2
                                        }
                                        switch (fi30 | (v1 << ((2) & 63))) {
                                        case 0:
                                            v2 = 42
                                            v16 += v16
                                            var v61 int64 = 42
                                            _ = v61
                                        case 32767:
                                            var v62 uint = 42
                                            _ = v62
                                        default:
                                            v2 = v2
                                            var v63 int32 = (v60 + 42)
                                            switch "hello" {
                                            case "hello":
                                                delete(v3, "hello")
                                            case "":
                                                v16 += (^(v16))
                                            }
                                            _ = v63
                                        }
                                        var v64 float32 = ((0.0 / (((float32(3.14159) * ((((-2.25 - 0.0) - 3.14159) - (((float32((3.14159 + 1.5)) - (1.5 - 0.0)) + (1e10 / 1.5)) * 1e10)) / (3.14159 + -2.25))) - (0.0 - -2.25)) / float32((1e10 - 1.5)))) + -2.25)
                                        _ = v64
                                    } else {
                                        var v65 int64 = 42
                                        _ = v65
                                    }
                                    var v66 int32 = (-(v60))
                                    fmt.Println(v17, v16)
                                    fmt.Println(v17, v66, fi30)
                                    _ = v60
                                    _ = v66
                                }
                                v67 := GenS12{F0: 677}
                                _ = v67
                            case "hello":
                                switch v31 {
                                case "":
                                    delete(v3, v31)
                                case "\n\t":
                                    var v68 uint = 7
                                    _ = v68
                                case "a":
                                    var v69 float32 = 3.14159
                                    _ = v69
                                }
                            default:
                                if ((1 != -746) || false) {
                                    var v70 uint = uint(v16)
                                    v1 = (fi30 & fi30)
                                    if false {
                                        v17 += (v17 + uint32(v32))
                                        switch ("hello" + v31) {
                                        case "hello":
                                            fmt.Println(v1, v31, v31)
                                        default:
                                            fmt.Println(v17)
                                        }
                                        v31 = "\n\t"
                                    }
                                    _ = v70
                                } else {
                                    if true {
                                        fmt.Println(v2, v16, v16)
                                    }
                                    if ((!true) || (!false)) {
                                        v16 += uint8(v17)
                                        switch v31 {
                                        case "a":
                                            v1 += (fi34 ^ (fi34 % (((fi34 & 669)) | 1)))
                                        default:
                                            var v71 int16 = -1
                                            _ = v71
                                        }
                                    } else {
                                        v2 = v2
                                    }
                                    var v72 uint8 = (v16 >> ((2) & 63))
                                    _ = v72
                                }
                            }
                            fmt.Println(v32, fi34)
                        }
                    } else {
                        fmt.Println(v31, fi30, v16)
                        var v73 string = "foo bar"
                        _ = v73
                    }
                    _ = fi34
                }
            case 128:
                var v74 uint = 65535
                _ = v74
            default:
                var v75 int = v1
                _ = v75
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        v2 -= v2
    }
    var v76 int64 = 1
    if (1 == 9223372036854775807) {
        v77 := GenS2{F0: v17, F1: (v16 - 1), F2: 65535}
        var v78 uint = 4294967295
        _ = v77
        _ = v78
    }
    var v79 int16 = 1
    v1 -= v1
    v3["\n\t"] = 1
    v2 -= v2
    var v80 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v76:", v76)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v79
    _ = v80
}
