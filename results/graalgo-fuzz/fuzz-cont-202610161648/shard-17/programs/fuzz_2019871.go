package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    if ((false && true) && true) {
        var v4 string = "hello"
        _ = v4
    }
    var v5 uint = 42
    if false {
        var v6 uint = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            if true {
                var v8 bool = false
                _ = v8
            }
            _ = fi7
        }
        v9 := GenS1{F0: -9223372036854775808, F1: 42, F2: 42}
        _ = v9
    }
    switch -547 {
    case -899:
        var v10 float32 = 3.14159
        _ = v10
    case 128:
        if ("hello" < "foo bar") {
            var v11 int8 = 0
            var v12 uint = (v5 >> ((v5) & 63))
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: -9223372036854775808, F1: 42, F2: 2}
                var v15 int16 = -103
                var v16 uint32 = 2147483647
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = v11
            _ = v12
        } else {
            v17 := GenS2{F0: 42, F1: 0, F2: -9223372036854775808}
            _ = v17
        }
    }
    switch "\n\t" {
    case "hello":
        switch 127 {
        case 32768:
            if (true && true) {
                v5 = 215
                switch "hello" {
                case "foo bar":
                    v3[-845] = false
                default:
                    if (!false) {
                        if (838 > -32768) {
                            var v18 uint16 = 32767
                            delete(v3, -2147483648)
                            _ = v18
                        }
                        var v19 float64 = 3.14159
                        switch -549 {
                        case 632:
                            v20 := []int{}
                            _ = v20
                        case 1:
                            switch "foo bar" {
                            case "foo bar":
                                v19 = v19
                            case "\n\t":
                                if (42 < -1) {
                                    v19 = v19
                                    if (1 >= 974) {
                                        var v21 float32 = 0.1
                                        var v22 uint = (v5 + (v5 ^ (v5 >> ((v5) & 63))))
                                        switch "foo bar" {
                                        case "a":
                                            v23 := []int{515, -472}
                                            _ = v23
                                        case "":
                                            var v24 uint = v22
                                            _ = v24
                                        case "hello":
                                            for fi25 := 0; fi25 < 16; fi25++ {
                                                delete(v3, -282)
                                                v21 += v21
                                                _ = fi25
                                            }
                                        }
                                        _ = v21
                                        _ = v22
                                    } else {
                                        v19 = float64(1e10)
                                        v26 := []int{32767, -689}
                                        _ = v26
                                    }
                                    var v27 float32 = ((0.1 * 0.0) - float32(v19))
                                    _ = v27
                                } else {
                                    if ((77 <= 100) && (42 >= -843)) {
                                        for fi28 := 0; fi28 < 2; fi28++ {
                                            if (v19 <= (0.1 * -2.25)) {
                                                for fi29 := 0; fi29 < 14; fi29++ {
                                                    v5 += (v5 / ((1) | 1))
                                                    _ = fi29
                                                }
                                            } else {
                                                fmt.Println(v5, v19)
                                                switch ("hello" + "hello") {
                                                case "\n\t":
                                                    v19 = float64(3.14159)
                                                case "":
                                                    var v30 int16 = 87
                                                    _ = v30
                                                case "a":
                                                    var v31 uint16 = 0
                                                    _ = v31
                                                default:
                                                    fmt.Println(fi28, fi28)
                                                }
                                                if ((!(!(!(false && (!true))))) && true) {
                                                    fmt.Println(fi28, v5, v5)
                                                    fmt.Println(v5, fi28, v19)
                                                }
                                            }
                                            _ = fi28
                                        }
                                        var v32 int64 = 1
                                        if false {
                                            v5 += v5
                                            v5 = (v5 & (v5 % (((v5 << (((v5 - v5)) & 63))) | 1)))
                                        } else {
                                            v33 := []int{}
                                            v5 = (v5 - 1)
                                            _ = v33
                                        }
                                        _ = v32
                                    } else {
                                        if (!false) {
                                            fmt.Println(v5, v19)
                                            switch -9223372036854775808 {
                                            case 0:
                                                fmt.Println(v19, v5)
                                            case 4294967295:
                                                v34 := []int{4294967295, -2147483648, 9223372036854775807}
                                                _ = v34
                                            case -1:
                                                if (false && (!true)) {
                                                    if (!false) {
                                                        var v35 uint64 = 18446744073709551615
                                                        v5 += 0
                                                        if false {
                                                            fmt.Println(v19)
                                                        }
                                                        _ = v35
                                                    }
                                                    v36 := GenS1{F0: -1, F1: 242, F2: 0}
                                                    _ = v36
                                                } else {
                                                    for fi37 := 0; fi37 < 6; fi37++ {
                                                        v38 := GenS1{F0: -1, F1: -961, F2: 128}
                                                        v39 := GenS2{F0: int32(v5), F1: 1, F2: 42}
                                                        var v40 uint64 = 42
                                                        _ = fi37
                                                        _ = v38
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                }
                                            }
                                            var v41 uint8 = 127
                                            _ = v41
                                        } else {
                                            v42 := GenS1{F0: 42, F1: -808, F2: 255}
                                            if (3.14159 <= float64(0.0)) {
                                                var v43 int32 = 759
                                                v44 := GenS1{F0: -9223372036854775808, F1: -544, F2: uint8(v5)}
                                                var v45 int16 = 42
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            } else {
                                                var v46 uint = v5
                                                if ("" >= "foo bar") {
                                                    _ = len(v3)
                                                } else {
                                                    var v47 uint32 = 945
                                                    _ = v47
                                                }
                                                var v48 int64 = -9223372036854775808
                                                _ = v46
                                                _ = v48
                                            }
                                            if (false || (!(!true))) {
                                                switch 606 {
                                                case 32767:
                                                    _ = len(v1)
                                                case 1:
                                                    for fi49 := 0; fi49 < 17; fi49++ {
                                                        v50 := GenS3{F0: 9223372036854775807, F1: 664}
                                                        _ = fi49
                                                        _ = v50
                                                    }
                                                }
                                            } else {
                                                v5 = 18446744073709551615
                                                var v51 uint64 = 2
                                                _ = v51
                                            }
                                            _ = v42
                                        }
                                    }
                                    var v52 bool = true
                                    _ = v52
                                }
                            }
                        case 4294967295:
                            v53 := []int{-137, 128, 30, 4294967295}
                            _ = v53
                        }
                        _ = v19
                    }
                }
                v5 += 313
            }
        case 0:
            if true {
                var v54 uint64 = uint64(v5)
                _ = v54
            }
        }
    case "\n\t":
        switch 128 {
        case -9223372036854775808:
            fmt.Println(v5)
        case 128:
            v55 := []int{-94, -1, 2147483647, 9223372036854775807}
            _ = v55
        case 733:
            v56 := GenS4{F0: -1, F1: 4294967295}
            _ = v56
        default:
            if (false && false) {
                for fi57 := 0; fi57 < 5; fi57++ {
                    if ((!false) || true) {
                        var v58 string = ("a" + "foo bar")
                        _ = v58
                    }
                    var v59 int64 = 848
                    _ = fi57
                    _ = v59
                }
                v60 := GenS2{F0: -2147483648, F1: 1, F2: -1}
                v61 := []int{}
                _ = v60
                _ = v61
            } else {
                v62 := GenS2{F0: 42, F1: 9223372036854775807, F2: -573}
                var v63 string = "\n\t"
                var v64 int = 420
                _ = v62
                _ = v63
                _ = v64
            }
        }
    case "foo bar":
        v5 -= 0
    default:
        v5 += (v5 >> ((v5) & 63))
    }
    if ((0 == 9223372036854775807) || (!(!false))) {
        v5 += v5
        var v65 int = -62
        v65 -= v65
        _ = v65
    } else {
        fmt.Println(v5)
        var v66 float32 = 1e10
        fmt.Println(v5, v5)
        _ = v66
    }
    v5 = 4294967295
    var v67 uint16 = 2
    fmt.Println(v67, v67)
    v68 := GenS4{F0: 817, F1: 1}
    v69 := []int{-2147483648, 131, 32767, -9223372036854775808}
    v67 = (v67 * (v67 | 1))
    var v70 float32 = 1.5
    if ((-431 == 1) && (1.5 <= 0.0)) {
        fmt.Println(v5)
        v71 := GenS2{F0: 42, F1: -9223372036854775808, F2: 0}
        _ = v71
    } else {
        switch "foo bar" {
        case "\n\t":
            var v72 int8 = 127
            _ = v72
        case "hello":
            if len(v69) >= 2 {
                v69 = v69[1:len(v69)]
            }
        case "a":
            v5 = v5
        }
    }
    _ = len(v69)
    v73 := GenS4{F0: 1, F1: 2}
    var v74 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v73
    _ = v74
}
