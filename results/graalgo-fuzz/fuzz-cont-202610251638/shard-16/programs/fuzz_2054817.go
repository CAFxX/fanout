package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 float64 = -2.25
    var v4 int16 = 347
    var v5 uint = 255
    v6 := []int{-1, 191}
    var v7 int64 = 9223372036854775807
    var v8 uint = 1
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int32 = 42
        v10 += v10
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v3)
    if true {
        if true {
            if (false || (false || true)) {
                if false {
                    var v11 float32 = 3.14159
                    _ = v11
                } else {
                    v8 -= v8
                    fmt.Println(v8)
                    delete(v2, int(v7))
                }
                v6 = append(v6, 1)
                var v12 float64 = v3
                _ = v12
            } else {
                v6 = append(v6, 32768)
                v8 -= (^(v5))
                fmt.Println(v4, v3, v5)
            }
            v13 := GenS1{F0: -9223372036854775808, F1: 0, F2: 662}
            _ = v13
        } else {
            if true {
                for fi14 := 0; fi14 < 12; fi14++ {
                    if ((!(!false)) && (true || true)) {
                        for fi15 := 0; fi15 < 14; fi15++ {
                            v5 -= 65535
                            switch ((("hello" + "foo bar") + (("hello" + "a") + "foo bar")) + "foo bar") {
                            case "\n\t":
                                v6 = append(v6, 4294967295)
                            case "a":
                                var v16 int32 = 42
                                _ = v16
                            default:
                                fmt.Println(v3, v5)
                            }
                            switch (("" + "\n\t") + (("foo bar" + "foo bar") + "hello")) {
                            case "":
                                v17 := GenS2{F0: 65535, F1: "a", F2: 0}
                                _ = v17
                            case "hello":
                                if false {
                                    var v18 uint16 = 2
                                    v19 := GenS1{F0: (v7 & int64(v4)), F1: 0, F2: -2147483648}
                                    _ = v18
                                    _ = v19
                                }
                            default:
                                switch (("hello" + "a") + "foo bar") {
                                case "foo bar":
                                    var v20 uint8 = 154
                                    _ = v20
                                case "a":
                                    var v21 int64 = v7
                                    _ = v21
                                case "\n\t":
                                    v7 -= int64(fi14)
                                }
                            }
                            _ = fi15
                        }
                        if (float64(3.14159) != 0.1) {
                            v4 -= 0
                        }
                    }
                    _ = fi14
                }
                if true {
                    v5 += v8
                    v8 = 436
                }
            } else {
                fmt.Println(v7, v7)
            }
            var v22 float64 = v3
            _ = v22
        }
        var v23 int8 = -128
        for fi24 := 0; fi24 < 7; fi24++ {
            if (316 != 129) {
                var v25 uint = v5
                if len(v6) > 0 {
                    v6[0] = (fi24 / ((fi24) | 1))
                }
                _ = v25
            }
            if false {
                v8 = v8
            } else {
                if true {
                    v7 -= (-(v7))
                    var v26 int64 = v7
                    _ = v26
                }
                v27 := GenS1{F0: v7, F1: 835, F2: 2147483647}
                _ = v27
            }
            v28 := GenS3{F0: (v7 | int64(v23)), F1: 2}
            _ = fi24
            _ = v28
        }
        _ = v23
    } else {
        v29 := GenS4{F0: 1}
        _ = v29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS3{F0: (v7 * (v7 + (v7 / (((v7 | v7)) | 1)))), F1: 1}
            v33 := GenS3{F0: (-(v7)), F1: 42}
            _ = fi31
            _ = v32
            _ = v33
        }
        v34 := GenS2{F0: 552, F1: ("" + "foo bar"), F2: 253}
        _ = fi30
        _ = v34
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        if false {
            for fi36 := 0; fi36 < 4; fi36++ {
                v4 -= v4
                _ = fi36
            }
            var v37 uint64 = 42
            _ = v37
        } else {
            v3 = v3
            v6 = append(v6, fi35)
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            var v39 uint = (v5 * v5)
            _ = fi38
            _ = v39
        }
        _ = fi35
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        switch "a" {
        case "hello":
            var v41 uint8 = 0
            _ = v41
        default:
            if true {
                for fi42 := 0; fi42 < 15; fi42++ {
                    v43 := GenS5{F0: v8}
                    var v44 uint32 = 0
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                v45 := GenS3{F0: v7, F1: 720}
                _ = v45
            } else {
                var v46 uint8 = 255
                var v47 uint = 18446744073709551615
                delete(v1, fi40)
                _ = v46
                _ = v47
            }
        }
        _ = fi40
    }
    var v48 float64 = (0.0 + (float64(((3.14159 + 0.1) * ((1.5 / 3.14159) * 0.1))) * (v3 * (((v3 + 1.5) * v3) / (float64(float32(v3)) - (float64(1e10) - v3))))))
    v4 += (v4 + 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
