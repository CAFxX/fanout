package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int64; F1 int8; F2 int64 }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 int; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("a" + ("hello" + "a"))
    v3 := make(map[int]string)
    v4 := []int{-81}
    v2 = ("foo bar" + v2)
    v2 = (((v2 + "\n\t") + v2) + "hello")
    _ = len(v4)
    v2 = v2
    v4 = append(v4, 476)
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                var v7 int8 = 100
                var v8 int16 = 32767
                switch (fi6 | (fi5 + (fi6 - -765))) {
                case 42:
                    fmt.Println(v2, v7)
                case 32767:
                    var v9 uint8 = 42
                    _ = v9
                }
                var v10 int = (fi6 & fi5)
                _ = v10
                _ = v7
                _ = v8
            } else {
                if (!(!(false || false))) {
                    v2 = "a"
                    if (float64((-2.25 + -2.25)) <= -2.25) {
                        if true {
                            if false {
                                var v11 uint32 = 362
                                _ = v11
                            }
                            if false {
                                var v12 uint16 = 65535
                                _ = v12
                            } else {
                                var v13 float64 = 0.0
                                var v14 bool = false
                                if v14 {
                                    var v15 int32 = 0
                                    v16 := GenS1{F0: 145}
                                    v1[577] = v2
                                    _ = v15
                                    _ = v16
                                } else {
                                    v14 = false
                                    v17 := GenS2{F0: 127, F1: (!false), F2: 1}
                                    _ = v17
                                }
                                _ = v13
                                _ = v14
                            }
                            switch v2 {
                            case "hello":
                                if true {
                                    v2 = (("foo bar" + v2) + v2)
                                } else {
                                    v2 = "a"
                                    switch (fi5 / (((-(fi5))) | 1)) {
                                    case 32767:
                                        if true {
                                            v2 = "foo bar"
                                            switch (^(fi5)) {
                                            case 2147483647:
                                                v18 := GenS3{F0: 0, F1: 1, F2: ("\n\t" + v2)}
                                                _ = v18
                                            case 32767:
                                                fmt.Println(fi6, v2, fi6)
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = (fi6 - (fi5 ^ fi6))
                                            }
                                        }
                                    case -223:
                                        var v19 int = 1
                                        _ = v19
                                    default:
                                        var v20 int64 = 0
                                        _ = v20
                                    }
                                    v21 := v1[167]
                                    _ = v21
                                }
                            case "\n\t":
                                var v22 uint64 = 4294967295
                                v23 := GenS1{F0: 9223372036854775807}
                                _ = v22
                                _ = v23
                            default:
                                v2 = v2
                            }
                        } else {
                            var v24 int32 = 42
                            _ = v24
                        }
                        var v25 int8 = 42
                        _ = v25
                    }
                }
                _ = len(v4)
                v2 = "hello"
                if true {
                    v26 := GenS3{F0: 127, F1: 1, F2: (((((v2 + (v2 + "a")) + (v2 + v2)) + "foo bar") + ((("a" + v2) + ("" + ((v2 + ("hello" + ((v2 + (v2 + "hello")) + v2))) + "hello"))) + ("" + "a"))) + v2)}
                    _ = v26
                } else {
                    var v27 int64 = 9223372036854775807
                    if false {
                        switch (fi6 + fi5) {
                        case 63:
                            if true {
                                fmt.Println(v2, fi6)
                            } else {
                                switch 42 {
                                case -9223372036854775808:
                                    v27 = (v27 >> ((4294967295) & 63))
                                case 128:
                                    v1[2] = (("" + ("hello" + ("hello" + "\n\t"))) + "\n\t")
                                }
                                var v28 string = "foo bar"
                                _ = v28
                            }
                        case 127:
                            v29 := GenS3{F0: 127, F1: 576, F2: v2}
                            _ = v29
                        case 9223372036854775807:
                            v30 := GenS2{F0: 1, F1: ((0 <= v27) || false), F2: 0}
                            _ = v30
                        }
                        fmt.Println(v27, fi6)
                    } else {
                        v31 := GenS3{F0: 33, F1: 939, F2: (v2 + v2)}
                        switch "\n\t" {
                        case "foo bar":
                            v32 := GenS1{F0: v27}
                            _ = v32
                        case "hello":
                            switch fi5 {
                            case 164:
                                v27 = v27
                            case 4294967295:
                                if (42 != 2147483647) {
                                    if false {
                                        fmt.Println(fi5, fi5)
                                        v33 := GenS4{F0: v27, F1: 2}
                                        var v34 int = 754
                                        if (2 > 0) {
                                            v4 = append(v4, 1)
                                            if false {
                                                fmt.Println(v34)
                                                v35 := GenS2{F0: 128, F1: true, F2: 42}
                                                _ = v35
                                            }
                                            fmt.Println(v34)
                                        } else {
                                            if (42 != 1) {
                                                fmt.Println(fi5, fi5, v34)
                                            }
                                        }
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v36 := v1[1]
                                        v27 += v27
                                        var v37 int = 127
                                        if len(v4) > 0 {
                                            v4[0] = v37
                                        }
                                        v27 -= v27
                                        v38 := GenS5{F0: 42, F1: (fi5 - fi5)}
                                        var v39 uint64 = 42
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                    if true {
                                        v2 = "hello"
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        fmt.Println(fi5)
                                    }
                                    if (1 > -9223372036854775808) {
                                        if true {
                                            v40 := GenS4{F0: v27, F1: 2}
                                            if ((-567 > (v27 & 1)) && (51 != 127)) {
                                                v27 += (v27 << ((2) & 63))
                                                v27 -= -393
                                                delete(v3, -7)
                                            } else {
                                                switch 0 {
                                                case 4294967295:
                                                    if len(v4) > 0 {
                                                        v4[0] = 2147483647
                                                    }
                                                case 546:
                                                    v41 := GenS1{F0: 461}
                                                    _ = v41
                                                case 128:
                                                    if (1000 < 1000) {
                                                        var v42 uint = 42
                                                        var v43 string = v2
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        var v44 float64 = ((0.0 - 1e10) / 1e10)
                                                        v45 := v1[611]
                                                        fmt.Println(v44, v2)
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    var v46 float64 = 0.1
                                                    _ = v46
                                                default:
                                                    v27 -= 126
                                                }
                                            }
                                            _ = v40
                                        } else {
                                            switch fi5 {
                                            case 1:
                                                if false {
                                                    if false {
                                                        if (-2147483648 <= 1) {
                                                            var v47 int = 350
                                                            _ = v47
                                                        }
                                                        v48 := v1[1]
                                                        _ = v48
                                                    }
                                                }
                                            case 4294967295:
                                                var v49 int = fi6
                                                _ = v49
                                            case 2147483647:
                                                v50 := GenS6{F0: (v27 ^ -9223372036854775808), F1: -74, F2: (v27 - v27)}
                                                _ = v50
                                            }
                                        }
                                    }
                                    v2 = "\n\t"
                                    switch fi5 {
                                    case 4294967295:
                                        var v51 uint32 = 2147483648
                                        _ = v51
                                    }
                                } else {
                                    var v52 uint32 = 65535
                                    if false {
                                        switch fi6 {
                                        case 2147483647:
                                            v53 := GenS7{F0: 2, F1: 124}
                                            if (true || true) {
                                                _ = len(v1)
                                                var v54 int32 = -1
                                                _ = v54
                                            }
                                            var v55 uint = 378
                                            v55 -= v55
                                            _ = v53
                                            _ = v55
                                        case 128:
                                            v52 += 42
                                        case -2147483648:
                                            v56 := GenS8{F0: fi6, F1: "foo bar"}
                                            _ = v56
                                        default:
                                            fmt.Println(v52)
                                        }
                                        v2 = (v2 + (("a" + ("a" + (v2 + "a"))) + "\n\t"))
                                        v3[fi6] = ("foo bar" + v2)
                                    } else {
                                        var v57 int8 = -1
                                        var v58 uint = 135
                                        _ = v57
                                        _ = v58
                                    }
                                    _ = v52
                                }
                            }
                        default:
                            v59 := v3[(^(fi5))]
                            _ = v59
                        }
                        _ = v31
                    }
                    var v60 int8 = -36
                    _ = v27
                    _ = v60
                }
            }
            v61 := GenS2{F0: 2, F1: (!false), F2: 127}
            fmt.Println(v2, fi5, v2)
            _ = fi6
            _ = v61
        }
        _ = fi5
    }
    for fi62 := 0; fi62 < 6; fi62++ {
        var v63 int16 = -823
        for fi64 := 0; fi64 < 6; fi64++ {
            var v65 float32 = (float32(1e10) * -2.25)
            v66 := v1[2147483648]
            _ = fi64
            _ = v65
            _ = v66
        }
        var v67 float64 = ((((3.14159 / 1.5) - 3.14159) + -2.25) + 1e10)
        _ = fi62
        _ = v63
        _ = v67
    }
    for fi68 := 0; fi68 < 18; fi68++ {
        var v69 int16 = 416
        var v70 bool = true
        _ = fi68
        _ = v69
        _ = v70
    }
    v2 = (v2 + "foo bar")
    v71 := GenS4{F0: 42, F1: 98}
    switch v2 {
    case "":
        v2 = "hello"
    case "foo bar":
        var v72 float64 = (0.0 - 1e10)
        _ = v72
    default:
        var v73 int64 = -1
        _ = v73
    }
    _ = len(v3)
    fmt.Println(v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
}
