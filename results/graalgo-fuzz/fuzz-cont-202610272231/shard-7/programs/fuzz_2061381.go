package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 bool; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 int16; F1 uint32; F2 uint8 }
func genFn0(p0 uint32, p1 int8, p2 bool) uint {
    if p2 {
        var v6 uint32 = p0
        _ = v6
    }
    v7 := []int{-1, 888, 32768, 267}
    var v8 uint32 = 2147483647
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16, p1 uint64, p2 uint) uint64 {
    var v9 float64 = 0.1
    var v10 uint = (p2 % ((147) | 1))
    v9 -= v9
    if (p2 <= v10) {
        var v11 int16 = p0
        v11 -= (v11 / ((689) | 1))
        _ = v11
    } else {
        switch -716 {
        case 862:
            if ((v9 != (0.0 + ((v9 + 1e10) / (v9 - 1.5)))) || ((p0 - (p0 | p0)) != -32768)) {
                var v12 uint = p2
                if (p0 != (^(p0))) {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        for fi14 := 0; fi14 < 12; fi14++ {
                            if (false || (false && (false && (255 != 65535)))) {
                                v15 := []int{(fi14 % ((fi13) | 1)), fi13, 128}
                                switch fi14 {
                                case -2147483648:
                                    var v16 float32 = 1.5
                                    _ = v16
                                default:
                                    _ = len(v15)
                                }
                                fmt.Println(v10, fi13, fi13)
                                if (p0 == 399) {
                                    fmt.Println(fi14)
                                    v15 = append(v15, (fi13 % (((fi14 | (fi13 & -772))) | 1)))
                                }
                                _ = v15
                            } else {
                                v17 := []int{-1, fi14, fi14, (-(fi14))}
                                var v18 string = "foo bar"
                                _ = v17
                                _ = v18
                            }
                            if true {
                                var v19 int = (fi14 ^ fi13)
                                fmt.Println(p2)
                                var v20 string = "a"
                                v21 := []int{v19, fi13, 32767, v19}
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                v22 := []int{561, (fi13 - (fi14 & -1))}
                                var v23 string = (("\n\t" + "a") + "hello")
                                v23 = "\n\t"
                                var v24 bool = ((((v9 - 0.1) + -2.25) - 0.1) > 0.1)
                                p0 = p0
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            _ = fi14
                        }
                        p2 += (v12 >> ((1) & 63))
                        if (53 == 1) {
                            for fi25 := 0; fi25 < 7; fi25++ {
                                var v26 uint32 = 1
                                v27 := []int{}
                                switch ("\n\t" + ("hello" + "foo bar")) {
                                case "hello":
                                    _ = len(v27)
                                case "a":
                                    fmt.Println(v26, v9, v9)
                                case "":
                                    fmt.Println(v9, v10)
                                default:
                                    if ((!(false && true)) && ((true && (1.5 > (0.0 * 1.5))) || true)) {
                                        fmt.Println(v9, fi13)
                                        if true {
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(v26)
                                            case "a":
                                                var v28 uint8 = 255
                                                _ = v28
                                            case "foo bar":
                                                var v29 string = ""
                                                _ = v29
                                            default:
                                                v30 := GenS1{F0: v26}
                                                _ = v30
                                            }
                                        } else {
                                            var v31 uint16 = 255
                                            _ = v31
                                        }
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        if ((true || (v26 >= v26)) || true) {
                                            if len(v27) > 0 {
                                                v27[0] = (-(fi25))
                                            }
                                            if false {
                                                v32 := GenS1{F0: (v26 % (((v26 | v26)) | 1))}
                                                if ("" == "") {
                                                    if true {
                                                        if (!((v10 <= (v10 << ((p2) & 63))) || (v26 != 4294967295))) {
                                                            if ("a" == ("\n\t" + (("a" + ("\n\t" + "hello")) + "hello"))) {
                                                                fmt.Println(p2, v10)
                                                            } else {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    v33 := GenS2{F0: 1, F1: (v26 - v26)}
                                                                    _ = v33
                                                                case "foo bar":
                                                                    if len(v27) > 0 {
                                                                        v27[0] = (-(fi13))
                                                                    }
                                                                }
                                                                v10 = (^(v12))
                                                                v12 += genFn0((v26 + 397), 29, (fi13 < (fi25 * (fi13 | (fi25 - fi13)))))
                                                            }
                                                            switch 32767 {
                                                            case 42:
                                                                switch (fi13 % (((fi25 + fi13)) | 1)) {
                                                                case -37:
                                                                    p1 -= 463
                                                                }
                                                            case 9223372036854775807:
                                                                p1 -= 4294967295
                                                                _ = len(v27)
                                                            case -174:
                                                                if ("a" <= "foo bar") {
                                                                    fmt.Println(v26)
                                                                } else {
                                                                    fmt.Println(p1, v12, v12)
                                                                    var v34 uint32 = (v26 / (((v26 >> ((genFn0(v26, 100, (("a" + "hello") < "foo bar"))) & 63))) | 1))
                                                                    _ = v34
                                                                }
                                                            }
                                                        }
                                                        var v35 int64 = -9223372036854775808
                                                        if len(v27) >= 2 {
                                                            v27 = v27[1:len(v27)]
                                                        }
                                                        _ = v35
                                                    }
                                                    v36 := GenS1{F0: (v26 - 192)}
                                                    var v37 string = ("a" + "foo bar")
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    switch fi25 {
                                                    case 998:
                                                        p2 -= (v12 | genFn0(2147483648, 127, ((!(!(uint64(p2) < (p1 << ((genFn0(65535, 0, (127 != 100))) & 63))))) || (1 == 121))))
                                                    case 42:
                                                        v10 = (v12 / ((uint(v26)) | 1))
                                                    }
                                                    if len(v27) >= 2 {
                                                        v27 = v27[1:len(v27)]
                                                    }
                                                    var v38 uint32 = (v26 % ((v26) | 1))
                                                    _ = v38
                                                }
                                                v10 -= (v12 >> ((4294967295) & 63))
                                                if false {
                                                    var v39 float64 = 0.1
                                                    var v40 int = 2147483647
                                                    switch fi25 {
                                                    case 32767:
                                                        var v41 uint16 = 2
                                                        fmt.Println(v26, v41)
                                                        _ = v41
                                                    }
                                                    var v42 uint32 = (^(v26))
                                                    var v43 int32 = -1
                                                    fmt.Println(v12, p0, v10)
                                                    _ = v39
                                                    _ = v40
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = v32
                                            }
                                            v44 := GenS1{F0: v26}
                                            _ = v44
                                        } else {
                                            p0 = p0
                                            if true {
                                                var v45 bool = (false && (32767 > fi13))
                                                fmt.Println(v45, v12, v45)
                                                _ = v45
                                            }
                                            if false {
                                                if true {
                                                    v46 := GenS1{F0: (v26 + (v26 * 2))}
                                                    _ = v46
                                                } else {
                                                    var v47 int64 = -350
                                                    v10 -= (p2 << ((65535) & 63))
                                                    _ = v47
                                                }
                                                var v48 uint16 = 0
                                                _ = v48
                                            } else {
                                                switch 2147483647 {
                                                case 522:
                                                    var v49 uint32 = v26
                                                    if ((fi25 / (((fi25 & 4294967295)) | 1)) >= (fi13 - (fi13 + fi25))) {
                                                        if false {
                                                            var v50 int32 = 2147483647
                                                            var v51 string = ((("foo bar" + "") + ("\n\t" + "foo bar")) + "\n\t")
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    } else {
                                                        switch (("\n\t" + "\n\t") + "a") {
                                                        case "hello":
                                                            if len(v27) > 0 {
                                                                v27[0] = (fi13 << ((v12) & 63))
                                                            }
                                                        case "\n\t":
                                                            var v52 int16 = p0
                                                            _ = v52
                                                        case "a":
                                                            v9 = 0.0
                                                        }
                                                        if (-1 < 127) {
                                                            if len(v27) >= 2 {
                                                                v27 = v27[1:len(v27)]
                                                            }
                                                            var v53 float32 = 0.0
                                                            _ = v53
                                                        } else {
                                                            p0 = p0
                                                        }
                                                    }
                                                    var v54 int8 = 42
                                                    _ = v49
                                                    _ = v54
                                                default:
                                                    var v55 uint = (v12 / ((v12) | 1))
                                                    _ = v55
                                                }
                                                v9 += float64(1e10)
                                                var v56 float32 = (1.5 / 1.5)
                                                _ = v56
                                            }
                                        }
                                    default:
                                        var v57 float32 = 0.1
                                        _ = v57
                                    }
                                    var v58 int64 = 42
                                    _ = v58
                                }
                                var v59 int16 = (-(p0))
                                if false {
                                    if true {
                                        v27 = append(v27, (fi13 / ((-577) | 1)))
                                        v60 := GenS3{F0: 1, F1: false, F2: "\n\t"}
                                        v61 := GenS4{F0: (v26 << (((p2 ^ (v10 << ((genFn0((v26 ^ v26), 127, true)) & 63)))) & 63))}
                                        var v62 uint64 = p1
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    } else {
                                        fmt.Println(fi25, fi25)
                                        p1 = (^(p1))
                                        v63 := GenS4{F0: 42}
                                        _ = v63
                                    }
                                }
                                _ = fi25
                                _ = v26
                                _ = v27
                                _ = v59
                            }
                        } else {
                            if ((!true) && false) {
                                p2 = (v10 ^ v12)
                                p0 = p0
                            } else {
                                fmt.Println(p1)
                            }
                            v64 := GenS1{F0: 2}
                            v10 = (v12 ^ (v10 / ((p2) | 1)))
                            _ = v64
                        }
                        _ = fi13
                    }
                } else {
                    var v65 uint = 65535
                    if ("\n\t" > ((("hello" + ("" + "\n\t")) + "a") + ("a" + "hello"))) {
                        for fi66 := 0; fi66 < 3; fi66++ {
                            v67 := []int{(^(fi66))}
                            var v68 string = "hello"
                            for fi69 := 0; fi69 < 17; fi69++ {
                                var v70 uint = (v10 + (v12 - uint(p1)))
                                _ = fi69
                                _ = v70
                            }
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                        v65 -= (^(v10))
                        v9 += ((0.0 + (v9 / float64((-2.25 - 3.14159)))) * (0.0 - ((3.14159 + 3.14159) / v9)))
                    }
                    var v71 int32 = 512
                    _ = v65
                    _ = v71
                }
                v10 -= (p2 >> (((^(v10))) & 63))
                _ = v12
            }
        case 0:
            switch 9223372036854775807 {
            case -122:
                v72 := GenS1{F0: 42}
                _ = v72
            case 128:
                v9 += v9
            default:
                if ((v9 > float64(1e10)) && (uint8(p2) != 28)) {
                    for fi73 := 0; fi73 < 1; fi73++ {
                        switch (fi73 % (((fi73 - int(p0))) | 1)) {
                        case 32768:
                            var v74 bool = false
                            _ = v74
                        case 772:
                            fmt.Println(p1, v9, v10)
                        case 0:
                            var v75 uint = p2
                            _ = v75
                        }
                        _ = fi73
                    }
                    var v76 string = "foo bar"
                    _ = v76
                }
            }
        }
        var v77 uint = (p2 << ((p2) & 63))
        _ = v77
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p1
}

func main() {
    var v1 uint64 = 87
    v2 := make(map[uint32]bool)
    var v3 int = 42
    var v4 float64 = -2.25
    v5 := make(map[int]bool)
    switch ("" + (("hello" + (("hello" + "") + ("a" + ""))) + "hello")) {
    case "hello":
        v78 := GenS2{F0: 0, F1: 2147483647}
        _ = v78
    case "a":
        var v79 float32 = 1e10
        _ = v79
    default:
        var v80 int32 = 1
        _ = v80
    }
    v3 += (v3 / (((-(v3))) | 1))
    for fi81 := 0; fi81 < 5; fi81++ {
        fmt.Println(v1, v1)
        _ = fi81
    }
    for fi82 := 0; fi82 < 14; fi82++ {
        v83 := GenS5{F0: (!(93 <= -71))}
        _ = fi82
        _ = v83
    }
    var v84 int16 = 42
    v85 := []int{v3, 4294967295}
    switch ("foo bar" + "a") {
    case "foo bar":
        fmt.Println(v1)
    case "\n\t":
        for fi86 := 0; fi86 < 10; fi86++ {
            for fi87 := 0; fi87 < 16; fi87++ {
                fmt.Println(v3)
                if false {
                    if false {
                        v3 -= v3
                        var v88 uint = 125
                        v84 += (-(v84))
                        _ = v88
                    } else {
                        fmt.Println(v84, fi86)
                        v89 := GenS6{F0: -59, F1: (^(v1))}
                        _ = v89
                    }
                    v90 := GenS7{F0: v84, F1: 42, F2: 2}
                    _ = v90
                } else {
                    v3 += -9223372036854775808
                    v3 += (v3 >> ((0) & 63))
                    var v91 float32 = 1.5
                    fmt.Println(v4, fi86, fi87)
                    _ = v91
                }
                v92 := v2[65535]
                _ = fi87
                _ = v92
            }
            _ = fi86
        }
    default:
        var v93 int64 = 0
        _ = v93
    }
    var v94 bool = true
    if false {
        v3 -= (v3 & 832)
        _ = len(v85)
    } else {
        var v95 uint16 = 255
        if v94 {
            if len(v85) >= 2 {
                v85 = v85[1:len(v85)]
            }
            var v96 string = ("foo bar" + "hello")
            _ = v96
        } else {
            var v97 uint64 = v1
            _ = v97
        }
        _ = v95
    }
    for fi98 := 0; fi98 < 1; fi98++ {
        v99 := GenS7{F0: (-(v84)), F1: 849, F2: 0}
        _ = fi98
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v84:", v84)
    fmt.Println("slicelen v85:", len(v85), v85)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
    _ = v94
}
