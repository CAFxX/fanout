package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint16, p1 uint8, p2 string) uint {
    v7 := []int{42, 127}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int16 = -1
    v2 := []uint32{2147483647}
    v3 := []uint32{4294967295}
    v4 := []uint32{2147483648, 2147483648, 1, 4294967295}
    v5 := make(map[string]bool)
    var v6 uint32 = 243
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(fi8)
        for fi9 := 0; fi9 < 14; fi9++ {
            v6 -= 0
            var v10 uint32 = v6
            _ = fi9
            _ = v10
        }
        if ((fi8 ^ (fi8 % ((-2147483648) | 1))) <= fi8) {
            var v11 uint8 = 205
            var v12 int = (fi8 / ((32767) | 1))
            switch -646 {
            case -9223372036854775808:
                v11 += v11
            case 32768:
                for fi13 := 0; fi13 < 7; fi13++ {
                    v11 = 2
                    if false {
                        v1 += 42
                        fmt.Println(v12, v1)
                    } else {
                        v12 = fi13
                        var v14 float64 = 1e10
                        _ = v14
                    }
                    _ = fi13
                }
            }
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    var v15 string = "foo bar"
    v15 = "\n\t"
    v5["a"] = (!(1 <= 0))
    v16 := GenS1{F0: (v15 + "\n\t")}
    v17 := GenS2{F0: 42, F1: 90}
    switch (v15 + "foo bar") {
    case "\n\t":
        var v18 bool = true
        _ = v18
    case "a":
        if (195 >= 255) {
            v6 = v6
            var v19 uint8 = 220
            switch (((v15 + "") + "") + v15) {
            case "":
                v6 += (v6 << ((uint(v19)) & 63))
            case "hello":
                v19 += 42
            case "foo bar":
                var v20 float32 = 1.5
                _ = v20
            default:
                for fi21 := 0; fi21 < 16; fi21++ {
                    v22 := GenS2{F0: 1, F1: -101}
                    if true {
                        var v23 int = fi21
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v6 += (v6 + 804)
                            v4 = append(v4, v6)
                            var v25 uint = 4294967295
                            switch -364 {
                            case -86:
                                var v26 int64 = -9223372036854775808
                                _ = v26
                            case -317:
                                v6 = 0
                            case -2147483648:
                                _ = len(v4)
                            default:
                                var v27 string = v15
                                _ = v27
                            }
                            var v28 string = v15
                            var v29 int32 = int32(v19)
                            fmt.Println(v29, v25, fi21)
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = (v6 % (((v6 >> ((v25) & 63))) | 1))
                                }
                                var v30 string = ((v28 + "a") + v15)
                                _ = v30
                            } else {
                                v15 = ((v28 + v28) + (v28 + ((v15 + ("\n\t" + ("foo bar" + v28))) + (v28 + v15))))
                                v1 += v1
                            }
                            fmt.Println(v1, v25, v15)
                            _ = fi24
                            _ = v25
                            _ = v28
                            _ = v29
                        }
                        _ = v23
                    } else {
                        v15 = "a"
                    }
                    _ = fi21
                    _ = v22
                }
            }
            _ = v19
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                for fi32 := 0; fi32 < 6; fi32++ {
                    var v33 int16 = (v1 | v1)
                    _ = fi32
                    _ = v33
                }
                fmt.Println(fi31, v15, v15)
                switch "" {
                case "hello":
                    v6 += (^(v6))
                }
                _ = fi31
            }
        }
    case "foo bar":
        var v34 uint32 = v6
        _ = v34
    }
    var v35 int16 = 1
    var v36 uint64 = 18446744073709551615
    if (!true) {
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 int64 = -755
            _ = fi37
            _ = v38
        }
        switch -1 {
        case -9223372036854775808:
            for fi39 := 0; fi39 < 17; fi39++ {
                if false {
                    for fi40 := 0; fi40 < 15; fi40++ {
                        v15 = "hello"
                        _ = fi40
                    }
                    if (true && false) {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            if (v6 <= v6) {
                                fmt.Println(v15, fi41, v36)
                                var v42 int8 = 100
                                if false {
                                    _ = len(v5)
                                }
                                switch (v15 + (v15 + "a")) {
                                case "hello":
                                    var v43 int64 = 0
                                    v44 := GenS2{F0: v6, F1: (^(v42))}
                                    _ = v43
                                    _ = v44
                                default:
                                    switch v15 {
                                    case "":
                                        v15 = ((v15 + v15) + "")
                                    case "foo bar":
                                        switch 0 {
                                        case 4294967295:
                                            var v45 uint64 = v36
                                            _ = v45
                                        case 127:
                                            var v46 uint8 = 2
                                            _ = v46
                                        }
                                    }
                                }
                                v6 = v6
                                if true {
                                    var v47 float64 = 0.1
                                    _ = v47
                                } else {
                                    if false {
                                        v48 := GenS2{F0: v6, F1: v42}
                                        v15 = v15
                                        if (9223372036854775807 == 42) {
                                            if len(v4) > 0 {
                                                v4[0] = (v6 >> ((4294967295) & 63))
                                            }
                                            v4 = append(v4, 2)
                                            var v49 float32 = (-2.25 * 1e10)
                                            _ = v49
                                        }
                                        _ = v48
                                    } else {
                                        v50 := GenS2{F0: v6, F1: (^(v42))}
                                        v6 = v6
                                        _ = v50
                                    }
                                    if (true && (v36 < (v36 * 119))) {
                                        v35 = -861
                                    }
                                }
                                v5["a"] = false
                                _ = v42
                            }
                            fmt.Println(fi39)
                            if true {
                                if true {
                                    fmt.Println(v35, fi39, v36)
                                } else {
                                    switch ("a" + (v15 + "foo bar")) {
                                    case "a":
                                        var v51 uint64 = v36
                                        _ = v51
                                    }
                                    var v52 uint64 = (v36 >> ((2) & 63))
                                    v53 := GenS3{F0: 0, F1: (^(v6))}
                                    _ = v52
                                    _ = v53
                                }
                                if ((v6 % ((v6) | 1)) < (v6 * (^(v6)))) {
                                    var v54 bool = (v36 > (^(v36)))
                                    v55 := GenS4{F0: 9223372036854775807}
                                    switch fi41 {
                                    case 471:
                                        v54 = (v6 != 2147483647)
                                    default:
                                        var v56 int16 = v1
                                        _ = v56
                                    }
                                    _ = v54
                                    _ = v55
                                }
                            } else {
                                v57 := GenS4{F0: 359}
                                if true {
                                    v1 = v1
                                    v15 = v15
                                    fmt.Println(fi39, v6, v1)
                                } else {
                                    v5[v15] = (128 == 214)
                                    var v58 bool = (true && false)
                                    v59 := GenS5{F0: 1, F1: 940}
                                    v6 += (v6 >> ((148) & 63))
                                    v4 = append(v4, (v6 + (^(v6))))
                                    _ = v58
                                    _ = v59
                                }
                                _ = v57
                            }
                            v36 = (^(v36))
                            var v60 int32 = 42
                            v36 -= v36
                            switch "foo bar" {
                            case "":
                                var v61 string = "a"
                                _ = v61
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "a":
                                var v62 int16 = (v1 % ((42) | 1))
                                _ = v62
                            }
                            _ = fi41
                            _ = v60
                        }
                    } else {
                        v1 = v1
                    }
                } else {
                    _ = len(v5)
                }
                switch fi39 {
                case -50:
                    switch "foo bar" {
                    case "":
                        if (!(uint64(v6) < v36)) {
                            switch (v15 + ((("" + v15) + v15) + (v15 + v15))) {
                            case "hello":
                                fmt.Println(v36, v15, fi39)
                            case "\n\t":
                                var v63 bool = false
                                _ = v63
                            case "foo bar":
                                v1 = v1
                            }
                            v64 := GenS6{F0: v1, F1: "\n\t"}
                            _ = v64
                        }
                    }
                default:
                    delete(v5, v15)
                }
                _ = fi39
            }
        case 32767:
            v5[v15] = ((0.0 + 0.0) < float64(0.1))
        }
    } else {
        if ((-2147483648 == 649) && true) {
            if false {
                fmt.Println(v6)
                if false {
                    if false {
                        for fi65 := 0; fi65 < 10; fi65++ {
                            switch (fi65 / ((-976) | 1)) {
                            case 849:
                                if true {
                                    v66 := GenS7{F0: -32768}
                                    _ = v66
                                } else {
                                    if (!(v15 > "\n\t")) {
                                        var v67 float32 = 0.0
                                        v35 -= -617
                                        _ = v67
                                    } else {
                                        switch "hello" {
                                        case "foo bar":
                                            v15 = (v15 + v15)
                                        case "":
                                            for fi68 := 0; fi68 < 13; fi68++ {
                                                v1 -= v35
                                                var v69 uint32 = (v6 & (v6 / (((v6 >> ((genFn0(2, 255, v15)) & 63))) | 1)))
                                                var v70 float64 = 1e10
                                                _ = fi68
                                                _ = v69
                                                _ = v70
                                            }
                                        case "\n\t":
                                            var v71 uint64 = v36
                                            _ = v71
                                        }
                                        var v72 float64 = float64(1e10)
                                        v73 := GenS3{F0: 184, F1: v6}
                                        _ = v72
                                        _ = v73
                                    }
                                    v74 := GenS8{F0: true}
                                    _ = len(v2)
                                    _ = v74
                                }
                            case -2147483648:
                                v35 -= int16(fi65)
                            case 128:
                                var v75 float32 = 1.5
                                _ = v75
                            }
                            _ = fi65
                        }
                        var v76 uint64 = v36
                        v77 := GenS9{F0: (v76 & (v36 % ((v36) | 1))), F1: 2}
                        _ = v76
                        _ = v77
                    }
                    var v78 string = "\n\t"
                    _ = v78
                }
                for fi79 := 0; fi79 < 2; fi79++ {
                    v80 := GenS5{F0: 9223372036854775807, F1: 42}
                    _ = fi79
                    _ = v80
                }
            } else {
                switch 728 {
                case -2147483648:
                    var v81 string = (("foo bar" + ("\n\t" + v15)) + (v15 + ("\n\t" + v15)))
                    _ = v81
                default:
                    fmt.Println(v36)
                }
            }
            var v82 float32 = 3.14159
            var v83 float64 = 3.14159
            _ = v82
            _ = v83
        }
        v15 = v15
    }
    var v84 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v84
}
