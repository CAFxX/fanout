package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -1
    v3 := make(map[string]int)
    var v4 float32 = ((0.0 * 3.14159) + float32(1.5))
    var v5 uint32 = 206
    fmt.Println(v2, v5, v2)
    for fi6 := 0; fi6 < 11; fi6++ {
        switch "" {
        case "foo bar":
            fmt.Println(v4, v5)
        default:
            v7 := v3[(("" + "hello") + "\n\t")]
            _ = v7
        }
        switch ("" + "a") {
        case "a":
            v8 := []int{0}
            _ = v8
        case "hello":
            v5 = 4294967295
        default:
            var v9 uint64 = 1
            _ = v9
        }
        fmt.Println(v5, v5)
        _ = fi6
    }
    if false {
        var v10 uint64 = 923
        _ = v10
    } else {
        var v11 int64 = v2
        _ = v11
    }
    var v12 uint16 = 0
    v13 := GenS1{F0: v12, F1: 787, F2: v12}
    v14 := GenS1{F0: v12, F1: -2147483648, F2: v12}
    v15 := GenS1{F0: v12, F1: 1, F2: (v12 << ((1) & 63))}
    fmt.Println(v2, v12)
    v16 := GenS1{F0: (^(v12)), F1: -725, F2: 42}
    if (!true) {
        v17 := []int{435, -323}
        switch "\n\t" {
        case "\n\t":
            switch "" {
            case "a":
                v2 = (v2 << ((18446744073709551615) & 63))
            }
        case "":
            fmt.Println(v12, v12, v2)
        default:
            switch 32767 {
            case 2147483647:
                if (true && ((v12 + (v12 % ((v12) | 1))) > (v12 * (v12 & v12)))) {
                    if ((v5 / ((2147483647) | 1)) != 0) {
                        _ = len(v17)
                        v17 = append(v17, 127)
                    } else {
                        var v18 int = 1
                        fmt.Println(v4, v2, v12)
                        _ = v18
                    }
                }
            case -56:
                v12 = 2
            case -9223372036854775808:
                if (1 > 523) {
                    if (v4 < (v4 + v4)) {
                        var v19 uint16 = v12
                        for fi20 := 0; fi20 < 2; fi20++ {
                            var v21 float64 = 1e10
                            v2 = 1
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                }
            }
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint = uint(v5)
            var v24 int32 = 42
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = v17
    }
    if (!false) {
        v25 := []int{-9223372036854775808, -928}
        v26 := GenS1{F0: v12, F1: 2147483647, F2: (v12 % ((v12) | 1))}
        _ = v25
        _ = v26
    } else {
        if true {
            var v27 uint8 = 22
            _ = v27
        } else {
            var v28 bool = false
            for fi29 := 0; fi29 < 7; fi29++ {
                if v28 {
                    if v28 {
                        switch (^(fi29)) {
                        case 32767:
                            for fi30 := 0; fi30 < 8; fi30++ {
                                if v28 {
                                    fmt.Println(fi29, v2, fi30)
                                    v31 := v1["a"]
                                    _ = v31
                                }
                                _ = fi30
                            }
                        case 1:
                            v4 += (v4 - float32(0.0))
                        case -714:
                            for fi32 := 0; fi32 < 6; fi32++ {
                                v12 -= v12
                                var v33 int16 = 546
                                if false {
                                    fmt.Println(v2)
                                    var v34 uint8 = 0
                                    switch (fi29 % (((fi29 ^ (fi32 ^ fi32))) | 1)) {
                                    case 32768:
                                        v35 := GenS1{F0: v12, F1: -1, F2: v12}
                                        _ = v35
                                    case -1:
                                        fmt.Println(fi29, v5)
                                    case 9223372036854775807:
                                        v36 := GenS1{F0: (^(v12)), F1: 1, F2: v12}
                                        _ = v36
                                    }
                                    _ = v34
                                }
                                _ = fi32
                                _ = v33
                            }
                        default:
                            v37 := []int{}
                            _ = v37
                        }
                    }
                    v2 = -9223372036854775808
                }
                var v38 uint64 = 341
                _ = fi29
                _ = v38
            }
            _ = v28
        }
        if ((v5 % ((v5) | 1)) < (^(v5))) {
            v2 = (v2 + v2)
            fmt.Println(v2, v2)
            v4 += 0.1
        }
        v39 := []int{2147483647}
        _ = v39
    }
    v2 -= v2
    v4 = 3.14159
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
