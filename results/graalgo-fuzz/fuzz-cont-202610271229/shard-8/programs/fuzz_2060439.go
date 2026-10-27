package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 uint; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int32 = 0
    var v2 int16 = -32768
    var v3 uint8 = 43
    v4 := GenS1{F0: -1, F1: 323, F2: (-(v1))}
    v5 := []int{46}
    if ((0.0 != 3.14159) || true) {
        switch "" {
        case "hello":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    } else {
        v2 = (v2 + 914)
        if true {
            v6 := GenS2{F0: 255, F1: "foo bar", F2: 42}
            _ = v6
        } else {
            if false {
                fmt.Println(v3, v3, v2)
            }
            for fi7 := 0; fi7 < 11; fi7++ {
                if len(v5) > 0 {
                    v5[0] = -9223372036854775808
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi7
            }
            v5 = append(v5, 984)
        }
        if (true || (1 <= 255)) {
            v1 += (v1 >> ((4294967295) & 63))
        }
    }
    var v8 uint16 = 2
    if (!(42 == 934)) {
        if (42 == -128) {
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 int16 = v2
                switch "a" {
                case "foo bar":
                    for fi11 := 0; fi11 < 12; fi11++ {
                        fmt.Println(v10)
                        var v12 uint32 = 2
                        _ = fi11
                        _ = v12
                    }
                case "hello":
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v2 -= (v2 >> ((42) & 63))
                    }
                }
                v13 := GenS2{F0: 255, F1: "hello", F2: 87}
                _ = fi9
                _ = v10
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 float64 = 0.1
            for fi16 := 0; fi16 < 20; fi16++ {
                switch (fi14 << ((4294967295) & 63)) {
                case 42:
                    v5 = append(v5, (fi14 >> ((210) & 63)))
                case 1:
                    var v17 int32 = 2147483647
                    _ = v17
                case -68:
                    if false {
                        v18 := GenS2{F0: 4294967295, F1: "", F2: 0}
                        switch "hello" {
                        case "hello":
                            v2 += 617
                        case "foo bar":
                            v3 -= v3
                        case "":
                            fmt.Println(fi14)
                        }
                        switch fi16 {
                        case 990:
                            var v19 float64 = ((v15 + 0.1) + ((3.14159 + v15) - (((v15 * v15) + (0.1 - float64(1.5))) - v15)))
                            _ = v19
                        case 128:
                            switch "foo bar" {
                            case "foo bar":
                                v20 := GenS1{F0: -42, F1: 42, F2: (v1 % ((42) | 1))}
                                _ = v20
                            default:
                                if true {
                                    v5 = append(v5, fi16)
                                    fmt.Println(v15)
                                    v2 -= (v2 >> ((880) & 63))
                                    _ = len(v5)
                                    var v21 string = ("" + "foo bar")
                                    _ = v21
                                } else {
                                    fmt.Println(v8, v3)
                                    var v22 float32 = 1.5
                                    if false {
                                        v23 := GenS3{F0: 807, F1: (v3 << ((649) & 63)), F2: -1}
                                        _ = v23
                                    }
                                    _ = v22
                                }
                            }
                            v3 = (^(v3))
                        default:
                            if false {
                                switch "foo bar" {
                                case "\n\t":
                                    var v24 uint = 943
                                    _ = v24
                                case "hello":
                                    v2 = (v2 & v2)
                                default:
                                    _ = len(v5)
                                    var v25 uint32 = 65535
                                    _ = v25
                                }
                            }
                            fmt.Println(v8, v1)
                            v3 += v3
                            var v26 uint32 = 1
                            v15 -= v15
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v1 = (v1 & v1)
                            }
                            if false {
                                var v27 uint8 = 1
                                var v28 uint32 = (v26 / (((v26 + v26)) | 1))
                                _ = v27
                                _ = v28
                            } else {
                                if (!(42 < (v2 | v2))) {
                                    var v29 bool = (!(true || ((fi14 * fi16) > 133)))
                                    v8 = (v8 ^ 18)
                                    _ = v29
                                } else {
                                    _ = len(v5)
                                    fmt.Println(fi16)
                                    var v30 int32 = (v1 / ((-424) | 1))
                                    _ = v30
                                }
                                var v31 string = "\n\t"
                                var v32 float32 = ((1e10 * 1e10) - ((1e10 - (1.5 + (3.14159 * 1.5))) * 1.5))
                                v32 -= 1e10
                                _ = v31
                                _ = v32
                            }
                            v8 = 330
                            _ = v26
                        }
                        _ = v18
                    }
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
    }
    v8 = 298
    fmt.Println(v8, v2, v1)
    var v33 uint8 = v3
    v34 := GenS4{F0: 63}
    v35 := GenS1{F0: 1, F1: 42, F2: v1}
    fmt.Println(v33, v8, v3)
    v1 -= v1
    v3 = 255
    fmt.Println(v33, v33)
    var v36 int8 = -1
    v33 = (v33 * (^(v3)))
    v37 := GenS5{F0: 94}
    var v38 float32 = ((-2.25 + (float32(0.1) / 0.1)) - ((float32(1.5) + (float32(1.5) / 0.1)) - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v8
}
