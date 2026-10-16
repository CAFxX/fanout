package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 int8; F2 int32 }
type GenS6 struct { F0 int32; F1 uint32 }

func main() {
    var v1 bool = (42 <= 1)
    var v2 int = -2147483648
    var v3 uint32 = 916
    var v4 int16 = -1
    var v5 int16 = (v4 << ((1) & 63))
    if (v1 || (!("\n\t" != ("hello" + "foo bar")))) {
        v6 := []int{42, v2}
        v5 -= -1
        for fi7 := 0; fi7 < 11; fi7++ {
            if len(v6) > 0 {
                v6[0] = (-(fi7))
            }
            fmt.Println(v2, v1, v5)
            v2 -= (fi7 << ((42) & 63))
            _ = fi7
        }
        _ = v6
    } else {
        var v8 float64 = -2.25
        _ = v8
    }
    switch "" {
    case "":
        var v9 uint8 = 0
        _ = v9
    }
    if (598 != -2147483648) {
        var v10 int16 = (v5 ^ v4)
        var v11 uint = 42
        _ = v10
        _ = v11
    }
    if v1 {
        if true {
            if true {
                v1 = false
                v5 = -801
            } else {
                v12 := GenS1{F0: 0}
                var v13 string = ("hello" + "\n\t")
                _ = v12
                _ = v13
            }
            if v1 {
                v14 := []int{227, (v2 % (((v2 ^ -9223372036854775808)) | 1))}
                _ = v14
            }
        } else {
            v3 -= v3
            if v1 {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if true {
                        switch ("\n\t" + "") {
                        case "":
                            if false {
                                v16 := GenS1{F0: (^(v3))}
                                v3 -= 534
                                fmt.Println(v3, v4)
                                _ = v16
                            }
                        case "hello":
                            v17 := GenS2{F0: v5, F1: -204}
                            _ = v17
                        case "\n\t":
                            v1 = v1
                        default:
                            v4 -= -32768
                        }
                        v4 = -791
                    }
                    fmt.Println(v2, v3)
                    _ = fi15
                }
                v18 := GenS2{F0: v4, F1: -1}
                if v1 {
                    var v19 float64 = (-2.25 - (3.14159 / (3.14159 + (((1e10 / 1.5) * 3.14159) * ((0.1 / 0.1) / 0.1)))))
                    if (false || ("\n\t" <= "")) {
                        var v20 int16 = v4
                        v1 = (true || (42 == 2))
                        _ = v20
                    } else {
                        if (255 > 127) {
                            v19 += (((v19 + 3.14159) / (v19 * v19)) * -2.25)
                        } else {
                            switch (v2 >> ((420) & 63)) {
                            case 1:
                                var v21 uint16 = 2
                                _ = v21
                            case -165:
                                v22 := GenS2{F0: 905, F1: -769}
                                _ = v22
                            case 2147483647:
                                v23 := GenS2{F0: (v5 * 1), F1: 2147483647}
                                _ = v23
                            default:
                                v24 := GenS3{F0: 139}
                                _ = v24
                            }
                        }
                    }
                    fmt.Println(v1, v4)
                    _ = v19
                } else {
                    var v25 uint64 = 1
                    if v1 {
                        if (!(false || true)) {
                            var v26 uint = 2
                            for fi27 := 0; fi27 < 6; fi27++ {
                                v28 := GenS4{F0: v25}
                                v1 = v1
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    if (2 <= v3) {
                                        v4 = (v4 >> ((v26) & 63))
                                        v30 := GenS2{F0: (v4 % (((v4 << ((v26) & 63))) | 1)), F1: 1}
                                        _ = v30
                                    }
                                    v31 := GenS1{F0: 421}
                                    _ = fi29
                                    _ = v31
                                }
                                _ = fi27
                                _ = v28
                            }
                            if (1 > v4) {
                                if (uint8(v26) <= 255) {
                                    var v32 int8 = 1
                                    var v33 uint = (v26 | v26)
                                    _ = v32
                                    _ = v33
                                }
                            } else {
                                var v34 uint8 = 180
                                _ = v34
                            }
                            _ = v26
                        } else {
                            v35 := []int{4294967295, (v2 + v2), (v2 / ((2147483647) | 1))}
                            _ = v35
                        }
                    }
                    if (-9223372036854775808 == 42) {
                        for fi36 := 0; fi36 < 11; fi36++ {
                            fmt.Println(fi36, v25, v4)
                            v37 := GenS5{F0: (!v1), F1: 100, F2: 1}
                            v38 := GenS5{F0: (((float64(-2.25) > (3.14159 + ((3.14159 - (1e10 + 1e10)) / 3.14159))) || (v1 || v1)) && (!v1)), F1: -1, F2: -1}
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        var v39 bool = (v2 != 32768)
                        var v40 float64 = 0.1
                        _ = v39
                        _ = v40
                    }
                    _ = v25
                }
                _ = v18
            }
        }
        var v41 int32 = 1
        v5 = -1
        _ = v41
    }
    var v42 uint = 255
    var v43 uint64 = uint64(v3)
    if (!true) {
        if v1 {
            v4 = v4
            for fi44 := 0; fi44 < 9; fi44++ {
                var v45 uint16 = 1
                for fi46 := 0; fi46 < 4; fi46++ {
                    v2 -= v2
                    _ = fi46
                }
                fmt.Println(v42, v42)
                _ = fi44
                _ = v45
            }
            var v47 uint64 = (v43 % ((42) | 1))
            _ = v47
        }
        for fi48 := 0; fi48 < 18; fi48++ {
            v49 := GenS1{F0: 237}
            if false {
                if (1.5 == 0.1) {
                    var v50 float32 = (-2.25 / 1.5)
                    for fi51 := 0; fi51 < 16; fi51++ {
                        if v1 {
                            var v52 uint8 = 42
                            _ = v52
                        } else {
                            if true {
                                if (!(v1 || v1)) {
                                    v53 := []int{fi48}
                                    _ = v53
                                } else {
                                    var v54 uint = 42
                                    _ = v54
                                }
                                v2 -= (v2 | 32768)
                                if true {
                                    switch fi48 {
                                    case -187:
                                        var v55 float64 = -2.25
                                        _ = v55
                                    default:
                                        fmt.Println(v2, v2)
                                    }
                                }
                            }
                        }
                        switch fi48 {
                        case -163:
                            v56 := GenS5{F0: v1, F1: 98, F2: -1}
                            _ = v56
                        default:
                            if (v43 > (v43 % ((915) | 1))) {
                                var v57 uint64 = (v43 * v43)
                                v58 := GenS3{F0: 1}
                                v59 := []int{v2, fi51, 35}
                                var v60 float64 = -2.25
                                _ = v57
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                fmt.Println(v2, v50)
                                var v61 int16 = (v5 + v4)
                                _ = v61
                            }
                        }
                        var v62 float64 = 1e10
                        var v63 bool = v1
                        _ = fi51
                        _ = v62
                        _ = v63
                    }
                    _ = v50
                }
                switch ((("" + "a") + "foo bar") + ("\n\t" + "\n\t")) {
                case "foo bar":
                    switch "" {
                    case "hello":
                        v3 += (^(v3))
                    }
                default:
                    var v64 int32 = 42
                    _ = v64
                }
            } else {
                v43 -= 429
                for fi65 := 0; fi65 < 5; fi65++ {
                    var v66 int16 = (v5 - (v5 ^ 1))
                    _ = fi65
                    _ = v66
                }
            }
            switch (fi48 * fi48) {
            case -538:
                v67 := []int{9223372036854775807, (^(v2)), v2, (-(v2))}
                _ = v67
            default:
                v68 := GenS3{F0: 887}
                _ = v68
            }
            _ = fi48
            _ = v49
        }
        v43 = (^(v43))
    }
    v42 += v42
    switch 1 {
    case 127:
        v69 := GenS6{F0: 1, F1: v3}
        _ = v69
    case -2147483648:
        if (!(((-9223372036854775808 > -9223372036854775808) || (-1 >= -69)) || true)) {
            fmt.Println(v2, v42)
        }
    }
    v4 += (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))
    if v1 {
        var v70 int64 = 0
        var v71 string = ("" + ("hello" + "hello"))
        v72 := GenS3{F0: 255}
        _ = v70
        _ = v71
        _ = v72
    }
    v73 := []int{(^(v2)), v2, v2, (v2 / (((v2 % (((-(v2))) | 1))) | 1))}
    v3 = (v3 & v3)
    var v74 uint = (v42 * v42)
    for fi75 := 0; fi75 < 1; fi75++ {
        v5 += v4
        _ = fi75
    }
    fmt.Println(v74, v4, v5)
    for fi76 := 0; fi76 < 5; fi76++ {
        for fi77 := 0; fi77 < 14; fi77++ {
            fmt.Println(v5, v74, v43)
            v42 = (^(v74))
            _ = fi77
        }
        _ = fi76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v73
    _ = v74
}
