package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []uint32{4294967295}
    v2 := []string{"\n\t", "a", "a", "a"}
    var v3 int16 = 42
    v4 := make(map[string]bool)
    v3 += 42
    if ("a" < ("hello" + ("\n\t" + "foo bar"))) {
        v5 := GenS1{F0: 9223372036854775807, F1: 42, F2: 2147483647}
        var v6 int32 = 1
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "\n\t" {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            fmt.Println(v3, fi7)
        }
        fmt.Println(v3, fi7)
        if false {
            var v8 int = 0
            v3 -= v3
            fmt.Println(v8, v3, fi7)
            _ = v8
        }
        _ = fi7
    }
    _ = len(v4)
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            switch fi9 {
            case 1:
                if (-751 != 0) {
                    var v10 uint32 = 919
                    v2 = append(v2, "foo bar")
                    if true {
                        if (!true) {
                            var v11 uint32 = 533
                            v3 = v3
                            _ = v11
                        }
                    } else {
                        _ = len(v4)
                        fmt.Println(v10)
                        v10 = v10
                    }
                    _ = v10
                } else {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        var v13 string = "hello"
                        var v14 uint = 178
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v15 := GenS1{F0: fi9, F1: 197, F2: fi9}
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if true {
                            if true {
                                var v17 int = -9223372036854775808
                                v3 = v3
                                v17 += fi9
                                _ = v17
                            } else {
                                fmt.Println(v3, fi16)
                            }
                        } else {
                            fmt.Println(fi9, v3, fi9)
                        }
                        var v18 bool = true
                        _ = fi16
                        _ = v18
                    }
                    _ = v15
                }
            }
            if (776 == 714) {
                v3 = v3
            } else {
                if ("a" != "foo bar") {
                    if ((1 == 245) && true) {
                        v3 += (v3 << ((255) & 63))
                        if true {
                            var v19 bool = false
                            v20 := GenS2{F0: 2, F1: 32767}
                            _ = v19
                            _ = v20
                        } else {
                            var v21 bool = (!(!(231 > 1)))
                            _ = v21
                        }
                        _ = len(v2)
                    }
                }
                _ = len(v1)
            }
            _ = fi9
        }
        fmt.Println(v3)
    } else {
        var v22 int = -9223372036854775808
        _ = v22
    }
    var v23 float64 = ((float64(0.0) / ((0.1 - (((1e10 - (0.0 - ((0.0 - (1.5 + float64(float32(0.1)))) - 0.0))) + float64(-2.25)) + 0.0)) * ((((1.5 - -2.25) / 0.1) + (1e10 * float64(-2.25))) - (3.14159 / -2.25)))) - (-2.25 + ((((((0.0 + 0.0) + (1e10 - float64(-2.25))) + 0.1) / (1e10 * float64(0.1))) * 1e10) + 1e10)))
    if (2 <= 2147483648) {
        fmt.Println(v3, v3)
    }
    v23 = (((0.0 / (0.0 / v23)) * v23) + v23)
    v23 += 0.1
    if true {
        if true {
            if (!(!(!(-17 < 127)))) {
                v23 = float64((float32(v23) + -2.25))
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            } else {
                _ = len(v2)
            }
            var v24 int8 = 1
            v25 := GenS1{F0: 128, F1: 128, F2: -347}
            _ = v24
            _ = v25
        } else {
            var v26 float32 = 1.5
            _ = v26
        }
        if false {
            var v27 float64 = float64(-2.25)
            var v28 int16 = v3
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 int = fi29
                switch (fi29 % ((-2147483648) | 1)) {
                case 127:
                    if false {
                        var v31 int16 = (v3 - v3)
                        var v32 float64 = v23
                        v33 := GenS3{F0: 68}
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
                for fi34 := 0; fi34 < 3; fi34++ {
                    var v35 uint16 = 218
                    fmt.Println(fi29, v23)
                    _ = fi34
                    _ = v35
                }
                _ = fi29
                _ = v30
            }
        }
        v36 := GenS3{F0: 55}
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
