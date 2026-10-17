package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int64; F2 int64 }
func genFn0() int {
    v7 := []int{}
    var v8 float32 = 3.14159
    if ((0 >= 127) || (!(!false))) {
        if true {
            var v9 uint64 = 0
            for fi10 := 0; fi10 < 17; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 float64 = 1.5
                    _ = fi11
                    _ = v12
                }
                v13 := GenS1{F0: 0, F1: 0}
                _ = fi10
                _ = v13
            }
            switch "\n\t" {
            case "hello":
                if true {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        for fi15 := 0; fi15 < 9; fi15++ {
                            switch "\n\t" {
                            case "":
                                if (!true) {
                                    if true {
                                        v16 := GenS1{F0: 127, F1: 226}
                                        switch fi14 {
                                        case -2147483648:
                                            v17 := GenS2{F0: -892}
                                            var v18 uint16 = 607
                                            var v19 bool = (507 != 18446744073709551615)
                                            _ = v17
                                            _ = v18
                                            _ = v19
                                        case -9223372036854775808:
                                            fmt.Println(v9)
                                        }
                                        v8 = ((0.1 + v8) / (1.5 * v8))
                                        _ = len(v7)
                                        v7 = append(v7, fi15)
                                        _ = v16
                                    } else {
                                        v9 = v9
                                        if (!true) {
                                            switch (fi14 % ((42) | 1)) {
                                            case 127:
                                                if false {
                                                    if true {
                                                        v9 = v9
                                                        var v20 uint8 = 127
                                                        _ = v20
                                                    } else {
                                                        fmt.Println(v8, v8, fi15)
                                                    }
                                                    v9 += v9
                                                } else {
                                                    if len(v7) > 0 {
                                                        v7[0] = fi15
                                                    }
                                                }
                                            case 9223372036854775807:
                                                v7 = append(v7, fi14)
                                            case 1:
                                                v9 -= v9
                                            default:
                                                var v21 int32 = 42
                                                _ = v21
                                            }
                                        } else {
                                            if (!(289 <= 32767)) {
                                                v8 = 1e10
                                                v8 -= (v8 * (v8 * (float32(1.5) - (v8 + (-2.25 - float32(1.5))))))
                                                if false {
                                                    var v22 int16 = -1
                                                    if (110 != 2) {
                                                        v8 = 1e10
                                                    } else {
                                                        v23 := GenS3{F0: 9223372036854775807, F1: 1}
                                                        switch fi15 {
                                                        case 4294967295:
                                                            v24 := GenS2{F0: v22}
                                                            _ = len(v7)
                                                            _ = v24
                                                        case 429:
                                                            var v25 int32 = 2147483647
                                                            _ = v25
                                                        default:
                                                            if (true || false) {
                                                                var v26 int = fi15
                                                                var v27 bool = false
                                                                _ = v26
                                                                _ = v27
                                                            }
                                                        }
                                                        v7 = append(v7, fi14)
                                                        _ = v23
                                                    }
                                                    _ = v22
                                                } else {
                                                    fmt.Println(fi15, fi14)
                                                }
                                                fmt.Println(fi14, v9)
                                            } else {
                                                v9 = 365
                                            }
                                            _ = len(v7)
                                        }
                                        if (0 > 0) {
                                            v8 -= 0.1
                                        }
                                    }
                                    v9 = v9
                                    switch (fi14 * (fi15 ^ (fi14 / ((-116) | 1)))) {
                                    case 1:
                                        v28 := GenS1{F0: 2, F1: 1}
                                        _ = v28
                                    }
                                }
                            default:
                                v9 -= v9
                            }
                            v9 = (^(v9))
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi15
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        _ = fi14
                    }
                    fmt.Println(v8, v9)
                } else {
                    fmt.Println(v9)
                    var v29 int32 = 646
                    var v30 uint = 0
                    _ = v29
                    _ = v30
                }
            case "":
                switch "" {
                case "":
                    var v31 int8 = -1
                    _ = v31
                case "hello":
                    for fi32 := 0; fi32 < 16; fi32++ {
                        for fi33 := 0; fi33 < 3; fi33++ {
                            var v34 int8 = 97
                            v35 := GenS1{F0: 26, F1: 0}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        v9 = (v9 + 42)
                        var v36 float64 = 1.5
                        _ = fi32
                        _ = v36
                    }
                default:
                    _ = len(v7)
                }
            case "foo bar":
                fmt.Println(v8)
            default:
                var v37 uint32 = 65535
                _ = v37
            }
            _ = v9
        } else {
            v8 += ((v8 + 3.14159) + (3.14159 + v8))
        }
        v8 = v8
    } else {
        if true {
            fmt.Println(v8)
            v38 := GenS2{F0: 953}
            _ = v38
        }
    }
    _ = v7
    _ = v8
    return 32768
}

func main() {
    var v1 float64 = (((0.1 + (0.0 + 0.1)) - 1e10) - (1.5 - 0.1))
    v2 := []int{}
    var v3 uint = 0
    var v4 int8 = int8(v3)
    v5 := make(map[uint32]uint32)
    var v6 float64 = 0.1
    switch "a" {
    case "a":
        var v39 int64 = 9223372036854775807
        _ = v39
    case "foo bar":
        for fi40 := 0; fi40 < 17; fi40++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v41 int = fi40
            _ = fi40
            _ = v41
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        var v43 bool = (1 > 42)
        if (!(uint8(v3) == 233)) {
            v1 += (v1 * v6)
            v43 = v43
            if ((^(fi42)) >= genFn0()) {
                v44 := GenS4{F0: -128, F1: 65535}
                fmt.Println(v4)
                var v45 uint32 = 2
                _ = v44
                _ = v45
            } else {
                v5[0] = 1
                var v46 bool = (0.1 != (-2.25 - ((-2.25 / 3.14159) - float32(((0.0 * ((float64(3.14159) - 0.1) * 0.1)) + 3.14159)))))
                _ = v46
            }
        } else {
            switch "foo bar" {
            case "":
                _ = len(v2)
            case "\n\t":
                for fi47 := 0; fi47 < 1; fi47++ {
                    v48 := GenS5{F0: -667, F1: 0, F2: 1}
                    _ = fi47
                    _ = v48
                }
            case "a":
                var v49 uint8 = 66
                _ = v49
            default:
                if (4294967295 <= 42) {
                    v50 := GenS6{F0: 127}
                    var v51 int64 = -9223372036854775808
                    _ = v50
                    _ = v51
                }
            }
            switch "\n\t" {
            case "foo bar":
                v52 := GenS5{F0: 42, F1: 0, F2: 890}
                _ = v52
            case "hello":
                v43 = true
            }
        }
        _ = fi42
        _ = v43
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v5[0] = 104
    var v53 uint = v3
    v54 := GenS5{F0: 1, F1: 42, F2: 629}
    if true {
        if false {
            v53 -= v53
        }
    }
    var v55 int32 = 488
    if false {
        delete(v5, 36)
        var v56 float32 = 1.5
        _ = v56
    } else {
        switch "foo bar" {
        case "":
            v5[2147483648] = 4294967295
        case "foo bar":
            v57 := GenS7{F0: 0, F1: -9223372036854775808, F2: 560}
            _ = v57
        default:
            v58 := GenS4{F0: (^(v4)), F1: 2147483647}
            _ = v58
        }
        if len(v2) > 0 {
            v2[0] = genFn0()
        }
    }
    v59 := GenS4{F0: v4, F1: uint32(v53)}
    var v60 float64 = v1
    fmt.Println(v1, v1, v1)
    v61 := GenS1{F0: 1, F1: 4294967295}
    var v62 uint = v53
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
