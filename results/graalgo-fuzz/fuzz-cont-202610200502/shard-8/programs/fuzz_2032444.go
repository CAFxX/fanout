package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 string; F1 int8; F2 uint }
type GenS11 struct { F0 int64; F1 int16; F2 bool }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8; F1 uint32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int8; F1 bool; F2 uint64 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]uint32)
    var v3 uint = 42
    var v4 float32 = (1.5 + -2.25)
    v5 := make(map[string]int)
    var v6 float32 = (((v4 / 3.14159) / float32((-2.25 - ((float64(v4) - (0.0 - 1e10)) + (1.5 + 0.0))))) - 1e10)
    var v7 int = 32768
    var v8 float64 = ((0.1 - (0.1 - 0.0)) + ((float64(-2.25) + 0.0) / 1e10))
    for fi9 := 0; fi9 < 16; fi9++ {
        switch v7 {
        case 2147483647:
            var v10 uint32 = 0
            _ = v10
        case 42:
            v8 = (1.5 * (v8 / (v8 - ((v8 + 1.5) - float64(v6)))))
        case 1:
            v11 := GenS1{F0: 71, F1: 1}
            _ = v11
        default:
            v7 -= (^(v7))
        }
        var v12 int8 = -68
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            if (2 >= 227) {
                v13 := GenS1{F0: (-(v12)), F1: 163}
                fmt.Println(v8, v12)
                _ = v13
            } else {
                v4 += (v6 - v4)
            }
        case "":
            if (!true) {
                v1 += 268
                v3 = v3
            }
        }
        _ = fi9
        _ = v12
    }
    v14 := []int{-280}
    switch "foo bar" {
    case "foo bar":
        v8 = v8
    case "hello":
        fmt.Println(v3)
    case "":
        v1 = 42
    }
    var v15 int32 = -2147483648
    _ = len(v2)
    switch "foo bar" {
    case "a":
        var v16 float32 = 1e10
        _ = v16
    case "":
        v15 = (-(v15))
    }
    var v17 uint8 = 2
    if true {
        v1 += (v1 >> ((uint(v17)) & 63))
        fmt.Println(v3, v4)
        v15 -= v15
    }
    var v18 float64 = v8
    v19 := GenS1{F0: 42, F1: 2147483647}
    switch ("a" + ("\n\t" + ("a" + ""))) {
    case "\n\t":
        if false {
            v20 := GenS1{F0: -1, F1: 1}
            if false {
                fmt.Println(v4, v15, v3)
                v3 -= v3
            }
            _ = v20
        }
    case "":
        _ = len(v14)
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if false {
            if len(v14) > 0 {
                v14[0] = 127
            }
            if true {
                switch fi21 {
                case -1:
                    v15 = v15
                case 42:
                    v22 := GenS2{F0: (v15 - v15), F1: -1}
                    _ = v22
                }
                fmt.Println(v1, v18)
            } else {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v8 = (1e10 * (v8 * v8))
                    if true {
                        var v24 string = "\n\t"
                        _ = v24
                    } else {
                        if true {
                            var v25 string = ""
                            if (1.5 >= v8) {
                                switch v25 {
                                case "\n\t":
                                    var v26 int32 = (v15 / ((v15) | 1))
                                    _ = v26
                                case "foo bar":
                                    v27 := GenS1{F0: 120, F1: 42}
                                    _ = v27
                                default:
                                    if false {
                                        fmt.Println(fi21, v18)
                                    } else {
                                        v28 := GenS1{F0: 0, F1: 0}
                                        v8 = 1.5
                                        v15 = (v15 ^ (v15 >> ((65535) & 63)))
                                        _ = v28
                                    }
                                }
                                _ = len(v5)
                                fmt.Println(v25)
                            } else {
                                switch fi21 {
                                case 1:
                                    if (210 > 1000) {
                                        switch v25 {
                                        case "foo bar":
                                            switch (fi21 >> ((v3) & 63)) {
                                            case -9223372036854775808:
                                                v7 += (fi23 ^ (fi21 & fi21))
                                                switch (v7 | 0) {
                                                case 4294967295:
                                                    fmt.Println(v1)
                                                case 42:
                                                    var v29 uint64 = 467
                                                    var v30 uint8 = v17
                                                    v31 := GenS3{F0: v30}
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                }
                                            case -628:
                                                switch (v25 + v25) {
                                                case "foo bar":
                                                    v8 = ((v18 + (v8 + (float64(float32(((float64((v4 - (((1.5 + (v4 * 1e10)) / v6) - v6))) / v8) * float64(v4)))) - (v18 - v8)))) / float64((float32(1e10) / v6)))
                                                case "a":
                                                    v32 := GenS2{F0: v15, F1: -32768}
                                                    _ = v32
                                                default:
                                                    var v33 float64 = v18
                                                    _ = v33
                                                }
                                            case 32768:
                                                switch v25 {
                                                case "a":
                                                    if len(v14) > 0 {
                                                        v14[0] = fi21
                                                    }
                                                case "\n\t":
                                                    v14 = append(v14, (v7 << (((v3 + 42)) & 63)))
                                                }
                                            }
                                        case "hello":
                                            v18 = v8
                                        case "a":
                                            v3 = 628
                                        default:
                                            v7 -= -60
                                        }
                                        if true {
                                            switch 127 {
                                            case 1:
                                                var v34 uint32 = 65535
                                                _ = v34
                                            case 128:
                                                var v35 int16 = -1
                                                _ = v35
                                            default:
                                                var v36 int = -140
                                                _ = v36
                                            }
                                        } else {
                                            var v37 int64 = 9223372036854775807
                                            _ = v37
                                        }
                                        if true {
                                            switch 448 {
                                            case 127:
                                                v38 := GenS3{F0: (v17 * v17)}
                                                _ = v38
                                            case 4294967295:
                                                v25 = "\n\t"
                                                var v39 string = "a"
                                                v4 = (0.0 / v6)
                                                v15 -= v15
                                                v40 := GenS4{F0: (v15 / (((v15 + (v15 >> ((4294967295) & 63)))) | 1)), F1: v1, F2: 100}
                                                _ = v39
                                                _ = v40
                                            case 667:
                                                if (176 >= 327) {
                                                    v3 = v3
                                                    switch v7 {
                                                    case 4294967295:
                                                        var v41 uint16 = 255
                                                        _ = v41
                                                    case 306:
                                                        v1 += 2
                                                    case 2147483647:
                                                        v5[v25] = fi23
                                                    default:
                                                        v25 = "foo bar"
                                                    }
                                                    v42 := GenS4{F0: v15, F1: (^(v1)), F2: -128}
                                                    _ = v42
                                                }
                                            default:
                                                _ = len(v14)
                                            }
                                        } else {
                                            v43 := GenS5{F0: 255, F1: v3}
                                            _ = len(v5)
                                            v44 := GenS6{F0: 814}
                                            _ = v43
                                            _ = v44
                                        }
                                        _ = len(v2)
                                        if ((v1 % ((65535) | 1)) == v1) {
                                            v45 := GenS1{F0: 100, F1: uint32(v3)}
                                            _ = v45
                                        } else {
                                            v4 += (v4 + v6)
                                            v15 -= 2147483647
                                            v46 := GenS5{F0: (v3 / (((v3 >> ((531) & 63))) | 1)), F1: v3}
                                            if true {
                                                v47 := GenS5{F0: (v3 * v3), F1: v3}
                                                var v48 uint8 = uint8(v3)
                                                _ = v47
                                                _ = v48
                                            } else {
                                                var v49 uint64 = 1
                                                if false {
                                                    fmt.Println(v1, v15, fi23)
                                                    switch "hello" {
                                                    case "\n\t":
                                                        fmt.Println(v15)
                                                    case "a":
                                                        if false {
                                                            if false {
                                                                v18 = v18
                                                            }
                                                        }
                                                    }
                                                    var v50 float64 = 1.5
                                                    _ = v50
                                                } else {
                                                    v14 = append(v14, -2147483648)
                                                }
                                                switch 127 {
                                                case 4294967295:
                                                    v51 := GenS4{F0: v15, F1: v1, F2: 1}
                                                    _ = v51
                                                default:
                                                    var v52 int64 = 42
                                                    _ = v52
                                                }
                                                _ = v49
                                            }
                                            if ((2147483648 != uint32(v17)) && false) {
                                                var v53 float32 = 0.1
                                                v17 -= (v17 % (((v17 + (v17 / ((v17) | 1)))) | 1))
                                                v54 := GenS7{F0: int64(v17), F1: 0}
                                                if true {
                                                    v55 := GenS1{F0: 100, F1: 1}
                                                    _ = v55
                                                }
                                                _ = v53
                                                _ = v54
                                            } else {
                                                var v56 uint = 42
                                                var v57 uint64 = 713
                                                v25 = (v25 + "")
                                                _ = v56
                                                _ = v57
                                            }
                                            switch v7 {
                                            case 766:
                                                v17 = (v17 * (^(v17)))
                                            }
                                            var v58 uint16 = v1
                                            _ = v46
                                            _ = v58
                                        }
                                        if len(v14) > 0 {
                                            v14[0] = (v7 ^ (-(v7)))
                                        }
                                    }
                                case 32767:
                                    v59 := GenS8{F0: v25, F1: v1}
                                    _ = v59
                                default:
                                    _ = len(v14)
                                }
                                if true {
                                    v2[417] = 728
                                    v1 = 2
                                } else {
                                    if (v18 == ((0.0 / v8) * 1e10)) {
                                        if (!(!true)) {
                                            v60 := GenS9{F0: v1}
                                            if false {
                                                v61 := GenS10{F0: v25, F1: 42, F2: (v3 + v3)}
                                                switch v7 {
                                                case -2147483648:
                                                    var v62 bool = (9223372036854775807 == 6)
                                                    _ = v62
                                                default:
                                                    var v63 int = (v7 - 42)
                                                    _ = v63
                                                }
                                                v15 = -2147483648
                                                _ = v61
                                            } else {
                                                v7 = (-(v7))
                                            }
                                            _ = v60
                                        }
                                    }
                                    var v64 uint32 = 2147483648
                                    _ = v64
                                }
                                var v65 uint64 = 291
                                _ = v65
                            }
                            if (!true) {
                                if ((3.14159 / 1.0) != (((v4 + 0.0) * 1.5) + ((v6 / v6) * v4))) {
                                    v66 := GenS11{F0: -194, F1: 1, F2: (v3 == 65535)}
                                    v67 := GenS12{F0: (((v25 + (v25 + v25)) + v25) + "hello")}
                                    _ = v66
                                    _ = v67
                                }
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                var v68 int16 = 152
                                _ = v68
                            }
                            var v69 string = "a"
                            _ = v25
                            _ = v69
                        } else {
                            var v70 uint64 = 146
                            _ = v70
                        }
                    }
                    v1 += (v1 + 32767)
                    var v71 float64 = v18
                    if (false && true) {
                        var v72 int = (fi21 / (((v7 << ((42) & 63))) | 1))
                        delete(v2, uint32(v17))
                        if false {
                            v73 := GenS13{F0: 127, F1: 83}
                            var v74 float32 = v6
                            v3 += 65535
                            v75 := GenS14{F0: 553}
                            _ = v73
                            _ = v74
                            _ = v75
                        } else {
                            var v76 uint8 = (v17 ^ (v17 >> (((v3 ^ (v3 & 453))) & 63)))
                            fmt.Println(fi21, v71)
                            v76 = v76
                            _ = v76
                        }
                        _ = v72
                    } else {
                        v77 := GenS15{F0: 42, F1: true, F2: 0}
                        if ((-920 != 32767) && (-1 != 67)) {
                            var v78 int16 = 32767
                            var v79 int64 = -1
                            _ = v78
                            _ = v79
                        }
                        _ = v77
                    }
                    _ = fi23
                    _ = v71
                }
            }
        }
        var v80 bool = true
        _ = fi21
        _ = v80
    }
    v17 = (v17 >> ((v3) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
