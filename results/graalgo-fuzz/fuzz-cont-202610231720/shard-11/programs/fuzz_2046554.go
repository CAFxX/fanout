package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 int32; F1 int8; F2 int32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -32768
    v3 := make(map[uint32]int)
    v4 := []string{"a", "", "", "a"}
    var v5 int32 = 407
    var v6 float32 = ((3.14159 + (0.1 / 1e10)) / ((1.5 * (1e10 * 1e10)) * float32((1e10 + 3.14159))))
    switch "a" {
    case "\n\t":
        v6 -= (1e10 + v6)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        if (-1 <= 1000) {
            v8 := GenS1{F0: 42, F1: -9223372036854775808}
            v9 := GenS2{F0: 27, F1: 9223372036854775807, F2: 65535}
            var v10 uint16 = 2
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    switch 127 {
    case 32767:
        if false {
            for fi11 := 0; fi11 < 18; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v2 += 1
                    var v13 float32 = (v6 + v6)
                    _ = fi12
                    _ = v13
                }
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                _ = fi11
            }
        } else {
            if false {
                v5 += (v5 ^ (v5 / ((0) | 1)))
            }
            v5 = -1
            var v14 uint16 = 897
            _ = v14
        }
    case 4294967295:
        var v15 int8 = 1
        _ = v15
    case 1:
        if (!(v6 <= 1e10)) {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 bool = ((!(127 <= -128)) && true)
                v17 = v17
                var v18 int = fi16
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    default:
        var v19 int64 = 915
        _ = v19
    }
    if false {
        v5 += (v5 ^ -1)
        var v20 uint = 2
        fmt.Println(v5, v20, v6)
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        if (1 == 4294967295) {
            if len(v4) > 0 {
                v4[0] = ""
            }
            delete(v1, (^(fi21)))
            for fi22 := 0; fi22 < 16; fi22++ {
                fmt.Println(fi21, v5)
                v23 := GenS1{F0: fi21, F1: (fi21 - 9223372036854775807)}
                _ = fi22
                _ = v23
            }
        }
        _ = fi21
    }
    switch int(v2) {
    case 42:
        if false {
            if (24 >= 1) {
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 uint64 = 4294967295
                    _ = fi24
                    _ = v25
                }
                v2 += 1
            }
            v5 += v5
            v2 -= (v2 & 32767)
        }
    default:
        v26 := GenS2{F0: 255, F1: int(v5), F2: 74}
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        fmt.Println(v2)
        v28 := GenS3{F0: v5, F1: 67, F2: v5}
        _ = fi27
        _ = v28
    }
    v1[499] = (v2 < (-(v2)))
    v29 := GenS3{F0: (v5 << ((18446744073709551615) & 63)), F1: -1, F2: v5}
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        var v30 uint = 1
        _ = v30
    case "a":
        for fi31 := 0; fi31 < 7; fi31++ {
            if false {
                if (9223372036854775807 == (^(fi31))) {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 string = "hello"
                        v2 = v2
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch (v33 + v33) {
                            case "foo bar":
                                v6 -= 0.1
                            case "hello":
                                switch v33 {
                                case "foo bar":
                                    if (1e10 == -2.25) {
                                        if (float32(((0.0 * ((-2.25 / 1.0) + (0.0 * (-2.25 * 1.5)))) - (0.0 / 1e10))) != ((v6 / v6) * 3.14159)) {
                                            fmt.Println(fi32, v2, v2)
                                            if false {
                                                var v34 string = v33
                                                _ = v34
                                            } else {
                                                fmt.Println(fi31)
                                                v2 = int16(fi32)
                                                var v35 uint32 = 0
                                                _ = v35
                                            }
                                        } else {
                                            var v36 uint64 = 0
                                            switch -373 {
                                            case 128:
                                                if (!true) {
                                                    v5 -= (v5 >> ((255) & 63))
                                                    var v37 int8 = 127
                                                    _ = v37
                                                }
                                            }
                                            var v38 int8 = -46
                                            _ = v36
                                            _ = v38
                                        }
                                        var v39 int32 = (v5 + -818)
                                        _ = len(v3)
                                        _ = v39
                                    } else {
                                        v40 := GenS4{F0: 255, F1: -808, F2: -1}
                                        v41 := GenS4{F0: 245, F1: -421, F2: 42}
                                        _ = v40
                                        _ = v41
                                    }
                                case "":
                                    var v42 float32 = v6
                                    _ = v42
                                case "a":
                                    v43 := GenS2{F0: 665, F1: 2147483647, F2: 2147483647}
                                    _ = v43
                                }
                            case "\n\t":
                                v33 = v33
                                if ((42 > 11) || false) {
                                    v44 := GenS4{F0: 840, F1: 746, F2: 0}
                                    _ = v44
                                } else {
                                    if (!(int64(v2) <= 42)) {
                                        fmt.Println(fi31, v2, v5)
                                        v4 = append(v4, v33)
                                        var v45 int16 = (v2 / ((v2) | 1))
                                        _ = v45
                                    } else {
                                        switch v33 {
                                        case "":
                                            var v46 int32 = -642
                                            _ = v46
                                        case "\n\t":
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = ("\n\t" + (v33 + v33))
                                                }
                                                v6 += float32(1.5)
                                            }
                                        }
                                        var v47 int16 = (^(v2))
                                        v47 = (v47 & (v2 >> ((255) & 63)))
                                        _ = v47
                                    }
                                }
                                var v48 uint64 = 18446744073709551615
                                _ = v48
                            default:
                                if false {
                                    if true {
                                        var v49 float32 = 0.1
                                        var v50 float64 = -2.25
                                        v6 = (v49 * v6)
                                        _ = v49
                                        _ = v50
                                    }
                                    fmt.Println(v2, v5)
                                } else {
                                    v2 = 0
                                    v2 = (-(v2))
                                    _ = len(v4)
                                    var v51 int8 = 31
                                    _ = v51
                                }
                            }
                            var v52 uint = 723
                            _ = v52
                        } else {
                            if false {
                                v33 = (v33 + v33)
                            }
                            var v53 string = "\n\t"
                            _ = v53
                        }
                        if len(v4) > 0 {
                            v4[0] = ("\n\t" + ("foo bar" + v33))
                        }
                        var v54 int8 = 100
                        var v55 uint32 = 0
                        _ = fi32
                        _ = v33
                        _ = v54
                        _ = v55
                    }
                } else {
                    if false {
                        var v56 uint32 = 793
                        v57 := GenS1{F0: 0, F1: 2147483647}
                        v56 += (v56 % ((2) | 1))
                        _ = v56
                        _ = v57
                    } else {
                        var v58 int16 = (^(v2))
                        v58 += 32767
                        fmt.Println(v5, v2, v6)
                        _ = v58
                    }
                    fmt.Println(v2)
                }
            }
            v2 -= v2
            for fi59 := 0; fi59 < 4; fi59++ {
                var v60 string = "\n\t"
                v1[(fi59 ^ (-(fi59)))] = (true && true)
                v61 := GenS5{F0: 2, F1: 1}
                _ = fi59
                _ = v60
                _ = v61
            }
            _ = fi31
        }
    }
    for fi62 := 0; fi62 < 2; fi62++ {
        v63 := GenS4{F0: 32767, F1: 42, F2: -1}
        _ = fi62
        _ = v63
    }
    var v64 string = "foo bar"
    if false {
        var v65 int = 32768
        _ = v65
    } else {
        fmt.Println(v6, v2)
        switch 127 {
        case 2147483647:
            var v66 int8 = -47
            _ = v66
        case -1:
            v64 = v64
        }
        v67 := GenS2{F0: 507, F1: -99, F2: 1}
        _ = v67
    }
    switch "foo bar" {
    case "foo bar":
        v6 -= float32(((0.0 + 3.14159) + 3.14159))
    case "":
        v68 := GenS5{F0: 551, F1: 2}
        _ = v68
    }
    v69 := GenS4{F0: 1, F1: 1, F2: -1}
    var v70 bool = ((32767 <= 128) && (false && false))
    v71 := GenS2{F0: 42, F1: -259, F2: 954}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v69
    _ = v70
    _ = v71
}
