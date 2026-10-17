package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 string }
type GenS6 struct { F0 string; F1 int }
func genFn0(p0 uint64) uint {
    p0 = 614
    p0 -= (p0 + 18446744073709551615)
    fmt.Println(p0)
    _ = p0
    return 385
}

func main() {
    var v1 int16 = 0
    var v2 int16 = -32768
    var v3 float64 = 0.1
    var v4 int32 = -2147483648
    v5 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 42}
    var v6 int8 = 42
    for fi7 := 0; fi7 < 15; fi7++ {
        switch (fi7 / ((32767) | 1)) {
        case -9223372036854775808:
            var v8 float32 = 1e10
            _ = v8
        default:
            v9 := GenS1{F0: 18446744073709551615, F1: -1, F2: 1}
            _ = v9
        }
        v10 := []int{}
        _ = fi7
        _ = v10
    }
    if false {
        switch ("\n\t" + (("hello" + "hello") + "foo bar")) {
        case "a":
            for fi11 := 0; fi11 < 16; fi11++ {
                v3 = v3
                var v12 int16 = v2
                var v13 uint = 42
                _ = fi11
                _ = v12
                _ = v13
            }
        case "foo bar":
            v14 := []int{-2}
            _ = v14
        case "hello":
            v4 = (v4 + (v4 << ((567) & 63)))
        default:
            for fi15 := 0; fi15 < 1; fi15++ {
                switch fi15 {
                case 32768:
                    v4 += (^(v4))
                case 42:
                    switch "\n\t" {
                    case "foo bar":
                        for fi16 := 0; fi16 < 18; fi16++ {
                            if false {
                                if (42 > 207) {
                                    var v17 int32 = v4
                                    switch "a" {
                                    case "\n\t":
                                        var v18 uint16 = 2
                                        _ = v18
                                    case "hello":
                                        if false {
                                            var v19 uint32 = 4294967295
                                            if (false || (float64(1.5) <= (v3 + v3))) {
                                                v6 -= v6
                                            } else {
                                                v20 := GenS1{F0: 394, F1: -139, F2: 2}
                                                _ = v20
                                            }
                                            _ = v19
                                        } else {
                                            var v21 float64 = (float64(1e10) / ((0.0 - -2.25) * (v3 / v3)))
                                            var v22 float64 = 0.1
                                            _ = v21
                                            _ = v22
                                        }
                                    default:
                                        v4 = (-(v17))
                                        if ((v3 < v3) && (!(true || (4294967295 == 962)))) {
                                            v23 := []int{(fi16 ^ fi15)}
                                            _ = v23
                                        } else {
                                            v24 := []int{(fi16 >> ((genFn0(2)) & 63)), 669, (fi16 / (((^(fi16))) | 1))}
                                            if false {
                                                var v25 int = 127
                                                switch 4294967295 {
                                                case -896:
                                                    var v26 uint16 = 65535
                                                    _ = v26
                                                case 434:
                                                    v25 += -2147483648
                                                default:
                                                    if (!((v2 >> ((563) & 63)) == v1)) {
                                                        if false {
                                                            v25 = v25
                                                            switch ("hello" + "") {
                                                            case "\n\t":
                                                                fmt.Println(v25)
                                                            case "":
                                                                v4 = v17
                                                            }
                                                        }
                                                        var v27 int16 = v1
                                                        if (float64(((float32(v3) * (1.5 * 3.14159)) + 1.5)) >= ((v3 + (1e10 - (0.0 - (float64(1e10) - 1.5)))) - 0.0)) {
                                                            v3 = float64((((-2.25 / -2.25) - (0.0 + (-2.25 * (((0.0 / (3.14159 + (1.5 * 3.14159))) / (1e10 + (-2.25 + 1e10))) / float32(v3))))) / ((1e10 + ((((-2.25 / 1.5) + 1.5) - 0.0) - float32(-2.25))) / float32(3.14159))))
                                                            v28 := GenS1{F0: 4294967295, F1: 42, F2: 255}
                                                            var v29 uint64 = 879
                                                            v6 = v6
                                                            _ = v28
                                                            _ = v29
                                                        } else {
                                                            v4 += (v17 ^ 1)
                                                            v25 -= fi16
                                                            v30 := GenS2{F0: (v4 << ((133) & 63)), F1: 42, F2: (v6 % ((v6) | 1))}
                                                            if true {
                                                                v4 = 2147483647
                                                            }
                                                            var v31 int32 = v4
                                                            if ((18446744073709551615 == 767) || (!((382 < (v17 - (v4 % ((v4) | 1)))) && true))) {
                                                                fmt.Println(v1, v4, v1)
                                                            } else {
                                                                switch "\n\t" {
                                                                case "":
                                                                    if true {
                                                                        var v32 uint = 203
                                                                        if (false && (!(((0 == 65535) && true) || (!true)))) {
                                                                            var v33 uint16 = 0
                                                                            v34 := GenS2{F0: (v31 >> ((v32) & 63)), F1: 215, F2: (v6 % (((v6 - -1)) | 1))}
                                                                            v17 += v31
                                                                            _ = v33
                                                                            _ = v34
                                                                        }
                                                                        var v35 uint = (v32 - (v32 | (v32 - v32)))
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v36 int32 = v17
                                                                            _ = v36
                                                                        case "a":
                                                                            v6 -= v6
                                                                        case "":
                                                                            var v37 bool = (1.5 == -2.25)
                                                                            _ = v37
                                                                        }
                                                                        switch (^(fi16)) {
                                                                        case 127:
                                                                            v17 += -1
                                                                            var v38 float32 = 0.0
                                                                            _ = v38
                                                                        case 32767:
                                                                            v24 = append(v24, fi16)
                                                                        case 42:
                                                                            if len(v24) >= 2 {
                                                                                v24 = v24[1:len(v24)]
                                                                            }
                                                                        }
                                                                        _ = v32
                                                                        _ = v35
                                                                    } else {
                                                                        v17 = (v17 / ((-498) | 1))
                                                                        if true {
                                                                            v39 := GenS3{F0: 236}
                                                                            var v40 uint = 0
                                                                            _ = v39
                                                                            _ = v40
                                                                        }
                                                                        v41 := GenS4{F0: (v6 * (v6 << ((730) & 63))), F1: v2}
                                                                        _ = v41
                                                                    }
                                                                case "a":
                                                                    var v42 string = "\n\t"
                                                                    _ = v42
                                                                case "foo bar":
                                                                    switch (v25 << ((661) & 63)) {
                                                                    case -2147483648:
                                                                        if (1 == 42) {
                                                                            v17 = (v31 - v31)
                                                                            fmt.Println(v27, v3, v17)
                                                                        } else {
                                                                            v27 += (v1 & (v27 % ((v27) | 1)))
                                                                        }
                                                                        v17 -= 2147483647
                                                                        fmt.Println(v2, v4, v27)
                                                                    default:
                                                                        v43 := GenS4{F0: 1, F1: 0}
                                                                        _ = v43
                                                                    }
                                                                    _ = len(v24)
                                                                    var v44 float32 = ((1e10 * (((3.14159 * (1e10 + 0.0)) - ((-2.25 - float32((v3 / (((v3 + (1.5 * v3)) * float64((0.1 * -2.25))) + v3)))) - ((0.0 - (1e10 / (float32(0.1) / 1.0))) - 1.5))) - -2.25)) / 3.14159)
                                                                    _ = v44
                                                                default:
                                                                    switch (v25 * v25) {
                                                                    case 4294967295:
                                                                        var v45 int64 = -9223372036854775808
                                                                        v1 -= (v2 / ((int16(fi15)) | 1))
                                                                        v17 = 30
                                                                        _ = v45
                                                                    case -379:
                                                                        var v46 int8 = (v6 ^ (^(v6)))
                                                                        _ = v46
                                                                    case 973:
                                                                        _ = len(v24)
                                                                    default:
                                                                        fmt.Println(v31, v25)
                                                                    }
                                                                }
                                                                fmt.Println(v27, fi16, v3)
                                                                v4 = -518
                                                            }
                                                            if len(v24) > 0 {
                                                                v24[0] = v25
                                                            }
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        var v47 uint16 = 1
                                                        v27 -= -32768
                                                        _ = v27
                                                        _ = v47
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v48 int16 = v1
                                                    _ = v48
                                                case "":
                                                    fmt.Println(v17, v6)
                                                case "foo bar":
                                                    switch fi16 {
                                                    case 1:
                                                        v6 = v6
                                                    case 42:
                                                        v3 = (float64(0.0) * ((-2.25 - float64((((float32(v3) * 0.1) - (float32(v3) / 0.1)) - (3.14159 - (0.0 * 0.1))))) * ((0.1 - (v3 / (v3 / (v3 - float64(1.5))))) + (v3 - float64(1.5)))))
                                                    default:
                                                        v49 := GenS3{F0: 426}
                                                        _ = v49
                                                    }
                                                default:
                                                    if len(v24) > 0 {
                                                        v24[0] = v25
                                                    }
                                                }
                                                _ = v25
                                            } else {
                                                var v50 int8 = v6
                                                v24 = append(v24, fi15)
                                                fmt.Println(v3)
                                                v6 = 42
                                                _ = v50
                                            }
                                            _ = v24
                                        }
                                    }
                                    if true {
                                        var v51 uint32 = 2147483648
                                        if (!(!(((!(!((v3 <= float64(-2.25)) && (9223372036854775807 < -498)))) && (!("\n\t" <= ""))) || false))) {
                                            fmt.Println(fi15, fi15, v4)
                                            fmt.Println(v2)
                                        }
                                        _ = v51
                                    }
                                    _ = v17
                                }
                                v2 += 923
                            }
                            var v52 uint = 170
                            _ = fi16
                            _ = v52
                        }
                    case "a":
                        for fi53 := 0; fi53 < 12; fi53++ {
                            var v54 uint32 = 42
                            v6 = v6
                            _ = fi53
                            _ = v54
                        }
                    default:
                        if (2 != genFn0(42)) {
                            v6 = v6
                            var v55 float32 = 1.5
                            _ = v55
                        }
                    }
                default:
                    v56 := []int{-192, (fi15 / (((fi15 | (fi15 << ((65535) & 63)))) | 1)), fi15}
                    _ = v56
                }
                if (65535 > 0) {
                    var v57 int8 = (-(v6))
                    _ = v57
                } else {
                    if (v6 == (v6 + 127)) {
                        if ((!(42 >= 42)) || false) {
                            v58 := GenS3{F0: 42}
                            v59 := []int{(fi15 & (fi15 ^ fi15)), int(v1), 956, (-(fi15))}
                            _ = v58
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 18; fi60++ {
                            if true {
                                v3 -= float64(0.0)
                                v4 = 0
                                var v61 uint16 = 299
                                _ = v61
                            }
                            var v62 string = "a"
                            v63 := GenS3{F0: 666}
                            _ = fi60
                            _ = v62
                            _ = v63
                        }
                    }
                }
                for fi64 := 0; fi64 < 18; fi64++ {
                    switch "" {
                    case "":
                        v65 := GenS1{F0: 810, F1: 633, F2: 22}
                        _ = v65
                    default:
                        if false {
                            v66 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 127}
                            _ = v66
                        }
                    }
                    _ = fi64
                }
                _ = fi15
            }
        }
        v67 := GenS5{F0: -2147483648, F1: 1, F2: ""}
        var v68 int8 = 0
        _ = v67
        _ = v68
    }
    v1 += (v2 + (v1 * (v2 >> ((1) & 63))))
    var v69 int64 = 42
    var v70 uint64 = 959
    v71 := GenS6{F0: "\n\t", F1: int(v4)}
    if (0 == 65535) {
        v69 = v69
    } else {
        var v72 int16 = v2
        _ = v72
    }
    var v73 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
    _ = v73
}
