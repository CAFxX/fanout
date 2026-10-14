package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int16; F2 bool }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS9 struct { F0 uint16; F1 uint; F2 uint }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 int16; F1 uint16 }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 uint64; F1 int; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-766, -509, -772}
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    v5 := []string{"a", "foo bar", (("a" + (("\n\t" + "hello") + "\n\t")) + "a"), "foo bar"}
    var v6 int8 = -128
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint32 = 980
        _ = fi7
        _ = v8
    }
    var v9 string = (("hello" + "foo bar") + "")
    var v10 uint8 = 1
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: 4294967295, F1: v9, F2: (v6 / ((-113) | 1))}
        var v13 float32 = 1e10
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int64 = 9223372036854775807
    switch 0 {
    case -2147483648:
        if (0 < 42) {
            var v15 float64 = 3.14159
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint16 = 32767
        switch (fi16 % ((int(v6)) | 1)) {
        case 341:
            var v18 string = ("\n\t" + "hello")
            _ = v18
        case -9223372036854775808:
            fmt.Println(v17, v9)
        case 42:
            delete(v4, ("hello" + ((v9 + v9) + v9)))
        }
        _ = fi16
        _ = v17
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        switch (fi19 << ((210) & 63)) {
        case 1:
            if (4294967295 < 876) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    if true {
                        var v21 int = (fi20 & (fi19 / ((fi20) | 1)))
                        if (fi20 >= int(v6)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch v21 {
                            case -1:
                                v22 := GenS2{F0: 4294967295, F1: true, F2: 32767}
                                _ = v22
                            case 221:
                                switch ("\n\t" + v9) {
                                case "hello":
                                    switch 4294967295 {
                                    case 296:
                                        var v23 int64 = 9223372036854775807
                                        _ = v23
                                    }
                                case "":
                                    if false {
                                        delete(v4, v9)
                                        var v24 int8 = 74
                                        var v25 uint16 = 119
                                        if (1 < 2147483647) {
                                            _ = len(v4)
                                            var v26 bool = ((941 < int16(v10)) && true)
                                            v9 = (v9 + v9)
                                            v4[("foo bar" + v9)] = fi19
                                            _ = v26
                                        }
                                        _ = len(v3)
                                        delete(v3, 65535)
                                        v14 -= 722
                                        var v27 float32 = 0.0
                                        var v28 bool = false
                                        _ = v24
                                        _ = v25
                                        _ = v27
                                        _ = v28
                                    } else {
                                        switch (fi19 & fi20) {
                                        case 128:
                                            v29 := GenS2{F0: 94, F1: true, F2: 67}
                                            var v30 int32 = 2147483647
                                            _ = len(v5)
                                            _ = v29
                                            _ = v30
                                        case -372:
                                            var v31 int8 = 1
                                            _ = v31
                                        case -9223372036854775808:
                                            fmt.Println(v14, v14, v21)
                                        }
                                        v9 = v9
                                        v21 -= 4294967295
                                        if true {
                                            v32 := GenS3{F0: 4294967295, F1: 32767, F2: (!false)}
                                            var v33 float64 = 1e10
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                }
                            default:
                                v21 = fi20
                            }
                            v9 = (("\n\t" + v9) + v9)
                        } else {
                            if (42 > 563) {
                                var v34 float64 = 1e10
                                v35 := GenS3{F0: 1, F1: 1000, F2: false}
                                v9 = (v9 + (((("hello" + ("hello" + v9)) + (v9 + v9)) + (v9 + v9)) + v9))
                                _ = v34
                                _ = v35
                            }
                            v6 -= 42
                            v36 := GenS3{F0: 224, F1: 1000, F2: (false || (v10 == (v10 & 40)))}
                            _ = v36
                        }
                        v14 = 0
                        var v37 bool = true
                        v38 := GenS4{F0: 4294967295, F1: -1}
                        _ = v21
                        _ = v37
                        _ = v38
                    } else {
                        var v39 float32 = 1.5
                        var v40 uint32 = 374
                        v41 := GenS5{F0: v9}
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    v9 = v9
                    v10 = v10
                    _ = fi20
                }
                if (1 != 32767) {
                    v9 = (v9 + "hello")
                } else {
                    v14 = v14
                    for fi42 := 0; fi42 < 4; fi42++ {
                        v43 := GenS1{F0: 0, F1: "a", F2: (^(v6))}
                        switch (fi19 << ((223) & 63)) {
                        case 533:
                            if (65535 > 2) {
                                if (1e10 == ((0.0 - 0.0) / 0.1)) {
                                    v44 := v1[2147483648]
                                    _ = v44
                                }
                                var v45 int16 = int16(v10)
                                if true {
                                    v3[65535] = (v9 + "\n\t")
                                } else {
                                    v46 := GenS6{F0: 0}
                                    fmt.Println(fi42, v14)
                                    if true {
                                        var v47 int32 = 1
                                        v14 -= v14
                                        _ = v47
                                    } else {
                                        var v48 int = (-(fi19))
                                        v48 -= fi42
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if ((v45 * (v45 + (^(v45)))) >= v45) {
                                            v49 := GenS3{F0: 1, F1: v45, F2: ((v9 + "") != v9)}
                                            v50 := GenS1{F0: 42, F1: ("hello" + (v9 + ("" + "a"))), F2: 1}
                                            var v51 uint32 = 42
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                        v45 += v45
                                        v52 := GenS3{F0: 42, F1: v45, F2: false}
                                        if (!false) {
                                            v53 := GenS2{F0: 2, F1: false, F2: 50}
                                            delete(v3, 987)
                                            v9 = "\n\t"
                                            _ = v53
                                        }
                                        _ = v48
                                        _ = v52
                                    }
                                    _ = v46
                                }
                                _ = v45
                            } else {
                                var v54 int8 = (v6 | v6)
                                v4[v9] = fi19
                                var v55 int8 = (v6 / (((v6 >> ((2) & 63))) | 1))
                                _ = v54
                                _ = v55
                            }
                        }
                        var v56 uint = 770
                        _ = fi42
                        _ = v43
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 9; fi57++ {
                        switch (fi57 % ((4294967295) | 1)) {
                        case 0:
                            var v58 string = ("foo bar" + "foo bar")
                            _ = v58
                        default:
                            if (!true) {
                                v59 := GenS7{F0: (18446744073709551615 <= 18446744073709551615)}
                                _ = v59
                            } else {
                                delete(v4, ("\n\t" + v9))
                            }
                            var v60 int64 = v14
                            _ = v60
                        }
                        if ((fi19 / ((fi19) | 1)) == (^(fi57))) {
                            v61 := GenS8{F0: 1, F1: 2, F2: 32767}
                            v6 += 58
                            var v62 float32 = (3.14159 / 0.1)
                            if false {
                                var v63 uint64 = 18446744073709551615
                                _ = v63
                            }
                            _ = v61
                            _ = v62
                        }
                        switch (-(fi19)) {
                        case 127:
                            v64 := GenS8{F0: 329, F1: 201, F2: 290}
                            _ = v64
                        case 1:
                            switch v9 {
                            case "hello":
                                if (true || true) {
                                    var v65 uint16 = 65535
                                    _ = v65
                                } else {
                                    if false {
                                        v66 := GenS9{F0: 42, F1: 65535, F2: 42}
                                        _ = v66
                                    } else {
                                        if false {
                                            v10 += 0
                                            if false {
                                                v14 = v14
                                                var v67 int64 = v14
                                                _ = v67
                                            }
                                            if false {
                                                v2 = append(v2, fi57)
                                                var v68 uint32 = 1
                                                switch v9 {
                                                case "":
                                                    var v69 uint = 4294967295
                                                    _ = v69
                                                case "a":
                                                    var v70 float64 = float64((((1.5 * 0.0) / 1.0) - (1e10 * (float32(-2.25) * 1.5))))
                                                    _ = v70
                                                case "foo bar":
                                                    var v71 int32 = -407
                                                    _ = v71
                                                }
                                                _ = v68
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (fi57 ^ -2147483648)
                                            }
                                        } else {
                                            if true {
                                                var v72 int = (fi19 * -117)
                                                fmt.Println(v10, v14)
                                                _ = v72
                                            } else {
                                                var v73 int = (-(fi57))
                                                fmt.Println(v9, fi19)
                                                v74 := GenS10{F0: 0, F1: v10}
                                                _ = v73
                                                _ = v74
                                            }
                                            v9 = v9
                                            v9 = ((v9 + ((v9 + (v9 + ("foo bar" + v9))) + "hello")) + "foo bar")
                                        }
                                    }
                                    v1[4294967295] = false
                                    v75 := GenS11{F0: 792, F1: 1}
                                    _ = v75
                                }
                            }
                        default:
                            var v76 uint = 65535
                            _ = v76
                        }
                        _ = fi57
                    }
                }
            }
        default:
            var v77 int32 = 297
            _ = v77
        }
        if true {
            v10 = (v10 >> ((0) & 63))
        }
        _ = fi19
    }
    var v78 uint16 = 255
    v9 = v9
    v9 = ("a" + v9)
    v78 -= v78
    fmt.Println(v10)
    if (v10 < 42) {
        for fi79 := 0; fi79 < 14; fi79++ {
            v78 -= (v78 << ((2) & 63))
            switch (v9 + v9) {
            case "\n\t":
                var v80 int64 = v14
                _ = v80
            case "":
                switch (-(fi79)) {
                case 1:
                    v9 = "foo bar"
                }
            }
            _ = fi79
        }
        for fi81 := 0; fi81 < 8; fi81++ {
            if true {
                var v82 float64 = 1.5
                v78 -= v78
                _ = v82
            } else {
                v6 = (v6 | -63)
                switch v9 {
                case "hello":
                    switch (fi81 | 812) {
                    case 128:
                        v83 := GenS12{F0: (-(fi81)), F1: (0.0 == (0.0 - 1e10))}
                        _ = v83
                    }
                case "\n\t":
                    v1[0] = (3.14159 < 0.0)
                case "a":
                    var v84 uint8 = v10
                    _ = v84
                default:
                    v6 -= v6
                }
                v85 := GenS13{F0: 4294967295, F1: fi81, F2: fi81}
                _ = v85
            }
            _ = fi81
        }
        v78 = 2
    } else {
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v9
}
