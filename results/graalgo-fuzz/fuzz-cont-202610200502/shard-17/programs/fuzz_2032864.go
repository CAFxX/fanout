package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint32 = 1
    var v2 int64 = 42
    v3 := make(map[string]bool)
    var v4 uint = 0
    v5 := make(map[int]string)
    var v6 uint32 = v1
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v6 += v1
            _ = len(v5)
            v9 := GenS1{F0: -52, F1: -1}
            _ = fi8
            _ = v9
        }
        var v10 int16 = -32768
        var v11 uint8 = 127
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 int16 = 1000
    if false {
        if ((v2 | v2) < (v2 >> ((v4) & 63))) {
            v4 = v4
            delete(v3, "foo bar")
            var v13 uint64 = 4294967295
            _ = v13
        } else {
            var v14 uint8 = 94
            _ = len(v3)
            _ = v14
        }
        var v15 int64 = v2
        var v16 uint16 = 339
        _ = v15
        _ = v16
    }
    var v17 float32 = 1.5
    switch ("hello" + "a") {
    case "hello":
        fmt.Println(v2, v12)
    default:
        v4 = 545
    }
    if false {
        switch -9223372036854775808 {
        case 42:
            v4 += v4
        case -9223372036854775808:
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 uint = (^(v4))
                v1 = (^(v6))
                _ = fi18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            v6 += v6
            for fi21 := 0; fi21 < 3; fi21++ {
                if (((42 != 32767) || true) && ((v4 == (^(v4))) || (false || false))) {
                    if false {
                        switch 127 {
                        case 2147483647:
                            var v22 int8 = 100
                            _ = v22
                        case 9223372036854775807:
                            v17 = v17
                        }
                        if (!true) {
                            if ((v2 % ((618) | 1)) > v2) {
                                v23 := GenS1{F0: 0, F1: (v12 + (v12 << (((v4 << ((558) & 63))) & 63)))}
                                var v24 uint32 = v6
                                switch fi20 {
                                case 2147483647:
                                    var v25 uint = (v4 >> (((v4 * v4)) & 63))
                                    _ = v25
                                case -644:
                                    v26 := []int{1}
                                    v27 := GenS1{F0: 127, F1: (v12 % (((v12 << ((65535) & 63))) | 1))}
                                    if false {
                                        v2 = v2
                                        if ((!true) || (!((v1 * v6) >= v1))) {
                                            v28 := GenS1{F0: -1, F1: (v12 % ((1) | 1))}
                                            var v29 bool = (0 <= -139)
                                            v6 = (v6 | (v1 + v24))
                                            _ = v28
                                            _ = v29
                                        } else {
                                            v30 := GenS1{F0: 100, F1: v12}
                                            var v31 int64 = v2
                                            var v32 int32 = 2147483647
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                    } else {
                                        v2 += -661
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        v33 := GenS1{F0: 13, F1: -32768}
                                        v24 = (^(v1))
                                        _ = v33
                                    }
                                    _ = v26
                                    _ = v27
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v34 int32 = 2147483647
                                        _ = v34
                                    case "hello":
                                        var v35 float32 = (((0.1 + 3.14159) * (v17 - (v17 * float32((-2.25 + float64(1e10)))))) - v17)
                                        _ = v35
                                    }
                                    var v36 string = "a"
                                    var v37 uint8 = 21
                                    fmt.Println(v1, v36, v2)
                                    if ((false || (false && false)) && (469 != 1)) {
                                        v1 = v24
                                        if (v36 == "\n\t") {
                                            var v38 int8 = 100
                                            _ = v38
                                        } else {
                                            if (!true) {
                                                v2 = (v2 ^ (v2 % ((9223372036854775807) | 1)))
                                            } else {
                                                var v39 float32 = v17
                                                _ = v39
                                            }
                                        }
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                        fmt.Println(v1, v2, fi21)
                        if true {
                            _ = len(v3)
                            if ((-910 <= int32(v4)) || ((false && false) || false)) {
                                v4 = 42
                                v40 := []int{}
                                _ = v40
                            }
                            var v41 float64 = 1e10
                            _ = v41
                        } else {
                            if (((float64(((float32(float64((1e10 * 0.0))) * v17) + (v17 - 0.1))) - 1.5) * (1.5 / 1.0)) < (0.0 * (0.1 / ((-2.25 - 3.14159) + (1.5 * (3.14159 * ((-2.25 / 0.1) + 1e10))))))) {
                                var v42 float64 = (-2.25 / 0.1)
                                _ = v42
                            }
                            fmt.Println(v1, v1, fi20)
                            var v43 float32 = (v17 + (((3.14159 / v17) - (v17 / v17)) / v17))
                            _ = v43
                        }
                    } else {
                        v44 := GenS2{F0: 1, F1: -1, F2: v6}
                        delete(v3, "\n\t")
                        var v45 float32 = -2.25
                        v46 := []int{}
                        fmt.Println(fi20, v2)
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    v12 -= v12
                    var v47 uint64 = 2
                    _ = v47
                } else {
                    v48 := GenS3{F0: v6, F1: 18446744073709551615, F2: (v2 | 42)}
                    if (!((!(!(v17 <= ((v17 + (v17 / ((v17 * -2.25) + (0.0 * -2.25)))) / 1.0)))) && (false || true))) {
                        v49 := GenS2{F0: v2, F1: -128, F2: v6}
                        switch fi20 {
                        case 38:
                            var v50 int64 = (v2 | (v2 & -183))
                            _ = v50
                        case -39:
                            if false {
                                var v51 int64 = v2
                                v52 := GenS2{F0: (v2 << ((v4) & 63)), F1: -1, F2: (^(v1))}
                                fmt.Println(fi21)
                                _ = v51
                                _ = v52
                            }
                        }
                        v17 = v17
                        _ = v49
                    } else {
                        delete(v3, ("\n\t" + "hello"))
                    }
                    v12 = (v12 + v12)
                    _ = v48
                }
                v4 += 1
                var v53 uint = (v4 << ((v4) & 63))
                switch (^(fi20)) {
                case 4294967295:
                    v54 := GenS2{F0: 9223372036854775807, F1: 1, F2: v1}
                    _ = v54
                case -521:
                    var v55 uint = 208
                    _ = v55
                case 177:
                    v1 -= (v1 + (v6 / ((v6) | 1)))
                default:
                    if false {
                        if false {
                            v56 := GenS2{F0: (v2 * 0), F1: -86, F2: (v1 << ((1) & 63))}
                            _ = v56
                        }
                        v57 := GenS3{F0: 609, F1: 4294967295, F2: v2}
                        if false {
                            v58 := GenS1{F0: -1, F1: (v12 >> ((42) & 63))}
                            _ = v58
                        } else {
                            var v59 int = fi21
                            _ = v59
                        }
                        _ = v57
                    }
                }
                _ = fi21
                _ = v53
            }
            v60 := []int{}
            _ = fi20
            _ = v60
        }
        v61 := GenS3{F0: v6, F1: 2, F2: v2}
        _ = v61
    }
    var v62 bool = false
    switch "foo bar" {
    case "a":
        v63 := GenS1{F0: 0, F1: v12}
        _ = v63
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        for fi65 := 0; fi65 < 19; fi65++ {
            var v66 int = 4294967295
            if (-9223372036854775808 == v2) {
                v67 := GenS2{F0: v2, F1: 100, F2: (^(v6))}
                v12 += v12
                _ = v67
            }
            fmt.Println(v4, fi64)
            _ = fi65
            _ = v66
        }
        if (v62 && (!false)) {
            _ = len(v3)
            var v68 int16 = v12
            _ = v68
        }
        for fi69 := 0; fi69 < 5; fi69++ {
            var v70 string = ("a" + "\n\t")
            _ = fi69
            _ = v70
        }
        _ = fi64
    }
    var v71 int16 = v12
    v1 = uint32(v4)
    var v72 uint8 = 255
    for fi73 := 0; fi73 < 13; fi73++ {
        var v74 int64 = (-(v2))
        fmt.Println(v1, v17)
        if (v2 >= int64(v6)) {
            var v75 int16 = v71
            _ = v75
        } else {
            if v62 {
                if v62 {
                    v76 := GenS4{F0: int32(v12)}
                    _ = v76
                }
            }
            v77 := []int{fi73}
            _ = v77
        }
        _ = fi73
        _ = v74
    }
    v12 -= (v71 % ((v12) | 1))
    if ((1e10 <= 3.14159) || v62) {
        switch -9223372036854775808 {
        case 2147483647:
            if ((!((!(!v62)) || true)) && false) {
                if v62 {
                    for fi78 := 0; fi78 < 13; fi78++ {
                        var v79 bool = v62
                        v12 -= v12
                        _ = fi78
                        _ = v79
                    }
                    switch -814 {
                    case 2147483647:
                        for fi80 := 0; fi80 < 6; fi80++ {
                            v81 := []int{(fi80 >> ((18446744073709551615) & 63)), (^(fi80)), -9223372036854775808, (^(fi80))}
                            var v82 int64 = (-(v2))
                            _ = fi80
                            _ = v81
                            _ = v82
                        }
                    case -290:
                        var v83 uint64 = 4294967295
                        _ = v83
                    case 553:
                        v17 = 1.5
                    }
                    var v84 float32 = v17
                    _ = v84
                } else {
                    var v85 uint16 = 42
                    switch "hello" {
                    case "hello":
                        if v62 {
                            _ = len(v5)
                        }
                    case "a":
                        for fi86 := 0; fi86 < 8; fi86++ {
                            for fi87 := 0; fi87 < 17; fi87++ {
                                var v88 uint64 = 2
                                switch (fi87 >> (((^(v4))) & 63)) {
                                case -9223372036854775808:
                                    v2 = (v2 & v2)
                                case 128:
                                    fmt.Println(v62, v17, v4)
                                    if v62 {
                                        v89 := []int{(fi87 << (((^(v4))) & 63)), (fi87 * fi86), (fi87 ^ (fi87 % (((fi87 % (((fi86 >> (((v4 >> (((v4 << ((483) & 63))) & 63))) & 63))) | 1))) | 1))), fi87}
                                        _ = v89
                                    }
                                    if v62 {
                                        v90 := GenS1{F0: -1, F1: (-(v71))}
                                        var v91 uint = uint(v88)
                                        var v92 float64 = float64(-2.25)
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    } else {
                                        v93 := GenS4{F0: 0}
                                        _ = v93
                                    }
                                case 4294967295:
                                    var v94 uint8 = v72
                                    _ = v94
                                default:
                                    v95 := GenS3{F0: uint32(v88), F1: v88, F2: (v2 / ((int64(v1)) | 1))}
                                    _ = v95
                                }
                                v62 = v62
                                _ = fi87
                                _ = v88
                            }
                            var v96 string = "hello"
                            v3[v96] = (!v62)
                            _ = fi86
                            _ = v96
                        }
                    case "foo bar":
                        for fi97 := 0; fi97 < 17; fi97++ {
                            for fi98 := 0; fi98 < 3; fi98++ {
                                v6 = (v6 % ((84) | 1))
                                v99 := []int{(fi98 / (((fi97 | (^(fi97)))) | 1)), fi98, fi97, (fi97 / ((fi97) | 1))}
                                _ = fi98
                                _ = v99
                            }
                            _ = fi97
                        }
                    default:
                        if false {
                            fmt.Println(v17)
                        }
                    }
                    _ = v85
                }
                v72 = (v72 | (v72 * (v72 % (((v72 / ((v72) | 1))) | 1))))
            }
        case -1:
            _ = len(v3)
        case 128:
            v72 -= v72
        default:
            v62 = (-44 != 100)
        }
        v6 = (v6 * v6)
    } else {
        var v100 uint32 = v1
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v71
    _ = v72
}
