package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 bool }

func main() {
    var v1 uint16 = 18
    var v2 uint32 = 2147483647
    v3 := []int{128, 32768, 128, int(v2)}
    var v4 int16 = 32767
    var v5 bool = (842 == v2)
    v6 := []int{-1, 4294967295}
    v1 = v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (0 != -1) {
        if false {
            var v7 int64 = 1
            _ = len(v6)
            switch "\n\t" {
            case "hello":
                v4 += v4
            default:
                var v8 string = "foo bar"
                _ = v8
            }
            _ = v7
        } else {
            switch -1 {
            case 32768:
                switch ("foo bar" + ("hello" + ("foo bar" + ("foo bar" + "")))) {
                case "\n\t":
                    v6 = append(v6, 0)
                case "foo bar":
                    var v9 uint16 = v1
                    _ = v9
                }
            case -1:
                fmt.Println(v5, v4, v2)
            case 4294967295:
                if (v5 && (v4 > (^(v4)))) {
                    if ("a" < "\n\t") {
                        fmt.Println(v5)
                        v2 -= 0
                        var v10 float64 = 0.0
                        _ = v10
                    }
                    if false {
                        if ("a" > "foo bar") {
                            if (((1e10 * ((3.14159 + 1e10) + 1e10)) * (0.1 / ((0.1 / 1.0) * ((-2.25 / (1e10 / 1e10)) * 3.14159)))) < 0.1) {
                                if (v5 && true) {
                                    for fi11 := 0; fi11 < 9; fi11++ {
                                        switch "foo bar" {
                                        case "hello":
                                            v1 = 32767
                                        case "a":
                                            if (v5 || (0.0 <= 0.1)) {
                                                _ = len(v3)
                                                if (0 <= 1) {
                                                    v4 = (v4 << ((0) & 63))
                                                    for fi12 := 0; fi12 < 1; fi12++ {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        fmt.Println(fi12, v5)
                                                        fmt.Println(v2, v5, fi11)
                                                        _ = fi12
                                                    }
                                                    v2 = 65535
                                                }
                                                fmt.Println(v1)
                                            }
                                        case "foo bar":
                                            v13 := GenS1{F0: (^(fi11)), F1: (v2 ^ (v2 * 174))}
                                            _ = v13
                                        }
                                        _ = fi11
                                    }
                                    if (true && false) {
                                        v4 -= v4
                                        if v5 {
                                            var v14 float64 = 1e10
                                            v15 := GenS1{F0: -596, F1: v2}
                                            _ = v14
                                            _ = v15
                                        }
                                    } else {
                                        v16 := GenS1{F0: 9223372036854775807, F1: (v2 % ((65535) | 1))}
                                        v1 = (v1 >> ((255) & 63))
                                        switch 128 {
                                        case 827:
                                            v5 = false
                                        }
                                        _ = v16
                                    }
                                }
                            } else {
                                v5 = v5
                                var v17 int32 = -29
                                _ = v17
                            }
                            v18 := GenS2{F0: 1}
                            _ = v18
                        }
                        v19 := GenS2{F0: 1}
                        var v20 float64 = 0.0
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint32 = 2147483648
                        _ = v21
                    }
                } else {
                    if false {
                        v22 := GenS1{F0: -752, F1: (v2 >> ((0) & 63))}
                        switch ("" + "") {
                        case "foo bar":
                            var v23 uint = 1
                            _ = v23
                        default:
                            v24 := GenS2{F0: 19}
                            _ = v24
                        }
                        var v25 float64 = 1.5
                        _ = v22
                        _ = v25
                    } else {
                        var v26 float64 = ((0.0 * -2.25) + ((1.5 + 0.0) / 1.5))
                        switch 127 {
                        case 42:
                            v6 = append(v6, 401)
                        case -2147483648:
                            v5 = v5
                        }
                        if (4294967295 >= 42) {
                            for fi27 := 0; fi27 < 11; fi27++ {
                                if true {
                                    var v28 uint16 = v1
                                    for fi29 := 0; fi29 < 19; fi29++ {
                                        if (!((v4 & v4) <= (^(v4)))) {
                                            v30 := GenS1{F0: fi29, F1: v2}
                                            _ = v30
                                        } else {
                                            if ((v2 & v2) < 338) {
                                                v26 = ((v26 + 1.5) - v26)
                                            }
                                        }
                                        v5 = false
                                        _ = fi29
                                    }
                                    var v31 float64 = 1.5
                                    _ = v28
                                    _ = v31
                                } else {
                                    if ((true && false) && (69 <= 42)) {
                                        var v32 float64 = ((((v26 / v26) * v26) / ((0.0 / (v26 - 0.1)) * (1e10 + ((0.1 * 0.1) * ((v26 - -2.25) * (1e10 / float64((((((0.1 + ((1e10 - 0.0) / float32(v26))) / 1.0) / 1e10) / 1.0) - 0.1)))))))) - v26)
                                        _ = v32
                                    }
                                    if false {
                                        var v33 int16 = -481
                                        _ = v33
                                    } else {
                                        for fi34 := 0; fi34 < 20; fi34++ {
                                            if ((v1 > v1) || v5) {
                                                v35 := GenS1{F0: 0, F1: 2}
                                                v36 := GenS2{F0: -183}
                                                _ = v35
                                                _ = v36
                                            } else {
                                                if v5 {
                                                    v37 := GenS3{F0: (-1 >= 0)}
                                                    fmt.Println(v4, v26)
                                                    _ = v37
                                                } else {
                                                    if ((v1 & 241) > (^(v1))) {
                                                        v4 += v4
                                                        fmt.Println(v2, v1, fi34)
                                                        v26 = v26
                                                        v2 = (v2 % (((v2 % ((2147483647) | 1))) | 1))
                                                    }
                                                    fmt.Println(fi34)
                                                    v2 = 785
                                                }
                                                if (42 == 9223372036854775807) {
                                                    switch ("hello" + "foo bar") {
                                                    case "foo bar":
                                                        var v38 uint32 = v2
                                                        _ = v38
                                                    default:
                                                        v5 = false
                                                    }
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (true || (!(false && (!v5)))) {
                                                            v39 := GenS4{F0: 0}
                                                            var v40 bool = v5
                                                            v2 += uint32(v1)
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            if v5 {
                                                                var v41 float32 = 3.14159
                                                                var v42 int16 = (v4 << ((42) & 63))
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            v26 = -2.25
                                                            if v5 {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    var v43 uint8 = 191
                                                                    _ = v43
                                                                default:
                                                                    v4 += v4
                                                                }
                                                            }
                                                            if ((!true) && v5) {
                                                                v3 = append(v3, int(v1))
                                                                v44 := GenS5{F0: v1, F1: 65535, F2: v2}
                                                                fmt.Println(fi34)
                                                                _ = v44
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = (fi34 | fi27)
                                                            }
                                                            v26 = 0.0
                                                        }
                                                    case "":
                                                        v45 := GenS6{F0: 948, F1: (v1 << ((841) & 63))}
                                                        _ = v45
                                                    case "hello":
                                                        var v46 int64 = 394
                                                        _ = v46
                                                    }
                                                    if (!v5) {
                                                        var v47 int32 = 2147483647
                                                        v48 := GenS7{F0: v5}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                }
                                            }
                                            v4 += (v4 & 202)
                                            switch "a" {
                                            case "a":
                                                if ((5 < 177) || v5) {
                                                    fmt.Println(fi34, v5)
                                                    var v49 string = "hello"
                                                    _ = v49
                                                } else {
                                                    v50 := GenS5{F0: (v1 / (((v1 & (^(v1)))) | 1)), F1: 2147483648, F2: v2}
                                                    var v51 int64 = 0
                                                    if ((v2 >> ((667) & 63)) <= v2) {
                                                        var v52 int16 = (v4 / (((v4 % (((v4 >> ((2) & 63))) | 1))) | 1))
                                                        _ = v52
                                                    } else {
                                                        fmt.Println(v4)
                                                        var v53 bool = (!v5)
                                                        _ = v53
                                                    }
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case "hello":
                                                var v54 float32 = float32(v26)
                                                _ = v54
                                            }
                                            _ = fi34
                                        }
                                        var v55 uint64 = 42
                                        for fi56 := 0; fi56 < 5; fi56++ {
                                            var v57 int16 = v4
                                            var v58 uint32 = v2
                                            v59 := GenS4{F0: int16(v58)}
                                            _ = fi56
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = v55
                                    }
                                }
                                for fi60 := 0; fi60 < 17; fi60++ {
                                    v61 := GenS4{F0: v4}
                                    var v62 uint8 = 2
                                    v62 -= (v62 % ((v62) | 1))
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                                if (!(1.5 > float32(3.14159))) {
                                    var v63 string = ("hello" + "foo bar")
                                    _ = v63
                                }
                                _ = fi27
                            }
                            var v64 uint8 = 127
                            _ = v64
                        }
                        _ = v26
                    }
                }
            }
            if true {
                v65 := GenS3{F0: v5}
                v4 = v4
                _ = v65
            }
        }
        v5 = false
    }
    var v66 int8 = -1
    v67 := GenS5{F0: (v1 | (v1 / ((v1) | 1))), F1: v2, F2: (v2 / ((v2) | 1))}
    fmt.Println(v66, v66, v66)
    v68 := GenS5{F0: 0, F1: v2, F2: 65535}
    v2 -= (v2 * v2)
    _ = len(v6)
    switch 9223372036854775807 {
    case -189:
        var v69 uint8 = 0
        _ = v69
    default:
        fmt.Println(v5, v2)
    }
    if v5 {
        switch -823 {
        case 9223372036854775807:
            v4 = 1000
        case 1:
            var v70 float32 = 0.1
            _ = v70
        }
        if (-1 <= -9223372036854775808) {
            _ = len(v3)
            switch 42 {
            case 32767:
                var v71 uint32 = 4294967295
                _ = v71
            default:
                v72 := GenS3{F0: false}
                _ = v72
            }
        } else {
            v1 = v1
            switch "\n\t" {
            case "":
                for fi73 := 0; fi73 < 11; fi73++ {
                    for fi74 := 0; fi74 < 6; fi74++ {
                        fmt.Println(v4)
                        _ = fi74
                    }
                    var v75 uint = 2
                    if (854 >= int32(v2)) {
                        v66 += int8(v75)
                    }
                    _ = fi73
                    _ = v75
                }
            case "foo bar":
                for fi76 := 0; fi76 < 16; fi76++ {
                    var v77 float64 = float64(float32(float64(0.1)))
                    var v78 uint8 = 141
                    var v79 uint64 = 4294967295
                    _ = fi76
                    _ = v77
                    _ = v78
                    _ = v79
                }
            default:
                if len(v6) > 0 {
                    v6[0] = int(v66)
                }
            }
            fmt.Println(v66, v2, v2)
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v1 = v1
    fmt.Println(v1)
    if (667 <= 4294967295) {
        v2 -= 4294967295
    }
    v2 -= (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
}
