package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0() uint64 {
    var v5 int32 = -526
    var v6 int = -571
    v5 = (v5 >> ((800) & 63))
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint16 = 694
    v2 := make(map[uint32]uint32)
    v3 := []uint32{860}
    v4 := []uint32{413, 65535}
    v7 := GenS1{F0: 1, F1: v1}
    var v8 float32 = 1.5
    switch 973 {
    case 907:
        v8 = v8
    case 130:
        var v9 float64 = 3.14159
        _ = v9
    }
    var v10 uint8 = 137
    if true {
        switch "hello" {
        case "":
            fmt.Println(v1, v8, v1)
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                if false {
                    v12 := GenS2{F0: "hello", F1: 389}
                    if (9223372036854775807 <= 489) {
                        v13 := GenS2{F0: (("a" + "") + (("" + "a") + "")), F1: 42}
                        var v14 uint16 = (^(v1))
                        switch (fi11 >> ((531) & 63)) {
                        case 432:
                            v14 = (v14 / (((^(v14))) | 1))
                        case 0:
                            v15 := GenS2{F0: ("a" + "a"), F1: uint32(v10)}
                            _ = v15
                        case 2147483647:
                            fmt.Println(v1, v8)
                        default:
                            v16 := GenS1{F0: 537, F1: (v1 + (v1 * (v14 | v14)))}
                            _ = v16
                        }
                        _ = v13
                        _ = v14
                    } else {
                        v17 := GenS3{F0: 4294967295}
                        fmt.Println(v8, v10, v8)
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 3; fi18++ {
                        fmt.Println(v8)
                        v19 := GenS1{F0: 4294967295, F1: v1}
                        var v20 float32 = ((float32((1.5 - 0.0)) - 1e10) + (v8 + v8))
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = v12
                }
                fmt.Println(v8, fi11)
                var v21 uint8 = (v10 % (((v10 + (v10 & v10))) | 1))
                _ = fi11
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 5; fi22++ {
            fmt.Println(v1)
            _ = fi22
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        fmt.Println(v1)
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 float32 = v8
            if true {
                fmt.Println(v24, v8)
                v25 := v2[39]
                v26 := GenS3{F0: 810}
                _ = v25
                _ = v26
            } else {
                var v27 uint32 = 559
                if len(v4) > 0 {
                    v4[0] = v27
                }
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    }
    if (1 > -893) {
        _ = len(v2)
    } else {
        switch 4294967295 {
        case 538:
            var v28 float64 = 3.14159
            _ = v28
        case -380:
            switch 1 {
            case 127:
                if false {
                    var v29 bool = ((v8 * (float32(1.5) - 3.14159)) <= (v8 * 3.14159))
                    var v30 int16 = 0
                    _ = v29
                    _ = v30
                }
            case 553:
                var v31 uint32 = 2147483648
                _ = v31
            case -9223372036854775808:
                var v32 int64 = -9223372036854775808
                _ = v32
            }
        }
        delete(v2, 2147483647)
    }
    var v33 int32 = 2147483647
    fmt.Println(v1, v33, v8)
    v34 := GenS4{F0: -176, F1: (v10 | 42)}
    fmt.Println(v33)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v7
    _ = v8
}
