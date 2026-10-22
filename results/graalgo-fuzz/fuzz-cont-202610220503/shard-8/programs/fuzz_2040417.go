package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 bool }
func genFn0() bool {
    var v7 string = "a"
    fmt.Println(v7)
    switch v7 {
    case "foo bar":
        if true {
            v7 = ("a" + v7)
        }
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: v7, F1: 1, F2: 32767}
        var v10 uint8 = 128
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return (458 <= 146)
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint16 = 346
    var v3 string = "hello"
    var v4 uint8 = 42
    v5 := make(map[string]uint32)
    var v6 uint8 = v4
    for fi11 := 0; fi11 < 4; fi11++ {
        _ = len(v1)
        switch v3 {
        case "":
            v3 = ""
        case "hello":
            var v12 int8 = 42
            _ = v12
        case "\n\t":
            var v13 uint32 = 2
            _ = v13
        }
        var v14 bool = false
        _ = fi11
        _ = v14
    }
    delete(v5, (v3 + v3))
    for fi15 := 0; fi15 < 19; fi15++ {
        if true {
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS2{F0: int64(fi15), F1: (v6 - (^(v4))), F2: 1}
                v18 := GenS1{F0: (v3 + ("" + "")), F1: v2, F2: (v2 + (v2 / (((v2 & (^(v2)))) | 1)))}
                var v19 uint64 = 1
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            if true {
                v6 = v4
                var v20 string = "a"
                _ = v20
            }
            switch fi15 {
            case 127:
                switch fi15 {
                case 252:
                    v3 = (v3 + "foo bar")
                case 791:
                    var v21 bool = false
                    _ = v21
                case 128:
                    v22 := GenS2{F0: int64(v6), F1: (v4 << ((0) & 63)), F2: -128}
                    _ = v22
                default:
                    if (false && false) {
                        v23 := GenS2{F0: 42, F1: v6, F2: 100}
                        _ = v23
                    } else {
                        v6 = (v6 / ((v4) | 1))
                    }
                }
            case 507:
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 int = fi24
                    _ = fi24
                    _ = v25
                }
            default:
                var v26 float32 = (float32(3.14159) * 0.1)
                _ = v26
            }
        }
        if (false && (-1 != 74)) {
            if (!((1000 < 32767) && true)) {
                switch (^(fi15)) {
                case 32767:
                    if (false && true) {
                        v27 := GenS1{F0: v3, F1: (v2 >> ((78) & 63)), F2: v2}
                        switch 4294967295 {
                        case 1:
                            v28 := GenS3{F0: v3, F1: 559}
                            _ = v28
                        case -50:
                            var v29 float32 = 1e10
                            _ = v29
                        case 4294967295:
                            v2 = (^(v2))
                        }
                        _ = v27
                    }
                }
                for fi30 := 0; fi30 < 16; fi30++ {
                    v31 := GenS1{F0: ("a" + v3), F1: v2, F2: v2}
                    switch v3 {
                    case "":
                        var v32 int = (fi15 % (((fi15 - fi15)) | 1))
                        _ = v32
                    case "foo bar":
                        v1 = append(v1, 0)
                    case "hello":
                        fmt.Println(v3, v4, fi15)
                    default:
                        _ = len(v1)
                    }
                    v3 = v3
                    _ = fi30
                    _ = v31
                }
            } else {
                switch (fi15 - fi15) {
                case 483:
                    v33 := GenS1{F0: v3, F1: (^(v2)), F2: 787}
                    _ = v33
                case -2147483648:
                    v3 = (v3 + ((v3 + (((((v3 + "\n\t") + v3) + v3) + ("a" + "foo bar")) + v3)) + v3))
                case 128:
                    for fi34 := 0; fi34 < 14; fi34++ {
                        var v35 uint16 = v2
                        delete(v5, (v3 + (((("" + "") + "\n\t") + ("hello" + (v3 + "\n\t"))) + "foo bar")))
                        var v36 uint32 = uint32(v6)
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                default:
                    var v37 int8 = -128
                    _ = v37
                }
            }
            if (!false) {
                v6 -= 128
                var v38 uint = 0
                _ = v38
            } else {
                if ((!(false && (1.5 <= 1.5))) && false) {
                    v2 -= (v2 + v2)
                    switch "hello" {
                    case "hello":
                        var v39 int32 = -69
                        _ = v39
                    }
                    v40 := GenS1{F0: v3, F1: (v2 ^ (v2 * v2)), F2: v2}
                    _ = v40
                } else {
                    if (v3 != v3) {
                        v41 := GenS2{F0: -1, F1: v4, F2: 0}
                        v1 = append(v1, 4294967295)
                        _ = v41
                    } else {
                        v2 -= (v2 & (^(v2)))
                        v2 = v2
                    }
                    if ((65535 <= 2) && (2 != 219)) {
                        var v42 uint32 = 2147483647
                        v6 = v4
                        _ = v42
                    }
                }
            }
        }
        v2 = 2
        _ = fi15
    }
    v43 := GenS4{F0: 0, F1: 403, F2: ((v2 >> ((255) & 63)) >= (v2 % ((65535) | 1)))}
    var v44 uint32 = 875
    if ((((128 < (^(v4))) && ((-9223372036854775808 <= 761) || (v3 >= "\n\t"))) && false) || true) {
        v4 = uint8(v2)
    }
    var v45 int16 = -32768
    if ((-9223372036854775808 != -679) || true) {
        v6 -= (v6 | v6)
        v46 := v5[v3]
        _ = v46
    } else {
        v4 += v4
    }
    v44 = v44
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
