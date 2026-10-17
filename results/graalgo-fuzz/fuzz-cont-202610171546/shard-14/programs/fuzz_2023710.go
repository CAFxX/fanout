package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint }
type GenS5 struct { F0 uint8; F1 uint16; F2 bool }

func main() {
    v1 := []string{}
    var v2 bool = false
    var v3 float64 = (float64(0.1) * 0.0)
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    var v6 bool = (-9223372036854775808 < 1)
    switch "" {
    case "foo bar":
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 0}
            if v6 {
                delete(v4, fi7)
                var v9 int8 = 7
                if false {
                    v10 := GenS2{F0: 127, F1: 2, F2: (^(v9))}
                    var v11 int8 = v9
                    var v12 uint64 = 2
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    switch 946 {
                    case 923:
                        switch fi7 {
                        case 4294967295:
                            if v2 {
                                var v13 int8 = -1
                                fmt.Println(v9, v3)
                                v1 = append(v1, ("a" + ("hello" + "a")))
                                _ = v13
                            }
                        case 162:
                            var v14 uint8 = 32
                            _ = v14
                        case -665:
                            for fi15 := 0; fi15 < 18; fi15++ {
                                var v16 uint = 327
                                fmt.Println(fi15)
                                _ = fi15
                                _ = v16
                            }
                        default:
                            var v17 string = ("a" + "\n\t")
                            _ = v17
                        }
                    case -362:
                        var v18 bool = v6
                        _ = v18
                    case 4294967295:
                        v19 := GenS3{F0: -591, F1: 91}
                        _ = v19
                    }
                }
                _ = v9
            } else {
                if v2 {
                    v20 := GenS4{F0: 765, F1: 429, F2: 4294967295}
                    fmt.Println(v3, fi7)
                    _ = v20
                } else {
                    var v21 uint = 255
                    var v22 uint64 = 0
                    _ = len(v1)
                    _ = v21
                    _ = v22
                }
            }
            _ = fi7
            _ = v8
        }
    case "":
        var v23 int64 = -1
        _ = v23
    }
    v3 = v3
    var v24 uint16 = 65535
    v3 = float64(((3.14159 + 3.14159) - float32(float64(0.0))))
    if v2 {
        for fi25 := 0; fi25 < 3; fi25++ {
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 uint16 = v24
                fmt.Println(fi26, fi25)
                v1 = append(v1, "")
                switch (-(fi25)) {
                case 728:
                    v27 = 42
                case -655:
                    var v28 int = 162
                    _ = v28
                case -37:
                    if v2 {
                        v5[4294967295] = v6
                        v3 += 0.1
                    }
                    if (-1 == -968) {
                        switch (fi25 / ((127) | 1)) {
                        case -1:
                            v29 := v4[32767]
                            _ = v29
                        case 1:
                            v27 = 21
                        }
                        var v30 float32 = (1.5 + 1e10)
                        _ = v30
                    } else {
                        if v2 {
                            switch (-(fi26)) {
                            case 912:
                                fmt.Println(v24, fi25, v2)
                            case -1:
                                var v31 float64 = ((v3 * ((0.1 + ((0.0 / v3) - (float64(0.1) - v3))) - (1.5 / (v3 - v3)))) * 0.0)
                                _ = v31
                            default:
                                if false {
                                    v1 = append(v1, ("foo bar" + ""))
                                    var v32 int64 = 772
                                    _ = v32
                                }
                                v33 := GenS1{F0: 42}
                                fmt.Println(v2, fi25)
                                _ = v33
                            }
                            v24 += (v24 / ((32767) | 1))
                        }
                    }
                default:
                    switch 128 {
                    case 32767:
                        var v34 int = fi26
                        _ = v34
                    default:
                        if (!(42 <= 17)) {
                            fmt.Println(v2)
                        }
                    }
                    fmt.Println(v27)
                    var v35 float32 = 3.14159
                    _ = v35
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        v36 := GenS4{F0: 791, F1: (v24 ^ v24), F2: 42}
        _ = v36
    }
    v37 := GenS5{F0: 106, F1: (v24 % (((v24 << ((4294967295) & 63))) | 1)), F2: (9223372036854775807 == 751)}
    var v38 int = 0
    v24 -= (^(v24))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
