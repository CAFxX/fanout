package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
func genFn0(p0 int16, p1 uint64) uint {
    v7 := GenS1{F0: "foo bar"}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int8, p1 string, p2 uint64) uint32 {
    p0 = -34
    p0 = (p0 / (((p0 ^ -1)) | 1))
    fmt.Println(p2, p0)
    if true {
        v9 := GenS1{F0: (p1 + "a")}
        switch "\n\t" {
        case "hello":
            switch -2147483648 {
            case -2147483648:
                var v10 uint64 = (^(p2))
                _ = v10
            default:
                v11 := GenS2{F0: 454}
                _ = v11
            }
        default:
            switch p1 {
            case "":
                fmt.Println(p1)
            case "\n\t":
                switch p1 {
                case "":
                    if false {
                        var v12 bool = ((-2.25 + 3.14159) > (((0.0 * (1e10 + (-2.25 - ((3.14159 + 0.1) + (0.1 - float32(3.14159)))))) + 0.0) / 1.0))
                        _ = v12
                    }
                default:
                    p2 += p2
                }
            }
        }
        _ = v9
    } else {
        var v13 int = -873
        _ = v13
    }
    var v14 int8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 894
    var v3 uint32 = 1
    v4 := []string{"a", "foo bar"}
    var v5 float64 = 0.1
    v6 := make(map[int]uint32)
    v5 -= float64((-2.25 - 0.0))
    if true {
        var v15 uint32 = (^(v3))
        _ = v15
    } else {
        v2 += v2
        v3 -= genFn1(103, "", 42)
        var v16 float64 = v5
        _ = v16
    }
    v2 -= v2
    var v17 uint16 = 259
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int16 = int16(v17)
        switch fi18 {
        case 892:
            for fi20 := 0; fi20 < 2; fi20++ {
                v5 = v5
                v3 += v3
                _ = fi20
            }
        default:
            var v21 int32 = 204
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v4 = append(v4, "hello")
    var v22 int = 0
    if false {
        v17 += (v17 ^ v17)
        var v23 string = "foo bar"
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v23
    } else {
        if (!true) {
            switch ("hello" + "a") {
            case "\n\t":
                switch v22 {
                case 42:
                    delete(v6, v22)
                case 32768:
                    fmt.Println(v3, v2, v5)
                case 2147483647:
                    if (-773 >= 42) {
                        fmt.Println(v3, v2)
                    }
                default:
                    for fi24 := 0; fi24 < 14; fi24++ {
                        delete(v6, fi24)
                        _ = fi24
                    }
                }
            case "":
                v25 := GenS2{F0: v3}
                _ = v25
            case "foo bar":
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 int8 = 0
                    if false {
                        _ = len(v6)
                    } else {
                        delete(v1, 127)
                        _ = len(v6)
                    }
                    if false {
                        if true {
                            for fi28 := 0; fi28 < 19; fi28++ {
                                var v29 bool = (v17 <= v17)
                                delete(v6, fi26)
                                _ = len(v6)
                                if (v5 <= v5) {
                                    _ = len(v1)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v2 -= v2
                                }
                                _ = fi28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 18; fi30++ {
                                switch fi26 {
                                case 127:
                                    if false {
                                        if ((v27 * (v27 >> ((42) & 63))) < v27) {
                                            v22 -= (v22 & v22)
                                            fmt.Println(v2, v22)
                                            var v31 string = "a"
                                            _ = v31
                                        }
                                    }
                                    v32 := GenS1{F0: ""}
                                    switch (v22 / ((4294967295) | 1)) {
                                    case -1:
                                        v5 -= float64((-2.25 * 1.5))
                                    default:
                                        if true {
                                            delete(v6, (fi30 - v22))
                                            v2 = v2
                                            v17 -= (^(v17))
                                        }
                                    }
                                    v5 = float64(0.1)
                                    _ = v32
                                }
                                v17 = v17
                                if (false || true) {
                                    v33 := GenS3{F0: -934, F1: 42, F2: 18446744073709551615}
                                    v34 := GenS3{F0: -695, F1: 0, F2: 4294967295}
                                    var v35 uint8 = 148
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                _ = fi30
                            }
                            v36 := GenS2{F0: v3}
                            _ = v36
                        } else {
                            fmt.Println(v3, v2, v3)
                            for fi37 := 0; fi37 < 15; fi37++ {
                                v22 += v22
                                _ = fi37
                            }
                            for fi38 := 0; fi38 < 8; fi38++ {
                                var v39 uint64 = 387
                                var v40 int32 = 378
                                v40 = -1
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                        }
                        var v41 string = ("a" + "\n\t")
                        v42 := GenS2{F0: genFn1((v27 % (((v27 & (v27 - (v27 << ((65535) & 63))))) | 1)), "", 2)}
                        _ = v41
                        _ = v42
                    }
                    _ = fi26
                    _ = v27
                }
            }
            v2 += 9223372036854775807
            var v43 uint32 = (v3 << ((18446744073709551615) & 63))
            _ = v43
        } else {
            fmt.Println(v17)
            if (true && false) {
                fmt.Println(v22, v2, v22)
                var v44 uint = 1
                _ = v44
            }
            v45 := GenS4{F0: 4294967295}
            _ = v45
        }
    }
    v46 := GenS5{F0: genFn0(0, 42)}
    var v47 float32 = (1e10 / float32(float64(0.1)))
    v48 := GenS6{F0: v2}
    for fi49 := 0; fi49 < 7; fi49++ {
        var v50 float64 = (v5 + 0.0)
        var v51 int32 = 799
        v52 := v6[v22]
        _ = fi49
        _ = v50
        _ = v51
        _ = v52
    }
    var v53 uint32 = 65535
    switch ("foo bar" + "\n\t") {
    case "a":
        v54 := GenS5{F0: 0}
        _ = v54
    case "hello":
        var v55 int64 = int64(v17)
        _ = v55
    case "":
        var v56 float32 = v47
        _ = v56
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        v47 += v47
        if ((1.5 + float64((0.0 * (float32((3.14159 - v5)) + v47)))) >= v5) {
            var v58 int64 = v2
            if false {
                var v59 int64 = v2
                var v60 float32 = 3.14159
                if true {
                    switch "" {
                    case "foo bar":
                        for fi61 := 0; fi61 < 14; fi61++ {
                            v2 = -1
                            _ = fi61
                        }
                    case "hello":
                        v59 = (v2 - (v2 | (v59 - v2)))
                    }
                    switch (v22 | v22) {
                    case 32767:
                        var v62 float32 = v47
                        _ = v62
                    case -1:
                        var v63 uint32 = (^(v53))
                        _ = v63
                    }
                    v58 += v58
                }
                _ = v59
                _ = v60
            } else {
                for fi64 := 0; fi64 < 18; fi64++ {
                    v2 = v2
                    _ = fi64
                }
                var v65 int64 = v58
                v17 = v17
                _ = v65
            }
            var v66 uint32 = (v3 >> ((483) & 63))
            _ = v58
            _ = v66
        } else {
            v6[(v22 - (-(v22)))] = 4294967295
            var v67 int = 852
            for fi68 := 0; fi68 < 10; fi68++ {
                var v69 float32 = (1.5 + 0.0)
                v69 -= 1e10
                switch 42 {
                case -2147483648:
                    if false {
                        v22 = v22
                        if false {
                            var v70 int8 = 0
                            v6[-9223372036854775808] = v53
                            _ = v70
                        } else {
                            var v71 uint8 = 128
                            _ = v71
                        }
                        if false {
                            switch "hello" {
                            case "a":
                                var v72 int32 = 1
                                _ = v72
                            default:
                                var v73 int = fi68
                                _ = v73
                            }
                            _ = len(v6)
                            if true {
                                var v74 bool = true
                                var v75 int64 = (^(v2))
                                v76 := GenS7{F0: ((true || (!(false || ((((v53 << ((genFn0(0, 4294967295)) & 63)) == v3) || v74) && (81 != 128))))) && v74)}
                                _ = v74
                                _ = v75
                                _ = v76
                            } else {
                                v17 = v17
                                v77 := v1[v67]
                                var v78 uint = 471
                                _ = v77
                                _ = v78
                            }
                            if (!true) {
                                v53 = (v53 * v53)
                                var v79 string = "hello"
                                _ = v79
                            } else {
                                var v80 int32 = int32(v17)
                                _ = v80
                            }
                            v81 := GenS3{F0: 1, F1: 4294967295, F2: 18446744073709551615}
                            _ = v81
                        } else {
                            _ = len(v1)
                            v1[v67] = (1000 != 32767)
                        }
                    }
                case 128:
                    switch v22 {
                    case -1:
                        v2 = (v2 >> ((42) & 63))
                    case 4294967295:
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                    case 2147483647:
                        v67 -= (v67 ^ (v67 - fi57))
                    }
                default:
                    var v82 float64 = (((v5 - 3.14159) / -2.25) * v5)
                    _ = v82
                }
                var v83 int32 = 1
                _ = fi68
                _ = v69
                _ = v83
            }
            _ = v67
        }
        _ = fi57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v53
    _ = v6
}
