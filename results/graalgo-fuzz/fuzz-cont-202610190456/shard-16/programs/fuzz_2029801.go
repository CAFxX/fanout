package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    v3 := []uint32{}
    v4 := []int{}
    v1 = (v1 * v1)
    if false {
        for fi5 := 0; fi5 < 20; fi5++ {
            v1 = 18446744073709551615
            _ = fi5
        }
        fmt.Println(v1)
        var v6 int16 = 0
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 float64 = -2.25
                var v10 int32 = 2147483647
                var v11 int8 = 43
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v1)
    default:
        var v12 uint8 = 2
        _ = v12
    }
    var v13 int = 808
    v14 := GenS1{F0: 0}
    if true {
        v15 := GenS1{F0: 42}
        if false {
            if true {
                var v16 uint = 4294967295
                if (!(-9223372036854775808 == -9223372036854775808)) {
                    var v17 int16 = 1000
                    v16 = (v16 ^ (v16 - (v16 + (^(v16)))))
                    fmt.Println(v16)
                    _ = v17
                }
                fmt.Println(v16, v1)
                _ = v16
            } else {
                switch "foo bar" {
                case "a":
                    v18 := GenS1{F0: 65535}
                    _ = v18
                case "foo bar":
                    for fi19 := 0; fi19 < 6; fi19++ {
                        switch 27 {
                        case -9223372036854775808:
                            v3 = append(v3, 2147483647)
                        case 228:
                            switch v13 {
                            case 4294967295:
                                fmt.Println(fi19)
                            }
                        }
                        var v20 float32 = (float32(1.5) / (3.14159 + 0.1))
                        _ = fi19
                        _ = v20
                    }
                }
                v21 := GenS2{F0: 79, F1: 255}
                _ = v21
            }
        }
        _ = v15
    }
    v1 = (^(v1))
    var v22 int32 = 1
    fmt.Println(v1, v22)
    _ = len(v2)
    var v23 uint32 = 2147483648
    var v24 int = v13
    var v25 int8 = 96
    if len(v4) > 0 {
        v4[0] = v13
    }
    v23 += (v23 >> ((42) & 63))
    var v26 uint16 = 980
    v27 := GenS2{F0: uint8(v26), F1: (v26 + 2)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
