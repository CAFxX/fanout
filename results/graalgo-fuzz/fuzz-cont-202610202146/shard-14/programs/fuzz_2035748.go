package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int64 }
func genFn0() int8 {
    if false {
        var v5 uint64 = 267
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: 32767, F2: 0}
    v7 := GenS1{F0: 4294967295, F1: 32767, F2: 223}
    var v8 uint32 = 2
    if (-288 <= 0) {
        v9 := GenS1{F0: 0, F1: 2, F2: -2147483648}
        var v10 uint32 = v8
        _ = v10
        _ = v9
    } else {
        v11 := []int{1, 32768, -9223372036854775808}
        fmt.Println(v8)
        fmt.Println(v8)
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    return -53
}

func main() {
    var v1 int8 = -20
    var v2 bool = ((42 > -1) && true)
    var v3 string = ("a" + "")
    var v4 int32 = 0
    v12 := []int{889, -2147483648, 2147483647}
    if (0.1 < 0.1) {
        v1 = v1
    } else {
        switch "\n\t" {
        case "a":
            var v13 uint16 = 0
            _ = v13
        }
        switch 32768 {
        case 464:
            switch 4294967295 {
            case -2147483648:
                for fi14 := 0; fi14 < 13; fi14++ {
                    if (42 < 1) {
                        v15 := GenS1{F0: 1, F1: 42, F2: -2147483648}
                        fmt.Println(v1)
                        _ = v15
                    } else {
                        v16 := GenS1{F0: 2, F1: 255, F2: -339}
                        v17 := GenS1{F0: 2, F1: 32767, F2: v4}
                        var v18 uint16 = 169
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS2{F0: v2, F1: -477}
                    for fi20 := 0; fi20 < 13; fi20++ {
                        if v2 {
                            switch v3 {
                            case "":
                                switch fi20 {
                                case -1:
                                    if ((v1 % ((1) | 1)) == v1) {
                                        switch v3 {
                                        case "\n\t":
                                            switch fi14 {
                                            case 0:
                                                v4 += v4
                                            default:
                                                if len(v12) > 0 {
                                                    v12[0] = 4294967295
                                                }
                                            }
                                        case "hello":
                                            v3 = v3
                                        case "a":
                                            fmt.Println(v1, fi20)
                                        default:
                                            fmt.Println(fi20, fi20)
                                        }
                                    } else {
                                        v2 = (((v3 + ("" + v3)) + "a") >= v3)
                                        fmt.Println(v3)
                                    }
                                case 9223372036854775807:
                                    var v21 uint = 255
                                    _ = v21
                                }
                            case "\n\t":
                                v22 := GenS3{F0: -1, F1: 0, F2: -149}
                                _ = v22
                            default:
                                var v23 uint16 = 852
                                _ = v23
                            }
                        } else {
                            _ = len(v12)
                        }
                        _ = fi20
                    }
                    _ = fi14
                    _ = v19
                }
            case 272:
                var v24 int32 = v4
                _ = v24
            case -852:
                var v25 uint64 = 4294967295
                _ = v25
            }
        default:
            v26 := GenS1{F0: 40, F1: 42, F2: 42}
            _ = v26
        }
    }
    v27 := GenS4{F0: 0}
    v28 := GenS5{F0: (v4 << ((18446744073709551615) & 63))}
    v1 -= (v1 + v1)
    if (411 <= 32767) {
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS6{F0: 9223372036854775807, F1: -9223372036854775808}
            fmt.Println(v4)
            _ = fi29
            _ = v30
        }
        fmt.Println(v3, v4)
        v12 = append(v12, 2147483647)
    }
    v3 = ("\n\t" + v3)
    if v2 {
        for fi31 := 0; fi31 < 12; fi31++ {
            if len(v12) > 0 {
                v12[0] = (^(fi31))
            }
            _ = fi31
        }
    }
    switch "foo bar" {
    case "\n\t":
        switch 0 {
        case 128:
            v2 = (46 > 162)
        default:
            if true {
                v32 := GenS2{F0: ((1e10 * 0.1) <= (-2.25 / (1.5 + 3.14159))), F1: -1}
                _ = v32
            }
        }
    }
    var v33 uint = 197
    var v34 float32 = (-2.25 / (-2.25 + -2.25))
    if true {
        fmt.Println(v1, v4, v4)
        if (!true) {
            var v35 float64 = ((0.0 * 1e10) * ((0.0 + 0.1) + (1e10 - 1e10)))
            _ = v35
        }
        fmt.Println(v4)
    } else {
        var v36 float32 = v34
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        v33 = v33
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
