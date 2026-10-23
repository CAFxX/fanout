package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 bool }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int8; F1 string; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 int32; F1 int8; F2 int16 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 uint16; F1 uint; F2 uint64 }
func genFn0() uint32 {
    v5 := GenS1{F0: -50}
    var v6 int8 = 0
    _ = v5
    _ = v6
    return 242
}

func main() {
    v1 := []uint32{}
    var v2 int16 = -32768
    v3 := make(map[int]bool)
    var v4 uint = 246
    if false {
        fmt.Println(v4)
        switch 4294967295 {
        case 2147483647:
            v7 := GenS2{F0: 2147483647, F1: v2, F2: (233 >= 2147483648)}
            _ = v7
        case 0:
            v8 := v3[267]
            _ = v8
        default:
            var v9 bool = (150 > 168)
            _ = v9
        }
        switch "hello" {
        case "":
            switch -895 {
            case 32768:
                if (!true) {
                    _ = len(v3)
                }
            default:
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS1{F0: 127}
                    var v12 string = ""
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        default:
            v13 := v3[-1]
            _ = v13
        }
    } else {
        if true {
            var v14 int32 = 170
            _ = v14
        } else {
            v4 += v4
            for fi15 := 0; fi15 < 1; fi15++ {
                v2 = (v2 | v2)
                v4 += v4
                _ = fi15
            }
            v16 := GenS3{F0: (v4 + 0), F1: uint64(v4)}
            _ = v16
        }
        fmt.Println(v4, v4)
        fmt.Println(v2)
    }
    var v17 uint16 = 255
    v18 := GenS1{F0: 0}
    v2 = (v2 - v2)
    delete(v3, 4294967295)
    var v19 int64 = -1
    var v20 float32 = 3.14159
    v4 = v4
    v17 = v17
    var v21 int64 = (^(v19))
    switch ("" + "hello") {
    case "\n\t":
        if true {
            v22 := GenS2{F0: genFn0(), F1: (v2 + -151), F2: (true && (true && (108 >= -1)))}
            var v23 float32 = (float32((-2.25 / 1.0)) / v20)
            _ = v22
            _ = v23
        } else {
            fmt.Println(v19)
            v20 += (v20 / (v20 - 3.14159))
            v4 += 1
        }
    case "a":
        switch "\n\t" {
        case "foo bar":
            v24 := GenS1{F0: 70}
            _ = v24
        default:
            switch 2147483647 {
            case 32768:
                if (2 == 2) {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v2 += v2
                        var v26 int = 9223372036854775807
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v19, v2)
                        case "\n\t":
                            for fi27 := 0; fi27 < 13; fi27++ {
                                v28 := GenS4{F0: 100, F1: ((("\n\t" + "") + (("foo bar" + "foo bar") + ("a" + ("foo bar" + "hello")))) + "\n\t"), F2: "hello"}
                                v29 := GenS3{F0: v4, F1: 468}
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        case "hello":
                            switch "a" {
                            case "":
                                var v30 string = "a"
                                _ = v30
                            case "foo bar":
                                switch "hello" {
                                case "hello":
                                    switch ((("hello" + "hello") + "hello") + "") {
                                    case "foo bar":
                                        fmt.Println(v17, v2, fi25)
                                    case "":
                                        if true {
                                            for fi31 := 0; fi31 < 7; fi31++ {
                                                _ = len(v3)
                                                fmt.Println(v26, v2)
                                                _ = fi31
                                            }
                                            if true {
                                                switch ("" + "a") {
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v32 uint = (^(v4))
                                                        _ = v32
                                                    case "":
                                                        var v33 uint64 = 403
                                                        _ = v33
                                                    case "hello":
                                                        v17 -= 922
                                                    }
                                                default:
                                                    var v34 uint8 = 161
                                                    _ = v34
                                                }
                                                v3[v26] = false
                                            }
                                            switch (v26 << ((1) & 63)) {
                                            case 42:
                                                switch "" {
                                                case "a":
                                                    for fi35 := 0; fi35 < 17; fi35++ {
                                                        var v36 uint16 = 965
                                                        var v37 uint32 = 1
                                                        var v38 int16 = 1000
                                                        _ = fi35
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                case "\n\t":
                                                    var v39 int = 741
                                                    _ = v39
                                                case "foo bar":
                                                    for fi40 := 0; fi40 < 5; fi40++ {
                                                        v3[(fi40 << ((v4) & 63))] = (!false)
                                                        v21 = v21
                                                        if (((1.5 / (1.5 * (((3.14159 / 3.14159) + 0.0) - float64(float32(3.14159))))) <= 1.5) && true) {
                                                            var v41 uint16 = (v17 + (v17 % (((^(v17))) | 1)))
                                                            _ = v41
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = 65535
                                                        }
                                                        var v42 uint64 = 2
                                                        fmt.Println(v21, v20)
                                                        _ = fi40
                                                        _ = v42
                                                    }
                                                }
                                            case -9223372036854775808:
                                                var v43 int8 = -128
                                                _ = v43
                                            default:
                                                var v44 uint16 = v17
                                                _ = v44
                                            }
                                        }
                                    }
                                case "":
                                    fmt.Println(v19)
                                case "foo bar":
                                    v45 := GenS3{F0: v4, F1: 42}
                                    _ = v45
                                default:
                                    var v46 string = "foo bar"
                                    _ = v46
                                }
                            case "\n\t":
                                for fi47 := 0; fi47 < 17; fi47++ {
                                    var v48 uint32 = 0
                                    v49 := GenS5{F0: (v19 ^ 9223372036854775807)}
                                    v21 += (v19 & (^(v21)))
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            default:
                                v26 += fi25
                            }
                        default:
                            if false {
                                if (1.5 >= 0.0) {
                                    var v50 int8 = 42
                                    _ = v50
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                fmt.Println(v20, v21)
                            } else {
                                v21 = (v21 & v19)
                                switch (("\n\t" + "") + ("hello" + "\n\t")) {
                                case "\n\t":
                                    for fi51 := 0; fi51 < 10; fi51++ {
                                        v52 := GenS6{F0: fi51, F1: fi25}
                                        _ = fi51
                                        _ = v52
                                    }
                                case "":
                                    var v53 uint = 42
                                    _ = v53
                                case "foo bar":
                                    if true {
                                        if false {
                                            if ((!((v19 >= (v19 & v21)) || (2147483648 >= 0))) || false) {
                                                if len(v1) > 0 {
                                                    v1[0] = 42
                                                }
                                                v4 += v4
                                                var v54 uint16 = v17
                                                _ = v54
                                            }
                                            v55 := GenS2{F0: 1, F1: v2, F2: (false && false)}
                                            _ = v55
                                        }
                                        for fi56 := 0; fi56 < 4; fi56++ {
                                            switch (fi56 % ((v26) | 1)) {
                                            case 471:
                                                if ((true && (v17 != (v17 & (v17 * v17)))) && (2147483647 != 179)) {
                                                    v2 = 302
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v57 bool = false
                                                    _ = v57
                                                }
                                            default:
                                                v4 = 125
                                            }
                                            var v58 int32 = -738
                                            var v59 bool = true
                                            _ = fi56
                                            _ = v58
                                            _ = v59
                                        }
                                        for fi60 := 0; fi60 < 1; fi60++ {
                                            var v61 int = (fi25 ^ fi25)
                                            if ((false || (127 >= 0)) && true) {
                                                fmt.Println(v20)
                                                v4 += (v4 + v4)
                                            } else {
                                                var v62 uint32 = 720
                                                _ = len(v1)
                                                _ = v62
                                            }
                                            _ = fi60
                                            _ = v61
                                        }
                                    } else {
                                        v63 := GenS3{F0: v4, F1: 616}
                                        var v64 uint64 = 1
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                v17 = v17
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v65 bool = false
                        _ = v65
                    case "\n\t":
                        _ = len(v1)
                    case "":
                        v66 := GenS7{F0: 1, F1: 1, F2: v2}
                        _ = v66
                    default:
                        v67 := GenS8{F0: 0, F1: 42}
                        _ = v67
                    }
                    v19 = (v21 * 66)
                }
            }
        }
    }
    switch -667 {
    case 32767:
        fmt.Println(v21)
    case -1:
        if true {
            switch -2147483648 {
            case 1:
                _ = len(v1)
            case 32767:
                var v68 int32 = 0
                _ = v68
            default:
                if true {
                    fmt.Println(v2)
                    for fi69 := 0; fi69 < 16; fi69++ {
                        var v70 uint16 = (v17 / ((v17) | 1))
                        _ = fi69
                        _ = v70
                    }
                    var v71 uint32 = 2
                    _ = v71
                }
            }
            v20 = (((0.0 - (0.1 / v20)) / (1e10 + v20)) / (v20 * 1e10))
        }
    case 32768:
        _ = len(v1)
    default:
        v72 := GenS7{F0: 2147483647, F1: 42, F2: v2}
        _ = v72
    }
    v2 += 1000
    v73 := GenS9{F0: (v17 << ((42) & 63)), F1: v4, F2: 42}
    fmt.Println(v4, v4, v19)
    var v74 uint16 = (v17 & 0)
    for fi75 := 0; fi75 < 16; fi75++ {
        fmt.Println(v20, v21, v17)
        switch (^(fi75)) {
        case 32768:
            fmt.Println(v17)
        default:
            for fi76 := 0; fi76 < 6; fi76++ {
                if (!(42 >= v17)) {
                    v19 += v19
                    var v77 uint64 = 42
                    _ = v77
                } else {
                    v21 = v19
                }
                _ = fi76
            }
        }
        v17 = 65535
        _ = fi75
    }
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v73
    _ = v74
}
