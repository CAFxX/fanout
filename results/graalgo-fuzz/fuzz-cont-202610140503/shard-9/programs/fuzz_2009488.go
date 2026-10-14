package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint; F2 bool }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 int64; F1 bool; F2 int32 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16 }
func genFn0() uint16 {
    v5 := GenS1{F0: 153, F1: 2}
    if (1 > 0) {
        var v6 uint16 = 255
        _ = v6
    }
    switch "foo bar" {
    case "foo bar":
        var v7 float32 = ((float32((float64(-2.25) * 3.14159)) + 1.5) * (0.0 * 1.5))
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 16; fi8++ {
            if ((false && ((false && true) && (10 >= 42))) || false) {
                if true {
                    v9 := GenS1{F0: 42, F1: 42}
                    _ = v9
                }
                if (0 != 42) {
                    var v10 string = "a"
                    v11 := []int{fi8}
                    switch "\n\t" {
                    case "":
                        switch ("" + v10) {
                        case "foo bar":
                            var v12 string = v10
                            _ = v12
                        case "a":
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                    case "a":
                        v10 = v10
                    default:
                        var v13 uint = 42
                        _ = v13
                    }
                    _ = v10
                    _ = v11
                }
                v14 := []int{(fi8 & 9223372036854775807)}
                _ = v14
            } else {
                v15 := GenS1{F0: 2147483647, F1: 0}
                _ = v15
            }
            v16 := GenS2{F0: 0, F1: 595, F2: false}
            if (fi8 <= (fi8 % (((fi8 - (fi8 - fi8))) | 1))) {
                if (!false) {
                    fmt.Println(fi8)
                    var v17 uint = 1
                    _ = v17
                } else {
                    var v18 bool = false
                    _ = v18
                }
                v19 := GenS3{F0: 1, F1: 243, F2: (((0.1 * 1e10) * (((float32(0.1) + -2.25) / -2.25) * (1e10 * 0.0))) >= -2.25)}
                fmt.Println(fi8)
                _ = v19
            } else {
                var v20 int = (fi8 >> ((42) & 63))
                var v21 int64 = 515
                _ = v20
                _ = v21
            }
            _ = fi8
            _ = v16
        }
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint = 1
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 int = fi24
            var v26 string = ("" + ("\n\t" + "hello"))
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    if true {
        if true {
            if (((!(1000 >= 42)) && true) || false) {
                v27 := []int{32768, -113}
                switch ("" + "") {
                case "hello":
                    if (717 != 0) {
                        v28 := GenS4{F0: 305, F1: (!(100 != -1)), F2: 2147483647}
                        _ = v28
                    }
                case "\n\t":
                    var v29 int8 = 1
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 17; fi30++ {
                        if (!false) {
                            var v31 uint32 = 2147483647
                            var v32 int64 = 42
                            _ = v31
                            _ = v32
                        } else {
                            for fi33 := 0; fi33 < 15; fi33++ {
                                fmt.Println(fi33, fi30)
                                var v34 int8 = 100
                                if len(v27) > 0 {
                                    v27[0] = 746
                                }
                                v35 := GenS3{F0: 42, F1: 19, F2: (42 <= 65535)}
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            v27 = append(v27, (fi30 ^ fi30))
                            var v36 uint32 = 2
                            _ = v36
                        }
                        _ = fi30
                    }
                }
                for fi37 := 0; fi37 < 9; fi37++ {
                    switch fi37 {
                    case -9223372036854775808:
                        for fi38 := 0; fi38 < 15; fi38++ {
                            if true {
                                var v39 uint = 4294967295
                                v39 += (v39 * 42)
                                _ = v39
                            } else {
                                var v40 float64 = 0.0
                                var v41 bool = true
                                var v42 int8 = 0
                                var v43 int64 = 42
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            v27 = append(v27, fi38)
                            _ = fi38
                        }
                    case 32768:
                        v44 := GenS5{F0: 128, F1: ((("foo bar" + ("" + ("foo bar" + "\n\t"))) + "foo bar") + "foo bar")}
                        _ = v44
                    case 9223372036854775807:
                        if (0 <= 147) {
                            var v45 uint64 = 18446744073709551615
                            if false {
                                switch (("\n\t" + ("\n\t" + "hello")) + "foo bar") {
                                case "\n\t":
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                case "":
                                    v45 += (v45 + (^(v45)))
                                case "foo bar":
                                    var v46 int16 = 354
                                    _ = v46
                                }
                                v45 = (v45 | 347)
                            } else {
                                if (((1e10 - (0.1 / (0.1 * 3.14159))) != (1e10 * 1e10)) || ("hello" < "foo bar")) {
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    switch -903 {
                                    case 9223372036854775807:
                                        v47 := GenS5{F0: 0, F1: (("foo bar" + "\n\t") + "hello")}
                                        _ = v47
                                    case -9223372036854775808:
                                        var v48 int16 = -1
                                        _ = v48
                                    default:
                                        var v49 float32 = -2.25
                                        _ = v49
                                    }
                                    for fi50 := 0; fi50 < 10; fi50++ {
                                        switch (fi37 + (fi50 - (fi37 % (((fi50 ^ -9223372036854775808)) | 1)))) {
                                        case -338:
                                            v51 := GenS6{F0: 2147483648}
                                            _ = v51
                                        case 32767:
                                            v52 := GenS7{F0: -9223372036854775808, F1: (!(false || (false || false)))}
                                            _ = v52
                                        case 995:
                                            v45 -= (^(v45))
                                        }
                                        if (255 < 42) {
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                        } else {
                                            v27 = append(v27, fi50)
                                            fmt.Println(fi50)
                                            var v53 uint32 = 0
                                            _ = v53
                                        }
                                        _ = fi50
                                    }
                                }
                                if ((!true) || (0 == 0)) {
                                    switch (fi37 - fi37) {
                                    case 197:
                                        v45 = (v45 + (v45 >> ((190) & 63)))
                                    case 116:
                                        v45 = (v45 & 694)
                                    case 1:
                                        v45 = 18446744073709551615
                                    default:
                                        var v54 uint32 = 1
                                        _ = v54
                                    }
                                    var v55 string = ("foo bar" + "a")
                                    _ = v55
                                } else {
                                    v45 -= (v45 >> ((4294967295) & 63))
                                    switch (fi37 * (fi37 + fi37)) {
                                    case 1:
                                        var v56 int32 = 391
                                        _ = v56
                                    case 42:
                                        var v57 uint8 = 0
                                        _ = v57
                                    }
                                    v45 = (^(v45))
                                }
                                var v58 uint64 = v45
                                _ = v58
                            }
                            _ = v45
                        } else {
                            for fi59 := 0; fi59 < 11; fi59++ {
                                var v60 uint16 = 895
                                var v61 uint = 65535
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                                fmt.Println(v61)
                                _ = fi59
                                _ = v60
                                _ = v61
                            }
                        }
                    }
                    _ = fi37
                }
                _ = v27
            } else {
                switch 4294967295 {
                case 778:
                    v62 := []int{4294967295, -769, 0}
                    _ = v62
                }
                if false {
                    v63 := []int{}
                    var v64 int16 = -32768
                    fmt.Println(v64)
                    _ = v63
                    _ = v64
                } else {
                    v65 := []int{42}
                    _ = v65
                }
                if ((false && true) || true) {
                    var v66 int32 = -2147483648
                    var v67 int8 = 42
                    _ = v66
                    _ = v67
                } else {
                    var v68 int16 = 1000
                    _ = v68
                }
            }
            if false {
                v69 := []int{42, 404, 32768, 903}
                for fi70 := 0; fi70 < 10; fi70++ {
                    if (2 != 0) {
                        var v71 int64 = -630
                        v72 := GenS8{F0: ("a" + "hello")}
                        fmt.Println(v71)
                        _ = v71
                        _ = v72
                    }
                    _ = fi70
                }
                var v73 int8 = 73
                _ = v69
                _ = v73
            }
            var v74 uint64 = 71
            _ = v74
        }
        var v75 int64 = 0
        var v76 float64 = -2.25
        _ = v75
        _ = v76
    } else {
        var v77 uint16 = 42
        var v78 float64 = -2.25
        for fi79 := 0; fi79 < 4; fi79++ {
            v80 := GenS1{F0: 1, F1: 18446744073709551615}
            fmt.Println(fi79, v78)
            _ = fi79
            _ = v80
        }
        _ = v77
        _ = v78
    }
    _ = v5
    return 255
}
func genFn1(p0 uint16) int64 {
    fmt.Println(p0)
    var v81 uint = 399
    v82 := []int{42}
    _ = len(v82)
    v83 := GenS8{F0: (("hello" + "") + (("\n\t" + (("foo bar" + ("hello" + ("" + "\n\t"))) + (("a" + (("hello" + "\n\t") + "foo bar")) + (((((("foo bar" + "") + ((("foo bar" + "foo bar") + "hello") + "")) + ("a" + "hello")) + "\n\t") + "a") + "a")))) + "\n\t"))}
    _ = p0
    _ = v81
    _ = v82
    _ = v83
    return -1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -100
    var v3 string = (("" + "hello") + ("hello" + ("foo bar" + "")))
    v4 := make(map[int]bool)
    var v84 uint16 = 1
    for fi85 := 0; fi85 < 5; fi85++ {
        v86 := GenS5{F0: 4294967295, F1: ((("\n\t" + ("" + "a")) + v3) + ((("a" + (v3 + ((v3 + (((v3 + v3) + (v3 + "foo bar")) + "\n\t")) + "\n\t"))) + "foo bar") + "\n\t"))}
        delete(v4, -339)
        _ = fi85
        _ = v86
    }
    for fi87 := 0; fi87 < 10; fi87++ {
        v88 := GenS2{F0: (v1 - v1), F1: 255, F2: true}
        if true {
            v89 := GenS6{F0: 496}
            _ = v89
        } else {
            v90 := GenS9{F0: v84}
            var v91 uint = 255
            _ = v90
            _ = v91
        }
        v2 = (v2 << ((65535) & 63))
        _ = fi87
        _ = v88
    }
    v84 = (v84 % ((65535) | 1))
    var v92 uint = 104
    v3 = v3
    v93 := GenS2{F0: (v1 << (((v92 << ((v92) & 63))) & 63)), F1: v92, F2: false}
    var v94 uint16 = 542
    var v95 uint8 = 1
    var v96 int64 = v1
    _ = len(v4)
    v84 += (v94 - v94)
    if (100 == -128) {
        if false {
            v97 := GenS3{F0: v92, F1: 128, F2: true}
            _ = v97
        }
        switch v3 {
        case "\n\t":
            if false {
                var v98 int32 = -2147483648
                for fi99 := 0; fi99 < 17; fi99++ {
                    v4[fi99] = false
                    v4[(fi99 + (fi99 | 2147483647))] = (65535 <= v92)
                    switch (fi99 - 128) {
                    case 42:
                        var v100 int16 = (v2 % (((v2 + -32768)) | 1))
                        _ = v100
                    case 32768:
                        v84 = 409
                    case 4294967295:
                        v101 := []int{(fi99 & 4294967295), (fi99 % ((fi99) | 1))}
                        _ = v101
                    default:
                        if (!false) {
                            v92 -= 339
                        }
                    }
                    _ = fi99
                }
                switch "hello" {
                case "hello":
                    v3 = v3
                default:
                    if true {
                        var v102 int = 2147483647
                        var v103 int64 = v1
                        _ = v102
                        _ = v103
                    }
                }
                _ = v98
            } else {
                switch 42 {
                case 1:
                    v1 = (v1 * -1)
                case -2147483648:
                    var v104 string = "foo bar"
                    _ = v104
                case 32767:
                    var v105 uint16 = 138
                    _ = v105
                }
            }
        }
        var v106 string = (v3 + v3)
        _ = v106
    } else {
        fmt.Println(v96, v84, v94)
        v92 -= (^(v92))
    }
    v107 := v4[4294967295]
    var v108 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v84:", v84)
    fmt.Println("v92:", v92)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v84
    _ = v92
    _ = v93
    _ = v94
    _ = v95
    _ = v96
}
