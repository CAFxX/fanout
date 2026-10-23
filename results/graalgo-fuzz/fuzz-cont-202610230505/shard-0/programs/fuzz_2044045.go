package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 int32 }

func main() {
    var v1 string = "a"
    var v2 uint64 = 1
    var v3 bool = ((v1 + v1) >= "\n\t")
    v4 := make(map[uint32]bool)
    v5 := []string{v1, "\n\t"}
    var v6 int64 = 9223372036854775807
    if (v3 || (!(836 != 255))) {
        v5 = append(v5, v1)
    } else {
        v7 := GenS1{F0: v3, F1: 2147483648}
        v8 := GenS1{F0: ((v2 << ((18446744073709551615) & 63)) >= (v2 ^ v2)), F1: 2}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: ((2147483647 > -2147483648) && (((v3 || false) && false) || ((2147483647 >= 1) && (!v3)))), F1: 2147483647}
    if (!v3) {
        var v10 float64 = (0.1 / (-2.25 * (3.14159 - 1e10)))
        v11 := GenS2{F0: 453, F1: -9223372036854775808}
        v10 -= 1e10
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 float64 = 1.5
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v3 = v3
        _ = fi14
    }
    var v15 int32 = -2147483648
    v3 = (v3 || (65535 <= 2147483647))
    switch v1 {
    case "foo bar":
        v16 := GenS3{F0: (v1 + (v1 + "")), F1: 150, F2: 1}
        _ = v16
    default:
        v3 = true
    }
    if (3.14159 <= 3.14159) {
        var v17 uint8 = 1
        switch int(v15) {
        case 9223372036854775807:
            var v18 uint8 = 255
            _ = v18
        case 1:
            var v19 uint16 = 1
            _ = v19
        case 4294967295:
            var v20 uint16 = 255
            _ = v20
        }
        if false {
            if v3 {
                v21 := GenS3{F0: v1, F1: 42, F2: 0}
                _ = v21
            } else {
                switch -495 {
                case -65:
                    v1 = "a"
                case -705:
                    if ((v15 * (-(v15))) < 42) {
                        v22 := v4[1]
                        for fi23 := 0; fi23 < 15; fi23++ {
                            for fi24 := 0; fi24 < 19; fi24++ {
                                v25 := v4[820]
                                fmt.Println(v2)
                                _ = fi24
                                _ = v25
                            }
                            if (-18 > 42) {
                                for fi26 := 0; fi26 < 10; fi26++ {
                                    if (2 < 255) {
                                        switch v1 {
                                        case "a":
                                            if len(v5) > 0 {
                                                v5[0] = ("" + v1)
                                            }
                                        default:
                                            var v27 uint8 = 1
                                            var v28 uint64 = (v2 / ((v2) | 1))
                                            _ = v27
                                            _ = v28
                                        }
                                        if v3 {
                                            v6 = v6
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        v2 -= v2
                                    } else {
                                        switch fi23 {
                                        case 9223372036854775807:
                                            v29 := v4[65535]
                                            _ = v29
                                        case 363:
                                            v30 := GenS4{F0: (v2 + v2)}
                                            _ = v30
                                        }
                                        v1 = v1
                                    }
                                    v22 = v3
                                    _ = len(v4)
                                    if false {
                                        v3 = (2147483648 != 990)
                                    } else {
                                        delete(v4, 266)
                                        v22 = v22
                                    }
                                    var v31 uint32 = 0
                                    v5 = append(v5, "")
                                    if true {
                                        v32 := GenS5{F0: (("hello" + v1) + (("a" + ((((v1 + "a") + (v1 + "hello")) + "hello") + (v1 + "\n\t"))) + (v1 + ((("" + "foo bar") + (v1 + ((v1 + ((v1 + v1) + "")) + v1))) + ((v1 + v1) + ((v1 + "a") + "")))))), F1: fi26}
                                        v3 = v22
                                        var v33 int64 = (v6 ^ v6)
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = fi26
                                    _ = v31
                                }
                            }
                            switch (fi23 / ((fi23) | 1)) {
                            case 1:
                                v34 := GenS6{F0: 919, F1: 4294967295}
                                _ = v34
                            default:
                                v22 = (v22 && (v2 < (v2 & (v2 << ((65535) & 63)))))
                            }
                            _ = fi23
                        }
                        _ = v22
                    }
                }
                v35 := GenS7{F0: v6}
                _ = v35
            }
            var v36 bool = ((81 != 255) || false)
            for fi37 := 0; fi37 < 13; fi37++ {
                v6 -= (v6 % (((v6 % ((v6) | 1))) | 1))
                fmt.Println(v3, v3, v6)
                for fi38 := 0; fi38 < 16; fi38++ {
                    v6 = (v6 - v6)
                    var v39 float64 = float64((-2.25 + (1.5 * 0.0)))
                    switch (fi37 % ((9223372036854775807) | 1)) {
                    case -1:
                        switch fi37 {
                        case 691:
                            fmt.Println(v17)
                        case -9223372036854775808:
                            switch 1 {
                            case 1:
                                v40 := v4[42]
                                _ = v40
                            case -124:
                                v39 = (v39 - (v39 * ((((v39 * ((0.1 - v39) + v39)) + v39) - v39) - 0.1)))
                            case 2147483647:
                                switch fi37 {
                                case -378:
                                    v41 := GenS8{F0: 0, F1: (v15 << ((uint(v2)) & 63))}
                                    _ = v41
                                default:
                                    var v42 uint8 = v17
                                    _ = v42
                                }
                            }
                        case 2:
                            fmt.Println(fi37, v3)
                        default:
                            var v43 uint16 = 42
                            _ = v43
                        }
                    case 9223372036854775807:
                        var v44 float64 = v39
                        _ = v44
                    case 128:
                        var v45 float32 = float32((float64(0.0) + v39))
                        _ = v45
                    }
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
            _ = v36
        }
        _ = v17
    }
    var v46 string = (v1 + "hello")
    v47 := GenS6{F0: 1, F1: 55}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v9
}
