package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []string{("hello" + "foo bar"), ("foo bar" + "hello")}
    var v2 string = "hello"
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    var v6 int16 = -1
    v6 = -1
    if true {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 128, F1: 42}
            _ = fi7
            _ = v8
        }
        fmt.Println(v2, v3, v2)
        switch v2 {
        case "\n\t":
            var v9 uint8 = 2
            _ = v9
        case "foo bar":
            v3 = v3
        }
    } else {
        switch 42 {
        case 469:
            fmt.Println(v3, v2)
        case 877:
            v10 := v5["a"]
            _ = v10
        case 4294967295:
            switch 77 {
            case 965:
                if false {
                    var v11 float64 = ((v3 + 1e10) / v3)
                    var v12 int32 = -786
                    _ = v11
                    _ = v12
                }
            }
        default:
            var v13 int = -2147483648
            _ = v13
        }
        v2 = v2
        _ = len(v1)
    }
    fmt.Println(v6, v6, v2)
    v14 := GenS2{F0: 87}
    v2 = v2
    v15 := v4[("a" + "\n\t")]
    var v16 uint32 = (v15 / (((^(v15))) | 1))
    v2 = "hello"
    v15 += 2
    v17 := GenS1{F0: 1, F1: 1}
    if (42 <= 42) {
        if true {
            var v18 uint32 = (v16 >> ((4294967295) & 63))
            if true {
                v16 = v15
                var v19 float32 = (0.1 + ((float32(v3) * (-2.25 + 1.5)) / ((-2.25 * 3.14159) + 1e10)))
                v3 = 3.14159
                _ = v19
            }
            fmt.Println(v15, v6)
            _ = v18
        }
        if true {
            var v20 uint = 65535
            _ = v20
        } else {
            switch "\n\t" {
            case "a":
                var v21 int = 127
                _ = v21
            case "\n\t":
                var v22 float64 = (1e10 + 3.14159)
                _ = v22
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                for fi24 := 0; fi24 < 6; fi24++ {
                    fmt.Println(v16)
                    var v25 uint = 292
                    if ((v16 + (v16 % ((v15) | 1))) > (v16 | (v16 / ((v15) | 1)))) {
                        if (-9223372036854775808 <= fi23) {
                            if (!(301 <= 791)) {
                                var v26 int8 = -128
                                _ = v26
                            }
                            v3 = float64(0.0)
                        } else {
                            switch -463 {
                            case -594:
                                fmt.Println(v2, v16)
                            case 9223372036854775807:
                                switch (v2 + v2) {
                                case "a":
                                    var v27 uint8 = 1
                                    _ = v27
                                default:
                                    fmt.Println(v2, v3)
                                }
                            case 207:
                                fmt.Println(v2, v15, fi23)
                            default:
                                var v28 string = (v2 + (v2 + v2))
                                _ = v28
                            }
                            switch "a" {
                            case "\n\t":
                                var v29 uint32 = (v15 & 269)
                                _ = v29
                            case "hello":
                                v30 := GenS3{F0: (v16 | v16), F1: v25, F2: 2}
                                _ = v30
                            default:
                                var v31 int8 = -1
                                _ = v31
                            }
                            if (((float32(v3) - 0.1) + (0.1 / 1.0)) <= 3.14159) {
                                v32 := GenS4{F0: 1, F1: 0}
                                fmt.Println(v2, fi24, fi23)
                                _ = v32
                            }
                        }
                        fmt.Println(v3)
                        switch ("hello" + (v2 + (v2 + "hello"))) {
                        case "\n\t":
                            var v33 float32 = 1e10
                            _ = v33
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                var v34 int64 = -9223372036854775808
                _ = fi23
                _ = v34
            }
        }
    } else {
        v35 := GenS5{F0: 978}
        v15 += (v15 - (v15 << ((0) & 63)))
        v36 := GenS1{F0: 35, F1: 0}
        _ = v35
        _ = v36
    }
    v1 = append(v1, "foo bar")
    _ = len(v1)
    var v37 int32 = int32(v6)
    v15 = (v16 | v16)
    if (-439 < 0) {
        v37 -= v37
        v3 = v3
    }
    fmt.Println(v37, v2, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
