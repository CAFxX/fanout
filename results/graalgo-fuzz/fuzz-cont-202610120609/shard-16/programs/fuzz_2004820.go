package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int64; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint; F1 int16; F2 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 uint16 = 2
    v4 := GenS1{F0: 100, F1: 1, F2: 127}
    var v5 uint32 = 469
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            v5 = 4294967295
            for fi7 := 0; fi7 < 3; fi7++ {
                if ((^(v5)) <= 42) {
                    if ((!true) || (fi7 <= (fi6 / (((fi6 ^ fi6)) | 1)))) {
                        if (0.1 <= (((float64(1.5) + -2.25) / (-2.25 * (0.1 + (-2.25 + 3.14159)))) + -2.25)) {
                            v8 := v2[(v5 * v5)]
                            var v9 int8 = -25
                            v3 += (^(v3))
                            _ = v8
                            _ = v9
                        }
                    } else {
                        var v10 string = ""
                        var v11 float64 = (0.1 + (((0.0 + (1.5 * ((3.14159 - 0.1) - 1.5))) / (((0.0 + (3.14159 - (0.0 + 3.14159))) / ((0.1 - 1.5) - ((-2.25 + 1.5) - 1.5))) - (((-2.25 - 0.1) / 1.0) / (-2.25 - 0.1)))) + 0.1))
                        _ = v10
                        _ = v11
                    }
                }
                if ((float32((-2.25 * 3.14159)) - -2.25) == float32(1e10)) {
                    v3 = 2
                    fmt.Println(v3, v3, fi7)
                } else {
                    v3 = v3
                    var v12 int64 = 1
                    v12 -= (v12 + 0)
                    _ = v12
                }
                _ = fi7
            }
            var v13 uint = 18446744073709551615
            _ = fi6
            _ = v13
        }
        v3 = v3
        if (!(true && true)) {
            var v14 int = 230
            var v15 uint = 42
            var v16 uint16 = 1
            _ = v14
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 uint16 = 83
                _ = fi17
                _ = v18
            }
            switch 127 {
            case 4294967295:
                fmt.Println(v3)
            case -869:
                if true {
                    v5 += v5
                } else {
                    v19 := GenS2{F0: -577}
                    fmt.Println(v5, v3)
                    _ = v19
                }
            case 9223372036854775807:
                switch 127 {
                case 42:
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 uint = 502
                        switch fi20 {
                        case 525:
                            if (-1 < 0) {
                                delete(v2, (^(v5)))
                                v5 = v5
                            } else {
                                var v22 uint32 = (v5 << ((uint(v5)) & 63))
                                switch 127 {
                                case -705:
                                    v22 = v22
                                case 9223372036854775807:
                                    if (true && (1 != -1)) {
                                        _ = len(v1)
                                        v23 := GenS3{F0: 32767, F1: -9223372036854775808, F2: 32767}
                                        v24 := []int{644, (fi20 % ((fi20) | 1))}
                                        _ = v23
                                        _ = v24
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            if (!false) {
                                                var v25 uint16 = v3
                                                var v26 bool = true
                                                v27 := GenS4{F0: 968, F1: (v21 & (^(v21))), F2: uint64(v25)}
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                        case "foo bar":
                                            for fi28 := 0; fi28 < 14; fi28++ {
                                                v22 = (v22 - 0)
                                                if false {
                                                    v3 += 32767
                                                    fmt.Println(fi28)
                                                    if true {
                                                        if (v5 == (^(v5))) {
                                                            v29 := GenS2{F0: -271}
                                                            _ = v29
                                                        } else {
                                                            if (0 != 410) {
                                                                v22 = 906
                                                                var v30 bool = false
                                                                _ = v30
                                                            }
                                                        }
                                                    } else {
                                                        if true {
                                                            v31 := GenS3{F0: 42, F1: -9223372036854775808, F2: 32767}
                                                            if (v22 >= (v22 * 2)) {
                                                                if (!(2147483647 > fi20)) {
                                                                    var v32 uint = 37
                                                                    v5 = (v5 + 178)
                                                                    var v33 uint = v21
                                                                    _ = v32
                                                                    _ = v33
                                                                } else {
                                                                    v22 += (v22 / ((v22) | 1))
                                                                    var v34 string = ("foo bar" + ("\n\t" + "a"))
                                                                    _ = v34
                                                                }
                                                                delete(v1, v22)
                                                                var v35 int16 = 1000
                                                                _ = v35
                                                            }
                                                            v36 := GenS1{F0: 53, F1: -2147483648, F2: 42}
                                                            _ = v31
                                                            _ = v36
                                                        } else {
                                                            fmt.Println(fi28)
                                                        }
                                                        v37 := GenS3{F0: 1000, F1: 559, F2: 52}
                                                        _ = v37
                                                    }
                                                }
                                                delete(v1, uint32(v21))
                                                switch ("foo bar" + "") {
                                                case "a":
                                                    v21 -= (v21 >> ((v21) & 63))
                                                default:
                                                    v38 := []int{9223372036854775807}
                                                    _ = v38
                                                }
                                                _ = fi28
                                            }
                                        }
                                    }
                                }
                                v5 = v22
                                _ = v22
                            }
                        case 42:
                            v39 := []int{(-(fi20)), (fi20 + -9223372036854775808)}
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 7; fi40++ {
                            v41 := v1[v5]
                            _ = fi40
                            _ = v41
                        }
                        _ = fi20
                        _ = v21
                    }
                case 135:
                    v5 = (v5 * (v5 * 1))
                }
            }
            v5 = 839
        }
    }
    v42 := GenS4{F0: 2, F1: uint(v3), F2: 0}
    if (564 > 1) {
        v43 := []int{32768}
        _ = v43
    } else {
        var v44 int32 = 2147483647
        for fi45 := 0; fi45 < 5; fi45++ {
            if (!(2147483648 <= v5)) {
                var v46 string = "hello"
                for fi47 := 0; fi47 < 16; fi47++ {
                    v5 -= v5
                    _ = fi47
                }
                _ = v46
            }
            var v48 uint16 = v3
            v44 -= (v44 + (v44 - v44))
            _ = fi45
            _ = v48
        }
        _ = v44
    }
    var v49 uint32 = uint32(v3)
    switch ("a" + "hello") {
    case "\n\t":
        v50 := GenS5{F0: "hello", F1: 127}
        _ = v50
    default:
        v1[(v49 ^ 472)] = "hello"
    }
    v51 := GenS3{F0: 42, F1: 0, F2: -116}
    switch 810 {
    case -9223372036854775808:
        v52 := GenS6{F0: 486, F1: -1, F2: ("" + "a")}
        _ = v52
    case 42:
        v53 := GenS4{F0: 2, F1: 979, F2: 454}
        _ = v53
    case -2147483648:
        v54 := GenS5{F0: "hello", F1: 127}
        _ = v54
    default:
        switch 32767 {
        case -1:
            if ((-1 > 42) && (1 < (^(v3)))) {
                v3 = (^(v3))
                if (false || ((v3 ^ v3) > 65535)) {
                    var v55 uint64 = 615
                    var v56 float64 = -2.25
                    _ = v55
                    _ = v56
                }
                for fi57 := 0; fi57 < 11; fi57++ {
                    var v58 uint32 = (v5 - (v49 / ((65535) | 1)))
                    var v59 uint = 819
                    for fi60 := 0; fi60 < 14; fi60++ {
                        v61 := GenS4{F0: 690, F1: (^(v59)), F2: 984}
                        delete(v2, v49)
                        v49 = (v5 >> ((2) & 63))
                        _ = fi60
                        _ = v61
                    }
                    _ = fi57
                    _ = v58
                    _ = v59
                }
            }
        case 1:
            for fi62 := 0; fi62 < 15; fi62++ {
                fmt.Println(fi62, v5)
                v5 = v49
                _ = fi62
            }
        default:
            v5 -= v49
        }
    }
    var v63 int16 = -1
    v3 = (v3 | (v3 | 255))
    var v64 int32 = 128
    var v65 uint16 = 780
    v65 = 255
    if (true || false) {
        for fi66 := 0; fi66 < 3; fi66++ {
            for fi67 := 0; fi67 < 14; fi67++ {
                v49 = (v49 - uint32(v65))
                fmt.Println(v63, v65)
                fmt.Println(fi67)
                _ = fi67
            }
            _ = fi66
        }
    } else {
        v68 := GenS7{F0: 210}
        var v69 int8 = -11
        v70 := GenS8{F0: 32767, F1: (v63 / ((int16(v64)) | 1))}
        _ = v68
        _ = v69
        _ = v70
    }
    var v71 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v51
    _ = v63
    _ = v64
    _ = v65
    _ = v71
}
