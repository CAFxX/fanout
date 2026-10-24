package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 int64) uint16 {
    var v5 bool = true
    v6 := GenS1{F0: (p1 << ((2) & 63))}
    var v7 float64 = 0.0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint16 = 10
    v2 := make(map[string]uint32)
    v3 := []uint32{0, 0, 4294967295}
    v4 := make(map[string]int)
    var v8 int64 = 307
    v8 = v8
    for fi9 := 0; fi9 < 6; fi9++ {
        if (0 != 1) {
            v10 := GenS2{F0: 255, F1: 31, F2: 623}
            var v11 float32 = 0.0
            _ = v10
            _ = v11
        } else {
            fmt.Println(fi9)
        }
        if true {
            var v12 float64 = 0.1
            _ = v12
        } else {
            var v13 int64 = 9223372036854775807
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(fi9)
                fmt.Println(fi9, v13)
                _ = fi14
            }
            switch (fi9 ^ (fi9 / ((fi9) | 1))) {
            case 9223372036854775807:
                v1 -= v1
            default:
                var v15 uint64 = 4294967295
                _ = v15
            }
            _ = v13
        }
        v16 := GenS2{F0: 128, F1: 1, F2: genFn0(162, (^(v8)))}
        _ = fi9
        _ = v16
    }
    fmt.Println(v1, v1)
    v1 = v1
    var v17 uint = 42
    var v18 int64 = v8
    if true {
        v19 := GenS3{F0: 514}
        _ = v19
    } else {
        v1 += genFn0(2, v8)
    }
    switch 42 {
    case 9223372036854775807:
        v3 = append(v3, 35)
    case 32767:
        v20 := GenS3{F0: 2}
        _ = v20
    case 78:
        if false {
            v8 = v18
            if (!((127 == 0) || false)) {
                var v21 uint8 = uint8(v17)
                v18 += (v18 * v8)
                _ = v21
            } else {
                v22 := GenS2{F0: 1, F1: 2, F2: genFn0(2, v8)}
                _ = len(v2)
                v8 = v18
                _ = v22
            }
        } else {
            fmt.Println(v8)
            switch ("a" + (("foo bar" + "hello") + ("foo bar" + "foo bar"))) {
            case "\n\t":
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 int32 = 42
                    var v25 uint64 = 2
                    var v26 int32 = v24
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            case "":
                var v27 int8 = 42
                _ = v27
            default:
                if (0.0 == 1e10) {
                    for fi28 := 0; fi28 < 15; fi28++ {
                        switch ("" + "foo bar") {
                        case "hello":
                            if (265 >= 69) {
                                fmt.Println(v18, v8)
                            } else {
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    var v30 uint32 = 0
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = (^(v30))
                                        }
                                        var v31 uint32 = v30
                                        _ = v31
                                    } else {
                                        v8 = v18
                                        v32 := GenS4{F0: 65535, F1: v1}
                                        _ = v32
                                    }
                                    v30 += v30
                                    _ = fi29
                                    _ = v30
                                }
                                v4["foo bar"] = fi28
                                fmt.Println(fi28, v18)
                            }
                        case "":
                            var v33 uint32 = 209
                            _ = v33
                        default:
                            var v34 uint32 = 2147483648
                            _ = v34
                        }
                        v1 = v1
                        v1 -= (v1 + 0)
                        _ = fi28
                    }
                }
            }
        }
    }
    switch "foo bar" {
    case "hello":
        var v35 int32 = 2147483647
        _ = v35
    case "\n\t":
        v36 := GenS2{F0: 131, F1: uint8(v1), F2: uint16(v17)}
        _ = v36
    case "foo bar":
        _ = len(v2)
    default:
        for fi37 := 0; fi37 < 19; fi37++ {
            v38 := GenS3{F0: uint(v1)}
            _ = fi37
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
