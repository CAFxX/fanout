package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 int64; F1 string; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 int = 32768
    var v2 string = "\n\t"
    var v3 uint = 10
    var v4 uint32 = 855
    v5 := make(map[int]string)
    v4 = v4
    v1 -= 128
    if false {
        v6 := GenS1{F0: v2, F1: v4, F2: 42}
        _ = v6
    }
    v7 := GenS1{F0: (v2 + ""), F1: 102, F2: -32768}
    if ((18446744073709551615 != 18446744073709551615) && (728 < 42)) {
        fmt.Println(v4, v2)
    } else {
        switch "foo bar" {
        case "hello":
            v3 = 4294967295
        case "foo bar":
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 uint64 = 481
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch fi8 {
                    case 983:
                        if true {
                            var v11 float32 = 3.14159
                            _ = v11
                        }
                    case 32768:
                        var v12 int = (fi8 & (fi10 | (v1 - (v1 - (fi8 & 9223372036854775807)))))
                        _ = v12
                    case 483:
                        if true {
                            _ = len(v5)
                            if true {
                                if (-2.25 <= 3.14159) {
                                    v3 -= 18446744073709551615
                                }
                            }
                        } else {
                            var v13 uint32 = 774
                            var v14 uint8 = 0
                            var v15 float64 = 1.5
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    default:
                        switch v2 {
                        case "a":
                            v16 := GenS1{F0: (((v2 + ((v2 + v2) + ("\n\t" + v2))) + "a") + v2), F1: v4, F2: 32767}
                            _ = v16
                        case "hello":
                            v17 := GenS2{F0: 42, F1: v2, F2: 0}
                            _ = v17
                        case "foo bar":
                            if true {
                                v1 -= (fi8 >> ((v3) & 63))
                            } else {
                                v18 := []int{v1}
                                v2 = "hello"
                                _ = v18
                            }
                        default:
                            v19 := GenS2{F0: 0, F1: v2, F2: 65535}
                            _ = v19
                        }
                    }
                    v2 = "foo bar"
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
        }
        v20 := GenS2{F0: -881, F1: v2, F2: 255}
        var v21 int32 = 2147483647
        _ = v20
        _ = v21
    }
    if (100 > 0) {
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := []int{}
            v1 -= v1
            _ = fi22
            _ = v23
        }
        v2 = ("hello" + v2)
    }
    v3 -= (v3 >> ((v3) & 63))
    switch v1 {
    case -1:
        fmt.Println(v3)
    case 32767:
        if ("" >= ((v2 + (v2 + ("" + (v2 + "\n\t")))) + "hello")) {
            if false {
                v24 := GenS3{F0: int64(v3), F1: 1, F2: uint(v4)}
                _ = v24
            } else {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 uint16 = 1
                    _ = fi25
                    _ = v26
                }
                fmt.Println(v4)
                var v27 int = v1
                _ = v27
            }
            v1 -= 2147483647
            delete(v5, (v1 & (v1 - (^(v1)))))
        }
    case -275:
        v4 = 0
    default:
        for fi28 := 0; fi28 < 7; fi28++ {
            fmt.Println(v1)
            v29 := GenS4{F0: 4294967295}
            _ = fi28
            _ = v29
        }
    }
    if (410 != 2147483647) {
        var v30 string = v2
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 uint32 = v4
            for fi33 := 0; fi33 < 11; fi33++ {
                var v34 int16 = 0
                _ = fi33
                _ = v34
            }
            fmt.Println(v30)
            _ = fi31
            _ = v32
        }
        _ = v30
    } else {
        v1 -= 1
        v2 = "foo bar"
    }
    fmt.Println(v3, v2, v1)
    v35 := GenS1{F0: v2, F1: v4, F2: 838}
    v36 := GenS2{F0: 604, F1: "foo bar", F2: 0}
    if (-2.25 > 1e10) {
        for fi37 := 0; fi37 < 2; fi37++ {
            v5[(v1 / ((32767) | 1))] = "\n\t"
            v2 = v2
            switch fi37 {
            case 407:
                for fi38 := 0; fi38 < 11; fi38++ {
                    if (42 == 100) {
                        v39 := v5[(v1 % (((fi38 - fi38)) | 1))]
                        switch v2 {
                        case "":
                            var v40 uint32 = v4
                            _ = v40
                        case "hello":
                            v41 := []int{(-(v1)), 9223372036854775807, fi37}
                            _ = v41
                        case "foo bar":
                            v42 := GenS4{F0: 42}
                            _ = v42
                        default:
                            v3 += v3
                        }
                        _ = v39
                    } else {
                        v43 := GenS1{F0: (((("foo bar" + v2) + ((v2 + v2) + (v2 + "foo bar"))) + v2) + ((v2 + ("a" + "foo bar")) + v2)), F1: 1, F2: -1}
                        var v44 uint32 = 2147483648
                        v2 = (("foo bar" + ((((v2 + (((v2 + "") + "") + v2)) + v2) + "a") + ((v2 + v2) + (v2 + ((v2 + "hello") + v2))))) + (v2 + "hello"))
                        _ = v43
                        _ = v44
                    }
                    v3 = (^(v3))
                    v1 += fi38
                    _ = fi38
                }
            case 127:
                var v45 uint8 = 30
                _ = v45
            }
            _ = fi37
        }
    } else {
        switch int(v4) {
        case 512:
            switch v2 {
            case "foo bar":
                var v46 uint64 = 18446744073709551615
                _ = v46
            case "hello":
                for fi47 := 0; fi47 < 18; fi47++ {
                    switch v2 {
                    case "\n\t":
                        var v48 uint64 = 4294967295
                        _ = v48
                    case "hello":
                        for fi49 := 0; fi49 < 7; fi49++ {
                            var v50 uint16 = 32767
                            if true {
                                v51 := GenS3{F0: -1, F1: v4, F2: (v3 & v3)}
                                if true {
                                    switch v2 {
                                    case "hello":
                                        if (128 > 179) {
                                            fmt.Println(v1, fi47)
                                            var v52 float64 = 3.14159
                                            _ = v52
                                        } else {
                                            delete(v5, (v1 / (((fi47 << ((v3) & 63))) | 1)))
                                            var v53 bool = true
                                            _ = v53
                                        }
                                        fmt.Println(fi49, fi47)
                                        if (127 > -33) {
                                            var v54 int16 = -527
                                            _ = v54
                                        } else {
                                            v55 := GenS2{F0: 42, F1: "", F2: (v50 % (((v50 >> (((^(v3))) & 63))) | 1))}
                                            _ = v55
                                        }
                                    case "foo bar":
                                        _ = len(v5)
                                    case "\n\t":
                                        var v56 string = v2
                                        var v57 int8 = int8(v50)
                                        if (true && (true && (1 == 9223372036854775807))) {
                                            var v58 bool = (42 >= 142)
                                            _ = v58
                                        } else {
                                            if (("a" + (v56 + v2)) > "\n\t") {
                                                var v59 string = "hello"
                                                _ = v59
                                            } else {
                                                v2 = "hello"
                                                switch v2 {
                                                case "hello":
                                                    if ((true && false) && true) {
                                                        var v60 float64 = 3.14159
                                                        if (1 <= 255) {
                                                            v56 = "\n\t"
                                                            v61 := GenS5{F0: false, F1: (v50 | (v50 ^ (v50 >> ((v3) & 63))))}
                                                            v62 := []int{(fi49 / (((fi47 - (v1 - 0))) | 1)), (fi49 - int(v50)), -416}
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        fmt.Println(v60, v2)
                                                        var v63 int32 = 2147483647
                                                        _ = v60
                                                        _ = v63
                                                    } else {
                                                        switch fi49 {
                                                        case 127:
                                                            if false {
                                                                v4 = (v4 - (v4 >> ((v3) & 63)))
                                                            }
                                                        case 32768:
                                                            var v64 int = v1
                                                            _ = v64
                                                        case 4294967295:
                                                            switch (fi49 / ((v1) | 1)) {
                                                            case -685:
                                                                var v65 float64 = (0.1 / (0.1 / 3.14159))
                                                                _ = v65
                                                            case -9223372036854775808:
                                                                v3 = 936
                                                            case 121:
                                                                v5[-9223372036854775808] = "foo bar"
                                                            default:
                                                                v66 := GenS6{F0: (^(v57)), F1: v4}
                                                                _ = v66
                                                            }
                                                        default:
                                                            fmt.Println(v3, v4)
                                                        }
                                                    }
                                                default:
                                                    v67 := GenS3{F0: -1, F1: v4, F2: (v3 - (v3 * 65535))}
                                                    _ = v67
                                                }
                                                fmt.Println(v50, v4)
                                            }
                                        }
                                        _ = v56
                                        _ = v57
                                    }
                                } else {
                                    v2 = (v2 + v2)
                                    v4 = (v4 << ((960) & 63))
                                }
                                switch (fi49 - 315) {
                                case 989:
                                    v4 -= uint32(v3)
                                default:
                                    v68 := []int{}
                                    v2 = ("foo bar" + v2)
                                    v5[v1] = v2
                                    _ = v68
                                }
                                _ = v51
                            } else {
                                _ = len(v5)
                                v69 := GenS7{F0: 716}
                                var v70 uint8 = 73
                                _ = v69
                                _ = v70
                            }
                            _ = fi49
                            _ = v50
                        }
                    case "":
                        v71 := v5[fi47]
                        _ = v71
                    }
                    _ = fi47
                }
            default:
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    switch ((v2 + v2) + v2) {
                    case "\n\t":
                        v72 := GenS6{F0: -1, F1: 1}
                        _ = v72
                    case "foo bar":
                        v73 := GenS7{F0: 473}
                        _ = v73
                    default:
                        v74 := []int{}
                        _ = v74
                    }
                }
            }
        case -9223372036854775808:
            var v75 int = (^(v1))
            _ = v75
        }
    }
    v76 := []int{(v1 / (((v1 + (-(v1)))) | 1)), (v1 << (((v3 - v3)) & 63))}
    v2 = v2
    for fi77 := 0; fi77 < 17; fi77++ {
        switch "foo bar" {
        case "\n\t":
            v2 = ("" + ((v2 + (v2 + v2)) + ""))
        case "":
            v4 -= 151
        }
        var v78 int32 = 126
        _ = fi77
        _ = v78
    }
    var v79 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v76:", len(v76), v76)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v7
    _ = v76
    _ = v79
}
