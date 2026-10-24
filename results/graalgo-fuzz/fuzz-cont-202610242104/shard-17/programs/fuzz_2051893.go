package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 bool; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool; F1 bool; F2 uint32 }
type GenS5 struct { F0 int32 }
func genFn0() string {
    if (false || false) {
        var v6 bool = false
        _ = v6
    } else {
        v7 := []int{}
        _ = v7
    }
    if false {
        var v8 int8 = -14
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if (727 >= 2) {
            v10 := []int{fi9}
            _ = v10
        } else {
            var v11 uint8 = 198
            _ = v11
        }
        v12 := GenS1{F0: 127, F1: (fi9 / ((fi9) | 1)), F2: 1}
        _ = fi9
        _ = v12
    }
    var v13 string = "hello"
    if (0 <= 0) {
        v14 := GenS1{F0: -1, F1: 0, F2: 42}
        switch 32768 {
        case 32767:
            for fi15 := 0; fi15 < 12; fi15++ {
                v13 = v13
                var v16 uint8 = 127
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 float64 = 1.5
                    switch (fi15 << ((18446744073709551615) & 63)) {
                    case 869:
                        var v19 string = (v13 + v13)
                        _ = v19
                    }
                    v13 = v13
                    if (!false) {
                        var v20 uint32 = 4294967295
                        v20 = v20
                        v16 = (^(v16))
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        case 127:
            var v21 int8 = 42
            _ = v21
        case -375:
            switch v13 {
            case "a":
                v13 = (((v13 + "a") + (v13 + (v13 + "\n\t"))) + "")
            case "\n\t":
                for fi22 := 0; fi22 < 17; fi22++ {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        var v24 uint = 570
                        var v25 uint = (^(v24))
                        if true {
                            if (18446744073709551615 >= v24) {
                                var v26 float32 = 1e10
                                if ("hello" <= v13) {
                                    v27 := GenS2{F0: 1, F1: 2147483647, F2: (v24 << ((255) & 63))}
                                    if (2147483647 >= 0) {
                                        fmt.Println(fi22, v13, v25)
                                        var v28 uint8 = 23
                                        _ = v28
                                    }
                                    v29 := GenS3{F0: false, F1: 42, F2: 42}
                                    _ = v27
                                    _ = v29
                                } else {
                                    v30 := []int{-541}
                                    _ = v30
                                }
                                v31 := GenS2{F0: 255, F1: 889, F2: 42}
                                _ = v26
                                _ = v31
                            } else {
                                var v32 string = v13
                                var v33 uint8 = 127
                                _ = v32
                                _ = v33
                            }
                            var v34 int32 = -1
                            _ = v34
                        }
                        v35 := GenS4{F0: false, F1: false, F2: 2147483648}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v35
                    }
                    fmt.Println(fi22)
                    v13 = v13
                    _ = fi22
                }
            default:
                v13 = ("a" + v13)
            }
        default:
            v13 = "a"
        }
        _ = v14
    } else {
        var v36 int32 = -220
        _ = v36
    }
    _ = v13
    return (v13 + (v13 + v13))
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    var v3 uint64 = v1
    var v4 int32 = -710
    var v5 uint8 = 128
    v5 += (v5 % (((v5 >> ((18446744073709551615) & 63))) | 1))
    for fi37 := 0; fi37 < 13; fi37++ {
        if (!(fi37 > (fi37 | (fi37 - fi37)))) {
            v2[fi37] = 749
            var v38 float32 = ((1e10 * 0.1) + 0.1)
            _ = v38
        }
        _ = fi37
    }
    if false {
        if false {
            if true {
                var v39 string = ("" + ("foo bar" + (((("a" + "a") + "foo bar") + ((((("hello" + "foo bar") + "\n\t") + ("hello" + "a")) + (("\n\t" + "") + ("a" + "hello"))) + "\n\t")) + "\n\t")))
                _ = v39
            } else {
                fmt.Println(v1, v4)
            }
        } else {
            for fi40 := 0; fi40 < 7; fi40++ {
                v5 = (v5 - (v5 >> ((2) & 63)))
                v41 := GenS5{F0: v4}
                _ = fi40
                _ = v41
            }
        }
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        var v43 int = (fi42 & (^(fi42)))
        _ = fi42
        _ = v43
    }
    if true {
        fmt.Println(v1, v3, v4)
        fmt.Println(v1, v3)
        var v44 uint = 18446744073709551615
        _ = v44
    } else {
        v1 = v3
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 float32 = 0.0
        _ = fi45
        _ = v46
    }
    v47 := []int{}
    switch (("a" + ("" + "foo bar")) + "\n\t") {
    case "hello":
        delete(v2, 0)
    case "\n\t":
        if (-1 != 32767) {
            delete(v2, 2147483647)
            v47 = append(v47, -9223372036854775808)
        }
    case "":
        var v48 int64 = 42
        _ = v48
    default:
        for fi49 := 0; fi49 < 1; fi49++ {
            switch (fi49 | fi49) {
            case 0:
                if true {
                    var v50 string = "foo bar"
                    _ = v50
                } else {
                    fmt.Println(v3, v4, v3)
                    fmt.Println(v1, v1, v4)
                    var v51 float32 = 1.5
                    _ = v51
                }
            }
            v4 -= v4
            _ = fi49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
