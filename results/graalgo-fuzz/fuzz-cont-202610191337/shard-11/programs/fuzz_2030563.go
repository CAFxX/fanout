package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 bool }
type GenS4 struct { F0 int64; F1 uint8; F2 int }
type GenS5 struct { F0 int64; F1 uint64 }

func main() {
    v1 := []uint32{421, 42, 397, 42}
    var v2 bool = true
    var v3 uint8 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = v3
    if ((v3 ^ (v3 ^ 203)) > (^(v3))) {
        for fi4 := 0; fi4 < 10; fi4++ {
            v5 := GenS1{F0: fi4}
            _ = fi4
            _ = v5
        }
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS2{F0: 255}
            _ = fi6
            _ = v7
        }
        v3 = (v3 | v3)
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint8 = (v3 * (v3 ^ (v3 & 2)))
            fmt.Println(v9, v2)
            _ = fi8
            _ = v9
        }
        if (491 >= 1) {
            var v10 float32 = 1.5
            if v2 {
                v11 := GenS3{F0: v3, F1: 18446744073709551615, F2: v2}
                if (v3 != (v3 % (((^(v3))) | 1))) {
                    var v12 uint8 = v3
                    if (!(((0.0 + 0.1) == float32(0.1)) || (65535 < 4294967295))) {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            var v14 uint16 = 255
                            var v15 uint8 = (v3 - (v3 - (v12 & 130)))
                            for fi16 := 0; fi16 < 15; fi16++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi16
                            }
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 = true
                    }
                    for fi17 := 0; fi17 < 19; fi17++ {
                        fmt.Println(fi17)
                        v1 = append(v1, uint32(v12))
                        _ = fi17
                    }
                    _ = v12
                } else {
                    var v18 float32 = v10
                    var v19 int8 = 1
                    _ = v18
                    _ = v19
                }
                switch "a" {
                case "a":
                    v1 = append(v1, 949)
                case "foo bar":
                    fmt.Println(v10)
                case "hello":
                    v3 = (v3 + (v3 * v3))
                default:
                    var v20 int = -879
                    _ = v20
                }
                _ = v11
            }
            fmt.Println(v10, v3, v10)
            _ = v10
        } else {
            if (227 <= 56) {
                v21 := GenS3{F0: (^(v3)), F1: 1, F2: (v2 || v2)}
                _ = v21
            } else {
                var v22 uint8 = (^(v3))
                _ = v22
            }
            fmt.Println(v3, v2)
        }
        v23 := GenS4{F0: -1, F1: 166, F2: 128}
        _ = v23
    }
    if v2 {
        fmt.Println(v3, v3)
    } else {
        if (65535 > 1) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = v3
            var v24 int8 = -16
            _ = v24
        } else {
            switch ((("" + "a") + "hello") + "hello") {
            case "\n\t":
                var v25 uint = 2
                _ = v25
            case "hello":
                v26 := GenS5{F0: -103, F1: 42}
                _ = v26
            default:
                v2 = (!v2)
            }
            var v27 float64 = -2.25
            _ = v27
        }
    }
    if false {
        fmt.Println(v3, v2)
    }
    fmt.Println(v3)
    switch 524 {
    case 10:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v28 := GenS4{F0: 1, F1: 255, F2: 0}
        _ = v28
    }
    var v29 string = "a"
    if true {
        var v30 uint16 = 0
        v1 = append(v1, 0)
        v31 := GenS5{F0: -253, F1: 0}
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        fmt.Println(fi32)
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
}
