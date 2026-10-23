package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 int16; F1 uint16; F2 string }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 bool; F1 bool }
type GenS12 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint16, p1 uint32) int16 {
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 uint8 = 1
            var v7 string = ((((("hello" + "") + ("\n\t" + ("a" + "hello"))) + "") + ("\n\t" + "a")) + "foo bar")
            p1 = (^(p1))
            if false {
                p1 = 0
            } else {
                var v8 float64 = float64((((0.0 - -2.25) - 3.14159) / 1e10))
                if ((v6 >> ((1) & 63)) != (v6 - (v6 + v6))) {
                    fmt.Println(fi4, fi5, v6)
                    v9 := []int{}
                    p1 = 65535
                    _ = v9
                } else {
                    var v10 int8 = 100
                    v7 = "\n\t"
                    var v11 float64 = -2.25
                    _ = v10
                    _ = v11
                }
                _ = v8
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        v12 := []int{(^(fi4)), 65, -702, fi4}
        for fi13 := 0; fi13 < 1; fi13++ {
            switch 32768 {
            case 1:
                p0 = (^(p0))
            }
            v14 := GenS1{F0: fi4, F1: "\n\t", F2: 255}
            _ = fi13
            _ = v14
        }
        _ = fi4
        _ = v12
    }
    p0 += (p0 + p0)
    if (int32(p0) >= 599) {
        var v15 int8 = -128
        switch "foo bar" {
        case "\n\t":
            p0 = 1
        case "a":
            var v16 uint8 = 127
            _ = v16
        default:
            v15 -= (v15 + v15)
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            p0 -= (p0 / (((p0 >> ((1) & 63))) | 1))
            v18 := []int{}
            _ = fi17
            _ = v18
        }
        _ = v15
    } else {
        v19 := []int{-9223372036854775808, 42}
        switch "hello" {
        case "\n\t":
            switch 4294967295 {
            case 1:
                p0 = (p0 - p0)
            case 42:
                if (!false) {
                    v20 := GenS1{F0: -2147483648, F1: ("a" + "\n\t"), F2: 4294967295}
                    p1 = 4294967295
                    _ = v20
                }
            }
        case "":
            p1 = (p1 / ((1) | 1))
        default:
            switch ("" + "hello") {
            case "\n\t":
                v21 := GenS1{F0: -1, F1: "foo bar", F2: 255}
                _ = v21
            case "a":
                switch 32768 {
                case 0:
                    var v22 float32 = 0.1
                    _ = v22
                default:
                    p1 += 42
                }
            default:
                v23 := GenS1{F0: -9223372036854775808, F1: "\n\t", F2: 4294967295}
                _ = v23
            }
        }
        _ = v19
    }
    fmt.Println(p1)
    p0 += p0
    _ = p0
    _ = p1
    return 0
}
func genFn1() string {
    if (false || false) {
        for fi24 := 0; fi24 < 4; fi24++ {
            if ((41 > 54) || (1 != 48)) {
                v25 := GenS1{F0: (fi24 | (fi24 + (fi24 + (fi24 & fi24)))), F1: "", F2: 65535}
                v26 := GenS2{F0: 47}
                switch "a" {
                case "hello":
                    if false {
                        if false {
                            if false {
                                fmt.Println(fi24)
                                v27 := GenS3{F0: 2, F1: 283}
                                _ = v27
                            } else {
                                v28 := []int{(-(fi24)), (fi24 << ((4294967295) & 63)), fi24, fi24}
                                switch ("foo bar" + "foo bar") {
                                case "hello":
                                    _ = len(v28)
                                case "":
                                    var v29 string = "foo bar"
                                    _ = v29
                                }
                                for fi30 := 0; fi30 < 11; fi30++ {
                                    var v31 float64 = (0.1 - -2.25)
                                    _ = fi30
                                    _ = v31
                                }
                                _ = v28
                            }
                            var v32 uint64 = 121
                            v33 := GenS4{F0: -1, F1: 4294967295, F2: 198}
                            _ = v32
                            _ = v33
                        } else {
                            var v34 float32 = (float32(-2.25) / 0.1)
                            var v35 float64 = (0.1 / -2.25)
                            _ = v34
                            _ = v35
                        }
                        if false {
                            var v36 float32 = -2.25
                            v36 = (0.1 / float32(0.1))
                            for fi37 := 0; fi37 < 9; fi37++ {
                                v36 -= (float32(3.14159) + v36)
                                v36 = 1e10
                                _ = fi37
                            }
                            _ = v36
                        }
                        switch 128 {
                        case 128:
                            for fi38 := 0; fi38 < 1; fi38++ {
                                var v39 string = "foo bar"
                                if (!(2147483647 == 345)) {
                                    v40 := []int{fi38}
                                    _ = v40
                                } else {
                                    var v41 uint = 591
                                    v42 := GenS5{F0: 100, F1: 197}
                                    v39 = v39
                                    _ = v41
                                    _ = v42
                                }
                                switch (^(fi38)) {
                                case 0:
                                    var v43 int32 = -2147483648
                                    _ = v43
                                case 32767:
                                    var v44 int64 = -1
                                    _ = v44
                                default:
                                    switch "\n\t" {
                                    case "hello":
                                        v45 := []int{(fi38 | (^(fi24))), fi24, fi24}
                                        _ = v45
                                    case "a":
                                        switch "a" {
                                        case "":
                                            v39 = ((v39 + "") + v39)
                                        case "foo bar":
                                            var v46 float64 = -2.25
                                            switch "foo bar" {
                                            case "\n\t":
                                                v39 = "a"
                                                if false {
                                                    v39 = v39
                                                    fmt.Println(fi24, fi38)
                                                } else {
                                                    var v47 float32 = 0.0
                                                    var v48 string = v39
                                                    v49 := []int{fi24}
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                            default:
                                                switch (v39 + v39) {
                                                case "":
                                                    var v50 int32 = 920
                                                    _ = v50
                                                case "a":
                                                    fmt.Println(v46)
                                                    v51 := []int{fi38, (-(fi24)), (^(fi38)), fi38}
                                                    _ = v51
                                                case "foo bar":
                                                    v52 := GenS3{F0: 141, F1: 2}
                                                    _ = v52
                                                default:
                                                    v46 += ((v46 + (v46 - v46)) - (1.5 - 1.5))
                                                }
                                            }
                                            _ = v46
                                        default:
                                            var v53 int = -9223372036854775808
                                            _ = v53
                                        }
                                    }
                                }
                                _ = fi38
                                _ = v39
                            }
                        case 32767:
                            switch (fi24 / ((fi24) | 1)) {
                            case -9223372036854775808:
                                var v54 uint32 = 65535
                                _ = v54
                            case -946:
                                v55 := GenS1{F0: (fi24 | (fi24 << ((255) & 63))), F1: "a", F2: 65535}
                                _ = v55
                            }
                        }
                    }
                }
                _ = v25
                _ = v26
            } else {
                var v56 bool = (255 < 4294967295)
                v56 = v56
                for fi57 := 0; fi57 < 1; fi57++ {
                    v56 = (true || v56)
                    fmt.Println(v56)
                    _ = fi57
                }
                _ = v56
            }
            fmt.Println(fi24)
            _ = fi24
        }
        var v58 uint32 = 217
        var v59 bool = (127 != 42)
        _ = v58
        _ = v59
    } else {
        v60 := []int{-385, 128, 0}
        _ = v60
    }
    if (42 < 1) {
        switch "" {
        case "\n\t":
            v61 := []int{128, 127, -663}
            _ = v61
        default:
            if false {
                if false {
                    var v62 int8 = 1
                    if true {
                        for fi63 := 0; fi63 < 20; fi63++ {
                            v64 := GenS6{F0: 1, F1: (fi63 - 32768)}
                            v65 := GenS7{F0: "hello"}
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    } else {
                        var v66 string = ("hello" + (("a" + "foo bar") + "foo bar"))
                        v66 = v66
                        _ = v66
                    }
                    _ = v62
                } else {
                    var v67 uint16 = 26
                    if (float64(1e10) < 3.14159) {
                        v67 -= 255
                    } else {
                        v68 := GenS8{F0: 42, F1: 2147483647}
                        for fi69 := 0; fi69 < 6; fi69++ {
                            v67 = 32767
                            v70 := []int{}
                            v67 += v67
                            _ = fi69
                            _ = v70
                        }
                        _ = v68
                    }
                    _ = v67
                }
                v71 := []int{32768, -1, 4294967295}
                _ = v71
            }
        }
    } else {
        var v72 float32 = float32(1.5)
        v72 += float32(1.5)
        v72 -= 1e10
        _ = v72
    }
    var v73 uint8 = 48
    v74 := GenS9{F0: 1000, F1: 42, F2: (("hello" + "hello") + "")}
    v73 += v73
    _ = v73
    _ = v74
    return "foo bar"
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 33
    var v75 uint8 = 128
    var v76 bool = true
    switch 819 {
    case -669:
        v75 -= (^(v75))
    case 9223372036854775807:
        var v77 int8 = 18
        _ = v77
    }
    var v78 float32 = 1.5
    var v79 float32 = float32(3.14159)
    if true {
        if (!v76) {
            var v80 int64 = -838
            v81 := []int{34, 32767, 1}
            _ = v80
            _ = v81
        }
        v82 := GenS10{F0: 32767, F1: 32767}
        var v83 bool = ((425 != 1) && false)
        _ = v82
        _ = v83
    }
    v84 := GenS5{F0: -1, F1: (^(v3))}
    v85 := GenS11{F0: v76, F1: v76}
    for fi86 := 0; fi86 < 3; fi86++ {
        var v87 uint = 18446744073709551615
        _ = fi86
        _ = v87
    }
    var v88 int32 = 0
    v79 = v79
    for fi89 := 0; fi89 < 6; fi89++ {
        v90 := []int{595}
        v91 := GenS12{F0: 255, F1: (v3 / ((42) | 1))}
        _ = fi89
        _ = v90
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v75
    _ = v76
    _ = v78
    _ = v79
    _ = v84
    _ = v85
    _ = v88
}
