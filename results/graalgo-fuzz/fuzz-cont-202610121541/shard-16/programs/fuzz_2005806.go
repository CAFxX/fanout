package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int64; F2 int8 }
type GenS5 struct { F0 int8; F1 uint; F2 string }
type GenS6 struct { F0 int8; F1 int32; F2 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 int16; F1 uint64; F2 int32 }
func genFn0(p0 uint, p1 bool, p2 int32) uint {
    for fi7 := 0; fi7 < 15; fi7++ {
        p2 = p2
        _ = fi7
    }
    var v8 float64 = -2.25
    v9 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int64 = 0
    var v2 string = "hello"
    v3 := make(map[string]int)
    var v4 float64 = 3.14159
    v5 := make(map[int]string)
    v6 := make(map[string]int)
    if true {
        for fi10 := 0; fi10 < 4; fi10++ {
            if ((18446744073709551615 >= 65535) || true) {
                var v11 int32 = 0
                _ = v11
            }
            v12 := GenS1{F0: 2}
            _ = fi10
            _ = v12
        }
        var v13 uint32 = 4294967295
        _ = v13
    } else {
        fmt.Println(v2)
        if (!(-1 >= -1)) {
            if (42 <= 42) {
                v1 = v1
                v2 = v2
            } else {
                v14 := []int{int(v1)}
                var v15 uint8 = 2
                if true {
                    switch 9223372036854775807 {
                    case -333:
                        if false {
                            for fi16 := 0; fi16 < 12; fi16++ {
                                if ((true || (-128 >= 58)) || true) {
                                    if true {
                                        var v17 float64 = 3.14159
                                        if false {
                                            if (18446744073709551615 <= 336) {
                                                v18 := GenS1{F0: 42}
                                                v1 = (v1 ^ (v1 - v1))
                                                _ = v18
                                            }
                                            _ = len(v5)
                                            v14 = append(v14, fi16)
                                        } else {
                                            var v19 uint16 = 42
                                            if (true || false) {
                                                v20 := GenS2{F0: 42, F1: ("a" + v2)}
                                                _ = v20
                                            }
                                            _ = v19
                                        }
                                        _ = v17
                                    } else {
                                        v21 := v3[(v2 + "\n\t")]
                                        v22 := GenS2{F0: 565, F1: v2}
                                        _ = v21
                                        _ = v22
                                    }
                                    v2 = "foo bar"
                                    var v23 uint16 = 255
                                    _ = v23
                                } else {
                                    v3["hello"] = -9223372036854775808
                                }
                                _ = fi16
                            }
                            if ((65 != 730) && ((!((v15 * 2) <= 1)) && true)) {
                                v1 = 9223372036854775807
                                if (false && ((-32768 == 1000) || false)) {
                                    v24 := GenS1{F0: 1}
                                    v2 = "a"
                                    _ = v24
                                }
                                v25 := GenS3{F0: -2147483648}
                                _ = v25
                            }
                        }
                    default:
                        if true {
                            v26 := v5[900]
                            for fi27 := 0; fi27 < 20; fi27++ {
                                v28 := GenS4{F0: (fi27 & (fi27 & 842)), F1: int64(v15), F2: -66}
                                _ = fi27
                                _ = v28
                            }
                            var v29 int8 = 1
                            _ = v26
                            _ = v29
                        } else {
                            switch 1 {
                            case 42:
                                switch v2 {
                                case "a":
                                    switch -2147483648 {
                                    case 127:
                                        for fi30 := 0; fi30 < 20; fi30++ {
                                            for fi31 := 0; fi31 < 1; fi31++ {
                                                v32 := GenS2{F0: 632, F1: v2}
                                                delete(v6, v2)
                                                v2 = v2
                                                _ = fi31
                                                _ = v32
                                            }
                                            for fi33 := 0; fi33 < 10; fi33++ {
                                                switch v2 {
                                                case "":
                                                    v4 -= v4
                                                case "\n\t":
                                                    v6[v2] = (fi33 ^ -473)
                                                case "hello":
                                                    v14 = append(v14, fi33)
                                                default:
                                                    var v34 uint32 = 1
                                                    var v35 int16 = 0
                                                    _ = v34
                                                    _ = v35
                                                }
                                                switch -989 {
                                                case 2147483647:
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(v2, v4, v1)
                                                    case "":
                                                        switch fi33 {
                                                        case 128:
                                                            v2 = (v2 + ("" + (v2 + v2)))
                                                        case 1:
                                                            var v36 uint16 = 197
                                                            _ = v36
                                                        }
                                                    case "hello":
                                                        v2 = (v2 + (v2 + v2))
                                                    default:
                                                        if ((-611 <= v1) && (0.1 == ((1e10 - (3.14159 - 1.5)) / 1.5))) {
                                                            var v37 int32 = 1
                                                            switch fi33 {
                                                            case 2147483647:
                                                                v4 -= 0.0
                                                            case 219:
                                                                if len(v14) >= 2 {
                                                                    v14 = v14[1:len(v14)]
                                                                }
                                                            default:
                                                                var v38 int16 = 32767
                                                                _ = v38
                                                            }
                                                            _ = v37
                                                        } else {
                                                            v39 := GenS5{F0: -61, F1: 4294967295, F2: "\n\t"}
                                                            var v40 uint = 1
                                                            var v41 float64 = 1.5
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        v6["foo bar"] = (fi30 - (-(fi30)))
                                                    }
                                                case 42:
                                                    fmt.Println(v2)
                                                case 4294967295:
                                                    fmt.Println(fi30, v2)
                                                }
                                                if true {
                                                    var v42 uint32 = 637
                                                    _ = v42
                                                } else {
                                                    v1 = (v1 ^ (v1 * v1))
                                                }
                                                var v43 int = fi30
                                                switch v2 {
                                                case "hello":
                                                    var v44 uint64 = 1
                                                    _ = v44
                                                default:
                                                    v45 := GenS1{F0: 42}
                                                    _ = v45
                                                }
                                                var v46 bool = true
                                                v47 := GenS6{F0: 100, F1: 42, F2: 255}
                                                _ = fi33
                                                _ = v43
                                                _ = v46
                                                _ = v47
                                            }
                                            _ = fi30
                                        }
                                    case -711:
                                        fmt.Println(v2)
                                    default:
                                        v2 = v2
                                    }
                                }
                            case 128:
                                for fi48 := 0; fi48 < 7; fi48++ {
                                    if (65535 != 2147483647) {
                                        v49 := GenS2{F0: 684, F1: v2}
                                        v50 := v6[("a" + "a")]
                                        switch (-(v50)) {
                                        case 0:
                                            var v51 int32 = 0
                                            _ = v51
                                        case 4294967295:
                                            for fi52 := 0; fi52 < 4; fi52++ {
                                                v4 -= (0.1 - v4)
                                                fmt.Println(v15, v4, v50)
                                                switch v2 {
                                                case "":
                                                    if len(v14) > 0 {
                                                        v14[0] = (fi52 + 680)
                                                    }
                                                case "\n\t":
                                                    var v53 int = (fi52 << ((65535) & 63))
                                                    _ = v53
                                                default:
                                                    switch v2 {
                                                    case "":
                                                        delete(v6, v2)
                                                    case "\n\t":
                                                        var v54 int32 = 0
                                                        _ = v54
                                                    case "foo bar":
                                                        v15 -= v15
                                                    }
                                                }
                                                _ = fi52
                                            }
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                    var v55 uint64 = 18446744073709551615
                                    _ = fi48
                                    _ = v55
                                }
                            }
                            switch 73 {
                            case -699:
                                v56 := GenS2{F0: 711, F1: v2}
                                _ = v56
                            case 9223372036854775807:
                                v15 = (v15 - (v15 | (v15 % (((v15 * (v15 << ((782) & 63)))) | 1))))
                            }
                            for fi57 := 0; fi57 < 1; fi57++ {
                                switch ("a" + "") {
                                case "foo bar":
                                    v4 -= v4
                                case "\n\t":
                                    for fi58 := 0; fi58 < 6; fi58++ {
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        v14 = append(v14, 42)
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        _ = fi58
                                    }
                                case "hello":
                                    v59 := GenS6{F0: -7, F1: -2147483648, F2: 65535}
                                    _ = v59
                                default:
                                    fmt.Println(v15, v4, v4)
                                }
                                var v60 uint8 = (v15 - v15)
                                if false {
                                    v5[fi57] = v2
                                    var v61 string = v2
                                    v4 = v4
                                    _ = v61
                                }
                                _ = fi57
                                _ = v60
                            }
                        }
                    }
                    delete(v5, 654)
                } else {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                _ = v14
                _ = v15
            }
        } else {
            v1 = -9223372036854775808
            var v62 int16 = 1000
            _ = v62
        }
    }
    if (!(!(1 != 42))) {
        var v63 int32 = 25
        v64 := GenS1{F0: 65535}
        _ = v63
        _ = v64
    }
    v65 := GenS7{F0: false, F1: 1}
    v66 := GenS4{F0: 364, F1: (-(v1)), F2: -1}
    switch -2147483648 {
    case -9223372036854775808:
        if (v2 <= (v2 + "")) {
            for fi67 := 0; fi67 < 7; fi67++ {
                v68 := GenS8{F0: 1000, F1: 0, F2: -2147483648}
                _ = fi67
                _ = v68
            }
            fmt.Println(v4)
            if true {
                if ((135 < 1) && (!((v1 + -9223372036854775808) > (v1 - -1)))) {
                    v4 = v4
                }
            } else {
                var v69 uint64 = 4294967295
                fmt.Println(v2, v69, v4)
                v69 += 869
                _ = v69
            }
        }
    }
    v1 -= (v1 * v1)
    switch 32768 {
    case -820:
        if true {
            switch (v2 + ("" + "hello")) {
            case "foo bar":
                v70 := []int{}
                _ = v70
            case "\n\t":
                var v71 float64 = v4
                _ = v71
            case "a":
                v4 += 3.14159
            }
            v72 := []int{32768}
            _ = v72
        }
    }
    fmt.Println(v4, v2)
    v2 = ("a" + v2)
    var v73 int = 32767
    v74 := GenS1{F0: 2}
    var v75 uint32 = 65535
    for fi76 := 0; fi76 < 4; fi76++ {
        var v77 string = "a"
        fmt.Println(v75, v4)
        fmt.Println(v4, v77, v2)
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v73
    _ = v74
    _ = v75
}
