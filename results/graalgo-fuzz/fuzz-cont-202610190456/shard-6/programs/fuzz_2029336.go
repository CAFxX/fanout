package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 int64, p1 string, p2 int8) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{-907}
        v8 := GenS1{F0: (^(p2)), F1: "\n\t", F2: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    if (0 >= 822) {
        var v9 uint64 = 4294967295
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    v3 := []int{-499, -9223372036854775808}
    v4 := make(map[int]bool)
    var v5 float32 = ((float32(1.5) / ((0.1 - (((0.1 / -2.25) / 0.1) + 0.0)) - -2.25)) - -2.25)
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 float32 = v5
            var v13 float64 = (1e10 - (1e10 - ((-2.25 + 3.14159) - 0.0)))
            fmt.Println(fi10, fi11, v13)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v1 = (v1 + "hello")
    switch v1 {
    case "hello":
        v3 = append(v3, -2147483648)
    case "":
        v1 = v1
    default:
        switch "\n\t" {
        case "a":
            fmt.Println(v5, v5)
        case "":
            if false {
                switch 75 {
                case 127:
                    if false {
                        switch 32768 {
                        case -2147483648:
                            for fi14 := 0; fi14 < 20; fi14++ {
                                v1 = ""
                                fmt.Println(v1, fi14)
                                _ = fi14
                            }
                        }
                        var v15 uint32 = 4294967295
                        if (987 < 18446744073709551615) {
                            var v16 int8 = 0
                            _ = v16
                        } else {
                            var v17 float64 = 0.0
                            _ = v17
                        }
                        _ = v15
                    }
                case 64:
                    v1 = (v1 + v1)
                }
                v1 = v1
            } else {
                v18 := GenS1{F0: 0, F1: ((v1 + v1) + v1), F2: 20}
                v5 = 0.1
                var v19 bool = false
                _ = v18
                _ = v19
            }
        case "\n\t":
            if ((false && true) && (157 == 394)) {
                var v20 int16 = -1
                _ = len(v4)
                _ = v20
            }
        }
    }
    var v21 int16 = -32768
    delete(v2, "hello")
    if (42 >= 42) {
        if (127 > -120) {
            v1 = v1
            var v22 uint32 = 495
            v21 += (^(v21))
            _ = v22
        }
        if (2 < 42) {
            var v23 float64 = 3.14159
            switch (v1 + ((v1 + v1) + "foo bar")) {
            case "hello":
                for fi24 := 0; fi24 < 13; fi24++ {
                    delete(v4, 128)
                    _ = fi24
                }
            case "":
                for fi25 := 0; fi25 < 4; fi25++ {
                    switch v1 {
                    case "a":
                        if (float32(float64((v5 + (3.14159 - v5)))) >= (-2.25 + 0.1)) {
                            switch (fi25 / ((fi25) | 1)) {
                            case 751:
                                v26 := GenS1{F0: -128, F1: v1, F2: 18446744073709551615}
                                _ = v26
                            case -2147483648:
                                v3 = append(v3, -1)
                            case 42:
                                var v27 uint16 = 0
                                _ = v27
                            default:
                                v1 = "a"
                            }
                            if (-1 >= 2147483647) {
                                v1 = ("a" + ((((v1 + v1) + v1) + (v1 + "foo bar")) + v1))
                                v23 += (v23 * (((v23 + 1e10) * 0.0) / v23))
                            } else {
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    v21 += v21
                                    var v29 bool = (!false)
                                    var v30 int32 = 515
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                var v31 int16 = (v21 * v21)
                                _ = v31
                            }
                            v23 += v23
                        }
                    case "":
                        v32 := GenS1{F0: 127, F1: v1, F2: 267}
                        _ = v32
                    }
                    v33 := GenS2{F0: 119, F1: -128}
                    _ = fi25
                    _ = v33
                }
            }
            var v34 int8 = 42
            _ = v23
            _ = v34
        }
        if (true || (true && true)) {
            fmt.Println(v1, v21)
        } else {
            v5 = 3.14159
        }
    } else {
        fmt.Println(v1, v21, v5)
    }
    v2["foo bar"] = v1
    fmt.Println(v21, v21)
    v35 := GenS3{F0: 42, F1: 789}
    v1 = v1
    v21 = v21
    v21 += (v21 | 0)
    if false {
        v5 += v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
