package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS6 struct { F0 int8; F1 string; F2 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool }
func genFn0() uint16 {
    v4 := GenS1{F0: 1, F1: 32768}
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 1, F1: (fi5 % ((fi5) | 1))}
        var v7 int16 = -260
        v8 := []int{(fi5 >> ((160) & 63)), 32768}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 127, F1: 560, F2: 240}
    var v10 uint16 = 2
    _ = v10
    _ = v4
    _ = v9
    return (^(v10))
}

func main() {
    var v1 int16 = 1000
    var v2 int32 = -817
    var v3 int16 = (v1 * v1)
    for fi11 := 0; fi11 < 7; fi11++ {
        switch "" {
        case "a":
            var v12 uint64 = 4294967295
            _ = v12
        case "foo bar":
            v3 += v3
        case "hello":
            v13 := GenS2{F0: 42, F1: (v3 | v1), F2: v2}
            _ = v13
        }
        v14 := []int{}
        _ = fi11
        _ = v14
    }
    var v15 bool = (49 <= 65535)
    for fi16 := 0; fi16 < 7; fi16++ {
        if v15 {
            v17 := []int{fi16, fi16}
            var v18 int8 = 100
            if (!((fi16 * fi16) >= fi16)) {
                for fi19 := 0; fi19 < 20; fi19++ {
                    if len(v17) > 0 {
                        v17[0] = 127
                    }
                    var v20 uint64 = 815
                    _ = fi19
                    _ = v20
                }
            }
            _ = v17
            _ = v18
        }
        var v21 int8 = 127
        v22 := GenS3{F0: (^(v3)), F1: 42}
        _ = fi16
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if false {
            v24 := []int{32768, fi23}
            _ = v24
        } else {
            v2 = 42
        }
        _ = fi23
    }
    v15 = true
    v25 := GenS4{F0: "\n\t", F1: 766}
    if (v15 && ((339 != 255) || ((v15 && v15) && v15))) {
        for fi26 := 0; fi26 < 11; fi26++ {
            for fi27 := 0; fi27 < 13; fi27++ {
                v1 += (v3 | v1)
                v3 += v1
                v2 = (v2 & (v2 << ((42) & 63)))
                _ = fi27
            }
            v28 := GenS5{F0: 42, F1: 1, F2: (v1 ^ (v1 / (((v3 & -22)) | 1)))}
            v29 := GenS6{F0: 37, F1: "", F2: 65535}
            _ = fi26
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 int = (^(fi30))
            _ = fi30
            _ = v31
        }
        var v32 int = 9223372036854775807
        _ = v32
    }
    v1 = 1
    if true {
        var v33 bool = false
        var v34 int32 = 792
        var v35 int64 = 767
        _ = v33
        _ = v34
        _ = v35
    } else {
        v2 += v2
        switch 0 {
        case -100:
            fmt.Println(v15)
        }
        if (-9223372036854775808 > 32768) {
            var v36 uint8 = 255
            _ = v36
        } else {
            var v37 string = ""
            switch 357 {
            case 54:
                v3 -= (v1 << ((2) & 63))
            default:
                v38 := []int{1, 1, 127, 512}
                _ = v38
            }
            v39 := GenS6{F0: int8(v3), F1: v37, F2: 255}
            _ = v37
            _ = v39
        }
    }
    fmt.Println(v3)
    switch "foo bar" {
    case "foo bar":
        var v40 int8 = 1
        _ = v40
    case "\n\t":
        v2 = 42
    }
    v41 := []int{4294967295, 128, 32767, -1}
    v3 = -32768
    v2 -= (v2 ^ -436)
    var v42 string = (("a" + "\n\t") + "\n\t")
    if (true || (42 != 268)) {
        v3 -= v3
        v43 := GenS4{F0: "\n\t", F1: 2147483648}
        for fi44 := 0; fi44 < 17; fi44++ {
            if ((v42 + ((v42 + ((v42 + ("" + (((v42 + ((v42 + (("" + "") + v42)) + ((v42 + v42) + "\n\t"))) + v42) + v42))) + (v42 + "\n\t"))) + "foo bar")) < (v42 + "\n\t")) {
                v41 = append(v41, fi44)
                v45 := GenS3{F0: v1, F1: 391}
                _ = v45
            } else {
                fmt.Println(fi44, v2, fi44)
            }
            _ = fi44
        }
        _ = v43
    } else {
        if (9223372036854775807 >= -9223372036854775808) {
            var v46 string = v42
            fmt.Println(v3, v2)
            fmt.Println(v46)
            _ = v46
        }
    }
    if (1 != -195) {
        fmt.Println(v2)
        if (818 > 18446744073709551615) {
            v3 -= (^(v1))
            _ = len(v41)
            if v15 {
                var v47 uint64 = 1
                var v48 string = v42
                switch (v42 + v42) {
                case "a":
                    var v49 uint32 = 2
                    _ = v49
                case "foo bar":
                    switch v42 {
                    case "":
                        var v50 float64 = 0.0
                        _ = v50
                    default:
                        v51 := GenS7{F0: 42}
                        _ = v51
                    }
                case "hello":
                    for fi52 := 0; fi52 < 2; fi52++ {
                        switch v42 {
                        case "a":
                            v41 = append(v41, 0)
                        case "hello":
                            fmt.Println(v42, fi52)
                        case "\n\t":
                            if v15 {
                                v48 = v48
                                var v53 bool = false
                                if v15 {
                                    if ((fi52 | fi52) > (fi52 * (fi52 % ((fi52) | 1)))) {
                                        if true {
                                            if len(v41) > 0 {
                                                v41[0] = -273
                                            }
                                            for fi54 := 0; fi54 < 5; fi54++ {
                                                if (255 >= 0) {
                                                    fmt.Println(fi54)
                                                    v55 := GenS8{F0: (-1 <= v1)}
                                                    var v56 int64 = -9223372036854775808
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    var v57 uint = 1
                                                    _ = v57
                                                }
                                                _ = fi54
                                            }
                                        }
                                        v42 = v42
                                        fmt.Println(fi52, v2)
                                    }
                                } else {
                                    for fi58 := 0; fi58 < 6; fi58++ {
                                        var v59 uint8 = uint8(v47)
                                        switch (v48 + v48) {
                                        case "\n\t":
                                            if len(v41) > 0 {
                                                v41[0] = fi52
                                            }
                                        case "hello":
                                            fmt.Println(v42)
                                        default:
                                            fmt.Println(fi52)
                                        }
                                        _ = fi58
                                        _ = v59
                                    }
                                    fmt.Println(v47, v2, v42)
                                    v60 := GenS4{F0: v42, F1: 0}
                                    _ = v60
                                }
                                _ = v53
                            }
                        default:
                            v61 := GenS3{F0: (v1 % ((v3) | 1)), F1: -1}
                            _ = v61
                        }
                        _ = fi52
                    }
                }
                _ = v47
                _ = v48
            }
        }
        v42 = v42
    } else {
        if (!v15) {
            var v62 uint32 = 783
            var v63 float64 = (-2.25 * 0.1)
            _ = v62
            _ = v63
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v41
    _ = v42
}
