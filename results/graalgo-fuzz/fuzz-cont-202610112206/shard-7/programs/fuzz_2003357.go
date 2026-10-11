package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int; F1 int }

func main() {
    var v1 int8 = 82
    var v2 uint64 = 0
    var v3 uint32 = 0
    v4 := make(map[uint32]string)
    var v5 int64 = 9223372036854775807
    v6 := GenS1{F0: 860, F1: 128, F2: 174}
    v2 = v2
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS2{F0: v2}
            _ = fi8
            _ = v9
        }
        v10 := []int{fi7, fi7}
        _ = fi7
        _ = v10
    }
    if ((156 != 179) || true) {
        v11 := GenS3{F0: 4294967295}
        var v12 uint = 255
        _ = v11
        _ = v12
    }
    v13 := []int{-552, -1, 42}
    var v14 int16 = -32768
    if true {
        var v15 float64 = (0.1 - 0.0)
        switch -1 {
        case -1:
            for fi16 := 0; fi16 < 17; fi16++ {
                if ((v2 << ((0) & 63)) < 4294967295) {
                    fmt.Println(v2, v1, v1)
                    if true {
                        var v17 int = (-(fi16))
                        for fi18 := 0; fi18 < 10; fi18++ {
                            if true {
                                if false {
                                    if (32767 == uint16(v3)) {
                                        fmt.Println(v1, v17, v14)
                                        if len(v13) > 0 {
                                            v13[0] = fi18
                                        }
                                    } else {
                                        v17 -= 42
                                    }
                                }
                                if (v5 == v5) {
                                    v19 := GenS4{F0: "a", F1: 127}
                                    if true {
                                        v20 := v4[(^(v3))]
                                        fmt.Println(v17, v5, fi16)
                                        _ = v20
                                    }
                                    v1 -= (v1 % (((v1 | v1)) | 1))
                                    _ = v19
                                } else {
                                    var v21 int = v17
                                    _ = v21
                                }
                            }
                            v15 -= (v15 - 3.14159)
                            _ = fi18
                        }
                        var v22 uint8 = 30
                        _ = v17
                        _ = v22
                    } else {
                        var v23 uint8 = 213
                        _ = v23
                    }
                    switch (fi16 * (fi16 - fi16)) {
                    case 127:
                        if false {
                            for fi24 := 0; fi24 < 15; fi24++ {
                                v15 -= v15
                                var v25 uint8 = 255
                                _ = fi24
                                _ = v25
                            }
                            v4[(v3 | v3)] = "foo bar"
                        } else {
                            fmt.Println(v15)
                            for fi26 := 0; fi26 < 3; fi26++ {
                                var v27 int = (fi16 * -152)
                                _ = fi26
                                _ = v27
                            }
                        }
                    case 101:
                        if len(v13) > 0 {
                            v13[0] = int(v3)
                        }
                    }
                }
                if true {
                    if true {
                        var v28 bool = ((true || (true || false)) && true)
                        _ = v28
                    }
                }
                _ = fi16
            }
        case 32768:
            v1 = -1
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            v2 = (v2 ^ v2)
            v1 = 0
            _ = fi29
        }
        _ = v15
    } else {
        v30 := GenS5{F0: 2}
        if true {
            if true {
                if true {
                    if (!false) {
                        v31 := GenS2{F0: v2}
                        fmt.Println(v14, v2, v2)
                        v32 := GenS3{F0: (v3 - 2147483647)}
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS5{F0: 255}
                    _ = v33
                } else {
                    if (100 < v1) {
                        v3 += v3
                    }
                }
                var v34 uint16 = uint16(v3)
                var v35 uint16 = v34
                _ = v34
                _ = v35
            } else {
                switch ("a" + "foo bar") {
                case "\n\t":
                    switch 545 {
                    case 128:
                        if true {
                            v36 := GenS6{F0: 98}
                            var v37 int = -2147483648
                            _ = v36
                            _ = v37
                        } else {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        }
                    default:
                        var v38 uint16 = 32767
                        _ = v38
                    }
                case "hello":
                    delete(v4, (v3 / ((2) | 1)))
                }
                v14 += v14
            }
            v5 -= int64(v1)
            if true {
                if (56 <= 127) {
                    var v39 string = ("\n\t" + "hello")
                    _ = v39
                } else {
                    switch ((("foo bar" + "\n\t") + "") + "foo bar") {
                    case "foo bar":
                        var v40 uint = 778
                        _ = v40
                    case "a":
                        v4[1] = ""
                    case "hello":
                        if false {
                            var v41 int16 = v14
                            v3 += v3
                            _ = v41
                        } else {
                            var v42 bool = true
                            v43 := GenS7{F0: 32768, F1: 127}
                            v44 := v4[v3]
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    default:
                        var v45 bool = ((!true) && false)
                        _ = v45
                    }
                }
                for fi46 := 0; fi46 < 10; fi46++ {
                    for fi47 := 0; fi47 < 8; fi47++ {
                        var v48 float32 = (0.1 + -2.25)
                        var v49 float64 = -2.25
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    _ = fi46
                }
                v14 += (-(v14))
            }
        } else {
            for fi50 := 0; fi50 < 18; fi50++ {
                switch "foo bar" {
                case "\n\t":
                    v2 = 4294967295
                case "hello":
                    fmt.Println(fi50, fi50, v5)
                default:
                    fmt.Println(v14)
                }
                var v51 int32 = 42
                v52 := GenS7{F0: (fi50 >> ((1) & 63)), F1: 42}
                _ = fi50
                _ = v51
                _ = v52
            }
            var v53 int = 608
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v53
        }
        if true {
            _ = len(v4)
            for fi54 := 0; fi54 < 2; fi54++ {
                v55 := GenS3{F0: 65535}
                fmt.Println(v3)
                var v56 int16 = (v14 << ((18446744073709551615) & 63))
                _ = fi54
                _ = v55
                _ = v56
            }
            _ = len(v4)
        }
        _ = v30
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        for fi58 := 0; fi58 < 15; fi58++ {
            var v59 uint16 = 387
            v5 = (v5 ^ (v5 - v5))
            _ = fi58
            _ = v59
        }
        _ = fi57
    }
    fmt.Println(v14, v2)
    if (!true) {
        v1 = 42
    }
    if (771 >= -1) {
        v14 = (-(v14))
    } else {
        for fi60 := 0; fi60 < 3; fi60++ {
            if len(v13) > 0 {
                v13[0] = 1
            }
            fmt.Println(v5)
            v4[uint32(v2)] = "hello"
            _ = fi60
        }
        if true {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
        var v61 uint16 = 65535
        _ = v61
    }
    if (837 > 0) {
        delete(v4, uint32(v2))
        if (10 > (v5 ^ v5)) {
            var v62 int16 = (v14 % (((v14 - (v14 << ((18446744073709551615) & 63)))) | 1))
            v4[(v3 & (^(v3)))] = "\n\t"
            _ = v62
        } else {
            var v63 uint32 = v3
            var v64 float32 = 0.1
            _ = v63
            _ = v64
        }
    } else {
        if (false || false) {
            var v65 uint32 = 4294967295
            _ = v65
        } else {
            if (765 < 255) {
                var v66 int32 = 2147483647
                switch 127 {
                case 42:
                    v2 = v2
                case -477:
                    if true {
                        _ = len(v13)
                        var v67 float32 = 0.0
                        v68 := GenS7{F0: 2147483647, F1: 255}
                        _ = v67
                        _ = v68
                    } else {
                        switch ("a" + ("\n\t" + (("a" + "foo bar") + "foo bar"))) {
                        case "\n\t":
                            var v69 float64 = 0.1
                            _ = v69
                        case "foo bar":
                            _ = len(v4)
                        default:
                            v5 = (v5 ^ 47)
                        }
                        var v70 int = 128
                        _ = v70
                    }
                case 1:
                    v1 = 127
                }
                v1 = (v1 >> ((65535) & 63))
                _ = v66
            }
            v5 -= v5
        }
        if true {
            for fi71 := 0; fi71 < 8; fi71++ {
                for fi72 := 0; fi72 < 12; fi72++ {
                    v1 = (v1 - (v1 >> ((403) & 63)))
                    switch "a" {
                    case "foo bar":
                        var v73 float32 = 1e10
                        _ = v73
                    }
                    _ = fi72
                }
                fmt.Println(v1, v1, v14)
                var v74 int32 = 0
                _ = fi71
                _ = v74
            }
        } else {
            var v75 string = "\n\t"
            _ = v75
        }
        v76 := GenS8{F0: 32767, F1: -237}
        _ = v76
    }
    for fi77 := 0; fi77 < 4; fi77++ {
        v78 := GenS3{F0: (v3 ^ v3)}
        v4[v3] = ""
        for fi79 := 0; fi79 < 7; fi79++ {
            v80 := GenS8{F0: (fi77 | (fi77 - (-(fi77)))), F1: (fi79 | int(v3))}
            _ = fi79
            _ = v80
        }
        _ = fi77
        _ = v78
    }
    v81 := GenS6{F0: 2}
    var v82 int64 = (-(v5))
    _ = len(v4)
    v83 := v4[v3]
    switch v83 {
    case "a":
        for fi84 := 0; fi84 < 5; fi84++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v85 := GenS1{F0: v3, F1: 43, F2: (^(v3))}
            _ = fi84
            _ = v85
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
}
