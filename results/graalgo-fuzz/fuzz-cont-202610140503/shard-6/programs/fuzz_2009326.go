package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS4 struct { F0 string; F1 int32; F2 uint64 }

func main() {
    var v1 uint = 42
    var v2 int16 = 32767
    v3 := make(map[string]int)
    var v4 uint64 = 0
    fmt.Println(v2, v2, v1)
    v5 := GenS1{F0: 0, F1: -9223372036854775808}
    var v6 uint32 = 1
    v7 := []int{}
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 = append(v7, fi8)
        switch (fi8 % ((-2147483648) | 1)) {
        case 127:
            var v9 int32 = 100
            _ = v9
        default:
            switch ("" + ("\n\t" + (("" + (("a" + "") + "hello")) + ""))) {
            case "\n\t":
                for fi10 := 0; fi10 < 15; fi10++ {
                    v6 += (v6 | (^(v6)))
                    if (2 != 61) {
                        var v11 uint8 = 1
                        v12 := GenS1{F0: 2147483647, F1: 0}
                        _ = v11
                        _ = v12
                    } else {
                        var v13 bool = (-821 == -695)
                        var v14 uint = (v1 >> ((v1) & 63))
                        fmt.Println(v14)
                        if v13 {
                            var v15 int32 = 42
                            v16 := GenS1{F0: v15, F1: 1}
                            v17 := GenS1{F0: (v15 % (((v15 << ((v14) & 63))) | 1)), F1: -601}
                            v18 := GenS1{F0: (^(v15)), F1: 0}
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        _ = v13
                        _ = v14
                    }
                    var v19 uint16 = 526
                    v19 = v19
                    fmt.Println(v19, v1, fi10)
                    _ = fi10
                    _ = v19
                }
            }
        }
        _ = fi8
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v7 = append(v7, -224)
        v4 = 0
        switch fi20 {
        case 128:
            switch 421 {
            case -91:
                if (1e10 <= (1.5 / float64(3.14159))) {
                    var v21 float64 = float64((3.14159 + 1.5))
                    _ = v21
                } else {
                    if false {
                        for fi22 := 0; fi22 < 7; fi22++ {
                            if (true && true) {
                                if (-80 < int8(v2)) {
                                    var v23 float32 = (0.0 + float32((0.0 / 1e10)))
                                    v4 = (v4 & 18446744073709551615)
                                    _ = v23
                                }
                                var v24 uint8 = 251
                                _ = v24
                            }
                            if (!true) {
                                v25 := GenS1{F0: -260, F1: -1}
                                _ = v25
                            } else {
                                v6 -= 2147483647
                                var v26 uint16 = 0
                                v27 := GenS2{F0: -128}
                                if false {
                                    if (!(1e10 > -2.25)) {
                                        v28 := GenS2{F0: 25}
                                        delete(v3, ("a" + ("\n\t" + "")))
                                        if ((v2 << ((2) & 63)) >= v2) {
                                            var v29 string = ""
                                            _ = v29
                                        }
                                        _ = v28
                                    }
                                    fmt.Println(v26, v2, v26)
                                } else {
                                    var v30 uint = (v1 | v1)
                                    var v31 uint8 = 1
                                    v26 -= 255
                                    _ = v30
                                    _ = v31
                                }
                                v26 = (v26 << (((^(v1))) & 63))
                                _ = v26
                                _ = v27
                            }
                            _ = fi22
                        }
                        v32 := GenS2{F0: 127}
                        _ = v32
                    }
                }
            case -9223372036854775808:
                for fi33 := 0; fi33 < 16; fi33++ {
                    v34 := GenS2{F0: 1}
                    _ = fi33
                    _ = v34
                }
            case 128:
                v4 -= v4
            default:
                for fi35 := 0; fi35 < 4; fi35++ {
                    v2 = (v2 * (v2 % (((^(v2))) | 1)))
                    var v36 uint64 = v4
                    _ = fi35
                    _ = v36
                }
            }
        default:
            v6 += (v6 / ((v6) | 1))
        }
        _ = fi20
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        for fi38 := 0; fi38 < 14; fi38++ {
            v39 := GenS3{F0: uint16(v1), F1: (v1 & (^(v1))), F2: (v6 * v6)}
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            if ("\n\t" >= "foo bar") {
                if (!true) {
                    var v41 uint64 = v4
                    _ = v41
                } else {
                    var v42 int = (fi37 / (((fi37 % ((fi40) | 1))) | 1))
                    v3["hello"] = -686
                    var v43 float64 = 3.14159
                    _ = v42
                    _ = v43
                }
                var v44 int64 = -1
                _ = v44
            }
            var v45 uint64 = 2
            fmt.Println(fi40, v4)
            _ = fi40
            _ = v45
        }
        _ = fi37
    }
    var v46 uint = v1
    v47 := GenS4{F0: ((("a" + "") + "a") + ("\n\t" + ("a" + "a"))), F1: -2147483648, F2: v4}
    var v48 float64 = (3.14159 - (1.5 / (((0.0 - (float64((1.5 / 1.0)) + (1.5 / 3.14159))) * 0.1) + 1e10)))
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
