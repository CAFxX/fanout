package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int; F2 int32 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int) bool {
    var v5 uint64 = 889
    switch ("foo bar" + "\n\t") {
    case "a":
        if false {
            p0 += p0
        }
    default:
        v6 := GenS1{F0: v5, F1: (p0 << ((1) & 63)), F2: 776}
        _ = v6
    }
    v7 := GenS1{F0: v5, F1: p0, F2: 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return true
}
func genFn1() int64 {
    switch 0 {
    case 640:
        switch "hello" {
        case "foo bar":
            v8 := GenS1{F0: 18446744073709551615, F1: 127, F2: 705}
            _ = v8
        default:
            if ((32767 > -1) || (-1 != 42)) {
                var v9 int16 = -476
                _ = v9
            }
        }
    case 1:
        var v10 int32 = 0
        _ = v10
    case 4294967295:
        for fi11 := 0; fi11 < 6; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int = (fi12 ^ fi11)
                v14 := []int{fi11}
                fmt.Println(v13)
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint = 2
                switch (fi11 | fi15) {
                case 397:
                    var v17 uint16 = 32767
                    _ = v17
                default:
                    v18 := GenS2{F0: (^(fi11)), F1: 42}
                    _ = v18
                }
                var v19 int8 = 5
                _ = fi15
                _ = v16
                _ = v19
            }
            var v20 uint64 = 274
            _ = fi11
            _ = v20
        }
    default:
        if (-83 <= -117) {
            for fi21 := 0; fi21 < 14; fi21++ {
                if (false && (!(!false))) {
                    var v22 int32 = 105
                    var v23 uint16 = 65535
                    var v24 int8 = 127
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    v25 := []int{(fi21 | fi21), fi21, fi21}
                    var v26 int16 = 523
                    _ = v25
                    _ = v26
                }
                var v27 float32 = 3.14159
                _ = fi21
                _ = v27
            }
            if false {
                switch 4294967295 {
                case 32767:
                    if true {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            if (590 != 0) {
                                var v29 string = "hello"
                                v30 := []int{(fi28 - fi28), (fi28 % ((fi28) | 1)), fi28, fi28}
                                _ = v29
                                _ = v30
                            }
                            var v31 int32 = -1
                            _ = fi28
                            _ = v31
                        }
                        var v32 bool = (42 < 42)
                        _ = v32
                    }
                case -116:
                    var v33 int8 = 122
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 4; fi34++ {
                        var v35 int32 = 2147483647
                        var v36 bool = true
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
                var v37 int16 = 32767
                v37 = (^(v37))
                _ = v37
            } else {
                var v38 uint = 255
                if (true || false) {
                    var v39 string = "a"
                    _ = v39
                } else {
                    v38 -= (v38 | 890)
                    var v40 uint = (v38 + (^(v38)))
                    v41 := GenS3{F0: false, F1: 9223372036854775807}
                    _ = v40
                    _ = v41
                }
                _ = v38
            }
        } else {
            v42 := []int{32768}
            if len(v42) >= 2 {
                v42 = v42[1:len(v42)]
            }
            var v43 uint64 = 0
            _ = v42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 13; fi44++ {
        switch (fi44 / (((fi44 - fi44)) | 1)) {
        case 726:
            var v45 float64 = float64(((1e10 - 1.5) + 0.1))
            _ = v45
        case -2147483648:
            v46 := GenS2{F0: 1, F1: 0}
            _ = v46
        case 723:
            v47 := []int{}
            _ = v47
        default:
            switch "foo bar" {
            case "":
                fmt.Println(fi44)
            case "hello":
                var v48 bool = (2 == 607)
                _ = v48
            case "foo bar":
                var v49 float32 = 3.14159
                _ = v49
            default:
                v50 := []int{fi44}
                _ = v50
            }
        }
        _ = fi44
    }
    if true {
        var v51 uint32 = 387
        var v52 string = ""
        fmt.Println(v52, v51)
        _ = v51
        _ = v52
    }
    return 9223372036854775807
}

func main() {
    v1 := []int{32768, 42, 32768}
    var v2 string = "a"
    var v3 bool = false
    var v4 int32 = 324
    var v53 bool = (((4294967295 >= 357) || v3) || v3)
    v3 = (v3 || ((float64(0.1) - (0.1 / (-2.25 + 0.1))) <= (1.5 / 1e10)))
    if len(v1) > 0 {
        v1[0] = 81
    }
    var v54 bool = v3
    v55 := GenS1{F0: 42, F1: 5, F2: 934}
    if len(v1) > 0 {
        v1[0] = -552
    }
    v56 := GenS3{F0: false, F1: 128}
    fmt.Println(v2, v53, v54)
    if v54 {
        fmt.Println(v4, v53)
        var v57 int64 = -9223372036854775808
        var v58 uint32 = 42
        _ = v57
        _ = v58
    } else {
        fmt.Println(v54, v54, v2)
        v3 = v3
    }
    if v3 {
        var v59 int8 = 100
        fmt.Println(v59, v4)
        for fi60 := 0; fi60 < 18; fi60++ {
            if v54 {
                switch 1 {
                case -9223372036854775808:
                    if (32767 != 924) {
                        if true {
                            var v61 int8 = (v59 / (((v59 >> ((18446744073709551615) & 63))) | 1))
                            _ = v61
                        } else {
                            v3 = false
                            v62 := GenS1{F0: 1, F1: 128, F2: 0}
                            _ = v62
                        }
                    }
                case 128:
                    v63 := GenS4{F0: 629}
                    _ = v63
                case 2147483647:
                    for fi64 := 0; fi64 < 11; fi64++ {
                        v2 = "a"
                        v59 = 100
                        v65 := GenS5{F0: 642, F1: fi60, F2: v4}
                        _ = fi64
                        _ = v65
                    }
                default:
                    _ = len(v1)
                }
            } else {
                _ = len(v1)
                _ = len(v1)
                v2 = (v2 + (((("\n\t" + (v2 + v2)) + v2) + v2) + ("foo bar" + v2)))
            }
            var v66 bool = true
            _ = fi60
            _ = v66
        }
        _ = v59
    }
    switch v2 {
    case "hello":
        v67 := GenS1{F0: 457, F1: 65, F2: 1}
        _ = v67
    case "foo bar":
        var v68 int64 = 770
        _ = v68
    case "a":
        if true {
            fmt.Println(v2, v53, v54)
            _ = len(v1)
            fmt.Println(v54)
        }
    }
    if (v3 && v54) {
        switch 0 {
        case 32768:
            var v69 bool = v54
            _ = v69
        default:
            switch v2 {
            case "a":
                switch -734 {
                case 1:
                    for fi70 := 0; fi70 < 20; fi70++ {
                        for fi71 := 0; fi71 < 19; fi71++ {
                            if ((v4 * v4) == (v4 % ((v4) | 1))) {
                                if v3 {
                                    fmt.Println(v53, v2)
                                    v72 := GenS4{F0: 345}
                                    _ = v72
                                } else {
                                    switch fi70 {
                                    case 4294967295:
                                        v53 = true
                                    case 32767:
                                        var v73 uint8 = 147
                                        _ = v73
                                    case 128:
                                        fmt.Println(v54, v3)
                                    default:
                                        var v74 float64 = (3.14159 - 1e10)
                                        _ = v74
                                    }
                                    if (!v54) {
                                        v75 := GenS6{F0: 4294967295, F1: 100, F2: 65535}
                                        _ = v75
                                    }
                                }
                            }
                            fmt.Println(v3, v53, v2)
                            var v76 int8 = 1
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi71
                            _ = v76
                        }
                        if (42 != 989) {
                            var v77 uint32 = 2147483648
                            v78 := GenS7{F0: fi70}
                            _ = v77
                            _ = v78
                        }
                        if true {
                            fmt.Println(v4, v53)
                            if true {
                                var v79 int64 = genFn1()
                                _ = v79
                            }
                        }
                        _ = fi70
                    }
                case -2147483648:
                    switch 247 {
                    case -2147483648:
                        var v80 uint = 2
                        _ = v80
                    }
                case 127:
                    v2 = (v2 + (v2 + (v2 + v2)))
                default:
                    v81 := GenS8{F0: 32767}
                    _ = v81
                }
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v82 uint16 = 98
        _ = v82
    }
    switch v2 {
    case "a":
        var v83 float64 = float64((1.5 * 1.5))
        _ = v83
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi84 := 0; fi84 < 14; fi84++ {
        if len(v1) > 0 {
            v1[0] = fi84
        }
        for fi85 := 0; fi85 < 15; fi85++ {
            fmt.Println(fi85, v2, v54)
            fmt.Println(v54, v4)
            if (v3 || (!false)) {
                var v86 int16 = -32768
                switch (^(fi84)) {
                case 2147483647:
                    v54 = v3
                }
                if true {
                    fmt.Println(v3, v53)
                } else {
                    v87 := GenS9{F0: 18446744073709551615, F1: genFn1()}
                    _ = v87
                }
                _ = v86
            }
            _ = fi85
        }
        v88 := GenS10{F0: 173}
        _ = fi84
        _ = v88
    }
    var v89 int32 = (v4 + v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v89
}
