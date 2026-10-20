package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    var v1 int64 = 342
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 int8 = 0
    var v5 string = "hello"
    var v6 int8 = 1
    v7 := []int{2147483647, -9223372036854775808, -2147483648, 9223372036854775807}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            if ((0.0 / 1.0) != -2.25) {
                var v10 int32 = 195
                var v11 float64 = 3.14159
                fmt.Println(v11)
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: (v6 - v6), F1: 127}
            v13 := GenS1{F0: (v4 ^ (v4 >> ((42) & 63))), F1: 127}
            _ = fi9
            _ = v12
            _ = v13
        }
        fmt.Println(v6, v4)
        _ = fi8
    }
    v4 -= (v6 | 42)
    var v14 bool = (false && false)
    var v15 string = "\n\t"
    var v16 float64 = (0.1 + (1.5 + 0.0))
    var v17 uint16 = 65535
    var v18 uint16 = (^(v17))
    switch (v15 + v15) {
    case "":
        switch 0 {
        case 127:
            v5 = ("hello" + "hello")
        default:
            fmt.Println(v17, v18)
        }
    case "foo bar":
        var v19 uint = 2
        _ = v19
    }
    var v20 int8 = -70
    var v21 uint = uint(v18)
    _ = len(v7)
    for fi22 := 0; fi22 < 19; fi22++ {
        if true {
            for fi23 := 0; fi23 < 13; fi23++ {
                v17 += v17
                switch (v15 + v5) {
                case "a":
                    v1 -= v1
                case "\n\t":
                    if len(v7) > 0 {
                        v7[0] = fi23
                    }
                }
                _ = fi23
            }
        } else {
            v24 := GenS1{F0: v4, F1: 183}
            v14 = false
            _ = v24
        }
        if false {
            if ((!(v6 == v6)) && (v20 <= (v4 - int8(v21)))) {
                v25 := v2["hello"]
                if v14 {
                    var v26 float64 = v16
                    var v27 uint8 = uint8(v18)
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(fi22)
                    if (-717 > fi22) {
                        if v14 {
                            _ = len(v7)
                        } else {
                            if v14 {
                                _ = len(v7)
                                switch (fi22 ^ fi22) {
                                case 1:
                                    fmt.Println(v4, v17, fi22)
                                case 4294967295:
                                    switch (v25 + v15) {
                                    case "\n\t":
                                        v1 -= (^(v1))
                                    }
                                }
                                _ = len(v7)
                            }
                        }
                        var v28 string = v5
                        var v29 uint = (v21 - v21)
                        _ = v28
                        _ = v29
                    }
                }
                _ = v25
            }
            fmt.Println(v17)
            switch fi22 {
            case 527:
                switch (^(fi22)) {
                case -378:
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v14 = false
                        if (-1 <= -817) {
                            switch fi30 {
                            case 32768:
                                v6 += 98
                                switch v15 {
                                case "foo bar":
                                    switch (fi22 | (fi30 | (fi22 << ((v21) & 63)))) {
                                    case 4294967295:
                                        v17 += v17
                                    case 128:
                                        var v31 int64 = (v1 & (v1 * v1))
                                        _ = v31
                                    }
                                case "a":
                                    v32 := GenS1{F0: v20, F1: 20}
                                    _ = v32
                                case "hello":
                                    v18 = v18
                                default:
                                    fmt.Println(v18, v5, v15)
                                }
                                if v14 {
                                    if false {
                                        v5 = v5
                                        var v33 int64 = 0
                                        v21 = v21
                                        _ = v33
                                    }
                                    fmt.Println(v4, v1, v14)
                                } else {
                                    v34 := GenS2{F0: 828, F1: 0}
                                    _ = v34
                                }
                                v35 := GenS1{F0: (v4 & v6), F1: 131}
                                _ = v35
                            }
                            v6 = int8(v17)
                        } else {
                            v1 -= (v1 & (^(v1)))
                        }
                        switch "\n\t" {
                        case "a":
                            switch (fi30 | fi22) {
                            case 32768:
                                v6 = (^(v20))
                            case 128:
                                v6 -= -83
                                var v36 uint64 = 42
                                v37 := GenS2{F0: (v36 / (((v36 + (v36 ^ (v36 * v36)))) | 1)), F1: -918}
                                _ = v36
                                _ = v37
                            case 0:
                                v38 := v3[-2147483648]
                                _ = v38
                            default:
                                if (((!(v16 >= v16)) || (113 != 192)) || false) {
                                    if v14 {
                                        fmt.Println(v6)
                                    } else {
                                        _ = len(v2)
                                        if (4294967295 <= 4294967295) {
                                            var v39 string = ""
                                            if v14 {
                                                var v40 int32 = int32(v21)
                                                v41 := GenS1{F0: 42, F1: 127}
                                                switch "hello" {
                                                case "\n\t":
                                                    var v42 uint = (^(v21))
                                                    _ = v42
                                                case "a":
                                                    v43 := GenS3{F0: -32768, F1: -1}
                                                    _ = v43
                                                case "hello":
                                                    var v44 int64 = -1
                                                    _ = v44
                                                default:
                                                    _ = len(v7)
                                                }
                                                _ = v40
                                                _ = v41
                                            } else {
                                                if ((!false) || v14) {
                                                    v6 = (v6 | (v20 - (-(v20))))
                                                } else {
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    var v45 string = "\n\t"
                                                    _ = v45
                                                }
                                                var v46 bool = v14
                                                var v47 uint = 65535
                                                _ = v46
                                                _ = v47
                                            }
                                            var v48 int16 = int16(v18)
                                            if len(v7) > 0 {
                                                v7[0] = 128
                                            }
                                            _ = v39
                                            _ = v48
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = (fi22 ^ fi30)
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = -9223372036854775808
                                        }
                                        v14 = true
                                    }
                                    v49 := GenS3{F0: -435, F1: v1}
                                    var v50 uint32 = 815
                                    _ = v49
                                    _ = v50
                                } else {
                                    var v51 uint64 = 231
                                    _ = v51
                                }
                            }
                        }
                        _ = fi30
                    }
                case -9223372036854775808:
                    _ = len(v3)
                }
            case 128:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case -1:
                for fi52 := 0; fi52 < 7; fi52++ {
                    if v14 {
                        v1 = 237
                        v14 = false
                    } else {
                        var v53 uint16 = 293
                        v21 += (v21 / ((65535) | 1))
                        _ = v53
                    }
                    _ = fi52
                }
            default:
                if (42 == -286) {
                    v1 = v1
                    var v54 float64 = v16
                    _ = v54
                } else {
                    var v55 uint = (v21 * v21)
                    v4 += v6
                    var v56 uint8 = 115
                    _ = v55
                    _ = v56
                }
            }
        }
        _ = fi22
    }
    if false {
        switch (("a" + ("a" + v15)) + (v15 + v15)) {
        case "hello":
            switch 0 {
            case 0:
                switch ((v5 + (v5 + v15)) + (v15 + v15)) {
                case "":
                    v14 = v14
                default:
                    var v57 uint64 = 1
                    _ = v57
                }
            case 32767:
                if false {
                    v58 := GenS3{F0: int16(v6), F1: (v1 / (((v1 & v1)) | 1))}
                    _ = v58
                }
            }
        case "":
            v4 -= 127
        }
    } else {
        v20 += (v6 + v20)
    }
    v21 += 480
    switch "hello" {
    case "foo bar":
        if v14 {
            fmt.Println(v6, v5, v21)
            for fi59 := 0; fi59 < 15; fi59++ {
                var v60 uint32 = 215
                fmt.Println(v14, v4, v17)
                _ = fi59
                _ = v60
            }
            switch (((v15 + (v5 + "\n\t")) + (("a" + (("a" + v15) + v15)) + "a")) + ("foo bar" + v5)) {
            case "foo bar":
                var v61 uint8 = 0
                _ = v61
            default:
                var v62 uint8 = 128
                _ = v62
            }
        } else {
            if (v14 && (118 <= 1000)) {
                if v14 {
                    switch "a" {
                    case "":
                        fmt.Println(v1)
                    case "foo bar":
                        var v63 string = "hello"
                        _ = v63
                    }
                    v64 := GenS1{F0: (v20 + v4), F1: 0}
                    var v65 string = ((v15 + (v5 + v15)) + "\n\t")
                    _ = v64
                    _ = v65
                }
                for fi66 := 0; fi66 < 1; fi66++ {
                    if (255 > v18) {
                        fmt.Println(v15, v5, v21)
                    }
                    var v67 int32 = -2147483648
                    var v68 int16 = 0
                    _ = fi66
                    _ = v67
                    _ = v68
                }
                var v69 int16 = 1
                _ = v69
            }
            var v70 string = "foo bar"
            var v71 uint = (^(v21))
            _ = v70
            _ = v71
        }
    case "hello":
        v18 += 42
    case "\n\t":
        var v72 uint32 = 4294967295
        _ = v72
    default:
        v73 := GenS1{F0: v4, F1: 0}
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
