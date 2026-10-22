package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 uint32; F1 int16 }
func genFn0(p0 int, p1 int64, p2 int8) int8 {
    p2 = (p2 % ((p2) | 1))
    p1 = p1
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return (^(p2))
}
func genFn1(p0 int8, p1 int16) int16 {
    v7 := []int{-1, -861}
    var v8 int64 = 9223372036854775807
    for fi9 := 0; fi9 < 1; fi9++ {
        if len(v7) > 0 {
            v7[0] = fi9
        }
        _ = fi9
    }
    p1 -= int16(p0)
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 1
    var v3 string = "hello"
    v4 := []string{v3, "hello", v3, ("foo bar" + (("a" + ("foo bar" + "hello")) + "a"))}
    var v5 int8 = 42
    v6 := make(map[int]uint32)
    v2 = (v2 / (((v2 * 52)) | 1))
    if (!(true && (1 <= (v5 % ((genFn0(2147483647, 96, (-(v5)))) | 1))))) {
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                switch (fi10 + (fi10 ^ -9223372036854775808)) {
                case 1:
                    fmt.Println(fi10, v5, v3)
                case 32768:
                    if true {
                        var v12 int8 = v5
                        _ = v12
                    } else {
                        var v13 uint32 = 65535
                        _ = v13
                    }
                default:
                    var v14 int16 = 42
                    _ = v14
                }
                switch "foo bar" {
                case "a":
                    if (genFn0(-2147483648, 1, genFn0((fi10 & -1), -86, (^(v5)))) < (-(v5))) {
                        switch (fi10 + fi11) {
                        case 4294967295:
                            if (2147483647 == 578) {
                                var v15 int = (fi10 ^ fi10)
                                v15 += (v15 * v15)
                                _ = v15
                            } else {
                                _ = len(v4)
                                var v16 float32 = 1.5
                                _ = v16
                            }
                        case 705:
                            var v17 float64 = 0.1
                            _ = v17
                        }
                        var v18 uint32 = 2147483648
                        _ = v18
                    } else {
                        var v19 uint = 4294967295
                        if false {
                            var v20 uint64 = 1
                            switch fi11 {
                            case 601:
                                switch v3 {
                                case "foo bar":
                                    switch fi11 {
                                    case 4294967295:
                                        fmt.Println(v19)
                                    }
                                case "":
                                    v21 := GenS1{F0: -503, F1: 118}
                                    _ = v21
                                case "\n\t":
                                    var v22 int16 = 42
                                    _ = v22
                                default:
                                    if true {
                                        var v23 uint32 = 42
                                        v2 = 0
                                        if (((((float32(3.14159) - float32((float64(3.14159) + 1.5))) * (float32(0.0) * 0.0)) * 0.1) - (((0.0 / 3.14159) - 0.0) + (0.0 - 0.0))) > 0.1) {
                                            var v24 uint = (^(v19))
                                            _ = v24
                                        }
                                        _ = v23
                                    }
                                }
                            case -980:
                                v4 = append(v4, (v3 + v3))
                                v19 = v19
                                var v25 bool = (!true)
                                fmt.Println(v5, v25, v19)
                                fmt.Println(fi10, fi11)
                                _ = v25
                            }
                            _ = v20
                        } else {
                            var v26 int8 = -95
                            if (!false) {
                                fmt.Println(fi11, v19, v3)
                                v27 := GenS1{F0: 42, F1: 42}
                                _ = v27
                            } else {
                                v28 := GenS1{F0: fi11, F1: -1}
                                if ((-2.25 != 0.0) || ((fi10 % ((fi11) | 1)) == fi10)) {
                                    v3 = "\n\t"
                                } else {
                                    switch "" {
                                    case "a":
                                        _ = len(v4)
                                    case "\n\t":
                                        v19 += v19
                                    case "foo bar":
                                        if (!(false && (2147483647 != 2))) {
                                            v2 = 112
                                            var v29 float64 = float64(((float32(3.14159) * (1.5 * -2.25)) * 0.0))
                                            v30 := GenS1{F0: (fi10 | -201), F1: 163}
                                            _ = v29
                                            _ = v30
                                        }
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                        if (false || ((true && ((v5 >> (((^(v19))) & 63)) != v26)) && true)) {
                                            if true {
                                                v3 = v3
                                                if (479 <= -1) {
                                                    v5 = (^(v5))
                                                } else {
                                                    switch ((v3 + v3) + v3) {
                                                    case "hello":
                                                        v31 := GenS2{F0: (fi10 - 32767)}
                                                        _ = v31
                                                    default:
                                                        v26 -= (v5 % ((-128) | 1))
                                                    }
                                                }
                                                switch v3 {
                                                case "hello":
                                                    v3 = (v3 + v3)
                                                    v32 := GenS2{F0: fi11}
                                                    _ = v32
                                                case "\n\t":
                                                    if true {
                                                        _ = len(v1)
                                                        fmt.Println(fi11, fi10, v26)
                                                    } else {
                                                        if true {
                                                            if true {
                                                                if (v26 >= (^(v26))) {
                                                                    var v33 int64 = 1
                                                                    v2 -= v2
                                                                    if (true || true) {
                                                                        v34 := GenS1{F0: -9223372036854775808, F1: v33}
                                                                        v19 = 954
                                                                        _ = v34
                                                                    } else {
                                                                        switch v3 {
                                                                        case "foo bar":
                                                                            if len(v4) > 0 {
                                                                                v4[0] = v3
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v33
                                                                }
                                                                v35 := GenS3{F0: "\n\t", F1: (-(v2)), F2: 0}
                                                                var v36 uint32 = 42
                                                                _ = v35
                                                                _ = v36
                                                            } else {
                                                                v37 := GenS3{F0: "\n\t", F1: 2147483647, F2: 2147483648}
                                                                v38 := GenS3{F0: "a", F1: v2, F2: 4294967295}
                                                                var v39 bool = false
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            v5 -= int8(v19)
                                                            v26 = (v26 | -1)
                                                        } else {
                                                            var v40 string = v3
                                                            _ = v40
                                                        }
                                                        if (false || false) {
                                                            v41 := GenS2{F0: fi11}
                                                            var v42 int16 = 1000
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            if false {
                                                                var v43 float64 = 1e10
                                                                v19 = v19
                                                                _ = v43
                                                            } else {
                                                                v6[-1] = 0
                                                                _ = len(v6)
                                                            }
                                                        }
                                                        v2 -= v2
                                                    }
                                                case "a":
                                                    v2 = v2
                                                }
                                            }
                                            var v44 int8 = 42
                                            v45 := v1[v3]
                                            var v46 uint64 = uint64(v19)
                                            var v47 float32 = 0.1
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v48 := GenS4{F0: 1, F1: (v3 + (v3 + v3))}
                                            v26 = (v5 + v5)
                                            var v49 uint64 = 18446744073709551615
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                            var v50 float32 = (3.14159 - (-2.25 * 1e10))
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                        v2 += (v2 - (v2 + v2))
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    fmt.Println(fi10, fi10)
                                }
                                _ = v28
                            }
                            if ((v19 ^ (v19 + (v19 - v19))) == v19) {
                                if true {
                                    var v51 uint32 = 42
                                    _ = v51
                                } else {
                                    fmt.Println(v26, v5)
                                }
                                fmt.Println(fi11)
                            } else {
                                var v52 int64 = -813
                                var v53 float32 = (((3.14159 - 0.1) + 1e10) / ((0.0 * 0.1) - ((0.1 / -2.25) - 0.1)))
                                if ((^(v26)) >= v5) {
                                    v52 = v52
                                    v54 := GenS5{F0: 18446744073709551615, F1: 4294967295, F2: 2147483647}
                                    v26 = v26
                                    _ = v54
                                }
                                _ = v52
                                _ = v53
                            }
                            _ = v26
                        }
                        _ = v19
                    }
                default:
                    fmt.Println(fi10)
                }
                fmt.Println(v2)
                _ = fi11
            }
            v5 -= genFn0(fi10, 42, v5)
            _ = fi10
        }
    } else {
        for fi55 := 0; fi55 < 15; fi55++ {
            fmt.Println(fi55)
            fmt.Println(v2, v5)
            v56 := GenS2{F0: fi55}
            _ = fi55
            _ = v56
        }
        for fi57 := 0; fi57 < 4; fi57++ {
            fmt.Println(v2)
            if false {
                v6[(fi57 % (((fi57 + (fi57 % ((32768) | 1)))) | 1))] = 4294967295
            } else {
                for fi58 := 0; fi58 < 10; fi58++ {
                    if false {
                        var v59 int8 = v5
                        _ = v59
                    } else {
                        v6[128] = 65535
                        var v60 uint = 42
                        _ = v60
                    }
                    _ = len(v4)
                    _ = fi58
                }
                switch fi57 {
                case 0:
                    if false {
                        switch fi57 {
                        case 4294967295:
                            var v61 bool = ((fi57 - fi57) < fi57)
                            _ = v61
                        case 449:
                            switch (-(fi57)) {
                            case 128:
                                v5 += v5
                            case -2147483648:
                                var v62 uint8 = 127
                                _ = v62
                            case 32767:
                                _ = len(v4)
                            }
                        default:
                            v63 := v6[(fi57 << ((319) & 63))]
                            _ = v63
                        }
                        v3 = v3
                    }
                case 32768:
                    var v64 int64 = 1
                    _ = v64
                }
                v65 := GenS6{F0: 262, F1: (false || (100 == 42))}
                _ = v65
            }
            v3 = "hello"
            _ = fi57
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v66 uint16 = 42
    fmt.Println(v66, v2, v5)
    switch ("foo bar" + v3) {
    case "hello":
        v66 -= 42
    default:
        if false {
            v5 -= genFn0(327, int64(v2), 119)
        }
    }
    v67 := GenS7{F0: 1}
    v68 := v1[v3]
    var v69 bool = (!(true || ((!(1e10 >= (1.5 - 1.5))) || true)))
    switch 128 {
    case 9223372036854775807:
        fmt.Println(v2, v3, v3)
    default:
        v1[("foo bar" + (v3 + ("hello" + ("a" + v3))))] = v68
    }
    fmt.Println(v2)
    if (!v69) {
        v70 := GenS8{F0: -1, F1: 18446744073709551615}
        fmt.Println(v5, v68, v3)
        _ = v70
    } else {
        v69 = (42 != 2147483647)
    }
    fmt.Println(v5, v5, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v3 {
    case "":
        switch "\n\t" {
        case "a":
            if v69 {
                v4 = append(v4, "\n\t")
                switch "hello" {
                case "hello":
                    for fi71 := 0; fi71 < 3; fi71++ {
                        v72 := GenS9{F0: v68, F1: genFn1(genFn0((fi71 * 403), 9223372036854775807, (^(v5))), 32767)}
                        switch 1 {
                        case 42:
                            var v73 bool = ((461 == 9223372036854775807) || (1 != 2))
                            _ = v73
                        default:
                            for fi74 := 0; fi74 < 1; fi74++ {
                                var v75 int8 = genFn0((fi74 << ((259) & 63)), int64(v66), (v5 + v5))
                                if v69 {
                                    v75 = (-(v5))
                                    v76 := v6[9223372036854775807]
                                    _ = v76
                                }
                                delete(v1, (("hello" + "\n\t") + ((v3 + ("hello" + ("foo bar" + "hello"))) + ("hello" + v3))))
                                _ = fi74
                                _ = v75
                            }
                        }
                        v3 = v3
                        _ = fi71
                        _ = v72
                    }
                }
            }
        case "hello":
            v2 -= v2
        case "\n\t":
            var v77 float64 = 0.1
            _ = v77
        }
    case "a":
        if (1e10 == 1.5) {
            var v78 uint32 = (^(v68))
            _ = v78
        } else {
            for fi79 := 0; fi79 < 10; fi79++ {
                v68 = (v68 | (^(v68)))
                v69 = false
                var v80 int8 = (v5 >> ((855) & 63))
                _ = fi79
                _ = v80
            }
            switch 128 {
            case -9223372036854775808:
                var v81 int16 = 0
                _ = v81
            }
        }
    default:
        v3 = "a"
    }
    v82 := GenS2{F0: 1}
    if ((!v69) && v69) {
        switch "\n\t" {
        case "hello":
            for fi83 := 0; fi83 < 14; fi83++ {
                for fi84 := 0; fi84 < 10; fi84++ {
                    v69 = v69
                    v85 := GenS8{F0: (fi83 ^ fi83), F1: 2}
                    _ = fi84
                    _ = v85
                }
                _ = fi83
            }
        case "":
            switch v3 {
            case "a":
                for fi86 := 0; fi86 < 6; fi86++ {
                    v5 = v5
                    var v87 int32 = (v2 + 750)
                    _ = fi86
                    _ = v87
                }
            case "foo bar":
                v66 = v66
            }
        default:
            var v88 float32 = 0.1
            _ = v88
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v82
}
