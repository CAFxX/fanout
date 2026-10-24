package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    var v4 uint16 = 223
    v5 := []uint32{65535, 1, 42, 2147483647}
    switch ("" + ("hello" + "a")) {
    case "\n\t":
        var v6 int8 = 42
        _ = v6
    case "":
        var v7 int64 = -414
        _ = v7
    case "a":
        fmt.Println(v1, v1)
    default:
        switch 403 {
        case 63:
            v8 := GenS1{F0: (v4 + 2), F1: 185}
            _ = v8
        }
    }
    v1 -= 506
    if false {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: (v4 + (v4 / ((v4) | 1))), F1: 128}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: (v4 ^ v4), F1: 1}
        v4 -= 96
        _ = v11
    } else {
        fmt.Println(v1)
        v2["a"] = ("a" + "")
        var v12 uint16 = v4
        _ = v12
    }
    var v13 uint = v1
    var v14 uint32 = 65535
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v15 int32 = -2147483648
    for fi16 := 0; fi16 < 1; fi16++ {
        switch fi16 {
        case -690:
            v17 := GenS1{F0: 42, F1: 127}
            _ = v17
        case 127:
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS2{F0: uint64(v13), F1: 32767}
                v20 := GenS1{F0: (v4 | 65535), F1: 255}
                switch "\n\t" {
                case "a":
                    v1 = v13
                case "hello":
                    var v21 int = (fi16 + (-(fi18)))
                    _ = v21
                default:
                    if (1.5 != float64(0.0)) {
                        v15 += (^(v15))
                    } else {
                        _ = len(v5)
                        _ = len(v3)
                    }
                    var v22 uint8 = 25
                    var v23 uint8 = 127
                    _ = v22
                    _ = v23
                }
                v4 = 233
                switch fi18 {
                case 0:
                    v24 := GenS1{F0: (v4 / ((2) | 1)), F1: 1}
                    _ = v24
                default:
                    if false {
                        v25 := GenS3{F0: (fi18 % (((fi18 % (((fi18 << (((v13 + (v1 - 478))) & 63))) | 1))) | 1))}
                        v4 -= v4
                        switch "foo bar" {
                        case "foo bar":
                            v14 += (v14 * v14)
                        case "hello":
                            v15 = (^(v15))
                        default:
                            switch (^(fi18)) {
                            case 0:
                                switch (fi16 | fi18) {
                                case -1:
                                    v3[v14] = true
                                case -727:
                                    v14 += 4294967295
                                    if (980 <= 18446744073709551615) {
                                        v14 = v14
                                        v13 += v13
                                    } else {
                                        v4 = (v4 % (((v4 / ((v4) | 1))) | 1))
                                        if ((!false) || true) {
                                            fmt.Println(v15, v15)
                                            var v26 int32 = -2147483648
                                            v27 := GenS4{F0: 0, F1: 0, F2: (v14 | (v14 | (v14 - (^(v14)))))}
                                            _ = len(v5)
                                            var v28 bool = (!true)
                                            v15 -= (v15 - -2147483648)
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    var v29 int8 = 0
                                    _ = v29
                                case 42:
                                    v30 := GenS2{F0: 4294967295, F1: 377}
                                    fmt.Println(v13, v4)
                                    _ = v30
                                }
                            case 2147483647:
                                var v31 uint = (v1 / ((216) | 1))
                                _ = v31
                            case 128:
                                var v32 float32 = (float32(3.14159) / 0.1)
                                _ = v32
                            }
                        }
                        _ = v25
                    } else {
                        v14 = v14
                        v33 := GenS5{F0: 42}
                        _ = v33
                    }
                }
                var v34 float64 = 3.14159
                if (true && ((v14 - 2) != 65535)) {
                    v35 := GenS1{F0: v4, F1: 1}
                    _ = v35
                }
                _ = fi18
                _ = v19
                _ = v20
                _ = v34
            }
        default:
            if (42 != uint32(v1)) {
                v36 := GenS5{F0: -1}
                var v37 int64 = 0
                _ = v36
                _ = v37
            }
        }
        _ = fi16
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        var v39 uint = 4294967295
        _ = fi38
        _ = v39
    }
    v2["hello"] = ""
    v14 += v14
    v15 += (v15 | (v15 + (v15 + (v15 % ((-2147483648) | 1)))))
    v40 := GenS1{F0: v4, F1: 0}
    if (false || (false && ("a" == "hello"))) {
        var v41 string = "a"
        _ = v41
    } else {
        fmt.Println(v15, v13, v1)
        _ = len(v2)
    }
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
