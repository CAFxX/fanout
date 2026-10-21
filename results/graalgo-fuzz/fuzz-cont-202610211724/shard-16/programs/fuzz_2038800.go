package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0() int32 {
    v5 := []int{-9223372036854775808, 128, -19, 32768}
    if (true || true) {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 222, F1: 9223372036854775807, F2: 42}
            _ = fi6
            _ = v7
        }
        var v8 uint8 = 0
        var v9 uint8 = v8
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 127
        var v11 uint = 993
        _ = v10
        _ = v11
    }
    _ = v5
    return 0
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{44, 4294967295, 4294967295}
    var v3 uint = 532
    var v4 uint = (v3 - (v3 | (v3 ^ (^(v3)))))
    for fi12 := 0; fi12 < 2; fi12++ {
        if len(v2) > 0 {
            v2[0] = 539
        }
        var v13 uint64 = 42
        _ = fi12
        _ = v13
    }
    v3 -= 549
    _ = len(v2)
    switch ("\n\t" + "\n\t") {
    case "":
        _ = len(v2)
    case "hello":
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    var v14 bool = false
    v3 = 767
    fmt.Println(v3, v14)
    var v15 int = 127
    fmt.Println(v3, v4, v1)
    switch "\n\t" {
    case "foo bar":
        var v16 uint16 = 908
        _ = v16
    }
    var v17 int32 = -2147483648
    for fi18 := 0; fi18 < 18; fi18++ {
        v1 = 203
        fmt.Println(fi18)
        v19 := GenS1{F0: (v17 % ((v17) | 1)), F1: (v15 / (((v15 / ((int(v1)) | 1))) | 1)), F2: (^(v1))}
        _ = fi18
        _ = v19
    }
    switch "foo bar" {
    case "\n\t":
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 float32 = 1e10
            _ = fi20
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 14; fi22++ {
            if v14 {
                v15 += -9223372036854775808
            } else {
                v15 = (fi22 | 0)
            }
            var v23 uint32 = 82
            _ = fi22
            _ = v23
        }
    }
    if v14 {
        v2 = append(v2, 487)
        var v24 int64 = 235
        switch (v15 % ((v15) | 1)) {
        case -1:
            switch "" {
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v25 int32 = genFn0()
                _ = v25
            }
        default:
            if (("\n\t" < "foo bar") || v14) {
                v26 := GenS2{F0: (v24 % (((-(v24))) | 1)), F1: (((float64(-2.25) + float64(0.1)) <= 0.1) && v14)}
                if true {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        if true {
                            switch v15 {
                            case -516:
                                var v28 float32 = -2.25
                                _ = v28
                            }
                        } else {
                            fmt.Println(v1)
                            v29 := GenS2{F0: v24, F1: (!(false && v14))}
                            _ = v29
                        }
                        v1 += (v1 >> ((v4) & 63))
                        _ = fi27
                    }
                    if (!v14) {
                        v15 = (v15 << (((v3 - v4)) & 63))
                    }
                    if true {
                        v17 = (^(v17))
                        if false {
                            if (v17 > 1) {
                                v14 = v14
                            } else {
                                v3 = v4
                                v30 := GenS2{F0: (v24 * 42), F1: v14}
                                v31 := GenS1{F0: genFn0(), F1: (v15 & (v15 - (v15 << (((v3 << (((^(v4))) & 63))) & 63)))), F2: (v1 << ((395) & 63))}
                                _ = v30
                                _ = v31
                            }
                            v2 = append(v2, 758)
                        }
                        v14 = v14
                    }
                }
                var v32 float32 = -2.25
                _ = v26
                _ = v32
            } else {
                v17 += genFn0()
                switch 32767 {
                case 2147483647:
                    var v33 uint64 = 0
                    _ = v33
                case 977:
                    switch "\n\t" {
                    case "foo bar":
                        v24 = v24
                    case "":
                        switch "" {
                        case "foo bar":
                            fmt.Println(v17, v17)
                        case "\n\t":
                            v14 = (genFn0() <= (v17 * 168))
                        default:
                            switch "hello" {
                            case "\n\t":
                                v3 += 17
                            case "a":
                                switch v15 {
                                case 1:
                                    var v34 uint16 = 493
                                    _ = v34
                                case 9223372036854775807:
                                    switch "\n\t" {
                                    case "hello":
                                        var v35 float64 = float64(1.5)
                                        _ = v35
                                    case "":
                                        v36 := GenS2{F0: 857, F1: v14}
                                        _ = v36
                                    default:
                                        var v37 uint32 = 475
                                        _ = v37
                                    }
                                }
                            }
                        }
                    }
                default:
                    var v38 uint = 41
                    _ = v38
                }
            }
        }
        _ = v24
    } else {
        fmt.Println(v15)
        fmt.Println(v4, v15)
    }
    var v39 float64 = 1.5
    _ = len(v2)
    v40 := GenS3{F0: 585, F1: (v17 << ((v4) & 63))}
    for fi41 := 0; fi41 < 17; fi41++ {
        v42 := GenS2{F0: -1, F1: (-32768 <= -676)}
        if true {
            v2 = append(v2, 996)
            fmt.Println(fi41, v4, v4)
            v17 += v17
        }
        var v43 uint64 = 210
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
