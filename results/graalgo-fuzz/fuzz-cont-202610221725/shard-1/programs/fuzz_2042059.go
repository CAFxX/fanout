package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 bool; F2 uint }
type GenS7 struct { F0 string; F1 uint; F2 int64 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS10 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (0 < 598)
    var v3 int16 = -216
    v4 := []int{9223372036854775807, -1, -208, -928}
    v5 := make(map[string]string)
    var v6 uint = 643
    if (3.14159 == (0.0 - 1e10)) {
        switch ("" + ("a" + "\n\t")) {
        case "a":
            for fi7 := 0; fi7 < 11; fi7++ {
                switch 1 {
                case 127:
                    if (("a" + "a") <= (("a" + "foo bar") + "a")) {
                        for fi8 := 0; fi8 < 15; fi8++ {
                            var v9 uint64 = 0
                            _ = fi8
                            _ = v9
                        }
                    } else {
                        v6 -= v6
                        for fi10 := 0; fi10 < 12; fi10++ {
                            v11 := GenS1{F0: 100, F1: 132, F2: 0}
                            _ = fi10
                            _ = v11
                        }
                    }
                }
                fmt.Println(v6, v6)
                _ = fi7
            }
        case "foo bar":
            v12 := GenS2{F0: 263}
            _ = v12
        }
        if v2 {
            var v13 int8 = 127
            _ = v13
        } else {
            v3 -= v3
            v3 = v3
            for fi14 := 0; fi14 < 13; fi14++ {
                v6 += (v6 - (^(v6)))
                switch "a" {
                case "a":
                    var v15 int16 = 42
                    _ = v15
                default:
                    if v2 {
                        for fi16 := 0; fi16 < 8; fi16++ {
                            v6 = 2
                            var v17 float64 = 1e10
                            fmt.Println(fi16, v6)
                            var v18 uint = (v6 | 0)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                }
                _ = fi14
            }
        }
        if (114 == 65535) {
            var v19 int = 1
            for fi20 := 0; fi20 < 11; fi20++ {
                switch (v19 ^ fi20) {
                case -9223372036854775808:
                    if false {
                        _ = len(v5)
                        v21 := GenS1{F0: 100, F1: 2, F2: -1}
                        _ = v21
                    } else {
                        var v22 int16 = v3
                        fmt.Println(v2)
                        var v23 int32 = 1
                        _ = v22
                        _ = v23
                    }
                case 196:
                    fmt.Println(v6)
                case 1:
                    fmt.Println(fi20, v6)
                default:
                    for fi24 := 0; fi24 < 2; fi24++ {
                        switch 869 {
                        case 0:
                            v19 = v19
                        default:
                            v25 := GenS3{F0: 42}
                            _ = v25
                        }
                        v26 := GenS3{F0: 42}
                        switch (-(fi24)) {
                        case 1:
                            if v2 {
                                if false {
                                    v27 := GenS2{F0: 1}
                                    var v28 string = "a"
                                    if (42 <= -1) {
                                        v29 := GenS1{F0: 93, F1: uint8(v6), F2: -235}
                                        var v30 float64 = 1e10
                                        var v31 string = (v28 + v28)
                                        _ = len(v4)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = len(v5)
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                    _ = v27
                                    _ = v28
                                } else {
                                    v32 := GenS4{F0: 580}
                                    _ = v32
                                }
                                switch "a" {
                                case "":
                                    v33 := GenS5{F0: fi20}
                                    _ = v33
                                case "a":
                                    var v34 uint = v6
                                    _ = v34
                                case "hello":
                                    var v35 int32 = -488
                                    _ = v35
                                default:
                                    v36 := GenS4{F0: 42}
                                    _ = v36
                                }
                            } else {
                                v19 += (v19 << (((v6 << ((v6) & 63))) & 63))
                                v6 -= 255
                            }
                        case 32768:
                            var v37 uint32 = 42
                            _ = v37
                        case -9223372036854775808:
                            var v38 int8 = 42
                            _ = v38
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v39 uint8 = 0
                        v40 := GenS2{F0: 4294967295}
                        _ = fi24
                        _ = v26
                        _ = v39
                        _ = v40
                    }
                }
                _ = fi20
            }
            for fi41 := 0; fi41 < 6; fi41++ {
                v42 := GenS5{F0: 2147483647}
                var v43 int8 = -128
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = v19
        } else {
            var v44 uint = (v6 | v6)
            if (false || (((551 > 131) || (-658 > 100)) && v2)) {
                v45 := GenS6{F0: 0, F1: (4294967295 != 2147483648), F2: (v6 | v6)}
                for fi46 := 0; fi46 < 3; fi46++ {
                    var v47 string = "foo bar"
                    v48 := GenS7{F0: "", F1: 0, F2: -690}
                    v6 -= 601
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = v45
            }
            v2 = true
            _ = v44
        }
    } else {
        switch "" {
        case "foo bar":
            fmt.Println(v6, v3, v3)
        case "a":
            v6 = (v6 % ((v6) | 1))
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = -9223372036854775808
            }
        default:
            v49 := GenS4{F0: 42}
            _ = v49
        }
        switch "foo bar" {
        case "hello":
            v50 := GenS8{F0: 444, F1: -100}
            _ = v50
        case "a":
            v51 := GenS3{F0: 691}
            _ = v51
        default:
            switch ("hello" + "foo bar") {
            case "":
                switch "" {
                case "a":
                    v2 = v2
                default:
                    var v52 float64 = (0.1 - -2.25)
                    _ = v52
                }
            case "a":
                v6 = (v6 + 572)
            }
        }
    }
    var v53 bool = (true || false)
    v54 := GenS9{F0: 206, F1: 1, F2: 2}
    v53 = v2
    var v55 string = ("foo bar" + "a")
    switch 2147483647 {
    case 32767:
        for fi56 := 0; fi56 < 14; fi56++ {
            v2 = false
            v6 = (^(v6))
            fmt.Println(v2, fi56, v53)
            _ = fi56
        }
    case -2147483648:
        for fi57 := 0; fi57 < 5; fi57++ {
            fmt.Println(v2)
            var v58 int32 = -2147483648
            _ = fi57
            _ = v58
        }
    case 2147483647:
        var v59 int16 = v3
        _ = v59
    }
    var v60 float64 = 3.14159
    for fi61 := 0; fi61 < 1; fi61++ {
        if len(v4) > 0 {
            v4[0] = (fi61 | (fi61 + fi61))
        }
        fmt.Println(v6)
        var v62 uint = 1
        _ = fi61
        _ = v62
    }
    v55 = ""
    v53 = (32767 == 1)
    fmt.Println(v6, v60, v3)
    for fi63 := 0; fi63 < 4; fi63++ {
        v6 -= (^(v6))
        v64 := GenS10{F0: 223}
        _ = fi63
        _ = v64
    }
    if ((v3 * (v3 / (((v3 / ((v3) | 1))) | 1))) == v3) {
        var v65 bool = false
        _ = v65
    } else {
        v66 := GenS8{F0: 32767, F1: 9223372036854775807}
        var v67 float64 = v60
        v55 = v55
        _ = v66
        _ = v67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v6
    _ = v60
}
