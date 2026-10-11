package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
func genFn0() uint16 {
    var v5 int8 = 100
    switch ("hello" + "hello") {
    case "":
        if true {
            v6 := []int{4294967295, 2147483647, -1, 4294967295}
            var v7 uint64 = 18446744073709551615
            if false {
                var v8 uint8 = 42
                _ = v8
            }
            _ = v6
            _ = v7
        }
    case "\n\t":
        fmt.Println(v5)
    default:
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(v5, v5)
            fmt.Println(fi9)
            v10 := GenS1{F0: 1000, F1: (fi9 ^ (fi9 | (fi9 >> ((2) & 63)))), F2: fi9}
            _ = fi9
            _ = v10
        }
    }
    if false {
        var v11 int = -961
        v12 := GenS1{F0: 32767, F1: (v11 + v11), F2: v11}
        fmt.Println(v11)
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := []int{}
            _ = len(v14)
            _ = fi13
            _ = v14
        }
    }
    _ = v5
    return 2
}

func main() {
    v1 := []string{("a" + "\n\t"), ("\n\t" + "a"), ("hello" + "foo bar"), "\n\t"}
    var v2 uint = 18446744073709551615
    var v3 int8 = 0
    v4 := make(map[uint32]uint32)
    var v15 int16 = -32768
    var v16 uint32 = 865
    if (9223372036854775807 > 42) {
        switch "\n\t" {
        case "\n\t":
            v3 += (v3 / (((v3 % ((1) | 1))) | 1))
        case "":
            if false {
                v17 := GenS2{F0: (v15 << (((v2 & v2)) & 63)), F1: (v3 & (v3 << (((v2 / ((220) | 1))) & 63))), F2: true}
                _ = v17
            } else {
                if (!(128 == 155)) {
                    v18 := GenS1{F0: -1, F1: 40, F2: -2147483648}
                    v15 = (v15 + (v15 / ((v15) | 1)))
                    switch 614 {
                    case 127:
                        switch -210 {
                        case 0:
                            if (!true) {
                                v3 = (v3 | int8(v15))
                                var v19 int32 = 0
                                _ = v19
                            }
                        case 747:
                            if true {
                                v1 = append(v1, "")
                                v3 = v3
                            }
                        case 42:
                            var v20 uint = 65535
                            _ = v20
                        }
                    case 416:
                        var v21 float32 = 0.1
                        _ = v21
                    }
                    _ = v18
                }
                var v22 int32 = -1
                v15 = (v15 << ((v2) & 63))
                _ = v22
            }
        default:
            v16 = (^(v16))
        }
    }
    switch -948 {
    case -1:
        for fi23 := 0; fi23 < 6; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 uint = v2
                _ = fi24
                _ = v25
            }
            var v26 int64 = 1
            _ = fi23
            _ = v26
        }
    case 9223372036854775807:
        fmt.Println(v3)
    case 986:
        v15 -= v15
    default:
        fmt.Println(v15, v2, v15)
    }
    v3 = v3
    var v27 uint = 65535
    v27 = (^(v27))
    if (19 > 127) {
        v28 := GenS2{F0: v15, F1: -46, F2: (9223372036854775807 > -1)}
        if (((0.0 + 0.0) - 1.5) < 0.1) {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 uint8 = 166
                v31 := GenS1{F0: (v15 % ((42) | 1)), F1: fi29, F2: (fi29 * fi29)}
                _ = fi29
                _ = v30
                _ = v31
            }
            v3 = v3
            var v32 int32 = 388
            _ = v32
        } else {
            var v33 uint32 = 1
            _ = v33
        }
        fmt.Println(v3)
        _ = v28
    }
    v34 := GenS2{F0: (v15 | v15), F1: v3, F2: false}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v35 := GenS1{F0: (v15 >> ((752) & 63)), F1: 9223372036854775807, F2: -2147483648}
    var v36 int32 = 430
    var v37 uint8 = 0
    v3 -= -128
    var v38 uint8 = 128
    var v39 int32 = (v36 ^ v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
