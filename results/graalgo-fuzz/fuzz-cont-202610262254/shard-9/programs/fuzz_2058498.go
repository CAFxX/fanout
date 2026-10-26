package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int16; F2 uint8 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 0.0
    var v3 string = ("a" + "foo bar")
    v4 := []int{-421, 127, 489}
    if true {
        v2 += 0.1
    } else {
        switch 42 {
        case -9223372036854775808:
            var v5 uint32 = 4294967295
            _ = v5
        case 2147483647:
            var v6 int8 = 7
            _ = v6
        case 500:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int64 = -1
            _ = fi7
            _ = v8
        }
        v3 = (v3 + ("a" + (v3 + ((v3 + ("hello" + (v3 + v3))) + v3))))
    }
    if len(v4) > 0 {
        v4[0] = -772
    }
    switch -853 {
    case -52:
        if true {
            if false {
                var v9 uint = 4294967295
                if (false || false) {
                    fmt.Println(v3)
                    _ = len(v4)
                }
                var v10 uint8 = 1
                _ = v10
                _ = v9
            } else {
                var v11 float64 = v2
                var v12 int16 = 1000
                if ((v1 ^ 9223372036854775807) <= (^(v1))) {
                    var v13 int8 = -1
                    _ = v13
                } else {
                    switch ("" + "") {
                    case "hello":
                        var v14 int8 = 109
                        _ = v14
                    default:
                        v15 := GenS1{F0: (v3 <= ("a" + "a")), F1: 0, F2: 0}
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v12 += (v12 + 32767)
                        _ = fi16
                    }
                }
                _ = v11
                _ = v12
            }
            fmt.Println(v3, v3, v3)
        }
    case 859:
        if (326 != 32767) {
            if false {
                for fi17 := 0; fi17 < 11; fi17++ {
                    v3 = v3
                    _ = fi17
                }
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v18 := GenS2{F0: 0}
            switch 160 {
            case -2147483648:
                v19 := GenS1{F0: true, F1: 0, F2: -2147483648}
                _ = v19
            default:
                if true {
                    v3 = "hello"
                }
            }
            _ = v18
        }
    case -380:
        switch -823 {
        case -97:
            if (127 >= -9223372036854775808) {
                for fi20 := 0; fi20 < 1; fi20++ {
                    if true {
                        fmt.Println(v2, v1, v2)
                        v21 := GenS2{F0: -1}
                        _ = v21
                    } else {
                        var v22 uint8 = 127
                        if (v2 > v2) {
                            _ = len(v4)
                        }
                        if false {
                            var v23 uint32 = 65535
                            var v24 float64 = v2
                            fmt.Println(v24)
                            _ = v23
                            _ = v24
                        }
                        _ = v22
                    }
                    v1 -= v1
                    _ = fi20
                }
            } else {
                v25 := GenS1{F0: ((-80 > 981) || false), F1: 32767, F2: -1}
                var v26 float64 = v2
                _ = v25
                _ = v26
            }
        default:
            fmt.Println(v3, v3)
        }
    default:
        v2 -= (v2 + ((1.5 + (((((((((-2.25 / (1e10 / v2)) * v2) * -2.25) * ((v2 * 1e10) - 3.14159)) - v2) - v2) * (0.0 - (v2 / float64((((0.0 + -2.25) / -2.25) / 1e10))))) - v2) * (float64(1.5) - (v2 - v2)))) * v2))
    }
    v2 = 0.1
    var v27 uint64 = 42
    switch v3 {
    case "\n\t":
        if false {
            var v28 uint32 = 2147483648
            _ = v28
        }
    default:
        fmt.Println(v1)
    }
    var v29 uint64 = v27
    for fi30 := 0; fi30 < 11; fi30++ {
        v31 := GenS3{F0: v1, F1: 27, F2: fi30}
        for fi32 := 0; fi32 < 2; fi32++ {
            if true {
                switch "a" {
                case "foo bar":
                    v2 += 1e10
                case "":
                    var v33 string = v3
                    _ = v33
                }
            } else {
                if false {
                    v3 = ""
                    switch (fi30 << ((4294967295) & 63)) {
                    case -9223372036854775808:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    var v34 int16 = -1
                    if true {
                        if false {
                            if ((-1 >= -96) && (((1e10 * 0.1) - 0.1) >= 1.5)) {
                                fmt.Println(v3, v3)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                if ("" < "\n\t") {
                                    var v35 uint16 = 377
                                    switch fi30 {
                                    case 629:
                                        v36 := GenS1{F0: (631 > 2), F1: 367, F2: (^(fi32))}
                                        _ = v36
                                    default:
                                        v37 := GenS4{F0: v27}
                                        _ = v37
                                    }
                                    var v38 float64 = float64(3.14159)
                                    v34 = v34
                                    v39 := GenS5{F0: "\n\t"}
                                    var v40 int16 = (v34 / (((v34 / ((v34) | 1))) | 1))
                                    _ = v35
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                                var v41 uint64 = (^(v29))
                                var v42 float64 = 3.14159
                                _ = v41
                                _ = v42
                            }
                            _ = len(v4)
                            if (!true) {
                                switch (fi32 ^ 761) {
                                case -2147483648:
                                    if len(v4) > 0 {
                                        v4[0] = (fi32 * 569)
                                    }
                                case 1:
                                    var v43 int = fi32
                                    _ = v43
                                case 32767:
                                    if (37 <= 1) {
                                        v44 := GenS3{F0: (v1 * (^(v1))), F1: 100, F2: fi30}
                                        var v45 int16 = (^(v34))
                                        v46 := GenS5{F0: v3}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                default:
                                    v47 := GenS3{F0: v1, F1: 42, F2: fi30}
                                    _ = v47
                                }
                                var v48 float64 = ((((v2 - float64(float32(((0.1 / (0.0 + (3.14159 + 0.1))) + (float64(((0.1 / 0.1) - 1e10)) / 0.1))))) / 1.0) - 3.14159) + ((float64(3.14159) * 1e10) / (0.1 / v2)))
                                v4 = append(v4, fi30)
                                v27 += v29
                                v4 = append(v4, fi32)
                                _ = v48
                            } else {
                                v1 = v1
                                v3 = v3
                            }
                            switch v3 {
                            case "":
                                fmt.Println(v34, v27, fi32)
                            case "hello":
                                v49 := GenS2{F0: 42}
                                _ = v49
                            case "\n\t":
                                if (255 > 0) {
                                    var v50 uint = 772
                                    _ = v50
                                }
                            }
                        }
                    } else {
                        if false {
                            var v51 string = "foo bar"
                            _ = v51
                        }
                        var v52 float32 = (3.14159 / ((3.14159 * (1e10 / 1e10)) + 1.5))
                        v27 = 4294967295
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v29 = v29
                        v53 := GenS6{F0: 255, F1: v34, F2: 1}
                        _ = v52
                        _ = v53
                    }
                    v4 = append(v4, (fi32 & (fi32 % ((128) | 1))))
                    v27 -= (v29 * 2)
                    v1 -= (v1 - (v1 / ((v1) | 1)))
                    _ = v34
                } else {
                    fmt.Println(v27, fi32)
                    v54 := GenS2{F0: int8(fi32)}
                    _ = v54
                }
            }
            v2 -= (((3.14159 / v2) + v2) + v2)
            if true {
                v55 := GenS5{F0: "\n\t"}
                _ = v55
            } else {
                if false {
                    switch (fi30 | fi30) {
                    case -1:
                        fmt.Println(fi30, v1)
                    default:
                        var v56 int32 = 0
                        _ = v56
                    }
                } else {
                    var v57 int = (fi30 << ((65535) & 63))
                    var v58 int8 = 100
                    _ = v57
                    _ = v58
                }
                fmt.Println(v29, v2, fi30)
            }
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    var v59 float64 = v2
    var v60 int64 = v1
    if ((true || ((("" == v3) && true) && false)) && (3.14159 != -2.25)) {
        v27 = v29
        fmt.Println(v60, v27, v59)
        for fi61 := 0; fi61 < 10; fi61++ {
            var v62 float32 = 1.5
            _ = fi61
            _ = v62
        }
    } else {
        var v63 int = -9223372036854775808
        _ = v63
    }
    var v64 uint16 = 42
    _ = len(v4)
    var v65 int64 = v60
    fmt.Println(v59, v27, v1)
    v4 = append(v4, 4294967295)
    v65 += (v1 * 42)
    v1 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v64
    _ = v65
}
