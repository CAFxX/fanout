package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 string }

func main() {
    var v1 uint64 = 1
    v2 := []uint32{2, 1, 118}
    var v3 uint8 = 128
    v4 := make(map[string]bool)
    switch 0 {
    case 0:
        if true {
            var v5 int16 = -32768
            v3 -= v3
            _ = v5
        }
    case 4294967295:
        delete(v4, "hello")
    default:
        if false {
            _ = len(v4)
            v6 := GenS1{F0: 0, F1: false, F2: 9223372036854775807}
            var v7 int32 = 0
            _ = v6
            _ = v7
        } else {
            var v8 uint64 = v1
            if false {
                if (!true) {
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    v9 := GenS2{F0: 1, F1: 1, F2: "a"}
                    switch 2147483647 {
                    case 689:
                        switch -889 {
                        case 9223372036854775807:
                            v8 = (v8 / (((v8 >> ((0) & 63))) | 1))
                        case -9223372036854775808:
                            for fi10 := 0; fi10 < 17; fi10++ {
                                if false {
                                    var v11 uint = 212
                                    _ = v11
                                }
                                var v12 float64 = 3.14159
                                v12 += (v12 + (v12 / v12))
                                _ = fi10
                                _ = v12
                            }
                        default:
                            fmt.Println(v1, v3, v8)
                        }
                    }
                    _ = v9
                } else {
                    if (false && false) {
                        for fi13 := 0; fi13 < 4; fi13++ {
                            var v14 int64 = 1
                            _ = fi13
                            _ = v14
                        }
                        switch "foo bar" {
                        case "":
                            for fi15 := 0; fi15 < 16; fi15++ {
                                var v16 int16 = 419
                                _ = len(v2)
                                var v17 uint8 = v3
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        case "\n\t":
                            v18 := GenS1{F0: 181, F1: false, F2: -9223372036854775808}
                            _ = v18
                        default:
                            v19 := GenS2{F0: 214, F1: 588, F2: ""}
                            _ = v19
                        }
                    }
                    v8 = v1
                    v3 -= (v3 ^ (v3 - v3))
                }
                if (489 <= 42) {
                    v2 = append(v2, 174)
                }
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 float64 = 0.0
                var v22 bool = (-1 != 127)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v8
        }
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            v1 = 4294967295
            fmt.Println(v1, v3)
        }
    } else {
        fmt.Println(v1, v3)
    }
    v1 += v1
    switch -873 {
    case 842:
        if true {
            for fi23 := 0; fi23 < 6; fi23++ {
                if (0 != 0) {
                    switch (^(fi23)) {
                    case 1:
                        v4[("\n\t" + ("\n\t" + "hello"))] = false
                    default:
                        v3 = 2
                    }
                    var v24 uint8 = 42
                    _ = v24
                } else {
                    var v25 int32 = -1
                    _ = v25
                }
                _ = fi23
            }
        }
    case 646:
        fmt.Println(v3)
    default:
        var v26 int64 = -950
        _ = v26
    }
    if (true && (!false)) {
        var v27 uint16 = 65535
        _ = v27
    } else {
        v1 += (v1 ^ (^(v1)))
        var v28 float64 = (1.5 / 0.1)
        for fi29 := 0; fi29 < 10; fi29++ {
            v28 -= (v28 - (v28 * v28))
            if false {
                fmt.Println(v28, v3, fi29)
                _ = len(v4)
            } else {
                fmt.Println(v3)
            }
            switch "hello" {
            case "hello":
                var v30 bool = true
                _ = v30
            case "\n\t":
                v3 = (^(v3))
            }
            _ = fi29
        }
        _ = v28
    }
    fmt.Println(v1, v3)
    if len(v2) > 0 {
        v2[0] = 42
    }
    v31 := GenS1{F0: 1, F1: true, F2: 664}
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v32 string = "\n\t"
    v32 = v32
    fmt.Println(v32, v3, v3)
    var v33 int64 = -316
    v1 = (v1 << ((4294967295) & 63))
    var v34 int32 = 42
    fmt.Println(v33, v32, v3)
    v3 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
