package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int32; F1 bool; F2 int8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 int8 }

func main() {
    var v1 uint16 = 65535
    var v2 uint16 = (^(v1))
    v3 := []uint32{792, 380}
    var v4 bool = (42 <= 857)
    var v5 uint = 18446744073709551615
    v1 -= v2
    fmt.Println(v2, v1, v2)
    v2 = (v2 | 2)
    if (!(v4 && (42 > 255))) {
        v4 = true
    } else {
        for fi6 := 0; fi6 < 17; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                fmt.Println(v5, v1)
                var v8 uint64 = 966
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        fmt.Println(v4, v4)
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v4 = true
        fmt.Println(v4)
        var v10 int32 = 524
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "hello":
        if v4 {
            if len(v3) > 0 {
                v3[0] = 461
            }
            v11 := GenS1{F0: (v5 * v5)}
            for fi12 := 0; fi12 < 13; fi12++ {
                v1 = 0
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 int = 1
                    if ("a" >= "hello") {
                        if (!((-1 == 1000) && true)) {
                            _ = len(v3)
                        } else {
                            fmt.Println(v4, fi12, fi12)
                            v5 -= 2
                            if (v4 && v4) {
                                var v15 string = "\n\t"
                                fmt.Println(fi12, fi12)
                                _ = v15
                            }
                        }
                        fmt.Println(v14, v4, fi12)
                        switch 127 {
                        case 127:
                            v16 := GenS1{F0: v5}
                            _ = v16
                        case -815:
                            switch "foo bar" {
                            case "foo bar":
                                if v4 {
                                    v3 = append(v3, 2147483648)
                                    if (-1 > 127) {
                                        switch int(v5) {
                                        case 4294967295:
                                            v1 -= 0
                                        case -492:
                                            if (1 > -517) {
                                                fmt.Println(v2, v2)
                                                if (((v4 && (0 >= 141)) || false) && v4) {
                                                    if ((v5 << (((^(v5))) & 63)) <= (^(v5))) {
                                                        v3 = append(v3, 65535)
                                                        if (!true) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v17 := GenS2{F0: fi12, F1: 863, F2: (v2 * v2)}
                                                            _ = v17
                                                        }
                                                        var v18 int32 = -221
                                                        var v19 int64 = 0
                                                        var v20 int = 9223372036854775807
                                                        switch v14 {
                                                        case -1:
                                                            var v21 int = 128
                                                            _ = v21
                                                        case 1:
                                                            var v22 uint32 = 2
                                                            _ = v22
                                                        case 42:
                                                            _ = len(v3)
                                                        }
                                                        v18 -= (v18 & v18)
                                                        _ = v18
                                                        _ = v19
                                                        _ = v20
                                                    } else {
                                                        var v23 int16 = 32767
                                                        fmt.Println(v14, v2)
                                                        _ = v23
                                                    }
                                                    if v4 {
                                                        switch "a" {
                                                        case "\n\t":
                                                            v3 = append(v3, 4294967295)
                                                            v24 := GenS3{F0: 305, F1: 4294967295, F2: v1}
                                                            _ = v24
                                                        default:
                                                            var v25 int8 = 1
                                                            _ = v25
                                                        }
                                                        if false {
                                                            v3 = append(v3, 458)
                                                        }
                                                        v5 += (v5 << ((638) & 63))
                                                    } else {
                                                        v14 += -2147483648
                                                        if (v5 <= v5) {
                                                            v1 -= (^(v2))
                                                        } else {
                                                            v2 -= 2
                                                            if (!(v4 && (((((("hello" + (("a" + "") + ("a" + "foo bar"))) + "") + ("a" + "a")) + "hello") + (("a" + "") + "\n\t")) != "hello"))) {
                                                                if false {
                                                                    v26 := GenS3{F0: 0, F1: uint64(v5), F2: v2}
                                                                    var v27 float64 = (0.1 + 0.1)
                                                                    var v28 int16 = 1000
                                                                    var v29 float64 = float64(float32(0.0))
                                                                    _ = v26
                                                                    _ = v27
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = 118
                                                    }
                                                }
                                                var v30 string = "foo bar"
                                                _ = v30
                                            }
                                        case 127:
                                            switch v14 {
                                            case -806:
                                                var v31 uint = v5
                                                if (false || false) {
                                                    var v32 int32 = -1
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(v1, v5, fi13)
                                                    case "hello":
                                                        fmt.Println(fi13)
                                                    case "\n\t":
                                                        if (!true) {
                                                            v2 -= v2
                                                            if (!(true && v4)) {
                                                                var v33 uint32 = 1
                                                                _ = v33
                                                            }
                                                        }
                                                    }
                                                    _ = v32
                                                } else {
                                                    fmt.Println(v1, v2, v31)
                                                }
                                                _ = v31
                                            }
                                        }
                                        if (v4 && v4) {
                                            v34 := GenS1{F0: v5}
                                            v5 += (v5 << ((v5) & 63))
                                            v1 -= 0
                                            _ = v34
                                        }
                                    }
                                } else {
                                    v35 := GenS3{F0: 1, F1: 1, F2: v2}
                                    v5 += v5
                                    _ = v35
                                }
                            }
                        case 9223372036854775807:
                            switch (-(fi13)) {
                            case -2147483648:
                                if (v4 || true) {
                                    var v36 int8 = 100
                                    v2 = (v1 >> ((v5) & 63))
                                    if v4 {
                                        var v37 int64 = -9223372036854775808
                                        v37 = (v37 ^ v37)
                                        _ = v37
                                    } else {
                                        v38 := GenS4{F0: 0, F1: (float32(float64(-2.25)) > -2.25), F2: v36}
                                        var v39 int = (fi13 << ((v5) & 63))
                                        var v40 uint64 = 620
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v36
                                } else {
                                    if false {
                                        fmt.Println(v5, v14, fi12)
                                    }
                                    v41 := GenS1{F0: (^(v5))}
                                    var v42 uint8 = 0
                                    fmt.Println(v42)
                                    _ = v41
                                    _ = v42
                                }
                            case -593:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 2147483647:
                                if true {
                                    _ = len(v3)
                                } else {
                                    var v43 uint64 = 386
                                    v43 -= (v43 / ((4294967295) | 1))
                                    if v4 {
                                        fmt.Println(v43, fi12)
                                        v4 = v4
                                        switch fi13 {
                                        case -9223372036854775808:
                                            if v4 {
                                                if len(v3) > 0 {
                                                    v3[0] = uint32(v2)
                                                }
                                                var v44 uint = (^(v5))
                                                _ = v44
                                            } else {
                                                v1 = (^(v1))
                                                var v45 int8 = -26
                                                switch (v14 * v14) {
                                                case 2147483647:
                                                    v46 := GenS5{F0: -1, F1: (fi13 / ((0) | 1))}
                                                    v5 = (^(v5))
                                                    _ = v46
                                                case 905:
                                                    v47 := GenS2{F0: (-(fi13)), F1: 0, F2: v1}
                                                    _ = v47
                                                case 32767:
                                                    v43 = 2
                                                    var v48 bool = (!(v43 != (v43 >> (((v5 >> ((899) & 63))) & 63))))
                                                    _ = v48
                                                default:
                                                    var v49 uint16 = (^(v2))
                                                    _ = v49
                                                }
                                                var v50 uint32 = 1
                                                var v51 int32 = -1
                                                _ = v45
                                                _ = v50
                                                _ = v51
                                            }
                                        }
                                    }
                                    v52 := GenS6{F0: 2147483647, F1: 2}
                                    var v53 uint8 = 0
                                    _ = v43
                                    _ = v52
                                    _ = v53
                                }
                            default:
                                if (2147483647 <= 504) {
                                    v5 += (^(v5))
                                    fmt.Println(fi13, v1, v2)
                                    if ((("a" < "foo bar") || v4) && true) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v2 = 42
                                        if (false || (!v4)) {
                                            v14 = -228
                                        }
                                    }
                                } else {
                                    switch (fi12 + v14) {
                                    case 9223372036854775807:
                                        if len(v3) > 0 {
                                            v3[0] = 42
                                        }
                                    case -1:
                                        var v54 uint8 = 1
                                        _ = v54
                                    case 1:
                                        v1 = (v1 % ((uint16(v5)) | 1))
                                    default:
                                        v1 += (v2 / ((v1) | 1))
                                    }
                                    v1 -= (v1 << (((v5 ^ v5)) & 63))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                        }
                        var v55 float32 = 3.14159
                        var v56 int8 = 42
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v55
                        _ = v56
                    } else {
                        switch (("" + (("\n\t" + "") + "")) + ("a" + "foo bar")) {
                        case "a":
                            var v57 int32 = -1
                            _ = v57
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "\n\t":
                            var v58 uint32 = 4294967295
                            _ = v58
                        default:
                            if (v4 && v4) {
                                var v59 uint16 = (^(v2))
                                v1 = v59
                                v14 = 127
                                _ = v59
                            } else {
                                if ((^(v1)) > v1) {
                                    fmt.Println(v5, v4)
                                    var v60 int8 = -1
                                    v61 := GenS5{F0: 360, F1: (fi13 + v14)}
                                    _ = v60
                                    _ = v61
                                }
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                }
                                v62 := GenS3{F0: 1, F1: 42, F2: (^(v2))}
                                _ = v62
                            }
                            var v63 float64 = -2.25
                            var v64 int64 = 0
                            var v65 uint64 = 505
                            if (!v4) {
                                switch (fi13 - fi13) {
                                case -2147483648:
                                    var v66 float32 = 0.0
                                    _ = v66
                                }
                                v1 = 65535
                            } else {
                                v14 += (-(fi12))
                                v14 = 2147483647
                            }
                            v64 -= (^(v64))
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        fmt.Println(v4, v14, v1)
                        var v67 int64 = -781
                        v1 -= (v2 % (((v1 % ((uint16(v5)) | 1))) | 1))
                        v68 := GenS5{F0: 42, F1: -9223372036854775808}
                        v69 := GenS7{F0: "", F1: 32767}
                        v70 := GenS8{F0: -1}
                        _ = v67
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    v14 += 270
                    _ = fi13
                    _ = v14
                }
                var v71 uint64 = 1
                _ = fi12
                _ = v71
            }
            _ = v11
        }
    case "":
        v4 = (v4 || true)
    case "\n\t":
        for fi72 := 0; fi72 < 13; fi72++ {
            for fi73 := 0; fi73 < 13; fi73++ {
                v5 = (v5 | 170)
                _ = fi73
            }
            _ = len(v3)
            var v74 uint = 255
            _ = fi72
            _ = v74
        }
    default:
        for fi75 := 0; fi75 < 2; fi75++ {
            v2 = (v2 + v2)
            _ = fi75
        }
    }
    _ = len(v3)
    for fi76 := 0; fi76 < 20; fi76++ {
        var v77 int8 = -128
        _ = len(v3)
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
