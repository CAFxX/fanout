package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 string; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 string; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint8; F1 uint16; F2 uint8 }

func main() {
    var v1 uint8 = 128
    v2 := []uint32{}
    v3 := []string{"\n\t", "a"}
    var v4 uint8 = v1
    fmt.Println(v4, v4)
    var v5 string = ""
    fmt.Println(v4)
    fmt.Println(v1, v5, v4)
    for fi6 := 0; fi6 < 10; fi6++ {
        if (42 >= 2147483648) {
            if (!true) {
                v7 := GenS1{F0: 65535, F1: 1}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v7
            }
            v1 += v4
            for fi8 := 0; fi8 < 20; fi8++ {
                v4 = (^(v4))
                v1 = v1
                _ = len(v2)
                _ = fi8
            }
        }
        var v9 int64 = -9223372036854775808
        _ = fi6
        _ = v9
    }
    var v10 bool = true
    switch 4294967295 {
    case 2147483647:
        if v10 {
            var v11 uint8 = (v1 << ((533) & 63))
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v4, v4, v10)
                v10 = (v10 && (119 < (^(v1))))
                if v10 {
                    if true {
                        if (!(0.0 < 1e10)) {
                            if v10 {
                                v10 = v10
                                for fi13 := 0; fi13 < 5; fi13++ {
                                    _ = len(v3)
                                    v14 := GenS2{F0: (fi13 ^ (fi13 + (^(fi13))))}
                                    var v15 uint32 = 994
                                    var v16 uint32 = (v15 % (((v15 - v15)) | 1))
                                    v15 = (v15 << ((377) & 63))
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                            }
                            var v17 float64 = ((0.1 / 3.14159) + 0.0)
                            _ = v17
                        } else {
                            v18 := GenS1{F0: 627, F1: 18446744073709551615}
                            switch "" {
                            case "a":
                                fmt.Println(v4, v4, v5)
                            case "\n\t":
                                for fi19 := 0; fi19 < 18; fi19++ {
                                    switch "foo bar" {
                                    case "a":
                                        if v10 {
                                            var v20 int = fi12
                                            v21 := GenS3{F0: (v20 << ((2) & 63)), F1: v5, F2: 2}
                                            _ = v20
                                            _ = v21
                                        } else {
                                            v22 := GenS4{F0: 18446744073709551615, F1: (v4 + 127), F2: 2147483647}
                                            var v23 uint8 = v4
                                            _ = v22
                                            _ = v23
                                        }
                                    case "":
                                        var v24 bool = ((514 > 1) || false)
                                        _ = v24
                                    default:
                                        var v25 int = (fi19 - (fi19 >> ((790) & 63)))
                                        _ = v25
                                    }
                                    var v26 int = 42
                                    v27 := GenS4{F0: 2, F1: 0, F2: -1}
                                    _ = fi19
                                    _ = v26
                                    _ = v27
                                }
                            case "":
                                switch fi12 {
                                case 32768:
                                    for fi28 := 0; fi28 < 18; fi28++ {
                                        v11 -= (v4 ^ 127)
                                        var v29 float64 = ((1.5 - 1e10) / 1.0)
                                        v11 = 127
                                        _ = fi28
                                        _ = v29
                                    }
                                case -1:
                                    for fi30 := 0; fi30 < 7; fi30++ {
                                        v11 += v11
                                        v2 = append(v2, 2)
                                        v31 := GenS1{F0: 255, F1: 18446744073709551615}
                                        switch (v5 + v5) {
                                        case "foo bar":
                                            v2 = append(v2, 2147483647)
                                        case "\n\t":
                                            if true {
                                                v32 := GenS1{F0: 186, F1: 0}
                                                fmt.Println(v10, fi12, v1)
                                                fmt.Println(v10, v11, v4)
                                                _ = v32
                                            } else {
                                                v2 = append(v2, 2)
                                                switch (fi12 * fi30) {
                                                case 9223372036854775807:
                                                    if true {
                                                        var v33 uint = 4294967295
                                                        _ = v33
                                                    } else {
                                                        if ((v5 + v5) > "hello") {
                                                            v1 -= (v4 | v1)
                                                            fmt.Println(v1)
                                                        }
                                                        if v10 {
                                                            v1 = 127
                                                        }
                                                        v5 = v5
                                                    }
                                                case 4294967295:
                                                    var v34 int64 = 1
                                                    _ = v34
                                                }
                                                switch 62 {
                                                case 1:
                                                    v1 = v1
                                                case 127:
                                                    var v35 int64 = -928
                                                    _ = v35
                                                case -288:
                                                    v1 = v11
                                                }
                                                if (963 > 255) {
                                                    v11 = (v1 * v11)
                                                    var v36 bool = true
                                                    v36 = v36
                                                    v3 = append(v3, v5)
                                                    _ = v36
                                                }
                                                v5 = v5
                                                v1 = (v1 * 0)
                                                v37 := GenS1{F0: 900, F1: 853}
                                                switch fi30 {
                                                case 1:
                                                    var v38 bool = true
                                                    _ = v38
                                                case -1:
                                                    if (1 < 2) {
                                                        var v39 uint32 = 2
                                                        _ = v39
                                                    } else {
                                                        var v40 int64 = 294
                                                        v40 = -1
                                                        var v41 string = ((v5 + v5) + v5)
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                case -82:
                                                    switch "hello" {
                                                    case "hello":
                                                        switch fi12 {
                                                        case 4294967295:
                                                            var v42 uint16 = 2
                                                            _ = v42
                                                        case 127:
                                                            switch fi30 {
                                                            case 96:
                                                                var v43 int32 = 2147483647
                                                                _ = v43
                                                            case -2147483648:
                                                                if (0 >= -108) {
                                                                    v10 = v10
                                                                    v1 = (v11 & 1)
                                                                    v4 += v4
                                                                    fmt.Println(v5, v5, v10)
                                                                    v1 = 0
                                                                }
                                                            default:
                                                                v5 = (v5 + "")
                                                            }
                                                            var v44 int64 = 0
                                                            _ = v44
                                                        }
                                                    default:
                                                        if len(v3) > 0 {
                                                            v3[0] = "a"
                                                        }
                                                    }
                                                default:
                                                    v45 := GenS5{F0: int32(v11), F1: 42}
                                                    _ = v45
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 65535
                                                }
                                                fmt.Println(v1)
                                                _ = v37
                                            }
                                        }
                                        v46 := GenS3{F0: 2147483647, F1: ("" + "foo bar"), F2: 1}
                                        _ = fi30
                                        _ = v31
                                        _ = v46
                                    }
                                }
                            }
                            if (!(0 > 1)) {
                                var v47 int64 = 42
                                _ = v47
                            } else {
                                fmt.Println(v11, v11, v1)
                                if v10 {
                                    v48 := GenS6{F0: ((v5 + (v5 + v5)) + v5), F1: (v11 % ((v4) | 1)), F2: 2147483648}
                                    fmt.Println(v5, v5)
                                    _ = v48
                                }
                            }
                            _ = v18
                        }
                    }
                    fmt.Println(v1, fi12, v5)
                    v5 = v5
                }
                _ = fi12
            }
            v49 := GenS5{F0: 1, F1: 42}
            _ = v11
            _ = v49
        } else {
            if (!(1e10 != -2.25)) {
                var v50 float64 = 1.5
                v10 = v10
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v50
            }
        }
    case -772:
        for fi51 := 0; fi51 < 15; fi51++ {
            v52 := GenS2{F0: fi51}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi51
            _ = v52
        }
    default:
        switch 9223372036854775807 {
        case 127:
            var v53 string = v5
            _ = v53
        case 1:
            v5 = (((("foo bar" + ("a" + (v5 + v5))) + v5) + "foo bar") + ("\n\t" + v5))
        case 4294967295:
            v10 = (1 == 100)
        default:
            for fi54 := 0; fi54 < 9; fi54++ {
                var v55 int32 = 0
                for fi56 := 0; fi56 < 11; fi56++ {
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if (!v10) {
                        fmt.Println(v5, v10)
                        if (!(4294967295 <= uint64(v4))) {
                            v1 += (v1 * v4)
                            var v57 int64 = -943
                            if (v57 == v57) {
                                v57 += v57
                                v1 += v1
                            }
                            _ = v57
                        } else {
                            if v10 {
                                switch -861 {
                                case 128:
                                    fmt.Println(v55)
                                    switch v5 {
                                    case "hello":
                                        if v10 {
                                            v58 := GenS6{F0: ("a" + (v5 + "a")), F1: (v1 % (((^(v4))) | 1)), F2: 194}
                                            if ((v55 ^ v55) != 0) {
                                                var v59 bool = ((2147483648 != 1) || v10)
                                                v10 = (v5 >= (("hello" + "a") + ("\n\t" + ((v5 + v5) + ((v5 + v5) + (v5 + ("\n\t" + "hello")))))))
                                                if v10 {
                                                    if (!v59) {
                                                        var v60 float64 = 0.1
                                                        _ = v60
                                                    }
                                                }
                                                var v61 int8 = 0
                                                switch fi54 {
                                                case 4294967295:
                                                    if v10 {
                                                        v1 = 0
                                                        v61 -= (v61 | (^(v61)))
                                                    }
                                                default:
                                                    var v62 int16 = -1
                                                    _ = v62
                                                }
                                                _ = v59
                                                _ = v61
                                            }
                                            v63 := GenS6{F0: v5, F1: 1, F2: 2147483647}
                                            _ = v58
                                            _ = v63
                                        } else {
                                            v5 = ""
                                        }
                                    case "a":
                                        switch v5 {
                                        case "":
                                            var v64 bool = false
                                            _ = v64
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = 42
                                            }
                                        }
                                        if (float64(3.14159) != float64(1.5)) {
                                            v55 += v55
                                            v65 := GenS3{F0: (fi56 / ((-2147483648) | 1)), F1: v5, F2: 2147483648}
                                            _ = v65
                                        } else {
                                            if v10 {
                                                var v66 int = (fi54 << ((65535) & 63))
                                                _ = v66
                                            }
                                        }
                                    case "foo bar":
                                        if v10 {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v55 = v55
                                            var v67 int = fi56
                                            _ = v67
                                        } else {
                                            v68 := GenS4{F0: 2, F1: 2, F2: 1}
                                            v69 := GenS7{F0: (^(v1)), F1: 260, F2: (v1 | 0)}
                                            var v70 uint32 = 2147483648
                                            _ = v68
                                            _ = v69
                                            _ = v70
                                        }
                                    default:
                                        v1 += 19
                                    }
                                case 9223372036854775807:
                                    if len(v3) > 0 {
                                        v3[0] = (v5 + v5)
                                    }
                                }
                                var v71 string = v5
                                var v72 int16 = 42
                                _ = v71
                                _ = v72
                            }
                        }
                    } else {
                        var v73 int16 = 377
                        var v74 float32 = float32(-2.25)
                        v4 = (v4 / (((v1 ^ v1)) | 1))
                        _ = v73
                        _ = v74
                    }
                    _ = fi56
                }
                var v75 uint32 = 2
                _ = fi54
                _ = v55
                _ = v75
            }
        }
    }
    fmt.Println(v5, v4)
    _ = len(v3)
    var v76 uint8 = 149
    v77 := GenS3{F0: 9223372036854775807, F1: v5, F2: 4294967295}
    v4 = v76
    var v78 bool = (1 != 1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
    _ = v78
}
