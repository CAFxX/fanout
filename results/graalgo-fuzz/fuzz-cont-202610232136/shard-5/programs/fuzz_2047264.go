package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }
type GenS3 struct { F0 int; F1 int32; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 int }

func main() {
    var v1 int64 = 73
    v2 := []uint32{2147483648, 42, 0}
    var v3 int = 9223372036854775807
    v4 := make(map[string]string)
    var v5 string = ""
    if true {
        v2 = append(v2, 2)
        var v6 int = v3
        _ = v6
    } else {
        switch (v5 + v5) {
        case "":
            switch v3 {
            case 127:
                v7 := v4[(v5 + v5)]
                _ = v7
            case -9223372036854775808:
                v1 += v1
            case -1:
                switch "foo bar" {
                case "hello":
                    var v8 uint64 = 18446744073709551615
                    _ = v8
                case "foo bar":
                    v5 = (("hello" + v5) + v5)
                case "":
                    var v9 int64 = (v1 ^ v1)
                    _ = v9
                default:
                    v5 = "\n\t"
                }
            default:
                var v10 int8 = 42
                _ = v10
            }
        }
        delete(v4, v5)
        var v11 uint = 18446744073709551615
        _ = v11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v1 += (v1 >> ((18446744073709551615) & 63))
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float32 = (1.5 + 3.14159)
        fmt.Println(v13, v13, v5)
        _ = fi12
        _ = v13
    }
    var v14 uint32 = 42
    delete(v4, ((v5 + "hello") + ("hello" + "\n\t")))
    v15 := v4["hello"]
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 bool = (995 >= 2)
        var v18 int16 = 1000
        if false {
            var v19 uint32 = v14
            _ = v19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v20 float64 = (0.0 + (0.1 + (1e10 + (((float64(1.5) + -2.25) / 0.1) + (1e10 / (float64(1.5) * (-2.25 * (0.0 + (3.14159 - (1.5 * (1.5 + (-2.25 - float64((0.0 + 3.14159))))))))))))))
        var v21 uint = 886
        var v22 int32 = 709
        _ = v20
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 19; fi23++ {
            fmt.Println(v14, v15, v14)
            v3 = -2147483648
            _ = fi23
        }
        if false {
            switch v3 {
            case -2147483648:
                if (2 != 690) {
                    if true {
                        for fi24 := 0; fi24 < 9; fi24++ {
                            for fi25 := 0; fi25 < 20; fi25++ {
                                if true {
                                    var v26 float32 = (3.14159 * 0.0)
                                    if (((1.5 > ((float32((((float64((-2.25 / v26)) / 1.5) + (1.5 * 3.14159)) - (1e10 + float64(v26)))) * v26) / (-2.25 * v26))) && (214 > 97)) && false) {
                                        var v27 uint = 0
                                        _ = v27
                                    } else {
                                        v2 = append(v2, v14)
                                        var v28 uint64 = 18446744073709551615
                                        v28 -= (^(v28))
                                        _ = v28
                                    }
                                    if false {
                                        v15 = "a"
                                        if false {
                                            v29 := GenS1{F0: (v14 ^ v14)}
                                            v14 += v14
                                            switch fi24 {
                                            case 9223372036854775807:
                                                v30 := GenS1{F0: v14}
                                                _ = v30
                                            case 1:
                                                if (((3.14159 < (float32(((0.0 / ((3.14159 + 0.1) * (((1.5 / (3.14159 - 0.0)) - ((1.5 / (0.1 + 1.5)) + (0.1 / ((3.14159 * 0.1) * 1.5)))) / 3.14159))) / 3.14159)) * v26)) || (((1.5 > 0.0) || (-1 >= (v3 ^ v3))) && false)) || ((^(v14)) >= (v14 + (v14 * v14)))) {
                                                    v31 := GenS2{F0: true, F1: (!true), F2: -32768}
                                                    var v32 bool = false
                                                    fmt.Println(v14)
                                                    var v33 uint = 2
                                                    v14 -= (v14 * 65535)
                                                    v1 -= (-(v1))
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                }
                                            default:
                                                var v34 float64 = (1.5 + (float64(float32(float64(v26))) / 1.5))
                                                _ = v34
                                            }
                                            _ = v29
                                        } else {
                                            if true {
                                                v15 = v15
                                                switch (v3 * (fi25 | -310)) {
                                                case 116:
                                                    var v35 float32 = (v26 + 0.0)
                                                    _ = v35
                                                case 32767:
                                                    v26 += float32(-2.25)
                                                case 1:
                                                    if false {
                                                        switch (v3 - (v3 | int(v14))) {
                                                        case 686:
                                                            var v36 int32 = 0
                                                            _ = v36
                                                        case 0:
                                                            var v37 int64 = -386
                                                            _ = v37
                                                        case 128:
                                                            v1 += (v1 % ((0) | 1))
                                                        default:
                                                            var v38 int8 = -58
                                                            _ = v38
                                                        }
                                                    }
                                                default:
                                                    v14 = v14
                                                }
                                                switch fi25 {
                                                case 1:
                                                    v14 = 417
                                                case 965:
                                                    switch ("foo bar" + (v15 + "foo bar")) {
                                                    case "hello":
                                                        v39 := GenS1{F0: 65535}
                                                        var v40 int32 = -2147483648
                                                        switch int(v40) {
                                                        case -315:
                                                            var v41 float32 = 3.14159
                                                            _ = v41
                                                        case 128:
                                                            v14 = (^(v14))
                                                        }
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                case 128:
                                                    var v42 int64 = (v1 | v1)
                                                    _ = v42
                                                default:
                                                    var v43 uint64 = 2
                                                    _ = v43
                                                }
                                                v44 := GenS3{F0: int(v1), F1: 42, F2: 19}
                                                v45 := GenS2{F0: false, F1: (0.0 != float64(((v26 + 1.5) - v26))), F2: 316}
                                                if (true || false) {
                                                    var v46 int16 = 32767
                                                    fmt.Println(v14)
                                                    _ = v46
                                                } else {
                                                    var v47 int = fi24
                                                    _ = v47
                                                }
                                                v1 = (v1 | -9223372036854775808)
                                                switch (fi24 - int(v14)) {
                                                case 9223372036854775807:
                                                    v14 += v14
                                                case -443:
                                                    switch fi25 {
                                                    case -2147483648:
                                                        v26 = float32((((((0.0 * (0.0 / 1e10)) * ((-2.25 * 0.0) / ((3.14159 / ((1e10 * (float64(((((((3.14159 - v26) - v26) + 1.5) + (-2.25 - float32(1.5))) - ((((3.14159 - float32(1.5)) + 1e10) / ((((((v26 - float32(float64((v26 / (float32(1.5) + float32((-2.25 + ((float64((v26 * 0.0)) + 1e10) * float64(v26))))))))) / -2.25) / v26) * (0.1 / v26)) - v26) - (((1.5 - v26) * 0.1) - 3.14159))) / (v26 / float32(1.5)))) + -2.25)) - -2.25)) * 0.1)) - float64(1e10)))) + (-2.25 / 1e10)) * -2.25) + (-2.25 + 1.5)))
                                                    default:
                                                        var v48 float64 = -2.25
                                                        _ = v48
                                                    }
                                                }
                                                _ = v44
                                                _ = v45
                                            }
                                            if false {
                                                v49 := GenS4{F0: v1}
                                                v1 = 42
                                                _ = v49
                                            } else {
                                                if ((^(v14)) >= v14) {
                                                    if ((1e10 - 1.5) == 1.5) {
                                                        var v50 uint8 = 128
                                                        v15 = ("a" + v15)
                                                        var v51 uint8 = v50
                                                        var v52 float64 = -2.25
                                                        if false {
                                                            var v53 int8 = 100
                                                            v54 := GenS5{F0: -517, F1: v14, F2: v14}
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                    } else {
                                                        switch ((v5 + "hello") + v5) {
                                                        case "\n\t":
                                                            fmt.Println(v26, v15, v1)
                                                            switch (fi25 % ((fi24) | 1)) {
                                                            case 250:
                                                                var v55 uint16 = 65535
                                                                _ = v55
                                                            case 32768:
                                                                v3 += (v3 | (fi24 + v3))
                                                            case 127:
                                                                if true {
                                                                    v5 = ""
                                                                    fmt.Println(fi24, v1)
                                                                }
                                                            }
                                                            v3 += (fi25 >> ((0) & 63))
                                                            v3 = fi25
                                                        }
                                                        var v56 int = (v3 % (((fi25 / ((v3) | 1))) | 1))
                                                        v15 = ((v15 + v5) + ("hello" + v15))
                                                        v57 := GenS6{F0: 313}
                                                        v58 := GenS6{F0: 65535}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    fmt.Println(v26, fi24, fi25)
                                                }
                                            }
                                            var v59 bool = true
                                            _ = v59
                                        }
                                        var v60 int16 = 1
                                        _ = v60
                                    }
                                    if true {
                                        var v61 int8 = -92
                                        switch ("a" + v5) {
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = 2147483647
                                            }
                                        case "\n\t":
                                            if false {
                                                if false {
                                                    v15 = v15
                                                } else {
                                                    v62 := GenS7{F0: 811}
                                                    var v63 uint32 = v14
                                                    _ = v62
                                                    _ = v63
                                                }
                                                v3 -= (fi25 & -72)
                                                if (float32(-2.25) >= v26) {
                                                    _ = len(v2)
                                                    switch (fi25 - (v3 << ((65535) & 63))) {
                                                    case 9223372036854775807:
                                                        v3 -= -1
                                                    }
                                                    var v64 int64 = (v1 & 411)
                                                    _ = v64
                                                }
                                            } else {
                                                fmt.Println(v5)
                                                if ((!(!(((float32(0.0) > v26) || (0 == 255)) && false))) && true) {
                                                    var v65 float64 = 0.1
                                                    _ = v65
                                                }
                                                v66 := GenS8{F0: (v1 >> ((42) & 63)), F1: (fi25 ^ 42)}
                                                _ = v66
                                            }
                                        case "hello":
                                            var v67 uint64 = 1
                                            _ = v67
                                        }
                                        if (!false) {
                                            v68 := v4["hello"]
                                            _ = v68
                                        }
                                        _ = v61
                                    } else {
                                        var v69 float32 = v26
                                        _ = v69
                                    }
                                    v15 = v15
                                    var v70 bool = true
                                    _ = v26
                                    _ = v70
                                }
                                v71 := GenS8{F0: -1, F1: fi24}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi25
                                _ = v71
                            }
                            _ = fi24
                        }
                        v5 = v5
                        v4[v5] = "hello"
                    }
                    v15 = v5
                }
            case -425:
                if false {
                    _ = len(v4)
                    if (false || (0.1 != 1.5)) {
                        v72 := GenS3{F0: v3, F1: -1, F2: 4294967295}
                        _ = v72
                    }
                }
            default:
                v73 := GenS5{F0: 282, F1: 2, F2: (v14 ^ (v14 ^ (v14 & (v14 >> ((749) & 63)))))}
                _ = v73
            }
            fmt.Println(v3, v5, v15)
        } else {
            v74 := GenS6{F0: 65535}
            var v75 bool = (false || (!false))
            var v76 uint = 18446744073709551615
            _ = v74
            _ = v75
            _ = v76
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
