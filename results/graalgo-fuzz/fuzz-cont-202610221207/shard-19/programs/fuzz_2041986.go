package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS8 struct { F0 int32; F1 int; F2 string }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 string }

func main() {
    var v1 uint64 = 900
    var v2 int64 = -1
    var v3 uint16 = 119
    var v4 uint64 = (v1 & v1)
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{fi6, (-(fi6))}
            var v8 int64 = v2
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint64 = 157
        if (true && true) {
            v10 := GenS1{F0: (!(true && true)), F1: v2, F2: (fi5 <= (fi5 * fi5))}
            switch ("" + "") {
            case "":
                switch ("\n\t" + "hello") {
                case "foo bar":
                    for fi11 := 0; fi11 < 2; fi11++ {
                        var v12 int = fi11
                        var v13 float64 = (0.0 + float64(float32((float64(1.5) + (float64(-2.25) * 3.14159)))))
                        v1 += 1
                        v14 := GenS2{F0: v2, F1: -1}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case "\n\t":
                    if true {
                        if false {
                            v15 := []int{}
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v3)
                            case "foo bar":
                                v16 := GenS3{F0: -1}
                                _ = v16
                            default:
                                var v17 uint64 = v4
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 11; fi18++ {
                                v9 += 4294967295
                                _ = fi18
                            }
                            _ = v15
                        } else {
                            fmt.Println(v9)
                            var v19 int32 = 1
                            _ = v19
                        }
                        if true {
                            v20 := GenS2{F0: (^(v2)), F1: 0}
                            _ = v20
                        } else {
                            v4 = v9
                            v21 := GenS3{F0: 9223372036854775807}
                            _ = v21
                        }
                        v22 := GenS1{F0: false, F1: v2, F2: (("hello" + "hello") >= "\n\t")}
                        _ = v22
                    }
                case "hello":
                    if true {
                        v23 := GenS3{F0: fi5}
                        v3 += (v3 | (^(v3)))
                        v24 := []int{fi5, (fi5 - fi5), fi5, fi5}
                        _ = v23
                        _ = v24
                    } else {
                        switch ("" + (("\n\t" + "hello") + ("a" + "hello"))) {
                        case "\n\t":
                            switch fi5 {
                            case -432:
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    var v26 float32 = 1.5
                                    switch fi25 {
                                    case -776:
                                        var v27 uint64 = (v1 * (v1 << ((42) & 63)))
                                        _ = v27
                                    default:
                                        if (148 >= 108) {
                                            var v28 uint64 = v4
                                            _ = v28
                                        }
                                    }
                                    v29 := GenS4{F0: (v2 >> ((278) & 63))}
                                    _ = fi25
                                    _ = v26
                                    _ = v29
                                }
                            }
                        case "hello":
                            v30 := []int{fi5}
                            _ = v30
                        }
                        v31 := []int{42, fi5, (^(fi5))}
                        _ = v31
                    }
                default:
                    var v32 uint16 = (v3 & v3)
                    _ = v32
                }
            case "foo bar":
                var v33 bool = true
                _ = v33
            case "\n\t":
                for fi34 := 0; fi34 < 3; fi34++ {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        var v35 int32 = -945
                        _ = v35
                    }
                    v36 := GenS4{F0: 0}
                    v9 = v1
                    var v37 uint16 = (v3 - (v3 / (((v3 - (v3 | v3))) | 1)))
                    _ = fi34
                    _ = v36
                    _ = v37
                }
            }
            _ = v10
        }
        _ = fi5
        _ = v9
    }
    v4 += v4
    if true {
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 string = "a"
            if (((!(-1 == 2147483647)) && true) && ((489 == fi38) && true)) {
                v40 := GenS5{F0: 920}
                _ = v40
            } else {
                for fi41 := 0; fi41 < 18; fi41++ {
                    v2 = (v2 << ((1) & 63))
                    if (-2147483648 != 2147483647) {
                        v42 := GenS4{F0: 1}
                        if false {
                            if (-9223372036854775808 > 251) {
                                var v43 bool = (334 < 1)
                                switch fi38 {
                                case -406:
                                    if (!(!((!(!v43)) && true))) {
                                        fmt.Println(v3, v1)
                                    } else {
                                        var v44 int = fi38
                                        v3 -= (v3 % ((32767) | 1))
                                        _ = v44
                                    }
                                default:
                                    if ((v2 - int64(v3)) == (v2 * (v2 % ((v2) | 1)))) {
                                        fmt.Println(v1)
                                        var v45 uint = 629
                                        v46 := []int{fi38, fi41}
                                        _ = v45
                                        _ = v46
                                    } else {
                                        fmt.Println(v4, v1, v39)
                                        var v47 int32 = 2147483647
                                        _ = v47
                                    }
                                }
                                _ = v43
                            }
                        } else {
                            v48 := GenS6{F0: 42}
                            _ = v48
                        }
                        _ = v42
                    } else {
                        v39 = v39
                    }
                    if true {
                        v49 := GenS7{F0: 2147483648, F1: -54, F2: v3}
                        switch ("a" + v39) {
                        case "hello":
                            if ("a" < "foo bar") {
                                var v50 uint = 2
                                _ = v50
                            } else {
                                switch v39 {
                                case "\n\t":
                                    var v51 uint16 = (v3 / (((^(v3))) | 1))
                                    _ = v51
                                case "foo bar":
                                    v3 = v3
                                default:
                                    v52 := []int{}
                                    if len(v52) >= 2 {
                                        v52 = v52[1:len(v52)]
                                    }
                                    v4 -= v4
                                    _ = v52
                                }
                            }
                        case "foo bar":
                            v53 := GenS5{F0: (v4 >> ((2) & 63))}
                            _ = v53
                        default:
                            v54 := []int{(fi41 * fi38)}
                            _ = v54
                        }
                        _ = v49
                    }
                    _ = fi41
                }
                var v55 int = (fi38 ^ 32768)
                for fi56 := 0; fi56 < 20; fi56++ {
                    switch (v39 + (("hello" + "a") + v39)) {
                    case "a":
                        var v57 uint32 = 1
                        switch ((v39 + v39) + ("hello" + v39)) {
                        case "":
                            v39 = v39
                        default:
                            v58 := []int{}
                            if len(v58) > 0 {
                                v58[0] = fi56
                            }
                            _ = v58
                        }
                        _ = v57
                    default:
                        var v59 bool = false
                        _ = v59
                    }
                    v3 = 702
                    _ = fi56
                }
                _ = v55
            }
            _ = fi38
            _ = v39
        }
        v60 := GenS6{F0: 42}
        v61 := GenS8{F0: 42, F1: -704, F2: (((("hello" + "foo bar") + ("foo bar" + "")) + "\n\t") + "a")}
        _ = v60
        _ = v61
    }
    v62 := []int{4294967295, 128}
    for fi63 := 0; fi63 < 1; fi63++ {
        v62 = append(v62, fi63)
        if (4294967295 > 32768) {
            switch (^(fi63)) {
            case 32768:
                _ = len(v62)
            case 128:
                for fi64 := 0; fi64 < 9; fi64++ {
                    var v65 uint = 315
                    _ = fi64
                    _ = v65
                }
            }
            for fi66 := 0; fi66 < 2; fi66++ {
                v3 -= 834
                _ = fi66
            }
            if (("\n\t" <= "foo bar") && true) {
                v67 := GenS2{F0: v2, F1: 630}
                _ = v67
            }
        } else {
            v68 := GenS9{F0: 0, F1: v3}
            _ = v68
        }
        _ = fi63
    }
    if (1.5 >= -2.25) {
        var v69 int = 32768
        _ = v69
    } else {
        v62 = append(v62, 127)
    }
    for fi70 := 0; fi70 < 17; fi70++ {
        for fi71 := 0; fi71 < 6; fi71++ {
            if false {
                v1 += 2
                var v72 uint16 = (v3 % ((v3) | 1))
                fmt.Println(v3)
                _ = v72
            } else {
                if true {
                    v73 := GenS3{F0: (fi71 - fi71)}
                    if false {
                        var v74 float32 = float32(-2.25)
                        v2 = v2
                        _ = v74
                    }
                    var v75 uint64 = (v1 % ((v4) | 1))
                    switch "hello" {
                    case "hello":
                        fmt.Println(v75, v75)
                    default:
                        if (4294967295 < 4294967295) {
                            v3 = (^(v3))
                            var v76 uint = 1
                            _ = v76
                        } else {
                            fmt.Println(v75, fi70, fi71)
                            v62 = append(v62, (^(fi70)))
                            if len(v62) > 0 {
                                v62[0] = (^(fi71))
                            }
                        }
                        v1 = 606
                        var v77 uint32 = 2147483647
                        _ = v77
                    }
                    var v78 uint16 = v3
                    fmt.Println(fi70, v78)
                    _ = v73
                    _ = v75
                    _ = v78
                }
                var v79 int64 = (v2 / (((v2 >> ((1) & 63))) | 1))
                v3 = (v3 >> ((459) & 63))
                _ = v79
            }
            _ = fi71
        }
        _ = fi70
    }
    v3 -= 32767
    if (587 >= 32767) {
        for fi80 := 0; fi80 < 12; fi80++ {
            _ = len(v62)
            v81 := GenS10{F0: "\n\t"}
            _ = fi80
            _ = v81
        }
        if len(v62) > 0 {
            v62[0] = 35
        }
        fmt.Println(v4, v3, v1)
    } else {
        if ((2147483647 < 42) || (2147483647 == 751)) {
            for fi82 := 0; fi82 < 13; fi82++ {
                v1 = (v1 * (v4 - (v1 << ((257) & 63))))
                v2 = v2
                if (v1 == 18446744073709551615) {
                    var v83 string = "foo bar"
                    _ = v83
                }
                _ = fi82
            }
        }
        var v84 int = 430
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v62
}
