package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
type GenS4 struct { F0 string; F1 int; F2 string }

func main() {
    var v1 uint = 2
    var v2 int32 = 1
    var v3 float64 = 3.14159
    var v4 float32 = 1e10
    var v5 int16 = 1000
    var v6 string = "hello"
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint32 = 4294967295
        _ = fi7
        _ = v8
    }
    switch 32768 {
    case 42:
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint32 = 2147483648
            v11 := GenS1{F0: 2, F1: (v10 - v10)}
            _ = fi9
            _ = v10
            _ = v11
        }
    case 128:
        var v12 uint32 = 0
        _ = v12
    case -9223372036854775808:
        switch "foo bar" {
        case "hello":
            v5 = -453
        case "":
            v13 := []int{369, 9223372036854775807, 170}
            _ = v13
        default:
            for fi14 := 0; fi14 < 16; fi14++ {
                switch fi14 {
                case -9223372036854775808:
                    for fi15 := 0; fi15 < 19; fi15++ {
                        switch (v6 + (v6 + "\n\t")) {
                        case "a":
                            switch (-(fi14)) {
                            case -2147483648:
                                switch (v6 + v6) {
                                case "foo bar":
                                    if (fi15 != (fi15 + 42)) {
                                        v5 = v5
                                    } else {
                                        v16 := GenS1{F0: 1, F1: 2147483648}
                                        var v17 uint = (v1 % ((4294967295) | 1))
                                        _ = v16
                                        _ = v17
                                    }
                                case "hello":
                                    v18 := GenS1{F0: 497, F1: 4294967295}
                                    _ = v18
                                case "\n\t":
                                    var v19 int32 = v2
                                    _ = v19
                                }
                            case 7:
                                switch v6 {
                                case "foo bar":
                                    v5 += v5
                                case "\n\t":
                                    var v20 int64 = 231
                                    _ = v20
                                case "hello":
                                    if (1 <= 2147483647) {
                                        v3 = float64((1e10 / 1.0))
                                        v5 -= -485
                                    } else {
                                        switch 580 {
                                        case -61:
                                            if (96 >= 255) {
                                                v6 = v6
                                                v21 := []int{fi15, -1, (fi15 / (((fi15 & fi14)) | 1)), (fi15 * (fi14 % ((fi14) | 1)))}
                                                v22 := GenS1{F0: 4294967295, F1: 65535}
                                                switch (^(fi15)) {
                                                case 168:
                                                    var v23 string = v6
                                                    _ = v23
                                                }
                                                v24 := GenS1{F0: 0, F1: 577}
                                                _ = v21
                                                _ = v22
                                                _ = v24
                                            } else {
                                                var v25 uint32 = 2147483648
                                                var v26 int8 = -1
                                                _ = v25
                                                _ = v26
                                            }
                                        case 128:
                                            var v27 int16 = v5
                                            _ = v27
                                        }
                                        if false {
                                            v4 = 0.1
                                            v6 = v6
                                        } else {
                                            v3 = (((v3 * (3.14159 - (v3 / (v3 / (((1.5 + ((v3 / 1e10) / -2.25)) + v3) * (v3 - v3)))))) / float64(0.1)) + v3)
                                            v28 := []int{}
                                            fmt.Println(v3, fi15, v5)
                                            _ = v28
                                        }
                                        v3 += 0.1
                                    }
                                }
                            default:
                                fmt.Println(v2, v6, v6)
                            }
                        case "hello":
                            var v29 bool = (255 <= 225)
                            _ = v29
                        default:
                            v3 = 3.14159
                        }
                        var v30 float32 = v4
                        if (!true) {
                            if (v30 >= v4) {
                                if true {
                                    fmt.Println(v4, v5)
                                    fmt.Println(v3, v1)
                                }
                            }
                            if false {
                                switch ("\n\t" + (v6 + "\n\t")) {
                                case "a":
                                    v31 := []int{32768, (fi15 & (fi14 << (((v1 - (v1 & 18446744073709551615))) & 63)))}
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    v32 := GenS2{F0: "foo bar"}
                                    _ = v31
                                    _ = v32
                                case "hello":
                                    v1 = (v1 << ((4294967295) & 63))
                                }
                                var v33 int = fi14
                                var v34 float64 = -2.25
                                _ = v33
                                _ = v34
                            } else {
                                var v35 uint = v1
                                var v36 int32 = 42
                                _ = v35
                                _ = v36
                            }
                            v37 := []int{fi15, int(v5), (fi15 - int(v5))}
                            _ = v37
                        } else {
                            v3 -= v3
                            v38 := GenS2{F0: v6}
                            _ = v38
                        }
                        _ = fi15
                        _ = v30
                    }
                case 4294967295:
                    v39 := GenS1{F0: 42, F1: 181}
                    _ = v39
                case 381:
                    fmt.Println(v1, v5, v4)
                }
                _ = fi14
            }
        }
    default:
        var v40 int16 = (v5 - v5)
        _ = v40
    }
    switch (v6 + v6) {
    case "foo bar":
        switch v6 {
        case "\n\t":
            var v41 bool = (!false)
            _ = v41
        case "":
            var v42 string = "foo bar"
            _ = v42
        case "hello":
            for fi43 := 0; fi43 < 13; fi43++ {
                v44 := GenS3{F0: v5, F1: (v2 - (v2 - (v2 ^ (v2 / ((v2) | 1))))), F2: 114}
                _ = fi43
                _ = v44
            }
        default:
            if (!false) {
                v45 := []int{-805, 127, 503, 647}
                var v46 bool = false
                _ = v45
                _ = v46
            }
        }
    case "\n\t":
        for fi47 := 0; fi47 < 7; fi47++ {
            v4 = 1e10
            if (((v6 + (v6 + "")) < v6) || false) {
                for fi48 := 0; fi48 < 5; fi48++ {
                    v4 = 0.0
                    _ = fi48
                }
                fmt.Println(v6, v6)
            } else {
                for fi49 := 0; fi49 < 20; fi49++ {
                    var v50 int16 = v5
                    fmt.Println(v2, fi49)
                    if (v3 == ((v3 / float64(v4)) + 3.14159)) {
                        v51 := []int{}
                        v1 = v1
                        switch v6 {
                        case "":
                            _ = len(v51)
                        }
                        v50 = (v5 * v5)
                        _ = v51
                    } else {
                        fmt.Println(v6, v5, v3)
                        v52 := []int{int(v1), (^(fi49))}
                        var v53 string = v6
                        _ = v52
                        _ = v53
                    }
                    _ = fi49
                    _ = v50
                }
                var v54 int64 = -838
                _ = v54
            }
            _ = fi47
        }
    }
    if false {
        var v55 string = v6
        _ = v55
    } else {
        switch -9223372036854775808 {
        case 1:
            v2 = (v2 & 1)
        }
        var v56 int8 = 127
        fmt.Println(v4)
        _ = v56
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        v6 = "hello"
        if false {
            if false {
                v5 = 1
                var v58 int8 = 100
                _ = v58
            } else {
                var v59 string = ((("a" + "") + "foo bar") + ("foo bar" + v6))
                var v60 uint8 = 215
                var v61 int32 = v2
                _ = v59
                _ = v60
                _ = v61
            }
        }
        var v62 bool = (397 != 0)
        _ = fi57
        _ = v62
    }
    fmt.Println(v2, v6, v6)
    if true {
        for fi63 := 0; fi63 < 11; fi63++ {
            if (true || (true && (float32((v3 + v3)) > float32((float64(v4) / 1.0))))) {
                if false {
                    var v64 int16 = 32767
                    _ = v64
                } else {
                    if true {
                        v6 = v6
                    } else {
                        v65 := GenS3{F0: (v5 + -32768), F1: -2147483648, F2: 255}
                        var v66 float64 = (v3 - v3)
                        v3 += 0.1
                        _ = v65
                        _ = v66
                    }
                    var v67 string = ((v6 + "\n\t") + "foo bar")
                    _ = v67
                }
                fmt.Println(v6, v3)
            } else {
                var v68 uint = (v1 << (((v1 << (((v1 & (v1 << (((v1 | v1)) & 63)))) & 63))) & 63))
                _ = v68
            }
            v69 := []int{fi63, -116}
            _ = fi63
            _ = v69
        }
    }
    if (!true) {
        v2 = v2
        var v70 int32 = (v2 ^ 1)
        _ = v70
    }
    v2 += v2
    v5 = (v5 / ((v5) | 1))
    for fi71 := 0; fi71 < 4; fi71++ {
        if ((v3 * v3) < (0.1 - v3)) {
            v2 = (v2 | v2)
            v1 -= 177
        } else {
            if (!(true && false)) {
                var v72 uint16 = 255
                _ = v72
            } else {
                var v73 int = 9223372036854775807
                if (1e10 == v3) {
                    fmt.Println(v5, fi71, v2)
                    var v74 uint8 = 25
                    var v75 uint = (v1 + uint(v74))
                    _ = v74
                    _ = v75
                } else {
                    var v76 uint32 = 79
                    fmt.Println(v4)
                    v77 := []int{(fi71 / (((v73 + (fi71 >> (((v1 * v1)) & 63)))) | 1)), 127, fi71, -1}
                    _ = v76
                    _ = v77
                }
                v78 := GenS2{F0: v6}
                _ = v73
                _ = v78
            }
        }
        v5 -= (v5 << ((0) & 63))
        if (-85 != 2) {
            var v79 float32 = v4
            var v80 bool = (1 <= 41)
            _ = v79
            _ = v80
        } else {
            var v81 int64 = 9223372036854775807
            _ = v81
        }
        _ = fi71
    }
    v4 += v4
    for fi82 := 0; fi82 < 4; fi82++ {
        v83 := []int{fi82, (fi82 | fi82), (^(fi82))}
        _ = fi82
        _ = v83
    }
    switch v6 {
    case "a":
        switch ("hello" + "foo bar") {
        case "a":
            if ((v4 - v4) <= float32(v3)) {
                v84 := []int{32768, 42}
                _ = v84
            }
        case "foo bar":
            v85 := GenS4{F0: "a", F1: 32767, F2: v6}
            _ = v85
        }
    case "foo bar":
        var v86 bool = false
        _ = v86
    case "\n\t":
        v87 := []int{0, 314, -1}
        _ = v87
    }
    switch 32768 {
    case 127:
        v88 := []int{107, -9223372036854775808, 128}
        _ = v88
    }
    v2 -= 42
    v89 := []int{32767, 283}
    var v90 int16 = (v5 * v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v89:", len(v89), v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
}
