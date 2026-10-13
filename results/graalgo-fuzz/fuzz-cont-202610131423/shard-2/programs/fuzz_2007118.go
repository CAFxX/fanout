package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int16; F1 int8; F2 string }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 1
    v3 := []string{"\n\t", ""}
    v4 := []string{("hello" + "\n\t")}
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint8 = v2
            v2 -= (v2 * (v2 - 2))
            _ = fi7
            _ = v8
        }
        var v9 uint64 = 42
        fmt.Println(v2)
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v5[399] = 127
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := v5[0]
            fmt.Println(v12)
            v13 := GenS1{F0: 32767, F1: (0.0 > -2.25)}
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint = 0
            v16 := GenS1{F0: uint16(v15), F1: (!true)}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi10
    }
    v17 := GenS2{F0: 0}
    v18 := GenS2{F0: 1}
    if ((0.0 / 1.0) == (0.1 - (0.0 * ((1.5 - 1.5) * 0.1)))) {
        v2 += v2
        v5[1] = 32767
        delete(v5, 1)
    }
    if true {
        v19 := GenS1{F0: 32767, F1: (!((false && false) || ((0.0 * ((0.0 / 3.14159) - float64(1e10))) <= 1e10)))}
        _ = v19
    } else {
        if false {
            for fi20 := 0; fi20 < 5; fi20++ {
                v2 = (v2 * (v2 % (((v2 / ((27) | 1))) | 1)))
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS1{F0: 0, F1: false}
                    _ = fi21
                    _ = v22
                }
                if (1 > 1) {
                    switch ("hello" + "\n\t") {
                    case "":
                        v2 += (v2 + v2)
                    case "\n\t":
                        v2 -= v2
                    }
                } else {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if (false && true) {
                            _ = len(v4)
                            switch "hello" {
                            case "\n\t":
                                v2 = v2
                            default:
                                v2 = (v2 / ((v2) | 1))
                            }
                            v24 := GenS2{F0: 0}
                            _ = v24
                        } else {
                            var v25 int16 = -1
                            var v26 bool = ((18446744073709551615 <= 2) && false)
                            v27 := GenS2{F0: 0}
                            v25 = 1
                            v28 := GenS3{F0: 9223372036854775807, F1: 0}
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        fmt.Println(fi23, fi23, fi20)
                        v2 += (v2 % ((v2) | 1))
                        _ = fi23
                    }
                    if ((1 < 42) && ("\n\t" == ("\n\t" + "\n\t"))) {
                        var v29 string = "a"
                        _ = v29
                    } else {
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v2 = v2
                            _ = fi30
                        }
                        v31 := GenS4{F0: 32767, F1: 127, F2: "\n\t"}
                        _ = v31
                    }
                    var v32 int32 = -532
                    _ = v32
                }
                _ = fi20
            }
            fmt.Println(v2)
        }
    }
    v2 = 201
    v33 := GenS1{F0: 65535, F1: true}
    var v34 int8 = 1
    switch "\n\t" {
    case "a":
        var v35 int8 = v34
        _ = v35
    case "":
        var v36 int64 = -9223372036854775808
        _ = v36
    case "hello":
        v1 = append(v1, 65535)
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v34 -= 0
        v38 := GenS1{F0: 86, F1: (1.5 < 0.1)}
        _ = fi37
        _ = v38
    }
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
