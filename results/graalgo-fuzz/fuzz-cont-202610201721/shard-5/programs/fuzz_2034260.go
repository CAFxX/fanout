package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 bool }
func genFn0() uint8 {
    v5 := GenS1{F0: 0, F1: 214, F2: 255}
    var v6 int64 = -582
    v6 = v6
    fmt.Println(v6)
    for fi7 := 0; fi7 < 14; fi7++ {
        v6 = v6
        v8 := GenS1{F0: 936, F1: 2147483648, F2: 42}
        if (!true) {
            var v9 string = "foo bar"
            _ = v9
        } else {
            for fi10 := 0; fi10 < 11; fi10++ {
                v6 = v6
                v11 := []int{(fi7 >> ((18446744073709551615) & 63)), fi7}
                if (1e10 > (0.0 / 1e10)) {
                    v6 -= (v6 - v6)
                    switch fi10 {
                    case 9223372036854775807:
                        if true {
                            if (722 == 2) {
                                v12 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: 2}
                                _ = v12
                            }
                            if true {
                                var v13 int64 = (v6 - 68)
                                if len(v11) > 0 {
                                    v11[0] = fi7
                                }
                                var v14 uint8 = 127
                                _ = v13
                                _ = v14
                            } else {
                                var v15 uint8 = 255
                                v15 = v15
                                var v16 float64 = 0.1
                                v6 = (v6 | (v6 ^ v6))
                                _ = len(v11)
                                _ = v15
                                _ = v16
                            }
                            fmt.Println(v6)
                            v11 = append(v11, (fi7 - (fi10 * (fi10 + (fi7 | fi10)))))
                            v17 := GenS1{F0: 2, F1: 4294967295, F2: 42}
                            _ = v17
                        }
                    case 4294967295:
                        v18 := GenS1{F0: 440, F1: 725, F2: 255}
                        _ = v18
                    case 0:
                        switch "" {
                        case "hello":
                            v6 = v6
                        default:
                            v6 -= (v6 * (v6 | (^(v6))))
                        }
                    default:
                        var v19 uint16 = 255
                        _ = v19
                    }
                    fmt.Println(v6, fi10, fi10)
                } else {
                    if len(v11) > 0 {
                        v11[0] = fi10
                    }
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        v20 := GenS1{F0: 931, F1: 156, F2: 42}
                        _ = v20
                    case "":
                        v6 -= 445
                    default:
                        switch "" {
                        case "\n\t":
                            var v21 float64 = 0.1
                            _ = v21
                        }
                    }
                    var v22 uint = 42
                    _ = v22
                }
                _ = fi10
                _ = v11
            }
            v23 := []int{fi7}
            fmt.Println(fi7, v6)
            _ = v23
        }
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 0
    var v3 bool = false
    var v4 uint = 371
    if (int(v2) < 9223372036854775807) {
        fmt.Println(v3, v2)
        v24 := GenS2{F0: 550, F1: 4294967295, F2: false}
        _ = v24
    }
    v25 := GenS2{F0: 80, F1: 457, F2: (0.0 > 0.0)}
    var v26 uint = v4
    fmt.Println(v3)
    v27 := GenS1{F0: (v26 / (((v26 << ((65535) & 63))) | 1)), F1: v2, F2: 255}
    v26 = v26
    v3 = v3
    fmt.Println(v4, v26)
    switch 601 {
    case 32768:
        v2 = (^(v2))
    case 1:
        v3 = (!false)
    case -842:
        v28 := GenS1{F0: 255, F1: 0, F2: 134}
        _ = v28
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            v3 = (1 == 18446744073709551615)
            v30 := []int{148, (fi29 & (fi29 % (((fi29 >> (((v26 ^ 255)) & 63))) | 1))), (^(fi29))}
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v4)
    var v31 string = "hello"
    v31 = "a"
    v32 := []int{127, 4294967295}
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 uint64 = 2
        var v35 int = fi33
        v31 = v31
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v2, v31, v3)
    v36 := GenS3{F0: 65535, F1: -1, F2: (1 <= -13)}
    var v37 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v4
}
