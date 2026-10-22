package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 int16 }
type GenS5 struct { F0 int64; F1 string; F2 int16 }
func genFn0() int {
    if true {
        var v5 uint64 = 0
        var v6 float32 = (0.0 * ((1.5 - ((1.5 * 0.0) / (((((1.5 - (float32(1e10) - (3.14159 - -2.25))) + (0.0 / -2.25)) / (3.14159 - (1e10 / 1.0))) / 3.14159) + (-2.25 * (float32(1.5) - (1e10 / float32(0.1))))))) * 1e10))
        v6 += (v6 / v6)
        _ = v5
        _ = v6
    } else {
        v7 := []int{32767, 42}
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    if true {
        if (!false) {
            var v10 string = "a"
            _ = v10
        }
        var v11 int8 = -128
        _ = v11
    }
    v12 := GenS1{F0: 32767, F1: 4294967295}
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int = (fi13 >> ((943) & 63))
        v14 = 9223372036854775807
        if ((fi13 / ((fi13) | 1)) != fi13) {
            var v15 uint64 = 135
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    _ = v12
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    v3 := []uint32{549, 42, 794, 777}
    var v4 uint8 = 0
    v4 = v4
    v4 = v4
    for fi16 := 0; fi16 < 8; fi16++ {
        v4 = (v4 & v4)
        var v17 int64 = -710
        v18 := GenS1{F0: 0, F1: 4294967295}
        _ = fi16
        _ = v17
        _ = v18
    }
    if false {
        var v19 float32 = 1.5
        v4 -= v4
        _ = v19
    } else {
        v20 := GenS2{F0: 127, F1: 4294967295, F2: "\n\t"}
        var v21 string = "\n\t"
        _ = v20
        _ = v21
    }
    if ("a" >= ("" + "hello")) {
        var v22 uint = 2
        var v23 int32 = 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v22
        _ = v23
    } else {
        var v24 int32 = -1
        v4 = 2
        if (((!(true || true)) || (((("hello" + "\n\t") + "") + "a") <= "\n\t")) || true) {
            for fi25 := 0; fi25 < 7; fi25++ {
                v4 += v4
                _ = fi25
            }
            v2[-487] = "a"
        } else {
            var v26 uint16 = 65535
            _ = v26
        }
        _ = v24
    }
    if (true || true) {
        var v27 uint8 = (v4 / (((^(v4))) | 1))
        _ = v27
    }
    fmt.Println(v4)
    for fi28 := 0; fi28 < 3; fi28++ {
        if (1 > 582) {
            v29 := GenS3{F0: -32768, F1: 26}
            if ((^(fi28)) >= genFn0()) {
                if (1 == 255) {
                    v30 := GenS2{F0: -1, F1: 42, F2: "\n\t"}
                    v4 = 255
                    _ = v30
                } else {
                    fmt.Println(v4, fi28)
                    if true {
                        for fi31 := 0; fi31 < 18; fi31++ {
                            fmt.Println(v4, v4)
                            v4 += v4
                            _ = fi31
                        }
                        if true {
                            fmt.Println(fi28, fi28)
                            for fi32 := 0; fi32 < 10; fi32++ {
                                var v33 float64 = 0.1
                                var v34 int8 = -128
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            for fi35 := 0; fi35 < 3; fi35++ {
                                var v36 bool = false
                                _ = fi35
                                _ = v36
                            }
                        }
                    }
                }
            }
            if len(v3) > 0 {
                v3[0] = 0
            }
            _ = v29
        } else {
            v37 := GenS4{F0: 42, F1: (v4 * 178), F2: 788}
            _ = v37
        }
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 int8 = 73
            _ = fi38
            _ = v39
        }
        v40 := GenS5{F0: -1, F1: "\n\t", F2: 0}
        _ = fi28
        _ = v40
    }
    switch "hello" {
    case "hello":
        v41 := GenS5{F0: -823, F1: "a", F2: 0}
        _ = v41
    case "\n\t":
        v4 += 128
    case "":
        var v42 string = "\n\t"
        _ = v42
    default:
        v4 -= (v4 | (v4 % ((v4) | 1)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
