package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []uint32{2147483648, 87}
    var v2 int8 = 65
    v3 := make(map[int]bool)
    var v4 int32 = -2147483648
    var v5 uint32 = 1
    v6 := make(map[string]bool)
    if false {
        v2 = -128
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint32 = (v5 & v5)
            v9 := GenS1{F0: v4, F1: 248, F2: 128}
            delete(v3, 4294967295)
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            switch (fi10 + (fi10 * (fi10 >> ((18446744073709551615) & 63)))) {
            case 32768:
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 uint = 2
                    switch "foo bar" {
                    case "hello":
                        if true {
                            if (int16(v2) == 0) {
                                if len(v1) > 0 {
                                    v1[0] = v5
                                }
                                var v13 uint32 = v5
                                _ = v13
                            }
                            if ((v2 - (v2 >> (((v12 | (^(v12)))) & 63))) < v2) {
                                v14 := v6[("" + "foo bar")]
                                var v15 float64 = (0.1 * (-2.25 * 0.1))
                                var v16 int64 = -9223372036854775808
                                _ = v14
                                _ = v15
                                _ = v16
                            } else {
                                fmt.Println(fi10, fi11)
                            }
                            var v17 uint8 = uint8(v5)
                            _ = v17
                        } else {
                            v2 = (v2 - v2)
                            v2 += 100
                        }
                        var v18 bool = false
                        var v19 string = "\n\t"
                        _ = v18
                        _ = v19
                    }
                    switch ("\n\t" + "hello") {
                    case "":
                        v6["\n\t"] = false
                    default:
                        var v20 uint64 = 2
                        _ = v20
                    }
                    if (!false) {
                        if true {
                            var v21 uint8 = 2
                            _ = v21
                        }
                        var v22 int16 = 42
                        _ = v22
                    } else {
                        v23 := GenS2{F0: 255, F1: 32768, F2: (fi11 - (fi11 % ((fi10) | 1)))}
                        switch (fi11 ^ 127) {
                        case -363:
                            fmt.Println(fi11, v12, v2)
                        case 4294967295:
                            if (((0 > 0) || (((((((3.14159 - float32(-2.25)) - (float32(1e10) * 0.0)) * -2.25) * float32((0.1 + 3.14159))) * (0.1 * 1.5)) + 1e10) <= (float32(((0.1 * (1.5 / (1.5 + 0.0))) * 0.1)) + (0.1 + 0.1)))) || (23 <= 440)) {
                                if (true || false) {
                                    fmt.Println(fi10, v4, v5)
                                    v12 += v12
                                } else {
                                    v24 := GenS3{F0: (^(v5)), F1: (v4 / (((v4 | -2147483648)) | 1))}
                                    v5 = v5
                                    _ = v24
                                }
                            }
                        case 32767:
                            var v25 uint32 = v5
                            _ = v25
                        default:
                            var v26 uint8 = 128
                            _ = v26
                        }
                        _ = v23
                    }
                    _ = fi11
                    _ = v12
                }
            case 127:
                v27 := GenS1{F0: (v4 << ((0) & 63)), F1: 573, F2: 36}
                _ = v27
            case 32767:
                for fi28 := 0; fi28 < 16; fi28++ {
                    v5 = v5
                    var v29 uint32 = (v5 / (((v5 + v5)) | 1))
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi10
        }
    }
    if (v4 > v4) {
        fmt.Println(v5, v2, v5)
        fmt.Println(v4, v4, v5)
        v2 += (v2 - 74)
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            v2 = -122
            _ = fi31
        }
        switch -526 {
        case 9223372036854775807:
            if (!(2 > (^(v5)))) {
                for fi32 := 0; fi32 < 3; fi32++ {
                    v33 := GenS1{F0: v4, F1: 675, F2: 1}
                    _ = fi32
                    _ = v33
                }
            }
        case 32768:
            for fi34 := 0; fi34 < 3; fi34++ {
                v4 = (v4 + v4)
                _ = fi34
            }
        }
        v5 = (v5 / ((4294967295) | 1))
        _ = fi30
    }
    _ = len(v3)
    for fi35 := 0; fi35 < 9; fi35++ {
        v2 = v2
        _ = fi35
    }
    v2 = v2
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
