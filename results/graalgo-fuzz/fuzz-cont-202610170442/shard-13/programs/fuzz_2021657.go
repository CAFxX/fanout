package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 int8; F1 uint16; F2 string }
type GenS10 struct { F0 int; F1 uint64 }

func main() {
    var v1 int = 592
    var v2 int = v1
    var v3 float64 = (((float64(-2.25) * 1.5) - -2.25) - 1.5)
    v4 := []string{"foo bar", "a"}
    v5 := make(map[int]int)
    var v6 int64 = -1
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v6, v1)
        _ = fi7
    }
    switch 426 {
    case 128:
        var v8 string = "foo bar"
        _ = v8
    default:
        switch -9223372036854775808 {
        case 128:
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        case -275:
            if false {
                if true {
                    v9 := GenS1{F0: 65535, F1: 1}
                    if false {
                        v2 -= v1
                    }
                    v6 = (^(v6))
                    _ = v9
                } else {
                    var v10 uint16 = 744
                    _ = v10
                }
                switch 32768 {
                case -1:
                    v3 = v3
                case 32768:
                    fmt.Println(v1)
                case 4294967295:
                    var v11 float64 = (((v3 - 0.0) / v3) * v3)
                    _ = v11
                }
            } else {
                fmt.Println(v3)
            }
        case -9223372036854775808:
            var v12 float64 = (v3 / v3)
            _ = v12
        default:
            v13 := GenS2{F0: 295, F1: 127}
            _ = v13
        }
    }
    v14 := GenS3{F0: 255, F1: -2147483648}
    fmt.Println(v2)
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v2 = (fi15 / (((^(fi16))) | 1))
            switch (v2 + fi15) {
            case 9223372036854775807:
                var v17 uint8 = 128
                _ = v17
            case -9223372036854775808:
                if (0 <= 100) {
                    v3 = (v3 - float64((0.1 - 0.1)))
                    v18 := v5[(fi15 - (v1 * 128))]
                    _ = v18
                }
            case -2147483648:
                var v19 bool = true
                _ = v19
            default:
                var v20 float64 = v3
                _ = v20
            }
            _ = fi16
        }
        var v21 uint16 = 42
        _ = fi15
        _ = v21
    }
    v6 -= v6
    for fi22 := 0; fi22 < 6; fi22++ {
        switch "\n\t" {
        case "\n\t":
            v23 := GenS3{F0: 42, F1: 42}
            _ = v23
        case "a":
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS4{F0: int32(v2), F1: 18446744073709551615}
                _ = fi24
                _ = v25
            }
        case "":
            for fi26 := 0; fi26 < 4; fi26++ {
                switch v2 {
                case 9223372036854775807:
                    if false {
                        fmt.Println(fi22, v1)
                        _ = len(v5)
                        fmt.Println(v6)
                    } else {
                        if true {
                            v27 := GenS5{F0: 1}
                            fmt.Println(v3)
                            fmt.Println(v6, v1, fi22)
                            v3 = v3
                            _ = v27
                        } else {
                            v6 = v6
                            v6 -= (v6 & (v6 % ((1) | 1)))
                        }
                        switch "" {
                        case "hello":
                            switch 0 {
                            case 2147483647:
                                _ = len(v4)
                                v3 += ((v3 - v3) * (-2.25 + float64(-2.25)))
                                var v28 int64 = v6
                                _ = v28
                            }
                            var v29 string = ""
                            v30 := GenS4{F0: -1, F1: 4294967295}
                            _ = v29
                            _ = v30
                        case "\n\t":
                            if true {
                                switch ("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) {
                                case "foo bar":
                                    var v31 uint16 = 151
                                    _ = v31
                                }
                                var v32 int32 = -289
                                switch v1 {
                                case 127:
                                    if (0 != 255) {
                                        var v33 float64 = (v3 * (float64(0.0) + (v3 + ((((1e10 / ((float64(1e10) / (0.0 - 1e10)) / v3)) / v3) - v3) + float64(1e10)))))
                                        v34 := GenS3{F0: 127, F1: (-(v32))}
                                        v35 := GenS4{F0: (^(v32)), F1: 18446744073709551615}
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                                _ = v32
                            } else {
                                v5[-9223372036854775808] = v2
                                var v36 uint32 = 42
                                _ = v36
                            }
                        case "foo bar":
                            if false {
                                switch "foo bar" {
                                case "\n\t":
                                    v3 += ((v3 / ((v3 + 3.14159) - 1.5)) / 3.14159)
                                }
                            } else {
                                v3 += (v3 - ((float64(((float32(((0.0 / v3) * v3)) + 3.14159) / 1.0)) - (((1e10 + v3) / -2.25) / 1.5)) * v3))
                                v37 := GenS3{F0: 2, F1: 0}
                                var v38 bool = (((("\n\t" + "hello") + "hello") + (("a" + (("hello" + "a") + ("" + ""))) + "a")) <= ((("" + ("\n\t" + ("a" + ("a" + "a")))) + "a") + ""))
                                _ = len(v4)
                                if (float64(0.0) == (float64(-2.25) / 0.1)) {
                                    v39 := GenS6{F0: -960}
                                    if true {
                                        fmt.Println(fi26)
                                    }
                                    _ = v39
                                }
                                if true {
                                    fmt.Println(fi22)
                                    v6 = (v6 + (v6 >> ((66) & 63)))
                                }
                                _ = v37
                                _ = v38
                            }
                        }
                        v40 := GenS7{F0: 65535, F1: 255, F2: 2}
                        _ = v40
                    }
                case 567:
                    v41 := GenS8{F0: 813, F1: -1}
                    _ = v41
                }
                v2 += 4294967295
                switch ("foo bar" + "\n\t") {
                case "a":
                    v42 := v5[296]
                    _ = v42
                case "\n\t":
                    var v43 uint8 = 128
                    _ = v43
                default:
                    if false {
                        fmt.Println(fi22)
                        v2 = v1
                        if (40 > 0) {
                            if (!(!(128 >= 127))) {
                                v5[v2] = (v2 * v1)
                                v2 -= 32767
                                v44 := GenS9{F0: 66, F1: 2, F2: "a"}
                                _ = v44
                            }
                            if false {
                                v2 -= v1
                                _ = len(v4)
                                delete(v5, (v2 % ((v1) | 1)))
                            }
                        } else {
                            v1 += v1
                            if false {
                                v45 := GenS10{F0: -1, F1: 2}
                                var v46 int32 = -1
                                fmt.Println(v2, v3, fi26)
                                v47 := GenS3{F0: 2, F1: (-(v46))}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v48 int = v1
                                _ = v45
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                        }
                        v49 := GenS5{F0: -1}
                        var v50 float32 = -2.25
                        var v51 int32 = 1
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                }
                _ = fi26
            }
        default:
            switch (fi22 << ((18446744073709551615) & 63)) {
            case 449:
                v52 := GenS9{F0: -1, F1: 255, F2: "\n\t"}
                _ = v52
            case 9223372036854775807:
                var v53 int16 = 503
                _ = v53
            case -995:
                var v54 uint16 = 0
                _ = v54
            }
        }
        _ = fi22
    }
    var v55 int8 = 73
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
