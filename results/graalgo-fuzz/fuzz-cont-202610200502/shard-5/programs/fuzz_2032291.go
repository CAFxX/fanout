package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 int8 = 77
    v2 += ((float64(((3.14159 - ((3.14159 - 1.5) - float32(v2))) + (3.14159 * 0.0))) - v2) - float64((0.1 - ((float32(-2.25) + (0.0 * 3.14159)) * 1e10))))
    switch ("\n\t" + "a") {
    case "foo bar":
        v6 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v6
    }
    switch "\n\t" {
    case "\n\t":
        var v7 float64 = v2
        _ = v7
    case "hello":
        v8 := GenS1{F0: 255, F1: -9223372036854775808}
        _ = v8
    case "foo bar":
        switch "foo bar" {
        case "a":
            v9 := GenS2{F0: -9223372036854775808, F1: v5}
            _ = v9
        case "":
            fmt.Println(v2, v2)
        case "\n\t":
            switch "foo bar" {
            case "":
                if true {
                    var v10 uint16 = 0
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    v2 = 0.1
                    _ = v10
                } else {
                    fmt.Println(v5)
                    if false {
                        switch "\n\t" {
                        case "a":
                            delete(v1, 127)
                        }
                        v11 := GenS2{F0: 1, F1: (v5 / ((v5) | 1))}
                        switch 0 {
                        case 128:
                            var v12 int16 = 3
                            _ = v12
                        case -398:
                            var v13 int64 = 9223372036854775807
                            _ = v13
                        case -2147483648:
                            v5 = v5
                        }
                        _ = v11
                    }
                    v14 := GenS3{F0: -1, F1: -1, F2: 65535}
                    _ = v14
                }
            case "\n\t":
                v2 -= v2
            }
        default:
            if true {
                if (0 < 127) {
                    var v15 int = 9223372036854775807
                    var v16 int64 = 9223372036854775807
                    if (!true) {
                        var v17 float32 = (0.1 * (float32(1e10) * 0.1))
                        switch v15 {
                        case 4294967295:
                            var v18 uint16 = 787
                            _ = v18
                        }
                        _ = v17
                    } else {
                        v16 += v16
                    }
                    _ = v15
                    _ = v16
                } else {
                    var v19 int32 = 956
                    switch 2147483647 {
                    case 0:
                        var v20 int16 = 0
                        _ = v20
                    case -622:
                        var v21 int32 = (v19 >> ((255) & 63))
                        _ = v21
                    case 4294967295:
                        if false {
                            for fi22 := 0; fi22 < 19; fi22++ {
                                var v23 int16 = 404
                                var v24 uint = 42
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            v5 = v5
                            switch 956 {
                            case 623:
                                fmt.Println(v2, v19)
                            }
                        } else {
                            v25 := GenS1{F0: 2, F1: 1}
                            _ = v25
                        }
                    default:
                        var v26 uint32 = 956
                        _ = v26
                    }
                    _ = v19
                }
                var v27 int8 = v5
                switch -1 {
                case -2147483648:
                    fmt.Println(v27, v2, v27)
                case -1:
                    v2 = (float64((0.1 / 1e10)) + v2)
                case 128:
                    var v28 uint8 = 0
                    _ = v28
                }
                _ = v27
            } else {
                v2 = (v2 + v2)
                v2 = (v2 - v2)
                for fi29 := 0; fi29 < 13; fi29++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v30 uint8 = 255
                        _ = v30
                    case "a":
                        fmt.Println(v5, v2, v2)
                    default:
                        var v31 int = -9223372036854775808
                        _ = v31
                    }
                    _ = fi29
                }
            }
        }
    }
    var v32 int = 669
    for fi33 := 0; fi33 < 3; fi33++ {
        fmt.Println(v2, v2)
        var v34 float64 = v2
        var v35 int64 = 0
        _ = fi33
        _ = v34
        _ = v35
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
