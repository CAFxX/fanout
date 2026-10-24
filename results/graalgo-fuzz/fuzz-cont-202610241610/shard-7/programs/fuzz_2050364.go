package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int32; F2 string }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int64; F1 uint8 }

func main() {
    var v1 bool = true
    var v2 int16 = 42
    var v3 float32 = (((1e10 - 3.14159) - 3.14159) + float32(-2.25))
    var v4 float64 = 0.0
    var v5 float64 = float64(3.14159)
    var v6 bool = (46 >= 26)
    var v7 uint32 = 65535
    if (-195 <= -410) {
        var v8 uint8 = 2
        switch -1 {
        case -317:
            switch 127 {
            case -339:
                v2 += 32767
            case 32767:
                switch "\n\t" {
                case "":
                    var v9 int = 4294967295
                    _ = v9
                case "hello":
                    v10 := GenS1{F0: -1, F1: 622}
                    _ = v10
                case "foo bar":
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 uint8 = v8
                        var v13 int8 = 38
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                default:
                    switch -2147483648 {
                    case 2147483647:
                        var v14 string = ""
                        _ = v14
                    case 9223372036854775807:
                        if (false || v1) {
                            v15 := GenS1{F0: 34, F1: (^(v2))}
                            _ = v15
                        }
                    case -268:
                        fmt.Println(v1)
                    default:
                        var v16 uint8 = (v8 ^ (v8 << ((2) & 63)))
                        _ = v16
                    }
                }
            case -1:
                if false {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            if (((!(((v3 + (1.5 - float32(v5))) / (v3 - -2.25)) >= 1e10)) || false) && v6) {
                                var v19 float64 = v4
                                if (uint16(v8) == 255) {
                                    v8 = 113
                                    switch "foo bar" {
                                    case "hello":
                                        switch int(v7) {
                                        case 32767:
                                            if (-9223372036854775808 > 890) {
                                                var v20 uint8 = v8
                                                _ = v20
                                            }
                                        case 2147483647:
                                            v6 = v1
                                            v2 = v2
                                            v21 := GenS1{F0: -128, F1: v2}
                                            _ = v21
                                        case 4294967295:
                                            v22 := GenS2{F0: 42, F1: "hello", F2: 9223372036854775807}
                                            _ = v22
                                        default:
                                            v3 = v3
                                        }
                                    case "":
                                        v23 := []int{fi17}
                                        _ = v23
                                    case "a":
                                        v24 := GenS2{F0: -890, F1: "", F2: -41}
                                        _ = v24
                                    default:
                                        var v25 int = (fi17 % ((fi17) | 1))
                                        _ = v25
                                    }
                                } else {
                                    var v26 uint = 42
                                    _ = v26
                                }
                                var v27 int16 = v2
                                if ((1 > 100) || ("" < ("hello" + ""))) {
                                    var v28 uint32 = 667
                                    switch "foo bar" {
                                    case "hello":
                                        switch (fi17 / (((fi18 / ((fi17) | 1))) | 1)) {
                                        case 0:
                                            v29 := []int{-105, 32768, (fi18 % ((fi17) | 1))}
                                            _ = v29
                                        case -961:
                                            if v6 {
                                                v30 := GenS3{F0: 65535}
                                                v31 := GenS3{F0: (^(v28))}
                                                v5 += v5
                                                _ = v30
                                                _ = v31
                                            } else {
                                                switch "foo bar" {
                                                case "":
                                                    v32 := []int{231, (fi18 >> ((42) & 63))}
                                                    _ = v32
                                                }
                                                v33 := GenS4{F0: -1, F1: 1, F2: "a"}
                                                _ = v33
                                            }
                                        case 4294967295:
                                            switch ("foo bar" + ((("\n\t" + "\n\t") + "foo bar") + "\n\t")) {
                                            case "a":
                                                v6 = ((v5 + -2.25) >= (1e10 + float64(float32(v4))))
                                            case "\n\t":
                                                v34 := []int{fi17}
                                                _ = v34
                                            default:
                                                var v35 bool = (894 != 283)
                                                _ = v35
                                            }
                                        default:
                                            switch fi18 {
                                            case -153:
                                                v27 -= (v2 ^ (v2 | 1000))
                                            case 42:
                                                v36 := []int{fi17, (fi17 << ((42) & 63)), (fi18 % ((fi18) | 1))}
                                                _ = v36
                                            case 1:
                                                switch "" {
                                                case "":
                                                    if true {
                                                        v37 := GenS4{F0: -505, F1: int32(v8), F2: "a"}
                                                        v6 = (2147483647 < 42)
                                                        _ = v37
                                                    } else {
                                                        fmt.Println(v5)
                                                        v19 = v19
                                                        v38 := []int{32768}
                                                        _ = v38
                                                    }
                                                case "foo bar":
                                                    v39 := GenS1{F0: 0, F1: v2}
                                                    _ = v39
                                                default:
                                                    switch 2147483647 {
                                                    case 32767:
                                                        v5 = ((1e10 - v5) / 0.1)
                                                    case 2147483647:
                                                        var v40 string = "\n\t"
                                                        _ = v40
                                                    case 128:
                                                        if ((v2 + (-(v27))) == v27) {
                                                            fmt.Println(fi17)
                                                            switch "a" {
                                                            case "foo bar":
                                                                if ((v3 / float32((float64((v3 + (v3 + v3))) * v4))) >= v3) {
                                                                    var v41 int64 = 42
                                                                    v42 := GenS2{F0: -367, F1: "", F2: 42}
                                                                    fmt.Println(v3, v41)
                                                                    _ = v41
                                                                    _ = v42
                                                                } else {
                                                                    fmt.Println(v3)
                                                                    v43 := GenS4{F0: 1, F1: -2147483648, F2: "hello"}
                                                                    _ = v43
                                                                }
                                                            case "a":
                                                                v27 -= v27
                                                            case "hello":
                                                                v44 := []int{(fi18 / ((fi18) | 1)), -1}
                                                                _ = v44
                                                            }
                                                        } else {
                                                            var v45 uint16 = 0
                                                            v46 := GenS2{F0: -2147483648, F1: "hello", F2: -9223372036854775808}
                                                            v27 -= (^(v27))
                                                            v47 := []int{(fi17 ^ fi18), 0, (fi18 * fi17), fi17}
                                                            var v48 int = (fi17 / ((fi18) | 1))
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    }
                                                }
                                            default:
                                                var v49 int32 = 710
                                                _ = v49
                                            }
                                            var v50 int = fi18
                                            _ = v50
                                        }
                                    default:
                                        if false {
                                            v3 = ((0.1 - v3) + -2.25)
                                            var v51 float64 = v19
                                            _ = v51
                                        }
                                    }
                                    var v52 uint64 = 18446744073709551615
                                    _ = v28
                                    _ = v52
                                } else {
                                    var v53 int = int(v27)
                                    _ = v53
                                }
                                fmt.Println(fi17)
                                if v6 {
                                    var v54 float32 = v3
                                    _ = v54
                                }
                                _ = v19
                                _ = v27
                            }
                            _ = fi18
                        }
                        var v55 uint32 = (v7 / ((42) | 1))
                        for fi56 := 0; fi56 < 14; fi56++ {
                            v57 := []int{(fi17 / ((128) | 1)), fi17, fi17}
                            var v58 int16 = v2
                            v8 = 2
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                        _ = fi17
                        _ = v55
                    }
                    var v59 int64 = -1
                    var v60 int8 = -128
                    _ = v59
                    _ = v60
                } else {
                    var v61 float64 = (v4 * v5)
                    _ = v61
                }
            default:
                v4 += float64((v3 * v3))
            }
        default:
            for fi62 := 0; fi62 < 15; fi62++ {
                if v6 {
                    if (778 >= 42) {
                        var v63 string = ("a" + "a")
                        _ = v63
                    } else {
                        var v64 uint32 = (v7 + 821)
                        switch "\n\t" {
                        case "hello":
                            v65 := []int{32768}
                            _ = v65
                        }
                        if v6 {
                            v2 = (v2 % ((v2) | 1))
                            if ((^(v8)) < 42) {
                                v8 -= v8
                                var v66 uint64 = 4294967295
                                _ = v66
                            } else {
                                if (4294967295 != 2) {
                                    v7 = v64
                                } else {
                                    var v67 float32 = (v3 * (((v3 + v3) + ((3.14159 / 0.1) + (v3 + 3.14159))) + v3))
                                    _ = v67
                                }
                            }
                            switch fi62 {
                            case 127:
                                v68 := []int{fi62, fi62}
                                _ = v68
                            case -1:
                                var v69 uint16 = 65535
                                _ = v69
                            case 559:
                                v70 := GenS2{F0: 770, F1: "foo bar", F2: -91}
                                _ = v70
                            default:
                                var v71 uint8 = v8
                                _ = v71
                            }
                        }
                        _ = v64
                    }
                    v8 = (^(v8))
                }
                _ = fi62
            }
        }
        for fi72 := 0; fi72 < 6; fi72++ {
            var v73 string = "hello"
            for fi74 := 0; fi74 < 12; fi74++ {
                fmt.Println(v2)
                v75 := GenS5{F0: 5, F1: -1}
                _ = fi74
                _ = v75
            }
            var v76 bool = (v7 < v7)
            _ = fi72
            _ = v73
            _ = v76
        }
        _ = v8
    }
    switch (("foo bar" + "") + ("hello" + "foo bar")) {
    case "":
        switch "hello" {
        case "a":
            fmt.Println(v2, v6)
        case "hello":
            v7 += (v7 + (v7 >> ((42) & 63)))
        default:
            var v77 float64 = v5
            _ = v77
        }
    default:
        var v78 uint = 42
        _ = v78
    }
    v6 = (v7 > v7)
    switch "" {
    case "":
        v79 := GenS6{F0: 0, F1: 0}
        _ = v79
    default:
        var v80 bool = (v6 || (v3 < (v3 - 1e10)))
        _ = v80
    }
    v3 = ((v3 + (1.5 * v3)) + v3)
    if v6 {
        v81 := GenS1{F0: -1, F1: v2}
        switch "hello" {
        case "foo bar":
            var v82 float32 = v3
            _ = v82
        case "hello":
            var v83 int = 127
            _ = v83
        default:
            v1 = (v6 && false)
        }
        _ = v81
    } else {
        v2 += v2
        switch "" {
        case "":
            if false {
                v84 := []int{}
                v3 += v3
                _ = v84
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
