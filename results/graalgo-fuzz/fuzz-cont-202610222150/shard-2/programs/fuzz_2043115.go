package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS6 struct { F0 int; F1 uint64; F2 string }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint16; F1 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 int32; F1 int; F2 uint8 }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint64; F1 int16; F2 uint }

func main() {
    var v1 string = ((("a" + "") + "a") + ("hello" + "hello"))
    v2 := make(map[int]string)
    var v3 uint16 = 65535
    var v4 float64 = (-2.25 - 1.5)
    var v5 float32 = 0.0
    fmt.Println(v5, v1, v5)
    v3 = v3
    var v6 int64 = 0
    v1 = v1
    fmt.Println(v3, v6, v1)
    for fi7 := 0; fi7 < 1; fi7++ {
        switch v1 {
        case "foo bar":
            var v8 int16 = 42
            _ = v8
        case "a":
            switch v1 {
            case "a":
                switch fi7 {
                case 1:
                    if (1e10 < (v4 * v4)) {
                        var v9 int16 = 1
                        var v10 int16 = (v9 % (((v9 & (v9 % ((1000) | 1)))) | 1))
                        v10 = (v9 | (v10 ^ (v9 - v10)))
                        _ = v10
                        _ = v9
                    } else {
                        var v11 bool = (2147483648 >= 2147483647)
                        for fi12 := 0; fi12 < 15; fi12++ {
                            _ = len(v2)
                            switch fi7 {
                            case 32768:
                                v13 := GenS1{F0: 5}
                                _ = v13
                            case -9223372036854775808:
                                v14 := GenS2{F0: 42, F1: (fi7 * -1)}
                                _ = v14
                            case -2147483648:
                                v15 := GenS1{F0: 160}
                                _ = v15
                            default:
                                v16 := []int{}
                                _ = v16
                            }
                            _ = fi12
                        }
                        for fi17 := 0; fi17 < 14; fi17++ {
                            switch -2147483648 {
                            case 32768:
                                var v18 float32 = 0.0
                                _ = v18
                            case 129:
                                v19 := []int{}
                                _ = v19
                            case 32767:
                                v20 := GenS1{F0: 18446744073709551615}
                                _ = v20
                            default:
                                v21 := GenS2{F0: 1, F1: (fi17 | fi7)}
                                _ = v21
                            }
                            _ = fi17
                        }
                        _ = v11
                    }
                }
            case "\n\t":
                fmt.Println(v4, v4, v3)
            default:
                var v22 float32 = 3.14159
                _ = v22
            }
        default:
            _ = len(v2)
        }
        v23 := GenS3{F0: (fi7 & fi7), F1: -985}
        if false {
            if (true || false) {
                var v24 int = fi7
                _ = v24
            }
            var v25 int8 = 42
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := []int{(fi26 - 2147483647), fi7}
                var v28 float64 = 1.5
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = v25
        } else {
            for fi29 := 0; fi29 < 13; fi29++ {
                v30 := GenS4{F0: 535}
                var v31 float64 = v4
                _ = fi29
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                if false {
                    var v33 float32 = v5
                    switch (v1 + (((v1 + ("a" + v1)) + v1) + ((v1 + "\n\t") + ((v1 + v1) + "a")))) {
                    case "\n\t":
                        fmt.Println(v4)
                    case "foo bar":
                        if true {
                            if (false && ((139 < 127) || false)) {
                                v34 := GenS5{F0: 131, F1: false, F2: (v3 | v3)}
                                _ = v34
                            } else {
                                var v35 uint64 = 120
                                switch (fi32 | (fi7 | fi7)) {
                                case -36:
                                    fmt.Println(v35, v3, v5)
                                }
                                var v36 uint = 42
                                _ = v35
                                _ = v36
                            }
                            v37 := GenS2{F0: 42, F1: (fi7 * (fi32 + fi32))}
                            _ = v37
                        }
                    case "":
                        if (!true) {
                            if (!((!true) || (!(false || true)))) {
                                if (!(!((32767 != (v3 + v3)) || false))) {
                                    if true {
                                        switch v1 {
                                        case "hello":
                                            v3 = v3
                                        case "":
                                            var v38 uint = 0
                                            _ = v38
                                        case "a":
                                            var v39 int8 = 0
                                            _ = v39
                                        }
                                        v40 := GenS5{F0: 255, F1: false, F2: (^(v3))}
                                        _ = v40
                                    }
                                    var v41 uint = 1
                                    _ = v41
                                } else {
                                    fmt.Println(fi7, v33, v3)
                                    v42 := GenS4{F0: 0}
                                    v33 = v5
                                    fmt.Println(v33, v3)
                                    _ = v42
                                }
                                if true {
                                    if false {
                                        if (!true) {
                                            var v43 uint8 = 48
                                            fmt.Println(v33, v3, v5)
                                            _ = v43
                                        }
                                        v33 = float32((v4 - v4))
                                        v44 := []int{-462, fi32, 2147483647, fi32}
                                        _ = v44
                                    }
                                    switch v1 {
                                    case "\n\t":
                                        v6 -= v6
                                    }
                                    v45 := GenS5{F0: 255, F1: (false && true), F2: 2}
                                    _ = v45
                                } else {
                                    var v46 float64 = v4
                                    v47 := GenS6{F0: fi32, F1: 1, F2: v1}
                                    if (!false) {
                                        v4 -= (v4 - -2.25)
                                        var v48 int8 = -128
                                        switch ("foo bar" + "\n\t") {
                                        case "hello":
                                            v46 = v4
                                        case "":
                                            var v49 float32 = 0.0
                                            _ = v49
                                        }
                                        _ = v48
                                    } else {
                                        v4 = 0.0
                                    }
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                var v50 bool = true
                                switch fi32 {
                                case -1:
                                    var v51 uint8 = 42
                                    _ = v51
                                case 963:
                                    v52 := []int{920, (fi7 - -1)}
                                    _ = v52
                                case 988:
                                    if true {
                                        fmt.Println(v5, v4, fi32)
                                        var v53 int = fi7
                                        switch (v53 ^ v53) {
                                        case -9223372036854775808:
                                            if (127 == -1) {
                                                var v54 bool = (!true)
                                                _ = v54
                                            } else {
                                                v55 := []int{fi32, fi32}
                                                if (!false) {
                                                    v5 += (float32((((v4 + ((v4 * (1e10 / v4)) + v4)) / ((v4 * (float64(-2.25) - (v4 - v4))) + 0.1)) + v4)) / ((3.14159 / 1.5) + v33))
                                                } else {
                                                    v4 -= v4
                                                }
                                                if (0.0 == v4) {
                                                    v50 = v50
                                                    var v56 uint32 = 102
                                                    _ = v56
                                                }
                                                _ = v55
                                            }
                                        case 9223372036854775807:
                                            v57 := []int{fi7, fi7, -9223372036854775808}
                                            _ = v57
                                        case -979:
                                            var v58 uint8 = 128
                                            _ = v58
                                        }
                                        _ = v53
                                    } else {
                                        v59 := []int{fi7, int(v6), fi7}
                                        fmt.Println(v6, v33)
                                        _ = v59
                                    }
                                default:
                                    if v50 {
                                        var v60 int16 = 42
                                        if v50 {
                                            if (!(true && (((-9223372036854775808 != (fi7 | 4294967295)) || (42 > 190)) && (v3 <= (v3 - v3))))) {
                                                fmt.Println(v6, v33, v1)
                                                if true {
                                                    fmt.Println(v6)
                                                    v60 -= v60
                                                    v5 -= 1e10
                                                }
                                                v61 := GenS3{F0: int(v3), F1: 519}
                                                _ = v61
                                            }
                                            v62 := []int{fi7, fi7}
                                            _ = v62
                                        }
                                        v63 := GenS7{F0: v3, F1: 1}
                                        _ = v60
                                        _ = v63
                                    } else {
                                        fmt.Println(v33, v4)
                                    }
                                }
                                if ((0.1 + 0.0) > (-2.25 * v33)) {
                                    v50 = (51 > 255)
                                    v64 := GenS8{F0: v3, F1: 18}
                                    _ = v64
                                } else {
                                    var v65 float64 = ((((1.5 * (((0.0 / float64((-2.25 * v5))) * (1e10 + v4)) * -2.25)) + (0.1 / v4)) - float64(v33)) * 1.5)
                                    _ = v65
                                }
                                if (v50 && (65535 == 1)) {
                                    if (float32(v4) == (0.1 + ((v5 + (float32(((3.14159 / 1.5) - ((v4 / 3.14159) - 0.0))) / -2.25)) * v33))) {
                                        if (1 < -2147483648) {
                                            v66 := GenS9{F0: 382}
                                            _ = v66
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            v67 := []int{fi7, 0, fi32, 32767}
                                            _ = v67
                                        case "a":
                                            var v68 float32 = v5
                                            _ = v68
                                        }
                                    } else {
                                        v69 := GenS10{F0: (v1 + "a"), F1: v1}
                                        if false {
                                            v70 := GenS4{F0: 103}
                                            _ = v70
                                        } else {
                                            v71 := GenS11{F0: -1, F1: fi32, F2: 42}
                                            v72 := []int{fi32, (^(fi7)), 2147483647, 0}
                                            _ = v71
                                            _ = v72
                                        }
                                        var v73 int = (-(fi32))
                                        v6 -= (v6 ^ v6)
                                        _ = v69
                                        _ = v73
                                    }
                                    v74 := GenS12{F0: (^(fi7)), F1: -87}
                                    if (0 >= 719) {
                                        v3 = v3
                                        v1 = ("a" + (v1 + "a"))
                                        v75 := GenS13{F0: v50}
                                        _ = v75
                                    } else {
                                        v76 := GenS14{F0: -2147483648}
                                        var v77 int = fi7
                                        v78 := []int{0, -911, fi7, 1}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                    _ = v74
                                }
                                if v50 {
                                    var v79 uint64 = 0
                                    v5 = v5
                                    v80 := GenS9{F0: 851}
                                    var v81 int16 = 1000
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                }
                                _ = v50
                            }
                            var v82 uint8 = 128
                            v83 := GenS15{F0: 2}
                            _ = v82
                            _ = v83
                        } else {
                            var v84 uint = 1
                            if true {
                                switch v1 {
                                case "foo bar":
                                    var v85 uint16 = (v3 * v3)
                                    _ = v85
                                case "a":
                                    if true {
                                        v86 := GenS4{F0: 42}
                                        _ = v86
                                    } else {
                                        v1 = v1
                                        v87 := GenS5{F0: 162, F1: (false || ((v4 * (((v4 + ((((v4 - (v4 / v4)) * 3.14159) / v4) * v4)) * 3.14159) / -2.25)) <= v4)), F2: v3}
                                        _ = v87
                                    }
                                default:
                                    v84 += v84
                                }
                                var v88 bool = true
                                _ = v88
                            }
                            _ = v84
                        }
                    }
                    _ = v33
                }
                v89 := []int{(fi7 / (((fi7 & -328)) | 1)), (fi32 | fi32), -1, fi32}
                var v90 uint16 = v3
                _ = fi32
                _ = v89
                _ = v90
            }
            for fi91 := 0; fi91 < 3; fi91++ {
                var v92 int32 = -2147483648
                var v93 uint8 = 1
                fmt.Println(v4, fi91, v93)
                var v94 uint16 = (v3 / ((508) | 1))
                _ = fi91
                _ = v92
                _ = v93
                _ = v94
            }
        }
        _ = fi7
        _ = v23
    }
    v95 := GenS16{F0: 516, F1: 42, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v95
}
