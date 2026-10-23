package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint8; F2 int16 }

func main() {
    var v1 uint8 = 231
    var v2 int = 4294967295
    var v3 float32 = 0.1
    v2 = (v2 / ((-1) | 1))
    var v4 int16 = int16(v1)
    v5 := GenS1{F0: 4294967295}
    switch ("foo bar" + "\n\t") {
    case "":
        v6 := GenS2{F0: (v4 + v4), F1: -239, F2: ((53 > 1) && (!false))}
        _ = v6
    case "a":
        v1 += (v1 / ((86) | 1))
    case "foo bar":
        switch (v2 | 4294967295) {
        case 32768:
            var v7 float64 = (1e10 - ((float64(v3) - 1.5) / 3.14159))
            _ = v7
        case -845:
            switch v2 {
            case -2147483648:
                for fi8 := 0; fi8 < 2; fi8++ {
                    v9 := []int{v2, (v2 & (v2 ^ 1))}
                    _ = fi8
                    _ = v9
                }
            case 32768:
                switch 32767 {
                case 127:
                    var v10 uint64 = 600
                    _ = v10
                case 9223372036854775807:
                    for fi11 := 0; fi11 < 15; fi11++ {
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 int16 = (v4 * (^(v4)))
                            var v14 string = "a"
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS1{F0: 913}
                        switch "foo bar" {
                        case "":
                            var v16 uint32 = 0
                            _ = v16
                        case "hello":
                            v3 -= 0.1
                        case "foo bar":
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v2 = 0
                                _ = fi17
                            }
                        }
                        _ = fi11
                        _ = v15
                    }
                default:
                    v2 = 464
                }
            default:
                var v18 uint = 1
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 11; fi19++ {
                v20 := GenS3{F0: 4294967295, F1: -128}
                var v21 float32 = v3
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    }
    if (1 < -1) {
        var v22 uint32 = 42
        if true {
            v23 := GenS3{F0: 42, F1: int8(v2)}
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 uint16 = uint16(v22)
                var v26 int16 = 42
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(v2, v1)
            _ = v23
        } else {
            for fi27 := 0; fi27 < 1; fi27++ {
                v28 := GenS2{F0: -580, F1: -462, F2: true}
                _ = fi27
                _ = v28
            }
            v29 := GenS3{F0: (v22 + 2147483648), F1: int8(v22)}
            _ = v29
        }
        _ = v22
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        v31 := GenS2{F0: (^(v4)), F1: 1, F2: (!true)}
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 bool = false
            v1 = 246
            _ = fi32
            _ = v33
        }
        if (!true) {
            v34 := GenS2{F0: (v4 + (v4 << ((18446744073709551615) & 63))), F1: 527, F2: (65535 < 4294967295)}
            v4 = (v4 * (v4 / ((v4) | 1)))
            _ = v34
        } else {
            v3 -= (3.14159 / (((v3 * v3) + 0.0) - 0.1))
        }
        _ = fi30
        _ = v31
    }
    switch (("" + "foo bar") + "\n\t") {
    case "hello":
        if true {
            if true {
                v35 := []int{(v2 * (v2 ^ (-(v2)))), (-(v2))}
                _ = len(v35)
                v36 := GenS3{F0: 1, F1: -78}
                _ = v35
                _ = v36
            }
            v3 = 1.5
        } else {
            fmt.Println(v1, v1)
        }
    default:
        switch ("\n\t" + ("foo bar" + "")) {
        case "foo bar":
            var v37 int8 = -57
            _ = v37
        case "\n\t":
            for fi38 := 0; fi38 < 13; fi38++ {
                switch "a" {
                case "foo bar":
                    var v39 uint8 = v1
                    _ = v39
                case "a":
                    for fi40 := 0; fi40 < 16; fi40++ {
                        var v41 uint64 = 1
                        _ = fi40
                        _ = v41
                    }
                case "":
                    fmt.Println(v4)
                }
                _ = fi38
            }
        case "a":
            if (!false) {
                fmt.Println(v4, v2, v1)
            } else {
                fmt.Println(v3)
            }
        default:
            v42 := []int{v2, (v2 << ((132) & 63)), 128, v2}
            _ = v42
        }
    }
    var v43 uint64 = 1
    for fi44 := 0; fi44 < 16; fi44++ {
        for fi45 := 0; fi45 < 18; fi45++ {
            if false {
                switch "\n\t" {
                case "foo bar":
                    v2 = 0
                }
            }
            v46 := GenS4{F0: v4, F1: v43, F2: 4294967295}
            if (!false) {
                v1 += (v1 * 1)
                fmt.Println(v43)
                v4 = (v4 ^ 1000)
            } else {
                fmt.Println(v43, v4)
                v43 += (v43 / ((v43) | 1))
                var v47 uint8 = (^(v1))
                _ = v47
            }
            if (float64(0.1) == (float64(1e10) * (1.5 * -2.25))) {
                v1 += 134
                switch "foo bar" {
                case "a":
                    v48 := GenS5{F0: -9223372036854775808}
                    v49 := GenS5{F0: 385}
                    switch 32768 {
                    case 1:
                        var v50 bool = false
                        v2 = fi44
                        fmt.Println(v1, v1, v1)
                        _ = v50
                    default:
                        if true {
                            switch (^(v2)) {
                            case 2147483647:
                                var v51 uint = 1
                                _ = v51
                            case 474:
                                fmt.Println(v3)
                            default:
                                if (42 == (v43 << ((329) & 63))) {
                                    if (v4 < v4) {
                                        var v52 uint16 = 1
                                        _ = v52
                                    }
                                    var v53 uint = 526
                                    _ = v53
                                }
                            }
                            v54 := []int{(fi44 % (((fi45 + fi45)) | 1)), 32768, fi44, -1}
                            if len(v54) > 0 {
                                v54[0] = fi44
                            }
                            var v55 uint64 = 2
                            var v56 uint16 = 42
                            switch "a" {
                            case "\n\t":
                                v57 := GenS5{F0: 1}
                                _ = v57
                            }
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                    }
                    var v58 uint64 = 42
                    _ = v48
                    _ = v49
                    _ = v58
                case "hello":
                    if true {
                        var v59 float64 = 0.0
                        v60 := []int{fi45, fi44}
                        _ = v59
                        _ = v60
                    }
                default:
                    fmt.Println(fi45, v3, fi44)
                }
                switch "" {
                case "\n\t":
                    switch -2147483648 {
                    case 32767:
                        v3 -= v3
                    case 0:
                        v1 = 194
                    case 127:
                        var v61 float32 = v3
                        _ = v61
                    }
                case "a":
                    v62 := GenS5{F0: 381}
                    _ = v62
                }
            } else {
                var v63 uint64 = 350
                _ = v63
            }
            var v64 bool = (true && false)
            fmt.Println(v43, v3, v2)
            _ = fi45
            _ = v46
            _ = v64
        }
        switch "foo bar" {
        case "foo bar":
            v43 -= v43
        case "":
            v65 := []int{}
            _ = v65
        case "a":
            var v66 int64 = -541
            _ = v66
        default:
            var v67 bool = (!(0.0 > 1.5))
            _ = v67
        }
        _ = fi44
    }
    if (v43 != 4294967295) {
        var v68 float64 = 1e10
        var v69 int = 0
        v70 := GenS6{F0: v69, F1: (^(v1)), F2: v4}
        _ = v68
        _ = v69
        _ = v70
    } else {
        var v71 int64 = 42
        _ = v71
    }
    if false {
        var v72 int = -1
        v4 = (v4 % (((v4 << ((1) & 63))) | 1))
        v73 := GenS2{F0: (v4 >> ((978) & 63)), F1: -640, F2: false}
        _ = v72
        _ = v73
    }
    v74 := GenS6{F0: v2, F1: 93, F2: (v4 * 1000)}
    v75 := GenS2{F0: v4, F1: -2147483648, F2: true}
    if true {
        v1 = (v1 >> ((4294967295) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v74
    _ = v75
}
