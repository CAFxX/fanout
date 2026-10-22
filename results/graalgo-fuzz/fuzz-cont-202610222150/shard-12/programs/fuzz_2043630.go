package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 int16 }
func genFn0() bool {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int16 = 0
        _ = fi5
        _ = v6
    }
    var v7 int32 = -696
    var v8 bool = ((false && false) && false)
    var v9 int64 = 0
    v9 = v9
    _ = v7
    _ = v8
    _ = v9
    return (-108 < 1)
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    var v3 int64 = 898
    v4 := []uint32{242, 1, 65535, 1}
    var v10 int32 = 0
    var v11 int32 = 1
    _ = len(v1)
    var v12 uint64 = 2
    v3 += v3
    var v13 uint32 = 1
    var v14 uint16 = 255
    v15 := GenS1{F0: 561}
    if false {
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 1; fi17++ {
                if false {
                    if (!(254 < 255)) {
                        delete(v1, (-(fi16)))
                    } else {
                        if (v13 != 70) {
                            fmt.Println(fi16, v13)
                        } else {
                            switch "" {
                            case "hello":
                                var v18 string = ("" + "")
                                _ = v18
                            case "\n\t":
                                var v19 int32 = (v10 % ((v11) | 1))
                                _ = v19
                            default:
                                v10 += (-(v10))
                            }
                            switch 491 {
                            case 2147483647:
                                if (!((-(fi17)) > -9223372036854775808)) {
                                    if (((!true) || (489 > 1000)) && false) {
                                        if len(v4) > 0 {
                                            v4[0] = (v13 + 1)
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        v20 := v1[(fi17 % (((fi16 << ((0) & 63))) | 1))]
                                        if ((!false) && true) {
                                            if true {
                                                v21 := GenS2{F0: -1, F1: 100}
                                                if ((v14 * (v14 * (v14 ^ v14))) == 0) {
                                                    v13 = v13
                                                    v14 = (v14 & 248)
                                                }
                                                _ = v21
                                            } else {
                                                var v22 uint32 = (v13 + v13)
                                                _ = v22
                                            }
                                            v2 = append(v2, v20)
                                            var v23 float64 = 0.0
                                            _ = v23
                                        }
                                        var v24 int8 = -51
                                        _ = v20
                                        _ = v24
                                    }
                                }
                            }
                        }
                        v10 += (v11 - v11)
                        var v25 int8 = -128
                        _ = v25
                    }
                }
                _ = fi17
            }
            if (!((false && (true || (v14 <= (v14 * 42)))) || false)) {
                var v26 int64 = (v3 ^ (v3 ^ v3))
                v10 = (v10 & v11)
                _ = v26
            } else {
                v27 := GenS1{F0: 42}
                var v28 float32 = (1.5 / (0.1 / 3.14159))
                v10 = 222
                _ = v27
                _ = v28
            }
            _ = fi16
        }
    } else {
        fmt.Println(v11, v14, v10)
        switch 2147483647 {
        case 1:
            var v29 uint8 = 220
            _ = v29
        default:
            if false {
                v30 := GenS2{F0: 1, F1: 127}
                for fi31 := 0; fi31 < 14; fi31++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v12 = (^(v12))
                    if true {
                        v32 := GenS2{F0: -68, F1: 109}
                        var v33 float64 = ((0.0 / 0.1) * 1e10)
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int = 9223372036854775807
                        if true {
                            var v35 int64 = (v3 * 1)
                            _ = v35
                        }
                        _ = v34
                    }
                    _ = fi31
                }
                _ = v30
            } else {
                for fi36 := 0; fi36 < 7; fi36++ {
                    var v37 uint32 = 1
                    _ = fi36
                    _ = v37
                }
                var v38 uint = 621
                var v39 uint64 = v12
                _ = v38
                _ = v39
            }
        }
        if true {
            var v40 bool = false
            for fi41 := 0; fi41 < 10; fi41++ {
                if ((!(-2.25 != 0.0)) && v40) {
                    v42 := GenS3{F0: uint8(v14), F1: v11, F2: 0}
                    _ = v42
                } else {
                    v14 += (^(v14))
                    fmt.Println(v10, v13)
                    fmt.Println(fi41)
                }
                switch "foo bar" {
                case "a":
                    delete(v1, fi41)
                case "foo bar":
                    v11 += (v11 % ((-363) | 1))
                case "":
                    v43 := GenS2{F0: 16, F1: 127}
                    _ = v43
                }
                _ = fi41
            }
            _ = v40
        }
    }
    v1[32767] = "hello"
    fmt.Println(v10)
    v44 := GenS3{F0: 2, F1: v10, F2: 42}
    for fi45 := 0; fi45 < 8; fi45++ {
        v13 = (v13 | v13)
        _ = fi45
    }
    v46 := GenS2{F0: 42, F1: int8(v12)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
}
