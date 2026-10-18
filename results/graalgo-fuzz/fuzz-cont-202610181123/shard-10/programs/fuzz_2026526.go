package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32; F1 string; F2 int16 }
type GenS7 struct { F0 string; F1 int8 }

func main() {
    var v1 uint16 = 491
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    var v4 uint = 1
    v4 -= 42
    v4 += v4
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint32 = 65535
        var v7 uint64 = 2
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        fmt.Println(v1, v1)
    case "a":
        fmt.Println(v4, v1)
    }
    var v8 int16 = 42
    v9 := GenS1{F0: 217, F1: v8}
    v8 += v8
    if (-416 <= 28) {
        for fi10 := 0; fi10 < 18; fi10++ {
            v4 = 18446744073709551615
            if true {
                v8 = (v8 & (-(v8)))
                fmt.Println(v8)
                v8 = (v8 << ((v4) & 63))
            }
            v4 = v4
            _ = fi10
        }
        var v11 int8 = -72
        var v12 bool = false
        _ = v11
        _ = v12
    } else {
        fmt.Println(v1, v4)
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint = v4
            v15 := GenS2{F0: "foo bar", F1: -1}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    if false {
        delete(v3, 42)
        fmt.Println(v1, v4)
    }
    if ((0.0 + (-2.25 * -2.25)) > float32(1.5)) {
        var v16 int64 = 9223372036854775807
        if false {
            var v17 uint32 = 2
            _ = len(v2)
            for fi18 := 0; fi18 < 10; fi18++ {
                v1 -= (^(v1))
                v1 = v1
                _ = fi18
            }
            _ = v17
        }
        _ = v16
    }
    v19 := GenS3{F0: uint8(v1), F1: 1}
    v1 = v1
    var v20 uint = 198
    switch 42 {
    case -749:
        var v21 int8 = 42
        _ = v21
    }
    if false {
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint8 = 0
            fmt.Println(v1)
            switch "" {
            case "\n\t":
                for fi24 := 0; fi24 < 13; fi24++ {
                    switch -2147483648 {
                    case -2147483648:
                        switch "foo bar" {
                        case "\n\t":
                            v23 = uint8(v1)
                            var v25 float64 = (1.5 - (1e10 * float64(float32(1e10))))
                            switch "foo bar" {
                            case "\n\t":
                                switch ("foo bar" + "foo bar") {
                                case "\n\t":
                                    var v26 int64 = 9223372036854775807
                                    v27 := GenS3{F0: (v23 + (v23 % ((128) | 1))), F1: 0}
                                    var v28 string = "a"
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                case "a":
                                    var v29 uint16 = v1
                                    _ = v29
                                }
                            case "a":
                                var v30 uint = (v4 % ((v4) | 1))
                                _ = v30
                            case "hello":
                                v31 := GenS3{F0: 255, F1: 1}
                                _ = v31
                            default:
                                if true {
                                    if true {
                                        fmt.Println(v23)
                                        var v32 bool = false
                                        _ = v32
                                    }
                                } else {
                                    var v33 float32 = 3.14159
                                    if false {
                                        v1 += v1
                                        if (true && ((false || false) || true)) {
                                            v4 -= 2
                                            switch (-(fi22)) {
                                            case 42:
                                                var v34 int32 = -1
                                                _ = v34
                                            case 995:
                                                if true {
                                                    v4 += v20
                                                    var v35 int32 = -1
                                                    if ((v23 ^ v23) != 170) {
                                                        var v36 string = ("a" + ("hello" + "foo bar"))
                                                        var v37 float64 = (0.1 - (((0.0 / v25) * (0.1 + v25)) / 0.1))
                                                        var v38 float64 = (((3.14159 + v25) / 1e10) - (0.0 * 1.5))
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    _ = v35
                                                } else {
                                                    fmt.Println(v8, v33, v20)
                                                    if false {
                                                        var v39 string = ("a" + ((("a" + "foo bar") + "hello") + ("" + (("hello" + ("" + (("" + ((("\n\t" + "a") + ("a" + "foo bar")) + ("hello" + "\n\t"))) + "foo bar"))) + ("a" + "")))))
                                                        v40 := GenS4{F0: (fi24 - fi22), F1: (fi22 >> (((^(v4))) & 63))}
                                                        fmt.Println(v39)
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    if false {
                                                        delete(v3, (fi24 | int(v8)))
                                                        fmt.Println(v20, v23, v25)
                                                    }
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            v3[(^(fi24))] = uint32(v23)
                                            switch "" {
                                            case "\n\t":
                                                if false {
                                                    if false {
                                                        v1 += (v1 * v1)
                                                        var v41 float32 = v33
                                                        v42 := GenS5{F0: 42, F1: 602, F2: 42}
                                                        _ = v41
                                                        _ = v42
                                                    } else {
                                                        if ((!(176 > 713)) && false) {
                                                            fmt.Println(v33, v25, v20)
                                                        }
                                                        v33 += v33
                                                        if (1.5 >= ((1e10 * ((3.14159 + v25) + v25)) / (v25 - ((v25 - (1e10 + -2.25)) / (v25 + v25))))) {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v43 int32 = 42
                                                                _ = v43
                                                            case "hello":
                                                                v3[fi24] = 192
                                                            case "":
                                                                var v44 uint16 = v1
                                                                _ = v44
                                                            }
                                                            var v45 uint = 42
                                                            var v46 uint64 = 4294967295
                                                            if (!(42 <= int8(v20))) {
                                                                v47 := GenS1{F0: (^(fi22)), F1: (v8 + v8)}
                                                                _ = v47
                                                            }
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            if true {
                                                                v48 := v3[(fi22 ^ (fi22 << (((v4 & v20)) & 63)))]
                                                                _ = v48
                                                            }
                                                            v25 -= -2.25
                                                        }
                                                    }
                                                    fmt.Println(v20, v1)
                                                } else {
                                                    var v49 float32 = 1e10
                                                    v2 = append(v2, 1)
                                                    var v50 string = ""
                                                    switch v50 {
                                                    case "":
                                                        v23 = (^(v23))
                                                    case "\n\t":
                                                        if len(v2) > 0 {
                                                            v2[0] = 2147483647
                                                        }
                                                    case "foo bar":
                                                        v49 = v49
                                                    }
                                                    var v51 int64 = 0
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case "foo bar":
                                                if ((4294967295 != 42) || (float32(1.5) != (v33 * (((0.1 / v33) + ((-2.25 / 3.14159) + float32(0.1))) / 1.5)))) {
                                                    switch -9223372036854775808 {
                                                    case 32768:
                                                        fmt.Println(v25)
                                                    case -2147483648:
                                                        var v52 float64 = v25
                                                        _ = v52
                                                    case -603:
                                                        var v53 bool = (true || false)
                                                        _ = v53
                                                    default:
                                                        fmt.Println(v4, v1)
                                                    }
                                                    v54 := GenS3{F0: 1, F1: 9223372036854775807}
                                                    fmt.Println(v33, v4, v4)
                                                    v4 = (v20 | v20)
                                                    v1 = 0
                                                    _ = v54
                                                } else {
                                                    if ((760 != 2147483647) && true) {
                                                        v4 = v20
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    fmt.Println(v4)
                                                }
                                                if (v20 == 255) {
                                                    v55 := GenS4{F0: fi22, F1: (fi22 ^ fi24)}
                                                    v25 = ((v25 + v25) * v25)
                                                    _ = len(v3)
                                                    _ = v55
                                                }
                                                fmt.Println(v1, v23, fi22)
                                                var v56 bool = true
                                                _ = v56
                                            case "":
                                                fmt.Println(v33, fi22)
                                            default:
                                                v57 := GenS4{F0: fi22, F1: (fi22 | (fi24 ^ fi22))}
                                                _ = v57
                                            }
                                            switch -9223372036854775808 {
                                            case -319:
                                                v58 := GenS3{F0: v23, F1: 9223372036854775807}
                                                _ = v58
                                            }
                                        }
                                        if (-35 < -744) {
                                            if ((1.5 > (v33 / 1.0)) || (!false)) {
                                                v59 := GenS6{F0: -1, F1: "hello", F2: (v8 << ((v20) & 63))}
                                                _ = v59
                                            }
                                            fmt.Println(v33, v33)
                                        }
                                    } else {
                                        switch "" {
                                        case "":
                                            var v60 int64 = -9223372036854775808
                                            _ = v60
                                        case "a":
                                            v20 -= (v20 % ((v4) | 1))
                                        default:
                                            v33 += v33
                                        }
                                        var v61 int64 = 42
                                        v62 := GenS6{F0: 2147483647, F1: ("hello" + "\n\t"), F2: 365}
                                        _ = v61
                                        _ = v62
                                    }
                                    _ = v33
                                }
                            }
                            fmt.Println(v20, v4, v20)
                            _ = v25
                        case "foo bar":
                            var v63 int16 = (v8 & (v8 * (v8 & v8)))
                            _ = v63
                        case "":
                            var v64 uint16 = 205
                            _ = v64
                        default:
                            v65 := GenS5{F0: 100, F1: uint64(v1), F2: (v23 - v23)}
                            _ = v65
                        }
                    }
                    v20 += 789
                    var v66 int16 = -851
                    _ = fi24
                    _ = v66
                }
            default:
                v4 = (^(v20))
            }
            _ = fi22
            _ = v23
        }
    }
    v67 := GenS7{F0: "hello", F1: 49}
    v68 := GenS1{F0: -588, F1: v8}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v8
    _ = v9
}
