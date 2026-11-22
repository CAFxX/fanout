package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 string; F1 string; F2 bool }
type GenS9 struct { F0 int; F1 int64; F2 int8 }
type GenS10 struct { F0 string; F1 uint32; F2 string }
type GenS11 struct { F0 uint }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    v2[916] = false
    if true {
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 int = (fi5 | (fi5 % ((-1) | 1)))
            _ = fi5
            _ = v6
        }
        var v7 float64 = (float64((float32((1.5 / float64(1.5))) * float32(3.14159))) + 0.0)
        _ = v7
    } else {
        var v8 float64 = 1e10
        _ = v8
    }
    var v9 uint32 = 2
    v9 = 65535
    v10 := GenS1{F0: 0, F1: 32767}
    v11 := []int{-2147483648}
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            fmt.Println(fi13, fi12)
            v9 = v9
            switch "a" {
            case "foo bar":
                v9 += v9
            }
            _ = fi13
        }
        v9 += (v9 * (^(v9)))
        var v14 uint32 = 0
        _ = fi12
        _ = v14
    }
    fmt.Println(v9)
    if (9223372036854775807 > 466) {
        var v15 int16 = -1
        var v16 float64 = 3.14159
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 2, F1: 42, F2: ""}
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "hello" {
        case "hello":
            var v19 int64 = -9223372036854775808
            _ = v19
        case "a":
            for fi20 := 0; fi20 < 15; fi20++ {
                if true {
                    if false {
                        if true {
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                if (42 > 756) {
                                    var v21 int64 = 0
                                    var v22 float64 = (3.14159 / (0.1 + 3.14159))
                                    v22 += v22
                                    _ = v21
                                    _ = v22
                                }
                            case "a":
                                v23 := GenS1{F0: -46, F1: fi20}
                                _ = v23
                            case "hello":
                                v24 := v1[(fi18 << ((322) & 63))]
                                _ = v24
                            }
                        } else {
                            v9 = 2147483648
                            v9 -= (v9 * 0)
                        }
                        var v25 uint16 = 2
                        v25 -= (v25 ^ 70)
                        _ = v25
                    }
                } else {
                    v9 += 199
                    if (531 < 42) {
                        v9 -= v9
                    }
                    v26 := GenS3{F0: 213}
                    switch "a" {
                    case "":
                        if ((false || (("a" < "") && (433 > 42))) || (!(false && false))) {
                            switch fi20 {
                            case 2147483647:
                                if false {
                                    fmt.Println(fi18, v9)
                                    fmt.Println(fi20, fi18)
                                    v9 = (v9 + v9)
                                }
                            case 42:
                                v9 = (v9 * (^(v9)))
                            default:
                                var v27 int = -1
                                fmt.Println(fi18, fi20)
                                _ = v27
                            }
                            if (1 > 18446744073709551615) {
                                if (-1 > -9223372036854775808) {
                                    switch (fi20 - fi20) {
                                    case 32768:
                                        v28 := GenS2{F0: 1, F1: 18446744073709551615, F2: "\n\t"}
                                        _ = v28
                                    default:
                                        v9 += v9
                                    }
                                    var v29 int32 = -33
                                    v29 += v29
                                    _ = v29
                                }
                                v30 := GenS4{F0: 1}
                                v31 := GenS5{F0: (float32(1.5) < ((1e10 + 3.14159) + -2.25))}
                                var v32 uint8 = 128
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        } else {
                            switch 128 {
                            case -1:
                                var v33 uint16 = 664
                                _ = v33
                            case 639:
                                v9 -= v9
                            case 4294967295:
                                var v34 int = 32767
                                _ = v34
                            }
                            switch fi20 {
                            case 0:
                                v9 = 1
                            case -494:
                                switch ("foo bar" + "foo bar") {
                                case "a":
                                    if false {
                                        v9 -= v9
                                        var v35 uint64 = 116
                                        _ = v35
                                    } else {
                                        v36 := GenS2{F0: 42, F1: 813, F2: "a"}
                                        if true {
                                            v37 := GenS6{F0: 195}
                                            v9 -= 425
                                            _ = v37
                                        }
                                        v1[(fi18 & (fi18 & (fi18 >> ((42) & 63))))] = false
                                        _ = v36
                                    }
                                default:
                                    var v38 uint8 = 101
                                    _ = v38
                                }
                            }
                            switch "hello" {
                            case "foo bar":
                                v9 -= 1
                            case "hello":
                                v9 += (v9 + 1)
                            }
                        }
                    case "a":
                        fmt.Println(fi20, fi18)
                    case "hello":
                        delete(v2, (^(v9)))
                    default:
                        switch 32767 {
                        case 127:
                            var v39 int = (-(fi20))
                            _ = v39
                        case 1:
                            v9 = 1
                            fmt.Println(fi20, fi18, fi18)
                        case 32768:
                            v9 += v9
                        default:
                            v40 := GenS7{F0: 1, F1: 4294967295}
                            _ = v40
                        }
                    }
                    _ = v26
                }
                switch "foo bar" {
                case "":
                    var v41 float64 = 3.14159
                    _ = v41
                case "foo bar":
                    v9 -= (v9 + v9)
                case "a":
                    v9 = v9
                }
                _ = fi20
            }
        default:
            var v42 int16 = -911
            _ = v42
        }
        _ = fi18
    }
    var v43 int64 = 0
    v43 = (v43 + -1)
    for fi44 := 0; fi44 < 6; fi44++ {
        for fi45 := 0; fi45 < 9; fi45++ {
            switch (fi44 / (((fi45 - fi45)) | 1)) {
            case 1:
                if (-128 > 21) {
                    v43 = v43
                    var v46 float32 = (float32(3.14159) * (1e10 - 3.14159))
                    _ = v46
                }
            case 2147483647:
                v47 := GenS8{F0: "\n\t", F1: "a", F2: false}
                if (false && false) {
                    v48 := GenS4{F0: 127}
                    _ = v48
                } else {
                    var v49 int8 = 100
                    _ = v49
                }
                _ = v47
            case 813:
                v4[("" + "")] = true
                v43 = (-(v43))
                var v50 int = fi44
                _ = v50
            default:
                v51 := GenS9{F0: (fi45 ^ (fi44 + (fi44 & 128))), F1: v43, F2: -1}
                _ = v51
            }
            _ = fi45
        }
        var v52 uint8 = 111
        if ((735 >= 2) && true) {
            for fi53 := 0; fi53 < 3; fi53++ {
                v54 := GenS10{F0: "a", F1: 4294967295, F2: (((("\n\t" + "\n\t") + "a") + "foo bar") + ("\n\t" + "foo bar"))}
                v55 := GenS11{F0: 2}
                _ = fi53
                _ = v54
                _ = v55
            }
        }
        _ = fi44
        _ = v52
    }
    if true {
        var v56 int16 = -677
        v9 = (v9 + v9)
        _ = v56
    } else {
        if (196 != 42) {
            var v57 int32 = int32(v43)
            v9 = v9
            v9 = v9
            _ = v57
        }
        if (v43 >= v43) {
            v43 -= v43
        } else {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        var v58 uint16 = 728
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v9
}
