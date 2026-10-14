package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 bool }

func main() {
    var v1 uint8 = 128
    var v2 float64 = float64((0.1 - (float32(0.0) + 0.1)))
    var v3 bool = (!true)
    var v4 int8 = 1
    var v5 float64 = 0.1
    var v6 float32 = (1.5 / 3.14159)
    var v7 int32 = -634
    v4 -= v4
    fmt.Println(v7, v4, v2)
    fmt.Println(v3, v7)
    var v8 uint32 = 2
    fmt.Println(v4, v2)
    v6 += (v6 * 0.1)
    v7 = -2147483648
    if (!(1 >= 0)) {
        fmt.Println(v8, v4, v8)
        for fi9 := 0; fi9 < 18; fi9++ {
            if (!(2 == v1)) {
                if (-32768 < 32767) {
                    v10 := GenS1{F0: "a", F1: (v8 / ((4294967295) | 1))}
                    v2 -= 3.14159
                    _ = v10
                }
            }
            var v11 uint16 = 255
            if (v3 && ((v4 - (v4 & 50)) <= (v4 % (((v4 & 1)) | 1)))) {
                if v3 {
                    switch int(v1) {
                    case 128:
                        v1 += (^(v1))
                    case -1:
                        if (!(42 != -9223372036854775808)) {
                            for fi12 := 0; fi12 < 11; fi12++ {
                                v13 := GenS1{F0: "foo bar", F1: v8}
                                _ = fi12
                                _ = v13
                            }
                            for fi14 := 0; fi14 < 19; fi14++ {
                                v15 := GenS2{F0: 18446744073709551615, F1: 105}
                                fmt.Println(v4)
                                var v16 uint32 = (v8 & v8)
                                if false {
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        v17 := []int{fi9}
                                        _ = v17
                                    case "a":
                                        var v18 uint = uint(v16)
                                        var v19 uint8 = (v1 / ((1) | 1))
                                        if v3 {
                                            switch "hello" {
                                            case "a":
                                                v11 = (v11 | v11)
                                            case "foo bar":
                                                fmt.Println(v16)
                                            default:
                                                v19 = v19
                                            }
                                            var v20 uint16 = (v11 | v11)
                                            switch "foo bar" {
                                            case "a":
                                                switch (fi9 * 127) {
                                                case 32768:
                                                    v21 := GenS2{F0: 956, F1: 2147483648}
                                                    _ = v21
                                                default:
                                                    var v22 int64 = -9223372036854775808
                                                    _ = v22
                                                }
                                            default:
                                                v23 := GenS3{F0: "hello"}
                                                _ = v23
                                            }
                                            _ = v20
                                        }
                                        _ = v18
                                        _ = v19
                                    default:
                                        var v24 bool = ((fi9 & fi14) <= -1)
                                        _ = v24
                                    }
                                }
                                var v25 uint8 = v1
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v25
                            }
                            fmt.Println(v11, v5, v2)
                        }
                    default:
                        v26 := []int{(fi9 - (fi9 * fi9)), (fi9 | (fi9 * fi9)), fi9}
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 17; fi27++ {
                        switch fi9 {
                        case 9223372036854775807:
                            if (!v3) {
                                if false {
                                    v28 := GenS3{F0: "a"}
                                    _ = v28
                                } else {
                                    v29 := GenS3{F0: ("hello" + ("" + "\n\t"))}
                                    v30 := GenS1{F0: "hello", F1: 0}
                                    if v3 {
                                        var v31 float64 = 1.5
                                        v32 := []int{(fi27 + -77), (fi27 >> ((317) & 63))}
                                        if ((v4 ^ v4) >= (v4 & (v4 % ((v4) | 1)))) {
                                            var v33 uint64 = 4294967295
                                            if len(v32) > 0 {
                                                v32[0] = int(v11)
                                            }
                                            _ = v33
                                        }
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = v29
                                    _ = v30
                                }
                                fmt.Println(v6)
                                v1 -= (v1 | (v1 - v1))
                            }
                        }
                        _ = fi27
                    }
                } else {
                    fmt.Println(v7)
                    if true {
                        v34 := []int{(fi9 << ((945) & 63))}
                        _ = v34
                    } else {
                        fmt.Println(v3, v4, v6)
                    }
                    var v35 uint64 = 0
                    _ = v35
                }
            }
            _ = fi9
            _ = v11
        }
        v36 := []int{32768, 32767, 887}
        _ = v36
    } else {
        v6 = 1e10
        if true {
            for fi37 := 0; fi37 < 3; fi37++ {
                switch "a" {
                case "":
                    var v38 int16 = 32767
                    _ = v38
                case "foo bar":
                    if v3 {
                        v39 := []int{(fi37 << ((485) & 63))}
                        _ = v39
                    }
                default:
                    var v40 int64 = 0
                    _ = v40
                }
                v41 := GenS2{F0: 4294967295, F1: v8}
                _ = fi37
                _ = v41
            }
            v8 = 2
            v42 := GenS2{F0: 4294967295, F1: v8}
            _ = v42
        }
    }
    v43 := GenS4{F0: v4, F1: 4294967295}
    for fi44 := 0; fi44 < 3; fi44++ {
        switch "\n\t" {
        case "hello":
            for fi45 := 0; fi45 < 4; fi45++ {
                var v46 int32 = (v7 * (v7 / (((v7 & v7)) | 1)))
                v1 += 127
                _ = fi45
                _ = v46
            }
        case "foo bar":
            var v47 int64 = 0
            _ = v47
        }
        if (139 > 2) {
            for fi48 := 0; fi48 < 18; fi48++ {
                v7 += v7
                _ = fi48
            }
            switch "a" {
            case "a":
                for fi49 := 0; fi49 < 13; fi49++ {
                    fmt.Println(fi44, fi49)
                    v50 := GenS5{F0: (!v3)}
                    switch ("" + "a") {
                    case "":
                        if v3 {
                            fmt.Println(fi44)
                        } else {
                            v51 := GenS3{F0: "\n\t"}
                            v52 := []int{1, (fi44 % ((509) | 1)), (fi49 >> ((4294967295) & 63))}
                            var v53 int64 = -1
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                    default:
                        var v54 uint64 = uint64(v8)
                        _ = v54
                    }
                    _ = fi49
                    _ = v50
                }
            case "\n\t":
                if (0.1 < v2) {
                    v4 = (v4 | (^(v4)))
                    var v55 int8 = v4
                    _ = v55
                }
            default:
                if (42 <= 0) {
                    v1 = (^(v1))
                    for fi56 := 0; fi56 < 15; fi56++ {
                        if false {
                            v7 += (v7 ^ (v7 ^ 2147483647))
                            v57 := []int{}
                            var v58 float32 = float32(v5)
                            _ = v57
                            _ = v58
                        } else {
                            if (v1 != (^(v1))) {
                                var v59 float32 = (v6 / -2.25)
                                _ = v59
                            }
                            fmt.Println(v6)
                        }
                        _ = fi56
                    }
                } else {
                    for fi60 := 0; fi60 < 9; fi60++ {
                        v61 := []int{fi44, (fi60 & 127), fi44, fi44}
                        v61 = append(v61, 158)
                        v2 += (v2 + float64((v6 + 3.14159)))
                        _ = fi60
                        _ = v61
                    }
                    var v62 int8 = (v4 * 0)
                    _ = v62
                }
            }
            v8 = 2147483647
        }
        for fi63 := 0; fi63 < 2; fi63++ {
            var v64 uint32 = v8
            _ = fi63
            _ = v64
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
