package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 bool) uint {
    if p0 {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 int64 = -112
            var v7 float64 = 0.1
            _ = fi5
            _ = v6
            _ = v7
        }
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
        var v8 uint = 4294967295
        v9 := []int{1, 32768}
        _ = v8
        _ = v9
    }
    var v10 int32 = -852
    _ = p0
    _ = v10
    return 894
}
func genFn1(p0 uint16, p1 int64) uint16 {
    v11 := GenS1{F0: 2, F1: 1, F2: (128 != 187)}
    var v12 int16 = 32767
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]string)
    var v3 uint64 = 623
    var v4 float32 = (1.5 / 1e10)
    fmt.Println(v4, v3)
    switch 4294967295 {
    case -506:
        if ((v1 << ((1) & 63)) >= v1) {
            if true {
                if true {
                    if false {
                        v4 = (0.1 * v4)
                        var v13 uint32 = 617
                        v1 = v1
                        _ = v13
                    }
                    _ = len(v2)
                    var v14 int64 = (v1 / (((v1 * (v1 + 1))) | 1))
                    _ = v14
                } else {
                    v3 += v3
                    v15 := GenS1{F0: 0, F1: 4294967295, F2: (-32768 >= 1000)}
                    _ = v15
                }
            }
            switch "foo bar" {
            case "foo bar":
                var v16 int64 = v1
                _ = v16
            }
        } else {
            v4 -= v4
            if false {
                var v17 bool = (!true)
                _ = v17
            }
        }
    case 1:
        v3 = 18446744073709551615
    case -2147483648:
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: genFn0(false), F1: 2147483648, F2: ((true || ((-(v1)) < v1)) && true)}
            _ = fi18
            _ = v19
        }
    }
    var v20 float32 = 1e10
    var v21 uint = 42
    v22 := []int{}
    _ = len(v22)
    if (!false) {
        for fi23 := 0; fi23 < 20; fi23++ {
            for fi24 := 0; fi24 < 11; fi24++ {
                fmt.Println(v21, v1, v3)
                v20 += (v20 / v4)
                _ = fi24
            }
            v20 -= float32(0.1)
            _ = fi23
        }
        if false {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 int16 = 1
                fmt.Println(v20, fi25)
                switch "" {
                case "":
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 int64 = -9223372036854775808
                        v26 = 1000
                        _ = fi27
                        _ = v28
                    }
                case "a":
                    var v29 int16 = (v26 >> (((v21 % (((v21 & 18446744073709551615)) | 1))) & 63))
                    _ = v29
                case "foo bar":
                    var v30 uint64 = (v3 / (((v3 - (v3 - (^(v3))))) | 1))
                    _ = v30
                }
                _ = fi25
                _ = v26
            }
            fmt.Println(v3)
            if (4294967295 == 0) {
                switch -2147483648 {
                case -1:
                    fmt.Println(v4)
                case 9:
                    for fi31 := 0; fi31 < 12; fi31++ {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            var v33 float32 = (0.1 * float32(0.0))
                            var v34 uint = v21
                            fmt.Println(v34, v1)
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        if true {
                            if ((3.14159 / v4) < v4) {
                                switch -210 {
                                case 348:
                                    v20 = v4
                                case 741:
                                    for fi35 := 0; fi35 < 18; fi35++ {
                                        if true {
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                            v36 := GenS2{F0: 160}
                                            _ = v36
                                        }
                                        v4 += v20
                                        var v37 float32 = (float32(3.14159) / (1e10 - 0.1))
                                        _ = fi35
                                        _ = v37
                                    }
                                case 137:
                                    delete(v2, 320)
                                default:
                                    v20 = (float32(1e10) / ((v4 / 1.5) * 1e10))
                                }
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    var v39 uint = genFn0((-2.25 > 0.1))
                                    if (!(false || false)) {
                                        switch (fi38 | 9223372036854775807) {
                                        case -1:
                                            var v40 int8 = 123
                                            _ = v40
                                        case 2147483647:
                                            v3 = 18446744073709551615
                                        default:
                                            v41 := GenS1{F0: (v21 ^ v39), F1: 0, F2: ((!(!(3.14159 != 1.5))) || (!false))}
                                            _ = v41
                                        }
                                        var v42 float32 = v20
                                        var v43 int16 = 42
                                        _ = v42
                                        _ = v43
                                    }
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            if (true && true) {
                                v4 -= ((v20 - v20) + ((((0.1 * v4) + float32((-2.25 + -2.25))) / (1e10 - v4)) + (1.5 / v4)))
                                v21 = 18446744073709551615
                            }
                        }
                        _ = fi31
                    }
                }
                if false {
                    var v44 uint32 = 374
                    switch "a" {
                    case "hello":
                        if false {
                            if (v44 >= (v44 + v44)) {
                                for fi45 := 0; fi45 < 13; fi45++ {
                                    if false {
                                        var v46 float64 = (float64((v4 - 1e10)) + 3.14159)
                                        var v47 float32 = v20
                                        v21 += v21
                                        _ = v46
                                        _ = v47
                                    }
                                    switch ("foo bar" + "foo bar") {
                                    case "hello":
                                        v3 = 750
                                    default:
                                        v1 += (v1 - v1)
                                    }
                                    _ = fi45
                                }
                                v3 -= v3
                                v48 := GenS1{F0: v21, F1: (v44 % (((v44 * v44)) | 1)), F2: true}
                                _ = v48
                            } else {
                                var v49 int = 560
                                for fi50 := 0; fi50 < 4; fi50++ {
                                    v22 = append(v22, v49)
                                    for fi51 := 0; fi51 < 16; fi51++ {
                                        v52 := GenS1{F0: (v21 - v21), F1: v44, F2: false}
                                        v21 = v21
                                        _ = fi51
                                        _ = v52
                                    }
                                    _ = fi50
                                }
                                var v53 string = "\n\t"
                                _ = v49
                                _ = v53
                            }
                        }
                    case "":
                        var v54 uint = (v21 | genFn0(true))
                        _ = v54
                    default:
                        v22 = append(v22, 1)
                    }
                    v4 += (3.14159 * -2.25)
                    _ = v44
                }
                switch "" {
                case "\n\t":
                    var v55 bool = (v4 == float32(3.14159))
                    _ = v55
                case "":
                    var v56 int16 = 42
                    _ = v56
                case "a":
                    var v57 uint8 = 2
                    _ = v57
                default:
                    var v58 uint32 = 2147483647
                    _ = v58
                }
            }
        } else {
            if (809 < genFn0(false)) {
                for fi59 := 0; fi59 < 16; fi59++ {
                    fmt.Println(v1)
                    _ = fi59
                }
                if true {
                    if len(v22) > 0 {
                        v22[0] = 4294967295
                    }
                    v4 -= (((v4 / float32(-2.25)) * v20) + (3.14159 - float32(3.14159)))
                    _ = len(v2)
                } else {
                    v3 += v3
                    switch ("hello" + "\n\t") {
                    case "a":
                        if (1 == -280) {
                            v3 = (^(v3))
                        }
                    case "":
                        var v60 int = 32768
                        _ = v60
                    case "hello":
                        delete(v2, 716)
                    default:
                        v3 = v3
                    }
                }
                var v61 uint8 = 227
                _ = v61
            } else {
                _ = len(v2)
            }
            v62 := GenS1{F0: v21, F1: 2147483648, F2: (!false)}
            var v63 uint16 = 531
            _ = v62
            _ = v63
        }
        if (false && true) {
            fmt.Println(v1, v20)
            v64 := GenS3{F0: 275, F1: -946, F2: v3}
            var v65 int8 = 1
            _ = v64
            _ = v65
        } else {
            var v66 float32 = ((1.5 / v4) - v20)
            _ = v66
        }
    } else {
        switch ("a" + "a") {
        case "hello":
            delete(v2, 2)
        default:
            var v67 uint64 = v3
            _ = v67
        }
        if true {
            switch "\n\t" {
            case "a":
                v2[2147483648] = "hello"
            default:
                switch -1 {
                case -33:
                    for fi68 := 0; fi68 < 11; fi68++ {
                        v69 := GenS4{F0: 98}
                        _ = fi68
                        _ = v69
                    }
                case 32767:
                    var v70 uint = (^(v21))
                    _ = v70
                }
            }
        }
    }
    for fi71 := 0; fi71 < 14; fi71++ {
        v72 := GenS3{F0: 42, F1: 1000, F2: v3}
        v20 -= ((v20 / -2.25) - 1.5)
        _ = fi71
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
