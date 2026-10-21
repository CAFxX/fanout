package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int64; F1 uint64; F2 uint }
type GenS9 struct { F0 string; F1 int32; F2 int8 }
type GenS10 struct { F0 int64; F1 string; F2 uint16 }

func main() {
    v1 := []string{("\n\t" + "foo bar"), "hello", "hello"}
    v2 := []int{-126, 0, 437, -2147483648}
    var v3 uint = 2
    var v4 int16 = 32767
    var v5 uint64 = 478
    v4 = v4
    var v6 string = ""
    v6 = v6
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint = (v3 ^ 65535)
        switch v6 {
        case "hello":
            switch fi7 {
            case -2147483648:
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 int64 = 42
                    _ = fi9
                    _ = v10
                }
            case 128:
                _ = len(v2)
            }
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = (fi7 ^ -566)
            }
        case "a":
            if ((^(v4)) == (v4 - v4)) {
                if ((true || (true && true)) && false) {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        if (v3 != (v8 & (v3 >> (((v8 * (v3 << ((v8) & 63)))) & 63)))) {
                            switch (-(fi11)) {
                            case 243:
                                v12 := GenS1{F0: v5, F1: (v5 & 107)}
                                _ = v12
                            case -1:
                                var v13 uint64 = v5
                                _ = v13
                            case 128:
                                if (v6 > (v6 + "")) {
                                    v14 := GenS2{F0: fi7}
                                    fmt.Println(v6, v5, v8)
                                    _ = v14
                                } else {
                                    v15 := GenS3{F0: 1, F1: (v5 ^ 885), F2: 42}
                                    var v16 string = "hello"
                                    v17 := GenS2{F0: 0}
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = (v6 + v6)
                            }
                        } else {
                            var v18 int8 = 0
                            if false {
                                switch "hello" {
                                case "foo bar":
                                    v6 = "foo bar"
                                case "\n\t":
                                    if true {
                                        _ = len(v1)
                                        var v19 uint16 = 596
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = v6
                                            }
                                            var v20 uint16 = v19
                                            _ = v20
                                        } else {
                                            v18 = int8(fi7)
                                        }
                                        _ = v19
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                default:
                                    v6 = v6
                                }
                            } else {
                                if true {
                                    v21 := GenS3{F0: (^(v18)), F1: 1, F2: 602}
                                    _ = v21
                                }
                                var v22 int32 = 1
                                v23 := GenS4{F0: v6}
                                _ = v22
                                _ = v23
                            }
                            var v24 uint16 = 2
                            _ = v18
                            _ = v24
                        }
                        if true {
                            var v25 int64 = 42
                            _ = v25
                        } else {
                            v5 += v5
                        }
                        _ = fi11
                    }
                    var v26 uint = 0
                    v3 = 18446744073709551615
                    _ = v26
                } else {
                    fmt.Println(v3)
                    v8 -= v3
                }
                if (int32(v5) > 1) {
                    var v27 float32 = float32(float64((1.5 - (((1e10 / 1.5) / 1e10) * 1.5))))
                    _ = v27
                }
                v6 = "foo bar"
            } else {
                if false {
                    var v28 int8 = 0
                    v28 = 127
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v30 := GenS4{F0: (v6 + "hello")}
                        _ = fi29
                        _ = v30
                    }
                    v5 += v5
                    v6 = "a"
                }
                v31 := GenS4{F0: (v6 + v6)}
                _ = v31
            }
        default:
            if (1 <= 1) {
                v5 = v5
                if true {
                    if false {
                        var v32 int = (fi7 % ((fi7) | 1))
                        fmt.Println(v4, v5)
                        _ = v32
                    }
                    v2 = append(v2, fi7)
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v1 = append(v1, v6)
                        var v34 int32 = 42
                        v35 := GenS2{F0: (fi33 >> ((65535) & 63))}
                        if true {
                            var v36 uint8 = 255
                            v37 := GenS5{F0: (v36 & v36)}
                            _ = v36
                            _ = v37
                        } else {
                            v5 = v5
                        }
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            }
        }
        _ = fi7
        _ = v8
    }
    if ((v6 + v6) >= ((v6 + (v6 + v6)) + "hello")) {
        v5 = (v5 + v5)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = v5
    } else {
        var v38 uint64 = 2
        var v39 float64 = -2.25
        v40 := GenS6{F0: -2147483648, F1: 0, F2: 2147483648}
        _ = v38
        _ = v39
        _ = v40
    }
    var v41 int32 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        v3 = v3
        for fi43 := 0; fi43 < 15; fi43++ {
            v2 = append(v2, fi42)
            var v44 int8 = 1
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    switch v6 {
    case "":
        if true {
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 float32 = ((0.0 + 3.14159) - 0.1)
                if true {
                    var v47 int16 = v4
                    var v48 string = "\n\t"
                    _ = v47
                    _ = v48
                }
                _ = fi45
                _ = v46
            }
            var v49 int64 = 854
            for fi50 := 0; fi50 < 11; fi50++ {
                _ = len(v2)
                _ = fi50
            }
            _ = v49
        } else {
            if (v4 >= v4) {
                var v51 int = 600
                v52 := GenS7{F0: 168, F1: (v41 | (-(v41)))}
                v53 := GenS7{F0: 128, F1: 0}
                _ = v51
                _ = v52
                _ = v53
            }
            if true {
                v5 -= v5
                v3 -= v3
            }
            var v54 int16 = 1
            _ = v54
        }
    }
    switch -124 {
    case -916:
        for fi55 := 0; fi55 < 10; fi55++ {
            var v56 float32 = 1e10
            var v57 int8 = 104
            _ = fi55
            _ = v56
            _ = v57
        }
    case 2147483647:
        if true {
            var v58 float32 = 0.0
            _ = v58
        }
    case -9223372036854775808:
        for fi59 := 0; fi59 < 17; fi59++ {
            fmt.Println(v6, v4, v4)
            _ = fi59
        }
    }
    var v60 uint64 = 2
    v60 = (v5 + (v60 << ((700) & 63)))
    for fi61 := 0; fi61 < 13; fi61++ {
        if false {
            var v62 int16 = v4
            var v63 int32 = (v41 << ((uint(v60)) & 63))
            _ = v62
            _ = v63
        }
        if (0.1 > 0.1) {
            if ((true && (42 < uint8(v3))) && (((v41 % ((42) | 1)) >= (v41 + (v41 + v41))) || true)) {
                if false {
                    fmt.Println(fi61, v4, v3)
                    for fi64 := 0; fi64 < 6; fi64++ {
                        v5 = 18446744073709551615
                        v3 = (v3 - (^(v3)))
                        v65 := GenS1{F0: v5, F1: (^(v5))}
                        _ = fi64
                        _ = v65
                    }
                    var v66 float64 = 1e10
                    _ = v66
                }
                for fi67 := 0; fi67 < 20; fi67++ {
                    var v68 uint16 = 827
                    var v69 int64 = 42
                    if (((v3 & v3) == v3) || ((0.1 / 1.0) <= (float64((((1e10 + 3.14159) * 0.0) / 3.14159)) + 1.5))) {
                        var v70 int64 = v69
                        var v71 int = -950
                        _ = v70
                        _ = v71
                    } else {
                        v2 = append(v2, (^(fi67)))
                        v72 := GenS8{F0: (v69 & 42), F1: v5, F2: v3}
                        _ = v72
                    }
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                switch v6 {
                case "a":
                    v41 = v41
                case "foo bar":
                    fmt.Println(v5, v60, v60)
                case "\n\t":
                    var v73 uint32 = 856
                    _ = v73
                }
            } else {
                v74 := GenS2{F0: -727}
                v75 := GenS9{F0: ("\n\t" + "\n\t"), F1: -2147483648, F2: int8(v5)}
                _ = v74
                _ = v75
            }
        } else {
            fmt.Println(v6, v3)
        }
        _ = fi61
    }
    if false {
        if false {
            var v76 string = v6
            _ = v76
        } else {
            var v77 float64 = (3.14159 - 0.1)
            _ = v77
        }
        if false {
            v3 = v3
        } else {
            var v78 bool = false
            v78 = (!((387 >= 0) && false))
            for fi79 := 0; fi79 < 8; fi79++ {
                for fi80 := 0; fi80 < 8; fi80++ {
                    fmt.Println(v5)
                    fmt.Println(v60, v41, fi79)
                    _ = fi80
                }
                v81 := GenS10{F0: 9223372036854775807, F1: (v6 + ("" + ("" + (v6 + v6)))), F2: 1}
                _ = fi79
                _ = v81
            }
            _ = v78
        }
        v2 = append(v2, 0)
    } else {
        var v82 int32 = v41
        v3 = (^(v3))
        _ = v82
    }
    v4 -= -32768
    switch 9223372036854775807 {
    case 0:
        v60 -= v5
    case 32767:
        switch 42 {
        case 128:
            if false {
                v60 = (v60 - (v60 - v5))
            }
        default:
            for fi83 := 0; fi83 < 1; fi83++ {
                v41 = (v41 * v41)
                var v84 float64 = 1.5
                _ = fi83
                _ = v84
            }
        }
    }
    var v85 uint32 = 437
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v60
    _ = v85
}
