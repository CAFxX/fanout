package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    v3 := []int{32768}
    v2 = (("" + v2) + "foo bar")
    switch ("hello" + (v2 + v2)) {
    case "":
        if (!(4294967295 < 4294967295)) {
            var v4 uint16 = 0
            var v5 float32 = (1.5 * 1.5)
            _ = v4
            _ = v5
        }
    case "\n\t":
        switch 0 {
        case 128:
            if true {
                var v6 int32 = -601
                v3 = append(v3, 1)
                _ = v6
            }
        case 1:
            for fi7 := 0; fi7 < 6; fi7++ {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 uint64 = 42
                    fmt.Println(v2, v2, fi8)
                    var v10 uint32 = 2
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 int32 = 0
                switch v2 {
                case "foo bar":
                    v2 = "hello"
                case "":
                    if (!(false && (-550 > 1))) {
                        if (!(42 == -1)) {
                            if (2 == 4294967295) {
                                var v12 uint32 = 0
                                v13 := GenS1{F0: v12, F1: ((true && true) || true)}
                                if (!false) {
                                    for fi14 := 0; fi14 < 20; fi14++ {
                                        var v15 bool = false
                                        v16 := GenS2{F0: v12, F1: 42, F2: v12}
                                        if (int16(v11) > 1) {
                                            var v17 string = v2
                                            var v18 uint16 = 42
                                            _ = v17
                                            _ = v18
                                        } else {
                                            var v19 uint8 = 87
                                            if v15 {
                                                v20 := GenS2{F0: v12, F1: 18446744073709551615, F2: (v12 ^ v12)}
                                                fmt.Println(v2)
                                                var v21 bool = (true && v15)
                                                _ = v20
                                                _ = v21
                                            }
                                            _ = v19
                                        }
                                        _ = fi14
                                        _ = v15
                                        _ = v16
                                    }
                                    v12 = 472
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    for fi22 := 0; fi22 < 3; fi22++ {
                                        v12 += (^(v12))
                                        v23 := GenS1{F0: (v12 >> ((42) & 63)), F1: true}
                                        v24 := v1[v2]
                                        v25 := GenS2{F0: v12, F1: 4294967295, F2: 338}
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                _ = v12
                                _ = v13
                            } else {
                                v2 = "hello"
                            }
                            if (-2.25 <= (float32(1.5) * ((((3.14159 - 0.0) * 0.1) - 0.1) + 1.5))) {
                                v11 = 2147483647
                                fmt.Println(v11, fi7)
                                var v26 int32 = 143
                                _ = v26
                            }
                        }
                        var v27 uint16 = 42
                        _ = v27
                    }
                case "hello":
                    for fi28 := 0; fi28 < 2; fi28++ {
                        switch (-(fi7)) {
                        case -597:
                            v11 = (v11 % (((v11 << ((18446744073709551615) & 63))) | 1))
                        case 9223372036854775807:
                            switch v2 {
                            case "":
                                var v29 float64 = 1e10
                                _ = v29
                            default:
                                v11 -= 2147483647
                            }
                        }
                        if true {
                            v30 := GenS2{F0: 2147483648, F1: 4294967295, F2: 560}
                            switch v2 {
                            case "":
                                var v31 float64 = float64(float32(0.1))
                                _ = v31
                            default:
                                var v32 uint32 = 65535
                                _ = v32
                            }
                            _ = v30
                        }
                        var v33 int64 = 1
                        _ = fi28
                        _ = v33
                    }
                default:
                    var v34 int8 = -128
                    _ = v34
                }
                _ = fi7
                _ = v11
            }
        default:
            v35 := GenS1{F0: 503, F1: true}
            _ = v35
        }
    default:
        v36 := GenS2{F0: 2, F1: 18446744073709551615, F2: 42}
        _ = v36
    }
    if (false && false) {
        if (-107 == 0) {
            delete(v1, v2)
            v37 := GenS2{F0: 0, F1: 224, F2: 703}
            _ = v37
        } else {
            for fi38 := 0; fi38 < 15; fi38++ {
                v2 = "hello"
                _ = fi38
            }
        }
    }
    v39 := v1[((v2 + ("" + v2)) + "\n\t")]
    fmt.Println(v39, v39)
    v40 := v1[("hello" + (v2 + v2))]
    v39 = ((1e10 + 1.5) == (((0.1 + (3.14159 + (0.1 + (1e10 * (((1e10 / 1.0) + 0.0) / 1.5))))) - 0.1) / 0.1))
    if len(v3) > 0 {
        v3[0] = -501
    }
    v40 = true
    switch 128 {
    case 4294967295:
        delete(v1, "foo bar")
    case -878:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -1:
        delete(v1, "a")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}
