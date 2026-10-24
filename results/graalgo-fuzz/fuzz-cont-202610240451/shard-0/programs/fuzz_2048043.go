package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }

func main() {
    v1 := []uint32{2}
    var v2 int32 = -983
    v3 := make(map[string]uint32)
    var v4 int8 = 1
    var v5 uint = 255
    if true {
        if false {
            v6 := GenS1{F0: 314, F1: 133, F2: (!false)}
            for fi7 := 0; fi7 < 2; fi7++ {
                v2 = (-(v2))
                fmt.Println(v2, v4, fi7)
                v4 = (^(v4))
                _ = fi7
            }
            var v8 uint = v5
            _ = v6
            _ = v8
        } else {
            fmt.Println(v4, v5, v4)
        }
    } else {
        v5 -= v5
        if (false || (0 >= 2)) {
            var v9 float32 = 0.0
            _ = v9
        }
    }
    if ((v2 ^ v2) >= (v2 / ((v2) | 1))) {
        switch -8 {
        case 9223372036854775807:
            var v10 float64 = ((1e10 / 0.1) - 0.0)
            _ = v10
        default:
            var v11 int16 = 0
            _ = v11
        }
        if (false || (464 < 4294967295)) {
            var v12 int8 = 1
            _ = v12
        }
    } else {
        v2 = v2
    }
    v2 = (^(v2))
    v4 = v4
    fmt.Println(v4, v2, v4)
    fmt.Println(v5, v5, v4)
    v3["\n\t"] = 4294967295
    v5 = (^(v5))
    var v13 string = "foo bar"
    var v14 int64 = 42
    if (!true) {
        switch ((v13 + v13) + (v13 + (v13 + v13))) {
        case "hello":
            switch -2147483648 {
            case -9223372036854775808:
                var v15 uint32 = 0
                _ = v15
            case 128:
                switch 766 {
                case 32767:
                    if false {
                        _ = len(v1)
                        var v16 string = "hello"
                        _ = v16
                    }
                case -9223372036854775808:
                    var v17 uint8 = 127
                    _ = v17
                default:
                    fmt.Println(v4, v2, v5)
                }
            }
        case "foo bar":
            var v18 float64 = -2.25
            _ = v18
        default:
            v5 += (v5 << ((v5) & 63))
        }
    }
    v13 = v13
    var v19 float32 = 1.5
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v5 = 4294967295
            _ = fi21
        }
        v5 += (^(v5))
        for fi22 := 0; fi22 < 5; fi22++ {
            v19 -= v19
            _ = len(v3)
            _ = fi22
        }
        _ = fi20
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint = v5
        v24 = (^(v24))
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v14 += int64(v4)
        var v26 int32 = (v2 / (((v2 + (v2 & int32(fi25)))) | 1))
        _ = fi25
        _ = v26
    }
    v2 -= v2
    fmt.Println(v14, v4, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
