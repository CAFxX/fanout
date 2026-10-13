package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 bool; F2 uint8 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]string)
    var v3 int16 = 0
    var v4 bool = true
    var v5 uint = 0
    var v6 uint8 = 42
    v4 = true
    v7 := GenS1{F0: v6, F1: 32768, F2: 829}
    v4 = v4
    if (2147483647 <= 2) {
        if (((v5 | (v5 % (((v5 | 18446744073709551615)) | 1))) == v5) || v4) {
            var v8 uint32 = 452
            _ = v8
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v5, v4, v1)
            _ = len(v2)
            switch fi9 {
            case -9223372036854775808:
                for fi10 := 0; fi10 < 3; fi10++ {
                    v4 = v4
                    v11 := []int{(fi9 & (fi10 - (fi9 / ((9223372036854775807) | 1)))), (fi10 / ((-2147483648) | 1))}
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi9
        }
        var v12 uint16 = 32767
        _ = v12
    } else {
        var v13 int = 4294967295
        v14 := []int{(^(v13)), v13}
        var v15 int16 = -32768
        _ = v13
        _ = v14
        _ = v15
    }
    if (42 >= 228) {
        var v16 int = 383
        v2[(v16 - (v16 % (((v16 % ((v16) | 1))) | 1)))] = ""
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := []int{(fi17 ^ (v16 & (fi17 ^ fi17)))}
            v16 += (fi17 / (((fi17 * (v16 * v16))) | 1))
            var v19 bool = v4
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        var v20 float32 = (3.14159 * ((1e10 * 0.1) / 1.0))
        fmt.Println(v5, v1, v3)
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                switch ("foo bar" + "a") {
                case "\n\t":
                    switch "a" {
                    case "hello":
                        v23 := GenS2{F0: 36, F1: 285}
                        _ = v23
                    case "foo bar":
                        switch ("" + "\n\t") {
                        case "hello":
                            v24 := GenS2{F0: 116, F1: -1}
                            _ = v24
                        default:
                            v25 := []int{fi22}
                            _ = v25
                        }
                    case "a":
                        v26 := GenS3{F0: 1, F1: "", F2: 1}
                        _ = v26
                    default:
                        fmt.Println(fi21)
                    }
                }
                v27 := GenS4{F0: 4294967295, F1: -1}
                v4 = true
                var v28 float32 = float32(-2.25)
                switch -9223372036854775808 {
                case 128:
                    var v29 string = "\n\t"
                    _ = v29
                case 511:
                    if v4 {
                        switch "\n\t" {
                        case "a":
                            switch (^(fi21)) {
                            case -2147483648:
                                switch "hello" {
                                case "hello":
                                    switch fi21 {
                                    case -9223372036854775808:
                                        switch "a" {
                                        case "\n\t":
                                            v30 := GenS3{F0: 9223372036854775807, F1: "a", F2: (^(v1))}
                                            _ = v30
                                        }
                                    default:
                                        fmt.Println(v1, v1, v20)
                                    }
                                case "a":
                                    switch fi21 {
                                    case -1:
                                        var v31 int64 = 9223372036854775807
                                        _ = v31
                                    }
                                case "\n\t":
                                    var v32 uint64 = 4294967295
                                    _ = v32
                                default:
                                    var v33 float64 = 3.14159
                                    _ = v33
                                }
                            }
                        default:
                            var v34 float64 = 1.5
                            _ = v34
                        }
                        v3 -= v3
                    }
                case 9223372036854775807:
                    delete(v2, 2147483647)
                }
                fmt.Println(fi21)
                if (v6 < v6) {
                    v6 = 0
                    fmt.Println(v4, fi21, fi22)
                } else {
                    v35 := []int{(fi21 | 32767), fi21, 42, (fi21 + (^(fi22)))}
                    v20 = v28
                    _ = v35
                }
                _ = fi22
                _ = v27
                _ = v28
            }
            _ = fi21
        }
        _ = v20
    }
    v36 := GenS5{F0: (^(v6)), F1: (^(v1))}
    for fi37 := 0; fi37 < 6; fi37++ {
        if (("\n\t" + ("foo bar" + (("foo bar" + "hello") + ""))) < "\n\t") {
            var v38 float32 = 1e10
            _ = v38
        } else {
            switch "hello" {
            case "foo bar":
                v4 = false
            case "hello":
                for fi39 := 0; fi39 < 11; fi39++ {
                    fmt.Println(v3)
                    v40 := []int{-9223372036854775808, (fi39 >> ((0) & 63)), fi39, fi37}
                    _ = fi39
                    _ = v40
                }
            }
            if (!v4) {
                var v41 uint16 = v1
                v1 = v1
                if (!((v3 | v3) <= (v3 | v3))) {
                    for fi42 := 0; fi42 < 12; fi42++ {
                        v43 := GenS6{F0: -2147483648}
                        v6 = (v6 ^ 166)
                        fmt.Println(v5, v41, v1)
                        _ = fi42
                        _ = v43
                    }
                    v6 = (v6 / ((85) | 1))
                }
                _ = v41
            }
        }
        _ = fi37
    }
    switch -122 {
    case 949:
        fmt.Println(v3, v5, v4)
    default:
        if (v4 && (0 != 2147483647)) {
            fmt.Println(v6, v1)
            if (!((v4 || v4) && ((v6 << (((v5 / ((4294967295) | 1))) & 63)) != v6))) {
                v6 -= 0
            }
        } else {
            for fi44 := 0; fi44 < 19; fi44++ {
                var v45 uint32 = uint32(v5)
                _ = fi44
                _ = v45
            }
            v6 -= v6
        }
    }
    if v4 {
        v46 := []int{-16, 2147483647}
        if (!true) {
            v47 := v2[-362]
            fmt.Println(v3)
            v1 = (v1 & (v1 + 863))
            _ = v47
        } else {
            v6 -= 128
            var v48 int16 = (v3 * (v3 % ((138) | 1)))
            _ = v48
        }
        _ = v46
    } else {
        for fi49 := 0; fi49 < 9; fi49++ {
            for fi50 := 0; fi50 < 8; fi50++ {
                v51 := []int{-248}
                if (((("\n\t" + "foo bar") + "") <= "foo bar") && v4) {
                    switch fi49 {
                    case -2147483648:
                        var v52 int8 = -103
                        _ = v52
                    case 42:
                        v53 := GenS4{F0: 93, F1: -54}
                        _ = v53
                    default:
                        if len(v51) > 0 {
                            v51[0] = (fi50 & (fi49 + (^(fi50))))
                        }
                    }
                    v3 = (v3 | v3)
                }
                fmt.Println(fi49, fi50, v3)
                _ = fi50
                _ = v51
            }
            _ = fi49
        }
    }
    v1 += (v1 * v1)
    v4 = true
    var v54 int = -1
    var v55 int32 = 1
    for fi56 := 0; fi56 < 20; fi56++ {
        switch (fi56 + fi56) {
        case 528:
            fmt.Println(v54, v6, fi56)
        case -2147483648:
            v55 += v55
        case 32768:
            v5 -= 421
        default:
            v4 = v4
        }
        _ = fi56
    }
    switch (v54 >> ((v5) & 63)) {
    case 32767:
        var v57 string = "\n\t"
        _ = v57
    case 128:
        v58 := GenS7{F0: 296, F1: true, F2: (v6 - (v6 % ((v6) | 1)))}
        _ = v58
    case 306:
        var v59 int = 128
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
