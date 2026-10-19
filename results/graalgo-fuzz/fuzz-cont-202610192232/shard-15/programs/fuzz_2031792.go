package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 uint; F1 int64; F2 bool }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 uint64; F1 string; F2 uint64 }
type GenS11 struct { F0 uint; F1 int8 }
type GenS12 struct { F0 uint16 }
func genFn0() uint16 {
    v7 := []int{9223372036854775807, 927, -9223372036854775808}
    var v8 int8 = -128
    v8 += v8
    for fi9 := 0; fi9 < 14; fi9++ {
        v8 -= v8
        if true {
            if false {
                var v10 float64 = (0.0 + 1e10)
                switch "foo bar" {
                case "\n\t":
                    var v11 uint16 = 65535
                    _ = v11
                case "a":
                    for fi12 := 0; fi12 < 20; fi12++ {
                        fmt.Println(fi9)
                        if len(v7) > 0 {
                            v7[0] = (fi12 - fi9)
                        }
                        _ = fi12
                    }
                }
                _ = len(v7)
                _ = v10
            } else {
                var v13 uint64 = 900
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v13
            }
            v14 := GenS1{F0: 65535, F1: 2}
            if (2 >= 65535) {
                var v15 int8 = (v8 ^ -1)
                v15 -= v15
                for fi16 := 0; fi16 < 4; fi16++ {
                    v15 += 100
                    var v17 bool = (true || false)
                    fmt.Println(fi9)
                    var v18 int8 = 17
                    switch fi9 {
                    case 0:
                        var v19 uint8 = 0
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
            _ = v14
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            if true {
                v21 := GenS2{F0: 42, F1: -174}
                v22 := GenS1{F0: 500, F1: 4294967295}
                if len(v7) > 0 {
                    v7[0] = 1
                }
                _ = v21
                _ = v22
            }
            _ = fi20
        }
        _ = fi9
    }
    _ = v7
    _ = v8
    return 937
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{307, uint32(v1), 4294967295}
    v3 := []string{"a"}
    v4 := make(map[string]bool)
    var v5 int16 = 0
    v6 := make(map[string]bool)
    var v23 float64 = 1e10
    if ((v5 * v5) == v5) {
        for fi24 := 0; fi24 < 5; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                if false {
                    var v26 uint16 = 32767
                    v23 -= v23
                    var v27 uint16 = v26
                    _ = v26
                    _ = v27
                }
                _ = len(v6)
                var v28 uint16 = 536
                _ = fi25
                _ = v28
            }
            fmt.Println(v1, v5, v5)
            _ = fi24
        }
        v5 = 42
        delete(v6, ("hello" + "a"))
    }
    switch "hello" {
    case "a":
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 float32 = 0.1
            _ = fi29
            _ = v30
        }
    case "hello":
        v31 := GenS1{F0: 255, F1: 0}
        _ = v31
    case "":
        var v32 int8 = -76
        _ = v32
    default:
        v5 -= (v5 % (((v5 / (((v5 >> ((42) & 63))) | 1))) | 1))
    }
    var v33 float32 = float32(v23)
    var v34 float32 = 1e10
    _ = len(v6)
    for fi35 := 0; fi35 < 7; fi35++ {
        v36 := GenS3{F0: v1, F1: 1, F2: true}
        v37 := GenS4{F0: v1, F1: (v1 * 0)}
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int16 = -1
    var v39 int32 = 204
    v40 := GenS5{F0: 0, F1: (^(v1))}
    v41 := GenS6{F0: 42, F1: 0}
    for fi42 := 0; fi42 < 20; fi42++ {
        switch "\n\t" {
        case "hello":
            var v43 uint64 = (v1 + v1)
            _ = v43
        case "":
            switch "a" {
            case "\n\t":
                v44 := v6["hello"]
                _ = v44
            case "foo bar":
                v4["foo bar"] = ((v23 + (0.0 - 1.5)) != v23)
            case "hello":
                if len(v2) > 0 {
                    v2[0] = 0
                }
            default:
                fmt.Println(v23, v23)
            }
        default:
            if true {
                switch "" {
                case "foo bar":
                    for fi45 := 0; fi45 < 17; fi45++ {
                        v39 += (v39 % ((v39) | 1))
                        _ = fi45
                    }
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        for fi46 := 0; fi46 < 7; fi46++ {
                            delete(v4, "hello")
                            _ = fi46
                        }
                    case "foo bar":
                        _ = len(v2)
                    case "hello":
                        for fi47 := 0; fi47 < 6; fi47++ {
                            v48 := v6[""]
                            var v49 int = 711
                            var v50 uint16 = 0
                            _ = fi47
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    default:
                        if (!(-920 == 9223372036854775807)) {
                            v38 = v38
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                        }
                    }
                default:
                    fmt.Println(v23, v5)
                }
                if (!false) {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        switch (-(fi51)) {
                        case 32767:
                            _ = len(v3)
                        case 270:
                            v1 = 2
                        case -337:
                            _ = len(v3)
                        default:
                            delete(v4, "a")
                        }
                        v34 = (0.1 * 0.0)
                        _ = fi51
                    }
                    _ = len(v4)
                    for fi52 := 0; fi52 < 9; fi52++ {
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        fmt.Println(v1, fi42, v23)
                        if (!(false && true)) {
                            switch fi52 {
                            case -99:
                                var v53 int8 = 100
                                _ = v53
                            case 9223372036854775807:
                                v54 := v4["\n\t"]
                                _ = v54
                            }
                        } else {
                            v33 += ((0.1 / v33) / v33)
                            var v55 uint = 4294967295
                            var v56 int64 = -935
                            _ = v55
                            _ = v56
                        }
                        _ = fi52
                    }
                } else {
                    v57 := GenS7{F0: "a", F1: (v1 / ((0) | 1))}
                    v33 += (0.1 / v33)
                    v5 = (v38 / ((-888) | 1))
                    _ = v57
                }
            } else {
                var v58 float64 = v23
                for fi59 := 0; fi59 < 12; fi59++ {
                    var v60 float32 = v34
                    v61 := GenS2{F0: -887, F1: -833}
                    if ((1e10 + v58) < v58) {
                        v62 := GenS8{F0: 0, F1: 1, F2: (false && (fi42 >= (fi42 | fi59)))}
                        var v63 uint64 = v1
                        _ = v62
                        _ = v63
                    } else {
                        v60 = v60
                        if false {
                            var v64 uint16 = 255
                            if false {
                                var v65 uint32 = 1
                                v23 -= float64(v34)
                                _ = v65
                            } else {
                                v4["hello"] = true
                                v34 -= v34
                                var v66 uint64 = (v1 + (v1 * v1))
                                _ = v66
                            }
                            _ = v64
                        }
                    }
                    v67 := GenS5{F0: 255, F1: v1}
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v67
                }
                for fi68 := 0; fi68 < 20; fi68++ {
                    v2 = append(v2, 2147483647)
                    v39 += v39
                    _ = fi68
                }
                _ = v58
            }
        }
        _ = fi42
    }
    if (!false) {
        for fi69 := 0; fi69 < 1; fi69++ {
            for fi70 := 0; fi70 < 10; fi70++ {
                if ((v23 * v23) <= (v23 * ((v23 - v23) / 3.14159))) {
                    v71 := GenS6{F0: v38, F1: 42}
                    var v72 int32 = (v39 | (v39 * v39))
                    _ = v71
                    _ = v72
                }
                if true {
                    v73 := GenS1{F0: 4294967295, F1: 758}
                    v23 = v23
                    _ = v73
                } else {
                    if (128 < 87) {
                        _ = len(v6)
                    } else {
                        v1 += v1
                        var v74 float32 = float32((v23 - v23))
                        _ = v74
                    }
                    var v75 uint64 = (v1 >> ((445) & 63))
                    if true {
                        var v76 uint32 = 0
                        var v77 float32 = v33
                        v78 := GenS4{F0: (^(v75)), F1: 35}
                        v79 := GenS9{F0: 0, F1: fi69}
                        _ = v76
                        _ = v77
                        _ = v78
                        _ = v79
                    }
                    _ = v75
                }
                _ = fi70
            }
            switch (fi69 ^ (^(fi69))) {
            case 920:
                if true {
                    v80 := GenS5{F0: 457, F1: (v1 & (^(v1)))}
                    var v81 string = ("hello" + ("a" + ""))
                    for fi82 := 0; fi82 < 1; fi82++ {
                        v38 = v5
                        v83 := GenS10{F0: v1, F1: "\n\t", F2: v1}
                        _ = fi82
                        _ = v83
                    }
                    _ = v80
                    _ = v81
                } else {
                    var v84 string = ("hello" + "")
                    _ = v84
                }
            case -9223372036854775808:
                var v85 int16 = v5
                _ = v85
            default:
                if false {
                    for fi86 := 0; fi86 < 6; fi86++ {
                        v33 = v33
                        var v87 float64 = (v23 / 1.0)
                        _ = fi86
                        _ = v87
                    }
                } else {
                    v88 := GenS11{F0: 1, F1: 100}
                    v33 += v34
                    _ = v88
                }
            }
            for fi89 := 0; fi89 < 16; fi89++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (-(fi69)) {
                case -2147483648:
                    fmt.Println(fi69, fi69)
                case 32767:
                    fmt.Println(v33)
                case -9223372036854775808:
                    fmt.Println(v1, fi69)
                default:
                    v38 = v5
                }
                _ = fi89
            }
            _ = fi69
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    v39 = v39
    v90 := GenS12{F0: genFn0()}
    v39 -= (-(v39))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v90
}
