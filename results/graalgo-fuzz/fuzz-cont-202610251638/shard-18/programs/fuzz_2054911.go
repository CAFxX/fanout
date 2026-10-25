package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 uint; F1 int16; F2 uint }
type GenS10 struct { F0 string }
func genFn0() int64 {
    v6 := []int{0, 32768, 396, 4294967295}
    v7 := GenS1{F0: 32767}
    switch 0 {
    case 127:
        var v8 int32 = -2147483648
        _ = v8
    case -2147483648:
        var v9 bool = false
        _ = v9
    case 128:
        if true {
            var v10 int32 = 1
            _ = v10
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 uint = 1
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                switch fi13 {
                case 0:
                    var v14 uint = 192
                    _ = v14
                default:
                    var v15 uint32 = 2
                    _ = v15
                }
                switch "hello" {
                case "\n\t":
                    if true {
                        if len(v6) > 0 {
                            v6[0] = fi13
                        }
                    } else {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            if len(v6) > 0 {
                                v6[0] = fi13
                            }
                            if (((0.1 / 1.5) * (0.1 - 1e10)) != float64((0.0 + 0.0))) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            var v17 float32 = 1.5
                            _ = fi16
                            _ = v17
                        }
                    }
                default:
                    fmt.Println(fi13)
                }
                _ = fi13
            }
        }
    }
    var v18 uint64 = 1
    if (v18 < v18) {
        if true {
            _ = len(v6)
        } else {
            v18 += 829
            v19 := GenS2{F0: 137, F1: (^(v18))}
            v18 -= (v18 | v18)
            _ = v19
        }
        switch 48 {
        case 4294967295:
            fmt.Println(v18)
        case 228:
            v6 = append(v6, 2147483647)
        case 127:
            v6 = append(v6, -1)
        default:
            if (65535 >= 2147483647) {
                var v20 uint32 = 458
                _ = v20
            }
        }
        if (true || (!(65535 != uint(v18)))) {
            v18 = (v18 + v18)
            v6 = append(v6, 128)
            for fi21 := 0; fi21 < 19; fi21++ {
                if true {
                    _ = len(v6)
                }
                for fi22 := 0; fi22 < 9; fi22++ {
                    v6 = append(v6, fi21)
                    var v23 uint32 = 4294967295
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        }
    }
    _ = v18
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 uint16, p2 uint8) string {
    p2 = 208
    p1 += (p1 + p1)
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS3{F0: genFn0()}
        var v26 uint16 = (p1 / ((32767) | 1))
        switch int(v26) {
        case 127:
            p0 = p0
        default:
            p2 += 2
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    return ("" + "\n\t")
}

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    var v4 float32 = 3.14159
    var v5 float64 = (-2.25 / -2.25)
    _ = len(v1)
    if false {
        fmt.Println(v4, v4, v5)
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    } else {
        if true {
            if ((v3 * v3) > (v3 << ((4294967295) & 63))) {
                switch -9223372036854775808 {
                case 312:
                    switch 416 {
                    case 2147483647:
                        if false {
                            v27 := GenS2{F0: 255, F1: 1}
                            var v28 int = 128
                            switch "foo bar" {
                            case "a":
                                v4 = v4
                            case "hello":
                                fmt.Println(v4, v3)
                            case "\n\t":
                                var v29 uint32 = 2
                                _ = v29
                            default:
                                fmt.Println(v5, v3, v4)
                            }
                            _ = v27
                            _ = v28
                        }
                    case 32767:
                        fmt.Println(v3, v3)
                    case -5:
                        for fi30 := 0; fi30 < 20; fi30++ {
                            for fi31 := 0; fi31 < 17; fi31++ {
                                delete(v2, 42)
                                if len(v1) > 0 {
                                    v1[0] = -1
                                }
                                var v32 int = (fi31 >> ((640) & 63))
                                v33 := GenS1{F0: 735}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            v34 := GenS4{F0: 138}
                            v3 += 1
                            _ = fi30
                            _ = v34
                        }
                    }
                case 9223372036854775807:
                    fmt.Println(v4, v3)
                case -501:
                    for fi35 := 0; fi35 < 20; fi35++ {
                        if true {
                            v36 := v2[2147483647]
                            delete(v2, 42)
                            _ = v36
                        } else {
                            v3 -= -1
                            if false {
                                for fi37 := 0; fi37 < 11; fi37++ {
                                    _ = len(v1)
                                    v5 = v5
                                    if (!(v5 <= -2.25)) {
                                        if true {
                                            if false {
                                                v5 += v5
                                                switch ("hello" + "a") {
                                                case "hello":
                                                    v4 = (float32(v5) / float32(v5))
                                                }
                                                delete(v2, 42)
                                                delete(v2, 1)
                                            } else {
                                                var v38 int = -240
                                                var v39 int16 = 32767
                                                _ = v38
                                                _ = v39
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = (fi37 % ((fi35) | 1))
                                            }
                                            delete(v2, 2)
                                        }
                                        v4 = v4
                                    } else {
                                        fmt.Println(fi35, fi37, v3)
                                        var v40 uint16 = 0
                                        _ = v40
                                    }
                                    _ = fi37
                                }
                                if true {
                                    var v41 uint = 18446744073709551615
                                    var v42 int = (fi35 + (fi35 << ((v41) & 63)))
                                    var v43 uint8 = 127
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                                v44 := GenS5{F0: 360}
                                _ = v44
                            }
                        }
                        _ = fi35
                    }
                default:
                    if true {
                        for fi45 := 0; fi45 < 12; fi45++ {
                            for fi46 := 0; fi46 < 12; fi46++ {
                                var v47 uint8 = 0
                                var v48 float32 = float32(v5)
                                v49 := GenS6{F0: 0}
                                v5 = v5
                                switch -2147483648 {
                                case 127:
                                    v3 += v3
                                case -2147483648:
                                    var v50 uint16 = 935
                                    _ = v50
                                }
                                _ = fi46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            fmt.Println(v5)
                            v51 := GenS7{F0: 0}
                            _ = fi45
                            _ = v51
                        }
                        v5 = v5
                    } else {
                        var v52 uint8 = 213
                        v3 = (v3 << ((450) & 63))
                        v53 := GenS8{F0: 2, F1: 42}
                        _ = v52
                        _ = v53
                    }
                }
            } else {
                v3 += (v3 * -2147483648)
                var v54 bool = ((float64(float32(1e10)) * v5) <= v5)
                v3 += (v3 >> ((2) & 63))
                _ = v54
            }
            delete(v2, 952)
            v55 := GenS7{F0: 10}
            _ = v55
        } else {
            v4 += (v4 * v4)
            var v56 int32 = -822
            _ = v56
        }
    }
    if ((v4 / v4) > v4) {
        v5 = ((v5 + -2.25) * v5)
    } else {
        v3 -= 42
        for fi57 := 0; fi57 < 8; fi57++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi57
        }
    }
    v58 := GenS9{F0: 1, F1: 481, F2: 727}
    var v59 float64 = (1.5 * ((0.0 * ((v5 + (1.5 * (1.5 + (v5 - v5)))) / v5)) / -2.25))
    for fi60 := 0; fi60 < 5; fi60++ {
        v5 = (v59 + v59)
        _ = fi60
    }
    switch 832 {
    case 1:
        var v61 uint16 = 42
        _ = v61
    default:
        var v62 uint8 = 2
        _ = v62
    }
    _ = len(v1)
    var v63 int32 = -2147483648
    if true {
        fmt.Println(v3, v59, v59)
        v64 := GenS10{F0: ""}
        _ = v64
    } else {
        for fi65 := 0; fi65 < 10; fi65++ {
            if (true || true) {
                v4 += 3.14159
                v66 := v2[2147483647]
                _ = v66
            } else {
                _ = len(v2)
            }
            v3 += v3
            v59 += v5
            _ = fi65
        }
    }
    for fi67 := 0; fi67 < 7; fi67++ {
        v1 = append(v1, fi67)
        fmt.Println(v3)
        _ = fi67
    }
    v59 += ((-2.25 / v5) - (0.1 * v59))
    switch 32768 {
    case 0:
        if len(v1) > 0 {
            v1[0] = 0
        }
    case 401:
        if (!true) {
            fmt.Println(v4, v59)
        }
    }
    fmt.Println(v5)
    _ = len(v2)
    switch 9223372036854775807 {
    case -2147483648:
        if (9223372036854775807 <= 9223372036854775807) {
            switch 32768 {
            case 128:
                switch "\n\t" {
                case "hello":
                    v68 := GenS7{F0: 798}
                    _ = v68
                case "a":
                    fmt.Println(v59, v3, v63)
                case "foo bar":
                    v4 += v4
                }
            case 1:
                if (false || false) {
                    switch "\n\t" {
                    case "a":
                        for fi69 := 0; fi69 < 13; fi69++ {
                            var v70 uint8 = 42
                            fmt.Println(v63, v3)
                            var v71 uint64 = 2
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                    case "foo bar":
                        v1 = append(v1, -1)
                    default:
                        fmt.Println(v59, v63)
                    }
                }
            case 2147483647:
                v4 -= v4
            default:
                switch 128 {
                case -1:
                    _ = len(v1)
                case 916:
                    var v72 uint32 = 65535
                    _ = v72
                case -9223372036854775808:
                    if (2147483648 >= 2) {
                        fmt.Println(v5, v59, v59)
                        switch "\n\t" {
                        case "hello":
                            var v73 string = ("hello" + "\n\t")
                            _ = v73
                        default:
                            if true {
                                var v74 string = "\n\t"
                                switch 0 {
                                case 2147483647:
                                    v75 := GenS6{F0: 42}
                                    _ = v75
                                case -9223372036854775808:
                                    v63 += (v63 | v3)
                                }
                                _ = v74
                            } else {
                                v76 := v2[365]
                                var v77 uint16 = 2
                                _ = v76
                                _ = v77
                            }
                        }
                    } else {
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            var v78 uint32 = 0
                            _ = v78
                        }
                        for fi79 := 0; fi79 < 15; fi79++ {
                            var v80 uint = 826
                            for fi81 := 0; fi81 < 17; fi81++ {
                                switch fi79 {
                                case 397:
                                    v80 += (v80 - 1)
                                }
                                v4 = ((1e10 - v4) / -2.25)
                                v63 -= v63
                                _ = fi81
                            }
                            _ = fi79
                            _ = v80
                        }
                    }
                }
            }
        }
    case -618:
        var v82 int32 = 42
        _ = v82
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v63
}
