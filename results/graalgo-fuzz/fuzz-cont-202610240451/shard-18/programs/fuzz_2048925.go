package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int16; F2 string }
func genFn0() uint64 {
    var v6 bool = (true || (!false))
    v7 := GenS1{F0: 1, F1: 410, F2: -32768}
    _ = v6
    _ = v7
    return 15
}
func genFn1(p0 uint64, p1 int64) string {
    fmt.Println(p1, p1)
    var v8 bool = false
    _ = p0
    _ = p1
    _ = v8
    return ""
}

func main() {
    var v1 int32 = -419
    v2 := make(map[string]int)
    var v3 float32 = 3.14159
    var v4 string = "foo bar"
    v5 := make(map[uint32]uint32)
    switch 641 {
    case 128:
        delete(v5, 42)
    case 42:
        v4 = v4
    case 2147483647:
        if (!(!(255 > 767))) {
            fmt.Println(v3, v3, v3)
            v9 := GenS2{F0: -1, F1: v4}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 9; fi11++ {
                    switch fi11 {
                    case 42:
                        var v12 uint16 = 42
                        _ = v12
                    }
                    _ = fi11
                }
                _ = fi10
            }
        }
    }
    if ((-386 != 42) && (730 != 936)) {
        var v13 int16 = 442
        if (519 == 9223372036854775807) {
            if (((v13 % ((v13) | 1)) != v13) || ((!true) && (!(1 >= 0)))) {
                v14 := GenS3{F0: 100, F1: 1}
                v4 = v4
                for fi15 := 0; fi15 < 15; fi15++ {
                    fmt.Println(v13, v13)
                    for fi16 := 0; fi16 < 18; fi16++ {
                        fmt.Println(v4, fi15)
                        v17 := GenS2{F0: 1, F1: (v4 + v4)}
                        _ = fi16
                        _ = v17
                    }
                    var v18 uint8 = 210
                    _ = fi15
                    _ = v18
                }
                _ = v14
            } else {
                v1 = (v1 << ((4294967295) & 63))
                if true {
                    v19 := GenS3{F0: 1, F1: 2}
                    v2[""] = -1
                    var v20 float32 = v3
                    _ = v19
                    _ = v20
                } else {
                    v1 -= v1
                    var v21 int32 = (v1 - (v1 >> ((1) & 63)))
                    _ = v21
                }
                switch int(v1) {
                case 0:
                    v22 := GenS4{F0: 255, F1: 211, F2: 42}
                    _ = v22
                case 807:
                    v13 = v13
                default:
                    v4 = (("a" + "hello") + v4)
                }
            }
            v23 := GenS1{F0: 255, F1: -9223372036854775808, F2: 1}
            for fi24 := 0; fi24 < 19; fi24++ {
                v25 := GenS2{F0: 9223372036854775807, F1: v4}
                var v26 int64 = -275
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        }
        _ = v13
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 int = (^(fi27))
        _ = fi27
        _ = v28
    }
    v3 = (v3 - (v3 / ((float32(0.1) - (0.0 - v3)) + v3)))
    delete(v5, 506)
    delete(v2, ("hello" + "a"))
    var v29 float64 = 1e10
    v29 = ((0.0 - float64(0.1)) * v29)
    for fi30 := 0; fi30 < 8; fi30++ {
        v1 = v1
        v31 := GenS5{F0: 1, F1: -1, F2: v4}
        _ = fi30
        _ = v31
    }
    var v32 uint64 = 0
    var v33 float32 = ((0.1 / 1e10) + (v3 + v3))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
