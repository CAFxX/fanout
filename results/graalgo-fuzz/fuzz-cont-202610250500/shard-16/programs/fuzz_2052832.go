package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    var v1 float64 = 0.1
    var v2 int32 = -2147483648
    var v3 int32 = 0
    var v4 uint = 621
    var v5 uint = (^(v4))
    var v6 string = ("" + "a")
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            v1 = v1
        }
        v5 -= (v4 >> ((v5) & 63))
        v8 := GenS1{F0: 331, F1: 1, F2: 8}
        _ = fi7
        _ = v8
    }
    var v9 int = 1
    switch ("" + v6) {
    case "a":
        v2 = (v3 * (v2 + (v2 << ((v4) & 63))))
    case "hello":
        v10 := GenS2{F0: (^(v9)), F1: 2147483647}
        _ = v10
    case "foo bar":
        switch (v9 / (((v9 * (v9 + (-(v9))))) | 1)) {
        case 4294967295:
            if true {
                v5 = 4294967295
                v11 := GenS2{F0: (v9 - -1), F1: 1}
                _ = v11
            } else {
                v1 = v1
                for fi12 := 0; fi12 < 17; fi12++ {
                    if ((128 != (^(v9))) && false) {
                        var v13 bool = (42 > 86)
                        switch v6 {
                        case "foo bar":
                            var v14 uint64 = 349
                            _ = v14
                        case "hello":
                            for fi15 := 0; fi15 < 7; fi15++ {
                                var v16 bool = (v2 == v2)
                                var v17 uint64 = 496
                                v6 = ("\n\t" + "foo bar")
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                        if v13 {
                            v18 := GenS3{F0: -1, F1: v5, F2: "foo bar"}
                            for fi19 := 0; fi19 < 5; fi19++ {
                                if true {
                                    var v20 uint64 = 0
                                    if (!true) {
                                        var v21 bool = v13
                                        _ = v21
                                    }
                                    _ = v20
                                } else {
                                    var v22 int32 = v3
                                    var v23 bool = false
                                    _ = v22
                                    _ = v23
                                }
                                if (!(v2 <= -389)) {
                                    if v13 {
                                        if v13 {
                                            fmt.Println(v1, v9)
                                            v24 := []int{}
                                            v4 = v5
                                            v25 := GenS4{F0: -1}
                                            var v26 int64 = 0
                                            switch v6 {
                                            case "a":
                                                v2 = (v3 << ((v4) & 63))
                                            case "":
                                                if v13 {
                                                    if false {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v27 int = v9
                                                            _ = v27
                                                        case "a":
                                                            if len(v24) >= 2 {
                                                                v24 = v24[1:len(v24)]
                                                            }
                                                        case "hello":
                                                            fmt.Println(v6, v1, fi19)
                                                        }
                                                        v26 = v26
                                                        v28 := GenS5{F0: 1}
                                                        if v13 {
                                                            v6 = (v6 + ("" + (v6 + ("" + "hello"))))
                                                            switch fi19 {
                                                            case -1:
                                                                v29 := GenS4{F0: -1}
                                                                _ = v29
                                                            case 32767:
                                                                if len(v24) >= 2 {
                                                                    v24 = v24[1:len(v24)]
                                                                }
                                                            case 42:
                                                                var v30 int8 = 127
                                                                _ = v30
                                                            }
                                                            var v31 uint8 = uint8(v5)
                                                            _ = v31
                                                        }
                                                        v26 -= (v26 / ((0) | 1))
                                                        _ = v28
                                                    }
                                                    v32 := GenS4{F0: -607}
                                                    _ = v32
                                                } else {
                                                    if (2 <= 1) {
                                                        var v33 int16 = -32768
                                                        _ = v33
                                                    }
                                                }
                                            default:
                                                var v34 int8 = -128
                                                v24 = append(v24, fi19)
                                                _ = v34
                                            }
                                            var v35 float32 = 0.1
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                            _ = v35
                                        }
                                        if true {
                                            v13 = (127 >= 62)
                                            var v36 int8 = -54
                                            fmt.Println(fi12)
                                            if true {
                                                v37 := []int{}
                                                v2 = (v2 << ((v5) & 63))
                                                _ = v37
                                            } else {
                                                switch v6 {
                                                case "hello":
                                                    if (152 <= 29) {
                                                        v2 += v2
                                                        var v38 uint64 = 2
                                                        if (479 == 306) {
                                                            fmt.Println(v6, v36, fi12)
                                                        } else {
                                                            v4 -= 592
                                                        }
                                                        var v39 int32 = v3
                                                        v40 := GenS5{F0: 1}
                                                        _ = v38
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        var v41 uint64 = 1
                                                        v42 := []int{(fi19 ^ fi12)}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                case "a":
                                                    v43 := GenS6{F0: 0}
                                                    _ = v43
                                                default:
                                                    var v44 int = (fi19 % (((fi12 ^ (fi19 - (-(v9))))) | 1))
                                                    _ = v44
                                                }
                                            }
                                            v45 := GenS6{F0: -1}
                                            _ = v36
                                            _ = v45
                                        }
                                        v1 = (-2.25 * ((v1 + 0.0) / float64(float32(v1))))
                                    } else {
                                        v46 := GenS7{F0: v9}
                                        if true {
                                            fmt.Println(fi19, v9)
                                            var v47 int32 = 1
                                            fmt.Println(v9, v4)
                                            if v13 {
                                                if v13 {
                                                    fmt.Println(v6, v47)
                                                    if ("hello" >= (v6 + v6)) {
                                                        v48 := GenS8{F0: 105}
                                                        _ = v48
                                                    } else {
                                                        fmt.Println(v2, v5, v4)
                                                        if false {
                                                            var v49 int8 = 59
                                                            var v50 int64 = 1
                                                            var v51 uint16 = 32767
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        } else {
                                                            v52 := []int{(fi19 + (fi12 & (v9 + 721))), (fi12 - (fi12 % ((v9) | 1))), v9, (fi12 << (((v5 % (((v5 << ((2) & 63))) | 1))) & 63))}
                                                            var v53 int64 = -923
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        fmt.Println(v47, v5)
                                                    }
                                                } else {
                                                    var v54 float64 = (0.1 / v1)
                                                    if (v13 || ((v6 + ("foo bar" + v6)) == v6)) {
                                                        v55 := []int{127, 128, (fi12 % (((v9 - fi12)) | 1))}
                                                        _ = v55
                                                    }
                                                    _ = v54
                                                }
                                                v47 -= (v3 - v47)
                                            }
                                            v56 := []int{}
                                            _ = v47
                                            _ = v56
                                        }
                                        var v57 uint8 = 39
                                        _ = v46
                                        _ = v57
                                    }
                                    fmt.Println(v2, v9)
                                } else {
                                    var v58 int16 = -1
                                    v59 := []int{177, 42, 828}
                                    _ = v58
                                    _ = v59
                                }
                                _ = fi19
                            }
                            _ = v18
                        } else {
                            for fi60 := 0; fi60 < 13; fi60++ {
                                fmt.Println(v13)
                                _ = fi60
                            }
                            for fi61 := 0; fi61 < 13; fi61++ {
                                var v62 uint = (^(v5))
                                var v63 uint = 1
                                if (!(!(1 <= 132))) {
                                    switch (v6 + v6) {
                                    case "foo bar":
                                        v64 := []int{v9, v9, (fi61 >> ((v62) & 63)), (^(v9))}
                                        _ = v64
                                    case "":
                                        if false {
                                            var v65 uint32 = 731
                                            _ = v65
                                        }
                                    }
                                    v5 += (^(v5))
                                    v66 := []int{(^(fi61)), (fi61 | 42), (-(fi12))}
                                    v66 = append(v66, v9)
                                    _ = v66
                                }
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                            v1 = float64(0.1)
                        }
                        _ = v13
                    }
                    _ = fi12
                }
            }
        case 127:
            var v67 float32 = (1e10 * 1e10)
            _ = v67
        }
    }
    var v68 uint32 = 0
    v69 := GenS1{F0: 255, F1: -922, F2: 2147483648}
    for fi70 := 0; fi70 < 1; fi70++ {
        for fi71 := 0; fi71 < 10; fi71++ {
            var v72 bool = (true && true)
            fmt.Println(v5)
            if v72 {
                var v73 float32 = 0.0
                _ = v73
            } else {
                v72 = (-32768 < -32768)
            }
            _ = fi71
            _ = v72
        }
        for fi74 := 0; fi74 < 20; fi74++ {
            fmt.Println(v5)
            var v75 int8 = -1
            v76 := GenS9{F0: 255, F1: 2}
            _ = fi74
            _ = v75
            _ = v76
        }
        _ = fi70
    }
    v77 := []int{(v9 - (v9 * (v9 >> ((v4) & 63))))}
    var v78 int = (v9 >> ((255) & 63))
    if (255 != (^(v4))) {
        var v79 string = ""
        switch "\n\t" {
        case "\n\t":
            v79 = v6
        default:
            fmt.Println(v3)
        }
        v1 += v1
        _ = v79
    }
    v80 := GenS10{F0: 1, F1: 128, F2: 18446744073709551615}
    v4 = v4
    fmt.Println(v78, v5, v78)
    for fi81 := 0; fi81 < 2; fi81++ {
        for fi82 := 0; fi82 < 7; fi82++ {
            fmt.Println(v4, v2, v68)
            _ = fi82
        }
        fmt.Println(v6, fi81, v4)
        _ = fi81
    }
    var v83 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v77
    _ = v78
    _ = v80
    _ = v83
    _ = v9
}
