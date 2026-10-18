package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 bool; F2 int32 }
type GenS9 struct { F0 int32; F1 int64; F2 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 int64; F1 bool; F2 int64 }
func genFn0() uint8 {
    if true {
        if true {
            v6 := GenS1{F0: 1000, F1: 32767}
            _ = v6
        } else {
            if true {
                v7 := GenS1{F0: 32767, F1: 1}
                var v8 uint32 = 65535
                _ = v7
                _ = v8
            } else {
                switch ("\n\t" + ((((("" + "foo bar") + ("foo bar" + ((("foo bar" + "a") + ("hello" + "\n\t")) + "hello"))) + "a") + "\n\t") + ("hello" + "foo bar"))) {
                case "hello":
                    if (438 > 2) {
                        v9 := GenS2{F0: 704}
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 19; fi10++ {
                            var v11 uint = 623
                            v12 := GenS2{F0: 42}
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    }
                case "\n\t":
                    var v13 uint16 = 2
                    _ = v13
                case "":
                    var v14 float64 = -2.25
                    _ = v14
                }
            }
        }
        v15 := GenS3{F0: 436, F1: 118}
        var v16 int8 = 127
        _ = v15
        _ = v16
    } else {
        if (!(1 == 2147483648)) {
            var v17 int64 = 1
            _ = v17
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := []int{fi18, (-(fi18))}
            switch fi18 {
            case 9223372036854775807:
                if (0 < 65535) {
                    switch "" {
                    case "hello":
                        for fi20 := 0; fi20 < 14; fi20++ {
                            fmt.Println(fi18, fi18)
                            _ = fi20
                        }
                    }
                } else {
                    if len(v19) > 0 {
                        v19[0] = (fi18 + (-(fi18)))
                    }
                    v21 := GenS4{F0: 128}
                    _ = v21
                }
            default:
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            }
            _ = fi18
            _ = v19
        }
    }
    v22 := []int{-126, 183}
    v23 := GenS3{F0: 42, F1: 951}
    _ = v22
    _ = v23
    return 1
}
func genFn1() uint {
    var v24 uint8 = 122
    v25 := GenS5{F0: 4294967295}
    var v26 uint32 = 4294967295
    _ = v24
    _ = v25
    _ = v26
    return uint(v26)
}

func main() {
    var v1 uint8 = 117
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = (v3 % ((v3) | 1))
    var v5 float32 = (1e10 * ((float32((3.14159 * (0.0 + 0.1))) - (float32((0.1 / 1.0)) / (1e10 + 0.0))) + 0.0))
    if (-2.25 < -2.25) {
        if false {
            if false {
                var v27 bool = (912 < 2)
                v28 := GenS6{F0: ((^(v4)) < v3)}
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 bool = true
                v31 := GenS7{F0: 1000}
                switch "a" {
                case "hello":
                    v32 := GenS8{F0: uint(v1), F1: (3.14159 != 1.5), F2: 1}
                    _ = v32
                case "foo bar":
                    _ = len(v2)
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        v33 := GenS5{F0: 42}
        v34 := []int{533, 4294967295}
        _ = v33
        _ = v34
    }
    fmt.Println(v5, v1, v4)
    switch "foo bar" {
    case "":
        for fi35 := 0; fi35 < 12; fi35++ {
            if true {
                fmt.Println(v1)
                v1 -= 202
            }
            v36 := GenS9{F0: 142, F1: 774, F2: -1}
            _ = fi35
            _ = v36
        }
    case "foo bar":
        v1 -= (v1 % (((v1 << ((18446744073709551615) & 63))) | 1))
    case "a":
        for fi37 := 0; fi37 < 8; fi37++ {
            if ((float64((((v5 / 1.0) * (float32(1e10) * 1.5)) / 1e10)) >= 3.14159) && (-2.25 <= 1.5)) {
                v38 := GenS8{F0: 152, F1: ((9223372036854775807 <= 903) && false), F2: -38}
                v39 := GenS5{F0: 255}
                _ = v38
                _ = v39
            } else {
                v3 = 1
                fmt.Println(fi37, fi37, v5)
            }
            _ = fi37
        }
    default:
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 int32 = -690
            _ = fi40
            _ = v41
        }
    }
    v42 := []int{849, -878, 32768, 128}
    v2[65535] = 0
    v43 := GenS10{F0: 0}
    v4 += (v4 - v4)
    var v44 uint32 = 42
    if true {
        v45 := GenS11{F0: 42, F1: v4}
        var v46 float32 = -2.25
        v47 := GenS6{F0: false}
        _ = v45
        _ = v46
        _ = v47
    }
    v5 -= 0.0
    for fi48 := 0; fi48 < 1; fi48++ {
        for fi49 := 0; fi49 < 17; fi49++ {
            v50 := v2[(v44 / ((v44) | 1))]
            var v51 float32 = ((-2.25 / (-2.25 / 1.5)) + -2.25)
            _ = fi49
            _ = v50
            _ = v51
        }
        fmt.Println(v1)
        _ = fi48
    }
    v4 -= v4
    if (uint8(v44) >= (v1 + (v1 >> ((372) & 63)))) {
        if true {
            switch -229 {
            case 9223372036854775807:
                var v52 uint32 = v44
                _ = v52
            default:
                if false {
                    fmt.Println(v5, v5, v3)
                    var v53 float64 = -2.25
                    fmt.Println(v4)
                    _ = v53
                }
            }
            switch 338 {
            case 0:
                var v54 int64 = -1
                _ = v54
            default:
                var v55 string = "foo bar"
                _ = v55
            }
            if (!true) {
                v56 := GenS12{F0: 23}
                v1 = genFn0()
                _ = len(v2)
                _ = v56
            } else {
                v3 = (v3 >> ((65535) & 63))
                for fi57 := 0; fi57 < 15; fi57++ {
                    for fi58 := 0; fi58 < 7; fi58++ {
                        if false {
                            if (true && true) {
                                v59 := GenS3{F0: 0, F1: 2}
                                v60 := v2[v44]
                                _ = v59
                                _ = v60
                            }
                            var v61 int32 = 725
                            v62 := GenS1{F0: 1000, F1: (fi58 ^ (fi58 & (-(fi57))))}
                            _ = v61
                            _ = v62
                        } else {
                            v44 = (v44 / (((v44 >> ((376) & 63))) | 1))
                            switch "\n\t" {
                            case "hello":
                                if (2 == 688) {
                                    var v63 uint = 315
                                    v63 = 516
                                    v64 := GenS1{F0: 1000, F1: fi57}
                                    v65 := GenS13{F0: 42}
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                } else {
                                    if true {
                                        var v66 uint32 = (v44 | 2147483647)
                                        _ = v66
                                    } else {
                                        var v67 int64 = -9223372036854775808
                                        var v68 uint8 = (v1 + v1)
                                        _ = v67
                                        _ = v68
                                    }
                                    v69 := GenS14{F0: 1}
                                    v70 := GenS3{F0: -334, F1: 51}
                                    _ = v69
                                    _ = v70
                                }
                            }
                            v3 -= (v3 - v3)
                        }
                        v5 -= v5
                        _ = fi58
                    }
                    v71 := GenS5{F0: 713}
                    _ = fi57
                    _ = v71
                }
            }
        } else {
            if len(v42) >= 2 {
                v42 = v42[1:len(v42)]
            }
            if (1 != 114) {
                _ = len(v42)
            } else {
                if (false || true) {
                    v3 = (v3 / ((18446744073709551615) | 1))
                    switch 1 {
                    case 225:
                        if false {
                            for fi72 := 0; fi72 < 7; fi72++ {
                                fmt.Println(v1, v3, v44)
                                for fi73 := 0; fi73 < 19; fi73++ {
                                    v3 -= (^(v4))
                                    v74 := GenS10{F0: -32768}
                                    v75 := GenS13{F0: 94}
                                    _ = fi73
                                    _ = v74
                                    _ = v75
                                }
                                _ = fi72
                            }
                            fmt.Println(v3, v3)
                            fmt.Println(v3, v4, v5)
                        }
                    default:
                        v76 := GenS15{F0: 0, F1: 1000}
                        _ = v76
                    }
                    var v77 int = 0
                    _ = v77
                } else {
                    if false {
                        var v78 int16 = 1000
                        var v79 uint32 = (v44 | v44)
                        if (!(!(true || (false && true)))) {
                            switch 2147483647 {
                            case -285:
                                if len(v42) >= 2 {
                                    v42 = v42[1:len(v42)]
                                }
                            case -2147483648:
                                var v80 string = "a"
                                _ = v80
                            }
                        }
                        _ = v78
                        _ = v79
                    }
                    fmt.Println(v4, v5, v3)
                }
            }
        }
        var v81 uint32 = 1
        var v82 int16 = -1
        _ = v81
        _ = v82
    } else {
        fmt.Println(v4)
        if (4294967295 < (v44 * (v44 << ((866) & 63)))) {
            v42 = append(v42, 4294967295)
        }
    }
    var v83 int = 127
    v1 = v1
    v3 -= (v4 - (^(v4)))
    v84 := GenS6{F0: false}
    v85 := GenS16{F0: 42, F1: (!(!true)), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v83
    _ = v84
    _ = v85
}
