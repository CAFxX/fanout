package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int64; F2 uint32 }
type GenS7 struct { F0 int8; F1 int16; F2 int16 }
type GenS8 struct { F0 string }

func main() {
    var v1 int = -1
    var v2 bool = (255 <= 973)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v5 uint8 = 167
    fmt.Println(v2, v5, v2)
    _ = len(v4)
    if v2 {
        for fi6 := 0; fi6 < 3; fi6++ {
            if true {
                v7 := []int{9223372036854775807, v1, v1}
                v1 = (^(v1))
                _ = v7
            }
            var v8 uint32 = 2147483648
            if ((-32768 == int16(v5)) && (!true)) {
                v9 := []int{(v1 | (fi6 + v1)), (fi6 + v1), 91, v1}
                v10 := GenS1{F0: v1, F1: (-(v1)), F2: 856}
                var v11 uint16 = 1
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi6
            _ = v8
        }
    } else {
        if (3.14159 <= (3.14159 - 1e10)) {
            var v12 float64 = 1.5
            _ = v12
        } else {
            v13 := []int{}
            _ = v13
        }
        _ = len(v4)
    }
    fmt.Println(v2, v2)
    v14 := GenS2{F0: (v1 / (((v1 << ((373) & 63))) | 1)), F1: 4294967295}
    v2 = false
    var v15 float64 = 3.14159
    var v16 bool = false
    v3[42] = true
    if false {
        var v17 int16 = 42
        v5 -= 160
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        switch "a" {
        case "hello":
            var v19 int16 = -1
            _ = v19
        default:
            var v20 float32 = 0.0
            _ = v20
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := []int{fi21}
            var v23 string = "\n\t"
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            v15 -= v15
            _ = fi24
        }
        _ = fi18
    }
    if true {
        for fi25 := 0; fi25 < 1; fi25++ {
            for fi26 := 0; fi26 < 14; fi26++ {
                switch "\n\t" {
                case "\n\t":
                    v27 := GenS2{F0: 728, F1: v1}
                    _ = v27
                case "foo bar":
                    _ = len(v4)
                case "a":
                    var v28 bool = v16
                    _ = v28
                }
                var v29 uint64 = 1
                _ = fi26
                _ = v29
            }
            _ = fi25
        }
        if v16 {
            _ = len(v3)
            if (1 <= 42) {
                fmt.Println(v2)
            } else {
                for fi30 := 0; fi30 < 8; fi30++ {
                    fmt.Println(v16, v15)
                    var v31 uint64 = 536
                    var v32 uint16 = 261
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                for fi33 := 0; fi33 < 8; fi33++ {
                    for fi34 := 0; fi34 < 14; fi34++ {
                        if (0 >= 4294967295) {
                            v35 := []int{(v1 * v1), -171}
                            _ = v35
                        } else {
                            switch "a" {
                            case "foo bar":
                                fmt.Println(v5, v16)
                            case "\n\t":
                                if (!(42 != 32767)) {
                                    if false {
                                        v36 := GenS1{F0: fi33, F1: v1, F2: -766}
                                        v2 = v2
                                        if v16 {
                                            v4["\n\t"] = v16
                                            fmt.Println(v5)
                                            v1 += (fi33 - -61)
                                            delete(v3, 42)
                                        } else {
                                            if (0.0 > -2.25) {
                                                fmt.Println(v2, fi34)
                                            } else {
                                                if v16 {
                                                    v5 = (v5 / (((v5 ^ v5)) | 1))
                                                    v37 := GenS2{F0: fi34, F1: (fi33 ^ 1)}
                                                    _ = v37
                                                }
                                                v38 := []int{-1, fi34, fi34}
                                                var v39 uint = 23
                                                v40 := GenS1{F0: fi33, F1: fi34, F2: 789}
                                                if ((((v5 << ((v39) & 63)) <= v5) && v16) && v2) {
                                                    v5 += 1
                                                    fmt.Println(v16, fi34)
                                                    if (1 < 42) {
                                                        fmt.Println(fi33, fi34)
                                                        v5 += (v5 | (v5 & (^(v5))))
                                                        switch (fi34 | fi33) {
                                                        case 652:
                                                            var v41 int = 2147483647
                                                            _ = v41
                                                        default:
                                                            v42 := GenS3{F0: 0}
                                                            _ = v42
                                                        }
                                                    } else {
                                                        switch "\n\t" {
                                                        case "":
                                                            var v43 int = fi34
                                                            v1 = (v1 - (fi33 | 9223372036854775807))
                                                            var v44 int32 = -1
                                                            _ = v43
                                                            _ = v44
                                                        case "hello":
                                                            v45 := GenS1{F0: (v1 >> ((v39) & 63)), F1: fi34, F2: 1}
                                                            _ = v45
                                                        }
                                                    }
                                                }
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        _ = v36
                                    } else {
                                        v46 := GenS3{F0: (v5 << ((951) & 63))}
                                        _ = v46
                                    }
                                    v47 := GenS4{F0: 392, F1: 36, F2: ""}
                                    v15 = float64(((-2.25 - float32((v15 * 0.1))) / 0.1))
                                    _ = v47
                                }
                            default:
                                if v2 {
                                    var v48 float64 = 0.0
                                    _ = len(v4)
                                    if (!true) {
                                        v49 := GenS2{F0: (fi34 << ((18446744073709551615) & 63)), F1: (fi33 >> ((4294967295) & 63))}
                                        if ((^(v5)) >= 0) {
                                            var v50 uint16 = 0
                                            _ = v50
                                        } else {
                                            v2 = ((^(v5)) >= (v5 - v5))
                                            if v16 {
                                                if v2 {
                                                    var v51 int8 = 127
                                                    _ = v51
                                                } else {
                                                    var v52 float64 = v15
                                                    v48 += v15
                                                    _ = v52
                                                }
                                                fmt.Println(v5)
                                                switch (fi33 / ((-9223372036854775808) | 1)) {
                                                case 4294967295:
                                                    fmt.Println(fi34)
                                                }
                                                if v16 {
                                                    switch ("hello" + ("foo bar" + "foo bar")) {
                                                    case "\n\t":
                                                        fmt.Println(fi34, v1)
                                                    default:
                                                        if (!("foo bar" == "\n\t")) {
                                                            if v2 {
                                                                fmt.Println(fi33)
                                                                var v53 string = ((("\n\t" + "hello") + ("" + (("hello" + "a") + "\n\t"))) + "hello")
                                                                fmt.Println(v53)
                                                                _ = v53
                                                            }
                                                            v54 := GenS3{F0: (v5 - v5)}
                                                            if false {
                                                                v55 := GenS1{F0: -1, F1: 1, F2: 0}
                                                                if true {
                                                                    var v56 string = ("" + "hello")
                                                                    fmt.Println(v56)
                                                                    _ = v56
                                                                } else {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        fmt.Println(v2)
                                                                    }
                                                                }
                                                                fmt.Println(v2)
                                                                _ = v55
                                                            }
                                                            _ = v54
                                                        } else {
                                                            switch (("foo bar" + "") + (((("" + "a") + "a") + "") + (("a" + "") + "a"))) {
                                                            case "hello":
                                                                var v57 int8 = 100
                                                                _ = v57
                                                            }
                                                            if v16 {
                                                                if (!v16) {
                                                                    switch ("foo bar" + "\n\t") {
                                                                    case "\n\t":
                                                                        fmt.Println(fi34, v2, fi33)
                                                                        v58 := GenS5{F0: 18446744073709551615}
                                                                        _ = v58
                                                                    default:
                                                                        fmt.Println(v5, fi33)
                                                                    }
                                                                    switch fi34 {
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v15)
                                                                    case 32767:
                                                                        var v59 int = fi34
                                                                        var v60 uint16 = 32767
                                                                        _ = v59
                                                                        _ = v60
                                                                    case 128:
                                                                        v16 = (v16 && (!(false || v2)))
                                                                    default:
                                                                        if false {
                                                                            if true {
                                                                                v16 = (2 >= 763)
                                                                            }
                                                                        } else {
                                                                            var v61 int8 = -1
                                                                            v62 := GenS6{F0: 399, F1: 0, F2: 2}
                                                                            fmt.Println(v61, v15, v5)
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                    }
                                                                    v16 = v2
                                                                    delete(v4, "")
                                                                    v63 := []int{}
                                                                    v3[715] = v16
                                                                    var v64 uint64 = 4294967295
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                            } else {
                                                                v2 = (!true)
                                                                v65 := GenS7{F0: 100, F1: 1, F2: 30}
                                                                if (255 <= 1) {
                                                                    var v66 int64 = -1
                                                                    _ = v66
                                                                }
                                                                _ = v65
                                                            }
                                                            v67 := []int{-54, (fi33 % ((v1) | 1)), (v1 << ((18446744073709551615) & 63)), fi33}
                                                            _ = v67
                                                        }
                                                    }
                                                } else {
                                                    var v68 uint8 = (v5 ^ v5)
                                                    var v69 uint64 = 18446744073709551615
                                                    if (!(v1 > fi33)) {
                                                        var v70 uint8 = 2
                                                        var v71 float32 = 0.0
                                                        _ = v70
                                                        _ = v71
                                                    } else {
                                                        switch (fi34 / (((fi33 >> ((215) & 63))) | 1)) {
                                                        case 182:
                                                            var v72 uint8 = (v68 + 134)
                                                            _ = v72
                                                        case 9223372036854775807:
                                                            if v16 {
                                                                var v73 bool = false
                                                                _ = v73
                                                            } else {
                                                                v5 -= v5
                                                                fmt.Println(v69, v15)
                                                            }
                                                        case 127:
                                                            var v74 uint32 = 2147483647
                                                            _ = v74
                                                        default:
                                                            delete(v4, (("" + (("\n\t" + ("foo bar" + "")) + "")) + "hello"))
                                                        }
                                                        var v75 uint64 = (v69 << ((107) & 63))
                                                        _ = v75
                                                    }
                                                    _ = v68
                                                    _ = v69
                                                }
                                                v15 += 3.14159
                                            } else {
                                                switch 1 {
                                                case 128:
                                                    delete(v3, 4294967295)
                                                    fmt.Println(v5, v16, v5)
                                                    v2 = true
                                                    v76 := GenS1{F0: fi33, F1: v1, F2: -910}
                                                    var v77 string = "a"
                                                    _ = v76
                                                    _ = v77
                                                case -404:
                                                    var v78 string = "\n\t"
                                                    _ = v78
                                                }
                                                var v79 uint8 = v5
                                                v80 := v3[569]
                                                if true {
                                                    v81 := GenS2{F0: (-(fi34)), F1: 435}
                                                    _ = v81
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    v80 = true
                                                }
                                                v80 = v16
                                                _ = v79
                                                _ = v80
                                            }
                                        }
                                        var v82 bool = (!v16)
                                        v3[2147483647] = false
                                        switch "hello" {
                                        case "foo bar":
                                            var v83 string = (("foo bar" + (("foo bar" + "foo bar") + "")) + "")
                                            _ = v83
                                        case "a":
                                            var v84 int8 = 127
                                            _ = v84
                                        case "\n\t":
                                            var v85 int16 = 0
                                            _ = v85
                                        default:
                                            if (65535 <= 845) {
                                                var v86 int16 = 1000
                                                var v87 float64 = v15
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                        _ = v49
                                        _ = v82
                                    }
                                    _ = v48
                                } else {
                                    v15 -= 1e10
                                    v15 -= 1.5
                                }
                            }
                            v2 = v16
                            v15 += (v15 / float64((float32(0.1) - 1.5)))
                        }
                        _ = fi34
                    }
                    if v16 {
                        v2 = ((v1 % ((v1) | 1)) != (fi33 ^ 9223372036854775807))
                        fmt.Println(v16, fi33, v15)
                        var v88 int32 = 1
                        _ = v88
                    }
                    _ = fi33
                }
                switch "hello" {
                case "foo bar":
                    v89 := GenS7{F0: 42, F1: 42, F2: 641}
                    _ = v89
                }
            }
        }
        v90 := []int{(v1 - (v1 * v1))}
        _ = v90
    } else {
        var v91 uint = 341
        _ = v91
    }
    var v92 int64 = -9223372036854775808
    switch "a" {
    case "":
        if ((2147483647 >= v1) && v2) {
            v93 := GenS8{F0: (("" + "hello") + "a")}
            v94 := []int{v1, v1, (v1 % (((v1 / ((v1) | 1))) | 1))}
            fmt.Println(v2)
            _ = v93
            _ = v94
        }
    }
    v2 = ((^(v92)) != v92)
    if (42 > v5) {
        var v95 uint16 = 5
        var v96 int16 = -32768
        _ = v95
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v92
}
