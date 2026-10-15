package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32; F1 bool }
func genFn0() uint32 {
    var v4 uint8 = 42
    var v5 int32 = 42
    _ = v4
    _ = v5
    return 527
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 int64 = -592
    var v3 int64 = (v2 * 233)
    var v6 float32 = -2.25
    v7 := GenS1{F0: 65535, F1: 127, F2: 2}
    if len(v1) > 0 {
        v1[0] = 32768
    }
    v1 = append(v1, 0)
    switch "\n\t" {
    case "":
        switch "" {
        case "foo bar":
            var v8 int16 = 32767
            _ = v8
        case "a":
            v3 -= v2
        }
    case "\n\t":
        for fi9 := 0; fi9 < 1; fi9++ {
            switch "" {
            case "":
                if false {
                    fmt.Println(v6, v3)
                    if len(v1) > 0 {
                        v1[0] = (fi9 + fi9)
                    }
                } else {
                    var v10 float64 = (((1.5 / float64(-2.25)) * float64((1.5 + v6))) / ((0.0 + float64(0.0)) + (0.1 * -2.25)))
                    v11 := GenS1{F0: 838, F1: 1, F2: 962}
                    _ = v10
                    _ = v11
                }
            case "a":
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 int8 = 42
                    v2 = (v2 - -9223372036854775808)
                    fmt.Println(fi9, v13, fi9)
                    var v14 uint16 = 1
                    _ = len(v1)
                    if (-32768 <= -32768) {
                        v13 = (v13 - -128)
                        if (true || (2 >= 1)) {
                            switch fi9 {
                            case -548:
                                var v15 uint64 = 0
                                _ = v15
                            }
                        } else {
                            if ((((v14 ^ 2) == (v14 * v14)) && false) && (161 <= uint32(v14))) {
                                if true {
                                    var v16 uint8 = 127
                                    _ = v16
                                }
                                var v17 float64 = ((1e10 - 3.14159) - 0.1)
                                _ = v17
                            } else {
                                if true {
                                    switch (fi9 ^ fi9) {
                                    case -927:
                                        switch "\n\t" {
                                        case "":
                                            v18 := GenS2{F0: (v3 % (((v3 | v2)) | 1)), F1: 42}
                                            _ = v18
                                        case "\n\t":
                                            fmt.Println(v6)
                                        default:
                                            if false {
                                                v14 += (v14 - (^(v14)))
                                                var v19 uint16 = v14
                                                _ = v19
                                            } else {
                                                fmt.Println(fi9, fi9, v2)
                                            }
                                        }
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v20 uint = 4294967295
                                    _ = v20
                                } else {
                                    if true {
                                        var v21 string = "\n\t"
                                        _ = v21
                                    } else {
                                        v22 := GenS1{F0: 4294967295, F1: -1, F2: 42}
                                        v14 = v14
                                        _ = v22
                                    }
                                    switch fi9 {
                                    case 4294967295:
                                        v23 := GenS2{F0: -9223372036854775808, F1: 1}
                                        _ = v23
                                    case 32767:
                                        var v24 uint16 = v14
                                        v6 -= (1.5 - -2.25)
                                        _ = v24
                                    }
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v25 float32 = v6
                        v26 := GenS3{F0: 2147483647, F1: false}
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS1{F0: 1, F1: 0, F2: 18446744073709551615}
                        v14 = (v14 % (((v14 + (^(v14)))) | 1))
                        _ = v27
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            _ = fi9
        }
    case "hello":
        v2 = 0
    default:
        if len(v1) > 0 {
            v1[0] = -79
        }
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    v2 += v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 += ((v6 * v6) + 0.1)
    switch (("\n\t" + "\n\t") + "a") {
    case "":
        switch -9223372036854775808 {
        case 2147483647:
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 string = "\n\t"
                v29 = v29
                _ = fi28
                _ = v29
            }
        case -9223372036854775808:
            for fi30 := 0; fi30 < 2; fi30++ {
                v6 += -2.25
                _ = fi30
            }
        }
    case "hello":
        _ = len(v1)
    case "a":
        fmt.Println(v2)
    default:
        var v31 int = 32767
        _ = v31
    }
    var v32 string = ("foo bar" + ("a" + ("hello" + "hello")))
    var v33 uint64 = 490
    var v34 int8 = -1
    var v35 int16 = -767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v6
    _ = v7
}
