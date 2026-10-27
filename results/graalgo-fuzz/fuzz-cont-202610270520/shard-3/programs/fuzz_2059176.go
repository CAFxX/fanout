package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 bool; F1 int32; F2 uint }

func main() {
    var v1 bool = false
    var v2 int32 = 171
    var v3 uint16 = 1
    v4 := make(map[string]int)
    var v5 uint32 = 907
    var v6 string = "foo bar"
    fmt.Println(v1, v1)
    fmt.Println(v6)
    fmt.Println(v5, v6, v1)
    for fi7 := 0; fi7 < 13; fi7++ {
        v1 = ((v1 || v1) || (((!(2 <= 4294967295)) && (1e10 > float32((((float64(0.0) * 0.0) * float64(3.14159)) - (1.5 * 0.1))))) && v1))
        if (int8(v3) > -1) {
            v3 = 65535
        } else {
            if (942 <= fi7) {
                switch fi7 {
                case 32767:
                    if v1 {
                        v8 := GenS1{F0: 42, F1: (v5 >> ((4294967295) & 63)), F2: (v5 % (((v5 | (v5 / ((65535) | 1)))) | 1))}
                        _ = v8
                    }
                case 402:
                    switch (fi7 & (^(fi7))) {
                    case -9223372036854775808:
                        var v9 string = v6
                        _ = v9
                    }
                }
                switch v6 {
                case "hello":
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v6 = ""
                        var v11 uint32 = 934
                        _ = fi10
                        _ = v11
                    }
                case "a":
                    v3 -= (v3 / (((v3 << ((2) & 63))) | 1))
                }
            } else {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS2{F0: (-(v2)), F1: (v2 / (((v2 >> ((65535) & 63))) | 1))}
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    v15 := GenS1{F0: 1000, F1: v5, F2: v5}
                    _ = fi14
                    _ = v15
                }
            }
        }
        fmt.Println(v6)
        _ = fi7
    }
    fmt.Println(v6, v1)
    fmt.Println(v2, v5)
    v16 := GenS2{F0: 42, F1: v2}
    var v17 float32 = float32(0.0)
    v18 := GenS3{F0: 159, F1: -823}
    if false {
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := []int{fi19, 32768, fi19, fi19}
            v5 -= 42
            var v21 int = (fi19 / ((-1) | 1))
            _ = fi19
            _ = v20
            _ = v21
        }
        v22 := v4[(v6 + "\n\t")]
        _ = v22
    }
    if (0 < (v3 | v3)) {
        fmt.Println(v1)
    } else {
        for fi23 := 0; fi23 < 14; fi23++ {
            fmt.Println(fi23, v17)
            for fi24 := 0; fi24 < 8; fi24++ {
                if v1 {
                    if false {
                        var v25 uint64 = 2
                        if (!(int64(v3) <= 502)) {
                            v26 := []int{(^(fi23)), (fi24 * (fi24 | fi24)), int(v3), (fi24 << ((0) & 63))}
                            fmt.Println(v2, v17)
                            _ = v26
                        }
                        v27 := GenS4{F0: 42, F1: v5, F2: v25}
                        var v28 uint = 18446744073709551615
                        var v29 float32 = (v17 / -2.25)
                        _ = v25
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        _ = len(v4)
                    }
                    v6 = "hello"
                    var v30 string = v6
                    _ = v30
                } else {
                    var v31 uint8 = 107
                    if v1 {
                        fmt.Println(v6, fi23, v17)
                        v32 := GenS2{F0: -454, F1: (^(v2))}
                        var v33 int16 = -32768
                        v4[(("foo bar" + "foo bar") + "a")] = (fi24 % ((fi24) | 1))
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(fi23, v5, fi24)
                    v6 = ((v6 + v6) + (v6 + (v6 + (v6 + ("a" + (v6 + v6))))))
                    fmt.Println(fi23, v3)
                    var v34 float64 = 0.1
                    if false {
                        v31 = v31
                        var v35 int = -736
                        var v36 int = (fi24 | fi24)
                        _ = v35
                        _ = v36
                    }
                    _ = v31
                    _ = v34
                }
                _ = fi24
            }
            _ = fi23
        }
    }
    if (1 == 1) {
        v37 := GenS4{F0: -757, F1: 2147483648, F2: 573}
        _ = v37
    }
    if true {
        if (42 == 0) {
            _ = len(v4)
            if v1 {
                if v1 {
                    v38 := GenS5{F0: int64(v5)}
                    fmt.Println(v5)
                    switch -187 {
                    case 1:
                        v39 := []int{32768, 9223372036854775807}
                        _ = v39
                    default:
                        switch (v6 + "a") {
                        case "hello":
                            switch v6 {
                            case "foo bar":
                                for fi40 := 0; fi40 < 12; fi40++ {
                                    for fi41 := 0; fi41 < 17; fi41++ {
                                        v3 = v3
                                        _ = fi41
                                    }
                                    delete(v4, v6)
                                    _ = fi40
                                }
                            case "hello":
                                v1 = (18446744073709551615 < 0)
                            case "":
                                switch 246 {
                                case 32768:
                                    v1 = true
                                case 263:
                                    v6 = "a"
                                case 2147483647:
                                    switch 42 {
                                    case -92:
                                        var v42 int = -9223372036854775808
                                        _ = v42
                                    }
                                default:
                                    if v1 {
                                        var v43 uint8 = 173
                                        _ = v43
                                    }
                                }
                            }
                        case "":
                            if false {
                                fmt.Println(v17, v5)
                            }
                        }
                    }
                    _ = v38
                } else {
                    v17 -= ((0.1 + float32(float64(1e10))) * (v17 / ((v17 + (v17 + v17)) / 1e10)))
                    v17 -= 3.14159
                    fmt.Println(v1, v2, v3)
                }
                v17 = v17
            } else {
                v44 := GenS1{F0: 42, F1: (v5 / ((v5) | 1)), F2: (v5 / ((v5) | 1))}
                _ = v44
            }
            var v45 float64 = -2.25
            _ = v45
        }
        fmt.Println(v3, v3, v5)
        switch 127 {
        case 917:
            v46 := GenS5{F0: 0}
            _ = v46
        }
    }
    switch -536 {
    case -706:
        for fi47 := 0; fi47 < 1; fi47++ {
            for fi48 := 0; fi48 < 20; fi48++ {
                if true {
                    v49 := GenS6{F0: v5, F1: (v3 & 0), F2: -72}
                    v50 := GenS6{F0: v5, F1: (v3 << ((0) & 63)), F2: 1}
                    fmt.Println(fi48)
                    var v51 int8 = 0
                    v52 := GenS7{F0: 42, F1: 42}
                    var v53 int8 = v51
                    _ = v49
                    _ = v50
                    _ = v51
                    _ = v52
                    _ = v53
                }
                var v54 uint = 442
                var v55 uint64 = 18446744073709551615
                _ = fi48
                _ = v54
                _ = v55
            }
            _ = fi47
        }
    default:
        for fi56 := 0; fi56 < 15; fi56++ {
            if ((v2 >> ((314) & 63)) > 42) {
                for fi57 := 0; fi57 < 7; fi57++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v58 int = (-(fi57))
                        _ = v58
                    case "foo bar":
                        v59 := GenS8{F0: v1, F1: int32(v3), F2: 255}
                        _ = v59
                    }
                    _ = fi57
                }
            } else {
                v2 = (^(v2))
                fmt.Println(v5)
            }
            v3 += 42
            v6 = v6
            _ = fi56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
