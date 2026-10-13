package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint8; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32 }

func main() {
    v1 := []uint32{0}
    var v2 float64 = float64(0.0)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    if len(v1) > 0 {
        v1[0] = 2
    }
    v5 := GenS1{F0: 1, F1: 216, F2: 1}
    var v6 uint8 = 255
    v7 := GenS1{F0: v6, F1: 32767, F2: -1}
    var v8 uint16 = 2
    delete(v4, 2)
    v6 = 42
    var v9 string = "foo bar"
    if (0 >= -128) {
        fmt.Println(v9)
        v10 := GenS2{F0: uint(v6), F1: -1}
        _ = v10
    }
    delete(v4, 2)
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS3{F0: 9223372036854775807}
        if (v2 != (1e10 / v2)) {
            if true {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS4{F0: 127, F1: v6, F2: 32767}
                    if false {
                        v2 = v2
                        var v15 float32 = (((1e10 / (0.1 + 1.5)) - 1e10) * (((0.0 - 0.0) - -2.25) - 0.1))
                        _ = v15
                    } else {
                        switch fi11 {
                        case 1:
                            switch (fi13 % ((fi13) | 1)) {
                            case 32767:
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                            default:
                                if ((true || true) && (318 > 1)) {
                                    v2 += v2
                                    v16 := GenS5{F0: "foo bar"}
                                    _ = v16
                                }
                            }
                        case -404:
                            var v17 int64 = 42
                            _ = v17
                        case 2147483647:
                            delete(v4, 4294967295)
                            v1 = append(v1, 0)
                            fmt.Println(fi13, v9, fi11)
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                fmt.Println(fi11)
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                if true {
                    switch "hello" {
                    case "a":
                        v19 := GenS4{F0: 128, F1: (v6 | v6), F2: 629}
                        _ = v19
                    }
                }
                if false {
                    var v20 float64 = v2
                    if ((729 != 2147483647) && true) {
                        var v21 int32 = 42
                        var v22 int32 = v21
                        _ = v21
                        _ = v22
                    }
                    v6 = (v6 - v6)
                    var v23 int32 = -1
                    v24 := GenS1{F0: v6, F1: int16(v6), F2: 1}
                    v23 -= (v23 + (v23 << ((uint(v6)) & 63)))
                    _ = v20
                    _ = v23
                    _ = v24
                }
                var v25 uint8 = 42
                _ = fi18
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := GenS6{F0: -1}
            _ = fi26
            _ = v27
        }
        _ = fi11
        _ = v12
    }
    v1 = append(v1, 637)
    var v28 int8 = -128
    v29 := v4[4294967295]
    var v30 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
