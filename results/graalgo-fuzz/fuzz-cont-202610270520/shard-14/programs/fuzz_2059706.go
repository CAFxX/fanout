package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int64; F1 uint32; F2 int }
type GenS9 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint16 = 308
    var v2 uint64 = 4294967295
    var v3 bool = true
    var v4 int16 = 1
    var v5 int = 599
    v6 := []uint32{42}
    v1 = v1
    v6 = append(v6, 992)
    v7 := GenS1{F0: -12}
    var v8 uint32 = 2147483647
    for fi9 := 0; fi9 < 20; fi9++ {
        switch (("foo bar" + "") + "foo bar") {
        case "\n\t":
            for fi10 := 0; fi10 < 3; fi10++ {
                v1 = 42
                if v3 {
                    var v11 int8 = 42
                    v12 := GenS1{F0: 2147483647}
                    v13 := GenS2{F0: (^(v11)), F1: 32767, F2: uint64(v8)}
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    v2 = v2
                }
                _ = fi10
            }
        case "a":
            v14 := GenS1{F0: 2147483647}
            _ = v14
        }
        v4 += -1
        if true {
            if ((!v3) || v3) {
                switch v5 {
                case -579:
                    v8 += (v8 | v8)
                case 128:
                    switch "" {
                    case "\n\t":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case "":
                        for fi15 := 0; fi15 < 3; fi15++ {
                            if true {
                                switch "a" {
                                case "foo bar":
                                    if v3 {
                                        fmt.Println(v3)
                                        v5 = v5
                                    }
                                    if true {
                                        var v16 int64 = -9223372036854775808
                                        v6 = append(v6, (v8 % (((v8 | (^(v8)))) | 1)))
                                        fmt.Println(v5, v2, v16)
                                        _ = v16
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            v17 := GenS3{F0: 1, F1: 42}
                                            _ = v17
                                        default:
                                            var v18 int16 = (v4 * (^(v4)))
                                            _ = v18
                                        }
                                        var v19 int = fi9
                                        fmt.Println(v8, v1, v3)
                                        v20 := GenS2{F0: 1, F1: 255, F2: v2}
                                        _ = v19
                                        _ = v20
                                    }
                                    v8 = 2147483648
                                case "hello":
                                    v1 += (v1 / (((v1 % ((v1) | 1))) | 1))
                                case "":
                                    if (!v3) {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v21 := GenS1{F0: int32(v4)}
                                        if v3 {
                                            v22 := GenS1{F0: 0}
                                            switch v5 {
                                            case 32768:
                                                var v23 int64 = 1
                                                _ = v23
                                            case -9223372036854775808:
                                                var v24 uint = 0
                                                _ = v24
                                            case 293:
                                                var v25 float32 = 1.5
                                                _ = v25
                                            default:
                                                var v26 uint8 = 1
                                                _ = v26
                                            }
                                            if true {
                                                v27 := GenS4{F0: 1, F1: (v4 - 32767)}
                                                if (1 != 710) {
                                                    _ = len(v6)
                                                    if len(v6) > 0 {
                                                        v6[0] = v8
                                                    }
                                                    if v3 {
                                                        v28 := GenS5{F0: fi9}
                                                        v3 = ((4294967295 != 42) && (((0.0 - 1e10) < (1.5 - 1e10)) && (v4 > (v4 * v4))))
                                                        if v3 {
                                                            if (int32(v1) >= 1) {
                                                                var v29 bool = v3
                                                                _ = v29
                                                            } else {
                                                                var v30 int16 = (v4 | v4)
                                                                _ = v30
                                                            }
                                                            var v31 float64 = 0.1
                                                            _ = v31
                                                        } else {
                                                            if len(v6) > 0 {
                                                                v6[0] = (v8 << ((1) & 63))
                                                            }
                                                        }
                                                        switch v5 {
                                                        case -2147483648:
                                                            v32 := GenS2{F0: 42, F1: v1, F2: (^(v2))}
                                                            _ = v32
                                                        case -48:
                                                            v6 = append(v6, (v8 << ((2) & 63)))
                                                        case 128:
                                                            v3 = (v3 && true)
                                                        }
                                                        var v33 int16 = -32768
                                                        fmt.Println(v3)
                                                        _ = v28
                                                        _ = v33
                                                    } else {
                                                        if v3 {
                                                            if (!((v4 << ((926) & 63)) == v4)) {
                                                                v34 := GenS6{F0: 128, F1: v4}
                                                                var v35 uint16 = v1
                                                                _ = v34
                                                                _ = v35
                                                            } else {
                                                                v36 := GenS7{F0: true, F1: fi15}
                                                                v37 := GenS8{F0: 42, F1: (v8 + v8), F2: (fi9 - (fi9 ^ fi15))}
                                                                v1 -= (v1 + (v1 | (^(v1))))
                                                                if len(v6) > 0 {
                                                                    v6[0] = v8
                                                                }
                                                                v38 := GenS4{F0: 0, F1: v4}
                                                                _ = v36
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        }
                                                        var v39 int = int(v8)
                                                        var v40 uint8 = 22
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                } else {
                                                    v2 = v2
                                                }
                                                _ = v27
                                            }
                                            _ = v22
                                        } else {
                                            if v3 {
                                                v8 -= (v8 * v8)
                                                fmt.Println(v5)
                                                var v41 uint8 = 2
                                                v2 -= 152
                                                if v3 {
                                                    if (!true) {
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                    v6 = append(v6, (^(v8)))
                                                    var v42 int32 = 2147483647
                                                    switch ("" + "hello") {
                                                    case "hello":
                                                        fmt.Println(fi15, v2)
                                                    }
                                                    v43 := GenS8{F0: -86, F1: (v8 % ((v8) | 1)), F2: (v5 ^ fi15)}
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    var v44 int = fi9
                                                    if v3 {
                                                        switch 1 {
                                                        case 0:
                                                            switch (fi9 - (fi15 + (fi9 | (^(fi15))))) {
                                                            case 128:
                                                                v45 := GenS3{F0: v1, F1: 9223372036854775807}
                                                                _ = v45
                                                            case 9223372036854775807:
                                                                var v46 int8 = 38
                                                                _ = v46
                                                            default:
                                                                v47 := GenS8{F0: -1, F1: (v8 ^ v8), F2: (^(fi9))}
                                                                _ = v47
                                                            }
                                                        default:
                                                            if len(v6) > 0 {
                                                                v6[0] = v8
                                                            }
                                                        }
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    } else {
                                                        var v48 float64 = 3.14159
                                                        v41 -= (v41 / ((uint8(v1)) | 1))
                                                        fmt.Println(fi15, v5, v44)
                                                        if v3 {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            var v49 uint16 = (^(v1))
                                                            _ = v49
                                                        }
                                                        var v50 uint64 = (v2 ^ v2)
                                                        _ = v48
                                                        _ = v50
                                                    }
                                                    _ = v44
                                                }
                                                _ = v41
                                            } else {
                                                if v3 {
                                                    var v51 uint32 = (^(v8))
                                                    v52 := GenS2{F0: 77, F1: (v1 >> ((0) & 63)), F2: 18446744073709551615}
                                                    var v53 float64 = 1.5
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    var v54 int8 = -1
                                                    if (v3 || v3) {
                                                        var v55 int = (^(v5))
                                                        _ = v55
                                                    }
                                                    if v3 {
                                                        var v56 bool = (v3 && (v1 <= (v1 - (^(v1)))))
                                                        fmt.Println(v3)
                                                        _ = v56
                                                    } else {
                                                        var v57 string = ("" + "a")
                                                        var v58 uint8 = 0
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    _ = v54
                                                }
                                                var v59 string = ("foo bar" + "\n\t")
                                                switch v59 {
                                                case "hello":
                                                    if false {
                                                        v2 += 179
                                                        switch 4294967295 {
                                                        case -942:
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            fmt.Println(v8, fi9)
                                                        case 127:
                                                            switch 988 {
                                                            case 128:
                                                                var v60 uint64 = 1
                                                                _ = v60
                                                            case 1:
                                                                var v61 int = fi15
                                                                _ = v61
                                                            case 27:
                                                                v62 := GenS9{F0: v2, F1: uint32(v2)}
                                                                _ = v62
                                                            default:
                                                                if (1 != v2) {
                                                                    v59 = ("\n\t" + "\n\t")
                                                                    var v63 int32 = int32(v4)
                                                                    _ = v63
                                                                }
                                                            }
                                                            fmt.Println(v5, fi15)
                                                        }
                                                    } else {
                                                        var v64 uint16 = v1
                                                        _ = v64
                                                    }
                                                case "":
                                                    v65 := GenS3{F0: (v1 ^ v1), F1: -1}
                                                    v1 = 292
                                                    fmt.Println(v4, fi9, fi9)
                                                    _ = v65
                                                case "foo bar":
                                                    switch (v59 + (v59 + ("foo bar" + (("" + ((v59 + v59) + v59)) + v59)))) {
                                                    case "hello":
                                                        _ = len(v6)
                                                    default:
                                                        fmt.Println(v3)
                                                    }
                                                default:
                                                    if len(v6) > 0 {
                                                        v6[0] = uint32(v1)
                                                    }
                                                }
                                                _ = v59
                                            }
                                            var v66 int8 = -96
                                            var v67 float64 = (-2.25 - 0.0)
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = v21
                                    }
                                default:
                                    if v3 {
                                        var v68 int = (fi9 % ((int(v1)) | 1))
                                        _ = v68
                                    } else {
                                        if len(v6) > 0 {
                                            v6[0] = 0
                                        }
                                    }
                                }
                                var v69 string = "hello"
                                var v70 uint64 = uint64(v8)
                                _ = v69
                                _ = v70
                            }
                            var v71 uint = 958
                            _ = fi15
                            _ = v71
                        }
                    default:
                        v8 -= v8
                    }
                default:
                    var v72 int8 = 2
                    _ = v72
                }
            } else {
                fmt.Println(v5, v3, v8)
            }
            for fi73 := 0; fi73 < 9; fi73++ {
                v74 := GenS7{F0: (2 <= (^(v8))), F1: fi73}
                _ = fi73
                _ = v74
            }
            var v75 float32 = (-2.25 / 0.1)
            _ = v75
        }
        _ = fi9
    }
    v5 += v5
    fmt.Println(v3, v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v76 := GenS8{F0: -9223372036854775808, F1: 4294967295, F2: (v5 % (((v5 & 127)) | 1))}
    for fi77 := 0; fi77 < 3; fi77++ {
        v5 += fi77
        for fi78 := 0; fi78 < 5; fi78++ {
            switch "hello" {
            case "hello":
                var v79 bool = false
                _ = v79
            case "\n\t":
                var v80 string = "foo bar"
                _ = v80
            case "a":
                if (float32(0.0) != float32(0.0)) {
                    var v81 int8 = 0
                    var v82 string = "hello"
                    _ = v81
                    _ = v82
                } else {
                    var v83 uint = 18446744073709551615
                    _ = v83
                }
            }
            _ = fi78
        }
        _ = fi77
    }
    var v84 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v8
    _ = v84
}
