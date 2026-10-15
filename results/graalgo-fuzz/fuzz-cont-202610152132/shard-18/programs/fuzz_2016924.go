package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int16; F2 int8 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 float64 = 1e10
    var v2 int16 = -32768
    v3 := []uint32{65535}
    var v4 float32 = (0.0 / (3.14159 + 1e10))
    v5 := GenS1{F0: 42, F1: 127, F2: -1}
    v6 := GenS1{F0: 127, F1: 127, F2: (v2 << ((2) & 63))}
    for fi7 := 0; fi7 < 5; fi7++ {
        switch (fi7 % (((fi7 >> ((42) & 63))) | 1)) {
        case -2147483648:
            switch "\n\t" {
            case "a":
                var v8 bool = true
                _ = v8
            case "\n\t":
                if false {
                    var v9 uint64 = 2
                    v10 := GenS1{F0: 127, F1: 2, F2: v2}
                    _ = v10
                    _ = v9
                }
            }
        case 32768:
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 uint16 = 2
                _ = fi11
                _ = v12
            }
        default:
            var v13 uint = 65535
            _ = v13
        }
        fmt.Println(v4)
        if true {
            fmt.Println(v1, v1, v1)
            for fi14 := 0; fi14 < 13; fi14++ {
                v3 = append(v3, 654)
                var v15 uint = 553
                _ = fi14
                _ = v15
            }
            var v16 uint16 = 553
            _ = v16
        } else {
            fmt.Println(v1)
            v1 = ((v1 * float64(v4)) - 1e10)
        }
        _ = fi7
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if (float64(0.0) != ((1e10 / (0.1 + v1)) + 0.0)) {
        if ((-(v2)) < (v2 >> ((20) & 63))) {
            for fi17 := 0; fi17 < 14; fi17++ {
                if (1 <= 42) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    var v18 int8 = 100
                    _ = v18
                }
                v1 = 3.14159
                _ = fi17
            }
            switch 128 {
            case 127:
                for fi19 := 0; fi19 < 5; fi19++ {
                    v1 = -2.25
                    switch "a" {
                    case "a":
                        if true {
                            if (false || ((^(fi19)) > 42)) {
                                if false {
                                    fmt.Println(fi19, v2)
                                    v20 := GenS2{F0: 126, F1: 255}
                                    if ((!(-1 > 127)) || true) {
                                        fmt.Println(v2, v1)
                                        switch "" {
                                        case "foo bar":
                                            var v21 int8 = 0
                                            _ = v21
                                        case "":
                                            v2 = (v2 | -32768)
                                        case "\n\t":
                                            if false {
                                                v4 -= v4
                                            } else {
                                                var v22 uint64 = 1
                                                _ = v22
                                            }
                                        default:
                                            var v23 uint16 = 2
                                            _ = v23
                                        }
                                    }
                                    _ = v20
                                } else {
                                    v24 := GenS1{F0: 3, F1: 255, F2: (-(v2))}
                                    v4 -= 0.0
                                    _ = v24
                                }
                                if ("\n\t" == ("\n\t" + "hello")) {
                                    switch (fi19 * fi19) {
                                    case 888:
                                        var v25 string = "foo bar"
                                        _ = v25
                                    case -9223372036854775808:
                                        v26 := GenS2{F0: 65535, F1: 165}
                                        _ = v26
                                    }
                                    var v27 uint64 = 784
                                    _ = v27
                                } else {
                                    for fi28 := 0; fi28 < 5; fi28++ {
                                        v2 = 118
                                        _ = fi28
                                    }
                                    v29 := GenS1{F0: -113, F1: 127, F2: -157}
                                    v30 := GenS2{F0: 412, F1: 63}
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                v31 := GenS1{F0: -128, F1: 1, F2: v2}
                                for fi32 := 0; fi32 < 8; fi32++ {
                                    v33 := GenS2{F0: 727, F1: 42}
                                    v2 = (v2 % ((v2) | 1))
                                    fmt.Println(fi32, fi32)
                                    _ = fi32
                                    _ = v33
                                }
                                _ = v31
                            }
                            var v34 string = "\n\t"
                            _ = v34
                        } else {
                            var v35 uint16 = 42
                            _ = v35
                        }
                    default:
                        fmt.Println(v1, v2, v2)
                    }
                    var v36 float64 = 0.1
                    _ = fi19
                    _ = v36
                }
            default:
                v1 = ((0.1 * (v1 - (3.14159 - ((v1 * float64(v4)) - v1)))) * (1.5 / v1))
            }
        }
    }
    var v37 int8 = 0
    v38 := GenS1{F0: v37, F1: 128, F2: (v2 + (v2 & (^(v2))))}
    for fi39 := 0; fi39 < 5; fi39++ {
        v40 := GenS2{F0: 2, F1: 2}
        _ = fi39
        _ = v40
    }
    v37 -= v37
    var v41 int16 = (^(v2))
    switch -1 {
    case 9223372036854775807:
        _ = len(v3)
    case -146:
        for fi42 := 0; fi42 < 15; fi42++ {
            if false {
                fmt.Println(v37)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v37 -= -103
                v37 = v37
            }
            if true {
                v43 := GenS1{F0: 26, F1: 94, F2: (v2 | v2)}
                if (-2.25 < 1.5) {
                    var v44 bool = true
                    _ = v44
                } else {
                    var v45 int32 = -698
                    _ = v45
                }
                _ = v43
            } else {
                var v46 int32 = 2147483647
                for fi47 := 0; fi47 < 14; fi47++ {
                    var v48 uint64 = 2
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v46 += v46
                    fmt.Println(v1, v41, v46)
                    v49 := GenS3{F0: 1, F1: (v41 - 32767), F2: v37}
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                _ = v46
            }
            _ = fi42
        }
    }
    v37 += (v37 ^ v37)
    v50 := GenS4{F0: 255}
    var v51 uint8 = 128
    fmt.Println(v1, v41)
    fmt.Println(v41)
    v52 := GenS4{F0: v51}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
