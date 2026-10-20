package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint }
type GenS3 struct { F0 bool; F1 uint32; F2 int64 }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }
type GenS5 struct { F0 string; F1 int16; F2 int32 }

func main() {
    var v1 int8 = -108
    var v2 string = ""
    var v3 string = "hello"
    v4 := []string{v2, ""}
    var v5 uint = 42
    var v6 float64 = (0.0 + 0.1)
    v1 += (v1 / ((v1) | 1))
    if (!(802 == 583)) {
        if true {
            v6 += 0.1
            var v7 uint16 = 32767
            _ = v7
        } else {
            for fi8 := 0; fi8 < 18; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 string = ("a" + (("hello" + v2) + v3))
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 int16 = int16(fi8)
                    if true {
                        fmt.Println(v3, fi8)
                        var v13 float32 = 1e10
                        var v14 int64 = 0
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 218, F1: (923 < 1), F2: (v1 + v1)}
                        _ = v15
                    }
                    _ = fi11
                    _ = v12
                }
                _ = fi8
            }
            v16 := GenS1{F0: 479, F1: (4294967295 > 696), F2: v1}
            if (float64(0.1) != -2.25) {
                if len(v4) > 0 {
                    v4[0] = ("hello" + (v3 + ((("\n\t" + "foo bar") + "foo bar") + "a")))
                }
                var v17 string = v3
                _ = v17
            } else {
                v18 := GenS2{F0: 0, F1: 821, F2: 439}
                v3 = "foo bar"
                var v19 int16 = 1
                _ = v18
                _ = v19
            }
            _ = v16
        }
        if (1 != 4294967295) {
            v20 := GenS2{F0: -32768, F1: 2, F2: 65535}
            var v21 int32 = 777
            _ = v20
            _ = v21
        } else {
            v1 = v1
            fmt.Println(v5, v6, v1)
            if (true || false) {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            } else {
                if (0.0 < 0.1) {
                    var v22 uint64 = 876
                    _ = v22
                }
                for fi23 := 0; fi23 < 6; fi23++ {
                    v2 = "hello"
                    var v24 uint32 = 900
                    v25 := GenS1{F0: -2147483648, F1: (false && true), F2: v1}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
    } else {
        if (42 != 0) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v26 := GenS3{F0: (!(true && true)), F1: 65535, F2: 9223372036854775807}
            var v27 float64 = (v6 + v6)
            _ = v26
            _ = v27
        }
    }
    var v28 int64 = -271
    _ = len(v4)
    v28 = (v28 / (((v28 & (v28 & v28))) | 1))
    var v29 float64 = v6
    switch 128 {
    case 1:
        if (994 != 42) {
            var v30 uint16 = 255
            v3 = "\n\t"
            _ = v30
        } else {
            _ = len(v4)
            for fi31 := 0; fi31 < 5; fi31++ {
                v32 := GenS4{F0: (v28 * v28), F1: (v5 + (^(v5))), F2: 1}
                _ = fi31
                _ = v32
            }
        }
    case -2147483648:
        v1 -= 127
    case 32768:
        v5 = 775
    }
    fmt.Println(v6)
    for fi33 := 0; fi33 < 8; fi33++ {
        if (160 != -1) {
            fmt.Println(v1)
            var v34 string = ""
            v3 = v2
            _ = v34
        }
        v4 = append(v4, v2)
        switch (^(fi33)) {
        case 127:
            v35 := GenS5{F0: v2, F1: -655, F2: 2147483647}
            _ = v35
        default:
            for fi36 := 0; fi36 < 2; fi36++ {
                var v37 uint32 = 381
                _ = len(v4)
                _ = len(v4)
                if (v28 <= (-(v28))) {
                    var v38 uint16 = 238
                    switch fi36 {
                    case 32767:
                        var v39 uint = (^(v5))
                        _ = v39
                    default:
                        var v40 uint = (v5 - (v5 / (((^(v5))) | 1)))
                        _ = v40
                    }
                    fmt.Println(v38, v1, fi33)
                    fmt.Println(v28, fi33)
                    _ = v38
                } else {
                    if (!(false && false)) {
                        var v41 int64 = v28
                        _ = v41
                    }
                    switch v2 {
                    case "hello":
                        v2 = v2
                    case "\n\t":
                        v42 := GenS5{F0: v2, F1: 42, F2: 725}
                        _ = v42
                    case "a":
                        switch fi33 {
                        case 127:
                            if false {
                                v3 = ("" + (v3 + v2))
                            }
                        case 9223372036854775807:
                            fmt.Println(v3, v6)
                        default:
                            v37 += (v37 / (((v37 & v37)) | 1))
                        }
                    }
                    var v43 uint32 = 42
                    _ = v43
                }
                v29 = 0.0
                _ = fi36
                _ = v37
            }
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
