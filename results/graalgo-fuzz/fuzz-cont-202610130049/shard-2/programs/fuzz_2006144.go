package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
type GenS4 struct { F0 uint64; F1 int16; F2 int8 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32; F1 int; F2 uint8 }
func genFn0(p0 uint32, p1 int64) uint16 {
    p0 -= p0
    var v5 int16 = -996
    _ = p0
    _ = p1
    _ = v5
    return 255
}

func main() {
    var v1 uint = 119
    var v2 int32 = 2147483647
    v3 := []int{42}
    var v4 float64 = -2.25
    if false {
        v6 := GenS1{F0: -73, F1: 1}
        if (42 >= 18446744073709551615) {
            v7 := GenS1{F0: -1, F1: 4294967295}
            switch 9223372036854775807 {
            case 42:
                if (9223372036854775807 < 0) {
                    if (9223372036854775807 < -997) {
                        v8 := GenS2{F0: v1}
                        v2 += v2
                        var v9 int = 32768
                        _ = v8
                        _ = v9
                    } else {
                        if (!true) {
                            fmt.Println(v1)
                            v4 = (float64(0.0) / v4)
                            if (-346 > 0) {
                                var v10 float64 = (((1e10 * v4) / float64(1e10)) + (v4 + 0.1))
                                _ = v10
                            }
                        } else {
                            for fi11 := 0; fi11 < 1; fi11++ {
                                v1 = v1
                                switch (fi11 << (((v1 / (((v1 / (((v1 + 1)) | 1))) | 1))) & 63)) {
                                case 42:
                                    var v12 float32 = (1e10 - 0.0)
                                    _ = v12
                                }
                                _ = fi11
                            }
                            v2 += (^(v2))
                            switch ("foo bar" + "foo bar") {
                            case "\n\t":
                                for fi13 := 0; fi13 < 5; fi13++ {
                                    v4 -= float64(3.14159)
                                    v14 := GenS2{F0: 140}
                                    var v15 uint16 = 32767
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            default:
                                switch "" {
                                case "\n\t":
                                    v16 := GenS3{F0: 42, F1: 44, F2: "a"}
                                    _ = v16
                                default:
                                    v17 := GenS4{F0: 18446744073709551615, F1: -918, F2: -82}
                                    _ = v17
                                }
                            }
                        }
                    }
                    fmt.Println(v4)
                } else {
                    v1 -= v1
                    v1 = v1
                    if true {
                        fmt.Println(v4, v4, v1)
                        var v18 uint64 = 675
                        var v19 uint64 = (v18 + 1)
                        _ = v18
                        _ = v19
                    } else {
                        v4 -= ((1.5 / v4) * v4)
                        _ = len(v3)
                    }
                }
            case 0:
                var v20 float64 = 0.0
                _ = v20
            case 1:
                if true {
                    v4 -= 0.0
                    fmt.Println(v2, v1)
                    var v21 float32 = (0.0 - 1.5)
                    _ = v21
                } else {
                    if false {
                        var v22 uint16 = 1
                        var v23 bool = (!(91 < 0))
                        if (2147483647 != 0) {
                            v24 := GenS5{F0: v23, F1: 9223372036854775807}
                            v1 += v1
                            if v23 {
                                v22 += uint16(v1)
                                if true {
                                    switch -648 {
                                    case 483:
                                        for fi25 := 0; fi25 < 9; fi25++ {
                                            var v26 bool = v23
                                            var v27 int32 = -2147483648
                                            var v28 uint64 = 1
                                            _ = fi25
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                    case 128:
                                        v29 := GenS2{F0: v1}
                                        _ = v29
                                    case 32767:
                                        fmt.Println(v4, v1)
                                    }
                                }
                            }
                            _ = v24
                        } else {
                            v30 := GenS5{F0: (90 < 2), F1: -1}
                            v31 := GenS1{F0: 42, F1: 32767}
                            _ = v30
                            _ = v31
                        }
                        _ = v22
                        _ = v23
                    } else {
                        v32 := GenS3{F0: 2, F1: 65535, F2: (("" + "hello") + "a")}
                        fmt.Println(v2, v1, v4)
                        _ = v32
                    }
                    var v33 int = -307
                    _ = v33
                }
            }
            var v34 float64 = (v4 + v4)
            _ = v34
            _ = v7
        }
        v2 += -1
        _ = v6
    } else {
        var v35 int64 = 42
        v1 = v1
        if false {
            fmt.Println(v1)
            switch "" {
            case "\n\t":
                if false {
                    var v36 int64 = v35
                    _ = v36
                }
            case "a":
                fmt.Println(v35, v1)
            case "hello":
                var v37 uint64 = 18446744073709551615
                _ = v37
            default:
                v1 = v1
            }
            fmt.Println(v1, v4)
        } else {
            if len(v3) > 0 {
                v3[0] = 0
            }
            v4 = v4
        }
        _ = v35
    }
    v4 += -2.25
    for fi38 := 0; fi38 < 2; fi38++ {
        switch "hello" {
        case "a":
            for fi39 := 0; fi39 < 12; fi39++ {
                fmt.Println(fi38, v1)
                var v40 uint64 = 18446744073709551615
                _ = fi39
                _ = v40
            }
        }
        _ = fi38
    }
    v41 := GenS6{F0: 0, F1: 429}
    var v42 uint64 = 42
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 uint64 = 992
        v2 = -2147483648
        _ = fi43
        _ = v44
    }
    fmt.Println(v2)
    v4 -= v4
    var v45 int = -473
    if true {
        switch "a" {
        case "hello":
            switch v45 {
            case 42:
                v46 := GenS7{F0: (v45 & (-(v45)))}
                _ = v46
            case 1:
                var v47 string = ""
                _ = v47
            default:
                for fi48 := 0; fi48 < 13; fi48++ {
                    if true {
                        switch ("foo bar" + ("" + "a")) {
                        case "foo bar":
                            v45 += v45
                        case "":
                            var v49 uint64 = v42
                            _ = v49
                        default:
                            var v50 bool = (((((!true) && true) || (v1 == v1)) || false) || ((v2 ^ v2) != 1))
                            _ = v50
                        }
                    } else {
                        var v51 float32 = -2.25
                        _ = v51
                    }
                    _ = fi48
                }
            }
        case "a":
            var v52 uint16 = genFn0(0, -9223372036854775808)
            _ = v52
        case "\n\t":
            if false {
                switch "hello" {
                case "\n\t":
                    v45 = (v45 ^ v45)
                }
                if true {
                    for fi53 := 0; fi53 < 11; fi53++ {
                        if false {
                            v54 := GenS5{F0: true, F1: -9223372036854775808}
                            v55 := GenS8{F0: genFn0(80, 225)}
                            _ = v54
                            _ = v55
                        }
                        _ = fi53
                    }
                    v56 := GenS9{F0: 42, F1: "hello"}
                    _ = v56
                }
            }
        }
    } else {
        var v57 uint8 = 42
        _ = v57
    }
    if (false || false) {
        for fi58 := 0; fi58 < 8; fi58++ {
            for fi59 := 0; fi59 < 17; fi59++ {
                fmt.Println(v42)
                var v60 int = fi59
                _ = fi59
                _ = v60
            }
            v45 = fi58
            v61 := GenS10{F0: 0}
            _ = fi58
            _ = v61
        }
    }
    var v62 bool = false
    v42 -= (v42 + (^(v42)))
    var v63 int32 = v2
    switch "hello" {
    case "\n\t":
        v1 += (v1 >> ((0) & 63))
    case "":
        v64 := GenS11{F0: "foo bar"}
        _ = v64
    case "hello":
        for fi65 := 0; fi65 < 3; fi65++ {
            v4 -= v4
            switch "foo bar" {
            case "foo bar":
                if false {
                    v62 = v62
                    var v66 float64 = ((v4 - float64(((float32(v4) + float32(0.0)) + 1.5))) * v4)
                    if v62 {
                        var v67 float64 = 1e10
                        fmt.Println(fi65)
                        switch fi65 {
                        case -2147483648:
                            var v68 uint64 = (v42 % (((v42 ^ (v42 << ((v1) & 63)))) | 1))
                            _ = v68
                        case 303:
                            var v69 string = ("\n\t" + "foo bar")
                            _ = v69
                        }
                        _ = v67
                    }
                    _ = v66
                } else {
                    v63 += v63
                }
            default:
                v3 = append(v3, v45)
            }
            _ = fi65
        }
    }
    var v70 int16 = 1000
    if (!v62) {
        v1 += v1
    } else {
        var v71 uint64 = 4294967295
        v70 = (v70 >> ((v1) & 63))
        _ = v71
    }
    switch (v45 % ((385) | 1)) {
    case -2147483648:
        var v72 int16 = v70
        _ = v72
    case -988:
        switch "a" {
        case "a":
            var v73 int32 = -1
            _ = v73
        case "":
            for fi74 := 0; fi74 < 2; fi74++ {
                v63 -= v2
                switch v45 {
                case 32767:
                    v75 := GenS6{F0: 32767, F1: -9223372036854775808}
                    _ = v75
                default:
                    v76 := GenS12{F0: -133, F1: (fi74 + -9223372036854775808), F2: uint8(v1)}
                    _ = v76
                }
                var v77 uint8 = 0
                _ = fi74
                _ = v77
            }
        }
    default:
        var v78 uint32 = 1
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v62
    _ = v63
    _ = v70
}
