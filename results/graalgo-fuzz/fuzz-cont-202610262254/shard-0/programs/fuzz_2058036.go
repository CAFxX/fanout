package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 uint32; F2 string }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint8) uint32 {
    v7 := []int{1, -666, 32768}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(p0)
    if false {
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                fmt.Println(fi8, fi8, p1)
                var v10 uint16 = 93
                if ((((1 >= 0) || (!false)) || (false || (true || (808 <= 18446744073709551615)))) || (!false)) {
                    if (1 != 900) {
                        if (869 < 32767) {
                            switch (fi9 & fi8) {
                            case 32767:
                                var v11 uint16 = (^(v10))
                                _ = v11
                            }
                            var v12 float32 = 0.1
                            var v13 int16 = 707
                            _ = v12
                            _ = v13
                        }
                    }
                } else {
                    var v14 float32 = float32(((((1.5 * 3.14159) * (-2.25 - (0.0 * -2.25))) + 1.5) + -2.25))
                    if (-299 >= 100) {
                        if (348 != 4294967295) {
                            var v15 uint16 = (v10 + v10)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v16 := GenS1{F0: (v15 & 764), F1: 799, F2: 4294967295}
                            fmt.Println(p0, v14, p1)
                            var v17 uint8 = p1
                            v10 = 32767
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint = 65535
                            v18 = (^(v18))
                            _ = v18
                        }
                        var v19 uint32 = 2
                        var v20 string = ("hello" + "hello")
                        _ = v19
                        _ = v20
                    }
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
            switch "\n\t" {
            case "a":
                v7 = append(v7, fi8)
            case "\n\t":
                switch "foo bar" {
                case "":
                    var v21 bool = (42 > 0)
                    _ = v21
                case "hello":
                    p1 -= p1
                }
            case "hello":
                _ = len(v7)
            }
            var v22 uint8 = (p1 << ((280) & 63))
            _ = fi8
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            if true {
                switch (fi23 & fi23) {
                case 128:
                    v7 = append(v7, 0)
                case -206:
                    switch (fi23 << ((905) & 63)) {
                    case 1:
                        for fi24 := 0; fi24 < 14; fi24++ {
                            fmt.Println(p1, fi24, p0)
                            v25 := GenS1{F0: 838, F1: 855, F2: 0}
                            p0 = p0
                            p1 = 127
                            switch fi24 {
                            case 0:
                                var v26 bool = (false && true)
                                _ = v26
                            case -2147483648:
                                v27 := GenS2{F0: 4294967295, F1: (p0 * (^(p0))), F2: uint32(p1)}
                                _ = v27
                            case 63:
                                if (((4294967295 < 2) || true) && true) {
                                    var v28 uint8 = (^(p1))
                                    v28 -= v28
                                    var v29 uint32 = 2
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                p0 = p0
                            }
                            v7 = append(v7, (fi23 ^ fi23))
                            _ = fi24
                            _ = v25
                        }
                    case 648:
                        switch (-(fi23)) {
                        case -1:
                            if (false && (-88 != 42)) {
                                v30 := GenS2{F0: 18446744073709551615, F1: 278, F2: 2147483648}
                                _ = v30
                            } else {
                                var v31 int64 = 42
                                _ = v31
                            }
                        case -9223372036854775808:
                            var v32 int = 906
                            _ = v32
                        }
                    }
                }
            } else {
                v33 := GenS2{F0: 4294967295, F1: p0, F2: 156}
                if len(v7) > 0 {
                    v7[0] = (fi23 | (fi23 * -904))
                }
                v34 := GenS3{F0: -9223372036854775808, F1: 239}
                _ = v33
                _ = v34
            }
            v35 := GenS1{F0: 1, F1: 468, F2: 18446744073709551615}
            for fi36 := 0; fi36 < 7; fi36++ {
                var v37 uint32 = 31
                _ = len(v7)
                _ = fi36
                _ = v37
            }
            _ = fi23
            _ = v35
        }
    } else {
        if false {
            p1 -= p1
        } else {
            p0 += 724
            if true {
                p1 = p1
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            } else {
                var v38 int64 = 0
                p1 = (p1 | 128)
                v39 := GenS1{F0: 2, F1: 18446744073709551615, F2: 42}
                _ = v38
                _ = v39
            }
        }
        var v40 string = ("foo bar" + "foo bar")
        v41 := GenS1{F0: 118, F1: 143, F2: 1}
        _ = v40
        _ = v41
    }
    v42 := GenS4{F0: p1}
    _ = p0
    _ = p1
    _ = v42
    _ = v7
    return 0
}

func main() {
    var v1 int = 9223372036854775807
    var v2 string = "\n\t"
    var v3 int8 = -76
    v4 := make(map[int]int)
    var v5 string = "hello"
    var v6 uint = 42
    var v43 int8 = v3
    switch 0 {
    case -306:
        v44 := GenS3{F0: 289, F1: -2147483648}
        _ = v44
    case 2147483647:
        for fi45 := 0; fi45 < 6; fi45++ {
            v3 += v3
            _ = fi45
        }
    case 1:
        var v46 bool = (!false)
        _ = v46
    default:
        if (!(3.14159 > ((1.5 - -2.25) + (3.14159 + (1e10 + (3.14159 - float64((-2.25 - ((1e10 / 0.1) / 3.14159))))))))) {
            v47 := GenS3{F0: -9223372036854775808, F1: int32(v1)}
            _ = v47
        } else {
            switch v5 {
            case "hello":
                if true {
                    v5 = (v5 + (v2 + ((v2 + v5) + ("" + (((v2 + "") + v2) + "a")))))
                    v6 += v6
                    switch 1 {
                    case 32768:
                        v5 = "foo bar"
                    case 4294967295:
                        var v48 uint64 = 404
                        _ = v48
                    case 0:
                        for fi49 := 0; fi49 < 11; fi49++ {
                            v3 = v3
                            v6 = (v6 + 255)
                            _ = fi49
                        }
                    default:
                        v50 := GenS4{F0: 128}
                        _ = v50
                    }
                } else {
                    switch (-(v1)) {
                    case -9223372036854775808:
                        var v51 uint8 = 16
                        _ = v51
                    default:
                        var v52 int = (v1 % ((v1) | 1))
                        _ = v52
                    }
                    v53 := GenS5{F0: 969}
                    v6 -= v6
                    _ = v53
                }
            case "\n\t":
                var v54 uint32 = 541
                _ = v54
            }
            v55 := GenS6{F0: 42, F1: 65535, F2: v2}
            _ = v55
        }
    }
    if false {
        v56 := []int{32767, (v1 >> (((v6 % ((v6) | 1))) & 63)), v1}
        v57 := GenS4{F0: 230}
        _ = v56
        _ = v57
    } else {
        v58 := GenS3{F0: 42, F1: 1}
        _ = v58
    }
    v59 := GenS3{F0: 1, F1: -1}
    for fi60 := 0; fi60 < 10; fi60++ {
        var v61 uint8 = 127
        v62 := []int{}
        if false {
            v4[v1] = v1
            switch (fi60 * 128) {
            case -220:
                switch v2 {
                case "\n\t":
                    for fi63 := 0; fi63 < 7; fi63++ {
                        fmt.Println(v2, v1)
                        _ = fi63
                    }
                case "a":
                    v43 -= (v43 - v43)
                case "hello":
                    v43 -= v43
                }
            case -812:
                fmt.Println(v43)
            case 127:
                v64 := GenS7{F0: 18446744073709551615}
                _ = v64
            }
        }
        _ = fi60
        _ = v61
        _ = v62
    }
    v4[(v1 + (v1 + v1))] = -437
    var v65 bool = true
    if ((((v6 << ((v6) & 63)) >= 65535) || v65) && ((!(v6 == (v6 - v6))) && v65)) {
        v66 := GenS7{F0: 618}
        v2 = v5
        v67 := GenS3{F0: -950, F1: int32(v6)}
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 15; fi68++ {
        v69 := []int{-1, (fi68 ^ fi68), (v1 | fi68)}
        _ = fi68
        _ = v69
    }
    var v70 int16 = 1
    v6 -= 669
    v4[v1] = -2147483648
    var v71 uint64 = 42
    v72 := []int{(^(v1)), (^(v1))}
    var v73 string = v5
    v74 := GenS3{F0: -9223372036854775808, F1: -33}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v59
    _ = v6
    _ = v65
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
}
