package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }
type GenS2 struct { F0 bool; F1 int32; F2 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 int8; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := []string{"a", "\n\t", "\n\t", "hello"}
    var v2 string = ("a" + "a")
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    var v5 string = v2
    var v6 uint16 = 32767
    var v7 int64 = 577
    var v8 string = (v2 + (((v2 + v2) + "a") + ""))
    var v9 string = ("foo bar" + v2)
    if false {
        v2 = v9
    }
    fmt.Println(v3, v6)
    v9 = ((v9 + ("foo bar" + v5)) + ((((v9 + v2) + v5) + "\n\t") + (v8 + (v2 + ((v8 + "a") + (v9 + "a"))))))
    var v10 int32 = -2147483648
    v10 = -340
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS1{F0: v2, F1: 170, F2: false}
        var v13 float32 = ((1.5 * 0.0) + (float32(((-2.25 + 1e10) * 0.1)) / float32(-2.25)))
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: true, F1: v10, F2: -214}
    if false {
        v5 = (v5 + v9)
        v7 = (v3 & -1)
        if false {
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 19; fi16++ {
                    v10 = (v10 | (v10 | 42))
                    _ = fi16
                }
                if ((uint64(v6) <= 18446744073709551615) || (0 < 1)) {
                    var v17 int = (-(fi15))
                    v18 := GenS3{F0: 1, F1: (^(v10)), F2: uint(v6)}
                    var v19 uint8 = 2
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi15
            }
        } else {
            if true {
                fmt.Println(v3, v8)
            }
            v20 := GenS1{F0: "hello", F1: 2147483647, F2: (1 == 32767)}
            v21 := GenS4{F0: (v6 / ((218) | 1))}
            _ = v20
            _ = v21
        }
    } else {
        var v22 float32 = -2.25
        if (v2 != ("\n\t" + "\n\t")) {
            v7 -= (v7 + (-(v7)))
            if (!((1 > 238) && false)) {
                v1 = append(v1, v8)
                switch (((("\n\t" + v2) + ("hello" + ("hello" + ("\n\t" + "\n\t")))) + (("foo bar" + v2) + v8)) + (v8 + v2)) {
                case "":
                    fmt.Println(v9)
                default:
                    switch 32768 {
                    case 509:
                        fmt.Println(v2, v7, v3)
                    case 507:
                        fmt.Println(v9, v5, v7)
                    case 32768:
                        _ = len(v4)
                    }
                }
            }
        }
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 uint16 = v6
                _ = len(v1)
                if (-1 != -128) {
                    v26 := GenS4{F0: 42}
                    var v27 uint64 = 846
                    _ = v26
                    _ = v27
                } else {
                    v8 = v2
                    v8 = v9
                    var v28 uint16 = 415
                    _ = v28
                }
                _ = fi24
                _ = v25
            }
            if (2 != 0) {
                if (2 <= 18446744073709551615) {
                    if (true && (128 == 9)) {
                        switch (("hello" + "") + ("hello" + "hello")) {
                        case "a":
                            v7 -= (v7 << ((255) & 63))
                        case "":
                            for fi29 := 0; fi29 < 1; fi29++ {
                                if (false || true) {
                                    _ = len(v1)
                                } else {
                                    v30 := GenS5{F0: 255}
                                    var v31 uint8 = 42
                                    _ = v30
                                    _ = v31
                                }
                                _ = fi29
                            }
                        case "hello":
                            for fi32 := 0; fi32 < 3; fi32++ {
                                v33 := GenS6{F0: "a", F1: 52, F2: v7}
                                var v34 float32 = 0.1
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        }
                        v10 -= 2147483647
                    }
                } else {
                    v4[(-(fi23))] = "hello"
                    v35 := GenS7{F0: 0, F1: (v6 + v6)}
                    _ = v35
                }
                var v36 int64 = (v3 >> ((2) & 63))
                for fi37 := 0; fi37 < 10; fi37++ {
                    if ((false && false) && (((1.5 / (1.5 * float32(1.5))) / (1.5 * ((1.5 + float32(1.5)) / -2.25))) > 0.1)) {
                        v38 := GenS5{F0: 2}
                        var v39 int16 = 42
                        _ = v38
                        _ = v39
                    } else {
                        v40 := GenS6{F0: "a", F1: -1, F2: v3}
                        v41 := v4[(fi23 % ((0) | 1))]
                        v6 -= 1
                        _ = v40
                        _ = v41
                    }
                    v7 = (v3 & (-(v3)))
                    v42 := GenS4{F0: 611}
                    _ = fi37
                    _ = v42
                }
                _ = v36
            }
            _ = fi23
        }
        switch 32768 {
        case 4294967295:
            var v43 int32 = (-(v10))
            _ = v43
        case 32768:
            v7 += v3
        case -252:
            fmt.Println(v6, v6)
        default:
            var v44 uint64 = 1
            _ = v44
        }
    } else {
        v2 = v5
        v7 -= -1
        var v45 int8 = 42
        _ = v45
    }
    if true {
        var v46 int16 = 549
        if true {
            v10 -= (v10 & v10)
            v47 := GenS3{F0: 4294967295, F1: v10, F2: 1}
            _ = v47
        } else {
            switch -124 {
            case 9223372036854775807:
                v3 += (v3 >> ((137) & 63))
            case -531:
                fmt.Println(v8, v46)
            case -450:
                for fi48 := 0; fi48 < 15; fi48++ {
                    fmt.Println(v46, v10)
                    fmt.Println(v5, v7)
                    switch (fi48 << ((4294967295) & 63)) {
                    case 96:
                        var v49 uint = 0
                        _ = v49
                    default:
                        var v50 string = v2
                        _ = v50
                    }
                    _ = fi48
                }
            default:
                var v51 int64 = 42
                _ = v51
            }
        }
        var v52 bool = true
        _ = v46
        _ = v52
    }
    _ = len(v4)
    switch 4294967295 {
    case 966:
        for fi53 := 0; fi53 < 20; fi53++ {
            var v54 int8 = 100
            _ = fi53
            _ = v54
        }
    }
    v2 = (v9 + "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
