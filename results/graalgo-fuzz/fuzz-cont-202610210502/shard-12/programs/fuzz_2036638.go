package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 int; F2 int32 }
func genFn0() uint8 {
    var v7 uint32 = 313
    v7 += (^(v7))
    _ = v7
    return 155
}
func genFn1(p0 int8, p1 bool) uint8 {
    var v8 bool = (("\n\t" + "foo bar") <= "\n\t")
    v9 := GenS1{F0: -1, F1: 2, F2: ("hello" + "foo bar")}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int64 = 0
    var v2 uint8 = 2
    var v3 int16 = 0
    var v4 float32 = 1e10
    var v5 uint32 = 1
    v6 := []int{-9223372036854775808, 2147483647, 127, 4294967295}
    if false {
        if true {
            v2 = (^(v2))
            v2 = 1
        } else {
            if (1 <= 314) {
                if len(v6) > 0 {
                    v6[0] = int(v3)
                }
                if true {
                    if (true || true) {
                        if ((2147483647 < 2147483647) && false) {
                            fmt.Println(v1, v3)
                        }
                        for fi10 := 0; fi10 < 5; fi10++ {
                            if false {
                                v11 := GenS2{F0: 58, F1: (!true)}
                                v12 := GenS1{F0: (v3 - v3), F1: 32767, F2: "a"}
                                fmt.Println(v5, fi10, v4)
                                _ = v11
                                _ = v12
                            }
                            for fi13 := 0; fi13 < 13; fi13++ {
                                var v14 bool = true
                                var v15 int16 = (v3 << ((255) & 63))
                                _ = len(v6)
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 10; fi16++ {
                                var v17 string = ("foo bar" + "foo bar")
                                switch v17 {
                                case "\n\t":
                                    v18 := GenS3{F0: -2147483648, F1: v2}
                                    v5 -= v5
                                    _ = v18
                                default:
                                    v17 = "foo bar"
                                }
                                v5 = v5
                                _ = fi16
                                _ = v17
                            }
                            _ = fi10
                        }
                    }
                } else {
                    if (!true) {
                        v19 := GenS1{F0: v3, F1: 1, F2: "\n\t"}
                        _ = v19
                    }
                    if len(v6) > 0 {
                        v6[0] = 42
                    }
                    if true {
                        fmt.Println(v2, v4)
                        v3 += v3
                    }
                }
                var v20 int64 = (v1 & 42)
                _ = v20
            } else {
                v21 := GenS4{F0: 1}
                _ = v21
            }
        }
    }
    v22 := GenS3{F0: 1, F1: v2}
    v23 := GenS2{F0: -1, F1: false}
    fmt.Println(v3, v1, v1)
    var v24 uint64 = 433
    for fi25 := 0; fi25 < 13; fi25++ {
        v24 = (v24 >> ((uint(v24)) & 63))
        _ = fi25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(v4)
        var v27 int = fi26
        _ = fi26
        _ = v27
    }
    v6 = append(v6, -1)
    if (65535 < 2) {
        switch "hello" {
        case "":
            var v28 bool = (24 <= 656)
            _ = v28
        default:
            for fi29 := 0; fi29 < 15; fi29++ {
                if false {
                    fmt.Println(v2, v5)
                } else {
                    v4 += v4
                }
                v30 := GenS5{F0: v1, F1: fi29, F2: 1}
                fmt.Println(v24, v5, v5)
                _ = fi29
                _ = v30
            }
        }
    } else {
        var v31 uint8 = (v2 >> ((0) & 63))
        v24 -= uint64(v5)
        v4 = v4
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
