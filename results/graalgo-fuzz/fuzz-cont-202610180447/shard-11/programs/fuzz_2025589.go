package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint }
type GenS10 struct { F0 int32; F1 bool; F2 int64 }
func genFn0(p0 int64, p1 int16) uint64 {
    var v5 int = 127
    var v6 float32 = 1.5
    var v7 bool = (false && false)
    v8 := []int{32768, v5, v5, -2147483648}
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int32 = -2147483648
            for fi11 := 0; fi11 < 2; fi11++ {
                if v7 {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        var v12 int32 = -2147483648
                        fmt.Println(v6, v10)
                        if len(v8) > 0 {
                            v8[0] = v5
                        }
                        var v13 float32 = v6
                        var v14 uint16 = 1
                        var v15 uint32 = 2
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v15
                    case "hello":
                        switch 32767 {
                        case 42:
                            p1 = (p1 | 32767)
                        case 4294967295:
                            v7 = v7
                        default:
                            v16 := GenS1{F0: -119, F1: 2, F2: p0}
                            if v7 {
                                fmt.Println(v7, p0, v6)
                            }
                            if (v7 || v7) {
                                p0 += p0
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                if v7 {
                                    v5 = (fi9 - (fi11 + (^(fi11))))
                                    v8 = append(v8, (v5 / ((9223372036854775807) | 1)))
                                    _ = len(v8)
                                }
                            }
                            _ = v16
                        }
                    case "\n\t":
                        var v17 uint16 = 2
                        _ = v17
                    default:
                        switch v5 {
                        case 32768:
                            fmt.Println(v6)
                        case 42:
                            if (p0 == (-(p0))) {
                                v5 = v5
                            } else {
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                if true {
                                    if (v7 || false) {
                                        var v18 uint = 65535
                                        _ = v18
                                    } else {
                                        fmt.Println(v5, v7)
                                        v19 := GenS1{F0: 1, F1: 937, F2: 806}
                                        v5 = fi11
                                        _ = v19
                                    }
                                    switch (fi9 / ((v5) | 1)) {
                                    case -1:
                                        switch ("\n\t" + "a") {
                                        case "foo bar":
                                            v6 += v6
                                        case "\n\t":
                                            v6 = v6
                                        case "":
                                            var v20 bool = (18446744073709551615 == 42)
                                            _ = v20
                                        }
                                    case 504:
                                        if v7 {
                                            v6 -= -2.25
                                            var v21 int16 = p1
                                            _ = v21
                                        }
                                        fmt.Println(p1, p1)
                                    }
                                }
                                switch "a" {
                                case "foo bar":
                                    var v22 int64 = (p0 - (p0 & (p0 << ((964) & 63))))
                                    _ = v22
                                default:
                                    p1 = (p1 % ((567) | 1))
                                }
                                var v23 uint = 42
                                _ = v23
                            }
                            var v24 float64 = (float64(v6) / (((1.5 - -2.25) / 1.5) / ((3.14159 + ((0.1 * 0.0) - (0.1 * 0.1))) - (3.14159 * (1e10 + -2.25)))))
                            var v25 int = fi9
                            v5 = 1
                            _ = v24
                            _ = v25
                        case -9223372036854775808:
                            if v7 {
                                v26 := GenS2{F0: (^(p1)), F1: 42, F2: 2147483647}
                                if v7 {
                                    v27 := GenS2{F0: p1, F1: 581, F2: 2}
                                    _ = v27
                                }
                                _ = v26
                            }
                        default:
                            if (((((^(p1)) != (p1 << ((2) & 63))) && (!false)) || (127 >= 0)) && v7) {
                                v28 := GenS1{F0: -128, F1: 42, F2: (p0 - (-(p0)))}
                                _ = v28
                            }
                        }
                        var v29 int64 = (p0 ^ (p0 + (p0 % (((-(p0))) | 1))))
                        var v30 float64 = 1e10
                        _ = v29
                        _ = v30
                    }
                }
                fmt.Println(fi9, p1)
                _ = fi11
            }
            var v31 bool = (v7 || (42 != -61))
            _ = fi9
            _ = v10
            _ = v31
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint) uint8 {
    v32 := GenS3{F0: -128}
    var v33 uint = 18446744073709551615
    p0 = (p0 / ((319) | 1))
    v34 := GenS1{F0: -61, F1: 48, F2: -1}
    for fi35 := 0; fi35 < 4; fi35++ {
        v36 := []int{}
        _ = fi35
        _ = v36
    }
    _ = p0
    _ = v32
    _ = v33
    _ = v34
    return 128
}

func main() {
    v1 := []int{1}
    var v2 int8 = -128
    var v3 bool = (v2 != v2)
    var v4 int8 = v2
    switch 127 {
    case -1:
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "hello":
            v37 := GenS3{F0: (v2 << ((670) & 63))}
            _ = v37
        default:
            var v38 int32 = 0
            _ = v38
        }
    case 32767:
        var v39 uint8 = 128
        _ = v39
    case -273:
        var v40 uint32 = 2147483648
        _ = v40
    default:
        v41 := GenS2{F0: -1, F1: 936, F2: 65535}
        _ = v41
    }
    var v42 uint16 = 0
    var v43 float32 = 3.14159
    if len(v1) > 0 {
        v1[0] = -482
    }
    var v44 int = 299
    if (!v3) {
        var v45 bool = (2 != 122)
        _ = v45
    } else {
        fmt.Println(v43)
        switch v44 {
        case 32768:
            var v46 int64 = 0
            _ = v46
        case 842:
            var v47 int16 = -1
            _ = v47
        }
    }
    v44 = (v44 + (v44 | v44))
    if false {
        var v48 int = 289
        _ = v48
    } else {
        v49 := GenS2{F0: 42, F1: 840, F2: 2147483648}
        _ = v49
    }
    v50 := GenS4{F0: v2}
    if v3 {
        for fi51 := 0; fi51 < 4; fi51++ {
            for fi52 := 0; fi52 < 1; fi52++ {
                v2 += 0
                if v3 {
                    fmt.Println(v3, v2)
                    v1 = append(v1, -659)
                    var v53 int16 = 135
                    _ = v53
                }
                _ = fi52
            }
            _ = fi51
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (v3 && (507 == 751)) {
            for fi54 := 0; fi54 < 5; fi54++ {
                if v3 {
                    for fi55 := 0; fi55 < 14; fi55++ {
                        switch 9223372036854775807 {
                        case -9223372036854775808:
                            v42 = 1
                        case 386:
                            var v56 uint = 255
                            _ = v56
                        default:
                            var v57 uint = 18446744073709551615
                            _ = v57
                        }
                        if (!false) {
                            if (v44 != (fi54 % (((fi54 & (fi54 / (((fi54 % ((int(v42)) | 1))) | 1)))) | 1))) {
                                var v58 int = fi54
                                _ = v58
                            } else {
                                v3 = v3
                                if v3 {
                                    var v59 int8 = v2
                                    v4 = 127
                                    if (4294967295 <= 0) {
                                        fmt.Println(v43, v44, fi55)
                                        v59 = (v2 ^ v4)
                                        if false {
                                            v44 += v44
                                            v43 = (0.1 + (float32(1e10) / (v43 - v43)))
                                        } else {
                                            var v60 uint32 = 0
                                            _ = v60
                                        }
                                    } else {
                                        if (!v3) {
                                            v59 = (v59 & v2)
                                            switch "\n\t" {
                                            case "hello":
                                                var v61 int64 = 192
                                                _ = v61
                                            default:
                                                switch 9223372036854775807 {
                                                case 127:
                                                    v62 := GenS5{F0: -1, F1: v3}
                                                    v3 = (42 != 415)
                                                    _ = v62
                                                default:
                                                    v4 = v2
                                                }
                                            }
                                            switch v44 {
                                            case 540:
                                                var v63 int = v44
                                                _ = v63
                                            default:
                                                switch v44 {
                                                case 9223372036854775807:
                                                    v64 := GenS6{F0: 255, F1: 2}
                                                    _ = v64
                                                case 127:
                                                    var v65 string = "hello"
                                                    _ = v65
                                                case 128:
                                                    var v66 uint8 = 42
                                                    _ = v66
                                                default:
                                                    var v67 uint64 = genFn0(-645, 51)
                                                    _ = v67
                                                }
                                            }
                                        } else {
                                            _ = len(v1)
                                            v2 = (v4 << ((65535) & 63))
                                            if v3 {
                                                var v68 bool = false
                                                _ = v68
                                            } else {
                                                if v3 {
                                                    v69 := GenS7{F0: (!("a" > ("hello" + ("hello" + "hello"))))}
                                                    fmt.Println(v3, v59, v59)
                                                    fmt.Println(v2, v59)
                                                    _ = v69
                                                } else {
                                                    fmt.Println(fi55, v43)
                                                    v43 = v43
                                                }
                                            }
                                            v4 = v2
                                            v1 = append(v1, (^(fi54)))
                                        }
                                        v70 := GenS2{F0: 849, F1: genFn0(-9223372036854775808, 57), F2: 1}
                                        if (v42 >= (v42 + (^(v42)))) {
                                            v43 -= v43
                                            v1 = append(v1, int(v59))
                                            v2 = (v59 ^ v4)
                                        }
                                        _ = v70
                                    }
                                    _ = v59
                                }
                            }
                            v4 -= 42
                            var v71 uint32 = 2147483648
                            _ = v71
                        }
                        _ = fi55
                    }
                } else {
                    v4 += (v2 | v2)
                    if (false || ((v42 % (((^(v42))) | 1)) == v42)) {
                        var v72 uint = 116
                        _ = v72
                    } else {
                        v42 += v42
                        fmt.Println(v44)
                    }
                }
                _ = fi54
            }
            if (!(v3 || false)) {
                var v73 uint = 542
                v74 := GenS3{F0: (v4 % ((0) | 1))}
                _ = v73
                _ = v74
            }
        } else {
            v4 = 87
        }
    }
    v75 := GenS8{F0: v2}
    if v3 {
        if len(v1) > 0 {
            v1[0] = v44
        }
    }
    v76 := GenS9{F0: 1, F1: -32768, F2: 4294967295}
    var v77 uint16 = 2
    if true {
        var v78 uint8 = 1
        for fi79 := 0; fi79 < 17; fi79++ {
            v80 := GenS10{F0: 0, F1: true, F2: 9223372036854775807}
            _ = fi79
            _ = v80
        }
        var v81 int64 = 42
        _ = v78
        _ = v81
    } else {
        v82 := GenS1{F0: v4, F1: v77, F2: 1}
        var v83 int8 = v4
        _ = v82
        _ = v83
    }
    v4 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v50
    _ = v75
    _ = v76
    _ = v77
}
