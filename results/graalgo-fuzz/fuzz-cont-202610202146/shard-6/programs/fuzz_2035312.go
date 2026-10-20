package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int64; F2 int8 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }
type GenS5 struct { F0 uint64; F1 bool; F2 int16 }

func main() {
    var v1 string = ("foo bar" + "")
    v2 := []uint32{2147483647}
    v3 := []string{"foo bar"}
    v4 := make(map[int]uint32)
    var v5 float32 = float32(float64(1.5))
    switch 4294967295 {
    case 128:
        _ = len(v4)
    case -9223372036854775808:
        if false {
            var v6 uint8 = 192
            var v7 uint32 = 2147483647
            _ = v6
            _ = v7
        } else {
            var v8 float64 = 0.1
            v5 -= ((0.1 - float32(v8)) - 3.14159)
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 float32 = -2.25
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    case -774:
        v11 := GenS1{F0: 100}
        _ = v11
    default:
        switch 32768 {
        case 0:
            v5 = v5
        }
    }
    var v12 int = 4294967295
    fmt.Println(v1, v1, v1)
    v1 = v1
    if (-1 < int64(v12)) {
        for fi13 := 0; fi13 < 1; fi13++ {
            fmt.Println(v1, v1, v1)
            _ = len(v3)
            v5 = v5
            _ = fi13
        }
        if false {
            for fi14 := 0; fi14 < 10; fi14++ {
                for fi15 := 0; fi15 < 16; fi15++ {
                    v12 = (v12 * (v12 / ((v12) | 1)))
                    var v16 int32 = -2147483648
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    _ = fi15
                    _ = v16
                }
                fmt.Println(fi14, fi14)
                if ((341 < fi14) || (!true)) {
                    v17 := GenS1{F0: 1}
                    v18 := GenS1{F0: 127}
                    _ = v17
                    _ = v18
                } else {
                    v12 = 32768
                }
                _ = fi14
            }
        }
        fmt.Println(v12, v12, v1)
    }
    if true {
        var v19 uint = 18446744073709551615
        if true {
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := GenS1{F0: 42}
                _ = fi20
                _ = v21
            }
        } else {
            var v22 bool = (false && (false && true))
            if true {
                v23 := GenS2{F0: 2}
                _ = v23
            }
            _ = v22
        }
        _ = v19
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS1{F0: -1}
            var v26 int64 = 1
            _ = fi24
            _ = v25
            _ = v26
        }
        v27 := GenS3{F0: 127, F1: -560, F2: -79}
        switch 1 {
        case 32768:
            v3 = append(v3, v1)
        case -239:
            v28 := GenS1{F0: -1}
            _ = v28
        case 32767:
            var v29 int = -1
            _ = v29
        default:
            for fi30 := 0; fi30 < 3; fi30++ {
                v31 := GenS1{F0: 100}
                v5 = float32((float64(v5) + (1e10 - (float64((1.5 / 1.0)) * ((0.1 * (((1e10 - 1e10) + 1.5) - (3.14159 * (1.5 * 0.0)))) / 1.0)))))
                _ = fi30
                _ = v31
            }
        }
        _ = v27
    }
    v12 = v12
    v1 = v1
    for fi32 := 0; fi32 < 16; fi32++ {
        delete(v4, 0)
        _ = fi32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 string = v1
        if true {
            v35 := GenS1{F0: -1}
            _ = v35
        }
        fmt.Println(v34, v5)
        _ = fi33
        _ = v34
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 int64 = 42
        if true {
            var v38 uint32 = 376
            _ = v38
        }
        _ = fi36
        _ = v37
    }
    v5 = ((v5 - v5) - v5)
    var v39 uint = 42
    v39 = v39
    if true {
        v40 := GenS2{F0: 255}
        _ = v40
    } else {
        var v41 uint8 = 198
        v42 := GenS4{F0: 2, F1: 0, F2: 18446744073709551615}
        var v43 float64 = ((-2.25 / 3.14159) + ((0.1 * ((((3.14159 * float64(v5)) - (1e10 - float64(v5))) / (0.0 * float64(v5))) / (0.1 / ((3.14159 - ((1.5 - -2.25) / (1e10 + -2.25))) - 1.5)))) + (-2.25 * (-2.25 * (-2.25 - (0.0 / 1.0))))))
        _ = v41
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        if (!(127 <= -128)) {
            switch "" {
            case "":
                for fi45 := 0; fi45 < 3; fi45++ {
                    v39 += 255
                    _ = fi45
                }
            }
        } else {
            if (1.5 != 0.1) {
                for fi46 := 0; fi46 < 11; fi46++ {
                    v39 = (v39 << ((20) & 63))
                    _ = fi46
                }
                v5 = ((v5 / 1.0) - (1e10 / 1e10))
                v39 = (v39 ^ (v39 >> (((v39 | 650)) & 63)))
            } else {
                switch v1 {
                case "foo bar":
                    for fi47 := 0; fi47 < 3; fi47++ {
                        var v48 int8 = 58
                        _ = fi47
                        _ = v48
                    }
                }
            }
            var v49 uint16 = 536
            v50 := GenS4{F0: 1, F1: 127, F2: v39}
            _ = v49
            _ = v50
        }
        _ = fi44
    }
    switch v12 {
    case 2147483647:
        for fi51 := 0; fi51 < 13; fi51++ {
            _ = len(v4)
            v52 := GenS4{F0: 246, F1: 0, F2: (v39 >> ((v39) & 63))}
            _ = fi51
            _ = v52
        }
    case -9223372036854775808:
        if false {
            switch 4294967295 {
            case 884:
                if (((v39 + (v39 >> ((v39) & 63))) != (v39 / ((18446744073709551615) | 1))) && false) {
                    for fi53 := 0; fi53 < 12; fi53++ {
                        for fi54 := 0; fi54 < 13; fi54++ {
                            var v55 int8 = -1
                            if (1 > 9223372036854775807) {
                                v55 = v55
                            } else {
                                var v56 uint32 = 4294967295
                                _ = v56
                            }
                            _ = fi54
                            _ = v55
                        }
                        v57 := v4[(fi53 % ((fi53) | 1))]
                        v39 = (v39 << (((v39 % ((v39) | 1))) & 63))
                        _ = fi53
                        _ = v57
                    }
                    switch ("a" + v1) {
                    case "\n\t":
                        var v58 int32 = 544
                        _ = v58
                    case "foo bar":
                        var v59 bool = true
                        _ = v59
                    }
                    var v60 int8 = 1
                    _ = v60
                } else {
                    _ = len(v4)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (false || false) {
                        for fi61 := 0; fi61 < 7; fi61++ {
                            v62 := GenS1{F0: -128}
                            switch v1 {
                            case "foo bar":
                                var v63 uint32 = 2
                                _ = v63
                            }
                            _ = fi61
                            _ = v62
                        }
                        v64 := GenS1{F0: -37}
                        delete(v4, v12)
                        _ = v64
                    } else {
                        var v65 int32 = -1
                        var v66 uint8 = 255
                        _ = v65
                        _ = v66
                    }
                }
            case 9223372036854775807:
                v67 := GenS2{F0: 127}
                _ = v67
            default:
                for fi68 := 0; fi68 < 19; fi68++ {
                    v1 = v1
                    var v69 float64 = (3.14159 / 1.0)
                    switch (v12 * (fi68 | (v12 + v12))) {
                    case 1:
                        v2 = append(v2, 405)
                    case 32768:
                        v2 = append(v2, 2147483648)
                    case -435:
                        if (0 >= 42) {
                            switch (fi68 % ((v12) | 1)) {
                            case -82:
                                v1 = (("a" + v1) + ("foo bar" + ("hello" + v1)))
                            case -2147483648:
                                v5 = float32(v69)
                            case -745:
                                var v70 bool = false
                                _ = v70
                            }
                            v71 := v4[fi68]
                            switch v1 {
                            case "foo bar":
                                v72 := GenS5{F0: 18446744073709551615, F1: true, F2: 1}
                                _ = v72
                            case "":
                                var v73 bool = (false && false)
                                _ = v73
                            }
                            _ = v71
                        }
                    }
                    _ = fi68
                    _ = v69
                }
            }
        } else {
            v39 -= (v39 * v39)
        }
    case 241:
        for fi74 := 0; fi74 < 3; fi74++ {
            switch v12 {
            case 32767:
                var v75 uint32 = 2
                _ = v75
            default:
                v5 = (3.14159 + float32(0.0))
            }
            _ = fi74
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
