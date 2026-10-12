package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    var v1 float64 = (0.0 - 1e10)
    var v2 int64 = -1
    v3 := []uint32{4294967295}
    var v4 float32 = 0.1
    var v5 string = ("" + "foo bar")
    v6 := make(map[string]uint32)
    v5 = "hello"
    switch -1 {
    case 4294967295:
        v7 := GenS1{F0: 387, F1: -32768, F2: 223}
        _ = v7
    default:
        if (32767 < 42) {
            var v8 string = (v5 + "foo bar")
            switch v5 {
            case "\n\t":
                v8 = (v8 + v5)
            case "a":
                var v9 int16 = 915
                _ = v9
            case "":
                v4 = ((v4 - v4) * (((1.5 + v4) / v4) * v4))
            }
            _ = v8
        }
    }
    var v10 int64 = 42
    if true {
        v11 := GenS1{F0: 1, F1: 0, F2: 18446744073709551615}
        var v12 string = (v5 + "foo bar")
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: 42, F1: 2}
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint64 = 4294967295
        _ = fi14
        _ = v15
    }
    var v16 int64 = (v10 % (((-(v10))) | 1))
    delete(v6, (("hello" + ((v5 + (v5 + v5)) + "hello")) + "hello"))
    var v17 uint16 = 255
    fmt.Println(v5, v1, v10)
    v16 = (v10 % ((42) | 1))
    var v18 uint32 = 4294967295
    if true {
        v5 = ("" + v5)
        switch 128 {
        case -2147483648:
            switch 42 {
            case 42:
                v18 -= v18
            }
        case 32767:
            switch 32767 {
            case -2147483648:
                fmt.Println(v2, v2, v4)
            case -1:
                v19 := GenS2{F0: -1, F1: (v18 + v18)}
                _ = v19
            }
        case 42:
            for fi20 := 0; fi20 < 5; fi20++ {
                _ = len(v6)
                if ((v16 - (-(v2))) > 1) {
                    if (!(127 >= 1)) {
                        fmt.Println(v17)
                        if false {
                            v1 = ((v1 * (v1 + 0.0)) - (v1 / 1.0))
                            var v21 uint16 = v17
                            v21 = 32767
                            _ = v21
                        }
                    }
                } else {
                    v22 := GenS2{F0: 127, F1: v18}
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v4 = v4
                        var v24 int32 = 42
                        switch "a" {
                        case "a":
                            var v25 uint = 65535
                            _ = v25
                        }
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
                _ = fi20
            }
        default:
            v6[v5] = (v18 >> ((791) & 63))
        }
        v26 := GenS1{F0: 2, F1: 32767, F2: 1}
        _ = v26
    } else {
        var v27 int64 = (v2 >> ((uint(v18)) & 63))
        v28 := GenS2{F0: -1, F1: v18}
        _ = v27
        _ = v28
    }
    v1 += v1
    _ = len(v3)
    v17 = (v17 * v17)
    switch 32768 {
    case -237:
        if (!true) {
            fmt.Println(v17, v2, v1)
            v29 := GenS1{F0: 4294967295, F1: -32768, F2: 508}
            _ = v29
        }
    case -636:
        var v30 int = 1
        _ = v30
    default:
        var v31 uint32 = 805
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
