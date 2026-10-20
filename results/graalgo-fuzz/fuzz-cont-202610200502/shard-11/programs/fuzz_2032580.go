package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS4 struct { F0 string; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 uint8; F1 string; F2 bool }
func genFn0() uint {
    var v5 int16 = 131
    var v6 int8 = 1
    if false {
        v7 := GenS1{F0: 51, F1: v5}
        _ = v7
    } else {
        var v8 uint32 = 4294967295
        _ = v8
    }
    var v9 int = 4294967295
    _ = v5
    _ = v6
    _ = v9
    return 42
}
func genFn1(p0 uint64, p1 string) uint {
    v10 := GenS1{F0: 18446744073709551615, F1: -1}
    var v11 int32 = -1
    var v12 int32 = v11
    v13 := []int{}
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: genFn0(), F1: 1000}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint = 824
    var v2 uint = 4294967295
    var v3 uint8 = uint8(v1)
    var v4 uint8 = v3
    for fi16 := 0; fi16 < 5; fi16++ {
        fmt.Println(v3)
        _ = fi16
    }
    var v17 string = ("foo bar" + "hello")
    v18 := GenS2{F0: 92, F1: 58}
    if true {
        switch "a" {
        case "\n\t":
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 uint32 = 706
                var v21 int16 = 32767
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        if true {
            var v22 uint32 = 490
            _ = v22
        }
        v23 := []int{-9223372036854775808, 42}
        _ = v23
    } else {
        v17 = v17
        var v24 float64 = ((0.0 * 1.5) * -2.25)
        _ = v24
    }
    v1 -= (v1 << ((65535) & 63))
    var v25 uint32 = 65535
    v4 -= v4
    var v26 int8 = 42
    v27 := GenS3{F0: 1, F1: 331, F2: 0}
    v1 = v2
    if (-1 <= -1) {
        switch 1 {
        case 42:
            var v28 uint8 = (^(v3))
            _ = v28
        case 138:
            for fi29 := 0; fi29 < 1; fi29++ {
                switch fi29 {
                case 712:
                    switch 127 {
                    case -454:
                        if ((860 <= 42) || ((v2 % ((18446744073709551615) | 1)) != v1)) {
                            var v30 int16 = -32768
                            _ = v30
                        }
                    case 250:
                        v31 := []int{(fi29 / ((127) | 1)), (fi29 | int(v2))}
                        _ = v31
                    case -9223372036854775808:
                        for fi32 := 0; fi32 < 18; fi32++ {
                            v33 := GenS1{F0: (^(v1)), F1: -32768}
                            fmt.Println(v1, v26)
                            v34 := GenS2{F0: 753, F1: -1}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    default:
                        switch fi29 {
                        case 2147483647:
                            v35 := []int{fi29, fi29}
                            _ = v35
                        default:
                            for fi36 := 0; fi36 < 19; fi36++ {
                                if true {
                                    v2 -= v2
                                    if ((fi29 / ((fi29) | 1)) >= 1) {
                                        var v37 int = 0
                                        _ = v37
                                    } else {
                                        var v38 int16 = -32768
                                        v2 += v1
                                        fmt.Println(v2, fi29, v25)
                                        _ = v38
                                    }
                                    fmt.Println(v17, fi36)
                                } else {
                                    if (!((1 == (v2 / ((v2) | 1))) || false)) {
                                        v1 = v2
                                        v17 = ("hello" + v17)
                                    }
                                    var v39 int = fi36
                                    v4 = (v4 + (v4 / (((^(v4))) | 1)))
                                    fmt.Println(fi29, v39, v17)
                                    _ = v39
                                }
                                fmt.Println(v25)
                                _ = fi36
                            }
                        }
                    }
                case 128:
                    v2 -= v2
                case 9223372036854775807:
                    v3 += v3
                default:
                    v40 := []int{}
                    _ = v40
                }
                if true {
                    var v41 uint64 = 575
                    if false {
                        if (0.1 > (1e10 / float64((-2.25 / 1.5)))) {
                            for fi42 := 0; fi42 < 15; fi42++ {
                                var v43 bool = true
                                _ = fi42
                                _ = v43
                            }
                            switch fi29 {
                            case 32767:
                                v44 := []int{-9223372036854775808, fi29, fi29, fi29}
                                _ = v44
                            case 32768:
                                for fi45 := 0; fi45 < 20; fi45++ {
                                    if (!false) {
                                        fmt.Println(v2)
                                        v46 := GenS1{F0: 42, F1: 0}
                                        fmt.Println(fi29, v26, v25)
                                        if (v3 < v3) {
                                            v47 := GenS1{F0: (v1 - (v2 % ((18446744073709551615) | 1))), F1: 32767}
                                            if ((v41 % ((v41) | 1)) < v41) {
                                                switch v17 {
                                                case "\n\t":
                                                    var v48 int32 = 2147483647
                                                    _ = v48
                                                }
                                            }
                                            var v49 float32 = 3.14159
                                            _ = v47
                                            _ = v49
                                        } else {
                                            v2 = v1
                                            if (v41 >= 0) {
                                                var v50 int8 = 0
                                                _ = v50
                                            }
                                            v51 := GenS3{F0: (v41 & (v41 | 42)), F1: (v41 % (((v41 >> ((v2) & 63))) | 1)), F2: v41}
                                            _ = v51
                                        }
                                        if (-280 > -109) {
                                            switch 709 {
                                            case 129:
                                                v52 := GenS4{F0: v17, F1: (v3 + (v3 << ((genFn0()) & 63))), F2: v25}
                                                _ = v52
                                            case 32767:
                                                v53 := GenS5{F0: (v3 & (^(v3))), F1: -2147483648}
                                                _ = v53
                                            }
                                            v1 = genFn0()
                                            var v54 uint = 65535
                                            _ = v54
                                        }
                                        fmt.Println(fi45, v25)
                                        var v55 int16 = -1
                                        _ = v46
                                        _ = v55
                                    }
                                    _ = fi45
                                }
                            case 2147483647:
                                fmt.Println(v26, v4)
                            }
                            for fi56 := 0; fi56 < 13; fi56++ {
                                v57 := GenS6{F0: v3, F1: "a", F2: false}
                                _ = fi56
                                _ = v57
                            }
                        }
                    }
                    _ = v41
                }
                v2 -= uint(v25)
                _ = fi29
            }
        case -2147483648:
            fmt.Println(v4)
        }
    }
    v25 = (^(v25))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
