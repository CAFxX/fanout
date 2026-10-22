package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (1e10 * 1e10)
    var v3 int16 = -1
    v4 := make(map[uint32]int)
    var v5 int64 = 0
    var v6 uint = 65535
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 float64 = float64(v2)
            if (v2 != 1e10) {
                var v10 string = ("" + "hello")
                v10 = ((v10 + "foo bar") + ("foo bar" + v10))
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        switch "foo bar" {
        case "\n\t":
            var v11 int = fi7
            _ = v11
        case "hello":
            var v12 float64 = 3.14159
            _ = v12
        case "a":
            var v13 float32 = v2
            _ = v13
        }
        v14 := v1[98]
        _ = fi7
        _ = v14
    }
    var v15 uint32 = 65535
    v15 += uint32(v6)
    switch 128 {
    case 32767:
        v15 += (v15 ^ uint32(v6))
    case 128:
        v2 = 0.0
    default:
        var v16 int64 = v5
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int32 = -1
        if (!false) {
            switch "foo bar" {
            case "\n\t":
                v6 += 0
            case "hello":
                v18 -= v18
            default:
                if false {
                    if (v5 > int64(v18)) {
                        switch 762 {
                        case 1:
                            var v19 float32 = 1e10
                            _ = v19
                        case -205:
                            v5 = -802
                        case -376:
                            v6 = (v6 ^ v6)
                        }
                    } else {
                        var v20 uint32 = 2147483648
                        switch int(v18) {
                        case 128:
                            v21 := GenS1{F0: "hello", F1: ""}
                            _ = v21
                        case -9223372036854775808:
                            if true {
                                v3 = v3
                                if true {
                                    v15 = 361
                                    fmt.Println(v2)
                                } else {
                                    var v22 bool = false
                                    _ = v22
                                }
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    var v24 int8 = -76
                                    v25 := GenS2{F0: 65535, F1: (("a" + (("foo bar" + ("a" + "foo bar")) + "a")) + "foo bar"), F2: 42}
                                    var v26 uint32 = (v15 / ((4294967295) | 1))
                                    fmt.Println(v2, fi23, v26)
                                    if ((0.0 * v2) < (0.0 + 1.5)) {
                                        if ((v18 << ((v6) & 63)) != (^(v18))) {
                                            var v27 string = "a"
                                            var v28 int32 = v18
                                            _ = v27
                                            _ = v28
                                        } else {
                                            fmt.Println(fi23, v26)
                                        }
                                        if true {
                                            var v29 float64 = 3.14159
                                            switch "hello" {
                                            case "a":
                                                v26 -= (v26 + 4294967295)
                                            case "\n\t":
                                                var v30 string = (("a" + "a") + "\n\t")
                                                _ = v30
                                            }
                                            v4[v15] = (fi23 >> ((136) & 63))
                                            _ = v29
                                        } else {
                                            var v31 float32 = v2
                                            if (!true) {
                                                switch "\n\t" {
                                                case "":
                                                    var v32 float32 = (v2 / float32(3.14159))
                                                    _ = v32
                                                case "foo bar":
                                                    var v33 uint64 = 0
                                                    _ = v33
                                                case "hello":
                                                    v20 -= (v26 + (v15 ^ 562))
                                                    fmt.Println(v20)
                                                default:
                                                    v34 := []int{fi23}
                                                    _ = v34
                                                }
                                                v35 := GenS2{F0: 32767, F1: (((("hello" + (("a" + "foo bar") + ("hello" + ""))) + "hello") + ((("" + "a") + (("" + "\n\t") + "a")) + "\n\t")) + "\n\t"), F2: 218}
                                                _ = v35
                                            }
                                            v15 += 2147483648
                                            v18 += v18
                                            v4[4294967295] = fi17
                                            v24 = v24
                                            v36 := GenS2{F0: 825, F1: "foo bar", F2: 119}
                                            _ = v31
                                            _ = v36
                                        }
                                    }
                                    if (!true) {
                                        v37 := GenS3{F0: 2, F1: true}
                                        _ = v37
                                    } else {
                                        var v38 int32 = v18
                                        v6 -= (v6 + (v6 % (((^(v6))) | 1)))
                                        _ = v38
                                    }
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            }
                        default:
                            var v39 int16 = (v3 << ((v6) & 63))
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 20; fi40++ {
                            v15 = (v15 & (^(v15)))
                            v41 := GenS3{F0: 65535, F1: true}
                            _ = fi40
                            _ = v41
                        }
                        _ = v20
                    }
                }
            }
        } else {
            v42 := GenS1{F0: "a", F1: ("a" + (("" + (("" + "hello") + "\n\t")) + "\n\t"))}
            _ = v42
        }
        _ = fi17
        _ = v18
    }
    switch -1 {
    case 1:
        v5 += -1
    }
    var v43 float64 = (0.0 * -2.25)
    if false {
        switch -1 {
        case 9223372036854775807:
            switch "hello" {
            case "foo bar":
                v5 = 743
            }
        default:
            if true {
                v44 := GenS4{F0: v6, F1: 2}
                switch "hello" {
                case "":
                    v3 -= (^(v3))
                default:
                    var v45 string = "foo bar"
                    _ = v45
                }
                for fi46 := 0; fi46 < 14; fi46++ {
                    var v47 int32 = -2147483648
                    _ = fi46
                    _ = v47
                }
                _ = v44
            } else {
                var v48 int = 4294967295
                switch v48 {
                case 9223372036854775807:
                    v49 := []int{v48}
                    _ = v49
                case 1:
                    switch "\n\t" {
                    case "hello":
                        if false {
                            v50 := GenS2{F0: 163, F1: "foo bar", F2: 0}
                            fmt.Println(v6, v43, v5)
                            _ = v50
                        } else {
                            v51 := GenS5{F0: 42, F1: 217}
                            fmt.Println(v6, v2)
                            _ = v51
                        }
                    case "a":
                        v4[v15] = 131
                    case "":
                        var v52 int8 = 127
                        _ = v52
                    default:
                        if true {
                            v6 = 65535
                            v53 := []int{(v48 + (v48 << ((v6) & 63))), (v48 >> (((v6 << ((uint(v15)) & 63))) & 63)), -2147483648, int(v5)}
                            v48 = (v48 << (((v6 << (((v6 & (v6 & (v6 >> ((588) & 63))))) & 63))) & 63))
                            _ = v53
                        } else {
                            v48 -= v48
                            for fi54 := 0; fi54 < 12; fi54++ {
                                if false {
                                    var v55 bool = ((false && true) || false)
                                    _ = v55
                                }
                                if (!true) {
                                    for fi56 := 0; fi56 < 7; fi56++ {
                                        if false {
                                            var v57 uint8 = 223
                                            if (false || (v6 <= (v6 + (v6 + 1)))) {
                                                v58 := GenS2{F0: 65535, F1: ("" + ""), F2: (^(v57))}
                                                _ = v58
                                            } else {
                                                if false {
                                                    v57 += v57
                                                    v59 := []int{}
                                                    _ = v59
                                                }
                                                fmt.Println(fi56)
                                                var v60 uint8 = v57
                                                _ = v60
                                            }
                                            _ = v57
                                        }
                                        v61 := GenS6{F0: 32768, F1: true}
                                        _ = fi56
                                        _ = v61
                                    }
                                    v62 := v1[886]
                                    fmt.Println(v5)
                                    _ = v62
                                } else {
                                    for fi63 := 0; fi63 < 1; fi63++ {
                                        if true {
                                            if false {
                                                v2 = v2
                                                fmt.Println(v15)
                                                if (v6 <= 4294967295) {
                                                    v15 += v15
                                                    var v64 uint16 = uint16(v15)
                                                    _ = v64
                                                } else {
                                                    v65 := GenS3{F0: 1, F1: true}
                                                    if false {
                                                        _ = len(v1)
                                                    }
                                                    _ = v65
                                                }
                                                if (false || (42 == -1)) {
                                                    fmt.Println(fi54, v43)
                                                    v66 := GenS4{F0: (v6 | (v6 - (v6 - (v6 * (v6 + (v6 / ((750) | 1))))))), F1: 2}
                                                    _ = v66
                                                }
                                                var v67 uint32 = v15
                                                v68 := []int{fi63, fi63, int(v5), fi63}
                                                _ = v67
                                                _ = v68
                                            }
                                            v5 -= v5
                                        }
                                        _ = fi63
                                    }
                                }
                                if false {
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        var v69 uint32 = 0
                                        _ = v69
                                    case "":
                                        for fi70 := 0; fi70 < 3; fi70++ {
                                            switch "a" {
                                            case "":
                                                v71 := GenS7{F0: 42, F1: 44}
                                                _ = v71
                                            }
                                            _ = fi70
                                        }
                                    case "foo bar":
                                        var v72 uint16 = 65535
                                        _ = v72
                                    default:
                                        var v73 float32 = v2
                                        _ = v73
                                    }
                                } else {
                                    var v74 uint16 = 1
                                    for fi75 := 0; fi75 < 1; fi75++ {
                                        v2 = (v2 - float32(((-2.25 / v43) + (v43 + float64((v2 - float32(1.5)))))))
                                        v6 = 4294967295
                                        _ = fi75
                                    }
                                    _ = v74
                                }
                                _ = fi54
                            }
                        }
                    }
                default:
                    var v76 int64 = v5
                    _ = v76
                }
                delete(v1, 2147483647)
                _ = v48
            }
        }
        if ((!(128 >= 2147483647)) || true) {
            v2 = (1e10 * v2)
        }
    }
    v15 += 892
    v77 := []int{}
    fmt.Println(v6)
    v5 -= (v5 & (-(v5)))
    v15 += 80
    var v78 uint16 = 964
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v77
    _ = v78
}
