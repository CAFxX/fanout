package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 int64; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 int; F1 int; F2 int64 }
type GenS8 struct { F0 int8; F1 uint16; F2 int8 }
func genFn0() int {
    var v6 uint64 = 18446744073709551615
    var v7 int8 = -1
    var v8 uint16 = 255
    v9 := []int{75}
    if len(v9) > 0 {
        v9[0] = 4294967295
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 923
}

func main() {
    var v1 float64 = 0.1
    var v2 int32 = 2147483647
    var v3 int = 0
    var v4 float32 = float32(3.14159)
    var v5 int32 = v2
    v2 += v2
    var v10 float32 = (float32(float64(3.14159)) - float32(v1))
    v11 := []int{(v3 % (((v3 | (v3 >> ((279) & 63)))) | 1)), (v3 << ((251) & 63))}
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float64 = v1
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if (float32(((0.0 * float64(v10)) * v13)) >= v10) {
            v14 := GenS1{F0: 865, F1: 860}
            _ = v14
        } else {
            switch 634 {
            case 500:
                var v15 bool = false
                _ = v15
            case -955:
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 int = -923
                    _ = fi16
                    _ = v17
                }
            case 9223372036854775807:
                v2 = (v5 * (v5 % ((1) | 1)))
            default:
                var v18 int32 = (v2 >> ((255) & 63))
                _ = v18
            }
            switch fi12 {
            case -1:
                for fi19 := 0; fi19 < 14; fi19++ {
                    if true {
                        v20 := GenS2{F0: 0, F1: false}
                        v2 += 394
                        v11 = append(v11, (fi19 * (^(fi19))))
                        _ = v20
                    }
                    _ = fi19
                }
            case -171:
                fmt.Println(v1)
            default:
                switch 1 {
                case 1:
                    v13 = 3.14159
                case -100:
                    var v21 uint32 = 2147483647
                    _ = v21
                }
            }
        }
        _ = fi12
        _ = v13
    }
    var v22 int8 = -99
    switch (("" + "a") + (("foo bar" + "\n\t") + ("foo bar" + ("a" + "foo bar")))) {
    case "foo bar":
        if true {
            for fi23 := 0; fi23 < 15; fi23++ {
                for fi24 := 0; fi24 < 6; fi24++ {
                    if len(v11) > 0 {
                        v11[0] = int(v22)
                    }
                    _ = fi24
                }
                _ = fi23
            }
            v2 += (v2 | 640)
        }
    }
    if true {
        v4 = 1e10
        v25 := GenS2{F0: 0, F1: true}
        v26 := GenS3{F0: 309, F1: "", F2: 580}
        _ = v25
        _ = v26
    }
    if (9223372036854775807 != 0) {
        if false {
            v1 = v1
            v22 += (v22 - v22)
        }
        if false {
            switch "a" {
            case "":
                v3 = genFn0()
            default:
                var v27 uint8 = 255
                _ = v27
            }
        } else {
            v2 -= (v2 ^ v5)
        }
        v22 += 27
    } else {
        if true {
            if (65535 == 65535) {
                v22 += 0
                _ = len(v11)
            }
        } else {
            if (!false) {
                v22 += v22
                if true {
                    if len(v11) > 0 {
                        v11[0] = (v3 & int(v22))
                    }
                } else {
                    v4 = ((float32(-2.25) + float32(v1)) + float32(3.14159))
                }
                var v28 bool = true
                _ = v28
            } else {
                if ((-1 == 9223372036854775807) || false) {
                    fmt.Println(v5)
                } else {
                    v29 := GenS4{F0: (v22 / (((v22 - (v22 / ((127) | 1)))) | 1)), F1: -259, F2: 42}
                    var v30 float32 = ((v4 / v4) + (float32(v1) - v10))
                    _ = v29
                    _ = v30
                }
            }
            v5 += v5
            v31 := GenS5{F0: false}
            _ = v31
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            v4 = v4
            if len(v11) > 0 {
                v11[0] = (v3 | (v3 & (fi32 ^ v3)))
            }
            var v33 float64 = -2.25
            _ = fi32
            _ = v33
        }
    }
    var v34 int16 = 0
    var v35 uint16 = 0
    if true {
        if (!(1 >= 564)) {
            if ((false && (!false)) && (true || (v35 >= 0))) {
                if (float64((((0.1 - v10) / 1.5) + v10)) < v1) {
                    fmt.Println(v1, v5)
                    switch (v3 | (v3 * genFn0())) {
                    case 32768:
                        var v36 bool = false
                        _ = v36
                    case 4294967295:
                        v4 += v10
                    case 743:
                        for fi37 := 0; fi37 < 6; fi37++ {
                            v3 = (fi37 % ((1) | 1))
                            _ = fi37
                        }
                    }
                }
                var v38 uint16 = v35
                v39 := GenS4{F0: (v22 ^ (v22 & (v22 / ((-35) | 1)))), F1: 1, F2: v35}
                _ = v38
                _ = v39
            }
        } else {
            v4 += (v10 * 1e10)
            v40 := GenS6{F0: (v3 | v3), F1: 0}
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 1; fi41++ {
            if true {
                v42 := GenS5{F0: true}
                _ = v42
            } else {
                _ = len(v11)
                v3 -= -468
                var v43 int32 = 128
                _ = v43
            }
            var v44 uint = 18446744073709551615
            fmt.Println(v35, v10, v35)
            _ = fi41
            _ = v44
        }
        for fi45 := 0; fi45 < 1; fi45++ {
            var v46 uint32 = 2147483648
            for fi47 := 0; fi47 < 20; fi47++ {
                if true {
                    switch "" {
                    case "hello":
                        if (95 > 0) {
                            v48 := GenS7{F0: v3, F1: -1, F2: 42}
                            if true {
                                if false {
                                    var v49 int8 = -128
                                    _ = v49
                                }
                            }
                            _ = v48
                        } else {
                            v5 = v5
                            switch "" {
                            case "":
                                v46 += (^(v46))
                            default:
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            }
                        }
                    case "":
                        var v50 uint32 = 42
                        _ = v50
                    default:
                        v5 -= (v5 - (v5 | 4))
                    }
                    var v51 int32 = v2
                    _ = v51
                } else {
                    switch "\n\t" {
                    case "hello":
                        if true {
                            _ = len(v11)
                            v11 = append(v11, fi45)
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                    }
                    v3 += (fi45 / (((v3 % ((int(v2)) | 1))) | 1))
                }
                if (!false) {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v34)
                    default:
                        v34 -= (v34 - -659)
                    }
                    var v52 int = fi45
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v3, v34)
                    case "":
                        v11 = append(v11, genFn0())
                    case "hello":
                        if (-460 != 0) {
                            v52 += v3
                            v53 := GenS3{F0: (v34 + (v34 - (v34 ^ (v34 * (^(v34)))))), F1: "foo bar", F2: v46}
                            v2 = 2147483647
                            _ = v53
                        }
                    }
                    _ = v52
                }
                if true {
                    var v54 int32 = v5
                    _ = v54
                }
                _ = fi47
            }
            var v55 bool = true
            _ = fi45
            _ = v46
            _ = v55
        }
    }
    _ = len(v11)
    switch genFn0() {
    case -332:
        v56 := GenS7{F0: genFn0(), F1: int(v22), F2: -132}
        _ = v56
    default:
        if len(v11) > 0 {
            v11[0] = 4294967295
        }
    }
    fmt.Println(v1)
    if (136 < 36) {
        switch "a" {
        case "hello":
            fmt.Println(v34, v5)
        default:
            v22 = (v22 >> ((18446744073709551615) & 63))
        }
    } else {
        v35 += (v35 + (v35 << ((756) & 63)))
        fmt.Println(v1, v1, v34)
        switch "foo bar" {
        case "a":
            switch ("\n\t" + "hello") {
            case "a":
                for fi57 := 0; fi57 < 8; fi57++ {
                    v1 -= v1
                    var v58 int32 = 2147483647
                    v59 := GenS4{F0: (v22 & (v22 * (v22 / ((-83) | 1)))), F1: 1, F2: (v35 | (v35 ^ 32767))}
                    _ = fi57
                    _ = v58
                    _ = v59
                }
            }
        }
    }
    v1 = v1
    switch "a" {
    case "a":
        v60 := GenS8{F0: v22, F1: v35, F2: v22}
        _ = v60
    case "hello":
        for fi61 := 0; fi61 < 17; fi61++ {
            var v62 bool = true
            _ = fi61
            _ = v62
        }
    }
    var v63 int64 = int64(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v63
}
