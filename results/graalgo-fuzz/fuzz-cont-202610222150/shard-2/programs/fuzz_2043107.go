package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint }
func genFn0() int32 {
    var v6 uint64 = 4294967295
    fmt.Println(v6)
    var v7 string = "hello"
    v6 = v6
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 int16 = 42
    v4 := make(map[uint32]int)
    v5 := make(map[string]string)
    v3 += (v3 << ((537) & 63))
    var v8 int64 = 9223372036854775807
    v1[42] = 0
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v3, v8, v8)
        v3 = (v3 / (((v3 % ((-1) | 1))) | 1))
        _ = fi9
    }
    v10 := []int{}
    fmt.Println(v3)
    v11 := GenS1{F0: 652}
    if false {
        var v12 int32 = -246
        _ = v12
    }
    if (!true) {
        if (-1 < (v8 ^ v8)) {
            v8 += -1
            for fi13 := 0; fi13 < 4; fi13++ {
                if false {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v8 -= v8
                        fmt.Println(fi14, fi13, fi14)
                        switch -1 {
                        case -67:
                            var v15 int = fi14
                            _ = v15
                        case 42:
                            var v16 string = "hello"
                            _ = v16
                        case 413:
                            var v17 int8 = 38
                            _ = v17
                        }
                        _ = fi14
                    }
                }
                _ = fi13
            }
            v18 := GenS1{F0: 255}
            _ = v18
        }
        delete(v2, "a")
        var v19 int32 = genFn0()
        _ = v19
    } else {
        if len(v10) > 0 {
            v10[0] = 9223372036854775807
        }
    }
    if (!true) {
        v4[42] = 9223372036854775807
        v20 := GenS1{F0: 0}
        var v21 bool = true
        _ = v20
        _ = v21
    } else {
        var v22 float64 = ((3.14159 / 0.1) / (-2.25 + (0.0 - float64((1.5 - ((1.5 + 1e10) - (0.1 / 1e10)))))))
        _ = v22
    }
    fmt.Println(v3, v3)
    v8 = v8
    switch 0 {
    case -931:
        delete(v1, 4294967295)
    case 32767:
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int = 42
            v10 = append(v10, (^(v24)))
            _ = fi23
            _ = v24
        }
    case 9223372036854775807:
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS1{F0: 42}
            var v27 float32 = (((3.14159 / 1.5) - 1e10) + (0.0 + 1e10))
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v28 := GenS2{F0: (v3 / ((1) | 1)), F1: -16, F2: 416}
    var v29 int64 = (v8 + (v8 >> ((361) & 63)))
    v3 -= -367
    v30 := GenS3{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
