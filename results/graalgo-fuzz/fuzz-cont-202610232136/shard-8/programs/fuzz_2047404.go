package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int16; F1 bool; F2 string }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint8; F1 int8; F2 int32 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0(p0 uint8, p1 bool) string {
    fmt.Println(p0, p1)
    var v4 float64 = -2.25
    var v5 int = 4294967295
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ""
}
func genFn1() uint8 {
    var v6 int8 = 1
    switch "\n\t" {
    case "hello":
        switch 32768 {
        case 0:
            v7 := []int{32768}
            _ = v7
        case 250:
            fmt.Println(v6)
        case 4294967295:
            v6 = 120
        default:
            if (828 >= 0) {
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 uint8 = 119
                    v6 = -1
                    v10 := []int{int(v6), 4294967295, (fi8 & (fi8 ^ (^(fi8)))), (fi8 & -1)}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        }
    case "\n\t":
        v11 := []int{-708, -9223372036854775808, 1, 630}
        _ = v11
    case "":
        switch ("a" + "a") {
        case "hello":
            fmt.Println(v6)
        case "a":
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: 4294967295, F1: 2}
                v14 := GenS2{F0: 2, F1: 2147483647, F2: ("foo bar" + "a")}
                _ = fi12
                _ = v13
                _ = v14
            }
        case "":
            for fi15 := 0; fi15 < 8; fi15++ {
                fmt.Println(v6, v6)
                var v16 int64 = 922
                for fi17 := 0; fi17 < 7; fi17++ {
                    switch (fi15 - (fi15 / ((fi17) | 1))) {
                    case -208:
                        switch (-(fi17)) {
                        case -2147483648:
                            var v18 uint64 = 18446744073709551615
                            _ = v18
                        case 127:
                            var v19 bool = true
                            var v20 uint32 = 0
                            _ = v19
                            _ = v20
                        case 593:
                            v16 += (v16 ^ -9223372036854775808)
                        default:
                            v16 -= (v16 ^ v16)
                        }
                        v16 -= -9223372036854775808
                        var v21 float64 = 1e10
                        _ = v21
                    case -73:
                        v22 := []int{9223372036854775807, fi17}
                        _ = v22
                    }
                    var v23 int8 = (v6 & v6)
                    v6 = v23
                    _ = fi17
                    _ = v23
                }
                _ = fi15
                _ = v16
            }
        default:
            fmt.Println(v6)
        }
    default:
        v24 := []int{-229, 129}
        _ = v24
    }
    v25 := GenS2{F0: 2, F1: -1, F2: "foo bar"}
    v26 := GenS3{F0: 1, F1: 0}
    var v27 uint16 = 1
    _ = v25
    _ = v26
    _ = v27
    _ = v6
    return 42
}

func main() {
    var v1 string = "a"
    var v2 int16 = 42
    var v3 int16 = 1
    v28 := []int{128, 2147483647, -900}
    v29 := GenS4{F0: v3, F1: (false && false), F2: ("" + (v1 + (v1 + v1)))}
    for fi30 := 0; fi30 < 15; fi30++ {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v3, fi30)
        case "\n\t":
            if (0.0 != 0.0) {
                v2 = (v2 << ((0) & 63))
                var v31 uint64 = 1
                _ = v31
            }
        }
        _ = fi30
    }
    v32 := GenS5{F0: -62, F1: 65535}
    v28 = append(v28, -204)
    switch (("" + v1) + v1) {
    case "\n\t":
        var v33 int = 0
        _ = v33
    }
    var v34 int32 = -1
    fmt.Println(v1)
    fmt.Println(v1, v3, v3)
    var v35 int = -194
    var v36 int64 = -1
    v28 = append(v28, -2147483648)
    switch (v1 + (((v1 + (v1 + (v1 + v1))) + ((v1 + "hello") + "foo bar")) + v1)) {
    case "":
        if false {
            v37 := GenS2{F0: 300, F1: int32(v35), F2: "\n\t"}
            _ = v37
        }
    case "a":
        if true {
            var v38 bool = (976 <= 0)
            _ = v38
        } else {
            v3 = (v2 >> ((1) & 63))
        }
    case "foo bar":
        for fi39 := 0; fi39 < 12; fi39++ {
            if ((^(v36)) <= v36) {
                var v40 uint = 0
                var v41 float64 = 0.1
                _ = v40
                _ = v41
            } else {
                v42 := GenS6{F0: v3, F1: 65535}
                v43 := GenS5{F0: v36, F1: 1}
                _ = v42
                _ = v43
            }
            if len(v28) > 0 {
                v28[0] = (v35 << ((18446744073709551615) & 63))
            }
            _ = fi39
        }
    default:
        var v44 uint32 = 65535
        _ = v44
    }
    v34 -= (-(v34))
    for fi45 := 0; fi45 < 3; fi45++ {
        switch (("hello" + (v1 + (v1 + "\n\t"))) + v1) {
        case "foo bar":
            var v46 int32 = v34
            _ = v46
        case "":
            for fi47 := 0; fi47 < 10; fi47++ {
                var v48 int16 = v2
                if false {
                    var v49 uint64 = 933
                    _ = v49
                } else {
                    v28 = append(v28, fi45)
                }
                v3 = v3
                _ = fi47
                _ = v48
            }
        case "a":
            if true {
                v50 := GenS7{F0: 255, F1: 1, F2: (v34 % ((-130) | 1))}
                _ = v50
            } else {
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
            }
        default:
            switch fi45 {
            case 127:
                for fi51 := 0; fi51 < 16; fi51++ {
                    if (((3.14159 + 0.0) / 1e10) >= -2.25) {
                        v52 := GenS5{F0: (v36 | (-(v36))), F1: 42}
                        v36 = (v36 - -9223372036854775808)
                        _ = v52
                    }
                    v53 := GenS5{F0: v36, F1: 18446744073709551615}
                    _ = fi51
                    _ = v53
                }
            case 0:
                if (false && true) {
                    if true {
                        var v54 uint8 = 152
                        for fi55 := 0; fi55 < 7; fi55++ {
                            v56 := GenS6{F0: (^(v2)), F1: 4294967295}
                            _ = fi55
                            _ = v56
                        }
                        switch ((v1 + v1) + v1) {
                        case "hello":
                            v57 := GenS8{F0: 2, F1: 0, F2: -96}
                            _ = v57
                        default:
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                        }
                        _ = v54
                    } else {
                        var v58 uint16 = 255
                        switch v35 {
                        case 127:
                            v34 = v34
                        }
                        for fi59 := 0; fi59 < 6; fi59++ {
                            switch fi45 {
                            case 4294967295:
                                var v60 float32 = 0.0
                                var v61 uint64 = 4294967295
                                if false {
                                    if false {
                                        var v62 float32 = float32(1.5)
                                        _ = v62
                                    }
                                    v36 = v36
                                    v60 = float32(-2.25)
                                } else {
                                    if true {
                                        v63 := GenS9{F0: (v58 ^ 1), F1: (^(v34))}
                                        v64 := GenS8{F0: (v58 >> ((2) & 63)), F1: 2, F2: 85}
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                _ = v60
                                _ = v61
                            case 32767:
                                v3 += int16(v34)
                            }
                            v65 := GenS10{F0: (v34 - 2147483647), F1: v58, F2: v34}
                            var v66 uint64 = 18446744073709551615
                            if len(v28) > 0 {
                                v28[0] = -338
                            }
                            _ = fi59
                            _ = v65
                            _ = v66
                        }
                        _ = v58
                    }
                    for fi67 := 0; fi67 < 17; fi67++ {
                        if (fi45 != (fi45 >> ((42) & 63))) {
                            if false {
                                var v68 string = v1
                                var v69 bool = (1 <= 65535)
                                _ = v68
                                _ = v69
                            }
                            v35 = -1
                            switch (^(v35)) {
                            case 9223372036854775807:
                                fmt.Println(v2, fi67, fi45)
                            case 42:
                                v2 = 64
                            default:
                                var v70 uint32 = 653
                                _ = v70
                            }
                            if true {
                                var v71 string = "a"
                                switch (v35 - fi45) {
                                case -1:
                                    if (!(false || (false || true))) {
                                        switch (fi67 - int(v34)) {
                                        case 0:
                                            if (!(!false)) {
                                                var v72 float64 = -2.25
                                                fmt.Println(v36, v36)
                                                _ = v72
                                            }
                                        case -9223372036854775808:
                                            var v73 float64 = 3.14159
                                            _ = v73
                                        default:
                                            var v74 uint64 = 2
                                            _ = v74
                                        }
                                    } else {
                                        v34 = (-(v34))
                                    }
                                case 32767:
                                    switch (v71 + (("hello" + v1) + (v1 + v71))) {
                                    case "a":
                                        var v75 uint8 = 1
                                        _ = v75
                                    default:
                                        if len(v28) > 0 {
                                            v28[0] = v35
                                        }
                                    }
                                case 127:
                                    var v76 int8 = -128
                                    _ = v76
                                }
                                _ = v71
                            } else {
                                _ = len(v28)
                                _ = len(v28)
                                v34 = v34
                            }
                        }
                        _ = fi67
                    }
                    if true {
                        v1 = "hello"
                        v28 = append(v28, v35)
                    }
                }
            default:
                fmt.Println(v34, fi45)
            }
        }
        var v77 uint32 = 1
        _ = fi45
        _ = v77
    }
    v78 := GenS8{F0: 0, F1: 4294967295, F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v78
}
