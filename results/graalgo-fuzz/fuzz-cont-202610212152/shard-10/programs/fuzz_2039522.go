package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int64; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 int; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint; F1 int32 }
func genFn0(p0 int) int32 {
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: v6, F1: 4294967295}
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    var v4 int16 = -32768
    var v5 uint = 1
    var v8 float64 = 0.1
    v9 := []int{-103}
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!true) {
            var v11 uint64 = 0
            var v12 int16 = (v4 / (((v4 | 32767)) | 1))
            for fi13 := 0; fi13 < 4; fi13++ {
                if (false || (!(!(0.1 < -2.25)))) {
                    v14 := GenS1{F0: -9223372036854775808, F1: 1}
                    _ = v14
                } else {
                    fmt.Println(fi10)
                    v3 = genFn0(fi10)
                    v4 = (v12 + (v4 | v12))
                    _ = len(v2)
                    switch (fi13 - fi10) {
                    case 0:
                        var v15 uint = v5
                        _ = v15
                    case 32767:
                        v16 := GenS2{F0: -128, F1: v3}
                        _ = v16
                    case 1:
                        var v17 float32 = float32(0.1)
                        _ = v17
                    }
                }
                v11 -= (v11 / ((v11) | 1))
                if false {
                    if (true && (genFn0(fi13) != v3)) {
                        var v18 uint64 = v11
                        _ = v18
                    } else {
                        switch "" {
                        case "\n\t":
                            v19 := GenS3{F0: fi13, F1: v12}
                            _ = v19
                        case "hello":
                            var v20 uint8 = 0
                            _ = v20
                        case "a":
                            v2[1] = "a"
                        }
                    }
                }
                fmt.Println(v4, v11)
                v21 := GenS4{F0: 68, F1: 2, F2: 2147483648}
                _ = fi13
                _ = v21
            }
            _ = v11
            _ = v12
        } else {
            if ((true || false) && (!(-9223372036854775808 != 0))) {
                v8 -= float64(3.14159)
                if true {
                    var v22 uint = v5
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 float32 = 0.0
                        fmt.Println(v24)
                        _ = fi23
                        _ = v24
                    }
                    if false {
                        if (532 < 1) {
                            if (4294967295 < 42) {
                                for fi25 := 0; fi25 < 18; fi25++ {
                                    switch -2147483648 {
                                    case 42:
                                        _ = len(v9)
                                    default:
                                        fmt.Println(fi25)
                                    }
                                    v2[42] = "a"
                                    v26 := GenS2{F0: 100, F1: -1}
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            for fi27 := 0; fi27 < 4; fi27++ {
                                var v28 int = (fi10 << ((v5) & 63))
                                _ = fi27
                                _ = v28
                            }
                        }
                        fmt.Println(fi10, v3, v4)
                    } else {
                        var v29 int = fi10
                        var v30 uint16 = 2
                        _ = v29
                        _ = v30
                    }
                    _ = v22
                }
            } else {
                fmt.Println(v5)
            }
            var v31 uint = (v5 ^ 2)
            v32 := GenS5{F0: 4294967295, F1: 1}
            _ = v31
            _ = v32
        }
        if (!true) {
            var v33 string = "hello"
            _ = v33
        }
        v3 = v3
        _ = fi10
    }
    if true {
        delete(v2, 705)
        v8 += v8
    } else {
        v8 += v8
        if (-2.25 != 0.0) {
            switch 4294967295 {
            case 32767:
                v3 -= genFn0(-2147483648)
            default:
                var v34 int = -197
                _ = v34
            }
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 bool = (!false)
                delete(v2, 448)
                _ = fi35
                _ = v36
            }
        }
        var v37 string = "hello"
        _ = v37
    }
    if false {
        for fi38 := 0; fi38 < 15; fi38++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            fmt.Println(v4, v8, fi38)
            _ = fi38
        }
        var v39 float32 = 0.1
        _ = v39
    } else {
        v40 := GenS4{F0: 1, F1: (v5 + (^(v5))), F2: 1}
        fmt.Println(v5, v8, v8)
        _ = v40
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        v5 = (v5 >> (((v5 / ((42) | 1))) & 63))
        v42 := GenS3{F0: fi41, F1: (v4 & (v4 & v4))}
        _ = fi41
        _ = v42
    }
    v4 += (v4 + (^(v4)))
    delete(v1, "")
    if (2147483648 >= 2147483648) {
        var v43 int = 9223372036854775807
        v44 := GenS6{F0: 87, F1: -389}
        var v45 bool = ((v4 << (((v5 - (v5 + (^(v5))))) & 63)) > (v4 & (v4 & 726)))
        _ = v43
        _ = v44
        _ = v45
    } else {
        fmt.Println(v3, v8, v3)
        fmt.Println(v8)
        fmt.Println(v4)
    }
    v5 = 4294967295
    for fi46 := 0; fi46 < 14; fi46++ {
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 string = "foo bar"
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    if true {
        v49 := GenS6{F0: -99, F1: 869}
        if true {
            fmt.Println(v3, v5, v5)
            var v50 float32 = 0.1
            switch 2147483647 {
            case -705:
                v51 := GenS3{F0: 127, F1: (v4 ^ v4)}
                _ = v51
            default:
                if false {
                    delete(v1, ("hello" + ""))
                    for fi52 := 0; fi52 < 9; fi52++ {
                        for fi53 := 0; fi53 < 17; fi53++ {
                            _ = len(v1)
                            _ = fi53
                        }
                        _ = fi52
                    }
                }
            }
            _ = v50
        }
        _ = v49
    } else {
        v5 -= v5
    }
    if true {
        for fi54 := 0; fi54 < 9; fi54++ {
            for fi55 := 0; fi55 < 6; fi55++ {
                delete(v2, 0)
                switch (-(fi54)) {
                case 32767:
                    var v56 int8 = 100
                    _ = v56
                case 127:
                    v4 = v4
                case 128:
                    v57 := GenS1{F0: -9223372036854775808, F1: 758}
                    _ = v57
                default:
                    if len(v9) > 0 {
                        v9[0] = (fi54 << (((v5 | v5)) & 63))
                    }
                }
                v5 += (v5 >> ((65535) & 63))
                var v58 bool = (127 <= -1)
                if v58 {
                    v8 = 0.1
                    if ("\n\t" > ("a" + "hello")) {
                        v59 := GenS2{F0: -3, F1: -2147483648}
                        v8 -= (v8 - v8)
                        v58 = (2 != 2)
                        switch fi55 {
                        case 722:
                            if ((127 <= 2) && (255 > 42)) {
                                if (!v58) {
                                    v58 = (915 > 42)
                                    var v60 string = (("a" + "hello") + ("a" + "foo bar"))
                                    v4 = v4
                                    _ = v60
                                }
                                var v61 int32 = 411
                                _ = v61
                            } else {
                                var v62 uint64 = 4294967295
                                var v63 uint64 = v62
                                fmt.Println(v62, v62)
                                _ = v62
                                _ = v63
                            }
                        default:
                            v8 -= 1.5
                        }
                        if v58 {
                            switch (fi54 - (-(fi54))) {
                            case 0:
                                v64 := GenS4{F0: -9223372036854775808, F1: v5, F2: 2147483648}
                                _ = v64
                            case -1:
                                v9 = append(v9, (fi55 | fi54))
                                v9 = append(v9, (-(fi55)))
                            default:
                                _ = len(v9)
                            }
                            v58 = false
                            var v65 uint16 = 516
                            v66 := GenS2{F0: -1, F1: (v3 << (((v5 * (v5 & 42))) & 63))}
                            _ = v65
                            _ = v66
                        }
                        _ = v59
                    } else {
                        var v67 int8 = int8(v3)
                        v67 += (v67 / (((v67 * v67)) | 1))
                        _ = v67
                    }
                    var v68 int64 = 11
                    _ = v68
                }
                _ = fi55
                _ = v58
            }
            var v69 uint32 = 1
            _ = fi54
            _ = v69
        }
        v4 = (v4 / ((v4) | 1))
        if (!true) {
            v70 := GenS6{F0: -1, F1: 1}
            fmt.Println(v5)
            var v71 string = "hello"
            _ = v70
            _ = v71
        } else {
            if true {
                for fi72 := 0; fi72 < 2; fi72++ {
                    var v73 bool = true
                    v5 += 2
                    v74 := GenS7{F0: fi72, F1: (^(v5)), F2: 0}
                    _ = fi72
                    _ = v73
                    _ = v74
                }
                fmt.Println(v4, v4, v3)
                v75 := v1[""]
                _ = v75
            }
        }
    }
    v8 -= v8
    if false {
        for fi76 := 0; fi76 < 4; fi76++ {
            v4 -= (v4 * v4)
            _ = fi76
        }
        if (832 != v5) {
            _ = len(v9)
            switch (("" + ("a" + ("a" + ("" + "hello")))) + "foo bar") {
            case "\n\t":
                for fi77 := 0; fi77 < 3; fi77++ {
                    for fi78 := 0; fi78 < 16; fi78++ {
                        _ = len(v9)
                        _ = fi78
                    }
                    for fi79 := 0; fi79 < 15; fi79++ {
                        if (!((true || (2147483647 <= 2147483647)) || ((v5 - (v5 - v5)) < v5))) {
                            v3 += genFn0((fi77 + (fi77 ^ 334)))
                        } else {
                            if (!false) {
                                var v80 uint64 = 4294967295
                                if false {
                                    switch (fi77 / ((2147483647) | 1)) {
                                    case 127:
                                        v3 = v3
                                    case -982:
                                        v81 := v2[288]
                                        _ = v81
                                    case 32767:
                                        if false {
                                            var v82 int = fi77
                                            var v83 uint64 = (v80 - (^(v80)))
                                            if (((v5 / ((v5) | 1)) <= v5) || (false || false)) {
                                                var v84 float32 = (float32((-2.25 / ((v8 - v8) / v8))) + ((float32(1.5) / 1e10) - float32((v8 * (v8 * ((v8 / 1.0) + v8))))))
                                                if true {
                                                    var v85 string = ""
                                                    v86 := GenS1{F0: 790, F1: 18446744073709551615}
                                                    _ = v85
                                                    _ = v86
                                                }
                                                delete(v1, "foo bar")
                                                _ = v84
                                            }
                                            v83 = 4294967295
                                            _ = v82
                                            _ = v83
                                        }
                                    default:
                                        v8 += v8
                                    }
                                    var v87 int8 = -92
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    v87 += (v87 >> ((v5) & 63))
                                    fmt.Println(fi77, v80, fi79)
                                    _ = v87
                                }
                                _ = len(v9)
                                _ = v80
                            }
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v88 int64 = 9223372036854775807
                            _ = v88
                        }
                        v4 += v4
                        if (!true) {
                            var v89 int16 = v4
                            v90 := GenS8{F0: 4294967295, F1: 2147483647}
                            v89 = 0
                            _ = v89
                            _ = v90
                        }
                        _ = fi79
                    }
                    _ = fi77
                }
            }
        }
        var v91 uint = (v5 % ((v5) | 1))
        _ = v91
    } else {
        v92 := GenS2{F0: -76, F1: v3}
        v3 -= genFn0(32768)
        _ = v92
    }
    v8 = (1.5 / float64(1e10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
