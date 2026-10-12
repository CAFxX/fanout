package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }

func main() {
    var v1 uint = 65535
    v2 := []string{("" + "foo bar"), ("hello" + "\n\t"), ""}
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 9; fi7++ {
        switch (fi7 >> ((v1) & 63)) {
        case 16:
            v2 = append(v2, "")
        case 507:
            v1 = 909
        case 32767:
            if (((1e10 + (((1.5 + 0.0) / 1.5) + -2.25)) == ((0.1 / (-2.25 + float32((1.5 + (3.14159 * -2.25))))) - 1.5)) || false) {
                _ = len(v6)
                v8 := GenS1{F0: (fi7 - 127), F1: 1, F2: v1}
                var v9 uint = (v1 & v1)
                _ = v8
                _ = v9
            } else {
                switch "\n\t" {
                case "hello":
                    v1 += (v1 / (((v1 * 18446744073709551615)) | 1))
                case "foo bar":
                    var v10 int32 = 2147483647
                    _ = v10
                }
                _ = len(v4)
            }
        default:
            var v11 uint16 = 920
            _ = v11
        }
        if (!((fi7 >> ((42) & 63)) == (^(fi7)))) {
            var v12 int64 = 1
            v2 = append(v2, ("a" + "hello"))
            _ = v12
        } else {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 uint8 = 127
                v15 := GenS1{F0: -2147483648, F1: 0, F2: 1}
                if (!(!(!false))) {
                    var v16 int = 115
                    v17 := GenS1{F0: (-(v16)), F1: 1, F2: v1}
                    _ = v16
                    _ = v17
                } else {
                    switch ("hello" + "a") {
                    case "a":
                        v18 := GenS1{F0: fi13, F1: 0, F2: 65535}
                        _ = v18
                    case "":
                        v2 = append(v2, (("foo bar" + "a") + "hello"))
                    case "foo bar":
                        if (!(!true)) {
                            fmt.Println(fi13, fi13, v1)
                            v3["\n\t"] = (("a" + ("\n\t" + "hello")) + (("hello" + "hello") + ""))
                            if (!false) {
                                var v19 float32 = 1.5
                                switch "" {
                                case "a":
                                    fmt.Println(fi7, fi7, v1)
                                }
                                var v20 int32 = -2147483648
                                _ = v19
                                _ = v20
                            }
                            v14 = v14
                            fmt.Println(v1, fi7, fi7)
                        } else {
                            fmt.Println(fi7)
                            var v21 uint64 = 90
                            var v22 float32 = ((float32((float64(1.5) + (-2.25 / ((0.1 / 1.0) + 0.0)))) + (0.1 / -2.25)) + (float32((((1.5 - 0.0) / float64(((1.5 - (1.5 + 0.1)) * -2.25))) * 1.5)) - 0.1))
                            _ = v21
                            _ = v22
                        }
                    }
                    if false {
                        fmt.Println(v1, fi7, v14)
                        v14 = 127
                    } else {
                        v1 -= v1
                        var v23 int64 = 9223372036854775807
                        v24 := GenS1{F0: (-(fi13)), F1: -324, F2: (v1 ^ (v1 & v1))}
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(fi13, fi13, fi13)
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            var v25 string = "foo bar"
            switch fi7 {
            case -551:
                v1 = 4294967295
            case 128:
                var v26 int16 = 32767
                _ = v26
            case 42:
                v2 = append(v2, ((("a" + ("\n\t" + ("hello" + "hello"))) + "hello") + (v25 + (v25 + v25))))
            }
            _ = v25
        }
        _ = fi7
    }
    var v27 uint = (v1 | v1)
    v1 = v1
    v28 := GenS1{F0: 1, F1: -1, F2: v27}
    var v29 uint32 = 2147483647
    for fi30 := 0; fi30 < 14; fi30++ {
        fmt.Println(v1)
        _ = fi30
    }
    v2 = append(v2, ("a" + "\n\t"))
    v1 += (v27 / ((692) | 1))
    v1 -= (^(v27))
    for fi31 := 0; fi31 < 17; fi31++ {
        switch fi31 {
        case -158:
            fmt.Println(v1)
        default:
            v32 := GenS1{F0: 128, F1: 42, F2: (^(v27))}
            _ = v32
        }
        _ = fi31
    }
    _ = len(v2)
    var v33 int16 = 1
    var v34 uint32 = v29
    switch 833 {
    case 430:
        var v35 uint64 = 56
        _ = v35
    case 773:
        v34 -= v29
    case 0:
        var v36 uint = v27
        _ = v36
    default:
        v27 = v1
    }
    v27 = v27
    v1 = 65535
    var v37 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
