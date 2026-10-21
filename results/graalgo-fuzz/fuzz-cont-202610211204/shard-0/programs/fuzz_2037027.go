package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 uint16) uint8 {
    if false {
        var v5 int = 128
        _ = v5
    }
    if (!(!(42 == 584))) {
        p0 = 42
        var v6 uint32 = 1
        v7 := GenS1{F0: 4294967295, F1: 255, F2: p0}
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    p0 -= (p0 << ((273) & 63))
    var v8 uint8 = 1
    _ = p0
    _ = v8
    return 200
}
func genFn1(p0 string, p1 uint8, p2 int32) bool {
    v9 := []int{}
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int64 = 286
        p1 -= genFn0(65535)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (32767 == 608)
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    v3 := []uint32{953, 2, 2147483647}
    var v4 uint8 = 93
    var v12 float32 = (((0.1 * (-2.25 * (0.0 / 1e10))) * 0.0) / 3.14159)
    fmt.Println(v4, v4, v1)
    v12 += float32(((1.5 + v1) / (0.0 - (v1 / 1e10))))
    v13 := GenS2{F0: 32767, F1: true}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            switch fi15 {
            case 42:
                if false {
                    if (v4 <= 193) {
                        if false {
                            fmt.Println(v12, v1, fi14)
                            var v16 bool = ((-45 <= 42) || (1 <= 4294967295))
                            v16 = false
                            _ = v16
                        }
                        v12 = float32(1.5)
                    }
                    fmt.Println(fi15, v1)
                }
            case 32768:
                var v17 int32 = 575
                _ = v17
            case -293:
                switch (fi15 % ((fi15) | 1)) {
                case 345:
                    v12 = v12
                case -576:
                    var v18 uint = 4294967295
                    _ = v18
                case 32768:
                    if ((4294967295 <= 42) && true) {
                        if false {
                            var v19 uint16 = 0
                            fmt.Println(v19, v12, v12)
                            _ = v19
                        } else {
                            var v20 uint16 = 621
                            _ = v20
                        }
                        var v21 uint64 = 42
                        _ = v21
                    } else {
                        var v22 int64 = 0
                        fmt.Println(v1, fi14, fi14)
                        _ = v22
                    }
                }
            default:
                if (false && (65535 == 4294967295)) {
                    v1 = v1
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
            }
            var v23 float32 = v12
            _ = fi15
            _ = v23
        }
        var v24 uint64 = 42
        if (!((^(v4)) >= v4)) {
            var v25 uint = 65535
            _ = v25
        } else {
            var v26 float32 = (((3.14159 / ((1.5 - 1.5) - float32(v1))) + (float32(-2.25) * v12)) - (v12 * v12))
            if true {
                var v27 float32 = v26
                _ = v27
            } else {
                for fi28 := 0; fi28 < 11; fi28++ {
                    if (42 == -1) {
                        fmt.Println(v24)
                        v1 = (v1 / v1)
                        var v29 float64 = (1e10 / v1)
                        _ = v29
                    }
                    if true {
                        switch (^(fi14)) {
                        case 9223372036854775807:
                            v30 := GenS1{F0: (^(fi28)), F1: 2, F2: 42}
                            _ = v30
                        }
                        switch (fi28 / ((fi14) | 1)) {
                        case -570:
                            v24 = v24
                        }
                    }
                    _ = fi28
                }
                fmt.Println(fi14, v1)
            }
            _ = v26
        }
        _ = fi14
        _ = v24
    }
    if true {
        if false {
            if false {
                v4 = v4
            }
            var v31 bool = true
            delete(v2, -9223372036854775808)
            _ = v31
        } else {
            var v32 uint8 = genFn0(935)
            _ = v32
        }
    } else {
        var v33 uint64 = uint64(v4)
        v34 := GenS2{F0: 65535, F1: (!(1000 > 42))}
        delete(v2, 127)
        _ = v33
        _ = v34
    }
    v1 = 1e10
    var v35 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
