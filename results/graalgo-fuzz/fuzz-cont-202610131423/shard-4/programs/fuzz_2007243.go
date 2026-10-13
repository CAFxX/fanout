package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 string; F1 uint }
type GenS10 struct { F0 bool; F1 int }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]int)
    var v3 float32 = 0.1
    v4 := []int{906, 0, 128, -58}
    var v5 string = "a"
    v5 = ("" + (v5 + ""))
    if ((0.1 - (3.14159 + v3)) != v3) {
        v4 = append(v4, -303)
        v3 -= (v3 - 0.1)
        if true {
            v4 = append(v4, 128)
            if false {
                for fi6 := 0; fi6 < 6; fi6++ {
                    for fi7 := 0; fi7 < 12; fi7++ {
                        if ((^(fi6)) <= 128) {
                            v8 := GenS1{F0: v5, F1: (-1 <= -128)}
                            var v9 bool = (42 >= 2)
                            v10 := GenS2{F0: fi6}
                            _ = v10
                            _ = v8
                            _ = v9
                        }
                        _ = fi7
                    }
                    _ = fi6
                }
                for fi11 := 0; fi11 < 3; fi11++ {
                    switch v5 {
                    case "a":
                        if (0 <= 0) {
                            if ((-9223372036854775808 > 1) && (!(65535 > 4294967295))) {
                                v1 -= ((v1 * v1) * v1)
                                if (1 == 18446744073709551615) {
                                    if true {
                                        for fi12 := 0; fi12 < 4; fi12++ {
                                            v13 := GenS3{F0: 185, F1: -9223372036854775808, F2: -1}
                                            v14 := GenS4{F0: -1}
                                            switch -9223372036854775808 {
                                            case 9223372036854775807:
                                                if len(v4) > 0 {
                                                    v4[0] = 2147483647
                                                }
                                            }
                                            _ = fi12
                                            _ = v13
                                            _ = v14
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            v3 = (((v3 - ((1.5 / v3) * ((1.5 + (v3 * 1e10)) + -2.25))) / v3) + ((float32(-2.25) * 1.5) - 1.5))
                                        case "hello":
                                            switch fi11 {
                                            case 32768:
                                                v15 := GenS2{F0: fi11}
                                                _ = v15
                                            case -2147483648:
                                                for fi16 := 0; fi16 < 18; fi16++ {
                                                    v4 = append(v4, fi11)
                                                    var v17 int16 = -30
                                                    v1 -= ((float64(v3) - 1.5) + (1e10 - v1))
                                                    _ = fi16
                                                    _ = v17
                                                }
                                            }
                                        default:
                                            v5 = v5
                                        }
                                    } else {
                                        v4 = append(v4, (^(fi11)))
                                    }
                                } else {
                                    if (-2147483648 > 439) {
                                        var v18 string = v5
                                        _ = v18
                                    } else {
                                        v19 := GenS3{F0: 0, F1: 0, F2: 42}
                                        v3 = ((1e10 / v3) * float32(v1))
                                        _ = v19
                                    }
                                    var v20 float64 = float64(3.14159)
                                    _ = v20
                                }
                                if (2 == 42) {
                                    for fi21 := 0; fi21 < 20; fi21++ {
                                        v3 = (v3 + 0.1)
                                        _ = fi21
                                    }
                                    switch fi11 {
                                    case -187:
                                        var v22 uint64 = 1
                                        _ = v22
                                    default:
                                        v23 := v2[(v5 + "hello")]
                                        _ = v23
                                    }
                                    fmt.Println(v3)
                                }
                            }
                            var v24 string = (v5 + ("foo bar" + "\n\t"))
                            switch fi11 {
                            case -2147483648:
                                if ((("foo bar" + v5) <= (v24 + ("foo bar" + v5))) && false) {
                                    v1 += v1
                                    var v25 uint16 = 0
                                    _ = v25
                                } else {
                                    v1 += -2.25
                                    v2["\n\t"] = (^(fi11))
                                }
                            case 0:
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    var v27 uint = 204
                                    switch fi26 {
                                    case -801:
                                        var v28 bool = true
                                        _ = v28
                                    case -968:
                                        var v29 bool = (v27 == (v27 << (((^(v27))) & 63)))
                                        _ = v29
                                    case 2147483647:
                                        var v30 string = (v5 + v5)
                                        _ = v30
                                    default:
                                        var v31 uint64 = 460
                                        v32 := GenS5{F0: 625}
                                        var v33 uint16 = 2
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    v3 = (float32(0.0) - v3)
                                    v5 = v5
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            _ = v24
                        } else {
                            v5 = ""
                        }
                    case "\n\t":
                        v4 = append(v4, (fi11 >> ((112) & 63)))
                    }
                    v34 := GenS6{F0: 610}
                    _ = fi11
                    _ = v34
                }
                v1 = v1
            }
        } else {
            var v35 int16 = 479
            delete(v2, "a")
            var v36 uint32 = 591
            _ = v35
            _ = v36
        }
    }
    v1 += v1
    if (v5 <= v5) {
        if len(v4) > 0 {
            v4[0] = 145
        }
        var v37 float64 = ((v1 + ((v1 / 3.14159) - v1)) / 1.0)
        _ = v37
    }
    var v38 int8 = 0
    v39 := GenS7{F0: -916}
    switch 32768 {
    case 32767:
        var v40 uint = 1
        _ = v40
    case 127:
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS8{F0: 18446744073709551615, F1: -1}
            for fi43 := 0; fi43 < 17; fi43++ {
                v4 = append(v4, fi43)
                if (!(!true)) {
                    fmt.Println(v1, fi41)
                    switch "foo bar" {
                    case "a":
                        switch v5 {
                        case "":
                            fmt.Println(fi41, fi41, v5)
                        case "\n\t":
                            v44 := v2["foo bar"]
                            _ = v44
                        }
                    case "hello":
                        var v45 int = (fi43 * fi43)
                        _ = v45
                    case "":
                        _ = len(v4)
                    }
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if (-1 <= -2147483648) {
                        switch 390 {
                        case -2147483648:
                            var v46 float32 = (v3 - 3.14159)
                            _ = v46
                        case 179:
                            v3 = ((1.5 / 1e10) / (((v3 - 1.5) / -2.25) + (v3 - v3)))
                        case 9223372036854775807:
                            fmt.Println(v1, fi41, v3)
                        default:
                            fmt.Println(v1, fi43)
                        }
                        v47 := GenS9{F0: v5, F1: 42}
                        _ = v47
                    } else {
                        v48 := GenS9{F0: v5, F1: 1}
                        var v49 uint32 = 452
                        _ = v48
                        _ = v49
                    }
                }
                var v50 uint16 = 65535
                v50 -= (v50 % (((v50 >> ((664) & 63))) | 1))
                _ = fi43
                _ = v50
            }
            _ = fi41
            _ = v42
        }
    case 9223372036854775807:
        v51 := GenS10{F0: ((610 == 1) && ((32767 != 0) && (-402 < -9223372036854775808))), F1: 128}
        _ = v51
    default:
        for fi52 := 0; fi52 < 11; fi52++ {
            switch 958 {
            case 611:
                fmt.Println(v5, v5)
            case -494:
                var v53 int8 = -44
                _ = v53
            default:
                v1 = v1
            }
            _ = fi52
        }
    }
    v5 = "hello"
    var v54 uint16 = 255
    v54 -= v54
    var v55 int16 = 32767
    v5 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v54
    _ = v55
}
