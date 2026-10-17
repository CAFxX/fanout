package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 string; F1 bool; F2 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 bool; F1 int }

func main() {
    var v1 int8 = 1
    var v2 uint16 = 585
    var v3 int8 = 88
    var v4 int8 = v1
    v5 := make(map[string]bool)
    if true {
        v5[("\n\t" + "hello")] = (42 < 4294967295)
        v6 := GenS1{F0: 54}
        _ = v6
    }
    switch "hello" {
    case "":
        var v7 int32 = -128
        _ = v7
    }
    v8 := GenS2{F0: -790}
    v3 = (v3 * v3)
    v4 = (v4 % ((v3) | 1))
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v4 = v4
            v11 := []int{(-(fi10)), fi9, (fi9 & (fi10 - fi9)), -735}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v1 += (v4 % (((v1 | (-(v4)))) | 1))
    if (0 == 4294967295) {
        if false {
            for fi12 := 0; fi12 < 20; fi12++ {
                v4 = (v4 & (v3 >> ((0) & 63)))
                _ = fi12
            }
        } else {
            v13 := GenS2{F0: 2147483647}
            fmt.Println(v1)
            v14 := []int{-391, -1, 32768}
            _ = v13
            _ = v14
        }
        fmt.Println(v1, v4)
        if false {
            var v15 uint32 = 2147483647
            if false {
                v3 = v4
                for fi16 := 0; fi16 < 8; fi16++ {
                    switch (^(fi16)) {
                    case 2147483647:
                        v17 := GenS3{F0: -1, F1: -1, F2: 42}
                        _ = v17
                    case -131:
                        v18 := GenS1{F0: 2}
                        _ = v18
                    }
                    _ = fi16
                }
                switch ("foo bar" + "\n\t") {
                case "a":
                    fmt.Println(v4, v1, v2)
                case "foo bar":
                    if false {
                        var v19 int32 = int32(v15)
                        delete(v5, (("\n\t" + "hello") + "foo bar"))
                        _ = v19
                    } else {
                        v20 := GenS4{F0: 42, F1: 1}
                        _ = v20
                    }
                default:
                    if (-1 > 42) {
                        if true {
                            var v21 uint64 = 4294967295
                            _ = v21
                        } else {
                            v22 := GenS2{F0: -1}
                            v15 = (v15 % ((v15) | 1))
                            _ = v22
                        }
                    } else {
                        v23 := GenS4{F0: (v3 << ((18446744073709551615) & 63)), F1: 344}
                        for fi24 := 0; fi24 < 6; fi24++ {
                            var v25 int8 = v4
                            switch ("\n\t" + "a") {
                            case "":
                                if false {
                                    fmt.Println(fi24)
                                }
                            case "a":
                                v26 := GenS5{F0: "", F1: false, F2: (!false)}
                                _ = v26
                            case "\n\t":
                                v27 := GenS6{F0: (v25 ^ 31)}
                                _ = v27
                            default:
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    v2 = (v2 + v2)
                                    fmt.Println(v15, v4, v2)
                                    var v29 float32 = float32(0.1)
                                    _ = fi28
                                    _ = v29
                                }
                            }
                            var v30 uint = 0
                            _ = fi24
                            _ = v25
                            _ = v30
                        }
                        v31 := GenS4{F0: (v1 - (v3 * 1)), F1: 4294967295}
                        _ = v23
                        _ = v31
                    }
                }
            } else {
                v2 -= 42
                var v32 uint8 = 128
                var v33 float32 = float32((float64((1e10 / 0.1)) * 0.1))
                _ = v32
                _ = v33
            }
            _ = v15
        }
    } else {
        if (1e10 < (0.1 + 1.5)) {
            var v34 uint = 42
            _ = v34
        }
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 uint8 = 127
            _ = fi35
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 1; fi38++ {
            v39 := GenS1{F0: 2}
            var v40 uint8 = 1
            delete(v5, "a")
            if false {
                if false {
                    switch (fi37 | fi38) {
                    case 0:
                        if true {
                            var v41 int32 = -1
                            switch (("\n\t" + (("\n\t" + ("foo bar" + "foo bar")) + (("foo bar" + ("a" + "hello")) + ""))) + "") {
                            case "":
                                if (!true) {
                                    var v42 int64 = -1
                                    _ = v42
                                } else {
                                    var v43 int8 = v3
                                    var v44 uint = uint(v2)
                                    if ((fi37 % ((fi38) | 1)) > (fi37 % ((fi38) | 1))) {
                                        _ = len(v5)
                                        v2 -= 2
                                        var v45 string = ("foo bar" + "a")
                                        _ = v45
                                    } else {
                                        v46 := []int{(-(fi38)), (^(fi37)), -899}
                                        v2 = uint16(v40)
                                        var v47 uint16 = v2
                                        v48 := GenS7{F0: 42}
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                    v4 -= (^(v4))
                                    _ = v43
                                    _ = v44
                                }
                            case "foo bar":
                                var v49 uint8 = v40
                                _ = v49
                            case "hello":
                                switch 9223372036854775807 {
                                case 0:
                                    var v50 float64 = float64(1e10)
                                    v51 := []int{fi37, (-(fi38)), 4294967295, (fi38 / (((fi38 | fi37)) | 1))}
                                    v52 := GenS1{F0: (v40 >> ((235) & 63))}
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                case -1:
                                    if true {
                                        fmt.Println(v4, fi38, v1)
                                    } else {
                                        var v53 string = "a"
                                        _ = v53
                                    }
                                case 9223372036854775807:
                                    var v54 float64 = (float64(1.5) * 0.1)
                                    _ = v54
                                default:
                                    var v55 float64 = float64(3.14159)
                                    _ = v55
                                }
                            default:
                                v40 += v40
                            }
                            var v56 uint8 = (^(v40))
                            _ = v41
                            _ = v56
                        }
                    case -9223372036854775808:
                        _ = len(v5)
                        if false {
                            if (true || false) {
                                if ((v3 / ((-1) | 1)) <= v1) {
                                    var v57 uint64 = uint64(v2)
                                    var v58 uint16 = v2
                                    var v59 float64 = -2.25
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                            }
                        } else {
                            v60 := []int{2147483647, fi38}
                            var v61 uint = uint(v40)
                            if true {
                                v3 += v4
                            } else {
                                if false {
                                    v61 -= v61
                                    var v62 float64 = 3.14159
                                    _ = v62
                                } else {
                                    _ = len(v60)
                                    var v63 float64 = (-2.25 + 0.0)
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v3)
                                    case "foo bar":
                                        if (v4 < 100) {
                                            v64 := GenS8{F0: v40, F1: 1}
                                            if false {
                                                var v65 uint64 = 1
                                                delete(v5, "foo bar")
                                                _ = v65
                                            }
                                            v63 = v63
                                            _ = v64
                                        }
                                    case "hello":
                                        fmt.Println(fi38, v40, v1)
                                    }
                                    _ = v63
                                }
                                v3 += v3
                                v1 += v1
                            }
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(fi37)
                            }
                            v40 -= v40
                            v61 = 255
                            if false {
                                v2 += (v2 * v2)
                            }
                            switch "" {
                            case "a":
                                v1 -= (v4 << (((v61 & 18446744073709551615)) & 63))
                                var v66 uint64 = 42
                                v61 += uint(v2)
                                fmt.Println(v2)
                                _ = v66
                            case "hello":
                                if false {
                                    var v67 int8 = (v4 & 103)
                                    if (!(!(v40 == v40))) {
                                        v68 := GenS9{F0: (("" + "\n\t") + ("a" + "a"))}
                                        _ = v68
                                    } else {
                                        var v69 uint16 = v2
                                        _ = len(v5)
                                        var v70 int64 = -9223372036854775808
                                        _ = v69
                                        _ = v70
                                    }
                                    _ = len(v5)
                                    delete(v5, "hello")
                                    _ = v67
                                } else {
                                    var v71 bool = true
                                    _ = v71
                                }
                            }
                            _ = v60
                            _ = v61
                        }
                        var v72 float64 = 0.1
                        var v73 string = "foo bar"
                        if false {
                            var v74 int8 = v1
                            var v75 int64 = -1
                            _ = v74
                            _ = v75
                        } else {
                            v72 -= (v72 + float64(0.0))
                            fmt.Println(v1, v72, v2)
                        }
                        _ = v72
                        _ = v73
                    }
                    v76 := []int{(-(fi37))}
                    var v77 uint64 = 1
                    _ = v76
                    _ = v77
                }
                switch "foo bar" {
                case "a":
                    var v78 string = "\n\t"
                    var v79 uint64 = 42
                    if (42 == 510) {
                        v80 := GenS10{F0: v40, F1: v79, F2: 1}
                        v81 := []int{(fi38 % ((127) | 1)), fi37}
                        _ = v80
                        _ = v81
                    }
                    _ = v78
                    _ = v79
                case "hello":
                    switch (("\n\t" + "") + "") {
                    case "\n\t":
                        v82 := []int{}
                        _ = v82
                    }
                default:
                    var v83 int32 = -1
                    _ = v83
                }
            } else {
                v84 := GenS11{F0: 255}
                v4 -= (v1 % (((-(v1))) | 1))
                _ = v84
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        v85 := GenS7{F0: 1}
        _ = fi37
        _ = v85
    }
    v86 := []int{4294967295}
    var v87 int8 = v4
    v88 := GenS12{F0: 779, F1: int8(v2)}
    if true {
        switch 4294967295 {
        case 216:
            v2 += 0
        default:
            v89 := GenS8{F0: 127, F1: 0}
            _ = v89
        }
    } else {
        var v90 float64 = 3.14159
        if (2147483648 >= 245) {
            var v91 int64 = -1
            var v92 string = "hello"
            _ = len(v86)
            _ = v91
            _ = v92
        } else {
            var v93 float64 = 1e10
            var v94 int32 = 564
            if (966 == -1) {
                for fi95 := 0; fi95 < 12; fi95++ {
                    if (757 <= v2) {
                        v3 += v4
                        for fi96 := 0; fi96 < 16; fi96++ {
                            if false {
                                v5["\n\t"] = (((-2.25 >= 1e10) || (true && false)) || (707 != 42))
                                if ((3.14159 * 0.1) < -2.25) {
                                    var v97 bool = (0 >= 42)
                                    _ = v97
                                } else {
                                    var v98 uint = 255
                                    v99 := GenS11{F0: uint(v2)}
                                    switch "a" {
                                    case "":
                                        switch "" {
                                        case "a":
                                            var v100 float32 = 1.5
                                            _ = v100
                                        case "":
                                            var v101 int32 = v94
                                            _ = v101
                                        case "\n\t":
                                            v93 = 3.14159
                                        default:
                                            v4 -= (v87 - v3)
                                        }
                                    case "foo bar":
                                        if len(v86) > 0 {
                                            v86[0] = fi95
                                        }
                                    case "a":
                                        fmt.Println(fi96, fi96, v93)
                                    }
                                    _ = v98
                                    _ = v99
                                }
                                fmt.Println(v90, v3, fi96)
                            } else {
                                v102 := v5["a"]
                                if len(v86) >= 2 {
                                    v86 = v86[1:len(v86)]
                                }
                                v103 := GenS13{F0: ((-(fi96)) > (fi96 % ((fi96) | 1))), F1: (fi95 * (fi95 - fi95))}
                                v86 = append(v86, -9223372036854775808)
                                _ = v102
                                _ = v103
                            }
                            v93 = v93
                            _ = fi96
                        }
                        v1 = v87
                    } else {
                        v104 := GenS13{F0: false, F1: 2147483647}
                        _ = v104
                    }
                    _ = fi95
                }
            } else {
                fmt.Println(v3)
                var v105 float64 = v93
                _ = v105
            }
            _ = v93
            _ = v94
        }
        var v106 int8 = 71
        _ = v106
        _ = v90
    }
    if (0 == 42) {
        v107 := GenS8{F0: 9, F1: 2147483647}
        _ = v107
    }
    if true {
        v2 -= v2
    }
    if false {
        if len(v86) > 0 {
            v86[0] = 621
        }
        for fi108 := 0; fi108 < 20; fi108++ {
            if false {
                v109 := GenS12{F0: -1, F1: (v1 | 127)}
                var v110 uint64 = 18446744073709551615
                v110 = (^(v110))
                _ = v109
                _ = v110
            } else {
                v4 += v4
                var v111 int = (^(fi108))
                _ = v111
            }
            _ = fi108
        }
    } else {
        if len(v86) > 0 {
            v86[0] = 936
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v86:", len(v86), v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v86
    _ = v87
    _ = v88
}
