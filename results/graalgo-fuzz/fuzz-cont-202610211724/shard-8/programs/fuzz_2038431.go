package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 int8) int {
    var v5 uint16 = 1
    v6 := []int{-572, 497, 720}
    switch "a" {
    case "\n\t":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 562
}
func genFn1() bool {
    var v7 uint = 4294967295
    v7 = v7
    var v8 float32 = -2.25
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    v3 := []uint32{2, 4294967295}
    v4 := make(map[uint32]bool)
    if false {
        var v9 uint32 = 392
        v10 := GenS1{F0: 90, F1: 32767, F2: 2147483647}
        _ = v10
        _ = v9
    }
    var v11 int32 = 991
    for fi12 := 0; fi12 < 20; fi12++ {
        switch fi12 {
        case 32768:
            v13 := GenS1{F0: 0, F1: -138, F2: (fi12 << ((392) & 63))}
            _ = v13
        case 9223372036854775807:
            var v14 uint64 = 530
            _ = v14
        case -2147483648:
            var v15 int32 = v11
            _ = v15
        default:
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 uint64 = 1
                v11 -= 512
                switch (fi12 & 1) {
                case 32768:
                    v17 += (v17 % ((v17) | 1))
                case 32767:
                    v18 := GenS1{F0: 1, F1: 32767, F2: fi12}
                    _ = v18
                default:
                    v11 -= v11
                }
                _ = fi16
                _ = v17
            }
        }
        _ = fi12
    }
    v19 := GenS1{F0: 100, F1: -822, F2: -1}
    var v20 uint32 = 4294967295
    fmt.Println(v11, v11)
    v11 = (^(v11))
    fmt.Println(v20, v11)
    var v21 float32 = -2.25
    if true {
        v22 := GenS1{F0: 1, F1: -985, F2: -9223372036854775808}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 float64 = 1.5
            var v25 int32 = 349
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 int64 = 9223372036854775807
        v27 := GenS1{F0: -76, F1: -32768, F2: -1}
        _ = v26
        _ = v27
    }
    if false {
        for fi28 := 0; fi28 < 19; fi28++ {
            switch "a" {
            case "\n\t":
                var v29 int16 = 1
                _ = v29
            default:
                for fi30 := 0; fi30 < 3; fi30++ {
                    if false {
                        fmt.Println(v20, fi28)
                        switch "foo bar" {
                        case "a":
                            v21 += float32((1.5 * (1.5 / ((-2.25 * 1e10) + 0.1))))
                        case "foo bar":
                            switch "hello" {
                            case "":
                                v31 := GenS2{F0: 1, F1: 0}
                                var v32 uint32 = 2147483648
                                v11 += (v11 & (v11 * v11))
                                _ = v31
                                _ = v32
                            }
                        }
                        if ((!false) || ((4294967295 < 0) && (true && false))) {
                            v2["hello"] = (fi28 * fi28)
                            v33 := GenS3{F0: 42}
                            fmt.Println(v21)
                            _ = v33
                        }
                    }
                    var v34 int8 = 100
                    var v35 int64 = 42
                    _ = len(v2)
                    _ = fi30
                    _ = v34
                    _ = v35
                }
            }
            fmt.Println(v21)
            var v36 float64 = (-2.25 - ((float64(v21) * 1e10) + 0.0))
            _ = fi28
            _ = v36
        }
        switch "foo bar" {
        case "a":
            if ((-(v11)) == 0) {
                _ = len(v1)
                if false {
                    var v37 bool = false
                    if ((true && true) || false) {
                        v38 := GenS1{F0: 100, F1: 0, F2: -276}
                        if true {
                            for fi39 := 0; fi39 < 15; fi39++ {
                                var v40 int32 = v11
                                var v41 float32 = v21
                                if v37 {
                                    var v42 uint8 = 7
                                    _ = v42
                                }
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            if v37 {
                                for fi43 := 0; fi43 < 1; fi43++ {
                                    for fi44 := 0; fi44 < 2; fi44++ {
                                        v21 += v21
                                        var v45 string = "\n\t"
                                        var v46 int16 = 485
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = len(v3)
                                    }
                                    _ = fi43
                                }
                            }
                            var v47 int64 = 1
                            _ = v47
                        } else {
                            var v48 uint32 = 566
                            if false {
                                for fi49 := 0; fi49 < 5; fi49++ {
                                    if (!(true && false)) {
                                        if v37 {
                                            v50 := GenS1{F0: 42, F1: 1000, F2: fi49}
                                            var v51 float32 = (1e10 * (((v21 - float32(0.1)) * 3.14159) - v21))
                                            _ = v50
                                            _ = v51
                                        } else {
                                            v20 = v20
                                            switch (fi49 & (-(fi49))) {
                                            case 10:
                                                var v52 int16 = 949
                                                _ = v52
                                            case 32768:
                                                if (550 >= 855) {
                                                    v20 = 4294967295
                                                }
                                            case 347:
                                                switch (fi49 << ((2) & 63)) {
                                                case -9223372036854775808:
                                                    for fi53 := 0; fi53 < 19; fi53++ {
                                                        v21 = (v21 + float32((float64((v21 - 1.5)) / 3.14159)))
                                                        if len(v3) > 0 {
                                                            v3[0] = v20
                                                        }
                                                        _ = fi53
                                                    }
                                                case 0:
                                                    var v54 int16 = 1000
                                                    _ = v54
                                                case 1:
                                                    var v55 int8 = -13
                                                    _ = v55
                                                }
                                            default:
                                                v37 = v37
                                            }
                                            if (false && (1 <= 32767)) {
                                                v20 = 2
                                                v21 += ((v21 * v21) / v21)
                                                var v56 string = "a"
                                                _ = v56
                                            }
                                        }
                                        var v57 int16 = -1
                                        _ = v57
                                    } else {
                                        v11 += -52
                                        v58 := GenS3{F0: 42}
                                        for fi59 := 0; fi59 < 3; fi59++ {
                                            v60 := GenS3{F0: 255}
                                            _ = fi59
                                            _ = v60
                                        }
                                        _ = v58
                                    }
                                    if v37 {
                                        var v61 uint8 = 116
                                        if (true || true) {
                                            var v62 uint64 = 1
                                            _ = v62
                                        } else {
                                            var v63 int16 = -32768
                                            v64 := GenS3{F0: 2}
                                            for fi65 := 0; fi65 < 13; fi65++ {
                                                var v66 string = "a"
                                                _ = fi65
                                                _ = v66
                                            }
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = v61
                                    }
                                    switch fi49 {
                                    case 2147483647:
                                        var v67 uint32 = 388
                                        _ = v67
                                    case 9223372036854775807:
                                        var v68 uint8 = 162
                                        _ = v68
                                    case -1:
                                        for fi69 := 0; fi69 < 19; fi69++ {
                                            var v70 int16 = 32767
                                            v71 := GenS3{F0: 42}
                                            var v72 int16 = 42
                                            _ = fi69
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        }
                                    default:
                                        if (v37 || v37) {
                                            v73 := GenS4{F0: 42, F1: 4294967295}
                                            for fi74 := 0; fi74 < 5; fi74++ {
                                                var v75 uint = 55
                                                _ = fi74
                                                _ = v75
                                            }
                                            _ = v73
                                        } else {
                                            var v76 int64 = 9223372036854775807
                                            var v77 float64 = ((0.1 / float64(v21)) + 1e10)
                                            fmt.Println(v76, v37, v20)
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                    _ = fi49
                                }
                                if v37 {
                                    v78 := GenS2{F0: -1, F1: -9223372036854775808}
                                    if false {
                                        switch (("a" + ((("" + "") + ("" + "hello")) + "")) + "") {
                                        case "":
                                            _ = len(v3)
                                        }
                                        var v79 int32 = v11
                                        v80 := v2[("hello" + "foo bar")]
                                        _ = v79
                                        _ = v80
                                    }
                                    var v81 int32 = v11
                                    _ = v78
                                    _ = v81
                                }
                            } else {
                                fmt.Println(v21)
                                switch "a" {
                                case "":
                                    if (42 > 18446744073709551615) {
                                        v82 := GenS3{F0: 2}
                                        fmt.Println(v11, v11)
                                        var v83 int32 = v11
                                        _ = v82
                                        _ = v83
                                    } else {
                                        v48 = v48
                                        fmt.Println(v11, v11)
                                    }
                                case "foo bar":
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v84 int16 = 680
                                        _ = v84
                                    }
                                case "a":
                                    for fi85 := 0; fi85 < 9; fi85++ {
                                        v21 = v21
                                        v86 := GenS5{F0: 122}
                                        v87 := GenS3{F0: 2}
                                        _ = fi85
                                        _ = v86
                                        _ = v87
                                    }
                                }
                                v88 := v1[("foo bar" + ("" + ""))]
                                _ = v88
                            }
                            _ = v48
                        }
                        _ = v38
                    } else {
                        v37 = ((!v37) && (!(!v37)))
                        var v89 uint8 = 102
                        var v90 uint64 = 1
                        _ = v89
                        _ = v90
                    }
                    _ = v37
                }
                fmt.Println(v21)
            } else {
                if len(v3) > 0 {
                    v3[0] = v20
                }
                v20 = (v20 + (v20 | 2147483648))
                switch 1 {
                case -9223372036854775808:
                    if false {
                        var v91 int64 = -1
                        for fi92 := 0; fi92 < 15; fi92++ {
                            var v93 uint64 = 1
                            _ = fi92
                            _ = v93
                        }
                        _ = v91
                    } else {
                        switch int(v11) {
                        case 4294967295:
                            for fi94 := 0; fi94 < 11; fi94++ {
                                delete(v1, (("foo bar" + (("\n\t" + "a") + "")) + "foo bar"))
                                if (586 < 255) {
                                    var v95 uint = 698
                                    switch ("\n\t" + ((("a" + "") + ("" + "\n\t")) + "\n\t")) {
                                    case "foo bar":
                                        var v96 string = "\n\t"
                                        _ = v96
                                    case "a":
                                        fmt.Println(v20)
                                    }
                                    _ = v95
                                } else {
                                    fmt.Println(v20, fi94)
                                    var v97 uint32 = v20
                                    _ = v97
                                }
                                _ = fi94
                            }
                        default:
                            v21 = (v21 + v21)
                        }
                    }
                case -931:
                    var v98 uint16 = 797
                    _ = v98
                case 1:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    _ = len(v1)
                }
            }
        case "hello":
            _ = len(v1)
        default:
            v11 = v11
        }
    }
    if (false || false) {
        v99 := GenS6{F0: v11, F1: 9223372036854775807, F2: 1}
        var v100 uint = 18446744073709551615
        v21 -= (((-2.25 - v21) + v21) - v21)
        _ = v100
        _ = v99
    }
    v11 = (-(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
