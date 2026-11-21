package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]string)
    var v3 bool = true
    var v4 float64 = -2.25
    var v5 uint16 = 255
    if ((v3 || v3) && (-128 >= 127)) {
        v5 = uint16(v1)
        v6 := GenS1{F0: (^(v1)), F1: 0, F2: 42}
        _ = v6
    }
    var v7 int = 32767
    for fi8 := 0; fi8 < 1; fi8++ {
        if v3 {
            v3 = true
            var v9 uint32 = 2
            fmt.Println(v4)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS2{F0: (^(v1))}
                v5 -= v5
                _ = fi10
                _ = v11
            }
            switch ("hello" + "foo bar") {
            case "\n\t":
                for fi12 := 0; fi12 < 6; fi12++ {
                    if v3 {
                        var v13 string = "hello"
                        _ = v13
                    } else {
                        var v14 int = (-(fi8))
                        var v15 uint32 = 980
                        _ = v14
                        _ = v15
                    }
                    var v16 uint16 = v5
                    _ = fi12
                    _ = v16
                }
            case "a":
                fmt.Println(v5)
            case "foo bar":
                v4 -= (((0.0 / v4) * v4) - ((((-2.25 + v4) - (0.0 + 3.14159)) * ((((v4 + 1e10) - float64(-2.25)) / (1.5 - v4)) - 0.0)) * (0.1 + (v4 * v4))))
            default:
                var v17 uint32 = 0
                _ = v17
            }
            v18 := GenS3{F0: (^(v7)), F1: (v5 >> ((42) & 63)), F2: (v1 | v1)}
            _ = v18
        }
        _ = fi8
    }
    if v3 {
        v19 := []int{}
        fmt.Println(v4, v5, v1)
        _ = v19
    } else {
        v20 := []int{0}
        switch (v7 / ((v7) | 1)) {
        case 32768:
            if ((v1 << ((2) & 63)) > v1) {
                var v21 int16 = 42
                v22 := GenS2{F0: (v1 >> ((343) & 63))}
                for fi23 := 0; fi23 < 16; fi23++ {
                    fmt.Println(v21, fi23)
                    _ = fi23
                }
                _ = v21
                _ = v22
            }
        default:
            v7 -= (v7 % ((v7) | 1))
        }
        _ = v20
    }
    if ((false || (v5 != (v5 % ((v5) | 1)))) || v3) {
        for fi24 := 0; fi24 < 20; fi24++ {
            switch (v7 * (fi24 / ((fi24) | 1))) {
            case 2147483647:
                v25 := GenS3{F0: -2147483648, F1: (v5 - v5), F2: (v1 % ((173) | 1))}
                _ = v25
            default:
                for fi26 := 0; fi26 < 7; fi26++ {
                    fmt.Println(v1, v3, v1)
                    if (uint16(v1) > (v5 % (((v5 & v5)) | 1))) {
                        v2[2] = "foo bar"
                    }
                    _ = fi26
                }
            }
            switch "\n\t" {
            case "a":
                for fi27 := 0; fi27 < 3; fi27++ {
                    var v28 int = (fi24 << ((414) & 63))
                    v29 := GenS1{F0: 0, F1: -2147483648, F2: 0}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case "hello":
                var v30 uint16 = 32767
                _ = v30
            case "":
                var v31 float32 = ((1e10 - (float32(v4) + -2.25)) + 0.0)
                _ = v31
            default:
                var v32 bool = (0.1 >= 3.14159)
                _ = v32
            }
            _ = fi24
        }
    }
    v33 := GenS3{F0: (v7 - v7), F1: v5, F2: 127}
    v3 = v3
    v34 := GenS1{F0: v1, F1: 2147483647, F2: 0}
    _ = len(v2)
    var v35 int8 = 126
    fmt.Println(v1, v35)
    if v3 {
        var v36 int64 = 9223372036854775807
        v37 := v2[449]
        for fi38 := 0; fi38 < 12; fi38++ {
            v7 = -1
            _ = fi38
        }
        _ = v36
        _ = v37
    } else {
        if v3 {
            fmt.Println(v3)
            v39 := []int{(v7 << ((1) & 63))}
            for fi40 := 0; fi40 < 8; fi40++ {
                if v3 {
                    switch -1 {
                    case -979:
                        delete(v2, 513)
                    }
                    if (1 == 2147483648) {
                        var v41 float64 = float64((((((1e10 * 0.1) - 1.5) + 0.0) * 3.14159) + 1.5))
                        fmt.Println(v5, v35, v3)
                        _ = v41
                    }
                    var v42 uint16 = (v5 | v5)
                    _ = v42
                } else {
                    v43 := GenS3{F0: v7, F1: (v5 << ((952) & 63)), F2: (v1 % ((v1) | 1))}
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                    _ = v43
                }
                v4 -= 0.1
                v44 := GenS2{F0: (v1 - v1)}
                _ = fi40
                _ = v44
            }
            _ = v39
        } else {
            v35 += (v35 ^ v35)
            v3 = (uint64(v5) > 18446744073709551615)
        }
        fmt.Println(v7, v4)
    }
    v45 := GenS4{F0: v1}
    v3 = (-2.25 <= float32(v4))
    var v46 uint = 255
    switch (^(v7)) {
    case 147:
        v7 += (^(v7))
    case 176:
        fmt.Println(v1, v46)
    case 1:
        for fi47 := 0; fi47 < 3; fi47++ {
            for fi48 := 0; fi48 < 18; fi48++ {
                v49 := GenS1{F0: v1, F1: -579, F2: 2147483647}
                switch v7 {
                case 128:
                    switch ((("hello" + "\n\t") + "") + ("\n\t" + ((("a" + "") + ("hello" + "hello")) + "hello"))) {
                    case "a":
                        if v3 {
                            delete(v2, 492)
                            v4 -= (v4 / 0.1)
                        }
                    case "foo bar":
                        var v50 int32 = 42
                        _ = v50
                    case "":
                        _ = len(v2)
                    default:
                        v1 = (^(v1))
                    }
                case 1:
                    if (3.14159 == 0.1) {
                        delete(v2, 4294967295)
                        var v51 uint32 = 65535
                        _ = v51
                    } else {
                        if v3 {
                            var v52 int = fi47
                            _ = v52
                        }
                        v3 = v3
                        fmt.Println(v46, fi47)
                        var v53 uint64 = 18446744073709551615
                        _ = v53
                    }
                case -1:
                    if (-939 >= -490) {
                        if (v3 || ((0.0 != 1e10) && (false || v3))) {
                            fmt.Println(v46, v5, v3)
                        } else {
                            if v3 {
                                switch ("hello" + "a") {
                                case "":
                                    if (v3 && v3) {
                                        v54 := []int{}
                                        switch "hello" {
                                        case "":
                                            switch "" {
                                            case "":
                                                v5 += v5
                                            }
                                        case "\n\t":
                                            v54 = append(v54, 1)
                                        case "hello":
                                            v3 = (float32(float64((((-2.25 / (-2.25 - float32(0.1))) / (float32(float64(0.1)) - float32(v4))) / 1.0))) < (-2.25 + 0.1))
                                        }
                                        switch fi47 {
                                        case -1:
                                            var v55 int16 = 1
                                            _ = v55
                                        case 771:
                                            if (1e10 >= (((float64(1e10) - v4) / v4) / v4)) {
                                                if (42 != 42) {
                                                    v56 := GenS4{F0: v1}
                                                    v57 := GenS3{F0: int(v35), F1: v5, F2: 0}
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        case 681:
                                            if true {
                                                var v58 uint8 = v1
                                                fmt.Println(v58, v4)
                                                _ = v58
                                            } else {
                                                var v59 float32 = (-2.25 * 0.1)
                                                v60 := GenS5{F0: int64(v46), F1: (^(v5))}
                                                if v3 {
                                                    _ = len(v54)
                                                    v61 := GenS5{F0: 559, F1: 992}
                                                    _ = v61
                                                } else {
                                                    _ = len(v54)
                                                    v62 := GenS6{F0: 2147483648, F1: 42}
                                                    if v3 {
                                                        if (v3 || v3) {
                                                            v63 := GenS5{F0: 582, F1: (^(v5))}
                                                            v54 = append(v54, (^(v7)))
                                                            _ = v63
                                                        }
                                                    }
                                                    if true {
                                                        fmt.Println(v1)
                                                        v46 = (v46 << ((0) & 63))
                                                        v59 = 0.1
                                                    }
                                                    v7 = fi47
                                                    v54 = append(v54, (-(fi48)))
                                                    _ = v62
                                                }
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                        _ = v54
                                    }
                                case "a":
                                    v64 := GenS7{F0: -587}
                                    _ = v64
                                case "hello":
                                    v65 := GenS7{F0: 1}
                                    _ = v65
                                }
                                var v66 int32 = 457
                                _ = v66
                            } else {
                                fmt.Println(v1, v5, v35)
                                var v67 int32 = 2147483647
                                v68 := GenS4{F0: (v1 >> (((^(v46))) & 63))}
                                _ = v67
                                _ = v68
                            }
                        }
                        fmt.Println(fi48, v4)
                        v69 := GenS7{F0: fi47}
                        _ = v69
                    }
                    var v70 int = (v7 - v7)
                    switch ("a" + "hello") {
                    case "\n\t":
                        if (42 != -32768) {
                            v70 = (v70 - v7)
                        }
                    case "hello":
                        v71 := GenS3{F0: (-(v7)), F1: v5, F2: v1}
                        _ = v71
                    case "foo bar":
                        var v72 int = 1
                        _ = v72
                    }
                    _ = v70
                }
                v46 = 1
                _ = fi48
                _ = v49
            }
            if false {
                v73 := v2[4294967295]
                _ = v73
            }
            _ = fi47
        }
    default:
        var v74 uint64 = 2
        _ = v74
    }
    v3 = v3
    delete(v2, 1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v7
}
