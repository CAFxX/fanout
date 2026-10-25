package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 int; F2 int }
type GenS3 struct { F0 int64; F1 int16; F2 string }

func main() {
    var v1 int = 42
    v2 := []uint32{65535, 0, 943, 0}
    var v3 uint8 = 42
    _ = len(v2)
    if true {
        for fi4 := 0; fi4 < 9; fi4++ {
            var v5 string = ("foo bar" + ("foo bar" + ""))
            if true {
                if true {
                    var v6 uint16 = 645
                    for fi7 := 0; fi7 < 13; fi7++ {
                        v8 := GenS1{F0: -905, F1: 4294967295}
                        var v9 int = -1
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    v1 -= v1
                    _ = v6
                } else {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    if false {
                        v10 := GenS1{F0: 0, F1: 255}
                        v1 += (fi4 & (fi4 & (^(v1))))
                        switch (v1 << ((42) & 63)) {
                        case 32768:
                            switch (v5 + v5) {
                            case "foo bar":
                                var v11 float64 = (((-2.25 - (0.1 / 1.0)) - 1.5) - ((float64((float32(1e10) * float32(1.5))) / 1.0) + 1e10))
                                _ = v11
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                            default:
                                fmt.Println(v5, v1)
                            }
                        }
                        _ = v10
                    } else {
                        var v12 int8 = -1
                        _ = v12
                    }
                }
                var v13 int64 = 0
                _ = v13
            }
            _ = fi4
            _ = v5
        }
        v14 := GenS2{F0: ("hello" + "hello"), F1: v1, F2: (v1 % ((v1) | 1))}
        var v15 uint64 = 0
        _ = v14
        _ = v15
    } else {
        var v16 uint32 = 1
        var v17 int8 = 1
        _ = v16
        _ = v17
    }
    if (!true) {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 int = (v1 - 32767)
            var v20 uint16 = 255
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 bool = ((!false) && true)
        _ = v21
    } else {
        v3 = (v3 + v3)
        switch (v1 | (v1 * v1)) {
        case 557:
            v22 := GenS1{F0: -249, F1: 1}
            _ = v22
        case 708:
            var v23 int64 = 1
            _ = v23
        default:
            if ((true && false) || (!("\n\t" >= ""))) {
                if true {
                    v24 := GenS3{F0: 42, F1: 1, F2: "a"}
                    fmt.Println(v1)
                    _ = v24
                }
            } else {
                v25 := GenS1{F0: -1, F1: 65535}
                var v26 bool = ((-2147483648 != -1) || (!true))
                _ = v25
                _ = v26
            }
        }
    }
    var v27 int = (v1 | (v1 + (v1 / (((-(v1))) | 1))))
    var v28 bool = (true && (9223372036854775807 >= 0))
    v3 -= (v3 % ((v3) | 1))
    v2 = append(v2, 517)
    if len(v2) > 0 {
        v2[0] = 42
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        fmt.Println(fi29)
        var v30 int = fi29
        var v31 uint16 = 255
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int16 = -1
    var v33 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
}
