package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 string; F1 uint; F2 int16 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32; F1 string }

func main() {
    var v1 int64 = 42
    var v2 uint32 = 4294967295
    var v3 int16 = 1
    var v4 int16 = 1000
    var v5 int8 = 0
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 7; fi7++ {
        v2 -= v2
        _ = fi7
    }
    var v8 uint16 = 597
    if true {
        for fi9 := 0; fi9 < 16; fi9++ {
            if (32767 == -1) {
                var v10 int16 = v3
                _ = v10
            }
            var v11 float32 = -2.25
            switch (fi9 & fi9) {
            case -9223372036854775808:
                switch (fi9 | -892) {
                case 997:
                    v12 := v6[fi9]
                    _ = v12
                case -2147483648:
                    v13 := GenS1{F0: 42, F1: (!false), F2: 1}
                    _ = v13
                default:
                    fmt.Println(v3)
                }
            case 32768:
                var v14 int = int(v3)
                _ = v14
            }
            _ = fi9
            _ = v11
        }
        v15 := GenS2{F0: (true || false)}
        _ = v15
    }
    _ = len(v6)
    v2 -= (v2 >> ((265) & 63))
    v16 := v6[2147483647]
    v17 := GenS1{F0: 0, F1: (42 <= 60), F2: 42}
    switch 32767 {
    case 32768:
        if true {
            v18 := GenS2{F0: (true || (float64((((0.0 + 0.1) + 1.5) / (1e10 / 0.1))) == 1e10))}
            if (!((v8 | v8) <= 368)) {
                switch (v16 + v16) {
                case "\n\t":
                    fmt.Println(v3)
                }
            } else {
                fmt.Println(v2)
                var v19 uint16 = (v8 * v8)
                if true {
                    var v20 int = int(v19)
                    v21 := GenS1{F0: uint8(v2), F1: (1 < 453), F2: 1}
                    _ = v20
                    _ = v21
                } else {
                    v22 := []int{845}
                    var v23 uint32 = v2
                    _ = v22
                    _ = v23
                }
                _ = v19
            }
            _ = v18
        }
    default:
        v24 := GenS1{F0: 2, F1: true, F2: 2}
        _ = v24
    }
    var v25 float32 = 1e10
    if (4294967295 == 508) {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 uint32 = 1
            v28 := GenS1{F0: 128, F1: ((fi26 * 4294967295) != fi26), F2: 128}
            _ = fi26
            _ = v27
            _ = v28
        }
        switch "a" {
        case "":
            if (150 < -2147483648) {
                v29 := []int{}
                var v30 uint8 = 42
                _ = v29
                _ = v30
            } else {
                v8 = v8
                v31 := GenS3{F0: v1, F1: v3, F2: 2}
                for fi32 := 0; fi32 < 17; fi32++ {
                    if (true || true) {
                        var v33 string = v16
                        _ = v33
                    } else {
                        var v34 int = (fi32 | (fi32 * (fi32 / ((fi32) | 1))))
                        if ((uint32(v8) >= (^(v2))) && (!true)) {
                            for fi35 := 0; fi35 < 14; fi35++ {
                                if false {
                                    var v36 uint32 = 2
                                    if true {
                                        v37 := GenS2{F0: ((0.1 * 3.14159) != float64(v25))}
                                        if true {
                                            var v38 string = "hello"
                                            fmt.Println(v25, v4, v16)
                                            fmt.Println(v5, fi32)
                                            if (!true) {
                                                if (18446744073709551615 < 1) {
                                                    v5 = (v5 % ((-68) | 1))
                                                } else {
                                                    v1 = 9223372036854775807
                                                    v6[32767] = v38
                                                    v39 := GenS4{F0: (18446744073709551615 < 444)}
                                                    _ = v39
                                                }
                                            }
                                            _ = v38
                                        } else {
                                            var v40 int8 = v5
                                            if (int(v1) <= v34) {
                                                var v41 string = v16
                                                switch (v41 + v16) {
                                                case "a":
                                                    if false {
                                                        fmt.Println(v25, v40)
                                                        if true {
                                                            fmt.Println(fi32)
                                                            v5 = (-(v40))
                                                            fmt.Println(v34, v25, v8)
                                                            v16 = ("foo bar" + "foo bar")
                                                            fmt.Println(v16, v4, v4)
                                                            var v42 int16 = -993
                                                            var v43 int8 = 1
                                                            v44 := GenS3{F0: (^(v1)), F1: -507, F2: 2}
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            v4 = 577
                                                            v45 := GenS2{F0: (2147483647 < -1)}
                                                            v8 -= (v8 * v8)
                                                            _ = v45
                                                        }
                                                    }
                                                case "":
                                                    if (128 < 42) {
                                                        var v46 float32 = v25
                                                        _ = v46
                                                    }
                                                case "foo bar":
                                                    v4 += (v3 & (v4 | v4))
                                                    v34 = (fi32 << ((2) & 63))
                                                default:
                                                    if (!((v40 >> ((0) & 63)) <= v40)) {
                                                        _ = len(v6)
                                                    }
                                                }
                                                fmt.Println(v4, v4, v4)
                                                v47 := []int{(v34 | (fi35 << ((1) & 63)))}
                                                if len(v47) > 0 {
                                                    v47[0] = 42
                                                }
                                                _ = v41
                                                _ = v47
                                            } else {
                                                v48 := GenS5{F0: (v34 / ((fi35) | 1))}
                                                switch v16 {
                                                case "foo bar":
                                                    v2 += (v36 & v2)
                                                case "":
                                                    if true {
                                                        v4 = (^(v4))
                                                        v1 = int64(v8)
                                                        if true {
                                                            switch "" {
                                                            case "":
                                                                fmt.Println(v4, v4, v25)
                                                            default:
                                                                v2 -= v2
                                                            }
                                                            var v49 int16 = (v3 + (v3 / (((^(v4))) | 1)))
                                                            _ = v49
                                                        } else {
                                                            var v50 bool = false
                                                            v51 := GenS4{F0: ((v2 << ((2) & 63)) == v2)}
                                                            v2 = v2
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    } else {
                                                        v8 -= (v8 << ((1) & 63))
                                                        fmt.Println(v1, v5)
                                                    }
                                                case "hello":
                                                    v52 := GenS6{F0: 172, F1: 255}
                                                    _ = v52
                                                default:
                                                    v40 = -17
                                                }
                                                if true {
                                                    v53 := GenS7{F0: (v16 + ("a" + v16)), F1: 18446744073709551615, F2: (v4 / ((v3) | 1))}
                                                    switch -9223372036854775808 {
                                                    case 127:
                                                        v54 := []int{v34, v34, (v34 * 4294967295), v34}
                                                        _ = v54
                                                    case 0:
                                                        if (true && false) {
                                                            v4 += v3
                                                            v2 = v2
                                                            v55 := GenS4{F0: false}
                                                            if (v34 == 9223372036854775807) {
                                                                v56 := []int{}
                                                                _ = v56
                                                            } else {
                                                                fmt.Println(fi35)
                                                            }
                                                            _ = v55
                                                        } else {
                                                            v1 -= v1
                                                            switch -358 {
                                                            case 1:
                                                                fmt.Println(v36)
                                                            case 9223372036854775807:
                                                                if true {
                                                                    v36 -= v36
                                                                    fmt.Println(v36, v4, v16)
                                                                }
                                                                v57 := GenS8{F0: false, F1: 42}
                                                                _ = v57
                                                            }
                                                            var v58 int16 = v3
                                                            v8 = (v8 * (v8 * 281))
                                                            _ = v58
                                                        }
                                                    default:
                                                        var v59 bool = (!true)
                                                        _ = v59
                                                    }
                                                    v60 := []int{}
                                                    var v61 uint8 = 127
                                                    _ = v53
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v40 += v40
                                                    v62 := GenS9{F0: 2}
                                                    _ = v62
                                                }
                                                _ = v48
                                            }
                                            v5 = (v40 / (((-(v40))) | 1))
                                            _ = v40
                                        }
                                        _ = v37
                                    } else {
                                        var v63 uint32 = uint32(v8)
                                        fmt.Println(v34, v8, v16)
                                        fmt.Println(v2, v25)
                                        _ = v63
                                    }
                                    _ = v36
                                }
                                v64 := GenS10{F0: v2}
                                var v65 int16 = (v4 / (((v3 ^ v3)) | 1))
                                v8 = (^(v8))
                                _ = fi35
                                _ = v64
                                _ = v65
                            }
                            switch v16 {
                            case "foo bar":
                                if ((v8 << ((671) & 63)) >= v8) {
                                    v4 = v3
                                    for fi66 := 0; fi66 < 9; fi66++ {
                                        var v67 uint = 1
                                        v4 = v4
                                        _ = fi66
                                        _ = v67
                                    }
                                    v34 = (fi32 ^ v34)
                                }
                            case "a":
                                v68 := GenS11{F0: v8, F1: 70}
                                _ = v68
                            case "\n\t":
                                v8 -= (v8 - 255)
                            }
                        } else {
                            v25 = ((v25 * v25) + ((3.14159 - (float32(1e10) / 3.14159)) * v25))
                            fmt.Println(v16, v2)
                            v69 := []int{(v34 + -1), v34}
                            _ = v69
                        }
                        for fi70 := 0; fi70 < 6; fi70++ {
                            v4 -= v3
                            v4 = 32767
                            if (2 >= 50) {
                                var v71 bool = (false && (((!true) || true) || false))
                                v5 = (v5 | 127)
                                _ = v71
                            } else {
                                var v72 int = (fi70 + (-(fi70)))
                                var v73 string = ""
                                _ = len(v6)
                                _ = v72
                                _ = v73
                            }
                            _ = fi70
                        }
                        _ = v34
                    }
                    _ = fi32
                }
                _ = v31
            }
        default:
            fmt.Println(v2)
        }
    } else {
        if false {
            v1 = 9223372036854775807
            fmt.Println(v5)
            var v74 int8 = -2
            _ = v74
        }
    }
    v75 := GenS3{F0: (v1 + (v1 + v1)), F1: v4, F2: 42}
    v76 := []int{-314, 4294967295, -2147483648}
    for fi77 := 0; fi77 < 1; fi77++ {
        v5 = (^(v5))
        switch (v16 + v16) {
        case "a":
            var v78 int32 = 1
            _ = v78
        case "hello":
            if ((v8 & (^(v8))) <= v8) {
                v5 = v5
            }
        case "foo bar":
            fmt.Println(v16, v2)
        }
        _ = fi77
    }
    if false {
        for fi79 := 0; fi79 < 12; fi79++ {
            v80 := GenS3{F0: v1, F1: v3, F2: 255}
            fmt.Println(v5, v2, v1)
            if false {
                if ((0.1 / (0.1 * ((0.1 / 0.1) * float64(float32(3.14159))))) < float64((0.0 * 0.1))) {
                    for fi81 := 0; fi81 < 14; fi81++ {
                        v82 := GenS12{F0: v2}
                        _ = fi81
                        _ = v82
                    }
                    var v83 uint8 = 1
                    _ = v83
                } else {
                    v8 = (^(v8))
                    if (v25 >= (v25 * v25)) {
                        v6[fi79] = "\n\t"
                    }
                }
                fmt.Println(fi79, v2)
                var v84 uint16 = v8
                _ = v84
            }
            _ = fi79
            _ = v80
        }
        if false {
            for fi85 := 0; fi85 < 1; fi85++ {
                if len(v76) > 0 {
                    v76[0] = (fi85 >> ((745) & 63))
                }
                var v86 int8 = v5
                for fi87 := 0; fi87 < 17; fi87++ {
                    fmt.Println(fi85)
                    _ = len(v76)
                    v88 := GenS12{F0: (v2 | v2)}
                    v86 += (v86 - 110)
                    _ = fi87
                    _ = v88
                }
                _ = fi85
                _ = v86
            }
            v89 := GenS13{F0: 65535, F1: v16}
            v16 = v16
            _ = v89
        }
        if len(v76) >= 2 {
            v76 = v76[1:len(v76)]
        }
    }
    v8 -= v8
    if ((v8 >> ((0) & 63)) == 646) {
        var v90 float64 = 1.5
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v76:", len(v76), v76)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v8
}
