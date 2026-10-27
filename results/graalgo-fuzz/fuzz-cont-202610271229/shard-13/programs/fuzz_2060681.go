package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int64; F1 uint32; F2 bool }
type GenS5 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 uint, p1 uint16) uint32 {
    if false {
        for fi4 := 0; fi4 < 19; fi4++ {
            fmt.Println(p1)
            _ = fi4
        }
    } else {
        fmt.Println(p1, p1)
        v5 := GenS1{F0: p0, F1: 4294967295}
        _ = v5
    }
    var v6 uint = 255
    p0 -= (p0 ^ (v6 % (((v6 * 0)) | 1)))
    _ = p0
    _ = p1
    _ = v6
    return 65535
}

func main() {
    var v1 uint = 992
    var v2 int = 0
    var v3 int32 = -2147483648
    if false {
        v7 := []int{(v2 & 101)}
        var v8 int16 = -32768
        _ = v7
        _ = v8
    } else {
        if ("\n\t" >= (("" + "hello") + "a")) {
            var v9 uint = v1
            var v10 string = ""
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := []int{v2, (v2 | -178), -279}
                var v13 uint8 = 118
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
            _ = v9
        }
        v2 -= (v2 - v2)
    }
    v14 := GenS1{F0: 217, F1: genFn0(18446744073709551615, 992)}
    v2 -= 127
    v2 = int(v1)
    if (false && ((!false) || true)) {
        for fi15 := 0; fi15 < 19; fi15++ {
            if (42 == -375) {
                v16 := []int{(v2 / ((2147483647) | 1)), (v2 / (((v2 ^ (-(v2)))) | 1)), (v2 << ((42) & 63))}
                var v17 uint = (^(v1))
                switch v2 {
                case -798:
                    fmt.Println(v2)
                case 4294967295:
                    fmt.Println(v2, v3)
                case -26:
                    if false {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            v16 = append(v16, fi18)
                            _ = fi18
                        }
                        if true {
                            var v19 int32 = v3
                            var v20 int8 = -17
                            _ = v19
                            _ = v20
                        }
                        if false {
                            var v21 uint8 = 42
                            if (false || (255 != 764)) {
                                if false {
                                    var v22 uint = 65535
                                    _ = v22
                                }
                                var v23 int = fi15
                                var v24 uint8 = (v21 % ((v21) | 1))
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        } else {
                            for fi25 := 0; fi25 < 18; fi25++ {
                                switch "" {
                                case "hello":
                                    v16 = append(v16, (fi25 | (fi25 / (((fi25 + 128)) | 1))))
                                case "":
                                    var v26 int8 = 3
                                    _ = v26
                                case "\n\t":
                                    fmt.Println(v17)
                                default:
                                    v27 := GenS2{F0: v3}
                                    _ = v27
                                }
                                _ = fi25
                            }
                            var v28 uint32 = 1
                            _ = v28
                        }
                    } else {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            if (-1 == 0) {
                                if (9 >= 1) {
                                    v30 := GenS2{F0: v3}
                                    var v31 int8 = int8(v2)
                                    var v32 uint = (v17 << (((v1 + (v1 | v1))) & 63))
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 float32 = float32(float64((float32(0.1) + 3.14159)))
                                    v16 = append(v16, -9223372036854775808)
                                    _ = v33
                                }
                                var v34 int8 = 42
                                v34 = 100
                                _ = v34
                            } else {
                                var v35 bool = (!false)
                                switch "a" {
                                case "hello":
                                    v36 := GenS2{F0: v3}
                                    _ = v36
                                }
                                fmt.Println(v17, v3)
                                _ = v35
                            }
                            _ = fi29
                        }
                        switch "\n\t" {
                        case "":
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        case "\n\t":
                            for fi37 := 0; fi37 < 16; fi37++ {
                                fmt.Println(fi15, v2, fi15)
                                _ = fi37
                            }
                        case "a":
                            v2 = (^(fi15))
                        }
                        _ = len(v16)
                    }
                default:
                    if false {
                        var v38 bool = (true && false)
                        _ = v38
                    } else {
                        v39 := GenS2{F0: v3}
                        v1 = 2
                        var v40 int32 = (v3 + v3)
                        _ = v39
                        _ = v40
                    }
                }
                _ = v16
                _ = v17
            }
            _ = fi15
        }
    }
    var v41 uint64 = 0
    switch "foo bar" {
    case "\n\t":
        v41 += v41
    case "a":
        if (false && (877 > 433)) {
            var v42 int = v2
            _ = v42
        }
    case "hello":
        v43 := []int{32768, 290, -1}
        _ = v43
    default:
        if (1 != 127) {
            v44 := []int{v2, (-(v2)), v2, int(v3)}
            for fi45 := 0; fi45 < 9; fi45++ {
                v2 = fi45
                if (0.0 >= 0.1) {
                    v2 = (fi45 * (v2 % ((461) | 1)))
                }
                _ = fi45
            }
            for fi46 := 0; fi46 < 4; fi46++ {
                switch "foo bar" {
                case "":
                    if true {
                        fmt.Println(v2, v1, v41)
                    }
                }
                v47 := GenS3{F0: -51, F1: "hello"}
                _ = fi46
                _ = v47
            }
            _ = v44
        } else {
            v48 := GenS2{F0: v3}
            if true {
                var v49 int32 = (v3 + v3)
                v50 := []int{v2}
                fmt.Println(v2, v2)
                _ = v49
                _ = v50
            } else {
                for fi51 := 0; fi51 < 18; fi51++ {
                    v41 += v41
                    for fi52 := 0; fi52 < 8; fi52++ {
                        if (0 != 168) {
                            v41 += 142
                            var v53 uint32 = 949
                            _ = v53
                        } else {
                            var v54 int64 = -9223372036854775808
                            _ = v54
                        }
                        v55 := GenS4{F0: 1, F1: 65535, F2: (fi51 > (fi52 << ((v1) & 63)))}
                        _ = fi52
                        _ = v55
                    }
                    _ = fi51
                }
                for fi56 := 0; fi56 < 16; fi56++ {
                    for fi57 := 0; fi57 < 10; fi57++ {
                        v3 += 0
                        _ = fi57
                    }
                    if false {
                        var v58 uint8 = 36
                        var v59 float32 = 0.1
                        _ = v58
                        _ = v59
                    }
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        var v60 string = "a"
                        _ = v60
                    }
                    _ = fi56
                }
                var v61 bool = true
                _ = v61
            }
            var v62 uint32 = 2147483647
            _ = v48
            _ = v62
        }
    }
    fmt.Println(v3)
    if false {
        switch "foo bar" {
        case "foo bar":
            v63 := []int{}
            _ = v63
        case "\n\t":
            for fi64 := 0; fi64 < 13; fi64++ {
                fmt.Println(v41, fi64, fi64)
                _ = fi64
            }
        default:
            v65 := []int{(v2 + (-(v2)))}
            _ = v65
        }
    }
    if true {
        switch "" {
        case "":
            v66 := []int{v2}
            _ = v66
        default:
            v67 := []int{(v2 << ((v1) & 63)), v2, 489, v2}
            _ = v67
        }
    } else {
        for fi68 := 0; fi68 < 17; fi68++ {
            for fi69 := 0; fi69 < 10; fi69++ {
                var v70 uint8 = 0
                _ = fi69
                _ = v70
            }
            v71 := GenS3{F0: 71, F1: ("a" + "a")}
            _ = fi68
            _ = v71
        }
        v1 += 2
        switch 42 {
        case 32767:
            var v72 int8 = 67
            _ = v72
        case 2147483647:
            for fi73 := 0; fi73 < 1; fi73++ {
                v74 := GenS5{F0: 32767, F1: 32767, F2: "foo bar"}
                _ = fi73
                _ = v74
            }
        case 189:
            v75 := []int{v2, v2, -328, v2}
            _ = v75
        default:
            v41 = (v41 ^ 4294967295)
        }
    }
    v76 := []int{(v2 + v2), 442}
    if true {
        v77 := GenS4{F0: 809, F1: 439, F2: (("a" + "foo bar") < "hello")}
        var v78 uint32 = 908
        _ = v77
        _ = v78
    }
    _ = len(v76)
    for fi79 := 0; fi79 < 12; fi79++ {
        if len(v76) > 0 {
            v76[0] = -115
        }
        _ = fi79
    }
    v41 = v41
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v76:", len(v76), v76)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v41
    _ = v76
}
