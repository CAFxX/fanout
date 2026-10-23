package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int; F2 uint32 }

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = -1
    var v3 uint = 4294967295
    var v4 uint64 = 4294967295
    v5 := make(map[int]string)
    if (v1 != 2147483647) {
        if true {
            v6 := GenS1{F0: 88, F1: 2147483648, F2: "hello"}
            _ = v6
        } else {
            if ((false || true) && ((^(v3)) <= 65535)) {
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 int32 = v1
                    delete(v5, (fi7 % ((fi7) | 1)))
                    switch "a" {
                    case "hello":
                        v2 -= v2
                    case "\n\t":
                        switch ("\n\t" + (("" + "hello") + (("foo bar" + "foo bar") + (("hello" + ("a" + "hello")) + ("a" + ""))))) {
                        case "hello":
                            v9 := GenS1{F0: 2, F1: 2, F2: "\n\t"}
                            _ = v9
                        default:
                            v8 += v1
                        }
                    case "foo bar":
                        if (1e10 != (0.1 / ((0.0 * (1e10 / float64(3.14159))) - -2.25))) {
                            if false {
                                for fi10 := 0; fi10 < 15; fi10++ {
                                    var v11 uint32 = 2147483647
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v12 string = ("a" + ((("a" + "foo bar") + ("foo bar" + (("\n\t" + "hello") + ("\n\t" + "foo bar")))) + "\n\t"))
                                        _ = v12
                                    case "":
                                        v8 -= 2147483647
                                    default:
                                        v13 := []int{fi10}
                                        switch (("foo bar" + "") + ("" + "hello")) {
                                        case "foo bar":
                                            if false {
                                                var v14 uint32 = 2
                                                if (float64(0.0) >= (0.0 - (0.0 / (0.1 - 3.14159)))) {
                                                    fmt.Println(fi7)
                                                    fmt.Println(v2, v4)
                                                    if (286 != v11) {
                                                        var v15 bool = true
                                                        if len(v13) > 0 {
                                                            v13[0] = (fi10 / ((884) | 1))
                                                        }
                                                        _ = v15
                                                    }
                                                }
                                                var v16 float32 = float32((-2.25 + -2.25))
                                                _ = v14
                                                _ = v16
                                            } else {
                                                switch 0 {
                                                case 1:
                                                    switch (fi7 >> (((v3 >> ((v3) & 63))) & 63)) {
                                                    case -9223372036854775808:
                                                        var v17 uint = v3
                                                        _ = v17
                                                    case 9223372036854775807:
                                                        v4 = (v4 + (v4 / (((v4 ^ v4)) | 1)))
                                                        v18 := GenS2{F0: 100}
                                                        v19 := GenS3{F0: false, F1: v2}
                                                        _ = v18
                                                        _ = v19
                                                    case 32768:
                                                        v20 := GenS2{F0: -1}
                                                        _ = v20
                                                    }
                                                case 4294967295:
                                                    _ = len(v5)
                                                case -506:
                                                    switch fi10 {
                                                    case 1:
                                                        fmt.Println(fi7, v11)
                                                    case 127:
                                                        if (v1 <= v8) {
                                                            v2 -= -578
                                                        }
                                                    default:
                                                        if ("foo bar" > ((("" + "hello") + "a") + "hello")) {
                                                            v21 := GenS4{F0: v11}
                                                            var v22 int = fi10
                                                            v13 = append(v13, 528)
                                                            _ = v21
                                                            _ = v22
                                                        } else {
                                                            fmt.Println(fi10, v8)
                                                        }
                                                    }
                                                    var v23 uint16 = 255
                                                    v24 := GenS3{F0: true, F1: v2}
                                                    _ = v23
                                                    _ = v24
                                                default:
                                                    var v25 int8 = -1
                                                    _ = v25
                                                }
                                                v11 = (v11 / (((v11 >> ((217) & 63))) | 1))
                                                fmt.Println(fi10, fi10)
                                                v26 := GenS5{F0: 127}
                                                _ = v26
                                            }
                                        case "hello":
                                            var v27 uint16 = 65535
                                            if true {
                                                var v28 string = ("hello" + "")
                                                _ = v28
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    var v29 float64 = 0.0
                                                    _ = v29
                                                case "a":
                                                    v30 := GenS6{F0: (v11 >> ((v3) & 63))}
                                                    var v31 bool = ("hello" == "a")
                                                    v5[(fi7 / ((fi10) | 1))] = ""
                                                    _ = v30
                                                    _ = v31
                                                case "hello":
                                                    v2 += v2
                                                }
                                                if true {
                                                    v11 -= v11
                                                    var v32 float32 = 0.0
                                                    _ = len(v13)
                                                    _ = v32
                                                }
                                                v8 = (v8 << (((^(v3))) & 63))
                                            }
                                            _ = v27
                                        case "\n\t":
                                            _ = len(v5)
                                        default:
                                            var v33 int16 = 983
                                            _ = v33
                                        }
                                        var v34 uint64 = v4
                                        var v35 uint = (^(v3))
                                        fmt.Println(v8)
                                        _ = v13
                                        _ = v34
                                        _ = v35
                                    }
                                    if true {
                                        var v36 int = (fi10 | fi7)
                                        v37 := GenS7{F0: (-(v1)), F1: -1, F2: v11}
                                        if true {
                                            fmt.Println(v3, v3, fi10)
                                        }
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = fi10
                                    _ = v11
                                }
                                v5[(fi7 & (-(fi7)))] = ("hello" + "foo bar")
                                switch (fi7 - (fi7 ^ (^(fi7)))) {
                                case -2147483648:
                                    var v38 float64 = 1.5
                                    _ = v38
                                case 1:
                                    switch ((("hello" + "") + ("\n\t" + "a")) + "hello") {
                                    case "a":
                                        if ((42 >= 114) || true) {
                                            v39 := GenS1{F0: 255, F1: 1, F2: "a"}
                                            for fi40 := 0; fi40 < 1; fi40++ {
                                                var v41 string = ("a" + ("\n\t" + ("foo bar" + "foo bar")))
                                                v8 = -2147483648
                                                v8 += v8
                                                _ = fi40
                                                _ = v41
                                            }
                                            _ = v39
                                        } else {
                                            v4 = 801
                                        }
                                    case "\n\t":
                                        v42 := GenS5{F0: int8(v4)}
                                        _ = v42
                                    case "hello":
                                        var v43 uint16 = 324
                                        _ = v43
                                    default:
                                        _ = len(v5)
                                    }
                                default:
                                    for fi44 := 0; fi44 < 15; fi44++ {
                                        fmt.Println(v4)
                                        v45 := []int{}
                                        _ = fi44
                                        _ = v45
                                    }
                                }
                            }
                            fmt.Println(fi7)
                        } else {
                            switch ("hello" + "foo bar") {
                            case "a":
                                var v46 int32 = (v8 | v8)
                                _ = v46
                            case "":
                                var v47 int = int(v2)
                                _ = v47
                            default:
                                if ((!false) && (false || true)) {
                                    v1 = (^(v1))
                                } else {
                                    for fi48 := 0; fi48 < 8; fi48++ {
                                        v49 := GenS2{F0: -62}
                                        var v50 int64 = 9223372036854775807
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            }
                            v5[(fi7 + (fi7 | fi7))] = "foo bar"
                        }
                    default:
                        fmt.Println(fi7, fi7, v2)
                    }
                    _ = fi7
                    _ = v8
                }
                v3 -= v3
            } else {
                v3 = v3
                v51 := GenS3{F0: ((v2 & 1000) != -32768), F1: 42}
                _ = v51
            }
            for fi52 := 0; fi52 < 1; fi52++ {
                if ((!(!true)) && true) {
                    v2 = v2
                }
                switch "hello" {
                case "hello":
                    v53 := []int{fi52, fi52}
                    _ = v53
                default:
                    var v54 int8 = -128
                    _ = v54
                }
                var v55 string = "foo bar"
                _ = fi52
                _ = v55
            }
        }
        var v56 string = ("\n\t" + "hello")
        _ = v56
    } else {
        for fi57 := 0; fi57 < 11; fi57++ {
            v58 := []int{fi57, (fi57 + (fi57 - fi57)), -463}
            var v59 uint64 = v4
            v5[(fi57 * (fi57 % ((42) | 1)))] = ("a" + (("hello" + ("foo bar" + "")) + "\n\t"))
            _ = fi57
            _ = v58
            _ = v59
        }
        for fi60 := 0; fi60 < 7; fi60++ {
            fmt.Println(v2)
            var v61 int16 = (v2 & 32767)
            _ = fi60
            _ = v61
        }
    }
    switch (("foo bar" + "foo bar") + "") {
    case "hello":
        v3 -= (v3 & uint(v4))
    case "foo bar":
        fmt.Println(v1)
    default:
        if true {
            fmt.Println(v3, v1, v2)
        }
    }
    if ((!false) && (193 != v4)) {
        v62 := []int{258}
        if (false && true) {
            var v63 uint = (v3 - 2)
            v3 -= (v63 << (((v3 * v3)) & 63))
            fmt.Println(v63, v2)
            _ = v63
        } else {
            var v64 int32 = v1
            switch ("foo bar" + ("foo bar" + "")) {
            case "foo bar":
                var v65 uint = v3
                _ = v65
            default:
                v62 = append(v62, 0)
            }
            _ = v64
        }
        v2 = v2
        _ = v62
    }
    fmt.Println(v4, v2)
    var v66 uint32 = 1
    var v67 int = -2147483648
    if true {
        v2 = (v2 / ((-1) | 1))
    } else {
        var v68 int = v67
        var v69 int64 = 9223372036854775807
        _ = v68
        _ = v69
    }
    v2 = 1
    for fi70 := 0; fi70 < 9; fi70++ {
        for fi71 := 0; fi71 < 15; fi71++ {
            v72 := []int{v67, (fi71 | 32768), -680}
            _ = fi71
            _ = v72
        }
        _ = fi70
    }
    var v73 string = "\n\t"
    if true {
        v5[694] = (v73 + "\n\t")
    }
    switch v73 {
    case "hello":
        delete(v5, v67)
    default:
        var v74 string = ((((v73 + ("" + (v73 + v73))) + v73) + v73) + v73)
        _ = v74
    }
    var v75 uint8 = 1
    var v76 int32 = -646
    if true {
        var v77 string = "hello"
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v73
    _ = v75
    _ = v76
}
